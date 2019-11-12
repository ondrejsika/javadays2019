[sika.io](https://sika.io) | <ondrej@sika.io>

# Docker & Kubernetes Talk @ Java Days 2019

    2019 Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/javadays2019


## Docker

### Build

```
docker build -t ondrejsika/hellojavadays2019 .
```

### Push to Docker Registry

```
docker push ondrejsika/hellojavadays2019
```

### Run

```
docker run --name hellojavadays -d -p 80:80 ondrejsika/hellojavadays2019
```

Go <http://127.0.0.1>

### Stop & Remove

```
docker stop hellojavadays
docker rm hellojavadays
```

## Kubernetes

### Run in Kubernetes

```
helm install hellojavadays ./helm/hellojavadays2019 --set host=hellojavadays2019.k8s.sikademo.com
```

Go <https://hellojavadays2019.k8s.sikademo.com>
