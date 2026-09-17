package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_entryBeginPlaceToPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v364 int32
	_ = v364
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	v5 = l4
	v9 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(_a_F_entryBeginPlaceToPage_0)
	m.G0 = v20
	if l1 < v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v40 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	if v41 == v40 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_entryBeginPlaceToPage[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v25+(l1^int32(-1))<<(uint(int32(2))%32))))
	v39 = v31
	goto L1
L3:
	;
	goto L4
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_entryBeginPlaceToPage[1]))
	v39 = v33 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39+v44<<(uint(int32(2))%32))+20))
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39+v48&int32(_a_F_entryBeginPlaceToPage_1))+6)))
	v61 = (v52&int32(_a_F_entryBeginPlaceToPage_2)+int32(7))&int32(_a_F_entryBeginPlaceToPage_3) | int32(4)
	goto L7
L6:
	;
	v61 = v9
	goto L7
L7:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
	v64 = int32(4)
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+14)))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+12)))
	v67 = v65 - v66
	if v67 <= v64 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L20
	} else {
		goto L70
	}
L9:
	;
	if base.Ui32(v70-int32(4)+v61) < base.Ui32((v63&int32(_a_F_entryBeginPlaceToPage_2)+int32(7))&int32(_a_F_entryBeginPlaceToPage_3)|int32(4)) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v70 = v64
	goto L12
L11:
	;
	v70 = v67
	goto L12
L12:
	;
	goto L9
L13:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	if l1 < int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v364 = v40
	goto L15
L15:
	;
	m.G0 = v20 + int32(_a_F_entryBeginPlaceToPage_0)
	return v364
L16:
	;
	v119 = F_PageGetTempPageCopy(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L20
	} else {
		goto L23
	}
L17:
	;
	v89 = (l1 ^ int32(-1)) << (uint(int32(2)) % 32)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_entryBeginPlaceToPage[0]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v91)))
	v94 = F_PageGetTempPageCopy(m, v93)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v103 = l1 << (uint(int32(13)) % 32)
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_entryBeginPlaceToPage[1]))
	v109 = F_PageGetTempPageCopy(m, v103+v105+int32(-8192))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L22
	}
L20:
	;
	return int32(0)
L21:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_entryBeginPlaceToPage[0]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99+v89)))
	v117 = v94
	v118 = v101
	goto L16
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_entryBeginPlaceToPage[1]))
	v117 = v109
	v118 = v112 + v103 + int32(-8192)
	goto L16
L23:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+19)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)))
	if v122 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_PageIndexTupleDelete(m, v117, v83)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L20
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v5 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v148 = int32(0)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v149) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+v129)+6)))
	if v131&int32(2) != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v117+v83<<(uint(int32(2))%32))+20))
	v140 = v117 + v137&int32(_a_F_entryBeginPlaceToPage_1)
	v141 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)) = uint16(v141)
	*(*uint16)(unsafe.Add(mBase, uint32(v140)+2)) = uint16(v5)
	v145 = int32(base.Ui32(v5) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v145)
	goto L28
L31:
	;
	if v83 == v157&int32(_a_F_entryBeginPlaceToPage_4)+int32(1) {
		goto L53
	} else {
		goto L54
	}
L32:
	;
	v157 = int32(base.Ui32(v149+int32(_a_F_entryBeginPlaceToPage_5)) >> (uint(int32(2)) % 32))
	goto L34
L33:
	;
	v157 = v148
	goto L34
L34:
	;
	if v157&int32(_a_F_entryBeginPlaceToPage_4) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v236 = v20 + int32(16)
	v242 = v148
	goto L31
L36:
	;
	goto L37
L37:
	;
	v164 = int32(2)
	v168 = (v157 + int32(1)) & int32(_a_F_entryBeginPlaceToPage_4)
	if base.Ui32(v168) <= base.Ui32(v164) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v171 = v164
	goto L40
L39:
	;
	v171 = v168
	goto L40
L40:
	;
	v178 = int32(1)
	v179 = v20 + int32(16)
	v185 = v148
	goto L41
