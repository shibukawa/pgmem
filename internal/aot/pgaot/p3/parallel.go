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
			v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
			mBase = m.M
			v195 = m.ExcPending
			if v195 != 0 {
				return int32(0)
			} else {
				return v194
			}
		default:
			v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
											v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
													v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
														v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
													v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
															v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
														v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																	v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
													v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
															v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
															v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																	v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																		v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																		v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																			v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
														v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																	v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																		v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																			v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																	v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																			v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
																				v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
						v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return int32(0)
						} else {
							return v194
						}
					}
				} else {
					v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
							v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
							v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
				v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
										v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
								v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
					v194 = F_planstate_tree_walker_impl(m, l0, int32(627), l1)
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
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v714 int32
	_ = v714
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v834 int32
	_ = v834
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v885 int32
	_ = v885
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v933 int32
	_ = v933
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v977 int32
	_ = v977
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1109 int32
	_ = v1109
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1149 int32
	_ = v1149
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1176 int32
	_ = v1176
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1252 int32
	_ = v1252
	var v1256 int32
	_ = v1256
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1385 int32
	_ = v1385
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1455 int32
	_ = v1455
	var v1463 int32
	_ = v1463
	var v1468 int32
	_ = v1468
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1483 int32
	_ = v1483
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1540 int32
	_ = v1540
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1557 int32
	_ = v1557
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1577 int32
	_ = v1577
	var v1583 int32
	_ = v1583
	var v1586 int32
	_ = v1586
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1615 int32
	_ = v1615
	var v1620 int32
	_ = v1620
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1631 int32
	_ = v1631
	var v1635 int32
	_ = v1635
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1650 int32
	_ = v1650
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1664 int32
	_ = v1664
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1702 int32
	_ = v1702
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1728 int32
	_ = v1728
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1746 int32
	_ = v1746
	var v1749 int32
	_ = v1749
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1784 int32
	_ = v1784
	var v1787 int32
	_ = v1787
	var v1790 int64
	_ = v1790
	var v1793 int64
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1811 int32
	_ = v1811
	var v1813 int32
	_ = v1813
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1846 int32
	_ = v1846
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1927 int32
	_ = v1927
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1932 int32
	_ = v1932
	var v1933 int32
	_ = v1933
	var v1934 int32
	_ = v1934
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1955 int32
	_ = v1955
	var v1958 int32
	_ = v1958
	var v1962 int32
	_ = v1962
	var v1963 int32
	_ = v1963
	var v1965 int32
	_ = v1965
	var v1973 int32
	_ = v1973
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2010 int32
	_ = v2010
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2026 int32
	_ = v2026
	var v2031 int32
	_ = v2031
	var v2055 int32
	_ = v2055
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2092 int32
	_ = v2092
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2104 float64
	_ = v2104
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2120 int32
	_ = v2120
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2140 int32
	_ = v2140
	var v2167 int32
	_ = v2167
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2185 int32
	_ = v2185
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2221 int32
	_ = v2221
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2231 int32
	_ = v2231
	var v2232 int32
	_ = v2232
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2275 int32
	_ = v2275
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2294 int32
	_ = v2294
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2301 int32
	_ = v2301
	var v2306 int32
	_ = v2306
	var v2312 int32
	_ = v2312
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2324 int32
	_ = v2324
	var v2351 int32
	_ = v2351
	var v2358 int32
	_ = v2358
	var v2393 int32
	_ = v2393
	var v2397 int32
	_ = v2397
	var v2402 int32
	_ = v2402
	var v2406 int32
	_ = v2406
	var v2410 int32
	_ = v2410
	var v2415 int32
	_ = v2415
	var v2419 int32
	_ = v2419
	var v2423 int32
	_ = v2423
	var v2428 int32
	_ = v2428
	var v2432 int32
	_ = v2432
	var v2436 int32
	_ = v2436
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2450 int32
	_ = v2450
	var v2453 int32
	_ = v2453
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2469 int32
	_ = v2469
	var v2473 int32
	_ = v2473
	var v2478 int32
	_ = v2478
	var v2479 int32
	_ = v2479
	var v2482 int32
	_ = v2482
	var v2484 int32
	_ = v2484
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2489 int32
	_ = v2489
	var v2496 int32
	_ = v2496
	var v2497 int32
	_ = v2497
	var v2498 int64
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2508 int32
	_ = v2508
	var v2509 int32
	_ = v2509
	var v2511 int32
	_ = v2511
	var v2514 int32
	_ = v2514
	var v2517 int32
	_ = v2517
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2534 int32
	_ = v2534
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2547 int64
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2560 int32
	_ = v2560
	var v2563 int32
	_ = v2563
	var v2566 int32
	_ = v2566
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2579 int32
	_ = v2579
	var v2580 int32
	_ = v2580
	var v2583 int32
	_ = v2583
	var v2584 int32
	_ = v2584
	var v2586 int32
	_ = v2586
	var v2587 int32
	_ = v2587
	var v2589 int32
	_ = v2589
	var v2592 int32
	_ = v2592
	var v2593 int32
	_ = v2593
	var v2594 int32
	_ = v2594
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2598 int32
	_ = v2598
	var v2601 int32
	_ = v2601
	var v2604 int64
	_ = v2604
	var v2607 int32
	_ = v2607
	var v2608 int64
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2614 int32
	_ = v2614
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2628 int32
	_ = v2628
	var v2639 int32
	_ = v2639
	var v2657 int32
	_ = v2657
	var v2659 int32
	_ = v2659
	var v2660 int32
	_ = v2660
	var v2661 int32
	_ = v2661
	var v2662 int32
	_ = v2662
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2680 int32
	_ = v2680
	var v2686 int32
	_ = v2686
	var v2704 int32
	_ = v2704
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2715 int32
	_ = v2715
	var v2718 int32
	_ = v2718
	var v2720 int32
	_ = v2720
	var v2721 int32
	_ = v2721
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2761 int32
	_ = v2761
	var v2765 int32
	_ = v2765
	var v2798 int32
	_ = v2798
	var v2801 int32
	_ = v2801
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2807 int32
	_ = v2807
	var v2810 int32
	_ = v2810
	var v2814 int32
	_ = v2814
	var v2818 int32
	_ = v2818
	var v2819 int32
	_ = v2819
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2822 int32
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2824 int32
	_ = v2824
	var v2825 int32
	_ = v2825
	var v2826 int32
	_ = v2826
	var v2827 int32
	_ = v2827
	var v2828 int32
	_ = v2828
	var v2829 int32
	_ = v2829
	var v2830 int32
	_ = v2830
	var v2831 int32
	_ = v2831
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2835 int32
	_ = v2835
	var v2836 int32
	_ = v2836
	var v2837 int32
	_ = v2837
	var v2838 int32
	_ = v2838
	var v2839 int32
	_ = v2839
	var v2840 int32
	_ = v2840
	var v2841 int32
	_ = v2841
	var v2842 int32
	_ = v2842
	var v2843 int32
	_ = v2843
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2849 int32
	_ = v2849
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2853 int32
	_ = v2853
	var v2885 int32
	_ = v2885
	var v2889 int32
	_ = v2889
	var v2890 int32
	_ = v2890
	var v2894 int32
	_ = v2894
	var v2896 int32
	_ = v2896
	var v2899 int32
	_ = v2899
	var v2900 int32
	_ = v2900
	var v2903 int32
	_ = v2903
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2972 int32
	_ = v2972
	var v2975 int32
	_ = v2975
	var v3004 int32
	_ = v3004
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3046 int32
	_ = v3046
	var v3049 int32
	_ = v3049
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3082 int64
	_ = v3082
	var v3084 int32
	_ = v3084
	var v3086 int32
	_ = v3086
	var v3087 int32
	_ = v3087
	var v3090 int32
	_ = v3090
	var v3091 int32
	_ = v3091
	var v3093 int32
	_ = v3093
	var v3124 int32
	_ = v3124
	var v3126 int32
	_ = v3126
	var v3161 int32
	_ = v3161
	var v3166 int32
	_ = v3166
	var v3168 int32
	_ = v3168
	var v3176 int32
	_ = v3176
	var v3178 int32
	_ = v3178
	var v3184 int32
	_ = v3184
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3189 int32
	_ = v3189
	var v3190 int32
	_ = v3190
	var v3191 int32
	_ = v3191
	var v3193 int32
	_ = v3193
	var v3196 int32
	_ = v3196
	var v3199 int32
	_ = v3199
	var v3204 int32
	_ = v3204
	var v3206 int32
	_ = v3206
	var v3217 int32
	_ = v3217
	var v3243 int32
	_ = v3243
	var v3245 int32
	_ = v3245
	var v3247 int32
	_ = v3247
	var v3250 int32
	_ = v3250
	var v3251 int32
	_ = v3251
	var v3315 int32
	_ = v3315
	var v3318 int32
	_ = v3318
	var v3319 int32
	_ = v3319
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3324 int32
	_ = v3324
	var v3326 int32
	_ = v3326
	var v3330 int32
	_ = v3330
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3336 int32
	_ = v3336
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3341 int32
	_ = v3341
	var v3344 int32
	_ = v3344
	var v3350 int32
	_ = v3350
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3358 int32
	_ = v3358
	var v3363 int32
	_ = v3363
	var v3388 int32
	_ = v3388
	var v3391 int32
	_ = v3391
	var v3394 int32
	_ = v3394
	var v3395 int32
	_ = v3395
	var v3397 int32
	_ = v3397
	var v3427 int32
	_ = v3427
	var v3430 int32
	_ = v3430
	var v3432 int32
	_ = v3432
	var v3438 int32
	_ = v3438
	var v3441 int32
	_ = v3441
	var v3442 int32
	_ = v3442
	var v3446 int32
	_ = v3446
	var v3451 int32
	_ = v3451
	var v3476 int32
	_ = v3476
	var v3479 int32
	_ = v3479
	var v3482 int32
	_ = v3482
	var v3483 int32
	_ = v3483
	var v3485 int32
	_ = v3485
	var v3520 int32
	_ = v3520
	var v3523 int32
	_ = v3523
	var v3524 int32
	_ = v3524
	var v3525 int32
	_ = v3525
	var v3526 int32
	_ = v3526
	var v3528 int32
	_ = v3528
	var v3530 int32
	_ = v3530
	var v3543 int32
	_ = v3543
	var v3548 int32
	_ = v3548
	var v3552 int32
	_ = v3552
	var v3557 int32
	_ = v3557
	var v3559 int32
	_ = v3559
	var v3563 int32
	_ = v3563
	var v3569 int32
	_ = v3569
	var v3572 int32
	_ = v3572
	var v3578 int32
	_ = v3578
	var v3582 int32
	_ = v3582
	var v3584 int32
	_ = v3584
	var v3592 int32
	_ = v3592
	var v3600 int32
	_ = v3600
	var v3602 int32
	_ = v3602
	var v3603 int32
	_ = v3603
	var v3606 int32
	_ = v3606
	var v3609 int32
	_ = v3609
	var v3610 int32
	_ = v3610
	var v3613 int32
	_ = v3613
	var v3614 int32
	_ = v3614
	var v3616 int32
	_ = v3616
	var v3618 int32
	_ = v3618
	var v3619 int32
	_ = v3619
	var v3620 int32
	_ = v3620
	var v3621 int32
	_ = v3621
	var v3622 int32
	_ = v3622
	var v3623 int32
	_ = v3623
	var v3630 int32
	_ = v3630
	var v3660 int32
	_ = v3660
	var v3662 int64
	_ = v3662
	var v3668 int32
	_ = v3668
	var v3678 int32
	_ = v3678
	var v3680 int32
	_ = v3680
	var v3681 int32
	_ = v3681
	var v3682 int32
	_ = v3682
	var v3683 int32
	_ = v3683
	var v3684 int32
	_ = v3684
	var v3690 int32
	_ = v3690
	var v3691 int32
	_ = v3691
	var v3724 int32
	_ = v3724
	var v3727 int32
	_ = v3727
	var v3728 int32
	_ = v3728
	var v3736 int32
	_ = v3736
	var v3741 int32
	_ = v3741
	var v3745 int32
	_ = v3745
	var v3750 int32
	_ = v3750
	var v3752 int32
	_ = v3752
	var v3756 int32
	_ = v3756
	var v3762 int32
	_ = v3762
	var v3765 int32
	_ = v3765
	var v3771 int32
	_ = v3771
	var v3775 int32
	_ = v3775
	var v3777 int32
	_ = v3777
	var v3785 int32
	_ = v3785
	var v3786 int32
	_ = v3786
	var v3787 int32
	_ = v3787
	var v3795 int32
	_ = v3795
	var v3800 int32
	_ = v3800
	var v3804 int32
	_ = v3804
	var v3809 int32
	_ = v3809
	var v3811 int32
	_ = v3811
	var v3815 int32
	_ = v3815
	var v3821 int32
	_ = v3821
	var v3824 int32
	_ = v3824
	var v3830 int32
	_ = v3830
	var v3834 int32
	_ = v3834
	var v3836 int32
	_ = v3836
	var v3844 int32
	_ = v3844
	var v3848 int32
	_ = v3848
	var v3849 int32
	_ = v3849
	var v3850 int32
	_ = v3850
	var v3856 int32
	_ = v3856
	var v3860 int32
	_ = v3860
	var v3862 int32
	_ = v3862
	var v3863 int32
	_ = v3863
	var v3867 int32
	_ = v3867
	var v3868 int32
	_ = v3868
	var v3870 int32
	_ = v3870
	var v3874 int32
	_ = v3874
	var v3876 int32
	_ = v3876
	var v3878 int32
	_ = v3878
	var v3881 int32
	_ = v3881
	var v3886 int32
	_ = v3886
	var v3887 int32
	_ = v3887
	var v3888 int32
	_ = v3888
	var v3890 int32
	_ = v3890
	var v3891 int32
	_ = v3891
	var v3892 int32
	_ = v3892
	var v3894 int32
	_ = v3894
	var v3898 int32
	_ = v3898
	var v3903 int32
	_ = v3903
	var v3904 int32
	_ = v3904
	var v3905 int32
	_ = v3905
	var v3912 int32
	_ = v3912
	var v3914 int32
	_ = v3914
	var v3915 int32
	_ = v3915
	var v3917 int32
	_ = v3917
	var v3927 int32
	_ = v3927
	var v3928 int32
	_ = v3928
	var v3934 int32
	_ = v3934
	var v3938 int32
	_ = v3938
	var v3940 int32
	_ = v3940
	var v3941 int32
	_ = v3941
	var v3945 int32
	_ = v3945
	var v3946 int32
	_ = v3946
	var v3948 int32
	_ = v3948
	var v3952 int32
	_ = v3952
	var v3954 int32
	_ = v3954
	var v3956 int32
	_ = v3956
	var v3959 int32
	_ = v3959
	var v3964 int32
	_ = v3964
	var v3965 int32
	_ = v3965
	var v3966 int32
	_ = v3966
	var v3968 int32
	_ = v3968
	var v3969 int32
	_ = v3969
	var v3970 int32
	_ = v3970
	var v3972 int32
	_ = v3972
	var v3976 int32
	_ = v3976
	var v3981 int32
	_ = v3981
	var v3982 int32
	_ = v3982
	var v3983 int32
	_ = v3983
	var v3990 int32
	_ = v3990
	var v3992 int32
	_ = v3992
	var v3993 int32
	_ = v3993
	var v3995 int32
	_ = v3995
	var v4003 int32
	_ = v4003
	var v4006 int32
	_ = v4006
	var v4007 int32
	_ = v4007
	var v4039 int32
	_ = v4039
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
	v35 = *(*int32)(unsafe.Add(mBase, _consts[104]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	goto L3
L3:
	;
	v37 = int32(4464496)
	v38 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v41 = *(*int32)(unsafe.Add(mBase, _consts[105]))
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
	v56 = *(*int32)(unsafe.Add(mBase, _consts[106]))
	if v56 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v1709 = F_shm_toc_estimate(m, v54)
	mBase = m.M
	v1710 = m.ExcPending
	if v1710 != 0 {
		goto L1
	} else {
		goto L326
	}
L7:
	;
	v68 = l0 + int32(12)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v69 <= int32(0) {
		v1682 = v2
		v1684 = v2
		v1686 = v2
		v1688 = v68
		v1690 = v2
		v1691 = v2
		v1692 = v2
		v1693 = v2
		v1694 = v2
		v1695 = v2
		v1696 = v2
		v1702 = v2
		v1706 = v2
		goto L6
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1682 = v2
	v1684 = v2
	v1686 = v2
	v1688 = l0 + int32(12)
	v1690 = v2
	v1691 = v2
	v1692 = v2
	v1693 = v2
	v1694 = v2
	v1695 = v2
	v1696 = v2
	v1702 = v2
	v1706 = v2
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
	v60 = *(*int32)(unsafe.Add(mBase, _consts[107]))
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
	v77 = *(*int32)(unsafe.Add(mBase, _consts[108]))
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
	v80 = int32(4464496)
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
	v141 = int32(4464496)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v145 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v145
	v148 = F_dshash_create(m, v129, int32(1733452), v129)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v152 = F_dshash_create(m, v129, int32(1733476), int32(0))
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
	v163 = *(*int32)(unsafe.Add(mBase, _consts[109]))
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
	v168 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	v174 = v163
	v188 = v168
	v189 = v2
	goto L40
L38:
	;
	goto L39
L39:
	;
	v291 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	*(*int32)(unsafe.Add(mBase, uint32(v291)+16)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v291)+12)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v291)+8)) = v135
	F_on_dsm_detach(m, v107, int32(1622), int32(0))
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
	v249 = *(*int32)(unsafe.Add(mBase, _consts[110]))
	v251 = *(*int32)(unsafe.Add(mBase, _consts[109]))
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
	F_errmsg_internal(m, int32(412868), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(487973), int32(2253), int32(98216))
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
	v322 = int32(4358956)
	v323 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = v107
	v326 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	*(*int32)(unsafe.Add(mBase, uint32(v326)+4)) = v129
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v81
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v362 = v330
	goto L13
