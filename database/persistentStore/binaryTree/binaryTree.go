// Online Go compiler to run Golang program online
// Print "Hello World!" message

package binaryTree
import "fmt"

type BTreeNode struct {
    key string
    begByte int64
    endByte int64
	midByte int64
	children []*BTreeNode
}

type BTree struct {
    head *BTreeNode
}

func initBTreeNode(key string, begByte int64, endByte int64)*BTreeNode{
	var midByte int64 = (int64)((begByte + endByte)/2)
    return &BTreeNode{key:key, begByte:begByte, midByte: midByte,endByte:endByte}
}

func addChild(parent *BTreeNode, child *BTreeNode){
	newArray = append(parent.children, child)
	if newArray != nil{
		parent.children = newArray
	}
}

func initBTree(key string, begByte int64, end int64)*BTree{
    return &BTree{tree:initBTreeNode(key, begByte, endByte)}
}

func appendToBTree(btree *BTree, node *BTreeNode){
    
}


func main() {
  fmt.Println("Hello World!")
}




