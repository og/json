package ogjson

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type User struct {
	Name string
	Age int
}
var user = User {
	Name: "nimo",
	Age: 27,
}
func ExampleString() {
	type User struct {
		Name string
		Age int
	}
	user := User {
		Name: "nimo",
		Age: 27,
	}
	String(user)
	// {"Name":"nimo","Age":27}
}
func ExampleParse() {
	type User struct {
		Name string
		Age int
	}
	var user User
	// In the value pointed to by user. If v test nil or not a pointer,
	// Parse returns an InvalidUnmarshalError.
	Parse(`{"Name":"nimo","Age":27}`, &user)
}

var userJSON = `{"Name":"nimo","Age":27}`
func TestString(t *testing.T) {
	{
		json := String(user)
		assert.Equal(t,userJSON, json)
	}
	{
		json, err := StringWithErr(user)
		assert.Equal(t,userJSON, json)
		assert.Equal(t,nil, err)
	}
	{
		json, err := StringWithErr(log.Print)
		assert.Equal(t,"", json)
		if err == nil {
			panic("ByteWithErr(log.Print) should return error")
		}
	}
}

func TestStringSpace(t *testing.T) {
	{
		json := StringSpace(user, 2)
		assert.Equal(t,"{\n  \"Name\": \"nimo\",\n  \"Age\": 27\n}", json)
	}
	{
		json, err := StringSpaceWithErr(user, 2)
		assert.Equal(t,"{\n  \"Name\": \"nimo\",\n  \"Age\": 27\n}", json)
		assert.Equal(t,nil, err)
	}
	{
		json, err := StringSpaceWithErr(log.Print, 2)
		assert.Equal(t,"", json)
		if err == nil {
			panic("StringIndentWithErr(log.Print) should return error")
		}
	}
}

func TestBytes(t *testing.T) {

	{
		json := Bytes(user)
		assert.Equal(t,[]byte(userJSON), json)
	}
	{
		user := User{
			Name: "nimo",
			Age: 27,
		}
		json, err := BytesWithErr(user)
		assert.Equal(t,[]byte(userJSON), json)
		assert.Equal(t,nil, err)
	}
	{
		json, err := BytesWithErr(log.Print)
		assert.Equal(t,[]byte(nil), json)
		if err == nil {
			panic("ByteWithErr(log.Print) should return error")
		}
	}
}


