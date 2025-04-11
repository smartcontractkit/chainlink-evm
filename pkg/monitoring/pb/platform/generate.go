package writetarget

//go:generate protoc -I=. --go_out=.  ./write_accepted.proto
//go:generate protoc -I=. --go_out=.  ./write_confirmed.proto
//go:generate protoc -I=. --go_out=.  ./write_error.proto
//go:generate protoc -I=. --go_out=.  ./write_initiated.proto
//go:generate protoc -I=. --go_out=.  ./write_sent.proto
//go:generate protoc -I=. --go_out=.  ./write_skipped.proto
