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
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v6 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_look_ahead[0]))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
	if v11 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v30 != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)) = uint8(v27)
	goto L3
L7:
	;
	return
L8:
	;
	F_errmsg_internal(m, int32(_a_F_read_stream_look_ahead_0), int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F_errfinish(m, int32(_a_F_read_stream_look_ahead_1), int32(1094), int32(_a_F_read_stream_look_ahead_2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
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
	v177 = m.ExcPending
	if v177 != 0 {
		goto L7
	} else {
		goto L68
	}
L13:
	;
	v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+64)))
	if v147 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L14:
	;
	v139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+14)) = v139
	v144 = v139
	goto L13
L15:
	;
	v33 = v30
	goto L18
L16:
	;
	goto L17
L17:
	;
	v144 = int32(0)
	goto L13
L18:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v37 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	if v37 <= v36 {
		v144 = v33
		goto L13
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+64)))
	v41 = int32(0)
	if base.B2i32(v40 <= v41)|v39 == v41 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	if v132 != 0 {
		v33 = v132
		goto L18
	} else {
		goto L57
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v71 != int32(-1) {
		goto L36
	} else {
		goto L37
	}
L23:
	;
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	v59 = int32(0)
	if base.B2i32(v59 < v40)|v39|base.B2i32(v40+v39 < base.I32_extend16_s(v33)) != 0 {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v40 < v47 {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	v49 = v39
	goto L26
L26:
	;
	if base.I32_extend16_s(v33) <= base.I32_extend16_s(v49)+v40 {
		v144 = v33
		goto L13
	} else {
		goto L28
	}
L27:
	;
	v49 = int32(0)
	goto L26
L28:
	;
	if v40 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	v67 = base.B2i32(v40 < v57)
	goto L32
L31:
	;
	v67 = v59
	goto L32
L32:
	;
	if v67 != 0 {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	v68 = F_read_stream_start_pending_read(m, l0)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	goto L21
L35:
	;
	if base.I32_extend16_s(v96) <= int32(0) {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(-1)
	v96 = v40
	v97 = v71
	goto L35
L37:
	;
	goto L38
L38:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+88)))
	v80 = v79 + v40
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v81 <= base.I32_extend16_s(v80) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v85 = v81
	goto L41
L40:
	;
	v85 = int32(0)
	goto L41
L41:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v91 = m.T0[v90].(func(*base.Module, int32, int32, int32) int32)(m, l0, v76, v77+v78*base.I32_extend16_s(v80-v85))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	if v91 == int32(-1) {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)))
	v96 = v95
	v97 = v91
	goto L35
L44:
	;
	goto L48
L45:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v101+v96&int32(_a_F_read_stream_look_ahead_3) != v97 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v107 = v96 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v107)
	goto L21
L47:
	;
	v124 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+64)) = uint16(v124)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v97
	goto L21
L48:
	;
	v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+64)))
	if v114 <= int32(0) {
		goto L47
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v97
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v123 != 0 {
		goto L12
	} else {
		goto L56
	}
L50:
	;
	v117 = F_read_stream_start_pending_read(m, l0)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	if v117 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if v119 != v120 {
		goto L48
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L49
L55:
	;
	goto L54
L56:
	;
	goto L11
L57:
	;
	goto L19
L58:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+33)))
	if v168 != int32(1) {
		goto L11
	} else {
		goto L67
	}
L59:
	;
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	if v151 <= v150 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	if v144 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v165 = F_read_stream_start_pending_read(m, l0)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L7
	} else {
		goto L66
	}
L62:
	;
	v155 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+14)))
	if v155 <= v147 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if int32(0) < v147 {
		goto L58
	} else {
		goto L64
	}
L64:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if v159|base.B2i32(v147+v159 < base.I32_extend16_s(v144)) != 0 {
		goto L58
	} else {
		goto L65
	}
L65:
	;
	goto L61
L66:
	;
	goto L58
L67:
	;
	goto L12
L68:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _c_F_read_stream_look_ahead[0]))
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+20)) = uint8(v180)
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = int32(_a_F_stream_change_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(1058)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v16
	v20 = int32(_a_F_stream_change_cb_wrapper_1)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_stream_change_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_stream_change_cb_wrapper[0])) = v10 + int32(4)
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
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_stream_change_cb_wrapper_2)
				F_errmsg(m, int32(_a_F_stream_change_cb_wrapper_3), v10)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_stream_change_cb_wrapper_4), int32(1614), int32(_a_F_stream_change_cb_wrapper_5))
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
			*(*int32)(unsafe.Add(mBase, _c_F_stream_change_cb_wrapper[0])) = v61
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
		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(_a_F_stream_truncate_cb_wrapper_0)
		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v14
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l4)))
		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = int32(1058)
		*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v19
		v23 = int32(_a_F_stream_truncate_cb_wrapper_1)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_stream_truncate_cb_wrapper[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_stream_truncate_cb_wrapper[0])) = v12 + int32(4)
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
			*(*int32)(unsafe.Add(mBase, _c_F_stream_truncate_cb_wrapper[0])) = v44
			m.G0 = v12 + int32(32)
			return
		}
	} else {
		m.G0 = v12 + int32(32)
		return
	}
}
