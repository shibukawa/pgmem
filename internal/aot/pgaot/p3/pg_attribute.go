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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
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
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
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
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v285 int32
	_ = v285
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v320 int32
	_ = v320
	v6 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = int32(655)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v17) <= base.Ui32(v18) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = v17
	goto L3
L2:
	;
	v21 = v18
	goto L3
L3:
	;
	v24 = F_palloc(m, v21<<(uint(int32(2))%32))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = v6
	goto L9
L7:
	;
	goto L8
L8:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v66 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v45 = F_MakeTupleTableSlot(m, v16, int32(_a_F_InsertPgAttributeTuples_0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24+v31<<(uint(int32(2))%32)))) = v45
	v49 = v31 + int32(1)
	if v49 != v21 {
		v31 = v49
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if v18 != 0 {
		goto L42
	} else {
		goto L43
	}
L14:
	;
	v73 = l4
	v75 = v66
	v77 = v6
	v78 = v6
	v83 = v6
	goto L15
L15:
	;
	v86 = v24 + v77<<(uint(int32(2))%32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	m.T0[v89].(func(*base.Module, int32))(m, v87)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	if v256 == int32(0) {
		goto L13
	} else {
		goto L40
	}
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)+12))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+20))
	base.MemoryFill(m, v95, int32(0), v94)
	goto L20
L19:
	;
	goto L20
L20:
	;
	v103 = l1 + v75<<(uint(int32(4))%32) + v78*int32(100)
	v105 = v103 + int32(20)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+16))
	if l2 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v109 = l2
	goto L23
L22:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v109 = v108
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v103 + int32(24)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+16))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v105)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v117)+8)) = v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+16))
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+12)) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	v126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+16)) = v126
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v105)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+20)) = v130
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+16))
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+80)))
	*(*int32)(unsafe.Add(mBase, uint32(v133)+24)) = v134
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)+16))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+82)))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+28)) = v138
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
	v142 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105)+83)))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+32)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+16))
	v146 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105)+84)))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+36)) = v146
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+16))
	v150 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105)+85)))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+40)) = v150
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+86)))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+44)) = v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+87)))
	*(*int32)(unsafe.Add(mBase, uint32(v157)+48)) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+88)))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+52)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+16))
	v166 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105)+89)))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+56)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	v170 = int32(*(*int8)(unsafe.Add(mBase, uint32(v105)+90)))
	*(*int32)(unsafe.Add(mBase, uint32(v169)+60)) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+91)))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+64)) = v174
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+92)))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+68)) = v178
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	v182 = int32(*(*int16)(unsafe.Add(mBase, uint32(v105)+94)))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+72)) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v105)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+76)) = v186
	if l3 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v212)+22)) = uint8(v210)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+20))
	v216 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+21)) = uint8(v216)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+23)) = uint8(v216)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+24)) = uint8(v216)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+4)))
	v229 = v227 & int32(_a_F_InsertPgAttributeTuples_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v226)+4)) = uint16(v229)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	*(*uint16)(unsafe.Add(mBase, uint32(v226)+6)) = uint16(v232)
	goto L28
L25:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	v192 = l3 + v78<<(uint(int32(4))%32)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+80)) = v193
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+20))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v196)+20)) = uint8(v197)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+16))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+88)) = v201
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+12)))
	v210 = v203
	goto L24
L26:
	;
	goto L27
L27:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v205)+20)) = uint8(v206)
	v210 = v206
	goto L24
L28:
	;
	v235 = v77 + int32(1)
	if v21 != v235 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v258 = v78 + int32(1)
	if v258 < v254 {
		v73 = v253
		v75 = v254
		v77 = v255
		v78 = v258
		v83 = v256
		goto L15
	} else {
		goto L39
	}
L30:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v78 != v237-int32(1) {
		v253 = v73
		v254 = v237
		v255 = v235
		v256 = v83
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v73 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L32
L34:
	;
	v245 = F_CatalogOpenIndexes(m, l0)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v247 = v73
	v248 = v83
	goto L36
L36:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, l0, v24, v235, v247)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L38
	}
L37:
	;
	v247 = v245
	v248 = int32(1)
	goto L36
L38:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v253 = v247
	v254 = v251
	v255 = int32(0)
	v256 = v248
	goto L29
L39:
	;
	goto L16
L40:
	;
	F_CatalogCloseIndexes(m, v253)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	goto L13
L42:
	;
	v285 = int32(0)
	goto L45
L43:
	;
	goto L44
L44:
	;
	F_pfree(m, v24)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L49
	}
L45:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v24+v285<<(uint(int32(2))%32))))
	F_ExecDropSingleTupleTableSlot(m, v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L44
L47:
	;
	v302 = v285 + int32(1)
	if v302 != v21 {
		v285 = v302
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
