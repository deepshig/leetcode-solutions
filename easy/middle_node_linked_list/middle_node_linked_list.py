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
    def middleNode(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        slow, fast = head, head
        while slow != None and fast != None and fast.next != None:
            slow = slow.next
            fast = fast.next.next
        return slow


s = Solution()

print("Solution 1")
head = Node(1)
head.push(2)
head.push(3)
head.push(4)
head.push(5)
head.print()
mid = s.middleNode(head)
mid.print()

print("Solution 2")
head = Node(1)
head.push(2)
head.push(3)
head.push(4)
head.push(5)
head.push(6)
head.print()
mid = s.middleNode(head)
mid.print()
