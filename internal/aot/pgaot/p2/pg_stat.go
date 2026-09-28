package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_stat_get_autovacuum_count(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+144))
			return v11
		}
	}
}
func F_pg_stat_get_autovacuum_scores(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v245 int64
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	v11 = m.G0
	v13 = v11 - int32(208)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v21 = F_MultiXactMemberFreezeThreshold(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = F_ReadNextFullTransactionId(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*uint32)(unsafe.Add(mBase, _c_F_pg_stat_get_autovacuum_scores[0])) = uint32(v24)
	v28 = F_ReadNextMultiXactId(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_autovacuum_scores[1])) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v13)+168)) = int64(481036337156)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_autovacuum_scores[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+196)) = v34
	v41 = F_hash_create(m, int32(_a_F_pg_stat_get_autovacuum_scores_0), int64(100), v13+int32(160), int32(1064))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v45 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v47 = int32(0)
	v49 = F_table_beginscan_catalog(m, v45, v47, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v51 = F_heap_getnext(m, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v51 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v53 = v51
	goto L13
L11:
	;
	goto L12
L12:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+188))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	m.T0[v124].(func(*base.Module, int32))(m, v49)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L27
	}
L13:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+22)))
	v65 = v63 + v64
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+119)))
	switch v66 - int32(109) {
	case 0, 5:
		goto L16
	default:
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v110 = F_heap_getnext(m, v49)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L25
	}
L16:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65)+118)))
	if v69 == int32(116) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+112))
	if v72 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v77 = F_extractRelOptions(m, v53, v75, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v77 == int32(0) {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	v82 = F_palloc(m, int32(96))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.MemoryCopy(m, v82, v77+int32(16), int32(96))
	F_pfree(m, v77)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v95 = F_hash_search(m, v41, v65+int32(112), int32(1), v13+int32(16))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v98 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v95)+8)) = uint8(v98)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = v97
	base.MemoryCopy(m, v95+int32(16), v82, int32(96))
	F_pfree(m, v82)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L15
L25:
	;
	if v110 != 0 {
		v53 = v110
		goto L13
	} else {
		goto L26
	}
L26:
	;
	goto L14
L27:
	;
	v127 = int32(0)
	v129 = F_table_beginscan_catalog(m, v45, v127, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v131 = F_heap_getnext(m, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v131 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v133 = v131
	goto L33
L31:
	;
	goto L32
L32:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+188))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v269)+12))
	m.T0[v270].(func(*base.Module, int32))(m, v129)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L58
	}
L33:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+22)))
	v145 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+8)) = uint16(v145)
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = int64(0)
	v149 = v143 + v144
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+119)))
	v152 = v150 - int32(109)
	if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v152))|base.B2i32(int32(1)<<(uint(v152)%32)&int32(161) == v145) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L32
L35:
	;
	v256 = F_heap_getnext(m, v129)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L56
	}
L36:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+118)))
	if v162 == int32(116) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
	v167 = F_extractRelOptions(m, v133, v165, int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v227 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v149))))
	v228 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v228
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v227
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v13)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v231
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v13)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v233
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v13)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+48)) = v235
	v237 = *(*int64)(unsafe.Add(mBase, uint32(v13)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+56)) = v237
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v13)+144))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v239
	v241 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v13)+159)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v241
	v243 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v13)+158)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+80)) = v243
	v245 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v13)+157)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	F_tuplestore_putvalues(m, v247, v248, v13+int32(16), v13)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L55
	}
L39:
	;
	if v167 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v171 = int32(0)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+119)))
	if v172 != int32(116) {
		v188 = v171
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	v203 = F_palloc(m, int32(96))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L51
	}
L43:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	F_relation_needs_vacanalyze(m, v190, v188, v149, v21, int32(0), v13+int32(159), v13+int32(158), v13+int32(157), v13+int32(104))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	v178 = F_hash_search(m, v41, v149, int32(0), v13+int32(16))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)))
	if v180 != int32(1) {
		v188 = v171
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+8)))
	if v186 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v187 = v178 + int32(16)
	goto L49
