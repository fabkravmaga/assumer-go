package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
)

func main() {
	var mfa string
	var err error
	opts := parseArgs()
	if opts.Version {
		fmt.Println("VERSION = ", version)
		os.Exit(1)
	}

	if opts.Region == "" {
		opts.Region = "us-east-1"
	}

	serial, err := getSerial(opts.Username, opts.ControlAccount)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	controlRole := getCompositeRole(opts.ControlRole, opts.ControlAccount)
	targetRole := getCompositeRole(opts.TargetRole, opts.TargetAccount)
	fmt.Println(targetRole + "\nvia\n" + controlRole)

	if opts.MFA == "" {
		mfa, err = getMFA()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		mfa = opts.MFA
	}

	creds := credentials.NewSharedCredentials("", opts.Profile)
	config := aws.NewConfig().WithCredentials(creds).WithRegion(opts.Region)
	sess := session.New(config)

	controlCreds, err := getControlCreds(opts.ControlAccount, controlRole, serial, mfa, sess)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	getTempCreds(opts.TargetAccount, targetRole, controlCreds, opts.Region)
}
