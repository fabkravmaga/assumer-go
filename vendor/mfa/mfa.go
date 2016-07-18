package mfa

import (
  "bufio"
  "strings"
  "fmt"
  "os"
  "bytes"
  "time"
)

var _ time.Duration
var _ bytes.Buffer

func GetMFA()(mfa string) {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("Enter 6 digit MFA: ")
  mfa_input, _ := reader.ReadString('\n')
  mfa_input_no_nextline := strings.Trim(mfa_input, "\n")
  fmt.Println(mfa_input_no_nextline)
  return mfa_input_no_nextline
}
