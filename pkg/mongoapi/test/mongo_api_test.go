package mongoapi

import (
	"context"
	"crayontool-go/pkg/mongoapi"
	"fmt"
	"os"
	"sync"
)

const uri = "mongodb://%s:%s@mongo-test.sucrayon.top:27017"

const (
	envUsername = "mongo_username"
	envPassword = "mongo_password"
)

var (
	once      sync.Once
	globalCtl mongoapi.MongoCtl
)

func ctlInit() {
	once.Do(func() {
		username := os.Getenv(envUsername)
		password := os.Getenv(envPassword)
		ctl := mongoapi.NewMongoCtl(context.Background())
		err := ctl.ApplyURIAndConnect(fmt.Sprintf(uri, username, password))
		if err != nil {
			panic(err)
		}
		globalCtl = ctl
	})
}
