package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTN2QT(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v2
	F_cntsize(m, l0, v9+int32(28), v9+int32(24))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v28 = base.I32_div_u_s(int32(1073741815)-v25, int32(12))
		if base.Ui32(v28) < base.Ui32(v23) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(381252), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(474355), int32(376), int32(494057))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v47 = v23 * int32(12)
			v50 = v47 + v25 + int32(8)
			v51 = F_palloc0(m, v50)
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v51))) = v50 << (uint(int32(2)) % 32)
				v58 = v51 + int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v58
				v60 = v58 + v47
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v60
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v60
				F_fillQT(m, v9+int32(12), l0)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(32)
					return v51
				}
			}
		}
	}
}
func F_QueueFKConstraintValidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l4)+16))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
	v21 = v19 + v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+119)))
	if v23 == int32(114) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	v249 = F_heap_copytuple(m, l4)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L40
	}
L2:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	F_ScanKeyInit(m, v17, int32(12), int32(3), int32(184), v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L27
	}
L3:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
	v175 = F_get_rel_relkind(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L25
	}
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
	if v26 != l3 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	v155 = v23
	goto L6
L6:
	;
	if v155&int32(255) == int32(112) {
		goto L2
	} else {
		goto L24
	}
L7:
	;
	v29 = F_palloc0(m, int32(108))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(161)
	v35 = F_pstrdup(m, v21+int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+8)) = v35
	v39 = F_palloc0(m, int32(32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v41
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v47
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v53 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+64))
	v136 = F_lappend(m, v135, v39)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L23
	}
L13:
	;
	v99 = F_palloc0(m, int32(140))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L20
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v56 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v68 = int32(0)
	goto L16
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v59+v68<<(uint(int32(2))%32))))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 == v52 {
		v127 = v78
		goto L12
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	v82 = v68 + int32(1)
	if v56 != v82 {
		v68 = v82
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99))) = v52
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+4)) = uint8(v105)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v108 = F_CreateTupleDescCopyConstr(m, v107)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+8)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v99)+88)) = int64(0)
	v113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+84)) = uint8(v113)
	v115 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+96)) = uint16(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v118 = F_lappend(m, v117, v99)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v118
	v127 = v99
	goto L12
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127)+64)) = v136
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+119)))
	v155 = v140
	goto L6
L24:
	;
	goto L3
L25:
	;
	if v175 != int32(112) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L2
L27:
	;
	v200 = int32(1)
	v203 = F_systable_beginscan(m, l1, int32(2579), v200, int32(0), v200, v17)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	goto L29
L29:
	;
	v219 = F_systable_getnext(m, v203)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L31
	}
L30:
	;
	F_systable_endscan(m, v203)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L8
	} else {
		goto L39
	}
L31:
	;
	if v219 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+22)))
	v223 = v221 + v222
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+76)))
	if v224 != 0 {
		goto L29
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L30
L35:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223)+80))
	v226 = F_table_open(m, v225, l5)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	F_QueueFKConstraintValidation(m, l0, l1, v226, l3, v219, l5)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	F_sequence_close(m, v226, int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	goto L29
L39:
	;
	goto L1
L40:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v249)+16))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+22)))
	v254 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v251+v252)+76)) = uint8(v254)
	F_CatalogTupleUpdate(m, l1, v249+int32(4), v249)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _consts[206]))
	if v261 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v264 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v263, v264, v264, v264)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_pfree(m, v249)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	m.G0 = v17 + int32(48)
	return
}
func F_quote_identifier(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v16 = base.B2i32(v7 == int32(95)) | base.B2i32(base.Ui32((v7-int32(97))&int32(255)) < base.Ui32(int32(26)))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v7
	v19 = l0
	v21 = int32(0)
	v22 = v16
	goto L4
L2:
	;
	v57 = v16
	v58 = int32(3)
	goto L3
L3:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, _consts[845])))
	if v60 != 0 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if base.Ui32((v18-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		v44 = v21
		v45 = v22
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v57 = v45
	v58 = v44 + int32(3)
	goto L3
L6:
	;
	v47 = v19 + int32(1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
	if v48 != 0 {
		v18 = v48
		v19 = v47
		v21 = v44
		v22 = v45
		goto L4
	} else {
		goto L10
	}
L7:
	;
	v30 = v18 & int32(255)
	if v30 == int32(95) {
		v44 = v21
		v45 = v22
		goto L6
	} else {
		goto L8
	}
L8:
	;
	if base.Ui32((v18-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		v44 = v21
		v45 = v22
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v44 = v21 + base.B2i32(v30 == int32(34))
	v45 = int32(0)
	goto L6
L10:
	;
	goto L5
L11:
	;
	if l0&int32(3) == int32(0) {
		v99 = l0
		goto L22
	} else {
		goto L23
	}
L12:
	;
	if v57 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v64 = F_ScanKeywordLookup(m, l0, int32(1819932))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v64 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return l0
L17:
	;
	goto L18
L18:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+uint32(_consts[846]))))
	if v73 != 0 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	return l0
L20:
	;
	v134 = F_palloc(m, v132+v58)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L14
	} else {
		goto L37
	}
L21:
	;
	v132 = v124 - l0
	goto L20
L22:
	;
	v103 = v99
	goto L31
L23:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v83 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v132 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v88 = l0
	goto L27
L27:
	;
	v92 = v88 + int32(1)
	if v92&int32(3) == int32(0) {
		v99 = v92
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v124 = v92
	goto L21
L29:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	if v97 != 0 {
		v88 = v92
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v112 = int32(-2139062144)
	if (int32(16843008)-v109|v109)&v112 == v112 {
		v103 = v103 + int32(4)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v118 = v103
	goto L34
L33:
	;
	goto L32
L34:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	if v122 != 0 {
		v118 = v118 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v124 = v118
	goto L21
L36:
	;
	goto L35
L37:
	;
	v136 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v136)
	v138 = l0
	v139 = v134
	goto L38
L38:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v144 != int32(34) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v159 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v139)+1)) = uint16(v159)
	return v134
L40:
	;
	goto L39
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v155))) = uint8(v144)
	v138 = v138 + int32(1)
	v139 = v155
	goto L38
L42:
	;
	if v144 == int32(0) {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v151 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)) = uint8(v151)
	v155 = v139 + int32(2)
	goto L41
L45:
	;
	v155 = v139 + int32(1)
	goto L41
}
func F_quote_nullable(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v2 == int32(1) {
		v6 = F_cstring_to_text(m, int32(509370))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v6
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = F_DirectFunctionCall1Coll(m, int32(1482), int32(0), v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
