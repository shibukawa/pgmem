package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_a_f8_f8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v82 int32
	_ = v82
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v89 int32
	_ = v89
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 float64
	_ = v168
	var v171 int32
	_ = v171
	var v174 float64
	_ = v174
	var v177 int32
	_ = v177
	var v180 float64
	_ = v180
	var v183 int32
	_ = v183
	var v186 float64
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v209 int32
	_ = v209
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v227 float64
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v282 float64
	_ = v282
	var v285 int32
	_ = v285
	var v288 float64
	_ = v288
	var v291 int32
	_ = v291
	var v294 float64
	_ = v294
	var v297 int32
	_ = v297
	var v300 float64
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v338 int32
	_ = v338
	var v341 float64
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = F_array_contains_nulls(m, v20)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L1
	} else {
		goto L68
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L63
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L59
	}
L7:
	;
	if v27 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = F_array_contains_nulls(m, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v29 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v34 = F_ArrayGetNItemsSafe(m, v31, v20+int32(16))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if int32(101) <= v34 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v41 = F_ArrayGetNItemsSafe(m, v38, v25+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v41 != v34 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v52 = v44
	goto L17
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L17
L17:
	;
	v53 = v52 + v20
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v54 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v62 = v54
	goto L20
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v62 = (v55<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L20
L20:
	;
	v63 = v62 + v25
	v64 = int32(0)
	if v34 <= v64 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v129 = F_palloc0(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L30
	}
L22:
	;
	v122 = v2
	v128 = v34<<(uint(int32(3))%32) + int32(8)
	goto L21
L23:
	;
	v67 = v64
	goto L24
L24:
	;
	v82 = v67 << (uint(int32(3)) % 32)
	v84 = *(*float64)(unsafe.Add(mBase, uint32(v53+v82)))
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v63+v82)))
	if base.F64_eq(v84, v86) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v122 = int32(1)
	v128 = v34<<(uint(int32(4))%32) | int32(8)
	goto L21
L26:
	;
	v89 = v67 + int32(1)
	if v34 != v89 {
		v67 = v89
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	goto L25
L29:
	;
	goto L22
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v128 << (uint(int32(2)) % 32)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v134&int32(-2147483648) | v34
	if int32(0) < v34 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	m.G0 = v17 + int32(16)
	return v129
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v34 | int32(-2147483648)
	goto L31
L33:
	;
	v142 = v34 & int32(3)
	v144 = v129 + int32(8)
	v145 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v34) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L35
L35:
	;
	if v122 != 0 {
		goto L31
	} else {
		goto L58
	}
L36:
	;
	if v122 == int32(0) {
		goto L32
	} else {
		goto L47
	}
L37:
	;
	v150 = v145
	v163 = v2
	goto L40
L38:
	;
	v195 = v145
	goto L39
L39:
	;
	v209 = v195
	v218 = v2
	goto L44
L40:
	;
	v165 = v150 << (uint(int32(3)) % 32)
	v168 = *(*float64)(unsafe.Add(mBase, uint32(v53+v165)))
	*(*float64)(unsafe.Add(mBase, uint32(v144+v165))) = v168
	v171 = v165 | int32(8)
	v174 = *(*float64)(unsafe.Add(mBase, uint32(v53+v171)))
	*(*float64)(unsafe.Add(mBase, uint32(v144+v171))) = v174
	v177 = v165 | int32(16)
	v180 = *(*float64)(unsafe.Add(mBase, uint32(v53+v177)))
	*(*float64)(unsafe.Add(mBase, uint32(v144+v177))) = v180
	v183 = v165 | int32(24)
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v53+v183)))
	*(*float64)(unsafe.Add(mBase, uint32(v144+v183))) = v186
	v188 = int32(4)
	v189 = v150 + v188
	v191 = v163 + v188
	if v191 != v34&int32(2147483644) {
		v150 = v189
		v163 = v191
		goto L40
	} else {
		goto L42
	}
L41:
	;
	if v142 == int32(0) {
		goto L36
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v195 = v189
	goto L39
L44:
	;
	v224 = v209 << (uint(int32(3)) % 32)
	v227 = *(*float64)(unsafe.Add(mBase, uint32(v53+v224)))
	*(*float64)(unsafe.Add(mBase, uint32(v144+v224))) = v227
	v229 = int32(1)
	v232 = v218 + v229
	if v232 != v142 {
		v209 = v209 + v229
		v218 = v232
		goto L44
	} else {
		goto L46
	}
L45:
	;
	goto L36
L46:
	;
	goto L45
L47:
	;
	v250 = int32(3)
	v251 = v34 & v250
	v256 = v129 + v34<<(uint(v250)%32) + int32(8)
	v257 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v34) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v264 = v257
	v273 = int32(0)
	goto L51
