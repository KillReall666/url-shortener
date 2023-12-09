package storage

type Storage struct {
	URLStore map[string]string
}

func New() *Storage {
	return &Storage{
		URLStore: make(map[string]string),
	}
}
