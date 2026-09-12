package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecParallelHashIncreaseNumBuckets(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
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
	var v60 int32
	_ = v60
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v14 = v12 + int32(128)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v17 = base.I32_rem_s(v15, int32(3))
	switch v17 {
	case 0:
		goto L4
	case 1:
		goto L3
	case 2:
		goto L2
	default:
		goto L1
	}
L1:
	;
	return
L2:
	;
	F_ExecParallelHashEnsureBatchAccessors(m, l0)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L18
	}
L3:
	;
	v108 = F_BarrierArriveAndWait(m, v14, int32(134217755))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L17
	}
L4:
	;
	v19 = F_BarrierArriveAndWait(m, v14, int32(134217754))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v19 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v24 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v23 << (uint(v24) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+44))
	v31 = v23 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+44)) = v29 + int32(base.Ui32(v31)>>(uint(v24)%32))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	F_dsa_free(m, v36, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v42 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v45 = F_dsa_allocate_extended(m, v43, v31, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v45
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = F_dsa_get_address(m, v50, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if int32(0) < v56 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v60 = v42
	goto L14
L12:
	;
	goto L13
L13:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v92
	goto L3
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54+v60<<(uint(int32(2))%32)))) = int32(0)
	v76 = v60 + int32(1)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v76 < v77 {
		v60 = v76
		goto L14
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	goto L15
L17:
	;
	goto L2
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v129 = F_dsa_get_address(m, v125, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v129
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v133
	v136 = int32(1073741823)
	if v136 <= v133 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v147
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+24)) = uint8(v148)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v158 = v156 + int32(40)
	v160 = F_LWLockAcquire(m, v158, v148)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L27
	}
L21:
	;
	v139 = v136
	goto L23
L22:
	;
	v139 = v133
	goto L23
L23:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v139) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v147 = int32(32) - base.I32_clz(v139-int32(1))
	goto L26
L25:
	;
	v147 = int32(0)
	goto L26
L26:
	;
	goto L20
L27:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156)+24))
	if v162 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v166 = v162
	v167 = v156 + int32(24)
	v168 = v158
	goto L31
L29:
	;
	v284 = v158
	goto L30
L30:
	;
	F_LWLockRelease(m, v284)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L59
	}
L31:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v177 = F_dsa_get_address(m, v176, v166)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L33
	}
L32:
	;
	v284 = v274
	goto L30
