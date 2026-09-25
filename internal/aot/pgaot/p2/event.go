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
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v391 int64
	_ = v391
	var v392 int64
	_ = v392
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v416 int32
	_ = v416
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
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
	v452 = m.ExcPending
	if v452 != 0 {
		goto L30
	} else {
		goto L99
	}
L5:
	;
	v435 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v435
	F_errstart_cold(m, int32(21), v435)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L30
	} else {
		goto L96
	}
L6:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v55 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v424 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v424
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = v424
	m.G0 = v18 + int32(1040)
	return v416
L8:
	;
	goto L7
L9:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v106 = m.Env.Pgmem_poll(m, v104, v105, v103)
	mBase = m.M
	if int32(0) <= v106 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	v99 = v41
	v100 = v42
	v102 = int32(0)
	v103 = v52
	goto L9
L11:
	;
	goto L12
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v59 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = int32(1)
	v64 = int32(0)
	v68 = base.AtomicRmwOr32(m, v64, int32(_a_F_WaitEventSetWait_0), v64)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v69 == v64 {
		v99 = v41
		v100 = v42
		v102 = v64
		v103 = v52
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = int32(-1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v79
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81+v79<<(uint(int32(4))%32))+12))
	v86 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v42)+12)) = v85
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v90
	if l3 == v86 {
		v416 = v86
		goto L8
	} else {
		goto L18
	}
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v72 == int32(0) {
		v99 = v41
		v100 = v42
		v102 = v64
		v103 = v52
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v99 = int32(0)
	v100 = v42 + int32(16)
	v102 = v86
	v103 = v90
	goto L9
L19:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v373 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L20:
	;
	if v114 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v114 = v106
	goto L20
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[2])) = int32(0) - v106
	v114 = int32(-1)
	goto L20
L24:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[2]))
	if v118 == int32(27) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	if v114 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v369 = int32(0)
	goto L19
L28:
	;
	goto L29
L29:
	;
	v123 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v123
	F_errstart_cold(m, int32(21), v123)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	return int32(0)
L31:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(_a_F_WaitEventSetWait_1)
	F_errmsg(m, int32(_a_F_WaitEventSetWait_2), v18)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_3), int32(1492), int32(_a_F_WaitEventSetWait_4))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v369 = int32(-1)
	goto L19
L36:
	;
	goto L37
L37:
	;
	v146 = int32(0)
	v147 = l3 - v102
	if v147 <= v146 {
		v369 = v146
		goto L19
	} else {
		goto L38
	}
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v150 <= int32(0) {
		v369 = v146
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v160 = v100
	v164 = v153
	v165 = v154
	v166 = v146
	goto L40
L40:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+6)))
	if v170 == int32(0) {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v369 = v343
	goto L19
L42:
	;
	v348 = v165 + int32(16)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v349+v350<<(uint(int32(4))%32)) <= base.Ui32(v348) {
		v369 = v343
		goto L19
	} else {
		goto L86
	}
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v173
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+12)) = v175
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v181 = v179 - int32(1)
	if v181 != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v337 = v160 + int32(16)
	v343 = v166 + int32(1)
	goto L42
L45:
	;
	if v179&int32(134) == int32(0) {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L74
	}
L46:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+6)))
	if v250&int32(57) == int32(0) {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L70
	}
L47:
	;
	if v181 == int32(15) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+6)))
	if v184&int32(57) == int32(0) {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L53
	}
L50:
	;
	goto L46
L51:
	;
	goto L45
L53:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[3]))
	goto L54
L54:
	;
	v209 = F_read(m, v190, v18+int32(16), int32(1024))
	mBase = m.M
	if v209 < int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v239 == int32(0) {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L67
	}
L56:
	;
	goto L55
L57:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[2]))
	if v213 == int32(27) {
		goto L54
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v209 == int32(0) {
		goto L5
	} else {
		goto L65
	}
L60:
	;
	if v213 == int32(6) {
		goto L56
	} else {
		goto L61
	}
L61:
	;
	v219 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v219
	F_errstart_cold(m, int32(21), v219)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L30
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_WaitEventSetWait_5), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L30
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_3), int32(1970), int32(_a_F_WaitEventSetWait_6))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L30
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	if base.Ui32(int32(1023)) < base.Ui32(v209) {
		goto L54
	} else {
		goto L66
	}
L66:
	;
	goto L56
L67:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v239)+4))
	if v242 == int32(0) {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L68
	}
L68:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v239)))
	if v245 == int32(0) {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L69
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v160)+4)) = int64(-4294967295)
	goto L44
L70:
	;
	v255 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L30
	} else {
		goto L71
	}
L71:
	;
	if v255 != 0 {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L72
	}
L72:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v257 == int32(1) {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v160)+4)) = int64(-4294967280)
	goto L44
L74:
	;
	v266 = int32(0)
	if v179&int32(2) == v266 {
		v280 = v179
		v281 = v266
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v280&int32(4) == int32(0) {
		v295 = v280
		v296 = v281
		goto L78
	} else {
		goto L79
	}
L76:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+6)))
	if v271&int32(57) == int32(0) {
		v280 = v179
		v281 = v266
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v276 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v276
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v280 = v279
	v281 = v276
	goto L75
L78:
	;
	if v295&int32(128) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+6)))
	if v286&int32(60) == int32(0) {
		v295 = v280
		v296 = v281
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v292 = v281 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v295 = v294
	v296 = v292
	goto L78
L81:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+8)) = v311
	goto L44
L82:
	;
	if v296 == int32(0) {
		v337 = v160
		v343 = v166
		goto L42
	} else {
		goto L85
	}
L83:
	;
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+6)))
	if v301&int32(_a_F_WaitEventSetWait_7) == int32(0) {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v160)+4)) = v296 | int32(128)
	goto L81
L85:
	;
	goto L81
L86:
	;
	if v343 < v147 {
		v160 = v337
		v164 = v164 + int32(8)
		v165 = v348
		v166 = v343
		goto L40
	} else {
		goto L87
	}
L87:
	;
	goto L41
L88:
	;
	if v369 == int32(-1) {
		v416 = v102
		goto L8
	} else {
		goto L91
	}
L89:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v376 == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+4)) = int32(0)
	goto L88
L91:
	;
	v383 = v102 + v369
	if v383|base.B2i32(v99 < int32(0)) != 0 {
		v405 = v103
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v383 == int32(0) {
		v41 = v99
		v42 = v100
		v52 = v405
		goto L6
	} else {
		goto L95
	}
L93:
	;
	F___clock_gettime(m, int32(1), v18+int32(16))
	mBase = m.M
	v391 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+24)))
	v392 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	v401 = v99 - base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i64_s(v391+(v392*int64(1000000000)+v32)), float64(1e+06)))
	if int32(0) < v401 {
		v405 = v401
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v416 = int32(0)
	goto L8
L95:
	;
	v416 = v383
	goto L8
L96:
	;
	F_errmsg_internal(m, int32(_a_F_WaitEventSetWait_8), int32(0))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L30
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_3), int32(1980), int32(_a_F_WaitEventSetWait_6))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L30
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
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
	v7 = Fn13854(m, l0, int32(_a_F_event_trigger_out_0), int32(367), int32(_a_F_event_trigger_out_1), int32(_a_F_event_trigger_out_2), int32(_a_F_event_trigger_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
