package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_ParameterAclLookup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_convert_GUC_name_for_parameter_acl(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_cstring_to_text(m, v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = int64(0)
			v20 = F_GetSysCacheOid(m, int32(43), base.I64_extend_i32_u(v14), v17, v17, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				if l1|v20 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(67137668))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg(m, int32(_a_F_ParameterAclLookup_0), v7)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ParameterAclLookup_1), int32(51), int32(_a_F_ParameterAclLookup_2))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_pfree(m, v10)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v20
					}
				}
			}
		}
	}
}
func F_PrepareSortSupportFromIndexRel(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v2 = l1
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+10)))
	if v12 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
			F_errmsg_internal(m, int32(_a_F_PrepareSortSupportFromIndexRel_0), v9)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_PrepareSortSupportFromIndexRel_1), int32(170), int32(_a_F_PrepareSortSupportFromIndexRel_2))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v30 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+10)))
		v34 = v30<<(uint(int32(2))%32) - int32(4)
		v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v34+v35)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v38+v34)))
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v2)
		F_FinishSortSupportFunction(m, v37, v40, l2)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_PrepareSortSupportFromOrderingOp(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v14 = F_get_ordering_op_properties(m, l0, v6+int32(12), v6+int32(8), v6+int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
				F_errmsg_internal(m, int32(_a_F_PrepareSortSupportFromOrderingOp_0), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PrepareSortSupportFromOrderingOp_1), int32(146), int32(_a_F_PrepareSortSupportFromOrderingOp_2))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+8)) = uint8(base.B2i32(v31 == int32(5)))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			F_FinishSortSupportFunction(m, v35, v36, l1)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
func F_p_isURLPath(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	v2 = int32(0)
	v7 = F_palloc0(m, int32(40))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v13 + v15
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 - v20
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v23 + v25<<(uint(int32(2))%32)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v31 = F_palloc(m, int32(32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v31)+8)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v31))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v46 = F_palloc(m, int32(32))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v46)+20)) = int32(57)
	F_check_stack_depth(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L12
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v44)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v46)+8)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = v52
	goto L8
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = int32(0)
	v56 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v46)+8)) = v56
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = v56
	goto L8
L12:
	;
	v68 = F_TParserGet(m, v7)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v102 != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	if v68 == int32(0) {
		v101 = v2
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
	if v72 != int32(18) {
		v101 = v2
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76 + v77
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = v81 + v82
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v86 + v87
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v90)+16)) = v91 + v92
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v95)+8)) = v97
	v101 = int32(1)
	goto L13
L17:
	;
	v103 = v102
	goto L20
L18:
	;
	goto L19
L19:
	;
	F_pfree(m, v7)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L24
	}
L20:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+24))
	F_pfree(m, v103)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	goto L19
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v108
	if v108 != 0 {
		v103 = v108
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	return v101
}
func F_packArcInfoCmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6 < v7 {
		return int32(-1)
	} else {
		v11 = int32(1)
		if v7 < v6 {
			v28 = v11
			return v28
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			if v13 < v14 {
				return int32(-1)
			} else {
				if v14 < v13 {
					v28 = v11
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v20 < v21 {
						v28 = int32(-1)
					} else {
						v28 = base.B2i32(v21 < v20)
					}
				}
				return v28
			}
		}
	}
}
func F_pagetable_delete(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var __phi52 int32
	_ = __phi52
	var v54 int32
	_ = v54
	var __phi54 int32
	_ = __phi54
	var v55 int32
	_ = v55
	var __phi55 int32
	_ = __phi55
	var v56 int32
	_ = v56
	var __phi56 int32
	_ = __phi56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	v8 = int32(16)
	v12 = (int32(base.Ui32(l1)>>(uint(v8)%32)) ^ l1) * int32(-2048144789)
	v17 = (int32(base.Ui32(v12)>>(uint(int32(13))%32)) ^ v12) * int32(-1028477387)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = int32(base.Ui32(v17)>>(uint(v8)%32)) ^ v17
	goto L1
L1:
	;
	v30 = v26 & v22
	v33 = v21 + v30*int32(48)
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
	switch v34 {
	case 0:
		v112 = int32(0)
		goto L4
	case 1:
		goto L5
	default:
		goto L3
	}
L3:
	;
	v26 = v30 + int32(1)
	goto L1
L4:
	;
	return v112
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v35 != l1 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v37 - v38
	v44 = v22 & (v30 + v38)
	v47 = v21 + v44*int32(48)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)))
	if v48 != v38 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v104 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+4)) = uint8(v104)
	v112 = v38
	goto L4
L8:
	;
	v99 = v33
	goto L7
L9:
	;
	goto L10
L10:
	;
	__phi52 = v44
	__phi54 = v33
	__phi55 = v47
	__phi56 = v22
	v52 = __phi52
	v54 = __phi54
	v55 = __phi55
	v56 = __phi56
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v59 = int32(16)
	v63 = (int32(base.Ui32(v58)>>(uint(v59)%32)) ^ v58) * int32(-2048144789)
	v68 = (int32(base.Ui32(v63)>>(uint(int32(13))%32)) ^ v63) * int32(-1028477387)
	if v52 == (int32(base.Ui32(v68)>>(uint(v59)%32))^v68)&v56 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v99 = v55
	goto L7
L13:
	;
	v99 = v54
	goto L7
L14:
	;
	goto L15
L15:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(v55)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+40)) = v74
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v55)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+32)) = v76
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v55)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+24)) = v78
	v80 = *(*int64)(unsafe.Add(mBase, uint32(v55)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+16)) = v80
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v55)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v55)))
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v88 = int32(1)
	v90 = v87 & (v52 + v88)
	v93 = v86 + v90*int32(48)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)))
	if v94 == v88 {
		__phi52 = v90
		__phi54 = v55
		__phi55 = v93
		__phi56 = v87
		v52 = __phi52
		v54 = __phi54
		v55 = __phi55
		v56 = __phi56
		goto L11
	} else {
		goto L16
	}
L16:
	;
	goto L12
}
func F_pagetable_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v68 int64
	_ = v68
	var v73 int32
	_ = v73
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v114 int64
	_ = v114
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v236 int64
	_ = v236
	var v238 int64
	_ = v238
	var v240 int64
	_ = v240
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v372 int64
	_ = v372
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int64
	_ = v420
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v426 int64
	_ = v426
	var v428 int64
	_ = v428
	var v430 int64
	_ = v430
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int64
	_ = v453
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v528 int32
	_ = v528
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	v15 = int32(16)
	v19 = (int32(base.Ui32(l1)>>(uint(v15)%32)) ^ l1) * int32(-2048144789)
	v24 = (int32(base.Ui32(v19)>>(uint(int32(13))%32)) ^ v19) * int32(-1028477387)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v34 = base.B2i32(base.Ui32(v28) < base.Ui32(v29))
	goto L1
L1:
	;
	if v34 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L26
	} else {
		goto L100
	}
L3:
	;
	goto L2
L4:
	;
	v528 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v528
	v34 = v528
	goto L1
L5:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v506)
	return v505
L6:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v491 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v490 + v491
	*(*uint8)(unsafe.Add(mBase, uint32(v483)+4)) = uint8(v491)
	*(*int32)(unsafe.Add(mBase, uint32(v483))) = l1
	v505 = v483
	v506 = int32(0)
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L26
	} else {
		goto L97
	}
L8:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v47 == int64(4294967296) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v299 = (int32(base.Ui32(v24)>>(uint(v15)%32)) ^ v24) & v298
	v302 = v297 + v299*int32(48)
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v302)+4)))
	if v303 == int32(0) {
		v483 = v302
		goto L6
	} else {
		goto L66
	}
L11:
	;
	v50 = int32(0)
	v52 = int64(2)
	v54 = v47 << (uint(int64(1)) % 64)
	if base.Ui64(v54) <= base.Ui64(v52) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v34 = int32(1)
	goto L1
L13:
	;
	v57 = v52
	goto L15
L14:
	;
	v57 = v54
	goto L15
L15:
	;
	v58 = int64(1)
	if v57&(v57-v58) == int64(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v68 = v57
	goto L18
L17:
	;
	v68 = v58 << (uint(int64(64)-base.I64_clz(v57)) % 64)
	goto L18
L18:
	;
	if base.Ui64(v68*int64(48)) < base.Ui64(int64(2147483647)) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v77 = base.I32_wrap_i64(v68) * int32(48)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+112))
	if v79 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	goto L3
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v102
	v104 = int64(1)
	if v68&(v68-v104) == int64(0) {
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v84 = F_MemoryContextAllocExtended(m, v82, v77, int32(5))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v78)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v78)+100)) = v88
	v93 = F_dsa_allocate_extended(m, v79, v77|int32(4), int32(5))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L26
	} else {
		goto L28
	}
L26:
	;
	return int32(0)
L27:
	;
	v102 = v84
	goto L22
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78)+96)) = v93
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v78)+112))
	v97 = F_dsa_get_address(m, v96, v93)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v102 = v97 + int32(4)
	goto L22
L30:
	;
	v114 = v68
	goto L32
L31:
	;
	v114 = v104 << (uint(int64(64)-base.I64_clz(v68)) % 64)
	goto L32
L32:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v114*int64(48)) {
		goto L3
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v114
	v122 = base.I32_wrap_i64(v114) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v122
	if v114 == int64(4294967296) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v131 = int32(-85899346)
	goto L36
L35:
	;
	v131 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v114), float64(0.9)))
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v131
	if v74 != int64(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v141 = v50
	goto L41
L38:
	;
	goto L39
L39:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)+112))
	if v285 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L40:
	;
	v184 = v177
	v188 = v50
	goto L46
L41:
	;
	v151 = v73 + v141*int32(48)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+4)))
	if v152 != int32(1) {
		v177 = v141
		goto L40
	} else {
		goto L43
	}
L42:
	;
	v177 = int32(0)
	goto L40
L43:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	v156 = int32(16)
	v160 = (int32(base.Ui32(v155)>>(uint(v156)%32)) ^ v155) * int32(-2048144789)
	v165 = (int32(base.Ui32(v160)>>(uint(int32(13))%32)) ^ v160) * int32(-1028477387)
	if (int32(base.Ui32(v165)>>(uint(v156)%32))^v165)&v122 == v141 {
		v177 = v141
		goto L40
	} else {
		goto L44
	}
L44:
	;
	v172 = v141 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v172)) < base.Ui64(v74) {
		v141 = v172
		goto L41
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	v194 = v73 + v184*int32(48)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+4)))
	if v195 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L39
L48:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v199 = int32(16)
	v203 = (int32(base.Ui32(v198)>>(uint(v199)%32)) ^ v198) * int32(-2048144789)
	v208 = (int32(base.Ui32(v203)>>(uint(int32(13))%32)) ^ v203) * int32(-1028477387)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v218 = int32(base.Ui32(v208)>>(uint(v199)%32)) ^ v208
	goto L51
L49:
	;
	goto L50
L50:
	;
	v261 = v184 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v261)) < base.Ui64(v74) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v227 = v212 & v218
	v232 = v102 + v227*int32(48)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+4)))
	if v233 != 0 {
		v218 = v227 + int32(1)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v194)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+40)) = v234
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v194)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+32)) = v236
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v194)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+24)) = v238
	v240 = *(*int64)(unsafe.Add(mBase, uint32(v194)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+16)) = v240
	v242 = *(*int64)(unsafe.Add(mBase, uint32(v194)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = v242
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v194)))
	*(*int64)(unsafe.Add(mBase, uint32(v232))) = v244
	goto L50
L53:
	;
	goto L52
L54:
	;
	v265 = v261
	goto L56
L55:
	;
	v265 = int32(0)
	goto L56
L56:
	;
	v267 = v188 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v267)) < base.Ui64(v74) {
		v184 = v265
		v188 = v267
		goto L46
	} else {
		goto L57
	}
L57:
	;
	goto L47
L58:
	;
	F_pfree(m, v73)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L26
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v284)+100))
	if v290 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L12
L62:
	;
	F_dsa_free(m, v285, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L26
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L12
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+100)) = int32(0)
	goto L64
L66:
	;
	v311 = v299
	v314 = int32(0)
	v315 = v302
	goto L67
L67:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	if v322 == l1 {
		v505 = v315
		v506 = int32(1)
		goto L5
	} else {
		goto L69
	}
L68:
	;
	v483 = v461
	goto L6
L69:
	;
	v325 = v311 + int32(1)
	v326 = int32(16)
	v330 = (int32(base.Ui32(v322)>>(uint(v326)%32)) ^ v322) * int32(-2048144789)
	v335 = (int32(base.Ui32(v330)>>(uint(int32(13))%32)) ^ v330) * int32(-1028477387)
	v339 = (int32(base.Ui32(v335)>>(uint(v326)%32)) ^ v335) & v298
	if base.Ui32(v311) < base.Ui32(v339) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v343 = v311 + v341
	goto L72
L71:
	;
	v343 = v311
	goto L72
L72:
	;
	if base.Ui32(v343-v339) < base.Ui32(v314) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v346 = v325 & v298
	v349 = v297 + v346*int32(48)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+4)))
	if v350 != 0 {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v448 = v314 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v448) {
		goto L92
	} else {
		goto L93
	}
L76:
	;
	v358 = int32(0)
	v360 = v346
	goto L79
L77:
	;
	v389 = v349
	v392 = v346
	goto L78
L78:
	;
	if v311 != v392 {
		goto L86
	} else {
		goto L87
	}
L79:
	;
	v367 = v358 + int32(1)
	if int32(151) <= v367 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v389 = v382
	v392 = v379
	goto L78
L81:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v372 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v370), base.F64_convert_i64_u(v372)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v379 = (v360 + int32(1)) & v298
	v382 = v297 + v379*int32(48)
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+4)))
	if v383 != 0 {
		v358 = v367
		v360 = v379
		goto L79
	} else {
		goto L85
	}
L84:
	;
	goto L83
L85:
	;
	goto L80
L86:
	;
	v404 = v389
	v407 = v392
	goto L89
L87:
	;
	goto L88
L88:
	;
	v483 = v315
	goto L6
L89:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v416 = v413 & (v407 - int32(1))
	v419 = v297 + v416*int32(48)
	v420 = *(*int64)(unsafe.Add(mBase, uint32(v419)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+40)) = v420
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v419)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+32)) = v422
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v419)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+24)) = v424
	v426 = *(*int64)(unsafe.Add(mBase, uint32(v419)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+16)) = v426
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v419)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v404)+8)) = v428
	v430 = *(*int64)(unsafe.Add(mBase, uint32(v419)))
	*(*int64)(unsafe.Add(mBase, uint32(v404))) = v430
	if v311 != v416 {
		v404 = v419
		v407 = v416
		goto L89
	} else {
		goto L91
	}
L90:
	;
	goto L88
L91:
	;
	goto L90
L92:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v453 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v451), base.F64_convert_i64_u(v453)), float64(0.1)) != 0 {
		goto L4
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v458 = v325 & v298
	v461 = v297 + v458*int32(48)
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461)+4)))
	if v462 != 0 {
		v311 = v458
		v314 = v448
		v315 = v461
		goto L67
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	goto L68
L97:
	;
	F_errmsg_internal(m, int32(_a_F_pagetable_insert_0), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L26
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_pagetable_insert_1), int32(635), int32(_a_F_pagetable_insert_2))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L26
	} else {
		goto L99
	}
L99:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L100:
	;
	F_errmsg_internal(m, int32(_a_F_pagetable_insert_3), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L26
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_pagetable_insert_1), int32(332), int32(_a_F_pagetable_insert_4))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L26
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pairingheap_SpGistSearchItem_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 float64
	_ = v42
	var v48 int64
	_ = v48
	var v53 float64
	_ = v53
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v108 int32
	_ = v108
	v11 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+42)))
	if v13 == v11 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return int32(0)
L2:
	;
	return v108
L3:
	;
	v108 = int32(-1)
	goto L2
L4:
	;
	v81 = int32(1)
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+43)))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+43)))
	if v83 == v81 {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	if v12&int32(1) != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if v12&int32(1) != 0 {
		v108 = v11
		goto L2
	} else {
		goto L9
	}
L8:
	;
	goto L3
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+112))
	if v20 <= int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v23 = int32(48)
	v31 = int32(0)
	goto L11
L11:
	;
	v39 = v31 << (uint(int32(3)) % 32)
	v40 = l1 + v23 + v39
	v42 = *(*float64)(unsafe.Add(mBase, uint32(l0+v23+v39)))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v42)&int64(9223372036854775807)) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L4
L13:
	;
	v69 = v31 + int32(1)
	if v69 != v20 {
		v31 = v69
		goto L11
	} else {
		goto L25
	}
L14:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	if base.Ui64(v48&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v53 = *(*float64)(unsafe.Add(mBase, uint32(v40)))
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v53)&int64(9223372036854775807)) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L13
L18:
	;
	return int32(1)
L19:
	;
	goto L20
L20:
	;
	if base.F64_eq(v42, v53) != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	if base.F64_lt(v42, v53) != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v65 = int32(1)
	goto L24
L23:
	;
	v65 = int32(-1)
	goto L24
L24:
	;
	return v65
L25:
	;
	goto L12
L26:
	;
	if v82&int32(1) == int32(0) {
		v108 = v81
		goto L2
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v82&int32(1) == int32(0) {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L1
L30:
	;
	goto L3
}
func F_palloc_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_palloc_extended[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+4)) = uint8(v3)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, v5, l0, l1)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if base.B2i32(l1&int32(4) == v3)|base.B2i32(v14 == int32(0)) != 0 {
			return v14
		} else {
			if l0&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(l0)) == int32(0) {
				if l0 == int32(0) {
					return v14
				} else {
					v32 = l0 + v14
					v34 = v14 + int32(4)
					if base.Ui32(v34) < base.Ui32(v32) {
						v36 = v32
					} else {
						v36 = v34
					}
					v41 = (v14^int32(-1)+v36)&int32(-4) + int32(4)
					if v41 == int32(0) {
						return v14
					} else {
						base.MemoryFill(m, v14, int32(0), v41)
						return v14
					}
				}
			} else {
				if l0 == int32(0) {
				} else {
					base.MemoryFill(m, v14, int32(0), l0)
				}
				return v14
			}
		}
	}
}
func F_parseXidFromText(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v9 = Fn14339(m, l0, l1, l2, int32(_a_F_parseXidFromText_0), int32(1391), int32(1386), int32(1381), int32(_a_F_parseXidFromText_1))
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		return v9
	}
}
func F_parse_format(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v414 int32
	_ = v414
	var v425 int32
	_ = v425
	var v445 int32
	_ = v445
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v599 int32
	_ = v599
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v667 int32
	_ = v667
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v690 int32
	_ = v690
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v734 int32
	_ = v734
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v804 int32
	_ = v804
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v972 int32
	_ = v972
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v19 == int32(0) {
		v720 = l0
		goto L15
	} else {
		goto L16
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v956 = m.ExcPending
	if v956 != 0 {
		goto L78
	} else {
		goto L265
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L78
	} else {
		goto L260
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L78
	} else {
		goto L256
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L78
	} else {
		goto L252
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L78
	} else {
		goto L248
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L78
	} else {
		goto L244
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L78
	} else {
		goto L240
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L78
	} else {
		goto L236
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L78
	} else {
		goto L232
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L78
	} else {
		goto L228
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L78
	} else {
		goto L224
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L78
	} else {
		goto L220
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L78
	} else {
		goto L216
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v744 = m.ExcPending
	if v744 != 0 {
		goto L78
	} else {
		goto L212
	}
L15:
	;
	v734 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v720)+9)) = uint8(v734)
	*(*int32)(unsafe.Add(mBase, uint32(v720))) = int32(1)
	m.G0 = v17 + int32(16)
	return
L16:
	;
	v25 = l5 & int32(1)
	v28 = l0
	v29 = l1
	v36 = v19
	goto L17
L17:
	;
	if v25 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v720 = v705
	goto L15
L19:
	;
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706))))
	if v719 != 0 {
		v28 = v705
		v29 = v706
		v36 = v719
		goto L17
	} else {
		goto L211
	}
L20:
	;
	if base.Ui32((v139-int32(126))&int32(255)) < base.Ui32(int32(163)) {
		goto L54
	} else {
		goto L55
	}
L21:
	;
	v132 = v29
	v139 = v36
	v140 = int32(0)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v45 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v132 = v29
	v139 = v36
	v140 = int32(0)
	goto L20
L25:
	;
	goto L26
L26:
	;
	v54 = l3
	v56 = v45
	goto L29
L27:
	;
	if v125&int32(255) == int32(0) {
		v705 = v28
		v706 = v122
		goto L19
	} else {
		goto L49
	}
L28:
	;
	v119 = v29 + v66
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v122 = v119
	v125 = v120
	v126 = v121
	goto L27
L29:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	if v63 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v122 = v29
	v125 = v36
	v126 = int32(0)
	goto L27
L31:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v66 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L33
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	if v115 != 0 {
		v54 = v54 + int32(16)
		v56 = v115
		goto L29
	} else {
		goto L48
	}
L34:
	;
	if v111 == int32(0) {
		goto L28
	} else {
		goto L47
	}
L35:
	;
	v111 = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v72 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v73 = v29
	v74 = v56
	v75 = v66
	v76 = v72
	goto L42
L39:
	;
	v99 = v56
	v103 = int32(0)
	goto L40
L40:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	v111 = v103 - v104
	goto L34
L41:
	;
	v99 = v94
	v103 = v96
	goto L40
L42:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if base.B2i32(v76 != v78)|base.B2i32(v78 == int32(0)) != 0 {
		v94 = v74
		v96 = v76
		goto L41
	} else {
		goto L44
	}
L43:
	;
	v94 = v88
	v96 = int32(0)
	goto L41
L44:
	;
	v84 = v75 - int32(1)
	if v84 == int32(0) {
		v94 = v74
		v96 = v76
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v87 = int32(1)
	v88 = v74 + v87
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v89 != 0 {
		v73 = v73 + v87
		v74 = v88
		v75 = v84
		v76 = v89
		goto L42
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	goto L33
L48:
	;
	goto L30
L49:
	;
	v132 = v122
	v139 = v125
	v140 = v126
	goto L20
L50:
	;
	v705 = v28 + int32(16)
	v706 = v690
	goto L19
L51:
	;
	if v250 != int32(92) {
		goto L179
	} else {
		goto L180
	}
L52:
	;
	if v250 == int32(32) {
		goto L175
	} else {
		goto L176
	}
L53:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)) = uint8(v140)
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v170
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if l5&int32(2) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L54:
	;
	v248 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v248
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v250 == v248 {
		v720 = v28
		goto L15
	} else {
		goto L75
	}
L55:
	;
	v152 = v139 & int32(255)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l4+v152<<(uint(int32(2))%32)-int32(128))))
	if v158 < int32(0) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v163 = l2 + v158*int32(20)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v170 = v163
	v172 = v164
	goto L57
L57:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	if v179 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L54
L59:
	;
	if v224 == int32(0) {
		goto L53
	} else {
		goto L72
	}
L60:
	;
	v224 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v185 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v186 = v132
	v187 = v172
	v188 = v179
	v189 = v185
	goto L67
L64:
	;
	v212 = v172
	v216 = int32(0)
	goto L65
L65:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212))))
	v224 = v216 - v217
	goto L59
L66:
	;
	v212 = v207
	v216 = v209
	goto L65
L67:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	if base.B2i32(v189 != v191)|base.B2i32(v191 == int32(0)) != 0 {
		v207 = v187
		v209 = v189
		goto L66
	} else {
		goto L69
	}
L68:
	;
	v207 = v201
	v209 = int32(0)
	goto L66
L69:
	;
	v197 = v188 - int32(1)
	if v197 == int32(0) {
		v207 = v187
		v209 = v189
		goto L66
	} else {
		goto L70
	}
L70:
	;
	v200 = int32(1)
	v201 = v187 + v200
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186)+1)))
	if v202 != 0 {
		v186 = v186 + v200
		v187 = v201
		v188 = v197
		v189 = v202
		goto L67
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v170)+20))
	if v227 == int32(0) {
		goto L54
	} else {
		goto L73
	}
L73:
	;
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227))))
	if v152 == v232 {
		v170 = v170 + int32(20)
		v172 = v227
		goto L57
	} else {
		goto L74
	}
L74:
	;
	goto L58
L75:
	;
	if base.B2i32(base.Ui32(l5) < base.Ui32(int32(4)))|base.B2i32(v250 == int32(34)) != 0 {
		goto L51
	} else {
		goto L76
	}
L76:
	;
	if base.B2i32(base.Ui32(v250) <= base.Ui32(int32(63)))&base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(v250))%64)&int64(864955565296582657) != int64(0)) != 0 {
		goto L52
	} else {
		goto L77
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	return
L79:
	;
	F_errcode(m, int32(117440642))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v273 = F_pg_mblen_cstr(m, v132)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	v275 = F_pnstrdup(m, v132, v273)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v275
	F_errmsg(m, int32(_a_F_parse_format_0), v17)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L78
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1443), int32(_a_F_parse_format_2))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v480 = v132 + v290
	if v25 == int32(0) {
		v690 = v480
		goto L50
	} else {
		goto L151
	}
L86:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v170)+8))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	if v294&int32(_a_F_parse_format_3) != 0 {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	if v469&int32(1024) == int32(0) {
		goto L85
	} else {
		goto L149
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v466
	v469 = v466
	goto L87
L89:
	;
	if v294&int32(4080) != 0 {
		goto L2
	} else {
		goto L148
	}
L90:
	;
	F_errmsg(m, int32(_a_F_parse_format_4), int32(0))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L78
	} else {
		goto L146
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L78
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	switch v293 - int32(1) {
	case 0:
		v379 = v294
		goto L108
	case 1:
		goto L111
	case 2:
		goto L112
	case 3:
		goto L110
	default:
		v469 = v294
		goto L87
	case 5:
		goto L109
	case 6:
		goto L89
	case 7:
		goto L107
	case 8, 9:
		goto L100
	case 10:
		goto L105
	case 11:
		goto L104
	case 12:
		goto L102
	case 13, 29:
		goto L101
	case 14:
		goto L103
	case 16:
		goto L106
	case 18:
		goto L99
	}
L94:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L78
	} else {
		goto L95
	}
L95:
	;
	if v293 == int32(7) {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(_a_F_parse_format_5), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L78
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1195), int32(_a_F_parse_format_6))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L78
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L99:
	;
	if v294&int32(2) != 0 {
		goto L3
	} else {
		goto L145
	}
L100:
	;
	v445 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6)+32)) = uint8(v445)
	v469 = v294
	goto L87
L101:
	;
	if v294&int32(1024) != 0 {
		goto L4
	} else {
		goto L144
	}
L102:
	;
	if v294&int32(832) != 0 {
		goto L5
	} else {
		goto L143
	}
L103:
	;
	if v294&int32(64) != 0 {
		goto L6
	} else {
		goto L142
	}
L104:
	;
	if v294&int32(64) != 0 {
		goto L7
	} else {
		goto L140
	}
L105:
	;
	if v294&int32(64) != 0 {
		goto L8
	} else {
		goto L138
	}
L106:
	;
	if v294&int32(64) != 0 {
		goto L10
	} else {
		goto L132
	}
L107:
	;
	v466 = v294 | int32(32)
	goto L88
L108:
	;
	if v379&int32(2) != 0 {
		goto L12
	} else {
		goto L130
	}