L64:
	;
	v368 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v368
	v1682 = v2
	v1684 = v2
	v1686 = v2
	v1688 = v68
	v1690 = v2
	v1691 = v2
	v1692 = v2
	v1693 = v2
	v1694 = v2
	v1695 = v368
	v1696 = v2
	v1702 = v2
	v1706 = v2
	goto L6
L65:
	;
	goto L66
L66:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v371 <= int32(0) {
		v1682 = v2
		v1684 = v2
		v1686 = v2
		v1688 = v68
		v1690 = v2
		v1691 = v2
		v1692 = v2
		v1693 = v2
		v1694 = v2
		v1695 = v362
		v1696 = v2
		v1702 = v2
		v1706 = v2
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v374 = int32(1)
	v376 = *(*int32)(unsafe.Add(mBase, _consts[111]))
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
	v478 = v374
	goto L70
L70:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v508 = F_add_size(m, v503, (v478+int32(31))&int32(-32))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L92
	}
L71:
	;
	v409 = v381 + int32(24)
	if v409&int32(3) == int32(0) {
		v433 = v409
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v478 = v469
	goto L70
L73:
	;
	v469 = F_add_size(m, v383, v466+int32(1))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L90
	}
L74:
	;
	v466 = v458 - v409
	goto L73
L75:
	;
	v437 = v433
	goto L84
L76:
	;
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v409))))
	if v417 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v466 = int32(0)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v422 = v409
	goto L80
L80:
	;
	v426 = v422 + int32(1)
	if v426&int32(3) == int32(0) {
		v433 = v426
		goto L75
	} else {
		goto L82
	}
L81:
	;
	v458 = v426
	goto L74
L82:
	;
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v426))))
	if v431 != 0 {
		v422 = v426
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v446 = int32(-2139062144)
	if (int32(16843008)-v443|v443)&v446 == v446 {
		v437 = v437 + int32(4)
		goto L84
	} else {
		goto L86
	}
L85:
	;
	v452 = v437
	goto L87
L86:
	;
	goto L85
L87:
	;
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v452))))
	if v456 != 0 {
		v452 = v452 + int32(1)
		goto L87
	} else {
		goto L89
	}
L88:
	;
	v458 = v452
	goto L74
L89:
	;
	goto L88
L90:
	;
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v471 != 0 {
		v381 = v471
		v383 = v469
		goto L71
	} else {
		goto L91
	}
L91:
	;
	goto L72
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v508
	v511 = m.G0
	v513 = v511 - int32(16)
	m.G0 = v513
	v515 = int32(4)
	v517 = *(*int32)(unsafe.Add(mBase, _consts[112]))
	if v517 == int32(0) {
		v977 = v515
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v1046 = F_add_size(m, v1041, (v977+int32(31))&int32(-32))
	mBase = m.M
	v1047 = m.ExcPending
	if v1047 != 0 {
		goto L1
	} else {
		goto L207
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L204
	}
L95:
	;
	m.G0 = v513 + int32(16)
	goto L93
L96:
	;
	if v517 == int32(4462524) {
		v977 = v515
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v534 = v517
	v538 = v515
	goto L98
L98:
	;
	v553 = int32(0)
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v534-int32(60))))
	if base.Ui32(v556) < base.Ui32(int32(2)) {
		v933 = v553
		goto L100
	} else {
		goto L101
	}
L99:
	;
	v977 = v956
	goto L95
L100:
	;
	v956 = F_add_size(m, v538, v933)
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L202
	}
L101:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v534-int32(32))))
	if v561 == int32(0) {
		v933 = v553
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v565 = v534 + int32(-64)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	if v566&int32(3) == int32(0) {
		v590 = v566
		goto L105
	} else {
		goto L106
	}
L103:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v534-int32(40))))
	switch v626 {
	case 0:
		goto L125
	case 1:
		goto L124
	case 2:
		goto L123
	case 3:
		goto L122
	case 4:
		goto L121
	default:
		v811 = v553
		goto L120
	}
L104:
	;
	v623 = v615 - v566
	goto L103
L105:
	;
	v594 = v590
	goto L114
L106:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v566))))
	if v574 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v623 = int32(0)
	goto L103
L108:
	;
	goto L109
L109:
	;
	v579 = v566
	goto L110
L110:
	;
	v583 = v579 + int32(1)
	if v583&int32(3) == int32(0) {
		v590 = v583
		goto L105
	} else {
		goto L112
	}
L111:
	;
	v615 = v583
	goto L104
L112:
	;
	v588 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v583))))
	if v588 != 0 {
		v579 = v583
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
	v603 = int32(-2139062144)
	if (int32(16843008)-v600|v600)&v603 == v603 {
		v594 = v594 + int32(4)
		goto L114
	} else {
		goto L116
	}
L115:
	;
	v609 = v594
	goto L117
L116:
	;
	goto L115
L117:
	;
	v613 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	if v613 != 0 {
		v609 = v609 + int32(1)
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v615 = v609
	goto L104
L119:
	;
	goto L118
L120:
	;
	v834 = int32(1)
	v838 = F_add_size(m, v623+v834, v811+v834)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L172
	}
L121:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v534)+28))
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v534)+36))
	if v703 == int32(0) {
		goto L94
	} else {
		goto L147
	}
L122:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v534)+28))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	if v641 == int32(0) {
		v811 = v553
		goto L120
	} else {
		goto L129
	}
L123:
	;
	v811 = int32(25)
	goto L120
L124:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v534)+28))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v630)))
	v633 = v631 >> (uint(int32(31)) % 32)
	if v631^v633-v633 < int32(1000) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v811 = int32(5)
	goto L120
L126:
	;
	v638 = int32(4)
	goto L128
L127:
	;
	v638 = int32(11)
	goto L128
L128:
	;
	v811 = v638
	goto L120
L129:
	;
	if v641&int32(3) == int32(0) {
		v667 = v641
		goto L132
	} else {
		goto L133
	}
L130:
	;
	v811 = v700
	goto L120
L131:
	;
	v700 = v692 - v641
	goto L130
L132:
	;
	v671 = v667
	goto L141
L133:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641))))
	if v651 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v700 = int32(0)
	goto L130
L135:
	;
	goto L136
L136:
	;
	v656 = v641
	goto L137
L137:
	;
	v660 = v656 + int32(1)
	if v660&int32(3) == int32(0) {
		v667 = v660
		goto L132
	} else {
		goto L139
	}
L138:
	;
	v692 = v660
	goto L131
L139:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v660))))
	if v665 != 0 {
		v656 = v660
		goto L137
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	v680 = int32(-2139062144)
	if (int32(16843008)-v677|v677)&v680 == v680 {
		v671 = v671 + int32(4)
		goto L141
	} else {
		goto L143
	}
L142:
	;
	v686 = v671
	goto L144
L143:
	;
	goto L142
L144:
	;
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686))))
	if v690 != 0 {
		v686 = v686 + int32(1)
		goto L144
	} else {
		goto L146
	}
L145:
	;
	v692 = v686
	goto L131
L146:
	;
	goto L145
L147:
	;
	v714 = v703
	goto L148
L148:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	if v737 == int32(0) {
		goto L94
	} else {
		goto L150
	}
