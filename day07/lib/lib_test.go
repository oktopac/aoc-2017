package day07

import "testing"

func TestFindTop(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test", args: args{fname: "test.txt"}, want: "tknk"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTop(tt.args.fname); got != tt.want {
				t.Errorf("FindTop() = %v, want %v", got, tt.want)
			}
		})
	}
}
