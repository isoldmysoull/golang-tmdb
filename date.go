package tmdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// tmdbDateLayouts lists all accepted date formats, in order of preference.
var tmdbDateLayouts = []string{
	time.DateOnly,
	time.RFC3339,
	"2006-01-02T15:04:05.000Z",
}

// TmdbDate represents TMDb date fields formatted as YYYY-MM-DD.
//
// Example:
//
//	"2020-07-10"
//
// It wraps time.Time to provide custom JSON marshaling and unmarshaling
// compatible with TMDb API responses.
type TmdbDate time.Time

// Time returns the underlying time.Time value.
func (d TmdbDate) Time() time.Time {
	return time.Time(d)
}

// IsZero reports whether the date is unset.
func (d TmdbDate) IsZero() bool {
	return d.Time().IsZero()
}

// String returns the TMDb formatted date (YYYY-MM-DD).
func (d TmdbDate) String() string {
	if d.IsZero() {
		return ""
	}

	return d.Time().Format(time.DateOnly)
}

// MarshalJSON implements json.Marshaler.
//
// Zero values are encoded as null.
func (d TmdbDate) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte("null"), nil
	}

	return json.Marshal(d.String())
}

// UnmarshalJSON implements json.Unmarshaler.
//
// Supported values:
//
//	"2020-07-10"
//	null
//	""
func (d *TmdbDate) UnmarshalJSON(b []byte) error {
	if d == nil {
		return fmt.Errorf("tmdbdate: nil receiver")
	}

	b = bytes.TrimSpace(b)

	if bytes.Equal(b, []byte("null")) {
		*d = TmdbDate{}
		return nil
	}

	var raw string
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("tmdbdate: invalid json value: %w", err)
	}

	if raw == "" {
		*d = TmdbDate{}
		return nil
	}

	var parsed time.Time
	var lastErr error
	for _, layout := range tmdbDateLayouts {
		parsed, lastErr = time.Parse(layout, raw)
		if lastErr == nil {
			break
		}
	}
	if lastErr != nil {
		return fmt.Errorf("tmdbdate: parse %q: %w", raw, lastErr)
	}

	*d = TmdbDate(parsed)
	return nil
}
