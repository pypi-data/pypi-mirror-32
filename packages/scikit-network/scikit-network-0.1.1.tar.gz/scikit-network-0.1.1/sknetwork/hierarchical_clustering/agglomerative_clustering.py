# -*- coding: utf-8 -*-
# agglomerative_clustering.py - functions for computing hierarchy using agglomerative approach
#
# Copyright 2018 Scikit-network Developers.
# Copyright 2018 Bertrand Charpentier <bertrand.charpentier@live.fr>
#
# This file is part of Scikit-network.
#
# NetworkX is distributed under a BSD license; see LICENSE.txt for more information.
"""
Agglomerative clustering algorithms
"""

import numpy as np
import networkx as nx


_AFFINITY = {'unitary', 'weighted'}
_LINKAGE = {'single', 'average', 'complete', 'modular'}


def linkage_clustering(graph, affinity='weighted', linkage='modular', f=lambda l: - np.log(l), check=True):
    """Compute a hierarchy of an undirected graph given a linkage.
    The graph can be weighted or unweighted

    Parameters
    ----------
    graph : NetworkX graph
        An undirected graph.
    affinity, optional : string (default: 'weighted')
        The affinity  can be either 'weighted' or 'unitary'.
        Value 'weighted' takes the attribute 'weight' on the edges.
        Value 'unitary' consider that the edges have a weight equal to 1
    linkage : string, optional (default: 'modular')
        The parameter linkage can be 'single', 'average', 'complete' or 'modular'.

        =============== ========================================
        Value           Linkage
        =============== ========================================
        'single'        max wij
        'average'       min wij
        'complete'      n^2/w*wij/(|i||j|)
        'modular'       w wij/(wi wj)
        =============== ========================================

    f : function, optional (default: lambda l: -np.log(l))
        The function f define the transformation to perform on the linkage
        to get a distance.
    check : bool, optional (default: True)
        If True, reorder the node labels and check that the edges have the
        'weight' attribute corresponding to the given affinity.

    Returns
    -------
    dendrogram : numpy array
        dendrogram.

    Raises
    ------
    ValueError
        If the affinity or the linkage is not known.
    KeyError
        If all the edges do not have the 'weight' attribute with the 'weighted' affinity.

    Notes
    -----
    Modular linkage is awesome !

    See Also
    --------
    laplacian_matrix
    """

    if affinity not in _AFFINITY:
        raise ValueError("Unknown affinity type %s."
                         "Valid options are %s" % (affinity, _AFFINITY))

    if linkage not in _LINKAGE:
        raise ValueError("Unknown linkage type %s."
                         "Valid options are %s" % (linkage, _LINKAGE))

    graph_copy = graph.copy()

    if check:

        graph_copy = nx.convert_node_labels_to_integers(graph_copy)

        if affinity == 'unitary':
            for e in graph_copy.edges:
                graph_copy.add_edge(e[0], e[1], weight=1)

        n_edges = len(list(graph_copy.edges()))
        n_weighted_edges = len(nx.get_edge_attributes(graph_copy, 'weight'))
        if affinity == 'weighted' and not n_weighted_edges == n_edges:
            raise KeyError("%s edges among %s do not have the attribute/key \'weigth\'."
                           % (n_edges - n_weighted_edges, n_edges))

    if linkage == 'single':
        dendrogram = single_linkage_hierarchy(graph_copy, f)
    elif linkage == 'average':
        dendrogram = average_linkage_hierarchy(graph_copy, f)
    elif linkage == 'complete':
        dendrogram = complete_linkage_hierarchy(graph_copy, f)
    elif linkage == 'modular':
        dendrogram = modular_linkage_hierarchy(graph_copy, f)

    return reorder_dendrogram(dendrogram)


