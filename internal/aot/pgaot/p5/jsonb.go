package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convertJsonbScalar(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	v12 = m.G0
	v14 = v12 - int32(144)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v16 {
	case 0:
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1073741824)
		m.G0 = v14 + int32(144)
		return
	case 1:
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		F_enlargeStringInfo(m, l0, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v22 = v18 + v21
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v26 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v24+v22))) = uint8(v26)
			if v18 != 0 {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				base.MemoryCopy(m, v28+v21, v17, v18)
			} else {
			}
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
			m.G0 = v14 + int32(144)
			return
		}
	case 2:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
		if v34 == int32(1) {
			v38 = int32(18)
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
			if v40 == v38 {
				v43 = v38
			} else {
				v43 = int32(2)
			}
			if base.Ui32((v40-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v50 = int32(6)
			} else {
				v50 = v43
			}
			v59 = v50
		} else {
			v51 = int32(1)
			if v34&v51 != 0 {
				v59 = int32(base.Ui32(v34) >> (uint(v51) % 32))
			} else {
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				v59 = int32(base.Ui32(v55) >> (uint(int32(2)) % 32))
			}
		}
		v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v64 = (v60 + int32(3)) & int32(-4)
		v65 = v64 - v60
		F_enlargeStringInfo(m, l0, v65)
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v69 = v65 + v68
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v69
			v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v73 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v71+v69))) = uint8(v73)
			if v65 <= v73 {
			} else {
				v78 = v65 & int32(3)
				v79 = int32(0)
				if base.Ui32(v60-v64) <= base.Ui32(int32(-4)) {
					v89 = v79
					v92 = int32(0)
					for {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v100 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v97+v68+v89))) = uint8(v100)
						v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*uint8)(unsafe.Add(mBase, uint32(v102+v68+v89)+1)) = uint8(v100)
						v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*uint8)(unsafe.Add(mBase, uint32(v107+v68+v89)+2)) = uint8(v100)
						v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*uint8)(unsafe.Add(mBase, uint32(v112+v68+v89)+3)) = uint8(v100)
						v117 = int32(4)
						v118 = v89 + v117
						v120 = v92 + v117
						if v120 != v65&int32(2147483644) {
							v89 = v118
							v92 = v120
							continue
						} else {
							break
						}
						break
					}
					if v78 == int32(0) {
					} else {
						v127 = v118
						v139 = v127
						v142 = int32(0)
						for {
							v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v150 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v147+v68+v139))) = uint8(v150)
							v152 = int32(1)
							v155 = v142 + v152
							if v155 != v78 {
								v139 = v139 + v152
								v142 = v155
								continue
							} else {
								break
							}
							break
						}
					}
				} else {
					v127 = v79
					v139 = v127
					v142 = int32(0)
					for {
						v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v150 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v147+v68+v139))) = uint8(v150)
						v152 = int32(1)
						v155 = v142 + v152
						if v155 != v78 {
							v139 = v139 + v152
							v142 = v155
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			F_enlargeStringInfo(m, l0, v59)
			mBase = m.M
			v170 = m.ExcPending
			if v170 != 0 {
				return
			} else {
				v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v172 = v171 + v59
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v172
				v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v176 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v174+v172))) = uint8(v176)
				if v59 != 0 {
					v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					base.MemoryCopy(m, v178+v171, v168, v59)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = base.I32_extend16_s(v65) + v59 | int32(268435456)
				m.G0 = v14 + int32(144)
				return
			}
		}
	case 3:
		v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		if v188 != 0 {
			v189 = int32(805306368)
		} else {
			v189 = int32(536870912)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v189
		m.G0 = v14 + int32(144)
		return
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v214 = m.ExcPending
		if v214 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_convertJsonbScalar_0), int32(0))
			mBase = m.M
			v218 = m.ExcPending
			if v218 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_convertJsonbScalar_1), int32(1990), int32(_a_F_convertJsonbScalar_2))
				mBase = m.M
				v223 = m.ExcPending
				if v223 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 32:
		v191 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
		v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		v195 = F_JsonEncodeDateTime(m, v14, v191, v192, l2+int32(24))
		mBase = m.M
		v196 = m.ExcPending
		if v196 != 0 {
			return
		} else {
			v197 = F_strlen(m, v14)
			mBase = m.M
			F_enlargeStringInfo(m, l0, v197)
			mBase = m.M
			v199 = m.ExcPending
			if v199 != 0 {
				return
			} else {
				v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v201 = v197 + v200
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v201
				v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v205 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v203+v201))) = uint8(v205)
				if v197 != 0 {
					v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					base.MemoryCopy(m, v207+v200, v14, v197)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v197
				m.G0 = v14 + int32(144)
				return
			}
		}
	}
}
func F_get_jsonb_path_all(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
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
			v19 = F_array_contains_nulls(m, v17)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				if v19 == int32(0) {
					F_deconstruct_array_builtin(m, v17, int32(25), v9+int32(12), v9+int32(8), v9)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
						v34 = F_jsonb_get_element(m, v12, v30, v31, v9+int32(7), l1)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)))
							if v36 != int32(1) {
								v43 = v34
							} else {
								v40 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
								v43 = int64(0)
							}
							m.G0 = v9 + int32(16)
							return v43
						}
					}
				} else {
					v40 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v40)
					v43 = int64(0)
					m.G0 = v9 + int32(16)
					return v43
				}
			}
		}
	}
}
func F_jsonb_agg_transfn(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_agg_transfn_worker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_array_length(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		if v7&int32(268435456) == int32(0) {
			if v7&int32(1073741824) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_jsonb_array_length_0), int32(0))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_jsonb_array_length_1), int32(1889), int32(_a_F_jsonb_array_length_2))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
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
				return base.I64_extend_i32_u(v7 & int32(268435455))
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_jsonb_array_length_3), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_jsonb_array_length_1), int32(1885), int32(_a_F_jsonb_array_length_2))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
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
func F_jsonb_build_array(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v91 int64
	_ = v91
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v18 = F_extract_variadic_args(m, l0, v10+int32(20), v10+int32(12), v10+int32(16))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(48)
	return v91
L2:
	;
	return int64(0)
L3:
	;
	if v18 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
	v91 = int64(0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
	v30 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v30
	v32 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v32
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v32
	F_pushJsonbValue(m, v10+int32(24), int32(4), v30)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v44 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	F_pushJsonbValue(m, v10+int32(24), int32(5), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L15
	}
L11:
	;
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v29+v44<<(uint(int32(3))%32))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v28))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v27+v44<<(uint(int32(2))%32))))
	F_add_jsonb(m, v52, v54, v10+int32(24), v60, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v65 = v44 + int32(1)
	if v65 != v18 {
		v44 = v65
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	v81 = F_JsonbValueToJsonb(m, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v91 = base.I64_extend_i32_u(v81)
	goto L1
}
func F_jsonb_build_object(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = F_extract_variadic_args(m, l0, v7+int32(12), v7+int32(4), v7+int32(8))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		if v15 < int32(0) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v31 = int64(0)
			m.G0 = v7 + int32(16)
			return v31
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v27 = int32(0)
			v29 = F_jsonb_build_object_worker(m, v15, v24, v25, v26, v27, v27)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = v29
				m.G0 = v7 + int32(16)
				return v31
			}
		}
	}
}
func F_jsonb_concat(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v306 int32
	_ = v306
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	v6 = m.G0
	v8 = v6 - int32(96)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
	v20 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v20
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if (v24^v25)&int32(536870912) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v8 + int32(96)
	return base.I64_extend_i32_u(v345)
L5:
	;
	v43 = F_JsonbIteratorInit(m, v11+int32(4))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	if v24&int32(268435456)|v25&int32(268435455) == int32(0) {
		v345 = v16
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v24&int32(268435455)|v25&int32(268435456) != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v345 = v11
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v43
	v48 = F_JsonbIteratorInit(m, v16+int32(4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v48
	v52 = v8 + int32(4)
	v54 = v8 - int32(-64)
	v56 = F_JsonbIteratorNext(m, v52, v54, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v343 = F_JsonbValueToJsonb(m, v342)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L105
	}
L12:
	;
	v315 = F_JsonbIteratorNext(m, v8, v8+int32(32), int32(1))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L95
	}
L13:
	;
	F_pushJsonbValue(m, v8+int32(8), int32(5), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L94
	}
L14:
	;
	v275 = F_JsonbIteratorNext(m, v8, v8+int32(32), int32(1))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L87
	}
L15:
	;
	v249 = v99
	goto L82
L16:
	;
	v231 = v77
	goto L77
L17:
	;
	v63 = F_JsonbIteratorNext(m, v8, v8+int32(32), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if base.B2i32(v56 != int32(6))|base.B2i32(v63 != int32(6)) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	F_pushJsonbValue(m, v8+int32(8), int32(6), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	F_pushJsonbValue(m, v8+int32(8), int32(4), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L25
	}
L22:
	;
	v77 = F_JsonbIteratorNext(m, v52, v54, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v77 != int32(7) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	v87 = int32(4)
	if base.B2i32(v56 != v87)|base.B2i32(v63 != v87) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v99 = F_JsonbIteratorNext(m, v8+int32(4), v8-int32(-64), int32(1))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v56 == int32(6) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v99 != int32(5) {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	goto L14
L31:
	;
	F_pushJsonbValue(m, v8+int32(8), int32(6), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v173 = F_JsonbIteratorNext(m, v8+int32(4), v8-int32(-64), int32(1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L57
	}
L34:
	;
	v116 = F_JsonbIteratorNext(m, v8+int32(4), v8-int32(-64), int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if v116 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v118 = v116
	goto L39
L37:
	;
	goto L38
L38:
	;
	v146 = F_JsonbIteratorNext(m, v8, v8+int32(32), int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L47
	}
L39:
	;
	v126 = v8 - int32(-64)
	if v118 != int32(7) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L38
L41:
	;
	v130 = v126
	goto L43
L42:
	;
	v130 = int32(0)
	goto L43
L43:
	;
	F_pushJsonbValue(m, v8+int32(8), v118, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v136 = F_JsonbIteratorNext(m, v8+int32(4), v126, int32(1))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	if v136 != 0 {
		v118 = v136
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L40
L47:
	;
	if v146 == int32(0) {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	v150 = v146
	goto L49
L49:
	;
	v158 = v8 + int32(32)
	if v150 != int32(5) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L11
L51:
	;
	v162 = v158
	goto L53
L52:
	;
	v162 = int32(0)
	goto L53
L53:
	;
	F_pushJsonbValue(m, v8+int32(8), v150, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v166 = F_JsonbIteratorNext(m, v8, v158, int32(1))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v166 != 0 {
		v150 = v166
		goto L49
	} else {
		goto L56
	}
L56:
	;
	goto L50
L57:
	;
	if v173 != int32(5) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v177 = v173
	goto L61
L59:
	;
	goto L60
L60:
	;
	F_pushJsonbValue(m, v8+int32(8), int32(6), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L66
	}
L61:
	;
	v185 = v8 - int32(-64)
	F_pushJsonbValue(m, v8+int32(8), v177, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v191 = F_JsonbIteratorNext(m, v8+int32(4), v185, int32(1))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v191 != int32(5) {
		v177 = v191
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v209 = F_JsonbIteratorNext(m, v8, v8+int32(32), int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v209 == int32(0) {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	v213 = v209
	goto L69
L69:
	;
	v221 = v8 + int32(32)
	if v213 != int32(7) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L13
L71:
	;
	v225 = v221
	goto L73
L72:
	;
	v225 = int32(0)
	goto L73
L73:
	;
	F_pushJsonbValue(m, v8+int32(8), v213, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v229 = F_JsonbIteratorNext(m, v8, v221, int32(1))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v229 != 0 {
		v213 = v229
		goto L69
	} else {
		goto L76
	}
L76:
	;
	goto L70
L77:
	;
	v239 = v8 - int32(-64)
	F_pushJsonbValue(m, v8+int32(8), v231, v239)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L12
L79:
	;
	v245 = F_JsonbIteratorNext(m, v8+int32(4), v239, int32(1))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v245 != int32(7) {
		v231 = v245
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	v257 = v8 - int32(-64)
	F_pushJsonbValue(m, v8+int32(8), v249, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	goto L14
L84:
	;
	v263 = F_JsonbIteratorNext(m, v8+int32(4), v257, int32(1))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v263 != int32(5) {
		v249 = v263
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	if v275 == int32(5) {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	goto L89
L89:
	;
	v288 = v8 + int32(32)
	F_pushJsonbValue(m, v8+int32(8), int32(3), v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	goto L13
L91:
	;
	v292 = F_JsonbIteratorNext(m, v8, v288, int32(1))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v292 != int32(5) {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	goto L90
L94:
	;
	goto L11
L95:
	;
	if v315 == int32(0) {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	v319 = v315
	goto L97
L97:
	;
	v327 = v8 + int32(32)
	if v319 != int32(7) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L11
L99:
	;
	v331 = v327
	goto L101
L100:
	;
	v331 = int32(0)
	goto L101
L101:
	;
	F_pushJsonbValue(m, v8+int32(8), v319, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v335 = F_JsonbIteratorNext(m, v8, v327, int32(1))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v335 != 0 {
		v319 = v335
		goto L97
	} else {
		goto L104
	}
L104:
	;
	goto L98
L105:
	;
	v345 = v343
	goto L4
}
func F_jsonb_exists(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
			v24 = v22 & v18
			if v24 != 0 {
				v25 = v18
			} else {
				v25 = int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16 + v25
			if v22 == int32(1) {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
				if v33 == int32(18) {
					v36 = int32(16)
				} else {
					v36 = int32(0)
				}
				if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v43 = int32(4)
				} else {
					v43 = v36
				}
				v54 = v43
			} else {
				v44 = int32(1)
				if v24 != 0 {
					v54 = int32(base.Ui32(v22)>>(uint(v44)%32)) - v44
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v54
			v59 = F_findJsonbValueFromContainer(m, v11+int32(4), int32(1610612736), v8)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int64(0)
			} else {
				m.G0 = v8 + int32(32)
				return base.I64_extend_i32_u(base.B2i32(v59 != int32(0)))
			}
		}
	}
}
func F_jsonb_exists_all(m *base.Module, l0 int32) int64 {
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
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int64
	_ = v110
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_deconstruct_array_builtin(m, v17, int32(25), v9+int32(44), v9+int32(40), v9+int32(36))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = int64(1)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	if v29 <= int32(0) {
		v110 = v28
		goto L5
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v9 + int32(48)
	return v110
L6:
	;
	v35 = int32(0)
	v37 = v29
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+40))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41+v35))))
	if v43 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v110 = int64(0)
	goto L5
L9:
	;
	goto L8
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v47
	v51 = v46 + v35<<(uint(int32(3))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v55&v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v99 = v37
	goto L12
L12:
	;
	v102 = v35 + int32(1)
	if v102 < v99 {
		v35 = v102
		v37 = v99
		goto L7
	} else {
		goto L29
	}
L13:
	;
	v58 = v47
	goto L15
L14:
	;
	v58 = int32(4)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v52 + v58
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
	if v62 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v91
	v94 = F_findJsonbValueFromContainer(m, v12+int32(4), int32(1610612736), v9)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L27
	}
L17:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v68 == int32(18) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v79 = int32(1)
	if v62&v79 != 0 {
		v91 = int32(base.Ui32(v62)>>(uint(v79)%32)) - v79
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v71 = int32(16)
	goto L22
L21:
	;
	v71 = int32(0)
	goto L22
L22:
	;
	if base.Ui32((v68-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = int32(4)
	goto L25
L24:
	;
	v78 = v71
	goto L25
L25:
	;
	v91 = v78
	goto L16
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	if v94 == int32(0) {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
	v99 = v98
	goto L12
L29:
	;
	v110 = v28
	goto L5
}
func F_jsonb_hash_extended(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v385 int64
	_ = v385
	var v386 int32
	_ = v386
	var v391 int64
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int64
	_ = v438
	var v444 int64
	_ = v444
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v20&int32(268435455) != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v25 = F_JsonbIteratorInit(m, v13+int32(4))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v444 = v17
	goto L5
L5:
	;
	m.G0 = v10 - int32(-64)
	return v444
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = v25
	goto L8
L7:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v434 != v13 {
		goto L78
	} else {
		goto L79
	}
L8:
	;
	v40 = F_JsonbIteratorNext(m, v8+int32(-4), v8+int32(-40), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L14
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L75
	}
L10:
	;
	goto L9
L11:
	;
	v51 = v8 + int32(-48)
	v53 = v8 + int32(-40)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	switch v54 {
	case 0:
		goto L16
	case 1:
		goto L20
	case 2:
		goto L19
	case 3:
		goto L18
	default:
		goto L17
	}
L12:
	;
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v46 ^ int64(2305843009750564864)
	goto L8
L13:
	;
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v42 ^ int64(4611686019501129728)
	goto L8
L14:
	;
	switch v40 {
	case 0:
		goto L7
	case 1, 2, 3:
		goto L11
	case 4:
		goto L13
	case 5, 7:
		goto L8
	case 6:
		goto L12
	default:
		goto L10
	}
L15:
	;
	v409 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v408 ^ (v409<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v409)>>(uint(int64(31))%64))&int64(4294967297))
	goto L8
L16:
	;
	v408 = v17 + int64(1)
	goto L15
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L72
	}
L18:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+8)))
	if v17 != int64(0) {
		goto L65
	} else {
		goto L66
	}
L19:
	;
	v374 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v53)+8)))
	v375 = F_DirectFunctionCall2Coll(m, int32(1470), int32(0), v374, v17)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L64
	}
L20:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	v62 = v56 - int32(1636608432)
	if v17 == int64(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v408 = base.I64_extend_i32_u(v362)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v362^v354-base.I32_rotl(v362, int32(24)))
	goto L15
L22:
	;
	if v55&int32(3) != 0 {
		goto L38
	} else {
		goto L39
	}
L23:
	;
	v99 = v62
	v101 = v62
	v103 = v62
	goto L22
L24:
	;
	goto L25
L25:
	;
	v66 = v62 + base.I32_wrap_i64(v17)
	v67 = v66 + v62
	v71 = int32(4)
	v73 = base.I32_wrap_i64(int64(base.Ui64(v17)>>(uint(int64(32))%64))) ^ base.I32_rotl(v62, v71)
	v77 = v66 - v73 ^ base.I32_rotl(v73, int32(6))
	v81 = v67 - v77 ^ base.I32_rotl(v77, int32(8))
	v82 = v67 + v73
	v83 = v77 + v82
	v84 = v81 + v83
	v88 = v82 - v81 ^ base.I32_rotl(v81, int32(16))
	v92 = v83 - v88 ^ base.I32_rotl(v88, int32(19))
	v97 = v84 + v88
	v99 = v97
	v101 = v84 - v92 ^ base.I32_rotl(v92, v71)
	v103 = v92 + v97
	goto L22
L26:
	;
	v340 = int32(14)
	v342 = v336 ^ v337 - base.I32_rotl(v336, v340)
	v346 = v342 ^ v335 - base.I32_rotl(v342, int32(11))
	v350 = v346 ^ v336 - base.I32_rotl(v346, int32(25))
	v354 = v350 ^ v342 - base.I32_rotl(v350, int32(16))
	v358 = v354 ^ v346 - base.I32_rotl(v354, int32(4))
	v362 = v358 ^ v350 - base.I32_rotl(v358, v340)
	goto L21
L27:
	;
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	v335 = v327 + v330
	v336 = v328
	v337 = v329
	goto L26
L28:
	;
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+1)))
	v327 = v323<<(uint(int32(8))%32) + v320
	v328 = v321
	v329 = v322
	goto L27
L29:
	;
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+2)))
	v320 = v316<<(uint(int32(16))%32) + v313
	v321 = v314
	v322 = v315
	goto L28
L30:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+3)))
	v313 = v309<<(uint(int32(24))%32) + v160
	v314 = v307
	v315 = v308
	goto L29
L31:
	;
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+4)))
	v307 = v303 + v305
	v308 = v304
	goto L30
L32:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+5)))
	v303 = v299<<(uint(int32(8))%32) + v297
	v304 = v298
	goto L31
