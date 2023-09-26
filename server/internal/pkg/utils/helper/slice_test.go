package helper

import (
	"reflect"
	"testing"
)

func TestSliceStringEqualUnordered(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "helper slice equal 1",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a", "b"},
			},
			want: true,
		},
		{
			name: "helper slice equal 2",
			args: args{
				a: []string{"a", "b"},
				b: []string{"b", "a"},
			},
			want: true,
		},
		{
			name: "helper slice equal 3",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a", "c"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceStringEqualUnordered(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SliceStringEqualUnordered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceIntEqualUnordered(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "helper slice int equal 1",
			args: args{
				a: []int{1, 2},
				b: []int{1, 2},
			},
			want: true,
		},
		{
			name: "helper slice int equal 2",
			args: args{
				a: []int{1, 2},
				b: []int{2, 1},
			},
			want: true,
		},
		{
			name: "helper slice int equal 3",
			args: args{
				a: []int{1, 2},
				b: []int{1, 3},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceIntEqualUnordered(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SliceIntEqualUnordered() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicates(t *testing.T) {
	type args struct {
		s []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "helper remove duplicates 1",
			args: args{
				s: []string{"a", "b", "c", "c"},
			},
			want: []string{"a", "b", "c"},
		},
		{
			name: "helper remove duplicates 2",
			args: args{
				s: []string{"a", "b", "c"},
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicates(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "helper TestIntersect 1",
			args: args{
				a: []string{"a", "b", "c"},
				b: []string{"a", "c"},
			},
			want: []string{"a", "c"},
		},
		{
			name: "helper TestIntersect 2",
			args: args{
				a: []string{"a", "b"},
				b: []string{"a", "c"},
			},
			want: []string{"a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIndexOfStrSlice(t *testing.T) {
	type args struct {
		s []string
		a string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 int
	}{
		// TODO: Add test cases.
		{
			name: "helper IndexOfStrSlice 1",
			args: args{
				s: []string{"a", "b"},
				a: "a",
			},
			want:  true,
			want1: 0,
		},
		{
			name: "helper IndexOfStrSlice 1",
			args: args{
				s: []string{"a", "b", "c"},
				a: "c",
			},
			want:  true,
			want1: 2,
		},
		{
			name: "helper IndexOfStrSlice 1",
			args: args{
				s: []string{"a", "b", "c"},
				a: "e",
			},
			want:  false,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := IndexOfStrSlice(tt.args.s, tt.args.a)
			if got != tt.want {
				t.Errorf("IndexOfStrSlice() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("IndexOfStrSlice() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
