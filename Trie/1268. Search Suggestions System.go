package Trie

import (
	"sort"
)

type TrieNodeSP struct {
	Children map[rune]*TrieNodeSP // Menyimpan child nodes berdasarkan karakter
	Words    []string             // Menyimpan hingga 3 saran produk
}

// Fungsi utama untuk memberikan saran produk
func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)

	// Inisialisasi root Trie
	root := &TrieNodeSP{Children: make(map[rune]*TrieNodeSP)}

	// Bangun Trie dengan produk
	for _, product := range products {
		node := root
		for _, char := range product {
			if node.Children[char] == nil {
				node.Children[char] = &TrieNodeSP{Children: make(map[rune]*TrieNodeSP)}
			}
			node = node.Children[char]
			// Simpan produk dalam node jika kurang dari 3 produk
			if len(node.Words) < 3 {
				node.Words = append(node.Words, product)
			}
		}
	}

	// Hasil untuk menyimpan daftar saran berdasarkan prefix dari searchWord
	result := make([][]string, 0)
	node := root
	for _, char := range searchWord {
		if node != nil {
			node = node.Children[char]
		}
		// Tambahkan saran produk atau [] jika node kosong
		if node == nil {
			result = append(result, []string{})
		} else {
			result = append(result, node.Words)
		}
	}

	return result
}
