// Package validator provides IP address validation and filtering.
package validator

import (
	"net"

	"github.com/rs/zerolog/log"
)

// ValidateIPs checks each IP string and returns a slice of valid IPs
// along with the count of skipped (invalid) entries.
// Invalid entries are logged as warnings.
func ValidateIPs(ips []string) ([]string, int) {
	var valid []string
	skipped := 0

	for _, ipStr := range ips {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			log.Warn().Str("ip", ipStr).Msg("skipping invalid IP entry")
			skipped++
			continue
		}
		valid = append(valid, ipStr)
	}

	return valid, skipped
}
