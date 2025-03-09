package room

// State represents the access state of a room.
type State int8

const (
	// Open means the room is accessible to everyone.
	Open State = iota

	// Locked means the room is restricted and requires manual approval.
	Locked

	// PasswordProtected means the room requires a password to enter.
	PasswordProtected

	// Invisible means the actual room is not being shown to default users.
	Invisible = 3
)
