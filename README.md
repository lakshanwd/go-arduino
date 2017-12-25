![Gobot](https://matt.aimonetti.net/images/gobotio.png)

# go-arduino

### Control your `arduino` with `golang`
As long as `arduino` is not supporting `multi threading`, we can use [gobot](https://gobot.io/) to perform `IO` operations and handle multi threading functionality with golang.

### prerequisites
* go
* arduino

### setup
Install `Firmata`
* Open `Arduino IDE`
* Go to `File` > `Examples` > `Firmata` > `StandardFirmata`

execute following     

    go run arduino.go
