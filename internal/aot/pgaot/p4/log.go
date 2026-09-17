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
	F_errdetail(m, int32(_a_F_LogChildExit_2), v13)
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
	v166 = int32(2839)
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
	v207 = int32(2861)
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
	v218 = int32(2872)
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
func F_assign_log_timezone(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_log_timezone[0])) = v4
	return
}
func F_log_heap_prune_and_freeze(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v304 int64
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v323 int32
	_ = v323
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	v9 = l8
	v11 = l10
	v13 = l12
	v14 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(_a_F_log_heap_prune_and_freeze_0)
	m.G0 = v24
	*(*int32)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_log_heap_prune_and_freeze[0]))) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_log_heap_prune_and_freeze[1]))) = uint8(v14)
	F_XLogBeginInsert(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < l6 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v37 = int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_log_heap_prune_and_freeze[1]))) = uint8(v37)
	F_pg_qsort(m, l5, l6, int32(12), int32(186))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v178 = v14
	goto L6
L6:
	;
	if int32(0) < v9 {
		goto L22
	} else {
		goto L23
	}
L7:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+608)) = v44
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+612)) = uint16(v46)
	v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+614)) = uint16(v48)
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	v51 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+618)) = uint16(v51)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+616)) = uint8(v50)
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)) = uint16(v54)
	if l6 != v51 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v76 = v24 + int32(608)
	v79 = v51
	v80 = int32(1)
	goto L11
L9:
	;
	v145 = v51
	goto L10
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+604)) = uint16(v145)
	F_XLogRegisterBufData(m, int32(0), v24+int32(604), int32(4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L20
	}
L11:
	;
	v85 = l5 + v80*int32(12)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v86 != v87 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v145 = v116
	goto L10
L13:
	;
	v120 = int32(1)
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v24+int32(16)+v80<<(uint(v120)%32)))) = uint16(v123)
	v126 = v80 + v120
	if v126 != l6 {
		v76 = v115
		v79 = v116
		v80 = v126
		goto L11
	} else {
		goto L19
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = v86
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)) = uint16(v103)
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+18)) = uint16(v105)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+8)))
	v108 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+22)) = uint16(v108)
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+20)) = uint8(v107)
	v115 = v76 + int32(12)
	v116 = v79 + v108
	goto L13
L15:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	if v89 != v90 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+6)))
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+6)))
	if v92 != v93 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+8)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+8)))
	if v95 != v96 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+10)))
	v100 = v98 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+10)) = uint16(v100)
	v115 = v76
	v116 = v79
	goto L13
L19:
	;
	goto L12
L20:
	;
	F_XLogRegisterBufData(m, int32(0), v24+int32(608), v145*int32(12))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v178 = v37
	goto L6
L22:
	;
	v187 = v178 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_log_heap_prune_and_freeze[1]))) = uint8(v187)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+602)) = uint16(v9)
	F_XLogRegisterBufData(m, int32(0), v24+int32(602), int32(2))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v201 = v178
	goto L24
L24:
	;
	if int32(0) < v11 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	F_XLogRegisterBufData(m, int32(0), l7, v9<<(uint(int32(2))%32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v201 = v187
	goto L24
L27:
	;
	v205 = v201 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_log_heap_prune_and_freeze[1]))) = uint8(v205)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+600)) = uint16(v11)
	F_XLogRegisterBufData(m, int32(0), v24+int32(600), int32(2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v219 = v201
	goto L29
L29:
	;
	if int32(0) < v13 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	F_XLogRegisterBufData(m, int32(0), l9, v11<<(uint(int32(1))%32))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v219 = v205
	goto L29
L32:
	;
	v223 = v219 | int32(-128)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_log_heap_prune_and_freeze[1]))) = uint8(v223)
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+598)) = uint16(v13)
	F_XLogRegisterBufData(m, int32(0), v24+int32(598), int32(2))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	v237 = v219
	goto L34
L34:
	;
	if int32(0) < l6 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	F_XLogRegisterBufData(m, int32(0), l11, v13<<(uint(int32(1))%32))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v237 = v223
	goto L34
