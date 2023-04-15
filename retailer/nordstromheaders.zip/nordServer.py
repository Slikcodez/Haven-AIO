from flask import Flask, request, jsonify

app = Flask(__name__)
headers = []

@app.route('/api/receiveHeaders', methods=['POST'])
def receive_headers():
    json_data = request.get_json()

    if not all(key in json_data for key in ['HeaderX', 'Special1', 'Special2', 'HeaderF', 'HeaderD', 'HeaderB', 'HeaderC', 'HeaderA', 'HeaderZ']):
        return jsonify({"message": "Missing keys in JSON data"}), 400

    headers.append(f"{json_data['HeaderX']}:::{json_data['HeaderX'].split('_')[0]}:::{json_data['Special1']}:::{json_data['Special2']}:::{json_data['HeaderF']}:::{json_data['HeaderD']}:::{json_data['HeaderB']}:::{json_data['HeaderC']}:::{json_data['HeaderA']}:::{json_data['HeaderZ']}")
    return '', 200

@app.route('/api/getHeaders', methods=['GET'])
def get_headers():
    if not headers:
        return jsonify({"message": "No headers available"}), 404

    header = headers.pop(0)
    return header, 200

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=6942)
