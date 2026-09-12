package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TParserGet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 <= v14 {
		v249 = v2
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	return v249
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v17 + v14
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v24 < v23 {
		v249 = v2
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v27 = v23
	v29 = v24
	v30 = v22
	goto L10
L9:
	;
	v249 = v244 & int32(1)
	goto L6
L10:
	;
	if v27 == v29 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+6)))
	v244 = v240
	goto L9
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45)+8)) = v44
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+28))
	if v48 != 0 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	v44 = int32(0)
	v45 = v30
	goto L12
L14:
	;
	goto L15
L15:
	;
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v35 == v34 {
		v44 = v34
		v45 = v30
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = F_pg_mblen_range(m, v38+v27, v38+v29)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v44 = v41
	v45 = v43
	goto L12
L18:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v60 == int32(0) {
		v79 = v59
		goto L18
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = int32(0)
	v59 = v48 + int32(20)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v53<<(uint(int32(3))%32))+uint32(_consts[896])))
	v59 = v58
	goto L19
L23:
	;
	v65 = v59
	goto L24
L24:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)) = uint8(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v72 = m.T0[v71].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L26
	}
L25:
	;
	v79 = v75
	goto L18
L26:
	;
	if v72 != 0 {
		v79 = v65
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v75 = v65 + int32(20)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v76 != 0 {
		v65 = v75
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	m.T0[v83].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+6)))
	if v86&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v92
	v94 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+16)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v94
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v99
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+6)))
	v102 = v101
	goto L35
L34:
	;
	v102 = v86
	goto L35
L35:
	;
	if v102&int32(2) != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v193 != int32(77) {
		goto L63
	} else {
		goto L64
	}
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	F_pfree(m, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v111 = v102 & int32(65535)
	if v111&int32(4) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v106
	goto L36
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v114)+28)) = v79
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v118 = F_palloc(m, int32(32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v111&int32(16) != 0 {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	if v116 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+24)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v118
	goto L36
L46:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v116)))
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = v120
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v116)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v118)+16)) = v122
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v116)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v118)+8)) = v124
	goto L45
L47:
	;
	goto L48
L48:
	;
	v126 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v118)+16)) = v126
	*(*int64)(unsafe.Add(mBase, uint32(v118)+8)) = v126
	goto L45
L49:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+24))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+24))
	F_pfree(m, v139)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v111&int32(64) != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+24)) = v140
	goto L36
L53:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+24))
	if v148 == int32(0) {
		goto L36
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if v111&int32(32) == int32(0) {
		goto L36
	} else {
		goto L61
	}
L56:
	;
	v152 = v148
	goto L57
L57:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v152)+24))
	F_pfree(m, v152)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L59
	}
L58:
	;
	goto L36
L59:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v160)+24)) = v157
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+24))
	if v163 != 0 {
		v152 = v163
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v174
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+8)) = v177
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v179)+12)) = v180
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v182)+16)) = v183
	F_pfree(m, v168)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	goto L36
L63:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+20)) = v193
	goto L65
L64:
	;
	goto L65
L65:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+6)))
	if v198&int32(1) != 0 {
		v244 = v198
		goto L9
	} else {
		goto L66
	}
L66:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.B2i32(v198&int32(8) == int32(0))&base.B2i32(v207 <= v206) != 0 {
		v244 = v198
		goto L9
	} else {
		goto L67
	}
L67:
	;
	if v198&int32(10) != 0 {
		v235 = v206
		v236 = v207
		v237 = v205
		goto L68
	} else {
		goto L69
	}
L68:
	;
	if v235 <= v236 {
		v27 = v235
		v29 = v236
		v30 = v237
		goto L10
	} else {
		goto L71
	}
L69:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v205)+8))
	if v212 == int32(0) {
		v235 = v206
		v236 = v207
		v237 = v205
		goto L68
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v205))) = v206 + v212
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+12)) = v218 + v219
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v224 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v223 + v224
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+16)) = v228 + v224
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = v234
	v236 = v232
	v237 = v233
	goto L68
L71:
	;
	goto L11
}
func F_TeardownHistoricSnapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1192])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[1186])) = v2
	return
}
func F_TerminateBufferIO(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(224563)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(484021)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v22 | v23
	if v22&v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L4
L2:
	;
	v50 = v22
	goto L3
L3:
	;
	if v50&int32(268435456) != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_perform_spin_delay(m, v10+int32(8))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v50 = v39
	goto L3
L6:
	;
	return
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v40 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v39 | v40
	if v39&v40 != 0 {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v58 = int32(-201326593)
	goto L11
L10:
	;
	v58 = int32(-1551892481)
	goto L11
L11:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v60 = v58
	goto L14
L13:
	;
	v60 = int32(-201326593)
	goto L14
L14:
	;
	v61 = (v50 | int32(4194304)) & v60
	v65 = int32(4083212)
	v66 = *(*int32)(unsafe.Add(mBase, _consts[324]))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8))+8))
	if v68 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if l4 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[324])) = v83
	goto L16
L18:
	;
	if int32(999) < v66 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v66 < int32(11) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v73 = int32(900)
	if v73 <= v66 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = v73
	goto L24
L23:
	;
	v76 = v66
	goto L24
L24:
	;
	v83 = v76 + int32(100)
	goto L17
L25:
	;
	v83 = v66 - int32(1)
	goto L17
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
	goto L29
L27:
	;
	v91 = v61
	goto L28
L28:
	;
	v92 = v91 | l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v92 & int32(-4194305)
	if l3 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v91 = v61 - int32(1)
	goto L28
L30:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[170]))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerForget(m, v97, v98+int32(1), int32(1599100))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _consts[802]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ConditionVariableBroadcast(m, v105+v106<<(uint(int32(4))%32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L6
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	if l4 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	m.G0 = v10 + int32(32)
	return
L36:
	;
	if v92&int32(536870912) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(224563)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(484021)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = int64(0)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v129 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v128 | v129
	if v128&v129 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	goto L41
L39:
	;
	v156 = v128
	goto L40
L40:
	;
	v161 = int32(4083212)
	v162 = *(*int32)(unsafe.Add(mBase, _consts[324]))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(8))+8))
	if v164 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L41:
	;
	F_perform_spin_delay(m, v10+int32(8))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L6
	} else {
		goto L43
	}
L42:
	;
	v156 = v145
	goto L40
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v146 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v145 | v146
	if v145&v146 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	if v156&int32(537133055) == int32(536870913) {
		goto L56
	} else {
		goto L57
	}
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[324])) = v179
	goto L46
