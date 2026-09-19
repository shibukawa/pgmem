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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int64
	_ = v443
	var v444 int64
	_ = v444
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v497 int32
	_ = v497
	var v506 int32
	_ = v506
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v522 int64
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
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
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v623 int32
	_ = v623
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v684 int32
	_ = v684
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v796 int32
	_ = v796
	var v801 int32
	_ = v801
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(2608)
	m.G0 = v15
	v17 = m.Env.Pgmem_getpid(m)
	mBase = m.M
	v18 = int32(_a_F_CreateLockFile_0)
	v24 = F___strchrnul(m, v18, int32(61))
	mBase = m.M
	if v18 == v24 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v66 != 0 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = v24 - v18
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_CreateLockFile[0]))))
	if v29 != 0 {
		v60 = v6
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v66 = v60
	goto L1
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[1]))
	if v31 == int32(0) {
		v60 = v6
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	if v34 == int32(0) {
		v60 = v6
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v38 = v31
	v39 = v34
	goto L9
L9:
	;
	v42 = F_strncmp(m, v18, v39, v27)
	mBase = m.M
	if v42 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v60 = v46 + int32(1)
	goto L5
L11:
	;
	goto L10
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v46 = v45 + v27
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v47 == int32(61) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v51 != 0 {
		v38 = v38 + int32(4)
		v39 = v51
		goto L9
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v60 = v6
	goto L5
L17:
	;
	v70 = v66
	goto L21
L18:
	;
	v115 = v6
	goto L19
L19:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+288)) = v117
	v122 = F_open(m, l0, int32(194), v15+int32(288))
	mBase = m.M
	if v122 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L20:
	;
	v115 = v114
	goto L19
L21:
	;
	v75 = v70 + int32(1)
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v70))))
	v77 = F___isspace(m, v76)
	mBase = m.M
	if v77 != 0 {
		v70 = v75
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v78 = int32(1)
	switch v76&int32(255) - int32(43) {
	case 0:
		v84 = v78
		goto L25
	default:
		v86 = v76
		v87 = v70
		v88 = v78
		goto L24
	case 2:
		goto L26
	}
L23:
	;
	goto L22
L24:
	;
	v89 = int32(0)
	v91 = v86 - int32(48)
	if base.Ui32(v91) <= base.Ui32(int32(9)) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v85 = int32(*(*int8)(unsafe.Add(mBase, uint32(v75))))
	v86 = v85
	v87 = v75
	v88 = v84
	goto L24
L26:
	;
	v84 = int32(0)
	goto L25
L27:
	;
	v94 = v89
	v95 = v91
	v96 = v87
	goto L30
L28:
	;
	v108 = v89
	goto L29
L29:
	;
	if v88 != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v98 = int32(10)
	v100 = v94*v98 - v95
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v96)+1)))
	v105 = v101 - int32(48)
	if base.Ui32(v105) < base.Ui32(v98) {
		v94 = v100
		v95 = v105
		v96 = v96 + int32(1)
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v108 = v100
	goto L29
L32:
	;
	goto L31
L33:
	;
	v114 = int32(0) - v108
	goto L35
L34:
	;
	v114 = v108
	goto L35
L35:
	;
	goto L20
L36:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L62
	} else {
		goto L232
	}
L37:
	;
	v765 = int32(_a_F_CreateLockFile_1)
	v766 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	v767 = F_unlink(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = v766
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L62
	} else {
		goto L228
	}
L38:
	;
	v742 = int32(_a_F_CreateLockFile_1)
	v743 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	v744 = F_close(m, v506)
	mBase = m.M
	v745 = F_unlink(m, l0)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = v743
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L62
	} else {
		goto L224
	}
L39:
	;
	v717 = int32(_a_F_CreateLockFile_1)
	v718 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	v719 = F_close(m, v506)
	mBase = m.M
	v720 = F_unlink(m, l0)
	mBase = m.M
	if v718 != 0 {
		goto L217
	} else {
		goto L218
	}
L40:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v699 = m.ExcPending
	if v699 != 0 {
		goto L62
	} else {
		goto L212
	}
L41:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L62
	} else {
		goto L207
	}
