package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TSDictionaryIsVisibleExt(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = Fn14223(m, l0, l1, int32(75), int32(_a_F_TSDictionaryIsVisibleExt_0), int32(3019), int32(_a_F_TSDictionaryIsVisibleExt_1), int32(76))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
func F_check_commit_ts_buffers(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_check_slru_buffers(m, int32(_a_F_check_commit_ts_buffers_0), l0)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_getTSCurrentConfig(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[0]))
	if v8 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[1]))
		if v12 != 0 {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v13 != 0 {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[2]))
				if v30 != 0 {
					v58 = *(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[1]))
					v60 = F_stringToQualifiedNameList(m, v58, int32(0))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v63 = F_get_ts_config_oid(m, v60, int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[0])) = v63
							v66 = v63
							m.G0 = v5 + int32(48)
							return v66
						}
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = int64(85899345924)
					v37 = F_hash_create(m, int32(_a_F_getTSCurrentConfig_0), int64(16), v5, int32(40))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[2])) = v37
						F_CacheRegisterSyscacheCallback(m, int32(74), int32(1811), base.I64_extend_i32_u(v37))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v48 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[2])))
							F_CacheRegisterSyscacheCallback(m, int32(72), int32(1811), v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[3]))
								if v52 != 0 {
									v58 = *(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[1]))
									v60 = F_stringToQualifiedNameList(m, v58, int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										v63 = F_get_ts_config_oid(m, v60, int32(0))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[0])) = v63
											v66 = v63
											m.G0 = v5 + int32(48)
											return v66
										}
									}
								} else {
									F_CreateCacheMemoryContext(m)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v58 = *(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[1]))
										v60 = F_stringToQualifiedNameList(m, v58, int32(0))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											v63 = F_get_ts_config_oid(m, v60, int32(0))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_getTSCurrentConfig[0])) = v63
												v66 = v63
												m.G0 = v5 + int32(48)
												return v66
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
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_getTSCurrentConfig_1), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_getTSCurrentConfig_2), int32(584), int32(_a_F_getTSCurrentConfig_3))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_getTSCurrentConfig_1), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_getTSCurrentConfig_2), int32(584), int32(_a_F_getTSCurrentConfig_3))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		v66 = v8
		m.G0 = v5 + int32(48)
		return v66
	}
}
func F_get_ts_dict_oid(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14294(m, l0, l1, int32(_a_F_get_ts_dict_oid_0), int32(2979), int32(_a_F_get_ts_dict_oid_1), int32(75))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_get_ts_template_func(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = F_defGetQualifiedName(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = int64(9796820404457)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v14
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v14
		v19 = int32(4)
		if l1 == v19 {
			v22 = int32(1)
		} else {
			v22 = v19
		}
		v24 = v8 + int32(16)
		v26 = F_LookupFuncName(m, v10, v22, v24, int32(0))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			v28 = F_get_func_rettype(m, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				if v28 != int32(2281) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v40 = F_func_signature_string(m, v10, v22, int32(0), v24)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int64(0)
							} else {
								v43 = F_format_type_be(m, int32(2281))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v43
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v40
									F_errmsg(m, int32(_a_F_get_ts_template_func_0), v8)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_get_ts_template_func_1), int32(643), int32(_a_F_get_ts_template_func_2))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
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
				} else {
					m.G0 = v8 + int32(32)
					return base.I64_extend_i32_u(v26)
				}
			}
		}
	}
}
func F_ts_dist(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v9 int64
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(2)) <= base.Ui64(v4-int64(9223372036854775807)) {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if base.Ui64(int64(1)) < base.Ui64(v9-int64(9223372036854775807)) {
			v28 = F_DirectFunctionCall2Coll(m, int32(1714), int32(0), v4, v9)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v31 = F_abs_interval(m, base.I32_wrap_i64(v28))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v31)
				}
			}
		} else {
			v16 = F_palloc(m, int32(16))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(9223372036854775807)
				*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(9223372034707292159)
				return base.I64_extend_i32_u(v16)
			}
		}
	} else {
		v16 = F_palloc(m, int32(16))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(9223372036854775807)
			*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(9223372034707292159)
			return base.I64_extend_i32_u(v16)
		}
	}
}
func F_ts_headline(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	v4 = F_getTSCurrentConfig(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_DirectFunctionCall3Coll(m, int32(1286), int32(0), base.I64_extend_i32_u(v4), v9, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			return v11
		}
	}
}
func F_ts_headline_json_byid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = F_DirectFunctionCall3Coll(m, int32(1289), int32(0), v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_ts_headline_json_byid_opt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int64
	_ = v42
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v14 = m.G0
	v16 = v14 - int32(48)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v26 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v26 < int32(4) {
		v36 = int32(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v38 = F_palloc0(m, int32(24))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v29 = int32(0)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v30 == v29 {
		v36 = v29
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v33 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v36 = v33
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = int32(0)
	v42 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+36)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v16)+28)) = v42
	*(*int64)(unsafe.Add(mBase, uint32(v16)+20)) = v42
	v48 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v48
	v52 = F_palloc_mul(m, int32(16), v48)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v16 + int32(12)
	v58 = F_lookup_ts_config_cache(m, v18)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v58
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v62 = F_lookup_ts_parser_cache(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v62
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v67 = F_deserialize_deflist(m, base.I64_extend_i32_u(v36))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v70 = v62
	v71 = int32(0)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v71
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	if v73 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v70 = v69
	v71 = v67
	goto L13
L15:
	;
	v74 = m.G0
	v76 = v74 - int32(96)
	m.G0 = v76
	v79 = F_palloc0(m, int32(40))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L63
	}
