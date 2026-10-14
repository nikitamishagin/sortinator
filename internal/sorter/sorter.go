// Package sorter provides numeric IP sorting, reverse, and deduplication.
package sorter

import (
	"bytes"
	"net"
	"sort"
)

// ipLess returns true if IP a is numerically less than IP b.
func ipLess(a, b net.IP) bool {
	return bytes.Compare(a, b) < 0
}

// SortIPs sorts a list of IP strings numerically by their binary representation.
// If reverse is true, sorts in descending order.
// Returns sorted IPs and the number of duplicates dropped (if dedup needed).
func SortIPs(ips []string, reverse bool) ([]string, int) {
	// Convert to 16-byte representations for unified sorting
	type ipEntry struct {
		str   string
		bytes net.IP
	}

	entries := make([]ipEntry, 0, len(ips))
	for _, ipStr := range ips {
		b := net.ParseIP(ipStr).To16()
		if b == nil {
			continue
		}
		entries = append(entries, ipEntry{str: ipStr, bytes: b})
	}

	sort.Slice(entries, func(i, j int) bool {
		return ipLess(entries[i].bytes, entries[j].bytes)
	})

	if reverse {
		for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
			entries[i], entries[j] = entries[j], entries[i]
		}
	}

	result := make([]string, len(entries))
	for i, e := range entries {
		result[i] = e.str
	}

	return result, 0
}

// Deduplicate removes duplicate IP strings from the sorted slice.
// Returns the deduplicated slice and the count of dropped duplicates.
func Deduplicate(ips []string) ([]string, int) {
	if len(ips) == 0 {
		return ips, 0
	}

	dropped := 0
	seen := make(map[string]bool, len(ips))
	result := make([]string, 0, len(ips))

	for _, ip := range ips {
		if seen[ip] {
			dropped++
			continue
		}
		seen[ip] = true
		result = append(result, ip)
	}

	return result, dropped
}
