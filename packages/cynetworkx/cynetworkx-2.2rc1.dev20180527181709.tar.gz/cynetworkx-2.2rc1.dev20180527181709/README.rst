cyNetworkX
========

cyNetworkX is a Cython port of the NetworkX package for the creation, manipulation, and
study of the structure, dynamics, and functions of complex networks.

- **Discussion:** https://github.com/pattern-inc/cynetworkx/issues
- **Bug reports:** https://github.com/pattern-inc/cynetworkx/issues
- **Source:** https://github.com/pattern-inc/cynetworkx

Install
-------

Install the latest version of cyNetworkX::

    $ pip install cynetworkx

Install with all optional dependencies::

    $ pip install cynetworkx[all]

For additional details, please see `INSTALL.rst`.

Simple example
--------------

Find the shortest path between two nodes in an undirected graph::

    >>> import cynetworkx as nx
    >>> G = nx.Graph()
    >>> G.add_edge('A', 'B', weight=4)
    >>> G.add_edge('B', 'D', weight=2)
    >>> G.add_edge('A', 'C', weight=3)
    >>> G.add_edge('C', 'D', weight=4)
    >>> nx.shortest_path(G, 'A', 'D', weight='weight')
    ['A', 'B', 'D']

Bugs
----

Please report any bugs that you find `here <https://github.com/pattern-inc/cynetworkx/issues>`_.
Or, even better, fork the repository on `GitHub <https://github.com/pattern-inc/cynetworkx>`_
and create a pull request (PR). We welcome all changes, big or small, and we
will help you make the PR if you are new to `git` (just ask on the issue and/or
see `CONTRIBUTING.rst`).

License
-------

Released under the Apache 2.0 license (see `LICENSE`)
