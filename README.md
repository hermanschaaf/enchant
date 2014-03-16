Go-Enchant
==========

Go-Enchant contains bindings for the C Enchant Spellcheck Library in Go. Instead of offering direct mappings to the Enchant functions, it abstracts away some complexity and makes it easier to do resource management in Go.

### Installation

First off, you will need to have the `libenchant` development files installed, along with any dictionaries you might want to use (aspell, hunspell, etc):

```bash
sudo aptitude install libenchant-dev
```

Then install this package with `go get`:

```bash
go get github.com/hermanschaaf/enchant
```

### Usage
