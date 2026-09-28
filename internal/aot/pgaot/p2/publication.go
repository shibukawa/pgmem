package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetPublication(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(51), base.I64_extend_i32_u(l0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_GetPublication_0), v9)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_GetPublication_1), int32(1308), int32(_a_F_GetPublication_2))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
			v35 = F_palloc(m, int32(20))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v35))) = l0
				v38 = v32 + v33
				v41 = F_pstrdup(m, v38+int32(4))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v41
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+72)))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+8)) = uint8(v44)
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+73)))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+9)) = uint8(v46)
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+74)))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)) = uint8(v48)
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+75)))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+17)) = uint8(v50)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+76)))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)) = uint8(v52)
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+77)))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+19)) = uint8(v54)
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+78)))
					*(*uint8)(unsafe.Add(mBase, uint32(v35)+10)) = uint8(v56)
					v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v38)+79)))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v58
					F_ReleaseCatCache(m, v13)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v35
					}
				}
			}
		}
	}
}
func F_PublicationAddTables(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
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
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int64
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int64
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int64
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int64
	_ = v334
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v339 int32
	_ = v339
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v342 int64
	_ = v342
	var v346 int64
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v354 int64
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v418 int32
	_ = v418
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v443 int32
	_ = v443
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v504 int32
	_ = v504
	var v529 int32
	_ = v529
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v629 int32
	_ = v629
	var v631 int32
	_ = v631
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v686 int32
	_ = v686
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v742 int32
	_ = v742
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int64
	_ = v808
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v858 int64
	_ = v858
	var v861 int64
	_ = v861
	var v864 int32
	_ = v864
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(48)
	m.G0 = v23
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v23 + int32(48)
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v27 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v44 = v5
	goto L4
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v44<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+56))
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_PublicationAddTables[0]))
	v60 = F_object_ownercheck(m, int32(1259), v57, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	if v60 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65)+119)))
	switch v66 - int32(73) {
	case 0, 32:
		v76 = int32(20)
		goto L12
	default:
		goto L13
	case 10:
		goto L17
	case 29:
		goto L14
	case 36:
		goto L15
	case 45:
		goto L16
	}
L9:
	;
	goto L10
L10:
	;
	v85 = v23 + int32(36)
	v86 = m.G0
	v88 = v86 - int32(192)
	m.G0 = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+56))
	v92 = F_GetPublication(m, l0)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L19
	}
L11:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v56)+48))
	F_aclcheck_error(m, int32(2), v78, v79+int32(4))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L18
	}
L12:
	;
	v78 = v76
	goto L11
L13:
	;
	v76 = int32(42)
	goto L12
L14:
	;
	v78 = int32(18)
	goto L11
L15:
	;
	v78 = int32(23)
	goto L11
L16:
	;
	v78 = int32(52)
	goto L11
L17:
	;
	v78 = int32(38)
	goto L11
L18:
	;
	goto L10
L19:
	;
	v96 = F_table_open(m, int32(_a_F_PublicationAddTables_0), int32(3))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v99 = base.I64_extend_i32_u(v91)
	v100 = base.I64_extend_i32_u(l0)
	v101 = int64(0)
	v103 = F_SearchSysCacheExists(m, int32(53), v99, v100, v101, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L24
	}
L21:
	;
	if l3 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L6
	} else {
		goto L190
	}
L23:
	;
	m.G0 = v88 + int32(192)
	goto L21
L24:
	;
	if v103 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_relation_close(m, v96, int32(3))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+48))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+12)))
	if v137 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L28:
	;
	if l2 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_PublicationAddTables[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v109
	v112 = *(*int64)(unsafe.Add(mBase, _c_F_PublicationAddTables[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v112
	goto L23
L30:
	;
	goto L31
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(_a_F_PublicationAddTables_1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v90)+48))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v121 + int32(4)
	F_errmsg(m, int32(_a_F_PublicationAddTables_2), v88)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_PublicationAddTables_3), int32(575), int32(_a_F_PublicationAddTables_4))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+119)))
	switch v207 - int32(112) {
	case 0, 2:
		goto L56
	default:
		goto L57
	}
