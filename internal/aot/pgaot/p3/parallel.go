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
	var v191 int32
	_ = v191
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
	v191 = int32(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBuckets[0]))
	if v269 != 0 {
		goto L53
	} else {
		goto L54
	}
L38:
	;
	v200 = v191 + (v177 + v184)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v209 = v201 + v202&(v203-int32(1))<<(uint(int32(2))%32)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v210
	v212 = v191 + (v166 + v184)
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
	v254 = (v249+int32(15))&int32(-8) + v191
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	if base.Ui32(v254) < base.Ui32(v255) {
		v191 = v254
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
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
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
				*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = int32(0)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				if v64 != 0 {
					base.MemoryFill(m, v61+int32(20), int32(0), v64)
				} else {
				}
			}
			v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
			mBase = m.M
			v194 = m.ExcPending
			if v194 != 0 {
				return int32(0)
			} else {
				return v193
			}
		default:
			v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
			mBase = m.M
			v194 = m.ExcPending
			if v194 != 0 {
				return int32(0)
			} else {
				return v193
			}
		case 6:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+36)))
			if v16 != int32(1) {
				v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					return v193
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
					v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						return v193
					}
				}
			}
		case 8:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+36)))
			if v29 != int32(1) {
				v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					return v193
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				F_index_parallelrescan(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						return v193
					}
				}
			}
		case 9:
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+36)))
			if v36 != int32(1) {
				v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					return v193
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
				F_index_parallelrescan(m, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						return v193
					}
				}
			}
		case 11:
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+36)))
			if v85 != int32(1) {
				v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					return v193
				}
			} else {
				v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+172))
				if v89 != 0 {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
					*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = int32(0)
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
					if v93 != 0 {
						v94 = F_dsa_get_address(m, v89, v93)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
							if v96 == int32(0) {
								v111 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
								if v111 == int32(0) {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
									if v126 == int32(0) {
										F_dsa_free(m, v89, v93)
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
											v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												return v193
											}
										}
									} else {
										v129 = F_dsa_get_address(m, v89, v126)
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int32(0)
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
											v132 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
											if v131 != v132 {
												F_dsa_free(m, v89, v93)
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
													v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														return v193
													}
												}
											} else {
												v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
												F_dsa_free(m, v89, v137)
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return int32(0)
												} else {
													F_dsa_free(m, v89, v93)
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
														v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int32(0)
														} else {
															return v193
														}
													}
												}
											}
										}
									}
								} else {
									v114 = F_dsa_get_address(m, v89, v111)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
										v117 = int32(1)
										*(*int32)(unsafe.Add(mBase, uint32(v114))) = v116 - v117
										if v116 != v117 {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
											if v126 == int32(0) {
												F_dsa_free(m, v89, v93)
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
													v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														return v193
													}
												}
											} else {
												v129 = F_dsa_get_address(m, v89, v126)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return int32(0)
												} else {
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
													v132 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
													if v131 != v132 {
														F_dsa_free(m, v89, v93)
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
															v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return int32(0)
															} else {
																return v193
															}
														}
													} else {
														v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														F_dsa_free(m, v89, v137)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return int32(0)
														} else {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v194 = m.ExcPending
																if v194 != 0 {
																	return int32(0)
																} else {
																	return v193
																}
															}
														}
													}
												}
											}
										} else {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
											F_dsa_free(m, v89, v122)
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return int32(0)
											} else {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
												if v126 == int32(0) {
													F_dsa_free(m, v89, v93)
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
														v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int32(0)
														} else {
															return v193
														}
													}
												} else {
													v129 = F_dsa_get_address(m, v89, v126)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
														v132 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
														if v131 != v132 {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v194 = m.ExcPending
																if v194 != 0 {
																	return int32(0)
																} else {
																	return v193
																}
															}
														} else {
															v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															F_dsa_free(m, v89, v137)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return int32(0)
															} else {
																F_dsa_free(m, v89, v93)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																	v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v194 = m.ExcPending
																	if v194 != 0 {
																		return int32(0)
																	} else {
																		return v193
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
								v99 = F_dsa_get_address(m, v89, v96)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
									v102 = int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v99))) = v101 - v102
									if v101 != v102 {
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
										if v111 == int32(0) {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
											if v126 == int32(0) {
												F_dsa_free(m, v89, v93)
												mBase = m.M
												v142 = m.ExcPending
												if v142 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
													v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														return v193
													}
												}
											} else {
												v129 = F_dsa_get_address(m, v89, v126)
												mBase = m.M
												v130 = m.ExcPending
												if v130 != 0 {
													return int32(0)
												} else {
													v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
													v132 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
													if v131 != v132 {
														F_dsa_free(m, v89, v93)
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
															v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return int32(0)
															} else {
																return v193
															}
														}
													} else {
														v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														F_dsa_free(m, v89, v137)
														mBase = m.M
														v139 = m.ExcPending
														if v139 != 0 {
															return int32(0)
														} else {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v194 = m.ExcPending
																if v194 != 0 {
																	return int32(0)
																} else {
																	return v193
																}
															}
														}
													}
												}
											}
										} else {
											v114 = F_dsa_get_address(m, v89, v111)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
												v117 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v114))) = v116 - v117
												if v116 != v117 {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
													if v126 == int32(0) {
														F_dsa_free(m, v89, v93)
														mBase = m.M
														v142 = m.ExcPending
														if v142 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
															v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v194 = m.ExcPending
															if v194 != 0 {
																return int32(0)
															} else {
																return v193
															}
														}
													} else {
														v129 = F_dsa_get_address(m, v89, v126)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int32(0)
														} else {
															v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
															v132 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
															if v131 != v132 {
																F_dsa_free(m, v89, v93)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																	v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v194 = m.ExcPending
																	if v194 != 0 {
																		return int32(0)
																	} else {
																		return v193
																	}
																}
															} else {
																v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																F_dsa_free(m, v89, v137)
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return int32(0)
																} else {
																	F_dsa_free(m, v89, v93)
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																		v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v194 = m.ExcPending
																		if v194 != 0 {
																			return int32(0)
																		} else {
																			return v193
																		}
																	}
																}
															}
														}
													}
												} else {
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
													F_dsa_free(m, v89, v122)
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return int32(0)
													} else {
														v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														if v126 == int32(0) {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v194 = m.ExcPending
																if v194 != 0 {
																	return int32(0)
																} else {
																	return v193
																}
															}
														} else {
															v129 = F_dsa_get_address(m, v89, v126)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return int32(0)
															} else {
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
																v132 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
																if v131 != v132 {
																	F_dsa_free(m, v89, v93)
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																		v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v194 = m.ExcPending
																		if v194 != 0 {
																			return int32(0)
																		} else {
																			return v193
																		}
																	}
																} else {
																	v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																	F_dsa_free(m, v89, v137)
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return int32(0)
																	} else {
																		F_dsa_free(m, v89, v93)
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																			v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v194 = m.ExcPending
																			if v194 != 0 {
																				return int32(0)
																			} else {
																				return v193
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
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
										F_dsa_free(m, v89, v107)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
											if v111 == int32(0) {
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
												if v126 == int32(0) {
													F_dsa_free(m, v89, v93)
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
														v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int32(0)
														} else {
															return v193
														}
													}
												} else {
													v129 = F_dsa_get_address(m, v89, v126)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
														v132 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
														if v131 != v132 {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v194 = m.ExcPending
																if v194 != 0 {
																	return int32(0)
																} else {
																	return v193
																}
															}
														} else {
															v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															F_dsa_free(m, v89, v137)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return int32(0)
															} else {
																F_dsa_free(m, v89, v93)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																	v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v194 = m.ExcPending
																	if v194 != 0 {
																		return int32(0)
																	} else {
																		return v193
																	}
																}
															}
														}
													}
												}
											} else {
												v114 = F_dsa_get_address(m, v89, v111)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return int32(0)
												} else {
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
													v117 = int32(1)
													*(*int32)(unsafe.Add(mBase, uint32(v114))) = v116 - v117
													if v116 != v117 {
														v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														if v126 == int32(0) {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v142 = m.ExcPending
															if v142 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v194 = m.ExcPending
																if v194 != 0 {
																	return int32(0)
																} else {
																	return v193
																}
															}
														} else {
															v129 = F_dsa_get_address(m, v89, v126)
															mBase = m.M
															v130 = m.ExcPending
															if v130 != 0 {
																return int32(0)
															} else {
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
																v132 = int32(1)
																*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
																if v131 != v132 {
																	F_dsa_free(m, v89, v93)
																	mBase = m.M
																	v142 = m.ExcPending
																	if v142 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																		v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v194 = m.ExcPending
																		if v194 != 0 {
																			return int32(0)
																		} else {
																			return v193
																		}
																	}
																} else {
																	v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																	F_dsa_free(m, v89, v137)
																	mBase = m.M
																	v139 = m.ExcPending
																	if v139 != 0 {
																		return int32(0)
																	} else {
																		F_dsa_free(m, v89, v93)
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																			v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v194 = m.ExcPending
																			if v194 != 0 {
																				return int32(0)
																			} else {
																				return v193
																			}
																		}
																	}
																}
															}
														}
													} else {
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
														F_dsa_free(m, v89, v122)
														mBase = m.M
														v124 = m.ExcPending
														if v124 != 0 {
															return int32(0)
														} else {
															v126 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															if v126 == int32(0) {
																F_dsa_free(m, v89, v93)
																mBase = m.M
																v142 = m.ExcPending
																if v142 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																	v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v194 = m.ExcPending
																	if v194 != 0 {
																		return int32(0)
																	} else {
																		return v193
																	}
																}
															} else {
																v129 = F_dsa_get_address(m, v89, v126)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return int32(0)
																} else {
																	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
																	v132 = int32(1)
																	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v131 - v132
																	if v131 != v132 {
																		F_dsa_free(m, v89, v93)
																		mBase = m.M
																		v142 = m.ExcPending
																		if v142 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																			v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v194 = m.ExcPending
																			if v194 != 0 {
																				return int32(0)
																			} else {
																				return v193
																			}
																		}
																	} else {
																		v137 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																		F_dsa_free(m, v89, v137)
																		mBase = m.M
																		v139 = m.ExcPending
																		if v139 != 0 {
																			return int32(0)
																		} else {
																			F_dsa_free(m, v89, v93)
																			mBase = m.M
																			v142 = m.ExcPending
																			if v142 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																				v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																				mBase = m.M
																				v194 = m.ExcPending
																				if v194 != 0 {
																					return int32(0)
																				} else {
																					return v193
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
						*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
						v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
						mBase = m.M
						v194 = m.ExcPending
						if v194 != 0 {
							return int32(0)
						} else {
							return v193
						}
					}
				} else {
					v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						return v193
					}
				}
			}
		case 21:
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+36)))
			if v43 != int32(1) {
				v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					return v193
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
							v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return int32(0)
							} else {
								return v193
							}
						}
					}
				} else {
					v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						return v193
					}
				}
			}
		case 22:
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+36)))
			if v70 != int32(1) {
				v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					return v193
				}
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+36))
				if v74 != 0 {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v77 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+40)))
					v79 = F_shm_toc_lookup(m, v75, v77, int32(0))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v73)+36))
						m.T0[v81].(func(*base.Module, int32, int32, int32))(m, l0, l1, v79)
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return int32(0)
							} else {
								return v193
							}
						}
					}
				} else {
					v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						return v193
					}
				}
			}
		case 26:
			v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+36)))
			if v152 != int32(1) {
				v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v194 = m.ExcPending
				if v194 != 0 {
					return int32(0)
				} else {
					return v193
				}
			} else {
				v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				if v155 != 0 {
					v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
					v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v158 = int64(*(*int32)(unsafe.Add(mBase, uint32(v157)+40)))
					v160 = F_shm_toc_lookup(m, v156, v158, int32(0))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return int32(0)
					} else {
						v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						if v162 != 0 {
							F_ExecHashTableDetachBatch(m, v162)
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return int32(0)
							} else {
								v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								F_ExecHashTableDetach(m, v165)
								mBase = m.M
								v167 = m.ExcPending
								if v167 != 0 {
									return int32(0)
								} else {
									F_FileSetDeleteAll(m, v160+int32(168))
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return int32(0)
									} else {
										v173 = v160 + int32(56)
										v174 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v174
										v176 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v173))) = v176
										*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = v176
										*(*uint8)(unsafe.Add(mBase, uint32(v173)+20)) = uint8(v174)
										F_ConditionVariableInit(m, v160+int32(80))
										mBase = m.M
										v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											return v193
										}
									}
								}
							}
						} else {
							F_FileSetDeleteAll(m, v160+int32(168))
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return int32(0)
							} else {
								v173 = v160 + int32(56)
								v174 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v174
								v176 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v173))) = v176
								*(*int64)(unsafe.Add(mBase, uint32(v173)+12)) = v176
								*(*uint8)(unsafe.Add(mBase, uint32(v173)+20)) = uint8(v174)
								F_ConditionVariableInit(m, v160+int32(80))
								mBase = m.M
								v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									return v193
								}
							}
						}
					}
				} else {
					v193 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						return v193
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
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
	var v258 int32
	_ = v258
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v603 int32
	_ = v603
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v656 int32
	_ = v656
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v836 int32
	_ = v836
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v843 int32
	_ = v843
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
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v866 int32
	_ = v866
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1061 int32
	_ = v1061
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
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
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1238 int32
	_ = v1238
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1285 int32
	_ = v1285
	var v1295 int32
	_ = v1295
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
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1324 int32
	_ = v1324
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1374 int32
	_ = v1374
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1386 int64
	_ = v1386
	var v1389 int64
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1443 int32
	_ = v1443
	var v1449 int32
	_ = v1449
	var v1453 int32
	_ = v1453
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1570 int32
	_ = v1570
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1612 int32
	_ = v1612
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1661 int32
	_ = v1661
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1681 int32
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 float64
	_ = v1710
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1779 int32
	_ = v1779
	var v1782 int32
	_ = v1782
	var v1792 int32
	_ = v1792
	var v1827 int32
	_ = v1827
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1870 int32
	_ = v1870
	var v1871 int32
	_ = v1871
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1903 int32
	_ = v1903
	var v1909 int32
	_ = v1909
	var v1911 int32
	_ = v1911
	var v1912 int32
	_ = v1912
	var v1918 int32
	_ = v1918
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v2024 int32
	_ = v2024
	var v2025 int32
	_ = v2025
	var v2032 int32
	_ = v2032
	var v2037 int32
	_ = v2037
	var v2038 int32
	_ = v2038
	var v2041 int32
	_ = v2041
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2049 int32
	_ = v2049
	var v2053 int32
	_ = v2053
	var v2061 int32
	_ = v2061
	var v2066 int32
	_ = v2066
	var v2070 int32
	_ = v2070
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2096 int64
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2110 int32
	_ = v2110
	var v2114 int32
	_ = v2114
	var v2119 int32
	_ = v2119
	var v2124 int32
	_ = v2124
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2135 int32
	_ = v2135
	var v2138 int32
	_ = v2138
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2152 int32
	_ = v2152
	var v2153 int32
	_ = v2153
	var v2154 int32
	_ = v2154
	var v2155 int64
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2178 int32
	_ = v2178
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2188 int32
	_ = v2188
	var v2194 int32
	_ = v2194
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2201 int32
	_ = v2201
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2221 int64
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2225 int64
	_ = v2225
	var v2228 int32
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2236 int32
	_ = v2236
	var v2242 int32
	_ = v2242
	var v2248 int32
	_ = v2248
	var v2250 int32
	_ = v2250
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2324 int32
	_ = v2324
	var v2331 int32
	_ = v2331
	var v2332 int32
	_ = v2332
	var v2336 int32
	_ = v2336
	var v2340 int32
	_ = v2340
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2382 int32
	_ = v2382
	var v2421 int32
	_ = v2421
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2430 int32
	_ = v2430
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2448 int32
	_ = v2448
	var v2449 int32
	_ = v2449
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2457 int32
	_ = v2457
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2474 int32
	_ = v2474
	var v2506 int32
	_ = v2506
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2513 int32
	_ = v2513
	var v2515 int32
	_ = v2515
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2524 int32
	_ = v2524
	var v2554 int32
	_ = v2554
	var v2555 int32
	_ = v2555
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2593 int32
	_ = v2593
	var v2598 int32
	_ = v2598
	var v2626 int32
	_ = v2626
	var v2631 int32
	_ = v2631
	var v2632 int32
	_ = v2632
	var v2633 int32
	_ = v2633
	var v2667 int32
	_ = v2667
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2704 int32
	_ = v2704
	var v2706 int64
	_ = v2706
	var v2708 int32
	_ = v2708
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2713 int32
	_ = v2713
	var v2717 int32
	_ = v2717
	var v2747 int32
	_ = v2747
	var v2751 int32
	_ = v2751
	var v2787 int32
	_ = v2787
	var v2790 int32
	_ = v2790
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2803 int32
	_ = v2803
	var v2805 int32
	_ = v2805
	var v2813 int32
	_ = v2813
	var v2843 int32
	_ = v2843
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2850 int32
	_ = v2850
	var v2851 int32
	_ = v2851
	var v2887 int32
	_ = v2887
	var v2890 int32
	_ = v2890
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2893 int32
	_ = v2893
	var v2895 int32
	_ = v2895
	var v2902 int32
	_ = v2902
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2907 int32
	_ = v2907
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2911 int32
	_ = v2911
	var v2914 int32
	_ = v2914
	var v2918 int32
	_ = v2918
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2926 int32
	_ = v2926
	var v2929 int32
	_ = v2929
	var v2957 int32
	_ = v2957
	var v2960 int32
	_ = v2960
	var v2963 int32
	_ = v2963
	var v2964 int32
	_ = v2964
	var v2969 int32
	_ = v2969
	var v2997 int32
	_ = v2997
	var v3000 int32
	_ = v3000
	var v3002 int32
	_ = v3002
	var v3006 int32
	_ = v3006
	var v3008 int32
	_ = v3008
	var v3009 int32
	_ = v3009
	var v3010 int32
	_ = v3010
	var v3014 int32
	_ = v3014
	var v3017 int32
	_ = v3017
	var v3045 int32
	_ = v3045
	var v3048 int32
	_ = v3048
	var v3051 int32
	_ = v3051
	var v3052 int32
	_ = v3052
	var v3057 int32
	_ = v3057
	var v3090 int32
	_ = v3090
	var v3093 int32
	_ = v3093
	var v3094 int32
	_ = v3094
	var v3095 int32
	_ = v3095
	var v3096 int32
	_ = v3096
	var v3098 int32
	_ = v3098
	var v3100 int32
	_ = v3100
	var v3106 int32
	_ = v3106
	var v3112 int32
	_ = v3112
	var v3118 int32
	_ = v3118
	var v3122 int32
	_ = v3122
	var v3125 int32
	_ = v3125
	var v3126 int32
	_ = v3126
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3132 int32
	_ = v3132
	var v3134 int32
	_ = v3134
	var v3135 int32
	_ = v3135
	var v3136 int32
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3144 int32
	_ = v3144
	var v3177 int32
	_ = v3177
	var v3179 int64
	_ = v3179
	var v3185 int32
	_ = v3185
	var v3195 int32
	_ = v3195
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3200 int32
	_ = v3200
	var v3201 int32
	_ = v3201
	var v3207 int32
	_ = v3207
	var v3208 int32
	_ = v3208
	var v3242 int32
	_ = v3242
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3247 int32
	_ = v3247
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3250 int32
	_ = v3250
	var v3254 int32
	_ = v3254
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3262 int32
	_ = v3262
	var v3266 int32
	_ = v3266
	var v3268 int32
	_ = v3268
	var v3269 int32
	_ = v3269
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3280 int32
	_ = v3280
	var v3282 int32
	_ = v3282
	var v3284 int32
	_ = v3284
	var v3287 int32
	_ = v3287
	var v3292 int32
	_ = v3292
	var v3293 int32
	_ = v3293
	var v3294 int32
	_ = v3294
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3299 int32
	_ = v3299
	var v3301 int32
	_ = v3301
	var v3304 int32
	_ = v3304
	var v3309 int32
	_ = v3309
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3318 int32
	_ = v3318
	var v3320 int32
	_ = v3320
	var v3321 int32
	_ = v3321
	var v3323 int32
	_ = v3323
	var v3333 int32
	_ = v3333
	var v3334 int32
	_ = v3334
	var v3340 int32
	_ = v3340
	var v3344 int32
	_ = v3344
	var v3346 int32
	_ = v3346
	var v3347 int32
	_ = v3347
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3354 int32
	_ = v3354
	var v3358 int32
	_ = v3358
	var v3360 int32
	_ = v3360
	var v3362 int32
	_ = v3362
	var v3365 int32
	_ = v3365
	var v3370 int32
	_ = v3370
	var v3371 int32
	_ = v3371
	var v3372 int32
	_ = v3372
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3377 int32
	_ = v3377
	var v3379 int32
	_ = v3379
	var v3382 int32
	_ = v3382
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3396 int32
	_ = v3396
	var v3398 int32
	_ = v3398
	var v3399 int32
	_ = v3399
	var v3401 int32
	_ = v3401
	var v3409 int32
	_ = v3409
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3446 int32
	_ = v3446
	var v3456 int32
	_ = v3456
	var v3460 int32
	_ = v3460
	var v3465 int32
	_ = v3465
	v2 = int32(0)
	v33 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[0]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	goto L3
L3:
	;
	v38 = int32(_a_F_InitializeParallelDSM_0)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1]))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v42
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v46 = F_add_size(m, v44, int32(96))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v46
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v51 = F_add_size(m, v49, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v51
	v55 = l0 + int32(36)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[3]))
	if v57 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v1305 = F_shm_toc_estimate(m, v55)
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L1
	} else {
		goto L191
	}
