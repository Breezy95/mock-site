package users

import (
	"fmt"
	"net"
	"path/filepath"
	"time"
	"strconv"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
)

type Server struct {
	logger         hclog.Logger
	userStore      *UserStore
	raft           *raft.Raft
	boltStore      *boltdb.BoltStore
	transportLayer *raft.NetworkTransport
	serveConfig    *ServerConfig
}

type ServerConfig struct {
	Config         raft.Config
	serverID       string
	address        string
	raftTobackport int
	rafttoraftPort int
	storeDir       string
	occupiedStore  bool
}

var (
	raftTime = 30 * time.Millisecond
	maxPool  = 5
)

func (serve *Server) Start() error {
	tcpaddress, err := net.ResolveTCPAddr("tcp", "localhost")
	if err != nil {
		return fmt.Errorf("error message %w", err)
	}

	serve.transportLayer, err = raft.NewTCPTransportWithLogger(
		serve.serveConfig.address,
		tcpaddress,
		maxPool,
		raftTime,
		serve.logger.Named("raftlogger"),
	)

	if err != nil {
		return fmt.Errorf("error in creating raft.TCPTransport with logger: %w", err)
	}

	snapstore, err := raft.NewFileSnapshotStoreWithLogger(
		serve.serveConfig.address,
		3,
		serve.logger.Named("snapstore"),
	)

	var stableStore raft.StableStore
	var logStore raft.LogStore

	if serve.serveConfig.occupiedStore {
		newRaftStore := raft.NewInmemStore()
		stableStore = newRaftStore
		logStore = newRaftStore

	} else {
		db, err := boltdb.NewBoltStore(filepath.Join("/storage"))
		if err != nil {
			return fmt.Errorf("error in creating boltdb boltstore %w", err)
		}

		stableStore = db
		serve.boltStore = db
		logStore, err = raft.NewLogCache(512, db)
		if err != nil {
			return fmt.Errorf("error in creating raft logstore %w", err)
		}
	}

	raft_cfg := raft.DefaultConfig()
	raft_cfg.LocalID = raft.ServerID(serve.serveConfig.serverID)
	raft_cfg.Logger = serve.logger.Named("created raft node")

	serve.raft, err = raft.NewRaft(raft_cfg, serve.userStore, logStore, stableStore, snapstore, serve.transportLayer)
	if err != nil {
		return fmt.Errorf("error in creating new raft struct  %w", err)
	}

	return nil
}

func (serve *Server) getLeaderID() (raft.ServerAddress, raft.ServerID) {
	return serve.raft.LeaderWithID()
}

func (serve *Server) BootstrapCluster() error {
	 	clusterCount,err := strconv.Atoi(serve.clusterCount())
	 	if err != nil{
			return fmt.Errorf("error in retrieving cluster count %w", err)
	 		}
	 	if clusterCount >1 {
			return fmt.Errorf("errorcluster created")
			}
	
		raft_cfg := raft.Configuration{ []raft.Server{{ID: raft.ServerID(serve.serveConfig.serverID) , Address: serve.transportLayer.LocalAddr(), }}}
		bootstrapped := serve.raft.BootstrapCluster(raft_cfg)
		if bootstrapped.Error()!= nil {
		return fmt.Errorf("error in bootstrapping cluster %w", err)
			}

	return nil
}

func (serve *Server) JoinCluster(serverID string, clusterAddress string) error {
	if serve.isLeader(){

	}
}

func (serve *Server) isLeader() bool{
	return serve.raft.State() == raft.Leader
}

func (serve *Server) clusterCount() string {
	return serve.raft.Stats()["num_peers"]
}
