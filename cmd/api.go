package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log/slog"
	pb "user-service-gRPC/gen/proto"
	"user-service-gRPC/models"
)

// userService struct implements the gRPC UserServiceServer interface.
type userService struct {
	pb.UnimplementedUserServiceServer
}

// Signup handles user signup requests.
func (us *userService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {

	// Create a new user in the models layer.
	p, err := models.CreateUser(req)

	// Log an error and return an appropriate gRPC status if user creation fails.
	if err != nil {
		slog.Error("creating user failed", slog.String("Error", err.Error()))
		return nil, status.Error(codes.Internal, "please send valid user details")
	}

	// Populate the response with the created user details.
	resPerson := pb.Person{
		Id:      p.ID,
		Fname:   p.FName,
		City:    p.City,
		Phone:   p.Phone,
		Height:  p.Height,
		Married: p.Married,
	}

	// Return the response.
	return &pb.SignupResponse{Person: &resPerson}, nil
}

// SearchUserById handles fetching user details by ID.
func (us *userService) SearchUserById(ctx context.Context, req *pb.SearchUserByIdRequest) (*pb.SearchUserByIdResponse, error) {
	// Check if the user exists in the models layer.
	user, exists := models.User[req.Id]
	if !exists {
		// Log an error and return a NotFound status if the user doesn't exist.
		slog.Error("user not found", slog.Bool("exists", exists))
		return nil, status.Errorf(codes.NotFound, "user with ID %s not found", req.Id)
	}

	// Populate the response with the user details.
	resPerson := pb.Person{
		Id:      user.ID,
		Fname:   user.FName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}

	// Return the response.
	return &pb.SearchUserByIdResponse{Person: &resPerson}, nil
}

// SearchUsersByIds handles fetching user details by a list of IDs.
func (us *userService) SearchUsersByIds(ctx context.Context, req *pb.SearchUsersByIdsRequest) (*pb.SearchUsersByIdsResponse, error) {
	var people []*pb.Person

	// Iterate over the provided IDs and fetch user details.
	for _, id := range req.Ids {
		user, exists := models.User[id]
		if exists {
			// Populate a person object with the user details.
			person := pb.Person{
				Id:      user.ID,
				Fname:   user.FName,
				City:    user.City,
				Phone:   user.Phone,
				Height:  user.Height,
				Married: user.Married,
			}
			// Add the person object to the list of people.
			people = append(people, &person)
		}
	}

	// Return a NotFound status if no users were found.
	if len(people) == 0 {
		return nil, status.Error(codes.NotFound, "no users found with provided IDs")
	}

	// Return the response.
	return &pb.SearchUsersByIdsResponse{People: people}, nil
}

// GetAllUserIds handles fetching all user IDs.
func (us *userService) GetAllUserIds(ctx context.Context, req *emptypb.Empty) (*pb.GetAllUserIdsResponse, error) {
	var ids []string

	// Collect all user IDs from the models layer.
	for id := range models.User {
		ids = append(ids, id)
	}

	// Return the response with the list of IDs.
	return &pb.GetAllUserIdsResponse{Ids: ids}, nil
}

// SearchUsers handles searching for users based on specific criteria.
func (us *userService) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	var people []*pb.Person

	// Iterate over all users in the models layer.
	for _, user := range models.User {
		// Check if the user matches the search criteria.
		if req.Fname != "" && user.FName != req.Fname {
			continue
		}
		if req.City != "" && user.City != req.City {
			continue
		}
		if req.Phone != 0 && user.Phone != req.Phone {
			continue
		}
		if req.Height != "" && user.Height != req.Height {
			continue
		}
		if req.Married && !user.Married {
			continue
		}

		// Add the user to the result if all criteria are met.
		person := pb.Person{
			Id:      user.ID,
			Fname:   user.FName,
			City:    user.City,
			Phone:   user.Phone,
			Height:  user.Height,
			Married: user.Married,
		}
		people = append(people, &person)
	}

	// Return the response with the list of matching users.
	return &pb.SearchUsersResponse{People: people}, nil
}
