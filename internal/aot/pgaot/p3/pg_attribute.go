package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InsertPgAttributeTuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int64
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int64
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int64
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
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
	var v163 int64
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v286 int32
	_ = v286
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v321 int32
	_ = v321
	v6 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = int32(655)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v18) <= base.Ui32(v19) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = v18
	goto L3
L2:
	;
	v22 = v19
	goto L3
L3:
	;
	v23 = F_palloc_mul(m, int32(4), v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v30 = v6
	goto L9
L7:
	;
	goto L8
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v65 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v44 = F_MakeSingleTupleTableSlot(m, v16, int32(_a_F_InsertPgAttributeTuples_0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23+v30<<(uint(int32(2))%32)))) = v44
	v48 = v30 + int32(1)
	if v48 != v22 {
		v30 = v48
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v19 != 0 {
		goto L42
	} else {
		goto L43
	}
L14:
	;
	v72 = l4
	v74 = v65
	v76 = v6
	v77 = v6
	v82 = v6
	goto L15
L15:
	;
	v85 = v23 + v76<<(uint(int32(2))%32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	m.T0[v88].(func(*base.Module, int32))(m, v86)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	if v257 == int32(0) {
		goto L13
	} else {
		goto L40
	}
L17:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	if v93 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v91)+20))
	base.MemoryFill(m, v94, int32(0), v93)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v102 = l1 + v74<<(uint(int32(3))%32) + v77*int32(100)
	v104 = v102 + int32(28)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	if l2 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v108 = l2
	goto L23
L22:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	v108 = v107
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106))) = base.I64_extend_i32_u(v108)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v112)+8)) = base.I64_extend_i32_u(v102 + int32(32))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	v119 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v104)+68)))
	*(*int64)(unsafe.Add(mBase, uint32(v118)+16)) = v119
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v123 = int64(*(*int16)(unsafe.Add(mBase, uint32(v104)+72)))
	*(*int64)(unsafe.Add(mBase, uint32(v122)+24)) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
	v127 = int64(*(*int16)(unsafe.Add(mBase, uint32(v104)+74)))
	*(*int64)(unsafe.Add(mBase, uint32(v126)+32)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+16))
	v131 = int64(*(*int32)(unsafe.Add(mBase, uint32(v104)+76)))
	*(*int64)(unsafe.Add(mBase, uint32(v130)+40)) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v135 = int64(*(*int16)(unsafe.Add(mBase, uint32(v104)+80)))
	*(*int64)(unsafe.Add(mBase, uint32(v134)+48)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
	v139 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104)+82)))
	*(*int64)(unsafe.Add(mBase, uint32(v138)+56)) = v139
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	v143 = int64(*(*int8)(unsafe.Add(mBase, uint32(v104)+83)))
	*(*int64)(unsafe.Add(mBase, uint32(v142)+64)) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
	v147 = int64(*(*int8)(unsafe.Add(mBase, uint32(v104)+84)))
	*(*int64)(unsafe.Add(mBase, uint32(v146)+72)) = v147
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+16))
	v151 = int64(*(*int8)(unsafe.Add(mBase, uint32(v104)+85)))
	*(*int64)(unsafe.Add(mBase, uint32(v150)+80)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
	v155 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104)+86)))
	*(*int64)(unsafe.Add(mBase, uint32(v154)+88)) = v155
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v159 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104)+87)))
	*(*int64)(unsafe.Add(mBase, uint32(v158)+96)) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	v163 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104)+88)))
	*(*int64)(unsafe.Add(mBase, uint32(v162)+104)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v167 = int64(*(*int8)(unsafe.Add(mBase, uint32(v104)+89)))
	*(*int64)(unsafe.Add(mBase, uint32(v166)+112)) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+16))
	v171 = int64(*(*int8)(unsafe.Add(mBase, uint32(v104)+90)))
	*(*int64)(unsafe.Add(mBase, uint32(v170)+120)) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	v175 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104)+91)))
	*(*int64)(unsafe.Add(mBase, uint32(v174)+128)) = v175
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+16))
	v179 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v104)+92)))
	*(*int64)(unsafe.Add(mBase, uint32(v178)+136)) = v179
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+16))
	v183 = int64(*(*int16)(unsafe.Add(mBase, uint32(v104)+94)))
	*(*int64)(unsafe.Add(mBase, uint32(v182)+144)) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	v187 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v104)+96)))
	*(*int64)(unsafe.Add(mBase, uint32(v186)+152)) = v187
	if l3 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v213)+22)) = uint8(v211)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	v217 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+21)) = uint8(v217)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v219)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v220)+23)) = uint8(v217)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+24)) = uint8(v217)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227)+4)))
	v230 = v228 & int32(_a_F_InsertPgAttributeTuples_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v227)+4)) = uint16(v230)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	*(*uint16)(unsafe.Add(mBase, uint32(v227)+6)) = uint16(v233)
	goto L28
L25:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+16))
	v193 = l3 + v77<<(uint(int32(5))%32)
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v193)))
	*(*int64)(unsafe.Add(mBase, uint32(v190)+160)) = v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+20))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v197)+20)) = uint8(v198)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+16))
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v193)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v201)+176)) = v202
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+24)))
	v211 = v204
	goto L24
L26:
	;
	goto L27
L27:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+20))
	v207 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+20)) = uint8(v207)
	v211 = v207
	goto L24
L28:
	;
	v236 = v76 + int32(1)
	if v22 != v236 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v259 = v77 + int32(1)
	if v259 < v255 {
		v72 = v254
		v74 = v255
		v76 = v256
		v77 = v259
		v82 = v257
		goto L15
	} else {
		goto L39
	}
L30:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v77 != v238-int32(1) {
		v254 = v72
		v255 = v238
		v256 = v236
		v257 = v82
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v72 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v246 = F_CatalogOpenIndexes(m, l0)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v248 = v72
	v249 = v82
	goto L36
L36:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, l0, v23, v236, v248)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L4
	} else {
		goto L38
	}
L37:
	;
	v248 = v246
	v249 = int32(1)
	goto L36
L38:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v254 = v248
	v255 = v252
	v256 = int32(0)
	v257 = v249
	goto L29
L39:
	;
	goto L16
L40:
	;
	F_CatalogCloseIndexes(m, v254)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	goto L13
L42:
	;
	v286 = int32(0)
	goto L45
L43:
	;
	goto L44
L44:
	;
	F_pfree(m, v23)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L4
	} else {
		goto L49
	}
L45:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v23+v286<<(uint(int32(2))%32))))
	F_ExecDropSingleTupleTableSlot(m, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	v303 = v286 + int32(1)
	if v303 != v22 {
		v286 = v303
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	return
}
func F_pg_attribute_aclcheck_all(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_pg_attribute_aclcheck_all_ext(m, l0, l1, l2, l3, int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_pg_attribute_aclcheck_ext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_pg_attribute_aclmask_ext(m, l0, l1, l2, l3, l4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
