package files

import (
	"embed"
	"io"
	"os"
)

//go:embed cache.db3
var local embed.FS

func Create(db3Path string) error {
	if _, err := os.Stat(db3Path); err == nil {
		return nil
	}

	src, err := local.Open("cache.db3")

	if err != nil {
		return err
	}

	dst, err := os.Create(db3Path)

	if err != nil {
		return err
	}

	defer dst.Close()
	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
