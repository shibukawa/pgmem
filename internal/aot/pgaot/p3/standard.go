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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
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
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v162 int64
	_ = v162
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v173 int64
	_ = v173
	var v175 int64
	_ = v175
	var v176 int64
	_ = v176
	var v180 int64
	_ = v180
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(288)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
	if v18 == int32(1) {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0]))
		v27 = F_AllocSetContextCreateInternal(m, v22, int32(_a_F_standard_ExplainOneQuery_0), int32(0), int32(_a_F_standard_ExplainOneQuery_1), int32(_a_F_standard_ExplainOneQuery_2))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = int32(_a_F_standard_ExplainOneQuery_3)
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0])) = v27
			v33 = v27
			v34 = v30
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
			if v35 == int32(1) {
				base.MemoryCopy(m, v16+int32(152), int32(_a_F_standard_ExplainOneQuery_4), int32(128))
			} else {
			}
			v45 = v16 + int32(24)
			F___clock_gettime(m, int32(1), v45)
			mBase = m.M
			v47 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
			v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+32)))
			v49 = F_pg_plan_query(m, l0, l4, l1, l5)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				v51 = int32(1)
				F___clock_gettime(m, v51, v45)
				mBase = m.M
				v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+32)))
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
				*(*int64)(unsafe.Add(mBase, uint32(v16)+280)) = v53 - v48 + (v55-v47)*int64(1000000000)
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
				if v61 == v51 {
					*(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0])) = v34
					F_MemoryContextMemConsumed(m, v33, v16+int32(8))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v72 == int32(1) {
							v76 = v16 + int32(24)
							base.MemoryFill(m, v76, int32(0), int32(128))
							v81 = v16 + int32(152)
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
							v84 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[1]))
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
							*(*int64)(unsafe.Add(mBase, uint32(v76))) = v82 + (v84 - v85)
							v89 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
							v91 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[2]))
							v92 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v89 + (v91 - v92)
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
							v98 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[3]))
							v99 = *(*int64)(unsafe.Add(mBase, uint32(v81)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+16)) = v96 + (v98 - v99)
							v103 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
							v105 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[4]))
							v106 = *(*int64)(unsafe.Add(mBase, uint32(v81)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+24)) = v103 + (v105 - v106)
							v110 = *(*int64)(unsafe.Add(mBase, uint32(v76)+32))
							v112 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[5]))
							v113 = *(*int64)(unsafe.Add(mBase, uint32(v81)+32))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+32)) = v110 + (v112 - v113)
							v117 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
							v119 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[6]))
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v81)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+40)) = v117 + (v119 - v120)
							v124 = *(*int64)(unsafe.Add(mBase, uint32(v76)+48))
							v126 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[7]))
							v127 = *(*int64)(unsafe.Add(mBase, uint32(v81)+48))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+48)) = v124 + (v126 - v127)
							v131 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
							v133 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[8]))
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v81)+56))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+56)) = v131 + (v133 - v134)
							v138 = *(*int64)(unsafe.Add(mBase, uint32(v76)+64))
							v140 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[9]))
							v141 = *(*int64)(unsafe.Add(mBase, uint32(v81)+64))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+64)) = v138 + (v140 - v141)
							v145 = *(*int64)(unsafe.Add(mBase, uint32(v76)+72))
							v147 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[10]))
							v148 = *(*int64)(unsafe.Add(mBase, uint32(v81)+72))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+72)) = v145 + (v147 - v148)
							v152 = *(*int64)(unsafe.Add(mBase, uint32(v76)+80))
							v154 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[11]))
							v155 = *(*int64)(unsafe.Add(mBase, uint32(v81)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+80)) = v152 + (v154 - v155)
							v159 = *(*int64)(unsafe.Add(mBase, uint32(v76)+88))
							v161 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[12]))
							v162 = *(*int64)(unsafe.Add(mBase, uint32(v81)+88))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+88)) = v159 + (v161 - v162)
							v166 = *(*int64)(unsafe.Add(mBase, uint32(v76)+96))
							v168 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[13]))
							v169 = *(*int64)(unsafe.Add(mBase, uint32(v81)+96))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+96)) = v166 + (v168 - v169)
							v173 = *(*int64)(unsafe.Add(mBase, uint32(v76)+104))
							v175 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[14]))
							v176 = *(*int64)(unsafe.Add(mBase, uint32(v81)+104))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+104)) = v173 + (v175 - v176)
							v180 = *(*int64)(unsafe.Add(mBase, uint32(v76)+112))
							v182 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[15]))
							v183 = *(*int64)(unsafe.Add(mBase, uint32(v81)+112))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+112)) = v180 + (v182 - v183)
							v187 = *(*int64)(unsafe.Add(mBase, uint32(v76)+120))
							v189 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[16]))
							v190 = *(*int64)(unsafe.Add(mBase, uint32(v81)+120))
							*(*int64)(unsafe.Add(mBase, uint32(v76)+120)) = v187 + (v189 - v190)
							v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
							if v195 != 0 {
								v196 = v76
							} else {
								v196 = int32(0)
							}
							v199 = v196
						} else {
							v199 = int32(0)
						}
						v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
						if v203 != 0 {
							v204 = v16 + int32(8)
						} else {
							v204 = int32(0)
						}
						F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v16+int32(280), v199, v204)
						mBase = m.M
						v206 = m.ExcPending
						if v206 != 0 {
							return
						} else {
							m.G0 = v16 + int32(288)
							return
						}
					}
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v72 == int32(1) {
						v76 = v16 + int32(24)
						base.MemoryFill(m, v76, int32(0), int32(128))
						v81 = v16 + int32(152)
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
						v84 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[1]))
						v85 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
						*(*int64)(unsafe.Add(mBase, uint32(v76))) = v82 + (v84 - v85)
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
						v91 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[2]))
						v92 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v89 + (v91 - v92)
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
						v98 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[3]))
						v99 = *(*int64)(unsafe.Add(mBase, uint32(v81)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+16)) = v96 + (v98 - v99)
						v103 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
						v105 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[4]))
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v81)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+24)) = v103 + (v105 - v106)
						v110 = *(*int64)(unsafe.Add(mBase, uint32(v76)+32))
						v112 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[5]))
						v113 = *(*int64)(unsafe.Add(mBase, uint32(v81)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+32)) = v110 + (v112 - v113)
						v117 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
						v119 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[6]))
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v81)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+40)) = v117 + (v119 - v120)
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v76)+48))
						v126 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[7]))
						v127 = *(*int64)(unsafe.Add(mBase, uint32(v81)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+48)) = v124 + (v126 - v127)
						v131 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
						v133 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[8]))
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v81)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+56)) = v131 + (v133 - v134)
						v138 = *(*int64)(unsafe.Add(mBase, uint32(v76)+64))
						v140 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[9]))
						v141 = *(*int64)(unsafe.Add(mBase, uint32(v81)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+64)) = v138 + (v140 - v141)
						v145 = *(*int64)(unsafe.Add(mBase, uint32(v76)+72))
						v147 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[10]))
						v148 = *(*int64)(unsafe.Add(mBase, uint32(v81)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+72)) = v145 + (v147 - v148)
						v152 = *(*int64)(unsafe.Add(mBase, uint32(v76)+80))
						v154 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[11]))
						v155 = *(*int64)(unsafe.Add(mBase, uint32(v81)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+80)) = v152 + (v154 - v155)
						v159 = *(*int64)(unsafe.Add(mBase, uint32(v76)+88))
						v161 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[12]))
						v162 = *(*int64)(unsafe.Add(mBase, uint32(v81)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+88)) = v159 + (v161 - v162)
						v166 = *(*int64)(unsafe.Add(mBase, uint32(v76)+96))
						v168 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[13]))
						v169 = *(*int64)(unsafe.Add(mBase, uint32(v81)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+96)) = v166 + (v168 - v169)
						v173 = *(*int64)(unsafe.Add(mBase, uint32(v76)+104))
						v175 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[14]))
						v176 = *(*int64)(unsafe.Add(mBase, uint32(v81)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+104)) = v173 + (v175 - v176)
						v180 = *(*int64)(unsafe.Add(mBase, uint32(v76)+112))
						v182 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[15]))
						v183 = *(*int64)(unsafe.Add(mBase, uint32(v81)+112))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+112)) = v180 + (v182 - v183)
						v187 = *(*int64)(unsafe.Add(mBase, uint32(v76)+120))
						v189 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[16]))
						v190 = *(*int64)(unsafe.Add(mBase, uint32(v81)+120))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+120)) = v187 + (v189 - v190)
						v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v195 != 0 {
							v196 = v76
						} else {
							v196 = int32(0)
						}
						v199 = v196
					} else {
						v199 = int32(0)
					}
					v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
					if v203 != 0 {
						v204 = v16 + int32(8)
					} else {
						v204 = int32(0)
					}
					F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v16+int32(280), v199, v204)
					mBase = m.M
					v206 = m.ExcPending
					if v206 != 0 {
						return
					} else {
						m.G0 = v16 + int32(288)
						return
					}
				}
			}
		}
	} else {
		v33 = v8
		v34 = v8
		v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
		if v35 == int32(1) {
			base.MemoryCopy(m, v16+int32(152), int32(_a_F_standard_ExplainOneQuery_4), int32(128))
		} else {
		}
		v45 = v16 + int32(24)
		F___clock_gettime(m, int32(1), v45)
		mBase = m.M
		v47 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
		v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+32)))
		v49 = F_pg_plan_query(m, l0, l4, l1, l5)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			v51 = int32(1)
			F___clock_gettime(m, v51, v45)
			mBase = m.M
			v53 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+32)))
			v55 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v16)+280)) = v53 - v48 + (v55-v47)*int64(1000000000)
			v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
			if v61 == v51 {
				*(*int32)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[0])) = v34
				F_MemoryContextMemConsumed(m, v33, v16+int32(8))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return
				} else {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v72 == int32(1) {
						v76 = v16 + int32(24)
						base.MemoryFill(m, v76, int32(0), int32(128))
						v81 = v16 + int32(152)
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
						v84 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[1]))
						v85 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
						*(*int64)(unsafe.Add(mBase, uint32(v76))) = v82 + (v84 - v85)
						v89 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
						v91 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[2]))
						v92 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v89 + (v91 - v92)
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
						v98 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[3]))
						v99 = *(*int64)(unsafe.Add(mBase, uint32(v81)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+16)) = v96 + (v98 - v99)
						v103 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
						v105 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[4]))
						v106 = *(*int64)(unsafe.Add(mBase, uint32(v81)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+24)) = v103 + (v105 - v106)
						v110 = *(*int64)(unsafe.Add(mBase, uint32(v76)+32))
						v112 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[5]))
						v113 = *(*int64)(unsafe.Add(mBase, uint32(v81)+32))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+32)) = v110 + (v112 - v113)
						v117 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
						v119 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[6]))
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v81)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+40)) = v117 + (v119 - v120)
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v76)+48))
						v126 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[7]))
						v127 = *(*int64)(unsafe.Add(mBase, uint32(v81)+48))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+48)) = v124 + (v126 - v127)
						v131 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
						v133 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[8]))
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v81)+56))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+56)) = v131 + (v133 - v134)
						v138 = *(*int64)(unsafe.Add(mBase, uint32(v76)+64))
						v140 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[9]))
						v141 = *(*int64)(unsafe.Add(mBase, uint32(v81)+64))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+64)) = v138 + (v140 - v141)
						v145 = *(*int64)(unsafe.Add(mBase, uint32(v76)+72))
						v147 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[10]))
						v148 = *(*int64)(unsafe.Add(mBase, uint32(v81)+72))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+72)) = v145 + (v147 - v148)
						v152 = *(*int64)(unsafe.Add(mBase, uint32(v76)+80))
						v154 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[11]))
						v155 = *(*int64)(unsafe.Add(mBase, uint32(v81)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+80)) = v152 + (v154 - v155)
						v159 = *(*int64)(unsafe.Add(mBase, uint32(v76)+88))
						v161 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[12]))
						v162 = *(*int64)(unsafe.Add(mBase, uint32(v81)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+88)) = v159 + (v161 - v162)
						v166 = *(*int64)(unsafe.Add(mBase, uint32(v76)+96))
						v168 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[13]))
						v169 = *(*int64)(unsafe.Add(mBase, uint32(v81)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+96)) = v166 + (v168 - v169)
						v173 = *(*int64)(unsafe.Add(mBase, uint32(v76)+104))
						v175 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[14]))
						v176 = *(*int64)(unsafe.Add(mBase, uint32(v81)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+104)) = v173 + (v175 - v176)
						v180 = *(*int64)(unsafe.Add(mBase, uint32(v76)+112))
						v182 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[15]))
						v183 = *(*int64)(unsafe.Add(mBase, uint32(v81)+112))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+112)) = v180 + (v182 - v183)
						v187 = *(*int64)(unsafe.Add(mBase, uint32(v76)+120))
						v189 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[16]))
						v190 = *(*int64)(unsafe.Add(mBase, uint32(v81)+120))
						*(*int64)(unsafe.Add(mBase, uint32(v76)+120)) = v187 + (v189 - v190)
						v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
						if v195 != 0 {
							v196 = v76
						} else {
							v196 = int32(0)
						}
						v199 = v196
					} else {
						v199 = int32(0)
					}
					v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
					if v203 != 0 {
						v204 = v16 + int32(8)
					} else {
						v204 = int32(0)
					}
					F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v16+int32(280), v199, v204)
					mBase = m.M
					v206 = m.ExcPending
					if v206 != 0 {
						return
					} else {
						m.G0 = v16 + int32(288)
						return
					}
				}
			} else {
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
				if v72 == int32(1) {
					v76 = v16 + int32(24)
					base.MemoryFill(m, v76, int32(0), int32(128))
					v81 = v16 + int32(152)
					v82 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
					v84 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[1]))
					v85 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
					*(*int64)(unsafe.Add(mBase, uint32(v76))) = v82 + (v84 - v85)
					v89 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
					v91 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[2]))
					v92 = *(*int64)(unsafe.Add(mBase, uint32(v81)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+8)) = v89 + (v91 - v92)
					v96 = *(*int64)(unsafe.Add(mBase, uint32(v76)+16))
					v98 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[3]))
					v99 = *(*int64)(unsafe.Add(mBase, uint32(v81)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+16)) = v96 + (v98 - v99)
					v103 = *(*int64)(unsafe.Add(mBase, uint32(v76)+24))
					v105 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[4]))
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v81)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+24)) = v103 + (v105 - v106)
					v110 = *(*int64)(unsafe.Add(mBase, uint32(v76)+32))
					v112 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[5]))
					v113 = *(*int64)(unsafe.Add(mBase, uint32(v81)+32))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+32)) = v110 + (v112 - v113)
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v76)+40))
					v119 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[6]))
					v120 = *(*int64)(unsafe.Add(mBase, uint32(v81)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+40)) = v117 + (v119 - v120)
					v124 = *(*int64)(unsafe.Add(mBase, uint32(v76)+48))
					v126 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[7]))
					v127 = *(*int64)(unsafe.Add(mBase, uint32(v81)+48))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+48)) = v124 + (v126 - v127)
					v131 = *(*int64)(unsafe.Add(mBase, uint32(v76)+56))
					v133 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[8]))
					v134 = *(*int64)(unsafe.Add(mBase, uint32(v81)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+56)) = v131 + (v133 - v134)
					v138 = *(*int64)(unsafe.Add(mBase, uint32(v76)+64))
					v140 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[9]))
					v141 = *(*int64)(unsafe.Add(mBase, uint32(v81)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+64)) = v138 + (v140 - v141)
					v145 = *(*int64)(unsafe.Add(mBase, uint32(v76)+72))
					v147 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[10]))
					v148 = *(*int64)(unsafe.Add(mBase, uint32(v81)+72))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+72)) = v145 + (v147 - v148)
					v152 = *(*int64)(unsafe.Add(mBase, uint32(v76)+80))
					v154 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[11]))
					v155 = *(*int64)(unsafe.Add(mBase, uint32(v81)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+80)) = v152 + (v154 - v155)
					v159 = *(*int64)(unsafe.Add(mBase, uint32(v76)+88))
					v161 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[12]))
					v162 = *(*int64)(unsafe.Add(mBase, uint32(v81)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+88)) = v159 + (v161 - v162)
					v166 = *(*int64)(unsafe.Add(mBase, uint32(v76)+96))
					v168 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[13]))
					v169 = *(*int64)(unsafe.Add(mBase, uint32(v81)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+96)) = v166 + (v168 - v169)
					v173 = *(*int64)(unsafe.Add(mBase, uint32(v76)+104))
					v175 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[14]))
					v176 = *(*int64)(unsafe.Add(mBase, uint32(v81)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+104)) = v173 + (v175 - v176)
					v180 = *(*int64)(unsafe.Add(mBase, uint32(v76)+112))
					v182 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[15]))
					v183 = *(*int64)(unsafe.Add(mBase, uint32(v81)+112))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+112)) = v180 + (v182 - v183)
					v187 = *(*int64)(unsafe.Add(mBase, uint32(v76)+120))
					v189 = *(*int64)(unsafe.Add(mBase, _c_F_standard_ExplainOneQuery[16]))
					v190 = *(*int64)(unsafe.Add(mBase, uint32(v81)+120))
					*(*int64)(unsafe.Add(mBase, uint32(v76)+120)) = v187 + (v189 - v190)
					v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+7)))
					if v195 != 0 {
						v196 = v76
					} else {
						v196 = int32(0)
					}
					v199 = v196
				} else {
					v199 = int32(0)
				}
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+11)))
				if v203 != 0 {
					v204 = v16 + int32(8)
				} else {
					v204 = int32(0)
				}
				F_ExplainOnePlan(m, v49, l2, l3, l4, l5, l6, v16+int32(280), v199, v204)
				mBase = m.M
				v206 = m.ExcPending
				if v206 != 0 {
					return
				} else {
					m.G0 = v16 + int32(288)
					return
				}
			}
		}
	}
}
