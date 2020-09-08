# bookstore_oauth-api

## Stack

- golang + gin
- cassandra db(v3.11.8)
  - https://hub.docker.com/_/cassandra

## go mod command

```
# go mod init
# go get golang.org/x/text
```

## cassandra query

```
> create keyspace oauth with replication = {'class': 'SimpleStrategy', 'replication_factor':1};
> use oauth;
> create table access_tokens(access_token varchar primary key, user_id bigint, client_id bigint, expires bigint);
```
