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
	var v19 float64
	_ = v19
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v49 float64
	_ = v49
	var v56 int32
	_ = v56
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_palloc(m, int32(24))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
		v16 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
		v17 = base.F64_add(v15, v16)
		v19 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v17), v19)|base.F64_eq(base.F64_abs(v15), v19) == int32(0))&base.F64_ne(base.F64_abs(v16), v19) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v56 = m.ExcPending
			if v56 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v31 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v9)+8))
			v33 = base.F64_add(v31, v32)
			v35 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(base.F64_abs(v33), v35)|base.F64_eq(base.F64_abs(v31), v35) == int32(0))&base.F64_ne(base.F64_abs(v32), v35) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v33
				*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
				v49 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v49
				return v11
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
	var v17 float64
	_ = v17
	var v30 int32
	_ = v30
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
		v17 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.F64_ne(base.F64_abs(v15), v17)|base.F64_eq(base.F64_abs(v13), v17)|base.F64_eq(base.F64_abs(v14), v17) == int32(0) {
			F_float_overflow_error(m)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.F64_le(v9, base.F64_add(v15, float64(1e-06)))
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
	var v29 int64
	_ = v29
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v45 float64
	_ = v45
	var v51 float64
	_ = v51
	var v60 float64
	_ = v60
	var v63 float64
	_ = v63
	var v65 float64
	_ = v65
	var v71 float64
	_ = v71
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
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
					v89 = m.ExcPending
					if v89 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v71 = float64(0)
					if base.B2i32(base.F64_eq(v17, v71)|base.F64_ne(v63, v71) == int32(0))&base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_underflow_error(m)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
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
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						v71 = float64(0)
						if base.B2i32(base.F64_eq(v17, v71)|base.F64_ne(v63, v71) == int32(0))&base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_underflow_error(m)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
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
				} else {
					v29 = int64(9218868437227405312)
					if base.B2i32(base.Ui64(v29) < base.Ui64(base.I64_reinterpret_f64(v20)))|base.B2i32(base.Ui64(v29) < base.Ui64(base.I64_reinterpret_f64(v24))) != 0 {
						v60 = math.Float64frombits(uint64(0x7ff8000000000000))
						v63 = base.F64_div(v17, v60)
						v65 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							v71 = float64(0)
							if base.B2i32(base.F64_eq(v17, v71)|base.F64_ne(v63, v71) == int32(0))&base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_underflow_error(m)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
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
					} else {
						v35 = base.F64_lt(v20, v24)
						if v35 != 0 {
							v36 = v24
						} else {
							v36 = v20
						}
						if v35 != 0 {
							v37 = v20
						} else {
							v37 = v24
						}
						if base.F64_ne(v37, float64(0)) != 0 {
							v40 = base.F64_div(v37, v36)
							v45 = base.F64_mul(v36, base.F64_sqrt(base.F64_add(base.F64_mul(v40, v40), float64(1))))
							if base.F64_eq(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.F64_eq(v45, float64(0)) != 0 {
									F_float_underflow_error(m)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v51 = v45
									if base.F64_ne(v51, float64(0)) != 0 {
										v60 = v51
										v63 = base.F64_div(v17, v60)
										v65 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v71 = float64(0)
											if base.B2i32(base.F64_eq(v17, v71)|base.F64_ne(v63, v71) == int32(0))&base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
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
									} else {
										if base.Ui64(base.I64_reinterpret_f64(v17)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
											F_float_zero_divide_error(m)
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
											v60 = v51
											v63 = base.F64_div(v17, v60)
											v65 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v71 = float64(0)
												if base.B2i32(base.F64_eq(v17, v71)|base.F64_ne(v63, v71) == int32(0))&base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
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
						} else {
							v51 = v36
							if base.F64_ne(v51, float64(0)) != 0 {
								v60 = v51
								v63 = base.F64_div(v17, v60)
								v65 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v71 = float64(0)
									if base.B2i32(base.F64_eq(v17, v71)|base.F64_ne(v63, v71) == int32(0))&base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
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
							} else {
								if base.Ui64(base.I64_reinterpret_f64(v17)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
									F_float_zero_divide_error(m)
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
									v60 = v51
									v63 = base.F64_div(v17, v60)
									v65 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v63), v65)&base.F64_ne(base.F64_abs(v17), v65) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v71 = float64(0)
										if base.B2i32(base.F64_eq(v17, v71)|base.F64_ne(v63, v71) == int32(0))&base.F64_ne(base.F64_abs(v60), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = v9 + int32(16)
	F_initStringInfo(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		F_appendStringInfoChar(m, v13, int32(60))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			F_appendStringInfoChar(m, v13, int32(40))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = *(*float64)(unsafe.Add(mBase, uint32(v11)+8))
				v25 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
				v26 = F_float8out_internal(m, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = F_float8out_internal(m, v24)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v28
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v26
						F_appendStringInfo(m, v13, int32(_a_F_circle_out_0), v9)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v26)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v28)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									F_appendStringInfoChar(m, v13, int32(41))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										F_appendStringInfoChar(m, v13, int32(44))
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
											return int32(0)
										} else {
											v45 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
											v46 = F_float8out_internal(m, v45)
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return int32(0)
											} else {
												F_appendStringInfoString(m, v13, v46)
												mBase = m.M
												v49 = m.ExcPending
												if v49 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v46)
													mBase = m.M
													v51 = m.ExcPending
													if v51 != 0 {
														return int32(0)
													} else {
														F_appendStringInfoChar(m, v13, int32(62))
														mBase = m.M
														v54 = m.ExcPending
														if v54 != 0 {
															return int32(0)
														} else {
															v55 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
															m.G0 = v9 + int32(32)
															return v55
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
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 float64
	_ = v46
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v68 int32
	_ = v68
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v84 int32
	_ = v84
	var v100 float64
	_ = v100
	var v105 float64
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 float64
	_ = v125
	var v129 int32
	_ = v129
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v136 float64
	_ = v136
	var v138 float64
	_ = v138
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v147 float64
	_ = v147
	var v151 float64
	_ = v151
	var v152 float64
	_ = v152
	var v153 float64
	_ = v153
	var v154 float64
	_ = v154
	var v160 float64
	_ = v160
	var v171 float64
	_ = v171
	var v173 float64
	_ = v173
	var v185 int32
	_ = v185
	var v187 float64
	_ = v187
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v207 float64
	_ = v207
	var v211 int32
	_ = v211
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v219 float64
	_ = v219
	var v220 float64
	_ = v220
	var v222 float64
	_ = v222
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v235 float64
	_ = v235
	var v241 float64
	_ = v241
	var v252 float64
	_ = v252
	var v254 float64
	_ = v254
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 float64
	_ = v272
	var v273 float64
	_ = v273
	var v275 float64
	_ = v275
	var v277 float64
	_ = v277
	var v288 int32
	_ = v288
	var v289 float64
	_ = v289
	var v294 int32
	_ = v294
	var v298 float64
	_ = v298
	var v304 float64
	_ = v304
	var v305 float64
	_ = v305
	var v306 float64
	_ = v306
	var v311 int32
	_ = v311
	var v315 float64
	_ = v315
	var v321 float64
	_ = v321
	var v322 float64
	_ = v322
	var v323 float64
	_ = v323
	var v325 float64
	_ = v325
	var v331 float64
	_ = v331
	var v332 float64
	_ = v332
	var v334 float64
	_ = v334
	var v340 float64
	_ = v340
	var v342 int32
	_ = v342
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v413 int32
	_ = v413
	var v430 int32
	_ = v430
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	if base.F64_le(base.F64_abs(v17), float64(1e-06)) == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_float_underflow_error(m)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L10
	} else {
		goto L103
	}
L2:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L10
	} else {
		goto L102
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L10
	} else {
		goto L98
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L10
	} else {
		goto L94
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v23 <= int32(1) {
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
	v352 = m.ExcPending
	if v352 != 0 {
		goto L10
	} else {
		goto L90
	}
L8:
	;
	v27 = v23 << (uint(int32(4)) % 32)
	v28 = base.I32_div_s(v27, v23)
	if base.B2i32(v28 != int32(16))|base.B2i32(int32(2147483607) < v27) != 0 {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v35 = v27 + int32(40)
	v36 = F_palloc0(m, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v35 << (uint(int32(2)) % 32)
	v46 = base.F64_div(float64(6.283185307179586), base.F64_convert_i32_u(v23))
	if base.F64_eq(v46, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	if base.F64_eq(v46, float64(0)) != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v51 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	v53 = base.F64_sub(v51, v52)
	v55 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v53), v55)|base.F64_eq(base.F64_abs(v51), v55) == int32(0))&base.F64_ne(base.F64_abs(v52), v55) != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v68 = v36 + int32(40)
	*(*float64)(unsafe.Add(mBase, uint32(v68))) = v53
	v70 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v71 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	v74 = base.F64_add(v70, base.F64_mul(v71, float64(0)))
	v76 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v74), v76)&base.F64_ne(base.F64_abs(v70), v76) != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v36)+48)) = v74
	v84 = int32(1)
	goto L16
L16:
	;
	v100 = base.F64_mul(v46, base.F64_convert_i32_u(v84))
	if base.F64_eq(v100, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L18
	}
L17:
	;
	v271 = int32(1)
	v272 = v74
	v273 = v53
	v275 = v74
	v277 = v53
	goto L51
L18:
	;
	if base.F64_eq(v100, float64(0)) != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	v109 = m.G0
	v111 = v109 - int32(16)
	m.G0 = v111
	v118 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v100))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v118) <= base.Ui32(int32(1072243195)) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v151 = *(*float64)(unsafe.Add(mBase, uint32(v16)))
	v152 = base.F64_mul(v147, v105)
	v153 = base.F64_abs(v152)
	v154 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v153, v154)&base.F64_ne(base.F64_abs(v105), v154) != 0 {
		goto L2
	} else {
		goto L31
	}
