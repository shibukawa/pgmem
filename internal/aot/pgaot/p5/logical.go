package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogicalSlotAdvanceAndCheckSnapState(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v220 int64
	_ = v220
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(224)
	m.G0 = v13
	v19 = v3
	v20 = int32(-1)
	v22 = v3
	v23 = v3
	v24 = v3
	goto L1
L1:
	;
	if v20 != int32(1) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v44 = v19
	v45 = v22
	v46 = v23
	v47 = v24
	goto L5
L5:
	;
	goto L11
L6:
	;
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v28)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v30 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[0]))
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1]))
	v38 = v13 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v13 + int32(12)
	goto L9
L9:
	;
	v44 = v30
	v45 = v34
	v46 = v36
	v47 = base.B2i32(l1 == v30)
	goto L5
L10:
	;
	goto L2
L11:
	;
	if v44 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	v219 = int32(m.ExcTag)
	v220 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v219 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(415)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(416)
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1])) = v13 + int32(32)
	v61 = int32(0)
	v68 = F_CreateDecodingContext(m, int64(0), v61, int32(1), v13+int32(20), v61, v61, v61)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[0])) = v45
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1])) = v46
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L13
	} else {
		goto L59
	}
L17:
	;
	F_WaitForStandbyConfirmation(m, l0)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[2]))
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)+104))
	F_XLogBeginRead(m, v72, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+40))
	if base.Ui64(v81) < base.Ui64(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v87 = v80
	goto L24
L22:
	;
	goto L23
L23:
	;
	if v47 != 0 {
		goto L42
	} else {
		goto L43
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(0)
	v97 = F_XLogReadRecord(m, v87, v13+int32(16))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L13
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v99 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v97 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v104
	F_errmsg_internal(m, int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_0), v13)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_1), int32(2245), int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L10
L33:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	F_LogicalDecodingProcessRecord(m, v68, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[3]))
	if v118 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L35
L37:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L13
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v121)+40))
	if base.Ui64(v122) < base.Ui64(l0) {
		v87 = v121
		goto L24
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	goto L25
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v140)+40))
	if v141 != int64(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	if v135 != int32(2) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v138 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v138)
	goto L42
L45:
	;
	F_LogicalConfirmReceivedLocation(m, l0)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[2]))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v149)+120))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v68)+52))
	if v151 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+216)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+212)) = int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_3)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+208)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = int32(1058)
	v159 = int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_4)
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[0])) = v13 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+196)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v13)+204)) = v13 + int32(208)
	v169 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+164)) = uint8(v169)
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+147)) = uint8(v169)
	m.T0[v151].(func(*base.Module, int32))(m, v68)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	F_ReorderBufferFree(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L13
	} else {
		goto L54
	}
L53:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[0])) = v176
	goto L52
L54:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
	F_FreeSnapshotBuilder(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	F_XLogReaderFree(m, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	F_MemoryContextDelete(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[0])) = v45
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1])) = v46
	m.G0 = v13 + int32(224)
	return v150
L59:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	goto L12
L61:
	;
	v224 = int32(v220)
	m.G0 = v13
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v227)))
	if v13+int32(12) == v230 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	m.ExcPending = 1
	goto L70
L63:
	;
	if v234 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v234 = v232
	goto L66
L65:
	;
	v234 = int32(0)
	goto L66
L66:
	;
	goto L63
L67:
	;
	F___wasm_longjmp(m, v227, v226)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v19 = v226
	v20 = v234
	v22 = v45
	v23 = v46
	v24 = v47
	goto L1
