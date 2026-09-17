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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int64
	_ = v95
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
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
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	v8 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(128)
	m.G0 = v22
	if l3 == v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v22 + int32(128)
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v26 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = l2
	v44 = v8
	goto L4
L4:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v44<<(uint(int32(2))%32))))
	v54 = F_palloc(m, int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
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
	v56 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+40)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v60
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v64
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v66
	F_parse_extension_control_file(m, v54, v52)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v72 = F_table_open(m, int32(3079), int32(3))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v75 = v22 + int32(80)
	F_ScanKeyInit(m, v75, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v82 = int32(1)
	v85 = F_systable_beginscan(m, v72, int32(3080), v82, int32(0), v82, v75)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	v225 = int32(3079)
	v228 = F_deleteDependencyRecordsForClass(m, v225, l0, v225, int32(110))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L40
	}
L12:
	;
	v204 = int32(0)
	v214 = v204
	v218 = v204
	goto L11
L13:
	;
	v87 = F_systable_getnext(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v87 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89+v90)+72))
	v93 = F_get_namespace_name(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
	v191 = m.ExcPending
	if v191 != 0 {
		goto L6
	} else {
		goto L37
	}
L18:
	;
	v95 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v95
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = int64(4294967296)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v107
	v109 = F_cstring_to_text(m, v52)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v111 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+37)) = uint8(v111)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v109
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v72)+52))
	v121 = F_heap_modify_tuple(m, v87, v114, v22+int32(48), v22+int32(40), v22+int32(32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	F_CatalogTupleUpdate(m, v72, v121+int32(4), v121)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	F_systable_endscan(m, v85)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	F_relation_close(m, v72, int32(3))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v54)+40))
	if v132 == int32(0) {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v135 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v138 <= v135 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v149 = v135
	v150 = v135
	v153 = v135
	goto L26
L26:
	;
	v160 = int32(0)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162+v150<<(uint(int32(2))%32))))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v169 = F_get_required_extension(m, v166, v167, l4, l5, v160, l6)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L6
	} else {
		goto L28
	}
L27:
	;
	v214 = v180
	v218 = v182
	goto L11
L28:
	;
	v171 = F_SearchSysCache1(m, int32(28), v169)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	if v171 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+22)))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v173+v174)+72))
	F_ReleaseCatCache(m, v171)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	v179 = v160
	goto L32
L32:
	;
	v180 = F_lappend_oid(m, v149, v169)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L6
	} else {
		goto L34
	}
L33:
	;
	v179 = v176
	goto L32
L34:
	;
	v182 = F_lappend_oid(m, v153, v179)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v185 = v150 + int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v185 < v186 {
		v149 = v180
		v150 = v185
		v153 = v182
		goto L26
	} else {
		goto L36
	}
L36:
	;
	goto L27
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
	F_errmsg_internal(m, int32(_a_F_ApplyExtensionUpdates_0), v22)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_ApplyExtensionUpdates_1), int32(3605), int32(_a_F_ApplyExtensionUpdates_2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
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
	v230 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(3079)
	if v214 == v230 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyExtensionUpdates[0]))
	if v301 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v237 = int32(0)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v238 <= v237 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v250 = v237
	goto L44
L44:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v260+v250<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(3079)
	F_recordDependencyOn(m, v22+int32(20), v22+int32(8), int32(110))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L6
	} else {
		goto L46
	}
L45:
	;
	goto L41
