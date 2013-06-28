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

func TestNew(t *testing.T) {
	vector := New(2.0, 4.0)
	if vector.X != 2.0 && vector.Y != 4.0 {
		t.Error("Creating new vector has failed")
	}
}

func TestIsEqual(t *testing.T) {
	vector := New(2.0, 4.0)
	the_same_vector := New(2.0, 4.0)
	completely_different_vector := New(1.0, 3.0)

	if !vector.IsEqual(the_same_vector) {
		t.Error("vector should be equal to the_save_vector")
	}

	if vector.IsEqual(completely_different_vector) {
		t.Error("vector should not be equal to completely_different_vector")
	}
}

func TestAngle(t *testing.T) {
	vector := New(2.0, 4.0)
	zero_vector := New(0.0, 0.0)

	if vector.Angle() != 63.434948822922 {
		t.Error("vector.Angle() appeared to be", vector.Angle())
	}

	if zero_vector.Angle() != 0.0 {
		t.Error("zero_vector.Angle() appeared to be", vector.Angle())
	}
}

func TestSetAngle(t *testing.T) {
	vector := New(2.0, 4.0)
	vector.SetAngle(45.0)

	if vector.Angle() != 45.0 {
		t.Error("After changing the angle vector.Angle() appeared to be", vector.Angle())
	}
	if vector.X != 3.1622776601683795 && vector.Y != 3.162277660168379 {
		t.Error("After changing the angle coordinates are X:", vector.X, " Y:", vector.Y)
	}
}

func TestLength(t *testing.T) {
	vector := New(2.0, 4.0)

	if vector.Length() != 4.47213595499958 {
		t.Error("vector.Length() appeared to be", vector.Length())
	}
}

func TestSetLength(t *testing.T) {
	vector := New(2.0, 4.0)
	vector.SetLength(5.0)

	if vector.Length() != 5.0 {
		t.Error("Vector's length is not the given one, but", vector.Length())
	}
	if vector.X != 2.23606797749979 && vector.Y != 4.47213595499958 {
		t.Error("After changing the length coordinates are X:", vector.X, " Y:", vector.Y)
	}
}

func TestAdd(t *testing.T) {
	vector := New(2.0, 4.0)
	other := New(3.0, 5.0)
	result := vector.Add(other)

	if result.X != 5.0 && result.Y != 9.0 {
		t.Error("Vector + Other vector gave X:", result.X, " Y:", result.Y)
	}
}

func TestAddToFloat64(t *testing.T) {
	vector := New(2.0, 4.0)
	result := vector.AddToFloat64(3.0)

	if result.X != 5.0 && result.Y != 6.0 {
		t.Error("Vector + float64 gave X:", result.X, " Y:", result.Y)
	}
}

func TestSub(t *testing.T) {
	vector := New(4.0, 4.0)
	other := New(2.0, 3.0)
	result := vector.Sub(other)

	if result.X != 2.0 && result.Y != 1.0 {
		t.Error("Vector - Other vector gave X:", result.X, " Y:", result.Y)
	}
}

func TestSubToFloat64(t *testing.T) {
	vector := New(2.0, 4.0)
	result := vector.SubToFloat64(1.0)

	if result.X != 1.0 && result.Y != 3.0 {
		t.Error("Vector - float64 gave X:", result.X, " Y:", result.Y)
	}
}

func TestMul(t *testing.T) {
	vector := New(2.0, 4.0)
	other := New(3.0, 5.0)
	result := vector.Mul(other)

	if result.X != 6.0 && result.Y != 20.0 {
		t.Error("Vector * Other gave X:", result.X, " Y:", result.Y)
	}
}

func TestMulToFloat64(t *testing.T) {
	vector := New(2.0, 4.0)
	result := vector.MulToFloat64(3.0)

	if result.X != 6.0 && result.Y != 12.0 {
		t.Error("Vector * Other gave X:", result.X, " Y:", result.Y)
	}
}

func TestDiv(t *testing.T) {
	vector := New(4.0, 8.0)
	other := New(2.0, 4.0)
	result := vector.Div(other)

	if result.X != 2.0 && result.Y != 2.0 {
		t.Error("Vector * Other gave X:", result.X, " Y:", result.Y)
	}
}

func TestDivToFloat64(t *testing.T) {
	vector := New(2.0, 4.0)
	result := vector.DivToFloat64(2.0)

	if result.X != 1.0 && result.Y != 2.0 {
		t.Error("Vector * Other gave X:", result.X, " Y:", result.Y)
	}
}

func TestRotate(t *testing.T) {
	vector := New(2.0, 4.0)
	vector.Rotate(2.0)

	if vector.X != 1.8591836672281876 && vector.Y != 4.067362301481385 {
		t.Error("Vector.Rotate(2.0) gave X:", vector.X, " Y:", vector.Y)
	}
}
