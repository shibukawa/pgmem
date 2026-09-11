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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 float64
	_ = v17
	var v20 int32
	_ = v20
	var v28 float64
	_ = v28
	var v36 float64
	_ = v36
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v51 float64
	_ = v51
	var v56 float64
	_ = v56
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 float64
	_ = v84
	var v85 int32
	_ = v85
	var v88 float64
	_ = v88
	var v89 int32
	_ = v89
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v4 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = l2 + int32(16)
	v17 = F_point_sl(m, l2, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		if base.F64_eq(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(-4616189618054758400)
			v28 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
			*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v28
			v78 = F_lseg_interpt_line(m, v13+int32(32), l1, v13+int32(8))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				if v78 == int32(0) {
					v107 = v4
					m.G0 = v13 + int32(48)
					return v107
				} else {
					v84 = F_point_dt(m, v13+int32(32), l2)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						v88 = F_point_dt(m, v13+int32(32), v16)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							v90 = base.F64_add(v84, v88)
							v91 = F_point_dt(m, l2, v16)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								if base.F64_ne(v90, v91) != 0 {
									if base.F64_le(base.F64_abs(base.F64_sub(v90, v91)), float64(1e-06)) == int32(0) {
										v107 = v4
									} else {
										v100 = int32(1)
										if l0 == int32(0) {
											v107 = v100
										} else {
											v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
											v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
											v107 = v100
										}
									}
								} else {
									v100 = int32(1)
									if l0 == int32(0) {
										v107 = v100
									} else {
										v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
										*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
										v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
										*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
										v107 = v100
									}
								}
								m.G0 = v13 + int32(48)
								return v107
							}
						}
					}
				}
			}
		} else {
			if base.F64_eq(v17, float64(0)) != 0 {
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(-4616189618054758400)
				*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
				v36 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v36
				v78 = F_lseg_interpt_line(m, v13+int32(32), l1, v13+int32(8))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					if v78 == int32(0) {
						v107 = v4
						m.G0 = v13 + int32(48)
						return v107
					} else {
						v84 = F_point_dt(m, v13+int32(32), l2)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							v88 = F_point_dt(m, v13+int32(32), v16)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int32(0)
							} else {
								v90 = base.F64_add(v84, v88)
								v91 = F_point_dt(m, l2, v16)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									if base.F64_ne(v90, v91) != 0 {
										if base.F64_le(base.F64_abs(base.F64_sub(v90, v91)), float64(1e-06)) == int32(0) {
											v107 = v4
										} else {
											v100 = int32(1)
											if l0 == int32(0) {
												v107 = v100
											} else {
												v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
												*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
												v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
												*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
												v107 = v100
											}
										}
									} else {
										v100 = int32(1)
										if l0 == int32(0) {
											v107 = v100
										} else {
											v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
											*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
											v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
											*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
											v107 = v100
										}
									}
									m.G0 = v13 + int32(48)
									return v107
								}
							}
						}
					}
				}
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = int64(-4616189618054758400)
				*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v17
				v41 = *(*float64)(unsafe.Add(mBase, uint32(l2)+8))
				v42 = *(*float64)(unsafe.Add(mBase, uint32(l2)))
				v43 = base.F64_mul(v17, v42)
				v44 = base.F64_abs(v43)
				v45 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_eq(v44, v45)&base.F64_ne(base.F64_abs(v42), v45) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v51 = float64(0)
					if base.F64_eq(v43, v51)&base.F64_ne(v42, v51) != 0 {
						F_float_underflow_error(m)
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
						v56 = base.F64_sub(v41, v43)
						if base.F64_ne(base.F64_abs(v56), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
							*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v56
							if base.F64_ne(v56, float64(0)) != 0 {
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(0)
							}
							v78 = F_lseg_interpt_line(m, v13+int32(32), l1, v13+int32(8))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								if v78 == int32(0) {
									v107 = v4
									m.G0 = v13 + int32(48)
									return v107
								} else {
									v84 = F_point_dt(m, v13+int32(32), l2)
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										v88 = F_point_dt(m, v13+int32(32), v16)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											v90 = base.F64_add(v84, v88)
											v91 = F_point_dt(m, l2, v16)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												if base.F64_ne(v90, v91) != 0 {
													if base.F64_le(base.F64_abs(base.F64_sub(v90, v91)), float64(1e-06)) == int32(0) {
														v107 = v4
													} else {
														v100 = int32(1)
														if l0 == int32(0) {
															v107 = v100
														} else {
															v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
															*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
															v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
															v107 = v100
														}
													}
												} else {
													v100 = int32(1)
													if l0 == int32(0) {
														v107 = v100
													} else {
														v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
														*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
														v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
														*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
														v107 = v100
													}
												}
												m.G0 = v13 + int32(48)
												return v107
											}
										}
									}
								}
							}
						} else {
							if base.F64_eq(base.F64_abs(v41), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
								*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v56
								if base.F64_ne(v56, float64(0)) != 0 {
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(0)
								}
								v78 = F_lseg_interpt_line(m, v13+int32(32), l1, v13+int32(8))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									if v78 == int32(0) {
										v107 = v4
										m.G0 = v13 + int32(48)
										return v107
									} else {
										v84 = F_point_dt(m, v13+int32(32), l2)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											v88 = F_point_dt(m, v13+int32(32), v16)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v90 = base.F64_add(v84, v88)
												v91 = F_point_dt(m, l2, v16)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													if base.F64_ne(v90, v91) != 0 {
														if base.F64_le(base.F64_abs(base.F64_sub(v90, v91)), float64(1e-06)) == int32(0) {
															v107 = v4
														} else {
															v100 = int32(1)
															if l0 == int32(0) {
																v107 = v100
															} else {
																v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
																*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
																v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
																*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
																v107 = v100
															}
														}
													} else {
														v100 = int32(1)
														if l0 == int32(0) {
															v107 = v100
														} else {
															v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
															*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
															v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
															*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
															v107 = v100
														}
													}
													m.G0 = v13 + int32(48)
													return v107
												}
											}
										}
									}
								}
							} else {
								if base.F64_ne(v44, math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
									F_float_overflow_error(m)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v56
									if base.F64_ne(v56, float64(0)) != 0 {
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = int64(0)
									}
									v78 = F_lseg_interpt_line(m, v13+int32(32), l1, v13+int32(8))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										if v78 == int32(0) {
											v107 = v4
											m.G0 = v13 + int32(48)
											return v107
										} else {
											v84 = F_point_dt(m, v13+int32(32), l2)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												v88 = F_point_dt(m, v13+int32(32), v16)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v90 = base.F64_add(v84, v88)
													v91 = F_point_dt(m, l2, v16)
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														if base.F64_ne(v90, v91) != 0 {
															if base.F64_le(base.F64_abs(base.F64_sub(v90, v91)), float64(1e-06)) == int32(0) {
																v107 = v4
															} else {
																v100 = int32(1)
																if l0 == int32(0) {
																	v107 = v100
																} else {
																	v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
																	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
																	v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
																	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
																	v107 = v100
																}
															}
														} else {
															v100 = int32(1)
															if l0 == int32(0) {
																v107 = v100
															} else {
																v103 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
																*(*int64)(unsafe.Add(mBase, uint32(l0))) = v103
																v105 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
																*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v105
																v107 = v100
															}
														}
														m.G0 = v13 + int32(48)
														return v107
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
