package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_box_area(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_box_ar(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_reinterpret_f64(v3)
	}
}
func F_box_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v18 float64
	_ = v18
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v28 float64
	_ = v28
	var v32 float64
	_ = v32
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v46 float64
	_ = v46
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
	var v58 float64
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v114 float64
	_ = v114
	var v115 float64
	_ = v115
	var v117 float64
	_ = v117
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v140 int64
	_ = v140
	var v142 int64
	_ = v142
	var v144 float64
	_ = v144
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v18 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.B2i32(base.F64_le(v13, v14) == v4)|base.B2i32(base.F64_ge(v13, v18) == v4) != 0 {
		v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v11))) = v18
		v42 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v40
		*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
		*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v42
		v46 = F_lseg_closept_point(m, l0, v11, l2)
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return float64(0)
		} else {
			v50 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v50
			v52 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v40
			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
			*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
			v58 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return float64(0)
			} else {
				v61 = int64(9223372036854775807)
				v63 = int64(9218868437227405312)
				v66 = int32(0)
				if base.B2i32(base.Ui64(v63) < base.Ui64(base.I64_reinterpret_f64(v58)&v61))|base.B2i32(base.F64_gt(v46, v58) == v66)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v46)&v61) <= base.Ui64(v63)) == v66 {
					if l0 != 0 {
						v77 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v77
						v79 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v79
					} else {
					}
					v81 = v58
				} else {
					v81 = v46
				}
				v82 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
				v83 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
				v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v84
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v83
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v84
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v82
				v91 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return float64(0)
				} else {
					v94 = int64(9223372036854775807)
					v96 = int64(9218868437227405312)
					v99 = int32(0)
					if base.B2i32(base.Ui64(v96) < base.Ui64(base.I64_reinterpret_f64(v91)&v94))|base.B2i32(base.F64_gt(v81, v91) == v99)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v81)&v94) <= base.Ui64(v96)) == v99 {
						if l0 != 0 {
							v110 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v110
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v112
						} else {
						}
						v114 = v91
					} else {
						v114 = v81
					}
					v115 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v115
					v117 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v84
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v83
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v117
					v123 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return float64(0)
					} else {
						v126 = int64(9223372036854775807)
						v128 = int64(9218868437227405312)
						if base.B2i32(base.Ui64(v128) < base.Ui64(base.I64_reinterpret_f64(v123)&v126))|base.B2i32(base.F64_gt(v114, v123) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v114)&v126) <= base.Ui64(v128)) != 0 {
							v144 = v114
						} else {
							if l0 != 0 {
								v140 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v140
								v142 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v142
							} else {
							}
							v144 = v123
						}
						m.G0 = v11 + int32(48)
						return v144
					}
				}
			}
		}
	} else {
		v23 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
		v24 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		if base.F64_le(v23, v24) == int32(0) {
			v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v18
			v42 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v40
			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
			*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v42
			v46 = F_lseg_closept_point(m, l0, v11, l2)
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return float64(0)
			} else {
				v50 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v50
				v52 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v40
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
				v58 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return float64(0)
				} else {
					v61 = int64(9223372036854775807)
					v63 = int64(9218868437227405312)
					v66 = int32(0)
					if base.B2i32(base.Ui64(v63) < base.Ui64(base.I64_reinterpret_f64(v58)&v61))|base.B2i32(base.F64_gt(v46, v58) == v66)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v46)&v61) <= base.Ui64(v63)) == v66 {
						if l0 != 0 {
							v77 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v77
							v79 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v79
						} else {
						}
						v81 = v58
					} else {
						v81 = v46
					}
					v82 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
					v83 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
					v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v84
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v83
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v84
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v82
					v91 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return float64(0)
					} else {
						v94 = int64(9223372036854775807)
						v96 = int64(9218868437227405312)
						v99 = int32(0)
						if base.B2i32(base.Ui64(v96) < base.Ui64(base.I64_reinterpret_f64(v91)&v94))|base.B2i32(base.F64_gt(v81, v91) == v99)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v81)&v94) <= base.Ui64(v96)) == v99 {
							if l0 != 0 {
								v110 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v110
								v112 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v112
							} else {
							}
							v114 = v91
						} else {
							v114 = v81
						}
						v115 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v115
						v117 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v84
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v83
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v117
						v123 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return float64(0)
						} else {
							v126 = int64(9223372036854775807)
							v128 = int64(9218868437227405312)
							if base.B2i32(base.Ui64(v128) < base.Ui64(base.I64_reinterpret_f64(v123)&v126))|base.B2i32(base.F64_gt(v114, v123) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v114)&v126) <= base.Ui64(v128)) != 0 {
								v144 = v114
							} else {
								if l0 != 0 {
									v140 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v140
									v142 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v142
								} else {
								}
								v144 = v123
							}
							m.G0 = v11 + int32(48)
							return v144
						}
					}
				}
			}
		} else {
			v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
			if base.F64_le(v28, v23) == int32(0) {
				v40 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v18
				v42 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v40
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v42
				v46 = F_lseg_closept_point(m, l0, v11, l2)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return float64(0)
				} else {
					v50 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v50
					v52 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v40
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v18
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v52
					v58 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return float64(0)
					} else {
						v61 = int64(9223372036854775807)
						v63 = int64(9218868437227405312)
						v66 = int32(0)
						if base.B2i32(base.Ui64(v63) < base.Ui64(base.I64_reinterpret_f64(v58)&v61))|base.B2i32(base.F64_gt(v46, v58) == v66)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v46)&v61) <= base.Ui64(v63)) == v66 {
							if l0 != 0 {
								v77 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v77
								v79 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v79
							} else {
							}
							v81 = v58
						} else {
							v81 = v46
						}
						v82 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
						v83 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v84
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v83
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v84
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v82
						v91 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return float64(0)
						} else {
							v94 = int64(9223372036854775807)
							v96 = int64(9218868437227405312)
							v99 = int32(0)
							if base.B2i32(base.Ui64(v96) < base.Ui64(base.I64_reinterpret_f64(v91)&v94))|base.B2i32(base.F64_gt(v81, v91) == v99)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v81)&v94) <= base.Ui64(v96)) == v99 {
								if l0 != 0 {
									v110 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v110
									v112 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v112
								} else {
								}
								v114 = v91
							} else {
								v114 = v81
							}
							v115 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v115
							v117 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v84
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v83
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v117
							v123 = F_lseg_closept_point(m, v11+int32(32), v11, l2)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return float64(0)
							} else {
								v126 = int64(9223372036854775807)
								v128 = int64(9218868437227405312)
								if base.B2i32(base.Ui64(v128) < base.Ui64(base.I64_reinterpret_f64(v123)&v126))|base.B2i32(base.F64_gt(v114, v123) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v114)&v126) <= base.Ui64(v128)) != 0 {
									v144 = v114
								} else {
									if l0 != 0 {
										v140 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v140
										v142 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v142
									} else {
									}
									v144 = v123
								}
								m.G0 = v11 + int32(48)
								return v144
							}
						}
					}
				}
			} else {
				v32 = float64(0)
				if l0 == int32(0) {
					v144 = v32
				} else {
					v35 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v35
					v37 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v37
					v144 = v32
				}
				m.G0 = v11 + int32(48)
				return v144
			}
		}
	}
}
func F_box_cn(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v23 float64
	_ = v23
	var v24 int32
	_ = v24
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v41 float64
	_ = v41
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v48 float64
	_ = v48
	var v49 int32
	_ = v49
	var v50 float64
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 float64
	_ = v59
	var v62 float64
	_ = v62
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v78 float64
	_ = v78
	var v79 int32
	_ = v79
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v117 float64
	_ = v117
	var v118 int32
	_ = v118
	var v119 float64
	_ = v119
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v127 float64
	_ = v127
	var v130 float64
	_ = v130
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v146 float64
	_ = v146
	var v147 int32
	_ = v147
	var v148 float64
	_ = v148
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v9 = base.F64_add(v7, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v7), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v25 = v9
	goto L3
L3:
	;
	if l2 != 0 {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return
L5:
	;
	v25 = v23
	goto L3
L6:
	;
	return
L7:
	;
	v83 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v84 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v85 = base.F64_add(v83, v84)
	v87 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v85), v87)|base.F64_eq(base.F64_abs(v83), v87)|base.F64_eq(base.F64_abs(v84), v87) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v26 == int32(453) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v59 = math.Float64frombits(uint64(0x7ff0000000000000))
	v62 = base.F64_mul(v25, float64(0.5))
	if base.F64_eq(base.F64_abs(v25), v59)|base.F64_ne(base.F64_abs(v62), v59) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v29 != 0 {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v31 = math.Float64frombits(uint64(0x7ff0000000000000))
	v34 = base.F64_mul(v25, float64(0.5))
	if base.F64_eq(base.F64_abs(v25), v31)|base.F64_ne(base.F64_abs(v34), v31) == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v52 != int32(453) {
		goto L7
	} else {
		goto L22
	}
L16:
	;
	v41 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v43 = float64(0)
	if base.F64_eq(v25, v43)|base.F64_ne(v34, v43) != 0 {
		v50 = v34
		goto L15
	} else {
		goto L20
	}
L19:
	;
	v50 = v41
	goto L15
L20:
	;
	v48 = F_float_underflow_error_ext(m, l2)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v50 = v48
	goto L15
L22:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v55 == int32(0) {
		goto L7
	} else {
		goto L23
	}
L23:
	;
	goto L6
L24:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v80
	goto L7
L25:
	;
	v70 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v72 = float64(0)
	if base.F64_eq(v25, v72)|base.F64_ne(v62, v72) != 0 {
		v80 = v62
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v80 = v70
	goto L24
L29:
	;
	v78 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v80 = v78
	goto L24
L31:
	;
	v99 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	v101 = v85
	goto L33
L33:
	;
	if l2 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v101 = v99
	goto L33
L35:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v148
	goto L6
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v102 == int32(453) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v127 = math.Float64frombits(uint64(0x7ff0000000000000))
	v130 = base.F64_mul(v101, float64(0.5))
	if base.F64_eq(base.F64_abs(v101), v127)|base.F64_ne(base.F64_abs(v130), v127) == int32(0) {
		goto L49
	} else {
		goto L50
	}
L39:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v105 != 0 {
		goto L6
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v107 = math.Float64frombits(uint64(0x7ff0000000000000))
	v110 = base.F64_mul(v101, float64(0.5))
	if base.F64_eq(base.F64_abs(v101), v107)|base.F64_ne(base.F64_abs(v110), v107) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	v117 = F_float_overflow_error_ext(m, l2)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v119 = float64(0)
	if base.F64_eq(v101, v119)|base.F64_ne(v110, v119) != 0 {
		v148 = v110
		goto L35
	} else {
		goto L47
	}
L46:
	;
	v148 = v117
	goto L35
L47:
	;
	v124 = F_float_underflow_error_ext(m, l2)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	v148 = v124
	goto L35
L49:
	;
	v138 = F_float_overflow_error_ext(m, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v140 = float64(0)
	if base.F64_eq(v101, v140)|base.F64_ne(v130, v140) != 0 {
		v148 = v130
		goto L35
	} else {
		goto L53
	}
L52:
	;
	v148 = v138
	goto L35
L53:
	;
	v146 = F_float_underflow_error_ext(m, int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v148 = v146
	goto L35
}
func F_box_contain_pt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v12 float64
	_ = v12
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v21 float64
	_ = v21
	var v25 int64
	_ = v25
	v4 = int64(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	if base.F64_le(v6, v8) == int32(0) {
		v25 = v4
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		if base.F64_le(v12, v6) == int32(0) {
			v25 = v4
		} else {
			v16 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			if base.F64_le(v16, v17) == int32(0) {
				v25 = v4
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
				v25 = base.I64_extend_i32_u(base.F64_le(v21, v16))
			}
		}
	}
	return v25
}
func F_box_copy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	v4 = F_palloc(m, int32(32))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+24)) = v8
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+16)) = v10
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v4)+8)) = v12
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v4))) = v14
		return v4
	}
}
func F_box_diagonal(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_palloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
		*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v3)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v11
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v3)+16))
		*(*float64)(unsafe.Add(mBase, uint32(v5)+16)) = v13
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v3)+24))
		*(*float64)(unsafe.Add(mBase, uint32(v5)+24)) = v15
		return base.I64_extend_i32_u(v5)
	}
}
func F_box_div(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 float64
	_ = v33
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v48 float64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v54 float64
	_ = v54
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	v8 = m.G0
	v9 = int32(32)
	v10 = v8 - v9
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_palloc(m, v9)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		F_point_div_point(m, v10+int32(16), v13, v12)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			F_point_div_point(m, v10, v13+int32(16), v12)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				v29 = int64(9223372036854775807)
				v31 = int64(9218868437227405312)
				v33 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
				if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v27)&v29) <= base.Ui64(v31))&(base.B2i32(base.Ui64(v31) < base.Ui64(base.I64_reinterpret_f64(v33)&v29))|base.F64_lt(v27, v33)) == int32(0) {
					v44 = v27
					v45 = v33
				} else {
					v44 = v33
					v45 = v27
				}
				*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v45
				*(*float64)(unsafe.Add(mBase, uint32(v15))) = v44
				v48 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				v50 = int64(9223372036854775807)
				v52 = int64(9218868437227405312)
				v54 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
				if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v48)&v50) <= base.Ui64(v52))&(base.B2i32(base.Ui64(v52) < base.Ui64(base.I64_reinterpret_f64(v54)&v50))|base.F64_gt(v54, v48)) == int32(0) {
					v65 = v54
					v66 = v48
				} else {
					v65 = v48
					v66 = v54
				}
				*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v65
				*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v66
				m.G0 = v10 + int32(32)
				return base.I64_extend_i32_u(v15)
			}
		}
	}
}
func F_box_overabove(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
	return base.I64_extend_i32_u(base.F64_le(v3, base.F64_add(v5, float64(1e-06))))
}
func F_box_overbelow(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	return base.I64_extend_i32_u(base.F64_le(v3, base.F64_add(v5, float64(1e-06))))
}
func F_box_penalty(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 float64
	_ = v11
	var v17 float64
	_ = v17
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v35 float64
	_ = v35
	var v41 float64
	_ = v41
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v48 float64
	_ = v48
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v76 float64
	_ = v76
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v81 float64
	_ = v81
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v104 float64
	_ = v104
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v17 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v17)&int64(9223372036854775807)) {
			v23 = v17
		} else {
			v23 = v11
		}
		if base.F64_lt(v11, v17) != 0 {
			v25 = v17
		} else {
			v25 = v23
		}
		v27 = v25
	} else {
		v27 = v11
	}
	*(*float64)(unsafe.Add(mBase, uint32(v9))) = v27
	v29 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.Ui64(base.I64_reinterpret_f64(v29)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v35 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v35)&int64(9223372036854775807)) {
			v41 = v35
		} else {
			v41 = v29
		}
		if base.F64_lt(v29, v35) != 0 {
			v43 = v35
		} else {
			v43 = v41
		}
		v45 = v43
	} else {
		v45 = v29
	}
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v45
	v47 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v48 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v47)&int64(9223372036854775807)) {
			v59 = v48
		} else {
			v59 = v47
		}
		if base.F64_gt(v47, v48) != 0 {
			v61 = v48
		} else {
			v61 = v59
		}
		v62 = v61
	} else {
		v62 = v47
	}
	*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v62
	v64 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v65 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(base.I64_reinterpret_f64(v65)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807)) {
			v76 = v65
		} else {
			v76 = v64
		}
		if base.F64_gt(v64, v65) != 0 {
			v78 = v65
		} else {
			v78 = v76
		}
		v79 = v78
	} else {
		v79 = v64
	}
	*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = v79
	v81 = F_size_box(m, v9)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		return float64(0)
	} else {
		v85 = F_size_box(m, l0)
		mBase = m.M
		v86 = m.ExcPending
		if v86 != 0 {
			return float64(0)
		} else {
			v87 = base.F64_sub(v81, v85)
			v89 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v87), v89)|base.F64_eq(base.F64_abs(v81), v89)|base.F64_eq(base.F64_abs(v85), v89) == int32(0) {
				v102 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return float64(0)
				} else {
					v104 = v102
					m.G0 = v9 + int32(32)
					return v104
				}
			} else {
				v104 = v87
				m.G0 = v9 + int32(32)
				return v104
			}
		}
	}
}
