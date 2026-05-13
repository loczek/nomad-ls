package store

import "errors"

type Store struct {
	files map[string]*Document
}

func NewStore() Store {
	return Store{
		files: make(map[string]*Document),
	}
}

func (s *Store) GetFile(path string) (*Document, error) {
	if file, ok := s.files[path]; ok {
		return file, nil
	}

	return nil, errors.New("file not found in store")
}

func (s *Store) AddFile(path string, content *Document) *Document {
	s.files[path] = content
	return s.files[path]
}

func (s *Store) RemoveFile(path string) {
	delete(s.files, path)
}

func (s *Store) Contains(path string) bool {
	if _, ok := s.files[path]; ok {
		return true
	}

	return false
}

func (s *Store) Files() map[string]*Document {
	return s.files
}
