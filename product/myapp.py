from flask import Flask

app = Flask(__name__)

@app.route('/')
def home():
    return "product service"


@app.route('/list')
def create():
    return "product list"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
