package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_above(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v43 int32
	_ = v43
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_sub(v9, v10)
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v23 = base.F64_add(v21, v22)
		if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
		} else {
			if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
			} else {
				if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
				}
			}
		}
	} else {
		if base.F64_eq(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v23 = base.F64_add(v21, v22)
			if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
			} else {
				if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
				} else {
					if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
					}
				}
			}
		} else {
			if base.F64_ne(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
				v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				v23 = base.F64_add(v21, v22)
				if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
				} else {
					if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
					} else {
						if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							return base.F64_gt(v11, base.F64_add(v23, float64(1e-06)))
						}
					}
				}
			}
		}
	}
}
func F_circle_area(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v15 float64
	_ = v15
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v28 float64
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)+16))
	v7 = base.F64_mul(v6, v6)
	v8 = base.F64_abs(v7)
	v9 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v8, v9)&base.F64_ne(base.F64_abs(v6), v9) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v15 = float64(0)
		if base.F64_eq(v7, v15)&base.F64_ne(v6, v15) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v21 = base.F64_mul(v7, float64(3.141592653589793))
			v23 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v21), v23)&base.F64_ne(v8, v23) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v28 = float64(0)
				if base.F64_eq(v21, v28)&base.F64_ne(v7, v28) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v33 = F_Float8GetDatum(m, v21)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v33
					}
				}
			}
		}
	}
}
func F_circle_center(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(16))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
		*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
		v11 = *(*float64)(unsafe.Add(mBase, uint32(v3)+8))
		*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v11
		return v5
	}
}
func F_circle_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = F_point_dt(m, v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v15 = base.F64_sub(v13, v14)
		if base.F64_ne(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			return base.F64_le(v9, base.F64_add(v15, float64(1e-06)))
		} else {
			if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_le(v9, base.F64_add(v15, float64(1e-06)))
			} else {
				if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_le(v9, base.F64_add(v15, float64(1e-06)))
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
func F_circle_diameter(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	var v8 float64
	_ = v8
	var v16 float64
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	v6 = base.F64_add(v5, v5)
	v8 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v6), v8)&base.F64_ne(base.F64_abs(v5), v8) == int32(0) {
		v16 = float64(0)
		if base.F64_eq(v6, v16)&base.F64_ne(v5, v16) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v21 = F_Float8GetDatum(m, v6)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v21
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_circle_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v26 float64
	_ = v26
	var v35 float64
	_ = v35
	var v38 float64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_point_dt(m, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
		v16 = base.F64_add(v14, v15)
		v17 = base.F64_abs(v16)
		if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v26 = base.F64_sub(v10, v16)
			if base.F64_ne(base.F64_abs(v26), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v35 = float64(0)
				if base.F64_lt(v26, v35) != 0 {
					v38 = v35
				} else {
					v38 = v26
				}
				v39 = F_Float8GetDatum(m, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					return v39
				}
			} else {
				if base.F64_eq(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v35 = float64(0)
					if base.F64_lt(v26, v35) != 0 {
						v38 = v35
					} else {
						v38 = v26
					}
					v39 = F_Float8GetDatum(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						return v39
					}
				} else {
					if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v35 = float64(0)
						if base.F64_lt(v26, v35) != 0 {
							v38 = v35
						} else {
							v38 = v26
						}
						v39 = F_Float8GetDatum(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							return v39
						}
					}
				}
			}
		} else {
			if base.F64_eq(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v26 = base.F64_sub(v10, v16)
				if base.F64_ne(base.F64_abs(v26), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v35 = float64(0)
					if base.F64_lt(v26, v35) != 0 {
						v38 = v35
					} else {
						v38 = v26
					}
					v39 = F_Float8GetDatum(m, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						return v39
					}
				} else {
					if base.F64_eq(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v35 = float64(0)
						if base.F64_lt(v26, v35) != 0 {
							v38 = v35
						} else {
							v38 = v26
						}
						v39 = F_Float8GetDatum(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							return v39
						}
					} else {
						if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v35 = float64(0)
							if base.F64_lt(v26, v35) != 0 {
								v38 = v35
							} else {
								v38 = v26
							}
							v39 = F_Float8GetDatum(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								return v39
							}
						}
					}
				}
			} else {
				if base.F64_ne(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v26 = base.F64_sub(v10, v16)
					if base.F64_ne(base.F64_abs(v26), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						v35 = float64(0)
						if base.F64_lt(v26, v35) != 0 {
							v38 = v35
						} else {
							v38 = v26
						}
						v39 = F_Float8GetDatum(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							return v39
						}
					} else {
						if base.F64_eq(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v35 = float64(0)
							if base.F64_lt(v26, v35) != 0 {
								v38 = v35
							} else {
								v38 = v26
							}
							v39 = F_Float8GetDatum(m, v38)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								return v39
							}
						} else {
							if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v35 = float64(0)
								if base.F64_lt(v26, v35) != 0 {
									v38 = v35
								} else {
									v38 = v26
								}
								v39 = F_Float8GetDatum(m, v38)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									return v39
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_circle_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_palloc(m, int32(24))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = int32(1)
	v25 = v16
	goto L6
L3:
	;
	v92 = F_pair_decode(m, v79, v18, v18+int32(8), v13+int32(44), int32(386859), v16, v15)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L15
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v69
	v79 = v69
	v84 = int32(0)
	goto L3
L5:
	;
	v69 = v25 + int32(1)
	goto L4
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v25
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if base.Ui32(v34-int32(9)) < base.Ui32(int32(5)) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v48 = v25
	goto L11
L8:
	;
	goto L7
L9:
	;
	v25 = v25 + int32(1)
	goto L6
L10:
	;
	switch v34 - int32(32) {
	case 0:
		goto L9
	default:
		v79 = v25
		v84 = v22
		goto L3
	case 8:
		goto L8
	case 28:
		goto L5
	}
L11:
	;
	v55 = v48 + int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if base.Ui32(v56-int32(9)) < base.Ui32(int32(5)) {
		v48 = v55
		goto L11
	} else {
		goto L13
	}
L13:
	;
	switch v56 - int32(32) {
	case 0:
		v48 = v55
		goto L11
	default:
		v79 = v25
		v84 = v22
		goto L3
	case 8:
		v69 = v55
		goto L4
	}
L14:
	;
	m.G0 = v13 + int32(48)
	return v232
L15:
	;
	if v92 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v96 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
	v232 = int32(0)
	goto L14
L17:
	;
	goto L18
L18:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v100 == int32(44) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v104 = v99 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v104
	v106 = v104
	goto L21
L20:
	;
	v106 = v99
	goto L21
L21:
	;
	v110 = F_float8in_internal(m, v106, v13+int32(44), int32(386859), v16, v15)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v18)+16)) = v110
	if v15 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if base.F64_lt(v110, float64(0)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L24:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v115 != int32(447) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)))
	if v118 != int32(1) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v121)
	v232 = int32(0)
	goto L14
