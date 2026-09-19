package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockDequeueSelf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(536870912)
	v14 = base.AtomicRmwOr32(m, l0, int32(4), v12)
	if v14&v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = v14
	goto L4
L2:
	;
	goto L3
L3:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+74)))
	if v94 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(_a_F_LWLockDequeueSelf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(_a_F_LWLockDequeueSelf_1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	if v18&int32(536870912) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	goto L9
L7:
	;
	goto L8
L8:
	;
	v60 = int32(_a_F_LWLockDequeueSelf_2)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[1]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8))+8))
	if v63 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v10+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v47&int32(536870912) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v80 = int32(536870912)
	v82 = base.AtomicRmwOr32(m, l0, int32(4), v80)
	if v82&v80 != 0 {
		v18 = v82
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[1])) = v78
	goto L15
L17:
	;
	if int32(999) < v61 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v61 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v68 = int32(900)
	if v68 <= v61 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v71 = v68
	goto L23
L22:
	;
	v71 = v61
	goto L23
L23:
	;
	v78 = v71 + int32(100)
	goto L16
L24:
	;
	v78 = v61 - int32(1)
	goto L16
L25:
	;
	goto L5
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[2]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[3]))
	v104 = v99 + v101*int32(640)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+76))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v104)+80))
	if v106 == int32(-1) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L28
L28:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v132 != int32(-1) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	if v105 == int32(-1) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v105
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104)+80))
	v115 = v110
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99+v106*int32(640))+76)) = v105
	v115 = v106
	goto L29
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v104)+76)) = int64(0)
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v115
	goto L33
L35:
	;
	goto L36
L36:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[2]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	*(*int32)(unsafe.Add(mBase, uint32(v121+v105*int32(640))+80)) = v115
	goto L33
L37:
	;
	v143 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-536870913))
	if v94 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) <= v135 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v140 = base.AtomicRmwAnd32(m, l0, int32(4), int32(2147483647))
	goto L37
L40:
	;
	m.G0 = v10 + int32(32)
	return
L41:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147)+74)) = uint8(v148)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v152 = base.AtomicRmwOr32(m, l0, int32(4), int32(1073741824))
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v157 = v154
	v158 = int32(0)
	goto L44
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	F_PGSemaphoreLock(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L11
	} else {
		goto L46
	}
L45:
	;
	if v158 <= int32(0) {
		goto L40
	} else {
		goto L48
	}
L46:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+74)))
	if v170 != 0 {
		v157 = v169
		v158 = v158 + int32(1)
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v173 = v158
	goto L49
L49:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	F_PGSemaphoreUnlock(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L11
	} else {
		goto L51
	}
L50:
	;
	goto L40
L51:
	;
	v185 = int32(1)
	if base.Ui32(v185) < base.Ui32(v173) {
		v173 = v173 - v185
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
}
func F_LWLockInitialize(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = l1
	*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-1)
	return
}
func F_LWLockUpdateVar(m *base.Module, l0 int32, l1 int32, l2 int64) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var __phi133 int32
	_ = __phi133
	var v135 int32
	_ = v135
	var __phi135 int32
	_ = __phi135
	var v136 int32
	_ = v136
	var __phi136 int32
	_ = __phi136
	var v138 int32
	_ = v138
	var __phi138 int32
	_ = __phi138
	var v139 int32
	_ = v139
	var __phi139 int32
	_ = __phi139
	var v142 int32
	_ = v142
	var __phi142 int32
	_ = __phi142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = base.AtomicRmwXchg64(m, l1, int32(0), l2)
	v18 = int32(536870912)
	v20 = base.AtomicRmwOr32(m, l0, int32(4), v18)
	if v20&v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v20
	goto L4
L2:
	;
	goto L3
L3:
	;
	v114 = int32(-1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v115 == v114 {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = int32(_a_F_LWLockUpdateVar_0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = int32(_a_F_LWLockUpdateVar_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
	if v24&int32(536870912) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	goto L9
L7:
	;
	goto L8
L8:
	;
	v78 = int32(_a_F_LWLockUpdateVar_2)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[0]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v14+int32(8))+8))
	if v81 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v14+int32(8))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	return
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v61&int32(536870912) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v98 = int32(536870912)
	v100 = base.AtomicRmwOr32(m, l0, int32(4), v98)
	if v100&v98 != 0 {
		v24 = v100
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[0])) = v96
	goto L15
L17:
	;
	if int32(999) < v79 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v79 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v86 = int32(900)
	if v86 <= v79 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v89 = v86
	goto L23
L22:
	;
	v89 = v79
	goto L23
L23:
	;
	v96 = v89 + int32(100)
	goto L16
L24:
	;
	v96 = v79 - int32(1)
	goto L16
L25:
	;
	goto L5
L26:
	;
	m.G0 = v14 + int32(32)
	return
L27:
	;
	v120 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-536870913))
	goto L26
L28:
	;
	goto L29
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v125 = v115 * int32(640)
	v126 = v123 + v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+75)))
	if v127 != int32(2) {
		v202 = v114
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v214 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-536870913))
	if v202 == int32(-1) {
		goto L26
	} else {
		goto L48
	}
L31:
	;
	__phi133 = v114
	__phi135 = v125 + v123
	__phi136 = v115
	__phi138 = int32(-1)
	__phi139 = v126
	__phi142 = v123
	v133 = __phi133
	v135 = __phi135
	v136 = __phi136
	v138 = __phi138
	v139 = __phi139
	v142 = __phi142
	goto L32
L32:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v139)+76))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v135)+76))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139)+80))
	if v145 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v202 = v185
	goto L30
L34:
	;
	if v143 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v143
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v139)+80))
	v154 = v149
	goto L34
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142+v145*int32(640))+76)) = v143
	v154 = v145
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v139)+76)) = int64(0)
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v172 = v169 + v136*int32(640)
	if v138 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v154
	goto L38
L40:
	;
	goto L41
L41:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v160+v143*int32(640))+80)) = v154
	goto L38
L42:
	;
	v186 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v172)+76)) = v186
	v188 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v139)+74)) = uint8(v188)
	if v144 == v186 {
		v202 = v185
		goto L30
	} else {
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+80)) = int32(-1)
	v185 = v136
	goto L42
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+80)) = v138
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)))
	*(*int32)(unsafe.Add(mBase, uint32(v180+v138*int32(640))+76)) = v136
	v185 = v133
	goto L42
L46:
	;
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v197 = v194 + v144*int32(640)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+75)))
	if v198 == int32(2) {
		__phi133 = v185
		__phi135 = v197
		__phi136 = v144
		__phi138 = v136
		__phi139 = v197
		__phi142 = v194
		v133 = __phi133
		v135 = __phi135
		v136 = __phi136
		v138 = __phi138
		v139 = __phi139
		v142 = __phi142
		goto L32
	} else {
		goto L47
	}
L47:
	;
	goto L33
L48:
	;
	v218 = v202
	goto L49
L49:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v233 = v230 + v218*int32(640)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+76))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233)+80))
	if v235 != int32(-1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L26
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v230+v235*int32(640))+76)) = v234
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v234 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	*(*int32)(unsafe.Add(mBase, uint32(v246+v234*int32(640))+80)) = v235
	goto L56
L55:
	;
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v233)+76)) = int64(0)
	v253 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v233)+74)) = uint8(v253)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	F_PGSemaphoreUnlock(m, v255)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	if v234 != int32(-1) {
		v218 = v234
		goto L49
	} else {
		goto L58
	}
L58:
	;
	goto L50
}