L149:
	;
	if v737&int32(3) == int32(0) {
		v769 = v737
		goto L157
	} else {
		goto L158
	}
L150:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v714)+4))
	if v702 != v740 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v743 = v714 + int32(12)
	if v743 == int32(0) {
		goto L94
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	goto L149
L154:
	;
	v714 = v743
	goto L148
L155:
	;
	v811 = v802
	goto L120
L156:
	;
	v802 = v794 - v737
	goto L155
L157:
	;
	v773 = v769
	goto L166
L158:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737))))
	if v753 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v802 = int32(0)
	goto L155
L160:
	;
	goto L161
L161:
	;
	v758 = v737
	goto L162
L162:
	;
	v762 = v758 + int32(1)
	if v762&int32(3) == int32(0) {
		v769 = v762
		goto L157
	} else {
		goto L164
	}
L163:
	;
	v794 = v762
	goto L156
L164:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v762))))
	if v767 != 0 {
		v758 = v762
		goto L162
	} else {
		goto L165
	}
L165:
	;
	goto L163
L166:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v773)))
	v782 = int32(-2139062144)
	if (int32(16843008)-v779|v779)&v782 == v782 {
		v773 = v773 + int32(4)
		goto L166
	} else {
		goto L168
	}
L167:
	;
	v788 = v773
	goto L169
L168:
	;
	goto L167
L169:
	;
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v788))))
	if v792 != 0 {
		v788 = v788 + int32(1)
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v794 = v788
	goto L156
L171:
	;
	goto L170
L172:
	;
	v841 = v534 + int32(20)
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	if v842 != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	if v842&int32(3) == int32(0) {
		v866 = v842
		goto L178
	} else {
		goto L179
	}
L174:
	;
	v902 = v838
	goto L175
L175:
	;
	v904 = F_add_size(m, v902, int32(1))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L1
	} else {
		goto L194
	}
L176:
	;
	v900 = F_add_size(m, v838, v899)
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L1
	} else {
		goto L193
	}
L177:
	;
	v899 = v891 - v842
	goto L176
L178:
	;
	v870 = v866
	goto L187
L179:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842))))
	if v850 == int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v899 = int32(0)
	goto L176
L181:
	;
	goto L182
L182:
	;
	v855 = v842
	goto L183
L183:
	;
	v859 = v855 + int32(1)
	if v859&int32(3) == int32(0) {
		v866 = v859
		goto L178
	} else {
		goto L185
	}
L184:
	;
	v891 = v859
	goto L177
L185:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v859))))
	if v864 != 0 {
		v855 = v859
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v870)))
	v879 = int32(-2139062144)
	if (int32(16843008)-v876|v876)&v879 == v879 {
		v870 = v870 + int32(4)
		goto L187
	} else {
		goto L189
	}
L188:
	;
	v885 = v870
	goto L190
L189:
	;
	goto L188
L190:
	;
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885))))
	if v889 != 0 {
		v885 = v885 + int32(1)
		goto L190
	} else {
		goto L192
	}
L191:
	;
	v891 = v885
	goto L177
L192:
	;
	goto L191
L193:
	;
	v902 = v900
	goto L175
L194:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	if v906 == int32(0) {
		v915 = v904
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v917 = F_add_size(m, v915, int32(4))
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L199
	}
L196:
	;
	v909 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906))))
	if v909 == int32(0) {
		v915 = v904
		goto L195
	} else {
		goto L197
	}
L197:
	;
	v913 = F_add_size(m, v904, int32(4))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v915 = v913
	goto L195
L199:
	;
	v920 = F_add_size(m, v917, int32(4))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v923 = F_add_size(m, v920, int32(4))
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v933 = v923
	goto L100
L202:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	if v958 != int32(4462524) {
		v534 = v958
		v538 = v956
		goto L98
	} else {
		goto L203
	}
L203:
	;
	goto L99
L204:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(v565)))
	*(*int32)(unsafe.Add(mBase, uint32(v513)+4)) = v1030
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = v702
	F_errmsg_internal(m, int32(177846), v513)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L1
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(488529), int32(3036), int32(337089))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v1046
	v1052 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	v1053 = F_mul_size(m, int32(8), v1052)
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	v1055 = F_add_size(m, int32(4), v1053)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v1062 = F_add_size(m, v1057, (v1055+int32(31))&int32(-32))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v1062
	v1066 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if int32(2) <= v1066 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1069 = F_EstimateSnapshotSpace(m, v32)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	v1079 = v2
	goto L213
L213:
	;
	v1080 = F_EstimateSnapshotSpace(m, v36)
	mBase = m.M
	v1081 = m.ExcPending
	if v1081 != 0 {
		goto L1
	} else {
		goto L216
	}
L214:
	;
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v1076 = F_add_size(m, v1071, (v1069+int32(31))&int32(-32))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v1076
	v1079 = v1069
	goto L213
L216:
	;
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1087 = F_add_size(m, v1082, (v1080+int32(31))&int32(-32))
	mBase = m.M
	v1088 = m.ExcPending
	if v1088 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1087
	v1090 = int32(0)
	v1092 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	if v1092 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1094 = v1092
	v1109 = v1090
	goto L221
L219:
	;
	v1149 = v1090
	goto L220
L220:
	;
	v1166 = F_mul_size(m, int32(4), v1149)
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L1
	} else {
		goto L229
	}
L221:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1094)))
	if v1124 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	v1149 = v1130
	goto L220
L223:
	;
	v1126 = F_add_size(m, v1109, int32(1))
	mBase = m.M
	v1127 = m.ExcPending
	if v1127 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	v1128 = v1109
	goto L225
L225:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+52))
	v1130 = F_add_size(m, v1128, v1129)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L1
	} else {
		goto L227
	}
L226:
	;
	v1128 = v1126
	goto L225
L227:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1094)+80))
	if v1132 != 0 {
		v1094 = v1132
		v1109 = v1130
		goto L221
	} else {
		goto L228
	}
L228:
	;
	goto L222
L229:
	;
	v1168 = F_add_size(m, int32(32), v1166)
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1175 = F_add_size(m, v1170, (v1168+int32(31))&int32(-32))
	mBase = m.M
	v1176 = m.ExcPending
	if v1176 != 0 {
		goto L1
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1175
	v1179 = F_add_size(m, v1175, int32(32))
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1179
	v1183 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v1183 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1183)))
	v1186 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+4))
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+412))
	if v1187 != 0 {
		goto L237
	} else {
		goto L238
	}
L234:
	;
	v1256 = int32(1)
	goto L235
L235:
	;
	v1258 = F_mul_size(m, v1256, int32(12))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L1
	} else {
		goto L240
	}
L236:
	;
	v1256 = v1252 + int32(1)
	goto L235
L237:
	;
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+376))
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+364))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+352))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+340))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+328))
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+316))
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+304))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+292))
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+280))
	v1197 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+268))
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+256))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+244))
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+232))
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+220))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+208))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+196))
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+184))
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+172))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+160))
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+148))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+136))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+124))
	v1210 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+112))
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+100))
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+88))
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+76))
	v1216 = *(*int32)(unsafe.Add(mBase, uint32(v1185-int32(-64))))
	v1217 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+52))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+40))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+28))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1185)+16))
	v1252 = v1188 + (v1189 + (v1190 + (v1191 + (v1192 + (v1193 + (v1194 + (v1195 + (v1196 + (v1197 + (v1198 + (v1199 + (v1200 + (v1201 + (v1202 + (v1203 + (v1204 + (v1205 + (v1206 + (v1207 + (v1208 + (v1209 + (v1210 + (v1211 + (v1212 + (v1213 + (v1216 + (v1217 + (v1218 + (v1219 + (v1220 + v1186))))))))))))))))))))))))))))))
	goto L239
L238:
	;
	v1252 = v1186
	goto L239
L239:
	;
	goto L236
L240:
	;
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1265 = F_add_size(m, v1260, (v1258+int32(31))&int32(-32))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1265
	v1270 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	if v1270 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(v1270)+4))
	v1273 = v1271
	goto L244
L243:
	;
	v1273 = int32(0)
	goto L244
L244:
	;
	v1274 = F_mul_size(m, int32(4), v1273)
	mBase = m.M
	v1275 = m.ExcPending
	if v1275 != 0 {
		goto L1
	} else {
		goto L245
	}
L245:
	;
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1283 = F_add_size(m, v1278, (v1274+int32(43))&int32(-32))
	mBase = m.M
	v1284 = m.ExcPending
	if v1284 != 0 {
		goto L1
	} else {
		goto L246
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1283
	v1289 = F_add_size(m, v1283, int32(1056))
	mBase = m.M
	v1290 = m.ExcPending
	if v1290 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1289
	v1294 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v1294 != 0 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v1294)))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+4))
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+412))
	if v1298 != 0 {
		goto L252
	} else {
		goto L253
	}
L249:
	;
	v1364 = int32(0)
	goto L250
L250:
	;
	v1366 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	if v1366 != 0 {
		goto L255
	} else {
		goto L256
	}
L251:
	;
	v1364 = v1363
	goto L250
L252:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+376))
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+364))
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+352))
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+340))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+328))
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+316))
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+304))
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+292))
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+280))
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+268))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+256))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+244))
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+232))
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+220))
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+208))
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+196))
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+184))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+172))
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+160))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+148))
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+136))
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+124))
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+112))
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+100))
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+88))
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+76))
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v1296-int32(-64))))
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+52))
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+40))
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+28))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v1296)+16))
	v1363 = v1299 + (v1300 + (v1301 + (v1302 + (v1303 + (v1304 + (v1305 + (v1306 + (v1307 + (v1308 + (v1309 + (v1310 + (v1311 + (v1312 + (v1313 + (v1314 + (v1315 + (v1316 + (v1317 + (v1318 + (v1319 + (v1320 + (v1321 + (v1322 + (v1323 + (v1324 + (v1327 + (v1328 + (v1329 + (v1330 + (v1331 + v1297))))))))))))))))))))))))))))))
	goto L254
L253:
	;
	v1363 = v1297
	goto L254
L254:
	;
	goto L251
L255:
	;
	v1368 = *(*int32)(unsafe.Add(mBase, uint32(v1366)))
	v1369 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+4))
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+412))
	if v1370 != 0 {
		goto L259
	} else {
		goto L260
	}
L256:
	;
	v1437 = v1364
	goto L257
L257:
	;
	v1439 = v1437 << (uint(int32(2)) % 32)
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1447 = F_add_size(m, v1442, (v1439+int32(39))&int32(-32))
	mBase = m.M
	v1448 = m.ExcPending
	if v1448 != 0 {
		goto L1
	} else {
		goto L262
	}
L258:
	;
	v1437 = v1435 + v1364
	goto L257
L259:
	;
	v1371 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+376))
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+364))
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+352))
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+340))
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+328))
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+316))
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+304))
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+292))
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+280))
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+268))
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+256))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+244))
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+232))
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+220))
	v1385 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+208))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+196))
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+184))
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+172))
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+160))
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+148))
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+136))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+124))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+112))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+100))
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+88))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+76))
	v1399 = *(*int32)(unsafe.Add(mBase, uint32(v1368-int32(-64))))
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+52))
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+40))
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+28))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+16))
	v1435 = v1371 + (v1372 + (v1373 + (v1374 + (v1375 + (v1376 + (v1377 + (v1378 + (v1379 + (v1380 + (v1381 + (v1382 + (v1383 + (v1384 + (v1385 + (v1386 + (v1387 + (v1388 + (v1389 + (v1390 + (v1391 + (v1392 + (v1393 + (v1394 + (v1395 + (v1396 + (v1399 + (v1400 + (v1401 + (v1402 + (v1403 + v1369))))))))))))))))))))))))))))))
	goto L261
L260:
	;
	v1435 = v1369
	goto L261
L261:
	;
	goto L258
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1447
	v1452 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v1453 = m.ExcPending
	if v1453 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v1455 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	if v1455 != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	if v1455&int32(3) == int32(0) {
		v1479 = v1455
		goto L269
	} else {
		goto L270
	}
L265:
	;
	v1517 = v1452
	goto L266
L266:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1523 = F_add_size(m, v1518, (v1517+int32(31))&int32(-32))
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L1
	} else {
		goto L285
	}
L267:
	;
	v1515 = F_add_size(m, v1452, v1512+int32(1))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L1
	} else {
		goto L284
	}
L268:
	;
	v1512 = v1504 - v1455
	goto L267
L269:
	;
	v1483 = v1479
	goto L278
L270:
	;
	v1463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1455))))
	if v1463 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1512 = int32(0)
	goto L267
L272:
	;
	goto L273
L273:
	;
	v1468 = v1455
	goto L274
