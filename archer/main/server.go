package main

import (
	"github.com/liu2hai/chive/archer/bows"
	"github.com/liu2hai/chive/config"
)

// RunServer  start servers
func RunServer() error {
	err := bows.InitKafkaClient(config.T.Broker)
	if err != nil {
		return err
	}

	bl := bows.InitBows()
	exchanges := config.T.Exchanges
	if err := bows.StartExArcher(exchanges, bl); err != nil {
		return err
	}
	bows.StartCmdLoop(bl)
	return nil
}
