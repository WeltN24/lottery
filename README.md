Lottery
=======

<img src="http://img.memecdn.com/winning-the-lottery-it-could-have-been-worse-i-guess_c_3237821.jpg" width=400>

DEV:

- install go
- get a token from github [tokens](https://github.com/settings/tokens) save it under .token
- use makefile


USE:
run the server:

```
$ ./lottery
```

get numbers:
```
$ curl localhost:8080/numbers
[42,37,26,14,33]
```

health:
```
$ curl localhost:8080/health
up
```

metrics:
```
$ curl localhost:8080/debug/metrics
{
	...
	"numbers.50-percentile": 0,
	"numbers.75-percentile": 0,
	"numbers.95-percentile": 0,
	"numbers.99-percentile": 0,
	"numbers.999-percentile": 0,
	"numbers.count": 0,
	"numbers.max": 0,
	"numbers.mean": 0,
	"numbers.min": 0,
	"numbers.std-dev": 0,
	"timer.50-percentile": 0,
	"timer.75-percentile": 0,
	"timer.95-percentile": 0,
	"timer.99-percentile": 0,
	"timer.999-percentile": 0,
	"timer.count": 0,
	"timer.fifteen-minute": 0,
	"timer.five-minute": 0,
	"timer.max": 0,
	"timer.mean": 0,
	"timer.mean-rate": 0,
	"timer.min": 0,
	"timer.one-minute": 0,
	"timer.std-dev": 0
}
```
