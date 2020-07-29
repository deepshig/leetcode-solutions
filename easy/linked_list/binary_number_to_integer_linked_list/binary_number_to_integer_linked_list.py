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
        return


class Solution(object):
    def getDecimalValue(self, head):
        """
        :type head: ListNode
        :rtype: int
        """
        binary_number_str = ""
        while head != None:
            binary_number_str += str(head.val)
            head = head.next

        return int(binary_number_str, 2)


s = Solution()

print("Solution 1")
head = Node(1)
head.push(0)
head.push(1)
head.print()
print("Number : ", s.getDecimalValue(head))

print("Solution 2")
head = Node(0)
head.print()
print("Number : ", s.getDecimalValue(head))

print("Solution 3")
head = Node(1)
head.push(0)
head.push(0)
head.push(1)
head.push(0)
head.push(0)
head.push(1)
head.push(1)
head.push(1)
head.push(0)
head.push(0)
head.push(0)
head.push(0)
head.push(0)
head.push(0)
head.print()
print("Number : ", s.getDecimalValue(head))

print("Solution 4")
head = Node(1)
head.print()
print("Number : ", s.getDecimalValue(head))

print("Solution 5")
head = Node(0)
head.push(0)
head.print()
print("Number : ", s.getDecimalValue(head))