L109:
	;
	v374 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6)+32)) = uint8(v374)
	v377 = v294 | int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v377
	v379 = v377
	goto L108
L110:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v367 != 0 {
		v469 = v294
		goto L87
	} else {
		goto L128
	}
L111:
	;
	if v294&int32(128) != 0 {
		goto L13
	} else {
		goto L120
	}
L112:
	;
	if v294&int32(128) != 0 {
		goto L14
	} else {
		goto L113
	}
L113:
	;
	if v294&int32(2048) != 0 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l6)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+20)) = v321 + int32(1)
	v469 = v294
	goto L87
L115:
	;
	goto L116
L116:
	;
	if v294&int32(2) != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v327 + int32(1)
	v469 = v294
	goto L87
L118:
	;
	goto L119
L119:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v331 + int32(1)
	v469 = v294
	goto L87
L120:
	;
	if v294&int32(10) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v342 = v294 | int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v342
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+24)) = v344 + int32(1)
	v348 = v342
	goto L123
L122:
	;
	v348 = v294
	goto L123
L123:
	;
	if v348&int32(2) == int32(0) {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+28)) = v363 + v364
	v469 = v348
	goto L87
L125:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v355 = v353 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v363 = v355
	v364 = v357
	goto L124
L126:
	;
	goto L127
L127:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	v360 = v358 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+4)) = v360
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v363 = v362
	v364 = v360
	goto L124
L128:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v368|v294&int32(8) != 0 {
		v469 = v294
		goto L87
	} else {
		goto L129
	}
L129:
	;
	v466 = v294 | int32(16)
	goto L88
L130:
	;
	if v379&int32(2048) != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	v466 = v379 | int32(2)
	goto L88
L132:
	;
	if v294&int32(896) != 0 {
		goto L9
	} else {
		goto L133
	}
L133:
	;
	if v294&int32(2) == int32(0) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = int32(-1)
	v398 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6)+32)) = uint8(v398)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	*(*int32)(unsafe.Add(mBase, uint32(l6)+16)) = v400
	v466 = v294 | int32(64)
	goto L88
L135:
	;
	goto L136
L136:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l6)+8))
	if v404 != 0 {
		v469 = v294
		goto L87
	} else {
		goto L137
	}
L137:
	;
	v405 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l6)+32)) = uint8(v405)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v405
	v466 = v294 | int32(64)
	goto L88
L138:
	;
	v414 = v294 | int32(256)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v414
	if v294&int32(2) == int32(0) {
		v469 = v414
		goto L87
	} else {
		goto L139
	}
L139:
	;
	v466 = v294 | int32(_a_F_parse_format_7)
	goto L88
L140:
	;
	v425 = v294 | int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v425
	if v294&int32(2) == int32(0) {
		v469 = v425
		goto L87
	} else {
		goto L141
	}
L141:
	;
	v466 = v294 | int32(_a_F_parse_format_8)
	goto L88
L142:
	;
	v466 = v294 | int32(768)
	goto L88
L143:
	;
	v466 = v294 | int32(128)
	goto L88
L144:
	;
	v466 = v294 | int32(1024)
	goto L88
L145:
	;
	v466 = v294 | int32(2048)
	goto L88
L146:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1345), int32(_a_F_parse_format_6))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L78
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	v466 = v294 | int32(_a_F_parse_format_3)
	goto L88
L149:
	;
	if v469&int32(-1057) != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	goto L85
L151:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v483 == int32(0) {
		v690 = v480
		goto L50
	} else {
		goto L152
	}
L152:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v486 == int32(0) {
		v690 = v480
		goto L50
	} else {
		goto L153
	}
L153:
	;
	v494 = l3
	v496 = v486
	goto L154
L154:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if v503 == int32(2) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)))
	v558 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+8)))
	v559 = v557 | v558
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)) = uint8(v559)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v690 = v480 + v561
	goto L50
L156:
	;
	goto L155
L157:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	if v506 == int32(0) {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	goto L159
L159:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v494)+16))
	if v554 != 0 {
		v494 = v494 + int32(16)
		v496 = v554
		goto L154
	} else {
		goto L174
	}
L160:
	;
	if v551 == int32(0) {
		goto L156
	} else {
		goto L173
	}
L161:
	;
	v551 = int32(0)
	goto L160
L162:
	;
	goto L163
L163:
	;
	v512 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	if v512 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v513 = v480
	v514 = v496
	v515 = v506
	v516 = v512
	goto L168
L165:
	;
	v539 = v496
	v543 = int32(0)
	goto L166
L166:
	;
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v539))))
	v551 = v543 - v544
	goto L160
L167:
	;
	v539 = v534
	v543 = v536
	goto L166
L168:
	;
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	if base.B2i32(v516 != v518)|base.B2i32(v518 == int32(0)) != 0 {
		v534 = v514
		v536 = v516
		goto L167
	} else {
		goto L170
	}
L169:
	;
	v534 = v528
	v536 = int32(0)
	goto L167
L170:
	;
	v524 = v515 - int32(1)
	if v524 == int32(0) {
		v534 = v514
		v536 = v516
		goto L167
	} else {
		goto L171
	}
L171:
	;
	v527 = int32(1)
	v528 = v514 + v527
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v513)+1)))
	if v529 != 0 {
		v513 = v513 + v527
		v514 = v528
		v515 = v524
		v516 = v529
		goto L168
	} else {
		goto L172
	}
L172:
	;
	goto L169
L173:
	;
	goto L159
L174:
	;
	v690 = v480
	goto L50
L175:
	;
	v567 = int32(5)
	goto L177
L176:
	;
	v567 = int32(4)
	goto L177
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v567
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	v570 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v570
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+5)) = uint8(v570)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+4)) = uint8(v569)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)) = uint8(v570)
	v690 = v132 + int32(1)
	goto L50
L178:
	;
	v637 = F_pg_mblen_cstr(m, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L78
	} else {
		goto L201
	}
L179:
	;
	if v250 != int32(34) {
		v636 = v132
		goto L178
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v632 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v632 == int32(34) {
		goto L198
	} else {
		goto L199
	}
L182:
	;
	v585 = v28
	v590 = v132 + int32(1)
	goto L183
L183:
	;
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if v599 != int32(92) {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v613 = F_pg_mblen_cstr(m, v612)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L78
	} else {
		goto L194
	}
L186:
	;
	if v599 == int32(0) {
		v720 = v585
		goto L15
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+1)))
	if v610 != 0 {
		goto L191
	} else {
		goto L192
	}
L189:
	;
	if v599 != int32(34) {
		v612 = v590
		goto L185
	} else {
		goto L190
	}
L190:
	;
	v705 = v585
	v706 = v590 + int32(1)
	goto L19
L191:
	;
	v611 = v590 + int32(1)
	goto L193
L192:
	;
	v611 = v590
	goto L193
L193:
	;
	v612 = v611
	goto L185
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v585))) = int32(3)
	v618 = v585 + int32(4)
	if v613 != 0 {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	base.MemoryCopy(m, v618, v612, v613)
	goto L197
L196:
	;
	goto L197
L197:
	;
	v621 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v613+v618))) = uint8(v621)
	*(*uint8)(unsafe.Add(mBase, uint32(v585)+9)) = uint8(v621)
	*(*int32)(unsafe.Add(mBase, uint32(v585)+12)) = v621
	v585 = v585 + int32(16)
	v590 = v613 + v612
	goto L183
L198:
	;
	v635 = v132 + int32(1)
	goto L200
L199:
	;
	v635 = v132
	goto L200
L200:
	;
	v636 = v635
	goto L178
L201:
	;
	v639 = int32(0)
	v641 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v636))))
	v644 = int32(255)
	if base.B2i32(v25 == v639)|base.B2i32(base.Ui32(int32(93)) < base.Ui32((v641-int32(33))&v644))|base.B2i32(base.Ui32(int32(229)) < base.Ui32((v641&int32(223)-int32(91))&v644)) == v639 {
		goto L203
	} else {
		goto L204
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v676
	v679 = v28 + int32(4)
	if v637 != 0 {
		goto L208
	} else {
		goto L209
	}
L203:
	;
	if base.Ui32((v641-int32(58))&int32(255)) < base.Ui32(int32(246)) {
		v676 = int32(4)
		goto L202
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	v667 = int32(5)
	if base.B2i32(v641 == int32(32))|base.B2i32(base.Ui32(v641-int32(9)) < base.Ui32(v667)) != 0 {
		v676 = v667
		goto L202
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	v676 = int32(3)
	goto L202
L208:
	;
	base.MemoryCopy(m, v679, v636, v637)
	goto L210
L209:
	;
	goto L210
L210:
	;
	v682 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v637+v679))) = uint8(v682)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+9)) = uint8(v682)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v682
	v690 = v636 + v637
	goto L50
L211:
	;
	goto L18
L212:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L78
	} else {
		goto L213
	}
L213:
	;
	F_errmsg(m, int32(_a_F_parse_format_9), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L78
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1203), int32(_a_F_parse_format_6))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L78
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L78
	} else {
		goto L217
	}
L217:
	;
	F_errmsg(m, int32(_a_F_parse_format_10), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L78
	} else {
		goto L218
	}
L218:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1219), int32(_a_F_parse_format_6))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L78
	} else {
		goto L219
	}
L219:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L220:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L78
	} else {
		goto L221
	}
L221:
	;
	F_errmsg(m, int32(_a_F_parse_format_11), int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L78
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1246), int32(_a_F_parse_format_6))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L78
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L78
	} else {
		goto L225
	}
L225:
	;
	F_errmsg(m, int32(_a_F_parse_format_12), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L78
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1250), int32(_a_F_parse_format_6))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L78
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L78
	} else {
		goto L229
	}
L229:
	;
	F_errmsg(m, int32(_a_F_parse_format_13), int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L78
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1262), int32(_a_F_parse_format_6))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L78
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L78
	} else {
		goto L233
	}
L233:
	;
	F_errmsg(m, int32(_a_F_parse_format_14), int32(0))
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L78
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1266), int32(_a_F_parse_format_6))
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L78
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L78
	} else {
		goto L237
	}
L237:
	;
	F_errmsg(m, int32(_a_F_parse_format_15), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L78
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1286), int32(_a_F_parse_format_6))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L78
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L78
	} else {
		goto L241
	}
L241:
	;
	F_errmsg(m, int32(_a_F_parse_format_16), int32(0))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L78
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1296), int32(_a_F_parse_format_6))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L78
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L78
	} else {
		goto L245
	}
L245:
	;
	F_errmsg(m, int32(_a_F_parse_format_17), int32(0))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L78
	} else {
		goto L246
	}
L246:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1306), int32(_a_F_parse_format_6))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L78
	} else {
		goto L247
	}
L247:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L248:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L78
	} else {
		goto L249
	}
L249:
	;
	F_errmsg(m, int32(_a_F_parse_format_18), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L78
	} else {
		goto L250
	}
L250:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1315), int32(_a_F_parse_format_6))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L78
	} else {
		goto L251
	}
L251:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L252:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v907 = m.ExcPending
	if v907 != 0 {
		goto L78
	} else {
		goto L253
	}
L253:
	;
	F_errmsg(m, int32(_a_F_parse_format_19), int32(0))
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L78
	} else {
		goto L254
	}
L254:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1324), int32(_a_F_parse_format_6))
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L78
	} else {
		goto L255
	}
L255:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L256:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L78
	} else {
		goto L257
	}
L257:
	;
	F_errmsg(m, int32(_a_F_parse_format_12), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L78
	} else {
		goto L258
	}
L258:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1337), int32(_a_F_parse_format_6))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L78
	} else {
		goto L259
	}
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L260:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L78
	} else {
		goto L261
	}
L261:
	;
	F_errmsg(m, int32(_a_F_parse_format_20), int32(0))
	mBase = m.M
	v943 = m.ExcPending
	if v943 != 0 {
		goto L78
	} else {
		goto L262
	}
L262:
	;
	v946 = F_errdetail(m, int32(_a_F_parse_format_21), int32(0))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L78
	} else {
		goto L263
	}
L263:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1352), int32(_a_F_parse_format_6))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L78
	} else {
		goto L264
	}
L264:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L265:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L78
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(_a_F_parse_format_22), int32(0))
	mBase = m.M
	v963 = m.ExcPending
	if v963 != 0 {
		goto L78
	} else {
		goto L267
	}
L267:
	;
	v966 = F_errdetail(m, int32(_a_F_parse_format_23), int32(0))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L78
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_parse_format_1), int32(1362), int32(_a_F_parse_format_6))
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L78
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_parse_scram_secret(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v301 int32
	_ = v301
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v399 int32
	_ = v399
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v500 int32
	_ = v500
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v697 int32
	_ = v697
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v746 int32
	_ = v746
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = F_pstrdup(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v17
	v23 = v15 + int32(12)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v38 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	v27 = F_strcspn(m, v26, int32(_a_F_parse_scram_secret_0))
	mBase = m.M
	v28 = v27 + v26
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	goto L3
L7:
	;
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v30)
	v35 = v28 + int32(1)
	goto L9
L8:
	;
	v35 = int32(0)
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v35
	goto L6
L10:
	;
	m.G0 = v15 + int32(16)
	return v746
L11:
	;
	v737 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v737
	v746 = v737
	goto L10
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v43 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v55 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v44 = F_strcspn(m, v43, int32(_a_F_parse_scram_secret_1))
	mBase = m.M
	v45 = v44 + v43
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v46 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	v47 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v45))) = uint8(v47)
	v52 = v45 + int32(1)
	goto L19
L18:
	;
	v52 = int32(0)
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v52
	goto L16
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v60 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v72 == int32(0) {
		goto L11
	} else {
		goto L28
	}
L22:
	;
	v61 = F_strcspn(m, v60, int32(_a_F_parse_scram_secret_0))
	mBase = m.M
	v62 = v61 + v60
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	if v63 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v62))) = uint8(v64)
	v69 = v62 + int32(1)
	goto L27
L26:
	;
	v69 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v69
	goto L24
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v77 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	if v89 == int32(0) {
		goto L11
	} else {
		goto L36
	}
L30:
	;
	v78 = F_strcspn(m, v77, int32(_a_F_parse_scram_secret_1))
	mBase = m.M
	v79 = v78 + v77
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v80 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	v81 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v81)
	v86 = v79 + int32(1)
	goto L35
L34:
	;
	v86 = int32(0)
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v86
	goto L32
L36:
	;
	v92 = int32(_a_F_parse_scram_secret_2)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_parse_scram_secret[0])))
	if base.B2i32(v95 == int32(0))|base.B2i32(v95 != v98) != 0 {
		v116 = v95
		v117 = v98
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v116-v117 != 0 {
		goto L11
	} else {
		goto L44
	}
L38:
	;
	goto L37
L39:
	;
	v101 = v26
	v102 = v92
	goto L40
L40:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+1)))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	if v106 == int32(0) {
		v116 = v106
		v117 = v105
		goto L38
	} else {
		goto L42
	}
L41:
	;
	v116 = v106
	v117 = v105
	goto L38
L42:
	;
	v109 = int32(1)
	if v106 == v105 {
		v101 = v101 + v109
		v102 = v102 + v109
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(32)
	*(*int32)(unsafe.Add(mBase, _c_F_parse_scram_secret[1])) = int32(0)
	v130 = F_strtox_2(m, v43, v15+int32(8), int32(10), int64(2147483648))
	mBase = m.M
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_wrap_i64(v130)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	if v134 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_parse_scram_secret[1]))
	if v136 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v137 = F_strlen(m, v60)
	mBase = m.M
	v141 = v137 * int32(3) >> (uint(int32(2)) % 32)
	goto L48
L48:
	;
	v142 = F_palloc(m, v141)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v144 = F_strlen(m, v60)
	mBase = m.M
	v145 = int32(0)
	if v145 < v144 {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	if v330 < int32(0) {
		goto L11
	} else {
		goto L96
	}
L51:
	;
	if v141 != 0 {
		goto L93
	} else {
		goto L94
	}
L52:
	;
	v154 = v60 + v144
	v155 = v60
	v159 = v145
	v160 = v142
	v162 = v145
	v163 = v145
	goto L55
L53:
	;
	v301 = v142
	goto L54
L54:
	;
	v330 = v301 - v142
	goto L50
L55:
	;
	v167 = v155 + int32(1)
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	if v168 != int32(61) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	if v289 != 0 {
		goto L51
	} else {
		goto L92
	}
L57:
	;
	if base.Ui32(v287) < base.Ui32(v154) {
		v155 = v287
		v159 = v289
		v160 = v290
		v162 = v292
		v163 = v293
		goto L55
	} else {
		goto L91
	}
L58:
	;
	if v141 < v160-v142+int32(1) {
		goto L51
	} else {
		goto L78
	}
L59:
	;
	v245 = v167
	v246 = int32(2)
	v249 = v163 << (uint(int32(6)) % 32)
	goto L58
L60:
	;
	v234 = v231 + v230<<(uint(int32(6))%32)
	v236 = v227 + int32(1)
	if v236 == int32(4) {
		v245 = v228
		v246 = v229
		v249 = v234
		goto L58
	} else {
		goto L77
	}
L61:
	;
	v227 = int32(3)
	v228 = v155 + int32(2)
	v229 = int32(1)
	v230 = v187
	v231 = v182
	goto L60
L62:
	;
	if base.Ui32(int32(125)) < base.Ui32((v206-int32(1))&int32(255)) {
		goto L51
	} else {
		goto L75
	}
L63:
	;
	v172 = v168 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v172))|base.B2i32(int32(1)<<(uint(v172)%32)&int32(_a_F_parse_scram_secret_3) == int32(0)) != 0 {
		v206 = v168
		v207 = v159
		v208 = v167
		v209 = v162
		v210 = v163
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v182 = int32(0)
	if v162 != 0 {
		v227 = v159
		v228 = v167
		v229 = v162
		v230 = v163
		v231 = v182
		goto L60
	} else {
		goto L67
	}
L66:
	;
	goto L51
L67:
	;
	switch v159 - int32(2) {
	case 0:
		goto L68
	case 1:
		goto L59
	default:
		goto L51
	}
L68:
	;
	if base.Ui32(v154) <= base.Ui32(v167) {
		goto L51
	} else {
		goto L69
	}
L69:
	;
	v187 = v163 << (uint(int32(6)) % 32)
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v188 == int32(61) {
		goto L61
	} else {
		goto L70
	}
L70:
	;
	v192 = v188 - int32(9)
	if int32(1)<<(uint(v192)%32)&int32(_a_F_parse_scram_secret_3) != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v200 = base.B2i32(base.Ui32(v192) <= base.Ui32(int32(23)))
	goto L73
L72:
	;
	v200 = int32(0)
	goto L73
L73:
	;
	if v200 != 0 {
		goto L51
	} else {
		goto L74
	}
L74:
	;
	v206 = v188
	v207 = int32(3)
	v208 = v155 + int32(2)
	v209 = int32(1)
	v210 = v187
	goto L62
L75:
	;
	v218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v206)+uint32(_c_F_parse_scram_secret[2]))))
	if v218 < int32(0) {
		goto L51
	} else {
		goto L76
	}
L76:
	;
	v227 = v207
	v228 = v208
	v229 = v209
	v230 = v210
	v231 = v218
	goto L60
L77:
	;
	v287 = v228
	v289 = v236
	v290 = v160
	v292 = v229
	v293 = v234
	goto L57
L78:
	;
	v255 = int32(base.Ui32(v249) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v255)
	v258 = v160 + int32(1)
	if base.Ui32(v246) < base.Ui32(int32(2)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v262 = v246
	goto L81
L80:
	;
	v262 = int32(0)
	goto L81
L81:
	;
	if v262 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v141 < v258-v142+int32(1) {
		goto L51
	} else {
		goto L85
	}
L83:
	;
	v274 = v258
	goto L84
L84:
	;
	if v246 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	v270 = int32(base.Ui32(v249) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v160)+1)) = uint8(v270)
	v274 = v160 + int32(2)
	goto L84
L86:
	;
	v287 = v245
	v289 = int32(0)
	v290 = v284
	v292 = v285
	v293 = int32(0)
	goto L57
L87:
	;
	v284 = v274
	v285 = v246
	goto L86
L88:
	;
	goto L89
L89:
	;
	if v141 < v274-v142+int32(1) {
		goto L51
	} else {
		goto L90
	}
L90:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v249)
	v284 = v274 + int32(1)
	v285 = int32(0)
	goto L86
L91:
	;
	goto L56
L92:
	;
	v301 = v290
	goto L54
L93:
	;
	base.MemoryFill(m, v142, int32(0), v141)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v330 = int32(-1)
	goto L50
L96:
	;
	v333 = F_pstrdup(m, v60)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v333
	v336 = F_strlen(m, v77)
	mBase = m.M
	v340 = v336 * int32(3) >> (uint(int32(2)) % 32)
	goto L98
L98:
	;
	v341 = F_palloc(m, v340)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v343 = F_strlen(m, v77)
	mBase = m.M
	v344 = int32(0)
	if v344 < v343 {
		goto L102
	} else {
		goto L103
	}
L100:
	;
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v529 != v530 {
		goto L11
	} else {
		goto L146
	}
L101:
	;
	if v340 != 0 {
		goto L143
	} else {
		goto L144
	}
L102:
	;
	v353 = v77 + v343
	v354 = v77
	v358 = v344
	v359 = v341
	v361 = v344
	v362 = v344
	goto L105
L103:
	;
	v500 = v341
	goto L104
L104:
	;
	v529 = v500 - v341
	goto L100
L105:
	;
	v366 = v354 + int32(1)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v354))))
	if v367 != int32(61) {
		goto L113
	} else {
		goto L114
	}
L106:
	;
	if v488 != 0 {
		goto L101
	} else {
		goto L142
	}
L107:
	;
	if base.Ui32(v486) < base.Ui32(v353) {
		v354 = v486
		v358 = v488
		v359 = v489
		v361 = v491
		v362 = v492
		goto L105
	} else {
		goto L141
	}
L108:
	;
	if v340 < v359-v341+int32(1) {
		goto L101
	} else {
		goto L128
	}
L109:
	;
	v444 = v366
	v445 = int32(2)
	v448 = v362 << (uint(int32(6)) % 32)
	goto L108
L110:
	;
	v433 = v430 + v429<<(uint(int32(6))%32)
	v435 = v426 + int32(1)
	if v435 == int32(4) {
		v444 = v427
		v445 = v428
		v448 = v433
		goto L108
	} else {
		goto L127
	}
L111:
	;
	v426 = int32(3)
	v427 = v354 + int32(2)
	v428 = int32(1)
	v429 = v386
	v430 = v381
	goto L110
L112:
	;
	if base.Ui32(int32(125)) < base.Ui32((v405-int32(1))&int32(255)) {
		goto L101
	} else {
		goto L125
	}
L113:
	;
	v371 = v367 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v371))|base.B2i32(int32(1)<<(uint(v371)%32)&int32(_a_F_parse_scram_secret_3) == int32(0)) != 0 {
		v405 = v367
		v406 = v358
		v407 = v366
		v408 = v361
		v409 = v362
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v381 = int32(0)
	if v361 != 0 {
		v426 = v358
		v427 = v366
		v428 = v361
		v429 = v362
		v430 = v381
		goto L110
	} else {
		goto L117
	}
L116:
	;
	goto L101
L117:
	;
	switch v358 - int32(2) {
	case 0:
		goto L118
	case 1:
		goto L109
	default:
		goto L101
	}
L118:
	;
	if base.Ui32(v353) <= base.Ui32(v366) {
		goto L101
	} else {
		goto L119
	}
L119:
	;
	v386 = v362 << (uint(int32(6)) % 32)
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366))))
	if v387 == int32(61) {
		goto L111
	} else {
		goto L120
	}
L120:
	;
	v391 = v387 - int32(9)
	if int32(1)<<(uint(v391)%32)&int32(_a_F_parse_scram_secret_3) != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v399 = base.B2i32(base.Ui32(v391) <= base.Ui32(int32(23)))
	goto L123
L122:
	;
	v399 = int32(0)
	goto L123
L123:
	;
	if v399 != 0 {
		goto L101
	} else {
		goto L124
	}
L124:
	;
	v405 = v387
	v406 = int32(3)
	v407 = v354 + int32(2)
	v408 = int32(1)
	v409 = v386
	goto L112
L125:
	;
	v417 = int32(*(*int8)(unsafe.Add(mBase, uint32(v405)+uint32(_c_F_parse_scram_secret[2]))))
	if v417 < int32(0) {
		goto L101
	} else {
		goto L126
	}
L126:
	;
	v426 = v406
	v427 = v407
	v428 = v408
	v429 = v409
	v430 = v417
	goto L110
L127:
	;
	v486 = v427
	v488 = v435
	v489 = v359
	v491 = v428
	v492 = v433
	goto L107
L128:
	;
	v454 = int32(base.Ui32(v448) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v359))) = uint8(v454)
	v457 = v359 + int32(1)
	if base.Ui32(v445) < base.Ui32(int32(2)) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v461 = v445
	goto L131
L130:
	;
	v461 = int32(0)
	goto L131
L131:
	;
	if v461 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if v340 < v457-v341+int32(1) {
		goto L101
	} else {
		goto L135
	}
L133:
	;
	v473 = v457
	goto L134
L134:
	;
	if v445 != 0 {
		goto L137
	} else {
		goto L138
	}
L135:
	;
	v469 = int32(base.Ui32(v448) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v359)+1)) = uint8(v469)
	v473 = v359 + int32(2)
	goto L134
L136:
	;
	v486 = v444
	v488 = int32(0)
	v489 = v483
	v491 = v484
	v492 = int32(0)
	goto L107
L137:
	;
	v483 = v473
	v484 = v445
	goto L136
L138:
	;
	goto L139
L139:
	;
	if v340 < v473-v341+int32(1) {
		goto L101
	} else {
		goto L140
	}
L140:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v473))) = uint8(v448)
	v483 = v473 + int32(1)
	v484 = int32(0)
	goto L136
L141:
	;
	goto L106
L142:
	;
	v500 = v489
	goto L104
L143:
	;
	base.MemoryFill(m, v341, int32(0), v340)
	goto L145
L144:
	;
	goto L145
L145:
	;
	v529 = int32(-1)
	goto L100
L146:
	;
	if v529 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	base.MemoryCopy(m, l5, v341, v529)
	goto L149
L148:
	;
	goto L149
L149:
	;
	v533 = F_strlen(m, v89)
	mBase = m.M
	v537 = v533 * int32(3) >> (uint(int32(2)) % 32)
	goto L150
