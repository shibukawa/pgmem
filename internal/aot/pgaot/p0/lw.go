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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
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
	v95 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+74)))
	if v96 == int32(1) {
		goto L26
	} else {
		goto L27
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(330303)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(518076)
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
	v61 = int32(4155244)
	v62 = *(*int32)(unsafe.Add(mBase, _consts[571]))
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
	*(*int32)(unsafe.Add(mBase, _consts[571])) = v79
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
	v100 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v103 = *(*int32)(unsafe.Add(mBase, _consts[246]))
	v106 = v101 + v103*int32(640)
	v108 = v106 + int32(76)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106)+80))
	if v110 == int32(-1) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L28
L28:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v136 != int32(-1) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	if v109 == int32(-1) {
		goto L34
	} else {
		goto L35
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v109
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v119 = v114
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101+v110*int32(640))+76)) = v109
	v119 = v110
	goto L29
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = int64(0)
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v119
	goto L33
L35:
	;
	goto L36
L36:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v125+v109*int32(640))+80)) = v119
	goto L33
L37:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v146 & int32(-536870913)
	if v96 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) <= v139 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v142 & int32(2147483647)
	goto L37
L40:
	;
	m.G0 = v10 + int32(32)
	return
L41:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	v154 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+74)) = uint8(v154)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v156 | int32(1073741824)
	v166 = int32(0)
	goto L44
L44:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _consts[88]))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+74)))
	if v175 != 0 {
		v166 = v166 + int32(1)
		goto L44
	} else {
		goto L46
	}
L45:
	;
	if v166 <= int32(0) {
		goto L40
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v178 = v166
	goto L48
L48:
	;
	v188 = int32(1)
	if base.Ui32(v188) < base.Ui32(v178) {
		v178 = v178 - v188
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var __phi155 int32
	_ = __phi155
	var v158 int32
	_ = v158
	var __phi158 int32
	_ = __phi158
	var v161 int32
	_ = v161
	var __phi161 int32
	_ = __phi161
	var v162 int32
	_ = v162
	var __phi162 int32
	_ = __phi162
	var v163 int32
	_ = v163
	var __phi163 int32
	_ = __phi163
	var v164 int32
	_ = v164
	var __phi164 int32
	_ = __phi164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = l2
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v20 | v21
	if v20&v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = v20
	goto L4
L2:
	;
	goto L3
L3:
	;
	v130 = int32(-1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v131 == v130 {
		goto L27
	} else {
		goto L28
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(330303)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(518076)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
	if v27&int32(536870912) != 0 {
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
	v90 = int32(4155244)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[571]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(8))+8))
	if v93 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L9:
	;
	F_perform_spin_delay(m, v17+int32(8))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
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
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v70&int32(536870912) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v111 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v110 | v111
	if v110&v111 != 0 {
		v27 = v110
		goto L4
	} else {
		goto L25
	}
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[571])) = v108
	goto L15
L17:
	;
	if int32(999) < v91 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v91 < int32(11) {
		goto L15
	} else {
		goto L24
	}
L20:
	;
	v98 = int32(900)
	if v98 <= v91 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v101 = v98
	goto L23
L22:
	;
	v101 = v91
	goto L23
L23:
	;
	v108 = v101 + int32(100)
	goto L16
L24:
	;
	v108 = v91 - int32(1)
	goto L16
L25:
	;
	goto L5
L26:
	;
	m.G0 = v17 + int32(32)
	return
L27:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v134 & int32(-536870913)
	goto L26
L28:
	;
	goto L29
L29:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v142 = v131 * int32(640)
	v143 = v140 + v142
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+75)))
	if v144 != int32(2) {
		v234 = v130
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v247 & int32(-536870913)
	if v234 == int32(-1) {
		goto L26
	} else {
		goto L48
	}
L31:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	__phi155 = v130
	__phi158 = v131
	__phi161 = v149 + v142 + int32(76)
	__phi162 = int32(-1)
	__phi163 = v140
	__phi164 = v143
	v155 = __phi155
	v158 = __phi158
	v161 = __phi161
	v162 = __phi162
	v163 = __phi163
	v164 = __phi164
	goto L32
L32:
	;
	v169 = v158 * int32(640)
	v170 = v163 + v169
	v172 = v170 + int32(76)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v170)+80))
	if v175 == int32(-1) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v234 = v215
	goto L30
L34:
	;
	if v173 == int32(-1) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v173
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	v184 = v179
	goto L34
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163+v175*int32(640))+76)) = v173
	v184 = v175
	goto L34
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v172))) = int64(0)
	v198 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v202 = v199 + v169 + int32(76)
	if v162 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v184
	goto L38
L40:
	;
	goto L41
L41:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	*(*int32)(unsafe.Add(mBase, uint32(v190+v173*int32(640))+80)) = v184
	goto L38
L42:
	;
	v216 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = v216
	v218 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+74)) = uint8(v218)
	if v174 == v216 {
		v234 = v215
		goto L30
	} else {
		goto L46
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = int32(-1)
	v215 = v158
	goto L42
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v162
	v209 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	*(*int32)(unsafe.Add(mBase, uint32(v210+v162*int32(640))+76)) = v158
	v215 = v155
	goto L42
L46:
	;
	v223 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v227 = v224 + v174*int32(640)
	v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+75)))
	if v230 == int32(2) {
		__phi155 = v215
		__phi158 = v174
		__phi161 = v227 + int32(76)
		__phi162 = v158
		__phi163 = v224
		__phi164 = v227
		v155 = __phi155
		v158 = __phi158
		v161 = __phi161
		v162 = __phi162
		v163 = __phi163
		v164 = __phi164
		goto L32
	} else {
		goto L47
	}
L47:
	;
	goto L33
L48:
	;
	v254 = v234
	goto L49
L49:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v272 = v269 + v254*int32(640)
	v274 = v272 + int32(76)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v272)+80))
	if v276 != int32(-1) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L26
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v269+v276*int32(640))+76)) = v275
	goto L53
L52:
	;
	goto L53
L53:
	;
	if v275 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _consts[91]))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)))
	*(*int32)(unsafe.Add(mBase, uint32(v287+v275*int32(640))+80)) = v276
	goto L56
L55:
	;
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v274))) = int64(0)
	v294 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v272)+74)) = uint8(v294)
	if v275 != int32(-1) {
		v254 = v275
		goto L49
	} else {
		goto L57
	}
L57:
	;
	goto L50
}