L48:
	;
	v187 = int32(0)
	goto L49
L49:
	;
	v188 = v187
	goto L43
L50:
	;
	goto L38
L51:
	;
	base.MemoryCopy(m, v203, v167+int32(16), int32(96))
	F_pfree(m, v167)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	F_relation_needs_vacanalyze(m, v211, v203, v149, v21, int32(0), v13+int32(159), v13+int32(158), v13+int32(157), v13+int32(104))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_pfree(m, v203)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	goto L38
L55:
	;
	goto L35
L56:
	;
	if v256 != 0 {
		v133 = v256
		goto L33
	} else {
		goto L57
	}
L57:
	;
	goto L34
L58:
	;
	F_relation_close(m, v45, int32(1))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_hash_destroy(m, v41)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	m.G0 = v13 + int32(208)
	return int64(0)
}
func F_pg_stat_get_backend_client_addr(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
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
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
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
	var v179 int32
	_ = v179
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v199 int64
	_ = v199
	v10 = m.G0
	v12 = v10 - int32(256)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pgstat_get_beentry_by_proc_number(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(256)
	return v199
L2:
	;
	return int64(0)
L3:
	;
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
	v199 = int64(0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_client_addr[0]))
	v27 = F_has_privs_of_role(m, v25, int32(3375))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v38 = v15 + int32(56)
	v42 = (int32(-56) - v15) & int32(3)
	if v42 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	if v27 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_client_addr[0]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v32 = F_has_privs_of_role(m, v30, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v32 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v34 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
	v199 = int64(0)
	goto L1
L12:
	;
	v187 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v187)
	v199 = int64(0)
	goto L1
L13:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	switch v154 - int32(2) {
	case 0, 8:
		goto L42
	default:
		goto L43
	}
L14:
	;
	v55 = v42 + v38
	v59 = (v15 + int32(188)) & int32(-4)
	v61 = v59 - int32(28)
	if base.Ui32(v55) < base.Ui32(v61) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38))))
	if v45 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v42 == int32(1) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+57)))
	if v48 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v42 == int32(2) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+58)))
	if v51|base.B2i32(v42 != int32(3)) != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	v64 = v42
	v65 = v55
	goto L24
L22:
	;
	v92 = v42
	goto L23
L23:
	;
	v100 = v92 + v38
	if base.Ui32(v100) < base.Ui32(v59) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v72|(v73|(v74|(v75|(v76|(v77|(v78|v79)))))) != 0 {
		goto L13
	} else {
		goto L26
	}
L25:
	;
	v92 = v88
	goto L23
L26:
	;
	v88 = v64 + int32(32)
	v89 = v38 + v88
	if base.Ui32(v89) < base.Ui32(v61) {
		v64 = v88
		v65 = v89
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v103 = v92
	v104 = v100
	goto L31
L29:
	;
	v117 = v92
	goto L30
L30:
	;
	v125 = int32(132)
	if base.Ui32(v117) <= base.Ui32(v125) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v111 != 0 {
		goto L13
	} else {
		goto L33
	}
L32:
	;
	v117 = v113
	goto L30
L33:
	;
	v113 = v103 + int32(4)
	v114 = v38 + v113
	if base.Ui32(v114) < base.Ui32(v59) {
		v103 = v113
		v104 = v114
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v128 = v125
	goto L37
L36:
	;
	v128 = v117
	goto L37
L37:
	;
	v130 = v117
	goto L38
L38:
	;
	if v130 == v128 {
		goto L12
	} else {
		goto L40
	}
L39:
	;
	goto L13
L40:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130+v38))))
	if v142 == int32(0) {
		v130 = v130 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v160)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
	v167 = F_pg_getnameinfo_all(m, v38, v162, v12, int32(255), v160, v160, int32(3))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	v157 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v157)
	v199 = int64(0)
	goto L1
