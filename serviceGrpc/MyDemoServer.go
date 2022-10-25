package serviceGrpc

import (
	"context"
	"fmt"
)

type MyDemoServer struct {
	UnimplementedDemoServiceServer
}

func (s *MyDemoServer) UnaryCall(ctx context.Context, req *DemoRequest) (*DemoReply, error) {
	fmt.Println("request:", req.Json)
	return &DemoReply{Message: "Hello " + req.Json}, nil
}