L46:
	;
	v278 = v250 + int32(1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v278 < v279 {
		v250 = v278
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v303 = int32(0)
	F_RunObjectPostAlterHook(m, int32(3079), l0, v303, v303, v303)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L6
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	F_execute_extension_script(m, l0, v54, v31, v52, v218, v93)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L6
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	v311 = v44 + int32(1)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v311 < v312 {
		v31 = v52
		v44 = v311
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
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v201 int32
	_ = v201
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v342 int32
	_ = v342
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v465 int32
	_ = v465
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v677 int32
	_ = v677
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v765 int32
	_ = v765
	var v781 int32
	_ = v781
	var v793 int32
	_ = v793
	var v807 int32
	_ = v807
	var v822 int32
	_ = v822
	var v823 int64
	_ = v823
	var v839 int32
	_ = v839
	var v852 int32
	_ = v852
	var v868 int32
	_ = v868
	var v883 int32
	_ = v883
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v912 int32
	_ = v912
	var v924 int32
	_ = v924
	var v940 int32
	_ = v940
	var v955 int32
	_ = v955
	var v966 int32
	_ = v966
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v994 int32
	_ = v994
	var v1012 int32
	_ = v1012
	var v1024 int32
	_ = v1024
	var v1040 int32
	_ = v1040
	var v1055 int32
	_ = v1055
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1244 int32
	_ = v1244
	var v1245 int32
	_ = v1245
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1270 int32
	_ = v1270
	var v1271 int32
	_ = v1271
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1303 int32
	_ = v1303
	var v1316 int32
	_ = v1316
	var v1333 int32
	_ = v1333
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1438 int32
	_ = v1438
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1470 int32
	_ = v1470
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1508 int32
	_ = v1508
	var v1511 int32
	_ = v1511
	var v1519 int32
	_ = v1519
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1586 int32
	_ = v1586
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1680 int32
	_ = v1680
	var v1687 int32
	_ = v1687
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1729 int32
	_ = v1729
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1785 int32
	_ = v1785
	var v1793 int32
	_ = v1793
	var v1797 int32
	_ = v1797
	var v1809 int32
	_ = v1809
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1860 int32
	_ = v1860
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1878 int32
	_ = v1878
	var v1892 int32
	_ = v1892
	var v1904 int32
	_ = v1904
	var v1916 int32
	_ = v1916
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1945 int32
	_ = v1945
	var v1958 int32
	_ = v1958
	var v1972 int32
	_ = v1972
	var v1987 int32
	_ = v1987
	var v1998 int32
	_ = v1998
	var v2004 int32
	_ = v2004
	var v2018 int32
	_ = v2018
	var v2020 int32
	_ = v2020
	var v2021 int32
	_ = v2021
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2098 int32
	_ = v2098
	var v2111 int32
	_ = v2111
	var v2117 int32
	_ = v2117
	var v2134 int32
	_ = v2134
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2197 int32
	_ = v2197
	var v2198 int32
	_ = v2198
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2249 int32
	_ = v2249
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2269 int32
	_ = v2269
	var v2282 int32
	_ = v2282
	var v2300 int32
	_ = v2300
	var v2315 int32
	_ = v2315
	var v2323 int32
	_ = v2323
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2372 int32
	_ = v2372
	var v2373 int64
	_ = v2373
	var v2377 int32
	_ = v2377
	var v2379 int32
	_ = v2379
	var v2380 int32
	_ = v2380
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2391 int32
	_ = v2391
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	v7 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(560)
	m.G0 = v31
	v41 = v7
	v42 = v7
	v43 = v7
	v44 = v7
	v45 = v7
	v46 = v7
	v47 = v7
	v48 = v7
	v51 = int32(-1)
	v52 = v7
	v53 = v7
	v54 = v7
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
	if v51 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L3
L6:
	;
	v2372 = int32(m.ExcTag)
	v2373 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v2372 == int32(0) {
		goto L266
	} else {
		goto L267
	}
L7:
	;
	if v733 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L8:
	;
	v720 = v41
	v721 = v42
	v722 = v43
	v723 = v44
	v724 = v45
	v730 = v52
	v732 = v53
	v733 = v54
	goto L7
L9:
	;
	goto L10
L10:
	;
	v64 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+388)) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v31)+384)) = v64
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+33)))
	if v69 != int32(1) {
		v252 = v64
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	v264 = v53 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	v266 = F_get_extension_script_filename(m, l1, l2, l3)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L6
	} else {
		goto L37
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	v80 = int32(1)
	v81 = v53 & v80
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	v84 = v52 & v80
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	v86 = F_superuser(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v86 != 0 {
		v252 = v64
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	if v88 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[0]))
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	v117 = F_object_aclcheck(m, int32(1262), v102, v104, int64(512))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	if v117 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	F_errcode(m, int32(16797828))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+176)) = v151
	F_errmsg(m, int32(_a_F_execute_extension_script_0), v31+int32(176))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+192)) = v151
	F_errmsg(m, int32(_a_F_execute_extension_script_1), v31+int32(192))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L6
	} else {
		goto L31
	}
