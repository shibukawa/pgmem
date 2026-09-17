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
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v124 int32
	_ = v124
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int64
	_ = v215
	var v219 int32
	_ = v219
	var v223 int64
	_ = v223
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v282 int64
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int64
	_ = v289
	var v293 int32
	_ = v293
	var v302 int64
	_ = v302
	var v308 int64
	_ = v308
	var v317 int64
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int64
	_ = v344
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v368 int64
	_ = v368
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v387 int64
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v405 int64
	_ = v405
	var v408 int32
	_ = v408
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int64
	_ = v425
	var v426 int64
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v441 int32
	_ = v441
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
	v441 = m.ExcPending
	if v441 != 0 {
		goto L11
	} else {
		goto L139
	}
L2:
	;
	v368 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	if base.Ui64(v18) <= base.Ui64(v368) {
		goto L116
	} else {
		goto L117
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
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
	goto L16
L15:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[4]))
	F_SetLatch(m, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L11
	} else {
		goto L115
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[5]))
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L11
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[6]))
	if v68 != 0 {
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
	v74 = m.ExcPending
	if v74 != 0 {
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
	v78 = m.ExcPending
	if v78 != 0 {
		goto L11
	} else {
		goto L26
	}
L24:
	;
	F_SyncRepInitConfig(m)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v81 = F_XLogBackgroundFlush(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])))
	if v88 == int32(1) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	if v128 != 0 {
		goto L48
	} else {
		goto L49
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0])) = v125
	goto L33
L35:
	;
	if v98 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+316))
	v96 = base.B2i32(v94 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])) = uint8(v96)
	v98 = v96
	goto L38
L37:
	;
	v98 = int32(0)
	goto L38
L38:
	;
	goto L35
L39:
	;
	v104 = int32(_a_F_logical_read_xlog_page_0)
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v105)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+280)) = v106
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9])) = v106
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v111)+272)) = v112
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[10])) = v112
	goto L44
L40:
	;
	goto L41
L41:
	;
	v123 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L11
	} else {
		goto L46
	}
L42:
	;
	v125 = v121
	goto L34
L44:
	;
	goto L45
L45:
	;
	v121 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[9]))
	goto L42
L46:
	;
	v125 = v123
	goto L34
L47:
	;
	goto L15
L48:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	v132 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[2]))
	if v134 == int32(0) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v148 = v52
	goto L50
L50:
	;
	v151 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[11]))
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[12]))
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v153)+32))
	if base.Ui64(v151) <= base.Ui64(v154) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[3]))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+202)))
	if v139 != int32(1) {
		goto L47
	} else {
		goto L52
	}
L52:
	;
	if v130 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v145 = int32(21)
	goto L55
L54:
	;
	v145 = int32(19)
	goto L55
L55:
	;
	v146 = F_StandbySlotsHaveCaughtup(m, v132, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	if v146 != 0 {
		goto L47
	} else {
		goto L57
	}
L57:
	;
	v148 = int32(100663302)
	goto L50
L58:
	;
	if v128 != 0 {
		v188 = v148
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v153)+24))
	if base.Ui64(v151) <= base.Ui64(v156) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[13])))
	if v159&int32(1) != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	F_WalSndKeepalive(m, int32(0), int64(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[14])) = uint8(v190)
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	v195 = m.T0[v194].(func(*base.Module) int32)(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L11
	} else {
		goto L75
	}
L64:
	;
	v167 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[0]))
	if base.Ui64(v167) < base.Ui64(v18) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v188 = int32(100663303)
	goto L63
L66:
	;
	goto L67
L67:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[1]))
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[2]))
	if v173 == int32(0) {
		goto L47
	} else {
		goto L68
	}
L68:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[3]))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177)+202)))
	if v178 != int32(1) {
		goto L47
	} else {
		goto L69
	}
L69:
	;
	if v171 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v183 = int32(21)
	goto L72
L71:
	;
	v183 = int32(19)
	goto L72
L72:
	;
	v184 = F_StandbySlotsHaveCaughtup(m, v167, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	if v184 != 0 {
		goto L47
	} else {
		goto L74
	}
L74:
	;
	v188 = int32(100663302)
	goto L63
L75:
	;
	if v195 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v198 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[16])))
	if v198 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v215 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[17]))
	if v215 <= int64(0) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[18])))
	if v202&int32(1) == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+12))
	v210 = m.T0[v209].(func(*base.Module) int32)(m)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	if v210 == int32(0) {
		goto L47
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	v268 = m.G0
	v269 = int32(16)
	v270 = v268 - v269
	m.G0 = v270
	F_gettimeofday(m, v270)
	mBase = m.M
	v273 = *(*int64)(unsafe.Add(mBase, uint32(v270)))
	v274 = int64(*(*int32)(unsafe.Add(mBase, uint32(v270)+8)))
	m.G0 = v270 + v269
	v282 = v274 + v273*int64(1000000) - int64(946684800000000)
	goto L96
L83:
	;
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[19]))
	if v219 <= int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v223 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[20]))
	if base.I64_extend_i32_u(v219)*int64(1000)+v215 <= v223 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v231 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L11
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[13])))
	if v245|base.B2i32(v223 < base.I64_extend_i32_u(int32(base.Ui32(v219)>>(uint(int32(1))%32)))*int64(1000)+v215) != 0 {
		goto L82
	} else {
		goto L92
	}
