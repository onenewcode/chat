package repository

import (
	"chat/config"
	"testing"
)

func TestInitDB(t *testing.T) {
	type args struct {
		c config.Config
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitDB(tt.args.c)
		})
	}
}
