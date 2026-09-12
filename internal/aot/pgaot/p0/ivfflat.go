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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[32]))
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
				v25 = *(*int32)(unsafe.Add(mBase, _consts[32]))
				v27 = base.I32_div_s(v25, int32(1024))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v27
				F_errmsg(m, int32(566446), v5)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_errfinish(m, int32(513653), int32(128), int32(421524))
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v5 = int32(0)
	v10 = int32(4549024)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = l3
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v5 < v14 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L11
	} else {
		goto L40
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L11
	} else {
		goto L37
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L11
	} else {
		goto L34
	}
L4:
	;
	v22 = v5
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v11
	return
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v26 <= v22 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v33 = F_DirectFunctionCall1Coll(m, v28, l1, v29+v30*v22)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if base.Ui32(v63) < base.Ui32(v62) {
		goto L2
	} else {
		goto L26
	}
L11:
	;
	return
L12:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v35 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v38 = int32(6)
	v40 = int32(18)
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v42 == v40 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v54 = int32(1)
	if v35&v54 != 0 {
		v62 = int32(base.Ui32(v35) >> (uint(v54) % 32))
		goto L10
	} else {
		goto L25
	}
L16:
	;
	v45 = v40
	goto L18
L17:
	;
	v45 = int32(2)
	goto L18
L18:
	;
	if v42&int32(254) == int32(2) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = v38
	goto L21
L20:
	;
	v50 = v45
	goto L21
L21:
	;
	if v42 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v53 = v38
	goto L24
L23:
	;
	v53 = v50
	goto L24
L24:
	;
	v62 = v53
	goto L10
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v62 = int32(base.Ui32(v58) >> (uint(int32(2)) % 32))
	goto L10
L26:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v65 <= v22 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v62 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	F_MemoryContextReset(m, l3)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L11
	} else {
		goto L32
	}
L29:
	;
	v70 = F__emscripten_memcpy_bulkmem(m, v67+v22*v63, v33, v62)
	mBase = m.M
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	v75 = v22 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v75 < v76 {
		v22 = v75
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L8
L34:
	;
	F_errmsg_internal(m, int32(471885), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(339984), int32(326), int32(115202))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	F_errmsg_internal(m, int32(471885), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(339984), int32(337), int32(114929))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errmsg_internal(m, int32(471885), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L11
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(339984), int32(326), int32(115202))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
