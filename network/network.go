package network

//Element of the Network
type Element struct {
	ID       int
	Next     *Element
	Previous *Element
}

// Network class
type Network struct {
	Elemetnts []Element
}

// MakeNetwork is the constructor func for Network
func MakeNetwork(numberOfElements int) *Network {
	if numberOfElements < 2 {
		panic("Invalid number")
	}
	network := new(Network)
	for i := 0; i < numberOfElements; i++ {
		network.Elemetnts = append(network.Elemetnts, Element{ID: i})
	}
	return network
}

// getting Element pointer
func (n Network) getElement(id int) *Element {
	var el *Element
	for key := range n.Elemetnts {
		if n.Elemetnts[key].ID == id {
			el = &n.Elemetnts[key]
		}
	}
	return el
}

// getting tail of the current segment of the Network
func (n Network) getTail(el *Element) *Element {
	if el.Next == nil {
		return el
	}
	return n.getTail(el.Next)
}

// getting head of the current segment of the Network
func (n Network) getHead(el *Element) *Element {
	if el.Previous == nil {
		return el
	}
	return n.getHead(el.Previous)
}

// Connect two Elements or segments of the Network
func (n Network) Connect(firstID, secondID int) {

	if isAlreadyConnected := n.Query(firstID, secondID); isAlreadyConnected {
		return
	}


	var (
		firstEl  *Element
		secondEl *Element
	)

	if firstID == secondID {
		panic("You entered same numbers")
	}

	firstEl = n.getElement(firstID)
	secondEl = n.getElement(secondID)

	if firstEl == nil || secondEl == nil {
		panic("Numbers are out of the Network range")
	}

	tail := n.getTail(firstEl)
	head := n.getHead(secondEl)

	tail.Next = head
	head.Previous = tail

}

// checks is elements are connected through the Next pointer
func (n Network) checkNext(el *Element, id int) bool {
	if el.Next == nil {
		return false
	}
	if el.Next.ID == id {
		return true
	}
	return n.checkNext(el.Next, id)
}

// checks is elements are connected through the Previous pointer
func (n Network) checkPrevious(el *Element, id int) bool {
	if el.Previous == nil {
		return false
	}
	if el.Previous.ID == id {
		return true
	}
	return n.checkPrevious(el.Previous, id)
}

// Query checks is elements are connected
func (n Network) Query(firstID, secondID int) bool {
	var (
		firstEl  *Element
		secondEl *Element
	)

	if firstID == secondID {
		panic("You entered same numbers")
	}

	firstEl = n.getElement(firstID)
	secondEl = n.getElement(secondID)

	if firstEl == nil || secondEl == nil {
		panic("Numbers are out of the Network range")
	}

	isConnectedTroughNext := n.checkNext(firstEl, secondID)
	isConnectedTroughPrevious := n.checkPrevious(firstEl, secondID)

	return isConnectedTroughNext || isConnectedTroughPrevious
}
