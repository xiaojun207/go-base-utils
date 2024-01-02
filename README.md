# go-base-utils
一个golang的常见功能函数的工具包，请大家star

## 项目地址
https://github.com/xiaojun207/go-base-utils

## min go sdk
1.18

## 这是一个golang的基本工具utils
零依赖，工具全部使用的是系统api

>包含：
* 1、类型转换工具单元，Float32ToString、Float64ToString、StrToFloat64、StrToInt、StrToBool、Random
* 2、struct、json、map的转换；
* 3、字符串截取基本单元，SubstrBetween，SubstrAfter，SubstrBefore；
* 4、http请求封装单元；
* 5、随机单元，包含随机密码（多种组合形式），如：RandomPassword(16, "mix")，范围随机：Random(60.0, 100.0)；
* 6、文件读写单元，写入文本文件；
* 7、DES加密，如：DESEncrypt，DESDecrypt；
* 8、AES加密，如：AESEncrypt，AESDecrypt；
* 9、RSA加、解密，如: RSAEncrypt("明文","公钥字符串")、RSADecrypt("密文", "私钥字符串")
* 10、日期单元，获取各种格式的当前日期，如：GetYYYYMMDDHHMMSS()；
* 11、排序方法：SortInterface(arr []interface{}, by func(p, q *interface{}) bool) 或 SortArrMap(maps []map[string]interface{}, by func(a, b map[string]interface{}) bool)
* 12、对象比较方法：CompareInterface(key string, a, b interface{}) (bool, interface{}, interface{})、CompareInter(a, b interface{})
* 13、增加通过reflect反射方式实现：赋值、clone：NewInterface(typ reflect.Type, data []byte) interface{} 、DeepClone(src interface{}) interface{}
* 14、增加SyncMap，使用泛型；
* 15、增加Array相关方法：IndexOf, Contains

## 说明
上面列举的这些共函数，都是我在项目中使用的。如果没有你需要的方法，可以提交pr；

## 联系
邮箱：xiaojun207@126.com