Cook
----------------

``` yaml
# cook.yml
target:
  - host: server
deploy:
  - type: docker-compose
    path: docker-compose.yml
    work_dir: /root
    project_name: test
```