L18:
	;
	v82 = F_palloc0(m, int32(16))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_initStringInfo(m, v76+int32(12))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v88 = F_pg_detoast_datum_packed(m, v20)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v121 = v76 + int32(28)
	v122 = int32(1)
	if v90&v122 != 0 {
		goto L33
	} else {
		goto L34
	}
L22:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v90 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	if v96 == int32(18) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v107 = int32(1)
	if v90&v107 != 0 {
		v119 = int32(base.Ui32(v90)>>(uint(v107)%32)) - v107
		goto L21
	} else {
		goto L32
	}
L26:
	;
	v99 = int32(16)
	goto L28
L27:
	;
	v99 = int32(0)
	goto L28
L28:
	;
	if base.Ui32((v96-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v106 = int32(4)
	goto L31
L30:
	;
	v106 = v99
	goto L31
L31:
	;
	v119 = v106
	goto L21
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v119 = int32(base.Ui32(v113)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L33:
	;
	v126 = v122
	goto L35
L34:
	;
	v126 = int32(4)
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_ts_headline_json_byid_opt[0]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	goto L36
L36:
	;
	v132 = F_makeJsonLexContextCstringLen(m, v121, v88+v126, v119, v130, int32(1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+12)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = int32(1287)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v76 + int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+36)) = int32(1520)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+16)) = int32(1521)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = int32(1522)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+8)) = int32(1523)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = int32(1524)
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v79)+28)) = int32(1525)
	*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = int32(1526)
	v156 = F_pg_parse_json(m, v121, v79)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v156 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_json_errsave_error(m, v156, v121, int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	F_freeJsonLexContext(m, v76+int32(28))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v168 = F_cstring_to_text_with_len(m, v166, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	m.G0 = v76 + int32(96)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v173 != v20 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_pfree(m, v20)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v177 != v24 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L47
L49:
	;
	F_pfree(m, v24)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	if v36 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L51
L53:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_pfree(m, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L57
	}
L54:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v36 == v183 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	F_pfree(m, v36)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+20)))
	if v190 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	F_pfree(m, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	m.G0 = v16 + int32(48)
	return base.I64_extend_i32_u(v168)
L61:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	F_pfree(m, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	F_errmsg(m, int32(_a_F_ts_headline_json_byid_opt_0), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_ts_headline_json_byid_opt_1), int32(471), int32(_a_F_ts_headline_json_byid_opt_2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ts_headline_jsonb_byid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = F_DirectFunctionCall3Coll(m, int32(1288), int32(0), v4, v5, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_ts_headline_jsonb_byid_opt(m *base.Module, l0 int32) int64 {
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int64
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = F_pg_detoast_datum(m, v16)
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v23 < int32(4) {
		v33 = int32(0)
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = F_palloc0(m, int32(24))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v27 == v26 {
		v33 = v26
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v33 = v30
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = int32(0)
	v39 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+36)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v13)+28)) = v39
	*(*int64)(unsafe.Add(mBase, uint32(v13)+20)) = v39
	v45 = int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v45
	v49 = F_palloc_mul(m, int32(16), v45)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v13 + int32(12)
	v55 = F_lookup_ts_config_cache(m, v15)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v59 = F_lookup_ts_parser_cache(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v59
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v64 = F_deserialize_deflist(m, base.I64_extend_i32_u(v33))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v67 = v59
	v68 = int32(0)
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	if v70 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v67 = v66
	v68 = v64
	goto L13
L15:
	;
	v71 = m.G0
	v73 = v71 + int32(-64)
	m.G0 = v73
	*(*int32)(unsafe.Add(mBase, uint32(v73)+16)) = int32(0)
	v77 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v73)+8)) = v77
	*(*int64)(unsafe.Add(mBase, uint32(v73))) = v77
	v83 = F_JsonbIteratorInit(m, v17+int32(4))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L70
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+60)) = v83
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+8)))
	v92 = F_JsonbIteratorNext(m, v71+int32(-4), v71+int32(-40), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v95 = v92
	goto L20
L20:
	;
	if v95&int32(-2) != int32(2) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	if v177 == int32(16) {
		goto L48
	} else {
		goto L49
	}
L22:
	;
	F_pushJsonbValue(m, v73, v95, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L45
	}
L23:
	;
	if base.Ui32(v95) < base.Ui32(int32(4)) {
		goto L42
	} else {
		goto L43
	}
L24:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v73)+24))
	if v108 != int32(1) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v73)+36))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v73)+32))
	v113 = F_headline_json_value(m, v35, v111, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v115 = F_pg_detoast_datum_packed(m, v113)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v117 = int32(1)
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v119&v117 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v122 = v117
	goto L30
L29:
	;
	v122 = int32(4)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+36)) = v115 + v122
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v125 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+32)) = v154
	v166 = v71 + int32(-40)
	goto L22
