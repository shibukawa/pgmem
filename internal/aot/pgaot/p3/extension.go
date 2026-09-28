package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ApplyExtensionUpdates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	v8 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(176)
	m.G0 = v23
	if l3 == v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v23 + int32(176)
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v27 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = l2
	v46 = v8
	goto L4
L4:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v46<<(uint(int32(2))%32))))
	v57 = F_palloc(m, int32(48))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+40)) = v59
	v61 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+32)) = v61
	v63 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+24)) = v63
	v65 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+16)) = v65
	v67 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = v67
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = v69
	F_parse_extension_control_file(m, v57, v55)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v75 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v78 = v23 + int32(112)
	F_ScanKeyInit(m, v78, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v85 = int32(1)
	v88 = F_systable_beginscan(m, v75, int32(3080), v85, int32(0), v85, v78)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	v240 = int32(3079)
	v243 = F_deleteDependencyRecordsForClass(m, v240, l0, v240, int32(110))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L6
	} else {
		goto L40
	}
L12:
	;
	v218 = int32(0)
	v228 = v218
	v232 = v218
	goto L11
L13:
	;
	v90 = F_systable_getnext(m, v88)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v90 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92+v93)+72))
	v96 = F_get_namespace_name(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L6
	} else {
		goto L37
	}
L18:
	;
	v98 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+104)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+88)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+56)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = int64(4294967296)
	v118 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v57)+32)))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v118
	v120 = F_cstring_to_text(m, v55)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v122 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+37)) = uint8(v122)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+88)) = base.I64_extend_i32_u(v120)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v75)+52))
	v133 = F_heap_modify_tuple(m, v90, v126, v23+int32(48), v23+int32(40), v23+int32(32))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	F_CatalogTupleUpdate(m, v75, v133+int32(4), v133)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	F_systable_endscan(m, v88)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	F_relation_close(m, v75, int32(3))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v57)+40))
	if v144 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v147 = int32(0)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v150 <= v147 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v161 = v147
	v162 = v147
	v165 = v147
	goto L26
L26:
	;
	v173 = int32(0)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v162<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v182 = F_get_required_extension(m, v179, v180, l4, l5, v173, l6)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L28
	}
L27:
	;
	v228 = v194
	v232 = v196
	goto L11
L28:
	;
	v185 = F_SearchSysCache1(m, int32(28), base.I64_extend_i32_u(v182))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	if v185 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)+16))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+22)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v187+v188)+72))
	F_ReleaseCatCache(m, v185)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	v193 = v173
	goto L32
L32:
	;
	v194 = F_lappend_oid(m, v161, v182)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L6
	} else {
		goto L34
	}
L33:
	;
	v193 = v190
	goto L32
L34:
	;
	v196 = F_lappend_oid(m, v165, v193)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v199 = v162 + int32(1)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v199 < v200 {
		v161 = v194
		v162 = v199
		v165 = v196
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L27
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l0
	F_errmsg_internal(m, int32(_a_F_ApplyExtensionUpdates_0), v23)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_ApplyExtensionUpdates_1), int32(3670), int32(_a_F_ApplyExtensionUpdates_2))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v245 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(3079)
	if v228 == v245 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyExtensionUpdates[0]))
	if v318 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v252 = int32(0)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v253 <= v252 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v265 = v252
	goto L44
L44:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v228)+12))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276+v265<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v280
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = int32(3079)
	F_recordDependencyOn(m, v23+int32(20), v23+int32(8), int32(110))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	goto L41
