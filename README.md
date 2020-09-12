# rap

`rap` is a dev environment management tool.

Based on `docker` and `docker-compose`, help developers to rapidly set up and easily manage development enviroments, e.g. `mysql` `mongodb` `redis` ...

## Config

`rap` use `rap.yml` to manage environment dependencies.

```yml
# rap.yml

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
# interactively get depns info to create rap.yml
$ rap init

# prepare enviroments, e.g. download or build docker images
$ rap prepare

# start|stop|restart all depns as rap.yml described
$ rap start|stop|restart

# interactivley add a new depn info
$ rap add

```

## Usage

```bash

# rap [FLAGS] [COMMAND]

## FLAGS

## COMMAND

```
