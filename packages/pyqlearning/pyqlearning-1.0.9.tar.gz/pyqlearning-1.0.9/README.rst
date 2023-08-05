Reinforcement Learning Library: pyqlearning
===========================================

``pyqlearning`` is Python library to implement Reinforcement Learning,
especially for Q-Learning.

Description
-----------

Considering many variable parts and functional extensions in the
Q-learning paradigm, I implemented these Python Scripts for
demonstrations of *commonality/variability analysis* in order to design
the models.

Documentation
-------------

Full documentation is available on
https://code.accel-brain.com/Reinforcement-Learning/ . This document
contains information on functionally reusability, functional scalability
and functional extensibility.

Installation
------------

Install using pip:

.. code:: sh

    pip install pyqlearning

Source code
~~~~~~~~~~~

The source code is currently hosted on GitHub.

-  `accel-brain-code/Reinforcement-Learning <https://github.com/chimera0/accel-brain-code/tree/master/Reinforcement-Learning>`__

Python package index(PyPI)
~~~~~~~~~~~~~~~~~~~~~~~~~~

Installers for the latest released version are available at the Python
package index.

-  `pyqlearning : Python Package
   Index <https://pypi.python.org/pypi/pyqlearning/>`__

Dependencies
~~~~~~~~~~~~

-  numpy: v1.13.3 or higher.
-  pandas: v0.22.0 or higher.

Demonstration: Simple Maze Solving by Q-Learning (Jupyter notebook)
-------------------------------------------------------------------

I have details of this library on my Jupyter notebook:
`search\_maze\_by\_q\_learning.ipynb <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/search_maze_by_q_learning.ipynb>`__.
This notebook demonstrates a simple maze solving algorithm based on
Epsilon-Greedy Q-Learning or Q-Learning, loosely coupled with Deep
Boltzmann Machine(DBM).

Demonstration: Q-Learning
-------------------------

Q-Learning is a kind of
``Temporal Difference learning``\ (``TD Learning``) that can be
considered as hybrid of ``Monte Carlo method`` and
``Dynamic Programming Method``. As ``Monte Carlo method``,
``TD Learning`` algorithm can learn by experience without model of
environment. And this learning algorithm is *functionally equivalent* of
bootstrap method as ``Dynamic Programming Method``.

``Epsilon Greedy Q-Leanring`` algorithm is ``off-policy``. In this
paradigm, *stochastic* searching and *deterministic* searching can
coexist by hyperparameter ε (0 < ε < 1) that is probability that agent
searches greedy. Greedy searching is *deterministic* in the sense that
policy of agent follows the selection that maximizes the Q-Value.

`demo\_maze\_greedy\_q\_learning.py <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/demo_maze_greedy_q_learning.py>`__
is a simple maze solving algorithm. ``MazeGreedyQLearning`` in
 `devsample/maze\_greedy\_q\_learning.py <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/devsample/maze_greedy_q_learning.py>`__
is a ``Concrete Class`` in ``Template Method Pattern`` to run the
Q-Learning algorithm for this task. ``GreedyQLearning`` in
`pyqlearning/qlearning/greedy\_q\_learning.py <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/pyqlearning/qlearning/greedy_q_learning.py>`__
is also ``Concreat Class`` for the epsilon-greedy-method. The
``Abstract Class`` that defines the skeleton of Q-Learning algorithm in
the operation and declares algorithm placeholders is
`pyqlearning/q\_learning.py <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/pyqlearning/q_learning.py>`__.
So
`demo\_maze\_greedy\_q\_learning.py <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/demo_maze_greedy_q_learning.py>`__
is a kind of ``Client`` in ``Template Method Pattern``.

This algorithm allow the *agent* to search the goal in maze by *reward
value* in each point in map.

The following is an example of map.