L46:
	;
	v294 = v265 + int32(1)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v228)+4))
	if v294 < v295 {
		v265 = v294
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v320 = int32(0)
	F_RunObjectPostAlterHook(m, int32(3079), l0, v320, v320, v320)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_execute_extension_script(m, l0, v57, v33, v55, v232, v96)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L6
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	v328 = v46 + int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v328 < v329 {
		v33 = v55
		v46 = v328
		goto L4
	} else {
		goto L53
	}
L53:
	;
	goto L5
}
func F_execute_extension_script(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v61 int64
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v203 int32
	_ = v203
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v344 int32
	_ = v344
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v377 int32
	_ = v377
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v470 int32
	_ = v470
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v535 int32
	_ = v535
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v684 int32
	_ = v684
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v789 int32
	_ = v789
	var v801 int32
	_ = v801
	var v815 int32
	_ = v815
	var v830 int32
	_ = v830
	var v831 int64
	_ = v831
	var v847 int32
	_ = v847
	var v860 int32
	_ = v860
	var v876 int32
	_ = v876
	var v891 int32
	_ = v891
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v920 int32
	_ = v920
	var v932 int32
	_ = v932
	var v948 int32
	_ = v948
	var v963 int32
	_ = v963
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v1002 int32
	_ = v1002
	var v1020 int32
	_ = v1020
	var v1032 int32
	_ = v1032
	var v1048 int32
	_ = v1048
	var v1063 int32
	_ = v1063
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1189 int64
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1284 int64
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1317 int32
	_ = v1317
	var v1330 int32
	_ = v1330
	var v1347 int32
	_ = v1347
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1367 int64
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1420 int64
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1434 int32
	_ = v1434
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1454 int32
	_ = v1454
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1486 int32
	_ = v1486
	var v1501 int32
	_ = v1501
	var v1504 int64
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1524 int32
	_ = v1524
	var v1536 int64
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1557 int32
	_ = v1557
	var v1560 int64
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1575 int32
	_ = v1575
	var v1586 int32
	_ = v1586
	var v1587 int32
	_ = v1587
	var v1602 int64
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int64
	_ = v1606
	var v1607 int64
	_ = v1607
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1633 int32
	_ = v1633
	var v1634 int32
	_ = v1634
	var v1659 int32
	_ = v1659
	var v1660 int32
	_ = v1660
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1700 int32
	_ = v1700
	var v1709 int32
	_ = v1709
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1729 int32
	_ = v1729
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1751 int32
	_ = v1751
	var v1762 int32
	_ = v1762
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1803 int32
	_ = v1803
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1832 int32
	_ = v1832
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1871 int32
	_ = v1871
	var v1872 int32
	_ = v1872
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1901 int32
	_ = v1901
	var v1915 int32
	_ = v1915
	var v1927 int32
	_ = v1927
	var v1939 int32
	_ = v1939
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1968 int32
	_ = v1968
	var v1981 int32
	_ = v1981
	var v1995 int32
	_ = v1995
	var v2010 int32
	_ = v2010
	var v2021 int32
	_ = v2021
	var v2027 int32
	_ = v2027
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2091 int32
	_ = v2091
	var v2123 int32
	_ = v2123
	var v2136 int32
	_ = v2136
	var v2142 int32
	_ = v2142
	var v2159 int32
	_ = v2159
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2222 int32
	_ = v2222
	var v2223 int32
	_ = v2223
	var v2234 int32
	_ = v2234
	var v2235 int32
	_ = v2235
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2262 int64
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2296 int32
	_ = v2296
	var v2309 int32
	_ = v2309
	var v2327 int32
	_ = v2327
	var v2342 int32
	_ = v2342
	var v2350 int32
	_ = v2350
	var v2363 int32
	_ = v2363
	var v2364 int32
	_ = v2364
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2400 int32
	_ = v2400
	var v2401 int64
	_ = v2401
	var v2405 int32
	_ = v2405
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2411 int32
	_ = v2411
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2417 int32
	_ = v2417
	var v2418 int32
	_ = v2418
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2425 int64
	_ = v2425
	var v2427 int32
	_ = v2427
	v7 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(576)
	m.G0 = v32
	v42 = int32(-1)
	v43 = v7
	v44 = v7
	v45 = v7
	v46 = v7
	v47 = v7
	v48 = v7
	v49 = v7
	v50 = v7
	v51 = v7
	v52 = v7
	v61 = int64(0)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v42 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v2400 = int32(m.ExcTag)
	v2401 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2400 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L7:
	;
	if v736 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L8:
	;
	v727 = v44
	v728 = v43
	v730 = v45
	v731 = v46
	v732 = v47
	v733 = v48
	v736 = v51
	v737 = v52
	goto L7
L9:
	;
	goto L10
L10:
	;
	v66 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v66
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v66
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+33)))
	if v71 != int32(1) {
		v252 = v66
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	v266 = v52 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	v268 = F_get_extension_script_filename(m, l1, l2, l3)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L37
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	v82 = int32(1)
	v83 = v52 & v82
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	v86 = v44 & v82
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	v88 = F_superuser(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v88 != 0 {
		v252 = v66
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	if v90 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[0]))
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	v119 = F_object_aclcheck(m, int32(1262), v104, v106, int64(512))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	if v119 == int32(0) {
		v252 = int32(1)
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	F_errcode(m, int32(16797828))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+176)) = v153
	F_errmsg(m, int32(_a_F_execute_extension_script_0), v32+int32(176))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+192)) = v153
	F_errmsg(m, int32(_a_F_execute_extension_script_1), v32+int32(192))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L6
	} else {
		goto L31
	}
