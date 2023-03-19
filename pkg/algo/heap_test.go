package algo

import "testing"

func TestHeap(t *testing.T) {
	data := []int{19, 3, 8, 20}
	h := BuildHeap(data)
	t.Log("Buil Heap")
	t.Log(h.data[1 : h.len+1])
	if h.Pop() != 20 {
		t.Errorf("Heap top expect 20")
	}
	t.Log("Pop 20")
	t.Log(h.data[1 : h.len+1])

	h.Add(30)
	t.Log("Add 30")
	t.Log(h.data[1 : h.len+1])
	if h.Pop() != 30 {
		t.Log(h.data[1 : h.len+1])
		t.Errorf("Heap top expect 30")
	}
	t.Log("Pop 30")
	t.Log(h.data[1 : h.len+1])

	j := NewHeap(2)
	j.Add(10)
	t.Log("Add 10")
	t.Log(j)
	j.Add(1)
	t.Log("Add 1")
	t.Log(j)
	j.Add(3)
	t.Log("Add 3")
	t.Log(j)
	j.Add(20)
	t.Log("Add 20")
	t.Log(j)
}
