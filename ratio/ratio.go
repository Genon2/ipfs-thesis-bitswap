package ratio


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
	list []*Get
}

// Add a new Get to the Ratio
func (rt *Ratio) Add(str string){
	var g = NewGet(str)
	rt.list = append(rt.list, g)
}

func NewRatio() Ratio {
	var ratio = &Ratio{
		list: make([]*Get, 0),
	}
	return *ratio
}

func (rt *Ratio) ToString() string{
	var last = rt.list[len(rt.list)-1]
	var str = last.CID +" DHT : " + ConvertBool(last.DHT)+ " BITSWAP : " + ConvertBool(last.BITSWAP)
	return str
}
// Create a New Get element
// Must be called in the Get Command of go-ipfs
func NewGet(str string) *Get{
	g := &Get{
		CID: str,
		DHT: false,
		BITSWAP: false,
	}
	return g
}

func (rt *Ratio) CheckBitswap(str string){
	last := rt.list[len(rt.list)-1]
	if (str == last.CID){
		last.BITSWAP = true
	}
	return // retourne une erreur TODO
}

func (rt *Ratio) CheckDHT(str string){
	last := rt.list[len(rt.list)-1]
	if (str == last.CID){
		if (last.BITSWAP){
			return // retourne une erreur TODO
		}
		last.DHT = true
	}
	return // retourne une erreur TODO
}

func ConvertBool(b bool) string{
	if b {
		return "True"
	}
	return "False"
}