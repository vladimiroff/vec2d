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
package vec2d

import "math"

// Rotated rotates a Vector by given angle degrees in float64 and returns a new one
func Rotated(vector *Vector, angle_degrees float64) *Vector {
	new_vector := New(vector.X, vector.Y)
	new_vector.Rotate(angle_degrees)
	return new_vector
}

// GetAngleBetween returns the angle between two Vectors in float64
func GetAngleBetween(first, second *Vector) float64 {
	cross := first.X*second.Y - first.Y*second.X
	dot := first.X*second.X + first.Y*second.Y
	return math.Atan2(cross, dot) / (math.Pi / 180)
}

// GetDistance returns the distance between two Vectors in float64
func GetDistance(first, second *Vector) float64 {
	return math.Sqrt(math.Pow(first.X-second.X, 2) + math.Pow(first.Y-second.Y, 2))
}

// Collect creates new Vector and sets its X and Y to the sum of these in first and second
func Collect(first, second *Vector) *Vector {
	return New(first.X+second.X, first.Y+second.Y)
}

// Collect returns new Vector adding given value to first's X and Y
func CollectToFloat64(vector *Vector, value float64) *Vector {
	return New(vector.X+value, vector.Y+value)
}

// Add creates a new Vector and sets its X and Y to the difference first-second
func Sub(first, second *Vector) *Vector {
	return New(first.X-second.X, first.Y-second.Y)
}

// Div returns a new Vector, taking given value from first's X and Y
func SubToFloat64(vector *Vector, value float64) *Vector {
	return New(vector.X-value, vector.Y-value)
}

// Mul returns a new Vector by multiplication of first and second
func Mul(first, second *Vector) *Vector {
	return New(first.X*second.X, first.Y*second.Y)
}

// Div returns a new Vector using multiplication of first's X and Y by given value
func MulToFloat64(vector *Vector, value float64) *Vector {
	return New(vector.X*value, vector.Y*value)
}

// Div returns a new Vector by dividing first and second
func Div(first, second *Vector) *Vector {
	return New(first.X/second.X, first.Y/second.Y)
}

// Div returns a new Vector by dividing first's X and Y by given value
func DivToFloat64(vector *Vector, value float64) *Vector {
	return New(vector.X/value, vector.Y/value)
}