L33:
	;
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
	v297 = v293<<(uint(int32(16))%32) + v291
	v298 = v292
	goto L32
L34:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+7)))
	v291 = v287<<(uint(int32(24))%32) + v161
	v292 = v286
	goto L33
L35:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+8)))
	v286 = v282<<(uint(int32(8))%32) + v281
	goto L34
L36:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+9)))
	v281 = v277<<(uint(int32(16))%32) + v276
	goto L35
L37:
	;
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+10)))
	v276 = v272<<(uint(int32(24))%32) + v162
	goto L36
L38:
	;
	if base.Ui32(int32(11)) < base.Ui32(v56) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v56) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v108 = v55
	v109 = v56
	v111 = v99
	v112 = v103
	v113 = v101
	goto L44
L42:
	;
	v157 = v55
	v158 = v56
	v160 = v99
	v161 = v103
	v162 = v101
	goto L43
L43:
	;
	switch v158 - int32(1) {
	case 0:
		v327 = v160
		v328 = v161
		v329 = v162
		goto L27
	case 1:
		v320 = v160
		v321 = v161
		v322 = v162
		goto L28
	case 2:
		v313 = v160
		v314 = v161
		v315 = v162
		goto L29
	case 3:
		v307 = v161
		v308 = v162
		goto L30
	case 4:
		v303 = v161
		v304 = v162
		goto L31
	case 5:
		v297 = v161
		v298 = v162
		goto L32
	case 6:
		v291 = v161
		v292 = v162
		goto L33
	case 7:
		v286 = v162
		goto L34
	case 8:
		v281 = v162
		goto L35
	case 9:
		v276 = v162
		goto L36
	case 10:
		goto L37
	default:
		v335 = v160
		v336 = v161
		v337 = v162
		goto L26
	}
