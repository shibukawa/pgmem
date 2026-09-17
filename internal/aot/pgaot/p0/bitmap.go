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
	var v73 float64
	_ = v73
	var v77 float64
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v104 int32
	_ = v104
	var v105 float64
	_ = v105
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v116 int32
	_ = v116
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v125 float64
	_ = v125
	var v129 float64
	_ = v129
	var v133 float64
	_ = v133
	var v138 float64
	_ = v138
	var v148 float64
	_ = v148
	var v150 float64
	_ = v150
	var v154 float64
	_ = v154
	var v163 float64
	_ = v163
	var v166 float64
	_ = v166
	var v169 float64
	_ = v169
	var v174 float64
	_ = v174
	var v178 float64
	_ = v178
	var v182 float64
	_ = v182
	var v191 float64
	_ = v191
	var v197 float64
	_ = v197
	var v205 float64
	_ = v205
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v24 - int32(280) {
	case 0:
		v50 = *(*float64)(unsafe.Add(mBase, _c_F_compute_bitmap_pages[0]))
		v53 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
		v55 = *(*float64)(unsafe.Add(mBase, uint32(l2)+96))
		v59 = base.F64_add(base.F64_mul(base.F64_mul(v50, float64(0.1)), v53), v55)
		v60 = l2 + int32(104)
		v61 = float64(1e+100)
		v62 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
		v63 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		v64 = base.F64_mul(v62, v63)
		if base.F64_gt(v64, v61)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807))) != 0 {
			v77 = v61
		} else {
			v73 = float64(1)
			if base.F64_le(v64, v73) != 0 {
				v77 = v73
			} else {
				v77 = base.F64_nearest(v64)
			}
		}
		v78 = int32(1)
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		if base.Ui32(v79) <= base.Ui32(v78) {
			v82 = v78
		} else {
			v82 = v79
		}
		v83 = base.F64_convert_i32_u(v82)
		v84 = base.F64_add(v83, v83)
		v87 = base.F64_div(base.F64_mul(v77, v84), base.F64_add(v84, v77))
		v88 = base.F64_convert_i32_u(v79)
		if base.F64_lt(v87, v88) != 0 {
			v90 = v87
		} else {
			v90 = v88
		}
		v91 = int32(16)
		v93 = *(*int32)(unsafe.Add(mBase, _c_F_compute_bitmap_pages[1]))
		v97 = base.I32_div_u_s(v93<<(uint(int32(10))%32), int32(56))
		if base.Ui32(v97) <= base.Ui32(v91) {
			v100 = v91
		} else {
			v100 = v97
		}
		v101 = base.F64_convert_i32_s(v100)
		if base.F64_gt(l3, float64(1)) != 0 {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
			v105 = F_get_indexpath_pages(m, l2)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return float64(0)
			} else {
				v107 = base.F64_mul(l3, v77)
				v108 = int32(1)
				if base.Ui32(v104) <= base.Ui32(v108) {
					v111 = v108
				} else {
					v111 = v104
				}
				v112 = base.F64_convert_i32_u(v111)
				v113 = base.F64_add(v112, v112)
				v114 = float64(1)
				v116 = *(*int32)(unsafe.Add(mBase, _c_F_compute_bitmap_pages[2]))
				v119 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
				v120 = base.F64_add(v105, v119)
				if base.F64_gt(v120, v114) != 0 {
					v124 = v120
				} else {
					v124 = v114
				}
				v125 = base.F64_div(base.F64_mul(v112, base.F64_convert_i32_s(v116)), v124)
				if base.F64_le(v125, float64(1)) != 0 {
					v129 = v114
				} else {
					v129 = base.F64_ceil(v125)
				}
				if base.F64_le(v112, v129) != 0 {
					v133 = base.F64_div(base.F64_mul(v107, v113), base.F64_add(v113, v107))
					if base.F64_ge(v133, v112) != 0 {
						v150 = v112
					} else {
						v150 = base.F64_ceil(v133)
					}
				} else {
					v138 = base.F64_div(base.F64_mul(v113, v129), base.F64_sub(v113, v129))
					if base.F64_ge(v138, v107) != 0 {
						v148 = base.F64_div(base.F64_mul(v107, v113), base.F64_add(v113, v107))
					} else {
						v148 = base.F64_add(v129, base.F64_div(base.F64_mul(base.F64_sub(v112, v129), base.F64_sub(v107, v138)), v112))
					}
					v150 = base.F64_ceil(v148)
				}
				v154 = base.F64_div(v150, l3)
				if base.F64_gt(v90, v101) == int32(0) {
					v197 = v77
				} else {
					v163 = float64(0)
					v166 = base.F64_add(v90, base.F64_mul(v101, float64(-0.5)))
					if base.F64_lt(v166, v163) != 0 {
						v169 = v163
					} else {
						v169 = v166
					}
					if base.F64_gt(v169, float64(0)) == int32(0) {
						v197 = v77
					} else {
						v174 = float64(1e+100)
						v178 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
						v182 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v90, v169), v90)), v178), base.F64_mul(base.F64_div(v169, v90), v178))
						if base.F64_gt(v182, v174)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v182)&int64(9223372036854775807))) != 0 {
							v197 = v174
						} else {
							v191 = float64(1)
							if base.F64_le(v182, v191) != 0 {
								v197 = v191
							} else {
								v197 = base.F64_nearest(v182)
							}
						}
					}
				}
				if l4 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
				} else {
				}
				if l5 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l5))) = v197
				} else {
				}
				m.G0 = v22 + int32(16)
				if base.F64_ge(v154, v83) != 0 {
					v205 = v83
				} else {
					v205 = base.F64_ceil(v154)
				}
				return v205
			}
		} else {
			v154 = v87
			if base.F64_gt(v90, v101) == int32(0) {
				v197 = v77
			} else {
				v163 = float64(0)
				v166 = base.F64_add(v90, base.F64_mul(v101, float64(-0.5)))
				if base.F64_lt(v166, v163) != 0 {
					v169 = v163
				} else {
					v169 = v166
				}
				if base.F64_gt(v169, float64(0)) == int32(0) {
					v197 = v77
				} else {
					v174 = float64(1e+100)
					v178 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
					v182 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v90, v169), v90)), v178), base.F64_mul(base.F64_div(v169, v90), v178))
					if base.F64_gt(v182, v174)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v182)&int64(9223372036854775807))) != 0 {
						v197 = v174
					} else {
						v191 = float64(1)
						if base.F64_le(v182, v191) != 0 {
							v197 = v191
						} else {
							v197 = base.F64_nearest(v182)
						}
					}
				}
			}
			if l4 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
			} else {
			}
			if l5 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l5))) = v197
			} else {
			}
			m.G0 = v22 + int32(16)
			if base.F64_ge(v154, v83) != 0 {
				v205 = v83
			} else {
				v205 = base.F64_ceil(v154)
			}
			return v205
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
			F_errmsg_internal(m, int32(_a_F_compute_bitmap_pages_0), v22)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return float64(0)
			} else {
				F_errfinish(m, int32(_a_F_compute_bitmap_pages_1), int32(1149), int32(_a_F_compute_bitmap_pages_2))
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
		if base.F64_gt(v64, v61)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807))) != 0 {
			v77 = v61
		} else {
			v73 = float64(1)
			if base.F64_le(v64, v73) != 0 {
				v77 = v73
			} else {
				v77 = base.F64_nearest(v64)
			}
		}
		v78 = int32(1)
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		if base.Ui32(v79) <= base.Ui32(v78) {
			v82 = v78
		} else {
			v82 = v79
		}
		v83 = base.F64_convert_i32_u(v82)
		v84 = base.F64_add(v83, v83)
		v87 = base.F64_div(base.F64_mul(v77, v84), base.F64_add(v84, v77))
		v88 = base.F64_convert_i32_u(v79)
		if base.F64_lt(v87, v88) != 0 {
			v90 = v87
		} else {
			v90 = v88
		}
		v91 = int32(16)
		v93 = *(*int32)(unsafe.Add(mBase, _c_F_compute_bitmap_pages[1]))
		v97 = base.I32_div_u_s(v93<<(uint(int32(10))%32), int32(56))
		if base.Ui32(v97) <= base.Ui32(v91) {
			v100 = v91
		} else {
			v100 = v97
		}
		v101 = base.F64_convert_i32_s(v100)
		if base.F64_gt(l3, float64(1)) != 0 {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
			v105 = F_get_indexpath_pages(m, l2)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return float64(0)
			} else {
				v107 = base.F64_mul(l3, v77)
				v108 = int32(1)
				if base.Ui32(v104) <= base.Ui32(v108) {
					v111 = v108
				} else {
					v111 = v104
				}
				v112 = base.F64_convert_i32_u(v111)
				v113 = base.F64_add(v112, v112)
				v114 = float64(1)
				v116 = *(*int32)(unsafe.Add(mBase, _c_F_compute_bitmap_pages[2]))
				v119 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
				v120 = base.F64_add(v105, v119)
				if base.F64_gt(v120, v114) != 0 {
					v124 = v120
				} else {
					v124 = v114
				}
				v125 = base.F64_div(base.F64_mul(v112, base.F64_convert_i32_s(v116)), v124)
				if base.F64_le(v125, float64(1)) != 0 {
					v129 = v114
				} else {
					v129 = base.F64_ceil(v125)
				}
				if base.F64_le(v112, v129) != 0 {
					v133 = base.F64_div(base.F64_mul(v107, v113), base.F64_add(v113, v107))
					if base.F64_ge(v133, v112) != 0 {
						v150 = v112
					} else {
						v150 = base.F64_ceil(v133)
					}
				} else {
					v138 = base.F64_div(base.F64_mul(v113, v129), base.F64_sub(v113, v129))
					if base.F64_ge(v138, v107) != 0 {
						v148 = base.F64_div(base.F64_mul(v107, v113), base.F64_add(v113, v107))
					} else {
						v148 = base.F64_add(v129, base.F64_div(base.F64_mul(base.F64_sub(v112, v129), base.F64_sub(v107, v138)), v112))
					}
					v150 = base.F64_ceil(v148)
				}
				v154 = base.F64_div(v150, l3)
				if base.F64_gt(v90, v101) == int32(0) {
					v197 = v77
				} else {
					v163 = float64(0)
					v166 = base.F64_add(v90, base.F64_mul(v101, float64(-0.5)))
					if base.F64_lt(v166, v163) != 0 {
						v169 = v163
					} else {
						v169 = v166
					}
					if base.F64_gt(v169, float64(0)) == int32(0) {
						v197 = v77
					} else {
						v174 = float64(1e+100)
						v178 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
						v182 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v90, v169), v90)), v178), base.F64_mul(base.F64_div(v169, v90), v178))
						if base.F64_gt(v182, v174)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v182)&int64(9223372036854775807))) != 0 {
							v197 = v174
						} else {
							v191 = float64(1)
							if base.F64_le(v182, v191) != 0 {
								v197 = v191
							} else {
								v197 = base.F64_nearest(v182)
							}
						}
					}
				}
				if l4 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
				} else {
				}
				if l5 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l5))) = v197
				} else {
				}
				m.G0 = v22 + int32(16)
				if base.F64_ge(v154, v83) != 0 {
					v205 = v83
				} else {
					v205 = base.F64_ceil(v154)
				}
				return v205
			}
		} else {
			v154 = v87
			if base.F64_gt(v90, v101) == int32(0) {
				v197 = v77
			} else {
				v163 = float64(0)
				v166 = base.F64_add(v90, base.F64_mul(v101, float64(-0.5)))
				if base.F64_lt(v166, v163) != 0 {
					v169 = v163
				} else {
					v169 = v166
				}
				if base.F64_gt(v169, float64(0)) == int32(0) {
					v197 = v77
				} else {
					v174 = float64(1e+100)
					v178 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
					v182 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v90, v169), v90)), v178), base.F64_mul(base.F64_div(v169, v90), v178))
					if base.F64_gt(v182, v174)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v182)&int64(9223372036854775807))) != 0 {
						v197 = v174
					} else {
						v191 = float64(1)
						if base.F64_le(v182, v191) != 0 {
							v197 = v191
						} else {
							v197 = base.F64_nearest(v182)
						}
					}
				}
			}
			if l4 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
			} else {
			}
			if l5 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l5))) = v197
			} else {
			}
			m.G0 = v22 + int32(16)
			if base.F64_ge(v154, v83) != 0 {
				v205 = v83
			} else {
				v205 = base.F64_ceil(v154)
			}
			return v205
		}
	case 4:
		v30 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
		v59 = v30
		v60 = l2 + int32(80)
		v61 = float64(1e+100)
		v62 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
		v63 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
		v64 = base.F64_mul(v62, v63)
		if base.F64_gt(v64, v61)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v64)&int64(9223372036854775807))) != 0 {
			v77 = v61
		} else {
			v73 = float64(1)
			if base.F64_le(v64, v73) != 0 {
				v77 = v73
			} else {
				v77 = base.F64_nearest(v64)
			}
		}
		v78 = int32(1)
		v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
		if base.Ui32(v79) <= base.Ui32(v78) {
			v82 = v78
		} else {
			v82 = v79
		}
		v83 = base.F64_convert_i32_u(v82)
		v84 = base.F64_add(v83, v83)
		v87 = base.F64_div(base.F64_mul(v77, v84), base.F64_add(v84, v77))
		v88 = base.F64_convert_i32_u(v79)
		if base.F64_lt(v87, v88) != 0 {
			v90 = v87
		} else {
			v90 = v88
		}
		v91 = int32(16)
		v93 = *(*int32)(unsafe.Add(mBase, _c_F_compute_bitmap_pages[1]))
		v97 = base.I32_div_u_s(v93<<(uint(int32(10))%32), int32(56))
		if base.Ui32(v97) <= base.Ui32(v91) {
			v100 = v91
		} else {
			v100 = v97
		}
		v101 = base.F64_convert_i32_s(v100)
		if base.F64_gt(l3, float64(1)) != 0 {
			v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
			v105 = F_get_indexpath_pages(m, l2)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return float64(0)
			} else {
				v107 = base.F64_mul(l3, v77)
				v108 = int32(1)
				if base.Ui32(v104) <= base.Ui32(v108) {
					v111 = v108
				} else {
					v111 = v104
				}
				v112 = base.F64_convert_i32_u(v111)
				v113 = base.F64_add(v112, v112)
				v114 = float64(1)
				v116 = *(*int32)(unsafe.Add(mBase, _c_F_compute_bitmap_pages[2]))
				v119 = *(*float64)(unsafe.Add(mBase, uint32(l0)+288))
				v120 = base.F64_add(v105, v119)
				if base.F64_gt(v120, v114) != 0 {
					v124 = v120
				} else {
					v124 = v114
				}
				v125 = base.F64_div(base.F64_mul(v112, base.F64_convert_i32_s(v116)), v124)
				if base.F64_le(v125, float64(1)) != 0 {
					v129 = v114
				} else {
					v129 = base.F64_ceil(v125)
				}
				if base.F64_le(v112, v129) != 0 {
					v133 = base.F64_div(base.F64_mul(v107, v113), base.F64_add(v113, v107))
					if base.F64_ge(v133, v112) != 0 {
						v150 = v112
					} else {
						v150 = base.F64_ceil(v133)
					}
				} else {
					v138 = base.F64_div(base.F64_mul(v113, v129), base.F64_sub(v113, v129))
					if base.F64_ge(v138, v107) != 0 {
						v148 = base.F64_div(base.F64_mul(v107, v113), base.F64_add(v113, v107))
					} else {
						v148 = base.F64_add(v129, base.F64_div(base.F64_mul(base.F64_sub(v112, v129), base.F64_sub(v107, v138)), v112))
					}
					v150 = base.F64_ceil(v148)
				}
				v154 = base.F64_div(v150, l3)
				if base.F64_gt(v90, v101) == int32(0) {
					v197 = v77
				} else {
					v163 = float64(0)
					v166 = base.F64_add(v90, base.F64_mul(v101, float64(-0.5)))
					if base.F64_lt(v166, v163) != 0 {
						v169 = v163
					} else {
						v169 = v166
					}
					if base.F64_gt(v169, float64(0)) == int32(0) {
						v197 = v77
					} else {
						v174 = float64(1e+100)
						v178 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
						v182 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v90, v169), v90)), v178), base.F64_mul(base.F64_div(v169, v90), v178))
						if base.F64_gt(v182, v174)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v182)&int64(9223372036854775807))) != 0 {
							v197 = v174
						} else {
							v191 = float64(1)
							if base.F64_le(v182, v191) != 0 {
								v197 = v191
							} else {
								v197 = base.F64_nearest(v182)
							}
						}
					}
				}
				if l4 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
				} else {
				}
				if l5 != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(l5))) = v197
				} else {
				}
				m.G0 = v22 + int32(16)
				if base.F64_ge(v154, v83) != 0 {
					v205 = v83
				} else {
					v205 = base.F64_ceil(v154)
				}
				return v205
			}
		} else {
			v154 = v87
			if base.F64_gt(v90, v101) == int32(0) {
				v197 = v77
			} else {
				v163 = float64(0)
				v166 = base.F64_add(v90, base.F64_mul(v101, float64(-0.5)))
				if base.F64_lt(v166, v163) != 0 {
					v169 = v163
				} else {
					v169 = v166
				}
				if base.F64_gt(v169, float64(0)) == int32(0) {
					v197 = v77
				} else {
					v174 = float64(1e+100)
					v178 = *(*float64)(unsafe.Add(mBase, uint32(l1)+120))
					v182 = base.F64_add(base.F64_mul(base.F64_mul(v62, base.F64_div(base.F64_sub(v90, v169), v90)), v178), base.F64_mul(base.F64_div(v169, v90), v178))
					if base.F64_gt(v182, v174)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v182)&int64(9223372036854775807))) != 0 {
						v197 = v174
					} else {
						v191 = float64(1)
						if base.F64_le(v182, v191) != 0 {
							v197 = v191
						} else {
							v197 = base.F64_nearest(v182)
						}
					}
				}
			}
			if l4 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l4))) = v59
			} else {
			}
			if l5 != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(l5))) = v197
			} else {
			}
			m.G0 = v22 + int32(16)
			if base.F64_ge(v154, v83) != 0 {
				v205 = v83
			} else {
				v205 = base.F64_ceil(v154)
			}
			return v205
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
	var v107 float64
	_ = v107
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
	var v134 float64
	_ = v134
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
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v166 float64
	_ = v166
	var v175 float64
	_ = v175
	var v179 float64
	_ = v179
	var v181 float64
	_ = v181
	var v184 float64
	_ = v184
	var v185 int32
	_ = v185
	var v186 float64
	_ = v186
	var v187 float64
	_ = v187
	var v189 int32
	_ = v189
	var v193 float64
	_ = v193
	var v196 float64
	_ = v196
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
	v139 = *(*float64)(unsafe.Add(mBase, _c_F_cost_bitmap_heap_scan[0]))
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
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v17)+56)) = v59
	v64 = float64(0)
	if v58 == int32(0) {
		v107 = v64
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
	v134 = v121
	v137 = v122
	goto L14
L18:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(l2)+192))
	v119 = *(*float64)(unsafe.Add(mBase, uint32(l2)+200))
	v134 = base.F64_add(v116, v117)
	v137 = base.F64_add(v107, v119)
	goto L14
