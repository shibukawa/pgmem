package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pg_replication_origin_create(m *base.Module, l0 int32) int64 {
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
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_origin_create[0])))
	if v10 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_origin_create[1]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+308))
	v18 = base.B2i32(v16 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_origin_create[0])) = uint8(v18)
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
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L46
	}
L9:
	;
	return int64(0)
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
	v43 = int32(_a_F_pg_replication_origin_create_0)
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
	v87 = int32(_a_F_pg_replication_origin_create_1)
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
	return base.I64_extend_i32_u(v127)
L46:
	;
	F_errcode(m, int32(100663618))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_pg_replication_origin_create_2), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_pg_replication_origin_create_3), int32(217), int32(_a_F_pg_replication_origin_create_4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
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
	v158 = m.ExcPending
	if v158 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v24
	F_errmsg(m, int32(_a_F_pg_replication_origin_create_5), v6+int32(16))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(_a_F_pg_replication_origin_create_0)
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_pg_replication_origin_create_1)
	v170 = F_errdetail(m, int32(_a_F_pg_replication_origin_create_6), v6)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_pg_replication_origin_create_3), int32(1417), int32(_a_F_pg_replication_origin_create_7))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
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
func F_pg_replication_origin_session_reset(m *base.Module, l0 int32) int64 {
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
		return int64(0)
	} else {
		F_replorigin_session_reset(m)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int64(0)
		} else {
			v9 = int64(0)
			*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_reset[0])) = v9
			*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_reset[1])) = v9
			v15 = int32(0)
			*(*uint16)(unsafe.Add(mBase, _c_F_pg_replication_origin_session_reset[2])) = uint16(v15)
			return v9
		}
	}
}
func F_pg_replication_slot_advance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v47 int64
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v66 int64
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int64
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int64
	_ = v194
	var v197 int64
	_ = v197
	var v198 int32
	_ = v198
	var v200 int64
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v282 int64
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v296 int64
	_ = v296
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_CheckSlotPermissions(m)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v10 != int64(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L79
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L74
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L71
	}
L6:
	;
	v21 = F_get_call_result_type(m, l0, int32(0), v8+int32(76))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
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
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L67
	}
L9:
	;
	if v21 != int32(1) {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[0])))
	if v27 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v103 = int32(1)
	F_ReplicationSlotAcquire(m, v11, v103, v103)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L12:
	;
	if v37 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[1]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+308))
	v35 = base.B2i32(v33 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[0])) = uint8(v35)
	v37 = v35
	goto L15
L14:
	;
	v37 = int32(0)
	goto L15
L15:
	;
	goto L12
L16:
	;
	v40 = int32(0)
	v42 = int32(_a_F_pg_replication_slot_advance_0)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[1]))
	v44 = int64(0)
	v47 = base.AtomicRmwCmpxchg64(m, v43, int32(272), v44, v44)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[2])) = v47
	v52 = base.AtomicRmwOr32(m, v40, int32(_a_F_pg_replication_slot_advance_1), v40)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[1]))
	v59 = base.AtomicRmwCmpxchg64(m, v55, int32(264), v44, v44)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[3])) = v59
	goto L21
L17:
	;
	goto L18
L18:
	;
	v96 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L28
	}
L19:
	;
	if base.Ui64(v10) < base.Ui64(v66) {
		v102 = v10
		goto L11
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v66 = *(*int64)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[2]))
	goto L19
L23:
	;
	v68 = int32(0)
	v70 = int32(_a_F_pg_replication_slot_advance_0)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[1]))
	v72 = int64(0)
	v75 = base.AtomicRmwCmpxchg64(m, v71, int32(272), v72, v72)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[2])) = v75
	v80 = base.AtomicRmwOr32(m, v68, int32(_a_F_pg_replication_slot_advance_1), v68)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[1]))
	v87 = base.AtomicRmwCmpxchg64(m, v83, int32(264), v72, v72)
	*(*int64)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[3])) = v87
	goto L26
L24:
	;
	v102 = v94
	goto L11
L26:
	;
	goto L27
L27:
	;
	v94 = *(*int64)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[2]))
	goto L24
L28:
	;
	if base.Ui64(v10) < base.Ui64(v96) {
		v102 = v10
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v100 = F_GetXLogReplayRecPtr(m, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v102 = v100
	goto L11
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[4]))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v108)+104))
	if v109 == int64(0) {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+88))
	if v112 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v201 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+46)) = uint8(v201)
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = base.I64_extend_i32_u(v204 + int32(24))
	F_ReplicationSlotsComputeRequiredXmin(m, v201)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L62
	}
L34:
	;
	if base.Ui64(v102) < base.Ui64(v109) {
		v282 = v109
		goto L3
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v108)+120))
	if base.Ui64(v102) < base.Ui64(v194) {
		v282 = v194
		goto L3
	} else {
		goto L60
	}
L37:
	;
	if base.Ui64(v102) <= base.Ui64(v109) {
		v200 = v109
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v119 = base.AtomicRmwXchg32(m, v108, int32(0), int32(1))
	if v119 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_s_lock(m, v108, int32(_a_F_pg_replication_slot_advance_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v124)+104)) = v102
	v126 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v124))), uint32(v126))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[0])))
	if v133 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v200 = v102
	goto L33
