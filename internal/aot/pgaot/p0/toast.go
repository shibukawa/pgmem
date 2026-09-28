package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_toast_tuple_externalize(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v267 int32
	_ = v267
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v376 int32
	_ = v376
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v30 = v27 + l1<<(uint(int32(3))%32)
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v35 = v32 + l1*int32(12)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)))
	v38 = v36 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)) = uint8(v38)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v42 = m.G0
	v44 = v42 - int32(2048)
	m.G0 = v44
	v47 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+112))
	v52 = F_table_open(m, v50, int32(3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+52))
	v60 = F_toast_open_indexes(m, v52, int32(3), v44+int32(2044), v44+int32(2040))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v62 = base.I32_wrap_i64(v31)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v63&int32(1) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v40)+264))
	if v98 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v66 = int32(1)
	v69 = int32(base.Ui32(v63) >> (uint(v66) % 32))
	v73 = v69 - v66
	v93 = v73
	v94 = v73
	v96 = v62 + v66
	v97 = v69 + int32(3)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v74 = int32(4)
	v75 = v62 + v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v77 = int32(2)
	v78 = int32(base.Ui32(v76) >> (uint(v77) % 32))
	v80 = v78 - v74
	if v63&v77 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v93 = v80
	v94 = v80
	v96 = v75
	v97 = v78
	goto L5
L10:
	;
	goto L11
L11:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v93 = v80
	v94 = v80 | v85&int32(-1073741824)
	v96 = v75
	v97 = v85&int32(1073741823) + int32(4)
	goto L5
L12:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v44)+2040))
	if int32(0) < v167 {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v52)+56))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v44)+2044))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v102+v60<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+56))
	v109 = F_GetNewOidWithIndex(m, v52, v107, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if v41 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v167 = v93
	v171 = v102
	v175 = v109
	v181 = v101
	goto L12
L17:
	;
	goto L28
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v41)+10))
	v116 = F_toastrel_valueid_exists(m, v52, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v41)+14))
	if v111 == v98 {
		goto L18
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v44)+2044))
	v120 = v93
	v121 = v113
	goto L17
L22:
	;
	goto L21
L23:
	;
	if v116 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v118 = int32(0)
	goto L26
L25:
	;
	v118 = v93
	goto L26
L26:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v44)+2044))
	if v115 != 0 {
		v167 = v118
		v171 = v119
		v175 = v115
		v181 = v98
		goto L12
	} else {
		goto L27
	}
L27:
	;
	v120 = v118
	v121 = v119
	goto L17
L28:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v121+v60<<(uint(int32(2))%32))))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+56))
	v155 = F_GetNewOidWithIndex(m, v52, v153, int32(1))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v167 = v120
	v171 = v121
	v175 = v155
	v181 = v98
	goto L12
L30:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v40)+264))
	v159 = F_table_open(m, v157, int32(1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v161 = F_toastrel_valueid_exists(m, v159, v155)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_relation_close(m, v159, int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v161 != 0 {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	v202 = v167
	v213 = v96
	v217 = int32(0)
	goto L38
L36:
	;
	goto L37
L37:
	;
	if int32(0) < v192 {
		goto L64
	} else {
		goto L65
	}
L38:
	;
	v227 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44)+2014)) = uint8(v227)
	*(*uint16)(unsafe.Add(mBase, uint32(v44)+2012)) = uint16(v227)
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
	*(*int64)(unsafe.Add(mBase, uint32(v44)+2016)) = base.I64_extend_i32_u(v175)
	*(*int64)(unsafe.Add(mBase, uint32(v44)+2024)) = base.I64_extend_i32_s(v217)
	v238 = int32(1996)
	if base.Ui32(v238) <= base.Ui32(v202) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	v241 = v238
	goto L46
L45:
	;
	v241 = v202
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+12)) = v241<<(uint(int32(2))%32) + int32(16)
	if v241 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	base.MemoryCopy(m, v44+int32(16), v213, v241)
	goto L49
L48:
	;
	goto L49
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v44)+2032)) = base.I64_extend_i32_u(v44 + int32(12))
	v253 = F_heap_form_tuple(m, v54, v44+int32(2016), v44+int32(2012))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_heap_insert(m, v52, v253, v47, l2, int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if int32(0) < v192 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v267 = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	F_pfree(m, v253)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L62
	}
L55:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v171+v267<<(uint(int32(2))%32))))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+192))
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+20)))
	if v294 == int32(1) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L54
L57:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293)+12)))
	v302 = int32(0)
	v304 = F_index_insert(m, v292, v44+int32(2016), v44+int32(2012), v253+int32(4), v52, v301, v302, v302)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v307 = v267 + int32(1)
	if v307 != v192 {
		v267 = v307
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
	v340 = v202 - v241
	if int32(0) < v340 {
		v202 = v340
		v213 = v241 + v213
		v217 = v217 + int32(1)
		goto L38
	} else {
		goto L63
	}
L63:
	;
	goto L39
L64:
	;
	v376 = int32(0)
	goto L67
L65:
	;
	goto L66
L66:
	;
	F_pfree(m, v171)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L71
	}
L67:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v171+v376<<(uint(int32(2))%32))))
	F_relation_close(m, v401, int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L66
L69:
	;
	v406 = v376 + int32(1)
	if v406 != v192 {
		v376 = v406
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	F_relation_close(m, v52, int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v440 = F_palloc(m, int32(18))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v440)+14)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v440)+10)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v440)+6)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v440)+2)) = v97
	v446 = int32(_a_F_toast_tuple_externalize_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v440))) = uint16(v446)
	m.G0 = v44 + int32(2048)
	*(*int64)(unsafe.Add(mBase, uint32(v30))) = base.I64_extend_i32_u(v440)
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)))
	if v453&int32(2) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_pfree(m, base.I32_wrap_i64(v31))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	v460 = v453
	goto L76
L76:
	;
	v462 = v460 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)) = uint8(v462)
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	v466 = v464 | int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v466)
	return
L77:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)))
	v460 = v459
	goto L76
}
