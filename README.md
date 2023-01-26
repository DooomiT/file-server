# file-server

This is a basic file server written in go.

## Setup

```bash
git clone https://github.com/DooomiT/file-server.git
go build
./file-server
```

## Usage

```plain
This file server serves files from a given directory

Usage:
  filer-server [command]

general
  serve       

user

Additional Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command

Flags:
  -h, --help   help for filer-server

Use "filer-server [command] --help" for more information about a command.
```

### Example

```bash
file-server serve testdir
```