L49:
	;
	v309 = v257
	goto L50
L50:
	;
	v323 = v309
	v326 = v257
	goto L55
L51:
	;
	v279 = v264 << (uint(int32(3)) % 32)
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v279+v63)))
	*(*float64)(unsafe.Add(mBase, uint32(v256+v279))) = v282
	v285 = v279 | int32(8)
	v288 = *(*float64)(unsafe.Add(mBase, uint32(v63+v285)))
	*(*float64)(unsafe.Add(mBase, uint32(v256+v285))) = v288
	v291 = v279 | int32(16)
	v294 = *(*float64)(unsafe.Add(mBase, uint32(v63+v291)))
	*(*float64)(unsafe.Add(mBase, uint32(v256+v291))) = v294
	v297 = v279 | int32(24)
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v297+v63)))
	*(*float64)(unsafe.Add(mBase, uint32(v256+v297))) = v300
	v302 = int32(4)
	v303 = v264 + v302
	v305 = v273 + v302
	if v305 != v34&int32(2147483644) {
		v264 = v303
		v273 = v305
		goto L51
	} else {
		goto L53
	}
L52:
	;
	if v251 == int32(0) {
		goto L31
	} else {
		goto L54
	}
L53:
	;
	goto L52
L54:
	;
	v309 = v303
	goto L50
L55:
	;
	v338 = v323 << (uint(int32(3)) % 32)
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v338+v63)))
	*(*float64)(unsafe.Add(mBase, uint32(v256+v338))) = v341
	v343 = int32(1)
	v346 = v326 + v343
	if v346 != v251 {
		v323 = v323 + v343
		v326 = v346
		goto L55
	} else {
		goto L57
	}
L56:
	;
	goto L31
L57:
	;
	goto L56
L58:
	;
	goto L32
