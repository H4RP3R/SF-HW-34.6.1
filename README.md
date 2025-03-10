# Mathematical expression parser

## Usage

| Flag    | Default    | Description                                      |
|---------|------------|--------------------------------------------------|
| -h      | —          | Show help                                        |
| -in     | *required* | Input file with mathematical expressions         |
| -out    | `out.txt`  | Output file for results                          |

```console
# Default output
go run cmd/main.go -in test_data/in.txt

# Specify output
go run cmd/main.go -in test_data/in.txt -out results.txt
```

## Features

- Skips lines with incorrect format
- Writes logs to the console

## Expression Format

Input lines must match pattern:  
`[number][operator][number]=?`  
Example valid expressions:  

- `-3.14+2.718=?`  
- `100/25=?`
