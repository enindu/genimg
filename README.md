# genimg

genimg is a lightweight tool for generating random images at custom sizes. I often need placeholder images for specific dimensions, and while there are many options available, I wanted a simple, self-hosted solution—so I built this.

## Install

You can install genimg using `go install` command.

```
go install github.com/enindu/genimg@latest
```

## Usage

You can run genimg using the following syntax.

```
genimg <command>:<subcommand> [arguments]
genimg [flags]
```

To display the version message:

```
genimg -v # or "genimg --version"
```

To display the help message:

```
genimg -h # or "genimg --help"
```

## License

This software is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html). You can view the full license [here](https://github.com/enindu/genimg/blob/master/COPYING.md).
