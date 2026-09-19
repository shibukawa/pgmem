package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v190 int32
	_ = v190
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
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
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
	v279 = v158
	goto L30
L30:
	;
	F_LWLockRelease(m, v279)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L53
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
	v279 = v269
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
	v190 = int32(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBuckets[0]))
	if v264 != 0 {
		goto L47
	} else {
		goto L48
	}
L38:
	;
	v200 = v190 + (v177 + v184)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200)+4))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v209 = v201 + v202&(v203-int32(1))<<(uint(int32(2))%32)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v210
	v212 = v190 + (v166 + v184)
	v214 = base.AtomicRmwCmpxchg32(m, v209, int32(0), v210, v212)
	if v210 != v214 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L37
L40:
	;
	v220 = v214
	goto L43
L41:
	;
	goto L42
L42:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v200)+8))
	v249 = (v244+int32(15))&int32(-8) + v190
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v177)+8))
	if base.Ui32(v249) < base.Ui32(v250) {
		v190 = v249
		goto L38
	} else {
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v220
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v228
	v231 = base.AtomicRmwCmpxchg32(m, v209, int32(0), v228, v212)
	if v228 != v231 {
		v220 = v231
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L42
L45:
	;
	goto L44
L46:
	;
	goto L39
L47:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v269 = v267 + int32(40)
	v271 = F_LWLockAcquire(m, v269, int32(0))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L5
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v267)+24))
	if v275 != 0 {
		v166 = v275
		v167 = v267 + int32(24)
		v168 = v269
		goto L31
	} else {
		goto L52
	}
L52:
	;
	goto L32
