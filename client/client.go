package main

import (
	"context"
	pb "go-grpc/usermgmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal("Did not connect")
	}
	defer conn.Close()
	c := pb.NewUserManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	var new_users = make(map[string]int32)

	new_users["Alice"] = 33
	new_users["Bob"] = 30

	for name, age := range new_users {
		r, err := c.CreateNewUser(ctx, &pb.NewUser{Name: name, Age: age})

		if err != nil {
			log.Fatalf("Could not create user. %v ", err)
		}

		log.Printf(`User Details:
		Name: %s
		Age: %d
		Id: %d`, r.GetName(), r.GetAge(), r.GetId())

	}

	params := &pb.GetUsersParams{}
	r, err := c.GetUsers(ctx, params)

	if err != nil {
		log.Fatalf("Could not retrieve users : %v", err)
	}
	log.Print("Users List: \n")

	log.Printf("r.GetUsers(): %v\n", r.GetUsers())

}
