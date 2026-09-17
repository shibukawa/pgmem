package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ProcessWalSndrMessage(m *base.Module, l0 int64, l1 int64) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int64
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v156 int64
	_ = v156
	var v162 int64
	_ = v162
	var v171 int64
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[0]))
	v18 = m.G0
	v19 = int32(16)
	v20 = v18 - v19
	m.G0 = v20
	F_gettimeofday(m, v20)
	mBase = m.M
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v24 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+8)))
	m.G0 = v20 + v19
	v32 = v24 + v23*int64(1000000) - int64(946684800000000)
	goto L1
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1456)) = int32(1)
	if v33 != 0 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	F_s_lock(m, v14+int32(1456), int32(_a_F_ProcessWalSndrMessage_0), int32(1263), int32(_a_F_ProcessWalSndrMessage_1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v14)+88))
	if base.Ui64(v43) < base.Ui64(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return
L6:
	;
	goto L4
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+96)) = l1
	goto L9
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+1456)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+80)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = l0
	v51 = int32(13)
	goto L12
L10:
	;
	if v88 != 0 {
		goto L23
	} else {
		goto L24
	}
L11:
	;
	goto L10
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[1]))
	goto L15
L13:
	;
	v71 = int32(0)
	goto L20
L15:
	;
	goto L16
L16:
	;
	if int32(0)|base.B2i32(v58 == int32(15)) != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v58 <= v51 {
		v88 = int32(1)
		goto L11
	} else {
		goto L19
	}
L19:
	;
	goto L13
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[2]))
	if v75 != int32(2) {
		v88 = v71
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[3])))
	if v79&int32(1) != 0 {
		v88 = v71
		goto L11
	} else {
		goto L22
	}
L22:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[4]))
	v88 = int32(0) | base.B2i32(v85 <= v51)
	goto L11
L23:
	;
	v90 = F_timestamptz_to_str(m, l1)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	m.G0 = v11 + int32(32)
	return
L26:
	;
	v92 = F_pstrdup(m, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v94 = F_timestamptz_to_str(m, v32)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	v96 = F_pstrdup(m, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[0]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+1456))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+1456)) = int32(1)
	if v100 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v184 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L5
	} else {
		goto L49
	}
L31:
	;
	F_s_lock(m, v99+int32(1456), int32(_a_F_ProcessWalSndrMessage_2), int32(372), int32(_a_F_ProcessWalSndrMessage_3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v110 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+1456)) = v110
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v99)+48))
	v114 = F_GetXLogReplayRecPtr(m, v110)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	if v112 != v114 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[5]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+96)) = int32(1)
	if v119 != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v178 = int32(0)
	goto L38
L38:
	;
	v181 = v178
	goto L30
L39:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[5]))
	F_s_lock(m, v123+int32(96), int32(_a_F_ProcessWalSndrMessage_4), int32(_a_F_ProcessWalSndrMessage_5), int32(_a_F_ProcessWalSndrMessage_6))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L5
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessWalSndrMessage[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v132)+96)) = int32(0)
	v136 = *(*int64)(unsafe.Add(mBase, uint32(v132)+72))
	if v136 == int64(0) {
		v181 = int32(-1)
		goto L30
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v142 = m.G0
	v143 = int32(16)
	v144 = v142 - v143
	m.G0 = v144
	F_gettimeofday(m, v144)
	mBase = m.M
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v144)))
	v148 = int64(*(*int32)(unsafe.Add(mBase, uint32(v144)+8)))
	m.G0 = v144 + v143
	v156 = v148 + v147*int64(1000000) - int64(946684800000000)
	goto L44
