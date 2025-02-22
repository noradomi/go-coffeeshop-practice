package shared_kernel

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)

	return id, err
}
