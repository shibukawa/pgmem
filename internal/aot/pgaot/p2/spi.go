package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SPI_cursor_open(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 <= v2 {
		v12 = int32(0)
		v15 = F_SPI_cursor_open_internal(m, v12, l0, v12, int32(1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v15
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_makeParamList(m, v9)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v26 = v2
			for {
				v35 = v21 + int32(32) + v26*int32(12)
				v37 = v26 << (uint(int32(2)) % 32)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = v38
				v40 = int32(1)
				*(*uint16)(unsafe.Add(mBase, uint32(v35)+6)) = uint16(v40)
				v42 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v35)+4)) = uint8(v42)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v20+v37)))
				*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v45
				v48 = v26 + v40
				if v48 != v9 {
					v26 = v48
					continue
				} else {
					break
				}
				break
			}
			v52 = F_SPI_cursor_open_internal(m, int32(0), l0, v21, int32(1))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v21)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					return v52
				}
			}
		}
	}
}
func F_SPI_cursor_open_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L27
	} else {
		goto L99
	}
L2:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v328 != int32(1) {
		goto L1
	} else {
		goto L90
	}
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[0])) = v35
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	if v37 == v35 {
		v326 = v22
		goto L2
	} else {
		goto L15
	}
L4:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v15 == int32(569278163) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[0])) = int32(-6)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v29 = v21
	goto L4
L9:
	;
	goto L8
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 == int32(1) {
		goto L3
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[0])) = int32(0)
	v29 = v22
	goto L4
L13:
	;
	goto L12
L14:
	;
	v326 = v29
	goto L2
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[1]))
	if v41 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L27
	} else {
		goto L85
	}
L17:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[2]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	goto L20
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L27
	} else {
		goto L82
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v45
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+24))
	v51 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[3])) = v51
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[4])) = v49
	v56 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[5])) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = v51
	if l0 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v74 = F_MemoryContextStrdup(m, v72, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L27
	} else {
		goto L30
	}
L22:
	;
	v67 = int32(0)
	v69 = F_CreatePortal(m, l0, v67, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L27
	} else {
		goto L29
	}
L23:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v62 != 0 {
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v63 = F_CreateNewPortal(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	return int32(0)
L28:
	;
	v71 = v63
	goto L21
L29:
	;
	v71 = v69
	goto L21
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v76
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(779)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v78
	v82 = int32(_a_F_SPI_cursor_open_internal_0)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[6])) = v13 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v13 + int32(40)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[1]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+36))
	v96 = F_GetCachedPlan(m, v42, l2, int32(0), v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v99 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v102 = int32(_a_F_SPI_cursor_open_internal_1)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[4]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[4])) = v105
	v107 = F_copyObjectImpl(m, v98)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L27
	} else {
		goto L35
	}
L33:
	;
	v115 = v98
	v116 = v96
	goto L34
L34:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v42)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+40)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v71)+32)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+80)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v71)+60)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v71)+56)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v71)+36)) = v119
	goto L37
L35:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[4])) = v103
	F_ReleaseCachedPlan(m, v96, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L27
	} else {
		goto L36
	}
L36:
	;
	v115 = v107
	v116 = int32(0)
	goto L34
L37:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = v130
	if v130&int32(6) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if v115 == int32(0) {
		v155 = v130
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v163 = v130
	goto L40
L40:
	;
	v165 = int32(0)
	if base.B2i32(v115 == v165)|base.B2i32(v163&int32(2) == v165) != 0 {
		goto L49
	} else {
		goto L50
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+76)) = v161
	v163 = v161
	goto L40
L42:
	;
	v161 = v155 | int32(4)
	goto L41
L43:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v138 != int32(1) {
		v155 = v130
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+4))
	if v143 == int32(6) {
		v155 = v130
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142)+72))
	if v146 != 0 {
		v155 = v130
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v142)+36))
	v148 = F_ExecSupportsBackwardScan(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L27
	} else {
		goto L47
	}
L47:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v71)+76))
	if v148 == int32(0) {
		v155 = v150
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v161 = v150 | int32(2)
	goto L41
L49:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[1]))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+68)) = v184
	if l3 != 0 {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v172 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
	if v177 == int32(6) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+72))
	if v180 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	if l2 != 0 {
		goto L76
	} else {
		goto L77
	}
L55:
	;
	if v115 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L27
	} else {
		goto L74
	}
L58:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[7]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)))
	goto L73
L59:
	;
	v188 = int32(0)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v189 <= v188 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v193 = v188
	goto L61
L61:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v193<<(uint(int32(2))%32))))
	v207 = F_CommandIsReadOnly(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L27
	} else {
		goto L63
	}
L62:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L27
	} else {
		goto L68
	}
