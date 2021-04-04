package main

import (
	"./bitcorn"
	"fmt"
)

func main() {
	fmt.Println("【比特币攻略】")
	bitcorn.Usd2CnyExchangeRate = 6.4552
	bitcorn.BtcCurrentDollarFee = 52795.19
	fmt.Println("    假定当前美元总的人民币汇率为：", bitcorn.Usd2CnyExchangeRate, "， BTC当前美元市价:", bitcorn.BtcCurrentDollarFee)
	fmt.Println("    所有火币费率按LV1默认交易费率")
	myTotalMoney := 2000.0
	left, rate := bitcorn.TopupThenWithdrowLeft(myTotalMoney)
	fmt.Print("【1.充值提现】充值", myTotalMoney, "后立马取出，还剩：", left, "   费率：百分之", rate*100, "\n")
	fmt.Println("【2.收益计算】幻想赚钱")
	fmt.Println("  2.0 一笔钱")
	fmt.Println("    现在，把本金6等分(每份约0.001个btc)，每一份在不同的水位买入，然后在高出买入位1000美元的高位抛出，来计算下这种操作的收益")

	oneTradeMonye := 0.001 * bitcorn.BtcCurrentDollarFee * bitcorn.Usd2CnyExchangeRate
	earned, fee, oneRate := bitcorn.BuyOnLowAndSellOnHighEarned(oneTradeMonye, 1000)
	oneEarned := earned
	fmt.Println("  2.1 有0.001比特币(当前约50美元)作为本金低吸高抛可以赚多少一次交易？", "本钱：", oneTradeMonye, " ，赚：", oneEarned, "收费：", fee, "费率：", oneRate)
	fmt.Println("      比特币市价越高，单笔1000涨幅低吸高抛赚的钱越少")

	nEarned, nearnRate := bitcorn.EarnedBy1000DifferForN(oneTradeMonye, 1000, 365)
	fmt.Println("  2.2 如果有人按这种方式平均每天赚1美金，1年后，他赚多少钱？", "本钱：", oneTradeMonye, " , 收益：", nEarned, "  ,收益率：", nearnRate)

}
