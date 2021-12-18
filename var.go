package main

import (
	"strings"
	"sync"
)

var (
	isMatch     = strings.Contains
	FATAL_CLASS = []string{
		"JndiLookup.class",
		"JndiManager.class",
	}

	wg  sync.WaitGroup
	opt options
)
