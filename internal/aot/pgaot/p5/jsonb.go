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
	var v191 int32
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
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
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
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
			m.G0 = v14 + int32(144)
			return
		}
	case 2:
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
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
			v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
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
		v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
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
				F_errfinish(m, int32(_a_F_convertJsonbScalar_1), int32(1876), int32(_a_F_convertJsonbScalar_2))
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
		v191 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v192 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v195 = F_JsonEncodeDateTime(m, v14, v191, v192, l2+int32(16))
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
func F_get_jsonb_path_all(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v16 = F_pg_detoast_datum(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = F_array_contains_nulls(m, v16)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					F_deconstruct_array_builtin(m, v16, int32(25), v8+int32(12), v8+int32(8), v8)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v33 = F_jsonb_get_element(m, v11, v29, v30, v8+int32(7), l1)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+7)))
							if v35 != int32(1) {
								v42 = v33
							} else {
								v39 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
								v42 = int32(0)
							}
							m.G0 = v8 + int32(16)
							return v42
						}
					}
				} else {
					v39 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v39)
					v42 = int32(0)
					m.G0 = v8 + int32(16)
					return v42
				}
			}
		}
	}
}
func F_jsonb_agg_transfn(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_agg_transfn_worker(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_array_length(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		if v7&int32(268435456) == int32(0) {
			if v7&int32(1073741824) == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_jsonb_array_length_0), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_jsonb_array_length_1), int32(1889), int32(_a_F_jsonb_array_length_2))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
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
				return v7 & int32(268435455)
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_jsonb_array_length_3), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_jsonb_array_length_1), int32(1885), int32(_a_F_jsonb_array_length_2))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
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
func F_jsonb_build_array(m *base.Module, l0 int32) int32 {
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
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v18 = F_extract_variadic_args(m, l0, v10+int32(12), v10+int32(4), v10+int32(8))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v10 + int32(32)
	return v87
L2:
	;
	return int32(0)
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
	v87 = int32(0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v30
	v38 = F_pushJsonbValue(m, v10+int32(16), int32(4), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v38
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v43 = int32(0)
	goto L11
L9:
	;
	goto L10
L10:
	;
	v75 = F_pushJsonbValue(m, v10+int32(16), int32(5), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L15
	}
L11:
	;
	v49 = v43 << (uint(int32(2)) % 32)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v29+v49)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v28))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v27+v49)))
	F_add_jsonb(m, v51, v53, v10+int32(16), v57, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L2
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v62 = v43 + int32(1)
	if v62 != v18 {
		v43 = v62
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v75
	v78 = F_JsonbValueToJsonb(m, v75)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v87 = v78
	goto L1
}
func F_jsonb_build_object(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = F_extract_variadic_args(m, l0, v7+int32(12), v7+int32(4), v7+int32(8))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 < int32(0) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v30 = int32(0)
			m.G0 = v7 + int32(16)
			return v30
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v26 = int32(0)
			v28 = F_jsonb_build_object_worker(m, v15, v23, v24, v25, v26, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = v28
				m.G0 = v7 + int32(16)
				return v30
			}
		}
	}
}
func F_jsonb_concat(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if (v20^v21)&int32(536870912) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v8 - int32(-64)
	return v360
L5:
	;
	v39 = F_JsonbIteratorInit(m, v11+int32(4))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	if v20&int32(268435456)|v21&int32(268435455) == int32(0) {
		v360 = v16
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if v20&int32(268435455)|v21&int32(268435456) != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	v360 = v11
	goto L4
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v39
	v44 = F_JsonbIteratorInit(m, v16+int32(4))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v44
	v48 = v6 + int32(-48)
	v50 = v6 + int32(-20)
	v52 = F_JsonbIteratorNext(m, v48, v50, int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v358 = F_JsonbValueToJsonb(m, v355)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L105
	}
L12:
	;
	v329 = F_JsonbIteratorNext(m, v6+int32(-52), v6+int32(-40), int32(1))
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L95
	}
L13:
	;
	v317 = F_pushJsonbValue(m, v6+int32(-44), int32(5), int32(0))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L94
	}
L14:
	;
	v285 = F_JsonbIteratorNext(m, v6+int32(-52), v6+int32(-40), int32(1))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L87
	}
L15:
	;
	v257 = v98
	goto L82
L16:
	;
	v239 = v76
	goto L77