L25:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	if v172 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v185 = int32(_a_F_execute_extension_script_2)
	goto L28
L27:
	;
	v185 = int32(_a_F_execute_extension_script_3)
	goto L28
L28:
	;
	F_errhint(m, v185, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1279), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L3
L31:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	if v220 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v233 = int32(_a_F_execute_extension_script_6)
	goto L34
L33:
	;
	v233 = int32(_a_F_execute_extension_script_7)
	goto L34
L34:
	;
	F_errhint(m, v233, int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v83)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v86)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1287), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	goto L3
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	v282 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	if l2 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v252 != 0 {
		goto L49
	} else {
		goto L50
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), v330, int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L6
	} else {
		goto L48
	}
L41:
	;
	if v282 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v282 == int32(0) {
		goto L39
	} else {
		goto L46
	}
L44:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+148)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v32)+144)) = v288
	F_errmsg_internal(m, int32(_a_F_execute_extension_script_8), v32+int32(144))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v330 = int32(1293)
	goto L40
L46:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+168)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v32)+164)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v32)+160)) = v309
	F_errmsg_internal(m, int32(_a_F_execute_extension_script_9), v32+int32(160))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v330 = int32(1295)
	goto L40
L48:
	;
	goto L39
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	v361 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(396)))) = v361
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v32+int32(392)))) = v364
	goto L52
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	v395 = int32(_a_F_execute_extension_script_10)
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[3]))
	v399 = v397 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[3])) = v399
	goto L54
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v32)+392))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[2])) = v377 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1])) = int32(10)
	goto L53
L53:
	;
	goto L51
L54:
	;
	v402 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[4]))
	if v402 <= int32(18) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	F_set_config_option(m, int32(_a_F_execute_extension_script_11), int32(_a_F_execute_extension_script_12), int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L6
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[5]))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v424<<(uint(int32(2))%32))+uint32(_c_F_execute_extension_script[6])))
	if v427 <= int32(18) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	F_set_config_option_ext(m, int32(_a_F_execute_extension_script_13), int32(_a_F_execute_extension_script_12), int32(5), int32(13), int32(10), int32(2), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v450 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[7])))
	if v450 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	F_set_config_option(m, int32(_a_F_execute_extension_script_14), int32(_a_F_execute_extension_script_15), int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	v482 = v32 + int32(376)
	F_initStringInfo(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L6
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	v495 = F_quote_identifier(m, l5)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v266)
	F_appendStringInfoString(m, v482, v495)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v510 = l4 + int32(4)
	v512 = base.B2i32(l4 == int32(0))
	if l4 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v512)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	F_appendStringInfoString(m, v32+int32(376), int32(_a_F_execute_extension_script_16))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L6
	} else {
		goto L89
	}