L274:
	;
	v1472 = v1468 + int32(1)
	if v1472&int32(3) == int32(0) {
		v1479 = v1472
		goto L269
	} else {
		goto L276
	}
L275:
	;
	v1504 = v1472
	goto L268
L276:
	;
	v1477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1472))))
	if v1477 != 0 {
		v1468 = v1472
		goto L274
	} else {
		goto L277
	}
L277:
	;
	goto L275
L278:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1483)))
	v1492 = int32(-2139062144)
	if (int32(16843008)-v1489|v1489)&v1492 == v1492 {
		v1483 = v1483 + int32(4)
		goto L278
	} else {
		goto L280
	}
L279:
	;
	v1498 = v1483
	goto L281
L280:
	;
	goto L279
L281:
	;
	v1502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498))))
	if v1502 != 0 {
		v1498 = v1498 + int32(1)
		goto L281
	} else {
		goto L283
	}
L282:
	;
	v1504 = v1498
	goto L268
L283:
	;
	goto L282
L284:
	;
	v1517 = v1515
	goto L266
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1523
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1528 = F_add_size(m, v1526, int32(12))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L1
	} else {
		goto L286
	}
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1528
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1534 = F_mul_size(m, int32(16384), v1533)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L1
	} else {
		goto L287
	}
L287:
	;
	v1540 = F_add_size(m, v1531, (v1534+int32(31))&int32(-32))
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1540
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1545 = F_add_size(m, v1543, int32(1))
	mBase = m.M
	v1546 = m.ExcPending
	if v1546 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1545
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v1549&int32(3) == int32(0) {
		v1573 = v1549
		goto L292
	} else {
		goto L293
	}
L290:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1607&int32(3) == int32(0) {
		v1631 = v1607
		goto L309
	} else {
		goto L310
	}
L291:
	;
	v1606 = v1598 - v1549
	goto L290
L292:
	;
	v1577 = v1573
	goto L301
L293:
	;
	v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1549))))
	if v1557 == int32(0) {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1606 = int32(0)
	goto L290
L295:
	;
	goto L296
L296:
	;
	v1562 = v1549
	goto L297
L297:
	;
	v1566 = v1562 + int32(1)
	if v1566&int32(3) == int32(0) {
		v1573 = v1566
		goto L292
	} else {
		goto L299
	}
L298:
	;
	v1598 = v1566
	goto L291
L299:
	;
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1566))))
	if v1571 != 0 {
		v1562 = v1566
		goto L297
	} else {
		goto L300
	}
L300:
	;
	goto L298
L301:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1577)))
	v1586 = int32(-2139062144)
	if (int32(16843008)-v1583|v1583)&v1586 == v1586 {
		v1577 = v1577 + int32(4)
		goto L301
	} else {
		goto L303
	}
L302:
	;
	v1592 = v1577
	goto L304
L303:
	;
	goto L302
L304:
	;
	v1596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1592))))
	if v1596 != 0 {
		v1592 = v1592 + int32(1)
		goto L304
	} else {
		goto L306
	}
L305:
	;
	v1598 = v1592
	goto L291
L306:
	;
	goto L305
L307:
	;
	v1670 = F_add_size(m, v1548, (v1606+v1664+int32(33))&int32(-32))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L1
	} else {
		goto L324
	}
L308:
	;
	v1664 = v1656 - v1607
	goto L307
L309:
	;
	v1635 = v1631
	goto L318
L310:
	;
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1607))))
	if v1615 == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1664 = int32(0)
	goto L307
L312:
	;
	goto L313
L313:
	;
	v1620 = v1607
	goto L314
L314:
	;
	v1624 = v1620 + int32(1)
	if v1624&int32(3) == int32(0) {
		v1631 = v1624
		goto L309
	} else {
		goto L316
	}
L315:
	;
	v1656 = v1624
	goto L308
L316:
	;
	v1629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1624))))
	if v1629 != 0 {
		v1620 = v1624
		goto L314
	} else {
		goto L317
	}
L317:
	;
	goto L315
L318:
	;
	v1641 = *(*int32)(unsafe.Add(mBase, uint32(v1635)))
	v1644 = int32(-2139062144)
	if (int32(16843008)-v1641|v1641)&v1644 == v1644 {
		v1635 = v1635 + int32(4)
		goto L318
	} else {
		goto L320
	}
L319:
	;
	v1650 = v1635
	goto L321
L320:
	;
	goto L319
L321:
	;
	v1654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1650))))
	if v1654 != 0 {
		v1650 = v1650 + int32(1)
		goto L321
	} else {
		goto L323
	}
L322:
	;
	v1656 = v1650
	goto L308
L323:
	;
	goto L322
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1670
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1675 = F_add_size(m, v1673, int32(1))
	mBase = m.M
	v1676 = m.ExcPending
	if v1676 != 0 {
		goto L1
	} else {
		goto L325
	}
L325:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1675
	v1682 = v977
	v1684 = v478
	v1686 = v1258
	v1688 = v68
	v1690 = v1274 + int32(12)
	v1691 = v1080
	v1692 = v1168
	v1693 = v1439 + int32(8)
	v1694 = int32(1048)
	v1695 = v362
	v1696 = v1517
	v1702 = v1055
	v1706 = v1079
	goto L6
L326:
	;
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v1688)))
	if v1711 <= int32(0) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	if v1719 != 0 {
		goto L333
	} else {
		goto L334
	}
L328:
	;
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1719 = v1714
	goto L327
L329:
	;
	goto L330
L330:
	;
	v1716 = F_dsm_create(m, v1709, int32(1))
	mBase = m.M
	v1717 = m.ExcPending
	if v1717 != 0 {
		goto L1
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1716
	v1719 = v1716
	goto L327
L332:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1728)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1728)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1728))) = int64(1346862204)
	*(*int32)(unsafe.Add(mBase, uint32(v1728)+12)) = v1709 & int32(-32)
	goto L337
L333:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(v1719)+24))
	v1728 = v1720
	goto L332
L334:
	;
	goto L335
L335:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1724 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1725 = F_MemoryContextAlloc(m, v1724, v1709)
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L1
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v1725
	v1728 = v1725
	goto L332
L337:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1728
	v1740 = F_shm_toc_allocate(m, v1728, int32(80))
	mBase = m.M
	v1741 = m.ExcPending
	if v1741 != 0 {
		goto L1
	} else {
		goto L338
	}
L338:
	;
	v1743 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740))) = v1743
	v1746 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+4)) = v1746
	v1749 = *(*int32)(unsafe.Add(mBase, _consts[121]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+8)) = v1749
	v1752 = *(*int32)(unsafe.Add(mBase, _consts[122]))
	v1755 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v1755 != 0 {
		goto L340
	} else {
		goto L341
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+12)) = v1756
	v1763 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740+int32(16)))) = v1763
	v1766 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740+int32(28)))) = v1766
	goto L343
L340:
	;
	v1756 = v1752
	goto L342
L341:
	;
	v1756 = int32(0)
	goto L342
L342:
	;
	goto L339
L343:
	;
	v1769 = int32(*(*uint8)(unsafe.Add(mBase, _consts[124])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1740)+32)) = uint8(v1769)
	v1772 = int32(*(*uint8)(unsafe.Add(mBase, _consts[125])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1740)+33)) = uint8(v1772)
	v1775 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+20)) = v1775
	v1778 = *(*int32)(unsafe.Add(mBase, _consts[127]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+24)) = v1778
	v1781 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+36)) = v1781
	v1784 = *(*int32)(unsafe.Add(mBase, _consts[129]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+40)) = v1784
	v1787 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+44)) = v1787
	v1790 = *(*int64)(unsafe.Add(mBase, _consts[130]))
	*(*int64)(unsafe.Add(mBase, uint32(v1740)+48)) = v1790
	v1793 = *(*int64)(unsafe.Add(mBase, _consts[131]))
	*(*int64)(unsafe.Add(mBase, uint32(v1740)+56)) = v1793
	v1796 = *(*int32)(unsafe.Add(mBase, _consts[132]))
	*(*int64)(unsafe.Add(mBase, uint32(v1740)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+68)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1740)+64)) = v1796
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1802, int64(-65535), v1740)
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L1
	} else {
		goto L344
	}
L344:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v1806 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1810 = F_shm_toc_allocate(m, v1809, v1684)
	mBase = m.M
	v1811 = m.ExcPending
	if v1811 != 0 {
		goto L1
	} else {
		goto L348
	}
L346:
	;
	v4039 = v1806
	goto L347
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v4039
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v38
	return
L348:
	;
	v1813 = *(*int32)(unsafe.Add(mBase, _consts[111]))
	if v1813 != 0 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1816 = v1813
	v1820 = v1684
	v1821 = v1810
	goto L352
L350:
	;
	v1973 = v1810
	goto L351
L351:
	;
	v1997 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1973))) = uint8(v1997)
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1999, int64(-65533), v1810)
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L1
	} else {
		goto L387
	}
L352:
	;
	v1846 = v1816 + int32(24)
	if v1820 == int32(0) {
		goto L356
	} else {
		goto L357
	}
L353:
	;
	v1973 = v1963
	goto L351
L354:
	;
	v1962 = v1958 + (v1955 - v1821) + int32(1)
	v1963 = v1821 + v1962
	v1965 = *(*int32)(unsafe.Add(mBase, uint32(v1816)))
	if v1965 != 0 {
		v1816 = v1965
		v1820 = v1820 - v1962
		v1821 = v1963
		goto L352
	} else {
		goto L386
	}
L355:
	;
	v1958 = F_strlen(m, v1954)
	mBase = m.M
	goto L354
L356:
	;
	v1954 = v1846
	v1955 = v1821
	goto L355
L357:
	;
	goto L358
L358:
	;
	v1852 = v1820 - int32(1)
	if (v1821^v1846)&int32(3) != 0 {
		goto L362
	} else {
		goto L363
	}
L359:
	;
	v1951 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1948))) = uint8(v1951)
	v1954 = v1947
	v1955 = v1948
	goto L355
L360:
	;
	v1932 = v1927
	v1933 = v1928
	v1934 = v1929
	goto L382
L361:
	;
	if v1922 == int32(0) {
		v1947 = v1920
		v1948 = v1921
		goto L359
	} else {
		goto L381
	}
L362:
	;
	v1920 = v1846
	v1921 = v1821
	v1922 = v1852
	goto L361
L363:
	;
	goto L364
L364:
	;
	v1856 = int32(0)
	if v1846&int32(3) == v1856 {
		goto L366
	} else {
		goto L367
	}
L365:
	;
	if v1889 == int32(0) {
		v1947 = v1886
		v1948 = v1887
		goto L359
	} else {
		goto L374
	}
L366:
	;
	v1886 = v1846
	v1887 = v1821
	v1888 = v1852
	v1889 = base.B2i32(v1852 != v1856)
	goto L365
L367:
	;
	if v1852 == int32(0) {
		goto L366
	} else {
		goto L368
	}
L368:
	;
	v1865 = v1846
	v1866 = v1821
	v1867 = v1852
	goto L369
L369:
	;
	v1869 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1865))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1866))) = uint8(v1869)
	if v1869 == int32(0) {
		v1927 = v1865
		v1928 = v1866
		v1929 = v1867
		goto L360
	} else {
		goto L371
	}
L370:
	;
	v1886 = v1880
	v1887 = v1874
	v1888 = v1876
	v1889 = v1878
	goto L365
L371:
	;
	v1873 = int32(1)
	v1874 = v1866 + v1873
	v1876 = v1867 - v1873
	v1877 = int32(0)
	v1878 = base.B2i32(v1876 != v1877)
	v1880 = v1865 + v1873
	if v1880&int32(3) == v1877 {
		v1886 = v1880
		v1887 = v1874
		v1888 = v1876
		v1889 = v1878
		goto L365
	} else {
		goto L372
	}
L372:
	;
	if v1876 != 0 {
		v1865 = v1880
		v1866 = v1874
		v1867 = v1876
		goto L369
	} else {
		goto L373
	}
L373:
	;
	goto L370
L374:
	;
	v1892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1886))))
	if v1892 == int32(0) {
		v1920 = v1886
		v1921 = v1887
		v1922 = v1888
		goto L361
	} else {
		goto L375
	}
L375:
	;
	if base.Ui32(v1888) < base.Ui32(int32(4)) {
		v1920 = v1886
		v1921 = v1887
		v1922 = v1888
		goto L361
	} else {
		goto L376
	}
L376:
	;
	v1898 = v1886
	v1899 = v1887
	v1900 = v1888
	goto L377
L377:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1898)))
	v1906 = int32(-2139062144)
	if (int32(16843008)-v1903|v1903)&v1906 != v1906 {
		v1927 = v1898
		v1928 = v1899
		v1929 = v1900
		goto L360
	} else {
		goto L379
	}