L44:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v116 = v115 + v112
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v120 = v119 + v113
	v122 = int32(4)
	v124 = v117 + v111 - v120 ^ base.I32_rotl(v120, v122)
	v128 = v116 - v124 ^ base.I32_rotl(v124, int32(6))
	v129 = v120 + v116
	v130 = v124 + v129
	v131 = v128 + v130
	v135 = v129 - v128 ^ base.I32_rotl(v128, int32(8))
	v139 = v130 - v135 ^ base.I32_rotl(v135, int32(16))
	v143 = v131 - v139 ^ base.I32_rotl(v139, int32(19))
	v144 = v135 + v131
	v145 = v139 + v144
	v146 = v143 + v145
	v150 = v144 - v143 ^ base.I32_rotl(v143, v122)
	v151 = int32(12)
	v152 = v108 + v151
	v154 = v109 - v151
	if base.Ui32(int32(11)) < base.Ui32(v154) {
		v108 = v152
		v109 = v154
		v111 = v145
		v112 = v146
		v113 = v150
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v157 = v152
	v158 = v154
	v160 = v145
	v161 = v146
	v162 = v150
	goto L43
L46:
	;
	goto L45
L47:
	;
	v168 = v55
	v169 = v56
	v171 = v99
	v172 = v103
	v173 = v101
	goto L50
L48:
	;
	v217 = v55
	v218 = v56
	v220 = v99
	v221 = v103
	v222 = v101
	goto L49
L49:
	;
	switch v218 - int32(1) {
	case 0:
		v269 = v220
		goto L53
	case 1:
		v264 = v220
		goto L54
	case 2:
		goto L55
	case 3:
		v257 = v221
		goto L56
	case 4:
		v254 = v221
		goto L57
	case 5:
		v249 = v221
		goto L58
	case 6:
		goto L59
	case 7:
		v240 = v222
		goto L60
	case 8:
		v235 = v222
		goto L61
	case 9:
		v230 = v222
		goto L62
	case 10:
		goto L63
	default:
		v335 = v220
		v336 = v221
		v337 = v222
		goto L26
	}
L50:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v176 = v175 + v172
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v168)+8))
	v180 = v179 + v173
	v182 = int32(4)
	v184 = v177 + v171 - v180 ^ base.I32_rotl(v180, v182)
	v188 = v176 - v184 ^ base.I32_rotl(v184, int32(6))
	v189 = v180 + v176
	v190 = v184 + v189
	v191 = v188 + v190
	v195 = v189 - v188 ^ base.I32_rotl(v188, int32(8))
	v199 = v190 - v195 ^ base.I32_rotl(v195, int32(16))
	v203 = v191 - v199 ^ base.I32_rotl(v199, int32(19))
	v204 = v195 + v191
	v205 = v199 + v204
	v206 = v203 + v205
	v210 = v204 - v203 ^ base.I32_rotl(v203, v182)
	v211 = int32(12)
	v212 = v168 + v211
	v214 = v169 - v211
	if base.Ui32(int32(11)) < base.Ui32(v214) {
		v168 = v212
		v169 = v214
		v171 = v205
		v172 = v206
		v173 = v210
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v217 = v212
	v218 = v214
	v220 = v205
	v221 = v206
	v222 = v210
	goto L49
L52:
	;
	goto L51
L53:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217))))
	v335 = v269 + v270
	v336 = v221
	v337 = v222
	goto L26
