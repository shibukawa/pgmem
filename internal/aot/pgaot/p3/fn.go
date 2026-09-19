package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func Fn13824(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if int32(0) < l0 {
		if base.Ui32(l7) <= base.Ui32(l0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l4
					F_errmsg(m, l3, v12)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, l2, l1, int32(_a_Fn13824_0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errcode(m, int32(130))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_errmsg(m, l6, int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errfinish(m, l2, l5, int32(_a_Fn13824_0))
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
}
func Fn13828(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = l2
	v14 = v10 + int32(16)
	v16 = F_pg_snprintf(m, v14, int32(32), l4, v10)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		F_ExplainProperty(m, l0, l1, v14, int32(1), l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v10 + int32(48)
			return
		}
	}
}
func Fn13837(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	goto L2
L1:
	;
	return
L2:
	;
	v13 = int32(0)
	v14 = m.Env.Pgmem_sem(m, l4, l0, v13)
	mBase = m.M
	if v13 <= v14 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	if int32(0) <= v22 {
		goto L1
	} else {
		goto L8
	}
L5:
	;
	v22 = v14
	goto L4
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_Fn13837[0])) = int32(0) - v14
	v22 = int32(-1)
	goto L4
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_Fn13837[0]))
	if v26 == int32(27) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	goto L3
L10:
	;
	return
L11:
	;
	F_errmsg_internal(m, l3, int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_Fn13837_0), l2, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn13840(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	v8 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = F_SearchSysCache1(m, l6, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v168
L2:
	;
	return int32(0)
L3:
	;
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	F_recomputeNamespacePath(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L13
	}
L7:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v24)
	v168 = v8
	goto L1
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg_internal(m, l5, v16)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_Fn13840_0), l4, l3)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v40 = v36 + v37
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+68))
	if v41 != int32(11) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_ReleaseCatCache(m, v18)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L2
	} else {
		goto L47
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_Fn13840[0]))
	v46 = int32(0)
	if v45 == v46 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _c_Fn13840[0]))
	if v88 == int32(0) {
		v153 = v8
		goto L14
	} else {
		goto L32
	}
L18:
	;
	if v84 == int32(0) {
		v153 = v8
		goto L14
	} else {
		goto L31
	}
