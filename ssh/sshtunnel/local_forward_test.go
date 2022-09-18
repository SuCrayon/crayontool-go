package sshtunnel

import (
	"crayontool-go/osutil"
	"fmt"
	"os"
	"testing"
)

func Test(t *testing.T) {
	lfConfig := NewDefaultConfig().SetSSHTunnelIP("192.168.124.12").SetSSHTunnelPort(10022).SetUser("crayon").SetPassword("slb614820984").SetLocalIP("127.0.0.1").SetLocalPort(9092).SetRemoteIP("127.0.0.1").SetRemotePort(9092)
	lfClient := NewClient(lfConfig)
	osutil.ListenSignalAsync(func() {
		lfClient.Stop()
	}, os.Interrupt, os.Kill)
	err := lfClient.Start()
	if err != nil {
		fmt.Printf("some error occur when start local forward client, err: %v\n", err)
		return
	}
}
