package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RangeVarAdjustRelationPersistence(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
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
	var v227 int32
	_ = v227
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
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
	switch v4 - int32(112) {
	case 0:
		goto L5
	default:
		goto L4
	case 4:
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L12
	} else {
		goto L104
	}
L2:
	;
	F_errmsg(m, int32(153029), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L12
	} else {
		goto L102
	}
L3:
	;
	return
L4:
	;
	v160 = F_get_namespace_name(m, l1)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L12
	} else {
		goto L66
	}
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v32 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L6:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	if v8 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v8 == l1 {
		goto L3
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v13 = F_isAnyTempNamespace(m, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v11 == l1 {
		goto L3
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	return
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v13 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(532209), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(524545), int32(860), int32(436460))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v41 = F_get_namespace_name(m, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L25
	}
L20:
	;
	if l1 != v32 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v37 != l1 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v39 = int32(116)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)) = uint8(v39)
	return
L24:
	;
	goto L23
L25:
	;
	if v41 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v45 = int32(533539)
	goto L30
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L62
	}
L28:
	;
	if v82-v83 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L30:
	;
	goto L31
L31:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v52 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v53 = v41
	v54 = v45
	v55 = int32(8)
	v56 = v52
	goto L36
L33:
	;
	v78 = v45
	v82 = int32(0)
	goto L34
L34:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	goto L28
L35:
	;
	v78 = v73
	v82 = v75
	goto L34
L36:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v56 != v58 {
		v73 = v54
		v75 = v56
		goto L35
	} else {
		goto L38
	}
L37:
	;
	v73 = v67
	v75 = int32(0)
	goto L35
L38:
	;
	if v58 == int32(0) {
		v73 = v54
		v75 = v56
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v63 = v55 - int32(1)
	if v63 == int32(0) {
		v73 = v54
		v75 = v56
		goto L35
	} else {
		goto L40
	}
L40:
	;
	v66 = int32(1)
	v67 = v54 + v66
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v68 != 0 {
		v53 = v53 + v66
		v54 = v67
		v55 = v63
		v56 = v68
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	F_pfree(m, v41)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v95 = int32(533524)
	goto L48
L45:
	;
	goto L27
L46:
	;
	F_pfree(m, v41)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L12
	} else {
		goto L60
	}
L48:
	;
	goto L49
L49:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v102 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v103 = v41
	v104 = v95
	v105 = int32(14)
	v106 = v102
	goto L54
L51:
	;
	v128 = v95
	v132 = int32(0)
	goto L52
L52:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	goto L46
L53:
	;
	v128 = v123
	v132 = v125
	goto L52
L54:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v106 != v108 {
		v123 = v104
		v125 = v106
		goto L53
	} else {
		goto L56
	}
L55:
	;
	v123 = v117
	v125 = int32(0)
	goto L53
L56:
	;
	if v108 == int32(0) {
		v123 = v104
		v125 = v106
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v113 = v105 - int32(1)
	if v113 == int32(0) {
		v123 = v104
		v125 = v106
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v116 = int32(1)
	v117 = v104 + v116
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v118 != 0 {
		v103 = v103 + v116
		v104 = v117
		v105 = v113
		v106 = v118
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	if v132-v133 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	goto L27
L62:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(153029), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(524545), int32(869), int32(436460))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	if v160 == int32(0) {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v164 = int32(533539)
	goto L70
L68:
	;
	if v201-v202 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L70:
	;
	goto L71
L71:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v171 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v172 = v160
	v173 = v164
	v174 = int32(8)
	v175 = v171
	goto L76
L73:
	;
	v197 = v164
	v201 = int32(0)
	goto L74
L74:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	goto L68
L75:
	;
	v197 = v192
	v201 = v194
	goto L74
L76:
	;
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173))))
	if v175 != v177 {
		v192 = v173
		v194 = v175
		goto L75
	} else {
		goto L78
	}
L77:
	;
	v192 = v186
	v194 = int32(0)
	goto L75
L78:
	;
	if v177 == int32(0) {
		v192 = v173
		v194 = v175
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v182 = v174 - int32(1)
	if v182 == int32(0) {
		v192 = v173
		v194 = v175
		goto L75
	} else {
		goto L80
	}
L80:
	;
	v185 = int32(1)
	v186 = v173 + v185
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+1)))
	if v187 != 0 {
		v172 = v172 + v185
		v173 = v186
		v174 = v182
		v175 = v187
		goto L76
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	F_pfree(m, v160)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L12
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v214 = int32(533524)
	goto L88
L85:
	;
	goto L1
L86:
	;
	F_pfree(m, v160)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L12
	} else {
		goto L100
	}
L88:
	;
	goto L89
L89:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	if v221 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v222 = v160
	v223 = v214
	v224 = int32(14)
	v225 = v221
	goto L94
L91:
	;
	v247 = v214
	v251 = int32(0)
	goto L92
L92:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247))))
	goto L86
L93:
	;
	v247 = v242
	v251 = v244
	goto L92
L94:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	if v225 != v227 {
		v242 = v223
		v244 = v225
		goto L93
	} else {
		goto L96
	}
L95:
	;
	v242 = v236
	v244 = int32(0)
	goto L93
L96:
	;
	if v227 == int32(0) {
		v242 = v223
		v244 = v225
		goto L93
	} else {
		goto L97
	}
