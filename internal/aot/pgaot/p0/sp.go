package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistInitBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v2 = l1
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistInitBuffer[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistInitBuffer[1]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if v20&int32(3) != 0 {
	} else {
	}
	v47 = F___memset(m, v20, int32(0), int32(_a_F_SpGistInitBuffer_0))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(_a_F_SpGistInitBuffer_1)
	v53 = int32(_a_F_SpGistInitBuffer_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v53)
	v59 = int32(_a_F_SpGistInitBuffer_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v59)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v59)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v63 = v20 + v62
	v64 = int32(_a_F_SpGistInitBuffer_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+6)) = uint16(v64)
	*(*uint16)(unsafe.Add(mBase, uint32(v63))) = uint16(v2)
	return
}
