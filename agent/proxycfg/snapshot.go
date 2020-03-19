package proxycfg

import (
	"context"

	"github.com/hashicorp/consul/agent/structs"
	"github.com/mitchellh/copystructure"
)

type configSnapshotConnectProxy struct {
	DiscoveryChain           map[string]*structs.CompiledDiscoveryChain // this is keyed by the Upstream.Identifier(), not the chain name
	WatchedUpstreams         map[string]map[string]context.CancelFunc
	WatchedUpstreamEndpoints map[string]map[string]structs.CheckServiceNodes
	WatchedGateways          map[string]map[string]context.CancelFunc
	WatchedGatewayEndpoints  map[string]map[string]structs.CheckServiceNodes
	WatchedServiceChecks     map[structs.ServiceID][]structs.CheckType // TODO: missing garbage collection

	PreparedQueryEndpoints map[string]structs.CheckServiceNodes // DEPRECATED:see:WatchedUpstreamEndpoints
}

func (c *configSnapshotConnectProxy) IsEmpty() bool {
	if c == nil {
		return true
	}
	return len(c.DiscoveryChain) == 0 &&
		len(c.WatchedUpstreams) == 0 &&
		len(c.WatchedUpstreamEndpoints) == 0 &&
		len(c.WatchedGateways) == 0 &&
		len(c.WatchedGatewayEndpoints) == 0 &&
		len(c.WatchedServiceChecks) == 0 &&
		len(c.PreparedQueryEndpoints) == 0
}

type configSnapshotMeshGateway struct {
	// WatchedServices is a map of service id to a cancel function. This cancel
	// function is tied to the watch of connect enabled services for the given
	// id. If the main datacenter services watch would indicate the removal of
	// a service all together we then cancel watching that service for its
	// connect endpoints.
	WatchedServices map[structs.ServiceID]context.CancelFunc

	// WatchedServicesSet indicates that the watch on the datacenters services
	// has completed. Even when there are no connect services, this being set
	// (and the Connect roots being available) will be enough for the config
	// snapshot to be considered valid. In the case of Envoy, this allows it to
	// start its listeners even when no services would be proxied and allow its
	// health check to pass.
	WatchedServicesSet bool

	// WatchedDatacenters is a map of datacenter name to a cancel function.
	// This cancel function is tied to the watch of mesh-gateway services in
	// that datacenter.
	WatchedDatacenters map[string]context.CancelFunc

	// ServiceGroups is a map of service id to the service instances of that
	// service in the local datacenter.
	ServiceGroups map[structs.ServiceID]structs.CheckServiceNodes

	// ServiceResolvers is a map of service id to an associated
	// service-resolver config entry for that service.
	ServiceResolvers map[structs.ServiceID]*structs.ServiceResolverConfigEntry

	// GatewayGroups is a map of datacenter names to services of kind
	// mesh-gateway in that datacenter.
	GatewayGroups map[string]structs.CheckServiceNodes

	// FedStateGateways is a map of datacenter names to mesh gateways in that
	// datacenter.
	FedStateGateways map[string]structs.CheckServiceNodes

	// ConsulServers is the list of consul servers in this datacenter.
	ConsulServers structs.CheckServiceNodes
}

func (c *configSnapshotMeshGateway) Datacenters() []string {
	sz1, sz2 := len(c.GatewayGroups), len(c.FedStateGateways)

	sz := sz1
	if sz2 > sz1 {
		sz = sz2
	}

	dcs := make([]string, 0, sz)
	for dc, _ := range c.GatewayGroups {
		dcs = append(dcs, dc)
	}
	for dc, _ := range c.FedStateGateways {
		if _, ok := c.GatewayGroups[dc]; !ok {
			dcs = append(dcs, dc)
		}
	}
	return dcs
}