L63:
	;
	if v207 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v210 = v193 + int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v210 < v211 {
		v193 = v210
		goto L61
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L62
L67:
	;
	goto L58
L68:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L27
	} else {
		goto L69
	}
L69:
	;
	v220 = F_CreateCommandName(m, v206)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L27
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v220
	F_errmsg(m, int32(_a_F_SPI_cursor_open_internal_2), v13+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L27
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_SPI_cursor_open_internal_3), int32(1745), int32(_a_F_SPI_cursor_open_internal_4))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L27
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	v260 = v245
	goto L54
L74:
	;
	v248 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L27
	} else {
		goto L75
	}
L75:
	;
	v260 = v248
	goto L54
L76:
	;
	v262 = int32(_a_F_SPI_cursor_open_internal_1)
	v263 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[4]))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[4])) = v265
	v267 = F_copyParamList(m, l2)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L27
	} else {
		goto L79
	}
L77:
	;
	v271 = int32(0)
	goto L78
L78:
	;
	F_PortalStart(m, v71, v271, int32(0), v260)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L27
	} else {
		goto L80
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[4])) = v263
	v271 = v267
	goto L78
L80:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[6])) = v277
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[1]))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_SPI_cursor_open_internal[4])) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v281)+12)) = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v281)+24))
	F_MemoryContextReset(m, v286)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L27
	} else {
		goto L81
	}
L81:
	;
	m.G0 = v13 + int32(48)
	return v71
L82:
	;
	F_errmsg_internal(m, int32(_a_F_SPI_cursor_open_internal_5), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L27
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_SPI_cursor_open_internal_3), int32(1620), int32(_a_F_SPI_cursor_open_internal_4))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L27
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L27
	} else {
		goto L86
	}
L86:
	;
	F_errmsg(m, int32(_a_F_SPI_cursor_open_internal_6), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L27
	} else {
		goto L87
	}
L87:
	;
	F_errdetail(m, int32(_a_F_SPI_cursor_open_internal_7), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L27
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_SPI_cursor_open_internal_3), int32(1719), int32(_a_F_SPI_cursor_open_internal_4))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L27
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v326)+12))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v333)+16))
	if v334 != int32(179) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v334<<(uint(int32(3))%32))+uint32(_c_F_SPI_cursor_open_internal[8])))
	goto L94
L92:
	;
	v340 = int32(_a_F_SPI_cursor_open_internal_8)
	goto L93
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L27
	} else {
		goto L95
	}
L94:
	;
	v340 = v339
	goto L93
