package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HandleFatalError(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, _consts[773]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(1)
	goto L1
L1:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _consts[774]))
	if v18 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _consts[775])))
	if v14 != 0 {
		v16 = int32(6)
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v16 = int32(3)
	goto L2
L6:
	;
	goto L5
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[776]))
	if v42 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	if v18 == int32(4357952) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v23 = v18
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23-int32(12))))
	if base.Ui32(v28) <= base.Ui32(int32(16)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L7
L12:
	;
	F_signal_child(m, v23-int32(20), v16)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v35 != int32(4357952) {
		v23 = v35
		goto L10
	} else {
		goto L17
	}
L15:
	;
	return
L16:
	;
	goto L14
L17:
	;
	goto L11
L18:
	;
	v51 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[777])) = uint8(v51)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	switch v54 - int32(2) {
	case 0, 1, 2, 3:
		goto L25
	default:
		goto L21
	case 5, 6, 7, 8:
		goto L24
	}
L19:
	;
	switch v16 - int32(3) {
	case 0, 3:
		goto L20
	default:
		goto L18
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[779])) = int32(2)
	goto L18
L21:
	;
	v116 = *(*int64)(unsafe.Add(mBase, _consts[780]))
	if v116 == int64(0) {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[778])) = v111
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, _consts[778]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96<<(uint(int32(2))%32))+uint32(_consts[781])))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v101
	F_errmsg_internal(m, int32(174299), v6)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L15
	} else {
		goto L36
	}
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[782]))
	if v66 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v57 = int32(6)
	v60 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v60 == int32(0) {
		v111 = v57
		goto L22
	} else {
		goto L27
	}
L27:
	;
	v92 = v57
	v93 = int32(503848)
	goto L23
L28:
	;
	F_FreeWaitEventSet(m, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L15
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v69 = int32(4358020)
	v70 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[782])) = v70
	v75 = F_CreateWaitEventSet(m, v70, int32(1))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L15
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[782])) = v75
	v81 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	F_AddWaitEventToSet(m, v75, int32(1), int32(-1), v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	v84 = int32(11)
	v87 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	if v87 == int32(0) {
		v111 = v84
		goto L22
	} else {
		goto L35
	}
L35:
	;
	v92 = v84
	v93 = int32(522668)
	goto L23
L36:
	;
	F_errfinish(m, int32(476110), int32(3272), int32(340010))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	v111 = v92
	goto L22
L38:
	;
	v120 = F___time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _consts[780])) = v120
	goto L40
L39:
	;
	goto L40
L40:
	;
	m.G0 = v6 + int32(16)
	return
}
func F_HoldingBufferPinThatDelaysRecovery(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
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
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
	if v10 < int32(0) {
		v70 = int32(0)
		m.G0 = v6 + int32(16)
		return v70
	} else {
		v14 = v10 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v14
		v17 = *(*int32)(unsafe.Add(mBase, _consts[921]))
		if v14 == v17 {
			v66 = int32(4365360)
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
			v70 = base.B2i32(int32(0) < v67)
			m.G0 = v6 + int32(16)
			return v70
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[922]))
			if v14 == v21 {
				v66 = int32(4365368)
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				v70 = base.B2i32(int32(0) < v67)
				m.G0 = v6 + int32(16)
				return v70
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _consts[923]))
				if v14 == v25 {
					v66 = int32(4365376)
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
					v70 = base.B2i32(int32(0) < v67)
					m.G0 = v6 + int32(16)
					return v70
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _consts[924]))
					if v14 == v29 {
						v66 = int32(4365384)
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
						v70 = base.B2i32(int32(0) < v67)
						m.G0 = v6 + int32(16)
						return v70
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _consts[925]))
						if v14 == v33 {
							v66 = int32(4365392)
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
							v70 = base.B2i32(int32(0) < v67)
							m.G0 = v6 + int32(16)
							return v70
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, _consts[926]))
							if v14 == v37 {
								v66 = int32(4365400)
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
								v70 = base.B2i32(int32(0) < v67)
								m.G0 = v6 + int32(16)
								return v70
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, _consts[927]))
								if v14 == v41 {
									v66 = int32(4365408)
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
									v70 = base.B2i32(int32(0) < v67)
									m.G0 = v6 + int32(16)
									return v70
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, _consts[928]))
									if v14 == v45 {
										v66 = int32(4365416)
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
										v70 = base.B2i32(int32(0) < v67)
										m.G0 = v6 + int32(16)
										return v70
									} else {
										v48 = int32(0)
										v50 = *(*int32)(unsafe.Add(mBase, _consts[929]))
										if v50 == v48 {
											v70 = v48
											m.G0 = v6 + int32(16)
											return v70
										} else {
											v54 = *(*int32)(unsafe.Add(mBase, _consts[930]))
											v57 = int32(0)
											v59 = F_hash_search(m, v54, v6+int32(12), v57, v57)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												if v59 == int32(0) {
													v70 = v48
												} else {
													v66 = v59
													v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
													v70 = base.B2i32(int32(0) < v67)
												}
												m.G0 = v6 + int32(16)
												return v70
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
func F_handle_sig_alarm(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var __phi83 int32
	_ = __phi83
	var v88 int32
	_ = v88
	var __phi88 int32
	_ = __phi88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int64
	_ = v127
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v153 int64
	_ = v153
	var v155 int32
	_ = v155
	var v159 int64
	_ = v159
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(4444172)
	v13 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	F_SetLatch(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1438])) = v22
	v25 = *(*int32)(unsafe.Add(mBase, _consts[636]))
	if v25 == v22 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v172 = int32(4444172)
	v174 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v174 - int32(1)
	m.G0 = v9 + int32(16)
	return
L5:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[636])) = v29
	v32 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	if v32 <= v29 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v38 = m.G0
	v39 = int32(16)
	v40 = v38 - v39
	m.G0 = v40
	F___gettimeofday(m, v40)
	mBase = m.M
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	v44 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+8)))
	m.G0 = v40 + v39
	v52 = v44 + v43*int64(1000000) - int64(946684800000000)
	goto L7
L7:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	if v54 <= int32(0) {
		v159 = v52
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_schedule_alarm(m, v159)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L30
	}
L9:
	;
	v58 = v52
	goto L10
L10:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[1439]))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v64)+24))
	if v58 < v65 {
		v159 = v58
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v159 = v153
	goto L8
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[1439]))
	v70 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	if v70 <= int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _consts[1439]))
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)) = uint8(v73)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	if int32(2) <= v80 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	__phi83 = int32(1)
	__phi88 = v73
	v83 = __phi83
	v88 = __phi88
	goto L17
L15:
	;
	goto L16
L16:
	;
	v110 = int32(4448456)
	v112 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	v113 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[637])) = v112 - v113
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+5)) = uint8(v113)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	m.T0[v118].(func(*base.Module))(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v89 = int32(2)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v83<<(uint(v89)%32))+uint32(_consts[1439])))
	*(*int32)(unsafe.Add(mBase, uint32(v88<<(uint(v89)%32))+uint32(_consts[1439]))) = v97
	v100 = v83 + int32(1)
	v102 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	if v100 < v102 {
		__phi83 = v100
		__phi88 = v83
		v83 = __phi83
		v88 = __phi88
		goto L17
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	goto L18
L20:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	if int32(0) < v121 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v127 = base.I64_extend_i32_u(v121) * int64(1000)
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v68)+24))
	v130 = v129 + v127
	if v130 < v58 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v139 = m.G0
	v140 = int32(16)
	v141 = v139 - v140
	m.G0 = v141
	F___gettimeofday(m, v141)
	mBase = m.M
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
	v145 = int64(*(*int32)(unsafe.Add(mBase, uint32(v141)+8)))
	m.G0 = v141 + v140
	v153 = v145 + v144*int64(1000000) - int64(946684800000000)
	goto L28
L24:
	;
	v132 = v127 + v58
	goto L26
L25:
	;
	v132 = v130
	goto L26
L26:
	;
	F_enable_timeout(m, v124, v58, v132, v121)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	if int32(0) < v155 {
		v58 = v153
		goto L10
	} else {
		goto L29
	}
L29:
	;
	goto L11
L30:
	;
	goto L4
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
	v188 = *(*int32)(unsafe.Add(mBase, _consts[637]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v188 - int32(1)
	F_errmsg_internal(m, int32(448841), v9)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(473581), int32(143), int32(26384))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_has_parameter_privilege_id_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_convert_any_priv_string(m, v12, int32(1620560))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_text_to_cstring(m, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = F_pg_parameter_aclcheck(m, v17, v5, v15)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v19 == int32(0))
					}
				}
			}
		}
	}
}
func F_has_parameter_privilege_name_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v15 = F_convert_any_priv_string(m, v12, int32(1620560))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_get_role_oid_or_public(m, v5)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = F_text_to_cstring(m, v7)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v21 = F_pg_parameter_aclcheck(m, v19, v17, v15)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v21 == int32(0))
						}
					}
				}
			}
		}
	}
}
func F_hashbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 float64
	_ = v171
	var v172 int32
	_ = v172
	var v175 float64
	_ = v175
	var v179 int64
	_ = v179
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int64
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 float64
	_ = v324
	var v326 float64
	_ = v326
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	v4 = int32(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v21 = F_RelationGetNumberOfBlocksInFork(m, l1, v4)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v21 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_estimate_rel_size(m, l0, int32(0), v16+int32(-4), v16+int32(-16), v16+int32(-24))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L62
	}
L6:
	;
	v36 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
	v38 = F__hash_init(m, l1, v36, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[58]))
	v45 = int32(base.Ui32(v41)>>(uint(int32(3))%32)) & int32(524287)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+118)))
	if v49 == int32(116) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v52 = int32(4365444)
	goto L10
L9:
	;
	v52 = int32(4069496)
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if base.Ui32(v45) < base.Ui32(v53) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = v45
	goto L13
L12:
	;
	v55 = v53
	goto L13
L13:
	;
	if base.Ui32(v55) <= base.Ui32(v38) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v58 = F_palloc0(m, int32(20))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v153 = v4
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v153
	v160 = int32(1)
	v161 = int32(0)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+140))
	v171 = m.T0[v170].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v160, v161, v160, v161, int32(-1), int32(137), v16+int32(-48), v161)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L29
	}
L17:
	;
	v60 = int32(1)
	v61 = v38 - v60
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = l1
	v64 = int32(-1)
	v67 = v38 + v60
	if v38&v67 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v74 = v64<<(uint(int32(32)-base.I32_clz(v67))%32) ^ v64
	goto L20
