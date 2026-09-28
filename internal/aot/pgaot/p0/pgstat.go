package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgstat_archiver_init_shmem_cb(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_LWLockInitialize(m, l0, int32(85))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
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
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int64
	_ = v91
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
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
	v82 = int32(0)
	v85 = base.AtomicRmwOr32(m, v82, int32(_a_F_pgstat_bestart_initial_1), v82)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v56
	v91 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v50)+24)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v50)+16)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v91
	base.MemoryCopy(m, v22+int32(56), v19+int32(184), int32(132))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+192)) = uint8(v82)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v23
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v19)+179))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+3)) = v106
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v19)+176))
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v108
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+200)) = uint8(v82)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v19)+171))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+3)) = v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v19)+168))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v114
	*(*int64)(unsafe.Add(mBase, uint32(v22)+220)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v75
	base.MemoryCopy(m, v38, v19+int32(4), int32(164))
	*(*int64)(unsafe.Add(mBase, uint32(v42)+8)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v42))) = v91
	*(*uint8)(unsafe.Add(mBase, uint32(v33))) = uint8(v82)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[4]))
	if v133 == v82 {
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
	v262 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v262)
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+63)) = uint8(v262)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+63)) = uint8(v262)
	v269 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[6]))
	v271 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34+v269-v271))) = uint8(v262)
	v278 = base.AtomicRmwOr32(m, v262, int32(_a_F_pgstat_bestart_initial_1), v262)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v279 + v271
	v283 = int32(_a_F_pgstat_bestart_initial_0)
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgstat_bestart_initial[5])) = v285 - v271
	m.G0 = v19 + int32(320)
	return
L6:
	;
	v259 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v259)
	goto L5
L7:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+280))
	if v136 == int32(0) {
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
	v255 = F_strlen(m, v244)
	mBase = m.M
	goto L9
L12:
	;
	goto L13
L13:
	;
	v145 = int32(63)
	if (v23^v136)&int32(3) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v248 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v245))) = uint8(v248)
	goto L10
L15:
	;
	v229 = v224
	v230 = v225
	v231 = v226
	goto L36
L16:
	;
	if v219 == int32(0) {
		v244 = v217
		v245 = v218
		goto L14
	} else {
		goto L35
	}
L17:
	;
	v217 = v136
	v218 = v23
	v219 = v145
	goto L16
L18:
	;
	goto L19
L19:
	;
	v149 = int32(0)
	if base.B2i32(v136&int32(3) == v149)|int32(0) == v149 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v185 == int32(0) {
		v244 = v182
		v245 = v183
		goto L14
	} else {
		goto L29
	}
L21:
	;
	v161 = v136
	v162 = v23
	v163 = v145
	goto L24
L22:
	;
	goto L23
L23:
	;
	v182 = v136
	v183 = v23
	v184 = v145
	v185 = int32(1)
	goto L20
L24:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v165)
	if v165 == int32(0) {
		v224 = v161
		v225 = v162
		v226 = v163
		goto L15
	} else {
		goto L26
	}
L25:
	;
	v182 = v176
	v183 = v170
	v184 = v172
	v185 = v174
	goto L20
L26:
	;
	v169 = int32(1)
	v170 = v162 + v169
	v172 = v163 - v169
	v173 = int32(0)
	v174 = base.B2i32(v172 != v173)
	v176 = v161 + v169
	if v176&int32(3) == v173 {
		v182 = v176
		v183 = v170
		v184 = v172
		v185 = v174
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if v172 != 0 {
		v161 = v176
		v162 = v170
		v163 = v172
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
	if base.B2i32(v188 == int32(0))|base.B2i32(base.Ui32(v184) < base.Ui32(int32(4))) != 0 {
		v217 = v182
		v218 = v183
		v219 = v184
		goto L16
	} else {
		goto L30
	}
L30:
	;
	v195 = v182
	v196 = v183
	v197 = v184
	goto L31
L31:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v203 = int32(-2139062144)
	if (int32(16843008)-v200|v200)&v203 != v203 {
		v224 = v195
		v225 = v196
		v226 = v197
		goto L15
	} else {
		goto L33
	}
L32:
	;
	v217 = v211
	v218 = v209
	v219 = v213
	goto L16
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v200
	v208 = int32(4)
	v209 = v196 + v208
	v211 = v195 + v208
	v213 = v197 - v208
	if base.Ui32(int32(3)) < base.Ui32(v213) {
		v195 = v211
		v196 = v209
		v197 = v213
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v224 = v217
	v225 = v218
	v226 = v219
	goto L15
L36:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v229))))
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v233)
	if v233 == int32(0) {
		v244 = v229
		v245 = v230
		goto L14
	} else {
		goto L38
	}
