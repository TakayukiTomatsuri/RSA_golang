package main

import (
	"fmt"
	"math/big"
)

/*
# Extended GCD
# args    : a,b   in (ax+by = c)
# ret val : c,x,y in (ax+by = c)
*/
func Egcd(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if a.Cmp(big.NewInt(0)) == 0 {
		return b, big.NewInt(0), big.NewInt(1)
	} else {
		g, x, y := Egcd(big.NewInt(0).Mod(b, a), a)
		//In golang, the quotient of int and int is int.
		return g, big.NewInt(0).Sub(y, big.NewInt(0).Mul(big.NewInt(0).Div(b, a), x)), x
	}
}

func Gcd(a, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) == 1 {
		newB := big.NewInt(0)
		newB.Mod(a, b)
		a.Add(b, big.NewInt(0))
		b.Add(newB, big.NewInt(0))
	}
	return a
}

func Lcm(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(big.NewInt(0).Mul(a, b), Gcd(a, b))
}

func Generate_keys(p, q, e *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int) {
	N := big.NewInt(0)
	N.Mul(p, q)
	L := big.NewInt(0)
	L.Add(big.NewInt(0), Lcm(big.NewInt(0).Sub(p, big.NewInt(1)), big.NewInt(0).Sub(q, big.NewInt(1))))

	x := big.NewInt(0)
	_, x, _ = Egcd(e, L)

	d := big.NewInt(0)
	d.Mod(x, L)

	// The return value N is duplicated so as to have the same format with Python3 version.
	//publick_key(e,N), private_key(d,N)
	return e, N, d, N
}

func Encrypt(plain_text string, e *big.Int, N *big.Int) []*big.Int {
	//Convert string to bytes slice, and convert it to big.Int slice.
	plain_bytes := []byte(plain_text)
	var plain_integers []*big.Int
	for _, item := range plain_bytes {
		plain_integers = append(plain_integers, big.NewInt(int64(item)))
	}

	//Encrypt big.Int slice
	var encrypted_integers []*big.Int
	for _, item := range plain_integers {
		newEncryptedInt := big.NewInt(0)
		newEncryptedInt.Exp(item, e, N)
		encrypted_integers = append(encrypted_integers, newEncryptedInt)
	}

	return encrypted_integers
}

func Decrypt(encrypted_integers []*big.Int, d *big.Int, N *big.Int) string {
	//Decrypt big.Int slice.
	var plain_integers []*big.Int
	for _, item := range encrypted_integers {
		newPlainInt := big.NewInt(0)
		newPlainInt.Exp(item, d, N)
		plain_integers = append(plain_integers, newPlainInt)
	}

	//Convert big.Int slice to bytes slice
	var plain_bytes []byte
	for _, item := range plain_integers {
		plain_bytes = append(plain_bytes, byte(item.Int64()))
	}

	return string(plain_bytes)
}

func main() {
	p := big.NewInt(29)
	q := big.NewInt(103)
	e := big.NewInt(13)

	var plain string
	plain = "FLAG{hello}"

	var N, d *big.Int
	e, N, d, N = Generate_keys(p, q, e)

	var encrypted_integers []*big.Int
	encrypted_integers = Encrypt(plain, e, N)
	fmt.Println(encrypted_integers)

	var decrypted_text string
	decrypted_text = Decrypt(encrypted_integers, d, N)
	fmt.Println(decrypted_text)
}
