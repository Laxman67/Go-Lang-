package main

type User struct {
	Name    string `json:name`
	Email   string `json:email`
	Address string `json:address`
}

func (u *User) Greet() string {
	return "Hello : " + u.Name
}

func main() {
	u := User{
		Name:    "Amit",
		Address: "Batala",
		Email:   "laxman@gmail.com",
	}

	println(u.Greet())

}
