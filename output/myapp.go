package output

import (
	"fmt"
	"time"
)

// myapp represents a generated struct
type myapp struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Newmyapp creates a new myapp
func Newmyapp(id string) *myapp {
	return &myapp{
		ID:        id,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Print prints the myapp details
func (s *myapp) Print() {
	fmt.Printf("ID: %s\nCreated: %v\nUpdated: %v\n", 
		s.ID, s.CreatedAt, s.UpdatedAt)
}