// Code generated by github.com/OldBigBuddha/gqlgen, DO NOT EDIT.

package out

import (
	"fmt"
	"io"
	"strconv"
)

type SomeContent string

const (
	SomeContentThis SomeContent = "This"
	SomeContentIs   SomeContent = "Is"
	SomeContentA    SomeContent = "A"
	SomeContentTest SomeContent = "Test"
)

var AllSomeContent = []SomeContent{
	SomeContentThis,
	SomeContentIs,
	SomeContentA,
	SomeContentTest,
}

func (e SomeContent) IsValid() bool {
	switch e {
	case SomeContentThis, SomeContentIs, SomeContentA, SomeContentTest:
		return true
	}
	return false
}

func (e SomeContent) String() string {
	return string(e)
}

func (e *SomeContent) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SomeContent(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SomeContent", str)
	}
	return nil
}

func (e SomeContent) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
