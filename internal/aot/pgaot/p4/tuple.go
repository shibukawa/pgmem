package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildTupleHashTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 float64, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32) int32 {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 float64
	_ = v80
	var v83 float64
	_ = v83
	var v86 float64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v101 int64
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v123 int64
	_ = v123
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(l8)&int64(9223372036854775807)))|base.F64_le(l8, float64(0)) != 0 {
		v32 = int32(1)
	} else {
		if base.F64_ge(l8, float64(4.294967295e+09)) != 0 {
			v32 = int32(-1)
		} else {
			v32 = base.I32_trunc_sat_f64_u(l8)
		}
	}
	v33 = int32(_a_F_BuildTupleHashTable_0)
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[0])) = l10
	v38 = F_palloc(m, int32(56))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return int32(0)
	} else {
		v42 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v38)+36)) = v42
		*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = (l9 + int32(7)) & int32(-8)
		*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = l12
		*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = l11
		*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = l7
		*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = l3
		*(*int64)(unsafe.Add(mBase, uint32(v38)+44)) = v42
		if l13 != 0 {
			v58 = *(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[1]))
			v59 = int32(16)
			v63 = (int32(base.Ui32(v58)>>(uint(v59)%32)) ^ v58) * int32(-2048144789)
			v68 = (int32(base.Ui32(v63)>>(uint(int32(13))%32)) ^ v63) * int32(-1028477387)
			v73 = int32(base.Ui32(v68)>>(uint(v59)%32)) ^ v68
		} else {
			v73 = int32(0)
		}
		v75 = F_MemoryContextAllocZero(m, l10, int32(32))
		mBase = m.M
		v76 = m.ExcPending
		if v76 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v75)+28)) = v38
			*(*int32)(unsafe.Add(mBase, uint32(v75)+24)) = l10
			v80 = float64(4.294967296e+09)
			v83 = base.F64_div(base.F64_convert_i32_u(v32), float64(0.9))
			if base.F64_ge(v83, v80) != 0 {
				v86 = v80
			} else {
				v86 = v83
			}
			v87 = base.I64_trunc_sat_f64_u(v86)
			if base.Ui64(v87) <= base.Ui64(int64(2)) {
				v90 = int64(2)
			} else {
				v90 = v87
			}
			v91 = int64(1)
			if v90&(v90-v91) == int64(0) {
				v101 = v90
			} else {
				v101 = v91 << (uint(int64(64)-base.I64_clz(v90)) % 64)
			}
			if base.Ui64(v101*int64(12)) < base.Ui64(int64(2147483647)) {
				v110 = F_MemoryContextAllocExtended(m, l10, base.I32_wrap_i64(v101)*int32(12), int32(5))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v75)+20)) = v110
					v113 = int64(1)
					if v101&(v101-v113) == int64(0) {
						v123 = v101
					} else {
						v123 = v113 << (uint(int64(64)-base.I64_clz(v101)) % 64)
					}
					if base.Ui64(int64(2147483647)) <= base.Ui64(v123*int64(12)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_BuildTupleHashTable_1), int32(0))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_BuildTupleHashTable_2), int32(332), int32(_a_F_BuildTupleHashTable_3))
								mBase = m.M
								v175 = m.ExcPending
								if v175 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v75))) = v123
						*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = base.I32_wrap_i64(v123) - int32(1)
						if v123 == int64(4294967296) {
							v140 = int32(-85899346)
						} else {
							v140 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v123), float64(0.9)))
						}
						*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = v140
						*(*int32)(unsafe.Add(mBase, uint32(v38))) = v75
						v143 = F_CreateTupleDescCopy(m, l1)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							v146 = F_MakeSingleTupleTableSlot(m, v143, int32(_a_F_BuildTupleHashTable_4))
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v146
								v149 = F_ExecBuildHash32FromAttrs(m, l1, l2, l6, l7, l3, l4, l0, v73)
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v149
									v153 = F_ExecBuildGroupingEqual(m, l1, l1, l2, int32(_a_F_BuildTupleHashTable_4), l3, l4, l5, l7, l0)
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v153
										v156 = F_CreateStandaloneExprContext(m)
										mBase = m.M
										v157 = m.ExcPending
										if v157 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v156
											*(*int32)(unsafe.Add(mBase, _c_F_BuildTupleHashTable[0])) = v34
											return v38
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
				v166 = m.ExcPending
				if v166 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_BuildTupleHashTable_1), int32(0))
					mBase = m.M
					v170 = m.ExcPending
					if v170 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_BuildTupleHashTable_2), int32(332), int32(_a_F_BuildTupleHashTable_3))
						mBase = m.M
						v175 = m.ExcPending
						if v175 != 0 {
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
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
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
	var v72 int32
	_ = v72
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
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
	v31 = m.T0[v30].(func(*base.Module, int32, int32, int32) int64)(m, v26, v27, v11+int32(14))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		v35 = base.I32_wrap_i64(v31)
		v36 = int32(16)
		v40 = (int32(base.Ui32(v35)>>(uint(v36)%32)) ^ v35) * int32(-2048144789)
		v45 = (int32(base.Ui32(v40)>>(uint(int32(13))%32)) ^ v40) * int32(-1028477387)
		v48 = int32(base.Ui32(v45)>>(uint(v36)%32)) ^ v45
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if l2 != 0 {
			v52 = F_tuplehash_insert_hash_internal(m, v49, v48, v11+int32(15))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
				if v54 == int32(1) {
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v57)
					v72 = v52
					if l3 != 0 {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
					} else {
					}
					*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v14
					m.G0 = v11 + int32(16)
					return v72
				} else {
					v59 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v59)
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v62
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+48))
					v67 = m.T0[v66].(func(*base.Module, int32, int32) int32)(m, l1, v64)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v52))) = v67
						v72 = v52
						if l3 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
						} else {
						}
						*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v14
						m.G0 = v11 + int32(16)
						return v72
					}
				}
			}
		} else {
			v70 = F_tuplehash_lookup_hash_internal(m, v49, v48)
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int32(0)
			} else {
				v72 = v70
				if l3 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v48
				} else {
				}
				*(*int32)(unsafe.Add(mBase, _c_F_LookupTupleHashEntry[0])) = v14
				m.G0 = v11 + int32(16)
				return v72
			}
		}
	}
}
func F_MakeTupleTableSlot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l0 == int32(0) {
		v29 = int32(2)
		v30 = v10
	} else {
		v15 = int32(7)
		v17 = int32(-8)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v29 = int32(18)
		v30 = (v10+v15)&v17 + v19<<(uint(int32(3))%32) + (v19+v15)&v17
	}
	v31 = F_palloc0(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = l0
		v36 = v29 | l2&int32(_a_F_MakeTupleTableSlot_0)
		*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)) = uint16(v36)
		*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(449)
		*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = l1
		v42 = *(*int32)(unsafe.Add(mBase, _c_F_MakeTupleTableSlot[0]))
		v43 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v31)+6)) = uint16(v43)
		*(*int32)(unsafe.Add(mBase, uint32(v31)+28)) = v42
		if l0 != 0 {
			v50 = v31 + (v10+int32(7))&int32(-8)
			*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v50
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v31)+20)) = v50 + v52<<(uint(int32(3))%32)
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if int32(0) <= v57 {
				F_IncrTupleDescRefCount(m, l0)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					if l2&int32(8) != 0 {
						v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v66 = v64
					} else {
						v66 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v66
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
					m.T0[v70].(func(*base.Module, int32))(m, v31)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						return v31
					}
				}
			} else {
				if l2&int32(8) != 0 {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v66 = v64
				} else {
					v66 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v31)+24)) = v66
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
				m.T0[v70].(func(*base.Module, int32))(m, v31)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					return v31
				}
			}
		} else {
			v69 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
			m.T0[v70].(func(*base.Module, int32))(m, v31)
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				return v31
			}
		}
	}
}
func F_TupleDescCopy(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v15 = v11*int32(108) + int32(28)
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	base.MemoryCopy(m, l0, l1, v15)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v17 < v18 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v121
	return
L5:
	;
	v22 = v17
	v23 = v18
	goto L8
L6:
	;
	goto L7
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	v121 = v17
	v122 = v18
	goto L4
L8:
	;
	v36 = l0 + v23<<(uint(int32(3))%32) + v22*int32(100)
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v36)+118)) = uint8(v37)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+114)) = v37
	F_populate_compact_attribute(m, l0, v22)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(-1)
	v49 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v49
	if v45 <= v49 {
		v121 = v49
		v122 = v45
		goto L4
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	v44 = v22 + int32(1)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v44 < v45 {
		v22 = v44
		v23 = v45
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
L13:
	;
	v55 = l0 + int32(28)
	v60 = v49
	v62 = v45
	v65 = int32(0)
	goto L15
L14:
	;
	v121 = v113
	v122 = v92
	goto L4
L15:
	;
	v71 = v55 + v45<<(uint(int32(3))%32) + v60*int32(100)
	v74 = v55 + v60<<(uint(int32(3))%32)
	if v45 != v62 {
		v92 = v62
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v113 = v45
	goto L14
L17:
	;
	v93 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74)+2)))
	if v93 <= int32(0) {
		v113 = v60
		goto L14
	} else {
		goto L25
	}