L70:
	;
	return int64(0)
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LogicalTapeCreate(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int64
	_ = v38
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 == int32(0) {
		v25 = F_palloc(m, int32(72))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			v27 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v27
			v29 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)) = uint8(v29)
			v31 = int32(1)
			*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v31)
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
			*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v27
			*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v27
			v38 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v38
			*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v38
			*(*int64)(unsafe.Add(mBase, uint32(v25)+52)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = int32(1073741823)
			*(*int64)(unsafe.Add(mBase, uint32(v25)+60)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = v29
			return v25
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v6 != int32(-1) {
			v25 = F_palloc(m, int32(72))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = int64(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v27
				v29 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)) = uint8(v29)
				v31 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+4)) = uint16(v31)
				*(*int32)(unsafe.Add(mBase, uint32(v25))) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v27
				*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v27
				v38 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v38
				*(*int64)(unsafe.Add(mBase, uint32(v25)+52)) = v38
				*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = int32(1073741823)
				*(*int64)(unsafe.Add(mBase, uint32(v25)+60)) = v38
				*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = v29
				return v25
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_LogicalTapeCreate_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_LogicalTapeCreate_1), int32(690), int32(_a_F_LogicalTapeCreate_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
func F_LogicalTapeRead(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	v4 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v7 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = F_palloc(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l2 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v11
	*(*int64)(unsafe.Add(mBase, uint32(l0)+52)) = int64(0)
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v18
	v20 = F_ltsReadFillBuffer(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	goto L3
L7:
	;
	return int32(0)
L8:
	;
	goto L9
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v28 = l1
	v29 = l2
	v31 = v26
	v32 = v4
	goto L10
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v33 <= v31 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return v59
L12:
	;
	goto L11
L13:
	;
	v35 = F_ltsReadFillBuffer(m, l0)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	v41 = v31
	v42 = v33
	goto L15
L15:
	;
	v43 = v42 - v41
	if base.Ui32(v43) < base.Ui32(v29) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v35 == int32(0) {
		v59 = v32
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v41 = v39
	v42 = v40
	goto L15
L18:
	;
	v45 = v43
	goto L20
L19:
	;
	v45 = v29
	goto L20
L20:
	;
	if v45 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	base.MemoryCopy(m, v28, v46+v41, v45)
	goto L23
L22:
	;
	goto L23
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v50 = v49 + v45
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v50
	v52 = v45 + v32
	v54 = v29 - v45
	if v54 != 0 {
		v28 = v28 + v45
		v29 = v54
		v31 = v50
		v32 = v52
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v59 = v52
	goto L12
}
func F_LogicalTapeSetClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_BufFileClose(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		F_pfree(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_LogicalTapeSetCreate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	v1 = l0
	v7 = m.G0
	v9 = v7 - int32(1024)
	m.G0 = v9
	v12 = F_palloc(m, int32(64))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)) = uint8(v16)
		v18 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(32)
		v27 = F_palloc(m, int32(256))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)) = uint8(v1)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v27
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l1
			v35 = int32(0)
			if base.B2i32(l1 == v35)|base.B2i32(l2 != int32(-1)) == v35 {
				*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
				m.G0 = v9 + int32(1024)
				return v12
			} else {
				if l1 != 0 {
					v44 = base.I32_extend16_s(l2)
					if int32(0) <= v44 {
						v54 = v44
						v55 = int32(0)
					} else {
						v49 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v49)
						v54 = int32(0) - v44
						v55 = int32(1)
					}
					v57 = F_pg_ultoa_n(m, v54, v9+v55)
					mBase = m.M
					v60 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v9+(v57+v55)))) = uint8(v60)
					v62 = F_BufFileCreateFileSet(m, l1, v9)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v62
						m.G0 = v9 + int32(1024)
						return v12
					}
				} else {
					v66 = F_BufFileCreateTemp(m, int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v66
						m.G0 = v9 + int32(1024)
						return v12
					}
				}
			}
		}
	}
}
func F_logical_read_xlog_page(m *base.Module, l0 int32, l1 int64, l2 int32, l3 int64, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int64
	_ = v147
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int64
	_ = v184
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int64
	_ = v249
	var v253 int32
	_ = v253
	var v257 int64
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int64
	_ = v286
	var v290 int32
	_ = v290
	var v292 int64
	_ = v292
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v327 int64
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int64
	_ = v334
	var v338 int32
	_ = v338
	var v347 int64
	_ = v347
	var v353 int64
	_ = v353
	var v362 int64
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int64
	_ = v369
	var v373 int32
	_ = v373
	var v379 int64
	_ = v379
	var v385 int64
	_ = v385
	var v394 int64
	_ = v394
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int64
	_ = v423
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v497 int64
	_ = v497
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v518 int64
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v537 int64
	_ = v537
	var v540 int32
	_ = v540
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int64
	_ = v557
	var v558 int64
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v572 int32
	_ = v572
	v10 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v18 = l1 + base.I64_extend_i32_s(l2)
	v21 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	if v21 == v10 {
		v46 = int32(0)
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_WalSndShutdown(m)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L11
	} else {
		goto L174
	}
L2:
	;
	v497 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	if base.Ui64(v18) <= base.Ui64(v497) {
		goto L148
	} else {
		goto L149
	}
L3:
	;
	v54 = v46
	v58 = v10
	goto L14
L4:
	;
	if base.Ui64(v21) < base.Ui64(v18) {
		v46 = int32(100663303)
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[2]))
	if v29 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[3]))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+202)))
	if v34 != int32(1) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v39 = int32(21)
	goto L10
