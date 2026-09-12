package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSimpleRelationUpdate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_CheckCmdReplicaIdentity(m, v15, int32(2))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = l3 + int32(28)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v21 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(16)
	return
L4:
	;
	v35 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)) = uint8(v35)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v38 == v35 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+13)))
	if v24 != int32(1) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(0)
	v31 = F_ExecBRUpdateTriggers(m, l1, l2, l0, v20, v27, l4, v27, v27, v27)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v31 == int32(0) {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+131)))
	if v54 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L10:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+17)))
	if v41 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_ExecComputeStoredGenerated(m, l0, l1, l4, int32(2))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_ExecConstraints(m, l0, l4, l1)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	if v48 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	goto L9
L17:
	;
	v58 = F_ExecPartitionCheck(m, l0, l4, l1, int32(1))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v63 = m.G0
	v65 = v63 - int32(32)
	m.G0 = v65
	v68 = F_GetCurrentCommandId(m, int32(1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v15)+188))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+100))
	v78 = m.T0[v77].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) int32)(m, v15, v20, l4, v68, v60, int32(0), int32(1), v65+int32(12), v65+int32(8), v13+int32(12))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	if v78 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	switch v78 - int32(2) {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	default:
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	m.G0 = v65 + int32(32)
	v137 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v138 <= v137 {
		v162 = v137
		goto L42
	} else {
		goto L43
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L39
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L36
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L33
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg_internal(m, int32(339834), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(498728), int32(355), int32(357075))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errmsg_internal(m, int32(449910), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(498728), int32(363), int32(357075))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errmsg_internal(m, int32(448119), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(498728), int32(367), int32(357075))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v78
	F_errmsg_internal(m, int32(58101), v65)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(498728), int32(371), int32(357075))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v164 = int32(0)
	F_ExecARUpdateTriggers(m, l1, l0, v164, v164, v20, v164, l4, v162, v164, v164)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L48
	}
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v141 == int32(0) {
		v162 = v137
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v152 = F_ExecInsertIndexTuples(m, l0, l4, l1, int32(1), base.B2i32(v145 != int32(0)), v13+int32(11), v145, base.B2i32(v141 == int32(2)))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+11)))
	if v154&int32(1) == int32(0) {
		v162 = v152
		goto L42
	} else {
		goto L46
	}
L46:
	;
	F_CheckAndReportConflict(m, l0, l1, int32(2), v152, l3, l4)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v162 = v152
	goto L42
