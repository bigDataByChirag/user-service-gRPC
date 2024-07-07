package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"testing"
	pb "user-service-gRPC/gen/proto"
)

// TestSignup tests the Signup function of the user service.
func TestSignup(t *testing.T) {
	// Create a new gRPC server and register the user service.
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userService{})

	// Create a signup request.
	req := &pb.SignupRequest{
		Fname:   "John",
		City:    "New York",
		Phone:   1234567890,
		Height:  "5.9",
		Married: false,
	}

	// Call the createUser helper function and assert no error occurred.
	resp, err := createUser(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Fname, resp.Person.Fname)
}

// createUser is a helper function to create a user by calling the Signup method.
func createUser(req *pb.SignupRequest) (*pb.SignupResponse, error) {
	ctx := context.Background()
	svc := &userService{}
	return svc.Signup(ctx, req)
}

// TestSearchUserById tests the SearchUserById function of the user service.
func TestSearchUserById(t *testing.T) {
	// Create a new gRPC server and register the user service.
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userService{})

	// Create a user first.
	req := &pb.SignupRequest{
		Fname:   "John",
		City:    "New York",
		Phone:   1234567890,
		Height:  "5.9",
		Married: false,
	}

	// Call the createUser helper function and assert no error occurred.
	resp, err := createUser(req)
	assert.NoError(t, err)

	// Create a search request using the user ID.
	searchReq := &pb.SearchUserByIdRequest{
		Id: resp.Person.Id,
	}

	// Call the searchUserById helper function and assert no error occurred.
	searchResp, err := searchUserById(searchReq)
	assert.NoError(t, err)
	assert.NotNil(t, searchResp)
	assert.Equal(t, req.Fname, searchResp.Person.Fname)
}

// searchUserById is a helper function to search for a user by ID.
func searchUserById(req *pb.SearchUserByIdRequest) (*pb.SearchUserByIdResponse, error) {
	ctx := context.Background()
	svc := &userService{}
	return svc.SearchUserById(ctx, req)
}

// TestSearchUsersByIds tests the SearchUsersByIds function of the user service.
func TestSearchUsersByIds(t *testing.T) {
	// Create a new gRPC server and register the user service.
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userService{})

	// Create users first.
	req1 := &pb.SignupRequest{
		Fname:   "John",
		City:    "New York",
		Phone:   1234567890,
		Height:  "5.9",
		Married: false,
	}
	resp1, err := createUser(req1)
	assert.NoError(t, err)

	req2 := &pb.SignupRequest{
		Fname:   "Jane",
		City:    "Los Angeles",
		Phone:   9876543210,
		Height:  "5.5",
		Married: true,
	}
	resp2, err := createUser(req2)
	assert.NoError(t, err)

	// Create a search request with the user IDs.
	searchReq := &pb.SearchUsersByIdsRequest{
		Ids: []string{resp1.Person.Id, resp2.Person.Id},
	}

	// Call the searchUsersByIds helper function and assert no error occurred.
	searchResp, err := searchUsersByIds(searchReq)
	assert.NoError(t, err)
	assert.Len(t, searchResp.People, 2)
}

// searchUsersByIds is a helper function to search for users by a list of IDs.
func searchUsersByIds(req *pb.SearchUsersByIdsRequest) (*pb.SearchUsersByIdsResponse, error) {
	ctx := context.Background()
	svc := &userService{}
	return svc.SearchUsersByIds(ctx, req)
}

// TestGetAllUserIds tests the GetAllUserIds function of the user service.
func TestGetAllUserIds(t *testing.T) {
	// Create a new gRPC server and register the user service.
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userService{})

	// Create users first.
	req1 := &pb.SignupRequest{
		Fname:   "John",
		City:    "New York",
		Phone:   1234567890,
		Height:  "5.9",
		Married: false,
	}
	_, err := createUser(req1)
	assert.NoError(t, err)

	req2 := &pb.SignupRequest{
		Fname:   "Jane",
		City:    "Los Angeles",
		Phone:   9876543210,
		Height:  "5.5",
		Married: true,
	}
	_, err = createUser(req2)
	assert.NoError(t, err)

	// Create an empty search request.
	searchReq := &emptypb.Empty{}

	// Call the getAllUserIds helper function and assert no error occurred.
	searchResp, err := getAllUserIds(searchReq)
	assert.NoError(t, err)
	assert.Len(t, searchResp.Ids, 6)
}

// getAllUserIds is a helper function to get all user IDs.
func getAllUserIds(req *emptypb.Empty) (*pb.GetAllUserIdsResponse, error) {
	ctx := context.Background()
	svc := &userService{}
	return svc.GetAllUserIds(ctx, req)
}

// TestSearchUsers tests the SearchUsers function of the user service.
func TestSearchUsers(t *testing.T) {
	// Create a new gRPC server and register the user service.
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userService{})

	// Create users first.
	req1 := &pb.SignupRequest{
		Fname:   "John",
		City:    "New York",
		Phone:   1234567890,
		Height:  "5.9",
		Married: false,
	}
	_, err := createUser(req1)
	assert.NoError(t, err)

	req2 := &pb.SignupRequest{
		Fname:   "Jane",
		City:    "Los Angeles",
		Phone:   9876543210,
		Height:  "5.5",
		Married: true,
	}
	_, err = createUser(req2)
	assert.NoError(t, err)

	// Create a search request with specific criteria.
	searchReq := &pb.SearchUsersRequest{
		Fname:   "Jane",
		City:    "Los Angeles",
		Phone:   9876543210,
		Height:  "5.5",
		Married: true,
	}

	// Call the searchUsers helper function and assert no error occurred.
	searchResp, err := searchUsers(searchReq)
	assert.NoError(t, err)
	assert.Len(t, searchResp.People, 3)
	assert.Equal(t, req2.Fname, searchResp.People[0].Fname)
}

// searchUsers is a helper function to search for users based on specific criteria.
func searchUsers(req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	ctx := context.Background()
	svc := &userService{}
	return svc.SearchUsers(ctx, req)
}
