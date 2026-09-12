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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v267 int32
	_ = v267
	var v268 int64
	_ = v268
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	v3 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v25 = v3
	v26 = int32(-1)
	v27 = v3
	v29 = v3
	v30 = v20
	v31 = v3
	v32 = v3
	v33 = v3
	v34 = v3
	v35 = v3
	v36 = v3
	v37 = v3
	goto L1
L1:
	;
	if v26 != int32(1) {
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
	v42 = int32(16)
	v43 = v30 - v42
	m.G0 = v43
	v46 = v43 - v42
	m.G0 = v46
	v49 = v46 - int32(160)
	m.G0 = v49
	v52 = v49 - v42
	m.G0 = v52
	v55 = v52 - v42
	m.G0 = v55
	v58 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v73 = v25
	v74 = v27
	v75 = v29
	v76 = v30
	v77 = v31
	v78 = v32
	v79 = v33
	v80 = v34
	v81 = v35
	v82 = v36
	v83 = v37
	goto L5
L5:
	;
	goto L11
L6:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v59)
	goto L8
L7:
	;
	goto L8
L8:
	;
	v61 = int32(0)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	v67 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v20 + int32(12)
	goto L9
L9:
	;
	v73 = v61
	v74 = v55
	v75 = v46
	v76 = v55
	v77 = v43
	v78 = v52
	v79 = v49
	v80 = v65
	v81 = v67
	v82 = v58
	v83 = base.B2i32(l1 == v61)
	goto L5
L10:
	;
	goto L2
L11:
	;
	if v73 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	v267 = int32(m.ExcTag)
	v268 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v267 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = int32(396)
	v95 = int32(0)
	v100 = F_CreateDecodingContext(m, int64(0), v95, int32(1), v78, v95, v95, v95)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v80
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v81
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L13
	} else {
		goto L59
	}
L17:
	;
	F_WaitForStandbyConfirmation(m, l0)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v106 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v106)+104))
	F_XLogBeginRead(m, v104, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v112)+40))
	if base.Ui64(v113) < base.Ui64(l0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v118 = v112
	goto L24
L22:
	;
	goto L23
L23:
	;
	if v83 != 0 {
		goto L42
	} else {
		goto L43
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74))) = int32(0)
	v134 = F_XLogReadRecord(m, v118, v74)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v136 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v134 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v141
	F_errmsg_internal(m, int32(196826), v20)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(489940), int32(2138), int32(348224))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	goto L10
L33:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	F_LogicalDecodingProcessRecord(m, v100, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v155 != 0 {
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
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v158)+40))
	if base.Ui64(v159) < base.Ui64(l0) {
		v118 = v158
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
	*(*int32)(unsafe.Add(mBase, _consts[10])) = v82
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v186)+40))
	if v187 != int64(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v179 != int32(2) {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v182 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v182)
	goto L42
L45:
	;
	F_LogicalConfirmReceivedLocation(m, l0)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L13
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v195)+120))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v100)+52))
	if v197 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = int32(240447)
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v75)+4)) = int32(992)
	v206 = int32(4463656)
	v207 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v207
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v75
	v211 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+164)) = uint8(v211)
	*(*uint8)(unsafe.Add(mBase, uint32(v100)+147)) = uint8(v211)
	m.T0[v197].(func(*base.Module, int32))(m, v100)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L13
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	F_ReorderBufferFree(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L13
	} else {
		goto L54
	}
L53:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v218
	goto L52
L54:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v100)+16))
	F_FreeSnapshotBuilder(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L13
	} else {
		goto L55
	}
L55:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v100)+8))
	F_XLogReaderFree(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L56
	}
L56:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	F_MemoryContextDelete(m, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L13
	} else {
		goto L57
	}
L57:
	;
	F_InvalidateSystemCachesExtended(m)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L13
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v80
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v81
	m.G0 = v20 + int32(16)
	return v196
L59:
	;
	F_pg_re_throw(m)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	goto L12