L7:
	;
	v69 = l0 + int32(12)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v70 <= int32(0) {
		v1274 = v2
		v1276 = v69
		v1277 = v2
		v1278 = v2
		v1280 = v2
		v1285 = v2
		v1295 = v2
		v1296 = v2
		v1297 = v2
		v1298 = v2
		v1299 = v2
		v1300 = v2
		v1304 = v2
		goto L6
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1274 = v2
	v1276 = l0 + int32(12)
	v1277 = v2
	v1278 = v2
	v1280 = v2
	v1285 = v2
	v1295 = v2
	v1296 = v2
	v1297 = v2
	v1298 = v2
	v1299 = v2
	v1300 = v2
	v1304 = v2
	goto L6
L9:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[4]))
	if v59 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[5]))
	if v61 == int32(0) {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v73 = m.G0
	v75 = v73 - int32(16)
	m.G0 = v75
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[6]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	m.G0 = v75 + int32(16)
	if v365 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v365 = v80
	goto L13
L15:
	;
	goto L16
L16:
	;
	v81 = int32(_a_F_InitializeParallelDSM_0)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1]))
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v85
	v89 = F_add_size(m, int32(0), int32(1))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v93 = F_add_size(m, int32(0), int32(_a_F_InitializeParallelDSM_1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v96 = F_add_size(m, v89, int32(1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v96
	v100 = F_add_size(m, v93, int32(32))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+8)) = v100
	v105 = F_shm_toc_estimate(m, v75+int32(8))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v108 = F_dsm_create(m, v105, int32(1))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v108 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v82
	v365 = int32(0)
	goto L13
L24:
	;
	goto L25
L25:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v116)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = int64(2880502729)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v105 & int32(-32)
	goto L26
L26:
	;
	v126 = F_shm_toc_allocate(m, v116, int32(_a_F_InitializeParallelDSM_1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v130 = F_dsa_create_in_place_ext(m, v126, int32(_a_F_InitializeParallelDSM_1), int32(72), v108)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_shm_toc_insert(m, v116, int64(-65535), v126)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v136 = F_shm_toc_allocate(m, v116, int32(12))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v138 = m.G0
	v140 = v138 - int32(16)
	m.G0 = v140
	v142 = int32(_a_F_InitializeParallelDSM_0)
	v143 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1]))
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v146
	v149 = F_dshash_create(m, v130, int32(_a_F_InitializeParallelDSM_2), v130)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v153 = F_dshash_create(m, v130, int32(_a_F_InitializeParallelDSM_3), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v143
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v149)+32))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v153)+32))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v161
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v164
	if int32(0) < v164 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_shm_toc_insert(m, v116, int64(-65534), v136)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L61
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L58
	}
