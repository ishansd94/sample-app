package service

import (
	"context"

	"github.com/gookit/gcli/v3"
	"github.com/go-ping/ping"
	"github.com/sanity-io/litter"

	"github.com/ishansd94/sample-app/internal/app/remote_client/pkg/constant"
	"github.com/ishansd94/sample-app/pkg/env"
)

func CheckUpdates(ctx context.Context) error {
	gcli.Print("Check for updates\n")
	//time.Sleep(time.Second*1)
	return nil
}

func CheckReachability(ctx context.Context) error {
	gcli.Print("Check for broker reachability\n")
	pinger, err := ping.NewPinger(env.Get(constant.REMOTE_BROKER_HOST, constant.REMOTE_BROKER_HOST_DEFAULT))
	if err != nil {
		return err
	}
	pinger.Count = 1
	err = pinger.Run()
	if err != nil {
		return err
	}
	stats := pinger.Statistics()
	litter.Dump(stats)
	return nil
}
