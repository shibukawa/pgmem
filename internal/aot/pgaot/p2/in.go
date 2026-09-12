package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_in_range_date_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v25 int64
	_ = v25
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = int64(-9223372036854775807 - 1)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v14 == int32(-2147483648) {
		v25 = v12
		if v11 == int32(-2147483648) {
			v36 = v12
			v39 = F_Int64GetDatum(m, v25)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = F_Int64GetDatum(m, v36)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = int32(0)
					v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						return v49
					}
				}
			}
		} else {
			if v11 == int32(2147483647) {
				v36 = int64(9223372036854775807)
				v39 = F_Int64GetDatum(m, v25)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = F_Int64GetDatum(m, v36)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = int32(0)
						v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							return v49
						}
					}
				}
			} else {
				if int32(106751983) <= v11 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(248335), int32(0))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(521473), int32(658), int32(32475))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
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
					v36 = base.I64_extend_i32_s(v11) * int64(86400000000)
					v39 = F_Int64GetDatum(m, v25)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = F_Int64GetDatum(m, v36)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = int32(0)
							v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								return v49
							}
						}
					}
				}
			}
		}
	} else {
		if v14 == int32(2147483647) {
			v25 = int64(9223372036854775807)
			if v11 == int32(-2147483648) {
				v36 = v12
				v39 = F_Int64GetDatum(m, v25)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = F_Int64GetDatum(m, v36)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = int32(0)
						v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							return v49
						}
					}
				}
			} else {
				if v11 == int32(2147483647) {
					v36 = int64(9223372036854775807)
					v39 = F_Int64GetDatum(m, v25)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = F_Int64GetDatum(m, v36)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = int32(0)
							v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								return v49
							}
						}
					}
				} else {
					if int32(106751983) <= v11 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(248335), int32(0))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(521473), int32(658), int32(32475))
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
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
						v36 = base.I64_extend_i32_s(v11) * int64(86400000000)
						v39 = F_Int64GetDatum(m, v25)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = F_Int64GetDatum(m, v36)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v45 = int32(0)
								v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									return v49
								}
							}
						}
					}
				}
			}
		} else {
			if int32(106751983) <= v14 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(248335), int32(0))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(521473), int32(658), int32(32475))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
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
				v25 = base.I64_extend_i32_s(v14) * int64(86400000000)
				if v11 == int32(-2147483648) {
					v36 = v12
					v39 = F_Int64GetDatum(m, v25)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						v43 = F_Int64GetDatum(m, v36)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v45 = int32(0)
							v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								return v49
							}
						}
					}
				} else {
					if v11 == int32(2147483647) {
						v36 = int64(9223372036854775807)
						v39 = F_Int64GetDatum(m, v25)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							v43 = F_Int64GetDatum(m, v36)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								v45 = int32(0)
								v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									return v49
								}
							}
						}
					} else {
						if int32(106751983) <= v11 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(248335), int32(0))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(521473), int32(658), int32(32475))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
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
							v36 = base.I64_extend_i32_s(v11) * int64(86400000000)
							v39 = F_Int64GetDatum(m, v25)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								v43 = F_Int64GetDatum(m, v36)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v45 = int32(0)
									v49 = F_DirectFunctionCall5Coll(m, int32(1282), int32(0), v39, v43, v10, base.B2i32(v9 != v45), base.B2i32(v8 != v45))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										return v49
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
