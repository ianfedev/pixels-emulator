package util

// Disposable represents an interface which can
// dynamically free unused resources if needed.
type Disposable interface {
	// Dispose free the resources or closes connections.
	Dispose() error
}