L33:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v179
	F_LWLockRelease(m, v168)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	if v183 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v184 = int32(16)
	v193 = int32(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v269 != 0 {
		goto L53
	} else {
		goto L54
	}
L38:
	;
	v200 = v193 + (v177 + v184)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v209 = v201 + v202&(v203-int32(1))<<(uint(int32(2))%32)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v210
	v212 = v193 + (v166 + v184)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v214 = base.B2i32(v213 == v210)
	if v213 == v210 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L37
L40:
	;
	v215 = v212
	goto L42
L41:
	;
	v215 = v213
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v215
	if v214 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v220 = v213
	goto L46
L44:
	;
	goto L45
L45:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v254 = (v249+int32(15))&int32(-8) + v193
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	if base.Ui32(v254) < base.Ui32(v255) {
		v193 = v254
		goto L38
	} else {
		goto L52
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v220
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v233 == v231 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	v235 = v212
	goto L50
L49:
	;
	v235 = v233
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v235
	if v233 != v231 {
		v220 = v233
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	goto L39
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v274 = v272 + int32(40)
	v276 = F_LWLockAcquire(m, v274, int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v272)+24))
	if v280 != 0 {
		v166 = v280
		v167 = v272 + int32(24)
		v168 = v274
		goto L31
	} else {
		goto L58
	}
L58:
	;
	goto L32
L59:
	;
	v295 = F_BarrierArriveAndWait(m, v14, int32(134217756))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	goto L1
}
func F_ExecParallelReInitializeDSM(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
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
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
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
	var v159 int64
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	if l0 == int32(0) {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v12 - int32(397) {
		case 0:
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+36)))
			if v58 != int32(1) {
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				v62 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v62
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				v69 = F__emscripten_memset_bulkmem(m, v61+int32(20), base.I32_extend8_s(v62), v67)
				mBase = m.M
			}
			v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
			mBase = m.M
			v195 = m.ExcPending
			if v195 != 0 {
				return int32(0)
			} else {
				return v194
			}
		default:
			v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
			mBase = m.M
			v195 = m.ExcPending
			if v195 != 0 {
				return int32(0)
			} else {
				return v194
			}
		case 6:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+36)))
			if v16 != int32(1) {
				v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return int32(0)
				} else {
					return v194
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+188))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
				m.T0[v23].(func(*base.Module, int32, int32))(m, v19, v21)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return int32(0)
					} else {
						return v194
					}
				}
			}
		case 8:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+36)))
			if v29 != int32(1) {
				v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return int32(0)
				} else {
					return v194
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				F_index_parallelrescan(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return int32(0)
					} else {
						return v194
					}
				}
			}
		case 9:
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+36)))
			if v36 != int32(1) {
				v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return int32(0)
				} else {
					return v194
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
				F_index_parallelrescan(m, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return int32(0)
					} else {
						return v194
					}
				}
			}
		case 11:
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+36)))
			if v86 != int32(1) {
				v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return int32(0)
				} else {
					return v194
				}
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+172))
				if v90 != 0 {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
					*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = int32(0)
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
					if v94 != 0 {
						v95 = F_dsa_get_address(m, v90, v94)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
							if v97 == int32(0) {
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
								if v112 == int32(0) {
									v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
									if v127 == int32(0) {
										F_dsa_free(m, v90, v94)
										mBase = m.M
										v143 = m.ExcPending
										if v143 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
											v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
											mBase = m.M
											v195 = m.ExcPending
											if v195 != 0 {
												return int32(0)
											} else {
												return v194
											}
										}
									} else {
										v130 = F_dsa_get_address(m, v90, v127)
										mBase = m.M
										v131 = m.ExcPending
										if v131 != 0 {
											return int32(0)
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
											v133 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
											if v132 != v133 {
												F_dsa_free(m, v90, v94)
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
													v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
														return int32(0)
													} else {
														return v194
													}
												}
											} else {
												v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
												F_dsa_free(m, v90, v138)
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													F_dsa_free(m, v90, v94)
													mBase = m.M
													v143 = m.ExcPending
													if v143 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
														v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															return v194
														}
													}
												}
											}
										}
									}
								} else {
									v115 = F_dsa_get_address(m, v90, v112)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
										v118 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v115))) = v117 - v118
										if v117 != v118 {
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
											if v127 == int32(0) {
												F_dsa_free(m, v90, v94)
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
													v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
														return int32(0)
													} else {
														return v194
													}
												}
											} else {
												v130 = F_dsa_get_address(m, v90, v127)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
													v133 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
													if v132 != v133 {
														F_dsa_free(m, v90, v94)
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
															v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v195 = m.ExcPending
															if v195 != 0 {
																return int32(0)
															} else {
																return v194
															}
														}
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
														F_dsa_free(m, v90, v138)
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	return v194
																}
															}
														}
													}
												}
											}
										} else {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
											F_dsa_free(m, v90, v123)
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
												if v127 == int32(0) {
													F_dsa_free(m, v90, v94)
													mBase = m.M
													v143 = m.ExcPending
													if v143 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
														v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															return v194
														}
													}
												} else {
													v130 = F_dsa_get_address(m, v90, v127)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int32(0)
													} else {
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
														v133 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
														if v132 != v133 {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	return v194
																}
															}
														} else {
															v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
															F_dsa_free(m, v90, v138)
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return int32(0)
															} else {
																F_dsa_free(m, v90, v94)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																	v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v195 = m.ExcPending
																	if v195 != 0 {
																		return int32(0)
																	} else {
																		return v194
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
							} else {
								v100 = F_dsa_get_address(m, v90, v97)
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
									v103 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v100))) = v102 - v103
									if v102 != v103 {
										v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
										if v112 == int32(0) {
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
											if v127 == int32(0) {
												F_dsa_free(m, v90, v94)
												mBase = m.M
												v143 = m.ExcPending
												if v143 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
													v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
														return int32(0)
													} else {
														return v194
													}
												}
											} else {
												v130 = F_dsa_get_address(m, v90, v127)
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return int32(0)
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
													v133 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
													if v132 != v133 {
														F_dsa_free(m, v90, v94)
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
															v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v195 = m.ExcPending
															if v195 != 0 {
																return int32(0)
															} else {
																return v194
															}
														}
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
														F_dsa_free(m, v90, v138)
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	return v194
																}
															}
														}
													}
												}
											}
										} else {
											v115 = F_dsa_get_address(m, v90, v112)
											mBase = m.M
											v116 = m.ExcPending
											if v116 != 0 {
												return int32(0)
											} else {
												v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
												v118 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v115))) = v117 - v118
												if v117 != v118 {
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
													if v127 == int32(0) {
														F_dsa_free(m, v90, v94)
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
															v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v195 = m.ExcPending
															if v195 != 0 {
																return int32(0)
															} else {
																return v194
															}
														}
													} else {
														v130 = F_dsa_get_address(m, v90, v127)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
															v133 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
															if v132 != v133 {
																F_dsa_free(m, v90, v94)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																	v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v195 = m.ExcPending
																	if v195 != 0 {
																		return int32(0)
																	} else {
																		return v194
																	}
																}
															} else {
																v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
																F_dsa_free(m, v90, v138)
																mBase = m.M
																v140 = m.ExcPending
																if v140 != 0 {
																	return int32(0)
																} else {
																	F_dsa_free(m, v90, v94)
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																		v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v195 = m.ExcPending
																		if v195 != 0 {
																			return int32(0)
																		} else {
																			return v194
																		}
																	}
																}
															}
														}
													}
												} else {
													v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
													F_dsa_free(m, v90, v123)
													mBase = m.M
													v125 = m.ExcPending
													if v125 != 0 {
														return int32(0)
													} else {
														v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
														if v127 == int32(0) {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	return v194
																}
															}
														} else {
															v130 = F_dsa_get_address(m, v90, v127)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
																v133 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
																if v132 != v133 {
																	F_dsa_free(m, v90, v94)
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																		v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v195 = m.ExcPending
																		if v195 != 0 {
																			return int32(0)
																		} else {
																			return v194
																		}
																	}
																} else {
																	v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
																	F_dsa_free(m, v90, v138)
																	mBase = m.M
																	v140 = m.ExcPending
																	if v140 != 0 {
																		return int32(0)
																	} else {
																		F_dsa_free(m, v90, v94)
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																			v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v195 = m.ExcPending
																			if v195 != 0 {
																				return int32(0)
																			} else {
																				return v194
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
									} else {
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
										F_dsa_free(m, v90, v108)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v112 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
											if v112 == int32(0) {
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
												if v127 == int32(0) {
													F_dsa_free(m, v90, v94)
													mBase = m.M
													v143 = m.ExcPending
													if v143 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
														v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v195 = m.ExcPending
														if v195 != 0 {
															return int32(0)
														} else {
															return v194
														}
													}
												} else {
													v130 = F_dsa_get_address(m, v90, v127)
													mBase = m.M
													v131 = m.ExcPending
													if v131 != 0 {
														return int32(0)
													} else {
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
														v133 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
														if v132 != v133 {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	return v194
																}
															}
														} else {
															v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
															F_dsa_free(m, v90, v138)
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return int32(0)
															} else {
																F_dsa_free(m, v90, v94)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																	v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v195 = m.ExcPending
																	if v195 != 0 {
																		return int32(0)
																	} else {
																		return v194
																	}
																}
															}
														}
													}
												}
											} else {
												v115 = F_dsa_get_address(m, v90, v112)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return int32(0)
												} else {
													v117 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
													v118 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v115))) = v117 - v118
													if v117 != v118 {
														v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
														if v127 == int32(0) {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v143 = m.ExcPending
															if v143 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return int32(0)
																} else {
																	return v194
																}
															}
														} else {
															v130 = F_dsa_get_address(m, v90, v127)
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
																v133 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
																if v132 != v133 {
																	F_dsa_free(m, v90, v94)
																	mBase = m.M
																	v143 = m.ExcPending
																	if v143 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																		v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v195 = m.ExcPending
																		if v195 != 0 {
																			return int32(0)
																		} else {
																			return v194
																		}
																	}
																} else {
																	v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
																	F_dsa_free(m, v90, v138)
																	mBase = m.M
																	v140 = m.ExcPending
																	if v140 != 0 {
																		return int32(0)
																	} else {
																		F_dsa_free(m, v90, v94)
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																			v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v195 = m.ExcPending
																			if v195 != 0 {
																				return int32(0)
																			} else {
																				return v194
																			}
																		}
																	}
																}
															}
														}
													} else {
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
														F_dsa_free(m, v90, v123)
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return int32(0)
														} else {
															v127 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
															if v127 == int32(0) {
																F_dsa_free(m, v90, v94)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																	v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v195 = m.ExcPending
																	if v195 != 0 {
																		return int32(0)
																	} else {
																		return v194
																	}
																}
															} else {
																v130 = F_dsa_get_address(m, v90, v127)
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return int32(0)
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
																	v133 = int32(1)
																	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v132 - v133
																	if v132 != v133 {
																		F_dsa_free(m, v90, v94)
																		mBase = m.M
																		v143 = m.ExcPending
																		if v143 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																			v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v195 = m.ExcPending
																			if v195 != 0 {
																				return int32(0)
																			} else {
																				return v194
																			}
																		}
																	} else {
																		v138 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
																		F_dsa_free(m, v90, v138)
																		mBase = m.M
																		v140 = m.ExcPending
																		if v140 != 0 {
																			return int32(0)
																		} else {
																			F_dsa_free(m, v90, v94)
																			mBase = m.M
																			v143 = m.ExcPending
																			if v143 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																				v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																				mBase = m.M
																				v195 = m.ExcPending
																				if v195 != 0 {
																					return int32(0)
																				} else {
																					return v194
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
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
						v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return int32(0)
						} else {
							return v194
						}
					}
				} else {
					v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return int32(0)
					} else {
						return v194
					}
				}
			}
		case 21:
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+36)))
			if v43 != int32(1) {
				v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return int32(0)
				} else {
					return v194
				}
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+152))
				if v47 != 0 {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v50 = int64(*(*int32)(unsafe.Add(mBase, uint32(v49)+40)))
					v52 = F_shm_toc_lookup(m, v48, v50, int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+152))
						m.T0[v54].(func(*base.Module, int32, int32, int32))(m, l0, l1, v52)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return int32(0)
							} else {
								return v194
							}
						}
					}
				} else {
					v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return int32(0)
					} else {
						return v194
					}
				}
			}
		case 22:
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+36)))
			if v71 != int32(1) {
				v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return int32(0)
				} else {
					return v194
				}
			} else {
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+36))
				if v75 != 0 {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v78 = int64(*(*int32)(unsafe.Add(mBase, uint32(v77)+40)))
					v80 = F_shm_toc_lookup(m, v76, v78, int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v74)+36))
						m.T0[v82].(func(*base.Module, int32, int32, int32))(m, l0, l1, v80)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return int32(0)
							} else {
								return v194
							}
						}
					}
				} else {
					v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return int32(0)
					} else {
						return v194
					}
				}
			}
		case 26:
			v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152)+36)))
			if v153 != int32(1) {
				v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v195 = m.ExcPending
				if v195 != 0 {
					return int32(0)
				} else {
					return v194
				}
			} else {
				v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				if v156 != 0 {
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
					v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v159 = int64(*(*int32)(unsafe.Add(mBase, uint32(v158)+40)))
					v161 = F_shm_toc_lookup(m, v157, v159, int32(0))
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
						return int32(0)
					} else {
						v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						if v163 != 0 {
							F_ExecHashTableDetachBatch(m, v163)
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return int32(0)
							} else {
								v166 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								F_ExecHashTableDetach(m, v166)
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return int32(0)
								} else {
									F_FileSetDeleteAll(m, v161+int32(168))
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return int32(0)
									} else {
										v174 = v161 + int32(56)
										v175 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v174)+8)) = v175
										v177 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v174))) = v177
										*(*int64)(unsafe.Add(mBase, uint32(v174)+12)) = v177
										*(*uint8)(unsafe.Add(mBase, uint32(v174)+20)) = uint8(v175)
										F_ConditionVariableInit(m, v161+int32(80))
										mBase = m.M
										v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
											return int32(0)
										} else {
											return v194
										}
									}
								}
							}
						} else {
							F_FileSetDeleteAll(m, v161+int32(168))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								v174 = v161 + int32(56)
								v175 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v174)+8)) = v175
								v177 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v174))) = v177
								*(*int64)(unsafe.Add(mBase, uint32(v174)+12)) = v177
								*(*uint8)(unsafe.Add(mBase, uint32(v174)+20)) = uint8(v175)
								F_ConditionVariableInit(m, v161+int32(80))
								mBase = m.M
								v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									return v194
								}
							}
						}
					}
				} else {
					v194 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v195 = m.ExcPending
					if v195 != 0 {
						return int32(0)
					} else {
						return v194
					}
				}
			}
		}
	}
}
func F_InitializeParallelDSM(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v546 int32
	_ = v546
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v587 int32
	_ = v587
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v697 int32
	_ = v697
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v829 int32
	_ = v829
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v869 int32
	_ = v869
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1162 int32
	_ = v1162
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1261 int32
	_ = v1261
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1304 int32
	_ = v1304
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1315 int32
	_ = v1315
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1342 int64
	_ = v1342
	var v1345 int64
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1398 int32
	_ = v1398
	var v1404 int32
	_ = v1404
	var v1408 int32
	_ = v1408
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1432 int32
	_ = v1432
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1444 int32
	_ = v1444
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1525 int32
	_ = v1525
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1569 int32
	_ = v1569
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1607 int32
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int32
	_ = v1644
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1656 float64
	_ = v1656
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1692 int32
	_ = v1692
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1725 int32
	_ = v1725
	var v1737 int32
	_ = v1737
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1781 int32
	_ = v1781
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1827 int32
	_ = v1827
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1846 int32
	_ = v1846
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1858 int32
	_ = v1858
	var v1864 int32
	_ = v1864
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1876 int32
	_ = v1876
	var v1903 int32
	_ = v1903
	var v1910 int32
	_ = v1910
	var v1945 int32
	_ = v1945
	var v1949 int32
	_ = v1949
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1962 int32
	_ = v1962
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1975 int32
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2021 int32
	_ = v2021
	var v2025 int32
	_ = v2025
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2048 int32
	_ = v2048
	var v2049 int32
	_ = v2049
	var v2050 int64
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2063 int32
	_ = v2063
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2082 int32
	_ = v2082
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int64
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2102 int32
	_ = v2102
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2115 int32
	_ = v2115
	var v2118 int32
	_ = v2118
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2153 int32
	_ = v2153
	var v2156 int64
	_ = v2156
	var v2159 int32
	_ = v2159
	var v2160 int64
	_ = v2160
	var v2163 int32
	_ = v2163
	var v2166 int32
	_ = v2166
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2180 int32
	_ = v2180
	var v2191 int32
	_ = v2191
	var v2209 int32
	_ = v2209
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2224 int32
	_ = v2224
	var v2232 int32
	_ = v2232
	var v2238 int32
	_ = v2238
	var v2256 int32
	_ = v2256
	var v2263 int32
	_ = v2263
	var v2264 int32
	_ = v2264
	var v2267 int32
	_ = v2267
	var v2270 int32
	_ = v2270
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2313 int32
	_ = v2313
	var v2317 int32
	_ = v2317
	var v2350 int32
	_ = v2350
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2355 int32
	_ = v2355
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2359 int32
	_ = v2359
	var v2362 int32
	_ = v2362
	var v2366 int32
	_ = v2366
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2378 int32
	_ = v2378
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2437 int32
	_ = v2437
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2446 int32
	_ = v2446
	var v2448 int32
	_ = v2448
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2455 int32
	_ = v2455
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2524 int32
	_ = v2524
	var v2527 int32
	_ = v2527
	var v2556 int32
	_ = v2556
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2563 int32
	_ = v2563
	var v2598 int32
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2602 int32
	_ = v2602
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2634 int64
	_ = v2634
	var v2636 int32
	_ = v2636
	var v2638 int32
	_ = v2638
	var v2639 int32
	_ = v2639
	var v2642 int32
	_ = v2642
	var v2643 int32
	_ = v2643
	var v2645 int32
	_ = v2645
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2713 int32
	_ = v2713
	var v2718 int32
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2728 int32
	_ = v2728
	var v2730 int32
	_ = v2730
	var v2736 int32
	_ = v2736
	var v2739 int32
	_ = v2739
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2743 int32
	_ = v2743
	var v2745 int32
	_ = v2745
	var v2748 int32
	_ = v2748
	var v2751 int32
	_ = v2751
	var v2756 int32
	_ = v2756
	var v2758 int32
	_ = v2758
	var v2769 int32
	_ = v2769
	var v2795 int32
	_ = v2795
	var v2797 int32
	_ = v2797
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2867 int32
	_ = v2867
	var v2870 int32
	_ = v2870
	var v2871 int32
	_ = v2871
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2876 int32
	_ = v2876
	var v2878 int32
	_ = v2878
	var v2882 int32
	_ = v2882
	var v2884 int32
	_ = v2884
	var v2887 int32
	_ = v2887
	var v2888 int32
	_ = v2888
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2893 int32
	_ = v2893
	var v2896 int32
	_ = v2896
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2910 int32
	_ = v2910
	var v2915 int32
	_ = v2915
	var v2940 int32
	_ = v2940
	var v2943 int32
	_ = v2943
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2949 int32
	_ = v2949
	var v2979 int32
	_ = v2979
	var v2982 int32
	_ = v2982
	var v2984 int32
	_ = v2984
	var v2990 int32
	_ = v2990
	var v2993 int32
	_ = v2993
	var v2994 int32
	_ = v2994
	var v2998 int32
	_ = v2998
	var v3003 int32
	_ = v3003
	var v3028 int32
	_ = v3028
	var v3031 int32
	_ = v3031
	var v3034 int32
	_ = v3034
	var v3035 int32
	_ = v3035
	var v3037 int32
	_ = v3037
	var v3072 int32
	_ = v3072
	var v3075 int32
	_ = v3075
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3080 int32
	_ = v3080
	var v3082 int32
	_ = v3082
	var v3088 int32
	_ = v3088
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3099 int32
	_ = v3099
	var v3102 int32
	_ = v3102
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3109 int32
	_ = v3109
	var v3110 int32
	_ = v3110
	var v3112 int32
	_ = v3112
	var v3114 int32
	_ = v3114
	var v3115 int32
	_ = v3115
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3119 int32
	_ = v3119
	var v3126 int32
	_ = v3126
	var v3156 int32
	_ = v3156
	var v3158 int64
	_ = v3158
	var v3164 int32
	_ = v3164
	var v3174 int32
	_ = v3174
	var v3176 int32
	_ = v3176
	var v3177 int32
	_ = v3177
	var v3178 int32
	_ = v3178
	var v3179 int32
	_ = v3179
	var v3180 int32
	_ = v3180
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3220 int32
	_ = v3220
	var v3223 int32
	_ = v3223
	var v3224 int32
	_ = v3224
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3232 int32
	_ = v3232
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3240 int32
	_ = v3240
	var v3244 int32
	_ = v3244
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3254 int32
	_ = v3254
	var v3258 int32
	_ = v3258
	var v3260 int32
	_ = v3260
	var v3262 int32
	_ = v3262
	var v3265 int32
	_ = v3265
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3272 int32
	_ = v3272
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3282 int32
	_ = v3282
	var v3287 int32
	_ = v3287
	var v3288 int32
	_ = v3288
	var v3289 int32
	_ = v3289
	var v3296 int32
	_ = v3296
	var v3298 int32
	_ = v3298
	var v3299 int32
	_ = v3299
	var v3301 int32
	_ = v3301
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3318 int32
	_ = v3318
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3336 int32
	_ = v3336
	var v3338 int32
	_ = v3338
	var v3340 int32
	_ = v3340
	var v3343 int32
	_ = v3343
	var v3348 int32
	_ = v3348
	var v3349 int32
	_ = v3349
	var v3350 int32
	_ = v3350
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3356 int32
	_ = v3356
	var v3360 int32
	_ = v3360
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3367 int32
	_ = v3367
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3379 int32
	_ = v3379
	var v3387 int32
	_ = v3387
	var v3390 int32
	_ = v3390
	var v3391 int32
	_ = v3391
	var v3423 int32
	_ = v3423
	v2 = int32(0)
	v32 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	goto L3
L3:
	;
	v37 = int32(4520272)
	v38 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v41 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v45 = F_add_size(m, v43, int32(96))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v50 = F_add_size(m, v48, int32(1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v50
	v54 = l0 + int32(36)
	v56 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v56 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v1261 = F_shm_toc_estimate(m, v54)
	mBase = m.M
	v1262 = m.ExcPending
	if v1262 != 0 {
		goto L1
	} else {
		goto L190
	}
L7:
	;
	v68 = l0 + int32(12)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v69 <= int32(0) {
		v1234 = v2
		v1236 = v2
		v1238 = v2
		v1240 = v68
		v1242 = v2
		v1243 = v2
		v1244 = v2
		v1245 = v2
		v1246 = v2
		v1247 = v2
		v1248 = v2
		v1254 = v2
		v1258 = v2
		goto L6
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1234 = v2
	v1236 = v2
	v1238 = v2
	v1240 = l0 + int32(12)
	v1242 = v2
	v1243 = v2
	v1244 = v2
	v1245 = v2
	v1246 = v2
	v1247 = v2
	v1248 = v2
	v1254 = v2
	v1258 = v2
	goto L6
L9:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v58 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v60 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v72 = m.G0
	v74 = v72 - int32(16)
	m.G0 = v74
	v77 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	if v78 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	m.G0 = v74 + int32(16)
	if v362 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L14:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v362 = v79
	goto L13
L15:
	;
	goto L16
L16:
	;
	v80 = int32(4520272)
	v81 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v84 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v84
	v88 = F_add_size(m, int32(0), int32(1))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v92 = F_add_size(m, int32(0), int32(196608))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v95 = F_add_size(m, v88, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+12)) = v95
	v99 = F_add_size(m, v92, int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v74)+8)) = v99
	v104 = F_shm_toc_estimate(m, v74+int32(8))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v107 = F_dsm_create(m, v104, int32(1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v107 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v81
	v362 = int32(0)
	goto L13
L24:
	;
	goto L25
L25:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v107)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v115)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v115))) = int64(2880502729)
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v104 & int32(-32)
	goto L26
