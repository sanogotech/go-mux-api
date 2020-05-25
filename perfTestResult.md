
*******************
sanogo@DESKTOP-OR1MPIC:~/tools/wrk$ ./wrk -t10 -c100 http://localhost:8010/hello  --timeout 10s
Running 10s test @ http://localhost:8010/hello
  10 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     8.85ms   16.77ms 315.49ms   97.34%
    Req/Sec     1.44k   374.47     6.45k    85.44%
  143881 requests in 10.10s, 17.84MB read
Requests/sec:  14248.02
Transfer/sec:      1.77MB

*******************

WRITE ERROR’s happen when attempting to write on a connection, but it fails because of a closed socket on the server.

READ ERROR’s happen when attempting to read on a connection, but it fails because of a closed socket on the server.

Happy if anybody can confirm or refute that.