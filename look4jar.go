package main

import (
	"sync"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
)

func main() {
	path := opt.Path

	gologger.Debug().Msgf("Digging into '%s'...", aurora.Magenta(path))
	files, err := lookup(path)
	if err != nil {
		gologger.Error().Msgf(ERR_LOOK, aurora.Red(path), aurora.Red(err.Error()))
	}

	for _, file := range files {
		gologger.Verbose().Msgf(FOUND_JAR, aurora.Green(file))

		wg.Add(1)
		go check(&wg, file)
	}

	wg.Wait()
	gologger.Info().Msg("Done!")
}

func check(wg *sync.WaitGroup, file string) {
	defer wg.Done()

	read, err := readJar(file)
	if err != nil {
		gologger.Error().Msgf(ERR_OPEN, aurora.Red(file))
	}

	if isVuln(read) {
		gologger.Warning().Msgf(POSSIBLE_VULN, aurora.Yellow(file))
	}
}