L17:
	;
	v61 = F_JsonbIteratorNext(m, v6+int32(-52), v6+int32(-40), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if base.B2i32(v52 != int32(6))|base.B2i32(v61 != int32(6)) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v68 = int32(0)
	v73 = F_pushJsonbValue(m, v6+int32(-44), int32(6), v68)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v84 = F_pushJsonbValue(m, v6+int32(-44), int32(4), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L25
	}
L22:
	;
	v76 = F_JsonbIteratorNext(m, v48, v50, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v76 != int32(7) {
		goto L16
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	v86 = int32(4)
	if base.B2i32(v52 != v86)|base.B2i32(v61 != v86) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v98 = F_JsonbIteratorNext(m, v6+int32(-48), v6+int32(-20), int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v52 == int32(6) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	if v98 != int32(5) {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	goto L14
L31:
	;
	v104 = int32(0)
	v109 = F_pushJsonbValue(m, v6+int32(-44), int32(6), v104)
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
	v177 = F_JsonbIteratorNext(m, v6+int32(-48), v6+int32(-20), int32(1))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L57
	}
L34:
	;
	v116 = F_JsonbIteratorNext(m, v6+int32(-48), v6+int32(-20), int32(1))
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
	v148 = F_JsonbIteratorNext(m, v6+int32(-52), v6+int32(-40), int32(1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L47
	}
L39:
	;
	v126 = v6 + int32(-20)
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
	v131 = F_pushJsonbValue(m, v6+int32(-44), v118, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v136 = F_JsonbIteratorNext(m, v6+int32(-48), v126, int32(1))
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
	if v148 == int32(0) {
		v355 = v104
		goto L11
	} else {
		goto L48
	}
L48:
	;
	v152 = v148
	goto L49
L49:
	;
	v160 = v6 + int32(-40)
	if v152 != int32(5) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v355 = v165
	goto L11
L51:
	;
	v164 = v160
	goto L53
L52:
	;
	v164 = int32(0)
	goto L53
L53:
	;
	v165 = F_pushJsonbValue(m, v6+int32(-44), v152, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v170 = F_JsonbIteratorNext(m, v6+int32(-52), v160, int32(1))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	if v170 != 0 {
		v152 = v170
		goto L49
	} else {
		goto L56
	}
L56:
	;
	goto L50
L57:
	;
	if v177 != int32(5) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v181 = v177
	goto L61
L59:
	;
	goto L60
L60:
	;
	v208 = F_pushJsonbValue(m, v6+int32(-44), int32(6), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L66
	}
L61:
	;
	v189 = v6 + int32(-20)
	v190 = F_pushJsonbValue(m, v6+int32(-44), v181, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L63
	}
L62:
	;
	goto L60
L63:
	;
	v195 = F_JsonbIteratorNext(m, v6+int32(-48), v189, int32(1))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	if v195 != int32(5) {
		v181 = v195
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	v215 = F_JsonbIteratorNext(m, v6+int32(-52), v6+int32(-40), int32(1))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v215 == int32(0) {
		goto L13
	} else {
		goto L68
	}
L68:
	;
	v219 = v215
	goto L69
L69:
	;
	v227 = v6 + int32(-40)
	if v219 != int32(7) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	goto L13
L71:
	;
	v231 = v227
	goto L73
L72:
	;
	v231 = int32(0)
	goto L73
L73:
	;
	v232 = F_pushJsonbValue(m, v6+int32(-44), v219, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v237 = F_JsonbIteratorNext(m, v6+int32(-52), v227, int32(1))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	if v237 != 0 {
		v219 = v237
		goto L69
	} else {
		goto L76
	}
L76:
	;
	goto L70
L77:
	;
	v247 = v6 + int32(-20)
	v248 = F_pushJsonbValue(m, v6+int32(-44), v239, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L12
L79:
	;
	v253 = F_JsonbIteratorNext(m, v6+int32(-48), v247, int32(1))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v253 != int32(7) {
		v239 = v253
		goto L77
	} else {
		goto L81
	}
L81:
	;
	goto L78
L82:
	;
	v265 = v6 + int32(-20)
	v266 = F_pushJsonbValue(m, v6+int32(-44), v257, v265)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L84
	}
L83:
	;
	goto L14
L84:
	;
	v271 = F_JsonbIteratorNext(m, v6+int32(-48), v265, int32(1))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v271 != int32(5) {
		v257 = v271
		goto L82
	} else {
		goto L86
	}
L86:
	;
	goto L83
L87:
	;
	if v285 == int32(5) {
		goto L13
	} else {
		goto L88
	}
L88:
	;
	goto L89
L89:
	;
	v298 = v6 + int32(-40)
	v299 = F_pushJsonbValue(m, v6+int32(-44), int32(3), v298)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	goto L13
L91:
	;
	v304 = F_JsonbIteratorNext(m, v6+int32(-52), v298, int32(1))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v304 != int32(5) {
		goto L89
	} else {
		goto L93
	}
L93:
	;
	goto L90
L94:
	;
	v355 = v317
	goto L11
L95:
	;
	if v329 == int32(0) {
		v355 = v68
		goto L11
	} else {
		goto L96
	}
L96:
	;
	v333 = v329
	goto L97
L97:
	;
	v341 = v6 + int32(-40)
	if v333 != int32(7) {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v355 = v346
	goto L11
L99:
	;
	v345 = v341
	goto L101
L100:
	;
	v345 = int32(0)
	goto L101
L101:
	;
	v346 = F_pushJsonbValue(m, v6+int32(-44), v333, v345)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	v351 = F_JsonbIteratorNext(m, v6+int32(-52), v341, int32(1))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v351 != 0 {
		v333 = v351
		goto L97
	} else {
		goto L104
	}
L104:
	;
	goto L98
L105:
	;
	v360 = v358
	goto L4
}
func F_jsonb_exists(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v19
			v22 = v17 + v19
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v27 = v25 & v19
			if v27 != 0 {
				v28 = v22
			} else {
				v28 = v17 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v28
			if v25 == int32(1) {
				v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				if v35 == int32(18) {
					v38 = int32(16)
				} else {
					v38 = int32(0)
				}
				if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v45 = int32(4)
				} else {
					v45 = v38
				}
				v56 = v45
			} else {
				v46 = int32(1)
				if v27 != 0 {
					v56 = int32(base.Ui32(v25)>>(uint(v46)%32)) - v46
				} else {
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v56
			v63 = F_findJsonbValueFromContainer(m, v12+int32(4), int32(1610612736), v9+int32(12))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(32)
				return base.B2i32(v63 != int32(0))
			}
		}
	}
}
func F_jsonb_exists_all(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	F_deconstruct_array_builtin(m, v16, int32(25), v8+int32(28), v8+int32(24), v8+int32(20))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	if v28 <= int32(0) {
		v106 = int32(1)
		goto L5
	} else {
		goto L6
	}
L5:
	;
	m.G0 = v8 + int32(32)
	return v106
L6:
	;
	v34 = int32(0)
	v37 = v28
	goto L7
L7:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v34))))
	if v41 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v106 = int32(0)
	goto L5
L9:
	;
	goto L8
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v45 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v45
	v49 = v44 + v34<<(uint(int32(2))%32)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v53&v45 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v98 = v37
	goto L12
L12:
	;
	v99 = int32(1)
	v101 = v34 + v99
	if v101 < v98 {
		v34 = v101
		v37 = v98
		goto L7
	} else {
		goto L29
	}
L13:
	;
	v56 = v45
	goto L15
L14:
	;
	v56 = int32(4)
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v50 + v56
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v60 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v89
	v92 = F_findJsonbValueFromContainer(m, v11+int32(4), int32(1610612736), v8)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L27
	}
L17:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+1)))
	if v66 == int32(18) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v77 = int32(1)
	if v60&v77 != 0 {
		v89 = int32(base.Ui32(v60)>>(uint(v77)%32)) - v77
		goto L16
	} else {
		goto L26
	}
L20:
	;
	v69 = int32(16)
	goto L22
L21:
	;
	v69 = int32(0)
	goto L22
L22:
	;
	if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v76 = int32(4)
	goto L25
L24:
	;
	v76 = v69
	goto L25
L25:
	;
	v89 = v76
	goto L16
L26:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L27:
	;
	if v92 == int32(0) {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
	v98 = v96
	goto L12
L29:
	;
	v106 = v99
	goto L5
}
func F_jsonb_hash_extended(m *base.Module, l0 int32) int32 {
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
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
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
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
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
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int64
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int64
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int64
	_ = v397
	var v400 int64
	_ = v400
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v417 int64
	_ = v417
	var v418 int64
	_ = v418
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int64
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21&int32(268435455) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v10 + int32(48)
	return v457
L4:
	;
	v26 = F_Int64GetDatum(m, v18)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v30 = F_JsonbIteratorInit(m, v13+int32(4))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v457 = v26
	goto L3
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v30
	goto L10
L9:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v443 != v13 {
		goto L83
	} else {
		goto L84
	}
L10:
	;
	v45 = F_JsonbIteratorNext(m, v10+int32(44), v10+int32(24), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L16
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L80
	}
L12:
	;
	goto L11
L13:
	;
	v56 = v10 + int32(16)
	v58 = v10 + int32(24)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	switch v59 {
	case 0:
		goto L18
	case 1:
		goto L22
	case 2:
		goto L21
	case 3:
		goto L20
	default:
		goto L19
	}
L14:
	;
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v51 ^ int64(2305843009750564864)
	goto L10
L15:
	;
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v47 ^ int64(4611686019501129728)
	goto L10
L16:
	;
	switch v45 {
	case 0:
		goto L9
	case 1, 2, 3:
		goto L13
	case 4:
		goto L15
	case 5, 7:
		goto L10
	case 6:
		goto L14
	default:
		goto L12
	}
L17:
	;
	v418 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
	*(*int64)(unsafe.Add(mBase, uint32(v56))) = v417 ^ (v418<<(uint(int64(1))%64)&int64(-4294967298) | int64(base.Ui64(v418)>>(uint(int64(31))%64))&int64(4294967297))
	goto L10
L18:
	;
	v417 = v18 + int64(1)
	goto L17
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L77
	}
L20:
	;
	v388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)))
	if v18 != int64(0) {
		goto L69
	} else {
		goto L70
	}
