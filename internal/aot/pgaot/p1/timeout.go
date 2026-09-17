package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_enable_timeout(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var __phi97 int32
	_ = __phi97
	var v98 int32
	_ = v98
	var __phi98 int32
	_ = __phi98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var __phi210 int32
	_ = __phi210
	var v213 int32
	_ = v213
	var __phi213 int32
	_ = __phi213
	var v214 int32
	_ = v214
	var __phi214 int32
	_ = __phi214
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = l0 * int32(40)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_enable_timeout[0]))))
	if v18 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	if v23 <= int32(0) {
		v62 = int32(-1)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	v147 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	if v149 <= v147 {
		v181 = v147
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_enable_timeout[2])))
	v86 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)) = uint8(v86)
	v89 = v30 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	if v89 < v91 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v30 = int32(0)
	goto L8
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	if v30 < v49 {
		goto L4
	} else {
		goto L12
	}
L8:
	;
	v38 = v30 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+uint32(_c_F_enable_timeout[2])))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 == l0 {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v62 = int32(-1)
	goto L5
L10:
	;
	v43 = v30 + int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	if v43 < v45 {
		v30 = v43
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v62 = v30
	goto L5
L13:
	;
	return
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v62
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v69 - int32(1)
	F_errmsg_internal(m, int32(_a_F_enable_timeout_0), v14+int32(16))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_enable_timeout_1), int32(143), int32(_a_F_enable_timeout_2))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	__phi97 = v30
	__phi98 = v89
	v97 = __phi97
	v98 = __phi98
	goto L20
L18:
	;
	goto L19
L19:
	;
	v128 = int32(_a_F_enable_timeout_3)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1])) = v130 - int32(1)
	goto L3
L20:
	;
	v104 = int32(2)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v98<<(uint(v104)%32))+uint32(_c_F_enable_timeout[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v97<<(uint(v104)%32))+uint32(_c_F_enable_timeout[2]))) = v110
	v113 = v98 + int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	if v113 < v115 {
		__phi97 = v98
		__phi98 = v113
		v97 = __phi97
		v98 = __phi98
		goto L20
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_enable_timeout[3]))) = l3
	*(*int64)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_enable_timeout[4]))) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_enable_timeout[5]))) = l1
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_enable_timeout[6]))) = uint8(v190)
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	if v181 <= v193 {
		goto L33
	} else {
		goto L34
	}
L24:
	;
	v157 = v147
	goto L25
L25:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v157<<(uint(int32(2))%32))+uint32(_c_F_enable_timeout[2])))
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v165)+24))
	if l2 < v166 {
		v181 = v157
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v181 = v172
	goto L23
L27:
	;
	if l2 == v166 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if base.Ui32(l0) < base.Ui32(v169) {
		v181 = v157
		goto L23
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v172 = v157 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	if v172 < v174 {
		v157 = v172
		goto L25
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	goto L26
L33:
	;
	v195 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_enable_timeout[0]))) = uint8(v195)
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	v201 = v199 - v195
	if v201 < v181 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L13
	} else {
		goto L51
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v181<<(uint(int32(2))%32))+uint32(_c_F_enable_timeout[2]))) = v17 + int32(_a_F_enable_timeout_4)
	v299 = int32(_a_F_enable_timeout_3)
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1])) = v301 + int32(1)
	m.G0 = v14 + int32(32)
	return
L37:
	;
	if base.Ui32(v181) < base.Ui32(v201) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v204 = v181
	goto L40
L39:
	;
	v204 = v201
	goto L40
L40:
	;
	v205 = v199 - v204
	v209 = v205 & int32(3)
	if v209 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	__phi210 = v201
	__phi213 = int32(0)
	__phi214 = v199
	v210 = __phi210
	v213 = __phi213
	v214 = __phi214
	goto L44
L42:
	;
	v234 = v201
	v238 = v199
	goto L43
L43:
	;
	if base.Ui32(v205-int32(1)) < base.Ui32(int32(3)) {
		goto L36
	} else {
		goto L47
	}
L44:
	;
	v221 = int32(2)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v210<<(uint(v221)%32))+uint32(_c_F_enable_timeout[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v214<<(uint(v221)%32))+uint32(_c_F_enable_timeout[2]))) = v227
	v229 = int32(1)
	v230 = v210 - v229
	v232 = v213 + v229
	if v232 != v209 {
		__phi210 = v230
		__phi213 = v232
		__phi214 = v210
		v210 = __phi210
		v213 = __phi213
		v214 = __phi214
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v234 = v230
	v238 = v210
	goto L43
L46:
	;
	goto L45
L47:
	;
	v247 = v234
	v251 = v238
	goto L48
L48:
	;
	v258 = int32(2)
	v263 = v247 << (uint(v258) % 32)
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_enable_timeout[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v251<<(uint(v258)%32))+uint32(_c_F_enable_timeout[2]))) = v266
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_enable_timeout[7])))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_enable_timeout[2]))) = v270
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_enable_timeout[8])))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_enable_timeout[7]))) = v274
	v277 = v247 - int32(3)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277<<(uint(v258)%32))+uint32(_c_F_enable_timeout[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v263)+uint32(_c_F_enable_timeout[8]))) = v280
	if base.Ui32(v181) < base.Ui32(v277) {
		v247 = v247 - int32(4)
		v251 = v277
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L36
L50:
	;
	goto L49
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v181
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_enable_timeout[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v314
	F_errmsg_internal(m, int32(_a_F_enable_timeout_0), v14)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L13
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_enable_timeout_1), int32(120), int32(_a_F_enable_timeout_5))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L13
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_timeout_active(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0*int32(40))+uint32(_c_F_get_timeout_active[0]))))
	return v4
}
