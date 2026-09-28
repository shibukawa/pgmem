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
	var v134 int32
	_ = v134
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
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
	v134 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v133
	if base.Ui32(int32(2)) <= base.Ui32(v133) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v147 = int32(32) - base.I32_clz(v133-int32(1))
	goto L22
L21:
	;
	v147 = v134
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149)+24)) = uint8(v150)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v154 = v152 + int32(40)
	v156 = F_LWLockAcquire(m, v154, v150)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+24))
	if v158 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v162 = v158
	v163 = v152 + int32(24)
	v164 = v154
	goto L27
L25:
	;
	v275 = v154
	goto L26
L26:
	;
	F_LWLockRelease(m, v275)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L49
	}
L27:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v173 = F_dsa_get_address(m, v172, v162)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L5
	} else {
		goto L29
	}
L28:
	;
	v275 = v265
	goto L26
L29:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v175
	F_LWLockRelease(m, v164)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if v179 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v180 = int32(16)
	v186 = int32(0)
	goto L34
L32:
	;
	goto L33
L33:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_ExecParallelHashIncreaseNumBuckets[0]))
	if v260 != 0 {
		goto L43
	} else {
		goto L44
	}
L34:
	;
	v196 = v186 + (v173 + v180)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v205 = v197 + v198&(v199-int32(1))<<(uint(int32(2))%32)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v206
	v208 = v186 + (v162 + v180)
	v210 = base.AtomicRmwCmpxchg32(m, v205, int32(0), v206, v208)
	if v206 != v210 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v216 = v210
	goto L39
L37:
	;
	goto L38
L38:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	v245 = (v240+int32(15))&int32(-8) + v186
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	if base.Ui32(v245) < base.Ui32(v246) {
		v186 = v245
		goto L34
	} else {
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v216
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v224
	v227 = base.AtomicRmwCmpxchg32(m, v205, int32(0), v224, v208)
	if v224 != v227 {
		v216 = v227
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L38
L41:
	;
	goto L40
L42:
	;
	goto L35
L43:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L5
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v265 = v263 + int32(40)
	v267 = F_LWLockAcquire(m, v265, int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v263)+24))
	if v271 != 0 {
		v162 = v271
		v163 = v263 + int32(24)
		v164 = v265
		goto L27
	} else {
		goto L48
	}
L48:
	;
	goto L28