L378:
	;
	v1920 = v1914
	v1921 = v1912
	v1922 = v1916
	goto L361
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1899))) = v1903
	v1911 = int32(4)
	v1912 = v1899 + v1911
	v1914 = v1898 + v1911
	v1916 = v1900 - v1911
	if base.Ui32(int32(3)) < base.Ui32(v1916) {
		v1898 = v1914
		v1899 = v1912
		v1900 = v1916
		goto L377
	} else {
		goto L380
	}
L380:
	;
	goto L378
L381:
	;
	v1927 = v1920
	v1928 = v1921
	v1929 = v1922
	goto L360
L382:
	;
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1932))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1933))) = uint8(v1936)
	if v1936 == int32(0) {
		v1947 = v1932
		v1948 = v1933
		goto L359
	} else {
		goto L384
	}
L383:
	;
	v1947 = v1943
	v1948 = v1941
	goto L359
L384:
	;
	v1940 = int32(1)
	v1941 = v1933 + v1940
	v1943 = v1932 + v1940
	v1945 = v1934 - v1940
	if v1945 != 0 {
		v1932 = v1943
		v1933 = v1941
		v1934 = v1945
		goto L382
	} else {
		goto L385
	}
L385:
	;
	goto L383
L386:
	;
	goto L353
L387:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2004 = F_shm_toc_allocate(m, v2003, v1682)
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L1
	} else {
		goto L388
	}
L388:
	;
	v2006 = m.G0
	v2008 = v2006 - int32(112)
	m.G0 = v2008
	v2010 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+108)) = v2004 + v2010
	v2014 = v1682 - v2010
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+104)) = v2014
	v2017 = *(*int32)(unsafe.Add(mBase, _consts[112]))
	if v2017 == int32(0) {
		v2358 = v2014
		goto L394
	} else {
		goto L395
	}
L389:
	;
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2442, int64(-65532), v2004)
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L1
	} else {
		goto L459
	}
L390:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L1
	} else {
		goto L456
	}
L391:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2419 = m.ExcPending
	if v2419 != 0 {
		goto L1
	} else {
		goto L453
	}
L392:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L1
	} else {
		goto L450
	}
L393:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2393 = m.ExcPending
	if v2393 != 0 {
		goto L1
	} else {
		goto L447
	}
L394:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2004))) = v2014 - v2358
	m.G0 = v2008 + int32(112)
	goto L389
L395:
	;
	if v2017 == int32(4462524) {
		v2358 = v2014
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v2026 = v2014
	v2031 = v2017
	goto L397
L397:
	;
	v2055 = *(*int32)(unsafe.Add(mBase, uint32(v2031-int32(60))))
	if base.Ui32(v2055) < base.Ui32(int32(2)) {
		v2324 = v2026
		goto L399
	} else {
		goto L400
	}
L398:
	;
	v2358 = v2324
	goto L394
L399:
	;
	v2351 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+4))
	if v2351 != int32(4462524) {
		v2026 = v2324
		v2031 = v2351
		goto L397
	} else {
		goto L446
	}
L400:
	;
	v2059 = v2031 - int32(32)
	v2060 = *(*int32)(unsafe.Add(mBase, uint32(v2059)))
	if v2060 == int32(0) {
		v2324 = v2026
		goto L399
	} else {
		goto L401
	}
L401:
	;
	v2064 = v2031 + int32(-64)
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v2064)))
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+96)) = v2065
	F_do_serialize(m, v2008+int32(108), v2008+int32(104), int32(202172), v2008+int32(96))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L1
	} else {
		goto L402
	}
L402:
	;
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v2031-int32(40))))
	switch v2078 {
	case 0:
		goto L411
	case 1:
		goto L410
	case 2:
		goto L409
	case 3:
		goto L408
	case 4:
		goto L407
	default:
		goto L406
	}
L403:
	;
	if base.Ui32(v2294) <= base.Ui32(int32(3)) {
		goto L392
	} else {
		goto L443
	}
L404:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v2008)+104))
	if base.Ui32(v2281) <= base.Ui32(int32(3)) {
		goto L393
	} else {
		goto L442
	}
L405:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2267 = m.ExcPending
	if v2267 != 0 {
		goto L1
	} else {
		goto L439
	}
L406:
	;
	v2218 = v2031 + int32(20)
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v2218)))
	if v2219 != 0 {
		goto L431
	} else {
		goto L432
	}
L407:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+28))
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v2131)))
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+36))
	if v2133 == int32(0) {
		goto L405
	} else {
		goto L422
	}
L408:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+28))
	v2118 = *(*int32)(unsafe.Add(mBase, uint32(v2117)))
	if v2118 != 0 {
		goto L418
	} else {
		goto L419
	}
L409:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+28))
	v2104 = *(*float64)(unsafe.Add(mBase, uint32(v2103)))
	*(*float64)(unsafe.Add(mBase, uint32(v2008)+40)) = v2104
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+32)) = int32(17)
	F_do_serialize(m, v2008+int32(108), v2008+int32(104), int32(410637), v2008+int32(32))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L1
	} else {
		goto L417
	}
L410:
	;
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+28))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2091)))
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+16)) = v2092
	F_do_serialize(m, v2008+int32(108), v2008+int32(104), int32(477471), v2008+int32(16))
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L1
	} else {
		goto L416
	}
L411:
	;
	v2085 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+28))
	v2086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2085))))
	if v2086 != 0 {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v2087 = int32(336570)
	goto L414
L413:
	;
	v2087 = int32(353292)
	goto L414
L414:
	;
	F_do_serialize(m, v2008+int32(108), v2008+int32(104), v2087, int32(0))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	goto L406
L416:
	;
	goto L406
L417:
	;
	goto L406
L418:
	;
	v2120 = v2118
	goto L420
L419:
	;
	v2120 = int32(728204)
	goto L420
L420:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+48)) = v2120
	F_do_serialize(m, v2008+int32(108), v2008+int32(104), int32(202172), v2008+int32(48))
	mBase = m.M
	v2130 = m.ExcPending
	if v2130 != 0 {
		goto L1
	} else {
		goto L421
	}
L421:
	;
	goto L406
L422:
	;
	v2140 = v2133
	goto L423
L423:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v2140)))
	if v2167 == int32(0) {
		goto L405
	} else {
		goto L425
	}
L424:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+80)) = v2167
	F_do_serialize(m, v2008+int32(108), v2008+int32(104), int32(202172), v2008+int32(80))
	mBase = m.M
	v2185 = m.ExcPending
	if v2185 != 0 {
		goto L1
	} else {
		goto L430
	}
L425:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v2140)+4))
	if v2132 != v2170 {
		goto L426
	} else {
		goto L427
	}
L426:
	;
	v2173 = v2140 + int32(12)
	if v2173 == int32(0) {
		goto L405
	} else {
		goto L429
	}
L427:
	;
	goto L428
L428:
	;
	goto L424
L429:
	;
	v2140 = v2173
	goto L423
L430:
	;
	goto L406
L431:
	;
	v2221 = v2219
	goto L433
L432:
	;
	v2221 = int32(728204)
	goto L433
L433:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2008))) = v2221
	F_do_serialize(m, v2008+int32(108), v2008+int32(104), int32(202172), v2008)
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L1
	} else {
		goto L434
	}
L434:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, uint32(v2218)))
	if v2230 != 0 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v2231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2230))))
	if v2231 != 0 {
		goto L404
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	v2232 = *(*int32)(unsafe.Add(mBase, uint32(v2008)+104))
	v2294 = v2232
	goto L403
L438:
	;
	goto L437
L439:
	;
	v2268 = *(*int32)(unsafe.Add(mBase, uint32(v2064)))
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+68)) = v2268
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+64)) = v2132
	F_errmsg_internal(m, int32(177846), v2008-int32(-64))
	mBase = m.M
	v2275 = m.ExcPending
	if v2275 != 0 {
		goto L1
	} else {
		goto L440
	}
L440:
	;
	F_errfinish(m, int32(488529), int32(3036), int32(337089))
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L442:
	;
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v2008)+108))
	v2285 = *(*int32)(unsafe.Add(mBase, uint32(v2031)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2284))) = v2285
	v2287 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+108)) = v2284 + v2287
	v2294 = v2281 - v2287
	goto L403
L443:
	;
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2008)+108))
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(v2059)))
	*(*int32)(unsafe.Add(mBase, uint32(v2297))) = v2298
	v2301 = v2294 & int32(-4)
	if v2301 == int32(4) {
		goto L391
	} else {
		goto L444
	}
L444:
	;
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v2031-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v2297)+4)) = v2306
	if v2301 == int32(8) {
		goto L390
	} else {
		goto L445
	}
L445:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(v2031-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v2297)+8)) = v2312
	v2314 = int32(12)
	v2315 = v2294 - v2314
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+104)) = v2315
	*(*int32)(unsafe.Add(mBase, uint32(v2008)+108)) = v2297 + v2314
	v2324 = v2315
	goto L399
L446:
	;
	goto L398
L447:
	;
	F_errmsg_internal(m, int32(345630), int32(0))
	mBase = m.M
	v2397 = m.ExcPending
	if v2397 != 0 {
		goto L1
	} else {
		goto L448
	}
L448:
	;
	F_errfinish(m, int32(488529), int32(6020), int32(17297))
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L450:
	;
	F_errmsg_internal(m, int32(345630), int32(0))
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	F_errfinish(m, int32(488529), int32(6020), int32(17297))
	mBase = m.M
	v2415 = m.ExcPending
	if v2415 != 0 {
		goto L1
	} else {
		goto L452
	}
L452:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L453:
	;
	F_errmsg_internal(m, int32(345630), int32(0))
	mBase = m.M
	v2423 = m.ExcPending
	if v2423 != 0 {
		goto L1
	} else {
		goto L454
	}
L454:
	;
	F_errfinish(m, int32(488529), int32(6020), int32(17297))
	mBase = m.M
	v2428 = m.ExcPending
	if v2428 != 0 {
		goto L1
	} else {
		goto L455
	}
L455:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L456:
	;
	F_errmsg_internal(m, int32(345630), int32(0))
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L1
	} else {
		goto L457
	}
L457:
	;
	F_errfinish(m, int32(488529), int32(6020), int32(17297))
	mBase = m.M
	v2441 = m.ExcPending
	if v2441 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L459:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2447 = F_shm_toc_allocate(m, v2446, v1702)
	mBase = m.M
	v2448 = m.ExcPending
	if v2448 != 0 {
		goto L1
	} else {
		goto L460
	}
L460:
	;
	v2450 = *(*int32)(unsafe.Add(mBase, _consts[113]))
	*(*int32)(unsafe.Add(mBase, uint32(v2447))) = v2450
	v2453 = v2447 + int32(4)
	v2455 = v2450 << (uint(int32(3)) % 32)
	v2456 = v2453 + v2455
	if base.Ui32(v2456) < base.Ui32(v2447) {
		goto L462
	} else {
		goto L463
	}
L461:
	;
	v2479 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2479, int64(-65531), v2447)
	mBase = m.M
	v2482 = m.ExcPending
	if v2482 != 0 {
		goto L1
	} else {
		goto L475
	}
L462:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2469 = m.ExcPending
	if v2469 != 0 {
		goto L1
	} else {
		goto L472
	}
L463:
	;
	if base.Ui32(v2447+v1702) < base.Ui32(v2456) {
		goto L462
	} else {
		goto L464
	}
L464:
	;
	if int32(0) < v2450 {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v2463 = *(*int32)(unsafe.Add(mBase, _consts[133]))
	if v2455 != 0 {
		goto L469
	} else {
		goto L470
	}
L466:
	;
	goto L467
L467:
	;
	goto L461
L468:
	;
	goto L467
L469:
	;
	v2464 = F__emscripten_memcpy_bulkmem(m, v2453, v2463, v2455)
	mBase = m.M
	goto L471
L470:
	;
	goto L471
L471:
	;
	goto L468
L472:
	;
	F_errmsg_internal(m, int32(345541), int32(0))
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L1
	} else {
		goto L473
	}
L473:
	;
	F_errfinish(m, int32(488489), int32(327), int32(346292))
	mBase = m.M
	v2478 = m.ExcPending
	if v2478 != 0 {
		goto L1
	} else {
		goto L474
	}
L474:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L475:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	if int32(2) <= v2484 {
		goto L476
	} else {
		goto L477
	}
L476:
	;
	v2487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2488 = F_shm_toc_allocate(m, v2487, v1706)
	mBase = m.M
	v2489 = m.ExcPending
	if v2489 != 0 {
		goto L1
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v2536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2537 = F_shm_toc_allocate(m, v2536, v1691)
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L1
	} else {
		goto L494
	}
