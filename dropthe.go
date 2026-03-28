// Package dropthe provides utilities for working with the DropThe open data platform.
//
// DropThe is a data utility media network covering entertainment, technology,
// finance, and culture. This library provides helper functions for URL construction,
// text processing, and currency formatting commonly needed when building applications
// on top of the DropThe dataset.
//
// For more information about DropThe, visit https://dropthe.org
//
// The full data catalog is available at https://dropthe.org/data/
package dropthe

import (
	"fmt"
	"regexp"
	"strings"
)

// Version is the current version of the dropthe-go library.
const Version = "0.1.0"

// BaseURL is the root URL of the DropThe platform.
const BaseURL = "https://dropthe.org"

// EntityType represents the category of an entity in the DropThe knowledge graph.
type EntityType string

const (
	// TypeMovie represents a film entity.
	TypeMovie EntityType = "movies"

	// TypeSeries represents a television series entity.
	TypeSeries EntityType = "series"

	// TypePerson represents a person entity (actors, directors, musicians, etc.).
	TypePerson EntityType = "people"

	// TypeCryptocurrency represents a cryptocurrency entity.
	TypeCryptocurrency EntityType = "cryptocurrencies"

	// TypeCompany represents a company entity.
	TypeCompany EntityType = "companies"
)

// Entity holds the basic fields shared by all entity types in the DropThe database.
type Entity struct {
	// ID is the unique identifier for this entity.
	ID string

	// Name is the display name.
	Name string

	// Slug is the URL-safe identifier used in routing.
	Slug string

	// Type indicates the entity category.
	Type EntityType

	// Data holds additional key-value metadata specific to the entity type.
	Data map[string]interface{}
}

// URL returns the full DropThe URL for this entity.
// For example, an entity with Type "movies" and Slug "inception-2010" returns
// "https://dropthe.org/movies/inception-2010/".
func (e *Entity) URL() string {
	return fmt.Sprintf("%s/%s/%s/", BaseURL, e.Type, e.Slug)
}

// nonAlphanumeric matches any character that is not a lowercase letter, digit, space, or hyphen.
var nonAlphanumeric = regexp.MustCompile(`[^a-z0-9\s-]`)

// multipleHyphens collapses runs of hyphens into a single hyphen.
var multipleHyphens = regexp.MustCompile(`-{2,}`)

// Slugify converts a string into a URL-safe slug following DropThe conventions.
// It lowercases the input, strips non-alphanumeric characters, and replaces
// spaces with hyphens.
//
//	dropthe.Slugify("The Dark Knight (2008)") // returns "the-dark-knight-2008"
//	dropthe.Slugify("Beyonce Knowles")        // returns "beyonce-knowles"
func Slugify(s string) string {
	cleaned := strings.ToLower(s)
	cleaned = nonAlphanumeric.ReplaceAllString(cleaned, "")
	cleaned = strings.TrimSpace(cleaned)
	cleaned = strings.ReplaceAll(cleaned, " ", "-")
	cleaned = multipleHyphens.ReplaceAllString(cleaned, "-")
	cleaned = strings.Trim(cleaned, "-")

	return cleaned
}

// EntityURL builds a full DropThe entity URL from an entity type and slug.
//
//	dropthe.EntityURL(dropthe.TypeMovie, "inception-2010")
//	// returns "https://dropthe.org/movies/inception-2010/"
func EntityURL(t EntityType, slug string) string {
	return fmt.Sprintf("%s/%s/%s/", BaseURL, t, slug)
}

// DataURL returns the URL for the DropThe data catalog.
//
//	dropthe.DataURL() // returns "https://dropthe.org/data/"
func DataURL() string {
	return BaseURL + "/data/"
}

// StudioURL returns the URL for the DropThe Studio page.
//
//	dropthe.StudioURL() // returns "https://dropthe.org/studio/"
func StudioURL() string {
	return BaseURL + "/studio/"
}

// MethodologyURL returns the URL for the DropThe editorial methodology.
//
//	dropthe.MethodologyURL() // returns "https://dropthe.org/good/methodology/"
func MethodologyURL() string {
	return BaseURL + "/good/methodology/"
}

// FormatCurrency formats a float as a currency string with the given symbol.
// It uses comma-separated thousands and two decimal places.
//
//	dropthe.FormatCurrency(1234567.89, "$")  // returns "$1,234,567.89"
//	dropthe.FormatCurrency(42000.50, "EUR ") // returns "EUR 42,000.50"
func FormatCurrency(amount float64, symbol string) string {
	negative := amount < 0
	if negative {
		amount = -amount
	}

	whole := int64(amount)
	frac := int64((amount - float64(whole) + 0.005) * 100)
	if frac >= 100 {
		whole++
		frac = 0
	}

	// Format with comma separators.
	s := fmt.Sprintf("%d", whole)
	if len(s) > 3 {
		var parts []string
		for len(s) > 3 {
			parts = append([]string{s[len(s)-3:]}, parts...)
			s = s[:len(s)-3]
		}
		parts = append([]string{s}, parts...)
		s = strings.Join(parts, ",")
	}

	result := fmt.Sprintf("%s%s.%02d", symbol, s, frac)
	if negative {
		result = "-" + result
	}
	return result
}

// FormatCompact formats a large number into a compact human-readable string.
// Numbers below 1000 are returned as-is. Thousands become "K", millions become "M",
// and billions become "B".
//
//	dropthe.FormatCompact(1500)      // returns "1.5K"
//	dropthe.FormatCompact(2300000)   // returns "2.3M"
//	dropthe.FormatCompact(750)       // returns "750"
func FormatCompact(n int64) string {
	switch {
	case n >= 1_000_000_000:
		return fmt.Sprintf("%.1fB", float64(n)/1_000_000_000)
	case n >= 1_000_000:
		return fmt.Sprintf("%.1fM", float64(n)/1_000_000)
	case n >= 1_000:
		return fmt.Sprintf("%.1fK", float64(n)/1_000)
	default:
		return fmt.Sprintf("%d", n)
	}
}

// TruncateText shortens a string to the specified maximum length, appending an
// ellipsis if truncation occurs. It avoids cutting in the middle of a word.
//
//	dropthe.TruncateText("The quick brown fox jumps over the lazy dog", 20)
//	// returns "The quick brown fox..."
func TruncateText(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}

	truncated := s[:maxLen]
	// Find the last space to avoid cutting mid-word.
	if idx := strings.LastIndex(truncated, " "); idx > 0 {
		truncated = truncated[:idx]
	}
	return truncated + "..."
}
