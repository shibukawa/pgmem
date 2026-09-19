package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_interval_avg_accum(m *base.Module, l0 int32) int32 {
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
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 != 0 {
			v60 = v12
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v62 != 0 {
				m.G0 = v7 + int32(16)
				return v60
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
				if v64 != int32(2147483647) {
					if v64 != int32(-2147483648) {
						v90 = v60 + int32(8)
						F_finite_interval_pl(m, v90, v63, v90)
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
							*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
							m.G0 = v7 + int32(16)
							return v60
						}
					} else {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
						if v69 != int32(-2147483648) {
							v90 = v60 + int32(8)
							F_finite_interval_pl(m, v90, v63, v90)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
								m.G0 = v7 + int32(16)
								return v60
							}
						} else {
							v72 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
							if v72 != int64(-9223372036854775807-1) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return v60
								}
							} else {
								v75 = *(*int64)(unsafe.Add(mBase, uint32(v60)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v75 + int64(1)
								m.G0 = v7 + int32(16)
								return v60
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
							return int32(0)
						} else {
							v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
							*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
							m.G0 = v7 + int32(16)
							return v60
						}
					} else {
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
						if v82 != int64(9223372036854775807) {
							v90 = v60 + int32(8)
							F_finite_interval_pl(m, v90, v63, v90)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
								m.G0 = v7 + int32(16)
								return v60
							}
						} else {
							v85 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v85 + int64(1)
							m.G0 = v7 + int32(16)
							return v60
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
				switch v20 - int32(429) {
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
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
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
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_interval_avg_accum_0), int32(0))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_interval_avg_accum_1), int32(3992), int32(_a_F_interval_avg_accum_2))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
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
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_accum[0])) = v49
					v60 = v54
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v62 != 0 {
						m.G0 = v7 + int32(16)
						return v60
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
						if v64 != int32(2147483647) {
							if v64 != int32(-2147483648) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return v60
								}
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
								if v69 != int32(-2147483648) {
									v90 = v60 + int32(8)
									F_finite_interval_pl(m, v90, v63, v90)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
										m.G0 = v7 + int32(16)
										return v60
									}
								} else {
									v72 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
									if v72 != int64(-9223372036854775807-1) {
										v90 = v60 + int32(8)
										F_finite_interval_pl(m, v90, v63, v90)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
											*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
											m.G0 = v7 + int32(16)
											return v60
										}
									} else {
										v75 = *(*int64)(unsafe.Add(mBase, uint32(v60)+32))
										*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v75 + int64(1)
										m.G0 = v7 + int32(16)
										return v60
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
									return int32(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return v60
								}
							} else {
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
								if v82 != int64(9223372036854775807) {
									v90 = v60 + int32(8)
									F_finite_interval_pl(m, v90, v63, v90)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
										m.G0 = v7 + int32(16)
										return v60
									}
								} else {
									v85 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
									*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v85 + int64(1)
									m.G0 = v7 + int32(16)
									return v60
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
			switch v20 - int32(429) {
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
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
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
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_interval_avg_accum_0), int32(0))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_interval_avg_accum_1), int32(3992), int32(_a_F_interval_avg_accum_2))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
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
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_interval_avg_accum[0])) = v49
				v60 = v54
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v62 != 0 {
					m.G0 = v7 + int32(16)
					return v60
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
					if v64 != int32(2147483647) {
						if v64 != int32(-2147483648) {
							v90 = v60 + int32(8)
							F_finite_interval_pl(m, v90, v63, v90)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
								m.G0 = v7 + int32(16)
								return v60
							}
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							if v69 != int32(-2147483648) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return v60
								}
							} else {
								v72 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
								if v72 != int64(-9223372036854775807-1) {
									v90 = v60 + int32(8)
									F_finite_interval_pl(m, v90, v63, v90)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int32(0)
									} else {
										v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
										*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
										m.G0 = v7 + int32(16)
										return v60
									}
								} else {
									v75 = *(*int64)(unsafe.Add(mBase, uint32(v60)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v75 + int64(1)
									m.G0 = v7 + int32(16)
									return v60
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
								return int32(0)
							} else {
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
								*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
								m.G0 = v7 + int32(16)
								return v60
							}
						} else {
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
							if v82 != int64(9223372036854775807) {
								v90 = v60 + int32(8)
								F_finite_interval_pl(m, v90, v63, v90)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int32(0)
								} else {
									v93 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
									*(*int64)(unsafe.Add(mBase, uint32(v60))) = v93 + int64(1)
									m.G0 = v7 + int32(16)
									return v60
								}
							} else {
								v85 = *(*int64)(unsafe.Add(mBase, uint32(v60)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v85 + int64(1)
								m.G0 = v7 + int32(16)
								return v60
							}
						}
					}
				}
			}
		}
	}
}
func F_interval_avg_accum_inv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v29 int32
	_ = v29
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int64
	_ = v53
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_interval_avg_accum_inv_0), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_interval_avg_accum_inv_1), int32(_a_F_interval_avg_accum_inv_2), int32(_a_F_interval_avg_accum_inv_3))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v6 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_interval_avg_accum_inv_0), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_interval_avg_accum_inv_1), int32(_a_F_interval_avg_accum_inv_2), int32(_a_F_interval_avg_accum_inv_3))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v9 == int32(0) {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				if v13 != int32(2147483647) {
					if v13 != int32(-2147483648) {
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
						v42 = v40 - int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v6))) = v42
						v45 = v6 + int32(8)
						if int64(0) < v42 {
							F_finite_interval_mi(m, v45, v12, v45)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								return v6
							}
						} else {
							v53 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
							*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
							return v6
						}
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v18 != int32(-2147483648) {
							v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
							v42 = v40 - int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v6))) = v42
							v45 = v6 + int32(8)
							if int64(0) < v42 {
								F_finite_interval_mi(m, v45, v12, v45)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									return v6
								}
							} else {
								v53 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
								*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
								return v6
							}
						} else {
							v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
							if v21 != int64(-9223372036854775807-1) {
								v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
								v42 = v40 - int64(1)
								*(*int64)(unsafe.Add(mBase, uint32(v6))) = v42
								v45 = v6 + int32(8)
								if int64(0) < v42 {
									F_finite_interval_mi(m, v45, v12, v45)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										return v6
									}
								} else {
									v53 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
									*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
									return v6
								}
							} else {
								v24 = *(*int64)(unsafe.Add(mBase, uint32(v6)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v24 - int64(1)
								return v6
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					if v29 != int32(2147483647) {
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
						v42 = v40 - int64(1)
						*(*int64)(unsafe.Add(mBase, uint32(v6))) = v42
						v45 = v6 + int32(8)
						if int64(0) < v42 {
							F_finite_interval_mi(m, v45, v12, v45)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								return v6
							}
						} else {
							v53 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
							*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
							return v6
						}
					} else {
						v32 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
						if v32 != int64(9223372036854775807) {
							v40 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
							v42 = v40 - int64(1)
							*(*int64)(unsafe.Add(mBase, uint32(v6))) = v42
							v45 = v6 + int32(8)
							if int64(0) < v42 {
								F_finite_interval_mi(m, v45, v12, v45)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									return v6
								}
							} else {
								v53 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
								*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
								return v6
							}
						} else {
							v35 = *(*int64)(unsafe.Add(mBase, uint32(v6)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v35 - int64(1)
							return v6
						}
					}
				}
			} else {
				return v6
			}
		}
	}
}
func F_interval_avg_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
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
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(429) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v42 = F_pg_detoast_datum_packed(m, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = int32(1)
			v47 = v42 + v46
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v50 = v48 & v46
			if v48 == v46 {
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
				if v56 == int32(18) {
					v59 = int32(16)
				} else {
					v59 = int32(0)
				}
				if base.Ui32((v56-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v66 = int32(4)
				} else {
					v66 = v59
				}
				v77 = v66
			} else {
				v67 = int32(1)
				if v50 != 0 {
					v77 = int32(base.Ui32(v48)>>(uint(v67)%32)) - v67
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v77 = int32(base.Ui32(v71)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v77
			if v50 != 0 {
				v83 = v47
			} else {
				v83 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v83
			v86 = F_palloc0(m, int32(40))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				v88 = F_pq_getmsgint64(m, v8)
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v86))) = v88
					v91 = F_pq_getmsgint64(m, v8)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = v91
						v95 = F_pq_getmsgint(m, v8, int32(4))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v95
							v99 = F_pq_getmsgint(m, v8, int32(4))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = v99
								v102 = F_pq_getmsgint64(m, v8)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v86)+24)) = v102
									v105 = F_pq_getmsgint64(m, v8)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v86)+32)) = v105
										F_pq_getmsgend(m, v8)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											m.G0 = v8 + int32(16)
											return v86
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
		v117 = m.ExcPending
		if v117 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_interval_avg_deserialize_0), int32(0))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_interval_avg_deserialize_1), int32(_a_F_interval_avg_deserialize_2), int32(_a_F_interval_avg_deserialize_3))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
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
func F_interval_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int64
	_ = v22
	var v23 int64
	_ = v23
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v53 int64
	_ = v53
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
	v17 = v14 + int32(16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+12)))
	v22 = int64(*(*int32)(unsafe.Add(mBase, uint32(v18)+8)))
	v23 = v19*int64(30) + v22
	v32 = int64(32)
	v33 = int64(20)
	v35 = int64(base.Ui64(v23) >> (uint(v32) % 64))
	v38 = int64(4294967295)
	v39 = int64(500654080)
	v41 = v23 & v38
	v42 = v39 * v41
	v46 = int64(base.Ui64(v42)>>(uint(v32)%64)) + v39*v35
	v53 = v41*v33 + v46&v38
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v23*int64(0) + v23>>(uint(int64(63))%64)*int64(86400000000) + v33*v35 + int64(base.Ui64(v46)>>(uint(v32)%64)) + int64(base.Ui64(v53)>>(uint(v32)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v42&v38 | v53<<(uint(v32)%64)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v69*int64(0) + v69>>(uint(int64(63))%64)*int64(86400000000) + v79*v81 + int64(base.Ui64(v92)>>(uint(v78)%64)) + int64(base.Ui64(v99)>>(uint(v78)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v88&v84 | v99<<(uint(v78)%64)
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
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
	return v136 - v139
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
	v4 = int32(8)
	v8 = F_range_cmp_bounds(m, l2, l0+v4, l1+v4)
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_interval_justify_interval(m *base.Module, l0 int32) int32 {
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
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L51
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	return v10
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
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_interval_justify_interval_0), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_interval_justify_interval_1), int32(2964), int32(_a_F_interval_justify_interval_2))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
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
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errmsg(m, int32(_a_F_interval_justify_interval_0), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_interval_justify_interval_1), int32(2981), int32(_a_F_interval_justify_interval_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
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
func F_interval_mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 float64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v91 int64
	_ = v91
	var v98 int64
	_ = v98
	var v109 int64
	_ = v109
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v145 float64
	_ = v145
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 float64
	_ = v165
	var v168 int32
	_ = v168
	var v187 float64
	_ = v187
	var v191 float64
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 float64
	_ = v201
	var v207 float64
	_ = v207
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v226 int32
	_ = v226
	var v227 float64
	_ = v227
	var v229 int32
	_ = v229
	var v235 int64
	_ = v235
	var v241 float64
	_ = v241
	var v244 int32
	_ = v244
	var v257 int64
	_ = v257
	var v261 int32
	_ = v261
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	v15 = m.G0
	v16 = int32(16)
	v17 = v15 - v16
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = *(*float64)(unsafe.Add(mBase, uint32(v19)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v25 = F_palloc(m, v16)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v20)&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v17 + int32(16)
	return v25
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L47
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v34 != int32(2147483647) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	if base.F64_eq(base.F64_abs(v20), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L21
	} else {
		goto L22
	}
L7:
	;
	if base.F64_eq(v20, float64(0)) != 0 {
		goto L4
	} else {
		goto L16
	}
L8:
	;
	if v34 != int32(-2147483648) {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v45 != int32(2147483647) {
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v39 != int32(-2147483648) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	if v42 == int64(-9223372036854775807-1) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L6
L14:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	if v48 != int64(9223372036854775807) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L7
L16:
	;
	if base.F64_lt(v20, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	F_interval_um_internal(m, v21, v25)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v59
	goto L3
L20:
	;
	goto L3
L21:
	;
	v64 = int64(*(*int32)(unsafe.Add(mBase, uint32(v21)+8)))
	v68 = v64 + base.I64_extend_i32_s(v34)*int64(30)
	v77 = int64(32)
	v78 = int64(20)
	v80 = int64(base.Ui64(v68) >> (uint(v77) % 64))
	v83 = int64(4294967295)
	v84 = int64(500654080)
	v86 = v68 & v83
	v87 = v84 * v86
	v91 = int64(base.Ui64(v87)>>(uint(v77)%64)) + v84*v80
	v98 = v86*v78 + v91&v83
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v68*int64(0) + v68>>(uint(int64(63))%64)*int64(86400000000) + v78*v80 + int64(base.Ui64(v91)>>(uint(v77)%64)) + int64(base.Ui64(v98)>>(uint(v77)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v87&v83 | v98<<(uint(v77)%64)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v145 = base.F64_mul(v20, base.F64_convert_i32_s(v34))
	v148 = int32(0)
	if base.B2i32(base.F64_lt(v145, float64(2.147483648e+09)) == v148)|base.B2i32(base.F64_ge(v145, float64(-2.147483648e+09)) == v148)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v145)&int64(9223372036854775807))) != 0 {
		goto L4
	} else {
		goto L32
	}
L24:
	;
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v111 = v109 + v110
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v118 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v111) < base.Ui64(v109))) + (v114 + v110>>(uint(int64(63))%64))
	if v111|v118 == int64(0) {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v122 = int64(0)
	if v118 == v122 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v128 = base.B2i32(v111 != v122)
	goto L28
L27:
	;
	v128 = base.B2i32(v122 < v118)
	goto L28
L28:
	;
	if base.F64_lt(base.F64_mul(v20, base.F64_convert_i32_s(v128-base.B2i32(v118 < int64(0)))), float64(0)) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = int64(-9223372034707292160)
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(-9223372036854775807 - 1)
	goto L3
L30:
	;
	goto L31
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = int64(9223372034707292159)
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = int64(9223372036854775807)
	goto L3
L32:
	;
	v161 = base.I32_trunc_sat_f64_s(v145)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v161
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v165 = base.F64_mul(v20, base.F64_convert_i32_s(v163))
	v168 = int32(0)
	if base.B2i32(base.F64_lt(v165, float64(2.147483648e+09)) == v168)|base.B2i32(base.F64_ge(v165, float64(-2.147483648e+09)) == v168)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v165)&int64(9223372036854775807))) != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v187 = float64(1e+06)
	v191 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_mul(base.F64_convert_i32_s(v23), v20), base.F64_convert_i32_s(v161)), float64(30)), v187)), v187)
	v194 = base.I32_trunc_sat_f64_s(v165)
	v198 = base.I32_trunc_sat_f64_s(v191)
	v201 = float64(86400)
	v207 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_add(v191, base.F64_sub(base.F64_mul(base.F64_convert_i32_s(v22), v20), base.F64_convert_i32_s(v194))), base.F64_convert_i32_s(v198)), v201), v187)), v187)
	if base.F64_ge(base.F64_abs(v207), v201) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v229 = v226 + v198
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v229
	if base.B2i32(v198 < int32(0))^base.B2i32(v229 < v226) != 0 {
		goto L4
	} else {
		goto L39
	}
