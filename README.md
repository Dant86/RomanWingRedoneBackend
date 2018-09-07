# Roman wing back-end library

## Build instructions:

__MacOS__: `go build -o libbackend.dylib -buildmode=c-shared main.go`

__Linux__: `go build -o libbackend.so -buildmode=c-shared main.go`

__Windows__: `go build -o libbackend.dll -buildmode=c-shared main.go`

## Library Function Documentation

Each library function returns a JSON string that still has to be parsed
wherever you import it.

`CreateUser(fName, lName, email, pword string) string`

Parameters
----------
fName: string
    User's first name
lName: string
    User's last name
email:
    User's email
pword: string
    User's password