L37:
	;
	v204 = v136
	v205 = v136 + int32(4)
	v206 = int32(_a_F_PublicationAddTables_5)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v136)+68))
	v144 = F_get_namespace_name(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v135)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+96)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v88)+100)) = v146 + int32(4)
	v154 = F_psprintf(m, int32(_a_F_PublicationAddTables_6), v88+int32(96))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v135)+48))
	v157 = int32(_a_F_PublicationAddTables_7)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+12)))
	if v158 != int32(1) {
		v204 = v156
		v205 = v154
		v206 = v157
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156)+131)))
	if v161 != int32(1) {
		v204 = v156
		v205 = v154
		v206 = v157
		goto L36
	} else {
		goto L43
	}
L43:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v135)+56))
	v165 = F_PartitionHasPendingDetach(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v154
	F_errmsg(m, int32(_a_F_PublicationAddTables_7), v88+int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	if v165 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v184 = F_errdetail(m, int32(_a_F_PublicationAddTables_8), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v193 = F_errdetail(m, int32(_a_F_PublicationAddTables_9), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L53
	}
L51:
	;
	F_errfinish(m, int32(_a_F_PublicationAddTables_3), int32(93), int32(_a_F_PublicationAddTables_10))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errhint(m, int32(_a_F_PublicationAddTables_11), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_PublicationAddTables_3), int32(88), int32(_a_F_PublicationAddTables_10))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v135)+56))
	goto L63
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+32)) = v205
	F_errmsg(m, v206, v88+int32(32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L6
	} else {
		goto L60
	}
L60:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v135)+48))
	v223 = int32(*(*int8)(unsafe.Add(mBase, uint32(v222)+119)))
	F_errdetail_relkind_not_supported(m, v223)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_PublicationAddTables_3), int32(102), int32(_a_F_PublicationAddTables_10))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L6
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
	if base.Ui32(v231) < base.Ui32(int32(_a_F_PublicationAddTables_12)) {
		goto L22
	} else {
		goto L64
	}
L64:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v135)+48))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+118)))
	switch v235 - int32(116) {
	case 0:
		goto L67
	case 1:
		goto L66
	default:
		goto L65
	}
L65:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v282 = F_pub_collist_validate(m, v280, v281)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L6
	} else {
		goto L78
	}
L66:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L6
	} else {
		goto L73
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+64)) = v205
	F_errmsg(m, v206, v88-int32(-64))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L70
	}
L70:
	;
	v252 = F_errdetail(m, int32(_a_F_PublicationAddTables_13), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_PublicationAddTables_3), int32(116), int32(_a_F_PublicationAddTables_10))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L6
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L6
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+80)) = v205
	F_errmsg(m, v206, v88+int32(80))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L6
	} else {
		goto L75
	}
L75:
	;
	v273 = F_errdetail(m, int32(_a_F_PublicationAddTables_14), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L6
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_PublicationAddTables_3), int32(121), int32(_a_F_PublicationAddTables_10))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v284 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+184)) = v284
	*(*int64)(unsafe.Add(mBase, uint32(v88)+176)) = v284
	v288 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+136)) = v288
	*(*uint16)(unsafe.Add(mBase, uint32(v88)+140)) = uint16(v288)
	v294 = F_GetNewOidWithIndex(m, v96, int32(_a_F_PublicationAddTables_15), int32(1))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L6
	} else {
		goto L79
	}
L79:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88)+160)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v88)+152)) = v100
	*(*int64)(unsafe.Add(mBase, uint32(v88)+144)) = base.I64_extend_i32_u(v294)
	v300 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v55)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v88)+168)) = v300
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v302 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v311 != 0 {
		goto L87
	} else {
		goto L88
	}
L81:
	;
	v303 = F_nodeToString(m, v302)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v309 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+140)) = uint8(v309)
	goto L80
