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
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v12 | v13
	if v12&v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = v12
	goto L4
L2:
	;
	goto L3
L3:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+74)))
	if v96 == int32(1) {
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
	if v19&int32(536870912) != 0 {
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
	v61 = int32(_a_F_LWLockDequeueSelf_2)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[1]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8))+8))
	if v64 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v10+int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
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
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v48&int32(536870912) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v82 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v81 | v82
	if v81&v82 != 0 {
		v19 = v81
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[1])) = v79
	goto L15
L17:
	;
	if int32(999) < v62 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v62 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v69 = int32(900)
	if v69 <= v62 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v72 = v69
	goto L23
L22:
	;
	v72 = v62
	goto L23
L23:
	;
	v79 = v72 + int32(100)
	goto L16
L24:
	;
	v79 = v62 - int32(1)
	goto L16
L25:
	;
	goto L5
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[2]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[3]))
	v106 = v101 + v103*int32(640)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+76))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)+80))
	if v108 == int32(-1) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L28
L28:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v134 != int32(-1) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	if v107 == int32(-1) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v107
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v106)+80))
	v117 = v112
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101+v108*int32(640))+76)) = v107
	v117 = v108
	goto L29
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v106)+76)) = int64(0)
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v117
	goto L33
L35:
	;
	goto L36
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[2]))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v123+v107*int32(640))+80)) = v117
	goto L33
L37:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v144 & int32(-536870913)
	if v96 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) <= v137 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v140 & int32(2147483647)
	goto L37
L40:
	;
	m.G0 = v10 + int32(32)
	return
L41:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v151)+74)) = uint8(v152)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154 | int32(1073741824)
	v163 = int32(0)
	goto L44
L44:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDequeueSelf[0]))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172)+74)))
	if v173 != 0 {
		v163 = v163 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	if v163 <= int32(0) {
		goto L40
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v176 = v163
	goto L48
L48:
	;
	v186 = int32(1)
	if base.Ui32(v186) < base.Ui32(v176) {
		v176 = v176 - v186
		goto L48
	} else {
		goto L50
	}
L49:
	;
	goto L40
L50:
	;
	goto L49
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v99 int32
	_ = v99
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var __phi135 int32
	_ = __phi135
	var v137 int32
	_ = v137
	var __phi137 int32
	_ = __phi137
	var v138 int32
	_ = v138
	var __phi138 int32
	_ = __phi138
	var v140 int32
	_ = v140
	var __phi140 int32
	_ = __phi140
	var v141 int32
	_ = v141
	var __phi141 int32
	_ = __phi141
	var v144 int32
	_ = v144
	var __phi144 int32
	_ = __phi144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v256 int32
	_ = v256
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = l2
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 | v18
	if v17&v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = v17
	goto L4
L2:
	;
	goto L3
L3:
	;
	v115 = int32(-1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v116 == v115 {
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
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v99 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v98 | v99
	if v98&v99 != 0 {
		v24 = v98
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
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119 & int32(-536870913)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v127 = v116 * int32(640)
	v128 = v125 + v127
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+75)))
	if v129 != int32(2) {
		v204 = v115
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v214 & int32(-536870913)
	if v204 == int32(-1) {
		goto L26
	} else {
		goto L48
	}
L31:
	;
	__phi135 = v115
	__phi137 = v127 + v125
	__phi138 = v116
	__phi140 = int32(-1)
	__phi141 = v128
	__phi144 = v125
	v135 = __phi135
	v137 = __phi137
	v138 = __phi138
	v140 = __phi140
	v141 = __phi141
	v144 = __phi144
	goto L32
L32:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141)+76))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v137)+76))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)+80))
	if v147 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v204 = v187
	goto L30
L34:
	;
	if v145 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v145
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v141)+80))
	v156 = v151
	goto L34
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144+v147*int32(640))+76)) = v145
	v156 = v147
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v141)+76)) = int64(0)
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v174 = v171 + v138*int32(640)
	if v140 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v156
	goto L38
L40:
	;
	goto L41
L41:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	*(*int32)(unsafe.Add(mBase, uint32(v162+v145*int32(640))+80)) = v156
	goto L38
L42:
	;
	v188 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v174)+76)) = v188
	v190 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+74)) = uint8(v190)
	if v146 == v188 {
		v204 = v187
		goto L30
	} else {
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+80)) = int32(-1)
	v187 = v138
	goto L42
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v174)+80)) = v140
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	*(*int32)(unsafe.Add(mBase, uint32(v182+v140*int32(640))+76)) = v138
	v187 = v135
	goto L42
L46:
	;
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v199 = v196 + v146*int32(640)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v199)+75)))
	if v200 == int32(2) {
		__phi135 = v187
		__phi137 = v199
		__phi138 = v146
		__phi140 = v138
		__phi141 = v199
		__phi144 = v196
		v135 = __phi135
		v137 = __phi137
		v138 = __phi138
		v140 = __phi140
		v141 = __phi141
		v144 = __phi144
		goto L32
	} else {
		goto L47
	}
L47:
	;
	goto L33
L48:
	;
	v221 = v204
	goto L49
L49:
	;
	v232 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v236 = v233 + v221*int32(640)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+76))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)+80))
	if v238 != int32(-1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L26
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v233+v238*int32(640))+76)) = v237
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v237 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockUpdateVar[1]))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	*(*int32)(unsafe.Add(mBase, uint32(v249+v237*int32(640))+80)) = v238
	goto L56
L55:
	;
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v236)+76)) = int64(0)
	v256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v236)+74)) = uint8(v256)
	if v237 != int32(-1) {
		v221 = v237
		goto L49
	} else {
		goto L57
	}
L57:
	;
	goto L50
}
