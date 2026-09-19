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
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
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
	v446 = m.ExcPending
	if v446 != 0 {
		goto L28
	} else {
		goto L97
	}
L5:
	;
	v429 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v429
	F_errstart_cold(m, int32(21), v429)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L28
	} else {
		goto L94
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
	v418 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v418
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v421))) = v418
	m.G0 = v18 + int32(1040)
	return v412
L8:
	;
	goto L7
L9:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = m.Env.Pgmem_poll(m, v98, v99, v97)
	mBase = m.M
	if int32(0) <= v100 {
		goto L19
	} else {
		goto L20
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
		v412 = v80
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
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v367 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L18:
	;
	if v108 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v108 = v100
	goto L18
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[2])) = int32(0) - v100
	v108 = int32(-1)
	goto L18
L22:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[2]))
	if v112 == int32(27) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	if v108 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	v363 = int32(0)
	goto L17
L26:
	;
	goto L27
L27:
	;
	v117 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v117
	F_errstart_cold(m, int32(21), v117)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return int32(0)
L29:
	;
	F_errcode_for_socket_access(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(_a_F_WaitEventSetWait_0)
	F_errmsg(m, int32(_a_F_WaitEventSetWait_1), v18)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_2), int32(1492), int32(_a_F_WaitEventSetWait_3))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v363 = int32(-1)
	goto L17
L34:
	;
	goto L35
L35:
	;
	v140 = int32(0)
	v141 = l3 - v96
	if v141 <= v140 {
		v363 = v140
		goto L17
	} else {
		goto L36
	}
L36:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v144 <= int32(0) {
		v363 = v140
		goto L17
	} else {
		goto L37
	}
L37:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v154 = v94
	v157 = v147
	v158 = v148
	v160 = v140
	goto L38
L38:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+6)))
	if v164 == int32(0) {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v363 = v337
	goto L17
L40:
	;
	v342 = v158 + int32(16)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v343+v344<<(uint(int32(4))%32)) <= base.Ui32(v342) {
		v363 = v337
		goto L17
	} else {
		goto L84
	}
L41:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v167
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+12)) = v169
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v175 = v173 - int32(1)
	if v175 != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v331 = v154 + int32(16)
	v337 = v160 + int32(1)
	goto L40
L43:
	;
	if v173&int32(134) == int32(0) {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L72
	}
L44:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
	if v244&int32(57) == int32(0) {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L68
	}
L45:
	;
	if v175 == int32(15) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
	if v178&int32(57) == int32(0) {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L51
	}
L48:
	;
	goto L44
L49:
	;
	goto L43
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[3]))
	goto L52
L52:
	;
	v203 = F_read(m, v184, v18+int32(16), int32(1024))
	mBase = m.M
	if v203 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v233 == int32(0) {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L65
	}
L54:
	;
	goto L53
L55:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[2]))
	if v207 == int32(27) {
		goto L52
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v203 == int32(0) {
		goto L5
	} else {
		goto L63
	}
L58:
	;
	if v207 == int32(6) {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v213 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitEventSetWait[1])) = v213
	F_errstart_cold(m, int32(21), v213)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L28
	} else {
		goto L60
	}
L60:
	;
	F_errmsg_internal(m, int32(_a_F_WaitEventSetWait_4), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L28
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_2), int32(1970), int32(_a_F_WaitEventSetWait_5))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L28
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	if base.Ui32(int32(1023)) < base.Ui32(v203) {
		goto L52
	} else {
		goto L64
	}
L64:
	;
	goto L54
L65:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v236 == int32(0) {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L66
	}
L66:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if v239 == int32(0) {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v154)+4)) = int64(-4294967295)
	goto L42
L68:
	;
	v249 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L28
	} else {
		goto L69
	}
L69:
	;
	if v249 != 0 {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L70
	}
L70:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v251 == int32(1) {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v154)+4)) = int64(-4294967280)
	goto L42
L72:
	;
	v260 = int32(0)
	if v173&int32(2) == v260 {
		v274 = v173
		v275 = v260
		goto L73
	} else {
		goto L74
	}
L73:
	;
	if v274&int32(4) == int32(0) {
		v289 = v274
		v290 = v275
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
	if v265&int32(57) == int32(0) {
		v274 = v173
		v275 = v260
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v270 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v270
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v274 = v273
	v275 = v270
	goto L73
L76:
	;
	if v289&int32(128) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
	if v280&int32(60) == int32(0) {
		v289 = v274
		v290 = v275
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v286 = v275 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v289 = v288
	v290 = v286
	goto L76
L79:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v154)+8)) = v305
	goto L42
L80:
	;
	if v290 == int32(0) {
		v331 = v154
		v337 = v160
		goto L40
	} else {
		goto L83
	}
L81:
	;
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+6)))
	if v295&int32(_a_F_WaitEventSetWait_6) == int32(0) {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154)+4)) = v290 | int32(128)
	goto L79
L83:
	;
	goto L79
L84:
	;
	if v337 < v141 {
		v154 = v331
		v157 = v157 + int32(8)
		v158 = v342
		v160 = v337
		goto L38
	} else {
		goto L85
	}
L85:
	;
	goto L39
L86:
	;
	if v363 == int32(-1) {
		v412 = v96
		goto L8
	} else {
		goto L89
	}
L87:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v367)+4))
	if v370 == int32(0) {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v367)+4)) = int32(0)
	goto L86
L89:
	;
	v377 = v96 + v363
	if v377|base.B2i32(v93 < int32(0)) != 0 {
		v399 = v97
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v377 == int32(0) {
		v41 = v93
		v42 = v94
		v52 = v399
		goto L6
	} else {
		goto L93
	}
L91:
	;
	F___clock_gettime(m, int32(1), v18+int32(16))
	mBase = m.M
	v385 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+24)))
	v386 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	v395 = v93 - base.I32_trunc_sat_f64_s(base.F64_div(base.F64_convert_i64_s(v385+(v386*int64(1000000000)+v32)), float64(1e+06)))
	if int32(0) < v395 {
		v399 = v395
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v412 = int32(0)
	goto L8
L93:
	;
	v412 = v377
	goto L8
L94:
	;
	F_errmsg_internal(m, int32(_a_F_WaitEventSetWait_7), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L28
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_WaitEventSetWait_2), int32(1980), int32(_a_F_WaitEventSetWait_5))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L28
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
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
