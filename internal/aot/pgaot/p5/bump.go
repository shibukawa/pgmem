package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BumpContextCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = F_emscripten_builtin_malloc(m, int32(8192))
	mBase = m.M
	if v13 != 0 {
		v14 = int32(8192)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v14
		*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v13 - int32(-8192)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v13 + int32(88)
		v26 = v13 + int32(60)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+76)) = v26
		v29 = v13 + int32(72)
		*(*int32)(unsafe.Add(mBase, uint32(v29))) = v26
		*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v29
		*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v29
		v33 = int32(1073741823)
		if base.Ui32(v33) <= base.Ui32(l2) {
			v36 = v33
		} else {
			v36 = l2
		}
		v45 = v36
		for {
			if base.Ui32(int32(base.Ui32(l2-int32(16))>>(uint(int32(3))%32))) < base.Ui32(v45) {
				v45 = int32(base.Ui32(v45) >> (uint(int32(1)) % 32))
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v45
		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
		v55 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v13)+4)) = uint8(v55)
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(477)
		v58 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = v58
		*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
		*(*int64)(unsafe.Add(mBase, uint32(v13)+20)) = v58
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(1805788)
		if l0 != 0 {
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v70
			if v70 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v70)+24)) = v13
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v13
			v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)) = uint8(v74)
		} else {
			v76 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v76
			*(*uint8)(unsafe.Add(mBase, uint32(v13)+5)) = uint8(v76)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(8192)
		m.G0 = v10 + int32(16)
		return v13
	} else {
		v88 = *(*int32)(unsafe.Add(mBase, _consts[145]))
		F_MemoryContextStats(m, v88)
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return int32(0)
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(8389))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(14086), int32(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
						F_errdetail(m, int32(696887), v10)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(521207), int32(185), int32(374203))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
