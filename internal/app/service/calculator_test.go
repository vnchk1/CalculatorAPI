package service

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name    string
		args    []int
		want    int
		wantErr bool
	}{
		{name: "negative", args: []int{-1, -2, -3}, want: -6, wantErr: false},
		{name: "positive", args: []int{1, 2, 3}, want: 6, wantErr: false},
		{name: "zero", args: []int{0}, want: 0, wantErr: false},
		{name: "empty", args: []int{}, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Sum(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum() error = %v, wantErr %v", err, tt.wantErr)
			}
			if result != tt.want {
				t.Errorf("Sum() = %v, want %v", result, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name    string
		args    []int
		want    int
		wantErr bool
	}{
		{name: "negative", args: []int{-1, -2, -4}, want: -8, wantErr: false},
		{name: "positive", args: []int{1, 2, 4}, want: 8, wantErr: false},
		{name: "zero", args: []int{0}, want: 0, wantErr: false},
		{name: "empty", args: []int{}, want: 0, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Multiply(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiply() error = %v, wantErr %v", err, tt.wantErr)
			}
			if result != tt.want {
				t.Errorf("Multiply() = %v, want %v", result, tt.want)
			}
		})
	}
}
