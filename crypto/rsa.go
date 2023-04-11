package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"os"
)

func GenerateRsaKey(bits int) error {
	// 1. 生成私钥文件
	// GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	// 参数1: Reader是一个全局、共享的密码用强随机数生成器
	// 参数2: 秘钥的位数 - bit
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	// 2. MarshalPKCS1PrivateKey将rsa私钥序列化为ASN.1 PKCS#1 DER编码
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 3. Block代表PEM编码的结构, 对其进行设置
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	// 4. 创建文件
	privFile, err := os.Create("private.pem")
	if err != nil {
		return err
	}
	// 5. 使用pem编码, 并将数据写入文件中
	err = pem.Encode(privFile, &block)
	if err != nil {
		return err
	}
	// 6. 最后的时候关闭文件
	defer privFile.Close()

	// 7. 生成公钥文件
	publicKey := privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return err
	}
	block = pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: derPkix,
	}
	pubFile, err := os.Create("public.pem")
	if err != nil {
		return err
	}
	// 8. 编码公钥, 写入文件
	err = pem.Encode(pubFile, &block)
	if err != nil {
		panic(err)
	}
	defer pubFile.Close()

	return nil
}

// RSAEncrypt rsa加密
// src 要加密的数据
// 公钥文件的路径
func RSAEncrypt(src, filename []byte) []byte {
	// 1. 根据文件名将文件内容从文件中读出
	file, err := os.Open(string(filename))
	if err != nil {
		return nil
	}
	// 2. 读文件
	info, _ := file.Stat()
	allText := make([]byte, info.Size())
	file.Read(allText)
	// 3. 关闭文件
	file.Close()

	// 4. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(allText)
	if block == nil {
		return nil
	}
	// 5. 解析一个DER编码的公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	// 6. 公钥加密
	result, _ := rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	return result
}

// RSADecrypt rsa加密
// src 要解密的数据
// 私钥文件的路径
func RSADecrypt(src, filename []byte) []byte {
	// 1. 根据文件名将文件内容从文件中读出
	file, err := os.Open(string(filename))
	if err != nil {
		return nil
	}
	// 2. 读文件
	info, _ := file.Stat()
	allText := make([]byte, info.Size())
	file.Read(allText)
	// 3. 关闭文件
	file.Close()
	// 4. 从数据中查找到下一个PEM格式的块
	block, _ := pem.Decode(allText)
	// 5. 解析一个pem格式的私钥
	privateKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	// 6. 私钥解密
	result, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, src)

	return result
}

func Rsa2PriSign(signContent string, privateKey string, hash crypto.Hash) string {
	shaNew := hash.New()
	shaNew.Write([]byte(signContent))
	hashed := shaNew.Sum(nil)
	priKey, err := parsePrivateKey(privateKey)
	if err != nil {
		return ""
	}

	signature, err := rsa.SignPKCS1v15(rand.Reader, priKey, hash, hashed)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(signature)
}
func parsePrivateKey(privateKey string) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return nil, errors.New("私钥信息错误！")
	}
	priKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return priKey, nil
}

// RSA2公钥验证签名
func Rsa2PubCheckSign(signContent, sign, publicKey string, hash crypto.Hash) bool {
	hashed := sha256.Sum256([]byte(signContent))
	pubKey, err := parsePublicKey(publicKey)
	if err != nil {
		return false
	}
	sig, _ := base64.StdEncoding.DecodeString(sign)
	err = rsa.VerifyPKCS1v15(pubKey, hash, hashed[:], sig)
	if err != nil {
		return false
	}
	return true
}

// 解析公钥
func parsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("公钥信息错误！")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pubKey.(*rsa.PublicKey), nil
}
