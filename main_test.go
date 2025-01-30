package main

import "testing"

func Test_unpack(t *testing.T) {
	type args struct {
		a1 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"a1"}, "a"},
		{"test2", args{"a0b"}, "b"},
		{"test3", args{"a2b3"}, "aabbb"},
		{"test4", args{"a2\n3"}, "aa\n\n\n"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if got = unpack(tt.args.a1); got != tt.want {
				t.Errorf("unpack() = %q, want %q", got, tt.want)
			} else {
				t.Logf("unpack(%q) = %q", tt.args.a1, got)
			}
		})
	}
}
