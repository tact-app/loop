> # 🌀 Loop
>
> Loom structure fetcher.

[![Build][build.icon]][build.page]
[![Documentation][docs.icon]][docs.page]
[![Quality][quality.icon]][quality.page]
[![Template][template.icon]][template.page]
[![Coverage][coverage.icon]][coverage.page]
[![Mirror][mirror.icon]][mirror.page]

## 💡 Idea

```bash
$ loop tree
OctoLab
└── Tact
    ├── directory
    │   ├── graphql
    │   │   ├── https://www.loom.com/share/abcdefghijklmopqrstuvwxyz1234567
    │   │   ├── https://www.loom.com/share/abcdefghijklmopqrstuvwxyz1234568
    │   │   └── https://www.loom.com/share/abcdefghijklmopqrstuvwxyz1234569
    │   └── responses
    │       ├── https://www.loom.com/share/abcdefghijklmopqrstuvwxyz1234560
    │       └── https://www.loom.com/share/abcdefghijklmopqrstuvwxyz1234561
    ├── https://www.loom.com/share/abcdefghijklmopqrstuvwxyz1234562
    └── https://www.loom.com/share/abcdefghijklmopqrstuvwxyz1234563

5 spaces, 24 directories, 102 looms
```

## 🏆 Motivation

...

## 🤼‍♂️ How to

...

## 🧩 Installation

### Homebrew

```bash
$ brew install tact-app/tap/loop
```

### Binary

```bash
$ curl -fsSL https://install.octolab.org/tact-app/loop | sh
# or
$ wget -qO-  https://install.octolab.org/tact-app/loop | sh
```

> Don't forget about [security](https://www.idontplaydarts.com/2016/04/detecting-curl-pipe-bash-server-side/).

### Source

```bash
# use standard go tools
$ go get go.octolab.org/tact/loop@latest
# or use egg tool
$ egg tools add go.octolab.org/tact/loop@latest
```

> [egg][] is the `extended go get`.

### Shell completions

```bash
$ loop completion > /path/to/completions/...
# or
$ source <(loop completion)
```

## 🤲 Outcomes

...

<p align="right">made with ❤️ for everyone</p>

[awesome.icon]:     https://awesome.re/mentioned-badge.svg
[build.page]:       https://github.com/tact-app/loop/actions/workflows/ci.yml
[build.icon]:       https://github.com/tact-app/loop/actions/workflows/ci.yml/badge.svg
[coverage.page]:    https://codeclimate.com/github/tact-app/loop/test_coverage
[coverage.icon]:    https://api.codeclimate.com/v1/badges/8491ba0aada439d2df0c/test_coverage
[design.page]:      https://www.notion.so/33715348cc114ea79dd350a25d16e0b0
[docs.page]:        https://pkg.go.dev/go.octolab.org/tact/loop
[docs.icon]:        https://img.shields.io/badge/docs-pkg.go.dev-blue
[mirror.page]:      https://bitbucket.org/kamilsk/go-tool
[mirror.icon]:      https://img.shields.io/badge/mirror-bitbucket-blue
[promo.page]:       https://github.com/tact-app/loop
[quality.page]:     https://goreportcard.com/report/go.octolab.org/tact/loop
[quality.icon]:     https://goreportcard.com/badge/go.octolab.org/tact/loop
[template.page]:    https://github.com/octomation/go-tool
[template.icon]:    https://img.shields.io/badge/template-go--tool-blue

[egg]:              https://github.com/kamilsk/egg
