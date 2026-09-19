package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_archiver_init_shmem_cb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-1)
	return
}
func F_pgstat_assoc_relation(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_assoc_relation[0]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+117)))
	if v10 != 0 {
		v11 = int32(0)
	} else {
		v11 = v8
	}
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = F_pgstat_prep_pending_entry(m, int32(2), v11, base.I64_extend_i32_u(v12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v10)
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = v12
		*(*int32)(unsafe.Add(mBase, uint32(l0)+272)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = l0
		return
	}
}
func F_pgstat_bestart_initial(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v87 int64
	_ = v87
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	v17 = m.G0
	v19 = v17 - int32(320)
	m.G0 = v19
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+188))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+179)) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+193))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+176)) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+201))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+204))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+171)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+212))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+216))
	v38 = v22 + int32(228)
	base.MemoryCopy(m, v19+int32(4), v38, int32(164))
	v42 = v22 + int32(392)
	v44 = v22 + int32(201)
	v46 = v22 + int32(193)
	v50 = v22 + int32(24)
	v52 = *(*int64)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[1]))
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[2]))
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[3]))
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[4]))
	if v58 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v72 = int32(_a_F_pgstat_bestart_initial_0)
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[5]))
	v75 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[5])) = v74 + v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v78 + v75
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v56
	v87 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+24)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v87
	base.MemoryCopy(m, v22+int32(56), v19+int32(184), int32(132))
	v99 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+192)) = uint8(v99)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v23
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v19)+179))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+3)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v104
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+200)) = uint8(v99)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v19)+171))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+3)) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v19)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v22)+220)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v75
	base.MemoryCopy(m, v38, v19+int32(4), int32(164))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = v87
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v87
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v99)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[4]))
	if v129 == v99 {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	base.MemoryFill(m, v19+int32(184), int32(0), int32(132))
	goto L1
L3:
	;
	goto L4
L4:
	;
	base.MemoryCopy(m, v19+int32(184), v58+int32(144), int32(132))
	goto L1
L5:
	;
	v258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v258)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+63)) = uint8(v258)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+63)) = uint8(v258)
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[6]))
	v267 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34+v265-v267))) = uint8(v258)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v271 + v267
	v275 = int32(_a_F_pgstat_bestart_initial_0)
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[5])) = v277 - v267
	m.G0 = v19 + int32(320)
	return
L6:
	;
	v255 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v255)
	goto L5
L7:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v129)+280))
	if v132 == int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L12
L9:
	;
	goto L5
L10:
	;
	v251 = F_strlen(m, v240)
	mBase = m.M
	goto L9
L12:
	;
	goto L13
L13:
	;
	v141 = int32(63)
	if (v23^v132)&int32(3) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v244 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v241))) = uint8(v244)
	goto L10
L15:
	;
	v225 = v220
	v226 = v221
	v227 = v222
	goto L36
L16:
	;
	if v215 == int32(0) {
		v240 = v213
		v241 = v214
		goto L14
	} else {
		goto L35
	}
L17:
	;
	v213 = v132
	v214 = v23
	v215 = v141
	goto L16
L18:
	;
	goto L19
L19:
	;
	v145 = int32(0)
	if base.B2i32(v132&int32(3) == v145)|int32(0) == v145 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v181 == int32(0) {
		v240 = v178
		v241 = v179
		goto L14
	} else {
		goto L29
	}
L21:
	;
	v157 = v132
	v158 = v23
	v159 = v141
	goto L24
L22:
	;
	goto L23
L23:
	;
	v178 = v132
	v179 = v23
	v180 = v141
	v181 = int32(1)
	goto L20
L24:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v161)
	if v161 == int32(0) {
		v220 = v157
		v221 = v158
		v222 = v159
		goto L15
	} else {
		goto L26
	}
L25:
	;
	v178 = v172
	v179 = v166
	v180 = v168
	v181 = v170
	goto L20
L26:
	;
	v165 = int32(1)
	v166 = v158 + v165
	v168 = v159 - v165
	v169 = int32(0)
	v170 = base.B2i32(v168 != v169)
	v172 = v157 + v165
	if v172&int32(3) == v169 {
		v178 = v172
		v179 = v166
		v180 = v168
		v181 = v170
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if v168 != 0 {
		v157 = v172
		v158 = v166
		v159 = v168
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	if base.B2i32(v184 == int32(0))|base.B2i32(base.Ui32(v180) < base.Ui32(int32(4))) != 0 {
		v213 = v178
		v214 = v179
		v215 = v180
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v191 = v178
	v192 = v179
	v193 = v180
	goto L31
L31:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	v199 = int32(-2139062144)
	if (int32(16843008)-v196|v196)&v199 != v199 {
		v220 = v191
		v221 = v192
		v222 = v193
		goto L15
	} else {
		goto L33
	}
L32:
	;
	v213 = v207
	v214 = v205
	v215 = v209
	goto L16
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v196
	v204 = int32(4)
	v205 = v192 + v204
	v207 = v191 + v204
	v209 = v193 - v204
	if base.Ui32(int32(3)) < base.Ui32(v209) {
		v191 = v207
		v192 = v205
		v193 = v209
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v220 = v213
	v221 = v214
	v222 = v215
	goto L15
L36:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225))))
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v229)
	if v229 == int32(0) {
		v240 = v225
		v241 = v226
		goto L14
	} else {
		goto L38
	}
