package utils

import (
	"testing"
)

func TestMD5Encode(t *testing.T) {
	Md5Encode("dsadas")

}

func TestMakePassword(t *testing.T) {
	MakePassword("Dsa", "sda")
}

func TestMd5Encode(t *testing.T) {
	Md5Encode("fds")
}

func TestValidPassword(t *testing.T) {
	ValidPassword("Fds", "FDsa", "FDSa")
}
