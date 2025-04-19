package room

import (
	"context"
	"pixels-emulator/user"
)

// GetUserRoom provides the related user room, if queuing, transitioning or in-game.
func GetUserRoom(ctx context.Context, rs Store, p *user.Player) (*Room, error) {

	r, err := rs.Records().GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for _, r := range r {

		if r.Queue.Contains(p.Id) {
			return r, nil
		}

		if r.IsTransitioning(p) {
			return r, nil
		}

		if r.IsOnline(p) {
			return r, nil
		}

	}

	return nil, nil

}
