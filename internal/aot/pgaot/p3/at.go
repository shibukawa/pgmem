package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ATExecDropColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v205 int32
	_ = v205
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int64
	_ = v240
	var v242 int32
	_ = v242
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	v19 = m.G0
	v21 = v19 - int32(128)
	m.G0 = v21
	if l5 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v35 = F_SearchSysCacheAttName(m, v34, l2)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L19
	}
L2:
	;
	F_ATSimplePermissions(m, int32(13), l1, int32(289))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L5
	} else {
		goto L8
	}
L5:
	;
	return
L6:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v33 = l8
	goto L1
L8:
	;
	v31 = F_new_object_addresses(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v33 = v31
	goto L1
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L112
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L109
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L105
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L100
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L5
	} else {
		goto L96
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L5
	} else {
		goto L92
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L5
	} else {
		goto L88
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L84
	}
L18:
	;
	m.G0 = v21 + int32(128)
	return
L19:
	;
	if v35 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if l6 == int32(0) {
		goto L17
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+22)))
	v67 = v65 + v66
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+74)))
	if v68 <= int32(0) {
		goto L16
	} else {
		goto L30
	}
L23:
	;
	v43 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	if v43 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v45 + int32(4)
	F_errmsg(m, int32(318711), v21)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v60 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v60
	v63 = *(*int32)(unsafe.Add(mBase, _consts[446]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v63
	goto L18
L28:
	;
	F_errfinish(m, int32(475396), int32(9328), int32(263587))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if l5 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v73 = int32(*(*int16)(unsafe.Add(mBase, uint32(v67)+94)))
	if int32(0) < v73 {
		goto L15
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v78 = F_bms_make_singleton(m, v68+int32(7))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	v82 = F_has_partition_attrs(m, l1, v78, v21+int32(115))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	if v82 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	F_ReleaseCatCache(m, v35)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v87 = F_find_inheritance_children(m, v86, l7)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	if v87 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if l4 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+116)) = int32(1259)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v226
	F_add_exact_object_address(m, v21+int32(116), v33)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L78
	}
L43:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+119)))
	if v92 == int32(112) {
		goto L13
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v97 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if int32(0) < v99 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v119 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	F_sequence_close(m, v97, int32(3))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L77
	}
L51:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v119<<(uint(int32(2))%32))))
	v126 = F_table_open(m, v124, int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v126)+48))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+118)))
	if v129 == int32(116) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+24)))
	if v132 == int32(0) {
		goto L12
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	F_CheckTableNotInUse(m, v126, int32(520745))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L5
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v138 = F_SearchSysCacheCopyAttName(m, v124, l2)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	if v138 == int32(0) {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+22)))
	v144 = v142 + v143
	v145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v144)+94)))
	if v145 <= int32(0) {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	if l4 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	F_pfree(m, v138)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L5
	} else {
		goto L74
	}
L63:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L73
	}
L64:
	;
	if v145 != int32(1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v165 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+92)) = uint8(v165)
	v168 = v145 - v165
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+94)) = uint16(v168)
	F_CatalogTupleUpdate(m, v97, v138+int32(4), v138)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L72
	}
L67:
	;
	v159 = v145 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+94)) = uint16(v159)
	F_CatalogTupleUpdate(m, v97, v138+int32(4), v138)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L5
	} else {
		goto L71
	}
L68:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+92)))
	if v150 != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v153 = int32(1)
	F_ATExecDropColumn(m, v21+int32(116), v126, l2, l3, v153, v153, int32(0), l7, v33)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	goto L62
L71:
	;
	goto L63
L72:
	;
	goto L63
L73:
	;
	goto L62
L74:
	;
	F_sequence_close(m, v126, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v182 = v119 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v182 < v183 {
		v119 = v182
		goto L51
	} else {
		goto L76
	}
L76:
	;
	goto L52
L77:
	;
	goto L42
L78:
	;
	if l5 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_performMultipleDeletions(m, v33, l3, int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v21)+116))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v21)+124))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v242
	goto L18
L82:
	;
	F_free_object_addresses(m, v33)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v272 + int32(4)
	F_errmsg(m, int32(68939), v21+int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L5
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(475396), int32(9322), int32(263587))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l2
	F_errmsg(m, int32(673673), v21+int32(32))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L5
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(475396), int32(9341), int32(263587))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = l2
	F_errmsg(m, int32(674189), v21+int32(96))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	F_errfinish(m, int32(475396), int32(9351), int32(263587))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L96:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L5
	} else {
		goto L97
	}
