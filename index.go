package main

import "fmt"

type addressType struct {
	country string
	city    string
	street  string
}
type Employee struct {
	firstName string
	lastName  string
	age       int
	addressType
}

func printEmployeeData(employee1 Employee) {

	// fmt.Printf("This is the employee data %+v", employee1)
}

func bulkSend(numMessages int) float64 {
	totalCost := 0.0

	for i := 0; i < numMessages; i++ {
		costToAdd := 1.0 + 0.01*float64(i)
		totalCost = totalCost + costToAdd
	}

	return totalCost
}

func sum(nums ...int) {
	fmt.Println(nums)
}

func main() {

	var employee1 Employee
	employee1.firstName = "Asmar"
	employee1.lastName = "Ali"
	employee1.age = 27
	employee1.country = "Pak"
	employee1.city = "Isb"
	employee1.street = "123 Baker street"

	printEmployeeData(employee1)
	mySlice := []int{1, 2, 3, 4}

	sum(mySlice...)

}