L59:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errmsg(m, int32(_a_F_cube_a_f8_f8_0), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_cube_a_f8_f8_1), int32(158), int32(_a_F_cube_a_f8_f8_2))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
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
	F_errcode(m, int32(261))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_cube_a_f8_f8_3), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(100)
	F_errdetail(m, int32(_a_F_cube_a_f8_f8_4), v17)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_cube_a_f8_f8_1), int32(166), int32(_a_F_cube_a_f8_f8_2))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_cube_a_f8_f8_5), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_cube_a_f8_f8_1), int32(171), int32(_a_F_cube_a_f8_f8_2))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_cube_c_f8_f8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v96 int32
	_ = v96
	var v99 float64
	_ = v99
	var v102 int32
	_ = v102
	var v105 float64
	_ = v105
	var v108 int32
	_ = v108
	var v111 float64
	_ = v111
	var v114 int32
	_ = v114
	var v117 float64
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v147 int32
	_ = v147
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v168 float64
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v271 int32
	_ = v271
	var v273 float64
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 float64
	_ = v279
	var v283 float64
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v302 int32
	_ = v302
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 float64
	_ = v315
	var v321 float64
	_ = v321
	var v334 int32
	_ = v334
	var v353 int32
	_ = v353
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	v2 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
		if base.Ui32(v29&int32(2147483644)) < base.Ui32(int32(100)) {
			v34 = int32(0)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v37 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v39 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
			if base.B2i32(v34 <= v29)|base.F64_ne(v37, v39) == v34 {
				v47 = v29<<(uint(int32(3))%32) + int32(16)
				v48 = F_palloc0(m, v47)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v48))) = v47 << (uint(int32(2)) % 32)
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					v55 = v53 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v55 | int32(-2147483648)
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					v61 = v59 & int32(2147483647)
					if v61 == int32(0) {
					} else {
						v65 = v59 & int32(3)
						v66 = int32(8)
						v67 = v48 + v66
						v69 = v25 + v66
						v70 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v61) {
							v78 = v70
							v80 = int32(0)
							for {
								v96 = v78 << (uint(int32(3)) % 32)
								v99 = *(*float64)(unsafe.Add(mBase, uint32(v96+v69)))
								*(*float64)(unsafe.Add(mBase, uint32(v67+v96))) = v99
								v102 = v96 | int32(8)
								v105 = *(*float64)(unsafe.Add(mBase, uint32(v69+v102)))
								*(*float64)(unsafe.Add(mBase, uint32(v67+v102))) = v105
								v108 = v96 | int32(16)
								v111 = *(*float64)(unsafe.Add(mBase, uint32(v69+v108)))
								*(*float64)(unsafe.Add(mBase, uint32(v67+v108))) = v111
								v114 = v96 | int32(24)
								v117 = *(*float64)(unsafe.Add(mBase, uint32(v114+v69)))
								*(*float64)(unsafe.Add(mBase, uint32(v67+v114))) = v117
								v119 = int32(4)
								v120 = v78 + v119
								v122 = v80 + v119
								if v122 != v59&int32(2147483644) {
									v78 = v120
									v80 = v122
									continue
								} else {
									break
								}
								break
							}
							if v65 == int32(0) {
							} else {
								v128 = v120
								v147 = v128
								v158 = v2
								for {
									v165 = v147 << (uint(int32(3)) % 32)
									v168 = *(*float64)(unsafe.Add(mBase, uint32(v165+v69)))
									*(*float64)(unsafe.Add(mBase, uint32(v67+v165))) = v168
									v170 = int32(1)
									v173 = v158 + v170
									if v173 != v65 {
										v147 = v147 + v170
										v158 = v173
										continue
									} else {
										break
									}
									break
								}
							}
						} else {
							v128 = v70
							v147 = v128
							v158 = v2
							for {
								v165 = v147 << (uint(int32(3)) % 32)
								v168 = *(*float64)(unsafe.Add(mBase, uint32(v165+v69)))
								*(*float64)(unsafe.Add(mBase, uint32(v67+v165))) = v168
								v170 = int32(1)
								v173 = v158 + v170
								if v173 != v65 {
									v147 = v147 + v170
									v158 = v173
									continue
								} else {
									break
								}
								break
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v48+v55<<(uint(int32(3))%32)))) = v37
					v353 = v48
					v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v369 != v25 {
						F_pfree(m, v25)
						mBase = m.M
						v372 = m.ExcPending
						if v372 != 0 {
							return int32(0)
						} else {
							m.G0 = v22 + int32(16)
							return v353
						}
					} else {
						m.G0 = v22 + int32(16)
						return v353
					}
				}
			} else {
				v201 = v29<<(uint(int32(4))%32) + int32(24)
				v202 = F_palloc0(m, v201)
				mBase = m.M
				v203 = m.ExcPending
				if v203 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v202))) = v201 << (uint(int32(2)) % 32)
					v207 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					v208 = int32(2147483647)
					v211 = v207&v208 + int32(1)
					v212 = *(*int32)(unsafe.Add(mBase, uint32(v202)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v211 | v212&int32(-2147483648)
					v217 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
					v219 = v217 & v208
					if v219 == int32(0) {
						v334 = v211 & int32(2147483647)
					} else {
						v225 = v202 + int32(8)
						v228 = v225 + v211<<(uint(int32(3))%32)
						v229 = int32(0)
						if v229 < v217 {
							v232 = v217
						} else {
							v232 = v229
						}
						v234 = v211 & int32(2147483647)
						v236 = v25 + int32(8)
						if v219 != int32(1) {
							v248 = int32(0)
							v254 = v2
							for {
								v263 = int32(3)
								v264 = v254 << (uint(v263) % 32)
								v266 = v264 + v236
								v267 = *(*float64)(unsafe.Add(mBase, uint32(v266)))
								*(*float64)(unsafe.Add(mBase, uint32(v225+v264))) = v267
								v271 = v232 << (uint(v263) % 32)
								v273 = *(*float64)(unsafe.Add(mBase, uint32(v266+v271)))
								*(*float64)(unsafe.Add(mBase, uint32(v264+v228))) = v273
								v276 = v264 | int32(8)
								v278 = v276 + v236
								v279 = *(*float64)(unsafe.Add(mBase, uint32(v278)))
								*(*float64)(unsafe.Add(mBase, uint32(v225+v276))) = v279
								v283 = *(*float64)(unsafe.Add(mBase, uint32(v278+v271)))
								*(*float64)(unsafe.Add(mBase, uint32(v276+v228))) = v283
								v285 = int32(2)
								v286 = v254 + v285
								v288 = v248 + v285
								if v288 != v217&int32(2147483646) {
									v248 = v288
									v254 = v286
									continue
								} else {
									break
								}
								break
							}
							if v217&int32(1) == int32(0) {
								v334 = v234
							} else {
								v302 = v286
								v311 = int32(3)
								v312 = v302 << (uint(v311) % 32)
								v314 = v312 + v236
								v315 = *(*float64)(unsafe.Add(mBase, uint32(v314)))
								*(*float64)(unsafe.Add(mBase, uint32(v225+v312))) = v315
								v321 = *(*float64)(unsafe.Add(mBase, uint32(v314+v232<<(uint(v311)%32))))
								*(*float64)(unsafe.Add(mBase, uint32(v312+v228))) = v321
								v334 = v234
							}
						} else {
							v302 = v2
							v311 = int32(3)
							v312 = v302 << (uint(v311) % 32)
							v314 = v312 + v236
							v315 = *(*float64)(unsafe.Add(mBase, uint32(v314)))
							*(*float64)(unsafe.Add(mBase, uint32(v225+v312))) = v315
							v321 = *(*float64)(unsafe.Add(mBase, uint32(v314+v232<<(uint(v311)%32))))
							*(*float64)(unsafe.Add(mBase, uint32(v312+v228))) = v321
							v334 = v234
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v202+v334<<(uint(int32(3))%32)))) = v37
					*(*float64)(unsafe.Add(mBase, uint32(v202+v211<<(uint(int32(4))%32)))) = v39
					v353 = v202
					v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v369 != v25 {
						F_pfree(m, v25)
						mBase = m.M
						v372 = m.ExcPending
						if v372 != 0 {
							return int32(0)
						} else {
							m.G0 = v22 + int32(16)
							return v353
						}
					} else {
						m.G0 = v22 + int32(16)
						return v353
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v380 = m.ExcPending
			if v380 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v383 = m.ExcPending
				if v383 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_cube_c_f8_f8_0), int32(0))
					mBase = m.M
					v387 = m.ExcPending
					if v387 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(100)
						F_errdetail(m, int32(_a_F_cube_c_f8_f8_1), v22)
						mBase = m.M
						v392 = m.ExcPending
						if v392 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cube_c_f8_f8_2), int32(1883), int32(_a_F_cube_c_f8_f8_3))
							mBase = m.M
							v397 = m.ExcPending
							if v397 != 0 {
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
		}
	}
}
func F_cube_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v58 float64
	_ = v58
	var v66 float64
	_ = v66
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v120 float64
	_ = v120
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v131 float64
	_ = v131
	var v134 float64
	_ = v134
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v165 int32
	_ = v165
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(0)
	if base.B2i32(v11 == v13)|base.B2i32(v6 == v13) != 0 {
		v165 = v13
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v184 != v6 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v183 = v165
	goto L4
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v31 = int32(2147483647)
	v32 = v30 & v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v35 = v33 & v31
	if base.Ui32(v32) < base.Ui32(v35) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v43 = v32
	goto L10
L8:
	;
	goto L9
L9:
	;
	if base.Ui32(v32) < base.Ui32(v35) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	v57 = v6 + int32(8) + v43<<(uint(int32(3))%32)
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v57)))
	if base.F64_ne(v58, float64(0)) != 0 {
		v165 = v13
		goto L5
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	if base.B2i32(v33 < int32(0)) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v66 = *(*float64)(unsafe.Add(mBase, uint32(v57+v35<<(uint(int32(3))%32))))
	if base.F64_ne(v66, float64(0)) != 0 {
		v165 = v13
		goto L5
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v70 = v43 + int32(1)
	if v70 != v35 {
		v43 = v70
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	goto L11
L18:
	;
	v87 = v32
	goto L20
L19:
	;
	v87 = v35
	goto L20
L20:
	;
	if v87 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v183 = int32(1)
	goto L4
L22:
	;
	goto L23
L23:
	;
	v91 = int32(8)
	v98 = int32(0)
	goto L24
L24:
	;
	v110 = int32(0)
	v112 = v98 << (uint(int32(3)) % 32)
	v113 = v11 + v91 + v112
	v114 = *(*float64)(unsafe.Add(mBase, uint32(v113)))
	v116 = base.B2i32(v30 < v110)
	if v30 < v110 {
		v123 = v114
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v165 = v151
	goto L5
L26:
	;
	v124 = v112 + (v6 + v91)
	v125 = *(*float64)(unsafe.Add(mBase, uint32(v124)))
	v127 = base.B2i32(v33 < int32(0))
	if v33 < int32(0) {
		v134 = v125
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v120 = *(*float64)(unsafe.Add(mBase, uint32(v113+v30<<(uint(int32(3))%32))))
	if base.F64_lt(v114, v120) != 0 {
		v123 = v114
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v123 = v120
	goto L26
L29:
	;
	if base.F64_gt(v123, v134) != 0 {
		v165 = v110
		goto L5
	} else {
		goto L32
	}
L30:
	;
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v124+v33<<(uint(int32(3))%32))))
	if base.F64_lt(v125, v131) != 0 {
		v134 = v125
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v134 = v131
	goto L29
L32:
	;
	if v30 < v110 {
		v141 = v114
		goto L33
	} else {
		goto L34
	}
L33:
	;
	if v33 < int32(0) {
		v148 = v125
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v139 = *(*float64)(unsafe.Add(mBase, uint32(v113+v30<<(uint(int32(3))%32))))
	if base.F64_gt(v114, v139) != 0 {
		v141 = v114
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v141 = v139
	goto L33
L36:
	;
	if base.F64_gt(v148, v141) != 0 {
		v165 = v110
		goto L5
	} else {
		goto L39
	}
L37:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v124+v33<<(uint(int32(3))%32))))
	if base.F64_gt(v125, v146) != 0 {
		v148 = v125
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v148 = v146
	goto L36
L39:
	;
	v151 = int32(1)
	v153 = v98 + v151
	if v153 != v87 {
		v98 = v153
		goto L24
	} else {
		goto L40
	}
L40:
	;
	goto L25
L41:
	;
	F_pfree(m, v6)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v188 != v11 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	F_pfree(m, v11)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	return v183
L48:
	;
	goto L47
}
func F_cube_coord_llur(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v50 float64
	_ = v50
	var v56 float64
	_ = v56
	var v60 float64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v13 != 0 {
			v16 = v13 >> (uint(int32(31)) % 32)
			v18 = v13 ^ v16 - v16
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			if base.Ui32(v19<<(uint(int32(1))%32)) < base.Ui32(v18) {
				v56 = float64(0)
			} else {
				v23 = int32(1)
				v24 = v18 - v23
				v26 = int32(base.Ui32(v24) >> (uint(v23) % 32))
				if v19 < int32(0) {
					v32 = *(*float64)(unsafe.Add(mBase, uint32(v9+v26<<(uint(int32(3))%32))+8))
					v56 = v32
				} else {
					v33 = int32(3)
					v35 = v9 + v26<<(uint(v33)%32)
					v43 = *(*float64)(unsafe.Add(mBase, uint32(v35+int32(8)+v19&int32(2147483647)<<(uint(v33)%32))))
					v44 = *(*float64)(unsafe.Add(mBase, uint32(v35)+8))
					if base.F64_lt(v43, v44) != 0 {
						v46 = v44
					} else {
						v46 = v43
					}
					if v24&int32(1) != 0 {
						v56 = v46
					} else {
						if base.F64_gt(v43, v44) != 0 {
							v50 = v44
						} else {
							v50 = v43
						}
						v56 = v50
					}
				}
			}
			if v13 < int32(0) {
				v60 = base.F64_neg(v56)
			} else {
				v60 = v56
			}
			v61 = F_Float8GetDatum(m, v60)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				return v61
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_cube_coord_llur_0), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_cube_coord_llur_1), int32(1659), int32(_a_F_cube_coord_llur_2))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
	}
}
func F_cube_f8_f8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.F64_eq(v5, v7) != 0 {
		v10 = F_palloc0(m, int32(16))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v5
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = int64(-9223372032559808448)
			return v10
		}
	} else {
		v19 = F_palloc0(m, int32(24))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(96)
			*(*float64)(unsafe.Add(mBase, uint32(v19)+16)) = v7
			*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v5
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v25&int32(-2147483648) | int32(1)
			return v19
		}
	}
}
func F_cube_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_cube_scanner_init(m, v7, v5+int32(8), v5+int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v21 = F_cube_yyparse(m, v5+int32(12), v18, v19, v20)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
			F_cube_scanner_finish(m, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				m.G0 = v5 + int32(16)
				return v26
			}
		}
	}
}
func F_cube_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 float64
	_ = v129
	var v131 int32
	_ = v131
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v142 int32
	_ = v142
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v156 float64
	_ = v156
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v167 int32
	_ = v167
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v212 int32
	_ = v212
	var v213 float64
	_ = v213
	var v220 float64
	_ = v220
	var v222 int32
	_ = v222
	var v226 float64
	_ = v226
	var v232 float64
	_ = v232
	var v238 float64
	_ = v238
	var v242 int32
	_ = v242
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v280 float64
	_ = v280
	var v286 float64
	_ = v286
	var v292 float64
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v335 float64
	_ = v335
	var v337 int32
	_ = v337
	var v341 float64
	_ = v341
	var v353 float64
	_ = v353
	var v357 float64
	_ = v357
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v386 float64
	_ = v386
	var v393 float64
	_ = v393
	var v395 int32
	_ = v395
	var v399 float64
	_ = v399
	var v411 float64
	_ = v411
	var v415 float64
	_ = v415
	var v420 int32
	_ = v420
	var v452 int32
	_ = v452
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v497 != v6 {
		goto L114
	} else {
		goto L115
	}
