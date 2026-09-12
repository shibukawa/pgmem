package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_replication_origin_create(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v10 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L9
	} else {
		goto L50
	}
L2:
	;
	if v20 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+316))
	v18 = base.B2i32(v16 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v18)
	v20 = v18
	goto L5
L4:
	;
	v20 = int32(0)
	goto L5
L5:
	;
	goto L2
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_text_to_cstring(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L46
	}
L9:
	;
	return int32(0)
L10:
	;
	v28 = int32(0)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v29 != int32(112) {
		v38 = v28
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v38 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	goto L11
L13:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v32 != int32(103) {
		v38 = v28
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)))
	v38 = base.B2i32(v35 == int32(95))
	goto L12
L15:
	;
	v42 = v24
	v43 = int32(356178)
	goto L17
L16:
	;
	if v80 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L17:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v46 == v47 {
		v69 = v46
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v80 = int32(0)
	goto L16
L19:
	;
	v71 = int32(1)
	if v69 != 0 {
		v42 = v42 + v71
		v43 = v43 + v71
		goto L17
	} else {
		goto L28
	}
L20:
	;
	if base.Ui32((v46-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v57 = v46 | int32(32)
	goto L23
L22:
	;
	v57 = v46
	goto L23
L23:
	;
	if base.Ui32((v47-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v66 = v47 | int32(32)
	goto L26
L25:
	;
	v66 = v47
	goto L26
L26:
	;
	if v57 == v66 {
		v69 = v57
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v80 = v57 - v66
	goto L16
L28:
	;
	goto L18
L29:
	;
	v86 = v24
	v87 = int32(17208)
	goto L31
L30:
	;
	if v124 == int32(0) {
		goto L1
	} else {
		goto L43
	}
L31:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v90 == v91 {
		v113 = v90
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v124 = int32(0)
	goto L30
L33:
	;
	v115 = int32(1)
	if v113 != 0 {
		v86 = v86 + v115
		v87 = v87 + v115
		goto L31
	} else {
		goto L42
	}
L34:
	;
	if base.Ui32((v90-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v101 = v90 | int32(32)
	goto L37
L36:
	;
	v101 = v90
	goto L37
L37:
	;
	if base.Ui32((v91-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v110 = v91 | int32(32)
	goto L40
L39:
	;
	v110 = v91
	goto L40
L40:
	;
	if v101 == v110 {
		v113 = v101
		goto L33
	} else {
		goto L41
	}
L41:
	;
	v124 = v101 - v110
	goto L30
L42:
	;
	goto L32
L43:
	;
	v127 = F_replorigin_create(m, v24)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L9
	} else {
		goto L44
	}
L44:
	;
	F_pfree(m, v24)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L9
	} else {
		goto L45
	}
L45:
	;
	m.G0 = v6 + int32(32)
	return v127
L46:
	;
	F_errcode(m, int32(100663618))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(13407), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(474952), int32(200), int32(150593))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(151818372))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v24
	F_errmsg(m, int32(420806), v6+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(356178)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(17208)
	F_errdetail(m, int32(607902), v6)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(474952), int32(1311), int32(339497))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_replication_origin_session_reset(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v15 int32
	_ = v15
	F_replorigin_check_prerequisites(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		F_replorigin_session_reset(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			v9 = int64(0)
			*(*int64)(unsafe.Add(mBase, _consts[171])) = v9
			*(*int64)(unsafe.Add(mBase, _consts[172])) = v9
			v15 = int32(0)
			*(*uint16)(unsafe.Add(mBase, _consts[170])) = uint16(v15)
			return v15
		}
	}
}
func F_pg_replication_slot_advance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v61 int64
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v290 int64
	_ = v290
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_CheckSlotPermissions(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v11 != int64(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L80
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L75
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L72
	}
L6:
	;
	v22 = F_get_call_result_type(m, l0, int32(0), v8+int32(44))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L68
	}
L9:
	;
	if v22 != int32(1) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v28 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v92 = int32(1)
	F_ReplicationSlotAcquire(m, v12, v92, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L31
	}
L12:
	;
	if v38 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+316))
	v36 = base.B2i32(v34 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v36)
	v38 = v36
	goto L15
L14:
	;
	v38 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v44 = int32(4339424)
	v45 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v45)+280)) = v46
	*(*int64)(unsafe.Add(mBase, _consts[181])) = v46
	v51 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v51)+272)) = v52
	*(*int64)(unsafe.Add(mBase, _consts[183])) = v52
	goto L21
L17:
	;
	goto L18
L18:
	;
	v85 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L28
	}
L19:
	;
	if base.Ui64(v11) < base.Ui64(v61) {
		v91 = v11
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v61 = *(*int64)(unsafe.Add(mBase, _consts[181]))
	goto L19
L23:
	;
	v66 = int32(4339424)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v67)+280))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+280)) = v68
	*(*int64)(unsafe.Add(mBase, _consts[181])) = v68
	v73 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v73)+272))
	*(*int64)(unsafe.Add(mBase, uint32(v73)+272)) = v74
	*(*int64)(unsafe.Add(mBase, _consts[183])) = v74
	goto L26
L24:
	;
	v91 = v83
	goto L11
L26:
	;
	goto L27
