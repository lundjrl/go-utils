package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ettle/strcase"
)

func processFile(arg string) (string, error) {
	extension := filepath.Ext(arg)
	temp := strings.SplitAfter(arg, ".")

	log.Info(temp)
	log.Info(extension)
	name := strcase.ToKebab(temp[0])

	filename := name + extension
	err := os.Rename(arg, filename)

	if err != nil {
		log.Error(err)
	}

	return filename, err
}

func processDirectory(directory []os.DirEntry) {
	for _, fileLike := range directory {
		isDir := fileLike.Type().IsDir()

		if isDir {
			log.Info(fileLike.Info())
			// processDirectory(fileLike.Name())
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

			if err != nil {
				log.Error(err)
				os.Exit(1)
			}

			name, err = processFile(arg)

			if err != nil {
				log.Error(err)
				os.Exit(1)
			}

			msg := arg + " moved to " + name
			log.Info(msg)
		}

		processDirectory(directory)
	}

	log.Info("done.")
}
