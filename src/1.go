
// This program calculates the area of a circle using the formula A = πr^2, where r is the radius.
package main
import (
	"fmt"
	"math"
)
func main() {
	var radius float64
	fmt.Println("Enter the value of the radius: ")
	fmt.Scanf("%f", &radius)
	area := math.Pi * radius * radius
	fmt.Println("The area of the circle is: ", area)
}