L44:
	;
	if v167 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v169)
	v199 = int64(0)
	goto L1
L46:
	;
	goto L47
L47:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if v172 != int32(10) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v185 = F_DirectFunctionCall1Coll(m, int32(1678), int32(0), base.I64_extend_i32_u(v12))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L2
	} else {
		goto L52
	}
L49:
	;
	goto L48
L50:
	;
	v176 = F_strchr(m, v12, int32(37))
	mBase = m.M
	if v176 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v179 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v179)
	goto L49
L52:
	;
	v199 = v185
	goto L1
}
func F_pg_stat_get_backend_dbid(m *base.Module, l0 int32) int64 {
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
	var v14 int64
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pgstat_get_beentry_by_proc_number(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int64(0)
		} else {
			v14 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v4)+48)))
			return v14
		}
	}
}
func F_pg_stat_get_blocks_fetched(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+104))
			return v11
		}
	}
}
func F_pg_stat_get_blocks_hit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+112))
			return v11
		}
	}
}
func F_pg_stat_get_checkpointer_restartpoints_timed(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+24))
		return v6
	}
}
func F_pg_stat_get_checkpointer_slru_written(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+72))
		return v6
	}
}
func F_pg_stat_get_checkpointer_stat_reset_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)+80))
		return v6
	}
}
func F_pg_stat_get_db_conflict_tablespace(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+80))
			return v11
		}
	}
}
func F_pg_stat_get_db_sessions_killed(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+232))
			return v11
		}
	}
}
func F_pg_stat_get_db_tuples_returned(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+32))
			return v11
		}
	}
}
func F_pg_stat_get_io(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v2 = int32(0)
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_pgstat_snapshot_fixed(m, int32(10))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = v2
	goto L4
L4:
	;
	if int32(base.Ui32(int32(_a_F_pg_stat_get_io_0))>>(uint(v13)%32))&base.B2i32(base.Ui32(v13) < base.Ui32(int32(17))) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	return int64(0)
L6:
	;
	v24 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_io[0]))
	F_pg_stat_io_build_tuples(m, v8, v13*int32(2880)+int32(_a_F_pg_stat_get_io_1), v13, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v28 = v13 + int32(1)
	if v28 != int32(18) {
		v13 = v28
		goto L4
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	goto L5
}
func F_pg_stat_get_last_autoanalyze_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+168))
			if v9 != int64(0) {
				v16 = v9
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				v16 = int64(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v16 = int64(0)
		}
		return v16
	}
}
func F_pg_stat_get_last_vacuum_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_fetch_stat_tabentry(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 != 0 {
			v9 = *(*int64)(unsafe.Add(mBase, uint32(v5)+120))
			if v9 != int64(0) {
				v16 = v9
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				v16 = int64(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v16 = int64(0)
		}
		return v16
	}
}
func F_pg_stat_get_total_autoanalyze_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+208))
			return base.I64_reinterpret_f64(base.F64_convert_i64_s(v11))
		}
	}
}
func F_pg_stat_get_total_vacuum_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+184))
			return base.I64_reinterpret_f64(base.F64_convert_i64_s(v11))
		}
	}
}
func F_pg_stat_get_wal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	F_pgstat_snapshot_fixed(m, int32(13))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[0]))
		v15 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[1]))
		*(*int64)(unsafe.Add(mBase, uint32(v5)+40)) = v15
		v18 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[2]))
		*(*int64)(unsafe.Add(mBase, uint32(v5)+32)) = v18
		v21 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[3]))
		*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v21
		v24 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[4]))
		*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v24
		v27 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_wal[5]))
		*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v27
		v31 = F_pg_stat_wal_build_tuple(m, v5+int32(8), v13)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int64(0)
		} else {
			m.G0 = v5 + int32(48)
			return v31
		}
	}
}
func F_pg_stat_get_xact_function_self_time(m *base.Module, l0 int32) int64 {
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
	var v14 int64
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_find_funcstat_entry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int64(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+16))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v14), float64(1e+06)))
		}
	}
}
func F_pg_stat_get_xact_function_total_time(m *base.Module, l0 int32) int64 {
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
	var v14 int64
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_find_funcstat_entry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int64(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
			return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v14), float64(1e+06)))
		}
	}
}
func F_pg_stat_get_xact_tuples_fetched(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+32))
			return v11
		}
	}
}
func F_pg_stat_get_xact_tuples_returned(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+24))
			return v11
		}
	}
}
func F_pg_stat_get_xact_tuples_updated(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+48))
			return v11
		}
	}
}
func F_pg_stat_io_build_tuples(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) {
	mBase := m.M
	_ = mBase
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v91 int32
	_ = v91
	var v109 int32
	_ = v109
	var v127 int32
	_ = v127
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v166 int32
	_ = v166
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v259 int32
	_ = v259
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v350 int64
	_ = v350
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v372 int32
	_ = v372
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v445 int32
	_ = v445
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v552 int32
	_ = v552
	var v573 int32
	_ = v573
	var v630 int32
	_ = v630
	var v656 int32
	_ = v656
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v706 int64
	_ = v706
	var v708 int32
	_ = v708
	var v712 int64
	_ = v712
	var v718 int64
	_ = v718
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v730 int64
	_ = v730
	var v731 int32
	_ = v731
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v752 int32
	_ = v752
	var v819 int32
	_ = v819
	var v823 int32
	_ = v823
	v66 = m.G0
	v68 = v66 - int32(464)
	m.G0 = v68
	v91 = v68 + int32(296)
	v109 = v68 + int32(272)
	v127 = v68 + int32(271)
	if base.Ui32(l2) <= base.Ui32(int32(17)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v151 = F_cstring_to_text(m, v150)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l2<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_io_build_tuples[0])))
	v150 = v148
	goto L4
