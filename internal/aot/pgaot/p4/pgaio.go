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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[910]))
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[903]))
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[903]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	return (l0 - v4) >> (uint(int32(7)) % 32)
}
func F_pgaio_io_get_target_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_consts[905])))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	return v8
}
func F_pgaio_io_process_completion(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v67 int32
	_ = v67
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
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
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	F_pgaio_io_update_state(m, l0, int32(5))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = m.G0
	v31 = v29 - int32(112)
	m.G0 = v31
	v33 = int32(4548900)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v35 + int32(1)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+104)) = int32(64)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+108)) = v39
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v43 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v67 = v43
	goto L6
L4:
	;
	goto L5
L5:
	;
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v31)+104))
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
	v89 = v67 - int32(1)
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+int32(5)+v89))))
	v93 = v91 << (uint(int32(3)) % 32)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_consts[904])))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v97 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+(l0+int32(9))))))
	v102 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(int32(1)) < base.Ui32(v67) {
		v67 = v89
		goto L6
	} else {
		goto L34
	}
L11:
	;
	if v102 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_errhidestmt(m)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v31)+104))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+40)) = v191
	m.T0[v190].(func(*base.Module, int32, int32, int32, int32))(m, v31+int32(96), l0, v31+int32(40), v99)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L33
	}
L15:
	;
	F_errhidecontext(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _consts[903]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+24))
	goto L17
L17:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v115) <= base.Ui32(int32(2)) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124<<(uint(int32(2))%32))+uint32(_consts[905])))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+8))
	goto L22
L19:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v115<<(uint(int32(2))%32))+uint32(_consts[906])))
	v123 = v122
	goto L21
L20:
	;
	v123 = int32(0)
	goto L21
L21:
	;
	goto L18
L22:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v132) <= base.Ui32(int32(7)) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v31)+104))
	v145 = int32(base.Ui32(v141)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v145) <= base.Ui32(int32(4)) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132<<(uint(int32(2))%32))+uint32(_consts[907])))
	v140 = v139
	goto L26
L25:
	;
	v140 = int32(0)
	goto L26
L26:
	;
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31-int32(-64)))) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(68)))) = v91
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v93)+uint32(_consts[908])))
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(72)))) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(76)))) = v99
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(80)))) = v154
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(92)))) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v31)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(84)))) = v163 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v31+int32(88)))) = int32(base.Ui32(v163) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+48)) = (l0 - v110) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+52)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v31)+56)) = v130
	*(*int32)(unsafe.Add(mBase, uint32(v31)+60)) = v140
	F_errmsg_internal(m, int32(710346), v31+int32(48))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L31
	}
L28:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v145<<(uint(int32(2))%32))+uint32(_consts[909])))
	v154 = v153
	goto L30
L29:
	;
	v154 = int32(0)
	goto L30
L30:
	;
	goto L27
L31:
	;
	F_errfinish(m, int32(521227), int32(257), int32(471905))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L14
L33:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v31)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v31)+104)) = v199
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
	v322 = int32(4548900)
	v324 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v324 - int32(1)
	m.G0 = v31 + int32(112)
	F_pgaio_io_update_state(m, l0, int32(6))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
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
	v246 = *(*int32)(unsafe.Add(mBase, _consts[903]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+24))
	goto L41
L41:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v252) <= base.Ui32(int32(2)) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v261<<(uint(int32(2))%32))+uint32(_consts[905])))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	goto L46
L43:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v252<<(uint(int32(2))%32))+uint32(_consts[906])))
	v260 = v259
	goto L45
L44:
	;
	v260 = int32(0)
	goto L45
L45:
	;
	goto L42
L46:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v269) <= base.Ui32(int32(7)) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v31)+104))
	v282 = int32(base.Ui32(v278)>>(uint(int32(6))%32)) & int32(7)
	if base.Ui32(v282) <= base.Ui32(int32(4)) {
		goto L52
	} else {
		goto L53
	}
L48:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v269<<(uint(int32(2))%32))+uint32(_consts[907])))
	v277 = v276
	goto L50
L49:
	;
	v277 = int32(0)
	goto L50
L50:
	;
	goto L47
L51:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v291
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v31)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(v31)+32)) = v292
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v31)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v297 & int32(63)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = int32(base.Ui32(v297) >> (uint(int32(9)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = (l0 - v247) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+4)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v277
	F_errmsg_internal(m, int32(502752), v31)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L55
	}
L52:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v282<<(uint(int32(2))%32))+uint32(_consts[909])))
	v291 = v290
	goto L54