L71:
	;
	v515 = int32(0)
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v516 <= v515 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v535 = v515
	goto L73
L73:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548+v535<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v512)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	v563 = F_get_namespace_name(m, v552)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L70
L75:
	;
	v638 = v535 + int32(1)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v510)))
	if v638 < v639 {
		v535 = v638
		goto L73
	} else {
		goto L88
	}
L76:
	;
	if v563 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v512)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	v577 = int32(_a_F_execute_extension_script_17)
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	v583 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[8])))
	if base.B2i32(v580 == int32(0))|base.B2i32(v580 != v583) != 0 {
		v601 = v580
		v602 = v583
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v601-v602 == int32(0) {
		goto L75
	} else {
		goto L85
	}
L79:
	;
	goto L78
L80:
	;
	v586 = v563
	v587 = v577
	goto L81
L81:
	;
	v590 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v587)+1)))
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586)+1)))
	if v591 == int32(0) {
		v601 = v591
		v602 = v590
		goto L79
	} else {
		goto L83
	}
L82:
	;
	v601 = v591
	v602 = v590
	goto L79
L83:
	;
	v594 = int32(1)
	if v591 == v590 {
		v586 = v586 + v594
		v587 = v587 + v594
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v512)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	v616 = F_quote_identifier(m, v563)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v512)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+128)) = v616
	F_appendStringInfo(m, v32+int32(376), int32(_a_F_execute_extension_script_18), v32+int32(128))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	goto L75
L88:
	;
	goto L74
L89:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v46
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v512)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v268
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v252)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v32)+376))
	F_set_config_option(m, int32(_a_F_execute_extension_script_19), v696, int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[9])) = l0
	v707 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[10])) = uint8(v707)
	v710 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11]))
	v712 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[12]))
	goto L91
L91:
	;
	v714 = v32 + int32(208)
	*(*int32)(unsafe.Add(mBase, uint32(v714)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v714))) = v32 + int32(204)
	goto L94
L92:
	;
	v727 = v252
	v728 = v268
	v730 = v712
	v731 = v710
	v732 = v510
	v733 = v399
	v736 = int32(0)
	v737 = v512
	goto L7
L94:
	;
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11])) = v32 + int32(208)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	v763 = int32(1)
	v764 = v737 & v763
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	v767 = v727 & v763
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v773 = F___fstatat(m, int32(-100), v728, v32+int32(400), int32(0))
	mBase = m.M
	goto L98
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[12])) = v730
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11])) = v731
	v2350 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[10])) = uint8(v2350)
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[9])) = v2350
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	v2363 = int32(1)
	v2364 = v737 & v2363
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v2364)
	v2367 = v727 & v2363
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v2367)
	F_pg_re_throw(m)
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L6
	} else {
		goto L265
	}
L98:
	;
	if v773 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L6
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v831 = *(*int64)(unsafe.Add(mBase, uint32(v32)+424))
	if int64(1073741823) <= v831 {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errcode_for_file_access(m)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v728
	F_errmsg(m, int32(_a_F_execute_extension_script_20), v32)
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(4014), int32(_a_F_execute_extension_script_21))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	goto L3
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L6
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v903 = F_AllocateFile(m, v728, int32(_a_F_execute_extension_script_22))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L6
	} else {
		goto L113
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errcode(m, int32(261))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v728
	F_errmsg(m, int32(_a_F_execute_extension_script_23), v32+int32(16))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(4019), int32(_a_F_execute_extension_script_21))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	goto L3
L113:
	;
	if v903 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v974 = base.I32_wrap_i64(v831)
	v977 = F_palloc(m, v974+int32(1))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L6
	} else {
		goto L121
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errcode_for_file_access(m)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v728
	F_errmsg(m, int32(_a_F_execute_extension_script_24), v32+int32(32))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(4026), int32(_a_F_execute_extension_script_21))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	goto L3
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v990 = F_fread(m, v977, int32(1), v974, v903)
	mBase = m.M
	v991 = m.ExcPending
	if v991 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v903)))
	goto L123