L32:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	if v131 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v142 = int32(1)
	if v125&v142 != 0 {
		v154 = int32(base.Ui32(v125)>>(uint(v142)%32)) - v142
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v134 = int32(16)
	goto L37
L36:
	;
	v134 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v131-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v141 = int32(4)
	goto L40
L39:
	;
	v141 = v134
	goto L40
L40:
	;
	v154 = v141
	goto L31
L41:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v154 = int32(base.Ui32(v148)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	v163 = v71 + int32(-40)
	goto L44
L43:
	;
	v163 = int32(0)
	goto L44
L44:
	;
	v166 = v163
	goto L22
L45:
	;
	v174 = F_JsonbIteratorNext(m, v71+int32(-4), v71+int32(-40), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v174 != 0 {
		v95 = v174
		goto L20
	} else {
		goto L47
	}
L47:
	;
	goto L21
L48:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v176)+16)) = uint8(v86)
	goto L50
L49:
	;
	goto L50
L50:
	;
	v181 = F_JsonbValueToJsonb(m, v176)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	m.G0 = v73 - int32(-64)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v186 != v17 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_pfree(m, v17)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v190 != v21 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	goto L54
L56:
	;
	F_pfree(m, v21)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v33 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L58
L60:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	F_pfree(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L64
	}
