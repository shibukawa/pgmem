package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_Async_Notify(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	goto L1
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	if v17 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L10
	} else {
		goto L96
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L10
	} else {
		goto L92
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L10
	} else {
		goto L88
	}
L5:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[350])))
	if v21 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L10
	} else {
		goto L85
	}
L8:
	;
	if l0 != 0 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v26 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v26 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	F_errmsg_internal(m, int32(674523), v11+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(499774), int32(602), int32(20832))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v41 = F_strlen(m, l0)
	mBase = m.M
	v42 = v41
	goto L17
L16:
	;
	v42 = v3
	goto L17
L17:
	;
	if l1 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v43 = F_strlen(m, l1)
	mBase = m.M
	v44 = v43
	goto L20
L19:
	;
	v44 = v3
	goto L20
L20:
	;
	if v42 == int32(0) {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	if base.Ui32(int32(64)) <= base.Ui32(v42) {
		goto L3
	} else {
		goto L22
	}
L22:
	;
	if base.Ui32(int32(8000)) <= base.Ui32(v44) {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v51 = int32(4515248)
	v52 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v55 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v55
	v60 = F_palloc(m, v42+v44+int32(6))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)) = uint16(v44)
	*(*uint16)(unsafe.Add(mBase, uint32(v60))) = uint16(v42)
	v65 = v60 + int32(4)
	if (l0^v65)&int32(3) != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	if l1 != 0 {
		goto L47
	} else {
		goto L48
	}
L26:
	;
	goto L25
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v119)
	if v119&int32(255) == int32(0) {
		goto L26
	} else {
		goto L42
	}
L28:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v118 = l0
	v119 = v71
	v120 = v65
	goto L27
L29:
	;
	goto L30
L30:
	;
	if l0&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v75 = l0
	v77 = v65
	goto L34
L32:
	;
	v89 = l0
	v91 = v65
	goto L33
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v96 = int32(-2139062144)
	if (int32(16843008)-v93|v93)&v96 != v96 {
		v118 = v89
		v119 = v93
		v120 = v91
		goto L27
	} else {
		goto L38
	}
L34:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	*(*uint8)(unsafe.Add(mBase, uint32(v77))) = uint8(v78)
	if v78 == int32(0) {
		goto L26
	} else {
		goto L36
	}
L35:
	;
	v89 = v85
	v91 = v83
	goto L33
L36:
	;
	v82 = int32(1)
	v83 = v77 + v82
	v85 = v75 + v82
	if v85&int32(3) != 0 {
		v75 = v85
		v77 = v83
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v101 = v89
	v102 = v93
	v103 = v91
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v102
	v105 = int32(4)
	v106 = v103 + v105
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v109 = v101 + v105
	v113 = int32(-2139062144)
	if (v107|(int32(16843008)-v107))&v113 == v113 {
		v101 = v109
		v102 = v107
		v103 = v106
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v118 = v109
	v119 = v107
	v120 = v106
	goto L27
L41:
	;
	goto L40
L42:
	;
	v127 = v118
	v129 = v120
	goto L43
L43:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v129)+1)) = uint8(v130)
	v132 = int32(1)
	if v130 != 0 {
		v127 = v127 + v132
		v129 = v129 + v132
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L26
L45:
	;
	goto L44
L46:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	if v221 != 0 {
		goto L73
	} else {
		goto L74
	}
L47:
	;
	v142 = v42 + v65 + int32(1)
	if (l1^v142)&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	goto L49
L49:
	;
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42+v65)+1)) = uint8(v218)
	goto L46
L50:
	;
	goto L46
L51:
	;
	goto L50
L52:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v197))) = uint8(v196)
	if v196&int32(255) == int32(0) {
		goto L51
	} else {
		goto L67
	}
L53:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v195 = l1
	v196 = v148
	v197 = v142
	goto L52
L54:
	;
	goto L55
L55:
	;
	if l1&int32(3) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v152 = l1
	v154 = v142
	goto L59
L57:
	;
	v166 = l1
	v168 = v142
	goto L58
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v173 = int32(-2139062144)
	if (int32(16843008)-v170|v170)&v173 != v173 {
		v195 = v166
		v196 = v170
		v197 = v168
		goto L52
	} else {
		goto L63
	}
L59:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v155)
	if v155 == int32(0) {
		goto L51
	} else {
		goto L61
	}
L60:
	;
	v166 = v162
	v168 = v160
	goto L58
L61:
	;
	v159 = int32(1)
	v160 = v154 + v159
	v162 = v152 + v159
	if v162&int32(3) != 0 {
		v152 = v162
		v154 = v160
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v178 = v166
	v179 = v170
	v180 = v168
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v179
	v182 = int32(4)
	v183 = v180 + v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v186 = v178 + v182
	v190 = int32(-2139062144)
	if (v184|(int32(16843008)-v184))&v190 == v190 {
		v178 = v186
		v179 = v184
		v180 = v183
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v195 = v186
	v196 = v184
	v197 = v183
	goto L52
L66:
	;
	goto L65
L67:
	;
	v204 = v195
	v206 = v197
	goto L68
L68:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)) = uint8(v207)
	v209 = int32(1)
	if v207 != 0 {
		v204 = v204 + v209
		v206 = v206 + v209
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L51
L70:
	;
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v52
	m.G0 = v11 + int32(32)
	return
L72:
	;
	v245 = F_AsyncExistsPendingNotify(m, v60)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L79
	}
