package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ettle/strcase"
)

func renameFile(oldPath string, isDir bool) (string, error) {
	base := filepath.Base(oldPath)
	dir := filepath.Dir(oldPath)

	if isDir {
		name := strcase.ToKebab(base)
		if name == base {
			return oldPath, nil
		}

		path := filepath.Join(dir, name)
		if err := os.Rename(oldPath, path); err != nil {
			return "", err
		}
		return path, nil
	}

	extension := filepath.Ext(base)
	baseName := strings.TrimSuffix(base, extension)
	name := strcase.ToKebab(baseName) + extension

	if name == base {
		return oldPath, nil
	}

	path := filepath.Join(dir, name)
	if err := os.Rename(oldPath, path); err != nil {
		return "", err
	}

	return path, nil
}

func processDirectory(dirPath string) error {
	log.Info("processing directory")
	entries, err := os.ReadDir(dirPath)

	log.Info(dirPath)

	if err != nil {
		return err
	}

	for _, fileLike := range entries {
		oldPath := filepath.Join(dirPath, fileLike.Name())

		if fileLike.IsDir() {
			newPath, err := renameFile(oldPath, true)
			log.Info("after first rename file")
			handlePossiblyNegatedError(err)

			if err := processDirectory(newPath); err != nil {
				log.Error(err)
			}
		} else {
			if _, err := renameFile(oldPath, false); err != nil {
				log.Info("after second rename file")
				log.Error(err)
			}
		}
	}

	return nil
}

func handlePossiblyNegatedError(err error) {
	if err != nil {
		log.Error(err)
	}
}

/*
 * Any file/directory passed to this package will be renamed as kebab-case.
 */
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Error("Please pass a file/directory name.")
		os.Exit(0)
	}

	for _, arg := range args {
		fileInfo, err := os.Stat(arg)
		handlePossiblyNegatedError(err)

		log.Info(fileInfo.IsDir())

		if fileInfo.IsDir() {
			if _, err := renameFile(arg, true); err != nil {
				log.Info("after third rename file")
				log.Error(err)
				continue
			}
			if err := processDirectory(arg); err != nil {
				log.Info("last process dir")
				log.Error(err)
			}
		} else {
			_, err := renameFile(arg, false)
			log.Info("after last rename file")
			handlePossiblyNegatedError(err)
		}
	}

	log.Info("done.")
}
