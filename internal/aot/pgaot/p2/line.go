package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_line_closept_point(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v44 float64
	_ = v44
	var v54 float64
	_ = v54
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v63 float64
	_ = v63
	var v69 float64
	_ = v69
	var v74 float64
	_ = v74
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 float64
	_ = v102
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v115 float64
	_ = v115
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v14 = base.F64_abs(v13)
	if base.F64_le(v14, float64(1e-06)) == int32(0) {
		v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
		v20 = base.F64_abs(v19)
		if base.F64_le(v20, float64(1e-06)) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(-4616189618054758400)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
			v54 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v54
			v96 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return float64(0)
			} else {
				if v96 == int32(0) {
					v102 = math.Float64frombits(uint64(0x7ff8000000000000))
					if l0 == int32(0) {
						v117 = v102
					} else {
						v105 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
						v107 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v107
						v117 = v102
					}
					m.G0 = v11 + int32(48)
					return v117
				} else {
					if l0 != 0 {
						v109 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v109
						v111 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
					} else {
					}
					v115 = F_point_dt(m, v11+int32(32), l2)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return float64(0)
					} else {
						v117 = v115
						m.G0 = v11 + int32(48)
						return v117
					}
				}
			}
		} else {
			v23 = base.F64_div(v19, v13)
			v24 = base.F64_abs(v23)
			v25 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(v24, v25)&base.F64_ne(v20, v25) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return float64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v23, float64(0))&base.F64_ne(v14, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v133 = m.ExcPending
					if v133 != 0 {
						return float64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if base.F64_ne(v24, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						if base.F64_ne(v23, float64(0)) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(-4616189618054758400)
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v23
							v59 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
							v60 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
							v61 = base.F64_mul(v23, v60)
							v62 = base.F64_abs(v61)
							v63 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(v62, v63)&base.F64_ne(base.F64_abs(v60), v63) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v69 = float64(0)
								if base.F64_eq(v61, v69)&base.F64_ne(v60, v69) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v74 = base.F64_sub(v59, v61)
									if base.F64_ne(base.F64_abs(v74), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v74
										if base.F64_ne(v74, float64(0)) != 0 {
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
										}
										v96 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return float64(0)
										} else {
											if v96 == int32(0) {
												v102 = math.Float64frombits(uint64(0x7ff8000000000000))
												if l0 == int32(0) {
													v117 = v102
												} else {
													v105 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
													v107 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v107
													v117 = v102
												}
												m.G0 = v11 + int32(48)
												return v117
											} else {
												if l0 != 0 {
													v109 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
													*(*int64)(unsafe.Add(mBase, uint32(l0))) = v109
													v111 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
													*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
												} else {
												}
												v115 = F_point_dt(m, v11+int32(32), l2)
												mBase = m.M
												v116 = m.ExcPending
												if v116 != 0 {
													return float64(0)
												} else {
													v117 = v115
													m.G0 = v11 + int32(48)
													return v117
												}
											}
										}
									} else {
										if base.F64_eq(base.F64_abs(v59), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v74
											if base.F64_ne(v74, float64(0)) != 0 {
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
											}
											v96 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
											mBase = m.M
											v99 = m.ExcPending
											if v99 != 0 {
												return float64(0)
											} else {
												if v96 == int32(0) {
													v102 = math.Float64frombits(uint64(0x7ff8000000000000))
													if l0 == int32(0) {
														v117 = v102
													} else {
														v105 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
														*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
														v107 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v107
														v117 = v102
													}
													m.G0 = v11 + int32(48)
													return v117
												} else {
													if l0 != 0 {
														v109 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0))) = v109
														v111 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
													} else {
													}
													v115 = F_point_dt(m, v11+int32(32), l2)
													mBase = m.M
													v116 = m.ExcPending
													if v116 != 0 {
														return float64(0)
													} else {
														v117 = v115
														m.G0 = v11 + int32(48)
														return v117
													}
												}
											}
										} else {
											if base.F64_ne(v62, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v74
												if base.F64_ne(v74, float64(0)) != 0 {
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = int64(0)
												}
												v96 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
												mBase = m.M
												v99 = m.ExcPending
												if v99 != 0 {
													return float64(0)
												} else {
													if v96 == int32(0) {
														v102 = math.Float64frombits(uint64(0x7ff8000000000000))
														if l0 == int32(0) {
															v117 = v102
														} else {
															v105 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
															*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
															v107 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v107
															v117 = v102
														}
														m.G0 = v11 + int32(48)
														return v117
													} else {
														if l0 != 0 {
															v109 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
															*(*int64)(unsafe.Add(mBase, uint32(l0))) = v109
															v111 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
														} else {
														}
														v115 = F_point_dt(m, v11+int32(32), l2)
														mBase = m.M
														v116 = m.ExcPending
														if v116 != 0 {
															return float64(0)
														} else {
															v117 = v115
															m.G0 = v11 + int32(48)
															return v117
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(-4616189618054758400)
							*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(0)
							v54 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v54
							v96 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return float64(0)
							} else {
								if v96 == int32(0) {
									v102 = math.Float64frombits(uint64(0x7ff8000000000000))
									if l0 == int32(0) {
										v117 = v102
									} else {
										v105 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
										v107 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v107
										v117 = v102
									}
									m.G0 = v11 + int32(48)
									return v117
								} else {
									if l0 != 0 {
										v109 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v109
										v111 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
									} else {
									}
									v115 = F_point_dt(m, v11+int32(32), l2)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return float64(0)
									} else {
										v117 = v115
										m.G0 = v11 + int32(48)
										return v117
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(-4616189618054758400)
						v44 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v44
						v96 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return float64(0)
						} else {
							if v96 == int32(0) {
								v102 = math.Float64frombits(uint64(0x7ff8000000000000))
								if l0 == int32(0) {
									v117 = v102
								} else {
									v105 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
									v107 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v107
									v117 = v102
								}
								m.G0 = v11 + int32(48)
								return v117
							} else {
								if l0 != 0 {
									v109 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v109
									v111 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
								} else {
								}
								v115 = F_point_dt(m, v11+int32(32), l2)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return float64(0)
								} else {
									v117 = v115
									m.G0 = v11 + int32(48)
									return v117
								}
							}
						}
					}
				}
			}
		}
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = int64(-4616189618054758400)
		v44 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
		*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v44
		v96 = F_line_interpt_line(m, v11+int32(32), v11+int32(8), l1)
		mBase = m.M
		v99 = m.ExcPending
		if v99 != 0 {
			return float64(0)
		} else {
			if v96 == int32(0) {
				v102 = math.Float64frombits(uint64(0x7ff8000000000000))
				if l0 == int32(0) {
					v117 = v102
				} else {
					v105 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v105
					v107 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v107
					v117 = v102
				}
				m.G0 = v11 + int32(48)
				return v117
			} else {
				if l0 != 0 {
					v109 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
					*(*int64)(unsafe.Add(mBase, uint32(l0))) = v109
					v111 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
					*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v111
				} else {
				}
				v115 = F_point_dt(m, v11+int32(32), l2)
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return float64(0)
				} else {
					v117 = v115
					m.G0 = v11 + int32(48)
					return v117
				}
			}
		}
	}
}
func F_line_construct(m *base.Module, l0 int32, l1 int32, l2 float64) {
	mBase := m.M
	_ = mBase
	var v14 float64
	_ = v14
	var v22 float64
	_ = v22
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v37 float64
	_ = v37
	var v42 float64
	_ = v42
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	if base.F64_eq(base.F64_abs(l2), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(-4616189618054758400)
		v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
		*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v14
		return
	} else {
		if base.F64_eq(l2, float64(0)) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4616189618054758400)
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v22
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4616189618054758400)
			*(*float64)(unsafe.Add(mBase, uint32(l0))) = l2
			v27 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			v28 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
			v29 = base.F64_mul(l2, v28)
			v30 = base.F64_abs(v29)
			v31 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(v30, v31)&base.F64_ne(base.F64_abs(v28), v31) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v37 = float64(0)
				if base.F64_eq(v29, v37)&base.F64_ne(v28, v37) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v42 = base.F64_sub(v27, v29)
					if base.F64_ne(base.F64_abs(v42), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v42
						if base.F64_eq(v42, float64(0)) != 0 {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
						} else {
						}
						return
					} else {
						if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v42
							if base.F64_eq(v42, float64(0)) != 0 {
								*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							} else {
							}
							return
						} else {
							if base.F64_ne(v30, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(l0)+16)) = v42
								if base.F64_eq(v42, float64(0)) != 0 {
									*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
								} else {
								}
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_line_interpt_line(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v40 float64
	_ = v40
	var v60 float64
	_ = v60
	var v61 float64
	_ = v61
	var v62 float64
	_ = v62
	var v70 float64
	_ = v70
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v116 float64
	_ = v116
	var v117 float64
	_ = v117
	var v125 float64
	_ = v125
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v144 float64
	_ = v144
	var v145 float64
	_ = v145
	var v146 float64
	_ = v146
	var v157 float64
	_ = v157
	var v158 float64
	_ = v158
	var v172 float64
	_ = v172
	var v177 float64
	_ = v177
	var v185 float64
	_ = v185
	var v189 float64
	_ = v189
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v200 float64
	_ = v200
	var v204 float64
	_ = v204
	var v205 float64
	_ = v205
	var v212 float64
	_ = v212
	var v214 float64
	_ = v214
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v241 float64
	_ = v241
	var v246 float64
	_ = v246
	var v247 float64
	_ = v247
	var v248 float64
	_ = v248
	var v249 float64
	_ = v249
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v269 float64
	_ = v269
	var v270 float64
	_ = v270
	var v271 float64
	_ = v271
	var v283 float64
	_ = v283
	var v284 float64
	_ = v284
	var v292 float64
	_ = v292
	var v297 float64
	_ = v297
	var v298 float64
	_ = v298
	var v311 float64
	_ = v311
	var v312 float64
	_ = v312
	var v313 float64
	_ = v313
	var v324 float64
	_ = v324
	var v325 float64
	_ = v325
	var v339 float64
	_ = v339
	var v344 float64
	_ = v344
	var v352 float64
	_ = v352
	var v356 float64
	_ = v356
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v375 int32
	_ = v375
	var v378 float64
	_ = v378
	var v381 float64
	_ = v381
	var v383 float64
	_ = v383
	var v386 float64
	_ = v386
	var v400 int32
	_ = v400
	var v417 int32
	_ = v417
	var v431 int32
	_ = v431
	var v445 int32
	_ = v445
	v16 = int32(0)
	v17 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	v18 = base.F64_abs(v17)
	if base.F64_le(v18, float64(1e-06)) == v16 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L152
	} else {
		goto L155
	}
