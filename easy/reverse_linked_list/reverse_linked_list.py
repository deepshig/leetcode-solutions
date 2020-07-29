class Node:
    def __init__(self, val):
        self.val = val
        self.next = None
        return

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


class Solution(object):
    def reverseList(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if head == None or head.next == None:
            return head
        prev, curr = None, head
        while curr.next != None:
            curr.next, prev, curr = prev, curr, curr.next
        curr.next = prev
        return curr


s = Solution()

print("Solution 1")
head = Node(1)
head.push(2)
head.push(3)
head.push(4)
head.push(5)
head.print()
reverse = s.reverseList(head)
reverse.print()
