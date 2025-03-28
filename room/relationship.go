package room

import (
	"context"
	"pixels-emulator/core/database"
	"pixels-emulator/core/model"
	"pixels-emulator/core/server"
	"pixels-emulator/role"
)

// Relationship defines possible relationship types
type Relationship int

const (
	OWNER       Relationship = iota // OWNER defines full control over the room
	RIGHTS                          // RIGHTS defines Nitro-granted permissions over the room.
	ACCESS                          // ACCESS defines granted access
	GUEST                           // GUEST defines no relationship at all.
	RESTRICTION                     // RESTRICTION defines room banning.
)

const AccessRoomPermissions = "pixels.room.access"
const OwnerRoomPermissions = "pixels.room.master"

// VerifyUserRoomRelationship verifies
func VerifyUserRoomRelationship(ctx context.Context, room model.Room, user model.User) (Relationship, error) {

	if role.HasPermission(user, OwnerRoomPermissions) || room.OwnerID == user.ID {
		return OWNER, nil
	}

	q := map[string]interface{}{"room_id": room.ID, "user_id": user.ID}
	pStore := &database.ModelService[model.RoomPermission]{DB: server.GetServer().Database()}
	pRes := <-pStore.FindByQuery(ctx, q)

	if pRes.Error == nil {
		return 4, pRes.Error
	}

	if len(pRes.Data) > 0 {
		return RIGHTS, nil
	}

	if role.HasPermission(user, AccessRoomPermissions) {
		return ACCESS, nil
	}

	if false { // TODO: Room restriction system
		return RESTRICTION, nil
	}

	return GUEST, nil

}