L49:
	;
	v286 = F_BarrierArriveAndWait(m, v14, int32(134217756))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	if l0 == int32(0) {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v12 - int32(403) {
		case 0:
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+36)))
			if v59 != int32(1) {
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = int32(0)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
				if v65 != 0 {
					base.MemoryFill(m, v62+int32(20), int32(0), v65)
				} else {
				}
			}
			v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
			mBase = m.M
			v204 = m.ExcPending
			if v204 != 0 {
				return int32(0)
			} else {
				return v203
			}
		default:
			v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
			mBase = m.M
			v204 = m.ExcPending
			if v204 != 0 {
				return int32(0)
			} else {
				return v203
			}
		case 6:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+36)))
			if v16 != int32(1) {
				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					return v203
				}
			} else {
				F_ExecSeqScanReInitializeDSM(m, l0)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						return v203
					}
				}
			}
		case 8:
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+36)))
			if v24 != int32(1) {
				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					return v203
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
				F_index_parallelrescan(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						return v203
					}
				}
			}
		case 9:
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+36)))
			if v31 != int32(1) {
				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					return v203
				}
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
				F_index_parallelrescan(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						return v203
					}
				}
			}
		case 11:
			v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+36)))
			if v86 != int32(1) {
				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					return v203
				}
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+172))
				if v90 != 0 {
					v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
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
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
								if v110 == int32(0) {
									v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
									if v123 == int32(0) {
										F_dsa_free(m, v90, v94)
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
											v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int32(0)
											} else {
												return v203
											}
										}
									} else {
										v126 = F_dsa_get_address(m, v90, v123)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											v128 = int32(1)
											v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
											if v130 != v128 {
												F_dsa_free(m, v90, v94)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
													v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int32(0)
													} else {
														return v203
													}
												}
											} else {
												v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
												F_dsa_free(m, v90, v133)
												mBase = m.M
												v135 = m.ExcPending
												if v135 != 0 {
													return int32(0)
												} else {
													F_dsa_free(m, v90, v94)
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
														v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int32(0)
														} else {
															return v203
														}
													}
												}
											}
										}
									}
								} else {
									v113 = F_dsa_get_address(m, v90, v110)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										v115 = int32(1)
										v117 = base.AtomicRmwSub32(m, v113, int32(0), v115)
										if v117 != v115 {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
											if v123 == int32(0) {
												F_dsa_free(m, v90, v94)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
													v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int32(0)
													} else {
														return v203
													}
												}
											} else {
												v126 = F_dsa_get_address(m, v90, v123)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													v128 = int32(1)
													v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
													if v130 != v128 {
														F_dsa_free(m, v90, v94)
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
															v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int32(0)
															} else {
																return v203
															}
														}
													} else {
														v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
														F_dsa_free(m, v90, v133)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int32(0)
														} else {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int32(0)
																} else {
																	return v203
																}
															}
														}
													}
												}
											}
										} else {
											v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
											F_dsa_free(m, v90, v120)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int32(0)
											} else {
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
												if v123 == int32(0) {
													F_dsa_free(m, v90, v94)
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
														v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int32(0)
														} else {
															return v203
														}
													}
												} else {
													v126 = F_dsa_get_address(m, v90, v123)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = int32(1)
														v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
														if v130 != v128 {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int32(0)
																} else {
																	return v203
																}
															}
														} else {
															v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
															F_dsa_free(m, v90, v133)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return int32(0)
															} else {
																F_dsa_free(m, v90, v94)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																	v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int32(0)
																	} else {
																		return v203
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
									v102 = int32(1)
									v104 = base.AtomicRmwSub32(m, v100, int32(0), v102)
									if v104 != v102 {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
										if v110 == int32(0) {
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
											if v123 == int32(0) {
												F_dsa_free(m, v90, v94)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
													v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int32(0)
													} else {
														return v203
													}
												}
											} else {
												v126 = F_dsa_get_address(m, v90, v123)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													v128 = int32(1)
													v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
													if v130 != v128 {
														F_dsa_free(m, v90, v94)
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
															v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int32(0)
															} else {
																return v203
															}
														}
													} else {
														v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
														F_dsa_free(m, v90, v133)
														mBase = m.M
														v135 = m.ExcPending
														if v135 != 0 {
															return int32(0)
														} else {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int32(0)
																} else {
																	return v203
																}
															}
														}
													}
												}
											}
										} else {
											v113 = F_dsa_get_address(m, v90, v110)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												v115 = int32(1)
												v117 = base.AtomicRmwSub32(m, v113, int32(0), v115)
												if v117 != v115 {
													v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
													if v123 == int32(0) {
														F_dsa_free(m, v90, v94)
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
															v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int32(0)
															} else {
																return v203
															}
														}
													} else {
														v126 = F_dsa_get_address(m, v90, v123)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int32(0)
														} else {
															v128 = int32(1)
															v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
															if v130 != v128 {
																F_dsa_free(m, v90, v94)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																	v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int32(0)
																	} else {
																		return v203
																	}
																}
															} else {
																v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
																F_dsa_free(m, v90, v133)
																mBase = m.M
																v135 = m.ExcPending
																if v135 != 0 {
																	return int32(0)
																} else {
																	F_dsa_free(m, v90, v94)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																		v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																		mBase = m.M
																		v204 = m.ExcPending
																		if v204 != 0 {
																			return int32(0)
																		} else {
																			return v203
																		}
																	}
																}
															}
														}
													}
												} else {
													v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
													F_dsa_free(m, v90, v120)
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int32(0)
													} else {
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
														if v123 == int32(0) {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int32(0)
																} else {
																	return v203
																}
															}
														} else {
															v126 = F_dsa_get_address(m, v90, v123)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																v128 = int32(1)
																v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
																if v130 != v128 {
																	F_dsa_free(m, v90, v94)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																		v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																		mBase = m.M
																		v204 = m.ExcPending
																		if v204 != 0 {
																			return int32(0)
																		} else {
																			return v203
																		}
																	}
																} else {
																	v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
																	F_dsa_free(m, v90, v133)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return int32(0)
																	} else {
																		F_dsa_free(m, v90, v94)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																			v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																			mBase = m.M
																			v204 = m.ExcPending
																			if v204 != 0 {
																				return int32(0)
																			} else {
																				return v203
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
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)+16))
										F_dsa_free(m, v90, v107)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
											if v110 == int32(0) {
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
												if v123 == int32(0) {
													F_dsa_free(m, v90, v94)
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
														v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int32(0)
														} else {
															return v203
														}
													}
												} else {
													v126 = F_dsa_get_address(m, v90, v123)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = int32(1)
														v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
														if v130 != v128 {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int32(0)
																} else {
																	return v203
																}
															}
														} else {
															v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
															F_dsa_free(m, v90, v133)
															mBase = m.M
															v135 = m.ExcPending
															if v135 != 0 {
																return int32(0)
															} else {
																F_dsa_free(m, v90, v94)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																	v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int32(0)
																	} else {
																		return v203
																	}
																}
															}
														}
													}
												}
											} else {
												v113 = F_dsa_get_address(m, v90, v110)
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return int32(0)
												} else {
													v115 = int32(1)
													v117 = base.AtomicRmwSub32(m, v113, int32(0), v115)
													if v117 != v115 {
														v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
														if v123 == int32(0) {
															F_dsa_free(m, v90, v94)
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int32(0)
																} else {
																	return v203
																}
															}
														} else {
															v126 = F_dsa_get_address(m, v90, v123)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																v128 = int32(1)
																v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
																if v130 != v128 {
																	F_dsa_free(m, v90, v94)
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																		v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																		mBase = m.M
																		v204 = m.ExcPending
																		if v204 != 0 {
																			return int32(0)
																		} else {
																			return v203
																		}
																	}
																} else {
																	v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
																	F_dsa_free(m, v90, v133)
																	mBase = m.M
																	v135 = m.ExcPending
																	if v135 != 0 {
																		return int32(0)
																	} else {
																		F_dsa_free(m, v90, v94)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																			v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																			mBase = m.M
																			v204 = m.ExcPending
																			if v204 != 0 {
																				return int32(0)
																			} else {
																				return v203
																			}
																		}
																	}
																}
															}
														}
													} else {
														v120 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
														F_dsa_free(m, v90, v120)
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															v123 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
															if v123 == int32(0) {
																F_dsa_free(m, v90, v94)
																mBase = m.M
																v137 = m.ExcPending
																if v137 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																	v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																	mBase = m.M
																	v204 = m.ExcPending
																	if v204 != 0 {
																		return int32(0)
																	} else {
																		return v203
																	}
																}
															} else {
																v126 = F_dsa_get_address(m, v90, v123)
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return int32(0)
																} else {
																	v128 = int32(1)
																	v130 = base.AtomicRmwSub32(m, v126, int32(0), v128)
																	if v130 != v128 {
																		F_dsa_free(m, v90, v94)
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																			v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																			mBase = m.M
																			v204 = m.ExcPending
																			if v204 != 0 {
																				return int32(0)
																			} else {
																				return v203
																			}
																		}
																	} else {
																		v133 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
																		F_dsa_free(m, v90, v133)
																		mBase = m.M
																		v135 = m.ExcPending
																		if v135 != 0 {
																			return int32(0)
																		} else {
																			F_dsa_free(m, v90, v94)
																			mBase = m.M
																			v137 = m.ExcPending
																			if v137 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(0)
																				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
																				mBase = m.M
																				v204 = m.ExcPending
																				if v204 != 0 {
																					return int32(0)
																				} else {
																					return v203
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
						v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
						mBase = m.M
						v204 = m.ExcPending
						if v204 != 0 {
							return int32(0)
						} else {
							return v203
						}
					}
				} else {
					v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						return v203
					}
				}
			}
		case 13:
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+36)))
			if v53 != int32(1) {
				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					return v203
				}
			} else {
				F_ExecSeqScanReInitializeDSM(m, l0)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						return v203
					}
				}
			}
		case 21:
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+36)))
			if v38 != int32(1) {
				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					return v203
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+156))
				if v42 != 0 {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v44)+40)))
					v47 = F_shm_toc_lookup(m, v43, v45, int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+156))
						m.T0[v49].(func(*base.Module, int32, int32, int32))(m, l0, l1, v47)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
							mBase = m.M
							v204 = m.ExcPending
							if v204 != 0 {
								return int32(0)
							} else {
								return v203
							}
						}
					}
				} else {
					v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						return v203
					}
				}
			}
		case 22:
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+36)))
			if v71 != int32(1) {
				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					return v203
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
							v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
							mBase = m.M
							v204 = m.ExcPending
							if v204 != 0 {
								return int32(0)
							} else {
								return v203
							}
						}
					}
				} else {
					v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						return v203
					}
				}
			}
		case 26:
			v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+36)))
			if v147 != int32(1) {
				v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
				mBase = m.M
				v204 = m.ExcPending
				if v204 != 0 {
					return int32(0)
				} else {
					return v203
				}
			} else {
				v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				if v150 != 0 {
					v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v153 = int64(*(*int32)(unsafe.Add(mBase, uint32(v152)+40)))
					v155 = F_shm_toc_lookup(m, v151, v153, int32(0))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int32(0)
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
						if v157 != 0 {
							F_ExecHashTableDetachBatch(m, v157)
							mBase = m.M
							v159 = m.ExcPending
							if v159 != 0 {
								return int32(0)
							} else {
								v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
								F_ExecHashTableDetach(m, v160)
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return int32(0)
								} else {
									F_FileSetDeleteAll(m, v155+int32(168))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
										return int32(0)
									} else {
										v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
										if v167 != 0 {
											F_tuplestore_end(m, v167)
											mBase = m.M
											v169 = m.ExcPending
											if v169 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
												v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+120))
												if v173 != 0 {
													F_tuplestore_end(m, v173)
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v172)+120)) = int32(0)
														v179 = v155 + int32(56)
														v180 = int32(0)
														atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v179))), uint32(v180))
														*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180
														*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
														*(*int64)(unsafe.Add(mBase, uint32(v179)+12)) = int64(0)
														*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v180
														F_ConditionVariableInit(m, v155+int32(80))
														mBase = m.M
														v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int32(0)
														} else {
															return v203
														}
													}
												} else {
													v179 = v155 + int32(56)
													v180 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v179))), uint32(v180))
													*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180
													*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
													*(*int64)(unsafe.Add(mBase, uint32(v179)+12)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v180
													F_ConditionVariableInit(m, v155+int32(80))
													mBase = m.M
													v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int32(0)
													} else {
														return v203
													}
												}
											}
										} else {
											v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+120))
											if v173 != 0 {
												F_tuplestore_end(m, v173)
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v172)+120)) = int32(0)
													v179 = v155 + int32(56)
													v180 = int32(0)
													atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v179))), uint32(v180))
													*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180
													*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
													*(*int64)(unsafe.Add(mBase, uint32(v179)+12)) = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v180
													F_ConditionVariableInit(m, v155+int32(80))
													mBase = m.M
													v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int32(0)
													} else {
														return v203
													}
												}
											} else {
												v179 = v155 + int32(56)
												v180 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v179))), uint32(v180))
												*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180
												*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
												*(*int64)(unsafe.Add(mBase, uint32(v179)+12)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v180
												F_ConditionVariableInit(m, v155+int32(80))
												mBase = m.M
												v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int32(0)
												} else {
													return v203
												}
											}
										}
									}
								}
							}
						} else {
							F_FileSetDeleteAll(m, v155+int32(168))
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
								return int32(0)
							} else {
								v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
								if v167 != 0 {
									F_tuplestore_end(m, v167)
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
										v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+120))
										if v173 != 0 {
											F_tuplestore_end(m, v173)
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v172)+120)) = int32(0)
												v179 = v155 + int32(56)
												v180 = int32(0)
												atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v179))), uint32(v180))
												*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180
												*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
												*(*int64)(unsafe.Add(mBase, uint32(v179)+12)) = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v180
												F_ConditionVariableInit(m, v155+int32(80))
												mBase = m.M
												v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int32(0)
												} else {
													return v203
												}
											}
										} else {
											v179 = v155 + int32(56)
											v180 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v179))), uint32(v180))
											*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180
											*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
											*(*int64)(unsafe.Add(mBase, uint32(v179)+12)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v180
											F_ConditionVariableInit(m, v155+int32(80))
											mBase = m.M
											v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int32(0)
											} else {
												return v203
											}
										}
									}
								} else {
									v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+120))
									if v173 != 0 {
										F_tuplestore_end(m, v173)
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v172)+120)) = int32(0)
											v179 = v155 + int32(56)
											v180 = int32(0)
											atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v179))), uint32(v180))
											*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180
											*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
											*(*int64)(unsafe.Add(mBase, uint32(v179)+12)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v180
											F_ConditionVariableInit(m, v155+int32(80))
											mBase = m.M
											v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int32(0)
											} else {
												return v203
											}
										}
									} else {
										v179 = v155 + int32(56)
										v180 = int32(0)
										atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v179))), uint32(v180))
										*(*int32)(unsafe.Add(mBase, uint32(v179)+8)) = v180
										*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
										*(*int64)(unsafe.Add(mBase, uint32(v179)+12)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v180
										F_ConditionVariableInit(m, v155+int32(80))
										mBase = m.M
										v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return int32(0)
										} else {
											return v203
										}
									}
								}
							}
						}
					}
				} else {
					v203 = F_planstate_tree_walker_impl(m, l0, int32(674), l1)
					mBase = m.M
					v204 = m.ExcPending
					if v204 != 0 {
						return int32(0)
					} else {
						return v203
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
	var v956 int64
	_ = v956
	var v957 int64
	_ = v957
	var v960 int64
	_ = v960
	var v961 int64
	_ = v961
	var v962 int64
	_ = v962
	var v963 int64
	_ = v963
	var v964 int64
	_ = v964
	var v965 int64
	_ = v965
	var v966 int64
	_ = v966
	var v967 int64
	_ = v967
	var v968 int64
	_ = v968
	var v969 int64
	_ = v969
	var v970 int64
	_ = v970
	var v971 int64
	_ = v971
	var v972 int64
	_ = v972
	var v973 int64
	_ = v973
	var v974 int64
	_ = v974
	var v975 int64
	_ = v975
	var v976 int64
	_ = v976
	var v977 int64
	_ = v977
	var v978 int64
	_ = v978
	var v979 int64
	_ = v979
	var v980 int64
	_ = v980
	var v981 int64
	_ = v981
	var v982 int64
	_ = v982
	var v983 int64
	_ = v983
	var v984 int64
	_ = v984
	var v985 int64
	_ = v985
	var v986 int64
	_ = v986
	var v987 int64
	_ = v987
	var v988 int64
	_ = v988
	var v989 int64
	_ = v989
	var v990 int64
	_ = v990
	var v1022 int64
	_ = v1022
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1068 int64
	_ = v1068
	var v1069 int64
	_ = v1069
	var v1072 int64
	_ = v1072
	var v1073 int64
	_ = v1073
	var v1074 int64
	_ = v1074
	var v1075 int64
	_ = v1075
	var v1076 int64
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1078 int64
	_ = v1078
	var v1079 int64
	_ = v1079
	var v1080 int64
	_ = v1080
	var v1081 int64
	_ = v1081
	var v1082 int64
	_ = v1082
	var v1083 int64
	_ = v1083
	var v1084 int64
	_ = v1084
	var v1085 int64
	_ = v1085
	var v1086 int64
	_ = v1086
	var v1087 int64
	_ = v1087
	var v1088 int64
	_ = v1088
	var v1089 int64
	_ = v1089
	var v1090 int64
	_ = v1090
	var v1091 int64
	_ = v1091
	var v1092 int64
	_ = v1092
	var v1093 int64
	_ = v1093
	var v1094 int64
	_ = v1094
	var v1095 int64
	_ = v1095
	var v1096 int64
	_ = v1096
	var v1097 int64
	_ = v1097
	var v1098 int64
	_ = v1098
	var v1099 int64
	_ = v1099
	var v1100 int64
	_ = v1100
	var v1101 int64
	_ = v1101
	var v1102 int64
	_ = v1102
	var v1134 int64
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int64
	_ = v1141
	var v1142 int64
	_ = v1142
	var v1145 int64
	_ = v1145
	var v1146 int64
	_ = v1146
	var v1147 int64
	_ = v1147
	var v1148 int64
	_ = v1148
	var v1149 int64
	_ = v1149
	var v1150 int64
	_ = v1150
	var v1151 int64
	_ = v1151
	var v1152 int64
	_ = v1152
	var v1153 int64
	_ = v1153
	var v1154 int64
	_ = v1154
	var v1155 int64
	_ = v1155
	var v1156 int64
	_ = v1156
	var v1157 int64
	_ = v1157
	var v1158 int64
	_ = v1158
	var v1159 int64
	_ = v1159
	var v1160 int64
	_ = v1160
	var v1161 int64
	_ = v1161
	var v1162 int64
	_ = v1162
	var v1163 int64
	_ = v1163
	var v1164 int64
	_ = v1164
	var v1165 int64
	_ = v1165
	var v1166 int64
	_ = v1166
	var v1167 int64
	_ = v1167
	var v1168 int64
	_ = v1168
	var v1169 int64
	_ = v1169
	var v1170 int64
	_ = v1170
	var v1171 int64
	_ = v1171
	var v1172 int64
	_ = v1172
	var v1173 int64
	_ = v1173
	var v1174 int64
	_ = v1174
	var v1175 int64
	_ = v1175
	var v1207 int64
	_ = v1207
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1226 int32
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
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
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1320 int32
	_ = v1320
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
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
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1394 int32
	_ = v1394
	var v1397 int64
	_ = v1397
	var v1400 int64
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1410 int32
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1455 int32
	_ = v1455
	var v1461 int32
	_ = v1461
	var v1465 int32
	_ = v1465
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1492 int32
	_ = v1492
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1519 int32
	_ = v1519
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
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
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1582 int32
	_ = v1582
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1618 int32
	_ = v1618
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1628 int32
	_ = v1628
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1640 int32
	_ = v1640
	var v1643 int32
	_ = v1643
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1693 int32
	_ = v1693
	var v1696 int32
	_ = v1696
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1722 float64
	_ = v1722
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1738 int32
	_ = v1738
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1757 int32
	_ = v1757
	var v1760 int32
	_ = v1760
	var v1791 int32
	_ = v1791
	var v1794 int32
	_ = v1794
	var v1804 int32
	_ = v1804
	var v1839 int32
	_ = v1839
	var v1872 int32
	_ = v1872
	var v1874 int32
	_ = v1874
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1900 int32
	_ = v1900
	var v1903 int32
	_ = v1903
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1910 int32
	_ = v1910
	var v1915 int32
	_ = v1915
	var v1921 int32
	_ = v1921
	var v1923 int32
	_ = v1923
	var v1924 int32
	_ = v1924
	var v1930 int32
	_ = v1930
	var v1961 int32
	_ = v1961
	var v1965 int32
	_ = v1965
	var v2036 int32
	_ = v2036
	var v2037 int32
	_ = v2037
	var v2044 int32
	_ = v2044
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2058 int32
	_ = v2058
	var v2061 int32
	_ = v2061
	var v2065 int32
	_ = v2065
	var v2073 int32
	_ = v2073
	var v2078 int32
	_ = v2078
	var v2082 int32
	_ = v2082
	var v2087 int32
	_ = v2087
	var v2088 int32
	_ = v2088
	var v2091 int32
	_ = v2091
	var v2093 int32
	_ = v2093
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int64
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2128 int32
	_ = v2128
	var v2133 int32
	_ = v2133
	var v2138 int32
	_ = v2138
	var v2140 int32
	_ = v2140
	var v2143 int32
	_ = v2143
	var v2149 int32
	_ = v2149
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2166 int32
	_ = v2166
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2169 int64
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2185 int32
	_ = v2185
	var v2189 int32
	_ = v2189
	var v2194 int32
	_ = v2194
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2204 int32
	_ = v2204
	var v2210 int32
	_ = v2210
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2219 int32
	_ = v2219
	var v2220 int32
	_ = v2220
	var v2222 int32
	_ = v2222
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2228 int32
	_ = v2228
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2234 int32
	_ = v2234
	var v2237 int64
	_ = v2237
	var v2240 int32
	_ = v2240
	var v2241 int64
	_ = v2241
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2252 int32
	_ = v2252
	var v2258 int32
	_ = v2258
	var v2264 int32
	_ = v2264
	var v2266 int32
	_ = v2266
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2298 int32
	_ = v2298
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2303 int32
	_ = v2303
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2311 int32
	_ = v2311
	var v2312 int32
	_ = v2312
	var v2340 int32
	_ = v2340
	var v2347 int32
	_ = v2347
	var v2348 int32
	_ = v2348
	var v2352 int32
	_ = v2352
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2398 int32
	_ = v2398
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	var v2441 int32
	_ = v2441
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2453 int32
	_ = v2453
	var v2457 int32
	_ = v2457
	var v2458 int64
	_ = v2458
	var v2459 int64
	_ = v2459
	var v2462 int64
	_ = v2462
	var v2463 int64
	_ = v2463
	var v2464 int64
	_ = v2464
	var v2465 int64
	_ = v2465
	var v2466 int64
	_ = v2466
	var v2467 int64
	_ = v2467
	var v2468 int64
	_ = v2468
	var v2469 int64
	_ = v2469
	var v2470 int64
	_ = v2470
	var v2471 int64
	_ = v2471
	var v2472 int64
	_ = v2472
	var v2473 int64
	_ = v2473
	var v2474 int64
	_ = v2474
	var v2475 int64
	_ = v2475
	var v2476 int64
	_ = v2476
	var v2477 int64
	_ = v2477
	var v2478 int64
	_ = v2478
	var v2479 int64
	_ = v2479
	var v2480 int64
	_ = v2480
	var v2481 int64
	_ = v2481
	var v2482 int64
	_ = v2482
	var v2483 int64
	_ = v2483
	var v2484 int64
	_ = v2484
	var v2485 int64
	_ = v2485
	var v2486 int64
	_ = v2486
	var v2487 int64
	_ = v2487
	var v2488 int64
	_ = v2488
	var v2489 int64
	_ = v2489
	var v2490 int64
	_ = v2490
	var v2491 int64
	_ = v2491
	var v2492 int64
	_ = v2492
	var v2524 int64
	_ = v2524
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2535 int32
	_ = v2535
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2542 int32
	_ = v2542
	var v2572 int32
	_ = v2572
	var v2573 int32
	_ = v2573
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2611 int32
	_ = v2611
	var v2616 int32
	_ = v2616
	var v2644 int32
	_ = v2644
	var v2649 int32
	_ = v2649
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2689 int32
	_ = v2689
	var v2693 int32
	_ = v2693
	var v2694 int32
	_ = v2694
	var v2722 int32
	_ = v2722
	var v2724 int64
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2730 int32
	_ = v2730
	var v2731 int32
	_ = v2731
	var v2735 int32
	_ = v2735
	var v2765 int32
	_ = v2765
	var v2769 int32
	_ = v2769
	var v2805 int32
	_ = v2805
	var v2808 int32
	_ = v2808
	var v2809 int32
	_ = v2809
	var v2810 int32
	_ = v2810
	var v2811 int32
	_ = v2811
	var v2814 int32
	_ = v2814
	var v2817 int32
	_ = v2817
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2831 int32
	_ = v2831
	var v2861 int32
	_ = v2861
	var v2863 int32
	_ = v2863
	var v2865 int32
	_ = v2865
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
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
	var v2913 int32
	_ = v2913
	var v2920 int32
	_ = v2920
	var v2923 int32
	_ = v2923
	var v2924 int32
	_ = v2924
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2929 int32
	_ = v2929
	var v2932 int32
	_ = v2932
	var v2936 int32
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2939 int32
	_ = v2939
	var v2940 int32
	_ = v2940
	var v2944 int32
	_ = v2944
	var v2947 int32
	_ = v2947
	var v2975 int32
	_ = v2975
	var v2978 int32
	_ = v2978
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2987 int32
	_ = v2987
	var v3015 int32
	_ = v3015
	var v3018 int32
	_ = v3018
	var v3020 int32
	_ = v3020
	var v3024 int32
	_ = v3024
	var v3026 int32
	_ = v3026
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3032 int32
	_ = v3032
	var v3035 int32
	_ = v3035
	var v3063 int32
	_ = v3063
	var v3066 int32
	_ = v3066
	var v3069 int32
	_ = v3069
	var v3070 int32
	_ = v3070
	var v3075 int32
	_ = v3075
	var v3108 int32
	_ = v3108
	var v3111 int32
	_ = v3111
	var v3112 int32
	_ = v3112
	var v3113 int32
	_ = v3113
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3118 int32
	_ = v3118
	var v3124 int32
	_ = v3124
	var v3130 int32
	_ = v3130
	var v3136 int32
	_ = v3136
	var v3140 int32
	_ = v3140
	var v3143 int32
	_ = v3143
	var v3145 int32
	_ = v3145
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3149 int32
	_ = v3149
	var v3151 int32
	_ = v3151
	var v3152 int32
	_ = v3152
	var v3153 int32
	_ = v3153
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3161 int32
	_ = v3161
	var v3194 int32
	_ = v3194
	var v3196 int32
	_ = v3196
	var v3199 int64
	_ = v3199
	var v3203 int32
	_ = v3203
	var v3213 int32
	_ = v3213
	var v3215 int32
	_ = v3215
	var v3216 int32
	_ = v3216
	var v3217 int32
	_ = v3217
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3225 int32
	_ = v3225
	var v3226 int32
	_ = v3226
	var v3260 int32
	_ = v3260
	var v3263 int32
	_ = v3263
	var v3264 int32
	_ = v3264
	var v3265 int32
	_ = v3265
	var v3266 int32
	_ = v3266
	var v3267 int32
	_ = v3267
	var v3268 int32
	_ = v3268
	var v3272 int32
	_ = v3272
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3280 int32
	_ = v3280
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3287 int32
	_ = v3287
	var v3291 int32
	_ = v3291
	var v3292 int32
	_ = v3292
	var v3294 int32
	_ = v3294
	var v3298 int32
	_ = v3298
	var v3300 int32
	_ = v3300
	var v3302 int32
	_ = v3302
	var v3305 int32
	_ = v3305
	var v3310 int32
	_ = v3310
	var v3311 int32
	_ = v3311
	var v3312 int32
	_ = v3312
	var v3314 int32
	_ = v3314
	var v3315 int32
	_ = v3315
	var v3317 int32
	_ = v3317
	var v3319 int32
	_ = v3319
	var v3322 int32
	_ = v3322
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3329 int32
	_ = v3329
	var v3336 int32
	_ = v3336
	var v3338 int32
	_ = v3338
	var v3339 int32
	_ = v3339
	var v3341 int32
	_ = v3341
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3358 int32
	_ = v3358
	var v3362 int32
	_ = v3362
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3369 int32
	_ = v3369
	var v3370 int32
	_ = v3370
	var v3372 int32
	_ = v3372
	var v3376 int32
	_ = v3376
	var v3378 int32
	_ = v3378
	var v3380 int32
	_ = v3380
	var v3383 int32
	_ = v3383
	var v3388 int32
	_ = v3388
	var v3389 int32
	_ = v3389
	var v3390 int32
	_ = v3390
	var v3392 int32
	_ = v3392
	var v3393 int32
	_ = v3393
	var v3395 int32
	_ = v3395
	var v3397 int32
	_ = v3397
	var v3400 int32
	_ = v3400
	var v3405 int32
	_ = v3405
	var v3406 int32
	_ = v3406
	var v3407 int32
	_ = v3407
	var v3414 int32
	_ = v3414
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3419 int32
	_ = v3419
	var v3427 int32
	_ = v3427
	var v3430 int32
	_ = v3430
	var v3431 int32
	_ = v3431
	var v3464 int32
	_ = v3464
	var v3474 int32
	_ = v3474
	var v3478 int32
	_ = v3478
	var v3483 int32
	_ = v3483
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
	v1315 = F_shm_toc_estimate(m, v55)
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L1
	} else {
		goto L191
	}
L7:
	;
	v69 = l0 + int32(12)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v70 <= int32(0) {
		v1284 = v2
		v1286 = v69
		v1287 = v2
		v1288 = v2
		v1290 = v2
		v1295 = v2
		v1305 = v2
		v1306 = v2
		v1307 = v2
		v1308 = v2
		v1309 = v2
		v1310 = v2
		v1314 = v2
		goto L6
	} else {
		goto L12
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1284 = v2
	v1286 = l0 + int32(12)
	v1287 = v2
	v1288 = v2
	v1290 = v2
	v1295 = v2
	v1305 = v2
	v1306 = v2
	v1307 = v2
	v1308 = v2
	v1309 = v2
	v1310 = v2
	v1314 = v2
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
	v131 = F_dsa_create_in_place_ext(m, v127, int32(_a_F_InitializeParallelDSM_1), int32(76), v108)
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
	F_on_dsm_detach(m, v108, int32(1821), int64(0))
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
	v213 = F_dsa_allocate_extended(m, v131, v207*int32(108)+int32(28), int32(0))
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
	v224 = F_dshash_find_or_insert_extended(m, v154, v206+int32(8), v223)
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
	v239 = F_dshash_find_or_insert_extended(m, v150, v141+int32(8), v223)
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
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_5), int32(2280), int32(_a_F_InitializeParallelDSM_6))
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
	v1284 = v2
	v1286 = v69
	v1287 = v2
	v1288 = v372
	v1290 = v2
	v1295 = v2
	v1305 = v2
	v1306 = v2
	v1307 = v2
	v1308 = v2
	v1309 = v2
	v1310 = v2
	v1314 = v2
	goto L6
