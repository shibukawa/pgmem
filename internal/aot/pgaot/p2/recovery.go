package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_GetRecoveryPauseState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_GetRecoveryPauseState[0]))
	v7 = base.AtomicRmwXchg32(m, v4, int32(96), int32(1))
	if v7 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_GetRecoveryPauseState[0]))
		F_s_lock(m, v9+int32(96), int32(_a_F_GetRecoveryPauseState_0), int32(3096), int32(_a_F_GetRecoveryPauseState_1))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetRecoveryPauseState[0]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
			v22 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v20)+96)), uint32(v22))
			return v21
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetRecoveryPauseState[0]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
		v22 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v20)+96)), uint32(v22))
		return v21
	}
}
func F_check_recovery_target_name(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = F_strlen(m, v8)
	mBase = m.M
	v11 = base.B2i32(base.Ui32(v9) < base.Ui32(int32(64)))
	if v11 == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_name[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_name[1])) = v15
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = int32(63)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(_a_F_check_recovery_target_name_0)
		v24 = F_format_elog_string(m, int32(_a_F_check_recovery_target_name_1), v6)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_name[2])) = v24
			m.G0 = v6 + int32(16)
			return v11
		}
	} else {
		m.G0 = v6 + int32(16)
		return v11
	}
}
func F_check_recovery_target_time(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int64
	_ = v194
	var v195 int64
	_ = v195
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v218 int64
	_ = v218
	var v219 int64
	_ = v219
	var v226 int32
	_ = v226
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	v6 = m.G0
	v8 = v6 - int32(464)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	if v11 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(464)
	return v267
L2:
	;
	v267 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v15 = int32(0)
	v16 = int32(_a_F_check_recovery_target_time_0)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_time[0])))
	if base.B2i32(v19 == v15)|base.B2i32(v19 != v22) != 0 {
		v40 = v19
		v41 = v22
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v40-v41 == int32(0) {
		v267 = v15
		goto L1
	} else {
		goto L12
	}
L6:
	;
	goto L5
L7:
	;
	v25 = v10
	v26 = v16
	goto L8
L8:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v40 = v30
		v41 = v29
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v40 = v30
	v41 = v29
	goto L6
L10:
	;
	v33 = int32(1)
	if v30 == v29 {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v45 = int32(_a_F_check_recovery_target_time_1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_time[1])))
	if base.B2i32(v48 == int32(0))|base.B2i32(v48 != v51) != 0 {
		v69 = v48
		v70 = v51
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v69-v70 == int32(0) {
		v267 = v15
		goto L1
	} else {
		goto L20
	}
L14:
	;
	goto L13
L15:
	;
	v54 = v10
	v55 = v45
	goto L16
L16:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v59 == int32(0) {
		v69 = v59
		v70 = v58
		goto L14
	} else {
		goto L18
	}
L17:
	;
	v69 = v59
	v70 = v58
	goto L14
L18:
	;
	v62 = int32(1)
	if v59 == v58 {
		v54 = v54 + v62
		v55 = v55 + v62
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v74 = int32(_a_F_check_recovery_target_time_2)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_time[2])))
	if base.B2i32(v77 == int32(0))|base.B2i32(v77 != v80) != 0 {
		v98 = v77
		v99 = v80
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v98-v99 == int32(0) {
		v267 = v15
		goto L1
	} else {
		goto L28
	}
L22:
	;
	goto L21
L23:
	;
	v83 = v10
	v84 = v74
	goto L24
L24:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v88
		v99 = v87
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v98 = v88
	v99 = v87
	goto L22
L26:
	;
	v91 = int32(1)
	if v88 == v87 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v103 = int32(_a_F_check_recovery_target_time_3)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_recovery_target_time[3])))
	if base.B2i32(v106 == int32(0))|base.B2i32(v106 != v109) != 0 {
		v127 = v106
		v128 = v109
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v127-v128 == int32(0) {
		v267 = v15
		goto L1
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v112 = v10
	v113 = v103
	goto L32
L32:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+1)))
	if v117 == int32(0) {
		v127 = v117
		v128 = v116
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v127 = v117
	v128 = v116
	goto L30
L34:
	;
	v120 = int32(1)
	if v117 == v116 {
		v112 = v112 + v120
		v113 = v113 + v120
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v136 = v8 + int32(304)
	v138 = v8 + int32(192)
	v141 = F_ParseDateTime(m, v10, v8+int32(32), int32(153), v136, v138, v8+int32(404))
	mBase = m.M
	if v141 != 0 {
		v267 = v15
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v8)+404))
	v146 = v8 + int32(416)
	v150 = v8 + int32(412)
	v153 = F_DecodeDateTime(m, v136, v138, v142, v8+int32(408), v146, v8+int32(460), v150, v8+int32(24))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(0)
