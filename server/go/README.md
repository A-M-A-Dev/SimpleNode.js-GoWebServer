# Golang Helloworld Server

- Served at path `/helloworld/go/`

## API

### HelloWorld

- `GET` API
- Routed at `/helloworld/go/`
- Returns a `text/plain` response containing **Hello World!**

### Adder

- `POST` API
- Routed at `/helloworld/go/adder`
- Request body example:

  ```HTTP
  POST /helloworld/go/adder application/json
  {
    "a": 523,
    "b": 10
  }
  ```
- `JSON` response body example:

  ```HTTP
  200 OK application/json
  {
     "sha256": "fb8a0d2da8683cec6cc64542f95ae11e085c72d56c744b2be5be335295976610"
  }
  ```

### Write

- `GET` API
- Routed at `/helloworld/go/write`
- request example:

  ```HTTP
  GET /helloworld/go/write?line=40
  ```
- response example:

  ```HTTP
  200 OK text/plain
  Line 40 : This is just a sample text
  ```

## Depoyment

- Build the server:

  ```bash
  go build main.go
  ```
- Systemd service configuration is available at [`/os/helloworld-go.service`](/os/helloworld-go.service)
- In the mentioned file, change `<path/to/go/executable/directory>` to absolute path of directory containing `main` executable file, and `<path/to/go/executable>` to absolute path of `main` executable file
- Save the service file at `/lib/systemd/system/helloworld-go.service`
- Enable the service:

  ```bash
  > sudo systemctl enable helloworld-go
  ```
- start the service:

  ```bash
  > sudo systemctl start helloworld-go
  ```
- Now, the server is listening on `0.0.0.0:8080`

## Use reverse proxy with nginx

- Add proxy config to Nginx server as bellow:

  ```NGINX
  location /helloworld/go {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_pass http://localhost:8080;
  }
  ```
- Restart Nginx service:

```bash
> sudo systemctl restart nginx
```
- Full Nginx config is available at [`/nginx/nginx-server-config`](/nginx/nginx-server-config)