L61:
	;
	v272 = int32(v268)
	m.G0 = v76
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v275)))
	if v20+int32(12) == v279 {
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
	if v282 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v282 = v281
	goto L66
L65:
	;
	v282 = int32(0)
	goto L66
L66:
	;
	goto L63
L67:
	;
	F___wasm_longjmp(m, v275, v274)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v25 = v274
	v26 = v282
	v27 = v74
	v29 = v75
	v30 = v76
	v31 = v77
	v32 = v78
	v33 = v79
	v34 = v80
	v35 = v81
	v36 = v82
	v37 = v83
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
	var v34 int64
	_ = v34
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
			v34 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v34
			*(*int64)(unsafe.Add(mBase, uint32(v25)+52)) = v34
			*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = int32(1073741823)
			*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v27
			*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v27
			*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v34
			*(*int64)(unsafe.Add(mBase, uint32(v25)+60)) = v34
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
				v34 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v25)+32)) = v34
				*(*int64)(unsafe.Add(mBase, uint32(v25)+52)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = int32(1073741823)
				*(*int64)(unsafe.Add(mBase, uint32(v25)+16)) = v27
				*(*int64)(unsafe.Add(mBase, uint32(v25)+24)) = v27
				*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = v34
				*(*int64)(unsafe.Add(mBase, uint32(v25)+60)) = v34
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
				F_errmsg_internal(m, int32(127717), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491141), int32(690), int32(350301))
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
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
	return v60
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
	v41 = v33
	v42 = v31
	goto L15
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v45 = v41 - v42
	if base.Ui32(v45) < base.Ui32(v29) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if v35 == int32(0) {
		v60 = v32
		goto L12
	} else {
		goto L17
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v41 = v40
	v42 = v39
	goto L15
L18:
	;
	v47 = v45
	goto L20
L19:
	;
	v47 = v29
	goto L20
L20:
	;
	if v47 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v51 = v50 + v47
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v51
	v53 = v47 + v32
	v55 = v29 - v47
	if v55 != 0 {
		v28 = v49 + v47
		v29 = v55
		v31 = v51
		v32 = v53
		goto L10
	} else {
		goto L25
	}
L22:
	;
	v48 = F__emscripten_memcpy_bulkmem(m, v28, v43+v42, v47)
	mBase = m.M
	v49 = v48
	goto L24
L23:
	;
	v49 = v28
	goto L24
L24:
	;
	goto L21
L25:
	;
	v60 = v53
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
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
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
		v16 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(32)
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)) = uint8(v20)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v16
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
			if l1 == int32(0) {
				if l1 != 0 {
					v41 = base.I32_extend16_s(l2)
					if int32(0) <= v41 {
						v51 = v41
						v52 = int32(0)
					} else {
						v46 = int32(45)
						*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v46)
						v51 = int32(0) - v41
						v52 = int32(1)
					}
					v54 = F_pg_ultoa_n(m, v51, v9+v52)
					mBase = m.M
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v9+(v54+v52)))) = uint8(v57)
					v59 = F_BufFileCreateFileSet(m, l1, v9)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v59
						m.G0 = v9 + int32(1024)
						return v12
					}
				} else {
					v63 = F_BufFileCreateTemp(m, int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v63
						m.G0 = v9 + int32(1024)
						return v12
					}
				}
			} else {
				if l2 != int32(-1) {
					if l1 != 0 {
						v41 = base.I32_extend16_s(l2)
						if int32(0) <= v41 {
							v51 = v41
							v52 = int32(0)
						} else {
							v46 = int32(45)
							*(*uint8)(unsafe.Add(mBase, uint32(v9))) = uint8(v46)
							v51 = int32(0) - v41
							v52 = int32(1)
						}
						v54 = F_pg_ultoa_n(m, v51, v9+v52)
						mBase = m.M
						v57 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v9+(v54+v52)))) = uint8(v57)
						v59 = F_BufFileCreateFileSet(m, l1, v9)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v59
							m.G0 = v9 + int32(1024)
							return v12
						}
					} else {
						v63 = F_BufFileCreateTemp(m, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v63
							m.G0 = v9 + int32(1024)
							return v12
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
					m.G0 = v9 + int32(1024)
					return v12
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
	var v57 int64
	_ = v57
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
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int64
	_ = v212
	var v216 int32
	_ = v216
	var v220 int64
	_ = v220
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int64
	_ = v285
	var v289 int32
	_ = v289
	var v298 int64
	_ = v298
	var v301 int32
	_ = v301
	var v304 int64
	_ = v304
	var v312 int64
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int64
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v363 int64
	_ = v363
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v382 int64
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int64
	_ = v400
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v437 int32
	_ = v437
	v10 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v18 = l1 + base.I64_extend_i32_s(l2)
	v21 = *(*int64)(unsafe.Add(mBase, _consts[571]))
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
	v437 = m.ExcPending
	if v437 != 0 {
		goto L11
	} else {
		goto L141
	}
L2:
	;
	v363 = *(*int64)(unsafe.Add(mBase, _consts[571]))
	if base.Ui64(v18) <= base.Ui64(v363) {
		goto L118
	} else {
		goto L119
	}
L3:
	;
	v52 = v46
	v57 = v10
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
	v27 = *(*int32)(unsafe.Add(mBase, _consts[572]))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[573]))
	if v29 == int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[517]))
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
	v60 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(0)
	goto L16
