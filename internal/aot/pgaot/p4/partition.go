package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSetupPartitionTupleRouting(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v5 = F_palloc0(m, int32(40))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l1
		v11 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
		v14 = int32(0)
		v17 = F_ExecInitPartitionDispatchInfo(m, l0, v5, v13, v14, v14, v14)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			return v5
		}
	}
}
func F_get_partition_ancestors(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(0)
	v12 = F_table_open(m, int32(2611), int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		F_get_partition_ancestors_worker(m, v12, l0, v6+int32(12))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			F_sequence_close(m, v12, int32(1))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				m.G0 = v6 + int32(16)
				return v23
			}
		}
	}
}
func F_get_partition_qual_relid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v2 = int32(0)
	v4 = F_get_rel_relispartition(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 != 0 {
			v9 = F_relation_open(m, l0, int32(1))
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v11 = F_generate_partition_qual(m, v9)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return int32(0)
				} else {
					if v11 == int32(0) {
						v24 = v2
						F_relation_close(m, v9, int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							v29 = v24
							return v29
						}
					} else {
						v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
						if int32(2) <= v15 {
							v20 = F_makeBoolExpr(m, int32(0), v11, int32(-1))
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return int32(0)
							} else {
								v24 = v20
								F_relation_close(m, v9, int32(0))
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return int32(0)
								} else {
									v29 = v24
									return v29
								}
							}
						} else {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
							v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
							v24 = v23
							F_relation_close(m, v9, int32(0))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								v29 = v24
								return v29
							}
						}
					}
				}
			}
		} else {
			v29 = v2
			return v29
		}
	}
}
func F_map_partition_varattnos(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)+52))
		v13 = F_build_attrmap_by_name(m, v10, v11, int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
			v21 = F_map_variable_attnos(m, l0, l1, v13, v18, v8+int32(15))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v24 = v21
				m.G0 = v8 + int32(16)
				return v24
			}
		}
	} else {
		v24 = int32(0)
		m.G0 = v8 + int32(16)
		return v24
	}
}
func F_partition_rbound_datum_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	v7 = int32(0)
	if l5 <= v7 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	v20 = v7
	goto L5
L4:
	;
	return v47
L5:
	;
	v24 = v20 << (uint(int32(2)) % 32)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3+v24)))
	switch v26 + int32(1) {
	case 0, 2:
		v47 = v26
		goto L4
	default:
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+v24)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2+v24)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l4+v24)))
	v38 = F_FunctionCall2Coll(m, l0+v20*int32(28), v33, v35, v37)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	if v38 != 0 {
		v47 = v38
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v43 = v20 + int32(1)
	if v43 != l5 {
		v20 = v43
		goto L5
	} else {
		goto L11
	}
L11:
	;
	goto L6
}
func F_release_partition(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v2 < v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = v2
	v10 = v5
	goto L4
L2:
	;
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+360))
	F_MemoryContextReset(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+v9*int32(56))+52))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v20 = v19
	goto L8
L7:
	;
	v20 = v10
	goto L8
L8:
	;
	v22 = v9 + int32(1)
	if v22 < v20 {
		v9 = v22
		v10 = v20
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	return
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	F_MemoryContextReset(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if int32(0) < v34 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v39 = int32(0)
	v40 = v34
	goto L16
L14:
	;
	goto L15
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v60 != 0 {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v39*int32(160))+128))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
	if v46 != v47 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	F_MemoryContextReset(m, v46)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L21
	}
L19:
	;
	v52 = v40
	goto L20
L20:
	;
	v54 = v39 + int32(1)
	if v54 < v52 {
		v39 = v54
		v40 = v52
		goto L16
	} else {
		goto L22
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v52 = v51
	goto L20
L22:
	;
	goto L17
L23:
	;
	F_tuplestore_clear(m, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L10
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v63 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+377)) = uint16(v63)
	return
L26:
	;
	goto L25
}
