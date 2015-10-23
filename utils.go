package sssaas

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
)

var prime *big.Int

func random() *big.Int {
	result := big.NewInt(0).Set(prime)
	result = result.Sub(result, big.NewInt(1))
	result, _ = rand.Int(rand.Reader, result)
	return result
}

func splitByteToInt(secret []byte) []*big.Int {
	count := int(math.Ceil(float64(len(secret)) / 32))

	var result []*big.Int = make([]*big.Int, count)
	for i := range result {
		data := ""
		if (i+1)*32 < len(secret) {
			data = hex.EncodeToString(secret[i*32 : (i+1)*32])
		} else {
			tmp := make([]byte, 32)
			copy(tmp, secret[i*32:])
			data = hex.EncodeToString(tmp)
		}

		result[i], _ = big.NewInt(0).SetString(data, 16)
	}

	return result
}

func mergeIntToByte(secret []*big.Int) []byte {
	var result []byte
	for i := range secret {
		tmp, _ := hex.DecodeString(fmt.Sprintf("%x", secret[i]))
		result = append(result, tmp...)
	}

	return result
}

func evaluatePolynomial(polynomial []*big.Int, value *big.Int) *big.Int {
	var result *big.Int = big.NewInt(0).Set(polynomial[0])

	for s := range polynomial[1:] {
		tmp := big.NewInt(0).Set(value)
		tmp = tmp.Exp(tmp, big.NewInt(int64(s)+1), prime)
		tmp = tmp.Mul(tmp, polynomial[s+1])
		result = result.Add(result, tmp)
		result = result.Mod(result, prime)
	}

	return result
}

func inNumbers(numbers []*big.Int, value *big.Int) bool {
	for n := range numbers {
		if numbers[n].Cmp(value) == 0 {
			return true
		}
	}

	return false
}

func toBase64(number *big.Int) string {
	hexdata := fmt.Sprintf("%x", number)
	for i := 0; len(hexdata) < 64; i++ {
		hexdata = "0" + hexdata
	}
	bytedata, success := hex.DecodeString(hexdata)
	if success != nil {
		fmt.Println("Error!")
		fmt.Println("hexdata: ", hexdata)
		fmt.Println("bytedata: ", bytedata)
		fmt.Println(success)
	}
	return base64.URLEncoding.EncodeToString(bytedata)
}

func fromBase64(number string) *big.Int {
	bytedata, _ := base64.URLEncoding.DecodeString(number)
	hexdata := hex.EncodeToString(bytedata)
	result, _ := big.NewInt(0).SetString(hexdata, 16)
	return result
}

func modInverse(number *big.Int) *big.Int {
	copy := big.NewInt(0).Set(number)
	copy = copy.Mod(copy, prime)
	pcopy := big.NewInt(0).Set(prime)
	x := big.NewInt(0)
	y := big.NewInt(0)
	negative := false

	if copy.Cmp(big.NewInt(0)) == -1 {
		negative = true
		copy = copy.Mul(copy, big.NewInt(-1))
	}

	gcd := copy.GCD(x, y, pcopy, copy)

	if gcd.Cmp(big.NewInt(1)) != 0 {
		fmt.Println("Broken GCD!")
	}

	result := big.NewInt(0).Set(prime)

	if negative {
		y = y.Mul(y, big.NewInt(-1))
	}

	result = result.Add(result, y)
	result = result.Mod(result, prime)
	return result
}