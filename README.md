# naming

[![asciicast](https://asciinema.org/a/0iYbY2jqHqqV6vxvcqz1tcMs5.svg)](https://asciinema.org/a/0iYbY2jqHqqV6vxvcqz1tcMs5)

## Introduction

`naming` is a command line tool designed to improve the readability of your code by suggesting intuitive and descriptive names for your functions and variables. Using the ChatGPT API, `naming` generates program naming suggestions that are tailored to your specific code, ensuring that your variable and function names are both accurate and easily understandable.

While the original design of Naming included an `overwrite` flag that allowed users to automatically apply the suggested names to their code, this feature was ultimately removed. This decision was made to encourage developers to evaluate the naming suggestions provided by Naming and manually make changes as necessary. In some cases, short variable names may be more appropriate, especially for variables with a short lifespan.

By removing the `overwrite` feature, `naming` empowers users to consider the suggested names thoughtfully and make the best decisions for their code. Whether you are a seasoned developer or a newcomer to programming, Naming is a valuable tool that can help you write more readable and maintainable code.

## Installation

You can install `naming` using `go install`

```sh
go install github.com/davidleitw/naming@latest
```

set `CHATGPT_API_KEY`

```sh
export CHATGPT_API_KEY=<your-api-key>
```

## Usage
To generate naming suggestions for a specific file, use the following command:

```sh
naming -f <file_path>
```

Replace <file_path> with the path to the file for which you want to generate naming suggestions. naming will analyze the file and provide a list of recommended variable and function names that you can use to improve the readability and maintainability of your code.

If you want to see more information about the naming tool, including a list of available flags and options, you can use the following command:

```sh
naming --help
```

This will display a detailed help message that describes all of the available options for the naming tool.

At present, only the simple command -f is available. If you have any useful features you would like naming to support, please contribute a pull request and let's work together to improve the readability of your code!

## License

This project is licensed under the terms of the GNU GENERAL PUBLIC LICENSE. Please see the LICENSE file for full details of the license.