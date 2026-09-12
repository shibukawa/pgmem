package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_circle_add_pt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v41 float64
	_ = v41
	var v48 int32
	_ = v48
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v17 = base.F64_add(v15, v16)
		if base.F64_ne(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v29 = base.F64_add(v27, v28)
			if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
				return v11
			} else {
				if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
					return v11
				} else {
					if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
						return v11
					}
				}
			}
		} else {
			if base.F64_eq(base.F64_abs(v15), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v29 = base.F64_add(v27, v28)
				if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
					v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
					return v11
				} else {
					if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
						return v11
					} else {
						if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
							v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
							return v11
						}
					}
				}
			} else {
				if base.F64_ne(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v27 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
					v28 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
					v29 = base.F64_add(v27, v28)
					if base.F64_ne(base.F64_abs(v29), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
						return v11
					} else {
						if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
							v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
							return v11
						} else {
							if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
								v41 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v41
								return v11
							}
						}
					}
				}
			}
		}
	}
}
func F_circle_contain(m *base.Module, l0 int32) int32 {
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
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v14 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
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
func F_circle_div_pt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v23 float64
	_ = v23
	var v24 float64
	_ = v24
	var v27 float64
	_ = v27
	var v34 int32
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v39 float64
	_ = v39
	var v44 float64
	_ = v44
	var v51 float64
	_ = v51
	var v60 float64
	_ = v60
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		F_point_div_point(m, v11, v9, v8)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v9)+16))
			v18 = math.Float64frombits(uint64(0x7ff0000000000000))
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
			v20 = base.F64_abs(v19)
			if base.F64_eq(v20, v18) != 0 {
				v60 = v18
				v63 = base.F64_div(v17, v60)
				v65 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					if base.F64_ne(v63, float64(0)) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
						return v11
					} else {
						if base.F64_eq(v17, float64(0)) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
							return v11
						} else {
							if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
								return v11
							}
						}
					}
				}
			} else {
				v23 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
				v24 = base.F64_abs(v23)
				if base.F64_eq(v24, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					v60 = v18
					v63 = base.F64_div(v17, v60)
					v65 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						if base.F64_ne(v63, float64(0)) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
							return v11
						} else {
							if base.F64_eq(v17, float64(0)) != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
								return v11
							} else {
								if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
									return v11
								}
							}
						}
					}
				} else {
					v27 = math.Float64frombits(uint64(0x7ff8000000000000))
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v20)) {
						v60 = v27
						v63 = base.F64_div(v17, v60)
						v65 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							if base.F64_ne(v63, float64(0)) != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
								return v11
							} else {
								if base.F64_eq(v17, float64(0)) != 0 {
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
									return v11
								} else {
									if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
										return v11
									}
								}
							}
						}
					} else {
						if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v24)) {
							v60 = v27
							v63 = base.F64_div(v17, v60)
							v65 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.F64_ne(v63, float64(0)) != 0 {
									*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
									return v11
								} else {
									if base.F64_eq(v17, float64(0)) != 0 {
										*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
										return v11
									} else {
										if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
											return v11
										}
									}
								}
							}
						} else {
							v34 = base.F64_lt(v20, v24)
							if v34 != 0 {
								v35 = v24
							} else {
								v35 = v20
							}
							if v34 != 0 {
								v36 = v20
							} else {
								v36 = v24
							}
							if base.F64_ne(v36, float64(0)) != 0 {
								v39 = base.F64_div(v36, v35)
								v44 = base.F64_mul(v35, base.F64_sqrt(base.F64_add(base.F64_mul(v39, v39), float64(1))))
								if base.F64_eq(base.F64_abs(v44), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if base.F64_eq(v44, float64(0)) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v51 = v44
										if base.F64_ne(v51, float64(0)) != 0 {
											v60 = v51
											v63 = base.F64_div(v17, v60)
											v65 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_ne(v63, float64(0)) != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
													return v11
												} else {
													if base.F64_eq(v17, float64(0)) != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
														return v11
													} else {
														if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
															F_float_underflow_error(m)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														} else {
															*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
															return v11
														}
													}
												}
											}
										} else {
											if base.Ui64(base.I64_reinterpret_f64(v17)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
												F_float_zero_divide_error(m)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v60 = v51
												v63 = base.F64_div(v17, v60)
												v65 = math.Float64frombits(uint64(0x7ff0000000000000))
												if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
													F_float_overflow_error(m)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													if base.F64_ne(v63, float64(0)) != 0 {
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
														return v11
													} else {
														if base.F64_eq(v17, float64(0)) != 0 {
															*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
															return v11
														} else {
															if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
																F_float_underflow_error(m)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															} else {
																*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
																return v11
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v51 = v35
								if base.F64_ne(v51, float64(0)) != 0 {
									v60 = v51
									v63 = base.F64_div(v17, v60)
									v65 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_ne(v63, float64(0)) != 0 {
											*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
											return v11
										} else {
											if base.F64_eq(v17, float64(0)) != 0 {
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
												return v11
											} else {
												if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
													return v11
												}
											}
										}
									}
								} else {
									if base.Ui64(base.I64_reinterpret_f64(v17)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
										F_float_zero_divide_error(m)
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v60 = v51
										v63 = base.F64_div(v17, v60)
										v65 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_ne(v63, float64(0)) != 0 {
												*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
												return v11
											} else {
												if base.F64_eq(v17, float64(0)) != 0 {
													*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
													return v11
												} else {
													if base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														F_float_underflow_error(m)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v63
														return v11
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
func F_circle_eq(m *base.Module, l0 int32) int32 {
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
									return base.F64_eq(v51, v24) | base.F64_le(base.F64_abs(base.F64_sub(v24, v51)), float64(1e-06))
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_circle_ge(m *base.Module, l0 int32) int32 {
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
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v10 = base.F64_mul(v9, v9)
	v11 = base.F64_abs(v10)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v11, v12)&base.F64_ne(base.F64_abs(v9), v12) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
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
			v80 = m.ExcPending
			if v80 != 0 {
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
				v74 = m.ExcPending
				if v74 != 0 {
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
					v80 = m.ExcPending
					if v80 != 0 {
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
						v74 = m.ExcPending
						if v74 != 0 {
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
							v80 = m.ExcPending
							if v80 != 0 {
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
								v74 = m.ExcPending
								if v74 != 0 {
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
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									return base.F64_ge(base.F64_add(v24, float64(1e-06)), v51)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_circle_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 float64
	_ = v27
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 float64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_initStringInfo(m, v8+int32(16))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_appendStringInfoChar(m, v8+int32(16), int32(60))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_appendStringInfoChar(m, v8+int32(16), int32(40))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*float64)(unsafe.Add(mBase, uint32(v10)+8))
				v28 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
				v29 = F_float8out_internal(m, v28)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = F_float8out_internal(m, v27)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v31
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
						F_appendStringInfo(m, v8+int32(16), int32(167206), v8)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v29)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v31)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									F_appendStringInfoChar(m, v8+int32(16), int32(41))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										F_appendStringInfoChar(m, v8+int32(16), int32(44))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											v56 = *(*float64)(unsafe.Add(mBase, uint32(v10)+16))
											v57 = F_float8out_internal(m, v56)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												F_appendStringInfoString(m, v8+int32(16), v57)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v57)
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_appendStringInfoChar(m, v8+int32(16), int32(62))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
															m.G0 = v8 + int32(32)
															return v68
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
func F_circle_poly(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 float64
	_ = v43
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v50 float64
	_ = v50
	var v61 int32
	_ = v61
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v74 float64
	_ = v74
	var v85 int32
	_ = v85
	var v99 float64
	_ = v99
	var v104 float64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v124 float64
	_ = v124
	var v128 int32
	_ = v128
	var v129 float64
	_ = v129
	var v130 float64
	_ = v130
	var v135 float64
	_ = v135
	var v137 float64
	_ = v137
	var v139 float64
	_ = v139
	var v142 float64
	_ = v142
	var v146 float64
	_ = v146
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v165 float64
	_ = v165
	var v176 int32
	_ = v176
	var v178 float64
	_ = v178
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v198 float64
	_ = v198
	var v202 int32
	_ = v202
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v213 float64
	_ = v213
	var v215 float64
	_ = v215
	var v217 float64
	_ = v217
	var v223 float64
	_ = v223
	var v224 float64
	_ = v224
	var v225 float64
	_ = v225
	var v226 float64
	_ = v226
	var v238 float64
	_ = v238
	var v249 int32
	_ = v249
	var v251 float64
	_ = v251
	var v252 float64
	_ = v252
	var v253 int32
	_ = v253
	var v261 float64
	_ = v261
	var v262 float64
	_ = v262
	var v263 float64
	_ = v263
	var v264 float64
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v280 int32
	_ = v280
	var v284 float64
	_ = v284
	var v290 float64
	_ = v290
	var v291 float64
	_ = v291
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v296 int64
	_ = v296
	var v300 float64
	_ = v300
	var v305 int32
	_ = v305
	var v309 float64
	_ = v309
	var v315 float64
	_ = v315
	var v316 float64
	_ = v316
	var v317 float64
	_ = v317
	var v318 float64
	_ = v318
	var v320 float64
	_ = v320
	var v326 float64
	_ = v326
	var v328 int32
	_ = v328
	var v332 float64
	_ = v332
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v335 float64
	_ = v335
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v410 int32
	_ = v410
	var v425 int32
	_ = v425
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	if base.F64_le(base.F64_abs(v15), float64(1e-06)) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L11
	} else {
		goto L127
	}
L2:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L11
	} else {
		goto L126
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L11
	} else {
		goto L122
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L11
	} else {
		goto L118
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v21 <= int32(1) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L11
	} else {
		goto L114
	}
L8:
	;
	v25 = v21 << (uint(int32(4)) % 32)
	v26 = base.I32_div_s(v25, v21)
	if int32(2147483607) < v25 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	if v26 != int32(16) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v32 = v25 + int32(40)
	v33 = F_palloc0(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v32 << (uint(int32(2)) % 32)
	v43 = base.F64_div(float64(6.283185307179586), base.F64_convert_i32_u(v21))
	if base.F64_eq(v43, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	if base.F64_eq(v43, float64(0)) != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v48 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
	v49 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v50 = base.F64_sub(v48, v49)
	if base.F64_ne(base.F64_abs(v50), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = v33 + int32(40)
	*(*float64)(unsafe.Add(mBase, uint32(v61))) = v50
	v63 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v66 = base.F64_mul(v64, float64(0))
	v67 = base.F64_abs(v66)
	v68 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v67, v68)&base.F64_ne(base.F64_abs(v64), v68) != 0 {
		goto L2
	} else {
		goto L19
	}
L16:
	;
	if base.F64_eq(base.F64_abs(v48), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if base.F64_ne(base.F64_abs(v49), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v74 = base.F64_add(v63, v66)
	if base.F64_ne(base.F64_abs(v74), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v33)+48)) = v74
	v85 = int32(1)
	goto L24
L21:
	;
	if base.F64_eq(base.F64_abs(v63), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	if base.F64_ne(v67, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L20
L24:
	;
	v99 = base.F64_mul(v43, base.F64_convert_i32_u(v85))
	if base.F64_eq(v99, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L26
	}
L25:
	;
	v251 = *(*float64)(unsafe.Add(mBase, uint32(v33)+48))
	v252 = *(*float64)(unsafe.Add(mBase, uint32(v33)+40))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v253 < int32(2) {
		goto L72
	} else {
		goto L73
	}
L26:
	;
	if base.F64_eq(v99, float64(0)) != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v104 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v108 = m.G0
	v110 = v108 - int32(16)
	m.G0 = v110
	v117 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v99))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v117) <= base.Ui32(int32(1072243195)) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
	v151 = base.F64_mul(v146, v104)
	v152 = base.F64_abs(v151)
	v153 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v152, v153)&base.F64_ne(base.F64_abs(v104), v153) != 0 {
		goto L2
	} else {
		goto L39
	}
L29:
	;
	m.G0 = v110 + int32(16)
	goto L28
L30:
	;
	if base.Ui32(v117) < base.Ui32(int32(1044816030)) {
		v146 = float64(1)
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v117) {
		v146 = base.F64_sub(v99, v99)
		goto L29
	} else {
		goto L34
	}
L33:
	;
	v124 = F___cos(m, v99, float64(0))
	mBase = m.M
	v146 = v124
	goto L29
L34:
	;
	v128 = F___rem_pio2(m, v99, v110)
	mBase = m.M
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v110)+8))
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v110)))
	switch v128&int32(3) - int32(1) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	default:
		goto L38
	}
L35:
	;
	v142 = F___sin(m, v130, v129, int32(1))
	mBase = m.M
	v146 = v142
	goto L29
L36:
	;
	v139 = F___cos(m, v130, v129)
	mBase = m.M
	v146 = base.F64_neg(v139)
	goto L29
L37:
	;
	v137 = F___sin(m, v130, v129, int32(1))
	mBase = m.M
	v146 = base.F64_neg(v137)
	goto L29
L38:
	;
	v135 = F___cos(m, v130, v129)
	mBase = m.M
	v146 = v135
	goto L29
L39:
	;
	if base.F64_ne(v151, float64(0)) != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v165 = base.F64_sub(v150, v151)
	if base.F64_ne(base.F64_abs(v165), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	if base.F64_eq(v146, float64(0)) != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	if base.F64_ne(v104, float64(0)) != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v176 = v61 + v85<<(uint(int32(4))%32)
	*(*float64)(unsafe.Add(mBase, uint32(v176))) = v165
	v178 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v182 = m.G0
	v184 = v182 - int32(16)
	m.G0 = v184
	v191 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v99))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v191) <= base.Ui32(int32(1072243195)) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	if base.F64_eq(base.F64_abs(v150), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	if base.F64_ne(v152, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v223 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	v224 = base.F64_mul(v217, v178)
	v225 = base.F64_abs(v224)
	v226 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v225, v226)&base.F64_ne(base.F64_abs(v178), v226) != 0 {
		goto L2
	} else {
		goto L61
	}
L49:
	;
	m.G0 = v184 + int32(16)
	goto L48
L50:
	;
	if base.Ui32(v191) < base.Ui32(int32(1045430272)) {
		v217 = v99
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v191) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v198 = F___sin(m, v99, float64(0), int32(0))
	mBase = m.M
	v217 = v198
	goto L49
L54:
	;
	v217 = base.F64_sub(v99, v99)
	goto L49
L55:
	;
	goto L56
L56:
	;
	v202 = F___rem_pio2(m, v99, v184)
	mBase = m.M
	v203 = *(*float64)(unsafe.Add(mBase, uint32(v184)+8))
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v184)))
	switch v202&int32(3) - int32(1) {
	case 0:
		goto L59
	case 1:
		goto L58
	case 2:
		goto L57
	default:
		goto L60
	}