L26:
	;
	v125 = F_shm_toc_allocate(m, v115, int32(196608))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v129 = F_dsa_create_in_place_ext(m, v125, int32(196608), int32(72), v107)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_shm_toc_insert(m, v115, int64(-65535), v125)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v135 = F_shm_toc_allocate(m, v115, int32(12))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v137 = m.G0
	v139 = v137 - int32(16)
	m.G0 = v139
	v141 = int32(4520272)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v145 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v145
	v148 = F_dshash_create(m, v129, int32(1766220), v129)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v152 = F_dshash_create(m, v129, int32(1766244), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v142
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v152)+32))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = v160
	v163 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v163
	if int32(0) < v163 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_shm_toc_insert(m, v115, int64(-65534), v135)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L61
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L58
	}
L37:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v174 = v163
	v188 = v168
	v189 = v2
	goto L40
L38:
	;
	goto L39
L39:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v291)+12)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v135
	F_on_dsm_detach(m, v107, int32(1623), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L57
	}
L40:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(int32(4))%32))+8))
	if v203 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v210 = F_dsa_allocate_extended(m, v129, v204*int32(116)+int32(20), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v253 = v174
	v254 = v188
	goto L44
L44:
	;
	v257 = v189 + int32(1)
	if v257 < v253 {
		v174 = v253
		v188 = v254
		v189 = v257
		goto L40
	} else {
		goto L56
	}
L45:
	;
	v212 = F_dsa_get_address(m, v129, v210)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_TupleDescCopy(m, v212, v203)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212)+8)) = v189
	v221 = F_dshash_find_or_insert(m, v152, v203+int32(8), v139+int32(7))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+7)))
	if v223 == int32(1) {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v221)+4)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v221))) = v226
	F_dshash_release_lock(m, v152, v221)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v203
	v232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+12)) = uint8(v232)
	v238 = F_dshash_find_or_insert(m, v148, v139+int32(8), v139+int32(7))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+7)))
	if v240 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v210
	v244 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+4)) = uint8(v244)
	goto L54
L53:
	;
	goto L54
L54:
	;
	F_dshash_release_lock(m, v148, v238)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v251 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	v253 = v251
	v254 = v249
	goto L44
L56:
	;
	goto L41
L57:
	;
	m.G0 = v139 + int32(16)
	goto L35
L58:
	;
	F_errmsg_internal(m, int32(423187), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(500408), int32(2253), int32(100752))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_dsm_pin_mapping(m, v107)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_dsa_pin_mapping(m, v129)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v322 = int32(4414732)
	v323 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v107
	v326 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	*(*int32)(unsafe.Add(mBase, uint32(v326)+4)) = v129
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v81
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v362 = v330
	goto L13
L64:
	;
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v368
	v1234 = v2
	v1236 = v2
	v1238 = v2
	v1240 = v68
	v1242 = v2
	v1243 = v2
	v1244 = v2
	v1245 = v2
	v1246 = v2
	v1247 = v368
	v1248 = v2
	v1254 = v2
	v1258 = v2
	goto L6
L65:
	;
	goto L66
L66:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v371 <= int32(0) {
		v1234 = v2
		v1236 = v2
		v1238 = v2
		v1240 = v68
		v1242 = v2
		v1243 = v2
		v1244 = v2
		v1245 = v2
		v1246 = v2
		v1247 = v362
		v1248 = v2
		v1254 = v2
		v1258 = v2
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v374 = int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	if v376 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v381 = v376
	v383 = v374
	goto L71
L69:
	;
	v422 = v374
	goto L70
L70:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v452 = F_add_size(m, v447, (v422+int32(31))&int32(-32))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L75
	}
L71:
	;
	v410 = F_strlen(m, v381+int32(24))
	mBase = m.M
	v413 = F_add_size(m, v383, v410+int32(1))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	v422 = v413
	goto L70
L73:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v415 != 0 {
		v381 = v415
		v383 = v413
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v452
	v455 = m.G0
	v457 = v455 - int32(16)
	m.G0 = v457
	v459 = int32(4)
	v461 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v461 == int32(0) {
		v697 = v459
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v766 = F_add_size(m, v761, (v697+int32(31))&int32(-32))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L122
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L119
	}
L78:
	;
	m.G0 = v457 + int32(16)
	goto L76
L79:
	;
	if v461 == int32(4518300) {
		v697 = v459
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v478 = v461
	v482 = v459
	goto L81
L81:
	;
	v497 = int32(0)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v478-int32(60))))
	if base.Ui32(v500) < base.Ui32(int32(2)) {
		v653 = v497
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v697 = v676
	goto L78
L83:
	;
	v676 = F_add_size(m, v482, v653)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L117
	}
L84:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v478-int32(32))))
	if v505 == int32(0) {
		v653 = v497
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v509 = v478 + int32(-64)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	v511 = F_strlen(m, v510)
	mBase = m.M
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v478-int32(40))))
	switch v514 {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	case 3:
		goto L88
	case 4:
		goto L87
	default:
		v587 = v497
		goto L86
	}
L86:
	;
	v610 = int32(1)
	v614 = F_add_size(m, v511+v610, v587+v610)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L104
	}
L87:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v478)+28))
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v533)))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v478)+36))
	if v535 == int32(0) {
		goto L77
	} else {
		goto L96
	}
L88:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v478)+28))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	if v529 == int32(0) {
		v587 = v497
		goto L86
	} else {
		goto L95
	}
L89:
	;
	v587 = int32(25)
	goto L86
L90:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v478)+28))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v521 = v519 >> (uint(int32(31)) % 32)
	if v519^v521-v521 < int32(1000) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v587 = int32(5)
	goto L86
L92:
	;
	v526 = int32(4)
	goto L94
L93:
	;
	v526 = int32(11)
	goto L94
L94:
	;
	v587 = v526
	goto L86
L95:
	;
	v532 = F_strlen(m, v529)
	mBase = m.M
	v587 = v532
	goto L86
L96:
	;
	v546 = v535
	goto L97
L97:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v546)))
	if v569 == int32(0) {
		goto L77
	} else {
		goto L99
	}
L98:
	;
	v578 = F_strlen(m, v569)
	mBase = m.M
	v587 = v578
	goto L86
L99:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v546)+4))
	if v534 != v572 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v575 = v546 + int32(12)
	if v575 == int32(0) {
		goto L77
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	goto L98
L103:
	;
	v546 = v575
	goto L97
L104:
	;
	v617 = v478 + int32(20)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	if v618 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v619 = F_strlen(m, v618)
	mBase = m.M
	v620 = F_add_size(m, v614, v619)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	v622 = v614
	goto L107
L107:
	;
	v624 = F_add_size(m, v622, int32(1))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	v622 = v620
	goto L107
L109:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	if v626 == int32(0) {
		v635 = v624
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v637 = F_add_size(m, v635, int32(4))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L1
	} else {
		goto L114
	}
L111:
	;
	v629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626))))
	if v629 == int32(0) {
		v635 = v624
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v633 = F_add_size(m, v624, int32(4))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v635 = v633
	goto L110
L114:
	;
	v640 = F_add_size(m, v637, int32(4))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v643 = F_add_size(m, v640, int32(4))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v653 = v643
	goto L83
L117:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	if v678 != int32(4518300) {
		v478 = v678
		v482 = v676
		goto L81
	} else {
		goto L118
	}
L118:
	;
	goto L82
L119:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v509)))
	*(*int32)(unsafe.Add(mBase, uint32(v457)+4)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v457))) = v534
	F_errmsg_internal(m, int32(182102), v457)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(500971), int32(3036), int32(345741))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v766
	v772 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v773 = F_mul_size(m, int32(8), v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	v775 = F_add_size(m, int32(4), v773)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v782 = F_add_size(m, v777, (v775+int32(31))&int32(-32))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v782
	v786 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	if int32(2) <= v786 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v789 = F_EstimateSnapshotSpace(m, v32)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L1
	} else {
		goto L129
	}
L127:
	;
	v799 = v2
	goto L128
L128:
	;
	v800 = F_EstimateSnapshotSpace(m, v36)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L131
	}
L129:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v796 = F_add_size(m, v791, (v789+int32(31))&int32(-32))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v796
	v799 = v789
	goto L128
