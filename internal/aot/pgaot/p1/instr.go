package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InstrAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	v3 = l2
	v4 = int32(0)
	v14 = F_palloc0(m, l0*int32(416))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if l1&int32(11) == int32(0) {
		} else {
			if l0 <= int32(0) {
			} else {
				v24 = int32(1)
				v25 = l1 & v24
				v26 = int32(3)
				v29 = int32(base.Ui32(l1)>>(uint(v26)%32)) & v24
				v33 = int32(base.Ui32(l1)>>(uint(v24)%32)) & v24
				v35 = l0 & v26
				v36 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(l0) {
					v42 = v36
					v50 = v4
					for {
						v52 = int32(416)
						v54 = v14 + v42*v52
						*(*uint8)(unsafe.Add(mBase, uint32(v54)+2)) = uint8(v29)
						*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)) = uint8(v33)
						*(*uint8)(unsafe.Add(mBase, uint32(v54)+3)) = uint8(v3)
						*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v25)
						v63 = v14 + (v42|int32(1))*v52
						*(*uint8)(unsafe.Add(mBase, uint32(v63)+2)) = uint8(v29)
						*(*uint8)(unsafe.Add(mBase, uint32(v63)+1)) = uint8(v33)
						*(*uint8)(unsafe.Add(mBase, uint32(v63)+3)) = uint8(v3)
						*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v25)
						v72 = v14 + (v42|int32(2))*v52
						*(*uint8)(unsafe.Add(mBase, uint32(v72)+2)) = uint8(v29)
						*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)) = uint8(v33)
						*(*uint8)(unsafe.Add(mBase, uint32(v72)+3)) = uint8(v3)
						*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v25)
						v81 = v14 + (v42|int32(3))*v52
						*(*uint8)(unsafe.Add(mBase, uint32(v81)+2)) = uint8(v29)
						*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)) = uint8(v33)
						*(*uint8)(unsafe.Add(mBase, uint32(v81)+3)) = uint8(v3)
						*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v25)
						v86 = int32(4)
						v87 = v42 + v86
						v89 = v50 + v86
						if v89 != l0&int32(2147483644) {
							v42 = v87
							v50 = v89
							continue
						} else {
							break
						}
						break
					}
					v92 = v87
				} else {
					v92 = v36
				}
				if v35 == int32(0) {
				} else {
					v105 = v92
					v112 = v4
					for {
						v117 = v14 + v105*int32(416)
						*(*uint8)(unsafe.Add(mBase, uint32(v117)+2)) = uint8(v29)
						*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)) = uint8(v33)
						*(*uint8)(unsafe.Add(mBase, uint32(v117)+3)) = uint8(v3)
						*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v25)
						v122 = int32(1)
						v125 = v112 + v122
						if v125 != v35 {
							v105 = v105 + v122
							v112 = v125
							continue
						} else {
							break
						}
						break
					}
				}
			}
		}
		return v14
	}
}
func F_InstrEndLoop(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v18 float64
	_ = v18
	var v22 int64
	_ = v22
	var v26 float64
	_ = v26
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v2 == int32(1) {
		v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v5 != int64(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(430509), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errfinish(m, int32(515548), int32(149), int32(246271))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v8 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v8)
			v10 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
			v11 = *(*float64)(unsafe.Add(mBase, uint32(l0)+200))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+200)) = base.F64_add(v10, v11)
			v14 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(l0)+216))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+216)) = base.F64_add(v14, v15)
			v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+232))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+232)) = base.F64_add(v18, float64(1))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
			v26 = *(*float64)(unsafe.Add(mBase, uint32(l0)+208))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+208)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v22), float64(1e+09)), v26)
			v30 = l0 + int32(8)
			v31 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v31
			*(*int64)(unsafe.Add(mBase, uint32(v30))) = v31
			return
		}
	} else {
		return
	}
}
