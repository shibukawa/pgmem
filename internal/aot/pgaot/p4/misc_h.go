package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HalfToFloat4(m *base.Module, l0 int32) float32 {
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v6 = l0 & int32(1023)
	v10 = l0 << (uint(int32(16)) % 32) & int32(-2147483648)
	v13 = int32(31)
	v14 = int32(base.Ui32(l0)>>(uint(int32(10))%32)) & v13
	if v14 != v13 {
		if v14 != 0 {
			v86 = v6
			v87 = v14<<(uint(int32(23))%32) + v10 + int32(939524096)
		} else {
			if v6 != 0 {
				if l0&int32(512) != 0 {
					v74 = v6 << (uint(int32(1)) % 32)
					v76 = int32(939524096)
				} else {
					if base.Ui32(int32(255)) < base.Ui32(v6) {
						v74 = v6 << (uint(int32(2)) % 32)
						v76 = int32(931135488)
					} else {
						if base.Ui32(int32(127)) < base.Ui32(v6) {
							v74 = v6 << (uint(int32(3)) % 32)
							v76 = int32(922746880)
						} else {
							if base.Ui32(int32(63)) < base.Ui32(v6) {
								v74 = v6 << (uint(int32(4)) % 32)
								v76 = int32(914358272)
							} else {
								if base.Ui32(int32(31)) < base.Ui32(v6) {
									v74 = v6 << (uint(int32(5)) % 32)
									v76 = int32(905969664)
								} else {
									if base.Ui32(int32(15)) < base.Ui32(v6) {
										v74 = v6 << (uint(int32(6)) % 32)
										v76 = int32(897581056)
									} else {
										if base.Ui32(int32(7)) < base.Ui32(v6) {
											v74 = v6 << (uint(int32(7)) % 32)
											v76 = int32(889192448)
										} else {
											if base.Ui32(int32(3)) < base.Ui32(v6) {
												v74 = v6 << (uint(int32(8)) % 32)
												v76 = int32(880803840)
											} else {
												v69 = base.B2i32(v6 == int32(1))
												if v6 == int32(1) {
													v70 = int32(1024)
												} else {
													v70 = v6 << (uint(int32(9)) % 32)
												}
												if v6 == int32(1) {
													v73 = int32(864026624)
												} else {
													v73 = int32(872415232)
												}
												v74 = v70
												v76 = v73
											}
										}
									}
								}
							}
						}
					}
				}
				v86 = v74 & int32(1022)
				v87 = v76 | v10
			} else {
				v86 = int32(0)
				v87 = v10
			}
		}
	} else {
		if v6 == int32(0) {
			v86 = int32(0)
			v87 = v10 | int32(2139095040)
		} else {
			v86 = v6
			v87 = v10 | int32(2143289344)
		}
	}
	return base.F32_reinterpret_i32(v87 | v86<<(uint(int32(13))%32))
}
func F_HandleFatalError(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = int32(1)
	goto L1
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[1]))
	v14 = int32(0)
	if base.B2i32(v13 == v14)|base.B2i32(v13 == int32(_a_F_HandleFatalError_0)) == v14 {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_HandleFatalError[2])))
	if v24&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[3]))
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L5:
	;
	v27 = int32(6)
	goto L7
L6:
	;
	v27 = int32(3)
	goto L7
L7:
	;
	if l0 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = v27
	goto L10
L9:
	;
	v29 = int32(3)
	goto L10
L10:
	;
	v31 = v13
	goto L11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31-int32(12))))
	if base.Ui32(v35) <= base.Ui32(int32(16)) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L4
L13:
	;
	F_signal_child(m, v31-int32(20), v29)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v42 != int32(_a_F_HandleFatalError_0) {
		v31 = v42
		goto L11
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
	goto L12
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[4])) = int32(2)
	goto L21
L20:
	;
	goto L21
L21:
	;
	v54 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_HandleFatalError[5])) = uint8(v54)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[6]))
	switch v57 - v54 {
	case 0, 1, 2, 3, 4:
		goto L26
	default:
		goto L22
	case 6, 7, 8, 9:
		goto L25
	}
L22:
	;
	v122 = *(*int64)(unsafe.Add(mBase, _c_F_HandleFatalError[7]))
	if v122 == int64(0) {
		goto L39
	} else {
		goto L40
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[6])) = v116
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v97
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[6]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100<<(uint(int32(2))%32))+uint32(_c_F_HandleFatalError[8])))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v105
	F_errmsg_internal(m, int32(_a_F_HandleFatalError_1), v6)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L16
	} else {
		goto L37
	}
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[9]))
	if v69 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v60 = int32(6)
	v63 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L16
	} else {
		goto L27
	}
L27:
	;
	if v63 == int32(0) {
		v116 = v60
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v96 = v60
	v97 = int32(_a_F_HandleFatalError_2)
	goto L24
L29:
	;
	F_FreeWaitEventSet(m, v69)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L16
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v72 = int32(_a_F_HandleFatalError_3)
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[9])) = v73
	v78 = F_CreateWaitEventSet(m, v73, int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L16
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[9])) = v78
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_HandleFatalError[10]))
	F_AddWaitEventToSet(m, v78, int32(1), int32(-1), v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L34
	}
L34:
	;
	v87 = int32(11)
	v90 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	if v90 == int32(0) {
		v116 = v87
		goto L23
	} else {
		goto L36
	}
L36:
	;
	v96 = v87
	v97 = int32(_a_F_HandleFatalError_4)
	goto L24
L37:
	;
	F_errfinish(m, int32(_a_F_HandleFatalError_5), int32(3320), int32(_a_F_HandleFatalError_6))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L16
	} else {
		goto L38
	}
L38:
	;
	v116 = v96
	goto L23
L39:
	;
	v126 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, _c_F_HandleFatalError[7])) = v126
	goto L41
L40:
	;
	goto L41
L41:
	;
	m.G0 = v6 + int32(16)
	return
}
func F_HoldingBufferPinThatDelaysRecovery(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	v1 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+76))
	if v6 < v1 {
		v35 = v1
		return v35
	} else {
		v10 = v6 + int32(1)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[1]))
		if v12 != int32(-1) {
			v16 = v12 << (uint(int32(4)) % 32)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+uint32(_c_F_HoldingBufferPinThatDelaysRecovery[2])))
			if v19 == v10 {
				v29 = v16 + int32(_a_F_HoldingBufferPinThatDelaysRecovery_0)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				v35 = base.B2i32(int32(0) < v30)
				return v35
			} else {
				v23 = F_GetPrivateRefCountEntrySlow(m, v10, int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v23 == int32(0) {
						v35 = v1
					} else {
						v29 = v23
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
						v35 = base.B2i32(int32(0) < v30)
					}
					return v35
				}
			}
		} else {
			v23 = F_GetPrivateRefCountEntrySlow(m, v10, int32(0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					v35 = v1
				} else {
					v29 = v23
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
					v35 = base.B2i32(int32(0) < v30)
				}
				return v35
			}
		}
	}
}
func F_handle_sig_alarm(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v108 int64
	_ = v108
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var __phi132 int32
	_ = __phi132
	var v133 int32
	_ = v133
	var __phi133 int32
	_ = __phi133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v177 int64
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v198 int64
	_ = v198
	var v200 int32
	_ = v200
	var v205 int64
	_ = v205
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(_a_F_handle_sig_alarm_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[0])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[1]))
	v22 = base.AtomicRmwOr32(m, v5, int32(_a_F_handle_sig_alarm_1), v5)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v71 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[2])) = v71
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[3]))
	if v74 == v71 {
		goto L16
	} else {
		goto L17
	}
