package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgaio_io_perform_synchronously(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v10 = int32(_a_F_pgaio_io_perform_synchronously_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1])) = v12 + int32(1)
	v18 = v8 + v9<<(uint(int32(3))%32)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	switch v20 {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_pgaio_io_perform_synchronously_1), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_pgaio_io_perform_synchronously_2), int32(141), int32(_a_F_pgaio_io_perform_synchronously_3))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 1:
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[2]))
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(167772181)
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+92)))
		if v27 != int32(1) {
			v59 = F_preadv(m, v26, v18, v27, v25)
			mBase = m.M
			v63 = v59
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v32 = F_pread(m, v26, v30, v31, v25)
			mBase = m.M
			v63 = v32
		}
		v65 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[2]))
		v66 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66
		if v66 <= v63 {
			v74 = v63
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[3]))
			v74 = int32(0) - v72
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74
		F_pgaio_io_process_completion(m, l0, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v80 = int32(_a_F_pgaio_io_perform_synchronously_0)
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1])) = v82 - int32(1)
			return
		}
	case 2:
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[2]))
		*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(167772184)
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+92)))
		if v39 == int32(1) {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v44 = F_pwrite(m, v38, v42, v43, v37)
			mBase = m.M
			v63 = v44
		} else {
			v45 = F_pwritev(m, v38, v18, v39, v37)
			mBase = m.M
			v63 = v45
		}
		v65 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[2]))
		v66 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66
		if v66 <= v63 {
			v74 = v63
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[3]))
			v74 = int32(0) - v72
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74
		F_pgaio_io_process_completion(m, l0, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v80 = int32(_a_F_pgaio_io_perform_synchronously_0)
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1])) = v82 - int32(1)
			return
		}
	default:
		v74 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74
		F_pgaio_io_process_completion(m, l0, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v80 = int32(_a_F_pgaio_io_perform_synchronously_0)
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1])) = v82 - int32(1)
			return
		}
	}
}
func F_pgaio_worker_needs_synchronous_execution(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgaio_worker_needs_synchronous_execution[0])))
	if v5 != v3 {
		v18 = v3
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
		if v8&int32(2) != 0 {
			v18 = v3
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(int32(2))%32))+uint32(_c_F_pgaio_worker_needs_synchronous_execution[1])))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v18 = base.B2i32(v15 == int32(0))
		}
	}
	return v18
}
func F_pgaio_worker_shmem_size(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_add_size(m, int32(272), int32(264))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_pgaio_worker_submit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v288 int32
	_ = v288
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(176)
	m.G0 = v17
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(176)
	return l0
L2:
	;
	v21 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[0]))
	v316 = F_LWLockAcquire(m, v312+int32(_a_F_pgaio_worker_submit_0), int32(0))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L7
	} else {
		goto L74
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1+v21<<(uint(int32(2))%32))))
	F_pgaio_io_update_state(m, v36, int32(4))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[0]))
	v72 = F_LWLockAcquire(m, v68+int32(_a_F_pgaio_worker_submit_0), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L13
	}
L7:
	;
	return int32(0)
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[1]))
	v45 = v43 + int32(152)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+156))
	if v46 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+160)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+156)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v43)+152)) = v45
	goto L11
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+28)) = v45
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v43)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v54
	v57 = v36 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v43)+152)) = v57
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v43)+160))
	v61 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+160)) = v60 + v61
	v65 = v21 + v61
	if v65 != l0 {
		v21 = v65
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v83 = v3
	v84 = v3
	v86 = v3
	goto L14
L14:
	;
	v92 = l1 + v86<<(uint(int32(2))%32)
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[2]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v96 = int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v101 = (v95 - v96) & (v98 + v96)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	if v101 == v102 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[0]))
	F_LWLockRelease(m, v236+int32(_a_F_pgaio_worker_submit_0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L51
	}
L16:
	;
	v232 = v86 + int32(1)
	if v232 != l0 {
		v83 = v225
		v84 = v226
		v86 = v232
		goto L14
	} else {
		goto L50
	}
L17:
	;
	v106 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[3]))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)+24))
	goto L28
