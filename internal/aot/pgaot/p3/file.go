package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddFileToBackupManifest(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	v9 = m.G0
	v11 = v9 - int32(1184)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+1176)) = l4
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L8
	} else {
		goto L75
	}
L2:
	;
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	m.G0 = v11 + int32(1184)
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(_a_F_AddFileToBackupManifest_0)
	v20 = v11 + int32(144)
	v25 = F_pg_snprintf(m, v20, int32(1024), int32(_a_F_AddFileToBackupManifest_1), v11+int32(48))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v28 = l2
	goto L7
L7:
	;
	v30 = v11 + int32(128)
	F_initStringInfo(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L10
	}
L8:
	;
	return
L9:
	;
	v28 = v20
	goto L7
L10:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v33 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v46 = F_strlen(m, v28)
	mBase = m.M
	v48 = v11 + int32(128)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v49 != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	F_appendStringInfoChar(m, v30, int32(10))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_appendStringInfoString(m, v11+int32(128), int32(_a_F_AddFileToBackupManifest_13))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L16
	}
L15:
	;
	v39 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)) = uint8(v39)
	goto L11
L16:
	;
	goto L11
L17:
	;
	F_appendStringInfoString(m, v48, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L33
	}
L18:
	;
	v63 = v11 + int32(128)
	F_appendStringInfoString(m, v63, int32(_a_F_AddFileToBackupManifest_2))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L24
	}
L19:
	;
	v52 = F_pg_verify_mbstr(m, int32(6), v28, v46, int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	if v52 == int32(0) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_appendStringInfoString(m, v48, int32(_a_F_AddFileToBackupManifest_14))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_escape_json_with_len(m, v48, v28, v46)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v104 = int32(_a_F_AddFileToBackupManifest_15)
	goto L17
L24:
	;
	F_enlargeStringInfo(m, v63, v46<<(uint(int32(1))%32))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	if v46 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v97 + base.I32_wrap_i64(base.I64_extend_i32_u(v46)<<(uint(int64(1))%64))
	v104 = int32(_a_F_AddFileToBackupManifest_3)
	goto L17
L27:
	;
	v76 = v28
	v78 = v71 + v72
	goto L30
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v81 = int32(1)
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80<<(uint(v81)%32))+uint32(_c_F_AddFileToBackupManifest[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v78))) = uint16(v83)
	v88 = v76 + v81
	if base.Ui32(v88) < base.Ui32(v28+v46) {
		v76 = v88
		v78 = v78 + int32(2)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	goto L29
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l3
	v109 = v11 + int32(128)
	F_appendStringInfo(m, v109, int32(_a_F_AddFileToBackupManifest_4), v11+int32(32))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	F_appendStringInfoString(m, v109, int32(_a_F_AddFileToBackupManifest_5))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	F_enlargeStringInfo(m, v109, int32(128))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	v128 = F_pg_gmtime(m, v11+int32(1176))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L8
	} else {
		goto L37
	}
L37:
	;
	v130 = F_pg_strftime(m, v121+v122, int32(128), int32(_a_F_AddFileToBackupManifest_6), v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v130 + v132
	F_appendStringInfoChar(m, v109, int32(34))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v138 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v140 = v11 - int32(-64)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	switch v142 - int32(1) {
	case 0:
		goto L49
	case 1:
		goto L48
	case 2:
		goto L47
	case 3:
		goto L46
	case 4:
		goto L45
	default:
		v191 = int32(0)
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	F_appendStringInfoString(m, v11+int32(128), int32(_a_F_AddFileToBackupManifest_12))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L72
	}
L43:
	;
	if v193 < int32(0) {
		goto L1
	} else {
		goto L58
	}
L44:
	;
	v193 = v191
	goto L43
L45:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v184 = F_pg_cryptohash_final(m, v182, v140, int32(64))
	mBase = m.M
	if v184 < int32(0) {
		v191 = int32(-1)
		goto L44
	} else {
		goto L56
	}
L46:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v174 = F_pg_cryptohash_final(m, v172, v140, int32(48))
	mBase = m.M
	if v174 < int32(0) {
		v191 = int32(-1)
		goto L44
	} else {
		goto L54
	}
L47:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v164 = F_pg_cryptohash_final(m, v162, v140, int32(32))
	mBase = m.M
	if v164 < int32(0) {
		v191 = int32(-1)
		goto L44
	} else {
		goto L52
	}
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v154 = F_pg_cryptohash_final(m, v152, v140, int32(28))
	mBase = m.M
	if v154 < int32(0) {
		v191 = int32(-1)
		goto L44
	} else {
		goto L50
	}
L49:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v147 = v145 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v147
	v193 = int32(4)
	goto L43
L50:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	v193 = int32(28)
	goto L43
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v193 = int32(32)
	goto L43
L54:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L55
	}