L19:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v69 <= int32(0) {
		v107 = v64
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
	v107 = v100
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
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v186 = *(*float64)(unsafe.Add(mBase, uint32(v185)+24))
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v185)+16))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_bitmap_heap_scan[1])))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v189 ^ int32(1)
	v193 = float64(0)
	v196 = base.F64_add(v187, base.F64_add(base.F64_add(v33, v193), v134))
	*(*float64)(unsafe.Add(mBase, uint32(l0)+48)) = v196
	*(*float64)(unsafe.Add(mBase, uint32(l0)+56)) = base.F64_add(v196, base.F64_add(base.F64_mul(v186, v184), base.F64_add(base.F64_add(base.F64_mul(v30, v57), v193), v181)))
	m.G0 = v17 - int32(-64)
	return
L26:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v181 = v142
	v184 = v146
	goto L25
L27:
	;
	goto L28
L28:
	;
	v147 = base.F64_convert_i32_u(v143)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_cost_bitmap_heap_scan[2])))
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
	v164 = float64(1e+100)
	v165 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v166 = base.F64_div(v165, v161)
	if base.F64_gt(v166, v164)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v166)&int64(9223372036854775807))) != 0 {
		v179 = v164
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
	*(*float64)(unsafe.Add(mBase, uint32(l0)+32)) = v179
	v181 = base.F64_div(v142, v161)
	v184 = v179
	goto L25
L36:
	;
	v175 = float64(1)
	if base.F64_le(v166, v175) != 0 {
		v179 = v175
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v179 = base.F64_nearest(v166)
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
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 float64
	_ = v83
	var v87 int32
	_ = v87
	var v91 float64
	_ = v91
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v116 float64
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v132 float64
	_ = v132
	var v134 float64
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 float64
	_ = v144
	var v147 float64
	_ = v147
	var v149 float64
	_ = v149
	var v153 float64
	_ = v153
	var v154 int32
	_ = v154
	var v156 float64
	_ = v156
	var v157 float64
	_ = v157
	var v158 float64
	_ = v158
	var v160 int32
	_ = v160
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v184 int32
	_ = v184
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
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
		v59 = v4
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v67 = F_get_baserel_parampathinfo(m, l0, l1, v59)
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
		v59 = v4
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
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v33<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v59 = v49
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
	v49 = F_bms_add_members(m, v32, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v52 = v33 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v52 < v53 {
		v32 = v49
		v33 = v52
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
	v79 = m.G0
	v81 = v79 - int32(16)
	m.G0 = v81
	v83 = float64(1)
	if l2 == v69 {
		v169 = v83
		v171 = v8
		goto L16
	} else {
		goto L17
	}
L14:
	;
	return v14
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L33
	}
L16:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14)+56)) = v171
	*(*float64)(unsafe.Add(mBase, uint32(v14)+48)) = v171
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = int64(0)
	*(*float64)(unsafe.Add(mBase, uint32(v14)+80)) = v169
	m.G0 = v81 + int32(16)
	goto L14
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v87 <= int32(0) {
		v169 = v83
		v171 = v8
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v91 = *(*float64)(unsafe.Add(mBase, _c_F_create_bitmap_and_path[0]))
	v93 = base.F64_mul(v91, float64(0.1))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	switch v96 - int32(280) {
	case 0:
		goto L20
	default:
		v184 = v95
		goto L15
	case 3:
		goto L21
	case 4:
		goto L22
	}
L19:
	;
	v114 = base.F64_add(v112, float64(0))
	v115 = int32(1)
	v116 = *(*float64)(unsafe.Add(mBase, uint32(v111)))
	if v87 == v115 {
		v169 = v116
		v171 = v114
		goto L16
	} else {
		goto L23
	}
L20:
	;
	v107 = *(*float64)(unsafe.Add(mBase, uint32(v95)+32))
	v109 = *(*float64)(unsafe.Add(mBase, uint32(v95)+96))
	v111 = v95 + int32(104)
	v112 = base.F64_add(base.F64_mul(v93, v107), v109)
	goto L19
L21:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v95)+56))
	v111 = v95 + int32(80)
	v112 = v104
	goto L19
