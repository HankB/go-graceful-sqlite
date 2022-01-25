# go-graceful-sqlite

Test signals and results with sqlite.

## Motivation

Explore what happens when a go process is updating an `sqlite` database and receives a signal. Does it get a chance to commit and/or close? Will recent operations be lost? Scenarios to check:

* `<ctrl><C>` to kill a program. Perhaps other signals as well.
* System shutdown. System crash?

The desire is to determine if it is necessary to commit every update of if these can be batched until some time limit is hit, number of records, program exit, etc.

Secondary desire is to explore signal handling in `go` on Linux.
