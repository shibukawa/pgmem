package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cmp_list_len_asc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	v3 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
		v7 = v6
	} else {
		v7 = v3
	}
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		v10 = v9
	} else {
		v10 = v3
	}
	return base.B2i32(v10 < v7) - base.B2i32(v7 < v10)
}
func F_list_concat_copy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	if l0 == int32(0) {
		if l1 == int32(0) {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v16 = int32(8)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v19 = v17 + int32(4)
			if v19 <= v16 {
				v22 = v16
			} else {
				v22 = v19
			}
			if v22&(v22-int32(1)) != 0 {
				v29 = int32(1) << (uint(int32(32)-base.I32_clz(v22)) % 32)
			} else {
				v29 = v22
			}
			v31 = v29 - int32(4)
			v36 = F_palloc(m, v31<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v31
				*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v17
				*(*int32)(unsafe.Add(mBase, uint32(v36))) = v13
				v44 = v36 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v44
				v47 = v17 << (uint(int32(2)) % 32)
				if v47 == int32(0) {
					v144 = v36
					return v144
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					base.MemoryCopy(m, v44, v50, v47)
					return v36
				}
			}
		}
	} else {
		if l1 == int32(0) {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v58 = int32(8)
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v61 = v59 + int32(4)
			if v61 <= v58 {
				v64 = v58
			} else {
				v64 = v61
			}
			if v64&(v64-int32(1)) != 0 {
				v71 = int32(1) << (uint(int32(32)-base.I32_clz(v64)) % 32)
			} else {
				v71 = v64
			}
			v73 = v71 - int32(4)
			v78 = F_palloc(m, v73<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v73
				*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v78))) = v55
				v84 = v78 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v78)+12)) = v84
				v87 = v59 << (uint(int32(2)) % 32)
				if v87 == int32(0) {
					v144 = v78
					return v144
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					base.MemoryCopy(m, v84, v90, v87)
					return v78
				}
			}
		} else {
			v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v96 = int32(8)
			v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v99 = v97 + v98
			v101 = v99 + int32(4)
			if v101 <= v96 {
				v104 = v96
			} else {
				v104 = v101
			}
			if v104&(v104-int32(1)) != 0 {
				v111 = int32(1) << (uint(int32(32)-base.I32_clz(v104)) % 32)
			} else {
				v111 = v104
			}
			v113 = v111 - int32(4)
			v118 = F_palloc(m, v113<<(uint(int32(2))%32)+int32(16))
			mBase = m.M
			v119 = m.ExcPending
			if v119 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = v113
				*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v99
				*(*int32)(unsafe.Add(mBase, uint32(v118))) = v93
				v124 = v118 + int32(16)
				*(*int32)(unsafe.Add(mBase, uint32(v118)+12)) = v124
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v128 = v126 << (uint(int32(2)) % 32)
				if v128 != 0 {
					v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					base.MemoryCopy(m, v124, v129, v128)
				} else {
				}
				v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v133 = v131 << (uint(int32(2)) % 32)
				if v133 == int32(0) {
					v144 = v118
				} else {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					base.MemoryCopy(m, v124+v136<<(uint(int32(2))%32), v140, v133)
					v144 = v118
				}
				return v144
			}
		}
	}
}
func F_list_delete_first(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	if l0 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v5 == int32(1) {
			if l0+int32(16) != v4 {
				F_pfree(m, v4)
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return int32(0)
					} else {
						return int32(0)
					}
				}
			} else {
				F_pfree(m, l0)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					return int32(0)
				}
			}
		} else {
			v22 = v5<<(uint(int32(2))%32) - int32(4)
			if v22 != 0 {
				base.MemoryCopy(m, v4, v4+int32(4), v22)
			} else {
			}
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v26 - int32(1)
			v33 = l0
			return v33
		}
	} else {
		v33 = int32(0)
		return v33
	}
}
func F_list_make1_impl(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	v5 = F_palloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v5)+4)) = int64(17179869185)
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v5 + int32(16)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v15
		return v5
	}
}
func F_list_member(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L6
L6:
	;
	v16 = v3
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18+v16<<(uint(int32(2))%32))))
	v23 = F_equal(m, v22, l1)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	return v23
L9:
	;
	return int32(0)
L10:
	;
	if v23 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v30 = v16 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v30 < v31 {
		v16 = v30
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L8
L14:
	;
	goto L13
}
func F_list_next_fn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v4 == int32(0) {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = v4 + int32(4)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		if base.Ui32(v11) < base.Ui32(v14+v15<<(uint(int32(2))%32)) {
			v20 = v11
		} else {
			v20 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v20
		return v9
	}
}
