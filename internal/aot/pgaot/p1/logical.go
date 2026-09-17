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
			F_errmsg_internal(m, int32(_a_F_LogicalOutputWrite_0), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_LogicalOutputWrite_1), int32(71), int32(_a_F_LogicalOutputWrite_2))
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v114 int64
	_ = v114
	var v121 int64
	_ = v121
	var v125 int64
	_ = v125
	var v136 int32
	_ = v136
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v11 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(_a_F_LogicalTapeRewindForRead_0)
	if base.Ui32(l1) <= base.Ui32(v14) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v24 = int32(_a_F_LogicalTapeRewindForRead_0)
	goto L3
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v26 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v17 = v14
	goto L6
L5:
	;
	v17 = l1
	goto L6
L6:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if base.Ui32(v17) < base.Ui32(v18) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v20 = v17
	goto L9
L8:
	;
	v20 = v18
	goto L9
L9:
	;
	v24 = v20 & int32(-8192)
	goto L3
L10:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	if v29 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v32)+uint32(_c_F_LogicalTapeRewindForRead[0]))) = base.I64_extend_i32_s(int32(0) - v34)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_ltsWriteBlock(m, v38, v39, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v43)
	goto L12
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	F_pfree(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v51 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if int32(0) < v52 {
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
	v57 = v52
	goto L28
L26:
	;
	v164 = v51
	goto L27
L27:
	;
	F_pfree(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L16
	} else {
		goto L49
	}
L28:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+40)))
	if v64 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v164 = v154
	goto L27
L30:
	;
	v150 = int32(1)
	if v150 < v57 {
		v57 = v57 - v150
		goto L28
	} else {
		goto L48
	}
L31:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v65+v57<<(uint(int32(3))%32)-int32(8))))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v25)+48))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v25)+56))
	if v72 < base.I64_extend_i32_u(v73) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+48)) = v90 + int64(1)
	if v90 == int64(0) {
		goto L38
	} else {
		goto L39
	}
L33:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v89 = v76
	v90 = v72
	goto L32
L34:
	;
	goto L35
L35:
	;
	v78 = v73 << (uint(int32(4)) % 32)
	if base.Ui32(int32(1073741823)) < base.Ui32(v78) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = v73 << (uint(int32(1)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v25)+44))
	v85 = F_repalloc(m, v84, v78)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+44)) = v85
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v25)+48))
	v89 = v85
	v90 = v88
	goto L32
L38:
	;
	v136 = int32(0)
	goto L40
L39:
	;
	v102 = v90
	goto L41
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89+v136<<(uint(int32(3))%32)))) = v71
	goto L30
L41:
	;
	v106 = int64(1)
	v109 = int64(base.Ui64(v102-v106) >> (uint(v106) % 64))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v89+base.I32_wrap_i64(v109)<<(uint(int32(3))%32))))
	if v114 < v71 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v136 = base.I32_wrap_i64(v125)
	goto L40
L43:
	;
	goto L42
L44:
	;
	v125 = v102
	goto L43
L45:
	;
	goto L46
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v89+base.I32_wrap_i64(v102)<<(uint(int32(3))%32)))) = v114
	v121 = int64(0)
	if v109 != v121 {
		v102 = v109
		goto L41
	} else {
		goto L47
	}
L47:
	;
	v125 = v121
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
	*(*uint8)(unsafe.Add(mBase, _c_F_ResetLogicalStreamingState[0])) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, _c_F_ResetLogicalStreamingState[1])) = v2
	return
}