L131:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v807 = F_add_size(m, v802, (v800+int32(31))&int32(-32))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v807
	v810 = int32(0)
	v812 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	if v812 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v814 = v812
	v829 = v810
	goto L136
L134:
	;
	v869 = v810
	goto L135
L135:
	;
	v886 = F_mul_size(m, int32(4), v869)
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L1
	} else {
		goto L144
	}
L136:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v814)))
	if v844 != 0 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v869 = v850
	goto L135
L138:
	;
	v846 = F_add_size(m, v829, int32(1))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L141
	}
L139:
	;
	v848 = v829
	goto L140
L140:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v814)+52))
	v850 = F_add_size(m, v848, v849)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L142
	}
L141:
	;
	v848 = v846
	goto L140
L142:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v814)+80))
	if v852 != 0 {
		v814 = v852
		v829 = v850
		goto L136
	} else {
		goto L143
	}
L143:
	;
	goto L137
L144:
	;
	v888 = F_add_size(m, int32(32), v886)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v895 = F_add_size(m, v890, (v888+int32(31))&int32(-32))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v895
	v899 = F_add_size(m, v895, int32(32))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v899
	v903 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v903 != 0 {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v903)))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v905)+4))
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v905)+412))
	if v907 != 0 {
		goto L152
	} else {
		goto L153
	}
L149:
	;
	v976 = int32(1)
	goto L150
L150:
	;
	v978 = F_mul_size(m, v976, int32(12))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L155
	}
L151:
	;
	v976 = v972 + int32(1)
	goto L150
L152:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v905)+376))
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v905)+364))
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v905)+352))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v905)+340))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v905)+328))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v905)+316))
	v914 = *(*int32)(unsafe.Add(mBase, uint32(v905)+304))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v905)+292))
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v905)+280))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v905)+268))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v905)+256))
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v905)+244))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v905)+232))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v905)+220))
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v905)+208))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v905)+196))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v905)+184))
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v905)+172))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v905)+160))
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v905)+148))
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v905)+136))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v905)+124))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v905)+112))
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v905)+100))
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v905)+88))
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v905)+76))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v905-int32(-64))))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v905)+52))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v905)+40))
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v905)+28))
	v940 = *(*int32)(unsafe.Add(mBase, uint32(v905)+16))
	v972 = v908 + (v909 + (v910 + (v911 + (v912 + (v913 + (v914 + (v915 + (v916 + (v917 + (v918 + (v919 + (v920 + (v921 + (v922 + (v923 + (v924 + (v925 + (v926 + (v927 + (v928 + (v929 + (v930 + (v931 + (v932 + (v933 + (v936 + (v937 + (v938 + (v939 + (v940 + v906))))))))))))))))))))))))))))))
	goto L154
L153:
	;
	v972 = v906
	goto L154
L154:
	;
	goto L151
L155:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v985 = F_add_size(m, v980, (v978+int32(31))&int32(-32))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v985
	v990 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	if v990 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v990)+4))
	v993 = v991
	goto L159
L158:
	;
	v993 = int32(0)
	goto L159
L159:
	;
	v994 = F_mul_size(m, int32(4), v993)
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1003 = F_add_size(m, v998, (v994+int32(43))&int32(-32))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1003
	v1009 = F_add_size(m, v1003, int32(1056))
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1009
	v1014 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v1014 != 0 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v1014)))
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+4))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+412))
	if v1018 != 0 {
		goto L167
	} else {
		goto L168
	}
L164:
	;
	v1084 = int32(0)
	goto L165
L165:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if v1086 != 0 {
		goto L170
	} else {
		goto L171
	}
L166:
	;
	v1084 = v1083
	goto L165
L167:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+376))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+364))
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+352))
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+340))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+328))
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+316))
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+304))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+292))
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+280))
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+268))
	v1029 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+256))
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+244))
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+232))
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+220))
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+208))
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+196))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+184))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+172))
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+160))
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+148))
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+136))
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+124))
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+112))
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+100))
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+88))
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+76))
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(v1016-int32(-64))))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+52))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+40))
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+28))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1016)+16))
	v1083 = v1019 + (v1020 + (v1021 + (v1022 + (v1023 + (v1024 + (v1025 + (v1026 + (v1027 + (v1028 + (v1029 + (v1030 + (v1031 + (v1032 + (v1033 + (v1034 + (v1035 + (v1036 + (v1037 + (v1038 + (v1039 + (v1040 + (v1041 + (v1042 + (v1043 + (v1044 + (v1047 + (v1048 + (v1049 + (v1050 + (v1051 + v1017))))))))))))))))))))))))))))))
	goto L169
L168:
	;
	v1083 = v1017
	goto L169
L169:
	;
	goto L166
L170:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1086)))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+4))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+412))
	if v1090 != 0 {
		goto L174
	} else {
		goto L175
	}
L171:
	;
	v1157 = v1084
	goto L172
L172:
	;
	v1159 = v1157 << (uint(int32(2)) % 32)
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1167 = F_add_size(m, v1162, (v1159+int32(39))&int32(-32))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L177
	}
L173:
	;
	v1157 = v1155 + v1084
	goto L172
L174:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+376))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+364))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+352))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+340))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+328))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+316))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+304))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+292))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+280))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+268))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+256))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+244))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+232))
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+220))
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+208))
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+196))
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+184))
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+172))
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+160))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+148))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+136))
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+124))
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+112))
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+100))
	v1115 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+88))
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+76))
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1088-int32(-64))))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+52))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+40))
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+28))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1088)+16))
	v1155 = v1091 + (v1092 + (v1093 + (v1094 + (v1095 + (v1096 + (v1097 + (v1098 + (v1099 + (v1100 + (v1101 + (v1102 + (v1103 + (v1104 + (v1105 + (v1106 + (v1107 + (v1108 + (v1109 + (v1110 + (v1111 + (v1112 + (v1113 + (v1114 + (v1115 + (v1116 + (v1119 + (v1120 + (v1121 + (v1122 + (v1123 + v1089))))))))))))))))))))))))))))))
	goto L176
L175:
	;
	v1155 = v1089
	goto L176
L176:
	;
	goto L173
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1167
	v1172 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L1
	} else {
		goto L178
	}
L178:
	;
	v1175 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	if v1175 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v1176 = F_strlen(m, v1175)
	mBase = m.M
	v1179 = F_add_size(m, v1172, v1176+int32(1))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	v1181 = v1172
	goto L181
L181:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1187 = F_add_size(m, v1182, (v1181+int32(31))&int32(-32))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L183
	}
L182:
	;
	v1181 = v1179
	goto L181
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1187
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1192 = F_add_size(m, v1190, int32(12))
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1192
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1198 = F_mul_size(m, int32(16384), v1197)
	mBase = m.M
	v1199 = m.ExcPending
	if v1199 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v1204 = F_add_size(m, v1195, (v1198+int32(31))&int32(-32))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1204
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1209 = F_add_size(m, v1207, int32(1))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1209
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1214 = F_strlen(m, v1213)
	mBase = m.M
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1216 = F_strlen(m, v1215)
	mBase = m.M
	v1222 = F_add_size(m, v1212, (v1214+v1216+int32(33))&int32(-32))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1222
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1227 = F_add_size(m, v1225, int32(1))
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1227
	v1234 = v697
	v1236 = v422
	v1238 = v978
	v1240 = v68
	v1242 = v994 + int32(12)
	v1243 = v800
	v1244 = v888
	v1245 = v1159 + int32(8)
	v1246 = int32(1048)
	v1247 = v362
	v1248 = v1181
	v1254 = v775
	v1258 = v799
	goto L6
L190:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1240)))
	if v1263 <= int32(0) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v1271 != 0 {
		goto L197
	} else {
		goto L198
	}
L192:
	;
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1271 = v1266
	goto L191
L193:
	;
	goto L194
L194:
	;
	v1268 = F_dsm_create(m, v1261, int32(1))
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1268
	v1271 = v1268
	goto L191
L196:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1280)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1280))) = int64(1346862204)
	*(*int32)(unsafe.Add(mBase, uint32(v1280)+12)) = v1261 & int32(-32)
	goto L201
L197:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v1271)+24))
	v1280 = v1272
	goto L196
L198:
	;
	goto L199
L199:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1276 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1277 = F_MemoryContextAlloc(m, v1276, v1261)
	mBase = m.M
	v1278 = m.ExcPending
	if v1278 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v1277
	v1280 = v1277
	goto L196
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1280
	v1292 = F_shm_toc_allocate(m, v1280, int32(80))
	mBase = m.M
	v1293 = m.ExcPending
	if v1293 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292))) = v1295
	v1298 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+4)) = v1298
	v1301 = *(*int32)(unsafe.Add(mBase, _consts[130]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+8)) = v1301
	v1304 = *(*int32)(unsafe.Add(mBase, _consts[131]))
	v1307 = int32(*(*uint8)(unsafe.Add(mBase, _consts[132])))
	if v1307 != 0 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+12)) = v1308
	v1315 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292+int32(16)))) = v1315
	v1318 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292+int32(28)))) = v1318
	goto L207
L204:
	;
	v1308 = v1304
	goto L206
L205:
	;
	v1308 = int32(0)
	goto L206
L206:
	;
	goto L203
L207:
	;
	v1321 = int32(*(*uint8)(unsafe.Add(mBase, _consts[133])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1292)+32)) = uint8(v1321)
	v1324 = int32(*(*uint8)(unsafe.Add(mBase, _consts[134])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1292)+33)) = uint8(v1324)
	v1327 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+20)) = v1327
	v1330 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+24)) = v1330
	v1333 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+36)) = v1333
	v1336 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+40)) = v1336
	v1339 = *(*int32)(unsafe.Add(mBase, _consts[109]))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+44)) = v1339
	v1342 = *(*int64)(unsafe.Add(mBase, _consts[139]))
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+48)) = v1342
	v1345 = *(*int64)(unsafe.Add(mBase, _consts[140]))
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+56)) = v1345
	v1348 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+68)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+64)) = v1348
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1354, int64(-65535), v1292)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v1358 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1362 = F_shm_toc_allocate(m, v1361, v1236)
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L212
	}
L210:
	;
	v3423 = v1358
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3423
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v38
	return
L212:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	if v1365 != 0 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v1368 = v1365
	v1372 = v1236
	v1373 = v1362
	goto L216
L214:
	;
	v1525 = v1362
	goto L215
L215:
	;
	v1549 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1525))) = uint8(v1549)
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1551, int64(-65533), v1362)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L1
	} else {
		goto L251
	}
L216:
	;
	v1398 = v1368 + int32(24)
	if v1372 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L217:
	;
	v1525 = v1515
	goto L215
L218:
	;
	v1514 = v1510 + (v1507 - v1373) + int32(1)
	v1515 = v1373 + v1514
	v1517 = *(*int32)(unsafe.Add(mBase, uint32(v1368)))
	if v1517 != 0 {
		v1368 = v1517
		v1372 = v1372 - v1514
		v1373 = v1515
		goto L216
	} else {
		goto L250
	}
L219:
	;
	v1510 = F_strlen(m, v1506)
	mBase = m.M
	goto L218
L220:
	;
	v1506 = v1398
	v1507 = v1373
	goto L219
L221:
	;
	goto L222
L222:
	;
	v1404 = v1372 - int32(1)
	if (v1373^v1398)&int32(3) != 0 {
		goto L226
	} else {
		goto L227
	}
L223:
	;
	v1503 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1500))) = uint8(v1503)
	v1506 = v1499
	v1507 = v1500
	goto L219
L224:
	;
	v1484 = v1479
	v1485 = v1480
	v1486 = v1481
	goto L246