L479:
	;
	v2496 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v2497 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v2498 = *(*int64)(unsafe.Add(mBase, uint32(v32)+4))
	v2499 = *(*int32)(unsafe.Add(mBase, uint32(v32)+32))
	v2500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+29)))
	v2501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2488)+16)) = uint8(v2501)
	*(*uint8)(unsafe.Add(mBase, uint32(v2488)+17)) = uint8(v2500)
	*(*int32)(unsafe.Add(mBase, uint32(v2488)+20)) = v2499
	*(*int64)(unsafe.Add(mBase, uint32(v2488))) = v2498
	*(*int32)(unsafe.Add(mBase, uint32(v2488)+8)) = v2497
	if v2500 != 0 {
		goto L481
	} else {
		goto L482
	}
L480:
	;
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2531, int64(-65530), v2488)
	mBase = m.M
	v2534 = m.ExcPending
	if v2534 != 0 {
		goto L1
	} else {
		goto L493
	}
L481:
	;
	v2508 = v2496
	goto L483
L482:
	;
	v2508 = int32(0)
	goto L483
L483:
	;
	if v2501 != 0 {
		goto L484
	} else {
		goto L485
	}
L484:
	;
	v2509 = v2508
	goto L486
L485:
	;
	v2509 = v2496
	goto L486
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2488)+12)) = v2509
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	if v2511 != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v2517 = F___memcpy(m, v2488+int32(24), v2514, v2511<<(uint(int32(2))%32))
	mBase = m.M
	goto L489
L488:
	;
	goto L489
L489:
	;
	if int32(0) < v2509 {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v2521 = int32(2)
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(v32)+20))
	v2527 = *(*int32)(unsafe.Add(mBase, uint32(v32)+24))
	v2530 = F___memcpy(m, v2488+v2520<<(uint(v2521)%32)+int32(24), v2526, v2527<<(uint(v2521)%32))
	mBase = m.M
	goto L492
L491:
	;
	goto L492
L492:
	;
	goto L480
L493:
	;
	goto L478
L494:
	;
	v2545 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v2546 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v2547 = *(*int64)(unsafe.Add(mBase, uint32(v36)+4))
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v2549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+29)))
	v2550 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2537)+16)) = uint8(v2550)
	*(*uint8)(unsafe.Add(mBase, uint32(v2537)+17)) = uint8(v2549)
	*(*int32)(unsafe.Add(mBase, uint32(v2537)+20)) = v2548
	*(*int64)(unsafe.Add(mBase, uint32(v2537))) = v2547
	*(*int32)(unsafe.Add(mBase, uint32(v2537)+8)) = v2546
	if v2549 != 0 {
		goto L496
	} else {
		goto L497
	}
L495:
	;
	v2580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2580, int64(-65529), v2537)
	mBase = m.M
	v2583 = m.ExcPending
	if v2583 != 0 {
		goto L1
	} else {
		goto L508
	}
L496:
	;
	v2557 = v2545
	goto L498
L497:
	;
	v2557 = int32(0)
	goto L498
L498:
	;
	if v2550 != 0 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v2558 = v2557
	goto L501
L500:
	;
	v2558 = v2545
	goto L501
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2537)+12)) = v2558
	v2560 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v2560 != 0 {
		goto L502
	} else {
		goto L503
	}
L502:
	;
	v2563 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v2566 = F___memcpy(m, v2537+int32(24), v2563, v2560<<(uint(int32(2))%32))
	mBase = m.M
	goto L504
L503:
	;
	goto L504
L504:
	;
	if int32(0) < v2558 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v2569 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	v2570 = int32(2)
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v2576 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	v2579 = F___memcpy(m, v2537+v2569<<(uint(v2570)%32)+int32(24), v2575, v2576<<(uint(v2570)%32))
	mBase = m.M
	goto L507
L506:
	;
	goto L507
L507:
	;
	goto L495
L508:
	;
	v2584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2586 = F_shm_toc_allocate(m, v2584, int32(4))
	mBase = m.M
	v2587 = m.ExcPending
	if v2587 != 0 {
		goto L1
	} else {
		goto L509
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2586))) = v1695
	v2589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2589, int64(-65526), v2586)
	mBase = m.M
	v2592 = m.ExcPending
	if v2592 != 0 {
		goto L1
	} else {
		goto L510
	}
L510:
	;
	v2593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2594 = F_shm_toc_allocate(m, v2593, v1692)
	mBase = m.M
	v2595 = m.ExcPending
	if v2595 != 0 {
		goto L1
	} else {
		goto L511
	}
L511:
	;
	v2596 = int32(0)
	v2598 = *(*int32)(unsafe.Add(mBase, _consts[114]))
	*(*int32)(unsafe.Add(mBase, uint32(v2594))) = v2598
	v2601 = int32(*(*uint8)(unsafe.Add(mBase, _consts[134])))
	*(*uint8)(unsafe.Add(mBase, uint32(v2594)+4)) = uint8(v2601)
	v2604 = *(*int64)(unsafe.Add(mBase, _consts[80]))
	*(*int64)(unsafe.Add(mBase, uint32(v2594)+8)) = v2604
	v2607 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	v2608 = *(*int64)(unsafe.Add(mBase, uint32(v2607)))
	*(*int64)(unsafe.Add(mBase, uint32(v2594)+16)) = v2608
	v2611 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	*(*int32)(unsafe.Add(mBase, uint32(v2594)+24)) = v2611
	v2614 = *(*int32)(unsafe.Add(mBase, _consts[136]))
	if v2596 < v2614 {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	v2798 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2798, int64(-65528), v2594)
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L1
	} else {
		goto L550
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2594)+28)) = v2614
	v2621 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	v2623 = v2614 << (uint(int32(2)) % 32)
	if v2623 != 0 {
		goto L517
	} else {
		goto L518
	}
L514:
	;
	goto L515
L515:
	;
	v2628 = v2596
	v2639 = v2607
	goto L520
L516:
	;
	goto L512
L517:
	;
	v2624 = F__emscripten_memcpy_bulkmem(m, v2594+int32(32), v2621, v2623)
	mBase = m.M
	goto L519
L518:
	;
	goto L519
L519:
	;
	goto L516
L520:
	;
	v2657 = *(*int32)(unsafe.Add(mBase, uint32(v2639)))
	if v2657 != 0 {
		goto L522
	} else {
		goto L523
	}
L521:
	;
	v2668 = v2663 << (uint(int32(2)) % 32)
	v2669 = F_palloc(m, v2668)
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L1
	} else {
		goto L528
	}
L522:
	;
	v2659 = F_add_size(m, v2628, int32(1))
	mBase = m.M
	v2660 = m.ExcPending
	if v2660 != 0 {
		goto L1
	} else {
		goto L525
	}
L523:
	;
	v2661 = v2628
	goto L524
L524:
	;
	v2662 = *(*int32)(unsafe.Add(mBase, uint32(v2639)+52))
	v2663 = F_add_size(m, v2661, v2662)
	mBase = m.M
	v2664 = m.ExcPending
	if v2664 != 0 {
		goto L1
	} else {
		goto L526
	}
L525:
	;
	v2661 = v2659
	goto L524
L526:
	;
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2639)+80))
	if v2665 != 0 {
		v2628 = v2663
		v2639 = v2665
		goto L520
	} else {
		goto L527
	}
L527:
	;
	goto L521
L528:
	;
	v2672 = *(*int32)(unsafe.Add(mBase, _consts[72]))
	if v2672 != 0 {
		goto L529
	} else {
		goto L530
	}
L529:
	;
	v2680 = int32(0)
	v2686 = v2672
	goto L532
L530:
	;
	goto L531
L531:
	;
	F_pg_qsort(m, v2669, v2663, int32(4), int32(185))
	mBase = m.M
	v2761 = m.ExcPending
	if v2761 != 0 {
		goto L1
	} else {
		goto L545
	}
L532:
	;
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v2686)))
	if v2704 != 0 {
		goto L534
	} else {
		goto L535
	}
L533:
	;
	goto L531
L534:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2669+v2680<<(uint(int32(2))%32)))) = v2704
	v2711 = v2680 + int32(1)
	goto L536
L535:
	;
	v2711 = v2680
	goto L536
L536:
	;
	v2712 = *(*int32)(unsafe.Add(mBase, uint32(v2686)+52))
	if int32(0) < v2712 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v2715 = int32(2)
	v2718 = *(*int32)(unsafe.Add(mBase, uint32(v2686)+48))
	v2720 = v2712 << (uint(v2715) % 32)
	if v2720 != 0 {
		goto L541
	} else {
		goto L542
	}
L538:
	;
	v2724 = v2712
	goto L539
L539:
	;
	v2726 = *(*int32)(unsafe.Add(mBase, uint32(v2686)+80))
	if v2726 != 0 {
		v2680 = v2724 + v2711
		v2686 = v2726
		goto L532
	} else {
		goto L544
	}
L540:
	;
	v2723 = *(*int32)(unsafe.Add(mBase, uint32(v2686)+52))
	v2724 = v2723
	goto L539
L541:
	;
	v2721 = F__emscripten_memcpy_bulkmem(m, v2669+v2711<<(uint(v2715)%32), v2718, v2720)
	mBase = m.M
	goto L543
L542:
	;
	goto L543
L543:
	;
	goto L540
L544:
	;
	goto L533
L545:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2594)+28)) = v2663
	if v2668 != 0 {
		goto L547
	} else {
		goto L548
	}
L546:
	;
	goto L512
L547:
	;
	v2765 = F__emscripten_memcpy_bulkmem(m, v2594+int32(32), v2669, v2668)
	mBase = m.M
	goto L549
L548:
	;
	goto L549
L549:
	;
	goto L546
L550:
	;
	v2802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2803 = F_shm_toc_allocate(m, v2802, v1686)
	mBase = m.M
	v2804 = m.ExcPending
	if v2804 != 0 {
		goto L1
	} else {
		goto L551
	}
L551:
	;
	v2805 = m.G0
	v2807 = v2805 - int32(80)
	m.G0 = v2807
	v2810 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	if v2810 != 0 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2807)+48)) = int64(51539607564)
	v2814 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v2807)+72)) = v2814
	v2818 = *(*int32)(unsafe.Add(mBase, uint32(v2810)))
	v2819 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+4))
	v2820 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+412))
	if v2820 != 0 {
		goto L556
	} else {
		goto L557
	}
L553:
	;
	v3126 = v2803
	goto L554
L554:
	;
	if v3126&int32(3) != 0 {
		goto L591
	} else {
		goto L592
	}
L555:
	;
	v2889 = F_hash_create(m, int32(128373), v2885, v2807+int32(32), int32(1064))
	mBase = m.M
	v2890 = m.ExcPending
	if v2890 != 0 {
		goto L1
	} else {
		goto L559
	}
L556:
	;
	v2821 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+376))
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+364))
	v2823 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+352))
	v2824 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+340))
	v2825 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+328))
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+316))
	v2827 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+304))
	v2828 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+292))
	v2829 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+280))
	v2830 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+268))
	v2831 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+256))
	v2832 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+244))
	v2833 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+232))
	v2834 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+220))
	v2835 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+208))
	v2836 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+196))
	v2837 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+184))
	v2838 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+172))
	v2839 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+160))
	v2840 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+148))
	v2841 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+136))
	v2842 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+124))
	v2843 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+112))
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+100))
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+88))
	v2846 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+76))
	v2849 = *(*int32)(unsafe.Add(mBase, uint32(v2818-int32(-64))))
	v2850 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+52))
	v2851 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+40))
	v2852 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+28))
	v2853 = *(*int32)(unsafe.Add(mBase, uint32(v2818)+16))
	v2885 = v2821 + (v2822 + (v2823 + (v2824 + (v2825 + (v2826 + (v2827 + (v2828 + (v2829 + (v2830 + (v2831 + (v2832 + (v2833 + (v2834 + (v2835 + (v2836 + (v2837 + (v2838 + (v2839 + (v2840 + (v2841 + (v2842 + (v2843 + (v2844 + (v2845 + (v2846 + (v2849 + (v2850 + (v2851 + (v2852 + (v2853 + v2819))))))))))))))))))))))))))))))
	goto L558
L557:
	;
	v2885 = v2819
	goto L558
L558:
	;
	goto L555
L559:
	;
	v2894 = *(*int32)(unsafe.Add(mBase, _consts[115]))
	F_hash_seq_init(m, v2807+int32(12), v2894)
	mBase = m.M
	v2896 = m.ExcPending
	if v2896 != 0 {
		goto L1
	} else {
		goto L560
	}
