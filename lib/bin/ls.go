package bin

import (
	"fmt"
	"io/fs"
	"strconv"
	"syscall"
	"time"

	"github.com/encse/altnet/lib/altnet"
	"github.com/encse/altnet/lib/io"
	"github.com/encse/altnet/lib/slices"
	"golang.org/x/term"
)

func Ls(files []altnet.FileInfo) error {
	columns, _, err := term.GetSize(int(syscall.Stdin))
	if err != nil {
		return err
	}

	names := slices.Map(files, func(fi altnet.FileInfo) string {
		return fi.Name()
	})

	maxWidth := 0
	for _, name := range names {
		if maxWidth < len(name) {
			maxWidth = len(name)
		}
	}

	lines := slices.Chunk(names, columns/(maxWidth+2))
	fmt.Print(io.Table(lines...))
	return nil
}

func LsWide(files []altnet.FileInfo) error {
	now := time.Now()

	flag := func(m fs.FileMode, mask uint32, ch string) string {
		if uint32(m)&mask != 0 {
			return ch
		} else {
			return "-"
		}
	}

	formatFileMode := func(fileMode fs.FileMode) string {
		return flag(fileMode, 01000, "d") +
			flag(fileMode, 0400, "r") +
			flag(fileMode, 0200, "w") +
			flag(fileMode, 0100, "x") +
			flag(fileMode, 0040, "r") +
			flag(fileMode, 0020, "w") +
			flag(fileMode, 0010, "x") +
			flag(fileMode, 0004, "r") +
			flag(fileMode, 0002, "w") +
			flag(fileMode, 0001, "x")
	}

	formatSize := func(size int64) string {
		return strconv.FormatInt(size, 10)
	}

	formatModTime := func(modTime time.Time) string {
		if now.Year() == modTime.Year() {
			return modTime.Format("Jan 01 15:04")
		} else {
			return modTime.Format("Jan 01  2006")
		}
	}

	lines := make([][]string, 0)

	for _, file := range files {
		lines = append(lines, []string{
			formatFileMode(file.Mode()),
			"1",
			string(file.User()),
			string(file.Group()),
			formatSize(file.Size()),
			formatModTime(file.ModTime()),
			file.Name(),
		})
	}

	fmt.Print(io.Table(lines...))
	return nil
}