L9:
	;
	v39 = int32(19)
	goto L10
L10:
	;
	v40 = F_StandbySlotsHaveCaughtup(m, v21, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	if v40 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v46 = int32(100663302)
	goto L3
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[4]))
	v61 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = v61
	v66 = base.AtomicRmwOr32(m, v61, int32(_a_F_logical_read_xlog_page_0), v61)
	goto L16
L15:
	;
	v431 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[4]))
	v432 = int32(0)
	v435 = base.AtomicRmwOr32(m, v432, int32(_a_F_logical_read_xlog_page_0), v432)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	if v436 != 0 {
		goto L135
	} else {
		goto L136
	}
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[5]))
	if v68 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[6]))
	if v72 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_ProcessRepliesIfAny(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L11
	} else {
		goto L27
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[6])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	F_SyncRepInitConfig(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])))
	if v84 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	F_SyncRepReleaseWaiters(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	goto L21
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	if v90 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v54 != int32(100663302) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8])))
	if v95 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	if v105 != 0 {
		goto L28
	} else {
		goto L34
	}
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+308))
	v103 = base.B2i32(v101 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8])) = uint8(v103)
	v105 = v103
	goto L33
L32:
	;
	v105 = int32(0)
	goto L33
L33:
	;
	goto L30
L34:
	;
	v106 = F_GetXLogInsertEndRecPtr(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	F_XLogFlush(m, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	goto L28
L37:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8])))
	if v115 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	goto L39
L39:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	if v161 != 0 {
		goto L54
	} else {
		goto L55
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0])) = v158
	goto L39
L41:
	;
	if v125 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+308))
	v123 = base.B2i32(v121 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8])) = uint8(v123)
	v125 = v123
	goto L44
L43:
	;
	v125 = int32(0)
	goto L44
L44:
	;
	goto L41
L45:
	;
	v128 = int32(0)
	v130 = int32(_a_F_logical_read_xlog_page_1)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9]))
	v132 = int64(0)
	v135 = base.AtomicRmwCmpxchg64(m, v131, int32(272), v132, v132)
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[10])) = v135
	v140 = base.AtomicRmwOr32(m, v128, int32(_a_F_logical_read_xlog_page_2), v128)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9]))
	v147 = base.AtomicRmwCmpxchg64(m, v143, int32(264), v132, v132)
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[11])) = v147
	goto L50
L46:
	;
	goto L47
L47:
	;
	v156 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L11
	} else {
		goto L52
	}
L48:
	;
	v158 = v154
	goto L40
L50:
	;
	goto L51
L51:
	;
	v154 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[10]))
	goto L48
L52:
	;
	v158 = v156
	goto L40
L53:
	;
	goto L15
L54:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	v165 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[2]))
	if v167 == int32(0) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	v182 = v54
	goto L56
L56:
	;
	v184 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[12]))
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[13]))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v186)+32))
	if base.Ui64(v184) <= base.Ui64(v187) {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[3]))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+202)))
	if v172 != int32(1) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	if v163 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v178 = int32(21)
	goto L61
L60:
	;
	v178 = int32(19)
	goto L61
L61:
	;
	v179 = F_StandbySlotsHaveCaughtup(m, v165, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	if v179 != 0 {
		goto L53
	} else {
		goto L63
	}
L63:
	;
	v182 = int32(100663302)
	goto L56
L64:
	;
	if v161 != 0 {
		v222 = v182
		goto L69
	} else {
		goto L70
	}
L65:
	;
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v186)+24))
	if base.Ui64(v184) <= base.Ui64(v189) {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[14])))
	if v192&int32(1) != 0 {
		goto L64
	} else {
		goto L67
	}
L67:
	;
	F_WalSndKeepalive(m, int32(0), int64(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	goto L64
L69:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15])) = uint8(v224)
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[16]))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+8))
	v229 = m.T0[v228].(func(*base.Module) int32)(m)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L11
	} else {
		goto L81
	}
