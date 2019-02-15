package main

import (
	"IsoMsg8583/isoFormatter"
	"bytes"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

func main() {
	//var msg Message
	amount := "20.00"
	if strings.Contains(amount, ".") {
		amount = strings.Replace(string(amount), ".", "", -2)
	} else {
		amount = amount + "00"
	}
	amount = PadLeft(amount, 12, "0")
	msg := isoFormatter.Message{}
	currentTime := time.Now()

	msg.MTI = "0200"
	msg.AddData(2, 255777630950)
	msg.AddData(3, "310000")
	msg.AddData(4, amount)
	msg.AddData(7, 303103707)
	msg.AddData(11, 133337)
	msg.AddData(13, 2015)
	msg.AddData(37, currentTime.Format("010215040512")[0:12])
	msg.AddData(49, "KES")
	msg.AddData(102, "01001030027089")

	isomsg := msg.Build()
	logISOMSG(msg)
	host := "192.168.12.28"
	port := "7003"

	fmt.Println(isomsg)
	len2 := len(isomsg)
	var buffer bytes.Buffer
	if len2 < 100 {
		buffer.WriteString("000")

	} else if len2 < 1000 {
		buffer.WriteString("00")
	} else {
		buffer.WriteString("0")

	}

	buffer.WriteString(strconv.Itoa(len2))
	buffer.WriteString(isomsg)
	requestISO := buffer.String()

	// connect to this socket
	conn, conerror := net.Dial("tcp", host+":"+port)
	if conerror != nil {
		fmt.Println(conerror)

	}
	conn.Write([]byte(requestISO))

	buff := make([]byte, 1024)
	response, resperror := conn.Read(buff)

	if resperror != nil {
		fmt.Println(resperror)

	}
	responses := buff[:response]
	fmt.Println("Receive: %s", responses)

	if responses == nil || len(responses) < 10 {
		fmt.Println(resperror)

	}
	//parse response
	results := fmt.Sprintf("%s", responses)
	fmt.Println("trim  ", results)
	parsed, parseError := isoFormatter.Parse(results[5:])

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
	logISOMSG(parsed)
}

func logISOMSG(res isoFormatter.Message) {
	k := 0
	for ; k <= 128; k++ {
		for _, f := range res.Fields {
			if k == f {
				fmt.Println(f, res.Data[f])
			}

		}
	}
}
func PadLeft(str string, length int, pad string) string {
	return times(pad, length-len(str)) + str
}

func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}