L2:
	;
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(1)
	v26 = int32(0)
	v29 = base.AtomicRmwOr32(m, v26, int32(_a_F_handle_sig_alarm_1), v26)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v30 == v26 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v33 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
	if v37 == v33 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v39 = m.G0
	v41 = v39 - int32(16)
	m.G0 = v41
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[5]))
	if v44 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v67 = F_pgmem_kill(m, v33, int32(23))
	mBase = m.M
	goto L2
L9:
	;
	m.G0 = v41 + int32(16)
	goto L1
L10:
	;
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+15)) = uint8(v47)
	goto L11
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[6]))
	v55 = F_write(m, v51, v41+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v55 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[7]))
	if v59 == int32(27) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L32
	} else {
		goto L44
	}
L16:
	;
	v217 = int32(_a_F_handle_sig_alarm_0)
	v219 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[0])) = v219 - int32(1)
	m.G0 = v9 + int32(16)
	return
L17:
	;
	v78 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[3])) = v78
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8]))
	if v81 <= v78 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v87 = m.G0
	v88 = int32(16)
	v89 = v87 - v88
	m.G0 = v89
	F_gettimeofday(m, v89)
	mBase = m.M
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
	v93 = int64(*(*int32)(unsafe.Add(mBase, uint32(v89)+8)))
	m.G0 = v89 + v88
	v101 = v93 + v92*int64(1000000) - int64(946684800000000)
	goto L19
L19:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8]))
	if v103 <= int32(0) {
		v205 = v101
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_schedule_alarm(m, v205)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L32
	} else {
		goto L43
	}
L21:
	;
	v108 = v101
	goto L22
L22:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[9]))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v113)+24))
	if v108 < v114 {
		v205 = v108
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v205 = v198
	goto L20
L24:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[9]))
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8]))
	if v119 <= int32(0) {
		goto L15
	} else {
		goto L25
	}
L25:
	;
	v122 = int32(0)
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[9]))
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+4)) = uint8(v122)
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8]))
	if int32(2) <= v129 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	__phi132 = v122
	__phi133 = int32(1)
	v132 = __phi132
	v133 = __phi133
	goto L29
L27:
	;
	goto L28
L28:
	;
	v155 = int32(_a_F_handle_sig_alarm_2)
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8]))
	v158 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8])) = v157 - v158
	*(*uint8)(unsafe.Add(mBase, uint32(v117)+5)) = uint8(v158)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	m.T0[v163].(func(*base.Module))(m)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v138 = int32(2)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v133<<(uint(v138)%32))+uint32(_c_F_handle_sig_alarm[9])))
	*(*int32)(unsafe.Add(mBase, uint32(v132<<(uint(v138)%32))+uint32(_c_F_handle_sig_alarm[9]))) = v142
	v145 = v133 + int32(1)
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8]))
	if v145 < v147 {
		__phi132 = v133
		__phi133 = v145
		v132 = __phi132
		v133 = __phi133
		goto L29
	} else {
		goto L31
	}
L30:
	;
	goto L28
L31:
	;
	goto L30
L32:
	;
	return
L33:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v117)+32))
	if int32(0) < v166 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	v172 = base.I64_extend_i32_u(v166) * int64(1000)
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v117)+24))
	v175 = v174 + v172
	if v175 < v108 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v184 = m.G0
	v185 = int32(16)
	v186 = v184 - v185
	m.G0 = v186
	F_gettimeofday(m, v186)
	mBase = m.M
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v186)))
	v190 = int64(*(*int32)(unsafe.Add(mBase, uint32(v186)+8)))
	m.G0 = v186 + v185
	v198 = v190 + v189*int64(1000000) - int64(946684800000000)
	goto L41
L37:
	;
	v177 = v172 + v108
	goto L39
L38:
	;
	v177 = v175
	goto L39
L39:
	;
	F_enable_timeout(m, v169, v108, v177, v166)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L32
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8]))
	if int32(0) < v200 {
		v108 = v198
		goto L22
	} else {
		goto L42
	}
L42:
	;
	goto L23
L43:
	;
	goto L16
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
	v233 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v233 - int32(1)
	F_errmsg_internal(m, int32(_a_F_handle_sig_alarm_3), v9)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_handle_sig_alarm_4), int32(143), int32(_a_F_handle_sig_alarm_5))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_has_parameter_privilege_id_name(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v15 = F_convert_any_priv_string(m, v12, int32(_a_F_has_parameter_privilege_id_name_0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				v17 = F_text_to_cstring(m, v7)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = F_pg_parameter_aclcheck(m, v17, v5, v15)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(v19 == int32(0)))
					}
				}
			}
		}
	}
}
func F_has_parameter_privilege_name_name(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v15 = F_convert_any_priv_string(m, v12, int32(_a_F_has_parameter_privilege_name_name_0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				v17 = F_get_role_oid_or_public(m, v5)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					v19 = F_text_to_cstring(m, v7)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						v21 = F_pg_parameter_aclcheck(m, v19, v17, v15)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v21 == int32(0)))
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 float64
	_ = v170
	var v171 int32
	_ = v171
	var v174 float64
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v240 int64
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int64
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 float64
	_ = v335
	var v337 float64
	_ = v337
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	v4 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v20 = F_RelationGetNumberOfBlocksInFork(m, l1, v4)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v20 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_estimate_rel_size(m, l0, int32(0), v15+int32(-4), v15+int32(-16), v15+int32(-24))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L58
	}
L6:
	;
	v35 = *(*float64)(unsafe.Add(mBase, uint32(v17)+48))
	v37 = F__hash_init(m, l1, v35, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[0]))
	v44 = int32(base.Ui32(v40)>>(uint(int32(3))%32)) & int32(_a_F_hashbuild_0)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+118)))
	if v48 == int32(116) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v51 = int32(_a_F_hashbuild_1)
	goto L10
L9:
	;
	v51 = int32(_a_F_hashbuild_2)
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if base.Ui32(v44) < base.Ui32(v52) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v54 = v44
	goto L13
L12:
	;
	v54 = v52
	goto L13
L13:
	;
	if base.Ui32(v54) <= base.Ui32(v37) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v57 = F_palloc0(m, int32(20))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v149 = v4
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v149
	v159 = int32(1)
	v160 = int32(0)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+140))
	v170 = m.T0[v169].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v159, v160, v159, v160, int32(-1), int32(122), v15+int32(-48), v160)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L29
	}
L17:
	;
	v59 = int32(1)
	v60 = v37 - v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = l1
	v63 = int32(-1)
	v66 = v37 + v59
	if v37&v66 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v73 = v63<<(uint(int32(32)-base.I32_clz(v66))%32) ^ v63
	goto L20
L19:
	;
	v73 = v37
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v73
	v76 = int32(base.Ui32(v73) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v76
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[0]))
	v80 = m.G0
	v82 = v80 - int32(32)
	m.G0 = v82
	v84 = int32(0)
	v86 = F_tuplesort_begin_common(m, v79, v84, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v88 = int32(_a_F_hashbuild_3)
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[1]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[1])) = v91
	v94 = F_palloc(m, int32(20))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_hashbuild[2])))
	if v97 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = int32(2054)
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v86)+60)) = v94
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+36)) = uint8(v122)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = int32(2055)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = int32(2056)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = int32(2059)
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(2060)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v94))) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[1])) = v89
	m.G0 = v82 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = v86
	v149 = v57
	goto L16
