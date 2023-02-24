package dataStructure

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
)

type NodeTree struct {
	value int
	left *NodeTree
	right *NodeTree
}


//递归遍历二叉树
func ForeachTree(head *NodeTree, ret []int) []int  {
	if head == nil {
		return ret
	}
	//这里打印（操作）是前序遍历
	//ret = append(ret, head.value)
	ret = ForeachTree(head.left, ret)
	//这里打印（操作）是中序遍历
	ret = append(ret, head.value)
	ret = ForeachTree(head.right, ret)
	//这里打印（操作）是后序遍历
	return ret
}

//非递归实现前序遍历 头左右
func PreOrderUnRecur(head *NodeTree) []int {
	nodeStack := stack.New()
	nodeStack.Push(head)
	var ret []int
	for nodeStack.Len() > 0 {
		cur := (nodeStack.Pop()).(*NodeTree)
		ret = append(ret, cur.value)
		if cur.right != nil {
			nodeStack.Push(cur.right)
		}
		if cur.left != nil {
			nodeStack.Push(cur.left)
		}
	}
	return ret
}

//非递归实现中序遍历 左头右
func InOrderUnRecur(head *NodeTree) []int {
	nodeStack := stack.New()
	var ret []int
	for nodeStack.Len() > 0 || head != nil{
		if head != nil {
			nodeStack.Push(head)
			head = head.left
		} else {
			head = (nodeStack.Pop()).(*NodeTree)
			ret = append(ret, head.value)
			head = head.right
		}
	}
	return ret
}
//非递归实现后序遍历 左右头
func PosOrderUnRecur(head *NodeTree) []int {

}

func Tree()  {
	tree := generateNodeTree()
	var ret []int
	//fmt.Println(PreOrderUnRecur(&tree))

	ret = ForeachTree(&tree,ret)
	fmt.Println(ret)

	fmt.Println(InOrderUnRecur(&tree))
}

func generateNodeTree() NodeTree {
	head := NodeTree{1,nil,nil}
	head.left = &NodeTree{2, nil, nil}
	head.right = &NodeTree{3, nil, nil}

	head.left.left = &NodeTree{4, nil, nil}
	head.left.right = &NodeTree{5, nil, nil}
	head.right.left = &NodeTree{6, nil, nil}
	head.right.right = &NodeTree{7, nil, nil}

	head.left.left.left = &NodeTree{8, nil, nil}
	head.left.left.right = &NodeTree{9, nil, nil}
	head.left.right.left = &NodeTree{10, nil, nil}
	head.left.right.right = &NodeTree{11, nil, nil}
	head.right.left.left = &NodeTree{12, nil, nil}
	head.right.left.right = &NodeTree{13, nil, nil}
	head.right.right.left = &NodeTree{14, nil, nil}
	head.right.right.right = &NodeTree{15, nil, nil}
	
	return head
}