L37:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[9]))
	v181 = v164
	v185 = v169
	v188 = v2
	goto L40
L38:
	;
	goto L39
L39:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v293)+16)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v293)+12)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v293)+8)) = v136
	F_on_dsm_detach(m, v108, int32(1607), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L57
	}
L40:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v185+v188<<(uint(int32(4))%32))+8))
	if v205 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v212 = F_dsa_allocate_extended(m, v130, v206*int32(116)+int32(20), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v253 = v181
	v254 = v185
	goto L44
L44:
	;
	v258 = v188 + int32(1)
	if v258 < v253 {
		v181 = v253
		v185 = v254
		v188 = v258
		goto L40
	} else {
		goto L56
	}
L45:
	;
	v214 = F_dsa_get_address(m, v130, v212)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_TupleDescCopy(m, v214, v205)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v214)+8)) = v188
	v222 = v140 + int32(7)
	v223 = F_dshash_find_or_insert(m, v153, v205+int32(8), v222)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+7)))
	if v225 == int32(1) {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v223)+4)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v223))) = v228
	F_dshash_release_lock(m, v153, v223)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+8)) = v205
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v140)+12)) = uint8(v234)
	v238 = F_dshash_find_or_insert(m, v149, v140+int32(8), v222)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+7)))
	if v240 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = v212
	v244 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v238)+4)) = uint8(v244)
	goto L54
L53:
	;
	goto L54
L54:
	;
	F_dshash_release_lock(m, v149, v238)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v249 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[9]))
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[8]))
	v253 = v251
	v254 = v249
	goto L44
L56:
	;
	goto L41
L57:
	;
	m.G0 = v140 + int32(16)
	goto L35
L58:
	;
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_4), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_5), int32(2253), int32(_a_F_InitializeParallelDSM_6))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
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
	F_dsm_pin_mapping(m, v108)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_dsa_pin_mapping(m, v130)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v324 = int32(_a_F_InitializeParallelDSM_7)
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v108
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v328)+4)) = v130
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v82
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v365 = v332
	goto L13
L64:
	;
	v371 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v371
	v1274 = v2
	v1276 = v69
	v1277 = v2
	v1278 = v371
	v1280 = v2
	v1285 = v2
	v1295 = v2
	v1296 = v2
	v1297 = v2
	v1298 = v2
	v1299 = v2
	v1300 = v2
	v1304 = v2
	goto L6
L65:
	;
	goto L66
L66:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v374 <= int32(0) {
		v1274 = v2
		v1276 = v69
		v1277 = v2
		v1278 = v365
		v1280 = v2
		v1285 = v2
		v1295 = v2
		v1296 = v2
		v1297 = v2
		v1298 = v2
		v1299 = v2
		v1300 = v2
		v1304 = v2
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v377 = int32(1)
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[10]))
	if v379 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v381 = v379
	v384 = v377
	goto L71
L69:
	;
	v424 = v377
	goto L70
L70:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v457 = F_add_size(m, v452, (v424+int32(31))&int32(-32))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L75
	}
L71:
	;
	v414 = F_strlen(m, v381+int32(24))
	mBase = m.M
	v417 = F_add_size(m, v384, v414+int32(1))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	v424 = v417
	goto L70
L73:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	if v419 != 0 {
		v381 = v419
		v384 = v417
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v457
	v460 = m.G0
	v462 = v460 - int32(16)
	m.G0 = v462
	v464 = int32(4)
	v466 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[11]))
	v467 = int32(0)
	if base.B2i32(v466 == v467)|base.B2i32(v466 == int32(_a_F_InitializeParallelDSM_8)) == v467 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v813 = F_add_size(m, v808, (v732+int32(31))&int32(-32))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L123
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L120
	}
L78:
	;
	v476 = v466
	v480 = v464
	goto L81
L79:
	;
	v732 = v464
	goto L80
L80:
	;
	m.G0 = v462 + int32(16)
	goto L76
L81:
	;
	v506 = int32(0)
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v476-int32(60))))
	if base.Ui32(v509) < base.Ui32(int32(2)) {
		v690 = v506
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v732 = v721
	goto L80
L83:
	;
	v721 = F_add_size(m, v480, v690)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L1
	} else {
		goto L118
	}
L84:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v476-int32(32))))
	if v514 == int32(0) {
		v690 = v506
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v518 = v476 + int32(-64)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v520 = F_strlen(m, v519)
	mBase = m.M
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v476-int32(40))))
	switch v523 {
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
		v625 = v506
		goto L86
	}
L86:
	;
	v656 = int32(1)
	v660 = F_add_size(m, v520+v656, v625+v656)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L105
	}
L87:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v476)+28))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v476)+36))
	if v544 == int32(0) {
		goto L77
	} else {
		goto L96
	}
L88:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v476)+28))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	if v538 == int32(0) {
		v625 = v506
		goto L86
	} else {
		goto L95
	}
L89:
	;
	v625 = int32(25)
	goto L86
L90:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v476)+28))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v527)))
	v530 = v528 >> (uint(int32(31)) % 32)
	if v528^v530-v530 < int32(1000) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v625 = int32(5)
	goto L86
L92:
	;
	v535 = int32(4)
	goto L94
L93:
	;
	v535 = int32(11)
	goto L94
L94:
	;
	v625 = v535
	goto L86
L95:
	;
	v541 = F_strlen(m, v538)
	mBase = m.M
	v625 = v541
	goto L86
L96:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v544)))
	if v547 == int32(0) {
		goto L77
	} else {
		goto L97
	}
L97:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	if v543 != v550 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v553 = v544
	goto L101
L99:
	;
	v603 = v547
	goto L100
L100:
	;
	v623 = F_strlen(m, v603)
	mBase = m.M
	v625 = v623
	goto L86
L101:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v553)+12))
	if v584 == int32(0) {
		goto L77
	} else {
		goto L103
	}
L102:
	;
	v603 = v584
	goto L100
L103:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v553)+16))
	if v543 != v587 {
		v553 = v553 + int32(12)
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v476)+20))
	if v662 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v663 = F_strlen(m, v662)
	mBase = m.M
	v664 = F_add_size(m, v660, v663)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	v666 = v660
	goto L108
L108:
	;
	v668 = F_add_size(m, v666, int32(1))
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	v666 = v664
	goto L108
L110:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v476)+20))
	if v670 == int32(0) {
		v679 = v668
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v681 = F_add_size(m, v679, int32(4))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L115
	}
