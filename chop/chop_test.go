package chop

import (
	"testing"
)

func TestChop(t *testing.T) {
	type args struct {
		num  int
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "corner case: empty list",
			args: args{
				num:  3,
				list: []int{},
			},
			want: -1,
		},
		{
			name: "corner case: one wrong element list",
			args: args{
				num:  3,
				list: []int{1},
			},
			want: -1,
		},
		{
			name: "corner case: one correct element list",
			args: args{
				num:  1,
				list: []int{1},
			},
			want: 0,
		},
		{
			name: "corner case: out of range low",
			args: args{
				num:  0,
				list: []int{1, 3, 5},
			},
			want: -1,
		},
		{
			name: "corner case: out of range high",
			args: args{
				num:  7,
				list: []int{1, 3, 5},
			},
			want: -1,
		},
		{
			name: "assert_equal(0,  chop(1, [1, 3, 5]))",
			args: args{
				num:  1,
				list: []int{1, 3, 5},
			},
			want: 0,
		},
		{
			name: "assert_equal(1,  chop(3, [1, 3, 5]))",
			args: args{
				num:  3,
				list: []int{1, 3, 5},
			},
			want: 1,
		},
		{
			name: "assert_equal(2,  chop(5, [1, 3, 5]))",
			args: args{
				num:  5,
				list: []int{1, 3, 5},
			},
			want: 2,
		},
		{
			name: "assert_equal(-1, chop(0, [1, 3, 5]))",
			args: args{
				num:  0,
				list: []int{1, 3, 5},
			},
			want: -1,
		},
		{
			name: "assert_equal(-1, chop(2, [1, 3, 5]))",
			args: args{
				num:  2,
				list: []int{1, 3, 5},
			},
			want: -1,
		},
		{
			name: "assert_equal(-1, chop(4, [1, 3, 5]))",
			args: args{
				num:  4,
				list: []int{1, 3, 5},
			},
			want: -1,
		},
		{
			name: "assert_equal(-1, chop(6, [1, 3, 5]))",
			args: args{
				num:  6,
				list: []int{1, 3, 5},
			},
			want: -1,
		},
		{
			name: "assert_equal(0,  chop(1, [1, 3, 5, 7]))",
			args: args{
				num:  1,
				list: []int{1, 3, 5, 7},
			},
			want: 0,
		},
		{
			name: "assert_equal(1,  chop(3, [1, 3, 5, 7]))",
			args: args{
				num:  3,
				list: []int{1, 3, 5, 7},
			},
			want: 1,
		},
		{
			name: "assert_equal(2,  chop(5, [1, 3, 5, 7]))",
			args: args{
				num:  5,
				list: []int{1, 3, 5, 7},
			},
			want: 2,
		},
		{
			name: "assert_equal(3,  chop(7, [1, 3, 5, 7]))",
			args: args{
				num:  7,
				list: []int{1, 3, 5, 7},
			},
			want: 3,
		},
		{
			name: "assert_equal(-1,  chop(0, [1, 3, 5, 7]))",
			args: args{
				num:  0,
				list: []int{1, 3, 5, 7},
			},
			want: -1,
		},
		{
			name: "assert_equal(-1,  chop(2, [1, 3, 5, 7]))",
			args: args{
				num:  2,
				list: []int{1, 3, 5, 7},
			},
			want: -1,
		},
		{
			name: "assert_equal(-1,  chop(4, [1, 3, 5, 7]))",
			args: args{
				num:  4,
				list: []int{1, 3, 5, 7},
			},
			want: -1,
		},
		{
			name: "assert_equal(-1,  chop(6, [1, 3, 5, 7]))",
			args: args{
				num:  6,
				list: []int{1, 3, 5, 7},
			},
			want: -1,
		},
		{
			name: "assert_equal(-1,  chop(8, [1, 3, 5, 7]))",
			args: args{
				num:  8,
				list: []int{1, 3, 5, 7},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chop(tt.args.num, tt.args.list); got != tt.want {
				t.Errorf("Chop() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chop2(tt.args.num, tt.args.list); got != tt.want {
				t.Errorf("Chop2() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chop3(tt.args.num, tt.args.list); got != tt.want {
				t.Errorf("Chop3() = %v, want %v", got, tt.want)
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chop4(tt.args.num, tt.args.list); got != tt.want {
				t.Errorf("Chop4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linearScale(t *testing.T) {
	type args struct {
		n    int
		list []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "scale 5 to interval [5,10]",
			args: args{
				n:    5,
				list: []int{5, 6, 7, 8, 9, 10},
			},
			want: 0,
		},
		{
			name: "scale 9 to interval [5,10]",
			args: args{
				n:    9,
				list: []int{5, 6, 7, 8, 9, 10},
			},
			want: 4,
		},
		{
			name: "scale 10 to interval [5,10]",
			args: args{
				n:    10,
				list: []int{5, 6, 7, 8, 9, 10},
			},
			want: 5,
		},
		{
			name: "scale 5 to interval [4,10]",
			args: args{
				n:    7,
				list: []int{4, 5, 6, 7, 8, 9, 10},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linearScale(tt.args.n, tt.args.list); got != tt.want {
				t.Errorf("linearScale() = %v, want %v", got, tt.want)
			}
		})
	}
}
