package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_dist_cpoly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_dist_ppoly_internal(m, v6, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
			v15 = base.F64_sub(v12, v14)
			v17 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v15), v17)|base.F64_eq(base.F64_abs(v12), v17)|base.F64_eq(base.F64_abs(v14), v17) == int32(0) {
				F_float_overflow_error(m)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v31 = float64(0)
				if base.F64_lt(v15, v31) != 0 {
					v34 = v31
				} else {
					v34 = v15
				}
				v35 = F_Float8GetDatum(m, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					return v35
				}
			}
		}
	}
}
func F_dist_ls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 float64
	_ = v14
	var v15 int32
	_ = v15
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v22 float64
	_ = v22
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_lseg_interpt_line(m, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v25 = float64(0)
			v26 = F_Float8GetDatum(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v26
			}
		} else {
			v14 = F_line_closept_point(m, int32(0), v7, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v19 = F_line_closept_point(m, int32(0), v7, v6+int32(16))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if base.F64_lt(v14, v19) != 0 {
						v22 = v14
					} else {
						v22 = v19
					}
					v25 = v22
					v26 = F_Float8GetDatum(m, v25)
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						return v26
					}
				}
			}
		}
	}
}
func F_dist_pathp(m *base.Module, l0 int32) int32 {
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
	var v9 float64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_dist_ppath_internal(m, v8, v4)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_Float8GetDatum(m, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_dist_ppath(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_dist_ppath_internal(m, v2, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_Float8GetDatum(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_dist_ppath_internal(m *base.Module, l0 int32, l1 int32) float64 {
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
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v30 float64
	_ = v30
	var v33 int32
	_ = v33
	var v36 float64
	_ = v36
	var v37 int32
	_ = v37
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v49 float64
	_ = v49
	var v50 int32
	_ = v50
	var v66 float64
	_ = v66
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v80 float64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v95 float64
	_ = v95
	var v96 int32
	_ = v96
	var v103 float64
	_ = v103
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v122 float64
	_ = v122
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 <= int32(0) {
		v122 = float64(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v122
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = l1 + v14<<(uint(int32(4))%32)
	v21 = *(*float64)(unsafe.Add(mBase, uint32(v20)))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = v21
	v23 = *(*float64)(unsafe.Add(mBase, uint32(v20)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v23
	v25 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v25
	v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v27
	v30 = F_lseg_closept_point(m, int32(0), v12, l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v36 = float64(0)
	goto L5
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 < int32(2) {
		v122 = v36
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return float64(0)
L7:
	;
	v36 = v30
	goto L5
L8:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = v40
	v42 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v42
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v44
	v46 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v46
	v49 = F_lseg_closept_point(m, int32(0), v12, l0)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v17 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v67 < int32(3) {
		v122 = v66
		goto L1
	} else {
		goto L16
	}
L11:
	;
	v66 = v49
	goto L10
L12:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v49)&int64(9223372036854775807)) {
		v66 = v36
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v36)&int64(9223372036854775807)) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	if base.F64_gt(v36, v49) == int32(0) {
		v66 = v36
		goto L10
	} else {
		goto L15
	}
L15:
	;
	goto L11
L16:
	;
	v76 = int32(2)
	v80 = v66
	goto L17
L17:
	;
	v83 = v76 << (uint(int32(4)) % 32)
	v84 = l1 + v83
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v84)))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = v85
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v84)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v87
	v89 = l1 + int32(16) + v83
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v89)))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v90
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v89)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v92
	v95 = F_lseg_closept_point(m, int32(0), v12, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L19
	}
L18:
	;
	v122 = v110
	goto L1
L19:
	;
	if base.Ui64(base.I64_reinterpret_f64(v95)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.F64_gt(v80, v95) != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v110 = v80
	goto L22
L22:
	;
	v112 = v76 + int32(1)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v112 < v113 {
		v76 = v112
		v80 = v110
		goto L17
	} else {
		goto L29
	}
L23:
	;
	v103 = v95
	goto L25
L24:
	;
	v103 = v80
	goto L25
L25:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v80)&int64(9223372036854775807)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v109 = v95
	goto L28
L27:
	;
	v109 = v103
	goto L28
L28:
	;
	v110 = v109
	goto L22
L29:
	;
	goto L18
}
func F_dist_ps(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_lseg_closept_point(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_Float8GetDatum(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
