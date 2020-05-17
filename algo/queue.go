/*
Pacakge algo queue
ArrayCycleQueue 循环队列实现:
1. 队列包括 enqueue，dequeue，入队和出队两个方法
2. 队列空，满和更新:
	- head 为队首的索引
	- 判断队列满: size == capbility
	- 判断队列空: size == 0
	- 下一个队尾位置: (head + size) % capbility
	- 下一个队首位置: (head + 1) % capbility
3. 队列扩容 Resize
4. 注意:
	- 在出队入队时，不要忘记增减 size
	- Resize 最后，不要忘记更新 head=0
*/

package algo

// ArrayCycleQueue 循环队列
type ArrayCycleQueue struct {
	data      []int
	size      int
	capbility int
	head      int
}

// NewArrayCycleQueue 创建循环队列
func NewArrayCycleQueue(capbility int) *ArrayCycleQueue {
	return &ArrayCycleQueue{
		data:      make([]int, capbility, capbility),
		capbility: capbility,
	}
}

func (c *ArrayCycleQueue) empty() bool {
	return c.size == 0
}

func (c *ArrayCycleQueue) full() bool {
	return c.size == c.capbility
}

// Resize 队列扩容
func (c *ArrayCycleQueue) Resize(capbility int) {
	if capbility < c.size {
		panic("resize error: capbility < size")
	}
	buf := make([]int, capbility, capbility)
	for i := 0; i < c.size; i++ {
		p := (c.head + i) % c.capbility
		buf[i] = c.data[p]
	}
	c.data = buf
	c.capbility = capbility
	c.head = 0 // 更新 head 为 0
}

// Enqueue 队尾入队
func (c *ArrayCycleQueue) Enqueue(value int) {
	if c.full() {
		c.Resize(2 * c.capbility)
	}
	avl := (c.head + c.size) % c.capbility
	c.data[avl] = value
	c.size++
}

// Dequeue 队首出队
func (c *ArrayCycleQueue) Dequeue() int {
	if c.empty() {
		panic("queue is empty, can not dequque")
	}
	value := c.data[c.head]
	c.head = (c.head + 1) % c.capbility
	if c.size < c.capbility/4 {
		c.Resize(c.capbility / 2)
	}
	c.size--
	return value
}
