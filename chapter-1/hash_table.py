# Run this file by copying 'python chapter-1/hash_table.py' into the Console box in the bottom righthand corner and hitting 'Enter'.

# Goal: Make a custom hash table implementation in python

# Implementation 1: No collision detection

# Creates an empty array of length 10
hash_table = [None] * 10

def hashing_func(key, hash_table):
    return key % len(hash_table)

def insert(hash_table, key, value):
    hash_key = hashing_func(key, hash_table)
    hash_table[hash_key] = value 

insert(hash_table, 10, 'Ten')
insert(hash_table, 20, 'Twenty')
insert(hash_table, 25, 'Twenty-five')

print(hash_table)
# -> ['Twenty', None, None, None, None, 'Twenty-five', None, None, None, None]

# Implementation 2: Chaining collision resolution

hash_table2 = [[] for _ in range(10)]

def insert2(hash_table, key, value):
    hash_key = hashing_func(key, hash_table)
    hash_table[hash_key].append(value)

insert2(hash_table2, 10, 'Ten')
insert2(hash_table2, 20, 'Twenty')
insert2(hash_table2, 25, 'Twenty-five')

print(hash_table2)
# -> [['Ten', 'Twenty'], [], [], [], [], ['Twenty-five'], [], [], [], []]

