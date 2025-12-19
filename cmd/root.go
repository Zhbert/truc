/*
Copyright © 2025 Zhbert zhbert@yandex.ru

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package cmd

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

const (
	Reset      = "\033[0m"
	Red        = "\033[31m"
	Green      = "\033[32m"
	DarkYellow = "\033[33m"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "truc",
	Short: "A simple CLI tool for creating true relative URLs",
	Long: `truc is a simple CLI tool for creating true relative URLs.

Allows you to quickly calculate a relative link based on the source and destination.`,
	Run: func(cmd *cobra.Command, args []string) {
		source, _ := cmd.Flags().GetString("source")
		target, _ := cmd.Flags().GetString("target")
		verbose, _ := cmd.Flags().GetBool("verbose")
		cpToClipboard, _ := cmd.Flags().GetBool("copy")
		ignoreDifferentDomains, _ := cmd.Flags().GetBool("stop-if-different-domains")

		if verbose {
			fmt.Printf("%15s %-s\n", "Source URL:", source)
			fmt.Printf("%15s %-s\n", "Target URL:", target)
		}

		sourceDomain, err := extractDomain(source)
		if err != nil {
			fmt.Println(err)
		}
		targetDomain, _ := extractDomain(target)
		if err != nil {
			fmt.Println(err)
		}

		if targetDomain != sourceDomain {
			if ignoreDifferentDomains {
				fmt.Println(Red + "URL domains do not match!" + Reset)
				return
			} else {
				fmt.Println(DarkYellow + "URL domains do not match!" + Reset)
			}
		}

		source, err = removeProtocol(source)
		if err != nil {
			fmt.Println(err)
		}
		target, err = removeProtocol(target)
		if err != nil {
			fmt.Println(err)
		}

		sourceParts := strings.Split(source, "/")
		targetParts := strings.Split(target, "/")

		discrepancyPosition := 0
		for i, part := range sourceParts {
			if i < len(targetParts) {
				if targetParts[i] == part {
					if verbose {
						fmt.Printf("%s%d: %s%s\n", Green, i, part, Reset)
					}
				} else {
					if verbose {
						if i <= len(targetParts) {
							fmt.Printf("%s%d: %s -> %s%s\n", Red, i, part, targetParts[i], Reset)
						} else {
							fmt.Printf("%s%d: %s%s\n", Red, i, part, Reset)
						}
					}
					if discrepancyPosition == 0 {
						discrepancyPosition = i
					}
				}
			}
		}
		delta := len(sourceParts) - discrepancyPosition - 1
		if verbose {
			fmt.Printf("Levels back: %d\n", delta)
		}

		var resultUrl string
		if delta == 0 {
			resultUrl += "./"
		} else {
			for i := 0; i < delta; i++ {
				resultUrl += "../"
			}
		}
		for i := discrepancyPosition; i < len(targetParts); i++ {
			resultUrl += targetParts[i]
			if i < len(targetParts)-1 {
				resultUrl += "/"
			}
		}
		fmt.Printf("%s %s%-s\n", "Result URL:", Green, resultUrl)
		if cpToClipboard {
			err := clipboard.WriteAll(resultUrl)
			if err != nil {
				fmt.Println(err)
			}
			if verbose {
				fmt.Println(Reset + "The URL has been copied to the clipboard.")
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.truc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("source", "s", "", "Specify the source URL")
	rootCmd.Flags().StringP("target", "t", "", "Specify the target URL")
	rootCmd.Flags().BoolP("verbose", "v", false, "Enable full log")
	rootCmd.Flags().BoolP("copy", "c", false, "Copy to clipboard")
	rootCmd.Flags().BoolP("stop-if-different-domains", "", false, "Enabling stops for different domains")
}
