package server

import (
	"JFFun/config"
	"JFFun/module"
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func Run(modules ...module.Module) error {
	config.ParseServerConfig()

	for i := 0; i < len(modules); i++ {
		if err := modules[i].Init(); err != nil {
			return err
		}
	}

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for _, mod := range modules {
		wg.Add(1)
		goRunModule(ctx, &wg, mod)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
	cancel()

	wg.Wait()
	return nil
}

func goRunModule(ctx context.Context, wg *sync.WaitGroup, mod module.Module) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
		mod.Run(ctx)
	}()
}