L65:
	;
	goto L66
L66:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v375 <= int32(0) {
		v1284 = v2
		v1286 = v69
		v1287 = v2
		v1288 = v366
		v1290 = v2
		v1295 = v2
		v1305 = v2
		v1306 = v2
		v1307 = v2
		v1308 = v2
		v1309 = v2
		v1310 = v2
		v1314 = v2
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
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v477+int32(-64))))
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
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v477-int32(36))))
	if v515 == int32(0) {
		v691 = v507
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v519 = v477 - int32(68)
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v519)))
	v521 = F_strlen(m, v520)
	mBase = m.M
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v477-int32(44))))
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
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(2938), int32(_a_F_InitializeParallelDSM_11))
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
	v956 = *(*int64)(unsafe.Add(mBase, uint32(v955)+8))
	v957 = *(*int64)(unsafe.Add(mBase, uint32(v955)+808))
	if v957 != int64(0) {
		goto L153
	} else {
		goto L154
	}
L150:
	;
	v1027 = int32(1)
	goto L151
L151:
	;
	v1029 = F_mul_size(m, v1027, int32(12))
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L1
	} else {
		goto L156
	}
L152:
	;
	v1027 = base.I32_wrap_i64(v1022) + int32(1)
	goto L151
L153:
	;
	v960 = *(*int64)(unsafe.Add(mBase, uint32(v955)+752))
	v961 = *(*int64)(unsafe.Add(mBase, uint32(v955)+728))
	v962 = *(*int64)(unsafe.Add(mBase, uint32(v955)+704))
	v963 = *(*int64)(unsafe.Add(mBase, uint32(v955)+680))
	v964 = *(*int64)(unsafe.Add(mBase, uint32(v955)+656))
	v965 = *(*int64)(unsafe.Add(mBase, uint32(v955)+632))
	v966 = *(*int64)(unsafe.Add(mBase, uint32(v955)+608))
	v967 = *(*int64)(unsafe.Add(mBase, uint32(v955)+584))
	v968 = *(*int64)(unsafe.Add(mBase, uint32(v955)+560))
	v969 = *(*int64)(unsafe.Add(mBase, uint32(v955)+536))
	v970 = *(*int64)(unsafe.Add(mBase, uint32(v955)+512))
	v971 = *(*int64)(unsafe.Add(mBase, uint32(v955)+488))
	v972 = *(*int64)(unsafe.Add(mBase, uint32(v955)+464))
	v973 = *(*int64)(unsafe.Add(mBase, uint32(v955)+440))
	v974 = *(*int64)(unsafe.Add(mBase, uint32(v955)+416))
	v975 = *(*int64)(unsafe.Add(mBase, uint32(v955)+392))
	v976 = *(*int64)(unsafe.Add(mBase, uint32(v955)+368))
	v977 = *(*int64)(unsafe.Add(mBase, uint32(v955)+344))
	v978 = *(*int64)(unsafe.Add(mBase, uint32(v955)+320))
	v979 = *(*int64)(unsafe.Add(mBase, uint32(v955)+296))
	v980 = *(*int64)(unsafe.Add(mBase, uint32(v955)+272))
	v981 = *(*int64)(unsafe.Add(mBase, uint32(v955)+248))
	v982 = *(*int64)(unsafe.Add(mBase, uint32(v955)+224))
	v983 = *(*int64)(unsafe.Add(mBase, uint32(v955)+200))
	v984 = *(*int64)(unsafe.Add(mBase, uint32(v955)+176))
	v985 = *(*int64)(unsafe.Add(mBase, uint32(v955)+152))
	v986 = *(*int64)(unsafe.Add(mBase, uint32(v955)+128))
	v987 = *(*int64)(unsafe.Add(mBase, uint32(v955)+104))
	v988 = *(*int64)(unsafe.Add(mBase, uint32(v955)+80))
	v989 = *(*int64)(unsafe.Add(mBase, uint32(v955)+56))
	v990 = *(*int64)(unsafe.Add(mBase, uint32(v955)+32))
	v1022 = v960 + (v961 + (v962 + (v963 + (v964 + (v965 + (v966 + (v967 + (v968 + (v969 + (v970 + (v971 + (v972 + (v973 + (v974 + (v975 + (v976 + (v977 + (v978 + (v979 + (v980 + (v981 + (v982 + (v983 + (v984 + (v985 + (v986 + (v987 + (v988 + (v989 + (v990 + v956))))))))))))))))))))))))))))))
	goto L155
