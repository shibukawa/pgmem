package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_lseg_horizontal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v4)+24))
	return base.F64_eq(v5, v6) | base.F64_le(base.F64_abs(base.F64_sub(v5, v6)), float64(1e-06))
}
func F_lseg_interpt_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v21 int32
	_ = v21
	var v29 float64
	_ = v29
	var v37 float64
	_ = v37
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v52 float64
	_ = v52
	var v58 float64
	_ = v58
	var v60 float64
	_ = v60
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	v4 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = l2 + int32(16)
	v18 = F_point_sl(m, l2, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		if base.F64_eq(base.F64_abs(v18), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(-4616189618054758400)
			v29 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			*(*float64)(unsafe.Add(mBase, uint32(v14)+24)) = v29
			v80 = v14 + int32(32)
			v83 = F_lseg_interpt_line(m, v80, l1, v14+int32(8))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				if v83 == int32(0) {
					v109 = v4
					m.G0 = v14 + int32(48)
					return v109
				} else {
					v87 = F_point_dt(m, v80, l2)
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						v89 = F_point_dt(m, v80, v17)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v91 = base.F64_add(v87, v89)
							v92 = F_point_dt(m, l2, v17)
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								if base.F64_ne(v91, v92)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v92)), float64(1e-06)) == int32(0)) != 0 {
									v109 = v4
								} else {
									v102 = int32(1)
									if l0 == int32(0) {
										v109 = v102
									} else {
										v105 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
										v107 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
										v109 = v102
									}
								}
								m.G0 = v14 + int32(48)
								return v109
							}
						}
					}
				}
			}
		} else {
			if base.F64_eq(v18, float64(0)) != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(-4616189618054758400)
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = int64(0)
				v37 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v14)+24)) = v37
				v80 = v14 + int32(32)
				v83 = F_lseg_interpt_line(m, v80, l1, v14+int32(8))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					if v83 == int32(0) {
						v109 = v4
						m.G0 = v14 + int32(48)
						return v109
					} else {
						v87 = F_point_dt(m, v80, l2)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							v89 = F_point_dt(m, v80, v17)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								v91 = base.F64_add(v87, v89)
								v92 = F_point_dt(m, l2, v17)
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									if base.F64_ne(v91, v92)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v92)), float64(1e-06)) == int32(0)) != 0 {
										v109 = v4
									} else {
										v102 = int32(1)
										if l0 == int32(0) {
											v109 = v102
										} else {
											v105 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
											v107 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
											v109 = v102
										}
									}
									m.G0 = v14 + int32(48)
									return v109
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = int64(-4616189618054758400)
				*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v18
				v42 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v43 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
				v44 = base.F64_mul(v18, v43)
				v45 = base.F64_abs(v44)
				v46 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(v45, v46)&base.F64_ne(base.F64_abs(v43), v46) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v52 = float64(0)
					if base.F64_eq(v44, v52)&base.F64_ne(v43, v52) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v58 = math.Float64frombits(uint64(0x7ff0000000000000))
						v60 = base.F64_sub(v42, v44)
						if base.B2i32(base.F64_eq(base.F64_abs(v42), v58)|base.F64_ne(base.F64_abs(v60), v58) == int32(0))&base.F64_ne(v45, v58) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v14)+24)) = v60
							if base.F64_ne(v60, float64(0)) != 0 {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = int64(0)
							}
							v80 = v14 + int32(32)
							v83 = F_lseg_interpt_line(m, v80, l1, v14+int32(8))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								if v83 == int32(0) {
									v109 = v4
									m.G0 = v14 + int32(48)
									return v109
								} else {
									v87 = F_point_dt(m, v80, l2)
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return int32(0)
									} else {
										v89 = F_point_dt(m, v80, v17)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											v91 = base.F64_add(v87, v89)
											v92 = F_point_dt(m, l2, v17)
											mBase = m.M
											v93 = m.ExcPending
											if v93 != 0 {
												return int32(0)
											} else {
												if base.F64_ne(v91, v92)&base.B2i32(base.F64_le(base.F64_abs(base.F64_sub(v91, v92)), float64(1e-06)) == int32(0)) != 0 {
													v109 = v4
												} else {
													v102 = int32(1)
													if l0 == int32(0) {
														v109 = v102
													} else {
														v105 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
														v107 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0))) = v107
														v109 = v102
													}
												}
												m.G0 = v14 + int32(48)
												return v109
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_lseg_length(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 float64
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_point_dt(m, v2, v2+int32(16))
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
func F_lseg_recv(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v12 float64
	_ = v12
	var v13 int32
	_ = v13
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_palloc(m, int32(32))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_pq_getmsgfloat8(m, v3)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v5))) = v9
			v12 = F_pq_getmsgfloat8(m, v3)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v5)+8)) = v12
				v15 = F_pq_getmsgfloat8(m, v3)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v5)+16)) = v15
					v18 = F_pq_getmsgfloat8(m, v3)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v5)+24)) = v18
						return v5
					}
				}
			}
		}
	}
}
