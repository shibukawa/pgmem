package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogicalOutputWrite(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(int32(1073741820)) <= base.Ui32(v12) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(436314), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errfinish(m, int32(514329), int32(71), int32(364308))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
		v29 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)) = uint8(v29)
		*(*uint16)(unsafe.Add(mBase, uint32(v9))) = uint16(v29)
		v33 = F_Int64GetDatum(m, l1)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v33
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
			v40 = F_cstring_to_text_with_len(m, v38, v39)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v40
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
				F_tuplestore_putvalues(m, v43, v44, v9+int32(4), v9)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					v49 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v49 + int64(1)
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	}
}
func F_LogicalTapeClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v3 != 0 {
		F_pfree(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			return
		}
	}
}
func F_LogicalTapeRewindForRead(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v108 int64
	_ = v108
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v121 int64
	_ = v121
	var v133 int32
	_ = v133
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = int32(8192)
	if base.Ui32(l1) <= base.Ui32(v15) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v25 = int32(8192)
	goto L3
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v27 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v18 = v15
	goto L6
L5:
	;
	v18 = l1
	goto L6
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if base.Ui32(v18) < base.Ui32(v19) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v21 = v18
	goto L9
L8:
	;
	v21 = v19
	goto L9
L9:
	;
	v25 = v21 & int32(-8192)
	goto L3
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v30 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v48 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+uint32(_consts[1235]))) = base.I64_extend_i32_s(int32(0) - v37)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ltsWriteBlock(m, v41, v42, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v46)
	goto L12
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	F_pfree(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if int32(0) < v55 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	return
L25:
	;
	v60 = v55
	goto L28
L26:
	;
	v168 = v54
	goto L27
L27:
	;
	F_pfree(m, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L16
	} else {
		goto L49
	}
L28:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+40)))
	if v68 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v168 = v157
	goto L27
L30:
	;
	v153 = int32(1)
	if v153 < v60 {
		v60 = v60 - v153
		goto L28
	} else {
		goto L48
	}
L31:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v69+v60<<(uint(int32(3))%32)-int32(8))))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v26)+56))
	if v76 < base.I64_extend_i32_u(v77) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v26)+48)) = v95 + int64(1)
	if v95 == int64(0) {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	v94 = v80
	v95 = v76
	goto L32
L34:
	;
	goto L35
L35:
	;
	v82 = v77 << (uint(int32(4)) % 32)
	if base.Ui32(int32(1073741823)) < base.Ui32(v82) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v77 << (uint(int32(1)) % 32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v26)+44))
	v89 = F_repalloc(m, v88, v82)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v89
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v26)+48))
	v94 = v89
	v95 = v92
	goto L32
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v94+v133<<(uint(int32(3))%32)))) = v75
	goto L30
L39:
	;
	v133 = int32(0)
	goto L38
L40:
	;
	goto L41
L41:
	;
	v108 = v95
	goto L42
L42:
	;
	v112 = base.I32_wrap_i64(v108)
	v113 = int64(1)
	v114 = v108 - v113
	v116 = int64(base.Ui64(v114) >> (uint(v113) % 64))
	v117 = base.I32_wrap_i64(v116)
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v94+v117<<(uint(int32(3))%32))))
	if v121 < v75 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v133 = v117
	goto L38
L44:
	;
	v133 = v112
	goto L38
L45:
	;
	goto L46
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v94+v112<<(uint(int32(3))%32)))) = v121
	if base.Ui64(int64(1)) < base.Ui64(v114) {
		v108 = v116
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	goto L29
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+60)) = int64(0)
	goto L24
}
func F_ResetLogicalStreamingState(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _consts[115])) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, _consts[114])) = v2
	return
}
