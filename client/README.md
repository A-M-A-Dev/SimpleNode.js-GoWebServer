# Helloworld client

- Served at `/helloworld/`
- An static `HTML/CSS/javascript` user interface

## Pages

### Home

- Routed at `/helloworld/`

### Adder

- Routed at `/helloworld/adder.html`

### Write

- Routed at `/helloworld/write.html`

### 404 Page

- Every 404 error redirects to this page (If request accepts `HTML` response)

## Serve staticly on nginx

- Add following configurations to nginx server (Change `</path/to/helloworld/root/dir>` to absolute path of client directory):

  ```NGINX
  location /helloworld {
    alias </path/to/helloworld/root/dir>;
    index helloworld.html;
  }

  error_page 404 /helloworld/404.html;
  ```
- Restart Nginx service:

```bash
> sudo systemctl restart nginx
```
- Full Nginx config is available at [`/nginx/nginx-server-config`](/nginx/nginx-server-config)
