package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
		pbUser "github.com/Kowiste/boilerPlate/doc/proto/user"
)

const (
    address = "localhost:50051"
)

func main() {
    // Set up a connection to the server
    conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pbUser.NewUserServiceClient(conn)

    // Call GetUserById
    getUserById(client, "1")

    // Call GetAllUsers
    getAllUsers(client)
}

func getUserById(client pbUser.UserServiceClient, id string) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    req := &pbUser.GetByIdRequest{Id: id}
    res, err := client.GetUserById(ctx, req)
    if err != nil {
        log.Fatalf("could not get user by ID: %v", err)
    }
    log.Printf("User: %v", res.User)
}

func getAllUsers(client pbUser.UserServiceClient) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    req := &pbUser.GetAllUsersRequest{}
    res, err := client.GetAllUsers(ctx, req)
    if err != nil {
        log.Fatalf("could not get all users: %v", err)
    }
    log.Printf("Users: %v", res.Users)
}