L15:
	;
	v346 = *(*int32)(unsafe.Add(mBase, _consts[83]))
	F_SetLatch(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L11
	} else {
		goto L117
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[44]))
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
	v68 = *(*int32)(unsafe.Add(mBase, _consts[574]))
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
	*(*int32)(unsafe.Add(mBase, _consts[574])) = int32(0)
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
	v80 = *(*int32)(unsafe.Add(mBase, _consts[572]))
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
	v88 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
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
	v128 = *(*int32)(unsafe.Add(mBase, _consts[572]))
	if v128 != 0 {
		goto L48
	} else {
		goto L49
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, _consts[571])) = v125
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
	v93 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+316))
	v96 = base.B2i32(v94 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v96)
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
	v104 = int32(4366320)
	v105 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v105)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v105)+280)) = v106
	*(*int64)(unsafe.Add(mBase, _consts[118])) = v106
	v111 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v111)+272)) = v112
	*(*int64)(unsafe.Add(mBase, _consts[117])) = v112
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
	v121 = *(*int64)(unsafe.Add(mBase, _consts[118]))
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
	v130 = *(*int32)(unsafe.Add(mBase, _consts[572]))
	v132 = *(*int64)(unsafe.Add(mBase, _consts[571]))
	v134 = *(*int32)(unsafe.Add(mBase, _consts[573]))
	if v134 == int32(0) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v149 = v52
	goto L50
L50:
	;
	v152 = *(*int64)(unsafe.Add(mBase, _consts[575]))
	v154 = *(*int32)(unsafe.Add(mBase, _consts[212]))
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)+32))
	if base.Ui64(v152) <= base.Ui64(v155) {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[517]))
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
	v149 = int32(100663302)
	goto L50
L58:
	;
	if v128 != 0 {
		v187 = v149
		goto L63
	} else {
		goto L64
	}
L59:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v154)+24))
	if base.Ui64(v152) <= base.Ui64(v157) {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, _consts[576])))
	if v160 != 0 {
		goto L58
	} else {
		goto L61
	}
L61:
	;
	F_WalSndKeepalive(m, int32(0), int64(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	goto L58
L63:
	;
	v189 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[577])) = uint8(v189)
	v192 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v194 = m.T0[v193].(func(*base.Module) int32)(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L11
	} else {
		goto L75
	}
L64:
	;
	v166 = *(*int64)(unsafe.Add(mBase, _consts[571]))
	if base.Ui64(v166) < base.Ui64(v18) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v187 = int32(100663303)
	goto L63
L66:
	;
	goto L67
L67:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _consts[572]))
	v172 = *(*int32)(unsafe.Add(mBase, _consts[573]))
	if v172 == int32(0) {
		goto L47
	} else {
		goto L68
	}
L68:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _consts[517]))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+202)))
	if v177 != int32(1) {
		goto L47
	} else {
		goto L69
	}
L69:
	;
	if v170 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v182 = int32(21)
	goto L72
L71:
	;
	v182 = int32(19)
	goto L72
L72:
	;
	v183 = F_StandbySlotsHaveCaughtup(m, v166, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	if v183 != 0 {
		goto L47
	} else {
		goto L74
	}
L74:
	;
	v187 = int32(100663302)
	goto L63
L75:
	;
	if v194 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, _consts[578])))
	if v197 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v212 = *(*int64)(unsafe.Add(mBase, _consts[579]))
	if v212 <= int64(0) {
		goto L82
	} else {
		goto L83
	}
L78:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, _consts[580])))
	if v201 == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	v207 = m.T0[v206].(func(*base.Module) int32)(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	if v207 == int32(0) {
		goto L47
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	v264 = m.G0
	v265 = int32(16)
	v266 = v264 - v265
	m.G0 = v266
	F___gettimeofday(m, v266)
	mBase = m.M
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v266)))
	v270 = int64(*(*int32)(unsafe.Add(mBase, uint32(v266)+8)))
	m.G0 = v266 + v265
	v278 = v270 + v269*int64(1000000) - int64(946684800000000)
	goto L97
L83:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	if v216 <= int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v220 = *(*int64)(unsafe.Add(mBase, _consts[582]))
	if base.I64_extend_i32_u(v216)*int64(1000)+v212 <= v220 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v228 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L11
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, _consts[576])))
	if v242 != 0 {
		goto L82
	} else {
		goto L92
	}
L88:
	;
	if v228 == int32(0) {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(65395), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(487926), int32(2789), int32(66955))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	goto L1