L37:
	;
	v240 = v236
	v241 = v234
	goto L14
L38:
	;
	v233 = int32(1)
	v234 = v226 + v233
	v236 = v225 + v233
	v238 = v227 - v233
	if v238 != 0 {
		v225 = v236
		v226 = v234
		v227 = v238
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
}
func F_pgstat_clear_snapshot(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int64
	_ = v3
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v1 = int32(0)
	v3 = int64(0)
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[0])) = v3
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[1])) = v3
	*(*int64)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[2])) = v3
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[3])) = uint8(v1)
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[4])) = v1
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[5])) = v1
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[6]))
	if v21 != 0 {
		F_MemoryContextDelete(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[6])) = int32(0)
			F_pgstat_clear_backend_activity_snapshot(m)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				v30 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[7])) = uint8(v30)
				return
			}
		}
	} else {
		F_pgstat_clear_backend_activity_snapshot(m)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_clear_snapshot[7])) = uint8(v30)
			return
		}
	}
}
func F_pgstat_count_io_op(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	v4 = int32(0)
	v12 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_pgstat_count_io_op[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_pgstat_count_io_op[0]))) = v13 + int64(1)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_pgstat_count_io_op[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+uint32(_c_F_pgstat_count_io_op[1]))) = v17
	v19 = int32(1)
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_count_io_op[2]))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v23))|base.B2i32(v19<<(uint(v23)%32)&int32(_a_F_pgstat_count_io_op_0) == v4) == v4 {
		v42 = l0*int32(320) + l1<<(uint(int32(6))%32) + l2<<(uint(int32(3))%32)
		v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_pgstat_count_io_op[3])))
		*(*int64)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_pgstat_count_io_op[3]))) = v43 + base.I64_extend_i32_u(v19)
		v47 = *(*int64)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_pgstat_count_io_op[4])))
		*(*int64)(unsafe.Add(mBase, uint32(v42)+uint32(_c_F_pgstat_count_io_op[4]))) = v47 + int64(0)
		v51 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op[5])) = uint8(v51)
		*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op[6])) = uint8(v51)
	} else {
	}
	v58 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op[5])) = uint8(v58)
	*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_count_io_op[7])) = uint8(v58)
	return
}
func F_pgstat_database_reset_timestamp_cb(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+280)) = l1
	return
}
func F_pgstat_drop_entry_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v14 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
	v21 = base.AtomicRmwSub32(m, l0, int32(20), v17)
	if v21 == v17 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L12
	} else {
		goto L24
	}
L4:
	;
	m.G0 = v12 + int32(32)
	return base.B2i32(v21 == int32(1))
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	if l1 != 0 {
		goto L4
	} else {
		goto L22
	}
L8:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry_internal[0]))
	F_dsa_free(m, v88, v24)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L12
	} else {
		goto L21
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry_internal[1]))
	F_dshash_delete_entry(m, v28, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	v45 = v34 + int32(base.Ui32(v36)>>(uint(int32(32)-v38)%32))<<(uint(int32(2))%32)
	goto L15
L12:
	;
	return int32(0)
L13:
	;
	goto L8
L14:
	;
	goto L8
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v53 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	F_dsa_free(m, v61, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L12
	} else {
		goto L20
	}
L17:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v57 = F_dsa_get_address(m, v56, v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if v57 != v35 {
		v45 = v57
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v45))) = v60
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v33)+32))
	v71 = v66 + int32(base.Ui32(v36)>>(uint(int32(25))%32))*int32(20)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+24)) = v72 - int32(1)
	goto L14
L21:
	;
	goto L4
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry_internal[1]))
	F_dshash_release_lock(m, v92, l0)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L4
L24:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v114-int32(1)) <= base.Ui32(int32(11)) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v144 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)+68))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v145
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v144
	F_errmsg_internal(m, int32(_a_F_pgstat_drop_entry_internal_0), v12)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L12
	} else {
		goto L32
	}
L26:
	;
	v143 = v114*int32(72) + int32(_a_F_pgstat_drop_entry_internal_1)
	goto L25
L27:
	;
	goto L28
L28:
	;
	if base.Ui32(int32(8)) < base.Ui32(v114-int32(24)) {
		v141 = int32(0)
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v143 = v141
	goto L25
L30:
	;
	v129 = int32(0)
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_drop_entry_internal[2]))
	if v131 == v129 {
		v141 = v129
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v131+v114<<(uint(int32(2))%32)-int32(96))))
	v141 = v139
	goto L29
L32:
	;
	F_errfinish(m, int32(_a_F_pgstat_drop_entry_internal_2), int32(908), int32(_a_F_pgstat_drop_entry_internal_3))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L12
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_fetch_stat_backend_by_pid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	v3 = int32(0)
	v7 = F_BackendPidGetProc(m, l0)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L5
L4:
	;
	goto L5
L5:
	;
	if v7 != 0 {
		v51 = v7
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_stat_backend_by_pid[0]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v57 = base.I32_div_s(v51-v54, int32(640))
	v58 = F_pgstat_get_beentry_by_proc_number(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L21
	}
L7:
	;
	v13 = int32(0)
	if l0 == v13 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if v48 != 0 {
		v51 = v48
		goto L6
	} else {
		goto L19
	}
L9:
	;
	v48 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_fetch_stat_backend_by_pid[1]))
	v25 = v13
	goto L14
