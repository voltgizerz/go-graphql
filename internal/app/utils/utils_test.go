package utils

import "testing"

func TestGetSafeInt(t *testing.T) {
	type args struct {
		num *int
	}

	number := 1
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "Test num is nil", args: args{num: nil}, want: 0},
		{name: "Test num is nil", args: args{num: &number}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSafeInt(tt.args.num); got != tt.want {
				t.Errorf("GetSafeInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
