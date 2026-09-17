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
	F_XLogRegisterBufData(m, int32(0), v49, v50&int32(_a_F_gistXLogUpdate_0))
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v164 int32
	_ = v164
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
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
	if v164 <= int32(0) {
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
	v45 = v15
	v46 = v31
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
		v164 = v121
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
	v164 = v56 - v57
	goto L4
L38:
	;
	goto L39
L39:
	;
	if v46 < int32(2) {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v126 = int32(1)
	v130 = int32(_a_F_gist_between_0)
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if v126 < v45 {
		v38 = v38 + (v133+int32(7))&v130
		v40 = v40 + (v56+int32(9))&v130
		v45 = v45 - v126
		v46 = v46 - v126
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
	v164 = v158 - v15
	goto L4
L45:
	;
	if v29 != 0 {
		v182 = v21
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v343 = int32(0)
	goto L47
L47:
	;
	return v343
L48:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v182)+4)))
	if v183 == int32(0) {
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
	v182 = v27
	goto L48
L51:
	;
	goto L52
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v182 = v27 + int32(base.Ui32(v178)>>(uint(int32(2))%32))
	goto L48
L53:
	;
	v343 = base.B2i32(int32(0) <= v316)
	goto L47
L54:
	;
	if base.Ui32(v183) < base.Ui32(v15) {
		goto L85
	} else {
		goto L86
	}
L55:
	;
	v190 = l1 + int32(16)
	v192 = v182 + int32(8)
	v197 = v15
	v198 = v183
	goto L56
L56:
	;
	v205 = v192 + int32(2)
	v207 = v190 + int32(23)
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v192))))
	v209 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+20)))
	if base.Ui32(v208) < base.Ui32(v209) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L54
L58:
	;
	v211 = v208
	goto L60
L59:
	;
	v211 = v209
	goto L60
L60:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v211) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if v273 != 0 {
		v316 = v273
		goto L53
	} else {
		goto L79
	}
L62:
	;
	v273 = int32(0)
	goto L61
L63:
	;
	v247 = v242
	v248 = v243
	v249 = v244
	goto L73
L64:
	;
	if (v205|v207)&int32(3) != 0 {
		v242 = v205
		v243 = v207
		v244 = v211
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v235 = v205
	v236 = v207
	v237 = v211
	goto L66
L66:
	;
	if v237 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v219 = v205
	v220 = v207
	v221 = v211
	goto L68
L68:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v224 != v225 {
		v242 = v219
		v243 = v220
		v244 = v221
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v235 = v230
	v236 = v228
	v237 = v232
	goto L66
L70:
	;
	v227 = int32(4)
	v228 = v220 + v227
	v230 = v219 + v227
	v232 = v221 - v227
	if base.Ui32(int32(3)) < base.Ui32(v232) {
		v219 = v230
		v220 = v228
		v221 = v232
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v242 = v235
	v243 = v236
	v244 = v237
	goto L63
L73:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248))))
	if v252 == v253 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v273 = v252 - v253
	goto L61
L75:
	;
	v255 = int32(1)
	v260 = v249 - v255
	if v260 != 0 {
		v247 = v247 + v255
		v248 = v248 + v255
		v249 = v260
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
	if v208 != v209 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v316 = v208 - v209
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
	v278 = int32(1)
	v282 = int32(_a_F_gist_between_0)
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190))))
	if v278 < v197 {
		v190 = v190 + (v285+int32(7))&v282
		v192 = v192 + (v208+int32(9))&v282
		v197 = v197 - v278
		v198 = v198 - v278
		goto L56
	} else {
		goto L84
	}
L84:
	;
	goto L57
L85:
	;
	v310 = v183
	goto L87
L86:
	;
	v310 = v15
	goto L87
