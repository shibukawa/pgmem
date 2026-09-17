package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_point_div_point(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v30 float64
	_ = v30
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v44 float64
	_ = v44
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v74 int32
	_ = v74
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v99 int32
	_ = v99
	var v101 float64
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 float64
	_ = v111
	var v113 float64
	_ = v113
	var v114 float64
	_ = v114
	var v124 int32
	_ = v124
	var v129 float64
	_ = v129
	var v131 float64
	_ = v131
	var v137 int32
	_ = v137
	var v138 float64
	_ = v138
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v152 int32
	_ = v152
	var v155 float64
	_ = v155
	var v163 float64
	_ = v163
	var v164 float64
	_ = v164
	var v177 float64
	_ = v177
	var v179 float64
	_ = v179
	var v180 float64
	_ = v180
	var v193 float64
	_ = v193
	var v195 float64
	_ = v195
	var v200 float64
	_ = v200
	var v230 int32
	_ = v230
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	v21 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
	v22 = base.F64_mul(v21, v21)
	v23 = base.F64_abs(v22)
	v24 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v23, v24)&base.F64_ne(base.F64_abs(v21), v24) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v230 = m.ExcPending
		if v230 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v30 = float64(0)
		if base.F64_eq(v22, v30)&base.F64_ne(v21, v30) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v251 = m.ExcPending
			if v251 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v35 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
			v36 = base.F64_mul(v35, v35)
			v37 = base.F64_abs(v36)
			v38 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_eq(v37, v38)&base.F64_ne(base.F64_abs(v35), v38) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v230 = m.ExcPending
				if v230 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				v44 = float64(0)
				if base.F64_eq(v36, v44)&base.F64_ne(v35, v44) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v251 = m.ExcPending
					if v251 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v49 = math.Float64frombits(uint64(0x7ff0000000000000))
					v51 = base.F64_add(v22, v36)
					v52 = base.F64_abs(v51)
					if base.B2i32(base.F64_eq(v23, v49)|base.F64_ne(v52, v49) == int32(0))&base.F64_ne(v37, v49) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v230 = m.ExcPending
						if v230 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v62 = math.Float64frombits(uint64(0x7ff0000000000000))
						v63 = base.F64_eq(base.F64_abs(v21), v62)
						v64 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
						v65 = base.F64_mul(v21, v64)
						v66 = base.F64_abs(v65)
						v74 = base.F64_ne(base.F64_abs(v64), v62)
						if base.B2i32(v63|base.F64_ne(v66, v62) == int32(0))&v74 != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v230 = m.ExcPending
							if v230 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v76 = float64(0)
							v77 = base.F64_ne(v64, v76)
							if v77&base.B2i32(base.F64_eq(v21, v76)|base.F64_ne(v65, v76) == int32(0)) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v251 = m.ExcPending
								if v251 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								v87 = math.Float64frombits(uint64(0x7ff0000000000000))
								v88 = base.F64_eq(base.F64_abs(v35), v87)
								v89 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
								v90 = base.F64_mul(v35, v89)
								v91 = base.F64_abs(v90)
								v99 = base.F64_ne(base.F64_abs(v89), v87)
								if base.B2i32(v88|base.F64_ne(v91, v87) == int32(0))&v99 != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v230 = m.ExcPending
									if v230 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v101 = float64(0)
									v102 = base.F64_ne(v89, v101)
									v104 = base.F64_eq(v35, v101)
									if v102&base.B2i32(v104|base.F64_ne(v90, v101) == int32(0)) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v251 = m.ExcPending
										if v251 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v111 = math.Float64frombits(uint64(0x7ff0000000000000))
										v113 = base.F64_add(v65, v90)
										v114 = base.F64_abs(v113)
										if base.B2i32(base.F64_eq(v66, v111)|base.F64_ne(v114, v111) == int32(0))&base.F64_ne(v91, v111) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v230 = m.ExcPending
											if v230 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v124 = base.F64_eq(v51, float64(0))
											if v124&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v114)) <= base.Ui64(int64(9218868437227405312))) != 0 {
												F_float_zero_divide_error(m)
												mBase = m.M
												v260 = m.ExcPending
												if v260 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v129 = base.F64_div(v113, v51)
												v131 = math.Float64frombits(uint64(0x7ff0000000000000))
												if base.F64_eq(base.F64_abs(v129), v131)&base.F64_ne(v114, v131) != 0 {
													F_float_overflow_error(m)
													mBase = m.M
													v230 = m.ExcPending
													if v230 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v137 = base.F64_eq(v52, math.Float64frombits(uint64(0x7ff0000000000000)))
													v138 = float64(0)
													if base.B2i32(v137|base.F64_ne(v129, v138) == int32(0))&base.F64_ne(v113, v138) != 0 {
														F_float_underflow_error(m)
														mBase = m.M
														v251 = m.ExcPending
														if v251 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v146 = base.F64_mul(v21, v89)
														v147 = base.F64_abs(v146)
														if v99 != 0 {
															v152 = base.F64_ne(v147, math.Float64frombits(uint64(0x7ff0000000000000))) | v63
														} else {
															v152 = int32(1)
														}
														if v152 == int32(0) {
															F_float_overflow_error(m)
															mBase = m.M
															v230 = m.ExcPending
															if v230 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															v155 = float64(0)
															if v102&base.B2i32(base.F64_eq(v21, v155)|base.F64_ne(v146, v155) == int32(0)) != 0 {
																F_float_underflow_error(m)
																mBase = m.M
																v251 = m.ExcPending
																if v251 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																v163 = base.F64_mul(v35, v64)
																v164 = base.F64_abs(v163)
																if v74&base.B2i32(base.F64_ne(v164, math.Float64frombits(uint64(0x7ff0000000000000)))|v88 == int32(0)) != 0 {
																	F_float_overflow_error(m)
																	mBase = m.M
																	v230 = m.ExcPending
																	if v230 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																} else {
																	if v77&base.B2i32(base.F64_ne(v163, float64(0))|v104 == int32(0)) != 0 {
																		F_float_underflow_error(m)
																		mBase = m.M
																		v251 = m.ExcPending
																		if v251 != 0 {
																			return
																		} else {
																			base.Wasm_trap_unreachable()
																			for {
																			}
																		}
																	} else {
																		v177 = math.Float64frombits(uint64(0x7ff0000000000000))
																		v179 = base.F64_sub(v146, v163)
																		v180 = base.F64_abs(v179)
																		if base.B2i32(base.F64_eq(v164, v177)|base.F64_ne(v180, v177) == int32(0))&base.F64_ne(v147, v177) != 0 {
																			F_float_overflow_error(m)
																			mBase = m.M
																			v230 = m.ExcPending
																			if v230 != 0 {
																				return
																			} else {
																				base.Wasm_trap_unreachable()
																				for {
																				}
																			}
																		} else {
																			if v124&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v180)) <= base.Ui64(int64(9218868437227405312))) != 0 {
																				F_float_zero_divide_error(m)
																				mBase = m.M
																				v260 = m.ExcPending
																				if v260 != 0 {
																					return
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			} else {
																				v193 = base.F64_div(v179, v51)
																				v195 = math.Float64frombits(uint64(0x7ff0000000000000))
																				if base.F64_eq(base.F64_abs(v193), v195)&base.F64_ne(v180, v195) != 0 {
																					F_float_overflow_error(m)
																					mBase = m.M
																					v230 = m.ExcPending
																					if v230 != 0 {
																						return
																					} else {
																						base.Wasm_trap_unreachable()
																						for {
																						}
																					}
																				} else {
																					v200 = float64(0)
																					if base.B2i32(base.F64_ne(v193, v200)|v137 == int32(0))&base.F64_ne(v179, v200) != 0 {
																						F_float_underflow_error(m)
																						mBase = m.M
																						v251 = m.ExcPending
																						if v251 != 0 {
																							return
																						} else {
																							base.Wasm_trap_unreachable()
																							for {
																							}
																						}
																					} else {
																						*(*float64)(unsafe.Add(mBase, uint32(l0)+8)) = v193
																						*(*float64)(unsafe.Add(mBase, uint32(l0))) = v129
																						return
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
						}
					}
				}
			}
		}
	}
}
func F_point_horiz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
	return base.F64_eq(v5, v7) | base.F64_le(base.F64_abs(base.F64_sub(v5, v7)), float64(1e-06))
}
func F_point_right(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	return base.F64_gt(v3, base.F64_add(v5, float64(1e-06)))
}
func F_point_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v5)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		F_pq_sendfloat8(m, v5, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			F_pq_sendfloat8(m, v5, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v19))) = v20 << (uint(int32(2)) % 32)
				m.G0 = v5 + int32(16)
				return v19
			}
		}
	}
}
