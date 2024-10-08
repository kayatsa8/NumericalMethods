# Numerical Methods

### Numerical Methods Implementations

Inspired by a university course by the same name, _Numerical Methods_ can calculate and solve tasks similar to those given in class.

The project is based on RESTful API Microservice architecture, where each Microservice is dedicated to a different field of Numerical Methods.


## The Microservices
* **Errors** (Python & Flask): The simplest Microservice, can calculate absolute error, relative error, and a given function's effect on the output.

* **Interpolations** (Go & Chi): Supports Newton's method, Vandermonde's method, Lagrange's method, Herrmite's method, Piecewise Linear interpolation & Cubic Spline interpolation. Uses _Polynomials_ and _LinearEquations_ Microservices for its functionality.

* **LinearEquations** (Go & Chi): Solves linear equations, using Gaussian Elimination method for efficiency.

* **Polynomials** (Go & Chi): A basic polynomials calculator, supports polynomial addition, polynomial multiplication and derivative.

* **Roots** (Python & Flask): Finds the root of a given function. Supports Bisection, Regula Flasi, Fixed Point and Newton's Method.

* **LeastSquares** (Python & Flask): Calculates a polynomial of degree n with minimum distance from all the given points. Uses _LinearEquations_ for its functionality.

* **LinearProgramming** (Python & Flask): solves linear programming problems and returns the value of each variable and the value of the target. 


## Paths
The basic path to each API is in _Paths_ file.

## Endpoints

* **Errors:**
    * Absolute Error:
        ```bash
        POST Errors:http://<host>:3001/api/errors/absolute
        ```
    * Relative Error:
        ```bash
        POST Errors:http://<host>:3001/api/errors/relative
        ```
    * Function's Effect on Error:
        ```bash
        POST Errors:http://<host>:3001/api/errors/effect
        ```

<br /><br /><br /><br />

**_Note:_** The Python Microservices uses Anaconda environment (_environment.yml_).

