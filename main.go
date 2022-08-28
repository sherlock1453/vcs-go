package main

import (
	"os"

	"github.com/sherlock1453/vcs-go/cmd"
	"github.com/sherlock1453/vcs-go/internal"
	logrus "github.com/sirupsen/logrus"
)



func main(){
    logger:=logrus.New()
    if internal.Contains(os.Args, "-v"){
        logger.SetLevel(logrus.DebugLevel)
    }
    command:= os.Args[1]
    logger.Infoln("command is init")
    if command == "init"{
        cmd.InitCmd(os.Args, logger)
    }
}