L20:
	;
	if v106 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_errhidestmt(m)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(48)+v83<<(uint(int32(2))%32)))) = v129
	v225 = v83 + int32(1)
	v226 = v84
	goto L16
L24:
	;
	F_errhidecontext(m)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[2]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v114
	F_errmsg_internal(m, int32(_a_F_pgaio_worker_submit_1), v17)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_pgaio_worker_submit_2), int32(191), int32(_a_F_pgaio_worker_submit_3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v94+v140<<(uint(int32(2))%32))+16)) = (v133 - v136) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v101
	if v84 != 0 {
		v225 = v83
		v226 = v84
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[4]))
	v149 = *(*int64)(unsafe.Add(mBase, uint32(v148)))
	if v149 == int64(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v168 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L34
	}
L31:
	;
	v163 = int32(-1)
	v164 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v154 = base.I64_ctz(v149)
	*(*int64)(unsafe.Add(mBase, uint32(v148))) = v149 & base.I64_rotl(int64(-2), v154)
	v158 = base.I32_wrap_i64(v154)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v148+v158<<(uint(int32(3))%32))+8))
	v163 = v158
	v164 = v162
	goto L30
L34:
	;
	if v168 == int32(0) {
		v225 = v83
		v226 = v164
		goto L16
	} else {
		goto L35
	}
L35:
	;
	F_errhidestmt(m)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	F_errhidecontext(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[3]))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+24))
	goto L38
L38:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+2)))
	if base.Ui32(v184) <= base.Ui32(int32(2)) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+1)))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v193<<(uint(int32(2))%32))+uint32(_c_F_pgaio_worker_submit[5])))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+8))
	goto L43
L40:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184<<(uint(int32(2))%32))+uint32(_c_F_pgaio_worker_submit[6])))
	v191 = v189
	goto L42
L41:
	;
	v191 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
	if base.Ui32(v199) <= base.Ui32(int32(7)) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(32)))) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v206
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v197
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = (v176 - v179) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_pgaio_worker_submit_4), v17+int32(16))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L48
	}
L45:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v199<<(uint(int32(2))%32))+uint32(_c_F_pgaio_worker_submit[7])))
	v206 = v204
	goto L47
L46:
	;
	v206 = int32(0)
	goto L47
L47:
	;
	goto L44
L48:
	;
	F_errfinish(m, int32(_a_F_pgaio_worker_submit_2), int32(276), int32(_a_F_pgaio_worker_submit_5))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v225 = v83
	v226 = v164
	goto L16
L50:
	;
	goto L15
L51:
	;
	if v226 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v226)))
	if v241 != 0 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	if v225 <= int32(0) {
		goto L1
	} else {
		goto L69
	}
L55:
	;
	goto L54
L56:
	;
	goto L55
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v226))) = int32(1)
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v244 == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	if v247 == int32(0) {
		goto L56
	} else {
		goto L59
	}
L59:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[8]))
	if v251 == v247 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v253 = m.G0
	v255 = v253 - int32(16)
	m.G0 = v255
	v258 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[9]))
	if v258 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	v281 = F_pgmem_kill(m, v247, int32(23))
	mBase = m.M
	goto L56
L63:
	;
	m.G0 = v255 + int32(16)
	goto L55
L64:
	;
	v261 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v255)+15)) = uint8(v261)
	goto L65
L65:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[10]))
	v269 = F_write(m, v265, v255+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v269 {
		goto L63
	} else {
		goto L67
	}
L66:
	;
	goto L63
L67:
	;
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[11]))
	if v273 == int32(27) {
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v288 = int32(0)
	goto L70
L70:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(48)+v288<<(uint(int32(2))%32))))
	F_pgaio_io_perform_synchronously(m, v305)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L7
	} else {
		goto L72
	}
L71:
	;
	goto L1
L72:
	;
	v309 = v288 + int32(1)
	if v309 != v225 {
		v288 = v309
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[0]))
	F_LWLockRelease(m, v319+int32(_a_F_pgaio_worker_submit_0))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L7
	} else {
		goto L75
	}
L75:
	;
	goto L1
}
