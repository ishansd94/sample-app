package main

import (
	"context"
	"os"
	"runtime"

	"github.com/StudioSol/async"
	"github.com/gookit/event"
	"github.com/sanity-io/litter"

	//"github.com/gookit/cache"
	//"github.com/gookit/goutil/cliutil"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/show"

	//"github.com/catmullet/go-workers"
	//"github.com:gookit/goutil"

	"github.com/ishansd94/sample-app/internal/app/remote_client"
	"github.com/ishansd94/sample-app/internal/app/version"
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

	fn1 := func(ctx context.Context) error {
		gcli.Print("Check for updates\n")
		//time.Sleep(time.Second*1)
		return nil
	}

	fn2 := func(ctx context.Context) error {
		gcli.Print("Check for broker reachability\n")
		//time.Sleep(time.Second*2)
		return nil
	}

	_ = async.Run(ctx, fn1, fn2)

	event.On(remote_client.COMMAND_SEND, event.ListenerFunc(func(e event.Event) error {
		gcli.Print("Received command..\n")
		litter.Dump(e.Get("command"))
		return nil
	}), event.Normal)

	app.Add(&gcli.Command{
		Name: "info",
		Desc: "Print information about the host system",
		Aliases: []string{"i"},
		Func: func (cmd *gcli.Command, args []string) error {
			exe, _ := os.Executable()
			info := map[string]interface{}{
				"os":       runtime.GOOS,
				"binName":  cmd.BinName(),
				"workDir":  cmd.WorkDir(),
				"rawArgs":  os.Args,
				"execAble": exe,
				"env":      os.Environ(),
			}
			show.JSON(&info)
			return nil
		},
	})

	app.Add(&gcli.Command{
		Name: "send",
		Desc: "Send a command to the receiver",
		Aliases: []string{"s"},
		Func: func (cmd *gcli.Command, args []string) error {
			command := cmd.Arg("command")
			commandValue := command.GetValue()
			gcli.Print("Sending command to receiver\n")
			event.MustFire(remote_client.COMMAND_SEND, event.M{"command": commandValue})
			return nil
		},
		Config: func(cmd *gcli.Command) {
			cmd.AddArg("command", "The command that should be executed in the receiver", true)
		},
	})

	app.Run(nil)
}