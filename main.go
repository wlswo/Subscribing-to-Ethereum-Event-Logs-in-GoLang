package main

import (
	"context"
	"flag"
	"log"

	conf "goEthSub/config"
	logger "goEthSub/logger"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	var configFlag = flag.String("config", "./config/config.toml", "toml file to use for configuration")
	flag.Parse()
	cf := conf.NewConfig(*configFlag)
	/* 로그 설정 */
	logger.InitLogger(cf.Log.Fpath)

	/* 1. 이벤트 로그를 구독하기 위해 webSocket 지원 Ethereum 클라이언트에 다이얼을 연결 */
	client, err := ethclient.Dial(cf.Goerli.ApiKey)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Transaction Daemon Server Start")
	}

	/* 2. 필터 쿼리를 생성 , 이벤트를 구독할 컨트랙트의 주소를 입력 */
	contractAddress := common.HexToAddress(cf.ContractAddress.Ca)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	/* 3. 이벤트를 수신하는 방식은 Go 채널을 이용 */
	logs := make(chan types.Log)

	/* 4. SubscribeFilterLogs 쿼리 옵션과 출력 채널을 받는 클라이언트에서 호출하여 구독 */
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	/* 5. 새 로그 이벤트를 읽기위해 select문으로 연속 루프를 설정 */
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			log.Printf("%+v", vLog)
		}
	}
}
