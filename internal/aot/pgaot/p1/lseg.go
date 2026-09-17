package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_center(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v28 float64
	_ = v28
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v46 float64
	_ = v46
	var v55 float64
	_ = v55
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		v14 = base.F64_add(v12, v13)
		if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v19 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v12), v19)&base.F64_ne(base.F64_abs(v13), v19) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v32 = base.F64_mul(v14, float64(0.5))
				v33 = float64(0)
				if base.F64_eq(v32, v33)&base.F64_ne(v14, v33) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v8))) = v32
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
					v41 = base.F64_add(v39, v40)
					if base.F64_eq(base.F64_abs(v41), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v46 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v39), v46)&base.F64_ne(base.F64_abs(v40), v46) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v59 = base.F64_mul(v41, float64(0.5))
							v60 = float64(0)
							if base.F64_eq(v59, v60)&base.F64_ne(v41, v60) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v59
								return v8
							}
						}
					} else {
						v55 = base.F64_mul(v41, float64(0.5))
						if base.F64_eq(base.F64_abs(v55), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v59 = v55
							v60 = float64(0)
							if base.F64_eq(v59, v60)&base.F64_ne(v41, v60) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v59
								return v8
							}
						}
					}
				}
			}
		} else {
			v28 = base.F64_mul(v14, float64(0.5))
			if base.F64_eq(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v32 = v28
				v33 = float64(0)
				if base.F64_eq(v32, v33)&base.F64_ne(v14, v33) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v8))) = v32
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
					v41 = base.F64_add(v39, v40)
					if base.F64_eq(base.F64_abs(v41), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v46 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v39), v46)&base.F64_ne(base.F64_abs(v40), v46) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v59 = base.F64_mul(v41, float64(0.5))
							v60 = float64(0)
							if base.F64_eq(v59, v60)&base.F64_ne(v41, v60) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v59
								return v8
							}
						}
					} else {
						v55 = base.F64_mul(v41, float64(0.5))
						if base.F64_eq(base.F64_abs(v55), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v59 = v55
							v60 = float64(0)
							if base.F64_eq(v59, v60)&base.F64_ne(v41, v60) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v8)+8)) = v59
								return v8
							}
						}
					}
				}
			}
		}
	}
}
func F_lseg_construct(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v13 float64
	_ = v13
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(32))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
		*(*float64)(unsafe.Add(mBase, uint32(v7))) = v11
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v5)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v13
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
		*(*float64)(unsafe.Add(mBase, uint32(v7)+16)) = v15
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v7)+24)) = v17
		return v7
	}
}
func F_lseg_inside_poly(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v91 float64
	_ = v91
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v97 float64
	_ = v97
	var v99 int32
	_ = v99
	var v100 float64
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v172 float64
	_ = v172
	var v177 float64
	_ = v177
	var v186 float64
	_ = v186
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v197 float64
	_ = v197
	var v198 float64
	_ = v198
	var v199 float64
	_ = v199
	var v204 float64
	_ = v204
	var v213 float64
	_ = v213
	var v217 float64
	_ = v217
	var v218 float64
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	F_check_stack_depth(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v26
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v30
	v33 = l2 + int32(40)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = l3
	goto L5
L4:
	;
	v35 = v34
	goto L5
L5:
	;
	v40 = v33 + v35<<(uint(int32(4))%32) - int32(16)
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v43
	if v34 <= l3 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L62
	}
L7:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L61
	}
L8:
	;
	m.G0 = v18 + int32(80)
	return v229
L9:
	;
	v229 = int32(1)
	goto L8
L10:
	;
	if base.B2i32(v152 == int32(0))|v161 != 0 {
		v229 = v152
		goto L8
	} else {
		goto L45
	}
L11:
	;
	v152 = int32(1)
	v157 = v34
	v161 = v5
	goto L10
L12:
	;
	goto L13
L13:
	;
	v48 = v18 + int32(32)
	v50 = v18 - int32(-64)
	v54 = l3
	v60 = v5
	goto L14
L14:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_lseg_inside_poly[0]))
	if v67 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v152 = v142
	v157 = v150
	v161 = v145
	goto L10
