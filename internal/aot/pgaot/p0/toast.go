package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_toast_tuple_externalize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v273 int32
	_ = v273
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v383 int32
	_ = v383
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	v4 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = v26 + l1<<(uint(int32(2))%32)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v34 = v31 + l1*int32(12)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)))
	v37 = v35 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)) = uint8(v37)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v41 = m.G0
	v43 = v41 - int32(2032)
	m.G0 = v43
	v46 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	base.MemoryFill(m, v43+int32(12), int32(0), int32(2000))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+112))
	v56 = F_table_open(m, v54, int32(3))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+52))
	v64 = F_toast_open_indexes(m, v56, int32(3), v43+int32(2028), v43+int32(8))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v66&int32(1) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v39)+264))
	if v101 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v69 = int32(1)
	v72 = int32(base.Ui32(v66) >> (uint(v69) % 32))
	v76 = v72 - v69
	v96 = v76
	v97 = v76
	v99 = v30 + v69
	v100 = v72 + int32(3)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v77 = int32(4)
	v78 = v30 + v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v80 = int32(2)
	v81 = int32(base.Ui32(v79) >> (uint(v80) % 32))
	v83 = v81 - v77
	if v66&v80 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v96 = v83
	v97 = v83
	v99 = v78
	v100 = v81
	goto L5
L10:
	;
	goto L11
L11:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v96 = v83
	v97 = v83 | v88&int32(-1073741824)
	v99 = v78
	v100 = v88&int32(1073741823) + int32(4)
	goto L5
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+2016)) = v177
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+2015)) = uint8(v194)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+2013)) = uint16(v194)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+2024)) = v43 + int32(12)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v194 < v169 {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v56)+56))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2028))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105+v64<<(uint(int32(2))%32))))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+56))
	v112 = F_GetNewOidWithIndex(m, v56, v110, int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v40 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v169 = v96
	v172 = v105
	v177 = v112
	v184 = v104
	goto L12
L17:
	;
	goto L28
L18:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v40)+10))
	v119 = F_toastrel_valueid_exists(m, v56, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v40)+14))
	if v114 == v101 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2028))
	v123 = v96
	v124 = v116
	goto L17
L22:
	;
	goto L21
L23:
	;
	if v119 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v121 = int32(0)
	goto L26
L25:
	;
	v121 = v96
	goto L26
L26:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2028))
	if v118 != 0 {
		v169 = v121
		v172 = v122
		v177 = v118
		v184 = v101
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v123 = v121
	v124 = v122
	goto L17
L28:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v124+v64<<(uint(int32(2))%32))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+56))
	v157 = F_GetNewOidWithIndex(m, v56, v155, int32(1))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v169 = v123
	v172 = v124
	v177 = v157
	v184 = v101
	goto L12
L30:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v39)+264))
	v161 = F_table_open(m, v159, int32(1))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v163 = F_toastrel_valueid_exists(m, v161, v157)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_relation_close(m, v161, int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v163 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	v207 = v169
	v218 = v99
	v223 = v4
	goto L38
L36:
	;
	goto L37
L37:
	;
	if int32(0) < v201 {
		goto L64
	} else {
		goto L65
	}
L38:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_toast_tuple_externalize[0]))
	if v232 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L37
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v235 = int32(1996)
	if base.Ui32(v235) <= base.Ui32(v207) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	v238 = v235
	goto L46
L45:
	;
	v238 = v207
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v238<<(uint(int32(2))%32) + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+2020)) = v223
	if v238 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	base.MemoryCopy(m, v43+int32(16), v218, v238)
	goto L49
L48:
	;
	goto L49
L49:
	;
	v250 = F_heap_form_tuple(m, v58, v43+int32(2016), v43+int32(2013))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_heap_insert(m, v56, v250, v46, l2, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if int32(0) < v201 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v273 = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	F_pfree(m, v250)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L62
	}
L55:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v172+v273<<(uint(int32(2))%32))))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+192))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+20)))
	if v290 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+12)))
	v298 = int32(0)
	v300 = F_index_insert(m, v288, v43+int32(2016), v43+int32(2013), v250+int32(4), v56, v297, v298, v298)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v303 = v273 + int32(1)
	if v303 != v201 {
		v273 = v303
		goto L55
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	goto L56
L62:
	;
	v335 = v207 - v238
	if int32(0) < v335 {
		v207 = v335
		v218 = v238 + v218
		v223 = v223 + int32(1)
		goto L38
	} else {
		goto L63
	}
L63:
	;
	goto L39
L64:
	;
	v383 = v4
	goto L67
L65:
	;
	goto L66
L66:
	;
	F_pfree(m, v172)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L71
	}
L67:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v172+v383<<(uint(int32(2))%32))))
	F_relation_close(m, v393, int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L66
L69:
	;
	v398 = v383 + int32(1)
	if v398 != v201 {
		v383 = v398
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	F_relation_close(m, v56, int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v431 = F_palloc(m, int32(18))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v431)+14)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v431)+10)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v431)+6)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v431)+2)) = v100
	v437 = int32(_a_F_toast_tuple_externalize_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v431))) = uint16(v437)
	m.G0 = v43 + int32(2032)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v431
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)))
	if v443&int32(2) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_pfree(m, v30)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v449 = v443
	goto L76
L76:
	;
	v451 = v449 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)) = uint8(v451)
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v455 = v453 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v455)
	return
L77:
	;
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)))
	v449 = v448
	goto L76
}