L48:
	;
	if int32(999) < v162 {
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v162 < int32(11) {
		goto L46
	} else {
		goto L55
	}
L51:
	;
	v169 = int32(900)
	if v169 <= v162 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v172 = v169
	goto L54
L53:
	;
	v172 = v162
	goto L54
L54:
	;
	v179 = v172 + int32(100)
	goto L47
L55:
	;
	v179 = v162 - int32(1)
	goto L47
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v156 & int32(-541327359)
	F_ProcSendSignal(m, v185)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L6
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v156 & int32(-4194305)
	goto L35
L59:
	;
	goto L35
}
func F_tan(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v50 float64
	_ = v50
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v97 float64
	_ = v97
	var v102 float64
	_ = v102
	var v104 float64
	_ = v104
	var v106 float64
	_ = v106
	var v131 float64
	_ = v131
	var v135 int32
	_ = v135
	var v136 float64
	_ = v136
	var v137 float64
	_ = v137
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v159 float64
	_ = v159
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v165 int32
	_ = v165
	var v166 float64
	_ = v166
	var v167 float64
	_ = v167
	var v170 float64
	_ = v170
	var v209 float64
	_ = v209
	var v210 float64
	_ = v210
	var v213 int32
	_ = v213
	var v217 float64
	_ = v217
	var v222 float64
	_ = v222
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v228 float64
	_ = v228
	var v230 int64
	_ = v230
	var v232 float64
	_ = v232
	var v236 float64
	_ = v236
	var v248 float64
	_ = v248
	var v251 float64
	_ = v251
	var v252 float64
	_ = v252
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v13 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v13) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v13) < base.Ui32(int32(1044381696)) {
			v252 = l0
		} else {
			v18 = float64(0)
			v19 = int32(0)
			v25 = base.I64_reinterpret_f64(l0)
			v29 = base.B2i32(base.Ui64(v25&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)))
			if v29 == v19 {
				v38 = base.B2i32(int64(0) <= v25)
				if int64(0) <= v25 {
					v39 = v18
				} else {
					v39 = base.F64_neg(v18)
				}
				v43 = base.F64_add(base.F64_sub(float64(0.7853981633974483), base.F64_abs(l0)), base.F64_sub(float64(3.061616997868383e-17), v39))
				v44 = float64(0)
				v45 = v38
			} else {
				v43 = l0
				v44 = v18
				v45 = v19
			}
			v46 = base.F64_mul(v43, v43)
			v47 = base.F64_mul(v43, v46)
			v50 = base.F64_mul(v46, v46)
			v89 = base.F64_add(base.F64_mul(v47, float64(0.3333333333333341)), base.F64_add(base.F64_mul(v46, base.F64_add(base.F64_mul(v47, base.F64_add(base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, float64(-1.8558637485527546e-05)), float64(7.817944429395571e-05))), float64(0.0005880412408202641))), float64(0.0035920791075913124))), float64(0.021869488294859542))), float64(0.13333333333320124)), base.F64_mul(v46, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, base.F64_add(base.F64_mul(v50, float64(2.590730518636337e-05)), float64(7.140724913826082e-05))), float64(0.0002464631348184699))), float64(0.0014562094543252903))), float64(0.0088632398235993))), float64(0.05396825397622605))))), v44)), v44))
			v90 = base.F64_add(v43, v89)
			if v29 == int32(0) {
				v97 = base.F64_convert_i32_s(int32(1))
				v102 = base.F64_add(v43, base.F64_sub(v89, base.F64_div(base.F64_mul(v90, v90), base.F64_add(v90, v97))))
				v104 = base.F64_sub(v97, base.F64_add(v102, v102))
				if v45 != 0 {
					v106 = v104
				} else {
					v106 = base.F64_neg(v104)
				}
				v131 = v106
			} else {
				v131 = v90
			}
			v252 = v131
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v13) {
			v252 = base.F64_sub(l0, l0)
		} else {
			v135 = F___rem_pio2(m, l0, v6)
			mBase = m.M
			v136 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
			v137 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			v139 = v135 & int32(1)
			v143 = int32(0)
			v145 = base.I64_reinterpret_f64(v136)
			v149 = base.B2i32(base.Ui64(v145&int64(9223372002495037440)) < base.Ui64(int64(4604249089280835585)))
			if v149 == v143 {
				v158 = base.B2i32(int64(0) <= v145)
				if int64(0) <= v145 {
					v159 = v137
				} else {
					v159 = base.F64_neg(v137)
				}
				v163 = base.F64_add(base.F64_sub(float64(0.7853981633974483), base.F64_abs(v136)), base.F64_sub(float64(3.061616997868383e-17), v159))
				v164 = float64(0)
				v165 = v158
			} else {
				v163 = v136
				v164 = v137
				v165 = v143
			}
			v166 = base.F64_mul(v163, v163)
			v167 = base.F64_mul(v163, v166)
			v170 = base.F64_mul(v166, v166)
			v209 = base.F64_add(base.F64_mul(v167, float64(0.3333333333333341)), base.F64_add(base.F64_mul(v166, base.F64_add(base.F64_mul(v167, base.F64_add(base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_mul(v170, float64(-1.8558637485527546e-05)), float64(7.817944429395571e-05))), float64(0.0005880412408202641))), float64(0.0035920791075913124))), float64(0.021869488294859542))), float64(0.13333333333320124)), base.F64_mul(v166, base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_mul(v170, base.F64_add(base.F64_mul(v170, float64(2.590730518636337e-05)), float64(7.140724913826082e-05))), float64(0.0002464631348184699))), float64(0.0014562094543252903))), float64(0.0088632398235993))), float64(0.05396825397622605))))), v164)), v164))
			v210 = base.F64_add(v163, v209)
			if v149 == int32(0) {
				v213 = int32(1)
				v217 = base.F64_convert_i32_s(v213 - v139<<(uint(v213)%32))
				v222 = base.F64_add(v163, base.F64_sub(v209, base.F64_div(base.F64_mul(v210, v210), base.F64_add(v210, v217))))
				v224 = base.F64_sub(v217, base.F64_add(v222, v222))
				if v165 != 0 {
					v226 = v224
				} else {
					v226 = base.F64_neg(v224)
				}
				v251 = v226
			} else {
				if v139 != 0 {
					v228 = base.F64_div(float64(-1), v210)
					v230 = int64(-4294967296)
					v232 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v228) & v230)
					v236 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v210) & v230)
					v248 = base.F64_add(base.F64_mul(v228, base.F64_add(base.F64_mul(v232, base.F64_sub(v209, base.F64_sub(v236, v163))), base.F64_add(base.F64_mul(v232, v236), float64(1)))), v232)
				} else {
					v248 = v210
				}
				v251 = v248
			}
			v252 = v251
		}
	}
	m.G0 = v6 + int32(16)
	return v252
}
func F_tblspc_identify(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v5 = l0 & int32(240)
	if v5 == int32(16) {
		v8 = int32(515006)
	} else {
		v8 = int32(0)
	}
	if v5 != 0 {
		v10 = v8
	} else {
		v10 = int32(526667)
	}
	return v10
}
func F_tblspc_redo(m *base.Module, l0 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+48)))
	v11 = v9 & int32(240)
	switch v11 {
	case 0:
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
		v76 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
		F_create_tablespace_directories(m, v73+int32(4), v76)
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			m.G0 = v6 + int32(32)
			return
		}
	default:
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v63 = m.ExcPending
		if v63 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v11
			F_errmsg_internal(m, int32(51745), v6)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				F_errfinish(m, int32(488308), int32(1568), int32(237345))
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 16:
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
		v13 = F_EmitProcSignalBarrier(m)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_WaitForProcSignalBarrier(m, v13)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v19 = F_destroy_tablespace_directories(m, v17, int32(1))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					if v19 != 0 {
						m.G0 = v6 + int32(32)
						return
					} else {
						v22 = int32(0)
						v24 = F_GetConflictingVirtualXIDs(m, v22, v22)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							F_ResolveRecoveryConflictWithVirtualXIDs(m, v24, int32(8), int32(134217773), int32(1))
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v33 = F_destroy_tablespace_directories(m, v31, int32(1))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									if v33 != 0 {
										m.G0 = v6 + int32(32)
										return
									} else {
										v37 = F_errstart(m, int32(15), int32(0))
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											if v37 == int32(0) {
												m.G0 = v6 + int32(32)
												return
											} else {
												F_errcode(m, int32(325))
												mBase = m.M
												v43 = m.ExcPending
												if v43 != 0 {
													return
												} else {
													v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
													*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v44
													F_errmsg(m, int32(430736), v6+int32(16))
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return
													} else {
														F_errhint(m, int32(546326), int32(0))
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return
														} else {
															F_errfinish(m, int32(488308), int32(1564), int32(237345))
															mBase = m.M
															v59 = m.ExcPending
															if v59 != 0 {
																return
															} else {
																m.G0 = v6 + int32(32)
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
func F_text2ltree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_DirectFunctionCall1Coll(m, int32(5635), int32(0), v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v12)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v18 != v8 {
						F_pfree(m, v8)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							return v14
						}
					} else {
						return v14
					}
				}
			}
		}
	}
}
func F_textlename(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1558), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 <= int32(0))
	}
}
func F_textnename(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
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
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
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
	var v181 int32
	_ = v181
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v14 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if v13&int32(3) == int32(0) {
		v68 = v13
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v17 = int32(4)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v19&int32(254) == int32(2) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v32 = int32(1)
	if v14&v32 != 0 {
		v44 = int32(base.Ui32(v14)>>(uint(v32)%32)) - v32
		goto L3
	} else {
		goto L13
	}
L7:
	;
	v28 = v17
	goto L9
L8:
	;
	v28 = base.B2i32(v19 == int32(18)) << (uint(v17) % 32)
	goto L9
L9:
	;
	if v19 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = v17
	goto L12
L11:
	;
	v31 = v28
	goto L12
L12:
	;
	v44 = v31
	goto L3
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L3
L14:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v102 != int32(950) {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	v101 = v93 - v13
	goto L14
L16:
	;
	v72 = v68
	goto L25
L17:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v52 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v101 = int32(0)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v57 = v13
	goto L21
L21:
	;
	v61 = v57 + int32(1)
	if v61&int32(3) == int32(0) {
		v68 = v61
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v93 = v61
	goto L15
L23:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v66 != 0 {
		v57 = v61
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v81 = int32(-2139062144)
	if (int32(16843008)-v78|v78)&v81 == v81 {
		v72 = v72 + int32(4)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v87 = v72
	goto L28
L27:
	;
	goto L26
L28:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v91 != 0 {
		v87 = v87 + int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v93 = v87
	goto L15
L30:
	;
	goto L29
L31:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v208 != v9 {
		goto L68
	} else {
		goto L69
	}
L32:
	;
	v197 = int32(1)
	if v14&v197 != 0 {
		goto L64
	} else {
		goto L65
	}
L33:
	;
	if v102 != 0 {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if v44 != v101 {
		v207 = int32(1)
		goto L31
	} else {
		goto L42
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(240014), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errhint(m, int32(546990), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(488865), int32(1648), int32(103562))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v127 = int32(1)
	if v14&v127 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v131 = v127
	goto L45
L44:
	;
	v131 = int32(4)
	goto L45
L45:
	;
	v132 = v9 + v131
	if base.Ui32(int32(4)) <= base.Ui32(v44) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v207 = base.B2i32(v194 != int32(0))
	goto L31
L47:
	;
	v194 = int32(0)
	goto L46
L48:
	;
	v168 = v163
	v169 = v164
	v170 = v165
	goto L58
L49:
	;
	if (v132|v13)&int32(3) != 0 {
		v163 = v132
		v164 = v13
		v165 = v44
		goto L48
	} else {
		goto L52
	}
L50:
	;
	v156 = v132
	v157 = v13
	v158 = v44
	goto L51
L51:
	;
	if v158 == int32(0) {
		goto L47
	} else {
		goto L57
	}
L52:
	;
	v140 = v132
	v141 = v13
	v142 = v44
	goto L53
L53:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	if v145 != v146 {
		v163 = v140
		v164 = v141
		v165 = v142
		goto L48
	} else {
		goto L55
	}
L54:
	;
	v156 = v151
	v157 = v149
	v158 = v153
	goto L51
L55:
	;
	v148 = int32(4)
	v149 = v141 + v148
	v151 = v140 + v148
	v153 = v142 - v148
	if base.Ui32(int32(3)) < base.Ui32(v153) {
		v140 = v151
		v141 = v149
		v142 = v153
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v163 = v156
	v164 = v157
	v165 = v158
	goto L48
L58:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	if v173 == v174 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v194 = v173 - v174
	goto L46
L60:
	;
	v176 = int32(1)
	v181 = v170 - v176
	if v181 != 0 {
		v168 = v168 + v176
		v169 = v169 + v176
		v170 = v181
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	goto L47
L64:
	;
	v201 = v197
	goto L66
L65:
	;
	v201 = int32(4)
	goto L66
L66:
	;
	v203 = F_varstr_cmp(m, v9+v201, v44, v13, v101, v102)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v207 = base.B2i32(v203 != int32(0))
	goto L31
L68:
	;
	F_pfree(m, v9)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	return v207
L71:
	;
	goto L70
}
func F_textoverlay(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v12 = F_text_overlay(m, v3, v8, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_textregexne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = v8 + int32(1)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = F_pg_detoast_datum_packed(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			v18 = int32(1)
			v19 = v17 & v18
			if v17 == v18 {
				v22 = int32(4)
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				if v24&int32(254) == int32(2) {
					v33 = v22
				} else {
					v33 = base.B2i32(v24 == int32(18)) << (uint(v22) % 32)
				}
				if v24 == int32(1) {
					v36 = v22
				} else {
					v36 = v33
				}
				v47 = v36
			} else {
				v37 = int32(1)
				if v19 != 0 {
					v47 = int32(base.Ui32(v17)>>(uint(v37)%32)) - v37
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v50 = F_RE_compile_and_cache(m, v15, int32(19), v49)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v56 = F_palloc(m, v47<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					if v19 != 0 {
						v60 = v13
					} else {
						v60 = v8 + int32(4)
					}
					v61 = F_pg_mb2wchar_with_len(m, v60, v56, v47)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = int32(0)
						v66 = F_RE_wchar_execute(m, v56, v61, v63, v63, v63)
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v56)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								return v66 ^ int32(1)
							}
						}
					}
				}
			}
		}
	}
}
func F_textregexreplace_extended_no_n(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_textregexreplace_extended(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_textregexreplace_noopt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v11 = F_pg_detoast_datum_packed(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v17 = F_replace_text_regexp(m, v3, v8, v11, int32(3), v14, int32(0), int32(1))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_textregexsubstr(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = F_RE_compile_and_cache(m, v18, int32(3), v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(1)
				v25 = v13 + v24
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
				v30 = v28 & v24
				if v30 != 0 {
					v31 = v25
				} else {
					v31 = v13 + int32(4)
				}
				if v28 == int32(1) {
					v34 = int32(4)
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
					if v36&int32(254) == int32(2) {
						v45 = v34
					} else {
						v45 = base.B2i32(v36 == int32(18)) << (uint(v34) % 32)
					}
					if v36 == int32(1) {
						v48 = v34
					} else {
						v48 = v45
					}
					v59 = v48
				} else {
					v49 = int32(1)
					if v30 != 0 {
						v59 = int32(base.Ui32(v28)>>(uint(v49)%32)) - v49
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v60 = int32(0)
				v65 = F_palloc(m, v59<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					v67 = F_pg_mb2wchar_with_len(m, v31, v65, v59)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v71 = F_RE_wchar_execute(m, v65, v67, int32(0), int32(2), v10)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v65)
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								if v71 == int32(0) {
									v77 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v77)
									v107 = v60
									m.G0 = v10 + int32(16)
									return v107
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, _consts[1025]))
									if v82 != 0 {
										v83 = v10 | int32(8)
									} else {
										v83 = v10
									}
									v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
									if int32(0) <= v84 {
										if v82 != 0 {
											v91 = v10 | int32(12)
										} else {
											v91 = v10 | int32(4)
										}
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
										if int32(0) <= v92 {
											v103 = F_DirectFunctionCall3Coll(m, int32(1492), int32(0), v13, v84+int32(1), v92-v84)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return int32(0)
											} else {
												v107 = v103
												m.G0 = v10 + int32(16)
												return v107
											}
										} else {
											v96 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
											v107 = v60
											m.G0 = v10 + int32(16)
											return v107
										}
									} else {
										v96 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
										v107 = v60
										m.G0 = v10 + int32(16)
										return v107
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
func F_tfuncFetchRows(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v307 int32
	_ = v307
	var v322 int32
	_ = v322
	var v338 int32
	_ = v338
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v391 int32
	_ = v391
	var v406 int32
	_ = v406
	var v422 int32
	_ = v422
	var v439 int32
	_ = v439
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v598 int32
	_ = v598
	var v613 int32
	_ = v613
	var v629 int32
	_ = v629
	var v649 int32
	_ = v649
	var v666 int32
	_ = v666
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v726 int32
	_ = v726
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v824 int32
	_ = v824
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v862 int32
	_ = v862
	var v875 int32
	_ = v875
	var v896 int64
	_ = v896
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v978 int32
	_ = v978
	var v1001 int32
	_ = v1001
	var v1016 int32
	_ = v1016
	var v1034 int32
	_ = v1034
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1105 int32
	_ = v1105
	var v1119 int32
	_ = v1119
	var v1120 int32
	_ = v1120
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1203 int32
	_ = v1203
	var v1228 int32
	_ = v1228
	var v1229 int32
	_ = v1229
	var v1243 int32
	_ = v1243
	var v1249 int32
	_ = v1249
	var v1263 int32
	_ = v1263
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1286 int32
	_ = v1286
	var v1301 int32
	_ = v1301
	var v1326 int32
	_ = v1326
	var v1338 int32
	_ = v1338
	var v1339 int64
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	v3 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(80)
	m.G0 = v39
	v43 = l0 + int32(156)
	v51 = v3
	v52 = v3
	v53 = v3
	v54 = v3
	v55 = v3
	v56 = v3
	v57 = v3
	v58 = v3
	v59 = v3
	v60 = v3
	v61 = v3
	v62 = v3
	v63 = v3
	v64 = int32(-1)
	v72 = v39
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
	if v64 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v1338 = int32(m.ExcTag)
	v1339 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1338 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L7:
	;
	v86 = int32(16)
	v87 = v72 - v86
	m.G0 = v87
	v90 = v87 - v86
	m.G0 = v90
	v93 = v90 - v86
	m.G0 = v93
	v96 = v93 - int32(160)
	m.G0 = v96
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v99 = int32(4464496)
	v100 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v87
	v116 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, _consts[523]))
	v120 = F_tuplestore_begin_heap(m, v116, v116, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		v1326 = v96
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v136 = v51
	v137 = v52
	v138 = v53
	v139 = v54
	v140 = v55
	v141 = v56
	v142 = v57
	v143 = v58
	v144 = v59
	v145 = v60
	v146 = v61
	v147 = v62
	v148 = v72
	goto L9
L9:
	;
	if v136 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v120
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v125
	v128 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	v130 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v39 + int32(28)
	goto L14
L12:
	;
	v136 = int32(0)
	v137 = v87
	v138 = v90
	v139 = v98
	v140 = v93
	v141 = v96
	v142 = v130
	v143 = v128
	v144 = l0 + int32(176)
	v145 = v100
	v146 = l0 + int32(180)
	v147 = v43
	v148 = v96
	goto L9
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v141
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	m.T0[v153].(func(*base.Module, int32, int32))(m, l0, v156)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v142
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v143
	v1271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v1271 != 0 {
		goto L138
	} else {
		goto L139
	}
L18:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v185 = m.T0[v172].(func(*base.Module, int32, int32, int32) int32)(m, v171, l1, v140)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v187 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+80))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+64))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	m.T0[v194].(func(*base.Module, int32, int32))(m, l0, v185)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L23
	}
L21:
	;
	v1203 = v63
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v142
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v143
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	if v1228 != 0 {
		goto L133
	} else {
		goto L134
	}
L23:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v215 = int32(0)
	goto L25
L24:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v506 == int32(0) {
		v714 = v505
		v726 = v63
		goto L59
	} else {
		goto L60
	}
L25:
	;
	v248 = int32(0)
	if v210 == v248 {
		v258 = v248
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v487 = F_text_to_cstring(m, v287)
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L57
	}
L27:
	;
	if v209 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L28:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v252 <= v215 {
		v258 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	v258 = v254 + v215<<(uint(int32(2))%32)
	goto L27
L30:
	;
	goto L26
L31:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v258)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v371 = m.T0[v358].(func(*base.Module, int32, int32, int32) int32)(m, v357, l1, v138)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L44
	}
L32:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v193)+12))
	if v270 == int32(0) {
		goto L24
	} else {
		goto L37
	}
L33:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v261 <= v215 {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	if v258 == int32(0) {
		goto L32
	} else {
		goto L35
	}
L35:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v209)+12))
	v268 = v265 + v215<<(uint(int32(2))%32)
	if v268 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v287 = m.T0[v274].(func(*base.Module, int32, int32, int32) int32)(m, v273, l1, v138)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v289 != int32(1) {
		goto L30
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errcode(m, int32(67108994))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errmsg(m, int32(296500), int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errfinish(m, int32(485797), int32(389), int32(334004))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L43
	}
L43:
	;
	goto L3
L44:
	;
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v373 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v452 = F_text_to_cstring(m, v371)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L52
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errcode(m, int32(67108994))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errmsg(m, int32(296700), int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errfinish(m, int32(485797), int32(370), int32(334004))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L51
	}
L51:
	;
	goto L3
L52:
	;
	if v356 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v356)+4))
	v456 = v454
	goto L55
L54:
	;
	v456 = int32(0)
	goto L55
L55:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	m.T0[v457].(func(*base.Module, int32, int32, int32))(m, l0, v456, v452)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L56
	}
L56:
	;
	v215 = v215 + int32(1)
	goto L25
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	m.T0[v474].(func(*base.Module, int32, int32))(m, l0, v487)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L58
	}
L58:
	;
	goto L24
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = int64(1)
	v749 = int32(4464496)
	v750 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v714)+20))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v714)+16))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)+80))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v754)+64))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v714)+12))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v760
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v758)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v775 = m.T0[v762].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L83
	}
L60:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	if v509 <= int32(0) {
		v714 = v505
		v726 = v63
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v505)+12))
	v519 = int32(0)
	v531 = v63
	v532 = v509
	goto L62
L62:
	;
	if v519 != v192 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v714 = v710
	v726 = v703
	goto L59
L64:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v559 = v512 + int32(20) + v553<<(uint(int32(4))%32) + v519*int32(100)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v560+v519<<(uint(int32(2))%32))))
	if v564 != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v703 = v531
	v704 = v532
	goto L66
L66:
	;
	v708 = v519 + int32(1)
	if v708 < v704 {
		v519 = v708
		v531 = v703
		v532 = v704
		goto L62
	} else {
		goto L82
	}
L67:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v683
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	m.T0[v687].(func(*base.Module, int32, int32, int32))(m, l0, v686, v519)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L81
	}
L68:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v564)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v578 = m.T0[v565].(func(*base.Module, int32, int32, int32) int32)(m, v564, l1, v138)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v683 = v531
	v686 = v559 + int32(4)
	goto L67
L71:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
	if v580 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v679 = F_text_to_cstring(m, v578)
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L80
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errcode(m, int32(67108994))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errmsg(m, int32(296539), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v559 + int32(4)
	F_errdetail(m, int32(593539), v39+int32(16))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errfinish(m, int32(485797), int32(418), int32(334004))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L79
	}
L79:
	;
	goto L3
L80:
	;
	v683 = v679
	v686 = v679
	goto L67
L81:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v703 = v683
	v704 = v702
	goto L66
L82:
	;
	goto L63
L83:
	;
	if v775 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	goto L87
L85:
	;
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v750
	v1203 = v726
	goto L22
L87:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v819 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L86
L89:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)+12))
	v822 = v820
	goto L91
L90:
	;
	v822 = int32(0)
	goto L91
L91:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v824 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_ProcessInterrupts(m)
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v839)+8))
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v840)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	m.T0[v841].(func(*base.Module, int32))(m, v839)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	v856 = int32(0)
	if v856 < v757 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v862 = v856
	v875 = v822
	goto L100
L98:
	;
	goto L99
L99:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_tuplestore_putvalues(m, v1105, v756, v752, v751)
	mBase = m.M
	v1119 = m.ExcPending
	if v1119 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L129
	}
L100:
	;
	if v862 == v755 {
		goto L106
	} else {
		goto L107
	}
L101:
	;
	goto L99
L102:
	;
	v1067 = v862 + int32(1)
	if v1067 != v757 {
		v862 = v1067
		v875 = v1065
		goto L100
	} else {
		goto L128
	}
L103:
	;
	v1053 = v875 + int32(4)
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+12))
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(v1055)+4))
	if base.Ui32(v1053) < base.Ui32(v1056+v1057<<(uint(int32(2))%32)) {
		goto L125
	} else {
		goto L126
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L121
	}
