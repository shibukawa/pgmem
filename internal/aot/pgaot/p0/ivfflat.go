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
func F_IvfflatCheckNorm(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_FunctionCall1Coll(m, l0, l1, l2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.F64_gt(base.F64_reinterpret_i64(v4), float64(0))
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
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
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
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
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
	v34 = F_DirectFunctionCall1Coll(m, v28, l1, base.I64_extend_i32_u(v29+v30*v22))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v63) < base.Ui32(v62) {
		goto L2
	} else {
		goto L22
	}
L10:
	;
	return
L11:
	;
	v36 = base.I32_wrap_i64(v34)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v37 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = int32(18)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v43 == v41 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v54 = int32(1)
	if v37&v54 != 0 {
		v62 = int32(base.Ui32(v37) >> (uint(v54) % 32))
		goto L9
	} else {
		goto L21
	}
L15:
	;
	v46 = v41
	goto L17
L16:
	;
	v46 = int32(2)
	goto L17
L17:
	;
	if base.Ui32((v43-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v53 = int32(6)
	goto L20
L19:
	;
	v53 = v46
	goto L20
L20:
	;
	v62 = v53
	goto L9
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v62 = int32(base.Ui32(v58) >> (uint(int32(2)) % 32))
	goto L9
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v65 <= v22 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v62 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	base.MemoryCopy(m, v67+v22*v63, v36, v62)
	goto L26
L25:
	;
	goto L26
L26:
	;
	F_MemoryContextReset(m, l3)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L10
	} else {
		goto L27
	}
L27:
	;
	v74 = v22 + int32(1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v74 < v75 {
		v22 = v74
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
	v95 = m.ExcPending
	if v95 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_IvfflatNormVectors_2), int32(337), int32(_a_F_IvfflatNormVectors_3))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
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
	v111 = m.ExcPending
	if v111 != 0 {
		goto L10
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_IvfflatNormVectors_2), int32(326), int32(_a_F_IvfflatNormVectors_4))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
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