L21:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v383 = F_Int64GetDatum(m, v18)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L67
	}
L22:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v67 = v61 - int32(1636608432)
	if v18 == int64(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v377 = F_Int64GetDatum(m, base.I64_extend_i32_u(v367)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v367^v359-base.I32_rotl(v367, int32(24))))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L66
	}
L24:
	;
	if v60&int32(3) != 0 {
		goto L40
	} else {
		goto L41
	}
L25:
	;
	v104 = v67
	v106 = v67
	v108 = v67
	goto L24
L26:
	;
	goto L27
L27:
	;
	v71 = v67 + base.I32_wrap_i64(v18)
	v72 = v71 + v67
	v76 = int32(4)
	v78 = base.I32_wrap_i64(int64(base.Ui64(v18)>>(uint(int64(32))%64))) ^ base.I32_rotl(v67, v76)
	v82 = v71 - v78 ^ base.I32_rotl(v78, int32(6))
	v86 = v72 - v82 ^ base.I32_rotl(v82, int32(8))
	v87 = v72 + v78
	v88 = v82 + v87
	v89 = v86 + v88
	v93 = v87 - v86 ^ base.I32_rotl(v86, int32(16))
	v97 = v88 - v93 ^ base.I32_rotl(v93, int32(19))
	v102 = v89 + v93
	v104 = v102
	v106 = v89 - v97 ^ base.I32_rotl(v97, v76)
	v108 = v97 + v102
	goto L24
