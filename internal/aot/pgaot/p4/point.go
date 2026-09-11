package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_point_add(m *base.Module, l0 int32) int32 {
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
	var v46 int32
	_ = v46
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_palloc(m, int32(16))
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
				return v11
			} else {
				if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
					return v11
				} else {
					if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
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
					return v11
				} else {
					if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
						*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
						return v11
					} else {
						if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
							return v11
						}
					}
				}
			} else {
				if base.F64_ne(base.F64_abs(v16), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
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
						return v11
					} else {
						if base.F64_eq(base.F64_abs(v27), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
							return v11
						} else {
							if base.F64_ne(base.F64_abs(v28), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
								*(*float64)(unsafe.Add(mBase, uint32(v11))) = v17
								return v11
							}
						}
					}
				}
			}
		}
	}
}
func F_point_div(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		F_point_div_point(m, v7, v5, v4)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_point_sl(m *base.Module, l0 int32, l1 int32) float64 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 float64
	_ = v22
	var v24 float64
	_ = v24
	var v25 float64
	_ = v25
	var v44 float64
	_ = v44
	var v46 float64
	_ = v46
	var v55 float64
	_ = v55
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	if base.F64_eq(v13, v14) != 0 {
		v55 = v12
		return v55
	} else {
		v16 = base.F64_sub(v13, v14)
		v17 = base.F64_abs(v16)
		if base.F64_le(v17, float64(1e-06)) != 0 {
			v55 = v12
			return v55
		} else {
			v20 = float64(0)
			v21 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
			v22 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
			if base.F64_eq(v21, v22) != 0 {
				v55 = v20
				return v55
			} else {
				v24 = base.F64_sub(v21, v22)
				v25 = base.F64_abs(v24)
				if base.F64_le(v25, float64(1e-06)) != 0 {
					v55 = v20
					return v55
				} else {
					if base.F64_ne(v25, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
						if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							v44 = base.F64_div(v24, v16)
							v46 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.F64_ne(v44, float64(0)) != 0 {
									v55 = v44
									return v55
								} else {
									if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_underflow_error(m)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v55 = v44
										return v55
									}
								}
							}
						} else {
							if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v44 = base.F64_div(v24, v16)
								v46 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if base.F64_ne(v44, float64(0)) != 0 {
										v55 = v44
										return v55
									} else {
										if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v55 = v44
											return v55
										}
									}
								}
							} else {
								if base.F64_ne(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									v44 = base.F64_div(v24, v16)
									v46 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_ne(v44, float64(0)) != 0 {
											v55 = v44
											return v55
										} else {
											if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v55 = v44
												return v55
											}
										}
									}
								}
							}
						}
					} else {
						if base.F64_eq(base.F64_abs(v21), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								v44 = base.F64_div(v24, v16)
								v46 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return float64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									if base.F64_ne(v44, float64(0)) != 0 {
										v55 = v44
										return v55
									} else {
										if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_underflow_error(m)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v55 = v44
											return v55
										}
									}
								}
							} else {
								if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v44 = base.F64_div(v24, v16)
									v46 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_ne(v44, float64(0)) != 0 {
											v55 = v44
											return v55
										} else {
											if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v55 = v44
												return v55
											}
										}
									}
								} else {
									if base.F64_ne(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										v44 = base.F64_div(v24, v16)
										v46 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_ne(v44, float64(0)) != 0 {
												v55 = v44
												return v55
											} else {
												if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return float64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v55 = v44
													return v55
												}
											}
										}
									}
								}
							}
						} else {
							if base.F64_ne(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								F_float_overflow_error(m)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							} else {
								if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									v44 = base.F64_div(v24, v16)
									v46 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
										F_float_overflow_error(m)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return float64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									} else {
										if base.F64_ne(v44, float64(0)) != 0 {
											v55 = v44
											return v55
										} else {
											if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
												F_float_underflow_error(m)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												v55 = v44
												return v55
											}
										}
									}
								} else {
									if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
										v44 = base.F64_div(v24, v16)
										v46 = math.Float64frombits(uint64(0x7ff0000000000000))
										if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											if base.F64_ne(v44, float64(0)) != 0 {
												v55 = v44
												return v55
											} else {
												if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
													F_float_underflow_error(m)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return float64(0)
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												} else {
													v55 = v44
													return v55
												}
											}
										}
									} else {
										if base.F64_ne(base.F64_abs(v14), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
											F_float_overflow_error(m)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return float64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										} else {
											v44 = base.F64_div(v24, v16)
											v46 = math.Float64frombits(uint64(0x7ff0000000000000))
											if base.F64_eq(base.F64_abs(v44), v46)&base.F64_ne(v25, v46) != 0 {
												F_float_overflow_error(m)
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return float64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											} else {
												if base.F64_ne(v44, float64(0)) != 0 {
													v55 = v44
													return v55
												} else {
													if base.F64_ne(v17, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
														F_float_underflow_error(m)
														mBase = m.M
														v69 = m.ExcPending
														if v69 != 0 {
															return float64(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													} else {
														v55 = v44
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
func F_point_vert(m *base.Module, l0 int32) int32 {
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
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	return base.F64_eq(v5, v7) | base.F64_le(base.F64_abs(base.F64_sub(v5, v7)), float64(1e-06))
}
