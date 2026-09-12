package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_standard_ExplainOneQuery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v186 int64
	_ = v186
	var v188 int64
	_ = v188
	var v189 int64
	_ = v189
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v196 int64
	_ = v196
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	v8 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(288)
	m.G0 = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
	if v17 == int32(1) {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v26 = F_AllocSetContextCreateInternal(m, v21, int32(66423), int32(0), int32(8192), int32(8388608))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = int32(4562096)
			v29 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
			v32 = v26
			v33 = v29
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
			if v34 == int32(1) {
				v41 = F__emscripten_memcpy_bulkmem(m, v15+int32(152), int32(4460328), int32(128))
				mBase = m.M
			} else {
			}
			F___clock_gettime(m, int32(1), v15+int32(24))
			mBase = m.M
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
			v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
			v49 = F_pg_plan_query(m, l0, l4, l1, l5)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v51 = int32(1)
				F___clock_gettime(m, v51, v15+int32(24))
				mBase = m.M
				v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
				v57 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v15)+280)) = v55 - v48 + (v57-v47)*int64(1000000000)
				v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
				if v63 == v51 {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
					F_MemoryContextMemConsumed(m, v32, v15+int32(8))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v75 == int32(1) {
							v83 = F__emscripten_memset_bulkmem(m, v15+int32(24), base.I32_extend8_s(int32(0)), int32(128))
							mBase = m.M
							v85 = v15 + int32(24)
							v87 = v15 + int32(152)
							v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v90 = *(*int64)(unsafe.Add(mBase, _consts[55]))
							v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
							v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							v97 = *(*int64)(unsafe.Add(mBase, _consts[56]))
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
							v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
							v104 = *(*int64)(unsafe.Add(mBase, _consts[57]))
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
							v111 = *(*int64)(unsafe.Add(mBase, _consts[58]))
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
							v118 = *(*int64)(unsafe.Add(mBase, _consts[59]))
							v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
							v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
							v125 = *(*int64)(unsafe.Add(mBase, _consts[60]))
							v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
							v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
							v132 = *(*int64)(unsafe.Add(mBase, _consts[61]))
							v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
							v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
							v139 = *(*int64)(unsafe.Add(mBase, _consts[62]))
							v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
							v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
							v146 = *(*int64)(unsafe.Add(mBase, _consts[63]))
							v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
							v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
							v153 = *(*int64)(unsafe.Add(mBase, _consts[64]))
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
							v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
							v160 = *(*int64)(unsafe.Add(mBase, _consts[65]))
							v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
							v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
							v167 = *(*int64)(unsafe.Add(mBase, _consts[66]))
							v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
							v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
							v174 = *(*int64)(unsafe.Add(mBase, _consts[67]))
							v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
							v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
							v181 = *(*int64)(unsafe.Add(mBase, _consts[68]))
							v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
							v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
							v188 = *(*int64)(unsafe.Add(mBase, _consts[69]))
							v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
							v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
							v195 = *(*int64)(unsafe.Add(mBase, _consts[70]))
							v196 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+120)) = v193 + (v195 - v196)
							v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
							if v203&int32(1) != 0 {
								v206 = v15 + int32(24)
							} else {
								v206 = int32(0)
							}
							v207 = v206
						} else {
							v207 = int32(0)
						}
						v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
						if v211 != 0 {
							v212 = v15 + int32(8)
						} else {
							v212 = int32(0)
						}
						F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v15+int32(280), v207, v212)
						mBase = m.M
						v214 = m.ExcPending
						if v214 != 0 {
							return
						} else {
							m.G0 = v15 + int32(288)
							return
						}
					}
				} else {
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v75 == int32(1) {
						v83 = F__emscripten_memset_bulkmem(m, v15+int32(24), base.I32_extend8_s(int32(0)), int32(128))
						mBase = m.M
						v85 = v15 + int32(24)
						v87 = v15 + int32(152)
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						v90 = *(*int64)(unsafe.Add(mBase, _consts[55]))
						v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
						v97 = *(*int64)(unsafe.Add(mBase, _consts[56]))
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
						v104 = *(*int64)(unsafe.Add(mBase, _consts[57]))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
						v111 = *(*int64)(unsafe.Add(mBase, _consts[58]))
						v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
						v118 = *(*int64)(unsafe.Add(mBase, _consts[59]))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
						v125 = *(*int64)(unsafe.Add(mBase, _consts[60]))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
						v132 = *(*int64)(unsafe.Add(mBase, _consts[61]))
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
						v139 = *(*int64)(unsafe.Add(mBase, _consts[62]))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
						v146 = *(*int64)(unsafe.Add(mBase, _consts[63]))
						v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
						v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
						v153 = *(*int64)(unsafe.Add(mBase, _consts[64]))
						v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
						v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
						v160 = *(*int64)(unsafe.Add(mBase, _consts[65]))
						v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
						v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
						v167 = *(*int64)(unsafe.Add(mBase, _consts[66]))
						v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
						v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
						v174 = *(*int64)(unsafe.Add(mBase, _consts[67]))
						v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
						v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
						v181 = *(*int64)(unsafe.Add(mBase, _consts[68]))
						v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
						v188 = *(*int64)(unsafe.Add(mBase, _consts[69]))
						v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
						v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
						v195 = *(*int64)(unsafe.Add(mBase, _consts[70]))
						v196 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+120)) = v193 + (v195 - v196)
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v203&int32(1) != 0 {
							v206 = v15 + int32(24)
						} else {
							v206 = int32(0)
						}
						v207 = v206
					} else {
						v207 = int32(0)
					}
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
					if v211 != 0 {
						v212 = v15 + int32(8)
					} else {
						v212 = int32(0)
					}
					F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v15+int32(280), v207, v212)
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return
					} else {
						m.G0 = v15 + int32(288)
						return
					}
				}
			}
		}
	} else {
		v32 = v8
		v33 = v8
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
		if v34 == int32(1) {
			v41 = F__emscripten_memcpy_bulkmem(m, v15+int32(152), int32(4460328), int32(128))
			mBase = m.M
		} else {
		}
		F___clock_gettime(m, int32(1), v15+int32(24))
		mBase = m.M
		v47 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
		v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
		v49 = F_pg_plan_query(m, l0, l4, l1, l5)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			v51 = int32(1)
			F___clock_gettime(m, v51, v15+int32(24))
			mBase = m.M
			v55 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+32)))
			v57 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v15)+280)) = v55 - v48 + (v57-v47)*int64(1000000000)
			v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
			if v63 == v51 {
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
				F_MemoryContextMemConsumed(m, v32, v15+int32(8))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v75 == int32(1) {
						v83 = F__emscripten_memset_bulkmem(m, v15+int32(24), base.I32_extend8_s(int32(0)), int32(128))
						mBase = m.M
						v85 = v15 + int32(24)
						v87 = v15 + int32(152)
						v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						v90 = *(*int64)(unsafe.Add(mBase, _consts[55]))
						v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
						v97 = *(*int64)(unsafe.Add(mBase, _consts[56]))
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
						v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
						v104 = *(*int64)(unsafe.Add(mBase, _consts[57]))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
						v111 = *(*int64)(unsafe.Add(mBase, _consts[58]))
						v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
						v118 = *(*int64)(unsafe.Add(mBase, _consts[59]))
						v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
						v125 = *(*int64)(unsafe.Add(mBase, _consts[60]))
						v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
						v132 = *(*int64)(unsafe.Add(mBase, _consts[61]))
						v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
						v139 = *(*int64)(unsafe.Add(mBase, _consts[62]))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
						v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
						v146 = *(*int64)(unsafe.Add(mBase, _consts[63]))
						v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
						v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
						v153 = *(*int64)(unsafe.Add(mBase, _consts[64]))
						v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
						v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
						v160 = *(*int64)(unsafe.Add(mBase, _consts[65]))
						v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
						v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
						v167 = *(*int64)(unsafe.Add(mBase, _consts[66]))
						v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
						v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
						v174 = *(*int64)(unsafe.Add(mBase, _consts[67]))
						v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
						v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
						v181 = *(*int64)(unsafe.Add(mBase, _consts[68]))
						v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
						v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
						v188 = *(*int64)(unsafe.Add(mBase, _consts[69]))
						v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
						v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
						v195 = *(*int64)(unsafe.Add(mBase, _consts[70]))
						v196 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+120)) = v193 + (v195 - v196)
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v203&int32(1) != 0 {
							v206 = v15 + int32(24)
						} else {
							v206 = int32(0)
						}
						v207 = v206
					} else {
						v207 = int32(0)
					}
					v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
					if v211 != 0 {
						v212 = v15 + int32(8)
					} else {
						v212 = int32(0)
					}
					F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v15+int32(280), v207, v212)
					mBase = m.M
					v214 = m.ExcPending
					if v214 != 0 {
						return
					} else {
						m.G0 = v15 + int32(288)
						return
					}
				}
			} else {
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
				if v75 == int32(1) {
					v83 = F__emscripten_memset_bulkmem(m, v15+int32(24), base.I32_extend8_s(int32(0)), int32(128))
					mBase = m.M
					v85 = v15 + int32(24)
					v87 = v15 + int32(152)
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
					v90 = *(*int64)(unsafe.Add(mBase, _consts[55]))
					v91 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
					*(*int64)(unsafe.Add(mBase, uint32(v85))) = v88 + (v90 - v91)
					v95 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
					v97 = *(*int64)(unsafe.Add(mBase, _consts[56]))
					v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = v95 + (v97 - v98)
					v102 = *(*int64)(unsafe.Add(mBase, uint32(v85)+16))
					v104 = *(*int64)(unsafe.Add(mBase, _consts[57]))
					v105 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v102 + (v104 - v105)
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
					v111 = *(*int64)(unsafe.Add(mBase, _consts[58]))
					v112 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+24)) = v109 + (v111 - v112)
					v116 = *(*int64)(unsafe.Add(mBase, uint32(v85)+32))
					v118 = *(*int64)(unsafe.Add(mBase, _consts[59]))
					v119 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+32)) = v116 + (v118 - v119)
					v123 = *(*int64)(unsafe.Add(mBase, uint32(v85)+40))
					v125 = *(*int64)(unsafe.Add(mBase, _consts[60]))
					v126 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+40)) = v123 + (v125 - v126)
					v130 = *(*int64)(unsafe.Add(mBase, uint32(v85)+48))
					v132 = *(*int64)(unsafe.Add(mBase, _consts[61]))
					v133 = *(*int64)(unsafe.Add(mBase, uint32(v87)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+48)) = v130 + (v132 - v133)
					v137 = *(*int64)(unsafe.Add(mBase, uint32(v85)+56))
					v139 = *(*int64)(unsafe.Add(mBase, _consts[62]))
					v140 = *(*int64)(unsafe.Add(mBase, uint32(v87)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+56)) = v137 + (v139 - v140)
					v144 = *(*int64)(unsafe.Add(mBase, uint32(v85)+64))
					v146 = *(*int64)(unsafe.Add(mBase, _consts[63]))
					v147 = *(*int64)(unsafe.Add(mBase, uint32(v87)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+64)) = v144 + (v146 - v147)
					v151 = *(*int64)(unsafe.Add(mBase, uint32(v85)+72))
					v153 = *(*int64)(unsafe.Add(mBase, _consts[64]))
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v87)+72))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+72)) = v151 + (v153 - v154)
					v158 = *(*int64)(unsafe.Add(mBase, uint32(v85)+80))
					v160 = *(*int64)(unsafe.Add(mBase, _consts[65]))
					v161 = *(*int64)(unsafe.Add(mBase, uint32(v87)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+80)) = v158 + (v160 - v161)
					v165 = *(*int64)(unsafe.Add(mBase, uint32(v85)+88))
					v167 = *(*int64)(unsafe.Add(mBase, _consts[66]))
					v168 = *(*int64)(unsafe.Add(mBase, uint32(v87)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+88)) = v165 + (v167 - v168)
					v172 = *(*int64)(unsafe.Add(mBase, uint32(v85)+96))
					v174 = *(*int64)(unsafe.Add(mBase, _consts[67]))
					v175 = *(*int64)(unsafe.Add(mBase, uint32(v87)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+96)) = v172 + (v174 - v175)
					v179 = *(*int64)(unsafe.Add(mBase, uint32(v85)+104))
					v181 = *(*int64)(unsafe.Add(mBase, _consts[68]))
					v182 = *(*int64)(unsafe.Add(mBase, uint32(v87)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+104)) = v179 + (v181 - v182)
					v186 = *(*int64)(unsafe.Add(mBase, uint32(v85)+112))
					v188 = *(*int64)(unsafe.Add(mBase, _consts[69]))
					v189 = *(*int64)(unsafe.Add(mBase, uint32(v87)+112))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+112)) = v186 + (v188 - v189)
					v193 = *(*int64)(unsafe.Add(mBase, uint32(v85)+120))
					v195 = *(*int64)(unsafe.Add(mBase, _consts[70]))
					v196 = *(*int64)(unsafe.Add(mBase, uint32(v87)+120))
					*(*int64)(unsafe.Add(mBase, uint32(v85)+120)) = v193 + (v195 - v196)
					v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v203&int32(1) != 0 {
						v206 = v15 + int32(24)
					} else {
						v206 = int32(0)
					}
					v207 = v206
				} else {
					v207 = int32(0)
				}
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
				if v211 != 0 {
					v212 = v15 + int32(8)
				} else {
					v212 = int32(0)
				}
				F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v15+int32(280), v207, v212)
				mBase = m.M
				v214 = m.ExcPending
				if v214 != 0 {
					return
				} else {
					m.G0 = v15 + int32(288)
					return
				}
			}
		}
	}
}