L54:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+1)))
	v269 = v265<<(uint(int32(8))%32) + v264
	goto L53
L55:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+2)))
	v264 = v260<<(uint(int32(16))%32) + v220
	goto L54
L56:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v335 = v258 + v220
	v336 = v257
	v337 = v222
	goto L26
L57:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+4)))
	v257 = v254 + v255
	goto L56
L58:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+5)))
	v254 = v250<<(uint(int32(8))%32) + v249
	goto L57
L59:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+6)))
	v249 = v245<<(uint(int32(16))%32) + v221
	goto L58
L60:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v335 = v241 + v220
	v336 = v243 + v221
	v337 = v240
	goto L26
L61:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+8)))
	v240 = v236<<(uint(int32(8))%32) + v235
	goto L60
L62:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+9)))
	v235 = v231<<(uint(int32(16))%32) + v230
	goto L61
L63:
	;
	v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+10)))
	v230 = v226<<(uint(int32(24))%32) + v222
	goto L62
L64:
	;
	v408 = v375
	goto L15
L65:
	;
	v385 = F_DirectFunctionCall2Coll(m, int32(1471), int32(0), base.I64_extend_i32_u(v377)&int64(1), v17)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	if v377&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v408 = v385
	goto L15
L69:
	;
	v391 = int64(2)
	goto L71
