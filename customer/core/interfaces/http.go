package interfaces

type HttpRequest struct {
	Params map[string]string
	Body   []byte
}

type HttpResponse struct {
	StausCode int
	Body      interface{}
}

type HttpHandler func(HttpRequest) HttpResponse