L560:
	;
	v2899 = F_hash_seq_search(m, v2807+int32(12))
	mBase = m.M
	v2900 = m.ExcPending
	if v2900 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	if v2899 != 0 {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	v2903 = v2899
	goto L565
L563:
	;
	goto L564
L564:
	;
	v2972 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	if v2972 != 0 {
		goto L570
	} else {
		goto L571
	}
L565:
	;
	v2934 = F_hash_search(m, v2889, v2903, int32(1), int32(0))
	mBase = m.M
	v2935 = m.ExcPending
	if v2935 != 0 {
		goto L1
	} else {
		goto L567
	}
L566:
	;
	goto L564
L567:
	;
	v2938 = F_hash_seq_search(m, v2807+int32(12))
	mBase = m.M
	v2939 = m.ExcPending
	if v2939 != 0 {
		goto L1
	} else {
		goto L568
	}
L568:
	;
	if v2938 != 0 {
		v2903 = v2938
		goto L565
	} else {
		goto L569
	}
L569:
	;
	goto L566
L570:
	;
	v2975 = v2972
	goto L573
L571:
	;
	goto L572
L572:
	;
	F_hash_seq_init(m, v2807+int32(12), v2889)
	mBase = m.M
	v3046 = m.ExcPending
	if v3046 != 0 {
		goto L1
	} else {
		goto L580
	}
L573:
	;
	v3004 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2975)+16)))
	if v3004 == int32(1) {
		goto L575
	} else {
		goto L576
	}
L574:
	;
	goto L572
L575:
	;
	v3009 = F_hash_search(m, v2889, v2975, int32(2), int32(0))
	mBase = m.M
	v3010 = m.ExcPending
	if v3010 != 0 {
		goto L1
	} else {
		goto L578
	}
L576:
	;
	goto L577
L577:
	;
	v3011 = *(*int32)(unsafe.Add(mBase, uint32(v2975)+24))
	if v3011 != 0 {
		v2975 = v3011
		goto L573
	} else {
		goto L579
	}
L578:
	;
	goto L577
L579:
	;
	goto L574
L580:
	;
	v3049 = F_hash_seq_search(m, v2807+int32(12))
	mBase = m.M
	v3050 = m.ExcPending
	if v3050 != 0 {
		goto L1
	} else {
		goto L581
	}
L581:
	;
	if v3049 != 0 {
		goto L582
	} else {
		goto L583
	}
L582:
	;
	v3052 = v2803
	v3053 = v3049
	goto L585
L583:
	;
	v3093 = v2803
	goto L584
L584:
	;
	F_hash_destroy(m, v2889)
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		goto L1
	} else {
		goto L589
	}
L585:
	;
	v3082 = *(*int64)(unsafe.Add(mBase, uint32(v3053)))
	*(*int64)(unsafe.Add(mBase, uint32(v3052))) = v3082
	v3084 = *(*int32)(unsafe.Add(mBase, uint32(v3053)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3052)+8)) = v3084
	v3086 = int32(12)
	v3087 = v3052 + v3086
	v3090 = F_hash_seq_search(m, v2807+v3086)
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L1
	} else {
		goto L587
	}
L586:
	;
	v3093 = v3087
	goto L584
L587:
	;
	if v3090 != 0 {
		v3052 = v3087
		v3053 = v3090
		goto L585
	} else {
		goto L588
	}
L588:
	;
	goto L586
L589:
	;
	v3126 = v3093
	goto L554
L590:
	;
	m.G0 = v2807 + int32(80)
	v3184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3184, int64(-65525), v2803)
	mBase = m.M
	v3187 = m.ExcPending
	if v3187 != 0 {
		goto L1
	} else {
		goto L599
	}
L591:
	;
	v3176 = int32(12)
	goto L593
L592:
	;
	v3161 = v3126 + int32(12)
	if base.Ui32(v3161) <= base.Ui32(v3126) {
		goto L590
	} else {
		goto L594
	}
L593:
	;
	v3178 = F__emscripten_memset_bulkmem(m, v3126, base.I32_extend8_s(int32(0)), v3176)
	mBase = m.M
	goto L598
L594:
	;
	v3166 = v3126 + int32(4)
	if base.Ui32(v3166) < base.Ui32(v3161) {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v3168 = v3161
	goto L597
L596:
	;
	v3168 = v3166
	goto L597
L597:
	;
	v3176 = (v3126^int32(-1)+v3168)&int32(-4) + int32(4)
	goto L593
L598:
	;
	goto L590
L599:
	;
	v3188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3189 = F_shm_toc_allocate(m, v3188, v1690)
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L1
	} else {
		goto L600
	}
L600:
	;
	v3191 = int32(0)
	v3193 = *(*int32)(unsafe.Add(mBase, _consts[139]))
	*(*int32)(unsafe.Add(mBase, uint32(v3189))) = v3193
	v3196 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	*(*int32)(unsafe.Add(mBase, uint32(v3189)+4)) = v3196
	v3199 = *(*int32)(unsafe.Add(mBase, _consts[47]))
	if v3199 == v3191 {
		goto L602
	} else {
		goto L603
	}
L601:
	;
	v3315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3315, int64(-65524), v3189)
	mBase = m.M
	v3318 = m.ExcPending
	if v3318 != 0 {
		goto L1
	} else {
		goto L611
	}
L602:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3189)+8)) = int32(0)
	goto L601
L603:
	;
	goto L604
L604:
	;
	v3204 = *(*int32)(unsafe.Add(mBase, uint32(v3199)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v3189)+8)) = v3204
	v3206 = *(*int32)(unsafe.Add(mBase, uint32(v3199)+4))
	if int32(0) < v3206 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	v3217 = v3191
	goto L608
L606:
	;
	goto L607
L607:
	;
	goto L601
L608:
	;
	v3243 = v3217 << (uint(int32(2)) % 32)
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(v3199)+12))
	v3247 = *(*int32)(unsafe.Add(mBase, uint32(v3245+v3243)))
	*(*int32)(unsafe.Add(mBase, uint32(v3189+int32(12)+v3243))) = v3247
	v3250 = v3217 + int32(1)
	v3251 = *(*int32)(unsafe.Add(mBase, uint32(v3199)+4))
	if v3250 < v3251 {
		v3217 = v3250
		goto L608
	} else {
		goto L610
	}
L609:
	;
	goto L607
L610:
	;
	goto L609
L611:
	;
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3320 = F_shm_toc_allocate(m, v3319, v1694)
	mBase = m.M
	v3321 = m.ExcPending
	if v3321 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	goto L614
L613:
	;
	v3326 = int32(524)
	goto L618
L614:
	;
	v3324 = F__emscripten_memcpy_bulkmem(m, v3320, int32(4452768), int32(524))
	mBase = m.M
	goto L616
L616:
	;
	goto L613
L617:
	;
	v3332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3332, int64(-65523), v3320)
	mBase = m.M
	v3335 = m.ExcPending
	if v3335 != 0 {
		goto L1
	} else {
		goto L621
	}
L618:
	;
	v3330 = F__emscripten_memcpy_bulkmem(m, v3324+v3326, int32(4453816), v3326)
	mBase = m.M
	goto L620
L620:
	;
	goto L617
L621:
	;
	v3336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3337 = F_shm_toc_allocate(m, v3336, v1693)
	mBase = m.M
	v3338 = m.ExcPending
	if v3338 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	v3339 = m.G0
	v3341 = v3339 - int32(32)
	m.G0 = v3341
	v3344 = *(*int32)(unsafe.Add(mBase, _consts[116]))
	if v3344 == int32(0) {
		v3397 = v3337
		goto L623
	} else {
		goto L624
	}
L623:
	;
	v3427 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3397))) = v3427
	v3430 = v3397 + int32(4)
	v3432 = *(*int32)(unsafe.Add(mBase, _consts[117]))
	if v3432 == v3427 {
		v3485 = v3430
		goto L632
	} else {
		goto L633
	}
L624:
	;
	F_hash_seq_init(m, v3341+int32(12), v3344)
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	v3353 = F_hash_seq_search(m, v3341+int32(12))
	mBase = m.M
	v3354 = m.ExcPending
	if v3354 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	if v3353 == int32(0) {
		v3397 = v3337
		goto L623
	} else {
		goto L627
	}
L627:
	;
	v3358 = v3337
	v3363 = v3353
	goto L628
L628:
	;
	v3388 = *(*int32)(unsafe.Add(mBase, uint32(v3363)))
	*(*int32)(unsafe.Add(mBase, uint32(v3358))) = v3388
	v3391 = v3358 + int32(4)
	v3394 = F_hash_seq_search(m, v3341+int32(12))
	mBase = m.M
	v3395 = m.ExcPending
	if v3395 != 0 {
		goto L1
	} else {
		goto L630
	}
L629:
	;
	v3397 = v3391
	goto L623
L630:
	;
	if v3394 != 0 {
		v3358 = v3391
		v3363 = v3394
		goto L628
	} else {
		goto L631
	}
L631:
	;
	goto L629
L632:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3485))) = int32(0)
	m.G0 = v3341 + int32(32)
	v3520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3520, int64(-65522), v3337)
	mBase = m.M
	v3523 = m.ExcPending
	if v3523 != 0 {
		goto L1
	} else {
		goto L641
	}
L633:
	;
	F_hash_seq_init(m, v3341+int32(12), v3432)
	mBase = m.M
	v3438 = m.ExcPending
	if v3438 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	v3441 = F_hash_seq_search(m, v3341+int32(12))
	mBase = m.M
	v3442 = m.ExcPending
	if v3442 != 0 {
		goto L1
	} else {
		goto L635
	}
L635:
	;
	if v3441 == int32(0) {
		v3485 = v3430
		goto L632
	} else {
		goto L636
	}
L636:
	;
	v3446 = v3430
	v3451 = v3441
	goto L637
L637:
	;
	v3476 = *(*int32)(unsafe.Add(mBase, uint32(v3451)))
	*(*int32)(unsafe.Add(mBase, uint32(v3446))) = v3476
	v3479 = v3446 + int32(4)
	v3482 = F_hash_seq_search(m, v3341+int32(12))
	mBase = m.M
	v3483 = m.ExcPending
	if v3483 != 0 {
		goto L1
	} else {
		goto L639
	}
L638:
	;
	v3485 = v3479
	goto L632
L639:
	;
	if v3482 != 0 {
		v3446 = v3479
		v3451 = v3482
		goto L637
	} else {
		goto L640
	}
L640:
	;
	goto L638
L641:
	;
	v3524 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3525 = F_shm_toc_allocate(m, v3524, v1696)
	mBase = m.M
	v3526 = m.ExcPending
	if v3526 != 0 {
		goto L1
	} else {
		goto L642
	}
L642:
	;
	v3528 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v3530 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	if v3530 == int32(0) {
		goto L644
	} else {
		goto L645
	}
L643:
	;
	v3606 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3606, int64(-65521), v3525)
	mBase = m.M
	v3609 = m.ExcPending
	if v3609 != 0 {
		goto L1
	} else {
		goto L671
	}
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3525)+4)) = v3528
	*(*int32)(unsafe.Add(mBase, uint32(v3525))) = int32(-1)
	goto L643
L645:
	;
	goto L646
L646:
	;
	if v3530&int32(3) == int32(0) {
		v3559 = v3530
		goto L649
	} else {
		goto L650
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3525)+4)) = v3528
	*(*int32)(unsafe.Add(mBase, uint32(v3525))) = v3592
	if int32(0) <= v3592 {
		goto L664
	} else {
		goto L665
	}
L648:
	;
	v3592 = v3584 - v3530
	goto L647
L649:
	;
	v3563 = v3559
	goto L658
L650:
	;
	v3543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3530))))
	if v3543 == int32(0) {
		goto L651
	} else {
		goto L652
	}
L651:
	;
	v3592 = int32(0)
	goto L647
L652:
	;
	goto L653
L653:
	;
	v3548 = v3530
	goto L654
L654:
	;
	v3552 = v3548 + int32(1)
	if v3552&int32(3) == int32(0) {
		v3559 = v3552
		goto L649
	} else {
		goto L656
	}
L655:
	;
	v3584 = v3552
	goto L648
L656:
	;
	v3557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3552))))
	if v3557 != 0 {
		v3548 = v3552
		goto L654
	} else {
		goto L657
	}
L657:
	;
	goto L655
L658:
	;
	v3569 = *(*int32)(unsafe.Add(mBase, uint32(v3563)))
	v3572 = int32(-2139062144)
	if (int32(16843008)-v3569|v3569)&v3572 == v3572 {
		v3563 = v3563 + int32(4)
		goto L658
	} else {
		goto L660
	}
