package main

import "fmt"

type USB interface {
	Name() string
	Connector()
}

type PCConnector struct {
	name string;
}

func (pc PCConnector) Name() string  {
	return pc.name;
}

func (pc PCConnector) Connector()  {
	fmt.Println("connector",pc.name)
}

func main() {
	_usb :=  PCConnector{"pc"}
	_usb.Connector()
	DisConnector(_usb)

}

func DisConnector(conn USB)  {

	if pc, ok := conn.(PCConnector);ok{
		fmt.Println("if pc,ok := conn.(PCConnector)",pc.name)
	}
	switch c:=conn.(type) {
	case PCConnector:
		fmt.Println("switch ", c.name)
	default:
		fmt.Println("UnKown")
	}
	fmt.Println(conn.Name());
	fmt.Println("disconnector:",conn.Name())
}