L225:
	;
	if v1474 == int32(0) {
		v1499 = v1472
		v1500 = v1473
		goto L223
	} else {
		goto L245
	}
L226:
	;
	v1472 = v1398
	v1473 = v1373
	v1474 = v1404
	goto L225
L227:
	;
	goto L228
L228:
	;
	v1408 = int32(0)
	if v1398&int32(3) == v1408 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	if v1441 == int32(0) {
		v1499 = v1438
		v1500 = v1439
		goto L223
	} else {
		goto L238
	}
L230:
	;
	v1438 = v1398
	v1439 = v1373
	v1440 = v1404
	v1441 = base.B2i32(v1404 != v1408)
	goto L229
L231:
	;
	if v1404 == int32(0) {
		goto L230
	} else {
		goto L232
	}
L232:
	;
	v1417 = v1398
	v1418 = v1373
	v1419 = v1404
	goto L233
L233:
	;
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1417))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1418))) = uint8(v1421)
	if v1421 == int32(0) {
		v1479 = v1417
		v1480 = v1418
		v1481 = v1419
		goto L224
	} else {
		goto L235
	}
L234:
	;
	v1438 = v1432
	v1439 = v1426
	v1440 = v1428
	v1441 = v1430
	goto L229
L235:
	;
	v1425 = int32(1)
	v1426 = v1418 + v1425
	v1428 = v1419 - v1425
	v1429 = int32(0)
	v1430 = base.B2i32(v1428 != v1429)
	v1432 = v1417 + v1425
	if v1432&int32(3) == v1429 {
		v1438 = v1432
		v1439 = v1426
		v1440 = v1428
		v1441 = v1430
		goto L229
	} else {
		goto L236
	}
L236:
	;
	if v1428 != 0 {
		v1417 = v1432
		v1418 = v1426
		v1419 = v1428
		goto L233
	} else {
		goto L237
	}
L237:
	;
	goto L234
L238:
	;
	v1444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1438))))
	if v1444 == int32(0) {
		v1472 = v1438
		v1473 = v1439
		v1474 = v1440
		goto L225
	} else {
		goto L239
	}
L239:
	;
	if base.Ui32(v1440) < base.Ui32(int32(4)) {
		v1472 = v1438
		v1473 = v1439
		v1474 = v1440
		goto L225
	} else {
		goto L240
	}
L240:
	;
	v1450 = v1438
	v1451 = v1439
	v1452 = v1440
	goto L241
L241:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, uint32(v1450)))
	v1458 = int32(-2139062144)
	if (int32(16843008)-v1455|v1455)&v1458 != v1458 {
		v1479 = v1450
		v1480 = v1451
		v1481 = v1452
		goto L224
	} else {
		goto L243
	}
L242:
	;
	v1472 = v1466
	v1473 = v1464
	v1474 = v1468
	goto L225
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1451))) = v1455
	v1463 = int32(4)
	v1464 = v1451 + v1463
	v1466 = v1450 + v1463
	v1468 = v1452 - v1463
	if base.Ui32(int32(3)) < base.Ui32(v1468) {
		v1450 = v1466
		v1451 = v1464
		v1452 = v1468
		goto L241
	} else {
		goto L244
	}
L244:
	;
	goto L242
L245:
	;
	v1479 = v1472
	v1480 = v1473
	v1481 = v1474
	goto L224
L246:
	;
	v1488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1484))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1485))) = uint8(v1488)
	if v1488 == int32(0) {
		v1499 = v1484
		v1500 = v1485
		goto L223
	} else {
		goto L248
	}
L247:
	;
	v1499 = v1495
	v1500 = v1493
	goto L223
L248:
	;
	v1492 = int32(1)
	v1493 = v1485 + v1492
	v1495 = v1484 + v1492
	v1497 = v1486 - v1492
	if v1497 != 0 {
		v1484 = v1495
		v1485 = v1493
		v1486 = v1497
		goto L246
	} else {
		goto L249
	}
L249:
	;
	goto L247
L250:
	;
	goto L217
L251:
	;
	v1555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1556 = F_shm_toc_allocate(m, v1555, v1234)
	mBase = m.M
	v1557 = m.ExcPending
	if v1557 != 0 {
		goto L1
	} else {
		goto L252
	}
L252:
	;
	v1558 = m.G0
	v1560 = v1558 - int32(112)
	m.G0 = v1560
	v1562 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+108)) = v1556 + v1562
	v1566 = v1234 - v1562
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+104)) = v1566
	v1569 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	if v1569 == int32(0) {
		v1910 = v1566
		goto L258
	} else {
		goto L259
	}
L253:
	;
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1994, int64(-65532), v1556)
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L1
	} else {
		goto L323
	}
L254:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L1
	} else {
		goto L320
	}
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L1
	} else {
		goto L317
	}
L256:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L1
	} else {
		goto L314
	}
L257:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L1
	} else {
		goto L311
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1556))) = v1566 - v1910
	m.G0 = v1560 + int32(112)
	goto L253
L259:
	;
	if v1569 == int32(4518300) {
		v1910 = v1566
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v1578 = v1566
	v1583 = v1569
	goto L261
L261:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1583-int32(60))))
	if base.Ui32(v1607) < base.Ui32(int32(2)) {
		v1876 = v1578
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v1910 = v1876
	goto L258
L263:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+4))
	if v1903 != int32(4518300) {
		v1578 = v1876
		v1583 = v1903
		goto L261
	} else {
		goto L310
	}
L264:
	;
	v1611 = v1583 - int32(32)
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1611)))
	if v1612 == int32(0) {
		v1876 = v1578
		goto L263
	} else {
		goto L265
	}
L265:
	;
	v1616 = v1583 + int32(-64)
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1616)))
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+96)) = v1617
	F_do_serialize(m, v1560+int32(108), v1560+int32(104), int32(206576), v1560+int32(96))
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1583-int32(40))))
	switch v1630 {
	case 0:
		goto L275
	case 1:
		goto L274
	case 2:
		goto L273
	case 3:
		goto L272
	case 4:
		goto L271
	default:
		goto L270
	}
L267:
	;
	if base.Ui32(v1846) <= base.Ui32(int32(3)) {
		goto L256
	} else {
		goto L307
	}
L268:
	;
	v1833 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+104))
	if base.Ui32(v1833) <= base.Ui32(int32(3)) {
		goto L257
	} else {
		goto L306
	}
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L1
	} else {
		goto L303
	}
L270:
	;
	v1770 = v1583 + int32(20)
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1770)))
	if v1771 != 0 {
		goto L295
	} else {
		goto L296
	}
L271:
	;
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+28))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1683)))
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+36))
	if v1685 == int32(0) {
		goto L269
	} else {
		goto L286
	}
L272:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+28))
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1669)))
	if v1670 != 0 {
		goto L282
	} else {
		goto L283
	}
L273:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+28))
	v1656 = *(*float64)(unsafe.Add(mBase, uint32(v1655)))
	*(*float64)(unsafe.Add(mBase, uint32(v1560)+40)) = v1656
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+32)) = int32(17)
	F_do_serialize(m, v1560+int32(108), v1560+int32(104), int32(420890), v1560+int32(32))
	mBase = m.M
	v1668 = m.ExcPending
	if v1668 != 0 {
		goto L1
	} else {
		goto L281
	}
L274:
	;
	v1643 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+28))
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v1643)))
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+16)) = v1644
	F_do_serialize(m, v1560+int32(108), v1560+int32(104), int32(489370), v1560+int32(16))
	mBase = m.M
	v1654 = m.ExcPending
	if v1654 != 0 {
		goto L1
	} else {
		goto L280
	}
L275:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+28))
	v1638 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1637))))
	if v1638 != 0 {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1639 = int32(345210)
	goto L278
L277:
	;
	v1639 = int32(362226)
	goto L278
L278:
	;
	F_do_serialize(m, v1560+int32(108), v1560+int32(104), v1639, int32(0))
	mBase = m.M
	v1642 = m.ExcPending
	if v1642 != 0 {
		goto L1
	} else {
		goto L279
	}
L279:
	;
	goto L270
L280:
	;
	goto L270
L281:
	;
	goto L270
L282:
	;
	v1672 = v1670
	goto L284
L283:
	;
	v1672 = int32(758841)
	goto L284
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+48)) = v1672
	F_do_serialize(m, v1560+int32(108), v1560+int32(104), int32(206576), v1560+int32(48))
	mBase = m.M
	v1682 = m.ExcPending
	if v1682 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	goto L270
L286:
	;
	v1692 = v1685
	goto L287
L287:
	;
	v1719 = *(*int32)(unsafe.Add(mBase, uint32(v1692)))
	if v1719 == int32(0) {
		goto L269
	} else {
		goto L289
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+80)) = v1719
	F_do_serialize(m, v1560+int32(108), v1560+int32(104), int32(206576), v1560+int32(80))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L1
	} else {
		goto L294
	}
L289:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1692)+4))
	if v1684 != v1722 {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1725 = v1692 + int32(12)
	if v1725 == int32(0) {
		goto L269
	} else {
		goto L293
	}
L291:
	;
	goto L292
L292:
	;
	goto L288
L293:
	;
	v1692 = v1725
	goto L287
L294:
	;
	goto L270
L295:
	;
	v1773 = v1771
	goto L297
L296:
	;
	v1773 = int32(758841)
	goto L297
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1560))) = v1773
	F_do_serialize(m, v1560+int32(108), v1560+int32(104), int32(206576), v1560)
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L1
	} else {
		goto L298
	}
L298:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1770)))
	if v1782 != 0 {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1782))))
	if v1783 != 0 {
		goto L268
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+104))
	v1846 = v1784
	goto L267
L302:
	;
	goto L301
L303:
	;
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1616)))
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+68)) = v1820
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+64)) = v1684
	F_errmsg_internal(m, int32(182102), v1560-int32(-64))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L1
	} else {
		goto L304
	}
L304:
	;
	F_errfinish(m, int32(500971), int32(3036), int32(345741))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L306:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+108))
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1583)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1836))) = v1837
	v1839 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+108)) = v1836 + v1839
	v1846 = v1833 - v1839
	goto L267
L307:
	;
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1560)+108))
	v1850 = *(*int32)(unsafe.Add(mBase, uint32(v1611)))
	*(*int32)(unsafe.Add(mBase, uint32(v1849))) = v1850
	v1853 = v1846 & int32(-4)
	if v1853 == int32(4) {
		goto L255
	} else {
		goto L308
	}
L308:
	;
	v1858 = *(*int32)(unsafe.Add(mBase, uint32(v1583-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1849)+4)) = v1858
	if v1853 == int32(8) {
		goto L254
	} else {
		goto L309
	}
L309:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(v1583-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v1849)+8)) = v1864
	v1866 = int32(12)
	v1867 = v1846 - v1866
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+104)) = v1867
	*(*int32)(unsafe.Add(mBase, uint32(v1560)+108)) = v1849 + v1866
	v1876 = v1867
	goto L263
L310:
	;
	goto L262
L311:
	;
	F_errmsg_internal(m, int32(354401), int32(0))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	F_errfinish(m, int32(500971), int32(6020), int32(17783))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L1
	} else {
		goto L313
	}
L313:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L314:
	;
	F_errmsg_internal(m, int32(354401), int32(0))
	mBase = m.M
	v1962 = m.ExcPending
	if v1962 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(500971), int32(6020), int32(17783))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L1
	} else {
		goto L316
	}
L316:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L317:
	;
	F_errmsg_internal(m, int32(354401), int32(0))
	mBase = m.M
	v1975 = m.ExcPending
	if v1975 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	F_errfinish(m, int32(500971), int32(6020), int32(17783))
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L1
	} else {
		goto L319
	}
