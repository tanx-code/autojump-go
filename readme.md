# autojump-go

[autojump](https://github.com/wting/autojump) golang implementation. (Not Fully Implemented)

## Speed Testing

On my MacbookPro (Intel Core i7 2.2 GHz) , with a directory database of 192 entries.

```
➜  autojump-go git:(master) ✗ time autojump-go api
/Users/xiaotan/Work/brm-api
autojump-go api  0.00s user 0.01s system 72% cpu 0.010 total
➜  autojump-go git:(master) ✗ time ~/.autojump/bin/autojump api
/Users/xiaotan/Work/brm-api
~/.autojump/bin/autojump api  0.13s user 0.14s system 82% cpu 0.327 total
```

Feels like just typed the shell command `cd`.

## Install

If you use zsh and has golang environment prepared, you can run these commands to install `autojump-go`.

```
go get -u github/tvytlx/autojump-go

echo "source $GOPATH/src/github.com/autojump-go/autojump-go.zsh" >> ~/.zshrc
```

## TODO

* more flag supports, it only supports `--add` so far
* more test cases

Pull requests welcome