L24:
	;
	v102 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v102 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+16)) = int32(102)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v73
	F_errmsg_internal(m, int32(_a_F_hashbuild_4), v82)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(_a_F_hashbuild_5), int32(468), int32(_a_F_hashbuild_6))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L23
L29:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v17)+48)) = v170
	v174 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[3]))
	if v178 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v219 != 0 {
		goto L34
	} else {
		goto L35
	}
L31:
	;
	goto L30
L32:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_hashbuild[4])))
	if v182&int32(1) == int32(0) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v187 = int32(_a_F_hashbuild_7)
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5]))
	v190 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5])) = v189 + v190
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v193 + v190
	v197 = int32(0)
	v199 = int32(_a_F_hashbuild_8)
	v200 = base.AtomicRmwOr32(m, v197, v199, v197)
	*(*int64)(unsafe.Add(mBase, uint32(v178+int32(88))+232)) = base.I64_trunc_sat_f64_s(v174)
	v208 = base.AtomicRmwOr32(m, v197, v199, v197)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v209 + v190
	v215 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5])) = v215 - v190
	goto L31
L34:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	F_tuplesort_performsort(m, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v333 = F_palloc(m, int32(16))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L57
	}
L37:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v225 = F_tuplesort_getheaptuple(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v225 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v227 = v225
	v240 = int64(0)
	goto L42
L40:
	;
	goto L41
L41:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)))
	F_tuplesort_end(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L55
	}
L42:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	F__hash_doinsert(m, v241, v227, v220, int32(1))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[6]))
	if v246 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v251 = v240 + int64(1)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[3]))
	if v254 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L47
L49:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	v296 = F_tuplesort_getheaptuple(m, v295)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	goto L49
L51:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_hashbuild[4])))
	if v258&int32(1) == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v263 = int32(_a_F_hashbuild_7)
	v265 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5]))
	v266 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5])) = v265 + v266
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v269 + v266
	v273 = int32(0)
	v275 = int32(_a_F_hashbuild_8)
	v276 = base.AtomicRmwOr32(m, v273, v275, v273)
	*(*int64)(unsafe.Add(mBase, uint32(v254+int32(96))+232)) = v251
	v284 = base.AtomicRmwOr32(m, v273, v275, v273)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	*(*int32)(unsafe.Add(mBase, uint32(v254))) = v285 + v266
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5])) = v291 - v266
	goto L50
L53:
	;
	if v296 != 0 {
		v227 = v296
		v240 = v251
		goto L42
	} else {
		goto L54
	}
L54:
	;
	goto L43
L55:
	;
	F_pfree(m, v312)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L36
L57:
	;
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v17)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v333))) = v335
	v337 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v333)+8)) = v337
	m.G0 = v17 - int32(-64)
	return v333
