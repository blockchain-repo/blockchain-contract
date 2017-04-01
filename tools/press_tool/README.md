## press tool
### [boom](https://github.com/rakyll/boom)

#### install
```
    https://github.com/rakyll/boom
```

#### Usage
```
usage: boom [-h] [--version] [-m {GET,POST,DELETE,PUT,HEAD,OPTIONS}]
              [--content-type CONTENT_TYPE] [-D DATA] [-c CONCURRENCY] [-a AUTH]
              [--header HEADER] [--pre-hook PRE_HOOK] [--post-hook POST_HOOK]
              [--json-output] [-n REQUESTS | -d DURATION]
              [url]

Simple HTTP Load runner.

positional arguments:
    url                   URL to hit

optional arguments:
    -h, --help            show this help message and exit
    --version             Displays version and exits.
    -m {GET,POST,DELETE,PUT,HEAD,OPTIONS}, --method {GET,POST,DELETE,PUT,HEAD,OPTIONS}
                          HTTP Method
    --content-type CONTENT_TYPE
                          Content-Type
    -D DATA, --data DATA  Data. Prefixed by "py:" to point a python callable.
    -c CONCURRENCY, --concurrency CONCURRENCY
                          Concurrency
    -a AUTH, --auth AUTH  Basic authentication user:password
    --header HEADER       Custom header. name:value
    --pre-hook PRE_HOOK   Python module path (eg: mymodule.pre_hook) to a
                          callable which will be executed before doing a request
                          for example: pre_hook(method, url, options). It must
                          return a tupple of parameters given in function
                          definition
    --post-hook POST_HOOK
                          Python module path (eg: mymodule.post_hook) to a
                          callable which will be executed after a request is
                          done for example: eg. post_hook(response). It must
                          return a given response parameter or raise an
                          `boom.boom.RequestException` for failed request.
    --json-output         Prints the results in JSON instead of the default
                          format
    -n REQUESTS, --requests REQUESTS
                          Number of requests
    -d DURATION, --duration DURATION
                          Duration in seconds
```

#### example
```
Run Result:
    % boom -n 1000 -c 100 https://google.com
     
    1000 / 1000 ∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎ 100.00 %
     
    Summary:
      Total:        21.1307 secs.
      Slowest:      2.9959 secs.
      Fastest:      0.9868 secs.
      Average:      2.0827 secs.
      Requests/sec: 47.3246
      Speed index:  Hahahaha
     
    Response time histogram:
      0.987 [1]     |
      1.188 [2]     |
      1.389 [3]     |
      1.590 [18]    |∎∎
      1.790 [85]    |∎∎∎∎∎∎∎∎∎∎∎
      1.991 [244]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
      2.192 [284]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
      2.393 [304]   |∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎∎
      2.594 [50]    |∎∎∎∎∎∎
      2.795 [5]     |
      2.996 [4]     |
     
    Latency distribution:
      10% in 1.7607 secs.
      25% in 1.9770 secs.
      50% in 2.0961 secs.
      75% in 2.2385 secs.
      90% in 2.3681 secs.
      95% in 2.4451 secs.
      99% in 2.5393 secs.
 
    Status code distribution:
      [200] 1000 responses
```