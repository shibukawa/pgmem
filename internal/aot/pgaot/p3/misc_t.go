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
	*(*int32)(unsafe.Add(mBase, _consts[1193])) = v2
	*(*int32)(unsafe.Add(mBase, _consts[1187])) = v2
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(228116)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(492025)
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
	v65 = int32(4102492)
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
	F_ResourceOwnerForget(m, v97, v98+int32(1), int32(1609628))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = int32(228116)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(492025)
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
	v161 = int32(4102492)
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
		v8 = int32(523443)
	} else {
		v8 = int32(0)
	}
	if v5 != 0 {
		v10 = v8
	} else {
		v10 = int32(535234)
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
			F_errmsg_internal(m, int32(52592), v6)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				F_errfinish(m, int32(496378), int32(1568), int32(241442))
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
													F_errmsg(m, int32(437791), v6+int32(16))
													mBase = m.M
													v50 = m.ExcPending
													if v50 != 0 {
														return
													} else {
														F_errhint(m, int32(555839), int32(0))
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return
														} else {
															F_errfinish(m, int32(496378), int32(1564), int32(241442))
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
			v14 = F_DirectFunctionCall1Coll(m, int32(5636), int32(0), v12)
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
	v6 = F_DirectFunctionCall2Coll(m, int32(1559), v3, v4, v5)
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
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
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
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
	v45 = F_strlen(m, v13)
	mBase = m.M
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v46 != int32(950) {
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
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v9 {
		goto L51
	} else {
		goto L52
	}
L15:
	;
	v141 = int32(1)
	if v14&v141 != 0 {
		goto L47
	} else {
		goto L48
	}
L16:
	;
	if v46 != 0 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if v44 != v45 {
		v151 = int32(1)
		goto L14
	} else {
		goto L25
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errmsg(m, int32(244111), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errhint(m, int32(556503), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(496954), int32(1648), int32(105597))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
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
	v71 = int32(1)
	if v14&v71 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v75 = v71
	goto L28
L27:
	;
	v75 = int32(4)
	goto L28
L28:
	;
	v76 = v9 + v75
	if base.Ui32(int32(4)) <= base.Ui32(v44) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v151 = base.B2i32(v138 != int32(0))
	goto L14
L30:
	;
	v138 = int32(0)
	goto L29
L31:
	;
	v112 = v107
	v113 = v108
	v114 = v109
	goto L41
L32:
	;
	if (v76|v13)&int32(3) != 0 {
		v107 = v76
		v108 = v13
		v109 = v44
		goto L31
	} else {
		goto L35
	}
L33:
	;
	v100 = v76
	v101 = v13
	v102 = v44
	goto L34
L34:
	;
	if v102 == int32(0) {
		goto L30
	} else {
		goto L40
	}
L35:
	;
	v84 = v76
	v85 = v13
	v86 = v44
	goto L36
L36:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v89 != v90 {
		v107 = v84
		v108 = v85
		v109 = v86
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v100 = v95
	v101 = v93
	v102 = v97
	goto L34
L38:
	;
	v92 = int32(4)
	v93 = v85 + v92
	v95 = v84 + v92
	v97 = v86 - v92
	if base.Ui32(int32(3)) < base.Ui32(v97) {
		v84 = v95
		v85 = v93
		v86 = v97
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v107 = v100
	v108 = v101
	v109 = v102
	goto L31
L41:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v117 == v118 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v138 = v117 - v118
	goto L29
L43:
	;
	v120 = int32(1)
	v125 = v114 - v120
	if v125 != 0 {
		v112 = v112 + v120
		v113 = v113 + v120
		v114 = v125
		goto L41
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	goto L30
L47:
	;
	v145 = v141
	goto L49
L48:
	;
	v145 = int32(4)
	goto L49
L49:
	;
	v147 = F_varstr_cmp(m, v9+v145, v44, v13, v45, v46)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v151 = base.B2i32(v147 != int32(0))
	goto L14
L51:
	;
	F_pfree(m, v9)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return v151
L54:
	;
	goto L53
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
											v103 = F_DirectFunctionCall3Coll(m, int32(1493), int32(0), v13, v84+int32(1), v92-v84)
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
	v99 = int32(4487040)
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
	F_errmsg(m, int32(301315), int32(0))
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
	F_errfinish(m, int32(493834), int32(389), int32(339645))
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
	F_errmsg(m, int32(301556), int32(0))
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
	F_errfinish(m, int32(493834), int32(370), int32(339645))
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
	v749 = int32(4487040)
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
	F_errmsg(m, int32(301354), int32(0))
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
	F_errdetail(m, int32(603596), v39+int32(16))
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
	F_errfinish(m, int32(493834), int32(418), int32(339645))
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
	F_errmsg(m, int32(692341), v39)
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
	F_errfinish(m, int32(493834), int32(508), int32(113477))
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
					F_errmsg(m, int32(400033), int32(0))
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(492756), int32(6509), int32(31034))
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
							F_errmsg(m, int32(400033), int32(0))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(492756), int32(6509), int32(31034))
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
						F_errmsg(m, int32(221384), int32(0))
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495500), int32(65), int32(275850))
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
	F_errmsg_internal(m, int32(114045), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(495701), int32(561), int32(13121))
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
	v25 = *(*int32)(unsafe.Add(mBase, uint32((l0-v15*v14)&v12<<(uint(v21)%32))+uint32(_consts[1355])))
	v26 = int32(8)
	v27 = int32(base.Ui32(l0) >> (uint(v26) % 32))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1356]))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30*int32(86)+v15)+uint32(_consts[1356]))))
	v41 = base.I32_rem_u_s(int32(base.Ui32(v25*v36)>>(uint(int32(11))%32)), int32(6))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_consts[1357]))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32((v41+v44)<<(uint(v21)%32))+uint32(_consts[1358])))
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
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1359]))))
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
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1360]))))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86<<(uint(int32(2))%32))+uint32(_consts[1358])))
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
	v50 = int32(73740)
	goto L21
L20:
	;
	v50 = int32(74014)
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
	F_errfinish(m, int32(495519), int32(3019), int32(357542))
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
	v122 = int32(93744)
	goto L42
