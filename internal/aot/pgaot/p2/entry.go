package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v175 int64
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
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
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 float64
	_ = v222
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v230 float64
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v264 float64
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int64
	_ = v335
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int64
	_ = v346
	var v347 int64
	_ = v347
	var v350 int64
	_ = v350
	var v351 int64
	_ = v351
	var v352 int64
	_ = v352
	var v353 int64
	_ = v353
	var v354 int64
	_ = v354
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v358 int64
	_ = v358
	var v359 int64
	_ = v359
	var v360 int64
	_ = v360
	var v361 int64
	_ = v361
	var v362 int64
	_ = v362
	var v363 int64
	_ = v363
	var v364 int64
	_ = v364
	var v365 int64
	_ = v365
	var v366 int64
	_ = v366
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v369 int64
	_ = v369
	var v370 int64
	_ = v370
	var v371 int64
	_ = v371
	var v372 int64
	_ = v372
	var v373 int64
	_ = v373
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v376 int64
	_ = v376
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v379 int64
	_ = v379
	var v380 int64
	_ = v380
	var v412 int64
	_ = v412
	var v414 int64
	_ = v414
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v445 int32
	_ = v445
	var v446 float64
	_ = v446
	var v448 float64
	_ = v448
	var v450 int32
	_ = v450
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int64
	_ = v464
	var v465 int64
	_ = v465
	var v473 int64
	_ = v473
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v21)+808))
	if v23 != int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v90 = int64(*(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[1])))
	if v90 <= v88 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v21)+752))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v21)+728))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v21)+704))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v21)+680))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v21)+656))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v21)+632))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v21)+608))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v21)+584))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v21)+560))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v21)+536))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v21)+512))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v21)+488))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v21)+464))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v21)+440))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v21)+416))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v21)+392))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v21)+368))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v21)+344))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v21)+320))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v21)+296))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v21)+272))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v21)+248))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v21)+224))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v21)+200))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v21)+176))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v21)+152))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v21)+128))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v21)+104))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v21)+80))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v21)+56))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	v88 = v26 + (v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + v22))))))))))))))))))))))))))))))
	goto L4
L3:
	;
	v88 = v22
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
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v434 = F_hash_search(m, v430, l0, int32(1), v16+int32(12))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L14
	} else {
		goto L58
	}
L8:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v108)+8))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v108)+808))
	if v110 != int64(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v179 = F_palloc(m, base.I32_wrap_i64(v175)<<(uint(int32(2))%32))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v108)+752))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v108)+728))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v108)+704))
	v116 = *(*int64)(unsafe.Add(mBase, uint32(v108)+680))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v108)+656))
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v108)+632))
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v108)+608))
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v108)+584))
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v108)+560))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v108)+536))
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v108)+512))
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v108)+488))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v108)+464))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v108)+440))
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v108)+416))
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v108)+392))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v108)+368))
	v130 = *(*int64)(unsafe.Add(mBase, uint32(v108)+344))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v108)+320))
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v108)+296))
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v108)+272))
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v108)+248))
	v135 = *(*int64)(unsafe.Add(mBase, uint32(v108)+224))
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v108)+200))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(v108)+176))
	v138 = *(*int64)(unsafe.Add(mBase, uint32(v108)+152))
	v139 = *(*int64)(unsafe.Add(mBase, uint32(v108)+128))
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v108)+104))
	v141 = *(*int64)(unsafe.Add(mBase, uint32(v108)+80))
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v108)+56))
	v143 = *(*int64)(unsafe.Add(mBase, uint32(v108)+32))
	v175 = v113 + (v114 + (v115 + (v116 + (v117 + (v118 + (v119 + (v120 + (v121 + (v122 + (v123 + (v124 + (v125 + (v126 + (v127 + (v128 + (v129 + (v130 + (v131 + (v132 + (v133 + (v134 + (v135 + (v136 + (v137 + (v138 + (v139 + (v140 + (v141 + (v142 + (v143 + v109))))))))))))))))))))))))))))))
	goto L13
L12:
	;
	v175 = v109
	goto L13
L13:
	;
	goto L10
L14:
	;
	return int32(0)
L15:
	;
	v184 = v16 + int32(12)
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	F_hash_seq_init(m, v184, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v189 = int32(0)
	v192 = F_hash_seq_search(m, v184)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L14
	} else {
		goto L18
	}
L17:
	;
	F_pfree(m, v179)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L14
	} else {
		goto L48
	}