L58:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v347 + int32(4)
	F_errmsg_internal(m, int32(_a_F_hashbuild_9), v17)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_hashbuild_10), int32(151), int32(_a_F_hashbuild_11))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
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
func F_hashinet(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
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
		return base.I64_extend_i32_u(v279 ^ v271 - base.I32_rotl(v279, int32(24)))
	}
}
func F_hashnameextended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_strlen(m, v3)
	mBase = m.M
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = v4 - int32(1636608432)
	if v5 == int64(0) {
		v48 = v11
		v50 = v11
		v52 = v11
	} else {
		v15 = v11 + base.I32_wrap_i64(v5)
		v16 = v15 + v11
		v20 = int32(4)
		v22 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) ^ base.I32_rotl(v11, v20)
		v26 = v15 - v22 ^ base.I32_rotl(v22, int32(6))
		v30 = v16 - v26 ^ base.I32_rotl(v26, int32(8))
		v31 = v16 + v22
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v33 + v37
		v48 = v46
		v50 = v33 - v41 ^ base.I32_rotl(v41, v20)
		v52 = v41 + v46
	}
	if v3&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v4) {
			v57 = v3
			v58 = v4
			v60 = v48
			v61 = v52
			v62 = v50
			for {
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
				v65 = v64 + v61
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
				v69 = v68 + v62
				v71 = int32(4)
				v73 = v66 + v60 - v69 ^ base.I32_rotl(v69, v71)
				v77 = v65 - v73 ^ base.I32_rotl(v73, int32(6))
				v78 = v69 + v65
				v79 = v73 + v78
				v80 = v77 + v79
				v84 = v78 - v77 ^ base.I32_rotl(v77, int32(8))
				v88 = v79 - v84 ^ base.I32_rotl(v84, int32(16))
				v92 = v80 - v88 ^ base.I32_rotl(v88, int32(19))
				v93 = v84 + v80
				v94 = v88 + v93
				v95 = v92 + v94
				v99 = v93 - v92 ^ base.I32_rotl(v92, v71)
				v100 = int32(12)
				v101 = v57 + v100
				v103 = v58 - v100
				if base.Ui32(int32(11)) < base.Ui32(v103) {
					v57 = v101
					v58 = v103
					v60 = v94
					v61 = v95
					v62 = v99
					continue
				} else {
					break
				}
				break
			}
			v106 = v101
			v107 = v103
			v109 = v94
			v110 = v95
			v111 = v99
		} else {
			v106 = v3
			v107 = v4
			v109 = v48
			v110 = v52
			v111 = v50
		}
		switch v107 - int32(1) {
		case 0:
			v276 = v109
			v277 = v110
			v278 = v111
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 1:
			v269 = v109
			v270 = v110
			v271 = v111
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 2:
			v262 = v109
			v263 = v110
			v264 = v111
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 3:
			v256 = v110
			v257 = v111
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+3)))
			v262 = v258<<(uint(int32(24))%32) + v109
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 4:
			v252 = v110
			v253 = v111
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+3)))
			v262 = v258<<(uint(int32(24))%32) + v109
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 5:
			v246 = v110
			v247 = v111
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+3)))
			v262 = v258<<(uint(int32(24))%32) + v109
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 6:
			v240 = v110
			v241 = v111
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+3)))
			v262 = v258<<(uint(int32(24))%32) + v109
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 7:
			v235 = v111
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+7)))
			v240 = v236<<(uint(int32(24))%32) + v110
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+3)))
			v262 = v258<<(uint(int32(24))%32) + v109
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 8:
			v230 = v111
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+7)))
			v240 = v236<<(uint(int32(24))%32) + v110
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+3)))
			v262 = v258<<(uint(int32(24))%32) + v109
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 9:
			v225 = v111
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+9)))
			v230 = v226<<(uint(int32(16))%32) + v225
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+7)))
			v240 = v236<<(uint(int32(24))%32) + v110
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+3)))
			v262 = v258<<(uint(int32(24))%32) + v109
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 10:
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+10)))
			v225 = v221<<(uint(int32(24))%32) + v111
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+9)))
			v230 = v226<<(uint(int32(16))%32) + v225
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+7)))
			v240 = v236<<(uint(int32(24))%32) + v110
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+3)))
			v262 = v258<<(uint(int32(24))%32) + v109
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		default:
			v284 = v109
			v285 = v110
			v286 = v111
		}
	} else {
		if base.Ui32(int32(12)) <= base.Ui32(v4) {
			v117 = v3
			v118 = v4
			v120 = v48
			v121 = v52
			v122 = v50
			for {
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
				v125 = v124 + v121
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
				v128 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
				v129 = v128 + v122
				v131 = int32(4)
				v133 = v126 + v120 - v129 ^ base.I32_rotl(v129, v131)
				v137 = v125 - v133 ^ base.I32_rotl(v133, int32(6))
				v138 = v129 + v125
				v139 = v133 + v138
				v140 = v137 + v139
				v144 = v138 - v137 ^ base.I32_rotl(v137, int32(8))
				v148 = v139 - v144 ^ base.I32_rotl(v144, int32(16))
				v152 = v140 - v148 ^ base.I32_rotl(v148, int32(19))
				v153 = v144 + v140
				v154 = v148 + v153
				v155 = v152 + v154
				v159 = v153 - v152 ^ base.I32_rotl(v152, v131)
				v160 = int32(12)
				v161 = v117 + v160
				v163 = v118 - v160
				if base.Ui32(int32(11)) < base.Ui32(v163) {
					v117 = v161
					v118 = v163
					v120 = v154
					v121 = v155
					v122 = v159
					continue
				} else {
					break
				}
				break
			}
			v166 = v161
			v167 = v163
			v169 = v154
			v170 = v155
			v171 = v159
		} else {
			v166 = v3
			v167 = v4
			v169 = v48
			v170 = v52
			v171 = v50
		}
		switch v167 - int32(1) {
		case 0:
			v218 = v169
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v284 = v218 + v219
			v285 = v170
			v286 = v171
		case 1:
			v213 = v169
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v218 = v214<<(uint(int32(8))%32) + v213
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v284 = v218 + v219
			v285 = v170
			v286 = v171
		case 2:
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+2)))
			v213 = v209<<(uint(int32(16))%32) + v169
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+1)))
			v218 = v214<<(uint(int32(8))%32) + v213
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166))))
			v284 = v218 + v219
			v285 = v170
			v286 = v171
		case 3:
			v206 = v170
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
			v284 = v207 + v169
			v285 = v206
			v286 = v171
		case 4:
			v203 = v170
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
			v284 = v207 + v169
			v285 = v206
			v286 = v171
		case 5:
			v198 = v170
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
			v284 = v207 + v169
			v285 = v206
			v286 = v171
		case 6:
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+6)))
			v198 = v194<<(uint(int32(16))%32) + v170
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+5)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
			v284 = v207 + v169
			v285 = v206
			v286 = v171
		case 7:
			v189 = v171
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
			v284 = v190 + v169
			v285 = v192 + v170
			v286 = v189
		case 8:
			v184 = v171
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
			v284 = v190 + v169
			v285 = v192 + v170
			v286 = v189
		case 9:
			v179 = v171
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
			v284 = v190 + v169
			v285 = v192 + v170
			v286 = v189
		case 10:
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+10)))
			v179 = v175<<(uint(int32(24))%32) + v171
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
			v284 = v190 + v169
			v285 = v192 + v170
			v286 = v189
		default:
			v284 = v169
			v285 = v170
			v286 = v171
		}
	}
	v289 = int32(14)
	v291 = v285 ^ v286 - base.I32_rotl(v285, v289)
	v295 = v291 ^ v284 - base.I32_rotl(v291, int32(11))
	v299 = v295 ^ v285 - base.I32_rotl(v295, int32(25))
	v303 = v299 ^ v291 - base.I32_rotl(v299, int32(16))
	v307 = v303 ^ v295 - base.I32_rotl(v303, int32(4))
	v311 = v307 ^ v299 - base.I32_rotl(v307, v289)
	return base.I64_extend_i32_u(v311)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v311^v303-base.I32_rotl(v311, int32(24)))
}
func F_hashoidvectorextended(m *base.Module, l0 int32) int64 {
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
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
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
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
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
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
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
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_check_valid_oidvector(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = v3 + int32(24)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+16))
		v12 = v10 << (uint(int32(2)) % 32)
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = v12 - int32(1636608432)
		if v13 == int64(0) {
			v56 = v19
			v58 = v19
			v60 = v19
		} else {
			v23 = v19 + base.I32_wrap_i64(v13)
			v24 = v23 + v19
			v28 = int32(4)
			v30 = base.I32_wrap_i64(int64(base.Ui64(v13)>>(uint(int64(32))%64))) ^ base.I32_rotl(v19, v28)
			v34 = v23 - v30 ^ base.I32_rotl(v30, int32(6))
			v38 = v24 - v34 ^ base.I32_rotl(v34, int32(8))
			v39 = v24 + v30
			v40 = v34 + v39
			v41 = v38 + v40
			v45 = v39 - v38 ^ base.I32_rotl(v38, int32(16))
			v49 = v40 - v45 ^ base.I32_rotl(v45, int32(19))
			v54 = v41 + v45
			v56 = v54
			v58 = v41 - v49 ^ base.I32_rotl(v49, v28)
			v60 = v49 + v54
		}
		if v9&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v12) {
				v65 = v9
				v66 = v12
				v68 = v56
				v69 = v60
				v70 = v58
				for {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
					v73 = v72 + v69
					v74 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
					v77 = v76 + v70
					v79 = int32(4)
					v81 = v74 + v68 - v77 ^ base.I32_rotl(v77, v79)
					v85 = v73 - v81 ^ base.I32_rotl(v81, int32(6))
					v86 = v77 + v73
					v87 = v81 + v86
					v88 = v85 + v87
					v92 = v86 - v85 ^ base.I32_rotl(v85, int32(8))
					v96 = v87 - v92 ^ base.I32_rotl(v92, int32(16))
					v100 = v88 - v96 ^ base.I32_rotl(v96, int32(19))
					v101 = v92 + v88
					v102 = v96 + v101
					v103 = v100 + v102
					v107 = v101 - v100 ^ base.I32_rotl(v100, v79)
					v108 = int32(12)
					v109 = v65 + v108
					v111 = v66 - v108
					if base.Ui32(int32(11)) < base.Ui32(v111) {
						v65 = v109
						v66 = v111
						v68 = v102
						v69 = v103
						v70 = v107
						continue
					} else {
						break
					}
					break
				}
				v114 = v109
				v115 = v111
				v117 = v102
				v118 = v103
				v119 = v107
			} else {
				v114 = v9
				v115 = v12
				v117 = v56
				v118 = v60
				v119 = v58
			}
			switch v115 - int32(1) {
			case 0:
				v284 = v117
				v285 = v118
				v286 = v119
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 1:
				v277 = v117
				v278 = v118
				v279 = v119
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 2:
				v270 = v117
				v271 = v118
				v272 = v119
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 3:
				v264 = v118
				v265 = v119
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
				v270 = v266<<(uint(int32(24))%32) + v117
				v271 = v264
				v272 = v265
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 4:
				v260 = v118
				v261 = v119
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
				v264 = v260 + v262
				v265 = v261
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
				v270 = v266<<(uint(int32(24))%32) + v117
				v271 = v264
				v272 = v265
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 5:
				v254 = v118
				v255 = v119
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
				v260 = v256<<(uint(int32(8))%32) + v254
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
				v264 = v260 + v262
				v265 = v261
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
				v270 = v266<<(uint(int32(24))%32) + v117
				v271 = v264
				v272 = v265
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 6:
				v248 = v118
				v249 = v119
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
				v254 = v250<<(uint(int32(16))%32) + v248
				v255 = v249
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
				v260 = v256<<(uint(int32(8))%32) + v254
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
				v264 = v260 + v262
				v265 = v261
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
				v270 = v266<<(uint(int32(24))%32) + v117
				v271 = v264
				v272 = v265
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 7:
				v243 = v119
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
				v248 = v244<<(uint(int32(24))%32) + v118
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
				v254 = v250<<(uint(int32(16))%32) + v248
				v255 = v249
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
				v260 = v256<<(uint(int32(8))%32) + v254
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
				v264 = v260 + v262
				v265 = v261
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
				v270 = v266<<(uint(int32(24))%32) + v117
				v271 = v264
				v272 = v265
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 8:
				v238 = v119
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+8)))
				v243 = v239<<(uint(int32(8))%32) + v238
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
				v248 = v244<<(uint(int32(24))%32) + v118
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
				v254 = v250<<(uint(int32(16))%32) + v248
				v255 = v249
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
				v260 = v256<<(uint(int32(8))%32) + v254
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
				v264 = v260 + v262
				v265 = v261
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
				v270 = v266<<(uint(int32(24))%32) + v117
				v271 = v264
				v272 = v265
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 9:
				v233 = v119
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+9)))
				v238 = v234<<(uint(int32(16))%32) + v233
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+8)))
				v243 = v239<<(uint(int32(8))%32) + v238
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
				v248 = v244<<(uint(int32(24))%32) + v118
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
				v254 = v250<<(uint(int32(16))%32) + v248
				v255 = v249
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
				v260 = v256<<(uint(int32(8))%32) + v254
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
				v264 = v260 + v262
				v265 = v261
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
				v270 = v266<<(uint(int32(24))%32) + v117
				v271 = v264
				v272 = v265
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			case 10:
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+10)))
				v233 = v229<<(uint(int32(24))%32) + v119
				v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+9)))
				v238 = v234<<(uint(int32(16))%32) + v233
				v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+8)))
				v243 = v239<<(uint(int32(8))%32) + v238
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
				v248 = v244<<(uint(int32(24))%32) + v118
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
				v254 = v250<<(uint(int32(16))%32) + v248
				v255 = v249
				v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+5)))
				v260 = v256<<(uint(int32(8))%32) + v254
				v261 = v255
				v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
				v264 = v260 + v262
				v265 = v261
				v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
				v270 = v266<<(uint(int32(24))%32) + v117
				v271 = v264
				v272 = v265
				v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
				v277 = v273<<(uint(int32(16))%32) + v270
				v278 = v271
				v279 = v272
				v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
				v284 = v280<<(uint(int32(8))%32) + v277
				v285 = v278
				v286 = v279
				v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
				v292 = v284 + v287
				v293 = v285
				v294 = v286
			default:
				v292 = v117
				v293 = v118
				v294 = v119
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v12) {
				v125 = v9
				v126 = v12
				v128 = v56
				v129 = v60
				v130 = v58
				for {
					v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
					v133 = v132 + v129
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
					v137 = v136 + v130
					v139 = int32(4)
					v141 = v134 + v128 - v137 ^ base.I32_rotl(v137, v139)
					v145 = v133 - v141 ^ base.I32_rotl(v141, int32(6))
					v146 = v137 + v133
					v147 = v141 + v146
					v148 = v145 + v147
					v152 = v146 - v145 ^ base.I32_rotl(v145, int32(8))
					v156 = v147 - v152 ^ base.I32_rotl(v152, int32(16))
					v160 = v148 - v156 ^ base.I32_rotl(v156, int32(19))
					v161 = v152 + v148
					v162 = v156 + v161
					v163 = v160 + v162
					v167 = v161 - v160 ^ base.I32_rotl(v160, v139)
					v168 = int32(12)
					v169 = v125 + v168
					v171 = v126 - v168
					if base.Ui32(int32(11)) < base.Ui32(v171) {
						v125 = v169
						v126 = v171
						v128 = v162
						v129 = v163
						v130 = v167
						continue
					} else {
						break
					}
					break
				}
				v174 = v169
				v175 = v171
				v177 = v162
				v178 = v163
				v179 = v167
			} else {
				v174 = v9
				v175 = v12
				v177 = v56
				v178 = v60
				v179 = v58
			}
			switch v175 - int32(1) {
			case 0:
				v226 = v177
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v292 = v226 + v227
				v293 = v178
				v294 = v179
			case 1:
				v221 = v177
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v226 = v222<<(uint(int32(8))%32) + v221
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v292 = v226 + v227
				v293 = v178
				v294 = v179
			case 2:
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)))
				v221 = v217<<(uint(int32(16))%32) + v177
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)))
				v226 = v222<<(uint(int32(8))%32) + v221
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174))))
				v292 = v226 + v227
				v293 = v178
				v294 = v179
			case 3:
				v214 = v178
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
				v292 = v215 + v177
				v293 = v214
				v294 = v179
			case 4:
				v211 = v178
				v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v214 = v211 + v212
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
				v292 = v215 + v177
				v293 = v214
				v294 = v179
			case 5:
				v206 = v178
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
				v211 = v207<<(uint(int32(8))%32) + v206
				v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v214 = v211 + v212
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
				v292 = v215 + v177
				v293 = v214
				v294 = v179
			case 6:
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+6)))
				v206 = v202<<(uint(int32(16))%32) + v178
				v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+5)))
				v211 = v207<<(uint(int32(8))%32) + v206
				v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+4)))
				v214 = v211 + v212
				v215 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
				v292 = v215 + v177
				v293 = v214
				v294 = v179
			case 7:
				v197 = v179
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
				v200 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
				v292 = v198 + v177
				v293 = v200 + v178
				v294 = v197
			case 8:
				v192 = v179
				v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
				v197 = v193<<(uint(int32(8))%32) + v192
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
				v200 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
				v292 = v198 + v177
				v293 = v200 + v178
				v294 = v197
			case 9:
				v187 = v179
				v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+9)))
				v192 = v188<<(uint(int32(16))%32) + v187
				v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
				v197 = v193<<(uint(int32(8))%32) + v192
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
				v200 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
				v292 = v198 + v177
				v293 = v200 + v178
				v294 = v197
			case 10:
				v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+10)))
				v187 = v183<<(uint(int32(24))%32) + v179
				v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+9)))
				v192 = v188<<(uint(int32(16))%32) + v187
				v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+8)))
				v197 = v193<<(uint(int32(8))%32) + v192
				v198 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
				v200 = *(*int32)(unsafe.Add(mBase, uint32(v174)+4))
				v292 = v198 + v177
				v293 = v200 + v178
				v294 = v197
			default:
				v292 = v177
				v293 = v178
				v294 = v179
			}
		}
		v297 = int32(14)
		v299 = v293 ^ v294 - base.I32_rotl(v293, v297)
		v303 = v299 ^ v292 - base.I32_rotl(v299, int32(11))
		v307 = v303 ^ v293 - base.I32_rotl(v303, int32(25))
		v311 = v307 ^ v299 - base.I32_rotl(v307, int32(16))
		v315 = v311 ^ v303 - base.I32_rotl(v311, int32(4))
		v319 = v315 ^ v307 - base.I32_rotl(v315, v297)
		return base.I64_extend_i32_u(v319)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v319^v311-base.I32_rotl(v319, int32(24)))
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
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
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v241 int32
	_ = v241
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+228))
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)+228))
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
	v96 = int32(0)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+232)))
	if v97 != int32(1) {
		v264 = v96
		goto L4
	} else {
		goto L34
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+228))
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
	v31 = v4
	goto L17
