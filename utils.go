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

// Rotate Vector by given angle degrees in float64 and returns a new one
func Rotated(vector *Vector, angle_degrees float64) *Vector {
	new_vector := New(vector.X, vector.Y)
	new_vector.Rotate(angle_degrees)
	return new_vector
}

// GetAngleBetween returns the angle between two Vectors in float64
func GetAngleBetween(first *Vector, second *Vector) float64 {
	cross := first.X*second.Y - first.Y*second.X
	dot := first.X*second.X + first.Y*second.Y
	return math.Atan2(cross, dot) / (math.Pi / 180)
}

// GetDistance returns the distance between two Vectors in float64
func GetDistance(first *Vector, second *Vector) float64 {
	return math.Sqrt(math.Pow(first.X-second.X, 2) + math.Pow(first.Y-second.Y, 2))
}

