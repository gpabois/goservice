package pg

import (
	"errors"
	"reflect"

	"github.com/gpabois/gostd/geojson"
	"github.com/gpabois/gostd/option"
	"github.com/gpabois/gostd/result"
	"github.com/gpabois/gostd/serde"
	"github.com/jackc/pgx/v5"
)

// Scan the next value of the rows
// Also process optional values, and geojson.
func Scan(rows pgx.Rows, dest any) result.Result[bool] {
	val := reflect.ValueOf(dest)
	typ := val.Elem().Type()

	// Optional value get a special treatment
	if option.Reflect_IsOptionType(typ) {
		innerVal := reflect.New(typ)
		rows.Scan(innerVal.Interface())
		return option.Reflect_TrySome(val, innerVal.Elem())
	}

	switch dest.(type) {
	case *geojson.Geometry:
		raw := make([]byte, 100)
		err := rows.Scan(&raw)
		if err != nil {
			return result.Failed[bool](err)
		}
		return serde.Reflect_DeserializeInto(raw, "application/json", dest)
	}

	if typ.Kind() == reflect.Struct {
		return result.Failed[bool](errors.New("Cannot scan struct-based values"))
	}

	err := rows.Scan(dest)
	if err != nil {
		return result.Failed[bool](err)
	}

	return result.Success(true)
}
