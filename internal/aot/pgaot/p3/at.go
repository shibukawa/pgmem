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
	var v60 int32
	_ = v60
	var v63 int64
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
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v198 int32
	_ = v198
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int64
	_ = v235
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
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
	v393 = m.ExcPending
	if v393 != 0 {
		goto L5
	} else {
		goto L110
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L5
	} else {
		goto L107
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L103
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L5
	} else {
		goto L98
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L94
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L5
	} else {
		goto L90
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L86
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L5
	} else {
		goto L82
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
	F_errmsg(m, int32(_a_F_ATExecDropColumn_0), v21)
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
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecDropColumn[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v60
	v63 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecDropColumn[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v63
	goto L18
L28:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_2), int32(_a_F_ATExecDropColumn_3))
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
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+124)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v219
	F_add_exact_object_address(m, v21+int32(116), v33)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L76
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
	F_relation_close(m, v97, int32(3))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L75
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
	F_CheckTableNotInUse(m, v126, int32(_a_F_ATExecDropColumn_4))
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
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L72
	}
L63:
	;
	v161 = v145 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v144)+94)) = uint16(v161)
	F_CatalogTupleUpdate(m, v97, v138+int32(4), v138)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L5
	} else {
		goto L70
	}
L64:
	;
	if v145 != int32(1) {
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v158 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v144)+92)) = uint8(v158)
	goto L63
L67:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+92)))
	if v150 != 0 {
		goto L63
	} else {
		goto L68
	}
L68:
	;
	v153 = int32(1)
	F_ATExecDropColumn(m, v21+int32(116), v126, l2, l3, v153, v153, int32(0), l7, v33)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	goto L62
L70:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	goto L62
L72:
	;
	F_relation_close(m, v126, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L5
	} else {
		goto L73
	}
L73:
	;
	v175 = v119 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v175 < v176 {
		v119 = v175
		goto L51
	} else {
		goto L74
	}
L74:
	;
	goto L52
L75:
	;
	goto L42
L76:
	;
	if l5 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_performMultipleDeletions(m, v33, l3, int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v21)+124))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v233
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v21)+116))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v235
	goto L18
L80:
	;
	F_free_object_addresses(m, v33)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v265 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecDropColumn_5), v21+int32(16))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_6), int32(_a_F_ATExecDropColumn_3))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L86:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = l2
	F_errmsg(m, int32(_a_F_ATExecDropColumn_7), v21+int32(32))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_8), int32(_a_F_ATExecDropColumn_3))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+96)) = l2
	F_errmsg(m, int32(_a_F_ATExecDropColumn_9), v21+int32(96))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_10), int32(_a_F_ATExecDropColumn_3))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L5
	} else {
		goto L95
	}
L95:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v323 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecDropColumn_11), v21+int32(48))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L5
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_12), int32(_a_F_ATExecDropColumn_3))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L5
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
	F_errcode(m, int32(101056644))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L5
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_ATExecDropColumn_13), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L100
	}
L100:
	;
	F_errhint(m, int32(_a_F_ATExecDropColumn_14), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_15), int32(_a_F_ATExecDropColumn_3))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L5
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(_a_F_ATExecDropColumn_16), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_17), int32(_a_F_ATExecDropColumn_18))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L5
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = l2
	F_errmsg_internal(m, int32(_a_F_ATExecDropColumn_19), v21-int32(-64))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L5
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_20), int32(_a_F_ATExecDropColumn_3))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L5
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v124
	F_errmsg_internal(m, int32(_a_F_ATExecDropColumn_21), v21+int32(80))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_ATExecDropColumn_1), int32(_a_F_ATExecDropColumn_22), int32(_a_F_ATExecDropColumn_3))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L5
	} else {
		goto L112
	}
