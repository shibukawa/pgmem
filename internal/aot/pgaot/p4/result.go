package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecInitResultTupleSlotTL(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+44))
	v6 = F_ExecTypeFromTLInternal(m, v4, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v6
		F_ExecInitResultSlot(m, l0, l1)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
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
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v167 int32
	_ = v167
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v203 int64
	_ = v203
	v6 = int32(0)
	base.MemoryFill(m, l0+int32(24), v6, int32(192))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v6
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(394)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v32 = base.B2i32(v27 == int32(1259)) | base.B2i32(v27 == int32(1262))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+49)) = uint8(v32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v35 = F_CopyTriggerDesc(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v35
		if v35 != 0 {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
			v40 = F_palloc0_mul(m, int32(28), v39)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v40
				v44 = F_palloc0_mul(m, int32(4), v39)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v44
					if l4 == int32(0) {
						v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+119)))
						if v187 == int32(102) {
							v191 = F_GetFdwRoutineForRelation(m, l1, int32(1))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return
							} else {
								v193 = v191
								v194 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v194)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v194
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v194)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v193
								v203 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v203
								*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v194
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v194)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v194
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v194)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v194
								return
							}
						} else {
							v193 = int32(0)
							v194 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v194)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v194
							*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v194
							*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v194)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v193
							v203 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v203
							*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v203
							*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v203
							*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v203
							*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v203
							*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v203
							*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v203
							*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v203
							*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v203
							*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v194
							*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v194
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v194)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v194
							*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v194
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v194)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v194
							return
						}
					} else {
						v50 = F_palloc0_mul(m, int32(368), v39)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							if v39 <= int32(0) {
							} else {
								v54 = int32(3)
								v55 = v39 & v54
								v56 = int32(1)
								v57 = l4 & v56
								v61 = int32(base.Ui32(l4)>>(uint(v54)%32)) & v56
								v65 = int32(base.Ui32(l4)>>(uint(v56)%32)) & v56
								v66 = int32(0)
								if base.Ui32(int32(4)) <= base.Ui32(v39) {
									v75 = v66
									v82 = v6
									for {
										v86 = v50 + v75*int32(368)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+2)) = uint8(v61)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)) = uint8(v65)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+370)) = uint8(v61)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+369)) = uint8(v65)
										*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v57)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+738)) = uint8(v61)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+737)) = uint8(v65)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+368)) = uint8(v57)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+1106)) = uint8(v61)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+1105)) = uint8(v65)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+736)) = uint8(v57)
										*(*uint8)(unsafe.Add(mBase, uint32(v86)+1104)) = uint8(v57)
										v99 = int32(4)
										v100 = v75 + v99
										v102 = v82 + v99
										if v102 != v39&int32(2147483644) {
											v75 = v100
											v82 = v102
											continue
										} else {
											break
										}
										break
									}
									if v55 == int32(0) {
									} else {
										v110 = v100
										v123 = v110
										v131 = v6
										for {
											v134 = v50 + v123*int32(368)
											*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)) = uint8(v61)
											*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)) = uint8(v65)
											*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v57)
											v138 = int32(1)
											v141 = v131 + v138
											if v141 != v55 {
												v123 = v123 + v138
												v131 = v141
												continue
											} else {
												break
											}
											break
										}
									}
								} else {
									v110 = v66
									v123 = v110
									v131 = v6
									for {
										v134 = v50 + v123*int32(368)
										*(*uint8)(unsafe.Add(mBase, uint32(v134)+2)) = uint8(v61)
										*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)) = uint8(v65)
										*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v57)
										v138 = int32(1)
										v141 = v131 + v138
										if v141 != v55 {
											v123 = v123 + v138
											v131 = v141
											continue
										} else {
											break
										}
										break
									}
								}
							}
							v167 = v50
							*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v167
							v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
							v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+119)))
							if v187 == int32(102) {
								v191 = F_GetFdwRoutineForRelation(m, l1, int32(1))
								mBase = m.M
								v192 = m.ExcPending
								if v192 != 0 {
									return
								} else {
									v193 = v191
									v194 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v194)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v194
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v194
									*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v194)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v193
									v203 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v203
									*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v203
									*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v203
									*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v203
									*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v203
									*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v203
									*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v203
									*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v203
									*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v203
									*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v194
									*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v194
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v194)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v194
									*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v194
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v194)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v194
									return
								}
							} else {
								v193 = int32(0)
								v194 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v194)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v194
								*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v194)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v193
								v203 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v203
								*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v203
								*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v194
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v194)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v194
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v194)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v194
								return
							}
						}
					}
				}
			}
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(0)
			v167 = v6
			*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v167
			v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+119)))
			if v187 == int32(102) {
				v191 = F_GetFdwRoutineForRelation(m, l1, int32(1))
				mBase = m.M
				v192 = m.ExcPending
				if v192 != 0 {
					return
				} else {
					v193 = v191
					v194 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v194)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v194
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v194
					*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v194)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v193
					v203 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v203
					*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v203
					*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v203
					*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v203
					*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v203
					*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v203
					*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v203
					*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v203
					*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v203
					*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v194
					*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v194
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v194)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v194
					*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v194
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v194)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v194
					return
				}
			} else {
				v193 = int32(0)
				v194 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)) = uint8(v194)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v194
				*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v194
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)) = uint16(v194)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v193
				v203 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+36)) = v203
				*(*int64)(unsafe.Add(mBase, uint32(l0)+41)) = v203
				*(*int64)(unsafe.Add(mBase, uint32(l0)+68)) = v203
				*(*int64)(unsafe.Add(mBase, uint32(l0)+76)) = v203
				*(*int64)(unsafe.Add(mBase, uint32(l0)+124)) = v203
				*(*int64)(unsafe.Add(mBase, uint32(l0)+132)) = v203
				*(*int64)(unsafe.Add(mBase, uint32(l0)+152)) = v203
				*(*int64)(unsafe.Add(mBase, uint32(l0)+160)) = v203
				*(*int64)(unsafe.Add(mBase, uint32(l0)+168)) = v203
				*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v194
				*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v194
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+196)) = uint8(v194)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v194
				*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v194
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+188)) = uint8(v194)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v194
				return
			}
		}
	}
}
func F_make_result_safe(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v17 = int32(_a_F_make_result_safe_0)
	if v16&v17 == v17 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L62
	}
