package models

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	pb "user-service-gRPC/gen/proto"
)

// CreateUser creates a new user based on the provided SignupRequest.
func CreateUser(request *pb.SignupRequest) (*Person, error) {
	// Generate a new unique ID for the user.
	id := uuid.NewString()

	// Create a new Person object with the details from the request.
	p := &Person{
		ID:      id,
		FName:   request.Fname,
		City:    request.City,
		Phone:   request.Phone,
		Height:  request.Height,
		Married: request.Married,
	}

	// Validate the Person object using the validator package.
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		// Return an error if validation fails.
		return nil, err
	}

	// Store the new user in the User map.
	User[p.ID] = p

	// Return the created Person object.
	return p, nil
}
