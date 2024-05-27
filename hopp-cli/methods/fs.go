package methods

import (
	"fmt"
	"log"

	"github.com/knadh/stuffbin"
)

func initFileSystem(binPath string, templateDir string) (stuffbin.FileSystem, error) {
	fs, err := stuffbin.UnStuff(binPath)
	// If files are not stuffed with the binary,
	// try to pick up files from local file system.
	if err == stuffbin.ErrNoID {
		// Running in local mode. Load the required static assets into
		// the in-memory stuffbin.FileSystem.

		files := []string{
			fmt.Sprintf("%s/index.html:/index.html", templateDir),
			fmt.Sprintf("%s/template.md:/template.md", templateDir),
			fmt.Sprintf("%s/_sidebar.md:/_sidebar.md", templateDir),
		}

		// mutates err object.
		fs, err = stuffbin.NewLocalFS("/", files...)
		if err != nil {
			log.Println("Error in Virtual FS", err)
		}
	}

	// Either unstuff or NewLocalFS throws error,
	// mutated error value will be returned
	return fs, err
}