L112:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v670))))
	if v673 == int32(0) {
		v679 = v668
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v677 = F_add_size(m, v668, int32(4))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v679 = v677
	goto L111
L115:
	;
	v684 = F_add_size(m, v681, int32(4))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v687 = F_add_size(m, v684, int32(4))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v690 = v687
	goto L83
L118:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v723 != int32(_a_F_InitializeParallelDSM_8) {
		v476 = v723
		v480 = v721
		goto L81
	} else {
		goto L119
	}
L119:
	;
	goto L82
L120:
	;
	v797 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	*(*int32)(unsafe.Add(mBase, uint32(v462)+4)) = v797
	*(*int32)(unsafe.Add(mBase, uint32(v462))) = v543
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_9), v462)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(3036), int32(_a_F_InitializeParallelDSM_11))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v813
	v819 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[12]))
	v820 = F_mul_size(m, int32(8), v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v822 = F_add_size(m, int32(4), v820)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v829 = F_add_size(m, v824, (v822+int32(31))&int32(-32))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v829
	v833 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[13]))
	if int32(2) <= v833 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v836 = F_EstimateSnapshotSpace(m, v33)
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	v846 = v2
	goto L129
L129:
	;
	v847 = F_EstimateSnapshotSpace(m, v37)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v843 = F_add_size(m, v838, (v836+int32(31))&int32(-32))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v843
	v846 = v836
	goto L129
L132:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v854 = F_add_size(m, v849, (v847+int32(31))&int32(-32))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v854
	v857 = int32(0)
	v859 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[14]))
	if v859 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v862 = v857
	v866 = v859
	goto L137
L135:
	;
	v903 = v857
	goto L136
L136:
	;
	v935 = F_mul_size(m, int32(4), v903)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L145
	}
L137:
	;
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v866)))
	if v892 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v903 = v898
	goto L136
L139:
	;
	v894 = F_add_size(m, v862, int32(1))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	v896 = v862
	goto L141
L141:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v866)+52))
	v898 = F_add_size(m, v896, v897)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	v896 = v894
	goto L141
L143:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v866)+80))
	if v900 != 0 {
		v862 = v898
		v866 = v900
		goto L137
	} else {
		goto L144
	}
L144:
	;
	goto L138
L145:
	;
	v937 = F_add_size(m, int32(32), v935)
	mBase = m.M
	v938 = m.ExcPending
	if v938 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v944 = F_add_size(m, v939, (v937+int32(31))&int32(-32))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v944
	v948 = F_add_size(m, v944, int32(32))
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v948
	v952 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[15]))
	if v952 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v952)))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v954)+4))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v954)+412))
	if v956 != 0 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v1023 = int32(1)
	goto L151
L151:
	;
	v1025 = F_mul_size(m, v1023, int32(12))
	mBase = m.M
	v1026 = m.ExcPending
	if v1026 != 0 {
		goto L1
	} else {
		goto L156
	}
L152:
	;
	v1023 = v1019 + int32(1)
	goto L151
L153:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v954)+376))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v954)+364))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v954)+352))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v954)+340))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v954)+328))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v954)+316))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v954)+304))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v954)+292))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v954)+280))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v954)+268))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v954)+256))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v954)+244))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v954)+232))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v954)+220))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v954)+208))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v954)+196))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v954)+184))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v954)+172))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v954)+160))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v954)+148))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v954)+136))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v954)+124))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v954)+112))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v954)+100))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v954)+88))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v954)+76))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v954)+64))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v954)+52))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v954)+40))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v954)+28))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v954)+16))
	v1019 = v957 + (v958 + (v959 + (v960 + (v961 + (v962 + (v963 + (v964 + (v965 + (v966 + (v967 + (v968 + (v969 + (v970 + (v971 + (v972 + (v973 + (v974 + (v975 + (v976 + (v977 + (v978 + (v979 + (v980 + (v981 + (v982 + (v983 + (v984 + (v985 + (v986 + (v987 + v955))))))))))))))))))))))))))))))
	goto L155
L154:
	;
	v1019 = v955
	goto L155
L155:
	;
	goto L152
L156:
	;
	v1027 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1032 = F_add_size(m, v1027, (v1025+int32(31))&int32(-32))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1032
	v1037 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[16]))
	if v1037 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+4))
	v1040 = v1038
	goto L160
L159:
	;
	v1040 = int32(0)
	goto L160
L160:
	;
	v1041 = F_mul_size(m, int32(4), v1040)
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1050 = F_add_size(m, v1045, (v1041+int32(43))&int32(-32))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1050
	v1056 = F_add_size(m, v1050, int32(1056))
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1056
	v1061 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[17]))
	if v1061 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v1061)))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+4))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+412))
	if v1065 != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	v1129 = int32(0)
	goto L166
L166:
	;
	v1131 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[18]))
	if v1131 != 0 {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v1129 = v1128
	goto L166
L168:
	;
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+376))
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+364))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+352))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+340))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+328))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+316))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+304))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+292))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+280))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+268))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+256))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+244))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+232))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+220))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+208))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+196))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+184))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+172))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+160))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+148))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+136))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+124))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+112))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+100))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+88))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+76))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+64))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+52))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+40))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+28))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1063)+16))
	v1128 = v1066 + (v1067 + (v1068 + (v1069 + (v1070 + (v1071 + (v1072 + (v1073 + (v1074 + (v1075 + (v1076 + (v1077 + (v1078 + (v1079 + (v1080 + (v1081 + (v1082 + (v1083 + (v1084 + (v1085 + (v1086 + (v1087 + (v1088 + (v1089 + (v1090 + (v1091 + (v1092 + (v1093 + (v1094 + (v1095 + (v1096 + v1064))))))))))))))))))))))))))))))
	goto L170
L169:
	;
	v1128 = v1064
	goto L170
L170:
	;
	goto L167
L171:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(v1131)))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+4))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+412))
	if v1135 != 0 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1200 = v1129
	goto L173
L173:
	;
	v1202 = v1200 << (uint(int32(2)) % 32)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1210 = F_add_size(m, v1205, (v1202+int32(39))&int32(-32))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L1
	} else {
		goto L178
	}
L174:
	;
	v1200 = v1198 + v1129
	goto L173
L175:
	;
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+376))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+364))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+352))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+340))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+328))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+316))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+304))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+292))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+280))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+268))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+256))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+244))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+232))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+220))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+208))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+196))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+184))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+172))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+160))
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+148))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+136))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+124))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+112))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+100))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+88))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+76))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+64))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+52))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+40))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+28))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1133)+16))
	v1198 = v1136 + (v1137 + (v1138 + (v1139 + (v1140 + (v1141 + (v1142 + (v1143 + (v1144 + (v1145 + (v1146 + (v1147 + (v1148 + (v1149 + (v1150 + (v1151 + (v1152 + (v1153 + (v1154 + (v1155 + (v1156 + (v1157 + (v1158 + (v1159 + (v1160 + (v1161 + (v1162 + (v1163 + (v1164 + (v1165 + (v1166 + v1134))))))))))))))))))))))))))))))
	goto L177
L176:
	;
	v1198 = v1134
	goto L177
L177:
	;
	goto L174
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1210
	v1215 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v1216 = m.ExcPending
	if v1216 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v1218 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	if v1218 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1219 = F_strlen(m, v1218)
	mBase = m.M
	v1222 = F_add_size(m, v1215, v1219+int32(1))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	v1224 = v1215
	goto L182
L182:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1230 = F_add_size(m, v1225, (v1224+int32(31))&int32(-32))
	mBase = m.M
	v1231 = m.ExcPending
	if v1231 != 0 {
		goto L1
	} else {
		goto L184
	}
L183:
	;
	v1224 = v1222
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1230
	v1233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1235 = F_add_size(m, v1233, int32(12))
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1235
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1241 = F_mul_size(m, int32(_a_F_InitializeParallelDSM_12), v1240)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v1247 = F_add_size(m, v1238, (v1241+int32(31))&int32(-32))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1247
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1252 = F_add_size(m, v1250, int32(1))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1252
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1257 = F_strlen(m, v1256)
	mBase = m.M
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1259 = F_strlen(m, v1258)
	mBase = m.M
	v1265 = F_add_size(m, v1255, (v1257+v1259+int32(33))&int32(-32))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1265
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1270 = F_add_size(m, v1268, int32(1))
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1270
	v1274 = v732
	v1276 = v69
	v1277 = v424
	v1278 = v365
	v1280 = v1224
	v1285 = v822
	v1295 = v846
	v1296 = v847
	v1297 = v937
	v1298 = v1025
	v1299 = v1041 + int32(12)
	v1300 = v1202 + int32(8)
	v1304 = int32(1048)
	goto L6
L191:
	;
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v1276)))
	if v1307 <= int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if v1315 != 0 {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1315 = v1310
	goto L192
L194:
	;
	goto L195
L195:
	;
	v1312 = F_dsm_create(m, v1305, int32(1))
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1312
	v1315 = v1312
	goto L192
L197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1324)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1324)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1324))) = int64(1346862204)
	*(*int32)(unsafe.Add(mBase, uint32(v1324)+12)) = v1305 & int32(-32)
	goto L202
L198:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+24))
	v1324 = v1316
	goto L197
L199:
	;
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1320 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[7]))
	v1321 = F_MemoryContextAlloc(m, v1320, v1305)
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v1321
	v1324 = v1321
	goto L197
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1324
	v1336 = F_shm_toc_allocate(m, v1324, int32(80))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v1339 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336))) = v1339
	v1342 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+4)) = v1342
	v1345 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+8)) = v1345
	v1348 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[23]))
	v1351 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[24])))
	if v1351 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+12)) = v1352
	v1359 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[25]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336+int32(16)))) = v1359
	v1362 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[26]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336+int32(28)))) = v1362
	goto L208
L205:
	;
	v1352 = v1348
	goto L207
L206:
	;
	v1352 = int32(0)
	goto L207
L207:
	;
	goto L204
L208:
	;
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[27])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1336)+32)) = uint8(v1365)
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[28])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1336)+33)) = uint8(v1368)
	v1371 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[29]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+20)) = v1371
	v1374 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+24)) = v1374
	v1377 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+36)) = v1377
	v1380 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+40)) = v1380
	v1383 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[33]))
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+44)) = v1383
	v1386 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[34]))
	*(*int64)(unsafe.Add(mBase, uint32(v1336)+48)) = v1386
	v1389 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[35]))
	*(*int64)(unsafe.Add(mBase, uint32(v1336)+56)) = v1389
	v1392 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[36]))
	*(*int64)(unsafe.Add(mBase, uint32(v1336)+72)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+68)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1336)+64)) = v1392
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1398, int64(-65535), v1336)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v1402 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3456 = m.ExcPending
	if v3456 != 0 {
		goto L1
	} else {
		goto L524
	}