L37:
	;
	F_XLogRegisterBufData(m, int32(0), v24+int32(16), l6<<(uint(int32(1))%32))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[2]))
	if v248 < int32(2) {
		v275 = v237
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	if l2|l3 != 0 {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v251)+118)))
	if v252 != int32(112) {
		v275 = v237
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L44
L44:
	;
	if base.B2i32(base.Ui32(v255) < base.Ui32(int32(_a_F_log_heap_prune_and_freeze_1))) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v260 == int32(0) {
		v275 = v237
		goto L41
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v272 = v237 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_log_heap_prune_and_freeze[1]))) = uint8(v272)
	v275 = v272
	goto L41
L48:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v263)+119)))
	switch v264 - int32(109) {
	case 0, 5:
		goto L49
	default:
		v275 = v237
		goto L41
	}
L49:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+104)))
	if v267 != int32(1) {
		v275 = v237
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	if l2 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	F_XLogRegisterData(m, v24+int32(_a_F_log_heap_prune_and_freeze_2), int32(2))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L60
	}
L54:
	;
	v279 = v275 | int32(8)
	goto L56
L55:
	;
	v279 = v275
	goto L56
L56:
	;
	if l3 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v282 = v279 | int32(4)
	goto L59
L58:
	;
	v282 = v279
	goto L59
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+uint32(_c_F_log_heap_prune_and_freeze[1]))) = uint8(v282)
	goto L53
L60:
	;
	if l2 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	F_XLogRegisterData(m, v24+int32(_a_F_log_heap_prune_and_freeze_3), int32(4))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if base.Ui32(l4) < base.Ui32(int32(3)) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	goto L63
L65:
	;
	v304 = F_XLogInsert(m, int32(9), (l4<<(uint(int32(4))%32)+int32(16))&int32(240))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L73
	}
L68:
	;
	if l1 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v323))) = base.I64_rotr(v304, int64(32))
	m.G0 = v24 + int32(_a_F_log_heap_prune_and_freeze_0)
	return
L70:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[3]))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v309+(l1^int32(-1))<<(uint(int32(2))%32))))
	v323 = v315
	goto L69
L71:
	;
	goto L72
