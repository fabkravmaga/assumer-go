package main

import (
  "fmt"
  "bytes"
  "time"
  "os"
  "strings"

  "mfa"
  "assumer"
)

var _ time.Duration
var _ bytes.Buffer

func main() {
  fmt.Println("Assumer-go")
  mfa := mfa.GetMFA()
  user := os.Getenv("USER")
  region := "us-west-2"
  controlRole := ""
  roleARN_base := ""
  s := []string{roleARN_base, user}
  roleARN := strings.Join(s, "")
  fmt.Println("Assuming Role:", roleARN)
  assumer.AssumeTargetAccount(region, controlRole, roleARN, mfa)
}