L97:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v330 + int32(4)
	F_errmsg(m, int32(669625), v21+int32(48))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L5
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(475396), int32(9364), int32(263587))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	F_errmsg(m, int32(70194), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L5
	} else {
		goto L102
	}
L102:
	;
	F_errhint(m, int32(607648), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(475396), int32(9389), int32(263587))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(135867), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(475396), int32(4460), int32(393756))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = l2
	F_errmsg_internal(m, int32(44623), v21-int32(-64))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(475396), int32(9405), int32(263587))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v124
	F_errmsg_internal(m, int32(675666), v21+int32(80))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L5
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(475396), int32(9410), int32(263587))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecDropIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+119)))
	if base.B2i32(l5 == v8)&base.B2i32(v21 == int32(112)) == v8 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L12
	} else {
		goto L73
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L12
	} else {
		goto L69
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L12
	} else {
		goto L65
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L12
	} else {
		goto L61
	}
L5:
	;
	if l6 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L12
	} else {
		goto L56
	}
L8:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+131)))
	if v29&int32(1) != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v34 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	return
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v37 = F_SearchSysCacheCopyAttName(m, v36, l2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	if v37 == int32(0) {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+22)))
	v43 = v41 + v42
	v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+74)))
	if v44 <= int32(0) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+89)))
	if v47 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	m.G0 = v16 + int32(80)
	return
L18:
	;
	if l3 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v43)+89)) = uint8(v83)
	F_CatalogTupleUpdate(m, v34, v37+int32(4), v37)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L30
	}
L21:
	;
	v54 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	if v54 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v56 + int32(4)
	F_errmsg(m, int32(319743), v16+int32(32))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_pfree(m, v37)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L28
	}
L26:
	;
	F_errfinish(m, int32(475396), int32(8540), int32(9812))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	F_sequence_close(m, v34, int32(3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[446]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78
	v81 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v81
	goto L17
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _consts[433]))
	if v90 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v43)+74)))
	v94 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1259), v92, v93, v94, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L12
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_pfree(m, v37)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L12
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	F_sequence_close(m, v34, int32(3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	if v21 != int32(112) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if l6 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L38:
	;
	if l5 == int32(0) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v109 = F_find_inheritance_children(m, v108, l4)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L40
	}
L40:
	;
	if v109 == int32(0) {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v113 <= int32(0) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v125 = int32(0)
	goto L43
L43:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v125<<(uint(int32(2))%32))))
	v138 = F_table_open(m, v136, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L12
	} else {
		goto L45
	}
L44:
	;
	goto L37
L45:
	;
	v141 = int32(1)
	F_ATExecDropIdentity(m, v16+int32(68), v138, l2, int32(0), l4, v141, v141)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	F_sequence_close(m, v138, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L12
	} else {
		goto L47
	}
L47:
	;
	v149 = v125 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v149 < v150 {
		v125 = v149
		goto L43
	} else {
		goto L48
	}
L48:
	;
	goto L44
L49:
	;
	v169 = F_getIdentitySequence(m, l1, v44, int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L12
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L17
L52:
	;
	v173 = F_deleteDependencyRecordsForClass(m, int32(1259), v169, int32(1259), int32(105))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	v177 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v177
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = int32(1259)
	F_performDeletion(m, v16+int32(68), v177, int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(379072), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errhint(m, int32(607648), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(475396), int32(8505), int32(9812))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	F_errmsg(m, int32(238763), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(475396), int32(8510), int32(9812))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L12
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v252 + int32(4)
	F_errmsg(m, int32(68939), v16)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(475396), int32(8518), int32(9812))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L12
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	F_errmsg(m, int32(673641), v16+int32(16))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(475396), int32(8527), int32(9812))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L12
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
	F_errcode(m, int32(325))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v290 + int32(4)
	F_errmsg(m, int32(261954), v16+int32(48))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(475396), int32(8535), int32(9812))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L12
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ATExecValidateConstraint(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
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
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
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
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int64
	_ = v372
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v642 int32
	_ = v642
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	v17 = m.G0
	v19 = v17 - int32(208)
	m.G0 = v19
	v23 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_ScanKeyInit(m, v19+int32(48), int32(9), int32(3), int32(184), v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v19+int32(96), int32(10), int32(3), int32(184), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_ScanKeyInit(m, v19+int32(144), int32(2), int32(3), int32(62), l3)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v54 = F_systable_beginscan(m, v23, int32(2665), int32(1), int32(0), int32(3), v19+int32(48))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	F_systable_endscan(m, v54)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L1
	} else {
		goto L139
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v642
	goto L6
L8:
	;
	v505 = int32(0)
	F_set_attnotnull(m, v505, l2, v291, v505)
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L1
	} else {
		goto L119
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L115
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L111
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L107
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L103
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L98
	}
L14:
	;
	v56 = F_systable_getnext(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+22)))
	v60 = v58 + v59
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+72)))
	v63 = v61 - int32(99)
	if base.Ui32(int32(11)) < base.Ui32(v63) {
		goto L13
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L94
	}