L70:
	;
	v200 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	if base.Ui64(v200) < base.Ui64(v18) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v222 = int32(100663303)
	goto L69
L72:
	;
	goto L73
L73:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[2]))
	if v206 == int32(0) {
		goto L53
	} else {
		goto L74
	}
L74:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[3]))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+202)))
	if v211 != int32(1) {
		goto L53
	} else {
		goto L75
	}
L75:
	;
	if v204 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v216 = int32(21)
	goto L78
L77:
	;
	v216 = int32(19)
	goto L78
L78:
	;
	v217 = F_StandbySlotsHaveCaughtup(m, v200, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L11
	} else {
		goto L79
	}
L79:
	;
	if v217 != 0 {
		goto L53
	} else {
		goto L80
	}
L80:
	;
	v222 = int32(100663302)
	goto L69
L81:
	;
	if v229 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[17])))
	if v232 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v249 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[18]))
	if v249 <= int64(0) {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[19])))
	if v236&int32(1) == int32(0) {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[16]))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v244 = m.T0[v243].(func(*base.Module) int32)(m)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L11
	} else {
		goto L86
	}
L86:
	;
	if v244 == int32(0) {
		goto L53
	} else {
		goto L87
	}
L87:
	;
	goto L83
L88:
	;
	F_WalSndCheckShutdownTimeout(m)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L11
	} else {
		goto L96
	}
L89:
	;
	v253 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[20]))
	if v253 <= int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v257 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[21]))
	if v257 < base.I64_extend_i32_u(v253)*int64(1000)+v249 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v265 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	if v265 == int32(0) {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(_a_F_logical_read_xlog_page_3), int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(_a_F_logical_read_xlog_page_4), int32(3008), int32(_a_F_logical_read_xlog_page_5))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	goto L1
L96:
	;
	v282 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[20]))
	if v282 <= int32(0) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v313 = m.G0
	v314 = int32(16)
	v315 = v313 - v314
	m.G0 = v315
	F_gettimeofday(m, v315)
	mBase = m.M
	v318 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
	v319 = int64(*(*int32)(unsafe.Add(mBase, uint32(v315)+8)))
	m.G0 = v315 + v314
	v327 = v319 + v318*int64(1000000) - int64(946684800000000)
	goto L105
L98:
	;
	v286 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[18]))
	if v286 <= int64(0) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[14])))
	if v290 != 0 {
		goto L97
	} else {
		goto L100
	}
L100:
	;
	v292 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[21]))
	if v292 < base.I64_extend_i32_u(int32(base.Ui32(v282)>>(uint(int32(1))%32)))*int64(1000)+v286 {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[16]))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+8))
	v307 = m.T0[v306].(func(*base.Module) int32)(m)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	if v307 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L97
L105:
	;
	v328 = int32(_a_F_logical_read_xlog_page_6)
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[20]))
	if v330 <= int32(0) {
		v366 = v328
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v369 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[22]))
	if v369 == int64(0) {
		v400 = v366
		goto L113
	} else {
		goto L114
	}
L107:
	;
	v334 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[18]))
	if v334 <= int64(0) {
		v366 = v328
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[14])))
	v347 = base.I64_extend_i32_u(int32(base.Ui32(v330)>>(uint((v338^int32(-1))&int32(1))%32)))*int64(1000) + v334
	if v347 <= v327 {
		v365 = int32(0)
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v366 = v365
	goto L106
L110:
	;
	goto L109
L111:
	;
	v353 = v347 - v327
	if base.B2i32(int64(0) < v327)^base.B2i32(v353 < v347)|base.B2i32(int64(2147483646000) < v353) != 0 {
		v365 = int32(2147483647)
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v362 = base.I64_div_s(v353+int64(999), int64(1000))
	v365 = base.I32_wrap_i64(v362)
	goto L110
L113:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[16]))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v407 = m.T0[v406].(func(*base.Module) int32)(m)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L11
	} else {
		goto L123
	}
