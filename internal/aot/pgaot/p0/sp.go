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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v2 = l1
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistInitBuffer[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_SpGistInitBuffer[1]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v21 = int32(_a_F_SpGistInitBuffer_0)
	v23 = int32(0)
	if v23|(v20&int32(3)|int32(1)) == v23 {
		v39 = v20 + v21
		v41 = v20 + int32(4)
		if base.Ui32(v41) < base.Ui32(v39) {
			v43 = v39
		} else {
			v43 = v41
		}
		v48 = (v20^int32(-1)+v43)&int32(-4) + int32(4)
		if v48 == int32(0) {
		} else {
			base.MemoryFill(m, v20, int32(0), v48)
		}
	} else {
		base.MemoryFill(m, v20, int32(0), v21)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(_a_F_SpGistInitBuffer_1)
	v62 = int32(_a_F_SpGistInitBuffer_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v62)
	v68 = int32(_a_F_SpGistInitBuffer_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v68)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v68)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v72 = v20 + v71
	v73 = int32(_a_F_SpGistInitBuffer_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)) = uint16(v73)
	*(*uint16)(unsafe.Add(mBase, uint32(v72))) = uint16(v2)
	return
}
