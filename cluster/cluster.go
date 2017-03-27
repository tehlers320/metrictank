package cluster

import (
	"time"

	"github.com/hashicorp/memberlist"
)

type ModeType string

const (
	ModeSingle = "single"
	ModeMulti  = "multi"
)

func validMode(m string) bool {
	if ModeType(m) == ModeSingle || ModeType(m) == ModeMulti {
		return true
	}
	return false
}

var (
	Mode    ModeType
	Manager ClusterManagerIf
	cfg     *memberlist.Config
)

func Init(name, version string, started time.Time, apiScheme string, apiPort int) {
	Manager = getClusterManager(name, version, started, apiScheme, apiPort)
}

func getClusterManager(name, version string, started time.Time, apiScheme string, apiPort int) *ClusterManager {
	manager := &ClusterManager{
		members: map[string]Node{
			name: {
				Name:          name,
				ApiPort:       apiPort,
				ApiScheme:     apiScheme,
				Started:       started,
				Version:       version,
				Primary:       primary,
				PrimaryChange: time.Now(),
				StateChange:   time.Now(),
				Updated:       time.Now(),
				local:         true,
			},
		},
		nodeName: name,
	}
	// initialize our "primary" state metric.
	nodePrimary.Set(primary)
	cfg = memberlist.DefaultLANConfig()
	cfg.BindPort = clusterPort
	cfg.BindAddr = clusterHost.String()
	cfg.AdvertisePort = clusterPort
	cfg.Events = manager
	cfg.Delegate = manager
	return manager
}

type partitionCandidates struct {
	priority int
	nodes    []Node
}
