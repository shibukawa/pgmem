package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v10 = int64(1000000)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v13 = base.I64_extend_i32_s(v8)*v10 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v20 = base.I64_extend_i32_s(v15)*v10 + v19
	if v20 < v13 {
		v30 = int32(1)
	} else {
		if v13 < v20 {
			v30 = int32(-1)
		} else {
			if v15 < v8 {
				v30 = int32(1)
			} else {
				if v8 < v15 {
					v29 = int32(-1)
				} else {
					v29 = int32(0)
				}
				v30 = v29
			}
		}
	}
	return v30
}
func F_timetz_hash_extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = F_DirectFunctionCall2Coll(m, int32(1285), int32(0), v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
		if v13 == int64(0) {
			v21 = int32(-1636608428)
			v59 = v21
			v61 = v21
			v64 = v21
		} else {
			v24 = base.I32_wrap_i64(v13)
			v29 = base.I32_wrap_i64(int64(base.Ui64(v13)>>(uint(int64(32))%64))) ^ int32(-415931063)
			v35 = v24 - v29 - int32(1636608428) ^ base.I32_rotl(v29, int32(6))
			v37 = v24 + int32(1021750440)
			v38 = v29 + v37
			v39 = v35 + v38
			v43 = v37 - v35 ^ base.I32_rotl(v35, int32(8))
			v47 = v38 - v43 ^ base.I32_rotl(v43, int32(16))
			v51 = v39 - v47 ^ base.I32_rotl(v47, int32(19))
			v52 = v43 + v39
			v53 = v47 + v52
			v59 = v51 + v53
			v61 = v53
			v64 = v52 - v51 ^ base.I32_rotl(v51, int32(4))
		}
		v66 = int32(14)
		v68 = v59 ^ v64 - base.I32_rotl(v59, v66)
		v73 = v68 ^ (v12 + v61) - base.I32_rotl(v68, int32(11))
		v77 = v73 ^ v59 - base.I32_rotl(v73, int32(25))
		v81 = v77 ^ v68 - base.I32_rotl(v77, int32(16))
		v85 = v81 ^ v73 - base.I32_rotl(v81, int32(4))
		v89 = v85 ^ v77 - base.I32_rotl(v85, v66)
		v99 = F_Int64GetDatum(m, base.I64_extend_i32_u(v89)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v89^v81-base.I32_rotl(v89, int32(24))))
		mBase = m.M
		v100 = m.ExcPending
		if v100 != 0 {
			return int32(0)
		} else {
			v101 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
			v103 = F_Int64GetDatum(m, v11^v101)
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int32(0)
			} else {
				return v103
			}
		}
	}
}
func F_timetz_pl_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v7 != int32(2147483647) {
		if v7 != int32(-2147483648) {
			v43 = F_palloc(m, int32(16))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
				v48 = int64(86400000000)
				v49 = base.I64_rem_s(v45+v46, v48)
				if v49 < int64(0) {
					v54 = v49 + v48
				} else {
					v54 = v49
				}
				*(*int64)(unsafe.Add(mBase, uint32(v43))) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v56
				return v43
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
			if v12 != int32(-2147483648) {
				v43 = F_palloc(m, int32(16))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
					v48 = int64(86400000000)
					v49 = base.I64_rem_s(v45+v46, v48)
					if v49 < int64(0) {
						v54 = v49 + v48
					} else {
						v54 = v49
					}
					*(*int64)(unsafe.Add(mBase, uint32(v43))) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v56
					return v43
				}
			} else {
				v15 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				if v15 == int64(-9223372036854775807-1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(370020), int32(0))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490962), int32(2670), int32(304493))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
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
					v43 = F_palloc(m, int32(16))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
						v48 = int64(86400000000)
						v49 = base.I64_rem_s(v45+v46, v48)
						if v49 < int64(0) {
							v54 = v49 + v48
						} else {
							v54 = v49
						}
						*(*int64)(unsafe.Add(mBase, uint32(v43))) = v54
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v56
						return v43
					}
				}
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v18 != int32(2147483647) {
			v43 = F_palloc(m, int32(16))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
				v48 = int64(86400000000)
				v49 = base.I64_rem_s(v45+v46, v48)
				if v49 < int64(0) {
					v54 = v49 + v48
				} else {
					v54 = v49
				}
				*(*int64)(unsafe.Add(mBase, uint32(v43))) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v56
				return v43
			}
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
			if v21 != int64(9223372036854775807) {
				v43 = F_palloc(m, int32(16))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					v45 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
					v48 = int64(86400000000)
					v49 = base.I64_rem_s(v45+v46, v48)
					if v49 < int64(0) {
						v54 = v49 + v48
					} else {
						v54 = v49
					}
					*(*int64)(unsafe.Add(mBase, uint32(v43))) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v56
					return v43
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(370020), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490962), int32(2670), int32(304493))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
	}
}
func F_timetz_zone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v8 = m.G0
	v10 = v8 - int32(336)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		F_text_to_cstring_buffer(m, v13, v10-int32(-64), int32(256))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v29 = F_DecodeTimezoneName(m, v10-int32(-64), v10+int32(60), v10+int32(56))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				switch v29 {
				case 0:
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+332)) = int32(0) - v32
					v58 = F_palloc(m, int32(16))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+332))
						*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v62
						v67 = base.I64_extend_i32_s(v61-v62) * int64(1000000)
						v68 = v60 + v67
						v69 = int64(0)
						if v69 < v68 {
							v72 = v68
						} else {
							v72 = v69
						}
						v73 = v72 - v60
						v75 = base.I64_extend_i32_u(base.B2i32(v73 != v67))
						v78 = int64(86400000000)
						v79 = base.I64_div_u_s(v73-(v67|v75), v78)
						v84 = v60 + (v79+v75)*v78 + v67
						v86 = base.I64_rem_u_s(v84, v78)
						if base.Ui64(int64(86399999999)) < base.Ui64(v84) {
							v89 = v86
						} else {
							v89 = v84
						}
						*(*int64)(unsafe.Add(mBase, uint32(v58))) = v89
						m.G0 = v10 + int32(336)
						return v58
					}
				case 1:
					v36 = *(*int64)(unsafe.Add(mBase, _consts[429]))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
					v42 = F_DetermineTimeZoneAbbrevOffsetTS(m, v36, v10-int32(-64), v39, v10+int32(12))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+332)) = v42
						v58 = F_palloc(m, int32(16))
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+332))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v62
							v67 = base.I64_extend_i32_s(v61-v62) * int64(1000000)
							v68 = v60 + v67
							v69 = int64(0)
							if v69 < v68 {
								v72 = v68
							} else {
								v72 = v69
							}
							v73 = v72 - v60
							v75 = base.I64_extend_i32_u(base.B2i32(v73 != v67))
							v78 = int64(86400000000)
							v79 = base.I64_div_u_s(v73-(v67|v75), v78)
							v84 = v60 + (v79+v75)*v78 + v67
							v86 = base.I64_rem_u_s(v84, v78)
							if base.Ui64(int64(86399999999)) < base.Ui64(v84) {
								v89 = v86
							} else {
								v89 = v84
							}
							*(*int64)(unsafe.Add(mBase, uint32(v58))) = v89
							m.G0 = v10 + int32(336)
							return v58
						}
					}
				default:
					v46 = *(*int64)(unsafe.Add(mBase, _consts[429]))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
					v55 = F_timestamp2tm(m, v46, v10+int32(332), v10+int32(12), v10+int32(8), int32(0), v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(396479), int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490962), int32(3166), int32(366671))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
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
							v58 = F_palloc(m, int32(16))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v60 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v10)+332))
								*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v62
								v67 = base.I64_extend_i32_s(v61-v62) * int64(1000000)
								v68 = v60 + v67
								v69 = int64(0)
								if v69 < v68 {
									v72 = v68
								} else {
									v72 = v69
								}
								v73 = v72 - v60
								v75 = base.I64_extend_i32_u(base.B2i32(v73 != v67))
								v78 = int64(86400000000)
								v79 = base.I64_div_u_s(v73-(v67|v75), v78)
								v84 = v60 + (v79+v75)*v78 + v67
								v86 = base.I64_rem_u_s(v84, v78)
								if base.Ui64(int64(86399999999)) < base.Ui64(v84) {
									v89 = v86
								} else {
									v89 = v84
								}
								*(*int64)(unsafe.Add(mBase, uint32(v58))) = v89
								m.G0 = v10 + int32(336)
								return v58
							}
						}
					}
				}
			}
		}
	}
}
