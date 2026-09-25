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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v230 int64
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(224)
	m.G0 = v14
	v20 = v3
	v21 = int32(-1)
	v23 = v3
	v24 = v3
	v25 = v3
	v26 = v3
	goto L1
L1:
	;
	if v21 != int32(1) {
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
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[0]))
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v48 = v20
	v49 = v23
	v50 = v24
	v51 = v25
	v52 = v26
	goto L5
L5:
	;
	goto L11
L6:
	;
	v32 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v32)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v34 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1]))
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[2]))
	v42 = v14 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v14 + int32(12)
	goto L9
L9:
	;
	v48 = v34
	v49 = v38
	v50 = v40
	v51 = v31
	v52 = base.B2i32(l1 == v34)
	goto L5
L10:
	;
	goto L2
L11:
	;
	if v48 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	v229 = int32(m.ExcTag)
	v230 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v229 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(396)
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[2])) = v14 + int32(32)
	v66 = int32(0)
	v73 = F_CreateDecodingContext(m, int64(0), v66, int32(1), v14+int32(20), v66, v66, v66)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1])) = v49
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[2])) = v50
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L13
	} else {
		goto L59
	}
L17:
	;
	F_WaitForStandbyConfirmation(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[3]))
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v79)+104))
	F_XLogBeginRead(m, v77, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
	if base.Ui64(v86) < base.Ui64(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v92 = v85
	goto L24
L22:
	;
	goto L23
L23:
	;
	if v52 != 0 {
		goto L42
	} else {
		goto L43
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	v103 = F_XLogReadRecord(m, v92, v14+int32(16))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v105 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v103 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v110
	F_errmsg_internal(m, int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_0), v14)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_1), int32(2138), int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L10
L33:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	F_LogicalDecodingProcessRecord(m, v73, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[4]))
	if v124 != 0 {
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
	v126 = m.ExcPending
	if v126 != 0 {
		goto L13
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v127)+40))
	if base.Ui64(v128) < base.Ui64(l0) {
		v92 = v127
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
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[0])) = v51
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v149)+40))
	if v150 != int64(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v142 != int32(2) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v145 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v145)
	goto L42
L45:
	;
	F_LogicalConfirmReceivedLocation(m, l0)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[3]))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v158)+120))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v73)+52))
	if v160 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+216)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_3)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v14)+200)) = int32(993)
	v168 = int32(_a_F_LogicalSlotAdvanceAndCheckSnapState_4)
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1])) = v14 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+196)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v14)+204)) = v14 + int32(208)
	v178 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+164)) = uint8(v178)
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+147)) = uint8(v178)
	m.T0[v160].(func(*base.Module, int32))(m, v73)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	F_ReorderBufferFree(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L54
	}
L53:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v14)+196))
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1])) = v185
	goto L52
L54:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
	F_FreeSnapshotBuilder(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	F_XLogReaderFree(m, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	F_MemoryContextDelete(m, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[1])) = v49
	*(*int32)(unsafe.Add(mBase, _c_F_LogicalSlotAdvanceAndCheckSnapState[2])) = v50
	m.G0 = v14 + int32(224)
	return v159
L59:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	goto L12
L61:
	;
	v234 = int32(v230)
	m.G0 = v14
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v237)))
	if v14+int32(12) == v240 {
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
	if v244 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	v244 = v242
	goto L66
L65:
	;
	v244 = int32(0)
	goto L66
L66:
	;
	goto L63
L67:
	;
	F___wasm_longjmp(m, v237, v236)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v20 = v236
	v21 = v244
	v23 = v49
	v24 = v50
	v25 = v51
	v26 = v52
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
	var v52 int32
	_ = v52
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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int64
	_ = v124
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v166 int64
	_ = v166
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v229 int32
	_ = v229
	var v233 int64
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int64
	_ = v283
	var v284 int64
	_ = v284
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int64
	_ = v299
	var v303 int32
	_ = v303
	var v312 int64
	_ = v312
	var v318 int64
	_ = v318
	var v327 int64
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int64
	_ = v354
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v427 int64
	_ = v427
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int64
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int64
	_ = v464
	var v467 int32
	_ = v467
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int64
	_ = v484
	var v485 int64
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
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
	v500 = m.ExcPending
	if v500 != 0 {
		goto L11
	} else {
		goto L152
	}
L2:
	;
	v427 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	if base.Ui64(v18) <= base.Ui64(v427) {
		goto L129
	} else {
		goto L130
	}
L3:
	;
	v52 = v46
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
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[4]))
	v362 = int32(0)
	v365 = base.AtomicRmwOr32(m, v362, int32(_a_F_logical_read_xlog_page_0), v362)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	if v366 != 0 {
		goto L116
	} else {
		goto L117
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
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[6])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_ProcessRepliesIfAny(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L26
	}
