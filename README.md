# serve

## Usage

Compile the code, put `serve` binary to your $PATH and then you can use it to serve
files from the current direcotry.

```
$ serve
Listening on http://127.0.0.1:3000
```

or you can specify port:

```
$ ./serve --port 8080
Listening on http://127.0.0.1:8080
```

Help:

```
$ serve --help
Usage of serve:
  -port int
        Port to run web server to (default 3000)
```

Other ways to do the same thing: https://gist.github.com/willurd/5720255
