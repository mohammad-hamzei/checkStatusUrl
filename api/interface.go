package api

type dispatcher interface {
	HandleUrl(url string)error
	Exists(url string) (bool, error)
	IsPending(url string) (bool, error)
	GetResult(url string)(string, error)
}