L17:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v31<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+32))
	v39 = int32(0)
	if base.B2i32(v24 == v39)|base.B2i32(v38 == v39) != 0 {
		v84 = v39
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L5
L19:
	;
	if v84 != 0 {
		v264 = int32(1)
		goto L4
	} else {
		goto L32
	}
L20:
	;
	goto L19
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v49 < v50 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v52 = v49
	goto L24
L23:
	;
	v52 = v50
	goto L24
L24:
	;
	if v52 <= int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v55 = int32(1)
	goto L27
L26:
	;
	v55 = v52
	goto L27
L27:
	;
	v56 = int32(8)
	v61 = int32(0)
	goto L28
L28:
	;
	v68 = v61 << (uint(int32(2)) % 32)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v38+v56+v68)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v24+v56+v68)))
	v73 = v70 & v72
	v75 = base.B2i32(v73 != int32(0))
	if v73 != 0 {
		v84 = v75
		goto L20
	} else {
		goto L30
	}
L29:
	;
	v84 = v75
	goto L20
L30:
	;
	v77 = v61 + int32(1)
	if v77 != v55 {
		v61 = v77
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v86 = v31 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v86 < v87 {
		v31 = v86
		goto L17
	} else {
		goto L33
	}
L33:
	;
	goto L18
L34:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+232)))
	if v100 != int32(1) {
		v264 = v96
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v105 = F_get_common_eclass_indexes(m, l0, v103, v104)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v264 = v259
	goto L4
L37:
	;
	return int32(0)
L38:
	;
	if v105 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if int32(0) <= v165 {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v165 = base.I32_ctz(v151) | v152<<(uint(int32(5))%32)
	goto L39
L41:
	;
	v165 = int32(-2)
	goto L39
L42:
	;
	v116 = int32(0)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v119 <= v116 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v122 = v105 + int32(8)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	v129 = v126 & int32(-1)
	if v129 != 0 {
		v151 = v129
		v152 = v116
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v130 = int32(1)
	if v130 == v119 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v134 = v130
	goto L46
L46:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v122+v134<<(uint(int32(2))%32))))
	if v141 != 0 {
		v151 = v141
		v152 = v134
		goto L40
	} else {
		goto L48
	}
