package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_Float8GetDatum(m *base.Module, l0 float64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_palloc(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v4))) = l0
		return v4
	}
}
func F_float8_corr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 float64
	_ = v25
	var v28 int32
	_ = v28
	var v31 float64
	_ = v31
	var v34 float64
	_ = v34
	var v38 int32
	_ = v38
	var v43 float64
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(197990)
				F_errmsg_internal(m, int32(24733), v8)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474438), int32(2938), int32(23273))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			if v18 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(197990)
					F_errmsg_internal(m, int32(24733), v8)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(474438), int32(2938), int32(23273))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(197990)
						F_errmsg_internal(m, int32(24733), v8)
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(474438), int32(2938), int32(23273))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
					if v22 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(197990)
							F_errmsg_internal(m, int32(24733), v8)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(474438), int32(2938), int32(23273))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
						if base.F64_lt(v25, float64(1)) != 0 {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							v51 = int32(0)
							m.G0 = v8 + int32(16)
							return v51
						} else {
							v31 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
							if base.F64_ne(v31, float64(0)) != 0 {
								v34 = *(*float64)(unsafe.Add(mBase, uint32(v11)+56))
								if base.F64_ne(v34, float64(0)) != 0 {
									v43 = *(*float64)(unsafe.Add(mBase, uint32(v11-int32(-64))))
									v47 = F_Float8GetDatum(m, base.F64_div(v43, base.F64_sqrt(base.F64_mul(v31, v34))))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										v51 = v47
										m.G0 = v8 + int32(16)
										return v51
									}
								} else {
									v38 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
									v51 = int32(0)
									m.G0 = v8 + int32(16)
									return v51
								}
							} else {
								v38 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
								v51 = int32(0)
								m.G0 = v8 + int32(16)
								return v51
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_mul(m *base.Module, l0 float64, l1 float64) float64 {
	var v4 float64
	_ = v4
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v4 = base.F64_mul(l0, l1)
	if base.F64_ne(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		if base.F64_ne(v4, float64(0)) != 0 {
			return v4
		} else {
			if base.F64_eq(l0, float64(0)) != 0 {
				return v4
			} else {
				if base.F64_ne(l1, float64(0)) != 0 {
					F_float_underflow_error(m)
					v26 = m.ExcPending
					if v26 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					return v4
				}
			}
		}
	} else {
		if base.F64_eq(base.F64_abs(l0), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			if base.F64_ne(v4, float64(0)) != 0 {
				return v4
			} else {
				if base.F64_eq(l0, float64(0)) != 0 {
					return v4
				} else {
					if base.F64_ne(l1, float64(0)) != 0 {
						F_float_underflow_error(m)
						v26 = m.ExcPending
						if v26 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return v4
					}
				}
			}
		} else {
			if base.F64_ne(base.F64_abs(l1), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				v24 = m.ExcPending
				if v24 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_ne(v4, float64(0)) != 0 {
					return v4
				} else {
					if base.F64_eq(l0, float64(0)) != 0 {
						return v4
					} else {
						if base.F64_ne(l1, float64(0)) != 0 {
							F_float_underflow_error(m)
							v26 = m.ExcPending
							if v26 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							return v4
						}
					}
				}
			}
		}
	}
}
func F_float8_qsort_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v11 int32
	_ = v11
	v7 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_lt(v7, v8) != 0 {
		v11 = int32(-1)
	} else {
		v11 = base.F64_ne(v7, v8)
	}
	return v11
}
func F_float8_regr_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v61 float64
	_ = v61
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v70 float64
	_ = v70
	var v73 float64
	_ = v73
	var v77 float64
	_ = v77
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v99 float64
	_ = v99
	var v106 float64
	_ = v106
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v120 float64
	_ = v120
	var v141 float64
	_ = v141
	var v147 float64
	_ = v147
	var v152 float64
	_ = v152
	var v161 int64
	_ = v161
	var v165 float64
	_ = v165
	var v167 float64
	_ = v167
	var v168 float64
	_ = v168
	var v169 float64
	_ = v169
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v187 int32
	_ = v187
	var v192 float64
	_ = v192
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v227 int32
	_ = v227
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	v17 = m.G0
	v19 = v17 - int32(96)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L3
	} else {
		goto L68
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L3
	} else {
		goto L65
	}
