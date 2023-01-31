<h1 align="center">
  S7CC02
</h1>

<br />

Cloud Computing second project. Create, Dockerise, and Deploy an entire cryptometer service 
with Golang, Docker, and Kubernetes. In this
project I build a Golang service with _Fiber_ framework to get crypto market values and cache them into a Redis database.

## How to run?

Run the application without any configs on your local.

```shell
go build -o main
./main
```

Make a copy of config file:

```shell
cp config.example.yml config.yml
```

Update config values:

```yaml
port: 6030

crypto:
  name: "bitcoin"
  value: "usd"

storage:
  host: "cryptometer-redis-cluster:6379"
  password: ""
  timeout: 5
```

## Endpoints

```json
[
  {
    "endpoint": "/api",
    "method": "get",
    "response": [
      {
        "code": "200",
        "body": {
          "name": "bitcoin",
          "value": 98.202
        }
      },
      {
        "code": "500",
        "body": "Error"
      }
    ]
  },
  {
    "endpoint": "/hlz",
    "method": "get",
    "response": [
      {
        "code": "200",
        "body": "OK [time]"
      },
      {
        "code": "500",
        "body": "Error"
      }
    ]
  }
]
```

## Docker

Use the following image:

```shell
docker pull amirhossein21/cryptometer:v0.2
```

Execute on container:

```shell
docker run --network=cryptonet --name cryptometer -d -p 6030:6030 --mount type=bind,source="$(pwd)"/config.yml,target=/app/config.yml amirhossein21/cryptometer:v0.2
```

## Kubernetes

Configmap:

```yaml
apiVersion: v1
kind: ConfigMap

metadata:
  name: cryptometer-configmap

data:
  config.yml:  |+
    port: 6030
    crypto:
      name: "bitcoin"
      value: "usd"
    storage:
      host: "redis-service:6379"
      password: "12345678"
      timeout: 5
```

Deployment:

```yaml
apiVersion: apps/v1
kind: Deployment

metadata:
  name: cryptometer

spec:
  replicas: 2
  selector:
    matchLabels:
      app: cryptometer
  template:
    metadata:
      labels:
        app: cryptometer
    spec:
      containers:
        - name: cryptometer
          image: amirhossein21/cryptometer:v0.2
          ports:
            - containerPort: 6030
              name: cryptometer
          resources:
            requests:
              memory: "64Mi"
              cpu: "250m"
            limits:
              memory: "128Mi"
              cpu: "500m"
          volumeMounts:
            - name: config
              mountPath: /app/config.yml
              subPath: config.yml
      volumes:
        - name: config
          configMap:
            name: cryptometer-configmap
```

Service:

```yaml
apiVersion: v1
kind: Service

metadata:
  name: cryptometer

spec:
  selector:
    app: cryptometer
  ports:
    - protocol: TCP
      port: 6030
      targetPort: 6030
```