L97:
	;
	v232 = v224 - int32(1)
	if v232 == int32(0) {
		v242 = v223
		v244 = v225
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v235 = int32(1)
	v236 = v223 + v235
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	if v237 != 0 {
		v222 = v222 + v235
		v223 = v236
		v224 = v232
		v225 = v237
		goto L94
	} else {
		goto L99
	}
L99:
	;
	goto L95
L100:
	;
	if v251-v252 == int32(0) {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	goto L3
L102:
	;
	F_errfinish(m, int32(524545), int32(856), int32(436460))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L12
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L12
	} else {
		goto L105
	}
L105:
	;
	F_errmsg(m, int32(184848), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L12
	} else {
		goto L106
	}
L106:
	;
	F_errfinish(m, int32(524545), int32(875), int32(436460))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L12
	} else {
		goto L107
	}
L107:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RangeVarCallbackForRenameAttribute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v6 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		if v6 != 0 {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+22)))
			F_renameatt_check(m, l1, v8+v9, int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F_ReleaseCatCache(m, v6)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			return
		}
	}
}
func F_RangeVarCallbackForRenameTrigger(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v16 = v14 + v15
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+119)))
			v19 = v17 - int32(102)
			v28 = (v19<<(uint(int32(7))%32) | int32(base.Ui32(v19&int32(254))>>(uint(int32(1))%32))) & int32(255)
			if base.Ui32(int32(8)) < base.Ui32(v28) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v90
						F_errmsg(m, int32(143173), v9)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return
						} else {
							v95 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+119)))
							F_errdetail_relkind_not_supported(m, v95)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								F_errfinish(m, int32(519958), int32(1440), int32(235974))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
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
			} else {
				if int32(1)<<(uint(v28)%32)&int32(353) == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v90
							F_errmsg(m, int32(143173), v9)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								v95 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+119)))
								F_errdetail_relkind_not_supported(m, v95)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									F_errfinish(m, int32(519958), int32(1440), int32(235974))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
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
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, _consts[4]))
					v40 = F_object_ownercheck(m, int32(1259), l1, v39)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						if v40 == int32(0) {
							v45 = F_get_rel_relkind(m, l1)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								switch v45 - int32(73) {
								case 0, 32:
									v58 = int32(20)
								default:
									v56 = int32(41)
									v58 = v56
								case 10:
									v58 = int32(37)
								case 29:
									v56 = int32(18)
									v58 = v56
								case 36:
									v58 = int32(23)
								case 45:
									v58 = int32(51)
								}
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_aclcheck_error(m, int32(2), v58, v59)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[270])))
									if v63 == int32(0) {
										v67 = int32(1)
										if base.Ui32(l1) < base.Ui32(int32(12000)) {
											v75 = v67
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
											if v70 == int32(99) {
												v75 = v67
											} else {
												v73 = F_isTempToastNamespace(m, v70)
												mBase = m.M
												v75 = v73
											}
										}
										if v75 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v106 = m.ExcPending
											if v106 != 0 {
												return
											} else {
												F_errcode(m, int32(16797828))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v110
													F_errmsg(m, int32(344134), v9+int32(16))
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return
													} else {
														F_errfinish(m, int32(519958), int32(1449), int32(235974))
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v12)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												m.G0 = v9 + int32(32)
												return
											}
										}
									} else {
										F_ReleaseCatCache(m, v12)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											m.G0 = v9 + int32(32)
											return
										}
									}
								}
							}
						} else {
							v63 = int32(*(*uint8)(unsafe.Add(mBase, _consts[270])))
							if v63 == int32(0) {
								v67 = int32(1)
								if base.Ui32(l1) < base.Ui32(int32(12000)) {
									v75 = v67
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
									if v70 == int32(99) {
										v75 = v67
									} else {
										v73 = F_isTempToastNamespace(m, v70)
										mBase = m.M
										v75 = v73
									}
								}
								if v75 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										F_errcode(m, int32(16797828))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v110
											F_errmsg(m, int32(344134), v9+int32(16))
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return
											} else {
												F_errfinish(m, int32(519958), int32(1449), int32(235974))
												mBase = m.M
												v121 = m.ExcPending
												if v121 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									F_ReleaseCatCache(m, v12)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							} else {
								F_ReleaseCatCache(m, v12)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(32)
			return
		}
	}
}
func F_RangeVarCallbackForTruncate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		v11 = F_SearchSysCache1(m, int32(57), l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			if v11 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1
					F_errmsg_internal(m, int32(50028), v8)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return
					} else {
						F_errfinish(m, int32(518745), int32(19541), int32(375562))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
				F_truncate_check_rel(m, l1, v15+v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+22)))
					v23 = *(*int32)(unsafe.Add(mBase, _consts[4]))
					v25 = F_pg_class_aclcheck(m, l1, v23, int64(16))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						if v25 != 0 {
							v27 = v20 + v21
							v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+119)))
							switch v28 - int32(73) {
							case 0, 32:
								v40 = int32(20)
							default:
								v38 = int32(41)
								v40 = v38
							case 10:
								v40 = int32(37)
							case 29:
								v38 = int32(18)
								v40 = v38
							case 36:
								v40 = int32(23)
							case 45:
								v40 = int32(51)
							}
							F_aclcheck_error(m, v25, v40, v27+int32(4))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v11)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
									return
								}
							}
						} else {
							F_ReleaseCatCache(m, v11)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func F_get_range_io_data(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 == l1 {
			v81 = v12
			m.G0 = v9 + int32(48)
			return v81
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
			v18 = F_MemoryContextAlloc(m, v16, int32(36))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v23 = F_lookup_type_cache(m, l1, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
							F_errmsg_internal(m, int32(388808), v9)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518394), int32(335), int32(529877))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
						F_get_type_io_data(m, v29, l2, v9+int32(46), v9+int32(45), v9+int32(44), v9+int32(43), v18+int32(32), v9+int32(36))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
							if v44 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(52461700))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+200))
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
										v57 = F_format_type_be(m, v56)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return int32(0)
										} else {
											if l2 == int32(2) {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v57
												F_errmsg(m, int32(200524), v9+int32(16))
												mBase = m.M
												v105 = m.ExcPending
												if v105 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(518394), int32(354), int32(529877))
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v57
												F_errmsg(m, int32(200435), v9+int32(32))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(518394), int32(359), int32(529877))
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
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
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
								F_fmgr_info_cxt(m, v44, v18+int32(4), v75)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v18
									v81 = v18
									m.G0 = v9 + int32(48)
									return v81
								}
							}
						}
					}
				}
			}
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
		v18 = F_MemoryContextAlloc(m, v16, int32(36))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v23 = F_lookup_type_cache(m, l1, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v23
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
						F_errmsg_internal(m, int32(388808), v9)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518394), int32(335), int32(529877))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					F_get_type_io_data(m, v29, l2, v9+int32(46), v9+int32(45), v9+int32(44), v9+int32(43), v18+int32(32), v9+int32(36))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
						if v44 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(52461700))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+200))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
									v57 = F_format_type_be(m, v56)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										if l2 == int32(2) {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v57
											F_errmsg(m, int32(200524), v9+int32(16))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(518394), int32(354), int32(529877))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v57
											F_errmsg(m, int32(200435), v9+int32(32))
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(518394), int32(359), int32(529877))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
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
						} else {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
							F_fmgr_info_cxt(m, v44, v18+int32(4), v75)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								*(*int32)(unsafe.Add(mBase, uint32(v78)+16)) = v18
								v81 = v18
								m.G0 = v9 + int32(48)
								return v81
							}
						}
					}
				}
			}
		}
	}
}
func F_get_range_nulltest(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = v8
	goto L3
L2:
	;
	v9 = v2
	goto L3
L3:
	;
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if int32(0) < v10 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L14
	} else {
		goto L24
	}
