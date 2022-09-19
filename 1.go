package kata

import (
  "strings"
  "strconv"
  "fmt"
)

func HighAndLow(in string) string {
  var min, max int

  numbers := strings.Fields(in)
  
  for i, number := range numbers {
    n, _ := strconv.Atoi(number)
    if i == 0 {
      max, min = n, n
    }
    
    if n > max {
      max = n
    }
    if n < min {
      min = n
    }
  }
  
  return fmt.Sprintf("%v %v", max, min)
}
