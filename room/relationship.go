package room

import (
	"context"
	"gorm.io/gorm"
	"pixels-emulator/core/database"
	"pixels-emulator/core/model"
	"pixels-emulator/role"
)

// Relationship defines possible relationship types
type Relationship int

const (
	Owner       Relationship = iota // Owner defines full control over the room
	Rights                          // Rights defines Nitro-granted permissions over the room.
	Access                          // Access defines granted access
	Guest                           // Guest defines no relationship at all.
	Restriction                     // Restriction defines room banning.
)

const AccessRoomPermissions = "pixels.room.access"
const OwnerRoomPermissions = "pixels.room.master"

// VerifyUserRoomRelationship verifies
func VerifyUserRoomRelationship(ctx context.Context, db *gorm.DB, room model.Room, user model.User) (Relationship, error) {

	if role.HasPermission(user, OwnerRoomPermissions) || room.OwnerID == user.ID {
		return Owner, nil
	}

	q := map[string]interface{}{"room_id": room.ID, "user_id": user.ID}
	pStore := &database.ModelService[model.RoomPermission]{DB: db}
	pRes := <-pStore.FindByQuery(ctx, q)

	if pRes.Error == nil {
		return 4, pRes.Error
	}

	if len(pRes.Data) > 0 {
		return Rights, nil
	}

	if role.HasPermission(user, AccessRoomPermissions) {
		return Access, nil
	}

	if false { // TODO: Room restriction system
		return Restriction, nil
	}

	return Guest, nil

}