L42:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L62
	} else {
		goto L204
	}
L43:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L62
	} else {
		goto L199
	}
L44:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L62
	} else {
		goto L195
	}
L45:
	;
	v133 = v6
	goto L48
L46:
	;
	v506 = v122
	goto L47
L47:
	;
	v515 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v515
	*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = l2
	v519 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v519
	v522 = *(*int64)(unsafe.Add(mBase, _c_F_CreateLockFile[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v522
	if l1 != 0 {
		goto L167
	} else {
		goto L168
	}
L48:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	if v138 != int32(20) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v506 = v497
	goto L47
L50:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+272)) = v148
	v150 = int32(0)
	v153 = F_open(m, l0, v150, v15+int32(272))
	mBase = m.M
	if v153 < v150 {
		goto L58
	} else {
		goto L59
	}
L51:
	;
	if v138 != int32(2) {
		goto L36
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if base.Ui32(int32(101)) <= base.Ui32(v133) {
		goto L36
	} else {
		goto L56
	}
L54:
	;
	if base.Ui32(v133) <= base.Ui32(int32(100)) {
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
	v490 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v490
	v497 = F_open(m, l0, int32(194), v15+int32(96))
	mBase = m.M
	if v497 < int32(0) {
		v133 = v133 + int32(1)
		goto L48
	} else {
		goto L166
	}
L58:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	if v157 == int32(44) {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = int32(167772190)
	v182 = v15 + int32(304)
	v184 = F_read(m, v153, v182, int32(2303))
	mBase = m.M
	if v184 < int32(0) {
		goto L44
	} else {
		goto L67
	}
L61:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
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
	v165 = m.ExcPending
	if v165 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_2), v15+int32(112))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1300), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
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
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	v189 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = v189
	v191 = F_close(m, v153)
	mBase = m.M
	if v184 == v189 {
		goto L43
	} else {
		goto L68
	}
L68:
	;
	v195 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v184+v182))) = uint8(v195)
	v200 = v182
	goto L70
L69:
	;
	v246 = v244 >> (uint(int32(31)) % 32)
	v248 = v244 ^ v246 - v246
	if v248 <= int32(0) {
		goto L42
	} else {
		goto L85
	}
L70:
	;
	v205 = v200 + int32(1)
	v206 = int32(*(*int8)(unsafe.Add(mBase, uint32(v200))))
	v207 = F___isspace(m, v206)
	mBase = m.M
	if v207 != 0 {
		v200 = v205
		goto L70
	} else {
		goto L72
	}
L71:
	;
	v208 = int32(1)
	switch v206&int32(255) - int32(43) {
	case 0:
		v214 = v208
		goto L74
	default:
		v216 = v206
		v217 = v200
		v218 = v208
		goto L73
	case 2:
		goto L75
	}
L72:
	;
	goto L71
L73:
	;
	v219 = int32(0)
	v221 = v216 - int32(48)
	if base.Ui32(v221) <= base.Ui32(int32(9)) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v215 = int32(*(*int8)(unsafe.Add(mBase, uint32(v205))))
	v216 = v215
	v217 = v205
	v218 = v214
	goto L73
L75:
	;
	v214 = int32(0)
	goto L74
L76:
	;
	v224 = v219
	v225 = v221
	v226 = v217
	goto L79
L77:
	;
	v238 = v219
	goto L78
L78:
	;
	if v218 != 0 {
		goto L82
	} else {
		goto L83
	}
L79:
	;
	v228 = int32(10)
	v230 = v224*v228 - v225
	v231 = int32(*(*int8)(unsafe.Add(mBase, uint32(v226)+1)))
	v235 = v231 - int32(48)
	if base.Ui32(v235) < base.Ui32(v228) {
		v224 = v230
		v225 = v235
		v226 = v226 + int32(1)
		goto L79
	} else {
		goto L81
	}
L80:
	;
	v238 = v230
	goto L78
L81:
	;
	goto L80
L82:
	;
	v244 = int32(0) - v238
	goto L84
L83:
	;
	v244 = v238
	goto L84
L84:
	;
	goto L69
