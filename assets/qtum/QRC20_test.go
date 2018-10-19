package qtum

import (
	"testing"
	"encoding/hex"
	"strconv"
	"github.com/shopspring/decimal"
	"github.com/blocktree/OpenWallet/common"
)


func Test_addressTo32bytesArg(t *testing.T) {
	address := "qP1VPw7RYm5qRuqcAvtiZ1cpurQpVWREu8"

	to32bytesArg, err := AddressTo32bytesArg(address)
	if err != nil {
		t.Errorf("To32bytesArg failed unexpected error: %v\n", err)
	}else {
		t.Logf("To32bytesArg success.")
	}

	t.Logf("This is to32bytesArg string for you to use: %s\n", hex.EncodeToString(to32bytesArg))
}


func Test_getUnspentByAddress(t *testing.T) {
	contractAddress := "91a6081095ef860d28874c9db613e7a4107b0281"
	address := "qdphfFinfJutJFvtnr2UaCwNAMxC3HbVxa"

	QRC20Utox, err := tw.GetUnspentByAddress(contractAddress, address)
	if err != nil {
		t.Errorf("GetUnspentByAddress failed unexpected error: %v\n", err)
	}

	sotashiUnspent, _ := strconv.ParseInt(QRC20Utox.Output,16,64)
	t.Logf("sotashiUnspent: %d\n", sotashiUnspent)

	sotashiUnspentDecimal, _ := decimal.NewFromString(common.NewString(sotashiUnspent).String())
	unspent := sotashiUnspentDecimal.Div(coinDecimal)

	if err != nil {
		t.Errorf("strconv.ParseInt failed unexpected error: %v\n", err)
	}else {
		t.Logf("QRC20Unspent %s: %s = %v\n", QRC20Utox.Address, address, unspent)
	}
}

func Test_AmountTo32bytesArg(t *testing.T){
	var amount int64= 100000000
	bytesArg, err := AmountTo32bytesArg(amount)
	if err != nil {
		t.Errorf("strconv.ParseInt failed unexpected error: %v\n", err)
	}else {
		t.Logf("hexAmount = %s\n", bytesArg)
	}
}

func Test_QRC20Transfer(t *testing.T) {
	contractAddress := "91a6081095ef860d28874c9db613e7a4107b0281"
	from := "qVT4jAoQDJ6E4FbjW1HPcwgXuF2ZdM2CAP"
	to := "qdphfFinfJutJFvtnr2UaCwNAMxC3HbVxa"
	gasPrice := "0.00000040"
	var gasLimit int64 = 250000
	var amount decimal.Decimal = decimal.NewFromFloat(9.999999)

	result, err := tw.QRC20Transfer(contractAddress, from, to, gasPrice, amount, gasLimit)
	if err != nil {
		t.Errorf("QRC20Transfer failed unexpected error: %v\n", err)
	}else {
		t.Logf("QRC20Transfer = %s\n", result)
	}
}