L24:
	;
	F_SyncRepInitConfig(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v85 = F_XLogBackgroundFlush(m)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L11
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v52 != int32(100663302) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])))
	if v92 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	if v138 != 0 {
		goto L48
	} else {
		goto L49
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0])) = v135
	goto L33
L35:
	;
	if v102 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+316))
	v100 = base.B2i32(v98 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])) = uint8(v100)
	v102 = v100
	goto L38
L37:
	;
	v102 = int32(0)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v105 = int32(0)
	v107 = int32(_a_F_logical_read_xlog_page_1)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v109 = int64(0)
	v112 = base.AtomicRmwCmpxchg64(m, v108, int32(280), v109, v109)
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9])) = v112
	v117 = base.AtomicRmwOr32(m, v105, int32(_a_F_logical_read_xlog_page_2), v105)
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v124 = base.AtomicRmwCmpxchg64(m, v120, int32(272), v109, v109)
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[10])) = v124
	goto L44
L40:
	;
	goto L41
L41:
	;
	v133 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L11
	} else {
		goto L46
	}
L42:
	;
	v135 = v131
	goto L34
L44:
	;
	goto L45
L45:
	;
	v131 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9]))
	goto L42
L46:
	;
	v135 = v133
	goto L34
L47:
	;
	goto L15
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	v142 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[2]))
	if v144 == int32(0) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v158 = v52
	goto L50
L50:
	;
	v161 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[11]))
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[12]))
	v164 = *(*int64)(unsafe.Add(mBase, uint32(v163)+32))
	if base.Ui64(v161) <= base.Ui64(v164) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[3]))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+202)))
	if v149 != int32(1) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	if v140 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v155 = int32(21)
	goto L55
L54:
	;
	v155 = int32(19)
	goto L55
L55:
	;
	v156 = F_StandbySlotsHaveCaughtup(m, v142, v155)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	if v156 != 0 {
		goto L47
	} else {
		goto L57
	}
L57:
	;
	v158 = int32(100663302)
	goto L50
L58:
	;
	if v138 != 0 {
		v198 = v158
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v163)+24))
	if base.Ui64(v161) <= base.Ui64(v166) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[13])))
	if v169&int32(1) != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	F_WalSndKeepalive(m, int32(0), int64(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[14])) = uint8(v200)
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	v205 = m.T0[v204].(func(*base.Module) int32)(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L75
	}
L64:
	;
	v177 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	if base.Ui64(v177) < base.Ui64(v18) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v198 = int32(100663303)
	goto L63
L66:
	;
	goto L67
L67:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[2]))
	if v183 == int32(0) {
		goto L47
	} else {
		goto L68
	}
L68:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[3]))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+202)))
	if v188 != int32(1) {
		goto L47
	} else {
		goto L69
	}
L69:
	;
	if v181 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v193 = int32(21)
	goto L72
L71:
	;
	v193 = int32(19)
	goto L72
L72:
	;
	v194 = F_StandbySlotsHaveCaughtup(m, v177, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	if v194 != 0 {
		goto L47
	} else {
		goto L74
	}
L74:
	;
	v198 = int32(100663302)
	goto L63
L75:
	;
	if v205 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[16])))
	if v208 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v225 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[17]))
	if v225 <= int64(0) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[18])))
	if v212&int32(1) == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v220 = m.T0[v219].(func(*base.Module) int32)(m)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	if v220 == int32(0) {
		goto L47
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	v278 = m.G0
	v279 = int32(16)
	v280 = v278 - v279
	m.G0 = v280
	F_gettimeofday(m, v280)
	mBase = m.M
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v280)))
	v284 = int64(*(*int32)(unsafe.Add(mBase, uint32(v280)+8)))
	m.G0 = v280 + v279
	v292 = v284 + v283*int64(1000000) - int64(946684800000000)
	goto L96
L83:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[19]))
	if v229 <= int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v233 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[20]))
	if base.I64_extend_i32_u(v229)*int64(1000)+v225 <= v233 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v241 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L11
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[13])))
	if v255|base.B2i32(v233 < base.I64_extend_i32_u(int32(base.Ui32(v229)>>(uint(int32(1))%32)))*int64(1000)+v225) != 0 {
		goto L82
	} else {
		goto L92
	}
L88:
	;
	if v241 == int32(0) {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_logical_read_xlog_page_3), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_logical_read_xlog_page_4), int32(2789), int32(_a_F_logical_read_xlog_page_5))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	goto L1