L25:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	if v170 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v183 = int32(_a_F_execute_extension_script_2)
	goto L28
L27:
	;
	v183 = int32(_a_F_execute_extension_script_3)
	goto L28
L28:
	;
	F_errhint(m, v183, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1227), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	goto L3
L31:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+34)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	if v218 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v231 = int32(_a_F_execute_extension_script_6)
	goto L34
L33:
	;
	v231 = int32(_a_F_execute_extension_script_7)
	goto L34
L34:
	;
	F_errhint(m, v231, int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v41
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v81)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v84)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1235), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	goto L3
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	v280 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), v328, int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L6
	} else {
		goto L48
	}
L41:
	;
	if v280 == int32(0) {
		goto L39
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v280 == int32(0) {
		goto L39
	} else {
		goto L46
	}
L44:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+148)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v31)+144)) = v286
	F_errmsg_internal(m, int32(_a_F_execute_extension_script_8), v31+int32(144))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	v328 = int32(1241)
	goto L40
L46:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+168)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v31)+164)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v31)+160)) = v307
	F_errmsg_internal(m, int32(_a_F_execute_extension_script_9), v31+int32(160))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L6
	} else {
		goto L47
	}
L47:
	;
	v328 = int32(1243)
	goto L40
L48:
	;
	goto L39
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	v359 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(388)))) = v359
	v362 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(384)))) = v362
	goto L52
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	v393 = int32(_a_F_execute_extension_script_10)
	v395 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[3]))
	v397 = v395 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[3])) = v397
	goto L54
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v31)+384))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[2])) = v375 | int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1])) = int32(10)
	goto L53
L53:
	;
	goto L51
L54:
	;
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[4]))
	if v400 <= int32(18) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	F_set_config_option(m, int32(_a_F_execute_extension_script_11), int32(_a_F_execute_extension_script_12), int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L6
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v422 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[5]))
	if v422 <= int32(18) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L57
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	F_set_config_option_ext(m, int32(_a_F_execute_extension_script_13), int32(_a_F_execute_extension_script_12), int32(5), int32(13), int32(10), int32(2), int32(0))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L6
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v445 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[6])))
	if v445 == int32(1) {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L61
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	F_set_config_option(m, int32(_a_F_execute_extension_script_14), int32(_a_F_execute_extension_script_15), int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L6
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	v477 = v31 + int32(368)
	F_initStringInfo(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L6
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	v490 = F_quote_identifier(m, l5)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v264)
	F_appendStringInfoString(m, v477, v490)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	v505 = l4 + int32(4)
	v507 = base.B2i32(l4 == int32(0))
	if l4 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	F_appendStringInfoString(m, v31+int32(368), int32(_a_F_execute_extension_script_16))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L6
	} else {
		goto L89
	}
L71:
	;
	v510 = int32(0)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v511 <= v510 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v534 = v510
	goto L73
L73:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v542+v534<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	v557 = F_get_namespace_name(m, v546)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L6
	} else {
		goto L76
	}
L74:
	;
	goto L70
L75:
	;
	v632 = v534 + int32(1)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v505)))
	if v632 < v633 {
		v534 = v632
		goto L73
	} else {
		goto L88
	}
