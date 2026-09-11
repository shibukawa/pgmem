package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpGistUpdateMetaPage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v4 != 0 {
		v6 = F_ReadBuffer(m, l0, int32(0))
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = F_ConditionalLockBuffer(m, v6)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				if v8 != 0 {
					if v6 < int32(0) {
						v13 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(v6^int32(-1))<<(uint(int32(2))%32))))
						v27 = v19
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, _consts[1]))
						v27 = v21 + v6<<(uint(int32(13))%32) + int32(-8192)
					}
					v28 = *(*int64)(unsafe.Add(mBase, uint32(v4)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+28)) = v28
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v4)+120))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+84)) = v30
					v32 = *(*int64)(unsafe.Add(mBase, uint32(v4)+112))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+76)) = v32
					v34 = *(*int64)(unsafe.Add(mBase, uint32(v4)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+68)) = v34
					v36 = *(*int64)(unsafe.Add(mBase, uint32(v4)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+60)) = v36
					v38 = *(*int64)(unsafe.Add(mBase, uint32(v4)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+52)) = v38
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v4)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+44)) = v40
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v4)+72))
					*(*int64)(unsafe.Add(mBase, uint32(v27)+36)) = v42
					v44 = int32(92)
					*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)) = uint16(v44)
					F_MarkBufferDirty(m, v6)
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						F_UnlockReleaseBuffer(m, v6)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_ReleaseBuffer(m, v6)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		return
	}
}