L70:
	;
	v391 = int64(4)
	goto L71
L71:
	;
	v408 = v391
	goto L15
L72:
	;
	F_errmsg_internal(m, int32(_a_F_jsonb_hash_extended_0), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_jsonb_hash_extended_1), int32(1516), int32(_a_F_jsonb_hash_extended_2))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v40
	F_errmsg_internal(m, int32(_a_F_jsonb_hash_extended_3), v10)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_jsonb_hash_extended_4), int32(329), int32(_a_F_jsonb_hash_extended_5))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_pfree(m, v13)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	v444 = v438
	goto L5
L81:
	;
	goto L80
}
func F_jsonb_in_array_start(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	F_pushJsonbValue(m, l0, int32(4), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_jsonb_object_keys(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	goto L3
L3:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	goto L42
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_jsonb_object_keys[0])) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v37
	goto L3
L5:
	;
	v103 = v99
	goto L28
L6:
	;
	v99 = int32(1)
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L24
	}
L8:
	;
	return int64(0)
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22&int32(268435456) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v22&int32(1073741824) != 0 {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L8
	} else {
		goto L20
	}
L13:
	;
	v29 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v31 = int32(_a_F_jsonb_object_keys_0)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_object_keys[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_jsonb_object_keys[0])) = v34
	v37 = F_palloc(m, int32(20))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v39 = int32(4)
	v40 = v18 + v39
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v37)+12)) = int64(0)
	v45 = v41 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v45
	v48 = F_palloc_mul(m, v39, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v48
	v51 = F_JsonbIteratorInit(m, v40)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v51
	v59 = F_JsonbIteratorNext(m, v9+int32(-4), v9+int32(-40), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v99 = int32(0)
	goto L5
L19:
	;
	switch v59 {
	case 0:
		goto L4
	case 1:
		goto L18
	default:
		goto L6
	}
L20:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_jsonb_object_keys_1)
	F_errmsg(m, int32(_a_F_jsonb_object_keys_2), v9+int32(-48))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_keys_3), int32(589), int32(_a_F_jsonb_object_keys_1))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_jsonb_object_keys_1)
	F_errmsg(m, int32(_a_F_jsonb_object_keys_4), v11)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_keys_3), int32(594), int32(_a_F_jsonb_object_keys_1))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	if v103 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v113 = F_palloc(m, v110+int32(1))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L8
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L37
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	if v115 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)+36))
	base.MemoryCopy(m, v113, v116, v115)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v120 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v113+v118))) = uint8(v120)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v123 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v122 + v123
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v126+v122<<(uint(int32(2))%32)))) = v113
	v103 = v123
	goto L28
