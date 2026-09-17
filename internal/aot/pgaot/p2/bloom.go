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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
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
			v47 = int32(_a_F_BloomFillMetapage_0)
			v49 = int32(0)
			if v49|(l1&int32(3)|int32(1)) == v49 {
				v65 = l1 + v47
				v67 = l1 + int32(4)
				if base.Ui32(v67) < base.Ui32(v65) {
					v69 = v65
				} else {
					v69 = v67
				}
				v74 = (l1^int32(-1)+v69)&int32(-4) + int32(4)
				if v74 == int32(0) {
				} else {
					base.MemoryFill(m, l1, int32(0), v74)
				}
			} else {
				base.MemoryFill(m, l1, int32(0), v47)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(_a_F_BloomFillMetapage_1)
			v88 = int32(_a_F_BloomFillMetapage_2)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v88)
			v94 = int32(_a_F_BloomFillMetapage_3)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v94)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v94)
			v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
			v98 = l1 + v97
			v99 = int32(_a_F_BloomFillMetapage_4)
			*(*uint16)(unsafe.Add(mBase, uint32(v98)+6)) = uint16(v99)
			v101 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v98)+2)) = uint16(v101)
			base.MemoryFill(m, l1+int32(28), int32(0), int32(_a_F_BloomFillMetapage_5))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(-609481235)
			base.MemoryCopy(m, l1+int32(32), v46, int32(136))
			v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
			v116 = v114 + int32(_a_F_BloomFillMetapage_6)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v116)
			return
		}
	} else {
		v46 = v4
		v47 = int32(_a_F_BloomFillMetapage_0)
		v49 = int32(0)
		if v49|(l1&int32(3)|int32(1)) == v49 {
			v65 = l1 + v47
			v67 = l1 + int32(4)
			if base.Ui32(v67) < base.Ui32(v65) {
				v69 = v65
			} else {
				v69 = v67
			}
			v74 = (l1^int32(-1)+v69)&int32(-4) + int32(4)
			if v74 == int32(0) {
			} else {
				base.MemoryFill(m, l1, int32(0), v74)
			}
		} else {
			base.MemoryFill(m, l1, int32(0), v47)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(_a_F_BloomFillMetapage_1)
		v88 = int32(_a_F_BloomFillMetapage_2)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v88)
		v94 = int32(_a_F_BloomFillMetapage_3)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v94)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v94)
		v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
		v98 = l1 + v97
		v99 = int32(_a_F_BloomFillMetapage_4)
		*(*uint16)(unsafe.Add(mBase, uint32(v98)+6)) = uint16(v99)
		v101 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v98)+2)) = uint16(v101)
		base.MemoryFill(m, l1+int32(28), int32(0), int32(_a_F_BloomFillMetapage_5))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = int32(-609481235)
		base.MemoryCopy(m, l1+int32(32), v46, int32(136))
		v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		v116 = v114 + int32(_a_F_BloomFillMetapage_6)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)) = uint16(v116)
		return
	}
}
