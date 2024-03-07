package main

import "fmt"

type rect struct {
	width, height int
}

// example-1
// receiver type
// pointer receiver, 改變地址指向的值
func (r *rect) area() int {
	return r.width * r.height
}

// value receiver, 不修改值
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

// example-2
type user struct {
	name, email string
}

// value recevier, 不修改值
// method
func (u user) notify() {
	// fmt.Println("This notify is method!")
	fmt.Println("user = ", u)
}

// 不會改變內部的值,僅改變複製變數的值
func (u user) changeName(name string) {
	u.name = name
}

// pointer recevier 修改值
func (u *user) changeEmail(email string) {
	u.email = email
}

// function
func notify() {
	fmt.Println("This notify is function !")
}

func main() {
	// example-1
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	// example-2
	notify()
	fmt.Println()

	tim := user{name: "Tim", email: "tim@gmail.com"}
	tim.notify()
	fmt.Println()

	tim.changeName("test")
	fmt.Println("After Call changeName() value receiver to change user name = test")
	tim.notify() // (&tim).notify()

	tim.changeEmail("test@abc.com")
	fmt.Println("After Call changeEmail() pointer receiver to change user email = test@abc.com")
	tim.notify()
}
