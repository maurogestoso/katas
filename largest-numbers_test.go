package katas

import (
	"reflect"
	"testing"
)

func TestLargestNumbers(t *testing.T) {
	type args struct {
		listOfLists [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Returns a slice of the largest numbers for a list of 1 list",
			args{listOfLists: [][]int{{1, 2, 3}}},
			[]int{3},
		},
		{
			"Returns a slice of the largest numbers for a list of 2 lists",
			args{listOfLists: [][]int{{3, 1, 2}, {-4, 6, 100}}},
			[]int{3, 100},
		},
		{
			"Returns a slice of the largest numbers for a list of 3 lists",
			args{listOfLists: [][]int{{2, 1, 3}, {-6, 6, 5}, {200, 300, 1000}}},
			[]int{3, 6, 1000},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestNumbers(tt.args.listOfLists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LargestNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