L27:
	;
	if v198 == int32(0) {
		v232 = v18
		goto L14
	} else {
		goto L49
	}
L28:
	;
	v176 = v128
	goto L45
L29:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	if v84 != 0 {
		v198 = v129
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v155 = int32(0)
	v156 = F_errsave_start(m, v15)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L40
	}
L32:
	;
	if v129 == int32(62) {
		goto L28
	} else {
		goto L33
	}
L33:
	;
	if v129 == int32(41) {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v134 = int32(0)
	v135 = F_errsave_start(m, v15)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v135 == int32(0) {
		v232 = v134
		goto L14
	} else {
		goto L36
	}
L36:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(386859)
	F_errmsg(m, int32(704449), v13+int32(32))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errsave_finish(m, v15, int32(488830), int32(4666), int32(277060))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v232 = v134
	goto L14
L40:
	;
	if v156 == int32(0) {
		v232 = v155
		goto L14
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(386859)
	F_errmsg(m, int32(704449), v13)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errsave_finish(m, v15, int32(488830), int32(4651), int32(277060))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v232 = v155
	goto L14
L45:
	;
	v185 = v176 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v185
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
	if base.Ui32(v187-int32(9)) < base.Ui32(int32(5)) {
		v176 = v185
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v198 = v187
	goto L27
L47:
	;
	if v187 == int32(32) {
		v176 = v185
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v206 = int32(0)
	v207 = F_errsave_start(m, v15)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v207 == int32(0) {
		v232 = v206
		goto L14
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(386859)
	F_errmsg(m, int32(704449), v13+int32(16))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errsave_finish(m, v15, int32(488830), int32(4673), int32(277060))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v232 = v206
	goto L14
}
func F_circle_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v12 float64
	_ = v12
	var v18 float64
	_ = v18
	var v24 float64
	_ = v24
	var v26 float64
	_ = v26
	var v31 float64
	_ = v31
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v45 float64
	_ = v45
	var v51 float64
	_ = v51
	var v53 float64
	_ = v53
	var v58 float64
	_ = v58
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v10 = base.F64_mul(v9, v9)
	v11 = base.F64_abs(v10)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v11, v12)&base.F64_ne(base.F64_abs(v9), v12) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v18 = float64(0)
		if base.F64_eq(v10, v18)&base.F64_ne(v9, v18) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v24 = base.F64_mul(v10, float64(3.141592653589793))
			v26 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(base.F64_abs(v24), v26)&base.F64_ne(v11, v26) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v31 = float64(0)
				if base.F64_eq(v24, v31)&base.F64_ne(v10, v31) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v36 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
					v37 = base.F64_mul(v36, v36)
					v38 = base.F64_abs(v37)
					v39 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_eq(v38, v39)&base.F64_ne(base.F64_abs(v36), v39) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v45 = float64(0)
						if base.F64_eq(v37, v45)&base.F64_ne(v36, v45) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v51 = base.F64_mul(v37, float64(3.141592653589793))
							v53 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v51), v53)&base.F64_ne(v38, v53) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v58 = float64(0)
								if base.F64_eq(v51, v58)&base.F64_ne(v37, v58) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									return base.F64_ne(v51, v24) & base.F64_gt(base.F64_abs(base.F64_sub(v24, v51)), float64(1e-06))
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_circle_overleft(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v11 float64
	_ = v11
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v43 int32
	_ = v43
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_add(v9, v10)
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v23 = base.F64_add(v21, v22)
		if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
		} else {
			if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
			} else {
				if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
				}
			}
		}
	} else {
		if base.F64_eq(base.F64_abs(v9), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v23 = base.F64_add(v21, v22)
			if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
			} else {
				if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
				} else {
					if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
					}
				}
			}
		} else {
			if base.F64_ne(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
				v22 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
				v23 = base.F64_add(v21, v22)
				if base.F64_ne(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
				} else {
					if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
					} else {
						if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							return base.F64_le(v11, base.F64_add(v23, float64(1e-06)))
						}
					}
				}
			}
		}
	}
}