L41:
	;
	v122 = int32(272476)
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
	F_errfinish(m, int32(495519), int32(3050), int32(357542))
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
	v79 = F_refnameNamespaceItem(m, l0, v75, int32(427322), int32(-1), v75)
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
	v81 = int32(427322)
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
	v92 = F_refnameNamespaceItem(m, l0, v88, int32(32047), int32(-1), v88)
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
	v94 = int32(32047)
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(539615)
	F_errmsg(m, int32(162287), v11+int32(-32))
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
	F_errfinish(m, int32(495217), int32(2669), int32(357634))
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
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(512992)
	F_errmsg(m, int32(162287), v11+int32(-16))
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
	F_errfinish(m, int32(495217), int32(2679), int32(357634))
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
	F_errmsg_internal(m, int32(479850), v13)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(495217), int32(2684), int32(357634))
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
	F_errmsg(m, int32(413080), v11+int32(-48))
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
	F_errfinish(m, int32(495217), int32(2692), int32(357634))
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
	F_errmsg(m, int32(272522), int32(0))
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
	F_errfinish(m, int32(495217), int32(2740), int32(357634))
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
			F_errcontext_msg(m, int32(709497), v8+int32(16))
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
			F_errcontext_msg(m, int32(697598), v8)
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
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v194 int32
	_ = v194
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v358 int32
	_ = v358
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
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
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
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
	v465 = m.ExcPending
	if v465 != 0 {
		goto L8
	} else {
		goto L89
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L8
	} else {
		goto L86
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L8
	} else {
		goto L83
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L8
	} else {
		goto L80
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L8
	} else {
		goto L77
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L8
	} else {
		goto L74
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L8
	} else {
		goto L71
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
	v369 = m.ExcPending
	if v369 != 0 {
		goto L8
	} else {
		goto L68
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
	v264 = v2
	goto L17
L17:
	;
	v275 = v264 + v38
	v276 = F_repalloc(m, v39, v275)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L51
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
	v264 = v241
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
	v255 = v55 + int32(1)
	if v255 != v25 {
		v48 = v48 + int32(12)
		v53 = v241
		v55 = v255
		goto L18
	} else {
		goto L50
	}
L22:
	;
	v214 = F_pq_getmsgint(m, v23, int32(1))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L42
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
	v80 = F_strlen(m, v78)
	mBase = m.M
	if v73&int32(240) != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(int32(2048)) <= base.Ui32(v80) {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	if int32(1048575) < v53 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	if v80 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v87 = int32(-1)
	if v80 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v194 = int32(0)
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v194
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+8)) = v80 | v53<<(uint(int32(12))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+2)) = uint8(base.B2i32(v76&int32(255) != int32(0)))
	*(*int32)(unsafe.Add(mBase, uint32(v33+v55<<(uint(int32(2))%32)))) = v78
	v241 = v80 + v53 + int32(1)
	goto L21
L33:
	;
	v93 = v78
	v94 = v87
	v101 = int32(0)
	goto L36
L34:
	;
	v140 = v78
	v141 = v87
	goto L35
L35:
	;
	if v80&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	v113 = int32(24)
	v116 = int32(2)
	v120 = *(*int32)(unsafe.Add(mBase, uint32((v112^int32(base.Ui32(v94)>>(uint(v113)%32)))<<(uint(v116)%32))+uint32(_consts[1047])))
	v121 = int32(8)
	v123 = v120 ^ v94<<(uint(v121)%32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32((v111^int32(base.Ui32(v123)>>(uint(v113)%32)))<<(uint(v116)%32))+uint32(_consts[1047])))
	v134 = v131 ^ v123<<(uint(v121)%32)
	v136 = v93 + v116
	v138 = v101 + v116
	if v138 != v80&int32(2046) {
		v93 = v136
		v94 = v134
		v101 = v138
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v140 = v136
	v141 = v134
	goto L35
L38:
	;
	goto L37
L39:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	v168 = *(*int32)(unsafe.Add(mBase, uint32((v160^int32(base.Ui32(v141)>>(uint(int32(24))%32)))<<(uint(int32(2))%32))+uint32(_consts[1047])))
	v172 = v168 ^ v141<<(uint(int32(8))%32)
	goto L41
L40:
	;
	v172 = v141
	goto L41
L41:
	;
	v194 = v172 ^ int32(-1)
	goto L32
L42:
	;
	v217 = v214 << (uint(int32(24)) % 32)
	if v214&int32(253) == int32(1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if v55 == v25-int32(1) {
		goto L3
	} else {
		goto L47
	}
L44:
	;
	if v217 == int32(33554432) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	if v217 != int32(67108864) {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+1)) = uint8(v214)
	if v217 != int32(67108864) {
		v241 = v53
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v231 = F_pq_getmsgint(m, v23, int32(2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L8
	} else {
		goto L49
	}
L49:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)) = uint16(v231)
	v241 = v53
	goto L21
L50:
	;
	goto L19
L51:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = int32(0)
	v282 = v276 + int32(8)
	F_findoprnd_recurse(m, v282, v21+int32(28), v25, v21+int32(27))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	if v289 != v25 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v25 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v295 = v282
	v296 = int32(0)
	v303 = v282 + v278*int32(12)
	goto L57
L55:
	;
	goto L56
L56:
	;
	F_pfree(m, v33)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L8
	} else {
		goto L67
	}
L57:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	if v313 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	goto L56
L59:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v33+v296<<(uint(int32(2))%32))))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v295)+8))
	v324 = v320&int32(4095) + int32(1)
	if v324 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v333 = v303
	goto L61
L61:
	;
	v337 = v296 + int32(1)
	if v337 != v25 {
		v295 = v295 + int32(12)
		v296 = v337
		v303 = v333
		goto L57
	} else {
		goto L66
	}
L62:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v295)+8))
	v333 = v326 + v327&int32(4095) + int32(1)
	goto L61
L63:
	;
	v325 = F__emscripten_memcpy_bulkmem(m, v303, v319, v324)
	mBase = m.M
	v326 = v325
	goto L65
L64:
	;
	v326 = v303
	goto L65
L65:
	;
	goto L62
L66:
	;
	goto L58
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v276))) = v275 << (uint(int32(2)) % 32)
	m.G0 = v21 + int32(32)
	return v276
L68:
	;
	F_errmsg_internal(m, int32(15278), int32(0))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(489270), int32(1241), int32(35667))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errmsg_internal(m, int32(236949), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(489270), int32(1274), int32(35667))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errmsg_internal(m, int32(326256), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(489270), int32(1277), int32(35667))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errmsg_internal(m, int32(459767), int32(0))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L8
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(489270), int32(1280), int32(35667))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v217 >> (uint(int32(24)) % 32)
	F_errmsg_internal(m, int32(473156), v21+int32(16))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(489270), int32(1309), int32(35667))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errmsg_internal(m, int32(424975), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L8
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(489270), int32(1311), int32(35667))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L8
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
	v452 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48))))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v452
	F_errmsg_internal(m, int32(482629), v21)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(489270), int32(1318), int32(35667))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errmsg_internal(m, int32(170405), int32(0))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(489270), int32(793), int32(422507))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L8
	} else {
		goto L91
	}
