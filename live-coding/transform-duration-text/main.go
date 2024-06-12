package main

import (
	"fmt"
	"strings"
)

func formatDuration(seconds int) string {
	if seconds == 0 {
		return "now"
	}

	units := []struct {
		name  string
		value int
	}{
		{"year", 365 * 24 * 60 * 60},
		{"day", 24 * 60 * 60},
		{"hour", 60 * 60},
		{"minute", 60},
		{"second", 1},
	}

	parts := []string{}
	for _, unit := range units {
		if seconds >= unit.value {
			count := seconds / unit.value
			seconds %= unit.value
			part := fmt.Sprintf("%d %s", count, unit.name)
			if count > 1 {
				part += "s"
			}
			parts = append(parts, part)
		}
	}

	// Modified output formatting
	if len(parts) > 1 {
		last := parts[len(parts)-1]
		parts = parts[:len(parts)-1] // Remove the last element
		return fmt.Sprintf("%s and %s", strings.Join(parts, ", "), last)
	}
	return parts[0]
}

func main() {
	// Example usage
	fmt.Println(formatDuration(62))   // Output: 1 minute and 2 seconds
	fmt.Println(formatDuration(3662)) // Output: 1 hour, 1 minute and 2 seconds
}