L47:
	;
	goto L41
L48:
	;
	v143 = v134 + int32(1)
	if v143 != v119 {
		v134 = v143
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v170 = v165
	goto L53
L51:
	;
	goto L52
L52:
	;
	v259 = int32(0)
	goto L36
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)+12))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v170<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	if v181 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	v182 = int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v182 < v183 {
		v259 = v182
		goto L36
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	if v105 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L57
L59:
	;
	if int32(0) <= v241 {
		v170 = v241
		goto L53
	} else {
		goto L70
	}
L60:
	;
	v241 = base.I32_ctz(v227) | v228<<(uint(int32(5))%32)
	goto L59
L61:
	;
	v241 = int32(-2)
	goto L59
L62:
	;
	v192 = v170 + int32(1)
	v194 = int32(base.Ui32(v192) >> (uint(int32(5)) % 32))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v195 <= v194 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v198 = v105 + int32(8)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v198+v194<<(uint(int32(2))%32))))
	v205 = v202 & (int32(-1) << (uint(v192) % 32))
	if v205 != 0 {
		v227 = v205
		v228 = v194
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v207 = v194 + int32(1)
	if v207 == v195 {
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v210 = v207
	goto L66
L66:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v198+v210<<(uint(int32(2))%32))))
	if v217 != 0 {
		v227 = v217
		v228 = v210
		goto L60
	} else {
		goto L68
	}
L67:
	;
	goto L61
L68:
	;
	v219 = v210 + int32(1)
	if v219 != v195 {
		v210 = v219
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	goto L54
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
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v278 int64
	_ = v278
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(176)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+64))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+48)))
	if v19&int32(80) != int32(16) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(176)
	return
L2:
	;
	switch v25 - int32(96) {
	case 0:
		goto L70
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
		goto L1
	case 16:
		goto L69
	default:
		goto L71
	}
L3:
	;
	v25 = v19 & int32(112)
	if v25 != int32(32) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18))))
	if v31&int32(8) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v18)+2))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v34
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_8), v15+int32(48))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v42 = v31
	goto L9
L9:
	;
	if v42&int32(2) != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18))))
	v42 = v41
	goto L9
L12:
	;
	v45 = int32(84)
	goto L14
L13:
	;
	v45 = int32(70)
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v45
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_9), v15+int32(32))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18))))
	if v52&int32(256) != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	if v52&int32(512) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+119)))
	if v67 != int32(1) {
		goto L1
	} else {
		goto L23
	}
L19:
	;
	v59 = int32(3)
	goto L21
L20:
	;
	v59 = int32(1)
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v59
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_6), v15+int32(16))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	v70 = int32(0)
	v72 = v15 + int32(172)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+72))
	if v75 < v70 {
		v97 = v70
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18))))
	if v101&int32(16) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L25:
	;
	v100 = v97
	goto L24
L26:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74+int32(0))+76)))
	if v80 != int32(1) {
		v97 = v70
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v84 = v74 + int32(76)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+43)))
	if v85 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v72 == int32(0) {
		v97 = v70
		goto L25
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v72 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v90 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v90
	v100 = v90
	goto L24
L32:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+48)))
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = v93
	goto L34
L33:
	;
	goto L34
L34:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v84)+44))
	v97 = v95
	goto L25
L35:
	;
	if v101&int32(32) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v113 = v100
	v114 = int32(0)
	v115 = v3
	goto L35
L37:
	;
	goto L38
L38:
	;
	v108 = v100 + int32(4)
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100))))
	v113 = v108 + v109*int32(12)
	v114 = v109
	v115 = v108
	goto L35
L39:
	;
	if v101&int32(64) != 0 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v127 = v113
	v128 = int32(0)
	v129 = v3
	goto L39
L41:
	;
	goto L42
L42:
	;
	v121 = int32(2)
	v122 = v113 + v121
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
	v127 = v122 + v123<<(uint(v121)%32)
	v128 = v123
	v129 = v122
	goto L39
L43:
	;
	v133 = v127 + int32(2)
	v134 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127))))
	v138 = v133 + v134<<(uint(int32(1))%32)
	v139 = v134
	v140 = v133
	goto L45
L44:
	;
	v138 = v127
	v139 = v3
	v140 = v3
	goto L45
L45:
	;
	if v101&int32(128) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v15)+168)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v153
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_10), v15)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L10
	} else {
		goto L50
	}
L47:
	;
	v152 = v138
	v153 = int32(0)
	v154 = v3
	goto L46
L48:
	;
	goto L49
L49:
	;
	v147 = v138 + int32(2)
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138))))
	v152 = v147 + v148<<(uint(int32(1))%32)
	v153 = v148
	v154 = v147
	goto L46