L22:
	;
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v95)+56))
	v111 = v95 + int32(80)
	v112 = v101
	goto L19
L23:
	;
	v119 = int32(0)
	if v119 < v87 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v122 = v87
	goto L26
L25:
	;
	v122 = v119
	goto L26
L26:
	;
	v130 = v115
	v132 = v116
	v134 = v114
	goto L27
L27:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v94+v130<<(uint(int32(2))%32))))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	switch v141 - int32(280) {
	case 0:
		goto L30
	default:
		v184 = v140
		goto L15
	case 3, 4:
		goto L31
	}
L28:
	;
	v169 = v158
	v171 = v156
	goto L16
L29:
	;
	v156 = base.F64_add(base.F64_mul(v91, float64(100)), base.F64_add(v134, v153))
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v154)))
	v158 = base.F64_mul(v132, v157)
	v160 = v130 + int32(1)
	if v160 != v122 {
		v130 = v160
		v132 = v158
		v134 = v156
		goto L27
	} else {
		goto L32
	}
L30:
	;
	v147 = *(*float64)(unsafe.Add(mBase, uint32(v140)+32))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v140)+96))
	v153 = base.F64_add(base.F64_mul(v93, v147), v149)
	v154 = v140 + int32(104)
	goto L29
L31:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v140)+56))
	v153 = v144
	v154 = v140 + int32(80)
	goto L29