L5:
	;
	v15 = v9
	v16 = v2
	v17 = v2
	goto L8
L6:
	;
	v81 = v2
	goto L7
L7:
	;
	return v81
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19+v17<<(uint(int32(1))%32)))))
	if v23 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v81 = v72
	goto L7
L10:
	;
	v61 = F_palloc0(m, int32(20))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L21
	}
L11:
	;
	v26 = v17 << (uint(int32(2)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26+v27)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30+v26)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33+v26)))
	v37 = F_makeVar(m, int32(1), v23, v29, v32, v35, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if v15 == int32(0) {
		goto L4
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	v58 = v15
	v59 = v37
	goto L10
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v44 = F_copyObjectImpl(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v47 = v15 + int32(4)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if base.Ui32(v47) < base.Ui32(v50+v51<<(uint(int32(2))%32)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v56 = v47
	goto L20
L19:
	;
	v56 = int32(0)
	goto L20
L20:
	;
	v58 = v56
	v59 = v44
	goto L10
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = int32(-1)
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v61)+12)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = int32(52)
	v72 = F_lappend(m, v16, v61)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v75 = v17 + int32(1)
	v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v75 < v76 {
		v15 = v58
		v16 = v72
		v17 = v75
		goto L8
	} else {
		goto L23
	}
L23:
	;
	goto L9
L24:
	;
	F_errmsg_internal(m, int32(153405), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L14
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(518482), int32(4700), int32(83313))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L14
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_range_partbound_string(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = F_makeStringInfo(m)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(0)
	v16 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v8)+28)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = v16
	*(*int64)(unsafe.Add(mBase, uint32(v8)+12)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v10
	F_appendStringInfoChar(m, v10, int32(40))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_appendStringInfoChar(m, v10, int32(41))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L27
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	F_appendStringInfoString(m, v10, int32(791891))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	switch v38 + int32(1) {
	case 0:
		goto L10
	default:
		goto L9
	case 2:
		goto L11
	}
L8:
	;
	v53 = int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v54 <= v53 {
		goto L4
	} else {
		goto L15
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	F_get_const_expr(m, v47, v8+int32(8), int32(-1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	F_appendStringInfoString(m, v10, int32(563986))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_appendStringInfoString(m, v10, int32(563933))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	goto L8
L14:
	;
	goto L8
L15:
	;
	v60 = v53
	goto L16
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v60<<(uint(int32(2))%32))))
	F_appendStringInfoString(m, v10, int32(780599))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L4
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	switch v70 + int32(1) {
	case 0:
		goto L22
	default:
		goto L20
	case 2:
		goto L21
	}
L19:
	;
	v86 = v60 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v86 < v87 {
		v60 = v86
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	F_get_const_expr(m, v79, v8+int32(8), int32(-1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L25
	}
L21:
	;
	F_appendStringInfoString(m, v10, int32(563933))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	F_appendStringInfoString(m, v10, int32(563986))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	goto L19
L25:
	;
	goto L19
L26:
	;
	goto L17
L27:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	m.G0 = v8 + int32(48)
	return v97
}
func F_makeRangeVar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v6 = F_palloc0(m, int32(28))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(0)
		v13 = int32(28673)
		*(*uint16)(unsafe.Add(mBase, uint32(v6)+16)) = uint16(v13)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(3)
		return v6
	}
}
func F_range_adjacent(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_adjacent_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(388808), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(518394), int32(1776), int32(418064))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_adjacent_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(388808), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518394), int32(1776), int32(418064))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_adjacent_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_after(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_after_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(388808), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(518394), int32(1776), int32(418064))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_after_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(388808), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518394), int32(1776), int32(418064))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_after_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_before_multirange(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
					if v41&int32(1) != 0 {
						v70 = v34
						m.G0 = v9 + int32(48)
						return v70
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v44 == int32(0) {
							v70 = v34
							m.G0 = v9 + int32(48)
							return v70
						} else {
							F_range_deserialize(m, v33, v12, v9+int32(40), v9+int32(32), v9+int32(15))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_multirange_get_bounds(m, v33, v17, int32(0), v9+int32(24), v9+int32(16))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v66 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(24))
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v70 = int32(base.Ui32(v66) >> (uint(int32(31)) % 32))
										m.G0 = v9 + int32(48)
										return v70
									}
								}
							}
						}
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(65536))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(388681), v9)
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(518389), int32(558), int32(418059))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = int32(0)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
							if v41&int32(1) != 0 {
								v70 = v34
								m.G0 = v9 + int32(48)
								return v70
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								if v44 == int32(0) {
									v70 = v34
									m.G0 = v9 + int32(48)
									return v70
								} else {
									F_range_deserialize(m, v33, v12, v9+int32(40), v9+int32(32), v9+int32(15))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										F_multirange_get_bounds(m, v33, v17, int32(0), v9+int32(24), v9+int32(16))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v66 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(24))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v70 = int32(base.Ui32(v66) >> (uint(int32(31)) % 32))
												m.G0 = v9 + int32(48)
												return v70
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(65536))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(388681), v9)
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518389), int32(558), int32(418059))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
						if v41&int32(1) != 0 {
							v70 = v34
							m.G0 = v9 + int32(48)
							return v70
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if v44 == int32(0) {
								v70 = v34
								m.G0 = v9 + int32(48)
								return v70
							} else {
								F_range_deserialize(m, v33, v12, v9+int32(40), v9+int32(32), v9+int32(15))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_multirange_get_bounds(m, v33, v17, int32(0), v9+int32(24), v9+int32(16))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v66 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(24))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v70 = int32(base.Ui32(v66) >> (uint(int32(31)) % 32))
											m.G0 = v9 + int32(48)
											return v70
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
func F_range_before_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
	if v16&int32(1) != 0 {
		v47 = v4
		m.G0 = v8 + int32(48)
		return v47
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v19 == int32(0) {
			v47 = v4
			m.G0 = v8 + int32(48)
			return v47
		} else {
			F_range_deserialize(m, l0, l1, v8+int32(40), v8+int32(32), v8+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_multirange_get_bounds(m, l0, l2, int32(0), v8+int32(24), v8+int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v43 = F_range_cmp_bounds(m, l0, v8+int32(32), v8+int32(24))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v47 = int32(base.Ui32(v43) >> (uint(int32(31)) % 32))
						m.G0 = v8 + int32(48)
						return v47
					}
				}
			}
		}
	}
}
func F_range_cmp_bounds(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v6 == int32(1) {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
		if v5&int32(1) != 0 {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
			if v12 == v9&int32(255) {
				return int32(0)
			} else {
				v19 = int32(1)
				if v9&v19 != 0 {
					v22 = int32(-1)
				} else {
					v22 = v19
				}
				return v22
			}
		} else {
			v25 = int32(1)
			if v9&v25 != 0 {
				v28 = int32(-1)
			} else {
				v28 = v25
			}
			return v28
		}
	} else {
		if v5&int32(1) != 0 {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
			if v34 != 0 {
				v35 = int32(1)
			} else {
				v35 = int32(-1)
			}
			return v35
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v42 = F_FunctionCall2Coll(m, l0+int32(212), v39, v40, v41)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				if v42 != 0 {
					v82 = v42
					return v82
				} else {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
					if v47 == int32(0) {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
						if v46&int32(1) == int32(0) {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
							if v55 == v50&int32(255) {
								return int32(0)
							} else {
								v61 = int32(1)
								if v50&v61 != 0 {
									v65 = v61
								} else {
									v65 = int32(-1)
								}
								return v65
							}
						} else {
							v67 = int32(1)
							if v50&v67 != 0 {
								v71 = v67
							} else {
								v71 = int32(-1)
							}
							return v71
						}
					} else {
						if v46&int32(1) != 0 {
							return int32(0)
						} else {
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
							if v79 != 0 {
								v80 = int32(-1)
							} else {
								v80 = int32(1)
							}
							v82 = v80
							return v82
						}
					}
				}
			}
		}
	}
}
func F_range_constructor2(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_get_fn_expr_rettype(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v14 {
				v30 = v19
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				v32 = int32(257)
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+29)) = uint16(v32)
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v31)
				if v31 != 0 {
					v36 = int32(0)
				} else {
					v36 = v12
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v36
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				v39 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v9)+21)) = uint16(v39)
				*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v38)
				if v38 != 0 {
					v43 = v39
				} else {
					v43 = v11
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v43
				v49 = int32(0)
				v51 = F_make_range(m, v30, v9+int32(24), v9+int32(16), v49, v49)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(32)
					return v51
				}
			} else {
				v23 = F_lookup_type_cache(m, v14, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
							F_errmsg_internal(m, int32(388808), v9)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518394), int32(1776), int32(418064))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
						v32 = int32(257)
						*(*uint16)(unsafe.Add(mBase, uint32(v9)+29)) = uint16(v32)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v31)
						if v31 != 0 {
							v36 = int32(0)
						} else {
							v36 = v12
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v36
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						v39 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v9)+21)) = uint16(v39)
						*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v38)
						if v38 != 0 {
							v43 = v39
						} else {
							v43 = v11
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v43
						v49 = int32(0)
						v51 = F_make_range(m, v30, v9+int32(24), v9+int32(16), v49, v49)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(32)
							return v51
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v14, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v14
						F_errmsg_internal(m, int32(388808), v9)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518394), int32(1776), int32(418064))
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					v32 = int32(257)
					*(*uint16)(unsafe.Add(mBase, uint32(v9)+29)) = uint16(v32)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)) = uint8(v31)
					if v31 != 0 {
						v36 = int32(0)
					} else {
						v36 = v12
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v36
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					v39 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v9)+21)) = uint16(v39)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)) = uint8(v38)
					if v38 != 0 {
						v43 = v39
					} else {
						v43 = v11
					}
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v43
					v49 = int32(0)
					v51 = F_make_range(m, v30, v9+int32(24), v9+int32(16), v49, v49)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(32)
						return v51
					}
				}
			}
		}
	}
}
func F_range_contains_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
		if v16&int32(1) != 0 {
			v65 = int32(0)
			m.G0 = v7 + int32(48)
			return v65
		} else {
			F_range_deserialize(m, l0, l1, v7+int32(40), v7+int32(32), v7+int32(7))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				F_multirange_get_bounds(m, l0, l2, v30, v7+int32(24), v7+int32(8))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					F_multirange_get_bounds(m, l0, l2, v38-int32(1), v7+int32(8), v7+int32(16))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v51 = F_range_cmp_bounds(m, l0, v7+int32(40), v7+int32(24))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							if int32(0) < v51 {
								v65 = v30
								m.G0 = v7 + int32(48)
								return v65
							} else {
								v59 = F_range_cmp_bounds(m, l0, v7+int32(32), v7+int32(16))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									if v59 < int32(0) {
										v65 = v30
									} else {
										v65 = int32(1)
									}
									m.G0 = v7 + int32(48)
									return v65
								}
							}
						}
					}
				}
			}
		}
	} else {
		v65 = int32(1)
		m.G0 = v7 + int32(48)
		return v65
	}
}
func F_range_gist_consistent(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v23)
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
		v26 = F_range_get_typcache(m, l0, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28+v29)+12)))
			if v31&int32(1) != 0 {
				if v13 != 0 {
					if v13 == int32(4537) {
						v42 = F_pg_detoast_datum(m, v15)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = F_range_gist_consistent_leaf_multirange(m, v26, v14, v19, v42)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								v79 = v44
								m.G0 = v11 + int32(32)
								return v79
							}
						}
					} else {
						if v13 != int32(3831) {
							if v14 == int32(16) {
								v77 = F_range_contains_elem_internal(m, v26, v19, v15)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v79 = v77
									m.G0 = v11 + int32(32)
									return v79
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
									F_errmsg_internal(m, int32(503658), v11+int32(16))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(516205), int32(1138), int32(102668))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
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
							v38 = F_pg_detoast_datum(m, v15)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								v40 = F_range_gist_consistent_leaf_range(m, v26, v14, v19, v38)
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v79 = v40
									m.G0 = v11 + int32(32)
									return v79
								}
							}
						}
					}
				} else {
					v38 = F_pg_detoast_datum(m, v15)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = F_range_gist_consistent_leaf_range(m, v26, v14, v19, v38)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v79 = v40
							m.G0 = v11 + int32(32)
							return v79
						}
					}
				}
			} else {
				if v13 != 0 {
					if v13 == int32(4537) {
						v71 = F_pg_detoast_datum(m, v15)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							v73 = F_range_gist_consistent_int_multirange(m, v26, v14, v19, v71)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								v79 = v73
								m.G0 = v11 + int32(32)
								return v79
							}
						}
					} else {
						if v13 != int32(3831) {
							if v14 != int32(16) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
									F_errmsg_internal(m, int32(503658), v11)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(516205), int32(1049), int32(102598))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v77 = F_range_contains_elem_internal(m, v26, v19, v15)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									v79 = v77
									m.G0 = v11 + int32(32)
									return v79
								}
							}
						} else {
							v67 = F_pg_detoast_datum(m, v15)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v69 = F_range_gist_consistent_int_range(m, v26, v14, v19, v67)
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v79 = v69
									m.G0 = v11 + int32(32)
									return v79
								}
							}
						}
					}
				} else {
					v67 = F_pg_detoast_datum(m, v15)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v69 = F_range_gist_consistent_int_range(m, v26, v14, v19, v67)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v79 = v69
							m.G0 = v11 + int32(32)
							return v79
						}
					}
				}
			}
		}
	}
}
func F_range_gist_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v21 = int32(*(*int8)(unsafe.Add(mBase, uint32(v7+int32(base.Ui32(v15)>>(uint(int32(2))%32))-int32(1)))))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v22)>>(uint(int32(2))%32))-int32(1)))))
			if v21 == v28 {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v31 = F_range_get_typcache(m, l0, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = F_range_eq_internal(m, v31, v7, v12)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = v33
						*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v35)
						return v14
					}
				}
			} else {
				v35 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v35)
				return v14
			}
		}
	}
}
func F_range_lt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2) >> (uint(int32(31)) % 32))
	}
}
func F_range_merge(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v34 = F_range_union_internal(m, v32, v12, v17, int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v34
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(388808), v9)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(518394), int32(1776), int32(418064))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v34 = F_range_union_internal(m, v32, v12, v17, int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(388808), v9)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518394), int32(1776), int32(418064))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v34 = F_range_union_internal(m, v32, v12, v17, int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func F_range_overlaps(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_overlaps_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(388808), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(518394), int32(1776), int32(418064))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_overlaps_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(388808), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518394), int32(1776), int32(418064))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_overlaps_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_overleft_multirange(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
					v34 = int32(0)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
					if v41&int32(1) != 0 {
						v72 = v34
						m.G0 = v9 + int32(48)
						return v72
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v44 == int32(0) {
							v72 = v34
							m.G0 = v9 + int32(48)
							return v72
						} else {
							F_range_deserialize(m, v33, v12, v9+int32(40), v9+int32(32), v9+int32(15))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								F_multirange_get_bounds(m, v33, v17, v55-int32(1), v9+int32(24), v9+int32(16))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									v68 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(16))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v72 = base.B2i32(v68 <= int32(0))
										m.G0 = v9 + int32(48)
										return v72
									}
								}
							}
						}
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(65536))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(388681), v9)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(518389), int32(558), int32(418059))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
							v34 = int32(0)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
							if v41&int32(1) != 0 {
								v72 = v34
								m.G0 = v9 + int32(48)
								return v72
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								if v44 == int32(0) {
									v72 = v34
									m.G0 = v9 + int32(48)
									return v72
								} else {
									F_range_deserialize(m, v33, v12, v9+int32(40), v9+int32(32), v9+int32(15))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
										F_multirange_get_bounds(m, v33, v17, v55-int32(1), v9+int32(24), v9+int32(16))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v68 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(16))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												v72 = base.B2i32(v68 <= int32(0))
												m.G0 = v9 + int32(48)
												return v72
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(65536))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+296))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(388681), v9)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518389), int32(558), int32(418059))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+296))
						v34 = int32(0)
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12+int32(base.Ui32(v35)>>(uint(int32(2))%32))-int32(1)))))
						if v41&int32(1) != 0 {
							v72 = v34
							m.G0 = v9 + int32(48)
							return v72
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							if v44 == int32(0) {
								v72 = v34
								m.G0 = v9 + int32(48)
								return v72
							} else {
								F_range_deserialize(m, v33, v12, v9+int32(40), v9+int32(32), v9+int32(15))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
									F_multirange_get_bounds(m, v33, v17, v55-int32(1), v9+int32(24), v9+int32(16))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										v68 = F_range_cmp_bounds(m, v33, v9+int32(32), v9+int32(16))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											v72 = base.B2i32(v68 <= int32(0))
											m.G0 = v9 + int32(48)
											return v72
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
func F_range_overright(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_overright_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(388808), v9)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(518394), int32(1776), int32(418064))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_overright_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(388808), v9)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(518394), int32(1776), int32(418064))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_overright_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33
						}
					}
				}
			}
		}
	}
}
func F_range_overright_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v21 int32
	_ = v21
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v9 == v10 {
		F_range_deserialize(m, l0, l1, v7+int32(40), v7+int32(24), v7+int32(15))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_range_deserialize(m, l0, l2, v7+int32(32), v7+int32(16), v7+int32(14))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = int32(0)
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
				if v31 != 0 {
					v104 = v30
					m.G0 = v7 + int32(48)
					return v104
				} else {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
					if v32 != 0 {
						v104 = v30
						m.G0 = v7 + int32(48)
						return v104
					} else {
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+36)))
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+44)))
						if v34 == int32(1) {
							v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
							if v33&int32(1) != 0 {
								v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
								if v40 == v37&int32(255) {
									v99 = int32(0)
								} else {
									v46 = int32(1)
									if v37&v46 != 0 {
										v49 = int32(-1)
									} else {
										v49 = v46
									}
									v99 = v49
								}
							} else {
								v51 = int32(1)
								if v37&v51 != 0 {
									v54 = int32(-1)
								} else {
									v54 = v51
								}
								v99 = v54
							}
							v104 = base.B2i32(int32(0) <= v99)
							m.G0 = v7 + int32(48)
							return v104
						} else {
							if v33&int32(1) != 0 {
								v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
								if v59 != 0 {
									v60 = int32(1)
								} else {
									v60 = int32(-1)
								}
								v99 = v60
								v104 = base.B2i32(int32(0) <= v99)
								m.G0 = v7 + int32(48)
								return v104
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
								v66 = F_FunctionCall2Coll(m, l0+int32(212), v63, v64, v65)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int32(0)
								} else {
									if v66 != 0 {
										v99 = v66
									} else {
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+37)))
										v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+45)))
										if v69 == int32(0) {
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
											if v68&int32(1) == int32(0) {
												v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
												if v77 == v72&int32(255) {
													v99 = int32(0)
												} else {
													v82 = int32(1)
													if v72&v82 != 0 {
														v86 = v82
													} else {
														v86 = int32(-1)
													}
													v99 = v86
												}
											} else {
												v87 = int32(1)
												if v72&v87 != 0 {
													v91 = v87
												} else {
													v91 = int32(-1)
												}
												v99 = v91
											}
										} else {
											if v68&int32(1) != 0 {
												v99 = int32(0)
											} else {
												v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
												if v97 != 0 {
													v98 = int32(-1)
												} else {
													v98 = int32(1)
												}
												v99 = v98
											}
										}
									}
									v104 = base.B2i32(int32(0) <= v99)
									m.G0 = v7 + int32(48)
									return v104
								}
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v112 = m.ExcPending
		if v112 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(341586), int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(518394), int32(941), int32(326159))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
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
func F_range_serialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	if l3 != 0 {
		v100 = int32(1)
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v332
L2:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v101)+12)))
	v103 = int32(*(*int8)(unsafe.Add(mBase, uint32(v101)+11)))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+10)))
	v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+8)))
	v108 = v100 & int32(9)
	if v108 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v15 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v81 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L5:
	;
	v70 = int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v71 != v70 {
		v100 = v70
		goto L2
	} else {
		goto L30
	}
