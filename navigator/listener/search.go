package listener

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"pixels-emulator/core/event"
	"pixels-emulator/core/server"
	eventNav "pixels-emulator/navigator/event"
)

func ProvideSearch() func(event event.Event) {
	return func(event event.Event) {
		OnNavigatorSearch(event)
	}
}

func OnNavigatorSearch(ev event.Event) {

	var err error
	defer func() {
		if err != nil {
			server.GetServer().Logger().Error("error during authentication handle", zap.Error(err))
		}
	}()

	queryEv, valid := ev.(*eventNav.NavigatorQueryEvent)
	if !valid {
		err = errors.New("event proportioned was not authentication")
		return
	}

	fmt.Println(queryEv.Realm())

}
