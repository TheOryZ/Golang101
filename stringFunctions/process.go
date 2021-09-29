package stringfunctions

//alias
import (
	"fmt"
	s "strings"
)

func Demo() {
	name := "John"
	fmt.Println(s.Count(name, "h"))
	fmt.Println(s.Contains(name, "n"))
	fmt.Println(s.Index(name, "n")) //first index. If not exist return -1
	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))
	fmt.Println(s.HasPrefix(name, "Jo"))
	fmt.Println(s.HasSuffix(name, "hn"))
	letters := []string{"j", "o", "h", "n"}
	fmt.Println(s.Join(letters, "-"))
	fmt.Println(s.Replace(name, "J", "j", -1)) //-1 means replace for all.
	fmt.Println(s.Split(name, "o"))
	fmt.Println(s.Repeat(name, 5))

}
