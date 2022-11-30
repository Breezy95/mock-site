package raft_nodes

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	rdb "github.com/hashicorp/raft-boltdb"
)

type raftNode struct {
	transport *raft.Transport
	store     *rdb.BoltStore
	raft      *raft.Raft
	log       hclog.Logger
	name      string
	dir       string
	fsm       *raft.FSM
}

func newRaftNode(logger hclog.Logger, transp *raft.Transport, nodeLst []string, name string) (*raftNode, error) {

}
