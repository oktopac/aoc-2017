package day02

import (
	"reflect"
	"testing"
)

func TestParseSpreadsheet(t *testing.T) {
	type args struct {
		fname string
	}
	tests := []struct {
		name    string
		args    args
		want    spreadsheet
		wantErr bool
	}{
		{name: "test", args: args{fname: "test1.txt"}, want: [][]int{{1, 2, 3}, {4, 5, 6}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSpreadsheet(tt.args.fname)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSpreadsheet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseSpreadsheet() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	type args struct {
		sheet spreadsheet
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{sheet: [][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}}}, want: 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checksum(tt.args.sheet); got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChecksum2(t *testing.T) {
	type args struct {
		sheet spreadsheet
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", args: args{sheet: [][]int{{5, 9, 2, 8}, {9, 4, 7, 3}, {3, 8, 6, 5}}}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Checksum2(tt.args.sheet); got != tt.want {
				t.Errorf("Checksum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
