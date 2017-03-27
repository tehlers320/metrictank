package cluster

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestPeersForQuery(t *testing.T) {
	manager := getClusterManager("node1", "test", time.Now(), "http", 6060)
	manager.SetPrimary(true)
	manager.SetPartitions([]int32{1, 2})
	manager.SetReady()
	Convey("when cluster in single mode", t, func() {
		selected := manager.MembersForQuery()
		So(selected, ShouldHaveLength, 1)
		So(selected[0], ShouldResemble, manager.ThisNode())
	})
	thisNode := manager.thisNode()
	manager.Lock()
	Mode = ModeMulti
	manager.members = map[string]Node{
		thisNode.GetName(): thisNode,
		"node2": {
			Name:       "node2",
			Primary:    true,
			Partitions: []int32{1, 2},
			State:      NodeReady,
		},
		"node3": {
			Name:       "node3",
			Primary:    true,
			Partitions: []int32{3, 4},
			State:      NodeReady,
		},
		"node4": {
			Name:       "node4",
			Primary:    true,
			Partitions: []int32{3, 4},
			State:      NodeReady,
		},
	}
	manager.Unlock()
	Convey("when cluster in multi mode", t, func() {
		selected := manager.membersForQuery()
		So(selected, ShouldHaveLength, 2)
		nodeNames := []string{}
		for _, n := range selected {
			nodeNames = append(nodeNames, n.Name)
			if n.Name == manager.thisNode().Name {
				So(n, ShouldResemble, manager.thisNode())
			}
		}

		So(nodeNames, ShouldContain, manager.thisNode().Name)
		Convey("members should be selected randomly with even distribution", func() {
			peerCount := make(map[string]int)
			for i := 0; i < 1000; i++ {
				selected = manager.membersForQuery()
				for _, p := range selected {
					peerCount[p.Name]++
				}
			}
			So(peerCount["node1"], ShouldEqual, 1000)
			So(peerCount["node2"], ShouldEqual, 0)
			So(peerCount["node3"], ShouldNotAlmostEqual, 500)
			So(peerCount["node4"], ShouldNotAlmostEqual, 500)
			for p, count := range peerCount {
				t.Logf("%s: %d", p, count)
			}
		})
	})

}
