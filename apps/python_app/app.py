from flask import Flask

app = Flask(__name__)


@app.route("/")
def me_api():
    return {
        "hello": "I'm a json response", "from": "app python"
    }
