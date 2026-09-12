package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_compute_bitmap_pages(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 float64
	_ = v27
	var v30 float64
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 float64
	_ = v50
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v59 float64
	_ = v59
	var v60 int32
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v72 float64
	_ = v72
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 float64
	_ = v100
	var v103 int32
	_ = v103
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v106 float64
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 float64
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v115 int32
	_ = v115
	var v118 float64
	_ = v118
	var v119 float64
	_ = v119
	var v123 float64
	_ = v123
	var v124 float64
	_ = v124
	var v128 float64
	_ = v128
	var v132 float64
	_ = v132
	var v137 float64
	_ = v137
	var v147 float64
	_ = v147
	var v149 float64
	_ = v149
	var v153 float64
	_ = v153
	var v162 float64
	_ = v162
	var v165 float64
	_ = v165
	var v168 float64
	_ = v168
	var v173 float64
	_ = v173
	var v177 float64
	_ = v177
	var v181 float64
	_ = v181
	var v189 float64
	_ = v189
	var v195 float64
	_ = v195
	var v204 float64
	_ = v204
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v24 - int32(280) {
	case 0:
		v50 = *(*float64)(unsafe.Add(mBase, _consts[384]))
		v53 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
		v55 = *(*float64)(unsafe.Add(mBase, uint32(l2)+96))
		v59 = base.F64_add(base.F64_mul(base.F64_mul(v50, float64(0.1)), v53), v55)
		v60 = l2 + int32(104)
		v61 = float64(1e+100)
		v62 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
		v63 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		v64 = base.F64_mul(v62, v63)
		if base.F64_gt(v64, v61) != 0 {
			v76 = v61
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807)) {
				v76 = v61
			} else {
				v72 = float64(1)
				if base.F64_le(v64, v72) != 0 {
					v76 = v72
				} else {
					v76 = base.F64_nearest(v64)
				}
			}
		}
		v77 = int32(1)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		if base.Ui32(v78) <= base.Ui32(v77) {
			v81 = v77
		} else {
			v81 = v78
		}
		v82 = base.F64_convert_i32_u(v81)
		v83 = base.F64_add(v82, v82)
		v86 = base.F64_div(base.F64_mul(v76, v83), base.F64_add(v83, v76))
		v87 = base.F64_convert_i32_u(v78)
		if base.F64_lt(v86, v87) != 0 {
			v89 = v86
		} else {
			v89 = v87
		}
		v90 = int32(16)
		v92 = *(*int32)(unsafe.Add(mBase, _consts[326]))
		v96 = base.I32_div_u_s(v92<<(uint(int32(10))%32), int32(56))
		if base.Ui32(v96) <= base.Ui32(v90) {
			v99 = v90
		} else {
			v99 = v96
		}
		v100 = base.F64_convert_i32_s(v99)
		if base.F64_gt(l3, float64(1)) != 0 {
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
			v104 = F_get_indexpath_pages(m, l2)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return float64(0)
			} else {
				v106 = base.F64_mul(l3, v76)
				v107 = int32(1)
				if base.Ui32(v103) <= base.Ui32(v107) {
					v110 = v107
				} else {
					v110 = v103
				}
				v111 = base.F64_convert_i32_u(v110)
				v112 = base.F64_add(v111, v111)
				v113 = float64(1)
				v115 = *(*int32)(unsafe.Add(mBase, _consts[385]))
				v118 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
				v119 = base.F64_add(v104, v118)
				if base.F64_gt(v119, v113) != 0 {
					v123 = v119
				} else {
					v123 = v113
				}
				v124 = base.F64_div(base.F64_mul(v111, base.F64_convert_i32_s(v115)), v123)
				if base.F64_le(v124, float64(1)) != 0 {
					v128 = v113
				} else {
					v128 = base.F64_ceil(v124)
				}
				if base.F64_le(v111, v128) != 0 {
					v132 = base.F64_div(base.F64_mul(v106, v112), base.F64_add(v112, v106))
					if base.F64_ge(v132, v111) != 0 {
						v149 = v111
					} else {
						v149 = base.F64_ceil(v132)
					}
				} else {
					v137 = base.F64_div(base.F64_mul(v112, v128), base.F64_sub(v112, v128))
					if base.F64_ge(v137, v106) != 0 {
						v147 = base.F64_div(base.F64_mul(v106, v112), base.F64_add(v112, v106))
					} else {
						v147 = base.F64_add(v128, base.F64_div(base.F64_mul(base.F64_sub(v111, v128), base.F64_sub(v106, v137)), v111))
					}
					v149 = base.F64_ceil(v147)
				}
				v153 = base.F64_div(v149, l3)
				if base.F64_gt(v89, v100) == int32(0) {
					v195 = v76
				} else {
					v162 = float64(0)
					v165 = base.F64_add(v89, base.F64_mul(v100, float64(-0.5)))
					if base.F64_lt(v165, v162) != 0 {
						v168 = v162
					} else {
						v168 = v165
					}
					if base.F64_gt(v168, float64(0)) == int32(0) {
						v195 = v76
					} else {
						v173 = float64(1e+100)
						v177 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
						v181 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v89, v168), v89)), v177), base.F64_mul(base.F64_div(v168, v89), v177))
						if base.F64_gt(v181, v173) != 0 {
							v195 = v173
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v181)&int64(9223372036854775807)) {
								v195 = v173
							} else {
								v189 = float64(1)
								if base.F64_le(v181, v189) != 0 {
									v195 = v189
								} else {
									v195 = base.F64_nearest(v181)
								}
							}
						}
					}
				}
				if l4 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
				} else {
				}
				if l5 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l5))) = v195
				} else {
				}
				m.G0 = v22 + int32(16)
				if base.F64_ge(v153, v82) != 0 {
					v204 = v82
				} else {
					v204 = base.F64_ceil(v153)
				}
				return v204
			}
		} else {
			v153 = v86
			if base.F64_gt(v89, v100) == int32(0) {
				v195 = v76
			} else {
				v162 = float64(0)
				v165 = base.F64_add(v89, base.F64_mul(v100, float64(-0.5)))
				if base.F64_lt(v165, v162) != 0 {
					v168 = v162
				} else {
					v168 = v165
				}
				if base.F64_gt(v168, float64(0)) == int32(0) {
					v195 = v76
				} else {
					v173 = float64(1e+100)
					v177 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
					v181 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v89, v168), v89)), v177), base.F64_mul(base.F64_div(v168, v89), v177))
					if base.F64_gt(v181, v173) != 0 {
						v195 = v173
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v181)&int64(9223372036854775807)) {
							v195 = v173
						} else {
							v189 = float64(1)
							if base.F64_le(v181, v189) != 0 {
								v195 = v189
							} else {
								v195 = base.F64_nearest(v181)
							}
						}
					}
				}
			}
			if l4 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
			} else {
			}
			if l5 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l5))) = v195
			} else {
			}
			m.G0 = v22 + int32(16)
			if base.F64_ge(v153, v82) != 0 {
				v204 = v82
			} else {
				v204 = base.F64_ceil(v153)
			}
			return v204
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return float64(0)
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = v39
			F_errmsg_internal(m, int32(463673), v22)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return float64(0)
			} else {
				F_errfinish(m, int32(475268), int32(1149), int32(391620))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v27 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
		v59 = v27
		v60 = l2 + int32(80)
		v61 = float64(1e+100)
		v62 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
		v63 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		v64 = base.F64_mul(v62, v63)
		if base.F64_gt(v64, v61) != 0 {
			v76 = v61
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807)) {
				v76 = v61
			} else {
				v72 = float64(1)
				if base.F64_le(v64, v72) != 0 {
					v76 = v72
				} else {
					v76 = base.F64_nearest(v64)
				}
			}
		}
		v77 = int32(1)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		if base.Ui32(v78) <= base.Ui32(v77) {
			v81 = v77
		} else {
			v81 = v78
		}
		v82 = base.F64_convert_i32_u(v81)
		v83 = base.F64_add(v82, v82)
		v86 = base.F64_div(base.F64_mul(v76, v83), base.F64_add(v83, v76))
		v87 = base.F64_convert_i32_u(v78)
		if base.F64_lt(v86, v87) != 0 {
			v89 = v86
		} else {
			v89 = v87
		}
		v90 = int32(16)
		v92 = *(*int32)(unsafe.Add(mBase, _consts[326]))
		v96 = base.I32_div_u_s(v92<<(uint(int32(10))%32), int32(56))
		if base.Ui32(v96) <= base.Ui32(v90) {
			v99 = v90
		} else {
			v99 = v96
		}
		v100 = base.F64_convert_i32_s(v99)
		if base.F64_gt(l3, float64(1)) != 0 {
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
			v104 = F_get_indexpath_pages(m, l2)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return float64(0)
			} else {
				v106 = base.F64_mul(l3, v76)
				v107 = int32(1)
				if base.Ui32(v103) <= base.Ui32(v107) {
					v110 = v107
				} else {
					v110 = v103
				}
				v111 = base.F64_convert_i32_u(v110)
				v112 = base.F64_add(v111, v111)
				v113 = float64(1)
				v115 = *(*int32)(unsafe.Add(mBase, _consts[385]))
				v118 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
				v119 = base.F64_add(v104, v118)
				if base.F64_gt(v119, v113) != 0 {
					v123 = v119
				} else {
					v123 = v113
				}
				v124 = base.F64_div(base.F64_mul(v111, base.F64_convert_i32_s(v115)), v123)
				if base.F64_le(v124, float64(1)) != 0 {
					v128 = v113
				} else {
					v128 = base.F64_ceil(v124)
				}
				if base.F64_le(v111, v128) != 0 {
					v132 = base.F64_div(base.F64_mul(v106, v112), base.F64_add(v112, v106))
					if base.F64_ge(v132, v111) != 0 {
						v149 = v111
					} else {
						v149 = base.F64_ceil(v132)
					}
				} else {
					v137 = base.F64_div(base.F64_mul(v112, v128), base.F64_sub(v112, v128))
					if base.F64_ge(v137, v106) != 0 {
						v147 = base.F64_div(base.F64_mul(v106, v112), base.F64_add(v112, v106))
					} else {
						v147 = base.F64_add(v128, base.F64_div(base.F64_mul(base.F64_sub(v111, v128), base.F64_sub(v106, v137)), v111))
					}
					v149 = base.F64_ceil(v147)
				}
				v153 = base.F64_div(v149, l3)
				if base.F64_gt(v89, v100) == int32(0) {
					v195 = v76
				} else {
					v162 = float64(0)
					v165 = base.F64_add(v89, base.F64_mul(v100, float64(-0.5)))
					if base.F64_lt(v165, v162) != 0 {
						v168 = v162
					} else {
						v168 = v165
					}
					if base.F64_gt(v168, float64(0)) == int32(0) {
						v195 = v76
					} else {
						v173 = float64(1e+100)
						v177 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
						v181 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v89, v168), v89)), v177), base.F64_mul(base.F64_div(v168, v89), v177))
						if base.F64_gt(v181, v173) != 0 {
							v195 = v173
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v181)&int64(9223372036854775807)) {
								v195 = v173
							} else {
								v189 = float64(1)
								if base.F64_le(v181, v189) != 0 {
									v195 = v189
								} else {
									v195 = base.F64_nearest(v181)
								}
							}
						}
					}
				}
				if l4 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
				} else {
				}
				if l5 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l5))) = v195
				} else {
				}
				m.G0 = v22 + int32(16)
				if base.F64_ge(v153, v82) != 0 {
					v204 = v82
				} else {
					v204 = base.F64_ceil(v153)
				}
				return v204
			}
		} else {
			v153 = v86
			if base.F64_gt(v89, v100) == int32(0) {
				v195 = v76
			} else {
				v162 = float64(0)
				v165 = base.F64_add(v89, base.F64_mul(v100, float64(-0.5)))
				if base.F64_lt(v165, v162) != 0 {
					v168 = v162
				} else {
					v168 = v165
				}
				if base.F64_gt(v168, float64(0)) == int32(0) {
					v195 = v76
				} else {
					v173 = float64(1e+100)
					v177 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
					v181 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v89, v168), v89)), v177), base.F64_mul(base.F64_div(v168, v89), v177))
					if base.F64_gt(v181, v173) != 0 {
						v195 = v173
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v181)&int64(9223372036854775807)) {
							v195 = v173
						} else {
							v189 = float64(1)
							if base.F64_le(v181, v189) != 0 {
								v195 = v189
							} else {
								v195 = base.F64_nearest(v181)
							}
						}
					}
				}
			}
			if l4 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
			} else {
			}
			if l5 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l5))) = v195
			} else {
			}
			m.G0 = v22 + int32(16)
			if base.F64_ge(v153, v82) != 0 {
				v204 = v82
			} else {
				v204 = base.F64_ceil(v153)
			}
			return v204
		}
	case 4:
		v30 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
		v59 = v30
		v60 = l2 + int32(80)
		v61 = float64(1e+100)
		v62 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
		v63 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		v64 = base.F64_mul(v62, v63)
		if base.F64_gt(v64, v61) != 0 {
			v76 = v61
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807)) {
				v76 = v61
			} else {
				v72 = float64(1)
				if base.F64_le(v64, v72) != 0 {
					v76 = v72
				} else {
					v76 = base.F64_nearest(v64)
				}
			}
		}
		v77 = int32(1)
		v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		if base.Ui32(v78) <= base.Ui32(v77) {
			v81 = v77
		} else {
			v81 = v78
		}
		v82 = base.F64_convert_i32_u(v81)
		v83 = base.F64_add(v82, v82)
		v86 = base.F64_div(base.F64_mul(v76, v83), base.F64_add(v83, v76))
		v87 = base.F64_convert_i32_u(v78)
		if base.F64_lt(v86, v87) != 0 {
			v89 = v86
		} else {
			v89 = v87
		}
		v90 = int32(16)
		v92 = *(*int32)(unsafe.Add(mBase, _consts[326]))
		v96 = base.I32_div_u_s(v92<<(uint(int32(10))%32), int32(56))
		if base.Ui32(v96) <= base.Ui32(v90) {
			v99 = v90
		} else {
			v99 = v96
		}
		v100 = base.F64_convert_i32_s(v99)
		if base.F64_gt(l3, float64(1)) != 0 {
			v103 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
			v104 = F_get_indexpath_pages(m, l2)
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return float64(0)
			} else {
				v106 = base.F64_mul(l3, v76)
				v107 = int32(1)
				if base.Ui32(v103) <= base.Ui32(v107) {
					v110 = v107
				} else {
					v110 = v103
				}
				v111 = base.F64_convert_i32_u(v110)
				v112 = base.F64_add(v111, v111)
				v113 = float64(1)
				v115 = *(*int32)(unsafe.Add(mBase, _consts[385]))
				v118 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
				v119 = base.F64_add(v104, v118)
				if base.F64_gt(v119, v113) != 0 {
					v123 = v119
				} else {
					v123 = v113
				}
				v124 = base.F64_div(base.F64_mul(v111, base.F64_convert_i32_s(v115)), v123)
				if base.F64_le(v124, float64(1)) != 0 {
					v128 = v113
				} else {
					v128 = base.F64_ceil(v124)
				}
				if base.F64_le(v111, v128) != 0 {
					v132 = base.F64_div(base.F64_mul(v106, v112), base.F64_add(v112, v106))
					if base.F64_ge(v132, v111) != 0 {
						v149 = v111
					} else {
						v149 = base.F64_ceil(v132)
					}
				} else {
					v137 = base.F64_div(base.F64_mul(v112, v128), base.F64_sub(v112, v128))
					if base.F64_ge(v137, v106) != 0 {
						v147 = base.F64_div(base.F64_mul(v106, v112), base.F64_add(v112, v106))
					} else {
						v147 = base.F64_add(v128, base.F64_div(base.F64_mul(base.F64_sub(v111, v128), base.F64_sub(v106, v137)), v111))
					}
					v149 = base.F64_ceil(v147)
				}
				v153 = base.F64_div(v149, l3)
				if base.F64_gt(v89, v100) == int32(0) {
					v195 = v76
				} else {
					v162 = float64(0)
					v165 = base.F64_add(v89, base.F64_mul(v100, float64(-0.5)))
					if base.F64_lt(v165, v162) != 0 {
						v168 = v162
					} else {
						v168 = v165
					}
					if base.F64_gt(v168, float64(0)) == int32(0) {
						v195 = v76
					} else {
						v173 = float64(1e+100)
						v177 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
						v181 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v89, v168), v89)), v177), base.F64_mul(base.F64_div(v168, v89), v177))
						if base.F64_gt(v181, v173) != 0 {
							v195 = v173
						} else {
							if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v181)&int64(9223372036854775807)) {
								v195 = v173
							} else {
								v189 = float64(1)
								if base.F64_le(v181, v189) != 0 {
									v195 = v189
								} else {
									v195 = base.F64_nearest(v181)
								}
							}
						}
					}
				}
				if l4 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
				} else {
				}
				if l5 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l5))) = v195
				} else {
				}
				m.G0 = v22 + int32(16)
				if base.F64_ge(v153, v82) != 0 {
					v204 = v82
				} else {
					v204 = base.F64_ceil(v153)
				}
				return v204
			}
		} else {
			v153 = v86
			if base.F64_gt(v89, v100) == int32(0) {
				v195 = v76
			} else {
				v162 = float64(0)
				v165 = base.F64_add(v89, base.F64_mul(v100, float64(-0.5)))
				if base.F64_lt(v165, v162) != 0 {
					v168 = v162
				} else {
					v168 = v165
				}
				if base.F64_gt(v168, float64(0)) == int32(0) {
					v195 = v76
				} else {
					v173 = float64(1e+100)
					v177 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
					v181 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v89, v168), v89)), v177), base.F64_mul(base.F64_div(v168, v89), v177))
					if base.F64_gt(v181, v173) != 0 {
						v195 = v173
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v181)&int64(9223372036854775807)) {
							v195 = v173
						} else {
							v189 = float64(1)
							if base.F64_le(v181, v189) != 0 {
								v195 = v189
							} else {
								v195 = base.F64_nearest(v181)
							}
						}
					}
				}
			}
			if l4 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
			} else {
			}
			if l5 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l5))) = v195
			} else {
			}
			m.G0 = v22 + int32(16)
			if base.F64_ge(v153, v82) != 0 {
				v204 = v82
			} else {
				v204 = base.F64_ceil(v153)
			}
			return v204
		}
	}
}
func F_cost_bitmap_heap_scan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 float64) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v64 float64
	_ = v64
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v109 float64
	_ = v109
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v119 float64
	_ = v119
	var v121 float64
	_ = v121
	var v122 float64
	_ = v122
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v143 int32
	_ = v143
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v149 int32
	_ = v149
	var v155 float64
	_ = v155
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v173 float64
	_ = v173
	var v177 float64
	_ = v177
	var v180 float64
	_ = v180
	var v182 float64
	_ = v182
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v188 int32
	_ = v188
	var v192 float64
	_ = v192
	var v195 float64
	_ = v195
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	if l3 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = l3 + int32(8)
	goto L3
