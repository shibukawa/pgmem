package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_IdleStatsUpdateTimeoutHandler(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	v2 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_IdleStatsUpdateTimeoutHandler[0])) = v2
	*(*int32)(unsafe.Add(mBase, _c_F_IdleStatsUpdateTimeoutHandler[1])) = v2
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_IdleStatsUpdateTimeoutHandler[2]))
	F_SetLatch(m, v8)
	mBase = m.M
	return
}
func F_IncrementVarSublevelsUp_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v3 = int32(0)
	if l0 == v3 {
		v95 = v3
		return v95
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v7 <= int32(57) {
			switch v7 - int32(6) {
			case 0:
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(v24) < base.Ui32(v25) {
					v95 = v3
					return v95
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v27 + v24
					return int32(0)
				}
			default:
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			case 3:
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(v48) < base.Ui32(v49) {
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v51 + v48
				}
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			case 4:
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if base.Ui32(v54) < base.Ui32(v55) {
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v57 + v54
				}
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			}
		} else {
			switch v7 - int32(58) {
			case 0:
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v32 != 0 {
					v95 = v3
					return v95
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_IncrementVarSublevelsUp_walker_0), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_IncrementVarSublevelsUp_walker_1), int32(818), int32(_a_F_IncrementVarSublevelsUp_walker_2))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			case 1, 2, 4, 5, 6, 7, 8:
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			case 3:
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v60 < v61 {
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v63 + v60
				}
				v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					return v68
				}
			case 9:
				v82 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v82 + int32(1)
				v88 = F_query_tree_walker_impl(m, l0, int32(1050), l1, int32(16))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v90 - int32(1)
					v95 = v88
					return v95
				}
			default:
				if v7 == int32(101) {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v71 != int32(6) {
						v95 = v3
						return v95
					} else {
						v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						v75 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if base.Ui32(v74) < base.Ui32(v75) {
							v95 = v3
							return v95
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+88)) = v77 + v74
							return int32(0)
						}
					}
				} else {
					if v7 != int32(319) {
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if base.Ui32(v18) < base.Ui32(v19) {
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v21 + v18
						}
					}
					v68 = F_expression_tree_walker_impl(m, l0, int32(1050), l1)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						return v68
					}
				}
			}
		}
	}
}
func F_InitBuildState_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v130 float64
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
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
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	v15 = F_HnswGetTypeInfo(m, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v15
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
		if v18 == int32(0) {
			v23 = int32(16)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v23 = v22
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
		if v25 == int32(0) {
			v30 = int32(64)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
			v30 = v29
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v30
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
		v34 = int32(4)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(v34)%32))+96))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v32+v39<<(uint(v34)%32))+88))
		if v43 != int32(1562) {
			if v37 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v189 = m.ExcPending
				if v189 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_InitBuildState_1_0), int32(0))
						mBase = m.M
						v196 = m.ExcPending
						if v196 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_InitBuildState_1_1), int32(706), int32(_a_F_InitBuildState_1_2))
							mBase = m.M
							v201 = m.ExcPending
							if v201 != 0 {
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
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				if v49 < v37 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v205 = m.ExcPending
					if v205 != 0 {
						return
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v208 = m.ExcPending
						if v208 != 0 {
							return
						} else {
							v209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v210
							F_errmsg(m, int32(_a_F_InitBuildState_1_3), v9)
							mBase = m.M
							v214 = m.ExcPending
							if v214 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InitBuildState_1_1), int32(711), int32(_a_F_InitBuildState_1_2))
								mBase = m.M
								v219 = m.ExcPending
								if v219 != 0 {
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
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v30 < v51<<(uint(int32(1))%32) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v226 = m.ExcPending
							if v226 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_InitBuildState_1_4), int32(0))
								mBase = m.M
								v230 = m.ExcPending
								if v230 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_InitBuildState_1_1), int32(716), int32(_a_F_InitBuildState_1_2))
									mBase = m.M
									v235 = m.ExcPending
									if v235 != 0 {
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
						v55 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v55
						*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v55
						F_HnswInitSupport(m, l0+int32(48), l2)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[0]))
							v66 = F_mul_size(m, v64, int32(1024))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_HnswInitLockTranche(m)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+156)) = uint8(v70)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+132)) = v70
									*(*int32)(unsafe.Add(mBase, uint32(l0)+112)) = v70
									*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v70
									*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = int64(0)
									v80 = int32(2147483647)
									if base.Ui32(v80) <= base.Ui32(v66) {
										v83 = v80
									} else {
										v83 = v66
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v83
									v85 = int32(0)
									atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0)+64)), uint32(v85))
									v89 = l0 + int32(80)
									v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[1]))
									*(*uint16)(unsafe.Add(mBase, uint32(v89))) = uint16(v91)
									*(*int32)(unsafe.Add(mBase, uint32(v89)+4)) = int32(1073741824)
									*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = int64(-1)
									v98 = l0 + int32(96)
									v100 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[1]))
									*(*uint16)(unsafe.Add(mBase, uint32(v98))) = uint16(v100)
									*(*int32)(unsafe.Add(mBase, uint32(v98)+4)) = int32(1073741824)
									*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = int64(-1)
									v107 = l0 + int32(116)
									v109 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[1]))
									*(*uint16)(unsafe.Add(mBase, uint32(v107))) = uint16(v109)
									*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = int32(1073741824)
									*(*int64)(unsafe.Add(mBase, uint32(v107)+8)) = int64(-1)
									v116 = l0 + int32(140)
									v118 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[1]))
									*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v118)
									*(*int32)(unsafe.Add(mBase, uint32(v116)+4)) = int32(1073741824)
									*(*int64)(unsafe.Add(mBase, uint32(v116)+8)) = int64(-1)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = l0 - int32(-64)
									v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v130 = F_log(m, base.F64_convert_i32_s(v128))
									mBase = m.M
									*(*float64)(unsafe.Add(mBase, uint32(l0)+168)) = base.F64_div(float64(1), v130)
									v133 = int32(63)
									v135 = base.I32_div_u_s(int32(1358), v128)
									v137 = v135 - int32(2)
									if base.Ui32(v133) <= base.Ui32(v137) {
										v140 = v133
									} else {
										v140 = v137
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v140
									v143 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[2]))
									v145 = int32(_a_F_InitBuildState_1_5)
									v148 = F_GenerationContextCreate(m, v143, int32(_a_F_InitBuildState_1_6), v145, v145, v145)
									mBase = m.M
									v149 = m.ExcPending
									if v149 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v148
										v152 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[2]))
										v157 = F_AllocSetContextCreateInternal(m, v152, int32(_a_F_InitBuildState_1_7), int32(0), int32(_a_F_InitBuildState_1_8), int32(_a_F_InitBuildState_1_9))
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = int32(0)
											*(*int64)(unsafe.Add(mBase, uint32(l0)+196)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = int32(_a_F_InitBuildState_1_10)
											*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v157
											*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = l0
											m.G0 = v9 + int32(16)
											return
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
			v173 = m.ExcPending
			if v173 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v176 = m.ExcPending
				if v176 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_InitBuildState_1_11), int32(0))
					mBase = m.M
					v180 = m.ExcPending
					if v180 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitBuildState_1_1), int32(700), int32(_a_F_InitBuildState_1_2))
						mBase = m.M
						v185 = m.ExcPending
						if v185 != 0 {
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
func F_InitProcess(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v116 int64
	_ = v116
	var v118 int32
	_ = v118
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
	if v9 != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
		if v11 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v291 = m.ExcPending
			if v291 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_InitProcess_0), int32(0))
				mBase = m.M
				v295 = m.ExcPending
				if v295 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InitProcess_1), int32(402), int32(_a_F_InitProcess_2))
					mBase = m.M
					v300 = m.ExcPending
					if v300 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
			if v13 == int32(1) {
				F_RegisterPostmasterChildActive(m)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
					v21 = v19 - int32(3)
					if base.Ui32(v21) <= base.Ui32(int32(4)) {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_InitProcess[4])))
						v28 = v26
					} else {
						v28 = int32(20)
					}
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
					v31 = v28 + v30
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
					v36 = base.AtomicRmwXchg32(m, v33, int32(0), int32(1))
					if v36 != 0 {
						v38 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
						F_s_lock(m, v38, int32(_a_F_InitProcess_1), int32(432), int32(_a_F_InitProcess_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[6])) = v46
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
							v50 = int32(0)
							if base.B2i32(v49 == v50)|base.B2i32(v49 == v31) == v50 {
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v57
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
								*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
								v61 = int32(_a_F_InitProcess_3)
								*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1])) = v49
								v64 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
								v65 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v64))), uint32(v65))
								v68 = int32(_a_F_InitProcess_4)
								v70 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
								v72 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
								v76 = base.I32_div_s(v70-v73, int32(640))
								*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7])) = v76
								v78 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v70))) = v78
								v81 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v81)+612)) = v65
								*(*uint8)(unsafe.Add(mBase, uint32(v81)+608)) = uint8(v65)
								*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v81)+36)) = v78
								v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[8]))
								*(*int32)(unsafe.Add(mBase, uint32(v81)+44)) = v91
								v94 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7]))
								*(*int64)(unsafe.Add(mBase, uint32(v81)+56)) = v78
								*(*int32)(unsafe.Add(mBase, uint32(v81)+52)) = v94
								*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = v78
								v101 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
								*(*int32)(unsafe.Add(mBase, uint32(v81)+120)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v81)+92)) = v78
								*(*uint16)(unsafe.Add(mBase, uint32(v81)+74)) = uint16(v65)
								*(*uint8)(unsafe.Add(mBase, uint32(v81)+124)) = uint8(base.B2i32(v101 == int32(4)))
								*(*uint8)(unsafe.Add(mBase, uint32(v81)+72)) = uint8(base.B2i32(v101 == int32(1)))
								v116 = base.AtomicRmwXchg64(m, v81, int32(112), v78)
								v118 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
								*(*int64)(unsafe.Add(mBase, uint32(v118)+560)) = v78
								*(*uint8)(unsafe.Add(mBase, uint32(v118)+536)) = uint8(v65)
								*(*uint8)(unsafe.Add(mBase, uint32(v118)+73)) = uint8(v65)
								*(*int64)(unsafe.Add(mBase, uint32(v118)+128)) = v78
								*(*int64)(unsafe.Add(mBase, uint32(v118)+136)) = v78
								*(*int32)(unsafe.Add(mBase, uint32(v118)+144)) = v65
								*(*int64)(unsafe.Add(mBase, uint32(v118)+544)) = v78
								*(*uint8)(unsafe.Add(mBase, uint32(v118)+552)) = uint8(v65)
								*(*int64)(unsafe.Add(mBase, uint32(v118)+576)) = v78
								*(*int64)(unsafe.Add(mBase, uint32(v118)+568)) = int64(-1)
								F_OwnLatch(m, v118+int32(20))
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return
								} else {
									F_SwitchToSharedLatch(m)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return
									} else {
										v146 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
										*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v146 + int32(548)
										v152 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
										v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
										F_PGSemaphoreReset(m, v153)
										mBase = m.M
										F_on_shmem_exit(m, int32(1116), int32(0))
										mBase = m.M
										v158 = m.ExcPending
										if v158 != 0 {
											return
										} else {
											v159 = int32(_a_F_InitProcess_5)
											v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10]))
											v163 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10])) = v163
											v167 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
											v170 = F_palloc(m, v167<<(uint(int32(2))%32))
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13])) = v170
												v175 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
												v178 = F_palloc(m, v175*int32(24))
												mBase = m.M
												v179 = m.ExcPending
												if v179 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[14])) = v178
													v183 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13]))
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[15])) = v183
													v187 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
													v190 = F_palloc(m, v187<<(uint(int32(2))%32))
													mBase = m.M
													v191 = m.ExcPending
													if v191 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[16])) = v190
														v195 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
														v198 = F_palloc(m, v195<<(uint(int32(2))%32))
														mBase = m.M
														v199 = m.ExcPending
														if v199 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[17])) = v198
															v203 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
															v205 = base.I32_div_s(v203, int32(2))
															v208 = F_palloc(m, v205*int32(12))
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[18])) = v208
																v213 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
																v216 = F_palloc(m, v213<<(uint(int32(2))%32))
																mBase = m.M
																v217 = m.ExcPending
																if v217 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[19])) = v216
																	v221 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[20])) = v221
																	v226 = F_palloc(m, v221*int32(20))
																	mBase = m.M
																	v227 = m.ExcPending
																	if v227 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[21])) = v226
																		v231 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[22])) = v231 << (uint(int32(2)) % 32)
																		v237 = F_palloc(m, v231*int32(80))
																		mBase = m.M
																		v238 = m.ExcPending
																		if v238 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10])) = v160
																			*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[23])) = v237
																			v244 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
																			if v244 != 0 {
																				F_AttachSharedMemoryStructs(m)
																				mBase = m.M
																				v246 = m.ExcPending
																				if v246 != 0 {
																					return
																				} else {
																					m.G0 = v6 + int32(16)
																					return
																				}
																			} else {
																				m.G0 = v6 + int32(16)
																				return
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
								v251 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
								v252 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v251))), uint32(v252))
								v256 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
								F_errstart_cold(m, int32(22), v252)
								mBase = m.M
								v260 = m.ExcPending
								if v260 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_InitProcess_6))
									mBase = m.M
									v263 = m.ExcPending
									if v263 != 0 {
										return
									} else {
										if v256 == int32(6) {
											v302 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[24]))
											*(*int32)(unsafe.Add(mBase, uint32(v6))) = v302
											F_errmsg(m, int32(_a_F_InitProcess_7), v6)
											mBase = m.M
											v306 = m.ExcPending
											if v306 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitProcess_1), int32(454), int32(_a_F_InitProcess_2))
												mBase = m.M
												v311 = m.ExcPending
												if v311 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										} else {
											F_errmsg(m, int32(_a_F_InitProcess_8), int32(0))
											mBase = m.M
											v269 = m.ExcPending
											if v269 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitProcess_1), int32(457), int32(_a_F_InitProcess_2))
												mBase = m.M
												v274 = m.ExcPending
												if v274 != 0 {
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
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
						*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[6])) = v46
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
						v50 = int32(0)
						if base.B2i32(v49 == v50)|base.B2i32(v49 == v31) == v50 {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v57
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
							v61 = int32(_a_F_InitProcess_3)
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1])) = v49
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
							v65 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v64))), uint32(v65))
							v68 = int32(_a_F_InitProcess_4)
							v70 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
							v72 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
							v76 = base.I32_div_s(v70-v73, int32(640))
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7])) = v76
							v78 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v70))) = v78
							v81 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v81)+612)) = v65
							*(*uint8)(unsafe.Add(mBase, uint32(v81)+608)) = uint8(v65)
							*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v65
							*(*int64)(unsafe.Add(mBase, uint32(v81)+36)) = v78
							v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[8]))
							*(*int32)(unsafe.Add(mBase, uint32(v81)+44)) = v91
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7]))
							*(*int64)(unsafe.Add(mBase, uint32(v81)+56)) = v78
							*(*int32)(unsafe.Add(mBase, uint32(v81)+52)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = v78
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v81)+120)) = v65
							*(*int64)(unsafe.Add(mBase, uint32(v81)+92)) = v78
							*(*uint16)(unsafe.Add(mBase, uint32(v81)+74)) = uint16(v65)
							*(*uint8)(unsafe.Add(mBase, uint32(v81)+124)) = uint8(base.B2i32(v101 == int32(4)))
							*(*uint8)(unsafe.Add(mBase, uint32(v81)+72)) = uint8(base.B2i32(v101 == int32(1)))
							v116 = base.AtomicRmwXchg64(m, v81, int32(112), v78)
							v118 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v118)+560)) = v78
							*(*uint8)(unsafe.Add(mBase, uint32(v118)+536)) = uint8(v65)
							*(*uint8)(unsafe.Add(mBase, uint32(v118)+73)) = uint8(v65)
							*(*int64)(unsafe.Add(mBase, uint32(v118)+128)) = v78
							*(*int64)(unsafe.Add(mBase, uint32(v118)+136)) = v78
							*(*int32)(unsafe.Add(mBase, uint32(v118)+144)) = v65
							*(*int64)(unsafe.Add(mBase, uint32(v118)+544)) = v78
							*(*uint8)(unsafe.Add(mBase, uint32(v118)+552)) = uint8(v65)
							*(*int64)(unsafe.Add(mBase, uint32(v118)+576)) = v78
							*(*int64)(unsafe.Add(mBase, uint32(v118)+568)) = int64(-1)
							F_OwnLatch(m, v118+int32(20))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								F_SwitchToSharedLatch(m)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return
								} else {
									v146 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
									*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v146 + int32(548)
									v152 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
									F_PGSemaphoreReset(m, v153)
									mBase = m.M
									F_on_shmem_exit(m, int32(1116), int32(0))
									mBase = m.M
									v158 = m.ExcPending
									if v158 != 0 {
										return
									} else {
										v159 = int32(_a_F_InitProcess_5)
										v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10]))
										v163 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
										*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10])) = v163
										v167 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
										v170 = F_palloc(m, v167<<(uint(int32(2))%32))
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13])) = v170
											v175 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
											v178 = F_palloc(m, v175*int32(24))
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[14])) = v178
												v183 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13]))
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[15])) = v183
												v187 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
												v190 = F_palloc(m, v187<<(uint(int32(2))%32))
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[16])) = v190
													v195 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
													v198 = F_palloc(m, v195<<(uint(int32(2))%32))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[17])) = v198
														v203 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
														v205 = base.I32_div_s(v203, int32(2))
														v208 = F_palloc(m, v205*int32(12))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[18])) = v208
															v213 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
															v216 = F_palloc(m, v213<<(uint(int32(2))%32))
															mBase = m.M
															v217 = m.ExcPending
															if v217 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[19])) = v216
																v221 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[20])) = v221
																v226 = F_palloc(m, v221*int32(20))
																mBase = m.M
																v227 = m.ExcPending
																if v227 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[21])) = v226
																	v231 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[22])) = v231 << (uint(int32(2)) % 32)
																	v237 = F_palloc(m, v231*int32(80))
																	mBase = m.M
																	v238 = m.ExcPending
																	if v238 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10])) = v160
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[23])) = v237
																		v244 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
																		if v244 != 0 {
																			F_AttachSharedMemoryStructs(m)
																			mBase = m.M
																			v246 = m.ExcPending
																			if v246 != 0 {
																				return
																			} else {
																				m.G0 = v6 + int32(16)
																				return
																			}
																		} else {
																			m.G0 = v6 + int32(16)
																			return
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
							v251 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
							v252 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v251))), uint32(v252))
							v256 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
							F_errstart_cold(m, int32(22), v252)
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_InitProcess_6))
								mBase = m.M
								v263 = m.ExcPending
								if v263 != 0 {
									return
								} else {
									if v256 == int32(6) {
										v302 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[24]))
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = v302
										F_errmsg(m, int32(_a_F_InitProcess_7), v6)
										mBase = m.M
										v306 = m.ExcPending
										if v306 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitProcess_1), int32(454), int32(_a_F_InitProcess_2))
											mBase = m.M
											v311 = m.ExcPending
											if v311 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errmsg(m, int32(_a_F_InitProcess_8), int32(0))
										mBase = m.M
										v269 = m.ExcPending
										if v269 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitProcess_1), int32(457), int32(_a_F_InitProcess_2))
											mBase = m.M
											v274 = m.ExcPending
											if v274 != 0 {
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
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
				v21 = v19 - int32(3)
				if base.Ui32(v21) <= base.Ui32(int32(4)) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v21<<(uint(int32(2))%32))+uint32(_c_F_InitProcess[4])))
					v28 = v26
				} else {
					v28 = int32(20)
				}
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
				v31 = v28 + v30
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
				v36 = base.AtomicRmwXchg32(m, v33, int32(0), int32(1))
				if v36 != 0 {
					v38 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
					F_s_lock(m, v38, int32(_a_F_InitProcess_1), int32(432), int32(_a_F_InitProcess_2))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
						*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[6])) = v46
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
						v50 = int32(0)
						if base.B2i32(v49 == v50)|base.B2i32(v49 == v31) == v50 {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v57
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
							*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
							v61 = int32(_a_F_InitProcess_3)
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1])) = v49
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
							v65 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v64))), uint32(v65))
							v68 = int32(_a_F_InitProcess_4)
							v70 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
							v72 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
							v76 = base.I32_div_s(v70-v73, int32(640))
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7])) = v76
							v78 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v70))) = v78
							v81 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v81)+612)) = v65
							*(*uint8)(unsafe.Add(mBase, uint32(v81)+608)) = uint8(v65)
							*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v65
							*(*int64)(unsafe.Add(mBase, uint32(v81)+36)) = v78
							v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[8]))
							*(*int32)(unsafe.Add(mBase, uint32(v81)+44)) = v91
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7]))
							*(*int64)(unsafe.Add(mBase, uint32(v81)+56)) = v78
							*(*int32)(unsafe.Add(mBase, uint32(v81)+52)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = v78
							v101 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v81)+120)) = v65
							*(*int64)(unsafe.Add(mBase, uint32(v81)+92)) = v78
							*(*uint16)(unsafe.Add(mBase, uint32(v81)+74)) = uint16(v65)
							*(*uint8)(unsafe.Add(mBase, uint32(v81)+124)) = uint8(base.B2i32(v101 == int32(4)))
							*(*uint8)(unsafe.Add(mBase, uint32(v81)+72)) = uint8(base.B2i32(v101 == int32(1)))
							v116 = base.AtomicRmwXchg64(m, v81, int32(112), v78)
							v118 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
							*(*int64)(unsafe.Add(mBase, uint32(v118)+560)) = v78
							*(*uint8)(unsafe.Add(mBase, uint32(v118)+536)) = uint8(v65)
							*(*uint8)(unsafe.Add(mBase, uint32(v118)+73)) = uint8(v65)
							*(*int64)(unsafe.Add(mBase, uint32(v118)+128)) = v78
							*(*int64)(unsafe.Add(mBase, uint32(v118)+136)) = v78
							*(*int32)(unsafe.Add(mBase, uint32(v118)+144)) = v65
							*(*int64)(unsafe.Add(mBase, uint32(v118)+544)) = v78
							*(*uint8)(unsafe.Add(mBase, uint32(v118)+552)) = uint8(v65)
							*(*int64)(unsafe.Add(mBase, uint32(v118)+576)) = v78
							*(*int64)(unsafe.Add(mBase, uint32(v118)+568)) = int64(-1)
							F_OwnLatch(m, v118+int32(20))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								F_SwitchToSharedLatch(m)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return
								} else {
									v146 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
									*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v146 + int32(548)
									v152 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
									F_PGSemaphoreReset(m, v153)
									mBase = m.M
									F_on_shmem_exit(m, int32(1116), int32(0))
									mBase = m.M
									v158 = m.ExcPending
									if v158 != 0 {
										return
									} else {
										v159 = int32(_a_F_InitProcess_5)
										v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10]))
										v163 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
										*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10])) = v163
										v167 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
										v170 = F_palloc(m, v167<<(uint(int32(2))%32))
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13])) = v170
											v175 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
											v178 = F_palloc(m, v175*int32(24))
											mBase = m.M
											v179 = m.ExcPending
											if v179 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[14])) = v178
												v183 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13]))
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[15])) = v183
												v187 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
												v190 = F_palloc(m, v187<<(uint(int32(2))%32))
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[16])) = v190
													v195 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
													v198 = F_palloc(m, v195<<(uint(int32(2))%32))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[17])) = v198
														v203 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
														v205 = base.I32_div_s(v203, int32(2))
														v208 = F_palloc(m, v205*int32(12))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[18])) = v208
															v213 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
															v216 = F_palloc(m, v213<<(uint(int32(2))%32))
															mBase = m.M
															v217 = m.ExcPending
															if v217 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[19])) = v216
																v221 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[20])) = v221
																v226 = F_palloc(m, v221*int32(20))
																mBase = m.M
																v227 = m.ExcPending
																if v227 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[21])) = v226
																	v231 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[22])) = v231 << (uint(int32(2)) % 32)
																	v237 = F_palloc(m, v231*int32(80))
																	mBase = m.M
																	v238 = m.ExcPending
																	if v238 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10])) = v160
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[23])) = v237
																		v244 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
																		if v244 != 0 {
																			F_AttachSharedMemoryStructs(m)
																			mBase = m.M
																			v246 = m.ExcPending
																			if v246 != 0 {
																				return
																			} else {
																				m.G0 = v6 + int32(16)
																				return
																			}
																		} else {
																			m.G0 = v6 + int32(16)
																			return
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
							v251 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
							v252 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v251))), uint32(v252))
							v256 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
							F_errstart_cold(m, int32(22), v252)
							mBase = m.M
							v260 = m.ExcPending
							if v260 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_InitProcess_6))
								mBase = m.M
								v263 = m.ExcPending
								if v263 != 0 {
									return
								} else {
									if v256 == int32(6) {
										v302 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[24]))
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = v302
										F_errmsg(m, int32(_a_F_InitProcess_7), v6)
										mBase = m.M
										v306 = m.ExcPending
										if v306 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitProcess_1), int32(454), int32(_a_F_InitProcess_2))
											mBase = m.M
											v311 = m.ExcPending
											if v311 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									} else {
										F_errmsg(m, int32(_a_F_InitProcess_8), int32(0))
										mBase = m.M
										v269 = m.ExcPending
										if v269 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitProcess_1), int32(457), int32(_a_F_InitProcess_2))
											mBase = m.M
											v274 = m.ExcPending
											if v274 != 0 {
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
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+68))
					*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[6])) = v46
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					v50 = int32(0)
					if base.B2i32(v49 == v50)|base.B2i32(v49 == v31) == v50 {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v56)+4)) = v57
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
						*(*int32)(unsafe.Add(mBase, uint32(v57))) = v59
						v61 = int32(_a_F_InitProcess_3)
						*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1])) = v49
						v64 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
						v65 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v64))), uint32(v65))
						v68 = int32(_a_F_InitProcess_4)
						v70 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
						v72 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
						v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
						v76 = base.I32_div_s(v70-v73, int32(640))
						*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7])) = v76
						v78 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v70))) = v78
						v81 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v81)+612)) = v65
						*(*uint8)(unsafe.Add(mBase, uint32(v81)+608)) = uint8(v65)
						*(*int32)(unsafe.Add(mBase, uint32(v81)+16)) = v65
						*(*int64)(unsafe.Add(mBase, uint32(v81)+36)) = v78
						v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[8]))
						*(*int32)(unsafe.Add(mBase, uint32(v81)+44)) = v91
						v94 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7]))
						*(*int64)(unsafe.Add(mBase, uint32(v81)+56)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(v81)+52)) = v94
						*(*int64)(unsafe.Add(mBase, uint32(v81)+64)) = v78
						v101 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v81)+120)) = v65
						*(*int64)(unsafe.Add(mBase, uint32(v81)+92)) = v78
						*(*uint16)(unsafe.Add(mBase, uint32(v81)+74)) = uint16(v65)
						*(*uint8)(unsafe.Add(mBase, uint32(v81)+124)) = uint8(base.B2i32(v101 == int32(4)))
						*(*uint8)(unsafe.Add(mBase, uint32(v81)+72)) = uint8(base.B2i32(v101 == int32(1)))
						v116 = base.AtomicRmwXchg64(m, v81, int32(112), v78)
						v118 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
						*(*int64)(unsafe.Add(mBase, uint32(v118)+560)) = v78
						*(*uint8)(unsafe.Add(mBase, uint32(v118)+536)) = uint8(v65)
						*(*uint8)(unsafe.Add(mBase, uint32(v118)+73)) = uint8(v65)
						*(*int64)(unsafe.Add(mBase, uint32(v118)+128)) = v78
						*(*int64)(unsafe.Add(mBase, uint32(v118)+136)) = v78
						*(*int32)(unsafe.Add(mBase, uint32(v118)+144)) = v65
						*(*int64)(unsafe.Add(mBase, uint32(v118)+544)) = v78
						*(*uint8)(unsafe.Add(mBase, uint32(v118)+552)) = uint8(v65)
						*(*int64)(unsafe.Add(mBase, uint32(v118)+576)) = v78
						*(*int64)(unsafe.Add(mBase, uint32(v118)+568)) = int64(-1)
						F_OwnLatch(m, v118+int32(20))
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return
						} else {
							F_SwitchToSharedLatch(m)
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return
							} else {
								v146 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
								*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v146 + int32(548)
								v152 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
								v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
								F_PGSemaphoreReset(m, v153)
								mBase = m.M
								F_on_shmem_exit(m, int32(1116), int32(0))
								mBase = m.M
								v158 = m.ExcPending
								if v158 != 0 {
									return
								} else {
									v159 = int32(_a_F_InitProcess_5)
									v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10]))
									v163 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
									*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10])) = v163
									v167 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
									v170 = F_palloc(m, v167<<(uint(int32(2))%32))
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13])) = v170
										v175 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
										v178 = F_palloc(m, v175*int32(24))
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[14])) = v178
											v183 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13]))
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[15])) = v183
											v187 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
											v190 = F_palloc(m, v187<<(uint(int32(2))%32))
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[16])) = v190
												v195 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
												v198 = F_palloc(m, v195<<(uint(int32(2))%32))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[17])) = v198
													v203 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
													v205 = base.I32_div_s(v203, int32(2))
													v208 = F_palloc(m, v205*int32(12))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[18])) = v208
														v213 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
														v216 = F_palloc(m, v213<<(uint(int32(2))%32))
														mBase = m.M
														v217 = m.ExcPending
														if v217 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[19])) = v216
															v221 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[20])) = v221
															v226 = F_palloc(m, v221*int32(20))
															mBase = m.M
															v227 = m.ExcPending
															if v227 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[21])) = v226
																v231 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[22])) = v231 << (uint(int32(2)) % 32)
																v237 = F_palloc(m, v231*int32(80))
																mBase = m.M
																v238 = m.ExcPending
																if v238 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10])) = v160
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[23])) = v237
																	v244 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
																	if v244 != 0 {
																		F_AttachSharedMemoryStructs(m)
																		mBase = m.M
																		v246 = m.ExcPending
																		if v246 != 0 {
																			return
																		} else {
																			m.G0 = v6 + int32(16)
																			return
																		}
																	} else {
																		m.G0 = v6 + int32(16)
																		return
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
						v251 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5]))
						v252 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v251))), uint32(v252))
						v256 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
						F_errstart_cold(m, int32(22), v252)
						mBase = m.M
						v260 = m.ExcPending
						if v260 != 0 {
							return
						} else {
							F_errcode(m, int32(_a_F_InitProcess_6))
							mBase = m.M
							v263 = m.ExcPending
							if v263 != 0 {
								return
							} else {
								if v256 == int32(6) {
									v302 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[24]))
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v302
									F_errmsg(m, int32(_a_F_InitProcess_7), v6)
									mBase = m.M
									v306 = m.ExcPending
									if v306 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_InitProcess_1), int32(454), int32(_a_F_InitProcess_2))
										mBase = m.M
										v311 = m.ExcPending
										if v311 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errmsg(m, int32(_a_F_InitProcess_8), int32(0))
									mBase = m.M
									v269 = m.ExcPending
									if v269 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_InitProcess_1), int32(457), int32(_a_F_InitProcess_2))
										mBase = m.M
										v274 = m.ExcPending
										if v274 != 0 {
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
		F_errstart_cold(m, int32(23), int32(0))
		mBase = m.M
		v278 = m.ExcPending
		if v278 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_InitProcess_9), int32(0))
			mBase = m.M
			v282 = m.ExcPending
			if v282 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_InitProcess_1), int32(399), int32(_a_F_InitProcess_2))
				mBase = m.M
				v287 = m.ExcPending
				if v287 != 0 {
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
func F_InitializeSessionUserIdStandalone(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	v4 = int32(10)
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0])) = v4
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[1])) = v4
	v10 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])) = uint8(v10)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])))
	if v13 == int32(0) {
		v17 = int32(10)
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v17
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v17
		F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), int32(_a_F_InitializeSessionUserIdStandalone_1), int32(0), int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[6]))
			if v29 == int32(0) {
				v57 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v57)
				v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
				if v60 != 0 {
					*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v60
					*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v60
					v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
					if v69 != 0 {
						v70 = int32(_a_F_InitializeSessionUserIdStandalone_1)
					} else {
						v70 = int32(_a_F_InitializeSessionUserIdStandalone_2)
					}
					F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), v70, int32(0), int32(1))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			} else {
				v34 = F_SearchSysCache1(m, int32(11), int32(10))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					if v34 == int32(0) {
						v57 = int32(0)
						*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v57)
						v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
						if v60 != 0 {
							*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v60
							*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v60
							v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
							if v69 != 0 {
								v70 = int32(_a_F_InitializeSessionUserIdStandalone_1)
							} else {
								v70 = int32(_a_F_InitializeSessionUserIdStandalone_2)
							}
							F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), v70, int32(0), int32(1))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return
							} else {
								return
							}
						} else {
							return
						}
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
						v43 = F_pstrdup(m, v38+v39+int32(4))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v34)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								if v43 == int32(0) {
									v57 = int32(0)
									*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v57)
									v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
									if v60 != 0 {
										*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v60
										*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v60
										v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
										if v69 != 0 {
											v70 = int32(_a_F_InitializeSessionUserIdStandalone_1)
										} else {
											v70 = int32(_a_F_InitializeSessionUserIdStandalone_2)
										}
										F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), v70, int32(0), int32(1))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											return
										}
									} else {
										return
									}
								} else {
									F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_3), v43, int32(4), int32(10))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return
									} else {
										v57 = int32(0)
										*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v57)
										v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
										if v60 != 0 {
											*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v60
											*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v60
											v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
											if v69 != 0 {
												v70 = int32(_a_F_InitializeSessionUserIdStandalone_1)
											} else {
												v70 = int32(_a_F_InitializeSessionUserIdStandalone_2)
											}
											F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), v70, int32(0), int32(1))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return
											} else {
												return
											}
										} else {
											return
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
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[6]))
		if v29 == int32(0) {
			v57 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v57)
			v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
			if v60 != 0 {
				*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v60
				*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v60
				v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
				if v69 != 0 {
					v70 = int32(_a_F_InitializeSessionUserIdStandalone_1)
				} else {
					v70 = int32(_a_F_InitializeSessionUserIdStandalone_2)
				}
				F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), v70, int32(0), int32(1))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		} else {
			v34 = F_SearchSysCache1(m, int32(11), int32(10))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				if v34 == int32(0) {
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v57)
					v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
					if v60 != 0 {
						*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v60
						*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v60
						v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
						if v69 != 0 {
							v70 = int32(_a_F_InitializeSessionUserIdStandalone_1)
						} else {
							v70 = int32(_a_F_InitializeSessionUserIdStandalone_2)
						}
						F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), v70, int32(0), int32(1))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							return
						}
					} else {
						return
					}
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
					v43 = F_pstrdup(m, v38+v39+int32(4))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_ReleaseCatCache(m, v34)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							if v43 == int32(0) {
								v57 = int32(0)
								*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v57)
								v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
								if v60 != 0 {
									*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v60
									*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v60
									v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
									if v69 != 0 {
										v70 = int32(_a_F_InitializeSessionUserIdStandalone_1)
									} else {
										v70 = int32(_a_F_InitializeSessionUserIdStandalone_2)
									}
									F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), v70, int32(0), int32(1))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
										return
									} else {
										return
									}
								} else {
									return
								}
							} else {
								F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_3), v43, int32(4), int32(10))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v57 = int32(0)
									*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v57)
									v60 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
									if v60 != 0 {
										*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v60
										*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v60
										v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
										if v69 != 0 {
											v70 = int32(_a_F_InitializeSessionUserIdStandalone_1)
										} else {
											v70 = int32(_a_F_InitializeSessionUserIdStandalone_2)
										}
										F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_0), v70, int32(0), int32(1))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return
										} else {
											return
										}
									} else {
										return
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
func F_InvalidationCallback(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InvalidationCallback[0])) = uint8(v5)
	v8 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_InvalidationCallback[1])) = uint8(v8)
	return
}
func F_IoWorkerMain(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v91 int32
	_ = v91
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v157 int32
	_ = v157
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v223 int32
	_ = v223
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v256 int32
	_ = v256
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v392 int32
	_ = v392
	var v398 int32
	_ = v398
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v501 int64
	_ = v501
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v591 int32
	_ = v591
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v640 int32
	_ = v640
	var v641 int64
	_ = v641
	var v644 int64
	_ = v644
	var v646 int64
	_ = v646
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v661 int32
	_ = v661
	var v662 int64
	_ = v662
	var v665 int64
	_ = v665
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v695 int32
	_ = v695
	var v699 int32
	_ = v699
	var v707 int64
	_ = v707
	var v712 int64
	_ = v712
	var v714 int64
	_ = v714
	var v725 int32
	_ = v725
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v805 int32
	_ = v805
	var v809 int32
	_ = v809
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v970 int32
	_ = v970
	var v971 int64
	_ = v971
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	v11 = m.G0
	v13 = v11 - int32(368)
	m.G0 = v13
	v18 = int32(-1)
	v20 = int32(0)
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v18 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v970 = int32(m.ExcTag)
	v971 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v970 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[0])) = int32(12)
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+204)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v33
	*(*int64)(unsafe.Add(mBase, uint32(v13)+192)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v33
	F_AuxiliaryProcessMainCommon(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v553 = v20
	goto L9
L9:
	;
	if v553 != 0 {
		goto L161
	} else {
		goto L162
	}
L10:
	;
	v44 = int32(914)
	v46 = m.G0
	v48 = v46 - int32(32)
	m.G0 = v48
	switch int32(916) {
	case 0, 2:
		v58 = v44
		goto L12
	default:
		goto L13
	}
L11:
	;
	v77 = int32(295)
	v79 = m.G0
	v81 = v79 - int32(32)
	m.G0 = v81
	switch int32(297) {
	case 0, 2:
		v91 = v77
		goto L18
	default:
		goto L19
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v58
	F_sigemptyset(m, v48+int32(16))
	mBase = m.M
	goto L15
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[1])) = v44
	v58 = int32(_a_F_IoWorkerMain_0)
	goto L12
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = int32(268435456)
	v72 = F___sigaction(m, int32(1), v48+int32(12), int32(0))
	mBase = m.M
	m.G0 = v48 + int32(32)
	goto L11
L17:
	;
	v110 = int32(-2)
	v112 = m.G0
	v114 = v112 - int32(32)
	m.G0 = v114
	switch int32(0) {
	case 0, 2:
		v124 = v110
		goto L24
	default:
		goto L25
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = v91
	F_sigemptyset(m, v81+int32(16))
	mBase = m.M
	goto L21
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[2])) = v77
	v91 = int32(_a_F_IoWorkerMain_0)
	goto L18
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+24)) = int32(268435456)
	v105 = F___sigaction(m, int32(2), v81+int32(12), int32(0))
	mBase = m.M
	m.G0 = v81 + int32(32)
	goto L17