L211:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1406 = F_shm_toc_allocate(m, v1405, v1277)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	v3446 = v1402
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3446
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v39
	return
L214:
	;
	v1409 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[10]))
	if v1409 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1412 = v1409
	v1413 = v1406
	v1414 = v1277
	goto L218
L216:
	;
	v1570 = v1406
	goto L217
L217:
	;
	v1599 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1570))) = uint8(v1599)
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1601, int64(-65533), v1406)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L1
	} else {
		goto L252
	}
L218:
	;
	v1443 = v1412 + int32(24)
	if v1414 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v1570 = v1564
	goto L217
L220:
	;
	v1563 = v1559 + (v1556 - v1413) + int32(1)
	v1564 = v1413 + v1563
	v1566 = *(*int32)(unsafe.Add(mBase, uint32(v1412)))
	if v1566 != 0 {
		v1412 = v1566
		v1413 = v1564
		v1414 = v1414 - v1563
		goto L218
	} else {
		goto L251
	}
L221:
	;
	v1559 = F_strlen(m, v1555)
	mBase = m.M
	goto L220
L222:
	;
	v1555 = v1443
	v1556 = v1413
	goto L221
L223:
	;
	goto L224
L224:
	;
	v1449 = v1414 - int32(1)
	if (v1413^v1443)&int32(3) != 0 {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	v1552 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1549))) = uint8(v1552)
	v1555 = v1548
	v1556 = v1549
	goto L221
L226:
	;
	v1533 = v1528
	v1534 = v1529
	v1535 = v1530
	goto L247
L227:
	;
	if v1523 == int32(0) {
		v1548 = v1521
		v1549 = v1522
		goto L225
	} else {
		goto L246
	}
L228:
	;
	v1521 = v1443
	v1522 = v1413
	v1523 = v1449
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1453 = int32(0)
	if base.B2i32(v1443&int32(3) == v1453)|base.B2i32(v1449 == v1453) == v1453 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	if v1489 == int32(0) {
		v1548 = v1486
		v1549 = v1487
		goto L225
	} else {
		goto L240
	}
L232:
	;
	v1465 = v1443
	v1466 = v1413
	v1467 = v1449
	goto L235
L233:
	;
	goto L234
L234:
	;
	v1486 = v1443
	v1487 = v1413
	v1488 = v1449
	v1489 = base.B2i32(v1449 != v1453)
	goto L231
L235:
	;
	v1469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1465))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1466))) = uint8(v1469)
	if v1469 == int32(0) {
		v1528 = v1465
		v1529 = v1466
		v1530 = v1467
		goto L226
	} else {
		goto L237
	}
L236:
	;
	v1486 = v1480
	v1487 = v1474
	v1488 = v1476
	v1489 = v1478
	goto L231
L237:
	;
	v1473 = int32(1)
	v1474 = v1466 + v1473
	v1476 = v1467 - v1473
	v1477 = int32(0)
	v1478 = base.B2i32(v1476 != v1477)
	v1480 = v1465 + v1473
	if v1480&int32(3) == v1477 {
		v1486 = v1480
		v1487 = v1474
		v1488 = v1476
		v1489 = v1478
		goto L231
	} else {
		goto L238
	}
L238:
	;
	if v1476 != 0 {
		v1465 = v1480
		v1466 = v1474
		v1467 = v1476
		goto L235
	} else {
		goto L239
	}
L239:
	;
	goto L236
L240:
	;
	v1492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1486))))
	if base.B2i32(v1492 == int32(0))|base.B2i32(base.Ui32(v1488) < base.Ui32(int32(4))) != 0 {
		v1521 = v1486
		v1522 = v1487
		v1523 = v1488
		goto L227
	} else {
		goto L241
	}
L241:
	;
	v1499 = v1486
	v1500 = v1487
	v1501 = v1488
	goto L242
L242:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1499)))
	v1507 = int32(-2139062144)
	if (int32(16843008)-v1504|v1504)&v1507 != v1507 {
		v1528 = v1499
		v1529 = v1500
		v1530 = v1501
		goto L226
	} else {
		goto L244
	}
L243:
	;
	v1521 = v1515
	v1522 = v1513
	v1523 = v1517
	goto L227
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1500))) = v1504
	v1512 = int32(4)
	v1513 = v1500 + v1512
	v1515 = v1499 + v1512
	v1517 = v1501 - v1512
	if base.Ui32(int32(3)) < base.Ui32(v1517) {
		v1499 = v1515
		v1500 = v1513
		v1501 = v1517
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v1528 = v1521
	v1529 = v1522
	v1530 = v1523
	goto L226
L247:
	;
	v1537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1533))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1534))) = uint8(v1537)
	if v1537 == int32(0) {
		v1548 = v1533
		v1549 = v1534
		goto L225
	} else {
		goto L249
	}
L248:
	;
	v1548 = v1544
	v1549 = v1542
	goto L225
L249:
	;
	v1541 = int32(1)
	v1542 = v1534 + v1541
	v1544 = v1533 + v1541
	v1546 = v1535 - v1541
	if v1546 != 0 {
		v1533 = v1544
		v1534 = v1542
		v1535 = v1546
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	goto L219
L252:
	;
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1606 = F_shm_toc_allocate(m, v1605, v1274)
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v1608 = m.G0
	v1610 = v1608 - int32(112)
	m.G0 = v1610
	v1612 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+108)) = v1606 + v1612
	v1616 = v1274 - v1612
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+104)) = v1616
	v1619 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[11]))
	v1620 = int32(0)
	if base.B2i32(v1619 == v1620)|base.B2i32(v1619 == int32(_a_F_InitializeParallelDSM_8)) == v1620 {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2038, int64(-65532), v1606)
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L1
	} else {
		goto L307
	}
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L1
	} else {
		goto L304
	}
L256:
	;
	v1628 = v1616
	v1631 = v1619
	goto L259
L257:
	;
	v1953 = v1616
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1606))) = v1616 - v1953
	m.G0 = v1610 + int32(112)
	goto L254
L259:
	;
	v1661 = *(*int32)(unsafe.Add(mBase, uint32(v1631-int32(60))))
	if base.Ui32(v1661) < base.Ui32(int32(2)) {
		v1918 = v1628
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1953 = v1918
	goto L258
L261:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+4))
	if v1949 != int32(_a_F_InitializeParallelDSM_8) {
		v1628 = v1918
		v1631 = v1949
		goto L259
	} else {
		goto L303
	}
L262:
	;
	v1665 = v1631 - int32(32)
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v1665)))
	if v1666 == int32(0) {
		v1918 = v1628
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1670 = v1631 + int32(-64)
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1670)))
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+96)) = v1671
	F_do_serialize(m, v1610+int32(108), v1610+int32(104), int32(_a_F_InitializeParallelDSM_13), v1610+int32(96))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v1631-int32(40))))
	switch v1684 {
	case 0:
		goto L270
	case 1:
		goto L269
	case 2:
		goto L268
	case 3:
		goto L267
	case 4:
		goto L266
	default:
		goto L265
	}
L265:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+20))
	if v1860 != 0 {
		goto L291
	} else {
		goto L292
	}
L266:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+28))
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v1737)))
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+36))
	if v1739 == int32(0) {
		goto L255
	} else {
		goto L281
	}
L267:
	;
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+28))
	v1724 = *(*int32)(unsafe.Add(mBase, uint32(v1723)))
	if v1724 != 0 {
		goto L277
	} else {
		goto L278
	}
L268:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+28))
	v1710 = *(*float64)(unsafe.Add(mBase, uint32(v1709)))
	*(*float64)(unsafe.Add(mBase, uint32(v1610)+40)) = v1710
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+32)) = int32(17)
	F_do_serialize(m, v1610+int32(108), v1610+int32(104), int32(_a_F_InitializeParallelDSM_14), v1610+int32(32))
	mBase = m.M
	v1722 = m.ExcPending
	if v1722 != 0 {
		goto L1
	} else {
		goto L276
	}
L269:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+28))
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v1697)))
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+16)) = v1698
	F_do_serialize(m, v1610+int32(108), v1610+int32(104), int32(_a_F_InitializeParallelDSM_15), v1610+int32(16))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L1
	} else {
		goto L275
	}
L270:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+28))
	v1692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1691))))
	if v1692 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1693 = int32(_a_F_InitializeParallelDSM_16)
	goto L273
L272:
	;
	v1693 = int32(_a_F_InitializeParallelDSM_17)
	goto L273
L273:
	;
	F_do_serialize(m, v1610+int32(108), v1610+int32(104), v1693, int32(0))
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	goto L265
L275:
	;
	goto L265
L276:
	;
	goto L265
L277:
	;
	v1726 = v1724
	goto L279
L278:
	;
	v1726 = int32(_a_F_InitializeParallelDSM_18)
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+48)) = v1726
	F_do_serialize(m, v1610+int32(108), v1610+int32(104), int32(_a_F_InitializeParallelDSM_13), v1610+int32(48))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	goto L265
L281:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1739)))
	if v1742 == int32(0) {
		goto L255
	} else {
		goto L282
	}
L282:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1739)+4))
	if v1738 != v1745 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1748 = v1739
	goto L286
L284:
	;
	v1792 = v1742
	goto L285
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+80)) = v1792
	F_do_serialize(m, v1610+int32(108), v1610+int32(104), int32(_a_F_InitializeParallelDSM_13), v1610+int32(80))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L1
	} else {
		goto L290
	}
L286:
	;
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+12))
	if v1779 == int32(0) {
		goto L255
	} else {
		goto L288
	}
L287:
	;
	v1792 = v1779
	goto L285
L288:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1748)+16))
	if v1738 != v1782 {
		v1748 = v1748 + int32(12)
		goto L286
	} else {
		goto L289
	}
L289:
	;
	goto L287
L290:
	;
	goto L265
L291:
	;
	v1862 = v1860
	goto L293
L292:
	;
	v1862 = int32(_a_F_InitializeParallelDSM_18)
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1610))) = v1862
	F_do_serialize(m, v1610+int32(108), v1610+int32(104), int32(_a_F_InitializeParallelDSM_13), v1610)
	mBase = m.M
	v1870 = m.ExcPending
	if v1870 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1871 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+20))
	if v1871 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	if base.Ui32(v1891) <= base.Ui32(int32(3)) {
		goto L210
	} else {
		goto L300
	}
L296:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+104))
	v1891 = v1888
	goto L295
