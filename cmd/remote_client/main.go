package main

import (
	"context"

	"github.com/StudioSol/async"
	"github.com/sanity-io/litter"

	//"github.com/gookit/cache"
	//"github.com/gookit/goutil/cliutil"
	"github.com/gookit/gcli/v3"
	//"github.com/catmullet/go-workers"
	//"github.com:gookit/goutil"

	"github.com/ishansd94/sample-app/internal/app/remote_client/cli"
	"github.com/ishansd94/sample-app/internal/app/remote_client/listener"
	"github.com/ishansd94/sample-app/internal/app/remote_client/pkg/constant"
	"github.com/ishansd94/sample-app/internal/app/remote_client/service"
	"github.com/ishansd94/sample-app/internal/app/version"
	"github.com/ishansd94/sample-app/pkg/env"
)

func main()  {

	app := gcli.NewApp()
	app.Version = version.Release
	app.Desc = "Remote CLI"
	app.On(gcli.EvtAppInit, func(data ...interface{}) bool {
		gcli.Print("init...\n")
		return false
	})

	ctx := context.Background()

	err := async.Run(ctx, service.CheckUpdates, service.CheckReachability)
	if err != nil {
		litter.Dump(err)
		gcli.Printf("Error connecting to broker at %s\n", env.Get(constant.REMOTE_BROKER_HOST, constant.REMOTE_BROKER_HOST_DEFAULT))
		return
	}

	listener.Register()

	app.Add(&gcli.Command{
		Name: "info",
		Desc: "Print information about the host system",
		Aliases: []string{"i"},
		Func: cli.Info,
	})

	app.Add(&gcli.Command{
		Name: "send",
		Desc: "Send a command to the receiver",
		Aliases: []string{"s"},
		Func: cli.Send,
		Config: func(cmd *gcli.Command) {
			cmd.AddArg("command", "The command that should be executed in the receiver", true)
		},
	})

	app.Add(&gcli.Command{
		Name: "session",
		Desc: "Manage sessions",
		Aliases: []string{"ss"},
		Subs: []*gcli.Command {
			&gcli.Command{
				Name: "init",
				Desc: "Initialize a session.",
				Aliases: []string{"i"},
				Config: func(cmd *gcli.Command) {
					cmd.AddArg("name", "Name of the session", true)
				},
				Func: cli.SessionInit,
			},
			&gcli.Command{
				Name: "use",
				Desc: "Use the session that is passed.",
				Aliases: []string{"u"},
				Config: func(cmd *gcli.Command) {
					cmd.AddArg("name", "Name of the session", true)
				},
				Func: cli.SessionUse,
			},
			&gcli.Command{
				Name: "show",
				Desc: "Show all available sessions.",
				Aliases: []string{"s"},
				Func: cli.Sessions,
			},
		},

	})

	app.Run(nil)
}