L41:
	;
	if v178 == v83 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v236 = v227
	v242 = v230
	goto L31
L43:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v195)+6)))
	v202 = (v196&int32(_a_F_entryBeginPlaceToPage_2) + int32(7)) & int32(_a_F_entryBeginPlaceToPage_3)
	if v202 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v208 = v179
	v209 = v185
	goto L45
L45:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v117+int32(20)+v178<<(uint(int32(2))%32))))
	v218 = v117 + v215&int32(_a_F_entryBeginPlaceToPage_1)
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v218)+6)))
	v225 = (v219&int32(_a_F_entryBeginPlaceToPage_2) + int32(7)) & int32(_a_F_entryBeginPlaceToPage_3)
	if v225 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	base.MemoryCopy(m, v179, v195, v202)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v208 = v179 + v202
	v209 = v185 + v202 + int32(4)
	goto L45
L49:
	;
	base.MemoryCopy(m, v208, v218, v225)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v227 = v208 + v225
	v230 = v209 + v225 + int32(4)
	v232 = v178 + int32(1)
	if v232 != v171 {
		v178 = v232
		v179 = v227
		v185 = v230
		goto L41
	} else {
		goto L52
	}
L52:
	;
	goto L42
L53:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+6)))
	v263 = (v257&int32(_a_F_entryBeginPlaceToPage_2) + int32(7)) & int32(_a_F_entryBeginPlaceToPage_3)
	if v263 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v270 = v242
	goto L55
L55:
	;
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)))
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117+v271)+6)))
	v274 = int32(8)
	v275 = v121 << (uint(v274) % 32)
	F_PageInit(m, v119, v275, v274)
	mBase = m.M
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+16)))
	v279 = v119 + v278
	*(*int32)(unsafe.Add(mBase, uint32(v279))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v279)+6)) = uint16(v273)
	goto L59
L56:
	;
	base.MemoryCopy(m, v236, v256, v263)
	goto L58
L57:
	;
	goto L58
L58:
	;
	v270 = v263 + v242 + int32(4)
	goto L55
L59:
	;
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119)+16)))
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119+v283)+6)))
	F_PageInit(m, v117, v275, int32(8))
	mBase = m.M
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+16)))
	v289 = v117 + v288
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = int32(-1)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+6)) = uint16(v285)
	goto L60
L60:
	;
	v293 = int32(1)
	v304 = v20 + int32(16)
	v305 = v117
	v311 = int32(0)
	v312 = v293
	goto L61
L61:
	;
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v304)+6)))
	v322 = v320 & int32(_a_F_entryBeginPlaceToPage_2)
	if base.Ui32(int32(base.Ui32(v270)>>(uint(v293)%32))) < base.Ui32(v311) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v117
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v119
	v364 = int32(2)
	goto L15
L63:
	;
	v333 = int32(0)
	v335 = F_PageAddItemExtended(m, v331, v304, v322, v333, v333)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L20
	} else {
		goto L67
	}
L64:
	;
	v331 = v119
	v332 = v311
	goto L63
L65:
	;
	goto L66
L66:
	;
	v331 = v305
	v332 = v311 + (v322+int32(7))&int32(_a_F_entryBeginPlaceToPage_3) + int32(4)
	goto L63
