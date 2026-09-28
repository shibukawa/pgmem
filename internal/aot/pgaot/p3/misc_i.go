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
				v68 = F_expression_tree_walker_impl(m, l0, int32(1128), l1)
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
				v68 = F_expression_tree_walker_impl(m, l0, int32(1128), l1)
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
				v68 = F_expression_tree_walker_impl(m, l0, int32(1128), l1)
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
							F_errfinish(m, int32(_a_F_IncrementVarSublevelsUp_walker_1), int32(836), int32(_a_F_IncrementVarSublevelsUp_walker_2))
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
				v68 = F_expression_tree_walker_impl(m, l0, int32(1128), l1)
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
				v68 = F_expression_tree_walker_impl(m, l0, int32(1128), l1)
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
				v88 = F_query_tree_walker_impl(m, l0, int32(1128), l1, int32(16))
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
					if v7 != int32(321) {
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
						if base.Ui32(v18) < base.Ui32(v19) {
						} else {
							v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v21 + v18
						}
					}
					v68 = F_expression_tree_walker_impl(m, l0, int32(1128), l1)
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 float64
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
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
		v34 = int32(3)
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(v34)%32))+104))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
		v43 = *(*int32)(unsafe.Add(mBase, uint32(v32+v39<<(uint(v34)%32))+96))
		if v43 != int32(1562) {
			if v37 < int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v177 = m.ExcPending
				if v177 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v180 = m.ExcPending
					if v180 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_InitBuildState_1_0), int32(0))
						mBase = m.M
						v184 = m.ExcPending
						if v184 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_InitBuildState_1_1), int32(706), int32(_a_F_InitBuildState_1_2))
							mBase = m.M
							v189 = m.ExcPending
							if v189 != 0 {
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
					v193 = m.ExcPending
					if v193 != 0 {
						return
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v196 = m.ExcPending
						if v196 != 0 {
							return
						} else {
							v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v198
							F_errmsg(m, int32(_a_F_InitBuildState_1_3), v9)
							mBase = m.M
							v202 = m.ExcPending
							if v202 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InitBuildState_1_1), int32(711), int32(_a_F_InitBuildState_1_2))
								mBase = m.M
								v207 = m.ExcPending
								if v207 != 0 {
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
						v211 = m.ExcPending
						if v211 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v214 = m.ExcPending
							if v214 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_InitBuildState_1_4), int32(0))
								mBase = m.M
								v218 = m.ExcPending
								if v218 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_InitBuildState_1_1), int32(716), int32(_a_F_InitBuildState_1_2))
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
									v91 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[1]))
									F_LWLockInitialize(m, l0+int32(80), v91)
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										v97 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[1]))
										F_LWLockInitialize(m, l0+int32(96), v97)
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											v103 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[1]))
											F_LWLockInitialize(m, l0+int32(116), v103)
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
												return
											} else {
												v109 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[1]))
												F_LWLockInitialize(m, l0+int32(140), v109)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = l0 - int32(-64)
													v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
													v118 = F_log(m, base.F64_convert_i32_s(v116))
													mBase = m.M
													*(*float64)(unsafe.Add(mBase, uint32(l0)+168)) = base.F64_div(float64(1), v118)
													v121 = int32(63)
													v123 = base.I32_div_u_s(int32(1358), v116)
													v125 = v123 - int32(2)
													if base.Ui32(v121) <= base.Ui32(v125) {
														v128 = v121
													} else {
														v128 = v125
													}
													*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v128
													v131 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[2]))
													v133 = int32(_a_F_InitBuildState_1_5)
													v136 = F_GenerationContextCreate(m, v131, int32(_a_F_InitBuildState_1_6), v133, v133, v133)
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v136
														v140 = *(*int32)(unsafe.Add(mBase, _c_F_InitBuildState_1[2]))
														v145 = F_AllocSetContextCreateInternal(m, v140, int32(_a_F_InitBuildState_1_7), int32(0), int32(_a_F_InitBuildState_1_8), int32(_a_F_InitBuildState_1_9))
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = int32(0)
															*(*int64)(unsafe.Add(mBase, uint32(l0)+196)) = int64(0)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = int32(_a_F_InitBuildState_1_10)
															*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v145
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
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v161 = m.ExcPending
			if v161 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v164 = m.ExcPending
				if v164 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_InitBuildState_1_11), int32(0))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitBuildState_1_1), int32(700), int32(_a_F_InitBuildState_1_2))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
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
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
	if v10 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
		if v12 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v284 = m.ExcPending
			if v284 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(_a_F_InitProcess_0), int32(0))
				mBase = m.M
				v288 = m.ExcPending
				if v288 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_InitProcess_1), int32(403), int32(_a_F_InitProcess_2))
					mBase = m.M
					v293 = m.ExcPending
					if v293 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
			if v14 == int32(1) {
				F_RegisterPostmasterChildActive(m)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
					v22 = v20 - int32(3)
					if base.Ui32(v22) <= base.Ui32(int32(4)) {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_InitProcess[4])))
						v29 = v27
					} else {
						v29 = int32(24)
					}
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
					v32 = v29 + v31
					v35 = base.AtomicRmwXchg32(m, v31, int32(20), int32(1))
					if v35 != 0 {
						F_s_lock(m, v31+int32(20), int32(_a_F_InitProcess_3))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v42 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5])) = v43
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
							v47 = int32(0)
							if base.B2i32(v46 == v47)|base.B2i32(v46 == v32) == v47 {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
								*(*int32)(unsafe.Add(mBase, uint32(v54))) = v56
								v58 = int32(_a_F_InitProcess_4)
								v59 = int32(4)
								v60 = v46 - v59
								*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1])) = v60
								v63 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
								v64 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v63)+20)), uint32(v64))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
								v71 = base.I32_div_s(v60-v68, int32(768))
								*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[6])) = v71
								*(*int32)(unsafe.Add(mBase, uint32(v46)+572)) = v64
								*(*uint8)(unsafe.Add(mBase, uint32(v46)+568)) = uint8(v64)
								*(*int32)(unsafe.Add(mBase, uint32(v46)+412)) = v64
								v79 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v46))) = v79
								*(*int64)(unsafe.Add(mBase, uint32(v46)+44)) = v79
								v84 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7]))
								*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(v46)+36)) = v71
								*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v84
								*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v64
								*(*int64)(unsafe.Add(mBase, uint32(v46)+16)) = v79
								v94 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
								*(*int32)(unsafe.Add(mBase, uint32(v46)+332)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v94
								*(*int64)(unsafe.Add(mBase, uint32(v46)+380)) = v79
								*(*uint16)(unsafe.Add(mBase, uint32(v46)+340)) = uint16(v64)
								*(*int64)(unsafe.Add(mBase, uint32(v46)+388)) = v79
								*(*uint8)(unsafe.Add(mBase, uint32(v46)+32)) = uint8(base.B2i32(v94 == v59))
								v109 = base.AtomicRmwXchg64(m, v46, int32(404), v79)
								v111 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
								*(*int32)(unsafe.Add(mBase, uint32(v111)+648)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(v111)+612)) = v64
								*(*int32)(unsafe.Add(mBase, uint32(v111)+340)) = v64
								*(*int64)(unsafe.Add(mBase, uint32(v111)+624)) = v79
								*(*uint8)(unsafe.Add(mBase, uint32(v111)+616)) = uint8(v64)
								*(*int64)(unsafe.Add(mBase, uint32(v111)+584)) = v79
								*(*int64)(unsafe.Add(mBase, uint32(v111)+592)) = v79
								*(*int64)(unsafe.Add(mBase, uint32(v111)+597)) = v79
								*(*int64)(unsafe.Add(mBase, uint32(v111)+640)) = v79
								*(*int64)(unsafe.Add(mBase, uint32(v111)+632)) = int64(-1)
								F_OwnLatch(m, v111+int32(316))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return
								} else {
									F_SwitchToSharedLatch(m)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return
									} else {
										v139 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
										*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[8])) = v139 + int32(648)
										v145 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+332))
										F_PGSemaphoreReset(m, v146)
										mBase = m.M
										F_on_shmem_exit(m, int32(1227), int64(0))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return
										} else {
											v152 = int32(_a_F_InitProcess_5)
											v153 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9]))
											v156 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10]))
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v156
											v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
											v163 = F_palloc(m, v160<<(uint(int32(2))%32))
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12])) = v163
												v168 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
												v171 = F_palloc(m, v168*int32(24))
												mBase = m.M
												v172 = m.ExcPending
												if v172 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13])) = v171
													v176 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[14])) = v176
													v180 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
													v183 = F_palloc(m, v180<<(uint(int32(2))%32))
													mBase = m.M
													v184 = m.ExcPending
													if v184 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[15])) = v183
														v188 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
														v191 = F_palloc(m, v188<<(uint(int32(2))%32))
														mBase = m.M
														v192 = m.ExcPending
														if v192 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[16])) = v191
															v196 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
															v198 = base.I32_div_s(v196, int32(2))
															v201 = F_palloc(m, v198*int32(12))
															mBase = m.M
															v202 = m.ExcPending
															if v202 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[17])) = v201
																v206 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
																v209 = F_palloc(m, v206<<(uint(int32(2))%32))
																mBase = m.M
																v210 = m.ExcPending
																if v210 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[18])) = v209
																	v214 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[19])) = v214
																	v219 = F_palloc(m, v214*int32(20))
																	mBase = m.M
																	v220 = m.ExcPending
																	if v220 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[20])) = v219
																		v224 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[21])) = v224 << (uint(int32(2)) % 32)
																		v230 = F_palloc(m, v224*int32(80))
																		mBase = m.M
																		v231 = m.ExcPending
																		if v231 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v153
																			*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[22])) = v230
																			v237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
																			if v237 != 0 {
																				F_AttachSharedMemoryStructs(m)
																				mBase = m.M
																				v239 = m.ExcPending
																				if v239 != 0 {
																					return
																				} else {
																					m.G0 = v7 + int32(16)
																					return
																				}
																			} else {
																				m.G0 = v7 + int32(16)
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
								v244 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
								v245 = int32(0)
								atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v244)+20)), uint32(v245))
								v249 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
								F_errstart_cold(m, int32(22), v245)
								mBase = m.M
								v253 = m.ExcPending
								if v253 != 0 {
									return
								} else {
									F_errcode(m, int32(_a_F_InitProcess_6))
									mBase = m.M
									v256 = m.ExcPending
									if v256 != 0 {
										return
									} else {
										if v249 == int32(6) {
											v295 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[23]))
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v295
											F_errmsg(m, int32(_a_F_InitProcess_7), v7)
											mBase = m.M
											v299 = m.ExcPending
											if v299 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitProcess_1), int32(455), int32(_a_F_InitProcess_2))
												mBase = m.M
												v304 = m.ExcPending
												if v304 != 0 {
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
											v262 = m.ExcPending
											if v262 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_InitProcess_1), int32(458), int32(_a_F_InitProcess_2))
												mBase = m.M
												v267 = m.ExcPending
												if v267 != 0 {
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
						v42 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
						*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5])) = v43
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
						v47 = int32(0)
						if base.B2i32(v46 == v47)|base.B2i32(v46 == v32) == v47 {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
							*(*int32)(unsafe.Add(mBase, uint32(v54))) = v56
							v58 = int32(_a_F_InitProcess_4)
							v59 = int32(4)
							v60 = v46 - v59
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1])) = v60
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
							v64 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v63)+20)), uint32(v64))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
							v71 = base.I32_div_s(v60-v68, int32(768))
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[6])) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v46)+572)) = v64
							*(*uint8)(unsafe.Add(mBase, uint32(v46)+568)) = uint8(v64)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+412)) = v64
							v79 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v46))) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v46)+44)) = v79
							v84 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7]))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v46)+36)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v84
							*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v64
							*(*int64)(unsafe.Add(mBase, uint32(v46)+16)) = v79
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+332)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v46)+380)) = v79
							*(*uint16)(unsafe.Add(mBase, uint32(v46)+340)) = uint16(v64)
							*(*int64)(unsafe.Add(mBase, uint32(v46)+388)) = v79
							*(*uint8)(unsafe.Add(mBase, uint32(v46)+32)) = uint8(base.B2i32(v94 == v59))
							v109 = base.AtomicRmwXchg64(m, v46, int32(404), v79)
							v111 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v111)+648)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v111)+612)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v111)+340)) = v64
							*(*int64)(unsafe.Add(mBase, uint32(v111)+624)) = v79
							*(*uint8)(unsafe.Add(mBase, uint32(v111)+616)) = uint8(v64)
							*(*int64)(unsafe.Add(mBase, uint32(v111)+584)) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v111)+592)) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v111)+597)) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v111)+640)) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v111)+632)) = int64(-1)
							F_OwnLatch(m, v111+int32(316))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								F_SwitchToSharedLatch(m)
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
									*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[8])) = v139 + int32(648)
									v145 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+332))
									F_PGSemaphoreReset(m, v146)
									mBase = m.M
									F_on_shmem_exit(m, int32(1227), int64(0))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return
									} else {
										v152 = int32(_a_F_InitProcess_5)
										v153 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9]))
										v156 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10]))
										*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v156
										v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
										v163 = F_palloc(m, v160<<(uint(int32(2))%32))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12])) = v163
											v168 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
											v171 = F_palloc(m, v168*int32(24))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13])) = v171
												v176 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[14])) = v176
												v180 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
												v183 = F_palloc(m, v180<<(uint(int32(2))%32))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[15])) = v183
													v188 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
													v191 = F_palloc(m, v188<<(uint(int32(2))%32))
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[16])) = v191
														v196 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
														v198 = base.I32_div_s(v196, int32(2))
														v201 = F_palloc(m, v198*int32(12))
														mBase = m.M
														v202 = m.ExcPending
														if v202 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[17])) = v201
															v206 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
															v209 = F_palloc(m, v206<<(uint(int32(2))%32))
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[18])) = v209
																v214 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[19])) = v214
																v219 = F_palloc(m, v214*int32(20))
																mBase = m.M
																v220 = m.ExcPending
																if v220 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[20])) = v219
																	v224 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[21])) = v224 << (uint(int32(2)) % 32)
																	v230 = F_palloc(m, v224*int32(80))
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v153
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[22])) = v230
																		v237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
																		if v237 != 0 {
																			F_AttachSharedMemoryStructs(m)
																			mBase = m.M
																			v239 = m.ExcPending
																			if v239 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(16)
																				return
																			}
																		} else {
																			m.G0 = v7 + int32(16)
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
							v244 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
							v245 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v244)+20)), uint32(v245))
							v249 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
							F_errstart_cold(m, int32(22), v245)
							mBase = m.M
							v253 = m.ExcPending
							if v253 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_InitProcess_6))
								mBase = m.M
								v256 = m.ExcPending
								if v256 != 0 {
									return
								} else {
									if v249 == int32(6) {
										v295 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[23]))
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v295
										F_errmsg(m, int32(_a_F_InitProcess_7), v7)
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitProcess_1), int32(455), int32(_a_F_InitProcess_2))
											mBase = m.M
											v304 = m.ExcPending
											if v304 != 0 {
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
										v262 = m.ExcPending
										if v262 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitProcess_1), int32(458), int32(_a_F_InitProcess_2))
											mBase = m.M
											v267 = m.ExcPending
											if v267 != 0 {
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
				v20 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
				v22 = v20 - int32(3)
				if base.Ui32(v22) <= base.Ui32(int32(4)) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v22<<(uint(int32(2))%32))+uint32(_c_F_InitProcess[4])))
					v29 = v27
				} else {
					v29 = int32(24)
				}
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
				v32 = v29 + v31
				v35 = base.AtomicRmwXchg32(m, v31, int32(20), int32(1))
				if v35 != 0 {
					F_s_lock(m, v31+int32(20), int32(_a_F_InitProcess_3))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
						*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5])) = v43
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
						v47 = int32(0)
						if base.B2i32(v46 == v47)|base.B2i32(v46 == v32) == v47 {
							v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
							*(*int32)(unsafe.Add(mBase, uint32(v54))) = v56
							v58 = int32(_a_F_InitProcess_4)
							v59 = int32(4)
							v60 = v46 - v59
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1])) = v60
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
							v64 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v63)+20)), uint32(v64))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
							v71 = base.I32_div_s(v60-v68, int32(768))
							*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[6])) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v46)+572)) = v64
							*(*uint8)(unsafe.Add(mBase, uint32(v46)+568)) = uint8(v64)
							*(*int32)(unsafe.Add(mBase, uint32(v46)+412)) = v64
							v79 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v46))) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v46)+44)) = v79
							v84 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7]))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v46)+36)) = v71
							*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v84
							*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v64
							*(*int64)(unsafe.Add(mBase, uint32(v46)+16)) = v79
							v94 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+332)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v46)+380)) = v79
							*(*uint16)(unsafe.Add(mBase, uint32(v46)+340)) = uint16(v64)
							*(*int64)(unsafe.Add(mBase, uint32(v46)+388)) = v79
							*(*uint8)(unsafe.Add(mBase, uint32(v46)+32)) = uint8(base.B2i32(v94 == v59))
							v109 = base.AtomicRmwXchg64(m, v46, int32(404), v79)
							v111 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
							*(*int32)(unsafe.Add(mBase, uint32(v111)+648)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v111)+612)) = v64
							*(*int32)(unsafe.Add(mBase, uint32(v111)+340)) = v64
							*(*int64)(unsafe.Add(mBase, uint32(v111)+624)) = v79
							*(*uint8)(unsafe.Add(mBase, uint32(v111)+616)) = uint8(v64)
							*(*int64)(unsafe.Add(mBase, uint32(v111)+584)) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v111)+592)) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v111)+597)) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v111)+640)) = v79
							*(*int64)(unsafe.Add(mBase, uint32(v111)+632)) = int64(-1)
							F_OwnLatch(m, v111+int32(316))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return
							} else {
								F_SwitchToSharedLatch(m)
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return
								} else {
									v139 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
									*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[8])) = v139 + int32(648)
									v145 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+332))
									F_PGSemaphoreReset(m, v146)
									mBase = m.M
									F_on_shmem_exit(m, int32(1227), int64(0))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
										return
									} else {
										v152 = int32(_a_F_InitProcess_5)
										v153 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9]))
										v156 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10]))
										*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v156
										v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
										v163 = F_palloc(m, v160<<(uint(int32(2))%32))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12])) = v163
											v168 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
											v171 = F_palloc(m, v168*int32(24))
											mBase = m.M
											v172 = m.ExcPending
											if v172 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13])) = v171
												v176 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[14])) = v176
												v180 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
												v183 = F_palloc(m, v180<<(uint(int32(2))%32))
												mBase = m.M
												v184 = m.ExcPending
												if v184 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[15])) = v183
													v188 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
													v191 = F_palloc(m, v188<<(uint(int32(2))%32))
													mBase = m.M
													v192 = m.ExcPending
													if v192 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[16])) = v191
														v196 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
														v198 = base.I32_div_s(v196, int32(2))
														v201 = F_palloc(m, v198*int32(12))
														mBase = m.M
														v202 = m.ExcPending
														if v202 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[17])) = v201
															v206 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
															v209 = F_palloc(m, v206<<(uint(int32(2))%32))
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[18])) = v209
																v214 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[19])) = v214
																v219 = F_palloc(m, v214*int32(20))
																mBase = m.M
																v220 = m.ExcPending
																if v220 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[20])) = v219
																	v224 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[21])) = v224 << (uint(int32(2)) % 32)
																	v230 = F_palloc(m, v224*int32(80))
																	mBase = m.M
																	v231 = m.ExcPending
																	if v231 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v153
																		*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[22])) = v230
																		v237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
																		if v237 != 0 {
																			F_AttachSharedMemoryStructs(m)
																			mBase = m.M
																			v239 = m.ExcPending
																			if v239 != 0 {
																				return
																			} else {
																				m.G0 = v7 + int32(16)
																				return
																			}
																		} else {
																			m.G0 = v7 + int32(16)
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
							v244 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
							v245 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v244)+20)), uint32(v245))
							v249 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
							F_errstart_cold(m, int32(22), v245)
							mBase = m.M
							v253 = m.ExcPending
							if v253 != 0 {
								return
							} else {
								F_errcode(m, int32(_a_F_InitProcess_6))
								mBase = m.M
								v256 = m.ExcPending
								if v256 != 0 {
									return
								} else {
									if v249 == int32(6) {
										v295 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[23]))
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v295
										F_errmsg(m, int32(_a_F_InitProcess_7), v7)
										mBase = m.M
										v299 = m.ExcPending
										if v299 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitProcess_1), int32(455), int32(_a_F_InitProcess_2))
											mBase = m.M
											v304 = m.ExcPending
											if v304 != 0 {
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
										v262 = m.ExcPending
										if v262 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_InitProcess_1), int32(458), int32(_a_F_InitProcess_2))
											mBase = m.M
											v267 = m.ExcPending
											if v267 != 0 {
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
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
					*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[5])) = v43
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
					v47 = int32(0)
					if base.B2i32(v46 == v47)|base.B2i32(v46 == v32) == v47 {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
						v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v54
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
						*(*int32)(unsafe.Add(mBase, uint32(v54))) = v56
						v58 = int32(_a_F_InitProcess_4)
						v59 = int32(4)
						v60 = v46 - v59
						*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1])) = v60
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
						v64 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v63)+20)), uint32(v64))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
						v71 = base.I32_div_s(v60-v68, int32(768))
						*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[6])) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v46)+572)) = v64
						*(*uint8)(unsafe.Add(mBase, uint32(v46)+568)) = uint8(v64)
						*(*int32)(unsafe.Add(mBase, uint32(v46)+412)) = v64
						v79 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v46))) = v79
						*(*int64)(unsafe.Add(mBase, uint32(v46)+44)) = v79
						v84 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[7]))
						*(*int32)(unsafe.Add(mBase, uint32(v46)+40)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v46)+36)) = v71
						*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v84
						*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = v64
						*(*int64)(unsafe.Add(mBase, uint32(v46)+16)) = v79
						v94 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v46)+332)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v94
						*(*int64)(unsafe.Add(mBase, uint32(v46)+380)) = v79
						*(*uint16)(unsafe.Add(mBase, uint32(v46)+340)) = uint16(v64)
						*(*int64)(unsafe.Add(mBase, uint32(v46)+388)) = v79
						*(*uint8)(unsafe.Add(mBase, uint32(v46)+32)) = uint8(base.B2i32(v94 == v59))
						v109 = base.AtomicRmwXchg64(m, v46, int32(404), v79)
						v111 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
						*(*int32)(unsafe.Add(mBase, uint32(v111)+648)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v111)+612)) = v64
						*(*int32)(unsafe.Add(mBase, uint32(v111)+340)) = v64
						*(*int64)(unsafe.Add(mBase, uint32(v111)+624)) = v79
						*(*uint8)(unsafe.Add(mBase, uint32(v111)+616)) = uint8(v64)
						*(*int64)(unsafe.Add(mBase, uint32(v111)+584)) = v79
						*(*int64)(unsafe.Add(mBase, uint32(v111)+592)) = v79
						*(*int64)(unsafe.Add(mBase, uint32(v111)+597)) = v79
						*(*int64)(unsafe.Add(mBase, uint32(v111)+640)) = v79
						*(*int64)(unsafe.Add(mBase, uint32(v111)+632)) = int64(-1)
						F_OwnLatch(m, v111+int32(316))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							F_SwitchToSharedLatch(m)
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								v139 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
								*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[8])) = v139 + int32(648)
								v145 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[1]))
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+332))
								F_PGSemaphoreReset(m, v146)
								mBase = m.M
								F_on_shmem_exit(m, int32(1227), int64(0))
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return
								} else {
									v152 = int32(_a_F_InitProcess_5)
									v153 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9]))
									v156 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[10]))
									*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v156
									v160 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
									v163 = F_palloc(m, v160<<(uint(int32(2))%32))
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12])) = v163
										v168 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
										v171 = F_palloc(m, v168*int32(24))
										mBase = m.M
										v172 = m.ExcPending
										if v172 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[13])) = v171
											v176 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[12]))
											*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[14])) = v176
											v180 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
											v183 = F_palloc(m, v180<<(uint(int32(2))%32))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[15])) = v183
												v188 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
												v191 = F_palloc(m, v188<<(uint(int32(2))%32))
												mBase = m.M
												v192 = m.ExcPending
												if v192 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[16])) = v191
													v196 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
													v198 = base.I32_div_s(v196, int32(2))
													v201 = F_palloc(m, v198*int32(12))
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[17])) = v201
														v206 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
														v209 = F_palloc(m, v206<<(uint(int32(2))%32))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[18])) = v209
															v214 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
															*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[19])) = v214
															v219 = F_palloc(m, v214*int32(20))
															mBase = m.M
															v220 = m.ExcPending
															if v220 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[20])) = v219
																v224 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[11]))
																*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[21])) = v224 << (uint(int32(2)) % 32)
																v230 = F_palloc(m, v224*int32(80))
																mBase = m.M
																v231 = m.ExcPending
																if v231 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[9])) = v153
																	*(*int32)(unsafe.Add(mBase, _c_F_InitProcess[22])) = v230
																	v237 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitProcess[2])))
																	if v237 != 0 {
																		F_AttachSharedMemoryStructs(m)
																		mBase = m.M
																		v239 = m.ExcPending
																		if v239 != 0 {
																			return
																		} else {
																			m.G0 = v7 + int32(16)
																			return
																		}
																	} else {
																		m.G0 = v7 + int32(16)
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
						v244 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[0]))
						v245 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v244)+20)), uint32(v245))
						v249 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[3]))
						F_errstart_cold(m, int32(22), v245)
						mBase = m.M
						v253 = m.ExcPending
						if v253 != 0 {
							return
						} else {
							F_errcode(m, int32(_a_F_InitProcess_6))
							mBase = m.M
							v256 = m.ExcPending
							if v256 != 0 {
								return
							} else {
								if v249 == int32(6) {
									v295 = *(*int32)(unsafe.Add(mBase, _c_F_InitProcess[23]))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v295
									F_errmsg(m, int32(_a_F_InitProcess_7), v7)
									mBase = m.M
									v299 = m.ExcPending
									if v299 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_InitProcess_1), int32(455), int32(_a_F_InitProcess_2))
										mBase = m.M
										v304 = m.ExcPending
										if v304 != 0 {
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
									v262 = m.ExcPending
									if v262 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_InitProcess_1), int32(458), int32(_a_F_InitProcess_2))
										mBase = m.M
										v267 = m.ExcPending
										if v267 != 0 {
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
		F_errstart_cold(m, int32(24), int32(0))
		mBase = m.M
		v271 = m.ExcPending
		if v271 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_InitProcess_9), int32(0))
			mBase = m.M
			v275 = m.ExcPending
			if v275 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_InitProcess_1), int32(400), int32(_a_F_InitProcess_2))
				mBase = m.M
				v280 = m.ExcPending
				if v280 != 0 {
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
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	v3 = int32(10)
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0])) = v3
	*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[1])) = v3
	v11 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])) = uint8(v11)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])))
	if v14 != 0 {
		v16 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v16)
		v43 = v3
		v44 = int32(_a_F_InitializeSessionUserIdStandalone_0)
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v43
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v43
		F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_1), v44, int32(0), int32(1))
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return
		} else {
			return
		}
	} else {
		v20 = int32(10)
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v20
		*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v20
		F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_1), int32(_a_F_InitializeSessionUserIdStandalone_0), int32(0), int32(1))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			v32 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[3])) = uint8(v32)
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[0]))
			if v35 == v32 {
				return
			} else {
				v41 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[2])))
				if v41 != 0 {
					v42 = int32(_a_F_InitializeSessionUserIdStandalone_0)
				} else {
					v42 = int32(_a_F_InitializeSessionUserIdStandalone_2)
				}
				v43 = v35
				v44 = v42
				*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[4])) = v43
				*(*int32)(unsafe.Add(mBase, _c_F_InitializeSessionUserIdStandalone[5])) = v43
				F_SetConfigOption(m, int32(_a_F_InitializeSessionUserIdStandalone_1), v44, int32(0), int32(1))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_InvalidationCallback(m *base.Module, l0 int64, l1 int32, l2 int32) {
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	v4 = l0
	goto L4
L1:
	;
	return v241
L2:
	;
	v241 = int32(1)
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
		v241 = int32(0)
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
		v241 = int32(1)
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
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v139 < int32(2) {
		goto L2
	} else {
		goto L56
	}
L19:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_IsSquashableConstant[0]))
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_IsSquashableConstant[1]))
	v49 = m.G0
	v50 = v48 - v49
	v52 = v50 >> (uint(int32(31)) % 32)
	goto L20
