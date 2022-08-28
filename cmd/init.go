package cmd

import (
	"os"
        "github.com/sherlock1453/vcs-go/internal"
	logrus "github.com/sirupsen/logrus"
)

func InitCmd(args []string, logger *logrus.Logger){
    wd, _ := os.Getwd()
    if internal.PathExists(wd){
        logger.Fatal("Repository already exists")
    } else {
        internal.InitRepo(wd, logger)
    }
}
