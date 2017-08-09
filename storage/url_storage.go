package storage

// URLStorage is an interface that sits in front
// of an implementation for URL storage.
type URLStorage interface {
	Add(identifier, url string) error
	Delete(identifier string) error
	Get(identifier string) (string, error)
}