func TestParse(t *testing.T) {

	{
		{
			var user User
			Parse(userJSON, &user)
			assert.Equal(t,User{
				Name: "nimo",
				Age: 27,
			}, user)
		}
		{
			var user User
			Parse(userJSON, user)  // not pointer
			assert.Equal(t,User{
				Name: "",
				Age: 0,
			}, user)
		}
	}
	{
		{
			var user User
			err := ParseWithErr(userJSON, &user)
			assert.Equal(t,User{
				Name: "nimo",
				Age: 27,
			}, user)
			assert.Equal(t,nil, err)
		}

		{
			var user User
			err := ParseWithErr(``, &user)
			assert.Equal(t,User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}

		{
			var user User
			err := ParseWithErr(userJSON, user) // not pointer
			assert.Equal(t,User{
				Name: "",
				Age: 0,
			}, user)
			assert.Equal(t,nil, err)
		}

		{
			var user User
			err := ParseWithErr(``, user) // not pointer
			assert.Equal(t,User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}
	}
}


func TestParseByte(t *testing.T) {

	{
		{
			var user User
			ParseBytes([]byte(userJSON), &user)
			assert.Equal(t,User{
				Name: "nimo",
				Age: 27,
			}, user)
		}
		{
			var user User
			ParseBytes([]byte(userJSON), user)  // not pointer
			assert.Equal(t,User{
				Name: "",
				Age: 0,
			}, user)
		}
	}
	{
		{
			var user User
			ParseBytes([]byte(userJSON), &user)
			assert.Equal(t,User{
				Name: "nimo",
				Age: 27,
			}, user)
		}

		{
			var user User
			err := ParseBytesWithErr([]byte(``), &user)
			assert.Equal(t,User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}

		{
			var user User
			err := ParseBytesWithErr([]byte(userJSON), user) // not pointer
			assert.Equal(t,User{
				Name: "",
				Age: 0,
			}, user)
			assert.Equal(t,nil, err)
		}

		{
			var user User
			err := ParseBytesWithErr([]byte(``), user) // not pointer
			assert.Equal(t,User{
				Name: "",
				Age: 0,
			}, user)
			if err == nil {
				panic("ParseWithErr(``) should return error")
			}
		}
	}
}

func TestStringUnfold(t *testing.T) {

	{
		var user User
		userUnfoldJSON := `{
  "Name": "",
  "Age": 0
}`
		assert.Equal(t,userUnfoldJSON, StringUnfold(user))
	}
}
func TestEmptyListMap (t *testing.T) {

	assert.Equal(t,`{"List":[],"Map":{}}`, String(struct {
		List []string
		Map map[string]interface{}
	}{}))
}

func TestStringConvInt (t *testing.T) {

	query := struct {
		Page int `json:"page"`
	}{}
	Parse(`{"page": "2"}`,&query)
	assert.Equal(t,2, query.Page)
}
func TestStringConvIntAndFloat (t *testing.T) {

	{
		query := struct {
			Page int `json:"page"`
		}{}
		Parse(`{"page": "2"}`,&query)
		assert.Equal(t,2, query.Page)
	}
	{
		query := struct {
			Page float64 `json:"page"`
		}{}
		Parse(`{"page": "2.2"}`,&query)
		assert.Equal(t,2.2, query.Page)
	}
}
// func TestInterface (t *testing.T) {
// 	{
// 		data := struct {
// 			Date SecondTime
// 			Name string
// 		}{}
// 		Parse(`{"Date":"2020-02-28 20:48:45"}`, &data)
// 		assert.Equal(t,data.Date.Format(gtime.Second), "2020-02-28 20:48:45")
// 		assert.Equal(t,String(data), `{"Date":"2020-02-28 20:48:45","Name":""}`)
// 	}
// 	{
// 		data := struct {
// 			Date MinuteTime
// 		}{}
// 		Parse(`{"Date":"2020-02-28 20:48"}`, &data)
// 		assert.Equal(t,data.Date.Format(gtime.Second), "2020-02-28 20:48:00")
// 		assert.Equal(t,String(data), `{"Date":"2020-02-28 20:48"}`)
// 	}
// 	{
// 		data := struct {
// 			Date HourTime
// 		}{}
// 		Parse(`{"Date":"2020-02-28 20"}`, &data)
// 		assert.Equal(t,data.Date.Format(gtime.Second), "2020-02-28 20:00:00")
// 		assert.Equal(t,String(data), `{"Date":"2020-02-28 20"}`)
// 	}
// 	{
// 		data := struct {
// 			Date DayTime
// 		}{}
// 		Parse(`{"Date":"2020-02-28"}`, &data)
// 		assert.Equal(t,data.Date.Format(gtime.Second), "2020-02-28 00:00:00")
// 		assert.Equal(t,String(data), `{"Date":"2020-02-28"}`)
// 	}
// 	{
// 		data := struct {
// 			Date MonthTime
// 		}{}
// 		Parse(`{"Date":"2020-02"}`, &data)
// 		assert.Equal(t,data.Date.Format(gtime.Second), "2020-02-01 00:00:00")
// 		assert.Equal(t,String(data), `{"Date":"2020-02"}`)
// 	}
// 	{
// 		data := struct {
// 			Date YearTime
// 		}{}
// 		Parse(`{"Date":"2020"}`, &data)
// 		assert.Equal(t,data.Date.Format(gtime.Second), "2020-01-01 00:00:00")
// 		assert.Equal(t,String(data), `{"Date":"2020"}`)
// 	}
// }

func Test_ParseSliceNil(t *testing.T) {

	{
		data := struct {
			List []string
		}{}
		Parse(`{"List":[]}`, &data)
		// binding 等库依赖了这个特性，所以不要改变这个特性 @nimoc
		assert.Equal(t,data.List, []string{})
	}
	{
		data := struct {
			List []string
		}{}
		Parse(`{}`, &data)
		// binding 等库依赖了这个特性，所以不要改变这个特性 @nimoc
		assert.Equal(t,data.List, []string(nil))
	}
}