L21:
	;
	m.G0 = v111 + int32(16)
	goto L20
L22:
	;
	if base.Ui32(v118) < base.Ui32(int32(1044816030)) {
		v147 = float64(1)
		goto L21
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v118) {
		v147 = base.F64_sub(v100, v100)
		goto L21
	} else {
		goto L26
	}
L25:
	;
	v125 = F___cos(m, v100, float64(0))
	mBase = m.M
	v147 = v125
	goto L21
L26:
	;
	v129 = F___rem_pio2(m, v100, v111)
	mBase = m.M
	v130 = *(*float64)(unsafe.Add(mBase, uint32(v111)+8))
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v111)))
	switch v129&int32(3) - int32(1) {
	case 0:
		goto L29
	case 1:
		goto L28
	case 2:
		goto L27
	default:
		goto L30
	}
L27:
	;
	v143 = F___sin(m, v131, v130, int32(1))
	mBase = m.M
	v147 = v143
	goto L21
L28:
	;
	v140 = F___cos(m, v131, v130)
	mBase = m.M
	v147 = base.F64_neg(v140)
	goto L21
L29:
	;
	v138 = F___sin(m, v131, v130, int32(1))
	mBase = m.M
	v147 = base.F64_neg(v138)
	goto L21
L30:
	;
	v136 = F___cos(m, v131, v130)
	mBase = m.M
	v147 = v136
	goto L21
