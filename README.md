## 사용법

[Infura Api Key 발급하기](https://medium.com/jelly-market/how-to-get-infura-api-key-e7d552dd396f)


```bash
$ git clone https://github.com/wlswo/Subscribing-to-Ethereum-Event-Logs-in-GoLang.git
```

```bash
$ cd config/
$ touch config.toml
```

### toml 파일 작성

```
[Goerli]
apikey = "wss://goerli.infura.io/ws/v3/{발급한 Api Key}"

[Log]
fpath = "./log/logfile.txt"

[ContractAddress]
ca = "{트랜잭션 발생 이벤트를 구독할 스마트 컨트랙트 주소}"
```

```bash 
$ cd ../
$ go run main.go
```