L19:
	;
	v74 = v38
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v74
	v77 = int32(base.Ui32(v74) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v77
	v80 = *(*int32)(unsafe.Add(mBase, _consts[58]))
	v81 = m.G0
	v83 = v81 - int32(32)
	m.G0 = v83
	v85 = int32(0)
	v87 = F_tuplesort_begin_common(m, v80, v85, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v89 = int32(4449520)
	v90 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v92
	v95 = F_palloc(m, int32(20))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _consts[82])))
	if v98 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = int32(1846)
	v123 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+40)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v87)+60)) = v95
	*(*uint8)(unsafe.Add(mBase, uint32(v87)+36)) = uint8(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = int32(1847)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = int32(1848)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = int32(1851)
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = int32(1852)
	*(*int32)(unsafe.Add(mBase, uint32(v95)+16)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v95)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v95)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = l0
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v90
	m.G0 = v83 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v58))) = v87
	v153 = v58
	goto L16
L24:
	;
	v103 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v103 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = int32(102)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v74
	F_errmsg_internal(m, int32(482456), v83)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(474530), int32(467), int32(309940))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L23
L29:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v18)+48)) = v171
	v175 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	if base.F64_lt(base.F64_abs(v175), float64(9.223372036854776e+18)) != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v184 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v179 = base.I64_trunc_f64_s(v175)
	v181 = v179
	goto L30
L32:
	;
	goto L33
L33:
	;
	v181 = int64(-9223372036854775807 - 1)
	goto L30
L34:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v215 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	goto L34
L36:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v188 != int32(1) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v191 = int32(4444180)
	v193 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v194 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v193 + v194
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v197 + v194
	*(*int64)(unsafe.Add(mBase, uint32(v184+int32(88))+232)) = v181
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v205 + v194
	v211 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v211 - v194
	goto L35
L38:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	F_tuplesort_performsort(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v322 = F_palloc(m, int32(16))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L61
	}
L41:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v221 = F_tuplesort_getheaptuple(m, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v221 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v223 = v221
	v237 = int64(0)
	goto L46
L44:
	;
	goto L45
L45:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	F_tuplesort_end(m, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L59
	}
L46:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	F__hash_doinsert(m, v238, v223, v216, int32(1))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v243 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v248 = v237 + int64(1)
	v251 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v251 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L51
L53:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v283 = F_tuplesort_getheaptuple(m, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	goto L53
L55:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v255 != int32(1) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v258 = int32(4444180)
	v260 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v261 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v260 + v261
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v264 + v261
	*(*int64)(unsafe.Add(mBase, uint32(v251+int32(96))+232)) = v248
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	*(*int32)(unsafe.Add(mBase, uint32(v251))) = v272 + v261
	v278 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v278 - v261
	goto L54
L57:
	;
	if v283 != 0 {
		v223 = v283
		v237 = v248
		goto L46
	} else {
		goto L58
	}
L58:
	;
	goto L47
L59:
	;
	F_pfree(m, v300)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L40
L61:
	;
	v324 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v322))) = v324
	v326 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v322)+8)) = v326
	m.G0 = v18 - int32(-64)
	return v322
L62:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v336 + int32(4)
	F_errmsg_internal(m, int32(485010), v18)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(478671), int32(138), int32(413596))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hashendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	if v6 == int32(-1) {
		F__hash_dropscanbuf(m, v5)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
			if v16 != 0 {
				F_pfree(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_pfree(m, v5)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
						return
					}
				}
			} else {
				F_pfree(m, v5)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
					return
				}
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+20))
		if v9 <= int32(0) {
			F__hash_dropscanbuf(m, v5)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
				if v16 != 0 {
					F_pfree(m, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						F_pfree(m, v5)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
							return
						}
					}
				} else {
					F_pfree(m, v5)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
						return
					}
				}
			}
		} else {
			F__hash_kill_items(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				F__hash_dropscanbuf(m, v5)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
					if v16 != 0 {
						F_pfree(m, v16)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							F_pfree(m, v5)
							mBase = m.M
							v20 = m.ExcPending
							if v20 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
								return
							}
						}
					} else {
						F_pfree(m, v5)
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
							return
						}
					}
				}
			}
		}
	}
}
func F_hashinet(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(1)
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
		if v9&v7 != 0 {
			v12 = v7
		} else {
			v12 = int32(4)
		}
		v13 = v3 + v12
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
		if v16 == int32(2) {
			v19 = int32(6)
		} else {
			v19 = int32(18)
		}
		v25 = v19 - int32(1636608432)
		if v13&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v19) {
				v134 = v13
				v135 = v19
				v136 = v25
				v137 = v25
				v138 = v25
				for {
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
					v141 = v140 + v137
					v142 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
					v145 = v144 + v138
					v147 = int32(4)
					v149 = v142 + v136 - v145 ^ base.I32_rotl(v145, v147)
					v153 = v141 - v149 ^ base.I32_rotl(v149, int32(6))
					v154 = v145 + v141
					v155 = v149 + v154
					v156 = v153 + v155
					v160 = v154 - v153 ^ base.I32_rotl(v153, int32(8))
					v164 = v155 - v160 ^ base.I32_rotl(v160, int32(16))
					v168 = v156 - v164 ^ base.I32_rotl(v164, int32(19))
					v169 = v160 + v156
					v170 = v164 + v169
					v171 = v168 + v170
					v175 = v169 - v168 ^ base.I32_rotl(v168, v147)
					v176 = int32(12)
					v177 = v134 + v176
					v179 = v135 - v176
					if base.Ui32(int32(11)) < base.Ui32(v179) {
						v134 = v177
						v135 = v179
						v136 = v170
						v137 = v171
						v138 = v175
						continue
					} else {
						break
					}
					break
				}
				v182 = v177
				v183 = v179
				v184 = v170
				v185 = v171
				v186 = v175
			} else {
				v182 = v13
				v183 = v19
				v184 = v25
				v185 = v25
				v186 = v25
			}
			switch v183 - int32(1) {
			case 0:
				v245 = v184
				v246 = v185
				v247 = v186
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 1:
				v238 = v184
				v239 = v185
				v240 = v186
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 2:
				v231 = v184
				v232 = v185
				v233 = v186
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 3:
				v225 = v185
				v226 = v186
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
				v231 = v227<<(uint(int32(24))%32) + v184
				v232 = v225
				v233 = v226
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 4:
				v221 = v185
				v222 = v186
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v225 = v221 + v223
				v226 = v222
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
				v231 = v227<<(uint(int32(24))%32) + v184
				v232 = v225
				v233 = v226
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 5:
				v215 = v185
				v216 = v186
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
				v221 = v217<<(uint(int32(8))%32) + v215
				v222 = v216
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v225 = v221 + v223
				v226 = v222
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
				v231 = v227<<(uint(int32(24))%32) + v184
				v232 = v225
				v233 = v226
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 6:
				v209 = v185
				v210 = v186
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+6)))
				v215 = v211<<(uint(int32(16))%32) + v209
				v216 = v210
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
				v221 = v217<<(uint(int32(8))%32) + v215
				v222 = v216
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v225 = v221 + v223
				v226 = v222
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
				v231 = v227<<(uint(int32(24))%32) + v184
				v232 = v225
				v233 = v226
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 7:
				v204 = v186
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
				v209 = v205<<(uint(int32(24))%32) + v185
				v210 = v204
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+6)))
				v215 = v211<<(uint(int32(16))%32) + v209
				v216 = v210
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
				v221 = v217<<(uint(int32(8))%32) + v215
				v222 = v216
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v225 = v221 + v223
				v226 = v222
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
				v231 = v227<<(uint(int32(24))%32) + v184
				v232 = v225
				v233 = v226
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 8:
				v199 = v186
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
				v204 = v200<<(uint(int32(8))%32) + v199
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
				v209 = v205<<(uint(int32(24))%32) + v185
				v210 = v204
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+6)))
				v215 = v211<<(uint(int32(16))%32) + v209
				v216 = v210
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
				v221 = v217<<(uint(int32(8))%32) + v215
				v222 = v216
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v225 = v221 + v223
				v226 = v222
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
				v231 = v227<<(uint(int32(24))%32) + v184
				v232 = v225
				v233 = v226
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 9:
				v194 = v186
				v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+9)))
				v199 = v195<<(uint(int32(16))%32) + v194
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
				v204 = v200<<(uint(int32(8))%32) + v199
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
				v209 = v205<<(uint(int32(24))%32) + v185
				v210 = v204
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+6)))
				v215 = v211<<(uint(int32(16))%32) + v209
				v216 = v210
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
				v221 = v217<<(uint(int32(8))%32) + v215
				v222 = v216
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v225 = v221 + v223
				v226 = v222
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
				v231 = v227<<(uint(int32(24))%32) + v184
				v232 = v225
				v233 = v226
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			case 10:
				v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+10)))
				v194 = v190<<(uint(int32(24))%32) + v186
				v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+9)))
				v199 = v195<<(uint(int32(16))%32) + v194
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
				v204 = v200<<(uint(int32(8))%32) + v199
				v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+7)))
				v209 = v205<<(uint(int32(24))%32) + v185
				v210 = v204
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+6)))
				v215 = v211<<(uint(int32(16))%32) + v209
				v216 = v210
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
				v221 = v217<<(uint(int32(8))%32) + v215
				v222 = v216
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v225 = v221 + v223
				v226 = v222
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+3)))
				v231 = v227<<(uint(int32(24))%32) + v184
				v232 = v225
				v233 = v226
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v238 = v234<<(uint(int32(16))%32) + v231
				v239 = v232
				v240 = v233
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v245 = v241<<(uint(int32(8))%32) + v238
				v246 = v239
				v247 = v240
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v252 = v245 + v248
				v253 = v246
				v254 = v247
			default:
				v252 = v184
				v253 = v185
				v254 = v186
			}
		} else {
			if base.Ui32(v19) < base.Ui32(int32(12)) {
				v80 = v13
				v81 = v19
				v82 = v25
				v83 = v25
				v84 = v25
			} else {
				v32 = v13
				v33 = v19
				v34 = v25
				v35 = v25
				v36 = v25
				for {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
					v39 = v38 + v35
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+8))
					v43 = v42 + v36
					v45 = int32(4)
					v47 = v40 + v34 - v43 ^ base.I32_rotl(v43, v45)
					v51 = v39 - v47 ^ base.I32_rotl(v47, int32(6))
					v52 = v43 + v39
					v53 = v47 + v52
					v54 = v51 + v53
					v58 = v52 - v51 ^ base.I32_rotl(v51, int32(8))
					v62 = v53 - v58 ^ base.I32_rotl(v58, int32(16))
					v66 = v54 - v62 ^ base.I32_rotl(v62, int32(19))
					v67 = v58 + v54
					v68 = v62 + v67
					v69 = v66 + v68
					v73 = v67 - v66 ^ base.I32_rotl(v66, v45)
					v74 = int32(12)
					v75 = v32 + v74
					v77 = v33 - v74
					if base.Ui32(int32(11)) < base.Ui32(v77) {
						v32 = v75
						v33 = v77
						v34 = v68
						v35 = v69
						v36 = v73
						continue
					} else {
						break
					}
					break
				}
				v80 = v75
				v81 = v77
				v82 = v68
				v83 = v69
				v84 = v73
			}
			switch v81 - int32(1) {
			case 0:
				v131 = v82
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
				v252 = v131 + v132
				v253 = v83
				v254 = v84
			case 1:
				v126 = v82
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
				v131 = v127<<(uint(int32(8))%32) + v126
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
				v252 = v131 + v132
				v253 = v83
				v254 = v84
			case 2:
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+2)))
				v126 = v122<<(uint(int32(16))%32) + v82
				v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
				v131 = v127<<(uint(int32(8))%32) + v126
				v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
				v252 = v131 + v132
				v253 = v83
				v254 = v84
			case 3:
				v119 = v83
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v252 = v120 + v82
				v253 = v119
				v254 = v84
			case 4:
				v116 = v83
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
				v119 = v116 + v117
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v252 = v120 + v82
				v253 = v119
				v254 = v84
			case 5:
				v111 = v83
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+5)))
				v116 = v112<<(uint(int32(8))%32) + v111
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
				v119 = v116 + v117
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v252 = v120 + v82
				v253 = v119
				v254 = v84
			case 6:
				v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+6)))
				v111 = v107<<(uint(int32(16))%32) + v83
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+5)))
				v116 = v112<<(uint(int32(8))%32) + v111
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
				v119 = v116 + v117
				v120 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v252 = v120 + v82
				v253 = v119
				v254 = v84
			case 7:
				v102 = v84
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
				v252 = v103 + v82
				v253 = v105 + v83
				v254 = v102
			case 8:
				v97 = v84
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
				v102 = v98<<(uint(int32(8))%32) + v97
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
				v252 = v103 + v82
				v253 = v105 + v83
				v254 = v102
			case 9:
				v92 = v84
				v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
				v97 = v93<<(uint(int32(16))%32) + v92
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
				v102 = v98<<(uint(int32(8))%32) + v97
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
				v252 = v103 + v82
				v253 = v105 + v83
				v254 = v102
			case 10:
				v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+10)))
				v92 = v88<<(uint(int32(24))%32) + v84
				v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+9)))
				v97 = v93<<(uint(int32(16))%32) + v92
				v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+8)))
				v102 = v98<<(uint(int32(8))%32) + v97
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
				v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
				v252 = v103 + v82
				v253 = v105 + v83
				v254 = v102
			default:
				v252 = v82
				v253 = v83
				v254 = v84
			}
		}
		v257 = int32(14)
		v259 = v253 ^ v254 - base.I32_rotl(v253, v257)
		v263 = v259 ^ v252 - base.I32_rotl(v259, int32(11))
		v267 = v263 ^ v253 - base.I32_rotl(v263, int32(25))
		v271 = v267 ^ v259 - base.I32_rotl(v267, int32(16))
		v275 = v271 ^ v263 - base.I32_rotl(v271, int32(4))
		v279 = v275 ^ v267 - base.I32_rotl(v275, v257)
		return v279 ^ v271 - base.I32_rotl(v279, int32(24))
	}
}
func F_hashnameextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
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
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3&int32(3) == int32(0) {
		v27 = v3
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
	v68 = v60 - int32(1636608432)
	if v62 == int64(0) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v60 = v52 - v3
	goto L1
L3:
	;
	v31 = v27
	goto L12
L4:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v11 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v60 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v16 = v3
	goto L8
L8:
	;
	v20 = v16 + int32(1)
	if v20&int32(3) == int32(0) {
		v27 = v20
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v52 = v20
	goto L2
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v25 != 0 {
		v16 = v20
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v40 = int32(-2139062144)
	if (int32(16843008)-v37|v37)&v40 == v40 {
		v31 = v31 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v46 = v31
	goto L15
L14:
	;
	goto L13
L15:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v50 != 0 {
		v46 = v46 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v52 = v46
	goto L2
L17:
	;
	goto L16
L18:
	;
	v378 = F_Int64GetDatum(m, base.I64_extend_i32_u(v368)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v360^v368-base.I32_rotl(v368, int32(24))))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L61
	} else {
		goto L62
	}
L19:
	;
	if v3&int32(3) != 0 {
		goto L35
	} else {
		goto L36
	}
L20:
	;
	v105 = v68
	v107 = v68
	v109 = v68
	goto L19
L21:
	;
	goto L22
L22:
	;
	v72 = v68 + base.I32_wrap_i64(v62)
	v73 = v72 + v68
	v77 = int32(4)
	v79 = base.I32_wrap_i64(int64(base.Ui64(v62)>>(uint(int64(32))%64))) ^ base.I32_rotl(v68, v77)
	v83 = v72 - v79 ^ base.I32_rotl(v79, int32(6))
	v87 = v73 - v83 ^ base.I32_rotl(v83, int32(8))
	v88 = v79 + v73
	v89 = v83 + v88
	v90 = v87 + v89
	v94 = v88 - v87 ^ base.I32_rotl(v87, int32(16))
	v98 = v89 - v94 ^ base.I32_rotl(v94, int32(19))
	v103 = v94 + v90
	v105 = v103
	v107 = v90 - v98 ^ base.I32_rotl(v98, v77)
	v109 = v98 + v103
	goto L19
L23:
	;
	v346 = int32(14)
	v348 = v342 ^ v343 - base.I32_rotl(v342, v346)
	v352 = v348 ^ v341 - base.I32_rotl(v348, int32(11))
	v356 = v352 ^ v342 - base.I32_rotl(v352, int32(25))
	v360 = v356 ^ v348 - base.I32_rotl(v356, int32(16))
	v364 = v360 ^ v352 - base.I32_rotl(v360, int32(4))
	v368 = v364 ^ v356 - base.I32_rotl(v364, v346)
	goto L18
L24:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	v341 = v333 + v336
	v342 = v334
	v343 = v335
	goto L23
L25:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	v333 = v329<<(uint(int32(8))%32) + v326
	v334 = v327
	v335 = v328
	goto L24
L26:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+2)))
	v326 = v322<<(uint(int32(16))%32) + v319
	v327 = v320
	v328 = v321
	goto L25