L84:
	;
	v305 = F_cstring_to_text(m, v303)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L85
	}
L85:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88)+176)) = base.I64_extend_i32_u(v305)
	goto L80
L86:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v96)+52))
	v556 = F_heap_form_tuple(m, v551, v88+int32(144), v88+int32(136))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L6
	} else {
		goto L134
	}
L87:
	;
	v312 = int32(0)
	v315 = int64(0)
	if v282 == v312 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	goto L89
L89:
	;
	v529 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+141)) = uint8(v529)
	goto L86
L90:
	;
	v360 = F_buildint2vector(m, v312, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L105
	}
L91:
	;
	v359 = int32(0)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v320 = v282 + int32(8)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v321 == int32(1) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	v359 = base.I32_popcnt(v324)
	goto L90
L95:
	;
	goto L96
L96:
	;
	v327 = v321 << (uint(int32(2)) % 32)
	if v327 <= int32(7) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v359 = base.I32_wrap_i64(v354)
	goto L90
L98:
	;
	if v327 == int32(0) {
		v354 = v315
		goto L97
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v351 = F_pg_popcount_optimized(m, v320, v327)
	mBase = m.M
	v354 = v351
	goto L97
L101:
	;
	v332 = v327
	v333 = v320
	v334 = v315
	goto L102
L102:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+3)))
	v336 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v335)+uint32(_c_F_PublicationAddTables[3]))))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+2)))
	v338 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v337)+uint32(_c_F_PublicationAddTables[3]))))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+1)))
	v340 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v339)+uint32(_c_F_PublicationAddTables[3]))))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	v342 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v341)+uint32(_c_F_PublicationAddTables[3]))))
	v346 = v336 + (v338 + (v340 + (v334 + v342)))
	v347 = int32(4)
	v350 = v332 - v347
	if v350 != 0 {
		v332 = v350
		v333 = v333 + v347
		v334 = v346
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v354 = v346
	goto L97
L104:
	;
	goto L103
L105:
	;
	if v282 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L106:
	;
	if int32(0) <= v418 {
		goto L117
	} else {
		goto L118
	}
L107:
	;
	v418 = base.I32_ctz(v404) | v405<<(uint(int32(5))%32)
	goto L106
L108:
	;
	v418 = int32(-2)
	goto L106
L109:
	;
	v369 = int32(0)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v372 <= v369 {
		goto L108
	} else {
		goto L110
	}
L110:
	;
	v375 = v282 + int32(8)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375)))
	v382 = v379 & int32(-1)
	if v382 != 0 {
		v404 = v382
		v405 = v369
		goto L107
	} else {
		goto L111
	}
L111:
	;
	v383 = int32(1)
	if v383 == v372 {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	v387 = v383
	goto L113
L113:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v375+v387<<(uint(int32(2))%32))))
	if v394 != 0 {
		v404 = v394
		v405 = v387
		goto L107
	} else {
		goto L115
	}
L114:
	;
	goto L108
L115:
	;
	v396 = v387 + int32(1)
	if v396 != v372 {
		v387 = v396
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v428 = v312
	v429 = v418
	goto L120
L118:
	;
	goto L119
L119:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v88)+184)) = base.I64_extend_i32_u(v360)
	goto L86
L120:
	;
	v443 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v360+int32(24)+v428<<(uint(v443)%32)))) = uint16(v429)
	if v282 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L119
L122:
	;
	if int32(0) <= v504 {
		v428 = v428 + v443
		v429 = v504
		goto L120
	} else {
		goto L133
	}
L123:
	;
	v504 = base.I32_ctz(v490) | v491<<(uint(int32(5))%32)
	goto L122
L124:
	;
	v504 = int32(-2)
	goto L122
