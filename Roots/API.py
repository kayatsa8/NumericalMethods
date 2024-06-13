from typing import List
import RootMethods as roots
from Response import Response

import math

from flask import Blueprint, Flask, json, request, jsonify
from flask_cors import CORS


bp = Blueprint('errors', __name__, url_prefix="/api/roots")

app: Flask = Flask(__name__)
CORS(app)



@bp.route("/")
def welcome() -> str:
    return "Welcome"


@bp.route("/bisection", methods=["POST"])
def bisection():
    data: dict = request.get_json()

    schema: List[str] = ["a", "b", "f", "delta"]

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    # convert f string to lambda
    lambda_str = data["f"]["lambda"]
    f = eval(lambda_str, {"math": math})
    
    response: Response[float]

    try:
        result: float = roots.bisection(a=data["a"], b=data["b"], f=f, delta=data["delta"])
        response = Response(value=result)
    except Exception as e:
        response = Response(error=True, message=str(e))
        app.log_exception(e)   


    return jsonify(response.toDict())


@bp.route("/regula_falsi", methods=["POST"])
def regulaFalsi():
    data: dict = request.get_json()

    schema: List[str] = ["a", "b", "f", "delta"]

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    # convert f string to lambda
    lambda_str = data["f"]["lambda"]
    f = eval(lambda_str, {"math": math})
    
    response: Response[float]

    try:
        result: float = roots.regulaFalsi(a=data["a"], b=data["b"], f=f, delta=data["delta"])
        response = Response(value=result)
    except Exception as e:
        response = Response(error=True, message=str(e))
        app.log_exception(e)   


    return jsonify(response.toDict())


@bp.route("/fixed_point", methods=["POST"])
def fixedPoint():
    data: dict = request.get_json()

    schema: List[str] = ["guess", "g", "delta"]

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    # convert g string to lambda
    lambda_str = data["g"]["lambda"]
    g = eval(lambda_str, {"math": math})
    
    response: Response[float]

    try:
        result: float = roots.fixedPoint(guess=data["guess"], g=g, delta=data["delta"])
        response = Response(value=result)
    except Exception as e:
        response = Response(error=True, message=str(e))
        app.log_exception(e)   


    return jsonify(response.toDict())


@bp.route("/newton_method", methods=["POST"])
def newronMethod():
    data: dict = request.get_json()

    schema: List[str] = ["guess", "f", "fTag", "delta"]

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    # convert f string to lambda
    lambda_str = data["f"]["lambda"]
    f = eval(lambda_str, {"math": math})

    # convert fTag string to lambda
    lambda_str = data["fTag"]["lambda"]
    fTag = eval(lambda_str, {"math": math})
    
    response: Response[float]

    try:
        result: float = roots.newtonMethod(guess=data["guess"], f=f, fTag=fTag, delta=data["delta"])
        response = Response(value=result)
    except Exception as e:
        response = Response(error=True, message=str(e))
        app.log_exception(e)   


    return jsonify(response.toDict())




def validateRequestSchema(request: dict, schema: List[str]) -> bool:
    for field in schema:
        if field not in request:
            return False
        
    for key in request:
        if key not in schema:
            return False
        
    return True




if __name__ == "__main__":
    app.register_blueprint(bp)
    app.run(debug=True, port=3002)