L48:
	;
	F_list_free(m, v162)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	goto L3
}
func F_assign_simple_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v4 = l3
	if v4 != 0 {
		v39 = l2
		v40 = l4
		v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
		if v43 != int32(1) {
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
			return
		} else {
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
			if v46 != 0 {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
				F_pfree(m, v61)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
					return
				}
			} else {
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
				v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
				if v48 != int32(65535) {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					F_pfree(m, v61)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
						return
					}
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
					if v52 != int32(1) {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						F_pfree(m, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						}
					} else {
						v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
						if v55 != int32(3) {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							F_pfree(m, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
								return
							}
						} else {
							F_DeleteExpandedObject(m, v51)
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
								return
							}
						}
					}
				}
			}
		}
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+31)))
		if v8&int32(1) != 0 {
			v39 = l2
			v40 = l4
			v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
			if v43 != int32(1) {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
				return
			} else {
				v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
				if v46 != 0 {
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
					F_pfree(m, v61)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
						return
					}
				} else {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
					v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
					if v48 != int32(65535) {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						F_pfree(m, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						}
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
						if v52 != int32(1) {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							F_pfree(m, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
								return
							}
						} else {
							v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
							if v55 != int32(3) {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								F_pfree(m, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
									return
								}
							} else {
								F_DeleteExpandedObject(m, v51)
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
									return
								}
							}
						}
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
			v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)))
			if v12 != int32(65535) {
				v39 = l2
				v40 = l4
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
				if v43 != int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
					return
				} else {
					v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
					if v46 != 0 {
						v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						F_pfree(m, v61)
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						}
					} else {
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
						v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
						if v48 != int32(65535) {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							F_pfree(m, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
								return
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							if v52 != int32(1) {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								F_pfree(m, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
									return
								}
							} else {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
								if v55 != int32(3) {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									F_pfree(m, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
										return
									}
								} else {
									F_DeleteExpandedObject(m, v51)
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
										return
									}
								}
							}
						}
					}
				}
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
				if v15 != int32(1) {
					v39 = l2
					v40 = l4
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
					if v43 != int32(1) {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
						return
					} else {
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
						if v46 != 0 {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
							F_pfree(m, v61)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
								return
							}
						} else {
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
							v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
							if v48 != int32(65535) {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								F_pfree(m, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
									return
								}
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
								if v52 != int32(1) {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									F_pfree(m, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
										return
									}
								} else {
									v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
									if v55 != int32(3) {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
										F_pfree(m, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
											return
										}
									} else {
										F_DeleteExpandedObject(m, v51)
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
											return
										}
									}
								}
							}
						}
					}
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
					if v18&int32(254) == int32(2) {
						v39 = l2
						v40 = l4
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
						if v43 != int32(1) {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						} else {
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
							if v46 != 0 {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
								F_pfree(m, v61)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
									return
								}
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
								v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
								if v48 != int32(65535) {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									F_pfree(m, v61)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
										return
									}
								} else {
									v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
									v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
									if v52 != int32(1) {
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
										F_pfree(m, v61)
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
											return
										}
									} else {
										v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
										if v55 != int32(3) {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
											F_pfree(m, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
												*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
												*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
												return
											}
										} else {
											F_DeleteExpandedObject(m, v51)
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
												*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
												*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
												return
											}
										}
									}
								}
							}
						}
					} else {
						v23 = int32(4520560)
						v24 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v26
						v28 = F_detoast_external_attr(m, l2)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v24
							if l4 != 0 {
								F_pfree(m, l2)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return
								} else {
									v36 = F_datumCopy(m, v28, int32(0), int32(-1))
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										v39 = v36
										v40 = int32(1)
										v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
										if v43 != int32(1) {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
											return
										} else {
											v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
											if v46 != 0 {
												v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
												F_pfree(m, v61)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
													*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
													*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
													return
												}
											} else {
												v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
												v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
												if v48 != int32(65535) {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
													F_pfree(m, v61)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
														*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
														*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
														return
													}
												} else {
													v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
													v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
													if v52 != int32(1) {
														v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
														F_pfree(m, v61)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
															*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
															*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
															return
														}
													} else {
														v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
														if v55 != int32(3) {
															v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
															F_pfree(m, v61)
															mBase = m.M
															v63 = m.ExcPending
															if v63 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
																*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
																*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
																return
															}
														} else {
															F_DeleteExpandedObject(m, v51)
															mBase = m.M
															v59 = m.ExcPending
															if v59 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
																*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
																*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
																return
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v36 = F_datumCopy(m, v28, int32(0), int32(-1))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return
								} else {
									v39 = v36
									v40 = int32(1)
									v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)))
									if v43 != int32(1) {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
										return
									} else {
										v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)))
										if v46 != 0 {
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
											F_pfree(m, v61)
											mBase = m.M
											v63 = m.ExcPending
											if v63 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
												*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
												*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
												return
											}
										} else {
											v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
											v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+12)))
											if v48 != int32(65535) {
												v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
												F_pfree(m, v61)
												mBase = m.M
												v63 = m.ExcPending
												if v63 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
													*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
													*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
													return
												}
											} else {
												v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
												v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
												if v52 != int32(1) {
													v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
													F_pfree(m, v61)
													mBase = m.M
													v63 = m.ExcPending
													if v63 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
														*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
														*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
														return
													}
												} else {
													v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
													if v55 != int32(3) {
														v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
														F_pfree(m, v61)
														mBase = m.M
														v63 = m.ExcPending
														if v63 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
															*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
															*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
															return
														}
													} else {
														F_DeleteExpandedObject(m, v51)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l1)+45)) = uint8(v40)
															*(*uint8)(unsafe.Add(mBase, uint32(l1)+44)) = uint8(v4)
															*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
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
		}
	}
}
func F_exec_is_simple_query(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 == v2 {
		v53 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		if v8 != int32(1) {
			v53 = v2
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
			if v13 == int32(0) {
				v53 = v2
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				if v16 != int32(1) {
					v53 = v2
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
					if v21 != int32(67) {
						v53 = v2
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
						if v24 != int32(1) {
							v53 = v2
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
							if v27 != 0 {
								v53 = v2
							} else {
								v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+36)))
								if v28 != 0 {
									v53 = v2
								} else {
									v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+37)))
									if v29 != 0 {
										v53 = v2
									} else {
										v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+38)))
										if v30 != 0 {
											v53 = v2
										} else {
											v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+39)))
											if v31 != 0 {
												v53 = v2
											} else {
												v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
												if v32 != 0 {
													v53 = v2
												} else {
													v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
													v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
													if v34 != 0 {
														v53 = v2
													} else {
														v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
														if v35 != 0 {
															v53 = v2
														} else {
															v36 = *(*int32)(unsafe.Add(mBase, uint32(v20)+100))
															if v36 != 0 {
																v53 = v2
															} else {
																v37 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
																if v37 != 0 {
																	v53 = v2
																} else {
																	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+112))
																	if v38 != 0 {
																		v53 = v2
																	} else {
																		v39 = *(*int32)(unsafe.Add(mBase, uint32(v20)+116))
																		if v39 != 0 {
																			v53 = v2
																		} else {
																			v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+120))
																			if v40 != 0 {
																				v53 = v2
																			} else {
																				v41 = *(*int32)(unsafe.Add(mBase, uint32(v20)+124))
																				if v41 != 0 {
																					v53 = v2
																				} else {
																					v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)+128))
																					if v42 != 0 {
																						v53 = v2
																					} else {
																						v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+132))
																						if v43 != 0 {
																							v53 = v2
																						} else {
																							v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+144))
																							if v44 != 0 {
																								v53 = v2
																							} else {
																								v45 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
																								if v45 == int32(0) {
																									v53 = v2
																								} else {
																									v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
																									v53 = base.B2i32(v48 == int32(1))
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
	return v53
}
func F_exec_save_simple_expr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v16 = v12 + int32(36)
	goto L3
L1:
	;
	F_errstart_cold(m, int32(21), int32(557877))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L13
	}
L2:
	;
	v35 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v35
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v35)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v24
	v42 = F_exprType(m, v24)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	switch v25 - int32(360) {
	case 0, 8:
		goto L6
	case 1, 2, 3, 4, 5, 6, 7:
		goto L1
	default:
		goto L5
	}
