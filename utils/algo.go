package utils

import (
	"bytes"
	"fmt"
	"sync"
)

/*
I am implementing very naive algorithm to generate short url.
Idea is generate sequential number and convert it to string.
Generate a pull of ids in batch and use them one by one.
*/

// I need number to byte conversion map[byte]byte number 97 should give 'a'

var Chars = [...]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't',
	'u', 'v', 'w', 'x', 'y', 'z', '0', '1', '2', '3',
	'4', '5', '6', '7', '8', '9', 'A', 'B', 'C', 'D',
	'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N',
	'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
	'Y', 'Z',
}

type Stack struct {
	s  [16]byte
	mu sync.Mutex
}

var stack Stack
var SmallBatch *Batch

func init() {
	stack = Stack{
		s: [16]uint8{
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
		mu: sync.Mutex{},
	}

	SmallBatch = NewBatch(2)
}

func (s *Stack) NextString() string {
	s.mu.Lock()
	defer s.mu.Unlock()
	var w bytes.Buffer
	for i := 15; i >= 0; i-- {
		fmt.Printf("i: %d, s[i]: %c\n", i, Chars[s.s[i]])
		if s.s[i] == 61 {
			s.s[i] = 0
			s.s[i-1]++
			continue
		} else {
			s.s[i]++
			break
		}
	}
	for i := 0; i < len(s.s); i++ {
		if e := w.WriteByte(Chars[s.s[i]]); e != nil {
			panic(e)
		}
	}
	return w.String()
}

const DefaultBatchSize = 100

type Node struct {
	Value string
	Used  bool
}

func NewBatch(size uint8) *Batch {
	if size == 0 {
		size = DefaultBatchSize
	}
	return &Batch{BatchSize: size, Used: true}
}

type Batch struct {
	Nodes     []Node
	Used      bool
	mu        sync.Mutex
	BatchSize uint8
}

func (b *Batch) NextBatch() {
	b.Nodes = []Node{}
	for i := 0; i < int(b.BatchSize); i++ {
		Node := Node{Value: stack.NextString(), Used: false}
		b.Nodes = append(b.Nodes, Node)
	}
}

func (b *Batch) GetNextShortUrl() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	fmt.Printf("is batch used: %t\n", b.Used)
	// All nodes are used
	if b.Used {
		b.NextBatch()
		b.Used = false
	}

	println("Batch is not used")
	for i, n := range b.Nodes {
		fmt.Printf("Node: %+v\n", n)
		if !n.Used {
			b.Nodes[i].Used = true
			if i == len(b.Nodes)-1 {
				b.Used = true
			}
			return b.Nodes[i].Value
		}
	}
	println("Batch is used")
	return ""
}
