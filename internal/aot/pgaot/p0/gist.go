package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_gistXLogUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	v3 = l2
	v5 = l4
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v5)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+12)) = uint16(v3)
	F_XLogBeginInsert(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_XLogRegisterData(m, v12+int32(12), int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_XLogRegisterBuffer(m, int32(0), l0, int32(8))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_XLogRegisterBufData(m, int32(0), l1, v3<<(uint(int32(1))%32))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if int32(0) < v5 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v43 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	if l5 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3+v43<<(uint(int32(2))%32))))
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+6)))
	F_XLogRegisterBufData(m, int32(0), v49, v50&int32(8191))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v56 = v43 + int32(1)
	if v56 != v5 {
		v43 = v56
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_XLogRegisterBuffer(m, int32(1), l5, int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v73 = F_XLogInsert(m, int32(14), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	m.G0 = v12 + int32(16)
	return v73
}
func F_gist_between(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
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
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v311 int32
	_ = v311
	var v322 int32
	_ = v322
	var v343 int32
	_ = v343
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v21 = l0 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v23&int32(2) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v174 = int32(0)
	if v169 <= v174 {
		goto L45
	} else {
		goto L46
	}
L5:
	;
	if base.Ui32(v31) < base.Ui32(v15) {
		goto L42
	} else {
		goto L43
	}
L6:
	;
	v26 = int32(0)
	goto L8
L7:
	;
	v26 = l2
	goto L8
L8:
	;
	v27 = v21 + v26
	v29 = v23 & int32(1)
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = v21
	goto L11
L10:
	;
	v30 = v27
	goto L11
L11:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+4)))
	if v31 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v38 = l1 + int32(16)
	v40 = v30 + int32(8)
	v44 = v15
	v45 = v31
	goto L13
L13:
	;
	v53 = v40 + int32(2)
	v55 = v38 + int32(23)
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40))))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+20)))
	if base.Ui32(v56) < base.Ui32(v57) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L5
L15:
	;
	v59 = v56
	goto L17
L16:
	;
	v59 = v57
	goto L17
L17:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v59) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if v121 != 0 {
		v169 = v121
		goto L4
	} else {
		goto L36
	}
L19:
	;
	v121 = int32(0)
	goto L18
L20:
	;
	v95 = v90
	v96 = v91
	v97 = v92
	goto L30
L21:
	;
	if (v53|v55)&int32(3) != 0 {
		v90 = v53
		v91 = v55
		v92 = v59
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v83 = v53
	v84 = v55
	v85 = v59
	goto L23
L23:
	;
	if v85 == int32(0) {
		goto L19
	} else {
		goto L29
	}
L24:
	;
	v67 = v53
	v68 = v55
	v69 = v59
	goto L25
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v72 != v73 {
		v90 = v67
		v91 = v68
		v92 = v69
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v83 = v78
	v84 = v76
	v85 = v80
	goto L23
L27:
	;
	v75 = int32(4)
	v76 = v68 + v75
	v78 = v67 + v75
	v80 = v69 - v75
	if base.Ui32(int32(3)) < base.Ui32(v80) {
		v67 = v78
		v68 = v76
		v69 = v80
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v90 = v83
	v91 = v84
	v92 = v85
	goto L20
L30:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v100 == v101 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v121 = v100 - v101
	goto L18
L32:
	;
	v103 = int32(1)
	v108 = v97 - v103
	if v108 != 0 {
		v95 = v95 + v103
		v96 = v96 + v103
		v97 = v108
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	goto L19
L36:
	;
	if v56 != v57 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v169 = v56 - v57
	goto L4
L38:
	;
	goto L39
L39:
	;
	if v45 < int32(2) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v126 = int32(1)
	v130 = int32(131064)
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if v126 < v44 {
		v38 = v38 + (v133+int32(7))&v130
		v40 = v40 + (v56+int32(9))&v130
		v44 = v44 - v126
		v45 = v45 - v126
		goto L13
	} else {
		goto L41
	}
L41:
	;
	goto L14
L42:
	;
	v158 = v31
	goto L44
L43:
	;
	v158 = v15
	goto L44
L44:
	;
	v169 = v158 - v15
	goto L4
L45:
	;
	if v29 != 0 {
		v183 = v21
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v343 = v174
	goto L47
L47:
	;
	return v343
L48:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+4)))
	if v184 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L49:
	;
	if v23&int32(4) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v183 = v27
	goto L48
L51:
	;
	goto L52
L52:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v183 = v27 + int32(base.Ui32(v179)>>(uint(int32(2))%32))
	goto L48
L53:
	;
	v343 = base.B2i32(int32(0) <= v322)
	goto L47
L54:
	;
	if base.Ui32(v184) < base.Ui32(v15) {
		goto L85
	} else {
		goto L86
	}
L55:
	;
	v191 = l1 + int32(16)
	v193 = v183 + int32(8)
	v197 = v15
	v198 = v184
	goto L56
L56:
	;
	v206 = v193 + int32(2)
	v208 = v191 + int32(23)
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193))))
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191)+20)))
	if base.Ui32(v209) < base.Ui32(v210) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L54
