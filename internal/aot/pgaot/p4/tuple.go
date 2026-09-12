package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildTupleHashTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32 {
	mBase := m.M
	_ = mBase
	var v23 float64
	_ = v23
	var v25 int32
	_ = v25
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v33 float64
	_ = v33
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 float64
	_ = v93
	var v96 float64
	_ = v96
	var v99 float64
	_ = v99
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v121 int64
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v143 int64
	_ = v143
	var v158 float64
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	v23 = *(*float64)(unsafe.Add(mBase, _consts[498]))
	v25 = *(*int32)(unsafe.Add(mBase, _consts[135]))
	v29 = base.F64_mul(base.F64_mul(v23, base.F64_convert_i32_s(v25)), float64(1024))
	v30 = float64(4.294967295e+09)
	if base.F64_lt(v29, v30) != 0 {
		v33 = v29
	} else {
		v33 = v30
	}
	if base.F64_lt(v33, float64(4.294967296e+09))&base.F64_ge(v33, float64(0)) != 0 {
		v39 = base.I32_trunc_f64_u(v33)
		v41 = v39
	} else {
		v41 = int32(0)
	}
	v42 = int32(4536272)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = l10
	v47 = F_palloc(m, int32(56))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return int32(0)
	} else {
		v51 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v47)+36)) = v51
		v56 = (l9 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(v47)+32)) = v56
		*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = l12
		*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = l11
		*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v47)+8)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = l3
		*(*int64)(unsafe.Add(mBase, uint32(v47)+44)) = v51
		if l13 != 0 {
			v67 = *(*int32)(unsafe.Add(mBase, _consts[181]))
			v68 = int32(16)
			v72 = (int32(base.Ui32(v67)>>(uint(v68)%32)) ^ v67) * int32(-2048144789)
			v77 = (int32(base.Ui32(v72)>>(uint(int32(13))%32)) ^ v72) * int32(-1028477387)
			v81 = int32(base.Ui32(v77)>>(uint(v68)%32)) ^ v77
		} else {
			v81 = int32(0)
		}
		v84 = base.I32_div_u_s(v41, v56+int32(12))
		if base.Ui32(l8) < base.Ui32(v84) {
			v86 = l8
		} else {
			v86 = v84
		}
		v88 = F_MemoryContextAllocZero(m, l10, int32(32))
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v88)+28)) = v47
			*(*int32)(unsafe.Add(mBase, uint32(v88)+24)) = l10
			v93 = float64(4.294967296e+09)
			v96 = base.F64_div(base.F64_convert_i32_u(v86), float64(0.9))
			if base.F64_ge(v96, v93) != 0 {
				v99 = v93
			} else {
				v99 = v96
			}
			if base.F64_lt(v99, float64(1.8446744073709552e+19))&base.F64_ge(v99, float64(0)) != 0 {
				v105 = base.I64_trunc_f64_u(v99)
				v107 = v105
			} else {
				v107 = int64(0)
			}
			if base.Ui64(v107) <= base.Ui64(int64(2)) {
				v110 = int64(2)
			} else {
				v110 = v107
			}
			v111 = int64(1)
			if v110&(v110-v111) == int64(0) {
				v121 = v110
			} else {
				v121 = v111 << (uint(int64(64)-base.I64_clz(v110)) % 64)
			}
			if base.Ui64(v121*int64(12)) < base.Ui64(int64(2147483647)) {
				v130 = F_MemoryContextAllocExtended(m, l10, base.I32_wrap_i64(v121)*int32(12), int32(5))
				mBase = m.M
				v131 = m.ExcPending
				if v131 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v88)+20)) = v130
					v133 = int64(1)
					if v121&(v121-v133) == int64(0) {
						v143 = v121
					} else {
						v143 = v133 << (uint(int64(64)-base.I64_clz(v121)) % 64)
					}
					if base.Ui64(int64(2147483647)) <= base.Ui64(v143*int64(12)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(409411), int32(0))
							mBase = m.M
							v189 = m.ExcPending
							if v189 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(334097), int32(327), int32(348898))
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v88))) = v143
						*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = base.I32_wrap_i64(v143) - int32(1)
						v158 = base.F64_mul(base.F64_convert_i64_u(v143), float64(0.9))
						if base.F64_lt(v158, float64(4.294967296e+09))&base.F64_ge(v158, float64(0)) != 0 {
							v164 = base.I32_trunc_f64_u(v158)
							v166 = v164
						} else {
							v166 = int32(0)
						}
						if v143 == int64(4294967296) {
							v167 = int32(-85899346)
						} else {
							v167 = v166
						}
						*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v167
						*(*int32)(unsafe.Add(mBase, uint32(v47))) = v88
						v196 = F_CreateTupleDescCopy(m, l1)
						mBase = m.M
						v197 = m.ExcPending
						if v197 != 0 {
							return int32(0)
						} else {
							v199 = F_MakeSingleTupleTableSlot(m, v196, int32(1632252))
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v47)+36)) = v199
								if l10 != l11 {
									v204 = l0
								} else {
									v204 = int32(0)
								}
								v205 = F_ExecBuildHash32FromAttrs(m, l1, l2, l6, l7, l3, l4, v204, v81)
								mBase = m.M
								v206 = m.ExcPending
								if v206 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v205
									v209 = F_ExecBuildGroupingEqual(m, l1, l1, l2, int32(1632252), l3, l4, l5, l7, v204)
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v47)+16)) = v209
										v212 = F_CreateStandaloneExprContext(m)
										mBase = m.M
										v213 = m.ExcPending
										if v213 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v47)+52)) = v212
											*(*int32)(unsafe.Add(mBase, _consts[28])) = v43
											return v47
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
				v172 = m.ExcPending
				if v172 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(409411), int32(0))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(334097), int32(327), int32(348898))
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
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
func F_ExecResetTupleTable(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 != 0 {
		goto L29
	} else {
		goto L30
	}
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = v3
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16+v15<<(uint(int32(2))%32))))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	m.T0[v22].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	return
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	m.T0[v26].(func(*base.Module, int32))(m, v20)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if int32(0) <= v30 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	if l1 != 0 {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	F_DecrTupleDescRefCount(m, v29)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = int32(0)
	goto L11
L15:
	;
	goto L14
L16:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
	if v37&int32(16) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v53 = v15 + int32(1)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v53 < v54 {
		v15 = v53
		goto L4
	} else {
		goto L28
	}
L19:
	;
	F_pfree(m, v20)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L27
	}
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
	if v40 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_pfree(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	if v43 == int32(0) {
		goto L19
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	F_pfree(m, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	goto L19
L27:
	;
	goto L18
L28:
	;
	goto L5
L29:
	;
	F_list_free(m, l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	return
L32:
	;
	goto L31
}
func F_LookupTupleHashEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
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
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(4536272)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v16
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+44)) = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+44))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
	v31 = m.T0[v30].(func(*base.Module, int32, int32, int32) int32)(m, v26, v27, v11+int32(14))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v35 = int32(16)
		v39 = (int32(base.Ui32(v31)>>(uint(v35)%32)) ^ v31) * int32(-2048144789)
		v44 = (int32(base.Ui32(v39)>>(uint(int32(13))%32)) ^ v39) * int32(-1028477387)
		v47 = int32(base.Ui32(v44)>>(uint(v35)%32)) ^ v44
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if l2 != 0 {
			v51 = F_tuplehash_insert_hash_internal(m, v48, v47, v11+int32(15))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int32(0)
			} else {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
				if v53 == int32(1) {
					v56 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
					v71 = v51
					if l3 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
					} else {
					}
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v14
					m.G0 = v11 + int32(16)
					return v71
				} else {
					v58 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v58)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v61
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+48))
					v66 = m.T0[v65].(func(*base.Module, int32, int32) int32)(m, l1, v63)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v51))) = v66
						v71 = v51
						if l3 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
						} else {
						}
						*(*int32)(unsafe.Add(mBase, _consts[28])) = v14
						m.G0 = v11 + int32(16)
						return v71
					}
				}
			}
		} else {
			v69 = F_tuplehash_lookup_hash_internal(m, v48, v47)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return int32(0)
			} else {
				v71 = v69
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
				} else {
				}
				*(*int32)(unsafe.Add(mBase, _consts[28])) = v14
				m.G0 = v11 + int32(16)
				return v71
			}
		}
	}
}
func F_MakeTupleTableSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l0 == int32(0) {
		v29 = int32(2)
		v30 = v6
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(7)
		v14 = int32(-8)
		v29 = int32(18)
		v30 = (v11+v12)&v14 + (v6+v12)&v14 + (v11<<(uint(int32(2))%32)+v12)&v14
	}
	v31 = F_palloc0(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(443)
		*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = l0
		v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)))
		v40 = v39 | v29
		*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)) = uint16(v40)
		v43 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		v44 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v31)+6)) = uint16(v44)
		*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v43
		if l0 == v44 {
			v70 = l1
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
			m.T0[v72].(func(*base.Module, int32))(m, v31)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				return v31
			}
		} else {
			v49 = int32(7)
			v51 = int32(-8)
			v53 = v31 + (v6+v49)&v51
			*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v53
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v53 + (v55<<(uint(int32(2))%32)+v49)&v51
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v64 < int32(0) {
				v70 = l1
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
				m.T0[v72].(func(*base.Module, int32))(m, v31)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					return v31
				}
			} else {
				F_IncrTupleDescRefCount(m, l0)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v70 = v69
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
					m.T0[v72].(func(*base.Module, int32))(m, v31)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						return v31
					}
				}
			}
		}
	}
}
func F_TupleDescCopy(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v9 = v5*int32(116) + int32(20)
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if int32(0) < v12 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v10 = F__emscripten_memcpy_bulkmem(m, l0, l1, v9)
	mBase = m.M
	v11 = v10
	goto L4
L3:
	;
	v11 = l0
	goto L4
L4:
	;
	goto L1
L5:
	;
	v19 = int32(0)
	v20 = v12
	goto L8
L6:
	;
	goto L7
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+12)) = int64(4294967295)
	return
L8:
	;
	v27 = v11 + int32(20) + v20<<(uint(int32(4))%32) + v19*int32(100)
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+90)) = uint8(v28)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+86)) = v28
	F_populate_compact_attribute(m, v11, v19)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	return
L11:
	;
	v35 = v19 + int32(1)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	if v35 < v36 {
		v19 = v35
		v20 = v36
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
}
func F_TupleDescCopyEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	v2 = l1
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(4)
	v10 = int32(100)
	v12 = l0 + v6<<(uint(v7)%32) + v2*v10
	v13 = int32(80)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v25 = F__emscripten_memcpy_bulkmem(m, v12-v13, l2+v15<<(uint(v7)%32)+l3*v10-v13, v10)
	mBase = m.M
	v27 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+6)) = v27
	*(*uint16)(unsafe.Add(mBase, uint32(v25)+74)) = uint16(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v27)
	F_populate_compact_attribute(m, l0, v2-int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		return
	}
}