L67:
	;
	if v335 == int32(0) {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v304)+6)))
	v348 = v312 + int32(1)
	if base.Ui32(v348&int32(_a_F_entryBeginPlaceToPage_4)) <= base.Ui32((v157+v293)&int32(_a_F_entryBeginPlaceToPage_4)) {
		v304 = v304 + (v339&int32(_a_F_entryBeginPlaceToPage_2)+int32(7))&int32(_a_F_entryBeginPlaceToPage_3)
		v305 = v331
		v311 = v332
		v312 = v348
		goto L61
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v381 + int32(4)
	F_errmsg_internal(m, int32(_a_F_entryBeginPlaceToPage_6), v20)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_entryBeginPlaceToPage_7), int32(689), int32(_a_F_entryBeginPlaceToPage_8))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_entry_alloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
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
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 float64
	_ = v220
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v228 float64
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 float64
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int64
	_ = v341
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v448 int32
	_ = v448
	var v449 float64
	_ = v449
	var v450 float64
	_ = v450
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v465 int64
	_ = v465
	var v466 int64
	_ = v466
	var v474 int64
	_ = v474
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+412))
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[1]))
	if v89 <= v87 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+376))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+364))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+352))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+340))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+328))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+316))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+304))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+292))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+280))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+268))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+256))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v22)+244))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+232))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)+220))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+208))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v22)+196))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v22)+184))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v22)+172))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+160))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v22)+148))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v22)+136))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+124))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v22)+112))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v22)+100))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v22)+88))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	v87 = v25 + (v26 + (v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + v23))))))))))))))))))))))))))))))
	goto L4
L3:
	;
	v87 = v23
	goto L4
L4:
	;
	goto L1
L5:
	;
	goto L8
L6:
	;
	goto L7
L7:
	;
	v432 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v436 = F_hash_search(m, v432, l0, int32(1), v17+int32(12))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L14
	} else {
		goto L58
	}
L8:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+412))
	if v110 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v176 = F_palloc(m, v173<<(uint(int32(2))%32))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+376))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+364))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+352))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)+340))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)+328))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+316))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v108)+304))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v108)+292))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)+280))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v108)+268))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v108)+256))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v108)+244))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v108)+232))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v108)+220))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v108)+208))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v108)+196))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v108)+184))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v108)+172))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v108)+160))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v108)+148))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v108)+136))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v108)+124))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v108)+112))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v108)+100))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v108)+88))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v108)+76))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v108)+64))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v108)+52))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v108)+40))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v108)+28))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v173 = v111 + (v112 + (v113 + (v114 + (v115 + (v116 + (v117 + (v118 + (v119 + (v120 + (v121 + (v122 + (v123 + (v124 + (v125 + (v126 + (v127 + (v128 + (v129 + (v130 + (v131 + (v132 + (v133 + (v134 + (v135 + (v136 + (v137 + (v138 + (v139 + (v140 + (v141 + v109))))))))))))))))))))))))))))))
	goto L13
L12:
	;
	v173 = v109
	goto L13
L13:
	;
	goto L10
L14:
	;
	return int32(0)
L15:
	;
	v181 = v17 + int32(12)
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	F_hash_seq_init(m, v181, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v186 = int32(0)
	v189 = F_hash_seq_search(m, v181)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	F_pfree(m, v176)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L14
	} else {
		goto L48
	}
L18:
	;
	if v189 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pg_qsort(m, v176, int32(0), int32(4), int32(_a_F_entry_alloc_0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L14
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v207 = v189
	v208 = v186
	v212 = v186
	v213 = v186
	goto L23
L22:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+16)) = int32(1024)
	goto L17
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176+v208<<(uint(int32(2))%32)))) = v207
	v220 = *(*float64)(unsafe.Add(mBase, uint32(v207)+256))
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v207)+24))
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v207)+32))
	if v223 == int64(0)-v225 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_pg_qsort(m, v176, v232, int32(4), int32(_a_F_entry_alloc_0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L14
	} else {
		goto L33
	}
L25:
	;
	v228 = float64(0.5)
	goto L27
L26:
	;
	v228 = float64(0.99)
	goto L27
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v207)+256)) = base.F64_mul(v220, v228)
	v232 = v208 + int32(1)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v207)+396))
	v234 = int32(-1)
	v238 = v212 + int32(base.Ui32(v233^v234)>>(uint(int32(31))%32))
	if v233 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v242 = v234
	goto L30
L29:
	;
	v242 = v233
	goto L30