L18:
	;
	if v192 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pg_qsort(m, v179, int32(0), int32(4), int32(_a_F_entry_alloc_0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L14
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v210 = v192
	v211 = v189
	v214 = v189
	v216 = v189
	goto L23
L22:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+136)) = int32(1024)
	goto L17
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179+v211<<(uint(int32(2))%32)))) = v210
	v222 = *(*float64)(unsafe.Add(mBase, uint32(v210)+256))
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v210)+24))
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v210)+32))
	if v225 == int64(0)-v227 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	F_pg_qsort(m, v179, v234, int32(4), int32(_a_F_entry_alloc_0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L14
	} else {
		goto L33
	}
L25:
	;
	v230 = float64(0.5)
	goto L27
L26:
	;
	v230 = float64(0.99)
	goto L27
L27:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v210)+256)) = base.F64_mul(v222, v230)
	v234 = v211 + int32(1)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v210)+412))
	v236 = int32(-1)
	v240 = v214 + int32(base.Ui32(v235^v236)>>(uint(int32(31))%32))
	if v235 < int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v244 = v236
	goto L30
L29:
	;
	v244 = v235
	goto L30
L30:
	;
	v247 = v216 + v244 + int32(1)
	v250 = F_hash_seq_search(m, v16+int32(12))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	if v250 != 0 {
		v210 = v250
		v211 = v234
		v214 = v240
		v216 = v247
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L24
L33:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v179+v234<<(uint(int32(1))%32)&int32(-4))))
	v264 = *(*float64)(unsafe.Add(mBase, uint32(v263)+256))
	*(*float64)(unsafe.Add(mBase, uint32(v257)+128)) = v264
	if v240 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v266 = base.I32_div_u_s(v247, v240)
	v268 = v266
	goto L36
L35:
	;
	v268 = int32(1024)
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v257)+136)) = v268
	v271 = int32(10)
	if base.Ui32(v271) <= base.Ui32(v234) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v274 = v271
	goto L39
L38:
	;
	v274 = v234
	goto L39
L39:
	;
	v276 = base.I32_div_u_s(v234, int32(20))
	if base.Ui32(v211) < base.Ui32(int32(199)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v279 = v274
	goto L42
L41:
	;
	v279 = v276
	goto L42
L42:
	;
	if v279 == int32(0) {
		goto L17
	} else {
		goto L43
	}
L43:
	;
	v287 = int32(0)
	goto L44
L44:
	;
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v297 = int32(2)
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v179+v287<<(uint(v297)%32))))
	v303 = F_hash_search(m, v296, v300, v297, int32(0))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L14
	} else {
		goto L46
	}
L45:
	;
	goto L17
L46:
	;
	v306 = v287 + int32(1)
	if v306 != v279 {
		v287 = v306
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
	v327 = base.AtomicRmwXchg32(m, v324, int32(140), int32(1))
	if v327 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_s_lock(m, v324+int32(140), int32(_a_F_entry_alloc_1))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L14
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v334 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v334)+160))
	*(*int64)(unsafe.Add(mBase, uint32(v334)+160)) = v335 + int64(1)
	v339 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v334)+140)), uint32(v339))
	v343 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[0]))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v345)+8))
	v347 = *(*int64)(unsafe.Add(mBase, uint32(v345)+808))
	if v347 != int64(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L51
L53:
	;
	v414 = int64(*(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[1])))
	if v414 <= v412 {
		goto L8
	} else {
		goto L57
	}
L54:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v345)+752))
	v351 = *(*int64)(unsafe.Add(mBase, uint32(v345)+728))
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v345)+704))
	v353 = *(*int64)(unsafe.Add(mBase, uint32(v345)+680))
	v354 = *(*int64)(unsafe.Add(mBase, uint32(v345)+656))
	v355 = *(*int64)(unsafe.Add(mBase, uint32(v345)+632))
	v356 = *(*int64)(unsafe.Add(mBase, uint32(v345)+608))
	v357 = *(*int64)(unsafe.Add(mBase, uint32(v345)+584))
	v358 = *(*int64)(unsafe.Add(mBase, uint32(v345)+560))
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v345)+536))
	v360 = *(*int64)(unsafe.Add(mBase, uint32(v345)+512))
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v345)+488))
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v345)+464))
	v363 = *(*int64)(unsafe.Add(mBase, uint32(v345)+440))
	v364 = *(*int64)(unsafe.Add(mBase, uint32(v345)+416))
	v365 = *(*int64)(unsafe.Add(mBase, uint32(v345)+392))
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v345)+368))
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v345)+344))
	v368 = *(*int64)(unsafe.Add(mBase, uint32(v345)+320))
	v369 = *(*int64)(unsafe.Add(mBase, uint32(v345)+296))
	v370 = *(*int64)(unsafe.Add(mBase, uint32(v345)+272))
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v345)+248))
	v372 = *(*int64)(unsafe.Add(mBase, uint32(v345)+224))
	v373 = *(*int64)(unsafe.Add(mBase, uint32(v345)+200))
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v345)+176))
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v345)+152))
	v376 = *(*int64)(unsafe.Add(mBase, uint32(v345)+128))
	v377 = *(*int64)(unsafe.Add(mBase, uint32(v345)+104))
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v345)+80))
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v345)+56))
	v380 = *(*int64)(unsafe.Add(mBase, uint32(v345)+32))
	v412 = v350 + (v351 + (v352 + (v353 + (v354 + (v355 + (v356 + (v357 + (v358 + (v359 + (v360 + (v361 + (v362 + (v363 + (v364 + (v365 + (v366 + (v367 + (v368 + (v369 + (v370 + (v371 + (v372 + (v373 + (v374 + (v375 + (v376 + (v377 + (v378 + (v379 + (v380 + v346))))))))))))))))))))))))))))))
	goto L56