L18:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+7)))
	if v76 != int32(118) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v92 = v60
	goto L17
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
	if v79 != int32(1) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+6)))
	if v82&int32(6) != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74)+2)))
	if v85 <= int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+90)))
	if v88 != int32(118) {
		v92 = v45
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+90)))
	if v96 == int32(118) {
		v113 = v60
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+5)))
	v105 = (v65 + v99 - int32(1)) & (int32(0) - v99)
	if int32(_a_F_TupleDescCopy_0) < v105 {
		v113 = v60
		goto L14
	} else {
		goto L27
	}
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v74))) = uint16(v105)
	v111 = v60 + int32(1)
	if v111 != v45 {
		v60 = v111
		v62 = v92
		v65 = v105 + v93
		goto L15
	} else {
		goto L28
	}
L28:
	;
	goto L16
}
func F_TupleDescCopyEntry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
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
	var v34 int32
	_ = v34
	v2 = l1
	v5 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = int32(3)
	v10 = int32(100)
	v12 = l0 + v6<<(uint(v7)%32) + v2*v10
	v13 = int32(72)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	base.MemoryCopy(m, v12-v13, l2+v15<<(uint(v7)%32)+l3*v10-v13, v10)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+14)) = v5
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+2)) = uint16(v2)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+18)) = uint8(v5)
	F_populate_compact_attribute(m, l0, v2-int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		return
	} else {
		return
	}
}
