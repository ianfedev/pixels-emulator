package sso

import (
	"errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"pixels-emulator/auth/grant"
	"pixels-emulator/core/config"
	"pixels-emulator/core/database"
	"pixels-emulator/core/event"
	"pixels-emulator/core/model"
	"pixels-emulator/core/protocol"
	"pixels-emulator/core/registry"
	"strconv"
)

type Handler struct {
	logger  *zap.Logger
	ssoSvc  *database.ModelService[model.SSOTicket]
	userSvc *database.ModelService[model.User]
	cs      *protocol.ConnectionStore
	em      *event.Manager
	debug   bool
}

func (h Handler) Handle(packet protocol.Packet, conn *protocol.Connection) {

	pack, ok := packet.(*Packet)
	if !ok {
		h.logger.Error("cannot cast packet, skipping processing")
		return
	}
	hLog := h.logger.With(zap.String("ticket", pack.Ticket))

	var closeConn error
	defer func() {
		if closeConn != nil {
			hLog.Warn("Error while authenticating SSO ticket", zap.Error(closeConn))
			if err := (*conn).Dispose(); err != nil {
				hLog.Error("Error while disposing connection")
			}
		}
	}()

	var assignedUser uint

	if h.debug {
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

		if closeConn != nil {
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
	(*conn).GrantIdentifier(id)
	h.logger.Debug("Connection upgraded", zap.String("identifier", (*conn).Identifier()))
	ev := grant.NewEvent(int(userRes.Entity.ID), 0, make(map[string]string))

	err := h.em.Fire(grant.AuthEventName, ev)
	if err != nil {
		closeConn = userRes.Error
		return
	}

}

func NewSSOTicketHandler(
	logger *zap.Logger,
	db *gorm.DB,
	cs *protocol.ConnectionStore,
	cfg *config.Config,
	em *event.Manager) registry.Handler[protocol.Packet] {
	return Handler{
		logger:  logger,
		ssoSvc:  &database.ModelService[model.SSOTicket]{DB: db},
		userSvc: &database.ModelService[model.User]{DB: db},
		cs:      cs,
		em:      em,
		debug:   cfg.Server.Environment == "DEVELOPMENT",
	}
}
