package mafsa

type node struct {
	first, last *node // linked list of children
	next        *node

	isWord bool

	b byte
}

type builder struct {
	root *node
}

func (b *builder) InsertWord(s string) {
	if b.root == nil {
		b.root = &node{}
	}

	current := b.root
	i := 0
	for ; i < len(s); i++ {
		b := s[i]

		if current.first == nil {
			break
		}

		if current.last.b < b {
			break
		}

		current = current.last
	}

	if i == len(s) {
		current.isWord = true
		return
	}

	for ; i < len(s); i++ {
		n := &node{b: s[i]}

		if current.last == nil {
			current.first = n
			current.last = n
			current = n
		} else {
			tmp := current.last
			current.last = n
			tmp.next = n
			current = n
		}
	}

	current.isWord = true
}

func (b *builder) Contains(s string) bool {
	if b.root == nil {
		return false
	}

	current := b.root

Outer:
	for len(s) > 0 {
		b := s[0]
		s = s[1:]

		for cand := current.first; cand != nil; cand = cand.next {
			if cand.b == b {
				current = cand
				continue Outer
			}
		}

		return false
	}

	return current.isWord
}
