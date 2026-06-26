# 1. Python - стек(LifoQueue)
from queue import LifoQueue

stack = LifoQueue()

stack.put(10)
stack.put(20)
stack.put(30)

top = stack.get()
print(top)          # 30

stack.put(40)

next_top = stack.get()
print(next_top)         # 40