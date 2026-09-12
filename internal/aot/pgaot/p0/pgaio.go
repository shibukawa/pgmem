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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v10 = int32(4465412)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v12 + int32(1)
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
			F_errmsg_internal(m, int32(257448), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_errfinish(m, int32(488517), int32(141), int32(19135))
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
		v22 = *(*int32)(unsafe.Add(mBase, _consts[165]))
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
		v65 = *(*int32)(unsafe.Add(mBase, _consts[165]))
		v66 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66
		if v66 <= v63 {
			v74 = v63
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, _consts[166]))
			v74 = int32(0) - v72
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74
		F_pgaio_io_process_completion(m, l0, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v80 = int32(4465412)
			v82 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v82 - int32(1)
			return
		}
	case 2:
		v34 = *(*int32)(unsafe.Add(mBase, _consts[165]))
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
		v65 = *(*int32)(unsafe.Add(mBase, _consts[165]))
		v66 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66
		if v66 <= v63 {
			v74 = v63
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, _consts[166]))
			v74 = int32(0) - v72
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74
		F_pgaio_io_process_completion(m, l0, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v80 = int32(4465412)
			v82 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v82 - int32(1)
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
			v80 = int32(4465412)
			v82 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v82 - int32(1)
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v3 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[123])))
	if v5 != v3 {
		v20 = v3
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
		if v8&int32(2) != 0 {
			v20 = v3
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(int32(2))%32))+uint32(_consts[557])))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v20 = base.B2i32(v17 == int32(0))
		}
	}
	return v20
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v149 int64
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(176)
	m.G0 = v16
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(176)
	return l0
L2:
	;
	v20 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	v271 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v275 = F_LWLockAcquire(m, v271+int32(6784), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L7
	} else {
		goto L61
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+v20<<(uint(int32(2))%32))))
	F_pgaio_io_update_state(m, v34, int32(4))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v70 = F_LWLockAcquire(m, v66+int32(6784), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L7
	} else {
		goto L13
	}
L7:
	;
	return int32(0)
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[558]))
	v43 = v41 + int32(152)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+156))
	if v44 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+160)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+156)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v41)+152)) = v43
	goto L11
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34)+28)) = v43
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v34)+24)) = v52
	v55 = v34 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v41)+152)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v41)+160))
	v59 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v41)+160)) = v58 + v59
	v63 = v20 + v59
	if v63 != l0 {
		v20 = v63
		goto L5
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v79 = v3
	v80 = v3
	v81 = v3
	goto L14
L14:
	;
	v87 = l1 + v81<<(uint(int32(2))%32)
	v89 = *(*int32)(unsafe.Add(mBase, _consts[559]))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v91 = int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	v96 = (v90 - v91) & (v93 + v91)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v96 == v97 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v237 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v237+int32(6784))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L7
	} else {
		goto L51
	}
L16:
	;
	v233 = v81 + int32(1)
	if v233 != l0 {
		v79 = v227
		v80 = v228
		v81 = v233
		goto L14
	} else {
		goto L50
	}
L17:
	;
	v101 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v130 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+24))
	goto L28
L20:
	;
	if v101 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_errhidestmt(m)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v16+int32(48)+v79<<(uint(int32(2))%32)))) = v124
	v227 = v79 + int32(1)
	v228 = v80
	goto L16
L24:
	;
	F_errhidecontext(m)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[559]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v109
	F_errmsg_internal(m, int32(121166), v16)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L7
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(487828), int32(191), int32(80870))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v89+v135<<(uint(int32(2))%32))+16)) = (v128 - v131) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v96
	if v80 != 0 {
		v227 = v79
		v228 = v80
		goto L16
	} else {
		goto L29
	}
L29:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[560]))
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v143)))
	if v144 == int64(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v163 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L7
	} else {
		goto L34
	}
L31:
	;
	v158 = int32(-1)
	v159 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v149 = base.I64_ctz(v144)
	*(*int64)(unsafe.Add(mBase, uint32(v143))) = v144 & base.I64_rotl(int64(-2), v149)
	v153 = base.I32_wrap_i64(v149)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v143+v153<<(uint(int32(3))%32))+8))
	v158 = v153
	v159 = v157
	goto L30
L34:
	;
	if v163 == int32(0) {
		v227 = v79
		v228 = v159
		goto L16
	} else {
		goto L35
	}
L35:
	;
	F_errhidestmt(m)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	F_errhidecontext(m)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v173 = *(*int32)(unsafe.Add(mBase, _consts[556]))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+24))
	goto L38
L38:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
	if base.Ui32(v180) <= base.Ui32(int32(2)) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v190<<(uint(int32(2))%32))+uint32(_consts[557])))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+8))
	goto L43
L40:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v180<<(uint(int32(2))%32))+uint32(_consts[561])))
	v188 = v187
	goto L42
L41:
	;
	v188 = int32(0)
	goto L42
L42:
	;
	goto L39
L43:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if base.Ui32(v199) <= base.Ui32(int32(7)) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v207
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = (v171 - v174) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(464144), v16+int32(16))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L7
	} else {
		goto L48
	}
L45:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v199<<(uint(int32(2))%32))+uint32(_consts[562])))
	v207 = v206
	goto L47
L46:
	;
	v207 = int32(0)
	goto L47
L47:
	;
	goto L44
L48:
	;
	F_errfinish(m, int32(487828), int32(276), int32(306201))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v227 = v79
	v228 = v159
	goto L16
L50:
	;
	goto L15
L51:
	;
	if v228 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_SetLatch(m, v228)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L7
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	if v227 <= int32(0) {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v248 = int32(0)
	goto L57
L57:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(48)+v248<<(uint(int32(2))%32))))
	F_pgaio_io_perform_synchronously(m, v264)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L59
	}
L58:
	;
	goto L1
L59:
	;
	v268 = v248 + int32(1)
	if v268 != v227 {
		v248 = v268
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v278 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v278+int32(6784))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	goto L1
}