L27:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+3)))
	v319 = v315<<(uint(int32(24))%32) + v166
	v320 = v313
	v321 = v314
	goto L26
L28:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+4)))
	v313 = v309 + v311
	v314 = v310
	goto L27
L29:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+5)))
	v309 = v305<<(uint(int32(8))%32) + v303
	v310 = v304
	goto L28
L30:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+6)))
	v303 = v299<<(uint(int32(16))%32) + v297
	v304 = v298
	goto L29
L31:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+7)))
	v297 = v293<<(uint(int32(24))%32) + v167
	v298 = v292
	goto L30
L32:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+8)))
	v292 = v288<<(uint(int32(8))%32) + v287
	goto L31
L33:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+9)))
	v287 = v283<<(uint(int32(16))%32) + v282
	goto L32
L34:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+10)))
	v282 = v278<<(uint(int32(24))%32) + v168
	goto L33
L35:
	;
	if base.Ui32(int32(11)) < base.Ui32(v60) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v60) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	v114 = v3
	v115 = v60
	v117 = v105
	v118 = v109
	v119 = v107
	goto L41
L39:
	;
	v163 = v3
	v164 = v60
	v166 = v105
	v167 = v109
	v168 = v107
	goto L40
L40:
	;
	switch v164 - int32(1) {
	case 0:
		v333 = v166
		v334 = v167
		v335 = v168
		goto L24
	case 1:
		v326 = v166
		v327 = v167
		v328 = v168
		goto L25
	case 2:
		v319 = v166
		v320 = v167
		v321 = v168
		goto L26
	case 3:
		v313 = v167
		v314 = v168
		goto L27
	case 4:
		v309 = v167
		v310 = v168
		goto L28
	case 5:
		v303 = v167
		v304 = v168
		goto L29
	case 6:
		v297 = v167
		v298 = v168
		goto L30
	case 7:
		v292 = v168
		goto L31
	case 8:
		v287 = v168
		goto L32
	case 9:
		v282 = v168
		goto L33
	case 10:
		goto L34
	default:
		v341 = v166
		v342 = v167
		v343 = v168
		goto L23
	}
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v122 = v121 + v118
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v114)+8))
	v126 = v125 + v119
	v128 = int32(4)
	v130 = v123 + v117 - v126 ^ base.I32_rotl(v126, v128)
	v134 = v122 - v130 ^ base.I32_rotl(v130, int32(6))
	v135 = v126 + v122
	v136 = v130 + v135
	v137 = v134 + v136
	v141 = v135 - v134 ^ base.I32_rotl(v134, int32(8))
	v145 = v136 - v141 ^ base.I32_rotl(v141, int32(16))
	v149 = v137 - v145 ^ base.I32_rotl(v145, int32(19))
	v150 = v141 + v137
	v151 = v145 + v150
	v152 = v149 + v151
	v156 = v150 - v149 ^ base.I32_rotl(v149, v128)
	v157 = int32(12)
	v158 = v114 + v157
	v160 = v115 - v157
	if base.Ui32(int32(11)) < base.Ui32(v160) {
		v114 = v158
		v115 = v160
		v117 = v151
		v118 = v152
		v119 = v156
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v163 = v158
	v164 = v160
	v166 = v151
	v167 = v152
	v168 = v156
	goto L40
L43:
	;
	goto L42
L44:
	;
	v174 = v3
	v175 = v60
	v177 = v105
	v178 = v109
	v179 = v107
	goto L47
L45:
	;
	v223 = v3
	v224 = v60
	v226 = v105
	v227 = v109
	v228 = v107
	goto L46
L46:
	;
	switch v224 - int32(1) {
	case 0:
		v275 = v226
		goto L50
	case 1:
		v270 = v226
		goto L51
	case 2:
		goto L52
	case 3:
		v263 = v227
		goto L53
	case 4:
		v260 = v227
		goto L54
	case 5:
		v255 = v227
		goto L55
	case 6:
		goto L56
	case 7:
		v246 = v228
		goto L57
	case 8:
		v241 = v228
		goto L58
	case 9:
		v236 = v228
		goto L59
	case 10:
		goto L60
	default:
		v341 = v226
		v342 = v227
		v343 = v228
		goto L23
	}
L47:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
	v182 = v181 + v178
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v174)+8))
	v186 = v185 + v179
	v188 = int32(4)
	v190 = v183 + v177 - v186 ^ base.I32_rotl(v186, v188)
	v194 = v182 - v190 ^ base.I32_rotl(v190, int32(6))
	v195 = v186 + v182
	v196 = v190 + v195
	v197 = v194 + v196
	v201 = v195 - v194 ^ base.I32_rotl(v194, int32(8))
	v205 = v196 - v201 ^ base.I32_rotl(v201, int32(16))
	v209 = v197 - v205 ^ base.I32_rotl(v205, int32(19))
	v210 = v201 + v197
	v211 = v205 + v210
	v212 = v209 + v211
	v216 = v210 - v209 ^ base.I32_rotl(v209, v188)
	v217 = int32(12)
	v218 = v174 + v217
	v220 = v175 - v217
	if base.Ui32(int32(11)) < base.Ui32(v220) {
		v174 = v218
		v175 = v220
		v177 = v211
		v178 = v212
		v179 = v216
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v223 = v218
	v224 = v220
	v226 = v211
	v227 = v212
	v228 = v216
	goto L46
L49:
	;
	goto L48
L50:
	;
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223))))
	v341 = v275 + v276
	v342 = v227
	v343 = v228
	goto L23
