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
    def deleteDuplicates(self, head):
        """
        :type head: ListNode
        :rtype: ListNode
        """
        if head == None or head.next == None:
            return head

        prev, curr = head, head.next
        while curr != None:
            if prev.val == curr.val:
                prev.next = curr.next
            else:
                prev = curr
            curr = curr.next
        return head


s = Solution()

print("Solution 1")
head = Node(1)
head.push(1)
head.push(2)
head.print()
noDuplicates = s.deleteDuplicates(head)
noDuplicates.print()

print("Solution 2")
head = Node(1)
head.push(1)
head.push(2)
head.push(3)
head.push(3)
head.print()
noDuplicates = s.deleteDuplicates(head)
noDuplicates.print()