L85:
	;
	if base.B2i32(v248 == v17)|base.B2i32(v248 == int32(1))|base.B2i32(v248 == v115) != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	if l3 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L87:
	;
	v257 = int32(0)
	v258 = F_pgmem_kill(m, v248, v257)
	mBase = m.M
	if v258 == v257 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L62
	} else {
		goto L90
	}
L89:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	switch v262 - int32(63) {
	case 0, 8:
		goto L86
	default:
		goto L88
	}
L90:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L62
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+256)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_5), v15+int32(256))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L62
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v248
	if l3 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v282 = int32(_a_F_CreateLockFile_6)
	goto L95
L94:
	;
	v282 = int32(_a_F_CreateLockFile_7)
	goto L95
L95:
	;
	if l3 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v285 = int32(_a_F_CreateLockFile_8)
	goto L98
L97:
	;
	v285 = int32(_a_F_CreateLockFile_9)
	goto L98
L98:
	;
	if v244 < int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v288 = v282
	goto L101
L100:
	;
	v288 = v285
	goto L101
L101:
	;
	F_errhint(m, v288, v15+int32(240))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L62
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1372), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L62
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L104:
	;
	v483 = F_unlink(m, l0)
	mBase = m.M
	if v483 < int32(0) {
		goto L40
	} else {
		goto L165
	}
L105:
	;
	v302 = int32(10)
	v303 = F___strchrnul(m, v15+int32(304), v302)
	mBase = m.M
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v305 == v302 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	if v309 == int32(0) {
		goto L104
	} else {
		goto L110
	}
L107:
	;
	v309 = v303
	goto L109
L108:
	;
	v309 = int32(0)
	goto L109
L109:
	;
	goto L106
L110:
	;
	v314 = int32(10)
	v315 = F___strchrnul(m, v309+int32(1), v314)
	mBase = m.M
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315))))
	if v317 == v314 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v321 == int32(0) {
		goto L104
	} else {
		goto L115
	}
L112:
	;
	v321 = v315
	goto L114
L113:
	;
	v321 = int32(0)
	goto L114
L114:
	;
	goto L111
L115:
	;
	v326 = int32(10)
	v327 = F___strchrnul(m, v321+int32(1), v326)
	mBase = m.M
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327))))
	if v329 == v326 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v333 == int32(0) {
		goto L104
	} else {
		goto L120
	}
L117:
	;
	v333 = v327
	goto L119
L118:
	;
	v333 = int32(0)
	goto L119
L119:
	;
	goto L116
L120:
	;
	v338 = int32(10)
	v339 = F___strchrnul(m, v333+int32(1), v338)
	mBase = m.M
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339))))
	if v341 == v338 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v345 == int32(0) {
		goto L104
	} else {
		goto L125
	}
L122:
	;
	v345 = v339
	goto L124
L123:
	;
	v345 = int32(0)
	goto L124
L124:
	;
	goto L121
L125:
	;
	v350 = int32(10)
	v351 = F___strchrnul(m, v345+int32(1), v350)
	mBase = m.M
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351))))
	if v353 == v350 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	if v357 == int32(0) {
		goto L104
	} else {
		goto L130
	}
L127:
	;
	v357 = v351
	goto L129
L128:
	;
	v357 = int32(0)
	goto L129
L129:
	;
	goto L126
L130:
	;
	v362 = int32(10)
	v363 = F___strchrnul(m, v357+int32(1), v362)
	mBase = m.M
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v363))))
	if v365 == v362 {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v369 == int32(0) {
		goto L104
	} else {
		goto L135
	}
L132:
	;
	v369 = v363
	goto L134
L133:
	;
	v369 = int32(0)
	goto L134
L134:
	;
	goto L131
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+228)) = v15 + int32(296)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v15 + int32(300)
	v383 = F_sscanf(m, v369+int32(1), int32(_a_F_CreateLockFile_10), v15+int32(224))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L62
	} else {
		goto L136
	}
L136:
	;
	if v383 != int32(2) {
		goto L104
	} else {
		goto L137
	}