L3:
	;
	v150 = int32(_a_F_pg_stat_io_build_tuples_0)
	goto L4
L4:
	;
	goto L1
L5:
	;
	return
L6:
	;
	v166 = int32(0)
	goto L7
L7:
	;
	v223 = v166 * int32(320)
	v225 = m.G0
	v227 = v225 - int32(16)
	m.G0 = v227
	if base.Ui32(int32(3)) <= base.Ui32(v166) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	m.G0 = v68 + int32(464)
	return
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v166<<(uint(int32(2))%32))+uint32(_c_F_pg_stat_io_build_tuples[1])))
	m.G0 = v227 + int32(16)
	v259 = int32(0)
	goto L15
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227))) = v166
	F_errmsg_internal(m, int32(_a_F_pg_stat_io_build_tuples_1), v227)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_pg_stat_io_build_tuples_2), int32(265), int32(_a_F_pg_stat_io_build_tuples_3))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L15:
	;
	v318 = m.G0
	v320 = v318 - int32(16)
	m.G0 = v320
	if base.Ui32(int32(5)) <= base.Ui32(v259) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v823 = v166 + int32(1)
	if v823 != int32(3) {
		v166 = v823
		goto L7
	} else {
		goto L115
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v337 = int32(2)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v259<<(uint(v337)%32))+uint32(_c_F_pg_stat_io_build_tuples[2])))
	v340 = int32(16)
	m.G0 = v320 + v340
	v345 = int32(0)
	base.MemoryFill(m, v68+int32(304), v345, int32(160))
	*(*int32)(unsafe.Add(mBase, uint32(v68)+288)) = v345
	v350 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v68)+280)) = v350
	*(*int64)(unsafe.Add(mBase, uint32(v68)+272)) = v350
	v356 = base.B2i32(base.Ui32(v340) < base.Ui32(l2))
	v357 = int32(1)
	v358 = v357 << (uint(l2) % 32)
	v372 = base.B2i32(v166 == v357)
	if v356|base.B2i32(v358&int32(_a_F_pg_stat_io_build_tuples_4) == v345)|(base.B2i32(v166 == v337)&base.B2i32(base.Ui32(v259-int32(4)) < base.Ui32(int32(-2)))|v372&base.B2i32(v259 != int32(3))) != 0 {
		v416 = v345
		goto L23
	} else {
		goto L24
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320))) = v259
	F_errmsg_internal(m, int32(_a_F_pg_stat_io_build_tuples_5), v320)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_pg_stat_io_build_tuples_2), int32(248), int32(_a_F_pg_stat_io_build_tuples_6))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	if v416 != 0 {
		goto L37
	} else {
		goto L38
	}
