package internal

import (
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func InitRepo(wd string, logger *logrus.Logger){
    err:= os.Mkdir(filepath.Join(wd, ".vcs"), 0777)
    if err!=nil {
        logger.Infoln("Repository Initialized")
    }
}