L2:
	;
	m.G0 = v14 + int32(16)
	return v219
L3:
	;
	if base.B2i32(base.B2i32(v16 == int32(_a_F_make_result_safe_0))|base.B2i32(v16 == int32(_a_F_make_result_safe_1)) == int32(0))&base.B2i32(v16 != int32(_a_F_make_result_safe_2)) != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v41 <= int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	v32 = F_palloc(m, int32(6))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)) = uint16(v16)
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = int32(24)
	v219 = v32
	goto L2
L9:
	;
	if v98 != 0 {
		goto L24
	} else {
		goto L25
	}
L10:
	;
	v98 = v41
	v100 = v40
	v101 = v39
	v103 = v3
	goto L9
L11:
	;
	goto L12
L12:
	;
	v50 = v39
	v51 = v40
	v52 = v41
	goto L13
L13:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
	if v59 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v98 = int32(0)
	v100 = v40 - v41
	v101 = v39 + v41<<(uint(int32(1))%32)
	v103 = v3
	goto L9
L15:
	;
	v64 = v52
	goto L19
L16:
	;
	goto L17
L17:
	;
	v85 = int32(1)
	if v85 < v52 {
		v50 = v50 + int32(2)
		v51 = v51 - v85
		v52 = v52 - v85
		goto L13
	} else {
		goto L23
	}
L18:
	;
	v98 = v84
	v100 = v51
	v101 = v50
	v103 = base.B2i32(v76 != int32(0))
	goto L9
