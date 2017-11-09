Cook
------

Cook provide an extremely simple and cheap deployment workflow for docker / compose, to save poor man's time

Install
------
``` bash
go install github.com/jjyr/cook
```

Build
-----

``` bash
go get github.com/jjyr/cook
cd $GOPATH/src/github.com/jjyr/cook
go test github.com/jjyr/cook/...
go build -v
```

Usage
-----

Cook provide a simple deployment workflow for *docker-compose* user.

You need put a *cook.yml* file into your project directory, type 'cook config --sample' to see how to write *cook.yml*.
``` yaml
# cook.yml sample
target:
- host: my_web_server
  user: web
deploy:
- path: docker-compose.yml

# 'target' is a list of servers
# make sure ssh web@my_web_server works
# cook use ssh to transfer images, and bring services up

# 'deploy' is a list of docker-compose
```

You can use 'cook config' to check current *cook.yml*

Now, *cd* into your project directory, create *cook.yml*, then type 'cook'. cook will start deployment workflow:

1) Run 'docker-compose build' on local-machine
2) Push local images to target servers (trough ssh)
3) Run 'docker-compose up' on target servers
4) Check services status, make sure they are 'up'

That's it!


**NOTICE**:

Typically *docker-compose* user use 'build' options in service, let 'docker-compose build' command auto generate a name for service image. But cook don't know how to push this type images, so the 'image' option need to clearly set. 

``` yaml
# docker-compose.yml
services:
  web:
    build: .
    image: mywebapp
```  

LICENSE
------
Apache License Version 2.0
