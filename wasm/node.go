package main

// Node 节点
type Node struct {
	State  []int // 当前状态
	Depth  int   // 当前深度
	Parent *Node // 父节点
	Value  int   // 代价估计值
	Zero   int   // 0 的位置
}

const (
	MoveUp = iota
	MoveDown
	MoveLeft
	MoveRight
)

// NewNode 新建节点
func NewNode(data []int) Node {
	var z int
	for i, c := range data {
		if c == 0 {
			z = i
			break
		}
	}
	return Node{
		State:  data,
		Depth:  0,
		Parent: nil,
		Value:  0,
		Zero:   z,
	}
}

// CopyNode 复制节点
func CopyNode(old Node) Node {
	newState := make([]int, 9)
	copy(newState, old.State)
	return Node{
		State:  newState,
		Depth:  old.Depth,
		Parent: old.Parent,
		Value:  old.Value,
		Zero:   old.Zero,
	}
}

// Equals 是否相等
func (n *Node) Equals(other *Node) bool {
	if len(n.State) != 9 || len(n.State) != len(other.State) {
		return false
	}
	for i := range n.State {
		if n.State[i] != other.State[i] {
			return false
		}
	}
	return true
}

// IsValid 是否有效
func (n *Node) IsValid() bool {
	if len(n.State) != 9 || n.Zero < 0 || n.Zero > 8 || n.State[n.Zero] != 0 {
		return false
	}
	var has [8]bool
	for _, n := range n.State {
		if n < 0 || n > 8 {
			return false
		}
		has[n] = true
	}
	for _, b := range has {
		if b == false {
			return false
		}
	}
	return true
}

// ToString 转化为字符串
func (n *Node) ToString() string {
	var str string
	for _, n := range n.State {
		str += string('0' + n)
		str += ","
	}
	return str
}

// CanMove 是否可以上下左右移动
func (n *Node) CanMove() []bool {
	movable := make([]bool, 4)
	for i := 0; i < 4; i++ {
		movable[i] = n.CanMoveTo(i)
	}
	return movable
}

// CanMoveTo 是否能移动
func (n *Node) CanMoveTo(where int) bool {
	switch where {
	case MoveUp:
		return n.Zero > 3 && n.Zero < 9
	case MoveDown:
		return n.Zero < 6 && n.Zero >= 0
	case MoveLeft:
		return n.Zero%3 > 0
	case MoveRight:
		return n.Zero%3 < 2
	}
	return false
}

// MoveTo 移动
func (n *Node) MoveTo(where int) bool {
	switch where {
	case MoveUp:
		return n.MoveEmptyUp()
	case MoveDown:
		return n.MoveEmptyDown()
	case MoveLeft:
		return n.MoveEmptyLeft()
	case MoveRight:
		return n.MoveEmptyRight()
	default:
		return false
	}
}

func (n *Node) MoveEmptyUp() bool {
	if !n.CanMoveTo(MoveUp) {
		return false
	}
	p := CopyNode(*n)
	n.Parent = &p
	n.Depth++
	n.State[n.Zero], n.State[n.Zero-3] = n.State[n.Zero-3], n.State[n.Zero]
	n.Zero -= 3
	return true
}

func (n *Node) MoveEmptyDown() bool {
	if !n.CanMoveTo(MoveDown) {
		return false
	}
	p := CopyNode(*n)
	n.Parent = &p
	n.Depth++
	n.State[n.Zero], n.State[n.Zero+3] = n.State[n.Zero+3], n.State[n.Zero]
	n.Zero += 3
	return true
}

func (n *Node) MoveEmptyLeft() bool {
	if !n.CanMoveTo(MoveLeft) {
		return false
	}
	p := CopyNode(*n)
	n.Parent = &p
	n.Depth++
	n.State[n.Zero], n.State[n.Zero-1] = n.State[n.Zero-1], n.State[n.Zero]
	n.Zero--
	return true
}

func (n *Node) MoveEmptyRight() bool {
	if !n.CanMoveTo(MoveRight) {
		return false
	}
	p := CopyNode(*n)
	n.Parent = &p
	n.Depth++
	n.State[n.Zero], n.State[n.Zero+1] = n.State[n.Zero+1], n.State[n.Zero]
	n.Zero++
	return true
}
