package math

import (
	"testing"
)

func TestSubtract(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1 : subtracting 1 from 2",
			args: args{x: 2, y: 1},
			want: 1,
		},
		{
			name: "case 2 : subtracting 3 from 4",
			args: args{x: 4, y: 3},
			want: 1,
		},
		{
			name: "case 3 : subtracting 5 from 6",
			args: args{x: 6, y: 5},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Subtract(tt.args.x, tt.args.y); gotRes != tt.want {
				t.Errorf("Subtract() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case 1 : adding 1 and 2",
			args: args{x: 1, y: 2},
			want: 3,
		},
		{
			name: "case 2 : adding 3 and 4",
			args: args{x: 3, y: 4},
			want: 7,
		},
		{
			name: "case 3 : adding 5 and 6",
			args: args{x: 5, y: 6},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := Add(tt.args.x, tt.args.y); gotRes != tt.want {
				t.Errorf("Add() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}
