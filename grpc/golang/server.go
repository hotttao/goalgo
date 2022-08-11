package grpc

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	pb "github.com/hotttao/goalgo/grpc/golang/api/helloworld"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *server) SayHelloReplyStream(in *pb.HelloRequest, stream pb.Greeter_SayHelloReplyStreamServer) error {
	for i := 1; i < 10; i++ {
		stream.Send(&pb.HelloReply{Message: "Hello " + string(in.GetName()) + strconv.Itoa(i)})
	}
	return nil
}

func (s *server) SayHelloRequestStream(stream pb.Greeter_SayHelloRequestStreamServer) error {
	i := 1
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloReply{Message: "RequestStream End " + strconv.Itoa(i)})
		}

		if err != nil {
			return err
		}
		log.Printf("Receive: %v", in)
		i++
	}

}
func (s *server) SayHelloStream(stream pb.Greeter_SayHelloStreamServer) error {
	i := 0
	for {
		i++
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		reply := &pb.HelloReply{Message: "Hello " + in.GetName() + strconv.Itoa(i)}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func StartGrpcServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {

		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
