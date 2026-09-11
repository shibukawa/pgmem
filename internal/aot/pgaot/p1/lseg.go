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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v45 int64
	_ = v45
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v85 float64
	_ = v85
	var v86 int32
	_ = v86
	var v89 float64
	_ = v89
	var v90 int32
	_ = v90
	var v93 float64
	_ = v93
	var v94 int32
	_ = v94
	var v97 float64
	_ = v97
	var v98 int32
	_ = v98
	var v99 float64
	_ = v99
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v182 float64
	_ = v182
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v189 float64
	_ = v189
	var v198 float64
	_ = v198
	var v202 float64
	_ = v202
	var v203 float64
	_ = v203
	var v209 float64
	_ = v209
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v216 float64
	_ = v216
	var v225 float64
	_ = v225
	var v229 float64
	_ = v229
	var v230 float64
	_ = v230
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
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
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v28
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v30
	v33 = v18 + int32(56)
	v35 = l2 + int32(40)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v37 = l3
	goto L5
L4:
	;
	v37 = v36
	goto L5
L5:
	;
	v42 = v35 + v37<<(uint(int32(4))%32) - int32(16)
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v45
	if v36 <= l3 {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L64
	}
L7:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L63
	}
L8:
	;
	m.G0 = v18 + int32(80)
	return v245
L9:
	;
	if v170 == int32(0) {
		v245 = v170
		goto L8
	} else {
		goto L46
	}
L10:
	;
	v170 = int32(1)
	v171 = v36
	v172 = v5
	goto L9
L11:
	;
	goto L12
L12:
	;
	v50 = v18 + int32(32)
	v52 = v18 - int32(-64)
	v56 = l3
	v60 = v5
	goto L13
L13:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_lseg_inside_poly[0]))
	if v69 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v170 = v157
	v171 = v163
	v172 = v158
	goto L9
L15:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v74 = v35 + v56<<(uint(int32(4))%32)
	v75 = *(*int64)(unsafe.Add(mBase, uint32(v74)))
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = v75
	v78 = v52 + int32(8)
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v74)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v78))) = v79
	v85 = F_point_dt(m, v18+int32(16), v18+int32(48))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L17
L19:
	;
	v89 = F_point_dt(m, v18+int32(16), v52)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v93 = F_point_dt(m, v18+int32(48), v52)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v97 = F_point_dt(m, v50, v18+int32(48))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v99 = F_point_dt(m, v50, v52)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v101 = base.F64_add(v97, v99)
	v104 = F_point_dt(m, v18+int32(48), v52)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v111 = base.F64_eq(v101, v104) | base.F64_le(base.F64_abs(base.F64_sub(v101, v104)), float64(1e-06))
	v112 = base.F64_add(v85, v89)
	if base.F64_ne(v112, v93) != 0 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v159 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	*(*int64)(unsafe.Add(mBase, uint32(v33))) = v159
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v52)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v163 <= v156 {
		v170 = v157
		v171 = v163
		v172 = v158
		goto L9
	} else {
		goto L44
	}
L26:
	;
	v154 = F_lseg_inside_poly(m, v50, v18, l2, v138)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L27:
	;
	v245 = int32(1)
	goto L8
L28:
	;
	if v111 != 0 {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v112, v93)), float64(1e-06)) == int32(0) {
		goto L28
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v111 != 0 {
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v125 = v56 + int32(1)
	v126 = F_touched_lseg_inside_poly(m, v18+int32(16), v50, v18+int32(48), l2, v125)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v156 = v125
	v157 = v126
	v158 = v60
	goto L25
L35:
	;
	v133 = v56 + int32(1)
	v134 = F_touched_lseg_inside_poly(m, v50, v18+int32(16), v18+int32(48), l2, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v136 = int32(1)
	v138 = v56 + v136
	v143 = F_lseg_interpt_lseg(m, v18, v18+int32(16), v18+int32(48))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v156 = v133
	v157 = v134
	v158 = v60
	goto L25
L39:
	;
	if v143 == int32(0) {
		v156 = v138
		v157 = v136
		v158 = v60
		goto L25
	} else {
		goto L40
	}
L40:
	;
	v149 = F_lseg_inside_poly(m, v18+int32(16), v18, l2, v138)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	if v149 != 0 {
		goto L26
	} else {
		goto L42
	}
L42:
	;
	v245 = int32(0)
	goto L8
L43:
	;
	v156 = v138
	v157 = v154
	v158 = int32(1)
	goto L25
L44:
	;
	if v157 != 0 {
		v56 = v156
		v60 = v158
		goto L13
	} else {
		goto L45
	}
L45:
	;
	goto L14
L46:
	;
	if v172 != 0 {
		v245 = v170
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v182 = *(*float64)(unsafe.Add(mBase, uint32(v18)+16))
	v183 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
	v184 = base.F64_add(v182, v183)
	if base.F64_eq(base.F64_abs(v184), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v203 = float64(0)
	if base.F64_eq(v202, v203)&base.F64_ne(v184, v203) != 0 {
		goto L6
	} else {
		goto L54
	}
L49:
	;
	v189 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v182), v189)&base.F64_ne(base.F64_abs(v183), v189) != 0 {
		goto L7
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v198 = base.F64_mul(v184, float64(0.5))
	if base.F64_eq(base.F64_abs(v198), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L53
	}
L52:
	;
	v202 = base.F64_mul(v184, float64(0.5))
	goto L48
L53:
	;
	v202 = v198
	goto L48
L54:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v18))) = v202
	v209 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
	v210 = *(*float64)(unsafe.Add(mBase, uint32(v18)+40))
	v211 = base.F64_add(v209, v210)
	if base.F64_eq(base.F64_abs(v211), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v230 = float64(0)
	if base.F64_eq(v229, v230)&base.F64_ne(v211, v230) != 0 {
		goto L6
	} else {
		goto L61
	}
L56:
	;
	v216 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v209), v216)&base.F64_ne(base.F64_abs(v210), v216) != 0 {
		goto L7
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v225 = base.F64_mul(v211, float64(0.5))
	if base.F64_eq(base.F64_abs(v225), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L60
	}
L59:
	;
	v229 = base.F64_mul(v211, float64(0.5))
	goto L55
L60:
	;
	v229 = v225
	goto L55
L61:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v18)+8)) = v229
	v236 = F_point_inside(m, v18, v171, v35)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v245 = base.B2i32(v236 != int32(0))
	goto L8
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
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
