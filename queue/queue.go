package queue

type CheckFunc func(string) (string, error)
type Queue struct {
	incoming chan string
	resultChan chan []string
	failChan chan string

	checkFunc CheckFunc
}

func NewQueue(checkFunc CheckFunc) *Queue {
	incoming := make(chan string, 1000)
	resultChan := make(chan []string, 1000)
	failChan := make(chan string, 1000)

	return &Queue{incoming: incoming, resultChan: resultChan, failChan: failChan}
}

func (q *Queue) Start() {
	for i:=0; i<3; i++ {
		go func() {
			for  {
				url := <-q.incoming
				result, err := q.checkFunc(url)
				if err != nil {
					q.failChan <- url
				} else {
					q.resultChan <- []string{url, result}
				}

			}
		}()
	}
	
}

func (q *Queue) Enqueue(url string) error {
	q.incoming <- url
	return nil
}

func (q *Queue) ResultChan() <-chan []string  {
	return q.resultChan
}

func (q *Queue) FailChan() <-chan string {
	return q.failChan
}
