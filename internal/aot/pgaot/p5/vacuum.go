package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_vacuum_buffer_usage_limit(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = int32(0)
	v15 = base.B2i32(v8 == v9) | base.B2i32(base.Ui32(v8-int32(128)) < base.Ui32(int32(16777089)))
	if v15 == v9 {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_check_vacuum_buffer_usage_limit[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_vacuum_buffer_usage_limit[1])) = v19
		*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(72057594037928064)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_check_vacuum_buffer_usage_limit_0)
		v28 = F_format_elog_string(m, int32(_a_F_check_vacuum_buffer_usage_limit_1), v6)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_vacuum_buffer_usage_limit[2])) = v28
			m.G0 = v6 + int32(16)
			return v15
		}
	} else {
		m.G0 = v6 + int32(16)
		return v15
	}
}
func F_vacuumRedirectAndPlaceholder(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
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
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(1648)
	m.G0 = v21
	if l2 < v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v41) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(l2^int32(-1))<<(uint(int32(2))%32))))
	v40 = v32
	goto L1
L3:
	;
	goto L4
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[1]))
	v40 = v34 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v49 = int32(base.Ui32(v41+int32(_a_F_vacuumRedirectAndPlaceholder_0)) >> (uint(int32(2)) % 32))
	goto L7
L6:
	;
	v49 = int32(0)
	goto L7
L7:
	;
	v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[2]))
	if v52 < int32(2) {
		v72 = v4
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v74 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v74
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v74)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)) = uint8(v72)
	v79 = F_GlobalVisHorizonKindForRel(m, l1)
	mBase = m.M
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79<<(uint(int32(2))%32))+uint32(_c_F_vacuumRedirectAndPlaceholder[3])))
	goto L15
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+118)))
	if v56 != int32(112) {
		v72 = v4
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	goto L11
L11:
	;
	if base.Ui32(v60) < base.Ui32(int32(_a_F_vacuumRedirectAndPlaceholder_1)) {
		v72 = int32(1)
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v64 == v63 {
		v72 = v63
		goto L8
	} else {
		goto L13
	}
L13:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+119)))
	switch v68 - int32(109) {
	case 0, 5:
		goto L14
	default:
		v72 = v63
		goto L8
	}
L14:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+104)))
	v72 = v71
	goto L8
L15:
	;
	v83 = int32(_a_F_vacuumRedirectAndPlaceholder_2)
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[4])) = v85 + int32(1)
	v90 = v49 & int32(_a_F_vacuumRedirectAndPlaceholder_3)
	if v90 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v363 = int32(_a_F_vacuumRedirectAndPlaceholder_2)
	v365 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[4])) = v365 - int32(1)
	m.G0 = v21 + int32(1648)
	return
L17:
	;
	v93 = v40 + v50
	v97 = v49
	v99 = v90
	v106 = v4
	v109 = v4
	v110 = v4
	v111 = v4
	v113 = v4
	goto L18
L18:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)))
	v115 = int32(0)
	if base.B2i32(v114 == v115)&(v109&int32(1)) == v115 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v217 = v210 & int32(_a_F_vacuumRedirectAndPlaceholder_3)
	if v217 != 0 {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v40+int32(20)+v99<<(uint(int32(2))%32))))
	v128 = v40 + v125&int32(_a_F_vacuumRedirectAndPlaceholder_4)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v129&int32(3) != int32(1) {
		v188 = v129
		v189 = v110
		v190 = v111
		v191 = v113
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v210 = v106
	v215 = v113
	goto L22
L22:
	;
	goto L19
L23:
	;
	v192 = int32(3)
	v196 = base.B2i32(v188&v192 != v192) | v109
	if v196&int32(1) != 0 {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if v134 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v135 = F_GlobalVisTestIsRemovableXid(m, v82, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v140 = v129
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v140 | int32(3)
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)))
	v145 = int32(1)
	v146 = v144 - v145
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+2)) = uint16(v146)
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
	v150 = v148 + v145
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)) = uint16(v150)
	if v110 != 0 {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	return
L29:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	if v135 == int32(0) {
		v188 = v137
		v189 = v110
		v190 = v111
		v191 = v113
		goto L23
	} else {
		goto L30
	}
L30:
	;
	v140 = v137
	goto L27
L31:
	;
	v170 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v128)+10)) = uint16(v170)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+6)) = int32(-1)
	v174 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(832)+v111&int32(_a_F_vacuumRedirectAndPlaceholder_3)<<(uint(v174)%32)))) = uint16(v99)
	v184 = v111 + v174
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)) = uint16(v184)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v188 = v186
	v189 = v169
	v190 = v184
	v191 = v174
	goto L23
