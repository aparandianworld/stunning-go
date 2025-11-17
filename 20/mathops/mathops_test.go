package mathops

import (
	"testing"
)

func TestAdd(t *testing.T) {

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 2.0, 3.0, 5.0},
		{"negative numbers", -2.0, -3.0, -5.0},
		{"mixed numbers", -2.0, 3.0, 1.0},
		{"zeros", 0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Add(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Add(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestSub(t *testing.T) {

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 2.0, 3.0, -1.0},
		{"negative numbers", -2.0, -3.0, 1.0},
		{"mixed numbers", -2.0, 3.0, -5.0},
		{"zeros", 0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Sub(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Sub(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 2.0, 3.0, 6.0},
		{"negative numbers", -2.0, -3.0, 6.0},
		{"mixed numbers", -2.0, 3.0, -6.0},
		{"zeros", 0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Mul(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Mul(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {

	tests := []struct {
		name string
		a, b float64
		want float64
	}{
		{"positive numbers", 2.0, 3.0, 0.6666666666666666},
		{"negative numbers", -2.0, -3.0, 0.6666666666666666},
		{"mixed numbers", -2.0, 3.0, -0.6666666666666666},
		{"zeros", 0.0, 0.0, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Div(tt.a, tt.b)
			if err != nil {
				t.Errorf("error division by zero")
			}
			if got != tt.want {
				t.Errorf("Div(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
