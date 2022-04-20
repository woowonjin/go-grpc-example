package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/woowonjin/go-grpc-example/data"
	postpb "github.com/woowonjin/go-grpc-example/protos/v1/post"
	userpb "github.com/woowonjin/go-grpc-example/protos/v1/user"
	client "github.com/woowonjin/go-grpc-example/simple-client-server"
)

const portNumber = "9001"

type postServer struct {
	postpb.PostServer

	userCli userpb.UserClient
}

func (s *postServer) ListPostsByUserId(ctx context.Context, req *postpb.ListPostsByUserIdRequest) (*postpb.ListPostsResponse, error) {
	var postMessages []*postpb.PostMessage
	for _, up := range data.UserPosts {
		resp, err := s.userCli.GetUser(ctx, &userpb.GetUserRequest{UserId: up.UserId})
		if err != nil {
			return nil, err
		}

		for _, p := range up.Posts {
			p.Author = resp.UserMessage.Name
		}
		
		postMessages = append(postMessages, up.Posts)
	}

	return &postpb.ListPostsResponse{
		PostMessages: postMessages,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":"+portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	userCli := client.GetUserClient("localhost:9000")
	grpcServer := grpc.NewServer()
	postpb.RegisterPostServer(grpcServer, &postServer{
		userCli: userCli,
	})

	log.Printf("start gRPC server on %s port", portNumber)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}