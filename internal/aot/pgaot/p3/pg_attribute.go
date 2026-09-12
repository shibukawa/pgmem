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
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
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
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
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
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
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
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
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
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v289 int32
	_ = v289
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v324 int32
	_ = v324
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
	v45 = F_MakeSingleTupleTableSlot(m, v16, int32(1599208))
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
		goto L43
	} else {
		goto L44
	}
L14:
	;
	v75 = l4
	v77 = v66
	v80 = v6
	v81 = v6
	v84 = v6
	goto L15
L15:
	;
	v88 = v24 + v81<<(uint(int32(2))%32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	m.T0[v91].(func(*base.Module, int32))(m, v89)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	if v260 == int32(0) {
		goto L13
	} else {
		goto L41
	}
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v100 = F__emscripten_memset_bulkmem(m, v95, base.I32_extend8_s(int32(0)), v98)
	mBase = m.M
	goto L18
L18:
	;
	v101 = int32(4)
	v106 = l1 + int32(20) + v77<<(uint(v101)%32) + v80*int32(100)
	if l3 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v111 = l3 + v80<<(uint(v101)%32)
	goto L21
L20:
	;
	v111 = int32(0)
	goto L21
L21:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+16))
	if l2 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v115 = l2
	goto L24
L23:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v115 = v114
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v106 + int32(4)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+16))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v106)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+8)) = v124
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106)+72)))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+12)) = v128
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106)+74)))
	*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = v132
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v106)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+20)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106)+80)))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = v140
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)+16))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+82)))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+28)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+16))
	v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+83)))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+32)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+84)))
	*(*int32)(unsafe.Add(mBase, uint32(v151)+36)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	v156 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+85)))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+40)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+86)))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+44)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+16))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+87)))
	*(*int32)(unsafe.Add(mBase, uint32(v163)+48)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+16))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+88)))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+52)) = v168
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+16))
	v172 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+89)))
	*(*int32)(unsafe.Add(mBase, uint32(v171)+56)) = v172
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	v176 = int32(*(*int8)(unsafe.Add(mBase, uint32(v106)+90)))
	*(*int32)(unsafe.Add(mBase, uint32(v175)+60)) = v176
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+91)))
	*(*int32)(unsafe.Add(mBase, uint32(v179-int32(-64)))) = v182
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+16))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+92)))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+68)) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+16))
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106)+94)))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+72)) = v190
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)+16))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v106)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+76)) = v194
	if v111 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v216)+22)) = uint8(v214)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+20))
	v220 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+21)) = uint8(v220)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v223)+23)) = uint8(v220)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v227)+24)) = uint8(v220)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)))
	v233 = v231 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)) = uint16(v233)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)))
	*(*uint16)(unsafe.Add(mBase, uint32(v230)+6)) = uint16(v236)
	goto L29
L26:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+16))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	*(*int32)(unsafe.Add(mBase, uint32(v197)+80)) = v198
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)+20))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v201)+20)) = uint8(v202)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+88)) = v206
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+12)))
	v214 = v208
	goto L25
L27:
	;
	goto L28
L28:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
	v211 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v210)+20)) = uint8(v211)
	v214 = v211
	goto L25
L29:
	;
	v239 = v81 + int32(1)
	if v21 != v239 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v262 = v80 + int32(1)
	if v262 < v258 {
		v75 = v257
		v77 = v258
		v80 = v262
		v81 = v259
		v84 = v260
		goto L15
	} else {
		goto L40
	}
L31:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v80 != v241-int32(1) {
		v257 = v75
		v258 = v241
		v259 = v239
		v260 = v84
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v75 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v249 = F_CatalogOpenIndexes(m, l0)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L38
	}
L36:
	;
	v251 = v75
	v252 = v84
	goto L37
L37:
	;
	F_CatalogTuplesMultiInsertWithInfo(m, l0, v24, v239, v251)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L4
	} else {
		goto L39
	}
L38:
	;
	v251 = v249
	v252 = int32(1)
	goto L37
L39:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v257 = v251
	v258 = v255
	v259 = int32(0)
	v260 = v252
	goto L30
L40:
	;
	goto L16
L41:
	;
	F_CatalogCloseIndexes(m, v257)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L13
L43:
	;
	v289 = int32(0)
	goto L46
L44:
	;
	goto L45
L45:
	;
	F_pfree(m, v24)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L4
	} else {
		goto L50
	}
L46:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v24+v289<<(uint(int32(2))%32))))
	F_ExecDropSingleTupleTableSlot(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	v306 = v289 + int32(1)
	if v306 != v21 {
		v289 = v306
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
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
