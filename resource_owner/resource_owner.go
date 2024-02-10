package resourceowner

import (
	"encoding/json"
	"fmt"
)

type ResourceOwner struct {
	ID       string
	Username string
	Email    string
}

func NewResourceOwner(id, username, email string) *ResourceOwner {
	return &ResourceOwner{
		ID:       id,
		Username: username,
		Email:    email,
	}
}

func (ro *ResourceOwner) String() string {
	return fmt.Sprintf("ID: %s, Username: %s, Email: %s", ro.ID, ro.Username, ro.Email)
}

func (ro *ResourceOwner) CsvString() string {
	return fmt.Sprintf("%s, %s, %s", ro.ID, ro.Username, ro.Email)
}

func (ro *ResourceOwner) PrepareForPost() ([]byte, error) {
	jsonData, err := json.Marshal(ro)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
