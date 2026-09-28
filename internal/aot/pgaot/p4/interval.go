package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_interval_avg_accum(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v12 != 0 {
			v60 = v12
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v62 != 0 {
				m.G0 = v7 + int32(16)
				return base.I64_extend_i32_u(v60)
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
				if v64 != int32(2147483647) {
					if v64 != int32(-2147483648) {
						v90 = v60 + int32(8)
						F_finite_interval_pl(m, v90, v63, v90)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int64(0)
						} else {
							v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
							*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
							m.G0 = v7 + int32(16)
							return base.I64_extend_i32_u(v60)
						}
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
						if v69 != int32(-2147483648) {
							v90 = v60 + int32(8)
							F_finite_interval_pl(m, v90, v63, v90)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v60)
							}
						} else {
							v72 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
							if v72 != int64(-9223372036854775807-1) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return base.I64_extend_i32_u(v60)
								}
							} else {
								v75 = *(*int64)(unsafe.Add(mBase, uint32(v60)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v75 + int64(1)
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v60)
							}
						}
					}
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
					if v79 != int32(2147483647) {
						v90 = v60 + int32(8)
						F_finite_interval_pl(m, v90, v63, v90)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int64(0)
						} else {
							v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
							*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
							m.G0 = v7 + int32(16)
							return base.I64_extend_i32_u(v60)
						}
					} else {
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
						if v82 != int64(9223372036854775807) {
							v90 = v60 + int32(8)
							F_finite_interval_pl(m, v90, v63, v90)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v60)
							}
						} else {
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v85 + int64(1)
							m.G0 = v7 + int32(16)
							return base.I64_extend_i32_u(v60)
						}
					}
				}
			}
		} else {
			v15 = v7 + int32(12)
			v16 = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 == v16 {
				v34 = int32(0)
				if v15 == v34 {
					v42 = v34
				} else {
					v37 = v34
					v38 = v16
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
				}
				v45 = v42
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				switch v20 - int32(435) {
				case 0:
					if v15 == int32(0) {
						v45 = int32(1)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
						v37 = v27
						v38 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
						v45 = v42
					}
				case 1:
					if v15 == int32(0) {
						v45 = int32(2)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+376))
						v37 = v32
						v38 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
						v45 = v42
					}
				default:
					v34 = int32(0)
					if v15 == v34 {
						v42 = v34
					} else {
						v37 = v34
						v38 = v16
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
					}
					v45 = v42
				}
			}
			if v45 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v107 = m.ExcPending
				if v107 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_interval_avg_accum_0), int32(0))
					mBase = m.M
					v111 = m.ExcPending
					if v111 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_interval_avg_accum_1), int32(3993), int32(_a_F_interval_avg_accum_2))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = int32(_a_F_interval_avg_accum_3)
				v49 = *(*int32)(unsafe.Add(mBase, _c_F_interval_avg_accum[0]))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_accum[0])) = v51
				v54 = F_palloc0(m, int32(40))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_accum[0])) = v49
					v60 = v54
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v62 != 0 {
						m.G0 = v7 + int32(16)
						return base.I64_extend_i32_u(v60)
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
						if v64 != int32(2147483647) {
							if v64 != int32(-2147483648) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return base.I64_extend_i32_u(v60)
								}
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
								if v69 != int32(-2147483648) {
									v90 = v60 + int32(8)
									F_finite_interval_pl(m, v90, v63, v90)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
										m.G0 = v7 + int32(16)
										return base.I64_extend_i32_u(v60)
									}
								} else {
									v72 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
									if v72 != int64(-9223372036854775807-1) {
										v90 = v60 + int32(8)
										F_finite_interval_pl(m, v90, v63, v90)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int64(0)
										} else {
											v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
											*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
											m.G0 = v7 + int32(16)
											return base.I64_extend_i32_u(v60)
										}
									} else {
										v75 = *(*int64)(unsafe.Add(mBase, uint32(v60)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v75 + int64(1)
										m.G0 = v7 + int32(16)
										return base.I64_extend_i32_u(v60)
									}
								}
							}
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							if v79 != int32(2147483647) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return base.I64_extend_i32_u(v60)
								}
							} else {
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
								if v82 != int64(9223372036854775807) {
									v90 = v60 + int32(8)
									F_finite_interval_pl(m, v90, v63, v90)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
										m.G0 = v7 + int32(16)
										return base.I64_extend_i32_u(v60)
									}
								} else {
									v85 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v85 + int64(1)
									m.G0 = v7 + int32(16)
									return base.I64_extend_i32_u(v60)
								}
							}
						}
					}
				}
			}
		}
	} else {
		v15 = v7 + int32(12)
		v16 = int32(0)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v17 == v16 {
			v34 = int32(0)
			if v15 == v34 {
				v42 = v34
			} else {
				v37 = v34
				v38 = v16
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
				v42 = v38
			}
			v45 = v42
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			switch v20 - int32(435) {
			case 0:
				if v15 == int32(0) {
					v45 = int32(1)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v37 = v27
					v38 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
					v45 = v42
				}
			case 1:
				if v15 == int32(0) {
					v45 = int32(2)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+376))
					v37 = v32
					v38 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
					v45 = v42
				}
			default:
				v34 = int32(0)
				if v15 == v34 {
					v42 = v34
				} else {
					v37 = v34
					v38 = v16
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
				}
				v45 = v42
			}
		}
		if v45 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v107 = m.ExcPending
			if v107 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_interval_avg_accum_0), int32(0))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_interval_avg_accum_1), int32(3993), int32(_a_F_interval_avg_accum_2))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v48 = int32(_a_F_interval_avg_accum_3)
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_interval_avg_accum[0]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_accum[0])) = v51
			v54 = F_palloc0(m, int32(40))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_accum[0])) = v49
				v60 = v54
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
				if v62 != 0 {
					m.G0 = v7 + int32(16)
					return base.I64_extend_i32_u(v60)
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
					if v64 != int32(2147483647) {
						if v64 != int32(-2147483648) {
							v90 = v60 + int32(8)
							F_finite_interval_pl(m, v90, v63, v90)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v60)
							}
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							if v69 != int32(-2147483648) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return base.I64_extend_i32_u(v60)
								}
							} else {
								v72 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
								if v72 != int64(-9223372036854775807-1) {
									v90 = v60 + int32(8)
									F_finite_interval_pl(m, v90, v63, v90)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
										m.G0 = v7 + int32(16)
										return base.I64_extend_i32_u(v60)
									}
								} else {
									v75 = *(*int64)(unsafe.Add(mBase, uint32(v60)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v75 + int64(1)
									m.G0 = v7 + int32(16)
									return base.I64_extend_i32_u(v60)
								}
							}
						}
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
						if v79 != int32(2147483647) {
							v90 = v60 + int32(8)
							F_finite_interval_pl(m, v90, v63, v90)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v60)
							}
						} else {
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
							if v82 != int64(9223372036854775807) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return base.I64_extend_i32_u(v60)
								}
							} else {
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v85 + int64(1)
								m.G0 = v7 + int32(16)
								return base.I64_extend_i32_u(v60)
							}
						}
					}
				}
			}
		}
	}
}
func F_interval_avg_accum_inv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v6 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v71 = m.ExcPending
		if v71 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_interval_avg_accum_inv_0), int32(0))
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_interval_avg_accum_inv_1), int32(_a_F_interval_avg_accum_inv_2), int32(_a_F_interval_avg_accum_inv_3))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v8 = base.I32_wrap_i64(v7)
		if v8 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_interval_avg_accum_inv_0), int32(0))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_interval_avg_accum_inv_1), int32(_a_F_interval_avg_accum_inv_2), int32(_a_F_interval_avg_accum_inv_3))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v12 = v7 & int64(4294967295)
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v13 == int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				if v17 != int32(2147483647) {
					if v17 != int32(-2147483648) {
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
						v46 = v44 - int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = v46
						v49 = v8 + int32(8)
						if int64(0) < v46 {
							F_finite_interval_mi(m, v49, v16, v49)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								return v12
							}
						} else {
							v57 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v57
							*(*int64)(unsafe.Add(mBase, uint32(v49))) = v57
							return v12
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						if v22 != int32(-2147483648) {
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
							v46 = v44 - int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v46
							v49 = v8 + int32(8)
							if int64(0) < v46 {
								F_finite_interval_mi(m, v49, v16, v49)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									return v12
								}
							} else {
								v57 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v57
								*(*int64)(unsafe.Add(mBase, uint32(v49))) = v57
								return v12
							}
						} else {
							v25 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							if v25 != int64(-9223372036854775807-1) {
								v44 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
								v46 = v44 - int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v46
								v49 = v8 + int32(8)
								if int64(0) < v46 {
									F_finite_interval_mi(m, v49, v16, v49)
									mBase = m.M
									v55 = m.ExcPending
									if v55 != 0 {
										return int64(0)
									} else {
										return v12
									}
								} else {
									v57 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v57
									*(*int64)(unsafe.Add(mBase, uint32(v49))) = v57
									return v12
								}
							} else {
								v28 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v28 - int64(1)
								return v12
							}
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
					if v33 != int32(2147483647) {
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
						v46 = v44 - int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v8))) = v46
						v49 = v8 + int32(8)
						if int64(0) < v46 {
							F_finite_interval_mi(m, v49, v16, v49)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int64(0)
							} else {
								return v12
							}
						} else {
							v57 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v57
							*(*int64)(unsafe.Add(mBase, uint32(v49))) = v57
							return v12
						}
					} else {
						v36 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						if v36 != int64(9223372036854775807) {
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
							v46 = v44 - int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v46
							v49 = v8 + int32(8)
							if int64(0) < v46 {
								F_finite_interval_mi(m, v49, v16, v49)
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int64(0)
								} else {
									return v12
								}
							} else {
								v57 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v57
								*(*int64)(unsafe.Add(mBase, uint32(v49))) = v57
								return v12
							}
						} else {
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v39 - int64(1)
							return v12
						}
					}
				}
			} else {
				return v12
			}
		}
	}
}
func F_interval_avg_deserialize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v11 == int32(0) {
		v39 = int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
		switch v14 - int32(435) {
		case 0:
			v39 = int32(1)
		case 1:
			v39 = int32(2)
		default:
			v39 = int32(0)
		}
	}
	if v39 != 0 {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v41 = F_pg_detoast_datum_packed(m, v40)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int64(0)
		} else {
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
			if v45 == int32(1) {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
				if v51 == int32(18) {
					v54 = int32(16)
				} else {
					v54 = int32(0)
				}
				if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v61 = int32(4)
				} else {
					v61 = v54
				}
				v74 = v61
			} else {
				v62 = int32(1)
				if v45&v62 != 0 {
					v74 = int32(base.Ui32(v45)>>(uint(v62)%32)) - v62
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v74
			v78 = int32(1)
			if v45&v78 != 0 {
				v82 = v78
			} else {
				v82 = int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v41 + v82
			v86 = F_palloc0(m, int32(40))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int64(0)
			} else {
				v88 = F_pq_getmsgint64(m, v7)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v86))) = v88
					v91 = F_pq_getmsgint64(m, v7)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int64(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = v91
						v95 = F_pq_getmsgint(m, v7, int32(4))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v95
							v99 = F_pq_getmsgint(m, v7, int32(4))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = v99
								v102 = F_pq_getmsgint64(m, v7)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = v102
									v105 = F_pq_getmsgint64(m, v7)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v86)+32)) = v105
										F_pq_getmsgend(m, v7)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int64(0)
										} else {
											m.G0 = v7 + int32(16)
											return base.I64_extend_i32_u(v86)
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v118 = m.ExcPending
		if v118 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_interval_avg_deserialize_0), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_interval_avg_deserialize_1), int32(_a_F_interval_avg_deserialize_2), int32(_a_F_interval_avg_deserialize_3))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_interval_cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v44 int64
	_ = v44
	var v51 int64
	_ = v51
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int64
	_ = v65
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int64
	_ = v84
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v99 int64
	_ = v99
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v115 int64
	_ = v115
	var v119 int64
	_ = v119
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+12)))
	v20 = int64(*(*int32)(unsafe.Add(mBase, uint32(v16)+8)))
	v21 = v17*int64(30) + v20
	v30 = int64(32)
	v31 = int64(20)
	v33 = int64(base.Ui64(v21) >> (uint(v30) % 64))
	v36 = int64(4294967295)
	v37 = int64(500654080)
	v39 = v21 & v36
	v40 = v37 * v39
	v44 = int64(base.Ui64(v40)>>(uint(v30)%64)) + v37*v33
	v51 = v39*v31 + v44&v36
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v21*int64(0) + v21>>(uint(int64(63))%64)*int64(86400000000) + v31*v33 + int64(base.Ui64(v44)>>(uint(v30)%64)) + int64(base.Ui64(v51)>>(uint(v30)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v40&v36 | v51<<(uint(v30)%64)
	v63 = v14 + int32(16)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+12)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	v69 = v65*int64(30) + v68
	v78 = int64(32)
	v79 = int64(20)
	v81 = int64(base.Ui64(v69) >> (uint(v78) % 64))
	v84 = int64(4294967295)
	v85 = int64(500654080)
	v87 = v69 & v84
	v88 = v85 * v87
	v92 = int64(base.Ui64(v88)>>(uint(v78)%64)) + v85*v81
	v99 = v87*v79 + v92&v84
	*(*int64)(unsafe.Add(mBase, uint32(v63)+8)) = v69*int64(0) + v69>>(uint(int64(63))%64)*int64(86400000000) + v79*v81 + int64(base.Ui64(v92)>>(uint(v78)%64)) + int64(base.Ui64(v99)>>(uint(v78)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v63))) = v88&v84 | v99<<(uint(v78)%64)
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	m.G0 = v14 + int32(32)
	v119 = v113 + v115
	v120 = v110 + v112
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	v135 = base.B2i32(v133 == v127)
	if v133 == v127 {
		v136 = base.B2i32(base.Ui64(v120) < base.Ui64(v119))
	} else {
		v136 = base.B2i32(v133 < v127)
	}
	if v133 == v127 {
		v139 = base.B2i32(base.Ui64(v119) < base.Ui64(v120))
	} else {
		v139 = base.B2i32(v127 < v133)
	}
	return base.I64_extend_i32_s(v136 - v139)
}
func F_interval_cmp_lower_1(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 float64
	_ = v3
	var v4 float64
	_ = v4
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	v3 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v3) & v9
	v13 = base.I64_reinterpret_f64(v4) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v13) {
		v23 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v32 = int32(0) - v23&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
	} else {
		v18 = int32(1)
		if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v10))|base.F64_gt(v3, v4) != 0 {
			v32 = v18
		} else {
			v23 = v18
			v32 = int32(0) - v23&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
		}
	}
	return v32
}
func F_interval_cmp_upper_2(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = int32(16)
	v8 = F_range_cmp_bounds(m, l2, l0+v4, l1+v4)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_interval_justify_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v22 int32
	_ = v22
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_palloc(m, int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v18
	if v14 != int32(2147483647) {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L51
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	return base.I64_extend_i32_u(v10)
L6:
	;
	v55 = base.I64_div_s(v18, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v18+int64(86399999999)) {
		goto L23
	} else {
		goto L24
	}
L7:
	;
	v40 = base.I32_div_s(v16, int32(30))
	v41 = v14 + v40
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v41
	v45 = v40*int32(-30) + v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v45
	if base.B2i32(v40 < int32(0)) != base.B2i32(v41 < v14) {
		goto L4
	} else {
		goto L22
	}
L8:
	;
	if v18 < int64(0) {
		goto L7
	} else {
		goto L21
	}
L9:
	;
	if int64(0) < v18 {
		goto L7
	} else {
		goto L20
	}
L10:
	;
	if int32(0) < v16 {
		goto L9
	} else {
		goto L18
	}
L11:
	;
	v22 = int32(-2147483648)
	if base.B2i32(v14 != v22)|base.B2i32(v16 != v22) != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v16 != int32(2147483647) {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	if v18 == int64(-9223372036854775807-1) {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	goto L10
L16:
	;
	if v18 == int64(9223372036854775807) {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	goto L9
L18:
	;
	if v16 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v51 = v14
	v52 = v16
	goto L6
L20:
	;
	v51 = v14
	v52 = v16
	goto L6
L21:
	;
	v51 = v14
	v52 = v16
	goto L6
L22:
	;
	v51 = v41
	v52 = v45
	goto L6
L23:
	;
	v62 = v55*int64(-86400000000) + v18
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v62
	v64 = v62
	goto L25
L24:
	;
	v64 = v18
	goto L25
L25:
	;
	v66 = v52 + base.I32_wrap_i64(v55)
	v68 = base.I32_div_s(v66, int32(30))
	v69 = v51 + v68
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v69
	v73 = v68*int32(-30) + v66
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v73
	if base.B2i32(v68 < int32(0)) != base.B2i32(v69 < v51) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	if int32(0) < v69 {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v131
	goto L5
L28:
	;
	if v64 <= int64(0) {
		goto L5
	} else {
		goto L46
	}
L29:
	;
	if int32(0) <= v102 {
		goto L5
	} else {
		goto L45
	}
L30:
	;
	if int64(0) <= v64 {
		goto L5
	} else {
		goto L44
	}
L31:
	;
	if v102 <= int32(0) {
		goto L29
	} else {
		goto L43
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v69 + v96
	v99 = v95 + v73
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v99
	v102 = v99
	goto L31
L33:
	;
	v81 = int32(-1)
	v82 = int32(30)
	if v73 < int32(0) {
		v95 = v82
		v96 = v81
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	if int32(0) <= v69 {
		v102 = v73
		goto L31
	} else {
		goto L39
	}
L36:
	;
	if v73 != 0 {
		v107 = v73
		goto L30
	} else {
		goto L37
	}
L37:
	;
	if v64 < int64(0) {
		v95 = v82
		v96 = v81
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L5
L39:
	;
	v89 = int32(1)
	v90 = int32(-30)
	if int32(0) < v73 {
		v95 = v90
		v96 = v89
		goto L32
	} else {
		goto L40
	}
L40:
	;
	if v73 != 0 {
		v119 = v73
		goto L28
	} else {
		goto L41
	}
L41:
	;
	if v64 <= int64(0) {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v95 = v90
	v96 = v89
	goto L32
L43:
	;
	v107 = v102
	goto L30
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v64 + int64(86400000000)
	v131 = v107 - int32(1)
	goto L27
L45:
	;
	v119 = v102
	goto L28
L46:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v64 - int64(86400000000)
	v131 = v119 + int32(1)
	goto L27
L47:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_interval_justify_interval_0), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_interval_justify_interval_1), int32(2959), int32(_a_F_interval_justify_interval_2))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_interval_justify_interval_0), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_interval_justify_interval_1), int32(2976), int32(_a_F_interval_justify_interval_2))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_interval_mul(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v97 int64
	_ = v97
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v127 int32
	_ = v127
	var v144 float64
	_ = v144
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 float64
	_ = v164
	var v167 int32
	_ = v167
	var v186 float64
	_ = v186
	var v190 float64
	_ = v190
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 float64
	_ = v200
	var v206 float64
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v225 int32
	_ = v225
	var v227 float64
	_ = v227
	var v228 int32
	_ = v228
	var v234 int64
	_ = v234
	var v240 float64
	_ = v240
	var v243 int32
	_ = v243
	var v256 int64
	_ = v256
	var v260 int32
	_ = v260
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	v15 = m.G0
	v16 = int32(16)
	v17 = v15 - v16
	m.G0 = v17
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v24 = F_palloc(m, v16)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(v19&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v17 + int32(16)
	return base.I64_extend_i32_u(v24)
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	v32 = base.F64_reinterpret_i64(v19)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v33 != int32(2147483647) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if base.F64_eq(base.F64_abs(v32), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	if base.F64_eq(v32, float64(0)) != 0 {
		goto L4
	} else {
		goto L16
	}
L8:
	;
	if v33 != int32(-2147483648) {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v44 != int32(2147483647) {
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v38 != int32(-2147483648) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	if v41 == int64(-9223372036854775807-1) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L6
L14:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	if v47 != int64(9223372036854775807) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	if base.F64_lt(v32, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_interval_um_internal(m, v20, v24)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v56
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v58
	goto L3
L20:
	;
	goto L3
L21:
	;
	v63 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+8)))
	v67 = v63 + base.I64_extend_i32_s(v33)*int64(30)
	v76 = int64(32)
	v77 = int64(20)
	v79 = int64(base.Ui64(v67) >> (uint(v76) % 64))
	v82 = int64(4294967295)
	v83 = int64(500654080)
	v85 = v67 & v82
	v86 = v83 * v85
	v90 = int64(base.Ui64(v86)>>(uint(v76)%64)) + v83*v79
	v97 = v85*v77 + v90&v82
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v67*int64(0) + v67>>(uint(int64(63))%64)*int64(86400000000) + v77*v79 + int64(base.Ui64(v90)>>(uint(v76)%64)) + int64(base.Ui64(v97)>>(uint(v76)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v86&v82 | v97<<(uint(v76)%64)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v144 = base.F64_mul(v32, base.F64_convert_i32_s(v33))
	v147 = int32(0)
	if base.B2i32(base.F64_lt(v144, float64(2.147483648e+09)) == v147)|base.B2i32(base.F64_ge(v144, float64(-2.147483648e+09)) == v147)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v144)&int64(9223372036854775807))) != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v110 = v108 + v109
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v117 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v110) < base.Ui64(v108))) + (v113 + v109>>(uint(int64(63))%64))
	if v110|v117 == int64(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v121 = int64(0)
	if v117 == v121 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v127 = base.B2i32(v110 != v121)
	goto L28
L27:
	;
	v127 = base.B2i32(v121 < v117)
	goto L28
L28:
	;
	if base.F64_lt(base.F64_mul(v32, base.F64_convert_i32_s(v127-base.B2i32(v117 < int64(0)))), float64(0)) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(-9223372034707292160)
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = int64(-9223372036854775807 - 1)
	goto L3
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = int64(9223372034707292159)
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = int64(9223372036854775807)
	goto L3
L32:
	;
	v160 = base.I32_trunc_sat_f64_s(v144)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v164 = base.F64_mul(v32, base.F64_convert_i32_s(v162))
	v167 = int32(0)
	if base.B2i32(base.F64_lt(v164, float64(2.147483648e+09)) == v167)|base.B2i32(base.F64_ge(v164, float64(-2.147483648e+09)) == v167)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v164)&int64(9223372036854775807))) != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v186 = float64(1e+06)
	v190 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_mul(base.F64_convert_i32_s(v22), v32), base.F64_convert_i32_s(v160)), float64(30)), v186)), v186)
	v193 = base.I32_trunc_sat_f64_s(v164)
	v197 = base.I32_trunc_sat_f64_s(v190)
	v200 = float64(86400)
	v206 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_add(v190, base.F64_sub(base.F64_mul(base.F64_convert_i32_s(v21), v32), base.F64_convert_i32_s(v193))), base.F64_convert_i32_s(v197)), v200), v186)), v186)
	if base.F64_ge(base.F64_abs(v206), v200) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v228 = v225 + v197
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v228
	if base.B2i32(v197 < int32(0))^base.B2i32(v228 < v225) != 0 {
		goto L4
	} else {
		goto L39
	}