L39:
	;
	if v153 != 0 {
		v267 = v15
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v8)+408))
	if v157 != int32(2) {
		v267 = v15
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v8)+460))
	v161 = int32(16)
	v162 = v8 + v161
	v169 = m.G0
	v171 = v169 - v161
	m.G0 = v171
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v146)+20))
	if v173 <= int32(-4713) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	if v250 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L43:
	;
	m.G0 = v171 + int32(16)
	goto L42
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v162))) = int64(0)
	v250 = int32(-1)
	goto L43
L45:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	v191 = F_date2j(m, v173, v189, v190)
	mBase = m.M
	v194 = base.I64_extend_i32_s(v191 - int32(_a_F_check_recovery_target_time_4))
	v195 = int64(63)
	F___multi3(m, v171, v194, v194>>(uint(v195)%64), int64(86400000000), int64(0))
	mBase = m.M
	v200 = *(*int64)(unsafe.Add(mBase, uint32(v171)+8))
	v201 = *(*int64)(unsafe.Add(mBase, uint32(v171)))
	if v200 != v201>>(uint(v195)%64) {
		goto L44
	} else {
		goto L56
	}
L46:
	;
	if v173 != int32(-4713) {
		goto L44
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v173 <= int32(_a_F_check_recovery_target_time_5) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v146)+16))
	if int32(10) < v178 {
		v189 = v178
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L44
L51:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v146)+16))
	v189 = v183
	goto L45
L52:
	;
	goto L53
L53:
	;
	if v173 != int32(_a_F_check_recovery_target_time_6) {
		goto L44
	} else {
		goto L54
	}
L54:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v146)+16))
	if int32(5) < v186 {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	v189 = v186
	goto L45
L56:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v209 = int32(60)
	v218 = base.I64_extend_i32_s(v160) + base.I64_extend_i32_s(v206+(v207+v208*v209)*v209)*int64(1000000)
	v219 = v201 + v218
	*(*int64)(unsafe.Add(mBase, uint32(v162))) = v219
	if base.B2i32(v218 < int64(0))^base.B2i32(v219 < v201) != 0 {
		goto L44
	} else {
		goto L57
	}
L57:
	;
	if v150 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v231 = base.I64_extend_i32_s(int32(0)-v226)*int64(-1000000) + v219
	*(*int64)(unsafe.Add(mBase, uint32(v162))) = v231
	v233 = v231
	goto L60
L59:
	;
	v233 = v219
	goto L60
L60:
	;
	if base.Ui64(v233+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
		v250 = int32(0)
		goto L43
	} else {
		goto L61
	}
L61:
	;
	goto L44
L62:
	;
	v267 = int32(1)
	goto L1
L63:
	;
	goto L64
L64:
	;
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_time[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_time[5])) = v258
	goto L65
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
	v264 = F_format_elog_string(m, int32(_a_F_check_recovery_target_time_7), v8)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_recovery_target_time[6])) = v264
	v267 = v15
	goto L1
}
func F_recovery_create_dbdir(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	v4 = m.G0
	v6 = v4 - int32(160)
	m.G0 = v6
	v12 = F___fstatat(m, int32(-100), l0, v6-int32(-64), int32(0))
	mBase = m.M
	goto L4
L1:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L19
	} else {
		goto L61
	}