L37:
	;
	v145 = F_JsonbIteratorNext(m, v9+int32(-4), v9+int32(-40), int32(1))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L8
	} else {
		goto L40
	}
L38:
	;
	v103 = int32(0)
	goto L28
L39:
	;
	goto L38
L40:
	;
	switch v145 {
	case 0:
		goto L4
	case 1:
		goto L39
	default:
		goto L37
	}
L41:
	;
	m.G0 = v11 - int32(-64)
	return v199
L42:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+16))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	if v170 < v171 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v173 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+16)) = v170 + v173
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v170<<(uint(int32(2))%32))))
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = v181 + int64(1)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+20)) = v173
	v188 = F_cstring_to_text(m, v180)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L8
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	v199 = base.I64_extend_i32_u(v188)
	goto L41
L47:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+20)) = int32(2)
	v196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v196)
	v199 = int64(0)
	goto L41
}
func F_jsonb_out(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v14 = F_JsonbToCStringWorker(m, int32(0), v4+int32(4), int32(base.Ui32(v10)>>(uint(int32(2))%32)), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v14)
		}
	}
}
func F_jsonb_path_exists(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_exists_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_path_exists_internal(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
			if v15 == int32(4) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v19 = F_pg_detoast_datum(m, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
					v24 = base.B2i32(v21 == int64(0))
					v25 = v19
					v29 = F_executeJsonPath(m, v13, v25, int32(1534), int32(1535), v8, v24, int32(0), l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						if v31 != v8 {
							F_pfree(m, v8)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int64(0)
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v35 != v13 {
									F_pfree(m, v13)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int64(0)
									} else {
										if v29 == int32(2) {
											v41 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
											return int64(0)
										} else {
											return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
										}
									}
								} else {
									if v29 == int32(2) {
										v41 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
										return int64(0)
									} else {
										return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
									}
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v35 != v13 {
								F_pfree(m, v13)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int64(0)
								} else {
									if v29 == int32(2) {
										v41 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
										return int64(0)
									} else {
										return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
									}
								}
							} else {
								if v29 == int32(2) {
									v41 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
									return int64(0)
								} else {
									return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
								}
							}
						}
					}
				}
			} else {
				v24 = v3
				v25 = v3
				v29 = F_executeJsonPath(m, v13, v25, int32(1534), int32(1535), v8, v24, int32(0), l1)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v31 != v8 {
						F_pfree(m, v8)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
							if v35 != v13 {
								F_pfree(m, v13)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int64(0)
								} else {
									if v29 == int32(2) {
										v41 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
										return int64(0)
									} else {
										return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
									}
								}
							} else {
								if v29 == int32(2) {
									v41 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
									return int64(0)
								} else {
									return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v35 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								if v29 == int32(2) {
									v41 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
									return int64(0)
								} else {
									return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
								}
							}
						} else {
							if v29 == int32(2) {
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
								return int64(0)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v29 == int32(0)))
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_path_match_tz(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_match_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_path_query_array(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_array_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_jsonb_path_query_internal(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
	if v11 == int32(0) {
		v14 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			v18 = int32(_a_F_jsonb_path_query_internal_0)
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_path_query_internal[0]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_jsonb_path_query_internal[0])) = v21
			v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v24 = F_pg_detoast_datum_copy(m, v23)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v27 = F_pg_detoast_datum_copy(m, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v30 = F_pg_detoast_datum_copy(m, v29)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
						v34 = F_palloc(m, int32(80))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v34)+8)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v34))) = int64(8589934592)
							*(*int32)(unsafe.Add(mBase, uint32(v34)+12)) = v34
							v45 = F_executeJsonPath(m, v27, v30, int32(1534), int32(1535), v24, base.B2i32(v32 == int64(0)), v34, l1)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								v48 = F_palloc(m, int32(8))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v48))) = v34
									*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v48
									*(*int32)(unsafe.Add(mBase, _c_F_jsonb_path_query_internal[0])) = v19
									v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
									if v65 != 0 {
										v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
										if v66 < v67 {
											v83 = v65
											v84 = v66
											v85 = int32(1)
											*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v84 + v85
											v88 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
											*(*int64)(unsafe.Add(mBase, uint32(v63))) = v88 + int64(1)
											v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v85
											v100 = F_JsonbValueToJsonb(m, v83+v84<<(uint(int32(5))%32)+int32(16))
											mBase = m.M
											v101 = m.ExcPending
											if v101 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v100)
											}
										} else {
											v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v64))) = v69
											if v69 != 0 {
												v83 = v69
												v84 = int32(0)
												v85 = int32(1)
												*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v84 + v85
												v88 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
												*(*int64)(unsafe.Add(mBase, uint32(v63))) = v88 + int64(1)
												v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v85
												v100 = F_JsonbValueToJsonb(m, v83+v84<<(uint(int32(5))%32)+int32(16))
												mBase = m.M
												v101 = m.ExcPending
												if v101 != 0 {
													return int64(0)
												} else {
													return base.I64_extend_i32_u(v100)
												}
											} else {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int64(0)
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
													v79 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
													return int64(0)
												}
											}
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int64(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
											v79 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
											return int64(0)
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
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
		v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
		if v65 != 0 {
			v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
			if v66 < v67 {
				v83 = v65
				v84 = v66
				v85 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v84 + v85
				v88 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
				*(*int64)(unsafe.Add(mBase, uint32(v63))) = v88 + int64(1)
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v85
				v100 = F_JsonbValueToJsonb(m, v83+v84<<(uint(int32(5))%32)+int32(16))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(v100)
				}
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v64))) = v69
				if v69 != 0 {
					v83 = v69
					v84 = int32(0)
					v85 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = v84 + v85
					v88 = *(*int64)(unsafe.Add(mBase, uint32(v63)))
					*(*int64)(unsafe.Add(mBase, uint32(v63))) = v88 + int64(1)
					v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = v85
					v100 = F_JsonbValueToJsonb(m, v83+v84<<(uint(int32(5))%32)+int32(16))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v100)
					}
				} else {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
						v79 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
						return int64(0)
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int64(0)
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
				v79 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
				return int64(0)
			}
		}
	}
}
func F_jsonb_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pq_getmsgint(m, v10, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v23 = F_pq_getmsgtext(m, v10, v18-v19, v8+int32(12))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v26 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v26
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v30
				*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v30
				v37 = v8 + int32(76)
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_recv[0]))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
				v42 = F_makeJsonLexContextCstringLen(m, v37, v23, v25, v40, int32(1))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v44 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v44
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+72)) = uint8(v44)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(1451)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(1452)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(1453)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(1454)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(1455)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(1456)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(56)
					v66 = F_pg_parse_json_or_errsave(m, v37, v8+int32(16), v44)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int64(0)
					} else {
						if v66 != 0 {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+56))
							v69 = F_JsonbValueToJsonb(m, v68)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int64(0)
							} else {
								v72 = base.I64_extend_i32_u(v69)
								m.G0 = v8 + int32(144)
								return v72
							}
						} else {
							v72 = int64(0)
							m.G0 = v8 + int32(144)
							return v72
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
				F_errmsg_internal(m, int32(_a_F_jsonb_recv_0), v8)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_jsonb_recv_1), int32(90), int32(_a_F_jsonb_recv_2))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
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
func F_jsonb_set_lax(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int64
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	if v13 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v9 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v9)
	return int64(0)
