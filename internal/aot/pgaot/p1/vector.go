package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_InitVector(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14216(m, l0, int32(4))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_VectorSumCenter(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 float32
	_ = v34
	var v36 float32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 float32
	_ = v42
	var v44 float32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 float32
	_ = v50
	var v52 float32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 float32
	_ = v58
	var v60 float32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 float32
	_ = v97
	var v99 float32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	v3 = int32(0)
	v11 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v11 <= v3 {
	} else {
		v15 = l0 + int32(8)
		v16 = int32(0)
		if base.Ui32(int32(4)) <= base.Ui32(v11) {
			v21 = v16
			v29 = v3
			for {
				v32 = v21 << (uint(int32(2)) % 32)
				v33 = l1 + v32
				v34 = *(*float32)(unsafe.Add(mBase, uint32(v33)))
				v36 = *(*float32)(unsafe.Add(mBase, uint32(v32+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v33))) = base.F32_add(v34, v36)
				v39 = int32(4)
				v40 = v32 | v39
				v41 = l1 + v40
				v42 = *(*float32)(unsafe.Add(mBase, uint32(v41)))
				v44 = *(*float32)(unsafe.Add(mBase, uint32(v40+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v41))) = base.F32_add(v42, v44)
				v48 = v32 | int32(8)
				v49 = l1 + v48
				v50 = *(*float32)(unsafe.Add(mBase, uint32(v49)))
				v52 = *(*float32)(unsafe.Add(mBase, uint32(v48+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v49))) = base.F32_add(v50, v52)
				v56 = v32 | int32(12)
				v57 = l1 + v56
				v58 = *(*float32)(unsafe.Add(mBase, uint32(v57)))
				v60 = *(*float32)(unsafe.Add(mBase, uint32(v56+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v57))) = base.F32_add(v58, v60)
				v64 = v21 + v39
				v66 = v29 + v39
				if v66 != v11&int32(_a_F_VectorSumCenter_0) {
					v21 = v64
					v29 = v66
					continue
				} else {
					break
				}
				break
			}
			if v11&int32(3) == int32(0) {
			} else {
				v72 = v64
				v84 = v72
				v93 = v3
				for {
					v95 = v84 << (uint(int32(2)) % 32)
					v96 = l1 + v95
					v97 = *(*float32)(unsafe.Add(mBase, uint32(v96)))
					v99 = *(*float32)(unsafe.Add(mBase, uint32(v95+v15)))
					*(*float32)(unsafe.Add(mBase, uint32(v96))) = base.F32_add(v97, v99)
					v102 = int32(1)
					v105 = v93 + v102
					if v105 != v11&int32(3) {
						v84 = v84 + v102
						v93 = v105
						continue
					} else {
						break
					}
					break
				}
			}
		} else {
			v72 = v16
			v84 = v72
			v93 = v3
			for {
				v95 = v84 << (uint(int32(2)) % 32)
				v96 = l1 + v95
				v97 = *(*float32)(unsafe.Add(mBase, uint32(v96)))
				v99 = *(*float32)(unsafe.Add(mBase, uint32(v95+v15)))
				*(*float32)(unsafe.Add(mBase, uint32(v96))) = base.F32_add(v97, v99)
				v102 = int32(1)
				v105 = v93 + v102
				if v105 != v11&int32(3) {
					v84 = v84 + v102
					v93 = v105
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return
}
func F_vector_gt(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14402(m, l0, int64(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_vector_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v94 float32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	v10 = m.G0
	v12 = v10 - int32(_a_F_vector_in_0)
	m.G0 = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = v15
	goto L1
L1:
	;
	v27 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18))))
	goto L3
L2:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v37 == int32(91) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	if base.B2i32(v27 == int32(32))|base.B2i32(base.Ui32((v27-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v18 = v18 + int32(1)
		goto L1
	} else {
		goto L4
	}
L4:
	;
	goto L2
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L25
	} else {
		goto L87
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L25
	} else {
		goto L82
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L25
	} else {
		goto L77
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L25
	} else {
		goto L73
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L25
	} else {
		goto L69
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L25
	} else {
		goto L65
	}
L11:
	;
	v40 = base.I32_wrap_i64(v14)
	v41 = v18
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L25
	} else {
		goto L60
	}
L14:
	;
	v51 = v41 + int32(1)
	v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41)+1)))
	goto L16
L15:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v62 == int32(93) {
		goto L10
	} else {
		goto L18
	}
L16:
	;
	if base.B2i32(v52 == int32(32))|base.B2i32(base.Ui32((v52-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v41 = v51
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v65 = v51
	v71 = int32(0)
	goto L20
L19:
	;
	v183 = v126
	goto L47
L20:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65))))
	goto L22
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L25
	} else {
		goto L43
	}
