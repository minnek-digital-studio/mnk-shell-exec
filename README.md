# MNKShellExec

MNKShellExec is a Go package developed by `Minnek Digital Studio` that provides functions for executing system shell commands on various operating systems, including Windows and Unix-based systems.

## Features

- Detects and uses the appropriate system shell for the current operating system
- Executes shell commands and returns standard output, error output, and execution errors
- Real-time output streaming for live monitoring of command execution
- Optional output printing for convenience

## Installation
```bash
go get github.com/minnek-digital-studio/mnk-shell-exec
```

## Usage

```go
import "github.com/minnek-digital-studio/mnk-shell-exec"
```

### Execute a command

```go
output := mnkShellExec.ExecuteCommand("ls", true)
```

### Execute a command with real-time output

```go
outStr, errStr, err := mnkShellExec.OutLive("ping google.com")
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

We welcome contributions! Please submit a pull request or create an issue if you find a bug or want to suggest a new feature.

## Contact

Minnek Digital Studio
- Website: [https://minnekdigital.com/](https://minnekdigital.com/)

`mnkShellExec v0.1.0`
