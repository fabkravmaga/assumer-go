package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/jessevdk/go-flags"
	"os"
)

const (
	version = "0.0.1"
)

type options struct {
	TargetAccount  string `short:"a" long:"target-account" description:"Target AWS account to assume into"`
	TargetRole     bool   `short:"r" long:"target-role" description:"The role in the target account"`
	ControlAccount string `short:"A" long:"control-account" description:"Control Plane AWS account"`
	ControlRole    string `short:"R" long:"control-role" description:"The role in the control account"`
	Region         string `short:"e" long:"region" description:"AWS region to operate in (default: us-east-1)"`
	Username       string `short:"u" long:"username" description:"Your IAM username"`
	Profile        string `short:"o" long:"profile" description:"Profile name from ~/.aws/credentials"`
	GUI            string `short:"g" long:"gui" description:"Open a web browser to the AWS console with these credentials"`
	Pry            string `short:"p" long:"pry" description:"Open a pry shell with these credentials"`
	Debug          string `short:"d" long:"debug" description:"Output debugging information"`
	Version        bool   `short:"v" long:"version" description:"Print the Version of the CLI"`
}

func getCredentials(profile string, region string) *aws.Config {
	creds := credentials.NewSharedCredentials("", profile)
	return aws.NewConfig().WithCredentials(creds).WithRegion(region)
}

func parseArgs() *options {
	opts := &options{}
	parser := flags.NewParser(opts, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		fmt.Println("Error parsing command line args")
		os.Exit(1)
	}
	return opts
}

func main() {

	opts := parseArgs()
	if opts.Version {
		fmt.Println("VERSION = ", version)
		os.Exit(1)
	}

	fmt.Println("Hello World")
}
