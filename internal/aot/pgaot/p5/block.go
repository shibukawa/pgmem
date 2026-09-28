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
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
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
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
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
	v150 = v48 << (uint(int32(1)) % 32)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150+v151))))
	if v153 == int32(0) {
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
	v57 = v54
	goto L12
L12:
	;
	v69 = v57 << (uint(int32(1)) % 32)
	if base.Ui32(v57) <= base.Ui32(v48) {
		v57 = v69
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
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v57
	goto L8
L16:
	;
	v74 = F_palloc0_mul(m, int32(2), v57)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v86 = F_repalloc(m, v85, v69)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v74
	v78 = F_palloc0_mul(m, int32(2), v57)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v78
	v82 = F_palloc0_mul(m, int32(4), v57)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v82
	goto L15
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v86
	v89 = v57 - v49
	v91 = v89 << (uint(int32(1)) % 32)
	v92 = int32(0)
	v93 = base.B2i32(v91 == v92)
	if v93 == v92 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	base.MemoryFill(m, v86+v96<<(uint(int32(1))%32), int32(0), v91)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v103 = F_repalloc(m, v102, v69)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v103
	if v93 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	base.MemoryFill(m, v103+v108<<(uint(int32(1))%32), int32(0), v91)
	goto L29
L28:
	;
	goto L29
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v117 = F_repalloc(m, v114, v57<<(uint(int32(2))%32))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+36)) = v117
	v121 = v89 << (uint(int32(2)) % 32)
	if v121 == int32(0) {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v36)+24))
	base.MemoryFill(m, v117+v124<<(uint(int32(2))%32), int32(0), v121)
	goto L15
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_BlockRefTableMarkBlockModified[0])) = v19
	m.G0 = v16 + int32(48)
	return
L33:
	;
	v158 = F_palloc_mul(m, int32(2), int32(16))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v178 = v4 & int32(_a_F_BlockRefTableMarkBlockModified_1)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179+v150))))
	if v181 != int32(_a_F_BlockRefTableMarkBlockModified_2) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v161 = v48 << (uint(int32(2)) % 32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v161+v162))) = v158
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v167 = int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v165+v150))) = uint16(v167)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169+v161)))
	*(*uint16)(unsafe.Add(mBase, uint32(v171))) = uint16(v4)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v175 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v173+v150))) = uint16(v175)
	goto L32
L37:
	;
	goto L32
L38:
	;
	if v181 == v153 {
		goto L57
	} else {
		goto L58
	}
L39:
	;
	v210 = int32(0)
	goto L44
L40:
	;
	if v181 == int32(0) {
		goto L38
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v192+v48<<(uint(int32(2))%32))))
	v201 = v196 + int32(base.Ui32(v178)>>(uint(int32(3))%32))&int32(_a_F_BlockRefTableMarkBlockModified_3)
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201))))
	v207 = v202 | int32(1)<<(uint(v4&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v201))) = uint16(v207)
	goto L32
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v186+v48<<(uint(int32(2))%32))))
	goto L39
L44:
	;
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190+v210<<(uint(int32(1))%32)))))
	if v225 == v178 {
		goto L37
	} else {
		goto L46
	}
L45:
	;
	if v181 != int32(4095) {
		goto L38
	} else {
		goto L48
	}
L46:
	;
	v228 = v210 + int32(1)
	if v228 != v181 {
		v210 = v228
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v233 = F_palloc0(m, int32(_a_F_BlockRefTableMarkBlockModified_4))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v236 = v48 << (uint(int32(1)) % 32)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v236+v237))))
	if v239 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v244 = int32(0)
	goto L53
L51:
	;
	goto L52
L52:
	;
	v298 = v233 + int32(base.Ui32(v178)>>(uint(int32(3))%32))&int32(_a_F_BlockRefTableMarkBlockModified_3)
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298))))
	v304 = v299 | int32(1)<<(uint(v4&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v298))) = uint16(v304)
	v307 = v48 << (uint(int32(2)) % 32)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v307+v308)))
	F_pfree(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L56
	}
L53:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256+v48<<(uint(int32(2))%32))))
	v259 = int32(1)
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258+v244<<(uint(v259)%32)))))
	v267 = v233 + int32(base.Ui32(v262)>>(uint(int32(3))%32))&int32(_a_F_BlockRefTableMarkBlockModified_3)
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v267))))
	v273 = v268 | v259<<(uint(v262&int32(15))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v267))) = uint16(v273)
	v276 = v244 + v259
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277+v236))))
	if base.Ui32(v276) < base.Ui32(v279) {
		v244 = v276
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
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v313+v307))) = v233
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v318 = int32(_a_F_BlockRefTableMarkBlockModified_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v316+v236))) = uint16(v318)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	*(*uint16)(unsafe.Add(mBase, uint32(v320+v236))) = uint16(v318)
	goto L32
L57:
	;
	v338 = int32(2)
	v339 = v48 << (uint(v338) % 32)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v339+v340)))
	v345 = F_repalloc(m, v342, v153<<(uint(v338)%32))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	v362 = v181
	goto L59
L59:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v363+v48<<(uint(int32(2))%32))))
	v368 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v367+v362<<(uint(v368)%32)))) = uint16(v4)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v375 = v372 + v48<<(uint(v368)%32)
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375))))
	v378 = v376 + v368
	*(*uint16)(unsafe.Add(mBase, uint32(v375))) = uint16(v378)
	goto L37
L60:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v36)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v347+v339))) = v345
	v350 = int32(1)
	v351 = v48 << (uint(v350) % 32)
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v36)+28))
	v355 = v153 << (uint(v350) % 32)
	*(*uint16)(unsafe.Add(mBase, uint32(v351+v352))) = uint16(v355)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v36)+32))
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v357+v351))))
	v362 = v359
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
