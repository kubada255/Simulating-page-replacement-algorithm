package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type czestotliwosc struct {
	nr            int
	czestotliwosc int
}

func zawiera(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func zawiera_nfu(s []czestotliwosc, e int) bool {
	for _, a := range s {
		if a.nr == e {
			return true
		}
	}
	return false
}
func index(s []int, e int) int {
	for i, a := range s {
		if a == e {

			return i
		}
	}
	return math.MaxInt32
}

// FIFO algorithm
func FIFO(numer_strony []int, ilosc_ramek int) {
	fmt.Println("FIFO")

	pominiete := 0
	stos := []int{}
	for _, numer := range numer_strony {
		if len(stos) < ilosc_ramek {
			if zawiera(stos, numer) {
				pominiete++
			} else {
				stos = append(stos, numer)
			}
		} else {
			if !zawiera(stos, numer) {
				stos = stos[1:]
				stos = append(stos, numer)
			} else {
				pominiete++
			}
		}

	}
	fmt.Println("Zamienione: ", len(numer_strony)-pominiete)
	fmt.Println("Pominiete: ", pominiete)
}

// LRU algorithm
func LRU(numer_strony []int, ilosc_ramek int) {
	fmt.Println("LRU")
	pominiete := 0
	stos := []int{}
	for _, numer := range numer_strony {
		if len(stos) < ilosc_ramek {
			if zawiera(stos, numer) {
				pominiete++
			} else {
				stos = append(stos, numer)
			}
		} else {
			if !zawiera(stos, numer) {
				stos = stos[1:]
				stos = append(stos, numer)

			} else {

				if stos[0] == numer {
					stos = stos[1:]
					stos = append(stos, numer)
				} else {
					for i := 1; i < len(stos); i++ {
						if stos[i] == numer {
							stos = append(stos[:i], stos[i+1:]...)
							stos = append(stos, numer)
							break
						}
					}
				}
				pominiete++
			}

		}

	}

	fmt.Println("Zamienione: ", len(numer_strony)-pominiete)
	fmt.Println("Pominiete: ", pominiete)

}

// OPT algorithm
func OPT(numer_strony []int, ilosc_ramek int) {
	fmt.Println("OPT")
	pominiete := 0
	stos := []int{}
	nastepne_uzycie := make([]int, ilosc_ramek)
	dlugosc := len(numer_strony)
	max := 0
	najpozniej := 0
	for i := range nastepne_uzycie {
		nastepne_uzycie[i] = math.MaxInt32
	}
	for _, numer := range numer_strony {
		max = 0
		if len(stos) < ilosc_ramek {
			if zawiera(stos, numer) {
				pominiete++
			} else {
				stos = append(stos, numer)
			}
		} else {
			if zawiera(stos, numer) {
				pominiete++
			} else {
				for i, str := range stos {
					if index(numer_strony, str) > max {
						max = index(numer_strony, str)
						najpozniej = i
					}
				}
				stos[najpozniej] = numer

			}

		}
		numer_strony = numer_strony[1:]

	}
	fmt.Println("Zamienione: ", dlugosc-pominiete)
	fmt.Println("Pominiete: ", pominiete)
}

// NFU algorithm
func NFU(numer_strony []int, ilosc_ramek int) {
	fmt.Println("NFU")
	pominiete := 0
	pamiec := []czestotliwosc{}

	stos := []czestotliwosc{}
	min := 0
	for _, numer := range numer_strony {
		min = 0
		if len(stos) < ilosc_ramek || zawiera_nfu(stos, numer) {
			if !zawiera_nfu(stos, numer) {
				stos = append(stos, czestotliwosc{numer, 1})
			} else {
				for i := 0; i < len(stos); i++ {
					if stos[i].nr == numer {
						stos[i].czestotliwosc++
					}
				}
				pominiete++
			}
		} else {
			for i := 0; i < len(stos); i++ {
				if stos[i].czestotliwosc < stos[min].czestotliwosc {
					min = i
				}
			}
			if zawiera_nfu(pamiec, stos[min].nr) {
				for i := 0; i < len(pamiec); i++ {
					if pamiec[i].nr == stos[min].nr {
						pamiec[i].czestotliwosc = stos[min].czestotliwosc
					}
				}
			} else {
				pamiec = append(pamiec, stos[min])
			}
			if zawiera_nfu(pamiec, numer) {
				for i := 0; i < len(pamiec); i++ {
					if pamiec[i].nr == numer {
						stos[min].nr = numer
						stos[min].czestotliwosc = pamiec[i].czestotliwosc + 1
						pamiec = append(pamiec[:i], pamiec[i+1:]...)
					}
				}
			} else {
				stos[min].nr = numer
				stos[min].czestotliwosc = 1
			}
		}
	}
	fmt.Println("Zamienione: ", len(numer_strony)-pominiete)
	fmt.Println("Pominiete: ", pominiete)

}

func main() {
	var plik string
	fmt.Print("Podaj nazwę pliku: ")
	fmt.Scanln(&plik)
	numer_strony := []int{}
	ilosc_ramek := 0

	//Read file
	data, err := ioutil.ReadFile(plik)
	if err != nil {

		fmt.Println("Blad czytania pliku", err)
		fmt.Scanf(" ")
		return
	}

	lines := strings.Split(string(data), "\r\n")

	//Split lines
	for i, linia := range lines {
		words := strings.Split(linia, " ")
		for _, word := range words {

			liczba, err := strconv.Atoi(word)
			if err != nil {
				fmt.Println("Blad czytania pliku", err)
				fmt.Scanf(" ")
				return
			}
			if i == 0 {
				numer_strony = append(numer_strony, liczba)
			}
			if i == 1 {
				ilosc_ramek = liczba
			}

		}

	}

	FIFO(numer_strony, ilosc_ramek)
	LRU(numer_strony, ilosc_ramek)
	OPT(numer_strony, ilosc_ramek)
	NFU(numer_strony, ilosc_ramek)
	fmt.Scanf(" ")
}
