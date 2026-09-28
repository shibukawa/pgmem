package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogChildExit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	v11 = m.G0
	v13 = v11 - int32(1088)
	m.G0 = v13
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v13 + int32(1088)
	return
L2:
	;
	F_errfinish(m, int32(_a_F_LogChildExit_0), v238, int32(_a_F_LogChildExit_1))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L39
	} else {
		goto L66
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v226
	v233 = F_errdetail(m, int32(_a_F_LogChildExit_2), v13)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L39
	} else {
		goto L65
	}
L4:
	;
	v168 = F_errstart(m, l0, int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L39
	} else {
		goto L44
	}
L5:
	;
	v16 = v13 - int32(-64)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_LogChildExit[0]))
	if v18 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v144 = int32(0)
	goto L7
L7:
	;
	v150 = F_errstart(m, l0, int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L39
	} else {
		goto L40
	}
L8:
	;
	v138 = l3 & int32(127)
	if v138 != 0 {
		goto L4
	} else {
		goto L38
	}
L9:
	;
	v136 = int32(0)
	goto L8
L10:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_LogChildExit[1]))
	if v22 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_LogChildExit[2]))
	if v27 <= int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v35 = v18
	v38 = int32(1)
	goto L13
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if l2 == v40 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L9
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+216))
	if base.Ui32(v42) < base.Ui32(v22) {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v113 = v38 + int32(1)
	if v113 <= v27 {
		v35 = v35 + int32(408)
		v38 = v113
		goto L13
	} else {
		goto L37
	}
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_LogChildExit[3]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_LogChildExit[4]))
	if base.Ui32(v22+v45-v48) < base.Ui32(v42) {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v51 == int32(0) {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v54 = int32(1024)
	if v54 < v48 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v57 = v54
	goto L23
L22:
	;
	v57 = v48
	goto L23
L23:
	;
	if v57 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v136 = v16
	goto L8
L25:
	;
	v61 = v57 - int32(1)
	if v61 == int32(0) {
		v98 = v16
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v103)
	goto L27
L29:
	;
	v64 = v16
	v65 = v42
	v67 = v61
	goto L30
L30:
	;
	v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == int32(0) {
		v98 = v64
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v98 = v95
	goto L28
L32:
	;
	if int32(31) < v69 {
		v89 = v69
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v91 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v89)
	v95 = v64 + v91
	v97 = v67 - v91
	if v97 != 0 {
		v64 = v95
		v65 = v65 + v91
		v67 = v97
		goto L30
	} else {
		goto L36
	}
L34:
	;
	v75 = v69 - int32(9)
	if base.Ui32(int32(4)) < base.Ui32(v75&int32(255)) {
		v89 = int32(63)
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v89 = base.I32_wrap_i64(int64(base.Ui64(int64(56895670793)) >> (uint(base.I64_extend_i32_u(v75<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L33
L36:
	;
	goto L31
L37:
	;
	goto L14
L38:
	;
	v144 = v136
	goto L7
L39:
	;
	return
L40:
	;
	if v150 == int32(0) {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(base.Ui32(l3)>>(uint(int32(8))%32)) & int32(255)
	F_errmsg(m, int32(_a_F_LogChildExit_3), v13+int32(16))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v166 = int32(2880)
	if v144 != 0 {
		v224 = v166
		v226 = v144
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v238 = v166
	goto L2
L44:
	;
	if base.Ui32(l3&int32(_a_F_LogChildExit_4)-int32(1)) <= base.Ui32(int32(254)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if v168 == int32(0) {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v168 == int32(0) {
		goto L1
	} else {
		goto L62
	}
L48:
	;
	v180 = int32(_a_F_LogChildExit_5)
	if base.Ui32(int32(-64)) <= base.Ui32(v138-int32(65)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
	F_errmsg(m, int32(_a_F_LogChildExit_6), v13+int32(32))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L39
	} else {
		goto L60
	}
L50:
	;
	v185 = v138
	v186 = v180
	goto L53
L51:
	;
	v194 = v180
	goto L52
L52:
	;
	if v194 != 0 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v189 = v186 + int32(1)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
	if v190 != 0 {
		v186 = v189
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v194 = v189
	goto L52
L55:
	;
	v192 = v185 - int32(1)
	if v192 != 0 {
		v185 = v192
		v186 = v189
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v197 = v194
	goto L59
L58:
	;
	v197 = int32(_a_F_LogChildExit_7)
	goto L59
L59:
	;
	goto L49
L60:
	;
	v207 = int32(2902)
	if v136 != 0 {
		v224 = v207
		v226 = v136
		goto L3
	} else {
		goto L61
	}
L61:
	;
	v238 = v207
	goto L2
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
	F_errmsg(m, int32(_a_F_LogChildExit_8), v13+int32(48))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L39
	} else {
		goto L63
	}
L63:
	;
	v218 = int32(2913)
	if v136 == int32(0) {
		v238 = v218
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v224 = v218
	v226 = v136
	goto L3
L65:
	;
	v238 = v224
	goto L2
L66:
	;
	goto L1
}
func F_assign_log_connections(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_connections[0])) = v4
	return
}
func F_assign_log_min_messages(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[0])) = v4
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[1])) = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[2])) = v10
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[3])) = v13
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[4])) = v16
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[5])) = v19
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[6])) = v22
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[7])) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[8])) = v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[9])) = v31
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[10])) = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[11])) = v37
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[12])) = v40
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[13])) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[14])) = v46
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[15])) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[16])) = v52
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_min_messages[17])) = v55
	return
}
func F_assign_log_timezone(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_timezone[0])) = v4
	return
}
func F_log_heap_prune_and_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v241 int32
	_ = v241
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v394 int64
	_ = v394
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	v11 = l10
	v13 = l12
	v15 = l14
	v16 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(_a_F_log_heap_prune_and_freeze_0)
	m.G0 = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[0]))) = l4
	if l1 < v16 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v49 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v49)
	v52 = l3 & int32(3)
	v53 = int32(1)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
	if v54 == int64(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[1]))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v34+(l1^int32(-1))<<(uint(int32(2))%32))))
	v48 = v40
	goto L1
