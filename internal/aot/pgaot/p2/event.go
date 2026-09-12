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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[388]))
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v83 int32
	_ = v83
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
	return v83
L2:
	;
	return int32(0)
L3:
	;
	if v10 == int32(0) {
		v83 = v5
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
		v83 = v5
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
	v42 = *(*int32)(unsafe.Add(mBase, _consts[387]))
	if v42 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if v62 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L12:
	;
	v64 = v27 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if v64 < v65 {
		v27 = v64
		v33 = v62
		goto L10
	} else {
		goto L25
	}
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v53 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	if v40&int32(255) != int32(79) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v40&int32(255) == int32(82) {
		v62 = v33
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v62 = v33
	goto L12
L18:
	;
	goto L13
L19:
	;
	v54 = F_bms_is_member(m, v21, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v59 = F_lappend_oid(m, v33, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L24
	}
L22:
	;
	if v54 == int32(0) {
		v62 = v33
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v62 = v59
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
	v83 = v62
	goto L1
}
func F_EventTriggerUndoInhibitCommandCollection(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _consts[388]))
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int64
	_ = v28
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
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
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v383 int64
	_ = v383
	var v384 int64
	_ = v384
	var v391 float64
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v415 int32
	_ = v415
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	v17 = m.G0
	v19 = v17 - int32(1040)
	m.G0 = v19
	if l1 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = int64(0)
	v34 = int32(-1)
	goto L3
L2:
	;
	F___clock_gettime(m, int32(1), v19+int32(16))
	mBase = m.M
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v31 = int64(*(*int32)(unsafe.Add(mBase, uint32(v19)+24)))
	v33 = v28*int64(-1000000000) - v31
	v34 = l1
	goto L3
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = l4
	*(*int32)(unsafe.Add(mBase, _consts[772])) = int32(1)
	v42 = l1
	v43 = l2
	v53 = v34
	goto L6
L4:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L24
	} else {
		goto L93
	}
L5:
	;
	v434 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[772])) = v434
	F_errstart_cold(m, int32(21), v434)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L24
	} else {
		goto L90
	}
L6:
	;
	v57 = int32(0)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v58 == v57 {
		v95 = v42
		v96 = v43
		v98 = v57
		v99 = v53
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v423 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[772])) = v423
	v426 = *(*int32)(unsafe.Add(mBase, _consts[192]))
	*(*int32)(unsafe.Add(mBase, uint32(v426))) = v423
	m.G0 = v19 + int32(1040)
	return v415
L8:
	;
	goto L7
L9:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v102 = F_pgmem_poll(m, v100, v101, v99)
	mBase = m.M
	if v102 < int32(0) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v61 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v66 == int32(0) {
		v95 = v42
		v96 = v43
		v98 = v57
		v99 = v53
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = int32(-1)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v75<<(uint(int32(4))%32))+12))
	v82 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v81
	v86 = int32(0)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v86
	if l3 == v82 {
		v415 = v82
		goto L8
	} else {
		goto L16
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v69 == int32(0) {
		v95 = v42
		v96 = v43
		v98 = v57
		v99 = v53
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v95 = int32(0)
	v96 = v43 + int32(16)
	v98 = v82
	v99 = v86
	goto L9
L17:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v366 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L18:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v106 == int32(27) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if v102 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v360 = int32(0)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[772])) = v111
	F_errstart_cold(m, int32(21), v111)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
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
	v120 = m.ExcPending
	if v120 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(310225)
	F_errmsg(m, int32(301833), v19)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(504279), int32(1492), int32(323600))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
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
	v360 = int32(-1)
	goto L17
L30:
	;
	goto L31
L31:
	;
	v134 = int32(0)
	v135 = l3 - v98
	if v135 <= v134 {
		v360 = v134
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v138+v139<<(uint(int32(4))%32)) <= base.Ui32(v138) {
		v360 = v134
		goto L17
	} else {
		goto L33
	}
L33:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v150 = v96
	v152 = v138
	v153 = v144
	v155 = v134
	goto L34
L34:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+6)))
	if v161 == int32(0) {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v360 = v333
	goto L17
L36:
	;
	v340 = v152 + int32(16)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v341+v342<<(uint(int32(4))%32)) <= base.Ui32(v340) {
		v360 = v333
		goto L17
	} else {
		goto L75
	}
L37:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+12)) = v166
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	switch v170 - int32(1) {
	case 0:
		goto L41
	default:
		goto L39
	case 15:
		goto L40
	}
