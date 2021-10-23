package remote_client

import (
	"fmt"

	"github.com/gookit/event"

	"github.com/ishansd94/sample-app/internal/app/remote_client/pkg/constant"
)

func Register()  {
	event.On(constant.FILES_CREATED, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		return nil
	}), event.Normal)

	event.On(constant.FILES_MODIFIED, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		return nil
	}), event.Normal)

	event.On(constant.FILES_DELETED, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		return nil
	}), event.Normal)
}
