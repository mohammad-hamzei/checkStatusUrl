package dispatcher

type dataStore interface {
	Upsert(url string, status string) error

	Exists(url string)(bool, error)
	GetStatus(url string)(string, error)
	GetResult(url string)(string, error)
	NewDatastore()(interface{})
}

type queue interface {
	Enqueue(url string) error

	ResultChan() <-chan []string
	FailChan() <-chan string

}


