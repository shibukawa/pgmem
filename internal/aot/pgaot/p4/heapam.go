package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_heapam_index_fetch_end(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		F_ReleaseBuffer(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
			F_pfree(m, l0)
			mBase = m.M
			v9 = m.ExcPending
			if v9 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			return
		}
	}
}
func F_heapam_index_fetch_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		F_ReleaseBuffer(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
			return
		}
	} else {
		return
	}
}
func F_heapam_relation_copy_data(m *base.Module, l0 int32, l1 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
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
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int64
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int64
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	F_FlushRelationBuffers(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v13 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+118)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+112)) = v16
	v21 = F_RelationCreateStorage(m, v8+int32(112), v13, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v49 = v23
	goto L6
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+104)) = v25
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+96)) = v27
	v31 = F_smgropen(m, v8+int32(96), v24)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v52 = int32(*(*int8)(unsafe.Add(mBase, uint32(v51)+118)))
	F_RelationCopyStorage(m, v49, v21, int32(0), v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v31
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	if v35 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v49 = v47
	goto L6
L9:
	;
	v43 = v35
	goto L11
L10:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v37
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+72))
	v43 = v41
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+72)) = v43 + int32(1)
	goto L8
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v55 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v81 = v55
	goto L15
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = v59
	v63 = F_smgropen(m, v8+int32(80), v56)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v83 = F_smgrexists(m, v81, int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v63
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	if v67 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v81 = v79
	goto L15
L18:
	;
	v75 = v67
	goto L20
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+76))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v69
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)+72))
	v75 = v73
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+72)) = v75 + int32(1)
	goto L17
L21:
	;
	if v83 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_smgrcreate(m, v21, int32(1), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v129 != 0 {
		goto L39
	} else {
		goto L40
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+118)))
	if v90 == int32(112) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_log_smgrcreate(m, l1, int32(1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v96 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v122 = v96
	goto L32
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v98
	v100 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v100
	v104 = F_smgropen(m, v8-int32(-64), v97)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v125 = int32(*(*int8)(unsafe.Add(mBase, uint32(v124)+118)))
	F_RelationCopyStorage(m, v122, v21, int32(1), v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L38
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v104
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+72))
	if v108 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v122 = v120
	goto L32
L35:
	;
	v116 = v108
	goto L37
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v104)+76))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v104)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v104)+72))
	v116 = v114
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v104)+72)) = v116 + int32(1)
	goto L34
L38:
	;
	goto L24
L39:
	;
	v155 = v129
	goto L41
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v133
	v137 = F_smgropen(m, v8+int32(48), v130)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	v157 = F_smgrexists(m, v155, int32(2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L47
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v137
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)+72))
	if v141 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v155 = v153
	goto L41
L44:
	;
	v149 = v141
	goto L46
L45:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v137)+76))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v137)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v142)+4)) = v143
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v137)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v137)+72))
	v149 = v147
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+72)) = v149 + int32(1)
	goto L43
L47:
	;
	if v157 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	F_smgrcreate(m, v21, int32(2), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v203 != 0 {
		goto L65
	} else {
		goto L66
	}
L51:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+118)))
	if v164 == int32(112) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_log_smgrcreate(m, l1, int32(2))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v170 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	v196 = v170
	goto L58
L57:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v172
	v174 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v174
	v178 = F_smgropen(m, v8+int32(32), v171)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v199 = int32(*(*int8)(unsafe.Add(mBase, uint32(v198)+118)))
	F_RelationCopyStorage(m, v196, v21, int32(2), v199)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L64
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v178
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178)+72))
	if v182 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v196 = v194
	goto L58
L61:
	;
	v190 = v182
	goto L63
L62:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v178)+76))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v178)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v178)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v178)+72))
	v190 = v188
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+72)) = v190 + int32(1)
	goto L60
L64:
	;
	goto L50
L65:
	;
	v229 = v203
	goto L67
L66:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v207
	v211 = F_smgropen(m, v8+int32(16), v204)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v231 = F_smgrexists(m, v229, int32(3))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L73
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v211
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)+72))
	if v215 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v229 = v227
	goto L67
L70:
	;
	v223 = v215
	goto L72
L71:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)+76))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v211)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+4)) = v217
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v211)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v211)+72))
	v223 = v221
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v211)+72)) = v223 + int32(1)
	goto L69
L73:
	;
	if v231 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	F_smgrcreate(m, v21, int32(3), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	F_RelationDropStorage(m, l0)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L90
	}
L77:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+118)))
	switch v238 - int32(112) {
	case 0, 5:
		goto L79
	default:
		goto L78
	}
L78:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v244 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	F_log_smgrcreate(m, l1, int32(3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v268 = v244
	goto L83
L82:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v246
	v248 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v248
	v250 = F_smgropen(m, v8, v245)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v271 = int32(*(*int8)(unsafe.Add(mBase, uint32(v270)+118)))
	F_RelationCopyStorage(m, v268, v21, int32(3), v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L89
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v250
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250)+72))
	if v254 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v268 = v266
	goto L83
L86:
	;
	v262 = v254
	goto L88
L87:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v250)+76))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v250)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = v256
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v250)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v258
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v250)+72))
	v262 = v260
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v250)+72)) = v262 + int32(1)
	goto L85
L89:
	;
	goto L76
L90:
	;
	F_smgrclose(m, v21)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	m.G0 = v8 + int32(128)
	return
}
func F_heapam_relation_set_new_filelocator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_heapam_relation_set_new_filelocator[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v11
	v13 = F_GetOldestMultiXactId(m)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v13
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v16
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v18
		v21 = F_RelationCreateStorage(m, v8, l2, int32(1))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if l2 == int32(117) {
				F_smgrcreate(m, v21, int32(3), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_log_smgrcreate(m, l1, int32(3))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return
					} else {
						F_smgrclose(m, v21)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							m.G0 = v8 + int32(16)
							return
						}
					}
				}
			} else {
				F_smgrclose(m, v21)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_heapam_relation_toast_am(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+84))
	return v3
}
