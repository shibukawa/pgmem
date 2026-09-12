package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_read_stream_end(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	F_read_stream_reset(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		F_pfree(m, l0)
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			return
		}
	}
}
func F_read_stream_look_ahead(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v4 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)))
	if v9 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	if v29 <= v28 {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(v25)
	goto L3
L7:
	;
	return
L8:
	;
	F_errmsg_internal(m, int32(125790), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(488484), int32(1084), int32(405298))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L11:
	;
	return
L12:
	;
	F_pgaio_submit_staged(m)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L57
	}
L13:
	;
	v110 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+52)))
	if v110 <= int32(0) {
		goto L47
	} else {
		goto L48
	}
L14:
	;
	goto L15
L15:
	;
	v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+52)))
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+12)))
	if v34 <= v35+v36 {
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v105 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v105)
	goto L13
L17:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	if v39 == v35&int32(65535) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L16
L19:
	;
	v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	if v102 < v103 {
		goto L15
	} else {
		goto L46
	}
L20:
	;
	v43 = F_read_stream_start_pending_read(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v45 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L19
L24:
	;
	if base.I32_extend16_s(v70) <= int32(0) {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(-1)
	v70 = v35
	v71 = v45
	goto L24
L26:
	;
	goto L27
L27:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+76)))
	v54 = v53 + v35
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v55 <= base.I32_extend16_s(v54) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v59 = v55
	goto L30
L29:
	;
	v59 = int32(0)
	goto L30
L30:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v65 = m.T0[v64].(func(*base.Module, int32, int32, int32) int32)(m, l0, v50, v51+v52*base.I32_extend16_s(v54-v59))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v65 == int32(-1) {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)))
	v70 = v69
	v71 = v65
	goto L24
L33:
	;
	goto L37
L34:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v75+v70&int32(65535) != v71 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v81 = v70 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v81)
	goto L19
L36:
	;
	v96 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+52)) = uint16(v96)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v71
	goto L19
L37:
	;
	v86 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+52)))
	if v86 <= int32(0) {
		goto L36
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v71
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v95 != 0 {
		goto L12
	} else {
		goto L45
	}
L39:
	;
	v89 = F_read_stream_start_pending_read(m, l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L40
	}
L40:
	;
	if v89 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v91 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v91 != v92 {
		goto L37
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L38
L44:
	;
	goto L43
L45:
	;
	goto L11
L46:
	;
	goto L13
L47:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v129 != int32(1) {
		goto L11
	} else {
		goto L56
	}
L48:
	;
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	if v113 == v110&int32(65535) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v124 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	if v124 <= v123 {
		goto L47
	} else {
		goto L54
	}
L50:
	;
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v110 < v117 {
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if v117 == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if v121 != 0 {
		goto L47
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	v126 = F_read_stream_start_pending_read(m, l0)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L7
	} else {
		goto L55
	}
L55:
	;
	goto L47
L56:
	;
	goto L12
L57:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _consts[720]))
	v139 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v138)+20)) = uint8(v139)
	goto L11
}
func F_stream_change_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(396888)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(992)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v16
	v20 = int32(4463656)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v10 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v10 + int32(16)
	v30 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+147)) = uint8(v30)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v32
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+164)) = uint8(v5)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+152)) = v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	if v38 == v5 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(495819)
				F_errmsg(m, int32(314020), v10)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errfinish(m, int32(489940), int32(1534), int32(215218))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		m.T0[v38].(func(*base.Module, int32, int32, int32, int32))(m, v12, l1, l2, l3)
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return
		} else {
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			*(*int32)(unsafe.Add(mBase, _consts[337])) = v61
			m.G0 = v10 + int32(32)
			return
		}
	}
}
func F_stream_truncate_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+104))
	if v15 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(351963)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v14
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(992)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v19
		v23 = int32(4463656)
		v24 = *(*int32)(unsafe.Add(mBase, _consts[337]))
		*(*int32)(unsafe.Add(mBase, _consts[337])) = v12 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v24
		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v12 + int32(16)
		v33 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+147)) = uint8(v33)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+160)) = v35
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		v38 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14)+164)) = uint8(v38)
		*(*int64)(unsafe.Add(mBase, uint32(v14)+152)) = v37
		m.T0[v15].(func(*base.Module, int32, int32, int32, int32, int32))(m, v14, l1, l2, l3, l4)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return
		} else {
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			*(*int32)(unsafe.Add(mBase, _consts[337])) = v44
			m.G0 = v12 + int32(32)
			return
		}
	} else {
		m.G0 = v12 + int32(32)
		return
	}
}
