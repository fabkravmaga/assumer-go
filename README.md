assumer-go
==================
## Overview ##
The Assumer gem is an interface allowing the user to assume role into an account.  Credentials can be loaded anywhere the AWS STS Client knows to load credentials from (ENV, profile, etc.)

## Content ##
Assumer will assume-role on a target account and generate temporary API keys for that account.

## Installation ##

Cross-compiled binaries work on Linux, MAC OSX, and Windows.  Download location to be TBD.  

### Build from Source ###

1. Clone the repository
1. Change directory into the newly-cloned repository
  * `cd assumer`
1. Run the build script (assumes that you have go 1.6 installed, will cross compile for linux, MAC OSX, and Windows)
  * `./build`
1. Alternately, you can simply run MAKE 
  * `make`

## Usage ##

To see help text:
  * `assumer -h` To see help text

Standard usage:
  * `assumer -a 123456789012 -r "target/role-name" -A 987654321098 -R "control/role-name" `

To request a console, pass the `-g` flag:
  * `assumer -a 123456789012 -r "target_role/target_path" -A 987654321098 -R "control_role/control_path" -g`

## Notes ##
 1. To be able to use this utility you will need to have permission to assume-role against the role you specify!
