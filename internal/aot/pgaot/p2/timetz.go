package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_cmp(m *base.Module, l0 int32) int64 {
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
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v10 = int64(1000000)
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v13 = base.I64_extend_i32_s(v8)*v10 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
	v20 = base.I64_extend_i32_s(v15)*v10 + v19
	if v20 < v13 {
		v30 = int64(1)
	} else {
		if v13 < v20 {
			v30 = int64(-1)
		} else {
			if v15 < v8 {
				v30 = int64(1)
			} else {
				if v8 < v15 {
					v29 = int64(-1)
				} else {
					v29 = int64(0)
				}
				v30 = v29
			}
		}
	}
	return v30
}
func F_timetz_hash_extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = F_DirectFunctionCall2Coll(m, int32(1398), int32(0), v7, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
		if v8 == int64(0) {
			v20 = int32(-1636608428)
			v59 = v20
			v60 = v20
			v63 = int32(0)
		} else {
			v23 = base.I32_wrap_i64(v8)
			v25 = v23 + int32(1021750440)
			v30 = base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(32))%64))) ^ int32(-415931063)
			v36 = v23 - v30 - int32(1636608428) ^ base.I32_rotl(v30, int32(6))
			v40 = v25 - v36 ^ base.I32_rotl(v36, int32(8))
			v41 = v30 + v25
			v42 = v36 + v41
			v43 = v40 + v42
			v47 = v41 - v40 ^ base.I32_rotl(v40, int32(16))
			v51 = v42 - v47 ^ base.I32_rotl(v47, int32(19))
			v56 = v47 + v43
			v57 = v51 + v56
			v59 = v57
			v60 = v56
			v63 = v43 - v51 ^ base.I32_rotl(v51, int32(4)) ^ v57
		}
		v64 = int32(14)
		v66 = v63 - base.I32_rotl(v59, v64)
		v71 = v66 ^ (v13 + v60) - base.I32_rotl(v66, int32(11))
		v75 = v59 ^ v71 - base.I32_rotl(v71, int32(25))
		v79 = v75 ^ v66 - base.I32_rotl(v75, int32(16))
		v83 = v79 ^ v71 - base.I32_rotl(v79, int32(4))
		v87 = v83 ^ v75 - base.I32_rotl(v83, v64)
		return v9 ^ (base.I64_extend_i32_u(v87)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v87^v79-base.I32_rotl(v87, int32(24))))
	}
}
func F_timetz_pl_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int32
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
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	if v6 != int32(2147483647) {
		if v6 != int32(-2147483648) {
			v42 = F_palloc(m, int32(16))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				v44 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
				v45 = base.I32_wrap_i64(v4)
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
				v48 = int64(86400000000)
				v49 = base.I64_rem_s(v44+v46, v48)
				if v49 < int64(0) {
					v54 = v49 + v48
				} else {
					v54 = v49
				}
				*(*int64)(unsafe.Add(mBase, uint32(v42))) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v56
				return base.I64_extend_i32_u(v42)
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
			if v11 != int32(-2147483648) {
				v42 = F_palloc(m, int32(16))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
					v45 = base.I32_wrap_i64(v4)
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
					v48 = int64(86400000000)
					v49 = base.I64_rem_s(v44+v46, v48)
					if v49 < int64(0) {
						v54 = v49 + v48
					} else {
						v54 = v49
					}
					*(*int64)(unsafe.Add(mBase, uint32(v42))) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v56
					return base.I64_extend_i32_u(v42)
				}
			} else {
				v14 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
				if v14 == int64(-9223372036854775807-1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_timetz_pl_interval_0), int32(0))
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_timetz_pl_interval_1), int32(2727), int32(_a_F_timetz_pl_interval_2))
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v42 = F_palloc(m, int32(16))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
						v45 = base.I32_wrap_i64(v4)
						v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
						v48 = int64(86400000000)
						v49 = base.I64_rem_s(v44+v46, v48)
						if v49 < int64(0) {
							v54 = v49 + v48
						} else {
							v54 = v49
						}
						*(*int64)(unsafe.Add(mBase, uint32(v42))) = v54
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v56
						return base.I64_extend_i32_u(v42)
					}
				}
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		if v17 != int32(2147483647) {
			v42 = F_palloc(m, int32(16))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				v44 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
				v45 = base.I32_wrap_i64(v4)
				v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
				v48 = int64(86400000000)
				v49 = base.I64_rem_s(v44+v46, v48)
				if v49 < int64(0) {
					v54 = v49 + v48
				} else {
					v54 = v49
				}
				*(*int64)(unsafe.Add(mBase, uint32(v42))) = v54
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v56
				return base.I64_extend_i32_u(v42)
			}
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
			if v20 != int64(9223372036854775807) {
				v42 = F_palloc(m, int32(16))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
					v45 = base.I32_wrap_i64(v4)
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v45)))
					v48 = int64(86400000000)
					v49 = base.I64_rem_s(v44+v46, v48)
					if v49 < int64(0) {
						v54 = v49 + v48
					} else {
						v54 = v49
					}
					*(*int64)(unsafe.Add(mBase, uint32(v42))) = v54
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = v56
					return base.I64_extend_i32_u(v42)
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_timetz_pl_interval_0), int32(0))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timetz_pl_interval_1), int32(2727), int32(_a_F_timetz_pl_interval_2))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
		}
	}
}
func F_timetz_zone(m *base.Module, l0 int32) int64 {
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
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int64
	_ = v71
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v88 int64
	_ = v88
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
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = v10 - int32(-64)
		F_text_to_cstring_buffer(m, v13, v19, int32(256))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int64(0)
		} else {
			v27 = F_DecodeTimezoneName(m, v19, v10+int32(60), v10+int32(56))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				switch v27 {
				case 0:
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+332)) = int32(0) - v30
					v56 = F_palloc(m, int32(16))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int64(0)
					} else {
						v58 = base.I32_wrap_i64(v17)
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+332))
						*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v61
						v66 = base.I64_extend_i32_s(v60-v61) * int64(1000000)
						v67 = v59 + v66
						v68 = int64(0)
						if v68 < v67 {
							v71 = v67
						} else {
							v71 = v68
						}
						v72 = v71 - v59
						v74 = base.I64_extend_i32_u(base.B2i32(v72 != v66))
						v77 = int64(86400000000)
						v78 = base.I64_div_u_s(v72-(v66|v74), v77)
						v83 = v59 + (v78+v74)*v77 + v66
						v85 = base.I64_rem_u_s(v83, v77)
						if base.Ui64(int64(86399999999)) < base.Ui64(v83) {
							v88 = v85
						} else {
							v88 = v83
						}
						*(*int64)(unsafe.Add(mBase, uint32(v56))) = v88
						m.G0 = v10 + int32(336)
						return base.I64_extend_i32_u(v56)
					}
				case 1:
					v34 = *(*int64)(unsafe.Add(mBase, _c_F_timetz_zone[0]))
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
					v40 = F_DetermineTimeZoneAbbrevOffsetTS(m, v34, v10-int32(-64), v37, v10+int32(12))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+332)) = v40
						v56 = F_palloc(m, int32(16))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							v58 = base.I32_wrap_i64(v17)
							v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+332))
							*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v61
							v66 = base.I64_extend_i32_s(v60-v61) * int64(1000000)
							v67 = v59 + v66
							v68 = int64(0)
							if v68 < v67 {
								v71 = v67
							} else {
								v71 = v68
							}
							v72 = v71 - v59
							v74 = base.I64_extend_i32_u(base.B2i32(v72 != v66))
							v77 = int64(86400000000)
							v78 = base.I64_div_u_s(v72-(v66|v74), v77)
							v83 = v59 + (v78+v74)*v77 + v66
							v85 = base.I64_rem_u_s(v83, v77)
							if base.Ui64(int64(86399999999)) < base.Ui64(v83) {
								v88 = v85
							} else {
								v88 = v83
							}
							*(*int64)(unsafe.Add(mBase, uint32(v56))) = v88
							m.G0 = v10 + int32(336)
							return base.I64_extend_i32_u(v56)
						}
					}
				default:
					v44 = *(*int64)(unsafe.Add(mBase, _c_F_timetz_zone[0]))
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
					v53 = F_timestamp2tm(m, v44, v10+int32(332), v10+int32(12), v10+int32(8), int32(0), v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						if v53 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timetz_zone_0), int32(0))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timetz_zone_1), int32(3229), int32(_a_F_timetz_zone_2))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v56 = F_palloc(m, int32(16))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								v58 = base.I32_wrap_i64(v17)
								v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+332))
								*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v61
								v66 = base.I64_extend_i32_s(v60-v61) * int64(1000000)
								v67 = v59 + v66
								v68 = int64(0)
								if v68 < v67 {
									v71 = v67
								} else {
									v71 = v68
								}
								v72 = v71 - v59
								v74 = base.I64_extend_i32_u(base.B2i32(v72 != v66))
								v77 = int64(86400000000)
								v78 = base.I64_div_u_s(v72-(v66|v74), v77)
								v83 = v59 + (v78+v74)*v77 + v66
								v85 = base.I64_rem_u_s(v83, v77)
								if base.Ui64(int64(86399999999)) < base.Ui64(v83) {
									v88 = v85
								} else {
									v88 = v83
								}
								*(*int64)(unsafe.Add(mBase, uint32(v56))) = v88
								m.G0 = v10 + int32(336)
								return base.I64_extend_i32_u(v56)
							}
						}
					}
				}
			}
		}
	}
}