L123:
	;
	if int32(base.Ui32(v1002)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L6
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1074 = F_FreeFile(m, v903)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L6
	} else {
		goto L131
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errcode_for_file_access(m)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+112)) = v728
	F_errmsg(m, int32(_a_F_execute_extension_script_25), v32+int32(112))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(4035), int32(_a_F_execute_extension_script_21))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	goto L3
L131:
	;
	v1077 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v977+v990))) = uint8(v1077)
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v1079 < v1077 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1093 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[13]))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v1093)+4))
	goto L135
L133:
	;
	v1095 = v50
	v1096 = v1079
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1108 = F_pg_verify_mbstr(m, v1096, v977, v990, int32(0))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L6
	} else {
		goto L136
	}
L135:
	;
	v1095 = v1094
	v1096 = v1094
	goto L134
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1120 = F_pg_any_to_server(m, v977, v990, v1096)
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1132 = F_cstring_to_text(m, v1120)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1145 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_26))
	mBase = m.M
	v1146 = m.ExcPending
	if v1146 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1158 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_27))
	mBase = m.M
	v1159 = m.ExcPending
	if v1159 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1171 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_28))
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1189 = F_DirectFunctionCall4Coll(m, int32(595), int32(950), base.I64_extend_i32_u(v1132), base.I64_extend_i32_u(v1145), base.I64_extend_i32_u(v1158), base.I64_extend_i32_u(v1171))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1202 = F_strstr(m, v1120, int32(_a_F_execute_extension_script_29))
	mBase = m.M
	if v1202 == int32(0) {
		v1363 = v49
		v1367 = v1189
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v1368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v1368 != 0 {
		v1504 = v1367
		goto L163
	} else {
		goto L164
	}
L144:
	;
	if v767 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1231 = F_GetUserNameFromId(m, v1219, int32(0))
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L6
	} else {
		goto L149
	}
L146:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v32)+396))
	v1218 = v49
	v1219 = v1205
	goto L145
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v49
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1217 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1]))
	v1218 = v1217
	v1219 = v1217
	goto L145
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1243 = F_quote_identifier(m, v1231)
	mBase = m.M
	v1244 = m.ExcPending
	if v1244 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1256 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_29))
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1268 = F_cstring_to_text(m, v1243)
	mBase = m.M
	v1269 = m.ExcPending
	if v1269 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1284 = F_DirectFunctionCall3Coll(m, int32(596), int32(950), v1189, base.I64_extend_i32_u(v1256), base.I64_extend_i32_u(v1268))
	mBase = m.M
	v1285 = m.ExcPending
	if v1285 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1297 = F_strcspn(m, v1231, int32(_a_F_execute_extension_script_30))
	mBase = m.M
	v1298 = v1297 + v1231
	v1300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1298))))
	if v1300 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v1301 == int32(0) {
		v1363 = v1218
		v1367 = v1284
		goto L143
	} else {
		goto L158
	}
L155:
	;
	v1301 = v1298
	goto L157
L156:
	;
	v1301 = int32(0)
	goto L157
L157:
	;
	goto L154
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1317 = m.ExcPending
	if v1317 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1330 = m.ExcPending
	if v1330 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+96)) = int32(_a_F_execute_extension_script_30)
	F_errmsg(m, int32(_a_F_execute_extension_script_31), v32+int32(96))
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1218
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1421), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	goto L3
L163:
	;
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v1524 = int32(0)
	v1536 = v1504
	goto L179
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1379 = F_quote_identifier(m, l5)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L6
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1392 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_32))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1404 = F_cstring_to_text(m, v1379)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1420 = F_DirectFunctionCall3Coll(m, int32(596), int32(950), v1367, base.I64_extend_i32_u(v1392), base.I64_extend_i32_u(v1404))
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	if v1367 == v1420 {
		v1504 = v1367
		goto L163
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1434 = F_strcspn(m, l5, int32(_a_F_execute_extension_script_30))
	mBase = m.M
	v1435 = v1434 + l5
	v1437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1435))))
	if v1437 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v1438 == int32(0) {
		v1504 = v1420
		goto L163
	} else {
		goto L174
	}
