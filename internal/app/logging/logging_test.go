package logging

import (
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
)

func TestConvertLogLevel(t *testing.T) {
	tests := []struct {
		name string
		args string
		want slog.Level
	}{
		{
			name: "debug case",
			args: "debug",
			want: slog.LevelDebug,
		},
		{
			name: "warn case",
			args: "warn",
			want: slog.LevelWarn,
		},
		{
			name: "error case",
			args: "error",
			want: slog.LevelError,
		},
		{
			name: "info case",
			args: "info",
			want: slog.LevelInfo,
		},
		{
			name: "empty string",
			args: "",
			want: slog.LevelInfo,
		},
		{
			name: "invalid case",
			args: "something",
			want: slog.LevelInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertLogLevel(tt.args)
			assert.Equal(t, tt.want, got)
		})
	}
}
