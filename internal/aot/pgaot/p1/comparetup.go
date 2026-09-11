package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_comparetup_cluster(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+36)))
	if v6 != int32(1) {
		v45 = F_comparetup_cluster_tiebreak(m, l0, l1, l2)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			v47 = v45
			return v47
		}
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
		if v11 == int32(1) {
			if v9&int32(1) != 0 {
				v45 = F_comparetup_cluster_tiebreak(m, l0, l1, l2)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = v45
					return v47
				}
			} else {
				v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+9)))
				if v18 != 0 {
					v19 = int32(-1)
				} else {
					v19 = int32(1)
				}
				return v19
			}
		} else {
			if v9&int32(1) != 0 {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+9)))
				if v25 != 0 {
					v26 = int32(1)
				} else {
					v26 = int32(-1)
				}
				return v26
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
				v31 = m.T0[v30].(func(*base.Module, int32, int32, int32) int32)(m, v28, v29, v10)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)))
					if v35 == int32(1) {
						if v31 < int32(0) {
							return int32(1)
						} else {
							v42 = int32(0) - v31
							if v42 != 0 {
								v47 = v42
								return v47
							} else {
								v45 = F_comparetup_cluster_tiebreak(m, l0, l1, l2)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									v47 = v45
									return v47
								}
							}
						}
					} else {
						v42 = v31
						if v42 != 0 {
							v47 = v42
							return v47
						} else {
							v45 = F_comparetup_cluster_tiebreak(m, l0, l1, l2)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								v47 = v45
								return v47
							}
						}
					}
				}
			}
		}
	}
}
func F_comparetup_cluster_tiebreak(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
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
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v265 int32
	_ = v265
	v12 = m.G0
	v14 = v12 - int32(336)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+36)))
	if v22 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v14 + int32(336)
	return v265
L2:
	;
	v265 = int32(1)
	goto L1
L3:
	;
	v265 = int32(0)
	goto L1
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	if v25 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v78 = int32(0)
	v79 = v16
	goto L6
L6:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+76))
	if v82 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L7:
	;
	v71 = int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v72 == v71 {
		goto L3
	} else {
		goto L31
	}
L8:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+12)))
	v32 = F_heap_getattr_1(m, v20, v29, v18, v14+int32(335))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v38 = F_heap_getattr_1(m, v19, v29, v18, v14+int32(334))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+334)))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+335)))
	if v41 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v40&int32(1) != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v40&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)))
	if v48 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = int32(-1)
	goto L18
L17:
	;
	v49 = int32(1)
	goto L18
L18:
	;
	v265 = v49
	goto L1
L19:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)))
	if v54 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v57 = m.T0[v56].(func(*base.Module, int32, int32, int32) int32)(m, v32, v38, v16)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L9
	} else {
		goto L25
	}
L22:
	;
	v55 = int32(1)
	goto L24
L23:
	;
	v55 = int32(-1)
	goto L24
L24:
	;
	v265 = v55
	goto L1
L25:
	;
	v59 = int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
	if v60 != v59 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v67 = v57
	goto L28
L27:
	;
	if v57 < int32(0) {
		v265 = v59
		goto L1
	} else {
		goto L29
	}
L28:
	;
	if v67 != 0 {
		v265 = v67
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v67 = int32(0) - v57
	goto L28
L30:
	;
	goto L7
L31:
	;
	v78 = v71
	v79 = v16 + int32(36)
	goto L6
L32:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v85 <= v78 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+152))
	if v146 != 0 {
		goto L62
	} else {
		goto L63
	}
L35:
	;
	v88 = v78
	v90 = v79
	goto L36
L36:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v98+v88<<(uint(int32(1))%32))+12)))
	v105 = F_heap_getattr_1(m, v20, v102, v18, v14+int32(335))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L38
	}
L37:
	;
	goto L3
L38:
	;
	v109 = F_heap_getattr_1(m, v19, v102, v18, v14+int32(334))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+334)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+335)))
	if v112 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v142 = v88 + int32(1)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v142 < v143 {
		v88 = v142
		v90 = v90 + int32(36)
		goto L36
	} else {
		goto L60
	}
L41:
	;
	if v111&int32(1) != 0 {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v111&int32(1) != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+9)))
	if v119 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v120 = int32(-1)
	goto L47
L46:
	;
	v120 = int32(1)
	goto L47
L47:
	;
	v265 = v120
	goto L1
L48:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+9)))
	if v125 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v128 = m.T0[v127].(func(*base.Module, int32, int32, int32) int32)(m, v105, v109, v90)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L54
	}
L51:
	;
	v126 = int32(1)
	goto L53
L52:
	;
	v126 = int32(-1)
	goto L53
L53:
	;
	v265 = v126
	goto L1
L54:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+8)))
	if v130 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v128 < int32(0) {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	v137 = v128
	goto L57
L57:
	;
	if v137 != 0 {
		v265 = v137
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v137 = int32(0) - v128
	goto L57
L59:
	;
	goto L40
L60:
	;
	goto L37
L61:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v160 = F_ExecStoreHeapTuple(m, v20, v158, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L68
	}
L62:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+20))
	F_MemoryContextReset(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L9
	} else {
		goto L65
	}
L63:
	;
	v152 = v145
	goto L64
L64:
	;
	v154 = F_MakePerTupleExprContext(m, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L67
	}
L65:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+152))
	if v151 != 0 {
		v157 = v151
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v152 = v150
	goto L64
L67:
	;
	v157 = v154
	goto L61
L68:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_FormIndexDatum(m, v162, v158, v163, v14+int32(192), v14+int32(160))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v171 = F_ExecStoreHeapTuple(m, v19, v158, int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_FormIndexDatum(m, v173, v158, v174, v14+int32(32), v14)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v179 <= v78 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	v182 = v78
	v184 = v79
	v186 = v179
	goto L73
L73:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182+v14))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(160)+v182))))
	if v197 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L3
L75:
	;
	v239 = v182 + int32(1)
	if v239 < v235 {
		v182 = v239
		v184 = v184 + int32(36)
		v186 = v235
		goto L73
	} else {
		goto L95
	}
L76:
	;
	if v193&int32(1) != 0 {
		v235 = v186
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v193&int32(1) != 0 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+9)))
	if v204 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v205 = int32(-1)
	goto L82
L81:
	;
	v205 = int32(1)
	goto L82
L82:
	;
	v265 = v205
	goto L1
L83:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+9)))
	if v210 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v213 = v182 << (uint(int32(2)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+(v14+int32(192)))))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(32)+v213)))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v223 = m.T0[v222].(func(*base.Module, int32, int32, int32) int32)(m, v217, v221, v184)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L89
	}
L86:
	;
	v211 = int32(1)
	goto L88
L87:
	;
	v211 = int32(-1)
	goto L88
L88:
	;
	v265 = v211
	goto L1
L89:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+8)))
	if v225 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v223 < int32(0) {
		goto L2
	} else {
		goto L93
	}
L91:
	;
	v232 = v223
	goto L92
L92:
	;
	if v232 != 0 {
		v265 = v232
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v232 = int32(0) - v223
	goto L92
L94:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v235 = v233
	goto L75
L95:
	;
	goto L74
}
func F_comparetup_index_brin(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	return base.B2i32(base.Ui32(v5) < base.Ui32(v4)) - base.B2i32(base.Ui32(v4) < base.Ui32(v5))
}
