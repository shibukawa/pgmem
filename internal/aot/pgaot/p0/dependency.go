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
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
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
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
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
	return v264
L2:
	;
	if v188 == int32(27) {
		goto L64
	} else {
		goto L65
	}
L3:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v187 = v184
	v188 = v186
	goto L2
L4:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)))
	if v161 != int32(1) {
		v264 = v4
		goto L1
	} else {
		goto L57
	}
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v17 != 0 {
		v264 = v4
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v72 = l0
	v73 = v14
	goto L7
L7:
	;
	switch v73 - int32(17) {
	case 0:
		goto L30
	default:
		v187 = v72
		v188 = v73
		goto L2
	case 3:
		v159 = v72
		goto L4
	case 4:
		goto L29
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = int32(0)
	if v18 == v19 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v64 != int32(1) {
		v264 = v4
		goto L1
	} else {
		goto L25
	}
L10:
	;
	v64 = int32(0)
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
	v35 = int32(0)
	v37 = v19
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(8)+v35<<(uint(int32(2))%32))))
	if v44 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v64 = v56
	goto L9
L18:
	;
	goto L17
L19:
	;
	v45 = int32(2)
	if v37 != 0 {
		v56 = v45
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
	v53 = v35 + int32(1)
	if v53 != v31 {
		v35 = v53
		v37 = v51
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v46 = int32(1)
	if base.Ui32(v46) < base.Ui32(base.I32_popcnt(v44)) {
		v56 = v45
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v51 = v46
	goto L21
L24:
	;
	v56 = v51
	goto L18
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v159 = int32(0)
	goto L4
L27:
	;
	goto L28
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v72 = v67
	v73 = v71
	goto L7
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	switch v108 - int32(1) {
	case 0:
		goto L43
	case 1:
		goto L42
	default:
		v184 = v72
		goto L3
	}
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	if v76 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 != int32(2) {
		v264 = v4
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v84 = F_is_pseudo_constant_clause(m, v83)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	if v84 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v93 = F_is_pseudo_constant_clause(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v101 = v89
	goto L37
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v104 = F_get_oprrest(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	if v93 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v101 = v98 + int32(4)
	goto L37
L40:
	;
	if v104 == int32(101) {
		v184 = v103
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v264 = v4
	goto L1
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v184 = v158
	goto L3
L43:
	;
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v111
	v114 = int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v115 == v111 {
		v264 = v114
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v118 <= int32(0) {
		v264 = v114
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v124 = v111
	goto L46
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+v124<<(uint(int32(2))%32))))
	v139 = F_dependency_is_compatible_expression(m, v136, l1, v12+int32(12))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L33
	} else {
		goto L49
	}
L47:
	;
	v264 = int32(0)
	goto L1
L48:
	;
	goto L47
L49:
	;
	if v139 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v144 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v146 = v144
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v143
	v146 = v143
	goto L53
L53:
	;
	v147 = F_equal(m, v143, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L33
	} else {
		goto L54
	}
L54:
	;
	if v147 == int32(0) {
		v264 = v147
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v152 = v124 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v152 < v153 {
		v124 = v152
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v264 = v147
	goto L1
L57:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	if v164 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v167 != int32(2) {
		v264 = v4
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	v172 = F_is_pseudo_constant_clause(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L33
	} else {
		goto L60
	}
L60:
	;
	if v172 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v180 = F_get_oprrest(m, v176)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	if v180 != int32(101) {
		v264 = v4
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v184 = v179
	goto L3
L64:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v192 = v191
	goto L66
L65:
	;
	v192 = v187
	goto L66
L66:
	;
	if l1 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v195 <= int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v205 = v4
	goto L69
L69:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207+v205<<(uint(int32(2))%32))))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+16)))
	if v212 != int32(102) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v264 = v4
	goto L1
L71:
	;
	v256 = v205 + int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v256 < v257 {
		v205 = v256
		goto L69
	} else {
		goto L82
	}
L72:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	if v215 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v218 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v219 <= v218 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v225 = v218
	goto L75
L75:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v225<<(uint(int32(2))%32))))
	v236 = F_equal(m, v192, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L33
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v235
	v264 = int32(1)
	goto L1
L77:
	;
	if v236 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v241 = v225 + int32(1)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v241 < v242 {
		v225 = v241
		goto L75
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	goto L76
L81:
	;
	goto L71
L82:
	;
	goto L70
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
	var v53 int32
	_ = v53
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
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
								v53 = v49
								v57 = v28
								for {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
									v64 = v61 + v52*int32(12)
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
									if v60 != v65 {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v76
										v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
										*(*int64)(unsafe.Add(mBase, uint32(v53)+12)) = v78
										v84 = v53 + int32(12)
										v85 = v57 + int32(1)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										if v67 != v68 {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v76
											v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
											*(*int64)(unsafe.Add(mBase, uint32(v53)+12)) = v78
											v84 = v53 + int32(12)
											v85 = v57 + int32(1)
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
											if v70 == v71 {
												v84 = v53
												v85 = v57
											} else {
												if v70 != 0 {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v76
													v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
													*(*int64)(unsafe.Add(mBase, uint32(v53)+12)) = v78
													v84 = v53 + int32(12)
													v85 = v57 + int32(1)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v71
													v84 = v53
													v85 = v57
												}
											}
										}
									}
									v89 = v52 + int32(1)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
									if v89 < v90 {
										v52 = v89
										v53 = v84
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
							v106 = v103
							v107 = v102
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
							F_recordMultipleDependencies(m, l0, v113, v106, int32(110))
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
						v106 = v38
						v107 = v37
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
						F_recordMultipleDependencies(m, l0, v113, v106, int32(110))
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