L12:
	;
	v48 = v42
	goto L8
L13:
	;
	v42 = v32 + int32(640)
	goto L12
L14:
	;
	v28 = v25 * int32(640)
	v29 = v21 + v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+44))
	if v30 == l0 {
		v42 = v29
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v48 = int32(0)
	goto L8
L16:
	;
	v32 = v21 + v28
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+684))
	if v33 == l0 {
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v36 = v25 + int32(2)
	if v36 != int32(38) {
		v25 = v36
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	return int32(0)
L20:
	;
	return v87
L21:
	;
	if v58 == int32(0) {
		v87 = v3
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v62))|base.B2i32(int32(1)<<(uint(v62)%32)&int32(_a_F_pgstat_fetch_stat_backend_by_pid_0) == int32(0)) != 0 {
		v87 = v3
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v72 != l0 {
		v87 = v3
		goto L20
	} else {
		goto L24
	}
L24:
	;
	if l1 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v62
	goto L27
L26:
	;
	goto L27
L27:
	;
	v78 = F_pgstat_fetch_entry(m, int32(6), int32(0), base.I64_extend_i32_s(v57))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v78|base.B2i32(l1 == int32(0)) != 0 {
		v87 = v78
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83
	v87 = v83
	goto L20
}
func F_pgstat_get_kind_info(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(11)) {
		return l0*int32(72) + int32(_a_F_pgstat_get_kind_info_0)
	} else {
		if base.Ui32(int32(8)) < base.Ui32(l0-int32(24)) {
			v29 = int32(0)
		} else {
			v17 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_kind_info[0]))
			if v19 == v17 {
				v29 = v17
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v19+l0<<(uint(int32(2))%32)-int32(96))))
				v29 = v27
			}
		}
		return v29
	}
}
func F_pgstat_get_my_query_id(m *base.Module) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_get_my_query_id[0]))
	if v3 == int32(0) {
		return int64(0)
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(v3)+392))
		return v8
	}
}
func F_pgstat_index(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v12
	v24 = int32(1)
	v26 = F_GetAccessStrategy(m, v24)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_LockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v34 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if base.Ui32(int32(1)) < base.Ui32(v34) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v45 = v34
	v46 = v24
	goto L9
L7:
	;
	v71 = v34
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = base.I64_extend_i32_u(v71) << (uint(int64(13)) % 64)
	F_relation_close(m, l0, int32(1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L21
	}
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_index[0]))
	if v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v71 = v61
	goto L8
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	m.T0[l1].(func(*base.Module, int32, int32, int32, int32))(m, v10, l0, v46, v26)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v55 = v46 + int32(1)
	if base.Ui32(v55) < base.Ui32(v45) {
		v46 = v55
		goto L9
	} else {
		goto L16
	}
L16:
	;
	F_LockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v61 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v55) < base.Ui32(v61) {
		v45 = v61
		v46 = v55
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	v81 = F_build_pgstattuple_type(m, v10, l2)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v10 + int32(48)
	return v81
}
func F_pgstat_io_init_shmem_cb(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	v2 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-1)
	v9 = l0 + int32(16)
	v10 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v9))) = uint16(v10)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(-1)
	v17 = l0 + int32(32)
	v18 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v17))) = uint16(v18)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(-1)
	v25 = l0 + int32(48)
	v26 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v26)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = int64(-1)
	v33 = l0 - int32(-64)
	v34 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v33))) = uint16(v34)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = int64(-1)
	v41 = l0 + int32(80)
	v42 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v41))) = uint16(v42)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+8)) = int64(-1)
	v49 = l0 + int32(96)
	v50 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v50)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = int64(-1)
	v57 = l0 + int32(112)
	v58 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v57))) = uint16(v58)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v57)+8)) = int64(-1)
	v65 = l0 + int32(128)
	v66 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v65))) = uint16(v66)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = int64(-1)
	v73 = l0 + int32(144)
	v74 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v73))) = uint16(v74)
	*(*int32)(unsafe.Add(mBase, uint32(v73)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = int64(-1)
	v81 = l0 + int32(160)
	v82 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v81))) = uint16(v82)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = int64(-1)
	v89 = l0 + int32(176)
	v90 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v90)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = int64(-1)
	v97 = l0 + int32(192)
	v98 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v97))) = uint16(v98)
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v97)+8)) = int64(-1)
	v105 = l0 + int32(208)
	v106 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v105))) = uint16(v106)
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v105)+8)) = int64(-1)
	v113 = l0 + int32(224)
	v114 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v113))) = uint16(v114)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v113)+8)) = int64(-1)
	v121 = l0 + int32(240)
	v122 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v122)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v121)+8)) = int64(-1)
	v129 = l0 + int32(256)
	v130 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v129))) = uint16(v130)
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v129)+8)) = int64(-1)
	v137 = l0 + int32(272)
	v138 = int32(81)
	*(*uint16)(unsafe.Add(mBase, uint32(v137))) = uint16(v138)
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = int64(-1)
	return
}
func F_pgstat_progress_incr_param(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_incr_param[0]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_progress_incr_param[1])))
		if v9&int32(1) == int32(0) {
		} else {
			v14 = int32(_a_F_pgstat_progress_incr_param_0)
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_incr_param[2]))
			v17 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_incr_param[2])) = v16 + v17
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v20 + v17
			v28 = v5 + l0<<(uint(int32(3))%32) + int32(232)
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
			*(*int64)(unsafe.Add(mBase, uint32(v28))) = v29 + l1
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v32 + v17
			v38 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_incr_param[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_incr_param[2])) = v38 - v17
		}
	}
	return
}
func F_pgstat_relation(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v74 int64
	_ = v74
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v120 int64
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int64
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v197 int32
	_ = v197
	var v201 int64
	_ = v201
	var v213 int32
	_ = v213
	var v217 int64
	_ = v217
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int64
	_ = v263
	var v265 int32
	_ = v265
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	v10 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(160)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+118)))
	if v21 == int32(116) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L16
	} else {
		goto L118
	}
