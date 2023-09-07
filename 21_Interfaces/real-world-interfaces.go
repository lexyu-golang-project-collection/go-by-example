package main

import (
	"net/http"
)

type User struct {
	id   int
	name string
}

// func GetEntity(*http.Request) (interface{}, error)

// func GetUser(*http.Request) (User, error)

type Entity interface {
	UnmarshalHTTP(*http.Request) error
}

func GetEntity(r *http.Request, v Entity) error {
	return v.UnmarshalHTTP(r)
}

func (u *User) UnmarshalHTTP(r *http.Request) error {
	// ...
	return nil
}

func main() {

	/*
		var u User
		if err := GetEntity(req, &u); err != nil {
			// ...
		}
	*/

	// 1. create abstractions by considering the functionality
	//    that is common between datatypes, instead of the fields that are common between datatypes
	// 2. an interface{} value is not of any type
	// 3. interfaces are two words wide; schematically they look like (type, value)
	// 4. it is better to accept an interface{} value than it is to return an interface{} value
	// 5. a pointer type may call the methods of its associated value type, but not vice versa
	// 6. everything is pass by value, even the receiver of a method
	// 7. an interface value isn’t strictly a pointer or not a pointer, it’s just an interface
	// 8. if you need to completely overwrite a value inside of a method, use the * operator to manually dereference a pointer

}
