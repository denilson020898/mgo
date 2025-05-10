package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {
	switch n {
	case "pass":
		n = "password"
		break
	case "ps":
		n = "password"
		break
	}
	return pflag.NormalizedName(n)
}

func main() {
	pflag.StringP("name", "n", "Denilson", "Name parameter")
	pflag.StringP("password", "p", "hardtoguess", "Password")
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	name := viper.GetString("name")
	password := viper.GetString("password")

	fmt.Println(name, password)

	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	if val != nil {
		fmt.Println("GOMAXPROCS:", val)
	}

	viper.Set("GOMAXPROCS", 16)
	val = viper.Get("GOMAXPROCS")
	if val != nil {
		fmt.Println("GOMAXPROCS:", val)
	}
}
