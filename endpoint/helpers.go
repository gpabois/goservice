package endpoint

import (
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde/decoder"
	"github.com/gpabois/gostd/serde/norm"
)

func Reflect_Inject(dest any, values map[string]any) result.Result[bool] {
	// Decode the normalised map into the endpoint request
	d := norm.NewDecoder(values)
	res := decoder.Reflect_DecodeInto(d, dest)
	if res.HasFailed() {
		return result.Result[bool]{}.Failed(res.UnwrapError())
	}
	return result.Success(true)
}
