package encode

// Door represents the access state of a room.
type Door int8

const (
	// Open means the room is accessible to everyone.
	Open Door = iota

	// Locked means the room is restricted and requires manual approval.
	Locked

	// PasswordProtected means the room requires a password to enter.
	PasswordProtected

	// Invisible means the actual room is not being shown to default users.
	Invisible = 3

	// Noob represents newbie status on Nitro client (INVESTIGATION).
	Noob = 4
)
