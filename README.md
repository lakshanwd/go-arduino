![Gobot](https://matt.aimonetti.net/images/gobotio.png)

# go-arduino

### Control your `arduino` with `golang`
As long as `arduino` is not supporting `multi threading`, we can use [gobot](https://gobot.io/) to perform `IO` operations and handle multi threading functionality with golang.

### prerequisites
* go
* arduino

### setup
* Install `Firmata`
    * Remove all attached components from `Arduino UNO`
    * Open `Arduino IDE`
    * Go to `File` > `Examples` > `Firmata` > `StandardFirmata`
* Setup Circuit
    * Create following Circuit ![Image](../master/circuit.svg?raw=true)
* execute following
    * `go run arduino.go`