L51:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	v275 = v271<<(uint(int32(8))%32) + v270
	goto L50
L52:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+2)))
	v270 = v266<<(uint(int32(16))%32) + v226
	goto L51
L53:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v341 = v264 + v226
	v342 = v263
	v343 = v228
	goto L23
L54:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+4)))
	v263 = v260 + v261
	goto L53
L55:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+5)))
	v260 = v256<<(uint(int32(8))%32) + v255
	goto L54
L56:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+6)))
	v255 = v251<<(uint(int32(16))%32) + v227
	goto L55
L57:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v223)+4))
	v341 = v247 + v226
	v342 = v249 + v227
	v343 = v246
	goto L23
L58:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+8)))
	v246 = v242<<(uint(int32(8))%32) + v241
	goto L57
L59:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+9)))
	v241 = v237<<(uint(int32(16))%32) + v236
	goto L58
L60:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+10)))
	v236 = v232<<(uint(int32(24))%32) + v228
	goto L59
L61:
	;
	return int32(0)
L62:
	;
	return v378
}
func F_hashoidvectorextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
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
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_valid_oidvector(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = v3 + int32(24)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		v12 = v10 << (uint(int32(2)) % 32)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v20 = v12 - int32(1636608432)
		if v14 == int64(0) {
			v57 = v20
			v59 = v20
			v61 = v20
		} else {
			v24 = v20 + base.I32_wrap_i64(v14)
			v25 = v24 + v20
			v29 = int32(4)
			v31 = base.I32_wrap_i64(int64(base.Ui64(v14)>>(uint(int64(32))%64))) ^ base.I32_rotl(v20, v29)
			v35 = v24 - v31 ^ base.I32_rotl(v31, int32(6))
			v39 = v25 - v35 ^ base.I32_rotl(v35, int32(8))
			v40 = v31 + v25
			v41 = v35 + v40
			v42 = v39 + v41
			v46 = v40 - v39 ^ base.I32_rotl(v39, int32(16))
			v50 = v41 - v46 ^ base.I32_rotl(v46, int32(19))
			v55 = v46 + v42
			v57 = v55
			v59 = v42 - v50 ^ base.I32_rotl(v50, v29)
			v61 = v50 + v55
		}
		if v9&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v12) {
				v66 = v9
				v67 = v12
				v69 = v57
				v70 = v61
				v71 = v59
				for {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
					v74 = v73 + v70
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
					v78 = v77 + v71
					v80 = int32(4)
					v82 = v75 + v69 - v78 ^ base.I32_rotl(v78, v80)
					v86 = v74 - v82 ^ base.I32_rotl(v82, int32(6))
					v87 = v78 + v74
					v88 = v82 + v87
					v89 = v86 + v88
					v93 = v87 - v86 ^ base.I32_rotl(v86, int32(8))
					v97 = v88 - v93 ^ base.I32_rotl(v93, int32(16))
					v101 = v89 - v97 ^ base.I32_rotl(v97, int32(19))
					v102 = v93 + v89
					v103 = v97 + v102
					v104 = v101 + v103
					v108 = v102 - v101 ^ base.I32_rotl(v101, v80)
					v109 = int32(12)
					v110 = v66 + v109
					v112 = v67 - v109
					if base.Ui32(int32(11)) < base.Ui32(v112) {
						v66 = v110
						v67 = v112
						v69 = v103
						v70 = v104
						v71 = v108
						continue
					} else {
						break
					}
					break
				}
				v115 = v110
				v116 = v112
				v118 = v103
				v119 = v104
				v120 = v108
			} else {
				v115 = v9
				v116 = v12
				v118 = v57
				v119 = v61
				v120 = v59
			}
			switch v116 - int32(1) {
			case 0:
				v285 = v118
				v286 = v119
				v287 = v120
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 1:
				v278 = v118
				v279 = v119
				v280 = v120
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 2:
				v271 = v118
				v272 = v119
				v273 = v120
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 3:
				v265 = v119
				v266 = v120
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
				v271 = v267<<(uint(int32(24))%32) + v118
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 4:
				v261 = v119
				v262 = v120
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
				v265 = v261 + v263
				v266 = v262
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
				v271 = v267<<(uint(int32(24))%32) + v118
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 5:
				v255 = v119
				v256 = v120
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)))
				v261 = v257<<(uint(int32(8))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
				v265 = v261 + v263
				v266 = v262
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
				v271 = v267<<(uint(int32(24))%32) + v118
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 6:
				v249 = v119
				v250 = v120
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+6)))
				v255 = v251<<(uint(int32(16))%32) + v249
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)))
				v261 = v257<<(uint(int32(8))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
				v265 = v261 + v263
				v266 = v262
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
				v271 = v267<<(uint(int32(24))%32) + v118
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 7:
				v244 = v120
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+7)))
				v249 = v245<<(uint(int32(24))%32) + v119
				v250 = v244
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+6)))
				v255 = v251<<(uint(int32(16))%32) + v249
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)))
				v261 = v257<<(uint(int32(8))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
				v265 = v261 + v263
				v266 = v262
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
				v271 = v267<<(uint(int32(24))%32) + v118
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 8:
				v239 = v120
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+8)))
				v244 = v240<<(uint(int32(8))%32) + v239
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+7)))
				v249 = v245<<(uint(int32(24))%32) + v119
				v250 = v244
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+6)))
				v255 = v251<<(uint(int32(16))%32) + v249
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)))
				v261 = v257<<(uint(int32(8))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
				v265 = v261 + v263
				v266 = v262
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
				v271 = v267<<(uint(int32(24))%32) + v118
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 9:
				v234 = v120
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+9)))
				v239 = v235<<(uint(int32(16))%32) + v234
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+8)))
				v244 = v240<<(uint(int32(8))%32) + v239
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+7)))
				v249 = v245<<(uint(int32(24))%32) + v119
				v250 = v244
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+6)))
				v255 = v251<<(uint(int32(16))%32) + v249
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)))
				v261 = v257<<(uint(int32(8))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
				v265 = v261 + v263
				v266 = v262
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
				v271 = v267<<(uint(int32(24))%32) + v118
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			case 10:
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+10)))
				v234 = v230<<(uint(int32(24))%32) + v120
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+9)))
				v239 = v235<<(uint(int32(16))%32) + v234
				v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+8)))
				v244 = v240<<(uint(int32(8))%32) + v239
				v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+7)))
				v249 = v245<<(uint(int32(24))%32) + v119
				v250 = v244
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+6)))
				v255 = v251<<(uint(int32(16))%32) + v249
				v256 = v250
				v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+5)))
				v261 = v257<<(uint(int32(8))%32) + v255
				v262 = v256
				v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+4)))
				v265 = v261 + v263
				v266 = v262
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
				v271 = v267<<(uint(int32(24))%32) + v118
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
				v278 = v274<<(uint(int32(16))%32) + v271
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
				v285 = v281<<(uint(int32(8))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
				v293 = v285 + v288
				v294 = v286
				v295 = v287
			default:
				v293 = v118
				v294 = v119
				v295 = v120
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v12) {
				v126 = v9
				v127 = v12
				v129 = v57
				v130 = v61
				v131 = v59
				for {
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
					v134 = v133 + v130
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v126)))
					v137 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
					v138 = v137 + v131
					v140 = int32(4)
					v142 = v135 + v129 - v138 ^ base.I32_rotl(v138, v140)
					v146 = v134 - v142 ^ base.I32_rotl(v142, int32(6))
					v147 = v138 + v134
					v148 = v142 + v147
					v149 = v146 + v148
					v153 = v147 - v146 ^ base.I32_rotl(v146, int32(8))
					v157 = v148 - v153 ^ base.I32_rotl(v153, int32(16))
					v161 = v149 - v157 ^ base.I32_rotl(v157, int32(19))
					v162 = v153 + v149
					v163 = v157 + v162
					v164 = v161 + v163
					v168 = v162 - v161 ^ base.I32_rotl(v161, v140)
					v169 = int32(12)
					v170 = v126 + v169
					v172 = v127 - v169
					if base.Ui32(int32(11)) < base.Ui32(v172) {
						v126 = v170
						v127 = v172
						v129 = v163
						v130 = v164
						v131 = v168
						continue
					} else {
						break
					}
					break
				}
				v175 = v170
				v176 = v172
				v178 = v163
				v179 = v164
				v180 = v168
			} else {
				v175 = v9
				v176 = v12
				v178 = v57
				v179 = v61
				v180 = v59
			}
			switch v176 - int32(1) {
			case 0:
				v227 = v178
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
				v293 = v227 + v228
				v294 = v179
				v295 = v180
			case 1:
				v222 = v178
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
				v227 = v223<<(uint(int32(8))%32) + v222
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
				v293 = v227 + v228
				v294 = v179
				v295 = v180
			case 2:
				v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
				v222 = v218<<(uint(int32(16))%32) + v178
				v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
				v227 = v223<<(uint(int32(8))%32) + v222
				v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
				v293 = v227 + v228
				v294 = v179
				v295 = v180
			case 3:
				v215 = v179
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
				v293 = v216 + v178
				v294 = v215
				v295 = v180
			case 4:
				v212 = v179
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
				v215 = v212 + v213
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
				v293 = v216 + v178
				v294 = v215
				v295 = v180
			case 5:
				v207 = v179
				v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
				v212 = v208<<(uint(int32(8))%32) + v207
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
				v215 = v212 + v213
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
				v293 = v216 + v178
				v294 = v215
				v295 = v180
			case 6:
				v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+6)))
				v207 = v203<<(uint(int32(16))%32) + v179
				v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+5)))
				v212 = v208<<(uint(int32(8))%32) + v207
				v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+4)))
				v215 = v212 + v213
				v216 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
				v293 = v216 + v178
				v294 = v215
				v295 = v180
			case 7:
				v198 = v180
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
				v293 = v199 + v178
				v294 = v201 + v179
				v295 = v198
			case 8:
				v193 = v180
				v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)))
				v198 = v194<<(uint(int32(8))%32) + v193
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
				v293 = v199 + v178
				v294 = v201 + v179
				v295 = v198
			case 9:
				v188 = v180
				v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+9)))
				v193 = v189<<(uint(int32(16))%32) + v188
				v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)))
				v198 = v194<<(uint(int32(8))%32) + v193
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
				v293 = v199 + v178
				v294 = v201 + v179
				v295 = v198
			case 10:
				v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+10)))
				v188 = v184<<(uint(int32(24))%32) + v180
				v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+9)))
				v193 = v189<<(uint(int32(16))%32) + v188
				v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)))
				v198 = v194<<(uint(int32(8))%32) + v193
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
				v201 = *(*int32)(unsafe.Add(mBase, uint32(v175)+4))
				v293 = v199 + v178
				v294 = v201 + v179
				v295 = v198
			default:
				v293 = v178
				v294 = v179
				v295 = v180
			}
		}
		v298 = int32(14)
		v300 = v294 ^ v295 - base.I32_rotl(v294, v298)
		v304 = v300 ^ v293 - base.I32_rotl(v300, int32(11))
		v308 = v304 ^ v294 - base.I32_rotl(v304, int32(25))
		v312 = v308 ^ v300 - base.I32_rotl(v308, int32(16))
		v316 = v312 ^ v304 - base.I32_rotl(v312, int32(4))
		v320 = v316 ^ v308 - base.I32_rotl(v316, v298)
		v330 = F_Int64GetDatum(m, base.I64_extend_i32_u(v320)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v312^v320-base.I32_rotl(v320, int32(24))))
		mBase = m.M
		v331 = m.ExcPending
		if v331 != 0 {
			return int32(0)
		} else {
			return v330
		}
	}
}
func F_hashtranslatecmptype(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(l0 == int32(3))
}
func F_have_relevant_joinclause(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v10 = v9
	goto L3
L2:
	;
	v10 = v4
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+212))
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v264
L5:
	;
	v95 = int32(0)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
	if v96 != int32(1) {
		v264 = v95
		goto L4
	} else {
		goto L35
	}
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v14 = v12
	goto L8