L154:
	;
	v1022 = v956
	goto L155
L155:
	;
	goto L152
L156:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1036 = F_add_size(m, v1031, (v1029+int32(31))&int32(-32))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1036
	v1041 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[16]))
	if v1041 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1041)+4))
	v1044 = v1042
	goto L160
L159:
	;
	v1044 = int32(0)
	goto L160
L160:
	;
	v1045 = F_mul_size(m, int32(4), v1044)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1054 = F_add_size(m, v1049, (v1045+int32(43))&int32(-32))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1054
	v1060 = F_add_size(m, v1054, int32(1056))
	mBase = m.M
	v1061 = m.ExcPending
	if v1061 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1060
	v1065 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[17]))
	if v1065 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v1065)))
	v1068 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+8))
	v1069 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+808))
	if v1069 != int64(0) {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	v1136 = int32(0)
	goto L166
L166:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[18]))
	if v1138 != 0 {
		goto L171
	} else {
		goto L172
	}
L167:
	;
	v1136 = base.I32_wrap_i64(v1134)
	goto L166
L168:
	;
	v1072 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+752))
	v1073 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+728))
	v1074 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+704))
	v1075 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+680))
	v1076 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+656))
	v1077 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+632))
	v1078 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+608))
	v1079 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+584))
	v1080 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+560))
	v1081 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+536))
	v1082 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+512))
	v1083 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+488))
	v1084 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+464))
	v1085 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+440))
	v1086 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+416))
	v1087 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+392))
	v1088 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+368))
	v1089 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+344))
	v1090 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+320))
	v1091 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+296))
	v1092 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+272))
	v1093 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+248))
	v1094 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+224))
	v1095 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+200))
	v1096 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+176))
	v1097 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+152))
	v1098 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+128))
	v1099 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+104))
	v1100 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+80))
	v1101 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+56))
	v1102 = *(*int64)(unsafe.Add(mBase, uint32(v1067)+32))
	v1134 = v1072 + (v1073 + (v1074 + (v1075 + (v1076 + (v1077 + (v1078 + (v1079 + (v1080 + (v1081 + (v1082 + (v1083 + (v1084 + (v1085 + (v1086 + (v1087 + (v1088 + (v1089 + (v1090 + (v1091 + (v1092 + (v1093 + (v1094 + (v1095 + (v1096 + (v1097 + (v1098 + (v1099 + (v1100 + (v1101 + (v1102 + v1068))))))))))))))))))))))))))))))
	goto L170
L169:
	;
	v1134 = v1068
	goto L170
L170:
	;
	goto L167
L171:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1138)))
	v1141 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+8))
	v1142 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+808))
	if v1142 != int64(0) {
		goto L175
	} else {
		goto L176
	}
L172:
	;
	v1210 = v1136
	goto L173
L173:
	;
	v1212 = v1210 << (uint(int32(2)) % 32)
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1220 = F_add_size(m, v1215, (v1212+int32(39))&int32(-32))
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L1
	} else {
		goto L178
	}
L174:
	;
	v1210 = v1136 + base.I32_wrap_i64(v1207)
	goto L173
L175:
	;
	v1145 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+752))
	v1146 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+728))
	v1147 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+704))
	v1148 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+680))
	v1149 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+656))
	v1150 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+632))
	v1151 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+608))
	v1152 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+584))
	v1153 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+560))
	v1154 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+536))
	v1155 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+512))
	v1156 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+488))
	v1157 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+464))
	v1158 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+440))
	v1159 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+416))
	v1160 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+392))
	v1161 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+368))
	v1162 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+344))
	v1163 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+320))
	v1164 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+296))
	v1165 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+272))
	v1166 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+248))
	v1167 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+224))
	v1168 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+200))
	v1169 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+176))
	v1170 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+152))
	v1171 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+128))
	v1172 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+104))
	v1173 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+80))
	v1174 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+56))
	v1175 = *(*int64)(unsafe.Add(mBase, uint32(v1140)+32))
	v1207 = v1145 + (v1146 + (v1147 + (v1148 + (v1149 + (v1150 + (v1151 + (v1152 + (v1153 + (v1154 + (v1155 + (v1156 + (v1157 + (v1158 + (v1159 + (v1160 + (v1161 + (v1162 + (v1163 + (v1164 + (v1165 + (v1166 + (v1167 + (v1168 + (v1169 + (v1170 + (v1171 + (v1172 + (v1173 + (v1174 + (v1175 + v1141))))))))))))))))))))))))))))))
	goto L177
L176:
	;
	v1207 = v1141
	goto L177
L177:
	;
	goto L174
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1220
	v1225 = F_add_size(m, int32(0), int32(8))
	mBase = m.M
	v1226 = m.ExcPending
	if v1226 != 0 {
		goto L1
	} else {
		goto L179
	}
L179:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	if v1228 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v1229 = F_strlen(m, v1228)
	mBase = m.M
	v1232 = F_add_size(m, v1225, v1229+int32(1))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L183
	}
L181:
	;
	v1234 = v1225
	goto L182
L182:
	;
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1240 = F_add_size(m, v1235, (v1234+int32(31))&int32(-32))
	mBase = m.M
	v1241 = m.ExcPending
	if v1241 != 0 {
		goto L1
	} else {
		goto L184
	}
L183:
	;
	v1234 = v1232
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1240
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1245 = F_add_size(m, v1243, int32(12))
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1245
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1251 = F_mul_size(m, int32(_a_F_InitializeParallelDSM_12), v1250)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	v1257 = F_add_size(m, v1248, (v1251+int32(31))&int32(-32))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1257
	v1260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1262 = F_add_size(m, v1260, int32(1))
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1262
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v1266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1267 = F_strlen(m, v1266)
	mBase = m.M
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1269 = F_strlen(m, v1268)
	mBase = m.M
	v1275 = F_add_size(m, v1265, (v1267+v1269+int32(33))&int32(-32))
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v1275
	v1278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v1280 = F_add_size(m, v1278, int32(1))
	mBase = m.M
	v1281 = m.ExcPending
	if v1281 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v1280
	v1284 = v733
	v1286 = v69
	v1287 = v425
	v1288 = v366
	v1290 = v1234
	v1295 = v823
	v1305 = v847
	v1306 = v848
	v1307 = v938
	v1308 = v1029
	v1309 = v1045 + int32(12)
	v1310 = v1212 + int32(8)
	v1314 = int32(1048)
	goto L6
L191:
	;
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v1286)))
	if v1317 <= int32(0) {
		goto L193
	} else {
		goto L194
	}
L192:
	;
	if v1325 != 0 {
		goto L198
	} else {
		goto L199
	}
L193:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v1325 = v1320
	goto L192
L194:
	;
	goto L195
L195:
	;
	v1322 = F_dsm_create(m, v1315, int32(1))
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v1322
	v1325 = v1322
	goto L192
L197:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1334))) = int64(1346862204)
	v1337 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1334)+8)), uint32(v1337))
	*(*int64)(unsafe.Add(mBase, uint32(v1334)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1334)+12)) = v1315 & int32(-32)
	goto L202
L198:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1325)+24))
	v1334 = v1326
	goto L197
L199:
	;
	goto L200
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	v1330 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[7]))
	v1331 = F_MemoryContextAlloc(m, v1330, v1315)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v1331
	v1334 = v1331
	goto L197
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1334
	v1347 = F_shm_toc_allocate(m, v1334, int32(80))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	v1350 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[20]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347))) = v1350
	v1353 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[21]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+4)) = v1353
	v1356 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[22]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+8)) = v1356
	v1359 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[23]))
	v1362 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[24])))
	if v1362 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+12)) = v1363
	v1370 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[25]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347+int32(16)))) = v1370
	v1373 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[26]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347+int32(28)))) = v1373
	goto L208
