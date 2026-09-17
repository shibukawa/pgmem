package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FreePageManagerInitialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	v3 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = v3
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v9
	v11 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)) = uint8(v11)
	if l0 != 0 {
		v17 = l0 - l1 + v11
	} else {
		v17 = v9
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v17
	base.MemoryFill(m, l0+int32(36), int32(0), int32(516))
	return
}
func F_PageAddItemExtended(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v259 int32
	_ = v259
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v17) < base.Ui32(int32(24)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L19
	} else {
		goto L59
	}
L2:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if base.Ui32(v20) < base.Ui32(v17) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if base.B2i32(base.Ui32(v22) < base.Ui32(v20))|base.B2i32(base.Ui32(int32(_a_F_PageAddItemExtended_0)) <= base.Ui32(v22)) != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v17 != int32(24) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = int32(base.Ui32(v17+int32(_a_F_PageAddItemExtended_1)) >> (uint(int32(2)) % 32))
	goto L7
L6:
	;
	v34 = int32(0)
	goto L7
L7:
	;
	v35 = int32(1)
	v36 = v34 + v35
	if base.Ui32((l3-v35)&int32(_a_F_PageAddItemExtended_2)) <= base.Ui32(int32(2047)) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	m.G0 = v15 + int32(16)
	return v259 & int32(_a_F_PageAddItemExtended_2)
L9:
	;
	v188 = int32(0)
	v191 = v177 & int32(_a_F_PageAddItemExtended_2)
	if base.B2i32(l4&int32(2) == v188)|base.B2i32(base.Ui32(v191) < base.Ui32(int32(292))) == v188 {
		goto L39
	} else {
		goto L40
	}
L10:
	;
	v153 = int32(_a_F_PageAddItemExtended_2)
	if base.Ui32(v144&v153) <= base.Ui32(v36&v153) {
		v177 = v144
		v183 = v150
		goto L9
	} else {
		goto L34
	}
L11:
	;
	if l4&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)))
	if v75&int32(1) == int32(0) {
		v177 = v36
		v183 = v6
		goto L9
	} else {
		goto L24
	}
L14:
	;
	if base.Ui32(v34&int32(_a_F_PageAddItemExtended_2)) < base.Ui32(l3) {
		v144 = l3
		v150 = v6
		goto L10
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v144 = l3
	v150 = base.B2i32(base.Ui32(l3) <= base.Ui32(v34&int32(_a_F_PageAddItemExtended_2)))
	goto L10
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0+l3<<(uint(int32(2))%32))+20))
	if base.Ui32(v51) < base.Ui32(int32(_a_F_PageAddItemExtended_3)) {
		v144 = l3
		v150 = v6
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v54 = int32(0)
	v57 = F_errstart(m, int32(19), v54)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v57 == int32(0) {
		v259 = v54
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_errmsg_internal(m, int32(_a_F_PageAddItemExtended_4), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_PageAddItemExtended_5), int32(235), int32(_a_F_PageAddItemExtended_6))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v259 = v54
	goto L8
L24:
	;
	v81 = v34 & int32(_a_F_PageAddItemExtended_2)
	if v81 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v144 = v131
	v150 = int32(0)
	goto L10
L26:
	;
	v126 = v75 & int32(_a_F_PageAddItemExtended_7)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+10)) = uint16(v126)
	v131 = v116
	goto L25
L27:
	;
	v116 = int32(1)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v91 = int32(1)
	goto L30
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(20)+v91&int32(_a_F_PageAddItemExtended_2)<<(uint(int32(2))%32))))
	if base.Ui32(v105) < base.Ui32(int32(_a_F_PageAddItemExtended_3)) {
		v131 = v91
		goto L25
	} else {
		goto L32
	}
L31:
	;
	v116 = v109
	goto L26
L32:
	;
	v109 = v91 + int32(1)
	if base.Ui32(v109&int32(_a_F_PageAddItemExtended_2)) <= base.Ui32(v81) {
		v91 = v109
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v158 = int32(0)
	v161 = F_errstart(m, int32(19), v158)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L19
	} else {
		goto L35
	}
