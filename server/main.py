import requests
import os

import yaml
from flask import Flask, jsonify, Response, request, send_from_directory
from flask_cors import CORS


app = Flask(__name__)

PORT = 8000

# Note: Setting CORS to allow chat.openapi.com is required for ChatGPT to access your plugin
CORS(app, origins=[f"http://localhost:{PORT}", "https://chat.openai.com"])

api_url = 'http://localhost:8001'


@app.route('/.well-known/ai-plugin.json')
def serve_manifest():
    return send_from_directory(os.path.dirname(__file__)+"/.well-known", 'ai-plugin.json')


@app.route('/.well-known/logo.png')
def serve_logo():
    return send_from_directory(os.path.dirname(__file__)+"/.well-known", 'logo.png')


@app.route('/.well-known/openapi.yaml')
def serve_openapi_yaml():
    with open(os.path.join(os.path.dirname(__file__)+"/.well-known", 'openapi.yaml'), 'r') as f:
        yaml_data = f.read()
    yaml_data = yaml.load(yaml_data, Loader=yaml.FullLoader)
    return jsonify(yaml_data)


@app.route('/<path:path>', methods=['GET', 'POST'])
def wrapper(path):

    headers = {
        'Content-Type': 'application/json',
    }

    url = f'{api_url}/{path}'
    print(f'Forwarding call: {request.method} {path} -> {url}')

    if request.method == 'GET':
        response = requests.get(url, headers=headers, params=request.args)
    elif request.method == 'POST':
        print(request.headers)
        response = requests.post(url, headers=headers, params=request.args, json=request.json)
    else:
        raise NotImplementedError(f'Method {request.method} not implemented in wrapper for {path=}')
    return response.content


if __name__ == '__main__':
    app.run(port=PORT)