L112:
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
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
	v287 = m.ExcPending
	if v287 != 0 {
		goto L12
	} else {
		goto L72
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L12
	} else {
		goto L68
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L12
	} else {
		goto L64
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L12
	} else {
		goto L60
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
	v213 = m.ExcPending
	if v213 != 0 {
		goto L12
	} else {
		goto L55
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
	F_errmsg(m, int32(_a_F_ATExecDropIdentity_0), v16+int32(32))
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
	F_errfinish(m, int32(_a_F_ATExecDropIdentity_1), int32(_a_F_ATExecDropIdentity_2), int32(_a_F_ATExecDropIdentity_3))
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
	F_relation_close(m, v34, int32(3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecDropIdentity[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v78
	v81 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecDropIdentity[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v81
	goto L17
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecDropIdentity[2]))
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
	F_relation_close(m, v34, int32(3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L12
	} else {
		goto L36
	}
L36:
	;
	if base.B2i32(l5 == int32(0))|base.B2i32(v21 != int32(112)) != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	if l6 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v110 = F_find_inheritance_children(m, v109, l4)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L12
	} else {
		goto L39
	}
L39:
	;
	if v110 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v114 <= int32(0) {
		goto L37
	} else {
		goto L41
	}
L41:
	;
	v123 = int32(0)
	goto L42
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v110)+12))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133+v123<<(uint(int32(2))%32))))
	v139 = F_table_open(m, v137, int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L12
	} else {
		goto L44
	}
L43:
	;
	goto L37
L44:
	;
	v142 = int32(1)
	F_ATExecDropIdentity(m, v16+int32(68), v139, l2, int32(0), l4, v142, v142)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L12
	} else {
		goto L45
	}
L45:
	;
	F_relation_close(m, v139, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L12
	} else {
		goto L46
	}
L46:
	;
	v150 = v123 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	if v150 < v151 {
		v123 = v150
		goto L42
	} else {
		goto L47
	}
L47:
	;
	goto L43
L48:
	;
	v170 = F_getIdentitySequence(m, l1, v44, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L17
L51:
	;
	v174 = F_deleteDependencyRecordsForClass(m, int32(1259), v170, int32(1259), int32(105))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L12
	} else {
		goto L53
	}
L53:
	;
	v178 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = int32(1259)
	F_performDeletion(m, v16+int32(68), v178, int32(1))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L12
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_ATExecDropIdentity_4), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	F_errhint(m, int32(_a_F_ATExecDropIdentity_5), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_ATExecDropIdentity_1), int32(_a_F_ATExecDropIdentity_6), int32(_a_F_ATExecDropIdentity_3))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L12
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(_a_F_ATExecDropIdentity_7), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_ATExecDropIdentity_1), int32(_a_F_ATExecDropIdentity_8), int32(_a_F_ATExecDropIdentity_3))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L12
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L12
	} else {
		goto L65
	}
L65:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v253 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecDropIdentity_9), v16)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L12
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_ATExecDropIdentity_1), int32(_a_F_ATExecDropIdentity_10), int32(_a_F_ATExecDropIdentity_3))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L12
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L12
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = l2
	F_errmsg(m, int32(_a_F_ATExecDropIdentity_11), v16+int32(16))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L12
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_ATExecDropIdentity_1), int32(_a_F_ATExecDropIdentity_12), int32(_a_F_ATExecDropIdentity_3))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L12
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L12
	} else {
		goto L73
	}
L73:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v291 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecDropIdentity_13), v16+int32(48))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L12
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_ATExecDropIdentity_1), int32(_a_F_ATExecDropIdentity_14), int32(_a_F_ATExecDropIdentity_3))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L12
	} else {
		goto L75
	}
L75:
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
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
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
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int64
	_ = v374
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v520 int32
	_ = v520
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v641 int32
	_ = v641
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
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
	v26 = v19 + int32(48)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	F_ScanKeyInit(m, v26, int32(9), int32(3), int32(184), v30)
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
	v52 = F_systable_beginscan(m, v23, int32(2665), int32(1), int32(0), int32(3), v26)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L6:
	;
	F_systable_endscan(m, v52)
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L1
	} else {
		goto L138
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2606)
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v641
	goto L6
L8:
	;
	v504 = int32(0)
	F_set_attnotnull(m, v504, l2, v290, v504)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L118
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L114
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L110
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L106
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L102
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L97
	}
