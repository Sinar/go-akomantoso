package state_assembly

import (
	"reflect"
	"testing"
)

func Test_extractDebaters(t *testing.T) {
	type args struct {
		allLines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractDebaters(tt.args.allLines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractDebaters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_looksLikeRep(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := looksLikeRep(tt.args.line)
			if got != tt.want {
				t.Errorf("looksLikeRep() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("looksLikeRep() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
