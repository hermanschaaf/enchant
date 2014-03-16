Go-Enchant
==========

Go-Enchant provides bindings for the C [Enchant Spellcheck Library](http://www.abisource.com/projects/enchant/) in Go. Instead of offering direct mappings to the Enchant functions, it abstracts away some complexity and makes it easier to do resource management in Go.

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

Basic usage is illustrated by the following example program:

```go
package main

import (
	"fmt"
	"github.com/hermanschaaf/enchant"
)

func main() {
	// create a new enchant instance
	enchant, err := enchant.NewEnchant()
	if err != nil {
		panic("Enchant error: " + err.Error())
	}

	// defer freeing memory to the end of this program
	defer enchant.Free()

	// check whether a certain dictionary exists on the system
	has_en := enchant.DictExists("en_GB")

	// load the english dictionary:
	if has_en {
		enchant.LoadDict("en_GB")

		// see if a word is in the dictionary:
		fmt.Println("hallo:", enchant.Check("hallo"))

		// and one that won't be in there:
		fmt.Println("wollo:", enchant.Check("wollo"))

		// now let's get some suggestions for "wollo":
		fmt.Println(enchant.Suggest("wollo"))
	}
}
```

1.  First, we create a new Enchant instance using the `NewEnchant` function.

2.  We defer a call to `enchant.Free()` to free memory allocation when our program ends. `Free()` handles the freeing of both the Enchant broker and loaded dictionaries.

3.  Next, we check whether a certain dictionary is installed on the system using a call to `enc.DictExists()`.

4.  We know the dictionary exists now, so we load it into our Enchant instance with a call to `LoadDict()`.

5.  Now we are free to make any calls to Enchant that we want. We call `Check`, which returns whether the given word is contained in the dictionary or not. We expect `"hallo"` to be in the dictionary, and `"wollo"` not to be. Indeed, our program output confirms this:

```
hallo: true
wollo: false
[Rollo woolly hollow follow woollen would worldly]
```

The final line is the result of our call to `enc.Suggest("wollo")`: it returns a slice of spelling suggestions for the given word.

### Documentation

Full documentation can be found at [godoc.org/github.com/hermanschaaf/enchant](http://godoc.org/github.com/hermanschaaf/enchant)