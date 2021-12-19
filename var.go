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
		"NetUtils.class",
	}
	SIGNATURES = []byte{
		0x06, 0x4C, 0x6F, 0x6F, 0x6B, 0x75, 0x70, // positive for: JndiLookup.class (VERIFY)
		0x6C, 0x6F, 0x67, 0x34, 0x6A, 0x32, 0x2E, // negative for: JndiManager.class (OLD CHECK)
		0x2F, 0x4C, 0x69, 0x73, 0x74, 0x3B, 0x01, // negative for: NetUtils.class (OLD CHECK)
	}

	wg  sync.WaitGroup
	opt options
)
