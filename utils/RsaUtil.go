package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
)

//生成RSA私钥和公钥，保存到文件中
// bits 证书大小
func GenerateRSAKey(bits int) {
	//GenerateKey函数使用随机数据生成器random生成一对具有指定字位数的RSA密钥
	//Reader是一个全局、共享的密码用强随机数生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}
	//保存私钥
	//通过x509标准将得到的ras私钥序列化为ASN.1 的 DER编码字符串
	X509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//使用pem格式对x509输出的内容进行编码
	//创建文件保存私钥
	privateFile, err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()
	//构建一个pem.Block结构体对象
	privateBlock := pem.Block{Type: "RSA Private Key", Bytes: X509PrivateKey}
	//将数据保存到文件
	pem.Encode(privateFile, &privateBlock)
	privateKeyData := pem.EncodeToMemory(&privateBlock)
	fmt.Println("privateKey:", string(privateKeyData))

	//保存公钥
	//获取公钥的数据
	publicKey := privateKey.PublicKey
	//X509对公钥编码
	X509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		panic(err)
	}
	//pem格式编码
	//创建用于保存公钥的文件
	publicFile, err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()
	//创建一个pem.Block结构体对象
	publicBlock := pem.Block{Type: "RSA Public Key", Bytes: X509PublicKey}
	//保存到文件
	pem.Encode(publicFile, &publicBlock)

	publiceKeyData := pem.EncodeToMemory(&publicBlock)
	fmt.Println("publiceKey:", string(publiceKeyData))
}

// RSAEncrypt RSA加密
// plainText 要加密的数据
// path 公钥
func RSAEncrypt(text string, rsaPublicKey string) string {
	if !strings.HasPrefix(rsaPublicKey, "-----BEGIN RSA Public Key-----") {
		rsaPublicKey = "-----BEGIN RSA Public Key-----\n" + rsaPublicKey + "\n-----END RSA Public Key-----"
	}

	//pem解码
	block, _ := pem.Decode([]byte(rsaPublicKey))
	//x509解码

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//类型断言
	publicKey := publicKeyInterface.(*rsa.PublicKey)
	//对明文进行加密
	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(text))
	if err != nil {
		panic(err)
	}
	//返回密文
	return base64.StdEncoding.EncodeToString(cipherText)
}

// RSADecrypt RSA解密
// cipherText 需要解密的byte数据
// rsaPrivateKey 私钥
func RSADecrypt(text string, rsaPrivateKey string) string {
	code, _ := base64.StdEncoding.DecodeString(text)

	if !strings.HasPrefix(rsaPrivateKey, "-----BEGIN RSA Private Key-----") {
		rsaPrivateKey = "-----BEGIN RSA Private Key-----\n" + rsaPrivateKey + "\n-----END RSA Private Key-----"
	}
	//pem解码
	block, _ := pem.Decode([]byte(rsaPrivateKey))
	//X509解码
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	//对密文进行解密
	plainText, _ := rsa.DecryptPKCS1v15(rand.Reader, privateKey, code)
	//返回明文
	return string(plainText)
}