L105:
	;
	if v875 != 0 {
		goto L103
	} else {
		goto L120
	}
L106:
	;
	v896 = *(*int64)(unsafe.Add(mBase, uint32(l0)+168))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v896 + int64(1)
	*(*uint32)(unsafe.Add(mBase, uint32(v752+v755<<(uint(int32(2))%32)))) = uint32(v896)
	v901 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v755+v751))) = uint8(v901)
	goto L105
L107:
	;
	goto L108
L108:
	;
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	v909 = v756 + int32(20) + v903<<(uint(int32(4))%32) + v862*int32(100)
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v909)+76))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v909)+68))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v758)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v925 = m.T0[v912].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, v862, v911, v910, v137)
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L109
	}
L109:
	;
	v929 = v752 + v862<<(uint(int32(2))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v929))) = v925
	v931 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	if v931 != int32(1) {
		v956 = v931
		goto L110
	} else {
		goto L111
	}
L110:
	;
	if v956&int32(1) != 0 {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	if v875 == int32(0) {
		v956 = v931
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v875)))
	if v936 == int32(0) {
		v956 = v931
		goto L110
	} else {
		goto L113
	}
L113:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v936)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v952 = m.T0[v939].(func(*base.Module, int32, int32, int32) int32)(m, v936, l1, v137)
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v929))) = v952
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v956 = v955
	goto L110
L115:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v974 = F_bms_is_member(m, v862, v961)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L118
	}
L116:
	;
	v978 = v956
	goto L117
L117:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v862+v751))) = uint8(v978)
	goto L105
L118:
	;
	if v974 != 0 {
		goto L104
	} else {
		goto L119
	}
L119:
	;
	v976 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	v978 = v976
	goto L117
L120:
	;
	v1065 = int32(0)
	goto L102
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errcode(m, int32(67108994))
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = v909 + int32(4)
	F_errmsg(m, int32(681919), v39)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_errfinish(m, int32(485797), int32(508), int32(111249))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L124
	}
L124:
	;
	goto L3
L125:
	;
	v1062 = v1053
	goto L127
L126:
	;
	v1062 = int32(0)
	goto L127
L127:
	;
	v1065 = v1062
	goto L102
L128:
	;
	goto L101
L129:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_MemoryContextReset(m, v1120)
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L130
	}
L130:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v758)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v726
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	v1148 = m.T0[v1135].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L131
	}
L131:
	;
	if v1148 != 0 {
		goto L87
	} else {
		goto L132
	}
L132:
	;
	goto L88
L133:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v1203
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	m.T0[v1229].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v145
	v1249 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v1203
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_MemoryContextReset(m, v1249)
	mBase = m.M
	v1263 = m.ExcPending
	if v1263 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = int32(0)
	goto L135
L137:
	;
	m.G0 = v39 + int32(80)
	return
L138:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(v139)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	m.T0[v1272].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v1286 = m.ExcPending
	if v1286 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v39)+40)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v39)+56)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v39)+60)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v39)+64)) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v39)+72)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v137
	F_pg_re_throw(m)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		v1326 = v148
		goto L6
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	goto L5
L143:
	;
	v1343 = int32(v1339)
	m.G0 = v1326
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v1343)+4))
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v1343)))
	v1350 = *(*int32)(unsafe.Add(mBase, uint32(v1346)))
	if v39+int32(28) == v1350 {
		goto L146
	} else {
		goto L147
	}
L144:
	;
	m.ExcPending = 1
	goto L152
L145:
	;
	if v1353 != 0 {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(v1346)+4))
	v1353 = v1352
	goto L148
L147:
	;
	v1353 = int32(0)
	goto L148
L148:
	;
	goto L145
L149:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v39)+76))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	v1356 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v39)+64))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v39)+60))
	v1359 = *(*int32)(unsafe.Add(mBase, uint32(v39)+56))
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1362 = *(*int32)(unsafe.Add(mBase, uint32(v39)+44))
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(v39)+40))
	v1364 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v51 = v1345
	v52 = v1354
	v53 = v1355
	v54 = v1359
	v55 = v1356
	v56 = v1357
	v57 = v1364
	v58 = v1363
	v59 = v1362
	v60 = v1360
	v61 = v1361
	v62 = v1358
	v63 = v1365
	v64 = v1353
	v72 = v1326
	goto L1
L150:
	;
	goto L151
L151:
	;
	F___wasm_longjmp(m, v1346, v1345)
	mBase = m.M
	v1367 = m.ExcPending
	if v1367 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	return
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tidge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return int32(base.Ui32(v27^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_tidlt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+2)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2))))
	v9 = int32(16)
	v11 = v7 | v8<<(uint(v9)%32)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+2)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3))))
	v16 = v12 | v13<<(uint(v9)%32)
	if base.Ui32(v11) < base.Ui32(v16) {
		v27 = int32(-1)
	} else {
		if base.Ui32(v16) < base.Ui32(v11) {
			v27 = int32(1)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2)+4)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3)+4)))
			if base.Ui32(v21) < base.Ui32(v22) {
				v27 = int32(-1)
			} else {
				v27 = base.B2i32(base.Ui32(v22) < base.Ui32(v21))
			}
		}
	}
	return int32(base.Ui32(v27) >> (uint(int32(31)) % 32))
}
func F_timestamp2timestamptz_opt_overflow(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int64
	_ = v17
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v129 int64
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v153 int64
	_ = v153
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int64
	_ = v168
	var v169 int64
	_ = v169
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	if l1 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	} else {
	}
	if base.Ui64(l0-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v169 = l0
		m.G0 = v8 + int32(48)
		return v169
	} else {
		v17 = base.I64_div_s(l0, int64(86400000000))
		if base.Ui64(int64(172799999999)) <= base.Ui64(l0+int64(86399999999)) {
			v25 = v17 * int64(-86400000000)
		} else {
			v25 = int64(0)
		}
		v26 = v25 + l0
		v29 = v26>>(uint(int64(63))%64) + v17
		if v29 < int64(-2451545) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v185 = m.ExcPending
			if v185 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v188 = m.ExcPending
				if v188 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(393503), int32(0))
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(484742), int32(6509), int32(30234))
						mBase = m.M
						v197 = m.ExcPending
						if v197 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v32 = base.I32_wrap_i64(v29)
			v44 = v32 + int32(2483589)
			v45 = int32(146097)
			v46 = base.I32_div_u_s(v44, v45)
			v47 = int32(3)
			v53 = int32(2)
			v58 = base.I32_div_u_s((v46*int32(1073595727)+v44)<<(uint(v53)%32)|v47, v45)
			v61 = v32 + int32(2451545) + v46*v47 + v58 + int32(32104)
			v62 = int32(1461)
			v63 = base.I32_div_u_s(v61, v62)
			v66 = v63*int32(-1461) + v61
			v68 = v66 << (uint(v53) % 32)
			if base.Ui32(v62) <= base.Ui32(v68) {
				v74 = base.I32_rem_u_s(v66+int32(305), int32(365))
				v79 = v74
			} else {
				v78 = base.I32_rem_u_s(v66+int32(306), int32(366))
				v79 = v78
			}
			v81 = base.I32_div_u_s(v68, int32(1461))
			*(*int32)(unsafe.Add(mBase, uint32(v8+int32(24)))) = v81 + v63<<(uint(int32(2))%32) - int32(4800)
			v89 = v79 + int32(123)
			v93 = int32(base.Ui32(v89*int32(2141)) >> (uint(int32(16)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v8+int32(16)))) = v89 - int32(base.Ui32(v93*int32(7834))>>(uint(int32(8))%32))
			v103 = base.I32_rem_u_s(v93+int32(10), int32(12))
			*(*int32)(unsafe.Add(mBase, uint32(v8+int32(20)))) = v103 + int32(1)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = int64(4294967295)
			if v26 < int64(0) {
				v113 = v26 + int64(86400000000)
			} else {
				v113 = v26
			}
			v115 = base.I64_div_s(v113, int64(3600000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v8)+12)) = uint32(v115)
			v120 = base.I64_extend32_s(v115)*int64(-3600000000) + v113
			v122 = base.I64_div_s(v120, int64(60000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v8)+8)) = uint32(v122)
			v129 = base.I64_div_s(base.I64_extend32_s(v122)*int64(-60000000)+v120, int64(1000000))
			*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)) = uint32(v129)
			v131 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v131
			v137 = *(*int32)(unsafe.Add(mBase, _consts[515]))
			v139 = m.G0
			v140 = int32(16)
			v141 = v139 - v140
			m.G0 = v141
			v145 = F_DetermineTimeZoneOffsetInternal(m, v8+int32(4), v137, v141+int32(8))
			mBase = m.M
			m.G0 = v141 + v140
			v153 = base.I64_extend_i32_s(v131-v145)*int64(-1000000) + l0
			if base.Ui64(v153+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
				v169 = v153
				m.G0 = v8 + int32(48)
				return v169
			} else {
				if l1 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v188 = m.ExcPending
						if v188 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(393503), int32(0))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(484742), int32(6509), int32(30234))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v163 = base.B2i32(v153 < int64(-211813488000000000))
					if v153 < int64(-211813488000000000) {
						v164 = int32(-1)
					} else {
						v164 = int32(1)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v164
					if v153 < int64(-211813488000000000) {
						v168 = int64(-9223372036854775807 - 1)
					} else {
						v168 = int64(9223372036854775807)
					}
					v169 = v168
					m.G0 = v8 + int32(48)
					return v169
				}
			}
		}
	}
}
func F_timetztypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v14 = F_ArrayGetIntegerTypmods(m, v8, v5+int32(12))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(218214), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(487430), int32(65), int32(271477))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v37 = F_anytime_typmod_check(m, int32(1), v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v37
				}
			}
		}
	}
}
func F_tliOfPointInHistory(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	v3 = int32(0)
	if l1 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	return v54
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v11 = int32(0)
	if v11 < v8 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v14 = v8
	goto L7
L6:
	;
	v14 = v11
	goto L7
L7:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v19 = v3
	goto L8
L8:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15+v19<<(uint(int32(2))%32))))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v24)+8))
	if base.Ui64(v25) <= base.Ui64(l0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)+16))
	if base.Ui64(l0) <= base.Ui64(v27-int64(1)) {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v32 = v19 + int32(1)
	if v32 != v14 {
		v19 = v32
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	goto L9
L15:
	;
	return int32(0)
L16:
	;
	F_errmsg_internal(m, int32(111817), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(487631), int32(561), int32(12827))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tlist_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v9 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v16 = v3
	goto L8
L7:
	;
	return v22
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v16<<(uint(int32(2))%32))))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v24 = F_equal(m, l0, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	return int32(0)
L11:
	;
	if v24 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v29 = v16 + int32(1)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v29 < v30 {
		v16 = v29
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
}
func F_towlower(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v129 int32
	_ = v129
	v2 = int32(0)
	if base.Ui32(int32(131071)) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v129
L2:
	;
	v129 = l0
	goto L1
L3:
	;
	v12 = int32(255)
	v13 = l0 & v12
	v14 = int32(3)
	v15 = base.I32_div_u_s(v13, v14)
	v21 = int32(2)
	v25 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_consts[1343])))
	v26 = int32(8)
	v27 = int32(base.Ui32(l0) >> (uint(v26) % 32))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1344]))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30*int32(86)+v15)+uint32(_consts[1344]))))
	v41 = base.I32_rem_u_s(int32(base.Ui32(v25*v36)>>(uint(int32(11))%32)), int32(6))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1345]))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32((v41+v44)<<(uint(v21)%32))+uint32(_consts[1346])))
	v52 = v50 >> (uint(v26) % 32)
	v54 = v50 & v12
	if base.Ui32(v54) <= base.Ui32(int32(1)) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v129 = v52&(int32(0)-(v2^v54)) + l0
	goto L1
L5:
	;
	goto L6
L6:
	;
	v63 = v52 & int32(255)
	if v63 == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v70 = v63
	v71 = int32(base.Ui32(v52) >> (uint(int32(8)) % 32))
	goto L8
L8:
	;
	v77 = int32(1)
	v78 = int32(base.Ui32(v70) >> (uint(v77) % 32))
	v79 = v78 + v71
	v81 = v79 << (uint(v77) % 32)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1347]))))
	if v84 == v13 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1348]))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86<<(uint(int32(2))%32))+uint32(_consts[1346])))
	v93 = v91 & int32(255)
	if base.Ui32(v93) <= base.Ui32(int32(1)) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v107 = base.B2i32(base.Ui32(v13) < base.Ui32(v84))
	if base.Ui32(v13) < base.Ui32(v84) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v129 = (int32(0)-(v2^v93))&(v91>>(uint(int32(8))%32)) + l0
	goto L1
L14:
	;
	goto L15
L15:
	;
	goto L17
L17:
	;
	goto L18
L18:
	;
	v129 = int32(1) + l0
	goto L1
L19:
	;
	v108 = v71
	goto L21
L20:
	;
	v108 = v79
	goto L21
L21:
	;
	if base.Ui32(v13) < base.Ui32(v84) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v110 = v78
	goto L24
L23:
	;
	v110 = v70 - v78
	goto L24
L24:
	;
	if v110 != 0 {
		v70 = v110
		v71 = v108
		goto L8
	} else {
		goto L25
	}
L25:
	;
	goto L9
}
func F_transfer_first_span(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	v4 = l3
	v11 = l1 + int32(16)
	v14 = v11 + l2<<(uint(int32(2))%32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+1468))
	if v16 != v18 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return base.B2i32(v15 != int32(0))
L4:
	;
	v23 = F_LWLockAcquire(m, v17+int32(1476), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v37 = l0 + int32(12)
	v39 = int32(base.Ui32(v15) >> (uint(int32(27)) % 32))
	v42 = v37 + v39*int32(20)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	return int32(0)
L8:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v29+int32(1476))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L6
L11:
	;
	v47 = v43
	goto L13
L12:
	;
	v44 = F_get_segment_by_index(m, l0, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L14
	}
L13:
	;
	v48 = v47 + v15&int32(134217727)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v49
	if v49 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v47 = v46
	goto L13
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+1468))
	if v51 != v53 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v87 = v11 + v4<<(uint(int32(2))%32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v15
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v91 != 0 {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v58 = F_LWLockAcquire(m, v52+int32(1476), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v70 = int32(base.Ui32(v49) >> (uint(int32(27)) % 32))
	v73 = v37 + v70*int32(20)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	if v74 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v62+int32(1476))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v78 = v74
	goto L26
L25:
	;
	v75 = F_get_segment_by_index(m, l0, v70)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L7
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49&int32(134217727)+v78)+4)) = int32(0)
	goto L17
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v78 = v77
	goto L26
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+652))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+1468))
	if v92 != v94 {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+30)) = uint16(v4)
	goto L3
L31:
	;
	v99 = F_LWLockAcquire(m, v93+int32(1476), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v111 = int32(base.Ui32(v91) >> (uint(int32(27)) % 32))
	v114 = v37 + v111*int32(20)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v115 != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	F_check_for_freed_segments_locked(m, l0)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_LWLockRelease(m, v103+int32(1476))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v119 = v115
	goto L39
L38:
	;
	v116 = F_get_segment_by_index(m, l0, v111)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L7
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91&int32(134217727)+v119)+4)) = v15
	goto L30
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v119 = v118
	goto L39
}
func F_transformDistinctClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v5 = int32(0)
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v71 == int32(0) {
		v108 = v67
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v16 = v5
	v17 = v5
	goto L7
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v9 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v67 = v5
	goto L1
L6:
	;
	goto L5
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+v17<<(uint(int32(2))%32))))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v26 = F_get_sortgroupclause_tle(m, v24, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L9
	} else {
		goto L17
	}