L23:
	;
	v143 = int32(-2)
	v145 = m.G0
	v147 = v145 - int32(32)
	m.G0 = v147
	switch int32(0) {
	case 0, 2:
		v157 = v143
		goto L30
	default:
		goto L31
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+12)) = v124
	F_sigemptyset(m, v114+int32(16))
	mBase = m.M
	goto L27
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[3])) = v110
	v124 = int32(_a_F_IoWorkerMain_0)
	goto L24
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(268435456)
	v138 = F___sigaction(m, int32(15), v114+int32(12), int32(0))
	mBase = m.M
	m.G0 = v114 + int32(32)
	goto L23
L29:
	;
	v176 = int32(-2)
	v178 = m.G0
	v180 = v178 - int32(32)
	m.G0 = v180
	switch int32(0) {
	case 0, 2:
		v190 = v176
		goto L36
	default:
		goto L37
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v157
	F_sigemptyset(m, v147+int32(16))
	mBase = m.M
	goto L33
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[4])) = v143
	v157 = int32(_a_F_IoWorkerMain_0)
	goto L30
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = int32(268435456)
	v171 = F___sigaction(m, int32(14), v147+int32(12), int32(0))
	mBase = m.M
	m.G0 = v147 + int32(32)
	goto L29
L35:
	;
	v209 = int32(917)
	v211 = m.G0
	v213 = v211 - int32(32)
	m.G0 = v213
	switch int32(919) {
	case 0, 2:
		v223 = v209
		goto L42
	default:
		goto L43
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v190
	F_sigemptyset(m, v180+int32(16))
	mBase = m.M
	goto L39
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[5])) = v176
	v190 = int32(_a_F_IoWorkerMain_0)
	goto L36
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+24)) = int32(268435456)
	v204 = F___sigaction(m, int32(13), v180+int32(12), int32(0))
	mBase = m.M
	m.G0 = v180 + int32(32)
	goto L35
L41:
	;
	v242 = int32(916)
	v244 = m.G0
	v246 = v244 - int32(32)
	m.G0 = v246
	switch int32(918) {
	case 0, 2:
		v256 = v242
		goto L48
	default:
		goto L49
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+12)) = v223
	F_sigemptyset(m, v213+int32(16))
	mBase = m.M
	goto L45
L43:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[6])) = v209
	v223 = int32(_a_F_IoWorkerMain_0)
	goto L42
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+24)) = int32(268435456)
	v237 = F___sigaction(m, int32(10), v213+int32(12), int32(0))
	mBase = m.M
	m.G0 = v213 + int32(32)
	goto L41
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7])) = int32(-1)
	v278 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[8]))
	v282 = F_LWLockAcquire(m, v278+int32(_a_F_IoWorkerMain_1), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L6
	} else {
		goto L53
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = v256
	F_sigemptyset(m, v246+int32(16))
	mBase = m.M
	goto L51
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[9])) = v242
	v256 = int32(_a_F_IoWorkerMain_0)
	goto L48
L51:
	;
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246)+24)) = int32(268435456)
	v270 = F___sigaction(m, int32(12), v246+int32(12), int32(0))
	mBase = m.M
	m.G0 = v246 + int32(32)
	goto L47
L53:
	;
	v284 = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[10]))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+12)))
	if v287 == v284 {
		v493 = v286
		v494 = v284
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v286)))
	*(*int64)(unsafe.Add(mBase, uint32(v286))) = v501 | int64(1)<<(uint(base.I64_extend_i32_u(v500))%64)
	v511 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v286+v500<<(uint(int32(3))%32))+8)) = v511
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[8]))
	F_LWLockRelease(m, v514+int32(_a_F_IoWorkerMain_1))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L6
	} else {
		goto L154
	}
L55:
	;
	v495 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v493)+12)) = uint8(v495)
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7])) = v494
	v500 = v494
	goto L54
L56:
	;
	v290 = int32(1)
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+20)))
	if v291 != v290 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v493 = v286 + int32(8)
	v494 = v290
	goto L55
L58:
	;
	goto L59
L59:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+28)))
	if v296 != int32(1) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v493 = v286 + int32(16)
	v494 = int32(2)
	goto L55
L61:
	;
	goto L62
L62:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+36)))
	if v302 != int32(1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v493 = v286 + int32(24)
	v494 = int32(3)
	goto L55
L64:
	;
	goto L65
L65:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+44)))
	if v308 != int32(1) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v493 = v286 + int32(32)
	v494 = int32(4)
	goto L55
L67:
	;
	goto L68
L68:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+52)))
	if v314 != int32(1) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v493 = v286 + int32(40)
	v494 = int32(5)
	goto L55
L70:
	;
	goto L71
L71:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+60)))
	if v320 != int32(1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v493 = v286 + int32(48)
	v494 = int32(6)
	goto L55
L73:
	;
	goto L74
L74:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+68)))
	if v326 != int32(1) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v493 = v286 + int32(56)
	v494 = int32(7)
	goto L55
L76:
	;
	goto L77
L77:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+76)))
	if v332 != int32(1) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v493 = v286 - int32(-64)
	v494 = int32(8)
	goto L55
L79:
	;
	goto L80
L80:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+84)))
	if v338 != int32(1) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v493 = v286 + int32(72)
	v494 = int32(9)
	goto L55
L82:
	;
	goto L83
L83:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+92)))
	if v344 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v493 = v286 + int32(80)
	v494 = int32(10)
	goto L55
L85:
	;
	goto L86
L86:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+100)))
	if v350 != int32(1) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v493 = v286 + int32(88)
	v494 = int32(11)
	goto L55
L88:
	;
	goto L89
L89:
	;
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+108)))
	if v356 != int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v493 = v286 + int32(96)
	v494 = int32(12)
	goto L55
L91:
	;
	goto L92
L92:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+116)))
	if v362 != int32(1) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v493 = v286 + int32(104)
	v494 = int32(13)
	goto L55
L94:
	;
	goto L95
L95:
	;
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+124)))
	if v368 != int32(1) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v493 = v286 + int32(112)
	v494 = int32(14)
	goto L55
L97:
	;
	goto L98
L98:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+132)))
	if v374 != int32(1) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v493 = v286 + int32(120)
	v494 = int32(15)
	goto L55
L100:
	;
	goto L101
L101:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+140)))
	if v380 != int32(1) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v493 = v286 + int32(128)
	v494 = int32(16)
	goto L55
L103:
	;
	goto L104
L104:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+148)))
	if v386 != int32(1) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v493 = v286 + int32(136)
	v494 = int32(17)
	goto L55
L106:
	;
	goto L107
L107:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+156)))
	if v392 != int32(1) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	v493 = v286 + int32(144)
	v494 = int32(18)
	goto L55
L109:
	;
	goto L110
L110:
	;
	v398 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+164)))
	if v398 != int32(1) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v493 = v286 + int32(152)
	v494 = int32(19)
	goto L55
L112:
	;
	goto L113
L113:
	;
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+172)))
	if v404 != int32(1) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v493 = v286 + int32(160)
	v494 = int32(20)
	goto L55
L115:
	;
	goto L116
L116:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+180)))
	if v410 != int32(1) {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v493 = v286 + int32(168)
	v494 = int32(21)
	goto L55
L118:
	;
	goto L119
L119:
	;
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+188)))
	if v416 != int32(1) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v493 = v286 + int32(176)
	v494 = int32(22)
	goto L55
L121:
	;
	goto L122
L122:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+196)))
	if v422 != int32(1) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v493 = v286 + int32(184)
	v494 = int32(23)
	goto L55
L124:
	;
	goto L125
L125:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+204)))
	if v428 != int32(1) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v493 = v286 + int32(192)
	v494 = int32(24)
	goto L55
L127:
	;
	goto L128
L128:
	;
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+212)))
	if v434 != int32(1) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v493 = v286 + int32(200)
	v494 = int32(25)
	goto L55
L130:
	;
	goto L131
L131:
	;
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+220)))
	if v440 != int32(1) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v493 = v286 + int32(208)
	v494 = int32(26)
	goto L55
L133:
	;
	goto L134
L134:
	;
	v446 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+228)))
	if v446 != int32(1) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v493 = v286 + int32(216)
	v494 = int32(27)
	goto L55
L136:
	;
	goto L137
L137:
	;
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+236)))
	if v452 != int32(1) {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v493 = v286 + int32(224)
	v494 = int32(28)
	goto L55
L139:
	;
	goto L140
L140:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+244)))
	if v458 != int32(1) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v493 = v286 + int32(232)
	v494 = int32(29)
	goto L55
L142:
	;
	goto L143
L143:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+252)))
	if v464 != int32(1) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v493 = v286 + int32(240)
	v494 = int32(30)
	goto L55
L145:
	;
	goto L146
L146:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+260)))
	if v470 != int32(1) {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	v493 = v286 + int32(248)
	v494 = int32(31)
	goto L55
L148:
	;
	goto L149
L149:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	if v477 != int32(-1) {
		v500 = v477
		goto L54
	} else {
		goto L150
	}
L150:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	F_errmsg_internal(m, int32(_a_F_IoWorkerMain_2), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L6
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_IoWorkerMain_3), int32(355), int32(_a_F_IoWorkerMain_4))
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	goto L3
L154:
	;
	F_on_shmem_exit(m, int32(1069), int32(0))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L6
	} else {
		goto L155
	}
L155:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v524
	v527 = v13 + int32(48)
	v531 = F_pg_sprintf(m, v527, int32(_a_F_IoWorkerMain_5), v13+int32(32))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	v533 = F_strlen(m, v527)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v13)+196)) = int32(1070)
	v537 = int32(_a_F_IoWorkerMain_6)
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[12])) = v13 + int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+192)) = v538
	goto L157
L157:
	;
	v545 = v13 + int32(208)
	*(*int32)(unsafe.Add(mBase, uint32(v545)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v13 + int32(36)
	goto L160
L158:
	;
	v553 = int32(0)
	goto L9
L160:
	;
	goto L158
L161:
	;
	v554 = int32(_a_F_IoWorkerMain_7)
	v556 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[13])) = v556 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[12])) = int32(0)
	F_EmitErrorReport(m)
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L6
	} else {
		goto L164
	}
L162:
	;
	goto L163
L163:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[14])) = v13 + int32(208)
	F_pgmem_sigprocmask(m, int32(_a_F_IoWorkerMain_8), int32(0))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L6
	} else {
		goto L171
	}
L164:
	;
	F_LWLockReleaseAll(m)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L6
	} else {
		goto L165
	}
L165:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v13)+204))
	if v567 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[15])) = v569
	v571 = int32(_a_F_IoWorkerMain_9)
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[16])) = v573 + int32(1)
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v13)+204))
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
	F_pgaio_io_process_completion(m, v577, int32(0)-v579)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L6
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	F_proc_exit(m, int32(1))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L6
	} else {
		goto L170
	}
L169:
	;
	v583 = int32(_a_F_IoWorkerMain_9)
	v585 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[16]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[16])) = v585 - int32(1)
	goto L168
L170:
	;
	goto L3
L171:
	;
	v601 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[17]))
	if v601 == int32(0) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	goto L175
L173:
	;
	goto L174
L174:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v13)+192))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[12])) = v955
	F_proc_exit(m, int32(0))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L6
	} else {
		goto L250
	}
L175:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[8]))
	v619 = F_LWLockAcquire(m, v615+int32(_a_F_IoWorkerMain_1), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L6
	} else {
		goto L177
	}
L176:
	;
	goto L174
L177:
	;
	v622 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[18]))
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v622)+12))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v622)+8))
	if v623 == v624 {
		goto L184
	} else {
		goto L185
	}
L178:
	;
	v929 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[19]))
	if v929 != 0 {
		goto L241
	} else {
		goto L242
	}
L179:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[20]))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+24))
	v832 = v829 + v629<<(uint(int32(7))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+204)) = v832
	*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = v832
	v837 = F_errstart(m, int32(11), int32(0))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L6
	} else {
		goto L221
	}
