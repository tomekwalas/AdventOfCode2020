package day12

import (
	"testing"
)

func TestFindBoatDirection(t *testing.T) {
	type args struct {
		inputs []direction
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find boat directions",
			args: args{
				inputs: []direction{
					{
						value: "F",
						units: 10,
					},
					{
						value: "N",
						units: 3,
					},
					{
						value: "F",
						units: 7,
					},
					{
						value: "R",
						units: 90,
					},
					{
						value: "F",
						units: 11,
					},
				},
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBoatDirection(tt.args.inputs); got != tt.want {
				t.Errorf("FindBoatDirection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindBoatDirectionFromWaypoint(t *testing.T) {
	type args struct {
		inputs []direction
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Should find boat directions with waypoint",
			args: args{
				inputs: []direction{
					{
						value: "F",
						units: 10,
					},
					{
						value: "N",
						units: 3,
					},
					{
						value: "F",
						units: 7,
					},
					{
						value: "R",
						units: 90,
					},
					{
						value: "F",
						units: 11,
					},
				},
			},
			want: 286,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBoatDirectionFromWaypoint(tt.args.inputs); got != tt.want {
				t.Errorf("FindBoatDirectionFromWaypoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
