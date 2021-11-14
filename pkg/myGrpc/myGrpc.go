package myGrpc

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
	network string
	address string
	timeout time.Duration
	listen  net.Listener
	ctx     context.Context
}

func NewServer() *Server {
	srv := &Server{
		ctx:     context.Background(),
		network: "tcp",
		address: "8080",
		timeout: 300 * time.Microsecond,
		Server:  grpc.NewServer(),
	}

	// if c.Grpc.Network != "" {
	// 	srv.network = c.Grpc.Network
	// }
	// if c.Grpc.Addr != "" {
	// 	srv.address = c.Grpc.Addr
	// }
	// if c.Grpc.Timeout != nil {
	// 	srv.timeout = c.Grpc.Timeout.AsDuration()
	// }
	return srv
}

func (srv *Server) Listen(lis net.Listener) {
	srv.listen = lis
}

// RegisterHelloService(new(HelloService))

// listener, err := net.Listen("tcp", ":1122")
// if err != nil {
// 	log.Fatal("listen TCP fail:", err)
// }

// for {
// 	conn, err := listener.Accept()
// 	if err != nil {
// 		log.Fatal("Accept fail:", err)
// 	}

// 	fmt.Println("service start")
// 	go rpc.ServeConn(conn)
// }
// return rpc.RegisterName(HelloServiceName, svc)
