package sdk

const CREATE_ADDRESS = "/mch/address/create"
const WITHDRAW = "/mch/withdraw"
const TRANSACTION = "/mch/transaction"
const AUTO_WITHDRAW = "/mch/withdraw/proxypay"
const SUPPORT_COIN = "/mch/support-coins"
const CHECK_PROXY = "/mch/check-proxy"
const CHECK_ADDRESS = "/mch/check/address"
const CREATE_BATCH_ADDRESS = "/mch/address/create/batch"

type Address struct {
	Address  string `json:"address"`
	CoinType int    `json:"coinType"`
}

type Coin struct {
	Name         string  `json:"name"`
	Symbol       string  `json:"symbol"`
	MainCoinType string  `json:"mainCoinType"`
	CoinType     string  `json:"coinType"`
	Decimals     string  `json:"decimals"`
	TokenStatus  int     `json:"tokenStatus"`
	MainSymbol   string  `json:"mainSymbol"`
	Balance      float64 `json:"balance"`
}

type ResultMsg struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Trade struct {
	TxId         string  `json:"txId"`         //交易Id
	TradeId      string  `json:"tradeId"`      //交易流水号
	Address      string  `json:"address"`      //交易地址
	MainCoinType string  `json:"mainCoinType"` //币种类型
	CoinType     string  `json:"coinType"`     //代币类型，erc20为合约地址
	Amount       float64 `json:"amount"`       //交易金额
	TradeType    int     `json:"tradeType"`    //交易类型  1-充值 2-提款(转账)
	Status       int     `json:"status"`       //交易状态 0-待审核 1-成功 2-失败,充值无审核
	Fee          float64 `json:"fee"`          //旷工费
	Decimals     int     `json:"decimals"`
	BusinessId   string  `json:"businessId"` //提币申请单号
	Memo         string  `json:"memo"`       //备注
}
