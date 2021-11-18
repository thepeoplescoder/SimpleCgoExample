// package main is nothing special.  This is essentially
// a glorified cgo hello world.
package main

/*
#include <stdio.h>
#include "example.h"
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

// getCString allocates memory for a C-style string via C.CString,
// returns a *C.char pointer to that address, along with a function
// to free the memory that was allocated.
func getCString(str string) (*C.char, func()) {
	s := C.CString(str)

	// Return the string and a closure to call the C function
	// responsible for freeing its memory
	return s, func() {
		C.free_memory(unsafe.Pointer(s))
	}
}

// main is the entry point of all Go programs.  In this
// particular example, we call mainWithErrorCode which
// provides a code to return to the operating system
// via os.Exit().
func main() {
	os.Exit(mainWithErrorCode())
}

// mainWithErrorCode is an attempt at providing a way to return
// an error code to the operating system via a C-ish main, where
// the integer value returned is that such code.
func mainWithErrorCode() int {

	// Say some nonsense
	fmt.Println("Look at me, I'm printing text from a Go function,")
	fmt.Println("but not for long! ;)")
	fmt.Println()

	// Essentially use this pattern to allocate memory for a C-style
	// string, ensuring that it gets freed before the function exits.
	// This simple example is a little convoluted, but I feel it
	// demonstrates what you can do with C and Go.
	cStr, freeThis := getCString("Hello world from CGO!!!\n")
	defer freeThis()

	// It's important to know that Go cannot call any variadic functions
	// such as printf, so to call such a function, we would have to wrap
	// the printf call into another C function, and call *that* C function
	// from Go instead of calling the variadic function directly.
	//
	// However, I don't see the point of using printf() from Go if you can't
	// use the formatting, so we're using puts() instead.
	//
	// If we didn't want the automatic newline that puts() appends before
	// printing to stdout, we can call fputs() by using C.fputs(cStr, C.stdout).
	C.puts(cStr)

	// Same thing as before, but using C.CString and
	// deferring the execution of the C function directly.
	cStr = C.CString("Here are the command line arguments:")
	defer C.free_memory(unsafe.Pointer(cStr))
	C.puts(cStr)

	// Show off the command line arguments.
	for index, value := range os.Args {
		cStr, freeThis = getCString(fmt.Sprintf("args[%d] = %v", index, value))
		defer freeThis()
		C.puts(cStr)
	}
	C.putchar('\n')

	// Saying goodbye...
	fmt.Println("Okay, I'm done playing with cgo.")
	fmt.Println("You can leave, but before you go, watch")
	fmt.Println("the magic of defer call the functions written")
	fmt.Println("in C to free the memory for the strings!")
	fmt.Println()

	// A value of zero indicates no error.
	return 0
}