L76:
	;
	if v557 == int32(0) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	v571 = int32(_a_F_execute_extension_script_17)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557))))
	v577 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[7])))
	if base.B2i32(v574 == int32(0))|base.B2i32(v574 != v577) != 0 {
		v595 = v574
		v596 = v577
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v595-v596 == int32(0) {
		goto L75
	} else {
		goto L85
	}
L79:
	;
	goto L78
L80:
	;
	v580 = v557
	v581 = v571
	goto L81
L81:
	;
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581)+1)))
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+1)))
	if v585 == int32(0) {
		v595 = v585
		v596 = v584
		goto L79
	} else {
		goto L83
	}
L82:
	;
	v595 = v585
	v596 = v584
	goto L79
L83:
	;
	v588 = int32(1)
	if v585 == v584 {
		v580 = v580 + v588
		v581 = v581 + v588
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	v610 = F_quote_identifier(m, v557)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+128)) = v610
	F_appendStringInfo(m, v31+int32(368), int32(_a_F_execute_extension_script_18), v31+int32(128))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v43
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v507)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v252)
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v31)+368))
	F_set_config_option(m, int32(_a_F_execute_extension_script_19), v689, int32(6), int32(13), int32(2), int32(1))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L6
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[8])) = l0
	v700 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[9])) = uint8(v700)
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[10]))
	v705 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11]))
	goto L91
L91:
	;
	v707 = v31 + int32(208)
	*(*int32)(unsafe.Add(mBase, uint32(v707)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v707))) = v31 + int32(204)
	goto L94
L92:
	;
	v720 = v266
	v721 = v705
	v722 = v703
	v723 = v505
	v724 = v397
	v730 = v252
	v732 = v507
	v733 = int32(0)
	goto L7
L94:
	;
	goto L92
L95:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[10])) = v31 + int32(208)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	v755 = int32(1)
	v756 = v732 & v755
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	v759 = v730 & v755
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v765 = F___fstatat(m, int32(-100), v720, v31+int32(392), int32(0))
	mBase = m.M
	goto L98
L96:
	;
	goto L97
L97:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11])) = v721
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[10])) = v722
	v2323 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[9])) = uint8(v2323)
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[8])) = v2323
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	v2336 = int32(1)
	v2337 = v732 & v2336
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v2337)
	v2340 = v730 & v2336
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v2340)
	F_pg_re_throw(m)
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L6
	} else {
		goto L265
	}
L98:
	;
	if v765 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L6
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v823 = *(*int64)(unsafe.Add(mBase, uint32(v31)+416))
	if int64(1073741823) <= v823 {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errcode_for_file_access(m)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L6
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v720
	F_errmsg(m, int32(_a_F_execute_extension_script_20), v31)
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L6
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(3949), int32(_a_F_execute_extension_script_21))
	mBase = m.M
	v822 = m.ExcPending
	if v822 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	goto L3
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L6
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v895 = F_AllocateFile(m, v720, int32(_a_F_execute_extension_script_22))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L6
	} else {
		goto L113
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errcode(m, int32(261))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L6
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v720
	F_errmsg(m, int32(_a_F_execute_extension_script_23), v31+int32(16))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L6
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(3954), int32(_a_F_execute_extension_script_21))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L6
	} else {
		goto L112
	}
L112:
	;
	goto L3
L113:
	;
	if v895 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v912 = m.ExcPending
	if v912 != 0 {
		goto L6
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v966 = base.I32_wrap_i64(v823)
	v969 = F_palloc(m, v966+int32(1))
	mBase = m.M
	v970 = m.ExcPending
	if v970 != 0 {
		goto L6
	} else {
		goto L121
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errcode_for_file_access(m)
	mBase = m.M
	v924 = m.ExcPending
	if v924 != 0 {
		goto L6
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v720
	F_errmsg(m, int32(_a_F_execute_extension_script_24), v31+int32(32))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(3961), int32(_a_F_execute_extension_script_21))
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	goto L3
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v982 = F_fread(m, v969, int32(1), v966, v895)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v895)))
	goto L123
