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

func TestUtilsRotated(t *testing.T) {
	vector := New(2.0, 4.0)
	result := Rotated(vector, 2.0)

	if result.X != 1.8591836672281876 && result.Y != 4.067362301481385 {
		t.Error("Vector.Rotate(2.0) gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("Rotated did not return a copy object")
	}
}

func TestUtilsGetAngleBetween(t *testing.T) {
	vector := New(2.0, 4.0)
	other := New(5.0, 12.0)
	angle_between := GetAngleBetween(vector, other)

	if angle_between != 3.945186229037563 {
		t.Error("Angle between vector and other happens to be:", angle_between)
	}
}

func TestUtilsGetDistance(t *testing.T) {
	vector := New(2.0, 4.0)
	other := New(5.0, 12.0)
	distance := GetDistance(vector, other)

	if distance != 8.54400374531753 {
		t.Error("Distance between vector and other happens to be:", distance)
	}
}

func TestUtilsCollect(t *testing.T) {
	vector := New(2.0, 4.0)
	other := New(3.0, 5.0)
	result := Collect(vector, other)

	if result.X != 5.0 && result.Y != 9.0 {
		t.Error("Vector + Other vector gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("Collect did not return a copy object")
	}
}

func TestUtilsCollectToFloat64(t *testing.T) {
	vector := New(2.0, 4.0)
	result := CollectToFloat64(vector, 3.0)

	if result.X != 5.0 && result.Y != 6.0 {
		t.Error("Vector + float64 gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("CollectToFloat64 did not return a copy object")
	}
}

func TestUtilsSub(t *testing.T) {
	vector := New(4.0, 4.0)
	other := New(2.0, 3.0)
	result := Sub(vector, other)

	if result.X != 2.0 && result.Y != 1.0 {
		t.Error("Vector - Other vector gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("Sub did not return a copy object")
	}
}

func TestUtilsSubToFloat64(t *testing.T) {
	vector := New(2.0, 4.0)
	result := SubToFloat64(vector, 1.0)

	if result.X != 1.0 && result.Y != 3.0 {
		t.Error("Vector - float64 gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("SubToFloat64 did not return a copy object")
	}
}

func TestUtilsMul(t *testing.T) {
	vector := New(2.0, 4.0)
	other := New(3.0, 5.0)
	result := Mul(vector, other)

	if result.X != 6.0 && result.Y != 20.0 {
		t.Error("Vector * Other gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("Mul did not return a copy object")
	}
}

func TestUtilsMulToFloat64(t *testing.T) {
	vector := New(2.0, 4.0)
	result := MulToFloat64(vector, 3.0)

	if result.X != 6.0 && result.Y != 12.0 {
		t.Error("Vector * Other gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("MulToFloat64 did not return a copy object")
	}
}

func TestUtilsDiv(t *testing.T) {
	vector := New(4.0, 8.0)
	other := New(2.0, 4.0)
	result := Div(vector, other)

	if result.X != 2.0 && result.Y != 2.0 {
		t.Error("Vector * Other gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("Div did not return a copy object")
	}
}

func TestUtilsDivToFloat64(t *testing.T) {
	vector := New(2.0, 4.0)
	result := DivToFloat64(vector, 2.0)

	if result.X != 1.0 && result.Y != 2.0 {
		t.Error("Vector * Other gave X:", result.X, " Y:", result.Y)
	}

	if vector == result {
		t.Error("DivToFloat64 did not return a copy object")
	}
}