L205:
	;
	v1363 = v1359
	goto L207
L206:
	;
	v1363 = int32(0)
	goto L207
L207:
	;
	goto L204
L208:
	;
	v1376 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[27])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1347)+32)) = uint8(v1376)
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[28])))
	*(*uint8)(unsafe.Add(mBase, uint32(v1347)+33)) = uint8(v1379)
	v1382 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[29]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+20)) = v1382
	v1385 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[30]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+24)) = v1385
	v1388 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[31]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+36)) = v1388
	v1391 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[32]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+40)) = v1391
	v1394 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[33]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+44)) = v1394
	v1397 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[34]))
	*(*int64)(unsafe.Add(mBase, uint32(v1347)+48)) = v1397
	v1400 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[35]))
	*(*int64)(unsafe.Add(mBase, uint32(v1347)+56)) = v1400
	v1403 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[36]))
	*(*int32)(unsafe.Add(mBase, uint32(v1347)+64)) = v1403
	v1405 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v1347)+68)), uint32(v1405))
	*(*int64)(unsafe.Add(mBase, uint32(v1347)+72)) = int64(0)
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1410, int64(-65535), v1347)
	mBase = m.M
	v1413 = m.ExcPending
	if v1413 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v1414 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3474 = m.ExcPending
	if v3474 != 0 {
		goto L1
	} else {
		goto L524
	}
L211:
	;
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1418 = F_shm_toc_allocate(m, v1417, v1287)
	mBase = m.M
	v1419 = m.ExcPending
	if v1419 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	v3464 = v1414
	goto L213
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v3464
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1])) = v39
	return
L214:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[10]))
	if v1421 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1424 = v1421
	v1425 = v1418
	v1426 = v1287
	goto L218
L216:
	;
	v1582 = v1418
	goto L217
L217:
	;
	v1611 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1582))) = uint8(v1611)
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v1613, int64(-65533), v1418)
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L1
	} else {
		goto L252
	}
L218:
	;
	v1455 = v1424 + int32(24)
	if v1426 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v1582 = v1576
	goto L217
L220:
	;
	v1575 = v1571 + (v1568 - v1425) + int32(1)
	v1576 = v1425 + v1575
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1424)))
	if v1578 != 0 {
		v1424 = v1578
		v1425 = v1576
		v1426 = v1426 - v1575
		goto L218
	} else {
		goto L251
	}
L221:
	;
	v1571 = F_strlen(m, v1567)
	mBase = m.M
	goto L220
L222:
	;
	v1567 = v1455
	v1568 = v1425
	goto L221
L223:
	;
	goto L224
L224:
	;
	v1461 = v1426 - int32(1)
	if (v1425^v1455)&int32(3) != 0 {
		goto L228
	} else {
		goto L229
	}
L225:
	;
	v1564 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1561))) = uint8(v1564)
	v1567 = v1560
	v1568 = v1561
	goto L221
L226:
	;
	v1545 = v1540
	v1546 = v1541
	v1547 = v1542
	goto L247
L227:
	;
	if v1535 == int32(0) {
		v1560 = v1533
		v1561 = v1534
		goto L225
	} else {
		goto L246
	}
L228:
	;
	v1533 = v1455
	v1534 = v1425
	v1535 = v1461
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1465 = int32(0)
	if base.B2i32(v1455&int32(3) == v1465)|base.B2i32(v1461 == v1465) == v1465 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	if v1501 == int32(0) {
		v1560 = v1498
		v1561 = v1499
		goto L225
	} else {
		goto L240
	}
L232:
	;
	v1477 = v1455
	v1478 = v1425
	v1479 = v1461
	goto L235
L233:
	;
	goto L234
L234:
	;
	v1498 = v1455
	v1499 = v1425
	v1500 = v1461
	v1501 = base.B2i32(v1461 != v1465)
	goto L231
L235:
	;
	v1481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1477))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1478))) = uint8(v1481)
	if v1481 == int32(0) {
		v1540 = v1477
		v1541 = v1478
		v1542 = v1479
		goto L226
	} else {
		goto L237
	}
L236:
	;
	v1498 = v1492
	v1499 = v1486
	v1500 = v1488
	v1501 = v1490
	goto L231
L237:
	;
	v1485 = int32(1)
	v1486 = v1478 + v1485
	v1488 = v1479 - v1485
	v1489 = int32(0)
	v1490 = base.B2i32(v1488 != v1489)
	v1492 = v1477 + v1485
	if v1492&int32(3) == v1489 {
		v1498 = v1492
		v1499 = v1486
		v1500 = v1488
		v1501 = v1490
		goto L231
	} else {
		goto L238
	}
L238:
	;
	if v1488 != 0 {
		v1477 = v1492
		v1478 = v1486
		v1479 = v1488
		goto L235
	} else {
		goto L239
	}
L239:
	;
	goto L236
L240:
	;
	v1504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498))))
	if base.B2i32(v1504 == int32(0))|base.B2i32(base.Ui32(v1500) < base.Ui32(int32(4))) != 0 {
		v1533 = v1498
		v1534 = v1499
		v1535 = v1500
		goto L227
	} else {
		goto L241
	}
L241:
	;
	v1511 = v1498
	v1512 = v1499
	v1513 = v1500
	goto L242
L242:
	;
	v1516 = *(*int32)(unsafe.Add(mBase, uint32(v1511)))
	v1519 = int32(-2139062144)
	if (int32(16843008)-v1516|v1516)&v1519 != v1519 {
		v1540 = v1511
		v1541 = v1512
		v1542 = v1513
		goto L226
	} else {
		goto L244
	}
L243:
	;
	v1533 = v1527
	v1534 = v1525
	v1535 = v1529
	goto L227
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1512))) = v1516
	v1524 = int32(4)
	v1525 = v1512 + v1524
	v1527 = v1511 + v1524
	v1529 = v1513 - v1524
	if base.Ui32(int32(3)) < base.Ui32(v1529) {
		v1511 = v1527
		v1512 = v1525
		v1513 = v1529
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v1540 = v1533
	v1541 = v1534
	v1542 = v1535
	goto L226
L247:
	;
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1545))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1546))) = uint8(v1549)
	if v1549 == int32(0) {
		v1560 = v1545
		v1561 = v1546
		goto L225
	} else {
		goto L249
	}
L248:
	;
	v1560 = v1556
	v1561 = v1554
	goto L225
L249:
	;
	v1553 = int32(1)
	v1554 = v1546 + v1553
	v1556 = v1545 + v1553
	v1558 = v1547 - v1553
	if v1558 != 0 {
		v1545 = v1556
		v1546 = v1554
		v1547 = v1558
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
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1618 = F_shm_toc_allocate(m, v1617, v1284)
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L1
	} else {
		goto L253
	}
L253:
	;
	v1620 = m.G0
	v1622 = v1620 - int32(112)
	m.G0 = v1622
	v1624 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+108)) = v1618 + v1624
	v1628 = v1284 - v1624
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+104)) = v1628
	v1631 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[11]))
	v1632 = int32(0)
	if base.B2i32(v1631 == v1632)|base.B2i32(v1631 == int32(_a_F_InitializeParallelDSM_8)) == v1632 {
		goto L256
	} else {
		goto L257
	}
L254:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2050, int64(-65532), v1618)
	mBase = m.M
	v2053 = m.ExcPending
	if v2053 != 0 {
		goto L1
	} else {
		goto L307
	}
L255:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L1
	} else {
		goto L304
	}
L256:
	;
	v1640 = v1628
	v1643 = v1631
	goto L259
L257:
	;
	v1965 = v1628
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1618))) = v1628 - v1965
	m.G0 = v1622 + int32(112)
	goto L254
L259:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v1643+int32(-64))))
	if base.Ui32(v1673) < base.Ui32(int32(2)) {
		v1930 = v1640
		goto L261
	} else {
		goto L262
	}
L260:
	;
	v1965 = v1930
	goto L258
L261:
	;
	v1961 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+4))
	if v1961 != int32(_a_F_InitializeParallelDSM_8) {
		v1640 = v1930
		v1643 = v1961
		goto L259
	} else {
		goto L303
	}
L262:
	;
	v1677 = v1643 - int32(36)
	v1678 = *(*int32)(unsafe.Add(mBase, uint32(v1677)))
	if v1678 == int32(0) {
		v1930 = v1640
		goto L261
	} else {
		goto L263
	}
L263:
	;
	v1682 = v1643 - int32(68)
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(v1682)))
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+96)) = v1683
	F_do_serialize(m, v1622+int32(108), v1622+int32(104), int32(_a_F_InitializeParallelDSM_13), v1622+int32(96))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L1
	} else {
		goto L264
	}
L264:
	;
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v1643-int32(44))))
	switch v1696 {
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
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+20))
	if v1872 != 0 {
		goto L291
	} else {
		goto L292
	}
L266:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+28))
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1749)))
	v1751 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+36))
	if v1751 == int32(0) {
		goto L255
	} else {
		goto L281
	}
L267:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+28))
	v1736 = *(*int32)(unsafe.Add(mBase, uint32(v1735)))
	if v1736 != 0 {
		goto L277
	} else {
		goto L278
	}
L268:
	;
	v1721 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+28))
	v1722 = *(*float64)(unsafe.Add(mBase, uint32(v1721)))
	*(*float64)(unsafe.Add(mBase, uint32(v1622)+40)) = v1722
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+32)) = int32(17)
	F_do_serialize(m, v1622+int32(108), v1622+int32(104), int32(_a_F_InitializeParallelDSM_14), v1622+int32(32))
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L1
	} else {
		goto L276
	}
L269:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+28))
	v1710 = *(*int32)(unsafe.Add(mBase, uint32(v1709)))
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+16)) = v1710
	F_do_serialize(m, v1622+int32(108), v1622+int32(104), int32(_a_F_InitializeParallelDSM_15), v1622+int32(16))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L1
	} else {
		goto L275
	}
L270:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+28))
	v1704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1703))))
	if v1704 != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1705 = int32(_a_F_InitializeParallelDSM_16)
	goto L273
L272:
	;
	v1705 = int32(_a_F_InitializeParallelDSM_17)
	goto L273
L273:
	;
	F_do_serialize(m, v1622+int32(108), v1622+int32(104), v1705, int32(0))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
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
	v1738 = v1736
	goto L279
L278:
	;
	v1738 = int32(_a_F_InitializeParallelDSM_18)
	goto L279
L279:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+48)) = v1738
	F_do_serialize(m, v1622+int32(108), v1622+int32(104), int32(_a_F_InitializeParallelDSM_13), v1622+int32(48))
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L1
	} else {
		goto L280
	}
L280:
	;
	goto L265
L281:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v1751)))
	if v1754 == int32(0) {
		goto L255
	} else {
		goto L282
	}
L282:
	;
	v1757 = *(*int32)(unsafe.Add(mBase, uint32(v1751)+4))
	if v1750 != v1757 {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v1760 = v1751
	goto L286
L284:
	;
	v1804 = v1754
	goto L285
L285:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+80)) = v1804
	F_do_serialize(m, v1622+int32(108), v1622+int32(104), int32(_a_F_InitializeParallelDSM_13), v1622+int32(80))
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L1
	} else {
		goto L290
	}
L286:
	;
	v1791 = *(*int32)(unsafe.Add(mBase, uint32(v1760)+12))
	if v1791 == int32(0) {
		goto L255
	} else {
		goto L288
	}