L37:
	;
	v244 = v240
	v245 = v238
	goto L14
L38:
	;
	v237 = int32(1)
	v238 = v230 + v237
	v240 = v229 + v237
	v242 = v231 - v237
	if v242 != 0 {
		v229 = v240
		v230 = v238
		v231 = v242
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
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
	v57 = base.I32_div_s(v51-v54, int32(768))
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
	v42 = v32 + int32(768)
	goto L12
L14:
	;
	v28 = v25 * int32(768)
	v29 = v21 + v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
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
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+780))
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
	return v88
L21:
	;
	if v58 == int32(0) {
		v88 = v3
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	if base.B2i32(base.Ui32(int32(16)) < base.Ui32(v62))|base.B2i32(int32(1)<<(uint(v62)%32)&int32(_a_F_pgstat_fetch_stat_backend_by_pid_0) == int32(0)) != 0 {
		v88 = v3
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v72 != l0 {
		v88 = v3
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
	v76 = int32(0)
	v79 = F_pgstat_fetch_entry(m, int32(6), v76, base.I64_extend_i32_s(v57), v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v79|base.B2i32(l1 == int32(0)) != 0 {
		v88 = v79
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v84 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v84
	v88 = v84
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
	if base.Ui32(l0-int32(1)) <= base.Ui32(int32(12)) {
		return l0*int32(84) + int32(_a_F_pgstat_get_kind_info_0)
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
func F_pgstat_index(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	v8 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v8
	v25 = int32(1)
	v27 = F_GetAccessStrategy(m, v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_LockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if base.Ui32(int32(1)) < base.Ui32(v35) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v46 = v35
	v47 = v25
	goto L9
L7:
	;
	v73 = v35
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = base.I64_extend_i32_u(v73) << (uint(int64(13)) % 64)
	F_relation_close(m, l0, int32(1))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L9:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_index[0]))
	if v51 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v73 = v63
	goto L8
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	m.T0[l1].(func(*base.Module, int32, int32, int32, int32))(m, v11, l0, v47, v27)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v57 = v47 + int32(1)
	if base.Ui32(v57) < base.Ui32(v46) {
		v47 = v57
		goto L9
	} else {
		goto L16
	}
L16:
	;
	F_LockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v63 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v57) < base.Ui32(v63) {
		v46 = v63
		v47 = v57
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L10
L21:
	;
	v84 = F_build_pgstattuple_type(m, v11, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v11 + int32(48)
	return v84
}
func F_pgstat_io_init_shmem_cb(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	F_LWLockInitialize(m, l0, int32(85))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_LWLockInitialize(m, l0+int32(16), int32(85))
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_LWLockInitialize(m, l0+int32(32), int32(85))
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_LWLockInitialize(m, l0+int32(48), int32(85))
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_LWLockInitialize(m, l0-int32(-64), int32(85))
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						F_LWLockInitialize(m, l0+int32(80), int32(85))
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							F_LWLockInitialize(m, l0+int32(96), int32(85))
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								F_LWLockInitialize(m, l0+int32(112), int32(85))
								v39 = m.ExcPending
								if v39 != 0 {
									return
								} else {
									F_LWLockInitialize(m, l0+int32(128), int32(85))
									v44 = m.ExcPending
									if v44 != 0 {
										return
									} else {
										F_LWLockInitialize(m, l0+int32(144), int32(85))
										v49 = m.ExcPending
										if v49 != 0 {
											return
										} else {
											F_LWLockInitialize(m, l0+int32(160), int32(85))
											v54 = m.ExcPending
											if v54 != 0 {
												return
											} else {
												F_LWLockInitialize(m, l0+int32(176), int32(85))
												v59 = m.ExcPending
												if v59 != 0 {
													return
												} else {
													F_LWLockInitialize(m, l0+int32(192), int32(85))
													v64 = m.ExcPending
													if v64 != 0 {
														return
													} else {
														F_LWLockInitialize(m, l0+int32(208), int32(85))
														v69 = m.ExcPending
														if v69 != 0 {
															return
														} else {
															F_LWLockInitialize(m, l0+int32(224), int32(85))
															v74 = m.ExcPending
															if v74 != 0 {
																return
															} else {
																F_LWLockInitialize(m, l0+int32(240), int32(85))
																v79 = m.ExcPending
																if v79 != 0 {
																	return
																} else {
																	F_LWLockInitialize(m, l0+int32(256), int32(85))
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return
																	} else {
																		F_LWLockInitialize(m, l0+int32(272), int32(85))
																		v89 = m.ExcPending
																		if v89 != 0 {
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
	}
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
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
			v24 = int32(0)
			v26 = int32(_a_F_pgstat_progress_incr_param_1)
			v27 = base.AtomicRmwOr32(m, v24, v26, v24)
			v32 = v5 + l0<<(uint(int32(3))%32) + int32(232)
			v33 = *(*int64)(unsafe.Add(mBase, uint32(v32)))
			*(*int64)(unsafe.Add(mBase, uint32(v32))) = v33 + l1
			v39 = base.AtomicRmwOr32(m, v24, v26, v24)
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v40 + v17
			v46 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_incr_param[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_progress_incr_param[2])) = v46 - v17
		}
	}
	return
}
func F_pgstat_relation(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int64
	_ = v81
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int64
	_ = v100
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v129 int64
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int64
	_ = v175
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v206 int32
	_ = v206
	var v210 int64
	_ = v210
	var v222 int32
	_ = v222
	var v226 int64
	_ = v226
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int64
	_ = v272
	var v274 int32
	_ = v274
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v390 int64
	_ = v390
	var v391 int32
	_ = v391
	var v401 int64
	_ = v401
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v429 int64
	_ = v429
	var v430 int32
	_ = v430
	var v446 int64
	_ = v446
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
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
	v454 = m.ExcPending
	if v454 != 0 {
		goto L22
	} else {
		goto L127
	}
L2:
	;
	m.G0 = v18 + int32(160)
	return v446
L3:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+18)))
	if v377 == int32(0) {
		goto L1
	} else {
		goto L104
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L22
	} else {
		goto L99
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L22
	} else {
		goto L95
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
	v324 = m.ExcPending
	if v324 != 0 {
		goto L22
	} else {
		goto L91
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
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[0]))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgstat_relation[1])))
	if v49&int32(1) != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v52 = int32(0)
	goto L18
L17:
	;
	v52 = v47
	goto L18
L18:
	;
	if v52 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v55 = int32(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v63 = m.T0[v62].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, int32(_a_F_pgstat_relation_0), v55, v55, v55, int32(321))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L22
	} else {
		goto L88
	}
L22:
	;
	return int64(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = int32(4)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+40))
	v70 = F_heap_getnext(m, v63)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if v70 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = v70
	v77 = v55
	v81 = v10
	v83 = v10
	v84 = v10
	v85 = v10
	v86 = v10
	goto L28
L26:
	;
	v206 = v55
	v210 = v10
	goto L27
L27:
	;
	if base.Ui32(v206) < base.Ui32(v69) {
		goto L64
	} else {
		goto L65
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[2]))
	if v88 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v109
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = v110
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v108
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v188
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v107
	v206 = v184
	v210 = v188
	goto L27
L30:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L22
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v63)+60))
	F_LockBufferInternal(m, v91, int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L22
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v63)+60))
	v98 = F_HeapTupleSatisfiesVisibility(m, v76, v18+int32(40), v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L22
	} else {
		goto L35
	}
