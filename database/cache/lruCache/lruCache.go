package lruCache

import (
	"fmt"
    "github.com/DaveSharma-Hub/Blazingly-Fast-Database/database/types"
)

type Node struct{
    key string
    next *Node
    prev *Node
    item globalTypes.Payload
}

type PersistedItemConversion func(key string) globalTypes.Payload

type LinkedList struct{
    head *Node
    end *Node
    size int64
}

type LRUCache struct{
    list *LinkedList
    lruMap map[string]*Node
    maxSize int64
}

func initNode(key string, nodeItem globalTypes.Payload)*Node{
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

func LLSize(list* LinkedList)int64{
    return int64(list.size)
}

func PrintItems(cache *LRUCache){
	list := cache.list
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

func InitLRUCache(maxSize int64)*LRUCache{
    return &LRUCache{ list: initLinkedList(), lruMap: make(map[string]*Node), maxSize: maxSize}
}

func GetItem(cache *LRUCache,key string, fnCallback PersistedItemConversion)globalTypes.Payload{
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
		value := fnCallback(key)
        newItem := initNode(key, value)
        cache.lruMap[key] = newItem
        pushToLinkedList(cache.list,newItem)
        return newItem.item
    }
}

func SetItem(cache *LRUCache, key string, value globalTypes.Payload){
    if LLSize(cache.list) == cache.maxSize {
        // remove head
        var removedItem *Node = removeLinkedListHead(cache.list);
        delete(cache.lruMap,removedItem.key)
    }
    //tmp item, would get from perisitent store
    newItem := initNode(key, value)
    cache.lruMap[key] = newItem
    pushToLinkedList(cache.list,newItem)
}

func UpdateItem(cache *LRUCache, key string, value globalTypes.Payload){
    var item *Node = cache.lruMap[key]
    newItem := initNode(key, value)
    if item!=nil {
        // remove from position
        removeItemFromLinkedList(cache.list, item)
        pushToLinkedList(cache.list,newItem)
    }else {
        if LLSize(cache.list) == cache.maxSize {
            // remove head
            var removedItem *Node = removeLinkedListHead(cache.list);
            delete(cache.lruMap,removedItem.key)
        }
        //tmp item, would get from perisitent store
        cache.lruMap[key] = newItem
        pushToLinkedList(cache.list,newItem)
    }
}

func RemoveItem(cache *LRUCache, key string){
    tmpNode := cache.lruMap[key]
    if tmpNode != nil{
        //does exist

        if cache.list.head == tmpNode {
            cache.list.head = tmpNode.next
            if tmpNode.next != nil {
                tmpNode.next.prev = nil
            }
        }else if cache.list.end == tmpNode {
            cache.list.end = tmpNode.prev
            if tmpNode.prev != nil {
                tmpNode.prev.next = nil
            }
        }else {
            tmpNode.prev.next = tmpNode.next
            tmpNode.next.prev = tmpNode.prev
        }
        delete(cache.lruMap, key)
        tmpNode.next = nil
        tmpNode.prev = nil
    }
}

// list *LinkedList
//     lruMap map[string]*Node

