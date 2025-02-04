package packer

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"test0", "", ""},
		{"test1", "a11", "a"},
		{"test2", "a0b", "b"},
		{"test3", "a2b3", "aabbb"},
		{"test4", "a2\n3", "aa\n\n\n"},
		{"test5", "a2b 3", "aab   "},
	}
	for _, tt := range tests {
		a1 := tt.args
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if got = Unpack(a1); got != tt.want {
				if got == "" {
					t.Errorf("WRONG calling: unpack(%q)", a1)
				} else {
					t.Errorf("unpack() != %q, want %q", got, tt.want)
				}
			} else {
				t.Logf("unpack(%q) == %q", a1, got)
			}
		})
	}
}

func TestPack(t *testing.T) {
	type args struct {
		a1 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test0", args{""}, ""},
		{"test1", args{"a1"}, "a"},
		{"test2", args{"aabbb"}, "a2b3"},
		{"test3", args{"a_bbb\n\n\t   "}, "a_b3\n2\t 3"},
	}
	for _, tt := range tests {
		a1 := tt.args.a1
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if got = Pack(a1); got != tt.want {
				if got == "" {
					t.Errorf("pack(%q)", a1)
					// t.SkipNow()
					return
				}
				t.Errorf("pack() != %q, want %q", got, tt.want)
			} else {
				t.Logf("pack(%q) == %q", a1, got)
			}
		})
	}
}