L44:
	;
	if v156 <= v136 {
		v174 = int32(0)
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v178 = v174
	goto L38
L46:
	;
	goto L45
L47:
	;
	v162 = v156 - v136
	if base.B2i32(int64(0) < v136)^base.B2i32(v162 < v156)|base.B2i32(int64(2147483646000) < v162) != 0 {
		v174 = int32(2147483647)
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v171 = base.I64_div_s(v162+int64(999), int64(1000))
	v174 = base.I32_wrap_i64(v171)
	goto L46
L49:
	;
	if v181 == int32(-1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	F_pfree(m, v92)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L5
	} else {
		goto L62
	}
L51:
	;
	F_errfinish(m, int32(_a_F_ProcessWalSndrMessage_0), v214, int32(_a_F_ProcessWalSndrMessage_1))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L61
	}
L52:
	;
	if v184 == int32(0) {
		goto L50
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v184 == int32(0) {
		goto L50
	} else {
		goto L58
	}
L55:
	;
	v191 = F_GetReplicationTransferLatency(m)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v92
	F_errmsg_internal(m, int32(_a_F_ProcessWalSndrMessage_7), v11)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	v214 = int32(1287)
	goto L51
L58:
	;
	v202 = F_GetReplicationTransferLatency(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v202
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v92
	F_errmsg_internal(m, int32(_a_F_ProcessWalSndrMessage_8), v11+int32(16))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v214 = int32(1293)
	goto L51
L61:
	;
	goto L50
L62:
	;
	F_pfree(m, v96)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L5
	} else {
		goto L63
	}
L63:
	;
	goto L25
}
func F_WalRcvFetchTimeLineHistoryFiles(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	v9 = m.G0
	v11 = v9 - int32(128)
	m.G0 = v11
	if base.Ui32(l0) <= base.Ui32(l1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L9
	} else {
		goto L83
	}
L2:
	;
	v14 = l0
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v11 + int32(128)
	return
L5:
	;
	if v14 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v277 = v14 + int32(1)
	if base.Ui32(v277) <= base.Ui32(l1) {
		v14 = v277
		goto L5
	} else {
		goto L82
	}
L8:
	;
	v24 = F_existsTimeLineHistory(m, v14)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	if v24 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v28 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v14
	F_errmsg(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_0), v11+int32(32))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L9
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[0]))
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[1]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
	m.T0[v51].(func(*base.Module, int32, int32, int32, int32, int32))(m, v42, v14, v11+int32(124), v11+int32(120), v11+int32(116))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L9
	} else {
		goto L18
	}
L16:
	;
	F_errfinish(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_1), int32(741), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v14
	v56 = v11 + int32(48)
	v61 = F_pg_snprintf(m, v56, int32(64), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_3), v11+int32(16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if base.B2i32(v66 == int32(0))|base.B2i32(v66 != v69) != 0 {
		v87 = v66
		v88 = v69
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v87-v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v72 = v63
	v73 = v56
	goto L23
L23:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v77 == int32(0) {
		v87 = v77
		v88 = v76
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v87 = v77
	v88 = v76
	goto L21
L25:
	;
	v80 = int32(1)
	if v77 == v76 {
		v72 = v72 + v80
		v73 = v73 + v80
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v11)+116))
	v92 = m.G0
	v94 = v92 - int32(2144)
	m.G0 = v94
	*(*int32)(unsafe.Add(mBase, uint32(v94)+80)) = int32(42)
	v99 = v94 + int32(96)
	v104 = F_pg_snprintf(m, v99, int32(1024), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_4), v94+int32(80))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v106 = F_unlink(m, v99)
	mBase = m.M
	v108 = F_OpenTransientFile(m, v99, int32(194))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L9
	} else {
		goto L32
	}
L29:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[2]))
	if v258 != int32(2) {
		goto L75
	} else {
		goto L76
	}
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L70
	}
L31:
	;
	v211 = int32(_a_F_WalRcvFetchTimeLineHistoryFiles_5)
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[3]))
	v214 = v94 + int32(96)
	v215 = F_unlink(m, v214)
	mBase = m.M
	if v212 != 0 {
		goto L63
	} else {
		goto L64
	}
L32:
	;
	if int32(0) <= v108 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[3])) = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(167772218)
	v119 = F_write(m, v108, v90, v91)
	mBase = m.M
	if v119 != v91 {
		goto L31
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L9
	} else {
		goto L59
	}
L36:
	;
	v121 = int32(_a_F_WalRcvFetchTimeLineHistoryFiles_6)
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[4]))
	v123 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v122))) = v123
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = int32(167772217)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[5])))
	if v131 != int32(1) {
		v145 = v123
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = int32(0)
	v175 = F_CloseTransientFile(m, v108)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L55
	}
L38:
	;
	if v145 == int32(0) {
		goto L37
	} else {
		goto L45
	}
L39:
	;
	goto L38
L40:
	;
	goto L41
L41:
	;
	v136 = F_fsync(m, v108)
	mBase = m.M
	if v136 != int32(-1) {
		v145 = v136
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v145 = int32(-1)
	goto L39
L43:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[3]))
	if v140 == int32(27) {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[6])))
	if v151 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v154 = F_errstart(m, v152, int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L50
	}
L47:
	;
	v152 = int32(21)
	goto L49
L48:
	;
	v152 = int32(23)
	goto L49
L49:
	;
	goto L46
L50:
	;
	if v154 == int32(0) {
		goto L37
	} else {
		goto L51
	}
L51:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+48)) = v99
	F_errmsg(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_7), v94+int32(48))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_8), int32(506), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_9))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	goto L37
