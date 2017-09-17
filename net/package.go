package net

import (
	kutils "github.com/heyuanlong/gosky/utils"
)

func ParsePackage(buf []byte) (msgLen int,msgType int,pBuf []byte) {
	zlen := len(buf)
	if zlen < 6{
	return 0,0,nil
	}
	msgLen = int(kutils.BytesToUint16Big(buf[0:2])) + 2
	if zlen < msgLen {
	return 0,0,nil
	}
	msgType = kutils.BytesToIntLittle(buf[2:6])
	pBuf = buf[6:msgLen]
	return
}

func SetPackage(msgType int , buf []byte) []byte  {
	pk := make([]byte,0)
	pkSize :=  4 + len(buf)

	pk = append(pk, kutils.Uint16ToBytesBig(uint16(pkSize))... )
	pk = append(pk, kutils.IntToBytesLittle( msgType)... )
	pk = append(pk, buf...)

	return pk
}