L53:
	;
	v291 = int32(0)
	goto L54
L54:
	;
	goto L51
L55:
	;
	F_errfinish(m, int32(521227), int32(270), int32(471905))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
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
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v340 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	if v338 == v340 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	F_pgaio_io_reclaim(m, l0)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
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
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v2 = l1
	v3 = l2
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	if base.Ui32(v2) < base.Ui32(int32(4)) {
		v17 = v2 << (uint(int32(3)) % 32)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[904])))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
		if v21 == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
			if v24 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v125 = m.ExcPending
				if v125 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v2
					F_errmsg_internal(m, int32(334087), v12+int32(16))
					mBase = m.M
					v131 = m.ExcPending
					if v131 != 0 {
						return
					} else {
						F_errfinish(m, int32(521227), int32(96), int32(164071))
						mBase = m.M
						v136 = m.ExcPending
						if v136 != 0 {
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
					v140 = m.ExcPending
					if v140 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(4)
						F_errmsg_internal(m, int32(490953), v12+int32(32))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return
						} else {
							F_errfinish(m, int32(521227), int32(99), int32(164071))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
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
									v42 = *(*int32)(unsafe.Add(mBase, _consts[903]))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
									v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
									if base.Ui32(v48) <= base.Ui32(int32(2)) {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v48<<(uint(int32(2))%32))+uint32(_consts[906])))
										v56 = v55
									} else {
										v56 = int32(0)
									}
									v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(2))%32))+uint32(_consts[905])))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
									v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
									if base.Ui32(v65) <= base.Ui32(int32(7)) {
										v72 = *(*int32)(unsafe.Add(mBase, uint32(v65<<(uint(int32(2))%32))+uint32(_consts[907])))
										v73 = v72
									} else {
										v73 = int32(0)
									}
									v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
									v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[908])))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v75
									*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v2
									*(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)))) = v74 + int32(1)
									*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v73
									*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v63
									*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = (l0 - v43) >> (uint(int32(7)) % 32)
									F_errmsg_internal(m, int32(187399), v12+int32(48))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return
									} else {
										F_errfinish(m, int32(521227), int32(106), int32(164071))
										mBase = m.M
										v96 = m.ExcPending
										if v96 != 0 {
											return
										} else {
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
											v104 = v102 + int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v104)
											m.G0 = v12 + int32(80)
											return
										}
									}
								}
							}
						} else {
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
							v104 = v102 + int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v104)
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
				v140 = m.ExcPending
				if v140 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(4)
					F_errmsg_internal(m, int32(490953), v12+int32(32))
					mBase = m.M
					v147 = m.ExcPending
					if v147 != 0 {
						return
					} else {
						F_errfinish(m, int32(521227), int32(99), int32(164071))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
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
								v42 = *(*int32)(unsafe.Add(mBase, _consts[903]))
								v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+24))
								v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
								if base.Ui32(v48) <= base.Ui32(int32(2)) {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(v48<<(uint(int32(2))%32))+uint32(_consts[906])))
									v56 = v55
								} else {
									v56 = int32(0)
								}
								v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(2))%32))+uint32(_consts[905])))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
								v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
								if base.Ui32(v65) <= base.Ui32(int32(7)) {
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v65<<(uint(int32(2))%32))+uint32(_consts[907])))
									v73 = v72
								} else {
									v73 = int32(0)
								}
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v17)+uint32(_consts[908])))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v75
								*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v2
								*(*int32)(unsafe.Add(mBase, uint32(v12-int32(-64)))) = v74 + int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v73
								*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = (l0 - v43) >> (uint(int32(7)) % 32)
								F_errmsg_internal(m, int32(187399), v12+int32(48))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return
								} else {
									F_errfinish(m, int32(521227), int32(106), int32(164071))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
										v104 = v102 + int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v104)
										m.G0 = v12 + int32(80)
										return
									}
								}
							}
						}
					} else {
						v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
						v104 = v102 + int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v104)
						m.G0 = v12 + int32(80)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v112 = m.ExcPending
		if v112 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v2
			F_errmsg_internal(m, int32(420626), v12)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return
			} else {
				F_errfinish(m, int32(521227), int32(93), int32(164071))
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[165]))
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
			F_errcontext_msg(m, int32(490252), v5)
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
