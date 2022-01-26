# go-graceful-sqlite

Test signals and results with sqlite.

## Motivation

Explore what happens when a go process is updating an `sqlite` database and receives a signal. Does it get a chance to commit and/or close? Will recent operations be lost? Scenarios to check:

* `<ctrl><C>` to kill a program. Perhaps other signals as well.
* System shutdown. System crash?

The desire is to determine if it is necessary to commit every update of if these can be batched until some time limit is hit, number of records, program exit, etc.

Secondary desire is to explore signal handling in `go` on Linux.

## Building

Environment: tested on Debian Bullseye, aarch64 and x86_64.

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
1. Looping over database writes (and console prints) it seems like no updates are lost. If the value is printed before the insert operation, the insert operation does not reflect the last value printed, but it is not clear if this is because the insert operation did not commit or if it never got called.
1. Shutting down the host (`shutdown -r now`) seems to not cause loss of data. Tested on a Prapberry Pi 4B. Last number seen in a terminal window was 117. Last value in the database was 122. Saw similar result with Gnome disabled, running from a text login and hitting <ctrl><alt><del> for a quicker (?) shutdown.

## Errata

The sqlite3 package is not required to build and run but may be convenient to examine the resulting database.