L55:
	;
	v193 = int32(48)
	goto L43
L56:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	F_pg_cryptohash_free(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	v191 = int32(64)
	goto L44
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if base.Ui32(v196) <= base.Ui32(int32(5)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v196<<(uint(int32(2))%32))+uint32(_c_F_AddFileToBackupManifest[1])))
	v203 = v201
	goto L61
L60:
	;
	v203 = int32(_a_F_AddFileToBackupManifest_10)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v203
	v206 = v11 + int32(128)
	F_appendStringInfo(m, v206, int32(_a_F_AddFileToBackupManifest_11), v11+int32(16))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L62
	}
L62:
	;
	F_enlargeStringInfo(m, v206, v193<<(uint(int32(1))%32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L63
	}
L63:
	;
	v217 = v11 - int32(-64)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	if v193 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v11)+132))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+132)) = v244 + base.I32_wrap_i64(base.I64_extend_i32_u(v193)<<(uint(int64(1))%64))
	F_appendStringInfoChar(m, v206, int32(34))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L71
	}
L65:
	;
	v223 = v217
	v225 = v218 + v219
	goto L68
L66:
	;
	goto L67
L67:
	;
	goto L64
L68:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v228 = int32(1)
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v227<<(uint(v228)%32))+uint32(_c_F_AddFileToBackupManifest[0]))))
	*(*uint16)(unsafe.Add(mBase, uint32(v225))) = uint16(v230)
	v235 = v223 + v228
	if base.Ui32(v235) < base.Ui32(v217+v193) {
		v223 = v235
		v225 = v225 + int32(2)
		goto L68
	} else {
		goto L70
	}
L69:
	;
	goto L67
L70:
	;
	goto L69
L71:
	;
	goto L42
L72:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_AppendStringToManifest(m, l0, v260)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v11)+128))
	F_pfree(m, v263)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	goto L4
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v28
	F_errmsg_internal(m, int32(_a_F_AddFileToBackupManifest_7), v11)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_AddFileToBackupManifest_8), int32(187), int32(_a_F_AddFileToBackupManifest_9))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateLockFile(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
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
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v430 int32
	_ = v430
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v466 int32
	_ = v466
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v481 int64
	_ = v481
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v643 int32
	_ = v643
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v749 int32
	_ = v749
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	v6 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(2608)
	m.G0 = v14
	v16 = int32(_a_F_CreateLockFile_0)
	v22 = F___strchrnul(m, v16, int32(61))
	mBase = m.M
	if v16 == v22 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v64 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v25 = v22 - v16
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+uint32(_c_F_CreateLockFile[0]))))
	if v27 != 0 {
		v58 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v64 = v58
	goto L1
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[1]))
	if v29 == int32(0) {
		v58 = v6
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32 == int32(0) {
		v58 = v6
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v36 = v29
	v37 = v32
	goto L9
L9:
	;
	v40 = F_strncmp(m, v16, v37, v25)
	mBase = m.M
	if v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v58 = v44 + int32(1)
	goto L5
L11:
	;
	goto L10
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v44 = v43 + v25
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45 == int32(61) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v49 != 0 {
		v36 = v36 + int32(4)
		v37 = v49
		goto L9
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v58 = v6
	goto L5
L17:
	;
	v68 = v64
	goto L21
L18:
	;
	v113 = v6
	goto L19
L19:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+288)) = v115
	v120 = F_open(m, l0, int32(194), v14+int32(288))
	mBase = m.M
	if v120 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v113 = v112
	goto L19
L21:
	;
	v73 = v68 + int32(1)
	v74 = int32(*(*int8)(unsafe.Add(mBase, uint32(v68))))
	v75 = F___isspace(m, v74)
	mBase = m.M
	if v75 != 0 {
		v68 = v73
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v76 = int32(1)
	switch v74&int32(255) - int32(43) {
	case 0:
		v82 = v76
		goto L25
	default:
		v84 = v74
		v85 = v68
		v86 = v76
		goto L24
	case 2:
		goto L26
	}
L23:
	;
	goto L22
L24:
	;
	v87 = int32(0)
	v89 = v84 - int32(48)
	if base.Ui32(v89) <= base.Ui32(int32(9)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73))))
	v84 = v83
	v85 = v73
	v86 = v82
	goto L24