L2:
	;
	v23 = l2 + int32(16)
	goto L3
L3:
	;
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v23)))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v24
	v30 = F_compute_bitmap_pages(m, l1, l2, l4, l5, v15+int32(-32), v15+int32(-40))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+116))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(v17)+32))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
	F_get_tablespace_page_costs(m, v34, v15+int32(-56), v15+int32(-48))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if base.F64_ge(v30, float64(2)) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if l3 != 0 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v17)+16))
	v46 = int32(1)
	if base.Ui32(v32) <= base.Ui32(v46) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v55 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
	v57 = v55
	goto L7
L11:
	;
	v49 = v46
	goto L13
L12:
	;
	v49 = v32
	goto L13
L13:
	;
	v57 = base.F64_sub(v43, base.F64_mul(base.F64_sub(v43, v44), base.F64_sqrt(base.F64_div(v30, base.F64_convert_i32_u(v49)))))
	goto L7
L14:
	;
	v139 = *(*float64)(unsafe.Add(mBase, _consts[380]))
	v141 = *(*float64)(unsafe.Add(mBase, uint32(v17)+24))
	v142 = base.F64_mul(base.F64_add(v137, v139), v141)
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v143 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v59 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = l1
	v64 = float64(0)
	if v58 == int32(0) {
		v109 = v64
		v116 = v64
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v122 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v135 = v121
	v137 = v122
	goto L14
L18:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v119 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v135 = base.F64_add(v116, v117)
	v137 = base.F64_add(v109, v119)
	goto L14
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v69 <= int32(0) {
		v109 = v64
		v116 = float64(0)
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v76 = int32(0)
	goto L21
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87+v76<<(uint(int32(2))%32))))
	v94 = F_cost_qual_eval_walker(m, v91, v15+int32(-24))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L23
	}
