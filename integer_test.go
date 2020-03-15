package roby

import "testing"

func TestExp(t *testing.T) {
	type args struct {
		num int
		n   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "1**0", args: args{num: 1, n: 0}, want: 1},
		{name: "2**0", args: args{num: 1, n: 0}, want: 1},
		{name: "1**1", args: args{num: 1, n: 1}, want: 1},
		{name: "2**1", args: args{num: 2, n: 1}, want: 2},
		{name: "2**2", args: args{num: 2, n: 2}, want: 4},
		{name: "3**2", args: args{num: 3, n: 2}, want: 9},
		{name: "4**2", args: args{num: 4, n: 2}, want: 16},
		{name: "2**3", args: args{num: 2, n: 3}, want: 8},
		{name: "3**3", args: args{num: 3, n: 3}, want: 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exp(tt.args.num, tt.args.n); got != tt.want {
				t.Errorf("Exp() = %v, want %v", got, tt.want)
			}
		})
	}
}
