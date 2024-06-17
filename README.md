# Numerical Methods

### Numerical Methods Implementations

Inspired by a university course by the same name, _Numerical Methods_ can calculate and solve tasks similar to those given in class.

The project is based on RESTful API Microservice architecture, where each Microservice is dedicated to a different field of Numerical Methods.


## The Microservices
* **Errors** (Python & Flask): The simplest Microservice, can calculate absolute error, relative error, and a given function's effect on the output.

* **Interpolations** (Go & Chi): Supports Newton's method and Vandermonde method. Uses _Polynomials_ and _LinearEquations_ Microservices for its functionality.
More methods will be implemented in the future.

* **LinearEquations** (Go & Chi): Solves linear equations, using Gaussian Elimination method for efficiency.

* **Polynomials** (Go & Chi): A basic polynomials calculator.

* **Roots** (Python & Flask): Finds the root of a given function. Supports Bisection, Regula Flasi, Fixed Point and Newton's Method.

* **LeastSquares** (Python & Flask): Calculates a polynomial of degree n with minimum distance from all the given points. Uses _LinearEquations_ for its functionality.


## Paths
The basic path to each API is in _Paths_ file.
<br /><br /><br /><br />

**_Note:_** The Python Microservices uses Anaconda environment (_environment.yml_).