L659:
	;
	v3578 = v3563
	goto L661
L660:
	;
	goto L659
L661:
	;
	v3582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3578))))
	if v3582 != 0 {
		v3578 = v3578 + int32(1)
		goto L661
	} else {
		goto L663
	}
L662:
	;
	v3584 = v3578
	goto L648
L663:
	;
	goto L662
L664:
	;
	v3600 = *(*int32)(unsafe.Add(mBase, _consts[118]))
	v3602 = v3592 + int32(1)
	if v3602 != 0 {
		goto L668
	} else {
		goto L669
	}
L665:
	;
	goto L666
L666:
	;
	goto L643
L667:
	;
	goto L666
L668:
	;
	v3603 = F__emscripten_memcpy_bulkmem(m, v3525+int32(8), v3600, v3602)
	mBase = m.M
	goto L670
L669:
	;
	goto L670
L670:
	;
	goto L667
L671:
	;
	v3610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3613 = F_palloc0(m, v3610<<(uint(int32(3))%32))
	mBase = m.M
	v3614 = m.ExcPending
	if v3614 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v3613
	v3616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3618 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3619 = F_mul_size(m, int32(16384), v3618)
	mBase = m.M
	v3620 = m.ExcPending
	if v3620 != 0 {
		goto L1
	} else {
		goto L673
	}
L673:
	;
	v3621 = F_shm_toc_allocate(m, v3616, v3619)
	mBase = m.M
	v3622 = m.ExcPending
	if v3622 != 0 {
		goto L1
	} else {
		goto L674
	}
L674:
	;
	v3623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v3623 {
		goto L675
	} else {
		goto L676
	}
L675:
	;
	v3630 = int32(0)
	goto L678
L676:
	;
	goto L677
L677:
	;
	v3724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3724, int64(-65534), v3621)
	mBase = m.M
	v3727 = m.ExcPending
	if v3727 != 0 {
		goto L1
	} else {
		goto L684
	}
L678:
	;
	v3660 = v3621 + v3630<<(uint(int32(14))%32)
	v3662 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3660)+16)) = v3662
	*(*int32)(unsafe.Add(mBase, uint32(v3660)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3660))) = v3662
	v3668 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v3660)+36)) = uint16(v3668)
	*(*int64)(unsafe.Add(mBase, uint32(v3660)+24)) = v3662
	*(*int32)(unsafe.Add(mBase, uint32(v3660)+32)) = int32(16344)
	goto L680
L679:
	;
	goto L677
L680:
	;
	v3678 = *(*int32)(unsafe.Add(mBase, _consts[128]))
	F_shm_mq_set_receiver(m, v3660, v3678)
	mBase = m.M
	v3680 = m.ExcPending
	if v3680 != 0 {
		goto L1
	} else {
		goto L681
	}
L681:
	;
	v3681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v3682 = F_shm_mq_attach(m, v3660, v3681)
	mBase = m.M
	v3683 = m.ExcPending
	if v3683 != 0 {
		goto L1
	} else {
		goto L682
	}
L682:
	;
	v3684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3684+v3630<<(uint(int32(3))%32))+4)) = v3682
	v3690 = v3630 + int32(1)
	v3691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3690 < v3691 {
		v3630 = v3690
		goto L678
	} else {
		goto L683
	}
L683:
	;
	goto L679
L684:
	;
	v3728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3728&int32(3) == int32(0) {
		v3752 = v3728
		goto L687
	} else {
		goto L688
	}
L685:
	;
	v3786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3787&int32(3) == int32(0) {
		v3811 = v3787
		goto L704
	} else {
		goto L705
	}
L686:
	;
	v3785 = v3777 - v3728
	goto L685
L687:
	;
	v3756 = v3752
	goto L696
L688:
	;
	v3736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3728))))
	if v3736 == int32(0) {
		goto L689
	} else {
		goto L690
	}
L689:
	;
	v3785 = int32(0)
	goto L685
L690:
	;
	goto L691
L691:
	;
	v3741 = v3728
	goto L692
L692:
	;
	v3745 = v3741 + int32(1)
	if v3745&int32(3) == int32(0) {
		v3752 = v3745
		goto L687
	} else {
		goto L694
	}
L693:
	;
	v3777 = v3745
	goto L686
L694:
	;
	v3750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3745))))
	if v3750 != 0 {
		v3741 = v3745
		goto L692
	} else {
		goto L695
	}
L695:
	;
	goto L693
L696:
	;
	v3762 = *(*int32)(unsafe.Add(mBase, uint32(v3756)))
	v3765 = int32(-2139062144)
	if (int32(16843008)-v3762|v3762)&v3765 == v3765 {
		v3756 = v3756 + int32(4)
		goto L696
	} else {
		goto L698
	}
L697:
	;
	v3771 = v3756
	goto L699
L698:
	;
	goto L697
L699:
	;
	v3775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3771))))
	if v3775 != 0 {
		v3771 = v3771 + int32(1)
		goto L699
	} else {
		goto L701
	}
L700:
	;
	v3777 = v3771
	goto L686
L701:
	;
	goto L700
L702:
	;
	v3848 = F_shm_toc_allocate(m, v3786, v3844+v3785+int32(2))
	mBase = m.M
	v3849 = m.ExcPending
	if v3849 != 0 {
		goto L1
	} else {
		goto L719
	}
L703:
	;
	v3844 = v3836 - v3787
	goto L702
L704:
	;
	v3815 = v3811
	goto L713
L705:
	;
	v3795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3787))))
	if v3795 == int32(0) {
		goto L706
	} else {
		goto L707
	}
L706:
	;
	v3844 = int32(0)
	goto L702
L707:
	;
	goto L708
L708:
	;
	v3800 = v3787
	goto L709
L709:
	;
	v3804 = v3800 + int32(1)
	if v3804&int32(3) == int32(0) {
		v3811 = v3804
		goto L704
	} else {
		goto L711
	}
L710:
	;
	v3836 = v3804
	goto L703
L711:
	;
	v3809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3804))))
	if v3809 != 0 {
		v3800 = v3804
		goto L709
	} else {
		goto L712
	}
L712:
	;
	goto L710
L713:
	;
	v3821 = *(*int32)(unsafe.Add(mBase, uint32(v3815)))
	v3824 = int32(-2139062144)
	if (int32(16843008)-v3821|v3821)&v3824 == v3824 {
		v3815 = v3815 + int32(4)
		goto L713
	} else {
		goto L715
	}
L714:
	;
	v3830 = v3815
	goto L716
L715:
	;
	goto L714
L716:
	;
	v3834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3830))))
	if v3834 != 0 {
		v3830 = v3830 + int32(1)
		goto L716
	} else {
		goto L718
	}
L717:
	;
	v3836 = v3830
	goto L703
L718:
	;
	goto L717
L719:
	;
	v3850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if (v3850^v3848)&int32(3) != 0 {
		goto L723
	} else {
		goto L724
	}
L720:
	;
	v3927 = v3785 + v3848 + int32(1)
	v3928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if (v3928^v3927)&int32(3) != 0 {
		goto L744
	} else {
		goto L745
	}
L721:
	;
	goto L720
L722:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3905))) = uint8(v3904)
	if v3904&int32(255) == int32(0) {
		goto L721
	} else {
		goto L737
	}
L723:
	;
	v3856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3850))))
	v3903 = v3850
	v3904 = v3856
	v3905 = v3848
	goto L722
L724:
	;
	goto L725
L725:
	;
	if v3850&int32(3) != 0 {
		goto L726
	} else {
		goto L727
	}
L726:
	;
	v3860 = v3850
	v3862 = v3848
	goto L729
L727:
	;
	v3874 = v3850
	v3876 = v3848
	goto L728
L728:
	;
	v3878 = *(*int32)(unsafe.Add(mBase, uint32(v3874)))
	v3881 = int32(-2139062144)
	if (int32(16843008)-v3878|v3878)&v3881 != v3881 {
		v3903 = v3874
		v3904 = v3878
		v3905 = v3876
		goto L722
	} else {
		goto L733
	}
L729:
	;
	v3863 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3860))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3862))) = uint8(v3863)
	if v3863 == int32(0) {
		goto L721
	} else {
		goto L731
	}
L730:
	;
	v3874 = v3870
	v3876 = v3868
	goto L728
L731:
	;
	v3867 = int32(1)
	v3868 = v3862 + v3867
	v3870 = v3860 + v3867
	if v3870&int32(3) != 0 {
		v3860 = v3870
		v3862 = v3868
		goto L729
	} else {
		goto L732
	}
L732:
	;
	goto L730
L733:
	;
	v3886 = v3874
	v3887 = v3878
	v3888 = v3876
	goto L734
L734:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3888))) = v3887
	v3890 = int32(4)
	v3891 = v3888 + v3890
	v3892 = *(*int32)(unsafe.Add(mBase, uint32(v3886)+4))
	v3894 = v3886 + v3890
	v3898 = int32(-2139062144)
	if (v3892|(int32(16843008)-v3892))&v3898 == v3898 {
		v3886 = v3894
		v3887 = v3892
		v3888 = v3891
		goto L734
	} else {
		goto L736
	}
L735:
	;
	v3903 = v3894
	v3904 = v3892
	v3905 = v3891
	goto L722
L736:
	;
	goto L735
L737:
	;
	v3912 = v3903
	v3914 = v3905
	goto L738
L738:
	;
	v3915 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3912)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3914)+1)) = uint8(v3915)
	v3917 = int32(1)
	if v3915 != 0 {
		v3912 = v3912 + v3917
		v3914 = v3914 + v3917
		goto L738
	} else {
		goto L740
	}
L739:
	;
	goto L721
L740:
	;
	goto L739
L741:
	;
	v4003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v4003, int64(-65527), v3848)
	mBase = m.M
	v4006 = m.ExcPending
	if v4006 != 0 {
		goto L1
	} else {
		goto L762
	}
L742:
	;
	goto L741
L743:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3983))) = uint8(v3982)
	if v3982&int32(255) == int32(0) {
		goto L742
	} else {
		goto L758
	}
L744:
	;
	v3934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3928))))
	v3981 = v3928
	v3982 = v3934
	v3983 = v3927
	goto L743
L745:
	;
	goto L746
L746:
	;
	if v3928&int32(3) != 0 {
		goto L747
	} else {
		goto L748
	}
L747:
	;
	v3938 = v3928
	v3940 = v3927
	goto L750
L748:
	;
	v3952 = v3928
	v3954 = v3927
	goto L749
L749:
	;
	v3956 = *(*int32)(unsafe.Add(mBase, uint32(v3952)))
	v3959 = int32(-2139062144)
	if (int32(16843008)-v3956|v3956)&v3959 != v3959 {
		v3981 = v3952
		v3982 = v3956
		v3983 = v3954
		goto L743
	} else {
		goto L754
	}
L750:
	;
	v3941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3938))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3940))) = uint8(v3941)
	if v3941 == int32(0) {
		goto L742
	} else {
		goto L752
	}
L751:
	;
	v3952 = v3948
	v3954 = v3946
	goto L749
L752:
	;
	v3945 = int32(1)
	v3946 = v3940 + v3945
	v3948 = v3938 + v3945
	if v3948&int32(3) != 0 {
		v3938 = v3948
		v3940 = v3946
		goto L750
	} else {
		goto L753
	}
L753:
	;
	goto L751
L754:
	;
	v3964 = v3952
	v3965 = v3956
	v3966 = v3954
	goto L755
L755:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3966))) = v3965
	v3968 = int32(4)
	v3969 = v3966 + v3968
	v3970 = *(*int32)(unsafe.Add(mBase, uint32(v3964)+4))
	v3972 = v3964 + v3968
	v3976 = int32(-2139062144)
	if (v3970|(int32(16843008)-v3970))&v3976 == v3976 {
		v3964 = v3972
		v3965 = v3970
		v3966 = v3969
		goto L755
	} else {
		goto L757
	}
L756:
	;
	v3981 = v3972
	v3982 = v3970
	v3983 = v3969
	goto L743
L757:
	;
	goto L756
L758:
	;
	v3990 = v3981
	v3992 = v3983
	goto L759
L759:
	;
	v3993 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3990)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3992)+1)) = uint8(v3993)
	v3995 = int32(1)
	if v3993 != 0 {
		v3990 = v3990 + v3995
		v3992 = v3992 + v3995
		goto L759
	} else {
		goto L761
	}
L760:
	;
	goto L742
L761:
	;
	goto L760
L762:
	;
	v4007 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4039 = v4007
	goto L347
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
				F_errmsg_internal(m, int32(664340), v10)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(486277), int32(904), int32(27386))
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