L53:
	;
	v290 = F_BarrierArriveAndWait(m, v14, int32(134217756))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
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
			v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
			mBase = m.M
			v191 = m.ExcPending
			if v191 != 0 {
				return int32(0)
			} else {
				return v190
			}
		default:
			v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
			mBase = m.M
			v191 = m.ExcPending
			if v191 != 0 {
				return int32(0)
			} else {
				return v190
			}
		case 6:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+36)))
			if v16 != int32(1) {
				v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					return v190
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
					v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						return v190
					}
				}
			}
		case 8:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+36)))
			if v29 != int32(1) {
				v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					return v190
				}
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				F_index_parallelrescan(m, v32)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						return v190
					}
				}
			}
		case 9:
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+36)))
			if v36 != int32(1) {
				v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					return v190
				}
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
				F_index_parallelrescan(m, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						return v190
					}
				}
			}
		case 11:
			v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+36)))
			if v85 != int32(1) {
				v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					return v190
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
								v109 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
								if v109 == int32(0) {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
									if v122 == int32(0) {
										F_dsa_free(m, v89, v93)
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
											v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return int32(0)
											} else {
												return v190
											}
										}
									} else {
										v125 = F_dsa_get_address(m, v89, v122)
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											v127 = int32(1)
											v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
											if v129 != v127 {
												F_dsa_free(m, v89, v93)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
													v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return int32(0)
													} else {
														return v190
													}
												}
											} else {
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
												F_dsa_free(m, v89, v132)
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													F_dsa_free(m, v89, v93)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
														v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															return v190
														}
													}
												}
											}
										}
									}
								} else {
									v112 = F_dsa_get_address(m, v89, v109)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v114 = int32(1)
										v116 = base.AtomicRmwSub32(m, v112, int32(0), v114)
										if v116 != v114 {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
											if v122 == int32(0) {
												F_dsa_free(m, v89, v93)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
													v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return int32(0)
													} else {
														return v190
													}
												}
											} else {
												v125 = F_dsa_get_address(m, v89, v122)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													v127 = int32(1)
													v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
													if v129 != v127 {
														F_dsa_free(m, v89, v93)
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
															v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v191 = m.ExcPending
															if v191 != 0 {
																return int32(0)
															} else {
																return v190
															}
														}
													} else {
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														F_dsa_free(m, v89, v132)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return int32(0)
														} else {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v191 = m.ExcPending
																if v191 != 0 {
																	return int32(0)
																} else {
																	return v190
																}
															}
														}
													}
												}
											}
										} else {
											v119 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
											F_dsa_free(m, v89, v119)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int32(0)
											} else {
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
												if v122 == int32(0) {
													F_dsa_free(m, v89, v93)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
														v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															return v190
														}
													}
												} else {
													v125 = F_dsa_get_address(m, v89, v122)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v127 = int32(1)
														v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
														if v129 != v127 {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v191 = m.ExcPending
																if v191 != 0 {
																	return int32(0)
																} else {
																	return v190
																}
															}
														} else {
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															F_dsa_free(m, v89, v132)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return int32(0)
															} else {
																F_dsa_free(m, v89, v93)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																	v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v191 = m.ExcPending
																	if v191 != 0 {
																		return int32(0)
																	} else {
																		return v190
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
									v101 = int32(1)
									v103 = base.AtomicRmwSub32(m, v99, int32(0), v101)
									if v103 != v101 {
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
										if v109 == int32(0) {
											v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
											if v122 == int32(0) {
												F_dsa_free(m, v89, v93)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
													v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return int32(0)
													} else {
														return v190
													}
												}
											} else {
												v125 = F_dsa_get_address(m, v89, v122)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													v127 = int32(1)
													v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
													if v129 != v127 {
														F_dsa_free(m, v89, v93)
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
															v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v191 = m.ExcPending
															if v191 != 0 {
																return int32(0)
															} else {
																return v190
															}
														}
													} else {
														v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														F_dsa_free(m, v89, v132)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return int32(0)
														} else {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v191 = m.ExcPending
																if v191 != 0 {
																	return int32(0)
																} else {
																	return v190
																}
															}
														}
													}
												}
											}
										} else {
											v112 = F_dsa_get_address(m, v89, v109)
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return int32(0)
											} else {
												v114 = int32(1)
												v116 = base.AtomicRmwSub32(m, v112, int32(0), v114)
												if v116 != v114 {
													v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
													if v122 == int32(0) {
														F_dsa_free(m, v89, v93)
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
															v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
															mBase = m.M
															v191 = m.ExcPending
															if v191 != 0 {
																return int32(0)
															} else {
																return v190
															}
														}
													} else {
														v125 = F_dsa_get_address(m, v89, v122)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return int32(0)
														} else {
															v127 = int32(1)
															v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
															if v129 != v127 {
																F_dsa_free(m, v89, v93)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																	v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v191 = m.ExcPending
																	if v191 != 0 {
																		return int32(0)
																	} else {
																		return v190
																	}
																}
															} else {
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																F_dsa_free(m, v89, v132)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return int32(0)
																} else {
																	F_dsa_free(m, v89, v93)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																		v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v191 = m.ExcPending
																		if v191 != 0 {
																			return int32(0)
																		} else {
																			return v190
																		}
																	}
																}
															}
														}
													}
												} else {
													v119 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
													F_dsa_free(m, v89, v119)
													mBase = m.M
													v121 = m.ExcPending
													if v121 != 0 {
														return int32(0)
													} else {
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														if v122 == int32(0) {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v191 = m.ExcPending
																if v191 != 0 {
																	return int32(0)
																} else {
																	return v190
																}
															}
														} else {
															v125 = F_dsa_get_address(m, v89, v122)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																v127 = int32(1)
																v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
																if v129 != v127 {
																	F_dsa_free(m, v89, v93)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																		v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v191 = m.ExcPending
																		if v191 != 0 {
																			return int32(0)
																		} else {
																			return v190
																		}
																	}
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																	F_dsa_free(m, v89, v132)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return int32(0)
																	} else {
																		F_dsa_free(m, v89, v93)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																			v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v191 = m.ExcPending
																			if v191 != 0 {
																				return int32(0)
																			} else {
																				return v190
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
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
										F_dsa_free(m, v89, v106)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
											if v109 == int32(0) {
												v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
												if v122 == int32(0) {
													F_dsa_free(m, v89, v93)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
														v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
														mBase = m.M
														v191 = m.ExcPending
														if v191 != 0 {
															return int32(0)
														} else {
															return v190
														}
													}
												} else {
													v125 = F_dsa_get_address(m, v89, v122)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return int32(0)
													} else {
														v127 = int32(1)
														v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
														if v129 != v127 {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v191 = m.ExcPending
																if v191 != 0 {
																	return int32(0)
																} else {
																	return v190
																}
															}
														} else {
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															F_dsa_free(m, v89, v132)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return int32(0)
															} else {
																F_dsa_free(m, v89, v93)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																	v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v191 = m.ExcPending
																	if v191 != 0 {
																		return int32(0)
																	} else {
																		return v190
																	}
																}
															}
														}
													}
												}
											} else {
												v112 = F_dsa_get_address(m, v89, v109)
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return int32(0)
												} else {
													v114 = int32(1)
													v116 = base.AtomicRmwSub32(m, v112, int32(0), v114)
													if v116 != v114 {
														v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
														if v122 == int32(0) {
															F_dsa_free(m, v89, v93)
															mBase = m.M
															v136 = m.ExcPending
															if v136 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																mBase = m.M
																v191 = m.ExcPending
																if v191 != 0 {
																	return int32(0)
																} else {
																	return v190
																}
															}
														} else {
															v125 = F_dsa_get_address(m, v89, v122)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																v127 = int32(1)
																v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
																if v129 != v127 {
																	F_dsa_free(m, v89, v93)
																	mBase = m.M
																	v136 = m.ExcPending
																	if v136 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																		v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																		mBase = m.M
																		v191 = m.ExcPending
																		if v191 != 0 {
																			return int32(0)
																		} else {
																			return v190
																		}
																	}
																} else {
																	v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																	F_dsa_free(m, v89, v132)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return int32(0)
																	} else {
																		F_dsa_free(m, v89, v93)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																			v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v191 = m.ExcPending
																			if v191 != 0 {
																				return int32(0)
																			} else {
																				return v190
																			}
																		}
																	}
																}
															}
														}
													} else {
														v119 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
														F_dsa_free(m, v89, v119)
														mBase = m.M
														v121 = m.ExcPending
														if v121 != 0 {
															return int32(0)
														} else {
															v122 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
															if v122 == int32(0) {
																F_dsa_free(m, v89, v93)
																mBase = m.M
																v136 = m.ExcPending
																if v136 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																	v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																	mBase = m.M
																	v191 = m.ExcPending
																	if v191 != 0 {
																		return int32(0)
																	} else {
																		return v190
																	}
																}
															} else {
																v125 = F_dsa_get_address(m, v89, v122)
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int32(0)
																} else {
																	v127 = int32(1)
																	v129 = base.AtomicRmwSub32(m, v125, int32(0), v127)
																	if v129 != v127 {
																		F_dsa_free(m, v89, v93)
																		mBase = m.M
																		v136 = m.ExcPending
																		if v136 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																			v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																			mBase = m.M
																			v191 = m.ExcPending
																			if v191 != 0 {
																				return int32(0)
																			} else {
																				return v190
																			}
																		}
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
																		F_dsa_free(m, v89, v132)
																		mBase = m.M
																		v134 = m.ExcPending
																		if v134 != 0 {
																			return int32(0)
																		} else {
																			F_dsa_free(m, v89, v93)
																			mBase = m.M
																			v136 = m.ExcPending
																			if v136 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v90))) = int32(0)
																				v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
																				mBase = m.M
																				v191 = m.ExcPending
																				if v191 != 0 {
																					return int32(0)
																				} else {
																					return v190
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
						v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
						mBase = m.M
						v191 = m.ExcPending
						if v191 != 0 {
							return int32(0)
						} else {
							return v190
						}
					}
				} else {
					v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						return v190
					}
				}
			}
		case 21:
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+36)))
			if v43 != int32(1) {
				v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					return v190
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
							v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
							mBase = m.M
							v191 = m.ExcPending
							if v191 != 0 {
								return int32(0)
							} else {
								return v190
							}
						}
					}
				} else {
					v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						return v190
					}
				}
			}
		case 22:
			v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+36)))
			if v70 != int32(1) {
				v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					return v190
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
							v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
							mBase = m.M
							v191 = m.ExcPending
							if v191 != 0 {
								return int32(0)
							} else {
								return v190
							}
						}
					}
				} else {
					v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						return v190
					}
				}
			}
		case 26:
			v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+36)))
			if v146 != int32(1) {
				v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
				mBase = m.M
				v191 = m.ExcPending
				if v191 != 0 {
					return int32(0)
				} else {
					return v190
				}
			} else {
				v149 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				if v149 != 0 {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
					v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v152 = int64(*(*int32)(unsafe.Add(mBase, uint32(v151)+40)))
					v154 = F_shm_toc_lookup(m, v150, v152, int32(0))
					mBase = m.M
					v155 = m.ExcPending
					if v155 != 0 {
						return int32(0)
					} else {
						v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						if v156 != 0 {
							F_ExecHashTableDetachBatch(m, v156)
							mBase = m.M
							v158 = m.ExcPending
							if v158 != 0 {
								return int32(0)
							} else {
								v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								F_ExecHashTableDetach(m, v159)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int32(0)
								} else {
									F_FileSetDeleteAll(m, v154+int32(168))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return int32(0)
									} else {
										v167 = v154 + int32(56)
										v168 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167))), uint32(v168))
										*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v168
										*(*uint8)(unsafe.Add(mBase, uint32(v167)+20)) = uint8(v168)
										*(*int64)(unsafe.Add(mBase, uint32(v167)+12)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v168
										F_ConditionVariableInit(m, v154+int32(80))
										mBase = m.M
										v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											return v190
										}
									}
								}
							}
						} else {
							F_FileSetDeleteAll(m, v154+int32(168))
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return int32(0)
							} else {
								v167 = v154 + int32(56)
								v168 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v167))), uint32(v168))
								*(*int32)(unsafe.Add(mBase, uint32(v167)+8)) = v168
								*(*uint8)(unsafe.Add(mBase, uint32(v167)+20)) = uint8(v168)
								*(*int64)(unsafe.Add(mBase, uint32(v167)+12)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v168
								F_ConditionVariableInit(m, v154+int32(80))
								mBase = m.M
								v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return int32(0)
								} else {
									return v190
								}
							}
						}
					}
				} else {
					v190 = F_planstate_tree_walker_impl(m, l0, int32(628), l1)
					mBase = m.M
					v191 = m.ExcPending
					if v191 != 0 {
						return int32(0)
					} else {
						return v190
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
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v604 int32
	_ = v604
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v657 int32
	_ = v657
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
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v733 int32
	_ = v733
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
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
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v893 int32
	_ = v893
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
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
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
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
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
	var v988 int32
	_ = v988
	var v1020 int32
	_ = v1020
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
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
	var v1097 int32
	_ = v1097
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1132 int32
	_ = v1132
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
	var v1167 int32
	_ = v1167
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1239 int32
	_ = v1239
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
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
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1344 int32
	_ = v1344
	var v1347 int32
	_ = v1347
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1361 int32
	_ = v1361
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1388 int64
	_ = v1388
	var v1391 int64
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1446 int32
	_ = v1446
	var v1452 int32
	_ = v1452
	var v1456 int32
	_ = v1456
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1483 int32
	_ = v1483
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1526 int32
	_ = v1526
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1538 int32
	_ = v1538
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1559 int32
	_ = v1559
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1573 int32
	_ = v1573
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1631 int32
	_ = v1631
	var v1634 int32
	_ = v1634
	var v1664 int32
	_ = v1664
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1684 int32
	_ = v1684
	var v1687 int32
	_ = v1687
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 float64
	_ = v1713
	var v1725 int32
	_ = v1725
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1782 int32
	_ = v1782
	var v1785 int32
	_ = v1785
	var v1795 int32
	_ = v1795
	var v1830 int32
	_ = v1830
	var v1863 int32
	_ = v1863
	var v1865 int32
	_ = v1865
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1880 int32
	_ = v1880
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1886 int32
	_ = v1886
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1906 int32
	_ = v1906
	var v1912 int32
	_ = v1912
	var v1914 int32
	_ = v1914
	var v1915 int32
	_ = v1915
	var v1921 int32
	_ = v1921
	var v1952 int32
	_ = v1952
	var v1956 int32
	_ = v1956
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2035 int32
	_ = v2035
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2046 int32
	_ = v2046
	var v2047 int32
	_ = v2047
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2056 int32
	_ = v2056
	var v2064 int32
	_ = v2064
	var v2069 int32
	_ = v2069
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2079 int32
	_ = v2079
	var v2082 int32
	_ = v2082
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2096 int32
	_ = v2096
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
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
	var v2117 int32
	_ = v2117
	var v2122 int32
	_ = v2122
	var v2127 int32
	_ = v2127
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2138 int32
	_ = v2138
	var v2141 int32
	_ = v2141
	var v2144 int32
	_ = v2144
	var v2146 int32
	_ = v2146
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2157 int32
	_ = v2157
	var v2158 int64
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2172 int32
	_ = v2172
	var v2176 int32
	_ = v2176
	var v2181 int32
	_ = v2181
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2197 int32
	_ = v2197
	var v2200 int32
	_ = v2200
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2209 int32
	_ = v2209
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
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2224 int64
	_ = v2224
	var v2227 int32
	_ = v2227
	var v2228 int64
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2234 int32
	_ = v2234
	var v2239 int32
	_ = v2239
	var v2245 int32
	_ = v2245
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2279 int32
	_ = v2279
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2287 int32
	_ = v2287
	var v2290 int32
	_ = v2290
	var v2291 int32
	_ = v2291
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2327 int32
	_ = v2327
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2339 int32
	_ = v2339
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2347 int32
	_ = v2347
	var v2349 int32
	_ = v2349
	var v2385 int32
	_ = v2385
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2428 int32
	_ = v2428
	var v2429 int32
	_ = v2429
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2440 int32
	_ = v2440
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
	var v2475 int32
	_ = v2475
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2509 int32
	_ = v2509
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2527 int32
	_ = v2527
	var v2557 int32
	_ = v2557
	var v2558 int32
	_ = v2558
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2596 int32
	_ = v2596
	var v2601 int32
	_ = v2601
	var v2629 int32
	_ = v2629
	var v2634 int32
	_ = v2634
	var v2635 int32
	_ = v2635
	var v2636 int32
	_ = v2636
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2707 int32
	_ = v2707
	var v2709 int64
	_ = v2709
	var v2711 int32
	_ = v2711
	var v2712 int32
	_ = v2712
	var v2715 int32
	_ = v2715
	var v2716 int32
	_ = v2716
	var v2720 int32
	_ = v2720
	var v2750 int32
	_ = v2750
	var v2754 int32
	_ = v2754
	var v2790 int32
	_ = v2790
	var v2793 int32
	_ = v2793
	var v2794 int32
	_ = v2794
	var v2795 int32
	_ = v2795
	var v2796 int32
	_ = v2796
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2805 int32
	_ = v2805
	var v2806 int32
	_ = v2806
	var v2808 int32
	_ = v2808
	var v2816 int32
	_ = v2816
	var v2846 int32
	_ = v2846
	var v2848 int32
	_ = v2848
	var v2850 int32
	_ = v2850
	var v2853 int32
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2890 int32
	_ = v2890
	var v2893 int32
	_ = v2893
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2896 int32
	_ = v2896
	var v2898 int32
	_ = v2898
	var v2905 int32
	_ = v2905
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2910 int32
	_ = v2910
	var v2911 int32
	_ = v2911
	var v2912 int32
	_ = v2912
	var v2914 int32
	_ = v2914
	var v2917 int32
	_ = v2917
	var v2921 int32
	_ = v2921
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2929 int32
	_ = v2929
	var v2932 int32
	_ = v2932
	var v2960 int32
	_ = v2960
	var v2963 int32
	_ = v2963
	var v2966 int32
	_ = v2966
	var v2967 int32
	_ = v2967
	var v2972 int32
	_ = v2972
	var v3000 int32
	_ = v3000
	var v3003 int32
	_ = v3003
	var v3005 int32
	_ = v3005
	var v3009 int32
	_ = v3009
	var v3011 int32
	_ = v3011
	var v3012 int32
	_ = v3012
	var v3013 int32
	_ = v3013
	var v3017 int32
	_ = v3017
	var v3020 int32
	_ = v3020
	var v3048 int32
	_ = v3048
	var v3051 int32
	_ = v3051
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3060 int32
	_ = v3060
	var v3093 int32
	_ = v3093
	var v3096 int32
	_ = v3096
	var v3097 int32
	_ = v3097
	var v3098 int32
	_ = v3098
	var v3099 int32
	_ = v3099
	var v3101 int32
	_ = v3101
	var v3103 int32
	_ = v3103
	var v3109 int32
	_ = v3109
	var v3115 int32
	_ = v3115
	var v3121 int32
	_ = v3121
	var v3125 int32
	_ = v3125
	var v3128 int32
	_ = v3128
	var v3129 int32
	_ = v3129
	var v3132 int32
	_ = v3132
	var v3133 int32
	_ = v3133
	var v3135 int32
	_ = v3135
	var v3137 int32
	_ = v3137
	var v3138 int32
	_ = v3138
	var v3139 int32
	_ = v3139
	var v3140 int32
	_ = v3140
	var v3141 int32
	_ = v3141
	var v3142 int32
	_ = v3142
	var v3147 int32
	_ = v3147
	var v3180 int32
	_ = v3180
	var v3182 int32
	_ = v3182
	var v3185 int64
	_ = v3185
	var v3189 int32
	_ = v3189
	var v3199 int32
	_ = v3199
	var v3201 int32
	_ = v3201
	var v3202 int32
	_ = v3202
	var v3203 int32
	_ = v3203
	var v3204 int32
	_ = v3204
	var v3205 int32
	_ = v3205
	var v3211 int32
	_ = v3211
	var v3212 int32
	_ = v3212
	var v3246 int32
	_ = v3246
	var v3249 int32
	_ = v3249
	var v3250 int32
	_ = v3250
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3253 int32
	_ = v3253
	var v3254 int32
	_ = v3254
	var v3258 int32
	_ = v3258
	var v3259 int32
	_ = v3259
	var v3260 int32
	_ = v3260
	var v3266 int32
	_ = v3266
	var v3270 int32
	_ = v3270
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3277 int32
	_ = v3277
	var v3278 int32
	_ = v3278
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3288 int32
	_ = v3288
	var v3291 int32
	_ = v3291
	var v3296 int32
	_ = v3296
	var v3297 int32
	_ = v3297
	var v3298 int32
	_ = v3298
	var v3300 int32
	_ = v3300
	var v3301 int32
	_ = v3301
	var v3303 int32
	_ = v3303
	var v3305 int32
	_ = v3305
	var v3308 int32
	_ = v3308
	var v3313 int32
	_ = v3313
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3325 int32
	_ = v3325
	var v3327 int32
	_ = v3327
	var v3337 int32
	_ = v3337
	var v3338 int32
	_ = v3338
	var v3344 int32
	_ = v3344
	var v3348 int32
	_ = v3348
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3355 int32
	_ = v3355
	var v3356 int32
	_ = v3356
	var v3358 int32
	_ = v3358
	var v3362 int32
	_ = v3362
	var v3364 int32
	_ = v3364
	var v3366 int32
	_ = v3366
	var v3369 int32
	_ = v3369
	var v3374 int32
	_ = v3374
	var v3375 int32
	_ = v3375
	var v3376 int32
	_ = v3376
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3381 int32
	_ = v3381
	var v3383 int32
	_ = v3383
	var v3386 int32
	_ = v3386
	var v3391 int32
	_ = v3391
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3400 int32
	_ = v3400
	var v3402 int32
	_ = v3402
	var v3403 int32
	_ = v3403
	var v3405 int32
	_ = v3405
	var v3413 int32
	_ = v3413
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3450 int32
	_ = v3450
	var v3460 int32
	_ = v3460
	var v3464 int32
	_ = v3464
	var v3469 int32
	_ = v3469
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
	v1306 = F_shm_toc_estimate(m, v55)
	mBase = m.M
	v1307 = m.ExcPending
	if v1307 != 0 {
		goto L1
	} else {
		goto L191
	}
L7:
	;
	v69 = l0 + int32(12)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v70 <= int32(0) {
		v1275 = v2
		v1277 = v69
		v1278 = v2
		v1279 = v2
		v1281 = v2
		v1286 = v2
		v1296 = v2
		v1297 = v2
		v1298 = v2
		v1299 = v2
		v1300 = v2
		v1301 = v2
		v1305 = v2
		goto L6
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1275 = v2
	v1277 = l0 + int32(12)
	v1278 = v2
	v1279 = v2
	v1281 = v2
	v1286 = v2
	v1296 = v2
	v1297 = v2
	v1298 = v2
	v1299 = v2
	v1300 = v2
	v1301 = v2
	v1305 = v2
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
	if v366 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v366 = v80
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
	v366 = int32(0)
	goto L13
L24:
	;
	goto L25
L25:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = int64(2880502729)
	v118 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v116)+8)), uint32(v118))
	*(*int64)(unsafe.Add(mBase, uint32(v116)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v105 & int32(-32)
	goto L26
L26:
	;
	v127 = F_shm_toc_allocate(m, v116, int32(_a_F_InitializeParallelDSM_1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v131 = F_dsa_create_in_place_ext(m, v127, int32(_a_F_InitializeParallelDSM_1), int32(72), v108)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_shm_toc_insert(m, v116, int64(-65535), v127)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v137 = F_shm_toc_allocate(m, v116, int32(12))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v139 = m.G0
	v141 = v139 - int32(16)
	m.G0 = v141
	v143 = int32(_a_F_InitializeParallelDSM_0)
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1]))
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v147
	v150 = F_dshash_create(m, v131, int32(_a_F_InitializeParallelDSM_2), v131)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v154 = F_dshash_create(m, v131, int32(_a_F_InitializeParallelDSM_3), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v144
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v150)+32))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v154)+32))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = v162
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = v165
	if int32(0) < v165 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_shm_toc_insert(m, v116, int64(-65534), v137)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L61
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L58
	}
