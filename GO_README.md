


## Features

- Concurrent scanning for faster results
- No external dependencies, like `dig` or `awk`
- Easy to use CLI interface

## Installation

1. Install [Go](https://golang.org/doc/install) if you haven't already.

2. Clone the repository:

```bash
git clone https://github.com/samhaxr/takeover.git
```

3. Change to the repository directory and build the binary:
```bash
cd takeover
go build takeover.go
```
4. (Optional) Add the takeover binary to your system's PATH to use it from any location.

## Usage
```bash
./takeover -file <input_file>
```
To enable performance info use the --perf flag
```bash
./takeover -file example_domains.txt --perf
```

## Output
TakeOver will display the scanned domain names along with their respective CNAME records in a formatted output. Each line will show the line number, the domain name, and the corresponding CNAME record.