L35:
	;
	v100 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v76))))
	if v98 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v63)+60))
	F_UnlockBuffer(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L22
	} else {
		goto L40
	}
L37:
	;
	v107 = v83 + int64(1)
	v108 = v100 + v84
	v109 = v85
	v110 = v86
	goto L36
L38:
	;
	goto L39
L39:
	;
	v107 = v83
	v108 = v84
	v109 = v85 + int64(1)
	v110 = v100 + v86
	goto L36
L40:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+6)))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
	v118 = v114 | v115<<(uint(int32(16))%32)
	if base.Ui32(v77) <= base.Ui32(v118) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v125 = v77
	v129 = v81
	goto L44
L42:
	;
	v184 = v77
	v188 = v81
	goto L43
L43:
	;
	v194 = F_heap_getnext(m, v63)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L22
	} else {
		goto L62
	}
L44:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[2]))
	if v136 != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v184 = v177
	v188 = v175
	goto L43
L46:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L22
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v139 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v63)+64))
	v142 = F_ReadBufferExtended(m, l0, v139, v125, v139, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L22
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	F_LockBufferInternal(m, v142, int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	if v142 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+14)))
	v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v164)+12)))
	v167 = v165 - v166
	v168 = int32(0)
	if v168 < v167 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[3]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150+(v142^int32(-1))<<(uint(int32(2))%32))))
	v164 = v156
	goto L52