L137:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v15)+296))
	v389 = m.G0
	v391 = v389 - int32(16)
	m.G0 = v391
	v393 = int32(0)
	v395 = v391 + int32(12)
	v398 = m.G0
	v400 = v398 - int32(192)
	m.G0 = v400
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v393
	v404 = int32(2)
	v408 = F_pgmem_shmctl(m, v388, v404, v400+int32(104))
	mBase = m.M
	if v408 < v393 {
		goto L140
	} else {
		goto L141
	}
L138:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v391)+12))
	if v455 == int32(0) {
		goto L157
	} else {
		goto L158
	}
L139:
	;
	m.G0 = v400 + int32(192)
	goto L138
L140:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	switch v412 - int32(2) {
	case 0:
		goto L144
	default:
		goto L143
	case 22, 26:
		v451 = v404
		goto L139
	}
L141:
	;
	goto L142
L142:
	;
	v417 = int32(0)
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[5]))
	v422 = F_stat(m, v419, v400+int32(8))
	mBase = m.M
	if v422 < v417 {
		v451 = v417
		goto L139
	} else {
		goto L145
	}
L143:
	;
	v451 = int32(0)
	goto L139
L144:
	;
	v451 = int32(3)
	goto L139
L145:
	;
	v425 = F_pgmem_shmat(m, v388, v393)
	mBase = m.M
	if v425 == int32(-1) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v428 = int32(2)
	v430 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	switch v430 - v428 {
	case 0:
		goto L150
	default:
		goto L149
	case 22, 26:
		v451 = v428
		goto L139
	}
L147:
	;
	goto L148
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395))) = v425
	v436 = int32(3)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v425)))
	if v437 != int32(679834894) {
		v451 = v436
		goto L139
	} else {
		goto L151
	}
L149:
	;
	v451 = int32(0)
	goto L139
L150:
	;
	v451 = int32(3)
	goto L139
L151:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v425)+24))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v400)+8))
	if v440 != v441 {
		v451 = v436
		goto L139
	} else {
		goto L152
	}
L152:
	;
	v443 = *(*int64)(unsafe.Add(mBase, uint32(v425)+32))
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v400)+96))
	if v443 != v444 {
		v451 = v436
		goto L139
	} else {
		goto L153
	}
L153:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v400)+176))
	if v448 != 0 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v449 = int32(1)
	goto L156
L155:
	;
	v449 = int32(4)
	goto L156
L156:
	;
	v451 = v449
	goto L139
L157:
	;
	m.G0 = v391 + int32(16)
	if base.Ui32(v451) < base.Ui32(int32(2)) {
		goto L41
	} else {
		goto L164
	}
L158:
	;
	v458 = F_pgmem_shmdt(m, v455)
	mBase = m.M
	if int32(0) <= v458 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v463 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L62
	} else {
		goto L160
	}
L160:
	;
	if v463 == int32(0) {
		goto L157
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391))) = v455
	F_errmsg_internal(m, int32(_a_F_CreateLockFile_11), v391)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L62
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_12), int32(324), int32(_a_F_CreateLockFile_13))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L62
	} else {
		goto L163
	}
L163:
	;
	goto L157
L164:
	;
	goto L104
L165:
	;
	goto L57
L166:
	;
	goto L49
L167:
	;
	v526 = v17
	goto L169
L168:
	;
	v526 = int32(0) - v17
	goto L169
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v526
	v529 = v15 + int32(304)
	v534 = F_pg_snprintf(m, v529, int32(2304), int32(_a_F_CreateLockFile_14), v15+int32(48))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L62
	} else {
		goto L170
	}
L170:
	;
	v536 = int32(0)
	if l1|base.B2i32(l3 == v536) == v536 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v541 = int32(_a_F_CreateLockFile_15)
	v542 = int32(2304)
	v544 = F_pg_ascii_verifystr(m, v529, v542)
	mBase = m.M
	if v544 == v542 {
		goto L176
	} else {
		goto L177
	}
L172:
	;
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = int32(0)
	v555 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v555))) = int32(167772192)
	v559 = v15 + int32(304)
	v560 = F_strlen(m, v559)
	mBase = m.M
	v561 = F_write(m, v506, v559, v560)
	mBase = m.M
	v562 = F_strlen(m, v559)
	mBase = m.M
	if v561 != v562 {
		goto L39
	} else {
		goto L179
	}
