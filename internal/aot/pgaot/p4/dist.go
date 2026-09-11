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
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
			if base.F64_ne(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v27 = float64(0)
				if base.F64_lt(v15, v27) != 0 {
					v30 = v27
				} else {
					v30 = v15
				}
				v31 = F_Float8GetDatum(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					return v31
				}
			} else {
				if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v27 = float64(0)
					if base.F64_lt(v15, v27) != 0 {
						v30 = v27
					} else {
						v30 = v15
					}
					v31 = F_Float8GetDatum(m, v30)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						return v31
					}
				} else {
					if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v27 = float64(0)
						if base.F64_lt(v15, v27) != 0 {
							v30 = v27
						} else {
							v30 = v15
						}
						v31 = F_Float8GetDatum(m, v30)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							return v31
						}
					} else {
						F_float_overflow_error(m)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
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
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v21 float64
	_ = v21
	var v24 float64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_lseg_interpt_line(m, int32(0), v6, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			v24 = float64(0)
			v25 = F_Float8GetDatum(m, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v25
			}
		} else {
			v13 = F_line_closept_point(m, int32(0), v7, v6)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v18 = F_line_closept_point(m, int32(0), v7, v6+int32(16))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if base.F64_lt(v13, v18) != 0 {
						v21 = v13
					} else {
						v21 = v18
					}
					v24 = v21
					v25 = F_Float8GetDatum(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return v25
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v37 int32
	_ = v37
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v50 float64
	_ = v50
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 float64
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v98 float64
	_ = v98
	var v99 int32
	_ = v99
	var v106 float64
	_ = v106
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 float64
	_ = v125
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v14 <= int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v125
L2:
	;
	v125 = float64(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v18 = l1 + int32(16)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v24 = v14<<(uint(int32(4))%32) + v18 - int32(16)
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v24)))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = v25
	v27 = *(*float64)(unsafe.Add(mBase, uint32(v24)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v27
	v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v29
	v31 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v31
	v34 = F_lseg_closept_point(m, int32(0), v12, l0)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v40 = float64(0)
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v41 < int32(2) {
		v125 = v40
		goto L1
	} else {
		goto L10
	}
L8:
	;
	return float64(0)
L9:
	;
	v40 = v34
	goto L7
L10:
	;
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = v44
	v46 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v46
	v48 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v48
	v50 = *(*float64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v50
	v53 = F_lseg_closept_point(m, int32(0), v12, l0)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v19 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v71 < int32(3) {
		v125 = v70
		goto L1
	} else {
		goto L18
	}
L13:
	;
	v70 = v53
	goto L12
L14:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v53)&int64(9223372036854775807)) {
		v70 = v40
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v40)&int64(9223372036854775807)) {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if base.F64_gt(v40, v53) == int32(0) {
		v70 = v40
		goto L12
	} else {
		goto L17
	}
L17:
	;
	goto L13
L18:
	;
	v78 = int32(2)
	v82 = v70
	goto L19
L19:
	;
	v86 = v18 + v78<<(uint(int32(4))%32)
	v88 = v86 - int32(16)
	v89 = *(*float64)(unsafe.Add(mBase, uint32(v88)))
	*(*float64)(unsafe.Add(mBase, uint32(v12))) = v89
	v91 = *(*float64)(unsafe.Add(mBase, uint32(v88)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+8)) = v91
	v93 = *(*float64)(unsafe.Add(mBase, uint32(v86)))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+16)) = v93
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v86)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v12)+24)) = v95
	v98 = F_lseg_closept_point(m, int32(0), v12, l0)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	v125 = v113
	goto L1
L21:
	;
	if base.Ui64(base.I64_reinterpret_f64(v98)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if base.F64_gt(v82, v98) != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v113 = v82
	goto L24
L24:
	;
	v115 = v78 + int32(1)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v115 < v116 {
		v78 = v115
		v82 = v113
		goto L19
	} else {
		goto L31
	}
L25:
	;
	v106 = v98
	goto L27
L26:
	;
	v106 = v82
	goto L27
L27:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v82)&int64(9223372036854775807)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v112 = v98
	goto L30
L29:
	;
	v112 = v106
	goto L30
L30:
	;
	v113 = v112
	goto L24
L31:
	;
	goto L20
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