L54:
	;
	goto L55
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[4]))
	v164 = v158 + v142<<(uint(int32(13))%32) + int32(-8192)
	goto L52
L56:
	;
	F_UnlockReleaseBuffer(m, v142)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L22
	} else {
		goto L60
	}
L57:
	;
	v171 = v167
	goto L59
L58:
	;
	v171 = v168
	goto L59
L59:
	;
	goto L56
L60:
	;
	v175 = v129 + base.I64_extend_i32_u(v171)
	v177 = v125 + int32(1)
	if base.Ui32(v177) <= base.Ui32(v118) {
		v125 = v177
		v129 = v175
		goto L44
	} else {
		goto L61
	}
L61:
	;
	goto L45
L62:
	;
	if v194 != 0 {
		v76 = v194
		v77 = v184
		v81 = v188
		v83 = v107
		v84 = v108
		v85 = v109
		v86 = v110
		goto L28
	} else {
		goto L63
	}
L63:
	;
	goto L29
L64:
	;
	v222 = v206
	v226 = v210
	goto L67
L65:
	;
	goto L66
L66:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)+188))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v293)+12))
	m.T0[v294].(func(*base.Module, int32))(m, v63)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L22
	} else {
		goto L85
	}
L67:
	;
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[2]))
	if v233 != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+152)) = v272
	goto L66
L69:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L22
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v236 = int32(0)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v63)+64))
	v239 = F_ReadBufferExtended(m, l0, v236, v222, v236, v238)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L22
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_LockBufferInternal(m, v239, int32(1))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L22
	} else {
		goto L74
	}
L74:
	;
	if v239 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+14)))
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+12)))
	v264 = v262 - v263
	v265 = int32(0)
	if v265 < v264 {
		goto L80
	} else {
		goto L81
	}
L76:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[3]))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v247+(v239^int32(-1))<<(uint(int32(2))%32))))
	v261 = v253
	goto L75
L77:
	;
	goto L78
L78:
	;
	v255 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_relation[4]))
	v261 = v255 + v239<<(uint(int32(13))%32) + int32(-8192)
	goto L75
L79:
	;
	F_UnlockReleaseBuffer(m, v239)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L22
	} else {
		goto L83
	}
L80:
	;
	v268 = v264
	goto L82
L81:
	;
	v268 = v265
	goto L82
L82:
	;
	goto L79