L28:
	;
	v345 = int32(14)
	v347 = v341 ^ v342 - base.I32_rotl(v341, v345)
	v351 = v347 ^ v340 - base.I32_rotl(v347, int32(11))
	v355 = v351 ^ v341 - base.I32_rotl(v351, int32(25))
	v359 = v355 ^ v347 - base.I32_rotl(v355, int32(16))
	v363 = v359 ^ v351 - base.I32_rotl(v359, int32(4))
	v367 = v363 ^ v355 - base.I32_rotl(v363, v345)
	goto L23
L29:
	;
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162))))
	v340 = v332 + v335
	v341 = v333
	v342 = v334
	goto L28
L30:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	v332 = v328<<(uint(int32(8))%32) + v325
	v333 = v326
	v334 = v327
	goto L29
L31:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+2)))
	v325 = v321<<(uint(int32(16))%32) + v318
	v326 = v319
	v327 = v320
	goto L30
L32:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+3)))
	v318 = v314<<(uint(int32(24))%32) + v165
	v319 = v312
	v320 = v313
	goto L31
L33:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+4)))
	v312 = v308 + v310
	v313 = v309
	goto L32
L34:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+5)))
	v308 = v304<<(uint(int32(8))%32) + v302
	v309 = v303
	goto L33
L35:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+6)))
	v302 = v298<<(uint(int32(16))%32) + v296
	v303 = v297
	goto L34
L36:
	;
	v292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+7)))
	v296 = v292<<(uint(int32(24))%32) + v166
	v297 = v291
	goto L35
L37:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+8)))
	v291 = v287<<(uint(int32(8))%32) + v286
	goto L36
L38:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+9)))
	v286 = v282<<(uint(int32(16))%32) + v281
	goto L37
L39:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+10)))
	v281 = v277<<(uint(int32(24))%32) + v167
	goto L38
