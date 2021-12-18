package conf

import "fmt"
import "os"

func Check(c string) string {
	if _, err := os.Stat(c); err == nil {
	  return fmt.Sprintf("%s EXISTS\n", c)
	} else {
		return fmt.Sprintf("%s DOES NOT EXIST\n", c)
	}
  }