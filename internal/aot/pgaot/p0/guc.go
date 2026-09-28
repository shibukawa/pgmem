package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GUC_flex_fatal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	*(*int32)(unsafe.Add(mBase, _c_F_GUC_flex_fatal[0])) = l0
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_GUC_flex_fatal[1]))
	F_pgl_longjmp(m, v5)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		base.Wasm_trap_unreachable()
		for {
		}
	}
}
func F_GUC_yy_create_buffer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	v9 = F_emscripten_builtin_malloc(m, int32(48))
	mBase = m.M
	if v9 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(_a_F_GUC_yy_create_buffer_0)
		v13 = F_emscripten_builtin_malloc(m, int32(_a_F_GUC_yy_create_buffer_1))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v13
		if v13 == int32(0) {
			F_GUC_flex_fatal(m, int32(_a_F_GUC_yy_create_buffer_2))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v17 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v17
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_GUC_yy_create_buffer[0]))
			v21 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v21
			*(*uint16)(unsafe.Add(mBase, uint32(v13))) = uint16(v21)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v13
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			if v30 == v21 {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(1)
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v36 = v30 + v33<<(uint(int32(2))%32)
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
				if v37 == v9 {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v42
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v42
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v46
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v48)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v30+v54<<(uint(int32(2))%32))))
				if v9 == v58 {
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = int64(1)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_GUC_yy_create_buffer[0])) = v20
			return v9
		}
	} else {
		F_GUC_flex_fatal(m, int32(_a_F_GUC_yy_create_buffer_2))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_ReportGUCOption(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_ShowGUCOption(m, l0, int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_pfree(m, v10)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L19
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.B2i32(v15 == int32(0))|base.B2i32(v15 != v18) != 0 {
		v36 = v15
		v37 = v18
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_pq_beginmessage(m, v7, int32(83))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L15
	}
L7:
	;
	if v36-v37 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v21 = v10
	v22 = v12
	goto L10
L10:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v26 == int32(0) {
		v36 = v26
		v37 = v25
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v36 = v26
	v37 = v25
	goto L8
L12:
	;
	v29 = int32(1)
	if v26 == v25 {
		v21 = v21 + v29
		v22 = v22 + v29
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L6
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pq_sendstring(m, v7, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_pq_sendstring(m, v7, v10)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_pq_endmessage(m, v7)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	goto L3
L19:
	;
	m.G0 = v7 + int32(16)
	return
}
func F_get_guc_variables(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v80 int64
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_get_guc_variables[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v13)+808))
	if v15 != int64(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v81 = base.I32_wrap_i64(v80)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v81
	v84 = F_palloc_mul(m, int32(4), v81)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v13)+752))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v13)+728))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v13)+704))
	v21 = *(*int64)(unsafe.Add(mBase, uint32(v13)+680))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v13)+656))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v13)+632))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v13)+608))
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v13)+584))
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v13)+560))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v13)+536))
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v13)+512))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v13)+488))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v13)+464))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v13)+440))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v13)+416))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v13)+392))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v13)+368))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v13)+344))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v13)+320))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v13)+296))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v13)+272))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v13)+248))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v13)+224))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v13)+200))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v13)+176))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v13)+152))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v13)+128))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v13)+80))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v13)+56))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	v80 = v18 + (v19 + (v20 + (v21 + (v22 + (v23 + (v24 + (v25 + (v26 + (v27 + (v28 + (v29 + (v30 + (v31 + (v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + v14))))))))))))))))))))))))))))))
	goto L4
L3:
	;
	v80 = v14
	goto L4
L4:
	;
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v89 = v8 + int32(12)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_get_guc_variables[0]))
	F_hash_seq_init(m, v89, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v94 = F_hash_seq_search(m, v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v94 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v98 = v94
	v100 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pg_qsort(m, v84, v117, int32(4), int32(1855))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L16
	}
L12:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v84+v100<<(uint(int32(2))%32)))) = v104
	v110 = F_hash_seq_search(m, v8+int32(12))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	if v110 != 0 {
		v98 = v110
		v100 = v100 + int32(1)
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	m.G0 = v8 + int32(32)
	return v84
}
func F_guc_malloc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_guc_malloc[0]))
	v5 = F_MemoryContextAllocExtended(m, v3, l0, int32(2))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 != 0 {
			return v5
		} else {
			v11 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				if v11 == int32(0) {
					return v5
				} else {
					F_errcode(m, int32(_a_F_guc_malloc_0))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_guc_malloc_1), int32(0))
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_guc_malloc_2), int32(646), int32(_a_F_guc_malloc_3))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int32(0)
							} else {
								return v5
							}
						}
					}
				}
			}
		}
	}
}
func F_set_guc_source(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v5 == int32(0) {
		if l1 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
			return
		} else {
			v11 = l0 + int32(68)
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_set_guc_source[0]))
			if v13 != 0 {
				v15 = *(*int32)(unsafe.Add(mBase, _c_F_set_guc_source[1]))
				v20 = v15
			} else {
				v17 = int32(_a_F_set_guc_source_0)
				*(*int32)(unsafe.Add(mBase, _c_F_set_guc_source[0])) = v17
				v20 = v17
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v20
			v22 = int32(_a_F_set_guc_source_0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v22
			*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v11
			*(*int32)(unsafe.Add(mBase, _c_F_set_guc_source[1])) = v11
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
			return
		}
	} else {
		if l1 != 0 {
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			*(*int32)(unsafe.Add(mBase, uint32(v29))) = v31
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
		return
	}
}