L171:
	;
	v1438 = v1435
	goto L173
L172:
	;
	v1438 = int32(0)
	goto L173
L173:
	;
	goto L170
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+84)) = int32(_a_F_execute_extension_script_30)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+80)) = v1468
	F_errmsg(m, int32(_a_F_execute_extension_script_33), v32+int32(80))
	mBase = m.M
	v1486 = m.ExcPending
	if v1486 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1445), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	goto L3
L179:
	;
	v1537 = int32(0)
	if v1506 == v1537 {
		v1547 = v1537
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if v764 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v1541 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+4))
	if v1541 <= v1524 {
		v1547 = int32(0)
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1506)+12))
	v1547 = v1543 + v1524<<(uint(int32(2))%32)
	goto L181
L184:
	;
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(v1557+v1524<<(uint(int32(2))%32))))
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v1547)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v2194 = F_get_namespace_name(m, v2182)
	mBase = m.M
	v2195 = m.ExcPending
	if v2195 != 0 {
		goto L6
	} else {
		goto L248
	}
L185:
	;
	v1550 = int32(0)
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v732)))
	if base.B2i32(v1547 == v1550)|base.B2i32(v1552 <= v1524) == v1550 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v1560 = v1504
	goto L187
L187:
	;
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1561 != 0 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v1557 != 0 {
		goto L184
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1560 = v1536
	goto L187
L191:
	;
	goto L190
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1573 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_34))
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L6
	} else {
		goto L195
	}
L193:
	;
	v1606 = v61
	v1607 = v1560
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1619 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v1607))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L6
	} else {
		goto L198
	}
L195:
	;
	v1575 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1586 = F_cstring_to_text(m, v1575)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1602 = F_DirectFunctionCall3Coll(m, int32(596), int32(950), v1560, base.I64_extend_i32_u(v1573), base.I64_extend_i32_u(v1586))
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	v1606 = v1602
	v1607 = v1602
	goto L194
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1631 = F_text_to_cstring(m, v1619)
	mBase = m.M
	v1632 = m.ExcPending
	if v1632 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	v1633 = int32(_a_F_execute_extension_script_35)
	v1634 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[12])) = v32 + int32(500)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+520)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+516)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v32)+512)) = v1631
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+504)) = int32(597)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+500)) = v1634
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+508)) = v32 + int32(512)
	v1659 = F_pg_parse_query(m, v1631)
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L6
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1672 = F_CreateDestReceiver(m, int32(0))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	if v1659 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v2123 = *(*int32)(unsafe.Add(mBase, uint32(v32)+500))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[12])) = v2123
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2136 = m.ExcPending
	if v2136 != 0 {
		goto L6
	} else {
		goto L242
	}
L203:
	;
	v1676 = int32(0)
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1659)+4))
	if v1677 <= v1676 {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1700 = v1676
	goto L205
L205:
	;
	v1709 = *(*int32)(unsafe.Add(mBase, uint32(v1659)+12))
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1709+v1700<<(uint(int32(2))%32))))
	v1714 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+520)) = v1714
	v1716 = *(*int32)(unsafe.Add(mBase, uint32(v1713)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+524)) = v1716
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1729 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[14]))
	v1734 = F_AllocSetContextCreateInternal(m, v1729, int32(_a_F_execute_extension_script_36), int32(0), int32(_a_F_execute_extension_script_37), int32(_a_F_execute_extension_script_38))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L6
	} else {
		goto L207
	}
L206:
	;
	goto L202
