package utils

import "testing"

func TestMD5Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5Encode(tt.args.data); got != tt.want {
				t.Errorf("MD5Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakePassword(t *testing.T) {
	type args struct {
		plainpwd string
		salt     string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakePassword(tt.args.plainpwd, tt.args.salt); got != tt.want {
				t.Errorf("MakePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMd5Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Md5Encode(tt.args.data); got != tt.want {
				t.Errorf("Md5Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidPassword(t *testing.T) {
	type args struct {
		plainpwd string
		salt     string
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidPassword(tt.args.plainpwd, tt.args.salt, tt.args.password); got != tt.want {
				t.Errorf("ValidPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
