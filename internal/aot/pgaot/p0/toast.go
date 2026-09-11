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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
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
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
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
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v383 int32
	_ = v383
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
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
	v53 = F__emscripten_memset_bulkmem(m, v43+int32(12), base.I32_extend8_s(int32(0)), int32(2000))
	mBase = m.M
	goto L3
L3:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+112))
	v57 = F_table_open(m, v55, int32(3))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v57)+52))
	v65 = F_toast_open_indexes(m, v57, int32(3), v43+int32(2028), v43+int32(8))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v67&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v39)+264))
	if v102 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L7:
	;
	v70 = int32(1)
	v73 = int32(base.Ui32(v67) >> (uint(v70) % 32))
	v77 = v73 - v70
	v98 = v77
	v99 = v77
	v100 = v30 + v70
	v101 = v73 + int32(3)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v78 = int32(4)
	v79 = v30 + v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v81 = int32(2)
	v82 = int32(base.Ui32(v80) >> (uint(v81) % 32))
	v84 = v82 - v78
	if v67&v81 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v98 = v84
	v99 = v84
	v100 = v79
	v101 = v82
	goto L6
L11:
	;
	goto L12
L12:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v98 = v84
	v99 = v84 | v89&int32(-1073741824)
	v100 = v79
	v101 = v89&int32(1073741823) + int32(4)
	goto L6
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+2016)) = v179
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+2015)) = uint8(v195)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+2013)) = uint16(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+2024)) = v43 + int32(12)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	if v195 < v174 {
		goto L36
	} else {
		goto L37
	}
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v57)+56))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2028))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v65<<(uint(int32(2))%32))))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+56))
	v113 = F_GetNewOidWithIndex(m, v57, v111, int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v40 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v174 = v98
	v178 = v106
	v179 = v113
	v187 = v105
	goto L13
L18:
	;
	goto L29
L19:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v40)+10))
	v120 = F_toastrel_valueid_exists(m, v57, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L24
	}
L20:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v40)+14))
	if v115 == v102 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2028))
	v124 = v98
	v125 = v117
	goto L18
L23:
	;
	goto L22
L24:
	;
	if v120 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v122 = int32(0)
	goto L27
L26:
	;
	v122 = v98
	goto L27
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v43)+2028))
	if v119 != 0 {
		v174 = v122
		v178 = v123
		v179 = v119
		v187 = v102
		goto L13
	} else {
		goto L28
	}
L28:
	;
	v124 = v122
	v125 = v123
	goto L18
L29:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v125+v65<<(uint(int32(2))%32))))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+56))
	v158 = F_GetNewOidWithIndex(m, v57, v156, int32(1))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	v174 = v124
	v178 = v125
	v179 = v158
	v187 = v102
	goto L13
L31:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v39)+264))
	v162 = F_table_open(m, v160, int32(1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v164 = F_toastrel_valueid_exists(m, v162, v158)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_sequence_close(m, v162, int32(1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	if v164 != 0 {
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	v212 = v174
	v220 = v100
	v222 = v4
	goto L39
L37:
	;
	goto L38
L38:
	;
	if int32(0) < v202 {
		goto L66
	} else {
		goto L67
	}
L39:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_toast_tuple_externalize[0]))
	if v233 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v236 = int32(1996)
	if base.Ui32(v236) <= base.Ui32(v212) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	v239 = v236
	goto L47
L46:
	;
	v239 = v212
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v239<<(uint(int32(2))%32) + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+2020)) = v222
	if v239 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v252 = F_heap_form_tuple(m, v59, v43+int32(2016), v43+int32(2013))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	v246 = F__emscripten_memcpy_bulkmem(m, v43+int32(16), v220, v239)
	mBase = m.M
	goto L51
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	F_heap_insert(m, v57, v252, v46, l2, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if int32(0) < v202 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v265 = int32(0)
	goto L57
L55:
	;
	goto L56
L56:
	;
	F_pfree(m, v252)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L64
	}
L57:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v178+v265<<(uint(int32(2))%32))))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v290)+192))
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+20)))
	if v292 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+12)))
	v300 = int32(0)
	v302 = F_index_insert(m, v290, v43+int32(2016), v43+int32(2013), v252+int32(4), v57, v299, v300, v300)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v305 = v265 + int32(1)
	if v305 != v202 {
		v265 = v305
		goto L57
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	goto L58
L64:
	;
	v337 = v212 - v239
	if int32(0) < v337 {
		v212 = v337
		v220 = v220 + v239
		v222 = v222 + int32(1)
		goto L39
	} else {
		goto L65
	}
L65:
	;
	goto L40
L66:
	;
	v383 = v4
	goto L69
L67:
	;
	goto L68
L68:
	;
	F_pfree(m, v178)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L73
	}
L69:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v178+v383<<(uint(int32(2))%32))))
	F_relation_close(m, v395, int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	v400 = v383 + int32(1)
	if v400 != v202 {
		v383 = v400
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	F_sequence_close(m, v57, int32(0))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v433 = F_palloc(m, int32(18))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433)+14)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v433)+10)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v433)+6)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v433)+2)) = v101
	v439 = int32(_a_F_toast_tuple_externalize_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v433))) = uint16(v439)
	m.G0 = v43 + int32(2032)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v433
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)))
	if v445&int32(2) != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_pfree(m, v30)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	v451 = v445
	goto L78
L78:
	;
	v453 = v451 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)) = uint8(v453)
	v455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v457 = v455 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v457)
	return
L79:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+8)))
	v451 = v450
	goto L78
}
