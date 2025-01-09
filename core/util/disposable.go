package util

// Disposable represents an interface that defines a method for freeing resources or closing connections when no longer needed.
type Disposable interface {
	// Dispose releases resources or closes connections associated with the object.
	Dispose() error
}
