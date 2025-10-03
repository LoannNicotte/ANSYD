import matplotlib.pyplot as plt
import networkx as nx

# Création du graphe dirigé
G = nx.DiGraph()

# Nœuds
nodes = [1, 2, 3, 4, 5]
G.add_nodes_from(nodes)

# Arêtes avec labels (deux canaux par arête)
edges = [
    (1, 3, "c1"), (3, 1, "c1"),
    (2, 3, "c2"), (3, 2, "c2"),
    (3, 5, "c3"), (5, 3, "c3"),
    (1, 4, "c4"), (4, 1, "c4"),
    (4, 5, "c5"), (5, 4, "c5"),
]
G.add_edges_from([(u, v) for u, v, _ in edges])

# Position des nœuds
pos = {
    1: (0, 2),
    2: (3, 2),
    3: (1.5, 1.5),
    4: (0, 0),
    5: (3, 0),
}

plt.figure(figsize=(8, 7))

# Dessiner les nœuds et leurs labels
nx.draw_networkx_nodes(G, pos, node_size=1000, node_color="lightblue")
nx.draw_networkx_labels(G, pos, font_size=14, font_weight="bold")

# Dessiner les arêtes avec flèches (une par sens)
nx.draw_networkx_edges(
    G, pos,
    edgelist=[(u, v) for u, v, _ in edges],
    arrowstyle="-|>", arrowsize=40, edge_color="black",
    connectionstyle="arc3,rad=0.25", width=2
)

# Labels des arêtes
edge_labels = {(u, v): label for u, v, label in edges}
print(edge_labels)
nx.draw_networkx_edge_labels(
    G, pos,
    edge_labels=edge_labels,
    font_size=11,
    rotate=False,
)
plt.title("Graphe avec 5 noeuds", fontsize=14)
plt.axis("off")
plt.show()