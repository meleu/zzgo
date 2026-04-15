# zzgo - miniapps for the terminal

An attempt to rewrite, in Go, some of the functionalities provided by
[funcoeszz](https://funcoeszz.net).

The [funcoeszz](https://github.com/funcoeszz/funcoeszz) is a cool project with a collection
of shell scripts with dozens miniapps to be used directly from the terminal.

The idea here is to rewrite some of those functionalities in Go and allow them
to be invoked as subcommands (`zz command` rather than `zzfunction`).
Also, I'm using English here, while the original project is in Brazilian Portuguese.

My motivation:

- practice Golang
- practice TDD, in general, and Go testing, in particular
- learn more about Cobra CLI framework
