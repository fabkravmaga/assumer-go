assumer-go
==================
## Overview ##
The Assumer gem is an interface allowing the user to assume role into an account.  Credentials can be loaded anywhere the AWS STS Client knows to load credentials from (ENV, profile, etc.)

## Content ##
Assumer will assume-role on a target account and generate temporary API keys for that account.

## Installation ##

1. wget the tgz file for your OS
  * `wget https://github.com/ccloes-intuit/assumer-go/releases/download/0.0.1/assumer-0.0.1-darwin-amd64.tgz`
  * `wget https://github.com/ccloes-intuit/assumer-go/releases/download/0.0.1/assumer-0.0.1-linux-amd64.tgz`
  * `wget https://github.com/ccloes-intuit/assumer-go/releases/download/0.0.1/assumer-0.0.1-windows-amd64.tgz`
2. untar the file
  * `tar xzvf assumer-*.tgz`
3. run the binary
  * `./assumer -h`

### Build from Source ###

1. Clone the repository
1. Change directory into the newly-cloned repository
  * `cd assumer`
1. Run the build script (assumes that you have go 1.6 installed, will cross compile for linux, MAC OSX, and Windows)
  * `./build`
1. Alternately, you can simply run MAKE 
  * `make`

## Options
```
Usage:
  assumer [OPTIONS]

Application Options:
  -a, --target-account=  Target AWS account to assume into
  -r, --target-role=     The role in the target account
  -A, --control-account= Control Plane AWS account
  -R, --control-role=    The role in the control account
  -e, --region=          AWS region to operate in (default: us-east-1)
  -t, --token-code=      MFA code
  -u, --username=        Your IAM username
  -o, --profile=         Profile name from ~/.aws/credentials
  -d, --debug=           Output debugging information
  -v, --version          Print the Version of the CLI

Help Options:
  -h, --help             Show this help message

```

## Example build
```
[15:54][ccloes@wdhl07e9fbdff:~/go/src/github.com/ccloes-intuit/assumer-go][ git:(upstream ⚡ initial)]$ ./build
GO ROOT: 

==> Formatting source code.
==> Running tests.
?       github.com/ccloes-intuit/assumer-go/assumer     [no test files]
==> Tests complete.
==> Building binaries 0.0.1
----> Building linux binary
----> Building darwin binary
----> Building windows binary

Creating artifact assumer-0.0.1-darwin-amd64.tgz...
Creating artifact assumer-0.0.1-linux-amd64.tgz...
Creating artifact assumer-0.0.1-windows-amd64.tgz...
Done

Complete
```

## Usage ##

To see help text:
  * `assumer -h` To see help text

Standard usage:
  * `assumer -a 123456789012 -r "target/role-name" -A 987654321098 -R "control/role-name" `

To request a console, pass the `-g` flag:
  * `assumer -a 123456789012 -r "target_role/target_path" -A 987654321098 -R "control_role/control_path" -g`

## Notes ##
 1. To be able to use this utility you will need to have permission to assume-role against the role you specify!
