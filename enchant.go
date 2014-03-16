package enchant

/*
#cgo LDFLAGS: -lenchant
#include <stdlib.h>
#include <sys/types.h>
#include "enchant/enchant.h"
*/
import "C"

import (
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
	if e.dict != nil {
		C.enchant_broker_free_dict(e.broker, e.dict)
	}
	C.enchant_broker_free(e.broker)
}

// DictExists wraps enchant_broker_dict_exists.
// It takes a language code name, such as "en_GB", as string
// argument, and it returns whether or not such a dictionary
// is installed on the system.
func (e *Enchant) DictExists(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	exists := C.enchant_broker_dict_exists(e.broker, cName)
	return exists > 0
}

// LoadDict wraps enchant_broker_request_dict, and adds
// the loaded dictionary to the Enchant instance.
// It takes a language code name, such as "en_GB", as string
// argument, and it returns a EnchantDict representation
// of this dictionary.
func (e *Enchant) LoadDict(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	if e.dict != nil {
		C.enchant_broker_free_dict(e.broker, e.dict)
	}
	dict := C.enchant_broker_request_dict(e.broker, cName)
	e.dict = dict
}

// Check whether a given word is in the currently loaded dictionary.
// This wraps enchant_dict_check.
// It returns a boolean value: true if the word is in the dictionary,
// false otherwise.
func (e *Enchant) Check(word string) bool {
	cWord := C.CString(word)
	defer C.free(unsafe.Pointer(cWord))

	size := uintptr(len(word))
	s := (*C.ssize_t)(unsafe.Pointer(&size))

	return C.enchant_dict_check(e.dict, cWord, *s) == 0
}
