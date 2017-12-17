package day13

import (
	"reflect"
	"testing"
)

func Test_parseFirewall(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint]uint
		wantErr bool
	}{
		{name: "test", args: args{fname: "test.txt"}, want: map[uint]uint{0: 3, 1: 2, 4: 4, 6: 4}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseFirewall(tt.args.fname)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseFirewall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseFirewall() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSeverity(t *testing.T) {
	type args struct {
		fname  string
		offset uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{name: "test1", args: args{fname: "test.txt", offset: uint(0)}, want: 24},
		{name: "test2", args: args{fname: "test.txt", offset: uint(10)}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Severity(tt.args.fname, tt.args.offset); got != tt.want {
				t.Errorf("Severity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLowestSeverity(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{name: "test", args: args{fname: "test.txt"}, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowestSeverity(tt.args.fname); got != tt.want {
				t.Errorf("LowestSeverity() = %v, want %v", got, tt.want)
			}
		})
	}
}
