class Graph:
    """A simple graph representation allowing node and edge attributes."""
    is_directed = ...
    adj_list = {}
    nodes = {}

    def __init__(self, graph):
        self.strung = graph
        self.graph = graph.split('\n')
        for i in self.graph:
            j = i.split()
            if len(j)==2:
                if j[1]=='true':
                    self.is_directed = True
                else:
                    self.is_directed = False
            if len(j)==5:
                n = Graph.Node(j[0],j[2],j[4])
                self.nodes[j[0]]=n
            if len(j)==3:
                ne = Graph.Neighbor(j[0], j[2])
                self.adj_list[j[0]]=ne

    def num_edges(self):
        return len(self.adj_list)

    def has_edge(self, n1, n2):
        point = n1+n2
        for i in self.adj_list:
            if i == point:
                return True
            else:
                return False

    def clear_edges(self):
        self.adj_list.clear()

    def __str__(self):
        return self.strung
    
    class Node:
        def __init__(self, name, x, y):
            self.name = name
            self.x = x
            self.y = y


    

    class Neighbor:
        def __init__(self, name, value):
            self.name = name
            self.value = value

def parse_graph(graph):
    g = Graph(graph)
    return g
