class Node(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next
        return

    def push(self, val):
        new_node = Node(val)
        head = self
        while head.next != None:
            head = head.next
        head.next = new_node
        return

    def print(self):
        head = self
        while head != None:
            print(head.val, " -> ", end=' ')
            head = head.next
        print()
        return


class Solution(object):
    def removeElements(self, head, val):
        """
        :type head: ListNode
        :type val: int
        :rtype: ListNode
        """
        if head == None or (head.next == None and head.val == val):
            return None
        if head.next == None and head.val != val:
            return head

        prev, curr = head, head.next
        while curr != None:
            if curr.val == val:
                prev.next = curr.next
            else:
                prev = curr
            curr = curr.next

        if head.val == val:
            head = head.next

        return head


s = Solution()

print("Solution 1")
head = Node(1)
head.push(2)
head.push(6)
head.push(3)
head.push(4)
head.push(5)
head.push(6)
head.print()
head = s.removeElements(head, 6)
if head != None:
    head.print()

print("Solution 2")
head = Node(1)
head.print()
head = s.removeElements(head, 2)
if head != None:
    head.print()

print("Solution 3")
head = Node(1)
head.print()
head = s.removeElements(head, 1)
if head != None:
    head.print()

print("Solution 4")
head = Node(1)
head.push(1)
head.print()
head = s.removeElements(head, 1)
if head != None:
    head.print()
