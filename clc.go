package main

import (
	"fmt"
   "os"
   "strconv"
	"github.com/shopspring/decimal"
)

func main() {
   usage()
   var choose int
   fmt.Scanln(&choose)
   fmt.Println("==============================")
   switch choose {
   case 1:
      calculateBuyPoint()
   case 2:
      calculateRate()
   default:
   }
}

func usage() {
   fmt.Printf("1 计算买点； 2 计算涨跌幅\n输入选择：")
}
func calculateBuyPoint() {
   var  rate,currnetPrice float64

   for {
      LOOP1: fmt.Printf("跌幅度(例如：-0.15)：")
         fmt.Scanln(&rate)
         if (rate == 1) {
            fmt.Printf("EXIT OUT!!!\n")
            os.Exit(0)
         }
         if (rate == 0.000000) {
            goto LOOP1
         }

      LOOP2: fmt.Printf("请输入价格(例如：1.1234)：")
         fmt.Scanln(&currnetPrice)
         if (currnetPrice == 0.000000) {
            goto LOOP2
         }
      fmt.Println()

   
      d1 := decimal.NewFromFloat(rate).Mul(decimal.NewFromFloat(currnetPrice))
      d2 := d1.Add(decimal.NewFromFloat(currnetPrice))
      fmt.Printf("买入价格：")
      d2Float64 := d2.InexactFloat64()
      value, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", d2Float64), 64)
      fmt.Println(value)
      fmt.Println("------------------------------")

   }
}

func calculateRate() {
   var  sPrice,ePrice float64

   for {
      LOOP1: fmt.Printf("请输入原值：")
         fmt.Scanln(&sPrice)
         if (sPrice == 1) {
            fmt.Printf("EXIT OUT!!!\n")
            os.Exit(0)
         }
         if (sPrice == 0.000000) {
            goto LOOP1
         }

      LOOP2: fmt.Printf("请输入现值：")
         fmt.Scanln(&ePrice)
         if (ePrice == 0.000000) {
            goto LOOP2
         }
      fmt.Println()
          
      d1 := decimal.NewFromFloat(ePrice).Div(decimal.NewFromFloat(sPrice))
      d2 := d1.Sub(decimal.NewFromFloat(1.0))

      fmt.Printf("涨跌幅：")
      d2Float64 := d2.InexactFloat64()
      value, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", d2Float64), 64)
      fmt.Printf("%.2f%%\n", value*100)
      //fmt.Println("%")
      fmt.Printf("-------------------------------\n")
   }
}
