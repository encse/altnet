package csokavar

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
)

func RunHiddenCommand(ctx context.Context, name string, arg ...string) {
	runAs(ctx, true, name, arg...)
}

func RunCommand(ctx context.Context, name string, arg ...string) {
	runAs(ctx, false, name, arg...)
}

func runAs(ctx context.Context, hidden bool, name string, arg ...string) {
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
	if hidden {
		cmd.Env = append(cmd.Env, fmt.Sprintf("ALTNET_EXE="))
	} else {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", "ALTNET_EXE", path.Base(name)))
	}
	err := cmd.Run()
	if err != nil {
		log.Error(err)
	}

	io.Freshline()
}