L20:
	;
	if base.B2i32(v46 < v50^v52-v52)&base.B2i32(v48 != int32(0)) != 0 {
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
	v63 = v39
	goto L28
L24:
	;
	if v135 != 0 {
		goto L18
	} else {
		goto L55
	}
L25:
	;
	v135 = v131
	goto L24
L26:
	;
	v131 = int32(1)
	goto L25
L27:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v78 = int32(1)
	if base.Ui32(v78) < base.Ui32(v77-v78) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if base.Ui32(int32(2)) <= base.Ui32(v66-int32(27)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v135 = base.B2i32(v74 == int32(0))
	goto L24
L30:
	;
	switch v66 - int32(7) {
	case 0:
		goto L26
	case 1:
		goto L33
	default:
		v131 = int32(0)
		goto L25
	case 8:
		goto L27
	}
L31:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v63 = v73
	goto L28
L32:
	;
	goto L29
L33:
	;
	goto L32
L34:
	;
	v135 = int32(0)
	goto L24
L35:
	;
	goto L36
L36:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if base.Ui32(int32(_a_F_IsSquashableConstant_0)) < base.Ui32(v83) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v135 = int32(0)
	goto L24
L38:
	;
	goto L39
L39:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
	if v87 == int32(0) {
		goto L26
	} else {
		goto L40
	}
L40:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v91 <= int32(0) {
		v131 = int32(1)
		goto L25
	} else {
		goto L41
	}
L41:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v96 == int32(7) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v104 < int32(2) {
		goto L26
	} else {
		goto L48
	}
L43:
	;
	v99 = F_stack_is_too_deep(m)
	mBase = m.M
	if v99 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v135 = int32(0)
	goto L24
L45:
	;
	goto L46
L46:
	;
	v101 = F_IsSquashableConstant(m, v95)
	mBase = m.M
	if v101 != 0 {
		goto L42
	} else {
		goto L47
	}
L47:
	;
	v135 = int32(0)
	goto L24
L48:
	;
	v107 = int32(1)
	goto L49
L49:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v110+v107<<(uint(int32(2))%32))))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v115 == int32(7) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v131 = v120
	goto L25