L3:
	;
	return int32(0)
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v26 != int32(1) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v29 != int32(6) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v32 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	if v33 != int32(701) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v41 = *(*float64)(unsafe.Add(mBase, uint32(v22)+24))
	v42 = *(*float64)(unsafe.Add(mBase, uint32(v22)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+72)) = v42
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v22)+48))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v22)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = v45
	v49 = *(*float64)(unsafe.Add(mBase, uint32(v22-int32(-64))))
	v51 = base.F64_add(v41, float64(1))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+88)) = v51
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v49
	v54 = base.F64_add(v37, v40)
	*(*float64)(unsafe.Add(mBase, uint32(v19)+80)) = v54
	v56 = base.F64_add(v39, v44)
	*(*float64)(unsafe.Add(mBase, uint32(v19)+64)) = v56
	if base.F64_gt(v41, float64(0)) != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v199 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v187))) = int64(9221120237041090560)
	v192 = math.Float64frombits(uint64(0x7ff8000000000000))
	v193 = v183
	v194 = v184
	goto L9
L11:
	;
	v61 = base.F64_sub(base.F64_mul(v39, v51), v56)
	v65 = base.F64_div(float64(1), base.F64_mul(v51, v41))
	v67 = base.F64_add(base.F64_mul(base.F64_mul(v61, v61), v65), v45)
	*(*float64)(unsafe.Add(mBase, uint32(v19)+56)) = v67
	v70 = base.F64_sub(base.F64_mul(v37, v51), v54)
	v73 = base.F64_add(base.F64_mul(base.F64_mul(v70, v70), v65), v42)
	*(*float64)(unsafe.Add(mBase, uint32(v19)+72)) = v73
	v77 = base.F64_add(base.F64_mul(base.F64_mul(v70, v61), v65), v49)
	*(*float64)(unsafe.Add(mBase, uint32(v19)+48)) = v77
	if base.F64_ne(base.F64_abs(v54), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	goto L13
L13:
	;
	v152 = base.F64_abs(v37)
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v152)) <= base.Ui64(int64(9218868437227405312)))&base.F64_ne(v152, math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L14:
	;
	v108 = math.Float64frombits(uint64(0x7ff0000000000000))
	v110 = base.F64_abs(v67)
	if base.F64_ne(v106, v108)&base.F64_ne(v110, v108) != 0 {
		goto L25
	} else {
		goto L26
	}
L15:
	;
	v82 = base.F64_abs(v56)
	v83 = base.F64_abs(v73)
	if base.F64_eq(v83, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v99 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v40), v99)&base.F64_ne(base.F64_abs(v37), v99) != 0 {
		goto L1
	} else {
		goto L24
	}
L18:
	;
	if base.F64_ne(v83, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v106 = v82
		goto L14
	} else {
		goto L23
	}
