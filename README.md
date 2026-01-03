# genignore
## Utility for setting up .gitignore files

### Usage

Usage: `genignore [template]` \
For a full list of templates, see `list.go`.

### Development

To setup the project for development, simply run `git clone https://github.com/regarager/genignore`.

To build, run `go build`. To build manpages, you can do `make man`, which builds the manpages.

### Installation

To install the file, run `go install [directory]`. If `directory` is left blank, it will be installed into `$GOPATH/bin` (probably).

Alternatively, you can move your binary to your desired directory in your PATH.

The manpage needs to go in the corresponding prefix's `share/man/man1` if you built it.