L19:
	;
	v76 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50+v64<<(uint(int32(1))%32)-int32(2)))))
	if v76 != 0 {
		v84 = v64
		goto L18
	} else {
		goto L21
	}
L20:
	;
	v84 = int32(0)
	goto L18
L21:
	;
	v79 = int32(1)
	if v79 < v64 {
		v64 = v64 - v79
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L14
L24:
	;
	v106 = v16
	goto L26
L25:
	;
	v106 = int32(0)
	goto L26
L26:
	;
	v108 = v98 << (uint(int32(1)) % 32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v98 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v157)+4)) = uint16(v159)
	if v103 != 0 {
		goto L40
	} else {
		goto L41
	}
L28:
	;
	v113 = v100
	goto L30
L29:
	;
	v113 = int32(0)
	goto L30
L30:
	;
	if base.B2i32(int32(63) < v109)|base.B2i32(base.Ui32(int32(127)) < base.Ui32(v113-int32(-64))) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v122 = v108 + int32(6)
	v123 = F_palloc(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v146 = v108 + int32(8)
	v147 = F_palloc(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L7
	} else {
		goto L38
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v122 << (uint(int32(2)) % 32)
	if v106 == int32(_a_F_make_result_safe_3) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v138 = int32(_a_F_make_result_safe_4)
	goto L37
L36:
	;
	v138 = int32(_a_F_make_result_safe_5)
	goto L37
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v157 = v123
	v159 = int32(base.Ui32(v113)>>(uint(int32(25))%32))&int32(64) | (v113&int32(63) | v138) | v141<<(uint(int32(7))%32)
	goto L27
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v146 << (uint(int32(2)) % 32)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*uint16)(unsafe.Add(mBase, uint32(v147)+6)) = uint16(v113)
	v157 = v147
	v159 = v152&int32(_a_F_make_result_safe_6) | v106
	goto L27
L39:
	;
	if v187 == v113 {
		goto L50
	} else {
		goto L51
	}
L40:
	;
	v162 = v98 << (uint(int32(1)) % 32)
	if v162 != 0 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v172 = v159
	goto L42
L42:
	;
	v183 = base.I32_extend16_s(v172)
	if v183 < int32(0) {
		v187 = v172<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v172&int32(63)
		goto L39
	} else {
		goto L49
	}
L43:
	;
	if base.I32_extend16_s(v159) < int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v157)+4)))
	v172 = v171
	goto L42
L46:
	;
	v168 = int32(6)
	goto L48
L47:
	;
	v168 = int32(8)
	goto L48
L48:
	;
	base.MemoryCopy(m, v157+v168, v101, v162)
	goto L45
L49:
	;
	v186 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+6)))
	v187 = v186
	goto L39
L50:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if int32(0) <= v183 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v200 = int32(0)
	v201 = F_errsave_start(m, l1)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L7
	} else {
		goto L57
	}
L53:
	;
	v198 = v172 & int32(_a_F_make_result_safe_6)
	goto L55
L54:
	;
	v198 = int32(base.Ui32(v172)>>(uint(int32(7))%32)) & int32(63)
	goto L55
L55:
	;
	if v189 == v198 {
		v219 = v157
		goto L2
	} else {
		goto L56
	}
L56:
	;
	goto L52
L57:
	;
	if v201 == int32(0) {
		v219 = v200
		goto L2
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L7
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_make_result_safe_7), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	F_errsave_finish(m, l1, int32(_a_F_make_result_safe_8), int32(_a_F_make_result_safe_9), int32(_a_F_make_result_safe_10))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L61
	}
L61:
	;
	v219 = v200
	goto L2
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v16
	F_errmsg_internal(m, int32(_a_F_make_result_safe_11), v14)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_make_result_safe_8), int32(_a_F_make_result_safe_12), int32(_a_F_make_result_safe_10))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