L319:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L320:
	;
	F_errmsg_internal(m, int32(354401), int32(0))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L1
	} else {
		goto L321
	}
L321:
	;
	F_errfinish(m, int32(500971), int32(6020), int32(17783))
	mBase = m.M
	v1993 = m.ExcPending
	if v1993 != 0 {
		goto L1
	} else {
		goto L322
	}
L322:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L323:
	;
	v1998 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1999 = F_shm_toc_allocate(m, v1998, v1254)
	mBase = m.M
	v2000 = m.ExcPending
	if v2000 != 0 {
		goto L1
	} else {
		goto L324
	}
L324:
	;
	v2002 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	*(*int32)(unsafe.Add(mBase, uint32(v1999))) = v2002
	v2005 = v1999 + int32(4)
	v2007 = v2002 << (uint(int32(3)) % 32)
	v2008 = v2005 + v2007
	if base.Ui32(v2008) < base.Ui32(v1999) {
		goto L326
	} else {
		goto L327
	}
L325:
	;
	v2031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2031, int64(-65531), v1999)
	mBase = m.M
	v2034 = m.ExcPending
	if v2034 != 0 {
		goto L1
	} else {
		goto L339
	}
L326:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2021 = m.ExcPending
	if v2021 != 0 {
		goto L1
	} else {
		goto L336
	}
L327:
	;
	if base.Ui32(v1999+v1254) < base.Ui32(v2008) {
		goto L326
	} else {
		goto L328
	}
L328:
	;
	if int32(0) < v2002 {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	v2015 = *(*int32)(unsafe.Add(mBase, _consts[142]))
	if v2007 != 0 {
		goto L333
	} else {
		goto L334
	}
L330:
	;
	goto L331
L331:
	;
	goto L325
L332:
	;
	goto L331
L333:
	;
	v2016 = F__emscripten_memcpy_bulkmem(m, v2005, v2015, v2007)
	mBase = m.M
	goto L335
L334:
	;
	goto L335
L335:
	;
	goto L332
L336:
	;
	F_errmsg_internal(m, int32(354312), int32(0))
	mBase = m.M
	v2025 = m.ExcPending
	if v2025 != 0 {
		goto L1
	} else {
		goto L337
	}
L337:
	;
	F_errfinish(m, int32(500931), int32(327), int32(355078))
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L339:
	;
	v2036 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	if int32(2) <= v2036 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2040 = F_shm_toc_allocate(m, v2039, v1258)
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L1
	} else {
		goto L343
	}
L341:
	;
	goto L342
L342:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2089 = F_shm_toc_allocate(m, v2088, v1243)
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L1
	} else {
		goto L358
	}
L343:
	;
	v2048 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v2050 = *(*int64)(unsafe.Add(mBase, uint32(v32)+4))
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v2052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+29)))
	v2053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2040)+16)) = uint8(v2053)
	*(*uint8)(unsafe.Add(mBase, uint32(v2040)+17)) = uint8(v2052)
	*(*int32)(unsafe.Add(mBase, uint32(v2040)+20)) = v2051
	*(*int64)(unsafe.Add(mBase, uint32(v2040))) = v2050
	*(*int32)(unsafe.Add(mBase, uint32(v2040)+8)) = v2049
	if v2052 != 0 {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	v2083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2083, int64(-65530), v2040)
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L1
	} else {
		goto L357
	}
L345:
	;
	v2060 = v2048
	goto L347
L346:
	;
	v2060 = int32(0)
	goto L347
L347:
	;
	if v2053 != 0 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v2061 = v2060
	goto L350
L349:
	;
	v2061 = v2048
	goto L350
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2040)+12)) = v2061
	v2063 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v2063 != 0 {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v2069 = F___memcpy(m, v2040+int32(24), v2066, v2063<<(uint(int32(2))%32))
	mBase = m.M
	goto L353
L352:
	;
	goto L353
L353:
	;
	if int32(0) < v2061 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v2072 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v2073 = int32(2)
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v2082 = F___memcpy(m, v2040+v2072<<(uint(v2073)%32)+int32(24), v2078, v2079<<(uint(v2073)%32))
	mBase = m.M
	goto L356
L355:
	;
	goto L356
L356:
	;
	goto L344
L357:
	;
	goto L342
L358:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v2099 = *(*int64)(unsafe.Add(mBase, uint32(v36)+4))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v2101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+29)))
	v2102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2089)+16)) = uint8(v2102)
	*(*uint8)(unsafe.Add(mBase, uint32(v2089)+17)) = uint8(v2101)
	*(*int32)(unsafe.Add(mBase, uint32(v2089)+20)) = v2100
	*(*int64)(unsafe.Add(mBase, uint32(v2089))) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v2089)+8)) = v2098
	if v2101 != 0 {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2132, int64(-65529), v2089)
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L1
	} else {
		goto L372
	}
L360:
	;
	v2109 = v2097
	goto L362
L361:
	;
	v2109 = int32(0)
	goto L362
L362:
	;
	if v2102 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v2110 = v2109
	goto L365
L364:
	;
	v2110 = v2097
	goto L365
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2089)+12)) = v2110
	v2112 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v2112 != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v2115 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v2118 = F___memcpy(m, v2089+int32(24), v2115, v2112<<(uint(int32(2))%32))
	mBase = m.M
	goto L368
L367:
	;
	goto L368
L368:
	;
	if int32(0) < v2110 {
		goto L369
	} else {
		goto L370
	}
L369:
	;
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v2122 = int32(2)
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v2131 = F___memcpy(m, v2089+v2121<<(uint(v2122)%32)+int32(24), v2127, v2128<<(uint(v2122)%32))
	mBase = m.M
	goto L371
L370:
	;
	goto L371
L371:
	;
	goto L359
L372:
	;
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2138 = F_shm_toc_allocate(m, v2136, int32(4))
	mBase = m.M
	v2139 = m.ExcPending
	if v2139 != 0 {
		goto L1
	} else {
		goto L373
	}
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2138))) = v1247
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2141, int64(-65526), v2138)
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L1
	} else {
		goto L374
	}
L374:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2146 = F_shm_toc_allocate(m, v2145, v1244)
	mBase = m.M
	v2147 = m.ExcPending
	if v2147 != 0 {
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v2148 = int32(0)
	v2150 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	*(*int32)(unsafe.Add(mBase, uint32(v2146))) = v2150
	v2153 = int32(*(*uint8)(unsafe.Add(mBase, _consts[143])))
	*(*uint8)(unsafe.Add(mBase, uint32(v2146)+4)) = uint8(v2153)
	v2156 = *(*int64)(unsafe.Add(mBase, _consts[80]))
	*(*int64)(unsafe.Add(mBase, uint32(v2146)+8)) = v2156
	v2159 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v2160 = *(*int64)(unsafe.Add(mBase, uint32(v2159)))
	*(*int64)(unsafe.Add(mBase, uint32(v2146)+16)) = v2160
	v2163 = *(*int32)(unsafe.Add(mBase, _consts[144]))
	*(*int32)(unsafe.Add(mBase, uint32(v2146)+24)) = v2163
	v2166 = *(*int32)(unsafe.Add(mBase, _consts[145]))
	if v2148 < v2166 {
		goto L377
	} else {
		goto L378
	}
L376:
	;
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2350, int64(-65528), v2146)
	mBase = m.M
	v2353 = m.ExcPending
	if v2353 != 0 {
		goto L1
	} else {
		goto L414
	}
L377:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2146)+28)) = v2166
	v2173 = *(*int32)(unsafe.Add(mBase, _consts[146]))
	v2175 = v2166 << (uint(int32(2)) % 32)
	if v2175 != 0 {
		goto L381
	} else {
		goto L382
	}
L378:
	;
	goto L379
L379:
	;
	v2180 = v2148
	v2191 = v2159
	goto L384
L380:
	;
	goto L376
L381:
	;
	v2176 = F__emscripten_memcpy_bulkmem(m, v2146+int32(32), v2173, v2175)
	mBase = m.M
	goto L383
L382:
	;
	goto L383
L383:
	;
	goto L380
L384:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(v2191)))
	if v2209 != 0 {
		goto L386
	} else {
		goto L387
	}
L385:
	;
	v2220 = v2215 << (uint(int32(2)) % 32)
	v2221 = F_palloc(m, v2220)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L1
	} else {
		goto L392
	}
L386:
	;
	v2211 = F_add_size(m, v2180, int32(1))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L1
	} else {
		goto L389
	}
L387:
	;
	v2213 = v2180
	goto L388
L388:
	;
	v2214 = *(*int32)(unsafe.Add(mBase, uint32(v2191)+52))
	v2215 = F_add_size(m, v2213, v2214)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L390
	}
L389:
	;
	v2213 = v2211
	goto L388
L390:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v2191)+80))
	if v2217 != 0 {
		v2180 = v2215
		v2191 = v2217
		goto L384
	} else {
		goto L391
	}
L391:
	;
	goto L385
L392:
	;
	v2224 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	if v2224 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v2232 = int32(0)
	v2238 = v2224
	goto L396
L394:
	;
	goto L395
L395:
	;
	F_pg_qsort(m, v2221, v2215, int32(4), int32(185))
	mBase = m.M
	v2313 = m.ExcPending
	if v2313 != 0 {
		goto L1
	} else {
		goto L409
	}
L396:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v2238)))
	if v2256 != 0 {
		goto L398
	} else {
		goto L399
	}
L397:
	;
	goto L395
L398:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2221+v2232<<(uint(int32(2))%32)))) = v2256
	v2263 = v2232 + int32(1)
	goto L400
L399:
	;
	v2263 = v2232
	goto L400
L400:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+52))
	if int32(0) < v2264 {
		goto L401
	} else {
		goto L402
	}
L401:
	;
	v2267 = int32(2)
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+48))
	v2272 = v2264 << (uint(v2267) % 32)
	if v2272 != 0 {
		goto L405
	} else {
		goto L406
	}
L402:
	;
	v2276 = v2264
	goto L403
L403:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+80))
	if v2278 != 0 {
		v2232 = v2276 + v2263
		v2238 = v2278
		goto L396
	} else {
		goto L408
	}
L404:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(v2238)+52))
	v2276 = v2275
	goto L403
L405:
	;
	v2273 = F__emscripten_memcpy_bulkmem(m, v2221+v2263<<(uint(v2267)%32), v2270, v2272)
	mBase = m.M
	goto L407
L406:
	;
	goto L407
L407:
	;
	goto L404
L408:
	;
	goto L397
L409:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2146)+28)) = v2215
	if v2220 != 0 {
		goto L411
	} else {
		goto L412
	}
L410:
	;
	goto L376
L411:
	;
	v2317 = F__emscripten_memcpy_bulkmem(m, v2146+int32(32), v2221, v2220)
	mBase = m.M
	goto L413
L412:
	;
	goto L413
L413:
	;
	goto L410
L414:
	;
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2355 = F_shm_toc_allocate(m, v2354, v1238)
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	v2357 = m.G0
	v2359 = v2357 - int32(80)
	m.G0 = v2359
	v2362 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	if v2362 != 0 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2359)+48)) = int64(51539607564)
	v2366 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v2359)+72)) = v2366
	v2370 = *(*int32)(unsafe.Add(mBase, uint32(v2362)))
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+4))
	v2372 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+412))
	if v2372 != 0 {
		goto L420
	} else {
		goto L421
	}
L417:
	;
	v2678 = v2355
	goto L418
L418:
	;
	if v2678&int32(3) != 0 {
		goto L455
	} else {
		goto L456
	}