L26:
	;
	v82 = int32(0)
	goto L25
L27:
	;
	v92 = v87
	v93 = v89
	v94 = v85
	goto L30
L28:
	;
	v106 = v87
	goto L29
L29:
	;
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v96 = int32(10)
	v98 = v92*v96 - v93
	v99 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94)+1)))
	v103 = v99 - int32(48)
	if base.Ui32(v103) < base.Ui32(v96) {
		v92 = v98
		v93 = v103
		v94 = v94 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v106 = v98
	goto L29
L32:
	;
	goto L31
L33:
	;
	v112 = int32(0) - v106
	goto L35
L34:
	;
	v112 = v106
	goto L35
L35:
	;
	goto L20
L36:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L62
	} else {
		goto L224
	}
L37:
	;
	v724 = int32(_a_F_CreateLockFile_1)
	v725 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	v726 = F_unlink(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = v725
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L62
	} else {
		goto L220
	}
L38:
	;
	v701 = int32(_a_F_CreateLockFile_1)
	v702 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	v703 = F_close(m, v466)
	mBase = m.M
	v704 = F_unlink(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = v702
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L62
	} else {
		goto L216
	}
L39:
	;
	v676 = int32(_a_F_CreateLockFile_1)
	v677 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	v678 = F_close(m, v466)
	mBase = m.M
	v679 = F_unlink(m, l0)
	mBase = m.M
	if v677 != 0 {
		goto L209
	} else {
		goto L210
	}
L40:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L62
	} else {
		goto L204
	}
L41:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L62
	} else {
		goto L199
	}
L42:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L62
	} else {
		goto L196
	}
L43:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L62
	} else {
		goto L191
	}
L44:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L62
	} else {
		goto L187
	}
L45:
	;
	v131 = v6
	goto L48
L46:
	;
	v466 = v120
	goto L47
L47:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v14-int32(-64)))) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v14)+68)) = l2
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+52)) = v478
	v481 = *(*int64)(unsafe.Add(mBase, _c_F_CreateLockFile[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+56)) = v481
	if l1 != 0 {
		goto L159
	} else {
		goto L160
	}
L48:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	if v135 != int32(20) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v466 = v457
	goto L47
L50:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+272)) = v145
	v147 = int32(0)
	v150 = F_open(m, l0, v147, v14+int32(272))
	mBase = m.M
	if v150 < v147 {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	if v135 != int32(2) {
		goto L36
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(101)) <= base.Ui32(v131) {
		goto L36
	} else {
		goto L56
	}
L54:
	;
	if base.Ui32(v131) <= base.Ui32(int32(100)) {
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L36
L56:
	;
	goto L50
L57:
	;
	v450 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+96)) = v450
	v457 = F_open(m, l0, int32(194), v14+int32(96))
	mBase = m.M
	if v457 < int32(0) {
		v131 = v131 + int32(1)
		goto L48
	} else {
		goto L158
	}
L58:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	if v154 == int32(44) {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = int32(167772190)
	v179 = v14 + int32(304)
	v181 = F_read(m, v150, v179, int32(2303))
	mBase = m.M
	if v181 < int32(0) {
		goto L44
	} else {
		goto L67
	}
L61:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	return
L63:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+112)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_2), v14+int32(112))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1317), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L62
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
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	v186 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v186
	v188 = F_close(m, v150)
	mBase = m.M
	if v181 == v186 {
		goto L43
	} else {
		goto L68
	}
