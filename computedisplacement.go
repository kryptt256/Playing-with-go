package done

import (
	"fmt"
)

func main() {
	var time, acceleration, iniVelocity, iniDisplacement float64

	fmt.Print("Enter a value for acceleration: ")
	fmt.Scanf("%f", &acceleration)
	fmt.Print("Enter a value for initial velocity: ")
	fmt.Scanf("%f", &iniVelocity)
	fmt.Print("Enter a value for initial displacement: ")
	fmt.Scanf("%f", &iniDisplacement)
	fmt.Print("Enter a value for time: ")
	fmt.Scanf("%f", &time)

	sOfTime := GenDisplaceFn(acceleration, iniVelocity, iniDisplacement)
	fmt.Println("The displacement S as a function of time T is: ", sOfTime(time))
}

func GenDisplaceFn(acceleration, iniVelocity, iniDisplacement float64) func(float64) float64 {
	return func(time float64) float64 { return 0.5*(time*time) + iniVelocity*time + iniDisplacement }
}