L58:
	;
	v212 = v209
	goto L60
L59:
	;
	v212 = v210
	goto L60
L60:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v212) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if v274 != 0 {
		v322 = v274
		goto L53
	} else {
		goto L79
	}
L62:
	;
	v274 = int32(0)
	goto L61
L63:
	;
	v248 = v243
	v249 = v244
	v250 = v245
	goto L73
L64:
	;
	if (v206|v208)&int32(3) != 0 {
		v243 = v206
		v244 = v208
		v245 = v212
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v236 = v206
	v237 = v208
	v238 = v212
	goto L66
L66:
	;
	if v238 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v220 = v206
	v221 = v208
	v222 = v212
	goto L68
L68:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v225 != v226 {
		v243 = v220
		v244 = v221
		v245 = v222
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v236 = v231
	v237 = v229
	v238 = v233
	goto L66
L70:
	;
	v228 = int32(4)
	v229 = v221 + v228
	v231 = v220 + v228
	v233 = v222 - v228
	if base.Ui32(int32(3)) < base.Ui32(v233) {
		v220 = v231
		v221 = v229
		v222 = v233
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v243 = v236
	v244 = v237
	v245 = v238
	goto L63
L73:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v253 == v254 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v274 = v253 - v254
	goto L61
L75:
	;
	v256 = int32(1)
	v261 = v250 - v256
	if v261 != 0 {
		v248 = v248 + v256
		v249 = v249 + v256
		v250 = v261
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L62
L79:
	;
	if v209 != v210 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v322 = v209 - v210
	goto L53
L81:
	;
	goto L82
L82:
	;
	if v198 < int32(2) {
		goto L54
	} else {
		goto L83
	}
L83:
	;
	v279 = int32(1)
	v283 = int32(131064)
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v191))))
	if v279 < v197 {
		v191 = v191 + (v286+int32(7))&v283
		v193 = v193 + (v209+int32(9))&v283
		v197 = v197 - v279
		v198 = v198 - v279
		goto L56
	} else {
		goto L84
	}
L84:
	;
	goto L57
L85:
	;
	v311 = v184
	goto L87
L86:
	;
	v311 = v15
	goto L87