L9:
	;
	return int32(0)
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+26)))
	if v30 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = F_copyObjectImpl(m, v24)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L8
L14:
	;
	v35 = F_lappend(m, v16, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v38 = v17 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v38 < v39 {
		v16 = v35
		v17 = v38
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v67 = v35
	goto L1
L17:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L18
	}
L18:
	;
	if l3 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = int32(72530)
	goto L21
L20:
	;
	v50 = int32(72804)
	goto L21
L21:
	;
	F_errmsg(m, v50, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v55 = F_exprLocation(m, v54)
	mBase = m.M
	F_parser_errposition(m, l0, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(487449), int32(3019), int32(351709))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	if v108 != 0 {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v74 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v75 <= v74 {
		v108 = v67
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v82 = v67
	v83 = v74
	goto L28
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v86+v83<<(uint(int32(2))%32))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+26)))
	if v91 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v108 = v99
	goto L25
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v96 = F_exprLocation(m, v95)
	mBase = m.M
	v97 = F_addTargetToGroupList(m, l0, v90, v82, v94, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	v99 = v82
	goto L32
L32:
	;
	v101 = v83 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v101 < v102 {
		v82 = v99
		v83 = v101
		goto L28
	} else {
		goto L34
	}
L33:
	;
	v99 = v97
	goto L32
L34:
	;
	goto L29
L35:
	;
	return v108
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	if l3 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v122 = int32(92167)
	goto L42
L41:
	;
	v122 = int32(268167)
	goto L42
L42:
	;
	F_errmsg(m, v122, int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(487449), int32(3050), int32(351709))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformReturningClause(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v18 = v16
	goto L3
L2:
	;
	v18 = int32(0)
	goto L3
L3:
	;
	if l2 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L22
	} else {
		goto L72
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L22
	} else {
		goto L67
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L22
	} else {
		goto L64
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L22
	} else {
		goto L59
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L22
	} else {
		goto L54
	}
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v19 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	m.G0 = v13 - int32(-64)
	return
L12:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v74 != 0 {
		goto L30
	} else {
		goto L31
	}
L13:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v22 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v31 = int32(0)
	goto L15
L15:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v31<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	switch v40 {
	case 0:
		goto L19
	case 1:
		goto L18
	default:
		goto L6
	}
L16:
	;
	goto L12
L17:
	;
	v48 = int32(0)
	v51 = F_refnameNamespaceItem(m, l0, v48, v47, int32(-1), v48)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	if v44 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+88))
	if v41 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v42
	v47 = v42
	goto L17
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v45
	v47 = v45
	goto L17
L22:
	;
	return
L23:
	;
	if v51 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v56 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v57 = int32(2)
	goto L27
L26:
	;
	v57 = int32(1)
	goto L27
L27:
	;
	F_addNSItemForReturning(m, l0, v53, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v61 = v31 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v61 < v62 {
		v31 = v61
		goto L15
	} else {
		goto L29
	}
L29:
	;
	goto L16
L30:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+92))
	if v87 != 0 {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v75 = int32(0)
	v79 = F_refnameNamespaceItem(m, l0, v75, int32(420458), int32(-1), v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L22
	} else {
		goto L32
	}
L32:
	;
	if v79 != 0 {
		goto L30
	} else {
		goto L33
	}
L33:
	;
	v81 = int32(420458)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+88)) = v81
	F_addNSItemForReturning(m, l0, v81, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	goto L30
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v104 = F_transformTargetList(m, l0, v103, l3)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L22
	} else {
		goto L40
	}
L36:
	;
	v88 = int32(0)
	v92 = F_refnameNamespaceItem(m, l0, v88, int32(31247), int32(-1), v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	if v92 != 0 {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v94 = int32(31247)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v94
	F_addNSItemForReturning(m, l0, v94, int32(2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L22
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+96)) = v104
	if v104 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_markTargetListOrigins(m, l0, v104)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+85)))
	if v111 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	F_resolveTargetListUnknowns(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L22
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v118 = int32(0)
	if v117 == v118 {
		v126 = v118
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v126
	goto L11
L48:
	;
	goto L47
L49:
	;
	if v18 <= int32(0) {
		v126 = v118
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v18 < v123 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v18
	goto L53
L52:
	;
	goto L53
L53:
	;
	v126 = v117
	goto L48
L54:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L22
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(531017)
	F_errmsg(m, int32(159536), v11+int32(-32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L22
	} else {
		goto L56
	}
L56:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	F_parser_errposition(m, l0, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L22
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(487147), int32(2669), int32(351801))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L22
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(504665)
	F_errmsg(m, int32(159536), v11+int32(-16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L22
	} else {
		goto L61
	}
L61:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	F_parser_errposition(m, l0, v178)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(487147), int32(2679), int32(351801))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L22
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
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v190
	F_errmsg_internal(m, int32(472130), v13)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(487147), int32(2684), int32(351801))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L22
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(33845380))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v207
	F_errmsg(m, int32(406502), v11+int32(-48))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L22
	} else {
		goto L69
	}
L69:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	F_parser_errposition(m, l0, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L22
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(487147), int32(2692), int32(351801))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L22
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L22
	} else {
		goto L73
	}
L73:
	;
	F_errmsg(m, int32(268213), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L22
	} else {
		goto L74
	}
L74:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)))
	v236 = F_exprLocation(m, v235)
	mBase = m.M
	F_parser_errposition(m, l0, v236)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L22
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(487147), int32(2740), int32(351801))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L22
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_traverse_lacons(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	F_check_stack_depth(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v9+l1<<(uint(int32(2))%32))))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	if v14 != int32(65535) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = v13
	v22 = v14
	goto L6
L4:
	;
	goto L5
L5:
	;
	return
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.I32_extend16_s(v22) < v23 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v43 = v18 + int32(8)
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43))))
	if v44 != int32(65535) {
		v18 = v43
		v22 = v44
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v26 + int32(1)
	if l4 <= v26 {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	F_traverse_lacons(m, l0, v38, l2, l3, l4)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v33 = l3 + v26<<(uint(int32(3))%32)
	v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18))))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v36
	goto L8
L13:
	;
	goto L8
L14:
	;
	goto L7
}
func F_trivial_subqueryscan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	v9 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	switch v10 - v9 {
	case 0:
		v95 = v9
		goto L1
	case 1:
		goto L3
	default:
		goto L2
	}
L1:
	;
	return v95
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = int32(2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return int32(0)
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v23 = v22
	goto L9
L8:
	;
	v23 = int32(0)
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = v26
	goto L12
L11:
	;
	v27 = v20
	goto L12
L12:
	;
	if v27 != v23 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v34 = int32(0)
	v39 = int32(1)
	goto L16
L16:
	;
	v41 = int32(0)
	if v21 == v41 {
		v51 = v41
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v89 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v89
	v95 = v89
	goto L1
L18:
	;
	if v25 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v45 <= v34 {
		v51 = int32(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v51 = v47 + v34<<(uint(int32(2))%32)
	goto L18
L21:
	;
	goto L17
L22:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v54 <= v34 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if v51 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v61 = v58 + v34<<(uint(int32(2))%32)
	if v61 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+26)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+26)))
	if v65 != v67 {
		v95 = v20
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v69 == int32(0) {
		v95 = v20
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	switch v72 - int32(6) {
	case 0:
		goto L30
	case 1:
		goto L29
	default:
		v95 = v20
		goto L1
	}
L28:
	;
	v84 = int32(1)
	v34 = v34 + v84
	v39 = v39 + v84
	goto L16
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v78 = F_equal(m, v69, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+8)))
	if v39 == v75 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v95 = v20
	goto L1
L32:
	;
	return int32(0)
L33:
	;
	if v78 == int32(0) {
		v95 = v20
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L28
}
func F_truncate_useless_pathkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v517 int32
	_ = v517
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v559 int32
	_ = v559
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	v4 = int32(0)
	if l2 == v4 {
		v378 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if l2 == v379 {
		goto L93
	} else {
		goto L94
	}
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v14 <= int32(0) {
		v378 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = v4
	goto L4
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v27<<(uint(int32(2))%32))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	if v33 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v378 = v365
	goto L1
L6:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v93 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v79 != v80 {
		v378 = v27
		goto L1
	} else {
		goto L19
	}
L8:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v76 == int32(1) {
		goto L6
	} else {
		goto L18
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v36 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v45 = int32(0)
	goto L11
L11:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v40+v45<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v57 == v39 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L8
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v59 == v60 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v63 = v45 + int32(1)
	if v36 != v63 {
		v45 = v63
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	goto L12
L18:
	;
	v378 = v27
	goto L1
L19:
	;
	goto L6
L20:
	;
	v365 = v27 + int32(1)
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v365 < v366 {
		v27 = v365
		goto L4
	} else {
		goto L90
	}
L21:
	;
	v96 = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+40)))
	if v98 != 0 {
		v256 = v96
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v271 == int32(0) {
		v378 = v27
		goto L1
	} else {
		goto L69
	}
L24:
	;
	if v256 != 0 {
		goto L20
	} else {
		goto L68
	}
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	if v99 == int32(0) {
		v256 = v96
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v102 < int32(2) {
		v256 = v96
		goto L24
	} else {
		goto L27
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v97)+36))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v108 = v106 - int32(2)
	if base.Ui32(v108) <= base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v108<<(uint(int32(2))%32))+uint32(_consts[593])))
	v117 = v115
	goto L30
L29:
	;
	v117 = int32(8)
	goto L30
L30:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l1+v117)))
	v120 = int32(0)
	if v105 == v120 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v173 != 0 {
		v256 = v96
		goto L24
	} else {
		goto L45
	}
L32:
	;
	v173 = int32(1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	if v119 == int32(0) {
		v164 = v120
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v173 = v164
	goto L31
L36:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v130 < v129 {
		v164 = v120
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v132 = int32(1)
	if v129 <= v132 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v135 = v132
	goto L40
L39:
	;
	v135 = v129
	goto L40
L40:
	;
	v136 = int32(8)
	v141 = int32(0)
	goto L41
L41:
	;
	v148 = v141 << (uint(int32(2)) % 32)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v105+v136+v148)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148+(v119+v136))))
	v155 = v150 & (v152 ^ int32(-1))
	v157 = base.B2i32(v155 == int32(0))
	if v155 != 0 {
		v164 = v157
		goto L35
	} else {
		goto L43
	}
L42:
	;
	v164 = v157
	goto L35
L43:
	;
	v159 = v141 + int32(1)
	if v159 != v135 {
		v141 = v159
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	if v174 == int32(0) {
		v256 = v96
		goto L24
	} else {
		goto L46
	}
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v177 <= int32(0) {
		v256 = v96
		goto L24
	} else {
		goto L47
	}
L47:
	;
	v187 = v96
	goto L48
L48:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191+v187<<(uint(int32(2))%32))))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	v197 = int32(0)
	if v196 == v197 {
		v238 = v197
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v256 = v238 ^ int32(1)
	goto L24
L50:
	;
	if v238 != 0 {
		goto L64
	} else {
		goto L65
	}
L51:
	;
	goto L50
L52:
	;
	if v119 == int32(0) {
		v238 = v197
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v206 < v207 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v209 = v206
	goto L56
L55:
	;
	v209 = v207
	goto L56
L56:
	;
	if v209 <= int32(1) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v212 = int32(1)
	goto L59
L58:
	;
	v212 = v209
	goto L59
L59:
	;
	v213 = int32(8)
	v218 = int32(0)
	goto L60
L60:
	;
	v225 = v218 << (uint(int32(2)) % 32)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v119+v213+v225)))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+(v196+v213))))
	v230 = v227 & v229
	v232 = base.B2i32(v230 != int32(0))
	if v230 != 0 {
		v238 = v232
		goto L51
	} else {
		goto L62
	}
L61:
	;
	v238 = v232
	goto L51
L62:
	;
	v234 = v218 + int32(1)
	if v234 != v212 {
		v218 = v234
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v243 = v187 + int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	if v243 < v244 {
		v187 = v243
		goto L48
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	goto L49
L67:
	;
	goto L66
L68:
	;
	goto L23
L69:
	;
	v274 = int32(0)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	if v275 <= v274 {
		v378 = v27
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v281 = v275
	v284 = v274
	goto L71
L71:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v289+v284<<(uint(int32(2))%32))))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+96))
	if v294 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v378 = v27
	goto L1
L73:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v293)+100))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+56))
	if v296 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v342 = v281
	goto L75
L75:
	;
	v351 = v284 + int32(1)
	if v351 < v342 {
		v281 = v342
		v284 = v351
		goto L71
	} else {
		goto L89
	}
L76:
	;
	v300 = v296
	goto L79
L77:
	;
	v317 = v295
	goto L78
L78:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v293)+104))
	v325 = v321
	goto L82
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+100)) = v300
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v300)+56))
	if v309 != 0 {
		v300 = v309
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v317 = v300
	goto L78
L81:
	;
	goto L80
L82:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)+56))
	if v333 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	if v335 == v317 {
		goto L20
	} else {
		goto L87
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+104)) = v333
	v325 = v333
	goto L82
L85:
	;
	goto L86
L86:
	;
	goto L83
L87:
	;
	if v325 == v335 {
		goto L20
	} else {
		goto L88
	}
L88:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v271)+4))
	v342 = v338
	goto L75
L89:
	;
	goto L72
L90:
	;
	goto L5
L91:
	;
	if v378 < v431 {
		goto L117
	} else {
		goto L118
	}
L92:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	v431 = v426
	goto L91
L93:
	;
	if v379 != 0 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v382 = int32(0)
	if l2 == v382 {
		v431 = v382
		goto L91
	} else {
		goto L97
	}
L96:
	;
	v431 = int32(0)
	goto L91
L97:
	;
	if v379 == int32(0) {
		v431 = v382
		goto L91
	} else {
		goto L98
	}
L98:
	;
	v387 = int32(0)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v387 < v388 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v392 = v388
	goto L101
L100:
	;
	v392 = v387
	goto L101
L101:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	v397 = v387
	goto L102
L102:
	;
	if v397 < v393 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v379)+12))
	v411 = v407 + v397<<(uint(int32(2))%32)
	goto L106
L105:
	;
	v411 = int32(0)
	goto L106
L106:
	;
	if v397 == v392 {
		v431 = v392
		goto L91
	} else {
		goto L107
	}
L107:
	;
	if v411 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v431 = v397
	goto L91
L109:
	;
	goto L110
L110:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v418 = v415 + v397<<(uint(int32(2))%32)
	if v418 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v431 = v397
	goto L91
L112:
	;
	goto L113
L113:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v418)))
	if v421 != v422 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v431 = v397
	goto L91
L115:
	;
	v397 = v397 + int32(1)
	goto L102
L117:
	;
	v439 = v431
	goto L119
L118:
	;
	v439 = v378
	goto L119
L119:
	;
	v440 = int32(0)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v442 == v440 {
		v517 = v440
		goto L124
	} else {
		goto L125
	}
L120:
	;
	if v626 < v625 {
		goto L175
	} else {
		goto L176
	}
L121:
	;
	v618 = int32(0)
	if v618 < v439 {
		goto L172
	} else {
		goto L173
	}
L122:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v553 <= int32(0) {
		v625 = v440
		v626 = v545
		goto L120
	} else {
		goto L154
	}
L123:
	;
	v532 = int32(0)
	if v532 < v439 {
		goto L150
	} else {
		goto L151
	}
L124:
	;
	if v439 < v517 {
		goto L145
	} else {
		goto L146
	}
L125:
	;
	if l2 == int32(0) {
		goto L121
	} else {
		goto L126
	}
L126:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v447 <= int32(0) {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v455 = v440
	goto L128
L128:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v462+v455<<(uint(int32(2))%32))))
	v467 = int32(0)
	if v461 == v467 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v517 = v509
	goto L124
