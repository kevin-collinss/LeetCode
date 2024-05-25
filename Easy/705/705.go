package Easy

// create a ListNode with a key and a next pointing to the node that comes after it
type MyListNode struct {
	Key  int
	Next *MyListNode
}

// Hashset is an array of pointers to ListNodes
type MyHashSet struct {
	data []*MyListNode
}

func Constructor() MyHashSet {
	return MyHashSet{
		//creates a slice of pointers to MyListNode with a length of 10000
		data: make([]*MyListNode, 10000),
	}
}

// keeps the index between 1 -> 10000
func hash(key int) int {
	return key % 10000
}

func (this *MyHashSet) Add(key int) {

	//hashes the key
	index := hash(key)

	//if index is empty
	if this.data[index] == nil {
		//populate with list node with the key being param key
		this.data[index] = &MyListNode{Key: key}
	} else {
		//cur is the Listnode at the index
		cur := this.data[index]
		//if cur isnt empty
		for cur != nil {
			//if key of current listnode = key
			if cur.Key == key {
				return
			}

			//if the next is nil
			if cur.Next == nil {
				//add the ListNode to end of the the LinkedList
				cur.Next = &MyListNode{Key: key}
				return
			}

			//cur then is the next node
			cur = cur.Next
		}
	}
}

func (this *MyHashSet) Remove(key int) {

	index := hash(key)
	cur := this.data[index]

	if cur == nil {
		return
	}
	if cur.Key == key {
		this.data[index] = cur.Next
		return
	}
	prev := cur
	for cur != nil {
		if cur.Key == key {
			prev.Next = cur.Next
			return
		}
		prev = cur
		cur = cur.Next
	}

}

func (this *MyHashSet) Contains(key int) bool {
	index := hash(key)
	cur := this.data[index]

	for cur != nil {
		if cur.Key == key {
			return true
		}
		cur = cur.Next
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
