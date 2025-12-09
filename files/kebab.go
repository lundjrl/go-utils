package main

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/ettle/strcase"
)

func processFile(arg string) (string, error) {
	extension := filepath.Ext(arg)
	log.Info(extension)
	name := strcase.ToKebab(arg)

	err := os.Rename(arg, name)

	if err != nil {
		log.Error(err)
	}

	return name, err
}

func processDirectory(directory []os.DirEntry) {
	for _, fileLike := range directory {
		isDir := fileLike.Type().IsDir()

		if isDir {
			processDirectory(fileLike.Name())
		} else {
			processFile(fileLike.Name())
		}
	}
}

/*
 * Any file/directory passed to this package will be renamed as kebab-case.
 */
func main() {
	// TODO: Read file names from argument (file or directory should be passed)

	args := os.Args[1:]

	for _, arg := range args {
		directory, err := os.ReadDir(arg)

		if err != nil {
			var name = ""
			_, err = os.ReadFile(arg)
			name, err = processFile(arg)
			msg := arg + " moved to " + name
			log.Info(msg)
		}

		processDirectory(directory)
	}

	log.Info("done.")
}
