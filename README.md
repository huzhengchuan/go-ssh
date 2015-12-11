# Elfgate

Batch exec command on servers &amp; written by golang.

## Features

- Exec command on cluster servers.
- Server clusters support.
- Hosts support simple preg(just IPs).
- sftp supported, can upload file to batch servers(NOTE: supported file only).

## Getting started

Get the package, add the src to workspace path and build it.

But, first you should get golang ssh packages

```shell
$ git clone github.com/youngsn/go-ssh
$ go build -o ssh-logins $GOPATH/go-ssh/src/main.go
$ cp go-ssh/hosts.toml /etc/elfgate.conf
```

Usage:

Execute command just like below, the outputs will print on screen or > to file.

Now also sudo command are supported, but at first you should enter password if not config pasword.

- -c, config file location, default: /etc/elfgate.conf
- -t, command timeout, default: 0, no timeout
- -g, groups that execute commands, default: default
- -d, execute command, need specify.

```shell
$ ssh-logins -c $CONF -t $TIMEOUT -g $GROUP -d $CMD
```

If you want to upload file to batch server, just do like below. Very easy and now only supported file, not directory.

```shell
$ ssh-logins -c $CONF -g $GROUP -d "sftp $LOCAL_PATH $REMOTE_PATH"
```

## Config syntax
```toml
Username  = ""       # server username
Password  = ""       # login password, if don't config, you will enter through stdin
PublicKey = ""       # ssh public authorized key path, if using this, add here

# Group default
[groups.default]
Hosts     = [
    "127.0.0.[1-5]",        # simple preg support
    "127.0.0.[6-7]:233",
    "127.0.0.8",            # default port 22
    "127.0.0.9:25"          # port 25
]

# Group example
[groups.example]
Hosts     = [
    "127.0.0.2",
    "127.0.0.3:25"
]
```


## Third packages

Uses packages.

- [toml config parser](https://github.com/BurntSushi/toml) master
- [golang.org/x/crypto/ssh](https://github.com/golang/crypto)
- [github.com/pkg/sftp](https://github.com/pkg/sftp)

## TODO

Interactive cmds solution.

## Author

**TangYang**
<youngsn.tang@gmail.com>

## License

Released under the [MIT License](https://github.com/youngsn/go-ssh/blob/master/LICENSE).
