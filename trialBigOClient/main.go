package main
import (
	"github.com/go-kata-starter/trialBigOClient/trialBigOWithElse"
	"github.com/go-kata-starter/trialBigOClient/trialBigOWithoutElse"
)


func main() {
	trialBigOWithElse.GetListUpto10WithElse()
	trialBigOWithoutElse.GetListUpto10WithoutElse()
}