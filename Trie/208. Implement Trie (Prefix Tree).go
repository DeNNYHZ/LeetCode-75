package Trie

type TrieNode struct {
	Children map[byte]*TrieNode // Child nodes untuk setiap karakter
	IsEnd    bool               // Penanda akhir kata

}

type Trie struct {
	Root *TrieNode // Node root dari Trie
}

// Fungsi constructor untuk inisialisasi Trie baru
func Constructor() Trie {
	return Trie{
		Root: &TrieNode{Children: make(map[byte]*TrieNode)},
	}
}

// Fungsi Insert untuk menambahkan kata baru ke dalam Trie
func (this *Trie) Insert(word string) {
	node := this.Root
	for i := 0; i < len(word); i++ {
		char := word[i]
		// Jika karakter belum ada dalam children, buat node baru
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = &TrieNode{Children: make(map[byte]*TrieNode)}
		}
		// Lanjut ke child node
		node = node.Children[char]
	}
	// Set IsEnd menjadi true di akhir kata
	node.IsEnd = true
}

// Fungsi Search untuk memeriksa apakah sebuah kata ada di dalam Trie
func (this *Trie) Search(word string) bool {
	node := this.Root
	for i := 0; i < len(word); i++ {
		char := word[i]
		// Jika karakter tidak ditemukan dalam children, return false
		if _, exists := node.Children[char]; !exists {
			return false
		}
		node = node.Children[char]
	}
	// Return true jika berada di akhir kata yang valid
	return node.IsEnd
}

// Fungsi StartsWith untuk memeriksa apakah ada kata yang dimulai dengan prefix tertentu
func (this *Trie) StartsWith(prefix string) bool {
	node := this.Root
	for i := 0; i < len(prefix); i++ {
		char := prefix[i]
		// Jika karakter tidak ditemukan dalam children, return false
		if _, exists := node.Children[char]; !exists {
			return false
		}
		node = node.Children[char]
	}
	// Return true karena prefix ditemukan
	return true
}

/**
 * Contoh cara menggunakan Trie:
 * obj := Constructor();
 * obj.Insert("hello");
 * param_2 := obj.Search("hello"); // true jika "hello" ada dalam Trie
 * param_3 := obj.StartsWith("hel"); // true jika ada kata yang dimulai dengan "hel"
 */