L38:
	;
	v328 = v150 + int32(16)
	v333 = v155 + int32(1)
	goto L36
L39:
	;
	if v170&int32(134) == int32(0) {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L63
	}
L40:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+6)))
	if v240&int32(57) == int32(0) {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L59
	}
L41:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+6)))
	if v173&int32(57) == int32(0) {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[801]))
	goto L43
L43:
	;
	v199 = F_read(m, v179, v19+int32(16), int32(1024))
	mBase = m.M
	if v199 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v229 == int32(0) {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L56
	}
L45:
	;
	goto L44
L46:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[158]))
	if v203 == int32(27) {
		goto L43
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v199 == int32(0) {
		goto L5
	} else {
		goto L54
	}
L49:
	;
	if v203 == int32(6) {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v209 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[772])) = v209
	F_errstart_cold(m, int32(21), v209)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	F_errmsg_internal(m, int32(301562), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L24
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(504279), int32(1970), int32(283226))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L24
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	if base.Ui32(int32(1023)) < base.Ui32(v199) {
		goto L43
	} else {
		goto L55
	}
L55:
	;
	goto L45
L56:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v229)+4))
	if v232 == int32(0) {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L57
	}
L57:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	if v235 == int32(0) {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L58
	}
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v150)+4)) = int64(-4294967295)
	goto L38
L59:
	;
	v245 = F_PostmasterIsAliveInternal(m)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L24
	} else {
		goto L60
	}
L60:
	;
	if v245 != 0 {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L61
	}
L61:
	;
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v247 == int32(1) {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v150)+4)) = int64(-4294967280)
	goto L38
L63:
	;
	v256 = int32(0)
	if v170&int32(2) == v256 {
		v270 = v170
		v271 = v256
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v270&int32(4) == int32(0) {
		v285 = v270
		v286 = v271
		goto L67
	} else {
		goto L68
	}
L65:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+6)))
	if v261&int32(57) == int32(0) {
		v270 = v170
		v271 = v256
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v266 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v266
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v270 = v269
	v271 = v266
	goto L64
L67:
	;
	if v285&int32(128) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L68:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+6)))
	if v276&int32(60) == int32(0) {
		v285 = v270
		v286 = v271
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v282 = v271 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v282
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	v285 = v284
	v286 = v282
	goto L67
L70:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v152)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v301
	goto L38
L71:
	;
	if v286 == int32(0) {
		v328 = v150
		v333 = v155
		goto L36
	} else {
		goto L74
	}
L72:
	;
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+6)))
	if v291&int32(8248) == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+4)) = v286 | int32(128)
	goto L70
L74:
	;
	goto L70
L75:
	;
	if v333 < v135 {
		v150 = v328
		v152 = v340
		v153 = v153 + int32(8)
		v155 = v333
		goto L34
	} else {
		goto L76
	}
L76:
	;
	goto L35
L77:
	;
	if v360 == int32(-1) {
		v415 = v98
		goto L8
	} else {
		goto L80
	}
L78:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v366)+4))
	if v369 == int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v366)+4)) = int32(0)
	goto L77
L80:
	;
	v376 = v98 + v360
	if v376 != 0 {
		v402 = v99
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if v376 == int32(0) {
		v42 = v95
		v43 = v96
		v53 = v402
		goto L6
	} else {
		goto L89
	}
L82:
	;
	if v95 < int32(0) {
		v402 = v99
		goto L81
	} else {
		goto L83
	}
L83:
	;
	F___clock_gettime(m, int32(1), v19+int32(16))
	mBase = m.M
	v383 = int64(*(*int32)(unsafe.Add(mBase, uint32(v19)+24)))
	v384 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v391 = base.F64_div(base.F64_convert_i64_s(v383+(v384*int64(1000000000)+v33)), float64(1e+06))
	if base.F64_lt(base.F64_abs(v391), float64(2.147483648e+09)) != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v398 = v95 - v397
	if int32(0) < v398 {
		v402 = v398
		goto L81
	} else {
		goto L88
	}
L85:
	;
	v395 = base.I32_trunc_f64_s(v391)
	v397 = v395
	goto L84
L86:
	;
	goto L87
L87:
	;
	v397 = int32(-2147483648)
	goto L84
L88:
	;
	v415 = int32(0)
	goto L8
L89:
	;
	v415 = v376
	goto L8
L90:
	;
	F_errmsg_internal(m, int32(380184), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L24
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(504279), int32(1980), int32(283226))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(228238)
			F_errmsg(m, int32(196530), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(505570), int32(367), int32(68390))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
