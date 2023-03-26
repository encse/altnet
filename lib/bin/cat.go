package bin

import (
	"context"

	"github.com/encse/altnet/lib/altnet"
)

func Cat(ctx context.Context, fi altnet.FileInfo) error {
	altnet.RunHiddenCommandWithStdErrRedirectedToStdout(ctx, "/bin/cat", fi.RealPath())
	return nil
}
