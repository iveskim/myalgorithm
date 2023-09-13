package main

import (
	"fmt"
)

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (node *TrieNode) Insert(word string) {
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			node.children[char] = NewTrieNode()
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (node *TrieNode) Search(word string) bool {
	for _, char := range word {
		if _, exists := node.children[char]; !exists {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func (node *TrieNode) PrintTrie(prefix string) {
	if node == nil {
		return
	}

	fmt.Printf("%s%s\n", prefix, string(' '))

	for char, child := range node.children {
		fmt.Printf("%s└──%s", prefix, string(char))
		if child.isEnd {
			fmt.Println(" (end)")
		} else {
			fmt.Println()
		}
		child.PrintTrie(prefix + "│  ")
	}
}

func main() {
	root := NewTrieNode()

	words := []string{"apple", "applet", "bat", "cat", "bee", "beer"}

	for _, word := range words {
		root.Insert(word)
	}

	fmt.Println(root.Search("apple")) // 输出 true
	fmt.Println(root.Search("app"))   // 输出 false
	fmt.Println(root.Search("beer"))  // 输出 true
	fmt.Println(root.Search("batt"))  // 输出 false

	root.PrintTrie("")
}
