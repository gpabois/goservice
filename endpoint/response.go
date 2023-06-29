package endpoint

import "io"

type Media interface {
	MimeType() string
}

type StreamMedia struct {
	MimeType string
	Stream   io.Reader
}

func PNG(stream io.Reader) StreamMedia {
	return StreamMedia{MimeType: "image/png", Stream: stream}
}
