package main

import (
	"ISO8583/isoFormatter"
	"fmt"
	"time"
)

func main() {

	amount := "20.00"
	amount = isoFormatter.PadAmount(amount, 12, "0")
	currentDate := time.Now().Format("20060102")
	currentTime := time.Now().Format("150405")
	msg := isoFormatter.Message{}
	rrn := currentDate + currentTime

	msg.MTI = "0200"
	msg.AddData(2, 12345678)
	msg.AddData(3, "310000")
	msg.AddData(4, amount)
	msg.AddData(7, currentDate+currentTime)
	msg.AddData(11, 133337)
	msg.AddData(12, currentTime)
	msg.AddData(13, time.Now().Format("0102"))
	msg.AddData(37, rrn[2:])
	msg.AddData(49, "404")
	msg.AddData(102, "123456789")

	isomsg := msg.Build()
	fmt.Println(isomsg)
	isoFormatter.LogISOMSG(msg)

	parsed, parseError := isoFormatter.Parse(isomsg)

	if parseError != nil {
		fmt.Println(parseError.Error())

	}
	//get fields value
	field47 := parsed.Data[47].String()
	field39 := parsed.Data[39].String()
	//field38 := parsed.Data[38].String()

	fmt.Println("MTI:", parsed.MTI)
	fmt.Println("39:", field39)
	fmt.Println("47:", field47)
	isoFormatter.LogISOMSG(parsed)
}
func buildISO() {
	amount := "20.00"
	amount = isoFormatter.PadAmount(amount, 12, "0")
	currentDate := time.Now().Format("20060102")
	currentTime := time.Now().Format("150405")
	msg := isoFormatter.Message{}
	rrn := currentDate + currentTime

	msg.MTI = "0200"
	msg.AddData(2, 12345678)
	msg.AddData(3, "310000")
	msg.AddData(4, amount)
	msg.AddData(7, currentDate+currentTime)
	msg.AddData(11, 133337)
	msg.AddData(12, currentTime)
	msg.AddData(13, time.Now().Format("0102"))
	msg.AddData(37, rrn[2:])
	msg.AddData(49, "404")
	msg.AddData(102, "123456789")

	isomsg := msg.Build()
	fmt.Println(isomsg)
	isoFormatter.LogISOMSG(msg)

}
func parseISO() {

	msg := "0200f2280000080080000000000004000000071234567310000000000000200030310370713333720150102150405124041112345678899"
	parsed, parseError := isoFormatter.Parse(msg)

	if parseError != nil {
		fmt.Println(parseError.Error())

	}
	//get fields value
	field37 := parsed.Data[37].String()
	field102 := parsed.Data[102].String()

	fmt.Println("MTI:", parsed.MTI)
	fmt.Println("37:", field37)
	fmt.Println("102:", field102)
	isoFormatter.LogISOMSG(parsed)
}
