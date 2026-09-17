package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EventTriggerAlterTableRelid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerAlterTableRelid[0]))
	if v4 == int32(0) {
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+20)))
		if v7 != 0 {
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
		}
	}
	return
}
func F_EventTriggerCommonSetup(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v79 int32
	_ = v79
	v5 = int32(0)
	v10 = F_EventCacheLookup(m, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v79
L2:
	;
	return int32(0)
L3:
	;
	if v10 == int32(0) {
		v79 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if l1 != int32(4) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v19 = F_CreateCommandTag(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L8
	}
L6:
	;
	v21 = int32(162)
	goto L7
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v22 <= int32(0) {
		v79 = v5
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v21 = v19
	goto L7
L9:
	;
	v27 = int32(0)
	v33 = v5
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v27<<(uint(int32(2))%32))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+4)))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerCommonSetup[0]))
	if v42 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if v58 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v60 = v27 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v60 < v61 {
		v27 = v60
		v33 = v58
		goto L10
	} else {
		goto L25
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	if v40 != int32(79) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v40 == int32(82) {
		v58 = v33
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v58 = v33
	goto L12
L18:
	;
	goto L13
L19:
	;
	v50 = F_bms_is_member(m, v21, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v55 = F_lappend_oid(m, v33, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	if v50 == int32(0) {
		v58 = v33
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v58 = v55
	goto L12
L25:
	;
	goto L11
L26:
	;
	return int32(0)
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(441)
	v79 = v58
	goto L1
}
func F_EventTriggerUndoInhibitCommandCollection(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_EventTriggerUndoInhibitCommandCollection[0]))
	if v3 != 0 {
		v4 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+20)) = uint8(v4)
	} else {
	}
	return
}
func F_WaitEventSetWait(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v27 int64
	_ = v27
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v297 int32
	_ = v297
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v369 int32
	_ = v369
	var v377 int64
	_ = v377
	var v378 int64
	_ = v378
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	v16 = m.G0
	v18 = v16 - int32(1040)
	m.G0 = v18
	if l1 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = int64(0)
	v33 = int32(-1)
	goto L3
L2:
	;
	F___clock_gettime(m, int32(1), v18+int32(16))
	mBase = m.M
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	v30 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+24)))
	v32 = v27*int64(-1000000000) - v30
	v33 = l1
	goto L3
L3:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = l4
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = int32(1)
	v41 = l1
	v42 = l2
	v52 = v33
	goto L6
L4:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L24
	} else {
		goto L93
	}
L5:
	;
	v421 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v421
	F_errstart_cold(m, int32(21), v421)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L24
	} else {
		goto L90
	}
L6:
	;
	v55 = int32(0)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v56 == v55 {
		v93 = v41
		v94 = v42
		v96 = v55
		v97 = v52
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v410 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v410
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v413))) = v410
	m.G0 = v18 + int32(1040)
	return v404
L8:
	;
	goto L7
L9:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = F_pgmem_poll(m, v98, v99, v97)
	mBase = m.M
	if v100 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v59 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v64 == int32(0) {
		v93 = v41
		v94 = v42
		v96 = v55
		v97 = v52
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(-1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v75+v73<<(uint(int32(4))%32))+12))
	v80 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v79
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v84
	if l3 == v80 {
		v404 = v80
		goto L8
	} else {
		goto L16
	}
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v67 == int32(0) {
		v93 = v41
		v94 = v42
		v96 = v55
		v97 = v52
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v93 = int32(0)
	v94 = v42 + int32(16)
	v96 = v80
	v97 = v84
	goto L9
L17:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v359 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L18:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[2]))
	if v104 == int32(27) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v100 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v355 = int32(0)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v109 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v109
	F_errstart_cold(m, int32(21), v109)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(_a_F_WaitEventSetWait_0)
	F_errmsg(m, int32(_a_F_WaitEventSetWait_1), v18)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_2), int32(1492), int32(_a_F_WaitEventSetWait_3))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v355 = int32(-1)
	goto L17
L30:
	;
	goto L31
L31:
	;
	v132 = int32(0)
	v133 = l3 - v96
	if v133 <= v132 {
		v355 = v132
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v136 <= int32(0) {
		v355 = v132
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v146 = v94
	v149 = v139
	v150 = v140
	v152 = v132
	goto L34
L34:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+6)))
	if v156 == int32(0) {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v355 = v329
	goto L17
L36:
	;
	v334 = v150 + int32(16)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v335+v336<<(uint(int32(4))%32)) <= base.Ui32(v334) {
		v355 = v329
		goto L17
	} else {
		goto L80
	}