L32:
	;
	goto L28
L33:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v200
	F_errmsg_internal(m, int32(_a_F_create_bitmap_and_path_0), v81)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_create_bitmap_and_path_1), int32(1149), int32(_a_F_create_bitmap_and_path_2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
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
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v314 int32
	_ = v314
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v402 int32
	_ = v402
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v636 int32
	_ = v636
	var v643 int32
	_ = v643
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
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v741 int32
	_ = v741
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
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
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 float64
	_ = v855
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 float64
	_ = v865
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v877 int32
	_ = v877
	var v883 int32
	_ = v883
	var v898 float64
	_ = v898
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 float64
	_ = v918
	var v919 float64
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v986 int32
	_ = v986
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1051 int32
	_ = v1051
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1119 int32
	_ = v1119
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1137 int32
	_ = v1137
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1174 float64
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1187 float64
	_ = v1187
	var v1189 float64
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1201 float64
	_ = v1201
	var v1203 float64
	_ = v1203
	var v1207 float64
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 float64
	_ = v1209
	var v1210 float64
	_ = v1210
	var v1211 float64
	_ = v1211
	var v1212 float64
	_ = v1212
	var v1214 float64
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1221 int32
	_ = v1221
	var v1232 int32
	_ = v1232
	var v1248 float64
	_ = v1248
	var v1249 float64
	_ = v1249
	var v1252 int32
	_ = v1252
	var v1255 int32
	_ = v1255
	var v1256 int32
	_ = v1256
	var v1265 float64
	_ = v1265
	var v1267 float64
	_ = v1267
	var v1272 float64
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1276 float64
	_ = v1276
	var v1277 float64
	_ = v1277
	var v1278 float64
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1306 float64
	_ = v1306
	var v1307 float64
	_ = v1307
	var v1314 float64
	_ = v1314
	var v1317 float64
	_ = v1317
	var v1326 int32
	_ = v1326
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1358 int32
	_ = v1358
	var v1363 int32
	_ = v1363
	var v1364 int32
	_ = v1364
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1368 int32
	_ = v1368
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1389 int32
	_ = v1389
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1409 int32
	_ = v1409
	var v1421 int32
	_ = v1421
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
	if l2 == int32(0) {
		v1409 = v31
		v1421 = v5
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v1409 + int32(32)
	return v1421
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v39 <= int32(0) {
		v1409 = v31
		v1421 = v5
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v42 = l0
	v43 = l1
	v44 = l2
	v53 = v31
	v56 = v33
	v63 = v5
	v65 = v5
	goto L6
L6:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v70+v63<<(uint(int32(2))%32))))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+52))
	goto L9
