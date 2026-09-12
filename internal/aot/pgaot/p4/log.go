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
	var v36 int32
	_ = v36
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
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
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
	F_errfinish(m, int32(492102), v240, int32(99236))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L39
	} else {
		goto L65
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v228
	F_errdetail(m, int32(201906), v13)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L39
	} else {
		goto L64
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[770]))
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[771]))
	if v22 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[772]))
	if v27 <= int32(0) {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v35 = v18
	v36 = int32(1)
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
	v113 = v36 + int32(1)
	if v113 <= v27 {
		v35 = v35 + int32(408)
		v36 = v113
		goto L13
	} else {
		goto L37
	}
L18:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	v48 = *(*int32)(unsafe.Add(mBase, _consts[774]))
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
	F_errmsg(m, int32(474710), v13+int32(16))
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
		v226 = v166
		v228 = v144
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v240 = v166
	goto L2
L44:
	;
	if base.Ui32(l3&int32(65535)-int32(1)) <= base.Ui32(int32(254)) {
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
		goto L61
	}
L48:
	;
	v180 = int32(4084368)
	if base.Ui32(v138-int32(65)) < base.Ui32(int32(-64)) {
		v197 = v180
		goto L50
	} else {
		goto L51
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = l1
	F_errmsg(m, int32(203511), v13+int32(32))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L39
	} else {
		goto L59
	}
L50:
	;
	if v197 != 0 {
		goto L56
	} else {
		goto L57
	}
L51:
	;
	v186 = v138
	v187 = v180
	goto L52
L52:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v191 = v187 + int32(1)
	if v189 != 0 {
		v187 = v191
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v197 = v191
	goto L50
L54:
	;
	v193 = v186 - int32(1)
	if v193 != 0 {
		v186 = v193
		v187 = v191
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v199 = v197
	goto L58
L57:
	;
	v199 = int32(311829)
	goto L58
L58:
	;
	goto L49
L59:
	;
	v209 = int32(2861)
	if v136 != 0 {
		v226 = v209
		v228 = v136
		goto L3
	} else {
		goto L60
	}
L60:
	;
	v240 = v209
	goto L2
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l1
	F_errmsg(m, int32(466275), v13+int32(48))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L39
	} else {
		goto L62
	}
L62:
	;
	v220 = int32(2872)
	if v136 == int32(0) {
		v240 = v220
		goto L2
	} else {
		goto L63
	}
L63:
	;
	v226 = v220
	v228 = v136
	goto L3
L64:
	;
	v240 = v226
	goto L2
L65:
	;
	goto L1
}
func F_assign_log_connections(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[1196])) = v4
	return
}
func F_assign_log_timezone(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[481])) = v4
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
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
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v141 int32
	_ = v141
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v300 int64
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	v9 = l8
	v11 = l10
	v13 = l12
	v14 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(4112)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[98]))) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))) = uint8(v14)
	F_XLogBeginInsert(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
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
	v33 = m.ExcPending
	if v33 != 0 {
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
	v36 = int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))) = uint8(v36)
	F_pg_qsort(m, l5, l6, int32(12), int32(186))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
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
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+608)) = v42
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+612)) = uint16(v44)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+614)) = uint16(v46)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+8)))
	v49 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+618)) = uint16(v49)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+616)) = uint8(v48)
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l5)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+16)) = uint16(v52)
	if l6 != v49 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v62 = v23 + int32(608)
	v75 = v49
	v76 = int32(1)
	goto L11
L9:
	;
	v141 = v49
	goto L10
L10:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+604)) = uint16(v141)
	F_XLogRegisterBufData(m, int32(0), v23+int32(604), int32(4))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L20
	}
L11:
	;
	v82 = l5 + v76*int32(12)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v83 != v84 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v141 = v113
	goto L10