L5:
	;
	v496 = int32(-1)
	goto L4
L6:
	;
	v496 = v452
	goto L4
L7:
	;
	v452 = int32(1)
	goto L6
L8:
	;
	v36 = v31
	goto L10
L9:
	;
	v36 = v34
	goto L10
L10:
	;
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = int32(8)
	v54 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v60 = v54 << (uint(int32(3)) % 32)
	v61 = v6 + v37 + v60
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v71 = v62
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v103 = int32(8)
	v121 = int32(0)
	goto L31
L16:
	;
	v72 = v60 + (v11 + v37)
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
	v75 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v82 = v73
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v61+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v68) != 0 {
		v71 = v62
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v71 = v68
	goto L16
L19:
	;
	if base.F64_gt(v71, v82) != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v72+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v79) != 0 {
		v82 = v73
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v82 = v79
	goto L19
L22:
	;
	if v29 < int32(0) {
		v89 = v62
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v32 < int32(0) {
		v96 = v73
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v61+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v87) != 0 {
		v89 = v62
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v89 = v87
	goto L23
L26:
	;
	v98 = int32(-1)
	if base.F64_lt(v89, v96) != 0 {
		v452 = v98
		goto L6
	} else {
		goto L29
	}
L27:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v72+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v94) != 0 {
		v96 = v73
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v96 = v94
	goto L26
L29:
	;
	v101 = v54 + int32(1)
	if v101 != v36 {
		v54 = v101
		goto L14
	} else {
		goto L30
	}
L30:
	;
	goto L15
L31:
	;
	v127 = v121 << (uint(int32(3)) % 32)
	v128 = v6 + v103 + v127
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
	v131 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v138 = v129
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L13
L33:
	;
	v139 = v127 + (v11 + v103)
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v139)))
	v142 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v149 = v140
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v128+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v129, v135) != 0 {
		v138 = v129
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v138 = v135
	goto L33
L36:
	;
	if base.F64_gt(v138, v149) != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v139+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v140, v146) != 0 {
		v149 = v140
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v149 = v146
	goto L36
L39:
	;
	if v29 < int32(0) {
		v156 = v129
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v32 < int32(0) {
		v163 = v140
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v128+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v129, v154) != 0 {
		v156 = v129
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v156 = v154
	goto L40
L43:
	;
	if base.F64_lt(v156, v163) != 0 {
		v452 = v98
		goto L6
	} else {
		goto L46
	}
L44:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v139+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v140, v161) != 0 {
		v163 = v140
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v163 = v161
	goto L43
L46:
	;
	v167 = v121 + int32(1)
	if v167 != v36 {
		v121 = v167
		goto L31
	} else {
		goto L47
	}
L47:
	;
	goto L32
L48:
	;
	v189 = v6 + int32(8)
	v192 = v36
	goto L51
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L81
	} else {
		goto L82
	}
L51:
	;
	v212 = v189 + v192<<(uint(int32(3))%32)
	v213 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v255 = v36
	goto L65
L53:
	;
	if base.F64_lt(v238, float64(0)) != 0 {
		goto L5
	} else {
		goto L63
	}
L54:
	;
	v220 = *(*float64)(unsafe.Add(mBase, uint32(v212+v31<<(uint(int32(3))%32))))
	if base.F64_lt(v213, v220) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if base.F64_gt(v213, float64(0)) != 0 {
		goto L7
	} else {
		goto L62
	}
L57:
	;
	v222 = int32(0)
	goto L59
L58:
	;
	v222 = v29
	goto L59
L59:
	;
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v212+v222<<(uint(int32(3))%32))))
	if base.F64_gt(v226, float64(0)) != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v212+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v213, v232) == int32(0) {
		v238 = v232
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v238 = v213
	goto L53
L62:
	;
	v238 = v213
	goto L53
L63:
	;
	v242 = v192 + int32(1)
	if v242 != v31 {
		v192 = v242
		goto L51
	} else {
		goto L64
	}
L64:
	;
	goto L52
L65:
	;
	v266 = v189 + v255<<(uint(int32(3))%32)
	v267 = *(*float64)(unsafe.Add(mBase, uint32(v266)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L5
L67:
	;
	if base.F64_lt(v292, float64(0)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L68:
	;
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v266+v31<<(uint(int32(3))%32))))
	if base.F64_gt(v267, v274) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if base.F64_gt(v267, float64(0)) != 0 {
		goto L7
	} else {
		goto L76
	}
L71:
	;
	v276 = int32(0)
	goto L73
L72:
	;
	v276 = v29
	goto L73
L73:
	;
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v266+v276<<(uint(int32(3))%32))))
	if base.F64_gt(v280, float64(0)) != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v266+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v267, v286) == int32(0) {
		v292 = v286
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v292 = v267
	goto L67
L76:
	;
	v292 = v267
	goto L67
L77:
	;
	v297 = int32(1)
	v299 = v255 + v297
	if v299 == v31 {
		v452 = v297
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L66
L80:
	;
	v255 = v299
	goto L65
L81:
	;
	v496 = int32(0)
	goto L4
L82:
	;
	goto L83
L83:
	;
	v304 = v11 + int32(8)
	v314 = v31
	goto L84
L84:
	;
	v327 = v304 + v314<<(uint(int32(3))%32)
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v327)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v374 = v36
	goto L99
L86:
	;
	if base.F64_lt(v357, float64(0)) != 0 {
		goto L7
	} else {
		goto L97
	}
L87:
	;
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v327+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v328, v353) == int32(0) {
		v357 = v353
		goto L86
	} else {
		goto L96
	}
