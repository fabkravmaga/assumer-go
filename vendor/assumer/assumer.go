package assumer

import (
  "bytes"
  "fmt"
  "time"

  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/sts"

)

var _ time.Duration
var _ bytes.Buffer
var debug = true

func AssumeTargetAccount(
                          region string,
                          roleArn string,
                          serialNumber string,
                          mfa string)(AssumeRoleOutput *sts.AssumeRoleOutput) {
  // First hope to control role
  controlRoleOutput := STS_AssumeControlRole(region, roleArn, serialNumber, mfa)
  if debug{
    fmt.Println(controlRoleOutput)
  }

  // Second hope to target role
  targetRoleRegion := region
  targetRoleArn := "arn:aws:iam::857967218396:role/Security/kaos-security-audit"

  targetRoleOutput := STS_AssumeTargetRole(targetRoleRegion, targetRoleArn)

  if debug{
    fmt.Println(targetRoleOutput)
  }
  return targetRoleOutput
}

func STS_AssumeControlRole(region string, roleArn string, serialNumber string, mfa string) (AssumeRoleOutput *sts.AssumeRoleOutput){

    svc := sts.New(session.New(), aws.NewConfig().WithRegion(region))

    params := &sts.AssumeRoleInput{
      RoleArn:         aws.String(roleArn),       // Required
      RoleSessionName: aws.String("AssumedRole"), // Required
      DurationSeconds: aws.Int64(3600),           // 1 hour
      SerialNumber:    aws.String(serialNumber),
      TokenCode:       aws.String(mfa),
    }

    output, err := svc.AssumeRole(params)

    if err != nil {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
        return
    }

    // Pretty-print the response data.
    if debug{
      fmt.Println(output)
    }
    return output
}

func STS_AssumeTargetRole(region string, roleArn string) (AssumeRoleOutput *sts.AssumeRoleOutput){

    svc := sts.New(session.New(), aws.NewConfig().WithRegion(region))

    params := &sts.AssumeRoleInput{
      RoleArn:         aws.String(roleArn),       // Required
      RoleSessionName: aws.String("AssumedRole"), // Required
      DurationSeconds: aws.Int64(3600),           // 1 hour
    }

    output, err := svc.AssumeRole(params)

    if err != nil {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        fmt.Println(err.Error())
        return
    }

    // Pretty-print the response data.
    if debug{
      fmt.Println(output)
    }
    return output
}
