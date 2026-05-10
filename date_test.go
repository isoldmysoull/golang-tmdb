package tmdb

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTmdbDate_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     string
		want      string
		wantZero  bool
		wantError bool
	}{
		{
			name:     "valid date",
			input:    `"2020-07-10"`,
			want:     "2020-07-10",
			wantZero: false,
		},
		{
			name:     "null value",
			input:    `null`,
			wantZero: true,
		},
		{
			name:     "empty string",
			input:    `""`,
			wantZero: true,
		},
		{
			name:      "invalid format",
			input:     `"10-07-2020"`,
			wantError: true,
		},
		{
			name:      "invalid json type",
			input:     `123`,
			wantError: true,
		},
		{
			name:      "invalid date",
			input:     `"2020-99-99"`,
			wantError: true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var got TmdbDate
			err := json.Unmarshal([]byte(tt.input), &got)

			if tt.wantError {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got.IsZero() != tt.wantZero {
				t.Fatalf("expected zero=%v, got zero=%v", tt.wantZero, got.IsZero())
			}

			if !tt.wantZero && got.String() != tt.want {
				t.Fatalf("expected %q, got %q", tt.want, got.String())
			}
		})
	}
}

func TestTmdbDate_MarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name  string
		input TmdbDate
		want  string
	}{
		{
			name:  "valid date",
			input: TmdbDate(time.Date(2020, 7, 10, 0, 0, 0, 0, time.UTC)),
			want:  `"2020-07-10"`,
		},
		{
			name:  "zero value",
			input: TmdbDate{},
			want:  `null`,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			b, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(b) != tt.want {
				t.Fatalf("expected %s, got %s", tt.want, string(b))
			}
		})
	}
}

func TestTmdbDate_String(t *testing.T) {
	t.Parallel()

	t.Run("zero value", func(t *testing.T) {
		var d TmdbDate

		if got := d.String(); got != "" {
			t.Fatalf("expected empty string, got %q", got)
		}
	})

	t.Run("formatted value", func(t *testing.T) {
		d := TmdbDate(time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC))

		if got := d.String(); got != "2024-01-15" {
			t.Fatalf("expected 2024-01-15, got %q", got)
		}
	})
}
