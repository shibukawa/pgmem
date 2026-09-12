package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CatalogCacheCreateEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v256 int32
	_ = v256
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
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
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v539 int32
	_ = v539
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v613 int32
	_ = v613
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v666 int32
	_ = v666
	var v680 int32
	_ = v680
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v703 int32
	_ = v703
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v759 int32
	_ = v759
	var v785 int32
	_ = v785
	var v818 int32
	_ = v818
	var v831 int32
	_ = v831
	var v857 int32
	_ = v857
	var v863 int32
	_ = v863
	var v864 int64
	_ = v864
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v893 int32
	_ = v893
	var v905 int32
	_ = v905
	v6 = int32(0)
	v32 = m.G0
	v34 = v32 + int32(-64)
	m.G0 = v34
	v43 = int32(-1)
	v45 = v6
	v46 = v6
	v48 = v6
	v49 = v6
	v50 = v6
	v51 = v6
	v52 = v6
	v53 = v6
	v55 = v6
	v56 = v6
	v57 = v6
	v59 = v6
	v64 = v34
	goto L2
L1:
	;
	m.G0 = v34 - int32(-64)
	return v905
L2:
	;
	goto L4
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v650
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v653
	v905 = v513
	goto L1
L4:
	;
	if v43 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L3
L6:
	;
	v863 = int32(m.ExcTag)
	v864 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v863 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+64)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v513))) = int32(1462113538)
	v539 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v513)+60)) = v539
	*(*uint8)(unsafe.Add(mBase, uint32(v513)+36)) = uint8(v539)
	*(*int32)(unsafe.Add(mBase, uint32(v513)+32)) = v539
	*(*int32)(unsafe.Add(mBase, uint32(v513)+4)) = l3
	v547 = v525 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v513)+37)) = uint8(v547)
	v550 = v513 + int32(24)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v552 = v551 + l4<<(uint(int32(3))%32)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+4))
	if v553 == v539 {
		goto L64
	} else {
		goto L65
	}
L8:
	;
	v358 = int32(4553888)
	v359 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v362 = *(*int32)(unsafe.Add(mBase, _consts[510]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v362
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v76
	v376 = F_palloc(m, int32(68))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		v857 = v82
		goto L6
	} else {
		goto L52
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v117
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v159)
	F_pfree(m, v164)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		v857 = v126
		goto L6
	} else {
		goto L51
	}
L10:
	;
	v185 = int32(4553888)
	v186 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v189 = *(*int32)(unsafe.Add(mBase, _consts[510]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v176
	v197 = v182 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v197)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v56
	v207 = F_palloc(m, v191+int32(76))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		v857 = v184
		goto L6
	} else {
		goto L28
	}
L11:
	;
	v73 = v64 + int32(-64)
	m.G0 = v73
	v75 = int32(16)
	v76 = v73 - v75
	m.G0 = v76
	v79 = v76 - int32(160)
	m.G0 = v79
	v82 = v79 - v75
	m.G0 = v82
	v85 = base.B2i32(l1 == int32(0))
	if l1 == int32(0) {
		goto L8
	} else {
		goto L14
	}
L12:
	;
	v116 = v45
	v117 = v46
	v118 = v48
	v119 = v49
	v120 = v50
	v121 = v51
	v122 = v52
	v123 = v53
	v124 = v59
	v126 = v64
	goto L13
L13:
	;
	if v116 != 0 {
		goto L22
	} else {
		goto L23
	}
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+20)))
	if v89&int32(4) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v173 = l1
	v175 = v76
	v176 = v82
	v177 = v49
	v178 = v79
	v179 = v51
	v180 = v52
	v181 = v53
	v182 = v85
	v184 = v82
	goto L10
L16:
	;
	goto L17
L17:
	;
	v94 = int32(0)
	v95 = int32(4540648)
	v96 = *(*int32)(unsafe.Add(mBase, _consts[1074]))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+8)) = uint16(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v96
	*(*int32)(unsafe.Add(mBase, _consts[1074])) = v76
	v105 = *(*int32)(unsafe.Add(mBase, _consts[53]))
	v107 = *(*int32)(unsafe.Add(mBase, _consts[52]))
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v32 + int32(-48)
	goto L21
L19:
	;
	v116 = v94
	v117 = v76
	v118 = v82
	v119 = v96
	v120 = v79
	v121 = v105
	v122 = v107
	v123 = v76 + int32(9)
	v124 = v85
	v126 = v82
	goto L13
L21:
	;
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v122
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v121
	*(*int32)(unsafe.Add(mBase, _consts[1074])) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v117
	v144 = v124 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v144)
	F_pg_re_throw(m)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		v857 = v126
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v120
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v119
	v159 = v124 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v159)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v117
	v164 = F_toast_flatten_tuple(m, l1, v150)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		v857 = v126
		goto L6
	} else {
		goto L26
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _consts[53])) = v121
	*(*int32)(unsafe.Add(mBase, _consts[1074])) = v119
	*(*int32)(unsafe.Add(mBase, _consts[52])) = v122
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if v172 != 0 {
		goto L9
	} else {
		goto L27
	}