L88:
	;
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v327+v34<<(uint(int32(3))%32))))
	if base.F64_lt(v328, v335) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if base.F64_gt(v328, float64(0)) == int32(0) {
		v357 = v328
		goto L86
	} else {
		goto L95
	}
L91:
	;
	v337 = int32(0)
	goto L93
L92:
	;
	v337 = v32
	goto L93
L93:
	;
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v327+v337<<(uint(int32(3))%32))))
	if base.F64_gt(v341, float64(0)) == int32(0) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	goto L5
L95:
	;
	goto L5
L96:
	;
	v357 = v328
	goto L86
L97:
	;
	v361 = v314 + int32(1)
	if v361 != v34 {
		v314 = v361
		goto L84
	} else {
		goto L98
	}
L98:
	;
	goto L85
L99:
	;
	v385 = v304 + v374<<(uint(int32(3))%32)
	v386 = *(*float64)(unsafe.Add(mBase, uint32(v385)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v452 = int32(-1)
	goto L6
L101:
	;
	if base.F64_lt(v415, float64(0)) != 0 {
		goto L7
	} else {
		goto L112
	}
L102:
	;
	v411 = *(*float64)(unsafe.Add(mBase, uint32(v385+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v386, v411) == int32(0) {
		v415 = v411
		goto L101
	} else {
		goto L111
	}
L103:
	;
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v385+v34<<(uint(int32(3))%32))))
	if base.F64_gt(v386, v393) != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if base.F64_gt(v386, float64(0)) == int32(0) {
		v415 = v386
		goto L101
	} else {
		goto L110
	}
L106:
	;
	v395 = int32(0)
	goto L108
L107:
	;
	v395 = v32
	goto L108
L108:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v385+v395<<(uint(int32(3))%32))))
	if base.F64_gt(v399, float64(0)) == int32(0) {
		goto L102
	} else {
		goto L109
	}
