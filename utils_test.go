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

import "testing"

func TestRotated(t *testing.T) {
	vector := New(2.0, 4.0)
	result := Rotated(vector, 2.0)

	if result.X != 1.8591836672281876 && result.Y != 4.067362301481385 {
		t.Error("Vector.Rotate(2.0) gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("Vector.Rotated did not return a copy object")
	}
}

func TestGetAngleBetween(t *testing.T) {
	vector := New(2.0, 4.0)
	other := New(5.0, 12.0)
	angle_between := GetAngleBetween(vector, other)

	if angle_between != 3.945186229037563 {
		t.Error("Angle between vector and other happens to be:", angle_between)
	}
}

func TestGetDistance(t *testing.T) {
	vector := New(2.0, 4.0)
	other := New(5.0, 12.0)
	distance := GetDistance(vector, other)

	if distance != 8.54400374531753 {
		t.Error("Distance between vector and other happens to be:", distance)
	}
}