L51:
	;
	v120 = int32(1)
	v122 = v107 + v120
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v122 < v123 {
		v107 = v122
		goto L49
	} else {
		goto L54
	}
L52:
	;
	v118 = F_IsSquashableConstant(m, v114)
	mBase = m.M
	if v118 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v135 = int32(0)
	goto L24
L54:
	;
	goto L50
L55:
	;
	return int32(0)
L56:
	;
	v142 = int32(1)
	goto L57
L57:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v145+v142<<(uint(int32(2))%32))))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	if v150 == int32(7) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v241 = v230
	goto L1
L59:
	;
	v230 = int32(1)
	v232 = v142 + v230
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v232 < v233 {
		v142 = v232
		goto L57
	} else {
		goto L93
	}
L60:
	;
	v155 = v149
	goto L65
L61:
	;
	if v227 != 0 {
		goto L59
	} else {
		goto L92
	}
L62:
	;
	v227 = v223
	goto L61
L63:
	;
	v223 = int32(1)
	goto L62
L64:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v155)+16))
	v170 = int32(1)
	if base.Ui32(v170) < base.Ui32(v169-v170) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	if base.Ui32(int32(2)) <= base.Ui32(v158-int32(27)) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v227 = base.B2i32(v166 == int32(0))
	goto L61
L67:
	;
	switch v158 - int32(7) {
	case 0:
		goto L63
	case 1:
		goto L70
	default:
		v223 = int32(0)
		goto L62
	case 8:
		goto L64
	}
L68:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v155 = v165
	goto L65
L69:
	;
	goto L66
L70:
	;
	goto L69
L71:
	;
	v227 = int32(0)
	goto L61
L72:
	;
	goto L73
L73:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	if base.Ui32(int32(_a_F_IsSquashableConstant_0)) < base.Ui32(v175) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v227 = int32(0)
	goto L61
L75:
	;
	goto L76
L76:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	if v179 == int32(0) {
		goto L63
	} else {
		goto L77
	}
L77:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v183 <= int32(0) {
		v223 = int32(1)
		goto L62
	} else {
		goto L78
	}
L78:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v188 == int32(7) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v196 < int32(2) {
		goto L63
	} else {
		goto L85
	}
L80:
	;
	v191 = F_stack_is_too_deep(m)
	mBase = m.M
	if v191 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v227 = int32(0)
	goto L61
L82:
	;
	goto L83
L83:
	;
	v193 = F_IsSquashableConstant(m, v187)
	mBase = m.M
	if v193 != 0 {
		goto L79
	} else {
		goto L84
	}
L84:
	;
	v227 = int32(0)
	goto L61
L85:
	;
	v199 = int32(1)
	goto L86
L86:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v179)+12))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202+v199<<(uint(int32(2))%32))))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)))
	if v207 == int32(7) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v223 = v212
	goto L62
L88:
	;
	v212 = int32(1)
	v214 = v199 + v212
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v214 < v215 {
		v199 = v214
		goto L86
	} else {
		goto L91
	}
L89:
	;
	v210 = F_IsSquashableConstant(m, v206)
	mBase = m.M
	if v210 != 0 {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v227 = int32(0)
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = base.I64_extend_i32_u(l0)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_IsThereCollationInNamespace[0]))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v17 = base.I64_extend_i32_u(l1)
	v19 = F_SearchSysCacheExists(m, int32(15), v12, base.I64_extend_i32_s(v15), v17, int64(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		if v19 == int32(0) {
			v26 = F_SearchSysCacheExists(m, int32(15), v12, int64(-1), v17, int64(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				if v26 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						F_errcode(m, int32(_a_F_IsThereCollationInNamespace_0))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							v61 = F_get_namespace_name(m, l1)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v61
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l0
								F_errmsg(m, int32(_a_F_IsThereCollationInNamespace_1), v9+int32(16))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_IsThereCollationInNamespace_2), int32(422), int32(_a_F_IsThereCollationInNamespace_3))
									mBase = m.M
									v74 = m.ExcPending
									if v74 != 0 {
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
					m.G0 = v9 + int32(32)
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				F_errcode(m, int32(_a_F_IsThereCollationInNamespace_0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, _c_F_IsThereCollationInNamespace[0]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
					v41 = F_get_namespace_name(m, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v40
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
						F_errmsg(m, int32(_a_F_IsThereCollationInNamespace_4), v9)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_IsThereCollationInNamespace_2), int32(412), int32(_a_F_IsThereCollationInNamespace_3))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
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
	Fn14218(m, l0, l1, l2, int32(_a_F_IsThereOpClassInNamespace_0), int32(1861), int32(_a_F_IsThereOpClassInNamespace_1), int32(13))
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
	Fn14218(m, l0, l1, l2, int32(_a_F_IsThereOpFamilyInNamespace_0), int32(1884), int32(_a_F_IsThereOpFamilyInNamespace_1), int32(41))
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_icregexeqsel(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14370(m, l0, int32(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_ind_fetch_func(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
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
	var v14 int64
	_ = v14
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+244))
	v5 = v4 * l1
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+240))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5+v6))))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v8)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v10+v5<<(uint(int32(3))%32))))
	return v14
}
func F_indonesian_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
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
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v482 int32
	_ = v482
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v618 int32
	_ = v618
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v647 int32
	_ = v647
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v798 int32
	_ = v798
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v836 int32
	_ = v836
	var v843 int32
	_ = v843
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v916 int32
	_ = v916
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v935 int32
	_ = v935
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v961 int32
	_ = v961
	var v973 int32
	_ = v973
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1054 int32
	_ = v1054
	var v1060 int32
	_ = v1060
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v2
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = v9
	goto L3
L1:
	;
	if int32(0) <= v126 {
		goto L26
	} else {
		goto L27
	}
L2:
	;
	v126 = v98
	goto L1
L3:
	;
	if v22 <= v31 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v126 = int32(-1)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v38 = int32(1)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31+v23))))
	if base.Ui32(v40) < base.Ui32(int32(192)) {
		v97 = v40
		v98 = v38
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if int32(117) < v97 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	v44 = v31 + int32(1)
	if v44 == v22 {
		v97 = v40
		v98 = v38
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44+v23))))
	v49 = v47 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v40) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v23))))
	v65 = v63 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v40) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v53 = v31 + int32(2)
	if v53 != v22 {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v97 = v40<<(uint(int32(6))%32)&int32(1984) | v49
	v98 = int32(2)
	goto L8
L15:
	;
	goto L14
L16:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v69))))
	v97 = v82&int32(63) | (v40<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v49<<(uint(int32(12))%32) | v65<<(uint(int32(6))%32))
	v98 = int32(4)
	goto L8
L17:
	;
	v69 = v31 + int32(3)
	if v69 != v22 {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v97 = v40<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v49<<(uint(int32(6))%32) | v65
	v98 = int32(3)
	goto L8
L20:
	;
	goto L19
L21:
	;
	v115 = v98 + v31
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v115
	v31 = v115
	goto L3
L22:
	;
	v102 = v97 - int32(97)
	if v102 < int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v102)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v108)>>(uint(v102&int32(7))%32))&int32(1) != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	goto L21
L26:
	;
	v132 = v126
	goto L29
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v268 < int32(3) {
		v1060 = v2
		goto L57
	} else {
		goto L58
	}
L29:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v136 = v135 + v132
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v136
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v138 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v163 = v136
	goto L33
L30:
	;
	goto L28
L31:
	;
	if int32(0) <= v258 {
		v132 = v258
		goto L29
	} else {
		goto L56
	}
L32:
	;
	v258 = v230
	goto L31
L33:
	;
	if v154 <= v163 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v258 = int32(-1)
	goto L31
L36:
	;
	goto L37
L37:
	;
	v170 = int32(1)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v155))))
	if base.Ui32(v172) < base.Ui32(int32(192)) {
		v229 = v172
		v230 = v170
		goto L38
	} else {
		goto L39
	}
L38:
	;
	if int32(117) < v229 {
		goto L51
	} else {
		goto L52
	}
L39:
	;
	v176 = v163 + int32(1)
	if v176 == v154 {
		v229 = v172
		v230 = v170
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v155))))
	v181 = v179 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v172) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185+v155))))
	v197 = v195 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v172) {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	v185 = v163 + int32(2)
	if v185 != v154 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v229 = v172<<(uint(int32(6))%32)&int32(1984) | v181
	v230 = int32(2)
	goto L38
L45:
	;
	goto L44
L46:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155+v201))))
	v229 = v214&int32(63) | (v172<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v181<<(uint(int32(12))%32) | v197<<(uint(int32(6))%32))
	v230 = int32(4)
	goto L38
L47:
	;
	v201 = v163 + int32(3)
	if v201 != v154 {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v229 = v172<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v181<<(uint(int32(6))%32) | v197
	v230 = int32(3)
	goto L38
L50:
	;
	goto L49
L51:
	;
	v247 = v230 + v163
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v247
	v163 = v247
	goto L33
L52:
	;
	v234 = v229 - int32(97)
	if v234 < int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v234)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v240)>>(uint(v234&int32(7))%32))&int32(1) != 0 {
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
	return v1060
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v9
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(0)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v274
	if v274-int32(2) <= v9 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v320
	v324 = v320 - int32(1)
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v324 <= v325 {
		goto L72
	} else {
		goto L73
	}
L60:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v316
	v320 = v316
	v321 = v2
	goto L59
L61:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280+v274-int32(1)))))
	switch v284 - int32(104) {
	case 0, 6:
		goto L62
	default:
		goto L60
	}
L62:
	;
	v290 = F_find_among_b(m, l0, int32(_a_F_indonesian_UTF_8_stem_2), int32(3), int32(0))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v310
	if int32(3) <= v308 {
		v320 = v310
		v321 = v309
		goto L59
	} else {
		goto L70
	}
L64:
	;
	return int32(0)
L65:
	;
	if v290 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v308 = v296
	v309 = v2
	goto L63
L67:
	;
	goto L68
L68:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v297
	v299 = F_slice_del(m, l0)
	mBase = m.M
	if v299 < int32(0) {
		v1060 = v299
		goto L57
	} else {
		goto L69
	}
L69:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v303 = int32(1)
	v304 = v302 - v303
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v304
	v308 = v304
	v309 = v303
	goto L63
L70:
	;
	return int32(0)
L71:
	;
	v366 = int32(0)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v367
	v370 = v367 + int32(1)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v371 <= v370 {
		v997 = v366
		goto L84
	} else {
		goto L85
	}
L72:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v361
	v365 = v361
	goto L71
L73:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v327+v324))))
	if base.B2i32(v329 != int32(117))&base.B2i32(v329 != int32(97)) != 0 {
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v338 = F_find_among_b(m, l0, int32(_a_F_indonesian_UTF_8_stem_3), int32(3), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L64
	} else {
		goto L76
	}
L75:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v354
	if int32(3) <= v353 {
		v365 = v354
		goto L71
	} else {
		goto L81
	}
L76:
	;
	if v338 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v353 = v342
	goto L75
L78:
	;
	goto L79
L79:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v343
	v345 = F_slice_del(m, l0)
	mBase = m.M
	if v345 < int32(0) {
		v1060 = v345
		goto L57
	} else {
		goto L80
	}
L80:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v350 = v348 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v350
	v353 = v350
	goto L75
L81:
	;
	return int32(0)
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365
	v1060 = int32(1)
	goto L57
L83:
	;
	if v1003 != 0 {
		goto L221
	} else {
		goto L222
	}
L84:
	;
	v1003 = v997
	goto L83
L85:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v370))))
	switch v375 - int32(101) {
	case 0, 4:
		goto L86
	default:
		v997 = v366
		goto L84
	}
L86:
	;
	v381 = F_find_among(m, l0, int32(_a_F_indonesian_UTF_8_stem_4), int32(10), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L64
	} else {
		goto L87
	}
L87:
	;
	if v381 == int32(0) {
		v997 = v366
		goto L84
	} else {
		goto L88
	}
L88:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v385
	v387 = int32(1)
	switch v381 - v387 {
	case 0:
		goto L94
	case 1:
		goto L93
	case 2:
		goto L92
	case 3:
		goto L91
	case 4:
		goto L90
	case 5:
		goto L89
	default:
		v997 = v387
		goto L84
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v859 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v859 - int32(1)
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v876 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L192
L90:
	;
	v720 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v720
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v722 - v720
	v737 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v738 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L161
L91:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v385 == v565 {
		goto L128
	} else {
		goto L129
	}
L92:
	;
	v555 = F_slice_del(m, l0)
	mBase = m.M
	if v555 < int32(0) {
		v997 = v555
		goto L84
	} else {
		goto L127
	}
L93:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v385 == v400 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	v390 = F_slice_del(m, l0)
	mBase = m.M
	if v390 < int32(0) {
		v997 = v390
		goto L84
	} else {
		goto L95
	}
L95:
	;
	v393 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v393
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v395 - v393
	v1003 = v393
	goto L83
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385
	v545 = F_slice_del(m, l0)
	mBase = m.M
	if v545 < int32(0) {
		v997 = v545
		goto L84
	} else {
		goto L126
	}
L97:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v385))))
	if v404 != int32(121) {
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v408 = v385 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v408
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L101
L99:
	;
	if v527 != 0 {
		goto L96
	} else {
		goto L123
	}
L100:
	;
	v527 = v520
	goto L99
L101:
	;
	if v422 <= v408 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	v520 = int32(0)
	goto L100
L103:
	;
	v527 = int32(-1)
	goto L99
L104:
	;
	goto L105
L105:
	;
	v438 = int32(1)
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v408+v423))))
	if base.Ui32(v440) < base.Ui32(int32(192)) {
		v497 = v440
		v498 = v438
		goto L106
	} else {
		goto L107
	}
L106:
	;
	if int32(117) < v497 {
		v520 = v498
		goto L100
	} else {
		goto L119
	}
L107:
	;
	v444 = v385 + int32(2)
	if v444 == v422 {
		v497 = v440
		v498 = v438
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444+v423))))
	v449 = v447 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v440) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453+v423))))
	v465 = v463 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v440) {
		goto L115
	} else {
		goto L116
	}
