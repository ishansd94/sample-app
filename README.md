# sample-app
Web service to create kubernetes secrets.

### Pre-requisites

```
1. https://github.com/go-task/task (Makefile alternative.)
2. GO
3. https://github.com/golang/dep (Go dependency manager)
```

### Installation

Clone the repo in your $GOPATH.
This project uses Dep (Golang vendoring tool) https://github.com/golang/dep

```sh
$ cd $GOPATH/src/github.com/ishansd94/sample-app
$ task install
$ task run
```
*NOTE: Default port is ```:8000```. Port can be changed by setting ```PORT``` environment variable*

### Build

Build parameters are available in the ```Taskfile.yml```
Change ```USERNAME``` and ```IMAGE``` parameters with your docker hub username and desired image name.

```sh
$ task build
```
If you use separate key for gitlab, change the location of the private key file.
```
SSH_PRIVATE_KEY: $(cat ~/.ssh/id_rsa)
```

### Usage
In order for kube-quotas to work the ```ServiceAccount``` within the pod where it's running should have the necessary RBAC permissions.

```
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
    name: sample-app
rules:
  - apiGroups: [""]
    resources: [""]
    verbs: ["get", "list", "watch", "create", "update", "patch", "delete"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: sample-app
subjects:
  - kind: Group
    name: system:serviceaccounts
    apiGroup: ""
roleRef:
  kind: ClusterRole
  name: sample-app
  apiGroup: rbac.authorization.k8s.io
```
*NOTE: For local clusters this is not needed.* 

##### Deploy to Kubernetes

Create a seperate ```namespace``` for kube-quotas ex: ```app``` and create a ```deployment```.

```
$ kubectl create ns app
$ kubectl create deployment sample-app --image=emzian7/sample-app -n app
```

##### Using sample-app web service

Get the pod ip using,
```
$ kubectl get pods -n app -o wide
```

##### Payloads
---
##### 1. Creating Items  

Expected payload as a ```POST``` request.

```
{
    "Field1": <string>,
    "Field2": <json obj>
}
```
*NOTE: Field2 is mapped to map[string]string, json obj expected is something like {"foo": "bar"}. If content field is not specified default uuid will be created*. 

```
$ curl -d '{"Field1":"foo", "Field2":{"foo":"bar"}}' -H "Content-Type: application/json" -X POST <sample-app pod ip>:8000
```

### Testing

```
$ task test
```