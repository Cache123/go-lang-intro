package main

import "fmt"

type car struct {
  gas_pedal uint16 //min to max 65535
  break_pedal uint16
  steerng_wheel int16 // -32k - 32k
  top_speed_kmh  float64
}

func main() {
  a_car := car{gas_pedal: 22341,
              break_pedal: 0,
              steerng_wheel: 12561,
              top_speed_kmh: 225.0}
  fmt.Println(a_car.gas_pedal)
}