L68:
	;
	v192 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v181+v179))) = uint8(v192)
	v197 = v179
	goto L70
L69:
	;
	v243 = v241 >> (uint(int32(31)) % 32)
	v245 = v241 ^ v243 - v243
	if v245 <= int32(0) {
		goto L42
	} else {
		goto L85
	}
L70:
	;
	v202 = v197 + int32(1)
	v203 = int32(*(*int8)(unsafe.Add(mBase, uint32(v197))))
	v204 = F___isspace(m, v203)
	mBase = m.M
	if v204 != 0 {
		v197 = v202
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v205 = int32(1)
	switch v203&int32(255) - int32(43) {
	case 0:
		v211 = v205
		goto L74
	default:
		v213 = v203
		v214 = v197
		v215 = v205
		goto L73
	case 2:
		goto L75
	}
L72:
	;
	goto L71
L73:
	;
	v216 = int32(0)
	v218 = v213 - int32(48)
	if base.Ui32(v218) <= base.Ui32(int32(9)) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v212 = int32(*(*int8)(unsafe.Add(mBase, uint32(v202))))
	v213 = v212
	v214 = v202
	v215 = v211
	goto L73
L75:
	;
	v211 = int32(0)
	goto L74
L76:
	;
	v221 = v216
	v222 = v218
	v223 = v214
	goto L79
L77:
	;
	v235 = v216
	goto L78
L78:
	;
	if v215 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v225 = int32(10)
	v227 = v221*v225 - v222
	v228 = int32(*(*int8)(unsafe.Add(mBase, uint32(v223)+1)))
	v232 = v228 - int32(48)
	if base.Ui32(v232) < base.Ui32(v225) {
		v221 = v227
		v222 = v232
		v223 = v223 + int32(1)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v235 = v227
	goto L78
L81:
	;
	goto L80
L82:
	;
	v241 = int32(0) - v235
	goto L84
L83:
	;
	v241 = v235
	goto L84
L84:
	;
	goto L69
L85:
	;
	if base.B2i32(v245 == int32(42))|base.B2i32(v245 == int32(1))|base.B2i32(v245 == v113) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if l3 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L87:
	;
	v256 = F_kill(m, v245, int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L62
	} else {
		goto L89
	}
L88:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L62
	} else {
		goto L91
	}
L89:
	;
	if v256 == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	switch v261 - int32(63) {
	case 0, 8:
		goto L86
	default:
		goto L88
	}
L91:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L62
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+256)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_5), v14+int32(256))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L62
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+244)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v14)+240)) = v245
	if l3 != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v281 = int32(_a_F_CreateLockFile_6)
	goto L96
L95:
	;
	v281 = int32(_a_F_CreateLockFile_7)
	goto L96
L96:
	;
	if l3 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v284 = int32(_a_F_CreateLockFile_8)
	goto L99
L98:
	;
	v284 = int32(_a_F_CreateLockFile_9)
	goto L99
L99:
	;
	if v241 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v287 = v281
	goto L102
L101:
	;
	v287 = v284
	goto L102
L102:
	;
	F_errhint(m, v287, v14+int32(240))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L62
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1389), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L62
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	v443 = F_unlink(m, l0)
	mBase = m.M
	if v443 < int32(0) {
		goto L40
	} else {
		goto L157
	}
L106:
	;
	v301 = int32(10)
	v302 = F___strchrnul(m, v14+int32(304), v301)
	mBase = m.M
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302))))
	if v304 == v301 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v308 == int32(0) {
		goto L105
	} else {
		goto L111
	}
L108:
	;
	v308 = v302
	goto L110
L109:
	;
	v308 = int32(0)
	goto L110
L110:
	;
	goto L107
L111:
	;
	v313 = int32(10)
	v314 = F___strchrnul(m, v308+int32(1), v313)
	mBase = m.M
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314))))
	if v316 == v313 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	if v320 == int32(0) {
		goto L105
	} else {
		goto L116
	}
L113:
	;
	v320 = v314
	goto L115
L114:
	;
	v320 = int32(0)
	goto L115
L115:
	;
	goto L112
