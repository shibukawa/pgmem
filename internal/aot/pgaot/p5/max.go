package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_set_max_safe_fds(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int64
	_ = v52
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	v1 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _consts[647]))
	v16 = F_palloc(m, int32(4096))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = v9 + int32(-16)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v96 = int32(1)
	if v14 <= v96 {
		goto L29
	} else {
		goto L30
	}
L4:
	;
	if v74 == int32(0) {
		goto L3
	} else {
		goto L24
	}
L5:
	;
	v28 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v28
	goto L7
L6:
	;
	goto L7
L7:
	;
	v32 = int32(0)
	v33 = F___syscall_ret(m, v32)
	mBase = m.M
	if v33 == v32 {
		v74 = int32(0)
		goto L8
	} else {
		goto L9
	}
L8:
	;
	m.G0 = v26 + int32(16)
	goto L4
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	if v37 != int32(52) {
		v74 = v33
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v41 = v26 + int32(8)
	v42 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v41))) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = v42
	v46 = int32(0)
	v47 = F___syscall_ret(m, v46)
	mBase = m.M
	if v47 < v46 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v74 = int32(-1)
	goto L8
L12:
	;
	goto L13
L13:
	;
	v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v26)+8)))
	if v52 == int64(4294967295) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v55 = int64(-1)
	goto L16
L15:
	;
	v55 = v52
	goto L16
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	if v58 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v62 = int64(-1)
	goto L19
L18:
	;
	v62 = base.I64_extend_i32_u(v58)
	goto L19
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v62
	if v52 == int64(4294967295) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(-1)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v68 = int32(0)
	if v58 != int32(-1) {
		v74 = v68
		goto L8
	} else {
		goto L23
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(-1)
	v74 = v68
	goto L8
L24:
	;
	v83 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v83 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(307186), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(520480), int32(988), int32(182602))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L3
L29:
	;
	v99 = v96
	goto L31
L30:
	;
	v99 = v14
	goto L31
L31:
	;
	v103 = v1
	v104 = v16
	v105 = int32(1024)
	v106 = v1
	goto L34
L32:
	;
	F_pfree(m, v193)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L64
	}
L33:
	;
	v174 = int32(0)
	goto L61
L34:
	;
	if v74 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v169 = v99
	v170 = v157
	v172 = v164
	goto L33
L36:
	;
	if v105 <= v103 {
		goto L53
	} else {
		goto L54
	}
L37:
	;
	if v103 != 0 {
		v169 = v103
		v170 = v104
		v172 = v106
		goto L33
	} else {
		goto L52
	}
L38:
	;
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
	if base.Ui64(v112-int64(1)) <= base.Ui64(base.I64_extend_i32_u(v106)) {
		goto L37
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v117 = m.Env.X__syscall_dup(m, int32(2))
	mBase = m.M
	if base.Ui32(int32(-4095)) <= base.Ui32(v117) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L40
L42:
	;
	if int32(0) <= v125 {
		goto L36
	} else {
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, _consts[87])) = int32(0) - v117
	v125 = int32(-1)
	goto L45
L44:
	;
	v125 = v117
	goto L45
L45:
	;
	goto L42
L46:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _consts[87]))
	switch v129 - int32(33) {
	case 0, 8:
		goto L37
	default:
		goto L47
	}
L47:
	;
	v134 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v134 == int32(0) {
		goto L37
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v103
	F_errmsg_internal(m, int32(305125), v9+int32(-32))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(520480), int32(1011), int32(182602))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	goto L37
L52:
	;
	v192 = v103
	v193 = v104
	v195 = v106
	goto L32
L53:
	;
	v153 = F_repalloc(m, v104, v105<<(uint(int32(3))%32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	v157 = v104
	v158 = v105
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157+v103<<(uint(int32(2))%32)))) = v125
	if v125 < v106 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v157 = v153
	v158 = v105 << (uint(int32(1)) % 32)
	goto L55
L57:
	;
	v164 = v106
	goto L59
L58:
	;
	v164 = v125
	goto L59
L59:
	;
	v166 = v103 + int32(1)
	if v166 != v99 {
		v103 = v166
		v104 = v157
		v105 = v158
		v106 = v164
		goto L34
	} else {
		goto L60
	}
L60:
	;
	goto L35
L61:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v170+v174<<(uint(int32(2))%32))))
	v186 = F_close(m, v185)
	mBase = m.M
	v188 = v174 + int32(1)
	if v188 != v169 {
		v174 = v188
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v192 = v169
	v193 = v170
	v195 = v172
	goto L32
L63:
	;
	goto L62
L64:
	;
	v202 = *(*int32)(unsafe.Add(mBase, _consts[647]))
	if v192 < v202 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v204 = v192
	goto L67
L66:
	;
	v204 = v202
	goto L67
L67:
	;
	v206 = v204 - int32(10)
	*(*int32)(unsafe.Add(mBase, _consts[648])) = v206
	v210 = v195 - v192 + int32(1)
	if int32(47) < v206 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v215 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L77
	}
L71:
	;
	if v215 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v192
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v210
	v220 = *(*int32)(unsafe.Add(mBase, _consts[648]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v220
	F_errmsg_internal(m, int32(498980), v9+int32(-48))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	m.G0 = v11 - int32(-64)
	return
L75:
	;
	F_errfinish(m, int32(520480), int32(1086), int32(182619))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errmsg(m, int32(136203), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(58)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v210
	v250 = *(*int32)(unsafe.Add(mBase, _consts[648]))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v250 + int32(10)
	F_errdetail(m, int32(642983), v11)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(520480), int32(1083), int32(182619))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