L50:
	;
	if v114 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_11))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L10
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v128 != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	F_array_desc(m, l0, v115, int32(12), v114, int32(249), v15+int32(168))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v139 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	F_array_desc(m, l0, v129, int32(4), v128, int32(250), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L10
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	if v153 == int32(0) {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	F_array_desc(m, l0, v140, int32(2), v139, int32(251), int32(0))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	F_array_desc(m, l0, v154, int32(2), v153, int32(251), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	goto L1
L69:
	;
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+30)))
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)))
	v263 = *(*int64)(unsafe.Add(mBase, uint32(v18)+16))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+32)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v15)+152)) = v264
	*(*int64)(unsafe.Add(mBase, uint32(v15)+144)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v15)+156)) = v261 | v262<<(uint(int32(16))%32)
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_3), v15+int32(144))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L10
	} else {
		goto L85
	}
L70:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+116)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = v241
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_0), v15+int32(112))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L82
	}
L71:
	;
	if v25 != int32(80) {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+2)))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v203
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_5), v15+int32(80))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L73
	}
L73:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v212&int32(32) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(3)
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_6), v15-int32(-64))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L10
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	if base.I32_extend8_s(v19) < int32(0) {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+119)))
	if v225&int32(1) == int32(0) {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_7))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+2)))
	F_array_desc(m, l0, v18+int32(4), int32(2), v236, int32(251), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L10
	} else {
		goto L81
	}
L81:
	;
	goto L1
L82:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+6)))
	F_infobits_desc(m, l0, v250, int32(_a_F_heap2_desc_1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+7)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = v254
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_2), v15+int32(96))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	goto L1
L85:
	;
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v18)+4))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+136)) = v279
	*(*int64)(unsafe.Add(mBase, uint32(v15)+128)) = v278
	F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_4), v15+int32(128))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	goto L1
}
func F_hemdist_3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	v5 = Fn14312(m, l0, l1, l2, int32(4))
	return v5
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
			F_errmsg_internal(m, int32(_a_F_hide_coercion_node_0), v6)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_hide_coercion_node_1), int32(831), int32(_a_F_hide_coercion_node_2))
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
func F_histogram_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) float64 {
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
	var v100 int64
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int64
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
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v14 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(112)
	return v162
L2:
	;
	F_free_attstatsslot(m, v12+int32(76))
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
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = l3
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
	v35 = F_get_attstatsslot(m, v12+int32(76), v31, int32(2), int32(0), int32(1))
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
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+64)) = uint8(v44)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+48)) = uint8(v44)
	v48 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+34)) = uint16(v48)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v44)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v12)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = l1
	if l4 == v44 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = l3
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
	F_errmsg_internal(m, int32(_a_F_histogram_selectivity_0), v12)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_histogram_selectivity_1), int32(_a_F_histogram_selectivity_2), int32(_a_F_histogram_selectivity_3))
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
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v96+v87<<(uint(int32(3))%32))))
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
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)) = uint8(v103)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v109 = m.T0[v108].(func(*base.Module, int32) int64)(m, v12+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L12
	} else {
		goto L33
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+40)) = v100
	goto L29
L31:
	;
	goto L32
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+56)) = v100
	goto L29
L33:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)))
	v117 = v92 + (v111^int32(-1))&base.B2i32(v109 != int64(0))
	v118 = int32(1)
	v119 = v87 + v118
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v48 int64
	_ = v48
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	v6 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v6
	v22 = F_lookup_ts_config_cache(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v25 = F_lookup_ts_parser_cache(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = F_FunctionCall2Coll(m, v25+int32(28), int32(0), base.I64_extend_i32_u(l3), base.I64_extend_i32_s(l4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+28)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = v34
	*(*int64)(unsafe.Add(mBase, uint32(v16)+44)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = int32(0)
	v48 = v32 & int64(4294967295)
	goto L5
L5:
	;
	v69 = F_FunctionCall3Coll(m, v25+int32(56), int32(0), v48, base.I64_extend_i32_u(v14+int32(-8)), base.I64_extend_i32_u(v14+int32(-4)))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v185 = F_FunctionCall1Coll(m, v25+int32(84), int32(0), v48)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L34
	}
L7:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
	v74 = base.I32_wrap_i64(v69)
	v75 = int32(0)
	if base.B2i32(v71 < int32(2048))|base.B2i32(v74 <= v75) == v75 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if int32(0) < v74 {
		goto L5
	} else {
		goto L33
	}
L9:
	;
	v82 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)+56))
	v105 = F_palloc(m, int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	if v82 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_hlparsetext_0), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(2047)
	v96 = F_errdetail(m, int32(_a_F_hlparsetext_1), v16)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_hlparsetext_2), int32(592), int32(_a_F_hlparsetext_3))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+8)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v103
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v74
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	if v110 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v105)+12)) = int32(0)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v116
	v122 = F_LexizeExec(m, v14+int32(-56), v14+int32(-60))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+12)) = v105
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v105
	goto L19
L23:
	;
	if v122 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v124 = v122
	goto L27
L25:
	;
	goto L26
L26:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_addHLParsedLex(m, l1, l2, v163, int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L32
	}
