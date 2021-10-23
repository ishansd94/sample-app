package listener

import (
	"github.com/gookit/event"
	"github.com/gookit/gcli/v3"
	"github.com/sanity-io/litter"

	"github.com/ishansd94/sample-app/internal/app/remote_client/pkg/constant"
)

func Register()  {
	event.On(constant.COMMAND_SEND, event.ListenerFunc(func(e event.Event) error {
		gcli.Print("Received command..\n")
		litter.Dump(e.Get("command"))
		return nil
	}), event.Normal)
}