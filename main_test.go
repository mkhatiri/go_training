package main

import "testing"

func TestAddition(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"33", args{3, 3}, 6},
		{"12", args{1, 2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Addition(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Addition() = %v, want %v", got, tt.want)
			}
		})
	}
}