L180:
	;
	v812 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[8]))
	F_LWLockRelease(m, v812+int32(_a_F_IoWorkerMain_1))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L6
	} else {
		goto L220
	}
L181:
	;
	v699 = int32(0)
	v707 = v646
	goto L196
L182:
	;
	if v690 <= int32(0) {
		goto L180
	} else {
		goto L193
	}
L183:
	;
	v686 = int32(2)
	v687 = v648 - v649
	if base.Ui32(v686) <= base.Ui32(v687) {
		v695 = v686
		goto L181
	} else {
		goto L192
	}
L184:
	;
	v661 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[10]))
	v662 = *(*int64)(unsafe.Add(mBase, uint32(v661)))
	v665 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7])))
	*(*int64)(unsafe.Add(mBase, uint32(v661))) = v662 | int64(1)<<(uint(v665)%64)
	v670 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[8]))
	F_LWLockRelease(m, v670+int32(_a_F_IoWorkerMain_1))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L6
	} else {
		goto L189
	}
L185:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v622+v623<<(uint(int32(2))%32))+16))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	v631 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v622)+12)) = (v630 - v631) & (v623 + v631)
	if v629 == int32(-1) {
		goto L184
	} else {
		goto L186
	}
L186:
	;
	v640 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[10]))
	v641 = *(*int64)(unsafe.Add(mBase, uint32(v640)))
	v644 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7])))
	v646 = v641 & base.I64_rotl(int64(-2), v644)
	*(*int64)(unsafe.Add(mBase, uint32(v640))) = v646
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v622)+8))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v622)+12))
	if base.Ui32(v649) <= base.Ui32(v648) {
		goto L183
	} else {
		goto L187
	}
L187:
	;
	v651 = int32(2)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v622)))
	if base.Ui32(v651) <= base.Ui32(v652+(v648-v649)) {
		v695 = v651
		goto L181
	} else {
		goto L188
	}
L188:
	;
	v690 = v652 + v648 - v649
	goto L182
L189:
	;
	v676 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[11]))
	v680 = F_WaitLatch(m, v676, int32(33), int32(-1), int32(83886086))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L6
	} else {
		goto L190
	}
L190:
	;
	v683 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = int32(0)
	goto L191
L191:
	;
	goto L178
L192:
	;
	v690 = v687
	goto L182
L193:
	;
	v695 = v690
	goto L181
L194:
	;
	v749 = int32(0)
	goto L203
L195:
	;
	v737 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[8]))
	F_LWLockRelease(m, v737+int32(_a_F_IoWorkerMain_1))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L6
	} else {
		goto L201
	}
L196:
	;
	if v707 == int64(0) {
		goto L195
	} else {
		goto L198
	}
L197:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[8]))
	F_LWLockRelease(m, v731+int32(_a_F_IoWorkerMain_1))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L6
	} else {
		goto L200
	}
L198:
	;
	v712 = base.I64_ctz(v707)
	v714 = v707 & base.I64_rotl(int64(-2), v712)
	*(*int64)(unsafe.Add(mBase, uint32(v640))) = v714
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v640+int32(8)+base.I32_wrap_i64(v712)<<(uint(int32(3))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(40)+v699<<(uint(int32(2))%32)))) = v725
	v728 = v699 + int32(1)
	if v728 != v695 {
		v699 = v728
		v707 = v714
		goto L196
	} else {
		goto L199
	}
L199:
	;
	goto L197
L200:
	;
	v745 = v695
	goto L194
L201:
	;
	if v699 == int32(0) {
		goto L179
	} else {
		goto L202
	}
L202:
	;
	v745 = v699
	goto L194
L203:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v13+int32(40)+v749<<(uint(int32(2))%32))))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)))
	if v765 != 0 {
		goto L206
	} else {
		goto L207
	}
L204:
	;
	goto L179
L205:
	;
	v809 = v749 + int32(1)
	if v809 != v745 {
		v749 = v809
		goto L203
	} else {
		goto L219
	}
L206:
	;
	goto L205
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v764))) = int32(1)
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v764)+4))
	if v768 == int32(0) {
		goto L206
	} else {
		goto L208
	}
L208:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v764)+12))
	if v771 == int32(0) {
		goto L206
	} else {
		goto L209
	}
L209:
	;
	v775 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[21]))
	if v775 == v771 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v777 = m.G0
	v779 = v777 - int32(16)
	m.G0 = v779
	v782 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[22]))
	if v782 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	goto L212
L212:
	;
	v805 = F_pgmem_kill(m, v771, int32(23))
	mBase = m.M
	goto L206
L213:
	;
	m.G0 = v779 + int32(16)
	goto L205
L214:
	;
	v785 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v779)+15)) = uint8(v785)
	goto L215
L215:
	;
	v789 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[23]))
	v793 = F_write(m, v789, v779+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v793 {
		goto L213
	} else {
		goto L217
	}
L216:
	;
	goto L213
L217:
	;
	v797 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[15]))
	if v797 == int32(27) {
		goto L215
	} else {
		goto L218
	}
L218:
	;
	goto L216
L219:
	;
	goto L204
L220:
	;
	goto L179
L221:
	;
	if v837 != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	F_errhidestmt(m)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L6
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v889 = int32(_a_F_IoWorkerMain_7)
	v891 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[13])) = v891 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = int32(44)
	v897 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832)+1)))
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v897<<(uint(int32(2))%32))+uint32(_c_F_IoWorkerMain[24])))
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v900)))
	m.T0[v901].(func(*base.Module, int32))(m, v832)
	mBase = m.M
	v903 = m.ExcPending
	if v903 != 0 {
		goto L6
	} else {
		goto L239
	}
L225:
	;
	F_errhidecontext(m)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L6
	} else {
		goto L226
	}
L226:
	;
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[20]))
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v844)+24))
	goto L227
L227:
	;
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832)+2)))
	if base.Ui32(v849) <= base.Ui32(int32(2)) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	v857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832)+1)))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v857<<(uint(int32(2))%32))+uint32(_c_F_IoWorkerMain[24])))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v860)+8))
	goto L232
L229:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v849<<(uint(int32(2))%32))+uint32(_c_F_IoWorkerMain[25])))
	v856 = v854
	goto L231
L230:
	;
	v856 = int32(0)
	goto L231
L231:
	;
	goto L228
L232:
	;
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v832))))
	if base.Ui32(v862) <= base.Ui32(int32(7)) {
		goto L234
	} else {
		goto L235
	}
L233:
	;
	v871 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(16)))) = v871
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v869
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v861
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v856
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = (v832 - v845) >> (uint(int32(7)) % 32)
	F_errmsg_internal(m, int32(_a_F_IoWorkerMain_10), v13)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L6
	} else {
		goto L237
	}
L234:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v862<<(uint(int32(2))%32))+uint32(_c_F_IoWorkerMain[26])))
	v869 = v867
	goto L236
L235:
	;
	v869 = int32(0)
	goto L236
L236:
	;
	goto L233
L237:
	;
	F_errfinish(m, int32(_a_F_IoWorkerMain_3), int32(513), int32(_a_F_IoWorkerMain_11))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L6
	} else {
		goto L238
	}
L238:
	;
	goto L224
L239:
	;
	v904 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+188)) = v904
	*(*int32)(unsafe.Add(mBase, uint32(v13)+204)) = v904
	F_pgaio_io_perform_synchronously(m, v832)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L6
	} else {
		goto L240
	}
L240:
	;
	v910 = int32(_a_F_IoWorkerMain_7)
	v912 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[13]))
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[13])) = v912 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+200)) = int32(0)
	goto L178
L241:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v931 = m.ExcPending
	if v931 != 0 {
		goto L6
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v933 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[27]))
	if v933 != 0 {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	goto L243
L245:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[27])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L6
	} else {
		goto L248
	}
L246:
	;
	goto L247
L247:
	;
	v941 = *(*int32)(unsafe.Add(mBase, _c_F_IoWorkerMain[17]))
	if v941 == int32(0) {
		goto L175
	} else {
		goto L249
	}
L248:
	;
	goto L247
L249:
	;
	goto L176
L250:
	;
	goto L5
L251:
	;
	v975 = int32(v971)
	m.G0 = v13
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v975)+4))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v975)))
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v978)))
	if v13+int32(36) == v981 {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	m.ExcPending = 1
	goto L258
L253:
	;
	if v985 != 0 {
		v18 = v985
		v20 = v977
		goto L1
	} else {
		goto L257
	}
L254:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(v978)+4))
	v985 = v983
	goto L256
L255:
	;
	v985 = int32(0)
	goto L256
L256:
	;
	goto L253
L257:
	;
	F___wasm_longjmp(m, v978, v977)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	return
L259:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_IsPinnedObject(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(base.B2i32(l0 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_IsPinnedObject_0)) < base.Ui32(l1)) == int32(0)) & ((base.B2i32(l0 != int32(2615)) | base.B2i32(l1 != int32(2200))) & base.B2i32(l0 != int32(1262)))
}
func F_IsReservedName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v2 = int32(0)
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v3 != int32(112) {
		v12 = v2
	} else {
		v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		if v6 != int32(103) {
			v12 = v2
		} else {
			v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v12 = base.B2i32(v9 == int32(95))
		}
	}
	return v12
}
func F_IsSquashableConstant(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	v4 = l0
	goto L4
L1:
	;
	return v243
L2:
	;
	v243 = int32(1)
	goto L1
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
	v20 = int32(1)
	if base.Ui32(v20) < base.Ui32(v19-v20) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui32(int32(2)) <= base.Ui32(v7-int32(27)) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	return base.B2i32(v15 == int32(0))
L6:
	;
	switch v7 - int32(7) {
	case 0:
		goto L2
	case 1:
		goto L9
	default:
		v243 = int32(0)
		goto L1
	case 8:
		goto L3
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	v4 = v14
	goto L4
L8:
	;
	goto L5
L9:
	;
	goto L8
L10:
	;
	return int32(0)
L11:
	;
	goto L12
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
	if base.Ui32(int32(_a_F_IsSquashableConstant_0)) < base.Ui32(v26) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	goto L15
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
	if v31 == int32(0) {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v35 <= int32(0) {
		v243 = int32(1)
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v40 == int32(7) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v141 < int32(2) {
		goto L2
	} else {
		goto L56
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_IsSquashableConstant[0]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_IsSquashableConstant[1]))
	v49 = m.G0
	v52 = v48 - (v49 - int32(1))
	v54 = v52 >> (uint(int32(31)) % 32)
	goto L20
L20:
	;
	if base.B2i32(v46 < v52^v54-v54)&base.B2i32(v48 != int32(0)) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	goto L23
L23:
	;
	v65 = v39
	goto L28
L24:
	;
	if v137 != 0 {
		goto L18
	} else {
		goto L55
	}
L25:
	;
	v137 = v133
	goto L24
L26:
	;
	v133 = int32(1)
	goto L25
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v65)+16))
	v80 = int32(1)
	if base.Ui32(v80) < base.Ui32(v79-v80) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if base.Ui32(int32(2)) <= base.Ui32(v68-int32(27)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v137 = base.B2i32(v76 == int32(0))
	goto L24
L30:
	;
	switch v68 - int32(7) {
	case 0:
		goto L26
	case 1:
		goto L33
	default:
		v133 = int32(0)
		goto L25
	case 8:
		goto L27
	}
L31:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v65 = v75
	goto L28
L32:
	;
	goto L29
L33:
	;
	goto L32
L34:
	;
	v137 = int32(0)
	goto L24
L35:
	;
	goto L36
L36:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if base.Ui32(int32(_a_F_IsSquashableConstant_0)) < base.Ui32(v85) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v137 = int32(0)
	goto L24
L38:
	;
	goto L39
L39:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v65)+28))
	if v89 == int32(0) {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v93 <= int32(0) {
		v133 = int32(1)
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v98 == int32(7) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v106 < int32(2) {
		goto L26
	} else {
		goto L48
	}
L43:
	;
	v101 = F_stack_is_too_deep(m)
	mBase = m.M
	if v101 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v137 = int32(0)
	goto L24
L45:
	;
	goto L46
L46:
	;
	v103 = F_IsSquashableConstant(m, v97)
	mBase = m.M
	if v103 != 0 {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v137 = int32(0)
	goto L24
L48:
	;
	v109 = int32(1)
	goto L49
L49:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112+v109<<(uint(int32(2))%32))))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	if v117 == int32(7) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v133 = v122
	goto L25
L51:
	;
	v122 = int32(1)
	v124 = v109 + v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v124 < v125 {
		v109 = v124
		goto L49
	} else {
		goto L54
	}
L52:
	;
	v120 = F_IsSquashableConstant(m, v116)
	mBase = m.M
	if v120 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v137 = int32(0)
	goto L24
L54:
	;
	goto L50
L55:
	;
	return int32(0)
L56:
	;
	v144 = int32(1)
	goto L57
L57:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v147+v144<<(uint(int32(2))%32))))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)))
	if v152 == int32(7) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v243 = v232
	goto L1
L59:
	;
	v232 = int32(1)
	v234 = v144 + v232
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v234 < v235 {
		v144 = v234
		goto L57
	} else {
		goto L93
	}
L60:
	;
	v157 = v151
	goto L65
L61:
	;
	if v229 != 0 {
		goto L59
	} else {
		goto L92
	}
L62:
	;
	v229 = v225
	goto L61
L63:
	;
	v225 = int32(1)
	goto L62
L64:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	v172 = int32(1)
	if base.Ui32(v172) < base.Ui32(v171-v172) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if base.Ui32(int32(2)) <= base.Ui32(v160-int32(27)) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v229 = base.B2i32(v168 == int32(0))
	goto L61
L67:
	;
	switch v160 - int32(7) {
	case 0:
		goto L63
	case 1:
		goto L70
	default:
		v225 = int32(0)
		goto L62
	case 8:
		goto L64
	}
L68:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v157 = v167
	goto L65
L69:
	;
	goto L66
L70:
	;
	goto L69
L71:
	;
	v229 = int32(0)
	goto L61
L72:
	;
	goto L73
L73:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if base.Ui32(int32(_a_F_IsSquashableConstant_0)) < base.Ui32(v177) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v229 = int32(0)
	goto L61
L75:
	;
	goto L76
L76:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v157)+28))
	if v181 == int32(0) {
		goto L63
	} else {
		goto L77
	}
L77:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v185 <= int32(0) {
		v225 = int32(1)
		goto L62
	} else {
		goto L78
	}
L78:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	if v190 == int32(7) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v198 < int32(2) {
		goto L63
	} else {
		goto L85
	}
L80:
	;
	v193 = F_stack_is_too_deep(m)
	mBase = m.M
	if v193 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v229 = int32(0)
	goto L61
L82:
	;
	goto L83
L83:
	;
	v195 = F_IsSquashableConstant(m, v189)
	mBase = m.M
	if v195 != 0 {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v229 = int32(0)
	goto L61
L85:
	;
	v201 = int32(1)
	goto L86
L86:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v201<<(uint(int32(2))%32))))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	if v209 == int32(7) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v225 = v214
	goto L62
L88:
	;
	v214 = int32(1)
	v216 = v201 + v214
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v181)+4))
	if v216 < v217 {
		v201 = v216
		goto L86
	} else {
		goto L91
	}
L89:
	;
	v212 = F_IsSquashableConstant(m, v208)
	mBase = m.M
	if v212 != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v229 = int32(0)
	goto L61
L91:
	;
	goto L87
L92:
	;
	return int32(0)
L93:
	;
	goto L58
}
func F_IsThereCollationInNamespace(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_IsThereCollationInNamespace[0]))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v14 = F_SearchSysCacheExists(m, int32(15), l0, v12, l1, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 == int32(0) {
			v21 = F_SearchSysCacheExists(m, int32(15), l0, int32(-1), l1, int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				if v21 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						F_errcode(m, int32(_a_F_IsThereCollationInNamespace_0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v56 = F_get_namespace_name(m, l1)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
								F_errmsg(m, int32(_a_F_IsThereCollationInNamespace_1), v7+int32(16))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_IsThereCollationInNamespace_2), int32(417), int32(_a_F_IsThereCollationInNamespace_3))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
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
					m.G0 = v7 + int32(32)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_IsThereCollationInNamespace_0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_IsThereCollationInNamespace[0]))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v36 = F_get_namespace_name(m, l1)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v36
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(_a_F_IsThereCollationInNamespace_4), v7)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_IsThereCollationInNamespace_2), int32(407), int32(_a_F_IsThereCollationInNamespace_3))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
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
func F_IsThereOpClassInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v9 int32
	_ = v9
	Fn13855(m, l0, l1, l2, int32(_a_F_IsThereOpClassInNamespace_0), int32(1843), int32(_a_F_IsThereOpClassInNamespace_1), int32(13))
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_IsThereOpFamilyInNamespace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v9 int32
	_ = v9
	Fn13855(m, l0, l1, l2, int32(_a_F_IsThereOpFamilyInNamespace_0), int32(1866), int32(_a_F_IsThereOpFamilyInNamespace_1), int32(41))
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_icregexeqsel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14010(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_ind_fetch_func(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v5 = v4 * l1
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v6))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10+v5<<(uint(int32(2))%32))))
	return v14
}
func F_indonesian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v83 int32
	_ = v83
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v260 int32
	_ = v260
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v519 int32
	_ = v519
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v545 int32
	_ = v545
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v615 int32
	_ = v615
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v659 int32
	_ = v659
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v697 int32
	_ = v697
	var v704 int32
	_ = v704
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v2
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = v10
	goto L3
L1:
	;
	if int32(0) <= v127 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v127 = v99
	goto L1
L3:
	;
	if v23 <= v32 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v127 = int32(-1)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v39 = int32(1)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32+v24))))
	if base.Ui32(v41) < base.Ui32(int32(192)) {
		v98 = v41
		v99 = v39
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if int32(117) < v98 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v45 = v32 + int32(1)
	if v45 == v23 {
		v98 = v41
		v99 = v39
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v24))))
	v50 = v48 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v41) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v24))))
	v66 = v64 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v41) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v54 = v32 + int32(2)
	if v54 != v23 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v98 = v41<<(uint(int32(6))%32)&int32(1984) | v50
	v99 = int32(2)
	goto L8
L15:
	;
	goto L14
L16:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v70))))
	v98 = v83&int32(63) | (v41<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v50<<(uint(int32(12))%32) | v66<<(uint(int32(6))%32))
	v99 = int32(4)
	goto L8
L17:
	;
	v70 = v32 + int32(3)
	if v70 != v23 {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v98 = v41<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v50<<(uint(int32(6))%32) | v66
	v99 = int32(3)
	goto L8
L20:
	;
	goto L19
L21:
	;
	v116 = v99 + v32
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v116
	v32 = v116
	goto L3
L22:
	;
	v103 = v98 - int32(97)
	if v103 < int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v103)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v109)>>(uint(v103&int32(7))%32))&int32(1) != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L21
L26:
	;
	v131 = v127
	goto L29
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v271 < int32(3) {
		v790 = v2
		goto L57
	} else {
		goto L58
	}
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v136 + v131
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v140 + int32(1)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = v155
	goto L33
L30:
	;
	goto L28
L31:
	;
	if int32(0) <= v260 {
		v131 = v260
		goto L29
	} else {
		goto L56
	}
L32:
	;
	v260 = v232
	goto L31
L33:
	;
	if v156 <= v165 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v260 = int32(-1)
	goto L31
L36:
	;
	goto L37
L37:
	;
	v172 = int32(1)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+v157))))
	if base.Ui32(v174) < base.Ui32(int32(192)) {
		v231 = v174
		v232 = v172
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if int32(117) < v231 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v178 = v165 + int32(1)
	if v178 == v156 {
		v231 = v174
		v232 = v172
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178+v157))))
	v183 = v181 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v174) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187+v157))))
	v199 = v197 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v174) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v187 = v165 + int32(2)
	if v187 != v156 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v231 = v174<<(uint(int32(6))%32)&int32(1984) | v183
	v232 = int32(2)
	goto L38
L45:
	;
	goto L44
L46:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v203))))
	v231 = v216&int32(63) | (v174<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v183<<(uint(int32(12))%32) | v199<<(uint(int32(6))%32))
	v232 = int32(4)
	goto L38
L47:
	;
	v203 = v165 + int32(3)
	if v203 != v156 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v231 = v174<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v183<<(uint(int32(6))%32) | v199
	v232 = int32(3)
	goto L38
L50:
	;
	goto L49
L51:
	;
	v249 = v232 + v165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v249
	v165 = v249
	goto L33
L52:
	;
	v236 = v231 - int32(97)
	if v236 < int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v236)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v242)>>(uint(v236&int32(7))%32))&int32(1) != 0 {
		goto L32
	} else {
		goto L54
	}
L54:
	;
	goto L51
L56:
	;
	goto L30
L57:
	;
	return v790
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = int32(0)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v276
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v278
	if v278-int32(2) <= v276 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v314)+4))
	if v319 < int32(3) {
		goto L71
	} else {
		goto L72
	}
L60:
	;
	v314 = v270
	v316 = int32(0)
	goto L59
L61:
	;
	goto L62
L62:
	;
	v285 = int32(0)
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286+v278-int32(1)))))
	switch v290 - int32(104) {
	case 0, 6:
		goto L63
	default:
		v314 = v270
		v316 = v285
		goto L59
	}
L63:
	;
	v295 = F_find_among_b(m, l0, int32(_a_F_indonesian_UTF_8_stem_2), int32(3))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return int32(0)
L65:
	;
	if v295 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v314 = v301
	v316 = v285
	goto L59
L67:
	;
	goto L68
L68:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v302
	v304 = F_slice_del(m, l0)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L64
	} else {
		goto L69
	}
L69:
	;
	if v304 < int32(0) {
		v790 = v304
		goto L57
	} else {
		goto L70
	}
L70:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	v310 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v308)+4)) = v309 - v310
	v314 = v308
	v316 = v310
	goto L59
L71:
	;
	return int32(0)
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v317
	v326 = v317 - int32(1)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v326 <= v327 {
		v355 = v314
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v357
	v359 = int32(0)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v355)+4))
	if v360 < int32(3) {
		v790 = v359
		goto L57
	} else {
		goto L83
	}
L75:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v329+v326))))
	if base.B2i32(v331 != int32(117))&base.B2i32(v331 != int32(97)) != 0 {
		v355 = v314
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v339 = F_find_among_b(m, l0, int32(_a_F_indonesian_UTF_8_stem_3), int32(3))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L64
	} else {
		goto L77
	}
L77:
	;
	if v339 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v355 = v343
	goto L74
L79:
	;
	goto L80
L80:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v344
	v346 = F_slice_del(m, l0)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L64
	} else {
		goto L81
	}
L81:
	;
	if v346 < int32(0) {
		v790 = v346
		goto L57
	} else {
		goto L82
	}
L82:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v350)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v350)+4)) = v351 - int32(1)
	v355 = v350
	goto L74
L83:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v363
	v366 = v363 + int32(1)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v367 <= v366 {
		v723 = v359
		goto L86
	} else {
		goto L87
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v357
	v790 = int32(1)
	goto L57
L85:
	;
	if v728 != 0 {
		goto L169
	} else {
		goto L170
	}
L86:
	;
	v728 = v723
	goto L85
L87:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369+v366))))
	switch v371 - int32(101) {
	case 0, 4:
		goto L88
	default:
		v723 = v359
		goto L86
	}
L88:
	;
	v376 = F_find_among(m, l0, int32(_a_F_indonesian_UTF_8_stem_4), int32(12))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L64
	} else {
		goto L89
	}
L89:
	;
	if v376 == int32(0) {
		v723 = v359
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v380
	v382 = int32(1)
	switch v376 - v382 {
	case 0:
		goto L96
	case 1:
		goto L95
	case 2:
		goto L94
	case 3:
		goto L93
	case 4:
		goto L92
	case 5:
		goto L91
	default:
		v723 = v382
		goto L86
	}
L91:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v579))) = int32(3)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v579)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v579)+4)) = v582 - int32(1)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L139
L92:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v440 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v439))) = v440
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v439)+4)) = v442 - v440
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L107
L93:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = int32(3)
	v429 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_UTF_8_stem_5))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L64
	} else {
		goto L103
	}
L94:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v410 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = v410
	v414 = F_slice_from_s(m, l0, v410, int32(_a_F_indonesian_UTF_8_stem_6))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L64
	} else {
		goto L101
	}
L95:
	;
	v397 = F_slice_del(m, l0)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L64
	} else {
		goto L99
	}
L96:
	;
	v385 = F_slice_del(m, l0)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L64
	} else {
		goto L97
	}
L97:
	;
	if v385 < int32(0) {
		v723 = v385
		goto L86
	} else {
		goto L98
	}
L98:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v390 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v392 - v390
	v728 = v390
	goto L85
L99:
	;
	if v397 < int32(0) {
		v723 = v397
		goto L86
	} else {
		goto L100
	}
L100:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v401))) = int32(3)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)+4))
	v405 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v401)+4)) = v404 - v405
	v728 = v405
	goto L85
L101:
	;
	if v414 < int32(0) {
		v723 = v414
		goto L86
	} else {
		goto L102
	}
L102:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v420 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v418)+4)) = v419 - v420
	v728 = v420
	goto L85
L103:
	;
	if v429 < int32(0) {
		v723 = v429
		goto L86
	} else {
		goto L104
	}
L104:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v433)+4))
	v435 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+4)) = v434 - v435
	v728 = v435
	goto L85
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v446
	if v564 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L106:
	;
	v564 = v557
	goto L105
