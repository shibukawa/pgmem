package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WaitReadBuffers(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v155 int32
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v167 int32
	_ = v167
	var v169 int64
	_ = v169
	var v176 int32
	_ = v176
	var v182 int64
	_ = v182
	var v186 int32
	_ = v186
	var v199 int32
	_ = v199
	var v205 int64
	_ = v205
	var v209 int32
	_ = v209
	var v221 int32
	_ = v221
	var v227 int64
	_ = v227
	var v233 int64
	_ = v233
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int64
	_ = v266
	var v271 int32
	_ = v271
	var v276 int64
	_ = v276
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
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
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v16 == int32(116) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v27 = l0 + int32(36)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	goto L8
L2:
	;
	v24 = int32(1)
	v25 = int32(3)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = F_IOContextForStrategy(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v24 = int32(0)
	v25 = v22
	goto L1
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L5
	} else {
		goto L74
	}
L8:
	;
	if base.B2i32(v28 != int32(-1)) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[738]))
	if v34 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v36 = l0 + int32(48)
	v38 = l0 + int32(56)
	goto L13
L12:
	;
	goto L11
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	goto L16
L14:
	;
	m.G0 = v14 + int32(32)
	return
L15:
	;
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)))
	if v307 != v303&int32(65535) {
		goto L66
	} else {
		goto L67
	}
L16:
	;
	if base.B2i32(v50 != int32(-1)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)))
	v303 = v55
	goto L15
L18:
	;
	goto L19
L19:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	if v56&int32(448) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v253 = int32(base.Ui32(v249)>>(uint(int32(6))%32)) & int32(7)
	if v253 == int32(4) {
		goto L55
	} else {
		goto L56
	}
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[746]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v65 = v61 + v62<<(uint(int32(7))%32)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v65)+48))
	v68 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+8)))
	v69 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+4)))
	if v67 != v68|v69<<(uint(int64(32))%64) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v88 != 0 {
		goto L20
	} else {
		goto L29
	}
L23:
	;
	v88 = int32(1)
	goto L22
L24:
	;
	if v66 == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	if v66&int32(254) != int32(6) {
		v88 = int32(0)
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v83 = *(*int32)(unsafe.Add(mBase, _consts[747]))
	if v81 != v83 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_pgaio_io_reclaim(m, v65)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	goto L23
L29:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, _consts[85])))
	v93 = m.G0
	v95 = v93 - int32(16)
	m.G0 = v95
	if v90 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	F_pgaio_wref_wait(m, v27)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L34
	}
L31:
	;
	F___clock_gettime(m, int32(1), v95)
	mBase = m.M
	v99 = int64(*(*int32)(unsafe.Add(mBase, uint32(v95)+8)))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
	v104 = v99 + v100*int64(1000000000)
	goto L33
L32:
	;
	v104 = int64(0)
	goto L33
L33:
	;
	m.G0 = v95 + int32(16)
	goto L30
L34:
	;
	v111 = int32(0)
	v112 = int64(0)
	v116 = m.G0
	v118 = v116 - int32(16)
	m.G0 = v118
	if v104 != v112 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L20
L36:
	;
	F___clock_gettime(m, int32(1), v118)
	mBase = m.M
	v124 = int64(*(*int32)(unsafe.Add(mBase, uint32(v118)+8)))
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v118)))
	v129 = v124 + (v125*int64(1000000000) - v104)
	if v24 == int32(2) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v221 = v24*int32(320) + v25<<(uint(int32(6))%32)
	v227 = *(*int64)(unsafe.Add(mBase, uint32(v221)+uint32(_consts[748])))
	*(*int64)(unsafe.Add(mBase, uint32(v221)+uint32(_consts[748]))) = v227 + base.I64_extend_i32_u(v111)
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v221)+uint32(_consts[749])))
	*(*int64)(unsafe.Add(mBase, uint32(v221)+uint32(_consts[749]))) = v233 + v112
	F_pgstat_count_backend_io_op(m, v24, v25, int32(6), v111, v112)
	mBase = m.M
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[245])) = uint8(v238)
	*(*uint8)(unsafe.Add(mBase, _consts[246])) = uint8(v238)
	m.G0 = v118 + int32(16)
	goto L35