L72:
	;
	v317 = *(*int32)(unsafe.Add(mBase, _c_F_log_heap_prune_and_freeze[4]))
	v323 = v317 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L69
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = l4
	F_errmsg_internal(m, int32(_a_F_log_heap_prune_and_freeze_4), v24)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_log_heap_prune_and_freeze_5), int32(2166), int32(_a_F_log_heap_prune_and_freeze_6))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
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
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int64
	_ = v151
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v4 = l3
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[0])))
	if v12 != int32(1) {
		v63 = int32(14)
		v70 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[1]))
		if int32(0)|base.B2i32(v70 == int32(15)) != 0 {
			v83 = int32(0)
			v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
			if v87 != int32(2) {
				v100 = v83
			} else {
				v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
				if v91&int32(1) != 0 {
					v100 = v83
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
					v100 = int32(0) | base.B2i32(v97 <= v63)
				}
			}
		} else {
			if v70 <= v63 {
				v100 = int32(1)
			} else {
				v83 = int32(0)
				v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
				if v87 != int32(2) {
					v100 = v83
				} else {
					v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
					if v91&int32(1) != 0 {
						v100 = v83
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
						v100 = int32(0) | base.B2i32(v97 <= v63)
					}
				}
			}
		}
		if v100 == int32(0) {
			v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
			if v134 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
				v142 = int32(40)
				v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
				mBase = m.M
				v146 = m.ExcPending
				if v146 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
					v148 = v145
					v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
					v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
					*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
					v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
					mBase = m.M
					v161 = m.ExcPending
					if v161 != 0 {
						return
					} else {
						v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
						if v162 == int32(0) {
							*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
						} else {
						}
						m.G0 = v9 + int32(112)
						return
					}
				}
			} else {
				v148 = v134
				v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
				v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
				*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
				v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
					return
				} else {
					v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
					if v162 == int32(0) {
						*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
					} else {
					}
					m.G0 = v9 + int32(112)
					return
				}
			}
		} else {
			v105 = v9 + int32(40)
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_GetRelationPath(m, v105, v106, v107, v108, int32(-1), l1)
			mBase = m.M
			v111 = m.ExcPending
			if v111 != 0 {
				return
			} else {
				v114 = F_errstart(m, int32(14), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return
				} else {
					if v114 == int32(0) {
						v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
						if v134 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
							v142 = int32(40)
							v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
								v148 = v145
								v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
								v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
								*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
								v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
									if v162 == int32(0) {
										*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
									} else {
									}
									m.G0 = v9 + int32(112)
									return
								}
							}
						} else {
							v148 = v134
							v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
							v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
							*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
							v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return
							} else {
								v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
								if v162 == int32(0) {
									*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
								} else {
								}
								m.G0 = v9 + int32(112)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v105
						if v4 != 0 {
							v122 = int32(_a_F_log_invalid_page_1)
						} else {
							v122 = int32(_a_F_log_invalid_page_2)
						}
						F_errmsg_internal(m, v122, v9)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return
						} else {
							if v4 != 0 {
								v128 = int32(93)
							} else {
								v128 = int32(96)
							}
							F_errfinish(m, int32(_a_F_log_invalid_page_3), v128, int32(_a_F_log_invalid_page_4))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return
							} else {
								v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
								if v134 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
									v142 = int32(40)
									v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
										v148 = v145
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
										v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
										*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
										v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return
										} else {
											v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
											if v162 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
											} else {
											}
											m.G0 = v9 + int32(112)
											return
										}
									}
								} else {
									v148 = v134
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
									v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
									*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
									v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
										if v162 == int32(0) {
											*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
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
							v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[6])))
							if v46 != 0 {
								v47 = int32(19)
							} else {
								v47 = int32(23)
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
									if int32(0)|base.B2i32(v70 == int32(15)) != 0 {
										v83 = int32(0)
										v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
										if v87 != int32(2) {
											v100 = v83
										} else {
											v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
											if v91&int32(1) != 0 {
												v100 = v83
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
												v100 = int32(0) | base.B2i32(v97 <= v63)
											}
										}
									} else {
										if v70 <= v63 {
											v100 = int32(1)
										} else {
											v83 = int32(0)
											v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
											if v87 != int32(2) {
												v100 = v83
											} else {
												v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
												if v91&int32(1) != 0 {
													v100 = v83
												} else {
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
													v100 = int32(0) | base.B2i32(v97 <= v63)
												}
											}
										}
									}
									if v100 == int32(0) {
										v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
										if v134 == int32(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
											v142 = int32(40)
											v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
											mBase = m.M
											v146 = m.ExcPending
											if v146 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
												v148 = v145
												v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
												v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return
												} else {
													v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v162 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											v148 = v134
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
											v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
											*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
											*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
											v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return
											} else {
												v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
												if v162 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
												} else {
												}
												m.G0 = v9 + int32(112)
												return
											}
										}
									} else {
										v105 = v9 + int32(40)
										v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_GetRelationPath(m, v105, v106, v107, v108, int32(-1), l1)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return
										} else {
											v114 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												if v114 == int32(0) {
													v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
													if v134 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
														v142 = int32(40)
														v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
															v148 = v145
															v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
															v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return
															} else {
																v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v162 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v148 = v134
														v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
														v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return
														} else {
															v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v162 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v105
													if v4 != 0 {
														v122 = int32(_a_F_log_invalid_page_1)
													} else {
														v122 = int32(_a_F_log_invalid_page_2)
													}
													F_errmsg_internal(m, v122, v9)
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														if v4 != 0 {
															v128 = int32(93)
														} else {
															v128 = int32(96)
														}
														F_errfinish(m, int32(_a_F_log_invalid_page_3), v128, int32(_a_F_log_invalid_page_4))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
															if v134 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
																v142 = int32(40)
																v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
																	v148 = v145
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
																	v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v161 = m.ExcPending
																	if v161 != 0 {
																		return
																	} else {
																		v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v162 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v148 = v134
																v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
																v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return
																} else {
																	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v162 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
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
											if int32(0)|base.B2i32(v70 == int32(15)) != 0 {
												v83 = int32(0)
												v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
												if v87 != int32(2) {
													v100 = v83
												} else {
													v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
													if v91&int32(1) != 0 {
														v100 = v83
													} else {
														v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
														v100 = int32(0) | base.B2i32(v97 <= v63)
													}
												}
											} else {
												if v70 <= v63 {
													v100 = int32(1)
												} else {
													v83 = int32(0)
													v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
													if v87 != int32(2) {
														v100 = v83
													} else {
														v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
														if v91&int32(1) != 0 {
															v100 = v83
														} else {
															v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
															v100 = int32(0) | base.B2i32(v97 <= v63)
														}
													}
												}
											}
											if v100 == int32(0) {
												v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
												if v134 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
													v142 = int32(40)
													v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
													mBase = m.M
													v146 = m.ExcPending
													if v146 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
														v148 = v145
														v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
														v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return
														} else {
															v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v162 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													v148 = v134
													v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
													v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
													*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
													*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
													v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
													mBase = m.M
													v161 = m.ExcPending
													if v161 != 0 {
														return
													} else {
														v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
														if v162 == int32(0) {
															*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
														} else {
														}
														m.G0 = v9 + int32(112)
														return
													}
												}
											} else {
												v105 = v9 + int32(40)
												v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_GetRelationPath(m, v105, v106, v107, v108, int32(-1), l1)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return
												} else {
													v114 = F_errstart(m, int32(14), int32(0))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														if v114 == int32(0) {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
															if v134 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
																v142 = int32(40)
																v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
																	v148 = v145
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
																	v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v161 = m.ExcPending
																	if v161 != 0 {
																		return
																	} else {
																		v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v162 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v148 = v134
																v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
																v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return
																} else {
																	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v162 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
																	} else {
																	}
																	m.G0 = v9 + int32(112)
																	return
																}
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v105
															if v4 != 0 {
																v122 = int32(_a_F_log_invalid_page_1)
															} else {
																v122 = int32(_a_F_log_invalid_page_2)
															}
															F_errmsg_internal(m, v122, v9)
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return
															} else {
																if v4 != 0 {
																	v128 = int32(93)
																} else {
																	v128 = int32(96)
																}
																F_errfinish(m, int32(_a_F_log_invalid_page_3), v128, int32(_a_F_log_invalid_page_4))
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return
																} else {
																	v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
																	if v134 == int32(0) {
																		*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
																		v142 = int32(40)
																		v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
																		mBase = m.M
																		v146 = m.ExcPending
																		if v146 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
																			v148 = v145
																			v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
																			v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																			*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																			v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
																			mBase = m.M
																			v161 = m.ExcPending
																			if v161 != 0 {
																				return
																			} else {
																				v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																				if v162 == int32(0) {
																					*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
																				} else {
																				}
																				m.G0 = v9 + int32(112)
																				return
																			}
																		}
																	} else {
																		v148 = v134
																		v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
																		v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																		*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																		v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
																		mBase = m.M
																		v161 = m.ExcPending
																		if v161 != 0 {
																			return
																		} else {
																			v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																			if v162 == int32(0) {
																				*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
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
					v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[6])))
					if v46 != 0 {
						v47 = int32(19)
					} else {
						v47 = int32(23)
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
							if int32(0)|base.B2i32(v70 == int32(15)) != 0 {
								v83 = int32(0)
								v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
								if v87 != int32(2) {
									v100 = v83
								} else {
									v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
									if v91&int32(1) != 0 {
										v100 = v83
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
										v100 = int32(0) | base.B2i32(v97 <= v63)
									}
								}
							} else {
								if v70 <= v63 {
									v100 = int32(1)
								} else {
									v83 = int32(0)
									v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
									if v87 != int32(2) {
										v100 = v83
									} else {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
										if v91&int32(1) != 0 {
											v100 = v83
										} else {
											v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
											v100 = int32(0) | base.B2i32(v97 <= v63)
										}
									}
								}
							}
							if v100 == int32(0) {
								v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
								if v134 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
									v142 = int32(40)
									v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
									mBase = m.M
									v146 = m.ExcPending
									if v146 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
										v148 = v145
										v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
										v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
										*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
										v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return
										} else {
											v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
											if v162 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
											} else {
											}
											m.G0 = v9 + int32(112)
											return
										}
									}
								} else {
									v148 = v134
									v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
									v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
									*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
									v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return
									} else {
										v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
										if v162 == int32(0) {
											*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
										} else {
										}
										m.G0 = v9 + int32(112)
										return
									}
								}
							} else {
								v105 = v9 + int32(40)
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								F_GetRelationPath(m, v105, v106, v107, v108, int32(-1), l1)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return
								} else {
									v114 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return
									} else {
										if v114 == int32(0) {
											v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
											if v134 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
												v142 = int32(40)
												v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
												mBase = m.M
												v146 = m.ExcPending
												if v146 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
													v148 = v145
													v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
													v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
													*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
													*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
													v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
													mBase = m.M
													v161 = m.ExcPending
													if v161 != 0 {
														return
													} else {
														v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
														if v162 == int32(0) {
															*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
														} else {
														}
														m.G0 = v9 + int32(112)
														return
													}
												}
											} else {
												v148 = v134
												v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
												v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return
												} else {
													v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v162 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v105
											if v4 != 0 {
												v122 = int32(_a_F_log_invalid_page_1)
											} else {
												v122 = int32(_a_F_log_invalid_page_2)
											}
											F_errmsg_internal(m, v122, v9)
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return
											} else {
												if v4 != 0 {
													v128 = int32(93)
												} else {
													v128 = int32(96)
												}
												F_errfinish(m, int32(_a_F_log_invalid_page_3), v128, int32(_a_F_log_invalid_page_4))
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return
												} else {
													v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
													if v134 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
														v142 = int32(40)
														v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
															v148 = v145
															v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
															v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return
															} else {
																v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v162 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v148 = v134
														v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
														v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return
														} else {
															v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v162 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
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
									if int32(0)|base.B2i32(v70 == int32(15)) != 0 {
										v83 = int32(0)
										v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
										if v87 != int32(2) {
											v100 = v83
										} else {
											v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
											if v91&int32(1) != 0 {
												v100 = v83
											} else {
												v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
												v100 = int32(0) | base.B2i32(v97 <= v63)
											}
										}
									} else {
										if v70 <= v63 {
											v100 = int32(1)
										} else {
											v83 = int32(0)
											v87 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[2]))
											if v87 != int32(2) {
												v100 = v83
											} else {
												v91 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_log_invalid_page[3])))
												if v91&int32(1) != 0 {
													v100 = v83
												} else {
													v97 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[4]))
													v100 = int32(0) | base.B2i32(v97 <= v63)
												}
											}
										}
									}
									if v100 == int32(0) {
										v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
										if v134 == int32(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
											v142 = int32(40)
											v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
											mBase = m.M
											v146 = m.ExcPending
											if v146 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
												v148 = v145
												v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
												v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return
												} else {
													v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v162 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											v148 = v134
											v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
											v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
											*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
											*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
											v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return
											} else {
												v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
												if v162 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
												} else {
												}
												m.G0 = v9 + int32(112)
												return
											}
										}
									} else {
										v105 = v9 + int32(40)
										v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_GetRelationPath(m, v105, v106, v107, v108, int32(-1), l1)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return
										} else {
											v114 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												if v114 == int32(0) {
													v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
													if v134 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
														v142 = int32(40)
														v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
															v148 = v145
															v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
															v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return
															} else {
																v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v162 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v148 = v134
														v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
														v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return
														} else {
															v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v162 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v105
													if v4 != 0 {
														v122 = int32(_a_F_log_invalid_page_1)
													} else {
														v122 = int32(_a_F_log_invalid_page_2)
													}
													F_errmsg_internal(m, v122, v9)
													mBase = m.M
													v124 = m.ExcPending
													if v124 != 0 {
														return
													} else {
														if v4 != 0 {
															v128 = int32(93)
														} else {
															v128 = int32(96)
														}
														F_errfinish(m, int32(_a_F_log_invalid_page_3), v128, int32(_a_F_log_invalid_page_4))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															v134 = *(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5]))
															if v134 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
																v142 = int32(40)
																v145 = F_hash_create(m, int32(_a_F_log_invalid_page_0), int32(100), v9+v142, v142)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_log_invalid_page[5])) = v145
																	v148 = v145
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
																	v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v161 = m.ExcPending
																	if v161 != 0 {
																		return
																	} else {
																		v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v162 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v148 = v134
																v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v149
																v151 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v151
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v160 = F_hash_search(m, v148, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return
																} else {
																	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v162 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v160)+20)) = uint8(v4)
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
					*(*int64)(unsafe.Add(mBase, uint32(l3))) = base.I64_rotr(v43, int64(32))
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
					F_errfinish(m, int32(_a_F_log_newpage_1), int32(320), int32(_a_F_log_newpage_2))
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
