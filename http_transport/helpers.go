package http_transport

import (
	"io"
	"net/http"

	"github.com/gpabois/goservice/endpoint"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
)

type HttpResult[T any] struct {
	Data  option.Option[T]
	Error string
}

func (res HttpResult[T]) Failed(err error, code option.Option[int]) HttpResult[T] {
	return HttpResult[T]{}
}

func WriteResult(res result.Result[any], w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Accept")

	if contentType == "" {
		contentType = "application/json"
	}

	if res.HasFailed() {
		err := HttpError_From(res.UnwrapError())
		w.WriteHeader(err.Code())

		w.Header().Set("Content-Type", contentType)
		encodedRes := serde.Serialize(HttpResult[any]{Error: err.Error()}, contentType)

		// Fallback, write the error as text/plain
		if encodedRes.HasFailed() {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte(err.Error()))
		}

		return
	}

	anyResp := res.Expect()
	switch resp := anyResp.(type) {
	case endpoint.StreamMedia:
		w.Header().Set("Content-Type", resp.MimeType)
		io.Copy(w, resp.Stream)
	default:
		res := serde.Serialize(r, contentType)
		if res.HasFailed() {
			WriteResult(res.ToAny(), w, r)
		}
		w.Write(res.Expect())
	}
}