L35:
	;
	v226 = v194
	v227 = v207
	goto L34
L36:
	;
	goto L37
L37:
	;
	v215 = base.I32_trunc_sat_f64_s(base.F64_div(v207, float64(86400)))
	v216 = v194 + v215
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v216
	if base.B2i32(v215 < int32(0))^base.B2i32(v216 < v194) != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v226 = v216
	v227 = base.F64_sub(v207, base.F64_convert_i32_s(v215*int32(_a_F_interval_mul_3)))
	goto L34
L39:
	;
	v235 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v241 = base.F64_nearest(base.F64_add(base.F64_mul(base.F64_convert_i64_s(v235), v20), base.F64_mul(v227, float64(1e+06))))
	v244 = int32(0)
	if base.B2i32(base.F64_lt(v241, float64(9.223372036854776e+18)) == v244)|base.B2i32(base.F64_ge(v241, float64(-9.223372036854776e+18)) == v244)|base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v241)&int64(9223372036854775807))) != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v257 = base.I64_trunc_sat_f64_s(v241)
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v257
	if v161 != int32(2147483647) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v261 = int32(-2147483648)
	if base.B2i32(v161 != v261)|base.B2i32(v229 != v261) != 0 {
		goto L3
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if base.B2i32(v229 != int32(2147483647))|base.B2i32(v257 != int64(9223372036854775807)) != 0 {
		goto L3
	} else {
		goto L46
	}
L44:
	;
	if v257 == int64(-9223372036854775807-1) {
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
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(_a_F_interval_mul_0), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_interval_mul_1), int32(3740), int32(_a_F_interval_mul_2))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
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
func F_interval_recv(m *base.Module, l0 int32) int32 {
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
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_palloc(m, int32(16))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_pq_getmsgint64(m, v5)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = v11
			v15 = F_pq_getmsgint(m, v5, int32(4))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v15
				v19 = F_pq_getmsgint(m, v5, int32(4))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v19
					F_AdjustIntervalForTypmod(m, v7, v4, int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v7
					}
				}
			}
		}
	}
}
func F_interval_to_char(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v122 int64
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
		if v21 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
				if v52 != int32(2147483647) {
					if v52 != int32(-2147483648) {
						v73 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
						*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
						v77 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
						v85 = v13 + int32(24)
						v87 = v13 + int32(8)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
						v89 = int32(12)
						v90 = base.I32_div_s(v88, v89)
						*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
						*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						v100 = base.I64_div_s(v98, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
						v104 = v98 + v100*int64(-3600000000)
						v106 = base.I64_div_s(v104, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
						v110 = v106*int64(-60000000) + v104
						v112 = base.I64_div_s(v110, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
						v116 = v112*int64(4293967296) + v110
						*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
						v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							if v141 == int32(0) {
								v145 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
								v148 = v77
							} else {
								v148 = v141
							}
							m.G0 = v13 + int32(112)
							return v148
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						if v57 != int32(-2147483648) {
							v73 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
							v85 = v13 + int32(24)
							v87 = v13 + int32(8)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
							v89 = int32(12)
							v90 = base.I32_div_s(v88, v89)
							*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
							*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							v100 = base.I64_div_s(v98, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
							v104 = v98 + v100*int64(-3600000000)
							v106 = base.I64_div_s(v104, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
							v110 = v106*int64(-60000000) + v104
							v112 = base.I64_div_s(v110, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
							v116 = v112*int64(4293967296) + v110
							*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
							v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								if v141 == int32(0) {
									v145 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
									v148 = v77
								} else {
									v148 = v141
								}
								m.G0 = v13 + int32(112)
								return v148
							}
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							if v60 == int64(-9223372036854775807-1) {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v148 = int32(0)
								m.G0 = v13 + int32(112)
								return v148
							} else {
								v73 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
								*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
								v85 = v13 + int32(24)
								v87 = v13 + int32(8)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
								v89 = int32(12)
								v90 = base.I32_div_s(v88, v89)
								*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
								v100 = base.I64_div_s(v98, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
								v104 = v98 + v100*int64(-3600000000)
								v106 = base.I64_div_s(v104, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
								v110 = v106*int64(-60000000) + v104
								v112 = base.I64_div_s(v110, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
								v116 = v112*int64(4293967296) + v110
								*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
								*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									if v141 == int32(0) {
										v145 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
										v148 = v77
									} else {
										v148 = v141
									}
									m.G0 = v13 + int32(112)
									return v148
								}
							}
						}
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					if v63 != int32(2147483647) {
						v73 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
						*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
						v77 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
						v85 = v13 + int32(24)
						v87 = v13 + int32(8)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
						v89 = int32(12)
						v90 = base.I32_div_s(v88, v89)
						*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
						*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						v100 = base.I64_div_s(v98, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
						v104 = v98 + v100*int64(-3600000000)
						v106 = base.I64_div_s(v104, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
						v110 = v106*int64(-60000000) + v104
						v112 = base.I64_div_s(v110, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
						v116 = v112*int64(4293967296) + v110
						*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
						v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							if v141 == int32(0) {
								v145 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
								v148 = v77
							} else {
								v148 = v141
							}
							m.G0 = v13 + int32(112)
							return v148
						}
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						if v66 != int64(9223372036854775807) {
							v73 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
							v85 = v13 + int32(24)
							v87 = v13 + int32(8)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
							v89 = int32(12)
							v90 = base.I32_div_s(v88, v89)
							*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
							*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							v100 = base.I64_div_s(v98, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
							v104 = v98 + v100*int64(-3600000000)
							v106 = base.I64_div_s(v104, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
							v110 = v106*int64(-60000000) + v104
							v112 = base.I64_div_s(v110, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
							v116 = v112*int64(4293967296) + v110
							*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
							v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								if v141 == int32(0) {
									v145 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
									v148 = v77
								} else {
									v148 = v141
								}
								m.G0 = v13 + int32(112)
								return v148
							}
						} else {
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
							v148 = int32(0)
							m.G0 = v13 + int32(112)
							return v148
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
					v148 = int32(0)
					m.G0 = v13 + int32(112)
					return v148
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
					if v52 != int32(2147483647) {
						if v52 != int32(-2147483648) {
							v73 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
							v85 = v13 + int32(24)
							v87 = v13 + int32(8)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
							v89 = int32(12)
							v90 = base.I32_div_s(v88, v89)
							*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
							*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							v100 = base.I64_div_s(v98, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
							v104 = v98 + v100*int64(-3600000000)
							v106 = base.I64_div_s(v104, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
							v110 = v106*int64(-60000000) + v104
							v112 = base.I64_div_s(v110, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
							v116 = v112*int64(4293967296) + v110
							*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
							v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								if v141 == int32(0) {
									v145 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
									v148 = v77
								} else {
									v148 = v141
								}
								m.G0 = v13 + int32(112)
								return v148
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							if v57 != int32(-2147483648) {
								v73 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
								*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
								v85 = v13 + int32(24)
								v87 = v13 + int32(8)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
								v89 = int32(12)
								v90 = base.I32_div_s(v88, v89)
								*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
								v100 = base.I64_div_s(v98, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
								v104 = v98 + v100*int64(-3600000000)
								v106 = base.I64_div_s(v104, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
								v110 = v106*int64(-60000000) + v104
								v112 = base.I64_div_s(v110, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
								v116 = v112*int64(4293967296) + v110
								*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
								*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									if v141 == int32(0) {
										v145 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
										v148 = v77
									} else {
										v148 = v141
									}
									m.G0 = v13 + int32(112)
									return v148
								}
							} else {
								v60 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								if v60 == int64(-9223372036854775807-1) {
									v70 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
									v148 = int32(0)
									m.G0 = v13 + int32(112)
									return v148
								} else {
									v73 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
									*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
									v77 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
									v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
									v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
									v85 = v13 + int32(24)
									v87 = v13 + int32(8)
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
									v89 = int32(12)
									v90 = base.I32_div_s(v88, v89)
									*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
									*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
									v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
									v100 = base.I64_div_s(v98, int64(3600000000))
									*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
									v104 = v98 + v100*int64(-3600000000)
									v106 = base.I64_div_s(v104, int64(60000000))
									*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
									v110 = v106*int64(-60000000) + v104
									v112 = base.I64_div_s(v110, int64(1000000))
									*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
									v116 = v112*int64(4293967296) + v110
									*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
									v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
									v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
									*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
									v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
									mBase = m.M
									v142 = m.ExcPending
									if v142 != 0 {
										return int32(0)
									} else {
										if v141 == int32(0) {
											v145 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
											v148 = v77
										} else {
											v148 = v141
										}
										m.G0 = v13 + int32(112)
										return v148
									}
								}
							}
						}
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						if v63 != int32(2147483647) {
							v73 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
							v85 = v13 + int32(24)
							v87 = v13 + int32(8)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
							v89 = int32(12)
							v90 = base.I32_div_s(v88, v89)
							*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
							*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							v100 = base.I64_div_s(v98, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
							v104 = v98 + v100*int64(-3600000000)
							v106 = base.I64_div_s(v104, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
							v110 = v106*int64(-60000000) + v104
							v112 = base.I64_div_s(v110, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
							v116 = v112*int64(4293967296) + v110
							*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
							v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								if v141 == int32(0) {
									v145 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
									v148 = v77
								} else {
									v148 = v141
								}
								m.G0 = v13 + int32(112)
								return v148
							}
						} else {
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							if v66 != int64(9223372036854775807) {
								v73 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
								*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
								v85 = v13 + int32(24)
								v87 = v13 + int32(8)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
								v89 = int32(12)
								v90 = base.I32_div_s(v88, v89)
								*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
								v100 = base.I64_div_s(v98, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
								v104 = v98 + v100*int64(-3600000000)
								v106 = base.I64_div_s(v104, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
								v110 = v106*int64(-60000000) + v104
								v112 = base.I64_div_s(v110, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
								v116 = v112*int64(4293967296) + v110
								*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
								*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									if v141 == int32(0) {
										v145 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
										v148 = v77
									} else {
										v148 = v141
									}
									m.G0 = v13 + int32(112)
									return v148
								}
							} else {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v148 = int32(0)
								m.G0 = v13 + int32(112)
								return v148
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
				v148 = int32(0)
				m.G0 = v13 + int32(112)
				return v148
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
				if v52 != int32(2147483647) {
					if v52 != int32(-2147483648) {
						v73 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
						*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
						v77 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
						v85 = v13 + int32(24)
						v87 = v13 + int32(8)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
						v89 = int32(12)
						v90 = base.I32_div_s(v88, v89)
						*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
						*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						v100 = base.I64_div_s(v98, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
						v104 = v98 + v100*int64(-3600000000)
						v106 = base.I64_div_s(v104, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
						v110 = v106*int64(-60000000) + v104
						v112 = base.I64_div_s(v110, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
						v116 = v112*int64(4293967296) + v110
						*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
						v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							if v141 == int32(0) {
								v145 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
								v148 = v77
							} else {
								v148 = v141
							}
							m.G0 = v13 + int32(112)
							return v148
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
						if v57 != int32(-2147483648) {
							v73 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
							v85 = v13 + int32(24)
							v87 = v13 + int32(8)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
							v89 = int32(12)
							v90 = base.I32_div_s(v88, v89)
							*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
							*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							v100 = base.I64_div_s(v98, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
							v104 = v98 + v100*int64(-3600000000)
							v106 = base.I64_div_s(v104, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
							v110 = v106*int64(-60000000) + v104
							v112 = base.I64_div_s(v110, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
							v116 = v112*int64(4293967296) + v110
							*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
							v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								if v141 == int32(0) {
									v145 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
									v148 = v77
								} else {
									v148 = v141
								}
								m.G0 = v13 + int32(112)
								return v148
							}
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							if v60 == int64(-9223372036854775807-1) {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v148 = int32(0)
								m.G0 = v13 + int32(112)
								return v148
							} else {
								v73 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
								*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
								v77 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
								v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
								v85 = v13 + int32(24)
								v87 = v13 + int32(8)
								v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
								v89 = int32(12)
								v90 = base.I32_div_s(v88, v89)
								*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
								*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
								v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
								v100 = base.I64_div_s(v98, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
								v104 = v98 + v100*int64(-3600000000)
								v106 = base.I64_div_s(v104, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
								v110 = v106*int64(-60000000) + v104
								v112 = base.I64_div_s(v110, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
								v116 = v112*int64(4293967296) + v110
								*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
								v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
								v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
								*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
								v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return int32(0)
								} else {
									if v141 == int32(0) {
										v145 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
										v148 = v77
									} else {
										v148 = v141
									}
									m.G0 = v13 + int32(112)
									return v148
								}
							}
						}
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					if v63 != int32(2147483647) {
						v73 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
						*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
						v77 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
						v85 = v13 + int32(24)
						v87 = v13 + int32(8)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
						v89 = int32(12)
						v90 = base.I32_div_s(v88, v89)
						*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
						*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
						v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
						v100 = base.I64_div_s(v98, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
						v104 = v98 + v100*int64(-3600000000)
						v106 = base.I64_div_s(v104, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
						v110 = v106*int64(-60000000) + v104
						v112 = base.I64_div_s(v110, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
						v116 = v112*int64(4293967296) + v110
						*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
						v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
						v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
						v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
						*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							if v141 == int32(0) {
								v145 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
								v148 = v77
							} else {
								v148 = v141
							}
							m.G0 = v13 + int32(112)
							return v148
						}
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						if v66 != int64(9223372036854775807) {
							v73 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v13)+96)) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v13)+88)) = v73
							v77 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v13)+108)) = v77
							v80 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v15)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
							v85 = v13 + int32(24)
							v87 = v13 + int32(8)
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
							v89 = int32(12)
							v90 = base.I32_div_s(v88, v89)
							*(*int32)(unsafe.Add(mBase, uint32(v85)+32)) = v90
							*(*int32)(unsafe.Add(mBase, uint32(v85)+28)) = v88 - v90*v89
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v85)+24)) = v96
							v98 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
							v100 = base.I64_div_s(v98, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v85)+16)) = v100
							v104 = v98 + v100*int64(-3600000000)
							v106 = base.I64_div_s(v104, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+8)) = uint32(v106)
							v110 = v106*int64(-60000000) + v104
							v112 = base.I64_div_s(v110, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v85)+4)) = uint32(v112)
							v116 = v112*int64(4293967296) + v110
							*(*uint32)(unsafe.Add(mBase, uint32(v85))) = uint32(v116)
							v118 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v118
							v120 = *(*int64)(unsafe.Add(mBase, uint32(v13)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v120
							v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v13)+72)) = v122
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v128
							*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v124 + (v126+v128*v89)*int32(30)
							v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v141 = F_datetime_to_char_body(m, v13-int32(-64), v17, int32(1), v140)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								if v141 == int32(0) {
									v145 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v145)
									v148 = v77
								} else {
									v148 = v141
								}
								m.G0 = v13 + int32(112)
								return v148
							}
						} else {
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
							v148 = int32(0)
							m.G0 = v13 + int32(112)
							return v148
						}
					}
				}
			}
		}
	}
}
func F_interval_trunc(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v162 int32
	_ = v162
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int64
	_ = v208
	var v209 int32
	_ = v209
	var v214 int64
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int64
	_ = v247
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v253 int32
	_ = v253
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v256 int64
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int64
	_ = v264
	var v270 int32
	_ = v270
	var v273 int64
	_ = v273
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v285 int64
	_ = v285
	var v286 int64
	_ = v286
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v301 int32
	_ = v301
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	v26 = v17 + v25
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v31 = v29 & v25
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	m.G0 = v14 + int32(48)
	return v23
L5:
	;
	v164 = base.I64_div_s(v161, int64(3600000000))
	v167 = v164*int64(-3600000000) + v161
	v169 = base.I64_div_s(v167, int64(60000000))
	v172 = v169*int64(-60000000) + v167
	v174 = base.I64_div_s(v172, int64(1000000))
	v178 = base.I32_wrap_i64(v174*int64(4293967296) + v172)
	v179 = int32(12)
	v180 = base.I32_div_s(v71, v179)
	v183 = v71 - v180*v179
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	switch v184 - int32(18) {
	case 0:
		goto L57
	case 1:
		v249 = v169
		v250 = v164
		goto L53
	case 2:
		goto L58
	case 3:
		goto L54
	default:
		goto L55
	case 5:
		v206 = v180
		v207 = v183
		goto L59
	case 6:
		v200 = v180
		v201 = v183
		goto L60
	case 7:
		v198 = v180
		goto L61
	case 8:
		v194 = v180
		goto L62
	case 9:
		v190 = v180
		goto L63
	case 10:
		goto L64
	case 11:
		goto L56
	case 12:
		v253 = v180
		v254 = v169
		v255 = v174
		v256 = v164
		v257 = v162
		v258 = v183
		v259 = v178
		goto L52
	}
L6:
	;
	v32 = v26
	goto L8
L7:
	;
	v32 = v17 + int32(4)
	goto L8
L8:
	;
	if v29 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v61 = F_downcase_truncate_identifier(m, v32, v59, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v38 == int32(18) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v49 = int32(1)
	if v31 != 0 {
		v59 = int32(base.Ui32(v29)>>(uint(v49)%32)) - v49
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v41 = int32(16)
	goto L15
L14:
	;
	v41 = int32(0)
	goto L15
L15:
	;
	if base.Ui32((v38-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v48 = int32(4)
	goto L18
L17:
	;
	v48 = v41
	goto L18
L18:
	;
	v59 = v48
	goto L9
L19:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L20:
	;
	v68 = Fn13846(m, v61, v14+int32(44), int32(_a_F_interval_trunc_0), int32(_a_F_interval_trunc_1), int32(_a_F_interval_trunc_2))
	mBase = m.M
	goto L21
L21:
	;
	if v68 == int32(17) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v71 != int32(-2147483648) {
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
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L47
	}
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if base.B2i32(base.Ui32(int32(8)) <= base.Ui32(v94-int32(23)))&base.B2i32(base.Ui32(int32(3)) < base.Ui32(v94-int32(18))) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L26:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v86 != int32(2147483647) {
		v161 = v85
		v162 = v86
		goto L5
	} else {
		goto L33
	}
L27:
	;
	if v71 == int32(2147483647) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v79 != int32(-2147483648) {
		v161 = v78
		v162 = v79
		goto L5
	} else {
		goto L31
	}
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v161 = v77
	v162 = v76
	goto L5
L31:
	;
	if v78 == int64(-9223372036854775807-1) {
		goto L25
	} else {
		goto L32
	}
L32:
	;
	v161 = v78
	v162 = int32(-2147483648)
	goto L5
L33:
	;
	if v85 == int64(9223372036854775807) {
		goto L25
	} else {
		goto L34
	}
L34:
	;
	v161 = v85
	v162 = int32(2147483647)
	goto L5
L35:
	;
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v108
	goto L4
L36:
	;
	goto L37
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v118 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v61
	F_errmsg(m, int32(_a_F_interval_trunc_7), v14+int32(16))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v127 == int32(22) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	F_errdetail(m, int32(_a_F_interval_trunc_8), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
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
	v138 = m.ExcPending
	if v138 != 0 {
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
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v147 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v61
	F_errmsg(m, int32(_a_F_interval_trunc_11), v14+int32(32))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_interval_trunc_4), int32(_a_F_interval_trunc_12), int32(_a_F_interval_trunc_6))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	v264 = base.I64_extend_i32_s(v258) + base.I64_extend_i32_s(v253)*int64(12)
	if base.Ui64(v264-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L74
	} else {
		goto L75
	}
L53:
	;
	v253 = v180
	v254 = v249
	v255 = int64(0)
	v256 = v250
	v257 = v162
	v258 = v183
	v259 = int32(0)
	goto L52
L54:
	;
	v247 = int64(0)
	v249 = v247
	v250 = v247
	goto L53
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L65
	}
L56:
	;
	v218 = base.I32_rem_s(v178, int32(1000))
	v253 = v180
	v254 = v169
	v255 = v174
	v256 = v164
	v257 = v162
	v258 = v183
	v259 = v178 - v218
	goto L52
L57:
	;
	v253 = v180
	v254 = v169
	v255 = v174
	v256 = v164
	v257 = v162
	v258 = v183
	v259 = int32(0)
	goto L52
L58:
	;
	v214 = int64(0)
	v253 = v180
	v254 = v214
	v255 = v214
	v256 = v164
	v257 = v162
	v258 = v183
	v259 = int32(0)
	goto L52
L59:
	;
	v208 = int64(0)
	v209 = int32(0)
	v253 = v206
	v254 = v208
	v255 = v208
	v256 = v208
	v257 = v209
	v258 = v207
	v259 = v209
	goto L52
L60:
	;
	v204 = base.I32_rem_s(base.I32_extend8_s(v201), int32(3))
	v206 = v200
	v207 = v201 - v204
	goto L59
L61:
	;
	v200 = v198
	v201 = int32(0)
	goto L60
L62:
	;
	v196 = base.I32_rem_s(v194, int32(10))
	v198 = v194 - v196
	goto L61
L63:
	;
	v192 = base.I32_rem_s(v190, int32(100))
	v194 = v190 - v192
	goto L62
L64:
	;
	v188 = base.I32_rem_s(v180, int32(1000))
	v190 = v180 - v188
	goto L63
L65:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v228 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v228
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v61
	F_errmsg(m, int32(_a_F_interval_trunc_7), v14)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v235 == int32(22) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_errdetail(m, int32(_a_F_interval_trunc_8), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
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
	v246 = m.ExcPending
	if v246 != 0 {
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
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L84
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v257
	v270 = base.I32_wrap_i64(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v270
	v273 = v256 * int64(3600000000)
	v276 = base.I64_extend32_s(v254) * int64(60000000)
	v277 = v273 + v276
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v277
	if base.B2i32(v276 < int64(0))^base.B2i32(v277 < v273) != 0 {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v285 = base.I64_extend32_s(v255) * int64(1000000)
	v286 = v277 + v285
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v286
	if base.B2i32(v285 < int64(0))^base.B2i32(v286 < v277) != 0 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v292 = base.I64_extend_i32_s(v259)
	v293 = v286 + v292
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v293
	if base.B2i32(v292 < int64(0))^base.B2i32(v293 < v286) != 0 {
		goto L74
	} else {
		goto L78
	}
L78:
	;
	if v270 != int32(2147483647) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v301 = int32(-2147483648)
	if base.B2i32(v270 != v301)|base.B2i32(v257 != v301)|base.B2i32(v293 != int64(-9223372036854775807-1)) != 0 {
		goto L4
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	if base.B2i32(v257 != int32(2147483647))|base.B2i32(v293 != int64(9223372036854775807)) != 0 {
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
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_interval_trunc_3), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_interval_trunc_4), int32(_a_F_interval_trunc_5), int32(_a_F_interval_trunc_6))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
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
								F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
									F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
										F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
												F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
												F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
							F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
								F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
									F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
											F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
											F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
							F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
								F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
									F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
											F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
											F_errfinish(m, int32(_a_F_interval_um_internal_1), int32(3459), int32(_a_F_interval_um_internal_2))
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
