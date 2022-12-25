# FullTextSearch

> A **Toy** Full Text Search

> No verify warning

## Run
Before you build it, you should modify `manifest/config/config.yml`
```yaml
server:
  #You can modify the address you like
  address:     ":8000"
  openapiPath: "/api.json"
  swaggerPath: "/swagger"

logger:
  level : "all"
  stdout: true

Jieba:
  # This option is decide Jieba mode
  useHMM: true
  # You can load Your dict(WIP)
  dictPath: ""

# Modify Address which is your Redis server
redis:
  TextRecord:
    address: 1.1.1.1
    db: 1
  IndexRecord:
    address: 1.1.1.1
    db: 2
```

And use this command to run it(By docker)
```shell
make run
```

Or use this command to build it
```shell
make cli.install
gf build
```
