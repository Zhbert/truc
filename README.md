# True Relative URL Creator

<p align="center">
  <img src="./logo.png" style="max-height:100%;" height="175">
</p>

A simple CLI utility for preparing the correct relative links.

## Installing

Download the appropriate version from the releases page and put it in the $PATH.

## Usage

Accepts two parameters as input:

- `--source` or `-s` — an original link;
- `--target` or `-t` — a target link.

For example:

```console
truc -s https://github.com/Zhbert/impomoro/tree/main/internal -t https://github.com/Zhbert/colligendis`
```

> You can use the `-c` (`--copy`) key to automatically copy the result to the clipboard.

Result:

```text
    Source URL: https://github.com/Zhbert/impomoro/tree/main/internal
    Target URL: https://github.com/Zhbert/colligendis
0: github.com
1: Zhbert
2: impomoro -> colligendis
Levels back: 3
    Result URL: ../../../colligendis
```

The resulting link can be used in, for example, static generators like Jekyll or Hugo.

## Reference help

Help on commands and flags can be viewed directly on the command line using the `help` command or the `-h` key.

```text
truc is a simple CLI tool for creating true relative URLs.

Allows you to quickly calculate a relative link based on the source and destination.

Usage:
  truc [flags]
  truc [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Show the utility version

Flags:
  -c, --copy            Copy to clipboard
  -h, --help            help for truc
  -s, --source string   Specify the source URL
  -t, --target string   Specify the target URL
  -v, --verbose         Enable full log

Use "truc [command] --help" for more information about a command.
```