L13:
	;
	v118 = int32(1)
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+10)))
	*(*uint16)(unsafe.Add(mBase, uint32(v23+int32(16)+v76<<(uint(v118)%32)))) = uint16(v121)
	v124 = v76 + v118
	if v124 != l6 {
		v62 = v112
		v75 = v113
		v76 = v124
		goto L11
	} else {
		goto L19
	}
L14:
	;
	v100 = v62 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v83
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+16)) = uint16(v102)
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+18)) = uint16(v104)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+8)))
	v107 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+22)) = uint16(v107)
	*(*uint8)(unsafe.Add(mBase, uint32(v62)+20)) = uint8(v106)
	v112 = v100
	v113 = v75 + v107
	goto L13
L15:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+4)))
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+4)))
	if v86 != v87 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+6)))
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v82)+6)))
	if v89 != v90 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+8)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+8)))
	if v92 != v93 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+10)))
	v97 = v95 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v62)+10)) = uint16(v97)
	v112 = v62
	v113 = v75
	goto L13
L19:
	;
	goto L12
L20:
	;
	F_XLogRegisterBufData(m, int32(0), v23+int32(608), v141*int32(12))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L6
L22:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))))
	v184 = v182 | int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))) = uint8(v184)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+602)) = uint16(v9)
	F_XLogRegisterBufData(m, int32(0), v23+int32(602), int32(2))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
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
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))))
	v202 = v200 | int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))) = uint8(v202)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+600)) = uint16(v11)
	F_XLogRegisterBufData(m, int32(0), v23+int32(600), int32(2))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
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
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))))
	v220 = v218 | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))) = uint8(v220)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+598)) = uint16(v13)
	F_XLogRegisterBufData(m, int32(0), v23+int32(598), int32(2))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
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
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	F_XLogRegisterBufData(m, int32(0), v23+int32(16), l6<<(uint(int32(1))%32))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v244 = *(*int32)(unsafe.Add(mBase, _consts[27]))
	if v244 < int32(2) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[98])))
	if v272 != 0 {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+118)))
	if v248 != int32(112) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L44
L44:
	;
	if base.B2i32(base.Ui32(v251) < base.Ui32(int32(12000))) == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v256 == int32(0) {
		goto L41
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))))
	v269 = v267 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))) = uint8(v269)
	goto L41
L48:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+119)))
	switch v260 - int32(109) {
	case 0, 5:
		goto L49
	default:
		goto L41
	}
L49:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+104)))
	if v263 != int32(1) {
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))))
	v275 = v273 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))) = uint8(v275)
	goto L53
L52:
	;
	goto L53
L53:
	;
	if l3 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))))
	v279 = v277 | int32(4)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+uint32(_consts[99]))) = uint8(v279)
	goto L56
L55:
	;
	goto L56
L56:
	;
	F_XLogRegisterData(m, v23+int32(4106), int32(2))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v272 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_XLogRegisterData(m, v23+int32(4108), int32(4))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if base.Ui32(l4) < base.Ui32(int32(3)) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L60
L62:
	;
	v300 = F_XLogInsert(m, int32(9), (l4<<(uint(int32(4))%32)+int32(16))&int32(240))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L70
	}
L65:
	;
	if l1 < int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v319))) = base.I64_rotr(v300, int64(32))
	m.G0 = v23 + int32(4112)
	return
L67:
	;
	v305 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v305+(l1^int32(-1))<<(uint(int32(2))%32))))
	v319 = v311
	goto L66
L68:
	;
	goto L69
