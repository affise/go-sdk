package affise

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_encoder_encodeSlice(t *testing.T) {
	t.Parallel()

	type T struct {
		Field1 int    `schema:"field1"`
		Field2 string `schema:"field2"`
	}
	v1 := T{1, "x"}
	v2 := T{2, "y"}
	alias := "list"

	tests := []struct {
		has  interface{}
		want string
	}{
		{[]T{v1, v2}, "list[0][field1]=1&list[0][field2]=x&list[1][field1]=2&list[1][field2]=y"},
		{[]*T{&v1}, "list[0][field1]=1&list[0][field2]=x"},
	}

	for _, tt := range tests {
		has, err := defaultEncoder.encodeSlice(alias, tt.has)
		require.NoError(t, err)

		want, err := url.ParseQuery(tt.want)
		require.NoError(t, err)

		require.Equal(t, want, has)
	}
}

func Test_encoder_encodeSliceErr(t *testing.T) {
	t.Parallel()

	tests := []interface{}{
		nil,
		42,
		"foo",
	}

	for _, tt := range tests {
		has, err := defaultEncoder.encodeSlice("bar", tt)
		require.Error(t, err)
		require.Nil(t, has)
	}
}

func Test_commaSeparatedInts(t *testing.T) {
	t.Parallel()

	tests := []struct {
		has  []int
		want string
	}{
		{[]int{1, 2, 3}, "1,2,3"},
		{[]int{1}, "1"},
		{[]int{}, ""},
	}

	for _, tt := range tests {
		want := commaSeparatedInts(tt.has)
		require.Equal(t, tt.want, want)
	}
}

func Test_commaSeparatedStrings(t *testing.T) {
	t.Parallel()

	tests := []struct {
		has  []string
		want string
	}{
		{[]string{"DE", "UA"}, "DE,UA"},
		{[]string{"DE"}, "DE"},
		{[]string{}, ""},
	}

	for _, tt := range tests {
		want := commaSeparatedStrings(tt.has)
		require.Equal(t, tt.want, want)
	}
}