L150:
	;
	v538 = F_palloc(m, v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v540 = F_strlen(m, v89)
	mBase = m.M
	v541 = int32(0)
	if v541 < v540 {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v726 != v727 {
		goto L11
	} else {
		goto L198
	}
L153:
	;
	if v537 != 0 {
		goto L195
	} else {
		goto L196
	}
L154:
	;
	v550 = v89 + v540
	v551 = v89
	v555 = v541
	v556 = v538
	v558 = v541
	v559 = v541
	goto L157
L155:
	;
	v697 = v538
	goto L156
L156:
	;
	v726 = v697 - v538
	goto L152
L157:
	;
	v563 = v551 + int32(1)
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	if v564 != int32(61) {
		goto L165
	} else {
		goto L166
	}
L158:
	;
	if v685 != 0 {
		goto L153
	} else {
		goto L194
	}
L159:
	;
	if base.Ui32(v683) < base.Ui32(v550) {
		v551 = v683
		v555 = v685
		v556 = v686
		v558 = v688
		v559 = v689
		goto L157
	} else {
		goto L193
	}
L160:
	;
	if v537 < v556-v538+int32(1) {
		goto L153
	} else {
		goto L180
	}
L161:
	;
	v641 = v563
	v642 = int32(2)
	v645 = v559 << (uint(int32(6)) % 32)
	goto L160
L162:
	;
	v630 = v627 + v626<<(uint(int32(6))%32)
	v632 = v623 + int32(1)
	if v632 == int32(4) {
		v641 = v624
		v642 = v625
		v645 = v630
		goto L160
	} else {
		goto L179
	}
L163:
	;
	v623 = int32(3)
	v624 = v551 + int32(2)
	v625 = int32(1)
	v626 = v583
	v627 = v578
	goto L162
L164:
	;
	if base.Ui32(int32(125)) < base.Ui32((v602-int32(1))&int32(255)) {
		goto L153
	} else {
		goto L177
	}
L165:
	;
	v568 = v564 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v568))|base.B2i32(int32(1)<<(uint(v568)%32)&int32(_a_F_parse_scram_secret_3) == int32(0)) != 0 {
		v602 = v564
		v603 = v555
		v604 = v563
		v605 = v558
		v606 = v559
		goto L164
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v578 = int32(0)
	if v558 != 0 {
		v623 = v555
		v624 = v563
		v625 = v558
		v626 = v559
		v627 = v578
		goto L162
	} else {
		goto L169
	}
L168:
	;
	goto L153
L169:
	;
	switch v555 - int32(2) {
	case 0:
		goto L170
	case 1:
		goto L161
	default:
		goto L153
	}
L170:
	;
	if base.Ui32(v550) <= base.Ui32(v563) {
		goto L153
	} else {
		goto L171
	}
L171:
	;
	v583 = v559 << (uint(int32(6)) % 32)
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v563))))
	if v584 == int32(61) {
		goto L163
	} else {
		goto L172
	}
L172:
	;
	v588 = v584 - int32(9)
	if int32(1)<<(uint(v588)%32)&int32(_a_F_parse_scram_secret_3) != 0 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v596 = base.B2i32(base.Ui32(v588) <= base.Ui32(int32(23)))
	goto L175
L174:
	;
	v596 = int32(0)
	goto L175
L175:
	;
	if v596 != 0 {
		goto L153
	} else {
		goto L176
	}
L176:
	;
	v602 = v584
	v603 = int32(3)
	v604 = v551 + int32(2)
	v605 = int32(1)
	v606 = v583
	goto L164
L177:
	;
	v614 = int32(*(*int8)(unsafe.Add(mBase, uint32(v602)+uint32(_c_F_parse_scram_secret[2]))))
	if v614 < int32(0) {
		goto L153
	} else {
		goto L178
	}
L178:
	;
	v623 = v603
	v624 = v604
	v625 = v605
	v626 = v606
	v627 = v614
	goto L162
L179:
	;
	v683 = v624
	v685 = v632
	v686 = v556
	v688 = v625
	v689 = v630
	goto L159
L180:
	;
	v651 = int32(base.Ui32(v645) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v556))) = uint8(v651)
	v654 = v556 + int32(1)
	if base.Ui32(v642) < base.Ui32(int32(2)) {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v658 = v642
	goto L183
L182:
	;
	v658 = int32(0)
	goto L183
L183:
	;
	if v658 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	if v537 < v654-v538+int32(1) {
		goto L153
	} else {
		goto L187
	}
L185:
	;
	v670 = v654
	goto L186
L186:
	;
	if v642 != 0 {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v666 = int32(base.Ui32(v645) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)) = uint8(v666)
	v670 = v556 + int32(2)
	goto L186
L188:
	;
	v683 = v641
	v685 = int32(0)
	v686 = v680
	v688 = v681
	v689 = int32(0)
	goto L159
L189:
	;
	v680 = v670
	v681 = v642
	goto L188
L190:
	;
	goto L191
L191:
	;
	if v537 < v670-v538+int32(1) {
		goto L153
	} else {
		goto L192
	}
L192:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v670))) = uint8(v645)
	v680 = v670 + int32(1)
	v681 = int32(0)
	goto L188
L193:
	;
	goto L158
L194:
	;
	v697 = v686
	goto L156
L195:
	;
	base.MemoryFill(m, v538, int32(0), v537)
	goto L197
L196:
	;
	goto L197
L197:
	;
	v726 = int32(-1)
	goto L152
L198:
	;
	if v726 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	base.MemoryCopy(m, l6, v538, v726)
	goto L201
L200:
	;
	goto L201
L201:
	;
	v746 = int32(1)
	goto L10
}
func F_parser_coercion_errposition(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	if l1 < int32(0) {
		v6 = F_exprLocation(m, l2)
		v7 = v6
	} else {
		v7 = l1
	}
	F_parser_errposition(m, l0, v7)
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_pathkeys_contained_in(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	if l0 == l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	goto L3
L3:
	;
	v12 = int32(0)
	goto L5
L4:
	;
	return v49
L5:
	;
	v16 = int32(0)
	if l0 == v16 {
		v26 = v16
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v49 = int32(0)
	goto L4
L7:
	;
	if l1 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 <= v12 {
		v26 = int32(0)
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v26 = v22 + v12<<(uint(int32(2))%32)
	goto L7
L10:
	;
	v33 = base.B2i32(v26 == int32(0))
	if v26 == int32(0) {
		v49 = v33
		goto L4
	} else {
		goto L15
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 < v27 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	return base.B2i32(v26 == int32(0))
L14:
	;
	goto L13
L15:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v36 == int32(0) {
		v49 = v33
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12<<(uint(int32(2))%32)+v36)))
	if v43 == v45 {
		v12 = v12 + int32(1)
		goto L5
	} else {
		goto L17
	}
L17:
	;
	goto L6
}
func F_pct_info_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v5 != v6 {
		if v5 < v6 {
			v11 = int32(-1)
		} else {
			v11 = int32(1)
		}
		return v11
	} else {
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
		if v13 != v14 {
			if v13 < v14 {
				v19 = int32(-1)
			} else {
				v19 = int32(1)
			}
			v21 = v19
		} else {
			v21 = int32(0)
		}
		return v21
	}
}
func F_pgl_system(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pgl_system[0]))
	if v4 == int32(0) {
		return int32(123)
	} else {
		v9 = m.T0[v4].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_pglz_decompress(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v198 int32
	_ = v198
	v6 = int32(0)
	v14 = l2 + l3
	v15 = l0 + l1
	if base.B2i32(l1 <= v6)|base.B2i32(l3 <= v6) == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v198
L2:
	;
	if l4 != 0 {
		goto L43
	} else {
		goto L44
	}
L3:
	;
	v23 = l0
	v26 = l2
	goto L6
L4:
	;
	goto L5
L5:
	;
	v172 = l0
	v175 = l2
	goto L2
L6:
	;
	v37 = v23 + int32(1)
	if base.Ui32(v15) <= base.Ui32(v37) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v172 = v157
	v175 = v160
	goto L2
L8:
	;
	if base.Ui32(v15) <= base.Ui32(v157) {
		v172 = v157
		v175 = v160
		goto L2
	} else {
		goto L41
	}
L9:
	;
	v157 = v37
	v160 = v26
	goto L8
L10:
	;
	goto L11
L11:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v41 = v37
	v44 = v26
	v50 = v39
	v51 = int32(0)
	goto L12
L12:
	;
	if v50&int32(1) != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v157 = v132
	v160 = v145
	goto L8
L14:
	;
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v51))|base.B2i32(base.Ui32(v15) <= base.Ui32(v132)) != 0 {
		v157 = v132
		v160 = v145
		goto L8
	} else {
		goto L39
	}
L15:
	;
	v56 = int32(-1)
	v58 = v41 + int32(2)
	if base.Ui32(v15) < base.Ui32(v58) {
		v198 = v56
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v126)
	v128 = int32(1)
	v132 = v41 + v128
	v145 = v44 + v128
	goto L14
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v65 = v61&int32(15) + int32(3)
	if v65 != int32(18) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v74 = v65
	v75 = v58
	goto L21
L20:
	;
	if base.Ui32(v15) <= base.Ui32(v58) {
		v198 = v56
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v80 = v61<<(uint(int32(4))%32)&int32(3840) | v60
	if base.B2i32(v80 == int32(0))|base.B2i32(v44-l2 < v80) != 0 {
		v198 = v56
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+2)))
	v74 = v69 + int32(18)
	v75 = v41 + int32(3)
	goto L21
L23:
	;
	v86 = v14 - v44
	if v74 < v86 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v88 = v74
	goto L26
L25:
	;
	v88 = v86
	goto L26
L26:
	;
	if v80 < v88 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v91 = v80
	v93 = v44
	v95 = v88
	goto L30
L28:
	;
	v111 = v80
	v113 = v44
	v115 = v88
	goto L29
L29:
	;
	if v115 != 0 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	if v91 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v111 = v108
	v113 = v105
	v115 = v106
	goto L29
L32:
	;
	base.MemoryCopy(m, v93, v93-v91, v91)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v105 = v91 + v93
	v106 = v95 - v91
	v108 = v91 << (uint(int32(1)) % 32)
	if v108 < v106 {
		v91 = v108
		v93 = v105
		v95 = v106
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	base.MemoryCopy(m, v113, v113-v111, v115)
	goto L38
L37:
	;
	goto L38
L38:
	;
	v132 = v75
	v145 = v113 + v115
	goto L14
L39:
	;
	v150 = int32(1)
	if base.Ui32(v145) < base.Ui32(v14) {
		v41 = v132
		v44 = v145
		v50 = int32(base.Ui32(v50&int32(254)) >> (uint(v150) % 32))
		v51 = v51 + v150
		goto L12
	} else {
		goto L40
	}
L40:
	;
	goto L13
L41:
	;
	if base.Ui32(v160) < base.Ui32(v14) {
		v23 = v157
		v26 = v160
		goto L6
	} else {
		goto L42
	}
L42:
	;
	goto L7
L43:
	;
	if base.B2i32(v172 != v15)|base.B2i32(v175 != v14) != 0 {
		v198 = int32(-1)
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v198 = v175 - l2
	goto L1
L46:
	;
	goto L45
}
func F_pgstatginindex_internal(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v3)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+28)) = uint16(v3)
	v18 = F_relation_open(m, l0, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+119)))
		if v23 != int32(105) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int64(0)
				} else {
					v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v102 + int32(4)
					F_errmsg(m, int32(_a_F_pgstatginindex_internal_0), v9+int32(-48))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(548), int32(_a_F_pgstatginindex_internal_2))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+84))
			if v26 != int32(2742) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int64(0)
					} else {
						v102 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v102 + int32(4)
						F_errmsg(m, int32(_a_F_pgstatginindex_internal_0), v9+int32(-48))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(548), int32(_a_F_pgstatginindex_internal_2))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
				if v29 == int32(116) {
					v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+24)))
					if v32 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_pgstatginindex_internal_3), int32(0))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(558), int32(_a_F_pgstatginindex_internal_2))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(325))
								mBase = m.M
								v138 = m.ExcPending
								if v138 != 0 {
									return int64(0)
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v11))) = v139 + int32(4)
									F_errmsg(m, int32(_a_F_pgstatginindex_internal_4), v11)
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(565), int32(_a_F_pgstatginindex_internal_2))
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v40 = F_ReadBuffer(m, v18, int32(0))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								F_LockBufferInternal(m, v40, int32(1))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int64(0)
								} else {
									if v40 < int32(0) {
										v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatginindex_internal[0]))
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+(v40^int32(-1))<<(uint(int32(2))%32))))
										v62 = v54
									} else {
										v56 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatginindex_internal[1]))
										v62 = v56 + v40<<(uint(int32(13))%32) + int32(-8192)
									}
									v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)+40))
									v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+36)))
									v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+72)))
									F_UnlockReleaseBuffer(m, v40)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int64(0)
									} else {
										F_relation_close(m, v18, int32(1))
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return int64(0)
										} else {
											v74 = F_get_call_result_type(m, l1, int32(0), v9+int32(-4))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int64(0)
											} else {
												if v74 != int32(1) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return int64(0)
													} else {
														F_errmsg_internal(m, int32(_a_F_pgstatginindex_internal_5), int32(0))
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(586), int32(_a_F_pgstatginindex_internal_2))
															mBase = m.M
															v163 = m.ExcPending
															if v163 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v63
													*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v64
													*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
													v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
													v86 = F_heap_form_tuple(m, v81, v9+int32(-32), v9+int32(-36))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return int64(0)
													} else {
														v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
														v89 = F_HeapTupleHeaderGetDatum(m, v88)
														mBase = m.M
														v90 = m.ExcPending
														if v90 != 0 {
															return int64(0)
														} else {
															m.G0 = v11 - int32(-64)
															return v89
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
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v18)+192))
					v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)))
					if v36 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v138 = m.ExcPending
							if v138 != 0 {
								return int64(0)
							} else {
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v18)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v139 + int32(4)
								F_errmsg(m, int32(_a_F_pgstatginindex_internal_4), v11)
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(565), int32(_a_F_pgstatginindex_internal_2))
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v40 = F_ReadBuffer(m, v18, int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							F_LockBufferInternal(m, v40, int32(1))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								if v40 < int32(0) {
									v48 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatginindex_internal[0]))
									v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+(v40^int32(-1))<<(uint(int32(2))%32))))
									v62 = v54
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, _c_F_pgstatginindex_internal[1]))
									v62 = v56 + v40<<(uint(int32(13))%32) + int32(-8192)
								}
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v62)+40))
								v64 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v62)+36)))
								v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+72)))
								F_UnlockReleaseBuffer(m, v40)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return int64(0)
								} else {
									F_relation_close(m, v18, int32(1))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int64(0)
									} else {
										v74 = F_get_call_result_type(m, l1, int32(0), v9+int32(-4))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int64(0)
										} else {
											if v74 != int32(1) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int64(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_pgstatginindex_internal_5), int32(0))
													mBase = m.M
													v158 = m.ExcPending
													if v158 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_pgstatginindex_internal_1), int32(586), int32(_a_F_pgstatginindex_internal_2))
														mBase = m.M
														v163 = m.ExcPending
														if v163 != 0 {
															return int64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v63
												*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v64
												*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
												v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
												v86 = F_heap_form_tuple(m, v81, v9+int32(-32), v9+int32(-36))
												mBase = m.M
												v87 = m.ExcPending
												if v87 != 0 {
													return int64(0)
												} else {
													v88 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
													v89 = F_HeapTupleHeaderGetDatum(m, v88)
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return int64(0)
													} else {
														m.G0 = v11 - int32(-64)
														return v89
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
func F_pgstatindex_v1_5(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = F_textToQualifiedNameList(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int64(0)
		} else {
			v9 = F_makeRangeVarFromNameList(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int64(0)
			} else {
				v12 = F_relation_openrv(m, v9, int32(1))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int64(0)
				} else {
					v14 = F_pgstatindex_impl(m, v12, l0)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int64(0)
					} else {
						return v14
					}
				}
			}
		}
	}
}
func F_pgstatindexbyid_v1_5(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_relation_open(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = F_pgstatindex_impl(m, v4, l0)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return v8
		}
	}
}
func F_pgstattuplebyid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pgstattuplebyid_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstattuplebyid_1), int32(218), int32(_a_F_pgstattuplebyid_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v28 = F_relation_open(m, base.I32_wrap_i64(v3), int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = F_pgstat_relation(m, v28, l0)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					return v30
				}
			}
		}
	}
}
func F_phraseto_tsquery_byid(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14390(m, l0, int32(1), int32(4))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_plan_elem_desc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v12
	F_appendStringInfo(m, l0, int32(_a_F_plan_elem_desc_0), v10)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		F_appendStringInfoString(m, l0, int32(_a_F_plan_elem_desc_1))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
			F_array_desc(m, l0, v26, int32(2), v28, int32(251), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+10)))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v33 + v34<<(uint(int32(1))%32)
				F_appendStringInfoString(m, l0, int32(_a_F_plan_elem_desc_2))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			}
		}
	}
}
func F_planstate_tree_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	return v231
L4:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v46 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v23 = int32(0)
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v23<<(uint(int32(2))%32))))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v33 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v32, l2)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	if v33 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v36 = v23 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v36 < v37 {
		v23 = v36
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v55 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v49 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v46, l2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v49 == int32(0) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	return int32(1)
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	switch v64 - int32(338) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L20
	case 3:
		goto L24
	case 4:
		goto L23
	case 13:
		goto L22
	case 21:
		goto L21
	}
L17:
	;
	v58 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v55, l2)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v58 == int32(0) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	return int32(1)
L20:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v195 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L21:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v162 == int32(0) {
		goto L20
	} else {
		goto L53
	}
L22:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v156 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v155, l2)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L51
	}
L23:
	;
	v133 = int32(0)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v134 <= v133 {
		goto L20
	} else {
		goto L45
	}
L24:
	;
	v111 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v112 <= v111 {
		goto L20
	} else {
		goto L39
	}
L25:
	;
	v89 = int32(0)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v90 <= v89 {
		goto L20
	} else {
		goto L33
	}
L26:
	;
	v67 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v68 <= v67 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v76 = v67
	goto L28
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v71+v76<<(uint(int32(2))%32))))
	v84 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v83, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L20
L30:
	;
	if v84 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L31
	}
L31:
	;
	v87 = v76 + int32(1)
	if v68 != v87 {
		v76 = v87
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v98 = v89
	goto L34
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v93+v98<<(uint(int32(2))%32))))
	v106 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v105, l2)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L20
L36:
	;
	if v106 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v109 = v98 + int32(1)
	if v90 != v109 {
		v98 = v109
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v120 = v111
	goto L40
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v115+v120<<(uint(int32(2))%32))))
	v128 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v127, l2)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L20
L42:
	;
	if v128 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L43
	}
L43:
	;
	v131 = v120 + int32(1)
	if v112 != v131 {
		v120 = v131
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v142 = v133
	goto L46
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v137+v142<<(uint(int32(2))%32))))
	v150 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v149, l2)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	goto L20
L48:
	;
	if v150 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L49
	}
L49:
	;
	v153 = v142 + int32(1)
	if v134 != v153 {
		v142 = v153
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	if v156 == int32(0) {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	return int32(1)
L53:
	;
	v165 = int32(0)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v166 <= v165 {
		goto L20
	} else {
		goto L54
	}
L54:
	;
	v173 = v165
	goto L55
L55:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v177+v173<<(uint(int32(2))%32))))
	v182 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v181, l2)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L57
	}
L56:
	;
	goto L20
L57:
	;
	if v182 != 0 {
		v231 = int32(1)
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v185 = v173 + int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	if v185 < v186 {
		v173 = v185
		goto L55
	} else {
		goto L59
	}
L59:
	;
	goto L56
L60:
	;
	return int32(0)
L61:
	;
	goto L62
L62:
	;
	v200 = int32(0)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v201 <= v200 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	return int32(0)
L64:
	;
	goto L65
L65:
	;
	v209 = v200
	goto L66
L66:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v209<<(uint(int32(2))%32))))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	v219 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v218, l2)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v231 = v219
	goto L3
L68:
	;
	if v219 != 0 {
		v231 = v219
		goto L3
	} else {
		goto L69
	}
L69:
	;
	v222 = v209 + int32(1)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v222 < v223 {
		v209 = v222
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
}
func F_porter_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v670 int32
	_ = v670
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v720 int32
	_ = v720
	var v731 int32
	_ = v731
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v763 int32
	_ = v763
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1008 int32
	_ = v1008
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1124 int32
	_ = v1124
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1200 int32
	_ = v1200
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1209 int32
	_ = v1209
	var v1211 int32
	_ = v1211
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1232 int32
	_ = v1232
	var v1238 int32
	_ = v1238
	var v1245 int32
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1303 int32
	_ = v1303
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	v9 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 == v10 {
		v31 = v9
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v1303
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v37 = v31
	goto L8
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12+v7))))
	if v14 != int32(121) {
		v31 = v9
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = int32(1)
	v18 = v7 + v17
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18
	v24 = F_slice_from_s(m, l0, v17, int32(_a_F_porter_ISO_8859_1_stem_0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v24 < int32(0) {
		v1303 = v24
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v31 = int32(0)
	goto L2
L8:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = v39
	goto L11
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v99
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v99
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v136 < v7 {
		goto L37
	} else {
		goto L38
	}
L10:
	;
	goto L9
L11:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v55 < v54 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v41
	v115 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v100 + v115
	v121 = F_slice_from_s(m, l0, v115, int32(_a_F_porter_ISO_8859_1_stem_1))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L33
	}
L13:
	;
	goto L12
L14:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v98 != 0 {
		goto L28
	} else {
		goto L29
	}
L15:
	;
	v57 = v54
	goto L17
L16:
	;
	v57 = v55
	goto L17
L17:
	;
	goto L19
L18:
	;
	v98 = v93
	goto L14
L19:
	;
	if v54 == v57 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v93 = int32(0)
	goto L18
L21:
	;
	v98 = int32(-1)
	goto L14
L22:
	;
	goto L23
L23:
	;
	v69 = int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v54))))
	if int32(121) < v72 {
		v93 = v69
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v74 = v72 - int32(97)
	if v74 < int32(0) {
		v93 = v69
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v74)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v80)>>(uint(v74&int32(7))%32))&int32(1) == int32(0) {
		v93 = v69
		goto L18
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54 + int32(1)
	goto L27
L27:
	;
	goto L20
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v41
	if v99 <= v41 {
		goto L10
	} else {
		goto L32
	}
L29:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v100
	if v99 == v100 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103+v100))))
	if v105 == int32(121) {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v112 = v41 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v112
	v41 = v112
	goto L11
L33:
	;
	if int32(0) <= v121 {
		v37 = int32(0)
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v1303 = v121
	goto L1
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v355
	if v355 <= v7 {
		goto L98
	} else {
		goto L99
	}
L36:
	;
	if v176 < int32(0) {
		goto L35
	} else {
		goto L51
	}
L37:
	;
	v138 = v7
	goto L39
L38:
	;
	v138 = v136
	goto L39
L39:
	;
	v145 = v7
	goto L41
L40:
	;
	v176 = v156
	goto L36
L41:
	;
	if v145 == v138 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v176 = int32(-1)
	goto L36
L44:
	;
	goto L45
L45:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149+v145))))
	if int32(121) < v151 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v168 = v145 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v168
	v145 = v168
	goto L41
L47:
	;
	v153 = v151 - int32(97)
	if v153 < int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v156 = int32(1)
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v153)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v160)>>(uint(v153&int32(7))%32))&v156 != 0 {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L46
L51:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v180 = v179 + v176
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v180
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v191 < v180 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v234 < int32(0) {
		goto L35
	} else {
		goto L66
	}
L53:
	;
	v193 = v180
	goto L55
L54:
	;
	v193 = v191
	goto L55
L55:
	;
	v199 = v180
	goto L57
L56:
	;
	v234 = int32(1)
	goto L52
L57:
	;
	if v199 == v193 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v234 = int32(-1)
	goto L52
L60:
	;
	goto L61
L61:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206+v199))))
	if int32(121) < v208 {
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v210 = v208 - int32(97)
	if v210 < int32(0) {
		goto L56
	} else {
		goto L63
	}
L63:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v210)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v216)>>(uint(v210&int32(7))%32))&int32(1) == int32(0) {
		goto L56
	} else {
		goto L64
	}
L64:
	;
	v225 = v199 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v225
	v199 = v225
	goto L57
L66:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v238 = v237 + v234
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v238
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v249 < v238 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v289 < int32(0) {
		goto L35
	} else {
		goto L82
	}
L68:
	;
	v251 = v238
	goto L70
L69:
	;
	v251 = v249
	goto L70
L70:
	;
	v258 = v238
	goto L72
L71:
	;
	v289 = v269
	goto L67
L72:
	;
	if v258 == v251 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v289 = int32(-1)
	goto L67
L75:
	;
	goto L76
L76:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262+v258))))
	if int32(121) < v264 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v281 = v258 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v281
	v258 = v281
	goto L72
L78:
	;
	v266 = v264 - int32(97)
	if v266 < int32(0) {
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v269 = int32(1)
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v266)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v273)>>(uint(v266&int32(7))%32))&v269 != 0 {
		goto L71
	} else {
		goto L80
	}
L80:
	;
	goto L77
L82:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v293 = v292 + v289
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v293
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v304 < v293 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v347 < int32(0) {
		goto L35
	} else {
		goto L97
	}
L84:
	;
	v306 = v293
	goto L86
L85:
	;
	v306 = v304
	goto L86
L86:
	;
	v312 = v293
	goto L88
L87:
	;
	v347 = int32(1)
	goto L83
L88:
	;
	if v312 == v306 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v347 = int32(-1)
	goto L83
L91:
	;
	goto L92
L92:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319+v312))))
	if int32(121) < v321 {
		goto L87
	} else {
		goto L93
	}
L93:
	;
	v323 = v321 - int32(97)
	if v323 < int32(0) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v323)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v329)>>(uint(v323&int32(7))%32))&int32(1) == int32(0) {
		goto L87
	} else {
		goto L95
	}
L95:
	;
	v338 = v312 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v338
	v312 = v338
	goto L88
L97:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v350 + v347
	goto L35
L98:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v393
	v397 = v393 - int32(1)
	v398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v397 <= v398 {
		goto L111
	} else {
		goto L112
	}
L99:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359+v355-int32(1)))))
	if v363 != int32(115) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v369 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_2), int32(4), int32(0))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L5
	} else {
		goto L101
	}
L101:
	;
	if v369 == int32(0) {
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v373
	switch v369 - int32(1) {
	case 0:
		goto L105
	case 1:
		goto L104
	case 2:
		goto L103
	default:
		goto L98
	}
L103:
	;
	v389 = F_slice_del(m, l0)
	mBase = m.M
	if v389 < int32(0) {
		v1303 = v389
		goto L1
	} else {
		goto L110
	}
L104:
	;
	v385 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_ISO_8859_1_stem_3))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L5
	} else {
		goto L108
	}
L105:
	;
	v379 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_4))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	if int32(0) <= v379 {
		goto L98
	} else {
		goto L107
	}
L107:
	;
	v1303 = v379
	goto L1
L108:
	;
	if int32(0) <= v385 {
		goto L98
	} else {
		goto L109
	}
L109:
	;
	v1303 = v385
	goto L1
L110:
	;
	goto L98
L111:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v705
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v705
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v705 <= v708 {
		goto L187
	} else {
		goto L188
	}
L112:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400+v397))))
	switch v402 - int32(100) {
	case 0, 3:
		goto L113
	default:
		goto L111
	}
L113:
	;
	v408 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_5), int32(3), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L5
	} else {
		goto L114
	}
L114:
	;
	if v408 == int32(0) {
		goto L111
	} else {
		goto L115
	}
L115:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v412
	switch v408 - int32(1) {
	case 0:
		goto L117
	case 1:
		goto L116
	default:
		goto L111
	}
L116:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v439 = v432
	goto L123
L117:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v412 < v416 {
		goto L111
	} else {
		goto L118
	}
L118:
	;
	v420 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_6))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L5
	} else {
		goto L119
	}
L119:
	;
	if int32(0) <= v420 {
		goto L111
	} else {
		goto L120
	}
L120:
	;
	v1303 = v420
	goto L1