L419:
	;
	v2441 = F_hash_create(m, int32(131979), v2437, v2359+int32(32), int32(1064))
	mBase = m.M
	v2442 = m.ExcPending
	if v2442 != 0 {
		goto L1
	} else {
		goto L423
	}
L420:
	;
	v2373 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+376))
	v2374 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+364))
	v2375 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+352))
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+340))
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+328))
	v2378 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+316))
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+304))
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+292))
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+280))
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+268))
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+256))
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+244))
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+232))
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+220))
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+208))
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+196))
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+184))
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+172))
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+160))
	v2392 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+148))
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+136))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+124))
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+112))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+100))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+88))
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+76))
	v2401 = *(*int32)(unsafe.Add(mBase, uint32(v2370-int32(-64))))
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+52))
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+40))
	v2404 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+28))
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(v2370)+16))
	v2437 = v2373 + (v2374 + (v2375 + (v2376 + (v2377 + (v2378 + (v2379 + (v2380 + (v2381 + (v2382 + (v2383 + (v2384 + (v2385 + (v2386 + (v2387 + (v2388 + (v2389 + (v2390 + (v2391 + (v2392 + (v2393 + (v2394 + (v2395 + (v2396 + (v2397 + (v2398 + (v2401 + (v2402 + (v2403 + (v2404 + (v2405 + v2371))))))))))))))))))))))))))))))
	goto L422
L421:
	;
	v2437 = v2371
	goto L422
L422:
	;
	goto L419
L423:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, _consts[124]))
	F_hash_seq_init(m, v2359+int32(12), v2446)
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L1
	} else {
		goto L424
	}
L424:
	;
	v2451 = F_hash_seq_search(m, v2359+int32(12))
	mBase = m.M
	v2452 = m.ExcPending
	if v2452 != 0 {
		goto L1
	} else {
		goto L425
	}
L425:
	;
	if v2451 != 0 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v2455 = v2451
	goto L429
L427:
	;
	goto L428
L428:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, _consts[147]))
	if v2524 != 0 {
		goto L434
	} else {
		goto L435
	}
L429:
	;
	v2486 = F_hash_search(m, v2441, v2455, int32(1), int32(0))
	mBase = m.M
	v2487 = m.ExcPending
	if v2487 != 0 {
		goto L1
	} else {
		goto L431
	}
L430:
	;
	goto L428
L431:
	;
	v2490 = F_hash_seq_search(m, v2359+int32(12))
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L1
	} else {
		goto L432
	}
L432:
	;
	if v2490 != 0 {
		v2455 = v2490
		goto L429
	} else {
		goto L433
	}
L433:
	;
	goto L430
L434:
	;
	v2527 = v2524
	goto L437
L435:
	;
	goto L436
L436:
	;
	F_hash_seq_init(m, v2359+int32(12), v2441)
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L1
	} else {
		goto L444
	}
L437:
	;
	v2556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2527)+16)))
	if v2556 == int32(1) {
		goto L439
	} else {
		goto L440
	}
L438:
	;
	goto L436
L439:
	;
	v2561 = F_hash_search(m, v2441, v2527, int32(2), int32(0))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L1
	} else {
		goto L442
	}
L440:
	;
	goto L441
L441:
	;
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v2527)+24))
	if v2563 != 0 {
		v2527 = v2563
		goto L437
	} else {
		goto L443
	}
L442:
	;
	goto L441
L443:
	;
	goto L438
L444:
	;
	v2601 = F_hash_seq_search(m, v2359+int32(12))
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L1
	} else {
		goto L445
	}
L445:
	;
	if v2601 != 0 {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v2604 = v2355
	v2605 = v2601
	goto L449
L447:
	;
	v2645 = v2355
	goto L448
L448:
	;
	F_hash_destroy(m, v2441)
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L1
	} else {
		goto L453
	}
L449:
	;
	v2634 = *(*int64)(unsafe.Add(mBase, uint32(v2605)))
	*(*int64)(unsafe.Add(mBase, uint32(v2604))) = v2634
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v2605)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2604)+8)) = v2636
	v2638 = int32(12)
	v2639 = v2604 + v2638
	v2642 = F_hash_seq_search(m, v2359+v2638)
	mBase = m.M
	v2643 = m.ExcPending
	if v2643 != 0 {
		goto L1
	} else {
		goto L451
	}
L450:
	;
	v2645 = v2639
	goto L448
L451:
	;
	if v2642 != 0 {
		v2604 = v2639
		v2605 = v2642
		goto L449
	} else {
		goto L452
	}
L452:
	;
	goto L450
L453:
	;
	v2678 = v2645
	goto L418
L454:
	;
	m.G0 = v2359 + int32(80)
	v2736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2736, int64(-65525), v2355)
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L1
	} else {
		goto L463
	}
L455:
	;
	v2728 = int32(12)
	goto L457
L456:
	;
	v2713 = v2678 + int32(12)
	if base.Ui32(v2713) <= base.Ui32(v2678) {
		goto L454
	} else {
		goto L458
	}
L457:
	;
	v2730 = F__emscripten_memset_bulkmem(m, v2678, base.I32_extend8_s(int32(0)), v2728)
	mBase = m.M
	goto L462
L458:
	;
	v2718 = v2678 + int32(4)
	if base.Ui32(v2718) < base.Ui32(v2713) {
		goto L459
	} else {
		goto L460
	}
L459:
	;
	v2720 = v2713
	goto L461
L460:
	;
	v2720 = v2718
	goto L461
L461:
	;
	v2728 = (v2678^int32(-1)+v2720)&int32(-4) + int32(4)
	goto L457
L462:
	;
	goto L454
L463:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2741 = F_shm_toc_allocate(m, v2740, v1242)
	mBase = m.M
	v2742 = m.ExcPending
	if v2742 != 0 {
		goto L1
	} else {
		goto L464
	}
L464:
	;
	v2743 = int32(0)
	v2745 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v2741))) = v2745
	v2748 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v2741)+4)) = v2748
	v2751 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	if v2751 == v2743 {
		goto L466
	} else {
		goto L467
	}
L465:
	;
	v2867 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2867, int64(-65524), v2741)
	mBase = m.M
	v2870 = m.ExcPending
	if v2870 != 0 {
		goto L1
	} else {
		goto L475
	}
L466:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2741)+8)) = int32(0)
	goto L465
L467:
	;
	goto L468
L468:
	;
	v2756 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2741)+8)) = v2756
	v2758 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+4))
	if int32(0) < v2758 {
		goto L469
	} else {
		goto L470
	}
L469:
	;
	v2769 = v2743
	goto L472
L470:
	;
	goto L471
L471:
	;
	goto L465
L472:
	;
	v2795 = v2769 << (uint(int32(2)) % 32)
	v2797 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+12))
	v2799 = *(*int32)(unsafe.Add(mBase, uint32(v2797+v2795)))
	*(*int32)(unsafe.Add(mBase, uint32(v2741+int32(12)+v2795))) = v2799
	v2802 = v2769 + int32(1)
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(v2751)+4))
	if v2802 < v2803 {
		v2769 = v2802
		goto L472
	} else {
		goto L474
	}
L473:
	;
	goto L471
L474:
	;
	goto L473
L475:
	;
	v2871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2872 = F_shm_toc_allocate(m, v2871, v1246)
	mBase = m.M
	v2873 = m.ExcPending
	if v2873 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	goto L478
L477:
	;
	v2878 = int32(524)
	goto L482
L478:
	;
	v2876 = F__emscripten_memcpy_bulkmem(m, v2872, int32(4508544), int32(524))
	mBase = m.M
	goto L480
L480:
	;
	goto L477
L481:
	;
	v2884 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2884, int64(-65523), v2872)
	mBase = m.M
	v2887 = m.ExcPending
	if v2887 != 0 {
		goto L1
	} else {
		goto L485
	}
L482:
	;
	v2882 = F__emscripten_memcpy_bulkmem(m, v2876+v2878, int32(4509592), v2878)
	mBase = m.M
	goto L484
L484:
	;
	goto L481
L485:
	;
	v2888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2889 = F_shm_toc_allocate(m, v2888, v1245)
	mBase = m.M
	v2890 = m.ExcPending
	if v2890 != 0 {
		goto L1
	} else {
		goto L486
	}
L486:
	;
	v2891 = m.G0
	v2893 = v2891 - int32(32)
	m.G0 = v2893
	v2896 = *(*int32)(unsafe.Add(mBase, _consts[125]))
	if v2896 == int32(0) {
		v2949 = v2889
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2979 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2949))) = v2979
	v2982 = v2949 + int32(4)
	v2984 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	if v2984 == v2979 {
		v3037 = v2982
		goto L496
	} else {
		goto L497
	}
L488:
	;
	F_hash_seq_init(m, v2893+int32(12), v2896)
	mBase = m.M
	v2902 = m.ExcPending
	if v2902 != 0 {
		goto L1
	} else {
		goto L489
	}
L489:
	;
	v2905 = F_hash_seq_search(m, v2893+int32(12))
	mBase = m.M
	v2906 = m.ExcPending
	if v2906 != 0 {
		goto L1
	} else {
		goto L490
	}
L490:
	;
	if v2905 == int32(0) {
		v2949 = v2889
		goto L487
	} else {
		goto L491
	}
L491:
	;
	v2910 = v2889
	v2915 = v2905
	goto L492
L492:
	;
	v2940 = *(*int32)(unsafe.Add(mBase, uint32(v2915)))
	*(*int32)(unsafe.Add(mBase, uint32(v2910))) = v2940
	v2943 = v2910 + int32(4)
	v2946 = F_hash_seq_search(m, v2893+int32(12))
	mBase = m.M
	v2947 = m.ExcPending
	if v2947 != 0 {
		goto L1
	} else {
		goto L494
	}
L493:
	;
	v2949 = v2943
	goto L487
L494:
	;
	if v2946 != 0 {
		v2910 = v2943
		v2915 = v2946
		goto L492
	} else {
		goto L495
	}
L495:
	;
	goto L493
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3037))) = int32(0)
	m.G0 = v2893 + int32(32)
	v3072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3072, int64(-65522), v2889)
	mBase = m.M
	v3075 = m.ExcPending
	if v3075 != 0 {
		goto L1
	} else {
		goto L505
	}
L497:
	;
	F_hash_seq_init(m, v2893+int32(12), v2984)
	mBase = m.M
	v2990 = m.ExcPending
	if v2990 != 0 {
		goto L1
	} else {
		goto L498
	}
L498:
	;
	v2993 = F_hash_seq_search(m, v2893+int32(12))
	mBase = m.M
	v2994 = m.ExcPending
	if v2994 != 0 {
		goto L1
	} else {
		goto L499
	}
L499:
	;
	if v2993 == int32(0) {
		v3037 = v2982
		goto L496
	} else {
		goto L500
	}
L500:
	;
	v2998 = v2982
	v3003 = v2993
	goto L501
L501:
	;
	v3028 = *(*int32)(unsafe.Add(mBase, uint32(v3003)))
	*(*int32)(unsafe.Add(mBase, uint32(v2998))) = v3028
	v3031 = v2998 + int32(4)
	v3034 = F_hash_seq_search(m, v2893+int32(12))
	mBase = m.M
	v3035 = m.ExcPending
	if v3035 != 0 {
		goto L1
	} else {
		goto L503
	}
L502:
	;
	v3037 = v3031
	goto L496
L503:
	;
	if v3034 != 0 {
		v2998 = v3031
		v3003 = v3034
		goto L501
	} else {
		goto L504
	}
L504:
	;
	goto L502
L505:
	;
	v3076 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3077 = F_shm_toc_allocate(m, v3076, v1248)
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L1
	} else {
		goto L506
	}
