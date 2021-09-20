from Graph import *

class Graph:
    """A simple graph representation allowing node and edge attributes."""

    class Node:
        pass

    nodes = {}

    class Neighbor:
        pass

    adj_list = {}

    def __init__(self, myfile):
        self.myfile = myfile.split

    def __str__(self):
        print(self.myfile)

def parse_graph(graph):
    graph.myfile.split('\n')

    for i in graph.myfile:
        if len(i) == 2:
            if i[1] == ('true').capitalize:
                Graph.is_directed = True
            else:
                Graph.is_directed = False
        elif len(i) == 5:
            i.split(i[0])
            nodes[i][0]=i[1]
        else:
            i.split(i[0])
            nodes[i][0]=i[1]


    return graph


s = open("input_graph").read()
g = parse_graph(s)
print(g)