L121:
	;
	if v473 < int32(0) {
		goto L111
	} else {
		goto L133
	}
L122:
	;
	v473 = v453
	goto L121
L123:
	;
	if v439 <= v433 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v473 = int32(-1)
	goto L121
L126:
	;
	goto L127
L127:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444+v439-int32(1)))))
	if int32(121) < v448 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v465 = v439 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v465
	v439 = v465
	goto L123
L129:
	;
	v450 = v448 - int32(97)
	if v450 < int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v453 = int32(1)
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v450)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v457)>>(uint(v450&int32(7))%32))&v453 != 0 {
		goto L122
	} else {
		goto L131
	}
L131:
	;
	goto L128
L133:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v476 + (v412 - v424)
	v480 = F_slice_del(m, l0)
	mBase = m.M
	if v480 < int32(0) {
		v1303 = v480
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v485 = v483 - v484
	v487 = v483 - int32(1)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v487 <= v488 {
		v532 = v483
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v532 != v533 {
		goto L111
	} else {
		goto L145
	}
L136:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490+v487))))
	if base.B2i32(v492&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v492)%32)&int32(68514004) == int32(0)) != 0 {
		v532 = v483
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v507 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_7), int32(13), int32(0))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L5
	} else {
		goto L138
	}
L138:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v510 = v509 + v485
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v510
	switch v507 - int32(1) {
	case 0:
		goto L140
	case 1:
		goto L139
	case 2:
		v532 = v510
		goto L135
	default:
		goto L111
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v510
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v510 <= v522 {
		goto L111
	} else {
		goto L143
	}
L140:
	;
	v516 = F_insert_s(m, l0, v510, v510, int32(1), int32(_a_F_porter_ISO_8859_1_stem_8))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L5
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v510
	if int32(0) <= v516 {
		goto L111
	} else {
		goto L142
	}
L142:
	;
	v1303 = v516
	goto L1
L143:
	;
	v525 = v510 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v525
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v525
	v528 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v528 {
		goto L111
	} else {
		goto L144
	}
L144:
	;
	v1303 = v528
	goto L1
L145:
	;
	v535 = int32(0)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L149
L146:
	;
	if v689 == int32(0) {
		goto L111
	} else {
		goto L184
	}
L147:
	;
	if v584 != 0 {
		v689 = v535
		goto L146
	} else {
		goto L159
	}
L148:
	;
	v584 = v581
	goto L147
L149:
	;
	if v543 <= v544 {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v581 = int32(0)
	goto L148
L151:
	;
	v584 = int32(-1)
	goto L147
L152:
	;
	goto L153
L153:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555+v543-int32(1)))))
	if int32(121) < v559 {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v543 - int32(1)
	goto L158
L155:
	;
	v561 = v559 - int32(89)
	if v561 < int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v564 = int32(1)
	v568 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v561)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v568)>>(uint(v561&int32(7))%32))&v564 != 0 {
		v581 = v564
		goto L148
	} else {
		goto L157
	}
L157:
	;
	goto L154
L158:
	;
	goto L150
L159:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L162
L160:
	;
	if v637 != 0 {
		v689 = v535
		goto L146
	} else {
		goto L171
	}
L161:
	;
	v637 = v633
	goto L160
L162:
	;
	if v593 <= v594 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	v633 = int32(0)
	goto L161
L164:
	;
	v637 = int32(-1)
	goto L160
L165:
	;
	goto L166
L166:
	;
	v606 = int32(1)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v607+v593-v606))))
	if int32(121) < v611 {
		v633 = v606
		goto L161
	} else {
		goto L167
	}
L167:
	;
	v613 = v611 - int32(97)
	if v613 < int32(0) {
		v633 = v606
		goto L161
	} else {
		goto L168
	}
L168:
	;
	v619 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v613)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v619)>>(uint(v613&int32(7))%32))&int32(1) == int32(0) {
		v633 = v606
		goto L161
	} else {
		goto L169
	}
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v593 - int32(1)
	goto L170
L170:
	;
	goto L163
L171:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L174
L172:
	;
	v689 = base.B2i32(v686 == int32(0))
	goto L146
L173:
	;
	v686 = v683
	goto L172
L174:
	;
	if v645 <= v646 {
		goto L176
	} else {
		goto L177
	}
L175:
	;
	v683 = int32(0)
	goto L173
L176:
	;
	v686 = int32(-1)
	goto L172
L177:
	;
	goto L178
L178:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v657+v645-int32(1)))))
	if int32(121) < v661 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v645 - int32(1)
	goto L183
L180:
	;
	v663 = v661 - int32(97)
	if v663 < int32(0) {
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v666 = int32(1)
	v670 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v663)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v670)>>(uint(v663&int32(7))%32))&v666 != 0 {
		v683 = v666
		goto L173
	} else {
		goto L182
	}
L182:
	;
	goto L179
L183:
	;
	goto L175
L184:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v693 = v692 + v485
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v693
	v697 = F_insert_s(m, l0, v693, v693, int32(1), int32(_a_F_porter_ISO_8859_1_stem_9))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L5
	} else {
		goto L185
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v693
	if v697 < int32(0) {
		v1303 = v697
		goto L1
	} else {
		goto L186
	}
L186:
	;
	goto L111
L187:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v784
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v784
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v784-int32(2) <= v787 {
		goto L205
	} else {
		goto L206
	}
L188:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v714 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v710+v705-int32(1)))))
	if v714|int32(32) != int32(121) {
		goto L187
	} else {
		goto L189
	}
L189:
	;
	v720 = v705 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v720
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v720
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v737 = v720
	goto L192
L190:
	;
	if v771 < int32(0) {
		goto L187
	} else {
		goto L202
	}
L191:
	;
	v771 = v751
	goto L190
L192:
	;
	if v737 <= v731 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v771 = int32(-1)
	goto L190
L195:
	;
	goto L196
L196:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742+v737-int32(1)))))
	if int32(121) < v746 {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v763 = v737 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v763
	v737 = v763
	goto L192
L198:
	;
	v748 = v746 - int32(97)
	if v748 < int32(0) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v751 = int32(1)
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v748)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v755)>>(uint(v748&int32(7))%32))&v751 != 0 {
		goto L191
	} else {
		goto L200
	}
L200:
	;
	goto L197
L202:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v774 - v771
	v779 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_ISO_8859_1_stem_10))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L5
	} else {
		goto L203
	}
L203:
	;
	if v779 < int32(0) {
		v1303 = v779
		goto L1
	} else {
		goto L204
	}
L204:
	;
	goto L187
L205:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v900
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v900
	v903 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v900-int32(2) <= v903 {
		goto L250
	} else {
		goto L251
	}
L206:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v793 = int32(1)
	v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v791+v784-v793))))
	if base.B2i32(v795&int32(224) != int32(96))|base.B2i32(v793<<(uint(v795)%32)&int32(_a_F_porter_ISO_8859_1_stem_11) == int32(0)) != 0 {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v810 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_12), int32(20), int32(0))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L5
	} else {
		goto L208
	}
L208:
	;
	if v810 == int32(0) {
		goto L205
	} else {
		goto L209
	}
L209:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v814
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v814 < v816 {
		goto L205
	} else {
		goto L210
	}
L210:
	;
	switch v810 - int32(1) {
	case 0:
		goto L223
	case 1:
		goto L222
	case 2:
		goto L221
	case 3:
		goto L220
	case 4:
		goto L219
	case 5:
		goto L218
	case 6:
		goto L217
	case 7:
		goto L216
	case 8:
		goto L215
	case 9:
		goto L214
	case 10:
		goto L213
	case 11:
		goto L212
	case 12:
		goto L211
	default:
		goto L205
	}
L211:
	;
	v894 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_13))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L5
	} else {
		goto L248
	}
L212:
	;
	v888 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_14))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L5
	} else {
		goto L246
	}
L213:
	;
	v882 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_15))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L5
	} else {
		goto L244
	}
L214:
	;
	v876 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_16))
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L5
	} else {
		goto L242
	}
L215:
	;
	v870 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_17))
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L5
	} else {
		goto L240
	}
L216:
	;
	v864 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_18))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L5
	} else {
		goto L238
	}
L217:
	;
	v858 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_19))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L5
	} else {
		goto L236
	}
L218:
	;
	v852 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_ISO_8859_1_stem_20))
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L5
	} else {
		goto L234
	}
L219:
	;
	v846 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_ISO_8859_1_stem_21))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L5
	} else {
		goto L232
	}
L220:
	;
	v840 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_ISO_8859_1_stem_22))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L5
	} else {
		goto L230
	}
L221:
	;
	v834 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_ISO_8859_1_stem_23))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L5
	} else {
		goto L228
	}
L222:
	;
	v828 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_ISO_8859_1_stem_24))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L5
	} else {
		goto L226
	}
L223:
	;
	v822 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_ISO_8859_1_stem_25))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L5
	} else {
		goto L224
	}
L224:
	;
	if int32(0) <= v822 {
		goto L205
	} else {
		goto L225
	}
L225:
	;
	v1303 = v822
	goto L1
L226:
	;
	if int32(0) <= v828 {
		goto L205
	} else {
		goto L227
	}
L227:
	;
	v1303 = v828
	goto L1
L228:
	;
	if int32(0) <= v834 {
		goto L205
	} else {
		goto L229
	}
L229:
	;
	v1303 = v834
	goto L1
L230:
	;
	if int32(0) <= v840 {
		goto L205
	} else {
		goto L231
	}
L231:
	;
	v1303 = v840
	goto L1
L232:
	;
	if int32(0) <= v846 {
		goto L205
	} else {
		goto L233
	}
L233:
	;
	v1303 = v846
	goto L1
L234:
	;
	if int32(0) <= v852 {
		goto L205
	} else {
		goto L235
	}
L235:
	;
	v1303 = v852
	goto L1
L236:
	;
	if int32(0) <= v858 {
		goto L205
	} else {
		goto L237
	}
L237:
	;
	v1303 = v858
	goto L1
L238:
	;
	if int32(0) <= v864 {
		goto L205
	} else {
		goto L239
	}
L239:
	;
	v1303 = v864
	goto L1
L240:
	;
	if int32(0) <= v870 {
		goto L205
	} else {
		goto L241
	}
L241:
	;
	v1303 = v870
	goto L1
L242:
	;
	if int32(0) <= v876 {
		goto L205
	} else {
		goto L243
	}
L243:
	;
	v1303 = v876
	goto L1
L244:
	;
	if int32(0) <= v882 {
		goto L205
	} else {
		goto L245
	}
L245:
	;
	v1303 = v882
	goto L1
L246:
	;
	if int32(0) <= v888 {
		goto L205
	} else {
		goto L247
	}
L247:
	;
	v1303 = v888
	goto L1
L248:
	;
	if v894 < int32(0) {
		v1303 = v894
		goto L1
	} else {
		goto L249
	}
L249:
	;
	goto L205
L250:
	;
	v953 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v953
	v955 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v953
	v959 = v953 - int32(1)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v959 <= v960 {
		v1014 = v955
		goto L264
	} else {
		goto L265
	}
L251:
	;
	v907 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v909 = int32(1)
	v911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v907+v900-v909))))
	if base.B2i32(v911&int32(224) != int32(96))|base.B2i32(v909<<(uint(v911)%32)&int32(_a_F_porter_ISO_8859_1_stem_26) == int32(0)) != 0 {
		goto L250
	} else {
		goto L252
	}
L252:
	;
	v926 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_27), int32(7), int32(0))
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L5
	} else {
		goto L253
	}
L253:
	;
	if v926 == int32(0) {
		goto L250
	} else {
		goto L254
	}
L254:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v930
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v930 < v932 {
		goto L250
	} else {
		goto L255
	}
L255:
	;
	switch v926 - int32(1) {
	case 0:
		goto L258
	case 1:
		goto L257
	case 2:
		goto L256
	default:
		goto L250
	}
L256:
	;
	v948 = F_slice_del(m, l0)
	mBase = m.M
	if v948 < int32(0) {
		v1303 = v948
		goto L1
	} else {
		goto L263
	}
L257:
	;
	v944 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_28))
	mBase = m.M
	v945 = m.ExcPending
	if v945 != 0 {
		goto L5
	} else {
		goto L261
	}
L258:
	;
	v938 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_ISO_8859_1_stem_29))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L5
	} else {
		goto L259
	}
L259:
	;
	if int32(0) <= v938 {
		goto L250
	} else {
		goto L260
	}
L260:
	;
	v1303 = v938
	goto L1
L261:
	;
	if int32(0) <= v944 {
		goto L250
	} else {
		goto L262
	}
L262:
	;
	v1303 = v944
	goto L1
L263:
	;
	goto L250
L264:
	;
	if v1014 < int32(0) {
		v1303 = v1014
		goto L1
	} else {
		goto L277
	}
L265:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962+v959))))
	if base.B2i32(v964&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v964)%32)&int32(_a_F_porter_ISO_8859_1_stem_30) == int32(0)) != 0 {
		v1014 = v955
		goto L264
	} else {
		goto L266
	}
L266:
	;
	v979 = F_find_among_b(m, l0, int32(_a_F_porter_ISO_8859_1_stem_31), int32(19), int32(0))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L5
	} else {
		goto L267
	}
L267:
	;
	if v979 == int32(0) {
		v1014 = v955
		goto L264
	} else {
		goto L268
	}
L268:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v983
	v985 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v983 < v985 {
		v1014 = v955
		goto L264
	} else {
		goto L269
	}
L269:
	;
	switch v979 - int32(1) {
	case 0:
		goto L272
	case 1:
		goto L271
	default:
		goto L270
	}
L270:
	;
	v1014 = int32(1)
	goto L264
L271:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v983 <= v992 {
		v1014 = v955
		goto L264
	} else {
		goto L274
	}
L272:
	;
	v989 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v989 {
		goto L270
	} else {
		goto L273
	}
L273:
	;
	v1014 = v989
	goto L264
L274:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v996 = int32(1)
	v998 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994+v983-v996))))
	if base.Ui32(v996) < base.Ui32((v998-int32(115))&int32(255)) {
		v1014 = v955
		goto L264
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v983 - int32(1)
	v1008 = F_slice_del(m, l0)
	mBase = m.M
	if v1008 < int32(0) {
		v1014 = v1008
		goto L264
	} else {
		goto L276
	}
L276:
	;
	goto L270
L277:
	;
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1018
	v1020 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1018
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1018 <= v1023 {
		v1204 = v1020
		goto L278
	} else {
		goto L279
	}
L278:
	;
	if v1204 < int32(0) {
		v1303 = v1204
		goto L1
	} else {
		goto L327
	}
L279:
	;
	v1025 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1029 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1025+v1018-int32(1)))))
	if v1029 != int32(101) {
		v1204 = v1020
		goto L278
	} else {
		goto L280
	}
L280:
	;
	v1033 = v1018 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1033
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1033
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1018 <= v1036 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1018 <= v1038 {
		v1204 = v1020
		goto L278
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	v1200 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1200 {
		goto L324
	} else {
		goto L325
	}
L284:
	;
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L288
L285:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1194 + (v1033 - v1040)
	goto L283
L286:
	;
	if v1089 != 0 {
		goto L285
	} else {
		goto L298
	}
L287:
	;
	v1089 = v1086
	goto L286
L288:
	;
	if v1048 <= v1049 {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	v1086 = int32(0)
	goto L287
L290:
	;
	v1089 = int32(-1)
	goto L286
L291:
	;
	goto L292
L292:
	;
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1064 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1060+v1048-int32(1)))))
	if int32(121) < v1064 {
		goto L293
	} else {
		goto L294
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1048 - int32(1)
	goto L297
L294:
	;
	v1066 = v1064 - int32(89)
	if v1066 < int32(0) {
		goto L293
	} else {
		goto L295
	}
L295:
	;
	v1069 = int32(1)
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1066)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v1073)>>(uint(v1066&int32(7))%32))&v1069 != 0 {
		v1086 = v1069
		goto L287
	} else {
		goto L296
	}
L296:
	;
	goto L293
L297:
	;
	goto L289
L298:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L301
L299:
	;
	if v1142 != 0 {
		goto L285
	} else {
		goto L310
	}
L300:
	;
	v1142 = v1138
	goto L299
L301:
	;
	if v1098 <= v1099 {
		goto L303
	} else {
		goto L304
	}
L302:
	;
	v1138 = int32(0)
	goto L300
L303:
	;
	v1142 = int32(-1)
	goto L299
L304:
	;
	goto L305
L305:
	;
	v1111 = int32(1)
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1112+v1098-v1111))))
	if int32(121) < v1116 {
		v1138 = v1111
		goto L300
	} else {
		goto L306
	}
L306:
	;
	v1118 = v1116 - int32(97)
	if v1118 < int32(0) {
		v1138 = v1111
		goto L300
	} else {
		goto L307
	}
L307:
	;
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1118)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v1124)>>(uint(v1118&int32(7))%32))&int32(1) == int32(0) {
		v1138 = v1111
		goto L300
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1098 - int32(1)
	goto L309
L309:
	;
	goto L302
L310:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L313
L311:
	;
	if v1191 == int32(0) {
		v1204 = v1020
		goto L278
	} else {
		goto L323
	}
L312:
	;
	v1191 = v1188
	goto L311
L313:
	;
	if v1150 <= v1151 {
		goto L315
	} else {
		goto L316
	}
L314:
	;
	v1188 = int32(0)
	goto L312
L315:
	;
	v1191 = int32(-1)
	goto L311
L316:
	;
	goto L317
L317:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1162+v1150-int32(1)))))
	if int32(121) < v1166 {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1150 - int32(1)
	goto L322
L319:
	;
	v1168 = v1166 - int32(97)
	if v1168 < int32(0) {
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1171 = int32(1)
	v1175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1168)>>(uint(int32(3))%32)))+uint32(_c_F_porter_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v1175)>>(uint(v1168&int32(7))%32))&v1171 != 0 {
		v1188 = v1171
		goto L312
	} else {
		goto L321
	}
L321:
	;
	goto L318
L322:
	;
	goto L314
L323:
	;
	goto L285
L324:
	;
	v1203 = int32(1)
	goto L326
L325:
	;
	v1203 = v1200
	goto L326
L326:
	;
	v1204 = v1203
	goto L278
L327:
	;
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1209
	v1211 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1209
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1209 <= v1218 {
		v1252 = v1211
		goto L329
	} else {
		goto L330
	}
L328:
	;
	if v1252 < int32(0) {
		v1303 = v1252
		goto L1
	} else {
		goto L337
	}
L329:
	;
	goto L328
L330:
	;
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1220+v1209-int32(1)))))
	if v1224 != int32(108) {
		v1252 = v1211
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v1228 = v1209 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1228
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1228
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.B2i32(v1228 <= v1218)|base.B2i32(v1209 <= v1232) != 0 {
		v1252 = v1211
		goto L329
	} else {
		goto L332
	}
L332:
	;
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228+v1220-int32(1)))))
	if v1238 != int32(108) {
		v1252 = v1211
		goto L329
	} else {
		goto L333
	}
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1209 - int32(2)
	v1245 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1245 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1248 = int32(1)
	goto L336
L335:
	;
	v1248 = v1245
	goto L336
L336:
	;
	v1252 = v1248
	goto L329
L337:
	;
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1255
	if v37 != 0 {
		goto L338
	} else {
		goto L339
	}
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1255
	v1303 = int32(1)
	goto L1
L339:
	;
	goto L340
L340:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1264 < v1263 {
		goto L342
	} else {
		goto L343
	}
L341:
	;
	v1303 = v1290
	goto L1
L342:
	;
	v1266 = v1263
	goto L344
L343:
	;
	v1266 = v1264
	goto L344
L344:
	;
	v1268 = v1263
	goto L345
L345:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1268
	if v1268 != v1264 {
		goto L348
	} else {
		goto L349
	}
L346:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1268
	v1285 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1268 + v1285
	v1290 = F_slice_from_s(m, l0, v1285, int32(_a_F_porter_ISO_8859_1_stem_32))
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L5
	} else {
		goto L353
	}
L347:
	;
	goto L346
L348:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1275+v1268))))
	if v1277 == int32(89) {
		goto L347
	} else {
		goto L351
	}
L349:
	;
	goto L350
L350:
	;
	if v1268 == v1266 {
		goto L338
	} else {
		goto L352
	}
L351:
	;
	goto L350
L352:
	;
	v1282 = v1268 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1282
	v1268 = v1282
	goto L345
L353:
	;
	if int32(0) <= v1290 {
		goto L340
	} else {
		goto L354
	}
L354:
	;
	goto L341
}
func F_porter_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v265 int32
	_ = v265
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v324 int32
	_ = v324
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v446 int32
	_ = v446
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v519 int32
	_ = v519
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v570 int32
	_ = v570
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v603 int32
	_ = v603
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v641 int32
	_ = v641
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v692 int32
	_ = v692
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v726 int32
	_ = v726
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v838 int32
	_ = v838
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v850 int32
	_ = v850
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v920 int32
	_ = v920
	var v929 int32
	_ = v929
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1046 int32
	_ = v1046
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1101 int32
	_ = v1101
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1181 int32
	_ = v1181
	var v1188 int32
	_ = v1188
	var v1201 int32
	_ = v1201
	var v1202 int32
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1282 int32
	_ = v1282
	var v1286 int32
	_ = v1286
	var v1288 int32
	_ = v1288
	var v1294 int32
	_ = v1294
	var v1310 int32
	_ = v1310
	var v1317 int32
	_ = v1317
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1332 int32
	_ = v1332
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1360 int32
	_ = v1360
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1411 int32
	_ = v1411
	var v1415 int32
	_ = v1415
	var v1417 int32
	_ = v1417
	var v1423 int32
	_ = v1423
	var v1440 int32
	_ = v1440
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1475 int32
	_ = v1475
	var v1481 int32
	_ = v1481
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1507 int32
	_ = v1507
	var v1514 int32
	_ = v1514
	var v1515 int32
	_ = v1515
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1577 int32
	_ = v1577
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1589 int32
	_ = v1589
	var v1598 int32
	_ = v1598
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1635 int32
	_ = v1635
	var v1637 int32
	_ = v1637
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1658 int32
	_ = v1658
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1724 int32
	_ = v1724
	var v1725 int32
	_ = v1725
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1745 int32
	_ = v1745
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1753 int32
	_ = v1753
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1786 int32
	_ = v1786
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1825 int32
	_ = v1825
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1850 int32
	_ = v1850
	var v1855 int32
	_ = v1855
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1913 int32
	_ = v1913
	var v1914 int32
	_ = v1914
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1944 int32
	_ = v1944
	var v1946 int32
	_ = v1946
	var v1948 int32
	_ = v1948
	var v1966 int32
	_ = v1966
	var v1968 int32
	_ = v1968
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1982 int32
	_ = v1982
	var v1988 int32
	_ = v1988
	var v2005 int32
	_ = v2005
	var v2012 int32
	_ = v2012
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2074 int32
	_ = v2074
	var v2076 int32
	_ = v2076
	var v2078 int32
	_ = v2078
	var v2096 int32
	_ = v2096
	var v2098 int32
	_ = v2098
	var v2106 int32
	_ = v2106
	var v2110 int32
	_ = v2110
	var v2112 int32
	_ = v2112
	var v2118 int32
	_ = v2118
	var v2134 int32
	_ = v2134
	var v2141 int32
	_ = v2141
	var v2154 int32
	_ = v2154
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2172 int32
	_ = v2172
	var v2173 int32
	_ = v2173
	var v2175 int32
	_ = v2175
	var v2177 int32
	_ = v2177
	var v2184 int32
	_ = v2184
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2203 int32
	_ = v2203
	var v2205 int32
	_ = v2205
	var v2207 int32
	_ = v2207
	var v2225 int32
	_ = v2225
	var v2227 int32
	_ = v2227
	var v2235 int32
	_ = v2235
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2247 int32
	_ = v2247
	var v2264 int32
	_ = v2264
	var v2271 int32
	_ = v2271
	var v2274 int32
	_ = v2274
	var v2280 int32
	_ = v2280
	var v2283 int32
	_ = v2283
	var v2286 int32
	_ = v2286
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2298 int32
	_ = v2298
	var v2300 int32
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2308 int32
	_ = v2308
	var v2312 int32
	_ = v2312
	var v2318 int32
	_ = v2318
	var v2325 int32
	_ = v2325
	var v2328 int32
	_ = v2328
	var v2332 int32
	_ = v2332
	var v2335 int32
	_ = v2335
	var v2344 int32
	_ = v2344
	var v2346 int32
	_ = v2346
	var v2353 int32
	_ = v2353
	var v2354 int32
	_ = v2354
	var v2357 int32
	_ = v2357
	var v2366 int32
	_ = v2366
	var v2368 int32
	_ = v2368
	var v2373 int32
	_ = v2373
	var v2375 int32
	_ = v2375
	var v2382 int32
	_ = v2382
	var v2385 int32
	_ = v2385
	var v2389 int32
	_ = v2389
	var v2396 int32
	_ = v2396
	var v2397 int32
	_ = v2397
	var v2411 int32
	_ = v2411
	var v2416 int32
	_ = v2416
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2435 int32
	_ = v2435
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v8
	v10 = int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 == v11 {
		v32 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v2435
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	v40 = v32
	goto L8
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v8))))
	if v15 != int32(121) {
		v32 = v10
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = int32(1)
	v19 = v8 + v18
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19
	v25 = F_slice_from_s(m, l0, v18, int32(_a_F_porter_UTF_8_stem_0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	if v25 < int32(0) {
		v2435 = v25
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v32 = int32(0)
	goto L2
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v43 = v41
	goto L11
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8
	v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v249
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v273 = v8
	goto L72
L10:
	;
	goto L9
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L17
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43
	v238 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v169 + v238
	v244 = F_slice_from_s(m, l0, v238, int32(_a_F_porter_UTF_8_stem_1))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L67
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43
	goto L48
L15:
	;
	if v166 != 0 {
		goto L39
	} else {
		goto L40
	}
L16:
	;
	v166 = v159
	goto L15
L17:
	;
	if v61 <= v60 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v159 = int32(0)
	goto L16
L19:
	;
	v166 = int32(-1)
	goto L15
L20:
	;
	goto L21
L21:
	;
	v77 = int32(1)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v62))))
	if base.Ui32(v79) < base.Ui32(int32(192)) {
		v136 = v79
		v137 = v77
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if int32(121) < v136 {
		v159 = v137
		goto L16
	} else {
		goto L35
	}
L23:
	;
	v83 = v60 + int32(1)
	if v83 == v61 {
		v136 = v79
		v137 = v77
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v62))))
	v88 = v86 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v79) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v62))))
	v104 = v102 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v79) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v92 = v60 + int32(2)
	if v92 != v61 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v136 = v79<<(uint(int32(6))%32)&int32(1984) | v88
	v137 = int32(2)
	goto L22
L29:
	;
	goto L28
L30:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v108))))
	v136 = v121&int32(63) | (v79<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v88<<(uint(int32(12))%32) | v104<<(uint(int32(6))%32))
	v137 = int32(4)
	goto L22
L31:
	;
	v108 = v60 + int32(3)
	if v108 != v61 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v136 = v79<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v88<<(uint(int32(6))%32) | v104
	v137 = int32(3)
	goto L22
L34:
	;
	goto L33
L35:
	;
	v141 = v136 - int32(97)
	if v141 < int32(0) {
		v159 = v137
		goto L16
	} else {
		goto L36
	}