L297:
	;
	v1874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1871))))
	if v1874 == int32(0) {
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+104))
	if base.Ui32(v1877) <= base.Ui32(int32(3)) {
		goto L210
	} else {
		goto L299
	}
L299:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+108))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v1631)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1880))) = v1881
	v1883 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+108)) = v1880 + v1883
	v1891 = v1877 - v1883
	goto L295
L300:
	;
	v1894 = *(*int32)(unsafe.Add(mBase, uint32(v1610)+108))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1665)))
	*(*int32)(unsafe.Add(mBase, uint32(v1894))) = v1895
	v1898 = v1891 & int32(-4)
	if v1898 == int32(4) {
		goto L210
	} else {
		goto L301
	}
L301:
	;
	v1903 = *(*int32)(unsafe.Add(mBase, uint32(v1631-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+4)) = v1903
	if v1898 == int32(8) {
		goto L210
	} else {
		goto L302
	}
L302:
	;
	v1909 = *(*int32)(unsafe.Add(mBase, uint32(v1631-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v1894)+8)) = v1909
	v1911 = int32(12)
	v1912 = v1891 - v1911
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+104)) = v1912
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+108)) = v1894 + v1911
	v1918 = v1912
	goto L261
L303:
	;
	goto L260
L304:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v1670)))
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+68)) = v2025
	*(*int32)(unsafe.Add(mBase, uint32(v1610)+64)) = v1738
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_9), v1610-int32(-64))
	mBase = m.M
	v2032 = m.ExcPending
	if v2032 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(3036), int32(_a_F_InitializeParallelDSM_11))
	mBase = m.M
	v2037 = m.ExcPending
	if v2037 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L307:
	;
	v2042 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2043 = F_shm_toc_allocate(m, v2042, v1285)
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v2046 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v2043))) = v2046
	v2049 = v2046 << (uint(int32(3)) % 32)
	if base.Ui32(v2049|int32(4)) <= base.Ui32(v1285) {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v2076 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2076, int64(-65531), v2043)
	mBase = m.M
	v2079 = m.ExcPending
	if v2079 != 0 {
		goto L1
	} else {
		goto L317
	}
L310:
	;
	v2053 = int32(0)
	if base.B2i32(v2049 == v2053)|base.B2i32(v2046 <= v2053) != 0 {
		goto L309
	} else {
		goto L313
	}
L311:
	;
	goto L312
L312:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2066 = m.ExcPending
	if v2066 != 0 {
		goto L1
	} else {
		goto L314
	}
L313:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[37]))
	base.MemoryCopy(m, v2043+int32(4), v2061, v2049)
	goto L309
L314:
	;
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_19), int32(0))
	mBase = m.M
	v2070 = m.ExcPending
	if v2070 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_20), int32(327), int32(_a_F_InitializeParallelDSM_21))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
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
	v2081 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[13]))
	if int32(2) <= v2081 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2085 = F_shm_toc_allocate(m, v2084, v1295)
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2144 = F_shm_toc_allocate(m, v2143, v1296)
	mBase = m.M
	v2145 = m.ExcPending
	if v2145 != 0 {
		goto L1
	} else {
		goto L336
	}
L321:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v2094 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+29)))
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v2096 = *(*int64)(unsafe.Add(mBase, uint32(v33)+4))
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v2098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2085)+16)) = uint8(v2098)
	*(*int32)(unsafe.Add(mBase, uint32(v2085)+20)) = v2097
	*(*int64)(unsafe.Add(mBase, uint32(v2085))) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v2085)+8)) = v2095
	*(*uint8)(unsafe.Add(mBase, uint32(v2085)+17)) = uint8(v2094)
	if v2094&int32(1) != 0 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2138, int64(-65530), v2085)
	mBase = m.M
	v2141 = m.ExcPending
	if v2141 != 0 {
		goto L1
	} else {
		goto L335
	}
L323:
	;
	v2107 = v2093
	goto L325
L324:
	;
	v2107 = int32(0)
	goto L325
L325:
	;
	if v2098 != 0 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2108 = v2107
	goto L328
L327:
	;
	v2108 = v2093
	goto L328
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2085)+12)) = v2108
	v2110 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v2110 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	if v2108 <= int32(0) {
		goto L332
	} else {
		goto L333
	}
L330:
	;
	v2114 = v2110 << (uint(int32(2)) % 32)
	if v2114 == int32(0) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v2119 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	base.MemoryCopy(m, v2085+int32(24), v2119, v2114)
	goto L329
L332:
	;
	goto L322
L333:
	;
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v2126 = v2124 << (uint(int32(2)) % 32)
	if v2126 == int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	base.MemoryCopy(m, v2085+v2129<<(uint(int32(2))%32)+int32(24), v2135, v2126)
	goto L332
L335:
	;
	goto L320
L336:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v2153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+29)))
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v2155 = *(*int64)(unsafe.Add(mBase, uint32(v37)+4))
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v2157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2144)+16)) = uint8(v2157)
	*(*int32)(unsafe.Add(mBase, uint32(v2144)+20)) = v2156
	*(*int64)(unsafe.Add(mBase, uint32(v2144))) = v2155
	*(*int32)(unsafe.Add(mBase, uint32(v2144)+8)) = v2154
	*(*uint8)(unsafe.Add(mBase, uint32(v2144)+17)) = uint8(v2153)
	if v2153&int32(1) != 0 {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2197, int64(-65529), v2144)
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L1
	} else {
		goto L350
	}
L338:
	;
	v2166 = v2152
	goto L340
L339:
	;
	v2166 = int32(0)
	goto L340
L340:
	;
	if v2157 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v2167 = v2166
	goto L343
L342:
	;
	v2167 = v2152
	goto L343
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2144)+12)) = v2167
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v2169 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	if v2167 <= int32(0) {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v2173 = v2169 << (uint(int32(2)) % 32)
	if v2173 == int32(0) {
		goto L344
	} else {
		goto L346
	}
L346:
	;
	v2178 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	base.MemoryCopy(m, v2144+int32(24), v2178, v2173)
	goto L344
L347:
	;
	goto L337
L348:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v2185 = v2183 << (uint(int32(2)) % 32)
	if v2185 == int32(0) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	base.MemoryCopy(m, v2144+v2188<<(uint(int32(2))%32)+int32(24), v2194, v2185)
	goto L347
L350:
	;
	v2201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2203 = F_shm_toc_allocate(m, v2201, int32(4))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2203))) = v1278
	v2206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2206, int64(-65526), v2203)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2211 = F_shm_toc_allocate(m, v2210, v1297)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v2213 = int32(0)
	v2215 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v2211))) = v2215
	v2218 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[38])))
	*(*uint8)(unsafe.Add(mBase, uint32(v2211)+4)) = uint8(v2218)
	v2221 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[39]))
	*(*int64)(unsafe.Add(mBase, uint32(v2211)+8)) = v2221
	v2224 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[14]))
	v2225 = *(*int64)(unsafe.Add(mBase, uint32(v2224)))
	*(*int64)(unsafe.Add(mBase, uint32(v2211)+16)) = v2225
	v2228 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v2211)+24)) = v2228
	v2231 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[41]))
	if v2213 < v2231 {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2421, int64(-65528), v2211)
	mBase = m.M
	v2424 = m.ExcPending
	if v2424 != 0 {
		goto L1
	} else {
		goto L385
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2211)+28)) = v2231
	v2236 = v2231 << (uint(int32(2)) % 32)
	if v2236 == int32(0) {
		goto L354
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v2248 = v2224
	v2250 = v2213
	goto L359
L358:
	;
	v2242 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[42]))
	base.MemoryCopy(m, v2211+int32(32), v2242, v2236)
	goto L354
L359:
	;
	v2276 = *(*int32)(unsafe.Add(mBase, uint32(v2248)))
	if v2276 != 0 {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	v2287 = v2282 << (uint(int32(2)) % 32)
	v2288 = F_palloc(m, v2287)
	mBase = m.M
	v2289 = m.ExcPending
	if v2289 != 0 {
		goto L1
	} else {
		goto L367
	}
L361:
	;
	v2278 = F_add_size(m, v2250, int32(1))
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L1
	} else {
		goto L364
	}
L362:
	;
	v2280 = v2250
	goto L363
L363:
	;
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v2248)+52))
	v2282 = F_add_size(m, v2280, v2281)
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L1
	} else {
		goto L365
	}
L364:
	;
	v2280 = v2278
	goto L363
L365:
	;
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v2248)+80))
	if v2284 != 0 {
		v2248 = v2284
		v2250 = v2282
		goto L359
	} else {
		goto L366
	}
L366:
	;
	goto L360
L367:
	;
	v2291 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[14]))
	if v2291 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v2295 = int32(0)
	v2296 = v2291
	goto L371
L369:
	;
	goto L370
L370:
	;
	F_pg_qsort(m, v2288, v2282, int32(4), int32(185))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L1
	} else {
		goto L383
	}
L371:
	;
	v2324 = *(*int32)(unsafe.Add(mBase, uint32(v2296)))
	if v2324 != 0 {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	goto L370
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2288+v2295<<(uint(int32(2))%32)))) = v2324
	v2331 = v2295 + int32(1)
	goto L375
L374:
	;
	v2331 = v2295
	goto L375
L375:
	;
	v2332 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+52))
	if int32(0) < v2332 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v2336 = v2332 << (uint(int32(2)) % 32)
	if v2336 != 0 {
		goto L379
	} else {
		goto L380
	}
L377:
	;
	v2344 = v2332
	goto L378
L378:
	;
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+80))
	if v2346 != 0 {
		v2295 = v2344 + v2331
		v2296 = v2346
		goto L371
	} else {
		goto L382
	}
L379:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+48))
	base.MemoryCopy(m, v2288+v2331<<(uint(int32(2))%32), v2340, v2336)
	goto L381
L380:
	;
	goto L381
L381:
	;
	v2342 = *(*int32)(unsafe.Add(mBase, uint32(v2296)+52))
	v2344 = v2342
	goto L378
L382:
	;
	goto L372
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2211)+28)) = v2282
	if v2287 == int32(0) {
		goto L354
	} else {
		goto L384
	}
L384:
	;
	base.MemoryCopy(m, v2211+int32(32), v2288, v2287)
	goto L354
