package main

import "fmt"

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		if node.next[word[i]-'a'] == nil {
			node.next[word[i]-'a'] = new(Trie)
		}
		node = node.next[word[i]-'a']
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for i := 0; i < len(word); i++ {
		if node.next[word[i]-'a'] == nil {
			return false
		}
		node = node.next[word[i]-'a']
	}
	return node.isEnd
}

func (t *Trie) StartWith(word string) bool {
	node := t
	for i := 0; i < len(word); i++ {
		if node.next[word[i]-'a'] == nil {
			return false
		}
		node = node.next[word[i]-'a']
	}
	return true
}

func main() {
	preTrie := Constructor()
	preTrie.Insert("xiaocai")
	preTrie.Insert("hello")

	fmt.Println("xiaocai: %v", preTrie.Search("xiaocai"))
	fmt.Println("xiaocai: %v", preTrie.Search("cai"))
}
