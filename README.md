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

### color output

to enable color set this config
```yaml
  color: true
```

color setting can be found in [main.go](main.go)

### todo
* make the config into args and easier to use, remove the config file