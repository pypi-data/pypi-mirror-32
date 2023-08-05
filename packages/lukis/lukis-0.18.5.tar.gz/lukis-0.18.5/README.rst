=====
Lukis
=====

Description:
~~~~~~~~~~~~
* Generate device independent vector graphic stream
  that can be converted to pdf, svg, or other vector graphic format.

* The output stream is in pure txt, and csv format.

* A simple stream looks like this::

    MoveTo,50,50
    LineTo,150,150
    Stroke

* "Lukis" is the Malay word for "Draw"

* show(GStream) converts the graphic stream to pdf, svg or other
  graphic file, and open it using the default viewer.

Example:
--------
::

    #To make a line from (0,0) to (300,300) and display the graphics:
    S = moveto(0,0)\\
      + lineto(300,300)\\
      + stroke()
    show(S)


Installation
~~~~~~~~~~~~
Python 3 as default Python::

	pip install lukis

Python as Python3::

	pip3 install kopi