L7:
	;
	v14 = int32(0)
	goto L8
L8:
	;
	v15 = base.B2i32(v14 < v10)
	if v14 < v10 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = l2
	goto L11
L10:
	;
	v16 = l1
	goto L11
L11:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+212))
	if v17 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	if v14 < v10 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v23 = l1
	goto L16
L15:
	;
	v23 = l2
	goto L16
L16:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v32 = v4
	goto L17
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v32<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v39 = int32(0)
	if v24 == v39 {
		v80 = v39
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L5
L19:
	;
	if v80 != 0 {
		v264 = int32(1)
		goto L4
	} else {
		goto L33
	}
L20:
	;
	goto L19
L21:
	;
	if v38 == int32(0) {
		v80 = v39
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v48 < v49 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v51 = v48
	goto L25
L24:
	;
	v51 = v49
	goto L25
L25:
	;
	if v51 <= int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v54 = int32(1)
	goto L28
L27:
	;
	v54 = v51
	goto L28
L28:
	;
	v55 = int32(8)
	v60 = int32(0)
	goto L29
L29:
	;
	v67 = v60 << (uint(int32(2)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v38+v55+v67)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+(v24+v55))))
	v72 = v69 & v71
	v74 = base.B2i32(v72 != int32(0))
	if v72 != 0 {
		v80 = v74
		goto L20
	} else {
		goto L31
	}
L30:
	;
	v80 = v74
	goto L20
L31:
	;
	v76 = v60 + int32(1)
	if v76 != v54 {
		v60 = v76
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v85 = v32 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v85 < v86 {
		v32 = v85
		goto L17
	} else {
		goto L34
	}
L34:
	;
	goto L18
L35:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+216)))
	if v99 != int32(1) {
		v264 = v95
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v104 = F_get_common_eclass_indexes(m, l0, v102, v103)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v264 = v258
	goto L4
L38:
	;
	return int32(0)
L39:
	;
	if v104 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	if int32(0) <= v164 {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v164 = base.I32_ctz(v150) | v151<<(uint(int32(5))%32)
	goto L40
L42:
	;
	v164 = int32(-2)
	goto L40
L43:
	;
	v117 = base.I32_div_s(int32(0), int32(32))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v118 <= v117 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v121 = v104 + int32(8)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121+v117<<(uint(int32(2))%32))))
	v128 = v125 & int32(-1)
	if v128 != 0 {
		v150 = v128
		v151 = v117
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v130 = v117 + int32(1)
	if v130 == v118 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v133 = v130
	goto L47
L47:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v121+v133<<(uint(int32(2))%32))))
	if v140 != 0 {
		v150 = v140
		v151 = v133
		goto L41
	} else {
		goto L49
	}
L48:
	;
	goto L42
L49:
	;
	v142 = v133 + int32(1)
	if v142 != v118 {
		v133 = v142
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v169 = v164
	goto L54
L52:
	;
	goto L53
L53:
	;
	v258 = int32(0)
	goto L37
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v169<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	if v180 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L53
L56:
	;
	v181 = int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v181 < v182 {
		v258 = v181
		goto L37
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v104 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	goto L58
L60:
	;
	if int32(0) <= v240 {
		v169 = v240
		goto L54
	} else {
		goto L71
	}
L61:
	;
	v240 = base.I32_ctz(v226) | v227<<(uint(int32(5))%32)
	goto L60
L62:
	;
	v240 = int32(-2)
	goto L60
L63:
	;
	v191 = v169 + int32(1)
	v193 = base.I32_div_s(v191, int32(32))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v194 <= v193 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v197 = v104 + int32(8)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v193<<(uint(int32(2))%32))))
	v204 = v201 & (int32(-1) << (uint(v191) % 32))
	if v204 != 0 {
		v226 = v204
		v227 = v193
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v206 = v193 + int32(1)
	if v206 == v194 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	v209 = v206
	goto L67
L67:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v197+v209<<(uint(int32(2))%32))))
	if v216 != 0 {
		v226 = v216
		v227 = v209
		goto L61
	} else {
		goto L69
	}
L68:
	;
	goto L62