L19:
	;
	v84 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v52 <= int32(0) {
		v78 = v46
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v84 = v78
	goto L18
L23:
	;
	v55 = int32(0)
	if v55 < v52 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v58 = v52
	goto L26
L25:
	;
	v58 = v55
	goto L26
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v61 = int32(0)
	goto L27
L27:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v59+v61<<(uint(int32(2))%32))))
	v70 = base.B2i32(v69 == v41)
	if v69 == v41 {
		v78 = v70
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v78 = v70
	goto L22
L29:
	;
	v72 = v61 + int32(1)
	if v72 != v58 {
		v61 = v72
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L17
L32:
	;
	v91 = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v91 < v92 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v97 = v91
	goto L36
L34:
	;
	goto L35
L35:
	;
	v153 = v8
	goto L14
L36:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v97<<(uint(int32(2))%32))))
	v116 = *(*int32)(unsafe.Add(mBase, _c_Fn13840[1]))
	if v114 != v116 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	if v114 == v41 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	v125 = v97 + int32(1)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v125 < v126 {
		v97 = v125
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v153 = int32(1)
	goto L14
L42:
	;
	goto L43
L43:
	;
	v120 = int32(0)
	v122 = F_SearchSysCacheExists(m, l2, v40+int32(4), v114, v120, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	if v122 != 0 {
		v153 = v8
		goto L14
	} else {
		goto L45
	}
L45:
	;
	goto L40
L46:
	;
	goto L37
L47:
	;
	v168 = v153
	goto L1
}
func Fn13846(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	v16 = F_query_or_expression_tree_mutator_impl(m, l0, l3, v8+int32(4))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(16)
		return v16
	}
}
func Fn13853(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v4 = F_palloc0(m, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)) = uint8(v8)
		*(*uint16)(unsafe.Add(mBase, uint32(v4))) = uint16(v8)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = (v4 + int32(19)) & int32(-8)
		v18 = F_lookup_type_cache(m, l1, int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = v18
			return v4
		}
	}
}
func Fn13859(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v8 = l1 << (uint(l3) % 32)
	v10 = v8 + int32(24)
	v11 = F_palloc0(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(0)
		if base.B2i32(v8 == v15)|(base.B2i32(l0 == v15)|base.B2i32(l1 <= v15)) == v15 {
			base.MemoryCopy(m, v11+int32(24), l0, v8)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l2
		*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(1)
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v10 << (uint(int32(2)) % 32)
		return v11
	}
}
func Fn13860(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v13 = v12 - v8
	v16 = int32(0)
	if base.B2i32(base.B2i32(int64(0) < v8)^base.B2i32(v13 < v12) == v16)&base.B2i32(v13 != int64(-9223372036854775807-1)) == v16 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, l3, int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, l2, int32(107), l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
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
		v39 = v13 >> (uint(int64(63)) % 64)
		v42 = F_Int64GetDatum(m, v13^v39-v39)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			return v42
		}
	}
}
func Fn13871(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = F_string2ean(m, v8, v9, v6+int32(8), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v24 = int32(0)
			m.G0 = v6 + int32(16)
			return v24
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			v22 = F_Int64GetDatum(m, v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				m.G0 = v6 + int32(16)
				return v24
			}
		}
	}
}
func Fn13877(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v6 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v6)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+20)))
	if v16&int32(1) == v6 {
		v21 = int32(4)
		v25 = l2 + l1<<(uint(v21)%32) + v21
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
		if v26 < int32(0) {
			v68 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v74 = v68
				m.G0 = v11 + int32(16)
				return v74
			}
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v31 = v15 + v29 + v26
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+6)))
			if v32 != int32(1) {
				v74 = v31
				m.G0 = v11 + int32(16)
				return v74
			} else {
				v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+4)))
				switch v35&int32(_a_Fn13877_0) - int32(1) {
				case 0:
					v40 = int32(*(*int8)(unsafe.Add(mBase, uint32(v31))))
					v74 = v40
					m.G0 = v11 + int32(16)
					return v74
				case 1:
					v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31))))
					v74 = v41
					m.G0 = v11 + int32(16)
					return v74
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v35
						F_errmsg_internal(m, int32(_a_Fn13877_1), v11)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l4, int32(70), int32(_a_Fn13877_2))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v74 = v42
					m.G0 = v11 + int32(16)
					return v74
				}
			}
		}
	} else {
		v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+23)))
		v58 = int32(1)
		if int32(base.Ui32(v57)>>(uint(l1-v58)%32))&v58 != 0 {
			v68 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				v74 = v68
				m.G0 = v11 + int32(16)
				return v74
			}
		} else {
			v63 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v63)
			v74 = int32(0)
			m.G0 = v11 + int32(16)
			return v74
		}
	}
}
func Fn13882(m *base.Module, l0 int32, l1 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F_gbt_var_picksplit(m, v4, v5, v6, l1, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func Fn13884(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_pg_detoast_datum(m, l1)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = v10 + int32(8)
			v21 = int32(4)
			v22 = v12 + v21
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			v25 = int32(2)
			v26 = int32(base.Ui32(v24) >> (uint(v25) % 32))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			if base.Ui32(v26+v21) < base.Ui32(int32(base.Ui32(v34)>>(uint(v25)%32))) {
				v38 = v22 + (v26+int32(3))&int32(2147483644)
			} else {
				v38 = v22
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v38
			v41 = int32(4)
			v42 = v16 + v41
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v42
			v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
			v45 = int32(2)
			v46 = int32(base.Ui32(v44) >> (uint(v45) % 32))
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if base.Ui32(v46+v41) < base.Ui32(int32(base.Ui32(v54)>>(uint(v45)%32))) {
				v58 = v42 + (v46+int32(3))&int32(2147483644)
			} else {
				v58 = v42
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v58
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v63 = F_DirectFunctionCall2Coll(m, l3, int32(0), v61, v62)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				if l0 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						if l1 != v16 {
							F_pfree(m, v16)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v63
							}
						} else {
							m.G0 = v10 + int32(16)
							return v63
						}
					}
				} else {
					if l1 != v16 {
						F_pfree(m, v16)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(16)
							return v63
						}
					} else {
						m.G0 = v10 + int32(16)
						return v63
					}
				}
			}
		}
	}
}
func Fn13888(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v8)+22)) = uint16(v14)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v16 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+16)))
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v31)+12)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = F_gbt_num_consistent(m, v8+int32(12), v8+int32(24), v8+int32(22), v33&int32(1), l1, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(32)
		return v37
	}
}
func Fn13897(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = F_SearchSysCache1(m, l1, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v17 = F_pstrdup(m, v12+v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func Fn13899(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, l1, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+8))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func Fn13907(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_palloc(m, int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v26 = F_palloc(m, int32(16))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v28
			v31 = F_palloc(m, v28)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v31
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v17
				*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v16)
				v42 = F_palloc(m, int32(4))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v42
					*(*int32)(unsafe.Add(mBase, uint32(v42))) = v26
					v47 = v16 & int32(_a_Fn13907_0)
					switch v47 - int32(1) {
					case 0, 1:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = l1
						v67 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v67)
						m.G0 = v13 + int32(16)
						return v21
					case 2:
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					case 3, 4:
						v50 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v50)
						*(*int32)(unsafe.Add(mBase, uint32(v21))) = v17
						m.G0 = v13 + int32(16)
						return v21
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = v47
							F_errmsg_internal(m, int32(_a_Fn13907_1), v13)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13907_2), int32(97), int32(_a_Fn13907_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
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
func Fn13912(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v16 == int32(-1) {
			m.G0 = v9 + int32(16)
			return v12
		} else {
			v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+4)))
			if v16 == v19 {
				m.G0 = v9 + int32(16)
				return v12
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(130))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
						F_errmsg(m, int32(_a_Fn13912_0), v9)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l2, l1, int32(_a_Fn13912_1))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	}
}
func Fn13927(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v7 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v11)+13)) = v15
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(-1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v7)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v13
	v30 = int32(3)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+22)) = uint16(v30)
	v34 = F_array_recv(m, v11+int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		return int32(0)
	} else {
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
		if v38 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50462850))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, l4, int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, l3, l2, l1)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
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
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
			if v41 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50462850))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, l4, int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, l3, l2, l1)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
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
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
				if v42 != l5 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50462850))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, l4, int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, l3, l2, l1)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
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
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
					if v44 == int32(0) {
						m.G0 = v11 + int32(48)
						return v34
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50462850))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, l4, int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, l3, l2, l1)
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
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
func Fn13932(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v14
		v20 = F_get_worker(m, v10, int32(0), v7+int32(12), int32(1), l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 == int32(0) {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				v27 = int32(0)
			} else {
				v27 = v20
			}
			m.G0 = v7 + int32(16)
			return v27
		}
	}
}
func Fn13934(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v10 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return v87
L2:
	;
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
	v87 = int32(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+4)) = int64(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = F_palloc(m, int32(32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v73 = int32(0)
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v73
	v77 = F_pushJsonbValue(m, v8, l1, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L17
	}
L8:
	;
	return int32(0)
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+28)) = uint8(v36)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+29)) = uint8(v38)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v41 = v24
	v43 = v40
	goto L13
