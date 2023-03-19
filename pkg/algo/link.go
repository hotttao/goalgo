/*
Package algo link
SingleLink 实现:
1. 使用 head，tail 指向链表首尾
2. 让 head 指向哨兵节点，简化链表操作
3. 定义 remove_head, add_head, add_tail 方法
4. 删除节点时要注意是否删除的是尾节点，如果是需要更新 tail 的指向
5. 单链表不适合实现删除尾节点，故不实现 remove_tail

DoubleLink 实现可以通过以下方法简化
1. 使用首位两个哨兵节点简化双链表操作
2. 定义 一个在两个节点之间插入节点的方法 InsertBetween(before, after) 简化插入操作
3. 定义删除节点 DeleteNode(node) 方法简化节点删除，因为双向链表能拿到前后节点，所以能实现此方法

CycleLinke 实现:
1. 因为是循环链表，我们使用一个 tail 指向链表的尾节点即可
2. 循环链表同样无法快速删除尾部节点
3. 插入头节点和插入尾节点的唯一区别是，最后是否更改 tail 的指向
4. 因此循环节点，只需要实现 AddTail 添加尾节点，和 RemoveHead 删除头节点即可*/
package algo

// LinkNode 链表节点
type LinkNode struct {
	Key   int
	Value int
	Pre   *LinkNode
	Nxt   *LinkNode
}

// SingleLink 单链表
type SingleLink struct {
	tail *LinkNode
	head *LinkNode
	size int
}

// NewSingleLink 创建一个单链表
func NewSingleLink() SingleLink {
	guard := &LinkNode{}
	s := SingleLink{
		head: guard,
		tail: guard,
	}
	return s
}

// AddHead 添加头节点
func (s *SingleLink) AddHead(k, v int) {
	node := LinkNode{
		Key:   k,
		Value: v,
	}
	node.Nxt = s.head.Nxt
	s.head.Nxt = &node
	s.size++
}

// RemoveHead 删除头节点
func (s *SingleLink) RemoveHead() *LinkNode {
	if s.size == 0 {
		return nil
	}
	pwd := s.head.Nxt // 当前头节点
	s.head.Nxt = pwd.Nxt
	if pwd == s.tail {
		s.tail = s.head
	}
	s.size--
	return pwd
}

// AddTail 添加尾节点
func (s *SingleLink) AddTail(k, v int) {
	node := LinkNode{
		Key:   k,
		Value: v,
	}
	s.tail.Nxt = &node
	s.tail = &node
	s.size++
}

// DoubleLink 双链表
type DoubleLink struct {
	tail *LinkNode
	head *LinkNode
	size int
}

// NewDoubleLink 创建一个 双向链表
func NewDoubleLink() *DoubleLink {
	guardHead := &LinkNode{}
	gurarTail := &LinkNode{}
	guardHead.Nxt = gurarTail
	gurarTail.Pre = guardHead
	d := &DoubleLink{
		head: guardHead,
		tail: gurarTail,
	}
	return d
}

// InsertBetween 在两个节点之间插入新的节点
func (d *DoubleLink) InsertBetween(k, v int, pre *LinkNode, next *LinkNode) {
	node := &LinkNode{
		Key:   k,
		Value: v,
		Pre:   pre,
		Nxt:   next,
	}
	pre.Nxt = node
	next.Pre = node
	d.size++
}

// DeleteNode 删除两个节点中间的节点
func (d *DoubleLink) DeleteNode(node *LinkNode) *LinkNode {
	pre := node.Pre
	next := node.Nxt
	pre.Nxt = next
	next.Pre = pre
	node.Pre = nil
	node.Nxt = nil
	d.size--
	return node
}

// AddHead 添加头节点
func (d *DoubleLink) AddHead(k, v int) {
	pre := d.head
	next := d.head.Nxt
	d.InsertBetween(k, v, pre, next)
}

