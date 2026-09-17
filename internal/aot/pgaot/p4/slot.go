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
						F_errdetail(m, int32(_a_F_CheckSlotPermissions_2), v4)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckSlotPermissions_3), int32(1553), int32(_a_F_CheckSlotPermissions_4))
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
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+36))
	if v5 != 0 {
		v7 = v5
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v4)+44))
		v7 = v6
	}
	v8 = m.T0[v7].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v13 = F_heap_copy_tuple_as_datum(m, v8, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			if v5 == int32(0) {
				F_pfree(m, v8)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v13
				}
			} else {
				return v13
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v48 int32
	_ = v48
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
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
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
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
	v28 = v16 << (uint(int32(2)) % 32)
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	base.MemoryCopy(m, v29, v30, v28)
	goto L9
L8:
	;
	goto L9
L9:
	;
	if v16 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	base.MemoryCopy(m, v32, v33, v16)
	goto L12
L11:
	;
	goto L12
L12:
	;
	if int32(0) < v16 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L35
	}
L14:
	;
	v39 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v162 = v160 & int32(_a_F_slot_modify_data_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v162)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v165)
	goto L34
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49+v39<<(uint(int32(1))%32)))))
	if v53 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	v148 = v39 + int32(1)
	if v148 != v16 {
		v39 = v148
		goto L17
	} else {
		goto L33
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56+v53))))
	if v58 == int32(117) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v63 = int32(4)
	v70 = v61 + v62<<(uint(v63)%32) + v39*int32(100) + int32(20)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, _c_F_slot_modify_data[0])) = v53
	v76 = v71 + v53<<(uint(v63)%32)
	v78 = v58 - int32(98)
	if v78 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_slot_modify_data[0])) = int32(-1)
	goto L19
L23:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v129+v39<<(uint(int32(2))%32)))) = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v137 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v135+v39))) = uint8(v137)
	goto L22
L24:
	;
	if v78 != int32(18) {
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = int32(0)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
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
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v70)+68))
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
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v70)+76))
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
	*(*int32)(unsafe.Add(mBase, uint32(v94+v39<<(uint(int32(2))%32)))) = v92
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99+v39))) = uint8(v101)
	goto L22
L30:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v70)+76))
	v115 = F_OidReceiveFunctionCall(m, v112, v76, v113, v114)
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
	*(*int32)(unsafe.Add(mBase, uint32(v117+v39<<(uint(int32(2))%32)))) = v115
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v122 != v123 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v127 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v39))) = uint8(v127)
	goto L22
L33:
	;
	goto L18
L34:
	;
	m.G0 = v13 + int32(16)
	return
L35:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v53 + int32(1)
	F_errmsg(m, int32(_a_F_slot_modify_data_1), v13)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_slot_modify_data_2), int32(962), int32(_a_F_slot_modify_data_3))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
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