L39:
	;
	v176 = v24*int32(320) + v25<<(uint(int32(6))%32)
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[750])))
	*(*int64)(unsafe.Add(mBase, uint32(v176)+uint32(_consts[750]))) = v182 + v129
	v186 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if base.Ui32(int32(16)) < base.Ui32(v186) {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	goto L42
L42:
	;
	goto L43
L43:
	;
	goto L46
L46:
	;
	v155 = int32(4495208)
	v157 = *(*int64)(unsafe.Add(mBase, _consts[87]))
	v159 = base.I64_div_s(v129, int64(1000))
	*(*int64)(unsafe.Add(mBase, _consts[87])) = v157 + v159
	switch v24 {
	case 0:
		goto L48
	case 1:
		goto L47
	default:
		goto L39
	}
L47:
	;
	v167 = int32(4413848)
	v169 = *(*int64)(unsafe.Add(mBase, _consts[116]))
	*(*int64)(unsafe.Add(mBase, _consts[116])) = v169 + v129
	goto L39
L48:
	;
	v162 = int32(4413832)
	v164 = *(*int64)(unsafe.Add(mBase, _consts[114]))
	*(*int64)(unsafe.Add(mBase, _consts[114])) = v164 + v129
	goto L39
L49:
	;
	goto L38
L50:
	;
	if int32(1)<<(uint(v186)%32)&int32(115186) == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v199 = v24*int32(320) + v25<<(uint(int32(6))%32)
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v199)+uint32(_consts[751])))
	*(*int64)(unsafe.Add(mBase, uint32(v199)+uint32(_consts[751]))) = v205 + v129
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[245])) = uint8(v209)
	*(*uint8)(unsafe.Add(mBase, _consts[249])) = uint8(v209)
	goto L49
L52:
	;
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)))
	v301 = v300 + v298
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)) = uint16(v301)
	v303 = v301
	goto L15
L53:
	;
	if v249&int32(448) != int32(128) {
		v298 = v258
		goto L52
	} else {
		goto L60
	}
L54:
	;
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v266
	F_pgaio_result_report(m, v14+int32(8), v38, v265)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L59
	}
L55:
	;
	v264 = int32(0)
	v265 = int32(21)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.Ui32(int32(1)) < base.Ui32(v253-int32(3)) {
		goto L53
	} else {
		goto L58
	}
L58:
	;
	v264 = v258
	v265 = int32(19)
	goto L54
L59:
	;
	v298 = v264
	goto L52
L60:
	;
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v276
	F_pgaio_result_report(m, v14+int32(16), v38, int32(14))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v285 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	if v285 == int32(0) {
		v298 = v258
		goto L52
	} else {
		goto L63
	}
L63:
	;
	F_errmsg_internal(m, int32(12768), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(495274), int32(1620), int32(98170))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L5
	} else {
		goto L65
	}
L65:
	;
	v298 = v258
	goto L52
L66:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v312 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	goto L14
L69:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v317 = F_AsyncReadBuffers(m, l0, v14+int32(28))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	goto L13
L74:
	;
	F_errmsg_internal(m, int32(464378), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(495274), int32(1661), int32(135437))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L5
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F___wasi_syscall_ret(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	if l0 == int32(0) {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[159])) = l0
		return int32(-1)
	}
}
func F___wasm_longjmp(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = int32(1)
	if base.Ui32(l1) <= base.Ui32(v3) {
		v6 = v3
	} else {
		v6 = l1
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l0
	{
		m.ExcTag = uint32(int32(0))
		m.ExcVals[0] = uint64(uint32(l0 + int32(8)))
		m.ExcPending = 1
	}
	return
}
func F___wasm_setjmp(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l2
	return
}
func F_walkdir(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	v8 = m.G0
	v10 = v8 - int32(2064)
	m.G0 = v10
	v12 = F_AllocateDir(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_ReadDirExtended(m, v12, l0, l3)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_FreeDir(m, v12)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L27
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v24 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+19)))
	if v27 != int32(46) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v65 = F_ReadDirExtended(m, v12, l0, l3)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L25
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v21 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
	v47 = F_pg_snprintf(m, v10+int32(16), int32(2048), int32(177081), v10)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L19
	}
