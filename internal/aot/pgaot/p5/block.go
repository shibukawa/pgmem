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
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	v4 = l3
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = int32(_a_F_BlockRefTableMarkBlockModified_0)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_BlockRefTableMarkBlockModified[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_BlockRefTableMarkBlockModified[0])) = v21
	v23 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = l2
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v23
	v36 = F_blockreftable_insert(m, v28, v16+int32(8), v16+int32(31))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+31)))
	if v38 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v41 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+32)) = v41
	goto L5
L4:
	;
	goto L5
L5:
	;
	v48 = int32(base.Ui32(v4) >> (uint(int32(16)) % 32))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	if base.Ui32(v49) <= base.Ui32(v48) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v51 = int32(16)
	if base.Ui32(v49) <= base.Ui32(v51) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v149 = v48 << (uint(int32(1)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v149+v150))))
	if v152 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L9:
	;
	v54 = v51
	goto L11
L10:
	;
	v54 = v49
	goto L11
L11:
	;
	v55 = v54
	goto L12
L12:
	;
	v69 = v55 << (uint(int32(1)) % 32)
	if base.Ui32(v55) <= base.Ui32(v48) {
		v55 = v69
		goto L12
	} else {
		goto L14
	}
L13:
	;
	if v49 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v55
	goto L8
L16:
	;
	v73 = F_palloc0(m, v69)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v85 = F_repalloc(m, v84, v69)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v73
	v76 = F_palloc0(m, v69)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v76
	v81 = F_palloc0(m, v55<<(uint(int32(2))%32))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v81
	goto L15
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v85
	v88 = v55 - v49
	v90 = v88 << (uint(int32(1)) % 32)
	v91 = int32(0)
	v92 = base.B2i32(v90 == v91)
	if v92 == v91 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	base.MemoryFill(m, v85+v95<<(uint(int32(1))%32), int32(0), v90)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v102 = F_repalloc(m, v101, v69)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v102
	if v92 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	base.MemoryFill(m, v102+v107<<(uint(int32(1))%32), int32(0), v90)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v116 = F_repalloc(m, v113, v55<<(uint(int32(2))%32))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v116
	v120 = v88 << (uint(int32(2)) % 32)
	if v120 == int32(0) {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	base.MemoryFill(m, v116+v123<<(uint(int32(2))%32), int32(0), v120)
	goto L15
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BlockRefTableMarkBlockModified[0])) = v19
	m.G0 = v16 + int32(48)
	return
L33:
	;
	v156 = F_palloc(m, int32(32))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v176 = v4 & int32(_a_F_BlockRefTableMarkBlockModified_1)
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v177+v149))))
	if v179 != int32(_a_F_BlockRefTableMarkBlockModified_2) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v159 = v48 << (uint(int32(2)) % 32)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v159+v160))) = v156
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v165 = int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v163+v149))) = uint16(v165)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v167+v159)))
	*(*uint16)(unsafe.Add(mBase, uint32(v169))) = uint16(v4)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v173 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v171+v149))) = uint16(v173)
	goto L32
L37:
	;
	goto L32
L38:
	;
	if v179 == v152 {
		goto L57
	} else {
		goto L58
	}
L39:
	;
	v207 = int32(0)
	goto L44
L40:
	;
	if v179 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v190+v48<<(uint(int32(2))%32))))
	v199 = v194 + int32(base.Ui32(v176)>>(uint(int32(3))%32))&int32(_a_F_BlockRefTableMarkBlockModified_3)
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v199))))
	v205 = v200 | int32(1)<<(uint(v4&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v199))) = uint16(v205)
	goto L32
L43:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184+v48<<(uint(int32(2))%32))))
	goto L39
L44:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v188+v207<<(uint(int32(1))%32)))))
	if v223 == v176 {
		goto L37
	} else {
		goto L46
	}
L45:
	;
	if v179 != int32(4095) {
		goto L38
	} else {
		goto L48
	}
L46:
	;
	v226 = v207 + int32(1)
	if v226 != v179 {
		v207 = v226
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v231 = F_palloc0(m, int32(_a_F_BlockRefTableMarkBlockModified_4))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v234 = v48 << (uint(int32(1)) % 32)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234+v235))))
	if v237 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v241 = int32(0)
	goto L53
L51:
	;
	goto L52
L52:
	;
	v296 = v231 + int32(base.Ui32(v176)>>(uint(int32(3))%32))&int32(_a_F_BlockRefTableMarkBlockModified_3)
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v296))))
	v302 = v297 | int32(1)<<(uint(v4&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v296))) = uint16(v302)
	v305 = v48 << (uint(int32(2)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v305+v306)))
	F_pfree(m, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L56
	}
L53:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v254+v48<<(uint(int32(2))%32))))
	v257 = int32(1)
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256+v241<<(uint(v257)%32)))))
	v265 = v231 + int32(base.Ui32(v260)>>(uint(int32(3))%32))&int32(_a_F_BlockRefTableMarkBlockModified_3)
	v266 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v265))))
	v271 = v266 | v257<<(uint(v260&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v265))) = uint16(v271)
	v274 = v241 + v257
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275+v234))))
	if base.Ui32(v274) < base.Ui32(v277) {
		v241 = v274
		goto L53
	} else {
		goto L55
	}
L54:
	;
	goto L52
L55:
	;
	goto L54
L56:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v311+v305))) = v231
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v316 = int32(_a_F_BlockRefTableMarkBlockModified_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v314+v234))) = uint16(v316)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v318+v234))) = uint16(v316)
	goto L32
L57:
	;
	v336 = int32(2)
	v337 = v48 << (uint(v336) % 32)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v337+v338)))
	v343 = F_repalloc(m, v340, v152<<(uint(v336)%32))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	v359 = v179
	goto L59
L59:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360+v48<<(uint(int32(2))%32))))
	v365 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v364+v359<<(uint(v365)%32)))) = uint16(v4)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v372 = v369 + v48<<(uint(v365)%32)
	v373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v372))))
	v375 = v373 + v365
	*(*uint16)(unsafe.Add(mBase, uint32(v372))) = uint16(v375)
	goto L37
L60:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v345+v337))) = v343
	v348 = int32(1)
	v349 = v48 << (uint(v348) % 32)
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v353 = v152 << (uint(v348) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v349+v350))) = uint16(v353)
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v355+v349))))
	v359 = v357
	goto L59
}
func F_block_range_read_stream_cb(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v4) < base.Ui32(v5) {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v4 + int32(1)
		v11 = v4
	} else {
		v11 = int32(-1)
	}
	return v11
}
