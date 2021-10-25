package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Jan-Ka/pikesies-srv/cmd"
	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	var wg sync.WaitGroup
	wg.Add(1)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	appCtx, appCancel := context.WithCancel(context.Background())

	cmdCtx := context.WithValue(appCtx, cmd.CtxWaitGroupKey{}, &wg)

	cmd.Execute(cmdCtx)

	<-signals
	// suppressing ^C
	fmt.Print("\r")
	appCancel()
	wg.Wait()
}
