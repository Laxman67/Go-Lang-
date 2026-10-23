package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// (u User ) only for Reading
// (u *User) for Update the data
func (u User) Greet(name string) string {
	u.Name = name
	fmt.Println(u)
	return "Hi I'm " + u.Name
}

func (u *User) HaveBirthday() {
	u.Age++
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}
type Employee struct {
	Name    string   `json:"name"`
	Address Address  `json:"address"`
	Skills  []string `json:"skills"`
}

func (u *User) SetEmail(e string) {
	u.Email = e
}

func main() {

	e := Employee{
		Name:    "Amit",
		Address: Address{City: "Pune", State: "MH"},
		Skills:  []string{"Go", "SQL", "Docker"},
	}
	data, _ := json.Marshal(e)
	fmt.Println("User Details: 1 ", e)

	fmt.Println("User Details: 2 ", string(data))
	jsonStr := `{"name":"Priya","address":{"city":"Delhi","state":"DL"},"skills":["Go"]}`

	var e2 Employee
	err := json.Unmarshal([]byte(jsonStr), &e2) // ⚠️ []byte + &e2
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(e2.Name)         // Priya
	fmt.Println(e2.Address.City) // Delhi
	fmt.Println(e2.Skills[0])    // Go

	for _, s := range e.Skills {
		fmt.Println(s)
	}

	// user3 := User{Name: "Laksh", Email: "laxman.webdev67@gmail.com", Age: 32}
	// data, _ := json.Marshal(user3)
	// fmt.Println(string(data))

	// var u2 User
	// json.Unmarshal(data, &u2)
	// fmt.Println(u2.Age)
}