L69:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v319 = v313 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L66
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = l4
	F_errmsg_internal(m, int32(479820), v23)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(492920), int32(2166), int32(340295))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	v4 = l3
	v7 = m.G0
	v9 = v7 - int32(112)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[305])))
	if v12 != int32(1) {
		v64 = int32(14)
		v71 = *(*int32)(unsafe.Add(mBase, _consts[215]))
		if v71 == int32(15) {
			v83 = int32(0)
			v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
			if v87 != int32(2) {
				v98 = v83
			} else {
				v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
				if v91 != 0 {
					v98 = v83
				} else {
					v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
					v98 = int32(0) | base.B2i32(v95 <= v64)
				}
			}
		} else {
			if v71 <= v64 {
				v98 = int32(1)
			} else {
				v83 = int32(0)
				v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
				if v87 != int32(2) {
					v98 = v83
				} else {
					v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
					if v91 != 0 {
						v98 = v83
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
						v98 = int32(0) | base.B2i32(v95 <= v64)
					}
				}
			}
		}
		if v98 == int32(0) {
			v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
			if v133 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
				v141 = int32(40)
				v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
					v147 = v144
					v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
					v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
					*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
					v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return
					} else {
						v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
						if v161 == int32(0) {
							*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
						} else {
						}
						m.G0 = v9 + int32(112)
						return
					}
				}
			} else {
				v147 = v133
				v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
				v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
				*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
				v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
				mBase = m.M
				v160 = m.ExcPending
				if v160 != 0 {
					return
				} else {
					v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
					if v161 == int32(0) {
						*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
					} else {
					}
					m.G0 = v9 + int32(112)
					return
				}
			}
		} else {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			F_GetRelationPath(m, v9+int32(40), v104, v105, v106, int32(-1), l1)
			mBase = m.M
			v109 = m.ExcPending
			if v109 != 0 {
				return
			} else {
				v112 = F_errstart(m, int32(14), int32(0))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return
				} else {
					if v112 == int32(0) {
						v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
						if v133 == int32(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
							v141 = int32(40)
							v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
								v147 = v144
								v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
								v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
								*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
								v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return
								} else {
									v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
									if v161 == int32(0) {
										*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
									} else {
									}
									m.G0 = v9 + int32(112)
									return
								}
							}
						} else {
							v147 = v133
							v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
							v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
							*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
							v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return
							} else {
								v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
								if v161 == int32(0) {
									*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
								} else {
								}
								m.G0 = v9 + int32(112)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(40)
						if v4 != 0 {
							v122 = int32(435637)
						} else {
							v122 = int32(69450)
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
							F_errfinish(m, int32(490815), v128, int32(404766))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return
							} else {
								v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
								if v133 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
									v141 = int32(40)
									v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
										v147 = v144
										v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
										v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
										*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
										v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return
										} else {
											v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
											if v161 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
											} else {
											}
											m.G0 = v9 + int32(112)
											return
										}
									}
								} else {
									v147 = v133
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
									v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
									*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
									v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
										if v161 == int32(0) {
											*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
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
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_GetRelationPath(m, v9+int32(40), v17, v18, v19, int32(-1), l1)
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
					*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v9 + int32(40)
					if v4 != 0 {
						v33 = int32(435637)
					} else {
						v33 = int32(69450)
					}
					F_errmsg_internal(m, v33, v9+int32(16))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						if v4 != 0 {
							v41 = int32(93)
						} else {
							v41 = int32(96)
						}
						F_errfinish(m, int32(490815), v41, int32(404766))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							v48 = int32(*(*uint8)(unsafe.Add(mBase, _consts[306])))
							if v48 != 0 {
								v49 = int32(19)
							} else {
								v49 = int32(23)
							}
							v51 = F_errstart(m, v49, int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if v51 == int32(0) {
									v64 = int32(14)
									v71 = *(*int32)(unsafe.Add(mBase, _consts[215]))
									if v71 == int32(15) {
										v83 = int32(0)
										v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
										if v87 != int32(2) {
											v98 = v83
										} else {
											v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
											if v91 != 0 {
												v98 = v83
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
												v98 = int32(0) | base.B2i32(v95 <= v64)
											}
										}
									} else {
										if v71 <= v64 {
											v98 = int32(1)
										} else {
											v83 = int32(0)
											v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
											if v87 != int32(2) {
												v98 = v83
											} else {
												v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
												if v91 != 0 {
													v98 = v83
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
													v98 = int32(0) | base.B2i32(v95 <= v64)
												}
											}
										}
									}
									if v98 == int32(0) {
										v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
										if v133 == int32(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
											v141 = int32(40)
											v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
												v147 = v144
												v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
												v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return
												} else {
													v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v161 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											v147 = v133
											v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
											v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
											*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
											*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
											v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return
											} else {
												v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
												if v161 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
												} else {
												}
												m.G0 = v9 + int32(112)
												return
											}
										}
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_GetRelationPath(m, v9+int32(40), v104, v105, v106, int32(-1), l1)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v112 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return
											} else {
												if v112 == int32(0) {
													v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
													if v133 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
														v141 = int32(40)
														v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
															v147 = v144
															v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
															v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return
															} else {
																v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v161 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v147 = v133
														v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
														v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return
														} else {
															v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v161 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(40)
													if v4 != 0 {
														v122 = int32(435637)
													} else {
														v122 = int32(69450)
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
														F_errfinish(m, int32(490815), v128, int32(404766))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
															if v133 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
																v141 = int32(40)
																v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
																	v147 = v144
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
																	v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return
																	} else {
																		v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v161 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v147 = v133
																v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
																v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return
																} else {
																	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v161 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
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
									F_errmsg_internal(m, int32(170046), int32(0))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return
									} else {
										F_errfinish(m, int32(490815), int32(120), int32(404786))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											v64 = int32(14)
											v71 = *(*int32)(unsafe.Add(mBase, _consts[215]))
											if v71 == int32(15) {
												v83 = int32(0)
												v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
												if v87 != int32(2) {
													v98 = v83
												} else {
													v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
													if v91 != 0 {
														v98 = v83
													} else {
														v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
														v98 = int32(0) | base.B2i32(v95 <= v64)
													}
												}
											} else {
												if v71 <= v64 {
													v98 = int32(1)
												} else {
													v83 = int32(0)
													v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
													if v87 != int32(2) {
														v98 = v83
													} else {
														v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
														if v91 != 0 {
															v98 = v83
														} else {
															v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
															v98 = int32(0) | base.B2i32(v95 <= v64)
														}
													}
												}
											}
											if v98 == int32(0) {
												v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
												if v133 == int32(0) {
													*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
													v141 = int32(40)
													v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
														v147 = v144
														v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
														v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return
														} else {
															v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v161 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													v147 = v133
													v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
													v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
													*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
													*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
													v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return
													} else {
														v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
														if v161 == int32(0) {
															*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
														} else {
														}
														m.G0 = v9 + int32(112)
														return
													}
												}
											} else {
												v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												F_GetRelationPath(m, v9+int32(40), v104, v105, v106, int32(-1), l1)
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v112 = F_errstart(m, int32(14), int32(0))
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return
													} else {
														if v112 == int32(0) {
															v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
															if v133 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
																v141 = int32(40)
																v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
																	v147 = v144
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
																	v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return
																	} else {
																		v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v161 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v147 = v133
																v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
																v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return
																} else {
																	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v161 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
																	} else {
																	}
																	m.G0 = v9 + int32(112)
																	return
																}
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(40)
															if v4 != 0 {
																v122 = int32(435637)
															} else {
																v122 = int32(69450)
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
																F_errfinish(m, int32(490815), v128, int32(404766))
																mBase = m.M
																v131 = m.ExcPending
																if v131 != 0 {
																	return
																} else {
																	v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
																	if v133 == int32(0) {
																		*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
																		v141 = int32(40)
																		v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
																		mBase = m.M
																		v145 = m.ExcPending
																		if v145 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
																			v147 = v144
																			v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
																			v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																			*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																			v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
																			mBase = m.M
																			v160 = m.ExcPending
																			if v160 != 0 {
																				return
																			} else {
																				v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																				if v161 == int32(0) {
																					*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
																				} else {
																				}
																				m.G0 = v9 + int32(112)
																				return
																			}
																		}
																	} else {
																		v147 = v133
																		v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
																		v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																		*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																		v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
																		mBase = m.M
																		v160 = m.ExcPending
																		if v160 != 0 {
																			return
																		} else {
																			v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																			if v161 == int32(0) {
																				*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
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
					v48 = int32(*(*uint8)(unsafe.Add(mBase, _consts[306])))
					if v48 != 0 {
						v49 = int32(19)
					} else {
						v49 = int32(23)
					}
					v51 = F_errstart(m, v49, int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						if v51 == int32(0) {
							v64 = int32(14)
							v71 = *(*int32)(unsafe.Add(mBase, _consts[215]))
							if v71 == int32(15) {
								v83 = int32(0)
								v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
								if v87 != int32(2) {
									v98 = v83
								} else {
									v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
									if v91 != 0 {
										v98 = v83
									} else {
										v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
										v98 = int32(0) | base.B2i32(v95 <= v64)
									}
								}
							} else {
								if v71 <= v64 {
									v98 = int32(1)
								} else {
									v83 = int32(0)
									v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
									if v87 != int32(2) {
										v98 = v83
									} else {
										v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
										if v91 != 0 {
											v98 = v83
										} else {
											v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
											v98 = int32(0) | base.B2i32(v95 <= v64)
										}
									}
								}
							}
							if v98 == int32(0) {
								v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
								if v133 == int32(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
									v141 = int32(40)
									v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
										v147 = v144
										v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
										v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
										*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
										*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
										v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return
										} else {
											v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
											if v161 == int32(0) {
												*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
											} else {
											}
											m.G0 = v9 + int32(112)
											return
										}
									}
								} else {
									v147 = v133
									v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
									v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
									*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
									v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return
									} else {
										v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
										if v161 == int32(0) {
											*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
										} else {
										}
										m.G0 = v9 + int32(112)
										return
									}
								}
							} else {
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								F_GetRelationPath(m, v9+int32(40), v104, v105, v106, int32(-1), l1)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v112 = F_errstart(m, int32(14), int32(0))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return
									} else {
										if v112 == int32(0) {
											v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
											if v133 == int32(0) {
												*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
												v141 = int32(40)
												v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
													v147 = v144
													v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
													v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
													*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
													*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
													v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
													mBase = m.M
													v160 = m.ExcPending
													if v160 != 0 {
														return
													} else {
														v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
														if v161 == int32(0) {
															*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
														} else {
														}
														m.G0 = v9 + int32(112)
														return
													}
												}
											} else {
												v147 = v133
												v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
												v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return
												} else {
													v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v161 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(40)
											if v4 != 0 {
												v122 = int32(435637)
											} else {
												v122 = int32(69450)
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
												F_errfinish(m, int32(490815), v128, int32(404766))
												mBase = m.M
												v131 = m.ExcPending
												if v131 != 0 {
													return
												} else {
													v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
													if v133 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
														v141 = int32(40)
														v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
															v147 = v144
															v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
															v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return
															} else {
																v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v161 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v147 = v133
														v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
														v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return
														} else {
															v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v161 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
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
							F_errmsg_internal(m, int32(170046), int32(0))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								F_errfinish(m, int32(490815), int32(120), int32(404786))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									v64 = int32(14)
									v71 = *(*int32)(unsafe.Add(mBase, _consts[215]))
									if v71 == int32(15) {
										v83 = int32(0)
										v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
										if v87 != int32(2) {
											v98 = v83
										} else {
											v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
											if v91 != 0 {
												v98 = v83
											} else {
												v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
												v98 = int32(0) | base.B2i32(v95 <= v64)
											}
										}
									} else {
										if v71 <= v64 {
											v98 = int32(1)
										} else {
											v83 = int32(0)
											v87 = *(*int32)(unsafe.Add(mBase, _consts[216]))
											if v87 != int32(2) {
												v98 = v83
											} else {
												v91 = int32(*(*uint8)(unsafe.Add(mBase, _consts[217])))
												if v91 != 0 {
													v98 = v83
												} else {
													v95 = *(*int32)(unsafe.Add(mBase, _consts[218]))
													v98 = int32(0) | base.B2i32(v95 <= v64)
												}
											}
										}
									}
									if v98 == int32(0) {
										v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
										if v133 == int32(0) {
											*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
											v141 = int32(40)
											v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
											mBase = m.M
											v145 = m.ExcPending
											if v145 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
												v147 = v144
												v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
												v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
												*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
												*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
												*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
												v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
												mBase = m.M
												v160 = m.ExcPending
												if v160 != 0 {
													return
												} else {
													v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
													if v161 == int32(0) {
														*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
													} else {
													}
													m.G0 = v9 + int32(112)
													return
												}
											}
										} else {
											v147 = v133
											v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
											v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
											*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
											*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
											v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return
											} else {
												v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
												if v161 == int32(0) {
													*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
												} else {
												}
												m.G0 = v9 + int32(112)
												return
											}
										}
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										F_GetRelationPath(m, v9+int32(40), v104, v105, v106, int32(-1), l1)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v112 = F_errstart(m, int32(14), int32(0))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
												return
											} else {
												if v112 == int32(0) {
													v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
													if v133 == int32(0) {
														*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
														v141 = int32(40)
														v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
															v147 = v144
															v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
															v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
															*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
															*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
															*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
															v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return
															} else {
																v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																if v161 == int32(0) {
																	*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
																} else {
																}
																m.G0 = v9 + int32(112)
																return
															}
														}
													} else {
														v147 = v133
														v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
														v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
														*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
														*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
														v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
														mBase = m.M
														v160 = m.ExcPending
														if v160 != 0 {
															return
														} else {
															v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
															if v161 == int32(0) {
																*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
															} else {
															}
															m.G0 = v9 + int32(112)
															return
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9))) = l2
													*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v9 + int32(40)
													if v4 != 0 {
														v122 = int32(435637)
													} else {
														v122 = int32(69450)
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
														F_errfinish(m, int32(490815), v128, int32(404766))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return
														} else {
															v133 = *(*int32)(unsafe.Add(mBase, _consts[290]))
															if v133 == int32(0) {
																*(*int64)(unsafe.Add(mBase, uint32(v9)+56)) = int64(103079215124)
																v141 = int32(40)
																v144 = F_hash_create(m, int32(391829), int32(100), v9+v141, v141)
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _consts[290])) = v144
																	v147 = v144
																	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
																	v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																	*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																	v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return
																	} else {
																		v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																		if v161 == int32(0) {
																			*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
																		} else {
																		}
																		m.G0 = v9 + int32(112)
																		return
																	}
																}
															} else {
																v147 = v133
																v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v148
																v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
																*(*int64)(unsafe.Add(mBase, uint32(v9)+40)) = v150
																*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = l2
																*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = l1
																v159 = F_hash_search(m, v147, v9+int32(40), int32(1), v9+int32(39))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return
																} else {
																	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+39)))
																	if v161 == int32(0) {
																		*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)) = uint8(v4)
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
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
		v11 = *(*int32)(unsafe.Add(mBase, _consts[130]))
		if v11 <= int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[130])) = int32(1)
		} else {
		}
		v18 = *(*int32)(unsafe.Add(mBase, _consts[131]))
		if int32(0) < v18 {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[132]))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v22)+4)) = v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v22 + int32(32)
			if l4 != 0 {
				v35 = int32(9)
			} else {
				v35 = int32(1)
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v35)
			v37 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v37
			v39 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v39)
			v43 = F_XLogInsert(m, v37, int32(176))
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
				F_errmsg_internal(m, int32(134508), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errfinish(m, int32(489916), int32(320), int32(315148))
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