L130:
	;
	if v505 == int32(0) {
		v517 = v455
		goto L124
	} else {
		goto L143
	}
L131:
	;
	v505 = int32(0)
	goto L130
L132:
	;
	goto L133
L133:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v461)+4))
	if v473 <= int32(0) {
		v498 = v467
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v505 = v498
	goto L130
L135:
	;
	v476 = int32(0)
	if v476 < v473 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v479 = v473
	goto L138
L137:
	;
	v479 = v476
	goto L138
L138:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v482 = int32(0)
	goto L139
L139:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v480+v482<<(uint(int32(2))%32))))
	v491 = base.B2i32(v490 == v466)
	if v490 == v466 {
		v498 = v491
		goto L134
	} else {
		goto L141
	}
L140:
	;
	v498 = v491
	goto L134
L141:
	;
	v493 = v482 + int32(1)
	if v493 != v479 {
		v482 = v493
		goto L139
	} else {
		goto L142
	}
L142:
	;
	goto L140
L143:
	;
	v509 = v455 + int32(1)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v509 < v510 {
		v455 = v509
		goto L128
	} else {
		goto L144
	}
L144:
	;
	goto L129
L145:
	;
	v524 = v517
	goto L147
L146:
	;
	v524 = v439
	goto L147
L147:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v525 == int32(0) {
		v625 = v440
		v626 = v524
		goto L120
	} else {
		goto L148
	}
L148:
	;
	if l2 == int32(0) {
		v625 = v440
		v626 = v524
		goto L120
	} else {
		goto L149
	}
L149:
	;
	v545 = v524
	v552 = l0 + int32(172)
	goto L122
L150:
	;
	v535 = v439
	goto L152
L151:
	;
	v535 = v532
	goto L152
L152:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	if v536 == int32(0) {
		v625 = v440
		v626 = v535
		goto L120
	} else {
		goto L153
	}
L153:
	;
	v545 = v535
	v552 = l0 + int32(172)
	goto L122
L154:
	;
	v559 = v440
	goto L155
L155:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v568+v559<<(uint(int32(2))%32))))
	v573 = int32(0)
	if v567 == v573 {
		goto L158
	} else {
		goto L159
	}
L156:
	;
	v625 = v615
	v626 = v545
	goto L120
L157:
	;
	if v611 == int32(0) {
		v625 = v559
		v626 = v545
		goto L120
	} else {
		goto L170
	}
L158:
	;
	v611 = int32(0)
	goto L157
L159:
	;
	goto L160
L160:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v567)+4))
	if v579 <= int32(0) {
		v604 = v573
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v611 = v604
	goto L157
L162:
	;
	v582 = int32(0)
	if v582 < v579 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v585 = v579
	goto L165
L164:
	;
	v585 = v582
	goto L165
L165:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v567)+12))
	v588 = int32(0)
	goto L166
L166:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v586+v588<<(uint(int32(2))%32))))
	v597 = base.B2i32(v596 == v572)
	if v596 == v572 {
		v604 = v597
		goto L161
	} else {
		goto L168
	}
L167:
	;
	v604 = v597
	goto L161
L168:
	;
	v599 = v588 + int32(1)
	if v599 != v585 {
		v588 = v599
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	v615 = v559 + int32(1)
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v615 < v616 {
		v559 = v615
		goto L155
	} else {
		goto L171
	}
L171:
	;
	goto L156
L172:
	;
	v621 = v439
	goto L174
L173:
	;
	v621 = v618
	goto L174
L174:
	;
	v625 = v440
	v626 = v621
	goto L120
L175:
	;
	v634 = v625
	goto L177
L176:
	;
	v634 = v626
	goto L177
L177:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if l2 == v635 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	if v634 < v687 {
		goto L204
	} else {
		goto L205
	}
L179:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	v687 = v682
	goto L178
L180:
	;
	if v635 != 0 {
		goto L179
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v638 = int32(0)
	if l2 == v638 {
		v687 = v638
		goto L178
	} else {
		goto L184
	}
L183:
	;
	v687 = int32(0)
	goto L178
L184:
	;
	if v635 == int32(0) {
		v687 = v638
		goto L178
	} else {
		goto L185
	}
L185:
	;
	v643 = int32(0)
	v644 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v643 < v644 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v648 = v644
	goto L188
L187:
	;
	v648 = v643
	goto L188
L188:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	v653 = v643
	goto L189
L189:
	;
	if v653 < v649 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
	v667 = v663 + v653<<(uint(int32(2))%32)
	goto L193
L192:
	;
	v667 = int32(0)
	goto L193
L193:
	;
	if v653 == v648 {
		v687 = v648
		goto L178
	} else {
		goto L194
	}
L194:
	;
	if v667 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v687 = v653
	goto L178
L196:
	;
	goto L197
L197:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v674 = v671 + v653<<(uint(int32(2))%32)
	if v674 == int32(0) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v687 = v653
	goto L178
L199:
	;
	goto L200
L200:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v667)))
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v674)))
	if v677 != v678 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	v687 = v653
	goto L178
L202:
	;
	v653 = v653 + int32(1)
	goto L189
L204:
	;
	v695 = v687
	goto L206
L205:
	;
	v695 = v634
	goto L206
L206:
	;
	if v695 == int32(0) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	return int32(0)
L208:
	;
	goto L209
L209:
	;
	if l2 != 0 {
		goto L211
	} else {
		goto L212
	}
L210:
	;
	return v706
L211:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v695 == v700 {
		v706 = l2
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v702 = F_list_copy_head(m, l2, v695)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	goto L213
L215:
	;
	return int32(0)
L216:
	;
	v706 = v702
	goto L210
}
func F_tsearch_readline_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v10 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v15
			F_errcontext_msg(m, int32(699037), v8+int32(16))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v15
			F_errcontext_msg(m, int32(687176), v8)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				m.G0 = v8 + int32(32)
				return
			}
		}
	}
}
func F_tsearch_readline_end(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v4 != v3 {
			F_pfree(m, v3)
			mBase = m.M
			v7 = m.ExcPending
			if v7 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_pfree(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v14 = F_FreeFile(m, v13)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return
					} else {
						v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						*(*int32)(unsafe.Add(mBase, _consts[49])) = v17
						return
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pfree(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v14 = F_FreeFile(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					*(*int32)(unsafe.Add(mBase, _consts[49])) = v17
					return
				}
			}
		}
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_pfree(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v14 = F_FreeFile(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				*(*int32)(unsafe.Add(mBase, _consts[49])) = v17
				return
			}
		}
	}
}
func F_tsqueryrecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v250 int32
	_ = v250
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v311 int32
	_ = v311
	var v320 int32
	_ = v320
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	v2 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_pq_getmsgint(m, v23, int32(4))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L8
	} else {
		goto L106
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L8
	} else {
		goto L103
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L8
	} else {
		goto L100
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L8
	} else {
		goto L97
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L8
	} else {
		goto L94
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L8
	} else {
		goto L91
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L8
	} else {
		goto L88
	}
L8:
	;
	return int32(0)
L9:
	;
	if base.Ui32(v25) < base.Ui32(int32(89478486)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = F_palloc(m, v25<<(uint(int32(2))%32))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L8
	} else {
		goto L85
	}
L13:
	;
	v38 = v25*int32(12) + int32(8)
	v39 = F_palloc0(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v25
	if v25 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v48 = v39 + int32(8)
	v53 = v2
	v55 = v2
	goto L18
L16:
	;
	v320 = v2
	goto L17
L17:
	;
	v331 = v320 + v38
	v332 = F_repalloc(m, v39, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L8
	} else {
		goto L68
	}
L18:
	;
	v65 = F_pq_getmsgint(m, v23, int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L20
	}
L19:
	;
	v320 = v297
	goto L17
L20:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v48))) = uint8(v65)
	switch v65&int32(255) - int32(1) {
	case 0:
		goto L23
	case 1:
		goto L22
	default:
		goto L2
	}
L21:
	;
	v311 = v55 + int32(1)
	if v311 != v25 {
		v48 = v48 + int32(12)
		v53 = v297
		v55 = v311
		goto L18
	} else {
		goto L67
	}
L22:
	;
	v270 = F_pq_getmsgint(m, v23, int32(1))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L59
	}
L23:
	;
	v73 = F_pq_getmsgint(m, v23, int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v76 = F_pq_getmsgint(m, v23, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v78 = F_pq_getmsgstring(m, v23)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	if v78&int32(3) == int32(0) {
		v103 = v78
		goto L29
	} else {
		goto L30
	}
L27:
	;
	if v73&int32(240) != 0 {
		goto L7
	} else {
		goto L44
	}
L28:
	;
	v136 = v128 - v78
	goto L27
L29:
	;
	v107 = v103
	goto L38
L30:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v87 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v136 = int32(0)
	goto L27
L32:
	;
	goto L33
L33:
	;
	v92 = v78
	goto L34
L34:
	;
	v96 = v92 + int32(1)
	if v96&int32(3) == int32(0) {
		v103 = v96
		goto L29
	} else {
		goto L36
	}
L35:
	;
	v128 = v96
	goto L28
L36:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	if v101 != 0 {
		v92 = v96
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v116 = int32(-2139062144)
	if (int32(16843008)-v113|v113)&v116 == v116 {
		v107 = v107 + int32(4)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v122 = v107
	goto L41
L40:
	;
	goto L39
L41:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v126 != 0 {
		v122 = v122 + int32(1)
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v128 = v122
	goto L28
L43:
	;
	goto L42
L44:
	;
	if base.Ui32(int32(2048)) <= base.Ui32(v136) {
		goto L6
	} else {
		goto L45
	}
L45:
	;
	if int32(1048575) < v53 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if v136 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v143 = int32(-1)
	if v136 != int32(1) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v250 = int32(0)
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v250
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v136 | v53<<(uint(int32(12))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)) = uint8(base.B2i32(v76&int32(255) != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v33+v55<<(uint(int32(2))%32)))) = v78
	v297 = v136 + v53 + int32(1)
	goto L21
L50:
	;
	v149 = v78
	v150 = v143
	v157 = int32(0)
	goto L53
L51:
	;
	v196 = v78
	v197 = v143
	goto L52
L52:
	;
	if v136&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	v169 = int32(24)
	v172 = int32(2)
	v176 = *(*int32)(unsafe.Add(mBase, uint32((v168^int32(base.Ui32(v150)>>(uint(v169)%32)))<<(uint(v172)%32))+uint32(_consts[1047])))
	v177 = int32(8)
	v179 = v176 ^ v150<<(uint(v177)%32)
	v187 = *(*int32)(unsafe.Add(mBase, uint32((v167^int32(base.Ui32(v179)>>(uint(v169)%32)))<<(uint(v172)%32))+uint32(_consts[1047])))
	v190 = v187 ^ v179<<(uint(v177)%32)
	v192 = v149 + v172
	v194 = v157 + v172
	if v194 != v136&int32(2046) {
		v149 = v192
		v150 = v190
		v157 = v194
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v196 = v192
	v197 = v190
	goto L52
L55:
	;
	goto L54
L56:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v224 = *(*int32)(unsafe.Add(mBase, uint32((v216^int32(base.Ui32(v197)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_consts[1047])))
	v228 = v224 ^ v197<<(uint(int32(8))%32)
	goto L58
L57:
	;
	v228 = v197
	goto L58
L58:
	;
	v250 = v228 ^ int32(-1)
	goto L49
L59:
	;
	v273 = v270 << (uint(int32(24)) % 32)
	if v270&int32(253) == int32(1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	if v55 == v25-int32(1) {
		goto L3
	} else {
		goto L64
	}
L61:
	;
	if v273 == int32(33554432) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if v273 != int32(67108864) {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v270)
	if v273 != int32(67108864) {
		v297 = v53
		goto L21
	} else {
		goto L65
	}
L65:
	;
	v287 = F_pq_getmsgint(m, v23, int32(2))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L66
	}
L66:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)) = uint16(v287)
	v297 = v53
	goto L21
L67:
	;
	goto L19
L68:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v332)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = int32(0)
	v338 = v332 + int32(8)
	F_findoprnd_recurse(m, v338, v21+int32(28), v25, v21+int32(27))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	if v345 != v25 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v25 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v351 = v338
	v352 = int32(0)
	v359 = v338 + v334*int32(12)
	goto L74
L72:
	;
	goto L73
L73:
	;
	F_pfree(m, v33)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L8
	} else {
		goto L84
	}
L74:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
	if v369 == int32(1) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	goto L73
L76:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v33+v352<<(uint(int32(2))%32))))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v351)+8))
	v380 = v376&int32(4095) + int32(1)
	if v380 != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v389 = v359
	goto L78
L78:
	;
	v393 = v352 + int32(1)
	if v393 != v25 {
		v351 = v351 + int32(12)
		v352 = v393
		v359 = v389
		goto L74
	} else {
		goto L83
	}
L79:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v351)+8))
	v389 = v382 + v383&int32(4095) + int32(1)
	goto L78
L80:
	;
	v381 = F__emscripten_memcpy_bulkmem(m, v359, v375, v380)
	mBase = m.M
	v382 = v381
	goto L82
L81:
	;
	v382 = v359
	goto L82
L82:
	;
	goto L79
L83:
	;
	goto L75
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v331 << (uint(int32(2)) % 32)
	m.G0 = v21 + int32(32)
	return v332
L85:
	;
	F_errmsg_internal(m, int32(14984), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L8
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(481320), int32(1241), int32(34839))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(232966), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(481320), int32(1274), int32(34839))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errmsg_internal(m, int32(320790), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(481320), int32(1277), int32(34839))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(452434), int32(0))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L8
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(481320), int32(1280), int32(34839))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L8
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
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v273 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(465572), v21+int32(16))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L8
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(481320), int32(1309), int32(34839))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(418132), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(481320), int32(1311), int32(34839))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L8
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
	v508 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48))))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v508
	F_errmsg_internal(m, int32(474909), v21)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	F_errfinish(m, int32(481320), int32(1318), int32(34839))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L8
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
	F_errmsg_internal(m, int32(167644), int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(481320), int32(793), int32(415683))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tsquerysend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v11+int32(16))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_enlargeStringInfo(m, v11+int32(16), int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v29 = int32(24)
	v31 = int32(65280)
	v33 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v26+v27))) = v20<<(uint(v29)%32) | v20&v31<<(uint(v33)%32) | (int32(base.Ui32(v20)>>(uint(v33)%32))&v31 | int32(base.Ui32(v20)>>(uint(v29)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v26 + int32(4)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v48 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v53 = v13 + int32(8)
	v56 = v53
	v60 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v180 != v13 {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	F_enlargeStringInfo(m, v11+int32(16), int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v68+v69))) = uint8(v62)
	v72 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v68 + v72
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	switch v75 - v72 {
	case 0:
		goto L11
	case 1:
		goto L13
	default:
		goto L12
	}
L10:
	;
	v169 = v60 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v169 < v170 {
		v56 = v56 + int32(12)
		v60 = v169
		goto L7
	} else {
		goto L23
	}
L11:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	F_enlargeStringInfo(m, v11+int32(16), int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L20
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L17
	}
L13:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	F_enlargeStringInfo(m, v11+int32(16), int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v84+v85))) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v84 + int32(1)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v91 != int32(4) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+2)))
	F_enlargeStringInfo(m, v11+int32(16), int32(2))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v103 = int32(8)
	v107 = v94<<(uint(v103)%32) | int32(base.Ui32(v94)>>(uint(v103)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v100+v101))) = uint16(v107)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v100 + int32(2)
	goto L10
L17:
	;
	v116 = int32(*(*int8)(unsafe.Add(mBase, uint32(v56))))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v116
	F_errmsg_internal(m, int32(474909), v11)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(481320), int32(1215), int32(416308))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v132+v133))) = uint8(v126)
	v136 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v132 + v136
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+2)))
	F_enlargeStringInfo(m, v11+int32(16), v136)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v145+v146))) = uint8(v139)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v145 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v155 = int32(12)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	F_pq_sendstring(m, v11+int32(16), v53+v154*v155+int32(base.Ui32(v158)>>(uint(v155)%32)))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L10
