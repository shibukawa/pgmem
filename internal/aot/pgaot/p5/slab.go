package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlabContextCreate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v10 = m.G0
	v11 = int32(32)
	v12 = v10 - v11
	m.G0 = v12
	v15 = l2 - v11
	v16 = int32(4)
	if base.Ui32(l3) <= base.Ui32(v16) {
		v19 = v16
	} else {
		v19 = l3
	}
	v25 = (v19+int32(7))&int32(-8) + int32(8)
	v26 = base.I32_div_u_s(v15, v25)
	if base.Ui32(v25) <= base.Ui32(v15) {
		v29 = F_emscripten_builtin_malloc(m, int32(104))
		mBase = m.M
		if v29 == int32(0) {
			v122 = *(*int32)(unsafe.Add(mBase, _consts[146]))
			F_MemoryContextStats(m, v122)
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return int32(0)
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(12790), int32(0))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
							F_errdetail(m, int32(623334), v12+int32(16))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(476678), int32(372), int32(338330))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
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
		} else {
			v32 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v32
			*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v26
			*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v19
			v42 = v32
			for {
				v48 = int32(1)
				if v48 < int32(base.Ui32(v26)>>(uint(v42)%32)) {
					v42 = v42 + v48
					continue
				} else {
					break
				}
				break
			}
			v53 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v29)+76)) = v53
			*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = v42
			v57 = v29 + int32(96)
			*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v57
			v60 = v29 + int32(88)
			*(*int32)(unsafe.Add(mBase, uint32(v29)+92)) = v60
			v63 = v29 + int32(80)
			*(*int32)(unsafe.Add(mBase, uint32(v29)+84)) = v63
			*(*int32)(unsafe.Add(mBase, uint32(v29)+80)) = v63
			v67 = v29 + int32(68)
			*(*int32)(unsafe.Add(mBase, uint32(v29)+72)) = v67
			*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v67
			*(*int32)(unsafe.Add(mBase, uint32(v57))) = v57
			*(*int32)(unsafe.Add(mBase, uint32(v60))) = v60
			*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = l0
			v75 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v29)+4)) = uint8(v75)
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(476)
			v78 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v29)+36)) = v78
			*(*int32)(unsafe.Add(mBase, uint32(v29)+32)) = l1
			*(*int64)(unsafe.Add(mBase, uint32(v29)+20)) = v78
			*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v53
			*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(1727892)
			if l0 != 0 {
				v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v90
				if v90 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v90)+24)) = v29
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v29
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)) = uint8(v94)
			} else {
				v96 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v29)+28)) = v96
				*(*uint8)(unsafe.Add(mBase, uint32(v29)+5)) = uint8(v96)
			}
			m.G0 = v12 + int32(32)
			return v29
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v110 = m.ExcPending
		if v110 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v19
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l2
			F_errmsg_internal(m, int32(143896), v12)
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(476678), int32(360), int32(338330))
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
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
