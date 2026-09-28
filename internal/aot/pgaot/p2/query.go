package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QueryRewrite(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int64
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	v2 = int32(0)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = F_RewriteQuery(m, l0, v2, v2, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v120
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		v120 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if int32(0) < v20 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v2
	v27 = v2
	goto L8
L6:
	;
	v51 = v2
	goto L7
L7:
	;
	if v51 == int32(0) {
		v120 = v2
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v24<<(uint(int32(2))%32))))
	v38 = F_fireRIRrules(m, v36, int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L10
	}
L9:
	;
	v51 = v41
	goto L7
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v10
	v41 = F_lappend(m, v27, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v44 = v24 + int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v44 < v45 {
		v24 = v44
		v27 = v41
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v58 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v58 < v59 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v120 = v51
	goto L1
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v63 = int32(0)
	if v63 < v59 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v97 = v58
	goto L17
L17:
	;
	if v97 == int32(0) {
		goto L14
	} else {
		goto L31
	}
L18:
	;
	v66 = v59
	goto L20
L19:
	;
	v66 = v63
	goto L20
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v69 = int32(0)
	v70 = v58
	goto L21
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v67+v69<<(uint(int32(2))%32))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v82 == int32(0) {
		goto L14
	} else {
		goto L23
	}
L22:
	;
	v97 = v92
	goto L17
L23:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v85 == v62 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v87 = v81
	goto L26
L25:
	;
	v87 = v70
	goto L26
L26:
	;
	if v82&int32(-2) == int32(2) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v92 = v87
	goto L29
L28:
	;
	v92 = v70
	goto L29
L29:
	;
	v94 = v69 + int32(1)
	if v94 != v66 {
		v69 = v94
		v70 = v92
		goto L21
	} else {
		goto L30
	}
L30:
	;
	goto L22
L31:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+24)) = uint8(v107)
	goto L14
}
func F_compareQueryOperand(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = int32(12)
	v9 = int32(4095)
	v10 = v5 & v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v17 = v12 & v9
	if v10 == int32(0) {
		v23 = int32(0)
		if v23 < v17 {
			v26 = int32(-1)
		} else {
			v26 = v23
		}
		v43 = v26
	} else {
		if v17 == int32(0) {
			v43 = base.B2i32(int32(0) < v10)
		} else {
			if base.Ui32(v10) < base.Ui32(v17) {
				v32 = v10
			} else {
				v32 = v17
			}
			v33 = F_memcmp(m, l2+int32(base.Ui32(v5)>>(uint(v6)%32)), l2+int32(base.Ui32(v12)>>(uint(v6)%32)), v32)
			mBase = m.M
			if v33 != 0 {
				v41 = v33
				v43 = v41
			} else {
				if v10 == v17 {
					v43 = int32(0)
				} else {
					if v10 < v17 {
						v40 = int32(-1)
					} else {
						v40 = int32(1)
					}
					v41 = v40
					v43 = v41
				}
			}
		}
	}
	return v43
}
func F_query_or_expression_tree_mutator_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	if l0 == int32(0) {
		v15 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v15
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 != int32(67) {
			v15 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		} else {
			v10 = F_query_tree_mutator_impl(m, l0, l1, l2, int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_query_requires_rewrite_plan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != int32(6) {
		v18 = int32(1)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v10 = v8 - int32(201)
		if base.Ui32(int32(41)) < base.Ui32(v10) {
			v18 = int32(0)
		} else {
			v18 = base.I32_wrap_i64(int64(base.Ui64(int64(3298534887425)) >> (uint(base.I64_extend_i32_u(v10)) % 64)))
		}
	}
	return v18 & int32(1)
}
func F_query_to_oid_list(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = F_SPI_execute(m, l0, int32(1), v2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v13 == int32(5) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_query_to_oid_list[0]))
	if v20 != int64(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v25 = v2
	v27 = int64(0)
	goto L9
L7:
	;
	v57 = v2
	goto L8
L8:
	;
	m.G0 = v9 + int32(16)
	return v57
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_query_to_oid_list[1]))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.I32_wrap_i64(v27)<<(uint(int32(2))%32))))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	v41 = F_SPI_getbinval(m, v36, v37, int32(1), v9+int32(15))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v57 = v49
	goto L8
L11:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
	if v43 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = F_lappend_oid(m, v25, base.I32_wrap_i64(v41))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v49 = v25
	goto L14
L14:
	;
	v51 = v27 + int64(1)
	v53 = *(*int64)(unsafe.Add(mBase, _c_F_query_to_oid_list[0]))
	if base.Ui64(v51) < base.Ui64(v53) {
		v25 = v49
		v27 = v51
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v49 = v47
	goto L14
L16:
	;
	goto L10
L17:
	;
	v69 = F_SPI_result_code_string(m, v13)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v69
	F_errmsg_internal(m, int32(_a_F_query_to_oid_list_0), v9)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_query_to_oid_list_1), int32(2837), int32(_a_F_query_to_oid_list_2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_query_uses_temp_object(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v3 = int32(0)
	v8 = m.G0
	v9 = int32(16)
	v10 = v8 - v9
	m.G0 = v10
	v13 = F_palloc(m, v9)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(137438953472)
	v21 = F_palloc_mul(m, int32(12), int32(32))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v13
	v32 = F_list_make1_impl(m, int32(1), v10)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v32
	v37 = F_find_expr_references_walker(m, l0, v10+int32(8))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v39 <= int32(0) {
		v77 = v3
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	F_pfree(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L16
	}
L7:
	;
	v43 = int32(0)
	goto L8
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v53 = v50 + v43*int32(12)
	v54 = F_get_object_namespace(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v77 = v3
	goto L6
L10:
	;
	v68 = v43 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v68 < v69 {
		v43 = v68
		goto L8
	} else {
		goto L15
	}
L11:
	;
	if v54 == int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v58 = F_isAnyTempNamespace(m, v54)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v58 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v64
	v77 = int32(1)
	goto L6
L15:
	;
	goto L9
L16:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v81 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_pfree(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_pfree(m, v13)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	m.G0 = v10 + int32(16)
	return v77
}
