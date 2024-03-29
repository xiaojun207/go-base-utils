package compress

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"io"
	"log"
	"os"
)

func WriteToFile(fileName string, data []byte) error {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	writer := zlib.NewWriter(file)
	writer.Write(data)
	writer.Flush()
	return writer.Close()
}

func ReadFromFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	r, err := zlib.NewReader(file)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(r)
}

func ZlibCompressFile(src, dest string) error {
	r, err := os.Open(src)
	if err != nil {
		return err
	}
	w, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	_, err = bufio.NewWriter(zlib.NewWriter(w)).ReadFrom(r)
	return err
}

// 进行zlib压缩
func ZlibCompress(src []byte) []byte {
	var in bytes.Buffer
	w, _ := zlib.NewWriterLevel(&in, zlib.BestCompression)
	w.Write(src)
	w.Flush()
	w.Close()
	return in.Bytes()
}

// 进行zlib解压缩
func ZlibUnCompress(compressSrc []byte) ([]byte, error) {
	b := bytes.NewReader(compressSrc)

	r, err := zlib.NewReader(b)
	if err != nil {
		log.Println("ZlibUnCompress.err:", err)
		return nil, err
	}
	io.ReadAll(r)
	var out bytes.Buffer
	io.Copy(&out, r)
	return out.Bytes(), err
}