L107:
	;
	if v459 <= v446 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v557 = int32(0)
	goto L106
L109:
	;
	v564 = int32(-1)
	goto L105
L110:
	;
	goto L111
L111:
	;
	v475 = int32(1)
	v477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v446+v460))))
	if base.Ui32(v477) < base.Ui32(int32(192)) {
		v534 = v477
		v535 = v475
		goto L112
	} else {
		goto L113
	}
L112:
	;
	if int32(117) < v534 {
		v557 = v535
		goto L106
	} else {
		goto L125
	}
L113:
	;
	v481 = v446 + int32(1)
	if v481 == v459 {
		v534 = v477
		v535 = v475
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481+v460))))
	v486 = v484 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v477) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v490+v460))))
	v502 = v500 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v477) {
		goto L121
	} else {
		goto L122
	}
L116:
	;
	v490 = v446 + int32(2)
	if v490 != v459 {
		goto L115
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	v534 = v477<<(uint(int32(6))%32)&int32(1984) | v486
	v535 = int32(2)
	goto L112
L119:
	;
	goto L118
L120:
	;
	v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460+v506))))
	v534 = v519&int32(63) | (v477<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v486<<(uint(int32(12))%32) | v502<<(uint(int32(6))%32))
	v535 = int32(4)
	goto L112
L121:
	;
	v506 = v446 + int32(3)
	if v506 != v459 {
		goto L120
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v534 = v477<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v486<<(uint(int32(6))%32) | v502
	v535 = int32(3)
	goto L112
L124:
	;
	goto L123
L125:
	;
	v539 = v534 - int32(97)
	if v539 < int32(0) {
		v557 = v535
		goto L106
	} else {
		goto L126
	}
L126:
	;
	v545 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v539)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v545)>>(uint(v539&int32(7))%32))&int32(1) == int32(0) {
		v557 = v535
		goto L106
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v535 + v446
	goto L128
L128:
	;
	goto L108
L129:
	;
	v728 = v578
	goto L85
L130:
	;
	v570 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_UTF_8_stem_7))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L64
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v574 = F_slice_del(m, l0)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L64
	} else {
		goto L135
	}
L133:
	;
	if v570 < int32(0) {
		v578 = v570
		goto L129
	} else {
		goto L134
	}
L134:
	;
	v723 = v382
	goto L86
L135:
	;
	if int32(0) <= v574 {
		v723 = v382
		goto L86
	} else {
		goto L136
	}
L136:
	;
	v578 = v574
	goto L129
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v586
	if v704 == int32(0) {
		goto L162
	} else {
		goto L163
	}
L138:
	;
	v704 = v697
	goto L137
L139:
	;
	if v599 <= v586 {
		goto L141
	} else {
		goto L142
	}
L140:
	;
	v697 = int32(0)
	goto L138
L141:
	;
	v704 = int32(-1)
	goto L137
L142:
	;
	goto L143
L143:
	;
	v615 = int32(1)
	v617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v586+v600))))
	if base.Ui32(v617) < base.Ui32(int32(192)) {
		v674 = v617
		v675 = v615
		goto L144
	} else {
		goto L145
	}
L144:
	;
	if int32(117) < v674 {
		v697 = v675
		goto L138
	} else {
		goto L157
	}
L145:
	;
	v621 = v586 + int32(1)
	if v621 == v599 {
		v674 = v617
		v675 = v615
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v621+v600))))
	v626 = v624 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v617) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v630+v600))))
	v642 = v640 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v617) {
		goto L153
	} else {
		goto L154
	}
L148:
	;
	v630 = v586 + int32(2)
	if v630 != v599 {
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v674 = v617<<(uint(int32(6))%32)&int32(1984) | v626
	v675 = int32(2)
	goto L144
L151:
	;
	goto L150
L152:
	;
	v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600+v646))))
	v674 = v659&int32(63) | (v617<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v626<<(uint(int32(12))%32) | v642<<(uint(int32(6))%32))
	v675 = int32(4)
	goto L144
L153:
	;
	v646 = v586 + int32(3)
	if v646 != v599 {
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v674 = v617<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v626<<(uint(int32(6))%32) | v642
	v675 = int32(3)
	goto L144
L156:
	;
	goto L155
L157:
	;
	v679 = v674 - int32(97)
	if v679 < int32(0) {
		v697 = v675
		goto L138
	} else {
		goto L158
	}
L158:
	;
	v685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v679)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v685)>>(uint(v679&int32(7))%32))&int32(1) == int32(0) {
		v697 = v675
		goto L138
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v675 + v586
	goto L160
L160:
	;
	goto L140
L161:
	;
	v723 = v720
	goto L86
L162:
	;
	v708 = int32(1)
	v711 = F_slice_from_s(m, l0, v708, int32(_a_F_indonesian_UTF_8_stem_8))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L64
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v716 = F_slice_del(m, l0)
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L64
	} else {
		goto L167
	}
L165:
	;
	if v711 < int32(0) {
		v720 = v711
		goto L161
	} else {
		goto L166
	}
L166:
	;
	v723 = v708
	goto L86
L167:
	;
	if int32(0) <= v716 {
		v723 = int32(1)
		goto L86
	} else {
		goto L168
	}
L168:
	;
	v720 = v716
	goto L161
L169:
	;
	if v728 < int32(0) {
		v790 = v728
		goto L57
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v357
	v760 = F_r_remove_second_order_prefix_2(m, l0)
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L64
	} else {
		goto L186
	}
L172:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+4))
	if v732 < int32(3) {
		goto L84
	} else {
		goto L173
	}
L173:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v735
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v737
	v739 = F_r_remove_suffix_2(m, l0)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L64
	} else {
		goto L174
	}
L174:
	;
	if v739 == int32(0) {
		goto L84
	} else {
		goto L175
	}
L175:
	;
	if v739 < int32(0) {
		v790 = v739
		goto L57
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v735
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	if v747 < int32(3) {
		goto L84
	} else {
		goto L177
	}
L177:
	;
	v750 = F_r_remove_second_order_prefix_2(m, l0)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L64
	} else {
		goto L178
	}
L178:
	;
	if int32(0) <= v750 {
		goto L84
	} else {
		goto L179
	}
L179:
	;
	if v750 < int32(0) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v756 = v750
	goto L182
L181:
	;
	v756 = v316
	goto L182
L182:
	;
	if v750 != 0 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v757 = v756
	goto L185
L184:
	;
	v757 = v316
	goto L185
L185:
	;
	return v757
L186:
	;
	if v760 < int32(0) {
		v790 = v760
		goto L57
	} else {
		goto L187
	}
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v357
	v765 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v765)+4))
	if v766 < int32(3) {
		goto L84
	} else {
		goto L188
	}
L188:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v357
	v770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v770
	v772 = F_r_remove_suffix_2(m, l0)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L64
	} else {
		goto L189
	}
L189:
	;
	if base.Ui32(int32(6)) < base.Ui32(int32(base.Ui32(v772)>>(uint(int32(31))%32))-int32(1)) {
		goto L84
	} else {
		goto L190
	}
L190:
	;
	if int32(0) <= v772 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	v783 = int32(1)
	goto L193
L192:
	;
	v783 = v772
	goto L193
L193:
	;
	return v783
}
func F_init_degree_constants(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 float64
	_ = v14
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v48 float64
	_ = v48
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v62 float64
	_ = v62
	var v67 float64
	_ = v67
	var v71 float64
	_ = v71
	var v80 float64
	_ = v80
	var v89 float64
	_ = v89
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v102 float64
	_ = v102
	var v108 int64
	_ = v108
	var v113 int32
	_ = v113
	var v126 float64
	_ = v126
	var v137 float64
	_ = v137
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v156 float64
	_ = v156
	var v161 float64
	_ = v161
	var v162 float64
	_ = v162
	var v163 float64
	_ = v163
	var v168 float64
	_ = v168
	var v174 float64
	_ = v174
	var v178 float64
	_ = v178
	var v181 float64
	_ = v181
	var v185 float64
	_ = v185
	var v192 int64
	_ = v192
	var v197 int32
	_ = v197
	var v207 float64
	_ = v207
	var v213 float64
	_ = v213
	var v244 float64
	_ = v244
	var v245 int32
	_ = v245
	var v246 float64
	_ = v246
	var v247 float64
	_ = v247
	var v261 float64
	_ = v261
	var v278 float64
	_ = v278
	var v285 int32
	_ = v285
	var v286 float64
	_ = v286
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v296 float64
	_ = v296
	var v297 float64
	_ = v297
	var v307 float64
	_ = v307
	var v311 float64
	_ = v311
	var v313 float64
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v333 float64
	_ = v333
	var v337 int32
	_ = v337
	var v338 float64
	_ = v338
	var v339 float64
	_ = v339
	var v345 float64
	_ = v345
	var v346 float64
	_ = v346
	var v348 float64
	_ = v348
	var v350 float64
	_ = v350
	var v352 float64
	_ = v352
	var v362 float64
	_ = v362
	var v364 float64
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v384 float64
	_ = v384
	var v388 int32
	_ = v388
	var v389 float64
	_ = v389
	var v390 float64
	_ = v390
	var v395 float64
	_ = v395
	var v397 float64
	_ = v397
	var v399 float64
	_ = v399
	var v402 float64
	_ = v402
	var v406 float64
	_ = v406
	var v410 float64
	_ = v410
	var v413 float64
	_ = v413
	var v417 float64
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v430 int32
	_ = v430
	var v437 float64
	_ = v437
	var v441 int32
	_ = v441
	var v442 float64
	_ = v442
	var v443 float64
	_ = v443
	var v449 float64
	_ = v449
	var v450 float64
	_ = v450
	var v452 float64
	_ = v452
	var v454 float64
	_ = v454
	var v456 float64
	_ = v456
	var v471 float64
	_ = v471
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v491 float64
	_ = v491
	var v495 int32
	_ = v495
	var v496 float64
	_ = v496
	var v497 float64
	_ = v497
	var v502 float64
	_ = v502
	var v504 float64
	_ = v504
	var v506 float64
	_ = v506
	var v509 float64
	_ = v509
	var v513 float64
	_ = v513
	var v517 float64
	_ = v517
	var v525 float64
	_ = v525
	var v530 float64
	_ = v530
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v543 int32
	_ = v543
	var v550 float64
	_ = v550
	var v554 int32
	_ = v554
	var v555 float64
	_ = v555
	var v556 float64
	_ = v556
	var v561 float64
	_ = v561
	var v563 float64
	_ = v563
	var v565 float64
	_ = v565
	var v568 float64
	_ = v568
	var v572 float64
	_ = v572
	var v576 float64
	_ = v576
	var v584 float64
	_ = v584
	var v594 float64
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v614 float64
	_ = v614
	var v618 int32
	_ = v618
	var v619 float64
	_ = v619
	var v620 float64
	_ = v620
	var v626 float64
	_ = v626
	var v627 float64
	_ = v627
	var v629 float64
	_ = v629
	var v631 float64
	_ = v631
	var v633 float64
	_ = v633
	var v644 float64
	_ = v644
	var v649 float64
	_ = v649
	var v651 float64
	_ = v651
	var v658 float64
	_ = v658
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v678 float64
	_ = v678
	var v682 int32
	_ = v682
	var v683 float64
	_ = v683
	var v684 float64
	_ = v684
	var v690 float64
	_ = v690
	var v691 float64
	_ = v691
	var v693 float64
	_ = v693
	var v695 float64
	_ = v695
	var v697 float64
	_ = v697
	var v712 float64
	_ = v712
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v725 int32
	_ = v725
	var v732 float64
	_ = v732
	var v736 int32
	_ = v736
	var v737 float64
	_ = v737
	var v738 float64
	_ = v738
	var v743 float64
	_ = v743
	var v745 float64
	_ = v745
	var v747 float64
	_ = v747
	var v750 float64
	_ = v750
	var v754 float64
	_ = v754
	var v758 float64
	_ = v758
	var v766 float64
	_ = v766
	var v768 int32
	_ = v768
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = *(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[0]))
	v20 = base.I64_reinterpret_f64(v14)
	v25 = base.I32_wrap_i64(int64(base.Ui64(v20)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v25) {
		if base.I32_wrap_i64(v20)|(v25-int32(1072693248)) == int32(0) {
			v102 = base.F64_add(base.F64_mul(v14, float64(1.5707963267948966)), float64(7.52316384526264e-37))
		} else {
			v102 = base.F64_div(float64(0), base.F64_sub(v14, v14))
		}
	} else {
		if base.Ui32(v25) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v25+int32(-1048576)) < base.Ui32(int32(1044381696)) {
				v94 = v14
				v102 = v94
			} else {
				v48 = F_R(m, base.F64_mul(v14, v14))
				mBase = m.M
				v102 = base.F64_add(base.F64_mul(v14, v48), v14)
			}
		} else {
			v55 = base.F64_mul(base.F64_sub(float64(1), base.F64_abs(v14)), float64(0.5))
			v56 = base.F64_sqrt(v55)
			v57 = F_R(m, v55)
			mBase = m.M
			if base.Ui32(int32(1072640819)) <= base.Ui32(v25) {
				v62 = base.F64_add(base.F64_mul(v56, v57), v56)
				v89 = base.F64_sub(float64(1.5707963267948966), base.F64_add(base.F64_add(v62, v62), float64(-6.123233995736766e-17)))
			} else {
				v67 = float64(0.7853981633974483)
				v71 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v56) & int64(-4294967296))
				v80 = base.F64_div(base.F64_sub(v55, base.F64_mul(v71, v71)), base.F64_add(v56, v71))
				v89 = base.F64_add(base.F64_sub(base.F64_sub(v67, base.F64_add(v71, v71)), base.F64_sub(base.F64_mul(base.F64_add(v56, v56), v57), base.F64_sub(float64(6.123233995736766e-17), base.F64_add(v80, v80)))), v67)
			}
			if v20 < int64(0) {
				v93 = base.F64_neg(v89)
			} else {
				v93 = v89
			}
			v94 = v93
			v102 = v94
		}
	}
	*(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[1])) = v102
	v108 = base.I64_reinterpret_f64(v14)
	v113 = base.I32_wrap_i64(int64(base.Ui64(v108)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1072693248)) <= base.Ui32(v113) {
		if base.I32_wrap_i64(v108)|(v113-int32(1072693248)) == int32(0) {
			if int64(0) <= v108 {
				v126 = float64(0)
			} else {
				v126 = float64(3.141592653589793)
			}
			v181 = v126
		} else {
			v181 = base.F64_div(float64(0), base.F64_sub(v14, v14))
		}
	} else {
		if base.Ui32(v113) <= base.Ui32(int32(1071644671)) {
			if base.Ui32(v113) < base.Ui32(int32(1012924417)) {
				v178 = float64(1.5707963267948966)
				v181 = v178
			} else {
				v137 = F_R(m, base.F64_mul(v14, v14))
				mBase = m.M
				v181 = base.F64_add(base.F64_sub(base.F64_sub(float64(6.123233995736766e-17), base.F64_mul(v14, v137)), v14), float64(1.5707963267948966))
			}
		} else {
			if v108 < int64(0) {
				v149 = base.F64_mul(base.F64_add(v14, float64(1)), float64(0.5))
				v150 = base.F64_sqrt(v149)
				v151 = F_R(m, v149)
				mBase = m.M
				v156 = base.F64_sub(float64(1.5707963267948966), base.F64_add(v150, base.F64_add(base.F64_mul(v150, v151), float64(-6.123233995736766e-17))))
				v181 = base.F64_add(v156, v156)
			} else {
				v161 = base.F64_mul(base.F64_sub(float64(1), v14), float64(0.5))
				v162 = base.F64_sqrt(v161)
				v163 = F_R(m, v161)
				mBase = m.M
				v168 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v162) & int64(-4294967296))
				v174 = base.F64_add(base.F64_add(base.F64_mul(v162, v163), base.F64_div(base.F64_sub(v161, base.F64_mul(v168, v168)), base.F64_add(v162, v168))), v168)
				v178 = base.F64_add(v174, v174)
				v181 = v178
			}
		}
	}
	*(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[2])) = v181
	v185 = *(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[3]))
	v192 = base.I64_reinterpret_f64(v185)
	v197 = base.I32_wrap_i64(int64(base.Ui64(v192)>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(1141899264)) <= base.Ui32(v197) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v185)&int64(9223372036854775807)) {
			v207 = v185
		} else {
			v207 = base.F64_copysign(float64(1.5707963267948966), v185)
		}
		v307 = v207
	} else {
		if base.Ui32(v197) <= base.Ui32(int32(1071382527)) {
			if base.Ui32(int32(1044381696)) <= base.Ui32(v197) {
				v244 = v185
				v245 = int32(-1)
				v246 = base.F64_mul(v244, v244)
				v247 = base.F64_mul(v246, v246)
				v261 = base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
				v278 = base.F64_mul(v246, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
				if base.Ui32(v197) <= base.Ui32(int32(1071382527)) {
					v307 = base.F64_sub(v244, base.F64_mul(v244, base.F64_add(v261, v278)))
				} else {
					v285 = v245 << (uint(int32(3)) % 32)
					v286 = *(*float64)(unsafe.Add(mBase, uint32(v285)+uint32(_c_F_init_degree_constants[4])))
					v289 = *(*float64)(unsafe.Add(mBase, uint32(v285)+uint32(_c_F_init_degree_constants[5])))
					v292 = base.F64_sub(v286, base.F64_sub(base.F64_sub(base.F64_mul(v244, base.F64_add(v261, v278)), v289), v244))
					if v192 < int64(0) {
						v296 = base.F64_neg(v292)
					} else {
						v296 = v292
					}
					v297 = v296
					v307 = v297
				}
			} else {
				v297 = v185
				v307 = v297
			}
		} else {
			v213 = base.F64_abs(v185)
			if base.Ui32(v197) <= base.Ui32(int32(1072889855)) {
				if base.Ui32(v197) <= base.Ui32(int32(1072037887)) {
					v244 = base.F64_div(base.F64_add(base.F64_add(v213, v213), float64(-1)), base.F64_add(v213, float64(2)))
					v245 = int32(0)
				} else {
					v244 = base.F64_div(base.F64_add(v213, float64(-1)), base.F64_add(v213, float64(1)))
					v245 = int32(1)
				}
			} else {
				if base.Ui32(v197) <= base.Ui32(int32(1073971199)) {
					v244 = base.F64_div(base.F64_add(v213, float64(-1.5)), base.F64_add(base.F64_mul(v213, float64(1.5)), float64(1)))
					v245 = int32(2)
				} else {
					v244 = base.F64_div(float64(-1), v213)
					v245 = int32(3)
				}
			}
			v246 = base.F64_mul(v244, v244)
			v247 = base.F64_mul(v246, v246)
			v261 = base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, float64(-0.036531572744216916)), float64(-0.058335701337905735))), float64(-0.0769187620504483))), float64(-0.11111110405462356))), float64(-0.19999999999876483)))
			v278 = base.F64_mul(v246, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, base.F64_add(base.F64_mul(v247, float64(0.016285820115365782)), float64(0.049768779946159324))), float64(0.06661073137387531))), float64(0.09090887133436507))), float64(0.14285714272503466))), float64(0.3333333333333293)))
			if base.Ui32(v197) <= base.Ui32(int32(1071382527)) {
				v307 = base.F64_sub(v244, base.F64_mul(v244, base.F64_add(v261, v278)))
			} else {
				v285 = v245 << (uint(int32(3)) % 32)
				v286 = *(*float64)(unsafe.Add(mBase, uint32(v285)+uint32(_c_F_init_degree_constants[4])))
				v289 = *(*float64)(unsafe.Add(mBase, uint32(v285)+uint32(_c_F_init_degree_constants[5])))
				v292 = base.F64_sub(v286, base.F64_sub(base.F64_sub(base.F64_mul(v244, base.F64_add(v261, v278)), v289), v244))
				if v192 < int64(0) {
					v296 = base.F64_neg(v292)
				} else {
					v296 = v292
				}
				v297 = v296
				v307 = v297
			}
		}
	}
	*(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[6])) = v307
	v311 = *(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[7]))
	v313 = base.F64_mul(v311, float64(0.017453292519943295))
	v317 = m.G0
	v319 = v317 - int32(16)
	m.G0 = v319
	v326 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v313))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v326) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v326) < base.Ui32(int32(1045430272)) {
			v352 = v313
		} else {
			v333 = F___sin(m, v313, float64(0), int32(0))
			mBase = m.M
			v352 = v333
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v326) {
			v352 = base.F64_sub(v313, v313)
		} else {
			v337 = F___rem_pio2(m, v313, v319)
			mBase = m.M
			v338 = *(*float64)(unsafe.Add(mBase, uint32(v319)+8))
			v339 = *(*float64)(unsafe.Add(mBase, uint32(v319)))
			switch v337&int32(3) - int32(1) {
			case 0:
				v346 = F___cos(m, v339, v338)
				mBase = m.M
				v352 = v346
			case 1:
				v348 = F___sin(m, v339, v338, int32(1))
				mBase = m.M
				v352 = base.F64_neg(v348)
			case 2:
				v350 = F___cos(m, v339, v338)
				mBase = m.M
				v352 = base.F64_neg(v350)
			default:
				v345 = F___sin(m, v339, v338, int32(1))
				mBase = m.M
				v352 = v345
			}
		}
	}
	m.G0 = v319 + int32(16)
	*(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[8])) = v352
	v362 = *(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[9]))
	v364 = base.F64_mul(v362, float64(0.017453292519943295))
	v368 = m.G0
	v370 = v368 - int32(16)
	m.G0 = v370
	v377 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v364))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v377) <= base.Ui32(int32(1072243195)) {
		if base.Ui32(v377) < base.Ui32(int32(1044816030)) {
			v406 = float64(1)
		} else {
			v384 = F___cos(m, v364, float64(0))
			mBase = m.M
			v406 = v384
		}
	} else {
		if base.Ui32(int32(2146435072)) <= base.Ui32(v377) {
			v406 = base.F64_sub(v364, v364)
		} else {
			v388 = F___rem_pio2(m, v364, v370)
			mBase = m.M
			v389 = *(*float64)(unsafe.Add(mBase, uint32(v370)+8))
			v390 = *(*float64)(unsafe.Add(mBase, uint32(v370)))
			switch v388&int32(3) - int32(1) {
			case 0:
				v397 = F___sin(m, v390, v389, int32(1))
				mBase = m.M
				v406 = base.F64_neg(v397)
			case 1:
				v399 = F___cos(m, v390, v389)
				mBase = m.M
				v406 = base.F64_neg(v399)
			case 2:
				v402 = F___sin(m, v390, v389, int32(1))
				mBase = m.M
				v406 = v402
			default:
				v395 = F___cos(m, v390, v389)
				mBase = m.M
				v406 = v395
			}
		}
	}
	m.G0 = v370 + int32(16)
	v410 = base.F64_sub(float64(1), v406)
	*(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[10])) = v410
	v413 = *(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[11]))
	if base.F64_le(v413, float64(30)) != 0 {
		v417 = base.F64_mul(v413, float64(0.017453292519943295))
		v421 = m.G0
		v423 = v421 - int32(16)
		m.G0 = v423
		v430 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v417))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v430) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v430) < base.Ui32(int32(1045430272)) {
				v456 = v417
			} else {
				v437 = F___sin(m, v417, float64(0), int32(0))
				mBase = m.M
				v456 = v437
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v430) {
				v456 = base.F64_sub(v417, v417)
			} else {
				v441 = F___rem_pio2(m, v417, v423)
				mBase = m.M
				v442 = *(*float64)(unsafe.Add(mBase, uint32(v423)+8))
				v443 = *(*float64)(unsafe.Add(mBase, uint32(v423)))
				switch v441&int32(3) - int32(1) {
				case 0:
					v450 = F___cos(m, v443, v442)
					mBase = m.M
					v456 = v450
				case 1:
					v452 = F___sin(m, v443, v442, int32(1))
					mBase = m.M
					v456 = base.F64_neg(v452)
				case 2:
					v454 = F___cos(m, v443, v442)
					mBase = m.M
					v456 = base.F64_neg(v454)
				default:
					v449 = F___sin(m, v443, v442, int32(1))
					mBase = m.M
					v456 = v449
				}
			}
		}
		m.G0 = v423 + int32(16)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v456
		v525 = base.F64_mul(base.F64_div(v456, v352), float64(0.5))
	} else {
		v471 = base.F64_mul(base.F64_sub(float64(90), v413), float64(0.017453292519943295))
		v475 = m.G0
		v477 = v475 - int32(16)
		m.G0 = v477
		v484 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v471))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v484) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v484) < base.Ui32(int32(1044816030)) {
				v513 = float64(1)
			} else {
				v491 = F___cos(m, v471, float64(0))
				mBase = m.M
				v513 = v491
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v484) {
				v513 = base.F64_sub(v471, v471)
			} else {
				v495 = F___rem_pio2(m, v471, v477)
				mBase = m.M
				v496 = *(*float64)(unsafe.Add(mBase, uint32(v477)+8))
				v497 = *(*float64)(unsafe.Add(mBase, uint32(v477)))
				switch v495&int32(3) - int32(1) {
				case 0:
					v504 = F___sin(m, v497, v496, int32(1))
					mBase = m.M
					v513 = base.F64_neg(v504)
				case 1:
					v506 = F___cos(m, v497, v496)
					mBase = m.M
					v513 = base.F64_neg(v506)
				case 2:
					v509 = F___sin(m, v497, v496, int32(1))
					mBase = m.M
					v513 = v509
				default:
					v502 = F___cos(m, v497, v496)
					mBase = m.M
					v513 = v502
				}
			}
		}
		m.G0 = v477 + int32(16)
		v517 = base.F64_sub(float64(1), v513)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v517
		v525 = base.F64_add(base.F64_mul(base.F64_div(v517, v410), float64(-0.5)), float64(1))
	}
	if base.F64_le(v413, float64(60)) != 0 {
		v530 = base.F64_mul(v413, float64(0.017453292519943295))
		v534 = m.G0
		v536 = v534 - int32(16)
		m.G0 = v536
		v543 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v530))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v543) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v543) < base.Ui32(int32(1044816030)) {
				v572 = float64(1)
			} else {
				v550 = F___cos(m, v530, float64(0))
				mBase = m.M
				v572 = v550
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v543) {
				v572 = base.F64_sub(v530, v530)
			} else {
				v554 = F___rem_pio2(m, v530, v536)
				mBase = m.M
				v555 = *(*float64)(unsafe.Add(mBase, uint32(v536)+8))
				v556 = *(*float64)(unsafe.Add(mBase, uint32(v536)))
				switch v554&int32(3) - int32(1) {
				case 0:
					v563 = F___sin(m, v556, v555, int32(1))
					mBase = m.M
					v572 = base.F64_neg(v563)
				case 1:
					v565 = F___cos(m, v556, v555)
					mBase = m.M
					v572 = base.F64_neg(v565)
				case 2:
					v568 = F___sin(m, v556, v555, int32(1))
					mBase = m.M
					v572 = v568
				default:
					v561 = F___cos(m, v556, v555)
					mBase = m.M
					v572 = v561
				}
			}
		}
		m.G0 = v536 + int32(16)
		v576 = base.F64_sub(float64(1), v572)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v576
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v576
		v584 = base.F64_sub(float64(1), base.F64_mul(base.F64_div(v576, v410), float64(0.5)))
		v649 = v584
		v651 = v584
	} else {
		v594 = base.F64_mul(base.F64_sub(float64(90), v413), float64(0.017453292519943295))
		v598 = m.G0
		v600 = v598 - int32(16)
		m.G0 = v600
		v607 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v594))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v607) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v607) < base.Ui32(int32(1045430272)) {
				v633 = v594
			} else {
				v614 = F___sin(m, v594, float64(0), int32(0))
				mBase = m.M
				v633 = v614
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v607) {
				v633 = base.F64_sub(v594, v594)
			} else {
				v618 = F___rem_pio2(m, v594, v600)
				mBase = m.M
				v619 = *(*float64)(unsafe.Add(mBase, uint32(v600)+8))
				v620 = *(*float64)(unsafe.Add(mBase, uint32(v600)))
				switch v618&int32(3) - int32(1) {
				case 0:
					v627 = F___cos(m, v620, v619)
					mBase = m.M
					v633 = v627
				case 1:
					v629 = F___sin(m, v620, v619, int32(1))
					mBase = m.M
					v633 = base.F64_neg(v629)
				case 2:
					v631 = F___cos(m, v620, v619)
					mBase = m.M
					v633 = base.F64_neg(v631)
				default:
					v626 = F___sin(m, v620, v619, int32(1))
					mBase = m.M
					v633 = v626
				}
			}
		}
		m.G0 = v600 + int32(16)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v633
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v633
		v644 = base.F64_mul(base.F64_div(v633, v352), float64(0.5))
		v649 = v644
		v651 = v644
	}
	*(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[12])) = base.F64_div(v525, v649)
	if base.F64_le(v413, float64(30)) != 0 {
		v658 = base.F64_mul(v413, float64(0.017453292519943295))
		v662 = m.G0
		v664 = v662 - int32(16)
		m.G0 = v664
		v671 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v658))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v671) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v671) < base.Ui32(int32(1045430272)) {
				v697 = v658
			} else {
				v678 = F___sin(m, v658, float64(0), int32(0))
				mBase = m.M
				v697 = v678
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v671) {
				v697 = base.F64_sub(v658, v658)
			} else {
				v682 = F___rem_pio2(m, v658, v664)
				mBase = m.M
				v683 = *(*float64)(unsafe.Add(mBase, uint32(v664)+8))
				v684 = *(*float64)(unsafe.Add(mBase, uint32(v664)))
				switch v682&int32(3) - int32(1) {
				case 0:
					v691 = F___cos(m, v684, v683)
					mBase = m.M
					v697 = v691
				case 1:
					v693 = F___sin(m, v684, v683, int32(1))
					mBase = m.M
					v697 = base.F64_neg(v693)
				case 2:
					v695 = F___cos(m, v684, v683)
					mBase = m.M
					v697 = base.F64_neg(v695)
				default:
					v690 = F___sin(m, v684, v683, int32(1))
					mBase = m.M
					v697 = v690
				}
			}
		}
		m.G0 = v664 + int32(16)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v697
		v766 = base.F64_mul(base.F64_div(v697, v352), float64(0.5))
	} else {
		v712 = base.F64_mul(base.F64_sub(float64(90), v413), float64(0.017453292519943295))
		v716 = m.G0
		v718 = v716 - int32(16)
		m.G0 = v718
		v725 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v712))>>(uint(int64(32))%64))) & int32(2147483647)
		if base.Ui32(v725) <= base.Ui32(int32(1072243195)) {
			if base.Ui32(v725) < base.Ui32(int32(1044816030)) {
				v754 = float64(1)
			} else {
				v732 = F___cos(m, v712, float64(0))
				mBase = m.M
				v754 = v732
			}
		} else {
			if base.Ui32(int32(2146435072)) <= base.Ui32(v725) {
				v754 = base.F64_sub(v712, v712)
			} else {
				v736 = F___rem_pio2(m, v712, v718)
				mBase = m.M
				v737 = *(*float64)(unsafe.Add(mBase, uint32(v718)+8))
				v738 = *(*float64)(unsafe.Add(mBase, uint32(v718)))
				switch v736&int32(3) - int32(1) {
				case 0:
					v745 = F___sin(m, v738, v737, int32(1))
					mBase = m.M
					v754 = base.F64_neg(v745)
				case 1:
					v747 = F___cos(m, v738, v737)
					mBase = m.M
					v754 = base.F64_neg(v747)
				case 2:
					v750 = F___sin(m, v738, v737, int32(1))
					mBase = m.M
					v754 = v750
				default:
					v743 = F___cos(m, v738, v737)
					mBase = m.M
					v754 = v743
				}
			}
		}
		m.G0 = v718 + int32(16)
		v758 = base.F64_sub(float64(1), v754)
		*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v758
		v766 = base.F64_add(base.F64_mul(base.F64_div(v758, v410), float64(-0.5)), float64(1))
	}
	v768 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_init_degree_constants[13])) = uint8(v768)
	*(*float64)(unsafe.Add(mBase, _c_F_init_degree_constants[14])) = base.F64_div(v651, v766)
	m.G0 = v10 + int32(16)
	return
}
func F_init_execution_state(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int64
	_ = v94
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v656 int32
	_ = v656
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v725 int32
	_ = v725
	var v740 int32
	_ = v740
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v768 int32
	_ = v768
	var v778 int32
	_ = v778
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v831 int32
	_ = v831
	var v834 int32
	_ = v834
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v880 int32
	_ = v880
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	F_ReleaseCachedPlan(m, v21, v22)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	if v34 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	return int32(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = int32(0)
	goto L3
L6:
	;
	m.G0 = v19 + int32(32)
	return v880
L7:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+68))
	if v477 < v479 {
		goto L114
	} else {
		goto L115
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L110
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L105
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L101
	}
