package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__jumbleA_Expr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	v4 = l1 + int32(4)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = v13
	} else {
		v14 = int32(4)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v16-int32(1021)) < base.Ui32(v14) {
			v25 = v16
			v27 = v14
			v29 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v25) {
					v34 = F_hash_bytes_extended(m, v15, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v15))) = v34
					v37 = int32(8)
				} else {
					v37 = v25
				}
				v39 = int32(1024) - v37
				if base.Ui32(v27) < base.Ui32(v39) {
					v41 = v27
				} else {
					v41 = v39
				}
				if v41 != 0 {
					base.MemoryCopy(m, v37+v15, v29, v41)
				} else {
				}
				v45 = v37 + v41
				v46 = v27 - v41
				if v46 != 0 {
					v25 = v45
					v27 = v46
					v29 = v41 + v29
					continue
				} else {
					break
				}
				break
			}
			v54 = v45
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16+v15))) = v10
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v54 = v49 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
		v64 = v54
	}
	v69 = int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v64-int32(1021)) < base.Ui32(v69) {
		v76 = v4
		v77 = v64
		v79 = v69
		for {
			if base.Ui32(int32(1024)) <= base.Ui32(v77) {
				v86 = F_hash_bytes_extended(m, v70, int32(1024), int64(0))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v70))) = v86
				v89 = int32(8)
			} else {
				v89 = v77
			}
			v91 = int32(1024) - v89
			if base.Ui32(v79) < base.Ui32(v91) {
				v93 = v79
			} else {
				v93 = v91
			}
			if v93 != 0 {
				base.MemoryCopy(m, v89+v70, v76, v93)
			} else {
			}
			v97 = v89 + v93
			v98 = v79 - v93
			if v98 != 0 {
				v76 = v76 + v93
				v77 = v97
				v79 = v98
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v97
	} else {
		v101 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		*(*int32)(unsafe.Add(mBase, uint32(v64+v70))) = v101
		v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v103 + int32(4)
	}
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		return
	} else {
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		F__jumbleNode(m, l0, v117)
		mBase = m.M
		v119 = m.ExcPending
		if v119 != 0 {
			return
		} else {
			v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			F__jumbleNode(m, l0, v120)
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F__jumbleA_Indices(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v4 = l1 + int32(4)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v10 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = v13
	} else {
		v14 = int32(4)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if base.Ui32(v16-int32(1021)) < base.Ui32(v14) {
			v25 = v16
			v26 = v14
			v29 = l0 + int32(28)
			for {
				if base.Ui32(int32(1024)) <= base.Ui32(v25) {
					v34 = F_hash_bytes_extended(m, v15, int32(1024), int64(0))
					mBase = m.M
					*(*int64)(unsafe.Add(mBase, uint32(v15))) = v34
					v37 = int32(8)
				} else {
					v37 = v25
				}
				v39 = int32(1024) - v37
				if base.Ui32(v26) < base.Ui32(v39) {
					v41 = v26
				} else {
					v41 = v39
				}
				if v41 != 0 {
					base.MemoryCopy(m, v37+v15, v29, v41)
				} else {
				}
				v45 = v37 + v41
				v46 = v26 - v41
				if v46 != 0 {
					v25 = v45
					v26 = v46
					v29 = v41 + v29
					continue
				} else {
					break
				}
				break
			}
			v54 = v45
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v16+v15))) = v10
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v54 = v49 + int32(4)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
		v64 = v54
	}
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v64 != int32(1024) {
		v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		*(*uint8)(unsafe.Add(mBase, uint32(v64+v69))) = uint8(v73)
		v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v75 + int32(1)
	} else {
		v81 = F_hash_bytes_extended(m, v69, int32(1024), int64(0))
		mBase = m.M
		*(*int64)(unsafe.Add(mBase, uint32(v69))) = v81
		v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		*(*uint8)(unsafe.Add(mBase, uint32(v69)+8)) = uint8(v83)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(9)
	}
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F__jumbleNode(m, l0, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		return
	} else {
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		F__jumbleNode(m, l0, v90)
		mBase = m.M
		v92 = m.ExcPending
		if v92 != 0 {
			return
		} else {
			return
		}
	}
}
func F_jit_compile_expr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
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
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v4 == v2 {
		v23 = v2
		return v23
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+176))
		v9 = int32(9)
		if v8&v9 != v9 {
			v23 = v2
			return v23
		} else {
			v13 = F_provider_init(m)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				if v13 == int32(0) {
					v23 = v2
					return v23
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_jit_compile_expr[0]))
					v21 = m.T0[v20].(func(*base.Module, int32) int32)(m, l0)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v23 = v21
						return v23
					}
				}
			}
		}
	}
}