L22:
	;
	v100 = *(*float64)(unsafe.Add(mBase, uint32(v17)+56))
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v17)+48))
	v109 = v100
	v116 = v101
	goto L18
L23:
	;
	v97 = v76 + int32(1)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v97 < v98 {
		v76 = v97
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v185 = *(*float64)(unsafe.Add(mBase, uint32(v184)+24))
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v184)+16))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, _consts[383])))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v188 ^ int32(1)
	v192 = float64(0)
	v195 = base.F64_add(v186, base.F64_add(base.F64_add(v33, v192), v135))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v195
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v195, base.F64_add(base.F64_mul(v185, v182), base.F64_add(base.F64_add(base.F64_mul(v30, v57), v192), v180)))
	m.G0 = v17 - int32(-64)
	return
L26:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v180 = v142
	v182 = v146
	goto L25
L27:
	;
	goto L28
L28:
	;
	v147 = base.F64_convert_i32_u(v143)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _consts[381])))
	if v149 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v155 = base.F64_add(base.F64_mul(v147, float64(-0.3)), float64(1))
	if base.F64_gt(v155, float64(0)) != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v161 = v147
	goto L31
L31:
	;
	v163 = float64(1e+100)
	v164 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v165 = base.F64_div(v164, v161)
	if base.F64_gt(v165, v163) != 0 {
		v177 = v163
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v159 = v155
	goto L34
L33:
	;
	v159 = math.Float64frombits(uint64(0x8000000000000000))
	goto L34
L34:
	;
	v161 = base.F64_add(v159, v147)
	goto L31
L35:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v177
	v180 = base.F64_div(v142, v161)
	v182 = v177
	goto L25
L36:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v165)&int64(9223372036854775807)) {
		v177 = v163
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v173 = float64(1)
	if base.F64_le(v165, v173) != 0 {
		v177 = v173
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v177 = base.F64_nearest(v165)
	goto L35
}
func F_create_bitmap_and_path(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 float64
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 float64
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 float64
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 float64
	_ = v109
	var v111 float64
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 float64
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v144 float64
	_ = v144
	var v145 int32
	_ = v145
	var v146 float64
	_ = v146
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v154 int32
	_ = v154
	var v163 float64
	_ = v163
	var v165 float64
	_ = v165
	v4 = int32(0)
	v8 = float64(0)
	v14 = F_palloc0(m, int32(88))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = int64(1447403979035)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v21
	if l2 == int32(0) {
		v60 = v4
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v67 = F_get_baserel_parampathinfo(m, l0, l1, v60)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L13
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v25 <= int32(0) {
		v60 = v4
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v32 = v4
	v33 = v4
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v32<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v60 = v49
	goto L3
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v48 = v46
	goto L10
L9:
	;
	v48 = int32(0)
	goto L10
L10:
	;
	v49 = F_bms_add_members(m, v33, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v52 = v32 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v52 < v53 {
		v32 = v52
		v33 = v49
		goto L6
	} else {
		goto L12
	}
L12:
	;
	goto L7
L13:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+20)) = uint8(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v67
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+72)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v14)+64)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v69
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+21)) = uint8(v72)
	v80 = m.G0
	v82 = v80 - int32(16)
	m.G0 = v82
	v84 = float64(1)
	if l2 == v69 {
		v163 = v8
		v165 = v84
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14)+56)) = v163
	*(*float64)(unsafe.Add(mBase, uint32(v14)+48)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v14)+80)) = v165
	m.G0 = v82 + int32(16)
	return v14
