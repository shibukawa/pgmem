package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_list_concat_unique(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	v3 = int32(0)
	if l1 == v3 {
		v67 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v67
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v9 <= int32(0) {
		v67 = l0
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v12 = l0
	v16 = v3
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = v18 + v16<<(uint(int32(2))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v12 == int32(0) {
		v51 = v22
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v67 = v57
	goto L1
L6:
	;
	v64 = v16 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v64 < v65 {
		v12 = v57
		v16 = v64
		goto L4
	} else {
		goto L17
	}
L7:
	;
	v55 = F_lappend(m, v12, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L12
	} else {
		goto L16
	}
L8:
	;
	v25 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v26 <= v25 {
		v51 = v22
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v32 = v25
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v32<<(uint(int32(2))%32))))
	v40 = F_equal(m, v39, v22)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v51 = v48
	goto L7
L12:
	;
	return int32(0)
L13:
	;
	if v40 != 0 {
		v57 = v12
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v45 = v32 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v45 < v46 {
		v32 = v45
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v57 = v55
	goto L6
L17:
	;
	goto L5
}
func F_list_copy_tail(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	v3 = int32(0)
	if l0 == v3 {
		v59 = v3
		return v59
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = int32(0)
		if v10 < l1 {
			v13 = l1
		} else {
			v13 = v10
		}
		if v9 <= v13 {
			v59 = v3
			return v59
		} else {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v18 = int32(8)
			v19 = v9 - v13
			v21 = v19 + int32(4)
			if v21 <= v18 {
				v24 = v18
			} else {
				v24 = v21
			}
			if v24&(v24-int32(1)) != 0 {
				v31 = int32(1) << (uint(int32(32)-base.I32_clz(v24)) % 32)
			} else {
				v31 = v24
			}
			v33 = v31 - int32(4)
			v38 = F_palloc(m, v33<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v19
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v15
				v46 = v38 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v46
				v49 = v19 << (uint(int32(2)) % 32)
				if v49 == int32(0) {
					v59 = v38
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					base.MemoryCopy(m, v46, v52+v13<<(uint(int32(2))%32), v49)
					v59 = v38
				}
				return v59
			}
		}
	}
}
func F_list_delete_cell(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 == int32(1) {
		if l0+int32(16) != v5 {
			F_pfree(m, v5)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return int32(0)
			}
		}
	} else {
		v23 = int32(2)
		v27 = (v6 + int32(base.Ui32(l1-v5^int32(-1))>>(uint(v23)%32))) << (uint(v23) % 32)
		if v27 != 0 {
			base.MemoryCopy(m, l1, l1+int32(4), v27)
		} else {
		}
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v31 - int32(1)
		return l0
	}
}
func F_list_difference_ptr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	v3 = int32(0)
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v119
L2:
	;
	if l0 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if l0 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L5:
	;
	return int32(0)
L6:
	;
	goto L7
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 <= int32(0) {
		v119 = v3
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v19 = v3
	v20 = v3
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v20<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v29 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v119 = v65
	goto L1
L11:
	;
	v71 = v20 + int32(1)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v71 < v72 {
		v19 = v65
		v20 = v71
		goto L9
	} else {
		goto L21
	}
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v58 = F_lappend(m, v19, v28)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32+v36<<(uint(int32(2))%32))))
	if v45 == v28 {
		v65 = v19
		goto L11
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v48 = v36 + int32(1)
	if v29 != v48 {
		v36 = v48
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	return int32(0)
L20:
	;
	v65 = v58
	goto L11
L21:
	;
	goto L10
L22:
	;
	return int32(0)
L23:
	;
	goto L24
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = int32(8)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v84 = v82 + int32(4)
	if v84 <= v81 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v87 = v81
	goto L27
L26:
	;
	v87 = v84
	goto L27
L27:
	;
	if v87&(v87-int32(1)) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v94 = int32(1) << (uint(int32(32)-base.I32_clz(v87)) % 32)
	goto L30
L29:
	;
	v94 = v87
	goto L30
L30:
	;
	v96 = v94 - int32(4)
	v101 = F_palloc(m, v96<<(uint(int32(2))%32)+int32(16))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101)+8)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v101)+4)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v78
	v107 = v101 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = v107
	v110 = v82 << (uint(int32(2)) % 32)
	if v110 == int32(0) {
		v119 = v101
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	base.MemoryCopy(m, v107, v113, v110)
	return v101
}
func F_list_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	if l0 != 0 {
		v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v3 != l0+int32(16) {
			F_pfree(m, v3)
			mBase = m.M
			v8 = m.ExcPending
			if v8 != 0 {
				return
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v10 = m.ExcPending
				if v10 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			F_pfree(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F_list_make3_impl(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v6 = F_palloc(m, int32(32))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v6)+4)) = int64(17179869187)
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v6 + int32(16)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v17
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v19
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v21
		return v6
	}
}
func F_list_sort(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	if l0 == int32(0) {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v6 < int32(2) {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_pg_qsort(m, v9, v6, int32(4), l1)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_list_startup_fn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = l0
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v8
		return
	}
}
