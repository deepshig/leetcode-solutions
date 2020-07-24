class Stack(object):
    def __init__(self):
        self.elements = []
        return

    def push(self, x):
        self.elements.append(x)
        return

    def pop(self):
        if self.empty():
            return 0
        return self.elements.pop()

    def peek(self):
        if self.empty():
            return 0
        return self.elements[len(self.elements)-1]

    def empty(self):
        return len(self.elements) == 0


class MyQueue(object):

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.stack = Stack()
        return

    def push(self, x):
        """
        Push element x to the back of queue.
        :type x: int
        :rtype: None
        """
        temp = Stack()
        while not self.stack.empty():
            element = self.stack.pop()
            temp.push(element)

        self.stack.push(x)
        while not temp.empty():
            element = temp.pop()
            self.stack.push(element)
        return

    def pop(self):
        """
        Removes the element from in front of queue and returns that element.
        :rtype: int
        """
        return self.stack.pop()

    def peek(self):
        """
        Get the front element.
        :rtype: int
        """
        return self.stack.peek()

    def empty(self):
        """
        Returns whether the queue is empty.
        :rtype: bool
        """
        return self.stack.empty()


# Your MyQueue object will be instantiated and called as such:
obj = MyQueue()
obj.push(1)
obj.push(2)
a = obj.pop()
b = obj.peek()
c = obj.empty()

print("a = ", a, ", b = ", b, ", c = ", c)