L506:
	;
	v3080 = *(*int32)(unsafe.Add(mBase, _consts[149]))
	v3082 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	if v3082 == int32(0) {
		goto L508
	} else {
		goto L509
	}
L507:
	;
	v3102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3102, int64(-65521), v3077)
	mBase = m.M
	v3105 = m.ExcPending
	if v3105 != 0 {
		goto L1
	} else {
		goto L518
	}
L508:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3077)+4)) = v3080
	*(*int32)(unsafe.Add(mBase, uint32(v3077))) = int32(-1)
	goto L507
L509:
	;
	goto L510
L510:
	;
	v3088 = F_strlen(m, v3082)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3077)+4)) = v3080
	*(*int32)(unsafe.Add(mBase, uint32(v3077))) = v3088
	if int32(0) <= v3088 {
		goto L511
	} else {
		goto L512
	}
L511:
	;
	v3096 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	v3098 = v3088 + int32(1)
	if v3098 != 0 {
		goto L515
	} else {
		goto L516
	}
L512:
	;
	goto L513
L513:
	;
	goto L507
L514:
	;
	goto L513
L515:
	;
	v3099 = F__emscripten_memcpy_bulkmem(m, v3077+int32(8), v3096, v3098)
	mBase = m.M
	goto L517
L516:
	;
	goto L517
L517:
	;
	goto L514
L518:
	;
	v3106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3109 = F_palloc0(m, v3106<<(uint(int32(3))%32))
	mBase = m.M
	v3110 = m.ExcPending
	if v3110 != 0 {
		goto L1
	} else {
		goto L519
	}
L519:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v3109
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3115 = F_mul_size(m, int32(16384), v3114)
	mBase = m.M
	v3116 = m.ExcPending
	if v3116 != 0 {
		goto L1
	} else {
		goto L520
	}
L520:
	;
	v3117 = F_shm_toc_allocate(m, v3112, v3115)
	mBase = m.M
	v3118 = m.ExcPending
	if v3118 != 0 {
		goto L1
	} else {
		goto L521
	}
L521:
	;
	v3119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v3119 {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v3126 = int32(0)
	goto L525
L523:
	;
	goto L524
L524:
	;
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3220, int64(-65534), v3117)
	mBase = m.M
	v3223 = m.ExcPending
	if v3223 != 0 {
		goto L1
	} else {
		goto L531
	}
L525:
	;
	v3156 = v3117 + v3126<<(uint(int32(14))%32)
	v3158 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3156)+16)) = v3158
	*(*int32)(unsafe.Add(mBase, uint32(v3156)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3156))) = v3158
	v3164 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v3156)+36)) = uint16(v3164)
	*(*int64)(unsafe.Add(mBase, uint32(v3156)+24)) = v3158
	*(*int32)(unsafe.Add(mBase, uint32(v3156)+32)) = int32(16344)
	goto L527
L526:
	;
	goto L524
L527:
	;
	v3174 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	F_shm_mq_set_receiver(m, v3156, v3174)
	mBase = m.M
	v3176 = m.ExcPending
	if v3176 != 0 {
		goto L1
	} else {
		goto L528
	}
L528:
	;
	v3177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v3178 = F_shm_mq_attach(m, v3156, v3177)
	mBase = m.M
	v3179 = m.ExcPending
	if v3179 != 0 {
		goto L1
	} else {
		goto L529
	}
L529:
	;
	v3180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3180+v3126<<(uint(int32(3))%32))+4)) = v3178
	v3186 = v3126 + int32(1)
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3186 < v3187 {
		v3126 = v3186
		goto L525
	} else {
		goto L530
	}
L530:
	;
	goto L526
L531:
	;
	v3224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3225 = F_strlen(m, v3224)
	mBase = m.M
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3228 = F_strlen(m, v3227)
	mBase = m.M
	v3232 = F_shm_toc_allocate(m, v3226, v3228+v3225+int32(2))
	mBase = m.M
	v3233 = m.ExcPending
	if v3233 != 0 {
		goto L1
	} else {
		goto L532
	}
L532:
	;
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if (v3234^v3232)&int32(3) != 0 {
		goto L536
	} else {
		goto L537
	}
L533:
	;
	v3311 = v3225 + v3232 + int32(1)
	v3312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if (v3312^v3311)&int32(3) != 0 {
		goto L557
	} else {
		goto L558
	}
L534:
	;
	goto L533
L535:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3289))) = uint8(v3288)
	if v3288&int32(255) == int32(0) {
		goto L534
	} else {
		goto L550
	}
L536:
	;
	v3240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3234))))
	v3287 = v3234
	v3288 = v3240
	v3289 = v3232
	goto L535
L537:
	;
	goto L538
L538:
	;
	if v3234&int32(3) != 0 {
		goto L539
	} else {
		goto L540
	}
L539:
	;
	v3244 = v3234
	v3246 = v3232
	goto L542
L540:
	;
	v3258 = v3234
	v3260 = v3232
	goto L541
L541:
	;
	v3262 = *(*int32)(unsafe.Add(mBase, uint32(v3258)))
	v3265 = int32(-2139062144)
	if (int32(16843008)-v3262|v3262)&v3265 != v3265 {
		v3287 = v3258
		v3288 = v3262
		v3289 = v3260
		goto L535
	} else {
		goto L546
	}
L542:
	;
	v3247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3244))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3246))) = uint8(v3247)
	if v3247 == int32(0) {
		goto L534
	} else {
		goto L544
	}
L543:
	;
	v3258 = v3254
	v3260 = v3252
	goto L541
L544:
	;
	v3251 = int32(1)
	v3252 = v3246 + v3251
	v3254 = v3244 + v3251
	if v3254&int32(3) != 0 {
		v3244 = v3254
		v3246 = v3252
		goto L542
	} else {
		goto L545
	}
L545:
	;
	goto L543
L546:
	;
	v3270 = v3258
	v3271 = v3262
	v3272 = v3260
	goto L547
L547:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3272))) = v3271
	v3274 = int32(4)
	v3275 = v3272 + v3274
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3270)+4))
	v3278 = v3270 + v3274
	v3282 = int32(-2139062144)
	if (v3276|(int32(16843008)-v3276))&v3282 == v3282 {
		v3270 = v3278
		v3271 = v3276
		v3272 = v3275
		goto L547
	} else {
		goto L549
	}
L548:
	;
	v3287 = v3278
	v3288 = v3276
	v3289 = v3275
	goto L535
L549:
	;
	goto L548
L550:
	;
	v3296 = v3287
	v3298 = v3289
	goto L551
L551:
	;
	v3299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3296)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3298)+1)) = uint8(v3299)
	v3301 = int32(1)
	if v3299 != 0 {
		v3296 = v3296 + v3301
		v3298 = v3298 + v3301
		goto L551
	} else {
		goto L553
	}
L552:
	;
	goto L534
L553:
	;
	goto L552
L554:
	;
	v3387 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3387, int64(-65527), v3232)
	mBase = m.M
	v3390 = m.ExcPending
	if v3390 != 0 {
		goto L1
	} else {
		goto L575
	}
L555:
	;
	goto L554
L556:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3367))) = uint8(v3366)
	if v3366&int32(255) == int32(0) {
		goto L555
	} else {
		goto L571
	}
L557:
	;
	v3318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3312))))
	v3365 = v3312
	v3366 = v3318
	v3367 = v3311
	goto L556
L558:
	;
	goto L559
L559:
	;
	if v3312&int32(3) != 0 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v3322 = v3312
	v3324 = v3311
	goto L563
L561:
	;
	v3336 = v3312
	v3338 = v3311
	goto L562
L562:
	;
	v3340 = *(*int32)(unsafe.Add(mBase, uint32(v3336)))
	v3343 = int32(-2139062144)
	if (int32(16843008)-v3340|v3340)&v3343 != v3343 {
		v3365 = v3336
		v3366 = v3340
		v3367 = v3338
		goto L556
	} else {
		goto L567
	}
L563:
	;
	v3325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3322))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3324))) = uint8(v3325)
	if v3325 == int32(0) {
		goto L555
	} else {
		goto L565
	}
L564:
	;
	v3336 = v3332
	v3338 = v3330
	goto L562
L565:
	;
	v3329 = int32(1)
	v3330 = v3324 + v3329
	v3332 = v3322 + v3329
	if v3332&int32(3) != 0 {
		v3322 = v3332
		v3324 = v3330
		goto L563
	} else {
		goto L566
	}
L566:
	;
	goto L564
L567:
	;
	v3348 = v3336
	v3349 = v3340
	v3350 = v3338
	goto L568
L568:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3350))) = v3349
	v3352 = int32(4)
	v3353 = v3350 + v3352
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(v3348)+4))
	v3356 = v3348 + v3352
	v3360 = int32(-2139062144)
	if (v3354|(int32(16843008)-v3354))&v3360 == v3360 {
		v3348 = v3356
		v3349 = v3354
		v3350 = v3353
		goto L568
	} else {
		goto L570
	}
L569:
	;
	v3365 = v3356
	v3366 = v3354
	v3367 = v3353
	goto L556
L570:
	;
	goto L569
L571:
	;
	v3374 = v3365
	v3376 = v3367
	goto L572
L572:
	;
	v3377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3374)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3376)+1)) = uint8(v3377)
	v3379 = int32(1)
	if v3377 != 0 {
		v3374 = v3374 + v3379
		v3376 = v3376 + v3379
		goto L572
	} else {
		goto L574
	}
L573:
	;
	goto L555
L574:
	;
	goto L573
L575:
	;
	v3391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3423 = v3391
	goto L211
}
func F_parallel_vacuum_process_one_index(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 float64
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l1
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v14
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+24)) = uint16(v4)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+26)) = uint8(v21)
	v23 = *(*float64)(unsafe.Add(mBase, uint32(v20)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v10)+32)) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v30 = F_pstrdup(m, v27+int32(4))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v30
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v33
		v36 = l2 + int32(8)
		if v12 != 0 {
			v38 = v36
		} else {
			v38 = int32(0)
		}
		switch v33 - int32(1) {
		case 0:
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v69 = F_vac_bulkdel_one_index(m, v10+int32(16), v38, v65, v66+int32(56))
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				v71 = v69
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
				if v72 != 0 {
					v89 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v89
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v89
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					F_pfree(m, v93)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
						F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					}
				} else {
					if v71 == int32(0) {
						v89 = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v89
						*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v89
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_pfree(m, v93)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
							F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					} else {
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
						*(*int64)(unsafe.Add(mBase, uint32(v36))) = v75
						v77 = *(*int64)(unsafe.Add(mBase, uint32(v71)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v77
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v71)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v79
						v81 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v81
						v83 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v83
						v85 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v85)
						F_pfree(m, v71)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							v89 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v89
							*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v89
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_pfree(m, v93)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
								F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									m.G0 = v10 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		case 1:
			v43 = F_vac_cleanup_one_index(m, v10+int32(16), v38)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v71 = v43
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
				if v72 != 0 {
					v89 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v89
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v89
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					F_pfree(m, v93)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
						F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
						}
					}
				} else {
					if v71 == int32(0) {
						v89 = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v89
						*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v89
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_pfree(m, v93)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
							F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					} else {
						v75 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
						*(*int64)(unsafe.Add(mBase, uint32(v36))) = v75
						v77 = *(*int64)(unsafe.Add(mBase, uint32(v71)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v77
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v71)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v79
						v81 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v81
						v83 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v83
						v85 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v85)
						F_pfree(m, v71)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							v89 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v89
							*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v89
							v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
							F_pfree(m, v93)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
								F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									m.G0 = v10 + int32(48)
									return
								}
							}
						}
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v50
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v49 + int32(4)
				F_errmsg_internal(m, int32(694562), v10)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(498630), int32(904), int32(28306))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
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
}
