package day01

import (
	"testing"
)

func Test_Solve(t *testing.T) {
	type args struct {
		challenge string
	}
	type test struct {
		name    string
		args    args
		want    int
		wantErr bool
	}
	tests := []test{
		test{name: "1122", args: args{challenge: "1122"}, want: 3, wantErr: false},
		test{name: "1111", args: args{challenge: "1111"}, want: 4, wantErr: false},
		test{name: "1234", args: args{challenge: "1234"}, want: 0, wantErr: false},
		test{name: "91212129", args: args{challenge: "91212129"}, want: 9, wantErr: false},
		test{name: "a", args: args{challenge: "a"}, want: 0, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solve(tt.args.challenge)
			if (err != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve2(t *testing.T) {
	type args struct {
		challenge string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "1212", args: args{challenge: "1212"}, want: 6, wantErr: false},
		{name: "1221", args: args{challenge: "1221"}, want: 0, wantErr: false},
		{name: "123425", args: args{challenge: "123425"}, want: 4, wantErr: false},
		{name: "123123", args: args{challenge: "123123"}, want: 12, wantErr: false},
		{name: "12131415", args: args{challenge: "12131415"}, want: 4, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Solve2(tt.args.challenge)
			if (err != nil) != tt.wantErr {
				t.Errorf("Solve2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Solve2() = %v, want %v", got, tt.want)
			}
		})
	}
}
