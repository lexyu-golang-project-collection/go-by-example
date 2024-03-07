package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

// start with a string representation of our JSON data
var input = `
{
	"created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	fmt.Println(string(b))
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v) // v is  time.Time value
	return nil
}

// The real world: getting a proper timestamp out of the Twitter API
func main() {
	// our target will be of type map[string]interface{}, which is a
	// pretty generic type that will give us a hashtable whose keys
	// are strings, and whose values are of type interface{}
	var val1 map[string]interface{}

	// var val2 map[string]time.Time

	// var val3 map[string]Timestamp

	if err := json.Unmarshal([]byte(input), &val1); err != nil {
		panic(err)
	}

	fmt.Println("val = ", val1)
	for k, v := range val1 {
		fmt.Println(k, reflect.TypeOf(v))
	}

	// fmt.Println(time.Time(val3["create_at"]))
}
