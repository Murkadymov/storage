package file

import (
	"fmt"
	"github.com/google/uuid"
)

type File struct {
	ID   uuid.UUID
	Name string
	Data []byte
}

func NewFile(filename string, blob []byte) (*File, error) {

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, fmt.Errorf("Got error generating uuid: %w", err)
	}

	return &File{
		ID:   id,
		Name: filename,
		Data: blob,
	}, nil
}