L27:
	;
	v173 = v164
	v175 = v117
	v176 = v118
	v177 = v119
	v178 = v120
	v179 = v121
	v180 = v122
	v181 = v123
	v182 = v124
	v184 = v126
	goto L10
L28:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+40)) = v209
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+44)) = v211
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v207)+48)) = uint16(v213)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v219 = (v207 + int32(75)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(v207)+56)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v207)+52)) = v215
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	if v223 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v186
	if l1 != v173 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v224 = F__emscripten_memcpy_bulkmem(m, v219, v222, v223)
	mBase = m.M
	goto L32
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v175
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v197)
	F_pfree(m, v173)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		v857 = v184
		goto L6
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v242 <= int32(0) {
		v512 = v175
		v513 = v207
		v514 = v176
		v515 = v177
		v516 = v178
		v517 = v179
		v518 = v180
		v519 = v181
		v521 = v55
		v522 = v56
		v523 = v57
		v525 = v182
		v530 = v184
		goto L7
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v246 = v207 + int32(40)
	v256 = int32(0)
	v268 = v55
	v269 = v56
	v270 = v57
	goto L38
L38:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v285 = v256 << (uint(int32(2)) % 32)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(48)+v285)))
	if int32(0) < v287 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	v512 = v175
	v513 = v207
	v514 = v176
	v515 = v177
	v516 = v178
	v517 = v179
	v518 = v180
	v519 = v181
	v521 = v334
	v522 = v335
	v523 = v336
	v525 = v182
	v530 = v184
	goto L7
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v285+(v207+int32(8))))) = v337
	v341 = v256 + int32(1)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v341 < v342 {
		v256 = v341
		v268 = v334
		v269 = v335
		v270 = v336
		goto L38
	} else {
		goto L50
	}
L41:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v207)+56))
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v290)+18)))
	if base.Ui32(v291&int32(2047)) < base.Ui32(v287) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v175
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v197)
	v332 = F_heap_getsysattr(m, v246, v287, v176)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		v857 = v184
		goto L6
	} else {
		goto L49
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v175
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v197)
	v306 = F_getmissingattr(m, v283, v287, v176)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		v857 = v184
		goto L6
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v268
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v180
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v179
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v175
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v197)
	v319 = F_fastgetattr_4(m, v246, v287, v283, v176)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		v857 = v184
		goto L6
	} else {
		goto L48
	}
L47:
	;
	v334 = v268
	v335 = v269
	v336 = v306
	v337 = v306
	goto L40
L48:
	;
	v334 = v268
	v335 = v319
	v336 = v270
	v337 = v319
	goto L40
L49:
	;
	v334 = v332
	v335 = v269
	v336 = v270
	v337 = v332
	goto L40
L50:
	;
	goto L39
L51:
	;
	v905 = int32(0)
	goto L1
L52:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if int32(0) < v378 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v393 = int32(0)
	goto L56
L54:
	;
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v359
	v512 = v76
	v513 = v376
	v514 = v82
	v515 = v49
	v516 = v79
	v517 = v51
	v518 = v52
	v519 = v53
	v521 = v55
	v522 = v56
	v523 = v57
	v525 = v85
	v530 = v82
	goto L7
L56:
	;
	v421 = v393 << (uint(int32(2)) % 32)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l2+v421)))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v385)))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v421+(l0+int32(48)))))
	v432 = v385 - int32(80) + v424<<(uint(int32(4))%32) + v429*int32(100)
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+68))
	if v433 == int32(19) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v76
	v448 = F_strncpy(m, v73, v423, int32(64))
	mBase = m.M
	v449 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v448)+63)) = uint8(v449)
	goto L61
L59:
	;
	v451 = v423
	goto L60
L60:
	;
	v452 = int32(*(*int16)(unsafe.Add(mBase, uint32(v432)+72)))
	v453 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+82)))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v49
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v85)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v76
	v466 = F_datumCopy(m, v451, v453, v452)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		v857 = v82
		goto L6
	} else {
		goto L62
	}
L61:
	;
	v451 = v73
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v421+(v376+int32(8))))) = v466
	v470 = v393 + int32(1)
	if v470 != v378 {
		v393 = v470
		goto L56
	} else {
		goto L63
	}
L63:
	;
	goto L57
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v552))) = v552
	v557 = v552
	goto L66
