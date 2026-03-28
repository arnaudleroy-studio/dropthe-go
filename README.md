# dropthe-go

[![Go Reference](https://pkg.go.dev/badge/github.com/arnaudleroy-studio/dropthe-go.svg)](https://pkg.go.dev/github.com/arnaudleroy-studio/dropthe-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/arnaudleroy-studio/dropthe-go)](https://goreportcard.com/report/github.com/arnaudleroy-studio/dropthe-go)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Go client library for the [DropThe](https://dropthe.org) open data platform. Provides URL construction, text processing, and formatting utilities for building applications on top of the DropThe dataset.

## About DropThe

[DropThe](https://dropthe.org) is a data utility media network covering entertainment, technology, finance, and culture. The platform maintains a knowledge graph of 1.8M+ entities including movies, series, people, cryptocurrencies, and companies.

Key resources:

- [Data Catalog](https://dropthe.org/data/) - Browse the full entity dataset
- [Studio](https://dropthe.org/studio/) - Content creation and editorial tools
- [Methodology](https://dropthe.org/good/methodology/) - How DropThe researches and validates data

## Installation

```bash
go get github.com/arnaudleroy-studio/dropthe-go
```

Requires Go 1.21 or later.

## Usage

### URL Construction

Build entity URLs following DropThe routing conventions:

```go
package main

import (
    "fmt"
    dropthe "github.com/arnaudleroy-studio/dropthe-go"
)

func main() {
    // Build a movie URL
    url := dropthe.EntityURL(dropthe.TypeMovie, "inception-2010")
    fmt.Println(url) // https://dropthe.org/movies/inception-2010/

    // Build a person URL
    url = dropthe.EntityURL(dropthe.TypePerson, "christopher-nolan")
    fmt.Println(url) // https://dropthe.org/people/christopher-nolan/

    // Access platform pages
    fmt.Println(dropthe.DataURL())        // https://dropthe.org/data/
    fmt.Println(dropthe.StudioURL())      // https://dropthe.org/studio/
    fmt.Println(dropthe.MethodologyURL()) // https://dropthe.org/good/methodology/
}
```

### Entity Type

Use the `Entity` struct to represent items from the DropThe knowledge graph:

```go
movie := dropthe.Entity{
    ID:   "abc-123",
    Name: "Inception",
    Slug: "inception-2010",
    Type: dropthe.TypeMovie,
    Data: map[string]interface{}{
        "year": 2010,
        "director": "Christopher Nolan",
    },
}

fmt.Println(movie.URL()) // https://dropthe.org/movies/inception-2010/
```

### Text Processing

```go
// Slugify text for URL construction
slug := dropthe.Slugify("The Dark Knight (2008)")
fmt.Println(slug) // the-dark-knight-2008

// Handles unicode normalization
slug = dropthe.Slugify("Penelope Cruz")
fmt.Println(slug) // penelope-cruz

// Truncate text without cutting words
text := dropthe.TruncateText("The quick brown fox jumps over the lazy dog", 20)
fmt.Println(text) // The quick brown fox...
```

### Formatting

```go
// Currency formatting with comma separators
fmt.Println(dropthe.FormatCurrency(1234567.89, "$"))  // $1,234,567.89
fmt.Println(dropthe.FormatCurrency(42000.50, "EUR ")) // EUR 42,000.50

// Compact number formatting
fmt.Println(dropthe.FormatCompact(1500))    // 1.5K
fmt.Println(dropthe.FormatCompact(2300000)) // 2.3M
```

## Entity Types

The library supports all entity types in the DropThe knowledge graph:

| Type | Constant | Example URL |
|------|----------|-------------|
| Movies | `TypeMovie` | [/movies/inception-2010/](https://dropthe.org/movies/inception-2010/) |
| Series | `TypeSeries` | [/series/breaking-bad/](https://dropthe.org/series/breaking-bad/) |
| People | `TypePerson` | [/people/christopher-nolan/](https://dropthe.org/people/christopher-nolan/) |
| Cryptocurrencies | `TypeCryptocurrency` | [/cryptocurrencies/bitcoin/](https://dropthe.org/cryptocurrencies/bitcoin/) |
| Companies | `TypeCompany` | [/companies/apple-inc/](https://dropthe.org/companies/apple-inc/) |

## Related Resources

- [DropThe Platform](https://dropthe.org) - Main website
- [Data Catalog](https://dropthe.org/data/) - Browse all entities
- [Studio](https://dropthe.org/studio/) - Editorial tools
- [Methodology](https://dropthe.org/good/methodology/) - Research and validation process

## License

MIT License. See [LICENSE](LICENSE) for details.
