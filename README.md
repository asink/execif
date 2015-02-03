# execif

A very tiny and simple application to execute a command
based on the existence of a file or directory.

### Building from source

Depends on:
  - Go (>= 1.2)
  - Git
  - Make
  - Linux only

```bash
$ git clone https://github.com/asink/execif.git
$ cd execif
$ make
$ sudo make install
```

### Usage

An example of running make in a directory that
does not initially exist:

```bash
$ execif /path/to/my/project make -C /path/to/my/project
```

#### Why is this useful?

I had some other processes that are responsible for mounting certain
files which don't initially exist. I needed
to do something with sed files as soon as they became available, so I
very quickly wrote this program to do so.

### License

MIT