L91:
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
	F_errmsg_internal(m, int32(482629), v11)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(489270), int32(1215), int32(423132))
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
						v69 = F_cstring_to_text(m, int32(518996))
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
							v69 = F_cstring_to_text(m, int32(518996))
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
									v69 = F_cstring_to_text(m, int32(518996))
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
	F_errmsg(m, int32(323508), int32(0))
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
	F_errdetail_plural(m, int32(634721), int32(634607), v85, v12+int32(32))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(496992), int32(954), int32(323199))
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
	F_errmsg(m, int32(323508), int32(0))
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
	F_errdetail(m, int32(585554), v12+int32(16))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L11
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(496992), int32(970), int32(323199))
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
	F_errmsg(m, int32(323508), int32(0))
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
	F_errdetail(m, int32(635467), v12)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(496992), int32(978), int32(323199))
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
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v93 int32
	_ = v93
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v436 int32
	_ = v436
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v689 int32
	_ = v689
	var v701 int32
	_ = v701
	var v720 int32
	_ = v720
	var v737 int64
	_ = v737
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v762 int64
	_ = v762
	var v765 int64
	_ = v765
	var v769 int64
	_ = v769
	var v770 int64
	_ = v770
	var v782 int32
	_ = v782
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v873 int64
	_ = v873
	var v878 int32
	_ = v878
	var v887 int32
	_ = v887
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v907 int32
	_ = v907
	var v923 int32
	_ = v923
	var v932 int32
	_ = v932
	var v942 int32
	_ = v942
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int64
	_ = v983
	var v986 int32
	_ = v986
	var v988 int32
	_ = v988
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1005 int32
	_ = v1005
	var v1018 int64
	_ = v1018
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1058 int32
	_ = v1058
	var v1060 int32
	_ = v1060
	var v1062 int32
	_ = v1062
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1071 int32
	_ = v1071
	var v1077 int32
	_ = v1077
	var v1084 int32
	_ = v1084
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1172 int32
	_ = v1172
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1190 int32
	_ = v1190
	var v1196 int32
	_ = v1196
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1212 int32
	_ = v1212
	var v1215 int32
	_ = v1215
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1234 int32
	_ = v1234
	var v1243 int32
	_ = v1243
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1255 int32
	_ = v1255
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1279 int32
	_ = v1279
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1306 int32
	_ = v1306
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1315 int32
	_ = v1315
	var v1321 int32
	_ = v1321
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1390 int32
	_ = v1390
	var v1396 int32
	_ = v1396
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1409 int32
	_ = v1409
	var v1413 int32
	_ = v1413
	var v1416 int32
	_ = v1416
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1429 int32
	_ = v1429
	var v1430 int32
	_ = v1430
	var v1434 int32
	_ = v1434
	var v1440 int32
	_ = v1440
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1459 int32
	_ = v1459
	var v1471 int32
	_ = v1471
	var v1474 int32
	_ = v1474
	var v1478 int32
	_ = v1478
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1499 int32
	_ = v1499
	var v1507 int32
	_ = v1507
	var v1509 int32
	_ = v1509
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1549 int64
	_ = v1549
	var v1552 int64
	_ = v1552
	var v1556 int64
	_ = v1556
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1574 int64
	_ = v1574
	var v1577 int64
	_ = v1577
	var v1581 int64
	_ = v1581
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1602 int64
	_ = v1602
	var v1605 int64
	_ = v1605
	var v1609 int64
	_ = v1609
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1648 int64
	_ = v1648
	var v1650 int64
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1668 int32
	_ = v1668
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1712 int32
	_ = v1712
	var v1718 int32
	_ = v1718
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
	return v1718
L2:
	;
	v351 = v332 + int32(1)
	if base.Ui32(int32(512)) < base.Ui32(v351) {
		goto L61
	} else {
		goto L62
	}
L3:
	;
	v33 = F_strlen(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v31)+44)) = int32(0)
	v322 = l0
	v332 = v33
	v333 = l0 + v33
	goto L2
L4:
	;
	goto L5
L5:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v37 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v132&int32(255) == int32(0) {
		v1718 = v4
		goto L1
	} else {
		goto L27
	}
L7:
	;
	v130 = l0
	v132 = v37
	v133 = l0
	v140 = l0 - l0
	goto L6
L8:
	;
	goto L9
L9:
	;
	if v37 != int32(60) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v45 = v37
	v46 = l0
	goto L14
L11:
	;
	goto L12
L12:
	;
	v90 = l0 + int32(1)
	v93 = v90
	goto L19
L13:
	;
	v130 = l0
	v132 = v86
	v133 = v87
	v140 = v87 - l0
	goto L6
L14:
	;
	if base.Ui32(int32(252)) < base.Ui32((v45-int32(46))&int32(255)) {
		v86 = v45
		v87 = v46
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v86 = int32(0)
	v87 = v83
	goto L13
L16:
	;
	if base.Ui32(int32(-11)) < base.Ui32(base.I32_extend8_s(v45)-int32(58)) {
		v86 = v45
		v87 = v46
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v83 = v46 + int32(1)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v84 != 0 {
		v45 = v84
		v46 = v83
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93))))
	if v119 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v128 = v93 + int32(1)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	v130 = v90
	v132 = v129
	v133 = v128
	v140 = v93 - v90
	goto L6
L21:
	;
	v1718 = v4
	goto L1
L22:
	;
	goto L23
L23:
	;
	if v119 != int32(62) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v93 = v93 + int32(1)
	goto L19
L25:
	;
	goto L26
L26:
	;
	goto L20
L27:
	;
	v163 = v31 + int32(44)
	v164 = int32(0)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	switch v171 - int32(43) {
	case 0:
		goto L32
	default:
		v181 = v133
		v182 = v171
		v184 = v164
		goto L30
	case 2:
		v175 = int32(1)
		goto L31
	}
L28:
	;
	if v316 == int32(0) {
		v1718 = v4
		goto L1
	} else {
		goto L60
	}
L29:
	;
	goto L28
L30:
	;
	if base.Ui32(int32(9)) < base.Ui32(base.I32_extend8_s(v182)-int32(48)) {
		v316 = v164
		goto L29
	} else {
		goto L34
	}