L92:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15]))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+8))
	v271 = m.T0[v270].(func(*base.Module) int32)(m)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	if v271 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L82
L96:
	;
	v293 = int32(_a_F_logical_read_xlog_page_6)
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[19]))
	if v295 <= int32(0) {
		v331 = v293
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v336 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15]))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v336)+12))
	v338 = m.T0[v337].(func(*base.Module) int32)(m)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L104
	}
L98:
	;
	v299 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[17]))
	if v299 <= int64(0) {
		v331 = v293
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[13])))
	v312 = base.I64_extend_i32_u(int32(base.Ui32(v295)>>(uint((v303^int32(-1))&int32(1))%32)))*int64(1000) + v299
	if v312 <= v292 {
		v330 = int32(0)
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v331 = v330
	goto L97
L101:
	;
	goto L100
L102:
	;
	v318 = v312 - v292
	if base.B2i32(int64(0) < v292)^base.B2i32(v318 < v312)|base.B2i32(int64(2147483646000) < v318) != 0 {
		v330 = int32(2147483647)
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v327 = base.I64_div_s(v318+int64(999), int64(1000))
	v330 = base.I32_wrap_i64(v327)
	goto L101
L104:
	;
	if v338 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v340 = int32(6)
	goto L107
L106:
	;
	v340 = int32(2)
	goto L107
L107:
	;
	goto L108
L108:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v292-v58 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L11
	} else {
		goto L112
	}
L110:
	;
	v354 = v58
	goto L111
L111:
	;
	F_WalSndWait(m, v340, v331, v198)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L11
	} else {
		goto L114
	}
L112:
	;
	v352 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	v354 = v292
	goto L111
L114:
	;
	v52 = v198
	v58 = v354
	goto L14
L115:
	;
	goto L2
L116:
	;
	goto L115
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v361))) = int32(1)
	v369 = int32(0)
	v372 = base.AtomicRmwOr32(m, v369, int32(_a_F_logical_read_xlog_page_0), v369)
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	if v373 == v369 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	if v376 == int32(0) {
		goto L116
	} else {
		goto L119
	}
L119:
	;
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[21]))
	if v380 == v376 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v382 = m.G0
	v384 = v382 - int32(16)
	m.G0 = v384
	v387 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[22]))
	if v387 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	v410 = F_pgmem_kill(m, v376, int32(23))
	mBase = m.M
	goto L116
L123:
	;
	m.G0 = v384 + int32(16)
	goto L115
L124:
	;
	v390 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v384)+15)) = uint8(v390)
	goto L125
L125:
	;
	v394 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[23]))
	v398 = F_write(m, v394, v384+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v398 {
		goto L123
	} else {
		goto L127
	}
L126:
	;
	goto L123
L127:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[24]))
	if v402 == int32(27) {
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	v432 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])))
	if v432 == int32(1) {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v490 = int32(-1)
	goto L131
L131:
	;
	m.G0 = v15 + int32(48)
	return v490
L132:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[25])) = uint8(v442)
	if v442 != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	v437 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v437)+316))
	v440 = base.B2i32(v438 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])) = uint8(v440)
	v442 = v440
	goto L135
L134:
	;
	v442 = int32(0)
	goto L135
L135:
	;
	goto L132
L136:
	;
	F_XLogReadDetermineTimeline(m, l0, l1, l2, v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L11
	} else {
		goto L142
	}
L137:
	;
	v446 = F_GetXLogReplayRecPtr(m, v15+int32(4))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L11
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)+308))
	goto L141
L140:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v453 = v448
	goto L136
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v451
	v453 = v451
	goto L136
L142:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
	*(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[26])) = v457
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[27])) = uint8(base.B2i32(v457 != v460))
	v464 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[28])) = v464
	v467 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1240))
	*(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[29])) = v467
	if base.Ui64(l1-int64(-8192)) <= base.Ui64(v427) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v475 = int32(_a_F_logical_read_xlog_page_7)
	goto L145
L144:
	;
	v475 = base.I32_wrap_i64(v427 - l1)
	goto L145
L145:
	;
	v477 = v15 + int32(8)
	v478 = F_WALRead(m, l0, l4, l1, v475, v460, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L11
	} else {
		goto L146
	}
L146:
	;
	if v478 == int32(0) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	F_WALReadRaiseError(m, v477)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L11
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v484 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	v485 = base.I64_div_u_s(l1, v484)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
	F_CheckXLogRemoved(m, v485, v486)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L11
	} else {
		goto L151
	}
L150:
	;
	goto L149
L151:
	;
	v490 = v475
	goto L131
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