L92:
	;
	if v220 < base.I64_extend_i32_u(int32(base.Ui32(v216)>>(uint(int32(1))%32)))*int64(1000)+v212 {
		goto L82
	} else {
		goto L93
	}
L93:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v257 = m.T0[v256].(func(*base.Module) int32)(m)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	if v257 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L82
L97:
	;
	v279 = int32(10000)
	v281 = *(*int32)(unsafe.Add(mBase, _consts[581]))
	if v281 <= int32(0) {
		v316 = v279
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[219]))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+12))
	v323 = m.T0[v322].(func(*base.Module) int32)(m)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L11
	} else {
		goto L106
	}
L99:
	;
	v285 = *(*int64)(unsafe.Add(mBase, _consts[579]))
	if v285 <= int64(0) {
		v316 = v279
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, _consts[576])))
	v298 = base.I64_extend_i32_u(int32(base.Ui32(v281)>>(uint((v289^int32(-1))&int32(1))%32)))*int64(1000) + v285
	if v298 <= v278 {
		v315 = int32(0)
		goto L102
	} else {
		goto L103
	}
L101:
	;
	v316 = v315
	goto L98
L102:
	;
	goto L101
L103:
	;
	v301 = int32(2147483647)
	v304 = v298 - v278
	if base.B2i32(int64(0) < v278)^base.B2i32(v304 < v298) != 0 {
		v315 = v301
		goto L102
	} else {
		goto L104
	}
L104:
	;
	if int64(2147483646000) < v304 {
		v315 = v301
		goto L102
	} else {
		goto L105
	}
L105:
	;
	v312 = base.I64_div_s(v304+int64(999), int64(1000))
	v315 = base.I32_wrap_i64(v312)
	goto L102
L106:
	;
	if v323 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v325 = int32(6)
	goto L109
L108:
	;
	v325 = int32(2)
	goto L109
L109:
	;
	goto L110
L110:
	;
	if base.I64_extend_i32_s(int32(1000))*int64(1000) <= v278-v57 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	F_pgstat_flush_io(m, int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L11
	} else {
		goto L114
	}
L112:
	;
	v339 = v57
	goto L113
L113:
	;
	F_WalSndWait(m, v325, v316, v187)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L11
	} else {
		goto L116
	}
L114:
	;
	v337 = F_pgstat_flush_backend(m, int32(0), int32(1))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	v339 = v278
	goto L113
L116:
	;
	v52 = v187
	v57 = v339
	goto L14
L117:
	;
	goto L2
L118:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, _consts[112])))
	if v368 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v427 = int32(-1)
	goto L120
L120:
	;
	m.G0 = v15 + int32(48)
	return v427
L121:
	;
	*(*uint8)(unsafe.Add(mBase, _consts[583])) = uint8(v378)
	if v378 != 0 {
		goto L126
	} else {
		goto L127
	}
L122:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+316))
	v376 = base.B2i32(v374 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[112])) = uint8(v376)
	v378 = v376
	goto L124
L123:
	;
	v378 = int32(0)
	goto L124
L124:
	;
	goto L121
L125:
	;
	F_XLogReadDetermineTimeline(m, l0, l1, l2, v389)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L11
	} else {
		goto L131
	}
L126:
	;
	v382 = F_GetXLogReplayRecPtr(m, v15+int32(4))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L11
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v386 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v386)+308))
	goto L130
L129:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v389 = v384
	goto L125
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v387
	v389 = v387
	goto L125
L131:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1224))
	*(*int32)(unsafe.Add(mBase, _consts[584])) = v393
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	*(*uint8)(unsafe.Add(mBase, _consts[585])) = uint8(base.B2i32(v393 != v396))
	v400 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1232))
	*(*int64)(unsafe.Add(mBase, _consts[586])) = v400
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1240))
	*(*int32)(unsafe.Add(mBase, _consts[587])) = v403
	if base.Ui64(l1-int64(-8192)) <= base.Ui64(v363) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v411 = int32(8192)
	goto L134
L133:
	;
	v411 = base.I32_wrap_i64(v363 - l1)
	goto L134
L134:
	;
	v414 = F_WALRead(m, l0, l4, l1, v411, v396, v15+int32(8))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L11
	} else {
		goto L135
	}
L135:
	;
	if v414 == int32(0) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	F_WALReadRaiseError(m, v15+int32(8))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L11
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v422 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+1160)))
	v423 = base.I64_div_u_s(l1, v422)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1184))
	F_CheckXLogRemoved(m, v423, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L11
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	v427 = v411
	goto L120
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