L88:
	;
	if v231 == int32(0) {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(_a_F_logical_read_xlog_page_1), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_logical_read_xlog_page_2), int32(2789), int32(_a_F_logical_read_xlog_page_3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
	v257 = m.ExcPending
	if v257 != 0 {
		goto L11
	} else {
		goto L93
	}
L93:
	;
	v259 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15]))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	v261 = m.T0[v260].(func(*base.Module) int32)(m)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	if v261 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	goto L82
L96:
	;
	v283 = int32(_a_F_logical_read_xlog_page_4)
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[19]))
	if v285 <= int32(0) {
		v321 = v283
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[15]))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v328 = m.T0[v327].(func(*base.Module) int32)(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L11
	} else {
		goto L104
	}
L98:
	;
	v289 = *(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[17]))
	if v289 <= int64(0) {
		v321 = v283
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[13])))
	v302 = base.I64_extend_i32_u(int32(base.Ui32(v285)>>(uint((v293^int32(-1))&int32(1))%32)))*int64(1000) + v289
	if v302 <= v282 {
		v320 = int32(0)
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v321 = v320
	goto L97
L101:
	;
	goto L100
L102:
	;
	v308 = v302 - v282
	if base.B2i32(int64(0) < v282)^base.B2i32(v308 < v302)|base.B2i32(int64(2147483646000) < v308) != 0 {
		v320 = int32(2147483647)
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v317 = base.I64_div_s(v308+int64(999), int64(1000))
	v320 = base.I32_wrap_i64(v317)
	goto L101
L104:
	;
	if v328 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v330 = int32(6)
	goto L107
L106:
	;
	v330 = int32(2)
	goto L107
L107:
	;
	goto L108
L108:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v282-v58 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L112
	}
L110:
	;
	v344 = v58
	goto L111
L111:
	;
	F_WalSndWait(m, v330, v321, v188)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L114
	}
L112:
	;
	v342 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	v344 = v282
	goto L111
L114:
	;
	v52 = v188
	v58 = v344
	goto L14
L115:
	;
	goto L2
L116:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])))
	if v373 == int32(1) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v431 = int32(-1)
	goto L118
L118:
	;
	m.G0 = v15 + int32(48)
	return v431
L119:
	;
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[21])) = uint8(v383)
	if v383 != 0 {
		goto L124
	} else {
		goto L125
	}
L120:
	;
	v378 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v378)+316))
	v381 = base.B2i32(v379 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[7])) = uint8(v381)
	v383 = v381
	goto L122
L121:
	;
	v383 = int32(0)
	goto L122
L122:
	;
	goto L119
L123:
	;
	F_XLogReadDetermineTimeline(m, l0, l1, l2, v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L129
	}
L124:
	;
	v387 = F_GetXLogReplayRecPtr(m, v15+int32(4))
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L11
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v391 = *(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[8]))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+308))
	goto L128
L127:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v394 = v389
	goto L123
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v392
	v394 = v392
	goto L123
L129:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
	*(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[22])) = v398
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*uint8)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[23])) = uint8(base.B2i32(v398 != v401))
	v405 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
	*(*int64)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[24])) = v405
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1240))
	*(*int32)(unsafe.Add(mBase, _c_F_logical_read_xlog_page[25])) = v408
	if base.Ui64(l1-int64(-8192)) <= base.Ui64(v368) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v416 = int32(_a_F_logical_read_xlog_page_5)
	goto L132
L131:
	;
	v416 = base.I32_wrap_i64(v368 - l1)
	goto L132
L132:
	;
	v418 = v15 + int32(8)
	v419 = F_WALRead(m, l0, l4, l1, v416, v401, v418)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L11
	} else {
		goto L133
	}
L133:
	;
	if v419 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	F_WALReadRaiseError(m, v418)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L11
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v425 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	v426 = base.I64_div_u_s(l1, v425)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
	F_CheckXLogRemoved(m, v426, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L11
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v431 = v416
	goto L118
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