L40:
	;
	if base.Ui32(int32(11)) < base.Ui32(v61) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if base.Ui32(int32(12)) <= base.Ui32(v61) {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	v113 = v60
	v114 = v61
	v116 = v104
	v117 = v108
	v118 = v106
	goto L46
L44:
	;
	v162 = v60
	v163 = v61
	v165 = v104
	v166 = v108
	v167 = v106
	goto L45
L45:
	;
	switch v163 - int32(1) {
	case 0:
		v332 = v165
		v333 = v166
		v334 = v167
		goto L29
	case 1:
		v325 = v165
		v326 = v166
		v327 = v167
		goto L30
	case 2:
		v318 = v165
		v319 = v166
		v320 = v167
		goto L31
	case 3:
		v312 = v166
		v313 = v167
		goto L32
	case 4:
		v308 = v166
		v309 = v167
		goto L33
	case 5:
		v302 = v166
		v303 = v167
		goto L34
	case 6:
		v296 = v166
		v297 = v167
		goto L35
	case 7:
		v291 = v167
		goto L36
	case 8:
		v286 = v167
		goto L37
	case 9:
		v281 = v167
		goto L38
	case 10:
		goto L39
	default:
		v340 = v165
		v341 = v166
		v342 = v167
		goto L28
	}
L46:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v121 = v120 + v117
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v125 = v124 + v118
	v127 = int32(4)
	v129 = v122 + v116 - v125 ^ base.I32_rotl(v125, v127)
	v133 = v121 - v129 ^ base.I32_rotl(v129, int32(6))
	v134 = v125 + v121
	v135 = v129 + v134
	v136 = v133 + v135
	v140 = v134 - v133 ^ base.I32_rotl(v133, int32(8))
	v144 = v135 - v140 ^ base.I32_rotl(v140, int32(16))
	v148 = v136 - v144 ^ base.I32_rotl(v144, int32(19))
	v149 = v140 + v136
	v150 = v144 + v149
	v151 = v148 + v150
	v155 = v149 - v148 ^ base.I32_rotl(v148, v127)
	v156 = int32(12)
	v157 = v113 + v156
	v159 = v114 - v156
	if base.Ui32(int32(11)) < base.Ui32(v159) {
		v113 = v157
		v114 = v159
		v116 = v150
		v117 = v151
		v118 = v155
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v162 = v157
	v163 = v159
	v165 = v150
	v166 = v151
	v167 = v155
	goto L45
L48:
	;
	goto L47
L49:
	;
	v173 = v60
	v174 = v61
	v176 = v104
	v177 = v108
	v178 = v106
	goto L52
L50:
	;
	v222 = v60
	v223 = v61
	v225 = v104
	v226 = v108
	v227 = v106
	goto L51
L51:
	;
	switch v223 - int32(1) {
	case 0:
		v274 = v225
		goto L55
	case 1:
		v269 = v225
		goto L56
	case 2:
		goto L57
	case 3:
		v262 = v226
		goto L58
	case 4:
		v259 = v226
		goto L59
	case 5:
		v254 = v226
		goto L60
	case 6:
		goto L61
	case 7:
		v245 = v227
		goto L62
	case 8:
		v240 = v227
		goto L63
	case 9:
		v235 = v227
		goto L64
	case 10:
		goto L65
	default:
		v340 = v225
		v341 = v226
		v342 = v227
		goto L28
	}
L52:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	v181 = v180 + v177
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v173)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v173)+8))
	v185 = v184 + v178
	v187 = int32(4)
	v189 = v182 + v176 - v185 ^ base.I32_rotl(v185, v187)
	v193 = v181 - v189 ^ base.I32_rotl(v189, int32(6))
	v194 = v185 + v181
	v195 = v189 + v194
	v196 = v193 + v195
	v200 = v194 - v193 ^ base.I32_rotl(v193, int32(8))
	v204 = v195 - v200 ^ base.I32_rotl(v200, int32(16))
	v208 = v196 - v204 ^ base.I32_rotl(v204, int32(19))
	v209 = v200 + v196
	v210 = v204 + v209
	v211 = v208 + v210
	v215 = v209 - v208 ^ base.I32_rotl(v208, v187)
	v216 = int32(12)
	v217 = v173 + v216
	v219 = v174 - v216
	if base.Ui32(int32(11)) < base.Ui32(v219) {
		v173 = v217
		v174 = v219
		v176 = v210
		v177 = v211
		v178 = v215
		goto L52
	} else {
		goto L54
	}
L53:
	;
	v222 = v217
	v223 = v219
	v225 = v210
	v226 = v211
	v227 = v215
	goto L51
L54:
	;
	goto L53
L55:
	;
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222))))
	v340 = v274 + v275
	v341 = v226
	v342 = v227
	goto L28
L56:
	;
	v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	v274 = v270<<(uint(int32(8))%32) + v269
	goto L55
L57:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+2)))
	v269 = v265<<(uint(int32(16))%32) + v225
	goto L56
L58:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v340 = v263 + v225
	v341 = v262
	v342 = v227
	goto L28
L59:
	;
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+4)))
	v262 = v259 + v260
	goto L58
L60:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+5)))
	v259 = v255<<(uint(int32(8))%32) + v254
	goto L59
L61:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+6)))
	v254 = v250<<(uint(int32(16))%32) + v226
	goto L60
L62:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v222)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	v340 = v246 + v225
	v341 = v248 + v226
	v342 = v245
	goto L28
L63:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+8)))
	v245 = v241<<(uint(int32(8))%32) + v240
	goto L62
L64:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+9)))
	v240 = v236<<(uint(int32(16))%32) + v235
	goto L63
L65:
	;
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+10)))
	v235 = v231<<(uint(int32(24))%32) + v227
	goto L64
L66:
	;
	v379 = *(*int64)(unsafe.Add(mBase, uint32(v377)))
	v417 = v379
	goto L17
L67:
	;
	v385 = F_DirectFunctionCall2Coll(m, int32(1330), int32(0), v382, v383)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v385)))
	v417 = v387
	goto L17
L69:
	;
	v393 = F_Int64GetDatum(m, v18)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v388 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v395 = F_DirectFunctionCall2Coll(m, int32(1331), int32(0), v388, v393)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v397 = *(*int64)(unsafe.Add(mBase, uint32(v395)))
	v417 = v397
	goto L17
L74:
	;
	v400 = int64(2)
	goto L76
L75:
	;
	v400 = int64(4)
	goto L76
L76:
	;
	v417 = v400
	goto L17
L77:
	;
	F_errmsg_internal(m, int32(_a_F_jsonb_hash_extended_0), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_jsonb_hash_extended_1), int32(1402), int32(_a_F_jsonb_hash_extended_2))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v45
	F_errmsg_internal(m, int32(_a_F_jsonb_hash_extended_3), v10)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_jsonb_hash_extended_4), int32(329), int32(_a_F_jsonb_hash_extended_5))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_pfree(m, v13)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
	v448 = F_Int64GetDatum(m, v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L87
	}
L86:
	;
	goto L85
