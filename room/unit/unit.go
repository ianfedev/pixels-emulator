package unit

import "pixels-emulator/room/path"

type Unit struct {
	Id         int32
	Status     map[Status]string // Status defines the actual movement control status.
	head, body path.Direction    // head, body defines the corporal rotation.
	Current    path.Coordinate
	Request    path.Request
}

func (u *Unit) GetCurrentTile(l *path.Layout) *path.Tile {
	return l.GetTile(int(u.Current.X()), int(u.Current.Y()))
}

// SetRotation updates the head and body rotation.
func (u *Unit) SetRotation(head path.Direction, body path.Direction) {
	u.head = head
	u.body = body
}

// Rotation provides the head and body rotation.
func (u *Unit) Rotation() (path.Direction, path.Direction) {
	return u.head, u.body
}

func NewUnit(id int32) *Unit {
	return &Unit{
		Id:     id,
		Status: make(map[Status]string),
		head:   0,
		body:   0,
	}
}
