# quoteguard

quoteguard is a static analysis tool for Go source code that helps developers choose the optimal quoting style for string literals. By analyzing the abstract syntax tree (AST), it detects string literals that could be more appropriately enclosed with either double quotes (") or back quotes (`). For example, if a string literal does not require escape sequences, quoteguard suggests using back quotes for better readability. Conversely, if a raw string can be safely represented with double quotes, it recommends switching to double quotes. This tool helps maintain clean and idiomatic Go code by guiding developers toward the most suitable quoting style for each string literal.

## Installation

```bash
go install github.com/otakakot/quoteguard/cmd/quoteguard@latest
```

## Usage

```bash
go vet -vettool=$(which quoteguard) ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