L385:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2426 = F_shm_toc_allocate(m, v2425, v1298)
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v2428 = m.G0
	v2430 = v2428 - int32(80)
	m.G0 = v2430
	v2433 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[15]))
	if v2433 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2430)+48)) = int64(51539607564)
	v2437 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v2430)+72)) = v2437
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(v2433)))
	v2442 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+4))
	v2443 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+412))
	if v2443 != 0 {
		goto L391
	} else {
		goto L392
	}
L388:
	;
	v2751 = v2426
	goto L389
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2751)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2751))) = int64(0)
	m.G0 = v2430 + int32(80)
	v2787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2787, int64(-65525), v2426)
	mBase = m.M
	v2790 = m.ExcPending
	if v2790 != 0 {
		goto L1
	} else {
		goto L425
	}
L390:
	;
	v2510 = F_hash_create(m, int32(_a_F_InitializeParallelDSM_22), v2506, v2430+int32(32), int32(1064))
	mBase = m.M
	v2511 = m.ExcPending
	if v2511 != 0 {
		goto L1
	} else {
		goto L394
	}
L391:
	;
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+376))
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+364))
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+352))
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+340))
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+328))
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+316))
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+304))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+292))
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+280))
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+268))
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+256))
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+244))
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+232))
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+220))
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+208))
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+196))
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+184))
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+172))
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+160))
	v2463 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+148))
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+136))
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+124))
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+112))
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+100))
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+88))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+76))
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+64))
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+52))
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+40))
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+28))
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v2441)+16))
	v2506 = v2444 + (v2445 + (v2446 + (v2447 + (v2448 + (v2449 + (v2450 + (v2451 + (v2452 + (v2453 + (v2454 + (v2455 + (v2456 + (v2457 + (v2458 + (v2459 + (v2460 + (v2461 + (v2462 + (v2463 + (v2464 + (v2465 + (v2466 + (v2467 + (v2468 + (v2469 + (v2470 + (v2471 + (v2472 + (v2473 + (v2474 + v2442))))))))))))))))))))))))))))))
	goto L393
L392:
	;
	v2506 = v2442
	goto L393
L393:
	;
	goto L390
L394:
	;
	v2513 = v2430 + int32(12)
	v2515 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[15]))
	F_hash_seq_init(m, v2513, v2515)
	mBase = m.M
	v2517 = m.ExcPending
	if v2517 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	v2518 = F_hash_seq_search(m, v2513)
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	if v2518 != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2524 = v2518
	goto L400
L398:
	;
	goto L399
L399:
	;
	v2593 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[43]))
	if v2593 != 0 {
		goto L405
	} else {
		goto L406
	}
L400:
	;
	v2554 = F_hash_search(m, v2510, v2524, int32(1), int32(0))
	mBase = m.M
	v2555 = m.ExcPending
	if v2555 != 0 {
		goto L1
	} else {
		goto L402
	}
L401:
	;
	goto L399
L402:
	;
	v2558 = F_hash_seq_search(m, v2430+int32(12))
	mBase = m.M
	v2559 = m.ExcPending
	if v2559 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	if v2558 != 0 {
		v2524 = v2558
		goto L400
	} else {
		goto L404
	}
L404:
	;
	goto L401
L405:
	;
	v2598 = v2593
	goto L408
L406:
	;
	goto L407
L407:
	;
	v2667 = v2430 + int32(12)
	F_hash_seq_init(m, v2667, v2510)
	mBase = m.M
	v2669 = m.ExcPending
	if v2669 != 0 {
		goto L1
	} else {
		goto L415
	}
L408:
	;
	v2626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2598)+16)))
	if v2626 == int32(1) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L407
L410:
	;
	v2631 = F_hash_search(m, v2510, v2598, int32(2), int32(0))
	mBase = m.M
	v2632 = m.ExcPending
	if v2632 != 0 {
		goto L1
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v2633 = *(*int32)(unsafe.Add(mBase, uint32(v2598)+24))
	if v2633 != 0 {
		v2598 = v2633
		goto L408
	} else {
		goto L414
	}
L413:
	;
	goto L412
L414:
	;
	goto L409
L415:
	;
	v2670 = F_hash_seq_search(m, v2667)
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	if v2670 != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v2675 = v2426
	v2676 = v2670
	goto L420
L418:
	;
	v2717 = v2426
	goto L419
L419:
	;
	F_hash_destroy(m, v2510)
	mBase = m.M
	v2747 = m.ExcPending
	if v2747 != 0 {
		goto L1
	} else {
		goto L424
	}
L420:
	;
	v2704 = *(*int32)(unsafe.Add(mBase, uint32(v2676)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2675)+8)) = v2704
	v2706 = *(*int64)(unsafe.Add(mBase, uint32(v2676)))
	*(*int64)(unsafe.Add(mBase, uint32(v2675))) = v2706
	v2708 = int32(12)
	v2709 = v2675 + v2708
	v2712 = F_hash_seq_search(m, v2430+v2708)
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L1
	} else {
		goto L422
	}
L421:
	;
	v2717 = v2709
	goto L419
L422:
	;
	if v2712 != 0 {
		v2675 = v2709
		v2676 = v2712
		goto L420
	} else {
		goto L423
	}
L423:
	;
	goto L421
L424:
	;
	v2751 = v2717
	goto L389
L425:
	;
	v2791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2792 = F_shm_toc_allocate(m, v2791, v1299)
	mBase = m.M
	v2793 = m.ExcPending
	if v2793 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v2796 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v2792))) = v2796
	v2799 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v2792)+4)) = v2799
	v2802 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[16]))
	if v2802 != 0 {
		goto L428
	} else {
		goto L429
	}
L427:
	;
	v2887 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2887, int64(-65524), v2792)
	mBase = m.M
	v2890 = m.ExcPending
	if v2890 != 0 {
		goto L1
	} else {
		goto L435
	}
L428:
	;
	v2803 = *(*int32)(unsafe.Add(mBase, uint32(v2802)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2792)+8)) = v2803
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(v2802)+4))
	if v2805 <= int32(0) {
		goto L427
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2792)+8)) = int32(0)
	goto L427
L431:
	;
	v2813 = int32(0)
	goto L432
L432:
	;
	v2843 = v2813 << (uint(int32(2)) % 32)
	v2845 = *(*int32)(unsafe.Add(mBase, uint32(v2802)+12))
	v2847 = *(*int32)(unsafe.Add(mBase, uint32(v2845+v2843)))
	*(*int32)(unsafe.Add(mBase, uint32(v2792+int32(12)+v2843))) = v2847
	v2850 = v2813 + int32(1)
	v2851 = *(*int32)(unsafe.Add(mBase, uint32(v2802)+4))
	if v2850 < v2851 {
		v2813 = v2850
		goto L432
	} else {
		goto L434
	}
L433:
	;
	goto L427
L434:
	;
	goto L433
L435:
	;
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2892 = F_shm_toc_allocate(m, v2891, v1304)
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	v2895 = int32(524)
	base.MemoryCopy(m, v2892, int32(_a_F_InitializeParallelDSM_23), v2895)
	base.MemoryCopy(m, v2892+v2895, int32(_a_F_InitializeParallelDSM_24), v2895)
	v2902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2902, int64(-65523), v2892)
	mBase = m.M
	v2905 = m.ExcPending
	if v2905 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	v2906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2907 = F_shm_toc_allocate(m, v2906, v1300)
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	v2909 = m.G0
	v2911 = v2909 - int32(32)
	m.G0 = v2911
	v2914 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[17]))
	if v2914 == int32(0) {
		v2969 = v2907
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v2997 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2969))) = v2997
	v3000 = v2969 + int32(4)
	v3002 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[18]))
	if v3002 == v2997 {
		v3057 = v3000
		goto L448
	} else {
		goto L449
	}
L440:
	;
	v2918 = v2911 + int32(12)
	F_hash_seq_init(m, v2918, v2914)
	mBase = m.M
	v2920 = m.ExcPending
	if v2920 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	v2921 = F_hash_seq_search(m, v2918)
	mBase = m.M
	v2922 = m.ExcPending
	if v2922 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	if v2921 == int32(0) {
		v2969 = v2907
		goto L439
	} else {
		goto L443
	}
L443:
	;
	v2926 = v2921
	v2929 = v2907
	goto L444
L444:
	;
	v2957 = *(*int32)(unsafe.Add(mBase, uint32(v2926)))
	*(*int32)(unsafe.Add(mBase, uint32(v2929))) = v2957
	v2960 = v2929 + int32(4)
	v2963 = F_hash_seq_search(m, v2911+int32(12))
	mBase = m.M
	v2964 = m.ExcPending
	if v2964 != 0 {
		goto L1
	} else {
		goto L446
	}
L445:
	;
	v2969 = v2960
	goto L439
L446:
	;
	if v2963 != 0 {
		v2926 = v2963
		v2929 = v2960
		goto L444
	} else {
		goto L447
	}
L447:
	;
	goto L445
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3057))) = int32(0)
	m.G0 = v2911 + int32(32)
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3090, int64(-65522), v2907)
	mBase = m.M
	v3093 = m.ExcPending
	if v3093 != 0 {
		goto L1
	} else {
		goto L457
	}
L449:
	;
	v3006 = v2911 + int32(12)
	F_hash_seq_init(m, v3006, v3002)
	mBase = m.M
	v3008 = m.ExcPending
	if v3008 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	v3009 = F_hash_seq_search(m, v3006)
	mBase = m.M
	v3010 = m.ExcPending
	if v3010 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	if v3009 == int32(0) {
		v3057 = v3000
		goto L448
	} else {
		goto L452
	}
L452:
	;
	v3014 = v3009
	v3017 = v3000
	goto L453
L453:
	;
	v3045 = *(*int32)(unsafe.Add(mBase, uint32(v3014)))
	*(*int32)(unsafe.Add(mBase, uint32(v3017))) = v3045
	v3048 = v3017 + int32(4)
	v3051 = F_hash_seq_search(m, v2911+int32(12))
	mBase = m.M
	v3052 = m.ExcPending
	if v3052 != 0 {
		goto L1
	} else {
		goto L455
	}
L454:
	;
	v3057 = v3048
	goto L448
L455:
	;
	if v3051 != 0 {
		v3014 = v3051
		v3017 = v3048
		goto L453
	} else {
		goto L456
	}
L456:
	;
	goto L454
L457:
	;
	v3094 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3095 = F_shm_toc_allocate(m, v3094, v1280)
	mBase = m.M
	v3096 = m.ExcPending
	if v3096 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v3098 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[46]))
	v3100 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	if v3100 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	v3122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3122, int64(-65521), v3095)
	mBase = m.M
	v3125 = m.ExcPending
	if v3125 != 0 {
		goto L1
	} else {
		goto L466
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3095)+4)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v3095))) = int32(-1)
	goto L459