L36:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v141)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v147)>>(uint(v141&int32(7))%32))&int32(1) == int32(0) {
		v159 = v137
		goto L16
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v137 + v60
	goto L38
L38:
	;
	goto L18
L39:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v179 = v167
	v180 = v168
	goto L14
L40:
	;
	goto L41
L41:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v172 == v169 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v179 = v169
	v180 = v171
	goto L14
L43:
	;
	goto L44
L44:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169+v171))))
	if v175 == int32(121) {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v179 = v172
	v180 = v171
	goto L14
L46:
	;
	if v233 < int32(0) {
		goto L10
	} else {
		goto L66
	}
L48:
	;
	goto L49
L49:
	;
	goto L50
L50:
	;
	v188 = v43
	v190 = int32(1)
	goto L53
L52:
	;
	v233 = v218
	goto L46
L53:
	;
	if v179 <= v188 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L52
L55:
	;
	v233 = int32(-1)
	goto L46
L56:
	;
	goto L57
L57:
	;
	v195 = v188 + int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v188))))
	if base.Ui32(v197) < base.Ui32(int32(192)) {
		v218 = v195
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v219 = int32(1)
	if v219 < v190 {
		v188 = v218
		v190 = v190 - v219
		goto L53
	} else {
		goto L65
	}
L59:
	;
	if v179 <= v195 {
		v218 = v195
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v204 = v195
	goto L61
L61:
	;
	v207 = int32(*(*int8)(unsafe.Add(mBase, uint32(v180+v204))))
	if int32(-65) < v207 {
		v218 = v204
		goto L58
	} else {
		goto L63
	}
L62:
	;
	v218 = v179
	goto L58
L63:
	;
	v211 = v204 + int32(1)
	if v211 != v179 {
		v204 = v211
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	goto L54
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v233
	v43 = v233
	goto L11
L67:
	;
	if int32(0) <= v244 {
		v40 = int32(0)
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v2435 = v244
	goto L1
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v8
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v745
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v745
	if v745 <= v8 {
		goto L172
	} else {
		goto L173
	}
L70:
	;
	if v368 < int32(0) {
		goto L69
	} else {
		goto L95
	}
L71:
	;
	v368 = v340
	goto L70
L72:
	;
	if v249 <= v273 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v368 = int32(-1)
	goto L70
L75:
	;
	goto L76
L76:
	;
	v280 = int32(1)
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273+v265))))
	if base.Ui32(v282) < base.Ui32(int32(192)) {
		v339 = v282
		v340 = v280
		goto L77
	} else {
		goto L78
	}
L77:
	;
	if int32(121) < v339 {
		goto L90
	} else {
		goto L91
	}
L78:
	;
	v286 = v273 + int32(1)
	if v286 == v249 {
		v339 = v282
		v340 = v280
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v265))))
	v291 = v289 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v282) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v265))))
	v307 = v305 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v282) {
		goto L86
	} else {
		goto L87
	}
L81:
	;
	v295 = v273 + int32(2)
	if v295 != v249 {
		goto L80
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v339 = v282<<(uint(int32(6))%32)&int32(1984) | v291
	v340 = int32(2)
	goto L77
L84:
	;
	goto L83
L85:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v265+v311))))
	v339 = v324&int32(63) | (v282<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v291<<(uint(int32(12))%32) | v307<<(uint(int32(6))%32))
	v340 = int32(4)
	goto L77
L86:
	;
	v311 = v273 + int32(3)
	if v311 != v249 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v339 = v282<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v291<<(uint(int32(6))%32) | v307
	v340 = int32(3)
	goto L77
L89:
	;
	goto L88
L90:
	;
	v357 = v340 + v273
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v357
	v273 = v357
	goto L72
L91:
	;
	v344 = v339 - int32(97)
	if v344 < int32(0) {
		goto L90
	} else {
		goto L92
	}
L92:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v344)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v350)>>(uint(v344&int32(7))%32))&int32(1) != 0 {
		goto L71
	} else {
		goto L93
	}
L93:
	;
	goto L90
L95:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v372 = v371 + v368
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v372
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v387 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v395 = v372
	goto L98
L96:
	;
	if v491 < int32(0) {
		goto L69
	} else {
		goto L120
	}
L97:
	;
	v491 = v462
	goto L96
L98:
	;
	if v386 <= v395 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v491 = int32(-1)
	goto L96
L101:
	;
	goto L102
L102:
	;
	v402 = int32(1)
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395+v387))))
	if base.Ui32(v404) < base.Ui32(int32(192)) {
		v461 = v404
		v462 = v402
		goto L103
	} else {
		goto L104
	}
L103:
	;
	if int32(121) < v461 {
		goto L97
	} else {
		goto L116
	}
L104:
	;
	v408 = v395 + int32(1)
	if v408 == v386 {
		v461 = v404
		v462 = v402
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v387))))
	v413 = v411 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v404) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v417+v387))))
	v429 = v427 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v404) {
		goto L112
	} else {
		goto L113
	}
L107:
	;
	v417 = v395 + int32(2)
	if v417 != v386 {
		goto L106
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	v461 = v404<<(uint(int32(6))%32)&int32(1984) | v413
	v462 = int32(2)
	goto L103
L110:
	;
	goto L109
L111:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v387+v433))))
	v461 = v446&int32(63) | (v404<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v413<<(uint(int32(12))%32) | v429<<(uint(int32(6))%32))
	v462 = int32(4)
	goto L103
L112:
	;
	v433 = v395 + int32(3)
	if v433 != v386 {
		goto L111
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v461 = v404<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v413<<(uint(int32(6))%32) | v429
	v462 = int32(3)
	goto L103
L115:
	;
	goto L114
L116:
	;
	v466 = v461 - int32(97)
	if v466 < int32(0) {
		goto L97
	} else {
		goto L117
	}
L117:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v466)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v472)>>(uint(v466&int32(7))%32))&int32(1) == int32(0) {
		goto L97
	} else {
		goto L118
	}
L118:
	;
	v480 = v462 + v395
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v480
	v395 = v480
	goto L98
L120:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v495 = v494 + v491
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v495
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v519 = v495
	goto L123
L121:
	;
	if v614 < int32(0) {
		goto L69
	} else {
		goto L146
	}
L122:
	;
	v614 = v586
	goto L121
L123:
	;
	if v510 <= v519 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v614 = int32(-1)
	goto L121
L126:
	;
	goto L127
L127:
	;
	v526 = int32(1)
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519+v511))))
	if base.Ui32(v528) < base.Ui32(int32(192)) {
		v585 = v528
		v586 = v526
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if int32(121) < v585 {
		goto L141
	} else {
		goto L142
	}
L129:
	;
	v532 = v519 + int32(1)
	if v532 == v510 {
		v585 = v528
		v586 = v526
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532+v511))))
	v537 = v535 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v528) {
		goto L132
	} else {
		goto L133
	}
L131:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v541+v511))))
	v553 = v551 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v528) {
		goto L137
	} else {
		goto L138
	}
L132:
	;
	v541 = v519 + int32(2)
	if v541 != v510 {
		goto L131
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v585 = v528<<(uint(int32(6))%32)&int32(1984) | v537
	v586 = int32(2)
	goto L128
L135:
	;
	goto L134
L136:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v557))))
	v585 = v570&int32(63) | (v528<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v537<<(uint(int32(12))%32) | v553<<(uint(int32(6))%32))
	v586 = int32(4)
	goto L128
L137:
	;
	v557 = v519 + int32(3)
	if v557 != v510 {
		goto L136
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v585 = v528<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v537<<(uint(int32(6))%32) | v553
	v586 = int32(3)
	goto L128
L140:
	;
	goto L139
L141:
	;
	v603 = v586 + v519
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v603
	v519 = v603
	goto L123
L142:
	;
	v590 = v585 - int32(97)
	if v590 < int32(0) {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v596 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v590)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v596)>>(uint(v590&int32(7))%32))&int32(1) != 0 {
		goto L122
	} else {
		goto L144
	}
L144:
	;
	goto L141
L146:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v618 = v617 + v614
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v618
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v633 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v641 = v618
	goto L149
L147:
	;
	if v737 < int32(0) {
		goto L69
	} else {
		goto L171
	}
L148:
	;
	v737 = v708
	goto L147
L149:
	;
	if v632 <= v641 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	v737 = int32(-1)
	goto L147
L152:
	;
	goto L153
L153:
	;
	v648 = int32(1)
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641+v633))))
	if base.Ui32(v650) < base.Ui32(int32(192)) {
		v707 = v650
		v708 = v648
		goto L154
	} else {
		goto L155
	}
L154:
	;
	if int32(121) < v707 {
		goto L148
	} else {
		goto L167
	}
L155:
	;
	v654 = v641 + int32(1)
	if v654 == v632 {
		v707 = v650
		v708 = v648
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v654+v633))))
	v659 = v657 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v650) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v663+v633))))
	v675 = v673 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v650) {
		goto L163
	} else {
		goto L164
	}
L158:
	;
	v663 = v641 + int32(2)
	if v663 != v632 {
		goto L157
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v707 = v650<<(uint(int32(6))%32)&int32(1984) | v659
	v708 = int32(2)
	goto L154
L161:
	;
	goto L160
L162:
	;
	v692 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633+v679))))
	v707 = v692&int32(63) | (v650<<(uint(int32(18))%32)&int32(_a_F_porter_UTF_8_stem_2) | v659<<(uint(int32(12))%32) | v675<<(uint(int32(6))%32))
	v708 = int32(4)
	goto L154
L163:
	;
	v679 = v641 + int32(3)
	if v679 != v632 {
		goto L162
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v707 = v650<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v659<<(uint(int32(6))%32) | v675
	v708 = int32(3)
	goto L154
L166:
	;
	goto L165
L167:
	;
	v712 = v707 - int32(97)
	if v712 < int32(0) {
		goto L148
	} else {
		goto L168
	}
L168:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v712)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v718)>>(uint(v712&int32(7))%32))&int32(1) == int32(0) {
		goto L148
	} else {
		goto L169
	}
L169:
	;
	v726 = v708 + v641
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v726
	v641 = v726
	goto L149
L171:
	;
	v740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v740 + v737
	goto L69
L172:
	;
	v783 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v783
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v783
	v787 = v783 - int32(1)
	v788 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v787 <= v788 {
		goto L185
	} else {
		goto L186
	}
L173:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v749+v745-int32(1)))))
	if v753 != int32(115) {
		goto L172
	} else {
		goto L174
	}
L174:
	;
	v759 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_4), int32(4), int32(0))
	mBase = m.M
	v760 = m.ExcPending
	if v760 != 0 {
		goto L5
	} else {
		goto L175
	}
L175:
	;
	if v759 == int32(0) {
		goto L172
	} else {
		goto L176
	}
L176:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v763
	switch v759 - int32(1) {
	case 0:
		goto L179
	case 1:
		goto L178
	case 2:
		goto L177
	default:
		goto L172
	}
L177:
	;
	v779 = F_slice_del(m, l0)
	mBase = m.M
	if v779 < int32(0) {
		v2435 = v779
		goto L1
	} else {
		goto L184
	}
L178:
	;
	v775 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_UTF_8_stem_5))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L5
	} else {
		goto L182
	}
L179:
	;
	v769 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_6))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L5
	} else {
		goto L180
	}
L180:
	;
	if int32(0) <= v769 {
		goto L172
	} else {
		goto L181
	}
L181:
	;
	v2435 = v769
	goto L1
L182:
	;
	if int32(0) <= v775 {
		goto L172
	} else {
		goto L183
	}
L183:
	;
	v2435 = v775
	goto L1
L184:
	;
	goto L172
L185:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1466
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1466
	v1469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1466 <= v1469 {
		goto L310
	} else {
		goto L311
	}
L186:
	;
	v790 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v790+v787))))
	switch v792 - int32(100) {
	case 0, 3:
		goto L187
	default:
		goto L185
	}
L187:
	;
	v798 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_7), int32(3), int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L5
	} else {
		goto L188
	}
L188:
	;
	if v798 == int32(0) {
		goto L185
	} else {
		goto L189
	}
L189:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v802
	switch v798 - int32(1) {
	case 0:
		goto L191
	case 1:
		goto L190
	default:
		goto L185
	}
L190:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v838 = v827
	goto L197
L191:
	;
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v802 < v806 {
		goto L185
	} else {
		goto L192
	}
L192:
	;
	v810 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_8))
	mBase = m.M
	v811 = m.ExcPending
	if v811 != 0 {
		goto L5
	} else {
		goto L193
	}
L193:
	;
	if int32(0) <= v810 {
		goto L185
	} else {
		goto L194
	}
L194:
	;
	v2435 = v810
	goto L1
L195:
	;
	if v944 < int32(0) {
		goto L185
	} else {
		goto L213
	}
L196:
	;
	v944 = int32(-1)
	goto L195
L197:
	;
	if v838 <= v828 {
		goto L196
	} else {
		goto L199
	}
L199:
	;
	v845 = int32(1)
	v846 = v838 - v845
	v848 = int32(*(*int8)(unsafe.Add(mBase, uint32(v829+v846))))
	v850 = v848 & int32(255)
	if base.B2i32(v846 == v828)|base.B2i32(int32(0) <= v848) != 0 {
		v908 = v850
		v912 = v845
		goto L200
	} else {
		goto L201
	}
L200:
	;
	if int32(121) < v908 {
		goto L208
	} else {
		goto L209
	}
L201:
	;
	v857 = v850 & int32(63)
	v859 = v838 - int32(2)
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829+v859))))
	v863 = v861 << (uint(int32(6)) % 32)
	if base.B2i32(v859 != v828)&base.B2i32(base.Ui32(v861) < base.Ui32(int32(192))) == int32(0) {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v908 = v863&int32(1984) | v857
	v912 = int32(2)
	goto L200
L203:
	;
	goto L204
L204:
	;
	v876 = v863&int32(4032) | v857
	v878 = v838 - int32(3)
	v880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v829+v878))))
	if base.B2i32(v878 != v828)&base.B2i32(base.Ui32(v880) < base.Ui32(int32(224))) == int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v908 = v880<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v876
	v912 = int32(3)
	goto L200
L206:
	;
	goto L207
L207:
	;
	v898 = int32(4)
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838+v829-v898))))
	v908 = v880<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v900&int32(7)<<(uint(int32(18))%32) | v876
	v912 = v898
	goto L200
L208:
	;
	v929 = v838 - v912
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v929
	v838 = v929
	goto L197
L209:
	;
	v914 = v908 - int32(97)
	if v914 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v914)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v920)>>(uint(v914&int32(7))%32))&int32(1) == int32(0) {
		goto L208
	} else {
		goto L211
	}
L211:
	;
	v944 = v912
	goto L195
L213:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v947 + (v802 - v814)
	v951 = F_slice_del(m, l0)
	mBase = m.M
	if v951 < int32(0) {
		v2435 = v951
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v956 = v954 - v955
	v958 = v954 - int32(1)
	v959 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v958 <= v959 {
		v1055 = v954
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1055 != v1056 {
		goto L185
	} else {
		goto L244
	}
L216:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v961+v958))))
	if base.B2i32(v963&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v963)%32)&int32(68514004) == int32(0)) != 0 {
		v1055 = v954
		goto L215
	} else {
		goto L217
	}
L217:
	;
	v978 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_10), int32(13), int32(0))
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L5
	} else {
		goto L218
	}
L218:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v981 = v980 + v956
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v981
	switch v978 - int32(1) {
	case 0:
		goto L220
	case 1:
		goto L219
	case 2:
		v1055 = v981
		goto L215
	default:
		goto L185
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v981
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L225
L220:
	;
	v987 = F_insert_s(m, l0, v981, v981, int32(1), int32(_a_F_porter_UTF_8_stem_11))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L5
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v981
	if int32(0) <= v987 {
		goto L185
	} else {
		goto L222
	}
L222:
	;
	v2435 = v987
	goto L1
L223:
	;
	if v1046 < int32(0) {
		goto L185
	} else {
		goto L242
	}
L225:
	;
	goto L226
L226:
	;
	goto L227
L227:
	;
	v1001 = v981
	v1003 = int32(1)
	goto L230
L229:
	;
	v1046 = v1028
	goto L223
L230:
	;
	if v1001 <= v994 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L229
L232:
	;
	v1046 = int32(-1)
	goto L223
L233:
	;
	goto L234
L234:
	;
	v1008 = v1001 - int32(1)
	v1010 = int32(*(*int8)(unsafe.Add(mBase, uint32(v993+v1008))))
	if base.B2i32(int32(0) <= v1010)|base.B2i32(v1008 <= v994) != 0 {
		v1028 = v1008
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1032 = int32(1)
	if v1032 < v1003 {
		v1001 = v1028
		v1003 = v1003 - v1032
		goto L230
	} else {
		goto L241
	}
L236:
	;
	v1016 = v1008
	goto L237
L237:
	;
	v1021 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v993+v1016))))
	if base.Ui32(int32(191)) < base.Ui32(v1021) {
		v1028 = v1016
		goto L235
	} else {
		goto L239
	}
L238:
	;
	v1028 = v994
	goto L235
L239:
	;
	v1025 = v1016 - int32(1)
	if v994 < v1025 {
		v1016 = v1025
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	goto L231
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1046
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1046
	v1051 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1051 {
		goto L185
	} else {
		goto L243
	}
L243:
	;
	v2435 = v1051
	goto L1
L244:
	;
	v1058 = int32(0)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L248
L245:
	;
	if v1450 == int32(0) {
		goto L185
	} else {
		goto L307
	}
L246:
	;
	if v1188 != 0 {
		v1450 = v1058
		goto L245
	} else {
		goto L264
	}
L247:
	;
	v1188 = v1181
	goto L246
L248:
	;
	if v1071 <= v1072 {
		v1181 = int32(-1)
		goto L247
	} else {
		goto L250
	}
L249:
	;
	v1181 = int32(0)
	goto L247
L250:
	;
	v1089 = int32(1)
	v1090 = v1071 - v1089
	v1092 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1073+v1090))))
	v1094 = v1092 & int32(255)
	if base.B2i32(v1090 == v1072)|base.B2i32(int32(0) <= v1092) != 0 {
		v1152 = v1094
		v1156 = v1089
		goto L251
	} else {
		goto L252
	}
L251:
	;
	if int32(121) < v1152 {
		goto L259
	} else {
		goto L260
	}
L252:
	;
	v1101 = v1094 & int32(63)
	v1103 = v1071 - int32(2)
	v1105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+v1103))))
	v1107 = v1105 << (uint(int32(6)) % 32)
	if base.B2i32(v1103 != v1072)&base.B2i32(base.Ui32(v1105) < base.Ui32(int32(192))) == int32(0) {
		goto L253
	} else {
		goto L254
	}
L253:
	;
	v1152 = v1107&int32(1984) | v1101
	v1156 = int32(2)
	goto L251
L254:
	;
	goto L255
L255:
	;
	v1120 = v1107&int32(4032) | v1101
	v1122 = v1071 - int32(3)
	v1124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073+v1122))))
	if base.B2i32(v1122 != v1072)&base.B2i32(base.Ui32(v1124) < base.Ui32(int32(224))) == int32(0) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1152 = v1124<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1120
	v1156 = int32(3)
	goto L251
L257:
	;
	goto L258
L258:
	;
	v1142 = int32(4)
	v1144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1071+v1073-v1142))))
	v1152 = v1124<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1144&int32(7)<<(uint(int32(18))%32) | v1120
	v1156 = v1142
	goto L251
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1071 - v1156
	goto L263
L260:
	;
	v1158 = v1152 - int32(89)
	if v1158 < int32(0) {
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1158)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[1]))))
	if int32(base.Ui32(v1164)>>(uint(v1158&int32(7))%32))&int32(1) == int32(0) {
		goto L259
	} else {
		goto L262
	}
L262:
	;
	v1188 = v1156
	goto L246
L263:
	;
	goto L249
L264:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1202 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L267
L265:
	;
	if v1317 != 0 {
		v1450 = v1058
		goto L245
	} else {
		goto L288
	}
L266:
	;
	v1317 = v1310
	goto L265
L267:
	;
	if v1201 <= v1202 {
		v1310 = int32(-1)
		goto L266
	} else {
		goto L269
	}
L268:
	;
	v1310 = int32(0)
	goto L266
L269:
	;
	v1219 = int32(1)
	v1220 = v1201 - v1219
	v1222 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1203+v1220))))
	v1224 = v1222 & int32(255)
	if base.B2i32(v1220 == v1202)|base.B2i32(int32(0) <= v1222) != 0 {
		v1282 = v1224
		v1286 = v1219
		goto L270
	} else {
		goto L271
	}
L270:
	;
	if int32(121) < v1282 {
		goto L278
	} else {
		goto L279
	}
L271:
	;
	v1231 = v1224 & int32(63)
	v1233 = v1201 - int32(2)
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203+v1233))))
	v1237 = v1235 << (uint(int32(6)) % 32)
	if base.B2i32(v1233 != v1202)&base.B2i32(base.Ui32(v1235) < base.Ui32(int32(192))) == int32(0) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1282 = v1237&int32(1984) | v1231
	v1286 = int32(2)
	goto L270
L273:
	;
	goto L274
L274:
	;
	v1250 = v1237&int32(4032) | v1231
	v1252 = v1201 - int32(3)
	v1254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1203+v1252))))
	if base.B2i32(v1252 != v1202)&base.B2i32(base.Ui32(v1254) < base.Ui32(int32(224))) == int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1282 = v1254<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1250
	v1286 = int32(3)
	goto L270
L276:
	;
	goto L277
L277:
	;
	v1272 = int32(4)
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1201+v1203-v1272))))
	v1282 = v1254<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1274&int32(7)<<(uint(int32(18))%32) | v1250
	v1286 = v1272
	goto L270
L278:
	;
	v1317 = v1286
	goto L265
L279:
	;
	goto L280
L280:
	;
	v1288 = v1282 - int32(97)
	if v1288 < int32(0) {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1317 = v1286
	goto L265
L282:
	;
	goto L283
L283:
	;
	v1294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1288)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v1294)>>(uint(v1288&int32(7))%32))&int32(1) == int32(0) {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1317 = v1286
	goto L265
L285:
	;
	goto L286
L286:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1201 - v1286
	goto L287
L287:
	;
	goto L268
L288:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L291
L289:
	;
	v1450 = base.B2i32(v1447 == int32(0))
	goto L245
L290:
	;
	v1447 = v1440
	goto L289
L291:
	;
	if v1330 <= v1331 {
		v1440 = int32(-1)
		goto L290
	} else {
		goto L293
	}
L292:
	;
	v1440 = int32(0)
	goto L290
L293:
	;
	v1348 = int32(1)
	v1349 = v1330 - v1348
	v1351 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1332+v1349))))
	v1353 = v1351 & int32(255)
	if base.B2i32(v1349 == v1331)|base.B2i32(int32(0) <= v1351) != 0 {
		v1411 = v1353
		v1415 = v1348
		goto L294
	} else {
		goto L295
	}
L294:
	;
	if int32(121) < v1411 {
		goto L302
	} else {
		goto L303
	}
L295:
	;
	v1360 = v1353 & int32(63)
	v1362 = v1330 - int32(2)
	v1364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1332+v1362))))
	v1366 = v1364 << (uint(int32(6)) % 32)
	if base.B2i32(v1362 != v1331)&base.B2i32(base.Ui32(v1364) < base.Ui32(int32(192))) == int32(0) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	v1411 = v1366&int32(1984) | v1360
	v1415 = int32(2)
	goto L294
L297:
	;
	goto L298
L298:
	;
	v1379 = v1366&int32(4032) | v1360
	v1381 = v1330 - int32(3)
	v1383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1332+v1381))))
	if base.B2i32(v1381 != v1331)&base.B2i32(base.Ui32(v1383) < base.Ui32(int32(224))) == int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1411 = v1383<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1379
	v1415 = int32(3)
	goto L294
L300:
	;
	goto L301
L301:
	;
	v1401 = int32(4)
	v1403 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330+v1332-v1401))))
	v1411 = v1383<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1403&int32(7)<<(uint(int32(18))%32) | v1379
	v1415 = v1401
	goto L294
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1330 - v1415
	goto L306
L303:
	;
	v1417 = v1411 - int32(97)
	if v1417 < int32(0) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1417)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v1423)>>(uint(v1417&int32(7))%32))&int32(1) == int32(0) {
		goto L302
	} else {
		goto L305
	}
L305:
	;
	v1447 = v1415
	goto L289
L306:
	;
	goto L292
L307:
	;
	v1453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1454 = v1453 + v956
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1454
	v1458 = F_insert_s(m, l0, v1454, v1454, int32(1), int32(_a_F_porter_UTF_8_stem_12))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L5
	} else {
		goto L308
	}
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1454
	if v1458 < int32(0) {
		v2435 = v1458
		goto L1
	} else {
		goto L309
	}
L309:
	;
	goto L185
L310:
	;
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1626
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1626
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1626-int32(2) <= v1629 {
		goto L334
	} else {
		goto L335
	}
L311:
	;
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1471+v1466-int32(1)))))
	if v1475|int32(32) != int32(121) {
		goto L310
	} else {
		goto L312
	}
L312:
	;
	v1481 = v1466 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1481
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1481
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1507 = v1481
	goto L315
L313:
	;
	if v1613 < int32(0) {
		goto L310
	} else {
		goto L331
	}
L314:
	;
	v1613 = int32(-1)
	goto L313
L315:
	;
	if v1507 <= v1497 {
		goto L314
	} else {
		goto L317
	}
L317:
	;
	v1514 = int32(1)
	v1515 = v1507 - v1514
	v1517 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1498+v1515))))
	v1519 = v1517 & int32(255)
	if base.B2i32(v1515 == v1497)|base.B2i32(int32(0) <= v1517) != 0 {
		v1577 = v1519
		v1581 = v1514
		goto L318
	} else {
		goto L319
	}
L318:
	;
	if int32(121) < v1577 {
		goto L326
	} else {
		goto L327
	}
L319:
	;
	v1526 = v1519 & int32(63)
	v1528 = v1507 - int32(2)
	v1530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498+v1528))))
	v1532 = v1530 << (uint(int32(6)) % 32)
	if base.B2i32(v1528 != v1497)&base.B2i32(base.Ui32(v1530) < base.Ui32(int32(192))) == int32(0) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v1577 = v1532&int32(1984) | v1526
	v1581 = int32(2)
	goto L318
L321:
	;
	goto L322
L322:
	;
	v1545 = v1532&int32(4032) | v1526
	v1547 = v1507 - int32(3)
	v1549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1498+v1547))))
	if base.B2i32(v1547 != v1497)&base.B2i32(base.Ui32(v1549) < base.Ui32(int32(224))) == int32(0) {
		goto L323
	} else {
		goto L324
	}
L323:
	;
	v1577 = v1549<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1545
	v1581 = int32(3)
	goto L318
L324:
	;
	goto L325
L325:
	;
	v1567 = int32(4)
	v1569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1507+v1498-v1567))))
	v1577 = v1549<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1569&int32(7)<<(uint(int32(18))%32) | v1545
	v1581 = v1567
	goto L318
L326:
	;
	v1598 = v1507 - v1581
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1598
	v1507 = v1598
	goto L315
L327:
	;
	v1583 = v1577 - int32(97)
	if v1583 < int32(0) {
		goto L326
	} else {
		goto L328
	}
