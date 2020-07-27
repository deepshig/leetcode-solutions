# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None

    def push(self, x):
        new_node = ListNode(x)
        self.next = new_node
        return new_node

    def print(self):
        node = self
        while node != None:
            print(node.val)
            node = node.next


def array_to_linked_list(arr, cycle_start_pos):
    if len(arr) < 1:
        return None

    head = ListNode(arr[0])
    start, tail = head, head
    for val in arr[1:]:
        new_node = ListNode(val)
        tail.next = new_node
        tail = new_node

    if cycle_start_pos >= 0:
        for i in range(0, cycle_start_pos):
            start = start.next
        tail.next = start

    return head


class Solution(object):
    def hasCycle(self, head):
        """
        :type head: ListNode
        :rtype: bool
        """
        slow, fast = head, head
        while slow and fast and fast.next:
            slow = slow.next
            fast = fast.next.next
            if slow == fast:
                return True

        return False


s = Solution()
head = array_to_linked_list([3, 2, 0, -4], 1)
print("Solution 1 : ", s.hasCycle(head))

head = array_to_linked_list([1, 2], 0)
print("Solution 2 : ", s.hasCycle(head))

head = array_to_linked_list([1], -1)
print("Solution 3 : ", s.hasCycle(head))

head = array_to_linked_list([], -1)
print("Solution 4 : ", s.hasCycle(head))
