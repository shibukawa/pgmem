package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_close_lseg(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
			if base.F64_eq(v9, v15) != 0 {
				v29 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
				v32 = int32(0)
				return v32
			} else {
				v19 = F_palloc(m, int32(16))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = F_lseg_closept_lseg(m, v19, v5, v6)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						if base.Ui64(base.I64_reinterpret_f64(v21)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							v32 = v19
						} else {
							v29 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
							v32 = int32(0)
						}
						return v32
					}
				}
			}
		}
	}
}
func F_lseg_closept_lseg(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 float64
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 float64
	_ = v42
	var v44 float64
	_ = v44
	var v45 int32
	_ = v45
	var v47 int64
	_ = v47
	var v52 int32
	_ = v52
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 float64
	_ = v67
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v74 int64
	_ = v74
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 float64
	_ = v93
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_lseg_interpt_lseg(m, l0, l1, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return float64(0)
	} else {
		if v11 != 0 {
			v93 = float64(0)
			m.G0 = v9 + int32(16)
			return v93
		} else {
			v15 = F_lseg_closept_point(m, l0, l1, l2)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return float64(0)
			} else {
				v19 = F_lseg_closept_point(m, v9, l1, l2+int32(16))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return float64(0)
				} else {
					v22 = int64(9223372036854775807)
					v27 = int32(0)
					if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&v22))|base.B2i32(base.F64_gt(v15, v19) == v27)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v15)&v22) < base.Ui64(int64(9218868437227405313))) == v27 {
						if l0 != 0 {
							v38 = *(*int64)(unsafe.Add(mBase, uint32(v9)+8))
							*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v38
							v40 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
							*(*int64)(unsafe.Add(mBase, uint32(l0))) = v40
						} else {
						}
						v42 = v19
					} else {
						v42 = v15
					}
					v44 = F_lseg_closept_point(m, int32(0), l2, l1)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return float64(0)
					} else {
						v47 = int64(9223372036854775807)
						v52 = int32(0)
						if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v44)&v47))|base.B2i32(base.F64_gt(v42, v44) == v52)&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v42)&v47) < base.Ui64(int64(9218868437227405313))) == v52 {
							if l0 != 0 {
								v63 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
								*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v63
								v65 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
								*(*int64)(unsafe.Add(mBase, uint32(l0))) = v65
							} else {
							}
							v67 = v44
						} else {
							v67 = v42
						}
						v70 = l1 + int32(16)
						v71 = F_lseg_closept_point(m, int32(0), l2, v70)
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return float64(0)
						} else {
							v74 = int64(9223372036854775807)
							if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v71)&v74))|base.B2i32(base.F64_gt(v67, v71) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v67)&v74) < base.Ui64(int64(9218868437227405313))) != 0 {
								v93 = v67
							} else {
								if l0 != 0 {
									v88 = *(*int64)(unsafe.Add(mBase, uint32(v70)+8))
									*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v88
									v90 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
									*(*int64)(unsafe.Add(mBase, uint32(l0))) = v90
								} else {
								}
								v93 = v71
							}
							m.G0 = v9 + int32(16)
							return v93
						}
					}
				}
			}
		}
	}
}
func F_lseg_crossing(m *base.Module, l0 float64, l1 float64, l2 float64, l3 float64) int32 {
	var v10 int32
	_ = v10
	var v14 float64
	_ = v14
	var v22 float64
	_ = v22
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v68 int32
	_ = v68
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v92 int32
	_ = v92
	var v110 float64
	_ = v110
	var v111 int32
	_ = v111
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v125 float64
	_ = v125
	var v126 float64
	_ = v126
	var v135 float64
	_ = v135
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v142 float64
	_ = v142
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v160 float64
	_ = v160
	var v170 float64
	_ = v170
	var v172 float64
	_ = v172
	var v173 float64
	_ = v173
	var v210 int32
	_ = v210
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	v10 = int32(0)
	v14 = base.F64_abs(l1)
	if base.F64_le(v14, float64(1e-06)) != 0 {
		if base.F64_le(base.F64_abs(l0), float64(1e-06)) != 0 {
			return int32(2147483647)
		} else {
			v22 = base.F64_abs(l3)
			if base.F64_gt(l0, float64(1e-06)) != 0 {
				if base.F64_le(v22, float64(1e-06)) != 0 {
					if base.F64_gt(l2, float64(1e-06)) != 0 {
						v31 = int32(0)
					} else {
						v31 = int32(2147483647)
					}
					return v31
				} else {
					if base.F64_lt(base.F64_add(l3, float64(1e-06)), float64(0)) != 0 {
						v39 = int32(1)
					} else {
						v39 = int32(-1)
					}
					return v39
				}
			} else {
				if base.F64_le(v22, float64(1e-06)) == int32(0) {
					return int32(0)
				} else {
					if base.F64_lt(base.F64_add(l2, float64(1e-06)), float64(0)) != 0 {
						v53 = int32(0)
					} else {
						v53 = int32(2147483647)
					}
					return v53
				}
			}
		}
	} else {
		if base.F64_gt(l1, float64(1e-06)) != 0 {
			v59 = int32(1)
		} else {
			v59 = int32(-1)
		}
		v60 = base.F64_abs(l3)
		if base.F64_le(v60, float64(1e-06)) != 0 {
			if base.F64_lt(base.F64_add(l2, float64(1e-06)), float64(0)) != 0 {
				v68 = int32(0)
			} else {
				v68 = v59
			}
			return v68
		} else {
			if base.F64_gt(l1, float64(1e-06)) == int32(0) {
				if base.F64_lt(base.F64_add(l3, float64(1e-06)), float64(0)) == int32(0) {
					v88 = float64(1e-06)
					v89 = base.F64_add(l0, v88)
					v92 = int32(0)
					if base.B2i32(base.F64_ge(v89, float64(0)) == v92)|base.B2i32(base.F64_gt(l2, v88) == v92) == v92 {
						return v59 << (uint(int32(1)) % 32)
					} else {
						if base.F64_lt(v89, float64(0))&base.F64_le(l2, float64(1e-06)) != 0 {
							v210 = v10
							return v210
						} else {
							v110 = math.Float64frombits(uint64(0x7ff0000000000000))
							v111 = base.F64_eq(base.F64_abs(l0), v110)
							v112 = base.F64_sub(l0, l2)
							v113 = base.F64_abs(v112)
							if base.B2i32(v111|base.F64_ne(v113, v110) == int32(0))&base.F64_ne(base.F64_abs(l2), v110) != 0 {
								F_float_overflow_error(m)
								v223 = m.ExcPending
								if v223 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v123 = math.Float64frombits(uint64(0x7ff0000000000000))
								v124 = base.F64_eq(v14, v123)
								v125 = base.F64_mul(l1, v112)
								v126 = base.F64_abs(v125)
								if base.B2i32(v124|base.F64_ne(v126, v123) == int32(0))&base.F64_ne(v113, v123) != 0 {
									F_float_overflow_error(m)
									v223 = m.ExcPending
									if v223 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v135 = float64(0)
									if base.F64_eq(v125, v135)&base.F64_ne(v112, v135) != 0 {
										F_float_underflow_error(m)
										v229 = m.ExcPending
										if v229 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v140 = base.F64_sub(l1, l3)
										v141 = base.F64_abs(v140)
										v142 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.B2i32(base.F64_ne(v141, v142)|v124 == int32(0))&base.F64_ne(v60, v142) != 0 {
											F_float_overflow_error(m)
											v223 = m.ExcPending
											if v223 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v150 = base.F64_mul(l0, v140)
											v151 = base.F64_abs(v150)
											v152 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.B2i32(base.F64_ne(v151, v152)|v111 == int32(0))&base.F64_ne(v141, v152) != 0 {
												F_float_overflow_error(m)
												v223 = m.ExcPending
												if v223 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v160 = float64(0)
												if base.B2i32(base.F64_eq(l0, v160)|base.F64_ne(v150, v160) == int32(0))&base.F64_ne(v140, v160) != 0 {
													F_float_underflow_error(m)
													v229 = m.ExcPending
													if v229 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v170 = math.Float64frombits(uint64(0x7ff0000000000000))
													v172 = base.F64_sub(v125, v150)
													v173 = base.F64_abs(v172)
													if base.B2i32(base.F64_eq(v126, v170)|base.F64_ne(v173, v170) == int32(0))&base.F64_ne(v151, v170) != 0 {
														F_float_overflow_error(m)
														v223 = m.ExcPending
														if v223 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														if base.F64_le(v173, float64(1e-06)) != 0 {
															return int32(2147483647)
														} else {
															if base.F64_gt(l1, float64(1e-06)) == int32(0) {
																if base.F64_lt(base.F64_add(v172, float64(1e-06)), float64(0)) == int32(0) {
																	v210 = v59 << (uint(int32(1)) % 32)
																} else {
																	v210 = v10
																}
															} else {
																if base.F64_gt(v172, float64(1e-06)) != 0 {
																	v210 = v10
																} else {
																	v210 = v59 << (uint(int32(1)) % 32)
																}
															}
															return v210
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
				} else {
					return int32(0)
				}
			} else {
				if base.F64_gt(l3, float64(1e-06)) == int32(0) {
					v88 = float64(1e-06)
					v89 = base.F64_add(l0, v88)
					v92 = int32(0)
					if base.B2i32(base.F64_ge(v89, float64(0)) == v92)|base.B2i32(base.F64_gt(l2, v88) == v92) == v92 {
						return v59 << (uint(int32(1)) % 32)
					} else {
						if base.F64_lt(v89, float64(0))&base.F64_le(l2, float64(1e-06)) != 0 {
							v210 = v10
							return v210
						} else {
							v110 = math.Float64frombits(uint64(0x7ff0000000000000))
							v111 = base.F64_eq(base.F64_abs(l0), v110)
							v112 = base.F64_sub(l0, l2)
							v113 = base.F64_abs(v112)
							if base.B2i32(v111|base.F64_ne(v113, v110) == int32(0))&base.F64_ne(base.F64_abs(l2), v110) != 0 {
								F_float_overflow_error(m)
								v223 = m.ExcPending
								if v223 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v123 = math.Float64frombits(uint64(0x7ff0000000000000))
								v124 = base.F64_eq(v14, v123)
								v125 = base.F64_mul(l1, v112)
								v126 = base.F64_abs(v125)
								if base.B2i32(v124|base.F64_ne(v126, v123) == int32(0))&base.F64_ne(v113, v123) != 0 {
									F_float_overflow_error(m)
									v223 = m.ExcPending
									if v223 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v135 = float64(0)
									if base.F64_eq(v125, v135)&base.F64_ne(v112, v135) != 0 {
										F_float_underflow_error(m)
										v229 = m.ExcPending
										if v229 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v140 = base.F64_sub(l1, l3)
										v141 = base.F64_abs(v140)
										v142 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.B2i32(base.F64_ne(v141, v142)|v124 == int32(0))&base.F64_ne(v60, v142) != 0 {
											F_float_overflow_error(m)
											v223 = m.ExcPending
											if v223 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v150 = base.F64_mul(l0, v140)
											v151 = base.F64_abs(v150)
											v152 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.B2i32(base.F64_ne(v151, v152)|v111 == int32(0))&base.F64_ne(v141, v152) != 0 {
												F_float_overflow_error(m)
												v223 = m.ExcPending
												if v223 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v160 = float64(0)
												if base.B2i32(base.F64_eq(l0, v160)|base.F64_ne(v150, v160) == int32(0))&base.F64_ne(v140, v160) != 0 {
													F_float_underflow_error(m)
													v229 = m.ExcPending
													if v229 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v170 = math.Float64frombits(uint64(0x7ff0000000000000))
													v172 = base.F64_sub(v125, v150)
													v173 = base.F64_abs(v172)
													if base.B2i32(base.F64_eq(v126, v170)|base.F64_ne(v173, v170) == int32(0))&base.F64_ne(v151, v170) != 0 {
														F_float_overflow_error(m)
														v223 = m.ExcPending
														if v223 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														if base.F64_le(v173, float64(1e-06)) != 0 {
															return int32(2147483647)
														} else {
															if base.F64_gt(l1, float64(1e-06)) == int32(0) {
																if base.F64_lt(base.F64_add(v172, float64(1e-06)), float64(0)) == int32(0) {
																	v210 = v59 << (uint(int32(1)) % 32)
																} else {
																	v210 = v10
																}
															} else {
																if base.F64_gt(v172, float64(1e-06)) != 0 {
																	v210 = v10
																} else {
																	v210 = v59 << (uint(int32(1)) % 32)
																}
															}
															return v210
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
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_lseg_interpt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_palloc(m, int32(16))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_lseg_interpt_lseg(m, v8, v6, v5)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v17 = v8
			} else {
				v14 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
				v17 = int32(0)
			}
			return v17
		}
	}
}
func F_lseg_vertical(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v4)+16))
	return base.F64_eq(v5, v6) | base.F64_le(base.F64_abs(base.F64_sub(v5, v6)), float64(1e-06))
}
