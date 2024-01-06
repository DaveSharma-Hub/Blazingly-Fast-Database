package binaryTree
import (
	"fmt"
	"github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/cache/types"
)

type BTreeNode struct {
    key string
    value globalTypes.Payload
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