L3:
	;
	goto L4
L4:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[2]))
	v48 = v42 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L17
	}
L6:
	;
	v77 = int32(9)
	v78 = v53
	goto L5
L7:
	;
	goto L8
L8:
	;
	v58 = int32(8)
	v59 = int32(0)
	if base.B2i32(v59 < v11)|base.B2i32(v59 < v13)|(l8|base.B2i32(v59 < v15)) != 0 {
		v77 = v58
		v78 = v53
		goto L5
	} else {
		goto L9
	}
L9:
	;
	if v52 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[4])))
	if v69 != 0 {
		v77 = v58
		v78 = v53
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v77 = int32(10)
	v78 = int32(0)
	goto L5
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[5]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+268))
	goto L14
L14:
	;
	if v72 != int32(0) {
		v77 = v58
		v78 = v53
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	return
L17:
	;
	F_XLogRegisterBuffer(m, int32(0), l1, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v52 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_XLogRegisterBuffer(m, int32(1), l2, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if int32(0) < l8 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L21
L23:
	;
	v90 = int32(16)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v90)
	F_pg_qsort(m, l7, l8, int32(12), int32(188))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L16
	} else {
		goto L26
	}
L24:
	;
	v241 = v16
	goto L25
L25:
	;
	if int32(0) < v11 {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+608)) = v97
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+612)) = uint16(v99)
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+614)) = uint16(v101)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7)+8)))
	v104 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+618)) = uint16(v104)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+616)) = uint8(v103)
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l7)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)) = uint16(v107)
	if l8 != v104 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v116 = v28 + int32(608)
	v133 = v104
	v134 = int32(1)
	goto L30
L28:
	;
	v203 = v104
	goto L29
L29:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+604)) = uint16(v203)
	F_XLogRegisterBufData(m, int32(0), v28+int32(604), int32(4))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L16
	} else {
		goto L39
	}