// RemoveHead 删除头节点
func (d *DoubleLink) RemoveHead() *LinkNode {
	if d.size == 0 {
		panic("DoubleLink empty")
	}
	return d.DeleteNode(d.head.Nxt)
}

// CycleLink 循环链表
type CycleLink struct {
	tail *LinkNode
	size int
}

// NewCycleLinke 创建一个循环链表
func NewCycleLinke() CycleLink {
	c := CycleLink{}
	return c
}

// AddTail 添加尾节点
func (c *CycleLink) AddTail(k, v int) {
	node := &LinkNode{
		Key:   k,
		Value: v,
	}
	if c.size == 0 {
		node.Nxt = node
	} else {
		node.Nxt = c.tail.Nxt
		c.tail.Nxt = node
	}
	c.tail = node
	c.size++
}

// RemoveHead 删除头节点
func (c *CycleLink) RemoveHead() *LinkNode {
	if c.size == 0 {
		return nil
	}
	node := c.tail.Nxt
	c.size--
	if c.size == 1 {
		c.tail = nil
	} else {
		c.tail.Nxt = node.Nxt
	}
	node.Nxt = nil
	return node
}

// LinkReverse 单链表反转
func LinkReverse(head *LinkNode) *LinkNode {
	if head == nil || head.Nxt == nil {
		return head
	}
	var pre *LinkNode // pre 必须为 nil，否则下面的循环无法终止
	pwd := head
	for pwd != nil {
		pre, pwd, pwd.Nxt = pwd, pwd.Nxt, pre
	}
	return pre
}

// LinkHasCycle 单链表环检测
func LinkHasCycle(head *LinkNode) bool {
	if head == nil {
		return false
	}
	s := head
	d := head.Nxt
	for d != nil && d.Nxt != nil {
		s = s.Nxt
		d = d.Nxt.Nxt
		if s == d {
			return true
		}
	}
	return false
}

// LinkMerge 两个有序单链表的合并
func LinkMerge(l1, l2 *LinkNode) *LinkNode {
	// 这里必须使用 New
	// var link LinkNode 会导致 c.Next 无法赋值
	// var link *LinkNode 第一循环 c.Next 为无效引用
	link := new(LinkNode)
	a := l1
	b := l2
	c := link // 当前节点的引用
	for a != nil && b != nil {
		if a.Value > b.Value {
			c.Nxt = b
			b = b.Nxt
		} else {
			c.Nxt = a
			a = a.Nxt
		}
		c = c.Nxt
	}
	if a == nil {
		c.Nxt = b
	}
	if b == nil {
		c.Nxt = a
	}
	return link.Nxt
}

// LinkDeleteLastN 删除单链表倒数第 N 个节点
func LinkDeleteLastN(head *LinkNode, n int) *LinkNode {
	if head == nil {
		return head
	}
	var pre *LinkNode
	pwd := head
	p := head
	// p 为 head，i 从 1 开始，循环后  p 指向第 n 个节点
	for i := 1; i < n; i++ {
		if head == nil {
			return nil
		}
		p = p.Nxt // 注意不是 head.Nxt
	}
	// p 指向最后一个节点时，pwd 就指向倒数第 N 个节点
	for p.Nxt != nil {
		pre = pwd
		pwd = pwd.Nxt
		p = p.Nxt
	}
	// pre 为 nil 则删除的是第一个节点，直接返回 pwd.Nxt
	if pre == nil {
		return pwd.Nxt
	}
	// 删除倒数第 N 个节点
	pre.Nxt = pwd.Nxt
	return head
}

// LinkMiddle 取链表的中间节点
func LinkMiddle(head *LinkNode) *LinkNode {
	single := head
	double := head
	for double.Nxt != nil && double.Nxt.Nxt != nil {
		single = single.Nxt
		double = double.Nxt.Nxt
	}
	// 奇数节点，double 会指向链表的尾部 double.Nxt 为 nil
	// 偶数节点，double 会指向链表的导数第二个节点 double.Nxt 不为 nil
	if double.Nxt == nil {
		return single
	}
	return single.Nxt
}
