package grpc

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	pb "github.com/hotttao/goalgo/grpc/golang/api/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func GetClinet() (pb.GreeterClient, error) {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// defer conn.Close()
	c := pb.NewGreeterClient(conn)
	return c, nil
}

func SayHello(c pb.GreeterClient) {
	log.Printf("run SayHello")
	// SayHello
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.HelloRequest{Name: *name}
	r, err := c.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Reciver: %s", r.GetMessage())
}

func SayHelloReplyStream(c pb.GreeterClient) {
	log.Printf("run SayHelloReplyStream")
	// SayHelloReplyStream
	req := &pb.HelloRequest{Name: *name}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.SayHelloReplyStream(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.SayHelloReplyStream(_) = _, %v", c, err)
		}
		log.Printf("Receiver: %s", reply.GetMessage())
	}
}

func SayHelloRequestStream(c pb.GreeterClient) {
	log.Printf("run SayHelloRequestStream")
	// SayHelloRequestStream

	stream, err := c.SayHelloRequestStream(context.Background())
	reqs := []*pb.HelloRequest{
		{Name: "tao", NumGreetings: "1"},
		{
			Name:         "he",
			NumGreetings: "2",
		},
	}
	if err != nil {
		log.Fatalf("%v.RecordRoute(_) = _, %v", c, err)
	}
	for _, req := range reqs {
		log.Printf("Send: %v", req)
		if err := stream.Send(req); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, req, err)
		}
	}
	reply, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Receiver: %v", reply)
}

func SayHelloStream(c pb.GreeterClient) {
	log.Printf("run SayHelloStream")
	// SayHelloStream
	waitic := make(chan struct{})
	stream, err := c.SayHelloStream(context.Background())
	if err != nil {
		log.Fatalf("%v.RecordRoute(_) = _, %v", c, err)
	}
	reqs := []*pb.HelloRequest{
		{Name: "tao", NumGreetings: "1"},
		{
			Name:         "he",
			NumGreetings: "2",
		},
	}

	go func() {
		for {
			reply, err := stream.Recv()
			if err == io.EOF {
				close(waitic)
				return
			}
			log.Printf("Response: %v", reply)
			if err != nil {
				log.Fatalf("Failed to receive : %v", err)
			}
		}
	}()

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, req, err)
		}
	}
	stream.CloseSend()
	<-waitic
}

func StartClient() {
	client, err := GetClinet()
	if err != nil {
		log.Fatalf("get client error: %v", err)
	}
	// SayHello(client)
	// SayHelloReplyStream(client)
	// SayHelloRequestStream(client)
	SayHelloStream(client)
}
