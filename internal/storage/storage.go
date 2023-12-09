package storage

type Storage struct {
	UrlStore map[string]string
}

func New() *Storage {
	return &Storage{
		UrlStore: make(map[string]string),
	}
}