L55:
	;
	v412 = v346
	goto L56
L56:
	;
	goto L53
L57:
	;
	goto L9
L58:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)))
	if v436 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryFill(m, v434+int32(24), int32(0), int32(384))
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
	m.G0 = v16 + int32(32)
	return v434
L62:
	;
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_entry_alloc[2]))
	v446 = *(*float64)(unsafe.Add(mBase, uint32(v445)+128))
	v448 = v446
	goto L64
L63:
	;
	v448 = float64(1)
	goto L64
L64:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v434)+256)) = v448
	v450 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v434)+440)), uint32(v450))
	*(*int32)(unsafe.Add(mBase, uint32(v434)+416)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v434)+412)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v434)+408)) = l1
	v459 = m.G0
	v460 = int32(16)
	v461 = v459 - v460
	m.G0 = v461
	F_gettimeofday(m, v461)
	mBase = m.M
	v464 = *(*int64)(unsafe.Add(mBase, uint32(v461)))
	v465 = int64(*(*int32)(unsafe.Add(mBase, uint32(v461)+8)))
	m.G0 = v461 + v460
	v473 = v465 + v464*int64(1000000) - int64(946684800000000)
	goto L65
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v434)+432)) = v473
	*(*int64)(unsafe.Add(mBase, uint32(v434)+424)) = v473
	goto L61
}
func F_entry_reset(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int64
	_ = v170
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
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
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int64
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v255 int64
	_ = v255
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int64
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v283 int64
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	v7 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	if v15 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
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
	v23 = F_LWLockAcquire(m, v15, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v30)+808))
	if v32 != int64(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v101 = m.G0
	v102 = int32(16)
	v103 = v101 - v102
	m.G0 = v103
	F_gettimeofday(m, v103)
	mBase = m.M
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	v107 = int64(*(*int32)(unsafe.Add(mBase, uint32(v103)+8)))
	m.G0 = v103 + v102
	v115 = v107 + v106*int64(1000000) - int64(946684800000000)
	goto L10
L7:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v30)+752))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v30)+728))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v30)+704))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v30)+680))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v30)+656))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v30)+632))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v30)+608))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v30)+584))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v30)+560))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v30)+536))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v30)+512))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v30)+488))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v30)+464))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v30)+440))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v30)+416))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v30)+392))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v30)+368))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v30)+344))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v30)+320))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v30)+296))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v30)+272))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v30)+248))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v30)+224))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v30)+200))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v30)+176))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v30)+152))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v30)+128))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v30)+104))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v30)+80))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v30)+56))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v30)+32))
	v97 = v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + (v57 + (v58 + (v59 + (v60 + (v61 + (v62 + (v63 + (v64 + (v65 + v31))))))))))))))))))))))))))))))
	goto L9
L8:
	;
	v97 = v31
	goto L9
L9:
	;
	goto L6
L10:
	;
	v118 = int32(0)
	if base.B2i32(l2 == int64(0))|(base.B2i32(l0 == v118)|base.B2i32(l1 == v118)) == v118 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	if v283 == v97 {
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
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v135 = int32(0)
	v137 = F_hash_search(m, v132, v12+int32(24), v135, v135)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v188 = v12 + int32(24)
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	F_hash_seq_init(m, v188, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L4
	} else {
		goto L28
	}
L15:
	;
	v158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)) = uint8(v158)
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v164 = int32(0)
	v166 = F_hash_search(m, v161, v12+int32(24), v164, v164)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L22
	}
L16:
	;
	if v137 == int32(0) {
		v157 = v7
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
	v141 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v137)+80)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v137)+72)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v137)+64)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v137)+56)) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v137)+432)) = v115
	v157 = v7
	goto L15
L19:
	;
	goto L20
L20:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v154 = F_hash_search(m, v151, v137, int32(2), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v157 = int64(1)
	goto L15
L22:
	;
	if v166 == int32(0) {
		v283 = v157
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
	v170 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v166)+80)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v166)+72)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v166)+64)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v166)+56)) = v170
	*(*int64)(unsafe.Add(mBase, uint32(v166)+432)) = v115
	v283 = v157
	goto L11
L25:
	;
	goto L26
