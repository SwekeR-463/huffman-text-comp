package main

import(
	"fmt"
	"os"
	"container/heap"
)



//HUFFMAN-ALGO
//huffman|c|
//n=|c|
//build min-heap Q with c
//for i=1 to n-1
//allocate a new node Z
//Z.left = X = extract_min(Q)
//Z.right = Y = extract_min(Q)
//Z.freq = X.freq + Y.freq
//insert(Z,Q)
//return (extract_min(Q))



//representing a node in the huffman tree
type HuffmanNode struct {
	char rune
	//rune -> represent any character in the unicode standard
	freq int
	left *HuffmanNode
	right *HuffmanNode
}


//priority queue for implementing a min-heap for node
type PriorityQueue []*HuffmanNode


func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i,j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq PriorityQueue) Swap(i,j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	//interface -> a type that lists methods without providing their code
	*pq = append(*pq, x.(*HuffmanNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}


//buiding the huffman tree
func buildHuffmanTree(freqMap map[rune]int) *HuffmanNode {
	pq := &PriorityQueue{}
	heap.Init(pq)

	//adding the chars to priority queue
	for char, freq := range freqMap {
		heap.Push(pq, &HuffmanNode{char : char, freq : freq})
	}

	//building the tree
	for pq.Len() > 1 {
		left := heap.Pop(pq).(*HuffmanNode)
		right := heap.Pop(pq).(*HuffmanNode)
		newNode := &HuffmanNode{
			freq: left.freq + right.freq,
			left: left,
			right: right,
		}
		heap.Push(pq, newNode)
	}
	return heap.Pop(pq).(*HuffmanNode)
}


//generating the huffman codes
func generateCodes(node *HuffmanNode, code string, codes map[rune]string) {
	if node == nil {
		return
	}
	if node.left == nil && node.right == nil {
		codes[node.char] = code
		return
	}
	generateCodes(node.left, code+"0", codes)
	generateCodes(node.right, code+"1", codes)
}


//encoding the text
func encodeText(text string, codes map[rune]string) string {
	encoded := ""
	for _, char := range text {
		encoded += codes[char]
	}
	return encoded
}

//decoding the text
func decodeText(encoded string, root *HuffmanNode) string {
	result := ""
	node := root
	for _, bit := range encoded {
		if bit == '0' {
			node = node.left
		} else {
			node = node.right
		}

		if node.left == nil && node.right == nil {
			result += string(node.char)
			node = root
		}
	}
	return result
}


//reading file and counting frequencies
func readFile(filename string) (string, map[rune]int) {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	text := string(data)

	freqMap := make(map[rune]int)
	for _, char := range text {
		freqMap[char]++
	}

	return text, freqMap
}


//writing file
func writeFile(filename, content string) {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		panic(fmt.Sprintf("Error writing to file: %v", err))
	}
}



func main() {
	
	//getting of the text file
	inputFile := "input.txt"

	text, freqMap := readFile(inputFile)
	
	//building of huffman tree
	root := buildHuffmanTree(freqMap)
	
	//generating codes
	codes := make(map[rune]string)
	generateCodes(root, "", codes)
	
	//encoding
	encoded := encodeText(text, codes)
	fmt.Println("Encoded Text:", encoded)
	
	//decoding
	decoded := decodeText(encoded, root)
	fmt.Println("Decoded Text:", decoded)
	
	//saving in another file
	outputFile := "encoded.txt"
	writeFile(outputFile, encoded)
	fmt.Println("Encoded data saved to", outputFile)
	
	if decoded == text {
		fmt.Println("Decoding successful!")
	} else {
		fmt.Println("Decoding failed!")
	}
}