L31:
	;
	v177 = v133 + int32(1)
	if v177 == int32(0) {
		v316 = v164
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v175 = int32(0)
	goto L31
L33:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	v181 = v177
	v182 = v180
	v184 = v175
	goto L30
L34:
	;
	v191 = v181
	v193 = v182
	v194 = int32(0)
	goto L35
L35:
	;
	v204 = base.I32_extend8_s(v193) + v194*int32(10) - int32(48)
	if int32(167) < v204 {
		v316 = v164
		goto L29
	} else {
		goto L37
	}
L36:
	;
	if v204 < int32(0) {
		v316 = v164
		goto L29
	} else {
		goto L39
	}
L37:
	;
	v208 = v191 + int32(1)
	v209 = int32(*(*int8)(unsafe.Add(mBase, uint32(v208))))
	if base.Ui32(v209-int32(48)) < base.Ui32(int32(10)) {
		v191 = v208
		v193 = v209
		v194 = v204
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v217 = v204 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v217
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
	if v219 != int32(58) {
		v301 = v208
		v306 = v217
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v184 != 0 {
		goto L57
	} else {
		goto L58
	}
L41:
	;
	v223 = v191 + int32(2)
	if v223 == int32(0) {
		v316 = v164
		goto L29
	} else {
		goto L42
	}
L42:
	;
	v226 = int32(*(*int8)(unsafe.Add(mBase, uint32(v223))))
	if base.Ui32(int32(9)) < base.Ui32(v226-int32(48)) {
		v316 = v164
		goto L29
	} else {
		goto L43
	}
L43:
	;
	v232 = v223
	v234 = int32(0)
	v235 = v226
	goto L44
L44:
	;
	v245 = base.I32_extend8_s(v235) + v234*int32(10) - int32(48)
	if int32(59) < v245 {
		v316 = v164
		goto L29
	} else {
		goto L46
	}
L45:
	;
	if v245 < int32(0) {
		v316 = v164
		goto L29
	} else {
		goto L48
	}
L46:
	;
	v249 = v232 + int32(1)
	v250 = int32(*(*int8)(unsafe.Add(mBase, uint32(v249))))
	if base.Ui32(v250-int32(48)) < base.Ui32(int32(10)) {
		v232 = v249
		v234 = v245
		v235 = v250
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v259 = v245*int32(60) + v217
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v259
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	if v261 != int32(58) {
		v301 = v249
		v306 = v259
		goto L40
	} else {
		goto L49
	}
L49:
	;
	v265 = v232 + int32(2)
	if v265 == int32(0) {
		v316 = v164
		goto L29
	} else {
		goto L50
	}
L50:
	;
	v268 = int32(*(*int8)(unsafe.Add(mBase, uint32(v265))))
	if base.Ui32(int32(9)) < base.Ui32(v268-int32(48)) {
		v316 = v164
		goto L29
	} else {
		goto L51
	}
L51:
	;
	v274 = v265
	v276 = int32(0)
	v277 = v268
	goto L52
L52:
	;
	v287 = base.I32_extend8_s(v277) + v276*int32(10) - int32(48)
	if int32(60) < v287 {
		v316 = v164
		goto L29
	} else {
		goto L54
	}
L53:
	;
	if v287 < int32(0) {
		v316 = v164
		goto L29
	} else {
		goto L56
	}
L54:
	;
	v291 = v274 + int32(1)
	v292 = int32(*(*int8)(unsafe.Add(mBase, uint32(v291))))
	if base.Ui32(v292-int32(48)) < base.Ui32(int32(10)) {
		v274 = v291
		v276 = v287
		v277 = v292
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v299 = v287 + v259
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = v299
	v301 = v291
	v306 = v299
	goto L40
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v163))) = int32(0) - v306
	goto L59
L58:
	;
	goto L59
L59:
	;
	v316 = v301
	goto L29
L60:
	;
	v322 = v130
	v332 = v140
	v333 = v316
	goto L2
L61:
	;
	v1718 = v4
	goto L1
L62:
	;
	goto L63
L63:
	;
	v354 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v354)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v354
	v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333))))
	if v358 != 0 {
		goto L69
	} else {
		goto L70
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v479
	v1700 = l1 + int32(22120)
	if v332 != 0 {
		goto L300
	} else {
		goto L301
	}
L65:
	;
	if v1655-v982 < int32(401) {
		goto L64
	} else {
		goto L298
	}
L66:
	;
	v1648 = *(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1194])))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1195]))) = v1648
	v1650 = *(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1196])))
	*(*int64)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1197]))) = v1650
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(1)
	goto L64
L67:
	;
	v986 = l1 + int32(16024)
	v988 = l1 + int32(24)
	v995 = v981
	v996 = v982
	v999 = v4
	v1005 = v720 + int32(400)
	v1018 = v983
	goto L176
L68:
	;
	if v720 < int32(2147483248) {
		v981 = v758
		v982 = v741
		v983 = v737
		goto L67
	} else {
		goto L175
	}
L69:
	;
	if v358 != int32(60) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	goto L71
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1)+4)) = int64(4294967296)
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v956 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1198]))) = v956
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1199]))) = uint16(v956)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1197]))) = v956
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1200]))) = uint8(v956)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1195]))) = v956 - v955
	v969 = l1 + int32(22120)
	if v332 != 0 {
		goto L172
	} else {
		goto L173
	}
L72:
	;
	if v474 == int32(0) {
		v1718 = v4
		goto L1
	} else {
		goto L90
	}
L73:
	;
	v363 = v358
	v364 = v333
	goto L76
L74:
	;
	goto L75
L75:
	;
	v407 = v333 + int32(1)
	v410 = v407
	goto L82
L76:
	;
	if base.Ui32(int32(252)) < base.Ui32((v363-int32(46))&int32(255)) {
		v404 = v364
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v449 = v404
	v457 = v333
	v474 = v404 - v333
	goto L72
L78:
	;
	goto L77
L79:
	;
	if base.Ui32(int32(-11)) < base.Ui32(base.I32_extend8_s(v363)-int32(58)) {
		v404 = v364
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v401 = v364 + int32(1)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v401))))
	if v402 != 0 {
		v363 = v402
		v364 = v401
		goto L76
	} else {
		goto L81
	}
L81:
	;
	v404 = v401
	goto L78
L82:
	;
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	if v436 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v449 = v410 + int32(1)
	v457 = v407
	v474 = v410 - v407
	goto L72
L84:
	;
	v1718 = v4
	goto L1
L85:
	;
	goto L86
L86:
	;
	if v436 != int32(62) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v410 = v410 + int32(1)
	goto L82
L88:
	;
	goto L89
L89:
	;
	goto L83
L90:
	;
	v479 = v474 + v332 + int32(2)
	if base.Ui32(int32(512)) < base.Ui32(v479) {
		v1718 = v4
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	switch v482 - int32(44) {
	case 0, 15:
		goto L93
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		goto L94
	default:
		goto L95
	}
L92:
	;
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v649))))
	if v651 != 0 {
		goto L134
	} else {
		goto L135
	}
L93:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+40)) = v645 - int32(3600)
	v649 = v449
	goto L92
L94:
	;
	v488 = v31 + int32(40)
	v489 = int32(0)
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449))))
	switch v496 - int32(43) {
	case 0:
		goto L101
	default:
		v506 = v449
		v507 = v496
		v509 = v489
		goto L99
	case 2:
		v500 = int32(1)
		goto L100
	}
