package service

type Storage interface {
	Get(hashID string) (string, error)
	Save(longURL string) string
}