L30:
	;
	v245 = v213 + v242 + int32(1)
	v248 = F_hash_seq_search(m, v17+int32(12))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	if v248 != 0 {
		v207 = v248
		v208 = v232
		v212 = v238
		v213 = v245
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L24
L33:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v176+v232<<(uint(int32(1))%32)&int32(-4))))
	v262 = *(*float64)(unsafe.Add(mBase, uint32(v261)+256))
	*(*float64)(unsafe.Add(mBase, uint32(v255)+8)) = v262
	if v238 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v264 = base.I32_div_u_s(v245, v238)
	v266 = v264
	goto L36
L35:
	;
	v266 = int32(1024)
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+16)) = v266
	v269 = int32(10)
	if base.Ui32(v269) <= base.Ui32(v232) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v272 = v269
	goto L39
L38:
	;
	v272 = v232
	goto L39
L39:
	;
	v274 = base.I32_div_u_s(v232, int32(20))
	if base.Ui32(v208) < base.Ui32(int32(199)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v277 = v272
	goto L42
L41:
	;
	v277 = v274
	goto L42
L42:
	;
	if v277 == int32(0) {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v285 = int32(0)
	goto L44
L44:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v296 = int32(2)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v176+v285<<(uint(v296)%32))))
	v302 = F_hash_search(m, v295, v299, v296, int32(0))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L14
	} else {
		goto L46
	}
L45:
	;
	goto L17
L46:
	;
	v305 = v285 + int32(1)
	if v305 != v277 {
		v285 = v305
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v324 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v324)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v324)+20)) = int32(1)
	if v325 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	F_s_lock(m, v329+int32(20), int32(_a_F_entry_alloc_1), int32(2205), int32(_a_F_entry_alloc_2))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L14
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v338)+20)) = int32(0)
	v341 = *(*int64)(unsafe.Add(mBase, uint32(v338)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v338)+40)) = v341 + int64(1)
	v346 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v348)+412))
	if v350 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L51
L53:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[1]))
	if v415 <= v413 {
		goto L8
	} else {
		goto L57
	}
L54:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v348)+376))
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348)+364))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v348)+352))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v348)+340))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v348)+328))
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v348)+316))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v348)+304))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v348)+292))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v348)+280))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v348)+268))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v348)+256))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v348)+244))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v348)+232))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v348)+220))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v348)+208))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v348)+196))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v348)+184))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v348)+172))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v348)+160))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v348)+148))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v348)+136))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v348)+124))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v348)+112))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v348)+100))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v348)+88))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v348)+76))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v348)+64))
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v348)+52))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v348)+40))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v348)+28))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v348)+16))
	v413 = v351 + (v352 + (v353 + (v354 + (v355 + (v356 + (v357 + (v358 + (v359 + (v360 + (v361 + (v362 + (v363 + (v364 + (v365 + (v366 + (v367 + (v368 + (v369 + (v370 + (v371 + (v372 + (v373 + (v374 + (v375 + (v376 + (v377 + (v378 + (v379 + (v380 + (v381 + v349))))))))))))))))))))))))))))))
	goto L56
L55:
	;
	v413 = v349
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L9
L58:
	;
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)))
	if v438 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryFill(m, v436+int32(24), int32(0), int32(368))
	if l4 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	m.G0 = v17 + int32(32)
	return v436
L62:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	v449 = *(*float64)(unsafe.Add(mBase, uint32(v448)+8))
	v450 = v449
	goto L64
