package ogjson

import (
	core_json "github.com/og/json/core"
	"strings"
)

// encode (without error)
func String(v interface{}) (jsonString string) {
	jsonString, err := StringWithErr(v); if err != nil { panic(err) }
	return
}
// encode (with error)
func StringWithErr (v interface{}) (jsonString string, err error) {
	bjson, err := core_json.Marshal(v, "json")
	if err != nil {
		return
	}
	jsonString = string(bjson)
	return
}
// encode  unfold (space 2)
func StringUnfold(v interface{}) (jsonString string) {
	jsonString, err := StringSpaceWithErr(v ,2)
	if err != nil { panic(err) }
	return
}
// encode  unfold (space 2) (with error)
func StringUnfoldWithErr(v interface{}) (jsonString string, err error) {
	jsonString, err = StringSpaceWithErr(v ,2)
	if err != nil {
		return
	}
	return
}
// encode pretty-print
func StringSpace(v interface{}, space int) (jsonString string) {
	jsonString, err := StringSpaceWithErr(v ,space)
	if err != nil { panic(err) }
	return
}
// encode pretty-print (with error)
func StringSpaceWithErr(v interface{}, space int) (jsonString string, err error) {
	bjson, err := core_json.MarshalIndent(v, "", strings.Repeat(" ", space), "json")
	if err != nil {
		return
	}
	jsonString = string(bjson)
	return
}
func Bytes(v interface{}) []byte {
	bjson, err := core_json.Marshal(v, "json")
	if err != nil { panic(err) }
	return bjson
}
// encode to []byte (with error)
func BytesWithErr(v interface{}) ([]byte, error) {
	return core_json.Marshal(v, "json")
}


// decode format string
// Parse(`{"name":"nimo"}`, &user)
// in the value pointed to by v. If v test nil or not a pointer,
// Parse returns an InvalidUnmarshalError.
func Parse(jsonString string, v interface{}) {
	err := ParseWithErr(jsonString, &v)
	if err != nil { panic(err) }
}
// decode format string (with error)
// in the value pointed to by v. If v test nil or not a pointer,
// Parse returns an InvalidUnmarshalError.
func ParseWithErr(jsonString string,  v interface{}) (err error) {
	err = core_json.Unmarshal([]byte(jsonString), &v, "json")
	return
}

// decode by []byte
// Parse([]byte(`{"name":"nimo"}`), &user)
// in the value pointed to by v. If v test nil or not a pointer,
// Parse returns an InvalidUnmarshalError.
func ParseBytes (data []byte, v interface{}) {
	err := ParseBytesWithErr(data, &v); if err != nil { panic(err) }
}
// decode by []byte (with error)
// equal json.Unmarshal(data []byte, v interface{}) error
// in the value pointed to by v. If v test nil or not a pointer,
// Parse returns an InvalidUnmarshalError.
func ParseBytesWithErr (data []byte, v interface{}) (err error) {
	err = core_json.Unmarshal(data, &v, "json")
	return
}
