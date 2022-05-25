# example-docker-kubernetes-helm

## build golang executable

```bash
go build -o bin/project
```

## build image

```bash
docker build $PWD --file docker/Dockerfile --tag sample-image:0.0.1
```


## run image 

```bash
docker run -p 8090:8090 sample-image:0.0.1
```

# load image to minikube

```bash
minikube image load sample-image:0.0.1
minikube image rm sample-image:0.0.1
```

# helm install

```bash
helm install sample-project chart/sample-helm
helm uninstall sample-project
helm list
```