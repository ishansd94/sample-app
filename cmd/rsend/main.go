package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gookit/event"
	//"github.com/gookit/cache"
	//"github.com/gookit/goutil/cliutil"
	"github.com/gookit/gcli/v3"
	//"https://github.com/go-co-op/gocron"
	//"https://github.com/emirpasic/gods"
	// "github.com/golang-module/carbon"
	"github.com/StudioSol/async"
	//"github.com/catmullet/go-workers"
)

const (
	FILES_MODIFIED = "file.modified"
	FILES_CREATED = "file.created"
	FILES_DELETED = "file.deleted"
)

func main()  {
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


	app := gcli.NewApp()
	app.Version = "1.0.0"
	app.Desc = "this is a sample cli application"

	ctx := context.Background()

	app.Add(&gcli.Command{
		Name: "test",
		Desc: "this is a description <info>message</> for {$cmd}",
		Aliases: []string{"t"},
		Func: func (cmd *gcli.Command, args []string) error {
			gcli.Print("Firing test event\n")
			_ = async.Run(ctx, func(ctx context.Context) error {
				time.Sleep(time.Second*2)
				gcli.Print("async process...\n")
				time.Sleep(time.Second*2)
				return nil
			})
			event.MustFire(FILES_CREATED, event.M{"foo": "val0", "bar": "val1"})
			return nil
		},
	})

	app.Run(nil)
}