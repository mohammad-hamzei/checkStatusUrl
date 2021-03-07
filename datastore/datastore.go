package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type Datastore struct {
	url string
	status uint
}

func NewDatastore() *Datastore  {
	return &Datastore{}
}




/*func (d *Datastore) upsert(url string, status string) error {
	store := []Datastore{}
	fi, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 1024)
	for {
		// read a chunk
		n, err := fi.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}
		store = append(url, status)
		if _, err := fi.Write(store ); err != nil {
			panic(err)
		}
	}
	return nil
}*/

func (d *Datastore) upsert(url string, status string) error {
	data := Datastore{
		url: status,
	}
	file, _ := json.MarshalIndent(data, "", " ")

	_ = ioutil.WriteFile("input.txt", file, 0644)
	return nil
}

func (d *Datastore) Exists(url string)(bool, error) {
	_, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		if !os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (d *Datastore) GetStatus(url string)(string, error) {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	result := strings.Split(string(data), "\n")
	return result[0], nil
}
func (d *Datastore) GetResult(url string)(string, error) {
	f, err := os.OpenFile("input.txt", os.O_RDONLY, 0)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	result := strings.Split(string(data), "\n")
	return result[1], nil
}