L11:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+76))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v278 + int32(1)
	v286 = int32(0)
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v292 = F_GetCachedPlan(m, v282, v290, v288, v286)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L57
	}
L12:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v37 = v35
	goto L14
L13:
	;
	v37 = int32(0)
	goto L14
L14:
	;
	if v37 <= v32 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
	if v39 <= v32 {
		v880 = v29
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v255 + int32(1)
	goto L11
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v41 + int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+72)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v33)+68))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	if v49 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v52 = v50
	goto L21
L20:
	;
	v52 = int32(0)
	goto L21
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v48+v52<<(uint(int32(2))%32))))
	v57 = F_copyObjectImpl(m, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
	if v45 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v209 = int32(0)
	v212 = base.B2i32(v52+int32(1) < v46)
	if v212 == v209 {
		goto L48
	} else {
		goto L49
	}
L24:
	;
	if v153 == int32(0) {
		goto L23
	} else {
		goto L39
	}
L25:
	;
	v62 = F_CreateCommandTag(m, v57)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v137 = F_CreateCommandTag(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L36
	}
L28:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1]))
	v70 = F_AllocSetContextCreateInternal(m, v65, int32(_a_F_init_execution_state_0), int32(0), int32(1024), int32(_a_F_init_execution_state_1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v72 = int32(_a_F_init_execution_state_2)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v70
	v77 = F_palloc0(m, int32(144))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(195726186)
	v82 = F_copyObjectImpl(m, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+4)) = v82
	v87 = F_pstrdup(m, v59)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+12)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v70)+36)) = v87
	v91 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+52)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v77)+16)) = v62
	v94 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+20)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+28)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+36)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+41)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+60)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v77)+56)) = v70
	*(*int64)(unsafe.Add(mBase, uint32(v77)+68)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+76)) = v94
	*(*uint16)(unsafe.Add(mBase, uint32(v77)+84)) = uint16(v91)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+88)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v77)+96)) = v91
	*(*int64)(unsafe.Add(mBase, uint32(v77)+120)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+112)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v77)+128)) = v94
	*(*int64)(unsafe.Add(mBase, uint32(v77)+136)) = v94
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v70
	v125 = F_copyObjectImpl(m, v57)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v77)+8)) = v125
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v73
	F_AcquireRewriteLocks(m, v57, int32(1), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v134 = F_pg_rewrite_query(m, v57)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v149 = v77
	v153 = v134
	goto L24
L36:
	;
	v139 = F_CreateCachedPlan(m, v57, v59, v137)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v33)+36))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	v145 = F_pg_analyze_and_rewrite_withcb(m, v57, v141, int32(474), v143, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v149 = v139
	v153 = v145
	goto L24
L39:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	if v156 <= int32(0) {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v162 = int32(0)
	goto L41
L41:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v159+v162<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)+4))
	if v181 != int32(6) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L23
L43:
	;
	v191 = v162 + int32(1)
	if v156 != v191 {
		v162 = v191
		goto L41
	} else {
		goto L47
	}
L44:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180)+28))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	if v185 != int32(213) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	if v188 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	goto L42
L48:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
	v217 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33)+58)))
	v219 = F_check_sql_stmt_retval(m, v153, v215, v216, v217, int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	v222 = v209
	goto L50
L50:
	;
	v223 = int32(0)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	F_CompleteCachedPlan(m, v149, v153, v223, v223, v223, int32(474), v227, int32(2052), v223)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+56)) = uint8(v219)
	v222 = v33
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v149)+40)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v149)+36)) = int32(687)
	v235 = int32(_a_F_init_execution_state_2)
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1]))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v238
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v241 = F_lappend(m, v240, v149)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v241
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v236
	F_SaveCachedPlan(m, v149)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	if v52+int32(1) < v46 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+64)) = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	F_MemoryContextDelete(m, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+80)) = int32(0)
	goto L11
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v292
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v295 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v297 = v296
	goto L60
L59:
	;
	v297 = v286
	goto L60
L60:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v298 < v297 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v300 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v318 = v295
	goto L63
L63:
	;
	if v318 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v312
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	v318 = v316
	goto L63
L65:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v306 = F_MemoryContextAlloc(m, v303, v297*int32(20))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v310 = F_repalloc(m, v300, v297*int32(20))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L69
	}
L68:
	;
	v312 = v306
	goto L64
L69:
	;
	v312 = v310
	goto L64
L70:
	;
	v466 = int32(0)
	goto L7
L71:
	;
	goto L72
L72:
	;
	v322 = int32(0)
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v323 <= v322 {
		v466 = v322
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v326 = int32(0)
	v330 = v326
	v331 = v326
	v333 = v322
	goto L74
L74:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v318)+12))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v344+v330<<(uint(int32(2))%32))))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v349 != int32(6) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v466 = v406
	goto L7
L76:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+57)))
	if v385 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L77:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v348)+88))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)))
	if v353 != int32(157) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if v353 != int32(225) {
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v352)+20))
	if v379 == int32(0) {
		goto L8
	} else {
		goto L87
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v348)+88))
	v366 = F_CreateCommandName(m, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v366
	F_errmsg(m, int32(_a_F_init_execution_state_3), v19+int32(16))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_init_execution_state_4), int32(744), int32(_a_F_init_execution_state_5))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	goto L76
L88:
	;
	v388 = F_CommandIsReadOnly(m, v348)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v395 = v392 + v330*int32(20)
	if v331 != 0 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	if v388 == int32(0) {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v398 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v395)+8)) = uint16(v398)
	*(*int64)(unsafe.Add(mBase, uint32(v395))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+16)) = v398
	*(*int32)(unsafe.Add(mBase, uint32(v395)+12)) = v348
	v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+26)))
	if v405 != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v331))) = v395
	goto L93
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v395
	goto L93
L97:
	;
	v406 = v395
	goto L99
L98:
	;
	v406 = v333
	goto L99
L99:
	;
	v408 = v330 + int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	if v408 < v409 {
		v330 = v408
		v331 = v395
		v333 = v406
		goto L74
	} else {
		goto L100
	}
L100:
	;
	goto L75
L101:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_init_execution_state_6), int32(0))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_init_execution_state_4), int32(2075), int32(_a_F_init_execution_state_7))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v434 = F_CreateCommandName(m, v348)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v434
	F_errmsg(m, int32(_a_F_init_execution_state_8), v19)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_init_execution_state_4), int32(752), int32(_a_F_init_execution_state_5))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_init_execution_state_9), int32(0))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_init_execution_state_4), int32(737), int32(_a_F_init_execution_state_5))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L4
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	v880 = int32(1)
	goto L6
L115:
	;
	goto L116
L116:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v478)+48))
	if v482 == int32(2278) {
		v834 = v478
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v848 = int32(1)
	v849 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+55)))
	if v849 != v848 {
		goto L201
	} else {
		goto L202
	}
L118:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v485 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+24))
	if v486 == v488 {
		v834 = v478
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v490 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L121
L123:
	;
	v505 = int32(_a_F_init_execution_state_2)
	v506 = *(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v504
	v511 = F_MakeTupleTableSlot(m, int32(0), int32(_a_F_init_execution_state_10))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L4
	} else {
		goto L129
	}
L124:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v498 = F_AllocSetContextCreateInternal(m, v493, int32(_a_F_init_execution_state_11), int32(0), int32(1024), int32(_a_F_init_execution_state_12))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	F_MemoryContextReset(m, v490)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L4
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v498
	v504 = v498
	goto L123
L128:
	;
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v504 = v503
	goto L123
L129:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v282)+60))
	if v513 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v671)+60))
	if v672 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L131:
	;
	v656 = int32(0)
	goto L130
L132:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	if v516 <= int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v520 = v516 & int32(3)
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v513)+12))
	v522 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v516) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v613 == int32(0) {
		goto L131
	} else {
		goto L160
	}
L135:
	;
	v531 = v522
	v532 = v522
	v539 = int32(0)
	goto L138
L136:
	;
	v569 = v522
	v570 = v522
	goto L137
L137:
	;
	v585 = v569
	v586 = v570
	v587 = v522
	goto L154
L138:
	;
	v548 = v521 + v531<<(uint(int32(2))%32)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+12))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v548)+8))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+24)))
	if v553 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v520 == int32(0) {
		v613 = v560
		goto L134
	} else {
		goto L153
	}
L140:
	;
	v554 = v552
	goto L142
L141:
	;
	v554 = v532
	goto L142
L142:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+24)))
	if v555 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v556 = v551
	goto L145
L144:
	;
	v556 = v554
	goto L145
L145:
	;
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550)+24)))
	if v557 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v558 = v550
	goto L148
L147:
	;
	v558 = v556
	goto L148
L148:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v549)+24)))
	if v559 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v560 = v549
	goto L151
L150:
	;
	v560 = v558
	goto L151
L151:
	;
	v561 = int32(4)
	v562 = v531 + v561
	v564 = v539 + v561
	if v564 != v516&int32(2147483644) {
		v531 = v562
		v532 = v560
		v539 = v564
		goto L138
	} else {
		goto L152
	}
L152:
	;
	goto L139
L153:
	;
	v569 = v562
	v570 = v560
	goto L137
L154:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v521+v585<<(uint(int32(2))%32))))
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v603)+24)))
	if v604 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v613 = v605
	goto L134
L156:
	;
	v605 = v603
	goto L158
L157:
	;
	v605 = v586
	goto L158
L158:
	;
	v606 = int32(1)
	v609 = v587 + v606
	if v609 != v520 {
		v585 = v585 + v606
		v586 = v605
		v587 = v609
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L155
L160:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	if v629 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v613)+76))
	v656 = v632
	goto L130
L162:
	;
	goto L163
L163:
	;
	if base.Ui32(int32(3)) < base.Ui32(v629-int32(2)) {
		goto L131
	} else {
		goto L164
	}
L164:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v613)+96))
	if v637 != 0 {
		v656 = v637
		goto L130
	} else {
		goto L165
	}
L165:
	;
	goto L131
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v813
	*(*int32)(unsafe.Add(mBase, uint32(v813)+4)) = int32(0)
	v817 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v818 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817)+56)))
	if v818 == int32(1) {
		goto L197
	} else {
		goto L198
	}
L167:
	;
	v795 = F_ExecInitJunkFilter(m, v656, v511)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L4
	} else {
		goto L196
	}
L168:
	;
	v675 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v671)+56)))
	if v675 != int32(1) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v678 = int32(0)
	if v511 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v672)))
	if int32(0) < v687 {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	F_ExecSetSlotDescriptor(m, v511, v672)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L4
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v684 = F_MakeTupleTableSlot(m, v672, int32(_a_F_init_execution_state_13))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L4
	} else {
		goto L175
	}
L174:
	;
	v686 = v511
	goto L170
L175:
	;
	v686 = v684
	goto L170
L176:
	;
	v692 = F_palloc0(m, v687<<(uint(int32(1))%32))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	v778 = v678
	goto L178
L178:
	;
	v787 = F_palloc0(m, int32(20))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L4
	} else {
		goto L195
	}
L179:
	;
	if v656 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v656)+12))
	v695 = v694
	goto L182
L181:
	;
	v695 = v678
	goto L182
L182:
	;
	v698 = v695
	v703 = v678
	goto L183
L183:
	;
	v715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v672+v703<<(uint(int32(4))%32))+29)))
	if v715 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v778 = v692
	goto L178
L185:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v656)+12))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	v725 = v698
	goto L188
L186:
	;
	v753 = v698
	goto L187
L187:
	;
	v768 = v703 + int32(1)
	if v768 != v687 {
		v698 = v753
		v703 = v768
		goto L183
	} else {
		goto L194
	}
L188:
	;
	v740 = v725 + int32(4)
	if base.Ui32(v740) < base.Ui32(v718+v719<<(uint(int32(2))%32)) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v749 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v744)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v692+v703<<(uint(int32(1))%32)))) = uint16(v749)
	v753 = v743
	goto L187
L190:
	;
	v743 = v740
	goto L192
L191:
	;
	v743 = int32(0)
	goto L192
L192:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v725)))
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+26)))
	if v745 != 0 {
		v725 = v743
		goto L188
	} else {
		goto L193
	}
L193:
	;
	goto L189
L194:
	;
	goto L184
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v787)+16)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v787)+12)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v787)+8)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v787)+4)) = v656
	*(*int32)(unsafe.Add(mBase, uint32(v787))) = int32(385)
	v813 = v787
	goto L166
L196:
	;
	v813 = v795
	goto L166
L197:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v821)+16))
	v823 = *(*int32)(unsafe.Add(mBase, uint32(v822)+12))
	v824 = F_BlessTupleDesc(m, v823)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L4
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v826)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v827
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v506
	v831 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v834 = v831
	goto L117
L200:
	;
	goto L199
L201:
	;
	if v466 == int32(0) {
		v880 = v848
		goto L6
	} else {
		goto L206
	}
L202:
	;
	v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+56)))
	if v852 != 0 {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v834)+48))
	v854 = F_type_is_rowtype(m, v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	if v854 == int32(0) {
		goto L201
	} else {
		goto L205
	}
L205:
	;
	v858 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v858)
	goto L201
L206:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v862 == int32(0) {
		v880 = v848
		goto L6
	} else {
		goto L207
	}
L207:
	;
	v865 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v466)+8)) = uint8(v865)
	v867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v867 != v865 {
		v880 = v848
		goto L6
	} else {
		goto L208
	}
L208:
	;
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v466)+12))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v870)+4))
	if v871 != int32(1) {
		v880 = v848
		goto L6
	} else {
		goto L209
	}
L209:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v870)+25)))
	if v874 != 0 {
		v880 = v848
		goto L6
	} else {
		goto L210
	}