L37:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[9]))
	v182 = v165
	v186 = v170
	v189 = v2
	goto L40
L38:
	;
	goto L39
L39:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v294)+16)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v294)+12)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v294)+8)) = v137
	F_on_dsm_detach(m, v108, int32(1607), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L57
	}
L40:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v186+v189<<(uint(int32(4))%32))+8))
	if v206 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L39
L42:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	v213 = F_dsa_allocate_extended(m, v131, v207*int32(116)+int32(20), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	v254 = v182
	v255 = v186
	goto L44
L44:
	;
	v259 = v189 + int32(1)
	if v259 < v254 {
		v182 = v254
		v186 = v255
		v189 = v259
		goto L40
	} else {
		goto L56
	}
L45:
	;
	v215 = F_dsa_get_address(m, v131, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_TupleDescCopy(m, v215, v206)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v215)+8)) = v189
	v223 = v141 + int32(7)
	v224 = F_dshash_find_or_insert(m, v154, v206+int32(8), v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+7)))
	if v226 == int32(1) {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v206)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v224)+4)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v229
	F_dshash_release_lock(m, v154, v224)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v206
	v235 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+12)) = uint8(v235)
	v239 = F_dshash_find_or_insert(m, v150, v141+int32(8), v223)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+7)))
	if v241 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v239))) = v213
	v245 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v239)+4)) = uint8(v245)
	goto L54
L53:
	;
	goto L54
L54:
	;
	F_dshash_release_lock(m, v150, v239)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v250 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[9]))
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[8]))
	v254 = v252
	v255 = v250
	goto L44
L56:
	;
	goto L41
L57:
	;
	m.G0 = v141 + int32(16)
	goto L35
L58:
	;
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_4), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_5), int32(2253), int32(_a_F_InitializeParallelDSM_6))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
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
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_dsa_pin_mapping(m, v131)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v325 = int32(_a_F_InitializeParallelDSM_7)
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v326))) = v108
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v329)+4)) = v131
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v82
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v366 = v333
	goto L13
L64:
	;
	v372 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v372
	v1275 = v2
	v1277 = v69
	v1278 = v2
	v1279 = v372
	v1281 = v2
	v1286 = v2
	v1296 = v2
	v1297 = v2
	v1298 = v2
	v1299 = v2
	v1300 = v2
	v1301 = v2
	v1305 = v2
	goto L6
L65:
	;
	goto L66
L66:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v375 <= int32(0) {
		v1275 = v2
		v1277 = v69
		v1278 = v2
		v1279 = v366
		v1281 = v2
		v1286 = v2
		v1296 = v2
		v1297 = v2
		v1298 = v2
		v1299 = v2
		v1300 = v2
		v1301 = v2
		v1305 = v2
		goto L6
	} else {
		goto L67
	}
L67:
	;
	v378 = int32(1)
	v380 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[10]))
	if v380 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v382 = v380
	v385 = v378
	goto L71
L69:
	;
	v425 = v378
	goto L70
L70:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v458 = F_add_size(m, v453, (v425+int32(31))&int32(-32))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L75
	}
L71:
	;
	v415 = F_strlen(m, v382+int32(24))
	mBase = m.M
	v418 = F_add_size(m, v385, v415+int32(1))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	v425 = v418
	goto L70
L73:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	if v420 != 0 {
		v382 = v420
		v385 = v418
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v458
	v461 = m.G0
	v463 = v461 - int32(16)
	m.G0 = v463
	v465 = int32(4)
	v467 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[11]))
	v468 = int32(0)
	if base.B2i32(v467 == v468)|base.B2i32(v467 == int32(_a_F_InitializeParallelDSM_8)) == v468 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v809 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v814 = F_add_size(m, v809, (v733+int32(31))&int32(-32))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L1
	} else {
		goto L123
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L120
	}
L78:
	;
	v477 = v467
	v481 = v465
	goto L81
L79:
	;
	v733 = v465
	goto L80
L80:
	;
	m.G0 = v463 + int32(16)
	goto L76
L81:
	;
	v507 = int32(0)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v477-int32(60))))
	if base.Ui32(v510) < base.Ui32(int32(2)) {
		v691 = v507
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v733 = v722
	goto L80
L83:
	;
	v722 = F_add_size(m, v481, v691)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L1
	} else {
		goto L118
	}
L84:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v477-int32(32))))
	if v515 == int32(0) {
		v691 = v507
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v519 = v477 + int32(-64)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v521 = F_strlen(m, v520)
	mBase = m.M
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v477-int32(40))))
	switch v524 {
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
		v626 = v507
		goto L86
	}
L86:
	;
	v657 = int32(1)
	v661 = F_add_size(m, v521+v657, v626+v657)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L105
	}
L87:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v477)+28))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v477)+36))
	if v545 == int32(0) {
		goto L77
	} else {
		goto L96
	}
L88:
	;
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v477)+28))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v538)))
	if v539 == int32(0) {
		v626 = v507
		goto L86
	} else {
		goto L95
	}
L89:
	;
	v626 = int32(25)
	goto L86