L63:
	;
	v450 = float64(1)
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v436)+424)) = int32(0)
	*(*float64)(unsafe.Add(mBase, uint32(v436)+256)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v436)+400)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v436)+396)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v436)+392)) = l1
	v460 = m.G0
	v461 = int32(16)
	v462 = v460 - v461
	m.G0 = v462
	F_gettimeofday(m, v462)
	mBase = m.M
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v462)))
	v466 = int64(*(*int32)(unsafe.Add(mBase, uint32(v462)+8)))
	m.G0 = v462 + v461
	v474 = v466 + v465*int64(1000000) - int64(946684800000000)
	goto L65
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v436)+416)) = v474
	*(*int64)(unsafe.Add(mBase, uint32(v436)+408)) = v474
	goto L61
}
func F_entry_reset(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int64
	_ = v169
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int64
	_ = v219
	var v221 int64
	_ = v221
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int64
	_ = v254
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	if v15 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L99
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = F_LWLockAcquire(m, v22, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+412))
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v100 = m.G0
	v101 = int32(16)
	v102 = v100 - v101
	m.G0 = v102
	F_gettimeofday(m, v102)
	mBase = m.M
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
	v106 = int64(*(*int32)(unsafe.Add(mBase, uint32(v102)+8)))
	m.G0 = v102 + v101
	v114 = v106 + v105*int64(1000000) - int64(946684800000000)
	goto L10
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+376))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+364))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+352))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+340))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+328))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+316))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+304))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+292))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+280))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v31)+268))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v31)+256))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+244))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)+232))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+220))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v31)+208))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v31)+196))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)+184))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v31)+172))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v31)+160))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v31)+148))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v31)+136))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+124))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v31)+112))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+100))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v31)+88))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v31)+64))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v31)+52))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v96 = v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + (v57 + (v58 + (v59 + (v60 + (v61 + (v62 + (v63 + (v64 + v32))))))))))))))))))))))))))))))
	goto L9
L8:
	;
	v96 = v32
	goto L9
L9:
	;
	goto L6
L10:
	;
	v117 = int32(0)
	if base.B2i32(l2 == int64(0))|(base.B2i32(l0 == v117)|base.B2i32(l1 == v117)) == v117 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	if v282 == v96 {
		goto L65
	} else {
		goto L66
	}
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = int64(0)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v134 = int32(0)
	v136 = F_hash_search(m, v131, v12+int32(24), v134, v134)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v187 = v12 + int32(24)
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	F_hash_seq_init(m, v187, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L28
	}
L15:
	;
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)) = uint8(v157)
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v163 = int32(0)
	v165 = F_hash_search(m, v160, v12+int32(24), v163, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L22
	}
L16:
	;
	if v136 == int32(0) {
		v156 = v5
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v140 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v136)+80)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v136)+72)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v136)+64)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v136)+56)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v136)+416)) = v114
	v156 = v5
	goto L15
L19:
	;
	goto L20
L20:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v153 = F_hash_search(m, v150, v136, int32(2), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v156 = int32(1)
	goto L15
L22:
	;
	if v165 == int32(0) {
		v282 = v156
		goto L11
	} else {
		goto L23
	}
L23:
	;
	if l3 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v169 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v165)+80)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(v165)+72)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(v165)+64)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(v165)+56)) = v169
	*(*int64)(unsafe.Add(mBase, uint32(v165)+416)) = v114
	v282 = v156
	goto L11
L25:
	;
	goto L26
L26:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v182 = F_hash_search(m, v179, v165, int32(2), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v282 = v156 + int32(1)
	goto L11
L28:
	;
	v192 = F_hash_seq_search(m, v187)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v195 = int32(0)
	if base.B2i32(l0|l1 == v195)&base.B2i32(l2 == int64(0)) == v195 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v192 == int32(0) {
		v282 = v5
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v192 == int32(0) {
		v282 = v5
		goto L11
	} else {
		goto L55
	}
L33:
	;
	v208 = v192
	v210 = v5
	goto L34
L34:
	;
	if l0 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v282 = v238
	goto L11
L36:
	;
	v241 = F_hash_seq_search(m, v12+int32(24))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L53
	}
L37:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	if v213 != l0 {
		v238 = v210
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if l1 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	if v215 != l1 {
		v238 = v210
		goto L36
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if l2 != int64(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	v219 = *(*int64)(unsafe.Add(mBase, uint32(v208)+8))
	if v219 != l2 {
		v238 = v210
		goto L36
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if l3 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	v221 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v208)+80)) = v221
	*(*int64)(unsafe.Add(mBase, uint32(v208)+72)) = v221
	*(*int64)(unsafe.Add(mBase, uint32(v208)+64)) = v221
	*(*int64)(unsafe.Add(mBase, uint32(v208)+56)) = v221
	*(*int64)(unsafe.Add(mBase, uint32(v208)+416)) = v114
	v238 = v210
	goto L36
L50:
	;
	goto L51
L51:
	;
	v231 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v234 = F_hash_search(m, v231, v208, int32(2), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v238 = v210 + int32(1)
	goto L36
L53:
	;
	if v241 != 0 {
		v208 = v241
		v210 = v238
		goto L34
	} else {
		goto L54
	}
L54:
	;
	goto L35
L55:
	;
	v249 = v192
	v251 = v5
	goto L56
L56:
	;
	if l3 != 0 {
		goto L59
	} else {
		goto L60
	}
L57:
	;
	v282 = v271
	goto L11
L58:
	;
	v274 = F_hash_seq_search(m, v12+int32(24))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L4
	} else {
		goto L63
	}
L59:
	;
	v254 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v249)+80)) = v254
	*(*int64)(unsafe.Add(mBase, uint32(v249)+72)) = v254
	*(*int64)(unsafe.Add(mBase, uint32(v249)+64)) = v254
	*(*int64)(unsafe.Add(mBase, uint32(v249)+56)) = v254
	*(*int64)(unsafe.Add(mBase, uint32(v249)+416)) = v114
	v271 = v251
	goto L58