L95:
	;
	if v482 == int32(0) {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	if v641 != 0 {
		v649 = v641
		goto L92
	} else {
		goto L129
	}
L98:
	;
	goto L97
L99:
	;
	if base.Ui32(int32(9)) < base.Ui32(base.I32_extend8_s(v507)-int32(48)) {
		v641 = v489
		goto L98
	} else {
		goto L103
	}
L100:
	;
	v502 = v449 + int32(1)
	if v502 == int32(0) {
		v641 = v489
		goto L98
	} else {
		goto L102
	}
L101:
	;
	v500 = int32(0)
	goto L100
L102:
	;
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	v506 = v502
	v507 = v505
	v509 = v500
	goto L99
L103:
	;
	v516 = v506
	v518 = v507
	v519 = int32(0)
	goto L104
L104:
	;
	v529 = base.I32_extend8_s(v518) + v519*int32(10) - int32(48)
	if int32(167) < v529 {
		v641 = v489
		goto L98
	} else {
		goto L106
	}
L105:
	;
	if v529 < int32(0) {
		v641 = v489
		goto L98
	} else {
		goto L108
	}
L106:
	;
	v533 = v516 + int32(1)
	v534 = int32(*(*int8)(unsafe.Add(mBase, uint32(v533))))
	if base.Ui32(v534-int32(48)) < base.Ui32(int32(10)) {
		v516 = v533
		v518 = v534
		v519 = v529
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v542 = v529 * int32(3600)
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = v542
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v533))))
	if v544 != int32(58) {
		v626 = v533
		v631 = v542
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v509 != 0 {
		goto L126
	} else {
		goto L127
	}
L110:
	;
	v548 = v516 + int32(2)
	if v548 == int32(0) {
		v641 = v489
		goto L98
	} else {
		goto L111
	}
L111:
	;
	v551 = int32(*(*int8)(unsafe.Add(mBase, uint32(v548))))
	if base.Ui32(int32(9)) < base.Ui32(v551-int32(48)) {
		v641 = v489
		goto L98
	} else {
		goto L112
	}
L112:
	;
	v557 = v548
	v559 = int32(0)
	v560 = v551
	goto L113
L113:
	;
	v570 = base.I32_extend8_s(v560) + v559*int32(10) - int32(48)
	if int32(59) < v570 {
		v641 = v489
		goto L98
	} else {
		goto L115
	}
L114:
	;
	if v570 < int32(0) {
		v641 = v489
		goto L98
	} else {
		goto L117
	}
L115:
	;
	v574 = v557 + int32(1)
	v575 = int32(*(*int8)(unsafe.Add(mBase, uint32(v574))))
	if base.Ui32(v575-int32(48)) < base.Ui32(int32(10)) {
		v557 = v574
		v559 = v570
		v560 = v575
		goto L113
	} else {
		goto L116
	}
L116:
	;
	goto L114
L117:
	;
	v584 = v570*int32(60) + v542
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = v584
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v574))))
	if v586 != int32(58) {
		v626 = v574
		v631 = v584
		goto L109
	} else {
		goto L118
	}
L118:
	;
	v590 = v557 + int32(2)
	if v590 == int32(0) {
		v641 = v489
		goto L98
	} else {
		goto L119
	}
L119:
	;
	v593 = int32(*(*int8)(unsafe.Add(mBase, uint32(v590))))
	if base.Ui32(int32(9)) < base.Ui32(v593-int32(48)) {
		v641 = v489
		goto L98
	} else {
		goto L120
	}
L120:
	;
	v599 = v590
	v601 = int32(0)
	v602 = v593
	goto L121
L121:
	;
	v612 = base.I32_extend8_s(v602) + v601*int32(10) - int32(48)
	if int32(60) < v612 {
		v641 = v489
		goto L98
	} else {
		goto L123
	}
L122:
	;
	if v612 < int32(0) {
		v641 = v489
		goto L98
	} else {
		goto L125
	}
L123:
	;
	v616 = v599 + int32(1)
	v617 = int32(*(*int8)(unsafe.Add(mBase, uint32(v616))))
	if base.Ui32(v617-int32(48)) < base.Ui32(int32(10)) {
		v599 = v616
		v601 = v612
		v602 = v617
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	v624 = v612 + v584
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = v624
	v626 = v616
	v631 = v624
	goto L109
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v488))) = int32(0) - v631
	goto L128
L127:
	;
	goto L128
L128:
	;
	v641 = v626
	goto L98
L129:
	;
	v1718 = v4
	goto L1
L130:
	;
	v923 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1199]))) = uint16(v923)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1197]))) = v923
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1200]))) = uint8(v923)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1195]))) = v923 - v907
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1201]))) = uint16(v923)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1196]))) = v351
	v942 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1202]))) = uint8(v942)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1198]))) = v923
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1194]))) = v923 - v932
	goto L64
L131:
	;
	v782 = v656
	goto L157
L132:
	;
	v665 = F_getrule(m, v652+int32(1), v31+int32(20))
	mBase = m.M
	if v665 == int32(0) {
		v1718 = v4
		goto L1
	} else {
		goto L139
	}
L133:
	;
	if v653 != 0 {
		v1718 = v4
		goto L1
	} else {
		goto L137
	}
L134:
	;
	v652 = v649
	goto L136
L135:
	;
	v652 = int32(551746)
	goto L136
L136:
	;
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652))))
	switch v653 - int32(44) {
	case 0, 15:
		goto L132
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
		v1718 = v4
		goto L1
	default:
		goto L133
	}
L137:
	;
	v656 = int32(0)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v656 < v657 {
		goto L131
	} else {
		goto L138
	}
L138:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v907 = v660
	goto L130
L139:
	;
	v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v665))))
	if v668 != int32(44) {
		v1718 = v4
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v673 = F_getrule(m, v665+int32(1), v31)
	mBase = m.M
	if v673 == int32(0) {
		v1718 = v4
		goto L1
	} else {
		goto L141
	}
L141:
	;
	v676 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
	if v676 != 0 {
		v1718 = v4
		goto L1
	} else {
		goto L142
	}
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(2)
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v680 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1199]))) = uint16(v680)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1197]))) = v680
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1200]))) = uint8(v680)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1195]))) = v680 - v679
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v31)+40))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1198]))) = v680
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1201]))) = uint16(v680)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1196]))) = v351
	v701 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1202]))) = uint8(v701)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[1194]))) = v680 - v689
	v720 = int32(1970)
	v737 = int64(0)
	goto L143
L143:
	;
	v741 = v720 - int32(1)
	if v741&int32(3) != 0 {
		v751 = int32(0)
		goto L146
	} else {
		goto L147
	}
L144:
	;
	v981 = int32(0)
	v982 = int32(1770)
	v983 = v770
	goto L67
L145:
	;
	v770 = v737 + v769
	if base.Ui32(int32(1771)) < base.Ui32(v720) {
		v720 = v741
		v737 = v770
		goto L143
	} else {
		goto L154
	}