L328:
	;
	v1589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1583)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v1589)>>(uint(v1583&int32(7))%32))&int32(1) == int32(0) {
		goto L326
	} else {
		goto L329
	}
L329:
	;
	v1613 = v1581
	goto L313
L331:
	;
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1616 - v1613
	v1621 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_UTF_8_stem_13))
	mBase = m.M
	v1622 = m.ExcPending
	if v1622 != 0 {
		goto L5
	} else {
		goto L332
	}
L332:
	;
	if v1621 < int32(0) {
		v2435 = v1621
		goto L1
	} else {
		goto L333
	}
L333:
	;
	goto L310
L334:
	;
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1742
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1742
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1742-int32(2) <= v1745 {
		goto L379
	} else {
		goto L380
	}
L335:
	;
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1635 = int32(1)
	v1637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1633+v1626-v1635))))
	if base.B2i32(v1637&int32(224) != int32(96))|base.B2i32(v1635<<(uint(v1637)%32)&int32(_a_F_porter_UTF_8_stem_14) == int32(0)) != 0 {
		goto L334
	} else {
		goto L336
	}
L336:
	;
	v1652 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_15), int32(20), int32(0))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L5
	} else {
		goto L337
	}
L337:
	;
	if v1652 == int32(0) {
		goto L334
	} else {
		goto L338
	}
L338:
	;
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1656
	v1658 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1656 < v1658 {
		goto L334
	} else {
		goto L339
	}
L339:
	;
	switch v1652 - int32(1) {
	case 0:
		goto L352
	case 1:
		goto L351
	case 2:
		goto L350
	case 3:
		goto L349
	case 4:
		goto L348
	case 5:
		goto L347
	case 6:
		goto L346
	case 7:
		goto L345
	case 8:
		goto L344
	case 9:
		goto L343
	case 10:
		goto L342
	case 11:
		goto L341
	case 12:
		goto L340
	default:
		goto L334
	}
L340:
	;
	v1736 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_16))
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L5
	} else {
		goto L377
	}
L341:
	;
	v1730 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_17))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L5
	} else {
		goto L375
	}
L342:
	;
	v1724 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_18))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L5
	} else {
		goto L373
	}
L343:
	;
	v1718 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_19))
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L5
	} else {
		goto L371
	}
L344:
	;
	v1712 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_20))
	mBase = m.M
	v1713 = m.ExcPending
	if v1713 != 0 {
		goto L5
	} else {
		goto L369
	}
L345:
	;
	v1706 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_21))
	mBase = m.M
	v1707 = m.ExcPending
	if v1707 != 0 {
		goto L5
	} else {
		goto L367
	}
L346:
	;
	v1700 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_22))
	mBase = m.M
	v1701 = m.ExcPending
	if v1701 != 0 {
		goto L5
	} else {
		goto L365
	}
L347:
	;
	v1694 = F_slice_from_s(m, l0, int32(1), int32(_a_F_porter_UTF_8_stem_23))
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L5
	} else {
		goto L363
	}
L348:
	;
	v1688 = F_slice_from_s(m, l0, int32(3), int32(_a_F_porter_UTF_8_stem_24))
	mBase = m.M
	v1689 = m.ExcPending
	if v1689 != 0 {
		goto L5
	} else {
		goto L361
	}
L349:
	;
	v1682 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_UTF_8_stem_25))
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L5
	} else {
		goto L359
	}
L350:
	;
	v1676 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_UTF_8_stem_26))
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L5
	} else {
		goto L357
	}
L351:
	;
	v1670 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_UTF_8_stem_27))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L5
	} else {
		goto L355
	}
L352:
	;
	v1664 = F_slice_from_s(m, l0, int32(4), int32(_a_F_porter_UTF_8_stem_28))
	mBase = m.M
	v1665 = m.ExcPending
	if v1665 != 0 {
		goto L5
	} else {
		goto L353
	}
L353:
	;
	if int32(0) <= v1664 {
		goto L334
	} else {
		goto L354
	}
L354:
	;
	v2435 = v1664
	goto L1
L355:
	;
	if int32(0) <= v1670 {
		goto L334
	} else {
		goto L356
	}
L356:
	;
	v2435 = v1670
	goto L1
L357:
	;
	if int32(0) <= v1676 {
		goto L334
	} else {
		goto L358
	}
L358:
	;
	v2435 = v1676
	goto L1
L359:
	;
	if int32(0) <= v1682 {
		goto L334
	} else {
		goto L360
	}
L360:
	;
	v2435 = v1682
	goto L1
L361:
	;
	if int32(0) <= v1688 {
		goto L334
	} else {
		goto L362
	}
L362:
	;
	v2435 = v1688
	goto L1
L363:
	;
	if int32(0) <= v1694 {
		goto L334
	} else {
		goto L364
	}
L364:
	;
	v2435 = v1694
	goto L1
L365:
	;
	if int32(0) <= v1700 {
		goto L334
	} else {
		goto L366
	}
L366:
	;
	v2435 = v1700
	goto L1
L367:
	;
	if int32(0) <= v1706 {
		goto L334
	} else {
		goto L368
	}
L368:
	;
	v2435 = v1706
	goto L1
L369:
	;
	if int32(0) <= v1712 {
		goto L334
	} else {
		goto L370
	}
L370:
	;
	v2435 = v1712
	goto L1
L371:
	;
	if int32(0) <= v1718 {
		goto L334
	} else {
		goto L372
	}
L372:
	;
	v2435 = v1718
	goto L1
L373:
	;
	if int32(0) <= v1724 {
		goto L334
	} else {
		goto L374
	}
L374:
	;
	v2435 = v1724
	goto L1
L375:
	;
	if int32(0) <= v1730 {
		goto L334
	} else {
		goto L376
	}
L376:
	;
	v2435 = v1730
	goto L1
L377:
	;
	if v1736 < int32(0) {
		v2435 = v1736
		goto L1
	} else {
		goto L378
	}
L378:
	;
	goto L334
L379:
	;
	v1795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1795
	v1797 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1795
	v1801 = v1795 - int32(1)
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1801 <= v1802 {
		v1855 = v1797
		goto L393
	} else {
		goto L394
	}
L380:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1751 = int32(1)
	v1753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1749+v1742-v1751))))
	if base.B2i32(v1753&int32(224) != int32(96))|base.B2i32(v1751<<(uint(v1753)%32)&int32(_a_F_porter_UTF_8_stem_29) == int32(0)) != 0 {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v1768 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_30), int32(7), int32(0))
	mBase = m.M
	v1769 = m.ExcPending
	if v1769 != 0 {
		goto L5
	} else {
		goto L382
	}
L382:
	;
	if v1768 == int32(0) {
		goto L379
	} else {
		goto L383
	}
L383:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1772
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1772 < v1774 {
		goto L379
	} else {
		goto L384
	}
L384:
	;
	switch v1768 - int32(1) {
	case 0:
		goto L387
	case 1:
		goto L386
	case 2:
		goto L385
	default:
		goto L379
	}
L385:
	;
	v1790 = F_slice_del(m, l0)
	mBase = m.M
	if v1790 < int32(0) {
		v2435 = v1790
		goto L1
	} else {
		goto L392
	}
L386:
	;
	v1786 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_31))
	mBase = m.M
	v1787 = m.ExcPending
	if v1787 != 0 {
		goto L5
	} else {
		goto L390
	}
L387:
	;
	v1780 = F_slice_from_s(m, l0, int32(2), int32(_a_F_porter_UTF_8_stem_32))
	mBase = m.M
	v1781 = m.ExcPending
	if v1781 != 0 {
		goto L5
	} else {
		goto L388
	}
L388:
	;
	if int32(0) <= v1780 {
		goto L379
	} else {
		goto L389
	}
L389:
	;
	v2435 = v1780
	goto L1
L390:
	;
	if int32(0) <= v1786 {
		goto L379
	} else {
		goto L391
	}
L391:
	;
	v2435 = v1786
	goto L1
L392:
	;
	goto L379
L393:
	;
	if v1855 < int32(0) {
		v2435 = v1855
		goto L1
	} else {
		goto L406
	}
L394:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1804+v1801))))
	if base.B2i32(v1806&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1806)%32)&int32(_a_F_porter_UTF_8_stem_33) == int32(0)) != 0 {
		v1855 = v1797
		goto L393
	} else {
		goto L395
	}
L395:
	;
	v1821 = F_find_among_b(m, l0, int32(_a_F_porter_UTF_8_stem_34), int32(19), int32(0))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L5
	} else {
		goto L396
	}
L396:
	;
	if v1821 == int32(0) {
		v1855 = v1797
		goto L393
	} else {
		goto L397
	}
L397:
	;
	v1825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1825
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1825 < v1827 {
		v1855 = v1797
		goto L393
	} else {
		goto L398
	}
L398:
	;
	switch v1821 - int32(1) {
	case 0:
		goto L401
	case 1:
		goto L400
	default:
		goto L399
	}
L399:
	;
	v1855 = int32(1)
	goto L393
L400:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1825 <= v1834 {
		v1855 = v1797
		goto L393
	} else {
		goto L403
	}
L401:
	;
	v1831 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1831 {
		goto L399
	} else {
		goto L402
	}
L402:
	;
	v1855 = v1831
	goto L393
L403:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1838 = int32(1)
	v1840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1836+v1825-v1838))))
	if base.Ui32(v1838) < base.Ui32((v1840-int32(115))&int32(255)) {
		v1855 = v1797
		goto L393
	} else {
		goto L404
	}
L404:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1825 - int32(1)
	v1850 = F_slice_del(m, l0)
	mBase = m.M
	if v1850 < int32(0) {
		v1855 = v1850
		goto L393
	} else {
		goto L405
	}
L405:
	;
	goto L399
L406:
	;
	v1860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1860
	v1862 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1860
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1860 <= v1865 {
		v2286 = v1862
		goto L407
	} else {
		goto L408
	}
L407:
	;
	if v2286 < int32(0) {
		v2435 = v2286
		goto L1
	} else {
		goto L480
	}
L408:
	;
	v1867 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1867+v1860-int32(1)))))
	if v1871 != int32(101) {
		v2286 = v1862
		goto L407
	} else {
		goto L409
	}
L409:
	;
	v1875 = v1860 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1875
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1875
	v1878 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1860 <= v1878 {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1860 <= v1880 {
		v2286 = v1862
		goto L407
	} else {
		goto L413
	}
L411:
	;
	goto L412
L412:
	;
	v2280 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2280 {
		goto L477
	} else {
		goto L478
	}
L413:
	;
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L417
L414:
	;
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2274 + (v1875 - v1882)
	goto L412
L415:
	;
	if v2012 != 0 {
		goto L414
	} else {
		goto L433
	}
L416:
	;
	v2012 = v2005
	goto L415
L417:
	;
	if v1895 <= v1896 {
		v2005 = int32(-1)
		goto L416
	} else {
		goto L419
	}
L418:
	;
	v2005 = int32(0)
	goto L416
L419:
	;
	v1913 = int32(1)
	v1914 = v1895 - v1913
	v1916 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1897+v1914))))
	v1918 = v1916 & int32(255)
	if base.B2i32(v1914 == v1896)|base.B2i32(int32(0) <= v1916) != 0 {
		v1976 = v1918
		v1980 = v1913
		goto L420
	} else {
		goto L421
	}
L420:
	;
	if int32(121) < v1976 {
		goto L428
	} else {
		goto L429
	}
L421:
	;
	v1925 = v1918 & int32(63)
	v1927 = v1895 - int32(2)
	v1929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1897+v1927))))
	v1931 = v1929 << (uint(int32(6)) % 32)
	if base.B2i32(v1927 != v1896)&base.B2i32(base.Ui32(v1929) < base.Ui32(int32(192))) == int32(0) {
		goto L422
	} else {
		goto L423
	}
L422:
	;
	v1976 = v1931&int32(1984) | v1925
	v1980 = int32(2)
	goto L420
L423:
	;
	goto L424
L424:
	;
	v1944 = v1931&int32(4032) | v1925
	v1946 = v1895 - int32(3)
	v1948 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1897+v1946))))
	if base.B2i32(v1946 != v1896)&base.B2i32(base.Ui32(v1948) < base.Ui32(int32(224))) == int32(0) {
		goto L425
	} else {
		goto L426
	}
L425:
	;
	v1976 = v1948<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v1944
	v1980 = int32(3)
	goto L420
L426:
	;
	goto L427
L427:
	;
	v1966 = int32(4)
	v1968 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1895+v1897-v1966))))
	v1976 = v1948<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v1968&int32(7)<<(uint(int32(18))%32) | v1944
	v1980 = v1966
	goto L420
L428:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1895 - v1980
	goto L432
L429:
	;
	v1982 = v1976 - int32(89)
	if v1982 < int32(0) {
		goto L428
	} else {
		goto L430
	}
L430:
	;
	v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1982)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[1]))))
	if int32(base.Ui32(v1988)>>(uint(v1982&int32(7))%32))&int32(1) == int32(0) {
		goto L428
	} else {
		goto L431
	}
L431:
	;
	v2012 = v1980
	goto L415
L432:
	;
	goto L418
L433:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2026 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2027 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L436
L434:
	;
	if v2141 != 0 {
		goto L414
	} else {
		goto L457
	}
L435:
	;
	v2141 = v2134
	goto L434
L436:
	;
	if v2025 <= v2026 {
		v2134 = int32(-1)
		goto L435
	} else {
		goto L438
	}
L437:
	;
	v2134 = int32(0)
	goto L435
L438:
	;
	v2043 = int32(1)
	v2044 = v2025 - v2043
	v2046 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2027+v2044))))
	v2048 = v2046 & int32(255)
	if base.B2i32(v2044 == v2026)|base.B2i32(int32(0) <= v2046) != 0 {
		v2106 = v2048
		v2110 = v2043
		goto L439
	} else {
		goto L440
	}
L439:
	;
	if int32(121) < v2106 {
		goto L447
	} else {
		goto L448
	}
L440:
	;
	v2055 = v2048 & int32(63)
	v2057 = v2025 - int32(2)
	v2059 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2027+v2057))))
	v2061 = v2059 << (uint(int32(6)) % 32)
	if base.B2i32(v2057 != v2026)&base.B2i32(base.Ui32(v2059) < base.Ui32(int32(192))) == int32(0) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v2106 = v2061&int32(1984) | v2055
	v2110 = int32(2)
	goto L439
L442:
	;
	goto L443
L443:
	;
	v2074 = v2061&int32(4032) | v2055
	v2076 = v2025 - int32(3)
	v2078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2027+v2076))))
	if base.B2i32(v2076 != v2026)&base.B2i32(base.Ui32(v2078) < base.Ui32(int32(224))) == int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v2106 = v2078<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v2074
	v2110 = int32(3)
	goto L439
L445:
	;
	goto L446
L446:
	;
	v2096 = int32(4)
	v2098 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2025+v2027-v2096))))
	v2106 = v2078<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v2098&int32(7)<<(uint(int32(18))%32) | v2074
	v2110 = v2096
	goto L439
L447:
	;
	v2141 = v2110
	goto L434
L448:
	;
	goto L449
L449:
	;
	v2112 = v2106 - int32(97)
	if v2112 < int32(0) {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	v2141 = v2110
	goto L434
L451:
	;
	goto L452
L452:
	;
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2112)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v2118)>>(uint(v2112&int32(7))%32))&int32(1) == int32(0) {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v2141 = v2110
	goto L434
L454:
	;
	goto L455
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2025 - v2110
	goto L456
L456:
	;
	goto L437
L457:
	;
	v2154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L460
L458:
	;
	if v2271 == int32(0) {
		v2286 = v1862
		goto L407
	} else {
		goto L476
	}
L459:
	;
	v2271 = v2264
	goto L458
L460:
	;
	if v2154 <= v2155 {
		v2264 = int32(-1)
		goto L459
	} else {
		goto L462
	}
L461:
	;
	v2264 = int32(0)
	goto L459
L462:
	;
	v2172 = int32(1)
	v2173 = v2154 - v2172
	v2175 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2156+v2173))))
	v2177 = v2175 & int32(255)
	if base.B2i32(v2173 == v2155)|base.B2i32(int32(0) <= v2175) != 0 {
		v2235 = v2177
		v2239 = v2172
		goto L463
	} else {
		goto L464
	}
L463:
	;
	if int32(121) < v2235 {
		goto L471
	} else {
		goto L472
	}
L464:
	;
	v2184 = v2177 & int32(63)
	v2186 = v2154 - int32(2)
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2156+v2186))))
	v2190 = v2188 << (uint(int32(6)) % 32)
	if base.B2i32(v2186 != v2155)&base.B2i32(base.Ui32(v2188) < base.Ui32(int32(192))) == int32(0) {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	v2235 = v2190&int32(1984) | v2184
	v2239 = int32(2)
	goto L463
L466:
	;
	goto L467
L467:
	;
	v2203 = v2190&int32(4032) | v2184
	v2205 = v2154 - int32(3)
	v2207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2156+v2205))))
	if base.B2i32(v2205 != v2155)&base.B2i32(base.Ui32(v2207) < base.Ui32(int32(224))) == int32(0) {
		goto L468
	} else {
		goto L469
	}
L468:
	;
	v2235 = v2207<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_3) | v2203
	v2239 = int32(3)
	goto L463
L469:
	;
	goto L470
L470:
	;
	v2225 = int32(4)
	v2227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2154+v2156-v2225))))
	v2235 = v2207<<(uint(int32(12))%32)&int32(_a_F_porter_UTF_8_stem_9) | v2227&int32(7)<<(uint(int32(18))%32) | v2203
	v2239 = v2225
	goto L463
L471:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2154 - v2239
	goto L475
L472:
	;
	v2241 = v2235 - int32(97)
	if v2241 < int32(0) {
		goto L471
	} else {
		goto L473
	}
L473:
	;
	v2247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2241)>>(uint(int32(3))%32)))+uint32(_c_F_porter_UTF_8_stem[0]))))
	if int32(base.Ui32(v2247)>>(uint(v2241&int32(7))%32))&int32(1) == int32(0) {
		goto L471
	} else {
		goto L474
	}
L474:
	;
	v2271 = v2239
	goto L458
L475:
	;
	goto L461
L476:
	;
	goto L414
L477:
	;
	v2283 = int32(1)
	goto L479
L478:
	;
	v2283 = v2280
	goto L479
L479:
	;
	v2286 = v2283
	goto L407
L480:
	;
	v2289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2289
	v2291 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2289
	v2298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2289 <= v2298 {
		v2332 = v2291
		goto L482
	} else {
		goto L483
	}
L481:
	;
	if v2332 < int32(0) {
		v2435 = v2332
		goto L1
	} else {
		goto L490
	}
L482:
	;
	goto L481
L483:
	;
	v2300 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2300+v2289-int32(1)))))
	if v2304 != int32(108) {
		v2332 = v2291
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v2308 = v2289 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2308
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2308
	v2312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if base.B2i32(v2308 <= v2298)|base.B2i32(v2289 <= v2312) != 0 {
		v2332 = v2291
		goto L482
	} else {
		goto L485
	}
L485:
	;
	v2318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2308+v2300-int32(1)))))
	if v2318 != int32(108) {
		v2332 = v2291
		goto L482
	} else {
		goto L486
	}
L486:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2289 - int32(2)
	v2325 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2325 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2328 = int32(1)
	goto L489
L488:
	;
	v2328 = v2325
	goto L489
L489:
	;
	v2332 = v2328
	goto L482
L490:
	;
	v2335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2335
	if v40 != 0 {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2335
	v2435 = int32(1)
	goto L1
L492:
	;
	goto L493
L493:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2346 = v2344
	goto L495
L494:
	;
	v2435 = v2421
	goto L1
L495:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2346
	v2353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2354 != v2346 {
		goto L498
	} else {
		goto L499
	}
L496:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2346
	v2416 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2346 + v2416
	v2421 = F_slice_from_s(m, l0, v2416, int32(_a_F_porter_UTF_8_stem_35))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L5
	} else {
		goto L523
	}
L497:
	;
	goto L496
L498:
	;
	v2357 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2346+v2353))))
	if v2357 == int32(89) {
		goto L497
	} else {
		goto L501
	}
L499:
	;
	goto L500
L500:
	;
	goto L504
L501:
	;
	goto L500
L502:
	;
	if v2411 < int32(0) {
		goto L491
	} else {
		goto L522
	}
L504:
	;
	goto L505
L505:
	;
	goto L506
L506:
	;
	v2366 = v2346
	v2368 = int32(1)
	goto L509
L508:
	;
	v2411 = v2396
	goto L502
L509:
	;
	if v2354 <= v2366 {
		goto L511
	} else {
		goto L512
	}
L510:
	;
	goto L508
L511:
	;
	v2411 = int32(-1)
	goto L502
L512:
	;
	goto L513
L513:
	;
	v2373 = v2366 + int32(1)
	v2375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2353+v2366))))
	if base.Ui32(v2375) < base.Ui32(int32(192)) {
		v2396 = v2373
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v2397 = int32(1)
	if v2397 < v2368 {
		v2366 = v2396
		v2368 = v2368 - v2397
		goto L509
	} else {
		goto L521
	}
L515:
	;
	if v2354 <= v2373 {
		v2396 = v2373
		goto L514
	} else {
		goto L516
	}
L516:
	;
	v2382 = v2373
	goto L517
L517:
	;
	v2385 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2353+v2382))))
	if int32(-65) < v2385 {
		v2396 = v2382
		goto L514
	} else {
		goto L519
	}
L518:
	;
	v2396 = v2354
	goto L514
L519:
	;
	v2389 = v2382 + int32(1)
	if v2389 != v2354 {
		v2382 = v2389
		goto L517
	} else {
		goto L520
	}
L520:
	;
	goto L518
L521:
	;
	goto L510
L522:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2411
	v2346 = v2411
	goto L495
L523:
	;
	if int32(0) <= v2421 {
		goto L493
	} else {
		goto L524
	}
L524:
	;
	goto L494
}
func F_predicatelock_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_predicatelock_hash[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_get_hash_value(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		return v6 ^ v10<<(uint(int32(4))%32)
	}
}
func F_prefix_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	v14 = v11 - int32(2)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v14))|base.B2i32(int32(base.Ui32(int32(487))>>(uint(v14)%32))&int32(1) == int32(0)) != 0 {
		v31 = int32(0)
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(v14<<(uint(int32(2))%32))+uint32(_c_F_prefix_init[0])))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
		v31 = v30
	}
	if int32(32) < v31 {
		v76 = int32(-12)
		m.G0 = v8 + int32(48)
		return v76
	} else {
		v35 = v31 + int32(2)
		v38 = F_pullf_read_max(m, l2, v35, v8+int32(44), v8)
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			if v38 < int32(0) {
				v76 = v38
				m.G0 = v8 + int32(48)
				return v76
			} else {
				if v38 != v35 {
					F_px_debug(m, int32(_a_F_prefix_init_0), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v71 = int32(-100)
						base.MemoryFill(m, v8, int32(0), int32(34))
						v76 = v71
						m.G0 = v8 + int32(48)
						return v76
					}
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v8)+44))
					v51 = v50 + v31
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51-int32(2)))))
					v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
					if v54 == v55 {
						v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51-int32(1)))))
						v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
						if v60 == v61 {
							v71 = int32(0)
							base.MemoryFill(m, v8, int32(0), int32(34))
							v76 = v71
							m.G0 = v8 + int32(48)
							return v76
						} else {
							v64 = int32(0)
							F_px_debug(m, int32(_a_F_prefix_init_1), v64)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+100)) = int32(1)
								v71 = v64
								base.MemoryFill(m, v8, int32(0), int32(34))
								v76 = v71
								m.G0 = v8 + int32(48)
								return v76
							}
						}
					} else {
						v64 = int32(0)
						F_px_debug(m, int32(_a_F_prefix_init_1), v64)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+100)) = int32(1)
							v71 = v64
							base.MemoryFill(m, v8, int32(0), int32(34))
							v76 = v71
							m.G0 = v8 + int32(48)
							return v76
						}
					}
				}
			}
		}
	}
}
func F_preprocess_groupclause(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v196
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v10 <= int32(0) {
		v196 = v3
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+124))
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L5:
	;
	v14 = int32(0)
	v17 = v3
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v14<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	v28 = F_get_sortgroupref_clause(m, v26, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v196 = v32
	goto L1
L8:
	;
	return int32(0)
L9:
	;
	v32 = F_lappend(m, v17, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v35 = v14 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 < v36 {
		v14 = v35
		v17 = v32
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	v191 = F_list_copy(m, v190)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L58
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if int32(0) < v39 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L15
L15:
	;
	v183 = v9 + int32(100)
	goto L12
L16:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	if v107 == int32(0) {
		v196 = v92
		goto L1
	} else {
		goto L35
	}
L17:
	;
	v45 = v3
	v47 = v3
	goto L20
L18:
	;
	goto L19
L19:
	;
	v183 = v9 + int32(100)
	goto L12
L20:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v9)+100))
	if v50 == int32(0) {
		v92 = v45
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v92 != 0 {
		goto L16
	} else {
		goto L34
	}
L22:
	;
	goto L21
L23:
	;
	v53 = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v54 <= v53 {
		v92 = v45
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v47<<(uint(int32(2))%32))))
	v62 = v53
	goto L25
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v62<<(uint(int32(2))%32))))
	v75 = F_equal(m, v74, v61)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L8
	} else {
		goto L27
	}
L26:
	;
	v83 = F_lappend(m, v45, v74)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L32
	}
L27:
	;
	if v75 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v80 = v62 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v80 < v81 {
		v62 = v80
		goto L25
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L26
L31:
	;
	v92 = v45
	goto L22
L32:
	;
	v86 = v47 + int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v86 < v87 {
		v45 = v83
		v47 = v86
		goto L20
	} else {
		goto L33
	}
L33:
	;
	v92 = v83
	goto L22
L34:
	;
	goto L19
L35:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v110 <= int32(0) {
		v196 = v92
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v116 = int32(0)
	v119 = v92
	goto L37
L37:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v116<<(uint(int32(2))%32))))
	v129 = int32(0)
	if v119 == v129 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	v196 = v175
	goto L1
L39:
	;
	if v167 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L40:
	;
	v167 = int32(0)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
	if v135 <= int32(0) {
		v161 = v129
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v167 = v161
	goto L39
L44:
	;
	v138 = int32(0)
	if v138 < v135 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v141 = v135
	goto L47
L46:
	;
	v141 = v138
	goto L47
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v119)+12))
	v144 = int32(0)
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v142+v144<<(uint(int32(2))%32))))
	v153 = base.B2i32(v152 == v128)
	if v152 == v128 {
		v161 = v153
		goto L43
	} else {
		goto L50
	}
L49:
	;
	v161 = v153
	goto L43
L50:
	;
	v155 = v144 + int32(1)
	if v155 != v141 {
		v144 = v155
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v128)+12))
	if v170 == int32(0) {
		v183 = v9 + int32(100)
		goto L12
	} else {
		goto L55
	}
L53:
	;
	v175 = v119
	goto L54
L54:
	;
	v177 = v116 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	if v177 < v178 {
		v116 = v177
		v119 = v175
		goto L37
	} else {
		goto L57
	}
L55:
	;
	v173 = F_lappend(m, v119, v128)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v175 = v173
	goto L54
L57:
	;
	goto L38
