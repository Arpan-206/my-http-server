# HTTP Server
This is a simple proof of concept of a HTTP server written in Go.


### How to run

1. Install Go
2. Clone this repository with `git clone https://github.com/Arpan-206/my-http-server.git`
3. Run `./your_server.sh --directory /tmp/` to start the server

### How to connect
There are many routes you can connect to. Here are some examples:

1. GET request to `http://localhost:4221/` will return a 200 status code without content.
2. GET request to `http://localhost:4221/echo/{some_text}` will return a 200 status code with the text you provided.
3. GET request to `http://localhost:4221/user_agent` will return a 200 status code with the user agent of the request.
4. GET request to `http://localhost:4221/files/{file_name}` will return a 200 status code with the content of the file you provided.
5. POST request to `http://localhost:4221/files/{file_name}` with a body will create a file with the content of the body and return a 200 status code.


### How to look around the code
To look around the code, you can open the `app/server.go` file. This is where the server is defined and the routes are set up. Rest of the files are just helper files.
