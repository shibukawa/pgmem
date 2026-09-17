package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ListComparatorForWalSummaryFiles(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	return base.B2i32(base.Ui64(v8) < base.Ui64(v6)) - base.B2i32(base.Ui64(v6) < base.Ui64(v8))
}
func F_cmp_list_len_contents_asc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = v8
	goto L3
L2:
	;
	v9 = v3
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v12 = v11
	goto L6
L5:
	;
	v12 = v3
	goto L6
L6:
	;
	v15 = base.B2i32(v12 < v9) - base.B2i32(v9 < v12)
	if v12 != v9 {
		v62 = v15
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return v62
L8:
	;
	v20 = int32(0)
	goto L9
L9:
	;
	v24 = int32(0)
	if v7 == v24 {
		v34 = v24
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v62 = int32(-1)
	goto L7
L11:
	;
	if v10 == int32(0) {
		v62 = v15
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v28 <= v20 {
		v34 = int32(0)
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v34 = v30 + v20<<(uint(int32(2))%32)
	goto L11
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	if base.B2i32(v34 == int32(0))|base.B2i32(v39 <= v20) != 0 {
		v62 = v15
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v42 == int32(0) {
		v62 = v15
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v42+v20<<(uint(int32(2))%32))))
	if v49 < v45 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(1)
L18:
	;
	goto L19
L19:
	;
	if v49 <= v45 {
		v20 = v20 + int32(1)
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L10
}
func F_list_append_unique_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v34
L2:
	;
	v30 = F_lappend_oid(m, l0, l1)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v14 = v3
	goto L5
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v11+v14<<(uint(int32(2))%32))))
	if v20 == l1 {
		v34 = l0
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L2
L7:
	;
	v23 = v14 + int32(1)
	if v8 != v23 {
		v14 = v23
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(0)
L10:
	;
	v34 = v30
	goto L1
}
func F_list_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	if l0 == int32(0) {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v13 = int32(8)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = v14 + int32(4)
		if v16 <= v13 {
			v19 = v13
		} else {
			v19 = v16
		}
		if v19&(v19-int32(1)) != 0 {
			v26 = int32(1) << (uint(int32(32)-base.I32_clz(v19)) % 32)
		} else {
			v26 = v19
		}
		v28 = v26 - int32(4)
		v33 = F_palloc(m, v28<<(uint(int32(2))%32)+int32(16))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v33))) = v10
			v41 = v33 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v33)+12)) = v41
			v44 = v14 << (uint(int32(2)) % 32)
			if v44 != 0 {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				base.MemoryCopy(m, v41, v45, v44)
			} else {
			}
			return v33
		}
	}
}
func F_list_delete_first_n(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	if int32(0) < l1 {
		v7 = int32(0)
		if l0 == v7 {
			v38 = v7
			return v38
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v11 <= l1 {
				if l0+int32(16) != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							return int32(0)
						}
					}
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			} else {
				v26 = (v11 - l1) << (uint(int32(2)) % 32)
				if v26 != 0 {
					base.MemoryCopy(m, v10, v10+l1<<(uint(int32(2))%32), v26)
				} else {
				}
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v31 - l1
				v38 = l0
				return v38
			}
		}
	} else {
		v38 = l0
		return v38
	}
}
