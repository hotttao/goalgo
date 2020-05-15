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
}

// DoubleLink 双链表
type DoubleLink struct {
}

// CycleLink 循环链表
type CycleLink struct {
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
