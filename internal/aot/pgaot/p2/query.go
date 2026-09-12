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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v50 int32
	_ = v50
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v122 int32
	_ = v122
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
	return v122
L2:
	;
	return int32(0)
L3:
	;
	if v14 == int32(0) {
		v122 = v2
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
	v25 = v2
	v26 = v2
	goto L8
L6:
	;
	v50 = v2
	goto L7
L7:
	;
	if v50 == int32(0) {
		v122 = v2
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v25<<(uint(int32(2))%32))))
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
	v50 = v41
	goto L7
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v10
	v41 = F_lappend(m, v26, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v44 = v25 + int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v44 < v45 {
		v25 = v44
		v26 = v41
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
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v58 < v59 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v122 = v50
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
	v98 = v58
	goto L17
L17:
	;
	if v98 == int32(0) {
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
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v70 = int32(0)
	v71 = v58
	goto L21
L21:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v67+v70<<(uint(int32(2))%32))))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v82 == int32(0) {
		goto L14
	} else {
		goto L23
	}
L22:
	;
	v98 = v92
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
	v87 = v71
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
	v92 = v71
	goto L29
L29:
	;
	v94 = v70 + int32(1)
	if v94 != v66 {
		v70 = v94
		v71 = v92
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
	*(*uint8)(unsafe.Add(mBase, uint32(v98)+24)) = uint8(v107)
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_SPI_execute(m, l0, int32(1), v2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v12 == int32(5) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = *(*int64)(unsafe.Add(mBase, _consts[461]))
	if v19 != int64(0) {
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
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v24 = v2
	v26 = int64(0)
	goto L9
L7:
	;
	v54 = v2
	goto L8
L8:
	;
	m.G0 = v8 + int32(16)
	return v54
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _consts[462]))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29+base.I32_wrap_i64(v26)<<(uint(int32(2))%32))))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v39 = F_SPI_getbinval(m, v34, v35, int32(1), v8+int32(15))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v54 = v46
	goto L8
L11:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
	if v41 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v44 = F_lappend_oid(m, v24, v39)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	v46 = v24
	goto L14
L14:
	;
	v48 = v26 + int64(1)
	v50 = *(*int64)(unsafe.Add(mBase, _consts[461]))
	if base.Ui64(v48) < base.Ui64(v50) {
		v24 = v46
		v26 = v48
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v46 = v44
	goto L14
L16:
	;
	goto L10
L17:
	;
	v65 = F_SPI_result_code_string(m, v12)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v65
	F_errmsg_internal(m, int32(181489), v8)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(497638), int32(2794), int32(74068))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
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
