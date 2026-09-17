package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgaio_error_cleanup(m *base.Module) {
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_error_cleanup[0]))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+20)))
	if v4 == int32(1) {
		v7 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v3)+20)) = uint8(v7)
		F_pgaio_submit_staged(m)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_pgaio_io_get_handle_data(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v3)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_get_handle_data[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	return v7 + v8<<(uint(int32(3))%32)
}
func F_pgaio_io_get_id(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_get_id[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	return (l0 - v4) >> (uint(int32(7)) % 32)
}
func F_pgaio_io_get_target_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_get_target_name[0])))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	return v6
}
func F_pgaio_io_process_completion(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v191 int32
	_ = v191
	var v192 int64
	_ = v192
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v235 int64
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int64
	_ = v287
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	F_pgaio_io_update_state(m, l0, int32(5))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v31 = m.G0
	v33 = v31 - int32(112)
	m.G0 = v33
	v35 = int32(_a_F_pgaio_io_process_completion_0)
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_process_completion[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_process_completion[0])) = v37 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+104)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v33)+108)) = v41
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v46 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v68 = v46
	v70 = v42
	v72 = v41
	goto L6
L4:
	;
	goto L5
L5:
	;
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v33)+104))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v235
	v239 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L35
	}
L6:
	;
	v94 = v68 - int32(1)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(5)+v94))))
	v98 = v96 << (uint(int32(3)) % 32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_pgaio_io_process_completion[1])))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v100 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+(l0+int32(9))))))
	v105 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v198 = v70
	v200 = v72
	goto L10
L10:
	;
	if base.Ui32(int32(1)) < base.Ui32(v68) {
		v68 = v94
		v70 = v198
		v72 = v200
		goto L6
	} else {
		goto L34
	}
L11:
	;
	if v105 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_errhidestmt(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v33)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = v184
	m.T0[v183].(func(*base.Module, int32, int32, int32, int32))(m, v33+int32(96), l0, v33+int32(40), v102)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L33
	}
L15:
	;
	F_errhidecontext(m)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_process_completion[2]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)+24))
	goto L17
L17:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v117) <= base.Ui32(int32(2)) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v125<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_process_completion[3])))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	goto L22
L19:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_process_completion[4])))
	v124 = v122
	goto L21
L20:
	;
	v124 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v130) <= base.Ui32(int32(7)) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v141 = int32(base.Ui32(v70)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v141) <= base.Ui32(int32(4)) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_process_completion[5])))
	v137 = v135
	goto L26
L25:
	;
	v137 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(92)))) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(88)))) = int32(base.Ui32(v70) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(84)))) = v70 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(80)))) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(76)))) = v102
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_pgaio_io_process_completion[6])))
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(72)))) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v33+int32(68)))) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v33-int32(-64)))) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v33)+60)) = v137
	*(*int32)(unsafe.Add(mBase, uint32(v33)+56)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v33)+52)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v33)+48)) = (l0 - v113) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_pgaio_io_process_completion_1), v33+int32(48))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v141<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_process_completion[7])))
	v148 = v146
	goto L30
L29:
	;
	v148 = int32(0)
	goto L30
L30:
	;
	goto L27
L31:
	;
	F_errfinish(m, int32(_a_F_pgaio_io_process_completion_2), int32(257), int32(_a_F_pgaio_io_process_completion_3))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L14
L33:
	;
	v192 = *(*int64)(unsafe.Add(mBase, uint32(v33)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+104)) = v192
	v198 = base.I32_wrap_i64(v192)
	v200 = base.I32_wrap_i64(int64(base.Ui64(v192) >> (uint(int64(32)) % 64)))
	goto L10
L34:
	;
	goto L7
L35:
	;
	if v239 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_errhidestmt(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v314 = int32(_a_F_pgaio_io_process_completion_0)
	v316 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_process_completion[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_process_completion[0])) = v316 - int32(1)
	m.G0 = v33 + int32(112)
	F_pgaio_io_update_state(m, l0, int32(6))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L57
	}
L39:
	;
	F_errhidecontext(m)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_process_completion[2]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+24))
	goto L41
L41:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v251) <= base.Ui32(int32(2)) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_process_completion[3])))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	goto L46
L43:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v251<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_process_completion[4])))
	v258 = v256
	goto L45
L44:
	;
	v258 = int32(0)
	goto L45
L45:
	;
	goto L42
L46:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v264) <= base.Ui32(int32(7)) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v272 = base.I32_wrap_i64(v235)
	v276 = int32(base.Ui32(v272)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v276) <= base.Ui32(int32(4)) {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v264<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_process_completion[5])))
	v271 = v269
	goto L50
L49:
	;
	v271 = int32(0)
	goto L50
L50:
	;
	goto L47
L51:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v284
	v287 = int64(base.Ui64(v235) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v33)+28)) = uint32(v287)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+24)) = int32(base.Ui32(v272) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = v272 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v283
	*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = (l0 - v247) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_pgaio_io_process_completion_4), v33)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L55
	}
L52:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v276<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_process_completion[7])))
	v283 = v281
	goto L54
L53:
	;
	v283 = int32(0)
	goto L54
L54:
	;
	goto L51
L55:
	;
	F_errfinish(m, int32(_a_F_pgaio_io_process_completion_2), int32(270), int32(_a_F_pgaio_io_process_completion_3))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L38
L57:
	;
	F_ConditionVariableBroadcast(m, l0+int32(56))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v332 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_process_completion[8]))
	if v330 == v332 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_pgaio_io_reclaim(m, l0)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	return
