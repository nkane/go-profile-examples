package main

import (
	"io"
	"log"
	"os"
	"unicode"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("couldn't open file %q: %v", os.Args[1], err)
	}
	defer f.Close()

	words := 0
	inword := false

	// NOTE(1): first change to increase performance
	// make sure to review each different profile type each time
	//b := bufio.NewReader(f)

	for {
		// NOTE(1): first change to increase performance
		//r, err := readbyte(b)

		r, err := readbyte(f)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("could not read file %q: %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inword {
			words++
			inword = false
		}
		inword = unicode.IsLetter(r)
	}
	log.Printf("%q: %d words\n", os.Args[1], words)
	// NOTE(nick): simulate a program that runs forever
	//select {}
}

// NOTE(2): second memory optimization, make it global instead of local
// var buf [1]byte
func readbyte(r io.Reader) (rune, error) {
	var buf [1]byte
	_, err := r.Read(buf[:])
	return rune(buf[0]), err
}