L90:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v477)+28))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	v531 = v529 >> (uint(int32(31)) % 32)
	if v529^v531-v531 < int32(1000) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v626 = int32(5)
	goto L86
L92:
	;
	v536 = int32(4)
	goto L94
L93:
	;
	v536 = int32(11)
	goto L94
L94:
	;
	v626 = v536
	goto L86
L95:
	;
	v542 = F_strlen(m, v539)
	mBase = m.M
	v626 = v542
	goto L86
L96:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	if v548 == int32(0) {
		goto L77
	} else {
		goto L97
	}
L97:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	if v544 != v551 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v554 = v545
	goto L101
L99:
	;
	v604 = v548
	goto L100
L100:
	;
	v624 = F_strlen(m, v604)
	mBase = m.M
	v626 = v624
	goto L86
L101:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v554)+12))
	if v585 == int32(0) {
		goto L77
	} else {
		goto L103
	}
L102:
	;
	v604 = v585
	goto L100
L103:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v554)+16))
	if v544 != v588 {
		v554 = v554 + int32(12)
		goto L101
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v477)+20))
	if v663 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v664 = F_strlen(m, v663)
	mBase = m.M
	v665 = F_add_size(m, v661, v664)
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	v667 = v661
	goto L108
L108:
	;
	v669 = F_add_size(m, v667, int32(1))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	v667 = v665
	goto L108
L110:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v477)+20))
	if v671 == int32(0) {
		v680 = v669
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v682 = F_add_size(m, v680, int32(4))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L1
	} else {
		goto L115
	}
L112:
	;
	v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671))))
	if v674 == int32(0) {
		v680 = v669
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v678 = F_add_size(m, v669, int32(4))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v680 = v678
	goto L111
L115:
	;
	v685 = F_add_size(m, v682, int32(4))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v688 = F_add_size(m, v685, int32(4))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v691 = v688
	goto L83
L118:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v477)+4))
	if v724 != int32(_a_F_InitializeParallelDSM_8) {
		v477 = v724
		v481 = v722
		goto L81
	} else {
		goto L119
	}
L119:
	;
	goto L82
L120:
	;
	v798 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	*(*int32)(unsafe.Add(mBase, uint32(v463)+4)) = v798
	*(*int32)(unsafe.Add(mBase, uint32(v463))) = v544
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_9), v463)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(3036), int32(_a_F_InitializeParallelDSM_11))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v814
	v820 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[12]))
	v821 = F_mul_size(m, int32(8), v820)
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	v823 = F_add_size(m, int32(4), v821)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v830 = F_add_size(m, v825, (v823+int32(31))&int32(-32))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v830
	v834 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[13]))
	if int32(2) <= v834 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	v837 = F_EstimateSnapshotSpace(m, v33)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	v847 = v2
	goto L129
L129:
	;
	v848 = F_EstimateSnapshotSpace(m, v37)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L1
	} else {
		goto L132
	}
L130:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v844 = F_add_size(m, v839, (v837+int32(31))&int32(-32))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v844
	v847 = v837
	goto L129
L132:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v855 = F_add_size(m, v850, (v848+int32(31))&int32(-32))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v855
	v858 = int32(0)
	v860 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[14]))
	if v860 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v863 = v858
	v867 = v860
	goto L137
L135:
	;
	v904 = v858
	goto L136
L136:
	;
	v936 = F_mul_size(m, int32(4), v904)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L1
	} else {
		goto L145
	}
L137:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	if v893 != 0 {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v904 = v899
	goto L136
L139:
	;
	v895 = F_add_size(m, v863, int32(1))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	v897 = v863
	goto L141
L141:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v867)+52))
	v899 = F_add_size(m, v897, v898)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L143
	}
L142:
	;
	v897 = v895
	goto L141
L143:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v867)+80))
	if v901 != 0 {
		v863 = v899
		v867 = v901
		goto L137
	} else {
		goto L144
	}
L144:
	;
	goto L138
L145:
	;
	v938 = F_add_size(m, int32(32), v936)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v945 = F_add_size(m, v940, (v938+int32(31))&int32(-32))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v945
	v949 = F_add_size(m, v945, int32(32))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v949
	v953 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[15]))
	if v953 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v953)))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v955)+4))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v955)+412))
	if v957 != 0 {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v1024 = int32(1)
	goto L151
L151:
	;
	v1026 = F_mul_size(m, v1024, int32(12))
	mBase = m.M
	v1027 = m.ExcPending
	if v1027 != 0 {
		goto L1
	} else {
		goto L156
	}
L152:
	;
	v1024 = v1020 + int32(1)
	goto L151
L153:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v955)+376))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v955)+364))
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v955)+352))
	v961 = *(*int32)(unsafe.Add(mBase, uint32(v955)+340))
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v955)+328))
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v955)+316))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v955)+304))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v955)+292))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v955)+280))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v955)+268))
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v955)+256))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v955)+244))
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v955)+232))
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v955)+220))
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v955)+208))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v955)+196))
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v955)+184))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v955)+172))
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v955)+160))
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v955)+148))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v955)+136))
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v955)+124))
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v955)+112))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v955)+100))
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v955)+88))
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v955)+76))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v955)+64))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v955)+52))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v955)+40))
	v987 = *(*int32)(unsafe.Add(mBase, uint32(v955)+28))
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v955)+16))
	v1020 = v958 + (v959 + (v960 + (v961 + (v962 + (v963 + (v964 + (v965 + (v966 + (v967 + (v968 + (v969 + (v970 + (v971 + (v972 + (v973 + (v974 + (v975 + (v976 + (v977 + (v978 + (v979 + (v980 + (v981 + (v982 + (v983 + (v984 + (v985 + (v986 + (v987 + (v988 + v956))))))))))))))))))))))))))))))
	goto L155
L154:
	;
	v1020 = v956
	goto L155
L155:
	;
	goto L152
L156:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1033 = F_add_size(m, v1028, (v1026+int32(31))&int32(-32))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1033
	v1038 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[16]))
	if v1038 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1038)+4))
	v1041 = v1039
	goto L160
L159:
	;
	v1041 = int32(0)
	goto L160
L160:
	;
	v1042 = F_mul_size(m, int32(4), v1041)
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1051 = F_add_size(m, v1046, (v1042+int32(43))&int32(-32))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1051
	v1057 = F_add_size(m, v1051, int32(1056))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1057
	v1062 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[17]))
	if v1062 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(v1062)))
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+4))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+412))
	if v1066 != 0 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	v1130 = int32(0)
	goto L166
L166:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[18]))
	if v1132 != 0 {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v1130 = v1129
	goto L166
L168:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+376))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+364))
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+352))
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+340))
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+328))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+316))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+304))
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+292))
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+280))
	v1076 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+268))
	v1077 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+256))
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+244))
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+232))
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+220))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+208))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+196))
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+184))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+172))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+160))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+148))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+136))
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+124))
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+112))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+100))
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+88))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+76))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+64))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+52))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+40))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+28))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1064)+16))
	v1129 = v1067 + (v1068 + (v1069 + (v1070 + (v1071 + (v1072 + (v1073 + (v1074 + (v1075 + (v1076 + (v1077 + (v1078 + (v1079 + (v1080 + (v1081 + (v1082 + (v1083 + (v1084 + (v1085 + (v1086 + (v1087 + (v1088 + (v1089 + (v1090 + (v1091 + (v1092 + (v1093 + (v1094 + (v1095 + (v1096 + (v1097 + v1065))))))))))))))))))))))))))))))
	goto L170
L169:
	;
	v1129 = v1065
	goto L170
L170:
	;
	goto L167
L171:
	;
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1132)))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+4))
	v1136 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+412))
	if v1136 != 0 {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1201 = v1130
	goto L173
L173:
	;
	v1203 = v1201 << (uint(int32(2)) % 32)
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1211 = F_add_size(m, v1206, (v1203+int32(39))&int32(-32))
	mBase = m.M
	v1212 = m.ExcPending
	if v1212 != 0 {
		goto L1
	} else {
		goto L178
	}
L174:
	;
	v1201 = v1199 + v1130
	goto L173
L175:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+376))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+364))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+352))
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+340))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+328))
	v1142 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+316))
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+304))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+292))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+280))
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+268))
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+256))
	v1148 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+244))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+232))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+220))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+208))
	v1152 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+196))
	v1153 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+184))
	v1154 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+172))
	v1155 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+160))
	v1156 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+148))
	v1157 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+136))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+124))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+112))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+100))
	v1161 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+88))
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+76))
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+64))
	v1164 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+52))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+40))
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+28))
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v1134)+16))
	v1199 = v1137 + (v1138 + (v1139 + (v1140 + (v1141 + (v1142 + (v1143 + (v1144 + (v1145 + (v1146 + (v1147 + (v1148 + (v1149 + (v1150 + (v1151 + (v1152 + (v1153 + (v1154 + (v1155 + (v1156 + (v1157 + (v1158 + (v1159 + (v1160 + (v1161 + (v1162 + (v1163 + (v1164 + (v1165 + (v1166 + (v1167 + v1135))))))))))))))))))))))))))))))
	goto L177
L176:
	;
	v1199 = v1135
	goto L177
L177:
	;
	goto L174
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1211
	v1216 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	if v1219 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1220 = F_strlen(m, v1219)
	mBase = m.M
	v1223 = F_add_size(m, v1216, v1220+int32(1))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	v1225 = v1216
	goto L182
L182:
	;
	v1226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1231 = F_add_size(m, v1226, (v1225+int32(31))&int32(-32))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L1
	} else {
		goto L184
	}
L183:
	;
	v1225 = v1223
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1231
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1236 = F_add_size(m, v1234, int32(12))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1236
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1242 = F_mul_size(m, int32(_a_F_InitializeParallelDSM_12), v1241)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v1248 = F_add_size(m, v1239, (v1242+int32(31))&int32(-32))
	mBase = m.M
	v1249 = m.ExcPending
	if v1249 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1248
	v1251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1253 = F_add_size(m, v1251, int32(1))
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1253
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1258 = F_strlen(m, v1257)
	mBase = m.M
	v1259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1260 = F_strlen(m, v1259)
	mBase = m.M
	v1266 = F_add_size(m, v1256, (v1258+v1260+int32(33))&int32(-32))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1266
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1271 = F_add_size(m, v1269, int32(1))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1271
	v1275 = v733
	v1277 = v69
	v1278 = v425
	v1279 = v366
	v1281 = v1225
	v1286 = v823
	v1296 = v847
	v1297 = v848
	v1298 = v938
	v1299 = v1026
	v1300 = v1042 + int32(12)
	v1301 = v1203 + int32(8)
	v1305 = int32(1048)
	goto L6