L207:
	;
	v1736 = int32(_a_F_execute_extension_script_39)
	v1737 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[14]))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[14])) = v1734
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1751 = m.ExcPending
	if v1751 != 0 {
		goto L6
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1762 = int32(0)
	v1765 = F_pg_analyze_and_rewrite_fixedparams(m, v1713, v1631, v1762, v1762, v1762)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1779 = F_pg_plan_queries(m, v1765, v1631, int32(2048), int32(0))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L6
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[14])) = v1737
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_MemoryContextDelete(m, v1734)
	mBase = m.M
	v2088 = m.ExcPending
	if v2088 != 0 {
		goto L6
	} else {
		goto L240
	}
L211:
	;
	if v1779 == int32(0) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v1783 = int32(0)
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v1779)+4))
	if v1784 <= v1783 {
		goto L210
	} else {
		goto L213
	}
L213:
	;
	v1803 = v1783
	goto L214
L214:
	;
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1779)+12))
	v1820 = *(*int32)(unsafe.Add(mBase, uint32(v1816+v1803<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L6
	} else {
		goto L216
	}
L215:
	;
	goto L210
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1843 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L6
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_PushActiveSnapshot(m, v1843)
	mBase = m.M
	v1856 = m.ExcPending
	if v1856 != 0 {
		goto L6
	} else {
		goto L218
	}
L218:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1820)+100))
	if v1857 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2041 = m.ExcPending
	if v2041 != 0 {
		goto L6
	} else {
		goto L238
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1871 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[15]))
	v1872 = *(*int32)(unsafe.Add(mBase, uint32(v1871)))
	goto L223
L221:
	;
	goto L222
L222:
	;
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(v1857)))
	if v1952 == int32(225) {
		goto L230
	} else {
		goto L231
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v1883 = int32(0)
	v1887 = F_CreateQueryDesc(m, v1820, v1631, v1872, v1883, v1672, v1883, v1883, v1883)
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L6
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_ExecutorStart(m, v1887, int32(0))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_ExecutorRun(m, v1887, int32(1), int64(0))
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L6
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_ExecutorFinish(m, v1887)
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L6
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_ExecutorEnd(m, v1887)
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L6
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_FreeQueryDesc(m, v1887)
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	goto L219
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1968 = m.ExcPending
	if v1968 != 0 {
		goto L6
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v2021 = int32(0)
	F_ProcessUtility(m, v1820, v1631, v2021, int32(1), v2021, v2021, v1672, v2021)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L6
	} else {
		goto L237
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errcode(m, int32(1088))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
		goto L6
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errmsg(m, int32(_a_F_execute_extension_script_40), int32(0))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L6
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1193), int32(_a_F_execute_extension_script_41))
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L6
	} else {
		goto L236
	}
L236:
	;
	goto L3
L237:
	;
	goto L219
L238:
	;
	v2043 = v1803 + int32(1)
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v1779)+4))
	if v2043 < v2044 {
		v1803 = v2043
		goto L214
	} else {
		goto L239
	}
L239:
	;
	goto L215
L240:
	;
	v2090 = v1700 + int32(1)
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(v1659)+4))
	if v2090 < v2091 {
		v1700 = v2090
		goto L205
	} else {
		goto L241
	}
L241:
	;
	goto L206
L242:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11])) = v731
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[12])) = v730
	v2142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[10])) = uint8(v2142)
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[9])) = v2142
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_AtEOXact_GUC(m, int32(1), v733)
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L6
	} else {
		goto L243
	}
L243:
	;
	if v767 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v1606
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v32)+396))
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v32)+392))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[2])) = v2171
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1])) = v2170
	goto L247
L245:
	;
	goto L246
L246:
	;
	m.G0 = v32 + int32(576)
	return