L24:
	;
	if v356|base.B2i32(v358&int32(_a_F_pg_stat_io_build_tuples_7) == int32(0)) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	if l2&int32(-2) == int32(10) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	if base.B2i32(v259 == int32(3))&v372 != 0 {
		v416 = v345
		goto L23
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(int32(2)) < base.Ui32(l2-int32(14)) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if v166 != int32(2) {
		v416 = v345
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v416 = base.B2i32(v259 != int32(1)) | base.B2i32(base.Ui32(l2-int32(5)) < base.Ui32(int32(-2)))
	goto L23
L31:
	;
	if base.B2i32(int32(1)<<(uint(v259)%32)&int32(19) == int32(0))|base.B2i32(base.Ui32(int32(4)) < base.Ui32(v259)) != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if l2 != int32(3) {
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v416 = v345
	goto L23
L35:
	;
	if v259 == int32(4) {
		v416 = v345
		goto L23
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v68)+304)) = base.I64_extend_i32_u(v151)
	v418 = F_cstring_to_text(m, v339)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v819 = v259 + int32(1)
	if v819 != int32(5) {
		v259 = v819
		goto L15
	} else {
		goto L114
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v68)+320)) = base.I64_extend_i32_u(v418)
	v422 = F_cstring_to_text(m, v248)
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v68)+312)) = base.I64_extend_i32_u(v422)
	if l3 != int64(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v432 = v259 << (uint(int32(6)) % 32)
	v445 = int32(0)
	goto L46
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v68)+456)) = l3
	goto L42
L44:
	;
	goto L45
L45:
	;
	v429 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+291)) = uint8(v429)
	goto L42
L46:
	;
	v502 = int32(0)
	switch v445 - int32(1) {
	case 0:
		goto L54
	case 1:
		goto L58
	case 2:
		goto L57
	case 3:
		goto L56
	case 4:
		goto L55
	case 5:
		v526 = v68 + int32(328)
		v527 = v68 + int32(344)
		v528 = v109 | int32(5)
		v529 = v109 | int32(3)
		v530 = v109 | int32(4)
		v531 = v502
		v532 = v68 + int32(336)
		v533 = v502
		goto L48
	case 6:
		goto L52
	default:
		goto L53
	}
L47:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_tuplestore_putvalues(m, v745, v746, v68+int32(304), v68+int32(272))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L5
	} else {
		goto L113
	}
L48:
	;
	v534 = int32(0)
	v536 = base.B2i32(base.Ui32(int32(16)) < base.Ui32(l2))
	v537 = int32(1)
	v538 = v537 << (uint(l2) % 32)
	v552 = base.B2i32(v166 == v537)
	if v536|base.B2i32(v538&int32(_a_F_pg_stat_io_build_tuples_4) == v534)|(base.B2i32(v166 == int32(2))&base.B2i32(base.Ui32(v259-int32(4)) < base.Ui32(int32(-2)))|v552&base.B2i32(v259 != int32(3))) != 0 {
		v698 = v534
		goto L59
	} else {
		goto L60
	}
L49:
	;
	v526 = v522
	v527 = v91
	v528 = v127
	v529 = v523
	v530 = v127
	v531 = v524
	v532 = v91
	v533 = int32(1)
	goto L48
