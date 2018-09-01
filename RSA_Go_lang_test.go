package main

import (
	"testing"
	"math/big"
)

func TestEgcd(t *testing.T) {
	var c, x, y *big.Int
	c, x, y = Egcd(big.NewInt(180), big.NewInt(150))

	if c.Cmp(big.NewInt(30)) != 0 {
		t.Fatal("failed test")
	}
	if x.Cmp(big.NewInt(1)) != 0 {
		t.Fatal("failed test")
	}
	if y.Cmp(big.NewInt(-1)) != 0 {
		t.Fatal("failed test")
	}
}

func TestGcd(t *testing.T) {
	ans := Gcd(big.NewInt(630), big.NewInt(300))
	if ans.Cmp(big.NewInt(30)) != 0 {
		t.Fatal("failed test")
	}
}

func TestLcm(t *testing.T) {
	ans := Lcm(big.NewInt(630), big.NewInt(300))
	if ans.Cmp(big.NewInt(6300)) != 0 {
		t.Fatal("failed test")
	}
}

func TestGenerate_keys(t *testing.T) {
	var e, N, d *big.Int
	e, N, d, N = Generate_keys(big.NewInt(29), big.NewInt(103), big.NewInt(13))
	if e.Cmp(big.NewInt(13)) != 0 {
		t.Fatal("failed test")
	}
	if N.Cmp(big.NewInt(2987)) != 0 {
		t.Fatal("failed test")
	}
	if d.Cmp(big.NewInt(769)) != 0 {
		t.Fatal("failed test")
	}
}

func TestDecrypt(t *testing.T) {

}

func TestEncrypt(t *testing.T) {

}
