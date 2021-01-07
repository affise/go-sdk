package affise

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
)

var (
	defaultEncoder        = newEncoder()
	errEncodeNilInterface = errors.New("encode err: interface must be not nil")
	errEncodeNotSlice     = errors.New("encode err: interface must be slice")
)

type valuer interface {
	values() (url.Values, error)
}

type encoder struct {
	encoder *schema.Encoder
}

func newEncoder() *encoder {
	e := schema.NewEncoder()

	return &encoder{encoder: e}
}

func (e *encoder) encode(src interface{}) (url.Values, error) {
	values := url.Values{}
	if err := e.encoder.Encode(src, values); err != nil {
		err = fmt.Errorf("encode err: %w", err)

		return nil, err
	}

	return values, nil
}

func (e *encoder) encodeSlice(alias string, src interface{}) (url.Values, error) {
	if src == nil {
		return nil, errEncodeNilInterface
	}

	typ := reflect.TypeOf(src).Kind()
	if typ != reflect.Slice {
		return nil, errEncodeNotSlice
	}

	ret := url.Values{}
	s := reflect.ValueOf(src)
	for i := 0; i < s.Len(); i++ {
		values, err := e.encode(s.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		for k, v := range values {
			param := fmt.Sprintf("%s[%d][%s]", alias, i, k)
			ret[param] = v
		}
	}

	return ret, nil
}

func (e *encoder) encodeMap(alias string, src map[string]string) (url.Values, error) {
	if src == nil {
		return nil, nil
	}

	ret := url.Values{}
	for k, v := range src {
		param := fmt.Sprintf("%s[%s]", alias, k)
		ret.Add(param, v)
	}

	return ret, nil
}

func commaSeparatedInts(src []int) string {
	s := make([]string, 0, len(src))
	for _, v := range src {
		s = append(s, strconv.Itoa(v))
	}

	return strings.Join(s, ",")
}

func commaSeparatedStrings(src []string) string {
	return strings.Join(src, ",")
}

func mergeValues(args ...url.Values) url.Values {
	res := url.Values{}
	for _, values := range args {
		for k, v := range values {
			for _, s := range v {
				res.Add(k, s)
			}
		}
	}

	return res
}
