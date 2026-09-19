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
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int64
	_ = v71
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v160 int64
	_ = v160
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v177 int32
	_ = v177
	var v181 int64
	_ = v181
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v197 int64
	_ = v197
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v218 int64
	_ = v218
	var v222 int64
	_ = v222
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int64
	_ = v255
	var v260 int32
	_ = v260
	var v265 int64
	_ = v265
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
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
	v314 = m.ExcPending
	if v314 != 0 {
		goto L5
	} else {
		goto L73
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
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[0]))
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
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)))
	if v296 != v292&int32(_a_F_WaitReadBuffers_0) {
		goto L65
	} else {
		goto L66
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
	v292 = v55
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
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v242 = int32(base.Ui32(v238)>>(uint(int32(6))%32)) & int32(7)
	if v242 == int32(4) {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[1]))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v65 = v61 + v62<<(uint(int32(7))%32)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(v65)+48))
	v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+8)))
	v71 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v27)+4)))
	if base.B2i32(v66 == int32(0))|base.B2i32(v69 != v70|v71<<(uint(int64(32))%64)) != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v89 != 0 {
		goto L20
	} else {
		goto L28
	}
L23:
	;
	v89 = int32(1)
	goto L22
L24:
	;
	if v66&int32(254) != int32(6) {
		v89 = int32(0)
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[2]))
	if v82 != v84 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	F_pgaio_io_reclaim(m, v65)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[3])))
	v94 = m.G0
	v96 = v94 - int32(16)
	m.G0 = v96
	if v91 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_pgaio_wref_wait(m, v27)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L33
	}
L30:
	;
	F___clock_gettime(m, int32(1), v96)
	mBase = m.M
	v100 = int64(*(*int32)(unsafe.Add(mBase, uint32(v96)+8)))
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v96)))
	v105 = v100 + v101*int64(1000000000)
	goto L32
L31:
	;
	v105 = int64(0)
	goto L32
L32:
	;
	m.G0 = v96 + int32(16)
	goto L29
L33:
	;
	v112 = int32(0)
	v113 = int64(0)
	v117 = m.G0
	v119 = v117 - int32(16)
	m.G0 = v119
	if v105 != v113 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L20
L35:
	;
	F___clock_gettime(m, int32(1), v119)
	mBase = m.M
	v125 = int64(*(*int32)(unsafe.Add(mBase, uint32(v119)+8)))
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v119)))
	v130 = v125 + (v126*int64(1000000000) - v105)
	if v24 == int32(2) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	v214 = v24*int32(320) + v25<<(uint(int32(6))%32)
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v214)+uint32(_c_F_WaitReadBuffers[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v214)+uint32(_c_F_WaitReadBuffers[4]))) = v218 + base.I64_extend_i32_u(v112)
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v214)+uint32(_c_F_WaitReadBuffers[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v214)+uint32(_c_F_WaitReadBuffers[5]))) = v222 + v113
	F_pgstat_count_backend_io_op(m, v24, v25, int32(6), v112, v113)
	mBase = m.M
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[6])) = uint8(v227)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[7])) = uint8(v227)
	m.G0 = v119 + int32(16)
	goto L34
L38:
	;
	v177 = v24*int32(320) + v25<<(uint(int32(6))%32)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_WaitReadBuffers[8])))
	*(*int64)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_WaitReadBuffers[8]))) = v181 + v130
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[9]))
	v192 = int32(0)
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v185))|base.B2i32(int32(1)<<(uint(v185)%32)&int32(_a_F_WaitReadBuffers_1) == v192) == v192 {
		goto L48
	} else {
		goto L49
	}
L39:
	;
	goto L41
L41:
	;
	goto L42
L42:
	;
	goto L45
L45:
	;
	v156 = int32(_a_F_WaitReadBuffers_2)
	v158 = *(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[10]))
	v160 = base.I64_div_s(v130, int64(1000))
	*(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[10])) = v158 + v160
	switch v24 {
	case 0:
		goto L47
	case 1:
		goto L46
	default:
		goto L38
	}
L46:
	;
	v168 = int32(_a_F_WaitReadBuffers_3)
	v170 = *(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[11]))
	*(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[11])) = v170 + v130
	goto L38
L47:
	;
	v163 = int32(_a_F_WaitReadBuffers_4)
	v165 = *(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[12]))
	*(*int64)(unsafe.Add(mBase, _c_F_WaitReadBuffers[12])) = v165 + v130
	goto L38
L48:
	;
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_WaitReadBuffers[13])))
	*(*int64)(unsafe.Add(mBase, uint32(v177)+uint32(_c_F_WaitReadBuffers[13]))) = v197 + v130
	v201 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[6])) = uint8(v201)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitReadBuffers[14])) = uint8(v201)
	goto L50
L49:
	;
	goto L50
L50:
	;
	goto L37
L51:
	;
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)))
	v290 = v289 + v287
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+34)) = uint16(v290)
	v292 = v290
	goto L15