L4:
	;
	if v25 != int32(331) {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	goto L4
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 == int32(7) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v16 = v20 + int32(52)
	goto L3
L8:
	;
	goto L2
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v42
	v45 = F_exprTypmod(m, v24)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v45
	v48 = F_contain_mutable_functions(m, v24)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v48)
	m.G0 = v8 + int32(16)
	return
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v60
	F_errmsg_internal(m, int32(487239), v8)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(501532), int32(8334), int32(207579))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_isSimpleNode(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var __phi10 int32
	_ = __phi10
	var v11 int32
	_ = v11
	var __phi11 int32
	_ = __phi11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v70 int64
	_ = v70
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	v4 = int32(0)
	if l0 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	__phi10 = l0
	__phi11 = l1
	v10 = __phi10
	v11 = __phi11
	goto L9
L3:
	;
	return v128
L4:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v128 = base.B2i32(base.Ui32(v122-int32(4)) < base.Ui32(int32(-3)))
	goto L3
L5:
	;
	return int32(1)
L6:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	switch v102 - int32(9) {
	case 0, 1, 2, 5, 10, 23, 26, 27, 29, 30, 32, 39:
		v128 = int32(1)
		goto L3
	default:
		goto L1
	case 6:
		goto L4
	case 12:
		goto L34
	}
L7:
	;
	switch v91 - int32(9) {
	case 0, 1, 2, 5, 10, 12, 23, 26, 27, 29, 30, 32:
		v128 = int32(1)
		goto L3
	default:
		goto L1
	case 6:
		goto L33
	}
L8:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v91 = v90
	goto L7
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	switch v18 - int32(6) {
	case 0, 1, 2, 3, 4, 5, 7, 8, 9, 13, 26, 29, 30, 32, 33, 34, 35, 39, 42, 50, 51, 52, 53:
		goto L5
	default:
		v128 = v4
		goto L3
	case 11:
		goto L11
	case 12, 16, 40, 46, 47:
		goto L8
	case 15:
		goto L6
	case 19:
		goto L15
	case 20:
		goto L14
	case 21, 22, 23, 24, 38, 49:
		v30 = int32(4)
		goto L12
	case 55:
		goto L13
	}
L10:
	;
	if l2&int32(1) == int32(0) {
		goto L8
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30+v10)))
	if v32 != 0 {
		__phi10 = v32
		__phi11 = v10
		v10 = __phi10
		v11 = __phi11
		goto L9
	} else {
		goto L16
	}