L50:
	;
	v526 = v517
	v527 = v518
	v528 = v519
	v529 = v520
	v530 = v127
	v531 = v502
	v532 = v91
	v533 = v521
	goto L48
L51:
	;
	v526 = v511
	v527 = v512
	v528 = v513
	v529 = v516
	v530 = v514
	v531 = v502
	v532 = v515
	v533 = v502
	goto L48
L52:
	;
	v511 = v68 + int32(352)
	v512 = v68 + int32(368)
	v513 = v109 | int32(8)
	v514 = v109 | int32(7)
	v515 = v68 + int32(360)
	v516 = v109 | int32(6)
	goto L51
L53:
	;
	v522 = v68 + int32(424)
	v523 = v109 | int32(15)
	v524 = int32(1)
	goto L49
L54:
	;
	v517 = v68 + int32(440)
	v518 = v68 + int32(448)
	v519 = v68 + int32(290)
	v520 = v68 + int32(289)
	v521 = int32(1)
	goto L50
L55:
	;
	v511 = v68 + int32(392)
	v512 = v68 + int32(408)
	v513 = v109 | int32(13)
	v514 = v109 | int32(12)
	v515 = v68 + int32(400)
	v516 = v109 | int32(11)
	goto L51
L56:
	;
	v517 = v68 + int32(376)
	v518 = v68 + int32(384)
	v519 = v109 | int32(10)
	v520 = v109 | int32(9)
	v521 = int32(1)
	goto L50
L57:
	;
	v522 = v68 + int32(432)
	v523 = v68 + int32(288)
	v524 = int32(1)
	goto L49
L58:
	;
	v522 = v68 + int32(416)
	v523 = v109 | int32(14)
	v524 = int32(1)
	goto L49
L59:
	;
	if v698 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L60:
	;
	if v536|base.B2i32(v538&int32(_a_F_pg_stat_io_build_tuples_7) == int32(0)) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v573 = l2 & int32(-2)
	if v573 == int32(10) {
		goto L67
	} else {
		goto L68
	}
L62:
	;
	if base.B2i32(v259 == int32(3))&v552 != 0 {
		v698 = v534
		goto L59
	} else {
		goto L63
	}
