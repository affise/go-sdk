package affise

import (
	"errors"
	"fmt"
)

type CustomBool bool

var errCustomBoolParsing = errors.New("CustomBool: parsing err")

func (b *CustomBool) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"true"`, `true`, `"1"`, `1`:
		*b = true

		return nil
	case `"false"`, `false`, `"0"`, `0`, `""`:
		*b = false

		return nil
	default:

		return fmt.Errorf("%w: unknown value %q", errCustomBoolParsing, string(data))
	}
}
