from typing import Dict, List
import LeastSquares as LS
from Polynomial import Polynomial
from Response import Response

from flask import Blueprint, Flask, request, jsonify
from flask_cors import CORS

bp = Blueprint('LeastSquares', __name__, url_prefix="/api/least_squares")

app: Flask = Flask(__name__)
CORS(app)


@bp.route("/")
def welcome() -> str:
    return "Welcome"

@bp.route("", methods=["POST"])
def leastSquares():
    data: dict = request.get_json()

    schema: List[str] = ["points", "degree"]

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    points: Dict[float, float] = makePointsDict(data["points"])

    result: Polynomial = LS.LeastSquares(points, data["degree"])
    
    response: Response[float] = Response(value=result.toDict())

    return jsonify(response.toDict())


def makePointsDict(pointsDictStr: Dict[str, float]) -> Dict[float, float]:
    return {float(x): y for x, y in pointsDictStr.items()}


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
    app.run(debug=True, port=3005)