L116:
	;
	v325 = int32(10)
	v326 = F___strchrnul(m, v320+int32(1), v325)
	mBase = m.M
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v326))))
	if v328 == v325 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v332 == int32(0) {
		goto L105
	} else {
		goto L121
	}
L118:
	;
	v332 = v326
	goto L120
L119:
	;
	v332 = int32(0)
	goto L120
L120:
	;
	goto L117
L121:
	;
	v337 = int32(10)
	v338 = F___strchrnul(m, v332+int32(1), v337)
	mBase = m.M
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338))))
	if v340 == v337 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	if v344 == int32(0) {
		goto L105
	} else {
		goto L126
	}
L123:
	;
	v344 = v338
	goto L125
L124:
	;
	v344 = int32(0)
	goto L125
L125:
	;
	goto L122
L126:
	;
	v349 = int32(10)
	v350 = F___strchrnul(m, v344+int32(1), v349)
	mBase = m.M
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if v352 == v349 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	if v356 == int32(0) {
		goto L105
	} else {
		goto L131
	}
L128:
	;
	v356 = v350
	goto L130
L129:
	;
	v356 = int32(0)
	goto L130
L130:
	;
	goto L127
L131:
	;
	v361 = int32(10)
	v362 = F___strchrnul(m, v356+int32(1), v361)
	mBase = m.M
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	if v364 == v361 {
		goto L133
	} else {
		goto L134
	}
L132:
	;
	if v368 == int32(0) {
		goto L105
	} else {
		goto L136
	}
L133:
	;
	v368 = v362
	goto L135
L134:
	;
	v368 = int32(0)
	goto L135
L135:
	;
	goto L132
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+228)) = v14 + int32(296)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+224)) = v14 + int32(300)
	v382 = F_sscanf(m, v368+int32(1), int32(_a_F_CreateLockFile_10), v14+int32(224))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L62
	} else {
		goto L137
	}
L137:
	;
	if v382 != int32(2) {
		goto L105
	} else {
		goto L138
	}
L138:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v14)+296))
	v388 = m.G0
	v390 = v388 - int32(16)
	m.G0 = v390
	v394 = F_PGSharedMemoryAttach(m, v387, v390+int32(12))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L62
	} else {
		goto L139
	}
L139:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	if v396 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	m.G0 = v390 + int32(16)
	if base.Ui32(v394) < base.Ui32(int32(2)) {
		goto L41
	} else {
		goto L156
	}
L141:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[8]))
	if v401 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L142:
	;
	if int32(0) <= v418 {
		goto L140
	} else {
		goto L151
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = int32(28)
	v418 = int32(-1)
	goto L142
L144:
	;
	v405 = v401
	goto L145
L145:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	if v396 != v406 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v418 = int32(0)
	goto L142
L147:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v405)+20))
	if v408 != 0 {
		v405 = v408
		goto L145
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	goto L146
L150:
	;
	goto L143
L151:
	;
	v423 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L62
	} else {
		goto L152
	}
L152:
	;
	if v423 == int32(0) {
		goto L140
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v390))) = v396
	F_errmsg_internal(m, int32(_a_F_CreateLockFile_11), v390)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L62
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_12), int32(324), int32(_a_F_CreateLockFile_13))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L62
	} else {
		goto L155
	}
L155:
	;
	goto L140
L156:
	;
	goto L105
L157:
	;
	goto L57
L158:
	;
	goto L49
L159:
	;
	v485 = int32(42)
	goto L161
L160:
	;
	v485 = int32(-42)
	goto L161
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = v485
	v488 = v14 + int32(304)
	v493 = F_pg_snprintf(m, v488, int32(2304), int32(_a_F_CreateLockFile_14), v14+int32(48))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L62
	} else {
		goto L162
	}
L162:
	;
	v495 = int32(0)
	if l1|base.B2i32(l3 == v495) == v495 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v500 = int32(_a_F_CreateLockFile_15)
	v501 = int32(2304)
	v503 = F_pg_ascii_verifystr(m, v488, v501)
	mBase = m.M
	if v503 == v501 {
		goto L168
	} else {
		goto L169
	}
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = int32(0)
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v514))) = int32(167772192)
	v518 = v14 + int32(304)
	v519 = F_strlen(m, v518)
	mBase = m.M
	v520 = F_write(m, v466, v518, v519)
	mBase = m.M
	v521 = F_strlen(m, v518)
	mBase = m.M
	if v520 != v521 {
		goto L39
	} else {
		goto L171
	}