L31:
	;
	v160 = float64(0)
	if base.B2i32(base.F64_eq(v147, v160)|base.F64_ne(v152, v160) == int32(0))&base.F64_ne(v105, v160) != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v171 = math.Float64frombits(uint64(0x7ff0000000000000))
	v173 = base.F64_sub(v151, v152)
	if base.B2i32(base.F64_eq(base.F64_abs(v151), v171)|base.F64_ne(base.F64_abs(v173), v171) == int32(0))&base.F64_ne(v153, v171) != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v185 = v68 + v84<<(uint(int32(4))%32)
	*(*float64)(unsafe.Add(mBase, uint32(v185))) = v173
	v187 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	v191 = m.G0
	v193 = v191 - int32(16)
	m.G0 = v193
	v200 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v100))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v200) <= base.Ui32(int32(1072243195)) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	v233 = base.F64_mul(v226, v187)
	v234 = base.F64_abs(v233)
	v235 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(v234, v235)&base.F64_ne(base.F64_abs(v187), v235) != 0 {
		goto L2
	} else {
		goto L47
	}
L35:
	;
	m.G0 = v193 + int32(16)
	goto L34
L36:
	;
	if base.Ui32(v200) < base.Ui32(int32(1045430272)) {
		v226 = v100
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v200) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v207 = F___sin(m, v100, float64(0), int32(0))
	mBase = m.M
	v226 = v207
	goto L35
L40:
	;
	v226 = base.F64_sub(v100, v100)
	goto L35
L41:
	;
	goto L42
L42:
	;
	v211 = F___rem_pio2(m, v100, v193)
	mBase = m.M
	v212 = *(*float64)(unsafe.Add(mBase, uint32(v193)+8))
	v213 = *(*float64)(unsafe.Add(mBase, uint32(v193)))
	switch v211&int32(3) - int32(1) {
	case 0:
		goto L45
	case 1:
		goto L44
	case 2:
		goto L43
	default:
		goto L46
	}
L43:
	;
	v224 = F___cos(m, v213, v212)
	mBase = m.M
	v226 = base.F64_neg(v224)
	goto L35
L44:
	;
	v222 = F___sin(m, v213, v212, int32(1))
	mBase = m.M
	v226 = base.F64_neg(v222)
	goto L35
L45:
	;
	v220 = F___cos(m, v213, v212)
	mBase = m.M
	v226 = v220
	goto L35
L46:
	;
	v219 = F___sin(m, v213, v212, int32(1))
	mBase = m.M
	v226 = v219
	goto L35