L83:
	;
	v272 = v226 + base.I64_extend_i32_u(v268)
	v274 = v222 + int32(1)
	if v274 != v69 {
		v222 = v274
		v226 = v272
		goto L67
	} else {
		goto L84
	}
L84:
	;
	goto L68
L85:
	;
	F_relation_close(m, l0, int32(1))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L22
	} else {
		goto L86
	}
L86:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = base.I64_extend_i32_u(v69) << (uint(int64(13)) % 64)
	v306 = F_build_pgstattuple_type(m, v18+int32(112), l1)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L22
	} else {
		goto L87
	}
L87:
	;
	v446 = v306
	goto L2
L88:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_relation_1), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L22
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_2), int32(931), int32(_a_F_pgstat_relation_3))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L22
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L22
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(_a_F_pgstat_relation_4), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L22
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_5), int32(335), int32(_a_F_pgstat_relation_6))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L22
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L22
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_pgstat_relation_7), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L22
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_5), int32(255), int32(_a_F_pgstat_relation_8))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L22
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L22
	} else {
		goto L100
	}
L100:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v360 + int32(4)
	F_errmsg(m, int32(_a_F_pgstat_relation_9), v18)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L22
	} else {
		goto L101
	}
L101:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v368 = int32(*(*int8)(unsafe.Add(mBase, uint32(v367)+119)))
	F_errdetail_relkind_not_supported(m, v368)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L22
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_5), int32(306), int32(_a_F_pgstat_relation_8))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L22
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
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
	if v380 <= int32(782) {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	v429 = F_pgstat_index(m, l0, int32(_a_F_pgstat_relation_10), l1)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L22
	} else {
		goto L126
	}
L106:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L22
	} else {
		goto L122
	}
L107:
	;
	v405 = int32(_a_F_pgstat_relation_11)
	goto L106
L108:
	;
	v405 = int32(_a_F_pgstat_relation_12)
	goto L106
L109:
	;
	v401 = F_pgstat_index(m, l0, int32(_a_F_pgstat_relation_13), l1)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L22
	} else {
		goto L121
	}
L110:
	;
	switch v380 - int32(403) {
	case 0:
		goto L105
	default:
		goto L107
	case 2:
		goto L109
	}
L111:
	;
	goto L112
L112:
	;
	if v380 <= int32(2741) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	if v380 != int32(783) {
		goto L107
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v380 == int32(2742) {
		v405 = int32(_a_F_pgstat_relation_14)
		goto L106
	} else {
		goto L118
	}
L116:
	;
	v390 = F_pgstat_index(m, l0, int32(_a_F_pgstat_relation_15), l1)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L22
	} else {
		goto L117
	}
L117:
	;
	v446 = v390
	goto L2
L118:
	;
	if v380 == int32(3580) {
		goto L108
	} else {
		goto L119
	}
L119:
	;
	if v380 != int32(4000) {
		goto L107
	} else {
		goto L120
	}
L120:
	;
	v405 = int32(_a_F_pgstat_relation_16)
	goto L106
L121:
	;
	v446 = v401
	goto L2
L122:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L22
	} else {
		goto L123
	}
L123:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v413 + int32(4)
	F_errmsg(m, int32(_a_F_pgstat_relation_17), v18+int32(16))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L22
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_5), int32(298), int32(_a_F_pgstat_relation_8))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L22
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	v446 = v429
	goto L2
L127:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L22
	} else {
		goto L128
	}
L128:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v458 + int32(4)
	F_errmsg(m, int32(_a_F_pgstat_relation_18), v18+int32(32))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L22
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_pgstat_relation_5), int32(269), int32(_a_F_pgstat_relation_8))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L22
	} else {
		goto L130
	}