L11:
	;
	v63 = v24
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63)+24)) = int32(0)
	v73 = v24
	goto L7
L13:
	;
	v47 = F_palloc(m, int32(32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v63 = v47
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v47
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v50
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v43)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v47)+8)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v43)))
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v56
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+28)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+28)) = uint8(v58)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+29)))
	*(*uint8)(unsafe.Add(mBase, uint32(v47)+29)) = uint8(v60)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
	if v62 != 0 {
		v41 = v47
		v43 = v62
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v77
	v80 = F_JsonbValueToJsonb(m, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v87 = v80
	goto L1
}
func Fn13943(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
	v8 = int32(8)
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v11 = int32(16)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+2)))
	v15 = v7<<(uint(v8)%32) | v10<<(uint(v11)%32) | v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+2)))
	v25 = v17<<(uint(v8)%32) | v20<<(uint(v11)%32) | v24
	if base.Ui32(v15) < base.Ui32(v25) {
		v51 = l1
	} else {
		if base.Ui32(v25) < base.Ui32(v15) {
			v51 = int32(1)
		} else {
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+5)))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)))
			v31 = int32(8)
			v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+3)))
			v34 = int32(16)
			v37 = v29 | (v30<<(uint(v31)%32) | v33<<(uint(v34)%32))
			v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+5)))
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)))
			v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+3)))
			v46 = v38 | (v39<<(uint(v31)%32) | v42<<(uint(v34)%32))
			if base.Ui32(v37) < base.Ui32(v46) {
				v51 = l1
			} else {
				v51 = base.B2i32(base.Ui32(v46) < base.Ui32(v37))
			}
		}
	}
	return v51
}
func Fn13956(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = F_strlen(m, l0)
	mBase = m.M
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L19
	} else {
		goto L35
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L31
	}
