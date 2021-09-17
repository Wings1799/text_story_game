clothes = ['shirt', 'pants', 'socks', 'shoes', 'hat', 'pants']
#print(clothes.count('pants'))
clothes.append('shorts')

print(clothes)

#for i in enumerate(clothes):
   # print(i)

def change_list(lst, wrd):
    my_list = lst.copy()
    for i in range(len(my_list)):
        if my_list[i]==wrd:
            my_list[i]='CHANGED'
    print(my_list)

#change_list(clothes, 'pants')

def update_list(lst, wrd):
    my_list = lst.copy()
    for i in range(len(my_list)):
        if my_list[i]==wrd:
            print(wrd, 'DELETED')
            del my_list[i]
    print(my_list)

update_list(clothes, 'shorts')