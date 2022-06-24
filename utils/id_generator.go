package utils

import "github.com/google/uuid"

func GenerateId() string {
	id := uuid.New()
	return id.String()
}