L114:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[23]))
	if v373 <= int32(0) {
		v400 = v366
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v379 = base.I64_extend_i32_u(v373)*int64(1000) + v369
	if v379 <= v327 {
		v397 = int32(0)
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v397 < v366 {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	goto L116
L118:
	;
	v385 = v379 - v327
	if base.B2i32(int64(0) < v327)^base.B2i32(v385 < v379)|base.B2i32(int64(2147483646000) < v385) != 0 {
		v397 = int32(2147483647)
		goto L117
	} else {
		goto L119
	}
L119:
	;
	v394 = base.I64_div_s(v385+int64(999), int64(1000))
	v397 = base.I32_wrap_i64(v394)
	goto L117
L120:
	;
	v399 = v397
	goto L122
L121:
	;
	v399 = v366
	goto L122
L122:
	;
	v400 = v399
	goto L113
L123:
	;
	if v407 != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v409 = int32(6)
	goto L126
L125:
	;
	v409 = int32(2)
	goto L126
L126:
	;
	goto L127
L127:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v327-v58 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L11
	} else {
		goto L131
	}
L129:
	;
	v423 = v58
	goto L130
L130:
	;
	F_WalSndWait(m, v409, v400, v222)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L11
	} else {
		goto L133
	}
L131:
	;
	v421 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	v423 = v327
	goto L130
L133:
	;
	v54 = v222
	v58 = v423
	goto L14
L134:
	;
	goto L2
L135:
	;
	goto L134
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431))) = int32(1)
	v439 = int32(0)
	v442 = base.AtomicRmwOr32(m, v439, int32(_a_F_logical_read_xlog_page_0), v439)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v431)+4))
	if v443 == v439 {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	if v446 == int32(0) {
		goto L135
	} else {
		goto L138
	}
L138:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[24]))
	if v450 == v446 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v452 = m.G0
	v454 = v452 - int32(16)
	m.G0 = v454
	v457 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[25]))
	if v457 == int32(0) {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	goto L141
L141:
	;
	v480 = F_pgmem_kill(m, v446, int32(23))
	mBase = m.M
	goto L135
L142:
	;
	m.G0 = v454 + int32(16)
	goto L134
L143:
	;
	v460 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v454)+15)) = uint8(v460)
	goto L144
L144:
	;
	v464 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[26]))
	v468 = F_write(m, v464, v454+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v468 {
		goto L142
	} else {
		goto L146
	}
L145:
	;
	goto L142
L146:
	;
	v472 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[27]))
	if v472 == int32(27) {
		goto L144
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8])))
	if v502 == int32(1) {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v563 = int32(-1)
	goto L150
L150:
	;
	m.G0 = v15 + int32(48)
	return v563
L151:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])) = uint8(v512)
	if v512 != 0 {
		goto L157
	} else {
		goto L158
	}
L152:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9]))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v507)+308))
	v510 = base.B2i32(v508 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8])) = uint8(v510)
	v512 = v510
	goto L154
L153:
	;
	v512 = int32(0)
	goto L154
L154:
	;
	goto L151
L155:
	;
	F_XLogReadDetermineTimeline(m, l0, l1, l2, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L11
	} else {
		goto L164
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v524
	v526 = v524
	goto L155
L157:
	;
	v514 = F_GetWALInsertionTimeLineIfSet(m)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L11
	} else {
		goto L160
	}
L158:
	;
	goto L159
L159:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9]))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+300))
	goto L163
L160:
	;
	if v514 != 0 {
		v524 = v514
		goto L156
	} else {
		goto L161
	}
L161:
	;
	v518 = F_GetXLogReplayRecPtr(m, v15+int32(4))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L11
	} else {
		goto L162
	}
L162:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v526 = v520
	goto L155
L163:
	;
	v524 = v523
	goto L156
L164:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
	*(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[28])) = v530
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[29])) = uint8(base.B2i32(v530 != v533))
	v537 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[30])) = v537
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1240))
	*(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[31])) = v540
	if base.Ui64(l1-int64(-8192)) <= base.Ui64(v497) {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v548 = int32(_a_F_logical_read_xlog_page_7)
	goto L167
L166:
	;
	v548 = base.I32_wrap_i64(v497 - l1)
	goto L167
L167:
	;
	v550 = v15 + int32(8)
	v551 = F_WALRead(m, l0, l4, l1, v548, v533, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L11
	} else {
		goto L168
	}
L168:
	;
	if v551 == int32(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	F_WALReadRaiseError(m, v550)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L11
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v557 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	v558 = base.I64_div_u_s(l1, v557)
	v559 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
	F_CheckXLogRemoved(m, v558, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L11
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v563 = v548
	goto L150
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