L19:
	;
	if int32(1)<<(uint(v63)%32)&int32(2057) == int32(0) {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+75)))
	if v72 == int32(0) {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+76)))
	if v75 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	switch v61 - int32(99) {
	case 0:
		goto L26
	default:
		goto L7
	case 3:
		goto L27
	case 11:
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v372 = *(*int64)(unsafe.Add(mBase, _consts[447]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v372
	v375 = *(*int32)(unsafe.Add(mBase, _consts[446]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v375
	goto L6
L25:
	;
	v291 = F_extractNotNullColumn(m, v56)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L70
	}
L26:
	;
	if l5 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v60)+96))
	F_QueueFKConstraintValidation(m, l1, v23, l2, v80, v56, l6)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L7
L29:
	;
	v153 = F_palloc0(m, int32(32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L45
	}
L30:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+106)))
	if v83 != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v87 = F_find_all_inheritors(m, v85, l6, v84)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	if v87 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v91 <= int32(0) {
		goto L29
	} else {
		goto L34
	}
L34:
	;
	v101 = v84
	goto L35
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v101<<(uint(int32(2))%32))))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v114 != v115 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L29
L37:
	;
	if l4 == int32(0) {
		goto L11
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v133 = v101 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v133 < v134 {
		v101 = v133
		goto L35
	} else {
		goto L44
	}
L40:
	;
	v122 = F_table_open(m, v114, int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_ATExecValidateConstraint(m, v19+int32(196), l1, v122, l3, int32(0), int32(1), l6)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_sequence_close(m, v122, int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L39
L44:
	;
	goto L36
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v153)+4)) = int64(5)
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = l3
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = v160
	v164 = F_SysCacheGetAttrNotNull(m, int32(19), v56, int32(28))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v166 = F_text_to_cstring(m, v164)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v168 = F_stringToNode(m, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v171 = F_expand_generated_columns_in_expr(m, v168, l2, int32(1))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = v171
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v175 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v254)+64))
	v264 = F_lappend(m, v263, v153)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L61
	}
L51:
	;
	v225 = F_palloc0(m, int32(140))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L58
	}
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
	if v178 <= int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v186 = int32(0)
	goto L54
L54:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v181+v186<<(uint(int32(2))%32))))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	if v203 == v174 {
		v254 = v202
		goto L50
	} else {
		goto L56
	}
L55:
	;
	goto L51
L56:
	;
	v206 = v186 + int32(1)
	if v178 != v206 {
		v186 = v206
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v225))) = v174
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+4)) = uint8(v231)
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v234 = F_CreateTupleDescCopyConstr(m, v233)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = v234
	*(*int64)(unsafe.Add(mBase, uint32(v225)+88)) = int64(0)
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v225)+84)) = uint8(v239)
	v241 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v225)+96)) = uint16(v241)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v244 = F_lappend(m, v243, v225)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v244
	v254 = v225
	goto L50
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v254)+64)) = v264
	F_CacheInvalidateRelcache(m, l2)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v269 = F_heap_copytuple(m, v56)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v269)+16))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+22)))
	v274 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v271+v272)+76)) = uint8(v274)
	F_CatalogTupleUpdate(m, v23, v269+int32(4), v269)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v281 = *(*int32)(unsafe.Add(mBase, _consts[433]))
	if v281 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v284 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v283, v284, v284, v284)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_pfree(m, v269)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	goto L67
L69:
	;
	goto L7
L70:
	;
	if l5 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v302 = int32(0)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v305 = F_find_all_inheritors(m, v303, l6, v302)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L77
	}
