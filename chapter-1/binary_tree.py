# Run this file by copying 'python chapter-1/balanced_binary_tree.py' into the Console box in the bottom righthand corner and hitting 'Enter'.

# Goal: Make a custom hash table implementation in python, using a balanced binary tree
# See hash_table.py for a simpler version

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

  def find(self, key):
    if key == self.key:
      return self.value
    elif key < self.key:
      if self.left:
        return self.left.find(key)
    else:
      if self.right:
        return self.right.find(key)

root = Node(0)

root.insert(10, 'Ten')
root.insert(20, 'Twenty')
root.insert(25, 'Twenty-five')
root.insert(-10, 'Negative')

print(root.find(10))
print(root.find(20))
print(root.find(25))
print(root.find(-10))

printTree(root)