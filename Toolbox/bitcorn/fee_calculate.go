package bitcorn

/**
[一笔钱流经火币还剩多少]
向比特币平台(火币)充一笔钱，没有做任何交易，然后提现到银行，还剩多少钱
*/
func TopupThenWithdrowLeft(cnMoney float64) (float64, float64) {
	// 每充值一美元 火币收费 0.1美元
	oneDoFee := TOPUP_FEE_RATE - WITHDROW_FEE_RATE
	usdMoney := cnMoney / Usd2CnyExchangeRate
	totalFee := usdMoney * oneDoFee
	rate := totalFee / cnMoney
	left := cnMoney - totalFee
	return left, rate
}

/**
【一定的本金低吸高抛能赚多少钱】
买入扣0.2%比特币
卖出扣0.2%美刀
p1: 单笔交易本金
p2: 高抛时涨幅 如100美元
*/
func BuyOnLowAndSellOnHighEarned(cnMoney float64, usdRise float64) (earned, fee, rate float64) {
	// 买入手续费
	buyFee := cnMoney * HUOBI_TRADE_FEE
	left := cnMoney - buyFee
	// 涨的百分比
	raseRate := usdRise / BtcCurrentDollarFee
	raseCnMoney := raseRate * left
	left += raseCnMoney
	soldFee := left * HUOBI_TRADE_FEE
	left -= soldFee
	// 赚多少钱
	earned = left - cnMoney
	fee = buyFee + soldFee
	rate = fee / cnMoney
	return
}

/*
以${cnMoney}为本金，以【1000美元差价高抛低吸】交易n次 最终赚多少钱，收益率多少
${raiseUsdMoney} 表示1000美元，可自己设定
*/
func EarnedBy1000DifferForN(cnMoney float64, raiseUsdMoney float64, n int) (earned, earnRate float64) {
	var totalEarn float64
	for i := 0; i < n; i++ {
		highEarned, _, _ := BuyOnLowAndSellOnHighEarned(cnMoney, raiseUsdMoney)
		totalEarn += highEarned
	}
	earned = totalEarn
	earnRate = totalEarn / cnMoney
	return earned, earnRate
}