L87:
	;
	v457 = v448
	goto L3
}
func F_jsonb_in_array_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_pushJsonbValue(m, l0, int32(4), int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v4
		return int32(0)
	}
}
func F_jsonb_object_keys(m *base.Module, l0 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
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
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int64
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v13 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	goto L3
L3:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	goto L42
L4:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_jsonb_object_keys[0])) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v36
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
	return int32(0)
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21&int32(268435456) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if v21&int32(1073741824) != 0 {
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
	v28 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v30 = int32(_a_F_jsonb_object_keys_0)
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_object_keys[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_jsonb_object_keys[0])) = v33
	v36 = F_palloc(m, int32(20))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v39 = v17 + int32(4)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int64)(unsafe.Add(mBase, uint32(v36)+12)) = int64(0)
	v44 = v40 & int32(268435455)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v44
	v48 = F_palloc(m, v44<<(uint(int32(2))%32))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v48
	v51 = F_JsonbIteratorInit(m, v39)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v51
	v59 = F_JsonbIteratorNext(m, v10+int32(44), v10+int32(24), int32(0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_jsonb_object_keys_1)
	F_errmsg(m, int32(_a_F_jsonb_object_keys_2), v10+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_keys_3), int32(586), int32(_a_F_jsonb_object_keys_1))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_jsonb_object_keys_1)
	F_errmsg(m, int32(_a_F_jsonb_object_keys_4), v10)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_jsonb_object_keys_3), int32(591), int32(_a_F_jsonb_object_keys_1))
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
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v112 = F_palloc(m, v109+int32(1))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	if v114 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
	base.MemoryCopy(m, v112, v115, v114)
	goto L36
L35:
	;
	goto L36
L36:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	v119 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112+v117))) = uint8(v119)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v121 + v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v125+v121<<(uint(int32(2))%32)))) = v112
	v103 = v122
	goto L28
L37:
	;
	v143 = F_JsonbIteratorNext(m, v10+int32(44), v10+int32(24), int32(1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
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
	switch v143 {
	case 0:
		goto L4
	case 1:
		goto L39
	default:
		goto L37
	}
L41:
	;
	m.G0 = v10 + int32(48)
	return v194
L42:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+16))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	if v166 < v167 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v169 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v166 + v169
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v172+v166<<(uint(int32(2))%32))))
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v164)))
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v177 + int64(1)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v181)+20)) = v169
	v184 = F_cstring_to_text(m, v176)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
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
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L47
	}
L46:
	;
	v194 = v184
	goto L41
