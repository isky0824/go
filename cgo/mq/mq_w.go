package main 

/*
	#include <stdio.h>

	typedef void (*HelloCB)(char *);

	extern void Wrapper( HelloCB cb, char*msg ) {
		cb( msg );
	}
*/
import "C"