L191:
	;
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1277)))
	if v1308 <= int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if v1316 != 0 {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	v1311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1316 = v1311
	goto L192
L194:
	;
	goto L195
L195:
	;
	v1313 = F_dsm_create(m, v1306, int32(1))
	mBase = m.M
	v1314 = m.ExcPending
	if v1314 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1313
	v1316 = v1313
	goto L192
L197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1325))) = int64(1346862204)
	v1328 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1325)+8)), uint32(v1328))
	*(*int64)(unsafe.Add(mBase, uint32(v1325)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1325)+12)) = v1306 & int32(-32)
	goto L202
L198:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+24))
	v1325 = v1317
	goto L197
L199:
	;
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1321 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[7]))
	v1322 = F_MemoryContextAlloc(m, v1321, v1306)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v1322
	v1325 = v1322
	goto L197
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1325
	v1338 = F_shm_toc_allocate(m, v1325, int32(80))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338))) = v1341
	v1344 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+4)) = v1344
	v1347 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+8)) = v1347
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[23]))
	v1353 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[24])))
	if v1353 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+12)) = v1354
	v1361 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[25]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338+int32(16)))) = v1361
	v1364 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[26]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338+int32(28)))) = v1364
	goto L208
L205:
	;
	v1354 = v1350
	goto L207
L206:
	;
	v1354 = int32(0)
	goto L207
L207:
	;
	goto L204
L208:
	;
	v1367 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[27])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1338)+32)) = uint8(v1367)
	v1370 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[28])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1338)+33)) = uint8(v1370)
	v1373 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[29]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+20)) = v1373
	v1376 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+24)) = v1376
	v1379 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+36)) = v1379
	v1382 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+40)) = v1382
	v1385 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[33]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+44)) = v1385
	v1388 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[34]))
	*(*int64)(unsafe.Add(mBase, uint32(v1338)+48)) = v1388
	v1391 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[35]))
	*(*int64)(unsafe.Add(mBase, uint32(v1338)+56)) = v1391
	v1394 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v1338)+64)) = v1394
	v1396 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1338)+68)), uint32(v1396))
	*(*int64)(unsafe.Add(mBase, uint32(v1338)+72)) = int64(0)
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1401, int64(-65535), v1338)
	mBase = m.M
	v1404 = m.ExcPending
	if v1404 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v1405 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L1
	} else {
		goto L524
	}
L211:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1409 = F_shm_toc_allocate(m, v1408, v1278)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	v3450 = v1405
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3450
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v39
	return
L214:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[10]))
	if v1412 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1415 = v1412
	v1416 = v1409
	v1417 = v1278
	goto L218
L216:
	;
	v1573 = v1409
	goto L217
L217:
	;
	v1602 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1573))) = uint8(v1602)
	v1604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1604, int64(-65533), v1409)
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L1
	} else {
		goto L252
	}
L218:
	;
	v1446 = v1415 + int32(24)
	if v1417 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v1573 = v1567
	goto L217
L220:
	;
	v1566 = v1562 + (v1559 - v1416) + int32(1)
	v1567 = v1416 + v1566
	v1569 = *(*int32)(unsafe.Add(mBase, uint32(v1415)))
	if v1569 != 0 {
		v1415 = v1569
		v1416 = v1567
		v1417 = v1417 - v1566
		goto L218
	} else {
		goto L251
	}
L221:
	;
	v1562 = F_strlen(m, v1558)
	mBase = m.M
	goto L220
L222:
	;
	v1558 = v1446
	v1559 = v1416
	goto L221
L223:
	;
	goto L224
L224:
	;
	v1452 = v1417 - int32(1)
	if (v1416^v1446)&int32(3) != 0 {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	v1555 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1552))) = uint8(v1555)
	v1558 = v1551
	v1559 = v1552
	goto L221
L226:
	;
	v1536 = v1531
	v1537 = v1532
	v1538 = v1533
	goto L247
L227:
	;
	if v1526 == int32(0) {
		v1551 = v1524
		v1552 = v1525
		goto L225
	} else {
		goto L246
	}
L228:
	;
	v1524 = v1446
	v1525 = v1416
	v1526 = v1452
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1456 = int32(0)
	if base.B2i32(v1446&int32(3) == v1456)|base.B2i32(v1452 == v1456) == v1456 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	if v1492 == int32(0) {
		v1551 = v1489
		v1552 = v1490
		goto L225
	} else {
		goto L240
	}
L232:
	;
	v1468 = v1446
	v1469 = v1416
	v1470 = v1452
	goto L235
L233:
	;
	goto L234
L234:
	;
	v1489 = v1446
	v1490 = v1416
	v1491 = v1452
	v1492 = base.B2i32(v1452 != v1456)
	goto L231
L235:
	;
	v1472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1468))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1469))) = uint8(v1472)
	if v1472 == int32(0) {
		v1531 = v1468
		v1532 = v1469
		v1533 = v1470
		goto L226
	} else {
		goto L237
	}
L236:
	;
	v1489 = v1483
	v1490 = v1477
	v1491 = v1479
	v1492 = v1481
	goto L231
L237:
	;
	v1476 = int32(1)
	v1477 = v1469 + v1476
	v1479 = v1470 - v1476
	v1480 = int32(0)
	v1481 = base.B2i32(v1479 != v1480)
	v1483 = v1468 + v1476
	if v1483&int32(3) == v1480 {
		v1489 = v1483
		v1490 = v1477
		v1491 = v1479
		v1492 = v1481
		goto L231
	} else {
		goto L238
	}
L238:
	;
	if v1479 != 0 {
		v1468 = v1483
		v1469 = v1477
		v1470 = v1479
		goto L235
	} else {
		goto L239
	}
L239:
	;
	goto L236
L240:
	;
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1489))))
	if base.B2i32(v1495 == int32(0))|base.B2i32(base.Ui32(v1491) < base.Ui32(int32(4))) != 0 {
		v1524 = v1489
		v1525 = v1490
		v1526 = v1491
		goto L227
	} else {
		goto L241
	}
L241:
	;
	v1502 = v1489
	v1503 = v1490
	v1504 = v1491
	goto L242
L242:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1502)))
	v1510 = int32(-2139062144)
	if (int32(16843008)-v1507|v1507)&v1510 != v1510 {
		v1531 = v1502
		v1532 = v1503
		v1533 = v1504
		goto L226
	} else {
		goto L244
	}
L243:
	;
	v1524 = v1518
	v1525 = v1516
	v1526 = v1520
	goto L227
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1503))) = v1507
	v1515 = int32(4)
	v1516 = v1503 + v1515
	v1518 = v1502 + v1515
	v1520 = v1504 - v1515
	if base.Ui32(int32(3)) < base.Ui32(v1520) {
		v1502 = v1518
		v1503 = v1516
		v1504 = v1520
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v1531 = v1524
	v1532 = v1525
	v1533 = v1526
	goto L226
L247:
	;
	v1540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1536))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1537))) = uint8(v1540)
	if v1540 == int32(0) {
		v1551 = v1536
		v1552 = v1537
		goto L225
	} else {
		goto L249
	}
L248:
	;
	v1551 = v1547
	v1552 = v1545
	goto L225
L249:
	;
	v1544 = int32(1)
	v1545 = v1537 + v1544
	v1547 = v1536 + v1544
	v1549 = v1538 - v1544
	if v1549 != 0 {
		v1536 = v1547
		v1537 = v1545
		v1538 = v1549
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
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1609 = F_shm_toc_allocate(m, v1608, v1275)
	mBase = m.M
	v1610 = m.ExcPending
	if v1610 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v1611 = m.G0
	v1613 = v1611 - int32(112)
	m.G0 = v1613
	v1615 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+108)) = v1609 + v1615
	v1619 = v1275 - v1615
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+104)) = v1619
	v1622 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[11]))
	v1623 = int32(0)
	if base.B2i32(v1622 == v1623)|base.B2i32(v1622 == int32(_a_F_InitializeParallelDSM_8)) == v1623 {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	v2041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2041, int64(-65532), v1609)
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L307
	}
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L304
	}
L256:
	;
	v1631 = v1619
	v1634 = v1622
	goto L259
L257:
	;
	v1956 = v1619
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1609))) = v1619 - v1956
	m.G0 = v1613 + int32(112)
	goto L254
L259:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v1634-int32(60))))
	if base.Ui32(v1664) < base.Ui32(int32(2)) {
		v1921 = v1631
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1956 = v1921
	goto L258
L261:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+4))
	if v1952 != int32(_a_F_InitializeParallelDSM_8) {
		v1631 = v1921
		v1634 = v1952
		goto L259
	} else {
		goto L303
	}
L262:
	;
	v1668 = v1634 - int32(32)
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1668)))
	if v1669 == int32(0) {
		v1921 = v1631
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1673 = v1634 + int32(-64)
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v1673)))
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+96)) = v1674
	F_do_serialize(m, v1613+int32(108), v1613+int32(104), int32(_a_F_InitializeParallelDSM_13), v1613+int32(96))
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1634-int32(40))))
	switch v1687 {
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
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+20))
	if v1863 != 0 {
		goto L291
	} else {
		goto L292
	}
L266:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+28))
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v1740)))
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+36))
	if v1742 == int32(0) {
		goto L255
	} else {
		goto L281
	}
L267:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+28))
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v1726)))
	if v1727 != 0 {
		goto L277
	} else {
		goto L278
	}
L268:
	;
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+28))
	v1713 = *(*float64)(unsafe.Add(mBase, uint32(v1712)))
	*(*float64)(unsafe.Add(mBase, uint32(v1613)+40)) = v1713
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+32)) = int32(17)
	F_do_serialize(m, v1613+int32(108), v1613+int32(104), int32(_a_F_InitializeParallelDSM_14), v1613+int32(32))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L1
	} else {
		goto L276
	}
L269:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+28))
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v1700)))
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+16)) = v1701
	F_do_serialize(m, v1613+int32(108), v1613+int32(104), int32(_a_F_InitializeParallelDSM_15), v1613+int32(16))
	mBase = m.M
	v1711 = m.ExcPending
	if v1711 != 0 {
		goto L1
	} else {
		goto L275
	}
L270:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+28))
	v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1694))))
	if v1695 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1696 = int32(_a_F_InitializeParallelDSM_16)
	goto L273
