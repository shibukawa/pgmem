package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_entryIsMoveRight(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1+v11)))
	if v13 == int32(-1) {
		v84 = int32(0)
		m.G0 = v9 + int32(16)
		return v84
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+12)))
		if base.Ui32(int32(25)) <= base.Ui32(v18) {
			v28 = int32(base.Ui32(v18+int32(_a_F_entryIsMoveRight_0))>>(uint(int32(2))%32)) & int32(_a_F_entryIsMoveRight_1)
		} else {
			v28 = int32(0)
		}
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l1+v28<<(uint(int32(2))%32))+20))
		v35 = l1 + v32&int32(_a_F_entryIsMoveRight_2)
		v36 = F_gintuple_get_attrnum(m, v17, v35)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v43 = F_gintuple_get_key(m, v40, v35, v9+int32(15))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+54)))
				if v45 != v36 {
					if base.Ui32(v45) < base.Ui32(v36) {
						v50 = int32(-1)
					} else {
						v50 = int32(1)
					}
					v77 = v50
					v84 = base.B2i32(int32(0) < v77)
					m.G0 = v9 + int32(16)
					return v84
				} else {
					v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+64)))
					v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v51 != v52 {
						if v51 < v52 {
							v57 = int32(-1)
						} else {
							v57 = int32(1)
						}
						v77 = v57
						v84 = base.B2i32(int32(0) < v77)
						m.G0 = v9 + int32(16)
						return v84
					} else {
						if v51 != 0 {
							v77 = int32(0)
							v84 = base.B2i32(int32(0) < v77)
							m.G0 = v9 + int32(16)
							return v84
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v36<<(uint(int32(2))%32)+v59)+uint32(_c_F_entryIsMoveRight[0])))
							v71 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
							v72 = F_FunctionCall2Coll(m, v59+v36*int32(28)+int32(112), v70, v71, v43)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v77 = base.I32_wrap_i64(v72)
								v84 = base.B2i32(int32(0) < v77)
								m.G0 = v9 + int32(16)
								return v84
							}
						}
					}
				}
			}
		}
	}
}
