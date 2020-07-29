class Node(object):
    def __init__(self, val):
        self.val = val
        self.next = None

    def push(self, val):
        new_node = Node(val)
        curr = self
        while curr.next != None:
            curr = curr.next
        curr.next = new_node
        return

    def print(self):
        curr = self
        while curr != None:
            print(curr.val, end=" -> ")
            curr = curr.next
        print()
        return


class Solution(object):
    def deleteNode(self, node):
        """
        :type node: ListNode
        :rtype: void Do not return anything, modify node in-place instead.
        """
        if node.next != None:
            node.val = node.next.val
            node.next = node.next.next
        return


s = Solution()

print("Solution 1")
head = Node(4)
head.push(5)
head.push(1)
head.push(9)
head.print()
s.deleteNode(head.next)
head.print()
