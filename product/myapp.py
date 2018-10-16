from flask import Flask
# import redis

app = Flask(__name__)

# pool = redis.ConnectionPool(host='redis', port=6379)
# rd = redis.Redis(connection_pool=pool)

# http://localhost:5000/create
@app.route('/create')
def create():
    # count = rd.incr('hits')
    # return 'Hello World! {}'.format(count)
    return "product"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