L123:
	;
	if int32(base.Ui32(v994)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L6
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1066 = F_FreeFile(m, v895)
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L6
	} else {
		goto L131
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errcode_for_file_access(m)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+112)) = v720
	F_errmsg(m, int32(_a_F_execute_extension_script_25), v31+int32(112))
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L6
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(3970), int32(_a_F_execute_extension_script_21))
	mBase = m.M
	v1055 = m.ExcPending
	if v1055 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	goto L3
L131:
	;
	v1069 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v969+v982))) = uint8(v1069)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v1071 < v1069 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1085 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[12]))
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1085)+4))
	goto L135
L133:
	;
	v1087 = v48
	v1088 = v1071
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1100 = F_pg_verify_mbstr(m, v1088, v969, v982, int32(0))
	mBase = m.M
	v1101 = m.ExcPending
	if v1101 != 0 {
		goto L6
	} else {
		goto L136
	}
L135:
	;
	v1087 = v1086
	v1088 = v1086
	goto L134
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1112 = F_pg_any_to_server(m, v969, v982, v1088)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1124 = F_cstring_to_text(m, v1112)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L6
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1137 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_26))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1150 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_27))
	mBase = m.M
	v1151 = m.ExcPending
	if v1151 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1163 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_28))
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L6
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1177 = F_DirectFunctionCall4Coll(m, int32(554), int32(950), v1124, v1137, v1150, v1163)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L6
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1190 = F_strstr(m, v1112, int32(_a_F_execute_extension_script_29))
	mBase = m.M
	if v1190 == int32(0) {
		v1349 = v46
		v1351 = v1177
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v1354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	if v1354 != 0 {
		v1487 = v1351
		goto L163
	} else {
		goto L164
	}
L144:
	;
	if v759 != 0 {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1219 = F_GetUserNameFromId(m, v1207, int32(0))
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L6
	} else {
		goto L149
	}
L146:
	;
	v1193 = *(*int32)(unsafe.Add(mBase, uint32(v31)+388))
	v1206 = v46
	v1207 = v1193
	goto L145
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1205 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1]))
	v1206 = v1205
	v1207 = v1205
	goto L145
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1231 = F_quote_identifier(m, v1219)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1244 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_29))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1256 = F_cstring_to_text(m, v1231)
	mBase = m.M
	v1257 = m.ExcPending
	if v1257 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1270 = F_DirectFunctionCall3Coll(m, int32(555), int32(950), v1177, v1244, v1256)
	mBase = m.M
	v1271 = m.ExcPending
	if v1271 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1283 = F_strcspn(m, v1219, int32(_a_F_execute_extension_script_30))
	mBase = m.M
	v1284 = v1283 + v1219
	v1286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1284))))
	if v1286 != 0 {
		goto L155
	} else {
		goto L156
	}
L154:
	;
	if v1287 == int32(0) {
		v1349 = v1206
		v1351 = v1270
		goto L143
	} else {
		goto L158
	}
L155:
	;
	v1287 = v1284
	goto L157
L156:
	;
	v1287 = int32(0)
	goto L157
L157:
	;
	goto L154
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1316 = m.ExcPending
	if v1316 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+96)) = int32(_a_F_execute_extension_script_30)
	F_errmsg(m, int32(_a_F_execute_extension_script_31), v31+int32(96))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L6
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1206
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1369), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v1348 = m.ExcPending
	if v1348 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	goto L3
L163:
	;
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v1508 = v1487
	v1511 = int32(0)
	goto L179
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1365 = F_quote_identifier(m, l5)
	mBase = m.M
	v1366 = m.ExcPending
	if v1366 != 0 {
		goto L6
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1378 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_32))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L6
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1390 = F_cstring_to_text(m, v1365)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L6
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1404 = F_DirectFunctionCall3Coll(m, int32(555), int32(950), v1351, v1378, v1390)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L6
	} else {
		goto L168
	}
