package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"flag"
)

const (
	f1K  = "./data/1000.prime"
	f10K = "./data/10000.prime"
	f1M  = "./data/1M.prime"
)

var (
	path = flag.String("path", "./data/10000.prime", "open path prime store")
)

type PrimeStore struct {
	path  string
	store []string
	index Uint64Slice
	count uint64
	base  uint64
	first uint64
	last  uint64
	low   uint64
	high  uint64
}

func NewPrimeStore(file string) (ps *PrimeStore, err error) {
	ps = new(PrimeStore)
	ps.path = file
	err = ps.loadFromFile()
	if err != nil {
		return nil, err
	}

	ps.atoi()

	if !ps.validate() {
		return nil, errors.New("Invalidate: unsorted")
	}

	ps.first = ps.index[0]
	ps.last = ps.index[ps.count-1]
	ps.low = 1
	ps.high = ps.last

	return
}

func (ps *PrimeStore) IsPrime(n uint64) (bool, error) {
	if n < ps.low || n > ps.high {
		return false, errors.New("out of range")
	}

	i, j := uint64(0), ps.count
	for i < j {
		h := i + (j-i)/2
		if n == ps.index[h] {
			return true, nil
		} else if n > ps.index[h] {
			i = h + 1
		} else {
			j = h
		}
	}
	return false, nil
}

/* GetByIndex return nth primer, eg:
2, 3, 5, 7, 9...
GetByIndex(3) return 5
OutOfRange in Store as return 0
*/
func (ps *PrimeStore) GetByIndex(nth uint64) (n uint64) {
	n = 0
	if nth < ps.base || nth >= (ps.base+ps.count) {
		log.Print("out of range.", nth, " ", ps)
		return
	}

	n, _ = strconv.ParseUint(ps.store[nth-ps.base], 0, 64)
	return
}

//No zero
func isPan(n string, nozero bool) bool {
	bitmap := []int{0,0,0,0,0,0,0,0,0,0}
	for _, v := range n {
		bitmap[v-'0']++
		if bitmap[v-'0'] > 1 {
			return false
		}
		if nozero && bitmap[0] > 0 {
			return false
		}
	}
	for i:=1; i <=8; i++ {
		if bitmap[i]==0 && bitmap[i+1]==1 {
			return false
		}
	}
	return true
}

func (ps *PrimeStore) GetPandigitalLargest() (n uint64) {
	for i := ps.count-1; i >= ps.base; i-- {
		if isPan(ps.store[i], true) {
			return ps.index[i]
		}
	}
	return 0
}

func (ps *PrimeStore) String() string {
	return fmt.Sprint("bsae,count,first,last,low,high=", ps.base, ps.count, ps.first, ps.last, ps.low, ps.high)
}

func (ps *PrimeStore) loadFromFile() (err error) {
	file, err := ioutil.ReadFile(ps.path)
	if err != nil {
		log.Printf("Open file.%s error.%s", ps.path, err)
		return err
	}

	reg := regexp.MustCompile("[0-9]+")
	ps.store = reg.FindAllString(string(file), len(string(file)))

	ps.count = uint64(len(ps.store))
	ps.base = 1 //TODO:base my update by file content in fact
	return nil
}

func (ps *PrimeStore) dump() {
	for i, v := range ps.index {
		fmt.Println(i, v)
	}
	fmt.Print(ps)
}

func (ps *PrimeStore) atoi() {
	ps.index = make([]uint64, ps.count, ps.count)
	for i, v := range ps.store {
		ps.index[i], _ = strconv.ParseUint(v, 0, 64)
	}
}

func (ps *PrimeStore) validate() bool {
	return sort.IsSorted(ps.index)
}

type Uint64Slice []uint64

func (index Uint64Slice) Len() int {
	return len(index)
}

func (index Uint64Slice) Less(i, j int) bool {
	return index[i] <= index[j]
}

func (index Uint64Slice) Swap(i, j int) {
	index[i], index[j] = index[j], index[i]
}

func main() {
	flag.Parse()

	ps, err := NewPrimeStore(*path)
	if err != nil {
		return
	}
	/*
	fmt.Println(ps)

	fmt.Println(ps.GetByIndex(0))
	fmt.Println(ps.GetByIndex(10))
	fmt.Println(ps.GetByIndex(1000))
	fmt.Println(ps.GetByIndex(10000))
	fmt.Println(ps.IsPrime(1))
	fmt.Println(ps.IsPrime(2))
	fmt.Println(ps.IsPrime(3))
	fmt.Println(ps.IsPrime(4))
	fmt.Println(ps.IsPrime(100))
	fmt.Println(ps.IsPrime(101))
	*/

	fmt.Println(ps.GetPandigitalLargest())
}
