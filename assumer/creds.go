package main

import (
	"bufio"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func getMFA() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter MFA: ")
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text, err
}

func getCredentials(profile string, region string) *aws.Config {
	creds := credentials.NewSharedCredentials("", profile)
	return aws.NewConfig().WithCredentials(creds).WithRegion(region)
}

func getSerial(user string, account string) (string, error) {
	if user == "" {
		user = os.Getenv("USER")
	}
	fmt.Println(user + " is assuming")
	return "arn:aws:iam::" + account + ":mfa/" + user, nil
}

func getCompositeRole(role string, account string) string {
	return "arn:aws:iam::" + account + ":role/" + role
}

func getControlCreds(account string, role string, serial string, token string, sess *session.Session) (credentials.Value, error) {
	stsCreds := credentials.NewCredentials(&stscreds.AssumeRoleProvider{
		Client:       sts.New(sess),
		Duration:     time.Hour,
		RoleARN:      role,
		ExpiryWindow: 5 * time.Minute,
		SerialNumber: aws.String(serial),
		TokenCode:    aws.String(token),
	})

	key, err := stsCreds.Get()
	if err != nil {
		fmt.Println("ERROR ", err)
		return key, err
	}
	return key, nil
}

func getTempCreds(account string, role string, control credentials.Value, region string) {
	tempDir, _ := ioutil.TempDir("/var/folders", "")
	tempFile, _ := ioutil.TempFile(tempDir, "")
	creds := credentials.NewStaticCredentials(control.AccessKeyID, control.SecretAccessKey, control.SessionToken)
	config := aws.NewConfig().WithCredentials(creds).WithRegion(region)
	sess := session.New(config)
	stsCreds := credentials.NewCredentials(&stscreds.AssumeRoleProvider{
		Client:       sts.New(sess),
		Duration:     time.Hour,
		RoleARN:      role,
		ExpiryWindow: 5 * time.Minute,
	})

	key, err := stsCreds.Get()
	if err != nil {
		fmt.Println("ERROR ", err)
	}
	values := ("export AWS_REGION=" + region + "\n" +
		"export AWS_ACCESS_KEY_ID=" + key.AccessKeyID + "\n" +
		"export AWS_SECRET_ACCESS_KEY=" + key.SecretAccessKey + "\n" +
		"export AWS_SESSION_TOKEN=" + key.SessionToken + "\n")
	file := []byte(values)
	_, err = tempFile.Write(file)
	if err != nil {
		fmt.Println("There was an error writing the file... " + err.Error())
		os.Exit(1)
	}
	fmt.Println("To import these values into the shell, source '" + tempFile.Name() + "'")

}