L210:
	;
	v875 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v466)+9)) = uint8(v875)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v875)
	v880 = v848
	goto L6
}
func F_init_sexpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v16 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v19 = v17
	} else {
		v19 = int32(0)
	}
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[0]))
	v24 = F_object_aclcheck(m, int32(1255), l0, v22, int64(128))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return
	} else {
		if v24 != 0 {
			v27 = F_get_func_name(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_aclcheck_error(m, v24, int32(19), v27)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[1]))
					if v32 != 0 {
						F_RunFunctionExecuteHook(m, l0)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return
						} else {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							if v35 != 0 {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
								if int32(101) <= v36 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v162 = m.ExcPending
									if v162 != 0 {
										return
									} else {
										F_errcode(m, int32(50856197))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return
										} else {
											v166 = int32(100)
											*(*int32)(unsafe.Add(mBase, uint32(v14))) = v166
											F_errmsg_plural(m, int32(_a_F_init_sexpr_0), int32(_a_F_init_sexpr_1), v166, v14)
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_init_sexpr_2), int32(721), int32(_a_F_init_sexpr_3))
												mBase = m.M
												v177 = m.ExcPending
												if v177 != 0 {
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
									v40 = l3 + int32(16)
									F_fmgr_info_cxt(m, l0, v40, l5)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
										v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
											*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
											v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
											v54 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
											v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
											v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
											v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
											*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
											v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
											*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
											v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
											if l6|base.B2i32(v66&int32(1) == v54) == v54 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return
													} else {
														F_errmsg(m, int32(_a_F_init_sexpr_4), int32(0))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															if l4 != 0 {
																v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
																v86 = F_exprLocation(m, l2)
																mBase = m.M
																F_executor_errposition(m, v85, v86)
																mBase = m.M
																v88 = m.ExcPending
																if v88 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															} else {
																F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
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
												v94 = int32(0)
												if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
													v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
													v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
													mBase = m.M
													v109 = m.ExcPending
													if v109 != 0 {
														return
													} else {
														v110 = int32(_a_F_init_sexpr_5)
														v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
														*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
														v114 = int32(1)
														if base.Ui32(v108-v114) <= base.Ui32(v114) {
															v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
															v119 = F_CreateTupleDescCopy(m, v118)
															mBase = m.M
															v120 = m.ExcPending
															if v120 != 0 {
																return
															} else {
																v121 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
																*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
																*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
																v152 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
																*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
																m.G0 = v14 + int32(16)
																return
															}
														} else {
															switch v108 {
															case 0:
																v125 = F_CreateTemplateTupleDesc(m, int32(1))
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
																	v129 = int32(0)
																	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
																	F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
																	mBase = m.M
																	v134 = m.ExcPending
																	if v134 != 0 {
																		return
																	} else {
																		v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
																		v136 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v136)
																		*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v135
																		*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
																		v152 = int32(0)
																		*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
																		*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
																		m.G0 = v14 + int32(16)
																		return
																	}
																}
															default:
																*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
																*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
																v152 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
																*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
																m.G0 = v14 + int32(16)
																return
															case 3:
																v139 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v139)
																*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
																*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
																v152 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
																*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
																m.G0 = v14 + int32(16)
																return
															}
														}
													}
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
													v152 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
													*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
													m.G0 = v14 + int32(16)
													return
												}
											}
										}
									}
								}
							} else {
								v40 = l3 + int32(16)
								F_fmgr_info_cxt(m, l0, v40, l5)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
									v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
										*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
										v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										v54 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
										v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
										if l6|base.B2i32(v66&int32(1) == v54) == v54 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_init_sexpr_4), int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														if l4 != 0 {
															v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
															v86 = F_exprLocation(m, l2)
															mBase = m.M
															F_executor_errposition(m, v85, v86)
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														} else {
															F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
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
											v94 = int32(0)
											if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
												v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v110 = int32(_a_F_init_sexpr_5)
													v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
													*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
													v114 = int32(1)
													if base.Ui32(v108-v114) <= base.Ui32(v114) {
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
														v119 = F_CreateTupleDescCopy(m, v118)
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return
														} else {
															v121 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
															*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
															*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
															v152 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
															*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
															m.G0 = v14 + int32(16)
															return
														}
													} else {
														switch v108 {
														case 0:
															v125 = F_CreateTemplateTupleDesc(m, int32(1))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
																v129 = int32(0)
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
																F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return
																} else {
																	v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
																	v136 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v136)
																	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v135
																	*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
																	v152 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
																	*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
																	m.G0 = v14 + int32(16)
																	return
																}
															}
														default:
															*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
															*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
															v152 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
															*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
															m.G0 = v14 + int32(16)
															return
														case 3:
															v139 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v139)
															*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
															*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
															v152 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
															*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
															m.G0 = v14 + int32(16)
															return
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
												v152 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
												*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
												m.G0 = v14 + int32(16)
												return
											}
										}
									}
								}
							}
						}
					} else {
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
						if v35 != 0 {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
							if int32(101) <= v36 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v162 = m.ExcPending
								if v162 != 0 {
									return
								} else {
									F_errcode(m, int32(50856197))
									mBase = m.M
									v165 = m.ExcPending
									if v165 != 0 {
										return
									} else {
										v166 = int32(100)
										*(*int32)(unsafe.Add(mBase, uint32(v14))) = v166
										F_errmsg_plural(m, int32(_a_F_init_sexpr_0), int32(_a_F_init_sexpr_1), v166, v14)
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_init_sexpr_2), int32(721), int32(_a_F_init_sexpr_3))
											mBase = m.M
											v177 = m.ExcPending
											if v177 != 0 {
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
								v40 = l3 + int32(16)
								F_fmgr_info_cxt(m, l0, v40, l5)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
									v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
										*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
										v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										v54 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
										v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
										v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
										v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
										v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
										*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
										if l6|base.B2i32(v66&int32(1) == v54) == v54 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return
												} else {
													F_errmsg(m, int32(_a_F_init_sexpr_4), int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														if l4 != 0 {
															v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
															v86 = F_exprLocation(m, l2)
															mBase = m.M
															F_executor_errposition(m, v85, v86)
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														} else {
															F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
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
											v94 = int32(0)
											if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
												v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
												v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
												mBase = m.M
												v109 = m.ExcPending
												if v109 != 0 {
													return
												} else {
													v110 = int32(_a_F_init_sexpr_5)
													v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
													*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
													v114 = int32(1)
													if base.Ui32(v108-v114) <= base.Ui32(v114) {
														v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
														v119 = F_CreateTupleDescCopy(m, v118)
														mBase = m.M
														v120 = m.ExcPending
														if v120 != 0 {
															return
														} else {
															v121 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
															*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
															*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
															v152 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
															*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
															m.G0 = v14 + int32(16)
															return
														}
													} else {
														switch v108 {
														case 0:
															v125 = F_CreateTemplateTupleDesc(m, int32(1))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
																v129 = int32(0)
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
																F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return
																} else {
																	v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
																	v136 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v136)
																	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v135
																	*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
																	v152 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
																	*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
																	m.G0 = v14 + int32(16)
																	return
																}
															}
														default:
															*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
															*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
															v152 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
															*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
															m.G0 = v14 + int32(16)
															return
														case 3:
															v139 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v139)
															*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
															*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
															v152 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
															*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
															m.G0 = v14 + int32(16)
															return
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
												v152 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
												*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
												m.G0 = v14 + int32(16)
												return
											}
										}
									}
								}
							}
						} else {
							v40 = l3 + int32(16)
							F_fmgr_info_cxt(m, l0, v40, l5)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
								v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
									*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									v54 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
									if l6|base.B2i32(v66&int32(1) == v54) == v54 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_init_sexpr_4), int32(0))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													if l4 != 0 {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
														v86 = F_exprLocation(m, l2)
														mBase = m.M
														F_executor_errposition(m, v85, v86)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													} else {
														F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
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
										v94 = int32(0)
										if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
											v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												v110 = int32(_a_F_init_sexpr_5)
												v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
												*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
												v114 = int32(1)
												if base.Ui32(v108-v114) <= base.Ui32(v114) {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
													v119 = F_CreateTupleDescCopy(m, v118)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														v121 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
														*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
														*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
														v152 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
														*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
														m.G0 = v14 + int32(16)
														return
													}
												} else {
													switch v108 {
													case 0:
														v125 = F_CreateTemplateTupleDesc(m, int32(1))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
															v129 = int32(0)
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
															F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return
															} else {
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
																v136 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v136)
																*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v135
																*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
																v152 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
																*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
																m.G0 = v14 + int32(16)
																return
															}
														}
													default:
														*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
														v152 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
														*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
														m.G0 = v14 + int32(16)
														return
													case 3:
														v139 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v139)
														*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
														v152 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
														*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
														m.G0 = v14 + int32(16)
														return
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
											v152 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
											*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
											m.G0 = v14 + int32(16)
											return
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[1]))
			if v32 != 0 {
				F_RunFunctionExecuteHook(m, l0)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					if v35 != 0 {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
						if int32(101) <= v36 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v162 = m.ExcPending
							if v162 != 0 {
								return
							} else {
								F_errcode(m, int32(50856197))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return
								} else {
									v166 = int32(100)
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v166
									F_errmsg_plural(m, int32(_a_F_init_sexpr_0), int32(_a_F_init_sexpr_1), v166, v14)
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_init_sexpr_2), int32(721), int32(_a_F_init_sexpr_3))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
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
							v40 = l3 + int32(16)
							F_fmgr_info_cxt(m, l0, v40, l5)
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return
							} else {
								v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
								v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
									*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
									v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									v54 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
									v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
									v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
									v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
									v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
									*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
									v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
									if l6|base.B2i32(v66&int32(1) == v54) == v54 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											F_errcode(m, int32(1088))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												F_errmsg(m, int32(_a_F_init_sexpr_4), int32(0))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													if l4 != 0 {
														v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
														v86 = F_exprLocation(m, l2)
														mBase = m.M
														F_executor_errposition(m, v85, v86)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return
														} else {
															F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													} else {
														F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
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
										v94 = int32(0)
										if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
											v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
											v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												v110 = int32(_a_F_init_sexpr_5)
												v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
												*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
												v114 = int32(1)
												if base.Ui32(v108-v114) <= base.Ui32(v114) {
													v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
													v119 = F_CreateTupleDescCopy(m, v118)
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														v121 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
														*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
														*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
														v152 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
														*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
														m.G0 = v14 + int32(16)
														return
													}
												} else {
													switch v108 {
													case 0:
														v125 = F_CreateTemplateTupleDesc(m, int32(1))
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
															v129 = int32(0)
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
															F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
															mBase = m.M
															v134 = m.ExcPending
															if v134 != 0 {
																return
															} else {
																v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
																v136 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v136)
																*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v135
																*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
																v152 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
																*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
																m.G0 = v14 + int32(16)
																return
															}
														}
													default:
														*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
														v152 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
														*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
														m.G0 = v14 + int32(16)
														return
													case 3:
														v139 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v139)
														*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
														*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
														v152 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
														*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
														m.G0 = v14 + int32(16)
														return
													}
												}
											}
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
											v152 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
											*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
											m.G0 = v14 + int32(16)
											return
										}
									}
								}
							}
						}
					} else {
						v40 = l3 + int32(16)
						F_fmgr_info_cxt(m, l0, v40, l5)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
							v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								v54 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
								if l6|base.B2i32(v66&int32(1) == v54) == v54 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_init_sexpr_4), int32(0))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												if l4 != 0 {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
													v86 = F_exprLocation(m, l2)
													mBase = m.M
													F_executor_errposition(m, v85, v86)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												} else {
													F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
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
									v94 = int32(0)
									if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
										v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = int32(_a_F_init_sexpr_5)
											v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
											*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
											v114 = int32(1)
											if base.Ui32(v108-v114) <= base.Ui32(v114) {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
												v119 = F_CreateTupleDescCopy(m, v118)
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
													*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
													*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
													v152 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
													*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
													m.G0 = v14 + int32(16)
													return
												}
											} else {
												switch v108 {
												case 0:
													v125 = F_CreateTemplateTupleDesc(m, int32(1))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
														v129 = int32(0)
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
														F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
															v136 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v136)
															*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v135
															*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
															v152 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
															*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
															m.G0 = v14 + int32(16)
															return
														}
													}
												default:
													*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
													v152 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
													*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
													m.G0 = v14 + int32(16)
													return
												case 3:
													v139 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v139)
													*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
													v152 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
													*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
													m.G0 = v14 + int32(16)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
										v152 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
										*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
										m.G0 = v14 + int32(16)
										return
									}
								}
							}
						}
					}
				}
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				if v35 != 0 {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
					if int32(101) <= v36 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v162 = m.ExcPending
						if v162 != 0 {
							return
						} else {
							F_errcode(m, int32(50856197))
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return
							} else {
								v166 = int32(100)
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v166
								F_errmsg_plural(m, int32(_a_F_init_sexpr_0), int32(_a_F_init_sexpr_1), v166, v14)
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_init_sexpr_2), int32(721), int32(_a_F_init_sexpr_3))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
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
						v40 = l3 + int32(16)
						F_fmgr_info_cxt(m, l0, v40, l5)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
							v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
								*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
								v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								v54 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
								v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
								v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
								*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
								v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
								if l6|base.B2i32(v66&int32(1) == v54) == v54 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_init_sexpr_4), int32(0))
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return
											} else {
												if l4 != 0 {
													v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
													v86 = F_exprLocation(m, l2)
													mBase = m.M
													F_executor_errposition(m, v85, v86)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return
													} else {
														F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												} else {
													F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
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
									v94 = int32(0)
									if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
										v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = int32(_a_F_init_sexpr_5)
											v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
											*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
											v114 = int32(1)
											if base.Ui32(v108-v114) <= base.Ui32(v114) {
												v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
												v119 = F_CreateTupleDescCopy(m, v118)
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
													*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
													*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
													v152 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
													*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
													m.G0 = v14 + int32(16)
													return
												}
											} else {
												switch v108 {
												case 0:
													v125 = F_CreateTemplateTupleDesc(m, int32(1))
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
														v129 = int32(0)
														v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
														F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
														mBase = m.M
														v134 = m.ExcPending
														if v134 != 0 {
															return
														} else {
															v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
															v136 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v136)
															*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v135
															*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
															v152 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
															*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
															m.G0 = v14 + int32(16)
															return
														}
													}
												default:
													*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
													v152 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
													*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
													m.G0 = v14 + int32(16)
													return
												case 3:
													v139 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v139)
													*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
													*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
													v152 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
													*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
													m.G0 = v14 + int32(16)
													return
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
										v152 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
										*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
										m.G0 = v14 + int32(16)
										return
									}
								}
							}
						}
					}
				} else {
					v40 = l3 + int32(16)
					F_fmgr_info_cxt(m, l0, v40, l5)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
						v49 = F_palloc(m, v19<<(uint(int32(3))%32)+int32(20))
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3)+60)) = v49
							*(*int32)(unsafe.Add(mBase, uint32(v49))) = v40
							v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
							v54 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
							*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v54
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
							*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = l1
							v61 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
							*(*uint8)(unsafe.Add(mBase, uint32(v61)+16)) = uint8(v54)
							v64 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
							*(*uint16)(unsafe.Add(mBase, uint32(v64)+18)) = uint16(v19)
							v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+27)))
							if l6|base.B2i32(v66&int32(1) == v54) == v54 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										F_errmsg(m, int32(_a_F_init_sexpr_4), int32(0))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return
										} else {
											if l4 != 0 {
												v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
												v86 = F_exprLocation(m, l2)
												mBase = m.M
												F_executor_errposition(m, v85, v86)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											} else {
												F_errfinish(m, int32(_a_F_init_sexpr_2), int32(740), int32(_a_F_init_sexpr_3))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
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
								v94 = int32(0)
								if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
									v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v110 = int32(_a_F_init_sexpr_5)
										v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
										*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
										v114 = int32(1)
										if base.Ui32(v108-v114) <= base.Ui32(v114) {
											v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
											v119 = F_CreateTupleDescCopy(m, v118)
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												v121 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
												*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
												*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
												v152 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
												*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
												m.G0 = v14 + int32(16)
												return
											}
										} else {
											switch v108 {
											case 0:
												v125 = F_CreateTemplateTupleDesc(m, int32(1))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
													v129 = int32(0)
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
													F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return
													} else {
														v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
														v136 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v136)
														*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v135
														*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
														v152 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
														*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
														m.G0 = v14 + int32(16)
														return
													}
												}
											default:
												*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
												*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
												v152 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
												*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
												m.G0 = v14 + int32(16)
												return
											case 3:
												v139 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v139)
												*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
												*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
												v152 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
												*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
												m.G0 = v14 + int32(16)
												return
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
									v152 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v152)
									*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
									m.G0 = v14 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_instrumentSortedGroup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int64
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v90 int64
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v9 + int64(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v17 == v3 {
		v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)))
		v34 = v23
		v35 = v20 - v21
		if v34&int32(255) != base.B2i32(v17 != int32(0)) {
			if v34&int32(1) != 0 {
				v55 = v3
			} else {
				v55 = int32(1)
			}
		} else {
			v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
			if v35 <= v41 {
				if v34&int32(1) != 0 {
					v55 = v3
				} else {
					v55 = int32(1)
				}
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)) = uint8(v34)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v35
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v45
				if v34&int32(1) == int32(0) {
					v55 = int32(1)
				} else {
					v55 = v3
				}
			}
		}
	} else {
		v24 = F_LogicalTapeSetBlocks(m, v17)
		mBase = m.M
		v26 = v24 << (uint(int64(13)) % 64)
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)))
		if v28 != 0 {
			v34 = int32(1)
			v35 = v26
			if v34&int32(255) != base.B2i32(v17 != int32(0)) {
				if v34&int32(1) != 0 {
					v55 = v3
				} else {
					v55 = int32(1)
				}
			} else {
				v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
				if v35 <= v41 {
					if v34&int32(1) != 0 {
						v55 = v3
					} else {
						v55 = int32(1)
					}
				} else {
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)) = uint8(v34)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v35
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v45
					if v34&int32(1) == int32(0) {
						v55 = int32(1)
					} else {
						v55 = v3
					}
				}
			}
		} else {
			v29 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+120)) = uint8(v29)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+112)) = v26
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v32
			v55 = v3
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v55
	v58 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	v62 = base.I64_div_s(v58+int64(1023), int64(1024))
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	switch v64 - int32(3) {
	case 0:
		v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+69)))
		if v69 != 0 {
			v70 = int32(1)
		} else {
			v70 = int32(2)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v70
	case 1:
		v75 = v64
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v75
	case 2:
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(8)
	default:
		v75 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = v75
	}
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	switch v78 {
	case 0:
		v79 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v79 + v80
		v83 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		if v79 <= v83 {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v79
		}
	case 1:
		v86 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v86 + v87
		v90 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		if v86 <= v90 {
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v86
		}
	default:
	}
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v94 | v95
	m.G0 = v7 + int32(16)
	return
}
func F_int24ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v2 <= v3)
}
func F_int24lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v3 < v2)
}
func F_int28lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v4 < v3)
}
func F_int28mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = v8 - v5
	if base.B2i32(int64(0) < v5) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int28mi_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int28mi_1), int32(1136), int32(_a_F_int28mi_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int2abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2&int32(_a_F_int2abs_0) == int32(_a_F_int2abs_1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int2abs_2), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int2abs_3), int32(1242), int32(_a_F_int2abs_4))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v25 = base.I32_extend16_s(v2)
		v27 = v25 >> (uint(int32(31)) % 32)
		return v25 ^ v27 - v27
	}
}
func F_int2eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 == v3)
}
func F_int2hashfast(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	v2 = base.I32_extend16_s(l0)
	v3 = int32(16)
	v7 = (int32(base.Ui32(v2)>>(uint(v3)%32)) ^ v2) * int32(-2048144789)
	v12 = (int32(base.Ui32(v7)>>(uint(int32(13))%32)) ^ v7) * int32(-1028477387)
	return int32(base.Ui32(v12)>>(uint(v3)%32)) ^ v12
}
func F_int2lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 < v3)
}
func F_int2recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pq_getmsgint(m, v2, int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return base.I32_extend16_s(v4)
	}
}
func F_int2smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	if v3 < v4 {
		v6 = v3
	} else {
		v6 = v4
	}
	return v6
}
func F_int2vectorin(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_palloc0(m, int32(88))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = v14
	v23 = v16
	v24 = int32(32)
	v26 = int32(0)
	goto L3
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if base.B2i32(base.Ui32(v29-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v29 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v148
L5:
	;
	v21 = v21 + int32(1)
	goto L3
L6:
	;
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	m.G0 = v11 + int32(48)
	goto L7
L9:
	;
	if v26 < v24 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = int32(21)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v26<<(uint(int32(3))%32) + int32(96)
	v148 = v23
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_int2vectorin[0])) = int32(0)
	v57 = F_strtox_2(m, v21, v11+int32(44), int32(10), int64(2147483648))
	mBase = m.M
	v58 = base.I32_wrap_i64(v57)
	goto L17
L13:
	;
	v48 = v23
	v49 = v24
	goto L12
L14:
	;
	goto L15
L15:
	;
	v46 = F_repalloc(m, v23, v24<<(uint(int32(2))%32)+int32(24))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v48 = v46
	v49 = v24 << (uint(int32(1)) % 32)
	goto L12
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v60 == v21 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v129 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v48+v26<<(uint(v129)%32))+24)) = uint16(v58)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	v21 = v135
	v23 = v48
	v24 = v49
	v26 = v26 + v129
	goto L3
L19:
	;
	v148 = int32(0)
	goto L8
L20:
	;
	F_errsave_finish(m, v13, int32(_a_F_int2vectorin_0), v124, int32(_a_F_int2vectorin_1))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v62 = F_errsave_start(m, v13)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_int2vectorin[0]))
	if base.B2i32(v77 != int32(68))&base.B2i32(base.Ui32(int32(-65537)) < base.Ui32(v58-int32(_a_F_int2vectorin_2))) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v62 == int32(0) {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_int2vectorin_3)
	F_errmsg(m, int32(_a_F_int2vectorin_4), v11)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v124 = int32(199)
	goto L20
L28:
	;
	v87 = F_errsave_start(m, v13)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	v104 = int32(32)
	if v103|v104 == v104 {
		goto L18
	} else {
		goto L35
	}