L57:
	;
	v215 = F___cos(m, v204, v203)
	mBase = m.M
	v217 = base.F64_neg(v215)
	goto L49
L58:
	;
	v213 = F___sin(m, v204, v203, int32(1))
	mBase = m.M
	v217 = base.F64_neg(v213)
	goto L49
L59:
	;
	v211 = F___cos(m, v204, v203)
	mBase = m.M
	v217 = v211
	goto L49
L60:
	;
	v210 = F___sin(m, v204, v203, int32(1))
	mBase = m.M
	v217 = v210
	goto L49
L61:
	;
	if base.F64_ne(v224, float64(0)) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v238 = base.F64_add(v223, v224)
	if base.F64_ne(base.F64_abs(v238), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	if base.F64_eq(v217, float64(0)) != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	if base.F64_ne(v178, float64(0)) != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v176)+8)) = v238
	v249 = v85 + int32(1)
	if v249 != v21 {
		v85 = v249
		goto L24
	} else {
		goto L70
	}
L67:
	;
	if base.F64_eq(base.F64_abs(v223), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	if base.F64_ne(v225, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	goto L25
L71:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v33)+32)) = v335
	*(*float64)(unsafe.Add(mBase, uint32(v33)+8)) = v333
	*(*float64)(unsafe.Add(mBase, uint32(v33)+24)) = v334
	*(*float64)(unsafe.Add(mBase, uint32(v33)+16)) = v332
	return v33