L58:
	;
	v196 = v191
	goto L1
}
func F_preprocess_pub_all_objtype_list(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	v6 = int32(0)
	if l0 == v6 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L12
	} else {
		goto L22
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L16
	}
L3:
	;
	return
L4:
	;
	v10 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v10)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v10)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= v10 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v23 = v6
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v23<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	switch v29 {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	v46 = v23 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v46 < v47 {
		v23 = v46
		goto L6
	} else {
		goto L15
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v40 == int32(1) {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v30 == int32(1) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	v37 = F_list_concat(m, v35, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v37
	goto L8
L14:
	;
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v43)
	goto L8
L15:
	;
	goto L7
L16:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	F_errmsg(m, int32(_a_F_preprocess_pub_all_objtype_list_0), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v69 = F_errdetail(m, int32(_a_F_preprocess_pub_all_objtype_list_1), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	F_scanner_errposition(m, v71, l4)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_preprocess_pub_all_objtype_list_2), int32(_a_F_preprocess_pub_all_objtype_list_3), int32(_a_F_preprocess_pub_all_objtype_list_4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L12
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_preprocess_pub_all_objtype_list_0), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L12
	} else {
		goto L24
	}
L24:
	;
	v92 = F_errdetail(m, int32(_a_F_preprocess_pub_all_objtype_list_5), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L12
	} else {
		goto L25
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	F_scanner_errposition(m, v94, l4)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_preprocess_pub_all_objtype_list_2), int32(_a_F_preprocess_pub_all_objtype_list_6), int32(_a_F_preprocess_pub_all_objtype_list_4))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_printSubscripts(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v12 = v10
	goto L3
L2:
	;
	v12 = int32(0)
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v23 = v12
	v26 = int32(0)
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_appendStringInfoChar(m, v19, int32(91))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L4
L9:
	;
	return
L10:
	;
	if v23 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	F_get_rule_expr(m, v33, l1, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	v51 = int32(0)
	goto L13
L13:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28+v26<<(uint(int32(2))%32))))
	F_get_rule_expr(m, v56, l1, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	F_appendStringInfoChar(m, v19, int32(58))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v41 = v23 + int32(4)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	if base.Ui32(v41) < base.Ui32(v44+v45<<(uint(int32(2))%32)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v50 = v41
	goto L18
L17:
	;
	v50 = int32(0)
	goto L18
L18:
	;
	v51 = v50
	goto L13
L19:
	;
	F_appendStringInfoChar(m, v19, int32(93))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v64 = v26 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v64 < v65 {
		v23 = v51
		v26 = v64
		goto L7
	} else {
		goto L21
	}
L21:
	;
	goto L8
}
func F_printf_core(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v567 int64
	_ = v567
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v585 int64
	_ = v585
	var v587 int32
	_ = v587
	var v591 int64
	_ = v591
	var v593 int32
	_ = v593
	var v597 int64
	_ = v597
	var v599 int32
	_ = v599
	var v603 int64
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v613 float64
	_ = v613
	var v616 int32
	_ = v616
	var v620 int64
	_ = v620
	var v622 int32
	_ = v622
	var v626 int64
	_ = v626
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v636 int64
	_ = v636
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v671 int64
	_ = v671
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v678 int32
	_ = v678
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int64
	_ = v700
	var v706 int32
	_ = v706
	var v730 int64
	_ = v730
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v740 int64
	_ = v740
	var v744 int32
	_ = v744
	var v781 int64
	_ = v781
	var v785 int32
	_ = v785
	var v809 int64
	_ = v809
	var v811 int32
	_ = v811
	var v816 int32
	_ = v816
	var v819 int64
	_ = v819
	var v823 int32
	_ = v823
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int64
	_ = v855
	var v859 int64
	_ = v859
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int64
	_ = v873
	var v874 int32
	_ = v874
	var v880 int64
	_ = v880
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v887 int64
	_ = v887
	var v888 int64
	_ = v888
	var v894 int32
	_ = v894
	var v898 int64
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v954 int32
	_ = v954
	var v958 int64
	_ = v958
	var v965 int32
	_ = v965
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v995 int64
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1069 int32
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1101 int32
	_ = v1101
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1123 int32
	_ = v1123
	var v1148 int32
	_ = v1148
	var v1150 int32
	_ = v1150
	var v1155 float64
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1171 int32
	_ = v1171
	var v1195 int32
	_ = v1195
	var v1198 int32
	_ = v1198
	var v1201 int32
	_ = v1201
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1211 int64
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1217 int64
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1223 int64
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1229 int64
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1239 float64
	_ = v1239
	var v1242 int32
	_ = v1242
	var v1246 int64
	_ = v1246
	var v1248 int32
	_ = v1248
	var v1252 int64
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1258 int32
	_ = v1258
	var v1262 int64
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1278 int32
	_ = v1278
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1335 int32
	_ = v1335
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1359 int32
	_ = v1359
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1385 int32
	_ = v1385
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1467 int32
	_ = v1467
	var v1519 int32
	_ = v1519
	v6 = int32(0)
	v27 = m.G0
	v29 = v27 + int32(-64)
	m.G0 = v29
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l1
	v37 = v27 + int32(-24)
	v39 = l1
	v50 = v6
	v55 = v6
	goto L5
L1:
	;
	m.G0 = v29 - int32(-64)
	return v1519
L2:
	;
	v1519 = int32(-1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_printf_core[0])) = v1467
	goto L2
L4:
	;
	v1467 = int32(61)
	goto L3
L5:
	;
	v66 = v39
	v70 = int32(0)
	v77 = v50
	v82 = v55
	goto L7
L6:
	;
	v1519 = int32(0)
	goto L1
L7:
	;
	if v77^int32(2147483647) < v70 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v94 = v70 + v77
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v95 != 0 {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	goto L8
L11:
	;
	v1365 = v1354 - v1350
	if v1365 < v1346 {
		goto L311
	} else {
		goto L312
	}
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+39)) = uint8(v1335)
	v1346 = int32(1)
	v1349 = v649
	v1350 = v27 + int32(-25)
	v1353 = v653
	v1354 = v37
	v1359 = v654
	goto L11
L13:
	;
	v1467 = int32(28)
	goto L3
L14:
	;
	v101 = v66
	v105 = v95
	goto L17
L15:
	;
	goto L16
L16:
	;
	if l0 != 0 {
		v1519 = v94
		goto L1
	} else {
		goto L284
	}
L17:
	;
	v123 = v105 & int32(255)
	if v123 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v101 = v101 + int32(1)
	v105 = v1160
	goto L17
L20:
	;
	v190 = v169 - v66
	v192 = v94 ^ int32(2147483647)
	if v192 < v190 {
		goto L4
	} else {
		goto L31
	}
L21:
	;
	v165 = v101
	v169 = v101
	goto L20
L22:
	;
	goto L23
L23:
	;
	if v123 != int32(37) {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	v133 = v101
	v137 = v101
	goto L25
L25:
	;
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)))
	if v154 != int32(37) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v165 = v161
	v169 = v158
	goto L20
L27:
	;
	v165 = v137
	v169 = v133
	goto L20
L28:
	;
	goto L29
L29:
	;
	v158 = v133 + int32(1)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137)+2)))
	v161 = v137 + int32(2)
	if v159 == int32(37) {
		v133 = v158
		v137 = v161
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	if l0 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_out(m, l0, v66, v190)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	if v190 != 0 {
		v66 = v165
		v70 = v190
		v77 = v94
		goto L7
	} else {
		goto L37
	}
L35:
	;
	return int32(0)
L36:
	;
	goto L34
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v165
	v200 = v165 + int32(1)
	v201 = int32(-1)
	v202 = int32(*(*int8)(unsafe.Add(mBase, uint32(v165)+1)))
	v204 = v202 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v204) {
		v213 = v200
		v214 = v201
		v215 = v82
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v213
	v217 = int32(0)
	v218 = int32(*(*int8)(unsafe.Add(mBase, uint32(v213))))
	v220 = v218 - int32(32)
	if base.Ui32(int32(31)) < base.Ui32(v220) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+2)))
	if v207 != int32(36) {
		v213 = v200
		v214 = v201
		v215 = v82
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v213 = v165 + int32(3)
	v214 = v204
	v215 = int32(1)
	goto L38
L41:
	;
	if v277 == int32(42) {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	v276 = v213
	v277 = v218
	v278 = v217
	goto L41
L43:
	;
	goto L44
L44:
	;
	v224 = int32(1) << (uint(v220) % 32)
	if v224&int32(_a_F_printf_core_0) == int32(0) {
		v276 = v213
		v277 = v218
		v278 = v217
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v230 = v224
	v234 = v213
	v239 = v217
	goto L46
L46:
	;
	v256 = v234 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v256
	v258 = v230 | v239
	v259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v234)+1)))
	v260 = int32(32)
	v261 = v259 - v260
	if base.Ui32(v260) <= base.Ui32(v261) {
		v276 = v256
		v277 = v259
		v278 = v258
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v276 = v256
	v277 = v259
	v278 = v258
	goto L41
L48:
	;
	v265 = int32(1) << (uint(v261) % 32)
	if v265&int32(_a_F_printf_core_0) != 0 {
		v230 = v265
		v234 = v256
		v239 = v258
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395))))
	if v402 != int32(46) {
		goto L82
	} else {
		goto L83
	}
L51:
	;
	v296 = int32(*(*int8)(unsafe.Add(mBase, uint32(v276)+1)))
	v298 = v296 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v298) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	goto L53
L53:
	;
	v345 = v27 + int32(-4)
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v345)))
	v352 = int32(*(*int8)(unsafe.Add(mBase, uint32(v351))))
	v354 = v352 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v354) {
		goto L68
	} else {
		goto L69
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v333
	if int32(0) <= v335 {
		v395 = v333
		v397 = v278
		v398 = v335
		v399 = v336
		goto L50
	} else {
		goto L66
	}
L55:
	;
	if v215 != 0 {
		goto L13
	} else {
		goto L62
	}
L56:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+2)))
	if v301 != int32(36) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	if l0 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v333 = v276 + int32(3)
	v335 = v316
	v336 = int32(1)
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v298<<(uint(int32(2))%32)))) = int32(10)
	v316 = int32(0)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l3+v298<<(uint(int32(3))%32))))
	v316 = v315
	goto L58
L62:
	;
	v321 = v276 + int32(1)
	if l0 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v321
	v325 = int32(0)
	v395 = v321
	v397 = v278
	v398 = v325
	v399 = v325
	goto L50
L64:
	;
	goto L65
L65:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v327 + int32(4)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	v333 = v321
	v335 = v331
	v336 = int32(0)
	goto L54
L66:
	;
	v395 = v333
	v397 = v278 | int32(_a_F_printf_core_1)
	v398 = int32(0) - v335
	v399 = v336
	goto L50
L67:
	;
	if v391 < int32(0) {
		goto L4
	} else {
		goto L80
	}
L68:
	;
	v391 = int32(0)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v359 = v354
	v360 = int32(0)
	v361 = v351
	goto L71
L71:
	;
	if base.Ui32(v360) <= base.Ui32(int32(214748364)) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v391 = v376
	goto L67
L73:
	;
	v369 = v360 * int32(10)
	if base.Ui32(v369^int32(2147483647)) < base.Ui32(v359) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v376 = int32(-1)
	goto L75
L75:
	;
	v378 = v361 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v345))) = v378
	v380 = int32(*(*int8)(unsafe.Add(mBase, uint32(v361)+1)))
	v382 = v380 - int32(48)
	if base.Ui32(v382) < base.Ui32(int32(10)) {
		v359 = v382
		v360 = v376
		v361 = v378
		goto L71
	} else {
		goto L79
	}
L76:
	;
	v374 = int32(-1)
	goto L78
L77:
	;
	v374 = v359 + v369
	goto L78
L78:
	;
	v376 = v374
	goto L75
L79:
	;
	goto L72
L80:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v395 = v394
	v397 = v278
	v398 = v391
	v399 = v215
	goto L50
L81:
	;
	v506 = v501
	v510 = int32(0)
	goto L111
L82:
	;
	v501 = v395
	v502 = int32(-1)
	v504 = int32(0)
	goto L81
L83:
	;
	goto L84
L84:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+1)))
	if v406 == int32(42) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v409 = int32(*(*int8)(unsafe.Add(mBase, uint32(v395)+2)))
	v411 = v409 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v411) {
		goto L89
	} else {
		goto L90
	}
L86:
	;
	goto L87
L87:
	;
	v448 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v395 + v448
	v453 = v27 + int32(-4)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	v460 = int32(*(*int8)(unsafe.Add(mBase, uint32(v459))))
	v462 = v460 - int32(48)
	if base.Ui32(int32(9)) < base.Ui32(v462) {
		goto L99
	} else {
		goto L100
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v442
	v501 = v442
	v502 = v444
	v504 = base.B2i32(int32(0) <= v444)
	goto L81
L89:
	;
	if v399 != 0 {
		goto L13
	} else {
		goto L96
	}
L90:
	;
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v395)+3)))
	if v414 != int32(36) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	if l0 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v442 = v395 + int32(4)
	v444 = v431
	goto L88
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v411<<(uint(int32(2))%32)))) = int32(10)
	v431 = int32(0)
	goto L92
L94:
	;
	goto L95
L95:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l3+v411<<(uint(int32(3))%32))))
	v431 = v430
	goto L92
L96:
	;
	v433 = v395 + int32(2)
	v434 = int32(0)
	if l0 == v434 {
		v442 = v433
		v444 = v434
		goto L88
	} else {
		goto L97
	}
L97:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v437 + int32(4)
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	v442 = v433
	v444 = v441
	goto L88
L98:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v501 = v500
	v502 = v499
	v504 = v448
	goto L81
L99:
	;
	v499 = int32(0)
	goto L98
L100:
	;
	goto L101
L101:
	;
	v467 = v462
	v468 = int32(0)
	v469 = v459
	goto L102
L102:
	;
	if base.Ui32(v468) <= base.Ui32(int32(214748364)) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v499 = v484
	goto L98
L104:
	;
	v477 = v468 * int32(10)
	if base.Ui32(v477^int32(2147483647)) < base.Ui32(v467) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v484 = int32(-1)
	goto L106
L106:
	;
	v486 = v469 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v453))) = v486
	v488 = int32(*(*int8)(unsafe.Add(mBase, uint32(v469)+1)))
	v490 = v488 - int32(48)
	if base.Ui32(v490) < base.Ui32(int32(10)) {
		v467 = v490
		v468 = v484
		v469 = v486
		goto L102
	} else {
		goto L110
	}
L107:
	;
	v482 = int32(-1)
	goto L109
L108:
	;
	v482 = v467 + v477
	goto L109
L109:
	;
	v484 = v482
	goto L106
L110:
	;
	goto L103
L111:
	;
	v531 = int32(28)
	v532 = int32(*(*int8)(unsafe.Add(mBase, uint32(v506))))
	if base.Ui32(v532-int32(123)) < base.Ui32(int32(-58)) {
		v1467 = v531
		goto L3
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = v538
	if v544 != int32(27) {
		goto L116
	} else {
		goto L117
	}
L113:
	;
	v537 = int32(1)
	v538 = v506 + v537
	v544 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510*int32(58)+v532)+uint32(_c_F_printf_core[1]))))
	if base.Ui32((v544-v537)&int32(255)) < base.Ui32(int32(8)) {
		v506 = v538
		v510 = v544
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v645&int32(32) != 0 {
		goto L2
	} else {
		goto L141
	}
L116:
	;
	if v544 == int32(0) {
		v1467 = v531
		goto L3
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	if int32(0) <= v214 {
		v1467 = v531
		goto L3
	} else {
		goto L139
	}
L119:
	;
	if int32(0) <= v214 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	if l0 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	if l0 == int32(0) {
		goto L10
	} else {
		goto L126
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4+v214<<(uint(int32(2))%32)))) = v544
	v39 = v538
	v50 = v94
	v55 = v399
	goto L5
L124:
	;
	goto L125
L125:
	;
	v567 = *(*int64)(unsafe.Add(mBase, uint32(l3+v214<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v567
	goto L115
L126:
	;
	v572 = v27 + int32(-16)
	switch v544 - int32(9) {
	case 0:
		goto L138
	case 1, 4, 14:
		goto L130
	case 2, 5, 11, 15:
		goto L129
	case 3, 10, 12, 13:
		goto L128
	case 6:
		goto L137
	case 7:
		goto L136
	case 8:
		goto L135
	case 9:
		goto L134
	case 16:
		goto L133
	case 17:
		goto L132
	default:
		goto L131
	}
L127:
	;
	goto L115
L128:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v632 = (v628 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v632 + int32(8)
	v636 = *(*int64)(unsafe.Add(mBase, uint32(v632)))
	*(*int64)(unsafe.Add(mBase, uint32(v572))) = v636
	goto L127
L129:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v622 + int32(4)
	v626 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v622))))
	*(*int64)(unsafe.Add(mBase, uint32(v572))) = v626
	goto L127
L130:
	;
	v616 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v616 + int32(4)
	v620 = int64(*(*int32)(unsafe.Add(mBase, uint32(v616))))
	*(*int64)(unsafe.Add(mBase, uint32(v572))) = v620
	goto L127
L131:
	;
	goto L127
L132:
	;
	F_pop_arg_long_double(m, v572, l2)
	mBase = m.M
	goto L131
L133:
	;
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v609 = (v605 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v609 + int32(8)
	v613 = *(*float64)(unsafe.Add(mBase, uint32(v609)))
	*(*float64)(unsafe.Add(mBase, uint32(v572))) = v613
	goto L127
L134:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v599 + int32(4)
	v603 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	*(*int64)(unsafe.Add(mBase, uint32(v572))) = v603
	goto L127
L135:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v593 + int32(4)
	v597 = int64(*(*int8)(unsafe.Add(mBase, uint32(v593))))
	*(*int64)(unsafe.Add(mBase, uint32(v572))) = v597
	goto L127
L136:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v587 + int32(4)
	v591 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v587))))
	*(*int64)(unsafe.Add(mBase, uint32(v572))) = v591
	goto L127
L137:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v581 + int32(4)
	v585 = int64(*(*int16)(unsafe.Add(mBase, uint32(v581))))
	*(*int64)(unsafe.Add(mBase, uint32(v572))) = v585
	goto L127
L138:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v575 + int32(4)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	*(*int32)(unsafe.Add(mBase, uint32(v572))) = v579
	goto L127
L139:
	;
	v641 = int32(0)
	if l0 == v641 {
		v66 = v538
		v70 = v641
		v77 = v94
		v82 = v399
		goto L7
	} else {
		goto L140
	}
L140:
	;
	goto L115
L141:
	;
	v649 = v397 & int32(-65537)
	if v397&int32(_a_F_printf_core_1) != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v652 = v649
	goto L144
L143:
	;
	v652 = v397
	goto L144
L144:
	;
	v653 = int32(0)
	v654 = int32(_a_F_printf_core_2)
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	v656 = base.I32_extend8_s(v655)
	if v655&int32(15) == int32(3) {
		goto L162
	} else {
		goto L163
	}
L145:
	;
	if v504&base.B2i32(v502 < int32(0)) != 0 {
		goto L4
	} else {
		goto L281
	}
L146:
	;
	F_pad(m, l0, int32(32), v398, v1123, v652^int32(_a_F_printf_core_1))
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L35
	} else {
		goto L277
	}
L147:
	;
	v1022 = int32(0)
	v1026 = v1015
	goto L251
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(0)
	*(*uint32)(unsafe.Add(mBase, uint32(v29)+8)) = uint32(v995)
	v1009 = v27 + int32(-56)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+48)) = v1009
	v1012 = int32(-1)
	v1015 = v1009
	goto L147
L149:
	;
	if v502 != 0 {
		goto L247
	} else {
		goto L248
	}
L150:
	;
	v995 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if v995 != int64(0) {
		goto L148
	} else {
		goto L246
	}
L151:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	if v979 != 0 {
		goto L232
	} else {
		goto L233
	}
L152:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+48)))
	v1335 = v978
	goto L12
L153:
	;
	if v504&base.B2i32(v941 < int32(0)) != 0 {
		goto L4
	} else {
		goto L222
	}
L154:
	;
	if base.Ui64(int64(4294967296)) <= base.Ui64(v873) {
		goto L207
	} else {
		goto L208
	}
L155:
	;
	v855 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if v855 < int64(0) {
		goto L197
	} else {
		goto L198
	}
L156:
	;
	v781 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if v781 != int64(0) {
		goto L187
	} else {
		goto L188
	}
L157:
	;
	v700 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	if v700 != int64(0) {
		goto L180
	} else {
		goto L181
	}
L158:
	;
	v690 = int32(8)
	if base.Ui32(v502) <= base.Ui32(v690) {
		goto L177
	} else {
		goto L178
	}
L159:
	;
	v673 = int32(0)
	switch v510 {
	case 0:
		goto L176
	case 1:
		goto L175
	case 2:
		goto L174
	case 3:
		goto L173
	case 4:
		goto L172
	default:
		v66 = v538
		v70 = v673
		v77 = v94
		v82 = v399
		goto L7
	case 6:
		goto L171
	case 7:
		goto L170
	}
L160:
	;
	v671 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	v872 = v653
	v873 = v671
	v874 = int32(_a_F_printf_core_2)
	goto L154
L161:
	;
	switch v664 - int32(65) {
	case 0, 4, 5, 6:
		goto L145
	case 1, 3:
		v1346 = v502
		v1349 = v652
		v1350 = v66
		v1353 = v653
		v1354 = v37
		v1359 = v654
		goto L11
	case 2:
		goto L150
	default:
		goto L168
	}
L162:
	;
	v663 = v656 & int32(-45)
	goto L164
L163:
	;
	v663 = v656
	goto L164
L164:
	;
	if v510 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v664 = v663
	goto L167
L166:
	;
	v664 = v656
	goto L167
L167:
	;
	switch v664 - int32(88) {
	case 0, 32:
		v697 = v664
		v698 = v502
		v699 = v652
		goto L157
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 16, 18, 19, 20, 21, 25, 26, 28, 30, 31:
		v1346 = v502
		v1349 = v652
		v1350 = v66
		v1353 = v653
		v1354 = v37
		v1359 = v654
		goto L11
	case 9, 13, 14, 15:
		goto L145
	case 11:
		goto L152
	case 12, 17:
		goto L155
	case 22:
		goto L159
	case 23:
		goto L156
	case 24:
		goto L158
	case 27:
		goto L151
	case 29:
		goto L160
	default:
		goto L161
	}
L168:
	;
	if v664 == int32(83) {
		goto L149
	} else {
		goto L169
	}
L169:
	;
	v1346 = v502
	v1349 = v652
	v1350 = v66
	v1353 = v653
	v1354 = v37
	v1359 = v654
	goto L11
L170:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v687))) = base.I64_extend_i32_s(v94)
	v66 = v538
	v70 = v673
	v77 = v94
	v82 = v399
	goto L7
L171:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v685))) = v94
	v66 = v538
	v70 = v673
	v77 = v94
	v82 = v399
	goto L7
L172:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v683))) = uint8(v94)
	v66 = v538
	v70 = v673
	v77 = v94
	v82 = v399
	goto L7
L173:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*uint16)(unsafe.Add(mBase, uint32(v681))) = uint16(v94)
	v66 = v538
	v70 = v673
	v77 = v94
	v82 = v399
	goto L7
L174:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v678))) = base.I64_extend_i32_s(v94)
	v66 = v538
	v70 = v673
	v77 = v94
	v82 = v399
	goto L7
L175:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v676))) = v94
	v66 = v538
	v70 = v673
	v77 = v94
	v82 = v399
	goto L7
L176:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v674))) = v94
	v66 = v538
	v70 = v673
	v77 = v94
	v82 = v399
	goto L7
L177:
	;
	v693 = v690
	goto L179
L178:
	;
	v693 = v502
	goto L179
L179:
	;
	v697 = int32(120)
	v698 = v693
	v699 = v652 | int32(8)
	goto L157
L180:
	;
	v706 = v37
	v730 = v700
	goto L183
L181:
	;
	v744 = v37
	goto L182
L182:
	;
	if base.B2i32(v699&int32(8) == int32(0))|base.B2i32(v700 == int64(0)) != 0 {
		v941 = v698
		v944 = v699
		v945 = v744
		v948 = v653
		v954 = v654
		v958 = v700
		goto L153
	} else {
		goto L186
	}
L183:
	;
	v732 = v706 - int32(1)
	v736 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v730)&int32(15))+uint32(_c_F_printf_core[2]))))
	v737 = v736 | v697&int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v732))) = uint8(v737)
	v740 = int64(base.Ui64(v730) >> (uint(int64(4)) % 64))
	if v740 != int64(0) {
		v706 = v732
		v730 = v740
		goto L183
	} else {
		goto L185
	}
L184:
	;
	v744 = v732
	goto L182
L185:
	;
	goto L184
L186:
	;
	v941 = v698
	v944 = v699
	v945 = v744
	v948 = int32(2)
	v954 = int32(base.Ui32(v697)>>(uint(int32(4))%32)) + int32(_a_F_printf_core_2)
	v958 = v700
	goto L153
L187:
	;
	v785 = v37
	v809 = v781
	goto L190
L188:
	;
	v823 = v37
	goto L189
L189:
	;
	if v652&int32(8) == int32(0) {
		v941 = v502
		v944 = v652
		v945 = v823
		v948 = v653
		v954 = v654
		v958 = v781
		goto L153
	} else {
		goto L193
	}
L190:
	;
	v811 = v785 - int32(1)
	v816 = base.I32_wrap_i64(v809)&int32(7) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v811))) = uint8(v816)
	v819 = int64(base.Ui64(v809) >> (uint(int64(3)) % 64))
	if v819 != int64(0) {
		v785 = v811
		v809 = v819
		goto L190
	} else {
		goto L192
	}
L191:
	;
	v823 = v811
	goto L189
L192:
	;
	goto L191
L193:
	;
	v852 = v27 + int32(-23) - v823
	if v852 < v502 {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v854 = v502
	goto L196
L195:
	;
	v854 = v852
	goto L196
L196:
	;
	v941 = v854
	v944 = v652
	v945 = v823
	v948 = v653
	v954 = v654
	v958 = v781
	goto L153
L197:
	;
	v859 = int64(0) - v855
	*(*int64)(unsafe.Add(mBase, uint32(v29)+48)) = v859
	v872 = int32(1)
	v873 = v859
	v874 = int32(_a_F_printf_core_2)
	goto L154
L198:
	;
	goto L199
L199:
	;
	if v652&int32(2048) != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v872 = int32(1)
	v873 = v855
	v874 = int32(_a_F_printf_core_3)
	goto L154
L201:
	;
	goto L202
L202:
	;
	v870 = v652 & int32(1)
	if v870 != 0 {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v871 = int32(_a_F_printf_core_4)
	goto L205
L204:
	;
	v871 = int32(_a_F_printf_core_2)
	goto L205
L205:
	;
	v872 = v870
	v873 = v855
	v874 = v871
	goto L154
L206:
	;
	v941 = v502
	v944 = v652
	v945 = v933
	v948 = v872
	v954 = v874
	v958 = v873
	goto L153
L207:
	;
	v880 = v873
	v881 = v37
	goto L210
L208:
	;
	v898 = v873
	v899 = v37
	goto L209
L209:
	;
	v903 = base.I32_wrap_i64(v898)
	if base.Ui64(int64(10)) <= base.Ui64(v898) {
		goto L213
	} else {
		goto L214
	}
L210:
	;
	v886 = v881 - int32(1)
	v887 = int64(10)
	v888 = base.I64_div_u_s(v880, v887)
	v894 = base.I32_wrap_i64(v880-v888*v887) | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v886))) = uint8(v894)
	if base.Ui64(int64(42949672959)) < base.Ui64(v880) {
		v880 = v888
		v881 = v886
		goto L210
	} else {
		goto L212
	}