L35:
	;
	if v161 == int32(0) {
		v259 = v158
		goto L8
	} else {
		goto L36
	}
L36:
	;
	F_errmsg_internal(m, int32(_a_F_PageAddItemExtended_8), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_PageAddItemExtended_5), int32(292), int32(_a_F_PageAddItemExtended_6))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	v259 = v158
	goto L8
L39:
	;
	v197 = int32(0)
	v200 = F_errstart(m, int32(19), v197)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L19
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v215 = v17 + int32(4)
	v217 = v36 & int32(_a_F_PageAddItemExtended_2)
	if v191 == v217 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	if v200 == int32(0) {
		v259 = v197
		goto L8
	} else {
		goto L43
	}
L43:
	;
	F_errmsg_internal(m, int32(_a_F_PageAddItemExtended_9), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L19
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(_a_F_PageAddItemExtended_5), int32(299), int32(_a_F_PageAddItemExtended_6))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	v259 = v197
	goto L8
L46:
	;
	v219 = v215
	goto L48
L47:
	;
	v219 = v17
	goto L48
L48:
	;
	if v183 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v220 = v215
	goto L51
L50:
	;
	v220 = v219
	goto L51
L51:
	;
	v225 = v20 - (l2+int32(7))&int32(-8)
	if v225 < v220 {
		v259 = int32(0)
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v229 = l0 + v191<<(uint(int32(2))%32)
	v231 = v229 + int32(20)
	if v183 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v225&int32(_a_F_PageAddItemExtended_10) | l2<<(uint(int32(17))%32) | int32(_a_F_PageAddItemExtended_3)
	if l2 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v236 = (v217 - v191) << (uint(int32(2)) % 32)
	if v236 == int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	base.MemoryCopy(m, v229+int32(24), v231, v236)
	goto L53
L56:
	;
	base.MemoryCopy(m, l0+v225, l1, l2)
	goto L58
L57:
	;
	goto L58
L58:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v225)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v220)
	v259 = v177
	goto L8
