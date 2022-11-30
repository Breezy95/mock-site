package main

import (
	"fmt"

	"github.com/hashicorp/raft"
)

func getDefaultRaftConfig() *raft.Config {
	return raft.DefaultConfig()
}

type RaftConfig struct {
	Bootstrap          bool
	NodeIdentifier     string
	StorageLocation    string
	BindAddr           string
	BindPort           int
	FiniteStateMachine raft.FSM
}

func createDefaultRaftNode() *raft.Raft {

	//r_conf := getDefaultRaftConfig()

	/* when creating a node with NewRaft()
	func(conf *raft.Config, fsm raft.FSM, logs raft.LogStore, stable raft.StableStore,
		 snaps raft.SnapshotStore, trans raft.Transport) (*raft.Raft, error)*/
	return &raft.Raft{}
}

func main() {
	r_conf := getDefaultRaftConfig()

	fmt.Println()
}