L125:
	;
	v455 = v429 + int32(1)
	v457 = int32(base.Ui32(v455) >> (uint(int32(5)) % 32))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v458 <= v457 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v461 = v282 + int32(8)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v461+v457<<(uint(int32(2))%32))))
	v468 = v465 & (int32(-1) << (uint(v455) % 32))
	if v468 != 0 {
		v490 = v468
		v491 = v457
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v470 = v457 + int32(1)
	if v470 == v458 {
		goto L124
	} else {
		goto L128
	}
L128:
	;
	v473 = v470
	goto L129
L129:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v461+v473<<(uint(int32(2))%32))))
	if v480 != 0 {
		v490 = v480
		v491 = v473
		goto L123
	} else {
		goto L131
	}
L130:
	;
	goto L124
L131:
	;
	v482 = v473 + int32(1)
	if v482 != v458 {
		v473 = v482
		goto L129
	} else {
		goto L132
	}
L132:
	;
	goto L130
L133:
	;
	goto L121
L134:
	;
	F_CatalogTupleInsert(m, v96, v556)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	F_pfree(m, v556)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	v562 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+132)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v88)+128)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v88)+124)) = int32(_a_F_PublicationAddTables_0)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+120)) = v562
	*(*int32)(unsafe.Add(mBase, uint32(v88)+116)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v88)+112)) = int32(_a_F_PublicationAddTables_16)
	v573 = v88 + int32(124)
	v575 = v88 + int32(112)
	F_recordDependencyOn(m, v573, v575, int32(97))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+120)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+116)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v88)+112)) = int32(1259)
	F_recordDependencyOn(m, v573, v575, int32(97))
	mBase = m.M
	v586 = m.ExcPending
	if v586 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v587 != 0 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v589 = *(*int32)(unsafe.Add(mBase, _c_F_PublicationAddTables[0]))
	F_CheckUsageOnTypesInSingleRelExpr(m, v587, v91, v589)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if v282 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L142:
	;
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	F_recordDependencyOnSingleRelExpr(m, v573, v592, v91, int32(110), int32(0))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L6
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	if int32(0) <= v653 {
		goto L155
	} else {
		goto L156
	}
L145:
	;
	v653 = base.I32_ctz(v639) | v640<<(uint(int32(5))%32)
	goto L144
L146:
	;
	v653 = int32(-2)
	goto L144
L147:
	;
	v604 = int32(0)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v607 <= v604 {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v610 = v282 + int32(8)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v610)))
	v617 = v614 & int32(-1)
	if v617 != 0 {
		v639 = v617
		v640 = v604
		goto L145
	} else {
		goto L149
	}
L149:
	;
	v618 = int32(1)
	if v618 == v607 {
		goto L146
	} else {
		goto L150
	}
L150:
	;
	v622 = v618
	goto L151
L151:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v610+v622<<(uint(int32(2))%32))))
	if v629 != 0 {
		v639 = v629
		v640 = v622
		goto L145
	} else {
		goto L153
	}
L152:
	;
	goto L146
L153:
	;
	v631 = v622 + int32(1)
	if v631 != v607 {
		v622 = v631
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
L155:
	;
	v661 = v653
	goto L158
L156:
	;
	goto L157
L157:
	;
	F_relation_close(m, v96, int32(3))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L6
	} else {
		goto L173
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+120)) = v661
	*(*int32)(unsafe.Add(mBase, uint32(v88)+116)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v88)+112)) = int32(1259)
	F_recordDependencyOn(m, v88+int32(124), v88+int32(112), int32(110))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L6
	} else {
		goto L160
	}
L159:
	;
	goto L157
L160:
	;
	if v282 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L161:
	;
	if int32(0) <= v742 {
		v661 = v742
		goto L158
	} else {
		goto L172
	}
L162:
	;
	v742 = base.I32_ctz(v728) | v729<<(uint(int32(5))%32)
	goto L161
L163:
	;
	v742 = int32(-2)
	goto L161