L37:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+12)) = v161
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v167 = v165 - int32(1)
	if v167 != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v323 = v146 + int32(16)
	v329 = v152 + int32(1)
	goto L36
L39:
	;
	if v165&int32(134) == int32(0) {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L68
	}
L40:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
	if v236&int32(57) == int32(0) {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L64
	}
L41:
	;
	if v167 == int32(15) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
	if v170&int32(57) == int32(0) {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L47
	}
L44:
	;
	goto L40
L45:
	;
	goto L39
L47:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[3]))
	goto L48
L48:
	;
	v195 = F_read(m, v176, v18+int32(16), int32(1024))
	mBase = m.M
	if v195 < int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v225 == int32(0) {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L61
	}
L50:
	;
	goto L49
L51:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[2]))
	if v199 == int32(27) {
		goto L48
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v195 == int32(0) {
		goto L5
	} else {
		goto L59
	}
L54:
	;
	if v199 == int32(6) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	v205 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v205
	F_errstart_cold(m, int32(21), v205)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L24
	} else {
		goto L56
	}
L56:
	;
	F_errmsg_internal(m, int32(_a_F_WaitEventSetWait_4), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L24
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_2), int32(1970), int32(_a_F_WaitEventSetWait_5))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L24
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	if base.Ui32(int32(1023)) < base.Ui32(v195) {
		goto L48
	} else {
		goto L60
	}
L60:
	;
	goto L50
L61:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v225)+4))
	if v228 == int32(0) {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L62
	}
L62:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v225)))
	if v231 == int32(0) {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L63
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v146)+4)) = int64(-4294967295)
	goto L38
L64:
	;
	v241 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L24
	} else {
		goto L65
	}
L65:
	;
	if v241 != 0 {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L66
	}
L66:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v243 == int32(1) {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v146)+4)) = int64(-4294967280)
	goto L38
L68:
	;
	v252 = int32(0)
	if v165&int32(2) == v252 {
		v266 = v165
		v267 = v252
		goto L69
	} else {
		goto L70
	}
L69:
	;
	if v266&int32(4) == int32(0) {
		v281 = v266
		v282 = v267
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
	if v257&int32(57) == int32(0) {
		v266 = v165
		v267 = v252
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v262 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v262
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v266 = v265
	v267 = v262
	goto L69
L72:
	;
	if v281&int32(128) == int32(0) {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+6)))
	if v272&int32(60) == int32(0) {
		v281 = v266
		v282 = v267
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v278 = v267 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v281 = v280
	v282 = v278
	goto L72
L75:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v146)+8)) = v297
	goto L38
L76:
	;
	if v282 == int32(0) {
		v323 = v146
		v329 = v152
		goto L36
	} else {
		goto L79
	}
L77:
	;
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149)+6)))
	if v287&int32(_a_F_WaitEventSetWait_6) == int32(0) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v146)+4)) = v282 | int32(128)
	goto L75
L79:
	;
	goto L75
L80:
	;
	if v329 < v133 {
		v146 = v323
		v149 = v149 + int32(8)
		v150 = v334
		v152 = v329
		goto L34
	} else {
		goto L81
	}
L81:
	;
	goto L35
L82:
	;
	if v355 == int32(-1) {
		v404 = v96
		goto L8
	} else {
		goto L85
	}
L83:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v359)+4))
	if v362 == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v359)+4)) = int32(0)
	goto L82
L85:
	;
	v369 = v96 + v355
	if v369|base.B2i32(v93 < int32(0)) != 0 {
		v391 = v97
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if v369 == int32(0) {
		v41 = v93
		v42 = v94
		v52 = v391
		goto L6
	} else {
		goto L89
	}
L87:
	;
	F___clock_gettime(m, int32(1), v18+int32(16))
	mBase = m.M
	v377 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+24)))
	v378 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	v387 = v93 - base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i64_s(v377+(v378*int64(1000000000)+v32)), float64(1e+06)))
	if int32(0) < v387 {
		v391 = v387
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v404 = int32(0)
	goto L8
L89:
	;
	v404 = v369
	goto L8
L90:
	;
	F_errmsg_internal(m, int32(_a_F_WaitEventSetWait_7), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L24
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_2), int32(1980), int32(_a_F_WaitEventSetWait_5))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L24
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_event_trigger_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13852(m, l0, int32(_a_F_event_trigger_out_0), int32(367), int32(_a_F_event_trigger_out_1), int32(_a_F_event_trigger_out_2), int32(_a_F_event_trigger_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