L55:
	;
	if v175 != 0 {
		goto L30
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v14
	v179 = v94 + int32(1120)
	v184 = F_pg_snprintf(m, v179, int32(1024), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_10), v94+int32(16))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L9
	} else {
		goto L57
	}
L57:
	;
	v189 = F_durable_rename(m, v94+int32(96), v179, int32(21))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	m.G0 = v94 + int32(2144)
	goto L29
L59:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = v94 + int32(96)
	F_errmsg(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_11), v94)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_8), int32(481), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_9))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	v218 = v212
	goto L65
L64:
	;
	v218 = int32(51)
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_WalRcvFetchTimeLineHistoryFiles[3])) = v218
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L9
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+64)) = v214
	F_errmsg(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_12), v94-int32(-64))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_8), int32(498), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_9))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+32)) = v94 + int32(96)
	F_errmsg(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_13), v94+int32(32))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L9
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_8), int32(512), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_9))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L9
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
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v11)+124))
	F_pfree(m, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L9
	} else {
		goto L80
	}
L75:
	;
	F_XLogArchiveForceDone(m, v256)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	F_XLogArchiveNotify(m, v256)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L9
	} else {
		goto L79
	}
L78:
	;
	goto L74
L79:
	;
	goto L74
L80:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v11)+120))
	F_pfree(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	goto L7
L82:
	;
	goto L6
L83:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
	F_errmsg_internal(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_14), v11)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_WalRcvFetchTimeLineHistoryFiles_1), int32(755), int32(_a_F_WalRcvFetchTimeLineHistoryFiles_2))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WalSummarizerShutdown(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerShutdown[0]))
	v8 = F_LWLockAcquire(m, v4+int32(_a_F_WalSummarizerShutdown_0), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerShutdown[1]))
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(-1)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_WalSummarizerShutdown[0]))
		F_LWLockRelease(m, v15+int32(_a_F_WalSummarizerShutdown_0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			return
		}
	}
}
func F_WalUsageAccumDiff(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int64)(unsafe.Add(mBase, _c_F_WalUsageAccumDiff[0]))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v3 + (v5 - v6)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_WalUsageAccumDiff[1]))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v10 + (v12 - v13)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int64)(unsafe.Add(mBase, _c_F_WalUsageAccumDiff[2]))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v17 + (v19 - v20)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_WalUsageAccumDiff[3]))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v24 + (v26 - v27)
	return
}
func F_WriteWalSummary(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int64
	_ = v26
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = l1
	v19 = F_FileWriteV(m, v12, v9+int32(40), int32(1), v11, int32(167772237))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v19 {
			if v19 != l2 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					F_errcode_for_file_access(m)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_WriteWalSummary[0]))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+v62*int32(48))+32))
						v69 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
						*(*uint32)(unsafe.Add(mBase, uint32(v9)+28)) = uint32(v69)
						*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v68
						F_errmsg(m, int32(_a_F_WriteWalSummary_0), v9+int32(16))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_WriteWalSummary_1), int32(0))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_WriteWalSummary_2), int32(312), int32(_a_F_WriteWalSummary_3))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				}
			} else {
				v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v26 + base.I64_extend_i32_u(l2)
				m.G0 = v9 + int32(48)
				return l2
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return int32(0)
			} else {
				F_errcode_for_file_access(m)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_WriteWalSummary[0]))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v40*int32(48))+32))
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
					F_errmsg(m, int32(_a_F_WriteWalSummary_4), v9)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_WriteWalSummary_2), int32(305), int32(_a_F_WriteWalSummary_3))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_check_wal_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(-1) {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_check_wal_buffers[0]))
		if v8 == int32(-1) {
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_check_wal_buffers[1]))
			v15 = base.I32_div_s(v13, int32(32))
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_check_wal_buffers[2]))
			v19 = base.I32_div_s(v17, int32(_a_F_check_wal_buffers_0))
			if v15 < v19 {
				v21 = v15
			} else {
				v21 = v19
			}
			if v21 <= int32(8) {
				v24 = int32(8)
			} else {
				v24 = v21
			}
			v29 = v24
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
		}
	} else {
		if int32(3) < v4 {
		} else {
			v29 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
		}
	}
	return int32(1)
}
func F_compareWalFileNames(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = int32(8)
	v5 = v3 + v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = v6 + v4
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
	if base.B2i32(v11 == int32(0))|base.B2i32(v11 != v14) != 0 {
		v32 = v11
		v33 = v14
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v32 - v33
L2:
	;
	goto L1
L3:
	;
	v17 = v5
	v18 = v8
	goto L4
L4:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v22
		v33 = v21
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v32 = v22
	v33 = v21
	goto L2
L6:
	;
	v25 = int32(1)
	if v22 == v21 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