L52:
	;
	if v238&int32(448) != int32(128) {
		v287 = v247
		goto L51
	} else {
		goto L59
	}
L53:
	;
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v255
	F_pgaio_result_report(m, v14+int32(8), v38, v254)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L5
	} else {
		goto L58
	}
L54:
	;
	v253 = int32(0)
	v254 = int32(21)
	goto L53
L55:
	;
	goto L56
L56:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.Ui32(int32(1)) < base.Ui32(v242-int32(3)) {
		goto L52
	} else {
		goto L57
	}
L57:
	;
	v253 = v247
	v254 = int32(19)
	goto L53
L58:
	;
	v287 = v253
	goto L51
L59:
	;
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v265
	F_pgaio_result_report(m, v14+int32(16), v38, int32(14))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v274 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	if v274 == int32(0) {
		v287 = v247
		goto L51
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_WaitReadBuffers_5), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_WaitReadBuffers_6), int32(1620), int32(_a_F_WaitReadBuffers_7))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L64
	}
L64:
	;
	v287 = v247
	goto L51
L65:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_WaitReadBuffers[15]))
	if v301 != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	goto L14
L68:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L5
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v306 = F_AsyncReadBuffers(m, l0, v14+int32(28))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	goto L13
L73:
	;
	F_errmsg_internal(m, int32(_a_F_WaitReadBuffers_8), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_WaitReadBuffers_6), int32(1661), int32(_a_F_WaitReadBuffers_9))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
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
		*(*int32)(unsafe.Add(mBase, _c_F___wasi_syscall_ret[0])) = l0
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v9 = m.G0
	v11 = v9 - int32(2064)
	m.G0 = v11
	v13 = F_AllocateDir(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = F_ReadDirExtended(m, v13, l0, l3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = v15
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_FreeDir(m, v13)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L27
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_walkdir[0]))
	if v26 != 0 {
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
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+19)))
	if v29 != int32(46) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v66 = F_ReadDirExtended(m, v13, l0, l3)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L25
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v22 + int32(19)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	v46 = v11 + int32(16)
	v49 = F_pg_snprintf(m, v46, int32(2048), int32(_a_F_walkdir_0), v11)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L19
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)))
	if v32 == int32(0) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)))
	if v35 != int32(46) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+21)))
	if v38 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L14
L19:
	;
	v51 = F_get_dirent_type(m, v46, v22, l2, l3)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	F_walkdir(m, v11+int32(16), l1, int32(0), l3)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	m.T0[l1].(func(*base.Module, int32, int32, int32))(m, v11+int32(16), int32(0), l3)
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
	if v66 != 0 {
		v22 = v66
		goto L7
	} else {
		goto L26
	}
L26:
	;
	goto L8
