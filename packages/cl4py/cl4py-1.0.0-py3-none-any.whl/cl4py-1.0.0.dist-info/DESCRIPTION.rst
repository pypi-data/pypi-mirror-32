#+TITLE: cl4py - Common Lisp for Python

The cl4py library allows Python programs to call Common Lisp libraries.

** Motivation
You are a Python programmer, but you want access to some of the powerful
features of Lisp, for example to compile code at run time? Or you want to
use some of the awesome, free [[https://www.quicklisp.org/beta/releases.html][Lisp libraries]]? In that case, cl4py is here
to help you.

** Tutorial

You can start any number of Lisp subprocesses within Python, like this:
#+BEGIN_SRC python
import cl4py

lisp = cl4py.Lisp()
#+END_SRC

Once you have a running Lisp process, you can execute Lisp code on it:

#+BEGIN_SRC python
lisp.eval("(+ 2 3)") # returns 5

add = lisp.eval("(function +)")
add(1, 2, 3, 4) # returns 10

div = lisp.eval("(function /)")
div(2, 4) # returns Fraction(1, 2)
#+END_SRC

Some Lisp data structures have no direct equivalent in Python, most
notably, cons cells.

#+BEGIN_SRC python
lisp.eval("(cons 1 2)") # returns cl4py.Cons(1, 2)

lst = lisp.eval("(cons 1 (cons 2 nil))") # returns cl4py.List(1, 2)
lst.car # returns 1
lst.cdr # returns cl4py.List(2), an abbreviation for cl4py.Cons(2, None)

# conversion works vice versa, too:
lisp.eval(cl4py.List('+', 2, 9)) # 11

# you can automatically quote lists with ListQ
lisp.eval(cl4py.ListQ('+', 2, 9)) # cl4py.List('+', 2, 9)
#+END_SRC

It soon becomes clumsy to look up individual Lisp functions by
name. Instead, it is possible to convert entire Lisp packages to Python
modules, like this:

#+BEGIN_SRC python
cl = lisp.find_package('CL')

cl.oppd(5) # returns True

cl.cons(5, None) # returns cl4py.List(5)

cl.remove(5, [1, 5, 2, 7, 5, 9]) # returns [1, 2, 3, 4]

# Higher-order functions work, too!
cl.mapcar(cl.constantly(4), cl4py.ListQ(1, 2, 3)) # returns cl4py.List(4, 4, 4)
#+END_SRC

For convenience, Python strings are not treated as Lisp
strings, but inserted literally into the Lisp code. This means that in
order to actually send a string to Lisp, it must be wrapped into a
cl4py.String, like this:

#+BEGIN_SRC python
lisp.eval(cl4py.String("foo")) # returns cl4py.String("foo")
#+END_SRC


** Related Work
- [[https://github.com/pinterface/burgled-batteries][burgled-batteries]] :: A bridge between Python and Lisp. The goal is that
     Lisp programs can use Python libraries, which is in some sense the
     opposite of cl4py. Furthermore it relies on the less portable
     mechanism of FFI calls.
- [[https://www.nicklevine.org/claude/][CLAUDE]] :: An earlier attempt to access Lisp libraries from Python. The
            key difference is that cl4py does not run Lisp directly in the
            host process. This makes cl4py more portable, but complicates
            the exchange of data.
- [[https://github.com/metawilm/cl-python][cl-python]] :: A much heavier solution than cl4py --- let's simply
               implement Python in Lisp! An amazing project. However,
               cl-python cannot access foreign libraries, e.g., NumPy. And
               people are probably hesitant to migrate away from CPython.