L69:
	;
	v218 = v209 + int32(1)
	if v218 != v194 {
		v209 = v218
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L55
}
func F_heap2_desc(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v266 int32
	_ = v266
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(160)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+48)))
	if v19&int32(80) != int32(16) {
		v25 = v19 & int32(112)
		if v25 != int32(32) {
			switch int32(base.Ui32((v25+int32(-64))&int32(240)) >> (uint(int32(4)) % 32)) {
			case 0:
				v192 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v193
				*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v192
				F_appendStringInfo(m, l0, int32(490626), v15+int32(48))
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
					return
				} else {
					m.G0 = v15 + int32(160)
					return
				}
			case 1:
				v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+2)))
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+68)) = v202
				*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v201
				F_appendStringInfo(m, l0, int32(490692), v15-int32(-64))
				mBase = m.M
				v209 = m.ExcPending
				if v209 != 0 {
					return
				} else {
					if base.I32_extend8_s(v19) < int32(0) {
						m.G0 = v15 + int32(160)
						return
					} else {
						v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+119)))
						if v214&int32(1) == int32(0) {
							m.G0 = v15 + int32(160)
							return
						} else {
							F_appendStringInfoString(m, l0, int32(526975))
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return
							} else {
								v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+2)))
								F_array_desc(m, l0, v18+int32(4), int32(2), v225, int32(243), int32(0))
								mBase = m.M
								v229 = m.ExcPending
								if v229 != 0 {
									return
								} else {
									m.G0 = v15 + int32(160)
									return
								}
							}
						}
					}
				}
			case 2:
				v230 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+100)) = v231
				*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v230
				F_appendStringInfo(m, l0, int32(708186), v15+int32(96))
				mBase = m.M
				v238 = m.ExcPending
				if v238 != 0 {
					return
				} else {
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+6)))
					F_infobits_desc(m, l0, v239, int32(116996))
					mBase = m.M
					v242 = m.ExcPending
					if v242 != 0 {
						return
					} else {
						v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+7)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v243
						F_appendStringInfo(m, l0, int32(490703), v15+int32(80))
						mBase = m.M
						v249 = m.ExcPending
						if v249 != 0 {
							return
						} else {
							m.G0 = v15 + int32(160)
							return
						}
					}
				}
			case 3:
				v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+30)))
				v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)))
				v252 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
				v253 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
				v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+32)))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = v254
				*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v253
				*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v252
				*(*int32)(unsafe.Add(mBase, uint32(v15)+140)) = v250 | v251<<(uint(int32(16))%32)
				F_appendStringInfo(m, l0, int32(37137), v15+int32(128))
				mBase = m.M
				v266 = m.ExcPending
				if v266 != 0 {
					return
				} else {
					v267 = *(*int64)(unsafe.Add(mBase, uint32(v18)+4))
					v268 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v268
					*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v267
					F_appendStringInfo(m, l0, int32(55981), v15+int32(112))
					mBase = m.M
					v275 = m.ExcPending
					if v275 != 0 {
						return
					} else {
						m.G0 = v15 + int32(160)
						return
					}
				}
			default:
				m.G0 = v15 + int32(160)
				return
			}
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
			if v31&int32(8) != 0 {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+2))
				*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v34
				F_appendStringInfo(m, l0, int32(56013), v15+int32(32))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
					v42 = v41
					if v42&int32(2) != 0 {
						v45 = int32(84)
					} else {
						v45 = int32(70)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v45
					F_appendStringInfo(m, l0, int32(483042), v15+int32(16))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+119)))
						if v53 != int32(1) {
							m.G0 = v15 + int32(160)
							return
						} else {
							v56 = int32(0)
							v58 = v15 + int32(156)
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
							if v61 < v56 {
								v83 = v56
								v86 = v83
							} else {
								v67 = v60 + int32(76)
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
								if v68 != int32(1) {
									v83 = v56
									v86 = v83
								} else {
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+43)))
									if v71 == int32(0) {
										if v58 == int32(0) {
											v83 = v56
											v86 = v83
										} else {
											v76 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v58))) = v76
											v86 = v76
										}
									} else {
										if v58 != 0 {
											v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+48)))
											*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79
										} else {
										}
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
										v83 = v81
										v86 = v83
									}
								}
							}
							v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
							if v87&int32(16) == int32(0) {
								v99 = v86
								v100 = int32(0)
								v101 = v3
							} else {
								v94 = v86 + int32(4)
								v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
								v99 = v94 + v95*int32(12)
								v100 = v95
								v101 = v94
							}
							if v87&int32(32) == int32(0) {
								v113 = v99
								v114 = int32(0)
								v115 = v3
							} else {
								v107 = int32(2)
								v108 = v99 + v107
								v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
								v113 = v108 + v109<<(uint(v107)%32)
								v114 = v109
								v115 = v108
							}
							if v87&int32(64) == int32(0) {
								v128 = v113
								v129 = int32(0)
								v130 = v3
							} else {
								v123 = v113 + int32(2)
								v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
								v128 = v123 + v124<<(uint(int32(1))%32)
								v129 = v124
								v130 = v123
							}
							if int32(0) <= base.I32_extend8_s(v87) {
								v140 = v128
								v141 = int32(0)
								v142 = v3
							} else {
								v135 = v128 + int32(2)
								v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
								v140 = v135 + v136<<(uint(int32(1))%32)
								v141 = v136
								v142 = v135
							}
							*(*int32)(unsafe.Add(mBase, uint32(v15))) = v100
							*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v140
							*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v114
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v129
							*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v141
							F_appendStringInfo(m, l0, int32(56904), v15)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return
							} else {
								if v100 != 0 {
									F_appendStringInfoString(m, l0, int32(527053))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										F_array_desc(m, l0, v101, int32(12), v100, int32(241), v15+int32(152))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return
										} else {
											if v114 != 0 {
												F_appendStringInfoString(m, l0, int32(527234))
												mBase = m.M
												v162 = m.ExcPending
												if v162 != 0 {
													return
												} else {
													F_array_desc(m, l0, v115, int32(4), v114, int32(242), int32(0))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return
													} else {
														if v129 != 0 {
															F_appendStringInfoString(m, l0, int32(527258))
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return
															} else {
																F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
																mBase = m.M
																v175 = m.ExcPending
																if v175 != 0 {
																	return
																} else {
																	if v141 == int32(0) {
																		m.G0 = v15 + int32(160)
																		return
																	} else {
																		F_appendStringInfoString(m, l0, int32(527248))
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
																			return
																		} else {
																			F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				m.G0 = v15 + int32(160)
																				return
																			}
																		}
																	}
																}
															}
														} else {
															if v141 == int32(0) {
																m.G0 = v15 + int32(160)
																return
															} else {
																F_appendStringInfoString(m, l0, int32(527248))
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return
																} else {
																	F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		m.G0 = v15 + int32(160)
																		return
																	}
																}
															}
														}
													}
												}
											} else {
												if v129 != 0 {
													F_appendStringInfoString(m, l0, int32(527258))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return
													} else {
														F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return
														} else {
															if v141 == int32(0) {
																m.G0 = v15 + int32(160)
																return
															} else {
																F_appendStringInfoString(m, l0, int32(527248))
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return
																} else {
																	F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		m.G0 = v15 + int32(160)
																		return
																	}
																}
															}
														}
													}
												} else {
													if v141 == int32(0) {
														m.G0 = v15 + int32(160)
														return
													} else {
														F_appendStringInfoString(m, l0, int32(527248))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return
														} else {
															F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return
															} else {
																m.G0 = v15 + int32(160)
																return
															}
														}
													}
												}
											}
										}
									}
								} else {
									if v114 != 0 {
										F_appendStringInfoString(m, l0, int32(527234))
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											F_array_desc(m, l0, v115, int32(4), v114, int32(242), int32(0))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return
											} else {
												if v129 != 0 {
													F_appendStringInfoString(m, l0, int32(527258))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return
													} else {
														F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return
														} else {
															if v141 == int32(0) {
																m.G0 = v15 + int32(160)
																return
															} else {
																F_appendStringInfoString(m, l0, int32(527248))
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return
																} else {
																	F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		m.G0 = v15 + int32(160)
																		return
																	}
																}
															}
														}
													}
												} else {
													if v141 == int32(0) {
														m.G0 = v15 + int32(160)
														return
													} else {
														F_appendStringInfoString(m, l0, int32(527248))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return
														} else {
															F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return
															} else {
																m.G0 = v15 + int32(160)
																return
															}
														}
													}
												}
											}
										}
									} else {
										if v129 != 0 {
											F_appendStringInfoString(m, l0, int32(527258))
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return
											} else {
												F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													if v141 == int32(0) {
														m.G0 = v15 + int32(160)
														return
													} else {
														F_appendStringInfoString(m, l0, int32(527248))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return
														} else {
															F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return
															} else {
																m.G0 = v15 + int32(160)
																return
															}
														}
													}
												}
											}
										} else {
											if v141 == int32(0) {
												m.G0 = v15 + int32(160)
												return
											} else {
												F_appendStringInfoString(m, l0, int32(527248))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return
												} else {
													F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return
													} else {
														m.G0 = v15 + int32(160)
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
			} else {
				v42 = v31
				if v42&int32(2) != 0 {
					v45 = int32(84)
				} else {
					v45 = int32(70)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v45
				F_appendStringInfo(m, l0, int32(483042), v15+int32(16))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+119)))
					if v53 != int32(1) {
						m.G0 = v15 + int32(160)
						return
					} else {
						v56 = int32(0)
						v58 = v15 + int32(156)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
						if v61 < v56 {
							v83 = v56
							v86 = v83
						} else {
							v67 = v60 + int32(76)
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
							if v68 != int32(1) {
								v83 = v56
								v86 = v83
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+43)))
								if v71 == int32(0) {
									if v58 == int32(0) {
										v83 = v56
										v86 = v83
									} else {
										v76 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v58))) = v76
										v86 = v76
									}
								} else {
									if v58 != 0 {
										v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+48)))
										*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79
									} else {
									}
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
									v83 = v81
									v86 = v83
								}
							}
						}
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
						if v87&int32(16) == int32(0) {
							v99 = v86
							v100 = int32(0)
							v101 = v3
						} else {
							v94 = v86 + int32(4)
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
							v99 = v94 + v95*int32(12)
							v100 = v95
							v101 = v94
						}
						if v87&int32(32) == int32(0) {
							v113 = v99
							v114 = int32(0)
							v115 = v3
						} else {
							v107 = int32(2)
							v108 = v99 + v107
							v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
							v113 = v108 + v109<<(uint(v107)%32)
							v114 = v109
							v115 = v108
						}
						if v87&int32(64) == int32(0) {
							v128 = v113
							v129 = int32(0)
							v130 = v3
						} else {
							v123 = v113 + int32(2)
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
							v128 = v123 + v124<<(uint(int32(1))%32)
							v129 = v124
							v130 = v123
						}
						if int32(0) <= base.I32_extend8_s(v87) {
							v140 = v128
							v141 = int32(0)
							v142 = v3
						} else {
							v135 = v128 + int32(2)
							v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
							v140 = v135 + v136<<(uint(int32(1))%32)
							v141 = v136
							v142 = v135
						}
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v100
						*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v140
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v114
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v141
						F_appendStringInfo(m, l0, int32(56904), v15)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							if v100 != 0 {
								F_appendStringInfoString(m, l0, int32(527053))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									F_array_desc(m, l0, v101, int32(12), v100, int32(241), v15+int32(152))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return
									} else {
										if v114 != 0 {
											F_appendStringInfoString(m, l0, int32(527234))
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												F_array_desc(m, l0, v115, int32(4), v114, int32(242), int32(0))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return
												} else {
													if v129 != 0 {
														F_appendStringInfoString(m, l0, int32(527258))
														mBase = m.M
														v170 = m.ExcPending
														if v170 != 0 {
															return
														} else {
															F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
															mBase = m.M
															v175 = m.ExcPending
															if v175 != 0 {
																return
															} else {
																if v141 == int32(0) {
																	m.G0 = v15 + int32(160)
																	return
																} else {
																	F_appendStringInfoString(m, l0, int32(527248))
																	mBase = m.M
																	v180 = m.ExcPending
																	if v180 != 0 {
																		return
																	} else {
																		F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
																			return
																		} else {
																			m.G0 = v15 + int32(160)
																			return
																		}
																	}
																}
															}
														}
													} else {
														if v141 == int32(0) {
															m.G0 = v15 + int32(160)
															return
														} else {
															F_appendStringInfoString(m, l0, int32(527248))
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return
															} else {
																F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(160)
																	return
																}
															}
														}
													}
												}
											}
										} else {
											if v129 != 0 {
												F_appendStringInfoString(m, l0, int32(527258))
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return
												} else {
													F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return
													} else {
														if v141 == int32(0) {
															m.G0 = v15 + int32(160)
															return
														} else {
															F_appendStringInfoString(m, l0, int32(527248))
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return
															} else {
																F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(160)
																	return
																}
															}
														}
													}
												}
											} else {
												if v141 == int32(0) {
													m.G0 = v15 + int32(160)
													return
												} else {
													F_appendStringInfoString(m, l0, int32(527248))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return
													} else {
														F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return
														} else {
															m.G0 = v15 + int32(160)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								if v114 != 0 {
									F_appendStringInfoString(m, l0, int32(527234))
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										F_array_desc(m, l0, v115, int32(4), v114, int32(242), int32(0))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											if v129 != 0 {
												F_appendStringInfoString(m, l0, int32(527258))
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return
												} else {
													F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return
													} else {
														if v141 == int32(0) {
															m.G0 = v15 + int32(160)
															return
														} else {
															F_appendStringInfoString(m, l0, int32(527248))
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return
															} else {
																F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(160)
																	return
																}
															}
														}
													}
												}
											} else {
												if v141 == int32(0) {
													m.G0 = v15 + int32(160)
													return
												} else {
													F_appendStringInfoString(m, l0, int32(527248))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return
													} else {
														F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return
														} else {
															m.G0 = v15 + int32(160)
															return
														}
													}
												}
											}
										}
									}
								} else {
									if v129 != 0 {
										F_appendStringInfoString(m, l0, int32(527258))
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return
										} else {
											F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return
											} else {
												if v141 == int32(0) {
													m.G0 = v15 + int32(160)
													return
												} else {
													F_appendStringInfoString(m, l0, int32(527248))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return
													} else {
														F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return
														} else {
															m.G0 = v15 + int32(160)
															return
														}
													}
												}
											}
										}
									} else {
										if v141 == int32(0) {
											m.G0 = v15 + int32(160)
											return
										} else {
											F_appendStringInfoString(m, l0, int32(527248))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return
											} else {
												F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
													return
												} else {
													m.G0 = v15 + int32(160)
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
	} else {
		v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
		if v31&int32(8) != 0 {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+2))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v34
			F_appendStringInfo(m, l0, int32(56013), v15+int32(32))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
				v42 = v41
				if v42&int32(2) != 0 {
					v45 = int32(84)
				} else {
					v45 = int32(70)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v45
				F_appendStringInfo(m, l0, int32(483042), v15+int32(16))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+119)))
					if v53 != int32(1) {
						m.G0 = v15 + int32(160)
						return
					} else {
						v56 = int32(0)
						v58 = v15 + int32(156)
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
						if v61 < v56 {
							v83 = v56
							v86 = v83
						} else {
							v67 = v60 + int32(76)
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
							if v68 != int32(1) {
								v83 = v56
								v86 = v83
							} else {
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+43)))
								if v71 == int32(0) {
									if v58 == int32(0) {
										v83 = v56
										v86 = v83
									} else {
										v76 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v58))) = v76
										v86 = v76
									}
								} else {
									if v58 != 0 {
										v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+48)))
										*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79
									} else {
									}
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
									v83 = v81
									v86 = v83
								}
							}
						}
						v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
						if v87&int32(16) == int32(0) {
							v99 = v86
							v100 = int32(0)
							v101 = v3
						} else {
							v94 = v86 + int32(4)
							v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
							v99 = v94 + v95*int32(12)
							v100 = v95
							v101 = v94
						}
						if v87&int32(32) == int32(0) {
							v113 = v99
							v114 = int32(0)
							v115 = v3
						} else {
							v107 = int32(2)
							v108 = v99 + v107
							v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
							v113 = v108 + v109<<(uint(v107)%32)
							v114 = v109
							v115 = v108
						}
						if v87&int32(64) == int32(0) {
							v128 = v113
							v129 = int32(0)
							v130 = v3
						} else {
							v123 = v113 + int32(2)
							v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
							v128 = v123 + v124<<(uint(int32(1))%32)
							v129 = v124
							v130 = v123
						}
						if int32(0) <= base.I32_extend8_s(v87) {
							v140 = v128
							v141 = int32(0)
							v142 = v3
						} else {
							v135 = v128 + int32(2)
							v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
							v140 = v135 + v136<<(uint(int32(1))%32)
							v141 = v136
							v142 = v135
						}
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v100
						*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v140
						*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v114
						*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v129
						*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v141
						F_appendStringInfo(m, l0, int32(56904), v15)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							if v100 != 0 {
								F_appendStringInfoString(m, l0, int32(527053))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
									return
								} else {
									F_array_desc(m, l0, v101, int32(12), v100, int32(241), v15+int32(152))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return
									} else {
										if v114 != 0 {
											F_appendStringInfoString(m, l0, int32(527234))
											mBase = m.M
											v162 = m.ExcPending
											if v162 != 0 {
												return
											} else {
												F_array_desc(m, l0, v115, int32(4), v114, int32(242), int32(0))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return
												} else {
													if v129 != 0 {
														F_appendStringInfoString(m, l0, int32(527258))
														mBase = m.M
														v170 = m.ExcPending
														if v170 != 0 {
															return
														} else {
															F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
															mBase = m.M
															v175 = m.ExcPending
															if v175 != 0 {
																return
															} else {
																if v141 == int32(0) {
																	m.G0 = v15 + int32(160)
																	return
																} else {
																	F_appendStringInfoString(m, l0, int32(527248))
																	mBase = m.M
																	v180 = m.ExcPending
																	if v180 != 0 {
																		return
																	} else {
																		F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
																			return
																		} else {
																			m.G0 = v15 + int32(160)
																			return
																		}
																	}
																}
															}
														}
													} else {
														if v141 == int32(0) {
															m.G0 = v15 + int32(160)
															return
														} else {
															F_appendStringInfoString(m, l0, int32(527248))
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return
															} else {
																F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(160)
																	return
																}
															}
														}
													}
												}
											}
										} else {
											if v129 != 0 {
												F_appendStringInfoString(m, l0, int32(527258))
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return
												} else {
													F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return
													} else {
														if v141 == int32(0) {
															m.G0 = v15 + int32(160)
															return
														} else {
															F_appendStringInfoString(m, l0, int32(527248))
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return
															} else {
																F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(160)
																	return
																}
															}
														}
													}
												}
											} else {
												if v141 == int32(0) {
													m.G0 = v15 + int32(160)
													return
												} else {
													F_appendStringInfoString(m, l0, int32(527248))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return
													} else {
														F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return
														} else {
															m.G0 = v15 + int32(160)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								if v114 != 0 {
									F_appendStringInfoString(m, l0, int32(527234))
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										F_array_desc(m, l0, v115, int32(4), v114, int32(242), int32(0))
										mBase = m.M
										v167 = m.ExcPending
										if v167 != 0 {
											return
										} else {
											if v129 != 0 {
												F_appendStringInfoString(m, l0, int32(527258))
												mBase = m.M
												v170 = m.ExcPending
												if v170 != 0 {
													return
												} else {
													F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
													mBase = m.M
													v175 = m.ExcPending
													if v175 != 0 {
														return
													} else {
														if v141 == int32(0) {
															m.G0 = v15 + int32(160)
															return
														} else {
															F_appendStringInfoString(m, l0, int32(527248))
															mBase = m.M
															v180 = m.ExcPending
															if v180 != 0 {
																return
															} else {
																F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																mBase = m.M
																v185 = m.ExcPending
																if v185 != 0 {
																	return
																} else {
																	m.G0 = v15 + int32(160)
																	return
																}
															}
														}
													}
												}
											} else {
												if v141 == int32(0) {
													m.G0 = v15 + int32(160)
													return
												} else {
													F_appendStringInfoString(m, l0, int32(527248))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return
													} else {
														F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return
														} else {
															m.G0 = v15 + int32(160)
															return
														}
													}
												}
											}
										}
									}
								} else {
									if v129 != 0 {
										F_appendStringInfoString(m, l0, int32(527258))
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return
										} else {
											F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
											mBase = m.M
											v175 = m.ExcPending
											if v175 != 0 {
												return
											} else {
												if v141 == int32(0) {
													m.G0 = v15 + int32(160)
													return
												} else {
													F_appendStringInfoString(m, l0, int32(527248))
													mBase = m.M
													v180 = m.ExcPending
													if v180 != 0 {
														return
													} else {
														F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
														mBase = m.M
														v185 = m.ExcPending
														if v185 != 0 {
															return
														} else {
															m.G0 = v15 + int32(160)
															return
														}
													}
												}
											}
										}
									} else {
										if v141 == int32(0) {
											m.G0 = v15 + int32(160)
											return
										} else {
											F_appendStringInfoString(m, l0, int32(527248))
											mBase = m.M
											v180 = m.ExcPending
											if v180 != 0 {
												return
											} else {
												F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
												mBase = m.M
												v185 = m.ExcPending
												if v185 != 0 {
													return
												} else {
													m.G0 = v15 + int32(160)
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
		} else {
			v42 = v31
			if v42&int32(2) != 0 {
				v45 = int32(84)
			} else {
				v45 = int32(70)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v45
			F_appendStringInfo(m, l0, int32(483042), v15+int32(16))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+119)))
				if v53 != int32(1) {
					m.G0 = v15 + int32(160)
					return
				} else {
					v56 = int32(0)
					v58 = v15 + int32(156)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
					if v61 < v56 {
						v83 = v56
						v86 = v83
					} else {
						v67 = v60 + int32(76)
						v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
						if v68 != int32(1) {
							v83 = v56
							v86 = v83
						} else {
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+43)))
							if v71 == int32(0) {
								if v58 == int32(0) {
									v83 = v56
									v86 = v83
								} else {
									v76 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v58))) = v76
									v86 = v76
								}
							} else {
								if v58 != 0 {
									v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+48)))
									*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79
								} else {
								}
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v67)+44))
								v83 = v81
								v86 = v83
							}
						}
					}
					v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
					if v87&int32(16) == int32(0) {
						v99 = v86
						v100 = int32(0)
						v101 = v3
					} else {
						v94 = v86 + int32(4)
						v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
						v99 = v94 + v95*int32(12)
						v100 = v95
						v101 = v94
					}
					if v87&int32(32) == int32(0) {
						v113 = v99
						v114 = int32(0)
						v115 = v3
					} else {
						v107 = int32(2)
						v108 = v99 + v107
						v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99))))
						v113 = v108 + v109<<(uint(v107)%32)
						v114 = v109
						v115 = v108
					}
					if v87&int32(64) == int32(0) {
						v128 = v113
						v129 = int32(0)
						v130 = v3
					} else {
						v123 = v113 + int32(2)
						v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
						v128 = v123 + v124<<(uint(int32(1))%32)
						v129 = v124
						v130 = v123
					}
					if int32(0) <= base.I32_extend8_s(v87) {
						v140 = v128
						v141 = int32(0)
						v142 = v3
					} else {
						v135 = v128 + int32(2)
						v136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128))))
						v140 = v135 + v136<<(uint(int32(1))%32)
						v141 = v136
						v142 = v135
					}
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v100
					*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v140
					*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v114
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v129
					*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v141
					F_appendStringInfo(m, l0, int32(56904), v15)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return
					} else {
						if v100 != 0 {
							F_appendStringInfoString(m, l0, int32(527053))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return
							} else {
								F_array_desc(m, l0, v101, int32(12), v100, int32(241), v15+int32(152))
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return
								} else {
									if v114 != 0 {
										F_appendStringInfoString(m, l0, int32(527234))
										mBase = m.M
										v162 = m.ExcPending
										if v162 != 0 {
											return
										} else {
											F_array_desc(m, l0, v115, int32(4), v114, int32(242), int32(0))
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return
											} else {
												if v129 != 0 {
													F_appendStringInfoString(m, l0, int32(527258))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return
													} else {
														F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
														mBase = m.M
														v175 = m.ExcPending
														if v175 != 0 {
															return
														} else {
															if v141 == int32(0) {
																m.G0 = v15 + int32(160)
																return
															} else {
																F_appendStringInfoString(m, l0, int32(527248))
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return
																} else {
																	F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		m.G0 = v15 + int32(160)
																		return
																	}
																}
															}
														}
													}
												} else {
													if v141 == int32(0) {
														m.G0 = v15 + int32(160)
														return
													} else {
														F_appendStringInfoString(m, l0, int32(527248))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return
														} else {
															F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return
															} else {
																m.G0 = v15 + int32(160)
																return
															}
														}
													}
												}
											}
										}
									} else {
										if v129 != 0 {
											F_appendStringInfoString(m, l0, int32(527258))
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return
											} else {
												F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													if v141 == int32(0) {
														m.G0 = v15 + int32(160)
														return
													} else {
														F_appendStringInfoString(m, l0, int32(527248))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return
														} else {
															F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return
															} else {
																m.G0 = v15 + int32(160)
																return
															}
														}
													}
												}
											}
										} else {
											if v141 == int32(0) {
												m.G0 = v15 + int32(160)
												return
											} else {
												F_appendStringInfoString(m, l0, int32(527248))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return
												} else {
													F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return
													} else {
														m.G0 = v15 + int32(160)
														return
													}
												}
											}
										}
									}
								}
							}
						} else {
							if v114 != 0 {
								F_appendStringInfoString(m, l0, int32(527234))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return
								} else {
									F_array_desc(m, l0, v115, int32(4), v114, int32(242), int32(0))
									mBase = m.M
									v167 = m.ExcPending
									if v167 != 0 {
										return
									} else {
										if v129 != 0 {
											F_appendStringInfoString(m, l0, int32(527258))
											mBase = m.M
											v170 = m.ExcPending
											if v170 != 0 {
												return
											} else {
												F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
												mBase = m.M
												v175 = m.ExcPending
												if v175 != 0 {
													return
												} else {
													if v141 == int32(0) {
														m.G0 = v15 + int32(160)
														return
													} else {
														F_appendStringInfoString(m, l0, int32(527248))
														mBase = m.M
														v180 = m.ExcPending
														if v180 != 0 {
															return
														} else {
															F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
															mBase = m.M
															v185 = m.ExcPending
															if v185 != 0 {
																return
															} else {
																m.G0 = v15 + int32(160)
																return
															}
														}
													}
												}
											}
										} else {
											if v141 == int32(0) {
												m.G0 = v15 + int32(160)
												return
											} else {
												F_appendStringInfoString(m, l0, int32(527248))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return
												} else {
													F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return
													} else {
														m.G0 = v15 + int32(160)
														return
													}
												}
											}
										}
									}
								}
							} else {
								if v129 != 0 {
									F_appendStringInfoString(m, l0, int32(527258))
									mBase = m.M
									v170 = m.ExcPending
									if v170 != 0 {
										return
									} else {
										F_array_desc(m, l0, v130, int32(2), v129, int32(243), int32(0))
										mBase = m.M
										v175 = m.ExcPending
										if v175 != 0 {
											return
										} else {
											if v141 == int32(0) {
												m.G0 = v15 + int32(160)
												return
											} else {
												F_appendStringInfoString(m, l0, int32(527248))
												mBase = m.M
												v180 = m.ExcPending
												if v180 != 0 {
													return
												} else {
													F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
													mBase = m.M
													v185 = m.ExcPending
													if v185 != 0 {
														return
													} else {
														m.G0 = v15 + int32(160)
														return
													}
												}
											}
										}
									}
								} else {
									if v141 == int32(0) {
										m.G0 = v15 + int32(160)
										return
									} else {
										F_appendStringInfoString(m, l0, int32(527248))
										mBase = m.M
										v180 = m.ExcPending
										if v180 != 0 {
											return
										} else {
											F_array_desc(m, l0, v142, int32(2), v141, int32(243), int32(0))
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
												return
											} else {
												m.G0 = v15 + int32(160)
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
func F_hex_decode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_hex_decode_safe(m, l0, l1, l2, int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_hide_coercion_node(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v4 = m.G0
	v5 = int32(16)
	v6 = v4 - v5
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v9 - int32(15) {
	case 0, 13:
		v29 = v5
		*(*int32)(unsafe.Add(mBase, uint32(l0+v29))) = int32(2)
		m.G0 = v6 + int32(16)
		return
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v18
			F_errmsg_internal(m, int32(467408), v6)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errfinish(m, int32(480205), int32(826), int32(394775))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 12, 40:
		v29 = int32(20)
		*(*int32)(unsafe.Add(mBase, uint32(l0+v29))) = int32(2)
		m.G0 = v6 + int32(16)
		return
	case 14:
		v29 = int32(24)
		*(*int32)(unsafe.Add(mBase, uint32(l0+v29))) = int32(2)
		m.G0 = v6 + int32(16)
		return
	case 15, 21:
		v29 = int32(12)
		*(*int32)(unsafe.Add(mBase, uint32(l0+v29))) = int32(2)
		m.G0 = v6 + int32(16)
		return
	}
}
func F_histogram_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v135 float64
	_ = v135
	var v149 float64
	_ = v149
	var v153 int32
	_ = v153
	var v162 float64
	_ = v162
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(80)
	return v162
L2:
	;
	F_free_attstatsslot(m, v12+int32(44))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L12
	} else {
		goto L35
	}
L3:
	;
	v81 = int32(1)
	if v81 < v39-v81 {
		goto L24
	} else {
		goto L25
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l3
	goto L3
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
	v162 = float64(-1)
	goto L1
L6:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v19 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v61 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L19
	}
L8:
	;
	v31 = v14
	goto L10
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v20 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v35 = F_get_attstatsslot(m, v12+int32(44), v31, int32(2), int32(0), int32(1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v23 = F_get_func_leakproof(m, v20)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return float64(0)
L13:
	;
	if v23 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = v29
	goto L10
L15:
	;
	if v35 == int32(0) {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v39
	if v39 < int32(10) {
		v149 = float64(-1)
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v44 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+40)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v44)
	v48 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+26)) = uint16(v48)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)) = uint8(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v12)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l1
	if l4 == v44 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
	goto L3
L19:
	;
	if v61 == int32(0) {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v65 = F_get_func_name(m, v20)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v65
	F_errmsg_internal(m, int32(324567), v12)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(475691), int32(6242), int32(305252))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L5
L24:
	;
	v87 = int32(1)
	v92 = int32(0)
	goto L27
L25:
	;
	v133 = v39
	v135 = float64(0)
	goto L26
L26:
	;
	v149 = base.F64_div(v135, base.F64_convert_i32_s(v133-int32(2)))
	goto L2
L27:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96+v87<<(uint(int32(2))%32))))
	if l4 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v133 = v120
	v135 = base.F64_convert_i32_s(v117)
	goto L26
