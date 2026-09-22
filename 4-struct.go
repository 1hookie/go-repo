package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func (u *User) Birthday() {
	u.Age++
}

func main() {
	user := User{
		ID:   1,
		Name: "Step",
		Age:  20,
	}

	fmt.Println(user.Name)
	fmt.Println(user.IsAdult())
	user.Birthday()
	fmt.Println(user.Age)
}
