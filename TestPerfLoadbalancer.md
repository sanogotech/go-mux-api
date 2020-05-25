C:\Users\User>ab -c10  -n1000  -k http://localhost:8010/hello/

Requests per second:    13513.88 [#/sec] (mean)
Time per request:       0.740 [ms] (mean)
Time per request:       0.074 [ms] (mean, across all concurrent requests)
Transfer rate:          2639.43 [Kbytes/sec] received

-----------
C:\Users\User>ab -c10  -n1000  -k http://localhost:8090/hello/
Requests per second:    11111.36 [#/sec] (mean)
Time per request:       0.900 [ms] (mean)
Time per request:       0.090 [ms] (mean, across all concurrent requests)
Transfer rate:          2170.19 [Kbytes/sec] received

----