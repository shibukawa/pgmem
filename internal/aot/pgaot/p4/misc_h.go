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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = int32(1)
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
	switch v57 - int32(2) {
	case 0, 1, 2, 3:
		goto L26
	default:
		goto L22
	case 5, 6, 7, 8:
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
	F_errfinish(m, int32(_a_F_HandleFatalError_5), int32(3272), int32(_a_F_HandleFatalError_6))
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
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	v1 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
	if v10 < int32(0) {
		v69 = v1
		m.G0 = v6 + int32(16)
		return v69
	} else {
		v14 = v10 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v14
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[1]))
		if v14 == v17 {
			v64 = int32(_a_F_HoldingBufferPinThatDelaysRecovery_0)
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
			v69 = base.B2i32(int32(0) < v65)
			m.G0 = v6 + int32(16)
			return v69
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[2]))
			if v14 == v21 {
				v64 = int32(_a_F_HoldingBufferPinThatDelaysRecovery_1)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
				v69 = base.B2i32(int32(0) < v65)
				m.G0 = v6 + int32(16)
				return v69
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[3]))
				if v14 == v25 {
					v64 = int32(_a_F_HoldingBufferPinThatDelaysRecovery_2)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
					v69 = base.B2i32(int32(0) < v65)
					m.G0 = v6 + int32(16)
					return v69
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[4]))
					if v14 == v29 {
						v64 = int32(_a_F_HoldingBufferPinThatDelaysRecovery_3)
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
						v69 = base.B2i32(int32(0) < v65)
						m.G0 = v6 + int32(16)
						return v69
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[5]))
						if v14 == v33 {
							v64 = int32(_a_F_HoldingBufferPinThatDelaysRecovery_4)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
							v69 = base.B2i32(int32(0) < v65)
							m.G0 = v6 + int32(16)
							return v69
						} else {
							v37 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[6]))
							if v14 == v37 {
								v64 = int32(_a_F_HoldingBufferPinThatDelaysRecovery_5)
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
								v69 = base.B2i32(int32(0) < v65)
								m.G0 = v6 + int32(16)
								return v69
							} else {
								v41 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[7]))
								if v14 == v41 {
									v64 = int32(_a_F_HoldingBufferPinThatDelaysRecovery_6)
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
									v69 = base.B2i32(int32(0) < v65)
									m.G0 = v6 + int32(16)
									return v69
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[8]))
									if v14 == v45 {
										v64 = int32(_a_F_HoldingBufferPinThatDelaysRecovery_7)
										v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										v69 = base.B2i32(int32(0) < v65)
										m.G0 = v6 + int32(16)
										return v69
									} else {
										v49 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[9]))
										if v49 == int32(0) {
											v69 = v1
											m.G0 = v6 + int32(16)
											return v69
										} else {
											v53 = *(*int32)(unsafe.Add(mBase, _c_F_HoldingBufferPinThatDelaysRecovery[10]))
											v56 = int32(0)
											v58 = F_hash_search(m, v53, v6+int32(12), v56, v56)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return int32(0)
											} else {
												if v58 == int32(0) {
													v69 = v1
												} else {
													v64 = v58
													v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
													v69 = base.B2i32(int32(0) < v65)
												}
												m.G0 = v6 + int32(16)
												return v69
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
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v149 int64
	_ = v149
	var v151 int32
	_ = v151
	var v155 int64
	_ = v155
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(_a_F_handle_sig_alarm_0)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[0])) = v13 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[1]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[2])) = v22
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[3]))
	if v25 == v22 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v168 = int32(_a_F_handle_sig_alarm_0)
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[0])) = v170 - int32(1)
	m.G0 = v9 + int32(16)
	return
L5:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[3])) = v29
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
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
	F_gettimeofday(m, v40)
	mBase = m.M
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	v44 = int64(*(*int32)(unsafe.Add(mBase, uint32(v40)+8)))
	m.G0 = v40 + v39
	v52 = v44 + v43*int64(1000000) - int64(946684800000000)
	goto L7
L7:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
	if v54 <= int32(0) {
		v155 = v52
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_schedule_alarm(m, v155)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
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
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[5]))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v64)+24))
	if v58 < v65 {
		v155 = v58
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v155 = v149
	goto L8
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[5]))
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
	if v70 <= int32(0) {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v73 = int32(0)
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[5]))
	*(*uint8)(unsafe.Add(mBase, uint32(v75)+4)) = uint8(v73)
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
	if int32(2) <= v80 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	__phi83 = v73
	__phi88 = int32(1)
	v83 = __phi83
	v88 = __phi88
	goto L17
