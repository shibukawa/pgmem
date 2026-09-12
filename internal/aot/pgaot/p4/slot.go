package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckSlotPermissions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _consts[159]))
	v8 = F_has_rolreplication(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					F_errmsg(m, int32(112117), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(509734)
						F_errdetail(m, int32(548121), v4)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_errfinish(m, int32(473983), int32(1553), int32(135603))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
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
		} else {
			m.G0 = v4 + int32(16)
			return
		}
	}
}
func F_ExecFetchSlotHeapTupleDatum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
	if v6 != 0 {
		v8 = v6
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
		v8 = v7
	}
	v9 = m.T0[v8].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_heap_copy_tuple_as_datum(m, v9, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v6 == int32(0) {
				F_pfree(m, v9)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					return v14
				}
			} else {
				return v14
			}
		}
	}
}
func F_slot_modify_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	m.T0[v18].(func(*base.Module, int32))(m, l0)
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v23 < v22 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_slot_getsomeattrs_int(m, l1, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v30 = v16 << (uint(int32(2)) % 32)
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L5
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v16 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v31 = F__emscripten_memcpy_bulkmem(m, v27, v28, v30)
	mBase = m.M
	goto L10
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	if int32(0) < v16 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v35 = F__emscripten_memcpy_bulkmem(m, v33, v34, v16)
	mBase = m.M
	goto L14
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L35
	}
L16:
	;
	v41 = int32(0)
	goto L19
L17:
	;
	goto L18
L18:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v163 = v161 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v163)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v166)
	goto L34
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51+v41<<(uint(int32(1))%32)))))
	if v55 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v149 = v41 + int32(1)
	if v149 != v16 {
		v41 = v149
		goto L19
	} else {
		goto L33
	}
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58+v55))))
	if v60 == int32(117) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v65 = int32(4)
	v72 = v63 + v64<<(uint(v65)%32) + v41*int32(100) + int32(20)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, _consts[846])) = v55
	v78 = v73 + v55<<(uint(v65)%32)
	switch v60 - int32(98) {
	case 0:
		goto L26
	default:
		goto L25
	case 18:
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[846])) = int32(-1)
	goto L21
L25:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v41<<(uint(int32(2))%32)))) = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v137 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v41))) = uint8(v137)
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v72)+68))
	F_getTypeBinaryInputInfo(m, v105, v13+int32(12), v13+int32(8))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v72)+68))
	F_getTypeInputInfo(m, v81, v13+int32(12), v13+int32(8))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v72)+76))
	v92 = F_OidInputFunctionCall(m, v88, v89, v90, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v94+v41<<(uint(int32(2))%32)))) = v92
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99+v41))) = uint8(v101)
	goto L24
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v72)+76))
	v115 = F_OidReceiveFunctionCall(m, v112, v78, v113, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v117+v41<<(uint(int32(2))%32)))) = v115
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v122 != v123 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v41))) = uint8(v127)
	goto L24
L33:
	;
	goto L20
L34:
	;
	m.G0 = v13 + int32(16)
	return
L35:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v55 + int32(1)
	F_errmsg(m, int32(454931), v13)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(476327), int32(962), int32(484566))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
