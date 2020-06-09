```docker run -d --name postgres -v ~/docker-data/postgres/96/:/var/lib/postgresql/data -p 5432:5432 postgres:9.6```

ALTER USER postgres PASSWORD 'password';