L63:
	;
	if base.Ui32(int32(2)) < base.Ui32(l2-int32(14)) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	if v166 != int32(2) {
		v698 = v534
		goto L59
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	switch l2 - int32(10) {
	case 0:
		goto L76
	case 1:
		goto L75
	default:
		goto L74
	}
L67:
	;
	if base.B2i32(int32(1)<<(uint(v259)%32)&int32(19) == int32(0))|base.B2i32(base.Ui32(int32(4)) < base.Ui32(v259)) != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if base.B2i32(l2 == int32(3))&base.B2i32(v259 == int32(4)) != 0 {
		v698 = v534
		goto L59
	} else {
		goto L71
	}
L70:
	;
	v698 = v534
	goto L59
L71:
	;
	if base.Ui32(l2-int32(5)) < base.Ui32(int32(-2)) {
		goto L66
	} else {
		goto L72
	}
L72:
	;
	if v259 == int32(1) {
		v698 = v534
		goto L59
	} else {
		goto L73
	}
L73:
	;
	goto L66
L74:
	;
	if base.B2i32(v573 == int32(10))&base.B2i32(v445 == int32(5)) != 0 {
		v698 = v534
		goto L59
	} else {
		goto L79
	}
L75:
	;
	if base.B2i32(v445&int32(-3) == int32(0))|base.B2i32(v445 == int32(6))&base.B2i32(v166 != int32(2)) != 0 {
		v698 = v534
		goto L59
	} else {
		goto L78
	}
L76:
	;
	if base.B2i32(int32(1)<<(uint(v445)%32)&int32(69) == int32(0))|base.B2i32(base.Ui32(int32(6)) < base.Ui32(v445)) != 0 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v698 = v534
	goto L59
L78:
	;
	goto L74
L79:
	;
	if base.B2i32(v166 != int32(2))|base.B2i32(v445 != int32(6)) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v670 = int32(2)
	v671 = base.B2i32(v166 != v670)
	if v671|base.B2i32(v259 != v670) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L81:
	;
	v630 = l2 - int32(3)
	if base.B2i32(base.Ui32(v630) <= base.Ui32(int32(13)))&(int32(base.Ui32(int32(_a_F_pg_stat_io_build_tuples_8))>>(uint(v630)%32))&int32(1)) != 0 {
		v698 = v534
		goto L59
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	if v166 != int32(1) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v669 = base.B2i32(v259 == int32(4)) | base.B2i32(base.Ui32(v259) < base.Ui32(int32(2)))
	goto L80
L85:
	;
	if base.B2i32(v259 == int32(0))&base.B2i32(v445 == int32(5)) != 0 {
		v698 = v534
		goto L59
	} else {
		goto L87
	}
L86:
	;
	switch v445 - int32(1) {
	case 0, 3:
		v698 = v534
		goto L59
	default:
		goto L85
	}
L87:
	;
	v656 = base.B2i32(v259 == int32(4)) | base.B2i32(base.Ui32(v259) < base.Ui32(int32(2)))
	if v445 != int32(3) {
		v669 = v656
		goto L80
	} else {
		goto L88
	}
L88:
	;
	if base.B2i32(int32(1)<<(uint(v259)%32)&int32(19) == int32(0))|base.B2i32(base.Ui32(int32(4)) < base.Ui32(v259)) != 0 {
		v698 = v534
		goto L59
	} else {
		goto L89
	}
L89:
	;
	v669 = v656
	goto L80
L90:
	;
	v698 = base.B2i32(v669 == int32(0)) | base.B2i32(v445 != int32(1))
	goto L59
L91:
	;
	switch v445 - int32(1) {
	case 0, 6:
		goto L90
	default:
		v698 = v534
		goto L59
	}
L92:
	;
	goto L93
L93:
	;
	if v671|base.B2i32(v259 != int32(3)) != 0 {
		goto L90
	} else {
		goto L94
	}
L94:
	;
	if base.B2i32(int32(1)<<(uint(v445)%32)&int32(194) == int32(0))|base.B2i32(base.Ui32(int32(7)) < base.Ui32(v445)) != 0 {
		v698 = v534
		goto L59
	} else {
		goto L95
	}
L95:
	;
	goto L90
L96:
	;
	v742 = v445 + int32(1)
	if v742 != int32(8) {
		v445 = v742
		goto L46
	} else {
		goto L112
	}
L97:
	;
	if v531 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L98:
	;
	v701 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v529))) = uint8(v701)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v704 = v445 << (uint(int32(3)) % 32)
	v706 = *(*int64)(unsafe.Add(mBase, uint32(v432+(v223+(l1+int32(960)))+v704)))
	*(*int64)(unsafe.Add(mBase, uint32(v526))) = v706
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
	if v708 != 0 {
		goto L97
	} else {
		goto L101
	}
L101:
	;
	if v531 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v712 = *(*int64)(unsafe.Add(mBase, uint32(v704+(v432+(v223+(l1+int32(1920)))))))
	*(*float64)(unsafe.Add(mBase, uint32(v527))) = base.F64_mul(base.F64_convert_i64_s(v712), float64(0.001))
	goto L104
L103:
	;
	goto L104
L104:
	;
	if v533 != 0 {
		goto L96
	} else {
		goto L105
	}
L105:
	;
	v718 = *(*int64)(unsafe.Add(mBase, uint32(v704+(l1+v223+v432))))
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v718
	v724 = F_pg_snprintf(m, v68+int32(16), int32(256), int32(_a_F_pg_stat_io_build_tuples_9), v68)
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	v730 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v68+int32(16)), int64(0), int64(-1))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L5
	} else {
		goto L107
	}
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v532))) = v730
	goto L96
