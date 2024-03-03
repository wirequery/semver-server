package versioning

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	want := SemanticVersion{Major: 0, Minor: 1, Patch: 0}
	if got := New(); !reflect.DeepEqual(got, want) {
		t.Errorf("New() = %v, want %v", got, want)
	}
}

func TestIncrementMajor(t *testing.T) {
	type args struct {
		s SemanticVersion
	}
	tests := []struct {
		name string
		args args
		want SemanticVersion
	}{
		{
			name: "all versions are reset, except major, which increments",
			args: args{SemanticVersion{0, 1, 2}},
			want: SemanticVersion{1, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncrementMajor(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncrementMajor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncrementMinor(t *testing.T) {
	type args struct {
		s SemanticVersion
	}
	tests := []struct {
		name string
		args args
		want SemanticVersion
	}{
		{
			name: "patch is reset, minor increments and major remains the same",
			args: args{SemanticVersion{1, 1, 2}},
			want: SemanticVersion{1, 2, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncrementMinor(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncrementMinor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncrementPatch(t *testing.T) {
	type args struct {
		s SemanticVersion
	}
	tests := []struct {
		name string
		args args
		want SemanticVersion
	}{
		{
			name: "patch is incremented",
			args: args{SemanticVersion{1, 1, 2}},
			want: SemanticVersion{1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncrementPatch(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IncrementPatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateString(t *testing.T) {
	type args struct {
		s SemanticVersion
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "dots are added between major, minor and patch",
			args: args{SemanticVersion{1, 2, 3}},
			want: "1.2.3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateString(tt.args.s); got != tt.want {
				t.Errorf("CreateString() = %v, want %v", got, tt.want)
			}
		})
	}
}