L30:
	;
	v142 = l7 + v134*int32(12)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v142)))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v143 != v144 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v203 = v173
	goto L29
L32:
	;
	v177 = int32(1)
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v28+int32(16)+v134<<(uint(v177)%32)))) = uint16(v180)
	v183 = v134 + v177
	if v183 != l8 {
		v116 = v172
		v133 = v173
		v134 = v183
		goto L30
	} else {
		goto L38
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v116)+12)) = v143
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+16)) = uint16(v160)
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+18)) = uint16(v162)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+8)))
	v165 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+22)) = uint16(v165)
	*(*uint8)(unsafe.Add(mBase, uint32(v116)+20)) = uint8(v164)
	v172 = v116 + int32(12)
	v173 = v133 + v165
	goto L32
L34:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+4)))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+4)))
	if v146 != v147 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+6)))
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v142)+6)))
	if v149 != v150 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+8)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+8)))
	if v152 != v153 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v116)+10)))
	v157 = v155 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v116)+10)) = uint16(v157)
	v172 = v116
	v173 = v133
	goto L32
L38:
	;
	goto L31
L39:
	;
	F_XLogRegisterBufData(m, int32(0), v28+int32(608), v203*int32(12))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v241 = v90
	goto L25
L41:
	;
	v252 = v241 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v252)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+602)) = uint16(v11)
	F_XLogRegisterBufData(m, int32(0), v28+int32(602), int32(2))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L16
	} else {
		goto L44
	}
L42:
	;
	v266 = v241
	goto L43
L43:
	;
	if int32(0) < v13 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	F_XLogRegisterBufData(m, int32(0), l9, v11<<(uint(int32(2))%32))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	v266 = v252
	goto L43
L46:
	;
	v270 = v266 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v270)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+600)) = uint16(v13)
	F_XLogRegisterBufData(m, int32(0), v28+int32(600), int32(2))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L16
	} else {
		goto L49
	}
L47:
	;
	v284 = v266
	goto L48
L48:
	;
	if int32(0) < v15 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	F_XLogRegisterBufData(m, int32(0), l11, v13<<(uint(int32(1))%32))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	v284 = v270
	goto L48
L51:
	;
	v288 = v284 | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v288)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+598)) = uint16(v15)
	F_XLogRegisterBufData(m, int32(0), v28+int32(598), int32(2))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L16
	} else {
		goto L54
	}
L52:
	;
	v302 = v284
	goto L53
L53:
	;
	if int32(0) < l8 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	F_XLogRegisterBufData(m, int32(0), l13, v15<<(uint(int32(1))%32))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	v302 = v288
	goto L53
L56:
	;
	F_XLogRegisterBufData(m, int32(0), v28+int32(16), l8<<(uint(int32(1))%32))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L16
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if l3&int32(1) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L58
L60:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[6]))
	if v328 <= int32(1) {
		goto L66
	} else {
		goto L67
	}
L61:
	;
	v326 = v302
	goto L60
L62:
	;
	goto L63
L63:
	;
	v317 = v302 | int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v317)
	if l3&int32(2) == int32(0) {
		v326 = v317
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v324 = v302 | int32(768)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v324)
	v326 = v324
	goto L60
L65:
	;
	if l4|l5 != 0 {
		goto L83
	} else {
		goto L84
	}
L66:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[7])))
	if v332&int32(1) == int32(0) {
		v364 = v326
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v337)+118)))
	if v338 != int32(112) {
		v364 = v326
		goto L65
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	if v328 <= int32(0) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v343 != 0 {
		v364 = v326
		goto L65
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L76
L74:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v344 != 0 {
		v364 = v326
		goto L65
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	if base.B2i32(base.Ui32(v345) < base.Ui32(int32(_a_F_log_heap_prune_and_freeze_1))) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v350 == int32(0) {
		v364 = v326
		goto L65
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v362 = v326 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v362)
	v364 = v362
	goto L65
L80:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353)+119)))
	switch v354 - int32(109) {
	case 0, 5:
		goto L81
	default:
		v364 = v326
		goto L65
	}