L15:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v88 <= int32(0) {
		v163 = v8
		v165 = v84
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v91 = int32(0)
	if v91 < v88 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v94 = v88
	goto L19
L18:
	;
	v94 = v91
	goto L19
L19:
	;
	v96 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v104 = v69
	v109 = v8
	v111 = v84
	goto L20
L20:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v99+v104<<(uint(int32(2))%32))))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	switch v118 - int32(280) {
	case 0:
		goto L23
	default:
		goto L24
	case 3, 4:
		goto L25
	}
L21:
	;
	v163 = v150
	v165 = v152
	goto L14
L22:
	;
	v146 = base.F64_add(v109, v144)
	if v104&int32(1073741823) != 0 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	v138 = *(*float64)(unsafe.Add(mBase, uint32(v117)+32))
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v117)+96))
	v144 = base.F64_add(base.F64_mul(base.F64_mul(v96, float64(0.1)), v138), v140)
	v145 = v117 + int32(104)
	goto L22
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v121 = *(*float64)(unsafe.Add(mBase, uint32(v117)+56))
	v144 = v121
	v145 = v117 + int32(80)
	goto L22
L26:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v128
	F_errmsg_internal(m, int32(463673), v82)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(475268), int32(1149), int32(391620))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	v150 = base.F64_add(base.F64_mul(v96, float64(100)), v146)
	goto L31
