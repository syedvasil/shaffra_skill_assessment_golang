package config

import (
	"reflect"
	"testing"
)

func TestConfig(t *testing.T) {
	t.Parallel()
	validConf := Config()

	tests := []struct {
		name string
		want *AppConfig
	}{
		{
			"valid",
			validConf,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Config(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config() = %v, want %v", got, tt.want)
			}
		})
	}
}
