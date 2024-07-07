package models

// Person represents a user with validation tags.
type Person struct {
	ID      string `validate:"required"` // ID must be present.
	FName   string `validate:"required"` // First name must be present.
	City    string `validate:"required"` // City must be present.
	Phone   uint64 `validate:"required"` // Phone must be present.
	Height  string // Height is optional.
	Married bool   // Married status is optional.
}

// Users is a map of user IDs to Person objects.
type Users map[string]*Person

// User is a global variable that stores all users.
var User Users = make(Users)