L130:
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
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v127 int64
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v141 int64
	_ = v141
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v146 int64
	_ = v146
	var v151 int64
	_ = v151
	var v157 int64
	_ = v157
	var v163 int64
	_ = v163
	var v168 int64
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v194 int64
	_ = v194
	var v196 int64
	_ = v196
	var v197 int64
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var __phi215 int32
	_ = __phi215
	var v220 int32
	_ = v220
	var __phi220 int32
	_ = __phi220
	var v223 int32
	_ = v223
	var __phi223 int32
	_ = __phi223
	var v224 int32
	_ = v224
	var __phi224 int32
	_ = __phi224
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v229 int64
	_ = v229
	var v230 int64
	_ = v230
	var v231 int64
	_ = v231
	var v233 int64
	_ = v233
	var v238 int64
	_ = v238
	var v244 int64
	_ = v244
	var v249 int64
	_ = v249
	var v254 int64
	_ = v254
	var v264 int64
	_ = v264
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
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
	v329 = m.ExcPending
	if v329 != 0 {
		goto L16
	} else {
		goto L63
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L16
	} else {
		goto L60
	}
L3:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[0]))
	v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v137
	v139 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v139
	v141 = int64(23)
	v144 = int64(2388976653695081527)
	v145 = (v139 ^ int64(base.Ui64(v139)>>(uint(v141)%64))) * v144
	v146 = int64(47)
	v151 = int64(-8645972361240307355)
	v157 = (v137 ^ int64(base.Ui64(v137)>>(uint(v141)%64))) * v144
	v163 = ((v145^int64(base.Ui64(v145)>>(uint(v146)%64))^int64(-9208349263878056368))*v151 ^ int64(base.Ui64(v157)>>(uint(v146)%64)) ^ v157) * v151
	v168 = (int64(base.Ui64(v163)>>(uint(v141)%64)) ^ v163) * v144
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v136)+20))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v178 = base.I32_wrap_i64(int64(base.Ui64(v168)>>(uint(v146)%64)) ^ v168 - int64(base.Ui64(v168)>>(uint(int64(32))%64)))
	goto L37
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
	if base.Ui32(v22-int32(1)) <= base.Ui32(int32(12)) {
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
	v39 = v22*int32(84) + int32(_a_F_pgstat_release_entry_ref_0)
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
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)+28))
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[2]))
	F_dshash_delete_entry(m, v77, v66)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L16
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[2]))
	F_dshash_release_lock(m, v129, v66)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L16
	} else {
		goto L36
	}
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[3]))
	F_dsa_free(m, v81, v75)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	if base.Ui32(v74-int32(1)) <= base.Ui32(int32(12)) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112))))
	if v113&int32(8) == int32(0) {
		goto L3
	} else {
		goto L35
	}
L29:
	;
	v112 = v74*int32(84) + int32(_a_F_pgstat_release_entry_ref_0)
	goto L28
L30:
	;
	goto L31
L31:
	;
	if base.Ui32(int32(8)) < base.Ui32(v74-int32(24)) {
		v110 = int32(0)
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v112 = v110
	goto L28
L33:
	;
	v98 = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[1]))
	if v100 == v98 {
		v110 = v98
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v100+v74<<(uint(int32(2))%32)-int32(96))))
	v110 = v108
	goto L32
L35:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_release_entry_ref[4]))
	v127 = base.AtomicRmwSub64(m, v119+v74<<(uint(int32(3))%32)+int32(16), int32(0), int64(1))
	goto L3
L36:
	;
	goto L3
L37:
	;
	v188 = v178 & v177
	v191 = v176 + v188*int32(24)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+16)))
	switch v192 {
	case 0:
		goto L40
	case 1:
		goto L41
	default:
		goto L39
	}
L39:
	;
	v178 = v188 + int32(1)
	goto L37
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L16
	} else {
		goto L57
	}
L41:
	;
	v193 = *(*int64)(unsafe.Add(mBase, uint32(v191)))
	v194 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v196 = *(*int64)(unsafe.Add(mBase, uint32(v191)+8))
	v197 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	if v193^v194|(v196^v197) != int64(0) {
		goto L39
	} else {
		goto L42
	}
