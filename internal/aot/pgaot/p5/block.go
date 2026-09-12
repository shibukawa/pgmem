package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BlockRefTableMarkBlockModified(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	v4 = l3
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = int32(4442992)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v21
	v24 = v16 + int32(40)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = l2
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v31
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v27
	v38 = F_blockreftable_insert(m, v30, v16+int32(8), v16+int32(31))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+31)))
	if v40 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v43 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v38)+32)) = v43
	goto L5
L4:
	;
	goto L5
L5:
	;
	v50 = int32(base.Ui32(v4) >> (uint(int32(16)) % 32))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	if base.Ui32(v51) <= base.Ui32(v50) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v53 = int32(16)
	if base.Ui32(v51) <= base.Ui32(v53) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v144 = v50 << (uint(int32(1)) % 32)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144+v145))))
	if v147 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L9:
	;
	v56 = v53
	goto L11
L10:
	;
	v56 = v51
	goto L11
L11:
	;
	v57 = v56
	goto L12
L12:
	;
	v71 = v57 << (uint(int32(1)) % 32)
	if base.Ui32(v57) <= base.Ui32(v50) {
		v57 = v71
		goto L12
	} else {
		goto L14
	}
L13:
	;
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v57
	goto L8
L16:
	;
	v75 = F_palloc0(m, v71)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v87 = F_repalloc(m, v86, v71)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v75
	v78 = F_palloc0(m, v71)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v78
	v83 = F_palloc0(m, v57<<(uint(int32(2))%32))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v83
	goto L15
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = v87
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v91 = int32(1)
	v95 = v57 - v51
	v97 = v95 << (uint(v91) % 32)
	v99 = F__emscripten_memset_bulkmem(m, v87+v90<<(uint(v91)%32), base.I32_extend8_s(int32(0)), v97)
	mBase = m.M
	goto L23
L23:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v101 = F_repalloc(m, v100, v71)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v101
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v110 = F__emscripten_memset_bulkmem(m, v101+v104<<(uint(int32(1))%32), base.I32_extend8_s(int32(0)), v97)
	mBase = m.M
	goto L25
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v114 = F_repalloc(m, v111, v57<<(uint(int32(2))%32))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v114
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	v118 = int32(2)
	v125 = F__emscripten_memset_bulkmem(m, v114+v117<<(uint(v118)%32), base.I32_extend8_s(int32(0)), v95<<(uint(v118)%32))
	mBase = m.M
	goto L27
L27:
	;
	goto L15
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[9])) = v19
	m.G0 = v16 + int32(48)
	return
L29:
	;
	v151 = F_palloc(m, int32(32))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v171 = v4 & int32(65535)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v172+v144))))
	if v174 != int32(4096) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v154 = v50 << (uint(int32(2)) % 32)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v154+v155))) = v151
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v160 = int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v158+v144))) = uint16(v160)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v162+v154)))
	*(*uint16)(unsafe.Add(mBase, uint32(v164))) = uint16(v4)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v168 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v166+v144))) = uint16(v168)
	goto L28
L33:
	;
	goto L28
L34:
	;
	if v174 == v147 {
		goto L53
	} else {
		goto L54
	}
L35:
	;
	v204 = int32(0)
	goto L40
L36:
	;
	if v174 == int32(0) {
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187+v50<<(uint(int32(2))%32))))
	v196 = v191 + int32(base.Ui32(v171)>>(uint(int32(3))%32))&int32(8190)
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196))))
	v202 = v197 | int32(1)<<(uint(v4&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v196))) = uint16(v202)
	goto L28
L39:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v50<<(uint(int32(2))%32))))
	goto L35
L40:
	;
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183+v204<<(uint(int32(1))%32)))))
	if v220 == v4&int32(65535) {
		goto L33
	} else {
		goto L42
	}
L41:
	;
	if v174 != int32(4095) {
		goto L34
	} else {
		goto L44
	}
L42:
	;
	v223 = v204 + int32(1)
	if v223 != v174 {
		v204 = v223
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	v228 = F_palloc0(m, int32(8192))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v231 = v50 << (uint(int32(1)) % 32)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231+v232))))
	if v234 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v238 = int32(0)
	goto L49
L47:
	;
	goto L48
L48:
	;
	v293 = v228 + int32(base.Ui32(v171)>>(uint(int32(3))%32))&int32(8190)
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v293))))
	v299 = v294 | int32(1)<<(uint(v4&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v293))) = uint16(v299)
	v302 = v50 << (uint(int32(2)) % 32)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v302+v303)))
	F_pfree(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v251+v50<<(uint(int32(2))%32))))
	v254 = int32(1)
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v253+v238<<(uint(v254)%32)))))
	v262 = v228 + int32(base.Ui32(v257)>>(uint(int32(3))%32))&int32(8190)
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262))))
	v268 = v263 | v254<<(uint(v257&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v262))) = uint16(v268)
	v271 = v238 + v254
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v272+v231))))
	if base.Ui32(v271) < base.Ui32(v274) {
		v238 = v271
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L48
L51:
	;
	goto L50
L52:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v308+v302))) = v228
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v313 = int32(4096)
	*(*uint16)(unsafe.Add(mBase, uint32(v311+v231))) = uint16(v313)
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v315+v231))) = uint16(v313)
	goto L28
L53:
	;
	v333 = int32(2)
	v334 = v50 << (uint(v333) % 32)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v334+v335)))
	v340 = F_repalloc(m, v337, v147<<(uint(v333)%32))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	v356 = v174
	goto L55
L55:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v357+v50<<(uint(int32(2))%32))))
	v362 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v361+v356<<(uint(v362)%32)))) = uint16(v4)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v369 = v366 + v50<<(uint(v362)%32)
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369))))
	v372 = v370 + v362
	*(*uint16)(unsafe.Add(mBase, uint32(v369))) = uint16(v372)
	goto L33
L56:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v38)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v342+v334))) = v340
	v345 = int32(1)
	v346 = v50 << (uint(v345) % 32)
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v350 = v147 << (uint(v345) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v346+v347))) = uint16(v350)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v352+v346))))
	v356 = v354
	goto L55
}
func F_block_range_read_stream_cb(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v5) < base.Ui32(v6) {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5 + int32(1)
		v11 = v5
	} else {
		v11 = int32(-1)
	}
	return v11
}
