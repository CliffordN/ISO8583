# ISO8583
Go for iso8583

This is a go lang project that assist in processing ISO 8583 message standards.

 go get github.com/CliffordN/ISO8583
 
 # Example
 
 func main() {
	// Creates a new Iso8583 message
	msg := isoFormatter.Message{}
	
	msg.MTI = "0200"
	msg.AddData(2, 1234567)
	msg.AddData(3, "310000")
	msg.AddData(4, amount)
	msg.AddData(7, 303103707)
	msg.AddData(11, 133337)
	msg.AddData(13, 2015)
	msg.AddData(37, 010215040512)	
	msg.AddData(49, "404")
	msg.AddData(102, "12345678899")
	
	//build the message
	isomsg := msg.Build()
	
	//print ISO data elements
	isoFormatter.LogISOMSG(msg)
  
  }
  

