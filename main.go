package main

import (
	"fmt"
	"github.com/mohammad-hamzei/yourypto/api"
	"github.com/mohammad-hamzei/yourypto/check"
	"github.com/mohammad-hamzei/yourypto/datastore"
	"github.com/mohammad-hamzei/yourypto/dispatcher"

	//"github.com/mohammad-hamzei/yourypto/dispatcher"
	queue2 "github.com/mohammad-hamzei/yourypto/queue"
)

func main()  {
	ds := datastore.NewDatastore()
	checkFunc := check.Check
	q := queue2.NewQueue(checkFunc)
	dsp := dispatcher.NewDispatcher(ds, q)
	fmt.Println("test:",dsp)
	fmt.Println("port: 8080")
	api.Start()

}