L2:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L19
	} else {
		goto L58
	}
L3:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L19
	} else {
		goto L55
	}
L4:
	;
	if v12 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l1 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	m.G0 = v6 + int32(160)
	return
L8:
	;
	v14 = F_strstr(m, l0, int32(_a_F_recovery_create_dbdir_0))
	mBase = m.M
	if v14 == int32(0) {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recovery_create_dbdir[0])))
	if v18 == int32(1) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_recovery_create_dbdir[1])))
	if v22&int32(1) == int32(0) {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v18 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v29 = int32(19)
	goto L18
L17:
	;
	v29 = int32(14)
	goto L18
L18:
	;
	v31 = F_errstart(m, v29, int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return
L20:
	;
	if v31 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	F_errmsg_internal(m, int32(_a_F_recovery_create_dbdir_1), v6+int32(32))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_recovery_create_dbdir[2]))
	v51 = m.G0
	v53 = v51 - int32(96)
	m.G0 = v53
	v56 = F_umask(m, int32(0))
	mBase = m.M
	v59 = F_umask(m, v56&int32(-193))
	mBase = m.M
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v66 = l0 + base.B2i32(v60 == int32(47))
	goto L28
L24:
	;
	F_errfinish(m, int32(_a_F_recovery_create_dbdir_2), int32(3298), int32(_a_F_recovery_create_dbdir_3))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if v116 != 0 {
		goto L1
	} else {
		goto L54
	}
L27:
	;
	v117 = F_umask(m, v56)
	mBase = m.M
	m.G0 = v53 + int32(96)
	goto L26
L28:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v71 != int32(47) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v116 = v100 >> (uint(int32(31)) % 32)
	goto L27
L30:
	;
	goto L29
L31:
	;
	v66 = v66 + int32(1)
	goto L28
L32:
	;
	v83 = F_stat(m, l0, v53)
	mBase = m.M
	if v83 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L33:
	;
	v80 = F_umask(m, v56)
	mBase = m.M
	v82 = int32(0)
	goto L32
L34:
	;
	if v71 != 0 {
		goto L31
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v76 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v76)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v79 != 0 {
		v82 = int32(1)
		goto L32
	} else {
		goto L38
	}
L37:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v74)
	goto L33
L38:
	;
	goto L33
L39:
	;
	v107 = int32(47)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v107)
	goto L31
L40:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v86&int32(_a_F_recovery_create_dbdir_4) != int32(_a_F_recovery_create_dbdir_5) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v82 != 0 {
		goto L50
	} else {
		goto L51
	}
L43:
	;
	if v82 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	if v82 != 0 {
		goto L39
	} else {
		goto L49
	}
L46:
	;
	v94 = int32(54)
	goto L48
L47:
	;
	v94 = int32(20)
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_recovery_create_dbdir[3])) = v94
	v116 = int32(-1)
	goto L27
L49:
	;
	v116 = int32(0)
	goto L27
L50:
	;
	v99 = int32(511)
	goto L52
L51:
	;
	v99 = v45
	goto L52
L52:
	;
	v100 = F_mkdir(m, l0, v99)
	mBase = m.M
	v101 = int32(0)
	if v82&base.B2i32(v101 <= v100) == v101 {
		goto L30
	} else {
		goto L53
	}
L53:
	;
	goto L39
L54:
	;
	goto L7
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg_internal(m, int32(_a_F_recovery_create_dbdir_6), v6)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_recovery_create_dbdir_2), int32(3291), int32(_a_F_recovery_create_dbdir_3))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L19
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l0
	F_errmsg(m, int32(_a_F_recovery_create_dbdir_7), v6+int32(48))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_recovery_create_dbdir_2), int32(3295), int32(_a_F_recovery_create_dbdir_3))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	F_errmsg(m, int32(_a_F_recovery_create_dbdir_8), v6+int32(16))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L19
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_recovery_create_dbdir_2), int32(3302), int32(_a_F_recovery_create_dbdir_3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L19
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
