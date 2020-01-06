package main

import "testing"

func Test_fuelConsumption(t *testing.T) {
	type args struct {
		neededLiters int
		haveLiters   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Needed Liters is 0", args{0, 4}, 0},
		{"Needed Liters is less than 0", args{ -1, 4}, 0},
		{"Both of them are more than 0", args{5,1}, 18},
		{"Both of them are more than 0", args{5,3}, 54},
		{"Have Liters is 0", args{4, 0}, 0},
		{"Have Liters is less than 0", args{ 4, -1}, 0},
		{"Both of them are  0", args{0,0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fuelConsumption(tt.args.neededLiters, tt.args.haveLiters); got != tt.want {
				t.Errorf("fuelConsumption() = %v, want %v", got, tt.want)
			}
		})
	}
}