L72:
	;
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+106)))
	if v295 != int32(1) {
		goto L71
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v300 = F_get_attname(m, v298, v291, int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	goto L8
L77:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v309 = F_get_attname(m, v307, v291, int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v305 == int32(0) {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v313 <= int32(0) {
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v319 = v302
	goto L81
L81:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v305)+12))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v332+v319<<(uint(int32(2))%32))))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v336 == v337 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L8
L83:
	;
	v368 = v319 + int32(1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v305)+4))
	if v368 < v369 {
		v319 = v368
		goto L81
	} else {
		goto L93
	}
L84:
	;
	if l4 == int32(0) {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	v341 = F_findNotNullConstraint(m, v336, v309)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	if v341 == int32(0) {
		goto L9
	} else {
		goto L87
	}
L87:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341)+16))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+22)))
	v347 = v345 + v346
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+76)))
	if v348 != 0 {
		goto L83
	} else {
		goto L88
	}
L88:
	;
	v352 = F_table_open(m, v336, int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v356 = F_pstrdup(m, v347+int32(4))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_ATExecValidateConstraint(m, v19+int32(196), l1, v352, v356, int32(0), int32(1), l6)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	F_sequence_close(m, v352, int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	goto L83
L93:
	;
	goto L82
L94:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v384 + int32(4)
	F_errmsg(m, int32(68873), v19)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(475396), int32(12943), int32(87110))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v404 + int32(4)
	F_errmsg(m, int32(670532), v19+int32(16))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errdetail(m, int32(543882), int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(475396), int32(12953), int32(87110))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(86642), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(475396), int32(12958), int32(87110))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L107:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg(m, int32(230289), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(475396), int32(13167), int32(254413))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(230289), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(475396), int32(13267), int32(254444))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	v475 = F_get_rel_name(m, v336)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v309
	F_errmsg_internal(m, int32(670945), v19+int32(32))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(475396), int32(13276), int32(254444))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v510 == int32(0) {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v598 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v589)+76)) = uint8(v598)
	F_CacheInvalidateRelcache(m, l2)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L131
	}
L121:
	;
	v560 = F_palloc0(m, int32(140))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L128
	}
L122:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	if v513 <= int32(0) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v510)+12))
	v521 = int32(0)
	goto L124
L124:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v516+v521<<(uint(int32(2))%32))))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	if v538 == v509 {
		v589 = v537
		goto L120
	} else {
		goto L126
	}
L125:
	;
	goto L121
L126:
	;
	v541 = v521 + int32(1)
	if v513 != v541 {
		v521 = v541
		goto L124
	} else {
		goto L127
	}
L127:
	;
	goto L125
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v560)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v560))) = v509
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v565)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v560)+4)) = uint8(v566)
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v569 = F_CreateTupleDescCopyConstr(m, v568)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v560)+8)) = v569
	*(*int64)(unsafe.Add(mBase, uint32(v560)+88)) = int64(0)
	v574 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v560)+84)) = uint8(v574)
	v576 = int32(28672)
	*(*uint16)(unsafe.Add(mBase, uint32(v560)+96)) = uint16(v576)
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v579 = F_lappend(m, v578, v560)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v579
	v589 = v560
	goto L120