L47:
	;
	v241 = float64(0)
	if base.B2i32(base.F64_eq(v226, v241)|base.F64_ne(v233, v241) == int32(0))&base.F64_ne(v187, v241) != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v252 = math.Float64frombits(uint64(0x7ff0000000000000))
	v254 = base.F64_add(v232, v233)
	if base.B2i32(base.F64_eq(base.F64_abs(v232), v252)|base.F64_ne(base.F64_abs(v254), v252) == int32(0))&base.F64_ne(v234, v252) != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v185)+8)) = v254
	v266 = v84 + int32(1)
	if v266 != v23 {
		v84 = v266
		goto L16
	} else {
		goto L50
	}
L50:
	;
	goto L17
L51:
	;
	v288 = v36 + int32(40) + v271<<(uint(int32(4))%32)
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v288)))
	v294 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v289)&int64(9223372036854775807)))
	if v294 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v36)+32)) = v322
	*(*float64)(unsafe.Add(mBase, uint32(v36)+8)) = v331
	*(*float64)(unsafe.Add(mBase, uint32(v36)+24)) = v305
	*(*float64)(unsafe.Add(mBase, uint32(v36)+16)) = v340
	return v36
L53:
	;
	if base.F64_gt(v273, v289) != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v305 = v273
	goto L55
L55:
	;
	v306 = *(*float64)(unsafe.Add(mBase, uint32(v288)+8))
	v311 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v306)&int64(9223372036854775807)))
	if v311 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	v298 = v289
	goto L58
L57:
	;
	v298 = v273
	goto L58
L58:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v273)&int64(9223372036854775807)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v304 = v289
	goto L61
L60:
	;
	v304 = v298
	goto L61
L61:
	;
	v305 = v304
	goto L55
L62:
	;
	if base.F64_gt(v272, v306) != 0 {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v322 = v272
	goto L64
L64:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v289)&int64(9223372036854775807)) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v315 = v306
	goto L67
L66:
	;
	v315 = v272
	goto L67
L67:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v272)&int64(9223372036854775807)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v321 = v306
	goto L70
L69:
	;
	v321 = v315
	goto L70
L70:
	;
	v322 = v321
	goto L64
L71:
	;
	v323 = v289
	goto L73
L72:
	;
	v323 = v277
	goto L73
L73:
	;
	if base.F64_gt(v289, v277) != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v325 = v289
	goto L76
L75:
	;
	v325 = v323
	goto L76
L76:
	;
	if base.Ui64(base.I64_reinterpret_f64(v277)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v331 = v325
	goto L79
L78:
	;
	v331 = v277
	goto L79
L79:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v306)&int64(9223372036854775807)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v332 = v306
	goto L82
L81:
	;
	v332 = v275
	goto L82
L82:
	;
	if base.F64_gt(v306, v275) != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v334 = v306
	goto L85
L84:
	;
	v334 = v332
	goto L85
L85:
	;
	if base.Ui64(base.I64_reinterpret_f64(v275)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v340 = v334
	goto L88
L87:
	;
	v340 = v275
	goto L88
L88:
	;
	v342 = v271 + int32(1)
	if v342 != v23 {
		v271 = v342
		v272 = v322
		v273 = v305
		v275 = v340
		v277 = v331
		goto L51
	} else {
		goto L89
	}
L89:
	;
	goto L52
L90:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	F_errmsg(m, int32(_a_F_circle_poly_0), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	F_errfinish(m, int32(_a_F_circle_poly_1), int32(_a_F_circle_poly_2), int32(_a_F_circle_poly_3))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L94:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_circle_poly_4), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_circle_poly_1), int32(_a_F_circle_poly_5), int32(_a_F_circle_poly_3))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L10
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(_a_F_circle_poly_6), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L10
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_circle_poly_1), int32(_a_F_circle_poly_7), int32(_a_F_circle_poly_3))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L10
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
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
								F_errmsg(m, int32(_a_F_circle_recv_0), int32(0))
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_circle_recv_1), int32(_a_F_circle_recv_2), int32(_a_F_circle_recv_3))
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
	var v13 float64
	_ = v13
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v51 int32
	_ = v51
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*float64)(unsafe.Add(mBase, uint32(v8)))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
	v11 = base.F64_sub(v9, v10)
	v13 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.B2i32(base.F64_ne(base.F64_abs(v11), v13)|base.F64_eq(base.F64_abs(v9), v13) == int32(0))&base.F64_ne(base.F64_abs(v10), v13) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v25 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		v26 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v27 = base.F64_add(v25, v26)
		v29 = math.Float64frombits(uint64(0x7ff0000000000000))
		if base.B2i32(base.F64_ne(base.F64_abs(v27), v29)|base.F64_eq(base.F64_abs(v25), v29) == int32(0))&base.F64_ne(base.F64_abs(v26), v29) != 0 {
			F_float_overflow_error(m)
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.F64_gt(v11, base.F64_add(v27, float64(1e-06)))
		}
	}
}
