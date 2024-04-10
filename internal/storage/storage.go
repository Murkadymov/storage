package storage

import "modstorage/internal/file"

type Storage struct {
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	finalfile, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}

	return finalfile, nil

}
