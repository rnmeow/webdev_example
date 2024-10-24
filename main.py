from datetime import datetime
from flask import Flask, jsonify, request

app = Flask(__name__)

@app.route('/', methods = ['GET'])

def ReturnJSON():
	if (request.method == 'GET'):
		dat = {
			'text': 'The backend is now on. 後端已連接。',
			'timestamp': datetime.now(),
		}

	resp = jsonify(dat)

	resp.headers['Content-Type'] = 'application/json; charset=utf-8'
	resp.headers['Access-Control-Allow-Origin'] = '*'

	return resp

if __name__ == '__main__':
	app.run(debug = True)

# dev server runs on `localhost:5000` in my environment
