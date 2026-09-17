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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
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
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v262 int32
	_ = v262
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
	return v262
L2:
	;
	v262 = int32(1)
	goto L1
L3:
	;
	v262 = int32(0)
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
	v77 = int32(0)
	v78 = v16
	goto L6
L6:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+76))
	if v80 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L7:
	;
	v70 = int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v71 == v70 {
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
	v262 = v49
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
	v262 = v55
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
		v262 = v59
		goto L1
	} else {
		goto L29
	}
L28:
	;
	if v67 != 0 {
		v262 = v67
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
	v77 = v70
	v78 = v16 + int32(36)
	goto L6
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v83 <= v77 {
		goto L3
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)+152))
	if v144 != 0 {
		goto L62
	} else {
		goto L63
	}
L35:
	;
	v86 = v77
	v88 = v78
	goto L36
L36:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(v96+v86<<(uint(int32(1))%32))+12)))
	v103 = F_heap_getattr_1(m, v20, v100, v18, v14+int32(335))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L9
	} else {
		goto L38
	}
L37:
	;
	goto L3
L38:
	;
	v107 = F_heap_getattr_1(m, v19, v100, v18, v14+int32(334))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+334)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+335)))
	if v110 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v140 = v86 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v140 < v141 {
		v86 = v140
		v88 = v88 + int32(36)
		goto L36
	} else {
		goto L60
	}
L41:
	;
	if v109&int32(1) != 0 {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v109&int32(1) != 0 {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+9)))
	if v117 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v118 = int32(-1)
	goto L47
L46:
	;
	v118 = int32(1)
	goto L47
L47:
	;
	v262 = v118
	goto L1
L48:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+9)))
	if v123 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v88)+16))
	v126 = m.T0[v125].(func(*base.Module, int32, int32, int32) int32)(m, v103, v107, v88)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L54
	}
L51:
	;
	v124 = int32(1)
	goto L53
L52:
	;
	v124 = int32(-1)
	goto L53
L53:
	;
	v262 = v124
	goto L1
L54:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+8)))
	if v128 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v126 < int32(0) {
		goto L2
	} else {
		goto L58
	}
L56:
	;
	v135 = v126
	goto L57
L57:
	;
	if v135 != 0 {
		v262 = v135
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v135 = int32(0) - v126
	goto L57
L59:
	;
	goto L40
L60:
	;
	goto L37
L61:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v158 = F_ExecStoreHeapTuple(m, v20, v156, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L9
	} else {
		goto L68
	}
L62:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+20))
	F_MemoryContextReset(m, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L9
	} else {
		goto L65
	}
L63:
	;
	v150 = v143
	goto L64
L64:
	;
	v152 = F_MakePerTupleExprContext(m, v150)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L9
	} else {
		goto L67
	}
L65:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+152))
	if v149 != 0 {
		v155 = v149
		goto L61
	} else {
		goto L66
	}
L66:
	;
	v150 = v148
	goto L64
L67:
	;
	v155 = v152
	goto L61
L68:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_FormIndexDatum(m, v160, v156, v161, v14+int32(192), v14+int32(160))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v169 = F_ExecStoreHeapTuple(m, v19, v156, int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L70
	}
L70:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	F_FormIndexDatum(m, v171, v156, v172, v14+int32(32), v14)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	if v177 <= v77 {
		goto L3
	} else {
		goto L72
	}
L72:
	;
	v179 = v177
	v180 = v77
	v182 = v78
	goto L73
L73:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v14))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+int32(160)+v180))))
	if v195 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L3
L75:
	;
	v236 = v180 + int32(1)
	if v236 < v232 {
		v179 = v232
		v180 = v236
		v182 = v182 + int32(36)
		goto L73
	} else {
		goto L95
	}
L76:
	;
	if v191&int32(1) != 0 {
		v232 = v179
		goto L75
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	if v191&int32(1) != 0 {
		goto L83
	} else {
		goto L84
	}
L79:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+9)))
	if v202 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v203 = int32(-1)
	goto L82
L81:
	;
	v203 = int32(1)
	goto L82
L82:
	;
	v262 = v203
	goto L1
L83:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+9)))
	if v208 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	v211 = v180 << (uint(int32(2)) % 32)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211+(v14+int32(192)))))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(32)+v211)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v182)+16))
	v221 = m.T0[v220].(func(*base.Module, int32, int32, int32) int32)(m, v215, v219, v182)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L9
	} else {
		goto L89
	}
L86:
	;
	v209 = int32(1)
	goto L88
L87:
	;
	v209 = int32(-1)
	goto L88
L88:
	;
	v262 = v209
	goto L1
L89:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
	if v223 == int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v221 < int32(0) {
		goto L2
	} else {
		goto L93
	}
L91:
	;
	v230 = v221
	goto L92
L92:
	;
	if v230 != 0 {
		v262 = v230
		goto L1
	} else {
		goto L94
	}
L93:
	;
	v230 = int32(0) - v221
	goto L92
L94:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v232 = v231
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
