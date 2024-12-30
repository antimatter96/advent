import sys

from networkx.algorithms import clique
import networkx as nx

edges = []
for line in sys.stdin:
    line = line.strip()
    edges += [(line.split("-")[0], line.split("-")[1])]

G = nx.Graph()
G.add_edges_from(edges)

cliques = clique.find_cliques(G)

max_v = 0
max_g = []
for index, clq in enumerate(cliques):
    if len(clq) > max_v:
        max_v = len(clq)
        max_g = ",".join(sorted(clq))

print(max_g)
