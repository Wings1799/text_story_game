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
        self.myfile = myfile

    def __str__(self):
        print(self.myfile)

def parse_graph(graph):
    graph.split('\n')

    for i in graph:
        if len(i) == 2:
            if i[1] == ('true').capitalize:
                g = Graph(graph)
                g.is_directed = True
            else:
                g = Graph(graph)
                g.is_directed = False
        elif len(i) == 5:
            i.split(' ', 1)
            nodes[i][0]=i[1]
        else:
            i.split(i[0])
            nodes[i][0]=i[1]


    return graph


s = open("input_graph").read()
g = parse_graph(s)
print(g)