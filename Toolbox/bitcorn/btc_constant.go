package bitcorn

/*
【火币交易费率】
当前火币交易的费率为0.2% 如果是买入 则扣除0.2%个btc
如果是卖出则扣除0.2%的美元
*/
const HUOBI_TRADE_FEE = 0.002

/*
充值美元时 美元的一般最高价格
*/
const TOPUP_FEE_RATE = 6.48

/*
从美元提现人民币时 一般最低价格
*/
const WITHDROW_FEE_RATE = 6.38

// 变量区
/**
美元兑换人民币
*/
var Usd2CnyExchangeRate = 6.4552

/*
比特币当前市价 美元
*/
var BtcCurrentDollarFee = 52000.0