L59:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v284
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v282
	F_errmsg(m, int32(_a_F_PageAddItemExtended_11), v15)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L19
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_PageAddItemExtended_5), int32(217), int32(_a_F_PageAddItemExtended_6))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PageRestoreTempPage(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+19)))
	v6 = v4 << (uint(int32(8)) % 32)
	if v6 != 0 {
		base.MemoryCopy(m, l1, l0, v6)
	} else {
	}
	F_pfree(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_PageSetChecksumInplace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v3 == int32(0) {
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, _c_F_PageSetChecksumInplace[0]))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+252))
		if base.B2i32(v8 != int32(0)) == int32(0) {
		} else {
			v13 = int32(0)
			v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v13)
			v49 = m.G0
			v50 = int32(128)
			v51 = v49 - v50
			base.MemoryCopy(m, v51, int32(_a_F_PageSetChecksumInplace_0), v50)
			v58 = v13
			for {
				v92 = l0 + v58<<(uint(int32(7))%32)
				v99 = int32(0)
				for {
					v129 = int32(2)
					v130 = v99 << (uint(v129) % 32)
					v131 = v51 + v130
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v92+v130)))
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
					v135 = v133 ^ v134
					v136 = int32(16777619)
					v138 = int32(17)
					*(*int32)(unsafe.Add(mBase, uint32(v131))) = v135*v136 ^ int32(base.Ui32(v135)>>(uint(v138)%32))
					v143 = v130 | int32(4)
					v144 = v51 + v143
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v92+v143)))
					v147 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
					v148 = v146 ^ v147
					*(*int32)(unsafe.Add(mBase, uint32(v144))) = v148*v136 ^ int32(base.Ui32(v148)>>(uint(v138)%32))
					v156 = v99 + v129
					if v156 != int32(32) {
						v99 = v156
						continue
					} else {
						break
					}
					break
				}
				v160 = v58 + int32(1)
				if v160 != int32(64) {
					v58 = v160
					continue
				} else {
					break
				}
				break
			}
			v169 = int32(0)
			for {
				v201 = v51 + v169<<(uint(int32(2))%32)
				v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
				v203 = int32(16777619)
				v205 = int32(17)
				*(*int32)(unsafe.Add(mBase, uint32(v201))) = v202*v203 ^ int32(base.Ui32(v202)>>(uint(v205)%32))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v201)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v201)+4)) = v209*v203 ^ int32(base.Ui32(v209)>>(uint(v205)%32))
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v201)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v201)+8)) = v216*v203 ^ int32(base.Ui32(v216)>>(uint(v205)%32))
				v223 = *(*int32)(unsafe.Add(mBase, uint32(v201)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v201)+12)) = v223*v203 ^ int32(base.Ui32(v223)>>(uint(v205)%32))
				v231 = v169 + int32(4)
				if v231 != int32(32) {
					v169 = v231
					continue
				} else {
					break
				}
				break
			}
			v240 = int32(0)
			for {
				v272 = v51 + v240<<(uint(int32(2))%32)
				v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
				v274 = int32(16777619)
				v276 = int32(17)
				*(*int32)(unsafe.Add(mBase, uint32(v272))) = v273*v274 ^ int32(base.Ui32(v273)>>(uint(v276)%32))
				v280 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v272)+4)) = v280*v274 ^ int32(base.Ui32(v280)>>(uint(v276)%32))
				v287 = *(*int32)(unsafe.Add(mBase, uint32(v272)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v272)+8)) = v287*v274 ^ int32(base.Ui32(v287)>>(uint(v276)%32))
				v294 = *(*int32)(unsafe.Add(mBase, uint32(v272)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v272)+12)) = v294*v274 ^ int32(base.Ui32(v294)>>(uint(v276)%32))
				v302 = v240 + int32(4)
				if v302 != int32(32) {
					v240 = v302
					continue
				} else {
					break
				}
				break
			}
			v305 = *(*int32)(unsafe.Add(mBase, uint32(v51)+124))
			v306 = *(*int32)(unsafe.Add(mBase, uint32(v51)+120))
			v307 = *(*int32)(unsafe.Add(mBase, uint32(v51)+116))
			v308 = *(*int32)(unsafe.Add(mBase, uint32(v51)+112))
			v309 = *(*int32)(unsafe.Add(mBase, uint32(v51)+108))
			v310 = *(*int32)(unsafe.Add(mBase, uint32(v51)+104))
			v311 = *(*int32)(unsafe.Add(mBase, uint32(v51)+100))
			v312 = *(*int32)(unsafe.Add(mBase, uint32(v51)+96))
			v313 = *(*int32)(unsafe.Add(mBase, uint32(v51)+92))
			v314 = *(*int32)(unsafe.Add(mBase, uint32(v51)+88))
			v315 = *(*int32)(unsafe.Add(mBase, uint32(v51)+84))
			v316 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
			v317 = *(*int32)(unsafe.Add(mBase, uint32(v51)+76))
			v318 = *(*int32)(unsafe.Add(mBase, uint32(v51)+72))
			v319 = *(*int32)(unsafe.Add(mBase, uint32(v51)+68))
			v320 = *(*int32)(unsafe.Add(mBase, uint32(v51)+64))
			v321 = *(*int32)(unsafe.Add(mBase, uint32(v51)+60))
			v322 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
			v323 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
			v324 = *(*int32)(unsafe.Add(mBase, uint32(v51)+48))
			v325 = *(*int32)(unsafe.Add(mBase, uint32(v51)+44))
			v326 = *(*int32)(unsafe.Add(mBase, uint32(v51)+40))
			v327 = *(*int32)(unsafe.Add(mBase, uint32(v51)+36))
			v328 = *(*int32)(unsafe.Add(mBase, uint32(v51)+32))
			v329 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
			v330 = *(*int32)(unsafe.Add(mBase, uint32(v51)+24))
			v331 = *(*int32)(unsafe.Add(mBase, uint32(v51)+20))
			v332 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
			v333 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
			v334 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
			v335 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
			v336 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v46)
			v370 = int32(_a_F_PageSetChecksumInplace_1)
			v371 = base.I32_rem_u_s(v305^(v306^(v307^(v308^(v309^(v310^(v311^(v312^(v313^(v314^(v315^(v316^(v317^(v318^(v319^(v320^(v321^(v322^(v323^(v324^(v325^(v326^(v327^(v328^(v329^(v330^(v331^(v332^(v333^(v334^(v335^(l1^v336))))))))))))))))))))))))))))))), v370)
			v375 = (v371 + int32(1)) & v370
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)) = uint16(v375)
		}
	}
	return
}
func F_ReadPageInternal(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1160))
	v14 = base.I32_wrap_i64(l1) & (v11 - int32(1))
	v16 = base.I64_div_u_s(l1, base.I64_extend_i32_s(v11))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1176))
	if v16 != v17 {
		v24 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v24
		if base.B2i32(v14 == v24)|base.B2i32(v16 == v17) == v24 {
			v34 = l1 - base.I64_extend_i32_u(v14)
			v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v39 = m.T0[v38].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, v34, int32(_a_F_ReadPageInternal_0), v36, v37)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				if v39 == int32(-2) {
					v108 = int32(-2)
					return v108
				} else {
					if v39 < int32(0) {
						v94 = int32(0)
						v97 = v94
						v98 = int32(-1)
						v99 = v94
						v102 = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
						*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
						*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
						v108 = v98
						return v108
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v48 = F_XLogReaderValidatePageHeader(m, l0, v34, v47)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							if v48 == int32(0) {
								v94 = int32(0)
								v97 = v94
								v98 = int32(-1)
								v99 = v94
								v102 = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
								*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
								v108 = v98
								return v108
							} else {
								v55 = int32(-2)
								v56 = int32(24)
								if base.Ui32(l2) <= base.Ui32(v56) {
									v59 = v56
								} else {
									v59 = l2
								}
								v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v63 = m.T0[v62].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v59, v60, v61)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									if v63 == int32(-2) {
										v108 = v55
										return v108
									} else {
										if v63 < int32(25) {
											v94 = int32(0)
											v97 = v94
											v98 = int32(-1)
											v99 = v94
											v102 = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
											*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
											*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
											v108 = v98
											return v108
										} else {
											v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
											if v72&int32(2) != 0 {
												v75 = int32(40)
											} else {
												v75 = int32(24)
											}
											if base.Ui32(v63) < base.Ui32(v75) {
												v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
												v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v79 = m.T0[v78].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v75, v77, v71)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int32(0)
												} else {
													if v79 == int32(-2) {
														v108 = v55
														return v108
													} else {
														if v79 < int32(0) {
															v94 = int32(0)
															v97 = v94
															v98 = int32(-1)
															v99 = v94
															v102 = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
															*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
															*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
															v108 = v98
															return v108
														} else {
															v85 = v79
															v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
															mBase = m.M
															v87 = m.ExcPending
															if v87 != 0 {
																return int32(0)
															} else {
																if v86 != 0 {
																	v97 = v85
																	v98 = v85
																	v99 = v14
																	v102 = v16
																} else {
																	v94 = int32(0)
																	v97 = v94
																	v98 = int32(-1)
																	v99 = v94
																	v102 = int64(0)
																}
																*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
																*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
																*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
																v108 = v98
																return v108
															}
														}
													}
												}
											} else {
												v85 = v63
												v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													if v86 != 0 {
														v97 = v85
														v98 = v85
														v99 = v14
														v102 = v16
													} else {
														v94 = int32(0)
														v97 = v94
														v98 = int32(-1)
														v99 = v94
														v102 = int64(0)
													}
													*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
													*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
													*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
													v108 = v98
													return v108
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v55 = int32(-2)
			v56 = int32(24)
			if base.Ui32(l2) <= base.Ui32(v56) {
				v59 = v56
			} else {
				v59 = l2
			}
			v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
			v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v63 = m.T0[v62].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v59, v60, v61)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				if v63 == int32(-2) {
					v108 = v55
					return v108
				} else {
					if v63 < int32(25) {
						v94 = int32(0)
						v97 = v94
						v98 = int32(-1)
						v99 = v94
						v102 = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
						*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
						*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
						v108 = v98
						return v108
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
						if v72&int32(2) != 0 {
							v75 = int32(40)
						} else {
							v75 = int32(24)
						}
						if base.Ui32(v63) < base.Ui32(v75) {
							v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v79 = m.T0[v78].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v75, v77, v71)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								if v79 == int32(-2) {
									v108 = v55
									return v108
								} else {
									if v79 < int32(0) {
										v94 = int32(0)
										v97 = v94
										v98 = int32(-1)
										v99 = v94
										v102 = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
										*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
										v108 = v98
										return v108
									} else {
										v85 = v79
										v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
										mBase = m.M
										v87 = m.ExcPending
										if v87 != 0 {
											return int32(0)
										} else {
											if v86 != 0 {
												v97 = v85
												v98 = v85
												v99 = v14
												v102 = v16
											} else {
												v94 = int32(0)
												v97 = v94
												v98 = int32(-1)
												v99 = v94
												v102 = int64(0)
											}
											*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
											*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
											*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
											v108 = v98
											return v108
										}
									}
								}
							}
						} else {
							v85 = v63
							v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return int32(0)
							} else {
								if v86 != 0 {
									v97 = v85
									v98 = v85
									v99 = v14
									v102 = v16
								} else {
									v94 = int32(0)
									v97 = v94
									v98 = int32(-1)
									v99 = v94
									v102 = int64(0)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
								*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
								v108 = v98
								return v108
							}
						}
					}
				}
			}
		}
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1192))
		if v14 != v19 {
			v24 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v24
			if base.B2i32(v14 == v24)|base.B2i32(v16 == v17) == v24 {
				v34 = l1 - base.I64_extend_i32_u(v14)
				v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v39 = m.T0[v38].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, v34, int32(_a_F_ReadPageInternal_0), v36, v37)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					if v39 == int32(-2) {
						v108 = int32(-2)
						return v108
					} else {
						if v39 < int32(0) {
							v94 = int32(0)
							v97 = v94
							v98 = int32(-1)
							v99 = v94
							v102 = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
							*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
							*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
							v108 = v98
							return v108
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v48 = F_XLogReaderValidatePageHeader(m, l0, v34, v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								if v48 == int32(0) {
									v94 = int32(0)
									v97 = v94
									v98 = int32(-1)
									v99 = v94
									v102 = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
									*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
									v108 = v98
									return v108
								} else {
									v55 = int32(-2)
									v56 = int32(24)
									if base.Ui32(l2) <= base.Ui32(v56) {
										v59 = v56
									} else {
										v59 = l2
									}
									v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v63 = m.T0[v62].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v59, v60, v61)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										if v63 == int32(-2) {
											v108 = v55
											return v108
										} else {
											if v63 < int32(25) {
												v94 = int32(0)
												v97 = v94
												v98 = int32(-1)
												v99 = v94
												v102 = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
												*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
												*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
												v108 = v98
												return v108
											} else {
												v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
												v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
												if v72&int32(2) != 0 {
													v75 = int32(40)
												} else {
													v75 = int32(24)
												}
												if base.Ui32(v63) < base.Ui32(v75) {
													v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
													v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v79 = m.T0[v78].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v75, v77, v71)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														if v79 == int32(-2) {
															v108 = v55
															return v108
														} else {
															if v79 < int32(0) {
																v94 = int32(0)
																v97 = v94
																v98 = int32(-1)
																v99 = v94
																v102 = int64(0)
																*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
																*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
																*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
																v108 = v98
																return v108
															} else {
																v85 = v79
																v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
																mBase = m.M
																v87 = m.ExcPending
																if v87 != 0 {
																	return int32(0)
																} else {
																	if v86 != 0 {
																		v97 = v85
																		v98 = v85
																		v99 = v14
																		v102 = v16
																	} else {
																		v94 = int32(0)
																		v97 = v94
																		v98 = int32(-1)
																		v99 = v94
																		v102 = int64(0)
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
																	*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
																	v108 = v98
																	return v108
																}
															}
														}
													}
												} else {
													v85 = v63
													v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int32(0)
													} else {
														if v86 != 0 {
															v97 = v85
															v98 = v85
															v99 = v14
															v102 = v16
														} else {
															v94 = int32(0)
															v97 = v94
															v98 = int32(-1)
															v99 = v94
															v102 = int64(0)
														}
														*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
														*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
														*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
														v108 = v98
														return v108
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v55 = int32(-2)
				v56 = int32(24)
				if base.Ui32(l2) <= base.Ui32(v56) {
					v59 = v56
				} else {
					v59 = l2
				}
				v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v63 = m.T0[v62].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v59, v60, v61)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					if v63 == int32(-2) {
						v108 = v55
						return v108
					} else {
						if v63 < int32(25) {
							v94 = int32(0)
							v97 = v94
							v98 = int32(-1)
							v99 = v94
							v102 = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
							*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
							*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
							v108 = v98
							return v108
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
							if v72&int32(2) != 0 {
								v75 = int32(40)
							} else {
								v75 = int32(24)
							}
							if base.Ui32(v63) < base.Ui32(v75) {
								v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v79 = m.T0[v78].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v75, v77, v71)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									if v79 == int32(-2) {
										v108 = v55
										return v108
									} else {
										if v79 < int32(0) {
											v94 = int32(0)
											v97 = v94
											v98 = int32(-1)
											v99 = v94
											v102 = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
											*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
											*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
											v108 = v98
											return v108
										} else {
											v85 = v79
											v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
											mBase = m.M
											v87 = m.ExcPending
											if v87 != 0 {
												return int32(0)
											} else {
												if v86 != 0 {
													v97 = v85
													v98 = v85
													v99 = v14
													v102 = v16
												} else {
													v94 = int32(0)
													v97 = v94
													v98 = int32(-1)
													v99 = v94
													v102 = int64(0)
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
												*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
												*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
												v108 = v98
												return v108
											}
										}
									}
								}
							} else {
								v85 = v63
								v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									if v86 != 0 {
										v97 = v85
										v98 = v85
										v99 = v14
										v102 = v16
									} else {
										v94 = int32(0)
										v97 = v94
										v98 = int32(-1)
										v99 = v94
										v102 = int64(0)
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
									*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
									*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
									v108 = v98
									return v108
								}
							}
						}
					}
				}
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
			if base.Ui32(l2) <= base.Ui32(v21) {
				v108 = v21
				return v108
			} else {
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v24
				if base.B2i32(v14 == v24)|base.B2i32(v16 == v17) == v24 {
					v34 = l1 - base.I64_extend_i32_u(v14)
					v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v39 = m.T0[v38].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, v34, int32(_a_F_ReadPageInternal_0), v36, v37)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						if v39 == int32(-2) {
							v108 = int32(-2)
							return v108
						} else {
							if v39 < int32(0) {
								v94 = int32(0)
								v97 = v94
								v98 = int32(-1)
								v99 = v94
								v102 = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
								*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
								v108 = v98
								return v108
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v48 = F_XLogReaderValidatePageHeader(m, l0, v34, v47)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 == int32(0) {
										v94 = int32(0)
										v97 = v94
										v98 = int32(-1)
										v99 = v94
										v102 = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
										*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
										v108 = v98
										return v108
									} else {
										v55 = int32(-2)
										v56 = int32(24)
										if base.Ui32(l2) <= base.Ui32(v56) {
											v59 = v56
										} else {
											v59 = l2
										}
										v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
										v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v63 = m.T0[v62].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v59, v60, v61)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											if v63 == int32(-2) {
												v108 = v55
												return v108
											} else {
												if v63 < int32(25) {
													v94 = int32(0)
													v97 = v94
													v98 = int32(-1)
													v99 = v94
													v102 = int64(0)
													*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
													*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
													*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
													v108 = v98
													return v108
												} else {
													v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
													v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
													if v72&int32(2) != 0 {
														v75 = int32(40)
													} else {
														v75 = int32(24)
													}
													if base.Ui32(v63) < base.Ui32(v75) {
														v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
														v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v79 = m.T0[v78].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v75, v77, v71)
														mBase = m.M
														v80 = m.ExcPending
														if v80 != 0 {
															return int32(0)
														} else {
															if v79 == int32(-2) {
																v108 = v55
																return v108
															} else {
																if v79 < int32(0) {
																	v94 = int32(0)
																	v97 = v94
																	v98 = int32(-1)
																	v99 = v94
																	v102 = int64(0)
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
																	*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
																	v108 = v98
																	return v108
																} else {
																	v85 = v79
																	v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
																	mBase = m.M
																	v87 = m.ExcPending
																	if v87 != 0 {
																		return int32(0)
																	} else {
																		if v86 != 0 {
																			v97 = v85
																			v98 = v85
																			v99 = v14
																			v102 = v16
																		} else {
																			v94 = int32(0)
																			v97 = v94
																			v98 = int32(-1)
																			v99 = v94
																			v102 = int64(0)
																		}
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
																		*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
																		*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
																		v108 = v98
																		return v108
																	}
																}
															}
														}
													} else {
														v85 = v63
														v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return int32(0)
														} else {
															if v86 != 0 {
																v97 = v85
																v98 = v85
																v99 = v14
																v102 = v16
															} else {
																v94 = int32(0)
																v97 = v94
																v98 = int32(-1)
																v99 = v94
																v102 = int64(0)
															}
															*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
															*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
															*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
															v108 = v98
															return v108
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				} else {
					v55 = int32(-2)
					v56 = int32(24)
					if base.Ui32(l2) <= base.Ui32(v56) {
						v59 = v56
					} else {
						v59 = l2
					}
					v60 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v63 = m.T0[v62].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v59, v60, v61)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						if v63 == int32(-2) {
							v108 = v55
							return v108
						} else {
							if v63 < int32(25) {
								v94 = int32(0)
								v97 = v94
								v98 = int32(-1)
								v99 = v94
								v102 = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
								*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
								*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
								v108 = v98
								return v108
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
								if v72&int32(2) != 0 {
									v75 = int32(40)
								} else {
									v75 = int32(24)
								}
								if base.Ui32(v63) < base.Ui32(v75) {
									v77 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1216))
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v79 = m.T0[v78].(func(*base.Module, int32, int64, int32, int64, int32) int32)(m, l0, l1, v75, v77, v71)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										if v79 == int32(-2) {
											v108 = v55
											return v108
										} else {
											if v79 < int32(0) {
												v94 = int32(0)
												v97 = v94
												v98 = int32(-1)
												v99 = v94
												v102 = int64(0)
												*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
												*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
												*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
												v108 = v98
												return v108
											} else {
												v85 = v79
												v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int32(0)
												} else {
													if v86 != 0 {
														v97 = v85
														v98 = v85
														v99 = v14
														v102 = v16
													} else {
														v94 = int32(0)
														v97 = v94
														v98 = int32(-1)
														v99 = v94
														v102 = int64(0)
													}
													*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
													*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
													*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
													v108 = v98
													return v108
												}
											}
										}
									}
								} else {
									v85 = v63
									v86 = F_XLogReaderValidatePageHeader(m, l0, l1, v71)
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										if v86 != 0 {
											v97 = v85
											v98 = v85
											v99 = v14
											v102 = v16
										} else {
											v94 = int32(0)
											v97 = v94
											v98 = int32(-1)
											v99 = v94
											v102 = int64(0)
										}
										*(*int32)(unsafe.Add(mBase, uint32(l0)+1192)) = v99
										*(*int64)(unsafe.Add(mBase, uint32(l0)+1176)) = v102
										*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v97
										v108 = v98
										return v108
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
