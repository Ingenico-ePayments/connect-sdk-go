package communicator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	// ErrNoName occurs when the given name is empty
	ErrNoName = errors.New("name is required")
)

// RequestParam represents a single request parameter. Immutable.
type RequestParam struct {
	name, value string
}

// RequestParams represents a slice of RequestParam
type RequestParams []RequestParam

// NewRequestParam creates a RequestParam with the given name and value
func NewRequestParam(name, value string) (*RequestParam, error) {
	if len(name) == 0 {
		return nil, ErrNoName
	}
	return &RequestParam{name, value}, nil
}

// Name returns the name of the RequestParam
func (rp RequestParam) Name() string {
	return rp.name
}

// Value returns the value of the RequestParam
func (rp RequestParam) Value() string {
	return rp.value
}

// String is the implmenetation of the Stringer interface
// Format: 'name:value'
func (rp RequestParam) String() string {
	return rp.name + ":" + rp.value
}

// AddRequestParameter adds a request parameter with the given name and value to the given list, unless if the value if nil.
// The types string, int, in64, bool and slices are supported
func AddRequestParameter(requestParameters *RequestParams, name string, value interface{}) error {
	return addParameter(requestParameters, name, value, true)
}

func addParameter(requestParameters *RequestParams, name string, value interface{}, allowCollection bool) error {
	maybeValue := reflect.ValueOf(value)

	if maybeValue == reflect.Zero(maybeValue.Type()) {
		return nil
	}

	actualIndirectValue := reflect.Indirect(maybeValue)

	switch actualIndirectValue.Kind() {
	case reflect.String:
		{
			var actualValue string
			actualValue, _ = actualIndirectValue.Interface().(string)

			if len(actualValue) < 1 {
				return nil
			}

			var newReqParam *RequestParam
			newReqParam, err := NewRequestParam(name, actualValue)

			if err != nil {
				return err
			}

			*requestParameters = append(*requestParameters, *newReqParam)

			break
		}
	case reflect.Int:
		{
			var actualValue int
			actualValue, _ = actualIndirectValue.Interface().(int)

			var newReqParam *RequestParam
			newReqParam, err := NewRequestParam(name, strings.ToLower(strconv.Itoa(actualValue)))

			if err != nil {
				return err
			}

			*requestParameters = append(*requestParameters, *newReqParam)

			break
		}
	case reflect.Int64:
		{
			var actualValue int64
			actualValue, _ = actualIndirectValue.Interface().(int64)

			var newReqParam *RequestParam
			newReqParam, err := NewRequestParam(name, strings.ToLower(strconv.FormatInt(actualValue, 10)))

			if err != nil {
				return err
			}

			*requestParameters = append(*requestParameters, *newReqParam)

			break
		}
	case reflect.Bool:
		{
			var actualValue bool
			actualValue, _ = actualIndirectValue.Interface().(bool)

			var newReqParam *RequestParam
			newReqParam, err := NewRequestParam(name, strings.ToLower(strconv.FormatBool(actualValue)))

			if err != nil {
				return err
			}

			*requestParameters = append(*requestParameters, *newReqParam)

			break
		}
	case reflect.Slice:
		{
			if allowCollection {
				collection := reflect.ValueOf(actualIndirectValue.Interface())

				for i := 0; i < collection.Len(); i++ {
					err := addParameter(requestParameters, name, collection.Index(i).Interface(), false)

					if err != nil {
						return err
					}
				}
			} else {
				return fmt.Errorf("nested slices are not allowed %s", actualIndirectValue.Kind().String())
			}

			break
		}
	default:
		{
			return fmt.Errorf("unsupported requestParameter type %s", actualIndirectValue.Kind().String())
		}
	}

	return nil
}
