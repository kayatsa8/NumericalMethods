from typing import Set, List
from Simplex import simplex
from Response import Response

from flask import Blueprint, Flask, request, jsonify
from flask_cors import CORS


bp = Blueprint('LinearProgramming', __name__, url_prefix="/api/lin_prog")

app: Flask = Flask(__name__)
CORS(app)


@bp.route("/")
def welcome() -> str:
    return "Welcome"


@bp.route("/simplex", methods=["POST"])
def bisection():
    data: dict = request.get_json()

    schema: Set[str] = {"objective", "constraints"}

    if not validateRequestSchema(data, schema):
        return Response(error=True, message="bad request body"), 400
    
    response: Response[float]

    try:
        result: List[float] = simplex(objective=data["objective"], constraints=data["constraints"])
        response = Response(value=result)
    except Exception as e:
        response = Response(error=True, message=str(e))
        app.log_exception(e)   

    return jsonify(response.toDict())


def validateRequestSchema(request: dict, schema: Set[str]) -> bool:
    for field in schema:
        if field not in request:
            return False
        
    for key in request:
        if key not in schema:
            return False
        
    return True




if __name__ == "__main__":
    app.register_blueprint(bp)
    app.run(debug=True, port=3006)


