package bin

import (
	"context"

	"github.com/encse/altnet/lib/altnet"
)

func More(ctx context.Context, fi altnet.FileInfo) error {
	altnet.RunHiddenCommandWithStdErrRedirectedToStdout(ctx, "/bin/more", fi.RealPath())
	return nil
}