L32:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v152))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v110)) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v167
	v169 = v167
	goto L31
L35:
	;
	if v164 == int32(0) {
		v169 = v110
		goto L31
	} else {
		goto L39
	}
L36:
	;
	v164 = base.B2i32(base.Ui32(v110) < base.Ui32(v152))
	goto L35
L37:
	;
	goto L38
L38:
	;
	v164 = int32(base.Ui32(v110-v152) >> (uint(int32(31)) % 32))
	goto L35
L39:
	;
	goto L34
L40:
	;
	v199 = v106
	goto L42
L41:
	;
	v199 = v99
	goto L42
L42:
	;
	v200 = int32(1)
	v203 = v97 - v200
	if v203&int32(_a_F_vacuumRedirectAndPlaceholder_3) != 0 {
		v97 = v203
		v99 = v99 - v200
		v106 = v199
		v109 = v196
		v110 = v189
		v111 = v190
		v113 = v191
		goto L18
	} else {
		goto L43
	}
L43:
	;
	v210 = v199
	v215 = v191
	goto L22
L44:
	;
	F_MarkBufferDirty(m, l2)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L28
	} else {
		goto L56
	}
L45:
	;
	if base.Ui32(v217) <= base.Ui32(v49&int32(_a_F_vacuumRedirectAndPlaceholder_3)) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v286 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v286)
	if v215 == v286 {
		goto L16
	} else {
		goto L55
	}
L48:
	;
	v224 = v210
	goto L51
L49:
	;
	goto L50
L50:
	;
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
	v276 = v49 - v210 + int32(1)
	v277 = v273 - v276
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)) = uint16(v277)
	F_PageIndexMultiDelete(m, v40, v21+int32(16), v276&int32(_a_F_vacuumRedirectAndPlaceholder_3))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L28
	} else {
		goto L54
	}
L51:
	;
	v241 = int32(_a_F_vacuumRedirectAndPlaceholder_3)
	v244 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v21+int32(16)+(v224&v241-v217)<<(uint(v244)%32)))) = uint16(v224)
	v249 = v224 + v244
	if base.Ui32(v249&v241) <= base.Ui32(v49&v241) {
		v224 = v249
		goto L51
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	goto L52
L54:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)) = uint16(v210)
	goto L44
L55:
	;
	goto L44
L56:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+118)))
	if v311 != int32(112) {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	v315 = *(*int32)(unsafe.Add(mBase, _c_F_vacuumRedirectAndPlaceholder[2]))
	if v315 <= int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v318 != 0 {
		goto L16
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L28
	} else {
		goto L63
	}
L61:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v319 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	F_XLogRegisterData(m, v21+int32(4), int32(10))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L28
	} else {
		goto L64
	}
L64:
	;
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	F_XLogRegisterData(m, v21+int32(832), v329<<(uint(int32(1))%32))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L28
	} else {
		goto L65
	}
L65:
	;
	F_XLogRegisterBuffer(m, int32(0), l2, int32(8))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L28
	} else {
		goto L66
	}
L66:
	;
	v340 = F_XLogInsert(m, int32(16), int32(128))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L28
	} else {
		goto L67
	}
L67:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v40))) = base.I64_rotr(v340, int64(32))
	goto L16
}
