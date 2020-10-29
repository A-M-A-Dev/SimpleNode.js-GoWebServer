# Load Test

- Load test was done with **Locust** tools.

## Installation

```bash
> pip install locust
> locust --version
```

## Run load test for each API

### Go Adder

- Request:
  
  ```HTTP
  POST /helloworld/go/adder application/json
  {
    "a": 2,
    "b": 3
  }
  ```
  
- Load test command:

  ```bash
  > locust -f loadtests/go_adder.py
  ```

- Result as `http://localhost:8089` with 500 users:

  ![](/loadtest/results/go_adder.png)

## #Go Write
    
- Request:

  ```HTTP
  GET /helloworld/go/write?line=68
  ```

- Load test command:

  ```bash
  > locust -f loadtests/go_write.py
  ```

- Result as `http://localhost:8089` with 500 users:

  ![](/loadtest/results/go_write.png)
  
### Node.js Adder
    
- Request:

  ```HTTP
  POST /helloworld/node/adder application/json
  {
    "a": 2,
    "b": 3
  }
  ```

- Load test command:

  ```bash
  > locust -f loadtests/node_adder.py
  ```
  
- Result as `http://localhost:8089` with 500 users:

  ![](/loadtest/results/node_adder.png)
  
### Node.js Write
    
- Request:

  ```HTTP
  GET /helloworld/node/write?line=68
  ```

- Load test command:

  ```bash
  > locust -f loadtests/node_write.py
  ```
  
- Result as `http://localhost:8089` with 500 users:

  ![](/loadtest/results/node_write.png)
