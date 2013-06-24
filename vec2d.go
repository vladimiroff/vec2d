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

// Package vec2d offers a two-dimensional vector implementation for Go.
package vec2d

import "math"

// Vector type defines vector using exported float64 values: X and Y
type Vector struct {
	X, Y float64
}

// Init initializes already created vector
func (self *Vector) Init(x, y float64) *Vector {
	self.X = x
	self.Y = y
	return self
}

// New returns a new vector
func New(x, y float64) *Vector {
	return new(Vector).Init(x, y)
}

// IsEqual compares а Vector with another and returns true if they're equal
func (self *Vector) IsEqual(other *Vector) bool {
	return self.X == other.X && self.Y == other.Y
}

// Angle returns the Vector's angle in float64
func (self *Vector) Angle() float64 {
	return math.Atan2(self.Y, self.X) / (math.Pi / 180)
}

// SetAngle changes Vector's angle using vector rotation
func (self *Vector) SetAngle(angle_degrees float64) {
	self.X = self.Length()
	self.Y = 0.0
	self.Rotate(angle_degrees)
}

// Length returns... well the Vector's length
func (self *Vector) Length() float64 {
	return math.Sqrt(math.Pow(self.X, 2) + math.Pow(self.Y, 2))
}

// SetLength changes Vector's length, which obviously changes
// the values of Vector.X and Vector.Y
func (self *Vector) SetLength(value float64) {
	length := self.Length()
	self.X *= value / length
	self.Y *= value / length
}

// Rotate Vector by given angle degrees in float64
func (self *Vector) Rotate(angle_degrees float64) {
	radians := (math.Pi / 180) * angle_degrees
	sin := math.Sin(radians)
	cos := math.Cos(radians)

	x := self.X*cos - self.Y*sin
	y := self.X*sin + self.Y*cos
	self.X = x
	self.Y = y
}

// Rotated does the same as Vector.Rotate, except does not change
// the given instance, but creates and returns a new one, instead.
func (self *Vector) Rotated(angle_degrees float64) *Vector {
	vector := New(self.X, self.Y)
	vector.Rotate(angle_degrees)
	return vector
}

// GetAngleBetween returns the angle between to Vectors in float64
func (self *Vector) GetAngleBetween(other *Vector) float64 {
	cross := self.X*other.Y - self.Y*other.X
	dot := self.X*other.X + self.Y*other.Y
	return math.Atan2(cross, dot) / (math.Pi / 180)
}

// GetDistance returns the distance betweeen to Vectors in float64
func (self *Vector) GetDistance(other *Vector) float64 {
	return math.Sqrt(math.Pow(self.X-other.X, 2) + math.Pow(self.Y-other.Y, 2))
}

// Add creates new Vector and sets its X and Y to the sum of these in self and other
func (self *Vector) Add(other *Vector) *Vector {
	return New(self.X+other.X, self.Y+other.Y)
}

// Div returns new Vector adding given value to self's X and Y
func (self *Vector) AddToFloat64(value float64) *Vector {
	return New(self.X+value, self.Y+value)
}

// Add creates a new Vector and sets its X and Y to the difference self-other
func (self *Vector) Sub(other *Vector) *Vector {
	return New(self.X-other.X, self.Y-other.Y)
}

// Div returns a new Vector, taking given value from self's X and Y
func (self *Vector) SubToFloat64(value float64) *Vector {
	return New(self.X-value, self.Y-value)
}

// Mul returns a new Vector by multiplication of self and other
func (self *Vector) Mul(other *Vector) *Vector {
	return New(self.X*other.X, self.Y*other.Y)
}

// Div returns a new Vector using multiplication of self's X and Y by given value
func (self *Vector) MulToFloat64(value float64) *Vector {
	return New(self.X*value, self.Y*value)
}

// Div returns a new Vector by dividing self and other
func (self *Vector) Div(other *Vector) *Vector {
	return New(self.X/other.X, self.Y/other.Y)
}

// Div returns a new Vector by dividing self's X and Y by given value
func (self *Vector) DivToFloat64(value float64) *Vector {
	return New(self.X/value, self.Y/value)
}