L7:
	;
	v1409 = v1377
	v1421 = v1389
	goto L3
L8:
	;
	v1395 = v63 + int32(1)
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+4))
	if v1395 < v1396 {
		v42 = v1366
		v43 = v1367
		v44 = v1368
		v53 = v1377
		v56 = v1380
		v63 = v1395
		v65 = v1389
		goto L6
	} else {
		goto L225
	}
L9:
	;
	if base.B2i32(v75 != int32(0)) == int32(0) {
		v1366 = v42
		v1367 = v43
		v1368 = v44
		v1377 = v53
		v1380 = v56
		v1389 = v65
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v74)+52))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	if v81 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	F_pfree(m, v752)
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L1
	} else {
		goto L115
	}
L12:
	;
	v84 = int32(0)
	v86 = F_palloc(m, v84)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v43)+68))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v92 = F_palloc(m, v89*int32(24))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v751 = v84
	v752 = v86
	goto L11
L16:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v94 <= int32(0) {
		v751 = v81
		v752 = v92
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v98 = int32(0)
	v103 = int32(-1)
	v113 = v98
	v122 = v98
	goto L18
L18:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128+v113<<(uint(int32(2))%32))))
	v134 = v103 + int32(1)
	v137 = v92 + v134*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v137)+20)) = v134
	*(*int32)(unsafe.Add(mBase, uint32(v137)+16)) = v134
	*(*int64)(unsafe.Add(mBase, uint32(v137)+8)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v137))) = int64(-1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v144 != int32(318) {
		v402 = v122
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v402 == int32(0) {
		v751 = v81
		v752 = v92
		goto L11
	} else {
		goto L68
	}
L20:
	;
	v409 = v113 + int32(1)
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v409 < v410 {
		v103 = v134
		v113 = v409
		v122 = v402
		goto L18
	} else {
		goto L67
	}
L21:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v148 != int32(17) {
		v402 = v122
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	if v151 == int32(0) {
		v402 = v122
		goto L20
	} else {
		goto L23
	}
L23:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v154 != int32(2) {
		v402 = v122
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v159 == int32(27) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v163 = v162
	goto L27
L26:
	;
	v163 = v158
	goto L27
L27:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v165 == int32(27) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	v169 = v168
	goto L30
L29:
	;
	v169 = v164
	goto L30
L30:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v132)+48))
	v172 = F_bms_is_member(m, v88, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v43)+108))
	if v195 == int32(0) {
		v402 = v122
		goto L20
	} else {
		goto L47
	}
L32:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
	v184 = F_bms_is_member(m, v88, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L41
	}