L14:
	;
	v54 = F_systable_getnext(m, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v54 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+22)))
	v58 = v56 + v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+72)))
	v61 = v59 - int32(99)
	if base.B2i32(base.Ui32(int32(11)) < base.Ui32(v61))|base.B2i32(int32(1)<<(uint(v61)%32)&int32(2057) == int32(0)) != 0 {
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
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L93
	}
L19:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+75)))
	if v71 == int32(0) {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+76)))
	if v74 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	switch v59 - int32(99) {
	case 0:
		goto L25
	default:
		goto L7
	case 3:
		goto L26
	case 11:
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v371 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecValidateConstraint[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v371
	v374 = *(*int64)(unsafe.Add(mBase, _c_F_ATExecValidateConstraint[1]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v374
	goto L6
L24:
	;
	v290 = F_extractNotNullColumn(m, v54)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L69
	}
L25:
	;
	if l5 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v58)+96))
	F_QueueFKConstraintValidation(m, l1, v23, l2, v79, v54, l6)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L7
L28:
	;
	v152 = F_palloc0(m, int32(32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L44
	}
L29:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+106)))
	if v82 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v83 = int32(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v86 = F_find_all_inheritors(m, v84, l6, v83)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v86 == int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v90 <= int32(0) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v98 = v83
	goto L34
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v98<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v113 != v114 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L28
L36:
	;
	if l4 == int32(0) {
		goto L11
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v132 = v98 + int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v132 < v133 {
		v98 = v132
		goto L34
	} else {
		goto L43
	}
L39:
	;
	v121 = F_table_open(m, v113, int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_ATExecValidateConstraint(m, v19+int32(196), l1, v121, l3, int32(0), int32(1), l6)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_relation_close(m, v121, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	goto L35
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v152)+4)) = int64(5)
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = l3
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	*(*int32)(unsafe.Add(mBase, uint32(v152)+20)) = v159
	v163 = F_SysCacheGetAttrNotNull(m, int32(19), v54, int32(28))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v165 = F_text_to_cstring(m, v163)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v167 = F_stringToNode(m, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v170 = F_expand_generated_columns_in_expr(m, v167, l2, int32(1))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+24)) = v170
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v174 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v251)+64))
	v263 = F_lappend(m, v262, v152)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L60
	}
L50:
	;
	v224 = F_palloc0(m, int32(140))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L57
	}
L51:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v177 <= int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v185 = int32(0)
	goto L53
L53:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v180+v185<<(uint(int32(2))%32))))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)))
	if v202 == v173 {
		v251 = v201
		goto L49
	} else {
		goto L55
	}
L54:
	;
	goto L50
L55:
	;
	v205 = v185 + int32(1)
	if v177 != v205 {
		v185 = v205
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v173
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+4)) = uint8(v230)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v233 = F_CreateTupleDescCopyConstr(m, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v224)+8)) = v233
	*(*int64)(unsafe.Add(mBase, uint32(v224)+88)) = int64(0)
	v238 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v224)+84)) = uint8(v238)
	v240 = int32(_a_F_ATExecValidateConstraint_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+96)) = uint16(v240)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v243 = F_lappend(m, v242, v224)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v243
	v251 = v224
	goto L49
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+64)) = v263
	F_CacheInvalidateRelcache(m, l2)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v268 = F_heap_copytuple(m, v54)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v268)+16))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+22)))
	v273 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v270+v271)+76)) = uint8(v273)
	F_CatalogTupleUpdate(m, v23, v268+int32(4), v268)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecValidateConstraint[2]))
	if v280 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v283 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v282, v283, v283, v283)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	F_pfree(m, v268)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	goto L66
L68:
	;
	goto L7
L69:
	;
	if l5 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v301 = int32(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v304 = F_find_all_inheritors(m, v302, l6, v301)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L76
	}
L71:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+106)))
	if v294 != int32(1) {
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v299 = F_get_attname(m, v297, v290, int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	goto L73
L75:
	;
	goto L8
L76:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v308 = F_get_attname(m, v306, v290, int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v304 == int32(0) {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v312 <= int32(0) {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	v318 = v301
	goto L80
L80:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v304)+12))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v331+v318<<(uint(int32(2))%32))))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	if v335 == v336 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	goto L8
L82:
	;
	v367 = v318 + int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v367 < v368 {
		v318 = v367
		goto L80
	} else {
		goto L92
	}