L47:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+20)) = int32(2)
	v191 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v191)
	v194 = int32(0)
	goto L41
}
func F_jsonb_out(m *base.Module, l0 int32) int32 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v14 = F_JsonbToCStringWorker(m, int32(0), v4+int32(4), int32(base.Ui32(v10)>>(uint(int32(2))%32)), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			return v14
		}
	}
}
func F_jsonb_path_exists(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_exists_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_path_exists_internal(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v21 int32
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = F_pg_detoast_datum(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
			if v15 == int32(4) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v19 = F_pg_detoast_datum(m, v18)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v24 = base.B2i32(v21 == int32(0))
					v25 = v19
					v29 = F_executeJsonPath(m, v13, v25, int32(1394), int32(1395), v8, v24, int32(0), l1)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v31 != v8 {
							F_pfree(m, v8)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
								if v35 != v13 {
									F_pfree(m, v13)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return int32(0)
									} else {
										if v29 == int32(2) {
											v41 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
											return int32(0)
										} else {
											return base.B2i32(v29 == int32(0))
										}
									}
								} else {
									if v29 == int32(2) {
										v41 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
										return int32(0)
									} else {
										return base.B2i32(v29 == int32(0))
									}
								}
							}
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v35 != v13 {
								F_pfree(m, v13)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									if v29 == int32(2) {
										v41 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
										return int32(0)
									} else {
										return base.B2i32(v29 == int32(0))
									}
								}
							} else {
								if v29 == int32(2) {
									v41 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
									return int32(0)
								} else {
									return base.B2i32(v29 == int32(0))
								}
							}
						}
					}
				}
			} else {
				v24 = v3
				v25 = v3
				v29 = F_executeJsonPath(m, v13, v25, int32(1394), int32(1395), v8, v24, int32(0), l1)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v31 != v8 {
						F_pfree(m, v8)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							if v35 != v13 {
								F_pfree(m, v13)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return int32(0)
								} else {
									if v29 == int32(2) {
										v41 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
										return int32(0)
									} else {
										return base.B2i32(v29 == int32(0))
									}
								}
							} else {
								if v29 == int32(2) {
									v41 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
									return int32(0)
								} else {
									return base.B2i32(v29 == int32(0))
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v35 != v13 {
							F_pfree(m, v13)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								if v29 == int32(2) {
									v41 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
									return int32(0)
								} else {
									return base.B2i32(v29 == int32(0))
								}
							}
						} else {
							if v29 == int32(2) {
								v41 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
								return int32(0)
							} else {
								return base.B2i32(v29 == int32(0))
							}
						}
					}
				}
			}
		}
	}
}
func F_jsonb_path_match_tz(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_match_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_path_query_array(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_jsonb_path_query_array_internal(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_jsonb_path_query_internal(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v12 == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = int64(0)
		v17 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(_a_F_jsonb_path_query_internal_0)
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_path_query_internal[0]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_jsonb_path_query_internal[0])) = v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v27 = F_pg_detoast_datum_copy(m, v26)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v30 = F_pg_detoast_datum_copy(m, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v33 = F_pg_detoast_datum_copy(m, v32)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
						v42 = F_executeJsonPath(m, v30, v33, int32(1394), int32(1395), v27, base.B2i32(v37 == int32(0)), v9+int32(16), l1)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
							if v44 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v44
								*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v44
								v50 = F_list_make1_impl(m, int32(1), v9+int32(12))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v53 = v50
									*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v53
									*(*int32)(unsafe.Add(mBase, _c_F_jsonb_path_query_internal[0])) = v22
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
									if v63 != 0 {
										v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
										if v64 != 0 {
											v74 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
											v75 = F_list_delete_first(m, v63)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v75
												v78 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
												*(*int64)(unsafe.Add(mBase, uint32(v62))) = v78 + int64(1)
												v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = int32(1)
												v85 = F_JsonbValueToJsonb(m, v74)
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													v88 = v85
													m.G0 = v9 + int32(32)
													return v88
												}
											}
										} else {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = int32(2)
												v71 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
												v88 = int32(0)
												m.G0 = v9 + int32(32)
												return v88
											}
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = int32(2)
											v71 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
											v88 = int32(0)
											m.G0 = v9 + int32(32)
											return v88
										}
									}
								}
							} else {
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
								v53 = v52
								*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v53
								*(*int32)(unsafe.Add(mBase, _c_F_jsonb_path_query_internal[0])) = v22
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
								if v63 != 0 {
									v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
									if v64 != 0 {
										v74 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
										v75 = F_list_delete_first(m, v63)
										mBase = m.M
										v76 = m.ExcPending
										if v76 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v75
											v78 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
											*(*int64)(unsafe.Add(mBase, uint32(v62))) = v78 + int64(1)
											v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = int32(1)
											v85 = F_JsonbValueToJsonb(m, v74)
											mBase = m.M
											v86 = m.ExcPending
											if v86 != 0 {
												return int32(0)
											} else {
												v88 = v85
												m.G0 = v9 + int32(32)
												return v88
											}
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = int32(2)
											v71 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
											v88 = int32(0)
											m.G0 = v9 + int32(32)
											return v88
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int32(0)
									} else {
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = int32(2)
										v71 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
										v88 = int32(0)
										m.G0 = v9 + int32(32)
										return v88
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
		if v63 != 0 {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
			if v64 != 0 {
				v74 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
				v75 = F_list_delete_first(m, v63)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v75
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
					*(*int64)(unsafe.Add(mBase, uint32(v62))) = v78 + int64(1)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v82)+20)) = int32(1)
					v85 = F_JsonbValueToJsonb(m, v74)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						v88 = v85
						m.G0 = v9 + int32(32)
						return v88
					}
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = int32(2)
					v71 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
					v88 = int32(0)
					m.G0 = v9 + int32(32)
					return v88
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = int32(2)
				v71 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
				v88 = int32(0)
				m.G0 = v9 + int32(32)
				return v88
			}
		}
	}
}
func F_jsonb_recv(m *base.Module, l0 int32) int32 {
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
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v65 int32
	_ = v65
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
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	v6 = m.G0
	v8 = v6 - int32(144)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pq_getmsgint(m, v10, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			v23 = F_pq_getmsgtext(m, v10, v18-v19, v8+int32(12))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
				v26 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = v26
				v28 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v28
				*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v28
				v36 = v8 + int32(76)
				v38 = *(*int32)(unsafe.Add(mBase, _c_F_jsonb_recv[0]))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
				v41 = F_makeJsonLexContextCstringLen(m, v36, v23, v25, v39, int32(1))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					v43 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v43
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+64)) = uint8(v43)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(1310)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(1311)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = int32(1312)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(1313)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(1314)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(1315)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(56)
					v65 = F_pg_parse_json_or_errsave(m, v36, v8+int32(16), v43)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						if v65 != 0 {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
							v68 = F_JsonbValueToJsonb(m, v67)
							mBase = m.M
							v69 = m.ExcPending
							if v69 != 0 {
								return int32(0)
							} else {
								v70 = v68
								m.G0 = v8 + int32(144)
								return v70
							}
						} else {
							v70 = v28
							m.G0 = v8 + int32(144)
							return v70
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
				F_errmsg_internal(m, int32(_a_F_jsonb_recv_0), v8)
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_jsonb_recv_1), int32(99), int32(_a_F_jsonb_recv_2))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
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
func F_jsonb_set_lax(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
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
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v3 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	if v12 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	v8 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
	return int32(0)
L3:
	;
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v4 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	if v5 != int32(1) {
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
	v203 = m.ExcPending
	if v203 != 0 {
		goto L14
	} else {
		goto L68
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L14
	} else {
		goto L62
	}
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	if v15 == int32(0) {
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
	v163 = m.ExcPending
	if v163 != 0 {
		goto L14
	} else {
		goto L58
	}
L11:
	;
	v18 = F_jsonb_set(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	return v18
L16:
	;
	v26 = F_text_to_cstring(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v28 = int32(_a_F_jsonb_set_lax_0)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_jsonb_set_lax[0])))
	if base.B2i32(v31 == int32(0))|base.B2i32(v31 != v34) != 0 {
		v52 = v31
		v53 = v34
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v52-v53 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L19:
	;
	goto L18
L20:
	;
	v37 = v26
	v38 = v28
	goto L21
L21:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+1)))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+1)))
	if v42 == int32(0) {
		v52 = v42
		v53 = v41
		goto L19
	} else {
		goto L23
	}
