package main

type bankDetail interface {
	getBalance(accNo string) int
}

type personalAccount struct {
}

func (p personalAccount) getBalance(accNo string) int {
	if accNo == "0001" {
		return 1200
	}
	return 1000
}

func main() {
	var account personalAccount

	bal := queryBalance(account, "0001")
	println(bal)

}

func queryBalance(a bankDetail, accNo string) int {
	return a.getBalance(accNo)
}