L81:
	;
	v357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350)+112)))
	if v357 != int32(1) {
		v364 = v326
		goto L65
	} else {
		goto L82
	}
L82:
	;
	goto L79
L83:
	;
	if l4 != 0 {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	F_XLogRegisterData(m, v28+int32(_a_F_log_heap_prune_and_freeze_2), int32(2))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L16
	} else {
		goto L92
	}
L86:
	;
	v369 = v364 | int32(8)
	goto L88
L87:
	;
	v369 = v364
	goto L88
L88:
	;
	if l5 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v372 = v369 | int32(4)
	goto L91
L90:
	;
	v372 = v369
	goto L91
L91:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_log_heap_prune_and_freeze[3]))) = uint16(v372)
	goto L85
L92:
	;
	if l4 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	F_XLogRegisterData(m, v28+int32(_a_F_log_heap_prune_and_freeze_3), int32(4))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L16
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if base.Ui32(l6) < base.Ui32(int32(3)) {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L95
L97:
	;
	v394 = F_XLogInsert(m, int32(9), (l6<<(uint(int32(4))%32)+int32(16))&int32(240))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L16
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L16
	} else {
		goto L111
	}
L100:
	;
	if v52 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if l2 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	if v78 != 0 {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v413))) = base.I64_rotl(v394, int64(32))
	goto L103
L105:
	;
	v399 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[1]))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v399+(l2^int32(-1))<<(uint(int32(2))%32))))
	v413 = v405
	goto L104
L106:
	;
	goto L107
L107:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[2]))
	v413 = v407 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L104
L108:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v48))) = base.I64_rotl(v394, int64(32))
	goto L110
L109:
	;
	goto L110
