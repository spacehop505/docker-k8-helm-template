# Installation

## Golang
* build golang executable

```bash
go build -o bin/project
```

## Docker
* build image

```bash
docker build $PWD --file docker/Dockerfile --tag my-docker-image:0.0.1
```

* run image 

```bash
docker run -p 8080:8080 my-docker-image:0.0.1
```

## Minikube
* load local docker image to minikube

```bash
minikube image load my-docker-image:0.0.1
```

## Helm

* helm install

```bash
helm install -n <namespace> sample-project chart/my-helm
```

* render chart template locally

```bash
helm template sample-project chart/my-helm --debug
```

* list chart

```bash
helm list
```

# Uninstallation 

* uninstall helm chart
```bash
helm uninstall -n <namespace> sample-project
```

* remove docker image from minikube

```bash
minikube image rm my-docker-image:0.0.1
```

* remove docker image from docker 

```bash
docker rmi my-docker-image:0.0.1
```