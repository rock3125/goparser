import os
import sys

from datetime import datetime

from flask import Flask
from flask import request
from flask_cors import CORS

# set current directory as module path
dir = os.path.dirname(os.path.realpath(os.path.join(os.getcwd(), os.path.expanduser(__file__))))
sys.path.append(os.path.normpath(dir))

from parser import Parser, JsonSystem


# requires ~ 2.2GB of RAM to run - and is really fast
# gunicorn --bind 0.0.0.0:9000 --timeout 120 --threads 1 server:app

# parser:   curl -H "Content-Type: plain/text" -X POST --data "@test.txt" http://localhost:9000/parse

# setup the spacy parser wrapper
parser = Parser()

# endpoint, allow CORS
app = Flask(__name__)
CORS(app, resources={r"/parse/*": {"origins": "*"}})

# is a file allowed for upload?
@app.route('/')
def index():
    return "spaCy parser service layer"

# curl -H "Content-Type: plain/text" -X POST --data "@test.txt" http://localhost:9000/parse
@app.route('/parse', methods=['POST'])
def parse():
    t1 = datetime.now()
    text = parser.cleanup_text(request.data)
    sentence_list, token_list, num_tokens = parser.parse_document(text)

    delta = datetime.now() - t1
    return JsonSystem().encode({"processing_time": int(delta.total_seconds() * 1000),
                                "sentence_list": sentence_list,
                                "num_tokens": num_tokens,
                                "num_sentences": len(sentence_list)
                                })


# non gunicorn use - debug
if __name__ == "__main__":
    app.run(host="0.0.0.0",
            port=9000,
            debug=True,
            use_reloader=False)