L65:
	;
	v557 = v553
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v513)+24)) = v552
	*(*int32)(unsafe.Add(mBase, uint32(v513)+28)) = v557
	*(*int32)(unsafe.Add(mBase, uint32(v557))) = v550
	*(*int32)(unsafe.Add(mBase, uint32(v552)+4)) = v550
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v563 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v562 + v563
	v567 = *(*int32)(unsafe.Add(mBase, _consts[1075]))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v567)+4)) = v568 + v563
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v572 <= v573<<(uint(v563)%32) {
		v905 = v513
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v512
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v547)
	v590 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		v857 = v530
		goto L6
	} else {
		goto L68
	}
L68:
	;
	if v590 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v515
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v547)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v595
	*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v593
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v592
	F_errmsg_internal(m, int32(132262), v34)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		v857 = v530
		goto L6
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v515
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v547)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v512
	v647 = *(*int32)(unsafe.Add(mBase, _consts[510]))
	v650 = F_MemoryContextAllocZero(m, v647, v634<<(uint(int32(4))%32))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		v857 = v530
		goto L6
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v512
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v547)
	F_errfinish(m, int32(522025), int32(992), int32(417509))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		v857 = v530
		goto L6
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v653 = v634 << (uint(int32(1)) % 32)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v654 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v666 = v654
	v680 = int32(0)
	goto L78
L76:
	;
	goto L77
L77:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v34)+20)) = v521
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v523
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v518
	*(*int32)(unsafe.Add(mBase, uint32(v34)+36)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v34)+40)) = v519
	*(*int32)(unsafe.Add(mBase, uint32(v34)+44)) = v515
	*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)) = uint8(v547)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+52)) = v514
	*(*int32)(unsafe.Add(mBase, uint32(v34)+56)) = v516
	*(*int32)(unsafe.Add(mBase, uint32(v34)+60)) = v512
	F_pfree(m, v818)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		v857 = v530
		goto L6
	} else {
		goto L90
	}
L78:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v694 = v691 + v680<<(uint(int32(3))%32)
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+4))
	if v695 == int32(0) {
		v759 = v666
		goto L80
	} else {
		goto L81
	}
L79:
	;
	goto L77
L80:
	;
	v785 = v680 + int32(1)
	if v785 < v759 {
		v666 = v759
		v680 = v785
		goto L78
	} else {
		goto L89
	}
L81:
	;
	if v695 == v694 {
		v759 = v666
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v703 = v695
	goto L83
L83:
	;
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v703-int32(20))))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v703)))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v703)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v733)+4)) = v734
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v703)))
	*(*int32)(unsafe.Add(mBase, uint32(v734))) = v736
	v741 = v650 + v732&(v653-int32(1))<<(uint(int32(3))%32)
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)+4))
	if v742 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v759 = v752
	goto L80
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v741))) = v741
	v746 = v741
	goto L87
L86:
	;
	v746 = v742
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v703))) = v741
	*(*int32)(unsafe.Add(mBase, uint32(v703)+4)) = v746
	*(*int32)(unsafe.Add(mBase, uint32(v746))) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v741)+4)) = v703
	if v694 != v734 {
		v703 = v734
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	goto L79
L90:
	;
	goto L5
L91:
	;
	v868 = int32(v864)
	m.G0 = v857
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v868)))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v871)))
	if v32+int32(-48) == v875 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	m.ExcPending = 1
	goto L100
L93:
	;
	if v878 != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v871)+4))
	v878 = v877
	goto L96
L95:
	;
	v878 = int32(0)
	goto L96
L96:
	;
	goto L93
L97:
	;
	v879 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+51)))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	v884 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v885 = *(*int32)(unsafe.Add(mBase, uint32(v34)+36))
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v34)+32))
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v34)+28))
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v34)+24))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v43 = v878
	v45 = v870
	v46 = v879
	v48 = v881
	v49 = v883
	v50 = v880
	v51 = v885
	v52 = v886
	v53 = v884
	v55 = v889
	v56 = v888
	v57 = v887
	v59 = v882
	v64 = v857
	goto L2
L98:
	;
	goto L99
L99:
	;
	F___wasm_longjmp(m, v871, v870)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	return int32(0)
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CatalogOpenIndexes(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	v4 = F_palloc0(m, int32(216))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+52)) = v8
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = int64(388)
		F_ExecOpenIndices(m, v4, v8)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v4
		}
	}
}
func F_GetCatalogSnapshot(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1205]))
	if v4 != 0 {
		v9 = v4
		return v9
	} else {
		v5 = F_GetNonHistoricCatalogSnapshot(m, l0)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = v5
			return v9
		}
	}
}
