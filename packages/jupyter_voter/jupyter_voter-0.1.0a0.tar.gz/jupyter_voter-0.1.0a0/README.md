jupyter-voter
===============================

A Jupyter widget for voting on simple ML training data.

Installation
------------

To install use pip:

    $ pip install jupyter_voter
    $ jupyter nbextension enable --py --sys-prefix jupyter_voter


For a development installation (requires npm),

    $ git clone https://github.com/GeoBigData/jupyter-voter.git
    $ cd jupyter-voter
    $ pip install -e .
    $ jupyter nbextension install --py --symlink --sys-prefix jupyter_voter
    $ jupyter nbextension enable --py --sys-prefix jupyter_voter
