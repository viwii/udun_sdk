package sdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type UdunClient struct {
	gateway            string //网关
	merchantId         string //商户编号
	merchantKey        string //商户key
	defaultCallBackUrl string // 默认回调地址
}

/**
 * 创建币种地址，别名和钱包编号默认，回调地址使用统一配置
 *
 * @param mainCoinType 主币种编号,使用获取商户币种信息接口
 * @return 地址
 */
func (uc *UdunClient) CreateAddressWithMainCoinType(mainCoinType string) (*Address, error) {
	return uc._createAddress(mainCoinType, "", "", uc.defaultCallBackUrl)
}

/**
 * 创建币种地址，别名和钱包编号自定义，回调地址使用统一配置
 *
 * @param mainCoinType 主币种编号,使用获取商户币种信息接口
 * @param alias        地址别名
 * @param walletId     钱包编号
 * @return 地址
 */
func (uc *UdunClient) CreateAddress(mainCoinType string, alias string, walletId string) (*Address, error) {
	return uc._createAddress(mainCoinType, alias, walletId, uc.defaultCallBackUrl)
}

/**
 * 创建币种地址，别名和钱包编号自定义，回调地址自定义
 *
 * @param mainCoinType 主币种编号,使用获取商户币种信息接口
 * @param alias        地址别名
 * @param walletId     钱包编号
 * @param callUrl      回调地址
 * @return 地址
 */
func (uc *UdunClient) _createAddress(mainCoinType string, alias string, walletId string, callUrl string) (*Address, error) {
	mp := make(map[string]string)
	mp["merchantId"] = uc.merchantId
	mp["coinType"] = mainCoinType
	mp["callUrl"] = callUrl
	mp["walletId"] = walletId
	mp["alias"] = alias
	data, _ := json.Marshal(mp)

	var retMsg ResultMsg
	retMsg.Data = &Address{}
	err := UdunPost(uc.gateway, uc.merchantKey, CREATE_ADDRESS, "["+string(data)+"]", &retMsg)
	if err != nil {
		return nil, err
	}

	if retMsg.Code != http.StatusOK {
		return nil, errors.New("error status")
	}

	if retMsg.Data != nil {
		return retMsg.Data.(*Address), nil
	}

	return nil, errors.New("error formate")
}

/**
 * 提币，回调地址自定义
 *
 * @param address      提币地址
 * @param amount       提币数量
 * @param mainCoinType 主币种编号,使用获取商户币种信息接口
 * @param coinType     子币种编号,使用获取商户币种信息接口
 * @param businessId   业务编号，必须保证该字段在系统内唯一，如果重复，则该笔提币钱包将不会进行接收
 * @param memo         备注,XRP和EOS，这两种币的提币申请该字段可选，其他类型币种不填
 * @return 返回信息
 */
func (uc *UdunClient) Withdraw(address string, amount float64, mainCoinType, coinType, businessId, memo string) (*ResultMsg, error) {
	return uc._withdraw(address, amount, mainCoinType, coinType, businessId, memo, uc.defaultCallBackUrl)
}

func (uc *UdunClient) _withdraw(address string, amount float64, mainCoinType string, coinType string, businessId string, memo string, callUrl string) (*ResultMsg, error) {
	params := make(map[string]string)
	params["address"] = address
	params["amount"] = strconv.FormatFloat(amount, 'f', 10, 64)
	params["merchantId"] = uc.merchantId
	params["mainCoinType"] = mainCoinType
	params["coinType"] = coinType
	params["callUrl"] = callUrl
	params["businessId"] = businessId
	params["memo"] = memo

	data, _ := json.Marshal(params)
	var retMsg ResultMsg
	err := UdunPost(uc.gateway, uc.merchantKey, WITHDRAW, "["+string(data)+"]", &retMsg)
	if err != nil {
		return nil, err
	}

	return &retMsg, nil
}

/**
 * 检验地址合法性
 *
 * @param mainCoinType 主币种编号,使用获取商户币种信息接口
 * @param address      币种地址
 * @return 是否合法
 */
func (uc *UdunClient) CheckAddress(mainCoinType string, address string) bool {
	params := make(map[string]string)
	params["merchantId"] = uc.merchantId
	params["mainCoinType"] = mainCoinType
	params["address"] = address

	data, _ := json.Marshal(params)
	var retMsg ResultMsg
	err := UdunPost(uc.gateway, uc.merchantKey, CHECK_ADDRESS, "["+string(data)+"]", &retMsg)
	if err != nil {
		return false
	}

	return retMsg.Code == http.StatusOK
}

func (uc *UdunClient) ListSupportCoin(showBalance bool) []Coin {
	params := make(map[string]interface{})
	params["merchantId"] = uc.merchantId
	if showBalance {
		params["showBalance"] = true
	} else {
		params["showBalance"] = false
	}
	var coins []Coin

	data, _ := json.Marshal(params)

	var retMsg ResultMsg
	retMsg.Data = &coins
	err := UdunPost(uc.gateway, uc.merchantKey, SUPPORT_COIN, string(data), &retMsg)
	if err != nil {
		return coins
	}

	if retMsg.Code != http.StatusOK {
		return coins
	}

	//err = json.Unmarshal([]byte(retMsg.Data), &coins)

	if err != nil {
		return coins
	}
	return coins
}

func NewUdunClient(gateway string, merchantId string, merchantKey string, defaultCallBackUrl string) *UdunClient {
	return &UdunClient{
		gateway:            gateway,
		merchantId:         merchantId,
		merchantKey:        merchantKey,
		defaultCallBackUrl: defaultCallBackUrl,
	}
}