L2:
	;
	m.G0 = v18 + int32(160)
	return v424
L3:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354)+18)))
	if v355 == int32(0) {
		goto L1
	} else {
		goto L95
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L16
	} else {
		goto L90
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L16
	} else {
		goto L86
	}
L6:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v24 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+119)))
	switch v27 - int32(83) {
	case 0, 26, 31, 33:
		goto L10
	default:
		goto L4
	case 22:
		goto L3
	}
L9:
	;
	goto L8
L10:
	;
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v30
	if v27 != int32(83) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L16
	} else {
		goto L82
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	if v42 != int32(2) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v45 = int32(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v53 = m.T0[v52].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, int32(_a_F_pgstat_relation_0), v45, v45, v45, int32(321))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	return int32(0)
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = int32(4)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+36))
	v60 = F_heap_getnext(m, v53)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v66 = v60
	v67 = v45
	v71 = v10
	v73 = v10
	v74 = v10
	v75 = v10
	v76 = v10
	goto L22
L20:
	;
	v197 = v45
	v201 = v10
	goto L21
L21:
	;
	if base.Ui32(v197) < base.Ui32(v59) {
		goto L58
	} else {
		goto L59
	}
L22:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[0]))
	if v78 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v100
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v98
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v179
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v97
	v197 = v175
	v201 = v179
	goto L21
L24:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
	F_LockBuffer(m, v81, int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L16
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
	v88 = F_HeapTupleSatisfiesVisibility(m, v66, v18+int32(40), v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v90 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v66))))
	if v88 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v53)+56))
	F_LockBuffer(m, v101, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L16
	} else {
		goto L34
	}
L31:
	;
	v97 = v73 + int64(1)
	v98 = v90 + v74
	v99 = v75
	v100 = v76
	goto L30
L32:
	;
	goto L33
L33:
	;
	v97 = v73
	v98 = v74
	v99 = v75 + int64(1)
	v100 = v90 + v76
	goto L30
L34:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+6)))
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)))
	v109 = v105 | v106<<(uint(int32(16))%32)
	if base.Ui32(v67) <= base.Ui32(v109) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v116 = v67
	v120 = v71
	goto L38
L36:
	;
	v175 = v67
	v179 = v71
	goto L37
L37:
	;
	v185 = F_heap_getnext(m, v53)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L16
	} else {
		goto L56
	}
L38:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[0]))
	if v127 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v175 = v168
	v179 = v166
	goto L37
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L16
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v130 = int32(0)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v53)+60))
	v133 = F_ReadBufferExtended(m, l0, v130, v116, v130, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L16
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	F_LockBuffer(m, v133, int32(1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	if v133 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+14)))
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+12)))
	v158 = v156 - v157
	v159 = int32(0)
	if v159 < v158 {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[1]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141+(v133^int32(-1))<<(uint(int32(2))%32))))
	v155 = v147
	goto L46
L48:
	;
	goto L49
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[2]))
	v155 = v149 + v133<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	F_UnlockReleaseBuffer(m, v133)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L16
	} else {
		goto L54
	}
L51:
	;
	v162 = v158
	goto L53
L52:
	;
	v162 = v159
	goto L53
L53:
	;
	goto L50
L54:
	;
	v166 = v120 + base.I64_extend_i32_u(v162)
	v168 = v116 + int32(1)
	if base.Ui32(v168) <= base.Ui32(v109) {
		v116 = v168
		v120 = v166
		goto L38
	} else {
		goto L55
	}
L55:
	;
	goto L39
L56:
	;
	if v185 != 0 {
		v66 = v185
		v67 = v175
		v71 = v179
		v73 = v97
		v74 = v98
		v75 = v99
		v76 = v100
		goto L22
	} else {
		goto L57
	}
L57:
	;
	goto L23
L58:
	;
	v213 = v197
	v217 = v201
	goto L61
L59:
	;
	goto L60
L60:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+188))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	m.T0[v285].(func(*base.Module, int32))(m, v53)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L16
	} else {
		goto L79
	}
L61:
	;
	v224 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[0]))
	if v224 != 0 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v263
	goto L60
L63:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v227 = int32(0)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v53)+60))
	v230 = F_ReadBufferExtended(m, l0, v227, v213, v227, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L16
	} else {
		goto L67
	}
L66:
	;
	goto L65
L67:
	;
	F_LockBuffer(m, v230, int32(1))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L16
	} else {
		goto L68
	}
L68:
	;
	if v230 < int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+14)))
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v252)+12)))
	v255 = v253 - v254
	v256 = int32(0)
	if v256 < v255 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[1]))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v238+(v230^int32(-1))<<(uint(int32(2))%32))))
	v252 = v244
	goto L69
