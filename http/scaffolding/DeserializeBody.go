package http_scaffolding

import (
	http_links "github.com/gpabois/goservice/http/links"
	"github.com/gpabois/goservice/utils"
	"github.com/gpabois/gostd/option"
)

func WithContentTypeFallback(fallback string) utils.Configurator[http_links.ReflectDeserializeBodyArgs] {
	return func(s *http_links.ReflectDeserializeBodyArgs) {
		s.FallbackContentType = option.Some(fallback)
	}
}
