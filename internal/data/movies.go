package data

import (
	"time"
)

/*
 * All the fields in our Movie struct are exported (i.e. start with a capital letter),
 * which is necessary for them to be visible to Go’s encoding/json package. Any fields
 * which aren’t exported won’t be included when encoding a struct to JSON.
 */
type Movie struct {
	ID        int64     `json:"id"`             // Unique integer ID for the movie
	CreatedAt time.Time `json:"-"`              // Timestamp for when the movie is added to our database
	Title     string    `json:"title"`          // Movie title
	Year      int32     `json:"year,omitempty"` // Movie release year

	// Use the Runtime type instead of int32. Note that the omitempty directive will
	// still work on this: if the Runtime field has the underlying value 0, then it will
	// be considered empty and omitted -- and the MarshalJSON() method we just made
	// won't be called at all.
	Runtime Runtime `json:"runtime,omitempty"`

	Genres []string `json:"genres,omitempty"` // Slice of genres for the movie (romance, comedy, etc.)

	// The version number starts at 1 and will be incremented each time the movie information is updated
	Version int32 `json:"version"`
}
