# 🚀 Go DevOps Lab 🛠️

A hands-on laboratory for DevOps practices with Go Project 🔧

## 🏗️ Prerequisites

- [Go](https://golang.org/doc/install) installed
- [Docker](https://docs.docker.com/get-docker/) installed and running.
- [Minikube](https://minikube.sigs.k8s.io/docs/start/) installed for running Kubernetes locally.
- [kubectl](https://kubernetes.io/docs/tasks/tools/) installed to manage Kubernetes clusters.
- [Kustomize](https://kubectl.docs.kubernetes.io/installation/kustomize/) is used for managing Kubernetes configurations. Install it to apply templates effectively.

## 🚦 Getting Started

### Clone the Repository

```bash
git clone https://github.com/pandakn/go-devops-lab.git
cd go-devops-lab
```

### Start Local Development

##### 1. Start Minikube and Configuration:

```bash
# Start Minikube
minikube start

# Enable ingress addon
minikube addons enable ingress
```

##### 2. Deploy to local cluster:

```bash
kubectl apply -k infrastructure/k8s/overlays/dev
```

##### 3. Run Minikube Tunnel:

```bash
# To expose LoadBalancer services and make them accessible from your local machine
minikube tunnel
```

##### 4. Map Localhost IP to Custom Domain:

```bash
sudo vim /etc/hosts

# Add the following line to map dns
127.0.0.1 api.hello.local
```

##### 5. Test your API:

```bash
curl http://api.hello.local

curl http://api.hello.local?name=world
```

## 📁 Project Structure

```plaintext
.
├── infrastructure
│   └── k8s
│       ├── base
│       │   ├── deployment.yaml
│       │   ├── ingress.yaml
│       │   ├── kustomization.yaml
│       │   └── service.yaml
│       └── overlays
│           ├── dev
│           │   └── kustomization.yaml
│           └── prod
│               └── kustomization.yaml
├── Dockerfile
├── Makefile
├── go.mod
└── main.go
```