L87:
	;
	v322 = v311 - v15
	goto L53
}
func F_gist_box_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 == v2 {
		v49 = v2
		m.G0 = v9 + int32(16)
		return v49
	} else {
		if v12 == int32(0) {
			v49 = v2
			m.G0 = v9 + int32(16)
			return v49
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+16)))
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22+v23)+12)))
			if v25&int32(1) != 0 {
				v31 = (v11 - int32(1)) & int32(65535)
				if base.Ui32(int32(12)) <= base.Ui32(v31) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
						F_errmsg_internal(m, int32(480638), v9)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(498030), int32(939), int32(92320))
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v31<<(uint(int32(2))%32))+uint32(_consts[22])))
					v40 = F_DirectFunctionCall2Coll(m, v38, int32(0), v17, v12)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v49 = base.B2i32(v40 != int32(0))
						m.G0 = v9 + int32(16)
						return v49
					}
				}
			} else {
				v46 = F_rtree_internal_consistent(m, v17, v12, v11)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v49 = v46
					m.G0 = v9 + int32(16)
					return v49
				}
			}
		}
	}
}
func F_gist_box_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 float64
	_ = v15
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 float64
	_ = v19
	var v27 int32
	_ = v27
	var v33 float64
	_ = v33
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 float64
	_ = v37
	var v45 int32
	_ = v45
	var v51 float64
	_ = v51
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 float64
	_ = v55
	var v63 int32
	_ = v63
	var v69 float64
	_ = v69
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v73 float64
	_ = v73
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == v2 {
		v92 = base.B2i32(v10|v9 == int32(0))
		*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
		return v8
	} else {
		if v9 == int32(0) {
			v92 = base.B2i32(v10|v9 == int32(0))
			*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
			return v8
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
			v17 = int64(9223372036854775807)
			v18 = base.I64_reinterpret_f64(v15) & v17
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
			if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v19)&v17) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(v18) {
					v33 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
					v35 = int64(9223372036854775807)
					v36 = base.I64_reinterpret_f64(v33) & v35
					v37 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
					if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v37)&v35) {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
							v51 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
							v53 = int64(9223372036854775807)
							v54 = base.I64_reinterpret_f64(v51) & v53
							v55 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
							if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v55)&v53) {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
									v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
									v71 = int64(9223372036854775807)
									v72 = base.I64_reinterpret_f64(v69) & v71
									v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
										return v8
									} else {
										v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
										return v8
									}
								} else {
									v63 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v63)
									return v8
								}
							} else {
								if base.F64_ne(v51, v55) != 0 {
									v92 = v2
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
									return v8
								} else {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
										v92 = v2
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
										return v8
									} else {
										v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
										v71 = int64(9223372036854775807)
										v72 = base.I64_reinterpret_f64(v69) & v71
										v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
											return v8
										} else {
											v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
											return v8
										}
									}
								}
							}
						} else {
							v45 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v45)
							return v8
						}
					} else {
						if base.F64_ne(v33, v37) != 0 {
							v92 = v2
							*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
							return v8
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
								v92 = v2
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
								return v8
							} else {
								v51 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
								v53 = int64(9223372036854775807)
								v54 = base.I64_reinterpret_f64(v51) & v53
								v55 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v55)&v53) {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
										v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
										v71 = int64(9223372036854775807)
										v72 = base.I64_reinterpret_f64(v69) & v71
										v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
											return v8
										} else {
											v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
											return v8
										}
									} else {
										v63 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v63)
										return v8
									}
								} else {
									if base.F64_ne(v51, v55) != 0 {
										v92 = v2
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
										return v8
									} else {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
											v92 = v2
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
											return v8
										} else {
											v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
											v71 = int64(9223372036854775807)
											v72 = base.I64_reinterpret_f64(v69) & v71
											v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
												return v8
											} else {
												v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
												return v8
											}
										}
									}
								}
							}
						}
					}
				} else {
					v27 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v27)
					return v8
				}
			} else {
				if base.F64_ne(v15, v19) != 0 {
					v92 = v2
					*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
					return v8
				} else {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v18) {
						v92 = v2
						*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
						return v8
					} else {
						v33 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
						v35 = int64(9223372036854775807)
						v36 = base.I64_reinterpret_f64(v33) & v35
						v37 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v37)&v35) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
								v51 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
								v53 = int64(9223372036854775807)
								v54 = base.I64_reinterpret_f64(v51) & v53
								v55 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v55)&v53) {
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
										v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
										v71 = int64(9223372036854775807)
										v72 = base.I64_reinterpret_f64(v69) & v71
										v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
										if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
											return v8
										} else {
											v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
											return v8
										}
									} else {
										v63 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v63)
										return v8
									}
								} else {
									if base.F64_ne(v51, v55) != 0 {
										v92 = v2
										*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
										return v8
									} else {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
											v92 = v2
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
											return v8
										} else {
											v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
											v71 = int64(9223372036854775807)
											v72 = base.I64_reinterpret_f64(v69) & v71
											v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
												return v8
											} else {
												v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
												return v8
											}
										}
									}
								}
							} else {
								v45 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v45)
								return v8
							}
						} else {
							if base.F64_ne(v33, v37) != 0 {
								v92 = v2
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
								return v8
							} else {
								if base.Ui64(int64(9218868437227405312)) < base.Ui64(v36) {
									v92 = v2
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
									return v8
								} else {
									v51 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
									v53 = int64(9223372036854775807)
									v54 = base.I64_reinterpret_f64(v51) & v53
									v55 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
									if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v55)&v53) {
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
											v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
											v71 = int64(9223372036854775807)
											v72 = base.I64_reinterpret_f64(v69) & v71
											v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
											if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
												return v8
											} else {
												v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
												return v8
											}
										} else {
											v63 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v63)
											return v8
										}
									} else {
										if base.F64_ne(v51, v55) != 0 {
											v92 = v2
											*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
											return v8
										} else {
											if base.Ui64(int64(9218868437227405312)) < base.Ui64(v54) {
												v92 = v2
												*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v92)
												return v8
											} else {
												v69 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
												v71 = int64(9223372036854775807)
												v72 = base.I64_reinterpret_f64(v69) & v71
												v73 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
												if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v73)&v71) {
													*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v72)))
													return v8
												} else {
													v86 = base.B2i32(base.Ui64(v72) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v69, v73)
													*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v86)
													return v8
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_gist_circle_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
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
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	if base.Ui32(int32(20)) <= base.Ui32(v9) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
			F_errmsg_internal(m, int32(480638), v7)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(498030), int32(1492), int32(415906))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v32 = F_computeDistance(m, int32(0), v30, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v34)
			v36 = F_Float8GetDatum(m, v32)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v36
			}
		}
	}
}
func F_gist_identify(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = int32(0)
	if v2 <= base.I32_extend8_s(l0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_consts[57])))
		v13 = v12
	} else {
		v13 = v2
	}
	return v13
}
func F_gist_point_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	if base.Ui32(v9) <= base.Ui32(int32(19)) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+16)))
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13+v14)+12)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v21 = F_computeDistance(m, v16&int32(1), v19, v20)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = F_Float8GetDatum(m, v21)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				m.G0 = v7 + int32(16)
				return v25
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
			F_errmsg_internal(m, int32(480638), v7)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(498030), int32(1470), int32(415970))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
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
