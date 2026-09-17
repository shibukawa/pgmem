package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IvfflatCheckMemoryUsage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatCheckMemoryUsage[0]))
	if base.Ui32(v8) < base.Ui32(int32(base.Ui32(l0)>>(uint(int32(10))%32))) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(base.Ui32(l0)>>(uint(int32(20))%32)) + int32(1)
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatCheckMemoryUsage[0]))
				v27 = base.I32_div_s(v25, int32(1024))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v27
				F_errmsg(m, int32(_a_F_IvfflatCheckMemoryUsage_0), v5)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_IvfflatCheckMemoryUsage_1), int32(128), int32(_a_F_IvfflatCheckMemoryUsage_2))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_IvfflatCheckNorm(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	v4 = F_FunctionCall1Coll(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
		return base.F64_gt(v8, float64(0))
	}
}
func F_IvfflatNormVectors(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v5 = int32(0)
	v10 = int32(_a_F_IvfflatNormVectors_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_IvfflatNormVectors[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_IvfflatNormVectors[0])) = l3
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v5 < v14 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L10
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L10
	} else {
		goto L29
	}
L3:
	;
	v22 = v5
	goto L6
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IvfflatNormVectors[0])) = v11
	return
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v26 <= v22 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v33 = F_DirectFunctionCall1Coll(m, v28, l1, v29+v30*v22)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v61) < base.Ui32(v60) {
		goto L2
	} else {
		goto L22
	}
L10:
	;
	return
L11:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v35 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v39 = int32(18)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v41 == v39 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v52 = int32(1)
	if v35&v52 != 0 {
		v60 = int32(base.Ui32(v35) >> (uint(v52) % 32))
		goto L9
	} else {
		goto L21
	}
L15:
	;
	v44 = v39
	goto L17
L16:
	;
	v44 = int32(2)
	goto L17
L17:
	;
	if base.Ui32((v41-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v51 = int32(6)
	goto L20
L19:
	;
	v51 = v44
	goto L20
L20:
	;
	v60 = v51
	goto L9
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v60 = int32(base.Ui32(v56) >> (uint(int32(2)) % 32))
	goto L9
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v63 <= v22 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v60 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	base.MemoryCopy(m, v65+v22*v61, v33, v60)
	goto L26
L25:
	;
	goto L26
L26:
	;
	F_MemoryContextReset(m, l3)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v72 = v22 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v72 < v73 {
		v22 = v72
		goto L6
	} else {
		goto L28
	}
L28:
	;
	goto L7
L29:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatNormVectors_1), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_IvfflatNormVectors_2), int32(337), int32(_a_F_IvfflatNormVectors_3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_errmsg_internal(m, int32(_a_F_IvfflatNormVectors_1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_IvfflatNormVectors_2), int32(326), int32(_a_F_IvfflatNormVectors_4))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
