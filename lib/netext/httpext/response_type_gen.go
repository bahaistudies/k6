// Code generated by "enumer -type=ResponseType -transform=snake -json -text -trimprefix ResponseType -output response_type_gen.go"; DO NOT EDIT.

package httpext

import (
	"encoding/json"
	"fmt"
)

const _ResponseTypeName = "textbinarynone"

var _ResponseTypeIndex = [...]uint8{0, 4, 10, 14}

func (i ResponseType) String() string {
	if i >= ResponseType(len(_ResponseTypeIndex)-1) {
		return fmt.Sprintf("ResponseType(%d)", i)
	}
	return _ResponseTypeName[_ResponseTypeIndex[i]:_ResponseTypeIndex[i+1]]
}

var _ResponseTypeValues = []ResponseType{0, 1, 2}

var _ResponseTypeNameToValueMap = map[string]ResponseType{
	_ResponseTypeName[0:4]:   0,
	_ResponseTypeName[4:10]:  1,
	_ResponseTypeName[10:14]: 2,
}

// ResponseTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResponseTypeString(s string) (ResponseType, error) {
	if val, ok := _ResponseTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResponseType values", s)
}

// ResponseTypeValues returns all values of the enum
func ResponseTypeValues() []ResponseType {
	return _ResponseTypeValues
}

// IsAResponseType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResponseType) IsAResponseType() bool {
	for _, v := range _ResponseTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for ResponseType
func (i ResponseType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseType
func (i *ResponseType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ResponseType should be a string, got %s", data)
	}

	var err error
	*i, err = ResponseTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for ResponseType
func (i ResponseType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ResponseType
func (i *ResponseType) UnmarshalText(text []byte) error {
	var err error
	*i, err = ResponseTypeString(string(text))
	return err
}