L3:
	;
	if v61 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L4:
	;
	v61 = int32(0)
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v23 = v15
	v24 = l0
	v25 = v16
	v26 = v22
	goto L11
L8:
	;
	v49 = l0
	v53 = int32(0)
	goto L9
L9:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	v61 = v53 - v54
	goto L3
L10:
	;
	v49 = v44
	v53 = v46
	goto L9
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if base.B2i32(v26 != v28)|base.B2i32(v28 == int32(0)) != 0 {
		v44 = v24
		v46 = v26
		goto L10
	} else {
		goto L13
	}
L12:
	;
	v44 = v38
	v46 = int32(0)
	goto L10
L13:
	;
	v34 = v25 - int32(1)
	if v34 == int32(0) {
		v44 = v24
		v46 = v26
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v37 = int32(1)
	v38 = v24 + v37
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v39 != 0 {
		v23 = v23 + v37
		v24 = v38
		v25 = v34
		v26 = v39
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L12
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v11 + int32(-4)
	v67 = v16 + v15
	v70 = F_sscanf(m, v67, l7, v11+int32(-32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L19
	} else {
		goto L27
	}
L19:
	;
	return int32(0)
L20:
	;
	if v70 != int32(1) {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v76 = int32(10)
	v77 = F___strchrnul(m, v67, v76)
	mBase = m.M
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v79 == v76 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v83 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v83 = v77
	goto L25
L24:
	;
	v83 = int32(0)
	goto L25
L25:
	;
	goto L22
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v83 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	m.G0 = v13 - int32(-64)
	return v89
L27:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l2
	F_errmsg(m, int32(_a_Fn13956_0), v11+int32(-16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L19
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_Fn13956_1), l6, l3)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l2
	F_errmsg(m, int32(_a_Fn13956_0), v11+int32(-48))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_Fn13956_1), l5, l3)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = l2
	F_errmsg(m, int32(_a_Fn13956_0), v13)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_Fn13956_1), l4, l3)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func Fn13961(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v35 int64
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = int32(_a_Fn13961_0)
			v17 = *(*int32)(unsafe.Add(mBase, _c_Fn13961[0]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
			*(*int32)(unsafe.Add(mBase, _c_Fn13961[0])) = v19
			v21 = F_collect_corrupt_items(m, v11, l2, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v21
				*(*int32)(unsafe.Add(mBase, _c_Fn13961[0])) = v17
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				if base.Ui32(v32) < base.Ui32(v33) {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
					*(*int64)(unsafe.Add(mBase, uint32(v30))) = v35 + int64(1)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v40 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v40
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32 + v40
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					return v45 + v32*int32(6)
				} else {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = int32(2)
						v55 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
						return int32(0)
					}
				}
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
		if base.Ui32(v32) < base.Ui32(v33) {
			v35 = *(*int64)(unsafe.Add(mBase, uint32(v30)))
			*(*int64)(unsafe.Add(mBase, uint32(v30))) = v35 + int64(1)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v40 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v31))) = v32 + v40
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
			return v45 + v32*int32(6)
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = int32(2)
				v55 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
				return int32(0)
			}
		}
	}
}
func Fn13967(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v39 int32
	_ = v39
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	switch v14 - int32(98) {
	case 0:
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
		v35 = int32(0)
		v37 = F_convert_case(m, l0, l1, l2, l3, l5, v34, v35, v35)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = v37
			m.G0 = v12 + int32(16)
			return v39
		}
	case 1:
		v17 = F_strlower_libc(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v39 = v17
			m.G0 = v12 + int32(16)
			return v39
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4))))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l6
			F_errmsg_internal(m, int32(_a_Fn13967_0), v12)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_Fn13967_1), l7, l6)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func Fn13970(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	if base.Ui32(int32(127)) < base.Ui32(l0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = l4
	v18 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return base.B2i32(base.Ui32(l0-l1) < base.Ui32(int32(26)))
L4:
	;
	v23 = base.I32_div_s(v17+v18, int32(2))
	v25 = v23 << (uint(int32(3)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25+l3)))
	if base.Ui32(v27) < base.Ui32(l0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	if v39 <= v38 {
		v17 = v38
		v18 = v39
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v38 = v17
	v39 = v23 + int32(1)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25+l2)))
	if base.Ui32(v32) <= base.Ui32(l0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(1)
L11:
	;
	goto L12
L12:
	;
	v38 = v23 - int32(1)
	v39 = v18
	goto L6
L13:
	;
	goto L5
}
func Fn13976(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 float64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 float64
	_ = v146
	var v150 int64
	_ = v150
	var v155 float64
	_ = v155
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	v3 = int32(0)
	v13 = float64(0)
	v16 = m.G0
	v18 = v16 + int32(-64)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = F_palloc0(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return
	} else {
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if base.B2i32(v25 == int64(0)) == int32(0) {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v36 = v3
			v40 = v3
			v42 = v3
			for {
				v49 = v31 + v36<<(uint(int32(3))%32)
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+4)))
				if v50 == int32(1) {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
					v55 = int32(16)
					v59 = (int32(base.Ui32(v54)>>(uint(v55)%32)) ^ v54) * int32(-2048144789)
					v64 = (int32(base.Ui32(v59)>>(uint(int32(13))%32)) ^ v59) * int32(-1028477387)
					v68 = v53 & (int32(base.Ui32(v64)>>(uint(v55)%32)) ^ v64)
					v71 = v23 + v68<<(uint(int32(2))%32)
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
					*(*int32)(unsafe.Add(mBase, uint32(v71))) = v72 + int32(1)
					if base.Ui32(v36) < base.Ui32(v68) {
						v79 = base.I32_wrap_i64(v25)
					} else {
						v79 = int32(0)
					}
					v80 = v36 - v68 + v79
					if base.Ui32(v42) < base.Ui32(v80) {
						v82 = v80
					} else {
						v82 = v42
					}
					v85 = v80 + v40
					v87 = v82
				} else {
					v85 = v40
					v87 = v42
				}
				v89 = v36 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v89)) < base.Ui64(v25) {
					v36 = v89
					v40 = v85
					v42 = v87
					continue
				} else {
					break
				}
				break
			}
			v92 = int32(0)
			v97 = v92
			v99 = v92
			v101 = v92
			for {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v23+v101<<(uint(int32(2))%32))))
				v115 = v113 - int32(1)
				if base.Ui32(v99) < base.Ui32(v115) {
					v117 = v115
				} else {
					v117 = v99
				}
				if v113 != 0 {
					v118 = v117
				} else {
					v118 = v99
				}
				v122 = v113 - base.B2i32(v113 != int32(0)) + v97
				v124 = v101 + int32(1)
				if base.Ui64(base.I64_extend_i32_u(v124)) < base.Ui64(v25) {
					v97 = v122
					v99 = v118
					v101 = v124
					continue
				} else {
					break
				}
				break
			}
			v129 = v122
			v131 = v118
			v135 = v85
			v137 = v87
		} else {
			v129 = v3
			v131 = v3
			v135 = v3
			v137 = v3
		}
		F_pfree(m, v23)
		mBase = m.M
		v143 = m.ExcPending
		if v143 != 0 {
			return
		} else {
			v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v144 != 0 {
				v146 = base.F64_convert_i32_u(v144)
				v150 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				v155 = base.F64_div(base.F64_convert_i32_u(v129), v146)
				v156 = base.F64_div(base.F64_convert_i32_u(v135), v146)
				v157 = base.F64_div(v146, base.F64_convert_i64_u(v150))
			} else {
				v155 = v13
				v156 = v13
				v157 = float64(0)
			}
			v160 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v161 = m.ExcPending
			if v161 != 0 {
				return
			} else {
				if v160 != 0 {
					v162 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v18)+48)) = v155
					*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v131
					*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v129
					*(*float64)(unsafe.Add(mBase, uint32(v18)+32)) = v156
					*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v137
					*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v135
					*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v157
					*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v163
					*(*int64)(unsafe.Add(mBase, uint32(v18))) = v162
					F_errmsg_internal(m, int32(_a_Fn13976_0), v18)
					mBase = m.M
					v175 = m.ExcPending
					if v175 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_Fn13976_1), int32(1144), l1)
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return
						} else {
							m.G0 = v18 - int32(-64)
							return
						}
					}
				} else {
					m.G0 = v18 - int32(-64)
					return
				}
			}
		}
	}
}
func Fn13987(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4+int32(base.Ui32(v8)>>(uint(int32(2))%32))-v12))))
		return int32(base.Ui32(v14)>>(uint(l1)%32)) & v12
	}
}
func Fn13989(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
			if v22 != 0 {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				if v23 == v20 {
					v33 = v22
					v34 = F_range_union_internal(m, v33, v13, v18, l1)
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(16)
						return v34
					}
				} else {
					v26 = F_lookup_type_cache(m, v20, int32(2048))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+200))
						if v28 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
								F_errmsg_internal(m, int32(_a_Fn13989_0), v10)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_Fn13989_1), int32(1776), int32(_a_Fn13989_2))
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
							v33 = v26
							v34 = F_range_union_internal(m, v33, v13, v18, l1)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								m.G0 = v10 + int32(16)
								return v34
							}
						}
					}
				}
			} else {
				v26 = F_lookup_type_cache(m, v20, int32(2048))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+200))
					if v28 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v20
							F_errmsg_internal(m, int32(_a_Fn13989_0), v10)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_Fn13989_1), int32(1776), int32(_a_Fn13989_2))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v26
						v33 = v26
						v34 = F_range_union_internal(m, v33, v13, v18, l1)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							m.G0 = v10 + int32(16)
							return v34
						}
					}
				}
			}
		}
	}
}
func Fn13992(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+48)))
	if base.Ui32(v11) <= base.Ui32(int32(15)) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v16
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v15
		F_appendStringInfo(m, l0, l2, v8)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v8 + int32(16)
			return
		}
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
func Fn13998(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v14 = v9 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			v19 = int32(1)
			v20 = v18 & v19
			if v18 == v19 {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				if v26 == int32(18) {
					v29 = int32(16)
				} else {
					v29 = int32(0)
				}
				if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v36 = int32(4)
				} else {
					v36 = v29
				}
				v47 = v36
			} else {
				v37 = int32(1)
				if v20 != 0 {
					v47 = int32(base.Ui32(v18)>>(uint(v37)%32)) - v37
				} else {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v47 = int32(base.Ui32(v41)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v49 = F_RE_compile_and_cache(m, v16, l1, v48)
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				v55 = F_palloc(m, v47<<(uint(int32(2))%32)+int32(4))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v20 != 0 {
						v59 = v14
					} else {
						v59 = v9 + int32(4)
					}
					v60 = F_pg_mb2wchar_with_len(m, v59, v55, v47)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v62 = int32(0)
						v65 = F_RE_wchar_execute(m, v55, v60, v62, v62, v62)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v55)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								return v65 ^ int32(1)
							}
						}
					}
				}
			}
		}
	}
}
func Fn14003(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_ArrayGetIntegerTypmods(m, v9, v6+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v17 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_Fn14003_0), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_Fn14003_1), int32(65), int32(_a_Fn14003_2))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
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
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v37 = F_anytime_typmod_check(m, l1, v36)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 + int32(16)
					return v37
				}
			}
		}
	}
}
func Fn14016(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v23, v24, v25, int32(6), int32(-1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		return int32(0)
	} else {
		v32 = v24 - l8
		if base.B2i32(base.Ui32(v32) <= base.Ui32(l7))&(int32(base.Ui32(l6)>>(uint(v32)%32))&int32(1)) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(2600))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v24
					F_errmsg(m, l5, v18)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, l4, l3, l2)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
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
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v32<<(uint(int32(2))%32)+l1)))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
			v57 = int32(0)
			v62 = F_UtfToLocal(m, v22, v25, v21, v56, v57, v57, v57, v24, base.B2i32(v20 != v57))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return int32(0)
			} else {
				m.G0 = v18 + int32(16)
				return v62
			}
		}
	}
}
func Fn14021(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 float32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v17 = v10 + int32(1)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum_packed(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
			v23 = v21 & int32(1)
			if v23 != 0 {
				v24 = v17
			} else {
				v24 = v10 + int32(4)
			}
			if v21 == int32(1) {
				v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
				if v30 == int32(18) {
					v33 = int32(16)
				} else {
					v33 = int32(0)
				}
				if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v40 = int32(4)
				} else {
					v40 = v33
				}
				v51 = v40
			} else {
				v41 = int32(1)
				if v23 != 0 {
					v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v52 = int32(1)
			v53 = v19 + v52
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
			v58 = v56 & v52
			if v58 != 0 {
				v59 = v53
			} else {
				v59 = v19 + int32(4)
			}
			if v56 == int32(1) {
				v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v65 == int32(18) {
					v68 = int32(16)
				} else {
					v68 = int32(0)
				}
				if base.Ui32((v65-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v75 = int32(4)
				} else {
					v75 = v68
				}
				v86 = v75
			} else {
				v76 = int32(1)
				if v58 != 0 {
					v86 = int32(base.Ui32(v56)>>(uint(v76)%32)) - v76
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v87 = F_calc_word_similarity(m, v24, v51, v59, v86, l1)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v89 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v93 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								return base.I32_reinterpret_f32(v87)
							}
						} else {
							return base.I32_reinterpret_f32(v87)
						}
					}
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v93 != v19 {
						F_pfree(m, v19)
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							return base.I32_reinterpret_f32(v87)
						}
					} else {
						return base.I32_reinterpret_f32(v87)
					}
				}
			}
		}
	}
}
