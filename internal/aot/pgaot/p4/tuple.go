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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v92 float64
	_ = v92
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v107 int64
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v129 int64
	_ = v129
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v23 = *(*float64)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[0]))
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[1]))
	v29 = base.F64_mul(base.F64_mul(v23, base.F64_convert_i32_s(v25)), float64(1024))
	v30 = float64(4.294967295e+09)
	if base.F64_lt(v29, v30) != 0 {
		v33 = v29
	} else {
		v33 = v30
	}
	v35 = int32(_a_F_BuildTupleHashTable_0)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[2])) = l10
	v40 = F_palloc(m, int32(56))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		return int32(0)
	} else {
		v44 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v40)+36)) = v44
		v49 = (l9 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(v40)+32)) = v49
		*(*int32)(unsafe.Add(mBase, uint32(v40)+28)) = l12
		*(*int32)(unsafe.Add(mBase, uint32(v40)+24)) = l11
		*(*int32)(unsafe.Add(mBase, uint32(v40)+20)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v40)+8)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = l3
		*(*int64)(unsafe.Add(mBase, uint32(v40)+44)) = v44
		if l13 != 0 {
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[3]))
			v61 = int32(16)
			v65 = (int32(base.Ui32(v60)>>(uint(v61)%32)) ^ v60) * int32(-2048144789)
			v70 = (int32(base.Ui32(v65)>>(uint(int32(13))%32)) ^ v65) * int32(-1028477387)
			v74 = int32(base.Ui32(v70)>>(uint(v61)%32)) ^ v70
		} else {
			v74 = int32(0)
		}
		v77 = base.I32_div_u_s(base.I32_trunc_sat_f64_u(v33), v49+int32(12))
		if base.Ui32(l8) < base.Ui32(v77) {
			v79 = l8
		} else {
			v79 = v77
		}
		v81 = F_MemoryContextAllocZero(m, l10, int32(32))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v81)+28)) = v40
			*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = l10
			v86 = float64(4.294967296e+09)
			v89 = base.F64_div(base.F64_convert_i32_u(v79), float64(0.9))
			if base.F64_ge(v89, v86) != 0 {
				v92 = v86
			} else {
				v92 = v89
			}
			v93 = base.I64_trunc_sat_f64_u(v92)
			if base.Ui64(v93) <= base.Ui64(int64(2)) {
				v96 = int64(2)
			} else {
				v96 = v93
			}
			v97 = int64(1)
			if v96&(v96-v97) == int64(0) {
				v107 = v96
			} else {
				v107 = v97 << (uint(int64(64)-base.I64_clz(v96)) % 64)
			}
			if base.Ui64(v107*int64(12)) < base.Ui64(int64(2147483647)) {
				v116 = F_MemoryContextAllocExtended(m, l10, base.I32_wrap_i64(v107)*int32(12), int32(5))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = v116
					v119 = int64(1)
					if v107&(v107-v119) == int64(0) {
						v129 = v107
					} else {
						v129 = v119 << (uint(int64(64)-base.I64_clz(v107)) % 64)
					}
					if base.Ui64(int64(2147483647)) <= base.Ui64(v129*int64(12)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_BuildTupleHashTable_1), int32(0))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_BuildTupleHashTable_2), int32(327), int32(_a_F_BuildTupleHashTable_3))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v81))) = v129
						*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = base.I32_wrap_i64(v129) - int32(1)
						if v129 == int64(4294967296) {
							v146 = int32(-85899346)
						} else {
							v146 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v129), float64(0.9)))
						}
						*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v146
						*(*int32)(unsafe.Add(mBase, uint32(v40))) = v81
						v149 = F_CreateTupleDescCopy(m, l1)
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							v152 = F_MakeTupleTableSlot(m, v149, int32(_a_F_BuildTupleHashTable_4))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v40)+36)) = v152
								if l10 != l11 {
									v157 = l0
								} else {
									v157 = int32(0)
								}
								v158 = F_ExecBuildHash32FromAttrs(m, l1, l2, l6, l7, l3, l4, v157, v74)
								mBase = m.M
								v159 = m.ExcPending
								if v159 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v40)+12)) = v158
									v162 = F_ExecBuildGroupingEqual(m, l1, l1, l2, int32(_a_F_BuildTupleHashTable_4), l3, l4, l5, l7, v157)
									mBase = m.M
									v163 = m.ExcPending
									if v163 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v40)+16)) = v162
										v165 = F_CreateStandaloneExprContext(m)
										mBase = m.M
										v166 = m.ExcPending
										if v166 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v40)+52)) = v165
											*(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[2])) = v36
											return v40
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
				v175 = m.ExcPending
				if v175 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_BuildTupleHashTable_1), int32(0))
					mBase = m.M
					v179 = m.ExcPending
					if v179 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_BuildTupleHashTable_2), int32(327), int32(_a_F_BuildTupleHashTable_3))
						mBase = m.M
						v184 = m.ExcPending
						if v184 != 0 {
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
	v13 = int32(_a_F_LookupTupleHashEntry_0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0]))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v16
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
					*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v14
					m.G0 = v11 + int32(16)
					return v71
				} else {
					v58 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v58)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v61
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
						*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v14
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
				*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v14
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
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTupleTableSlot[0]))
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
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = v4*int32(116) + int32(20)
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	base.MemoryCopy(m, l0, l1, v8)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v10 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v15 = int32(0)
	v16 = v10
	goto L7
L5:
	;
	goto L6
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967295)
	return
L7:
	;
	v22 = l0 + v16<<(uint(int32(4))%32) + v15*int32(100)
	v23 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+110)) = uint8(v23)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+106)) = v23
	F_populate_compact_attribute(m, l0, v15)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	return
L10:
	;
	v30 = v15 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v30 < v31 {
		v15 = v30
		v16 = v31
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
}
func F_TupleDescCopyEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v35 int32
	_ = v35
	v2 = l1
	v5 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = int32(4)
	v11 = int32(100)
	v13 = l0 + v7<<(uint(v8)%32) + v2*v11
	v14 = int32(80)
	v15 = v13 - v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	base.MemoryCopy(m, v15, l2+v16<<(uint(v8)%32)+l3*v11-v14, v11)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+6)) = v5
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+74)) = uint16(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+10)) = uint8(v5)
	F_populate_compact_attribute(m, l0, v2-int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return
	} else {
		return
	}
}
