package main

import (
	"context"
	pb "go-grpc/usermgmt"

	// "io/ioutil"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
)

const port = ":50051"

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{user_list: &pb.UserList{}}
}

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

func (server *UserManagementServer) Run() error {
	lis, err := net.Listen("tcp", port)
	// readb,err := ioutil.WriteFile()
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, server)

	log.Printf("Server listening at : %v ", lis.Addr())

	return s.Serve(lis)
}

func (s *UserManagementServer) CreateNewUser(c context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received : %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000))
	created_user := &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}
	s.user_list.Users = append(s.user_list.Users, created_user)
	return created_user, nil
}

func (s *UserManagementServer) GetUsers(c context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.user_list, nil
}
func main() {

	server := NewUserManagementServer()

	if err := server.Run(); err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

}
