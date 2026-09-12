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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v97 int32
	_ = v97
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
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
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L29
	}
L2:
	;
	m.G0 = v15 + int32(16)
	return v97
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
		v97 = int32(0)
		goto L2
	} else {
		goto L27
	}
L5:
	;
	v52 = int32(0)
	v54 = v5
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
	v28 = int32(0)
	v30 = F_hash_search(m, v18, v15+int32(12), v28, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
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
	if v30 == int32(0) {
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v97 = v44 + v43*int32(216)
	goto L2
L14:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v43 = v38
	goto L13
L15:
	;
	goto L16
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v41
	v43 = v41
	goto L13
L17:
	;
	v62 = v24 + v52*int32(216)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+56))
	v65 = base.B2i32(v64 == l1)
	if v64 == l1 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L4
L19:
	;
	v66 = v62
	goto L21
L20:
	;
	v66 = v54
	goto L21
L21:
	;
	if l3 == int32(0) {
		v72 = v66
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v64 == l1 {
		v97 = v72
		goto L2
	} else {
		goto L25
	}
L23:
	;
	if l1 != v64 {
		v72 = v66
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = l1
	v72 = v62
	goto L22
L25:
	;
	v74 = v52 + int32(1)
	if v74 != v21 {
		v52 = v74
		v54 = v72
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
	v97 = v5
	goto L2
L29:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v123
	F_errmsg_internal(m, int32(56270), v15)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(500592), int32(4623), int32(438833))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
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
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
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
	var v98 int64
	_ = v98
	if l0&int32(3) == int32(0) {
		v13 = l0 + int32(216)
		if base.Ui32(v13) <= base.Ui32(l0) {
		} else {
			v19 = l0 + int32(4)
			if base.Ui32(v19) < base.Ui32(v13) {
				v21 = v13
			} else {
				v21 = v19
			}
			v28 = F__emscripten_memset_bulkmem(m, l0, base.I32_extend8_s(int32(0)), (l0^int32(-1)+v21)&int32(-4)+int32(4))
			mBase = m.M
		}
	} else {
		v34 = F__emscripten_memset_bulkmem(m, l0+int32(24), base.I32_extend8_s(int32(0)), int32(192))
		mBase = m.M
	}
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(388)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v51 = base.B2i32(v46 == int32(1259)) | base.B2i32(v46 == int32(1262))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+49)) = uint8(v51)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v54 = F_CopyTriggerDesc(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v54
		if v54 != 0 {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
			v60 = F_palloc0(m, v57*int32(28))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v60
				v65 = F_palloc0(m, v57<<(uint(int32(2))%32))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v65
					if l4 == int32(0) {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
						if v82 == int32(102) {
							v86 = F_GetFdwRoutineForRelation(m, l1, int32(1))
							mBase = m.M
							v87 = m.ExcPending
							if v87 != 0 {
								return
							} else {
								v88 = v86
								v89 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v89)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v89)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v88
								v98 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v89
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v89)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v89
								*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v89
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v89)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v89
								return
							}
						} else {
							v88 = int32(0)
							v89 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v89)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v89
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v89)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v88
							v98 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v98
							*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v98
							*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v98
							*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v98
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v89
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v89)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v89
							*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v98
							*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v98
							*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v98
							*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v98
							*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v98
							*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v89
							*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v89
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v89)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v89
							return
						}
					} else {
						v71 = F_InstrAlloc(m, v57, l4, int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							v76 = v71
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v76
							v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
							v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
							if v82 == int32(102) {
								v86 = F_GetFdwRoutineForRelation(m, l1, int32(1))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return
								} else {
									v88 = v86
									v89 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v89)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v89
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v89)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v88
									v98 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v98
									*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v98
									*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v98
									*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v89
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v89)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v89
									*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v98
									*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v98
									*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v98
									*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v98
									*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v98
									*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v89
									*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v89
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v89)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v89
									return
								}
							} else {
								v88 = int32(0)
								v89 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v89)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v89)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v88
								v98 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v89
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v89)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v89
								*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v98
								*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v98
								*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v89
								*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v89
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v89)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v89
								return
							}
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(0)
			v76 = v37
			*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v76
			v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+119)))
			if v82 == int32(102) {
				v86 = F_GetFdwRoutineForRelation(m, l1, int32(1))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return
				} else {
					v88 = v86
					v89 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v89)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v89
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v89)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v88
					v98 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v98
					*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v98
					*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v98
					*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v98
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v89
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v89)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v89
					*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v98
					*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v98
					*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v98
					*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v98
					*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v98
					*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v89
					*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v89
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v89)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v89
					return
				}
			} else {
				v88 = int32(0)
				v89 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v89)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v89
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v89
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v89)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v88
				v98 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v98
				*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v98
				*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v98
				*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v98
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v89
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v89)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v89
				*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v98
				*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v98
				*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v98
				*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v98
				*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v98
				*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v89
				*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v89
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v89)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v89
				return
			}
		}
	}
}
