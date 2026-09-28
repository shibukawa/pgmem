package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_window_frame_options(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	if l0&int32(1) != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if l0&int32(2) != 0 {
		v21 = int32(_a_F_get_window_frame_options_0)
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v25 = l0 & int32(16)
	if v25 != 0 {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	F_appendStringInfoString(m, v9, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	if l0&int32(4) != 0 {
		v21 = int32(_a_F_get_window_frame_options_18)
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if l0&int32(8) == int32(0) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v21 = int32(_a_F_get_window_frame_options_19)
	goto L5
L9:
	;
	return
L10:
	;
	goto L4
L11:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_get_window_frame_options_1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l0&int32(32) != 0 {
		v50 = int32(_a_F_get_window_frame_options_2)
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	if v25 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L16:
	;
	F_appendStringInfoString(m, v9, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L9
	} else {
		goto L23
	}
L17:
	;
	if l0&int32(512) != 0 {
		v50 = int32(_a_F_get_window_frame_options_11)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if l0&int32(_a_F_get_window_frame_options_17) == int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	F_get_rule_expr(m, l1, l3, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if l0&int32(2048) != 0 {
		v50 = int32(_a_F_get_window_frame_options_14)
		goto L16
	} else {
		goto L21
	}
L21:
	;
	if l0&int32(_a_F_get_window_frame_options_20) == int32(0) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v50 = int32(_a_F_get_window_frame_options_16)
	goto L16
L23:
	;
	goto L15
L24:
	;
	if l0&int32(_a_F_get_window_frame_options_3) != 0 {
		v93 = int32(_a_F_get_window_frame_options_4)
		goto L36
	} else {
		goto L37
	}
L25:
	;
	F_appendStringInfoString(m, v9, int32(_a_F_get_window_frame_options_9))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	if l0&int32(256) != 0 {
		v79 = int32(_a_F_get_window_frame_options_10)
		goto L27
	} else {
		goto L28
	}
L27:
	;
	F_appendStringInfoString(m, v9, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L9
	} else {
		goto L34
	}
L28:
	;
	if l0&int32(1024) != 0 {
		v79 = int32(_a_F_get_window_frame_options_11)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if l0&int32(_a_F_get_window_frame_options_12) == int32(0) {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	F_get_rule_expr(m, l2, l3, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	if l0&int32(_a_F_get_window_frame_options_13) != 0 {
		v79 = int32(_a_F_get_window_frame_options_14)
		goto L27
	} else {
		goto L32
	}
L32:
	;
	if l0&int32(_a_F_get_window_frame_options_15) == int32(0) {
		goto L24
	} else {
		goto L33
	}
L33:
	;
	v79 = int32(_a_F_get_window_frame_options_16)
	goto L27
L34:
	;
	goto L24
L35:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v98 = v96 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v98
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100+v98))) = uint8(v102)
	goto L3
L36:
	;
	F_appendStringInfoString(m, v9, v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L40
	}
L37:
	;
	if l0&int32(_a_F_get_window_frame_options_5) != 0 {
		v93 = int32(_a_F_get_window_frame_options_6)
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if l0&int32(_a_F_get_window_frame_options_7) == int32(0) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v93 = int32(_a_F_get_window_frame_options_8)
	goto L36
L40:
	;
	goto L35
}
func F_window_cume_dist(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v8 = F_WinGetPartitionRowCount(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_WinCheckAndInitializeNullTreatment(m, v7, int32(0), l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)+176))
	goto L4
L4:
	;
	v18 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	F_WinSetMarkPosition(m, v7, v16)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L11
	}
L6:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	if v20 == int64(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(1)
	v31 = int32(0)
	goto L5
L8:
	;
	goto L9
L9:
	;
	v27 = F_WinRowsArePeers(m, v7, v16-int64(1), v16)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v31 = v27 ^ int32(1)
	goto L5
L11:
	;
	v35 = F_WinGetPartitionLocalMemory(m, v7, int32(8))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v31 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v70), base.F64_convert_i64_s(v8)))
L14:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	if v39 != int64(1) {
		v70 = v39
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)+176))
	goto L18
L17:
	;
	goto L16
L18:
	;
	v46 = v44 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v46
	if v8 <= v46 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v70 = v46
	goto L13
L20:
	;
	goto L21
L21:
	;
	v50 = v46
	goto L22
L22:
	;
	v57 = F_WinRowsArePeers(m, v7, v50-int64(1), v50)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v70 = v63
	goto L13
L24:
	;
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	if v57 == int32(0) {
		v70 = v59
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v62 = int64(1)
	v63 = v59 + v62
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v63
	v66 = v50 + v62
	if v66 != v8 {
		v50 = v66
		goto L22
	} else {
		goto L26
	}
L26:
	;
	goto L23
}
func F_window_lag_with_offset(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int64
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_WinCheckAndInitializeNullTreatment(m, v10, int32(1), l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v18 = v8 + int32(15)
		v19 = F_WinGetFuncArgCurrent(m, v10, int32(1), v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			if v21 == int32(0) {
				v24 = int32(0)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				if v27 == v24 {
					v75 = v24
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
					if v33 == int32(0) {
						v75 = v24
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v38 = v36 - int32(11)
						if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v38))|base.B2i32(int32(base.Ui32(int32(977))>>(uint(v38)%32))&int32(1) == int32(0))|int32(0) != 0 {
							v75 = v24
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v38<<(uint(int32(2))%32))+uint32(_c_F_window_lag_with_offset[0])))
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v33+v53)))
							if v55 == int32(0) {
								v75 = v24
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
								if v58 <= int32(1) {
									v75 = v24
								} else {
									v60 = int32(1)
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(4))))
									v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
									switch v66 - int32(7) {
									case 0:
										v75 = v60
									case 1:
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
										if v69 == int32(0) {
											v75 = v60
										} else {
											v75 = int32(0)
										}
									default:
										v75 = int32(0)
									}
								}
							}
						}
					}
				}
				v78 = F_WinGetFuncArgInPartition(m, v10, v24-base.I32_wrap_i64(v19), v75, v18, v8+int32(14))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int64(0)
				} else {
					v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
					if v80 != int32(1) {
						v87 = v78
					} else {
						v84 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
						v87 = int64(0)
					}
					m.G0 = v8 + int32(16)
					return v87
				}
			} else {
				v84 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
				v87 = int64(0)
				m.G0 = v8 + int32(16)
				return v87
			}
		}
	}
}