L45:
	;
	if v143 != 0 {
		goto L44
	} else {
		goto L49
	}
L46:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[1]))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+308))
	v141 = base.B2i32(v139 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[0])) = uint8(v141)
	v143 = v141
	goto L48
L47:
	;
	v143 = int32(0)
	goto L48
L48:
	;
	goto L45
L49:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[4]))
	v148 = int32(0)
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[5]))
	if v154 == v148 {
		v184 = v148
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if v184 == int32(0) {
		goto L44
	} else {
		goto L58
	}
L51:
	;
	goto L50
L52:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if v157 <= int32(0) {
		v184 = v148
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v163 = v154 + int32(4)
	v167 = v148
	goto L54
L54:
	;
	v168 = F_strcmp(m, v163, v145+int32(24))
	mBase = m.M
	v170 = base.B2i32(v168 == int32(0))
	if v168 == int32(0) {
		v184 = v170
		goto L51
	} else {
		goto L56
	}
L55:
	;
	v184 = v170
	goto L51
L56:
	;
	v173 = F_strlen(m, v163)
	mBase = m.M
	v175 = int32(1)
	v178 = v167 + v175
	if v178 != v157 {
		v163 = v173 + v163 + v175
		v167 = v178
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_pg_replication_slot_advance[6]))
	F_ConditionVariableBroadcast(m, v189+int32(76))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	goto L44
L60:
	;
	v197 = F_LogicalSlotAdvanceAndCheckSnapState(m, v102, int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v200 = v197
	goto L33
L62:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+47)) = uint8(v216)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v200
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	v224 = F_heap_form_tuple(m, v219, v8+int32(48), v8+int32(46))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
	v227 = F_HeapTupleHeaderGetDatum(m, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	m.G0 = v8 + int32(80)
	return v227
L67:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_pg_replication_slot_advance_3), int32(0))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_pg_replication_slot_advance_4), int32(563), int32(_a_F_pg_replication_slot_advance_5))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_pg_replication_slot_advance_6), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_pg_replication_slot_advance_4), int32(567), int32(_a_F_pg_replication_slot_advance_5))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
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
	F_errcode(m, int32(325))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	F_errmsg(m, int32(_a_F_pg_replication_slot_advance_7), v8)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v275 = F_errdetail(m, int32(_a_F_pg_replication_slot_advance_8), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_pg_replication_slot_advance_4), int32(587), int32(_a_F_pg_replication_slot_advance_5))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+28)) = uint32(v282)
	v291 = int64(32)
	v292 = int64(base.Ui64(v282) >> (uint(v291) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+24)) = uint32(v292)
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+20)) = uint32(v102)
	v296 = int64(base.Ui64(v102) >> (uint(v291) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v8)+16)) = uint32(v296)
	F_errmsg(m, int32(_a_F_pg_replication_slot_advance_9), v8+int32(16))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_pg_replication_slot_advance_4), int32(604), int32(_a_F_pg_replication_slot_advance_5))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_show_replication_origin_status(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_pg_show_replication_origin_status[0]))
	v23 = F_LWLockAcquire(m, v19+int32(_a_F_pg_show_replication_origin_status_0), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_show_replication_origin_status[1]))
	if int32(0) < v26 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_show_replication_origin_status[2]))
	v33 = v26
	v34 = v30
	v36 = v2
	goto L7
L5:
	;
	goto L6
L6:
	;
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_pg_show_replication_origin_status[0]))
	F_LWLockRelease(m, v118+int32(_a_F_pg_show_replication_origin_status_0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	v40 = v34 + v36<<(uint(int32(6))%32)
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40))))
	if v41 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v42 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(16843009)
	v52 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v40))))
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)) = uint8(v53)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v52
	v57 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v40))))
	v58 = F_SearchSysCache1(m, int32(58), v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v104 = v33
	v105 = v34
	goto L11
L11:
	;
	v108 = v36 + int32(1)
	if v108 < v104 {
		v33 = v104
		v34 = v105
		v36 = v108
		goto L7
	} else {
		goto L22
	}
L12:
	;
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+22)))
	v65 = F_text_to_cstring(m, v60+v61+int32(4))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v78 = v40 + int32(44)
	v80 = F_LWLockAcquire(m, v78, int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L19
	}
L16:
	;
	F_ReleaseCatCache(m, v58)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v69 = F_cstring_to_text(m, v65)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v71 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)) = uint8(v71)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = base.I64_extend_i32_u(v69)
	goto L15
L19:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	v83 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v83)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v82
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v40)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v83)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v86
	F_LWLockRelease(m, v78)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	F_tuplestore_putvalues(m, v92, v93, v10+int32(16), v10+int32(12))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_pg_show_replication_origin_status[2]))
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_pg_show_replication_origin_status[1]))
	v104 = v103
	v105 = v101
	goto L11
L22:
	;
	goto L8
L23:
	;
	m.G0 = v10 + int32(48)
	return int64(0)
}
