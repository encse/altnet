package altnet

import (
	"context"
	"fmt"
	stdio "io"
	"os"
	"os/exec"
	"os/signal"
	"path"

	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/log"
)

func RunHiddenCommandWithStdErrRedirectedToStdout(ctx context.Context, name string, arg ...string) {
	runAs(ctx, os.Stdin, os.Stdout, os.Stdout, true, name, arg...)
}

func RunHiddenCommand(ctx context.Context, name string, arg ...string) {
	runAs(ctx, os.Stdin, os.Stdout, os.Stderr, true, name, arg...)
}

func RunCommand(ctx context.Context, name string, arg ...string) {
	runAs(ctx, os.Stdin, os.Stdout, os.Stderr, false, name, arg...)
}

func runAs(
	ctx context.Context,
	stdin stdio.Reader,
	stdout stdio.Writer,
	stderr stdio.Writer,
	hidden bool,
	name string,
	arg ...string,
) {
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
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Env = ContextToEnv(os.Environ(), ctx)
	if hidden {
		cmd.Env = append(cmd.Env, fmt.Sprintf("ALTNET_EXE="))
	} else {
		cmd.Env = append(cmd.Env, fmt.Sprintf("%s=%s", "ALTNET_EXE", path.Base(name)))
	}
	err := io.RunWithSavedTerminalState(cmd.Run)
	if err != nil {
		log.Error(err)
	}
	io.Freshline()

}
