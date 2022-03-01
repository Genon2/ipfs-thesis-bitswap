package ratio

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"testing"
	"time"
)

// Element to Get 
// DHT & BITSWAP are initialized a true 
// when we pass a certain part of the code 
// Cannot be initialized a true both
type Get struct {
	CID        string
	DHT        bool
	BITSWAP    bool
}

// Object that contains all the result
type Ratio struct{
	list []Get
}

// Add a new Get to the Ratio
func (rt *Ratio) Add(str String){
	g = NewGet(str)
	rt.list = append(rt.list, g)
}

func NewRatio() Ratio {
	ratio = &Ratio{
		list: make([]Get, nil),
	}
	return ratio
}

// Create a New Get element
// Must be called in the Get Command of go-ipfs
func NewGet(str String){
	g := &Get{
		CID: str,
		DHT: false,
		BITSWAP: false,
	}
}

func (rt *Ratio) checkBitswap(str String){
	last := rt.list[len(rt.list)-1]
	if (str == last.CID){
		g.BITSWAP = true
	}
	return // retourne une erreur TODO
}

func (rt *Ratio) checkDHT(str String){
	last := rt.list[len(rt.list)-1]
	if (str == last.CID){
		if (last.BITSWAP){
			return // retourne une erreur TODO
		}
		last.DHT = true
	}
	return // retourne une erreur TODO
}