L27:
	;
	if v13 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	m.T0[l1].(func(*base.Module, int32, int32, int32))(m, l0, int32(1), l3)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	m.G0 = v11 + int32(2064)
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	v2 = l1
	if l0 != 0 {
		if base.Ui32(v2) <= base.Ui32(int32(127)) {
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
			return int32(1)
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _c_F_wcrtomb[0]))
			v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			if v7 == int32(0) {
				if v2&int32(-128) == int32(_a_F_wcrtomb_0) {
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
					return int32(1)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_wcrtomb[1])) = int32(25)
					v91 = int32(-1)
					return v91
				}
			} else {
				if base.Ui32(v2) <= base.Ui32(int32(2047)) {
					v19 = v2&int32(63) | int32(128)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v19)
					v24 = int32(base.Ui32(v2)>>(uint(int32(6))%32)) | int32(192)
					*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v24)
					return int32(2)
				} else {
					if base.B2i32(v2&int32(-8192) != int32(_a_F_wcrtomb_1))&base.B2i32(base.Ui32(int32(_a_F_wcrtomb_2)) <= base.Ui32(v2)) == int32(0) {
						v37 = int32(63)
						v39 = int32(128)
						v40 = v2&v37 | v39
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v40)
						v45 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v45)
						v52 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v37 | v39
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v52)
						return int32(3)
					} else {
						if base.Ui32(v2-int32(_a_F_wcrtomb_3)) <= base.Ui32(int32(_a_F_wcrtomb_4)) {
							v60 = int32(63)
							v62 = int32(128)
							v63 = v2&v60 | v62
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v63)
							v68 = int32(base.Ui32(v2)>>(uint(int32(18))%32)) | int32(240)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v68)
							v75 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v60 | v62
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v75)
							v82 = int32(base.Ui32(v2)>>(uint(int32(12))%32))&v60 | v62
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v82)
							return int32(4)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_wcrtomb[1])) = int32(25)
							v91 = int32(-1)
							return v91
						}
					}
				}
			}
		}
	} else {
		v91 = int32(1)
		return v91
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
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v2 = l1
	if l0 == int32(0) {
		return int32(0)
	} else {
		if l0 != 0 {
			if base.Ui32(v2) <= base.Ui32(int32(127)) {
				*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
				v95 = int32(1)
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, _c_F_wctomb[0]))
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				if v11 == int32(0) {
					if v2&int32(-128) == int32(_a_F_wctomb_0) {
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v2)
						v95 = int32(1)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_wctomb[1])) = int32(25)
						v92 = int32(-1)
						v95 = v92
					}
				} else {
					if base.Ui32(v2) <= base.Ui32(int32(2047)) {
						v23 = v2&int32(63) | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v23)
						v28 = int32(base.Ui32(v2)>>(uint(int32(6))%32)) | int32(192)
						*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v28)
						v95 = int32(2)
					} else {
						if base.B2i32(v2&int32(-8192) != int32(_a_F_wctomb_1))&base.B2i32(base.Ui32(int32(_a_F_wctomb_2)) <= base.Ui32(v2)) == int32(0) {
							v40 = int32(63)
							v42 = int32(128)
							v43 = v2&v40 | v42
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v43)
							v48 = int32(base.Ui32(v2)>>(uint(int32(12))%32)) | int32(224)
							*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v48)
							v55 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v40 | v42
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v55)
							v95 = int32(3)
						} else {
							if base.Ui32(v2-int32(_a_F_wctomb_3)) <= base.Ui32(int32(_a_F_wctomb_4)) {
								v62 = int32(63)
								v64 = int32(128)
								v65 = v2&v62 | v64
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v65)
								v70 = int32(base.Ui32(v2)>>(uint(int32(18))%32)) | int32(240)
								*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v70)
								v77 = int32(base.Ui32(v2)>>(uint(int32(6))%32))&v62 | v64
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v77)
								v84 = int32(base.Ui32(v2)>>(uint(int32(12))%32))&v62 | v64
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v84)
								v95 = int32(4)
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_wctomb[1])) = int32(25)
								v92 = int32(-1)
								v95 = v92
							}
						}
					}
				}
			}
		} else {
			v92 = int32(1)
			v95 = v92
		}
		return v95
	}
}
func F_win1250_to_mic(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13951(m, l0, int32(_a_F_win1250_to_mic_0), int32(29), int32(130))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_win1251_to_win866(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13958(m, l0, int32(_a_F_win1251_to_win866_0), int32(20), int32(23))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_win866_to_iso(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13958(m, l0, int32(_a_F_win866_to_iso_0), int32(25), int32(20))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
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
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[0]))
		if v9 != 0 {
			v21 = v9
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[1]))
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[2]))
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
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[1]))
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_write_syslogger_file[2]))
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
			F_write_stderr(m, int32(_a_F_write_syslogger_file_0), int32(0))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
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
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v12 != 0 {
		v25 = v4
		v26 = v4
		v27 = int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v26 + v27
		v31 = v9 + int32(12)
		F_LogicalTapeWrite(m, l1, v31, v27)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			F_LogicalTapeWrite(m, l1, v25, v26)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v37&int32(1) != 0 {
					F_LogicalTapeWrite(m, l1, v31, int32(4))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				} else {
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
		if v13 == int32(0) {
			v16 = int32(4)
			v25 = l2 + v16
			v26 = v16
			v27 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v26 + v27
			v31 = v9 + int32(12)
			F_LogicalTapeWrite(m, l1, v31, v27)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_LogicalTapeWrite(m, l1, v25, v26)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
					if v37&int32(1) != 0 {
						F_LogicalTapeWrite(m, l1, v31, int32(4))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							m.G0 = v9 + int32(16)
							return
						}
					} else {
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			v23 = F_datumGetSize(m, v19, int32(0), v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = v19
				v26 = v23
				v27 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v26 + v27
				v31 = v9 + int32(12)
				F_LogicalTapeWrite(m, l1, v31, v27)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					F_LogicalTapeWrite(m, l1, v25, v26)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
						if v37&int32(1) != 0 {
							F_LogicalTapeWrite(m, l1, v31, int32(4))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						} else {
							m.G0 = v9 + int32(16)
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v10 - int32(6)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v16 = v8 + int32(12)
	F_BufFileWrite(m, v14, v16, int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v21 = int32(10)
		F_BufFileWrite(m, v20, l1+v21, v10-v21)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
			if v27 == int32(1) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				F_BufFileWrite(m, v30, v16, int32(4))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					v34 = F_GetMemoryChunkSpace(m, l1)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v36 + base.I64_extend_i32_u(v34)
						F_pfree(m, l1)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				v34 = F_GetMemoryChunkSpace(m, l1)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v36 + base.I64_extend_i32_u(v34)
					F_pfree(m, l1)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						m.G0 = v8 + int32(16)
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v11 + v12
	v16 = v8 + int32(12)
	F_LogicalTapeWrite(m, l1, v16, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		F_LogicalTapeWrite(m, l1, v10, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v23&int32(1) != 0 {
				F_LogicalTapeWrite(m, l1, v16, int32(4))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
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
