package main

import (
	"fmt"
	"math/big"
)

/*
# 拡張ユークリッドの互除法
# 引数:ax+by = c の a,b
# 戻値:リスト(c,x,y)
*/
func Egcd(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if a.Cmp(big.NewInt(0)) == 0 {
		return b, big.NewInt(0), big.NewInt(1)
	} else {
		g, x, y := Egcd(big.NewInt(0).Mod(b, a), a)
		//Goではintとintの商はint
		return g, big.NewInt(0).Sub(y,  big.NewInt(0).Mul( big.NewInt(0).Div(b, a), x) ), x
	}
}

func Gcd(a, b *big.Int) *big.Int {
	for b.Cmp(big.NewInt(0)) == 1 {
		var newB *big.Int
        newB.Mod(a, b)
		a.Add(b, big.NewInt(0))
        b.Add(newB, big.NewInt(0))
	}
    return a
}

func Lcm(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div( big.NewInt(0).Mul(a,b), Gcd(a, b) )
}

/*
#鍵生成
*/
func Generate_keys(p, q, e *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int) {
	var N *big.Int
	N.Mul(p, q)
	var L *big.Int
	L.Add(big.NewInt(0), Lcm(big.NewInt(0).Sub(p, big.NewInt(1)), big.NewInt(0).Sub(q, big.NewInt(1))))
    
    var  x *big.Int
	_, x, _ = Egcd(e, L)

	var d *big.Int
	d.Mod(x, L)

	//publick_key, private_key
	return e, N, d, N
}

/*
#暗号化
*/
func Encrypt(plain_text string, e *big.Int, N *big.Int) []*big.Int {
	plain_bytes := []byte(plain_text)
	var plain_integers []*big.Int
	for _, item := range plain_bytes {
		plain_integers = append(plain_integers, big.NewInt(int64(item)))
	}

	var encrypted_integers []*big.Int
	for _, item := range plain_integers {
		var newEncryptedInt *big.Int
        newEncryptedInt = big.NewInt(0)
		newEncryptedInt.Exp(item, e, N)
		encrypted_integers = append(encrypted_integers, newEncryptedInt)
	}

	return encrypted_integers
}

/*
// Bytes2str converts []byte to string("00 00 00 00 00 00 00 00")
func Bytes2str(bytes ...byte) string {
    strs := []string{}
    for _, b := range bytes {
        strs = append(strs, fmt.Sprintf("%02x", b))
    }
    return strings.Join(strs, " ")
}
*/

/*
#復号
*/
func Decrypt(encrypted_integers []*big.Int, d *big.Int, N *big.Int) string {
	var plain_integers []*big.Int
	for _, item := range encrypted_integers {
		var newPlainInt *big.Int
        newPlainInt = big.NewInt(0)
		newPlainInt.Exp(item, d, N)
		plain_integers = append(plain_integers, newPlainInt)
	}

	var plain_bytes []byte
	for _, item := range plain_integers {
		plain_bytes = append(plain_bytes, byte(item.Int64()))
	}

	return string(plain_bytes)
	/*
		    plain_bytes = make([]byte, len(encrypted_bytes) )
		    copy(encrypted_bytes, plain_bytes)
			for ind, item = range encrypted_bytes{
				plain_bytes.Exp(encrypted_bytes[ind], d, N )
			}

			plain_text = Bytes2str(plain_bytes)
			return plain_text
	*/
}

func testmain() {
	p := big.NewInt(29)
	q := big.NewInt(103)
	e := big.NewInt(13)
    
    var plain string
	plain = "FLAG{hello}"
    var N, d *big.Int
	e, N, d, N = Generate_keys(p, q, e)

    var encrypted_integers []*big.Int
	encrypted_integers = Encrypt(plain, e, N)
	fmt.Print(encrypted_integers)

	var decrypted_text string
    decrypted_text = Decrypt(encrypted_integers, d, N)
	fmt.Print(decrypted_text)
}

func main() {
	testmain()
}

/*
#############
#! -*- coding:utf-8 -*-

def Egcd(a, b):
    if a == 0:
        return (b, 0, 1)
    else:
        g, x, y = Egcd(b % a, a)
        return (g, y - (b // a) * x, x)

#最大公約数
def gcd(a,b):
    while b > 0:
        a, b = b, a % b
    return a

#最小公倍数
def Lcm(a, b):
    #整数割り算にしないとfloatがオーバーフローするとかでてとまる
    return a * b // gcd(a, b)

#(借り物)
#def modinv(a, m):
#    g, x, y = Egcd(a, m)
#    if g != 1:
#        raise Exception('No modular inverse')
#    return x%m

#鍵生成
def Generate_keys(p, q, e=65537):
    N = p * q
    L = Lcm(p - 1, q - 1)
    #もしくは、
    #L = (p-1)*(q-1)//gcd(p-1, q-1)

    c, x, y = Egcd(e, L)
    d = x % L
    #こちらでもよい?
    #d = inverse(e, L)
    #d = modinv(e, (p-1)*(q-1))

    #publick_key, private_key
    return (e, N), (d, N)

#暗号化
def Encrypt(plain_text, public_key):
    e, N = public_key
    plain_bytes = plain_text.encode("UTF-8")
    plain_integer = int.from_bytes(plain_bytes, 'big')
    encrypted_integer = pow(plain_integer, e, N)
    encrypted_bytes = encrypted_integer.to_bytes((encrypted_integer.bit_length() // 8) + 1, 'big')

    return encrypted_bytes


#復号
def Decrypt(encrypted_bytes, private_key):
    d, N = private_key
    encrypted_integer = int.from_bytes(encrypted_bytes, 'big')
    plain_integer = pow(encrypted_integer, d, N)
    plain_bytes = plain_integer.to_bytes((plain_integer.bit_length() // 8) + 1, byteorder='big')
    plain_text = plain_bytes.decode(encoding='UTF-8',errors='strict')
    return(plain_text)


#--------------------------------------


p = 54311
q = 158304142767773473275973624083670689370769915077762416888835511454118432478825486829242855992134819928313346652550326171670356302948444602468194484069516892927291240140200374848857608566129161693687407393820501709299228594296583862100570595789385365606706350802643746830710894411204232176703046334374939501731

plain = "FLAG{hello}"
pub_key, priv_key = Generate_keys(p, q)

encrypted_bytes = Encrypt(plain, pub_key )
print(encrypted_bytes)

decrypted_text = Decrypt(encrypted_bytes, priv_key)
print(decrypted_text)
*/
