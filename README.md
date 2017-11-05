Cook
----------------

``` yaml
# cook.yml
build:
  - path: ./Dockerfile
    image: example
    build_dir: .
Target:
  - host: server
deploy:
  - type: docker-compose
    path: docker-compose.yml
    work_dir: /root

```