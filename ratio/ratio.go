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
	var list []Get
}

// Add a new Get to the Ratio
func (rt *Ratio) Add(str String){
	g = NewGet(str)
	rt.list = append(rt.list, g)
}

// Create a New Get element
// Must be called in the Get Command of go-ipfs
func NewGet(str String){
	g := Get{
		CID: str,
		DHT: false,
		BITSWAP: false,
	}
}

func (g *Get) checkBitswap(str String){
	if (str == g.CID){
		if (g.DHT){
			return // retourne une erreur TODO
		}
		g.BITSWAP = true
	}
	return // retourne une erreur TODO
}

func (g *Get) checkDHT(str String){
	if (str == g.CID){
		if (g.BITSWAP){
			return // retourne une erreur TODO
		}
		g.DHT = true
	}
	return // retourne une erreur TODO
}