L19:
	;
	if base.F64_eq(v82, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if base.F64_eq(base.F64_abs(v67), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if base.F64_ne(base.F64_abs(v77), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v192 = v77
		v193 = v73
		v194 = v67
		goto L9
	} else {
		goto L22
	}
L22:
	;
	goto L18
L23:
	;
	goto L17
L24:
	;
	v106 = base.F64_abs(v56)
	goto L14
L25:
	;
	v120 = base.F64_abs(v77)
	if base.F64_ne(v120, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	if base.F64_eq(base.F64_abs(v44), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	if base.F64_ne(base.F64_abs(v39), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	if base.F64_eq(base.F64_abs(v73), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L35
	} else {
		goto L36
	}
L30:
	;
	if base.F64_eq(base.F64_abs(v40), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if base.F64_eq(base.F64_abs(v37), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if base.F64_eq(base.F64_abs(v44), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L29
	} else {
		goto L33
	}
L33:
	;
	if base.F64_ne(base.F64_abs(v39), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L29
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = int64(9221120237041090560)
	v141 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L37
L36:
	;
	v141 = v73
	goto L37
L37:
	;
	if base.F64_eq(v110, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = int64(9221120237041090560)
	v147 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L40
L39:
	;
	v147 = v67
	goto L40
L40:
	;
	if base.F64_eq(v120, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v183 = v141
		v184 = v147
		v187 = v19 + int32(48)
		goto L10
	} else {
		goto L41
	}
L41:
	;
	v192 = v77
	v193 = v141
	v194 = v147
	goto L9
L42:
	;
	v161 = int64(9221120237041090560)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v161
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v161
	v165 = math.Float64frombits(uint64(0x7ff8000000000000))
	v167 = v165
	v168 = v165
	goto L44
L43:
	;
	v167 = v49
	v168 = v42
	goto L44
L44:
	;
	v169 = base.F64_abs(v39)
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v169)) <= base.Ui64(int64(9218868437227405312)))&base.F64_ne(v169, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v192 = v167
		v193 = v168
		v194 = v45
		goto L9
	} else {
		goto L45
	}
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = int64(9221120237041090560)
	v183 = v168
	v184 = math.Float64frombits(uint64(0x7ff8000000000000))
	v187 = v19 + int32(56)
	goto L10
L46:
	;
	m.G0 = v19 + int32(96)
	return v260
L47:
	;
	if v227 != 0 {
		goto L61
	} else {
		goto L62
	}
L48:
	;
	v227 = int32(0)
	goto L47
L50:
	;
	goto L48
L51:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	switch v202 - int32(429) {
	case 0:
		goto L53
	case 1:
		goto L52
	default:
		goto L50
	}
L52:
	;
	goto L57
L53:
	;
	goto L54
L54:
	;
	v227 = int32(1)
	goto L47
L57:
	;
	v227 = int32(2)
	goto L47
L61:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v22-int32(-64)))) = v192
	*(*float64)(unsafe.Add(mBase, uint32(v22)+56)) = v194
	*(*float64)(unsafe.Add(mBase, uint32(v22)+48)) = v56
	*(*float64)(unsafe.Add(mBase, uint32(v22)+40)) = v193
	*(*float64)(unsafe.Add(mBase, uint32(v22)+32)) = v54
	*(*float64)(unsafe.Add(mBase, uint32(v22)+24)) = v51
	v260 = v22
	goto L46
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v19 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v19 + int32(56)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v19 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v19 + int32(72)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v19 + int32(80)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v19 + int32(88)
	v258 = F_construct_array_builtin(m, v19+int32(16), int32(6), int32(701))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v260 = v258
	goto L46
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(275420)
	F_errmsg_internal(m, int32(24733), v19)
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(474438), int32(2938), int32(23273))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_float8_regr_avgy(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(19272)
				F_errmsg_internal(m, int32(24733), v7)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474438), int32(2938), int32(23273))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(19272)
					F_errmsg_internal(m, int32(24733), v7)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(474438), int32(2938), int32(23273))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(19272)
						F_errmsg_internal(m, int32(24733), v7)
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(474438), int32(2938), int32(23273))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(19272)
							F_errmsg_internal(m, int32(24733), v7)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(474438), int32(2938), int32(23273))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_lt(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v34 = int32(0)
							m.G0 = v7 + int32(16)
							return v34
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+48))
							v32 = F_Float8GetDatum(m, base.F64_div(v30, v24))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								v34 = v32
								m.G0 = v7 + int32(16)
								return v34
							}
						}
					}
				}
			}
		}
	}
}
func F_float8_regr_slope(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 float64
	_ = v24
	var v27 int32
	_ = v27
	var v30 float64
	_ = v30
	var v33 int32
	_ = v33
	var v38 float64
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		if v14 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(357217)
				F_errmsg_internal(m, int32(24733), v7)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(474438), int32(2938), int32(23273))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			if v17 != int32(6) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(357217)
					F_errmsg_internal(m, int32(24733), v7)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(474438), int32(2938), int32(23273))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
				if v20 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(357217)
						F_errmsg_internal(m, int32(24733), v7)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(474438), int32(2938), int32(23273))
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
					if v21 != int32(701) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(357217)
							F_errmsg_internal(m, int32(24733), v7)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(474438), int32(2938), int32(23273))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*float64)(unsafe.Add(mBase, uint32(v10)+24))
						if base.F64_lt(v24, float64(1)) != 0 {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
							v43 = int32(0)
							m.G0 = v7 + int32(16)
							return v43
						} else {
							v30 = *(*float64)(unsafe.Add(mBase, uint32(v10)+40))
							if base.F64_eq(v30, float64(0)) != 0 {
								v33 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
								v43 = int32(0)
								m.G0 = v7 + int32(16)
								return v43
							} else {
								v38 = *(*float64)(unsafe.Add(mBase, uint32(v10-int32(-64))))
								v40 = F_Float8GetDatum(m, base.F64_div(v38, v30))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v43 = v40
									m.G0 = v7 + int32(16)
									return v43
								}
							}
						}
					}
				}
			}
		}
	}
}
