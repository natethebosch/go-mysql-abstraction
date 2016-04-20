# go-mysql-abstraction

## Functionality
This package provides two functions

```
QueryOneRow( <query-string> ) returns mysql.Row`

Performs a single row fetch. Connections are managed by the library 
in pool of 20 connections which are replenished as needed.

On error, nil is returned
```
  
```
BulkQuery ( <query-string> ) returns chan mysql.Row

Uses a buffered channel and LIMIT <offset>, <count> to get 10 rows at a time.
Connections are managed in the same pool as QueryOneRow.

On error, the channel will be closed
```

## Setup
Setup server, user, password and database with
```
SetConnectionInfo("127.0.0.1:3306", "<user>", "<password>", "<database>")
```

## Caviates

There cannot be more than <connection-pool-size> queries at any given time.