L166:
	;
	goto L165
L167:
	;
	goto L166
L168:
	;
	v505 = F_strlen(m, v500)
	mBase = m.M
	goto L167
L169:
	;
	goto L170
L170:
	;
	v508 = F_strlcpy(m, v488+v503, v500, v501-v503)
	mBase = m.M
	goto L167
L171:
	;
	v523 = int32(_a_F_CreateLockFile_16)
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	v525 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v524))) = v525
	v528 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v528))) = int32(167772191)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateLockFile[9])))
	if v533 != int32(1) {
		v547 = v525
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v547 != 0 {
		goto L38
	} else {
		goto L179
	}
L173:
	;
	goto L172
L174:
	;
	goto L175
L175:
	;
	v538 = F_fsync(m, v466)
	mBase = m.M
	if v538 != int32(-1) {
		v547 = v538
		goto L173
	} else {
		goto L177
	}
L176:
	;
	v547 = int32(-1)
	goto L173
L177:
	;
	v542 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	if v542 == int32(27) {
		goto L175
	} else {
		goto L178
	}
L178:
	;
	goto L176
L179:
	;
	v549 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v549))) = int32(0)
	v552 = F_close(m, v466)
	mBase = m.M
	if v552 != 0 {
		goto L37
	} else {
		goto L180
	}
L180:
	;
	v554 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[10]))
	if v554 == int32(0) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	F_on_proc_exit(m, int32(1623))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L62
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v561 = F_pstrdup(m, l0)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L62
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	v564 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[10]))
	v565 = F_lcons(m, v561, v564)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L62
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[10])) = v565
	m.G0 = v14 + int32(2608)
	return
L187:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L62
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+128)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_17), v14+int32(128))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L62
	} else {
		goto L189
	}
L189:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1324), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L62
	} else {
		goto L190
	}
L190:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L191:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L62
	} else {
		goto L192
	}
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+144)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_18), v14+int32(144))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L62
	} else {
		goto L193
	}
L193:
	;
	F_errhint(m, int32(_a_F_CreateLockFile_19), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L62
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1333), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L62
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v14)+164)) = v14 + int32(304)
	F_errmsg_internal(m, int32(_a_F_CreateLockFile_20), v14+int32(160))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L62
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1344), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L62
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L62
	} else {
		goto L200
	}
L200:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v14)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+208)) = v635
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v14)+296))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+212)) = v637
	F_errmsg(m, int32(_a_F_CreateLockFile_21), v14+int32(208))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L62
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+192)) = l4
	F_errhint(m, int32(_a_F_CreateLockFile_22), v14+int32(192))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L62
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1427), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L62
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L62
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+176)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_23), v14+int32(176))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L62
	} else {
		goto L206
	}
L206:
	;
	F_errhint(m, int32(_a_F_CreateLockFile_24), int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L62
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1443), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L62
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	v682 = v677
	goto L211
L210:
	;
	v682 = int32(51)
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = v682
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L62
	} else {
		goto L212
	}
L212:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L62
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_25), v14+int32(32))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L62
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1478), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L62
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L62
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_25), v14+int32(16))
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L62
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1492), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L62
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L62
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_25), v14)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L62
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1503), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L62
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L62
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_26), v14+int32(80))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L62
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1303), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L62
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FileSync(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = F_FileAccess(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v4 < int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(-1)
L4:
	;
	goto L5
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_FileSync[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_FileSync[1])))
	if v16 != int32(1) {
		v38 = int32(0)
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_FileSync[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v40))) = int32(0)
	return v38
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_FileSync[2]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20+l0*int32(48))))
	goto L8
L8:
	;
	v28 = F_fsync(m, v24)
	mBase = m.M
	if v28 != int32(-1) {
		v38 = v28
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(-1)
	goto L6
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_FileSync[3]))
	if v32 == int32(27) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
}
func F_load_file(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v81 = F_expand_dynamic_library_name(m, l0)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L21
	} else {
		goto L26
	}