L83:
	;
	if l4 == int32(0) {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v340 = F_findNotNullConstraint(m, v335, v308)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v340 == int32(0) {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v340)+16))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+22)))
	v346 = v344 + v345
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v346)+76)))
	if v347 != 0 {
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v351 = F_table_open(m, v335, int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v355 = F_pstrdup(m, v346+int32(4))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_ATExecValidateConstraint(m, v19+int32(196), l1, v351, v355, int32(0), int32(1), l6)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	F_relation_close(m, v351, int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L82
L92:
	;
	goto L81
L93:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v383 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecValidateConstraint_1), v19)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_ATExecValidateConstraint_2), int32(_a_F_ATExecValidateConstraint_3), int32(_a_F_ATExecValidateConstraint_4))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v403 + int32(4)
	F_errmsg(m, int32(_a_F_ATExecValidateConstraint_5), v19+int32(16))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	F_errdetail(m, int32(_a_F_ATExecValidateConstraint_6), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_ATExecValidateConstraint_2), int32(_a_F_ATExecValidateConstraint_7), int32(_a_F_ATExecValidateConstraint_4))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errmsg(m, int32(_a_F_ATExecValidateConstraint_8), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(_a_F_ATExecValidateConstraint_2), int32(_a_F_ATExecValidateConstraint_9), int32(_a_F_ATExecValidateConstraint_4))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L106:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(_a_F_ATExecValidateConstraint_10), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_ATExecValidateConstraint_2), int32(_a_F_ATExecValidateConstraint_11), int32(_a_F_ATExecValidateConstraint_12))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(101056644))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_ATExecValidateConstraint_10), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_ATExecValidateConstraint_2), int32(_a_F_ATExecValidateConstraint_13), int32(_a_F_ATExecValidateConstraint_14))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v474 = F_get_rel_name(m, v335)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v308
	F_errmsg_internal(m, int32(_a_F_ATExecValidateConstraint_15), v19+int32(32))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_ATExecValidateConstraint_2), int32(_a_F_ATExecValidateConstraint_16), int32(_a_F_ATExecValidateConstraint_14))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v509 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v597 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v586)+76)) = uint8(v597)
	F_CacheInvalidateRelcache(m, l2)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L130
	}
L120:
	;
	v559 = F_palloc0(m, int32(140))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L1
	} else {
		goto L127
	}
L121:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v509)+4))
	if v512 <= int32(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v509)+12))
	v520 = int32(0)
	goto L123
L123:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v515+v520<<(uint(int32(2))%32))))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v536)))
	if v537 == v508 {
		v586 = v536
		goto L119
	} else {
		goto L125
	}
L124:
	;
	goto L120
L125:
	;
	v540 = v520 + int32(1)
	if v512 != v540 {
		v520 = v540
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v559))) = v508
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v564)+119)))
	*(*uint8)(unsafe.Add(mBase, uint32(v559)+4)) = uint8(v565)
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
	v568 = F_CreateTupleDescCopyConstr(m, v567)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v559)+8)) = v568
	*(*int64)(unsafe.Add(mBase, uint32(v559)+88)) = int64(0)
	v573 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v559)+84)) = uint8(v573)
	v575 = int32(_a_F_ATExecValidateConstraint_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v559)+96)) = uint16(v575)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v578 = F_lappend(m, v577, v559)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v578
	v586 = v559
	goto L119
L130:
	;
	v601 = F_heap_copytuple(m, v54)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v601)+16))
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+22)))
	v606 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v603+v604)+76)) = uint8(v606)
	F_CatalogTupleUpdate(m, v23, v601+int32(4), v601)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_ATExecValidateConstraint[2]))
	if v613 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v616 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2606), v615, v616, v616, v616)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	F_pfree(m, v601)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L137
	}
L136:
	;
	goto L135
L137:
	;
	goto L7
