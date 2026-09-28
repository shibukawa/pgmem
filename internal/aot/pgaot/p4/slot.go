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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_CheckSlotPermissions[0]))
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
					F_errmsg(m, int32(_a_F_CheckSlotPermissions_0), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_CheckSlotPermissions_1)
						v26 = F_errdetail(m, int32(_a_F_CheckSlotPermissions_2), v4)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckSlotPermissions_3), int32(1699), int32(_a_F_CheckSlotPermissions_4))
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
func F_ExecFetchSlotHeapTupleDatum(m *base.Module, l0 int32) int64 {
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
	var v14 int64
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
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v14 = F_heap_copy_tuple_as_datum(m, v9, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			if v6 == int32(0) {
				F_pfree(m, v9)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					return v14
				}
			} else {
				return v14
			}
		}
	}
}
func F_SlotSyncShmemRequest(m *base.Module, l0 int32) {
	var v6 int32
	_ = v6
	Fn14222(m, l0, int32(_a_F_SlotSyncShmemRequest_0), int64(24), int32(_a_F_SlotSyncShmemRequest_1))
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_slot_modify_data(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int64
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
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
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	m.T0[v19].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v24 < v23 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	m.T0[v27].(func(*base.Module, int32, int32))(m, l1, v23)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v31 = v17 << (uint(int32(3)) % 32)
	if v31 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	base.MemoryCopy(m, v32, v33, v31)
	goto L9
L8:
	;
	goto L9
L9:
	;
	if v17 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	base.MemoryCopy(m, v35, v36, v17)
	goto L12
L11:
	;
	goto L12
L12:
	;
	if int32(0) < v17 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L41
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L37
	}
L15:
	;
	v42 = int32(0)
	goto L18
L16:
	;
	goto L17
L17:
	;
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v169 = v167 & int32(_a_F_slot_modify_data_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v169)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v172)
	goto L36
L18:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53+v42<<(uint(int32(1))%32)))))
	if v57 < int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v154 = v42 + int32(1)
	if v154 != v17 {
		v42 = v154
		goto L18
	} else {
		goto L35
	}
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v60 <= v57 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v57))))
	if v64 == int32(117) {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v76 = v67 + v68<<(uint(int32(3))%32) + v42*int32(100) + int32(28)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, _c_F_slot_modify_data[0])) = v57
	v82 = v77 + v57<<(uint(int32(4))%32)
	v84 = v64 - int32(98)
	if v84 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_slot_modify_data[0])) = int32(-1)
	goto L20
L25:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v135+v42<<(uint(int32(3))%32)))) = int64(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v143 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v141+v42))) = uint8(v143)
	goto L24
L26:
	;
	if v84 != int32(18) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v76)+68))
	F_getTypeBinaryInputInfo(m, v111, v14+int32(28), v14+int32(24))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L32
	}
L29:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v76)+68))
	F_getTypeInputInfo(m, v87, v14+int32(28), v14+int32(24))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v76)+76))
	v98 = F_OidInputFunctionCall(m, v94, v95, v96, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v100+v42<<(uint(int32(3))%32)))) = v98
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v105+v42))) = uint8(v107)
	goto L24
L32:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v76)+76))
	v121 = F_OidReceiveFunctionCall(m, v118, v82, v119, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v123+v42<<(uint(int32(3))%32)))) = v121
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v128 != v129 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v131+v42))) = uint8(v133)
	goto L24
L35:
	;
	goto L19
L36:
	;
	m.G0 = v14 + int32(32)
	return
L37:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57 + int32(1)
	F_errmsg_plural(m, int32(_a_F_slot_modify_data_1), int32(_a_F_slot_modify_data_2), v184, v14+int32(16))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_slot_modify_data_3), int32(1173), int32(_a_F_slot_modify_data_4))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v57 + int32(1)
	F_errmsg(m, int32(_a_F_slot_modify_data_5), v14)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_slot_modify_data_3), int32(1214), int32(_a_F_slot_modify_data_4))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