L29:
	;
	v103 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)) = uint8(v103)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = m.T0[v108].(func(*base.Module, int32) int32)(m, v12+int32(8))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L33
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v100
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v100
	goto L29
L33:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
	v117 = v92 + (v111^int32(-1))&base.B2i32(v109 != int32(0))
	v118 = int32(1)
	v119 = v87 + v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	if v119 < v120-v118 {
		v87 = v119
		v92 = v117
		goto L27
	} else {
		goto L34
	}
L34:
	;
	goto L28
L35:
	;
	v162 = v149
	goto L1
}
func F_hlparsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+60)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v6
	v19 = F_lookup_ts_config_cache(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v22 = F_lookup_ts_parser_cache(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = F_FunctionCall2Coll(m, v22+int32(28), int32(0), l3, l4)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v13)+44)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v19
	goto L5
L5:
	;
	v57 = F_FunctionCall3Coll(m, v22+int32(56), int32(0), v27, v11+int32(-8), v11+int32(-4))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v160 = F_FunctionCall1Coll(m, v22+int32(84), int32(0), v27)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L34
	}
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	if v57 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if int32(0) < v57 {
		goto L5
	} else {
		goto L33
	}
L9:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v89 = F_palloc(m, int32(16))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L18
	}
L10:
	;
	if v59 < int32(2047) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v66 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v66 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(421568), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(2047)
	F_errdetail(m, int32(612148), v13)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(479418), int32(575), int32(60542))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+8)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v89))) = v57
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if v94 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = int32(0)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v100
	v106 = F_LexizeExec(m, v11+int32(-56), v11+int32(-60))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v89
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v89
	goto L19