L22:
	;
	v52 = v42
	v53 = v41
	goto L19
L23:
	;
	v45 = int32(1)
	if v42 == v41 {
		v37 = v37 + v45
		v38 = v38 + v45
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v57 = int32(_a_F_jsonb_set_lax_1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_jsonb_set_lax[1])))
	if base.B2i32(v60 == int32(0))|base.B2i32(v60 != v63) != 0 {
		v81 = v60
		v82 = v63
		goto L27
	} else {
		goto L28
	}
L26:
	;
	if v81-v82 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	v66 = v26
	v67 = v57
	goto L29
L29:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 == int32(0) {
		v81 = v71
		v82 = v70
		goto L27
	} else {
		goto L31
	}
L30:
	;
	v81 = v71
	v82 = v70
	goto L27
L31:
	;
	v74 = int32(1)
	if v71 == v70 {
		v66 = v66 + v74
		v67 = v67 + v74
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v89 = F_DirectFunctionCall1Coll(m, int32(486), int32(0), int32(_a_F_jsonb_set_lax_2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L14
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v97 = int32(_a_F_jsonb_set_lax_3)
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_jsonb_set_lax[2])))
	if base.B2i32(v100 == int32(0))|base.B2i32(v100 != v103) != 0 {
		v121 = v100
		v122 = v103
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v89
	v94 = F_jsonb_set(m, l0)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	return v94
L38:
	;
	if v121-v122 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	goto L38
L40:
	;
	v106 = v26
	v107 = v97
	goto L41
L41:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
	if v111 == int32(0) {
		v121 = v111
		v122 = v110
		goto L39
	} else {
		goto L43
	}
L42:
	;
	v121 = v111
	v122 = v110
	goto L39
L43:
	;
	v114 = int32(1)
	if v111 == v110 {
		v106 = v106 + v114
		v107 = v107 + v114
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v126 = F_jsonb_delete_path(m, l0)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L14
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v129 = int32(_a_F_jsonb_set_lax_4)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_jsonb_set_lax[3])))
	if base.B2i32(v132 == int32(0))|base.B2i32(v132 != v135) != 0 {
		v153 = v132
		v154 = v135
		goto L50
	} else {
		goto L51
	}
L48:
	;
	return v126
L49:
	;
	if v153-v154 != 0 {
		goto L6
	} else {
		goto L56
	}
L50:
	;
	goto L49
L51:
	;
	v138 = v26
	v139 = v129
	goto L52
L52:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v143 == int32(0) {
		v153 = v143
		v154 = v142
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v153 = v143
	v154 = v142
	goto L50
L54:
	;
	v146 = int32(1)
	if v143 == v142 {
		v138 = v138 + v146
		v139 = v139 + v146
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v157 = F_pg_detoast_datum(m, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	return v157
L58:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L14
	} else {
		goto L59
	}
L59:
	;
	F_errmsg(m, int32(_a_F_jsonb_set_lax_5), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L14
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_jsonb_set_lax_6), int32(_a_F_jsonb_set_lax_7), int32(_a_F_jsonb_set_lax_8))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
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
	v182 = m.ExcPending
	if v182 != 0 {
		goto L14
	} else {
		goto L63
	}
L63:
	;
	F_errmsg(m, int32(_a_F_jsonb_set_lax_9), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L14
	} else {
		goto L64
	}
L64:
	;
	F_errdetail(m, int32(_a_F_jsonb_set_lax_10), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L14
	} else {
		goto L65
	}
L65:
	;
	F_errhint(m, int32(_a_F_jsonb_set_lax_11), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L14
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_jsonb_set_lax_6), int32(_a_F_jsonb_set_lax_12), int32(_a_F_jsonb_set_lax_8))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
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
	v206 = m.ExcPending
	if v206 != 0 {
		goto L14
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_jsonb_set_lax_5), int32(0))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L14
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_jsonb_set_lax_6), int32(_a_F_jsonb_set_lax_13), int32(_a_F_jsonb_set_lax_8))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
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
func F_jsonb_string_to_tsvector(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_getTSCurrentConfig(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
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
				return int32(0)
			} else {
				v29 = F_make_tsvector(m, v22)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v31 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(32)
							return v29
						}
					} else {
						m.G0 = v7 + int32(32)
						return v29
					}
				}
			}
		}
	}
}
