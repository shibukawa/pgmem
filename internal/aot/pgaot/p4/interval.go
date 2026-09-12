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
					F_errmsg_internal(m, int32(66371), int32(0))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515856), int32(3992), int32(368746))
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
				v48 = int32(4549024)
				v49 = *(*int32)(unsafe.Add(mBase, _consts[28]))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v51
				v54 = F_palloc0(m, int32(40))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v49
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
				F_errmsg_internal(m, int32(66371), int32(0))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515856), int32(3992), int32(368746))
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
			v48 = int32(4549024)
			v49 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, _consts[28])) = v51
			v54 = F_palloc0(m, int32(40))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v49
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
			F_errmsg_internal(m, int32(367967), int32(0))
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(515856), int32(4216), int32(33794))
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
				F_errmsg_internal(m, int32(367967), int32(0))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515856), int32(4216), int32(33794))
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
							*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
							*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
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
								*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
								*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
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
									*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
									*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
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
							*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
							*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
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
								*(*int64)(unsafe.Add(mBase, uint32(v45))) = v53
								*(*int64)(unsafe.Add(mBase, uint32(v45)+8)) = v53
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
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
				v53 = int32(4)
				v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47))))
				if v55&int32(254) == int32(2) {
					v64 = v53
				} else {
					v64 = base.B2i32(v55 == int32(18)) << (uint(v53) % 32)
				}
				if v55 == int32(1) {
					v67 = v53
				} else {
					v67 = v64
				}
				v78 = v67
			} else {
				v68 = int32(1)
				if v50 != 0 {
					v78 = int32(base.Ui32(v48)>>(uint(v68)%32)) - v68
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v78 = int32(base.Ui32(v72)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v78
			if v50 != 0 {
				v84 = v47
			} else {
				v84 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v84
			v87 = F_palloc0(m, int32(40))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				v89 = F_pq_getmsgint64(m, v8)
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v87))) = v89
					v92 = F_pq_getmsgint64(m, v8)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v92
						v96 = F_pq_getmsgint(m, v8, int32(4))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = v96
							v100 = F_pq_getmsgint(m, v8, int32(4))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v87)+20)) = v100
								v103 = F_pq_getmsgint64(m, v8)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v87)+24)) = v103
									v106 = F_pq_getmsgint64(m, v8)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v87)+32)) = v106
										F_pq_getmsgend(m, v8)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											m.G0 = v8 + int32(16)
											return v87
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
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66371), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(515856), int32(4172), int32(356457))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
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
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(v14)+24))
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v114 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
	v115 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	m.G0 = v14 + int32(32)
	v119 = v110 + v112
	v120 = v113 + v115
	v124 = int64(63)
	v127 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v119) < base.Ui64(v112))) + (v111 + v110>>(uint(v124)%64))
	v133 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v120) < base.Ui64(v115))) + (v114 + v113>>(uint(v124)%64))
	v135 = base.B2i32(v127 == v133)
	if v127 == v133 {
		v136 = base.B2i32(base.Ui64(v120) < base.Ui64(v119))
	} else {
		v136 = base.B2i32(v133 < v127)
	}
	if v127 == v133 {
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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	v3 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l1)))
	v9 = int64(9223372036854775807)
	v10 = base.I64_reinterpret_f64(v3) & v9
	v13 = base.I64_reinterpret_f64(v4) & v9
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v13) {
		v22 = base.B2i32(base.Ui64(v10) < base.Ui64(int64(9218868437227405313)))
		v30 = int32(0) - v22&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
	} else {
		v18 = int32(1)
		if base.F64_gt(v3, v4) != 0 {
			v30 = v18
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v10) {
				v30 = v18
			} else {
				v22 = v18
				v30 = int32(0) - v22&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v13))|base.F64_lt(v3, v4))
			}
		}
	}
	return v30
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
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
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L52
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L48
	}
L5:
	;
	return v10
