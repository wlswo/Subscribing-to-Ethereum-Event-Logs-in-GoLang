package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	conf "goEthSub/config"
	logger "goEthSub/logger"
	"goEthSub/models"
	"goEthSub/utils"

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
	logger.InitLogger(cf)

	/* model 초기화 */
	md, err := models.NewModel(cf.DB.Host)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(cf.Wemix.Url)
	if err != nil {
		log.Fatal(err)
	} else {
		logger.Info("Transaction Daemon Server Start")
		log.Println("Transaction Daemon Server Start")
	}

	/* 배포한 컨트랙트의 이벤트를 구독할 컨트랙트의 주소를 입력 */
	contractAddress := common.HexToAddress(cf.ContractAddress.Ca)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	/* 블록 감지 */
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	/* SubscribeFilterLogs 쿼리 옵션과 출력 채널을 받는 클라이언트에서 호출하여 구독 */
	logsContract := make(chan types.Log)
	subContract, err := client.SubscribeFilterLogs(context.Background(), query, logsContract)
	if err != nil {
		log.Fatal(err)
	}

	/* 새 로그 이벤트를 읽기위해 select문으로 연속 루프를 설정 */
	for {
		select {
		case err := <-subContract.Err():
			log.Fatal(err)
		case vLog := <-logsContract:
			TransactionLog, _ := utils.Log(vLog).MarshalJSON()
			fmt.Println("CA :", string(TransactionLog))
			logger.Event(string(TransactionLog))
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByNumber(context.Background(), header.Number)
			if err != nil {
				log.Fatal(err)
			}

			/* 블록 구조체 생성 */
			b, err := utils.BindingBlock(block)
			if err != nil {
				log.Fatal(err)
			}
			/* 트랜잭션 추출 */
			err = utils.GetTransactionsFromBlock(block.Transactions(), &b, block)
			if err != nil {
				log.Fatal(err)
			}

			/* 트랜잭션이 존재할 경우만 DB에 저장 */
			if len(b.Transactions) > 0 {
				if err := md.SaveBlock(&b); err != nil {
					log.Fatal(err)
				}
			}
		}
	}

}