L15:
	;
	goto L16
L16:
	;
	v106 = int32(_a_F_handle_sig_alarm_1)
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
	v109 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4])) = v108 - v109
	*(*uint8)(unsafe.Add(mBase, uint32(v68)+5)) = uint8(v109)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	m.T0[v114].(func(*base.Module))(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L20
	}
L17:
	;
	v89 = int32(2)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88<<(uint(v89)%32))+uint32(_c_F_handle_sig_alarm[5])))
	*(*int32)(unsafe.Add(mBase, uint32(v83<<(uint(v89)%32))+uint32(_c_F_handle_sig_alarm[5]))) = v93
	v96 = v88 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
	if v96 < v98 {
		__phi83 = v88
		__phi88 = v96
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
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v68)+32))
	if int32(0) < v117 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v123 = base.I64_extend_i32_u(v117) * int64(1000)
	v125 = *(*int64)(unsafe.Add(mBase, uint32(v68)+24))
	v126 = v125 + v123
	if v126 < v58 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v135 = m.G0
	v136 = int32(16)
	v137 = v135 - v136
	m.G0 = v137
	F_gettimeofday(m, v137)
	mBase = m.M
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v137)))
	v141 = int64(*(*int32)(unsafe.Add(mBase, uint32(v137)+8)))
	m.G0 = v137 + v136
	v149 = v141 + v140*int64(1000000) - int64(946684800000000)
	goto L28
L24:
	;
	v128 = v123 + v58
	goto L26
L25:
	;
	v128 = v126
	goto L26
L26:
	;
	F_enable_timeout(m, v120, v58, v128, v117)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L23
L28:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
	if int32(0) < v151 {
		v58 = v149
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
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_handle_sig_alarm[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v184 - int32(1)
	F_errmsg_internal(m, int32(_a_F_handle_sig_alarm_2), v9)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_handle_sig_alarm_3), int32(143), int32(_a_F_handle_sig_alarm_4))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
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
			v15 = F_convert_any_priv_string(m, v12, int32(_a_F_has_parameter_privilege_id_name_0))
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
			v15 = F_convert_any_priv_string(m, v12, int32(_a_F_has_parameter_privilege_name_name_0))
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
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int64
	_ = v243
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 float64
	_ = v319
	var v321 float64
	_ = v321
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
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
	v330 = m.ExcPending
	if v330 != 0 {
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
	v170 = m.T0[v169].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v159, v160, v159, v160, int32(-1), int32(137), v15+int32(-48), v160)
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
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = int32(1831)
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+40)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v86)+60)) = v94
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+36)) = uint8(v122)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = int32(1832)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = int32(1833)
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = int32(1836)
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(1837)
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
	F_errfinish(m, int32(_a_F_hashbuild_5), int32(467), int32(_a_F_hashbuild_6))
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
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v211 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v178+int32(88))+232)) = base.I64_trunc_sat_f64_s(v174)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v201 + v190
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5])) = v207 - v190
	goto L31
L34:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	F_tuplesort_performsort(m, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v317 = F_palloc(m, int32(16))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L57
	}
L37:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v217 = F_tuplesort_getheaptuple(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v217 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v219 = v217
	v232 = int64(0)
	goto L42
L40:
	;
	goto L41
L41:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	F_tuplesort_end(m, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L55
	}
L42:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	F__hash_doinsert(m, v233, v219, v212, int32(1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[6]))
	if v238 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v243 = v232 + int64(1)
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[3]))
	if v246 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L47
L49:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	v280 = F_tuplesort_getheaptuple(m, v279)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	goto L49
L51:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_hashbuild[4])))
	if v250&int32(1) == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v255 = int32(_a_F_hashbuild_7)
	v257 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5]))
	v258 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5])) = v257 + v258
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = v261 + v258
	*(*int64)(unsafe.Add(mBase, uint32(v246+int32(96))+232)) = v243
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = v269 + v258
	v275 = *(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbuild[5])) = v275 - v258
	goto L50
L53:
	;
	if v280 != 0 {
		v219 = v280
		v232 = v243
		goto L42
	} else {
		goto L54
	}
L54:
	;
	goto L43
L55:
	;
	F_pfree(m, v296)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L36
L57:
	;
	v319 = *(*float64)(unsafe.Add(mBase, uint32(v17)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v317))) = v319
	v321 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v317)+8)) = v321
	m.G0 = v17 - int32(-64)
	return v317