L6:
	;
	v54 = base.I64_div_s(v18, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v18+int64(86399999999)) {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v39 = base.I32_div_s(v16, int32(30))
	v40 = v14 + v39
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v40
	v44 = v39*int32(-30) + v16
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v44
	if base.B2i32(v39 < int32(0)) != base.B2i32(v40 < v14) {
		goto L4
	} else {
		goto L23
	}
L8:
	;
	if v18 < int64(0) {
		goto L7
	} else {
		goto L22
	}
L9:
	;
	if int64(0) < v18 {
		goto L7
	} else {
		goto L21
	}
L10:
	;
	if int32(0) < v16 {
		goto L9
	} else {
		goto L19
	}
L11:
	;
	if v14 != int32(-2147483648) {
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
		goto L17
	}
L14:
	;
	if v16 != int32(-2147483648) {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	if v18 == int64(-9223372036854775807-1) {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v18 == int64(9223372036854775807) {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	if v16 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v50 = v16
	v52 = v14
	goto L6
L21:
	;
	v50 = v16
	v52 = v14
	goto L6
L22:
	;
	v50 = v16
	v52 = v14
	goto L6
L23:
	;
	v50 = v44
	v52 = v40
	goto L6
L24:
	;
	v61 = v54*int64(-86400000000) + v18
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v61
	v63 = v61
	goto L26
L25:
	;
	v63 = v18
	goto L26
L26:
	;
	v65 = v50 + base.I32_wrap_i64(v54)
	v67 = base.I32_div_s(v65, int32(30))
	v68 = v52 + v67
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v68
	v72 = v67*int32(-30) + v65
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v72
	if base.B2i32(v67 < int32(0)) != base.B2i32(v68 < v52) {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if int32(0) < v68 {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v130
	goto L5
L29:
	;
	if v63 <= int64(0) {
		goto L5
	} else {
		goto L47
	}
L30:
	;
	if int32(0) <= v100 {
		goto L5
	} else {
		goto L46
	}
L31:
	;
	if int64(0) <= v63 {
		goto L5
	} else {
		goto L45
	}
L32:
	;
	if v100 <= int32(0) {
		goto L30
	} else {
		goto L44
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v94 + v68
	v98 = v72 + v95
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v98
	v100 = v98
	goto L32
L34:
	;
	v80 = int32(-1)
	v81 = int32(30)
	if v72 < int32(0) {
		v94 = v80
		v95 = v81
		goto L33
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if int32(0) <= v68 {
		v100 = v72
		goto L32
	} else {
		goto L40
	}
L37:
	;
	if v72 != 0 {
		v105 = v72
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v63 < int64(0) {
		v94 = v80
		v95 = v81
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L5
L40:
	;
	v88 = int32(1)
	v89 = int32(-30)
	if int32(0) < v72 {
		v94 = v88
		v95 = v89
		goto L33
	} else {
		goto L41
	}
L41:
	;
	if v72 != 0 {
		v117 = v72
		goto L29
	} else {
		goto L42
	}
L42:
	;
	if v63 <= int64(0) {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v94 = v88
	v95 = v89
	goto L33
L44:
	;
	v105 = v100
	goto L31
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v63 + int64(86400000000)
	v130 = v105 - int32(1)
	goto L28
L46:
	;
	v117 = v100
	goto L29
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v10))) = v63 - int64(86400000000)
	v130 = v117 + int32(1)
	goto L28
L48:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(418557), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(515856), int32(2964), int32(321337))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
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
	F_errcode(m, int32(134217858))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errmsg(m, int32(418557), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(515856), int32(2981), int32(321337))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
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
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 float64
	_ = v168
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 float64
	_ = v196
	var v200 float64
	_ = v200
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 float64
	_ = v212
	var v214 float64
	_ = v214
	var v218 float64
	_ = v218
	var v225 float64
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v242 float64
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v252 int64
	_ = v252
	var v258 float64
	_ = v258
	var v275 int64
	_ = v275
	var v277 int64
	_ = v277
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
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
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L75
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
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = v59
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
	if base.F64_lt(v145, float64(2.147483648e+09)) == int32(0) {
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
	if base.F64_ge(v145, float64(-2.147483648e+09)) == int32(0) {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v145)&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if base.F64_lt(base.F64_abs(v145), float64(2.147483648e+09)) != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v168 = base.F64_mul(v20, base.F64_convert_i32_s(v166))
	if base.F64_lt(v168, float64(2.147483648e+09)) == int32(0) {
		goto L4
	} else {
		goto L39
	}
L36:
	;
	v162 = base.I32_trunc_f64_s(v145)
	v164 = v162
	goto L35
L37:
	;
	goto L38
L38:
	;
	v164 = int32(-2147483648)
	goto L35
L39:
	;
	if base.F64_ge(v168, float64(-2.147483648e+09)) == int32(0) {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v168)&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	if base.F64_lt(base.F64_abs(v168), float64(2.147483648e+09)) != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v196 = float64(1e+06)
	v200 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_mul(base.F64_convert_i32_s(v23), v20), base.F64_convert_i32_s(v164)), float64(30)), v196)), v196)
	if base.F64_lt(base.F64_abs(v200), float64(2.147483648e+09)) != 0 {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	v187 = base.I32_trunc_f64_s(v168)
	v189 = v187
	goto L42
L44:
	;
	goto L45
L45:
	;
	v189 = int32(-2147483648)
	goto L42
L46:
	;
	v246 = v244 + v209
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v246
	if base.B2i32(v209 < int32(0))^base.B2i32(v246 < v244) != 0 {
		goto L4
	} else {
		goto L59
	}
L47:
	;
	v212 = float64(86400)
	v214 = float64(1e+06)
	v218 = base.F64_div(base.F64_nearest(base.F64_mul(base.F64_mul(base.F64_sub(base.F64_add(v200, base.F64_sub(base.F64_mul(base.F64_convert_i32_s(v22), v20), base.F64_convert_i32_s(v189))), base.F64_convert_i32_s(v209)), v212), v214)), v214)
	if base.F64_ge(base.F64_abs(v218), v212) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	v207 = base.I32_trunc_f64_s(v200)
	v209 = v207
	goto L47
L49:
	;
	goto L50
L50:
	;
	v209 = int32(-2147483648)
	goto L47
L51:
	;
	v242 = v218
	v244 = v189
	goto L46
L52:
	;
	goto L53
L53:
	;
	v225 = base.F64_div(v218, float64(86400))
	if base.F64_lt(base.F64_abs(v225), float64(2.147483648e+09)) != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v232 = v231 + v189
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v232
	if base.B2i32(v231 < int32(0))^base.B2i32(v232 < v189) != 0 {
		goto L4
	} else {
		goto L58
	}
L55:
	;
	v229 = base.I32_trunc_f64_s(v225)
	v231 = v229
	goto L54
L56:
	;
	goto L57
L57:
	;
	v231 = int32(-2147483648)
	goto L54
L58:
	;
	v242 = base.F64_sub(v218, base.F64_convert_i32_s(v231*int32(86400)))
	v244 = v232
	goto L46
L59:
	;
	v252 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v258 = base.F64_nearest(base.F64_add(base.F64_mul(base.F64_convert_i64_s(v252), v20), base.F64_mul(v242, float64(1e+06))))
	if base.F64_lt(v258, float64(9.223372036854776e+18)) == int32(0) {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	if base.F64_ge(v258, float64(-9.223372036854776e+18)) == int32(0) {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v258)&int64(9223372036854775807)) {
		goto L4
	} else {
		goto L62
	}
L62:
	;
	if base.F64_lt(base.F64_abs(v258), float64(9.223372036854776e+18)) != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v277
	if v164 != int32(2147483647) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v275 = base.I64_trunc_f64_s(v258)
	v277 = v275
	goto L63
L65:
	;
	goto L66
L66:
	;
	v277 = int64(-9223372036854775807 - 1)
	goto L63
L67:
	;
	if v164 != int32(-2147483648) {
		goto L3
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	if v246 != int32(2147483647) {
		goto L3
	} else {
		goto L73
	}
L70:
	;
	if v246 != int32(-2147483648) {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	if v277 == int64(-9223372036854775807-1) {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	goto L3
L73:
	;
	if v277 != int64(9223372036854775807) {
		goto L3
	} else {
		goto L74
	}
L74:
	;
	goto L4
L75:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(418557), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(515856), int32(3740), int32(311844))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int64
	_ = v100
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v118 int64
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		if v22 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				if v52 != int32(2147483647) {
					if v52 != int32(-2147483648) {
						v74 = v14 + int32(96)
						v75 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
						*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
						v79 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
						v87 = v14 + int32(24)
						v89 = v14 + int32(8)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
						v91 = int32(12)
						v92 = base.I32_div_s(v90, v91)
						*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
						*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
						v102 = base.I64_div_s(v100, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
						v106 = v100 + v102*int64(-3600000000)
						v108 = base.I64_div_s(v106, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
						v112 = v108*int64(-60000000) + v106
						v114 = base.I64_div_s(v112, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
						v118 = v114*int64(4293967296) + v112
						*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
						*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							if v143 == int32(0) {
								v147 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
								v150 = v79
							} else {
								v150 = v143
							}
							m.G0 = v14 + int32(112)
							return v150
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						if v57 != int32(-2147483648) {
							v74 = v14 + int32(96)
							v75 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
							*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
							v79 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
							v87 = v14 + int32(24)
							v89 = v14 + int32(8)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
							v91 = int32(12)
							v92 = base.I32_div_s(v90, v91)
							*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
							v102 = base.I64_div_s(v100, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
							v106 = v100 + v102*int64(-3600000000)
							v108 = base.I64_div_s(v106, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
							v112 = v108*int64(-60000000) + v106
							v114 = base.I64_div_s(v112, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
							v118 = v114*int64(4293967296) + v112
							*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
							v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
							*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
							*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								if v143 == int32(0) {
									v147 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
									v150 = v79
								} else {
									v150 = v143
								}
								m.G0 = v14 + int32(112)
								return v150
							}
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							if v60 == int64(-9223372036854775807-1) {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v150 = int32(0)
								m.G0 = v14 + int32(112)
								return v150
							} else {
								v74 = v14 + int32(96)
								v75 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
								*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
								v79 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
								v87 = v14 + int32(24)
								v89 = v14 + int32(8)
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
								v91 = int32(12)
								v92 = base.I32_div_s(v90, v91)
								*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
								*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
								v102 = base.I64_div_s(v100, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
								v106 = v100 + v102*int64(-3600000000)
								v108 = base.I64_div_s(v106, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
								v112 = v108*int64(-60000000) + v106
								v114 = base.I64_div_s(v112, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
								v118 = v114*int64(4293967296) + v112
								*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
								*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
								*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									if v143 == int32(0) {
										v147 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
										v150 = v79
									} else {
										v150 = v143
									}
									m.G0 = v14 + int32(112)
									return v150
								}
							}
						}
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
					if v63 != int32(2147483647) {
						v74 = v14 + int32(96)
						v75 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
						*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
						v79 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
						v87 = v14 + int32(24)
						v89 = v14 + int32(8)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
						v91 = int32(12)
						v92 = base.I32_div_s(v90, v91)
						*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
						*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
						v102 = base.I64_div_s(v100, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
						v106 = v100 + v102*int64(-3600000000)
						v108 = base.I64_div_s(v106, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
						v112 = v108*int64(-60000000) + v106
						v114 = base.I64_div_s(v112, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
						v118 = v114*int64(4293967296) + v112
						*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
						*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							if v143 == int32(0) {
								v147 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
								v150 = v79
							} else {
								v150 = v143
							}
							m.G0 = v14 + int32(112)
							return v150
						}
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						if v66 != int64(9223372036854775807) {
							v74 = v14 + int32(96)
							v75 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
							*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
							v79 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
							v87 = v14 + int32(24)
							v89 = v14 + int32(8)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
							v91 = int32(12)
							v92 = base.I32_div_s(v90, v91)
							*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
							v102 = base.I64_div_s(v100, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
							v106 = v100 + v102*int64(-3600000000)
							v108 = base.I64_div_s(v106, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
							v112 = v108*int64(-60000000) + v106
							v114 = base.I64_div_s(v112, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
							v118 = v114*int64(4293967296) + v112
							*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
							v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
							*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
							*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								if v143 == int32(0) {
									v147 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
									v150 = v79
								} else {
									v150 = v143
								}
								m.G0 = v14 + int32(112)
								return v150
							}
						} else {
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
							v150 = int32(0)
							m.G0 = v14 + int32(112)
							return v150
						}
					}
				}
			} else {
				v48 = base.B2i32(v25 == int32(18)) << (uint(int32(4)) % 32)
				if v48 == int32(0) {
					v70 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
					v150 = int32(0)
					m.G0 = v14 + int32(112)
					return v150
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
					if v52 != int32(2147483647) {
						if v52 != int32(-2147483648) {
							v74 = v14 + int32(96)
							v75 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
							*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
							v79 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
							v87 = v14 + int32(24)
							v89 = v14 + int32(8)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
							v91 = int32(12)
							v92 = base.I32_div_s(v90, v91)
							*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
							v102 = base.I64_div_s(v100, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
							v106 = v100 + v102*int64(-3600000000)
							v108 = base.I64_div_s(v106, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
							v112 = v108*int64(-60000000) + v106
							v114 = base.I64_div_s(v112, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
							v118 = v114*int64(4293967296) + v112
							*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
							v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
							*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
							*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								if v143 == int32(0) {
									v147 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
									v150 = v79
								} else {
									v150 = v143
								}
								m.G0 = v14 + int32(112)
								return v150
							}
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
							if v57 != int32(-2147483648) {
								v74 = v14 + int32(96)
								v75 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
								*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
								v79 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
								v87 = v14 + int32(24)
								v89 = v14 + int32(8)
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
								v91 = int32(12)
								v92 = base.I32_div_s(v90, v91)
								*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
								*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
								v102 = base.I64_div_s(v100, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
								v106 = v100 + v102*int64(-3600000000)
								v108 = base.I64_div_s(v106, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
								v112 = v108*int64(-60000000) + v106
								v114 = base.I64_div_s(v112, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
								v118 = v114*int64(4293967296) + v112
								*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
								*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
								*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									if v143 == int32(0) {
										v147 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
										v150 = v79
									} else {
										v150 = v143
									}
									m.G0 = v14 + int32(112)
									return v150
								}
							} else {
								v60 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								if v60 == int64(-9223372036854775807-1) {
									v70 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
									v150 = int32(0)
									m.G0 = v14 + int32(112)
									return v150
								} else {
									v74 = v14 + int32(96)
									v75 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
									*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
									v79 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
									v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
									v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
									*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
									v87 = v14 + int32(24)
									v89 = v14 + int32(8)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
									v91 = int32(12)
									v92 = base.I32_div_s(v90, v91)
									*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
									*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
									v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
									v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
									v102 = base.I64_div_s(v100, int64(3600000000))
									*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
									v106 = v100 + v102*int64(-3600000000)
									v108 = base.I64_div_s(v106, int64(60000000))
									*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
									v112 = v108*int64(-60000000) + v106
									v114 = base.I64_div_s(v112, int64(1000000))
									*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
									v118 = v114*int64(4293967296) + v112
									*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
									v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
									*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
									v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
									*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
									*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
									*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
									*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
									v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return int32(0)
									} else {
										if v143 == int32(0) {
											v147 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
											v150 = v79
										} else {
											v150 = v143
										}
										m.G0 = v14 + int32(112)
										return v150
									}
								}
							}
						}
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						if v63 != int32(2147483647) {
							v74 = v14 + int32(96)
							v75 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
							*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
							v79 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
							v87 = v14 + int32(24)
							v89 = v14 + int32(8)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
							v91 = int32(12)
							v92 = base.I32_div_s(v90, v91)
							*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
							v102 = base.I64_div_s(v100, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
							v106 = v100 + v102*int64(-3600000000)
							v108 = base.I64_div_s(v106, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
							v112 = v108*int64(-60000000) + v106
							v114 = base.I64_div_s(v112, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
							v118 = v114*int64(4293967296) + v112
							*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
							v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
							*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
							*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								if v143 == int32(0) {
									v147 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
									v150 = v79
								} else {
									v150 = v143
								}
								m.G0 = v14 + int32(112)
								return v150
							}
						} else {
							v66 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							if v66 != int64(9223372036854775807) {
								v74 = v14 + int32(96)
								v75 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
								*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
								v79 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
								v87 = v14 + int32(24)
								v89 = v14 + int32(8)
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
								v91 = int32(12)
								v92 = base.I32_div_s(v90, v91)
								*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
								*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
								v102 = base.I64_div_s(v100, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
								v106 = v100 + v102*int64(-3600000000)
								v108 = base.I64_div_s(v106, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
								v112 = v108*int64(-60000000) + v106
								v114 = base.I64_div_s(v112, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
								v118 = v114*int64(4293967296) + v112
								*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
								*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
								*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									if v143 == int32(0) {
										v147 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
										v150 = v79
									} else {
										v150 = v143
									}
									m.G0 = v14 + int32(112)
									return v150
								}
							} else {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v150 = int32(0)
								m.G0 = v14 + int32(112)
								return v150
							}
						}
					}
				}
			}
		} else {
			v36 = int32(1)
			if v22&v36 != 0 {
				v48 = int32(base.Ui32(v22)>>(uint(v36)%32)) - v36
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v48 = int32(base.Ui32(v42)>>(uint(int32(2))%32)) - int32(4)
			}
			if v48 == int32(0) {
				v70 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
				v150 = int32(0)
				m.G0 = v14 + int32(112)
				return v150
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
				if v52 != int32(2147483647) {
					if v52 != int32(-2147483648) {
						v74 = v14 + int32(96)
						v75 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
						*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
						v79 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
						v87 = v14 + int32(24)
						v89 = v14 + int32(8)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
						v91 = int32(12)
						v92 = base.I32_div_s(v90, v91)
						*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
						*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
						v102 = base.I64_div_s(v100, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
						v106 = v100 + v102*int64(-3600000000)
						v108 = base.I64_div_s(v106, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
						v112 = v108*int64(-60000000) + v106
						v114 = base.I64_div_s(v112, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
						v118 = v114*int64(4293967296) + v112
						*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
						*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							if v143 == int32(0) {
								v147 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
								v150 = v79
							} else {
								v150 = v143
							}
							m.G0 = v14 + int32(112)
							return v150
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						if v57 != int32(-2147483648) {
							v74 = v14 + int32(96)
							v75 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
							*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
							v79 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
							v87 = v14 + int32(24)
							v89 = v14 + int32(8)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
							v91 = int32(12)
							v92 = base.I32_div_s(v90, v91)
							*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
							v102 = base.I64_div_s(v100, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
							v106 = v100 + v102*int64(-3600000000)
							v108 = base.I64_div_s(v106, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
							v112 = v108*int64(-60000000) + v106
							v114 = base.I64_div_s(v112, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
							v118 = v114*int64(4293967296) + v112
							*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
							v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
							*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
							*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								if v143 == int32(0) {
									v147 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
									v150 = v79
								} else {
									v150 = v143
								}
								m.G0 = v14 + int32(112)
								return v150
							}
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							if v60 == int64(-9223372036854775807-1) {
								v70 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
								v150 = int32(0)
								m.G0 = v14 + int32(112)
								return v150
							} else {
								v74 = v14 + int32(96)
								v75 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
								*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
								v79 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
								v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
								v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
								v87 = v14 + int32(24)
								v89 = v14 + int32(8)
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
								v91 = int32(12)
								v92 = base.I32_div_s(v90, v91)
								*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
								*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
								v102 = base.I64_div_s(v100, int64(3600000000))
								*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
								v106 = v100 + v102*int64(-3600000000)
								v108 = base.I64_div_s(v106, int64(60000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
								v112 = v108*int64(-60000000) + v106
								v114 = base.I64_div_s(v112, int64(1000000))
								*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
								v118 = v114*int64(4293967296) + v112
								*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
								v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
								v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
								v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
								*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
								*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
								v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return int32(0)
								} else {
									if v143 == int32(0) {
										v147 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
										v150 = v79
									} else {
										v150 = v143
									}
									m.G0 = v14 + int32(112)
									return v150
								}
							}
						}
					}
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
					if v63 != int32(2147483647) {
						v74 = v14 + int32(96)
						v75 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
						*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
						v79 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
						v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
						v87 = v14 + int32(24)
						v89 = v14 + int32(8)
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
						v91 = int32(12)
						v92 = base.I32_div_s(v90, v91)
						*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
						*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
						v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
						v102 = base.I64_div_s(v100, int64(3600000000))
						*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
						v106 = v100 + v102*int64(-3600000000)
						v108 = base.I64_div_s(v106, int64(60000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
						v112 = v108*int64(-60000000) + v106
						v114 = base.I64_div_s(v112, int64(1000000))
						*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
						v118 = v114*int64(4293967296) + v112
						*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
						*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
						*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
						*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
						*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
						v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							if v143 == int32(0) {
								v147 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
								v150 = v79
							} else {
								v150 = v143
							}
							m.G0 = v14 + int32(112)
							return v150
						}
					} else {
						v66 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						if v66 != int64(9223372036854775807) {
							v74 = v14 + int32(96)
							v75 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v74))) = v75
							*(*int64)(unsafe.Add(mBase, uint32(v14)+88)) = v75
							v79 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v14)+108)) = v79
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v82
							v84 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v84
							v87 = v14 + int32(24)
							v89 = v14 + int32(8)
							v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
							v91 = int32(12)
							v92 = base.I32_div_s(v90, v91)
							*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v92
							*(*int32)(unsafe.Add(mBase, uint32(v87)+28)) = v90 - v92*v91
							v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
							v102 = base.I64_div_s(v100, int64(3600000000))
							*(*int64)(unsafe.Add(mBase, uint32(v87)+16)) = v102
							v106 = v100 + v102*int64(-3600000000)
							v108 = base.I64_div_s(v106, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+8)) = uint32(v108)
							v112 = v108*int64(-60000000) + v106
							v114 = base.I64_div_s(v112, int64(1000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v87)+4)) = uint32(v114)
							v118 = v114*int64(4293967296) + v112
							*(*uint32)(unsafe.Add(mBase, uint32(v87))) = uint32(v118)
							v120 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
							*(*int32)(unsafe.Add(mBase, uint32(v74))) = v120 + (v121+v122*v91)*int32(30)
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+104)) = v130
							v132 = *(*int64)(unsafe.Add(mBase, uint32(v14)+28))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+64)) = v132
							v134 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v14)+72)) = v134
							*(*int32)(unsafe.Add(mBase, uint32(v14)+80)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v14)+84)) = v121
							*(*int32)(unsafe.Add(mBase, uint32(v14)+88)) = v122
							v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v143 = F_datetime_to_char_body(m, v14-int32(-64), v18, int32(1), v142)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int32(0)
							} else {
								if v143 == int32(0) {
									v147 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
									v150 = v79
								} else {
									v150 = v143
								}
								m.G0 = v14 + int32(112)
								return v150
							}
						} else {
							v70 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v70)
							v150 = int32(0)
							m.G0 = v14 + int32(112)
							return v150
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
	var v152 int64
	_ = v152
	var v157 int32
	_ = v157
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v272 int32
	_ = v272
	var v276 int64
	_ = v276
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int64
	_ = v321
	var v322 int64
	_ = v322
	var v323 int64
	_ = v323
	var v328 int64
	_ = v328
	var v334 int32
	_ = v334
	var v337 int64
	_ = v337
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v349 int64
	_ = v349
	var v350 int64
	_ = v350
	var v356 int64
	_ = v356
	var v357 int64
	_ = v357
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
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
	v225 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v227 = base.I64_div_s(v225, int64(3600000000))
	v230 = v227*int64(-3600000000) + v225
	v232 = base.I64_div_s(v230, int64(60000000))
	v235 = v232*int64(-60000000) + v230
	v237 = base.I64_div_s(v235, int64(1000000))
	v241 = base.I32_wrap_i64(v237*int64(4293967296) + v235)
	v242 = int32(12)
	v243 = base.I32_div_s(v136, v242)
	v246 = v136 - v243*v242
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	switch v247 - int32(18) {
	case 0:
		goto L75
	case 1:
		goto L72
	case 2:
		goto L76
	case 3:
		goto L77
	default:
		goto L73
	case 5:
		v269 = v243
		v270 = v246
		goto L78
	case 6:
		v263 = v243
		v264 = v246
		goto L79
	case 7:
		v261 = v243
		goto L80
	case 8:
		v257 = v243
		goto L81
	case 9:
		v253 = v243
		goto L82
	case 10:
		goto L83
	case 11:
		goto L74
	case 12:
		v317 = v243
		v318 = v246
		v319 = v224
		v320 = v241
		v321 = v232
		v322 = v227
		v323 = v237
		goto L70
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
	v62 = F_downcase_truncate_identifier(m, v32, v60, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v35 = int32(4)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v37&int32(254) == int32(2) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v50 = int32(1)
	if v31 != 0 {
		v60 = int32(base.Ui32(v29)>>(uint(v50)%32)) - v50
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v46 = v35
	goto L15
L14:
	;
	v46 = base.B2i32(v37 == int32(18)) << (uint(v35) % 32)
	goto L15
L15:
	;
	if v37 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v49 = v35
	goto L18
L17:
	;
	v49 = v46
	goto L18
L18:
	;
	v60 = v49
	goto L9
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L20:
	;
	v65 = v14 + int32(44)
	v72 = *(*int32)(unsafe.Add(mBase, _consts[1256]))
	if v72 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	if v133 == int32(17) {
		goto L40
	} else {
		goto L41
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1256])) = v116
	v123 = int32(*(*int8)(unsafe.Add(mBase, uint32(v116)+11)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v124
	v133 = v123
	goto L21
L23:
	;
	v74 = F_strncmp(m, v62, v72, int32(10))
	mBase = m.M
	if v74 == int32(0) {
		v116 = v72
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v77 = int32(*(*int8)(unsafe.Add(mBase, uint32(v62))))
	v83 = int32(1688224)
	v85 = int32(1689184)
	goto L27
L26:
	;
	goto L25
L27:
	;
	v92 = v83 + (v85-v83)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v93 = int32(*(*int8)(unsafe.Add(mBase, uint32(v92))))
	v94 = v77 - v93
	if v94 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = int32(0)
	v133 = int32(31)
	goto L21
L29:
	;
	v98 = F_strncmp(m, v62, v92, int32(10))
	mBase = m.M
	if v98 == int32(0) {
		v116 = v92
		goto L22
	} else {
		goto L32
	}
L30:
	;
	v101 = v94
	goto L31
L31:
	;
	v105 = base.B2i32(v101 < int32(0))
	if v101 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v101 = v98
	goto L31
L33:
	;
	v106 = v92 - int32(16)
	goto L35
L34:
	;
	v106 = v85
	goto L35
L35:
	;
	if v101 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v109 = v83
	goto L38
L37:
	;
	v109 = v92 + int32(16)
	goto L38
L38:
	;
	if base.Ui32(v109) <= base.Ui32(v106) {
		v83 = v109
		v85 = v106
		goto L27
	} else {
		goto L39
	}
L39:
	;
	goto L28
L40:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v136 != int32(-2147483648) {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	goto L42
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L65
	}
L43:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if base.B2i32(base.Ui32(int32(8)) <= base.Ui32(v157-int32(23)))&base.B2i32(base.Ui32(int32(3)) < base.Ui32(v157-int32(18))) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v149 != int32(2147483647) {
		v224 = v149
		goto L5
	} else {
		goto L51
	}
L45:
	;
	if v136 == int32(2147483647) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v142 != int32(-2147483648) {
		v224 = v142
		goto L5
	} else {
		goto L49
	}
L48:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v224 = v141
	goto L5
L49:
	;
	v145 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	if v145 == int64(-9223372036854775807-1) {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	v224 = int32(-2147483648)
	goto L5
L51:
	;
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	if v152 == int64(9223372036854775807) {
		goto L43
	} else {
		goto L52
	}
L52:
	;
	v224 = int32(2147483647)
	goto L5
L53:
	;
	v169 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v169
	v171 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+8)) = v171
	goto L4
L54:
	;
	goto L55
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v181 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v62
	F_errmsg(m, int32(199023), v14+int32(16))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v190 == int32(22) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	F_errdetail(m, int32(617402), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_errfinish(m, int32(515856), int32(5181), int32(508396))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v210 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v210
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v62
	F_errmsg(m, int32(198986), v14+int32(32))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(515856), int32(5246), int32(508396))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v328 = base.I64_extend_i32_s(v318) + base.I64_extend_i32_s(v317)*int64(12)
	if base.Ui64(v328-int64(2147483648)) < base.Ui64(int64(-4294967296)) {
		goto L93
	} else {
		goto L94
	}
L71:
	;
	v317 = v243
	v318 = v246
	v319 = v224
	v320 = v315
	v321 = v313
	v322 = v314
	v323 = int64(0)
	goto L70
L72:
	;
	v313 = v232
	v314 = v227
	v315 = int32(0)
	goto L71
L73:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L84
	}
L74:
	;
	v283 = base.I32_rem_s(v241, int32(1000))
	v317 = v243
	v318 = v246
	v319 = v224
	v320 = v241 - v283
	v321 = v232
	v322 = v227
	v323 = v237
	goto L70
L75:
	;
	v317 = v243
	v318 = v246
	v319 = v224
	v320 = int32(0)
	v321 = v232
	v322 = v227
	v323 = v237
	goto L70
L76:
	;
	v313 = int64(0)
	v314 = v227
	v315 = int32(0)
	goto L71
L77:
	;
	v276 = int64(0)
	v313 = v276
	v314 = v276
	v315 = int32(0)
	goto L71
L78:
	;
	v271 = int64(0)
	v272 = int32(0)
	v317 = v269
	v318 = v270
	v319 = v272
	v320 = v272
	v321 = v271
	v322 = v271
	v323 = v271
	goto L70
L79:
	;
	v267 = base.I32_rem_s(base.I32_extend8_s(v264), int32(3))
	v269 = v263
	v270 = v264 - v267
	goto L78
L80:
	;
	v263 = v261
	v264 = int32(0)
	goto L79
L81:
	;
	v259 = base.I32_rem_s(v257, int32(10))
	v261 = v257 - v259
	goto L80
L82:
	;
	v255 = base.I32_rem_s(v253, int32(100))
	v257 = v253 - v255
	goto L81
L83:
	;
	v251 = base.I32_rem_s(v243, int32(1000))
	v253 = v243 - v251
	goto L82
L84:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v293 = F_format_type_be(m, int32(1186))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v62
	F_errmsg(m, int32(199023), v14)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	if v300 == int32(22) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_errdetail(m, int32(617402), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	F_errfinish(m, int32(515856), int32(5233), int32(508396))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L92
	}
L91:
	;
	goto L90
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L106
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v319
	v334 = base.I32_wrap_i64(v328)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v334
	v337 = v322 * int64(3600000000)
	v340 = base.I64_extend32_s(v321) * int64(60000000)
	v341 = v337 + v340
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v341
	if base.B2i32(v340 < int64(0))^base.B2i32(v341 < v337) != 0 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v349 = base.I64_extend32_s(v323) * int64(1000000)
	v350 = v341 + v349
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v350
	if base.B2i32(v349 < int64(0))^base.B2i32(v350 < v341) != 0 {
		goto L93
	} else {
		goto L96
	}
L96:
	;
	v356 = base.I64_extend_i32_s(v320)
	v357 = v350 + v356
	*(*int64)(unsafe.Add(mBase, uint32(v23))) = v357
	if base.B2i32(v356 < int64(0))^base.B2i32(v357 < v350) != 0 {
		goto L93
	} else {
		goto L97
	}
L97:
	;
	if v334 != int32(2147483647) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	if v334 != int32(-2147483648) {
		goto L4
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v319 != int32(2147483647) {
		goto L4
	} else {
		goto L104
	}
L101:
	;
	if v319 != int32(-2147483648) {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	if v357 != int64(-9223372036854775807-1) {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	goto L93
L104:
	;
	if v357 != int64(9223372036854775807) {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	goto L93
L106:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errmsg(m, int32(418557), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(515856), int32(5239), int32(508396))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
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
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 != int32(-2147483648) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return
L2:
	;
	v35 = int64(0)
	v36 = v35 - v34
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v36
	if base.B2i32(v36 < v35)^base.B2i32(v35 < v34) != 0 {
		goto L13
	} else {
		goto L14
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v28
	v31 = v7 ^ int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v31
	goto L1
L4:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v21 != int32(2147483647) {
		v34 = v20
		goto L2
	} else {
		goto L11
	}
L5:
	;
	if v7 == int32(2147483647) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v14 != int32(-2147483648) {
		v34 = v13
		goto L2
	} else {
		goto L9
	}
L8:
	;
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v34 = v12
	goto L2
L9:
	;
	if v13 != int64(-9223372036854775807-1) {
		v34 = v13
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v28 = int64(9223372036854775807)
	goto L3
L11:
	;
	if v20 != int64(9223372036854775807) {
		v34 = v20
		goto L2
	} else {
		goto L12
	}
L12:
	;
	v28 = int64(-9223372036854775807 - 1)
	goto L3
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v43 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v45 = v43 - v44
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v45
	if base.B2i32(v45 < v43)^base.B2i32(v43 < v44) != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v52 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v54 = v52 - v53
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v54
	if base.B2i32(v54 < v52)^base.B2i32(v52 < v53) != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v54 != int32(2147483647) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	if v54 != int32(-2147483648) {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v45 != int32(2147483647) {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	if v45 != int32(-2147483648) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v36 != int64(-9223372036854775807-1) {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L13
L23:
	;
	if v36 != int64(9223372036854775807) {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L13
L25:
	;
	return
L26:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(418557), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(515856), int32(3459), int32(323982))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L25
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