L33:
	;
	if v172 == int32(0) {
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v132)+44))
	v177 = F_bms_is_member(m, v88, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v177 != 0 {
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v179 = F_contain_volatile_functions(m, v163)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v179 != 0 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v181 = F_get_commutator(m, v170)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v181 != 0 {
		v193 = v169
		v194 = v181
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v402 = v122
	goto L20
L41:
	;
	if v184 == int32(0) {
		v402 = v122
		goto L20
	} else {
		goto L42
	}
L42:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v132)+48))
	v189 = F_bms_is_member(m, v88, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v189 != 0 {
		v402 = v122
		goto L20
	} else {
		goto L44
	}
L44:
	;
	v191 = F_contain_volatile_functions(m, v169)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v191 != 0 {
		v402 = v122
		goto L20
	} else {
		goto L46
	}
L46:
	;
	v193 = v163
	v194 = v170
	goto L31
L47:
	;
	v198 = int32(0)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	if v200 <= v198 {
		v402 = v122
		goto L20
	} else {
		goto L48
	}
L48:
	;
	v207 = v198
	v209 = v200
	v223 = v198
	v225 = v122
	goto L49
L49:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v195)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v223<<(uint(int32(2))%32))))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+110)))
	if v236 != int32(1) {
		v353 = v207
		v355 = v209
		v371 = v225
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v402 = v371
	goto L20
L51:
	;
	v378 = v223 + int32(1)
	if v378 < v355 {
		v207 = v353
		v209 = v355
		v223 = v378
		v225 = v371
		goto L49
	} else {
		goto L66
	}
L52:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+107)))
	if v239 != int32(1) {
		v353 = v207
		v355 = v209
		v371 = v225
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v242 = int32(0)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v235)+40))
	if v242 < v243 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if int32(0) <= v343 {
		v402 = v337
		goto L20
	} else {
		goto L65
	}
L55:
	;
	v252 = v242
	goto L58
L56:
	;
	goto L57
L57:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v337 = v225
	v343 = v314
	goto L54
L58:
	;
	v274 = F_match_index_to_operand(m, v193, v252, v235)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	goto L57
L60:
	;
	if v274 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137)+8)) = v194
	*(*int32)(unsafe.Add(mBase, uint32(v137)+4)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v207
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v147)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v137)+12)) = v279
	v337 = int32(1)
	v343 = v207
	goto L54
L62:
	;
	goto L63
L63:
	;
	v283 = v252 + int32(1)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v235)+40))
	if v283 < v284 {
		v252 = v283
		goto L58
	} else {
		goto L64
	}
L64:
	;
	goto L59
L65:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	v353 = v207 + int32(1)
	v355 = v348
	v371 = v337
	goto L51
L66:
	;
	goto L50
L67:
	;
	goto L19
L68:
	;
	F_pg_qsort(m, v92, v89, int32(24), int32(822))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if int32(2) <= v89 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v424 = int32(1)
	goto L73
L71:
	;
	goto L72
L72:
	;
	F_pg_qsort(m, v92, v89, int32(24), int32(823))
	mBase = m.M
	v513 = m.ExcPending
	if v513 != 0 {
		goto L1
	} else {
		goto L81
	}
L73:
	;
	v449 = int32(24)
	v451 = v92 + v424*v449
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v451-v449)))
	if v452 != v455 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L72
L75:
	;
	v480 = v424 + int32(1)
	if v480 != v89 {
		v424 = v480
		goto L73
	} else {
		goto L80
	}
L76:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v451)+4))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v451-int32(20))))
	if v457 != v460 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v451)+8))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v451-int32(16))))
	if base.B2i32(v452 == int32(-1))|base.B2i32(v464 != v467) != 0 {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v451)+12))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v451-int32(12))))
	if v470 != v473 {
		goto L75
	} else {
		goto L79
	}
L79:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v451-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v451)+20)) = v477
	goto L75
L80:
	;
	goto L74
L81:
	;
	if v89 <= int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v751 = int32(0)
	v752 = v92
	goto L11
L83:
	;
	goto L84
L84:
	;
	v517 = int32(0)
	v523 = v517
	v524 = int32(1)
	v525 = v517
	goto L85
L85:
	;
	if v523 < int32(0) {
		v715 = v523
		v717 = v525
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v751 = v717
	v752 = v92
	goto L11
L87:
	;
	v741 = v524 + int32(1)
	if v741 <= v89 {
		v523 = v715
		v524 = v741
		v525 = v717
		goto L85
	} else {
		goto L114
	}
L88:
	;
	if v524 == v89 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	if v524-v523 != int32(1) {
		goto L96
	} else {
		goto L97
	}
L90:
	;
	v551 = int32(24)
	v553 = v92 + v524*v551
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)))
	v557 = v92 + v523*v551
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v557)))
	if v554 != v558 {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v557)+4))
	if v560 != v561 {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v557)+8))
	if base.B2i32(v554 == int32(-1))|base.B2i32(v565 != v566) != 0 {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v553)+12))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v557)+12))
	if v569 == v570 {
		v715 = v523
		v717 = v525
		goto L87
	} else {
		goto L94
	}
L94:
	;
	goto L89
L95:
	;
	v710 = F_lappend(m, v525, v709)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L113
	}
L96:
	;
	v578 = int32(0)
	if v523 < v524 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L98
L98:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v92+v523*int32(24))+16))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v672+v676<<(uint(int32(2))%32))))
	v709 = v680
	goto L95
L99:
	;
	v584 = v523
	v587 = v578
	v594 = v578
	goto L102
L100:
	;
	v636 = v578
	v643 = v578
	goto L101
L101:
	;
	v658 = F_make_orclause(m, v636)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L110
	}
L102:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v81)+12))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v92+v584*int32(24))+16))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v609+v613<<(uint(int32(2))%32))))
	v618 = F_lappend(m, v594, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L104
	}
L103:
	;
	v636 = v625
	v643 = v618
	goto L101
L104:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v617)))
	if v620 == int32(318) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v617)+4))
	v624 = v623
	goto L107
L106:
	;
	v624 = v617
	goto L107
L107:
	;
	v625 = F_lappend(m, v587, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v628 = v584 + int32(1)
	if v628 != v524 {
		v584 = v628
		v587 = v625
		v594 = v618
		goto L102
	} else {
		goto L109
	}
L109:
	;
	goto L103
L110:
	;
	v660 = F_make_orclause(m, v643)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+8)))
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+11)))
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+12)))
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+10)))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v74)+32))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v74)+36))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v74)+40))
	v670 = F_make_plain_restrictinfo(m, v42, v658, v660, v662, v663, v664, v665, v666, v667, v668, v669)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v709 = v670
	goto L95
L113:
	;
	v715 = v524
	v717 = v710
	goto L87
L114:
	;
	goto L86
