package main

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		i int
		j int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{"asd", args{1,2}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.i, tt.args.j); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
