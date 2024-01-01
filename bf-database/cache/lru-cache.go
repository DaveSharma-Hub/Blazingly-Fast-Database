// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main
import "fmt"

type payload struct{
    value string
}

type Node struct{
    key string
    next *Node
    prev *Node
    item payload
}

type LinkedList struct{
    head *Node
    end *Node
    size int
}

type LRUCache struct{
    list *LinkedList
    lruMap map[string]*Node
    maxSize int
}

func initNode(key string, nodeItem payload)*Node{
    return &Node{ key:key, next: nil, prev: nil, item: nodeItem}
}

func assignNext(node *Node, nextNode *Node){
    node.next = nextNode;
}

func assignPrev(node *Node, prevNode *Node){
    node.prev = prevNode;
}

func getNext(node *Node)*Node{
    return node.next
}
func getPrev(node *Node)*Node{
    return node.prev;
}

func initLinkedList()*LinkedList{
    return &LinkedList{head:nil, end:nil, size:0}
}

func pushToLinkedList(list *LinkedList, node *Node){
    if list.head==nil{
        list.head = node
        list.end = list.head
    }else {
        assignNext(list.end,node)
        list.end = node
    }
    list.size++
}

func LLSize(list* LinkedList)int{
    return list.size
}
func printLL(list *LinkedList){
    node := list.head
    for node!=nil {
        fmt.Println(node.item)
        node = getNext(node)
    }
}

func removeItemFromLinkedList(list *LinkedList, node *Node){
    prevNode := node.prev
    nextNode := node.next

    if prevNode != nil {
        prevNode.next = nextNode
    }
    if nextNode != nil {
        nextNode.prev = prevNode
    }

    if list.head == node {
        list.head = nextNode
    }
    node.next = nil
    node.prev = nil
}

func removeLinkedListHead(list *LinkedList)*Node{
    var head *Node = list.head
    list.head = list.head.next
    assignNext(head, nil)
    assignPrev(head, nil)
    list.size--
    return head
}

func InitLRUCache(maxSize int)*LRUCache{
    return &LRUCache{ list: initLinkedList(), lruMap: make(map[string]*Node), maxSize: maxSize}
}

func GetItem(cache *LRUCache,key string)payload{
    var item *Node = cache.lruMap[key]
    if item!=nil {
        // remove from position
        removeItemFromLinkedList(cache.list, item)
        pushToLinkedList(cache.list,item)
        return item.item
    }else {
        if LLSize(cache.list) == cache.maxSize {
            // remove head
            var removedItem *Node = removeLinkedListHead(cache.list);
            delete(cache.lruMap,removedItem.key)
        }
        //tmp item, would get from perisitent store
        newItem := initNode(key, payload{value:key})
        cache.lruMap[key] = newItem
        pushToLinkedList(cache.list,newItem)
        return newItem.item
    }
}

// list *LinkedList
//     lruMap map[string]*Node

func main() {
  fmt.Println("Hello World!")
//   node1 := initNode("1",payload{value:"1"})
//   node2 := initNode("2",payload{value:"2"})
//   node3 := initNode("3", payload{value:"3"})
//   node4 := initNode("4", payload{value:"4"})
//   node5 := initNode("5",payload{value:"5"})
//   node6 := initNode("6",payload{value:"6"})
//   LL := initLinkedList();
//   pushToLinkedList(LL,node1)
//   pushToLinkedList(LL,node2)
//   pushToLinkedList(LL,node3)
//   pushToLinkedList(LL,node4)
//   pushToLinkedList(LL,node5)
//   pushToLinkedList(LL,node6)
//   printLL(LL)
    lru := initLRUCache(3)
    getItem(lru,"1")
    getItem(lru,"2")
    getItem(lru,"3")
    getItem(lru,"4")
    getItem(lru,"5")
    getItem(lru,"6")
    getItem(lru,"1")
    getItem(lru,"2")
    printLL(lru.list)
}

