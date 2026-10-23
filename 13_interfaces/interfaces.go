package main

import "fmt"

// Interface ek contract bolti hai: "Jo bhi type mere saare methods implement kare, woh meri type ban jaayegi."
type Speaker interface {
	Speak() string // sirf signature, body NAHI
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d Dog) Speak() string { return d.Name + " : Bhookh!" }

func (c Cat) Speak() string { return c.Name + " : Meow" }

func describe(i interface{}) { // interface{} = kuch bhi ho sakta hai
	switch v := i.(type) {
	case int:
		fmt.Println("Integer:", v*2)
	case string:
		fmt.Println("String:", len(v))
	case Dog:
		fmt.Println("Dog:", v.Name)
	case bool:
		fmt.Println("Boolean:", v)
	case Cat:
		fmt.Println("Cat:", v)
	default:
		fmt.Println("Unknown type")
	}
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

// ✅ Data add kiya
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// ✅ Interface ke dono methods implement kiye (exact match!)
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

// ✅ Ek hi function — koi bhi Shape chalega (POLYMORPHISM!)
func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f | Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

type Vehicle interface {
	StartEngine() string
	StopEngine() string
	CalculateOil() string
	DisplayDetails()
}
type VehicleSpecification struct {
	Name     string
	Brand    string
	FuelType string
}

type Car struct {
	VehicleSpecification
}
type Bike struct {
	VehicleSpecification
}

func (s VehicleSpecification) StartEngine() string  { return s.Name + " Engine Started" }
func (s VehicleSpecification) StopEngine() string   { return s.Name + " Engine Stopped" }
func (s VehicleSpecification) CalculateOil() string { return s.Name + " " + s.FuelType + " Oil (ML)" }
func (s VehicleSpecification) DisplayDetails() {
	// ✅ Printf use kiya (Mistake 2 fix)
	fmt.Printf("Name: %s | Brand: %s | Fuel: %s\n", s.Name, s.Brand, s.FuelType)
}

// ye function sabhi Vehicle type k liye valid hai
func displayDetails(v Vehicle) {
	v.DisplayDetails()
	fmt.Println("  Start :", v.StartEngine())
	fmt.Println("  Oil   :", v.CalculateOil())
	fmt.Println("  Stop  :", v.StopEngine())
	fmt.Println()

}

func main() {
	car := Car{VehicleSpecification{Name: "Maruti", Brand: "Toyota", FuelType: "Diesel"}}
	bike := Bike{VehicleSpecification{Name: "Hybuzza", Brand: "FZ", FuelType: "Petrol"}}

	displayDetails(car)  // ✅ dono chalenge
	displayDetails(bike) // ✅ same function!
}

// var s Shape = Rectangle{Width: 9, Height: 4}
// 	printShapeDetails(s) // Area: 36.00 | Perimeter: 26.00

// 	printShapeDetails(Rectangle{Width: 10, Height: 8})
// 	printShapeDetails(Circle{Radius: 5})

// var anything interface{}
// anything = 12
// fmt.Println(anything)

// anything = "Hello"
// anything = true
// fmt.Println(anything)

// describe("Hello")
// describe(12)
// d := Dog{Name: "Doggy"}
// c := Cat{Name: "Catty"}

// describe(true)
// describe(d)
// describe(c)

// var s Speaker = Dog{Name: "Tommy"}
// d, ok := s.(Dog)
// if ok {
// 	fmt.Println("It's a Dog!", d.Name) // It's a Dog! Tommy
// } else {
// 	fmt.Println("Not a Dog")
// }

// }