L58:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v331 + int32(4)
	F_errmsg_internal(m, int32(_a_F_hashbuild_8), v17)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_hashbuild_9), int32(138), int32(_a_F_hashbuild_10))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
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
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
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
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
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
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
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
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
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
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_strlen(m, v3)
	mBase = m.M
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v12 = v4 - int32(1636608432)
	if v6 == int64(0) {
		v49 = v12
		v51 = v12
		v53 = v12
	} else {
		v16 = v12 + base.I32_wrap_i64(v6)
		v17 = v16 + v12
		v21 = int32(4)
		v23 = base.I32_wrap_i64(int64(base.Ui64(v6)>>(uint(int64(32))%64))) ^ base.I32_rotl(v12, v21)
		v27 = v16 - v23 ^ base.I32_rotl(v23, int32(6))
		v31 = v17 - v27 ^ base.I32_rotl(v27, int32(8))
		v32 = v17 + v23
		v33 = v27 + v32
		v34 = v31 + v33
		v38 = v32 - v31 ^ base.I32_rotl(v31, int32(16))
		v42 = v33 - v38 ^ base.I32_rotl(v38, int32(19))
		v47 = v34 + v38
		v49 = v47
		v51 = v34 - v42 ^ base.I32_rotl(v42, v21)
		v53 = v42 + v47
	}
	if v3&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v4) {
			v58 = v3
			v59 = v4
			v61 = v49
			v62 = v53
			v63 = v51
			for {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
				v66 = v65 + v62
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
				v70 = v69 + v63
				v72 = int32(4)
				v74 = v67 + v61 - v70 ^ base.I32_rotl(v70, v72)
				v78 = v66 - v74 ^ base.I32_rotl(v74, int32(6))
				v79 = v70 + v66
				v80 = v74 + v79
				v81 = v78 + v80
				v85 = v79 - v78 ^ base.I32_rotl(v78, int32(8))
				v89 = v80 - v85 ^ base.I32_rotl(v85, int32(16))
				v93 = v81 - v89 ^ base.I32_rotl(v89, int32(19))
				v94 = v85 + v81
				v95 = v89 + v94
				v96 = v93 + v95
				v100 = v94 - v93 ^ base.I32_rotl(v93, v72)
				v101 = int32(12)
				v102 = v58 + v101
				v104 = v59 - v101
				if base.Ui32(int32(11)) < base.Ui32(v104) {
					v58 = v102
					v59 = v104
					v61 = v95
					v62 = v96
					v63 = v100
					continue
				} else {
					break
				}
				break
			}
			v107 = v102
			v108 = v104
			v110 = v95
			v111 = v96
			v112 = v100
		} else {
			v107 = v3
			v108 = v4
			v110 = v49
			v111 = v53
			v112 = v51
		}
		switch v108 - int32(1) {
		case 0:
			v277 = v110
			v278 = v111
			v279 = v112
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 1:
			v270 = v110
			v271 = v111
			v272 = v112
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 2:
			v263 = v110
			v264 = v111
			v265 = v112
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 3:
			v257 = v111
			v258 = v112
			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
			v263 = v259<<(uint(int32(24))%32) + v110
			v264 = v257
			v265 = v258
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 4:
			v253 = v111
			v254 = v112
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
			v257 = v253 + v255
			v258 = v254
			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
			v263 = v259<<(uint(int32(24))%32) + v110
			v264 = v257
			v265 = v258
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 5:
			v247 = v111
			v248 = v112
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
			v253 = v249<<(uint(int32(8))%32) + v247
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
			v257 = v253 + v255
			v258 = v254
			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
			v263 = v259<<(uint(int32(24))%32) + v110
			v264 = v257
			v265 = v258
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 6:
			v241 = v111
			v242 = v112
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+6)))
			v247 = v243<<(uint(int32(16))%32) + v241
			v248 = v242
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
			v253 = v249<<(uint(int32(8))%32) + v247
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
			v257 = v253 + v255
			v258 = v254
			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
			v263 = v259<<(uint(int32(24))%32) + v110
			v264 = v257
			v265 = v258
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 7:
			v236 = v112
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+7)))
			v241 = v237<<(uint(int32(24))%32) + v111
			v242 = v236
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+6)))
			v247 = v243<<(uint(int32(16))%32) + v241
			v248 = v242
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
			v253 = v249<<(uint(int32(8))%32) + v247
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
			v257 = v253 + v255
			v258 = v254
			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
			v263 = v259<<(uint(int32(24))%32) + v110
			v264 = v257
			v265 = v258
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 8:
			v231 = v112
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+8)))
			v236 = v232<<(uint(int32(8))%32) + v231
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+7)))
			v241 = v237<<(uint(int32(24))%32) + v111
			v242 = v236
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+6)))
			v247 = v243<<(uint(int32(16))%32) + v241
			v248 = v242
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
			v253 = v249<<(uint(int32(8))%32) + v247
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
			v257 = v253 + v255
			v258 = v254
			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
			v263 = v259<<(uint(int32(24))%32) + v110
			v264 = v257
			v265 = v258
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 9:
			v226 = v112
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+9)))
			v231 = v227<<(uint(int32(16))%32) + v226
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+8)))
			v236 = v232<<(uint(int32(8))%32) + v231
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+7)))
			v241 = v237<<(uint(int32(24))%32) + v111
			v242 = v236
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+6)))
			v247 = v243<<(uint(int32(16))%32) + v241
			v248 = v242
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
			v253 = v249<<(uint(int32(8))%32) + v247
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
			v257 = v253 + v255
			v258 = v254
			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
			v263 = v259<<(uint(int32(24))%32) + v110
			v264 = v257
			v265 = v258
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		case 10:
			v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+10)))
			v226 = v222<<(uint(int32(24))%32) + v112
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+9)))
			v231 = v227<<(uint(int32(16))%32) + v226
			v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+8)))
			v236 = v232<<(uint(int32(8))%32) + v231
			v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+7)))
			v241 = v237<<(uint(int32(24))%32) + v111
			v242 = v236
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+6)))
			v247 = v243<<(uint(int32(16))%32) + v241
			v248 = v242
			v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
			v253 = v249<<(uint(int32(8))%32) + v247
			v254 = v248
			v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
			v257 = v253 + v255
			v258 = v254
			v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+3)))
			v263 = v259<<(uint(int32(24))%32) + v110
			v264 = v257
			v265 = v258
			v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
			v270 = v266<<(uint(int32(16))%32) + v263
			v271 = v264
			v272 = v265
			v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
			v277 = v273<<(uint(int32(8))%32) + v270
			v278 = v271
			v279 = v272
			v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
			v285 = v277 + v280
			v286 = v278
			v287 = v279
		default:
			v285 = v110
			v286 = v111
			v287 = v112
		}
	} else {
		if base.Ui32(int32(12)) <= base.Ui32(v4) {
			v118 = v3
			v119 = v4
			v121 = v49
			v122 = v53
			v123 = v51
			for {
				v125 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
				v126 = v125 + v122
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v118)+8))
				v130 = v129 + v123
				v132 = int32(4)
				v134 = v127 + v121 - v130 ^ base.I32_rotl(v130, v132)
				v138 = v126 - v134 ^ base.I32_rotl(v134, int32(6))
				v139 = v130 + v126
				v140 = v134 + v139
				v141 = v138 + v140
				v145 = v139 - v138 ^ base.I32_rotl(v138, int32(8))
				v149 = v140 - v145 ^ base.I32_rotl(v145, int32(16))
				v153 = v141 - v149 ^ base.I32_rotl(v149, int32(19))
				v154 = v145 + v141
				v155 = v149 + v154
				v156 = v153 + v155
				v160 = v154 - v153 ^ base.I32_rotl(v153, v132)
				v161 = int32(12)
				v162 = v118 + v161
				v164 = v119 - v161
				if base.Ui32(int32(11)) < base.Ui32(v164) {
					v118 = v162
					v119 = v164
					v121 = v155
					v122 = v156
					v123 = v160
					continue
				} else {
					break
				}
				break
			}
			v167 = v162
			v168 = v164
			v170 = v155
			v171 = v156
			v172 = v160
		} else {
			v167 = v3
			v168 = v4
			v170 = v49
			v171 = v53
			v172 = v51
		}
		switch v168 - int32(1) {
		case 0:
			v219 = v170
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v285 = v219 + v220
			v286 = v171
			v287 = v172
		case 1:
			v214 = v170
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v219 = v215<<(uint(int32(8))%32) + v214
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v285 = v219 + v220
			v286 = v171
			v287 = v172
		case 2:
			v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+2)))
			v214 = v210<<(uint(int32(16))%32) + v170
			v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+1)))
			v219 = v215<<(uint(int32(8))%32) + v214
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
			v285 = v219 + v220
			v286 = v171
			v287 = v172
		case 3:
			v207 = v171
			v208 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
			v285 = v208 + v170
			v286 = v207
			v287 = v172
		case 4:
			v204 = v171
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v207 = v204 + v205
			v208 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
			v285 = v208 + v170
			v286 = v207
			v287 = v172
		case 5:
			v199 = v171
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v204 = v200<<(uint(int32(8))%32) + v199
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v207 = v204 + v205
			v208 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
			v285 = v208 + v170
			v286 = v207
			v287 = v172
		case 6:
			v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+6)))
			v199 = v195<<(uint(int32(16))%32) + v171
			v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+5)))
			v204 = v200<<(uint(int32(8))%32) + v199
			v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+4)))
			v207 = v204 + v205
			v208 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
			v285 = v208 + v170
			v286 = v207
			v287 = v172
		case 7:
			v190 = v172
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
			v193 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
			v285 = v191 + v170
			v286 = v193 + v171
			v287 = v190
		case 8:
			v185 = v172
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v190 = v186<<(uint(int32(8))%32) + v185
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
			v193 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
			v285 = v191 + v170
			v286 = v193 + v171
			v287 = v190
		case 9:
			v180 = v172
			v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+9)))
			v185 = v181<<(uint(int32(16))%32) + v180
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v190 = v186<<(uint(int32(8))%32) + v185
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
			v193 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
			v285 = v191 + v170
			v286 = v193 + v171
			v287 = v190
		case 10:
			v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+10)))
			v180 = v176<<(uint(int32(24))%32) + v172
			v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+9)))
			v185 = v181<<(uint(int32(16))%32) + v180
			v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+8)))
			v190 = v186<<(uint(int32(8))%32) + v185
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
			v193 = *(*int32)(unsafe.Add(mBase, uint32(v167)+4))
			v285 = v191 + v170
			v286 = v193 + v171
			v287 = v190
		default:
			v285 = v170
			v286 = v171
			v287 = v172
		}
	}
	v290 = int32(14)
	v292 = v286 ^ v287 - base.I32_rotl(v286, v290)
	v296 = v292 ^ v285 - base.I32_rotl(v292, int32(11))
	v300 = v296 ^ v286 - base.I32_rotl(v296, int32(25))
	v304 = v300 ^ v292 - base.I32_rotl(v300, int32(16))
	v308 = v304 ^ v296 - base.I32_rotl(v304, int32(4))
	v312 = v308 ^ v300 - base.I32_rotl(v308, v290)
	v322 = F_Int64GetDatum(m, base.I64_extend_i32_u(v312)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v312^v304-base.I32_rotl(v312, int32(24))))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		return int32(0)
	} else {
		return v322
	}
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
			v40 = v25 + v31
			v41 = v35 + v40
			v42 = v39 + v41
			v46 = v40 - v39 ^ base.I32_rotl(v39, int32(16))
			v50 = v41 - v46 ^ base.I32_rotl(v46, int32(19))
			v55 = v42 + v46
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
		v330 = F_Int64GetDatum(m, base.I64_extend_i32_u(v320)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v320^v312-base.I32_rotl(v320, int32(24))))
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
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
	v96 = int32(0)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+216)))
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
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+216)))
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
	v118 = base.I32_div_s(int32(0), int32(32))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v119 <= v118 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v122 = v105 + int32(8)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v122+v118<<(uint(int32(2))%32))))
	v129 = v126 & int32(-1)
	if v129 != 0 {
		v151 = v129
		v152 = v118
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v131 = v118 + int32(1)
	if v131 == v119 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v134 = v131
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
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
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
	v194 = base.I32_div_s(v192, int32(32))
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
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
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
				F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_0), v15+int32(48))
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
				F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_1), v15-int32(-64))
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
							F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_2))
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
				F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_3), v15+int32(96))
				mBase = m.M
				v238 = m.ExcPending
				if v238 != 0 {
					return
				} else {
					v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+6)))
					F_infobits_desc(m, l0, v239, int32(_a_F_heap2_desc_4))
					mBase = m.M
					v242 = m.ExcPending
					if v242 != 0 {
						return
					} else {
						v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+7)))
						*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v243
						F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_5), v15+int32(80))
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
				F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_6), v15+int32(128))
				mBase = m.M
				v266 = m.ExcPending
				if v266 != 0 {
					return
				} else {
					v267 = *(*int64)(unsafe.Add(mBase, uint32(v18)+4))
					v268 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+120)) = v268
					*(*int64)(unsafe.Add(mBase, uint32(v15)+112)) = v267
					F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_7), v15+int32(112))
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
				F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_8), v15+int32(32))
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
					F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_9), v15+int32(16))
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
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(0))+76)))
								if v66 != int32(1) {
									v83 = v56
									v86 = v83
								} else {
									v70 = v60 + int32(76)
									v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+43)))
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
											v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+48)))
											*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79
										} else {
										}
										v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
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
							F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_10), v15)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return
							} else {
								if v100 != 0 {
									F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_11))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
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
															F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
																		F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
																F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
																F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
														F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
										F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
																F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
														F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
											F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
														F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
				F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_9), v15+int32(16))
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
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(0))+76)))
							if v66 != int32(1) {
								v83 = v56
								v86 = v83
							} else {
								v70 = v60 + int32(76)
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+43)))
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
										v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+48)))
										*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79
									} else {
									}
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
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
						F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_10), v15)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							if v100 != 0 {
								F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_11))
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
											F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
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
														F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
																	F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
															F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
															F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
									F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
															F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
										F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
											F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
			F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_8), v15+int32(32))
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
				F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_9), v15+int32(16))
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
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(0))+76)))
							if v66 != int32(1) {
								v83 = v56
								v86 = v83
							} else {
								v70 = v60 + int32(76)
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+43)))
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
										v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+48)))
										*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79
									} else {
									}
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
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
						F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_10), v15)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return
						} else {
							if v100 != 0 {
								F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_11))
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
											F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
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
														F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
																	F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
															F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
															F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
									F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
															F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
										F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
											F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
			F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_9), v15+int32(16))
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
						v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+int32(0))+76)))
						if v66 != int32(1) {
							v83 = v56
							v86 = v83
						} else {
							v70 = v60 + int32(76)
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+43)))
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
									v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+48)))
									*(*int32)(unsafe.Add(mBase, uint32(v58))) = v79
								} else {
								}
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
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
					F_appendStringInfo(m, l0, int32(_a_F_heap2_desc_10), v15)
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return
					} else {
						if v100 != 0 {
							F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_11))
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
										F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
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
													F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
																F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
														F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
											F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
														F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
								F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_12))
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
											F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
														F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
									F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_13))
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
												F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
										F_appendStringInfoString(m, l0, int32(_a_F_heap2_desc_14))
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
func F_hemdist_3(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	v5 = Fn13923(m, l0, l1, l2, int32(4))
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
				F_errfinish(m, int32(_a_F_hide_coercion_node_1), int32(826), int32(_a_F_hide_coercion_node_2))
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
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
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
	*(*int64)(unsafe.Add(mBase, uint32(v13)+28)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v13)+12)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v13)+44)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = int32(0)
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
	v163 = F_FunctionCall1Coll(m, v22+int32(84), int32(0), v27)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L34
	}