L22:
	;
	if base.B2i32(v76 == int32(32))|base.B2i32(base.Ui32((v76-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v65 = v65 + int32(1)
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v86 == int32(0) {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_vector_in[0])) = int32(0)
	v94 = F_strtof(m, v65, v12+int32(124))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	return int64(0)
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v12)+124))
	if v98 == v65 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_vector_in[0]))
	if base.B2i32(v101 == int32(68))&base.F32_eq(base.F32_abs(v94), math.Float32frombits(uint32(0x7f800000))) != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	F_CheckElement_3(m, v94)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	*(*float32)(unsafe.Add(mBase, uint32(v12+int32(128)+v71<<(uint(int32(2))%32)))) = v94
	v116 = v98
	goto L30
L30:
	;
	v126 = v116 + int32(1)
	v127 = int32(*(*int8)(unsafe.Add(mBase, uint32(v116))))
	goto L32
L31:
	;
	v138 = v71 + int32(1)
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v139 != int32(44) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	if base.B2i32(v127 == int32(32))|base.B2i32(base.Ui32((v127-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v116 = v126
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v139 == int32(93) {
		goto L19
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v138 != int32(_a_F_vector_in_1) {
		v65 = v126
		v71 = v138
		goto L20
	} else {
		goto L42
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L25
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v15
	F_errmsg(m, int32(_a_F_vector_in_2), v12+int32(48))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L25
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(260), int32(_a_F_vector_in_4))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L25
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	goto L21
L43:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L25
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = int32(_a_F_vector_in_1)
	F_errmsg(m, int32(_a_F_vector_in_5), v12-int32(-64))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(214), int32(_a_F_vector_in_4))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L25
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v194 = int32(*(*int8)(unsafe.Add(mBase, uint32(v183))))
	goto L49
L48:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v204 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	if base.B2i32(v194 == int32(32))|base.B2i32(base.Ui32((v194-int32(9))&int32(255)) < base.Ui32(int32(5))) != 0 {
		v183 = v183 + int32(1)
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	F_CheckDim_3(m, v138)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	if base.B2i32(v40 != int32(-1))&base.B2i32(v138 != v40) != 0 {
		goto L5
	} else {
		goto L53
	}
L53:
	;
	v213 = F_mul_size(m, int32(4), v138)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	v215 = F_add_size(m, int32(8), v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	v217 = F_palloc0(m, v215)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L25
	} else {
		goto L56
	}
L56:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v217)+4)) = uint16(v138)
	v220 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v215 << (uint(v220) % 32)
	v226 = v71<<(uint(v220)%32) + int32(4)
	if v226 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	base.MemoryCopy(m, v217+int32(8), v12+int32(128), v226)
	goto L59
L58:
	;
	goto L59
L59:
	;
	m.G0 = v12 + int32(_a_F_vector_in_0)
	return base.I64_extend_i32_u(v217)
L60:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L25
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v15
	F_errmsg(m, int32(_a_F_vector_in_2), v12+int32(112))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L25
	} else {
		goto L62
	}
L62:
	;
	v252 = F_errdetail(m, int32(_a_F_vector_in_6), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(194), int32(_a_F_vector_in_4))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L25
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
	F_errcode(m, int32(130))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L25
	} else {
		goto L66
	}
L66:
	;
	F_errmsg(m, int32(_a_F_vector_in_7), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(204), int32(_a_F_vector_in_4))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v15
	F_errmsg(m, int32(_a_F_vector_in_2), v12)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L25
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(223), int32(_a_F_vector_in_4))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L25
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L25
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v15
	F_errmsg(m, int32(_a_F_vector_in_2), v12+int32(16))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L25
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(234), int32(_a_F_vector_in_4))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L25
	} else {
		goto L78
	}
L78:
	;
	v317 = F_pnstrdup(m, v65, v98-v65)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L25
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v317
	F_errmsg(m, int32(_a_F_vector_in_8), v12+int32(32))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L25
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(240), int32(_a_F_vector_in_4))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L25
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v15
	F_errmsg(m, int32(_a_F_vector_in_2), v12+int32(96))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L25
	} else {
		goto L84
	}
L84:
	;
	v345 = F_errdetail(m, int32(_a_F_vector_in_9), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L25
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(271), int32(_a_F_vector_in_4))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L25
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L25
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v40
	F_errmsg(m, int32(_a_F_vector_in_10), v12+int32(80))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L25
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_vector_in_3), int32(88), int32(_a_F_vector_in_11))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L25
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_vector_typmod_in(m *base.Module, l0 int32) int64 {
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	v11 = Fn14301(m, l0, int32(_a_F_vector_typmod_in_0), int32(366), int32(_a_F_vector_typmod_in_1), int32(_a_F_vector_typmod_in_2), int32(_a_F_vector_typmod_in_3), int32(361), int32(_a_F_vector_typmod_in_4), int32(356), int32(_a_F_vector_typmod_in_5))
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		return v11
	}
}