L27:
	;
	v83 = *(*int64)(unsafe.Add(mBase, _consts[181]))
	goto L24
L28:
	;
	if base.Ui64(v11) < base.Ui64(v85) {
		v91 = v11
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v89 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v91 = v89
	goto L11
L31:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v97)+104))
	if v98 == int64(0) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)+88))
	if v101 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+34)) = uint8(v194)
	v197 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v197 + int32(24)
	F_ReplicationSlotsComputeRequiredXmin(m, v194)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L62
	}
L34:
	;
	if base.Ui64(v91) < base.Ui64(v98) {
		v276 = v98
		goto L3
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v97)+120))
	if base.Ui64(v91) < base.Ui64(v186) {
		v276 = v186
		goto L3
	} else {
		goto L60
	}
L37:
	;
	if base.Ui64(v91) <= base.Ui64(v98) {
		v193 = v98
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = int32(1)
	if v106 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	F_s_lock(m, v110, int32(473123), int32(474), int32(398773))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v117)+104)) = v91
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, _consts[182])))
	if v125 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v193 = v91
	goto L33
L45:
	;
	if v135 != 0 {
		goto L44
	} else {
		goto L49
	}
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[175]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+316))
	v133 = base.B2i32(v131 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _consts[182])) = uint8(v133)
	v135 = v133
	goto L48
L47:
	;
	v135 = int32(0)
	goto L48
L48:
	;
	goto L45
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	v140 = int32(0)
	v146 = *(*int32)(unsafe.Add(mBase, _consts[681]))
	if v146 == v140 {
		v174 = v140
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v174 == int32(0) {
		goto L44
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v149 <= int32(0) {
		v174 = v140
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v155 = v146 + int32(4)
	v157 = v140
	goto L54
L54:
	;
	v160 = F_strcmp(m, v155, v137+int32(24))
	mBase = m.M
	v162 = base.B2i32(v160 == int32(0))
	if v160 == int32(0) {
		v174 = v162
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v174 = v162
	goto L51
L56:
	;
	v165 = F_strlen(m, v155)
	mBase = m.M
	v167 = int32(1)
	v170 = v157 + v167
	if v170 != v149 {
		v155 = v165 + v155 + v167
		v157 = v170
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[682]))
	F_ConditionVariableBroadcast(m, v181+int32(76))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L44
L60:
	;
	v189 = F_LogicalSlotAdvanceAndCheckSnapState(m, v91, int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v193 = v189
	goto L33
L62:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v208 = F_Int64GetDatum(m, v193)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v210 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+35)) = uint8(v210)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v208
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
	v218 = F_heap_form_tuple(m, v213, v8+int32(36), v8+int32(34))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v218)+16))
	v221 = F_HeapTupleHeaderGetDatum(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	m.G0 = v8 + int32(48)
	return v221
L68:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(505989), int32(0))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(473123), int32(529), int32(398810))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(351051), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(473123), int32(533), int32(398810))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
	F_errmsg(m, int32(443333), v8)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errdetail(m, int32(608909), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(473123), int32(553), int32(398810))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
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
	F_errcode(m, int32(325))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+28)) = uint32(v276)
	v285 = int64(32)
	v286 = int64(base.Ui64(v276) >> (uint(v285) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+24)) = uint32(v286)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v91)
	v290 = int64(base.Ui64(v91) >> (uint(v285) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)) = uint32(v290)
	F_errmsg(m, int32(491927), v8+int32(16))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(473123), int32(570), int32(398810))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_show_replication_origin_status(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
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
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v22 = F_LWLockAcquire(m, v18+int32(5120), int32(1))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	if int32(0) < v25 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v32 = v25
	v33 = v29
	v35 = v2
	goto L7
L5:
	;
	goto L6
L6:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	F_LWLockRelease(m, v114+int32(5120))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L25
	}
L7:
	;
	v38 = v33 + v35*int32(56)
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v40 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v40
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v40
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(16843009)
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+12)) = uint8(v47)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v46
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v52 = F_SearchSysCache1(m, int32(58), v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v102 = v32
	v103 = v33
	goto L11
L11:
	;
	v105 = v35 + int32(1)
	if v105 < v102 {
		v32 = v102
		v33 = v103
		v35 = v105
		goto L7
	} else {
		goto L24
	}
L12:
	;
	if v52 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+22)))
	v59 = F_text_to_cstring(m, v54+v55+int32(4))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v71 = v38 + int32(40)
	v73 = F_LWLockAcquire(m, v71, int32(1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	F_ReleaseCatCache(m, v52)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v63 = F_cstring_to_text(m, v59)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+13)) = uint8(v65)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v63
	goto L15
L19:
	;
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v38)+8))
	v76 = F_Int64GetDatum(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v78)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v76
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v38)+16))
	v82 = F_Int64GetDatum(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v84 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v84)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v82
	F_LWLockRelease(m, v71)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	F_tuplestore_putvalues(m, v89, v90, v9+int32(16), v9+int32(12))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[644]))
	v100 = *(*int32)(unsafe.Add(mBase, _consts[643]))
	v102 = v100
	v103 = v98
	goto L11
L24:
	;
	goto L8
L25:
	;
	m.G0 = v9 + int32(32)
	return int32(0)
}