L30:
	;
	v150 = v146
	goto L31
L31:
	;
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v145)))
	v152 = base.F64_mul(v111, v151)
	v154 = v104 + int32(1)
	if v154 != v94 {
		v104 = v154
		v109 = v150
		v111 = v152
		goto L20
	} else {
		goto L32
	}
L32:
	;
	goto L21
}
func F_generate_bitmap_or_paths(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v29 int32
	_ = v29
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v251 int32
	_ = v251
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v313 int32
	_ = v313
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v401 int32
	_ = v401
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v531 int32
	_ = v531
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v588 int32
	_ = v588
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v637 int32
	_ = v637
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v714 int32
	_ = v714
	var v722 int32
	_ = v722
	var v738 int32
	_ = v738
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 float64
	_ = v852
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v862 float64
	_ = v862
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v874 int32
	_ = v874
	var v882 int32
	_ = v882
	var v895 float64
	_ = v895
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 float64
	_ = v915
	var v916 float64
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v930 int32
	_ = v930
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v986 int32
	_ = v986
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1048 int32
	_ = v1048
	var v1070 int32
	_ = v1070
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1136 int32
	_ = v1136
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1171 float64
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1189 float64
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1205 int32
	_ = v1205
	var v1220 float64
	_ = v1220
	var v1221 float64
	_ = v1221
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	var v1248 int32
	_ = v1248
	var v1251 float64
	_ = v1251
	var v1252 float64
	_ = v1252
	var v1256 float64
	_ = v1256
	var v1259 float64
	_ = v1259
	var v1261 float64
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1266 float64
	_ = v1266
	var v1267 float64
	_ = v1267
	var v1268 float64
	_ = v1268
	var v1270 int32
	_ = v1270
	var v1296 float64
	_ = v1296
	var v1297 float64
	_ = v1297
	var v1304 float64
	_ = v1304
	var v1307 float64
	_ = v1307
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1314 int32
	_ = v1314
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1337 int32
	_ = v1337
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1358 int32
	_ = v1358
	var v1369 int32
	_ = v1369
	v5 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(32)
	m.G0 = v31
	v33 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l2 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v1358 + int32(32)
	return v1369
L4:
	;
	v40 = l0
	v41 = l1
	v42 = l2
	v52 = v31
	v55 = v33
	v61 = v5
	v63 = v5
	goto L9
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if int32(0) < v37 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v1358 = v31
	v1369 = v5
	goto L3
L8:
	;
	goto L7
L9:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68+v61<<(uint(int32(2))%32))))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+52))
	goto L12
L10:
	;
	v1358 = v1326
	v1369 = v1337
	goto L3
L11:
	;
	v1343 = v61 + int32(1)
	v1344 = *(*int32)(unsafe.Add(mBase, uint32(v1316)+4))
	if v1343 < v1344 {
		v40 = v1314
		v41 = v1315
		v42 = v1316
		v52 = v1326
		v55 = v1329
		v61 = v1343
		v63 = v1337
		goto L9
	} else {
		goto L226
	}
L12:
	;
	if base.B2i32(v73 != int32(0)) == int32(0) {
		v1314 = v40
		v1315 = v41
		v1316 = v42
		v1326 = v52
		v1329 = v55
		v1337 = v63
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)+52))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v79 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_pfree(m, v751)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L120
	}
L15:
	;
	v82 = int32(0)
	v84 = F_palloc(m, v82)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v41)+68))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v90 = F_palloc(m, v87*int32(24))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v749 = v82
	v751 = v84
	goto L14
L19:
	;
	v92 = int32(0)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v93 <= v92 {
		v749 = v79
		v751 = v90
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v103 = int32(-1)
	v108 = v92
	v120 = int32(0)
	goto L21
L21:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v126+v108<<(uint(int32(2))%32))))
	v132 = v103 + int32(1)
	v135 = v90 + v132*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v135)+20)) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v135)+16)) = v132
	*(*int64)(unsafe.Add(mBase, uint32(v135)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v135))) = int64(-1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v142 != int32(318) {
		v401 = v120
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v401 == int32(0) {
		v749 = v79
		v751 = v90
		goto L14
	} else {
		goto L71
	}
L23:
	;
	v408 = v108 + int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v408 < v409 {
		v103 = v132
		v108 = v408
		v120 = v401
		goto L21
	} else {
		goto L70
	}
L24:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	if v146 != int32(17) {
		v401 = v120
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145)+28))
	if v149 == int32(0) {
		v401 = v120
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v152 != int32(2) {
		v401 = v120
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v157 == int32(27) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	v161 = v160
	goto L30
L29:
	;
	v161 = v156
	goto L30
L30:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v163 == int32(27) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v167 = v166
	goto L33
L32:
	;
	v167 = v162
	goto L33
L33:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v145)+4))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v130)+48))
	v170 = F_bms_is_member(m, v86, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v41)+108))
	if v194 == int32(0) {
		v401 = v120
		goto L23
	} else {
		goto L50
	}
L35:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	v182 = F_bms_is_member(m, v86, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L44
	}