L110:
	;
	v453 = v385 + int32(3)
	if v453 != v422 {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v497 = v440<<(uint(int32(6))%32)&int32(1984) | v449
	v498 = int32(2)
	goto L106
L113:
	;
	goto L112
L114:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423+v469))))
	v497 = v482&int32(63) | (v440<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v449<<(uint(int32(12))%32) | v465<<(uint(int32(6))%32))
	v498 = int32(4)
	goto L106
L115:
	;
	v469 = v385 + int32(4)
	if v469 != v422 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v497 = v440<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v449<<(uint(int32(6))%32) | v465
	v498 = int32(3)
	goto L106
L118:
	;
	goto L117
L119:
	;
	v502 = v497 - int32(97)
	if v502 < int32(0) {
		v520 = v498
		goto L100
	} else {
		goto L120
	}
L120:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v502)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v508)>>(uint(v502&int32(7))%32))&int32(1) == int32(0) {
		v520 = v498
		goto L100
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v498 + v408
	goto L122
L122:
	;
	goto L102
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v408
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v408
	v532 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_UTF_8_stem_5))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L64
	} else {
		goto L124
	}
L124:
	;
	if v532 < int32(0) {
		v997 = v532
		goto L84
	} else {
		goto L125
	}
L125:
	;
	v536 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v536
	v538 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v538 - v536
	v1003 = v536
	goto L83
L126:
	;
	v548 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v548
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v550 - v548
	v1003 = v548
	goto L83
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v561 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v560 - v561
	v1003 = v561
	goto L83
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385
	v710 = F_slice_del(m, l0)
	mBase = m.M
	if v710 < int32(0) {
		v997 = v710
		goto L84
	} else {
		goto L158
	}
L129:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v567+v385))))
	if v569 != int32(121) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v573 = v385 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v573
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L133
L131:
	;
	if v692 != 0 {
		goto L128
	} else {
		goto L155
	}
L132:
	;
	v692 = v685
	goto L131
L133:
	;
	if v587 <= v573 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v685 = int32(0)
	goto L132
L135:
	;
	v692 = int32(-1)
	goto L131
L136:
	;
	goto L137
L137:
	;
	v603 = int32(1)
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v573+v588))))
	if base.Ui32(v605) < base.Ui32(int32(192)) {
		v662 = v605
		v663 = v603
		goto L138
	} else {
		goto L139
	}
L138:
	;
	if int32(117) < v662 {
		v685 = v663
		goto L132
	} else {
		goto L151
	}
L139:
	;
	v609 = v385 + int32(2)
	if v609 == v587 {
		v662 = v605
		v663 = v603
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v612 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609+v588))))
	v614 = v612 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v605) {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v618+v588))))
	v630 = v628 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v605) {
		goto L147
	} else {
		goto L148
	}
L142:
	;
	v618 = v385 + int32(3)
	if v618 != v587 {
		goto L141
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	v662 = v605<<(uint(int32(6))%32)&int32(1984) | v614
	v663 = int32(2)
	goto L138
L145:
	;
	goto L144
L146:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v588+v634))))
	v662 = v647&int32(63) | (v605<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v614<<(uint(int32(12))%32) | v630<<(uint(int32(6))%32))
	v663 = int32(4)
	goto L138
L147:
	;
	v634 = v385 + int32(4)
	if v634 != v587 {
		goto L146
	} else {
		goto L150
	}
L148:
	;
	goto L149
L149:
	;
	v662 = v605<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v614<<(uint(int32(6))%32) | v630
	v663 = int32(3)
	goto L138
L150:
	;
	goto L149
L151:
	;
	v667 = v662 - int32(97)
	if v667 < int32(0) {
		v685 = v663
		goto L132
	} else {
		goto L152
	}
L152:
	;
	v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v667)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v673)>>(uint(v667&int32(7))%32))&int32(1) == int32(0) {
		v685 = v663
		goto L132
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v663 + v573
	goto L154
L154:
	;
	goto L134
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v573
	v697 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_UTF_8_stem_6))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L64
	} else {
		goto L156
	}
L156:
	;
	if v697 < int32(0) {
		v997 = v697
		goto L84
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v704 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v703 - v704
	v1003 = v704
	goto L83
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v716 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v715 - v716
	v1003 = v716
	goto L83
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385
	if v843 == int32(0) {
		goto L184
	} else {
		goto L185
	}
L160:
	;
	v843 = v836
	goto L159
L161:
	;
	if v738 <= v737 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v836 = int32(0)
	goto L160
L163:
	;
	v843 = int32(-1)
	goto L159
L164:
	;
	goto L165
L165:
	;
	v754 = int32(1)
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v737+v739))))
	if base.Ui32(v756) < base.Ui32(int32(192)) {
		v813 = v756
		v814 = v754
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if int32(117) < v813 {
		v836 = v814
		goto L160
	} else {
		goto L179
	}
L167:
	;
	v760 = v737 + int32(1)
	if v760 == v738 {
		v813 = v756
		v814 = v754
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v760+v739))))
	v765 = v763 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v756) {
		goto L170
	} else {
		goto L171
	}
L169:
	;
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v769+v739))))
	v781 = v779 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v756) {
		goto L175
	} else {
		goto L176
	}
L170:
	;
	v769 = v737 + int32(2)
	if v769 != v738 {
		goto L169
	} else {
		goto L173
	}
L171:
	;
	goto L172
L172:
	;
	v813 = v756<<(uint(int32(6))%32)&int32(1984) | v765
	v814 = int32(2)
	goto L166
L173:
	;
	goto L172
L174:
	;
	v798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739+v785))))
	v813 = v798&int32(63) | (v756<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v765<<(uint(int32(12))%32) | v781<<(uint(int32(6))%32))
	v814 = int32(4)
	goto L166
L175:
	;
	v785 = v737 + int32(3)
	if v785 != v738 {
		goto L174
	} else {
		goto L178
	}
L176:
	;
	goto L177
L177:
	;
	v813 = v756<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v765<<(uint(int32(6))%32) | v781
	v814 = int32(3)
	goto L166
L178:
	;
	goto L177
L179:
	;
	v818 = v813 - int32(97)
	if v818 < int32(0) {
		v836 = v814
		goto L160
	} else {
		goto L180
	}
L180:
	;
	v824 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v818)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v824)>>(uint(v818&int32(7))%32))&int32(1) == int32(0) {
		v836 = v814
		goto L160
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v814 + v737
	goto L182
L182:
	;
	goto L162
L183:
	;
	v1003 = v856
	goto L83
L184:
	;
	v849 = F_slice_from_s(m, l0, int32(1), int32(_a_F_indonesian_UTF_8_stem_7))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L64
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v853 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v853 {
		v997 = v387
		goto L84
	} else {
		goto L189
	}
L187:
	;
	if v849 < int32(0) {
		v856 = v849
		goto L183
	} else {
		goto L188
	}
L188:
	;
	v997 = v387
	goto L84
L189:
	;
	v856 = v853
	goto L183
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385
	if v980 == int32(0) {
		goto L215
	} else {
		goto L216
	}
L191:
	;
	v980 = v973
	goto L190
L192:
	;
	if v875 <= v874 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v973 = int32(0)
	goto L191
L194:
	;
	v980 = int32(-1)
	goto L190
L195:
	;
	goto L196
L196:
	;
	v891 = int32(1)
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v874+v876))))
	if base.Ui32(v893) < base.Ui32(int32(192)) {
		v950 = v893
		v951 = v891
		goto L197
	} else {
		goto L198
	}
L197:
	;
	if int32(117) < v950 {
		v973 = v951
		goto L191
	} else {
		goto L210
	}
L198:
	;
	v897 = v874 + int32(1)
	if v897 == v875 {
		v950 = v893
		v951 = v891
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v900 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v897+v876))))
	v902 = v900 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v893) {
		goto L201
	} else {
		goto L202
	}
L200:
	;
	v916 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906+v876))))
	v918 = v916 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v893) {
		goto L206
	} else {
		goto L207
	}
L201:
	;
	v906 = v874 + int32(2)
	if v906 != v875 {
		goto L200
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v950 = v893<<(uint(int32(6))%32)&int32(1984) | v902
	v951 = int32(2)
	goto L197
L204:
	;
	goto L203
L205:
	;
	v935 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876+v922))))
	v950 = v935&int32(63) | (v893<<(uint(int32(18))%32)&int32(_a_F_indonesian_UTF_8_stem_0) | v902<<(uint(int32(12))%32) | v918<<(uint(int32(6))%32))
	v951 = int32(4)
	goto L197
L206:
	;
	v922 = v874 + int32(3)
	if v922 != v875 {
		goto L205
	} else {
		goto L209
	}
L207:
	;
	goto L208
L208:
	;
	v950 = v893<<(uint(int32(12))%32)&int32(_a_F_indonesian_UTF_8_stem_1) | v902<<(uint(int32(6))%32) | v918
	v951 = int32(3)
	goto L197
L209:
	;
	goto L208
L210:
	;
	v955 = v950 - int32(97)
	if v955 < int32(0) {
		v973 = v951
		goto L191
	} else {
		goto L211
	}
L211:
	;
	v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v955)>>(uint(int32(3))%32)))+uint32(_c_F_indonesian_UTF_8_stem[0]))))
	if int32(base.Ui32(v961)>>(uint(v955&int32(7))%32))&int32(1) == int32(0) {
		v973 = v951
		goto L191
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v951 + v874
	goto L213
L213:
	;
	goto L193
L214:
	;
	v997 = v996
	goto L84
L215:
	;
	v984 = int32(1)
	v987 = F_slice_from_s(m, l0, v984, int32(_a_F_indonesian_UTF_8_stem_8))
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L64
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v992 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v992 {
		v997 = int32(1)
		goto L84
	} else {
		goto L220
	}
L218:
	;
	if v987 < int32(0) {
		v996 = v987
		goto L214
	} else {
		goto L219
	}
L219:
	;
	v997 = v984
	goto L84
L220:
	;
	v996 = v992
	goto L214
L221:
	;
	if v1003 < int32(0) {
		v1060 = v1003
		goto L57
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v365
	v1033 = F_r_remove_second_order_prefix_2(m, l0)
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L64
	} else {
		goto L238
	}
L224:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1006 < int32(3) {
		goto L82
	} else {
		goto L225
	}
L225:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1009
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1011
	v1013 = F_r_remove_suffix_2(m, l0)
	mBase = m.M
	v1014 = m.ExcPending
	if v1014 != 0 {
		goto L64
	} else {
		goto L226
	}
L226:
	;
	if v1013 == int32(0) {
		goto L82
	} else {
		goto L227
	}
L227:
	;
	if v1013 < int32(0) {
		v1060 = v1013
		goto L57
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1009
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1020 < int32(3) {
		goto L82
	} else {
		goto L229
	}
L229:
	;
	v1023 = F_r_remove_second_order_prefix_2(m, l0)
	mBase = m.M
	v1024 = m.ExcPending
	if v1024 != 0 {
		goto L64
	} else {
		goto L230
	}
L230:
	;
	if int32(0) <= v1023 {
		goto L82
	} else {
		goto L231
	}
L231:
	;
	if v1023 < int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1029 = v1023
	goto L234
L233:
	;
	v1029 = v321
	goto L234
L234:
	;
	if v1023 != 0 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v1030 = v1029
	goto L237
L236:
	;
	v1030 = v321
	goto L237
L237:
	;
	return v1030
L238:
	;
	if v1033 < int32(0) {
		v1060 = v1033
		goto L57
	} else {
		goto L239
	}
L239:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1037 < int32(3) {
		goto L82
	} else {
		goto L240
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v365
	v1041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1041
	v1043 = F_r_remove_suffix_2(m, l0)
	mBase = m.M
	v1044 = m.ExcPending
	if v1044 != 0 {
		goto L64
	} else {
		goto L241
	}
L241:
	;
	if base.Ui32(int32(7)) < base.Ui32(int32(base.Ui32(v1043)>>(uint(int32(31))%32))-int32(1)) {
		goto L82
	} else {
		goto L242
	}
L242:
	;
	if int32(0) <= v1043 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1054 = int32(1)
	goto L245
L244:
	;
	v1054 = v1043
	goto L245
L245:
	;
	return v1054
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
	var v2 int32
	_ = v2
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
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
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
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v612 int32
	_ = v612
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v652 int32
	_ = v652
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v671 int32
	_ = v671
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v766 int32
	_ = v766
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v811 int32
	_ = v811
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
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
	var v829 int32
	_ = v829
	var v835 int32
	_ = v835
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	v2 = int32(0)
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
	return v878
L7:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v475)+68))
	if v474 < v476 {
		goto L114
	} else {
		goto L115
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L110
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L105
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L101
	}
L11:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+76))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v276+v277<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v277 + int32(1)
	v285 = int32(0)
	v287 = *(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[0]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v287
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v291 = F_GetCachedPlan(m, v281, v289, v287, v285)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
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
		v878 = v29
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v254 + int32(1)
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
	v208 = int32(0)
	v211 = base.B2i32(v52+int32(1) < v46)
	if v211 == v208 {
		goto L48
	} else {
		goto L49
	}
L24:
	;
	if v152 == int32(0) {
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
	v148 = v77
	v152 = v134
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
	v145 = F_pg_analyze_and_rewrite_withcb(m, v57, v141, int32(509), v143, int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v148 = v139
	v152 = v145
	goto L24
L39:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v155 <= int32(0) {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v161 = int32(0)
	goto L41
L41:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v158+v161<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	if v180 != int32(6) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L23
L43:
	;
	v190 = v161 + int32(1)
	if v155 != v190 {
		v161 = v190
		goto L41
	} else {
		goto L47
	}
L44:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+28))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
	if v184 != int32(213) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v183)+12))
	if v187 != 0 {
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
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v33)+48))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v33)+60))
	v216 = int32(*(*int8)(unsafe.Add(mBase, uint32(v33)+58)))
	v218 = F_check_sql_stmt_retval(m, v152, v214, v215, v216, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	v221 = v208
	goto L50
L50:
	;
	v222 = int32(0)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v33)+40))
	F_CompleteCachedPlan(m, v148, v152, v222, v222, v222, int32(509), v226, int32(2052), v222)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L4
	} else {
		goto L52
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v33)+56)) = uint8(v218)
	v221 = v33
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v148)+36)) = int32(733)
	v234 = int32(_a_F_init_execution_state_2)
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1]))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v33)+84))
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v237
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v33)+76))
	v240 = F_lappend(m, v239, v148)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+76)) = v240
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v235
	F_SaveCachedPlan(m, v148)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
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
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v33)+80))
	F_MemoryContextDelete(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v291
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	if v294 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	v296 = v295
	goto L60
L59:
	;
	v296 = v285
	goto L60
L60:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v297 < v296 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v299 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v316 = v294
	goto L63
L63:
	;
	if v316 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v310
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v316 = v314
	goto L63
L65:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v305 = F_MemoryContextAlloc(m, v302, v296*int32(20))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v308 = F_repalloc_mul(m, v299, int32(20), v296)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L4
	} else {
		goto L69
	}
L68:
	;
	v310 = v305
	goto L64
L69:
	;
	v310 = v308
	goto L64
L70:
	;
	v461 = int32(0)
	goto L7
