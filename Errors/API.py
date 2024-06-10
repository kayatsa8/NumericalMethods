import inspect
from typing import Callable, Dict, List
import ErrorCalculator as claculator
from Response import Response

from flask import Blueprint, Flask, json, request, jsonify
from flask_cors import CORS

import marshal


bp = Blueprint('errors', __name__, url_prefix="/api/errors")

app: Flask = Flask(__name__)
CORS(app)



@bp.route("/")
def welcome() -> str:
    return "Welcome"


@bp.route("/absolute", methods=["POST"])
def absoluteError():
    data: dict = request.get_json()

    schema: List[str] = ["original", "estimated"]

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    response: Response[float] = Response(value=claculator.absoluteError(data["original"], data["estimated"]))

    return jsonify(response.toDict())


@bp.route("/relative", methods=["POST"])
def relativeError():
    data: dict = request.get_json()

    schema: List[str] = ["original", "estimated"]

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    response: Response[float] = Response(value=claculator.relativeError(data["original"], data["estimated"]))

    return jsonify(response.toDict())


@bp.route("/effect", methods=["POST"])
def fEffect():
    data: dict = request.get_json()

    schema: List[str] = ["original", "estimated", "f"]

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    # convert f string to lambda
    lambda_str = data["f"]["lambda"]
    f = eval(lambda_str)

    
    response: Response[float] = Response(value=claculator.functionEffectOnResult(data["original"], data["estimated"], f))

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
    app.run(debug=True, port=3001)