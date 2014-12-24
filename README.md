## Lessons for learning "Go" programing language

The lessons are broken into serveral sections. Each section offers details about the specific topic under discussion.

## Attributions

Some of the details have been shamelessly copied from:
- [Effective Go] (http://golang.org/doc/effective_go.html)
- [The Go Programming Language Specification] (http://golang.org/ref/spec)
- [The Go Blog] (http://blog.golang.org)

## Getting the Material

	git clone https://github.com/spuranam/go-learn

or

	go get github.com/spuranam/go-learn

or

[download the zip archive] (https://github.com/spuranam/go-learn/archive/master.zip)

## Running the code examples

**_Installing Go_**

[Vist this page to install Go development tools] (http://golang.org/doc/install)

**_Build & Run sample_**

Complie go source file and generate the binary

	go build name_of_program

Complie go source file, execute it send the output if any to stdout and throw away the generated binary

	go run name_of_program

 Compile go source file, and install it in bin subdirectory of your GOPATH

	go install name_of_program

**Running on go playground**

To run these samples on [Go playground, paste the code here] (http://play.golang.org)
Please note that go playground is sandboxed environment, with following limitations:
 - Time is constant, i.e. it does not change
 - No filesystem i.e. you do not have access to the filesystem
 - No network .i.e. you can not run an code that imports net package

However the playground allows you get started with the business of learning the language right away, with getting bogged down in setting up the development/learning environment.