L3:
	;
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v5 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	if v6 != int32(1) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L14
	} else {
		goto L68
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L14
	} else {
		goto L62
	}
L8:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	if v16 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L14
	} else {
		goto L58
	}
L11:
	;
	v19 = F_jsonb_set(m, l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return int64(0)
L15:
	;
	return v19
L16:
	;
	v27 = F_text_to_cstring(m, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v29 = int32(_a_F_jsonb_set_lax_0)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_jsonb_set_lax[0])))
	if base.B2i32(v32 == int32(0))|base.B2i32(v32 != v35) != 0 {
		v53 = v32
		v54 = v35
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v53-v54 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L19:
	;
	goto L18
L20:
	;
	v38 = v27
	v39 = v29
	goto L21
L21:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	if v43 == int32(0) {
		v53 = v43
		v54 = v42
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v53 = v43
	v54 = v42
	goto L19
L23:
	;
	v46 = int32(1)
	if v43 == v42 {
		v38 = v38 + v46
		v39 = v39 + v46
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v58 = int32(_a_F_jsonb_set_lax_1)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_jsonb_set_lax[1])))
	if base.B2i32(v61 == int32(0))|base.B2i32(v61 != v64) != 0 {
		v82 = v61
		v83 = v64
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v82-v83 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	v67 = v27
	v68 = v58
	goto L29
L29:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v72 == int32(0) {
		v82 = v72
		v83 = v71
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v82 = v72
	v83 = v71
	goto L27
L31:
	;
	v75 = int32(1)
	if v72 == v71 {
		v67 = v67 + v75
		v68 = v68 + v75
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v90 = F_DirectFunctionCall1Coll(m, int32(521), int32(0), int64(327844))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L14
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v98 = int32(_a_F_jsonb_set_lax_2)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v104 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_jsonb_set_lax[2])))
	if base.B2i32(v101 == int32(0))|base.B2i32(v101 != v104) != 0 {
		v122 = v101
		v123 = v104
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v92 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v92)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v90
	v95 = F_jsonb_set(m, l0)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	return v95