L31:
	;
	if v87 == int32(0) {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(_a_F_int2vectorin_3)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v21
	F_errmsg(m, int32(_a_F_int2vectorin_5), v11+int32(16))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v124 = int32(205)
	goto L20
L35:
	;
	v108 = F_errsave_start(m, v13)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v108 == int32(0) {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(_a_F_int2vectorin_3)
	F_errmsg(m, int32(_a_F_int2vectorin_4), v11+int32(32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v124 = int32(211)
	goto L20
L40:
	;
	goto L19
}
func F_int42ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 <= v2)
}
func F_int42lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v2 < v3)
}
func F_int42mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = v6 - v3
	if base.B2i32(int32(0) < v3) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int42mi_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int42mi_1), int32(1101), int32(_a_F_int42mi_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return v7
	}
}
func F_int42pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = v6 + v3
	if base.B2i32(v3 < int32(0)) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int42pl_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int42pl_1), int32(1087), int32(_a_F_int42pl_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return v7
	}
}
func F_int48lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	return base.B2i32(v4 < v3)
}
func F_int48mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = v8 - v5
	if base.B2i32(int64(0) < v5) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int48mi_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int48mi_1), int32(994), int32(_a_F_int48mi_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int48pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+20)))
	v9 = v5 + v8
	if base.B2i32(v5 < int64(0)) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int48pl_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int48pl_1), int32(980), int32(_a_F_int48pl_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int4abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v3 == int32(-2147483648) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int4abs_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4abs_1), int32(1228), int32(_a_F_int4abs_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = v3 >> (uint(int32(31)) % 32)
		return v3 ^ v25 - v25
	}
}
func F_int4eqfast(m *base.Module, l0 int32, l1 int32) int32 {
	return base.B2i32(l0 == l1)
}
func F_int4hashfast(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	v2 = int32(16)
	v6 = (int32(base.Ui32(l0)>>(uint(v2)%32)) ^ l0) * int32(-2048144789)
	v11 = (int32(base.Ui32(v6)>>(uint(int32(13))%32)) ^ v6) * int32(-1028477387)
	return int32(base.Ui32(v11)>>(uint(v2)%32)) ^ v11
}
func F_int4lcm(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var __phi32 int32
	_ = __phi32
	var v33 int32
	_ = v33
	var __phi33 int32
	_ = __phi33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == v2 {
		v66 = v2
		return v66
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v9 == int32(0) {
			v66 = v2
			return v66
		} else {
			v12 = int32(31)
			v13 = v6 >> (uint(v12) % 32)
			v17 = v9 >> (uint(v12) % 32)
			v20 = base.B2i32(base.Ui32(v17-(v17^v9)) < base.Ui32(v13-(v13^v6)))
			if base.Ui32(v17-(v17^v9)) < base.Ui32(v13-(v13^v6)) {
				v21 = v6
			} else {
				v21 = v9
			}
			if base.Ui32(v17-(v17^v9)) < base.Ui32(v13-(v13^v6)) {
				v22 = v9
			} else {
				v22 = v6
			}
			if v22 != int32(-2147483648) {
				__phi32 = v21
				__phi33 = v22
				v32 = __phi32
				v33 = __phi33
				for {
					v37 = base.I32_rem_s(v33, v32)
					if v37 != 0 {
						__phi32 = v37
						__phi33 = v32
						v32 = __phi32
						v33 = __phi33
						continue
					} else {
						break
					}
					break
				}
				v39 = v32 >> (uint(int32(31)) % 32)
				v47 = v32 ^ v39 - v39
				v48 = base.I32_div_s(v6, v47)
				v51 = base.I64_extend_i32_s(v48) * base.I64_extend_i32_s(v9)
				v55 = base.I32_wrap_i64(v51)
				if base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(32))%64))) != v55>>(uint(int32(31))%32) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_int4lcm_1), int32(1360), int32(_a_F_int4lcm_2))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
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
					if v55 == int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_int4lcm_1), int32(1366), int32(_a_F_int4lcm_2))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
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
						v62 = v55 >> (uint(int32(31)) % 32)
						v66 = v55 ^ v62 - v62
						return v66
					}
				}
			} else {
				if v21&int32(2147483647) == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_int4lcm_1), int32(1292), int32(_a_F_int4lcm_3))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
					if v21 != int32(-1) {
						__phi32 = v21
						__phi33 = v22
						v32 = __phi32
						v33 = __phi33
						for {
							v37 = base.I32_rem_s(v33, v32)
							if v37 != 0 {
								__phi32 = v37
								__phi33 = v32
								v32 = __phi32
								v33 = __phi33
								continue
							} else {
								break
							}
							break
						}
						v39 = v32 >> (uint(int32(31)) % 32)
						v47 = v32 ^ v39 - v39
					} else {
						v47 = int32(1)
					}
					v48 = base.I32_div_s(v6, v47)
					v51 = base.I64_extend_i32_s(v48) * base.I64_extend_i32_s(v9)
					v55 = base.I32_wrap_i64(v51)
					if base.I32_wrap_i64(int64(base.Ui64(v51)>>(uint(int64(32))%64))) != v55>>(uint(int32(31))%32) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_int4lcm_1), int32(1360), int32(_a_F_int4lcm_2))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
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
						if v55 == int32(-2147483648) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50331778))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_int4lcm_1), int32(1366), int32(_a_F_int4lcm_2))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
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
							v62 = v55 >> (uint(int32(31)) % 32)
							v66 = v55 ^ v62 - v62
							return v66
						}
					}
				}
			}
		}
	}
}
func F_int4mod(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	switch v4 + int32(1) {
	case 0:
		v27 = int32(0)
		return v27
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int4mod_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int4mod_1), int32(1168), int32(_a_F_int4mod_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v26 = base.I32_rem_s(v25, v4)
		v27 = v26
		return v27
	}
}
func F_int4range_canonical(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
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
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v17 {
				v30 = v19
				F_range_deserialize(m, v30, v12, v9+int32(24), v9+int32(16), v9+int32(15))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v39 == int32(1) {
						v112 = v12
						m.G0 = v9 + int32(32)
						return v112
					} else {
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
						if v42 != 0 {
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
							if v72 != 0 {
								v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									v112 = v109
									m.G0 = v9 + int32(32)
									return v112
								}
							} else {
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
								if v73&int32(1) == int32(0) {
									v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v112 = v109
										m.G0 = v9 + int32(32)
										return v112
									}
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
									if v78 == int32(2147483647) {
										v81 = int32(0)
										v82 = F_errsave_start(m, v16)
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int32(0)
										} else {
											if v82 == int32(0) {
												v112 = v81
												m.G0 = v9 + int32(32)
												return v112
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int32(0)
														} else {
															v112 = v81
															m.G0 = v9 + int32(32)
															return v112
														}
													}
												}
											}
										}
									} else {
										v98 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
										v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v112 = v109
											m.G0 = v9 + int32(32)
											return v112
										}
									}
								}
							}
						} else {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
							if v43&int32(1) != 0 {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								if v72 != 0 {
									v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v112 = v109
										m.G0 = v9 + int32(32)
										return v112
									}
								} else {
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
									if v73&int32(1) == int32(0) {
										v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v112 = v109
											m.G0 = v9 + int32(32)
											return v112
										}
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
										if v78 == int32(2147483647) {
											v81 = int32(0)
											v82 = F_errsave_start(m, v16)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												if v82 == int32(0) {
													v112 = v81
													m.G0 = v9 + int32(32)
													return v112
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																v112 = v81
																m.G0 = v9 + int32(32)
																return v112
															}
														}
													}
												}
											}
										} else {
											v98 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
											v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												v112 = v109
												m.G0 = v9 + int32(32)
												return v112
											}
										}
									}
								}
							} else {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
								if v46 == int32(2147483647) {
									v49 = int32(0)
									v50 = F_errsave_start(m, v16)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										if v50 == int32(0) {
											v112 = v49
											m.G0 = v9 + int32(32)
											return v112
										} else {
											F_errcode(m, int32(50331778))
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1552), int32(_a_F_int4range_canonical_2))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														v112 = v49
														m.G0 = v9 + int32(32)
														return v112
													}
												}
											}
										}
									}
								} else {
									v66 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)) = uint8(v66)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v46 + v66
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v72 != 0 {
										v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v112 = v109
											m.G0 = v9 + int32(32)
											return v112
										}
									} else {
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
										if v73&int32(1) == int32(0) {
											v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												v112 = v109
												m.G0 = v9 + int32(32)
												return v112
											}
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											if v78 == int32(2147483647) {
												v81 = int32(0)
												v82 = F_errsave_start(m, v16)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													if v82 == int32(0) {
														v112 = v81
														m.G0 = v9 + int32(32)
														return v112
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	v112 = v81
																	m.G0 = v9 + int32(32)
																	return v112
																}
															}
														}
													}
												}
											} else {
												v98 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
												v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													v112 = v109
													m.G0 = v9 + int32(32)
													return v112
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
				v23 = F_lookup_type_cache(m, v17, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(_a_F_int4range_canonical_3), v9)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_int4range_canonical_1), int32(1776), int32(_a_F_int4range_canonical_4))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						F_range_deserialize(m, v30, v12, v9+int32(24), v9+int32(16), v9+int32(15))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
							if v39 == int32(1) {
								v112 = v12
								m.G0 = v9 + int32(32)
								return v112
							} else {
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
								if v42 != 0 {
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v72 != 0 {
										v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v112 = v109
											m.G0 = v9 + int32(32)
											return v112
										}
									} else {
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
										if v73&int32(1) == int32(0) {
											v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												v112 = v109
												m.G0 = v9 + int32(32)
												return v112
											}
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											if v78 == int32(2147483647) {
												v81 = int32(0)
												v82 = F_errsave_start(m, v16)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													if v82 == int32(0) {
														v112 = v81
														m.G0 = v9 + int32(32)
														return v112
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	v112 = v81
																	m.G0 = v9 + int32(32)
																	return v112
																}
															}
														}
													}
												}
											} else {
												v98 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
												v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													v112 = v109
													m.G0 = v9 + int32(32)
													return v112
												}
											}
										}
									}
								} else {
									v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
									if v43&int32(1) != 0 {
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										if v72 != 0 {
											v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												v112 = v109
												m.G0 = v9 + int32(32)
												return v112
											}
										} else {
											v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
											if v73&int32(1) == int32(0) {
												v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													v112 = v109
													m.G0 = v9 + int32(32)
													return v112
												}
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
												if v78 == int32(2147483647) {
													v81 = int32(0)
													v82 = F_errsave_start(m, v16)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														if v82 == int32(0) {
															v112 = v81
															m.G0 = v9 + int32(32)
															return v112
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return int32(0)
																} else {
																	F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int32(0)
																	} else {
																		v112 = v81
																		m.G0 = v9 + int32(32)
																		return v112
																	}
																}
															}
														}
													}
												} else {
													v98 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
													v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
													} else {
														v112 = v109
														m.G0 = v9 + int32(32)
														return v112
													}
												}
											}
										}
									} else {
										v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										if v46 == int32(2147483647) {
											v49 = int32(0)
											v50 = F_errsave_start(m, v16)
											mBase = m.M
											v51 = m.ExcPending
											if v51 != 0 {
												return int32(0)
											} else {
												if v50 == int32(0) {
													v112 = v49
													m.G0 = v9 + int32(32)
													return v112
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v56 = m.ExcPending
													if v56 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1552), int32(_a_F_int4range_canonical_2))
															mBase = m.M
															v65 = m.ExcPending
															if v65 != 0 {
																return int32(0)
															} else {
																v112 = v49
																m.G0 = v9 + int32(32)
																return v112
															}
														}
													}
												}
											}
										} else {
											v66 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)) = uint8(v66)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v46 + v66
											v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
											if v72 != 0 {
												v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													v112 = v109
													m.G0 = v9 + int32(32)
													return v112
												}
											} else {
												v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
												if v73&int32(1) == int32(0) {
													v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
													} else {
														v112 = v109
														m.G0 = v9 + int32(32)
														return v112
													}
												} else {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
													if v78 == int32(2147483647) {
														v81 = int32(0)
														v82 = F_errsave_start(m, v16)
														mBase = m.M
														v83 = m.ExcPending
														if v83 != 0 {
															return int32(0)
														} else {
															if v82 == int32(0) {
																v112 = v81
																m.G0 = v9 + int32(32)
																return v112
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v88 = m.ExcPending
																if v88 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return int32(0)
																	} else {
																		F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
																		mBase = m.M
																		v97 = m.ExcPending
																		if v97 != 0 {
																			return int32(0)
																		} else {
																			v112 = v81
																			m.G0 = v9 + int32(32)
																			return v112
																		}
																	}
																}
															}
														}
													} else {
														v98 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
														v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return int32(0)
														} else {
															v112 = v109
															m.G0 = v9 + int32(32)
															return v112
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
			v23 = F_lookup_type_cache(m, v17, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(_a_F_int4range_canonical_3), v9)
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_int4range_canonical_1), int32(1776), int32(_a_F_int4range_canonical_4))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					F_range_deserialize(m, v30, v12, v9+int32(24), v9+int32(16), v9+int32(15))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v39 == int32(1) {
							v112 = v12
							m.G0 = v9 + int32(32)
							return v112
						} else {
							v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+28)))
							if v42 != 0 {
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
								if v72 != 0 {
									v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v112 = v109
										m.G0 = v9 + int32(32)
										return v112
									}
								} else {
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
									if v73&int32(1) == int32(0) {
										v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v112 = v109
											m.G0 = v9 + int32(32)
											return v112
										}
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
										if v78 == int32(2147483647) {
											v81 = int32(0)
											v82 = F_errsave_start(m, v16)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												if v82 == int32(0) {
													v112 = v81
													m.G0 = v9 + int32(32)
													return v112
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int32(0)
															} else {
																v112 = v81
																m.G0 = v9 + int32(32)
																return v112
															}
														}
													}
												}
											}
										} else {
											v98 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
											v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												v112 = v109
												m.G0 = v9 + int32(32)
												return v112
											}
										}
									}
								}
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)))
								if v43&int32(1) != 0 {
									v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
									if v72 != 0 {
										v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v112 = v109
											m.G0 = v9 + int32(32)
											return v112
										}
									} else {
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
										if v73&int32(1) == int32(0) {
											v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												v112 = v109
												m.G0 = v9 + int32(32)
												return v112
											}
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
											if v78 == int32(2147483647) {
												v81 = int32(0)
												v82 = F_errsave_start(m, v16)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													if v82 == int32(0) {
														v112 = v81
														m.G0 = v9 + int32(32)
														return v112
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int32(0)
																} else {
																	v112 = v81
																	m.G0 = v9 + int32(32)
																	return v112
																}
															}
														}
													}
												}
											} else {
												v98 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
												v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													v112 = v109
													m.G0 = v9 + int32(32)
													return v112
												}
											}
										}
									}
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
									if v46 == int32(2147483647) {
										v49 = int32(0)
										v50 = F_errsave_start(m, v16)
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											if v50 == int32(0) {
												v112 = v49
												m.G0 = v9 + int32(32)
												return v112
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
													mBase = m.M
													v60 = m.ExcPending
													if v60 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1552), int32(_a_F_int4range_canonical_2))
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return int32(0)
														} else {
															v112 = v49
															m.G0 = v9 + int32(32)
															return v112
														}
													}
												}
											}
										}
									} else {
										v66 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+29)) = uint8(v66)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v46 + v66
										v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+20)))
										if v72 != 0 {
											v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return int32(0)
											} else {
												v112 = v109
												m.G0 = v9 + int32(32)
												return v112
											}
										} else {
											v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)))
											if v73&int32(1) == int32(0) {
												v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													v112 = v109
													m.G0 = v9 + int32(32)
													return v112
												}
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
												if v78 == int32(2147483647) {
													v81 = int32(0)
													v82 = F_errsave_start(m, v16)
													mBase = m.M
													v83 = m.ExcPending
													if v83 != 0 {
														return int32(0)
													} else {
														if v82 == int32(0) {
															v112 = v81
															m.G0 = v9 + int32(32)
															return v112
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return int32(0)
																} else {
																	F_errsave_finish(m, v16, int32(_a_F_int4range_canonical_1), int32(1565), int32(_a_F_int4range_canonical_2))
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int32(0)
																	} else {
																		v112 = v81
																		m.G0 = v9 + int32(32)
																		return v112
																	}
																}
															}
														}
													}
												} else {
													v98 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v9)+21)) = uint8(v98)
													*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v78 + int32(1)
													v109 = F_range_serialize(m, v30, v9+int32(24), v9+int32(16), int32(0), v16)
													mBase = m.M
													v110 = m.ExcPending
													if v110 != 0 {
														return int32(0)
													} else {
														v112 = v109
														m.G0 = v9 + int32(32)
														return v112
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
	}
}
func F_int64_to_numeric(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v32 int64
	_ = v32
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	v16 = F_palloc(m, int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v16
		v21 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v16))) = uint16(v21)
		*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v16 + int32(2)
		if l0 < int64(0) {
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(16384)
			v36 = int64(0) - l0
			v39 = v36
			v42 = v2
			v43 = v16 + int32(12)
			for {
				v48 = v43 - int32(2)
				v50 = base.I64_div_u_s(v39, int64(10000))
				v53 = v50*int64(55536) + v39
				*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v53)
				v56 = v42 + int32(1)
				if base.Ui64(int64(9999)) < base.Ui64(v39) {
					v39 = v50
					v42 = v56
					v43 = v48
					continue
				} else {
					break
				}
				break
			}
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v48
			v63 = v56
			v65 = v42
		} else {
			v32 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v32
			if l0 == v32 {
				v63 = v2
				v65 = v2
			} else {
				v36 = l0
				v39 = v36
				v42 = v2
				v43 = v16 + int32(12)
				for {
					v48 = v43 - int32(2)
					v50 = base.I64_div_u_s(v39, int64(10000))
					v53 = v50*int64(55536) + v39
					*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v53)
					v56 = v42 + int32(1)
					if base.Ui64(int64(9999)) < base.Ui64(v39) {
						v39 = v50
						v42 = v56
						v43 = v48
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v48
				v63 = v56
				v65 = v42
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v65
		*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v63
		v73 = F_make_result_opt_error(m, v11+int32(8), int32(0))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			F_pfree(m, v16)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(32)
				return v73
			}
		}
	}
}
func F_int82eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 == v4)
}
func F_int82pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = v4 + v8
	if base.B2i32(v4 < int64(0)) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int82pl_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int82pl_1), int32(1041), int32(_a_F_int82pl_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int84(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(v4-int64(2147483648)) <= base.Ui64(int64(-4294967297)) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int84_0), int32(0))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int84_1), int32(1256), int32(_a_F_int84_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
		return base.I32_wrap_i64(v4)
	}
}
func F_int84lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int64
	_ = v4
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v3 < v4)
}
func F_int84mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	v9 = v8 - v4
	if base.B2i32(int64(0) < v4) != base.B2i32(v9 < v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int84mi_0), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int84mi_1), int32(913), int32(_a_F_int84mi_2))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
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
		v30 = F_Int64GetDatum(m, v9)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int8abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	if v5 == int64(-9223372036854775807-1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int8abs_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int8abs_1), int32(554), int32(_a_F_int8abs_2))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
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
		v27 = v5 >> (uint(int64(63)) % 64)
		v30 = F_Int64GetDatum(m, v5^v27-v27)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v30
		}
	}
}
func F_int8range_canonical(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int64
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		if v20 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			if v21 == v18 {
				v31 = v20
				F_range_deserialize(m, v31, v13, v10+int32(24), v10+int32(16), v10+int32(15))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v40 == int32(1) {
						v121 = v13
						m.G0 = v10 + int32(32)
						return v121
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
						if v43 != 0 {
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
							if v77 != 0 {
								v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									v121 = v118
									m.G0 = v10 + int32(32)
									return v121
								}
							} else {
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
								if v78&int32(1) == int32(0) {
									v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										v121 = v118
										m.G0 = v10 + int32(32)
										return v121
									}
								} else {
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
									v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
									if v84 == int64(9223372036854775807) {
										v87 = int32(0)
										v88 = F_errsave_start(m, v17)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return int32(0)
										} else {
											if v88 == int32(0) {
												v121 = v87
												m.G0 = v10 + int32(32)
												return v121
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
													mBase = m.M
													v98 = m.ExcPending
													if v98 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v121 = v87
															m.G0 = v10 + int32(32)
															return v121
														}
													}
												}
											}
										}
									} else {
										v106 = F_Int64GetDatum(m, v84+int64(1))
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											v108 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
											v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v121 = v118
												m.G0 = v10 + int32(32)
												return v121
											}
										}
									}
								}
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
							if v44&int32(1) != 0 {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
								if v77 != 0 {
									v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										v121 = v118
										m.G0 = v10 + int32(32)
										return v121
									}
								} else {
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
									if v78&int32(1) == int32(0) {
										v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											v121 = v118
											m.G0 = v10 + int32(32)
											return v121
										}
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
										v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
										if v84 == int64(9223372036854775807) {
											v87 = int32(0)
											v88 = F_errsave_start(m, v17)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												if v88 == int32(0) {
													v121 = v87
													m.G0 = v10 + int32(32)
													return v121
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v121 = v87
																m.G0 = v10 + int32(32)
																return v121
															}
														}
													}
												}
											}
										} else {
											v106 = F_Int64GetDatum(m, v84+int64(1))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												v108 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
												v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													v121 = v118
													m.G0 = v10 + int32(32)
													return v121
												}
											}
										}
									}
								}
							} else {
								v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
								v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
								if v48 == int64(9223372036854775807) {
									v51 = int32(0)
									v52 = F_errsave_start(m, v17)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										if v52 == int32(0) {
											v121 = v51
											m.G0 = v10 + int32(32)
											return v121
										} else {
											F_errcode(m, int32(50331778))
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
												mBase = m.M
												v62 = m.ExcPending
												if v62 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1599), int32(_a_F_int8range_canonical_2))
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														v121 = v51
														m.G0 = v10 + int32(32)
														return v121
													}
												}
											}
										}
									}
								} else {
									v70 = F_Int64GetDatum(m, v48+int64(1))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
										return int32(0)
									} else {
										v72 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v72)
										*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v70
										v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
										if v77 != 0 {
											v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v121 = v118
												m.G0 = v10 + int32(32)
												return v121
											}
										} else {
											v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
											if v78&int32(1) == int32(0) {
												v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													v121 = v118
													m.G0 = v10 + int32(32)
													return v121
												}
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
												v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
												if v84 == int64(9223372036854775807) {
													v87 = int32(0)
													v88 = F_errsave_start(m, v17)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														if v88 == int32(0) {
															v121 = v87
															m.G0 = v10 + int32(32)
															return v121
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return int32(0)
																} else {
																	F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v121 = v87
																		m.G0 = v10 + int32(32)
																		return v121
																	}
																}
															}
														}
													}
												} else {
													v106 = F_Int64GetDatum(m, v84+int64(1))
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int32(0)
													} else {
														v108 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
														*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
														v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return int32(0)
														} else {
															v121 = v118
															m.G0 = v10 + int32(32)
															return v121
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
				v24 = F_lookup_type_cache(m, v18, int32(2048))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
							F_errmsg_internal(m, int32(_a_F_int8range_canonical_3), v10)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_int8range_canonical_1), int32(1776), int32(_a_F_int8range_canonical_4))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
						v31 = v24
						F_range_deserialize(m, v31, v13, v10+int32(24), v10+int32(16), v10+int32(15))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v40 == int32(1) {
								v121 = v13
								m.G0 = v10 + int32(32)
								return v121
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
								if v43 != 0 {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
									if v77 != 0 {
										v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											v121 = v118
											m.G0 = v10 + int32(32)
											return v121
										}
									} else {
										v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
										if v78&int32(1) == int32(0) {
											v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v121 = v118
												m.G0 = v10 + int32(32)
												return v121
											}
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
											v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
											if v84 == int64(9223372036854775807) {
												v87 = int32(0)
												v88 = F_errsave_start(m, v17)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													if v88 == int32(0) {
														v121 = v87
														m.G0 = v10 + int32(32)
														return v121
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	v121 = v87
																	m.G0 = v10 + int32(32)
																	return v121
																}
															}
														}
													}
												}
											} else {
												v106 = F_Int64GetDatum(m, v84+int64(1))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													v108 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
													v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														v121 = v118
														m.G0 = v10 + int32(32)
														return v121
													}
												}
											}
										}
									}
								} else {
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
									if v44&int32(1) != 0 {
										v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
										if v77 != 0 {
											v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v121 = v118
												m.G0 = v10 + int32(32)
												return v121
											}
										} else {
											v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
											if v78&int32(1) == int32(0) {
												v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													v121 = v118
													m.G0 = v10 + int32(32)
													return v121
												}
											} else {
												v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
												v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
												if v84 == int64(9223372036854775807) {
													v87 = int32(0)
													v88 = F_errsave_start(m, v17)
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														if v88 == int32(0) {
															v121 = v87
															m.G0 = v10 + int32(32)
															return v121
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v94 = m.ExcPending
															if v94 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return int32(0)
																} else {
																	F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
																	mBase = m.M
																	v103 = m.ExcPending
																	if v103 != 0 {
																		return int32(0)
																	} else {
																		v121 = v87
																		m.G0 = v10 + int32(32)
																		return v121
																	}
																}
															}
														}
													}
												} else {
													v106 = F_Int64GetDatum(m, v84+int64(1))
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int32(0)
													} else {
														v108 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
														*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
														v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return int32(0)
														} else {
															v121 = v118
															m.G0 = v10 + int32(32)
															return v121
														}
													}
												}
											}
										}
									} else {
										v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
										v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
										if v48 == int64(9223372036854775807) {
											v51 = int32(0)
											v52 = F_errsave_start(m, v17)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												if v52 == int32(0) {
													v121 = v51
													m.G0 = v10 + int32(32)
													return v121
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
														mBase = m.M
														v62 = m.ExcPending
														if v62 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1599), int32(_a_F_int8range_canonical_2))
															mBase = m.M
															v67 = m.ExcPending
															if v67 != 0 {
																return int32(0)
															} else {
																v121 = v51
																m.G0 = v10 + int32(32)
																return v121
															}
														}
													}
												}
											}
										} else {
											v70 = F_Int64GetDatum(m, v48+int64(1))
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return int32(0)
											} else {
												v72 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v72)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v70
												v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
												if v77 != 0 {
													v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														v121 = v118
														m.G0 = v10 + int32(32)
														return v121
													}
												} else {
													v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
													if v78&int32(1) == int32(0) {
														v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return int32(0)
														} else {
															v121 = v118
															m.G0 = v10 + int32(32)
															return v121
														}
													} else {
														v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
														v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
														if v84 == int64(9223372036854775807) {
															v87 = int32(0)
															v88 = F_errsave_start(m, v17)
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int32(0)
															} else {
																if v88 == int32(0) {
																	v121 = v87
																	m.G0 = v10 + int32(32)
																	return v121
																} else {
																	F_errcode(m, int32(50331778))
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
																		mBase = m.M
																		v98 = m.ExcPending
																		if v98 != 0 {
																			return int32(0)
																		} else {
																			F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
																			mBase = m.M
																			v103 = m.ExcPending
																			if v103 != 0 {
																				return int32(0)
																			} else {
																				v121 = v87
																				m.G0 = v10 + int32(32)
																				return v121
																			}
																		}
																	}
																}
															}
														} else {
															v106 = F_Int64GetDatum(m, v84+int64(1))
															mBase = m.M
															v107 = m.ExcPending
															if v107 != 0 {
																return int32(0)
															} else {
																v108 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
																*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
																v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
																	return int32(0)
																} else {
																	v121 = v118
																	m.G0 = v10 + int32(32)
																	return v121
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
				}
			}
		} else {
			v24 = F_lookup_type_cache(m, v18, int32(2048))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
						F_errmsg_internal(m, int32(_a_F_int8range_canonical_3), v10)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_int8range_canonical_1), int32(1776), int32(_a_F_int8range_canonical_4))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v24
					v31 = v24
					F_range_deserialize(m, v31, v13, v10+int32(24), v10+int32(16), v10+int32(15))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v40 == int32(1) {
							v121 = v13
							m.G0 = v10 + int32(32)
							return v121
						} else {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+28)))
							if v43 != 0 {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
								if v77 != 0 {
									v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										v121 = v118
										m.G0 = v10 + int32(32)
										return v121
									}
								} else {
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
									if v78&int32(1) == int32(0) {
										v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											v121 = v118
											m.G0 = v10 + int32(32)
											return v121
										}
									} else {
										v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
										v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
										if v84 == int64(9223372036854775807) {
											v87 = int32(0)
											v88 = F_errsave_start(m, v17)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												if v88 == int32(0) {
													v121 = v87
													m.G0 = v10 + int32(32)
													return v121
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v94 = m.ExcPending
													if v94 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return int32(0)
															} else {
																v121 = v87
																m.G0 = v10 + int32(32)
																return v121
															}
														}
													}
												}
											}
										} else {
											v106 = F_Int64GetDatum(m, v84+int64(1))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												v108 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
												v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													v121 = v118
													m.G0 = v10 + int32(32)
													return v121
												}
											}
										}
									}
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)))
								if v44&int32(1) != 0 {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
									if v77 != 0 {
										v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											v121 = v118
											m.G0 = v10 + int32(32)
											return v121
										}
									} else {
										v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
										if v78&int32(1) == int32(0) {
											v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int32(0)
											} else {
												v121 = v118
												m.G0 = v10 + int32(32)
												return v121
											}
										} else {
											v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
											v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
											if v84 == int64(9223372036854775807) {
												v87 = int32(0)
												v88 = F_errsave_start(m, v17)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													if v88 == int32(0) {
														v121 = v87
														m.G0 = v10 + int32(32)
														return v121
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v94 = m.ExcPending
														if v94 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return int32(0)
																} else {
																	v121 = v87
																	m.G0 = v10 + int32(32)
																	return v121
																}
															}
														}
													}
												}
											} else {
												v106 = F_Int64GetDatum(m, v84+int64(1))
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													v108 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
													*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
													v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														v121 = v118
														m.G0 = v10 + int32(32)
														return v121
													}
												}
											}
										}
									}
								} else {
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
									v48 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
									if v48 == int64(9223372036854775807) {
										v51 = int32(0)
										v52 = F_errsave_start(m, v17)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											if v52 == int32(0) {
												v121 = v51
												m.G0 = v10 + int32(32)
												return v121
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
													mBase = m.M
													v62 = m.ExcPending
													if v62 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1599), int32(_a_F_int8range_canonical_2))
														mBase = m.M
														v67 = m.ExcPending
														if v67 != 0 {
															return int32(0)
														} else {
															v121 = v51
															m.G0 = v10 + int32(32)
															return v121
														}
													}
												}
											}
										}
									} else {
										v70 = F_Int64GetDatum(m, v48+int64(1))
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return int32(0)
										} else {
											v72 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+29)) = uint8(v72)
											*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v70
											v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+20)))
											if v77 != 0 {
												v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int32(0)
												} else {
													v121 = v118
													m.G0 = v10 + int32(32)
													return v121
												}
											} else {
												v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)))
												if v78&int32(1) == int32(0) {
													v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int32(0)
													} else {
														v121 = v118
														m.G0 = v10 + int32(32)
														return v121
													}
												} else {
													v83 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
													v84 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
													if v84 == int64(9223372036854775807) {
														v87 = int32(0)
														v88 = F_errsave_start(m, v17)
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															if v88 == int32(0) {
																v121 = v87
																m.G0 = v10 + int32(32)
																return v121
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v94 = m.ExcPending
																if v94 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return int32(0)
																	} else {
																		F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1612), int32(_a_F_int8range_canonical_2))
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return int32(0)
																		} else {
																			v121 = v87
																			m.G0 = v10 + int32(32)
																			return v121
																		}
																	}
																}
															}
														}
													} else {
														v106 = F_Int64GetDatum(m, v84+int64(1))
														mBase = m.M
														v107 = m.ExcPending
														if v107 != 0 {
															return int32(0)
														} else {
															v108 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v10)+21)) = uint8(v108)
															*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v106
															v118 = F_range_serialize(m, v31, v10+int32(24), v10+int32(16), int32(0), v17)
															mBase = m.M
															v119 = m.ExcPending
															if v119 != 0 {
																return int32(0)
															} else {
																v121 = v118
																m.G0 = v10 + int32(32)
																return v121
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
			}
		}
	}
}
func F_int8um(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(-9223372036854775807-1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_int8um_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_int8um_1), int32(448), int32(_a_F_int8um_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v27 = F_Int64GetDatum(m, int64(0)-v4)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_intarray_push_array(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = F_intarray_concat_arrays(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return v13
							}
						} else {
							return v13
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int32(0)
						} else {
							return v13
						}
					} else {
						return v13
					}
				}
			}
		}
	}
}
func F_inter_lb(m *base.Module, l0 int32) int32 {
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
	var v15 float64
	_ = v15
	var v16 float64
	_ = v16
	var v17 float64
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v16
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v16
	v22 = int32(1)
	v24 = F_lseg_interpt_line(m, int32(0), v11, v13)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		if v24 != 0 {
			v55 = v22
			m.G0 = v11 + int32(32)
			return v55
		} else {
			v28 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
			v29 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
			*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
			*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v16
			*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
			*(*float64)(unsafe.Add(mBase, uint32(v11))) = v28
			v35 = F_lseg_interpt_line(m, int32(0), v11, v13)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				if v35 != 0 {
					v55 = v22
					m.G0 = v11 + int32(32)
					return v55
				} else {
					v37 = *(*float64)(unsafe.Add(mBase, uint32(v14)))
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
					*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v38
					*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v37
					*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v29
					*(*float64)(unsafe.Add(mBase, uint32(v11))) = v28
					v44 = F_lseg_interpt_line(m, int32(0), v11, v13)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						if v44 != 0 {
							v55 = v22
							m.G0 = v11 + int32(32)
							return v55
						} else {
							v46 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
							v47 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
							*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v38
							*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v37
							*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v47
							*(*float64)(unsafe.Add(mBase, uint32(v11))) = v46
							v53 = F_lseg_interpt_line(m, int32(0), v11, v13)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								v55 = v53
								m.G0 = v11 + int32(32)
								return v55
							}
						}
					}
				}
			}
		}
	}
}
func F_internalerrposition(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_internalerrposition[0]))
	if v4 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_internalerrposition[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_internalerrposition_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_internalerrposition_1), int32(1489), int32(_a_F_internalerrposition_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4*int32(100))+uint32(_c_F_internalerrposition[1]))) = l0
		return
	}
}
func F_intervaltypmodleastfield(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 < v2 {
		v56 = v2
		m.G0 = v7 + int32(16)
		return v56
	} else {
		v12 = int32(base.Ui32(l0) >> (uint(int32(16)) % 32))
		if base.Ui32(v12) <= base.Ui32(int32(3071)) {
			switch v12 - int32(2) {
			case 0, 4:
				v56 = int32(4)
				m.G0 = v7 + int32(16)
				return v56
			case 1, 3, 5:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg_internal(m, int32(_a_F_intervaltypmodleastfield_0), v7)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1248), int32(_a_F_intervaltypmodleastfield_2))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 2:
				v56 = int32(5)
				m.G0 = v7 + int32(16)
				return v56
			case 6:
				v56 = int32(3)
				m.G0 = v7 + int32(16)
				return v56
			default:
				switch v12 - int32(1024) {
				case 0, 8:
					v56 = int32(2)
					m.G0 = v7 + int32(16)
					return v56
				case 1, 2, 3, 4, 5, 6, 7:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(_a_F_intervaltypmodleastfield_0), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1248), int32(_a_F_intervaltypmodleastfield_2))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				default:
					if v12 == int32(2048) {
						v56 = int32(1)
						m.G0 = v7 + int32(16)
						return v56
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg_internal(m, int32(_a_F_intervaltypmodleastfield_0), v7)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1248), int32(_a_F_intervaltypmodleastfield_2))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
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
		} else {
			if base.Ui32(v12) <= base.Ui32(int32(_a_F_intervaltypmodleastfield_3)) {
				switch v12 - int32(3072) {
				case 0, 8:
					v56 = int32(1)
					m.G0 = v7 + int32(16)
					return v56
				case 1, 2, 3, 4, 5, 6, 7:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(_a_F_intervaltypmodleastfield_0), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1248), int32(_a_F_intervaltypmodleastfield_2))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				default:
					if v12 != int32(_a_F_intervaltypmodleastfield_4) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
							F_errmsg_internal(m, int32(_a_F_intervaltypmodleastfield_0), v7)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1248), int32(_a_F_intervaltypmodleastfield_2))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v56 = v2
						m.G0 = v7 + int32(16)
						return v56
					}
				}
			} else {
				switch v12 - int32(_a_F_intervaltypmodleastfield_5) {
				case 0, 8:
					v56 = v2
					m.G0 = v7 + int32(16)
					return v56
				case 1, 2, 3, 4, 5, 6, 7:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg_internal(m, int32(_a_F_intervaltypmodleastfield_0), v7)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1248), int32(_a_F_intervaltypmodleastfield_2))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				default:
					if v12 == int32(_a_F_intervaltypmodleastfield_6) {
						v56 = v2
						m.G0 = v7 + int32(16)
						return v56
					} else {
						if v12 != int32(_a_F_intervaltypmodleastfield_7) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v42 = m.ExcPending
							if v42 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
								F_errmsg_internal(m, int32(_a_F_intervaltypmodleastfield_0), v7)
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1248), int32(_a_F_intervaltypmodleastfield_2))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v56 = v2
							m.G0 = v7 + int32(16)
							return v56
						}
					}
				}
			}
		}
	}
}
func F_invariant_l_offset(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
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
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v10 = F_PageGetItemIdCareful_2(m, l0, v8, v9, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v17 = F__bt_compare(m, v15, l1, v16, l2)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v14 == int32(0) {
				return base.B2i32(v17 <= int32(0))
			} else {
				if v17 == int32(0) {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v30 = v26 + v27&int32(_a_F_invariant_l_offset_0)
					v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+16)))
					v33 = v26 + v32
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+12)))
					if v34&int32(1) != 0 {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
						v45 = base.B2i32(v37 == int32(0)) | base.B2i32(base.Ui32(int32(1)) < base.Ui32(l2&int32(_a_F_invariant_l_offset_1)))
					} else {
						v45 = int32(0)
					}
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+192))
					v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+10)))
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+7)))
					if v49&int32(32) != 0 {
						v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+4)))
						if v52&int32(_a_F_invariant_l_offset_2) != 0 {
							v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+8)))
							if v48 < v55 {
								v66 = v48
							} else {
								v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+8)))
								v63 = v57
								v66 = base.I32_extend16_s(v63)
							}
						} else {
							v59 = v52 & int32(4095)
							if v48 < v59 {
								v66 = v48
							} else {
								v66 = v59
							}
						}
					} else {
						v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+8)))
						if v48 < v61 {
							v66 = v48
						} else {
							v63 = v61
							v66 = base.I32_extend16_s(v63)
						}
					}
					v69 = F_BTreeTupleGetHeapTIDCareful(m, l0, v30, v45)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if v66 == v71 {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							v74 = int32(0)
							return base.B2i32(v73 == v74) & base.B2i32(v69 != v74)
						} else {
							return base.B2i32(v71 < v66)
						}
					}
				} else {
					return int32(base.Ui32(v17) >> (uint(int32(31)) % 32))
				}
			}
		}
	}
}
func F_is_encoding_supported_by_icu(m *base.Module, l0 int32) int32 {
	return base.B2i32(base.Ui32(l0) < base.Ui32(int32(35))) & base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(l0))%64)&int64(34357509982) != int64(0))
}
func F_is_member_of_role_nosuper(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	if l0 != l1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v4 = int32(0)
	v7 = F_roles_is_member_of(m, l0, v4, v4, v4)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v51 = int32(1)
	goto L3
L3:
	;
	return v51
L4:
	;
	return int32(0)
L5:
	;
	v11 = int32(0)
	if v7 == v11 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v51 = v49
	goto L3
L7:
	;
	v49 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v17 <= int32(0) {
		v43 = v11
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v49 = v43
	goto L6
L11:
	;
	v20 = int32(0)
	if v20 < v17 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v23 = v17
	goto L14
L13:
	;
	v23 = v20
	goto L14
L14:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v26 = int32(0)
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v24+v26<<(uint(int32(2))%32))))
	v35 = base.B2i32(v34 == l1)
	if v34 == l1 {
		v43 = v35
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v43 = v35
	goto L10
L17:
	;
	v37 = v26 + int32(1)
	if v37 != v23 {
		v26 = v37
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_isn_out(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13949(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_iso8859_to_utf8(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn13950(m, l0, int32(_a_F_iso8859_to_utf8_0), int32(_a_F_iso8859_to_utf8_1), int32(133), int32(_a_F_iso8859_to_utf8_2), int32(_a_F_iso8859_to_utf8_3), int32(_a_F_iso8859_to_utf8_4), int32(19), int32(9))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_iso_to_koi8r(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13958(m, l0, int32(_a_F_iso_to_koi8r_0), int32(22), int32(25))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_iso_to_mic(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13951(m, l0, int32(_a_F_iso_to_mic_0), int32(25), int32(139))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_iso_to_win1251(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn13958(m, l0, int32(_a_F_iso_to_win1251_0), int32(23), int32(25))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_iswpunct(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	if base.Ui32(l0) <= base.Ui32(int32(_a_F_iswpunct_0)) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(8))%32)))+uint32(_c_F_iswpunct[0]))))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(l0)>>(uint(int32(3))%32))&int32(31)|v10<<(uint(int32(5))%32))+uint32(_c_F_iswpunct[0]))))
		v21 = int32(base.Ui32(v14)>>(uint(l0&int32(7))%32)) & int32(1)
	} else {
		v21 = int32(0)
	}
	return v21
}
func F_iswupper(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	v3 = F_casemap(m, l0, int32(0))
	return base.B2i32(v3 != l0)
}
func F_iteratorFromContainer(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v7 = F_palloc0(m, int32(40))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v15 = l0 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v15
		v18 = v13 & int32(268435455)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v22 = v20 & int32(1610612736)
		if v22 != int32(536870912) {
			if v22 == int32(1073741824) {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v15 + v18<<(uint(int32(2))%32)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v35 = int32(base.Ui32(v31)>>(uint(int32(28))%32)) & int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v7)+8)) = uint8(v35)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(0)
				return v7
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_iteratorFromContainer_0), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_iteratorFromContainer_1), int32(1043), int32(_a_F_iteratorFromContainer_2))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v15 + v18<<(uint(int32(3))%32)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(2)
			return v7
		}
	}
}
func F_ivfflatbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 float64
	_ = v334
	var v340 float64
	_ = v340
	var v344 int32
	_ = v344
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int64
	_ = v393
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v422 int32
	_ = v422
	var v448 int32
	_ = v448
	v20 = m.G0
	v22 = v20 - int32(_a_F_ivfflatbulkdelete_0)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v33 = F_palloc0(m, int32(40))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v35 = l1
	goto L5
