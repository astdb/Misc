
// Markov Chain Algorithm
// ----------------------
// set w1 and w2 to the first two words in the text
// print w1 and w2
// loop:
// 	randomly choose w3, one of the successors of prefix w1 w2 in the text
// 	print w3
// 	replace w1 and w2 by w2 and w3
// 	repeat loop

package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	// map of prefixes to suffixes
	statetab := map[string][]string{}

	// ensure input filename entered on command line
	if len(os.Args) <= 1 {
		fmt.Println("Usage: > go run markov.go <input.file>")
		return
	}

	// capture input file name
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return
	}
	defer file.Close()

	// setup scanner to read input file
	scanner := bufio.NewScanner(file)

	k := 0
	pref1, pref2, suf := "", "", ""
	pref1_set, pref2_set, suf_set, record_state, sufx, quit := false, false, false, false, false, false

	for scanner.Scan() {
		// read next line
		line := scanner.Text()
		// fmt.Println(line)
		// fmt.Println("----------------")
		words := strings.Split(line, " ")
		words_len := len(words)

		for i := 0; i < words_len; i++ {
			fmt.Printf("------------------\nPREF1 = <%s>, PREF2 = <%s>, SUF: <%s>\n", pref1, pref2, suf)
			if !pref1_set {
				w1 := strings.TrimSpace(words[i])

				if w1 == "" {
					continue
				}

				pref1 = w1
				pref1_set = true
				fmt.Printf("Setting <%s> to pref1 (init)\n", pref1)
			}

			ipp := 0
			if !pref2_set {
				ipp = i + 1
				if ipp < words_len {
					w2 := strings.TrimSpace(words[ipp])

					if w2 == "" {
						continue
					}

					pref2 = w2
					pref2_set = true
					fmt.Printf("Setting <%s> to pref2 (init)\n", pref2)
				} else {
					continue
				}
			}

			if !suf_set {
				ipp = ipp + 1
				if ipp < words_len {
					w3 := strings.TrimSpace(words[ipp])

					if w3 == "" {
						continue
					}

					suf = w3
					sufx = true
					suf_set = true
					record_state = true
					fmt.Printf("Setting <%s> to suf (init)\n", suf)
				} else {
					continue
				}
			}

			if pref1_set && pref2_set && !sufx {
				// wx := strings.TrimSpace(words[i])
				wx := ""
				if i + 2 < words_len {
					wx = strings.TrimSpace(words[i+2])
				} else {
					continue
				}

				if wx == "" {
					continue
				}

				suf = wx
				record_state = true
				fmt.Printf("Setting <%s> to suf\n", suf)
			}

			// if (pref1_set && pref2_set && suf_set && !quit) {
			if (record_state && !quit) {
				prefkey := pref1 + " " + pref2
				// fmt.Println(prefkey)
				_, found := statetab[prefkey]
				if found {
					fmt.Printf("Adding value <%s> to key <%s>\n", suf, prefkey)
					statetab[prefkey] = append(statetab[prefkey], suf)
				} else {
					fmt.Printf("Adding value <%s> to new key <%s>\n", suf, prefkey)
					statetab[prefkey] = []string{suf}
				}

				sufx = false
				pref1 = pref2
				pref2 = suf
				suf = ""

				// suf_set = false
				record_state = false
				// i += 1
			}
		}

		k++
	}

	// for k, v := range statetab {
	// 	fmt.Printf("<%s> -> %v\n", k, v)
	// }
}

// --