L73:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	if v15 <= v222 {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v227 = F_MemoryContextAlloc(m, v225, int32(16))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v60
	v235 = F_list_make1_impl(m, int32(1), v11+int32(12))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v227)+4)) = v235
	v240 = int32(4412200)
	v241 = *(*int32)(unsafe.Add(mBase, _consts[351]))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+12)) = v241
	*(*int32)(unsafe.Add(mBase, _consts[351])) = v227
	goto L71
L79:
	;
	if v245 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	F_pfree(m, v60)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	F_AddEventToPendingNotifies(m, v60)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L10
	} else {
		goto L84
	}
L83:
	;
	goto L71
L84:
	;
	goto L71
L85:
	;
	F_errmsg_internal(m, int32(220809), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(499774), int32(599), int32(20832))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	F_errmsg(m, int32(8861), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(499774), int32(611), int32(20832))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	F_errmsg(m, int32(327997), int32(0))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(499774), int32(617), int32(20832))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	F_errmsg(m, int32(327900), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(499774), int32(622), int32(20832))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_Async_UnlistenAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[350])))
	if v7 != int32(1) {
		v28 = *(*int32)(unsafe.Add(mBase, _consts[352]))
		if v28 == int32(0) {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[353])))
			if v32 == int32(0) {
				m.G0 = v4 + int32(16)
				return
			} else {
				F_queue_listen(m, int32(2), int32(757108))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					m.G0 = v4 + int32(16)
					return
				}
			}
		} else {
			F_queue_listen(m, int32(2), int32(757108))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				m.G0 = v4 + int32(16)
				return
			}
		}
	} else {
		v12 = F_errstart(m, int32(14), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			if v12 == int32(0) {
				v28 = *(*int32)(unsafe.Add(mBase, _consts[352]))
				if v28 == int32(0) {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[353])))
					if v32 == int32(0) {
						m.G0 = v4 + int32(16)
						return
					} else {
						F_queue_listen(m, int32(2), int32(757108))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							m.G0 = v4 + int32(16)
							return
						}
					}
				} else {
					F_queue_listen(m, int32(2), int32(757108))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						m.G0 = v4 + int32(16)
						return
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[354]))
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = v17
				F_errmsg_internal(m, int32(678417), v4)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(499774), int32(772), int32(305238))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, _consts[352]))
						if v28 == int32(0) {
							v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[353])))
							if v32 == int32(0) {
								m.G0 = v4 + int32(16)
								return
							} else {
								F_queue_listen(m, int32(2), int32(757108))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									m.G0 = v4 + int32(16)
									return
								}
							}
						} else {
							F_queue_listen(m, int32(2), int32(757108))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								m.G0 = v4 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_Async_UnlistenOnExit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _consts[350])))
	if v8 != int32(1) {
		m.Env.Pgmem_listen(m, int32(757108), int32(2))
		mBase = m.M
		v32 = *(*int32)(unsafe.Add(mBase, _consts[355]))
		F_list_free_deep(m, v32)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[355])) = int32(0)
			F_asyncQueueUnregister(m)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	} else {
		v13 = F_errstart(m, int32(14), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			if v13 == int32(0) {
				m.Env.Pgmem_listen(m, int32(757108), int32(2))
				mBase = m.M
				v32 = *(*int32)(unsafe.Add(mBase, _consts[355]))
				F_list_free_deep(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[355])) = int32(0)
					F_asyncQueueUnregister(m)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, _consts[354]))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v18
				F_errmsg_internal(m, int32(678271), v5)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errfinish(m, int32(499774), int32(1205), int32(100730))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						m.Env.Pgmem_listen(m, int32(757108), int32(2))
						mBase = m.M
						v32 = *(*int32)(unsafe.Add(mBase, _consts[355]))
						F_list_free_deep(m, v32)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[355])) = int32(0)
							F_asyncQueueUnregister(m)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return
							} else {
								m.G0 = v5 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecAsyncAppendResponse(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v5 == int32(1) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v9 != 0 {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
			if v10&int32(2) == int32(0) {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)+136))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+136)) = v19 + int32(1)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v8)+132))
				*(*int32)(unsafe.Add(mBase, uint32(v23+v19<<(uint(int32(2))%32)))) = v9
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+148))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v30 = F_bms_add_member(m, v28, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = v30
					return
				}
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v15 - int32(1)
				return
			}
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+144))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v15 - int32(1)
			return
		}
	} else {
		return
	}
}
