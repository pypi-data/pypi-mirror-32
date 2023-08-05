==========
WatchGhost
==========

Your invisible but loud monitoring pet

Quickstart
==========

.. code-block:: shell

  pew new watchghost
  pip install -U git+https://gitlab.com/localg-host/watchghost.git
  watchghost-server.py
  $NAVIGATOR http://localhost:8888
  $EDITOR ~/.config/watchghost/*

Code on watchghost
==================

.. code-block:: shell

  pew new watchghost -p python3
  git clone https://gitlab.com/localg-host/watchghost.git
  cd watchghost
  pip install -e .
  ./bin/watchghost-server.py
