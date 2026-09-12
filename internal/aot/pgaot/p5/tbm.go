package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tbm_add_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	if v5 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = int32(-1)
	v32 = v5
	v39 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v20 + int32(16)
	return
L4:
	;
	v44 = l1 + v39*int32(6)
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)))
	if base.Ui32(int32(65244)) < base.Ui32((v45-int32(292))&int32(65535)) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L3
L6:
	;
	v313 = v39 + int32(1)
	if v313 != l2 {
		v30 = v300
		v32 = v302
		v39 = v313
		goto L4
	} else {
		goto L56
	}
L7:
	;
	if v164 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L8:
	;
	v155 = F_tbm_get_pageentry(m, l0, v56)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L23
	} else {
		goto L27
	}
L9:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)))
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v44))))
	v56 = v52 | v53<<(uint(int32(16))%32)
	if v56 == v30 {
		v162 = v30
		v164 = v32
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v58 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+20))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v65 = v56 & int32(-256)
	v68 = (v65 ^ v53) * int32(-2048144789)
	v73 = (int32(base.Ui32(v68)>>(uint(int32(13))%32)) ^ v68) * int32(-1028477387)
	v77 = v63 & (int32(base.Ui32(v73)>>(uint(int32(16))%32)) ^ v73)
	v80 = v62 + v77*int32(48)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
	if v81 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v88 = v80
	v89 = v77
	goto L15
L15:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	if v65 != v101 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+5)))
	if v110 != int32(1) {
		goto L8
	} else {
		goto L21
	}
L17:
	;
	v105 = (v89 + int32(1)) & v63
	v108 = v62 + v105*int32(48)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+4)))
	if v109 != 0 {
		v88 = v108
		v89 = v105
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L8
L21:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v88+int32(base.Ui32(v52)>>(uint(int32(3))%32))&int32(28))+8))
	if int32(base.Ui32(v118)>>(uint(v52)%32))&int32(1) == int32(0) {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v300 = v56
	v302 = int32(0)
	goto L6
L23:
	;
	return
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v45
	F_errmsg_internal(m, int32(59076), v20)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(497138), int32(383), int32(164808))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v162 = v56
	v164 = v155
	goto L7
L28:
	;
	v300 = v162
	v302 = int32(0)
	goto L6
L29:
	;
	goto L30
L30:
	;
	v179 = v45 - int32(1)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+5)))
	if v182 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v183 = int32(0)
	goto L33
L32:
	;
	v183 = int32(base.Ui32(v179) >> (uint(int32(5)) % 32))
	goto L33
L33:
	;
	v188 = v164 + v183<<(uint(int32(2))%32) + int32(8)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = int32(1)
	if v182 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v193 = v190
	goto L36
L35:
	;
	v193 = v190 << (uint(v179) % 32)
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v189 | v193
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+6)))
	v197 = v196 | l3
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+6)) = uint8(v197)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v199 <= v200 {
		v300 = v162
		v302 = v164
		goto L6
	} else {
		goto L37
	}
L37:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v206 = v204 & v205
	v211 = v206
	v212 = int32(0)
	v216 = v199
	v217 = v203
	v223 = v200
	goto L39
L38:
	;
	v284 = int32(-1)
	v286 = base.I32_div_s(v283, int32(2))
	if v280 <= v286 {
		v300 = v284
		v302 = v164
		goto L6
	} else {
		goto L52
	}
L39:
	;
	v228 = v211
	v229 = v212
	v232 = v212
	goto L41
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v253
	v280 = v269
	v283 = v270
	goto L38
L41:
	;
	if v232&int32(1) != 0 {
		v280 = v216
		v283 = v223
		goto L38
	} else {
		goto L43
	}
L42:
	;
	if v255 == int32(0) {
		v280 = v216
		v283 = v223
		goto L38
	} else {
		goto L45
	}
L43:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v244 = int32(1)
	v245 = v228 - v244
	v249 = base.B2i32(v243&(v245^v206) == int32(0))
	v250 = v249 | v229
	v253 = v245 & v243
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v217)+20))
	v255 = v228*int32(48) + v254
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+4)))
	if v256 != v244 {
		v228 = v253
		v229 = v250
		v232 = v249
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+5)))
	if v261 != 0 {
		v211 = v253
		v212 = v250
		goto L39
	} else {
		goto L46
	}
L46:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	if v262&int32(255) == int32(0) {
		v211 = v253
		v212 = v250
		goto L39
	} else {
		goto L47
	}
L47:
	;
	F_tbm_mark_page_lossy(m, l0, v262)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L23
	} else {
		goto L48
	}
L48:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v272 = base.I32_div_s(v270, int32(2))
	if v272 < v269 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v211 = v253
	v212 = v250
	v216 = v269
	v217 = v274
	v223 = v270
	goto L39
L50:
	;
	goto L51
L51:
	;
	goto L40
L52:
	;
	v288 = int32(1073741823)
	if v288 <= v280 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v291 = v288
	goto L55