def single_linkage_hierarchy(graph, f):
    remaining_nodes = set(graph.nodes())
    n_nodes = len(remaining_nodes)

    cluster_size = {u: 1 for u in range(n_nodes)}
    connected_components = []
    dendrogram = []
    u = n_nodes

    while n_nodes > 0:
        for new_node in remaining_nodes:
            chain = [new_node]
            break
        while chain:
            a = chain.pop()
            linkage_max = - float("inf")
            b = -1
            neighbors_a = list(graph.neighbors(a))
            for v in neighbors_a:
                if v != a:
                    linkage = float(graph[a][v]['weight'])
                    if linkage > linkage_max:
                        b = v
                        linkage_max = linkage
                    elif linkage == linkage_max:
                        b = min(b, v)
            linkage = linkage_max
            if chain:
                c = chain.pop()
                if b == c:
                    dendrogram.append([a, b, f(linkage), cluster_size[a] + cluster_size[b]])
                    graph.add_node(u)
                    remaining_nodes.add(u)
                    neighbors_a = list(graph.neighbors(a))
                    neighbors_b = list(graph.neighbors(b))
                    for v in neighbors_a:
                        graph.add_edge(u, v, weight=graph[a][v]['weight'])
                    for v in neighbors_b:
                        if graph.has_edge(u, v):
                            graph[u][v]['weight'] = max(graph[b][v]['weight'], graph[u][v]['weight'])
                        else:
                            graph.add_edge(u, v, weight=graph[b][v]['weight'])
                    graph.remove_node(a)
                    remaining_nodes.remove(a)
                    graph.remove_node(b)
                    remaining_nodes.remove(b)
                    n_nodes -= 1
                    cluster_size[u] = cluster_size.pop(a) + cluster_size.pop(b)
                    u += 1
                else:
                    chain.append(c)
                    chain.append(a)
                    chain.append(b)
            elif b >= 0:
                chain.append(a)
                chain.append(b)
            else:
                connected_components.append((a, cluster_size[a]))
                graph.remove_node(a)
                cluster_size.pop(a)
                n_nodes -= 1

    a, cluster_size = connected_components.pop()
    for b, t in connected_components:
        cluster_size += t
        dendrogram.append([a, b, float("inf"), cluster_size])
        a = u
        u += 1

    return np.array(dendrogram)


def average_linkage_hierarchy(graph, f):
    remaining_nodes = set(graph.nodes())
    n_nodes = len(remaining_nodes)

    wtot = 0
    for (u, v) in graph.edges():
        weight = graph[u][v]['weight']
        wtot += 2 * weight
    n_nodes2_wtot = n_nodes * n_nodes / wtot
    cluster_size = {u: 1 for u in range(n_nodes)}
    connected_components = []
    dendrogram = []
    u = n_nodes

    while n_nodes > 0:
        for new_node in remaining_nodes:
            chain = [new_node]
            break
        while chain:
            a = chain.pop()
            linkage_max = - float("inf")
            b = -1
            neighbors_a = list(graph.neighbors(a))
            for v in neighbors_a:
                if v != a:
                    linkage = n_nodes2_wtot * float(graph[a][v]['weight'])/(cluster_size[a]*cluster_size[v])
                    if linkage > linkage_max:
                        b = v
                        linkage_max = linkage
                    elif linkage == linkage_max:
                        b = min(b, v)
            linkage = linkage_max
            if chain:
                c = chain.pop()
                if b == c:
                    dendrogram.append([a, b, f(linkage), cluster_size[a] + cluster_size[b]])
                    graph.add_node(u)
                    remaining_nodes.add(u)
                    neighbors_a = list(graph.neighbors(a))
                    neighbors_b = list(graph.neighbors(b))
                    for v in neighbors_a:
                        graph.add_edge(u, v, weight=graph[a][v]['weight'])
                    for v in neighbors_b:
                        if graph.has_edge(u, v):
                            graph[u][v]['weight'] += graph[b][v]['weight']
                        else:
                            graph.add_edge(u, v, weight=graph[b][v]['weight'])
                    graph.remove_node(a)
                    remaining_nodes.remove(a)
                    graph.remove_node(b)
                    remaining_nodes.remove(b)
                    n_nodes -= 1
                    cluster_size[u] = cluster_size.pop(a) + cluster_size.pop(b)
                    u += 1
                else:
                    chain.append(c)
                    chain.append(a)
                    chain.append(b)
            elif b >= 0:
                chain.append(a)
                chain.append(b)
            else:
                connected_components.append((a, cluster_size[a]))
                graph.remove_node(a)
                cluster_size.pop(a)
                n_nodes -= 1

    a, cluster_size = connected_components.pop()
    for b, t in connected_components:
        cluster_size += t
        dendrogram.append([a, b, float("inf"), cluster_size])
        a = u
        u += 1

    return np.array(dendrogram)