L71:
	;
	goto L72
L72:
	;
	v320 = int32(0)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	if v321 <= v320 {
		v461 = v320
		goto L7
	} else {
		goto L73
	}
L73:
	;
	v327 = int32(0)
	v328 = v320
	v330 = v2
	goto L74
L74:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v316)+12))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v341+v330<<(uint(int32(2))%32))))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v346 != int32(6) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v461 = v403
	goto L7
L76:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v382 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+57)))
	if v382 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L77:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345)+100))
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v349)))
	if v350 != int32(157) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if v350 != int32(225) {
		goto L76
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v349)+20))
	if v376 == int32(0) {
		goto L8
	} else {
		goto L87
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v345)+100))
	v363 = F_CreateCommandName(m, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v363
	F_errmsg(m, int32(_a_F_init_execution_state_3), v19+int32(16))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_init_execution_state_4), int32(745), int32(_a_F_init_execution_state_5))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
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
	v385 = F_CommandIsReadOnly(m, v345)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v392 = v389 + v330*int32(20)
	if v327 != 0 {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	if v385 == int32(0) {
		goto L9
	} else {
		goto L92
	}
L92:
	;
	goto L90
L93:
	;
	v395 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v392)+8)) = uint16(v395)
	*(*int64)(unsafe.Add(mBase, uint32(v392))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v392)+16)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v392)+12)) = v345
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+30)))
	if v402 != 0 {
		goto L97
	} else {
		goto L98
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v327))) = v392
	goto L93
L95:
	;
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v392
	goto L93
L97:
	;
	v403 = v392
	goto L99
L98:
	;
	v403 = v328
	goto L99
L99:
	;
	v405 = v330 + int32(1)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v316)+4))
	if v405 < v406 {
		v327 = v392
		v328 = v403
		v330 = v405
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
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(_a_F_init_execution_state_6), int32(0))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(_a_F_init_execution_state_4), int32(2076), int32(_a_F_init_execution_state_7))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
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
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L106
	}
L106:
	;
	v431 = F_CreateCommandName(m, v345)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v431
	F_errmsg(m, int32(_a_F_init_execution_state_8), v19)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_init_execution_state_4), int32(753), int32(_a_F_init_execution_state_5))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
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
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(_a_F_init_execution_state_9), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_init_execution_state_4), int32(738), int32(_a_F_init_execution_state_5))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
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
	v878 = int32(1)
	goto L6
L115:
	;
	goto L116
L116:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v475)+48))
	if v479 == int32(2278) {
		v835 = v475
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v846 = int32(1)
	v847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835)+55)))
	if v847 != v846 {
		goto L201
	} else {
		goto L202
	}
L118:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v482 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+24))
	if v483 == v485 {
		v835 = v475
		goto L117
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v487 == int32(0) {
		goto L124
	} else {
		goto L125
	}
L122:
	;
	goto L121
L123:
	;
	v502 = int32(_a_F_init_execution_state_2)
	v503 = *(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v501
	v508 = F_MakeSingleTupleTableSlot(m, int32(0), int32(_a_F_init_execution_state_10))
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L4
	} else {
		goto L129
	}
L124:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v495 = F_AllocSetContextCreateInternal(m, v490, int32(_a_F_init_execution_state_11), int32(0), int32(1024), int32(_a_F_init_execution_state_12))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L4
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	F_MemoryContextReset(m, v487)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L128
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+64)) = v495
	v501 = v495
	goto L123
L128:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v501 = v500
	goto L123
L129:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v281)+60))
	if v510 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v667)+60))
	if v668 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L131:
	;
	v652 = int32(0)
	goto L130
L132:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v510)+4))
	if v513 <= int32(0) {
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v517 = v513 & int32(3)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v510)+12))
	v519 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v513) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v612 == int32(0) {
		goto L131
	} else {
		goto L160
	}
L135:
	;
	v527 = v519
	v531 = v519
	v538 = v2
	goto L138
L136:
	;
	v565 = v519
	v569 = v519
	goto L137
L137:
	;
	v581 = v565
	v582 = v519
	v585 = v569
	goto L154
L138:
	;
	v544 = v518 + v527<<(uint(int32(2))%32)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v544)+12))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v544)+8))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v544)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v544)))
	v549 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548)+24)))
	if v549 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	if v517 == int32(0) {
		v612 = v556
		goto L134
	} else {
		goto L153
	}
L140:
	;
	v550 = v548
	goto L142
L141:
	;
	v550 = v531
	goto L142
L142:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547)+24)))
	if v551 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v552 = v547
	goto L145
L144:
	;
	v552 = v550
	goto L145
L145:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v546)+24)))
	if v553 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v554 = v546
	goto L148
L147:
	;
	v554 = v552
	goto L148
L148:
	;
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545)+24)))
	if v555 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v556 = v545
	goto L151
L150:
	;
	v556 = v554
	goto L151
L151:
	;
	v557 = int32(4)
	v558 = v527 + v557
	v560 = v538 + v557
	if v560 != v513&int32(2147483644) {
		v527 = v558
		v531 = v556
		v538 = v560
		goto L138
	} else {
		goto L152
	}
L152:
	;
	goto L139
L153:
	;
	v565 = v558
	v569 = v556
	goto L137
L154:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v518+v581<<(uint(int32(2))%32))))
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599)+24)))
	if v600 != 0 {
		goto L156
	} else {
		goto L157
	}
L155:
	;
	v612 = v601
	goto L134
L156:
	;
	v601 = v599
	goto L158
L157:
	;
	v601 = v585
	goto L158
L158:
	;
	v602 = int32(1)
	v605 = v582 + v602
	if v605 != v517 {
		v581 = v581 + v602
		v582 = v605
		v585 = v601
		goto L154
	} else {
		goto L159
	}
L159:
	;
	goto L155
L160:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v612)+4))
	if v625 == int32(1) {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v612)+76))
	v652 = v628
	goto L130
L162:
	;
	goto L163
L163:
	;
	if base.Ui32(int32(3)) < base.Ui32(v625-int32(2)) {
		goto L131
	} else {
		goto L164
	}
L164:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v612)+96))
	if v633 != 0 {
		v652 = v633
		goto L130
	} else {
		goto L165
	}
L165:
	;
	goto L131
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v811
	*(*int32)(unsafe.Add(mBase, uint32(v811)+4)) = int32(0)
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v815)+56)))
	if v816 == int32(1) {
		goto L197
	} else {
		goto L198
	}
L167:
	;
	v793 = F_ExecInitJunkFilter(m, v652, v508)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L4
	} else {
		goto L196
	}
L168:
	;
	v671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v667)+56)))
	if v671 != int32(1) {
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v674 = int32(0)
	if v508 != 0 {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v668)))
	if int32(0) < v683 {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	F_ExecSetSlotDescriptor(m, v508, v668)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L4
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v680 = F_MakeSingleTupleTableSlot(m, v668, int32(_a_F_init_execution_state_13))
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L4
	} else {
		goto L175
	}
L174:
	;
	v682 = v508
	goto L170
L175:
	;
	v682 = v680
	goto L170
L176:
	;
	v688 = F_palloc0(m, v683<<(uint(int32(1))%32))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	v780 = v674
	goto L178
L178:
	;
	v785 = F_palloc0(m, int32(20))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L4
	} else {
		goto L195
	}
L179:
	;
	if v652 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v652)+12))
	v691 = v690
	goto L182
L181:
	;
	v691 = v674
	goto L182
L182:
	;
	v693 = v691
	v699 = v674
	goto L183
L183:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668+v699<<(uint(int32(3))%32))+34)))
	if v711&int32(4) == int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	v780 = v688
	goto L178
L185:
	;
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v652)+12))
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v652)+4))
	v722 = v693
	goto L188
L186:
	;
	v750 = v693
	goto L187
L187:
	;
	v766 = v699 + int32(1)
	if v766 != v683 {
		v693 = v750
		v699 = v766
		goto L183
	} else {
		goto L194
	}
L188:
	;
	v738 = v722 + int32(4)
	if base.Ui32(v738) < base.Ui32(v716+v717<<(uint(int32(2))%32)) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	v747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v742)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v688+v699<<(uint(int32(1))%32)))) = uint16(v747)
	v750 = v741
	goto L187
L190:
	;
	v741 = v738
	goto L192
L191:
	;
	v741 = int32(0)
	goto L192
L192:
	;
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v722)))
	v743 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v742)+26)))
	if v743 != 0 {
		v722 = v741
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
	*(*int32)(unsafe.Add(mBase, uint32(v785)+16)) = v682
	*(*int32)(unsafe.Add(mBase, uint32(v785)+12)) = v780
	*(*int32)(unsafe.Add(mBase, uint32(v785)+8)) = v668
	*(*int32)(unsafe.Add(mBase, uint32(v785)+4)) = v652
	*(*int32)(unsafe.Add(mBase, uint32(v785))) = int32(391)
	v811 = v785
	goto L166
L196:
	;
	v811 = v793
	goto L166
L197:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v819)+16))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v820)+12))
	v822 = F_BlessTupleDesc(m, v821)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L4
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v824)+24))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v825
	*(*int32)(unsafe.Add(mBase, _c_F_init_execution_state[1])) = v503
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v835 = v829
	goto L117
L200:
	;
	goto L199
L201:
	;
	if v461 == int32(0) {
		v878 = v846
		goto L6
	} else {
		goto L206
	}
L202:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835)+56)))
	if v850 != 0 {
		goto L201
	} else {
		goto L203
	}
L203:
	;
	v851 = *(*int32)(unsafe.Add(mBase, uint32(v835)+48))
	v852 = F_type_is_rowtype(m, v851)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L4
	} else {
		goto L204
	}
L204:
	;
	if v852 == int32(0) {
		goto L201
	} else {
		goto L205
	}
L205:
	;
	v856 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)) = uint8(v856)
	goto L201
L206:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v860 == int32(0) {
		v878 = v846
		goto L6
	} else {
		goto L207
	}
L207:
	;
	v863 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v461)+8)) = uint8(v863)
	v865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	if v865 != v863 {
		v878 = v846
		goto L6
	} else {
		goto L208
	}
L208:
	;
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v461)+12))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v868)+4))
	if v869 != int32(1) {
		v878 = v846
		goto L6
	} else {
		goto L209
	}
L209:
	;
	v872 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v868)+29)))
	if v872 != 0 {
		v878 = v846
		goto L6
	} else {
		goto L210
	}
L210:
	;
	v873 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v461)+9)) = uint8(v873)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)) = uint8(v873)
	v878 = v846
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
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v19 = v17
	goto L3
L2:
	;
	v19 = int32(0)
	goto L3
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[0]))
	v24 = F_object_aclcheck(m, int32(1255), l0, v22, int64(128))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = F_get_func_name(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[1]))
	if v32 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_aclcheck_error(m, v24, int32(19), v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	F_RunFunctionExecuteHook(m, l0)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v35 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L4
	} else {
		goto L67
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if int32(101) <= v36 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v40 = l3 + int32(16)
	F_fmgr_info_cxt(m, l0, v40, l5)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+40)) = v43
	v49 = F_palloc(m, v19<<(uint(int32(4))%32)+int32(24))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
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
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v94 = int32(0)
	if base.B2i32(l7 == v94)|base.B2i32(v66&int32(1) == v94) == v94 {
		goto L34
	} else {
		goto L35
	}
L25:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	F_errmsg(m, int32(_a_F_init_sexpr_0), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L27
	}
L27:
	;
	if l4 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v86 = F_exprLocation(m, l2)
	mBase = m.M
	F_executor_errposition(m, v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_errfinish(m, int32(_a_F_init_sexpr_1), int32(742), int32(_a_F_init_sexpr_2))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	v240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+59)) = uint8(v240)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+44)) = int64(0)
	m.G0 = v14 + int32(16)
	return
L34:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v108 = F_get_expr_result_type(m, v103, v14+int32(12), v14+int32(8))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	goto L33
L37:
	;
	v110 = int32(_a_F_init_sexpr_3)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = l5
	v114 = int32(1)
	if base.Ui32(v108-v114) <= base.Ui32(v114) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_init_sexpr[2])) = v111
	goto L33
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v119 = F_CreateTupleDescCopy(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	switch v108 {
	case 0:
		goto L45
	default:
		goto L43
	case 3:
		goto L44
	}
L42:
	;
	v121 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v121)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v119
	goto L38
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	goto L38
L44:
	;
	v227 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v227)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = int32(0)
	goto L38
L45:
	;
	v125 = F_CreateTemplateTupleDesc(m, int32(1))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v125
	v129 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_TupleDescInitEntry(m, v125, int32(1), v129, v130, int32(-1), v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v136 = int32(0)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v136 < v145 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+56)) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+52)) = v223
	goto L38
L49:
	;
	v149 = v135 + int32(28)
	v156 = v136
	v157 = v145
	v159 = v136
	goto L53
L50:
	;
	v213 = v136
	v220 = v145
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+20)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v135)+16)) = v213
	goto L48
L52:
	;
	v213 = v207
	v220 = v186
	goto L51
L53:
	;
	v165 = v149 + v145<<(uint(int32(3))%32) + v156*int32(100)
	v168 = v149 + v156<<(uint(int32(3))%32)
	if v145 != v157 {
		v186 = v157
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v207 = v145
	goto L52
L55:
	;
	v187 = int32(*(*int16)(unsafe.Add(mBase, uint32(v168)+2)))
	if v187 <= int32(0) {
		v207 = v156
		goto L52
	} else {
		goto L63
	}
L56:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+7)))
	if v170 != int32(118) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v186 = v156
	goto L55
L58:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+4)))
	if v173 != int32(1) {
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+6)))
	if v176&int32(6) != 0 {
		goto L57
	} else {
		goto L60
	}
L60:
	;
	v179 = int32(*(*int16)(unsafe.Add(mBase, uint32(v168)+2)))
	if v179 <= int32(0) {
		goto L57
	} else {
		goto L61
	}
L61:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+90)))
	if v182 != int32(118) {
		v186 = v145
		goto L55
	} else {
		goto L62
	}
L62:
	;
	goto L57
L63:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+90)))
	if v190 == int32(118) {
		v207 = v156
		goto L52
	} else {
		goto L64
	}
L64:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+5)))
	v199 = (v159 + v193 - int32(1)) & (int32(0) - v193)
	if int32(_a_F_init_sexpr_4) < v199 {
		v207 = v156
		goto L52
	} else {
		goto L65
	}
L65:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v168))) = uint16(v199)
	v205 = v156 + int32(1)
	if v205 != v145 {
		v156 = v205
		v157 = v186
		v159 = v199 + v187
		goto L53
	} else {
		goto L66
	}
L66:
	;
	goto L54