L168:
	;
	if v1351 == v1404 {
		v1487 = v1351
		goto L163
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1418 = F_strcspn(m, l5, int32(_a_F_execute_extension_script_30))
	mBase = m.M
	v1419 = v1418 + l5
	v1421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1419))))
	if v1421 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	if v1422 == int32(0) {
		v1487 = v1404
		goto L163
	} else {
		goto L174
	}
L171:
	;
	v1422 = v1419
	goto L173
L172:
	;
	v1422 = int32(0)
	goto L173
L173:
	;
	goto L170
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L6
	} else {
		goto L175
	}
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L6
	} else {
		goto L176
	}
L176:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+84)) = int32(_a_F_execute_extension_script_30)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+80)) = v1452
	F_errmsg(m, int32(_a_F_execute_extension_script_33), v31+int32(80))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L6
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1393), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v1485 = m.ExcPending
	if v1485 != 0 {
		goto L6
	} else {
		goto L178
	}
L178:
	;
	goto L3
L179:
	;
	v1519 = int32(0)
	if v1489 == v1519 {
		v1529 = v1519
		goto L181
	} else {
		goto L182
	}
L181:
	;
	if v756 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L182:
	;
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+4))
	if v1523 <= v1511 {
		v1529 = int32(0)
		goto L181
	} else {
		goto L183
	}
L183:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1489)+12))
	v1529 = v1525 + v1511<<(uint(int32(2))%32)
	goto L181
L184:
	;
	v2157 = *(*int32)(unsafe.Add(mBase, uint32(v1539+v1511<<(uint(int32(2))%32))))
	v2158 = *(*int32)(unsafe.Add(mBase, uint32(v1529)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v2169 = F_get_namespace_name(m, v2157)
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L6
	} else {
		goto L248
	}
L185:
	;
	v1532 = int32(0)
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v723)))
	if base.B2i32(v1529 == v1532)|base.B2i32(v1534 <= v1511) == v1532 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v1541 = v1487
	goto L187
L187:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v1543 != 0 {
		goto L192
	} else {
		goto L193
	}
L188:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v1539 != 0 {
		goto L184
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1541 = v1508
	goto L187
L191:
	;
	goto L190
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1555 = F_cstring_to_text(m, int32(_a_F_execute_extension_script_34))
	mBase = m.M
	v1556 = m.ExcPending
	if v1556 != 0 {
		goto L6
	} else {
		goto L195
	}
L193:
	;
	v1584 = v47
	v1586 = v1541
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1598 = F_pg_detoast_datum_packed(m, v1586)
	mBase = m.M
	v1599 = m.ExcPending
	if v1599 != 0 {
		goto L6
	} else {
		goto L198
	}
L195:
	;
	v1557 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1568 = F_cstring_to_text(m, v1557)
	mBase = m.M
	v1569 = m.ExcPending
	if v1569 != 0 {
		goto L6
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1582 = F_DirectFunctionCall3Coll(m, int32(555), int32(950), v1541, v1555, v1568)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L6
	} else {
		goto L197
	}
L197:
	;
	v1584 = v1582
	v1586 = v1582
	goto L194
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1610 = F_text_to_cstring(m, v1598)
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L6
	} else {
		goto L199
	}
L199:
	;
	v1612 = int32(_a_F_execute_extension_script_35)
	v1613 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11])) = v31 + int32(492)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+512)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+508)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(v31)+504)) = v1610
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+496)) = int32(556)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+492)) = v1613
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+500)) = v31 + int32(504)
	v1638 = F_pg_parse_query(m, v1610)
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L6
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1651 = F_CreateDestReceiver(m, int32(0))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L6
	} else {
		goto L201
	}
L201:
	;
	if v1638 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(v31)+492))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11])) = v2098
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L6
	} else {
		goto L242
	}
L203:
	;
	v1655 = int32(0)
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	if v1656 <= v1655 {
		goto L202
	} else {
		goto L204
	}
L204:
	;
	v1680 = v1655
	goto L205