L62:
	;
	goto L61
}
func F_pgaio_io_register_callbacks(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v2 = l1
	v3 = l2
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	if base.Ui32(v2) < base.Ui32(int32(4)) {
		v17 = v2 << (uint(int32(3)) % 32)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_pgaio_io_register_callbacks[0])))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		if v21 == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
			if v24 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v2
					F_errmsg_internal(m, int32(_a_F_pgaio_io_register_callbacks_0), v12+int32(16))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pgaio_io_register_callbacks_1), int32(96), int32(_a_F_pgaio_io_register_callbacks_2))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
				if base.Ui32(int32(4)) <= base.Ui32(v27) {
					F_errstart_cold(m, int32(23), int32(0))
					mBase = m.M
					v134 = m.ExcPending
					if v134 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(4)
						F_errmsg_internal(m, int32(_a_F_pgaio_io_register_callbacks_3), v12+int32(32))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_pgaio_io_register_callbacks_1), int32(99), int32(_a_F_pgaio_io_register_callbacks_2))
							mBase = m.M
							v146 = m.ExcPending
							if v146 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v30 = l0 + v27
					*(*uint8)(unsafe.Add(mBase, uint32(v30)+9)) = uint8(v3)
					*(*uint8)(unsafe.Add(mBase, uint32(v30)+5)) = uint8(v2)
					v35 = F_errstart(m, int32(12), int32(0))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return
					} else {
						if v35 != 0 {
							F_errhidestmt(m)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								F_errhidecontext(m)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_register_callbacks[1]))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
									v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									if base.Ui32(v47) <= base.Ui32(int32(2)) {
										v52 = *(*int32)(unsafe.Add(mBase, uint32(v47<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_register_callbacks[2])))
										v54 = v52
									} else {
										v54 = int32(0)
									}
									v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
									v58 = *(*int32)(unsafe.Add(mBase, uint32(v55<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_register_callbacks[3])))
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
									if base.Ui32(v60) <= base.Ui32(int32(7)) {
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v60<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_register_callbacks[4])))
										v67 = v65
									} else {
										v67 = int32(0)
									}
									v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_pgaio_io_register_callbacks[5])))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v69
									*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v2
									*(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)))) = v68 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v67
									*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v59
									*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v54
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = (l0 - v43) >> (uint(int32(7)) % 32)
									F_errmsg_internal(m, int32(_a_F_pgaio_io_register_callbacks_4), v12+int32(48))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_pgaio_io_register_callbacks_1), int32(106), int32(_a_F_pgaio_io_register_callbacks_2))
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return
										} else {
											v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
											v98 = v96 + int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v98)
											m.G0 = v12 + int32(80)
											return
										}
									}
								}
							}
						} else {
							v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
							v98 = v96 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v98)
							m.G0 = v12 + int32(80)
							return
						}
					}
				}
			}
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
			if base.Ui32(int32(4)) <= base.Ui32(v27) {
				F_errstart_cold(m, int32(23), int32(0))
				mBase = m.M
				v134 = m.ExcPending
				if v134 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(4)
					F_errmsg_internal(m, int32(_a_F_pgaio_io_register_callbacks_3), v12+int32(32))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_pgaio_io_register_callbacks_1), int32(99), int32(_a_F_pgaio_io_register_callbacks_2))
						mBase = m.M
						v146 = m.ExcPending
						if v146 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v30 = l0 + v27
				*(*uint8)(unsafe.Add(mBase, uint32(v30)+9)) = uint8(v3)
				*(*uint8)(unsafe.Add(mBase, uint32(v30)+5)) = uint8(v2)
				v35 = F_errstart(m, int32(12), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					if v35 != 0 {
						F_errhidestmt(m)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							F_errhidecontext(m)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								v42 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_register_callbacks[1]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
								v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								if base.Ui32(v47) <= base.Ui32(int32(2)) {
									v52 = *(*int32)(unsafe.Add(mBase, uint32(v47<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_register_callbacks[2])))
									v54 = v52
								} else {
									v54 = int32(0)
								}
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v55<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_register_callbacks[3])))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
								v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if base.Ui32(v60) <= base.Ui32(int32(7)) {
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v60<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_register_callbacks[4])))
									v67 = v65
								} else {
									v67 = int32(0)
								}
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_c_F_pgaio_io_register_callbacks[5])))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v69
								*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v2
								*(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)))) = v68 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v67
								*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v59
								*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v54
								*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = (l0 - v43) >> (uint(int32(7)) % 32)
								F_errmsg_internal(m, int32(_a_F_pgaio_io_register_callbacks_4), v12+int32(48))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_pgaio_io_register_callbacks_1), int32(106), int32(_a_F_pgaio_io_register_callbacks_2))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return
									} else {
										v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
										v98 = v96 + int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v98)
										m.G0 = v12 + int32(80)
										return
									}
								}
							}
						}
					} else {
						v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
						v98 = v96 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v98)
						m.G0 = v12 + int32(80)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v106 = m.ExcPending
		if v106 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v2
			F_errmsg_internal(m, int32(_a_F_pgaio_io_register_callbacks_5), v12)
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_pgaio_io_register_callbacks_1), int32(93), int32(_a_F_pgaio_io_register_callbacks_2))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_pgaio_worker_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if l0 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_error_callback[0]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9+v10*int32(640))+44))
		F_set_errcontext_domain(m, int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v14
			F_errcontext_msg(m, int32(_a_F_pgaio_worker_error_callback_0), v5)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