L115:
	;
	v773 = int32(0)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v74)+52))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v775)+8))
	if v776 != v751 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v778 = F_list_copy(m, v56)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L119
	}
L117:
	;
	v782 = v773
	goto L118
L118:
	;
	if v751 == int32(0) {
		v1051 = v773
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v780 = F_list_delete(m, v778, v74)
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v782 = v780
	goto L118
L121:
	;
	if v782 != 0 {
		goto L179
	} else {
		goto L180
	}
L122:
	;
	v785 = int32(0)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v751)+4))
	if v786 <= v785 {
		v1051 = v773
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v793 = v785
	v796 = v773
	goto L124
L124:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v751)+12))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v817+v793<<(uint(int32(2))%32))))
	if v821 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L125:
	;
	v1051 = v1039
	goto L121
L126:
	;
	v1041 = v793 + int32(1)
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v751)+4))
	if v1041 < v1042 {
		v793 = v1041
		v796 = v1039
		goto L124
	} else {
		goto L178
	}
L127:
	;
	v1009 = F_list_concat(m, v796, v986)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L1
	} else {
		goto L177
	}
L128:
	;
	if v973 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L129:
	;
	v835 = *(*int32)(unsafe.Add(mBase, uint32(v821)+52))
	goto L136
L130:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v821)))
	if v824 != int32(21) {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v821)+4))
	if v827 != 0 {
		goto L129
	} else {
		goto L132
	}
L132:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v821)+8))
	v829 = F_build_paths_for_OR(m, v42, v43, v828, v56)
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v831 = F_generate_bitmap_or_paths(m, v42, v43, v828, v56)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v833 = F_list_concat(m, v829, v831)
	mBase = m.M
	v834 = m.ExcPending
	if v834 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v973 = v833
	goto L128
L136:
	;
	if v835 != int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(v821)+52))
	v839 = *(*int32)(unsafe.Add(mBase, uint32(v838)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v821
	*(*int32)(unsafe.Add(mBase, uint32(v53)+28)) = v821
	v845 = F_list_make1_impl(m, int32(1), v53+int32(8))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L1
	} else {
		goto L141
	}
L138:
	;
	goto L139
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v821
	*(*int32)(unsafe.Add(mBase, uint32(v53)+16)) = v821
	v968 = F_list_make1_impl(m, int32(1), v53+int32(12))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L170
	}
L140:
	;
	if v839 == int32(0) {
		goto L150
	} else {
		goto L151
	}
L141:
	;
	v847 = F_build_paths_for_OR(m, v42, v43, v845, v782)
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v847 == int32(0) {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v864 = int32(0)
	v865 = float64(0)
	goto L140
L144:
	;
	goto L145
L145:
	;
	v853 = F_choose_bitmap_and(m, v42, v43, v847)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v855 = *(*float64)(unsafe.Add(mBase, uint32(v853)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v853
	*(*int32)(unsafe.Add(mBase, uint32(v53)+24)) = v853
	v861 = F_list_make1_impl(m, int32(1), v53+int32(4))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	if v782 != 0 {
		v864 = v861
		v865 = v855
		goto L140
	} else {
		goto L148
	}
L148:
	;
	if v861 != 0 {
		v986 = v861
		goto L127
	} else {
		goto L149
	}
L149:
	;
	v864 = v861
	v865 = v855
	goto L140
L150:
	;
	if v864 != 0 {
		v986 = v864
		goto L127
	} else {
		goto L169
	}
L151:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v839)+4))
	if v868 <= int32(0) {
		goto L150
	} else {
		goto L152
	}
L152:
	;
	v871 = int32(0)
	v877 = v871
	v883 = v871
	v898 = float64(0)
	goto L153
L153:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v839)+12))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v902+v877<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v906
	*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v906
	v910 = F_list_make1_impl(m, int32(1), v53)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L1
	} else {
		goto L155
	}
L154:
	;
	if v920 == int32(0) {
		goto L150
	} else {
		goto L161
	}
L155:
	;
	v912 = F_build_paths_for_OR(m, v42, v43, v910, v782)
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	if v912 == int32(0) {
		goto L150
	} else {
		goto L157
	}
L157:
	;
	v916 = F_choose_bitmap_and(m, v42, v43, v912)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	v918 = *(*float64)(unsafe.Add(mBase, uint32(v916)+56))
	v919 = base.F64_add(v898, v918)
	v920 = F_lappend(m, v883, v916)
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	v923 = v877 + int32(1)
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v839)+4))
	if v923 < v924 {
		v877 = v923
		v883 = v920
		v898 = v919
		goto L153
	} else {
		goto L160
	}
L160:
	;
	goto L154
L161:
	;
	if v864 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v930 = F_list_concat(m, v796, v920)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L1
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	if base.F64_gt(v919, v865) != 0 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	v1039 = v930
	goto L126
L166:
	;
	v933 = v864
	goto L168
L167:
	;
	v933 = v920
	goto L168
L168:
	;
	v986 = v933
	goto L127
L169:
	;
	v1051 = int32(0)
	goto L121
L170:
	;
	v970 = F_build_paths_for_OR(m, v42, v43, v968, v56)
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	v973 = v970
	goto L128
L172:
	;
	v1051 = int32(0)
	goto L121
L173:
	;
	goto L174
L174:
	;
	v977 = F_choose_bitmap_and(m, v42, v43, v973)
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v979 = F_lappend(m, v796, v977)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v1039 = v979
	goto L126
L177:
	;
	v1039 = v1009
	goto L126
L178:
	;
	goto L125