L164:
	;
	v693 = v661 + int32(1)
	v695 = int32(base.Ui32(v693) >> (uint(int32(5)) % 32))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	if v696 <= v695 {
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v699 = v282 + int32(8)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v699+v695<<(uint(int32(2))%32))))
	v706 = v703 & (int32(-1) << (uint(v693) % 32))
	if v706 != 0 {
		v728 = v706
		v729 = v695
		goto L162
	} else {
		goto L166
	}
L166:
	;
	v708 = v695 + int32(1)
	if v708 == v696 {
		goto L163
	} else {
		goto L167
	}
L167:
	;
	v711 = v708
	goto L168
L168:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v699+v711<<(uint(int32(2))%32))))
	if v718 != 0 {
		v728 = v718
		v729 = v711
		goto L162
	} else {
		goto L170
	}
L169:
	;
	goto L163
L170:
	;
	v720 = v711 + int32(1)
	if v720 != v696 {
		v711 = v720
		goto L168
	} else {
		goto L171
	}
L171:
	;
	goto L169
L172:
	;
	goto L159
L173:
	;
	v768 = int32(0)
	if l3 == v768 {
		v778 = v768
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v779 = int32(1)
	v781 = int32(0)
	v783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+12)))
	if base.B2i32(v778&v779 == v781)&base.B2i32(v783 == v779) == v781 {
		goto L178
	} else {
		goto L179
	}
L175:
	;
	v771 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)))
	if v771 != int32(1) {
		v778 = v768
		goto L174
	} else {
		goto L176
	}
L176:
	;
	v774 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+20)))
	if v774 != int32(1) {
		v778 = v768
		goto L174
	} else {
		goto L177
	}
L177:
	;
	v777 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+12)))
	v778 = v777
	goto L174
L178:
	;
	v789 = F_get_rel_relkind(m, v91)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L6
	} else {
		goto L182
	}
L179:
	;
	goto L180
L180:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v88)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v806
	v808 = *(*int64)(unsafe.Add(mBase, uint32(v88)+124))
	*(*int64)(unsafe.Add(mBase, uint32(v85))) = v808
	goto L23
L181:
	;
	F_InvalidatePublicationRels(m, v803)
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L6
	} else {
		goto L189
	}
L182:
	;
	if v789 == int32(112) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v793 = int32(0)
	v796 = F_find_all_inheritors(m, v91, v793, v793)
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L6
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v801 = F_lappend_oid(m, int32(0), v91)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L6
	} else {
		goto L188
	}
L186:
	;
	v798 = F_list_concat(m, v793, v796)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L6
	} else {
		goto L187
	}
L187:
	;
	v803 = v798
	goto L181
L188:
	;
	v803 = v801
	goto L181
L189:
	;
	goto L180
L190:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L6
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+48)) = v205
	F_errmsg(m, v206, v88+int32(48))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L6
	} else {
		goto L192
	}
L192:
	;
	v847 = F_errdetail(m, int32(_a_F_PublicationAddTables_17), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L6
	} else {
		goto L193
	}
L193:
	;
	F_errfinish(m, int32(_a_F_PublicationAddTables_3), int32(109), int32(_a_F_PublicationAddTables_10))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L6
	} else {
		goto L194
	}
L194:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L195:
	;
	v883 = v44 + int32(1)
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v883 < v884 {
		v44 = v883
		goto L4
	} else {
		goto L200
	}
L196:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v856
	v858 = *(*int64)(unsafe.Add(mBase, uint32(v23)+36))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+24)) = v858
	v861 = *(*int64)(unsafe.Add(mBase, _c_F_PublicationAddTables[2]))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v861
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_PublicationAddTables[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v864
	F_EventTriggerCollectSimpleCommand(m, v23+int32(24), v23+int32(8), l3)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	v873 = *(*int32)(unsafe.Add(mBase, _c_F_PublicationAddTables[4]))
	if v873 == int32(0) {
		goto L195
	} else {
		goto L198
	}
L198:
	;
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v878 = int32(0)
	F_RunObjectPostCreateHook(m, int32(_a_F_PublicationAddTables_0), v877, v878, v878)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	goto L195
L200:
	;
	goto L5
}