L26:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v183 = F_hash_search(m, v180, v166, int32(2), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v283 = v157 + int64(1)
	goto L11
L28:
	;
	v193 = F_hash_seq_search(m, v188)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v196 = int32(0)
	if base.B2i32(l0|l1 == v196)&base.B2i32(l2 == int64(0)) == v196 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v193 == int32(0) {
		v283 = v7
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if v193 == int32(0) {
		v283 = v7
		goto L11
	} else {
		goto L55
	}
L33:
	;
	v209 = v193
	v211 = v7
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
	v283 = v239
	goto L11
L36:
	;
	v242 = F_hash_seq_search(m, v12+int32(24))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L4
	} else {
		goto L53
	}
L37:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	if v214 != l0 {
		v239 = v211
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
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v216 != l1 {
		v239 = v211
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
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v209)+8))
	if v220 != l2 {
		v239 = v211
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
	v222 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v209)+80)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v209)+72)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v209)+64)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v209)+56)) = v222
	*(*int64)(unsafe.Add(mBase, uint32(v209)+432)) = v115
	v239 = v211
	goto L36
L50:
	;
	goto L51
L51:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v235 = F_hash_search(m, v232, v209, int32(2), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v239 = v211 + int64(1)
	goto L36
L53:
	;
	if v242 != 0 {
		v209 = v242
		v211 = v239
		goto L34
	} else {
		goto L54
	}
L54:
	;
	goto L35
L55:
	;
	v250 = v193
	v252 = v7
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
	v283 = v272
	goto L11
L58:
	;
	v275 = F_hash_seq_search(m, v12+int32(24))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L63
	}
L59:
	;
	v255 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v250)+80)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v250)+72)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v250)+64)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v250)+56)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v250)+432)) = v115
	v272 = v252
	goto L58
L60:
	;
	goto L61
L61:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[1]))
	v268 = F_hash_search(m, v265, v250, int32(2), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	v272 = v252 + int64(1)
	goto L58
L63:
	;
	if v275 != 0 {
		v250 = v275
		v252 = v272
		goto L56
	} else {
		goto L64
	}
L64:
	;
	goto L57
L65:
	;
	v291 = base.AtomicRmwXchg32(m, v287, int32(140), int32(1))
	if v291 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	v385 = v287
	goto L67
L67:
	;
	F_LWLockRelease(m, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L98
	}
L68:
	;
	F_s_lock(m, v287+int32(140), int32(_a_F_entry_reset_0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L4
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v298 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v298)+168)) = v115
	*(*int64)(unsafe.Add(mBase, uint32(v298)+160)) = int64(0)
	v302 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v298)+140)), uint32(v302))
	v307 = F_AllocateFile(m, int32(_a_F_entry_reset_1), int32(_a_F_entry_reset_2))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L70
L72:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v364)+144)) = int32(0)
	v369 = base.AtomicRmwXchg32(m, v364, int32(140), int32(1))
	if v369 != 0 {
		goto L94
	} else {
		goto L95
	}
L73:
	;
	if v307 == int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v313 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L4
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v307)+60))
	if v329 < int32(0) {
		goto L84
	} else {
		goto L85
	}
L77:
	;
	if v313 == int32(0) {
		goto L72
	} else {
		goto L78
	}
L78:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_entry_reset_1)
	F_errmsg(m, int32(_a_F_entry_reset_3), v12)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_entry_reset_4), int32(2768), int32(_a_F_entry_reset_5))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	goto L72
L82:
	;
	v361 = F_FreeFile(m, v307)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L4
	} else {
		goto L93
	}
L83:
	;
	v338 = F_ftruncate(m, v336, int64(0))
	mBase = m.M
	if v338 == int32(0) {
		goto L82
	} else {
		goto L87
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_entry_reset[2])) = int32(8)
	v336 = int32(-1)
	goto L86
L85:
	;
	v336 = v329
	goto L86
L86:
	;
	goto L83
L87:
	;
	v343 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	if v343 == int32(0) {
		goto L82
	} else {
		goto L89
	}
L89:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L4
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_entry_reset_1)
	F_errmsg(m, int32(_a_F_entry_reset_6), v12+int32(16))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L4
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_entry_reset_4), int32(2777), int32(_a_F_entry_reset_5))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
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
	F_s_lock(m, v364+int32(140), int32(_a_F_entry_reset_0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L4
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_entry_reset[0]))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v376)+152)) = v377 + int32(1)
	v381 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v376)+140)), uint32(v381))
	v385 = v376
	goto L67
L97:
	;
	goto L96
L98:
	;
	m.G0 = v12 + int32(48)
	return v115
L99:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	F_errmsg(m, int32(_a_F_entry_reset_7), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L4
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_entry_reset_4), int32(2691), int32(_a_F_entry_reset_5))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
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