L110:
	;
	m.G0 = v28 + int32(_a_F_log_heap_prune_and_freeze_0)
	return
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = l6
	F_errmsg_internal(m, int32(_a_F_log_heap_prune_and_freeze_4), v28)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L16
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_log_heap_prune_and_freeze_5), int32(2743), int32(_a_F_log_heap_prune_and_freeze_6))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L16
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_log_invalid_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
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
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	v4 = l3
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[0])))
	if v12 != int32(1) {
		v63 = int32(14)
		v70 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[1]))
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_c_F_log_invalid_page[2])))
		if int32(0)|base.B2i32(v73 == int32(15)) != 0 {
			v86 = int32(0)
			v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
			if v90 != int32(2) {
				v103 = v86
			} else {
				v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
				if v94&int32(1) != 0 {
					v103 = v86
				} else {
					v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
					v103 = int32(0) | base.B2i32(v100 <= v63)
				}
			}
		} else {
			if v73 <= v63 {
				v103 = int32(1)
			} else {
				v86 = int32(0)
				v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
				if v90 != int32(2) {
					v103 = v86
				} else {
					v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
					if v94&int32(1) != 0 {
						v103 = v86
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
						v103 = int32(0) | base.B2i32(v100 <= v63)
					}
				}
			}
		}
		if v103 == int32(0) {
			v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
			if v137 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
				v145 = int32(40)
				v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
				mBase = m.M
				v149 = m.ExcPending
				if v149 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
					v151 = v148
					v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
					v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
					*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
					v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
					mBase = m.M
					v164 = m.ExcPending
					if v164 != 0 {
						return
					} else {
						v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
						if v165 == int32(0) {
							*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
						} else {
						}
						m.G0 = v9 + int32(112)
						return
					}
				}
			} else {
				v151 = v137
				v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
				v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
				*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
				v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return
				} else {
					v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
					if v165 == int32(0) {
						*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
					} else {
					}
					m.G0 = v9 + int32(112)
					return
				}
			}
		} else {
			v108 = v9 + int32(40)
			v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_GetRelationPath(m, v108, v109, v110, v111, int32(-1), l1)
			mBase = m.M
			v114 = m.ExcPending
			if v114 != 0 {
				return
			} else {
				v117 = F_errstart(m, int32(14), int32(0))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return
				} else {
					if v117 == int32(0) {
						v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
						if v137 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
							v145 = int32(40)
							v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
								v151 = v148
								v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
								v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
								*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
								v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return
								} else {
									v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
									if v165 == int32(0) {
										*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
									} else {
									}
									m.G0 = v9 + int32(112)
									return
								}
							}
						} else {
							v151 = v137
							v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
							v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
							*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
							v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return
							} else {
								v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
								if v165 == int32(0) {
									*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
								} else {
								}
								m.G0 = v9 + int32(112)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v108
						if v4 != 0 {
							v125 = int32(_a_F_log_invalid_page_1)
						} else {
							v125 = int32(_a_F_log_invalid_page_2)
						}
						F_errmsg_internal(m, v125, v9)
						mBase = m.M
						v127 = m.ExcPending
						if v127 != 0 {
							return
						} else {
							if v4 != 0 {
								v131 = int32(93)
							} else {
								v131 = int32(96)
							}
							F_errfinish(m, int32(_a_F_log_invalid_page_3), v131, int32(_a_F_log_invalid_page_4))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return
							} else {
								v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
								if v137 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
									v145 = int32(40)
									v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
										v151 = v148
										v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
										v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
										*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
										v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
											if v165 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
											} else {
											}
											m.G0 = v9 + int32(112)
											return
										}
									}
								} else {
									v151 = v137
									v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
									v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
									*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
									v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return
									} else {
										v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
										if v165 == int32(0) {
											*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
										} else {
										}
										m.G0 = v9 + int32(112)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v16 = v9 + int32(40)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_GetRelationPath(m, v16, v17, v18, v19, int32(-1), l1)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			v25 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v16
					if v4 != 0 {
						v31 = int32(_a_F_log_invalid_page_1)
					} else {
						v31 = int32(_a_F_log_invalid_page_2)
					}
					F_errmsg_internal(m, v31, v9+int32(16))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						if v4 != 0 {
							v39 = int32(93)
						} else {
							v39 = int32(96)
						}
						F_errfinish(m, int32(_a_F_log_invalid_page_3), v39, int32(_a_F_log_invalid_page_4))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[7])))
							if v46 != 0 {
								v47 = int32(19)
							} else {
								v47 = int32(24)
							}
							v49 = F_errstart(m, v47, int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								if v49 == int32(0) {
									v63 = int32(14)
									v70 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[1]))
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_c_F_log_invalid_page[2])))
									if int32(0)|base.B2i32(v73 == int32(15)) != 0 {
										v86 = int32(0)
										v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
										if v90 != int32(2) {
											v103 = v86
										} else {
											v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
											if v94&int32(1) != 0 {
												v103 = v86
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
												v103 = int32(0) | base.B2i32(v100 <= v63)
											}
										}
									} else {
										if v73 <= v63 {
											v103 = int32(1)
										} else {
											v86 = int32(0)
											v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
											if v90 != int32(2) {
												v103 = v86
											} else {
												v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
												if v94&int32(1) != 0 {
													v103 = v86
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
													v103 = int32(0) | base.B2i32(v100 <= v63)
												}
											}
										}
									}
									if v103 == int32(0) {
										v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
										if v137 == int32(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
											v145 = int32(40)
											v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
												v151 = v148
												v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
												v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v165 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											v151 = v137
											v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
											v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
											*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
											*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
											v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
												if v165 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
												} else {
												}
												m.G0 = v9 + int32(112)
												return
											}
										}
									} else {
										v108 = v9 + int32(40)
										v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_GetRelationPath(m, v108, v109, v110, v111, int32(-1), l1)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											v117 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return
											} else {
												if v117 == int32(0) {
													v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
													if v137 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
														v145 = int32(40)
														v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
															v151 = v148
															v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
															v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return
															} else {
																v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v165 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v151 = v137
														v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
														v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return
														} else {
															v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v165 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v108
													if v4 != 0 {
														v125 = int32(_a_F_log_invalid_page_1)
													} else {
														v125 = int32(_a_F_log_invalid_page_2)
													}
													F_errmsg_internal(m, v125, v9)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return
													} else {
														if v4 != 0 {
															v131 = int32(93)
														} else {
															v131 = int32(96)
														}
														F_errfinish(m, int32(_a_F_log_invalid_page_3), v131, int32(_a_F_log_invalid_page_4))
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
															if v137 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
																v145 = int32(40)
																v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
																mBase = m.M
																v149 = m.ExcPending
																if v149 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
																	v151 = v148
																	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
																	v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v164 = m.ExcPending
																	if v164 != 0 {
																		return
																	} else {
																		v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v165 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v151 = v137
																v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
																v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v164 = m.ExcPending
																if v164 != 0 {
																	return
																} else {
																	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v165 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																	} else {
																	}
																	m.G0 = v9 + int32(112)
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									F_errmsg_internal(m, int32(_a_F_log_invalid_page_5), int32(0))
									mBase = m.M
									v56 = m.ExcPending
									if v56 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_log_invalid_page_3), int32(120), int32(_a_F_log_invalid_page_6))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											v63 = int32(14)
											v70 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[1]))
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_c_F_log_invalid_page[2])))
											if int32(0)|base.B2i32(v73 == int32(15)) != 0 {
												v86 = int32(0)
												v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
												if v90 != int32(2) {
													v103 = v86
												} else {
													v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
													if v94&int32(1) != 0 {
														v103 = v86
													} else {
														v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
														v103 = int32(0) | base.B2i32(v100 <= v63)
													}
												}
											} else {
												if v73 <= v63 {
													v103 = int32(1)
												} else {
													v86 = int32(0)
													v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
													if v90 != int32(2) {
														v103 = v86
													} else {
														v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
														if v94&int32(1) != 0 {
															v103 = v86
														} else {
															v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
															v103 = int32(0) | base.B2i32(v100 <= v63)
														}
													}
												}
											}
											if v103 == int32(0) {
												v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
												if v137 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
													v145 = int32(40)
													v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
														v151 = v148
														v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
														v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return
														} else {
															v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v165 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													v151 = v137
													v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
													v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
													*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
													*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
													v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
														if v165 == int32(0) {
															*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
														} else {
														}
														m.G0 = v9 + int32(112)
														return
													}
												}
											} else {
												v108 = v9 + int32(40)
												v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_GetRelationPath(m, v108, v109, v110, v111, int32(-1), l1)
												mBase = m.M
												v114 = m.ExcPending
												if v114 != 0 {
													return
												} else {
													v117 = F_errstart(m, int32(14), int32(0))
													mBase = m.M
													v118 = m.ExcPending
													if v118 != 0 {
														return
													} else {
														if v117 == int32(0) {
															v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
															if v137 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
																v145 = int32(40)
																v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
																mBase = m.M
																v149 = m.ExcPending
																if v149 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
																	v151 = v148
																	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
																	v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v164 = m.ExcPending
																	if v164 != 0 {
																		return
																	} else {
																		v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v165 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v151 = v137
																v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
																v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v164 = m.ExcPending
																if v164 != 0 {
																	return
																} else {
																	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v165 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																	} else {
																	}
																	m.G0 = v9 + int32(112)
																	return
																}
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v108
															if v4 != 0 {
																v125 = int32(_a_F_log_invalid_page_1)
															} else {
																v125 = int32(_a_F_log_invalid_page_2)
															}
															F_errmsg_internal(m, v125, v9)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return
															} else {
																if v4 != 0 {
																	v131 = int32(93)
																} else {
																	v131 = int32(96)
																}
																F_errfinish(m, int32(_a_F_log_invalid_page_3), v131, int32(_a_F_log_invalid_page_4))
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return
																} else {
																	v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
																	if v137 == int32(0) {
																		*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
																		v145 = int32(40)
																		v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
																		mBase = m.M
																		v149 = m.ExcPending
																		if v149 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
																			v151 = v148
																			v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
																			v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																			*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																			v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
																			mBase = m.M
																			v164 = m.ExcPending
																			if v164 != 0 {
																				return
																			} else {
																				v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																				if v165 == int32(0) {
																					*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																				} else {
																				}
																				m.G0 = v9 + int32(112)
																				return
																			}
																		}
																	} else {
																		v151 = v137
																		v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
																		v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																		*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																		v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
																		mBase = m.M
																		v164 = m.ExcPending
																		if v164 != 0 {
																			return
																		} else {
																			v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																			if v165 == int32(0) {
																				*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																			} else {
																			}
																			m.G0 = v9 + int32(112)
																			return
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
								}
							}
						}
					}
				} else {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[7])))
					if v46 != 0 {
						v47 = int32(19)
					} else {
						v47 = int32(24)
					}
					v49 = F_errstart(m, v47, int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return
					} else {
						if v49 == int32(0) {
							v63 = int32(14)
							v70 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[1]))
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_c_F_log_invalid_page[2])))
							if int32(0)|base.B2i32(v73 == int32(15)) != 0 {
								v86 = int32(0)
								v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
								if v90 != int32(2) {
									v103 = v86
								} else {
									v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
									if v94&int32(1) != 0 {
										v103 = v86
									} else {
										v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
										v103 = int32(0) | base.B2i32(v100 <= v63)
									}
								}
							} else {
								if v73 <= v63 {
									v103 = int32(1)
								} else {
									v86 = int32(0)
									v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
									if v90 != int32(2) {
										v103 = v86
									} else {
										v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
										if v94&int32(1) != 0 {
											v103 = v86
										} else {
											v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
											v103 = int32(0) | base.B2i32(v100 <= v63)
										}
									}
								}
							}
							if v103 == int32(0) {
								v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
								if v137 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
									v145 = int32(40)
									v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
										v151 = v148
										v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
										v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
										*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
										v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
											if v165 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
											} else {
											}
											m.G0 = v9 + int32(112)
											return
										}
									}
								} else {
									v151 = v137
									v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
									v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
									*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
									v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return
									} else {
										v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
										if v165 == int32(0) {
											*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
										} else {
										}
										m.G0 = v9 + int32(112)
										return
									}
								}
							} else {
								v108 = v9 + int32(40)
								v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								F_GetRelationPath(m, v108, v109, v110, v111, int32(-1), l1)
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return
								} else {
									v117 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v118 = m.ExcPending
									if v118 != 0 {
										return
									} else {
										if v117 == int32(0) {
											v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
											if v137 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
												v145 = int32(40)
												v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
													v151 = v148
													v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
													v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
													*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
													*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
													v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return
													} else {
														v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
														if v165 == int32(0) {
															*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
														} else {
														}
														m.G0 = v9 + int32(112)
														return
													}
												}
											} else {
												v151 = v137
												v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
												v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v165 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v108
											if v4 != 0 {
												v125 = int32(_a_F_log_invalid_page_1)
											} else {
												v125 = int32(_a_F_log_invalid_page_2)
											}
											F_errmsg_internal(m, v125, v9)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return
											} else {
												if v4 != 0 {
													v131 = int32(93)
												} else {
													v131 = int32(96)
												}
												F_errfinish(m, int32(_a_F_log_invalid_page_3), v131, int32(_a_F_log_invalid_page_4))
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return
												} else {
													v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
													if v137 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
														v145 = int32(40)
														v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
															v151 = v148
															v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
															v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return
															} else {
																v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v165 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v151 = v137
														v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
														v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return
														} else {
															v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v165 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							F_errmsg_internal(m, int32(_a_F_log_invalid_page_5), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_log_invalid_page_3), int32(120), int32(_a_F_log_invalid_page_6))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v63 = int32(14)
									v70 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[1]))
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v70<<(uint(int32(2))%32))+uint32(_c_F_log_invalid_page[2])))
									if int32(0)|base.B2i32(v73 == int32(15)) != 0 {
										v86 = int32(0)
										v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
										if v90 != int32(2) {
											v103 = v86
										} else {
											v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
											if v94&int32(1) != 0 {
												v103 = v86
											} else {
												v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
												v103 = int32(0) | base.B2i32(v100 <= v63)
											}
										}
									} else {
										if v73 <= v63 {
											v103 = int32(1)
										} else {
											v86 = int32(0)
											v90 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[3]))
											if v90 != int32(2) {
												v103 = v86
											} else {
												v94 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[4])))
												if v94&int32(1) != 0 {
													v103 = v86
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
													v103 = int32(0) | base.B2i32(v100 <= v63)
												}
											}
										}
									}
									if v103 == int32(0) {
										v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
										if v137 == int32(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
											v145 = int32(40)
											v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
											mBase = m.M
											v149 = m.ExcPending
											if v149 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
												v151 = v148
												v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
												v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return
												} else {
													v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v165 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											v151 = v137
											v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
											v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
											*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
											*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
											v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
												if v165 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
												} else {
												}
												m.G0 = v9 + int32(112)
												return
											}
										}
									} else {
										v108 = v9 + int32(40)
										v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_GetRelationPath(m, v108, v109, v110, v111, int32(-1), l1)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return
										} else {
											v117 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v118 = m.ExcPending
											if v118 != 0 {
												return
											} else {
												if v117 == int32(0) {
													v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
													if v137 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
														v145 = int32(40)
														v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
															v151 = v148
															v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
															v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return
															} else {
																v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v165 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v151 = v137
														v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
														v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v164 = m.ExcPending
														if v164 != 0 {
															return
														} else {
															v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v165 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v108
													if v4 != 0 {
														v125 = int32(_a_F_log_invalid_page_1)
													} else {
														v125 = int32(_a_F_log_invalid_page_2)
													}
													F_errmsg_internal(m, v125, v9)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return
													} else {
														if v4 != 0 {
															v131 = int32(93)
														} else {
															v131 = int32(96)
														}
														F_errfinish(m, int32(_a_F_log_invalid_page_3), v131, int32(_a_F_log_invalid_page_4))
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															v137 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6]))
															if v137 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = int64(103079215124)
																v145 = int32(40)
																v148 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int64(100), v9+v145, v145)
																mBase = m.M
																v149 = m.ExcPending
																if v149 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[6])) = v148
																	v151 = v148
																	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
																	v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v164 = m.ExcPending
																	if v164 != 0 {
																		return
																	} else {
																		v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v165 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v151 = v137
																v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v152
																v154 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v154
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v163 = F_hash_search(m, v151, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v164 = m.ExcPending
																if v164 != 0 {
																	return
																} else {
																	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v165 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v163)+20)) = uint8(v4)
																	} else {
																	}
																	m.G0 = v9 + int32(112)
																	return
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
						}
					}
				}
			}
		}
	}
}
func F_log_newpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int64
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	F_XLogBeginInsert(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage[0]))
		if v11 <= int32(0) {
			*(*int32)(unsafe.Add(mBase, _c_F_log_newpage[0])) = int32(1)
		} else {
		}
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage[1]))
		if int32(0) < v18 {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage[2]))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v23
			v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v22)+4)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
			if l4 != 0 {
				v32 = int32(9)
			} else {
				v32 = int32(1)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v32)
			v34 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v34
			v36 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v36)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v22 + int32(32)
			v43 = F_XLogInsert(m, v34, int32(176))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+14)))
				if v45 != 0 {
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = base.I64_rotl(v43, int64(32))
				} else {
				}
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_log_newpage_0), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_log_newpage_1), int32(328), int32(_a_F_log_newpage_2))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_log_smgrcreate(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v10
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	F_XLogBeginInsert(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		F_XLogRegisterData(m, v6, int32(16))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v20 = F_XLogInsert(m, int32(2), int32(17))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