L36:
	;
	if v170 == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v130)+44))
	v175 = F_bms_is_member(m, v86, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v175 != 0 {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v177 = F_contain_volatile_functions(m, v161)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v177 != 0 {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	v179 = F_get_commutator(m, v168)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v179 != 0 {
		v192 = v167
		v193 = v179
		goto L34
	} else {
		goto L43
	}
L43:
	;
	v401 = v120
	goto L23
L44:
	;
	if v182 == int32(0) {
		v401 = v120
		goto L23
	} else {
		goto L45
	}
L45:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v130)+48))
	v187 = F_bms_is_member(m, v86, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v187 != 0 {
		v401 = v120
		goto L23
	} else {
		goto L47
	}
L47:
	;
	v189 = F_contain_volatile_functions(m, v167)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	if v189 != 0 {
		v401 = v120
		goto L23
	} else {
		goto L49
	}
L49:
	;
	v192 = v161
	v193 = v168
	goto L34
L50:
	;
	v197 = int32(0)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	if v199 <= v197 {
		v401 = v120
		goto L23
	} else {
		goto L51
	}
L51:
	;
	v208 = v199
	v219 = v197
	v222 = v197
	v224 = v120
	goto L52
L52:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230+v222<<(uint(int32(2))%32))))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+110)))
	if v235 != int32(1) {
		v354 = v208
		v365 = v219
		v370 = v224
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v401 = v370
	goto L23
L54:
	;
	v377 = v222 + int32(1)
	if v377 < v354 {
		v208 = v354
		v219 = v365
		v222 = v377
		v224 = v370
		goto L52
	} else {
		goto L69
	}
L55:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v234)+107)))
	if v238 != int32(1) {
		v354 = v208
		v365 = v219
		v370 = v224
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v241 = int32(0)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v234)+40))
	if v241 < v242 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if int32(0) <= v342 {
		v401 = v336
		goto L23
	} else {
		goto L68
	}
L58:
	;
	v251 = v241
	goto L61
L59:
	;
	goto L60
L60:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v336 = v224
	v342 = v313
	goto L57
L61:
	;
	v273 = F_match_index_to_operand(m, v192, v251, v234)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	if v273 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+8)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v135)+4)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v219
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v145)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v135)+12)) = v278
	v336 = int32(1)
	v342 = v219
	goto L57
L65:
	;
	goto L66
L66:
	;
	v282 = v251 + int32(1)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v234)+40))
	if v282 < v283 {
		v251 = v282
		goto L61
	} else {
		goto L67
	}
L67:
	;
	goto L62
L68:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v194)+4))
	v354 = v347
	v365 = v219 + int32(1)
	v370 = v336
	goto L54
L69:
	;
	goto L53
L70:
	;
	goto L22
L71:
	;
	F_pg_qsort(m, v90, v87, int32(24), int32(821))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if int32(2) <= v87 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v425 = int32(1)
	goto L76
L74:
	;
	goto L75
L75:
	;
	F_pg_qsort(m, v90, v87, int32(24), int32(822))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L85
	}
L76:
	;
	v448 = int32(24)
	v450 = v90 + v425*v448
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v450)))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v450-v448)))
	if v451 != v454 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L75
L78:
	;
	v478 = v425 + int32(1)
	if v478 != v87 {
		v425 = v478
		goto L76
	} else {
		goto L84
	}
L79:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v450)+4))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v450-int32(20))))
	if v456 != v459 {
		goto L78
	} else {
		goto L80
	}
L80:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v450)+8))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v450-int32(16))))
	if v461 != v464 {
		goto L78
	} else {
		goto L81
	}
L81:
	;
	if v451 == int32(-1) {
		goto L78
	} else {
		goto L82
	}
L82:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v450)+12))
	v471 = *(*int32)(unsafe.Add(mBase, uint32(v450-int32(12))))
	if v468 != v471 {
		goto L78
	} else {
		goto L83
	}
L83:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v450-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v450)+20)) = v475
	goto L78
L84:
	;
	goto L77
L85:
	;
	if v87 <= int32(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v749 = int32(0)
	v751 = v90
	goto L14
L87:
	;
	goto L88
L88:
	;
	v515 = int32(0)
	v522 = int32(1)
	v523 = v515
	v531 = v515
	goto L89
L89:
	;
	if v523 < int32(0) {
		v714 = v523
		v722 = v531
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v749 = v722
	v751 = v90
	goto L14
L91:
	;
	v738 = v522 + int32(1)
	if v738 <= v87 {
		v522 = v738
		v523 = v714
		v531 = v722
		goto L89
	} else {
		goto L119
	}
L92:
	;
	if v522 == v87 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	if v522-v523 != int32(1) {
		goto L101
	} else {
		goto L102
	}
L94:
	;
	v549 = int32(24)
	v551 = v90 + v522*v549
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	v555 = v90 + v523*v549
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)))
	if v552 != v556 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v551)+4))
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	if v558 != v559 {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v551)+8))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v555)+8))
	if v561 != v562 {
		goto L93
	} else {
		goto L97
	}
L97:
	;
	if v552 == int32(-1) {
		goto L93
	} else {
		goto L98
	}
L98:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v551)+12))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v555)+12))
	if v566 == v567 {
		v714 = v523
		v722 = v531
		goto L91
	} else {
		goto L99
	}
L99:
	;
	goto L93
L100:
	;
	v707 = F_lappend(m, v531, v706)
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L1
	} else {
		goto L118
	}
L101:
	;
	v575 = int32(0)
	if v523 < v522 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v90+v523*int32(24))+16))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v669+v673<<(uint(int32(2))%32))))
	v706 = v677
	goto L100
L104:
	;
	v583 = v523
	v584 = v575
	v588 = v575
	goto L107
L105:
	;
	v633 = v575
	v637 = v575
	goto L106
L106:
	;
	v655 = F_make_orclause(m, v633)
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L1
	} else {
		goto L115
	}
L107:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v90+v583*int32(24))+16))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v606+v610<<(uint(int32(2))%32))))
	v615 = F_lappend(m, v588, v614)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	v633 = v622
	v637 = v615
	goto L106
L109:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v614)))
	if v617 == int32(318) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v614)+4))
	v621 = v620
	goto L112
