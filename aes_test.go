package crypto4go

//import (
//	"testing"
//	"fmt"
//	"encoding/hex"
//)
//
//var plain = "use this method to release shared resources, save user data, invalidate timers, and store enough application state information to restore your application to its current state in case it is terminated later."
//var key = "11111111111111111111111111111111"
//var iv  = "1111111111111111"
//
//func Test_AES_CBC(t *testing.T) {
//	var a = []byte(plain)
//	var r, _ = AESCBCEncrypt(a, []byte(key), []byte(iv))
//	fmt.Println("AES CBC Encrypt: ", hex.EncodeToString(r))
//
//	r, _ = AESCBCDecrypt(r, []byte(key), []byte(iv))
//	fmt.Println("AES CBC Decrypt: ", string(r))
//}
//
//func Test_AES_CFB(t *testing.T) {
//	var a = []byte(plain)
//	var r, _ = AESCFBEncrypt(a, []byte(key), []byte(iv))
//	fmt.Println("AES CFB Encrypt: ", hex.EncodeToString(r))
//
//	r, _ = AESCFBDecrypt(r, []byte(key), []byte(iv))
//	fmt.Println("AES CFB Decrypt: ", string(r))
//}