L16:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v72 = v33 + v54<<(uint(int32(4))%32)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = v73
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	*(*int64)(unsafe.Add(mBase, uint32(v50))) = v75
	v78 = v18 + int32(16)
	v80 = v18 + int32(48)
	v81 = F_point_dt(m, v78, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v83 = F_point_dt(m, v78, v50)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v85 = F_point_dt(m, v80, v50)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v87 = F_point_dt(m, v48, v80)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v89 = F_point_dt(m, v48, v50)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v91 = base.F64_add(v87, v89)
	v92 = F_point_dt(m, v80, v50)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v97 = float64(1e-06)
	v99 = base.F64_eq(v91, v92) | base.F64_le(base.F64_abs(base.F64_sub(v91, v92)), v97)
	v100 = base.F64_add(v81, v83)
	v106 = int32(0)
	if base.F64_ne(v100, v85)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v100, v85)), v97) == v106) == v106 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v146
	v148 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v150 <= v143 {
		v152 = v142
		v157 = v150
		v161 = v145
		goto L10
	} else {
		goto L43
	}
L27:
	;
	if v99 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v99 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v112 = v54 + int32(1)
	v113 = F_touched_lseg_inside_poly(m, v78, v48, v80, l2, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v142 = v113
	v143 = v112
	v145 = v60
	goto L26
L32:
	;
	v120 = v54 + int32(1)
	v121 = F_touched_lseg_inside_poly(m, v48, v18+int32(16), v18+int32(48), l2, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v123 = int32(1)
	v125 = v54 + v123
	v127 = v18 + int32(16)
	v130 = F_lseg_interpt_lseg(m, v18, v127, v18+int32(48))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	v142 = v121
	v143 = v120
	v145 = v60
	goto L26
L36:
	;
	if v130 == int32(0) {
		v142 = v123
		v143 = v125
		v145 = v60
		goto L26
	} else {
		goto L37
	}
L37:
	;
	v134 = F_lseg_inside_poly(m, v127, v18, l2, v125)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v134 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v229 = int32(0)
	goto L8
L40:
	;
	goto L41
L41:
	;
	v140 = F_lseg_inside_poly(m, v48, v18, l2, v125)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v142 = v140
	v143 = v125
	v145 = int32(1)
	goto L26
L43:
	;
	if v142 != 0 {
		v54 = v143
		v60 = v145
		goto L14
	} else {
		goto L44
	}
L44:
	;
	goto L15
L45:
	;
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
	v171 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
	v172 = base.F64_add(v170, v171)
	if base.F64_eq(base.F64_abs(v172), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v191 = float64(0)
	if base.F64_eq(v190, v191)&base.F64_ne(v172, v191) != 0 {
		goto L6
	} else {
		goto L52
	}
L47:
	;
	v177 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v170), v177)&base.F64_ne(base.F64_abs(v171), v177) != 0 {
		goto L7
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v186 = base.F64_mul(v172, float64(0.5))
	if base.F64_eq(base.F64_abs(v186), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L51
	}
L50:
	;
	v190 = base.F64_mul(v172, float64(0.5))
	goto L46
L51:
	;
	v190 = v186
	goto L46
L52:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v18))) = v190
	v197 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	v198 = *(*float64)(unsafe.Add(mBase, uint32(v18)+40))
	v199 = base.F64_add(v197, v198)
	if base.F64_eq(base.F64_abs(v199), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v218 = float64(0)
	if base.F64_eq(v217, v218)&base.F64_ne(v199, v218) != 0 {
		goto L6
	} else {
		goto L59
	}
L54:
	;
	v204 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v197), v204)&base.F64_ne(base.F64_abs(v198), v204) != 0 {
		goto L7
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v213 = base.F64_mul(v199, float64(0.5))
	if base.F64_eq(base.F64_abs(v213), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L58
	}
L57:
	;
	v217 = base.F64_mul(v199, float64(0.5))
	goto L53
L58:
	;
	v217 = v213
	goto L53
L59:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v18)+8)) = v217
	v224 = F_point_inside(m, v18, v157, v33)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v229 = base.B2i32(v224 != int32(0))
	goto L8
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lseg_parallel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_point_sl(m, v6, v6+int32(16))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v15 = F_point_sl(m, v5, v5+int32(16))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return base.F64_eq(v9, v15) | base.F64_le(base.F64_abs(base.F64_sub(v9, v15)), float64(1e-06))
		}
	}
}
