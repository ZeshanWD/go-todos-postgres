# Golang REST API with PostgreSQL

This code is part of an series of videos that I'm doing on my [Youtube Channel](https://www.youtube.com/channel/UCr-fU3wpAnV0HSYLhRXeJLQ)

In this series we are creating a basic todos CRUD Api with the [Golang Programming Language](https://golang.org/) and PostgreSQL as our database.

If you want to follow the youtube playlist, you can click [here](https://www.youtube.com/watch?v=iHssTgnVxDY&list=PLKzZ5IHxJJWcNHqhSCHETZXHMcr9gqM2T)!

## Dockerize Postgres

if you want to run postgres in a docker container like I did, you can follow the next instructions:


```bash

docker run -d --name postgres -v ~/docker-data/postgres/96/:/var/lib postgresql/data -p 5432:5432 postgres:9.6

```

after running the container, you need access psql and change the password for `postgres` user

```mysql

ALTER USER postgres PASSWORD 'password';

```