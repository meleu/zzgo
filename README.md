# zzgo - miniapps for the terminal

An attempt to rewrite, in Go, some of the functionalities provided by
[funcoeszz](https://funcoeszz.net).

The [funcoeszz](https://github.com/funcoeszz/funcoeszz) is a cool project with a collection
of shell scripts with dozens miniapps to be used directly from the terminal.

The idea here is to rewrite some of those functionalities in Go and allow them
to be invoked as subcommands (`zz command` rather than `zzfunction`).
Also, I'm using English here, while the original project is in Brazilian Portuguese.

My motivation/goal:

- practice Golang
- practice TDD, in general, and Go testing, in particular
- learn more about Cobra CLI framework

While building this project I'm taking [notes about my learnings](./LEARNING_NOTES.md).

## AI Software Engineering Tutor

**The code in this repository is 100% organic, typed with my own fingers.**
However, I'm using Claude Code to act like a Software Engineering tutor and
I'm finding it an enriching experience.

In order to start a session with the tutor, I run `claude --agent tutor`.

You can see in the [tutor agent prompt](/.claude/agents/tutor.md) that it doesn't
have permissions to edit/write code (also enforced via [settings](./.claude/settings.json)).
The prompt also instructs the agent to never provide copy-pasteable blocks, but
describe what to write in natural language. This forces me to think harder
about what code to write.

The tutor is also enforcing the TDD red->green->refactor approach, and after
some sessions I can confirm that writing tests first and implementing testable
production code is constant concern it has in mind.
