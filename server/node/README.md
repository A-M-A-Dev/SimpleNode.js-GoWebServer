# Node.js Helloworld Server

- Served at path `/helloworld/node/`
- Created by `express js`

## API

### HelloWorld

- `GET` API
- Routed at `/helloworld/node/`
- Returns a `text/plain` response containing **Hello World!!**

### Adder

- `POST` API
- Routed at `/helloworld/node/adder`
- Request body example:

  ```HTTP
  POST /helloworld/node/adder application/json
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
- Routed at `/helloworld/node/write`
- request example:

  ```HTTP
  GET /helloworld/node/write?line=40
  ```
- response example:

  ```HTTP
  200 OK text/plain
  Line 40 : This is just a sample text
  ```

## Depoyment

- Systemd service configuration is available at [`/os/helloworld-node.service`](/os/helloworld-node.service)
- In the mentioned file, change `<path/to/nodejs/main/file>` to absolute path of `indes.js` file
- Save the service file at `/lib/systemd/system/helloworld-node.service`
- Enable the service:

  ```bash
  > sudo systemctl enable helloworld-node
  ```
- start the service:

  ```bash
  > sudo systemctl start helloworld-node
  ```
- Now, the server is listening on `0.0.0.0:3000`

## Use reverse proxy with nginx

- Add proxy config to Nginx server as bellow:

  ```NGINX
  location /helloworld/node {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_pass http://localhost:3000;
  }
  ```
- Full Nginx config is available at [`/nginx/nginx-server-config`](/nginx/nginx-server-config)