L23:
	;
	goto L8
L24:
	;
	F_pfree(m, v13)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v185 = v11 + int32(16)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v187))) = v188 << (uint(int32(2)) % 32)
	goto L28
L27:
	;
	goto L26
L28:
	;
	m.G0 = v11 + int32(32)
	return v187
}
func F_tsquerytree(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
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
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 == v2 {
		v18 = F_palloc(m, int32(4))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(16)
			v104 = v18
			m.G0 = v11 + int32(32)
			return v104
		}
	} else {
		v24 = int32(8)
		v26 = m.G0
		v28 = v26 - int32(16)
		m.G0 = v28
		v31 = v13 + v24
		v32 = F_maketree(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			v34 = F_clean_NOT_intree(m, v32)
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = int64(16)
				v38 = int32(0)
				if v34 == v38 {
					v60 = v2
					v61 = v38
					*(*int32)(unsafe.Add(mBase, uint32(v11+v24))) = v60
					m.G0 = v28 + int32(16)
					if v61 == int32(0) {
						v69 = F_cstring_to_text(m, int32(510645))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							v99 = v69
							v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v13 == v100 {
								v104 = v99
								m.G0 = v11 + int32(32)
								return v104
							} else {
								F_pfree(m, v13)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									v104 = v99
									m.G0 = v11 + int32(32)
									return v104
								}
							}
						}
					} else {
						v71 = int32(32)
						*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v61
						v75 = F_palloc(m, v71)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v75
							*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v75
							v79 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v75))) = uint8(v79)
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
							v82 = int32(12)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31 + v81*v82
							F_infix_1(m, v11+v82, int32(-1), v79)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
								v95 = F_cstring_to_text_with_len(m, v92, v93-v92)
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v61)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										v99 = v95
										v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v13 == v100 {
											v104 = v99
											m.G0 = v11 + int32(32)
											return v104
										} else {
											F_pfree(m, v13)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v104 = v99
												m.G0 = v11 + int32(32)
												return v104
											}
										}
									}
								}
							}
						}
					}
				} else {
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
					v44 = int32(1)
					if base.Ui32(v44) < base.Ui32((v43-v44)&int32(255)) {
						v60 = v2
						v61 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v11+v24))) = v60
						m.G0 = v28 + int32(16)
						if v61 == int32(0) {
							v69 = F_cstring_to_text(m, int32(510645))
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								v99 = v69
								v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v13 == v100 {
									v104 = v99
									m.G0 = v11 + int32(32)
									return v104
								} else {
									F_pfree(m, v13)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										v104 = v99
										m.G0 = v11 + int32(32)
										return v104
									}
								}
							}
						} else {
							v71 = int32(32)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v61
							v75 = F_palloc(m, v71)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v75
								*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v75
								v79 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v75))) = uint8(v79)
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v82 = int32(12)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31 + v81*v82
								F_infix_1(m, v11+v82, int32(-1), v79)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
									v95 = F_cstring_to_text_with_len(m, v92, v93-v92)
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										F_pfree(m, v61)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											v99 = v95
											v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v13 == v100 {
												v104 = v99
												m.G0 = v11 + int32(32)
												return v104
											} else {
												F_pfree(m, v13)
												mBase = m.M
												v103 = m.ExcPending
												if v103 != 0 {
													return int32(0)
												} else {
													v104 = v99
													m.G0 = v11 + int32(32)
													return v104
												}
											}
										}
									}
								}
							}
						}
					} else {
						v51 = F_palloc(m, int32(192))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v51
							F_plainnode(m, v28+int32(4), v34)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
								v60 = v58
								v61 = v59
								*(*int32)(unsafe.Add(mBase, uint32(v11+v24))) = v60
								m.G0 = v28 + int32(16)
								if v61 == int32(0) {
									v69 = F_cstring_to_text(m, int32(510645))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v99 = v69
										v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v13 == v100 {
											v104 = v99
											m.G0 = v11 + int32(32)
											return v104
										} else {
											F_pfree(m, v13)
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v104 = v99
												m.G0 = v11 + int32(32)
												return v104
											}
										}
									}
								} else {
									v71 = int32(32)
									*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v71
									*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v61
									v75 = F_palloc(m, v71)
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v75
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v75
										v79 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v75))) = uint8(v79)
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v82 = int32(12)
										*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31 + v81*v82
										F_infix_1(m, v11+v82, int32(-1), v79)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
											v95 = F_cstring_to_text_with_len(m, v92, v93-v92)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v61)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return int32(0)
												} else {
													v99 = v95
													v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
													if v13 == v100 {
														v104 = v99
														m.G0 = v11 + int32(32)
														return v104
													} else {
														F_pfree(m, v13)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v104 = v99
															m.G0 = v11 + int32(32)
															return v104
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
func F_tupledesc_match(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 == v15 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L11
	} else {
		goto L32
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L11
	} else {
		goto L25
	}
L3:
	;
	if int32(0) < v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L11
	} else {
		goto L20
	}
L6:
	;
	v19 = int32(20)
	v26 = v14
	v27 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	m.G0 = v12 + int32(48)
	return
L9:
	;
	v33 = v27 * int32(100)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v35 = int32(4)
	v38 = v33 + (l1 + v19 + v34<<(uint(v35)%32))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	v43 = l0 + v19 + v26<<(uint(v35)%32) + v33
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
	v45 = F_IsBinaryCoercible(m, v39, v44)
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
	if v45 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+91)))
	if v49 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v59 = v27 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v59 < v60 {
		v26 = v60
		v27 = v59
		goto L9
	} else {
		goto L19
	}
L16:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+72)))
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+72)))
	if v52 != v53 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+83)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+83)))
	if v55 != v56 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	goto L10
L20:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(318094), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v85
	F_errdetail_plural(m, int32(624486), int32(624372), v85, v12+int32(32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(488903), int32(954), int32(317785))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L11
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L11
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(318094), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L11
	} else {
		goto L27
	}
L27:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	v112 = F_format_type_be(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L11
	} else {
		goto L28
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
	v115 = F_format_type_be(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v115
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v27 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v112
	F_errdetail(m, int32(575910), v12+int32(16))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(488903), int32(970), int32(317785))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L11
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(318094), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v27 + int32(1)
	F_errdetail(m, int32(625232), v12)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(488903), int32(978), int32(317785))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_tzparse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v492 int32
	_ = v492
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v745 int32
	_ = v745
	var v757 int32
	_ = v757
	var v776 int32
	_ = v776
	var v793 int64
	_ = v793
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v807 int32
	_ = v807
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v818 int64
	_ = v818
	var v821 int64
	_ = v821
	var v825 int64
	_ = v825
	var v826 int64
	_ = v826
	var v838 int32
	_ = v838
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v929 int64
	_ = v929
	var v934 int32
	_ = v934
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v963 int32
	_ = v963
	var v979 int32
	_ = v979
	var v988 int32
	_ = v988
	var v998 int32
	_ = v998
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1029 int32
	_ = v1029
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int64
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1055 int32
	_ = v1055
	var v1061 int32
	_ = v1061
	var v1074 int64
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1127 int32
	_ = v1127
	var v1133 int32
	_ = v1133
	var v1140 int32
	_ = v1140
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1202 int32
	_ = v1202
	var v1208 int32
	_ = v1208
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1221 int32
	_ = v1221
	var v1225 int32
	_ = v1225
	var v1228 int32
	_ = v1228
	var v1234 int32
	_ = v1234
	var v1238 int32
	_ = v1238
	var v1241 int32
	_ = v1241
	var v1242 int32
	_ = v1242
	var v1246 int32
	_ = v1246
	var v1252 int32
	_ = v1252
	var v1258 int32
	_ = v1258
	var v1262 int32
	_ = v1262
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1271 int32
	_ = v1271
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1290 int32
	_ = v1290
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1311 int32
	_ = v1311
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1335 int32
	_ = v1335
	var v1337 int32
	_ = v1337
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1371 int32
	_ = v1371
	var v1377 int32
	_ = v1377
	var v1384 int32
	_ = v1384
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1426 int32
	_ = v1426
	var v1435 int32
	_ = v1435
	var v1438 int32
	_ = v1438
	var v1440 int32
	_ = v1440
	var v1446 int32
	_ = v1446
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1456 int32
	_ = v1456
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1490 int32
	_ = v1490
	var v1496 int32
	_ = v1496
	var v1502 int32
	_ = v1502
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1515 int32
	_ = v1515
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1534 int32
	_ = v1534
	var v1543 int32
	_ = v1543
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1550 int32
	_ = v1550
	var v1555 int32
	_ = v1555
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1571 int32
	_ = v1571
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1595 int32
	_ = v1595
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1605 int64
	_ = v1605
	var v1608 int64
	_ = v1608
	var v1612 int64
	_ = v1612
	var v1620 int32
	_ = v1620
	var v1624 int32
	_ = v1624
	var v1626 int32
	_ = v1626
	var v1630 int64
	_ = v1630
	var v1633 int64
	_ = v1633
	var v1637 int64
	_ = v1637
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1658 int64
	_ = v1658
	var v1661 int64
	_ = v1661
	var v1665 int64
	_ = v1665
	var v1669 int32
	_ = v1669
	var v1672 int32
	_ = v1672
	var v1704 int64
	_ = v1704
	var v1706 int64
	_ = v1706
	var v1711 int32
	_ = v1711
	var v1724 int32
	_ = v1724
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1760 int32
	_ = v1760
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1774 int32
	_ = v1774
	v4 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(48)
	m.G0 = v31
	if l2 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v31 + int32(48)
	return v1774
L2:
	;
	v407 = v388 + int32(1)
	if base.Ui32(int32(512)) < base.Ui32(v407) {
		goto L78
	} else {
		goto L79
	}
L3:
	;
	if l0&int32(3) == int32(0) {
		v56 = l0
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L5
L5:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v93 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = int32(0)
	v378 = l0
	v388 = v89
	v389 = l0 + v89
	goto L2
L7:
	;
	v89 = v81 - l0
	goto L6
L8:
	;
	v60 = v56
	goto L17
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v40 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v89 = int32(0)
	goto L6
L11:
	;
	goto L12
L12:
	;
	v45 = l0
	goto L13
L13:
	;
	v49 = v45 + int32(1)
	if v49&int32(3) == int32(0) {
		v56 = v49
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v81 = v49
	goto L7
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v54 != 0 {
		v45 = v49
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v69 = int32(-2139062144)
	if (int32(16843008)-v66|v66)&v69 == v69 {
		v60 = v60 + int32(4)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v75 = v60
	goto L20
L19:
	;
	goto L18
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v79 != 0 {
		v75 = v75 + int32(1)
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v81 = v75
	goto L7
L22:
	;
	goto L21
L23:
	;
	if v188&int32(255) == int32(0) {
		v1774 = v4
		goto L1
	} else {
		goto L44
	}
L24:
	;
	v186 = l0
	v188 = v93
	v189 = l0
	v196 = l0 - l0
	goto L23
L25:
	;
	goto L26
L26:
	;
	if v93 != int32(60) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v101 = v93
	v102 = l0
	goto L31
L28:
	;
	goto L29
L29:
	;
	v146 = l0 + int32(1)
	v149 = v146
	goto L36
L30:
	;
	v186 = l0
	v188 = v142
	v189 = v143
	v196 = v143 - l0
	goto L23
L31:
	;
	if base.Ui32(int32(252)) < base.Ui32((v101-int32(46))&int32(255)) {
		v142 = v101
		v143 = v102
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v142 = int32(0)
	v143 = v139
	goto L30
L33:
	;
	if base.Ui32(int32(-11)) < base.Ui32(base.I32_extend8_s(v101)-int32(58)) {
		v142 = v101
		v143 = v102
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v139 = v102 + int32(1)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139))))
	if v140 != 0 {
		v101 = v140
		v102 = v139
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v175 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v184 = v149 + int32(1)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	v186 = v146
	v188 = v185
	v189 = v184
	v196 = v149 - v146
	goto L23
L38:
	;
	v1774 = v4
	goto L1
L39:
	;
	goto L40
L40:
	;
	if v175 != int32(62) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v149 = v149 + int32(1)
	goto L36
L42:
	;
	goto L43
L43:
	;
	goto L37
L44:
	;
	v219 = v31 + int32(44)
	v220 = int32(0)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	switch v227 - int32(43) {
	case 0:
		goto L49
	default:
		v237 = v189
		v238 = v227
		v240 = v220
		goto L47
	case 2:
		v231 = int32(1)
		goto L48
	}
L45:
	;
	if v372 == int32(0) {
		v1774 = v4
		goto L1
	} else {
		goto L77
	}
L46:
	;
	goto L45
L47:
	;
	if base.Ui32(int32(9)) < base.Ui32(base.I32_extend8_s(v238)-int32(48)) {
		v372 = v220
		goto L46
	} else {
		goto L51
	}
L48:
	;
	v233 = v189 + int32(1)
	if v233 == int32(0) {
		v372 = v220
		goto L46
	} else {
		goto L50
	}
L49:
	;
	v231 = int32(0)
	goto L48
L50:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233))))
	v237 = v233
	v238 = v236
	v240 = v231
	goto L47
L51:
	;
	v247 = v237
	v249 = v238
	v250 = int32(0)
	goto L52
L52:
	;
	v260 = base.I32_extend8_s(v249) + v250*int32(10) - int32(48)
	if int32(167) < v260 {
		v372 = v220
		goto L46
	} else {
		goto L54
	}
L53:
	;
	if v260 < int32(0) {
		v372 = v220
		goto L46
	} else {
		goto L56
	}
L54:
	;
	v264 = v247 + int32(1)
	v265 = int32(*(*int8)(unsafe.Add(mBase, uint32(v264))))
	if base.Ui32(v265-int32(48)) < base.Ui32(int32(10)) {
		v247 = v264
		v249 = v265
		v250 = v260
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v273 = v260 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v273
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264))))
	if v275 != int32(58) {
		v357 = v264
		v362 = v273
		goto L57
	} else {
		goto L58
	}
L57:
	;
	if v240 != 0 {
		goto L74
	} else {
		goto L75
	}
L58:
	;
	v279 = v247 + int32(2)
	if v279 == int32(0) {
		v372 = v220
		goto L46
	} else {
		goto L59
	}
L59:
	;
	v282 = int32(*(*int8)(unsafe.Add(mBase, uint32(v279))))
	if base.Ui32(int32(9)) < base.Ui32(v282-int32(48)) {
		v372 = v220
		goto L46
	} else {
		goto L60
	}
L60:
	;
	v288 = v279
	v290 = int32(0)
	v291 = v282
	goto L61
L61:
	;
	v301 = base.I32_extend8_s(v291) + v290*int32(10) - int32(48)
	if int32(59) < v301 {
		v372 = v220
		goto L46
	} else {
		goto L63
	}
L62:
	;
	if v301 < int32(0) {
		v372 = v220
		goto L46
	} else {
		goto L65
	}
L63:
	;
	v305 = v288 + int32(1)
	v306 = int32(*(*int8)(unsafe.Add(mBase, uint32(v305))))
	if base.Ui32(v306-int32(48)) < base.Ui32(int32(10)) {
		v288 = v305
		v290 = v301
		v291 = v306
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v315 = v301*int32(60) + v273
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v315
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305))))
	if v317 != int32(58) {
		v357 = v305
		v362 = v315
		goto L57
	} else {
		goto L66
	}
L66:
	;
	v321 = v288 + int32(2)
	if v321 == int32(0) {
		v372 = v220
		goto L46
	} else {
		goto L67
	}
