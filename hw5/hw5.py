import random

g = random.randrange(0,21)
print(g)

food = ['banana', 'apple', 'egg', 'milk', 'butter']
random.shuffle(food)
print(food)
print(random.sample(food, 2))

for i in food:
    print(f'Next on the shopping list is {i}.')

f = open('hello.txt', 'r+')
print(f.write('HIIII\n'))
print(f.read(20))
f.close()

while True:
    try:
        22+'22'
    except TypeError:
        print("ERRORRRRRR")
        break