L87:
	;
	v316 = v310 - v15
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
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
	if base.B2i32(v17 == v2)|base.B2i32(v12 == v2) != 0 {
		v48 = v2
		m.G0 = v9 + int32(16)
		return v48
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
		v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23)+16)))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v24)+12)))
		if v26&int32(1) != 0 {
			v32 = (v11 - int32(1)) & int32(_a_F_gist_box_consistent_0)
			if base.Ui32(int32(12)) <= base.Ui32(v32) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v11
					F_errmsg_internal(m, int32(_a_F_gist_box_consistent_1), v9)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_gist_box_consistent_2), int32(939), int32(_a_F_gist_box_consistent_3))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32))+uint32(_c_F_gist_box_consistent[0])))
				v39 = F_DirectFunctionCall2Coll(m, v37, int32(0), v17, v12)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v48 = base.B2i32(v39 != int32(0))
					m.G0 = v9 + int32(16)
					return v48
				}
			}
		} else {
			v45 = F_rtree_internal_consistent(m, v17, v12, v11)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				v48 = v45
				m.G0 = v9 + int32(16)
				return v48
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
	var v12 int32
	_ = v12
	var v18 float64
	_ = v18
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 float64
	_ = v22
	var v34 float64
	_ = v34
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 float64
	_ = v38
	var v50 float64
	_ = v50
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 float64
	_ = v54
	var v66 float64
	_ = v66
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 float64
	_ = v70
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.B2i32(v9 == v2)|base.B2i32(v12 == v2) == v2 {
		v18 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
		v20 = int64(9223372036854775807)
		v21 = base.I64_reinterpret_f64(v18) & v20
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v22)&v20) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v21) {
				v34 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
				v36 = int64(9223372036854775807)
				v37 = base.I64_reinterpret_f64(v34) & v36
				v38 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v38)&v36) {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v37) {
						v50 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
						v52 = int64(9223372036854775807)
						v53 = base.I64_reinterpret_f64(v50) & v52
						v54 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v54)&v52) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v53) {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
								v68 = int64(9223372036854775807)
								v69 = base.I64_reinterpret_f64(v66) & v68
								v70 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&v68) {
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v69)))
									return v8
								} else {
									v83 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v66, v70)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v83)
									return v8
								}
							} else {
								v98 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v98)
								return v8
							}
						} else {
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v53))|base.F64_ne(v50, v54) != 0 {
								v89 = v2
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v89)
								return v8
							} else {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
								v68 = int64(9223372036854775807)
								v69 = base.I64_reinterpret_f64(v66) & v68
								v70 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&v68) {
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v69)))
									return v8
								} else {
									v83 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v66, v70)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v83)
									return v8
								}
							}
						}
					} else {
						v98 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v98)
						return v8
					}
				} else {
					if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v37))|base.F64_ne(v34, v38) != 0 {
						v89 = v2
						*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v89)
						return v8
					} else {
						v50 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
						v52 = int64(9223372036854775807)
						v53 = base.I64_reinterpret_f64(v50) & v52
						v54 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v54)&v52) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v53) {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
								v68 = int64(9223372036854775807)
								v69 = base.I64_reinterpret_f64(v66) & v68
								v70 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&v68) {
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v69)))
									return v8
								} else {
									v83 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v66, v70)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v83)
									return v8
								}
							} else {
								v98 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v98)
								return v8
							}
						} else {
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v53))|base.F64_ne(v50, v54) != 0 {
								v89 = v2
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v89)
								return v8
							} else {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
								v68 = int64(9223372036854775807)
								v69 = base.I64_reinterpret_f64(v66) & v68
								v70 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&v68) {
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v69)))
									return v8
								} else {
									v83 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v66, v70)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v83)
									return v8
								}
							}
						}
					}
				}
			} else {
				v98 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v98)
				return v8
			}
		} else {
			if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v21))|base.F64_ne(v18, v22) != 0 {
				v89 = v2
				*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v89)
				return v8
			} else {
				v34 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
				v36 = int64(9223372036854775807)
				v37 = base.I64_reinterpret_f64(v34) & v36
				v38 = *(*float64)(unsafe.Add(mBase, uint32(v9)+24))
				if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v38)&v36) {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(v37) {
						v50 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
						v52 = int64(9223372036854775807)
						v53 = base.I64_reinterpret_f64(v50) & v52
						v54 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v54)&v52) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v53) {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
								v68 = int64(9223372036854775807)
								v69 = base.I64_reinterpret_f64(v66) & v68
								v70 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&v68) {
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v69)))
									return v8
								} else {
									v83 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v66, v70)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v83)
									return v8
								}
							} else {
								v98 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v98)
								return v8
							}
						} else {
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v53))|base.F64_ne(v50, v54) != 0 {
								v89 = v2
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v89)
								return v8
							} else {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
								v68 = int64(9223372036854775807)
								v69 = base.I64_reinterpret_f64(v66) & v68
								v70 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&v68) {
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v69)))
									return v8
								} else {
									v83 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v66, v70)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v83)
									return v8
								}
							}
						}
					} else {
						v98 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v98)
						return v8
					}
				} else {
					if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v37))|base.F64_ne(v34, v38) != 0 {
						v89 = v2
						*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v89)
						return v8
					} else {
						v50 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
						v52 = int64(9223372036854775807)
						v53 = base.I64_reinterpret_f64(v50) & v52
						v54 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
						if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v54)&v52) {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(v53) {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
								v68 = int64(9223372036854775807)
								v69 = base.I64_reinterpret_f64(v66) & v68
								v70 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&v68) {
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v69)))
									return v8
								} else {
									v83 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v66, v70)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v83)
									return v8
								}
							} else {
								v98 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v98)
								return v8
							}
						} else {
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v53))|base.F64_ne(v50, v54) != 0 {
								v89 = v2
								*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v89)
								return v8
							} else {
								v66 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
								v68 = int64(9223372036854775807)
								v69 = base.I64_reinterpret_f64(v66) & v68
								v70 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
								if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v70)&v68) {
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v69)))
									return v8
								} else {
									v83 = base.B2i32(base.Ui64(v69) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v66, v70)
									*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v83)
									return v8
								}
							}
						}
					}
				}
			}
		}
	} else {
		v89 = base.B2i32(v12|v9 == int32(0))
		*(*uint8)(unsafe.Add(mBase, uint32(v8))) = uint8(v89)
		return v8
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
			F_errmsg_internal(m, int32(_a_F_gist_circle_distance_0), v7)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gist_circle_distance_1), int32(1492), int32(_a_F_gist_circle_distance_2))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	if int32(0) <= base.I32_extend8_s(l0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(2))%32))&int32(60))+uint32(_c_F_gist_identify[0])))
		v11 = v9
	} else {
		v11 = int32(0)
	}
	return v11
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
			F_errmsg_internal(m, int32(_a_F_gist_point_distance_0), v7)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gist_point_distance_1), int32(1470), int32(_a_F_gist_point_distance_2))
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
