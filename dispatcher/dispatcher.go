package dispatcher

import (
	"fmt"
	"github.com/mohammad-hamzei/yourypto/datastore"
)

type Dispatcher struct {
	ds dataStore
	q  queue
}

func NewDispatcher(ds *datastore.Datastore, q queue) *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) HandleUrl(url string) error  {
	err := d.ds.Upsert(url, "pending")
	if err != nil {
		return err
	}

	err = d.q.Enqueue(url)
	if err != nil {
		return err
	}
	return nil
}

func (d *Dispatcher) Start() {
	
	go func() {
		for  {
			select {
			case r:= <-d.q.ResultChan():
				fmt.Println("Result Received: ",r)
				d.ds.Upsert(r[0], "finished")
			case r:= <-d.q.FailChan():
				fmt.Println("Fail Received: ",r)
				d.ds.Upsert(r, "failed")


			}

		}
	}()
}

func (d *Dispatcher)Exists(url string) (bool, error) {
	return d.ds.Exists(url)

}
func (d *Dispatcher)IsPending(url string) (bool, error) {
	status, err := d.ds.GetStatus(url)
	return status == "pending", err
}
func (d *Dispatcher)GetResult(url string)(string, error) {
	return d.ds.GetResult(url)
}