L35:
	;
	v225 = v193
	v227 = v206
	goto L34
L36:
	;
	goto L37
L37:
	;
	v214 = base.I32_trunc_sat_f64_s(base.F64_div(v206, float64(86400)))
	v215 = v193 + v214
	*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v215
	if base.B2i32(v214 < int32(0))^base.B2i32(v215 < v193) != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v225 = v215
	v227 = base.F64_sub(v206, base.F64_convert_i32_s(v214*int32(_a_F_interval_mul_3)))
	goto L34
L39:
	;
	v234 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v240 = base.F64_nearest(base.F64_add(base.F64_mul(base.F64_convert_i64_s(v234), v32), base.F64_mul(v227, float64(1e+06))))
	v243 = int32(0)
	if base.B2i32(base.F64_lt(v240, float64(9.223372036854776e+18)) == v243)|base.B2i32(base.F64_ge(v240, float64(-9.223372036854776e+18)) == v243)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v240)&int64(9223372036854775807))) != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v256 = base.I64_trunc_sat_f64_s(v240)
	*(*int64)(unsafe.Add(mBase, uint32(v24))) = v256
	if v160 != int32(2147483647) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v260 = int32(-2147483648)
	if base.B2i32(v160 != v260)|base.B2i32(v228 != v260) != 0 {
		goto L3
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if base.B2i32(v228 != int32(2147483647))|base.B2i32(v256 != int64(9223372036854775807)) != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	if v256 == int64(-9223372036854775807-1) {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	goto L3
L46:
	;
	goto L4
L47:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_interval_mul_0), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_interval_mul_1), int32(3741), int32(_a_F_interval_mul_2))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_interval_recv(m *base.Module, l0 int32) int64 {
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
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = F_pq_getmsgint64(m, v5)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v11
			v15 = F_pq_getmsgint(m, v5, int32(4))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
				v19 = F_pq_getmsgint(m, v5, int32(4))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v19
					v23 = F_AdjustIntervalForTypmod(m, v7, v4, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v7)
					}
				}
			}
		}
	}
}
func F_interval_to_char(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v120 int64
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v152 int64
	_ = v152
	v8 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
		if v21 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
				if v52 != int32(2147483647) {
					if v52 != int32(-2147483648) {
						v72 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
						*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
						v83 = v13 + int32(24)
						v85 = v13 + int32(8)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
						v87 = int32(12)
						v88 = base.I32_div_s(v86, v87)
						*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						v98 = base.I64_div_s(v96, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
						v102 = v96 + v98*int64(-3600000000)
						v104 = base.I64_div_s(v102, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
						v108 = v104*int64(-60000000) + v102
						v110 = base.I64_div_s(v108, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
						v114 = v110*int64(4293967296) + v108
						*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int64(0)
						} else {
							if v139 == int32(0) {
								v143 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
								v152 = v8
							} else {
								v152 = base.I64_extend_i32_u(v139)
							}
							m.G0 = v13 + int32(112)
							return v152
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						if v57 != int32(-2147483648) {
							v72 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
							v83 = v13 + int32(24)
							v85 = v13 + int32(8)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
							v87 = int32(12)
							v88 = base.I32_div_s(v86, v87)
							*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
							*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v98 = base.I64_div_s(v96, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
							v102 = v96 + v98*int64(-3600000000)
							v104 = base.I64_div_s(v102, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
							v108 = v104*int64(-60000000) + v102
							v110 = base.I64_div_s(v108, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
							v114 = v110*int64(4293967296) + v108
							*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								if v139 == int32(0) {
									v143 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
									v152 = v8
								} else {
									v152 = base.I64_extend_i32_u(v139)
								}
								m.G0 = v13 + int32(112)
								return v152
							}
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							if v60 == int64(-9223372036854775807-1) {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v152 = v8
								m.G0 = v13 + int32(112)
								return v152
							} else {
								v72 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
								v83 = v13 + int32(24)
								v85 = v13 + int32(8)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
								v87 = int32(12)
								v88 = base.I32_div_s(v86, v87)
								*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
								*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								v98 = base.I64_div_s(v96, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
								v102 = v96 + v98*int64(-3600000000)
								v104 = base.I64_div_s(v102, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
								v108 = v104*int64(-60000000) + v102
								v110 = base.I64_div_s(v108, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
								v114 = v110*int64(4293967296) + v108
								*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
								*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									if v139 == int32(0) {
										v143 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
										v152 = v8
									} else {
										v152 = base.I64_extend_i32_u(v139)
									}
									m.G0 = v13 + int32(112)
									return v152
								}
							}
						}
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					if v63 != int32(2147483647) {
						v72 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
						*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
						v83 = v13 + int32(24)
						v85 = v13 + int32(8)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
						v87 = int32(12)
						v88 = base.I32_div_s(v86, v87)
						*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						v98 = base.I64_div_s(v96, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
						v102 = v96 + v98*int64(-3600000000)
						v104 = base.I64_div_s(v102, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
						v108 = v104*int64(-60000000) + v102
						v110 = base.I64_div_s(v108, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
						v114 = v110*int64(4293967296) + v108
						*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int64(0)
						} else {
							if v139 == int32(0) {
								v143 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
								v152 = v8
							} else {
								v152 = base.I64_extend_i32_u(v139)
							}
							m.G0 = v13 + int32(112)
							return v152
						}
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						if v66 != int64(9223372036854775807) {
							v72 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
							v83 = v13 + int32(24)
							v85 = v13 + int32(8)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
							v87 = int32(12)
							v88 = base.I32_div_s(v86, v87)
							*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
							*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v98 = base.I64_div_s(v96, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
							v102 = v96 + v98*int64(-3600000000)
							v104 = base.I64_div_s(v102, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
							v108 = v104*int64(-60000000) + v102
							v110 = base.I64_div_s(v108, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
							v114 = v110*int64(4293967296) + v108
							*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								if v139 == int32(0) {
									v143 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
									v152 = v8
								} else {
									v152 = base.I64_extend_i32_u(v139)
								}
								m.G0 = v13 + int32(112)
								return v152
							}
						} else {
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
							v152 = v8
							m.G0 = v13 + int32(112)
							return v152
						}
					}
				}
			} else {
				if v24 == int32(18) {
					v35 = int32(16)
				} else {
					v35 = int32(0)
				}
				v48 = v35
				if v48 == int32(0) {
					v70 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
					v152 = v8
					m.G0 = v13 + int32(112)
					return v152
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					if v52 != int32(2147483647) {
						if v52 != int32(-2147483648) {
							v72 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
							v83 = v13 + int32(24)
							v85 = v13 + int32(8)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
							v87 = int32(12)
							v88 = base.I32_div_s(v86, v87)
							*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
							*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v98 = base.I64_div_s(v96, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
							v102 = v96 + v98*int64(-3600000000)
							v104 = base.I64_div_s(v102, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
							v108 = v104*int64(-60000000) + v102
							v110 = base.I64_div_s(v108, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
							v114 = v110*int64(4293967296) + v108
							*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								if v139 == int32(0) {
									v143 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
									v152 = v8
								} else {
									v152 = base.I64_extend_i32_u(v139)
								}
								m.G0 = v13 + int32(112)
								return v152
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							if v57 != int32(-2147483648) {
								v72 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
								v83 = v13 + int32(24)
								v85 = v13 + int32(8)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
								v87 = int32(12)
								v88 = base.I32_div_s(v86, v87)
								*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
								*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								v98 = base.I64_div_s(v96, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
								v102 = v96 + v98*int64(-3600000000)
								v104 = base.I64_div_s(v102, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
								v108 = v104*int64(-60000000) + v102
								v110 = base.I64_div_s(v108, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
								v114 = v110*int64(4293967296) + v108
								*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
								*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									if v139 == int32(0) {
										v143 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
										v152 = v8
									} else {
										v152 = base.I64_extend_i32_u(v139)
									}
									m.G0 = v13 + int32(112)
									return v152
								}
							} else {
								v60 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								if v60 == int64(-9223372036854775807-1) {
									v70 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
									v152 = v8
									m.G0 = v13 + int32(112)
									return v152
								} else {
									v72 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
									*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
									*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
									v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
									v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
									v83 = v13 + int32(24)
									v85 = v13 + int32(8)
									v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
									v87 = int32(12)
									v88 = base.I32_div_s(v86, v87)
									*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
									*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
									v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
									v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
									v98 = base.I64_div_s(v96, int64(3600000000))
									*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
									v102 = v96 + v98*int64(-3600000000)
									v104 = base.I64_div_s(v102, int64(60000000))
									*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
									v108 = v104*int64(-60000000) + v102
									v110 = base.I64_div_s(v108, int64(1000000))
									*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
									v114 = v110*int64(4293967296) + v108
									*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
									v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
									v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
									*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int64(0)
									} else {
										if v139 == int32(0) {
											v143 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
											v152 = v8
										} else {
											v152 = base.I64_extend_i32_u(v139)
										}
										m.G0 = v13 + int32(112)
										return v152
									}
								}
							}
						}
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						if v63 != int32(2147483647) {
							v72 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
							v83 = v13 + int32(24)
							v85 = v13 + int32(8)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
							v87 = int32(12)
							v88 = base.I32_div_s(v86, v87)
							*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
							*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v98 = base.I64_div_s(v96, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
							v102 = v96 + v98*int64(-3600000000)
							v104 = base.I64_div_s(v102, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
							v108 = v104*int64(-60000000) + v102
							v110 = base.I64_div_s(v108, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
							v114 = v110*int64(4293967296) + v108
							*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								if v139 == int32(0) {
									v143 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
									v152 = v8
								} else {
									v152 = base.I64_extend_i32_u(v139)
								}
								m.G0 = v13 + int32(112)
								return v152
							}
						} else {
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							if v66 != int64(9223372036854775807) {
								v72 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
								v83 = v13 + int32(24)
								v85 = v13 + int32(8)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
								v87 = int32(12)
								v88 = base.I32_div_s(v86, v87)
								*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
								*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								v98 = base.I64_div_s(v96, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
								v102 = v96 + v98*int64(-3600000000)
								v104 = base.I64_div_s(v102, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
								v108 = v104*int64(-60000000) + v102
								v110 = base.I64_div_s(v108, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
								v114 = v110*int64(4293967296) + v108
								*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
								*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									if v139 == int32(0) {
										v143 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
										v152 = v8
									} else {
										v152 = base.I64_extend_i32_u(v139)
									}
									m.G0 = v13 + int32(112)
									return v152
								}
							} else {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v152 = v8
								m.G0 = v13 + int32(112)
								return v152
							}
						}
					}
				}
			}
		} else {
			v36 = int32(1)
			if v21&v36 != 0 {
				v48 = int32(base.Ui32(v21)>>(uint(v36)%32)) - v36
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
			if v48 == int32(0) {
				v70 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
				v152 = v8
				m.G0 = v13 + int32(112)
				return v152
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
				if v52 != int32(2147483647) {
					if v52 != int32(-2147483648) {
						v72 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
						*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
						v83 = v13 + int32(24)
						v85 = v13 + int32(8)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
						v87 = int32(12)
						v88 = base.I32_div_s(v86, v87)
						*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						v98 = base.I64_div_s(v96, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
						v102 = v96 + v98*int64(-3600000000)
						v104 = base.I64_div_s(v102, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
						v108 = v104*int64(-60000000) + v102
						v110 = base.I64_div_s(v108, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
						v114 = v110*int64(4293967296) + v108
						*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int64(0)
						} else {
							if v139 == int32(0) {
								v143 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
								v152 = v8
							} else {
								v152 = base.I64_extend_i32_u(v139)
							}
							m.G0 = v13 + int32(112)
							return v152
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						if v57 != int32(-2147483648) {
							v72 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
							v83 = v13 + int32(24)
							v85 = v13 + int32(8)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
							v87 = int32(12)
							v88 = base.I32_div_s(v86, v87)
							*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
							*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v98 = base.I64_div_s(v96, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
							v102 = v96 + v98*int64(-3600000000)
							v104 = base.I64_div_s(v102, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
							v108 = v104*int64(-60000000) + v102
							v110 = base.I64_div_s(v108, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
							v114 = v110*int64(4293967296) + v108
							*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								if v139 == int32(0) {
									v143 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
									v152 = v8
								} else {
									v152 = base.I64_extend_i32_u(v139)
								}
								m.G0 = v13 + int32(112)
								return v152
							}
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							if v60 == int64(-9223372036854775807-1) {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v152 = v8
								m.G0 = v13 + int32(112)
								return v152
							} else {
								v72 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
								*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
								*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
								v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
								v83 = v13 + int32(24)
								v85 = v13 + int32(8)
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
								v87 = int32(12)
								v88 = base.I32_div_s(v86, v87)
								*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
								*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								v98 = base.I64_div_s(v96, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
								v102 = v96 + v98*int64(-3600000000)
								v104 = base.I64_div_s(v102, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
								v108 = v104*int64(-60000000) + v102
								v110 = base.I64_div_s(v108, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
								v114 = v110*int64(4293967296) + v108
								*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
								v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
								*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int64(0)
								} else {
									if v139 == int32(0) {
										v143 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
										v152 = v8
									} else {
										v152 = base.I64_extend_i32_u(v139)
									}
									m.G0 = v13 + int32(112)
									return v152
								}
							}
						}
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					if v63 != int32(2147483647) {
						v72 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
						*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
						*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
						v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
						v83 = v13 + int32(24)
						v85 = v13 + int32(8)
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
						v87 = int32(12)
						v88 = base.I32_div_s(v86, v87)
						*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
						*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
						v98 = base.I64_div_s(v96, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
						v102 = v96 + v98*int64(-3600000000)
						v104 = base.I64_div_s(v102, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
						v108 = v104*int64(-60000000) + v102
						v110 = base.I64_div_s(v108, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
						v114 = v110*int64(4293967296) + v108
						*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int64(0)
						} else {
							if v139 == int32(0) {
								v143 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
								v152 = v8
							} else {
								v152 = base.I64_extend_i32_u(v139)
							}
							m.G0 = v13 + int32(112)
							return v152
						}
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						if v66 != int64(9223372036854775807) {
							v72 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v72
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v72
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = int32(0)
							v78 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v78
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v80
							v83 = v13 + int32(24)
							v85 = v13 + int32(8)
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
							v87 = int32(12)
							v88 = base.I32_div_s(v86, v87)
							*(*int32)(unsafe.Add(mBase, uint32(v83)+32)) = v88
							*(*int32)(unsafe.Add(mBase, uint32(v83)+28)) = v86 - v88*v87
							v94 = *(*int32)(unsafe.Add(mBase, uint32(v85)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = v94
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							v98 = base.I64_div_s(v96, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v83)+16)) = v98
							v102 = v96 + v98*int64(-3600000000)
							v104 = base.I64_div_s(v102, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+8)) = uint32(v104)
							v108 = v104*int64(-60000000) + v102
							v110 = base.I64_div_s(v108, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v83)+4)) = uint32(v110)
							v114 = v110*int64(4293967296) + v108
							*(*uint32)(unsafe.Add(mBase, uint32(v83))) = uint32(v114)
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v116
							v118 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v120
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v126
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v122 + (v124+v126*v87)*int32(30)
							v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v139 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v138)
							mBase = m.M
							v140 = m.ExcPending
							if v140 != 0 {
								return int64(0)
							} else {
								if v139 == int32(0) {
									v143 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
									v152 = v8
								} else {
									v152 = base.I64_extend_i32_u(v139)
								}
								m.G0 = v13 + int32(112)
								return v152
							}
						} else {
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
							v152 = v8
							m.G0 = v13 + int32(112)
							return v152
						}
					}
				}
			}
		}
	}
}
func F_interval_trunc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v78 int32
	_ = v78
	var v84 int64
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v213 int64
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v263 int64
	_ = v263
	var v269 int32
	_ = v269
	var v272 int64
	_ = v272
	var v275 int64
	_ = v275
	var v276 int64
	_ = v276
	var v284 int64
	_ = v284
	var v285 int64
	_ = v285
	var v291 int64
	_ = v291
	var v292 int64
	_ = v292
	var v300 int32
	_ = v300
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = F_palloc(m, int32(16))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = int32(1)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v29 = v27 & v25
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v14 + int32(48)
	return base.I64_extend_i32_u(v23)
L5:
	;
	v163 = base.I64_div_s(v160, int64(3600000000))
	v166 = v163*int64(-3600000000) + v160
	v168 = base.I64_div_s(v166, int64(60000000))
	v171 = v168*int64(-60000000) + v166
	v173 = base.I64_div_s(v171, int64(1000000))
	v177 = base.I32_wrap_i64(v173*int64(4293967296) + v171)
	v178 = int32(12)
	v179 = base.I32_div_s(v70, v178)
	v182 = v70 - v179*v178
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	switch v183 - int32(18) {
	case 0:
		goto L57
	case 1:
		v248 = v168
		v249 = v163
		goto L53
	case 2:
		goto L58
	case 3:
		goto L54
	default:
		goto L55
	case 5:
		v205 = v179
		v206 = v182
		goto L59
	case 6:
		v199 = v179
		v200 = v182
		goto L60
	case 7:
		v197 = v179
		goto L61
	case 8:
		v193 = v179
		goto L62
	case 9:
		v189 = v179
		goto L63
	case 10:
		goto L64
	case 11:
		goto L56
	case 12:
		v252 = v179
		v253 = v168
		v254 = v173
		v255 = v163
		v256 = v161
		v257 = v177
		v258 = v182
		goto L52
	}
L6:
	;
	v30 = v25
	goto L8
L7:
	;
	v30 = int32(4)
	goto L8
L8:
	;
	if v27 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v60 = F_downcase_truncate_identifier(m, v17+v30, v58, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v37 == int32(18) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v48 = int32(1)
	if v29 != 0 {
		v58 = int32(base.Ui32(v27)>>(uint(v48)%32)) - v48
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v40 = int32(16)
	goto L15
L14:
	;
	v40 = int32(0)
	goto L15
L15:
	;
	if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v47 = int32(4)
	goto L18
L17:
	;
	v47 = v40
	goto L18
L18:
	;
	v58 = v47
	goto L9
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L20:
	;
	v67 = Fn14210(m, v60, v14+int32(44), int32(_a_F_interval_trunc_0), int32(_a_F_interval_trunc_1), int32(_a_F_interval_trunc_2))
	mBase = m.M
	goto L21
L21:
	;
	if v67 == int32(17) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v70 != int32(-2147483648) {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L47
	}
L25:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if base.B2i32(base.Ui32(int32(8)) <= base.Ui32(v93-int32(23)))&base.B2i32(base.Ui32(int32(3)) < base.Ui32(v93-int32(18))) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v84 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v85 != int32(2147483647) {
		v160 = v84
		v161 = v85
		goto L5
	} else {
		goto L33
	}
L27:
	;
	if v70 == int32(2147483647) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v78 != int32(-2147483648) {
		v160 = v77
		v161 = v78
		goto L5
	} else {
		goto L31
	}
L30:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v160 = v76
	v161 = v75
	goto L5
L31:
	;
	if v77 == int64(-9223372036854775807-1) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v160 = v77
	v161 = int32(-2147483648)
	goto L5
L33:
	;
	if v84 == int64(9223372036854775807) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v160 = v84
	v161 = int32(2147483647)
	goto L5
L35:
	;
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v105
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v107
	goto L4
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v117 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v60
	F_errmsg(m, int32(_a_F_interval_trunc_7), v14+int32(16))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v126 == int32(22) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v131 = F_errdetail(m, int32(_a_F_interval_trunc_8), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errfinish(m, int32(_a_F_interval_trunc_4), int32(_a_F_interval_trunc_10), int32(_a_F_interval_trunc_6))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L46
	}
L45:
	;
	goto L44
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v146 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v60
	F_errmsg(m, int32(_a_F_interval_trunc_11), v14+int32(32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_interval_trunc_4), int32(_a_F_interval_trunc_12), int32(_a_F_interval_trunc_6))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	v263 = base.I64_extend_i32_s(v258) + base.I64_extend_i32_s(v252)*int64(12)
	if base.Ui64(v263-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L74
	} else {
		goto L75
	}
L53:
	;
	v252 = v179
	v253 = v248
	v254 = int64(0)
	v255 = v249
	v256 = v161
	v257 = int32(0)
	v258 = v182
	goto L52
L54:
	;
	v246 = int64(0)
	v248 = v246
	v249 = v246
	goto L53
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L65
	}
L56:
	;
	v217 = base.I32_rem_s(v177, int32(1000))
	v252 = v179
	v253 = v168
	v254 = v173
	v255 = v163
	v256 = v161
	v257 = v177 - v217
	v258 = v182
	goto L52
L57:
	;
	v252 = v179
	v253 = v168
	v254 = v173
	v255 = v163
	v256 = v161
	v257 = int32(0)
	v258 = v182
	goto L52
L58:
	;
	v213 = int64(0)
	v252 = v179
	v253 = v213
	v254 = v213
	v255 = v163
	v256 = v161
	v257 = int32(0)
	v258 = v182
	goto L52
L59:
	;
	v207 = int64(0)
	v208 = int32(0)
	v252 = v205
	v253 = v207
	v254 = v207
	v255 = v207
	v256 = v208
	v257 = v208
	v258 = v206
	goto L52
L60:
	;
	v203 = base.I32_rem_s(base.I32_extend8_s(v200), int32(3))
	v205 = v199
	v206 = v200 - v203
	goto L59
L61:
	;
	v199 = v197
	v200 = int32(0)
	goto L60
L62:
	;
	v195 = base.I32_rem_s(v193, int32(10))
	v197 = v193 - v195
	goto L61
L63:
	;
	v191 = base.I32_rem_s(v189, int32(100))
	v193 = v189 - v191
	goto L62
L64:
	;
	v187 = base.I32_rem_s(v179, int32(1000))
	v189 = v179 - v187
	goto L63
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v227 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v60
	F_errmsg(m, int32(_a_F_interval_trunc_7), v14)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v234 == int32(22) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v239 = F_errdetail(m, int32(_a_F_interval_trunc_8), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_errfinish(m, int32(_a_F_interval_trunc_4), int32(_a_F_interval_trunc_9), int32(_a_F_interval_trunc_6))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L84
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v256
	v269 = base.I32_wrap_i64(v263)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v269
	v272 = v255 * int64(3600000000)
	v275 = base.I64_extend32_s(v253) * int64(60000000)
	v276 = v272 + v275
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v276
	if base.B2i32(v275 < int64(0))^base.B2i32(v276 < v272) != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v284 = base.I64_extend32_s(v254) * int64(1000000)
	v285 = v276 + v284
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v285
	if base.B2i32(v284 < int64(0))^base.B2i32(v285 < v276) != 0 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v291 = base.I64_extend_i32_s(v257)
	v292 = v285 + v291
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v292
	if base.B2i32(v291 < int64(0))^base.B2i32(v292 < v285) != 0 {
		goto L74
	} else {
		goto L78
	}
L78:
	;
	if v269 != int32(2147483647) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v300 = int32(-2147483648)
	if base.B2i32(v269 != v300)|base.B2i32(v256 != v300)|base.B2i32(v292 != int64(-9223372036854775807-1)) != 0 {
		goto L4
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if base.B2i32(v256 != int32(2147483647))|base.B2i32(v292 != int64(9223372036854775807)) != 0 {
		goto L4
	} else {
		goto L83
	}
L82:
	;
	goto L74
L83:
	;
	goto L74
L84:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_interval_trunc_3), int32(0))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_interval_trunc_4), int32(_a_F_interval_trunc_5), int32(_a_F_interval_trunc_6))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_interval_um_internal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v30 int64
	_ = v30
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != int32(-2147483648) {
		if v7 == int32(2147483647) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			if base.B2i32(v21 != int32(2147483647))|base.B2i32(v24 != int64(9223372036854775807)) != 0 {
				v36 = v24
				v37 = int64(0)
				v38 = v37 - v36
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v38
				if base.B2i32(v38 < v37)^base.B2i32(v37 < v36) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v45 = int32(0)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v47 = v45 - v46
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v47
					if base.B2i32(v47 < v45)^base.B2i32(v45 < v46) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v54 = int32(0)
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v56 = v54 - v55
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v56
						if base.B2i32(v56 < v54)^base.B2i32(v54 < v55) != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							if v56 != int32(2147483647) {
								v65 = int32(-2147483648)
								if base.B2i32(v56 != v65)|base.B2i32(v47 != v65)|base.B2i32(v38 != int64(-9223372036854775807-1)) != 0 {
									return
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							} else {
								if base.B2i32(v47 != int32(2147483647))|base.B2i32(v38 != int64(9223372036854775807)) != 0 {
									return
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
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
						}
					}
				}
			} else {
				v30 = int64(-9223372036854775807 - 1)
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v30
				v33 = v7 ^ int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v33
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v33
				return
			}
		} else {
			v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			v36 = v12
			v37 = int64(0)
			v38 = v37 - v36
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v38
			if base.B2i32(v38 < v37)^base.B2i32(v37 < v36) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v45 = int32(0)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v47 = v45 - v46
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v47
				if base.B2i32(v47 < v45)^base.B2i32(v45 < v46) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v54 = int32(0)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v56 = v54 - v55
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v56
					if base.B2i32(v56 < v54)^base.B2i32(v54 < v55) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if v56 != int32(2147483647) {
							v65 = int32(-2147483648)
							if base.B2i32(v56 != v65)|base.B2i32(v47 != v65)|base.B2i32(v38 != int64(-9223372036854775807-1)) != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						} else {
							if base.B2i32(v47 != int32(2147483647))|base.B2i32(v38 != int64(9223372036854775807)) != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
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
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		if base.B2i32(v13 != int32(-2147483648))|base.B2i32(v16 != int64(-9223372036854775807-1)) != 0 {
			v36 = v16
			v37 = int64(0)
			v38 = v37 - v36
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v38
			if base.B2i32(v38 < v37)^base.B2i32(v37 < v36) != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v45 = int32(0)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v47 = v45 - v46
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v47
				if base.B2i32(v47 < v45)^base.B2i32(v45 < v46) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v54 = int32(0)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v56 = v54 - v55
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v56
					if base.B2i32(v56 < v54)^base.B2i32(v54 < v55) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						if v56 != int32(2147483647) {
							v65 = int32(-2147483648)
							if base.B2i32(v56 != v65)|base.B2i32(v47 != v65)|base.B2i32(v38 != int64(-9223372036854775807-1)) != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							}
						} else {
							if base.B2i32(v47 != int32(2147483647))|base.B2i32(v38 != int64(9223372036854775807)) != 0 {
								return
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v88 = m.ExcPending
									if v88 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_interval_um_internal_0), int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3460), int32(_a_F_interval_um_internal_2))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
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
					}
				}
			}
		} else {
			v30 = int64(9223372036854775807)
			*(*int64)(unsafe.Add(mBase, uint32(l1))) = v30
			v33 = v7 ^ int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v33
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v33
			return
		}
	}
}