L179:
	;
	F_list_free(m, v782)
	mBase = m.M
	v1073 = m.ExcPending
	if v1073 != 0 {
		goto L1
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	if v1051 == int32(0) {
		v1366 = v42
		v1367 = v43
		v1368 = v44
		v1377 = v53
		v1380 = v56
		v1389 = v65
		goto L8
	} else {
		goto L183
	}
L182:
	;
	goto L181
L183:
	;
	v1076 = int32(0)
	v1079 = F_palloc0(m, int32(88))
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+8)) = v43
	*(*int64)(unsafe.Add(mBase, uint32(v1079))) = int64(1451698946332)
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v43)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+12)) = v1084
	if v1051 == int32(0) {
		v1137 = v1076
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v1162 = F_get_baserel_parampathinfo(m, v42, v43, v1137)
	mBase = m.M
	v1163 = m.ExcPending
	if v1163 != 0 {
		goto L1
	} else {
		goto L195
	}
L186:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	if v1088 <= int32(0) {
		v1137 = v1076
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v1094 = v1076
	v1095 = v1076
	goto L188
L188:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+12))
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v1119+v1095<<(uint(int32(2))%32))))
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1123)+16))
	if v1124 != 0 {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v1137 = v1128
	goto L185
L190:
	;
	v1125 = *(*int32)(unsafe.Add(mBase, uint32(v1124)+4))
	v1127 = v1125
	goto L192
L191:
	;
	v1127 = int32(0)
	goto L192
L192:
	;
	v1128 = F_bms_add_members(m, v1094, v1127)
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	v1131 = v1095 + int32(1)
	v1132 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	if v1131 < v1132 {
		v1094 = v1128
		v1095 = v1131
		goto L188
	} else {
		goto L194
	}
L194:
	;
	goto L189
L195:
	;
	v1164 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+20)) = uint8(v1164)
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+16)) = v1162
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+26)))
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+72)) = v1051
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+64)) = v1164
	*(*int32)(unsafe.Add(mBase, uint32(v1079)+24)) = v1164
	*(*uint8)(unsafe.Add(mBase, uint32(v1079)+21)) = uint8(v1167)
	v1174 = float64(0)
	v1176 = m.G0
	v1178 = v1176 - int32(16)
	m.G0 = v1178
	if v1051 == v1164 {
		v1306 = v1174
		v1307 = v1174
		goto L198
	} else {
		goto L199
	}
L196:
	;
	v1364 = F_lappend(m, v65, v1079)
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L1
	} else {
		goto L224
	}
L197:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L1
	} else {
		goto L221
	}
L198:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1079)+56)) = v1306
	*(*float64)(unsafe.Add(mBase, uint32(v1079)+48)) = v1306
	*(*int64)(unsafe.Add(mBase, uint32(v1079)+32)) = int64(0)
	v1314 = float64(1)
	if base.F64_lt(v1307, v1314) != 0 {
		goto L218
	} else {
		goto L219
	}
L199:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	if v1183 <= int32(0) {
		v1306 = v1174
		v1307 = v1174
		goto L198
	} else {
		goto L200
	}
L200:
	;
	v1187 = *(*float64)(unsafe.Add(mBase, _c_F_generate_bitmap_or_paths[0]))
	v1189 = base.F64_mul(v1187, float64(0.1))
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+12))
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	v1192 = *(*int32)(unsafe.Add(mBase, uint32(v1191)))
	if base.Ui32(int32(2)) <= base.Ui32(v1192-int32(283)) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	v1210 = float64(0)
	v1211 = base.F64_add(v1209, v1210)
	v1212 = *(*float64)(unsafe.Add(mBase, uint32(v1208)))
	v1214 = base.F64_add(v1212, v1210)
	v1215 = int32(1)
	if v1183 == v1215 {
		v1306 = v1211
		v1307 = v1214
		goto L198
	} else {
		goto L206
	}
L202:
	;
	if v1192 != int32(280) {
		v1326 = v1191
		goto L197
	} else {
		goto L205
	}
L203:
	;
	goto L204
L204:
	;
	v1207 = *(*float64)(unsafe.Add(mBase, uint32(v1191)+56))
	v1208 = v1191 + int32(80)
	v1209 = v1207
	goto L201
L205:
	;
	v1201 = *(*float64)(unsafe.Add(mBase, uint32(v1191)+32))
	v1203 = *(*float64)(unsafe.Add(mBase, uint32(v1191)+96))
	v1208 = v1191 + int32(104)
	v1209 = base.F64_add(base.F64_mul(v1189, v1201), v1203)
	goto L201
L206:
	;
	v1218 = int32(0)
	if v1218 < v1183 {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v1221 = v1183
	goto L209
L208:
	;
	v1221 = v1218
	goto L209
L209:
	;
	v1232 = v1215
	v1248 = v1211
	v1249 = v1214
	goto L210
L210:
	;
	v1252 = int32(2)
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1190+v1232<<(uint(v1252)%32))))
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1255)))
	if base.Ui32(v1252) <= base.Ui32(v1256-int32(283)) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v1306 = v1276
	v1307 = v1278
	goto L198
L212:
	;
	v1277 = *(*float64)(unsafe.Add(mBase, uint32(v1275)))
	v1278 = base.F64_add(v1249, v1277)
	v1280 = v1232 + int32(1)
	if v1280 != v1221 {
		v1232 = v1280
		v1248 = v1276
		v1249 = v1278
		goto L210
	} else {
		goto L217
	}
L213:
	;
	if v1256 != int32(280) {
		v1326 = v1255
		goto L197
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1272 = *(*float64)(unsafe.Add(mBase, uint32(v1255)+56))
	v1275 = v1255 + int32(80)
	v1276 = base.F64_add(base.F64_mul(v1187, float64(100)), base.F64_add(v1248, v1272))
	goto L212
L216:
	;
	v1265 = *(*float64)(unsafe.Add(mBase, uint32(v1255)+32))
	v1267 = *(*float64)(unsafe.Add(mBase, uint32(v1255)+96))
	v1275 = v1255 + int32(104)
	v1276 = base.F64_add(v1248, base.F64_add(base.F64_mul(v1189, v1265), v1267))
	goto L212
L217:
	;
	goto L211
L218:
	;
	v1317 = v1307
	goto L220
L219:
	;
	v1317 = v1314
	goto L220
L220:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v1079)+80)) = v1317
	m.G0 = v1178 + int32(16)
	goto L196
L221:
	;
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1326)))
	*(*int32)(unsafe.Add(mBase, uint32(v1178))) = v1354
	F_errmsg_internal(m, int32(_a_F_generate_bitmap_or_paths_0), v1178)
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_generate_bitmap_or_paths_1), int32(1149), int32(_a_F_generate_bitmap_or_paths_2))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	v1366 = v42
	v1367 = v43
	v1368 = v44
	v1377 = v53
	v1380 = v56
	v1389 = v1364
	goto L8
L225:
	;
	goto L7
}
