package main

import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "strconv"
)


func convert(value float64, tani string) string {
  switch tani {
  case "m":
    return fmt.Sprintf("%gft.", value / 0.3048)
  case "ft.":
    return fmt.Sprintf("%gm", value * 0.3048)
  case "Kg":
    return fmt.Sprintf("%glb", value * 0.454)
  case "lb":
    return fmt.Sprintf("%gKg", value / 0.454)
  case "C":
    return fmt.Sprintf("%gF", value * 9 / 5 + 32)
  case "F":
    return fmt.Sprintf("%gC", (value - 32 ) * 5 / 9)
  default:
    return "Invalid unit type"
  }
}

func main() {
  regex := regexp.MustCompile(`(-?\d+)([A-Za-z]+)`)
  if len(os.Args[1:]) > 0 {
    for _, arg := range os.Args[1:] {
      parsed := regex.FindStringSubmatch(arg)
      value, _ := strconv.ParseFloat(parsed[1], 64)
      tani := parsed[2]
      fmt.Println(convert(value, tani))
    }
  } else {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
      parsed := regex.FindStringSubmatch(input.Text())
      value, _ := strconv.ParseFloat(parsed[1], 64)
      tani := parsed[2]
      fmt.Println(convert(value, tani))
    }
  }
}
