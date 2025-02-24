package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func executeMakefile(filePath string, arg ...string) error {
	cmd := exec.Command(filePath, arg...)
	cmd.Dir = filepath.Dir(filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("%s\n", filePath)
	return cmd.Run()
}

func build(cli *cli.Context, args []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	makefilePath := filepath.Join(cwd, "Makefile.bat")
	if _, err := os.Stat(makefilePath); err == nil {
		return executeMakefile(makefilePath, args...)
	}

	fileWalk := false
	err = filepath.Walk(cwd, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() == "Makefile.bat" {
			fileWalk = true
			return executeMakefile(path, args...)
		}
		return nil
	})

	if err != nil {
		return err
	}

	if !fileWalk {
		return NotFoundMakeFile(cli)
	}

	return nil
}