L27:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v137 + int32(1)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	F_addHLParsedLex(m, l1, l2, v141, v124)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v148 = F_LexizeExec(m, v14+int32(-56), v14+int32(-60))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v148 != 0 {
		v124 = v148
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
	m.G0 = v16 - int32(-64)
	return
}
func F_hmac_init(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v13 = m.T0[v12].(func(*base.Module, int32) int32)(m, v11)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_palloc0(m, v13)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if base.Ui32(v13) < base.Ui32(l2) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				m.T0[v18].(func(*base.Module, int32, int32, int32))(m, v11, l1, l2)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
					m.T0[v21].(func(*base.Module, int32, int32))(m, v11, v15)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
						m.T0[v24].(func(*base.Module, int32))(m, v11)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							if v13 == int32(0) {
							} else {
								v32 = int32(0)
								if v13 != int32(1) {
									v41 = v32
									v48 = int32(0)
									for {
										v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v51 = v41 + v15
										v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
										v53 = int32(54)
										v54 = v52 ^ v53
										*(*uint8)(unsafe.Add(mBase, uint32(v49+v41))) = uint8(v54)
										v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
										v59 = int32(92)
										v60 = v58 ^ v59
										*(*uint8)(unsafe.Add(mBase, uint32(v56+v41))) = uint8(v60)
										v63 = v41 | int32(1)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v66 = v63 + v15
										v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
										v69 = v67 ^ v53
										*(*uint8)(unsafe.Add(mBase, uint32(v63+v64))) = uint8(v69)
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
										v75 = v73 ^ v59
										*(*uint8)(unsafe.Add(mBase, uint32(v71+v63))) = uint8(v75)
										v77 = int32(2)
										v78 = v41 + v77
										v80 = v48 + v77
										if v80 != v13&int32(-2) {
											v41 = v78
											v48 = v80
											continue
										} else {
											break
										}
										break
									}
									if v13&int32(1) == int32(0) {
									} else {
										v86 = v78
										v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
										v96 = v86 + v15
										v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
										v99 = v97 ^ int32(54)
										*(*uint8)(unsafe.Add(mBase, uint32(v94+v86))) = uint8(v99)
										v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
										v105 = v103 ^ int32(92)
										*(*uint8)(unsafe.Add(mBase, uint32(v101+v86))) = uint8(v105)
									}
								} else {
									v86 = v32
									v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
									v96 = v86 + v15
									v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
									v99 = v97 ^ int32(54)
									*(*uint8)(unsafe.Add(mBase, uint32(v94+v86))) = uint8(v99)
									v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
									v105 = v103 ^ int32(92)
									*(*uint8)(unsafe.Add(mBase, uint32(v101+v86))) = uint8(v105)
								}
							}
							if v13 != 0 {
								base.MemoryFill(m, v15, int32(0), v13)
							} else {
							}
							F_pfree(m, v15)
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
								m.T0[v122].(func(*base.Module, int32, int32, int32))(m, v11, v121, v13)
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			} else {
				if l2 == int32(0) {
				} else {
					base.MemoryCopy(m, v15, l1, l2)
				}
				if v13 == int32(0) {
				} else {
					v32 = int32(0)
					if v13 != int32(1) {
						v41 = v32
						v48 = int32(0)
						for {
							v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v51 = v41 + v15
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							v53 = int32(54)
							v54 = v52 ^ v53
							*(*uint8)(unsafe.Add(mBase, uint32(v49+v41))) = uint8(v54)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							v59 = int32(92)
							v60 = v58 ^ v59
							*(*uint8)(unsafe.Add(mBase, uint32(v56+v41))) = uint8(v60)
							v63 = v41 | int32(1)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v66 = v63 + v15
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
							v69 = v67 ^ v53
							*(*uint8)(unsafe.Add(mBase, uint32(v63+v64))) = uint8(v69)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
							v75 = v73 ^ v59
							*(*uint8)(unsafe.Add(mBase, uint32(v71+v63))) = uint8(v75)
							v77 = int32(2)
							v78 = v41 + v77
							v80 = v48 + v77
							if v80 != v13&int32(-2) {
								v41 = v78
								v48 = v80
								continue
							} else {
								break
							}
							break
						}
						if v13&int32(1) == int32(0) {
						} else {
							v86 = v78
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
							v96 = v86 + v15
							v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
							v99 = v97 ^ int32(54)
							*(*uint8)(unsafe.Add(mBase, uint32(v94+v86))) = uint8(v99)
							v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
							v105 = v103 ^ int32(92)
							*(*uint8)(unsafe.Add(mBase, uint32(v101+v86))) = uint8(v105)
						}
					} else {
						v86 = v32
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
						v96 = v86 + v15
						v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
						v99 = v97 ^ int32(54)
						*(*uint8)(unsafe.Add(mBase, uint32(v94+v86))) = uint8(v99)
						v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
						v105 = v103 ^ int32(92)
						*(*uint8)(unsafe.Add(mBase, uint32(v101+v86))) = uint8(v105)
					}
				}
				if v13 != 0 {
					base.MemoryFill(m, v15, int32(0), v13)
				} else {
				}
				F_pfree(m, v15)
				mBase = m.M
				v120 = m.ExcPending
				if v120 != 0 {
					return
				} else {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					m.T0[v122].(func(*base.Module, int32, int32, int32))(m, v11, v121, v13)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_hnswbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 float64
	_ = v29
	var v31 int32
	_ = v31
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v41 float64
	_ = v41
	v5 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = F_palloc(m, int32(80))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_HnswGetTypeInfo(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
				F_HnswInitSupport(m, v10-int32(-64), l0)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbeginscan[0]))
					v25 = F_AllocSetContextCreateInternal(m, v20, int32(_a_F_hnswbeginscan_0), int32(0), int32(_a_F_hnswbeginscan_1), int32(_a_F_hnswbeginscan_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v25
						v29 = *(*float64)(unsafe.Add(mBase, _c_F_hnswbeginscan[1]))
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbeginscan[2]))
						v37 = base.F64_add(base.F64_mul(base.F64_mul(v29, base.F64_convert_i32_s(v31)), float64(1024)), float64(256))
						v38 = float64(2.147483647e+09)
						if base.F64_lt(v37, v38) != 0 {
							v41 = v37
						} else {
							v41 = v38
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = base.I32_trunc_sat_f64_u(v41)
						*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v10
						return v5
					}
				}
			}
		}
	}
}
func F_hs_concat(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_concat(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_hypot(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v26 float64
	_ = v26
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v42 float64
	_ = v42
	var v50 float64
	_ = v50
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v83 float64
	_ = v83
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v110 float64
	_ = v110
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = base.F64_abs(l0)
	v16 = base.F64_abs(l1)
	v19 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v15)) < base.Ui64(base.I64_reinterpret_f64(v16)))
	if base.Ui64(base.I64_reinterpret_f64(v15)) < base.Ui64(base.I64_reinterpret_f64(v16)) {
		v20 = v15
	} else {
		v20 = v16
	}
	v21 = base.I64_reinterpret_f64(v20)
	v23 = int64(base.Ui64(v21) >> (uint(int64(52)) % 64))
	if v23 == int64(2047) {
		v110 = v20
	} else {
		if base.Ui64(base.I64_reinterpret_f64(v15)) < base.Ui64(base.I64_reinterpret_f64(v16)) {
			v26 = v16
		} else {
			v26 = v15
		}
		if v21 == int64(0) {
			v110 = v26
		} else {
			v29 = base.I64_reinterpret_f64(v26)
			v31 = int64(base.Ui64(v29) >> (uint(int64(52)) % 64))
			if v31 == int64(2047) {
				v110 = v26
			} else {
				if int32(65) <= base.I32_wrap_i64(v31)-base.I32_wrap_i64(v23) {
					v110 = base.F64_add(v15, v16)
				} else {
					if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v29) {
						v42 = float64(1.90109156629516e-211)
						v55 = base.F64_mul(v26, v42)
						v56 = base.F64_mul(v20, v42)
						v57 = float64(5.260135901548374e+210)
					} else {
						if base.Ui64(int64(2580562586483294207)) < base.Ui64(v21) {
							v55 = v26
							v56 = v20
							v57 = float64(1)
						} else {
							v50 = float64(5.260135901548374e+210)
							v55 = base.F64_mul(v26, v50)
							v56 = base.F64_mul(v20, v50)
							v57 = float64(1.90109156629516e-211)
						}
					}
					v64 = base.F64_mul(v55, v55)
					*(*float64)(unsafe.Add(mBase, uint32(v13+int32(24)))) = v64
					v67 = base.F64_mul(v55, float64(1.34217729e+08))
					v69 = base.F64_add(v67, base.F64_sub(v55, v67))
					v70 = base.F64_sub(v55, v69)
					*(*float64)(unsafe.Add(mBase, uint32(v13+int32(16)))) = base.F64_add(base.F64_mul(v70, v70), base.F64_add(base.F64_mul(base.F64_add(v69, v69), v70), base.F64_sub(base.F64_mul(v69, v69), v64)))
					v83 = base.F64_mul(v56, v56)
					*(*float64)(unsafe.Add(mBase, uint32(v13+int32(8)))) = v83
					v86 = base.F64_mul(v56, float64(1.34217729e+08))
					v88 = base.F64_add(v86, base.F64_sub(v56, v86))
					v89 = base.F64_sub(v56, v88)
					*(*float64)(unsafe.Add(mBase, uint32(v13))) = base.F64_add(base.F64_mul(v89, v89), base.F64_add(base.F64_mul(base.F64_add(v88, v88), v89), base.F64_sub(base.F64_mul(v88, v88), v83)))
					v98 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
					v99 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
					v101 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
					v103 = *(*float64)(unsafe.Add(mBase, uint32(v13)+24))
					v110 = base.F64_mul(v57, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v98, v99), v101), v103)))
				}
			}
		}
	}
	m.G0 = v13 + int32(32)
	return v110
}