L7:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v62 = int32(0)
	if base.B2i32(v59 < int32(2047))|base.B2i32(v57 <= v62) == v62 {
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
	v69 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	v92 = F_palloc(m, int32(16))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	if v69 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_hlparsetext_0), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(2047)
	F_errdetail(m, int32(_a_F_hlparsetext_1), v13)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_hlparsetext_2), int32(575), int32(_a_F_hlparsetext_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+8)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v92))) = v57
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	if v97 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v92)+12)) = int32(0)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v13)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v103
	v109 = F_LexizeExec(m, v11+int32(-56), v11+int32(-60))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v92
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v92
	goto L19
L23:
	;
	if v109 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v111 = v109
	goto L27
L25:
	;
	goto L26
L26:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_addHLParsedLex(m, l1, l2, v144, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L32
	}
L27:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v121 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	F_addHLParsedLex(m, l1, l2, v125, v111)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v132 = F_LexizeExec(m, v11+int32(-56), v11+int32(-60))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if v132 != 0 {
		v111 = v132
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
		v10 = F_palloc(m, int32(72))
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
				F_HnswInitSupport(m, v10+int32(56), l0)
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
						*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v25
						v29 = *(*float64)(unsafe.Add(mBase, _c_F_hnswbeginscan[1]))
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbeginscan[2]))
						v37 = base.F64_add(base.F64_mul(base.F64_mul(v29, base.F64_convert_i32_s(v31)), float64(1024)), float64(256))
						v38 = float64(2.147483647e+09)
						if base.F64_lt(v37, v38) != 0 {
							v41 = v37
						} else {
							v41 = v38
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = base.I32_trunc_sat_f64_u(v41)
						*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v10
						return v5
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