L461:
	;
	goto L462
L462:
	;
	v3106 = F_strlen(m, v3100)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3095)+4)) = v3098
	*(*int32)(unsafe.Add(mBase, uint32(v3095))) = v3106
	if v3106 < int32(0) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	goto L459
L464:
	;
	v3112 = v3106 + int32(1)
	if v3112 == int32(0) {
		goto L463
	} else {
		goto L465
	}
L465:
	;
	v3118 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	base.MemoryCopy(m, v3095+int32(8), v3118, v3112)
	goto L463
L466:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3129 = F_palloc0(m, v3126<<(uint(int32(3))%32))
	mBase = m.M
	v3130 = m.ExcPending
	if v3130 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v3129
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3135 = F_mul_size(m, int32(_a_F_InitializeParallelDSM_12), v3134)
	mBase = m.M
	v3136 = m.ExcPending
	if v3136 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v3137 = F_shm_toc_allocate(m, v3132, v3135)
	mBase = m.M
	v3138 = m.ExcPending
	if v3138 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v3139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v3139 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v3144 = int32(0)
	goto L473
L471:
	;
	goto L472
L472:
	;
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3242, int64(-65534), v3137)
	mBase = m.M
	v3245 = m.ExcPending
	if v3245 != 0 {
		goto L1
	} else {
		goto L479
	}
L473:
	;
	v3177 = v3137 + v3144<<(uint(int32(14))%32)
	v3179 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3177)+16)) = v3179
	*(*int32)(unsafe.Add(mBase, uint32(v3177)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3177))) = v3179
	v3185 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v3177)+36)) = uint16(v3185)
	*(*int64)(unsafe.Add(mBase, uint32(v3177)+24)) = v3179
	*(*int32)(unsafe.Add(mBase, uint32(v3177)+32)) = int32(_a_F_InitializeParallelDSM_25)
	goto L475
L474:
	;
	goto L472
L475:
	;
	v3195 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[31]))
	F_shm_mq_set_receiver(m, v3177, v3195)
	mBase = m.M
	v3197 = m.ExcPending
	if v3197 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v3199 = F_shm_mq_attach(m, v3177, v3198)
	mBase = m.M
	v3200 = m.ExcPending
	if v3200 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	v3201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3201+v3144<<(uint(int32(3))%32))+4)) = v3199
	v3207 = v3144 + int32(1)
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3207 < v3208 {
		v3144 = v3207
		goto L473
	} else {
		goto L478
	}
L478:
	;
	goto L474
L479:
	;
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3247 = F_strlen(m, v3246)
	mBase = m.M
	v3248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3250 = F_strlen(m, v3249)
	mBase = m.M
	v3254 = F_shm_toc_allocate(m, v3248, v3250+v3247+int32(2))
	mBase = m.M
	v3255 = m.ExcPending
	if v3255 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if (v3256^v3254)&int32(3) != 0 {
		goto L484
	} else {
		goto L485
	}
L481:
	;
	v3333 = v3247 + v3254 + int32(1)
	v3334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if (v3334^v3333)&int32(3) != 0 {
		goto L505
	} else {
		goto L506
	}
L482:
	;
	goto L481
L483:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3311))) = uint8(v3310)
	if v3310&int32(255) == int32(0) {
		goto L482
	} else {
		goto L498
	}
L484:
	;
	v3262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3256))))
	v3309 = v3256
	v3310 = v3262
	v3311 = v3254
	goto L483
L485:
	;
	goto L486
L486:
	;
	if v3256&int32(3) != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v3266 = v3256
	v3268 = v3254
	goto L490
L488:
	;
	v3280 = v3256
	v3282 = v3254
	goto L489
L489:
	;
	v3284 = *(*int32)(unsafe.Add(mBase, uint32(v3280)))
	v3287 = int32(-2139062144)
	if (int32(16843008)-v3284|v3284)&v3287 != v3287 {
		v3309 = v3280
		v3310 = v3284
		v3311 = v3282
		goto L483
	} else {
		goto L494
	}
L490:
	;
	v3269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3266))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3268))) = uint8(v3269)
	if v3269 == int32(0) {
		goto L482
	} else {
		goto L492
	}
L491:
	;
	v3280 = v3276
	v3282 = v3274
	goto L489
L492:
	;
	v3273 = int32(1)
	v3274 = v3268 + v3273
	v3276 = v3266 + v3273
	if v3276&int32(3) != 0 {
		v3266 = v3276
		v3268 = v3274
		goto L490
	} else {
		goto L493
	}
L493:
	;
	goto L491
L494:
	;
	v3292 = v3280
	v3293 = v3284
	v3294 = v3282
	goto L495
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3294))) = v3293
	v3296 = int32(4)
	v3297 = v3294 + v3296
	v3299 = v3292 + v3296
	v3301 = *(*int32)(unsafe.Add(mBase, uint32(v3292)+4))
	v3304 = int32(-2139062144)
	if (int32(16843008)-v3301|v3301)&v3304 == v3304 {
		v3292 = v3299
		v3293 = v3301
		v3294 = v3297
		goto L495
	} else {
		goto L497
	}
L496:
	;
	v3309 = v3299
	v3310 = v3301
	v3311 = v3297
	goto L483
L497:
	;
	goto L496
L498:
	;
	v3318 = v3309
	v3320 = v3311
	goto L499
L499:
	;
	v3321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3318)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3320)+1)) = uint8(v3321)
	v3323 = int32(1)
	if v3321 != 0 {
		v3318 = v3318 + v3323
		v3320 = v3320 + v3323
		goto L499
	} else {
		goto L501
	}
L500:
	;
	goto L482
L501:
	;
	goto L500
L502:
	;
	v3409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3409, int64(-65527), v3254)
	mBase = m.M
	v3412 = m.ExcPending
	if v3412 != 0 {
		goto L1
	} else {
		goto L523
	}
L503:
	;
	goto L502
L504:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3389))) = uint8(v3388)
	if v3388&int32(255) == int32(0) {
		goto L503
	} else {
		goto L519
	}
L505:
	;
	v3340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3334))))
	v3387 = v3334
	v3388 = v3340
	v3389 = v3333
	goto L504
L506:
	;
	goto L507
L507:
	;
	if v3334&int32(3) != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3344 = v3334
	v3346 = v3333
	goto L511
L509:
	;
	v3358 = v3334
	v3360 = v3333
	goto L510
L510:
	;
	v3362 = *(*int32)(unsafe.Add(mBase, uint32(v3358)))
	v3365 = int32(-2139062144)
	if (int32(16843008)-v3362|v3362)&v3365 != v3365 {
		v3387 = v3358
		v3388 = v3362
		v3389 = v3360
		goto L504
	} else {
		goto L515
	}
L511:
	;
	v3347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3344))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3346))) = uint8(v3347)
	if v3347 == int32(0) {
		goto L503
	} else {
		goto L513
	}
L512:
	;
	v3358 = v3354
	v3360 = v3352
	goto L510
L513:
	;
	v3351 = int32(1)
	v3352 = v3346 + v3351
	v3354 = v3344 + v3351
	if v3354&int32(3) != 0 {
		v3344 = v3354
		v3346 = v3352
		goto L511
	} else {
		goto L514
	}
L514:
	;
	goto L512
L515:
	;
	v3370 = v3358
	v3371 = v3362
	v3372 = v3360
	goto L516
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3372))) = v3371
	v3374 = int32(4)
	v3375 = v3372 + v3374
	v3377 = v3370 + v3374
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3370)+4))
	v3382 = int32(-2139062144)
	if (int32(16843008)-v3379|v3379)&v3382 == v3382 {
		v3370 = v3377
		v3371 = v3379
		v3372 = v3375
		goto L516
	} else {
		goto L518
	}
L517:
	;
	v3387 = v3377
	v3388 = v3379
	v3389 = v3375
	goto L504
L518:
	;
	goto L517
L519:
	;
	v3396 = v3387
	v3398 = v3389
	goto L520
L520:
	;
	v3399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3396)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3398)+1)) = uint8(v3399)
	v3401 = int32(1)
	if v3399 != 0 {
		v3396 = v3396 + v3401
		v3398 = v3398 + v3401
		goto L520
	} else {
		goto L522
	}
L521:
	;
	goto L503
L522:
	;
	goto L521
L523:
	;
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3446 = v3413
	goto L213
L524:
	;
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_26), int32(0))
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(_a_F_InitializeParallelDSM_27), int32(_a_F_InitializeParallelDSM_28))
	mBase = m.M
	v3465 = m.ExcPending
	if v3465 != 0 {
		goto L1
	} else {
		goto L526
	}
L526:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v73 int32
	_ = v73
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
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
				v73 = int32(0)
				if v72|base.B2i32(v71 == v73) == v73 {
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v71)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v78
					v80 = *(*int64)(unsafe.Add(mBase, uint32(v71)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v80
					v82 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v82
					v84 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v84
					v86 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v86
					v88 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v88)
					F_pfree(m, v71)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						v92 = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v92
						*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v92
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_pfree(m, v96)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
							F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				} else {
					v92 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v92
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v92
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					F_pfree(m, v96)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
						F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
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
				v73 = int32(0)
				if v72|base.B2i32(v71 == v73) == v73 {
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v71)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v78
					v80 = *(*int64)(unsafe.Add(mBase, uint32(v71)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v80
					v82 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v82
					v84 = *(*int64)(unsafe.Add(mBase, uint32(v71)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v84
					v86 = *(*int64)(unsafe.Add(mBase, uint32(v71)))
					*(*int64)(unsafe.Add(mBase, uint32(v36))) = v86
					v88 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v88)
					F_pfree(m, v71)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						v92 = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v92
						*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v92
						v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
						F_pfree(m, v96)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
							F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								m.G0 = v10 + int32(48)
								return
							}
						}
					}
				} else {
					v92 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v92
					*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v92
					v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					F_pfree(m, v96)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = int32(0)
						F_pgstat_progress_parallel_incr_param(m, int32(9), int64(1))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return
						} else {
							m.G0 = v10 + int32(48)
							return
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
				F_errmsg_internal(m, int32(_a_F_parallel_vacuum_process_one_index_0), v10)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_parallel_vacuum_process_one_index_1), int32(904), int32(_a_F_parallel_vacuum_process_one_index_2))
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