L108:
	;
	v736 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v528))) = uint8(v736)
	goto L110
L109:
	;
	goto L110
L110:
	;
	if v533 != 0 {
		goto L96
	} else {
		goto L111
	}
L111:
	;
	v738 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v530))) = uint8(v738)
	goto L96
L112:
	;
	goto L47
L113:
	;
	goto L39
L114:
	;
	goto L16
L115:
	;
	goto L8
}
func F_pg_stat_reset(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	v6 = m.G0
	v7 = int32(16)
	v8 = v6 - v7
	m.G0 = v8
	F_gettimeofday(m, v8)
	mBase = m.M
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	v12 = int64(*(*int32)(unsafe.Add(mBase, uint32(v8)+8)))
	m.G0 = v8 + v7
	v23 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_pg_stat_reset[0])))
	F_pgstat_reset_matching_entries(m, int32(1316), v23, v12+v11*int64(1000000)-int64(946684800000000))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_reset_single_table_counters(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v32 int32
	_ = v32
	var v49 int32
	_ = v49
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = base.I32_wrap_i64(v3)
	v7 = int32(1)
	if v4 <= int32(3591) {
		if v4 <= int32(2670) {
			switch v4 - int32(1213) {
			case 0, 1, 19, 20, 47, 48, 49:
				v78 = v7
			case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
				v78 = int32(0)
			default:
				if base.Ui32(int32(2)) <= base.Ui32(v4-int32(2396)) {
					v78 = int32(0)
				} else {
					v78 = v7
				}
			}
		} else {
			v19 = v4 - int32(2671)
			if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v19))|base.B2i32(int32(1)<<(uint(v19)%32)&int32(226492515) == int32(0)) != 0 {
				if base.B2i32(base.Ui32(v4-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(v4-int32(2846)) < base.Ui32(int32(2))) != 0 {
					v78 = v7
				} else {
					v78 = int32(0)
				}
			} else {
				v78 = v7
			}
		}
	} else {
		if v4 <= int32(_a_F_pg_stat_reset_single_table_counters_0) {
			v32 = v4 - int32(_a_F_pg_stat_reset_single_table_counters_1)
			if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v32))|base.B2i32(int32(1)<<(uint(v32)%32)&int32(963) == int32(0)) != 0 {
				if base.Ui32(v4-int32(3592)) < base.Ui32(int32(2)) {
					v78 = v7
				} else {
					if base.Ui32(int32(2)) <= base.Ui32(v4-int32(4060)) {
						v78 = int32(0)
					} else {
						v78 = v7
					}
				}
			} else {
				v78 = v7
			}
		} else {
			switch v4 - int32(_a_F_pg_stat_reset_single_table_counters_2) {
			case 0, 1, 2, 3, 4, 59, 60:
				v78 = v7
			case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
				v78 = int32(0)
			default:
				if base.Ui32(v4-int32(_a_F_pg_stat_reset_single_table_counters_3)) < base.Ui32(int32(3)) {
					v78 = v7
				} else {
					v49 = v4 - int32(_a_F_pg_stat_reset_single_table_counters_4)
					if base.Ui32(int32(15)) < base.Ui32(v49) {
						v78 = int32(0)
					} else {
						if int32(1)<<(uint(v49)%32)&int32(_a_F_pg_stat_reset_single_table_counters_5) != 0 {
							v78 = v7
						} else {
							v78 = int32(0)
						}
					}
				}
			}
		}
	}
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_single_table_counters[0]))
	if v78 != 0 {
		v83 = int32(0)
	} else {
		v83 = v82
	}
	F_pgstat_reset(m, int32(2), v83, v3&int64(4294967295))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_1_11(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pg_stat_statements_internal(m, l0, int32(7), base.B2i32(v3 != int64(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
