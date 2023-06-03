package env

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

func PrintAllEnvsToFile() {
	if err := func() error {
		if Getenv("PRINT_ALL_ENVS") != "true" {
			logrus.Infof("skip printing envs by env var")
			return nil
		}
		var s = strings.Join(os.Environ(), "; ")
		if len(os.Args) == 0 {
			return errors.Errorf("couldn't get program name")
		}
		fileName := "/tmp/" + filepath.Base(os.Args[0]) + "_envs.txt"
		if err := os.WriteFile(fileName, []byte(s), 0777); err != nil {
			return errors.WithStack(err)
		}
		logrus.Tracef("all envs is printed to file %+v", fileName)
		logrus.Tracef("pbcopy < %+v", fileName)
		return nil
	}(); err != nil {
		logrus.Errorf("couldn't print all envs to file: %+v", err)
	}
}
