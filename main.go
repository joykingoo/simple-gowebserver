package main

func main() {

}

func Listen(network, address string) (Listener, error)

listener,error := net.Listen("tcp","8080")
if error != nil {
	//handle error
}