L71:
	;
	goto L72
L72:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[2]))
	v252 = v246 + v230<<(uint(int32(13))%32) + int32(-8192)
	goto L69
L73:
	;
	F_UnlockReleaseBuffer(m, v230)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L16
	} else {
		goto L77
	}
L74:
	;
	v259 = v255
	goto L76
L75:
	;
	v259 = v256
	goto L76
L76:
	;
	goto L73
L77:
	;
	v263 = v217 + base.I64_extend_i32_u(v259)
	v265 = v213 + int32(1)
	if v265 != v59 {
		v213 = v265
		v217 = v263
		goto L61
	} else {
		goto L78
	}
L78:
	;
	goto L62
L79:
	;
	F_relation_close(m, l0, int32(1))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = base.I64_extend_i32_u(v59) << (uint(int64(13)) % 64)
	v297 = F_build_pgstattuple_type(m, v18+int32(112), l1)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	v424 = v297
	goto L2
L82:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L16
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(_a_F_pgstat_relation_1), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L16
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_2), int32(335), int32(_a_F_pgstat_relation_3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L16
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L16
	} else {
		goto L87
	}
L87:
	;
	F_errmsg(m, int32(_a_F_pgstat_relation_4), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L16
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_2), int32(255), int32(_a_F_pgstat_relation_5))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L16
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L90:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L16
	} else {
		goto L91
	}
L91:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v338 + int32(4)
	F_errmsg(m, int32(_a_F_pgstat_relation_6), v18)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L16
	} else {
		goto L92
	}
L92:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v346 = int32(*(*int8)(unsafe.Add(mBase, uint32(v345)+119)))
	F_errdetail_relkind_not_supported(m, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L16
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_2), int32(306), int32(_a_F_pgstat_relation_5))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L16
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	if v358 <= int32(782) {
		goto L101
	} else {
		goto L102
	}
L96:
	;
	v407 = F_pgstat_index(m, l0, int32(_a_F_pgstat_relation_7), l1)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L16
	} else {
		goto L117
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L16
	} else {
		goto L113
	}
L98:
	;
	v383 = int32(_a_F_pgstat_relation_8)
	goto L97
L99:
	;
	v383 = int32(_a_F_pgstat_relation_9)
	goto L97
L100:
	;
	v379 = F_pgstat_index(m, l0, int32(_a_F_pgstat_relation_10), l1)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L16
	} else {
		goto L112
	}
L101:
	;
	switch v358 - int32(403) {
	case 0:
		goto L96
	default:
		goto L98
	case 2:
		goto L100
	}
L102:
	;
	goto L103
L103:
	;
	if v358 <= int32(2741) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	if v358 != int32(783) {
		goto L98
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	if v358 == int32(2742) {
		v383 = int32(_a_F_pgstat_relation_11)
		goto L97
	} else {
		goto L109
	}
L107:
	;
	v368 = F_pgstat_index(m, l0, int32(_a_F_pgstat_relation_12), l1)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L16
	} else {
		goto L108
	}
L108:
	;
	v424 = v368
	goto L2
L109:
	;
	if v358 == int32(3580) {
		goto L99
	} else {
		goto L110
	}
L110:
	;
	if v358 != int32(4000) {
		goto L98
	} else {
		goto L111
	}
L111:
	;
	v383 = int32(_a_F_pgstat_relation_13)
	goto L97
L112:
	;
	v424 = v379
	goto L2
L113:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L16
	} else {
		goto L114
	}
L114:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v391 + int32(4)
	F_errmsg(m, int32(_a_F_pgstat_relation_14), v18+int32(16))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L16
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_2), int32(298), int32(_a_F_pgstat_relation_5))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L16
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v424 = v407
	goto L2
L118:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L16
	} else {
		goto L119
	}
L119:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v436 + int32(4)
	F_errmsg(m, int32(_a_F_pgstat_relation_15), v18+int32(32))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L16
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_2), int32(269), int32(_a_F_pgstat_relation_5))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L16
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_release_entry_ref(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v105 int64
	_ = v105
	var v111 int64
	_ = v111
	var v117 int64
	_ = v117
	var v122 int64
	_ = v122
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var __phi169 int32
	_ = __phi169
	var v174 int32
	_ = v174
	var __phi174 int32
	_ = __phi174
	var v177 int32
	_ = v177
	var __phi177 int32
	_ = __phi177
	var v178 int32
	_ = v178
	var __phi178 int32
	_ = __phi178
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v187 int64
	_ = v187
	var v192 int64
	_ = v192
	var v198 int64
	_ = v198
	var v203 int64
	_ = v203
	var v208 int64
	_ = v208
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L16
	} else {
		goto L55
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L16
	} else {
		goto L52
	}
