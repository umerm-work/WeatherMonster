# Weather Monster

## Setting Up the project

First need to settup postgres database and user credentials. Then update the config.yaml file with host,port,username and password.

To download the dependencies run the following commmand

`go get ./...`

## Running the tests
To run test use this command

`go test .\pkg\http`

or
Run `test.sh`

## Running the app
To run the service execute the following command

`go run cmd/main.go`


NOTE : Didn't implemented webhook functionality properly. Callback URL will not work on localhost.