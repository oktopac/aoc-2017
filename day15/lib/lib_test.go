package day15

import (
	"reflect"
	"testing"
)

func Test_exhaustGenerator(t *testing.T) {
	type args struct {
		start  uint
		factor uint
		number uint
	}
	tests := []struct {
		name string
		args args
		want []uint
	}{
		{name: "generatorA", args: args{start: 65, factor: 16807, number: 5}, want: []uint{1092455, 1181022009, 245556042, 1744312007, 1352636452}},
		{name: "generatorB", args: args{start: 8921, factor: 48271, number: 5}, want: []uint{430625591, 1233683848, 1431495498, 137874439, 285222916}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exhaustGenerator(tt.args.start, tt.args.factor, tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exhaustGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}