L205:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+12))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1687+v1680<<(uint(int32(2))%32))))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1691)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+512)) = v1692
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v1691)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+516)) = v1694
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1707 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[13]))
	v1712 = F_AllocSetContextCreateInternal(m, v1707, int32(_a_F_execute_extension_script_36), int32(0), int32(_a_F_execute_extension_script_37), int32(_a_F_execute_extension_script_38))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L6
	} else {
		goto L207
	}
L206:
	;
	goto L202
L207:
	;
	v1714 = int32(_a_F_execute_extension_script_39)
	v1715 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[13])) = v1712
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L6
	} else {
		goto L208
	}
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1740 = int32(0)
	v1743 = F_pg_analyze_and_rewrite_fixedparams(m, v1691, v1610, v1740, v1740, v1740)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L6
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1757 = F_pg_plan_queries(m, v1743, v1610, int32(2048), int32(0))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L6
	} else {
		goto L211
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[13])) = v1715
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_MemoryContextDelete(m, v1712)
	mBase = m.M
	v2064 = m.ExcPending
	if v2064 != 0 {
		goto L6
	} else {
		goto L240
	}
L211:
	;
	if v1757 == int32(0) {
		goto L210
	} else {
		goto L212
	}
L212:
	;
	v1761 = int32(0)
	v1762 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+4))
	if v1762 <= v1761 {
		goto L210
	} else {
		goto L213
	}
L213:
	;
	v1785 = v1761
	goto L214
L214:
	;
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+12))
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(v1793+v1785<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_CommandCounterIncrement(m)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L6
	} else {
		goto L216
	}
L215:
	;
	goto L210
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1820 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L6
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_PushActiveSnapshot(m, v1820)
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L6
	} else {
		goto L218
	}
L218:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1797)+88))
	if v1834 == int32(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_PopActiveSnapshot(m)
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L6
	} else {
		goto L238
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1848 = *(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[14]))
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(v1848)))
	goto L223
L221:
	;
	goto L222
L222:
	;
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1834)))
	if v1929 == int32(225) {
		goto L230
	} else {
		goto L231
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1860 = int32(0)
	v1864 = F_CreateQueryDesc(m, v1797, v1610, v1849, v1860, v1651, v1860, v1860, v1860)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L6
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_ExecutorStart(m, v1864, int32(0))
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L6
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_ExecutorRun(m, v1864, int32(1), int64(0))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L6
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_ExecutorFinish(m, v1864)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L6
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_ExecutorEnd(m, v1864)
	mBase = m.M
	v1916 = m.ExcPending
	if v1916 != 0 {
		goto L6
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_FreeQueryDesc(m, v1864)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L6
	} else {
		goto L229
	}
L229:
	;
	goto L219
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L6
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v1998 = int32(0)
	F_ProcessUtility(m, v1797, v1610, v1998, int32(1), v1998, v1998, v1651, v1998)
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L6
	} else {
		goto L237
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errcode(m, int32(1088))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L6
	} else {
		goto L234
	}
L234:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errmsg(m, int32(_a_F_execute_extension_script_40), int32(0))
	mBase = m.M
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L6
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1141), int32(_a_F_execute_extension_script_41))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
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
	v2020 = v1785 + int32(1)
	v2021 = *(*int32)(unsafe.Add(mBase, uint32(v1757)+4))
	if v2020 < v2021 {
		v1785 = v2020
		goto L214
	} else {
		goto L239
	}
L239:
	;
	goto L215
L240:
	;
	v2066 = v1680 + int32(1)
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v1638)+4))
	if v2066 < v2067 {
		v1680 = v2066
		goto L205
	} else {
		goto L241
	}
L241:
	;
	goto L206
L242:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[10])) = v722
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[11])) = v721
	v2117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_execute_extension_script[9])) = uint8(v2117)
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[8])) = v2117
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_AtEOXact_GUC(m, int32(1), v724)
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L6
	} else {
		goto L243
	}
L243:
	;
	if v759 != 0 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v1584
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v31)+388))
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v31)+384))
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[2])) = v2146
	*(*int32)(unsafe.Add(mBase, _c_F_execute_extension_script[1])) = v2145
	goto L247
