package main

import (
	"bytes"
	"strings"
	"sync"
)

var (
	isStr = strings.Contains
	isByt = bytes.Contains

	FATAL_CLASS = []string{
		"JndiLookup.class",
		"JndiManager.class",
	}
	SIGNATURES = []byte{
		0x06, 0x4C, 0x6F, 0x6F, 0x6B, 0x75, 0x70, // JndiLookup.class
		0x6C, 0x6F, 0x67, 0x34, 0x6A, 0x32, 0x2E, // JndiManager.class
	}

	wg  sync.WaitGroup
	opt options
)