L6:
	;
	if v41 != 0 {
		goto L4
	} else {
		goto L29
	}
L7:
	;
	v50 = F_errsave_start(m, l4)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L21
	} else {
		goto L24
	}
L8:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	if v14&int32(1) != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if v14&int32(1) != 0 {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v21 == v18&int32(255) {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v18&int32(1) == int32(0) {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	if v18&int32(1) == int32(0) {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L4
L16:
	;
	goto L4
L17:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)))
	if v35 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v41 = F_FunctionCall2Coll(m, l0+int32(212), v38, v39, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L4
L21:
	;
	return int32(0)
L22:
	;
	if v41 <= int32(0) {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	goto L7
L24:
	;
	if v50 == int32(0) {
		v332 = int32(0)
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(445184), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	F_errsave_finish(m, l4, int32(518394), int32(1821), int32(359801))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	return int32(0)
L29:
	;
	goto L5
L30:
	;
	v74 = int32(1)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v75 != v74 {
		v100 = v74
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L4
L32:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v86 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v88 = int32(8)
	goto L34
L34:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v91 == int32(1) {
		v100 = v88 | int32(16)
		goto L2
	} else {
		goto L38
	}
L35:
	;
	v87 = int32(2)
	goto L37
L36:
	;
	v87 = int32(0)
	goto L37
L37:
	;
	v88 = v87
	goto L34
L38:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v96 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v97 = v88 | int32(4)
	goto L41
L40:
	;
	v97 = v88
	goto L41
L41:
	;
	v100 = v97
	goto L2
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v105 == int32(-1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v206 = int32(8)
	goto L44
L44:
	;
	v208 = v100 & int32(17)
	if v208 != 0 {
		goto L78
	} else {
		goto L79
	}
L45:
	;
	v114 = F_pg_detoast_datum_packed(m, v111)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L21
	} else {
		goto L48
	}
L46:
	;
	v117 = v111
	goto L47
L47:
	;
	v118 = int32(8)
	v121 = base.B2i32(v105 != int32(-1))
	if v105 != int32(-1) {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v114
	v117 = v114
	goto L47
L49:
	;
	v206 = v205
	goto L44
L50:
	;
	v205 = v196 + v198
	goto L49
L51:
	;
	v193 = F_strlen(m, v117)
	mBase = m.M
	v196 = v154
	v198 = v193 + int32(1)
	goto L50
L52:
	;
	if v161&int32(255) == int32(1) {
		goto L68
	} else {
		goto L69
	}
L53:
	;
	switch v103 - int32(99) {
	case 0:
		v154 = v118
		goto L62
	case 1:
		goto L64
	default:
		goto L63
	case 6:
		goto L65
	}
L54:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v136&int32(1) != 0 {
		v160 = v118
		v161 = v136
		goto L52
	} else {
		goto L61
	}
L55:
	;
	if v105 != int32(-1) {
		goto L53
	} else {
		goto L60
	}
L56:
	;
	if v102 == int32(112) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v124&int32(3) != 0 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v131 = int32(base.Ui32(v127)>>(uint(int32(2))%32)) - int32(3)
	if base.Ui32(int32(127)) < base.Ui32(v131) {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v205 = v118 + v131
	goto L49
L60:
	;
	goto L54
L61:
	;
	goto L53
L62:
	;
	if int32(0) < v105 {
		v196 = v154
		v198 = v105
		goto L50
	} else {
		goto L66
	}
L63:
	;
	v154 = int32(8)
	goto L62
L64:
	;
	v154 = int32(8)
	goto L62
L65:
	;
	v154 = int32(8)
	goto L62
L66:
	;
	if v105 != int32(-1) {
		goto L51
	} else {
		goto L67
	}
L67:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	v160 = v154
	v161 = v159
	goto L52
L68:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if base.Ui32((v167-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v196 = v160
		v198 = int32(6)
		goto L50
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v161&int32(1) != 0 {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	v174 = int32(18)
	if v167&int32(255) == v174 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v180 = v174
	goto L74
L73:
	;
	v180 = int32(2)
	goto L74
L74:
	;
	v205 = v180 + v160
	goto L49
L75:
	;
	v205 = int32(base.Ui32(v161&int32(254))>>(uint(int32(1))%32)) + v160
	goto L49
L76:
	;
	goto L77
L77:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v205 = int32(base.Ui32(v189)>>(uint(int32(2))%32)) + v160
	goto L49
L78:
	;
	v304 = v206
	goto L80
L79:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v105 == int32(-1) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v306 = v304 + int32(1)
	v307 = F_palloc0(m, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L21
	} else {
		goto L114
	}
L81:
	;
	v212 = F_pg_detoast_datum_packed(m, v209)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L21
	} else {
		goto L84
	}
L82:
	;
	v215 = v209
	goto L83
L83:
	;
	v218 = base.B2i32(v105 != int32(-1))
	if v105 != int32(-1) {
		goto L91
	} else {
		goto L92
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v212
	v215 = v212
	goto L83
L85:
	;
	v304 = v302
	goto L80
L86:
	;
	v302 = v293 + v295
	goto L85
L87:
	;
	v290 = F_strlen(m, v215)
	mBase = m.M
	v293 = v251
	v295 = v290 + int32(1)
	goto L86
L88:
	;
	if v258&int32(255) == int32(1) {
		goto L104
	} else {
		goto L105
	}
L89:
	;
	switch v103 - int32(99) {
	case 0:
		v251 = v206
		goto L98
	case 1:
		goto L100
	default:
		goto L99
	case 6:
		goto L101
	}
L90:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v233&int32(1) != 0 {
		v257 = v206
		v258 = v233
		goto L88
	} else {
		goto L97
	}
L91:
	;
	if v105 != int32(-1) {
		goto L89
	} else {
		goto L96
	}
L92:
	;
	if v102 == int32(112) {
		goto L91
	} else {
		goto L93
	}
L93:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v221&int32(3) != 0 {
		goto L90
	} else {
		goto L94
	}
L94:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v228 = int32(base.Ui32(v224)>>(uint(int32(2))%32)) - int32(3)
	if base.Ui32(int32(127)) < base.Ui32(v228) {
		goto L90
	} else {
		goto L95
	}
L95:
	;
	v302 = v206 + v228
	goto L85
L96:
	;
	goto L90
L97:
	;
	goto L89
L98:
	;
	if int32(0) < v105 {
		v293 = v251
		v295 = v105
		goto L86
	} else {
		goto L102
	}
L99:
	;
	v251 = (v206 + int32(1)) & int32(-2)
	goto L98
L100:
	;
	v251 = (v206 + int32(7)) & int32(-8)
	goto L98
L101:
	;
	v251 = (v206 + int32(3)) & int32(-4)
	goto L98
L102:
	;
	if v105 != int32(-1) {
		goto L87
	} else {
		goto L103
	}
L103:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	v257 = v251
	v258 = v256
	goto L88
L104:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if base.Ui32((v264-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v293 = v257
		v295 = int32(6)
		goto L86
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v258&int32(1) != 0 {
		goto L111
	} else {
		goto L112
	}
L107:
	;
	v271 = int32(18)
	if v264&int32(255) == v271 {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v277 = v271
	goto L110
L109:
	;
	v277 = int32(2)
	goto L110
L110:
	;
	v302 = v277 + v257
	goto L85
L111:
	;
	v302 = int32(base.Ui32(v258&int32(254))>>(uint(int32(1))%32)) + v257
	goto L85
L112:
	;
	goto L113
L113:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v302 = int32(base.Ui32(v286)>>(uint(int32(2))%32)) + v257
	goto L85
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v307))) = v306 << (uint(int32(2)) % 32)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+4)) = v312
	v315 = v307 + int32(8)
	if v108 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v321 = F_datum_write(m, v315, v318, v104&int32(1), v103, v105, v102)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L21
	} else {
		goto L118
	}
L116:
	;
	v323 = v315
	goto L117
L117:
	;
	if v208 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v323 = v321
	goto L117
L119:
	;
	v329 = v323
	goto L121
L120:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v327 = F_datum_write(m, v323, v324, v104&int32(1), v103, v105, v102)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L21
	} else {
		goto L122
	}
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v329))) = uint8(v100)
	v332 = v307
	goto L1
L122:
	;
	v329 = v327
	goto L121
}
