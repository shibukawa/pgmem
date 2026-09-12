package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BloomFillMetapage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v46 int32
	_ = v46
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v4 == int32(0) {
		v8 = F_palloc0(m, int32(136))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+132)) = int32(2)
			v12 = int64(8589934594)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+124)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+116)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+108)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+100)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+92)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+84)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+76)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+68)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+60)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+52)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+44)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+28)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(8589934597)
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(544)
			v46 = v8
			if l1&int32(3) != 0 {
			} else {
			}
			v73 = F___memset(m, l1, int32(0), int32(_a_F_BloomFillMetapage_0))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(_a_F_BloomFillMetapage_1)
			v79 = int32(_a_F_BloomFillMetapage_2)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v79)
			v85 = int32(_a_F_BloomFillMetapage_3)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v85)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v85)
			v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
			v89 = l1 + v88
			v90 = int32(_a_F_BloomFillMetapage_4)
			*(*uint16)(unsafe.Add(mBase, uint32(v89)+6)) = uint16(v90)
			v92 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v89)+2)) = uint16(v92)
			v99 = F__emscripten_memset_bulkmem(m, l1+int32(28), base.I32_extend8_s(int32(0)), int32(_a_F_BloomFillMetapage_5))
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(-609481235)
			v105 = F__emscripten_memcpy_bulkmem(m, l1+int32(32), v46, int32(136))
			mBase = m.M
			v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
			v109 = v107 + int32(_a_F_BloomFillMetapage_6)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v109)
			return
		}
	} else {
		v46 = v4
		if l1&int32(3) != 0 {
		} else {
		}
		v73 = F___memset(m, l1, int32(0), int32(_a_F_BloomFillMetapage_0))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(_a_F_BloomFillMetapage_1)
		v79 = int32(_a_F_BloomFillMetapage_2)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v79)
		v85 = int32(_a_F_BloomFillMetapage_3)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v85)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v85)
		v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
		v89 = l1 + v88
		v90 = int32(_a_F_BloomFillMetapage_4)
		*(*uint16)(unsafe.Add(mBase, uint32(v89)+6)) = uint16(v90)
		v92 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v89)+2)) = uint16(v92)
		v99 = F__emscripten_memset_bulkmem(m, l1+int32(28), base.I32_extend8_s(int32(0)), int32(_a_F_BloomFillMetapage_5))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(-609481235)
		v105 = F__emscripten_memcpy_bulkmem(m, l1+int32(32), v46, int32(136))
		mBase = m.M
		v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		v109 = v107 + int32(_a_F_BloomFillMetapage_6)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v109)
		return
	}
}