L146:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v751<<(uint(int32(2))%32))+uint32(_consts[1203])))
	v758 = v756 * int32(-86400)
	if v758 < int32(0) {
		goto L149
	} else {
		goto L150
	}
L147:
	;
	v746 = base.I32_rem_u_s(v741, int32(100))
	if v746 != 0 {
		v751 = int32(1)
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v748 = base.I32_rem_u_s(v741, int32(400))
	v751 = base.B2i32(v748 == int32(0))
	goto L146
L149:
	;
	v762 = base.I64_extend_i32_s(v758)
	if int64(-9223372036854775807-1)-v762 <= v737 {
		v769 = v762
		goto L145
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v765 = base.I64_extend_i32_u(v758)
	if v765^int64(9223372036854775807) < v737 {
		goto L68
	} else {
		goto L153
	}
L152:
	;
	goto L68
L153:
	;
	v769 = v765
	goto L145
L154:
	;
	goto L144
L155:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v31)+44))
	v835 = int32(0)
	v841 = v823
	goto L161
L156:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v811)))
	v823 = int32(0) - v820
	goto L155
L157:
	;
	v808 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v782+(l1+int32(16024))))))
	v811 = l1 + int32(18024) + v808<<(uint(int32(4))%32)
	v812 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v811)+4)))
	if v812 == int32(0) {
		goto L156
	} else {
		goto L159
	}
L158:
	;
	v823 = int32(0)
	goto L155
L159:
	;
	v816 = v782 + int32(1)
	if v816 != v657 {
		v782 = v816
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	v860 = v835 + (l1 + int32(16024))
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860))))
	v864 = l1 + int32(18024) + v861<<(uint(int32(4))%32)
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v860))) = uint8(v865)
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+13)))
	if v867 == int32(0) {
		goto L165
	} else {
		goto L166
	}
L162:
	;
	v907 = v830
	goto L130
L163:
	;
	v892 = v835 + int32(1)
	v893 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v892 < v893 {
		v835 = v892
		v841 = v890
		goto L161
	} else {
		goto L170
	}
L164:
	;
	v887 = *(*int32)(unsafe.Add(mBase, uint32(v864)))
	v890 = int32(0) - v887
	goto L163
L165:
	;
	v872 = l1 + int32(24) + v835<<(uint(int32(3))%32)
	v873 = *(*int64)(unsafe.Add(mBase, uint32(v872)))
	*(*int64)(unsafe.Add(mBase, uint32(v872))) = v873 + base.I64_extend_i32_s(v830-v841)
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864)+4)))
	if v878&int32(1) == int32(0) {
		goto L164
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	if v865&int32(1) != 0 {
		v890 = v841
		goto L163
	} else {
		goto L169
	}
L168:
	;
	v890 = v841
	goto L163
L169:
	;
	goto L164
L170:
	;
	goto L162
L171:
	;
	v973 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v971+v332))) = uint8(v973)
	v1718 = int32(1)
	goto L1
L172:
	;
	v970 = F__emscripten_memcpy_bulkmem(m, v969, v322, v332)
	mBase = m.M
	v971 = v970
	goto L174
L173:
	;
	v971 = v969
	goto L174
L174:
	;
	goto L171
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(0)
	goto L66
L176:
	;
	v1021 = v31 + int32(20)
	v1022 = int32(0)
	if v996&int32(3) != 0 {
		v1040 = v1022
		goto L179
	} else {
		goto L180
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v1591
	if v1591 != 0 {
		v1655 = v1616
		goto L65
	} else {
		goto L297
	}
L178:
	;
	v1266 = int32(0)
	if v996&int32(3) != 0 {
		v1284 = v1266
		goto L219
	} else {
		goto L220
	}
L179:
	;
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(v1021)))
	switch v1041 {
	case 0:
		goto L185
	case 1:
		goto L184
	case 2:
		goto L183
	default:
		v1255 = v1022
		goto L182
	}
L180:
	;
	v1035 = base.I32_rem_s(v996, int32(100))
	if v1035 != 0 {
		v1040 = int32(1)
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v1037 = base.I32_rem_s(v996, int32(400))
	v1040 = base.B2i32(v1037 == int32(0))
	goto L179
L182:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+16))
	v1265 = v1263 + (v679 + v1255)
	goto L178
L183:
	;
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+12))
	v1058 = v996 - base.B2i32(v1055 < int32(3))
	v1060 = base.I32_div_s(v1058, int32(400))
	v1062 = base.I32_rem_s(v1058, int32(100))
	v1065 = base.I32_div_s(v1058, int32(-100))
	v1066 = int32(1)
	v1071 = base.I32_div_s(base.I32_extend8_s(v1062), int32(4))
	v1077 = base.I32_rem_s(v1055+int32(9), int32(12))
	v1084 = base.I32_div_s(base.I32_extend16_s(v1077*int32(26)+int32(24)), int32(10))
	v1089 = int32(7)
	v1090 = base.I32_rem_s(v1060+v1062+v1065<<(uint(v1066)%32)+base.I32_extend8_s(v1071)+base.I32_extend16_s(v1084+v1066), v1089)
	if v1090 < int32(0) {
		goto L192
	} else {
		goto L193
	}
L184:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	v1255 = v1051 * int32(86400)
	goto L182
L185:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+4))
	v1043 = int32(86400)
	v1044 = v1042 * v1043
	v1046 = v1044 - v1043
	if int32(59) < v1042 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v1049 = v1044
	goto L188
L187:
	;
	v1049 = v1046
	goto L188
L188:
	;
	if v1040 != 0 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v1050 = v1049
	goto L191
L190:
	;
	v1050 = v1046
	goto L191
L191:
	;
	v1255 = v1050
	goto L182
L192:
	;
	v1095 = v1090 + v1089
	goto L194
L193:
	;
	v1095 = v1090
	goto L194
L194:
	;
	v1096 = v1054 - v1095
	if v1096 < int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	v1101 = v1096 + int32(7)
	goto L197
L196:
	;
	v1101 = v1096
	goto L197
L197:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1021)+8))
	if v1102 <= int32(1) {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v1152 = v1140 * int32(86400)
	if v1146 <= int32(0) {
		v1255 = v1152
		goto L182
	} else {
		goto L206
	}
L199:
	;
	v1140 = v1101
	v1146 = v1055 - int32(1)
	goto L198
L200:
	;
	goto L201
L201:
	;
	v1109 = int32(1)
	v1110 = v1055 - v1109
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1040*int32(48)+v1110<<(uint(int32(2))%32))+uint32(_consts[1204])))
	v1117 = int32(7)
	v1123 = v1101
	v1126 = v1109
	goto L202