L60:
	;
	goto L61
L61:
	;
	v264 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v267 = F_hash_search(m, v264, v249, int32(2), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v271 = v251 + int32(1)
	goto L58
L63:
	;
	if v274 != 0 {
		v249 = v274
		v251 = v271
		goto L56
	} else {
		goto L64
	}
L64:
	;
	goto L57
L65:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v286)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v286)+20)) = int32(1)
	if v288 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v390 = v286
	goto L67
L67:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	F_LWLockRelease(m, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L4
	} else {
		goto L98
	}
L68:
	;
	v292 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	F_s_lock(m, v292+int32(20), int32(_a_F_entry_reset_0), int32(2749), int32(_a_F_entry_reset_1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L4
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+48)) = v114
	*(*int64)(unsafe.Add(mBase, uint32(v301)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v301)+20)) = int32(0)
	v309 = F_AllocateFile(m, int32(_a_F_entry_reset_2), int32(_a_F_entry_reset_3))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	v366 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+24)) = int32(0)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v366)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v366)+20)) = int32(1)
	if v369 != 0 {
		goto L94
	} else {
		goto L95
	}
L73:
	;
	if v309 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v315 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v309)+60))
	if v331 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	if v315 == int32(0) {
		goto L72
	} else {
		goto L78
	}
L78:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_entry_reset_2)
	F_errmsg(m, int32(_a_F_entry_reset_4), v12)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_entry_reset_0), int32(2764), int32(_a_F_entry_reset_1))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	goto L72
L82:
	;
	v363 = F_FreeFile(m, v309)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L93
	}
L83:
	;
	v340 = F_ftruncate(m, v338, int64(0))
	mBase = m.M
	if v340 == int32(0) {
		goto L82
	} else {
		goto L87
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_entry_reset[2])) = int32(8)
	v338 = int32(-1)
	goto L86
L85:
	;
	v338 = v331
	goto L86
L86:
	;
	goto L83
L87:
	;
	v345 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	if v345 == int32(0) {
		goto L82
	} else {
		goto L89
	}
L89:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_entry_reset_2)
	F_errmsg(m, int32(_a_F_entry_reset_5), v12+int32(16))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_entry_reset_0), int32(2773), int32(_a_F_entry_reset_1))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	goto L82
L93:
	;
	goto L72
L94:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	F_s_lock(m, v373+int32(20), int32(_a_F_entry_reset_0), int32(2780), int32(_a_F_entry_reset_1))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L4
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v382 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v382)+20)) = int32(0)
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v382)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v382)+32)) = v385 + int32(1)
	v390 = v382
	goto L67
L97:
	;
	goto L96
L98:
	;
	m.G0 = v12 + int32(48)
	return v114
L99:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_entry_reset_6), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_entry_reset_0), int32(2687), int32(_a_F_entry_reset_1))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
