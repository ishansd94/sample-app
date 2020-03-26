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

### Testing

```
$ task test
```