L202:
	;
	v1135 = v1123 + int32(7)
	if v1116 <= v1135 {
		v1140 = v1123
		v1146 = v1110
		goto L198
	} else {
		goto L204
	}
L203:
	;
	v1140 = v1101 + v1102*v1117 - v1117
	v1146 = v1110
	goto L198
L204:
	;
	v1138 = v1126 + int32(1)
	if v1138 != v1102 {
		v1123 = v1135
		v1126 = v1138
		goto L202
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	v1155 = int32(3)
	v1156 = v1146 & v1155
	if base.Ui32(v1055-int32(2)) < base.Ui32(v1155) {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	if v1156 == int32(0) {
		v1255 = v1215
		goto L182
	} else {
		goto L214
	}
L208:
	;
	v1212 = int32(0)
	v1215 = v1152
	goto L207
L209:
	;
	goto L210
L210:
	;
	v1165 = int32(0)
	v1169 = v1165
	v1172 = v1152
	v1178 = v1165
	goto L211
L211:
	;
	v1182 = v1040*int32(48) + v1169<<(uint(int32(2))%32)
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+uint32(_consts[1205])))
	v1186 = int32(86400)
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+uint32(_consts[1204])))
	v1196 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+uint32(_consts[1206])))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(v1182)+uint32(_consts[1207])))
	v1206 = v1185*v1186 + (v1190*v1186 + v1172 + v1196*v1186 + v1202*v1186)
	v1207 = int32(4)
	v1208 = v1169 + v1207
	v1210 = v1178 + v1207
	if v1210 != v1146&int32(2147483644) {
		v1169 = v1208
		v1172 = v1206
		v1178 = v1210
		goto L211
	} else {
		goto L213
	}
L212:
	;
	v1212 = v1208
	v1215 = v1206
	goto L207
L213:
	;
	goto L212
L214:
	;
	v1227 = v1212
	v1230 = v1215
	v1234 = int32(0)
	goto L215
L215:
	;
	v1243 = *(*int32)(unsafe.Add(mBase, uint32(v1040*int32(48)+v1227<<(uint(int32(2))%32))+uint32(_consts[1204])))
	v1246 = v1243*int32(86400) + v1230
	v1247 = int32(1)
	v1250 = v1234 + v1247
	if v1250 != v1156 {
		v1227 = v1227 + v1247
		v1230 = v1246
		v1234 = v1250
		goto L215
	} else {
		goto L217
	}
L216:
	;
	v1255 = v1246
	goto L182
L217:
	;
	goto L216
L218:
	;
	if v996&int32(3) != 0 {
		v1520 = int32(0)
		goto L258
	} else {
		goto L259
	}
L219:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	switch v1285 {
	case 0:
		goto L225
	case 1:
		goto L224
	case 2:
		goto L223
	default:
		v1499 = v1266
		goto L222
	}
L220:
	;
	v1279 = base.I32_rem_s(v996, int32(100))
	if v1279 != 0 {
		v1284 = int32(1)
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v1281 = base.I32_rem_s(v996, int32(400))
	v1284 = base.B2i32(v1281 == int32(0))
	goto L219
L222:
	;
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v1509 = v1507 + (v689 + v1499)
	goto L218
L223:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v1302 = v996 - base.B2i32(v1299 < int32(3))
	v1304 = base.I32_div_s(v1302, int32(400))
	v1306 = base.I32_rem_s(v1302, int32(100))
	v1309 = base.I32_div_s(v1302, int32(-100))
	v1310 = int32(1)
	v1315 = base.I32_div_s(base.I32_extend8_s(v1306), int32(4))
	v1321 = base.I32_rem_s(v1299+int32(9), int32(12))
	v1328 = base.I32_div_s(base.I32_extend16_s(v1321*int32(26)+int32(24)), int32(10))
	v1333 = int32(7)
	v1334 = base.I32_rem_s(v1304+v1306+v1309<<(uint(v1310)%32)+base.I32_extend8_s(v1315)+base.I32_extend16_s(v1328+v1310), v1333)
	if v1334 < int32(0) {
		goto L232
	} else {
		goto L233
	}
L224:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1499 = v1295 * int32(86400)
	goto L222
L225:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v1287 = int32(86400)
	v1288 = v1286 * v1287
	v1290 = v1288 - v1287
	if int32(59) < v1286 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1293 = v1288
	goto L228
L227:
	;
	v1293 = v1290
	goto L228
L228:
	;
	if v1284 != 0 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v1294 = v1293
	goto L231
L230:
	;
	v1294 = v1290
	goto L231
L231:
	;
	v1499 = v1294
	goto L222
L232:
	;
	v1339 = v1334 + v1333
	goto L234
L233:
	;
	v1339 = v1334
	goto L234
L234:
	;
	v1340 = v1298 - v1339
	if v1340 < int32(0) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1345 = v1340 + int32(7)
	goto L237
L236:
	;
	v1345 = v1340
	goto L237
L237:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	if v1346 <= int32(1) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1396 = v1384 * int32(86400)
	if v1390 <= int32(0) {
		v1499 = v1396
		goto L222
	} else {
		goto L246
	}
L239:
	;
	v1384 = v1345
	v1390 = v1299 - int32(1)
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1353 = int32(1)
	v1354 = v1299 - v1353
	v1360 = *(*int32)(unsafe.Add(mBase, uint32(v1284*int32(48)+v1354<<(uint(int32(2))%32))+uint32(_consts[1204])))
	v1361 = int32(7)
	v1367 = v1345
	v1370 = v1353
	goto L242
L242:
	;
	v1379 = v1367 + int32(7)
	if v1360 <= v1379 {
		v1384 = v1367
		v1390 = v1354
		goto L238
	} else {
		goto L244
	}
L243:
	;
	v1384 = v1345 + v1346*v1361 - v1361
	v1390 = v1354
	goto L238
L244:
	;
	v1382 = v1370 + int32(1)
	if v1382 != v1346 {
		v1367 = v1379
		v1370 = v1382
		goto L242
	} else {
		goto L245
	}
L245:
	;
	goto L243
L246:
	;
	v1399 = int32(3)
	v1400 = v1390 & v1399
	if base.Ui32(v1299-int32(2)) < base.Ui32(v1399) {
		goto L248
	} else {
		goto L249
	}
L247:
	;
	if v1400 == int32(0) {
		v1499 = v1459
		goto L222
	} else {
		goto L254
	}
L248:
	;
	v1456 = int32(0)
	v1459 = v1396
	goto L247
L249:
	;
	goto L250
L250:
	;
	v1409 = int32(0)
	v1413 = v1409
	v1416 = v1396
	v1422 = v1409
	goto L251
L251:
	;
	v1426 = v1284*int32(48) + v1413<<(uint(int32(2))%32)
	v1429 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+uint32(_consts[1205])))
	v1430 = int32(86400)
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+uint32(_consts[1204])))
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+uint32(_consts[1206])))
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(v1426)+uint32(_consts[1207])))
	v1450 = v1429*v1430 + (v1434*v1430 + v1416 + v1440*v1430 + v1446*v1430)
	v1451 = int32(4)
	v1452 = v1413 + v1451
	v1454 = v1422 + v1451
	if v1454 != v1390&int32(2147483644) {
		v1413 = v1452
		v1416 = v1450
		v1422 = v1454
		goto L251
	} else {
		goto L253
	}