L67:
	;
	v324 = int32(*(*int8)(unsafe.Add(mBase, uint32(v321))))
	if base.Ui32(int32(9)) < base.Ui32(v324-int32(48)) {
		v372 = v220
		goto L46
	} else {
		goto L68
	}
L68:
	;
	v330 = v321
	v332 = int32(0)
	v333 = v324
	goto L69
L69:
	;
	v343 = base.I32_extend8_s(v333) + v332*int32(10) - int32(48)
	if int32(60) < v343 {
		v372 = v220
		goto L46
	} else {
		goto L71
	}
L70:
	;
	if v343 < int32(0) {
		v372 = v220
		goto L46
	} else {
		goto L73
	}
L71:
	;
	v347 = v330 + int32(1)
	v348 = int32(*(*int8)(unsafe.Add(mBase, uint32(v347))))
	if base.Ui32(v348-int32(48)) < base.Ui32(int32(10)) {
		v330 = v347
		v332 = v343
		v333 = v348
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v355 = v343 + v315
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = v355
	v357 = v347
	v362 = v355
	goto L57
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v219))) = int32(0) - v362
	goto L76
L75:
	;
	goto L76
L76:
	;
	v372 = v357
	goto L46
L77:
	;
	v378 = v186
	v388 = v196
	v389 = v372
	goto L2
L78:
	;
	v1774 = v4
	goto L1
L79:
	;
	goto L80
L80:
	;
	v410 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v410)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v410
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v389))))
	if v414 != 0 {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v535
	v1756 = l1 + int32(22120)
	if v388 != 0 {
		goto L317
	} else {
		goto L318
	}
L82:
	;
	if v1711-v1038 < int32(401) {
		goto L81
	} else {
		goto L315
	}
L83:
	;
	v1704 = *(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1193])))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1194]))) = v1704
	v1706 = *(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1195])))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1196]))) = v1706
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	goto L81
L84:
	;
	v1042 = l1 + int32(16024)
	v1044 = l1 + int32(24)
	v1051 = v1037
	v1052 = v1038
	v1055 = v4
	v1061 = v776 + int32(400)
	v1074 = v1039
	goto L193
L85:
	;
	if v776 < int32(2147483248) {
		v1037 = v814
		v1038 = v797
		v1039 = v793
		goto L84
	} else {
		goto L192
	}
L86:
	;
	if v414 != int32(60) {
		goto L90
	} else {
		goto L91
	}
L87:
	;
	goto L88
L88:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(4294967296)
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v1012 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1197]))) = v1012
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1198]))) = uint16(v1012)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1196]))) = v1012
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1199]))) = uint8(v1012)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v407
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1194]))) = v1012 - v1011
	v1025 = l1 + int32(22120)
	if v388 != 0 {
		goto L189
	} else {
		goto L190
	}
L89:
	;
	if v530 == int32(0) {
		v1774 = v4
		goto L1
	} else {
		goto L107
	}
L90:
	;
	v419 = v414
	v420 = v389
	goto L93
L91:
	;
	goto L92
L92:
	;
	v463 = v389 + int32(1)
	v466 = v463
	goto L99
L93:
	;
	if base.Ui32(int32(252)) < base.Ui32((v419-int32(46))&int32(255)) {
		v460 = v420
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v505 = v460
	v513 = v389
	v530 = v460 - v389
	goto L89
L95:
	;
	goto L94
L96:
	;
	if base.Ui32(int32(-11)) < base.Ui32(base.I32_extend8_s(v419)-int32(58)) {
		v460 = v420
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v457 = v420 + int32(1)
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457))))
	if v458 != 0 {
		v419 = v458
		v420 = v457
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v460 = v457
	goto L95
L99:
	;
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v466))))
	if v492 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v505 = v466 + int32(1)
	v513 = v463
	v530 = v466 - v463
	goto L89
L101:
	;
	v1774 = v4
	goto L1
L102:
	;
	goto L103
L103:
	;
	if v492 != int32(62) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v466 = v466 + int32(1)
	goto L99
L105:
	;
	goto L106
L106:
	;
	goto L100
L107:
	;
	v535 = v530 + v388 + int32(2)
	if base.Ui32(int32(512)) < base.Ui32(v535) {
		v1774 = v4
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	switch v538 - int32(44) {
	case 0, 15:
		goto L110
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L111
	default:
		goto L112
	}
L109:
	;
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v705))))
	if v707 != 0 {
		goto L151
	} else {
		goto L152
	}
L110:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v701 - int32(3600)
	v705 = v505
	goto L109
L111:
	;
	v544 = v31 + int32(40)
	v545 = int32(0)
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	switch v552 - int32(43) {
	case 0:
		goto L118
	default:
		v562 = v505
		v563 = v552
		v565 = v545
		goto L116
	case 2:
		v556 = int32(1)
		goto L117
	}
L112:
	;
	if v538 == int32(0) {
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	if v697 != 0 {
		v705 = v697
		goto L109
	} else {
		goto L146
	}
L115:
	;
	goto L114
L116:
	;
	if base.Ui32(int32(9)) < base.Ui32(base.I32_extend8_s(v563)-int32(48)) {
		v697 = v545
		goto L115
	} else {
		goto L120
	}
L117:
	;
	v558 = v505 + int32(1)
	if v558 == int32(0) {
		v697 = v545
		goto L115
	} else {
		goto L119
	}
L118:
	;
	v556 = int32(0)
	goto L117
L119:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	v562 = v558
	v563 = v561
	v565 = v556
	goto L116
L120:
	;
	v572 = v562
	v574 = v563
	v575 = int32(0)
	goto L121
L121:
	;
	v585 = base.I32_extend8_s(v574) + v575*int32(10) - int32(48)
	if int32(167) < v585 {
		v697 = v545
		goto L115
	} else {
		goto L123
	}
L122:
	;
	if v585 < int32(0) {
		v697 = v545
		goto L115
	} else {
		goto L125
	}
L123:
	;
	v589 = v572 + int32(1)
	v590 = int32(*(*int8)(unsafe.Add(mBase, uint32(v589))))
	if base.Ui32(v590-int32(48)) < base.Ui32(int32(10)) {
		v572 = v589
		v574 = v590
		v575 = v585
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v598 = v585 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = v598
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	if v600 != int32(58) {
		v682 = v589
		v687 = v598
		goto L126
	} else {
		goto L127
	}
L126:
	;
	if v565 != 0 {
		goto L143
	} else {
		goto L144
	}
L127:
	;
	v604 = v572 + int32(2)
	if v604 == int32(0) {
		v697 = v545
		goto L115
	} else {
		goto L128
	}
L128:
	;
	v607 = int32(*(*int8)(unsafe.Add(mBase, uint32(v604))))
	if base.Ui32(int32(9)) < base.Ui32(v607-int32(48)) {
		v697 = v545
		goto L115
	} else {
		goto L129
	}
L129:
	;
	v613 = v604
	v615 = int32(0)
	v616 = v607
	goto L130
L130:
	;
	v626 = base.I32_extend8_s(v616) + v615*int32(10) - int32(48)
	if int32(59) < v626 {
		v697 = v545
		goto L115
	} else {
		goto L132
	}
L131:
	;
	if v626 < int32(0) {
		v697 = v545
		goto L115
	} else {
		goto L134
	}
L132:
	;
	v630 = v613 + int32(1)
	v631 = int32(*(*int8)(unsafe.Add(mBase, uint32(v630))))
	if base.Ui32(v631-int32(48)) < base.Ui32(int32(10)) {
		v613 = v630
		v615 = v626
		v616 = v631
		goto L130
	} else {
		goto L133
	}
L133:
	;
	goto L131
L134:
	;
	v640 = v626*int32(60) + v598
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = v640
	v642 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630))))
	if v642 != int32(58) {
		v682 = v630
		v687 = v640
		goto L126
	} else {
		goto L135
	}
L135:
	;
	v646 = v613 + int32(2)
	if v646 == int32(0) {
		v697 = v545
		goto L115
	} else {
		goto L136
	}
L136:
	;
	v649 = int32(*(*int8)(unsafe.Add(mBase, uint32(v646))))
	if base.Ui32(int32(9)) < base.Ui32(v649-int32(48)) {
		v697 = v545
		goto L115
	} else {
		goto L137
	}
L137:
	;
	v655 = v646
	v657 = int32(0)
	v658 = v649
	goto L138
L138:
	;
	v668 = base.I32_extend8_s(v658) + v657*int32(10) - int32(48)
	if int32(60) < v668 {
		v697 = v545
		goto L115
	} else {
		goto L140
	}
L139:
	;
	if v668 < int32(0) {
		v697 = v545
		goto L115
	} else {
		goto L142
	}
L140:
	;
	v672 = v655 + int32(1)
	v673 = int32(*(*int8)(unsafe.Add(mBase, uint32(v672))))
	if base.Ui32(v673-int32(48)) < base.Ui32(int32(10)) {
		v655 = v672
		v657 = v668
		v658 = v673
		goto L138
	} else {
		goto L141
	}
L141:
	;
	goto L139
L142:
	;
	v680 = v668 + v640
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = v680
	v682 = v672
	v687 = v680
	goto L126
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v544))) = int32(0) - v687
	goto L145
L144:
	;
	goto L145
L145:
	;
	v697 = v682
	goto L115
L146:
	;
	v1774 = v4
	goto L1
L147:
	;
	v979 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1198]))) = uint16(v979)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1196]))) = v979
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1199]))) = uint8(v979)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1194]))) = v979 - v963
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1200]))) = uint16(v979)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1195]))) = v407
	v998 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1201]))) = uint8(v998)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1197]))) = v979
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1193]))) = v979 - v988
	goto L81
L148:
	;
	v838 = v712
	goto L174
L149:
	;
	v721 = F_getrule(m, v708+int32(1), v31+int32(20))
	mBase = m.M
	if v721 == int32(0) {
		v1774 = v4
		goto L1
	} else {
		goto L156
	}
L150:
	;
	if v709 != 0 {
		v1774 = v4
		goto L1
	} else {
		goto L154
	}
L151:
	;
	v708 = v705
	goto L153
L152:
	;
	v708 = int32(542276)
	goto L153
L153:
	;
	v709 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v708))))
	switch v709 - int32(44) {
	case 0, 15:
		goto L149
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		v1774 = v4
		goto L1
	default:
		goto L150
	}
L154:
	;
	v712 = int32(0)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v712 < v713 {
		goto L148
	} else {
		goto L155
	}
L155:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v963 = v716
	goto L147
L156:
	;
	v724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v721))))
	if v724 != int32(44) {
		v1774 = v4
		goto L1
	} else {
		goto L157
	}
L157:
	;
	v729 = F_getrule(m, v721+int32(1), v31)
	mBase = m.M
	if v729 == int32(0) {
		v1774 = v4
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v729))))
	if v732 != 0 {
		v1774 = v4
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(2)
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v736 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1198]))) = uint16(v736)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1196]))) = v736
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1199]))) = uint8(v736)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1194]))) = v736 - v735
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1197]))) = v736
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1200]))) = uint16(v736)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1195]))) = v407
	v757 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1201]))) = uint8(v757)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1193]))) = v736 - v745
	v776 = int32(1970)
	v793 = int64(0)
	goto L160
L160:
	;
	v797 = v776 - int32(1)
	if v797&int32(3) != 0 {
		v807 = int32(0)
		goto L163
	} else {
		goto L164
	}
L161:
	;
	v1037 = int32(0)
	v1038 = int32(1770)
	v1039 = v826
	goto L84
L162:
	;
	v826 = v793 + v825
	if base.Ui32(int32(1771)) < base.Ui32(v776) {
		v776 = v797
		v793 = v826
		goto L160
	} else {
		goto L171
	}
L163:
	;
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v807<<(uint(int32(2))%32))+uint32(_consts[1202])))
	v814 = v812 * int32(-86400)
	if v814 < int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	v802 = base.I32_rem_u_s(v797, int32(100))
	if v802 != 0 {
		v807 = int32(1)
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v804 = base.I32_rem_u_s(v797, int32(400))
	v807 = base.B2i32(v804 == int32(0))
	goto L163
L166:
	;
	v818 = base.I64_extend_i32_s(v814)
	if int64(-9223372036854775807-1)-v818 <= v793 {
		v825 = v818
		goto L162
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	v821 = base.I64_extend_i32_u(v814)
	if v821^int64(9223372036854775807) < v793 {
		goto L85
	} else {
		goto L170
	}
L169:
	;
	goto L85
L170:
	;
	v825 = v821
	goto L162
L171:
	;
	goto L161
L172:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v891 = int32(0)
	v897 = v879
	goto L178
L173:
	;
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v867)))
	v879 = int32(0) - v876
	goto L172
L174:
	;
	v864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838+(l1+int32(16024))))))
	v867 = l1 + int32(18024) + v864<<(uint(int32(4))%32)
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867)+4)))
	if v868 == int32(0) {
		goto L173
	} else {
		goto L176
	}
L175:
	;
	v879 = int32(0)
	goto L172
L176:
	;
	v872 = v838 + int32(1)
	if v872 != v713 {
		v838 = v872
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v916 = v891 + (l1 + int32(16024))
	v917 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v916))))
	v920 = l1 + int32(18024) + v917<<(uint(int32(4))%32)
	v921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v916))) = uint8(v921)
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+13)))
	if v923 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	v963 = v886
	goto L147
L180:
	;
	v948 = v891 + int32(1)
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v948 < v949 {
		v891 = v948
		v897 = v946
		goto L178
	} else {
		goto L187
	}
L181:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v920)))
	v946 = int32(0) - v943
	goto L180
L182:
	;
	v928 = l1 + int32(24) + v891<<(uint(int32(3))%32)
	v929 = *(*int64)(unsafe.Add(mBase, uint32(v928)))
	*(*int64)(unsafe.Add(mBase, uint32(v928))) = v929 + base.I64_extend_i32_s(v886-v897)
	v934 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+4)))
	if v934&int32(1) == int32(0) {
		goto L181
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	if v921&int32(1) != 0 {
		v946 = v897
		goto L180
	} else {
		goto L186
	}
L185:
	;
	v946 = v897
	goto L180
L186:
	;
	goto L181
L187:
	;
	goto L179
L188:
	;
	v1029 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1027+v388))) = uint8(v1029)
	v1774 = int32(1)
	goto L1
L189:
	;
	v1026 = F__emscripten_memcpy_bulkmem(m, v1025, v378, v388)
	mBase = m.M
	v1027 = v1026
	goto L191
L190:
	;
	v1027 = v1025
	goto L191
L191:
	;
	goto L188
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	goto L83
L193:
	;
	v1077 = v31 + int32(20)
	v1078 = int32(0)
	if v1052&int32(3) != 0 {
		v1096 = v1078
		goto L196
	} else {
		goto L197
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1647
	if v1647 != 0 {
		v1711 = v1672
		goto L82
	} else {
		goto L314
	}
L195:
	;
	v1322 = int32(0)
	if v1052&int32(3) != 0 {
		v1340 = v1322
		goto L236
	} else {
		goto L237
	}
L196:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1077)))
	switch v1097 {
	case 0:
		goto L202
	case 1:
		goto L201
	case 2:
		goto L200
	default:
		v1311 = v1078
		goto L199
	}
L197:
	;
	v1091 = base.I32_rem_s(v1052, int32(100))
	if v1091 != 0 {
		v1096 = int32(1)
		goto L196
	} else {
		goto L198
	}
L198:
	;
	v1093 = base.I32_rem_s(v1052, int32(400))
	v1096 = base.B2i32(v1093 == int32(0))
	goto L196
L199:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+16))
	v1321 = v1319 + (v735 + v1311)
	goto L195
