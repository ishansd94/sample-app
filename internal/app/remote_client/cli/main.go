package cli

import (
	"os"
	"runtime"

	"github.com/gookit/event"
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/show"

	"github.com/ishansd94/sample-app/internal/app/remote_client/pkg/constant"
)

func Info(cmd *gcli.Command, args []string) error {
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
}

func Send(cmd *gcli.Command, args []string) error {
	command := cmd.Arg("command")
	commandValue := command.GetValue()
	gcli.Print("Sending command to receiver\n")
	event.MustFire(constant.COMMAND_SEND, event.M{"command": commandValue})
	return nil
}

func SessionInit(cmd *gcli.Command, args []string) error {
	command := cmd.Arg("command")
	commandValue := command.GetValue()
	gcli.Print("Sending command to receiver\n")
	event.MustFire(constant.COMMAND_SEND, event.M{"command": commandValue})
	return nil
}

func SessionUse(cmd *gcli.Command, args []string) error {
	command := cmd.Arg("command")
	commandValue := command.GetValue()
	gcli.Print("Sending command to receiver\n")
	event.MustFire(constant.COMMAND_SEND, event.M{"command": commandValue})
	return nil
}

func Sessions(cmd *gcli.Command, args []string) error {
	command := cmd.Arg("command")
	commandValue := command.GetValue()
	gcli.Print("Sending command to receiver\n")
	event.MustFire(constant.COMMAND_SEND, event.M{"command": commandValue})
	return nil
}