L252:
	;
	v1456 = v1452
	v1459 = v1450
	goto L247
L253:
	;
	goto L252
L254:
	;
	v1471 = v1456
	v1474 = v1459
	v1478 = int32(0)
	goto L255
L255:
	;
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1284*int32(48)+v1471<<(uint(int32(2))%32))+uint32(_consts[1204])))
	v1490 = v1487*int32(86400) + v1474
	v1491 = int32(1)
	v1494 = v1478 + v1491
	if v1494 != v1400 {
		v1471 = v1471 + v1491
		v1474 = v1490
		v1478 = v1494
		goto L255
	} else {
		goto L257
	}
L256:
	;
	v1499 = v1490
	goto L222
L257:
	;
	goto L256
L258:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(v1520<<(uint(int32(2))%32))+uint32(_consts[1203])))
	v1527 = v1525 * int32(86400)
	v1528 = base.B2i32(v1509 < v1265)
	if v1528 == int32(0) {
		goto L262
	} else {
		goto L263
	}
L259:
	;
	v1515 = base.I32_rem_u_s(v996, int32(100))
	if v1515 != 0 {
		v1520 = int32(1)
		goto L258
	} else {
		goto L260
	}
L260:
	;
	v1517 = base.I32_rem_u_s(v996, int32(400))
	v1520 = base.B2i32(v1517 == int32(0))
	goto L258
L261:
	;
	v1598 = v995 + v1527
	if v1598 < int32(0) {
		goto L291
	} else {
		goto L292
	}
L262:
	;
	if v1509 <= v1265 {
		v1591 = v999
		v1593 = v1005
		goto L261
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	if int32(1999) <= v999 {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	if v1527+(v679-v689) <= v1509-v1265 {
		v1591 = v999
		v1593 = v1005
		goto L261
	} else {
		goto L266
	}
L266:
	;
	goto L264
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v999
	v1655 = v996
	goto L65
L268:
	;
	goto L269
L269:
	;
	if v1265 < v1509 {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1539 = v1509
	goto L272
L271:
	;
	v1539 = v1265
	goto L272
L272:
	;
	v1542 = v988 + v999<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1542))) = v1018
	if v1509 < v1265 {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	v1568 = v988 + v1564<<(uint(int32(3))%32)
	*(*int64)(unsafe.Add(mBase, uint32(v1568))) = v1018
	v1570 = v995 + v1539
	if v1570 < int32(0) {
		goto L284
	} else {
		goto L285
	}
L274:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1542))) = v1018 + v1556
	*(*uint8)(unsafe.Add(mBase, uint32(v999+v986))) = uint8(base.B2i32(v1265 <= v1509))
	v1564 = v999 + int32(1)
	goto L273
L275:
	;
	v1544 = v1509
	goto L277
L276:
	;
	v1544 = v1265
	goto L277
L277:
	;
	v1545 = v1544 + v995
	if v1545 < int32(0) {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1549 = base.I64_extend_i32_s(v1545)
	if int64(-9223372036854775807-1)-v1549 <= v1018 {
		v1556 = v1549
		goto L274
	} else {
		goto L281
	}
L279:
	;
	goto L280
L280:
	;
	v1552 = base.I64_extend_i32_u(v1545)
	if v1552^int64(9223372036854775807) < v1018 {
		v1564 = v999
		goto L273
	} else {
		goto L282
	}
L281:
	;
	v1564 = v999
	goto L273
L282:
	;
	v1556 = v1552
	goto L274
L283:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1568))) = v1018 + v1581
	*(*uint8)(unsafe.Add(mBase, uint32(v1564+v986))) = uint8(v1528)
	v1591 = v1564 + int32(1)
	v1593 = v996 + int32(401)
	goto L261
L284:
	;
	v1574 = base.I64_extend_i32_s(v1570)
	if int64(-9223372036854775807-1)-v1574 <= v1018 {
		v1581 = v1574
		goto L283
	} else {
		goto L287
	}
L285:
	;
	goto L286
L286:
	;
	v1577 = base.I64_extend_i32_u(v1570)
	if v1577^int64(9223372036854775807) < v1018 {
		v1591 = v1564
		v1593 = v1005
		goto L261
	} else {
		goto L288
	}
L287:
	;
	v1591 = v1564
	v1593 = v1005
	goto L261
L288:
	;
	v1581 = v1577
	goto L283
L289:
	;
	goto L177
L290:
	;
	v1613 = v996 + int32(1)
	if v1613 < v1593 {
		v995 = int32(0)
		v996 = v1613
		v999 = v1591
		v1005 = v1593
		v1018 = v1018 + v1609
		goto L176
	} else {
		goto L296
	}
L291:
	;
	v1602 = base.I64_extend_i32_s(v1598)
	if int64(-9223372036854775807-1)-v1602 <= v1018 {
		v1609 = v1602
		goto L290
	} else {
		goto L294
	}
L292:
	;
	goto L293
L293:
	;
	v1605 = base.I64_extend_i32_u(v1598)
	if v1605^int64(9223372036854775807) < v1018 {
		v1616 = v996
		goto L289
	} else {
		goto L295
	}
L294:
	;
	v1616 = v996
	goto L289
L295:
	;
	v1609 = v1605
	goto L290
L296:
	;
	v1616 = v1613
	goto L289
L297:
	;
	goto L66
L298:
	;
	v1668 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v1668)
	goto L64
L299:
	;
	v1703 = v1702 + v332
	v1704 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1703))) = uint8(v1704)
	v1706 = int32(1)
	v1708 = v1703 + v1706
	if v474 != 0 {
		goto L304
	} else {
		goto L305
	}
L300:
	;
	v1701 = F__emscripten_memcpy_bulkmem(m, v1700, v322, v332)
	mBase = m.M
	v1702 = v1701
	goto L302
L301:
	;
	v1702 = v1700
	goto L302
L302:
	;
	goto L299
L303:
	;
	v1712 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1710+v474))) = uint8(v1712)
	v1718 = v1706
	goto L1
L304:
	;
	v1709 = F__emscripten_memcpy_bulkmem(m, v1708, v457, v474)
	mBase = m.M
	v1710 = v1709
	goto L306
L305:
	;
	v1710 = v1708
	goto L306
L306:
	;
	goto L303
}