L95:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L27
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v340
	F_errmsg(m, int32(_a_F_SPI_cursor_open_internal_9), v13)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L27
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_SPI_cursor_open_internal_3), int32(1612), int32(_a_F_SPI_cursor_open_internal_4))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L27
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errcode(m, int32(17170564))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L27
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_SPI_cursor_open_internal_10), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L27
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_SPI_cursor_open_internal_3), int32(1602), int32(_a_F_SPI_cursor_open_internal_4))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L27
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SPI_exec(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_SPI_execute(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_SPI_finish(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[0]))
	if v3 == int32(0) {
		return int32(-4)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
		*(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[1])) = v9
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
		F_MemoryContextDelete(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(0)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
			F_MemoryContextDelete(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = int32(_a_F_SPI_finish_0)
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[0]))
				v25 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v25
				v28 = *(*int64)(unsafe.Add(mBase, uint32(v24)+48))
				*(*int64)(unsafe.Add(mBase, _c_F_SPI_finish[2])) = v28
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[3])) = v31
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v24)+60))
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[4])) = v34
				v36 = int32(_a_F_SPI_finish_1)
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[5]))
				v40 = v38 - int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[5])) = v40
				v44 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[6]))
				if v25 <= v40 {
					v51 = v44 + v40<<(uint(int32(6))%32)
				} else {
					v51 = v25
				}
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_finish[0])) = v51
				return int32(2)
			}
		}
	}
}
func F_SPI_freeplan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v6 != int32(569278163) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v9 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_MemoryContextDelete(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L12
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = v2
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v17<<(uint(int32(2))%32))))
	F_DropCachedPlan(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	return
L10:
	;
	v26 = v17 + int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	if v26 < v27 {
		v17 = v26
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	goto L1
}
func F_SPI_palloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_palloc[0]))
	if v4 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_SPI_palloc_0), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_SPI_palloc_1), int32(1341), int32(_a_F_SPI_palloc_2))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
		v23 = F_MemoryContextAlloc(m, v22, l0)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	}
}
func F_SPI_result_code_string(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = l0 + int32(13)
	if base.B2i32(base.Ui32(int32(33)) <= base.Ui32(v9))|base.B2i32(base.I32_wrap_i64(int64(base.Ui64(int64(8589926143))>>(uint(base.I64_extend_i32_u(v9))%64)))&int32(1) == v2) == v2 {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v9<<(uint(int32(2))%32))+uint32(_c_F_SPI_result_code_string[0])))
		v34 = v25
		m.G0 = v6 + int32(16)
		return v34
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		v27 = int32(_a_F_SPI_result_code_string_0)
		v30 = F_pg_sprintf(m, v27, int32(_a_F_SPI_result_code_string_1), v6)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = v27
			m.G0 = v6 + int32(16)
			return v34
		}
	}
}
func F_SPI_sql_row_to_xmlelement(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(_a_F_SPI_sql_row_to_xmlelement_0)
		F_appendStringInfo(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_1), v7+int32(32))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			F_appendStringInfoString(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_2))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v19 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l2
					F_appendStringInfo(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_3), v7+int32(16))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						F_appendStringInfoString(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_4))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[0]))
							v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
							if int32(0) < v32 {
								v59 = v31
								v61 = int32(0)
								*(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[1])) = v61
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
								if v63 <= v61 {
									*(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[1])) = int32(-9)
									v76 = F_map_sql_identifier_to_xml_name(m)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v74 = F_pstrdup(m, v59+v63<<(uint(int32(4))%32)+int32(24))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return
									} else {
										v76 = F_map_sql_identifier_to_xml_name(m)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_SPI_sql_row_to_xmlelement_0)
								F_appendStringInfo(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_5), v7)
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_appendStringInfoChar(m, l0, int32(10))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										m.G0 = v7 + int32(48)
										return
									}
								}
							}
						}
					}
				} else {
					F_appendStringInfoString(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_4))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[0]))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						if int32(0) < v32 {
							v59 = v31
							v61 = int32(0)
							*(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[1])) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
							if v63 <= v61 {
								*(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[1])) = int32(-9)
								v76 = F_map_sql_identifier_to_xml_name(m)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v74 = F_pstrdup(m, v59+v63<<(uint(int32(4))%32)+int32(24))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return
								} else {
									v76 = F_map_sql_identifier_to_xml_name(m)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(_a_F_SPI_sql_row_to_xmlelement_0)
							F_appendStringInfo(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_5), v7)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								F_appendStringInfoChar(m, l0, int32(10))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									m.G0 = v7 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_appendStringInfoString(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_6))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[0]))
			v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
			if int32(0) < v49 {
				v59 = v48
				v61 = int32(0)
				*(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[1])) = v61
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				if v63 <= v61 {
					*(*int32)(unsafe.Add(mBase, _c_F_SPI_sql_row_to_xmlelement[1])) = int32(-9)
					v76 = F_map_sql_identifier_to_xml_name(m)
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v74 = F_pstrdup(m, v59+v63<<(uint(int32(4))%32)+int32(24))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return
					} else {
						v76 = F_map_sql_identifier_to_xml_name(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				F_appendStringInfoString(m, l0, int32(_a_F_SPI_sql_row_to_xmlelement_7))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					m.G0 = v7 + int32(48)
					return
				}
			}
		}
	}
}
func F__SPI_cursor_operation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	if l0 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[0]))
		if v9 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F__SPI_cursor_operation_0), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F__SPI_cursor_operation_1), int32(3018), int32(_a_F__SPI_cursor_operation_2))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[1]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v16 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v14
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
			v20 = int64(0)
			*(*int64)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[2])) = v20
			*(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[3])) = v18
			v25 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[4])) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v25
			*(*int64)(unsafe.Add(mBase, uint32(v16))) = v20
			v31 = F_PortalRunFetch(m, l0, l1, l2, l3)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[0]))
				*(*int64)(unsafe.Add(mBase, uint32(v34))) = v31
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
				if v37 == int32(5) {
					if v36 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(_a_F__SPI_cursor_operation_3), int32(0))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F__SPI_cursor_operation_1), int32(3043), int32(_a_F__SPI_cursor_operation_2))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v42 = *(*int64)(unsafe.Add(mBase, uint32(v36)+8))
						if v31 != v42 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F__SPI_cursor_operation_3), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F__SPI_cursor_operation_1), int32(3043), int32(_a_F__SPI_cursor_operation_2))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[4])) = v36
							*(*int64)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[2])) = v31
							v48 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v48
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
							*(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[3])) = v51
							*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v48
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
							F_MemoryContextReset(m, v55)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[4])) = v36
					*(*int64)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[2])) = v31
					v48 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v48
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
					*(*int32)(unsafe.Add(mBase, _c_F__SPI_cursor_operation[3])) = v51
					*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v48
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
					F_MemoryContextReset(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F__SPI_cursor_operation_4), int32(0))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F__SPI_cursor_operation_1), int32(3014), int32(_a_F__SPI_cursor_operation_2))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