L38:
	;
	if v122-v123 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	goto L38
L40:
	;
	v107 = v27
	v108 = v98
	goto L41
L41:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	if v112 == int32(0) {
		v122 = v112
		v123 = v111
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v122 = v112
	v123 = v111
	goto L39
L43:
	;
	v115 = int32(1)
	if v112 == v111 {
		v107 = v107 + v115
		v108 = v108 + v115
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v127 = F_jsonb_delete_path(m, l0)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L14
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v130 = int32(_a_F_jsonb_set_lax_3)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_jsonb_set_lax[3])))
	if base.B2i32(v133 == int32(0))|base.B2i32(v133 != v136) != 0 {
		v154 = v133
		v155 = v136
		goto L50
	} else {
		goto L51
	}
L48:
	;
	return v127
L49:
	;
	if v154-v155 != 0 {
		goto L6
	} else {
		goto L56
	}
L50:
	;
	goto L49
L51:
	;
	v139 = v27
	v140 = v130
	goto L52
L52:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+1)))
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)))
	if v144 == int32(0) {
		v154 = v144
		v155 = v143
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v154 = v144
	v155 = v143
	goto L50
L54:
	;
	v147 = int32(1)
	if v144 == v143 {
		v139 = v139 + v147
		v140 = v140 + v147
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v158 = F_pg_detoast_datum(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	return base.I64_extend_i32_u(v158)
L58:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_jsonb_set_lax_4), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_jsonb_set_lax_5), int32(_a_F_jsonb_set_lax_6), int32(_a_F_jsonb_set_lax_7))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L14
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_jsonb_set_lax_8), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L14
	} else {
		goto L64
	}
L64:
	;
	v191 = F_errdetail(m, int32(_a_F_jsonb_set_lax_9), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	F_errhint(m, int32(_a_F_jsonb_set_lax_10), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_jsonb_set_lax_5), int32(_a_F_jsonb_set_lax_11), int32(_a_F_jsonb_set_lax_7))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L14
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L14
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_jsonb_set_lax_4), int32(0))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L14
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_jsonb_set_lax_5), int32(_a_F_jsonb_set_lax_12), int32(_a_F_jsonb_set_lax_7))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L14
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_jsonb_string_to_tsvector(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = F_getTSCurrentConfig(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v14
			v17 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			v22 = v7 + int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v22
			F_iterate_jsonb_values(m, v10, int32(2), v7+int32(24))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				v29 = F_make_tsvector(m, v22)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v31 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int64(0)
						} else {
							m.G0 = v7 + int32(32)
							return base.I64_extend_i32_u(v29)
						}
					} else {
						m.G0 = v7 + int32(32)
						return base.I64_extend_i32_u(v29)
					}
				}
			}
		}
	}
}