::

    [['#' '#' '#' '#' '#' '#' '#' '#' '#' '#']
     ['#' 'S'  4   8   8   4   9   6   0  '#']
     ['#'  2  26   2   5   9   0   6   6  '#']
     ['#'  2  '@' 38   5   8   8   1   2  '#']
     ['#'  3   6   0  49   8   3   4   9  '#']
     ['#'  9   7   4   6  55   7   0   3  '#']
     ['#'  1   8   4   8   2  69   8   2  '#']
     ['#'  1   0   2   1   7   0  76   2  '#']
     ['#'  2   8   0   1   4   7   5  'G' '#']
     ['#' '#' '#' '#' '#' '#' '#' '#' '#' '#']]

-  ``#`` is wall in maze.
-  ``S`` is a start point.
-  ``G`` is a goal.
-  ``@`` is the agent.

In relation to reinforcement learning theory, the *state* of *agent* is
2D position coordinates and the *action* is to dicide the direction of
movement. Within the wall, the *agent* is movable in a cross direction
and can advance by one point at a time. After moving into a new
position, the *agent* can obtain a *reward*. On greedy searching, this
extrinsically motivated agent performs in order to obtain some *reward*
as high as possible. Each *reward value* is plot in map.

To see how *agent* can search and rearch the goal, run the batch
program:
`demo\_maze\_greedy\_q\_learning.py <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/demo_maze_greedy_q_learning.py>`__

.. code:: bash

    python demo_maze_greedy_q_learning.py

Demonstration: Q-Learning, loosely coupled with Deep Boltzmann Machine.
-----------------------------------------------------------------------

`demo\_maze\_deep\_boltzmann\_q\_learning.py <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/demo_maze_deep_boltzmann_q_learning.py>`__
is a demonstration of how the *Q-Learning* can be to *deepen*. A
so-called *Deep Q-Network* (DQN) is meant only as an example. In this
demonstration, let me cite the *Q-Learning* , loosely coupled with
**Deep Boltzmann Machine** (DBM). As API Documentation of
`pydbm <https://github.com/chimera0/accel-brain-code/tree/master/Deep-Learning-by-means-of-Design-Pattern>`__
library has pointed out, DBM is functionally equivalent to stacked
auto-encoder. The main function I observe is the same as dimensions
reduction(or pre-training). Then the function this DBM is dimensionality
reduction of *reward value* matrix.

Q-Learning, loosely coupled with Deep Boltzmann Machine (DBM), is a more
effective way to solve maze. The pre-training by DBM allow Q-Learning
*agent* to abstract feature of ``reward value`` matrix and to observe
the map in a bird's-eye view. Then *agent* can reache the goal with a
smaller number of trials.

To realize the power of DBM, I performed a simple experiment.

Feature engineering
~~~~~~~~~~~~~~~~~~~

For instance, a feature in each coordinate can be transformed and
extracted by reward value as so-called *observed data points* in its
adjoining points. More formally, see
`search\_maze\_by\_q\_learning.ipynb <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/search_maze_by_q_learning.ipynb>`__.

Then the feature representation can be as calculated. After this
pre-training, the DBM has extracted *feature points* below.

