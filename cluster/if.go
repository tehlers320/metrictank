package cluster

import (
	"time"
)

type NodeIf interface {
	IsLocal() bool
	Post(string, interface{}) ([]byte, error)
	GetName() string
}

type ClusterManagerIf interface {
	MemberList() []NodeIf
	ThisNode() NodeIf
	Start()
	Stop()
	SetPriority(int)
	SetPrimary(bool)
	IsPrimary() bool
	SetReady()
	SetReadyIn(time.Duration)
	IsReady() bool
	MembersForQuery() []NodeIf
	GetPartitions() []int32
	SetPartitions([]int32)
}