L42:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	v203 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v202 - v203
	v208 = (v188 + v203) & v177
	v211 = v176 + v208*int32(24)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+16)))
	if v212 != v203 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v291 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v283)+16)) = uint8(v291)
	if l1 != 0 {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v283 = v191
	goto L43
L45:
	;
	goto L46
L46:
	;
	__phi215 = v191
	__phi220 = v211
	__phi223 = v177
	__phi224 = v208
	v215 = __phi215
	v220 = __phi220
	v223 = __phi223
	v224 = __phi224
	goto L47
L47:
	;
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v220)+8))
	v226 = int64(23)
	v229 = int64(2388976653695081527)
	v230 = (int64(base.Ui64(v225)>>(uint(v226)%64)) ^ v225) * v229
	v231 = int64(47)
	v233 = *(*int64)(unsafe.Add(mBase, uint32(v220)))
	v238 = (int64(base.Ui64(v233)>>(uint(v226)%64)) ^ v233) * v229
	v244 = int64(-8645972361240307355)
	v249 = (int64(base.Ui64(v230)>>(uint(v231)%64)) ^ (v238^int64(base.Ui64(v238)>>(uint(v231)%64))^int64(-9208349263878056368))*v244 ^ v230) * v244
	v254 = (int64(base.Ui64(v249)>>(uint(v226)%64)) ^ v249) * v229
	if v224 == v223&base.I32_wrap_i64(int64(base.Ui64(v254)>>(uint(v231)%64))^v254-int64(base.Ui64(v254)>>(uint(int64(32))%64))) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v283 = v220
	goto L43
L49:
	;
	v283 = v215
	goto L43
L50:
	;
	goto L51
L51:
	;
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v220)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v215)+16)) = v264
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v220)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v215)+8)) = v266
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v220)))
	*(*int64)(unsafe.Add(mBase, uint32(v215))) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v136)+20))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v136)+12))
	v272 = int32(1)
	v274 = v271 & (v224 + v272)
	v277 = v270 + v274*int32(24)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+16)))
	if v278 == v272 {
		__phi215 = v220
		__phi220 = v277
		__phi223 = v271
		__phi224 = v274
		v215 = __phi215
		v220 = __phi220
		v223 = __phi223
		v224 = __phi224
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	F_pfree(m, l1)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L16
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	m.G0 = v13 + int32(16)
	return
L56:
	;
	goto L55
L57:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_release_entry_ref_1), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_pgstat_release_entry_ref_2), int32(684), int32(_a_F_pgstat_release_entry_ref_3))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errmsg_internal(m, int32(_a_F_pgstat_release_entry_ref_4), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L16
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_pgstat_release_entry_ref_2), int32(657), int32(_a_F_pgstat_release_entry_ref_3))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L16
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
	F_errmsg_internal(m, int32(_a_F_pgstat_release_entry_ref_5), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	F_errfinish(m, int32(_a_F_pgstat_release_entry_ref_2), int32(628), int32(_a_F_pgstat_release_entry_ref_3))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L16
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pgstat_replslot_reset_timestamp_cb(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = l1
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
		v22 = v12 + v10*int32(296)
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
						F_errfinish(m, int32(_a_F_pgstat_replslot_to_serialized_name_cb_2), int32(229), int32(_a_F_pgstat_replslot_to_serialized_name_cb_3))
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
				v30 = int32(0)
				v32 = int32(_a_F_pgstat_report_plan_id_1)
				v33 = base.AtomicRmwOr32(m, v30, v32, v30)
				*(*int64)(unsafe.Add(mBase, uint32(v5)+400)) = l0
				v38 = base.AtomicRmwOr32(m, v30, v32, v30)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v39 + v23
				v45 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_plan_id[2]))
				*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_plan_id[2])) = v45 - v23
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
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
			v21 = int32(0)
			v23 = int32(_a_F_pgstat_report_xact_timestamp_1)
			v24 = base.AtomicRmwOr32(m, v21, v23, v21)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = l0
			v29 = base.AtomicRmwOr32(m, v21, v23, v21)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v30 + v14
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_xact_timestamp[2]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgstat_report_xact_timestamp[2])) = v36 - v14
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
func F_pgstat_twophase_postabort(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) {
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