L54:
	;
	v291 = v280
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v291 << (uint(int32(1)) % 32)
	v300 = v284
	v302 = v164
	goto L6
L56:
	;
	goto L5
}
func F_tbm_union_page(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
	if v12 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v226 <= v227 {
		goto L1
	} else {
		goto L35
	}
L3:
	;
	v21 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v73 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(8)+v21<<(uint(int32(2))%32))))
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L2
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v38 = v31
	v39 = v32 + v21<<(uint(int32(5))%32)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v69 = v21 + int32(1)
	if v69 != int32(8) {
		v21 = v69
		goto L6
	} else {
		goto L19
	}
L11:
	;
	if v38&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	F_tbm_mark_page_lossy(m, l0, v39)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v51 = int32(1)
	if base.Ui32(v51) < base.Ui32(v38) {
		v38 = int32(base.Ui32(v38) >> (uint(v51) % 32))
		v39 = v39 + v51
		goto L11
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	goto L12
L19:
	;
	goto L7
L20:
	;
	v144 = F_tbm_get_pageentry(m, l0, v72)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L16
	} else {
		goto L31
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+20))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v80 = v72 & int32(-256)
	v81 = int32(16)
	v85 = (v80 ^ int32(base.Ui32(v72)>>(uint(v81)%32))) * int32(-2048144789)
	v90 = (int32(base.Ui32(v85)>>(uint(int32(13))%32)) ^ v85) * int32(-1028477387)
	v94 = v78 & (int32(base.Ui32(v90)>>(uint(v81)%32)) ^ v90)
	v97 = v77 + v94*int32(48)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v98 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v103 = v97
	v104 = v94
	goto L23
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v80 != v112 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+5)))
	if v121 != int32(1) {
		goto L20
	} else {
		goto L29
	}
L25:
	;
	v116 = (v104 + int32(1)) & v78
	v119 = v77 + v116*int32(48)
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+4)))
	if v120 != 0 {
		v103 = v119
		v104 = v116
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L20
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v103+int32(base.Ui32(v72)>>(uint(int32(3))%32))&int32(28))+8))
	if int32(base.Ui32(v129)>>(uint(v72)%32))&int32(1) != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L20
L31:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+5)))
	if v146 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v149 | v150
	v154 = v144 + int32(12)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v154))) = v155 | v156
	v160 = v144 + int32(16)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v160))) = v161 | v162
	v166 = v144 + int32(20)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v167 | v168
	v172 = v144 + int32(24)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v173 | v174
	v178 = v144 + int32(28)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v179 | v180
	v184 = v144 + int32(32)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v185 | v186
	v190 = v144 + int32(36)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v191 | v192
	v196 = v144 + int32(40)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v197 | v198
	v202 = v144 + int32(44)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v203 | v204
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+6)))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+6)))
	v209 = v207 | v208
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+6)) = uint8(v209)
	goto L2
L33:
	;
	goto L34
L34:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v211 | int32(1)
	goto L2
L35:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v233 = v231 & v232
	v236 = v233
	v237 = int32(0)
	v238 = v230
	v242 = v226
	v243 = v227
	goto L37
L36:
	;
	v300 = base.I32_div_s(v298, int32(2))
	if v297 <= v300 {
		goto L1
	} else {
		goto L50
	}
L37:
	;
	v247 = v236
	v248 = v237
	v251 = v237
	goto L39
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v268
	v297 = v284
	v298 = v285
	goto L36
L39:
	;
	if v251&int32(1) != 0 {
		v297 = v242
		v298 = v243
		goto L36
	} else {
		goto L41
	}
L40:
	;
	if v270 == int32(0) {
		v297 = v242
		v298 = v243
		goto L36
	} else {
		goto L43
	}
L41:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v259 = int32(1)
	v260 = v247 - v259
	v264 = base.B2i32(v258&(v260^v233) == int32(0))
	v265 = v264 | v248
	v268 = v260 & v258
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	v270 = v247*int32(48) + v269
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+4)))
	if v271 != v259 {
		v247 = v268
		v248 = v265
		v251 = v264
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+5)))
	if v276 != 0 {
		v236 = v268
		v237 = v265
		goto L37
	} else {
		goto L44
	}
L44:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v277&int32(255) == int32(0) {
		v236 = v268
		v237 = v265
		goto L37
	} else {
		goto L45
	}
L45:
	;
	F_tbm_mark_page_lossy(m, l0, v277)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L16
	} else {
		goto L46
	}
L46:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v287 = base.I32_div_s(v285, int32(2))
	if v287 < v284 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v236 = v268
	v237 = v265
	v238 = v289
	v242 = v284
	v243 = v285
	goto L37
L48:
	;
	goto L49
L49:
	;
	goto L38
L50:
	;
	v302 = int32(1073741823)
	if v302 <= v297 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v305 = v302
	goto L53
L52:
	;
	v305 = v297
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v305 << (uint(int32(1)) % 32)
	goto L1
}