L272:
	;
	v1696 = int32(_a_F_InitializeParallelDSM_17)
	goto L273
L273:
	;
	F_do_serialize(m, v1613+int32(108), v1613+int32(104), v1696, int32(0))
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
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
	v1729 = v1727
	goto L279
L278:
	;
	v1729 = int32(_a_F_InitializeParallelDSM_18)
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+48)) = v1729
	F_do_serialize(m, v1613+int32(108), v1613+int32(104), int32(_a_F_InitializeParallelDSM_13), v1613+int32(48))
	mBase = m.M
	v1739 = m.ExcPending
	if v1739 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	goto L265
L281:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1742)))
	if v1745 == int32(0) {
		goto L255
	} else {
		goto L282
	}
L282:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v1742)+4))
	if v1741 != v1748 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1751 = v1742
	goto L286
L284:
	;
	v1795 = v1745
	goto L285
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+80)) = v1795
	F_do_serialize(m, v1613+int32(108), v1613+int32(104), int32(_a_F_InitializeParallelDSM_13), v1613+int32(80))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L1
	} else {
		goto L290
	}
L286:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+12))
	if v1782 == int32(0) {
		goto L255
	} else {
		goto L288
	}
L287:
	;
	v1795 = v1782
	goto L285
L288:
	;
	v1785 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+16))
	if v1741 != v1785 {
		v1751 = v1751 + int32(12)
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
	v1865 = v1863
	goto L293
L292:
	;
	v1865 = int32(_a_F_InitializeParallelDSM_18)
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1613))) = v1865
	F_do_serialize(m, v1613+int32(108), v1613+int32(104), int32(_a_F_InitializeParallelDSM_13), v1613)
	mBase = m.M
	v1873 = m.ExcPending
	if v1873 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+20))
	if v1874 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	if base.Ui32(v1894) <= base.Ui32(int32(3)) {
		goto L210
	} else {
		goto L300
	}
L296:
	;
	v1891 = *(*int32)(unsafe.Add(mBase, uint32(v1613)+104))
	v1894 = v1891
	goto L295
L297:
	;
	v1877 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1874))))
	if v1877 == int32(0) {
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v1613)+104))
	if base.Ui32(v1880) <= base.Ui32(int32(3)) {
		goto L210
	} else {
		goto L299
	}
L299:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1613)+108))
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1883))) = v1884
	v1886 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+108)) = v1883 + v1886
	v1894 = v1880 - v1886
	goto L295
L300:
	;
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v1613)+108))
	v1898 = *(*int32)(unsafe.Add(mBase, uint32(v1668)))
	*(*int32)(unsafe.Add(mBase, uint32(v1897))) = v1898
	v1901 = v1894 & int32(-4)
	if v1901 == int32(4) {
		goto L210
	} else {
		goto L301
	}
L301:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1634-int32(24))))
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+4)) = v1906
	if v1901 == int32(8) {
		goto L210
	} else {
		goto L302
	}
L302:
	;
	v1912 = *(*int32)(unsafe.Add(mBase, uint32(v1634-int32(16))))
	*(*int32)(unsafe.Add(mBase, uint32(v1897)+8)) = v1912
	v1914 = int32(12)
	v1915 = v1894 - v1914
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+104)) = v1915
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+108)) = v1897 + v1914
	v1921 = v1915
	goto L261
L303:
	;
	goto L260
L304:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v1673)))
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+68)) = v2028
	*(*int32)(unsafe.Add(mBase, uint32(v1613)+64)) = v1741
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_9), v1613-int32(-64))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(3036), int32(_a_F_InitializeParallelDSM_11))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
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
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2046 = F_shm_toc_allocate(m, v2045, v1286)
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v2046))) = v2049
	v2052 = v2049 << (uint(int32(3)) % 32)
	if base.Ui32(v2052|int32(4)) <= base.Ui32(v1286) {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2079, int64(-65531), v2046)
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L1
	} else {
		goto L317
	}
L310:
	;
	v2056 = int32(0)
	if base.B2i32(v2052 == v2056)|base.B2i32(v2049 <= v2056) != 0 {
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
	v2069 = m.ExcPending
	if v2069 != 0 {
		goto L1
	} else {
		goto L314
	}
L313:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[37]))
	base.MemoryCopy(m, v2046+int32(4), v2064, v2052)
	goto L309
L314:
	;
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_19), int32(0))
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_20), int32(327), int32(_a_F_InitializeParallelDSM_21))
	mBase = m.M
	v2078 = m.ExcPending
	if v2078 != 0 {
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
	v2084 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[13]))
	if int32(2) <= v2084 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v2087 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2088 = F_shm_toc_allocate(m, v2087, v1296)
	mBase = m.M
	v2089 = m.ExcPending
	if v2089 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2147 = F_shm_toc_allocate(m, v2146, v1297)
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L1
	} else {
		goto L336
	}
L321:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v2097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+29)))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v2099 = *(*int64)(unsafe.Add(mBase, uint32(v33)+4))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v2101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2088)+16)) = uint8(v2101)
	*(*int32)(unsafe.Add(mBase, uint32(v2088)+20)) = v2100
	*(*int64)(unsafe.Add(mBase, uint32(v2088))) = v2099
	*(*int32)(unsafe.Add(mBase, uint32(v2088)+8)) = v2098
	*(*uint8)(unsafe.Add(mBase, uint32(v2088)+17)) = uint8(v2097)
	if v2097&int32(1) != 0 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v2141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2141, int64(-65530), v2088)
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L1
	} else {
		goto L335
	}
L323:
	;
	v2110 = v2096
	goto L325
L324:
	;
	v2110 = int32(0)
	goto L325
L325:
	;
	if v2101 != 0 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2111 = v2110
	goto L328
L327:
	;
	v2111 = v2096
	goto L328
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2088)+12)) = v2111
	v2113 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v2113 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	if v2111 <= int32(0) {
		goto L332
	} else {
		goto L333
	}
L330:
	;
	v2117 = v2113 << (uint(int32(2)) % 32)
	if v2117 == int32(0) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v2122 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	base.MemoryCopy(m, v2088+int32(24), v2122, v2117)
	goto L329
L332:
	;
	goto L322
L333:
	;
	v2127 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v2129 = v2127 << (uint(int32(2)) % 32)
	if v2129 == int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	base.MemoryCopy(m, v2088+v2132<<(uint(int32(2))%32)+int32(24), v2138, v2129)
	goto L332
L335:
	;
	goto L320
L336:
	;
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v2156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+29)))
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v2158 = *(*int64)(unsafe.Add(mBase, uint32(v37)+4))
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v2160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2147)+16)) = uint8(v2160)
	*(*int32)(unsafe.Add(mBase, uint32(v2147)+20)) = v2159
	*(*int64)(unsafe.Add(mBase, uint32(v2147))) = v2158
	*(*int32)(unsafe.Add(mBase, uint32(v2147)+8)) = v2157
	*(*uint8)(unsafe.Add(mBase, uint32(v2147)+17)) = uint8(v2156)
	if v2156&int32(1) != 0 {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v2200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2200, int64(-65529), v2147)
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L1
	} else {
		goto L350
	}
L338:
	;
	v2169 = v2155
	goto L340
L339:
	;
	v2169 = int32(0)
	goto L340
L340:
	;
	if v2160 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v2170 = v2169
	goto L343
L342:
	;
	v2170 = v2155
	goto L343
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2147)+12)) = v2170
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v2172 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	if v2170 <= int32(0) {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v2176 = v2172 << (uint(int32(2)) % 32)
	if v2176 == int32(0) {
		goto L344
	} else {
		goto L346
	}
L346:
	;
	v2181 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	base.MemoryCopy(m, v2147+int32(24), v2181, v2176)
	goto L344
L347:
	;
	goto L337
L348:
	;
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v2188 = v2186 << (uint(int32(2)) % 32)
	if v2188 == int32(0) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v2191 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v2197 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	base.MemoryCopy(m, v2147+v2191<<(uint(int32(2))%32)+int32(24), v2197, v2188)
	goto L347
L350:
	;
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2206 = F_shm_toc_allocate(m, v2204, int32(4))
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2206))) = v1279
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2209, int64(-65526), v2206)
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2214 = F_shm_toc_allocate(m, v2213, v1298)
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v2216 = int32(0)
	v2218 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v2214))) = v2218
	v2221 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[38])))
	*(*uint8)(unsafe.Add(mBase, uint32(v2214)+4)) = uint8(v2221)
	v2224 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[39]))
	*(*int64)(unsafe.Add(mBase, uint32(v2214)+8)) = v2224
	v2227 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[14]))
	v2228 = *(*int64)(unsafe.Add(mBase, uint32(v2227)))
	*(*int64)(unsafe.Add(mBase, uint32(v2214)+16)) = v2228
	v2231 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v2214)+24)) = v2231
	v2234 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[41]))
	if v2216 < v2234 {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2424, int64(-65528), v2214)
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L1
	} else {
		goto L385
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2214)+28)) = v2234
	v2239 = v2234 << (uint(int32(2)) % 32)
	if v2239 == int32(0) {
		goto L354
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v2251 = v2227
	v2253 = v2216
	goto L359
L358:
	;
	v2245 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[42]))
	base.MemoryCopy(m, v2214+int32(32), v2245, v2239)
	goto L354
L359:
	;
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v2251)))
	if v2279 != 0 {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	v2290 = v2285 << (uint(int32(2)) % 32)
	v2291 = F_palloc(m, v2290)
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L1
	} else {
		goto L367
	}
L361:
	;
	v2281 = F_add_size(m, v2253, int32(1))
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L1
	} else {
		goto L364
	}
L362:
	;
	v2283 = v2253
	goto L363
L363:
	;
	v2284 = *(*int32)(unsafe.Add(mBase, uint32(v2251)+52))
	v2285 = F_add_size(m, v2283, v2284)
	mBase = m.M
	v2286 = m.ExcPending
	if v2286 != 0 {
		goto L1
	} else {
		goto L365
	}
L364:
	;
	v2283 = v2281
	goto L363
L365:
	;
	v2287 = *(*int32)(unsafe.Add(mBase, uint32(v2251)+80))
	if v2287 != 0 {
		v2251 = v2287
		v2253 = v2285
		goto L359
	} else {
		goto L366
	}
L366:
	;
	goto L360
L367:
	;
	v2294 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[14]))
	if v2294 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v2298 = int32(0)
	v2299 = v2294
	goto L371