def complete_linkage_hierarchy(graph, f):
    remaining_nodes = set(graph.nodes())
    n_nodes = len(remaining_nodes)

    cluster_size = {u: 1 for u in range(n_nodes)}
    connected_components = []
    dendrogram = []
    u = n_nodes

    while n_nodes > 0:
        for new_node in remaining_nodes:
            chain = [new_node]
            break
        while chain:
            a = chain.pop()
            linkage_max = - float("inf")
            b = -1
            neighbors_a = list(graph.neighbors(a))
            for v in neighbors_a:
                if v != a:
                    linkage = float(graph[a][v]['weight'])
                    if linkage > linkage_max:
                        b = v
                        linkage_max = linkage
                    elif linkage == linkage_max:
                        b = min(b, v)
            linkage = linkage_max
            if chain:
                c = chain.pop()
                if b == c:
                    dendrogram.append([a, b, f(linkage), cluster_size[a] + cluster_size[b]])
                    graph.add_node(u)
                    remaining_nodes.add(u)
                    neighbors_a = list(graph.neighbors(a))
                    neighbors_b = list(graph.neighbors(b))
                    for v in neighbors_a:
                        graph.add_edge(u, v, weight=graph[a][v]['weight'])
                    for v in neighbors_b:
                        if graph.has_edge(u, v):
                            graph[u][v]['weight'] = min(graph[b][v]['weight'], graph[u][v]['weight'])
                        else:
                            graph.add_edge(u, v, weight=graph[b][v]['weight'])
                    graph.remove_node(a)
                    remaining_nodes.remove(a)
                    graph.remove_node(b)
                    remaining_nodes.remove(b)
                    n_nodes -= 1
                    cluster_size[u] = cluster_size.pop(a) + cluster_size.pop(b)
                    u += 1
                else:
                    chain.append(c)
                    chain.append(a)
                    chain.append(b)
            elif b >= 0:
                chain.append(a)
                chain.append(b)
            else:
                connected_components.append((a, cluster_size[a]))
                graph.remove_node(a)
                cluster_size.pop(a)
                n_nodes -= 1

    a, cluster_size = connected_components.pop()
    for b, t in connected_components:
        cluster_size += t
        dendrogram.append([a, b, float("inf"), cluster_size])
        a = u
        u += 1

    return np.array(dendrogram)


def modular_linkage_hierarchy(graph, f):
    remaining_nodes = set(graph.nodes())
    n_nodes = len(remaining_nodes)

    w = {u: 0 for u in range(n_nodes)}
    wtot = 0
    for (u, v) in graph.edges():
        weight = graph[u][v]['weight']
        w[u] += weight
        w[v] += weight
        wtot += 2 * weight
    cluster_size = {u: 1 for u in range(n_nodes)}
    connected_components = []
    dendrogram = []
    u = n_nodes

    while n_nodes > 0:
        for new_node in remaining_nodes:
            chain = [new_node]
            break
        while chain:
            a = chain.pop()
            linkage_max = - float("inf")
            b = -1
            neighbors_a = list(graph.neighbors(a))
            for v in neighbors_a:
                if v != a:
                    linkage = wtot * float(graph[a][v]['weight'])/(w[a]*w[v])
                    if linkage > linkage_max:
                        b = v
                        linkage_max = linkage
                    elif linkage == linkage_max:
                        b = min(b, v)
            linkage = linkage_max
            if chain:
                c = chain.pop()
                if b == c:
                    dendrogram.append([a, b, f(linkage), cluster_size[a] + cluster_size[b]])
                    graph.add_node(u)
                    remaining_nodes.add(u)
                    neighbors_a = list(graph.neighbors(a))
                    neighbors_b = list(graph.neighbors(b))
                    for v in neighbors_a:
                        graph.add_edge(u, v, weight=graph[a][v]['weight'])
                    for v in neighbors_b:
                        if graph.has_edge(u, v):
                            graph[u][v]['weight'] += graph[b][v]['weight']
                        else:
                            graph.add_edge(u, v, weight=graph[b][v]['weight'])
                    graph.remove_node(a)
                    remaining_nodes.remove(a)
                    graph.remove_node(b)
                    remaining_nodes.remove(b)
                    n_nodes -= 1
                    w[u] = w.pop(a) + w.pop(b)
                    cluster_size[u] = cluster_size.pop(a) + cluster_size.pop(b)
                    u += 1
                else:
                    chain.append(c)
                    chain.append(a)
                    chain.append(b)
            elif b >= 0:
                chain.append(a)
                chain.append(b)
            else:
                connected_components.append((a, cluster_size[a]))
                graph.remove_node(a)
                w.pop(a)
                cluster_size.pop(a)
                n_nodes -= 1

    a, cluster_size = connected_components.pop()
    for b, t in connected_components:
        cluster_size += t
        dendrogram.append([a, b, float("inf"), cluster_size])
        a = u
        u += 1

    return np.array(dendrogram)


def reorder_dendrogram(D):
    n = np.shape(D)[0] + 1
    order = np.zeros((2, n - 1), float)
    order[0] = range(n - 1)
    order[1] = np.array(D)[:, 2]
    index = np.lexsort(order)
    nindex = {i: i for i in range(n)}
    nindex.update({n + index[t]: n + t for t in range(n - 1)})
    return np.array([[nindex[int(D[t][0])], nindex[int(D[t][1])], D[t][2], D[t][3]] for t in range(n - 1)])[index, :]
