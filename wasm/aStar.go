package main

type Info struct {
	CurrentNode  Node
	SearchedNode int
	isEnd        bool
	isSuccess    bool
	Opened       int
}

func AStarNextStep() func() Info {
	start := NewNode([]int{2, 8, 3, 1, 6, 4, 7, 0, 5})
	end := NewNode([]int{1, 2, 3, 8, 0, 4, 7, 6, 5})
	var visitedList []Node
	var exploreList []Node
	var currentNode Node
	SearchedNode := 0

	start.Value = getValue(start, end)
	visitedList = append(visitedList, start)
	exploreList = append(exploreList, start)

	return func() Info {
		if len(exploreList) == 0 {
			return Info{
				SearchedNode: SearchedNode,
				isEnd:        true,
				isSuccess:    false,
				Opened:       len(exploreList),
			}
		}
		SearchedNode++
		currentNode = pollMinNode(&exploreList)
		if currentNode.Equals(&end) {
			return Info{
				isSuccess:    true,
				isEnd:        true,
				CurrentNode:  currentNode,
				SearchedNode: SearchedNode,
				Opened:       len(exploreList),
			}
		}

		// 进行搜索
		nextNodes := []Node{
			CopyNode(currentNode), CopyNode(currentNode), CopyNode(currentNode), CopyNode(currentNode),
		}
		for i := 0; i < 4; i++ {
			if nextNodes[i].MoveTo(i) && !contains(&visitedList, nextNodes[i]) {
				nextNodes[i].Value = getValue(nextNodes[i], end)
				pushNode(&exploreList, nextNodes[i])
				visitedList = append(visitedList, nextNodes[i])
			}
		}
		// .Println(exploreList)
		return Info{
			isSuccess:    false,
			isEnd:        false,
			CurrentNode:  currentNode,
			SearchedNode: SearchedNode,
			Opened:       len(exploreList),
		}
	}
}

func getValue(node Node, end Node) int {
	s := 0 // 不正确的码数
	d := 0 // 放错的码距离之和
	for i := 0; i < 9; i++ {
		if node.State[i] != end.State[i] {
			s++
		}
		if node.State[i] != 0 {
			x := node.State[i] % 3
			y := node.State[i] / 3
			cx := end.State[i] % 3
			cy := end.State[i] / 3
			dx := x - cx
			dy := y - cy
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			d += int(dx + dy)
		}
	}
	return end.Depth*5 + s*2 + d*4
}

func contains(src *[]Node, node Node) bool {
	for i := range *src {
		if (*src)[i].Equals(&node) {
			return true
		}
	}
	return false
}

func pushNode(src *[]Node, node Node) *[]Node {
	// fmt.Println(*src, node)
	if len(*src) == 0 {
		*src = append(*src, node)
		return src
	}
	for i := range *src {
		if (*src)[i].Value > node.Value {
			*src = append((*src)[:i], node)
			if len(*src) > i + 1 {
				*src = append(*src, (*src)[i + 1:]...)
			}
			break
		}
		if i == len(*src)-1 {
			*src = append(*src, node)
		}
	}
	// fmt.Println(*src)
	return src
}

func pollMinNode(list *[]Node) Node {
	min := (*list)[0]
	if len(*list) > 1 {
		*list = (*list)[1:]
	} else {
		*list = (*list)[:0]
	}
	return min
}
