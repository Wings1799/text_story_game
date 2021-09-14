"""
id = 123456
my_id = int(input("Enter ID: "))
if my_id == id:
    print("ACCESS GRANTED")
elif len(str(my_id)) != len(str(id)):
    print("INVALID ID FORMAT")
else:
    print("ACCESS DENIED")
"""
"""
for x in range(0,5):
    print(x)
"""
"""
names = ['Jack', 'John', 'James', 'Jeff', 'Jarvis']
for n in names:
    print(n, 'PRESENT')
"""
"""
for i in range(-2, 15, 4):
    print(i)
"""
"""
words = ['Eggs', 'Milk', 'Lettuce', 'Bananas']
#for i in enumerate(words):
    #print(i)
item = str(input('Enter item: '))
if item in words:
    for i in range(len(words)):
        if item == words[i]:
            words[i] = ''
else:
    words.append(item)
print(words)
"""
"""
def cubic(n):
    return n**3
print(cubic(3))
"""
"""
import math
def pythag(x,y):
    sum = x**2 + y**2
    return math.sqrt(sum)
print(pythag(3,4))
"""