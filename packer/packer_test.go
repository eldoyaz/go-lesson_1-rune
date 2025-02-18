package packer

import "testing"

func TestUnpack(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"a0b", "b"},
		{"a2b3", "aabbb"},
		{"a2\n3", "aa\n\n\n"},
		{"a2b 3", "aab   "},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {

			got, err := Unpack(tt.input)
			if err != nil {
				t.Errorf("текст ошибки: %q", err.Error())
				t.SkipNow()
			}
			if got != tt.want {
				t.Errorf("Unpack(%q) != %q", tt.input, got)
			} else {
				t.Logf("[OK] Unpack(%q) == %q", tt.input, tt.want)
			}
		})
	}

}

func TestUnpackErrors(t *testing.T) {

	errTests := []struct {
		wrongStr string
	}{ // каждый тест должен возвращать ОШИБКУ
		{"a11a"},
		{"a12"},
		{"1a"},
	}

	for _, tt := range errTests {
		t.Run(tt.wrongStr, func(t *testing.T) {

			if _, err := Unpack(tt.wrongStr); err != nil {
				t.Logf("текст ошибки: %q", err.Error())
			} else {
				t.Error("[error expected]")
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
