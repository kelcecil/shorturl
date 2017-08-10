package storage

// URLStorage is an interface that sits in front
// of an implementation for URL storage.
type URLStorage interface {
	Add(string) (int, error)
	Delete(int) error
	Get(int) (string, error)
}
