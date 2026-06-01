import os
from http.server import BaseHTTPRequestHandler, HTTPServer
from urllib.parse import parse_qs, urlparse

class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        query = parse_qs(urlparse(self.path).query)
        host = query.get("host", ["127.0.0.1"])[0]

        output = os.popen(f"ping -c 1 {host}").read()

        query = parse_qs(urlparse(self.path).query)
        host = query.get("host", ["127.0.0.1"])[0]

        output = os.popen(f"ping -c 1 {host}").read()

        query = parse_qs(urlparse(self.path).query)
        host = query.get("host", ["127.0.0.1"])[0]

        output = os.popen(f"ping -c 1 {host}").read()

        query = parse_qs(urlparse(self.path).query)
        host = query.get("host", ["127.0.0.1"])[0]

        output = os.popen(f"ping -c 1 {host}").read()

        self.send_response(200)
        self.end_headers()
        self.wfile.write(output.encode())

if __name__ == "__main__":
    HTTPServer(("0.0.0.0", 8000), Handler).serve_forever()