L5:
	;
	v48 = int32(1)
	goto L7
L6:
	;
	v35 = v33
	goto L5
L7:
	;
	v56 = F_ReadBuffer(m, v24, v48)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_bms_free(m, v26)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L75
	}
L9:
	;
	F_LockBuffer(m, v56, int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	if v56 < int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_ivfflatbulkdelete[0]))) = v48
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+16)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v78+v214)))
	F_UnlockReleaseBuffer(m, v56)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L30
	}
L12:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v79) {
		goto L16
	} else {
		goto L17
	}
L13:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbulkdelete[1]))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v64+(v56^int32(-1))<<(uint(int32(2))%32))))
	v78 = v70
	goto L12
L14:
	;
	goto L15
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatbulkdelete[2]))
	v78 = v72 + v56<<(uint(int32(13))%32) + int32(-8192)
	goto L12
L16:
	;
	v87 = int32(base.Ui32(v79+int32(_a_F_ivfflatbulkdelete_1)) >> (uint(int32(2)) % 32))
	goto L18
L17:
	;
	v87 = int32(0)
	goto L18
L18:
	;
	v89 = v87 & int32(_a_F_ivfflatbulkdelete_2)
	if v89 == int32(0) {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v92 = int32(1)
	v94 = v78 + int32(20)
	v98 = (v87 + v92) & int32(_a_F_ivfflatbulkdelete_2)
	if base.Ui32(int32(3)) <= base.Ui32(v98) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v101 = int32(2)
	if base.Ui32(v98) <= base.Ui32(v101) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v163 = v92
	goto L22
L22:
	;
	v183 = v163 << (uint(int32(2)) % 32)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v183+v94)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v78+v188&int32(_a_F_ivfflatbulkdelete_3))))
	*(*int32)(unsafe.Add(mBase, uint32(v22+v183)+uint32(_c_F_ivfflatbulkdelete[3]))) = v192
	goto L11
L23:
	;
	v104 = v101
	goto L25
L24:
	;
	v104 = v98
	goto L25
L25:
	;
	v105 = int32(1)
	v106 = v104 - v105
	v113 = v105
	v119 = int32(0)
	goto L26
L26:
	;
	v132 = int32(2)
	v133 = v113 << (uint(v132) % 32)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v94+v133)))
	v139 = int32(_a_F_ivfflatbulkdelete_3)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v78+v138&v139)))
	*(*int32)(unsafe.Add(mBase, uint32(v22+v133)+uint32(_c_F_ivfflatbulkdelete[3]))) = v142
	v145 = v133 + int32(4)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v94+v145)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v78+v150&v139)))
	*(*int32)(unsafe.Add(mBase, uint32(v22+v145)+uint32(_c_F_ivfflatbulkdelete[3]))) = v154
	v157 = v113 + v132
	v159 = v119 + v132
	if v159 != v106&int32(-2) {
		v113 = v157
		v119 = v159
		goto L26
	} else {
		goto L28
	}
L27:
	;
	if v106&v105 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v163 = v157
	goto L22
L30:
	;
	if v89 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v226 = int32(1)
	goto L34
L32:
	;
	goto L33
L33:
	;
	if v216 != int32(-1) {
		v48 = v216
		goto L7
	} else {
		goto L74
	}
L34:
	;
	v239 = int32(-1)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v226&int32(_a_F_ivfflatbulkdelete_2)<<(uint(int32(2))%32)+v22)+uint32(_c_F_ivfflatbulkdelete[3])))
	if v247 == v239 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L33
L36:
	;
	v422 = v226 + int32(1)
	if base.Ui32(v422&int32(_a_F_ivfflatbulkdelete_2)) <= base.Ui32(v89) {
		v226 = v422
		goto L34
	} else {
		goto L73
	}
L37:
	;
	v254 = v239
	v259 = v247
	goto L38
L38:
	;
	v269 = int32(0)
	F_vacuum_delay_point(m, v269)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	if v385 == int32(-1) {
		goto L36
	} else {
		goto L71
	}
L40:
	;
	v273 = int32(0)
	v275 = F_ReadBufferExtended(m, v24, v273, v259, v273, v26)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_LockBufferForCleanup(m, v275)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v279 = F_GenericXLogStart(m, v24)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	v368 = base.B2i32(v355 <= int32(0))
	if v355 <= int32(0) {
		goto L56
	} else {
		goto L57
	}
L44:
	;
	v282 = F_GenericXLogRegisterBuffer(m, v279, v275, int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282)+12)))
	if base.Ui32(v284) < base.Ui32(int32(25)) {
		v355 = v269
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v288 = v284 + int32(_a_F_ivfflatbulkdelete_1)
	if v288&int32(_a_F_ivfflatbulkdelete_4) == int32(0) {
		v355 = v269
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v300 = int32(1)
	v307 = v269
	goto L48
L48:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v282+int32(20)+v300<<(uint(int32(2))%32))))
	v326 = m.T0[l2].(func(*base.Module, int32, int32) int32)(m, v282+v322&int32(_a_F_ivfflatbulkdelete_3), l3)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	v355 = v344
	goto L43
L50:
	;
	if v300 != int32(base.Ui32(v288)>>(uint(int32(2))%32))&int32(_a_F_ivfflatbulkdelete_2) {
		v300 = v300 + int32(1)
		v307 = v344
		goto L48
	} else {
		goto L55
	}
L51:
	;
	if v326 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v330 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v22+int32(16)+v307<<(uint(v330)%32)))) = uint16(v300)
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v35)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+16)) = base.F64_add(v334, float64(1))
	v344 = v307 + v330
	goto L50
L53:
	;
	goto L54
L54:
	;
	v340 = *(*float64)(unsafe.Add(mBase, uint32(v35)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v35)+8)) = base.F64_add(v340, float64(1))
	v344 = v307
	goto L50
L55:
	;
	goto L49
L56:
	;
	v369 = v254
	goto L58
L57:
	;
	v369 = v259
	goto L58
L58:
	;
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v282)+16)))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v282+v370)))
	if v368 == int32(0) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v254 != int32(-1) {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	F_PageIndexMultiDelete(m, v282, v22+int32(16), v355)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_pfree(m, v279)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	F_GenericXLogFinish(m, v279)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L59
L65:
	;
	goto L59
L66:
	;
	v385 = v254
	goto L68
L67:
	;
	v385 = v369
	goto L68
L68:
	;
	F_UnlockReleaseBuffer(m, v275)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	if v372 != int32(-1) {
		v254 = v385
		v259 = v372
		goto L38
	} else {
		goto L70
	}
L70:
	;
	goto L39
L71:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_ivfflatbulkdelete[3]))) = uint16(v226)
	v393 = *(*int64)(unsafe.Add(mBase, uint32(v22)+uint32(_c_F_ivfflatbulkdelete[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v393
	v397 = int32(-1)
	F_IvfflatUpdateList(m, v24, v22+int32(8), v385, v397, v397, int32(0))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L36
L73:
	;
	goto L35
L74:
	;
	goto L8
L75:
	;
	m.G0 = v22 + int32(_a_F_ivfflatbulkdelete_0)
	return v35
}
func F_ivfflatinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v97 float64
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v157 float64
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 float64
	_ = v170
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 float64
	_ = v183
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v208 float64
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
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
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v387 int64
	_ = v387
	var v394 int32
	_ = v394
	var v419 int32
	_ = v419
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	v22 = m.G0
	v24 = v22 - int32(32)
	m.G0 = v24
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v26 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L5
	} else {
		goto L87
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[0]))
	v35 = F_AllocSetContextCreateInternal(m, v30, int32(_a_F_ivfflatinsert_0), int32(0), int32(_a_F_ivfflatinsert_1), int32(_a_F_ivfflatinsert_2))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	m.G0 = v24 + int32(32)
	return int32(0)
L5:
	;
	return int32(0)
L6:
	;
	v39 = int32(_a_F_ivfflatinsert_3)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[0])) = v35
	v43 = F_IvfflatGetTypeInfo(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v46 = F_pg_detoast_datum(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v46
	v50 = F_HnswOptionalProcInfo(m, l0, int32(2))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[0])) = v40
	F_MemoryContextDelete(m, v35)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L5
	} else {
		goto L86
	}
L10:
	;
	if v50 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v54 = F_IvfflatCheckNorm(m, v50, v53, v46)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	v62 = v46
	goto L13
L13:
	;
	v63 = int32(0)
	F_IvfflatGetMetaPageInfo(m, l0, v63, v63)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L17
	}
L14:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v58 = F_HnswNormValue(m, v43, v53, v46)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+28)) = v58
	v62 = v58
	goto L13
L17:
	;
	v67 = int32(1)
	v69 = F_index_getprocinfo(m, l0, v67, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v74 = int32(1)
	v82 = v74
	v83 = v74
	v84 = int32(-1)
	v85 = v74
	v97 = float64(1.7976931348623157e+308)
	goto L19
L19:
	;
	v99 = F_ReadBuffer(m, l0, v85)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L21
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+24)) = uint16(v193)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v194
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v222 = F_index_form_tuple(m, v219, v24+int32(28), l2)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L39
	}
L21:
	;
	F_LockBuffer(m, v99, int32(1))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	if v99 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+16)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v121+v210)))
	F_UnlockReleaseBuffer(m, v99)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L5
	} else {
		goto L37
	}
L24:
	;
	v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121)+12)))
	if base.Ui32(v122) < base.Ui32(int32(25)) {
		v193 = v82
		v194 = v83
		v195 = v84
		v208 = v97
		goto L23
	} else {
		goto L28
	}
L25:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[1]))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v107+(v99^int32(-1))<<(uint(int32(2))%32))))
	v121 = v113
	goto L24
L26:
	;
	goto L27
L27:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[2]))
	v121 = v115 + v99<<(uint(int32(13))%32) + int32(-8192)
	goto L24
L28:
	;
	v126 = v122 + int32(_a_F_ivfflatinsert_4)
	if v126&int32(_a_F_ivfflatinsert_5) == int32(0) {
		v193 = v82
		v194 = v83
		v195 = v84
		v208 = v97
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v139 = int32(1)
	v142 = v82
	v143 = v83
	v144 = v84
	v157 = v97
	goto L30
L30:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v121+int32(20)+v139<<(uint(int32(2))%32))))
	v165 = v121 + v162&int32(_a_F_ivfflatinsert_6)
	v168 = F_FunctionCall2Coll(m, v69, v72, v62, v165+int32(8))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v193 = v180
	v194 = v181
	v195 = v182
	v208 = v183
	goto L23
L32:
	;
	v170 = *(*float64)(unsafe.Add(mBase, uint32(v168)))
	v172 = int32(0)
	if base.B2i32(base.F64_lt(v170, v157) == v172)&base.B2i32(v144 != int32(-1)) == v172 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	v180 = v139
	v181 = v85
	v182 = v179
	v183 = v170
	goto L35
L34:
	;
	v180 = v142
	v181 = v143
	v182 = v144
	v183 = v157
	goto L35
L35:
	;
	if base.B2i32(v139 == int32(base.Ui32(v126)>>(uint(int32(2))%32))&int32(_a_F_ivfflatinsert_7)) == int32(0) {
		v139 = v139 + int32(1)
		v142 = v180
		v143 = v181
		v144 = v182
		v157 = v183
		goto L30
	} else {
		goto L36
	}
L36:
	;
	goto L31
L37:
	;
	if v212 != int32(-1) {
		v82 = v193
		v83 = v194
		v84 = v195
		v85 = v212
		v97 = v208
		goto L19
	} else {
		goto L38
	}
L38:
	;
	goto L20
L39:
	;
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v222)+4)) = uint16(v224)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v226
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+6)))
	v229 = F_ReadBuffer(m, l0, v195)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	F_LockBuffer(m, v229, int32(2))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v234 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L43
	}
L42:
	;
	v378 = int32(0)
	v380 = F_PageAddItemExtended(m, v360, v222, v253, v378, v378)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L81
	}
L43:
	;
	v237 = F_GenericXLogRegisterBuffer(m, v234, v229, int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v239 = int32(4)
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237)+14)))
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v237)+12)))
	v242 = v240 - v241
	if v242 <= v239 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v253 = (v228&int32(_a_F_ivfflatinsert_8) + int32(7)) & int32(_a_F_ivfflatinsert_9)
	if base.Ui32(v253) <= base.Ui32(v245-int32(4)) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v245 = v239
	goto L48
L47:
	;
	v245 = v242
	goto L48
L48:
	;
	goto L45
L49:
	;
	v358 = v229
	v360 = v237
	v361 = v195
	v362 = v234
	goto L42
L50:
	;
	goto L51
L51:
	;
	v256 = v229
	v258 = v237
	v260 = v234
	goto L52
L52:
	;
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+16)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v258+v276)))
	if v278 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_LockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L5
	} else {
		goto L68
	}
L54:
	;
	F_pfree(m, v260)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L5
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L53
L57:
	;
	F_UnlockReleaseBuffer(m, v256)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v285 = F_ReadBuffer(m, l0, v278)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	F_LockBuffer(m, v285, int32(2))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v290 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v293 = F_GenericXLogRegisterBuffer(m, v290, v285, int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v295 = int32(4)
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v293)+14)))
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v293)+12)))
	v298 = v296 - v297
	if v298 <= v295 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if base.Ui32(v301-int32(4)) < base.Ui32(v253) {
		v256 = v285
		v258 = v293
		v260 = v290
		goto L52
	} else {
		goto L67
	}
L64:
	;
	v301 = v295
	goto L66
L65:
	;
	v301 = v298
	goto L66
L66:
	;
	goto L63
L67:
	;
	v358 = v285
	v360 = v293
	v361 = v278
	v362 = v290
	goto L42
L68:
	;
	v309 = F_HnswNewBuffer(m, l0, int32(0))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	v315 = F_GenericXLogRegisterBuffer(m, v260, v309, int32(1))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	F_PageInit(m, v315, int32(_a_F_ivfflatinsert_1), int32(8))
	mBase = m.M
	v320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v315)+16)))
	v321 = v315 + v320
	v322 = int32(_a_F_ivfflatinsert_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v321)+6)) = uint16(v322)
	*(*int32)(unsafe.Add(mBase, uint32(v321))) = int32(-1)
	goto L72
L72:
	;
	if v309 < int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v345 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v258+v345))) = v344
	F_GenericXLogFinish(m, v260)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L77
	}
L74:
	;
	v329 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[3]))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329+(v309^int32(-1))<<(uint(int32(6))%32))+16))
	v344 = v335
	goto L73
L75:
	;
	goto L76
L76:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[4]))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v337+v309<<(uint(int32(6))%32)+int32(-64))+16))
	v344 = v343
	goto L73
L77:
	;
	F_UnlockReleaseBuffer(m, v256)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v352 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v355 = F_GenericXLogRegisterBuffer(m, v352, v309, int32(0))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v358 = v309
	v360 = v355
	v361 = v344
	v362 = v352
	goto L42
L81:
	;
	if v380 == int32(0) {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_IvfflatCommitBuffer(m, v358, v362)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	if v361 == v195 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	v387 = *(*int64)(unsafe.Add(mBase, uint32(v24)+20))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v387
	F_IvfflatUpdateList(m, l0, v24+int32(8), v361, v195, int32(-1), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	goto L9
L86:
	;
	goto L4
L87:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v450 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ivfflatinsert_11), v24)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_ivfflatinsert_12), int32(174), int32(_a_F_ivfflatinsert_13))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L5
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