L67:
	;
	F_errcode(m, int32(50856197))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v254 = int32(100)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v254
	F_errmsg_plural(m, int32(_a_F_init_sexpr_5), int32(_a_F_init_sexpr_6), v254, v14)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_init_sexpr_1), int32(723), int32(_a_F_init_sexpr_2))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
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
func F_int24ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v2 <= v3))
}
func F_int24lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v3 < v2))
}
func F_int28lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v3 < v2))
}
func F_int28mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v8 = v7 - v4
	if base.B2i32(int64(0) < v4) != base.B2i32(v8 < v7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int28mi_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int28mi_1), int32(1175), int32(_a_F_int28mi_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		return v8
	}
}
func F_int2abs(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	if v2 == int32(_a_F_int2abs_0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int2abs_1), int32(0))
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int2abs_2), int32(1243), int32(_a_F_int2abs_3))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
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
		v23 = base.I32_extend16_s(v2)
		v25 = v23 >> (uint(int32(31)) % 32)
		return base.I64_extend_i32_u(v23 ^ v25 - v25)
	}
}
func F_int2eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_int2hashfast(m *base.Module, l0 int64) int32 {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	v4 = base.I32_extend16_s(base.I32_wrap_i64(l0))
	v5 = int32(16)
	v9 = (int32(base.Ui32(v4)>>(uint(v5)%32)) ^ v4) * int32(-2048144789)
	v14 = (int32(base.Ui32(v9)>>(uint(int32(13))%32)) ^ v9) * int32(-1028477387)
	return int32(base.Ui32(v14)>>(uint(v5)%32)) ^ v14
}
func F_int2lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 < v3))
}
func F_int2recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pq_getmsgint(m, v2, int32(2))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend16_s(base.I64_extend_i32_u(v4))
	}
}
func F_int2smaller(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	if v3 < v4 {
		v6 = v3
	} else {
		v6 = v4
	}
	return base.I64_extend_i32_s(v6)
}
func F_int2vectorin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v155 int64
	_ = v155
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_palloc0(m, int32(88))
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
	v22 = v15
	v24 = v17
	v25 = int32(32)
	v27 = int32(0)
	goto L3
L3:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if base.B2i32(base.Ui32(v31-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v31 == int32(32)) != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v155
L5:
	;
	v22 = v22 + int32(1)
	goto L3
L6:
	;
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L4
L8:
	;
	m.G0 = v12 + int32(48)
	goto L7
L9:
	;
	if v27 < v25 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = int32(21)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+4)) = int64(1)
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v27<<(uint(int32(3))%32) + int32(96)
	v155 = base.I64_extend_i32_u(v24)
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_int2vectorin[0])) = int32(0)
	v59 = F_strtox_2(m, v22, v12+int32(44), int32(10), int64(2147483648))
	mBase = m.M
	v60 = base.I32_wrap_i64(v59)
	goto L17
L13:
	;
	v50 = v24
	v51 = v25
	goto L12
L14:
	;
	goto L15
L15:
	;
	v48 = F_repalloc(m, v24, v25<<(uint(int32(2))%32)+int32(24))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v50 = v48
	v51 = v25 << (uint(int32(1)) % 32)
	goto L12
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	if v62 == v22 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v131 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v50+v27<<(uint(v131)%32))+24)) = uint16(v60)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v22 = v137
	v24 = v50
	v25 = v51
	v27 = v27 + v131
	goto L3
L19:
	;
	v155 = int64(0)
	goto L8
L20:
	;
	F_errsave_finish(m, v14, int32(_a_F_int2vectorin_0), v126, int32(_a_F_int2vectorin_1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L40
	}
L21:
	;
	v64 = F_errsave_start(m, v14)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_int2vectorin[0]))
	if base.B2i32(v79 != int32(68))&base.B2i32(base.Ui32(int32(-65537)) < base.Ui32(v60-int32(_a_F_int2vectorin_2))) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	if v64 == int32(0) {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_int2vectorin_3)
	F_errmsg(m, int32(_a_F_int2vectorin_4), v12)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v126 = int32(199)
	goto L20
L28:
	;
	v89 = F_errsave_start(m, v14)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	v106 = int32(32)
	if v105|v106 == v106 {
		goto L18
	} else {
		goto L35
	}