L111:
	;
	v621 = v614
	goto L112
L112:
	;
	v622 = F_lappend(m, v584, v621)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v625 = v583 + int32(1)
	if v625 != v522 {
		v583 = v625
		v584 = v622
		v588 = v615
		goto L107
	} else {
		goto L114
	}
L114:
	;
	goto L108
L115:
	;
	v657 = F_make_orclause(m, v637)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+8)))
	v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+11)))
	v661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+12)))
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+10)))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v72)+20))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v72)+32))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v72)+36))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v72)+40))
	v667 = F_make_plain_restrictinfo(m, v40, v655, v657, v659, v660, v661, v662, v663, v664, v665, v666)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v706 = v667
	goto L100
L118:
	;
	v714 = v522
	v722 = v707
	goto L91
L119:
	;
	goto L90
L120:
	;
	v770 = int32(0)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v72)+52))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v772)+8))
	if v773 != v749 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v775 = F_list_copy(m, v55)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	v779 = v770
	goto L123
L123:
	;
	if v749 == int32(0) {
		v1048 = v770
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v777 = F_list_delete(m, v775, v72)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	v779 = v777
	goto L123
L126:
	;
	if v779 != 0 {
		goto L184
	} else {
		goto L185
	}
L127:
	;
	v782 = int32(0)
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
	if v783 <= v782 {
		v1048 = v770
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v790 = v782
	v793 = v770
	goto L129
L129:
	;
	v814 = *(*int32)(unsafe.Add(mBase, uint32(v749)+12))
	v818 = *(*int32)(unsafe.Add(mBase, uint32(v814+v790<<(uint(int32(2))%32))))
	if v818 == int32(0) {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	v1048 = v1036
	goto L126
L131:
	;
	v1038 = v790 + int32(1)
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v749)+4))
	if v1038 < v1039 {
		v790 = v1038
		v793 = v1036
		goto L129
	} else {
		goto L183
	}
L132:
	;
	v1006 = F_list_concat(m, v793, v986)
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L1
	} else {
		goto L182
	}
L133:
	;
	if v970 == int32(0) {
		goto L177
	} else {
		goto L178
	}
L134:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v818)+52))
	goto L141
L135:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v818)))
	if v821 != int32(21) {
		goto L134
	} else {
		goto L136
	}
L136:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v818)+4))
	if v824 != 0 {
		goto L134
	} else {
		goto L137
	}
L137:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v818)+8))
	v826 = F_build_paths_for_OR(m, v40, v41, v825, v55)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v828 = F_generate_bitmap_or_paths(m, v40, v41, v825, v55)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v830 = F_list_concat(m, v826, v828)
	mBase = m.M
	v831 = m.ExcPending
	if v831 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	v970 = v830
	goto L133
L141:
	;
	if v832 != int32(0) {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v818)+52))
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v835)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+8)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v52)+28)) = v818
	v842 = F_list_make1_impl(m, int32(1), v52+int32(8))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L1
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v52)+12)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v52)+16)) = v818
	v965 = F_list_make1_impl(m, int32(1), v52+int32(12))
	mBase = m.M
	v966 = m.ExcPending
	if v966 != 0 {
		goto L1
	} else {
		goto L175
	}
L145:
	;
	if v836 == int32(0) {
		goto L155
	} else {
		goto L156
	}
L146:
	;
	v844 = F_build_paths_for_OR(m, v40, v41, v842, v779)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v844 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v861 = int32(0)
	v862 = float64(0)
	goto L145
L149:
	;
	goto L150
L150:
	;
	v850 = F_choose_bitmap_and(m, v40, v41, v844)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v852 = *(*float64)(unsafe.Add(mBase, uint32(v850)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v850
	*(*int32)(unsafe.Add(mBase, uint32(v52)+24)) = v850
	v858 = F_list_make1_impl(m, int32(1), v52+int32(4))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if v779 != 0 {
		v861 = v858
		v862 = v852
		goto L145
	} else {
		goto L153
	}
L153:
	;
	if v858 != 0 {
		v986 = v858
		goto L132
	} else {
		goto L154
	}
L154:
	;
	v861 = v858
	v862 = v852
	goto L145
L155:
	;
	if v861 != 0 {
		v986 = v861
		goto L132
	} else {
		goto L174
	}
L156:
	;
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v836)+4))
	if v865 <= int32(0) {
		goto L155
	} else {
		goto L157
	}
L157:
	;
	v868 = int32(0)
	v874 = v868
	v882 = v868
	v895 = float64(0)
	goto L158
L158:
	;
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v836)+12))
	v903 = *(*int32)(unsafe.Add(mBase, uint32(v899+v874<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v903
	*(*int32)(unsafe.Add(mBase, uint32(v52)+20)) = v903
	v907 = F_list_make1_impl(m, int32(1), v52)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L1
	} else {
		goto L160
	}
L159:
	;
	if v917 == int32(0) {
		goto L155
	} else {
		goto L166
	}
L160:
	;
	v909 = F_build_paths_for_OR(m, v40, v41, v907, v779)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	if v909 == int32(0) {
		goto L155
	} else {
		goto L162
	}
L162:
	;
	v913 = F_choose_bitmap_and(m, v40, v41, v909)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v915 = *(*float64)(unsafe.Add(mBase, uint32(v913)+56))
	v916 = base.F64_add(v895, v915)
	v917 = F_lappend(m, v882, v913)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v920 = v874 + int32(1)
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v836)+4))
	if v920 < v921 {
		v874 = v920
		v882 = v917
		v895 = v916
		goto L158
	} else {
		goto L165
	}
L165:
	;
	goto L159
L166:
	;
	if v861 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v927 = F_list_concat(m, v793, v917)
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L1
	} else {
		goto L170
	}
L168:
	;
	goto L169
L169:
	;
	if base.F64_gt(v916, v862) != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v1036 = v927
	goto L131
L171:
	;
	v930 = v861
	goto L173
L172:
	;
	v930 = v917
	goto L173
L173:
	;
	v986 = v930
	goto L132
L174:
	;
	v1048 = int32(0)
	goto L126
L175:
	;
	v967 = F_build_paths_for_OR(m, v40, v41, v965, v55)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v970 = v967
	goto L133