L245:
	;
	goto L246
L246:
	;
	m.G0 = v31 + int32(560)
	return
L247:
	;
	goto L246
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v2181 = F_quote_identifier(m, v2169)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L6
	} else {
		goto L249
	}
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+64)) = v2158
	v2197 = F_psprintf(m, int32(_a_F_execute_extension_script_42), v31-int32(-64))
	mBase = m.M
	v2198 = m.ExcPending
	if v2198 != 0 {
		goto L6
	} else {
		goto L250
	}
L250:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v2209 = F_cstring_to_text(m, v2197)
	mBase = m.M
	v2210 = m.ExcPending
	if v2210 != 0 {
		goto L6
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v2221 = F_cstring_to_text(m, v2181)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L6
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v2235 = F_DirectFunctionCall3Coll(m, int32(555), int32(950), v1508, v2209, v2221)
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L6
	} else {
		goto L254
	}
L253:
	;
	v1508 = v2235
	v1511 = v1511 + int32(1)
	goto L179
L254:
	;
	if v1508 == v2235 {
		goto L253
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	v2249 = F_strcspn(m, v2169, int32(_a_F_execute_extension_script_30))
	mBase = m.M
	v2250 = v2249 + v2169
	v2252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2250))))
	if v2252 != 0 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v2253 == int32(0) {
		goto L253
	} else {
		goto L260
	}
L257:
	;
	v2253 = v2250
	goto L259
L258:
	;
	v2253 = int32(0)
	goto L259
L259:
	;
	goto L256
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L6
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errcode(m, int32(33685634))
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L6
	} else {
		goto L262
	}
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = int32(_a_F_execute_extension_script_30)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = v2158
	F_errmsg(m, int32(_a_F_execute_extension_script_33), v31+int32(48))
	mBase = m.M
	v2300 = m.ExcPending
	if v2300 != 0 {
		goto L6
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+524)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v31)+520)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v31)+528)) = v1087
	*(*int32)(unsafe.Add(mBase, uint32(v31)+532)) = v721
	*(*int32)(unsafe.Add(mBase, uint32(v31)+536)) = v722
	*(*int32)(unsafe.Add(mBase, uint32(v31)+544)) = v723
	*(*int32)(unsafe.Add(mBase, uint32(v31)+548)) = v724
	*(*int32)(unsafe.Add(mBase, uint32(v31)+552)) = v720
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)) = uint8(v756)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)) = uint8(v759)
	F_errfinish(m, int32(_a_F_execute_extension_script_4), int32(1420), int32(_a_F_execute_extension_script_5))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
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
	v2377 = int32(v2373)
	m.G0 = v31
	v2379 = *(*int32)(unsafe.Add(mBase, uint32(v2377)+4))
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v2377)))
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(v2380)))
	if v31+int32(204) == v2383 {
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
	if v2387 != 0 {
		goto L272
	} else {
		goto L273
	}
L269:
	;
	v2385 = *(*int32)(unsafe.Add(mBase, uint32(v2380)+4))
	v2387 = v2385
	goto L271
L270:
	;
	v2387 = int32(0)
	goto L271
L271:
	;
	goto L268
L272:
	;
	v2388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+559)))
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(v31)+552))
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v31)+548))
	v2391 = *(*int32)(unsafe.Add(mBase, uint32(v31)+544))
	v2392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+543)))
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v31)+536))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v31)+532))
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(v31)+528))
	v2396 = *(*int32)(unsafe.Add(mBase, uint32(v31)+524))
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v31)+520))
	v41 = v2389
	v42 = v2394
	v43 = v2393
	v44 = v2391
	v45 = v2390
	v46 = v2396
	v47 = v2397
	v48 = v2395
	v51 = v2387
	v52 = v2388
	v53 = v2392
	v54 = v2379
	goto L1
L273:
	;
	goto L274
L274:
	;
	F___wasm_longjmp(m, v2380, v2379)
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
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
