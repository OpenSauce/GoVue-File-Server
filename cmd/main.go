package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/OpenSauce/GoVue-File-Server/pkg/server"
)

var doneCh = make(chan struct{}, 1)

//Main execution function of the server.
func main() {
	sigs := make(chan os.Signal, 1)
	go server.Start(doneCh)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	doneCh<-struct{}{}
	
}
