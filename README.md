# rap

`rap` is a dev environment management tool.

Based on `docker` and `docker-compose`, help developers to rapidly set up and easily manage development enviroments, e.g. `mysql` `mongodb` `redis` ...

## Config

`rap` use `rap.yaml` to manage environment dependencies.

```yaml
# rap.yaml

# env. depns
depns:
  mysql: mysql@5.7
  mongo: mongo@3.6
  cache: redis@5.0
  ...

# [optional] extra configs for each env
configs:
  mysql:                               # depns name
    volumes:                           # specify storages, default is `.rap/vol/[name]`
      - .rap/vol/mysql:/var/lib/mysql
    ports:                             # specify ports map
      - "3306:3306"
    environment:                       # specify container enviroment variables
      MYSQL_ROOT_PASSWORD: 123
  ...
```

## Examples

```bash
# start|stop|restart all depns as rap.yaml described
$ rap start|stop|restart

# interactivley add a new depn info
$ rap add

# start a web server to manage rap
$ rap serve

# download & build images specified in docker-compose file
$ rap download

```

## Usage

```bash

# rap [FLAGS] [COMMAND]

## FLAGS

## COMMAND

```
