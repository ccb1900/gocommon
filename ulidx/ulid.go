package ulidx

import (
	"strings"

	"github.com/oklog/ulid/v2"
)

func Get() string {
	return strings.ToLower(ulid.Make().String())
}
