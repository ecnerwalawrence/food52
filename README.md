FourthLargest

==== SETUP ====

1. Install go

- https://golang.org/doc/install

Summary of instructions for macos
- Download https://golang.org/dl/
- Install "Apple macOS"
- Check if `/usr/local/go/bin` is in your $PATH by running `go`
	- If not, add `/usr/local/go/bin` to your $PATH

2. Retrieve code

git clone git@github.com:ecnerwalawrence/food52.git

==== EXECUTING THE CODE ====


OPTION 1:

1. To execute unit tests

go test -v ./...

2. To run (by default it uses sample_input.json).

go run -v fourthlargest.go -json sample_input2.json 

OR you can pass in your own test cases with your json files.

go run -v fourthlargest.go -json <your json file>

If you pass a non-json file, the code will error and exit.

OPTION 2:
Alternative if you trust the binary.

