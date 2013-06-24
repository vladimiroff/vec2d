// Copyright 2013 Kiril Vladimiroff
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package vec2d_test

import (
	"github.com/Vladimiroff/vec2d"
	"fmt"
)

// In this example we create new vector with X=2 and Y=4,
// then we get the length and angle of it.
func ExampleVector_create() {
	vector := vec2d.New(2.0, 4.0)
	length := vector.Length()
	angle := vector.Angle()

	fmt.Printf("Vector{X: %v, Y: %v}\n", vector.X, vector.Y)
	fmt.Printf("Length: %v\n", length)
	fmt.Printf("Angle: %v\n",angle)
	// Output:
	// Vector{X: 2, Y: 4}
	// Length: 4.47213595499958
	// Angle: 63.434948822922
}

// We're able to edit our Vector by changing its X, Y, angle or length.
// Obviously, if we change X or Y, then the angle and length
// most definately will be changes, also.
func ExampleVector_edit() {
	vector := vec2d.New(2.0, 4.0)
	vector.X = 4
	length := vector.Length()
	angle := vector.Angle()

	fmt.Printf("Vector{X: %v, Y: %v}\n", vector.X, vector.Y)
	fmt.Printf("Length: %v\n", length)
	fmt.Printf("Angle: %v\n",angle)
	// Output:
	// Vector{X: 4, Y: 4}
	// Length: 5.656854249492381
	// Angle: 45
}

func ExampleVector_GetAngleBetween() {
	vector := vec2d.New(2.0, 4.0)
	other := vec2d.New(5.0, 12.0)
	angle_between := vector.GetAngleBetween(other)

	fmt.Printf("Vector{X: %v, Y: %v}\n", vector.X, vector.Y)
	fmt.Printf("Vector{X: %v, Y: %v}\n", other.X, other.Y)
	fmt.Printf("Angle between: %v\n", angle_between)
	// Output:
	// Vector{X: 2, Y: 4}
	// Vector{X: 5, Y: 12}
	// Angle between: 3.945186229037563
}

func ExampleVector_GetDistance() {
	vector := vec2d.New(2.0, 4.0)
	other := vec2d.New(5.0, 12.0)
	distance := vector.GetDistance(other)

	fmt.Printf("Vector{X: %v, Y: %v}\n", vector.X, vector.Y)
	fmt.Printf("Vector{X: %v, Y: %v}\n", other.X, other.Y)
	fmt.Printf("Distance between: %v\n", distance)
	// Output:
	// Vector{X: 2, Y: 4}
	// Vector{X: 5, Y: 12}
	// Distance between: 8.54400374531753
}

// Note that operations like Add, Sub, Mul and Div return new Vector
func ExampleVector_Add() {
	vector := vec2d.New(2.0, 4.0)
	other := vec2d.New(3.0, 5.0)
	result := vector.Add(other)

	fmt.Printf("Vector + Other vector gave new vector with X: %v, Y: %v\n", result.X, result.Y)
	// Output: Vector + Other vector gave new vector with X: 5, Y: 9
}

// In this example we rotate our vector by 2 degrees, and this will change X and Y
// of the current instance, without creating a new one.
//
// Note: If you don't want to change the current instance take a look at Vector.Rotated
func ExampleVector_Rotate() {
	vector := vec2d.New(2.0, 4.0)
	vector.Rotate(2.0)

	fmt.Printf("Vector.Rotate(2.0) changed X to %v and Y to %v\n", vector.X, vector.Y)
	// Output: Vector.Rotate(2.0) changed X to 1.8591836672281876 and Y to 4.067362301481385
}