::

    [['#' '#' '#' '#' '#' '#' '#' '#' '#' '#']
     ['#' 'S' 0.22186305563593528 0.22170599483791015 0.2216928599218454
      0.22164807496640074 0.22170371283788584 0.22164021608623224
      0.2218165339471332 '#']
     ['#' 0.22174745260072407 0.221880094307873 0.22174244728061343
      0.2214709292493749 0.22174626768015263 0.2216756589222596
      0.22181057818975275 0.22174525714311788 '#']
     ['#' 0.22177496678085065 0.2219122743656551 0.22187543599733664
      0.22170745588799798 0.2215226084843615 0.22153827385193636
      0.22168466277729898 0.22179391402965035 '#']
     ['#' 0.2215341770250964 0.22174315536140118 0.22143149966676515
      0.22181685688674144 0.22178215385805333 0.2212249704384472
      0.22149210148879617 0.22185413678274837 '#']
     ['#' 0.22162363223483128 0.22171313373253035 0.2217109987501002
      0.22152432841656014 0.22175562457887335 0.22176040052504634
      0.22137688854285298 0.22175365642579478 '#']
     ['#' 0.22149515807715153 0.22169199881701832 0.22169558478042856
      0.2216904005450013 0.22145368271014734 0.2217144069625017
      0.2214896100292738 0.221398594191006 '#']
     ['#' 0.22139837944992058 0.22130176116356184 0.2215414328019404
      0.22146667964656613 0.22164354506366127 0.22148685616333666
      0.22162822887193126 0.22140174437162474 '#']
     ['#' 0.22140060918518528 0.22155145714201702 0.22162929776464463
      0.22147466752374162 0.22150300682310872 0.22162775291471243
      0.2214233075299188 'G' '#']
     ['#' '#' '#' '#' '#' '#' '#' '#' '#' '#']]

To see how *agent* can search and rearch the goal, install
`pydbm <https://github.com/chimera0/accel-brain-code/tree/master/Deep-Learning-by-means-of-Design-Pattern>`__
library and run the batch program:
`demo\_maze\_deep\_boltzmann\_q\_learning.py <https://github.com/chimera0/accel-brain-code/blob/master/Reinforcement-Learning/demo_maze_deep_boltzmann_q_learning.py>`__

.. code:: bash

    python demo_maze_deep_boltzmann_q_learning.py

More detail demos
~~~~~~~~~~~~~~~~~

-  `Webクローラ型人工知能：キメラ・ネットワークの仕様 <https://media.accel-brain.com/_chimera-network-is-web-crawling-ai/>`__

   -  20001 bots are running as 20001 web-crawlers and 20001
      web-scrapers.

Related PoC
~~~~~~~~~~~

-  `Webクローラ型人工知能によるパラドックス探索暴露機能の社会進化論 <https://accel-brain.com/social-evolution-of-exploration-and-exposure-of-paradox-by-web-crawling-type-artificial-intelligence/>`__
   (Japanese)

   -  `プロトタイプの開発：人工知能エージェント「キメラ・ネットワーク」 <https://accel-brain.com/social-evolution-of-exploration-and-exposure-of-paradox-by-web-crawling-type-artificial-intelligence/5/#i-8>`__

-  `深層強化学習のベイズ主義的な情報探索に駆動された自然言語処理の意味論 <https://accel-brain.com/semantics-of-natural-language-processing-driven-by-bayesian-information-search-by-deep-reinforcement-learning/>`__
   (Japanese)

   -  `プロトタイプの開発：深層学習と強化学習による「排除された第三項」の推論 <https://accel-brain.com/semantics-of-natural-language-processing-driven-by-bayesian-information-search-by-deep-reinforcement-learning/4/#i-5>`__

-  `ハッカー倫理に準拠した人工知能のアーキテクチャ設計 <https://accel-brain.com/architectural-design-of-artificial-intelligence-conforming-to-hacker-ethics/>`__
   (Japanese)

   -  `プロトタイプの開発：深層強化学習のアーキテクチャ設計 <https://accel-brain.com/architectural-design-of-artificial-intelligence-conforming-to-hacker-ethics/5/#i-2>`__

-  `ヴァーチャルリアリティにおける動物的「身体」の物神崇拝的なユースケース <https://accel-brain.com/cyborg-fetischismus-in-sammlung-von-animalisch-korper-in-virtual-reality/>`__
   (Japanese)

   -  `プロトタイプの開発：「人工天使ヒューズ＝ヒストリア」 <https://accel-brain.com/cyborg-fetischismus-in-sammlung-von-animalisch-korper-in-virtual-reality/4/#i-6>`__

Author
------

-  chimera0(RUM)

Author URI
----------

-  http://accel-brain.com/

License
-------

-  GNU General Public License v2.0
