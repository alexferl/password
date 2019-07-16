package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

var length int
var numDigits int
var numSymbols int
var noUpper bool
var allowRepeat bool
var lowerLetters string
var upperLetters string
var digits string
var symbols string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a password",
	Run: func(cmd *cobra.Command, args []string) {
		generate()
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().IntVarP(&length, "length", "l", 64,
		"total number of characters in the password")
	generateCmd.PersistentFlags().IntVarP(&numDigits, "num-digits", "d", 10,
		"number of digits to include in the result")
	generateCmd.PersistentFlags().IntVarP(&numSymbols, "num-symbols", "s", 10,
		"number of symbols to include in the result")
	generateCmd.PersistentFlags().BoolVarP(&noUpper, "no-upper", "u", false,
		"excludes uppercase letters")
	generateCmd.PersistentFlags().BoolVarP(&allowRepeat, "allow-repeat", "r", false,
		"allows characters to repeat")
	generateCmd.PersistentFlags().StringVarP(&lowerLetters, "lower-letters", "", password.LowerLetters,
		"list of permitted lowercase letters")
	generateCmd.PersistentFlags().StringVarP(&upperLetters, "upper-letters", "", password.UpperLetters,
		"list of permitted uppercase letters")
	generateCmd.PersistentFlags().StringVarP(&digits, "digits", "", password.Digits,
		"list of permitted digits")
	generateCmd.PersistentFlags().StringVarP(&symbols, "symbols", "", password.Symbols,
		"list of permitted symbols")

	_ = viper.BindPFlag("length", generateCmd.PersistentFlags().Lookup("length"))
	_ = viper.BindPFlag("num-digits", generateCmd.PersistentFlags().Lookup("num-digits"))
	_ = viper.BindPFlag("num-symbols", generateCmd.PersistentFlags().Lookup("num-symbols"))
	_ = viper.BindPFlag("no-upper", generateCmd.PersistentFlags().Lookup("no-upper"))
	_ = viper.BindPFlag("allow-repeat", generateCmd.PersistentFlags().Lookup("allow-repeat"))
	_ = viper.BindPFlag("lower-letters", generateCmd.PersistentFlags().Lookup("lower-letters"))
	_ = viper.BindPFlag("upper-letters", generateCmd.PersistentFlags().Lookup("upper-letters"))
	_ = viper.BindPFlag("digits", generateCmd.PersistentFlags().Lookup("digits"))
	_ = viper.BindPFlag("symbols", generateCmd.PersistentFlags().Lookup("symbols"))
}

func generate() {
	gi := &password.GeneratorInput{}

	if lowerLetters != "" {
		gi.LowerLetters = viper.GetString("lower-letter")
	}

	if upperLetters != "" {
		gi.UpperLetters = viper.GetString("upper-letters")
	}

	if digits != "" {
		gi.Digits = viper.GetString("digits")
	}

	if symbols != "" {
		gi.Symbols = viper.GetString("symbols")
	}

	gen, err := password.NewGenerator(gi)
	if err != nil {
		fmt.Printf("Error creating generator: %v\n", err)
		os.Exit(1)
	}

	res, err := gen.Generate(
		viper.GetInt("length"),
		viper.GetInt("num-digits"),
		viper.GetInt("num-symbols"),
		viper.GetBool("no-upper"),
		viper.GetBool("allow-repeat"),
	)
	if err != nil {
		fmt.Printf("Error generating password: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", res)
}