L174:
	;
	goto L173
L175:
	;
	goto L174
L176:
	;
	v546 = F_strlen(m, v541)
	mBase = m.M
	goto L175
L177:
	;
	goto L178
L178:
	;
	v549 = F_strlcpy(m, v529+v544, v541, v542-v544)
	mBase = m.M
	goto L175
L179:
	;
	v564 = int32(_a_F_CreateLockFile_16)
	v565 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	v566 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v565))) = v566
	v569 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v569))) = int32(167772191)
	v574 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CreateLockFile[8])))
	if v574 != int32(1) {
		v588 = v566
		goto L181
	} else {
		goto L182
	}
L180:
	;
	if v588 != 0 {
		goto L38
	} else {
		goto L187
	}
L181:
	;
	goto L180
L182:
	;
	goto L183
L183:
	;
	v579 = F_fsync(m, v506)
	mBase = m.M
	if v579 != int32(-1) {
		v588 = v579
		goto L181
	} else {
		goto L185
	}
L184:
	;
	v588 = int32(-1)
	goto L181
L185:
	;
	v583 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3]))
	if v583 == int32(27) {
		goto L183
	} else {
		goto L186
	}
L186:
	;
	goto L184
L187:
	;
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v590))) = int32(0)
	v593 = F_close(m, v506)
	mBase = m.M
	if v593 != 0 {
		goto L37
	} else {
		goto L188
	}
L188:
	;
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[9]))
	if v595 == int32(0) {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	F_on_proc_exit(m, int32(1623))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L62
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v602 = F_pstrdup(m, l0)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L62
	} else {
		goto L193
	}
L192:
	;
	goto L191
L193:
	;
	v605 = *(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[9]))
	v606 = F_lcons(m, v602, v605)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L62
	} else {
		goto L194
	}
L194:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[9])) = v606
	m.G0 = v15 + int32(2608)
	return
L195:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L62
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_17), v15+int32(128))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L62
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1307), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
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
	v635 = m.ExcPending
	if v635 != 0 {
		goto L62
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_18), v15+int32(144))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L62
	} else {
		goto L201
	}
L201:
	;
	F_errhint(m, int32(_a_F_CreateLockFile_19), int32(0))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L62
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1316), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v15)+164)) = v15 + int32(304)
	F_errmsg_internal(m, int32(_a_F_CreateLockFile_20), v15+int32(160))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L62
	} else {
		goto L205
	}
L205:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1327), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L62
	} else {
		goto L206
	}
L206:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L207:
	;
	F_errcode(m, int32(16777238))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L62
	} else {
		goto L208
	}
L208:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v15)+300))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v676
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v15)+296))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+212)) = v678
	F_errmsg(m, int32(_a_F_CreateLockFile_21), v15+int32(208))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L62
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = l4
	F_errhint(m, int32(_a_F_CreateLockFile_22), v15+int32(192))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L62
	} else {
		goto L210
	}
L210:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1410), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L62
	} else {
		goto L211
	}
L211:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L212:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L62
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_23), v15+int32(176))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L62
	} else {
		goto L214
	}
L214:
	;
	F_errhint(m, int32(_a_F_CreateLockFile_24), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L62
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1426), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L62
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	v723 = v718
	goto L219
L218:
	;
	v723 = int32(51)
	goto L219
L219:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_CreateLockFile[3])) = v723
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L62
	} else {
		goto L220
	}
L220:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v730 = m.ExcPending
	if v730 != 0 {
		goto L62
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_25), v15+int32(32))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L62
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1461), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
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
	v753 = m.ExcPending
	if v753 != 0 {
		goto L62
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_25), v15+int32(16))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L62
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1475), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v764 = m.ExcPending
	if v764 != 0 {
		goto L62
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L62
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_25), v15)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L62
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1486), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L62
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L62
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = l0
	F_errmsg(m, int32(_a_F_CreateLockFile_26), v15+int32(80))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L62
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_CreateLockFile_3), int32(1286), int32(_a_F_CreateLockFile_4))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L62
	} else {
		goto L235
	}
L235:
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
	v62 = Fn13880(m, l0+int32(16), int32(47))
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