L287:
	;
	v1804 = v1791
	goto L285
L288:
	;
	v1794 = *(*int32)(unsafe.Add(mBase, uint32(v1760)+16))
	if v1750 != v1794 {
		v1760 = v1760 + int32(12)
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
	v1874 = v1872
	goto L293
L292:
	;
	v1874 = int32(_a_F_InitializeParallelDSM_18)
	goto L293
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1622))) = v1874
	F_do_serialize(m, v1622+int32(108), v1622+int32(104), int32(_a_F_InitializeParallelDSM_13), v1622)
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L1
	} else {
		goto L294
	}
L294:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+20))
	if v1883 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	if base.Ui32(v1903) <= base.Ui32(int32(3)) {
		goto L210
	} else {
		goto L300
	}
L296:
	;
	v1900 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+104))
	v1903 = v1900
	goto L295
L297:
	;
	v1886 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1883))))
	if v1886 == int32(0) {
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v1889 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+104))
	if base.Ui32(v1889) <= base.Ui32(int32(3)) {
		goto L210
	} else {
		goto L299
	}
L299:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+108))
	v1893 = *(*int32)(unsafe.Add(mBase, uint32(v1643)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1892))) = v1893
	v1895 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+108)) = v1892 + v1895
	v1903 = v1889 - v1895
	goto L295
L300:
	;
	v1906 = *(*int32)(unsafe.Add(mBase, uint32(v1622)+108))
	v1907 = *(*int32)(unsafe.Add(mBase, uint32(v1677)))
	*(*int32)(unsafe.Add(mBase, uint32(v1906))) = v1907
	v1910 = v1903 & int32(-4)
	if v1910 == int32(4) {
		goto L210
	} else {
		goto L301
	}
L301:
	;
	v1915 = *(*int32)(unsafe.Add(mBase, uint32(v1643-int32(28))))
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+4)) = v1915
	if v1910 == int32(8) {
		goto L210
	} else {
		goto L302
	}
L302:
	;
	v1921 = *(*int32)(unsafe.Add(mBase, uint32(v1643-int32(20))))
	*(*int32)(unsafe.Add(mBase, uint32(v1906)+8)) = v1921
	v1923 = int32(12)
	v1924 = v1903 - v1923
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+104)) = v1924
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+108)) = v1906 + v1923
	v1930 = v1924
	goto L261
L303:
	;
	goto L260
L304:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v1682)))
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+68)) = v2037
	*(*int32)(unsafe.Add(mBase, uint32(v1622)+64)) = v1750
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_9), v1622-int32(-64))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(2938), int32(_a_F_InitializeParallelDSM_11))
	mBase = m.M
	v2049 = m.ExcPending
	if v2049 != 0 {
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
	v2054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2055 = F_shm_toc_allocate(m, v2054, v1295)
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v2058 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v2055))) = v2058
	v2061 = v2058 << (uint(int32(3)) % 32)
	if base.Ui32(v2061|int32(4)) <= base.Ui32(v1295) {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	v2088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2088, int64(-65531), v2055)
	mBase = m.M
	v2091 = m.ExcPending
	if v2091 != 0 {
		goto L1
	} else {
		goto L317
	}
L310:
	;
	v2065 = int32(0)
	if base.B2i32(v2061 == v2065)|base.B2i32(v2058 <= v2065) != 0 {
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
	v2078 = m.ExcPending
	if v2078 != 0 {
		goto L1
	} else {
		goto L314
	}
L313:
	;
	v2073 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[37]))
	base.MemoryCopy(m, v2055+int32(4), v2073, v2061)
	goto L309
L314:
	;
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_19), int32(0))
	mBase = m.M
	v2082 = m.ExcPending
	if v2082 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_20), int32(327), int32(_a_F_InitializeParallelDSM_21))
	mBase = m.M
	v2087 = m.ExcPending
	if v2087 != 0 {
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
	v2093 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[13]))
	if int32(2) <= v2093 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2097 = F_shm_toc_allocate(m, v2096, v1305)
	mBase = m.M
	v2098 = m.ExcPending
	if v2098 != 0 {
		goto L1
	} else {
		goto L321
	}
L319:
	;
	goto L320
L320:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2158 = F_shm_toc_allocate(m, v2157, v1306)
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L1
	} else {
		goto L336
	}
L321:
	;
	v2099 = int32(0)
	v2105 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v2106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+29)))
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v2108 = *(*int64)(unsafe.Add(mBase, uint32(v33)+4))
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v2110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2097)+16)) = uint8(v2110)
	*(*uint16)(unsafe.Add(mBase, uint32(v2097)+18)) = uint16(v2099)
	*(*int32)(unsafe.Add(mBase, uint32(v2097)+20)) = v2109
	*(*int64)(unsafe.Add(mBase, uint32(v2097))) = v2108
	*(*int32)(unsafe.Add(mBase, uint32(v2097)+8)) = v2107
	*(*uint8)(unsafe.Add(mBase, uint32(v2097)+17)) = uint8(v2106)
	if v2106&int32(1) != 0 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2152, int64(-65530), v2097)
	mBase = m.M
	v2155 = m.ExcPending
	if v2155 != 0 {
		goto L1
	} else {
		goto L335
	}
L323:
	;
	v2121 = v2105
	goto L325
L324:
	;
	v2121 = v2099
	goto L325
L325:
	;
	if v2110 != 0 {
		goto L326
	} else {
		goto L327
	}
L326:
	;
	v2122 = v2121
	goto L328
L327:
	;
	v2122 = v2105
	goto L328
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2097)+12)) = v2122
	v2124 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	if v2124 == int32(0) {
		goto L329
	} else {
		goto L330
	}
L329:
	;
	if v2122 <= int32(0) {
		goto L332
	} else {
		goto L333
	}
L330:
	;
	v2128 = v2124 << (uint(int32(2)) % 32)
	if v2128 == int32(0) {
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v2133 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	base.MemoryCopy(m, v2097+int32(24), v2133, v2128)
	goto L329
L332:
	;
	goto L322
L333:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v2140 = v2138 << (uint(int32(2)) % 32)
	if v2140 == int32(0) {
		goto L332
	} else {
		goto L334
	}
L334:
	;
	v2143 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
	base.MemoryCopy(m, v2097+v2143<<(uint(int32(2))%32)+int32(24), v2149, v2140)
	goto L332
L335:
	;
	goto L320
L336:
	;
	v2160 = int32(0)
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v2167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+29)))
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v2169 = *(*int64)(unsafe.Add(mBase, uint32(v37)+4))
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v2171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v2158)+16)) = uint8(v2171)
	*(*uint16)(unsafe.Add(mBase, uint32(v2158)+18)) = uint16(v2160)
	*(*int32)(unsafe.Add(mBase, uint32(v2158)+20)) = v2170
	*(*int64)(unsafe.Add(mBase, uint32(v2158))) = v2169
	*(*int32)(unsafe.Add(mBase, uint32(v2158)+8)) = v2168
	*(*uint8)(unsafe.Add(mBase, uint32(v2158)+17)) = uint8(v2167)
	if v2167&int32(1) != 0 {
		goto L338
	} else {
		goto L339
	}
L337:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2213, int64(-65529), v2158)
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L350
	}
L338:
	;
	v2182 = v2166
	goto L340
L339:
	;
	v2182 = v2160
	goto L340
L340:
	;
	if v2171 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v2183 = v2182
	goto L343
L342:
	;
	v2183 = v2166
	goto L343
L343:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2158)+12)) = v2183
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v2185 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	if v2183 <= int32(0) {
		goto L347
	} else {
		goto L348
	}
L345:
	;
	v2189 = v2185 << (uint(int32(2)) % 32)
	if v2189 == int32(0) {
		goto L344
	} else {
		goto L346
	}
L346:
	;
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	base.MemoryCopy(m, v2158+int32(24), v2194, v2189)
	goto L344
L347:
	;
	goto L337
L348:
	;
	v2199 = *(*int32)(unsafe.Add(mBase, uint32(v37)+24))
	v2201 = v2199 << (uint(int32(2)) % 32)
	if v2201 == int32(0) {
		goto L347
	} else {
		goto L349
	}
L349:
	;
	v2204 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v2210 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	base.MemoryCopy(m, v2158+v2204<<(uint(int32(2))%32)+int32(24), v2210, v2201)
	goto L347
L350:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2219 = F_shm_toc_allocate(m, v2217, int32(4))
	mBase = m.M
	v2220 = m.ExcPending
	if v2220 != 0 {
		goto L1
	} else {
		goto L351
	}
L351:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2219))) = v1288
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2222, int64(-65526), v2219)
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L1
	} else {
		goto L352
	}
L352:
	;
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2227 = F_shm_toc_allocate(m, v2226, v1307)
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L1
	} else {
		goto L353
	}
L353:
	;
	v2229 = int32(0)
	v2231 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v2227))) = v2231
	v2234 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[38])))
	*(*uint8)(unsafe.Add(mBase, uint32(v2227)+4)) = uint8(v2234)
	v2237 = *(*int64)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[39]))
	*(*int64)(unsafe.Add(mBase, uint32(v2227)+8)) = v2237
	v2240 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[14]))
	v2241 = *(*int64)(unsafe.Add(mBase, uint32(v2240)))
	*(*int64)(unsafe.Add(mBase, uint32(v2227)+16)) = v2241
	v2244 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[40]))
	*(*int32)(unsafe.Add(mBase, uint32(v2227)+24)) = v2244
	v2247 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[41]))
	if v2229 < v2247 {
		goto L355
	} else {
		goto L356
	}
L354:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2437, int64(-65528), v2227)
	mBase = m.M
	v2440 = m.ExcPending
	if v2440 != 0 {
		goto L1
	} else {
		goto L385
	}
L355:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2227)+28)) = v2247
	v2252 = v2247 << (uint(int32(2)) % 32)
	if v2252 == int32(0) {
		goto L354
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v2264 = v2240
	v2266 = v2229
	goto L359
L358:
	;
	v2258 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[42]))
	base.MemoryCopy(m, v2227+int32(32), v2258, v2252)
	goto L354
L359:
	;
	v2292 = *(*int32)(unsafe.Add(mBase, uint32(v2264)))
	if v2292 != 0 {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	v2303 = v2298 << (uint(int32(2)) % 32)
	v2304 = F_palloc(m, v2303)
	mBase = m.M
	v2305 = m.ExcPending
	if v2305 != 0 {
		goto L1
	} else {
		goto L367
	}
L361:
	;
	v2294 = F_add_size(m, v2266, int32(1))
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L1
	} else {
		goto L364
	}
L362:
	;
	v2296 = v2266
	goto L363
L363:
	;
	v2297 = *(*int32)(unsafe.Add(mBase, uint32(v2264)+52))
	v2298 = F_add_size(m, v2296, v2297)
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L1
	} else {
		goto L365
	}
L364:
	;
	v2296 = v2294
	goto L363
L365:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(v2264)+80))
	if v2300 != 0 {
		v2264 = v2300
		v2266 = v2298
		goto L359
	} else {
		goto L366
	}
L366:
	;
	goto L360
L367:
	;
	v2307 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[14]))
	if v2307 != 0 {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v2311 = int32(0)
	v2312 = v2307
	goto L371
L369:
	;
	goto L370
L370:
	;
	F_pg_qsort(m, v2304, v2298, int32(4), int32(187))
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L1
	} else {
		goto L383
	}
L371:
	;
	v2340 = *(*int32)(unsafe.Add(mBase, uint32(v2312)))
	if v2340 != 0 {
		goto L373
	} else {
		goto L374
	}