L3:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[0]))
	v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v91
	v93 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v93
	v95 = int64(23)
	v98 = int64(2388976653695081527)
	v99 = (v93 ^ int64(base.Ui64(v93)>>(uint(v95)%64))) * v98
	v100 = int64(47)
	v105 = int64(-8645972361240307355)
	v111 = (v91 ^ int64(base.Ui64(v91)>>(uint(v95)%64))) * v98
	v117 = ((v99^int64(base.Ui64(v99)>>(uint(v100)%64))^int64(-9208349263878056368))*v105 ^ int64(base.Ui64(v111)>>(uint(v100)%64)) ^ v111) * v105
	v122 = (int64(base.Ui64(v117)>>(uint(v95)%64)) ^ v117) * v98
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v132 = base.I32_wrap_i64(int64(base.Ui64(v122)>>(uint(v100)%64)) ^ v122 - int64(base.Ui64(v122)>>(uint(int64(32))%64)))
	goto L29
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v53 == int32(0) {
		goto L3
	} else {
		goto L19
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if base.Ui32(v22-int32(1)) <= base.Ui32(int32(11)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v39 = v22*int32(72) + int32(_a_F_pgstat_release_entry_ref_0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[1]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+v22<<(uint(int32(2))%32)-int32(96))))
	v39 = v38
	goto L9
L13:
	;
	m.T0[v40].(func(*base.Module, int32))(m, l1)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	F_pfree(m, v20)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L16
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v50
	goto L7
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v57 = int32(1)
	v59 = base.AtomicRmwSub32(m, v56, int32(20), v57)
	if v59 != v57 {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[2]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v66 = F_dshash_find(m, v63, v64, int32(1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	if v66 == int32(0) {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+24))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v71 == v72 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[2]))
	F_dshash_delete_entry(m, v76, v66)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[2]))
	F_dshash_release_lock(m, v84, v66)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L28
	}
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[3]))
	F_dsa_free(m, v80, v74)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L3
L28:
	;
	goto L3
L29:
	;
	v142 = v132 & v131
	v145 = v130 + v142*int32(24)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+16)))
	switch v146 {
	case 0:
		goto L32
	case 1:
		goto L33
	default:
		goto L31
	}
L31:
	;
	v132 = v142 + int32(1)
	goto L29
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L16
	} else {
		goto L49
	}
L33:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v145)+8))
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	if v147^v148|(v150^v151) != int64(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v90)+8))
	v157 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v90)+8)) = v156 - v157
	v162 = (v142 + v157) & v131
	v165 = v130 + v162*int32(24)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+16)))
	if v166 != v157 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v245 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v237)+16)) = uint8(v245)
	if l1 != 0 {
		goto L45
	} else {
		goto L46
	}
L36:
	;
	v237 = v145
	goto L35
L37:
	;
	goto L38
L38:
	;
	__phi169 = v145
	__phi174 = v165
	__phi177 = v131
	__phi178 = v162
	v169 = __phi169
	v174 = __phi174
	v177 = __phi177
	v178 = __phi178
	goto L39
L39:
	;
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v174)+8))
	v180 = int64(23)
	v183 = int64(2388976653695081527)
	v184 = (int64(base.Ui64(v179)>>(uint(v180)%64)) ^ v179) * v183
	v185 = int64(47)
	v187 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
	v192 = (int64(base.Ui64(v187)>>(uint(v180)%64)) ^ v187) * v183
	v198 = int64(-8645972361240307355)
	v203 = (int64(base.Ui64(v184)>>(uint(v185)%64)) ^ (v192^int64(base.Ui64(v192)>>(uint(v185)%64))^int64(-9208349263878056368))*v198 ^ v184) * v198
	v208 = (int64(base.Ui64(v203)>>(uint(v180)%64)) ^ v203) * v183
	if v178 == v177&base.I32_wrap_i64(int64(base.Ui64(v208)>>(uint(v185)%64))^v208-int64(base.Ui64(v208)>>(uint(int64(32))%64))) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v237 = v174
	goto L35
L41:
	;
	v237 = v169
	goto L35
L42:
	;
	goto L43