L369:
	;
	goto L370
L370:
	;
	F_pg_qsort(m, v2291, v2285, int32(4), int32(185))
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L1
	} else {
		goto L383
	}
L371:
	;
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2299)))
	if v2327 != 0 {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	goto L370
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2291+v2298<<(uint(int32(2))%32)))) = v2327
	v2334 = v2298 + int32(1)
	goto L375
L374:
	;
	v2334 = v2298
	goto L375
L375:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+52))
	if int32(0) < v2335 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v2339 = v2335 << (uint(int32(2)) % 32)
	if v2339 != 0 {
		goto L379
	} else {
		goto L380
	}
L377:
	;
	v2347 = v2335
	goto L378
L378:
	;
	v2349 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+80))
	if v2349 != 0 {
		v2298 = v2347 + v2334
		v2299 = v2349
		goto L371
	} else {
		goto L382
	}
L379:
	;
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+48))
	base.MemoryCopy(m, v2291+v2334<<(uint(int32(2))%32), v2343, v2339)
	goto L381
L380:
	;
	goto L381
L381:
	;
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2299)+52))
	v2347 = v2345
	goto L378
L382:
	;
	goto L372
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2214)+28)) = v2285
	if v2290 == int32(0) {
		goto L354
	} else {
		goto L384
	}
L384:
	;
	base.MemoryCopy(m, v2214+int32(32), v2291, v2290)
	goto L354
L385:
	;
	v2428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2429 = F_shm_toc_allocate(m, v2428, v1299)
	mBase = m.M
	v2430 = m.ExcPending
	if v2430 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v2431 = m.G0
	v2433 = v2431 - int32(80)
	m.G0 = v2433
	v2436 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[15]))
	if v2436 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2433)+48)) = int64(51539607564)
	v2440 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v2433)+72)) = v2440
	v2444 = *(*int32)(unsafe.Add(mBase, uint32(v2436)))
	v2445 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+4))
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+412))
	if v2446 != 0 {
		goto L391
	} else {
		goto L392
	}
L388:
	;
	v2754 = v2429
	goto L389
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2754)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2754))) = int64(0)
	m.G0 = v2433 + int32(80)
	v2790 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2790, int64(-65525), v2429)
	mBase = m.M
	v2793 = m.ExcPending
	if v2793 != 0 {
		goto L1
	} else {
		goto L425
	}
L390:
	;
	v2513 = F_hash_create(m, int32(_a_F_InitializeParallelDSM_22), v2509, v2433+int32(32), int32(1064))
	mBase = m.M
	v2514 = m.ExcPending
	if v2514 != 0 {
		goto L1
	} else {
		goto L394
	}
L391:
	;
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+376))
	v2448 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+364))
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+352))
	v2450 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+340))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+328))
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+316))
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+304))
	v2454 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+292))
	v2455 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+280))
	v2456 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+268))
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+256))
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+244))
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+232))
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+220))
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+208))
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+196))
	v2463 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+184))
	v2464 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+172))
	v2465 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+160))
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+148))
	v2467 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+136))
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+124))
	v2469 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+112))
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+100))
	v2471 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+88))
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+76))
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+64))
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+52))
	v2475 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+40))
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+28))
	v2477 = *(*int32)(unsafe.Add(mBase, uint32(v2444)+16))
	v2509 = v2447 + (v2448 + (v2449 + (v2450 + (v2451 + (v2452 + (v2453 + (v2454 + (v2455 + (v2456 + (v2457 + (v2458 + (v2459 + (v2460 + (v2461 + (v2462 + (v2463 + (v2464 + (v2465 + (v2466 + (v2467 + (v2468 + (v2469 + (v2470 + (v2471 + (v2472 + (v2473 + (v2474 + (v2475 + (v2476 + (v2477 + v2445))))))))))))))))))))))))))))))
	goto L393
L392:
	;
	v2509 = v2445
	goto L393
L393:
	;
	goto L390
L394:
	;
	v2516 = v2433 + int32(12)
	v2518 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[15]))
	F_hash_seq_init(m, v2516, v2518)
	mBase = m.M
	v2520 = m.ExcPending
	if v2520 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	v2521 = F_hash_seq_search(m, v2516)
	mBase = m.M
	v2522 = m.ExcPending
	if v2522 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	if v2521 != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2527 = v2521
	goto L400
L398:
	;
	goto L399
L399:
	;
	v2596 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[43]))
	if v2596 != 0 {
		goto L405
	} else {
		goto L406
	}
L400:
	;
	v2557 = F_hash_search(m, v2513, v2527, int32(1), int32(0))
	mBase = m.M
	v2558 = m.ExcPending
	if v2558 != 0 {
		goto L1
	} else {
		goto L402
	}
L401:
	;
	goto L399
L402:
	;
	v2561 = F_hash_seq_search(m, v2433+int32(12))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	if v2561 != 0 {
		v2527 = v2561
		goto L400
	} else {
		goto L404
	}
L404:
	;
	goto L401
L405:
	;
	v2601 = v2596
	goto L408
L406:
	;
	goto L407
L407:
	;
	v2670 = v2433 + int32(12)
	F_hash_seq_init(m, v2670, v2513)
	mBase = m.M
	v2672 = m.ExcPending
	if v2672 != 0 {
		goto L1
	} else {
		goto L415
	}
L408:
	;
	v2629 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2601)+16)))
	if v2629 == int32(1) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L407
L410:
	;
	v2634 = F_hash_search(m, v2513, v2601, int32(2), int32(0))
	mBase = m.M
	v2635 = m.ExcPending
	if v2635 != 0 {
		goto L1
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v2636 = *(*int32)(unsafe.Add(mBase, uint32(v2601)+24))
	if v2636 != 0 {
		v2601 = v2636
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
	v2673 = F_hash_seq_search(m, v2670)
	mBase = m.M
	v2674 = m.ExcPending
	if v2674 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	if v2673 != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v2678 = v2429
	v2679 = v2673
	goto L420
L418:
	;
	v2720 = v2429
	goto L419
L419:
	;
	F_hash_destroy(m, v2513)
	mBase = m.M
	v2750 = m.ExcPending
	if v2750 != 0 {
		goto L1
	} else {
		goto L424
	}
L420:
	;
	v2707 = *(*int32)(unsafe.Add(mBase, uint32(v2679)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2678)+8)) = v2707
	v2709 = *(*int64)(unsafe.Add(mBase, uint32(v2679)))
	*(*int64)(unsafe.Add(mBase, uint32(v2678))) = v2709
	v2711 = int32(12)
	v2712 = v2678 + v2711
	v2715 = F_hash_seq_search(m, v2433+v2711)
	mBase = m.M
	v2716 = m.ExcPending
	if v2716 != 0 {
		goto L1
	} else {
		goto L422
	}
L421:
	;
	v2720 = v2712
	goto L419
L422:
	;
	if v2715 != 0 {
		v2678 = v2712
		v2679 = v2715
		goto L420
	} else {
		goto L423
	}
L423:
	;
	goto L421
L424:
	;
	v2754 = v2720
	goto L389
L425:
	;
	v2794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2795 = F_shm_toc_allocate(m, v2794, v1300)
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v2799 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v2795))) = v2799
	v2802 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v2795)+4)) = v2802
	v2805 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[16]))
	if v2805 != 0 {
		goto L428
	} else {
		goto L429
	}
L427:
	;
	v2890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2890, int64(-65524), v2795)
	mBase = m.M
	v2893 = m.ExcPending
	if v2893 != 0 {
		goto L1
	} else {
		goto L435
	}
L428:
	;
	v2806 = *(*int32)(unsafe.Add(mBase, uint32(v2805)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2795)+8)) = v2806
	v2808 = *(*int32)(unsafe.Add(mBase, uint32(v2805)+4))
	if v2808 <= int32(0) {
		goto L427
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2795)+8)) = int32(0)
	goto L427
L431:
	;
	v2816 = int32(0)
	goto L432
L432:
	;
	v2846 = v2816 << (uint(int32(2)) % 32)
	v2848 = *(*int32)(unsafe.Add(mBase, uint32(v2805)+12))
	v2850 = *(*int32)(unsafe.Add(mBase, uint32(v2848+v2846)))
	*(*int32)(unsafe.Add(mBase, uint32(v2795+int32(12)+v2846))) = v2850
	v2853 = v2816 + int32(1)
	v2854 = *(*int32)(unsafe.Add(mBase, uint32(v2805)+4))
	if v2853 < v2854 {
		v2816 = v2853
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
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2895 = F_shm_toc_allocate(m, v2894, v1305)
	mBase = m.M
	v2896 = m.ExcPending
	if v2896 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	v2898 = int32(524)
	base.MemoryCopy(m, v2895, int32(_a_F_InitializeParallelDSM_23), v2898)
	base.MemoryCopy(m, v2895+v2898, int32(_a_F_InitializeParallelDSM_24), v2898)
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2905, int64(-65523), v2895)
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2910 = F_shm_toc_allocate(m, v2909, v1301)
	mBase = m.M
	v2911 = m.ExcPending
	if v2911 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	v2912 = m.G0
	v2914 = v2912 - int32(32)
	m.G0 = v2914
	v2917 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[17]))
	if v2917 == int32(0) {
		v2972 = v2910
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v3000 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2972))) = v3000
	v3003 = v2972 + int32(4)
	v3005 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[18]))
	if v3005 == v3000 {
		v3060 = v3003
		goto L448
	} else {
		goto L449
	}
L440:
	;
	v2921 = v2914 + int32(12)
	F_hash_seq_init(m, v2921, v2917)
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	v2924 = F_hash_seq_search(m, v2921)
	mBase = m.M
	v2925 = m.ExcPending
	if v2925 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	if v2924 == int32(0) {
		v2972 = v2910
		goto L439
	} else {
		goto L443
	}
L443:
	;
	v2929 = v2924
	v2932 = v2910
	goto L444
L444:
	;
	v2960 = *(*int32)(unsafe.Add(mBase, uint32(v2929)))
	*(*int32)(unsafe.Add(mBase, uint32(v2932))) = v2960
	v2963 = v2932 + int32(4)
	v2966 = F_hash_seq_search(m, v2914+int32(12))
	mBase = m.M
	v2967 = m.ExcPending
	if v2967 != 0 {
		goto L1
	} else {
		goto L446
	}
L445:
	;
	v2972 = v2963
	goto L439
