package main

/*
#cgo LDFLAGS: -lenchant
#include <stdlib.h>
#include "enchant/enchant.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

// Enchant is a type that encapsulates Enchant internals
type Enchant struct {
	broker *C.EnchantBroker
	dict   *C.EnchantDict
}

// NewEnchant creates a new Enchant instance for access
// to the rest of the Enchant API.
//
// The returned value is an Enchant struct.
//
// Example usage:
//
// 		enchant, err := enchant.NewEnchant()
// 		if err != nil {
// 			panic("Enchant error: " + err.Error())
// 		}
// 		defer enchant.Free()
//      fmt.Println(fmt.DictExists("zh"))
//
// Because the Enchant package is a binding to Enchant C library, memory
// allocated by the NewEnchant() call has to be disposed explicitly.
// This is why the above example contains a deferred call to Free().
func NewEnchant() (e *Enchant, err error) {
	broker := C.enchant_broker_init()
	e = &Enchant{broker, nil}
	return e, nil
}

// Free frees the Enchant broker and dictionary, and needs
// to be called when use of the library is no longer needed
// to prevent memory leaks.
func (e *Enchant) Free() {
	C.enchant_broker_free(e.broker)
}

// DictExists wraps enchant_broker_dict_exists.
// It takes a language code name, such as "en_GB", as string
// argument, and it returns whether or not such a dictionary
// is installed on the system.
func (e *Enchant) DictExists(name string) bool {
	// create CString representation
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	exists := C.enchant_broker_dict_exists(e.broker, cName)
	return exists > 0
}

func main() {
	e, err := NewEnchant()
	defer e.Free()

	// check that dictionary exists
	fmt.Println(e.DictExists("en_US"), e.DictExists("nl"), e.DictExists("af"), e.DictExists("bla"))
}
