# go-graceful-sqlite

Test signals and results with sqlite.

## Motivation

Explore what happens when a go process is updating an `sqlite` database and receives a signal. Does it get a chance to commit and/or close? Will recent operations be lost? Scenarios to check:

* `<ctrl><C>` to kill a program. Perhaps other signals as well.
* System shutdown. System crash?

The desire is to determine if it is necessary to commit every update of if these can be batched until some time limit is hit, number of records, program exit, etc.

Secondary desire is to explore signal handling in `go` on Linux.

## Building

After not having coded in `go` for about 4 years, I find I need to learn the build process all over again. Looks like the `sqlite3` library I used before is still maintained and comes up first in a search.

```text
go mod init go-sqlite
$EDITOR go-sqlite.go
go get github.com/mattn/go-sqlite3
go run .
```

Interact with resulting database

```text
sqlite3 test.db '.schema'
sqlite3 test.db 'select * from data;'
```

## Strategy

Keep 'pushing' on the program, trying to identify a situation where data written to the database can be lost and then tyr to mitigate that.

## Findings

1. If the program is interrupted with `<ctrl>C`, (during sleep following DB creation) the deferred function seems not to be called but the DB is created.
1. Writing a row to the DB and killing the program with `<ctrl>C` leaves the DB with the row just written.