L372:
	;
	goto L370
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2304+v2311<<(uint(int32(2))%32)))) = v2340
	v2347 = v2311 + int32(1)
	goto L375
L374:
	;
	v2347 = v2311
	goto L375
L375:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+52))
	if int32(0) < v2348 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v2352 = v2348 << (uint(int32(2)) % 32)
	if v2352 != 0 {
		goto L379
	} else {
		goto L380
	}
L377:
	;
	v2360 = v2348
	goto L378
L378:
	;
	v2362 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+80))
	if v2362 != 0 {
		v2311 = v2360 + v2347
		v2312 = v2362
		goto L371
	} else {
		goto L382
	}
L379:
	;
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+48))
	base.MemoryCopy(m, v2304+v2347<<(uint(int32(2))%32), v2356, v2352)
	goto L381
L380:
	;
	goto L381
L381:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+52))
	v2360 = v2358
	goto L378
L382:
	;
	goto L372
L383:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2227)+28)) = v2298
	if v2303 == int32(0) {
		goto L354
	} else {
		goto L384
	}
L384:
	;
	base.MemoryCopy(m, v2227+int32(32), v2304, v2303)
	goto L354
L385:
	;
	v2441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2442 = F_shm_toc_allocate(m, v2441, v1308)
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v2444 = m.G0
	v2446 = v2444 - int32(80)
	m.G0 = v2446
	v2449 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[15]))
	if v2449 != 0 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2446)+40)) = int64(51539607564)
	v2453 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v2446)+68)) = v2453
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(v2449)))
	v2458 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+8))
	v2459 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+808))
	if v2459 != int64(0) {
		goto L391
	} else {
		goto L392
	}
L388:
	;
	v2769 = v2442
	goto L389
L389:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2769)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v2769))) = int64(0)
	m.G0 = v2446 + int32(80)
	v2805 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2805, int64(-65525), v2442)
	mBase = m.M
	v2808 = m.ExcPending
	if v2808 != 0 {
		goto L1
	} else {
		goto L425
	}
L390:
	;
	v2528 = F_hash_create(m, int32(_a_F_InitializeParallelDSM_22), v2524, v2446+int32(32), int32(1064))
	mBase = m.M
	v2529 = m.ExcPending
	if v2529 != 0 {
		goto L1
	} else {
		goto L394
	}
L391:
	;
	v2462 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+752))
	v2463 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+728))
	v2464 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+704))
	v2465 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+680))
	v2466 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+656))
	v2467 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+632))
	v2468 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+608))
	v2469 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+584))
	v2470 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+560))
	v2471 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+536))
	v2472 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+512))
	v2473 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+488))
	v2474 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+464))
	v2475 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+440))
	v2476 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+416))
	v2477 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+392))
	v2478 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+368))
	v2479 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+344))
	v2480 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+320))
	v2481 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+296))
	v2482 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+272))
	v2483 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+248))
	v2484 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+224))
	v2485 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+200))
	v2486 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+176))
	v2487 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+152))
	v2488 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+128))
	v2489 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+104))
	v2490 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+80))
	v2491 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+56))
	v2492 = *(*int64)(unsafe.Add(mBase, uint32(v2457)+32))
	v2524 = v2462 + (v2463 + (v2464 + (v2465 + (v2466 + (v2467 + (v2468 + (v2469 + (v2470 + (v2471 + (v2472 + (v2473 + (v2474 + (v2475 + (v2476 + (v2477 + (v2478 + (v2479 + (v2480 + (v2481 + (v2482 + (v2483 + (v2484 + (v2485 + (v2486 + (v2487 + (v2488 + (v2489 + (v2490 + (v2491 + (v2492 + v2458))))))))))))))))))))))))))))))
	goto L393
L392:
	;
	v2524 = v2458
	goto L393
L393:
	;
	goto L390
L394:
	;
	v2531 = v2446 + int32(12)
	v2533 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[15]))
	F_hash_seq_init(m, v2531, v2533)
	mBase = m.M
	v2535 = m.ExcPending
	if v2535 != 0 {
		goto L1
	} else {
		goto L395
	}
L395:
	;
	v2536 = F_hash_seq_search(m, v2531)
	mBase = m.M
	v2537 = m.ExcPending
	if v2537 != 0 {
		goto L1
	} else {
		goto L396
	}
L396:
	;
	if v2536 != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2542 = v2536
	goto L400
L398:
	;
	goto L399
L399:
	;
	v2611 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[43]))
	if v2611 != 0 {
		goto L405
	} else {
		goto L406
	}
L400:
	;
	v2572 = F_hash_search(m, v2528, v2542, int32(1), int32(0))
	mBase = m.M
	v2573 = m.ExcPending
	if v2573 != 0 {
		goto L1
	} else {
		goto L402
	}
L401:
	;
	goto L399
L402:
	;
	v2576 = F_hash_seq_search(m, v2446+int32(12))
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L1
	} else {
		goto L403
	}
L403:
	;
	if v2576 != 0 {
		v2542 = v2576
		goto L400
	} else {
		goto L404
	}
L404:
	;
	goto L401
L405:
	;
	v2616 = v2611
	goto L408
L406:
	;
	goto L407
L407:
	;
	v2685 = v2446 + int32(12)
	F_hash_seq_init(m, v2685, v2528)
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L1
	} else {
		goto L415
	}
L408:
	;
	v2644 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2616)+16)))
	if v2644 == int32(1) {
		goto L410
	} else {
		goto L411
	}
L409:
	;
	goto L407
L410:
	;
	v2649 = F_hash_search(m, v2528, v2616, int32(2), int32(0))
	mBase = m.M
	v2650 = m.ExcPending
	if v2650 != 0 {
		goto L1
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v2651 = *(*int32)(unsafe.Add(mBase, uint32(v2616)+24))
	if v2651 != 0 {
		v2616 = v2651
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
	v2688 = F_hash_seq_search(m, v2685)
	mBase = m.M
	v2689 = m.ExcPending
	if v2689 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	if v2688 != 0 {
		goto L417
	} else {
		goto L418
	}
L417:
	;
	v2693 = v2442
	v2694 = v2688
	goto L420
L418:
	;
	v2735 = v2442
	goto L419
L419:
	;
	F_hash_destroy(m, v2528)
	mBase = m.M
	v2765 = m.ExcPending
	if v2765 != 0 {
		goto L1
	} else {
		goto L424
	}
L420:
	;
	v2722 = *(*int32)(unsafe.Add(mBase, uint32(v2694)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2693)+8)) = v2722
	v2724 = *(*int64)(unsafe.Add(mBase, uint32(v2694)))
	*(*int64)(unsafe.Add(mBase, uint32(v2693))) = v2724
	v2726 = int32(12)
	v2727 = v2693 + v2726
	v2730 = F_hash_seq_search(m, v2446+v2726)
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L1
	} else {
		goto L422
	}
L421:
	;
	v2735 = v2727
	goto L419
L422:
	;
	if v2730 != 0 {
		v2693 = v2727
		v2694 = v2730
		goto L420
	} else {
		goto L423
	}
L423:
	;
	goto L421
L424:
	;
	v2769 = v2735
	goto L389
L425:
	;
	v2809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2810 = F_shm_toc_allocate(m, v2809, v1309)
	mBase = m.M
	v2811 = m.ExcPending
	if v2811 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	v2814 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[44]))
	*(*int32)(unsafe.Add(mBase, uint32(v2810))) = v2814
	v2817 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[45]))
	*(*int32)(unsafe.Add(mBase, uint32(v2810)+4)) = v2817
	v2820 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[16]))
	if v2820 != 0 {
		goto L428
	} else {
		goto L429
	}
L427:
	;
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2905, int64(-65524), v2810)
	mBase = m.M
	v2908 = m.ExcPending
	if v2908 != 0 {
		goto L1
	} else {
		goto L435
	}
L428:
	;
	v2821 = *(*int32)(unsafe.Add(mBase, uint32(v2820)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2810)+8)) = v2821
	v2823 = *(*int32)(unsafe.Add(mBase, uint32(v2820)+4))
	if v2823 <= int32(0) {
		goto L427
	} else {
		goto L431
	}
L429:
	;
	goto L430
L430:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2810)+8)) = int32(0)
	goto L427
L431:
	;
	v2831 = int32(0)
	goto L432
L432:
	;
	v2861 = v2831 << (uint(int32(2)) % 32)
	v2863 = *(*int32)(unsafe.Add(mBase, uint32(v2820)+12))
	v2865 = *(*int32)(unsafe.Add(mBase, uint32(v2863+v2861)))
	*(*int32)(unsafe.Add(mBase, uint32(v2810+int32(12)+v2861))) = v2865
	v2868 = v2831 + int32(1)
	v2869 = *(*int32)(unsafe.Add(mBase, uint32(v2820)+4))
	if v2868 < v2869 {
		v2831 = v2868
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
	v2909 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2910 = F_shm_toc_allocate(m, v2909, v1314)
	mBase = m.M
	v2911 = m.ExcPending
	if v2911 != 0 {
		goto L1
	} else {
		goto L436
	}
L436:
	;
	v2913 = int32(524)
	base.MemoryCopy(m, v2910, int32(_a_F_InitializeParallelDSM_23), v2913)
	base.MemoryCopy(m, v2910+v2913, int32(_a_F_InitializeParallelDSM_24), v2913)
	v2920 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v2920, int64(-65523), v2910)
	mBase = m.M
	v2923 = m.ExcPending
	if v2923 != 0 {
		goto L1
	} else {
		goto L437
	}
L437:
	;
	v2924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v2925 = F_shm_toc_allocate(m, v2924, v1310)
	mBase = m.M
	v2926 = m.ExcPending
	if v2926 != 0 {
		goto L1
	} else {
		goto L438
	}
L438:
	;
	v2927 = m.G0
	v2929 = v2927 - int32(32)
	m.G0 = v2929
	v2932 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[17]))
	if v2932 == int32(0) {
		v2987 = v2925
		goto L439
	} else {
		goto L440
	}
L439:
	;
	v3015 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2987))) = v3015
	v3018 = v2987 + int32(4)
	v3020 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[18]))
	if v3020 == v3015 {
		v3075 = v3018
		goto L448
	} else {
		goto L449
	}
L440:
	;
	v2936 = v2929 + int32(12)
	F_hash_seq_init(m, v2936, v2932)
	mBase = m.M
	v2938 = m.ExcPending
	if v2938 != 0 {
		goto L1
	} else {
		goto L441
	}
L441:
	;
	v2939 = F_hash_seq_search(m, v2936)
	mBase = m.M
	v2940 = m.ExcPending
	if v2940 != 0 {
		goto L1
	} else {
		goto L442
	}
L442:
	;
	if v2939 == int32(0) {
		v2987 = v2925
		goto L439
	} else {
		goto L443
	}
L443:
	;
	v2944 = v2939
	v2947 = v2925
	goto L444
L444:
	;
	v2975 = *(*int32)(unsafe.Add(mBase, uint32(v2944)))
	*(*int32)(unsafe.Add(mBase, uint32(v2947))) = v2975
	v2978 = v2947 + int32(4)
	v2981 = F_hash_seq_search(m, v2929+int32(12))
	mBase = m.M
	v2982 = m.ExcPending
	if v2982 != 0 {
		goto L1
	} else {
		goto L446
	}
L445:
	;
	v2987 = v2978
	goto L439
L446:
	;
	if v2981 != 0 {
		v2944 = v2981
		v2947 = v2978
		goto L444
	} else {
		goto L447
	}