L446:
	;
	if v2966 != 0 {
		v2929 = v2966
		v2932 = v2963
		goto L444
	} else {
		goto L447
	}
L447:
	;
	goto L445
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3060))) = int32(0)
	m.G0 = v2914 + int32(32)
	v3093 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3093, int64(-65522), v2910)
	mBase = m.M
	v3096 = m.ExcPending
	if v3096 != 0 {
		goto L1
	} else {
		goto L457
	}
L449:
	;
	v3009 = v2914 + int32(12)
	F_hash_seq_init(m, v3009, v3005)
	mBase = m.M
	v3011 = m.ExcPending
	if v3011 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	v3012 = F_hash_seq_search(m, v3009)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	if v3012 == int32(0) {
		v3060 = v3003
		goto L448
	} else {
		goto L452
	}
L452:
	;
	v3017 = v3012
	v3020 = v3003
	goto L453
L453:
	;
	v3048 = *(*int32)(unsafe.Add(mBase, uint32(v3017)))
	*(*int32)(unsafe.Add(mBase, uint32(v3020))) = v3048
	v3051 = v3020 + int32(4)
	v3054 = F_hash_seq_search(m, v2914+int32(12))
	mBase = m.M
	v3055 = m.ExcPending
	if v3055 != 0 {
		goto L1
	} else {
		goto L455
	}
L454:
	;
	v3060 = v3051
	goto L448
L455:
	;
	if v3054 != 0 {
		v3017 = v3054
		v3020 = v3051
		goto L453
	} else {
		goto L456
	}
L456:
	;
	goto L454
L457:
	;
	v3097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3098 = F_shm_toc_allocate(m, v3097, v1281)
	mBase = m.M
	v3099 = m.ExcPending
	if v3099 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v3101 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[46]))
	v3103 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	if v3103 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	v3125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3125, int64(-65521), v3098)
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		goto L1
	} else {
		goto L466
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3098)+4)) = v3101
	*(*int32)(unsafe.Add(mBase, uint32(v3098))) = int32(-1)
	goto L459
L461:
	;
	goto L462
L462:
	;
	v3109 = F_strlen(m, v3103)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3098)+4)) = v3101
	*(*int32)(unsafe.Add(mBase, uint32(v3098))) = v3109
	if v3109 < int32(0) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	goto L459
L464:
	;
	v3115 = v3109 + int32(1)
	if v3115 == int32(0) {
		goto L463
	} else {
		goto L465
	}
L465:
	;
	v3121 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	base.MemoryCopy(m, v3098+int32(8), v3121, v3115)
	goto L463
L466:
	;
	v3129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3132 = F_palloc0(m, v3129<<(uint(int32(3))%32))
	mBase = m.M
	v3133 = m.ExcPending
	if v3133 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v3132
	v3135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3138 = F_mul_size(m, int32(_a_F_InitializeParallelDSM_12), v3137)
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v3140 = F_shm_toc_allocate(m, v3135, v3138)
	mBase = m.M
	v3141 = m.ExcPending
	if v3141 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v3142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v3142 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v3147 = int32(0)
	goto L473
L471:
	;
	goto L472
L472:
	;
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3246, int64(-65534), v3140)
	mBase = m.M
	v3249 = m.ExcPending
	if v3249 != 0 {
		goto L1
	} else {
		goto L479
	}
L473:
	;
	v3180 = v3140 + v3147<<(uint(int32(14))%32)
	v3182 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3180))), uint32(v3182))
	v3185 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3180)+16)) = v3185
	*(*int64)(unsafe.Add(mBase, uint32(v3180)+4)) = v3185
	v3189 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v3180)+36)) = uint16(v3189)
	*(*int64)(unsafe.Add(mBase, uint32(v3180)+24)) = v3185
	*(*int32)(unsafe.Add(mBase, uint32(v3180)+32)) = int32(_a_F_InitializeParallelDSM_25)
	goto L475
L474:
	;
	goto L472
L475:
	;
	v3199 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[31]))
	F_shm_mq_set_receiver(m, v3180, v3199)
	mBase = m.M
	v3201 = m.ExcPending
	if v3201 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	v3202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v3203 = F_shm_mq_attach(m, v3180, v3202)
	mBase = m.M
	v3204 = m.ExcPending
	if v3204 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	v3205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3205+v3147<<(uint(int32(3))%32))+4)) = v3203
	v3211 = v3147 + int32(1)
	v3212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3211 < v3212 {
		v3147 = v3211
		goto L473
	} else {
		goto L478
	}
L478:
	;
	goto L474
L479:
	;
	v3250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3251 = F_strlen(m, v3250)
	mBase = m.M
	v3252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3253 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3254 = F_strlen(m, v3253)
	mBase = m.M
	v3258 = F_shm_toc_allocate(m, v3252, v3254+v3251+int32(2))
	mBase = m.M
	v3259 = m.ExcPending
	if v3259 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if (v3260^v3258)&int32(3) != 0 {
		goto L484
	} else {
		goto L485
	}
L481:
	;
	v3337 = v3251 + v3258 + int32(1)
	v3338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if (v3338^v3337)&int32(3) != 0 {
		goto L505
	} else {
		goto L506
	}
L482:
	;
	goto L481
L483:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3315))) = uint8(v3314)
	if v3314&int32(255) == int32(0) {
		goto L482
	} else {
		goto L498
	}
L484:
	;
	v3266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3260))))
	v3313 = v3260
	v3314 = v3266
	v3315 = v3258
	goto L483
L485:
	;
	goto L486
L486:
	;
	if v3260&int32(3) != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v3270 = v3260
	v3272 = v3258
	goto L490
L488:
	;
	v3284 = v3260
	v3286 = v3258
	goto L489
L489:
	;
	v3288 = *(*int32)(unsafe.Add(mBase, uint32(v3284)))
	v3291 = int32(-2139062144)
	if (int32(16843008)-v3288|v3288)&v3291 != v3291 {
		v3313 = v3284
		v3314 = v3288
		v3315 = v3286
		goto L483
	} else {
		goto L494
	}
L490:
	;
	v3273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3270))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3272))) = uint8(v3273)
	if v3273 == int32(0) {
		goto L482
	} else {
		goto L492
	}
L491:
	;
	v3284 = v3280
	v3286 = v3278
	goto L489
L492:
	;
	v3277 = int32(1)
	v3278 = v3272 + v3277
	v3280 = v3270 + v3277
	if v3280&int32(3) != 0 {
		v3270 = v3280
		v3272 = v3278
		goto L490
	} else {
		goto L493
	}
L493:
	;
	goto L491
L494:
	;
	v3296 = v3284
	v3297 = v3288
	v3298 = v3286
	goto L495
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3298))) = v3297
	v3300 = int32(4)
	v3301 = v3298 + v3300
	v3303 = v3296 + v3300
	v3305 = *(*int32)(unsafe.Add(mBase, uint32(v3296)+4))
	v3308 = int32(-2139062144)
	if (int32(16843008)-v3305|v3305)&v3308 == v3308 {
		v3296 = v3303
		v3297 = v3305
		v3298 = v3301
		goto L495
	} else {
		goto L497
	}
L496:
	;
	v3313 = v3303
	v3314 = v3305
	v3315 = v3301
	goto L483
L497:
	;
	goto L496
L498:
	;
	v3322 = v3313
	v3324 = v3315
	goto L499
L499:
	;
	v3325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3322)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3324)+1)) = uint8(v3325)
	v3327 = int32(1)
	if v3325 != 0 {
		v3322 = v3322 + v3327
		v3324 = v3324 + v3327
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
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3413, int64(-65527), v3258)
	mBase = m.M
	v3416 = m.ExcPending
	if v3416 != 0 {
		goto L1
	} else {
		goto L523
	}
L503:
	;
	goto L502
L504:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3393))) = uint8(v3392)
	if v3392&int32(255) == int32(0) {
		goto L503
	} else {
		goto L519
	}
L505:
	;
	v3344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3338))))
	v3391 = v3338
	v3392 = v3344
	v3393 = v3337
	goto L504
L506:
	;
	goto L507
L507:
	;
	if v3338&int32(3) != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3348 = v3338
	v3350 = v3337
	goto L511
L509:
	;
	v3362 = v3338
	v3364 = v3337
	goto L510
L510:
	;
	v3366 = *(*int32)(unsafe.Add(mBase, uint32(v3362)))
	v3369 = int32(-2139062144)
	if (int32(16843008)-v3366|v3366)&v3369 != v3369 {
		v3391 = v3362
		v3392 = v3366
		v3393 = v3364
		goto L504
	} else {
		goto L515
	}
L511:
	;
	v3351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3348))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3350))) = uint8(v3351)
	if v3351 == int32(0) {
		goto L503
	} else {
		goto L513
	}
L512:
	;
	v3362 = v3358
	v3364 = v3356
	goto L510
L513:
	;
	v3355 = int32(1)
	v3356 = v3350 + v3355
	v3358 = v3348 + v3355
	if v3358&int32(3) != 0 {
		v3348 = v3358
		v3350 = v3356
		goto L511
	} else {
		goto L514
	}
L514:
	;
	goto L512
L515:
	;
	v3374 = v3362
	v3375 = v3366
	v3376 = v3364
	goto L516
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3376))) = v3375
	v3378 = int32(4)
	v3379 = v3376 + v3378
	v3381 = v3374 + v3378
	v3383 = *(*int32)(unsafe.Add(mBase, uint32(v3374)+4))
	v3386 = int32(-2139062144)
	if (int32(16843008)-v3383|v3383)&v3386 == v3386 {
		v3374 = v3381
		v3375 = v3383
		v3376 = v3379
		goto L516
	} else {
		goto L518
	}
L517:
	;
	v3391 = v3381
	v3392 = v3383
	v3393 = v3379
	goto L504
L518:
	;
	goto L517
L519:
	;
	v3400 = v3391
	v3402 = v3393
	goto L520
L520:
	;
	v3403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3400)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3402)+1)) = uint8(v3403)
	v3405 = int32(1)
	if v3403 != 0 {
		v3400 = v3400 + v3405
		v3402 = v3402 + v3405
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
	v3417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3450 = v3417
	goto L213
L524:
	;
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_26), int32(0))
	mBase = m.M
	v3464 = m.ExcPending
	if v3464 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(_a_F_InitializeParallelDSM_27), int32(_a_F_InitializeParallelDSM_28))
	mBase = m.M
	v3469 = m.ExcPending
	if v3469 != 0 {
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
