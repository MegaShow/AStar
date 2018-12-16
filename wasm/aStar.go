package main

type Info struct {
	CurrentNode  Node
	SearchedNode int
	isEnd        bool
	isSuccess    bool
	Opened       int
}

func getInvSum(arr []int) int {
	sum := 0
	for i := range arr {
		if arr[i] != 0 {
			for j := 0; j < i; j++ {
				if arr[j] > arr[i] {
					sum++
				}
			}
		}
	}
	// fmt.Println(sum)
	return sum
}

func isValid(start, end []int) bool {
	return getInvSum(start)%2 == getInvSum(end)%2
}

func AStarNextStep() func() Info {
	// 复杂搜索
	startArr := []int{3, 6, 1, 2, 8, 7, 4, 5, 0}
	endArr := []int{1, 2, 3, 8, 0, 4, 7, 6, 5}

	// 简单搜索
	// startArr := []int{2, 8, 3, 1, 6, 4, 7, 0, 5}
	// endArr := []int{1, 2, 3, 8, 0, 4, 7, 6, 5}


	if !isValid(startArr, endArr) {
		return func() Info {
			return Info{
				SearchedNode: 0,
				isEnd:        true,
				isSuccess:    false,
				Opened:       0,
			}
		}
	}
	start := NewNode(startArr)
	end := NewNode(endArr)
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
			currentNode.Value = getValue(currentNode, end)
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
		if end.State[i] != 0 && node.State[i] != end.State[i] {
			s++
		}
		if node.State[i] != 0 {
			x := i % 3
			y := i / 3
			var cx, cy int
			for j := 0; j < 9; j++ {
				if node.State[i] == end.State[j] {
					cx = j % 3
					cy = j / 3
					break
				}
			}
			dx := x - cx
			dy := y - cy
			if dx < 0 {
				dx = -dx
			}
			if dy < 0 {
				dy = -dy
			}
			d += dx + dy
		}
	}
	return node.Depth + d // node.Depth + s*2 + d*4
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
	// fmt.Println("s:", *src, node)
	if len(*src) == 0 {
		*src = append(*src, node)
		return src
	}
	for i := range *src {
		if (*src)[i].Value > node.Value {
			 ori :=  make([]Node, len(*src))
			copy(ori, *src)
			*src = append(ori[:i], node)
			if len(ori) >= i {
				*src = append(*src, ori[i:]...)
			}
			break
		}
		if i == len(*src)-1 {
			*src = append(*src, node)
		}
	}
	// fmt.Println("after:", *src)
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