L72:
	;
	v332 = v251
	v333 = v252
	v334 = v252
	v335 = v251
	goto L71
L73:
	;
	goto L74
L74:
	;
	v261 = v251
	v262 = v252
	v263 = v252
	v264 = v251
	v267 = int32(1)
	goto L75
L75:
	;
	v274 = v33 + int32(40) + v267<<(uint(int32(4))%32)
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)))
	v280 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v275)&int64(9223372036854775807)))
	if v280 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v332 = v326
	v333 = v317
	v334 = v291
	v335 = v316
	goto L71
L77:
	;
	if base.F64_lt(v275, v263) != 0 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v291 = v263
	goto L79
L79:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v275)&int64(9223372036854775807)) {
		goto L86
	} else {
		goto L87
	}
L80:
	;
	v284 = v275
	goto L82
L81:
	;
	v284 = v263
	goto L82
L82:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v263)&int64(9223372036854775807)) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v290 = v275
	goto L85
L84:
	;
	v290 = v284
	goto L85
L85:
	;
	v291 = v290
	goto L79
L86:
	;
	v292 = v275
	goto L88
L87:
	;
	v292 = v262
	goto L88
L88:
	;
	if base.F64_gt(v275, v262) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v294 = v275
	goto L91
L90:
	;
	v294 = v292
	goto L91
