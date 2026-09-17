package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitResultTupleSlotTL(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+44))
	v8 = F_ExecTypeFromTLInternal(m, v6, int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v8
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v12 = F_MakeTupleTableSlot(m, v8, l1)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
			v15 = F_lappend(m, v14, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+104)) = v15
				*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v12
				v19 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v19)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = l1
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(base.B2i32(v22 != int32(0)))
				return
			}
		}
	}
}
func F_ExecInitResultTypeTL(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+44))
	v5 = F_ExecTypeFromTLInternal(m, v3, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v5
		return
	}
}
func F_ExecLookupResultRelByOid(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v99 int32
	_ = v99
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = l1
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	if v18 == v5 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L29
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return v99
L3:
	;
	if l2 == int32(0) {
		goto L1
	} else {
		goto L28
	}
L4:
	;
	if l2 != 0 {
		v99 = int32(0)
		goto L2
	} else {
		goto L27
	}
L5:
	;
	v52 = v5
	v53 = v5
	goto L17
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v21 <= int32(0) {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v27 = int32(0)
	v29 = F_hash_search(m, v18, v15+int32(12), v27, v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	goto L5
L10:
	;
	return int32(0)
L11:
	;
	if v29 == int32(0) {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if l3 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v99 = v43 + v42*int32(216)
	goto L2
L14:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v42 = v37
	goto L13
L15:
	;
	goto L16
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v40
	v42 = v40
	goto L13
L17:
	;
	v61 = v24 + v52*int32(216)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+56))
	v64 = base.B2i32(v63 == l1)
	if v63 == l1 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L4
L19:
	;
	v65 = v61
	goto L21
L20:
	;
	v65 = v53
	goto L21
L21:
	;
	v66 = int32(0)
	if base.B2i32(l3 == v66)|base.B2i32(l1 != v63) == v66 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = l1
	v74 = v61
	goto L24
L23:
	;
	v74 = v65
	goto L24
L24:
	;
	if v63 == l1 {
		v99 = v74
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v76 = v52 + int32(1)
	if v76 != v21 {
		v52 = v76
		v53 = v74
		goto L17
	} else {
		goto L26
	}
L26:
	;
	goto L18
L27:
	;
	goto L1
L28:
	;
	v99 = v5
	goto L2
L29:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v125
	F_errmsg_internal(m, int32(_a_F_ExecLookupResultRelByOid_0), v15)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_ExecLookupResultRelByOid_1), int32(_a_F_ExecLookupResultRelByOid_2), int32(_a_F_ExecLookupResultRelByOid_3))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_InitResultRelInfo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int64
	_ = v72
	v6 = int32(0)
	base.MemoryFill(m, l0+int32(24), v6, int32(192))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(388)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v25 = base.B2i32(v20 == int32(1259)) | base.B2i32(v20 == int32(1262))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+49)) = uint8(v25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v28 = F_CopyTriggerDesc(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v28
		if v28 != 0 {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
			v34 = F_palloc0(m, v31*int32(28))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v34
				v39 = F_palloc0(m, v31<<(uint(int32(2))%32))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v39
					if l4 == int32(0) {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+119)))
						if v56 == int32(102) {
							v60 = F_GetFdwRoutineForRelation(m, l1, int32(1))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								v62 = v60
								v63 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v63)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v62
								v72 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v63
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v63
								return
							}
						} else {
							v62 = int32(0)
							v63 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v63)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v63)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v62
							v72 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v63)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v63
							*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v63
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v63)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v63
							return
						}
					} else {
						v45 = F_InstrAlloc(m, v31, l4, int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							v50 = v45
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v50
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+119)))
							if v56 == int32(102) {
								v60 = F_GetFdwRoutineForRelation(m, l1, int32(1))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									v62 = v60
									v63 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v63)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v63
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v63)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v62
									v72 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v72
									*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v63
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v63)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v63
									*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v63
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v63)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v63
									return
								}
							} else {
								v62 = int32(0)
								v63 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v63)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v62
								v72 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v63
								*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v63
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v63)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v63
								return
							}
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(0)
			v50 = v6
			*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v50
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+119)))
			if v56 == int32(102) {
				v60 = F_GetFdwRoutineForRelation(m, l1, int32(1))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					v62 = v60
					v63 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v63)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v63)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v62
					v72 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v72
					*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v72
					*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v63)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v63
					*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v63
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v63)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v63
					return
				}
			} else {
				v62 = int32(0)
				v63 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v63)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v63
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v63
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v63)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v62
				v72 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v72
				*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v72
				*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v63
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v63
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v63)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v63
				*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v63
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v63)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v63
				return
			}
		}
	}
}