L177:
	;
	v1048 = int32(0)
	goto L126
L178:
	;
	goto L179
L179:
	;
	v974 = F_choose_bitmap_and(m, v40, v41, v970)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L1
	} else {
		goto L180
	}
L180:
	;
	v976 = F_lappend(m, v793, v974)
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v1036 = v976
	goto L131
L182:
	;
	v1036 = v1006
	goto L131
L183:
	;
	goto L130
L184:
	;
	F_list_free(m, v779)
	mBase = m.M
	v1070 = m.ExcPending
	if v1070 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	if v1048 == int32(0) {
		v1314 = v40
		v1315 = v41
		v1316 = v42
		v1326 = v52
		v1329 = v55
		v1337 = v63
		goto L11
	} else {
		goto L188
	}
L187:
	;
	goto L186
L188:
	;
	v1073 = int32(0)
	v1076 = F_palloc0(m, int32(88))
	mBase = m.M
	v1077 = m.ExcPending
	if v1077 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+8)) = v41
	*(*int64)(unsafe.Add(mBase, uint32(v1076))) = int64(1451698946332)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+12)) = v1081
	if v1048 == int32(0) {
		v1136 = v1073
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v1159 = F_get_baserel_parampathinfo(m, v40, v41, v1136)
	mBase = m.M
	v1160 = m.ExcPending
	if v1160 != 0 {
		goto L1
	} else {
		goto L200
	}
L191:
	;
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1085 <= int32(0) {
		v1136 = v1073
		goto L190
	} else {
		goto L192
	}
L192:
	;
	v1092 = v1073
	v1093 = v1073
	goto L193
L193:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+12))
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1116+v1092<<(uint(int32(2))%32))))
	v1121 = *(*int32)(unsafe.Add(mBase, uint32(v1120)+16))
	if v1121 != 0 {
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v1136 = v1125
	goto L190
L195:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(v1121)+4))
	v1124 = v1122
	goto L197
L196:
	;
	v1124 = int32(0)
	goto L197
L197:
	;
	v1125 = F_bms_add_members(m, v1093, v1124)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L1
	} else {
		goto L198
	}
L198:
	;
	v1128 = v1092 + int32(1)
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1128 < v1129 {
		v1092 = v1128
		v1093 = v1125
		goto L193
	} else {
		goto L199
	}
L199:
	;
	goto L194
L200:
	;
	v1161 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1076)+20)) = uint8(v1161)
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+16)) = v1159
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+72)) = v1048
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+64)) = v1161
	*(*int32)(unsafe.Add(mBase, uint32(v1076)+24)) = v1161
	*(*uint8)(unsafe.Add(mBase, uint32(v1076)+21)) = uint8(v1164)
	v1171 = float64(0)
	v1174 = m.G0
	v1176 = v1174 - int32(16)
	m.G0 = v1176
	if v1048 == v1161 {
		v1296 = v1171
		v1297 = v1171
		goto L201
	} else {
		goto L202
	}
L201:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1076)+56)) = v1296
	*(*float64)(unsafe.Add(mBase, uint32(v1076)+48)) = v1296
	*(*int64)(unsafe.Add(mBase, uint32(v1076)+32)) = int64(0)
	v1304 = float64(1)
	if base.F64_lt(v1297, v1304) != 0 {
		goto L222
	} else {
		goto L223
	}
L202:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+4))
	if v1181 <= int32(0) {
		v1296 = v1171
		v1297 = v1171
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v1184 = int32(0)
	if v1184 < v1181 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v1187 = v1181
	goto L206
L205:
	;
	v1187 = v1184
	goto L206
L206:
	;
	v1189 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1048)+12))
	v1205 = v1161
	v1220 = float64(0)
	v1221 = v1171
	goto L207
L207:
	;
	v1224 = int32(2)
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v1192+v1205<<(uint(v1224)%32))))
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v1227)))
	if base.Ui32(v1224) <= base.Ui32(v1228-int32(283)) {
		goto L211
	} else {
		goto L212
	}
L208:
	;
	v1296 = v1266
	v1297 = v1268
	goto L201
L209:
	;
	v1267 = *(*float64)(unsafe.Add(mBase, uint32(v1264)))
	v1268 = base.F64_add(v1221, v1267)
	v1270 = v1205 + int32(1)
	if v1270 != v1187 {
		v1205 = v1270
		v1220 = v1266
		v1221 = v1268
		goto L207
	} else {
		goto L221
	}
L210:
	;
	v1259 = *(*float64)(unsafe.Add(mBase, uint32(v1227)+32))
	v1261 = *(*float64)(unsafe.Add(mBase, uint32(v1227)+96))
	v1264 = v1227 + int32(104)
	v1266 = base.F64_add(v1220, base.F64_add(base.F64_mul(base.F64_mul(v1189, float64(0.1)), v1259), v1261))
	goto L209
L211:
	;
	if v1228 == int32(280) {
		goto L210
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v1251 = *(*float64)(unsafe.Add(mBase, uint32(v1227)+56))
	v1252 = base.F64_add(v1220, v1251)
	if v1205&int32(1073741823) != 0 {
		goto L218
	} else {
		goto L219
	}
L214:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(v1227)))
	*(*int32)(unsafe.Add(mBase, uint32(v1176))) = v1239
	F_errmsg_internal(m, int32(463673), v1176)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(475268), int32(1149), int32(391620))
	mBase = m.M
	v1248 = m.ExcPending
	if v1248 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	v1256 = base.F64_add(base.F64_mul(v1189, float64(100)), v1252)
	goto L220
L219:
	;
	v1256 = v1252
	goto L220
L220:
	;
	v1264 = v1227 + int32(80)
	v1266 = v1256
	goto L209
L221:
	;
	goto L208
L222:
	;
	v1307 = v1297
	goto L224
L223:
	;
	v1307 = v1304
	goto L224
L224:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1076)+80)) = v1307
	m.G0 = v1176 + int32(16)
	v1312 = F_lappend(m, v63, v1076)
	mBase = m.M
	v1313 = m.ExcPending
	if v1313 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v1314 = v40
	v1315 = v41
	v1316 = v42
	v1326 = v52
	v1329 = v55
	v1337 = v1312
	goto L11
L226:
	;
	goto L10
}