L31:
	;
	if v89 == int32(0) {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(_a_F_int2vectorin_3)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v22
	F_errmsg(m, int32(_a_F_int2vectorin_5), v12+int32(16))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v126 = int32(205)
	goto L20
L35:
	;
	v110 = F_errsave_start(m, v14)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if v110 == int32(0) {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(_a_F_int2vectorin_3)
	F_errmsg(m, int32(_a_F_int2vectorin_4), v12+int32(32))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v126 = int32(211)
	goto L20
L40:
	;
	goto L19
}
func F_int42ge(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v3 <= v2))
}
func F_int42lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 < v3))
}
func F_int42mi(m *base.Module, l0 int32) int64 {
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
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = v6 - v3
	if base.B2i32(int32(0) < v3) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int42mi_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int42mi_1), int32(1101), int32(_a_F_int42mi_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return base.I64_extend_i32_s(v7)
	}
}
func F_int42pl(m *base.Module, l0 int32) int64 {
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
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = v6 + v3
	if base.B2i32(v3 < int32(0)) != base.B2i32(v7 < v6) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int42pl_0), int32(0))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int42pl_1), int32(1087), int32(_a_F_int42pl_2))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
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
		return base.I64_extend_i32_s(v7)
	}
}
func F_int48lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	return base.I64_extend_i32_u(base.B2i32(v3 < v2))
}
func F_int48mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	v8 = v7 - v4
	if base.B2i32(int64(0) < v4) != base.B2i32(v8 < v7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int48mi_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int48mi_1), int32(1033), int32(_a_F_int48mi_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		return v8
	}
}
func F_int48pl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+24)))
	v8 = v4 + v7
	if base.B2i32(v4 < int64(0)) != base.B2i32(v8 < v7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int48pl_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int48pl_1), int32(1019), int32(_a_F_int48pl_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		return v8
	}
}
func F_int4abs(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 == int32(-2147483648) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int4abs_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4abs_1), int32(1229), int32(_a_F_int4abs_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		v25 = v3 >> (uint(int32(31)) % 32)
		return base.I64_extend_i32_u(v3 ^ v25 - v25)
	}
}
func F_int4eqfast(m *base.Module, l0 int64, l1 int64) int32 {
	return base.B2i32(base.I32_wrap_i64(l0) == base.I32_wrap_i64(l1))
}
func F_int4hashfast(m *base.Module, l0 int64) int32 {
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	v3 = base.I32_wrap_i64(l0)
	v4 = int32(16)
	v8 = (int32(base.Ui32(v3)>>(uint(v4)%32)) ^ v3) * int32(-2048144789)
	v13 = (int32(base.Ui32(v8)>>(uint(int32(13))%32)) ^ v8) * int32(-1028477387)
	return int32(base.Ui32(v13)>>(uint(v4)%32)) ^ v13
}
func F_int4lcm(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
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
	var v70 int64
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v6 == int32(0) {
		v70 = v5
		return v70
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v9 == int32(0) {
			v70 = v5
			return v70
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
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_int4lcm_1), int32(1361), int32(_a_F_int4lcm_2))
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
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
					if v55 == int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_int4lcm_1), int32(1367), int32(_a_F_int4lcm_2))
									mBase = m.M
									v121 = m.ExcPending
									if v121 != 0 {
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
						v62 = v55 >> (uint(int32(31)) % 32)
						v70 = base.I64_extend_i32_u(v55 ^ v62 - v62)
						return v70
					}
				}
			} else {
				if v21&int32(2147483647) == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_int4lcm_1), int32(1293), int32(_a_F_int4lcm_3))
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
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_int4lcm_1), int32(1361), int32(_a_F_int4lcm_2))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
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
						if v55 == int32(-2147483648) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50331778))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_int4lcm_0), int32(0))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_int4lcm_1), int32(1367), int32(_a_F_int4lcm_2))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
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
							v62 = v55 >> (uint(int32(31)) % 32)
							v70 = base.I64_extend_i32_u(v55 ^ v62 - v62)
							return v70
						}
					}
				}
			}
		}
	}
}
func F_int4mod(m *base.Module, l0 int32) int64 {
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
	var v28 int64
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	switch v4 + int32(1) {
	case 0:
		v28 = int64(0)
		return v28
	case 1:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(33816706))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int4mod_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int4mod_1), int32(1168), int32(_a_F_int4mod_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	default:
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v26 = base.I32_rem_s(v25, v4)
		v28 = base.I64_extend_i32_s(v26)
		return v28
	}
}
func F_int4range_canonical(m *base.Module, l0 int32) int64 {
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
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		if v20 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			if v21 == v18 {
				v31 = v20
				F_range_deserialize(m, v31, v13, v10+int32(32), v10+int32(16), v10+int32(15))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v40 == int32(1) {
						v120 = v13
						m.G0 = v10 + int32(48)
						return base.I64_extend_i32_u(v120)
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)))
						if v43 != 0 {
							v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
							if v77 != 0 {
								v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int64(0)
								} else {
									v120 = v118
									m.G0 = v10 + int32(48)
									return base.I64_extend_i32_u(v120)
								}
							} else {
								v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
								if v78&int32(1) == int32(0) {
									v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int64(0)
									} else {
										v120 = v118
										m.G0 = v10 + int32(48)
										return base.I64_extend_i32_u(v120)
									}
								} else {
									v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
									if v83 == int64(2147483647) {
										v86 = int32(0)
										v87 = F_errsave_start(m, v17)
										mBase = m.M
										v88 = m.ExcPending
										if v88 != 0 {
											return int64(0)
										} else {
											if v87 == int32(0) {
												v120 = v86
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
														mBase = m.M
														v102 = m.ExcPending
														if v102 != 0 {
															return int64(0)
														} else {
															v120 = v86
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v120)
														}
													}
												}
											}
										}
									} else {
										v103 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
										v105 = int64(32)
										*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
										v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v120 = v118
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v120)
										}
									}
								}
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)))
							if v44&int32(1) != 0 {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
								if v77 != 0 {
									v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int64(0)
									} else {
										v120 = v118
										m.G0 = v10 + int32(48)
										return base.I64_extend_i32_u(v120)
									}
								} else {
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
									if v78&int32(1) == int32(0) {
										v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v120 = v118
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v120)
										}
									} else {
										v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
										if v83 == int64(2147483647) {
											v86 = int32(0)
											v87 = F_errsave_start(m, v17)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int64(0)
											} else {
												if v87 == int32(0) {
													v120 = v86
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int64(0)
															} else {
																v120 = v86
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v120)
															}
														}
													}
												}
											}
										} else {
											v103 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
											v105 = int64(32)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
											v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = v118
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											}
										}
									}
								}
							} else {
								v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)))
								if v47 == int64(2147483647) {
									v50 = int32(0)
									v51 = F_errsave_start(m, v17)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int64(0)
									} else {
										if v51 == int32(0) {
											v120 = v50
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v120)
										} else {
											F_errcode(m, int32(50331778))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1722), int32(_a_F_int4range_canonical_2))
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int64(0)
													} else {
														v120 = v50
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v120)
													}
												}
											}
										}
									}
								} else {
									v67 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)) = uint8(v67)
									v69 = int64(32)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = (v47<<(uint(v69)%64) + int64(4294967296)) >> (uint(v69) % 64)
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v77 != 0 {
										v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v120 = v118
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v120)
										}
									} else {
										v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v78&int32(1) == int32(0) {
											v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = v118
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											}
										} else {
											v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
											if v83 == int64(2147483647) {
												v86 = int32(0)
												v87 = F_errsave_start(m, v17)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int64(0)
												} else {
													if v87 == int32(0) {
														v120 = v86
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v120)
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return int64(0)
																} else {
																	v120 = v86
																	m.G0 = v10 + int32(48)
																	return base.I64_extend_i32_u(v120)
																}
															}
														}
													}
												}
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
												v105 = int64(32)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
												v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int64(0)
												} else {
													v120 = v118
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
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
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
							F_errmsg_internal(m, int32(_a_F_int4range_canonical_3), v10)
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_int4range_canonical_1), int32(1946), int32(_a_F_int4range_canonical_4))
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int64(0)
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
						F_range_deserialize(m, v31, v13, v10+int32(32), v10+int32(16), v10+int32(15))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v40 == int32(1) {
								v120 = v13
								m.G0 = v10 + int32(48)
								return base.I64_extend_i32_u(v120)
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)))
								if v43 != 0 {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v77 != 0 {
										v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v120 = v118
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v120)
										}
									} else {
										v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v78&int32(1) == int32(0) {
											v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = v118
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											}
										} else {
											v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
											if v83 == int64(2147483647) {
												v86 = int32(0)
												v87 = F_errsave_start(m, v17)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int64(0)
												} else {
													if v87 == int32(0) {
														v120 = v86
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v120)
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return int64(0)
																} else {
																	v120 = v86
																	m.G0 = v10 + int32(48)
																	return base.I64_extend_i32_u(v120)
																}
															}
														}
													}
												}
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
												v105 = int64(32)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
												v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int64(0)
												} else {
													v120 = v118
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
												}
											}
										}
									}
								} else {
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)))
									if v44&int32(1) != 0 {
										v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
										if v77 != 0 {
											v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = v118
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											}
										} else {
											v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
											if v78&int32(1) == int32(0) {
												v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int64(0)
												} else {
													v120 = v118
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
												}
											} else {
												v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
												if v83 == int64(2147483647) {
													v86 = int32(0)
													v87 = F_errsave_start(m, v17)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int64(0)
													} else {
														if v87 == int32(0) {
															v120 = v86
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v120)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int64(0)
																} else {
																	F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return int64(0)
																	} else {
																		v120 = v86
																		m.G0 = v10 + int32(48)
																		return base.I64_extend_i32_u(v120)
																	}
																}
															}
														}
													}
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
													v105 = int64(32)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
													v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int64(0)
													} else {
														v120 = v118
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v120)
													}
												}
											}
										}
									} else {
										v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)))
										if v47 == int64(2147483647) {
											v50 = int32(0)
											v51 = F_errsave_start(m, v17)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int64(0)
											} else {
												if v51 == int32(0) {
													v120 = v50
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1722), int32(_a_F_int4range_canonical_2))
															mBase = m.M
															v66 = m.ExcPending
															if v66 != 0 {
																return int64(0)
															} else {
																v120 = v50
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v120)
															}
														}
													}
												}
											}
										} else {
											v67 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)) = uint8(v67)
											v69 = int64(32)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = (v47<<(uint(v69)%64) + int64(4294967296)) >> (uint(v69) % 64)
											v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
											if v77 != 0 {
												v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int64(0)
												} else {
													v120 = v118
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
												}
											} else {
												v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
												if v78&int32(1) == int32(0) {
													v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int64(0)
													} else {
														v120 = v118
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v120)
													}
												} else {
													v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
													if v83 == int64(2147483647) {
														v86 = int32(0)
														v87 = F_errsave_start(m, v17)
														mBase = m.M
														v88 = m.ExcPending
														if v88 != 0 {
															return int64(0)
														} else {
															if v87 == int32(0) {
																v120 = v86
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v120)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return int64(0)
																	} else {
																		F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
																		mBase = m.M
																		v102 = m.ExcPending
																		if v102 != 0 {
																			return int64(0)
																		} else {
																			v120 = v86
																			m.G0 = v10 + int32(48)
																			return base.I64_extend_i32_u(v120)
																		}
																	}
																}
															}
														}
													} else {
														v103 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
														v105 = int64(32)
														*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
														v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
														mBase = m.M
														v119 = m.ExcPending
														if v119 != 0 {
															return int64(0)
														} else {
															v120 = v118
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v120)
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
				return int64(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
						F_errmsg_internal(m, int32(_a_F_int4range_canonical_3), v10)
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_int4range_canonical_1), int32(1946), int32(_a_F_int4range_canonical_4))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return int64(0)
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
					F_range_deserialize(m, v31, v13, v10+int32(32), v10+int32(16), v10+int32(15))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v40 == int32(1) {
							v120 = v13
							m.G0 = v10 + int32(48)
							return base.I64_extend_i32_u(v120)
						} else {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)))
							if v43 != 0 {
								v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
								if v77 != 0 {
									v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int64(0)
									} else {
										v120 = v118
										m.G0 = v10 + int32(48)
										return base.I64_extend_i32_u(v120)
									}
								} else {
									v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
									if v78&int32(1) == int32(0) {
										v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v120 = v118
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v120)
										}
									} else {
										v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
										if v83 == int64(2147483647) {
											v86 = int32(0)
											v87 = F_errsave_start(m, v17)
											mBase = m.M
											v88 = m.ExcPending
											if v88 != 0 {
												return int64(0)
											} else {
												if v87 == int32(0) {
													v120 = v86
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
														mBase = m.M
														v97 = m.ExcPending
														if v97 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
															mBase = m.M
															v102 = m.ExcPending
															if v102 != 0 {
																return int64(0)
															} else {
																v120 = v86
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v120)
															}
														}
													}
												}
											}
										} else {
											v103 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
											v105 = int64(32)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
											v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = v118
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											}
										}
									}
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)))
								if v44&int32(1) != 0 {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v77 != 0 {
										v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int64(0)
										} else {
											v120 = v118
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v120)
										}
									} else {
										v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v78&int32(1) == int32(0) {
											v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = v118
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											}
										} else {
											v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
											if v83 == int64(2147483647) {
												v86 = int32(0)
												v87 = F_errsave_start(m, v17)
												mBase = m.M
												v88 = m.ExcPending
												if v88 != 0 {
													return int64(0)
												} else {
													if v87 == int32(0) {
														v120 = v86
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v120)
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
																	return int64(0)
																} else {
																	v120 = v86
																	m.G0 = v10 + int32(48)
																	return base.I64_extend_i32_u(v120)
																}
															}
														}
													}
												}
											} else {
												v103 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
												v105 = int64(32)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
												v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int64(0)
												} else {
													v120 = v118
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
												}
											}
										}
									}
								} else {
									v47 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+32)))
									if v47 == int64(2147483647) {
										v50 = int32(0)
										v51 = F_errsave_start(m, v17)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int64(0)
										} else {
											if v51 == int32(0) {
												v120 = v50
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1722), int32(_a_F_int4range_canonical_2))
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return int64(0)
														} else {
															v120 = v50
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v120)
														}
													}
												}
											}
										}
									} else {
										v67 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)) = uint8(v67)
										v69 = int64(32)
										*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = (v47<<(uint(v69)%64) + int64(4294967296)) >> (uint(v69) % 64)
										v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
										if v77 != 0 {
											v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
												return int64(0)
											} else {
												v120 = v118
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v120)
											}
										} else {
											v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
											if v78&int32(1) == int32(0) {
												v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
													return int64(0)
												} else {
													v120 = v118
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v120)
												}
											} else {
												v83 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10)+16)))
												if v83 == int64(2147483647) {
													v86 = int32(0)
													v87 = F_errsave_start(m, v17)
													mBase = m.M
													v88 = m.ExcPending
													if v88 != 0 {
														return int64(0)
													} else {
														if v87 == int32(0) {
															v120 = v86
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v120)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_int4range_canonical_0), int32(0))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return int64(0)
																} else {
																	F_errsave_finish(m, v17, int32(_a_F_int4range_canonical_1), int32(1735), int32(_a_F_int4range_canonical_2))
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
																		return int64(0)
																	} else {
																		v120 = v86
																		m.G0 = v10 + int32(48)
																		return base.I64_extend_i32_u(v120)
																	}
																}
															}
														}
													}
												} else {
													v103 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v103)
													v105 = int64(32)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = (v83<<(uint(v105)%64) + int64(4294967296)) >> (uint(v105) % 64)
													v118 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
													mBase = m.M
													v119 = m.ExcPending
													if v119 != 0 {
														return int64(0)
													} else {
														v120 = v118
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v120)
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
		v73 = F_make_result_safe(m, v11+int32(8), int32(0))
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
func F_int82eq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3))
}
func F_int82pl(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = v4 + v7
	if base.B2i32(v4 < int64(0)) != base.B2i32(v8 < v7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int82pl_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int82pl_1), int32(1080), int32(_a_F_int82pl_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		return v8
	}
}
func F_int84(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(int64(-4294967297)) < base.Ui64(v3-int64(2147483648)) {
		v29 = v3
		return v29
	} else {
		v8 = int64(0)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = F_errsave_start(m, v9)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			if v10 == int32(0) {
				v29 = v8
				return v29
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_int84_0), int32(0))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						F_errsave_finish(m, v9, int32(_a_F_int84_1), int32(1296), int32(_a_F_int84_2))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int64(0)
						} else {
							v29 = v8
							return v29
						}
					}
				}
			}
		}
	}
}
func F_int84lt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_u(base.B2i32(v2 < v3))
}
func F_int84mi(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = v7 - v4
	if base.B2i32(int64(0) < v4) != base.B2i32(v8 < v7) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int84mi_0), int32(0))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int84mi_1), int32(952), int32(_a_F_int84mi_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		return v8
	}
}
func F_int8abs(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
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
	var v26 int64
	_ = v26
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v4 == int64(-9223372036854775807-1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int8abs_0), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8abs_1), int32(562), int32(_a_F_int8abs_2))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
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
		v26 = v4 >> (uint(int64(63)) % 64)
		return v4 ^ v26 - v26
	}
}
func F_int8range_canonical(m *base.Module, l0 int32) int64 {
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
	var v47 int64
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
		if v20 != 0 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			if v21 == v18 {
				v31 = v20
				F_range_deserialize(m, v31, v13, v10+int32(32), v10+int32(16), v10+int32(15))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
					if v40 == int32(1) {
						v112 = v13
						m.G0 = v10 + int32(48)
						return base.I64_extend_i32_u(v112)
					} else {
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)))
						if v43 != 0 {
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
							if v73 != 0 {
								v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int64(0)
								} else {
									v112 = v110
									m.G0 = v10 + int32(48)
									return base.I64_extend_i32_u(v112)
								}
							} else {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
								if v74&int32(1) == int32(0) {
									v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int64(0)
									} else {
										v112 = v110
										m.G0 = v10 + int32(48)
										return base.I64_extend_i32_u(v112)
									}
								} else {
									v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
									if v79 == int64(9223372036854775807) {
										v82 = int32(0)
										v83 = F_errsave_start(m, v17)
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int64(0)
										} else {
											if v83 == int32(0) {
												v112 = v82
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
													mBase = m.M
													v93 = m.ExcPending
													if v93 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
														mBase = m.M
														v98 = m.ExcPending
														if v98 != 0 {
															return int64(0)
														} else {
															v112 = v82
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v112)
														}
													}
												}
											}
										}
									} else {
										v99 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
										*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
										v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int64(0)
										} else {
											v112 = v110
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v112)
										}
									}
								}
							}
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)))
							if v44&int32(1) != 0 {
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
								if v73 != 0 {
									v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int64(0)
									} else {
										v112 = v110
										m.G0 = v10 + int32(48)
										return base.I64_extend_i32_u(v112)
									}
								} else {
									v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
									if v74&int32(1) == int32(0) {
										v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int64(0)
										} else {
											v112 = v110
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v112)
										}
									} else {
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
										if v79 == int64(9223372036854775807) {
											v82 = int32(0)
											v83 = F_errsave_start(m, v17)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int64(0)
											} else {
												if v83 == int32(0) {
													v112 = v82
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int64(0)
															} else {
																v112 = v82
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v112)
															}
														}
													}
												}
											}
										} else {
											v99 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
											v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int64(0)
											} else {
												v112 = v110
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											}
										}
									}
								}
							} else {
								v47 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
								if v47 == int64(9223372036854775807) {
									v50 = int32(0)
									v51 = F_errsave_start(m, v17)
									mBase = m.M
									v52 = m.ExcPending
									if v52 != 0 {
										return int64(0)
									} else {
										if v51 == int32(0) {
											v112 = v50
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v112)
										} else {
											F_errcode(m, int32(50331778))
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1769), int32(_a_F_int8range_canonical_2))
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int64(0)
													} else {
														v112 = v50
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v112)
													}
												}
											}
										}
									}
								} else {
									v67 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)) = uint8(v67)
									*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v47 + int64(1)
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v73 != 0 {
										v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int64(0)
										} else {
											v112 = v110
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v112)
										}
									} else {
										v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v74&int32(1) == int32(0) {
											v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int64(0)
											} else {
												v112 = v110
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											}
										} else {
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
											if v79 == int64(9223372036854775807) {
												v82 = int32(0)
												v83 = F_errsave_start(m, v17)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int64(0)
												} else {
													if v83 == int32(0) {
														v112 = v82
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v112)
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return int64(0)
																} else {
																	v112 = v82
																	m.G0 = v10 + int32(48)
																	return base.I64_extend_i32_u(v112)
																}
															}
														}
													}
												}
											} else {
												v99 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
												v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int64(0)
												} else {
													v112 = v110
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
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
					return int64(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
					if v26 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
							F_errmsg_internal(m, int32(_a_F_int8range_canonical_3), v10)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_int8range_canonical_1), int32(1946), int32(_a_F_int8range_canonical_4))
								mBase = m.M
								v131 = m.ExcPending
								if v131 != 0 {
									return int64(0)
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
						F_range_deserialize(m, v31, v13, v10+int32(32), v10+int32(16), v10+int32(15))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
							if v40 == int32(1) {
								v112 = v13
								m.G0 = v10 + int32(48)
								return base.I64_extend_i32_u(v112)
							} else {
								v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)))
								if v43 != 0 {
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v73 != 0 {
										v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int64(0)
										} else {
											v112 = v110
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v112)
										}
									} else {
										v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v74&int32(1) == int32(0) {
											v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int64(0)
											} else {
												v112 = v110
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											}
										} else {
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
											if v79 == int64(9223372036854775807) {
												v82 = int32(0)
												v83 = F_errsave_start(m, v17)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int64(0)
												} else {
													if v83 == int32(0) {
														v112 = v82
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v112)
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return int64(0)
																} else {
																	v112 = v82
																	m.G0 = v10 + int32(48)
																	return base.I64_extend_i32_u(v112)
																}
															}
														}
													}
												}
											} else {
												v99 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
												v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int64(0)
												} else {
													v112 = v110
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
												}
											}
										}
									}
								} else {
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)))
									if v44&int32(1) != 0 {
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
										if v73 != 0 {
											v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int64(0)
											} else {
												v112 = v110
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											}
										} else {
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
											if v74&int32(1) == int32(0) {
												v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int64(0)
												} else {
													v112 = v110
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
												}
											} else {
												v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
												if v79 == int64(9223372036854775807) {
													v82 = int32(0)
													v83 = F_errsave_start(m, v17)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int64(0)
													} else {
														if v83 == int32(0) {
															v112 = v82
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v112)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int64(0)
																} else {
																	F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return int64(0)
																	} else {
																		v112 = v82
																		m.G0 = v10 + int32(48)
																		return base.I64_extend_i32_u(v112)
																	}
																}
															}
														}
													}
												} else {
													v99 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
													v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int64(0)
													} else {
														v112 = v110
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v112)
													}
												}
											}
										}
									} else {
										v47 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
										if v47 == int64(9223372036854775807) {
											v50 = int32(0)
											v51 = F_errsave_start(m, v17)
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int64(0)
											} else {
												if v51 == int32(0) {
													v112 = v50
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1769), int32(_a_F_int8range_canonical_2))
															mBase = m.M
															v66 = m.ExcPending
															if v66 != 0 {
																return int64(0)
															} else {
																v112 = v50
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v112)
															}
														}
													}
												}
											}
										} else {
											v67 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)) = uint8(v67)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v47 + int64(1)
											v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
											if v73 != 0 {
												v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int64(0)
												} else {
													v112 = v110
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
												}
											} else {
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
												if v74&int32(1) == int32(0) {
													v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int64(0)
													} else {
														v112 = v110
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v112)
													}
												} else {
													v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
													if v79 == int64(9223372036854775807) {
														v82 = int32(0)
														v83 = F_errsave_start(m, v17)
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return int64(0)
														} else {
															if v83 == int32(0) {
																v112 = v82
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v112)
															} else {
																F_errcode(m, int32(50331778))
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return int64(0)
																	} else {
																		F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
																		mBase = m.M
																		v98 = m.ExcPending
																		if v98 != 0 {
																			return int64(0)
																		} else {
																			v112 = v82
																			m.G0 = v10 + int32(48)
																			return base.I64_extend_i32_u(v112)
																		}
																	}
																}
															}
														}
													} else {
														v99 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
														*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
														v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
														mBase = m.M
														v111 = m.ExcPending
														if v111 != 0 {
															return int64(0)
														} else {
															v112 = v110
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v112)
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
				return int64(0)
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+200))
				if v26 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
						F_errmsg_internal(m, int32(_a_F_int8range_canonical_3), v10)
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_int8range_canonical_1), int32(1946), int32(_a_F_int8range_canonical_4))
							mBase = m.M
							v131 = m.ExcPending
							if v131 != 0 {
								return int64(0)
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
					F_range_deserialize(m, v31, v13, v10+int32(32), v10+int32(16), v10+int32(15))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
						if v40 == int32(1) {
							v112 = v13
							m.G0 = v10 + int32(48)
							return base.I64_extend_i32_u(v112)
						} else {
							v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)))
							if v43 != 0 {
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
								if v73 != 0 {
									v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int64(0)
									} else {
										v112 = v110
										m.G0 = v10 + int32(48)
										return base.I64_extend_i32_u(v112)
									}
								} else {
									v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
									if v74&int32(1) == int32(0) {
										v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int64(0)
										} else {
											v112 = v110
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v112)
										}
									} else {
										v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
										if v79 == int64(9223372036854775807) {
											v82 = int32(0)
											v83 = F_errsave_start(m, v17)
											mBase = m.M
											v84 = m.ExcPending
											if v84 != 0 {
												return int64(0)
											} else {
												if v83 == int32(0) {
													v112 = v82
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
												} else {
													F_errcode(m, int32(50331778))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int64(0)
															} else {
																v112 = v82
																m.G0 = v10 + int32(48)
																return base.I64_extend_i32_u(v112)
															}
														}
													}
												}
											}
										} else {
											v99 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
											*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
											v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int64(0)
											} else {
												v112 = v110
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											}
										}
									}
								}
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)))
								if v44&int32(1) != 0 {
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
									if v73 != 0 {
										v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
										mBase = m.M
										v111 = m.ExcPending
										if v111 != 0 {
											return int64(0)
										} else {
											v112 = v110
											m.G0 = v10 + int32(48)
											return base.I64_extend_i32_u(v112)
										}
									} else {
										v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v74&int32(1) == int32(0) {
											v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int64(0)
											} else {
												v112 = v110
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											}
										} else {
											v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
											if v79 == int64(9223372036854775807) {
												v82 = int32(0)
												v83 = F_errsave_start(m, v17)
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return int64(0)
												} else {
													if v83 == int32(0) {
														v112 = v82
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v112)
													} else {
														F_errcode(m, int32(50331778))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return int64(0)
																} else {
																	v112 = v82
																	m.G0 = v10 + int32(48)
																	return base.I64_extend_i32_u(v112)
																}
															}
														}
													}
												}
											} else {
												v99 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
												*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
												v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int64(0)
												} else {
													v112 = v110
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
												}
											}
										}
									}
								} else {
									v47 = *(*int64)(unsafe.Add(mBase, uint32(v10)+32))
									if v47 == int64(9223372036854775807) {
										v50 = int32(0)
										v51 = F_errsave_start(m, v17)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int64(0)
										} else {
											if v51 == int32(0) {
												v112 = v50
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											} else {
												F_errcode(m, int32(50331778))
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1769), int32(_a_F_int8range_canonical_2))
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return int64(0)
														} else {
															v112 = v50
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v112)
														}
													}
												}
											}
										}
									} else {
										v67 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+41)) = uint8(v67)
										*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v47 + int64(1)
										v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+24)))
										if v73 != 0 {
											v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int64(0)
											} else {
												v112 = v110
												m.G0 = v10 + int32(48)
												return base.I64_extend_i32_u(v112)
											}
										} else {
											v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
											if v74&int32(1) == int32(0) {
												v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int64(0)
												} else {
													v112 = v110
													m.G0 = v10 + int32(48)
													return base.I64_extend_i32_u(v112)
												}
											} else {
												v79 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
												if v79 == int64(9223372036854775807) {
													v82 = int32(0)
													v83 = F_errsave_start(m, v17)
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return int64(0)
													} else {
														if v83 == int32(0) {
															v112 = v82
															m.G0 = v10 + int32(48)
															return base.I64_extend_i32_u(v112)
														} else {
															F_errcode(m, int32(50331778))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_int8range_canonical_0), int32(0))
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return int64(0)
																} else {
																	F_errsave_finish(m, v17, int32(_a_F_int8range_canonical_1), int32(1782), int32(_a_F_int8range_canonical_2))
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return int64(0)
																	} else {
																		v112 = v82
																		m.G0 = v10 + int32(48)
																		return base.I64_extend_i32_u(v112)
																	}
																}
															}
														}
													}
												} else {
													v99 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)) = uint8(v99)
													*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v79 + int64(1)
													v110 = F_range_serialize(m, v31, v10+int32(32), v10+int32(16), int32(0), v17)
													mBase = m.M
													v111 = m.ExcPending
													if v111 != 0 {
														return int64(0)
													} else {
														v112 = v110
														m.G0 = v10 + int32(48)
														return base.I64_extend_i32_u(v112)
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
func F_int8um(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v3 == int64(-9223372036854775807-1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50331778))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_int8um_0), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8um_1), int32(455), int32(_a_F_int8um_2))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
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
		return int64(0) - v3
	}
}
func F_intarray_push_array(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v13 = F_intarray_concat_arrays(m, v6, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				if v15 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return int64(0)
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v19 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v13)
							}
						} else {
							return base.I64_extend_i32_u(v13)
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v19 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v13)
						}
					} else {
						return base.I64_extend_i32_u(v13)
					}
				}
			}
		}
	}
}
func F_inter_lb(m *base.Module, l0 int32) int64 {
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
	var v22 int64
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
	var v60 int64
	_ = v60
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	v16 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v17 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v11)+24)) = v17
	*(*float64)(unsafe.Add(mBase, uint32(v11)+16)) = v16
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	*(*float64)(unsafe.Add(mBase, uint32(v11))) = v16
	v22 = int64(1)
	v24 = F_lseg_interpt_line(m, int32(0), v11, v13)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int64(0)
	} else {
		if v24 != 0 {
			v60 = v22
			m.G0 = v11 + int32(32)
			return v60
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
				return int64(0)
			} else {
				if v35 != 0 {
					v60 = v22
					m.G0 = v11 + int32(32)
					return v60
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
						return int64(0)
					} else {
						if v44 != 0 {
							v60 = v22
							m.G0 = v11 + int32(32)
							return v60
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
								return int64(0)
							} else {
								v60 = base.I64_extend_i32_u(v53)
								m.G0 = v11 + int32(32)
								return v60
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
				F_errfinish(m, int32(_a_F_internalerrposition_1), int32(1680), int32(_a_F_internalerrposition_2))
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
						F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1249), int32(_a_F_intervaltypmodleastfield_2))
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
							F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1249), int32(_a_F_intervaltypmodleastfield_2))
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
								F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1249), int32(_a_F_intervaltypmodleastfield_2))
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
							F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1249), int32(_a_F_intervaltypmodleastfield_2))
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
								F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1249), int32(_a_F_intervaltypmodleastfield_2))
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
							F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1249), int32(_a_F_intervaltypmodleastfield_2))
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
									F_errfinish(m, int32(_a_F_intervaltypmodleastfield_1), int32(1249), int32(_a_F_intervaltypmodleastfield_2))
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
	return base.B2i32(base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(34)) < base.Ui32(l0)) == int32(0)) & base.B2i32(int64(1)<<(uint(base.I64_extend_i32_u(l0))%64)&int64(34357509982) != int64(0))
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
func F_isn_out(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14316(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_iso8859_to_utf8(m *base.Module, l0 int32) int64 {
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14317(m, l0, int32(_a_F_iso8859_to_utf8_0), int32(_a_F_iso8859_to_utf8_1), int32(133), int32(_a_F_iso8859_to_utf8_2), int32(_a_F_iso8859_to_utf8_3), int32(_a_F_iso8859_to_utf8_4), int32(19), int32(9))
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		return v10
	}
}
func F_iso_to_koi8r(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14323(m, l0, int32(_a_F_iso_to_koi8r_0), int32(22), int32(25))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_iso_to_win1251(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14323(m, l0, int32(_a_F_iso_to_win1251_0), int32(23), int32(25))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
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
						F_errfinish(m, int32(_a_F_iteratorFromContainer_1), int32(1157), int32(_a_F_iteratorFromContainer_2))
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
	F_LockBufferInternal(m, v56, int32(1))
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
	var v48 int64
	_ = v48
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
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v98 float64
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v158 float64
	_ = v158
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v172 float64
	_ = v172
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v210 float64
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
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
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v389 int64
	_ = v389
	var v396 int32
	_ = v396
	var v421 int32
	_ = v421
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
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
	v451 = m.ExcPending
	if v451 != 0 {
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
	v48 = base.I64_extend_i32_u(v46)
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v48
	v51 = F_HnswOptionalProcInfo(m, l0, int32(2))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[0])) = v40
	F_MemoryContextDelete(m, v35)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L5
	} else {
		goto L86
	}
L10:
	;
	if v51 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	v55 = F_IvfflatCheckNorm(m, v51, v54, v48)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	v63 = v48
	goto L13
L13:
	;
	v64 = int32(0)
	F_IvfflatGetMetaPageInfo(m, l0, v64, v64)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L17
	}
