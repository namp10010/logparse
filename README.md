# logparse
a no-brainer json log parser

### install

```bash
  go install github.com/namp10010/logparse
```

copy the `config.yaml` to `~/.logparse`

### usage

```bash
  echo '{"message":"hello"}' | logparse
```

### todo

* make the config into args and easier to use, remove the config file
* enable color output
  * RED - error
  * YELLOW - warn
  * GREEN - info
  * DEBUG - blue