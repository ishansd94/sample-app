package remote_client

import (
	"fmt"

	"github.com/gookit/event"
)

func Register()  {
	event.On(FILES_CREATED, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		return nil
	}), event.Normal)

	event.On(FILES_MODIFIED, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		return nil
	}), event.Normal)

	event.On(FILES_DELETED, event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		return nil
	}), event.Normal)
}