L14:
	;
	if v55 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v59 = F_HnswNormValue(m, v43, v54, v48)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+24)) = v59
	v63 = v59
	goto L13
L17:
	;
	v68 = int32(1)
	v70 = F_index_getprocinfo(m, l0, v68, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v75 = int32(1)
	v83 = v75
	v84 = v75
	v85 = int32(-1)
	v86 = v75
	v98 = float64(1.7976931348623157e+308)
	goto L19
L19:
	;
	v100 = F_ReadBuffer(m, l0, v86)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L21
	}
L20:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v24)+20)) = uint16(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v196
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v224 = F_index_form_tuple(m, v221, v24+int32(24), l2)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L39
	}
L21:
	;
	F_LockBufferInternal(m, v100, int32(1))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	if v100 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+16)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v122+v212)))
	F_UnlockReleaseBuffer(m, v100)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L37
	}
L24:
	;
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+12)))
	if base.Ui32(v123) < base.Ui32(int32(25)) {
		v195 = v83
		v196 = v84
		v197 = v85
		v210 = v98
		goto L23
	} else {
		goto L28
	}
L25:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[1]))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108+(v100^int32(-1))<<(uint(int32(2))%32))))
	v122 = v114
	goto L24
L26:
	;
	goto L27
L27:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[2]))
	v122 = v116 + v100<<(uint(int32(13))%32) + int32(-8192)
	goto L24
L28:
	;
	v127 = v123 + int32(_a_F_ivfflatinsert_4)
	if v127&int32(_a_F_ivfflatinsert_5) == int32(0) {
		v195 = v83
		v196 = v84
		v197 = v85
		v210 = v98
		goto L23
	} else {
		goto L29
	}
L29:
	;
	v140 = int32(1)
	v143 = v83
	v144 = v84
	v145 = v85
	v158 = v98
	goto L30
L30:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v122+int32(20)+v140<<(uint(int32(2))%32))))
	v166 = v122 + v163&int32(_a_F_ivfflatinsert_6)
	v170 = F_FunctionCall2Coll(m, v70, v73, v63, base.I64_extend_i32_u(v166+int32(8)))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v195 = v182
	v196 = v183
	v197 = v184
	v210 = v185
	goto L23
L32:
	;
	v172 = base.F64_reinterpret_i64(v170)
	v174 = int32(0)
	if base.B2i32(base.F64_gt(v158, v172) == v174)&base.B2i32(v145 != int32(-1)) == v174 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v182 = v140
	v183 = v86
	v184 = v181
	v185 = v172
	goto L35
L34:
	;
	v182 = v143
	v183 = v144
	v184 = v145
	v185 = v158
	goto L35
L35:
	;
	if base.B2i32(v140 == int32(base.Ui32(v127)>>(uint(int32(2))%32))&int32(_a_F_ivfflatinsert_7)) == int32(0) {
		v140 = v140 + int32(1)
		v143 = v182
		v144 = v183
		v145 = v184
		v158 = v185
		goto L30
	} else {
		goto L36
	}
L36:
	;
	goto L31
L37:
	;
	if v214 != int32(-1) {
		v83 = v195
		v84 = v196
		v85 = v197
		v86 = v214
		v98 = v210
		goto L19
	} else {
		goto L38
	}
L38:
	;
	goto L20
L39:
	;
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v224)+4)) = uint16(v226)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v224))) = v228
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224)+6)))
	v231 = F_ReadBuffer(m, l0, v197)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	F_LockBufferInternal(m, v231, int32(3))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v236 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L43
	}
L42:
	;
	v380 = int32(0)
	v382 = F_PageAddItemExtended(m, v362, v224, v255, v380, v380)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L81
	}
L43:
	;
	v239 = F_GenericXLogRegisterBuffer(m, v236, v231, int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	v241 = int32(4)
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239)+14)))
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239)+12)))
	v244 = v242 - v243
	if v244 <= v241 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v255 = (v230&int32(_a_F_ivfflatinsert_8) + int32(7)) & int32(_a_F_ivfflatinsert_9)
	if base.Ui32(v255) <= base.Ui32(v247-int32(4)) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v247 = v241
	goto L48
L47:
	;
	v247 = v244
	goto L48
L48:
	;
	goto L45
L49:
	;
	v360 = v231
	v362 = v239
	v363 = v197
	v364 = v236
	goto L42
L50:
	;
	goto L51
L51:
	;
	v258 = v231
	v260 = v239
	v262 = v236
	goto L52
L52:
	;
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+16)))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v260+v278)))
	if v280 != int32(-1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	F_LockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L5
	} else {
		goto L68
	}
L54:
	;
	F_pfree(m, v262)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
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
	F_UnlockReleaseBuffer(m, v258)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L58
	}
L58:
	;
	v287 = F_ReadBuffer(m, l0, v280)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L5
	} else {
		goto L59
	}
L59:
	;
	F_LockBufferInternal(m, v287, int32(3))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v292 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	v295 = F_GenericXLogRegisterBuffer(m, v292, v287, int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	v297 = int32(4)
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+14)))
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+12)))
	v300 = v298 - v299
	if v300 <= v297 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	if base.Ui32(v303-int32(4)) < base.Ui32(v255) {
		v258 = v287
		v260 = v295
		v262 = v292
		goto L52
	} else {
		goto L67
	}
L64:
	;
	v303 = v297
	goto L66
L65:
	;
	v303 = v300
	goto L66
L66:
	;
	goto L63
L67:
	;
	v360 = v287
	v362 = v295
	v363 = v280
	v364 = v292
	goto L42
L68:
	;
	v311 = F_HnswNewBuffer(m, l0, int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L69
	}
L69:
	;
	F_UnlockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	v317 = F_GenericXLogRegisterBuffer(m, v262, v311, int32(1))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L5
	} else {
		goto L71
	}
L71:
	;
	F_PageInit(m, v317, int32(_a_F_ivfflatinsert_1), int32(8))
	mBase = m.M
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v317)+16)))
	v323 = v317 + v322
	v324 = int32(_a_F_ivfflatinsert_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v323)+6)) = uint16(v324)
	*(*int32)(unsafe.Add(mBase, uint32(v323))) = int32(-1)
	goto L72
L72:
	;
	if v311 < int32(0) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v260+v347))) = v346
	F_GenericXLogFinish(m, v262)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L77
	}
L74:
	;
	v331 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[3]))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v331+(v311^int32(-1))*int32(56))+16))
	v346 = v337
	goto L73
L75:
	;
	goto L76
L76:
	;
	v339 = *(*int32)(unsafe.Add(mBase, _c_F_ivfflatinsert[4]))
	v340 = int32(56)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v339+v311*v340-v340)+16))
	v346 = v345
	goto L73
L77:
	;
	F_UnlockReleaseBuffer(m, v258)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L5
	} else {
		goto L78
	}
L78:
	;
	v354 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v357 = F_GenericXLogRegisterBuffer(m, v354, v311, int32(0))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L5
	} else {
		goto L80
	}
L80:
	;
	v360 = v311
	v362 = v357
	v363 = v346
	v364 = v354
	goto L42
L81:
	;
	if v382 == int32(0) {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_IvfflatCommitBuffer(m, v360, v364)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L83
	}
L83:
	;
	if v363 == v197 {
		goto L9
	} else {
		goto L84
	}
L84:
	;
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v24)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v24)+8)) = v389
	F_IvfflatUpdateList(m, l0, v24+int32(8), v363, v197, int32(-1), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
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
	v452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v452 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ivfflatinsert_11), v24)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_ivfflatinsert_12), int32(174), int32(_a_F_ivfflatinsert_13))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
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