L23:
	;
	if v106 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v108 = v106
	goto L27
L25:
	;
	goto L26
L26:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_addHLParsedLex(m, l1, l2, v141, int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L32
	}
L27:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v118 + int32(1)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_addHLParsedLex(m, l1, l2, v122, v108)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v129 = F_LexizeExec(m, v11+int32(-56), v11+int32(-60))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v129 != 0 {
		v108 = v129
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	goto L8
L33:
	;
	goto L6
L34:
	;
	m.G0 = v13 - int32(-64)
	return
}
func F_hmac_init(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, v10)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v14 = F_palloc0(m, v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			if base.Ui32(v12) < base.Ui32(l2) {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				m.T0[v17].(func(*base.Module, int32, int32, int32))(m, v10, l1, l2)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
					m.T0[v20].(func(*base.Module, int32, int32))(m, v10, v14)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
						m.T0[v23].(func(*base.Module, int32))(m, v10)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							if v12 == int32(0) {
							} else {
								v30 = int32(0)
								if v12 != int32(1) {
									v37 = v30
									v41 = int32(0)
									for {
										v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v46 = v37 + v14
										v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
										v48 = int32(54)
										v49 = v47 ^ v48
										*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))) = uint8(v49)
										v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
										v54 = int32(92)
										v55 = v53 ^ v54
										*(*uint8)(unsafe.Add(mBase, uint32(v51+v37))) = uint8(v55)
										v58 = v37 | int32(1)
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v61 = v58 + v14
										v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
										v64 = v62 ^ v48
										*(*uint8)(unsafe.Add(mBase, uint32(v58+v59))) = uint8(v64)
										v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
										v70 = v68 ^ v54
										*(*uint8)(unsafe.Add(mBase, uint32(v66+v58))) = uint8(v70)
										v72 = int32(2)
										v73 = v37 + v72
										v75 = v41 + v72
										if v75 != v12&int32(-2) {
											v37 = v73
											v41 = v75
											continue
										} else {
											break
										}
										break
									}
									v79 = v73
								} else {
									v79 = v30
								}
								if v12&int32(1) == int32(0) {
								} else {
									v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v92 = v79 + v14
									v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
									v95 = v93 ^ int32(54)
									*(*uint8)(unsafe.Add(mBase, uint32(v90+v79))) = uint8(v95)
									v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
									v101 = v99 ^ int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v97+v79))) = uint8(v101)
								}
							}
							v113 = F___memset(m, v14, int32(0), v12)
							mBase = m.M
							F_pfree(m, v14)
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v117 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
								m.T0[v117].(func(*base.Module, int32, int32, int32))(m, v10, v116, v12)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			} else {
				if l2 != 0 {
					v26 = F__emscripten_memcpy_bulkmem(m, v14, l1, l2)
					mBase = m.M
				} else {
				}
				if v12 == int32(0) {
				} else {
					v30 = int32(0)
					if v12 != int32(1) {
						v37 = v30
						v41 = int32(0)
						for {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v46 = v37 + v14
							v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
							v48 = int32(54)
							v49 = v47 ^ v48
							*(*uint8)(unsafe.Add(mBase, uint32(v44+v37))) = uint8(v49)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
							v54 = int32(92)
							v55 = v53 ^ v54
							*(*uint8)(unsafe.Add(mBase, uint32(v51+v37))) = uint8(v55)
							v58 = v37 | int32(1)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v61 = v58 + v14
							v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
							v64 = v62 ^ v48
							*(*uint8)(unsafe.Add(mBase, uint32(v58+v59))) = uint8(v64)
							v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
							v70 = v68 ^ v54
							*(*uint8)(unsafe.Add(mBase, uint32(v66+v58))) = uint8(v70)
							v72 = int32(2)
							v73 = v37 + v72
							v75 = v41 + v72
							if v75 != v12&int32(-2) {
								v37 = v73
								v41 = v75
								continue
							} else {
								break
							}
							break
						}
						v79 = v73
					} else {
						v79 = v30
					}
					if v12&int32(1) == int32(0) {
					} else {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v92 = v79 + v14
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
						v95 = v93 ^ int32(54)
						*(*uint8)(unsafe.Add(mBase, uint32(v90+v79))) = uint8(v95)
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
						v101 = v99 ^ int32(92)
						*(*uint8)(unsafe.Add(mBase, uint32(v97+v79))) = uint8(v101)
					}
				}
				v113 = F___memset(m, v14, int32(0), v12)
				mBase = m.M
				F_pfree(m, v14)
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return
				} else {
					v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					m.T0[v117].(func(*base.Module, int32, int32, int32))(m, v10, v116, v12)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_hs_concat(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_concat(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_hungarian_ISO_8859_2_create_env(m *base.Module) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_create_env(m, int32(0), int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