L138:
	;
	F_relation_close(m, v23, int32(3))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if l3 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
		if v13 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_ATPrepAddColumn_0), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ATPrepAddColumn_1), int32(_a_F_ATPrepAddColumn_2), int32(_a_F_ATPrepAddColumn_3))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
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
					v19 = int32(0)
					if l4|base.B2i32(l2 == v19) == v19 {
						v24 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l5)+29)) = uint8(v24)
					} else {
					}
					return
				}
			} else {
				v19 = int32(0)
				if l4|base.B2i32(l2 == v19) == v19 {
					v24 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l5)+29)) = uint8(v24)
				} else {
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
				v19 = int32(0)
				if l4|base.B2i32(l2 == v19) == v19 {
					v24 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l5)+29)) = uint8(v24)
				} else {
				}
				return
			}
		} else {
			v19 = int32(0)
			if l4|base.B2i32(l2 == v19) == v19 {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l5)+29)) = uint8(v24)
			} else {
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
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
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_ATSimplePermissions[0])))
			v34 = v32
		} else {
			v34 = int32(0)
		}
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			if v34 != 0 {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return
				} else {
					v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v101 + int32(4)
					F_errmsg(m, int32(_a_F_ATSimplePermissions_0), v8+int32(16))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return
					} else {
						v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v112 = int32(*(*int8)(unsafe.Add(mBase, uint32(v111)+119)))
						F_errdetail_relkind_not_supported(m, v112)
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ATSimplePermissions_1), int32(_a_F_ATSimplePermissions_2), int32(_a_F_ATSimplePermissions_3))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
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
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v39 + int32(4)
				F_errmsg_internal(m, int32(_a_F_ATSimplePermissions_4), v8)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_ATSimplePermissions_1), int32(_a_F_ATSimplePermissions_5), int32(_a_F_ATSimplePermissions_3))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
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
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
		v54 = *(*int32)(unsafe.Add(mBase, _c_F_ATSimplePermissions[1]))
		v55 = F_object_ownercheck(m, int32(1259), v52, v54)
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return
		} else {
			if v55 == int32(0) {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				v61 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60)+119)))
				switch v61 - int32(73) {
				case 0, 32:
					v71 = int32(20)
					v73 = v71
				default:
					v71 = int32(41)
					v73 = v71
				case 10:
					v73 = int32(37)
				case 29:
					v73 = int32(18)
				case 36:
					v73 = int32(23)
				case 45:
					v73 = int32(51)
				}
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
				F_aclcheck_error(m, int32(2), v73, v74+int32(4))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ATSimplePermissions[2])))
					if v80 == int32(0) {
						v84 = int32(1)
						v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
						if base.Ui32(v85) < base.Ui32(int32(_a_F_ATSimplePermissions_6)) {
							v94 = v84
						} else {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+68))
							if v89 == int32(99) {
								v94 = v84
							} else {
								v92 = F_isTempToastNamespace(m, v89)
								mBase = m.M
								v94 = v92
							}
						}
						if v94 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return
								} else {
									v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v127 + int32(4)
									F_errmsg(m, int32(_a_F_ATSimplePermissions_7), v8+int32(32))
									mBase = m.M
									v135 = m.ExcPending
									if v135 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ATSimplePermissions_1), int32(_a_F_ATSimplePermissions_8), int32(_a_F_ATSimplePermissions_3))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
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
				v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ATSimplePermissions[2])))
				if v80 == int32(0) {
					v84 = int32(1)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
					if base.Ui32(v85) < base.Ui32(int32(_a_F_ATSimplePermissions_6)) {
						v94 = v84
					} else {
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+68))
						if v89 == int32(99) {
							v94 = v84
						} else {
							v92 = F_isTempToastNamespace(m, v89)
							mBase = m.M
							v94 = v92
						}
					}
					if v94 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								v127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v127 + int32(4)
								F_errmsg(m, int32(_a_F_ATSimplePermissions_7), v8+int32(32))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ATSimplePermissions_1), int32(_a_F_ATSimplePermissions_8), int32(_a_F_ATSimplePermissions_3))
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
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
	F_CheckTableNotInUse(m, v40, int32(_a_F_ATSimpleRecursion_0))
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
	F_errmsg(m, int32(_a_F_ATSimpleRecursion_1), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_ATSimpleRecursion_2), int32(_a_F_ATSimpleRecursion_3), int32(_a_F_ATSimpleRecursion_4))
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