L13:
	;
	v30 = int32(12)
	goto L12
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	return base.B2i32(v25 != int32(26))
L15:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	return base.B2i32(v21 != int32(25))
L16:
	;
	goto L1
L17:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v37 != int32(17) {
		v91 = v37
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v40 = F_get_simple_binary_op_name(m, v10)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v40 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v46 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if base.Ui64(int64(63)) < base.Ui64(v46) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v50 = int64(1) << (uint(v46) % 64)
	if v50&int64(189253438930945) == int64(0) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v55 = F_get_simple_binary_op_name(m, v11)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	if v55 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v59 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if base.Ui64(int64(63)) < base.Ui64(v59) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v63 = int64(1) << (uint(v59) % 64)
	if v63&int64(189253438930945) == int64(0) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v70 = int64(0)
	if base.B2i32(v50&int64(145272973819905) != v70)&base.B2i32(v63&int64(43980465111041) != v70) != 0 {
		goto L5
	} else {
		goto L28
	}
L28:
	;
	if v50&int64(43980465111041) != int64(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if v63&int64(145272973819905) != int64(0) {
		v128 = v4
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	return base.B2i32(v10 == v87)
L32:
	;
	goto L31
L33:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	return base.B2i32(base.Ui32(v95-int32(4)) < base.Ui32(int32(-3)))
L34:
	;
	if l2&int32(1) == int32(0) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v110 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	switch v111 {
	case 0, 2:
		goto L37
	case 1:
		goto L36
	default:
		v128 = v110
		goto L3
	}
L36:
	;
	if v109 != int32(1) {
		v128 = v110
		goto L3
	} else {
		goto L39
	}
L37:
	;
	if base.Ui32(v109) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v128 = v110
	goto L3
L39:
	;
	goto L5
}
func F_is_simple_subquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6 != int32(67) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L24
	} else {
		goto L47
	}
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v9 != int32(1) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+144))
	if v12 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v114
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)))
	if v13 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+37)))
	if v14 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)))
	if v15 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+100))
	if v16 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v17 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if v18 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	if v19 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+120))
	if v20 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v21 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+132))
	if v22 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)))
	if v23 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v24 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+40)))
	if v25 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L18
	}
L18:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+124)))
	if v26 != int32(1) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v109 = F_contain_volatile_functions(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L24
	} else {
		goto L46
	}
L20:
	;
	if l3 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v32 = int32(0)
	v34 = F_jointree_contains_lateral_outer_refs(m, l0, v31, v32, v32)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v40 = int32(1)
	v42 = F_get_relids_in_jointree(m, l3, v40, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L24
	} else {
		goto L27
	}
L24:
	;
	return int32(0)
L25:
	;
	if v34 == int32(0) {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	v114 = v5
	goto L4
L27:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v46 = F_jointree_contains_lateral_outer_refs(m, l0, v44, int32(1), v42)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L24
	} else {
		goto L28
	}
L28:
	;
	if v46 != 0 {
		v114 = v5
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	v49 = F_pull_varnos_of_level(m, l0, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v51 = int32(0)
	if v49 == v51 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v104 == int32(0) {
		v114 = v5
		goto L4
	} else {
		goto L45
	}
L32:
	;
	v104 = int32(1)
	goto L31
L33:
	;
	goto L34
L34:
	;
	if v42 == int32(0) {
		v95 = v51
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v104 = v95
	goto L31
L36:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v61 < v60 {
		v95 = v51
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v63 = int32(1)
	if v60 <= v63 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v66 = v63
	goto L40
L39:
	;
	v66 = v60
	goto L40
L40:
	;
	v67 = int32(8)
	v72 = int32(0)
	goto L41
L41:
	;
	v79 = v72 << (uint(int32(2)) % 32)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v49+v67+v79)))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79+(v42+v67))))
	v86 = v81 & (v83 ^ int32(-1))
	v88 = base.B2i32(v86 == int32(0))
	if v86 != 0 {
		v95 = v88
		goto L35
	} else {
		goto L43
	}
L42:
	;
	v95 = v88
	goto L35
L43:
	;
	v90 = v72 + int32(1)
	if v90 != v66 {
		v72 = v90
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	goto L19
L46:
	;
	v114 = v109 ^ int32(1)
	goto L4
L47:
	;
	F_errmsg_internal(m, int32(115620), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L24
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(500891), int32(1815), int32(15389))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L24
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
