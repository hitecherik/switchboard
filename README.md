# `switchboard`

A simple Go server that allows you to specify various redirects in
[TOML](https://toml.io).

While it's possible to write this directly in various reverse proxies, this is
simpler and easier to configure. I made this mainly since some other projects
involve setting up a lot of redirects, and the TOML notation makes it quicker
to define them.

## Setup

## Configuration

First, define a `config.toml` file that contains your redirects.

```toml
"/abcd" = "http://example.com"
"/search/duck/duck/go" = "https://duckduckgo.com"

["example.com"]
"/foo/*" = "https://github.com"
"/" = "https://wikipedia.org"
```

At the top of the file, we defined two redirects that will be performed
regardless of the host that the request is made to. Then, we define two
additional redirects for only the `example.com` host.

`"/foo/*" = "https://github.com"` means that any path beginning with `/foo/`
will be redirected to `https://github.com`.

The precedence of redirect rules is:

1. Any rules specified for a particular host.
2. Rules specified for all hosts in the order in which they're defined.

## Run the server

Next, we can install the program and run the HTTP server.

```bash
$ go get github.com/hitecherik/switchboard
$ switchboard -config path/to/config.toml
```

This will run the server on port 8000 and will use the `301 Moved Permanently`
HTTP code for its redirects.

The full options for `switchboard` are:

```
Usage of switchboard:
  -config string
      file defining redirects (default "config.toml")
  -port uint
      port to listen on (default 8000)
  -temporary
      whether the redirect should be temporary
```

If the `-temporary` flag is passed, `switchboard` will use the
`307 Temporary Redirect` HTTP code.

# License

This project is licensed under the [MIT License](LICENSE.txt).

Copyright &copy; Alexander Nielsen, 2020.
