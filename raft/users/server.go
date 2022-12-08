package users

import (
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

type Server struct {
	logger 			hclog.Logger
	userStore 		*UserStore
	raft 			*raft.Raft
	boltStore 		*raftboltdb.BoltStore
	transportLayer *raft.NetworkTransport

}


func (serve *Server) Start() error {
	tcpaddress, err := net.ResolveTCPAddr("tcp", "localhost")
	if err != nil {
		return fmt.Errorf("error message %w", err)
	}

	serve.transportLayer, err = raft.NewTCPTransportWithLogger(
		"localhost",
		tcpaddress,
		5,
		time.Second * 10,
		serve.logger.Named("raft"),
	)
	if err != nil{
		return fmt.Errorf("no tcp transport")
	}

	//create snapshot store


	return nil
}

func (serve *Server) getVoters() {}

func (serve *Server) BootstrapCluster()  {
	


}




