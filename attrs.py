from attrs import *
import sys

def has_attr(dict, a):
    if a in dict:
        return True
    else:
        return False

def remove_attr(d, a):
    dict = d.copy()
    if a in dict:
        del dict[a]
    return dict

def update_attr(dict, l):
    d = dict.copy()
    if l[1]==non_attr_val:
        for i in d:
            if l[0]==i:
                del d[i]
                break
    else:
        if l[0] in d:
            for i in d:
                if l[0]==i:
                    d[i]=l[1]
        else:
            d[l[0]]=l[1]
    return d

def update_multi_attr(dict, l):
    d=dict.copy()
    for i in l:
        if i[1]==non_attr_val:
            for k in d:
                if i[0]==k:
                    del d[k]
                    break
        else:
            if i[0] in d:
                for j in d:
                    if j[0]==i[0]:
                        d[i[0]]=j[1]
            else:
                d[i[0]]=i[1]
    return d

def attrs_from_string(word):
    dict={}
    for i in word:
        for n in range(1,len(word)):
            if type(i)==str:
                dict[i]=word[n]
            break
    return dict



non_attr_val = -sys.maxsize

my_list = {'hi':1, 'hello':3, 'hola':8}
#print(has_attr(my_list, 'ho'))
#print(remove_attr(my_list, 'hi'))
#print(update_attr(my_list, ['hi', 9]))
print(my_list)
print(attrs_from_string('a 3 b 4'))