L61:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v33 == v196 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	F_pfree(m, v33)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+20)))
	if v203 == int32(1) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	F_pfree(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	m.G0 = v13 + int32(48)
	return base.I64_extend_i32_u(v181)
L68:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	F_pfree(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_ts_headline_jsonb_byid_opt_0), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_ts_headline_jsonb_byid_opt_1), int32(394), int32(_a_F_ts_headline_jsonb_byid_opt_2))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ts_parse_byid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	if v5 != 0 {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
		v23 = F_prs_process_call(m, v22)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int64(0)
		} else {
			if v23 != int64(0) {
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
				*(*int64)(unsafe.Add(mBase, uint32(v22))) = v27 + int64(1)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
				return v23
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(2)
					v40 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
					return v23
				}
			}
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v7 = F_pg_detoast_datum_packed(m, v6)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v11 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				F_prs_setup_firstcall(m, v11, l0, v13, v7)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int64(0)
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v7 == v16 {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
						v23 = F_prs_process_call(m, v22)
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
						} else {
							if v23 != int64(0) {
								v27 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
								*(*int64)(unsafe.Add(mBase, uint32(v22))) = v27 + int64(1)
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
								return v23
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int64(0)
								} else {
									v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(2)
									v40 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
									return v23
								}
							}
						}
					} else {
						F_pfree(m, v7)
						mBase = m.M
						v19 = m.ExcPending
						if v19 != 0 {
							return int64(0)
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
							v23 = F_prs_process_call(m, v22)
							mBase = m.M
							v24 = m.ExcPending
							if v24 != 0 {
								return int64(0)
							} else {
								if v23 != int64(0) {
									v27 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
									*(*int64)(unsafe.Add(mBase, uint32(v22))) = v27 + int64(1)
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = int32(1)
									return v23
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int64(0)
									} else {
										v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v37)+20)) = int32(2)
										v40 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
										return v23
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
func F_ts_rank_wtt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 float32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			F_getWeights(m, v12, v9)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int64(0)
			} else {
				v23 = F_calc_rank(m, v9, v17, v19, int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v25 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int64(0)
						} else {
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v29 != v17 {
								F_pfree(m, v17)
								mBase = m.M
								v32 = m.ExcPending
								if v32 != 0 {
									return int64(0)
								} else {
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v33 != v19 {
										F_pfree(m, v19)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int64(0)
										} else {
											m.G0 = v9 + int32(16)
											return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
										}
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
									}
								}
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v33 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int64(0)
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
									}
								} else {
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v29 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								if v33 != v19 {
									F_pfree(m, v19)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int64(0)
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
									}
								} else {
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
								}
							}
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							if v33 != v19 {
								F_pfree(m, v19)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int64(0)
								} else {
									m.G0 = v9 + int32(16)
									return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
								}
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_s(base.I32_reinterpret_f32(v23))
							}
						}
					}
				}
			}
		}
	}
}
func F_ts_stat2(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = F_pg_detoast_datum_packed(m, v11)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v17 = F_pg_detoast_datum_packed(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					F_SPI_connect_ext(m, int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
						v25 = F_ts_stat_sql(m, v24, v12, v17)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v27 != v12 {
								F_pfree(m, v12)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int64(0)
								} else {
									v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v31 != v17 {
										F_pfree(m, v17)
										mBase = m.M
										v34 = m.ExcPending
										if v34 != 0 {
											return int64(0)
										} else {
											F_ts_setup_firstcall(m, l0, v19, v25)
											mBase = m.M
											v36 = m.ExcPending
											if v36 != 0 {
												return int64(0)
											} else {
												v37 = F_SPI_finish(m)
												mBase = m.M
												v38 = m.ExcPending
												if v38 != 0 {
													return int64(0)
												} else {
													v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
													v45 = F_ts_process_call(m, v44)
													mBase = m.M
													v46 = m.ExcPending
													if v46 != 0 {
														return int64(0)
													} else {
														if v45 != int64(0) {
															v49 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
															*(*int64)(unsafe.Add(mBase, uint32(v44))) = v49 + int64(1)
															v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(1)
															return v45
														} else {
															F_end_MultiFuncCall(m, l0)
															mBase = m.M
															v58 = m.ExcPending
															if v58 != 0 {
																return int64(0)
															} else {
																v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = int32(2)
																v62 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v62)
																return v45
															}
														}
													}
												}
											}
										}
									} else {
										F_ts_setup_firstcall(m, l0, v19, v25)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int64(0)
										} else {
											v37 = F_SPI_finish(m)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return int64(0)
											} else {
												v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
												v45 = F_ts_process_call(m, v44)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int64(0)
												} else {
													if v45 != int64(0) {
														v49 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
														*(*int64)(unsafe.Add(mBase, uint32(v44))) = v49 + int64(1)
														v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(1)
														return v45
													} else {
														F_end_MultiFuncCall(m, l0)
														mBase = m.M
														v58 = m.ExcPending
														if v58 != 0 {
															return int64(0)
														} else {
															v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = int32(2)
															v62 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v62)
															return v45
														}
													}
												}
											}
										}
									}
								}
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v31 != v17 {
									F_pfree(m, v17)
									mBase = m.M
									v34 = m.ExcPending
									if v34 != 0 {
										return int64(0)
									} else {
										F_ts_setup_firstcall(m, l0, v19, v25)
										mBase = m.M
										v36 = m.ExcPending
										if v36 != 0 {
											return int64(0)
										} else {
											v37 = F_SPI_finish(m)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return int64(0)
											} else {
												v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
												v45 = F_ts_process_call(m, v44)
												mBase = m.M
												v46 = m.ExcPending
												if v46 != 0 {
													return int64(0)
												} else {
													if v45 != int64(0) {
														v49 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
														*(*int64)(unsafe.Add(mBase, uint32(v44))) = v49 + int64(1)
														v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(1)
														return v45
													} else {
														F_end_MultiFuncCall(m, l0)
														mBase = m.M
														v58 = m.ExcPending
														if v58 != 0 {
															return int64(0)
														} else {
															v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = int32(2)
															v62 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v62)
															return v45
														}
													}
												}
											}
										}
									}
								} else {
									F_ts_setup_firstcall(m, l0, v19, v25)
									mBase = m.M
									v36 = m.ExcPending
									if v36 != 0 {
										return int64(0)
									} else {
										v37 = F_SPI_finish(m)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return int64(0)
										} else {
											v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
											v45 = F_ts_process_call(m, v44)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return int64(0)
											} else {
												if v45 != int64(0) {
													v49 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
													*(*int64)(unsafe.Add(mBase, uint32(v44))) = v49 + int64(1)
													v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(1)
													return v45
												} else {
													F_end_MultiFuncCall(m, l0)
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return int64(0)
													} else {
														v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = int32(2)
														v62 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v62)
														return v45
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
	} else {
		v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
		v45 = F_ts_process_call(m, v44)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int64(0)
		} else {
			if v45 != int64(0) {
				v49 = *(*int64)(unsafe.Add(mBase, uint32(v44)))
				*(*int64)(unsafe.Add(mBase, uint32(v44))) = v49 + int64(1)
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = int32(1)
				return v45
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int64(0)
				} else {
					v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = int32(2)
					v62 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v62)
					return v45
				}
			}
		}
	}
}
