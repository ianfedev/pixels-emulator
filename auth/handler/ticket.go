package handler

import (
	"errors"
	"go.uber.org/zap"
	grant "pixels-emulator/auth/event"
	"pixels-emulator/auth/message"
	"pixels-emulator/core/config"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"pixels-emulator/core/server"
	"strconv"
)

// AuthTicketHandler handles the authentication of tickets for users.
type AuthTicketHandler struct {
	logger    *zap.Logger                           // logger Logger instance for logging
	ssoSvc    database.DataService[model.SSOTicket] // ssoSvc Service for handling SSO tickets
	userSvc   database.DataService[model.User]      // userSvc Service for managing user data
	connStore protocol.ConnectionManager            // connStore Connection store for managing connections
	em        event.Manager                         // em Event manager for firing events
	cfg       *config.Config                        // cfg Configuration for server settings
}

// Handle processes the provided authentication ticket packet.
// This make security checks to validate the ticket handling or enabling development mode.
// Also, when SSO validation is successful, should broadcast a structured event.
func (h *AuthTicketHandler) Handle(packet protocol.Packet, conn protocol.Connection) {

	pack, ok := packet.(*message.AuthTicketPacket)
	if !ok {
		h.logger.Error("cannot cast packet, skipping processing")
		return
	}
	hLog := h.logger.With(zap.String("ticket", pack.Ticket))

	var closeConn error
	defer func() {
		if closeConn != nil {
			hLog.Warn("Error while authenticating SSO ticket", zap.Error(closeConn))
			if err := conn.Dispose(); err != nil {
				hLog.Error("Error while disposing connection")
			}
		}
	}()

	var assignedUser uint
	debug := h.cfg.Server.Environment == "DEVELOPMENT"

	if debug {
		uVal, err := strconv.ParseUint(pack.Ticket, 10, 32)
		if err != nil {
			closeConn = err
		}
		assignedUser = uint(uVal)
		hLog.Warn("Attempting to log in debug mode, SSO ticket will be taken as user id, switch to production to prevent this")
	} else {

		q := map[string]interface{}{"ticket": pack.Ticket}
		res := <-h.ssoSvc.FindByQuery(q)
		ssoRes, err := res.Entities, res.Error

		if err != nil {
			closeConn = err
			return
		}

		if len(ssoRes) < 1 {
			closeConn = errors.New("session is being created with not valid ticket")
			return
		} else if len(ssoRes) > 1 {
			closeConn = errors.New("session is being duplicated")
			return
		}

		assignedUser = ssoRes[0].UserID

	}

	userRes := <-h.userSvc.Get(assignedUser)

	if userRes.Error != nil {
		closeConn = userRes.Error
		return
	}

	id := strconv.Itoa(int(userRes.Entity.ID))
	conn.GrantIdentifier(id)
	h.logger.Debug("Connection upgraded", zap.String("identifier", conn.Identifier()))
	ev := grant.NewEvent(int(userRes.Entity.ID), 0, make(map[string]string))

	err := h.em.Fire(grant.AuthGrantEventName, ev)
	if err != nil {
		closeConn = userRes.Error
		return
	}

}

// NewAuthTicket creates a new AuthTicketHandler with necessary dependencies.
func NewAuthTicket() registry.Handler[protocol.Packet] {

	sv := server.GetServer()
	db := sv.Database()

	return &AuthTicketHandler{
		logger:    sv.Logger(),
		ssoSvc:    &database.ModelService[model.SSOTicket]{DB: db},
		userSvc:   &database.ModelService[model.User]{DB: db},
		connStore: sv.ConnStore(),
		em:        sv.EventManager(),
		cfg:       sv.Config(),
	}
}
