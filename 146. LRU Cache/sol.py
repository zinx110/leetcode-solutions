class Node:
    def __init__(self, key: int, value: int):
        self.key = key
        self.value = value
        self.prev = self.next = None
        
class LRUCache:

    def __init__(self, capacity: int):
        self.cache = {}
        
        self.left = Node(0, 0)
        self.right = Node(0, 0)
        self.left.next = self.right
        self.right.prev = self.left

        self.capacity = capacity
        self.size = 0
        
    def remove(self, node: Node) -> None:
        print("  remove :", node.key, node.value)
        prev = node.prev
        next = node.next

        prev.next = next
        next.prev = prev


    
    def insert(self, node: Node) -> None:
        print("  insert :", node.key, node.value)
        prev = self.right.prev
        prev.next = node
        self.right.prev = node
        node.prev = prev
        node.next = self.right



    def get(self, key: int) -> int:
        print("get :", key)
        if key in self.cache:
            self.remove(self.cache[key])
            self.insert(self.cache[key])
            return self.cache[key].value
        return -1

    def put(self, key: int, value: int) -> None:
        print("put :", key, value)
        if key in self.cache:
            self.remove(self.cache[key])
            self.insert(self.cache[key])
            self.cache[key].value = value
            return 

        if self.size < self.capacity:
            self.size += 1
        else:
            first = self.left.next
            print("   removing first : ", first.key, first.value)
            self.remove(first)
            del self.cache[first.key]

        node = Node(key, value)
        self.insert(node)
        self.cache[key] = node


    



        


# Your LRUCache object will be instantiated and called as such:
# obj = LRUCache(capacity)
# param_1 = obj.get(key)
# obj.put(key,value)