L211:
	;
	v898 = v888
	v899 = v886
	goto L209
L212:
	;
	goto L211
L213:
	;
	v907 = v899
	v908 = v903
	goto L216
L214:
	;
	v924 = v899
	v925 = v903
	goto L215
L215:
	;
	if v925 != 0 {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	v912 = v907 - int32(1)
	v913 = int32(10)
	v914 = base.I32_div_u_s(v908, v913)
	v919 = v908 - v914*v913 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v912))) = uint8(v919)
	if base.Ui32(int32(99)) < base.Ui32(v908) {
		v907 = v912
		v908 = v914
		goto L216
	} else {
		goto L218
	}
L217:
	;
	v924 = v912
	v925 = v914
	goto L215
L218:
	;
	goto L217
L219:
	;
	v929 = v924 - int32(1)
	v931 = v925 | int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v929))) = uint8(v931)
	v933 = v929
	goto L221
L220:
	;
	v933 = v924
	goto L221
L221:
	;
	goto L206
L222:
	;
	if v504 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v965 = v944 & int32(-65537)
	goto L225
L224:
	;
	v965 = v944
	goto L225
L225:
	;
	if base.B2i32(v958 != int64(0))|v941 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1346 = int32(0)
	v1349 = v965
	v1350 = v37
	v1353 = v948
	v1354 = v37
	v1359 = v954
	goto L11
L227:
	;
	goto L228
L228:
	;
	v975 = base.B2i32(v958 == int64(0)) + (v37 - v945)
	if v975 < v941 {
		goto L229
	} else {
		goto L230
	}
L229:
	;
	v977 = v941
	goto L231
L230:
	;
	v977 = v975
	goto L231
L231:
	;
	v1346 = v977
	v1349 = v965
	v1350 = v945
	v1353 = v948
	v1354 = v37
	v1359 = v954
	goto L11
L232:
	;
	v981 = v979
	goto L234
L233:
	;
	v981 = int32(_a_F_printf_core_5)
	goto L234
L234:
	;
	v982 = int32(2147483647)
	if base.Ui32(v982) <= base.Ui32(v502) {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v985 = v982
	goto L237
L236:
	;
	v985 = v502
	goto L237
L237:
	;
	v988 = F_memchr(m, v981, int32(0), v985)
	mBase = m.M
	if v988 != 0 {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v991 = v990 + v981
	if int32(0) <= v502 {
		goto L242
	} else {
		goto L243
	}
L239:
	;
	v990 = v988 - v981
	goto L241
L240:
	;
	v990 = v985
	goto L241
L241:
	;
	goto L238
L242:
	;
	v1346 = v990
	v1349 = v649
	v1350 = v981
	v1353 = v653
	v1354 = v991
	v1359 = v654
	goto L11
L243:
	;
	goto L244
L244:
	;
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v991))))
	if v994 != 0 {
		goto L4
	} else {
		goto L245
	}
L245:
	;
	v1346 = v990
	v1349 = v649
	v1350 = v981
	v1353 = v653
	v1354 = v991
	v1359 = v654
	goto L11
L246:
	;
	v1335 = int32(0)
	goto L12
L247:
	;
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v1012 = v502
	v1015 = v999
	goto L147
L248:
	;
	goto L249
L249:
	;
	v1000 = int32(0)
	F_pad(m, l0, int32(32), v398, v1000, v652)
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L35
	} else {
		goto L250
	}
L250:
	;
	v1123 = v1000
	goto L146
L251:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v1026)))
	if v1043 == int32(0) {
		v1061 = v1022
		goto L253
	} else {
		goto L254
	}
L252:
	;
	if v1061 < int32(0) {
		v1467 = int32(61)
		goto L3
	} else {
		goto L262
	}
L253:
	;
	goto L252
L254:
	;
	v1047 = v27 + int32(-60)
	if v1047 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L255:
	;
	if v1052 < int32(0) {
		goto L2
	} else {
		goto L259
	}
L256:
	;
	v1052 = int32(0)
	goto L255
L257:
	;
	goto L258
L258:
	;
	v1051 = F_wcrtomb(m, v1047, v1043)
	mBase = m.M
	v1052 = v1051
	goto L255
L259:
	;
	if base.Ui32(v1012-v1022) < base.Ui32(v1052) {
		v1061 = v1022
		goto L253
	} else {
		goto L260
	}
L260:
	;
	v1059 = v1022 + v1052
	if base.Ui32(v1059) < base.Ui32(v1012) {
		v1022 = v1059
		v1026 = v1026 + int32(4)
		goto L251
	} else {
		goto L261
	}
L261:
	;
	v1061 = v1059
	goto L253
L262:
	;
	F_pad(m, l0, int32(32), v398, v1061, v652)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L35
	} else {
		goto L263
	}
L263:
	;
	if v1061 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1123 = int32(0)
	goto L146
L265:
	;
	goto L266
L266:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v29)+48))
	v1083 = int32(0)
	v1084 = v1074
	goto L267
L267:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1084)))
	if v1101 == int32(0) {
		v1123 = v1061
		goto L146
	} else {
		goto L269
	}
L268:
	;
	v1123 = v1061
	goto L146
L269:
	;
	v1105 = v27 + int32(-60)
	if v1105 == int32(0) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1111 = v1110 + v1083
	if base.Ui32(v1061) < base.Ui32(v1111) {
		v1123 = v1061
		goto L146
	} else {
		goto L274
	}
L271:
	;
	v1110 = int32(0)
	goto L270
L272:
	;
	goto L273
L273:
	;
	v1109 = F_wcrtomb(m, v1105, v1101)
	mBase = m.M
	v1110 = v1109
	goto L270
L274:
	;
	F_out(m, l0, v1105, v1110)
	mBase = m.M
	v1114 = m.ExcPending
	if v1114 != 0 {
		goto L35
	} else {
		goto L275
	}
L275:
	;
	if base.Ui32(v1111) < base.Ui32(v1061) {
		v1083 = v1111
		v1084 = v1084 + int32(4)
		goto L267
	} else {
		goto L276
	}
L276:
	;
	goto L268
L277:
	;
	if v1123 < v398 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1150 = v398
	goto L280
L279:
	;
	v1150 = v1123
	goto L280
L280:
	;
	v66 = v538
	v70 = v1150
	v77 = v94
	v82 = v399
	goto L7
L281:
	;
	v1155 = *(*float64)(unsafe.Add(mBase, uint32(v29)+48))
	v1156 = F_fmt_fp(m, l0, v1155, v398, v502, v652, v664, v510)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L35
	} else {
		goto L282
	}
L282:
	;
	if int32(0) <= v1156 {
		v66 = v538
		v70 = v1156
		v77 = v94
		v82 = v399
		goto L7
	} else {
		goto L283
	}
L283:
	;
	v1467 = int32(61)
	goto L3
L284:
	;
	if v82 == int32(0) {
		goto L10
	} else {
		goto L285
	}
L285:
	;
	v1171 = int32(1)
	goto L286
L286:
	;
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(l4+v1171<<(uint(int32(2))%32))))
	if v1195 != 0 {
		goto L288
	} else {
		goto L289
	}
L287:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v1171) {
		goto L304
	} else {
		goto L305
	}
L288:
	;
	v1198 = l3 + v1171<<(uint(int32(3))%32)
	switch v1195 - int32(9) {
	case 0:
		goto L302
	case 1, 4, 14:
		goto L294
	case 2, 5, 11, 15:
		goto L293
	case 3, 10, 12, 13:
		goto L292
	case 6:
		goto L301
	case 7:
		goto L300
	case 8:
		goto L299
	case 9:
		goto L298
	case 16:
		goto L297
	case 17:
		goto L296
	default:
		goto L295
	}
L289:
	;
	goto L290
L290:
	;
	goto L287
L291:
	;
	v1265 = int32(1)
	v1267 = v1171 + v1265
	if v1267 != int32(10) {
		v1171 = v1267
		goto L286
	} else {
		goto L303
	}
L292:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1258 = (v1254 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1258 + int32(8)
	v1262 = *(*int64)(unsafe.Add(mBase, uint32(v1258)))
	*(*int64)(unsafe.Add(mBase, uint32(v1198))) = v1262
	goto L291
L293:
	;
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1248 + int32(4)
	v1252 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1248))))
	*(*int64)(unsafe.Add(mBase, uint32(v1198))) = v1252
	goto L291
L294:
	;
	v1242 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1242 + int32(4)
	v1246 = int64(*(*int32)(unsafe.Add(mBase, uint32(v1242))))
	*(*int64)(unsafe.Add(mBase, uint32(v1198))) = v1246
	goto L291
L295:
	;
	goto L291
L296:
	;
	F_pop_arg_long_double(m, v1198, l2)
	mBase = m.M
	goto L295
L297:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1235 = (v1231 + int32(7)) & int32(-8)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1235 + int32(8)
	v1239 = *(*float64)(unsafe.Add(mBase, uint32(v1235)))
	*(*float64)(unsafe.Add(mBase, uint32(v1198))) = v1239
	goto L291
L298:
	;
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1225 + int32(4)
	v1229 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v1225))))
	*(*int64)(unsafe.Add(mBase, uint32(v1198))) = v1229
	goto L291
L299:
	;
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1219 + int32(4)
	v1223 = int64(*(*int8)(unsafe.Add(mBase, uint32(v1219))))
	*(*int64)(unsafe.Add(mBase, uint32(v1198))) = v1223
	goto L291
L300:
	;
	v1213 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1213 + int32(4)
	v1217 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1213))))
	*(*int64)(unsafe.Add(mBase, uint32(v1198))) = v1217
	goto L291
L301:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1207 + int32(4)
	v1211 = int64(*(*int16)(unsafe.Add(mBase, uint32(v1207))))
	*(*int64)(unsafe.Add(mBase, uint32(v1198))) = v1211
	goto L291
L302:
	;
	v1201 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1201 + int32(4)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1201)))
	*(*int32)(unsafe.Add(mBase, uint32(v1198))) = v1205
	goto L291
L303:
	;
	v1519 = v1265
	goto L1
L304:
	;
	v1519 = int32(1)
	goto L1
L305:
	;
	goto L306
L306:
	;
	v1278 = v1171
	goto L307
L307:
	;
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(l4+v1278<<(uint(int32(2))%32))))
	if v1302 != 0 {
		goto L13
	} else {
		goto L309
	}
L308:
	;
	v1519 = v1303
	goto L1
L309:
	;
	v1303 = int32(1)
	v1305 = v1278 + v1303
	if v1305 != int32(10) {
		v1278 = v1305
		goto L307
	} else {
		goto L310
	}
L310:
	;
	goto L308
L311:
	;
	v1367 = v1346
	goto L313
L312:
	;
	v1367 = v1365
	goto L313
L313:
	;
	if v1353^int32(2147483647) < v1367 {
		goto L4
	} else {
		goto L314
	}
L314:
	;
	v1372 = v1367 + v1353
	if v1372 < v398 {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	v1374 = v398
	goto L317
L316:
	;
	v1374 = v1372
	goto L317
L317:
	;
	if base.Ui32(v192) < base.Ui32(v1374) {
		v1467 = int32(61)
		goto L3
	} else {
		goto L318
	}
L318:
	;
	F_pad(m, l0, int32(32), v1374, v1372, v1349)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L35
	} else {
		goto L319
	}
L319:
	;
	F_out(m, l0, v1359, v1353)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L35
	} else {
		goto L320
	}
L320:
	;
	F_pad(m, l0, int32(48), v1374, v1372, v1349^int32(_a_F_printf_core_6))
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L35
	} else {
		goto L321
	}
L321:
	;
	F_pad(m, l0, int32(48), v1367, v1365, int32(0))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L35
	} else {
		goto L322
	}
L322:
	;
	F_out(m, l0, v1350, v1365)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L35
	} else {
		goto L323
	}
L323:
	;
	F_pad(m, l0, int32(32), v1374, v1372, v1349^int32(_a_F_printf_core_1))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L35
	} else {
		goto L324
	}
L324:
	;
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v66 = v1397
	v70 = v1374
	v77 = v94
	v82 = v399
	goto L7
}
func F_printtup_shutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v3 != 0 {
		F_pfree(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			v6 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
			*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v6
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v10 != 0 {
				F_pfree(m, v10)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					if v15 != 0 {
						F_MemoryContextDelete(m, v15)
						mBase = m.M
						v17 = m.ExcPending
						if v17 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
						return
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if v15 != 0 {
					F_MemoryContextDelete(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
					return
				}
			}
		}
	} else {
		v6 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
		*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v6
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v10 != 0 {
			F_pfree(m, v10)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if v15 != 0 {
					F_MemoryContextDelete(m, v15)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
					return
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = int32(0)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			if v15 != 0 {
				F_MemoryContextDelete(m, v15)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
				return
			}
		}
	}
}
func F_process_owned_by(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L18
	} else {
		goto L63
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L18
	} else {
		goto L59
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L18
	} else {
		goto L55
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L18
	} else {
		goto L51
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L18
	} else {
		goto L46
	}
L6:
	;
	if l2 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L7:
	;
	v15 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v19 = int32(_a_F_process_owned_by_0)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_process_owned_by[0])))
	if base.B2i32(v22 == v15)|base.B2i32(v22 != v25) != 0 {
		v43 = v22
		v44 = v25
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	v70 = F_list_copy_head(m, l1, v12-int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L18
	} else {
		goto L24
	}
L10:
	;
	if v43-v44 == int32(0) {
		v121 = v15
		v123 = int32(0)
		goto L6
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v28 = v18
	v29 = v19
	goto L13
L13:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v33 == int32(0) {
		v43 = v33
		v44 = v32
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v43 = v33
	v44 = v32
	goto L11
L15:
	;
	v36 = int32(1)
	if v33 == v32 {
		v28 = v28 + v36
		v29 = v29 + v36
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_process_owned_by_1), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errhint(m, int32(_a_F_process_owned_by_2), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1616), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72+v73<<(uint(int32(2))%32)-int32(4))))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v81 = F_makeRangeVarFromNameList(m, v70)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v84 = F_relation_openrv(m, v81, int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+119)))
	v89 = v87 - int32(102)
	v94 = int32(1)
	v98 = (v89<<(uint(int32(7))%32) | int32(base.Ui32(v89&int32(254))>>(uint(v94)%32))) & int32(255)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v98))|base.B2i32(v94<<(uint(v98)%32)&int32(353) == int32(0)) != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+80))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v86)+80))
	if v109 != v110 {
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+68))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v86)+68))
	if v112 != v113 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v84)+56))
	v116 = F_get_attnum(m, v115, v80)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L18
	} else {
		goto L30
	}
L30:
	;
	if v116 == int32(0) {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v121 = v84
	v123 = v116
	goto L6
L32:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v132 = F_sequenceIsOwned(m, v126, int32(105), v8+int32(-12), v8+int32(-24))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L18
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v134 = int32(1259)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	if v132 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v139 = int32(105)
	goto L39
L38:
	;
	v139 = int32(97)
	goto L39
L39:
	;
	v140 = F_deleteDependencyRecordsForClass(m, v134, v135, v134, v139)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	if v121 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v142 = int32(1259)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v142
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v121)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v142
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v149
	F_recordDependencyOn(m, v8+int32(-24), v8+int32(-12), v139)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L18
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	m.G0 = v10 - int32(-64)
	return
L44:
	;
	F_relation_close(m, v121, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v174 + int32(4)
	F_errmsg(m, int32(_a_F_process_owned_by_5), v8+int32(-48))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	v184 = int32(*(*int8)(unsafe.Add(mBase, uint32(v183)+119)))
	F_errdetail_relkind_not_supported(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1643), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_process_owned_by_6), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1649), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L18
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_process_owned_by_7), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1653), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L18
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v84)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v231 + int32(4)
	F_errmsg(m, int32(_a_F_process_owned_by_8), v8+int32(-32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L18
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1661), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L18
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_process_owned_by_9), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	v259 = F_get_rel_name(m, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v257 + int32(4)
	v266 = F_errdetail(m, int32(_a_F_process_owned_by_10), v10)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_process_owned_by_3), int32(1678), int32(_a_F_process_owned_by_4))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pt_contained_circle(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v23 float64
	_ = v23
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v32 float64
	_ = v32
	var v43 float64
	_ = v43
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 float64
	_ = v69
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v85 float64
	_ = v85
	var v93 float64
	_ = v93
	var v98 float64
	_ = v98
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = base.F64_sub(v7, v9)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v10), v12)|base.F64_eq(base.F64_abs(v7), v12)|base.F64_eq(base.F64_abs(v9), v12) != 0 {
		v27 = v10
		v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
		v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
		v30 = base.F64_sub(v28, v29)
		v32 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
			v45 = v30
			v54 = m.G0
			v56 = v54 - int32(32)
			m.G0 = v56
			v58 = base.F64_abs(v27)
			v59 = base.F64_abs(v45)
			v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
			if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
				v63 = v58
			} else {
				v63 = v59
			}
			v64 = base.I64_reinterpret_f64(v63)
			v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
			if v66 == int64(2047) {
				v121 = v63
			} else {
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v69 = v59
				} else {
					v69 = v58
				}
				if v64 == int64(0) {
					v121 = v69
				} else {
					v72 = base.I64_reinterpret_f64(v69)
					v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
					if v74 == int64(2047) {
						v121 = v69
					} else {
						if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
							v121 = base.F64_add(v58, v59)
						} else {
							if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
								v85 = float64(1.90109156629516e-211)
								v98 = base.F64_mul(v69, v85)
								v99 = base.F64_mul(v63, v85)
								v100 = float64(5.260135901548374e+210)
							} else {
								if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
									v98 = v69
									v99 = v63
									v100 = float64(1)
								} else {
									v93 = float64(5.260135901548374e+210)
									v98 = base.F64_mul(v69, v93)
									v99 = base.F64_mul(v63, v93)
									v100 = float64(1.90109156629516e-211)
								}
							}
							F_sq(m, v56+int32(24), v56+int32(16), v98)
							mBase = m.M
							F_sq(m, v56+int32(8), v56, v99)
							mBase = m.M
							v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
							v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
							v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
							v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
						}
					}
				}
			}
			m.G0 = v56 + int32(32)
			v127 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
			return base.I64_extend_i32_u(base.F64_le(v121, v127))
		} else {
			v43 = F_float_overflow_error_ext(m, int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				v45 = v43
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				v127 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
				return base.I64_extend_i32_u(base.F64_le(v121, v127))
			}
		}
	} else {
		v23 = F_float_overflow_error_ext(m, int32(0))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = v23
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			v29 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v30 = base.F64_sub(v28, v29)
			v32 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v30), v32)|base.F64_eq(base.F64_abs(v28), v32)|base.F64_eq(base.F64_abs(v29), v32) != 0 {
				v45 = v30
				v54 = m.G0
				v56 = v54 - int32(32)
				m.G0 = v56
				v58 = base.F64_abs(v27)
				v59 = base.F64_abs(v45)
				v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
				if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
					v63 = v58
				} else {
					v63 = v59
				}
				v64 = base.I64_reinterpret_f64(v63)
				v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
				if v66 == int64(2047) {
					v121 = v63
				} else {
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v69 = v59
					} else {
						v69 = v58
					}
					if v64 == int64(0) {
						v121 = v69
					} else {
						v72 = base.I64_reinterpret_f64(v69)
						v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
						if v74 == int64(2047) {
							v121 = v69
						} else {
							if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
								v121 = base.F64_add(v58, v59)
							} else {
								if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
									v85 = float64(1.90109156629516e-211)
									v98 = base.F64_mul(v69, v85)
									v99 = base.F64_mul(v63, v85)
									v100 = float64(5.260135901548374e+210)
								} else {
									if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
										v98 = v69
										v99 = v63
										v100 = float64(1)
									} else {
										v93 = float64(5.260135901548374e+210)
										v98 = base.F64_mul(v69, v93)
										v99 = base.F64_mul(v63, v93)
										v100 = float64(1.90109156629516e-211)
									}
								}
								F_sq(m, v56+int32(24), v56+int32(16), v98)
								mBase = m.M
								F_sq(m, v56+int32(8), v56, v99)
								mBase = m.M
								v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
								v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
								v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
							}
						}
					}
				}
				m.G0 = v56 + int32(32)
				v127 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
				return base.I64_extend_i32_u(base.F64_le(v121, v127))
			} else {
				v43 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					v45 = v43
					v54 = m.G0
					v56 = v54 - int32(32)
					m.G0 = v56
					v58 = base.F64_abs(v27)
					v59 = base.F64_abs(v45)
					v62 = base.B2i32(base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)))
					if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
						v63 = v58
					} else {
						v63 = v59
					}
					v64 = base.I64_reinterpret_f64(v63)
					v66 = int64(base.Ui64(v64) >> (uint(int64(52)) % 64))
					if v66 == int64(2047) {
						v121 = v63
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v58)) < base.Ui64(base.I64_reinterpret_f64(v59)) {
							v69 = v59
						} else {
							v69 = v58
						}
						if v64 == int64(0) {
							v121 = v69
						} else {
							v72 = base.I64_reinterpret_f64(v69)
							v74 = int64(base.Ui64(v72) >> (uint(int64(52)) % 64))
							if v74 == int64(2047) {
								v121 = v69
							} else {
								if int32(65) <= base.I32_wrap_i64(v74)-base.I32_wrap_i64(v66) {
									v121 = base.F64_add(v58, v59)
								} else {
									if base.Ui64(int64(6908521828386340864)) <= base.Ui64(v72) {
										v85 = float64(1.90109156629516e-211)
										v98 = base.F64_mul(v69, v85)
										v99 = base.F64_mul(v63, v85)
										v100 = float64(5.260135901548374e+210)
									} else {
										if base.Ui64(int64(2580562586483294207)) < base.Ui64(v64) {
											v98 = v69
											v99 = v63
											v100 = float64(1)
										} else {
											v93 = float64(5.260135901548374e+210)
											v98 = base.F64_mul(v69, v93)
											v99 = base.F64_mul(v63, v93)
											v100 = float64(1.90109156629516e-211)
										}
									}
									F_sq(m, v56+int32(24), v56+int32(16), v98)
									mBase = m.M
									F_sq(m, v56+int32(8), v56, v99)
									mBase = m.M
									v109 = *(*float64)(unsafe.Add(mBase, uint32(v56)))
									v110 = *(*float64)(unsafe.Add(mBase, uint32(v56)+16))
									v112 = *(*float64)(unsafe.Add(mBase, uint32(v56)+8))
									v114 = *(*float64)(unsafe.Add(mBase, uint32(v56)+24))
									v121 = base.F64_mul(v100, base.F64_sqrt(base.F64_add(base.F64_add(base.F64_add(v109, v110), v112), v114)))
								}
							}
						}
					}
					m.G0 = v56 + int32(32)
					v127 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
					return base.I64_extend_i32_u(base.F64_le(v121, v127))
				}
			}
		}
	}
}
func F_pull_varnos(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14359(m, l0, l1, int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pull_vars_of_level(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14360(m, l0, l1, int32(949), int32(0))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_pull_vars_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v3 = int32(0)
	if l0 == v3 {
		v39 = v3
		return v39
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 != int32(321) {
			if v7 == int32(67) {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v20 + int32(1)
				v26 = F_query_tree_walker_impl(m, l0, int32(949), l1, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v30 - int32(1)
					return v26
				}
			} else {
				if v7 != int32(6) {
					v36 = F_expression_tree_walker_impl(m, l0, int32(949), l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v39 = v36
						return v39
					}
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					if v14 != v15 {
						v39 = v3
						return v39
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v42 = F_lappend(m, v41, l0)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
							return int32(0)
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v17 != v18 {
				v39 = v3
				return v39
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v42 = F_lappend(m, v41, l0)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v42
					return int32(0)
				}
			}
		}
	}
}
func F_pullf_read_max(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = l0
	goto L4
L2:
	;
	v35 = l0
	v36 = v14
	goto L3
L3:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	if l1 < v40 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	if v27 == int32(0) {
		v22 = v25
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v35 = v25
	v36 = v27
	goto L3
L6:
	;
	goto L5
L7:
	;
	m.G0 = v11 + int32(16)
	return v120
L8:
	;
	v42 = l1
	goto L10
L9:
	;
	v42 = v40
	goto L10
L10:
	;
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v42
	goto L13
L12:
	;
	v43 = l1
	goto L13
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v45 = m.T0[v36].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v38, v39, v43, l2, v44, v40)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if base.B2i32(v45 <= int32(0))|base.B2i32(l1 == v45) != 0 {
		v120 = v45
		goto L7
	} else {
		goto L16
	}
L16:
	;
	if v45 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	base.MemoryCopy(m, l3, v53, v45)
	goto L19
L18:
	;
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = l3
	v56 = l1 - v45
	if v56 <= int32(0) {
		v120 = v45
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v60 = v56
	v63 = v45
	goto L21
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v68 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v120 = v112
	goto L7
L23:
	;
	v76 = l0
	goto L26
L24:
	;
	v89 = l0
	v90 = v68
	goto L25
L25:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
	if v60 < v94 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	if v81 == int32(0) {
		v76 = v79
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v89 = v79
	v90 = v81
	goto L25
L28:
	;
	goto L27
L29:
	;
	v96 = v60
	goto L31
L30:
	;
	v96 = v94
	goto L31
L31:
	;
	if v94 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v97 = v96
	goto L34
L33:
	;
	v97 = v60
	goto L34
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v101 = m.T0[v90].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v92, v93, v97, v11+int32(12), v100, v94)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	if v101 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v63 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	goto L38
L38:
	;
	if v101 == int32(0) {
		v120 = v63
		goto L7
	} else {
		goto L43
	}
L39:
	;
	v120 = v101
	goto L7
L40:
	;
	base.MemoryFill(m, l3, int32(0), v63)
	goto L42
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	if v101 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	base.MemoryCopy(m, l3+v63, v110, v101)
	goto L46
L45:
	;
	goto L46
L46:
	;
	v112 = v101 + v63
	v113 = v60 - v101
	if int32(0) < v113 {
		v60 = v113
		v63 = v112
		goto L21
	} else {
		goto L47
	}
L47:
	;
	goto L22
}
func F_push_child_plan(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	base.MemoryCopy(m, l2, l0, int32(80))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v8 = F_lcons(m, v6, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v8
		F_set_deparse_plan(m, l0, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pushf_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		m.T0[v4].(func(*base.Module, int32))(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v8 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				if v10 != 0 {
					base.MemoryFill(m, v8, int32(0), v10)
				} else {
				}
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				F_pfree(m, v12)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					base.MemoryFill(m, l0, int32(0), int32(24))
					F_pfree(m, l0)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				base.MemoryFill(m, l0, int32(0), int32(24))
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v10 != 0 {
				base.MemoryFill(m, v8, int32(0), v10)
			} else {
			}
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pfree(m, v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				base.MemoryFill(m, l0, int32(0), int32(24))
				F_pfree(m, l0)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			base.MemoryFill(m, l0, int32(0), int32(24))
			F_pfree(m, l0)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				return
			}
		}
	}
}
