package cluster

import (
	"time"
)

type MockNode struct {
	isLocal      bool
	name         string
	postResponse []byte
}

func (n *MockNode) IsLocal() bool {
	return n.isLocal
}

func (n *MockNode) Post(contentType string, body interface{}) ([]byte, error) {
	return n.postResponse, nil
}

func (n *MockNode) GetName() string {
	return n.name
}

func NewMockNode(isLocal bool, name string, postResponse []byte) *MockNode {
	return &MockNode{
		isLocal:      isLocal,
		name:         name,
		postResponse: postResponse,
	}
}

type MockClusterManager struct {
	Peers           []*MockNode
	membersForQuery []*MockNode
	thisNode        int
	isPrimary       bool
	isReady         bool
	partitions      []int32
}

func (c *MockClusterManager) MemberList() []NodeIf {
	return mockToIf(c.Peers)
}

func (c *MockClusterManager) ThisNode() NodeIf {
	return c.Peers[c.thisNode]
}

func (c *MockClusterManager) Start()                     {}
func (c *MockClusterManager) Stop()                      {}
func (c *MockClusterManager) SetPriority(prio int)       {}
func (c *MockClusterManager) SetPrimary(primary bool)    {}
func (c *MockClusterManager) SetReady()                  {}
func (c *MockClusterManager) SetReadyIn(t time.Duration) {}

func (c *MockClusterManager) IsPrimary() bool {
	return c.isPrimary
}

func (c *MockClusterManager) IsReady() bool {
	return c.isReady
}

func (c *MockClusterManager) MembersForQuery() []NodeIf {
	return mockToIf(c.membersForQuery)
}

func (c *MockClusterManager) GetPartitions() []int32 {
	return c.partitions
}

func (c *MockClusterManager) SetPartitions(partitions []int32) {
	c.partitions = partitions
}

func InitMock() *MockClusterManager {
	manager := &MockClusterManager{}
	Manager = manager
	return manager
}

func mockToIf(in []*MockNode) []NodeIf {
	out := make([]NodeIf, len(in))
	for i, m := range in {
		out[i] = m
	}
	return out
}
