package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SlabAlloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v113 int64
	_ = v113
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v220 int64
	_ = v220
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v9 == l1 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
		if v11 == int32(0) {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			if v14 != 0 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v17
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v19
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v21 - int32(1)
				v26 = v15 - int32(8)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v26))) = v28
					v43 = v27
				} else {
					v31 = v15 - int32(4)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32 + v33
					v37 = v15 - int32(12)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					*(*int32)(unsafe.Add(mBase, uint32(v37))) = v38 - int32(1)
					v43 = v32
				}
				v47 = v15 - int32(16)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
				v50 = v48 - int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = v50
				v76 = v50
				v78 = v15 - int32(20)
				v79 = v43
				v80 = int32(0)
				v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
				v84 = (v80 - v76) >> (uint(v83) % 32)
				v88 = l0 - v84<<(uint(int32(3))%32)
				v90 = v88 + int32(80)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+84))
				if v91 == v80 {
					*(*int32)(unsafe.Add(mBase, uint32(v90))) = v90
					v95 = v90
				} else {
					v95 = v91
				}
				*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v90
				*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v95
				v99 = v78 + int32(20)
				*(*int32)(unsafe.Add(mBase, uint32(v95))) = v99
				*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v99
				*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v80 - v84
				v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v113 = int64(5)
				*(*int64)(unsafe.Add(mBase, uint32(v79))) = base.I64_extend_i32_u(v79-v78)<<(uint(int64(34))%64) | base.I64_extend_i32_u((v107+int32(7))&int32(-8))<<(uint(v113)%64) | v113
				v126 = v79 + int32(8)
				return v126
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v53 = F_emscripten_builtin_malloc(m, v52)
				mBase = m.M
				if v53 == int32(0) {
					v56 = F_MemoryContextAllocationFailure(m, l0, l1, l2)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v126 = v56
						return v126
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v53))) = l0
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v61 + v52
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v66 = v64 - int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v66
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v66
					v73 = v53 + int32(32)
					*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v68 + v73
					v76 = v66
					v78 = v53
					v79 = v73
					v80 = int32(0)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
					v84 = (v80 - v76) >> (uint(v83) % 32)
					v88 = l0 - v84<<(uint(int32(3))%32)
					v90 = v88 + int32(80)
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+84))
					if v91 == v80 {
						*(*int32)(unsafe.Add(mBase, uint32(v90))) = v90
						v95 = v90
					} else {
						v95 = v91
					}
					*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v90
					*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v95
					v99 = v78 + int32(20)
					*(*int32)(unsafe.Add(mBase, uint32(v95))) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v90)+4)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v80 - v84
					v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v113 = int64(5)
					*(*int64)(unsafe.Add(mBase, uint32(v79))) = base.I64_extend_i32_u(v79-v78)<<(uint(int64(34))%64) | base.I64_extend_i32_u((v107+int32(7))&int32(-8))<<(uint(v113)%64) | v113
					v126 = v79 + int32(8)
					return v126
				}
			}
		} else {
			v129 = l0 + int32(80)
			v132 = v129 + v11<<(uint(int32(3))%32)
			v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
			v135 = v133 - int32(8)
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
			if v136 != 0 {
				v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v135))) = v137
				v151 = v136
			} else {
				v140 = v133 - int32(4)
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
				v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v140))) = v141 + v142
				v146 = v133 - int32(12)
				v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
				*(*int32)(unsafe.Add(mBase, uint32(v146))) = v147 - int32(1)
				v151 = v141
			}
			v156 = v133 - int32(16)
			v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
			v158 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v156))) = v157 - v158
			v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
			v166 = int32(0) - (v158-v157)>>(uint(v164)%32)
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			if v166 == v167 {
			} else {
				v169 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
				v170 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v170
				v172 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
				*(*int32)(unsafe.Add(mBase, uint32(v170))) = v172
				v176 = v129 + v166<<(uint(int32(3))%32)
				v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+4))
				if v177 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(v176))) = v176
					v181 = v176
				} else {
					v181 = v177
				}
				*(*int32)(unsafe.Add(mBase, uint32(v133))) = v176
				*(*int32)(unsafe.Add(mBase, uint32(v133)+4)) = v181
				*(*int32)(unsafe.Add(mBase, uint32(v181))) = v133
				*(*int32)(unsafe.Add(mBase, uint32(v176)+4)) = v133
				v186 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
				if v186 != v132 {
					v189 = v186
				} else {
					v189 = int32(0)
				}
				if v189 != 0 {
				} else {
					v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
					if v190 != 0 {
						if v190 != l0+int32(88) {
							v204 = int32(1)
						} else {
							v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
							if v195 != 0 {
								v202 = base.B2i32(v195 != l0+int32(96)) << (uint(int32(1)) % 32)
							} else {
								v202 = int32(0)
							}
							v204 = v202
						}
					} else {
						v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
						if v195 != 0 {
							v202 = base.B2i32(v195 != l0+int32(96)) << (uint(int32(1)) % 32)
						} else {
							v202 = int32(0)
						}
						v204 = v202
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v204
				}
			}
			v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v220 = int64(5)
			*(*int64)(unsafe.Add(mBase, uint32(v151))) = base.I64_extend_i32_u(v151-(v133-int32(20)))<<(uint(int64(34))%64) | base.I64_extend_i32_u((v214+int32(7))&int32(-8))<<(uint(v220)%64) | v220
			return v151 + int32(8)
		}
	} else {
		v229 = m.G0
		v231 = v229 - int32(16)
		m.G0 = v231
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v236 = m.ExcPending
		if v236 != 0 {
			return int32(0)
		} else {
			v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v237
			*(*int32)(unsafe.Add(mBase, uint32(v231))) = l1
			F_errmsg_internal(m, int32(645992), v231)
			mBase = m.M
			v242 = m.ExcPending
			if v242 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(491082), int32(611), int32(336280))
				mBase = m.M
				v247 = m.ExcPending
				if v247 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_SlabReset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v5 == int32(0) {
	} else {
		v9 = l0 + int32(68)
		if v5 == v9 {
		} else {
			v12 = v5
			for {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v16
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				*(*int32)(unsafe.Add(mBase, uint32(v16))) = v18
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v20 - int32(1)
				F_emscripten_builtin_free(m, v12-int32(20))
				mBase = m.M
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v27 - v28
				if v16 != v9 {
					v12 = v16
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v36 == int32(0) {
	} else {
		v40 = l0 + int32(80)
		if v36 == v40 {
		} else {
			v43 = v36
			for {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = v49
				F_emscripten_builtin_free(m, v43-int32(20))
				mBase = m.M
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v54 - v55
				if v47 != v40 {
					v43 = v47
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v63 == int32(0) {
	} else {
		v67 = l0 + int32(88)
		if v63 == v67 {
		} else {
			v70 = v63
			for {
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = v74
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
				*(*int32)(unsafe.Add(mBase, uint32(v74))) = v76
				F_emscripten_builtin_free(m, v70-int32(20))
				mBase = m.M
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v81 - v82
				if v74 != v67 {
					v70 = v74
					continue
				} else {
					break
				}
				break
			}
		}
	}
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v90 == int32(0) {
	} else {
		v94 = l0 + int32(96)
		if v90 == v94 {
		} else {
			v97 = v90
			for {
				v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
				v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v100)+4)) = v101
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
				*(*int32)(unsafe.Add(mBase, uint32(v101))) = v103
				F_emscripten_builtin_free(m, v97-int32(20))
				mBase = m.M
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v108 - v109
				if v101 != v94 {
					v97 = v101
					continue
				} else {
					break
				}
				break
			}
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = int32(0)
	return
}
func F_SlabStats(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(240)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v25 = v21*v22 + int32(104)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v26 == v6 {
		v68 = v6
		v69 = v25
		v70 = v6
		v71 = v6
	} else {
		v30 = l0 + int32(80)
		if v26 == v30 {
			v68 = v6
			v69 = v25
			v70 = v6
			v71 = v6
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v38 = v26
			v40 = v6
			v41 = v25
			v42 = v6
			v43 = v6
			for {
				v49 = v41 + v21
				v51 = v42 + int32(1)
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v38-int32(16))))
				v55 = v54 + v43
				v57 = v32*v54 + v40
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				if v58 != v30 {
					v38 = v58
					v40 = v57
					v41 = v49
					v42 = v51
					v43 = v55
					continue
				} else {
					break
				}
				break
			}
			v68 = v57
			v69 = v49
			v70 = v51
			v71 = v55
		}
	}
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	if v77 == int32(0) {
		v118 = v68
		v119 = v69
		v120 = v70
		v121 = v71
	} else {
		v81 = l0 + int32(88)
		if v77 == v81 {
			v118 = v68
			v119 = v69
			v120 = v70
			v121 = v71
		} else {
			v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v89 = v77
			v91 = v68
			v92 = v69
			v93 = v70
			v94 = v71
			for {
				v100 = v92 + v21
				v102 = v93 + int32(1)
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v89-int32(16))))
				v106 = v105 + v94
				v108 = v83*v105 + v91
				v109 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
				if v109 != v81 {
					v89 = v109
					v91 = v108
					v92 = v100
					v93 = v102
					v94 = v106
					continue
				} else {
					break
				}
				break
			}
			v118 = v108
			v119 = v100
			v120 = v102
			v121 = v106
		}
	}
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v127 == int32(0) {
		v168 = v118
		v169 = v119
		v170 = v120
		v171 = v121
	} else {
		v131 = l0 + int32(96)
		if v127 == v131 {
			v168 = v118
			v169 = v119
			v170 = v120
			v171 = v121
		} else {
			v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v139 = v127
			v141 = v118
			v142 = v119
			v143 = v120
			v144 = v121
			for {
				v150 = v142 + v21
				v152 = v143 + int32(1)
				v155 = *(*int32)(unsafe.Add(mBase, uint32(v139-int32(16))))
				v156 = v155 + v144
				v158 = v133*v155 + v141
				v159 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
				if v159 != v131 {
					v139 = v159
					v141 = v158
					v142 = v150
					v143 = v152
					v144 = v156
					continue
				} else {
					break
				}
				break
			}
			v168 = v158
			v169 = v150
			v170 = v152
			v171 = v156
		}
	}
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v171
		*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v168
		*(*int32)(unsafe.Add(mBase, uint32(v19))) = v169
		*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v169 - v168
		*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v170
		*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v22
		v188 = F_pg_snprintf(m, v19+int32(32), int32(200), int32(441402), v19)
		mBase = m.M
		v189 = m.ExcPending
		if v189 != 0 {
			return
		} else {
			m.T0[l1].(func(*base.Module, int32, int32, int32, int32))(m, l0, l2, v19+int32(32), l4)
			mBase = m.M
			v193 = m.ExcPending
			if v193 != 0 {
				return
			} else {
				if l3 != 0 {
					v194 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v194 + v170
					v197 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v197 + v171
					v200 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v200 + v169
					v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v203 + v168
				} else {
				}
				m.G0 = v19 + int32(240)
				return
			}
		}
	} else {
		if l3 != 0 {
			v194 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v194 + v170
			v197 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v197 + v171
			v200 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v200 + v169
			v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v203 + v168
		} else {
		}
		m.G0 = v19 + int32(240)
		return
	}
}
