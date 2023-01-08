## 사용법

[Infura Api Key 발급하기](https://medium.com/jelly-market/how-to-get-infura-api-key-e7d552dd396f)


```bash
$ git clone https://github.com/wlswo/Subscribing-to-Ethereum-Event-Logs-in-GoLang.git
```

```bash
$ cd config/
$ touch config.toml
```

### config.toml 파일 작성

```
[Goerli]
apikey = "wss://goerli.infura.io/ws/v3/{발급한 Api Key}"

[log]
level = "debug" # debug or info
fpath = "./logs/go-loger" # 로그가 생성될 경로 : ./logs, 로그파일명 go-loger_xxx.log
msize = 2000    # 2g : megabytes
mage = 7        # 7days
mbackup = 5     # number of log files

[ContractAddress]
ca = "{감지할 스마트 컨트랙트 주소}"
```

```bash 
$ cd ../
$ go run main.go
```
