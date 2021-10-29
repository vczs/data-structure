package main

import "fmt"

/**
树有很多种 二叉树只是树的其中一种
二叉树：每个节点最多只能有两个子节点的一种形式称为二叉树。

第一个节点称为根节点，遍历时需要从根节点遍历。
一个节点的上一层是该节点的父节点，它的下一层是该节点的子节点。
没有子节点的称为叶节点。
*/
type People struct {
	id    int
	name  string
	left  *People
	right *People
}

//前序遍历[先输出根节点 再输出左子树 最后输出右子树]
func preOrder(people *People) {
	if people != nil {
		fmt.Printf("{id:%v  name:%v}\n", people.id, people.name)
		preOrder(people.left)
		preOrder(people.right)
	}
}

//中序遍历[先输出左子树 再输出根节点 最后输出右子树]
func infixOrder(people *People) {
	if people != nil {
		infixOrder(people.left)
		fmt.Printf("{id:%v  name:%v}\n", people.id, people.name)
		infixOrder(people.right)
	}
}

//后序遍历[先输出左子树 再输出右子树 最后输出根节点]
func postOrder(people *People) {
	if people != nil {
		postOrder(people.left)
		postOrder(people.right)
		fmt.Printf("{id:%v  name:%v}\n", people.id, people.name)
	}
}
func main() {

	peopleRoot := &People{id: 1, name: "一"}
	people02 := &People{id: 2, name: "二"}
	people03 := &People{id: 3, name: "三"}
	people04 := &People{id: 4, name: "四"}
	people05 := &People{id: 5, name: "五"}
	people06 := &People{id: 6, name: "六"}

	//建立二叉树
	peopleRoot.left = people02
	peopleRoot.right = people03
	people02.left = people04
	people02.right = people05
	people03.right = people06

	fmt.Println("前序遍历：")
	preOrder(peopleRoot) //前序遍历
	fmt.Println("中序遍历：")
	infixOrder(peopleRoot) //中序遍历
	fmt.Println("后序遍历：")
	postOrder(peopleRoot)
}