L109:
	;
	goto L5
L110:
	;
	goto L5
L111:
	;
	v415 = v386
	goto L101
L112:
	;
	v420 = v374 + int32(1)
	if v34 != v420 {
		v374 = v420
		goto L99
	} else {
		goto L113
	}
L113:
	;
	goto L100
L114:
	;
	F_pfree(m, v6)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v501 != v11 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	F_pfree(m, v11)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	return base.B2i32(v496 <= int32(0))
L121:
	;
	goto L120
}
func F_cube_scanner_init(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	v4 = int32(0)
	v13 = F_strlen(m, l0)
	mBase = m.M
	if l2 != 0 {
		v16 = F_palloc(m, int32(96))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v16
			if v16 != 0 {
				v37 = int32(0)
				base.MemoryFill(m, v16, v37, int32(96))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				*(*int32)(unsafe.Add(mBase, uint32(v40)+60)) = v37
				v43 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v40)+52)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v40)+44)) = v37
				*(*int64)(unsafe.Add(mBase, uint32(v40)+36)) = v43
				*(*int64)(unsafe.Add(mBase, uint32(v40)+4)) = v43
				*(*int64)(unsafe.Add(mBase, uint32(v40)+12)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = v37
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v57 = v13 + int32(2)
				v58 = F_palloc(m, v57)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					if v58 != 0 {
						if v13 <= int32(0) {
						} else {
							v63 = v13 & int32(3)
							v64 = int32(0)
							if base.Ui32(int32(4)) <= base.Ui32(v13) {
								v71 = v64
								v79 = v4
								for {
									v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v71))))
									*(*uint8)(unsafe.Add(mBase, uint32(v71+v58))) = uint8(v83)
									v86 = v71 | int32(1)
									v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v86))))
									*(*uint8)(unsafe.Add(mBase, uint32(v58+v86))) = uint8(v89)
									v92 = v71 | int32(2)
									v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v92))))
									*(*uint8)(unsafe.Add(mBase, uint32(v58+v92))) = uint8(v95)
									v98 = v71 | int32(3)
									v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v98))))
									*(*uint8)(unsafe.Add(mBase, uint32(v58+v98))) = uint8(v101)
									v103 = int32(4)
									v104 = v71 + v103
									v106 = v79 + v103
									if v106 != v13&int32(2147483644) {
										v71 = v104
										v79 = v106
										continue
									} else {
										break
									}
									break
								}
								if v63 == int32(0) {
								} else {
									v112 = v104
									v124 = v112
									v133 = v4
									for {
										v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v124))))
										*(*uint8)(unsafe.Add(mBase, uint32(v124+v58))) = uint8(v136)
										v138 = int32(1)
										v141 = v133 + v138
										if v141 != v63 {
											v124 = v124 + v138
											v133 = v141
											continue
										} else {
											break
										}
										break
									}
								}
							} else {
								v112 = v64
								v124 = v112
								v133 = v4
								for {
									v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v124))))
									*(*uint8)(unsafe.Add(mBase, uint32(v124+v58))) = uint8(v136)
									v138 = int32(1)
									v141 = v133 + v138
									if v141 != v63 {
										v124 = v124 + v138
										v133 = v141
										continue
									} else {
										break
									}
									break
								}
							}
						}
						v156 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v58+v13))) = uint16(v156)
						v158 = F_cube_yy_scan_buffer(m, v58, v57, v55)
						mBase = m.M
						v159 = m.ExcPending
						if v159 != 0 {
							return
						} else {
							if v158 == int32(0) {
								F_yy_fatal_error_6(m, int32(_a_F_cube_scanner_init_0))
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v158)+20)) = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v13
								return
							}
						}
					} else {
						F_yy_fatal_error_6(m, int32(_a_F_cube_scanner_init_1))
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v22 = int32(48)
				*(*int32)(unsafe.Add(mBase, _c_F_cube_scanner_init[0])) = v22
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F_cube_scanner_init_2), int32(0))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_cube_scanner_init_3), int32(108), int32(_a_F_cube_scanner_init_4))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		v22 = int32(28)
		*(*int32)(unsafe.Add(mBase, _c_F_cube_scanner_init[0])) = v22
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_cube_scanner_init_2), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_cube_scanner_init_3), int32(108), int32(_a_F_cube_scanner_init_4))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_cube_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v33 float64
	_ = v33
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v61 float64
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 float64
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v92 float64
	_ = v92
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	v2 = float64(0)
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			v92 = v2
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v19 <= int32(0) {
				v92 = v2
			} else {
				v23 = v13 + int32(8)
				if v19 == int32(1) {
					v70 = float64(1)
					v72 = v3
					v80 = int32(3)
					v82 = v23 + v72<<(uint(v80)%32)
					v86 = *(*float64)(unsafe.Add(mBase, uint32(v82+v19<<(uint(v80)%32))))
					v87 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
					v92 = base.F64_mul(v70, base.F64_abs(base.F64_sub(v86, v87)))
				} else {
					v33 = float64(1)
					v35 = v3
					v42 = v3
					for {
						v43 = int32(3)
						v45 = v23 + v35<<(uint(v43)%32)
						v47 = v19 << (uint(v43) % 32)
						v49 = *(*float64)(unsafe.Add(mBase, uint32(v45+v47)))
						v50 = *(*float64)(unsafe.Add(mBase, uint32(v45)))
						v57 = *(*float64)(unsafe.Add(mBase, uint32(v45+int32(8)+v47)))
						v58 = *(*float64)(unsafe.Add(mBase, uint32(v45)+8))
						v61 = base.F64_mul(base.F64_mul(v33, base.F64_abs(base.F64_sub(v49, v50))), base.F64_abs(base.F64_sub(v57, v58)))
						v62 = int32(2)
						v63 = v35 + v62
						v65 = v42 + v62
						if v65 != v19&int32(2147483646) {
							v33 = v61
							v35 = v63
							v42 = v65
							continue
						} else {
							break
						}
						break
					}
					if v19&int32(1) == int32(0) {
						v92 = v61
					} else {
						v70 = v61
						v72 = v63
						v80 = int32(3)
						v82 = v23 + v72<<(uint(v80)%32)
						v86 = *(*float64)(unsafe.Add(mBase, uint32(v82+v19<<(uint(v80)%32))))
						v87 = *(*float64)(unsafe.Add(mBase, uint32(v82)))
						v92 = base.F64_mul(v70, base.F64_abs(base.F64_sub(v86, v87)))
					}
				}
			}
		}
		v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v102 != v13 {
			F_pfree(m, v13)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int32(0)
			} else {
				v106 = F_Float8GetDatum(m, v92)
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int32(0)
				} else {
					return v106
				}
			}
		} else {
			v106 = F_Float8GetDatum(m, v92)
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return int32(0)
			} else {
				return v106
			}
		}
	}
}
func F_cube_yyalloc(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc(m, l0)
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