L91:
	;
	v296 = int64(9223372036854775807)
	v300 = *(*float64)(unsafe.Add(mBase, uint32(v274)+8))
	v305 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v300)&v296))
	if v305 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if base.F64_lt(v300, v264) != 0 {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v316 = v264
	goto L94
L94:
	;
	if base.Ui64(base.I64_reinterpret_f64(v262)&v296) < base.Ui64(int64(9218868437227405313)) {
		goto L101
	} else {
		goto L102
	}
L95:
	;
	v309 = v300
	goto L97
L96:
	;
	v309 = v264
	goto L97
L97:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v264)&int64(9223372036854775807)) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v315 = v300
	goto L100
L99:
	;
	v315 = v309
	goto L100
L100:
	;
	v316 = v315
	goto L94
L101:
	;
	v317 = v294
	goto L103
L102:
	;
	v317 = v262
	goto L103
L103:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v300)&v296) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v318 = v300
	goto L106
L105:
	;
	v318 = v261
	goto L106
L106:
	;
	if base.F64_gt(v300, v261) != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v320 = v300
	goto L109
L108:
	;
	v320 = v318
	goto L109
L109:
	;
	if base.Ui64(base.I64_reinterpret_f64(v261)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v326 = v320
	goto L112
L111:
	;
	v326 = v261
	goto L112
L112:
	;
	v328 = v267 + int32(1)
	if v328 != v253 {
		v261 = v326
		v262 = v317
		v263 = v291
		v264 = v316
		v267 = v328
		goto L75
	} else {
		goto L113
	}
L113:
	;
	goto L76
L114:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	F_errmsg(m, int32(259647), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(470966), int32(5239), int32(18124))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(111432), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	F_errfinish(m, int32(470966), int32(5244), int32(18124))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L122:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(420347), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(470966), int32(5253), int32(18124))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_circle_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v11 int32
	_ = v11
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v16 float64
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_pq_getmsgfloat8(m, v4)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			*(*float64)(unsafe.Add(mBase, uint32(v6))) = v10
			v13 = F_pq_getmsgfloat8(m, v4)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v13
				v16 = F_pq_getmsgfloat8(m, v4)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v6)+16)) = v16
					if base.F64_lt(v16, float64(0)) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50462850))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(330283), int32(0))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(470966), int32(4718), int32(34204))
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						return v6
					}
				}
			}
		}
	}
}
func F_circle_right(m *base.Module, l0 int32) int32 {
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
	v11 = base.F64_sub(v9, v10)
	if base.F64_ne(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
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
			v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
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
				v21 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
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