L15:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)))
	if v30 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+20)))
	if v33 != int32(46) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+21)))
	if v36 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	v51 = F_get_dirent_type(m, v10+int32(16), v21, l2, l3)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	F_walkdir(m, v10+int32(16), l1, int32(0), l3)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	m.T0[l1].(func(*base.Module, int32, int32, int32))(m, v10+int32(16), int32(0), l3)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	switch v51 - int32(2) {
	case 0:
		goto L21
	case 1:
		goto L20
	default:
		goto L13
	}
L23:
	;
	goto L13
L24:
	;
	goto L13
L25:
	;
	if v65 != 0 {
		v21 = v65
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L8
L27:
	;
	if v12 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	m.T0[l1].(func(*base.Module, int32, int32, int32))(m, l0, int32(1), l3)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	m.G0 = v10 + int32(2064)
	return
L31:
	;
	goto L30
}
func F_wcrtomb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	v2 = l1
	if l0 != 0 {
		if base.Ui32(v2) <= base.Ui32(int32(127)) {
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
			return int32(1)
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
			v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v7 == int32(0) {
				if v2&int32(-128) == int32(57216) {
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
					return int32(1)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[159])) = int32(25)
					v94 = int32(-1)
					return v94
				}
			} else {
				if base.Ui32(v2) <= base.Ui32(int32(2047)) {
					v22 = v2&int32(63) | int32(128)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v22)
					v27 = int32(base.Ui32(v2)>>(uint(int32(6))%32)) | int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v27)
					return int32(2)
				} else {
					if base.B2i32(v2&int32(-8192) != int32(57344))&base.B2i32(base.Ui32(int32(55296)) <= base.Ui32(v2)) == int32(0) {
						v40 = int32(63)
						v42 = int32(128)
						v43 = v2&v40 | v42
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v43)
						v48 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
						v55 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v40 | v42
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v55)
						return int32(3)
					} else {
						if base.Ui32(v2-int32(65536)) <= base.Ui32(int32(1048575)) {
							v63 = int32(63)
							v65 = int32(128)
							v66 = v2&v63 | v65
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v66)
							v71 = int32(base.Ui32(v2)>>(uint(int32(18))%32)) | int32(240)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v71)
							v78 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v63 | v65
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v78)
							v85 = int32(base.Ui32(v2)>>(uint(int32(12))%32))&v63 | v65
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v85)
							return int32(4)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[159])) = int32(25)
							v94 = int32(-1)
							return v94
						}
					}
				}
			}
		}
	} else {
		v94 = int32(1)
		return v94
	}
}
func F_wctomb(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v2 = l1
	if l0 == int32(0) {
		return int32(0)
	} else {
		if l0 != 0 {
			if base.Ui32(v2) <= base.Ui32(int32(127)) {
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
				v98 = int32(1)
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				if v11 == int32(0) {
					if v2&int32(-128) == int32(57216) {
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
						v98 = int32(1)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[159])) = int32(25)
						v95 = int32(-1)
						v98 = v95
					}
				} else {
					if base.Ui32(v2) <= base.Ui32(int32(2047)) {
						v26 = v2&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v26)
						v31 = int32(base.Ui32(v2)>>(uint(int32(6))%32)) | int32(192)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v31)
						v98 = int32(2)
					} else {
						if base.B2i32(v2&int32(-8192) != int32(57344))&base.B2i32(base.Ui32(int32(55296)) <= base.Ui32(v2)) == int32(0) {
							v43 = int32(63)
							v45 = int32(128)
							v46 = v2&v43 | v45
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v46)
							v51 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v51)
							v58 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v43 | v45
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v58)
							v98 = int32(3)
						} else {
							if base.Ui32(v2-int32(65536)) <= base.Ui32(int32(1048575)) {
								v65 = int32(63)
								v67 = int32(128)
								v68 = v2&v65 | v67
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v68)
								v73 = int32(base.Ui32(v2)>>(uint(int32(18))%32)) | int32(240)
								*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v73)
								v80 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v65 | v67
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v80)
								v87 = int32(base.Ui32(v2)>>(uint(int32(12))%32))&v65 | v67
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v87)
								v98 = int32(4)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[159])) = int32(25)
								v95 = int32(-1)
								v98 = v95
							}
						}
					}
				}
			}
		} else {
			v95 = int32(1)
			v98 = v95
		}
		return v98
	}
}
func F_win1250_to_mic(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(29), int32(7))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_latin2mic_with_table(m, v6, v5, v10, int32(130), int32(29), int32(2234720), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_win1251_to_win866(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(23), int32(20))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(23), int32(20), int32(2230064), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_win866_to_iso(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(20), int32(25))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_local2local(m, v6, v5, v10, int32(20), int32(25), int32(2230704), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_write_syslogger_file(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	if l2&int32(8) != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _consts[586]))
		if v9 != 0 {
			v21 = v9
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[587]))
			v14 = *(*int32)(unsafe.Add(mBase, _consts[588]))
			if v12 != 0 {
				v15 = v12
			} else {
				v15 = v14
			}
			if int32(base.Ui32(l2&int32(16))>>(uint(int32(4))%32)) != 0 {
				v20 = v15
			} else {
				v20 = v14
			}
			v21 = v20
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _consts[587]))
		v14 = *(*int32)(unsafe.Add(mBase, _consts[588]))
		if v12 != 0 {
			v15 = v12
		} else {
			v15 = v14
		}
		if int32(base.Ui32(l2&int32(16))>>(uint(int32(4))%32)) != 0 {
			v20 = v15
		} else {
			v20 = v14
		}
		v21 = v20
	}
	v24 = F_fwrite(m, l0, int32(1), l1, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		if v24 != l1 {
			F_write_stderr(m, int32(748979), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_writetup_datum(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v11 != 0 {
		v24 = v4
		v25 = v4
		v26 = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v25 + v26
		F_LogicalTapeWrite(m, l1, v8+int32(12), v26)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			F_LogicalTapeWrite(m, l1, v24, v25)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v36&int32(1) != 0 {
					F_LogicalTapeWrite(m, l1, v8+int32(12), int32(4))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v12 == int32(0) {
			v15 = int32(4)
			v24 = l2 + v15
			v25 = v15
			v26 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v25 + v26
			F_LogicalTapeWrite(m, l1, v8+int32(12), v26)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				F_LogicalTapeWrite(m, l1, v24, v25)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
					if v36&int32(1) != 0 {
						F_LogicalTapeWrite(m, l1, v8+int32(12), int32(4))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					} else {
						m.G0 = v8 + int32(16)
						return
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			v22 = F_datumGetSize(m, v18, int32(0), v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				v24 = v18
				v25 = v22
				v26 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v25 + v26
				F_LogicalTapeWrite(m, l1, v8+int32(12), v26)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					F_LogicalTapeWrite(m, l1, v24, v25)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
						if v36&int32(1) != 0 {
							F_LogicalTapeWrite(m, l1, v8+int32(12), int32(4))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
func F_writetup_heap_2(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v9 - int32(6)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	F_BufFileWrite(m, v13, v7+int32(12), int32(4))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v20 = int32(10)
		F_BufFileWrite(m, v19, l1+v20, v9-v20)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
			if v26 == int32(1) {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				F_BufFileWrite(m, v29, v7+int32(12), int32(4))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v35 = F_GetMemoryChunkSpace(m, l1)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v37 + base.I64_extend_i32_u(v35)
						F_pfree(m, l1)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			} else {
				v35 = F_GetMemoryChunkSpace(m, l1)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v37 + base.I64_extend_i32_u(v35)
					F_pfree(m, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_writetup_index_gin(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v11 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v10 + v11
	F_LogicalTapeWrite(m, l1, v7+int32(12), v11)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		F_LogicalTapeWrite(m, l1, v9, v19)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v22&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v7+int32(12), int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
