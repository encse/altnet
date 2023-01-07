package csokavar

import (
	"context"
	"os"
	"os/exec"
	"os/signal"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/log"
)

func RunCommand(ctx context.Context, name string, arg ...string) {
	log.Info("run", name, arg)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	defer signal.Stop(c)
	go func() {
		for range c {
			// pass
		}
	}()
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = altnet.ContextToEnv(os.Environ(), ctx)
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}
}