L200:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+4))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+12))
	v1114 = v1052 - base.B2i32(v1111 < int32(3))
	v1116 = base.I32_div_s(v1114, int32(400))
	v1118 = base.I32_rem_s(v1114, int32(100))
	v1121 = base.I32_div_s(v1114, int32(-100))
	v1122 = int32(1)
	v1127 = base.I32_div_s(base.I32_extend8_s(v1118), int32(4))
	v1133 = base.I32_rem_s(v1111+int32(9), int32(12))
	v1140 = base.I32_div_s(base.I32_extend16_s(v1133*int32(26)+int32(24)), int32(10))
	v1145 = int32(7)
	v1146 = base.I32_rem_s(v1116+v1118+v1121<<(uint(v1122)%32)+base.I32_extend8_s(v1127)+base.I32_extend16_s(v1140+v1122), v1145)
	if v1146 < int32(0) {
		goto L209
	} else {
		goto L210
	}
L201:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+4))
	v1311 = v1107 * int32(86400)
	goto L199
L202:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+4))
	v1099 = int32(86400)
	v1100 = v1098 * v1099
	v1102 = v1100 - v1099
	if int32(59) < v1098 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v1105 = v1100
	goto L205
L204:
	;
	v1105 = v1102
	goto L205
L205:
	;
	if v1096 != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v1106 = v1105
	goto L208
L207:
	;
	v1106 = v1102
	goto L208
L208:
	;
	v1311 = v1106
	goto L199
L209:
	;
	v1151 = v1146 + v1145
	goto L211
L210:
	;
	v1151 = v1146
	goto L211
L211:
	;
	v1152 = v1110 - v1151
	if v1152 < int32(0) {
		goto L212
	} else {
		goto L213
	}
L212:
	;
	v1157 = v1152 + int32(7)
	goto L214
L213:
	;
	v1157 = v1152
	goto L214
L214:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1077)+8))
	if v1158 <= int32(1) {
		goto L216
	} else {
		goto L217
	}
L215:
	;
	v1208 = v1196 * int32(86400)
	if v1202 <= int32(0) {
		v1311 = v1208
		goto L199
	} else {
		goto L223
	}
L216:
	;
	v1196 = v1157
	v1202 = v1111 - int32(1)
	goto L215
L217:
	;
	goto L218
L218:
	;
	v1165 = int32(1)
	v1166 = v1111 - v1165
	v1172 = *(*int32)(unsafe.Add(mBase, uint32(v1096*int32(48)+v1166<<(uint(int32(2))%32))+uint32(_consts[1203])))
	v1173 = int32(7)
	v1179 = v1157
	v1182 = v1165
	goto L219
L219:
	;
	v1191 = v1179 + int32(7)
	if v1172 <= v1191 {
		v1196 = v1179
		v1202 = v1166
		goto L215
	} else {
		goto L221
	}
L220:
	;
	v1196 = v1157 + v1158*v1173 - v1173
	v1202 = v1166
	goto L215
L221:
	;
	v1194 = v1182 + int32(1)
	if v1194 != v1158 {
		v1179 = v1191
		v1182 = v1194
		goto L219
	} else {
		goto L222
	}
L222:
	;
	goto L220
L223:
	;
	v1211 = int32(3)
	v1212 = v1202 & v1211
	if base.Ui32(v1111-int32(2)) < base.Ui32(v1211) {
		goto L225
	} else {
		goto L226
	}
L224:
	;
	if v1212 == int32(0) {
		v1311 = v1271
		goto L199
	} else {
		goto L231
	}
L225:
	;
	v1268 = int32(0)
	v1271 = v1208
	goto L224
L226:
	;
	goto L227
L227:
	;
	v1221 = int32(0)
	v1225 = v1221
	v1228 = v1208
	v1234 = v1221
	goto L228
L228:
	;
	v1238 = v1096*int32(48) + v1225<<(uint(int32(2))%32)
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+uint32(_consts[1204])))
	v1242 = int32(86400)
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+uint32(_consts[1203])))
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+uint32(_consts[1205])))
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(v1238)+uint32(_consts[1206])))
	v1262 = v1241*v1242 + (v1246*v1242 + v1228 + v1252*v1242 + v1258*v1242)
	v1263 = int32(4)
	v1264 = v1225 + v1263
	v1266 = v1234 + v1263
	if v1266 != v1202&int32(2147483644) {
		v1225 = v1264
		v1228 = v1262
		v1234 = v1266
		goto L228
	} else {
		goto L230
	}
L229:
	;
	v1268 = v1264
	v1271 = v1262
	goto L224
L230:
	;
	goto L229
L231:
	;
	v1283 = v1268
	v1286 = v1271
	v1290 = int32(0)
	goto L232
L232:
	;
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1096*int32(48)+v1283<<(uint(int32(2))%32))+uint32(_consts[1203])))
	v1302 = v1299*int32(86400) + v1286
	v1303 = int32(1)
	v1306 = v1290 + v1303
	if v1306 != v1212 {
		v1283 = v1283 + v1303
		v1286 = v1302
		v1290 = v1306
		goto L232
	} else {
		goto L234
	}
L233:
	;
	v1311 = v1302
	goto L199
L234:
	;
	goto L233
L235:
	;
	if v1052&int32(3) != 0 {
		v1576 = int32(0)
		goto L275
	} else {
		goto L276
	}
L236:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	switch v1341 {
	case 0:
		goto L242
	case 1:
		goto L241
	case 2:
		goto L240
	default:
		v1555 = v1322
		goto L239
	}
L237:
	;
	v1335 = base.I32_rem_s(v1052, int32(100))
	if v1335 != 0 {
		v1340 = int32(1)
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v1337 = base.I32_rem_s(v1052, int32(400))
	v1340 = base.B2i32(v1337 == int32(0))
	goto L236
L239:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v1565 = v1563 + (v745 + v1555)
	goto L235
L240:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1355 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v1358 = v1052 - base.B2i32(v1355 < int32(3))
	v1360 = base.I32_div_s(v1358, int32(400))
	v1362 = base.I32_rem_s(v1358, int32(100))
	v1365 = base.I32_div_s(v1358, int32(-100))
	v1366 = int32(1)
	v1371 = base.I32_div_s(base.I32_extend8_s(v1362), int32(4))
	v1377 = base.I32_rem_s(v1355+int32(9), int32(12))
	v1384 = base.I32_div_s(base.I32_extend16_s(v1377*int32(26)+int32(24)), int32(10))
	v1389 = int32(7)
	v1390 = base.I32_rem_s(v1360+v1362+v1365<<(uint(v1366)%32)+base.I32_extend8_s(v1371)+base.I32_extend16_s(v1384+v1366), v1389)
	if v1390 < int32(0) {
		goto L249
	} else {
		goto L250
	}
L241:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1555 = v1351 * int32(86400)
	goto L239
L242:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1343 = int32(86400)
	v1344 = v1342 * v1343
	v1346 = v1344 - v1343
	if int32(59) < v1342 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1349 = v1344
	goto L245
L244:
	;
	v1349 = v1346
	goto L245
L245:
	;
	if v1340 != 0 {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	v1350 = v1349
	goto L248
L247:
	;
	v1350 = v1346
	goto L248
L248:
	;
	v1555 = v1350
	goto L239
L249:
	;
	v1395 = v1390 + v1389
	goto L251
L250:
	;
	v1395 = v1390
	goto L251
L251:
	;
	v1396 = v1354 - v1395
	if v1396 < int32(0) {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v1401 = v1396 + int32(7)
	goto L254
L253:
	;
	v1401 = v1396
	goto L254
L254:
	;
	v1402 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v1402 <= int32(1) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	v1452 = v1440 * int32(86400)
	if v1446 <= int32(0) {
		v1555 = v1452
		goto L239
	} else {
		goto L263
	}
L256:
	;
	v1440 = v1401
	v1446 = v1355 - int32(1)
	goto L255
L257:
	;
	goto L258
L258:
	;
	v1409 = int32(1)
	v1410 = v1355 - v1409
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1340*int32(48)+v1410<<(uint(int32(2))%32))+uint32(_consts[1203])))
	v1417 = int32(7)
	v1423 = v1401
	v1426 = v1409
	goto L259
L259:
	;
	v1435 = v1423 + int32(7)
	if v1416 <= v1435 {
		v1440 = v1423
		v1446 = v1410
		goto L255
	} else {
		goto L261
	}
L260:
	;
	v1440 = v1401 + v1402*v1417 - v1417
	v1446 = v1410
	goto L255
L261:
	;
	v1438 = v1426 + int32(1)
	if v1438 != v1402 {
		v1423 = v1435
		v1426 = v1438
		goto L259
	} else {
		goto L262
	}
L262:
	;
	goto L260
L263:
	;
	v1455 = int32(3)
	v1456 = v1446 & v1455
	if base.Ui32(v1355-int32(2)) < base.Ui32(v1455) {
		goto L265
	} else {
		goto L266
	}
L264:
	;
	if v1456 == int32(0) {
		v1555 = v1515
		goto L239
	} else {
		goto L271
	}
L265:
	;
	v1512 = int32(0)
	v1515 = v1452
	goto L264
L266:
	;
	goto L267
L267:
	;
	v1465 = int32(0)
	v1469 = v1465
	v1472 = v1452
	v1478 = v1465
	goto L268
L268:
	;
	v1482 = v1340*int32(48) + v1469<<(uint(int32(2))%32)
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+uint32(_consts[1204])))
	v1486 = int32(86400)
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+uint32(_consts[1203])))
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+uint32(_consts[1205])))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1482)+uint32(_consts[1206])))
	v1506 = v1485*v1486 + (v1490*v1486 + v1472 + v1496*v1486 + v1502*v1486)
	v1507 = int32(4)
	v1508 = v1469 + v1507
	v1510 = v1478 + v1507
	if v1510 != v1446&int32(2147483644) {
		v1469 = v1508
		v1472 = v1506
		v1478 = v1510
		goto L268
	} else {
		goto L270
	}
L269:
	;
	v1512 = v1508
	v1515 = v1506
	goto L264
L270:
	;
	goto L269
L271:
	;
	v1527 = v1512
	v1530 = v1515
	v1534 = int32(0)
	goto L272
L272:
	;
	v1543 = *(*int32)(unsafe.Add(mBase, uint32(v1340*int32(48)+v1527<<(uint(int32(2))%32))+uint32(_consts[1203])))
	v1546 = v1543*int32(86400) + v1530
	v1547 = int32(1)
	v1550 = v1534 + v1547
	if v1550 != v1456 {
		v1527 = v1527 + v1547
		v1530 = v1546
		v1534 = v1550
		goto L272
	} else {
		goto L274
	}
L273:
	;
	v1555 = v1546
	goto L239
L274:
	;
	goto L273
L275:
	;
	v1581 = *(*int32)(unsafe.Add(mBase, uint32(v1576<<(uint(int32(2))%32))+uint32(_consts[1202])))
	v1583 = v1581 * int32(86400)
	v1584 = base.B2i32(v1565 < v1321)
	if v1584 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v1571 = base.I32_rem_u_s(v1052, int32(100))
	if v1571 != 0 {
		v1576 = int32(1)
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1573 = base.I32_rem_u_s(v1052, int32(400))
	v1576 = base.B2i32(v1573 == int32(0))
	goto L275
L278:
	;
	v1654 = v1051 + v1583
	if v1654 < int32(0) {
		goto L308
	} else {
		goto L309
	}
L279:
	;
	if v1565 <= v1321 {
		v1647 = v1055
		v1649 = v1061
		goto L278
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	if int32(1999) <= v1055 {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	if v1583+(v735-v745) <= v1565-v1321 {
		v1647 = v1055
		v1649 = v1061
		goto L278
	} else {
		goto L283
	}
L283:
	;
	goto L281
L284:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1055
	v1711 = v1052
	goto L82
L285:
	;
	goto L286
L286:
	;
	if v1321 < v1565 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1595 = v1565
	goto L289
L288:
	;
	v1595 = v1321
	goto L289
L289:
	;
	v1598 = v1044 + v1055<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1598))) = v1074
	if v1565 < v1321 {
		goto L292
	} else {
		goto L293
	}
L290:
	;
	v1624 = v1044 + v1620<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1624))) = v1074
	v1626 = v1051 + v1595
	if v1626 < int32(0) {
		goto L301
	} else {
		goto L302
	}
L291:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1598))) = v1074 + v1612
	*(*uint8)(unsafe.Add(mBase, uint32(v1055+v1042))) = uint8(base.B2i32(v1321 <= v1565))
	v1620 = v1055 + int32(1)
	goto L290
L292:
	;
	v1600 = v1565
	goto L294
L293:
	;
	v1600 = v1321
	goto L294
L294:
	;
	v1601 = v1600 + v1051
	if v1601 < int32(0) {
		goto L295
	} else {
		goto L296
	}
L295:
	;
	v1605 = base.I64_extend_i32_s(v1601)
	if int64(-9223372036854775807-1)-v1605 <= v1074 {
		v1612 = v1605
		goto L291
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	v1608 = base.I64_extend_i32_u(v1601)
	if v1608^int64(9223372036854775807) < v1074 {
		v1620 = v1055
		goto L290
	} else {
		goto L299
	}
L298:
	;
	v1620 = v1055
	goto L290
L299:
	;
	v1612 = v1608
	goto L291
L300:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1624))) = v1074 + v1637
	*(*uint8)(unsafe.Add(mBase, uint32(v1620+v1042))) = uint8(v1584)
	v1647 = v1620 + int32(1)
	v1649 = v1052 + int32(401)
	goto L278
L301:
	;
	v1630 = base.I64_extend_i32_s(v1626)
	if int64(-9223372036854775807-1)-v1630 <= v1074 {
		v1637 = v1630
		goto L300
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	v1633 = base.I64_extend_i32_u(v1626)
	if v1633^int64(9223372036854775807) < v1074 {
		v1647 = v1620
		v1649 = v1061
		goto L278
	} else {
		goto L305
	}
L304:
	;
	v1647 = v1620
	v1649 = v1061
	goto L278
L305:
	;
	v1637 = v1633
	goto L300
L306:
	;
	goto L194
L307:
	;
	v1669 = v1052 + int32(1)
	if v1669 < v1649 {
		v1051 = int32(0)
		v1052 = v1669
		v1055 = v1647
		v1061 = v1649
		v1074 = v1074 + v1665
		goto L193
	} else {
		goto L313
	}
L308:
	;
	v1658 = base.I64_extend_i32_s(v1654)
	if int64(-9223372036854775807-1)-v1658 <= v1074 {
		v1665 = v1658
		goto L307
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	v1661 = base.I64_extend_i32_u(v1654)
	if v1661^int64(9223372036854775807) < v1074 {
		v1672 = v1052
		goto L306
	} else {
		goto L312
	}
L311:
	;
	v1672 = v1052
	goto L306
L312:
	;
	v1665 = v1661
	goto L307
L313:
	;
	v1672 = v1669
	goto L306
L314:
	;
	goto L83
L315:
	;
	v1724 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v1724)
	goto L81
L316:
	;
	v1759 = v1758 + v388
	v1760 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1759))) = uint8(v1760)
	v1762 = int32(1)
	v1764 = v1759 + v1762
	if v530 != 0 {
		goto L321
	} else {
		goto L322
	}
L317:
	;
	v1757 = F__emscripten_memcpy_bulkmem(m, v1756, v378, v388)
	mBase = m.M
	v1758 = v1757
	goto L319
L318:
	;
	v1758 = v1756
	goto L319
L319:
	;
	goto L316
L320:
	;
	v1768 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1766+v530))) = uint8(v1768)
	v1774 = v1762
	goto L1
L321:
	;
	v1765 = F__emscripten_memcpy_bulkmem(m, v1764, v513, v530)
	mBase = m.M
	v1766 = v1765
	goto L323
L322:
	;
	v1766 = v1764
	goto L323
L323:
	;
	goto L320
}
