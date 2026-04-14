# zzgo - miniapps for the terminal - inspired by <https://funcoeszz.net>

The [funcoeszz](https://github.com/funcoeszz/funcoeszz) is a cool shell scripting project,
started by @aureliojargas more than two decades ago. It's a collection of more
than a hundred of functions with miniapps with utilities to be used directly
from the terminal.

This project here is an attempt to rewrite some of those functionalities in
Golang, where the functions are invoked as subcommands (I'm going to use
[Cobra](https://cobra.dev) to implement the subcommands.

The original funcoeszz project is written in pt-BR, while here I'm using English.

My motivation:

- practice Golang
- practice TDD
- learn more about Cobra CLI framework.
