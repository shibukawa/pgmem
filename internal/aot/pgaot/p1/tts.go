package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tts_buffer_heap_copy_heap_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v4 != 0 {
		v32 = v4
		v34 = F_heap_copytuple(m, v32)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int32(0)
		} else {
			return v34
		}
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v6&int32(4) != 0 {
			v32 = int32(0)
			v34 = F_heap_copytuple(m, v32)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				return v34
			}
		} else {
			v9 = int32(4442992)
			v10 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
			v14 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v14)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v14
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v21 = F_heap_form_tuple(m, v18, v19, v20)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v21
				v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
				v28 = v26 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v28)
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v10
				v32 = v21
				v34 = F_heap_copytuple(m, v32)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return v34
				}
			}
		}
	}
}
func F_tts_heap_is_current_xact_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if base.Ui32(v24) < base.Ui32(int32(3)) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	return int32(0)
L5:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	F_errmsg(m, int32(57712), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	F_errfinish(m, int32(471331), int32(391), int32(364663))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L9:
	;
	return v144
L10:
	;
	v144 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[69]))
	if v35 == v24 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v144 = int32(1)
	goto L9
L14:
	;
	goto L15
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[70]))
	if v39 <= int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v144 = v136
	goto L9
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	if v43 == int32(0) {
		v136 = int32(0)
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[71]))
	v107 = int32(0)
	v109 = v39 - int32(1)
	goto L39
L20:
	;
	v48 = v43
	goto L21
L21:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v53 == int32(4) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v136 = int32(0)
	goto L16
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	if v100 != 0 {
		v48 = v100
		goto L21
	} else {
		goto L38
	}
L24:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v56 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v59 = int32(1)
	if v24 == v56 {
		v136 = v59
		goto L16
	} else {
		goto L26
	}
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v63 = v61 - int32(1)
	if v63 < int32(0) {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v68 = int32(0)
	v70 = v63
	goto L28
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v76 = int32(2)
	v77 = base.I32_div_s(v70-v68, v76)
	v78 = v77 + v68
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v74+v78<<(uint(v76)%32))))
	if v82 == v24 {
		v136 = v59
		goto L16
	} else {
		goto L30
	}
L29:
	;
	goto L23
L30:
	;
	v86 = F_TransactionIdPrecedes(m, v82, v24)
	mBase = m.M
	if v86 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v87 = v78 + int32(1)
	goto L33
L32:
	;
	v87 = v68
	goto L33
L33:
	;
	if v86 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v90 = v70
	goto L36
L35:
	;
	v90 = v78 - int32(1)
	goto L36
L36:
	;
	if v87 <= v90 {
		v68 = v87
		v70 = v90
		goto L28
	} else {
		goto L37
	}
L37:
	;
	goto L29
L38:
	;
	goto L22
L39:
	;
	v114 = int32(2)
	v115 = base.I32_div_s(v109-v107, v114)
	v116 = v115 + v107
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v105+v116<<(uint(v114)%32))))
	v121 = base.B2i32(v120 == v24)
	if v120 == v24 {
		v136 = v121
		goto L16
	} else {
		goto L41
	}
L40:
	;
	v136 = v121
	goto L16
L41:
	;
	v124 = base.B2i32(base.Ui32(v120) < base.Ui32(v24))
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v125 = v116 + int32(1)
	goto L44
L43:
	;
	v125 = v107
	goto L44
L44:
	;
	if base.Ui32(v120) < base.Ui32(v24) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v128 = v109
	goto L47
L46:
	;
	v128 = v116 - int32(1)
	goto L47
L47:
	;
	if v125 <= v128 {
		v107 = v125
		v109 = v128
		goto L39
	} else {
		goto L48
	}
L48:
	;
	goto L40
}
func F_tts_virtual_copy_heap_tuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_heap_form_tuple(m, v2, v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