L447:
	;
	goto L445
L448:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3075))) = int32(0)
	m.G0 = v2929 + int32(32)
	v3108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3108, int64(-65522), v2925)
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L1
	} else {
		goto L457
	}
L449:
	;
	v3024 = v2929 + int32(12)
	F_hash_seq_init(m, v3024, v3020)
	mBase = m.M
	v3026 = m.ExcPending
	if v3026 != 0 {
		goto L1
	} else {
		goto L450
	}
L450:
	;
	v3027 = F_hash_seq_search(m, v3024)
	mBase = m.M
	v3028 = m.ExcPending
	if v3028 != 0 {
		goto L1
	} else {
		goto L451
	}
L451:
	;
	if v3027 == int32(0) {
		v3075 = v3018
		goto L448
	} else {
		goto L452
	}
L452:
	;
	v3032 = v3027
	v3035 = v3018
	goto L453
L453:
	;
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v3032)))
	*(*int32)(unsafe.Add(mBase, uint32(v3035))) = v3063
	v3066 = v3035 + int32(4)
	v3069 = F_hash_seq_search(m, v2929+int32(12))
	mBase = m.M
	v3070 = m.ExcPending
	if v3070 != 0 {
		goto L1
	} else {
		goto L455
	}
L454:
	;
	v3075 = v3066
	goto L448
L455:
	;
	if v3069 != 0 {
		v3032 = v3069
		v3035 = v3066
		goto L453
	} else {
		goto L456
	}
L456:
	;
	goto L454
L457:
	;
	v3112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3113 = F_shm_toc_allocate(m, v3112, v1290)
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L1
	} else {
		goto L458
	}
L458:
	;
	v3116 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[46]))
	v3118 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	if v3118 == int32(0) {
		goto L460
	} else {
		goto L461
	}
L459:
	;
	v3140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3140, int64(-65521), v3113)
	mBase = m.M
	v3143 = m.ExcPending
	if v3143 != 0 {
		goto L1
	} else {
		goto L466
	}
L460:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3113)+4)) = v3116
	*(*int32)(unsafe.Add(mBase, uint32(v3113))) = int32(-1)
	goto L459
L461:
	;
	goto L462
L462:
	;
	v3124 = F_strlen(m, v3118)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v3113)+4)) = v3116
	*(*int32)(unsafe.Add(mBase, uint32(v3113))) = v3124
	if v3124 < int32(0) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	goto L459
L464:
	;
	v3130 = v3124 + int32(1)
	if v3130 == int32(0) {
		goto L463
	} else {
		goto L465
	}
L465:
	;
	v3136 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[19]))
	base.MemoryCopy(m, v3113+int32(8), v3136, v3130)
	goto L463
L466:
	;
	v3145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3146 = F_palloc0_mul(m, int32(8), v3145)
	mBase = m.M
	v3147 = m.ExcPending
	if v3147 != 0 {
		goto L1
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v3146
	v3149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3152 = F_mul_size(m, int32(_a_F_InitializeParallelDSM_12), v3151)
	mBase = m.M
	v3153 = m.ExcPending
	if v3153 != 0 {
		goto L1
	} else {
		goto L468
	}
L468:
	;
	v3154 = F_shm_toc_allocate(m, v3149, v3152)
	mBase = m.M
	v3155 = m.ExcPending
	if v3155 != 0 {
		goto L1
	} else {
		goto L469
	}
L469:
	;
	v3156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) < v3156 {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v3161 = int32(0)
	goto L473
L471:
	;
	goto L472
L472:
	;
	v3260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3260, int64(-65534), v3154)
	mBase = m.M
	v3263 = m.ExcPending
	if v3263 != 0 {
		goto L1
	} else {
		goto L479
	}
L473:
	;
	v3194 = v3154 + v3161<<(uint(int32(14))%32)
	v3196 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v3194))), uint32(v3196))
	v3199 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v3194)+16)) = v3199
	*(*int64)(unsafe.Add(mBase, uint32(v3194)+4)) = v3199
	v3203 = int32(512)
	*(*uint16)(unsafe.Add(mBase, uint32(v3194)+36)) = uint16(v3203)
	*(*int64)(unsafe.Add(mBase, uint32(v3194)+24)) = v3199
	*(*int32)(unsafe.Add(mBase, uint32(v3194)+32)) = int32(_a_F_InitializeParallelDSM_25)
	goto L475
L474:
	;
	goto L472
L475:
	;
	v3213 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeParallelDSM[31]))
	F_shm_mq_set_receiver(m, v3194, v3213)
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L1
	} else {
		goto L476
	}
L476:
	;
	v3216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v3217 = F_shm_mq_attach(m, v3194, v3216)
	mBase = m.M
	v3218 = m.ExcPending
	if v3218 != 0 {
		goto L1
	} else {
		goto L477
	}
L477:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v3219+v3161<<(uint(int32(3))%32))+4)) = v3217
	v3225 = v3161 + int32(1)
	v3226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3225 < v3226 {
		v3161 = v3225
		goto L473
	} else {
		goto L478
	}
L478:
	;
	goto L474
L479:
	;
	v3264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3265 = F_strlen(m, v3264)
	mBase = m.M
	v3266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v3267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3268 = F_strlen(m, v3267)
	mBase = m.M
	v3272 = F_shm_toc_allocate(m, v3266, v3268+v3265+int32(2))
	mBase = m.M
	v3273 = m.ExcPending
	if v3273 != 0 {
		goto L1
	} else {
		goto L480
	}
L480:
	;
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if (v3274^v3272)&int32(3) != 0 {
		goto L484
	} else {
		goto L485
	}
L481:
	;
	v3351 = v3265 + v3272 + int32(1)
	v3352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if (v3352^v3351)&int32(3) != 0 {
		goto L505
	} else {
		goto L506
	}
L482:
	;
	goto L481
L483:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3329))) = uint8(v3328)
	if v3328&int32(255) == int32(0) {
		goto L482
	} else {
		goto L498
	}
L484:
	;
	v3280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3274))))
	v3327 = v3274
	v3328 = v3280
	v3329 = v3272
	goto L483
L485:
	;
	goto L486
L486:
	;
	if v3274&int32(3) != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v3284 = v3274
	v3286 = v3272
	goto L490
L488:
	;
	v3298 = v3274
	v3300 = v3272
	goto L489
L489:
	;
	v3302 = *(*int32)(unsafe.Add(mBase, uint32(v3298)))
	v3305 = int32(-2139062144)
	if (int32(16843008)-v3302|v3302)&v3305 != v3305 {
		v3327 = v3298
		v3328 = v3302
		v3329 = v3300
		goto L483
	} else {
		goto L494
	}
L490:
	;
	v3287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3284))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3286))) = uint8(v3287)
	if v3287 == int32(0) {
		goto L482
	} else {
		goto L492
	}
L491:
	;
	v3298 = v3294
	v3300 = v3292
	goto L489
L492:
	;
	v3291 = int32(1)
	v3292 = v3286 + v3291
	v3294 = v3284 + v3291
	if v3294&int32(3) != 0 {
		v3284 = v3294
		v3286 = v3292
		goto L490
	} else {
		goto L493
	}
L493:
	;
	goto L491
L494:
	;
	v3310 = v3298
	v3311 = v3302
	v3312 = v3300
	goto L495
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3312))) = v3311
	v3314 = int32(4)
	v3315 = v3312 + v3314
	v3317 = v3310 + v3314
	v3319 = *(*int32)(unsafe.Add(mBase, uint32(v3310)+4))
	v3322 = int32(-2139062144)
	if (int32(16843008)-v3319|v3319)&v3322 == v3322 {
		v3310 = v3317
		v3311 = v3319
		v3312 = v3315
		goto L495
	} else {
		goto L497
	}
L496:
	;
	v3327 = v3317
	v3328 = v3319
	v3329 = v3315
	goto L483
L497:
	;
	goto L496
L498:
	;
	v3336 = v3327
	v3338 = v3329
	goto L499
L499:
	;
	v3339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3336)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3338)+1)) = uint8(v3339)
	v3341 = int32(1)
	if v3339 != 0 {
		v3336 = v3336 + v3341
		v3338 = v3338 + v3341
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
	v3427 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_shm_toc_insert(m, v3427, int64(-65527), v3272)
	mBase = m.M
	v3430 = m.ExcPending
	if v3430 != 0 {
		goto L1
	} else {
		goto L523
	}
L503:
	;
	goto L502
L504:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v3407))) = uint8(v3406)
	if v3406&int32(255) == int32(0) {
		goto L503
	} else {
		goto L519
	}
L505:
	;
	v3358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3352))))
	v3405 = v3352
	v3406 = v3358
	v3407 = v3351
	goto L504
L506:
	;
	goto L507
L507:
	;
	if v3352&int32(3) != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v3362 = v3352
	v3364 = v3351
	goto L511
L509:
	;
	v3376 = v3352
	v3378 = v3351
	goto L510
L510:
	;
	v3380 = *(*int32)(unsafe.Add(mBase, uint32(v3376)))
	v3383 = int32(-2139062144)
	if (int32(16843008)-v3380|v3380)&v3383 != v3383 {
		v3405 = v3376
		v3406 = v3380
		v3407 = v3378
		goto L504
	} else {
		goto L515
	}
L511:
	;
	v3365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3362))))
	*(*uint8)(unsafe.Add(mBase, uint32(v3364))) = uint8(v3365)
	if v3365 == int32(0) {
		goto L503
	} else {
		goto L513
	}
L512:
	;
	v3376 = v3372
	v3378 = v3370
	goto L510
L513:
	;
	v3369 = int32(1)
	v3370 = v3364 + v3369
	v3372 = v3362 + v3369
	if v3372&int32(3) != 0 {
		v3362 = v3372
		v3364 = v3370
		goto L511
	} else {
		goto L514
	}
L514:
	;
	goto L512
L515:
	;
	v3388 = v3376
	v3389 = v3380
	v3390 = v3378
	goto L516
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3390))) = v3389
	v3392 = int32(4)
	v3393 = v3390 + v3392
	v3395 = v3388 + v3392
	v3397 = *(*int32)(unsafe.Add(mBase, uint32(v3388)+4))
	v3400 = int32(-2139062144)
	if (int32(16843008)-v3397|v3397)&v3400 == v3400 {
		v3388 = v3395
		v3389 = v3397
		v3390 = v3393
		goto L516
	} else {
		goto L518
	}
L517:
	;
	v3405 = v3395
	v3406 = v3397
	v3407 = v3393
	goto L504
L518:
	;
	goto L517
L519:
	;
	v3414 = v3405
	v3416 = v3407
	goto L520
L520:
	;
	v3417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3414)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v3416)+1)) = uint8(v3417)
	v3419 = int32(1)
	if v3417 != 0 {
		v3414 = v3414 + v3419
		v3416 = v3416 + v3419
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
	v3431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3464 = v3431
	goto L213
L524:
	;
	F_errmsg_internal(m, int32(_a_F_InitializeParallelDSM_26), int32(0))
	mBase = m.M
	v3478 = m.ExcPending
	if v3478 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	F_errfinish(m, int32(_a_F_InitializeParallelDSM_10), int32(_a_F_InitializeParallelDSM_27), int32(_a_F_InitializeParallelDSM_28))
	mBase = m.M
	v3483 = m.ExcPending
	if v3483 != 0 {
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
					F_errfinish(m, int32(_a_F_parallel_vacuum_process_one_index_1), int32(1112), int32(_a_F_parallel_vacuum_process_one_index_2))
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
