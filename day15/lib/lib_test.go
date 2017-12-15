package day15

import (
	"reflect"
	"testing"
)

func Test_exhaustGenerator(t *testing.T) {
	type args struct {
		generator <-chan uint64
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{name: "generatorA", args: args{generator: generator(65, 16807, 5)}, want: []uint64{1092455, 1181022009, 245556042, 1744312007, 1352636452}},
		{name: "generatorB", args: args{generator: generator(8921, 48271, 5)}, want: []uint64{430625591, 1233683848, 1431495498, 137874439, 285222916}},
		{name: "generatorB", args: args{generator: generator2(65, 16807, 5, 4)}, want: []uint64{1352636452, 1992081072, 530830436, 1980017072, 740335192}},
		{name: "generatorB", args: args{generator: generator2(8921, 48271, 5, 8)}, want: []uint64{1233683848, 862516352, 1159784568, 1616057672, 412269392}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exhaustGenerator(tt.args.generator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exhaustGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_match(t *testing.T) {
	type args struct {
		val1 uint64
		val2 uint64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "match", args: args{val1: 245556042, val2: 1431495498}, want: true},
		{name: "match", args: args{val1: 1092455, val2: 430625591}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := match(tt.args.val1, tt.args.val2); got != tt.want {
				t.Errorf("match() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countMatches(t *testing.T) {
	type args struct {
		generator1 <-chan uint64
		generator2 <-chan uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// {name: "test", args: args{generator1: generatorA(65, 40e6), generator2: generatorB(8921, 40e6)}, want: 588},
		// {name: "chal1", args: args{generator1: GeneratorA(65, 40e6), generator2: GeneratorB(8921, 40e6)}, want: 588},
		{name: "chal2", args: args{generator1: GeneratorA2(65, 5e6), generator2: GeneratorB2(8921, 5e6)}, want: 309},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountMatches(tt.args.generator1, tt.args.generator2); got != tt.want {
				t.Errorf("countMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
