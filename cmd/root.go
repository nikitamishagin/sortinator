package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/mishagin/sortinator/internal/sorter"
	"github.com/mishagin/sortinator/internal/validator"
)

var (
	version    = "dev"
	outputFile string
	reverse    bool
	unique     bool
	verbose    bool
)

// IPConfig represents the YAML structure.
type IPConfig struct {
	IPs []string `yaml:"ips"`
}

var rootCmd = &cobra.Command{
	Use:   "sortinator [file]",
	Short: "Sort IPs from a YAML file",
	Long: `Sortinator reads a YAML file with an ips: key, sorts the IPs numerically (IPv4 and IPv6), and prints the result.

Examples:
  sortinator test.yaml
  sortinator test.yaml -r -u -v
  sortinator test.yaml -f output.yaml`,
	Version: version,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			_ = cmd.Usage()
			return nil
		}

		filePath := args[0]

		// Read YAML file
		data, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal().Err(err).Str("file", filePath).Msg("failed to read file")
			return err
		}

		// Parse YAML
		var config IPConfig
		if err := yaml.Unmarshal(data, &config); err != nil {
			log.Fatal().Err(err).Str("file", filePath).Msg("failed to parse YAML")
			return err
		}

		// Validate IPs
		validIPs, skipped := validator.ValidateIPs(config.IPs)

		// Sort IPs
		sortedIPs, _ := sorter.SortIPs(validIPs, reverse)

		// Deduplicate if requested
		var dropped int
		if unique {
			sortedIPs, dropped = sorter.Deduplicate(sortedIPs)
		}

		// Build output
		var output strings.Builder
		for _, ip := range sortedIPs {
			output.WriteString(ip + "\n")
		}

		// Verbose stats
		if verbose {
			log.Info().
				Int("sorted", len(sortedIPs)).
				Int("skipped", skipped).
				Int("dropped", dropped).
				Msg("processing complete")
		}

		// Output
		if outputFile != "" {
			if err := os.WriteFile(outputFile, []byte(output.String()), 0644); err != nil {
				log.Fatal().Err(err).Str("file", outputFile).Msg("failed to write output file")
				return err
			}
			log.Info().Str("file", outputFile).Msg("output written to file")
		} else {
			fmt.Print(output.String())
		}

		return nil
	},
}

func init() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	rootCmd.Flags().StringVarP(&outputFile, "file", "f", "", "write output to specified file")
	rootCmd.Flags().BoolVarP(&reverse, "reverse", "r", false, "sort in descending order")
	rootCmd.Flags().BoolVarP(&unique, "unique", "u", false, "remove duplicate IPs")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show statistics summary")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("error")
		os.Exit(1)
	}
}