L2:
	;
	F_float_zero_divide_error(m)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L152
	} else {
		goto L154
	}
L3:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L152
	} else {
		goto L153
	}
L4:
	;
	return v400
L5:
	;
	v375 = int32(1)
	if l0 == int32(0) {
		v400 = v375
		goto L4
	} else {
		goto L145
	}
L6:
	;
	v23 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v24 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v25 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v26 = base.F64_div(v25, v17)
	v27 = base.F64_abs(v26)
	v28 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v27, v28)&base.F64_ne(base.F64_abs(v25), v28) != 0 {
		goto L3
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v196 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
	v197 = base.F64_abs(v196)
	if base.F64_le(v197, float64(1e-06)) != 0 {
		v400 = v16
		goto L4
	} else {
		goto L81
	}
L9:
	;
	if base.F64_ne(v26, float64(0)) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v40 = base.F64_mul(v23, v26)
	if base.F64_ne(base.F64_abs(v40), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if base.F64_eq(v18, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	if base.F64_ne(v25, float64(0)) != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	if base.F64_ne(v40, float64(0)) != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	if base.F64_eq(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if base.F64_ne(v27, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	if base.F64_eq(v40, v24) != 0 {
		v400 = v16
		goto L4
	} else {
		goto L22
	}
L19:
	;
	if base.F64_eq(v23, float64(0)) != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	if base.F64_ne(v26, float64(0)) != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v24, v40)), float64(1e-06)) != 0 {
		v400 = v16
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v60 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v61 = base.F64_mul(v17, v60)
	v62 = base.F64_abs(v61)
	if base.F64_ne(v62, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v70 = float64(0)
	if base.F64_eq(v61, v70)&base.F64_ne(v60, v70) != 0 {
		goto L1
	} else {
		goto L28
	}
L25:
	;
	if base.F64_eq(v18, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v75 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v76 = base.F64_mul(v25, v75)
	v77 = base.F64_abs(v76)
	if base.F64_ne(v77, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	if base.F64_ne(v76, float64(0)) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	if base.F64_eq(base.F64_abs(v25), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	if base.F64_ne(base.F64_abs(v75), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v92 = base.F64_sub(v61, v76)
	v93 = base.F64_abs(v92)
	if base.F64_ne(v93, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	if base.F64_eq(v25, float64(0)) != 0 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	if base.F64_ne(v75, float64(0)) != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v100 = base.F64_mul(v23, v25)
	v101 = base.F64_abs(v100)
	if base.F64_ne(v101, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if base.F64_eq(v62, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	if base.F64_ne(v77, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	if base.F64_ne(v100, float64(0)) != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	if base.F64_eq(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	if base.F64_ne(base.F64_abs(v25), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v116 = base.F64_mul(v17, v24)
	v117 = base.F64_abs(v116)
	if base.F64_ne(v117, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	if base.F64_eq(v23, float64(0)) != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	if base.F64_ne(v25, float64(0)) != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v125 = float64(0)
	if base.F64_eq(v116, v125)&base.F64_ne(v24, v125) != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	if base.F64_eq(v18, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	if base.F64_ne(base.F64_abs(v24), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	v130 = base.F64_sub(v100, v116)
	v131 = base.F64_abs(v130)
	if base.F64_ne(v131, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v93)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v130, float64(0)) != 0 {
		goto L2
	} else {
		goto L58
	}
L55:
	;
	if base.F64_eq(v117, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	if base.F64_ne(v101, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	v144 = base.F64_div(v92, v130)
	v145 = base.F64_abs(v144)
	v146 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v145, v146)&base.F64_ne(v93, v146) != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	if base.F64_ne(v144, float64(0)) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v157 = base.F64_mul(v23, v144)
	v158 = base.F64_abs(v157)
	if base.F64_ne(v158, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	if base.F64_eq(v131, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	if base.F64_ne(v92, float64(0)) != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	if base.F64_ne(v157, float64(0)) != 0 {
		goto L68
	} else {
		goto L69
	}
L65:
	;
	if base.F64_eq(base.F64_abs(v23), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	if base.F64_ne(v145, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	v172 = base.F64_add(v75, v157)
	if base.F64_eq(base.F64_abs(v172), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L73
	} else {
		goto L74
	}
L69:
	;
	if base.F64_eq(v23, float64(0)) != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	if base.F64_ne(v144, float64(0)) != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L68
L72:
	;
	if base.F64_ne(v189, float64(0)) != 0 {
		v363 = v144
		v364 = v189
		goto L5
	} else {
		goto L78
	}
L73:
	;
	v177 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v75), v177)&base.F64_ne(v158, v177) != 0 {
		goto L3
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v185 = base.F64_div(base.F64_neg(v172), v17)
	if base.F64_eq(base.F64_abs(v185), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L77
	}
L76:
	;
	v189 = base.F64_div(base.F64_neg(v172), v17)
	goto L72
L77:
	;
	v189 = v185
	goto L72
L78:
	;
	if base.F64_eq(v18, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v363 = v144
		v364 = v189
		goto L5
	} else {
		goto L79
	}
L79:
	;
	if base.F64_eq(v172, float64(0)) != 0 {
		v363 = v144
		v364 = v189
		goto L5
	} else {
		goto L80
	}
L80:
	;
	goto L1
L81:
	;
	v200 = base.F64_div(v17, v196)
	if base.F64_eq(base.F64_abs(v200), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L82
	}
L82:
	;
	v204 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v205 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_ne(v200, float64(0)) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v212 = base.F64_mul(v200, v204)
	v214 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v212), v214)&base.F64_ne(base.F64_abs(v204), v214) != 0 {
		goto L3
	} else {
		goto L87
	}
L84:
	;
	if base.F64_eq(v17, float64(0)) != 0 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	if base.F64_ne(v197, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	if base.F64_ne(v212, float64(0)) != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if base.F64_eq(v212, v205) != 0 {
		v400 = v16
		goto L4
	} else {
		goto L92
	}
L89:
	;
	if base.F64_eq(v200, float64(0)) != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	if base.F64_ne(v204, float64(0)) != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	if base.F64_le(base.F64_abs(base.F64_sub(v205, v212)), float64(1e-06)) != 0 {
		v400 = v16
		goto L4
	} else {
		goto L93
	}
L93:
	;
	v231 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	v232 = base.F64_mul(v196, v231)
	v233 = base.F64_abs(v232)
	if base.F64_ne(v233, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v241 = float64(0)
	if base.F64_eq(v232, v241)&base.F64_ne(v231, v241) != 0 {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	if base.F64_eq(v197, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L94
	} else {
		goto L96
	}
L96:
	;
	if base.F64_ne(base.F64_abs(v231), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L97
	}
L97:
	;
	goto L94
L98:
	;
	v246 = *(*float64)(unsafe.Add(mBase, uint32(l2)+16))
	v247 = base.F64_mul(v17, v246)
	v248 = base.F64_abs(v247)
	v249 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v248, v249)&base.F64_ne(base.F64_abs(v246), v249) != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	if base.F64_ne(v247, float64(0)) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v261 = base.F64_sub(v232, v247)
	v262 = base.F64_abs(v261)
	if base.F64_ne(v262, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	if base.F64_eq(v17, float64(0)) != 0 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	if base.F64_ne(v246, float64(0)) != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L100
L104:
	;
	v269 = base.F64_mul(v17, v204)
	v270 = base.F64_abs(v269)
	v271 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v270, v271)&base.F64_ne(base.F64_abs(v204), v271) != 0 {
		goto L3
	} else {
		goto L108
	}
L105:
	;
	if base.F64_eq(v233, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L104
	} else {
		goto L106
	}
L106:
	;
	if base.F64_ne(v248, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	goto L104
L108:
	;
	if base.F64_ne(v269, float64(0)) != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v283 = base.F64_mul(v196, v205)
	v284 = base.F64_abs(v283)
	if base.F64_ne(v284, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	if base.F64_eq(v17, float64(0)) != 0 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	if base.F64_ne(v204, float64(0)) != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	v292 = float64(0)
	if base.F64_eq(v283, v292)&base.F64_ne(v205, v292) != 0 {
		goto L1
	} else {
		goto L117
	}
L114:
	;
	if base.F64_eq(v197, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	if base.F64_ne(base.F64_abs(v205), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	goto L113
L117:
	;
	v297 = base.F64_sub(v269, v283)
	v298 = base.F64_abs(v297)
	if base.F64_ne(v298, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	if base.B2i32(base.Ui64(base.I64_reinterpret_f64(v262)) <= base.Ui64(int64(9218868437227405312)))&base.F64_eq(v297, float64(0)) != 0 {
		goto L2
	} else {
		goto L122
	}
L119:
	;
	if base.F64_eq(v284, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L118
	} else {
		goto L120
	}
L120:
	;
	if base.F64_ne(v270, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L121
	}
L121:
	;
	goto L118
L122:
	;
	v311 = base.F64_div(v261, v297)
	v312 = base.F64_abs(v311)
	v313 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v312, v313)&base.F64_ne(v262, v313) != 0 {
		goto L3
	} else {
		goto L123
	}
L123:
	;
	if base.F64_ne(v311, float64(0)) != 0 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v324 = base.F64_mul(v204, v311)
	v325 = base.F64_abs(v324)
	if base.F64_ne(v325, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	if base.F64_eq(v298, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	if base.F64_ne(v261, float64(0)) != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	if base.F64_ne(v324, float64(0)) != 0 {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	if base.F64_eq(base.F64_abs(v204), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	if base.F64_ne(v312, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	v339 = base.F64_add(v246, v324)
	if base.F64_eq(base.F64_abs(v339), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L137
	} else {
		goto L138
	}
L133:
	;
	if base.F64_eq(v204, float64(0)) != 0 {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	if base.F64_ne(v311, float64(0)) != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	goto L132
L136:
	;
	if base.F64_ne(v356, float64(0)) != 0 {
		v363 = v311
		v364 = v356
		goto L5
	} else {
		goto L142
	}
L137:
	;
	v344 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v246), v344)&base.F64_ne(v325, v344) != 0 {
		goto L3
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v352 = base.F64_div(base.F64_neg(v339), v196)
	if base.F64_eq(base.F64_abs(v352), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L3
	} else {
		goto L141
	}
L140:
	;
	v356 = base.F64_div(base.F64_neg(v339), v196)
	goto L136
L141:
	;
	v356 = v352
	goto L136
L142:
	;
	if base.F64_eq(v197, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v363 = v311
		v364 = v356
		goto L5
	} else {
		goto L143
	}
L143:
	;
	if base.F64_ne(v339, float64(0)) != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v363 = v311
	v364 = v356
	goto L5
L145:
	;
	v378 = float64(0)
	if base.F64_eq(v364, v378) != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v381 = v378
	goto L148
L147:
	;
	v381 = v364
	goto L148
L148:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v381
	v383 = float64(0)
	if base.F64_eq(v363, v383) != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v386 = v383
	goto L151
L150:
	;
	v386 = v363
	goto L151
L151:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v386
	v400 = v375
	goto L4
L152:
	;
	return int32(0)
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_line_intersect(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_line_interpt_line(m, int32(0), v3, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_line_perp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v17 float64
	_ = v17
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v31 int32
	_ = v31
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v58 float64
	_ = v58
	var v60 float64
	_ = v60
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v14 = base.F64_abs(v13)
	if base.F64_le(v14, float64(1e-06)) != 0 {
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
		return base.F64_le(base.F64_abs(v17), float64(1e-06))
	} else {
		v22 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
		v23 = base.F64_abs(v22)
		v24 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v25 = base.F64_abs(v24)
		if base.F64_le(v25, float64(1e-06)) != 0 {
			return base.F64_le(v23, float64(1e-06))
		} else {
			v31 = int32(0)
			if base.F64_le(v23, float64(1e-06)) != 0 {
				v78 = v31
				return v78
			} else {
				v34 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v35 = base.F64_abs(v34)
				if base.F64_le(v35, float64(1e-06)) != 0 {
					v78 = v31
					return v78
				} else {
					v38 = base.F64_mul(v13, v24)
					v39 = base.F64_abs(v38)
					if base.F64_ne(v39, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						if base.F64_eq(v38, float64(0)) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v48 = base.F64_mul(v22, v34)
							v49 = base.F64_abs(v48)
							if base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								if base.F64_eq(v48, float64(0)) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v58 = base.F64_div(v38, v48)
									v60 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
											return v78
										}
									}
								}
							} else {
								if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									if base.F64_eq(v48, float64(0)) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v58 = base.F64_div(v38, v48)
										v60 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
												return v78
											}
										}
									}
								} else {
									if base.F64_ne(v35, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_eq(v48, float64(0)) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v58 = base.F64_div(v38, v48)
											v60 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
													return v78
												}
											}
										}
									}
								}
							}
						}
					} else {
						if base.F64_eq(v14, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							if base.F64_eq(v38, float64(0)) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v48 = base.F64_mul(v22, v34)
								v49 = base.F64_abs(v48)
								if base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									if base.F64_eq(v48, float64(0)) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v58 = base.F64_div(v38, v48)
										v60 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
												return v78
											}
										}
									}
								} else {
									if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										if base.F64_eq(v48, float64(0)) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v58 = base.F64_div(v38, v48)
											v60 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
													return v78
												}
											}
										}
									} else {
										if base.F64_ne(v35, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_eq(v48, float64(0)) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v58 = base.F64_div(v38, v48)
												v60 = math.Float64frombits(uint64(0x7ff0000000000000))
												if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
													F_float_overflow_error(m)
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														F_float_underflow_error(m)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
														return v78
													}
												}
											}
										}
									}
								}
							}
						} else {
							if base.F64_ne(v25, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.F64_eq(v38, float64(0)) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v48 = base.F64_mul(v22, v34)
									v49 = base.F64_abs(v48)
									if base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										if base.F64_eq(v48, float64(0)) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v58 = base.F64_div(v38, v48)
											v60 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
													return v78
												}
											}
										}
									} else {
										if base.F64_eq(v23, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											if base.F64_eq(v48, float64(0)) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v58 = base.F64_div(v38, v48)
												v60 = math.Float64frombits(uint64(0x7ff0000000000000))
												if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
													F_float_overflow_error(m)
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														F_float_underflow_error(m)
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
														return v78
													}
												}
											}
										} else {
											if base.F64_ne(v35, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_eq(v48, float64(0)) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v58 = base.F64_div(v38, v48)
													v60 = math.Float64frombits(uint64(0x7ff0000000000000))
													if base.F64_eq(base.F64_abs(v58), v60)&base.F64_ne(v39, v60) != 0 {
														F_float_overflow_error(m)
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														if base.F64_eq(v58, float64(0))&base.F64_ne(v49, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
															F_float_underflow_error(m)
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v78 = base.F64_eq(v58, float64(-1)) | base.F64_le(base.F64_abs(base.F64_add(v58, float64(1))), float64(1e-06))
															return v78
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
		}
	}
}
