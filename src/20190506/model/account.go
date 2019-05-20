package model

import "fmt"

type account struct {
	Name string
	Age int
	accountNo string
	pwd string
	balance float64
}

func NewAccount(accountNo string, pwd string, balance float64) *account  {
	if len(accountNo) < 6 || len(accountNo) > 10 {
		fmt.Println("账号长队不对...")
		return nil
	}

	if len(pwd) != 6 {
		fmt.Println("密码长度不对...")
		return nil
	}

	if balance < 20 {
		fmt.Println("余额数目不对")
		return nil
	}

	return &account{
		accountNo:accountNo,
		pwd:pwd,
		balance:balance,
	}
}

func (account *account) Deposite(money float64, pwd string) {

}