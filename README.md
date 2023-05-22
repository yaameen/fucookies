# Fucookies

This is a simple proxy server written in Go that allows your local React app to communicate with an API protected by CORS. The proxy server handles the CORS headers and forwards the requests to the API.

## Prerequisites

- Go 1.16 or higher

## Getting Started

1. Clone the repository:

```bash
git clone https://github.com/yaamynu/fucookies.git
``` 

2. Navigate to the project directory:

```bash
cd fucookies
```

3. Install dependencies:

```bash
go get -d ./...
```

4. Configure the proxy server:

Open the config.yaml file and update the following values according to your requirements:

- `target`: The URL of the API that you want to proxy.
- `port`: The port on which the proxy server will listen.
- `allowed_headers`: List of allowed headers in the CORS requests.
- `cookie.domain.from`: The domain from which the cookie should be rewritten.
- `cookie.domain.to`: The domain to which the cookie should be rewritten.

Following is a snippet from the config.yaml file:

```yaml
listen_on: 0.0.0.0:3000
target: "https://example.com/"
allowed_origin: http://localhost:8080
allow_credentials: true
port: 443
allowed_headers:
  - Content-Type
  - Authorization
cookie:
  domain:
    from: example.com
    to: localhost
```

5. Build and run the proxy server:

```bash
go run main.go
```
The proxy server will start listening on the specified port.

6. Configure your React app:

Update the API URL in your React app to use the URL of the proxy server (http://localhost:<port>).

For example, if your proxy server is running on port 3000:

```javascript
const apiUrl = 'http://localhost:3000/api';
```
Make sure to adjust the URL path according to your API's endpoint structure.

7. Start your React app:

```bash
npm start
```
Your React app should now be able to communicate with the API through the proxy server.


## License
This project is licensed under the MIT License.


Feel free to customize the README file further to include additional details specific to your project or any usage instructions you find necessary.
