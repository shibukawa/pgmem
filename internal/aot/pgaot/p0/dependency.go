package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dependency_is_compatible_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
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
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == int32(318) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v266
L2:
	;
	if v190 == int32(27) {
		goto L64
	} else {
		goto L65
	}
L3:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v189 = v185
	v190 = v188
	goto L2
L4:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)))
	if v162 != int32(1) {
		v266 = v4
		goto L1
	} else {
		goto L57
	}
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v17 != 0 {
		v266 = v4
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v73 = l0
	v74 = v14
	goto L7
L7:
	;
	switch v74 - int32(17) {
	case 0:
		goto L30
	default:
		v189 = v73
		v190 = v74
		goto L2
	case 3:
		v160 = v73
		goto L4
	case 4:
		goto L29
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v18 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v65 != int32(1) {
		v266 = v4
		goto L1
	} else {
		goto L25
	}
L10:
	;
	v65 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v27 = int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v28 <= v27 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v31 = v27
	goto L15
L14:
	;
	v31 = v28
	goto L15
L15:
	;
	v34 = int32(0)
	v36 = v34
	v37 = v34
	goto L16
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(8)+v36<<(uint(int32(2))%32))))
	if v45 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v65 = v58
	goto L9
L18:
	;
	goto L17
L19:
	;
	v46 = int32(2)
	if v37 != 0 {
		v58 = v46
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v51 = v37
	goto L21
L21:
	;
	v54 = v36 + int32(1)
	if v54 != v31 {
		v36 = v54
		v37 = v51
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v47 = int32(1)
	if base.Ui32(v47) < base.Ui32(base.I32_popcnt(v45)) {
		v58 = v46
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v51 = v47
	goto L21
L24:
	;
	v58 = v51
	goto L18
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v68 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v160 = int32(0)
	goto L4
L27:
	;
	goto L28
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v73 = v68
	v74 = v72
	goto L7
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	switch v109 - int32(1) {
	case 0:
		goto L43
	case 1:
		goto L42
	default:
		v185 = v73
		goto L3
	}
L30:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	if v77 == int32(0) {
		v266 = v4
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	if v80 != int32(2) {
		v266 = v4
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v85 = F_is_pseudo_constant_clause(m, v84)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v85 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v94 = F_is_pseudo_constant_clause(m, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v102 = v90
	goto L37
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v105 = F_get_oprrest(m, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	if v94 == int32(0) {
		v266 = v4
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+12))
	v102 = v99 + int32(4)
	goto L37
L40:
	;
	if v105 == int32(101) {
		v185 = v104
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v266 = v4
	goto L1
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v185 = v159
	goto L3
L43:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v112
	v115 = int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
	if v116 == v112 {
		v266 = v115
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v119 <= int32(0) {
		v266 = v115
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v125 = v112
	goto L46
L46:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131+v125<<(uint(int32(2))%32))))
	v140 = F_dependency_is_compatible_expression(m, v137, l1, v12+int32(12))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L33
	} else {
		goto L49
	}
L47:
	;
	v266 = int32(0)
	goto L1
L48:
	;
	goto L47
L49:
	;
	if v140 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v145 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v147 = v145
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v144
	v147 = v144
	goto L53
L53:
	;
	v148 = F_equal(m, v144, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L33
	} else {
		goto L54
	}
L54:
	;
	if v148 == int32(0) {
		v266 = v148
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v153 = v125 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v153 < v154 {
		v125 = v153
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v266 = v148
	goto L1
L57:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)+28))
	if v165 == int32(0) {
		v266 = v4
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v168 != int32(2) {
		v266 = v4
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+4))
	v173 = F_is_pseudo_constant_clause(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L33
	} else {
		goto L60
	}
L60:
	;
	if v173 == int32(0) {
		v266 = v4
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v160)+28))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	v181 = F_get_oprrest(m, v177)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	if v181 != int32(101) {
		v266 = v4
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v185 = v180
	goto L3
L64:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	v195 = v194
	goto L66
L65:
	;
	v195 = v189
	goto L66
L66:
	;
	if l1 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v266 = v4
	goto L1
L68:
	;
	goto L69
L69:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v198 <= int32(0) {
		v266 = v4
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v208 = v4
	goto L71
L71:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210+v208<<(uint(int32(2))%32))))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+16)))
	if v215 != int32(102) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v266 = v4
	goto L1
L73:
	;
	v259 = v208 + int32(1)
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v259 < v260 {
		v208 = v259
		goto L71
	} else {
		goto L84
	}
L74:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214)+24))
	if v218 == int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v221 = int32(0)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v222 <= v221 {
		goto L73
	} else {
		goto L76
	}
L76:
	;
	v228 = v221
	goto L77
L77:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234+v228<<(uint(int32(2))%32))))
	v239 = F_equal(m, v195, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L33
	} else {
		goto L79
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v238
	v266 = int32(1)
	goto L1
L79:
	;
	if v239 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v244 = v228 + int32(1)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v244 < v245 {
		v228 = v244
		goto L77
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	goto L78
L83:
	;
	goto L73
L84:
	;
	goto L72
}
func F_recordDependencyOnExpr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	v10 = m.G0
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	v15 = F_palloc(m, v11)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(137438953472)
		v20 = F_palloc(m, int32(384))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v15
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l2
			v28 = int32(1)
			v30 = F_list_make1_impl(m, v28, v12)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v30
				v35 = F_find_expr_references_walker(m, l1, v12+int32(8))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
					if int32(2) <= v38 {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
						F_pg_qsort(m, v41, v38, int32(12), int32(462))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
							if int32(2) <= v46 {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
								v52 = int32(1)
								v54 = v49
								v57 = v28
								for {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
									v64 = v61 + v52*int32(12)
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
									if v60 != v65 {
										v77 = v54 + int32(12)
										v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
										*(*int64)(unsafe.Add(mBase, uint32(v77))) = v78
										v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v80
										v84 = v77
										v85 = v57 + int32(1)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										if v67 != v68 {
											v77 = v54 + int32(12)
											v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
											*(*int64)(unsafe.Add(mBase, uint32(v77))) = v78
											v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v80
											v84 = v77
											v85 = v57 + int32(1)
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
											if v70 == v71 {
												v84 = v54
												v85 = v57
											} else {
												if v70 != 0 {
													v77 = v54 + int32(12)
													v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
													*(*int64)(unsafe.Add(mBase, uint32(v77))) = v78
													v80 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v54)+20)) = v80
													v84 = v77
													v85 = v57 + int32(1)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v54)+8)) = v71
													v84 = v54
													v85 = v57
												}
											}
										}
									}
									v89 = v52 + int32(1)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
									if v89 < v90 {
										v52 = v89
										v54 = v84
										v57 = v85
										continue
									} else {
										break
									}
									break
								}
								v98 = v85
							} else {
								v98 = v28
							}
							*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v98
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
							v107 = v103
							v109 = v102
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
							F_recordMultipleDependencies(m, l0, v113, v107, int32(110))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return
							} else {
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
								F_pfree(m, v118)
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return
								} else {
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
									if v121 != 0 {
										F_pfree(m, v121)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											F_pfree(m, v117)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										}
									} else {
										F_pfree(m, v117)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								}
							}
						}
					} else {
						v107 = v38
						v109 = v37
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
						F_recordMultipleDependencies(m, l0, v113, v107, int32(110))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							v117 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
							F_pfree(m, v118)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
								if v121 != 0 {
									F_pfree(m, v121)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										F_pfree(m, v117)
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								} else {
									F_pfree(m, v117)
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
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
