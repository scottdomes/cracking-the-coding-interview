# Run this file by copying 'python chapter-1/binary_tree.py' into the Console box in the bottom righthand corner and hitting 'Enter'.

# Goal: Make a custom hash table implementation in python, using a binary tree (not balanced, for simplicity)
# See hash_table.py for a basic hash_table

def printTree(node, level=0):
    if node != None:
        printTree(node.left, level + 1)
        print(' ' * 4 * level + '->', node.key, ':', node.value)
        printTree(node.right, level + 1)

class Node:
  def __init__(self, key, value=None):
    self.left = None
    self.right = None
    self.key = key
    self.value = value

  def insert(self, key, value):
    if key < self.key:
      if self.left:
        self.left.insert(key, value)
      else:
        self.left = Node(key, value)
    else:
      if self.right:
        self.right.insert(key, value)
      else:
        self.right = Node(key, value)
  
  def search(self, key):
    if key == self.key:
      return self. value
    elif key < self.key:
      if self.left:
        return self.left.search(key)
    else:
      if self.right:
        return self.right.search(key)

root = Node(0)

root.insert(10, 'Ten')
root.insert(10, 'Tennn')
root.insert(20, 'Twenty')
root.insert(-10, 'Negative')

print(root.search(20))

# printTree(root)