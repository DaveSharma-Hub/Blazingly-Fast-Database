// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

type BTreeNode struct {
    key string
    begByte int64
    endByte int64
}

type BTree struct {
    tree []*BTreeNode
}

func initBTreeNode(key string, begByte int64, endByte int64)*BTreeNode{
    return &BTreeNode{key:key, begByte:begByte, endByte:endByte}
}

func initBTree(maxSize int)*BTree{
    return &BTree{tree:make([]*BTreeNode,maxSize)}
}

func appendToBTree(btree *BTree, node *BTreeNode){
    var index int = len(btree.tree)-1
    btree.tree[index] = node
}


func main() {
  fmt.Println("Hello World!")
}