L131:
	;
	v602 = F_heap_copytuple(m, v56)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v602)+16))
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v604)+22)))
	v607 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v604+v605)+76)) = uint8(v607)
	F_CatalogTupleUpdate(m, v23, v602+int32(4), v602)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v614 = *(*int32)(unsafe.Add(mBase, _consts[433]))
	if v614 != 0 {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v617 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v616, v617, v617, v617)
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	F_pfree(m, v602)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	goto L7
L139:
	;
	F_sequence_close(m, v23, int32(3))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	m.G0 = v19 + int32(208)
	return
}
func F_ATPrepAddColumn(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if l3 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
		if v13 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errmsg(m, int32(378542), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						F_errfinish(m, int32(475396), int32(7200), int32(263644))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
			if v14 == int32(99) {
				F_ATTypedTableRecursion(m, l0, l1, l5, l6, l7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					if l4 != 0 {
					} else {
						if l2 == int32(0) {
						} else {
							v21 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l5)+29)) = uint8(v21)
						}
					}
					return
				}
			} else {
				if l4 != 0 {
				} else {
					if l2 == int32(0) {
					} else {
						v21 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l5)+29)) = uint8(v21)
					}
				}
				return
			}
		}
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+119)))
		if v14 == int32(99) {
			F_ATTypedTableRecursion(m, l0, l1, l5, l6, l7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				if l4 != 0 {
				} else {
					if l2 == int32(0) {
					} else {
						v21 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l5)+29)) = uint8(v21)
					}
				}
				return
			}
		} else {
			if l4 != 0 {
			} else {
				if l2 == int32(0) {
				} else {
					v21 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l5)+29)) = uint8(v21)
				}
			}
			return
		}
	}
}
func F_ATSimplePermissions(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+119)))
	switch v12 - int32(73) {
	case 0:
		v24 = int32(64)
	default:
		v24 = int32(0)
	case 10:
		v24 = int32(128)
	case 26:
		v24 = int32(16)
	case 29:
		v24 = int32(32)
	case 32:
		v24 = int32(8)
	case 36:
		v24 = int32(4)
	case 39:
		v24 = int32(256)
	case 41:
		v24 = int32(1)
	case 45:
		v24 = int32(2)
	}
	if l2&v24 == int32(0) {
		if base.Ui32(l0) <= base.Ui32(int32(64)) {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[496])))
			v36 = v34
		} else {
			v36 = int32(0)
		}
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			if v36 != 0 {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return
				} else {
					v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v103 + int32(4)
					F_errmsg(m, int32(668110), v8+int32(16))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return
					} else {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v113)+119)))
						F_errdetail_relkind_not_supported(m, v114)
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							F_errfinish(m, int32(475396), int32(6788), int32(135624))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v41 + int32(4)
				F_errmsg_internal(m, int32(668062), v8)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errfinish(m, int32(475396), int32(6792), int32(135624))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
		v56 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v57 = F_object_ownercheck(m, int32(1259), v54, v56)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			if v57 == int32(0) {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				v63 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62)+119)))
				switch v63 - int32(73) {
				case 0, 32:
					v75 = int32(20)
				default:
					v73 = int32(41)
					v75 = v73
				case 10:
					v75 = int32(37)
				case 29:
					v73 = int32(18)
					v75 = v73
				case 36:
					v75 = int32(23)
				case 45:
					v75 = int32(51)
				}
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				F_aclcheck_error(m, int32(2), v75, v76+int32(4))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					v82 = int32(*(*uint8)(unsafe.Add(mBase, _consts[486])))
					if v82 == int32(0) {
						v86 = int32(1)
						v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
						if base.Ui32(v87) < base.Ui32(int32(12000)) {
							v96 = v86
						} else {
							v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
							v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+68))
							if v91 == int32(99) {
								v96 = v86
							} else {
								v94 = F_isTempToastNamespace(m, v91)
								mBase = m.M
								v96 = v94
							}
						}
						if v96 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return
								} else {
									v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v129 + int32(4)
									F_errmsg(m, int32(313736), v8+int32(32))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return
									} else {
										F_errfinish(m, int32(475396), int32(6804), int32(135624))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							m.G0 = v8 + int32(48)
							return
						}
					} else {
						m.G0 = v8 + int32(48)
						return
					}
				}
			} else {
				v82 = int32(*(*uint8)(unsafe.Add(mBase, _consts[486])))
				if v82 == int32(0) {
					v86 = int32(1)
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
					if base.Ui32(v87) < base.Ui32(int32(12000)) {
						v96 = v86
					} else {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+68))
						if v91 == int32(99) {
							v96 = v86
						} else {
							v94 = F_isTempToastNamespace(m, v91)
							mBase = m.M
							v96 = v94
						}
					}
					if v96 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return
							} else {
								v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v129 + int32(4)
								F_errmsg(m, int32(313736), v8+int32(32))
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return
								} else {
									F_errfinish(m, int32(475396), int32(6804), int32(135624))
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						m.G0 = v8 + int32(48)
						return
					}
				} else {
					m.G0 = v8 + int32(48)
					return
				}
			}
		}
	}
}
func F_ATSimpleRecursion(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	if l3 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L23
	}
L2:
	;
	return
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+126)))
	if v12 != int32(1) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v15 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v18 = F_find_all_inheritors(m, v16, l4, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	if v18 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 <= int32(0) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v28 = v15
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v28<<(uint(int32(2))%32))))
	if v16 != v37 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L2
L11:
	;
	v40 = F_relation_open(m, v37, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v61 = v28 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v61 < v62 {
		v28 = v61
		goto L9
	} else {
		goto L22
	}
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+118)))
	if v43 == int32(116) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+24)))
	if v46 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_CheckTableNotInUse(m, v40, int32(520745))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	F_ATPrepCmd(m, l0, v40, l2, int32(0), int32(1), l4, l5)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_relation_close(m, v40, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	goto L10
L23:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(135867), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(475396), int32(4460), int32(393756))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