L2:
	;
	v10 = int32(_a_F_load_file_0)
	goto L5
L3:
	;
	if v48-v49 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	goto L6
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v18 = l0
	v19 = v10
	v20 = int32(16)
	v21 = v17
	goto L11
L8:
	;
	v44 = v10
	v48 = int32(0)
	goto L9
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	goto L3
L10:
	;
	v44 = v39
	v48 = v41
	goto L9
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v21 != v23)|base.B2i32(v23 == int32(0)) != 0 {
		v39 = v19
		v41 = v21
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v39 = v33
	v41 = int32(0)
	goto L10
L13:
	;
	v29 = v20 - int32(1)
	if v29 == int32(0) {
		v39 = v19
		v41 = v21
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v32 = int32(1)
	v33 = v19 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v34 != 0 {
		v18 = v18 + v32
		v19 = v33
		v20 = v29
		v21 = v34
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	v62 = Fn13878(m, l0+int32(16), int32(47))
	mBase = m.M
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	if v62 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	return
L22:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(_a_F_load_file_1), v6)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_load_file_2), int32(528), int32(_a_F_load_file_3))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L26:
	;
	v83 = F_internal_load_library(m, v81)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	F_pfree(m, v81)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	m.G0 = v6 + int32(16)
	return
}
func F_sendFileWithContent(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v11 = m.G0
	v13 = v11 - int32(128)
	m.G0 = v13
	v16 = v13 + int32(24)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v18 = F_pg_checksum_init(m, v16, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L2
	} else {
		goto L40
	}
L2:
	;
	return
L3:
	;
	if int32(0) <= v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = F_strlen(m, l2)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = int32(123)
	v25 = m.Env.X__syscall_getegid32(m)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v25
	v27 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = base.I64_extend_i32_s(v22)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_sendFileWithContent[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v32
	v34 = int32(0)
	F__tarWriteHeader(m, l0, l1, v34, v13+int32(32), v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L2
	} else {
		goto L37
	}
L7:
	;
	v40 = F_pg_checksum_update(m, v16, l2, v22)
	mBase = m.M
	if v40 < int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if int32(0) < v22 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v47 = l2
	v52 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v82 = (v22+int32(511))&int32(-512) - v22
	if int32(0) < v82 {
		goto L22
	} else {
		goto L23
	}
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v56 = v22 - v52
	if base.Ui32(v55) < base.Ui32(v56) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v58 = v55
	goto L16
L15:
	;
	v58 = v56
	goto L16
L16:
	;
	if v58 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	base.MemoryCopy(m, v59, v47, v58)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	m.T0[v62].(func(*base.Module, int32, int32))(m, l0, v58)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v66 = v58 + v52
	if v66 < v22 {
		v47 = v47 + v58
		v52 = v66
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v85&int32(3) != 0 {
		v106 = v82
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	F_AddFileToBackupManifest(m, l3, int32(0), l1, v22, v27, v13+int32(24))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L2
	} else {
		goto L36
	}
L25:
	;
	if v106 != 0 {
		goto L32
	} else {
		goto L33
	}
L26:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v82) {
		v106 = v82
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if v82&int32(3) != 0 {
		v106 = v82
		goto L25
	} else {
		goto L28
	}
L28:
	;
	v94 = v82 + v85
	v96 = v85 + int32(4)
	if base.Ui32(v96) < base.Ui32(v94) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v98 = v94
	goto L31
L30:
	;
	v98 = v96
	goto L31
L31:
	;
	v106 = (v85^int32(-1)+v98)&int32(-4) + int32(4)
	goto L25
L32:
	;
	base.MemoryFill(m, v85, int32(0), v106)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	m.T0[v110].(func(*base.Module, int32, int32))(m, l0, v82)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	goto L24
L36:
	;
	m.G0 = v13 + int32(128)
	return
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l1
	F_errmsg_internal(m, int32(_a_F_sendFileWithContent_0), v13)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_sendFileWithContent_1), int32(1084), int32(_a_F_sendFileWithContent_2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_sendFileWithContent_3), v13+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_sendFileWithContent_1), int32(1109), int32(_a_F_sendFileWithContent_2))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
