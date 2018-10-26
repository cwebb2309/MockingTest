package main

type bankDetail interface {
	getBalance(accNo string, sortCode string) (int, int)
}

type personalAccount struct {
}

func (p personalAccount) getBalance(accNo string, sortCode string) (int, int) {
	if accNo == "0001" {
		return 1200, 2700
	}
	return 1000, 2200
}

func main() {
	var account personalAccount

	bal, available := queryBalance(account, "0001", "01-01-01")
	println(bal)
	print(available)

}

func queryBalance(a bankDetail, accNo string, sortCode string) (int, int) {
	bal, avail := a.getBalance(accNo, sortCode)
	return bal, avail
}