L247:
	;
	goto L246
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v2206 = F_quote_identifier(m, v2194)
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L6
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+64)) = v2183
	v2222 = F_psprintf(m, int32(_a_F_execute_extension_script_42), v32-int32(-64))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L6
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v2234 = F_cstring_to_text(m, v2222)
	mBase = m.M
	v2235 = m.ExcPending
	if v2235 != 0 {
		goto L6
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v2246 = F_cstring_to_text(m, v2206)
	mBase = m.M
	v2247 = m.ExcPending
	if v2247 != 0 {
		goto L6
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v2262 = F_DirectFunctionCall3Coll(m, int32(596), int32(950), v1536, base.I64_extend_i32_u(v2234), base.I64_extend_i32_u(v2246))
	mBase = m.M
	v2263 = m.ExcPending
	if v2263 != 0 {
		goto L6
	} else {
		goto L254
	}
L253:
	;
	v1524 = v1524 + int32(1)
	v1536 = v2262
	goto L179
L254:
	;
	if v1536 == v2262 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	v2276 = F_strcspn(m, v2194, int32(_a_F_execute_extension_script_30))
	mBase = m.M
	v2277 = v2276 + v2194
	v2279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2277))))
	if v2279 != 0 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v2280 == int32(0) {
		goto L253
	} else {
		goto L260
	}
L257:
	;
	v2280 = v2277
	goto L259
L258:
	;
	v2280 = int32(0)
	goto L259
L259:
	;
	goto L256
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L6
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errcode(m, int32(33685634))
	mBase = m.M
	v2309 = m.ExcPending
	if v2309 != 0 {
		goto L6
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = int32(_a_F_execute_extension_script_30)
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v2183
	F_errmsg(m, int32(_a_F_execute_extension_script_33), v32+int32(48))
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L6
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+540)) = v1363
	*(*int64)(unsafe.Add(mBase, uint32(v32)+528)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v32)+544)) = v1095
	*(*int32)(unsafe.Add(mBase, uint32(v32)+548)) = v730
	*(*int32)(unsafe.Add(mBase, uint32(v32)+552)) = v731
	*(*int32)(unsafe.Add(mBase, uint32(v32)+560)) = v732
	*(*int32)(unsafe.Add(mBase, uint32(v32)+564)) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v32)+568)) = v728
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)) = uint8(v764)
	*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)) = uint8(v767)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1472), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v2342 = m.ExcPending
	if v2342 != 0 {
		goto L6
	} else {
		goto L264
	}
L264:
	;
	goto L3
L265:
	;
	goto L5
L266:
	;
	v2405 = int32(v2401)
	m.G0 = v32
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v2405)+4))
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v2405)))
	v2411 = *(*int32)(unsafe.Add(mBase, uint32(v2408)))
	if v32+int32(204) == v2411 {
		goto L269
	} else {
		goto L270
	}
L267:
	;
	m.ExcPending = 1
	goto L275
L268:
	;
	if v2415 != 0 {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2408)+4))
	v2415 = v2413
	goto L271
L270:
	;
	v2415 = int32(0)
	goto L271
L271:
	;
	goto L268
L272:
	;
	v2416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+575)))
	v2417 = *(*int32)(unsafe.Add(mBase, uint32(v32)+568))
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(v32)+564))
	v2419 = *(*int32)(unsafe.Add(mBase, uint32(v32)+560))
	v2420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+559)))
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v32)+552))
	v2422 = *(*int32)(unsafe.Add(mBase, uint32(v32)+548))
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v32)+544))
	v2424 = *(*int32)(unsafe.Add(mBase, uint32(v32)+540))
	v2425 = *(*int64)(unsafe.Add(mBase, uint32(v32)+528))
	v42 = v2415
	v43 = v2417
	v44 = v2416
	v45 = v2422
	v46 = v2421
	v47 = v2419
	v48 = v2418
	v49 = v2424
	v50 = v2423
	v51 = v2407
	v52 = v2420
	v61 = v2425
	goto L1
L273:
	;
	goto L274
L274:
	;
	F___wasm_longjmp(m, v2408, v2407)
	mBase = m.M
	v2427 = m.ExcPending
	if v2427 != 0 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	return
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
