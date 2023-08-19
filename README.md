# pdf-tools
PDF Tools is a command-line utility written in Go that enables you to manipulate PDF files. This tool allows you to split a given PDF file into individual pages and save them to a specified output directory.

## Features

- Split PDF files into separate pages and store them in a designated output directory.
- Customizable output directory for saving the split PDF files.
- User-friendly command-line interface powered by the Cobra library.


## Usage
- Installation
```
go install github.com/linzhengen/pdf-tools@latest
```
- Execution
```
# help
pdf-tools help
# split pdf
pdf-tools split --in-file=hoge.pdf
```
