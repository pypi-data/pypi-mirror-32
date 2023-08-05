.. |build| image:: https://img.shields.io/travis/gjeusel/plotlyink/master.svg
    :target: https://travis-ci.org/gjeusel/plotlyink
.. |status| image:: https://img.shields.io/badge/State-Beta-yellow.svg
.. |license| image:: https://img.shields.io/github/license/gjeusel/plotlyink.svg

|build| |status| |license|


Plotlyink
==========
.. _plotly: http://www.plot.ly
.. _pandas: http://pandas.pydata.org/

Plotlyink aims at marrying plotly_ with pandas_ for quick and easy plotting.

Installation
------------

.. code:: bash

    pip install plotlyink


Overview
--------

.. code:: python

    import pandas as pd
    import plotlyink
    df = pd.DataFrame({
        'one' : [1., 2., 3., 4.],
        'two' : [4., 3., 2., 1.],
        })

    # open .html in your brother:
    df.iplot.scatter()

    # get figure:
    fig = df.iplot.scatter(as_figure=True)


For more, see: `Tutorial Notebook <https://github.com/gjeusel/plotlyink/blob/master/notebooks/tutorial.ipynb>`_.