L43:
	;
	v218 = *(*int64)(unsafe.Add(mBase, uint32(v174)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v169)+16)) = v218
	v220 = *(*int64)(unsafe.Add(mBase, uint32(v174)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v169)+8)) = v220
	v222 = *(*int64)(unsafe.Add(mBase, uint32(v174)))
	*(*int64)(unsafe.Add(mBase, uint32(v169))) = v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v90)+20))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v226 = int32(1)
	v228 = v225 & (v178 + v226)
	v231 = v224 + v228*int32(24)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+16)))
	if v232 == v226 {
		__phi169 = v174
		__phi174 = v231
		__phi177 = v225
		__phi178 = v228
		v169 = __phi169
		v174 = __phi174
		v177 = __phi177
		v178 = __phi178
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L40
L45:
	;
	F_pfree(m, l1)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L16
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	m.G0 = v13 + int32(16)
	return
L48:
	;
	goto L47
L49:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_release_entry_ref_1), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L16
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_pgstat_release_entry_ref_2), int32(667), int32(_a_F_pgstat_release_entry_ref_3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L16
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_release_entry_ref_4), int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L16
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_pgstat_release_entry_ref_2), int32(640), int32(_a_F_pgstat_release_entry_ref_3))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L16
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_release_entry_ref_5), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_pgstat_release_entry_ref_2), int32(611), int32(_a_F_pgstat_release_entry_ref_3))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L16
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_replslot_reset_timestamp_cb(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = l1
	return
}
func F_pgstat_replslot_to_serialized_name_cb(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_replslot_to_serialized_name_cb[0]))
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_replslot_to_serialized_name_cb[1]))
	v18 = F_LWLockAcquire(m, v14+int32(_a_F_pgstat_replslot_to_serialized_name_cb_0), int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		v22 = v12 + v10*int32(288)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)))
		if v23 == int32(1) {
			v29 = F_strncpy(m, l2, v22+int32(24), int32(64))
			mBase = m.M
			v30 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v29)+63)) = uint8(v30)
		} else {
		}
		v33 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_replslot_to_serialized_name_cb[1]))
		F_LWLockRelease(m, v33+int32(_a_F_pgstat_replslot_to_serialized_name_cb_0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			if v23 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v8))) = v44
					F_errmsg_internal(m, int32(_a_F_pgstat_replslot_to_serialized_name_cb_1), v8)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pgstat_replslot_to_serialized_name_cb_2), int32(198), int32(_a_F_pgstat_replslot_to_serialized_name_cb_3))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_pgstat_report_plan_id(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_plan_id[0]))
	if v5 == int32(0) {
	} else {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_plan_id[1])))
		if v9&int32(1) == int32(0) {
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v5)+400))
			if base.B2i32(l1 == int32(0))&base.B2i32(v16 != int64(0)) != 0 {
			} else {
				v20 = int32(_a_F_pgstat_report_plan_id_0)
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_plan_id[2]))
				v23 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_plan_id[2])) = v22 + v23
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26 + v23
				*(*int64)(unsafe.Add(mBase, uint32(v5)+400)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26 + int32(2)
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_plan_id[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_plan_id[2])) = v37 - v23
			}
		}
	}
	return
}
func F_pgstat_report_xact_timestamp(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_report_xact_timestamp[0])))
	if v4 != int32(1) {
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_xact_timestamp[1]))
		if v8 == int32(0) {
		} else {
			v11 = int32(_a_F_pgstat_report_xact_timestamp_0)
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_xact_timestamp[2]))
			v14 = int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_xact_timestamp[2])) = v13 + v14
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17 + v14
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v17 + int32(2)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_xact_timestamp[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_xact_timestamp[2])) = v28 - v14
		}
	}
	return
}
func F_pgstat_request_entry_refs_gc(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_request_entry_refs_gc[0]))
	v5 = base.AtomicRmwAdd64(m, v2, int32(16), int64(1))
	return
}
func F_pgstat_reset_wait_event_storage(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_reset_wait_event_storage[0])) = int32(_a_F_pgstat_reset_wait_event_storage_0)
	return
}
func F_pgstat_slru_reset_all_cb(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int64
	_ = v179
	var v195 int32
	_ = v195
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_reset_all_cb[0]))
	v7 = v5 + int32(_a_F_pgstat_slru_reset_all_cb_0)
	v9 = F_LWLockAcquire(m, v7, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_slru_reset_all_cb[1]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_slru_reset_all_cb[2]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_slru_reset_all_cb[3]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_slru_reset_all_cb[4]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_slru_reset_all_cb[5]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_slru_reset_all_cb[6]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_slru_reset_all_cb[7]))) = v11
		*(*int64)(unsafe.Add(mBase, uint32(v5)+uint32(_c_F_pgstat_slru_reset_all_cb[8]))) = l0
		F_LWLockRelease(m, v7)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_reset_all_cb[0]))
			v31 = v29 + int32(_a_F_pgstat_slru_reset_all_cb_0)
			v33 = F_LWLockAcquire(m, v31, int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v35 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pgstat_slru_reset_all_cb[9]))) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pgstat_slru_reset_all_cb[10]))) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pgstat_slru_reset_all_cb[11]))) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pgstat_slru_reset_all_cb[12]))) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pgstat_slru_reset_all_cb[13]))) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pgstat_slru_reset_all_cb[14]))) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pgstat_slru_reset_all_cb[15]))) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_pgstat_slru_reset_all_cb[16]))) = l0
				F_LWLockRelease(m, v31)
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_reset_all_cb[0]))
					v55 = v53 + int32(_a_F_pgstat_slru_reset_all_cb_0)
					v57 = F_LWLockAcquire(m, v55, int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						v59 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_pgstat_slru_reset_all_cb[17]))) = v59
						*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_pgstat_slru_reset_all_cb[18]))) = v59
						*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_pgstat_slru_reset_all_cb[19]))) = v59
						*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_pgstat_slru_reset_all_cb[20]))) = v59
						*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_pgstat_slru_reset_all_cb[21]))) = v59
						*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_pgstat_slru_reset_all_cb[22]))) = v59
						*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_pgstat_slru_reset_all_cb[23]))) = v59
						*(*int64)(unsafe.Add(mBase, uint32(v53)+uint32(_c_F_pgstat_slru_reset_all_cb[24]))) = l0
						F_LWLockRelease(m, v55)
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_reset_all_cb[0]))
							v79 = v77 + int32(_a_F_pgstat_slru_reset_all_cb_0)
							v81 = F_LWLockAcquire(m, v79, int32(0))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return
							} else {
								v83 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pgstat_slru_reset_all_cb[25]))) = v83
								*(*int64)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pgstat_slru_reset_all_cb[26]))) = v83
								*(*int64)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pgstat_slru_reset_all_cb[27]))) = v83
								*(*int64)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pgstat_slru_reset_all_cb[28]))) = v83
								*(*int64)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pgstat_slru_reset_all_cb[29]))) = v83
								*(*int64)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pgstat_slru_reset_all_cb[30]))) = v83
								*(*int64)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pgstat_slru_reset_all_cb[31]))) = v83
								*(*int64)(unsafe.Add(mBase, uint32(v77)+uint32(_c_F_pgstat_slru_reset_all_cb[32]))) = l0
								F_LWLockRelease(m, v79)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									v101 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_reset_all_cb[0]))
									v103 = v101 + int32(_a_F_pgstat_slru_reset_all_cb_0)
									v105 = F_LWLockAcquire(m, v103, int32(0))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return
									} else {
										v107 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pgstat_slru_reset_all_cb[33]))) = v107
										*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pgstat_slru_reset_all_cb[34]))) = v107
										*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pgstat_slru_reset_all_cb[35]))) = v107
										*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pgstat_slru_reset_all_cb[36]))) = v107
										*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pgstat_slru_reset_all_cb[37]))) = v107
										*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pgstat_slru_reset_all_cb[38]))) = v107
										*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pgstat_slru_reset_all_cb[39]))) = v107
										*(*int64)(unsafe.Add(mBase, uint32(v101)+uint32(_c_F_pgstat_slru_reset_all_cb[40]))) = l0
										F_LWLockRelease(m, v103)
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return
										} else {
											v125 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_reset_all_cb[0]))
											v127 = v125 + int32(_a_F_pgstat_slru_reset_all_cb_0)
											v129 = F_LWLockAcquire(m, v127, int32(0))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return
											} else {
												v131 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_pgstat_slru_reset_all_cb[41]))) = v131
												*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_pgstat_slru_reset_all_cb[42]))) = v131
												*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_pgstat_slru_reset_all_cb[43]))) = v131
												*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_pgstat_slru_reset_all_cb[44]))) = v131
												*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_pgstat_slru_reset_all_cb[45]))) = v131
												*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_pgstat_slru_reset_all_cb[46]))) = v131
												*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_pgstat_slru_reset_all_cb[47]))) = v131
												*(*int64)(unsafe.Add(mBase, uint32(v125)+uint32(_c_F_pgstat_slru_reset_all_cb[48]))) = l0
												F_LWLockRelease(m, v127)
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return
												} else {
													v149 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_reset_all_cb[0]))
													v151 = v149 + int32(_a_F_pgstat_slru_reset_all_cb_0)
													v153 = F_LWLockAcquire(m, v151, int32(0))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														v155 = int64(0)
														*(*int64)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_pgstat_slru_reset_all_cb[49]))) = v155
														*(*int64)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_pgstat_slru_reset_all_cb[50]))) = v155
														*(*int64)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_pgstat_slru_reset_all_cb[51]))) = v155
														*(*int64)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_pgstat_slru_reset_all_cb[52]))) = v155
														*(*int64)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_pgstat_slru_reset_all_cb[53]))) = v155
														*(*int64)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_pgstat_slru_reset_all_cb[54]))) = v155
														*(*int64)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_pgstat_slru_reset_all_cb[55]))) = v155
														*(*int64)(unsafe.Add(mBase, uint32(v149)+uint32(_c_F_pgstat_slru_reset_all_cb[56]))) = l0
														F_LWLockRelease(m, v151)
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return
														} else {
															v173 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_slru_reset_all_cb[0]))
															v175 = v173 + int32(_a_F_pgstat_slru_reset_all_cb_0)
															v177 = F_LWLockAcquire(m, v175, int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return
															} else {
																v179 = int64(0)
																*(*int64)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_pgstat_slru_reset_all_cb[57]))) = v179
																*(*int64)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_pgstat_slru_reset_all_cb[58]))) = v179
																*(*int64)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_pgstat_slru_reset_all_cb[59]))) = v179
																*(*int64)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_pgstat_slru_reset_all_cb[60]))) = v179
																*(*int64)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_pgstat_slru_reset_all_cb[61]))) = v179
																*(*int64)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_pgstat_slru_reset_all_cb[62]))) = v179
																*(*int64)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_pgstat_slru_reset_all_cb[63]))) = v179
																*(*int64)(unsafe.Add(mBase, uint32(v173)+uint32(_c_F_pgstat_slru_reset_all_cb[64]))) = l0
																F_LWLockRelease(m, v175)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return
																} else {
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
}
func F_pgstat_subscription_reset_timestamp_cb(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = l1
	return
}
func F_pgstat_twophase_postabort(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_twophase_postabort[0]))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+52)))
	if v10 != 0 {
		v11 = int32(0)
	} else {
		v11 = v9
	}
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v15 = F_pgstat_prep_pending_entry(m, int32(2), v11, base.I64_extend_i32_u(v12), int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		*(*uint8)(unsafe.Add(mBase, uint32(v17)+4)) = uint8(v10)
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = v12
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+53)))
		if v20 == int32(0) {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			v30 = v23
		} else {
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l2)+24))
			*(*int64)(unsafe.Add(mBase, uint32(l2))) = v24
			v26 = *(*int64)(unsafe.Add(mBase, uint32(l2)+32))
			*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v26
			v28 = *(*int64)(unsafe.Add(mBase, uint32(l2)+40))
			*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v28
			v30 = v24
		}
		v31 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v31 + v30
		v34 = *(*int64)(unsafe.Add(mBase, uint32(v17)+48))
		v35 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v34 + v35
		v38 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
		v39 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v38 + v39
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v17)+96))
		v43 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		v44 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = v42 + (v43 + v44)
		return
	}
}