func (c *configSnapshotMeshGateway) IsEmpty() bool {
	if c == nil {
		return true
	}
	return len(c.WatchedServices) == 0 &&
		!c.WatchedServicesSet &&
		len(c.WatchedDatacenters) == 0 &&
		len(c.ServiceGroups) == 0 &&
		len(c.ServiceResolvers) == 0 &&
		len(c.GatewayGroups) == 0 &&
		len(c.FedStateGateways) == 0 &&
		len(c.ConsulServers) == 0
}

type configSnapshotIngressGateway struct {
	Config                   *structs.IngressGatewayConfigEntry
	Services                 map[structs.ServiceID]struct{}
	WatchedDiscoveryChains   map[string]context.CancelFunc
	DiscoveryChain           map[string]*structs.CompiledDiscoveryChain
	WatchedUpstreams         map[string]map[string]context.CancelFunc
	WatchedUpstreamEndpoints map[string]map[string]structs.CheckServiceNodes
}

func (c *configSnapshotIngressGateway) IsEmpty() bool {
	if c == nil {
		return true
	}
	return c.Config == nil &&
		len(c.Services) == 0 &&
		len(c.DiscoveryChain) == 0 &&
		len(c.WatchedDiscoveryChains) == 0 &&
		len(c.WatchedUpstreams) == 0 &&
		len(c.WatchedUpstreamEndpoints) == 0
}

// ConfigSnapshot captures all the resulting config needed for a proxy instance.
// It is meant to be point-in-time coherent and is used to deliver the current
// config state to observers who need it to be pushed in (e.g. XDS server).
type ConfigSnapshot struct {
	Kind            structs.ServiceKind
	Service         string
	ProxyID         structs.ServiceID
	Address         string
	Port            int
	ServiceMeta     map[string]string
	TaggedAddresses map[string]structs.ServiceAddress
	Proxy           structs.ConnectProxyConfig
	Datacenter      string

	ServerSNIFn ServerSNIFunc
	Roots       *structs.IndexedCARoots
	// Used for connect-proxy and ingress-gateway proxies
	Leaf *structs.IssuedCert

	// connect-proxy specific
	ConnectProxy configSnapshotConnectProxy

	// mesh-gateway specific
	MeshGateway configSnapshotMeshGateway

	// ingress-gateway specific
	IngressGateway configSnapshotIngressGateway

	// Skip intentions for now as we don't push those down yet, just pre-warm them.
}

// Valid returns whether or not the snapshot has all required fields filled yet.
func (s *ConfigSnapshot) Valid() bool {
	switch s.Kind {
	case structs.ServiceKindConnectProxy:
		return s.Roots != nil && s.Leaf != nil
	case structs.ServiceKindMeshGateway:
		if s.ServiceMeta[structs.MetaWANFederationKey] == "1" {
			if len(s.MeshGateway.ConsulServers) == 0 {
				return false
			}
		}
		return s.Roots != nil && (s.MeshGateway.WatchedServicesSet || len(s.MeshGateway.ServiceGroups) > 0)
	case structs.ServiceKindIngressGateway:
		return s.Roots != nil &&
			s.Leaf != nil &&
			s.IngressGateway.Config != nil &&
			s.IngressGateway.Services != nil
	default:
		return false
	}
}

// Clone makes a deep copy of the snapshot we can send to other goroutines
// without worrying that they will racily read or mutate shared maps etc.
func (s *ConfigSnapshot) Clone() (*ConfigSnapshot, error) {
	snapCopy, err := copystructure.Copy(s)
	if err != nil {
		return nil, err
	}

	snap := snapCopy.(*ConfigSnapshot)

	// nil these out as anything receiving one of these clones does not need them and should never "cancel" our watches
	switch s.Kind {
	case structs.ServiceKindConnectProxy:
		snap.ConnectProxy.WatchedUpstreams = nil
		snap.ConnectProxy.WatchedGateways = nil
	case structs.ServiceKindMeshGateway:
		snap.MeshGateway.WatchedDatacenters = nil
		snap.MeshGateway.WatchedServices = nil
	}

	return snap, nil
}
