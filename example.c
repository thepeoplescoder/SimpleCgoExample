#include <stdio.h>
#include <stdlib.h>
#include "example.h"

// We don't need to do all of this extra stuff to free
// a block of memory.  I'm just adding all of this to
// demonstrate calling a C function via defer.
void free_memory(void *ptr)
{
	printf("Freeing memory at 0x%X...  ", ptr);
	printf("Value was \"%s\"\n", ptr);
	free(ptr);
}