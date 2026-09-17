package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckFunctionValidatorAccess(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = F_SearchSysCache1(m, int32(47), l1)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		if v13 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+22)))
			v20 = v18 + v19
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
			v22 = F_SearchSysCache1(m, int32(36), v21)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v83
						F_errmsg_internal(m, int32(_a_F_CheckFunctionValidatorAccess_0), v10+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_CheckFunctionValidatorAccess_1), int32(2170), int32(_a_F_CheckFunctionValidatorAccess_2))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+22)))
					v28 = v26 + v27
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
					if v29 != l0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v28)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v103
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v102
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
								F_errmsg(m, int32(_a_F_CheckFunctionValidatorAccess_3), v10+int32(32))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CheckFunctionValidatorAccess_1), int32(2178), int32(_a_F_CheckFunctionValidatorAccess_2))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
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
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_CheckFunctionValidatorAccess[0]))
						v36 = F_object_aclcheck(m, int32(2612), v32, v34, int64(256))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							if v36 != 0 {
								F_aclcheck_error(m, v36, int32(21), v28+int32(4))
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return int32(0)
								} else {
									v45 = *(*int32)(unsafe.Add(mBase, _c_F_CheckFunctionValidatorAccess[0]))
									v47 = F_object_aclcheck(m, int32(1255), l1, v45, int64(128))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return int32(0)
									} else {
										if v47 != 0 {
											F_aclcheck_error(m, v47, int32(19), v20+int32(4))
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v13)
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v22)
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
														return int32(0)
													} else {
														m.G0 = v10 + int32(48)
														return int32(1)
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v13)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v22)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(48)
													return int32(1)
												}
											}
										}
									}
								}
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, _c_F_CheckFunctionValidatorAccess[0]))
								v47 = F_object_aclcheck(m, int32(1255), l1, v45, int64(128))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									if v47 != 0 {
										F_aclcheck_error(m, v47, int32(19), v20+int32(4))
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v13)
											mBase = m.M
											v55 = m.ExcPending
											if v55 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v22)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(48)
													return int32(1)
												}
											}
										}
									} else {
										F_ReleaseCatCache(m, v13)
										mBase = m.M
										v55 = m.ExcPending
										if v55 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v22)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(48)
												return int32(1)
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
			v66 = m.ExcPending
			if v66 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg(m, int32(_a_F_CheckFunctionValidatorAccess_4), v10)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_CheckFunctionValidatorAccess_1), int32(2161), int32(_a_F_CheckFunctionValidatorAccess_2))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
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
func F_ExecFunctionScan(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(712), int32(713))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_FunctionCall6Coll(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v31 int32
	_ = v31
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
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v8 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+76)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+68)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+60)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+52)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l1
	v31 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+30)) = uint16(v31)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l0
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = m.T0[v42].(func(*base.Module, int32) int32)(m, v11+int32(12))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		return int32(0)
	} else {
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
		if v47 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v54
				F_errmsg_internal(m, int32(_a_F_FunctionCall6Coll_0), v11)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall6Coll_1), int32(1278), int32(_a_F_FunctionCall6Coll_2))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v11 + int32(80)
			return v43
		}
	}
}
func F_SendFunctionCall(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v8 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v6)+13)) = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+28)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = l1
	v16 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+22)) = uint16(v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = m.T0[v20].(func(*base.Module, int32) int32)(m, v6+int32(4))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+20)))
		if v25 != int32(1) {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
			if v28&int32(3) != 0 {
				v31 = F_detoast_attr(m, v21)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = v31
					m.G0 = v6 + int32(32)
					return v33
				}
			} else {
				v33 = v21
				m.G0 = v6 + int32(32)
				return v33
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v42
				F_errmsg_internal(m, int32(_a_F_SendFunctionCall_0), v6)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_SendFunctionCall_1), int32(1143), int32(_a_F_SendFunctionCall_2))
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
func F_add_function_cost(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v22 int64
	_ = v22
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v41 float64
	_ = v41
	var v42 float64
	_ = v42
	var v46 float32
	_ = v46
	var v49 float64
	_ = v49
	var v51 float64
	_ = v51
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v14 = F_SearchSysCache1(m, int32(47), l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v18 = v16 + v17
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+92))
			if v19 == int32(0) {
				v46 = *(*float32)(unsafe.Add(mBase, uint32(v18)+80))
				v49 = *(*float64)(unsafe.Add(mBase, _c_F_add_function_cost[0]))
				v51 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
				*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_promote_f32(v46), v49), v51)
				F_ReleaseCatCache(m, v14)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					m.G0 = v11 + int32(48)
					return
				}
			} else {
				v22 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(459)
				*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v22
				v33 = v11 + int32(16)
				v34 = F_OidFunctionCall1Coll(m, v19, int32(0), v33)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					if v34 != v33 {
						v46 = *(*float32)(unsafe.Add(mBase, uint32(v18)+80))
						v49 = *(*float64)(unsafe.Add(mBase, _c_F_add_function_cost[0]))
						v51 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_promote_f32(v46), v49), v51)
					} else {
						v37 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
						v38 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v37, v38)
						v41 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
						v42 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(v41, v42)
					}
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						m.G0 = v11 + int32(48)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
				F_errmsg_internal(m, int32(_a_F_add_function_cost_0), v11)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_add_function_cost_1), int32(2133), int32(_a_F_add_function_cost_2))
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
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
func F_check_transform_function(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
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
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+101)))
	if v7 != int32(118) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
		if v10 != int32(102) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return
			} else {
				F_errcode(m, int32(117833860))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_check_transform_function_0), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1811), int32(_a_F_check_transform_function_2))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
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
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+100)))
			if v13 == int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					F_errcode(m, int32(117833860))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_check_transform_function_3), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1815), int32(_a_F_check_transform_function_2))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
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
				v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+104)))
				if v16 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_check_transform_function_4), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1819), int32(_a_F_check_transform_function_2))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
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
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
					if v19 != int32(2281) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_check_transform_function_5)
								F_errmsg(m, int32(_a_F_check_transform_function_6), v5)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1824), int32(_a_F_check_transform_function_2))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
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
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			F_errcode(m, int32(117833860))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_check_transform_function_7), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1807), int32(_a_F_check_transform_function_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
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
func F_compute_function_hashkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	v7 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v7
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v17
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v17
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v24
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v26 != 0 {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
		v30 = base.B2i32(v27 == int32(442))
	} else {
		v30 = v7
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v30)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v32 != 0 {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
		v36 = base.B2i32(v33 == int32(441))
	} else {
		v36 = v7
	}
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v36)
	v39 = int32(0)
	if l5|base.B2i32(v30 == v39) == v39 {
		v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+20))
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v46
	} else {
	}
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v48
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+104)))
	if int32(0) < v50 {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v50
		v55 = l2 + int32(28)
		v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+104)))
		v58 = v56 << (uint(int32(2)) % 32)
		if v58 != 0 {
			base.MemoryCopy(m, v55, l1+int32(136), v58)
		} else {
		}
		v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+104)))
		v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
		F_cfunc_resolve_polymorphic_argtypes(m, v62, v55, int32(0), v65, l5, l1+int32(4))
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return
		} else {
			if l4 == int32(0) {
				m.G0 = v13 + int32(16)
				return
			} else {
				v78 = F_get_call_result_type(m, l0, v13+int32(12), v13+int32(8))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					v80 = int32(1)
					if base.Ui32(v80) < base.Ui32(v78-v80) {
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v84
					}
					m.G0 = v13 + int32(16)
					return
				}
			}
		}
	} else {
		if l4 == int32(0) {
			m.G0 = v13 + int32(16)
			return
		} else {
			v78 = F_get_call_result_type(m, l0, v13+int32(12), v13+int32(8))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return
			} else {
				v80 = int32(1)
				if base.Ui32(v80) < base.Ui32(v78-v80) {
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v84
				}
				m.G0 = v13 + int32(16)
				return
			}
		}
	}
}
func F_expand_function_arguments(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(400)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
	v24 = v22 + v23
	v26 = v24 + int32(136)
	v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+104)))
	if l1 == v5 {
		v52 = v27
		v53 = v26
		goto L9
	} else {
		goto L10
	}
L1:
	;
	m.G0 = v20 + int32(400)
	return v598
L2:
	;
	v484 = m.G0
	v486 = v484 - int32(800)
	m.G0 = v486
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+22)))
	if v471 == int32(0) {
		goto L111
	} else {
		goto L112
	}
L3:
	;
	if v52 <= v54 {
		goto L78
	} else {
		goto L79
	}
L4:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v62+v271<<(uint(int32(2))%32))))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	if v284 == int32(16) {
		goto L75
	} else {
		goto L76
	}
L5:
	;
	if v54&int32(1) == int32(0) {
		goto L3
	} else {
		goto L74
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L11
	} else {
		goto L71
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L11
	} else {
		goto L68
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L11
	} else {
		goto L65
	}
L9:
	;
	if l0 != 0 {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v32 = F_SysCacheGetAttr(m, int32(47), l3, int32(21), v20)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v36 != 0 {
		v52 = v27
		v53 = v26
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v37 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v39 != int32(1) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	if v42 < int32(0) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	if v45 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v46 != int32(26) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v52 = v42
	v53 = v37 + int32(24)
	goto L9
L19:
	;
	v200 = F_SysCacheGetAttrNotNull(m, int32(47), l3, int32(24))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L52
	}
L20:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v54 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v178 = int32(0)
	if v52 <= v178 {
		v598 = v178
		goto L1
	} else {
		goto L51
	}
L23:
	;
	if v54 < v52 {
		v185 = v54
		goto L19
	} else {
		goto L50
	}
L24:
	;
	v57 = int32(0)
	if v57 < v54 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v61 = v54
	goto L27
L26:
	;
	v61 = v57
	goto L27
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v64 = v57
	goto L28
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v62+v64<<(uint(int32(2))%32))))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v84 != int32(16) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if base.Ui32(int32(101)) <= base.Ui32(v52) {
		goto L7
	} else {
		goto L34
	}
L30:
	;
	v88 = v64 + int32(1)
	if v61 != v88 {
		v64 = v88
		goto L28
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L23
L34:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+22)))
	v95 = v52 << (uint(int32(2)) % 32)
	if v95 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	base.MemoryFill(m, v20, int32(0), v95)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v98 = int32(0)
	if v54 == int32(1) {
		v263 = v98
		v271 = v5
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v105 = v98
	v113 = v5
	v121 = v5
	goto L39
L39:
	;
	v124 = v62 + v113<<(uint(int32(2))%32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v126 != int32(16) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L5
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v134<<(uint(int32(2))%32)))) = v133
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	if v141 != int32(16) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v133 = v125
	v134 = v105
	v135 = v105 + int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v133 = v132
	v134 = v131
	v135 = v105
	goto L41
L45:
	;
	v151 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v20+v149<<(uint(v151)%32)))) = v148
	v156 = v113 + v151
	v158 = v121 + v151
	if v54&int32(2147483646) != v158 {
		v105 = v150
		v113 = v156
		v121 = v158
		goto L39
	} else {
		goto L49
	}
L46:
	;
	v148 = v140
	v149 = v135
	v150 = v135 + int32(1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v148 = v147
	v149 = v146
	v150 = v135
	goto L45
L49:
	;
	goto L40
L50:
	;
	v598 = l0
	goto L1
L51:
	;
	v185 = v178
	goto L19
L52:
	;
	v202 = F_text_to_cstring(m, v200)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v204 = F_stringToNode(m, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	F_pfree(m, v202)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	if v204 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v210 = v208
	goto L58
L57:
	;
	v210 = int32(0)
	goto L58
L58:
	;
	v211 = v210 + v185
	v212 = v211 - v52
	if v212 < int32(0) {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	if v211 != v52 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v216 = F_list_delete_first_n(m, v204, v212)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L11
	} else {
		goto L63
	}
L61:
	;
	v218 = v204
	goto L62
L62:
	;
	v219 = F_list_concat_copy(m, l0, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L11
	} else {
		goto L64
	}
L63:
	;
	v218 = v216
	goto L62
L64:
	;
	v471 = v219
	goto L2
L65:
	;
	F_errmsg_internal(m, int32(_a_F_expand_function_arguments_0), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_2), int32(_a_F_expand_function_arguments_3))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L11
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
	F_errmsg_internal(m, int32(_a_F_expand_function_arguments_4), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_5), int32(_a_F_expand_function_arguments_6))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errmsg_internal(m, int32(_a_F_expand_function_arguments_7), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_8), int32(_a_F_expand_function_arguments_9))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L11
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	v263 = v150
	v271 = v156
	goto L4
L75:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v283)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v289 = v287
	v290 = v288
	goto L77
L76:
	;
	v289 = v263
	v290 = v283
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v289<<(uint(int32(2))%32)))) = v290
	goto L3
L78:
	;
	if v52 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L79:
	;
	v315 = F_SysCacheGetAttrNotNull(m, int32(47), l3, int32(24))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	v317 = F_text_to_cstring(m, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	v319 = F_stringToNode(m, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	F_pfree(m, v317)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	if v319 == int32(0) {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v319)+4))
	if v325 <= int32(0) {
		goto L78
	} else {
		goto L85
	}
L85:
	;
	v329 = int32(*(*int16)(unsafe.Add(mBase, uint32(v92+v93)+106)))
	v330 = v52 - v329
	v331 = int32(0)
	if v325 != int32(1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v340 = v330
	v343 = v331
	v347 = int32(0)
	goto L89
L87:
	;
	v389 = v330
	v392 = v331
	goto L88
L88:
	;
	v407 = v20 + v389<<(uint(int32(2))%32)
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)))
	if v408 != 0 {
		goto L78
	} else {
		goto L99
	}
L89:
	;
	v358 = v20 + v340<<(uint(int32(2))%32)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	if v359 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v325&int32(1) == int32(0) {
		goto L78
	} else {
		goto L98
	}
L91:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v362+v343<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v358))) = v366
	goto L93
L92:
	;
	goto L93
L93:
	;
	v369 = v358 + int32(4)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v370 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373+v343<<(uint(int32(2))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v377
	goto L96
L95:
	;
	goto L96
L96:
	;
	v379 = int32(2)
	v380 = v343 + v379
	v382 = v340 + v379
	v384 = v347 + v379
	if v384 != v325&int32(2147483646) {
		v340 = v382
		v343 = v380
		v347 = v384
		goto L89
	} else {
		goto L97
	}
L97:
	;
	goto L90
L98:
	;
	v389 = v382
	v392 = v380
	goto L88
L99:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v319)+12))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v409+v392<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v407))) = v413
	goto L78
L100:
	;
	v471 = int32(0)
	goto L2
L101:
	;
	goto L102
L102:
	;
	v435 = int32(1)
	if v52 <= v435 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v438 = v435
	goto L105
L104:
	;
	v438 = v52
	goto L105
L105:
	;
	v439 = int32(0)
	v442 = v439
	v445 = v439
	goto L106
L106:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v20+v442<<(uint(int32(2))%32))))
	v462 = F_lappend(m, v445, v461)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L11
	} else {
		goto L108
	}
L107:
	;
	v471 = v462
	goto L2
L108:
	;
	v465 = v442 + int32(1)
	if v465 != v438 {
		v442 = v465
		v445 = v462
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v563 = v52 << (uint(int32(2)) % 32)
	if v563 != 0 {
		goto L125
	} else {
		goto L126
	}
L111:
	;
	v548 = int32(0)
	goto L110
L112:
	;
	goto L113
L113:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if int32(101) <= v493 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L11
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v509 = int32(0)
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v510 <= v509 {
		v548 = v509
		goto L110
	} else {
		goto L120
	}
L117:
	;
	F_errmsg_internal(m, int32(_a_F_expand_function_arguments_4), int32(0))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_10), int32(_a_F_expand_function_arguments_11))
	mBase = m.M
	v508 = m.ExcPending
	if v508 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v516 = v509
	goto L121
L121:
	;
	v531 = v516 << (uint(int32(2)) % 32)
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v471)+12))
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v535+v531)))
	v538 = F_exprType(m, v537)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L11
	} else {
		goto L123
	}
L122:
	;
	v548 = v542
	goto L110
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v531+(v486+int32(400))))) = v538
	v542 = v516 + int32(1)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v471)+4))
	if v542 < v543 {
		v516 = v542
		goto L121
	} else {
		goto L124
	}
L124:
	;
	goto L122
L125:
	;
	base.MemoryCopy(m, v486, v53, v563)
	goto L127
L126:
	;
	goto L127
L127:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v488+v489)+108))
	v570 = F_enforce_generic_type_consistency(m, v486+int32(400), v486, v548, v568, int32(0))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	if v570 != l2 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L11
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	F_make_fn_arguments(m, int32(0), v471, v486+int32(400), v486)
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L11
	} else {
		goto L135
	}
L132:
	;
	F_errmsg_internal(m, int32(_a_F_expand_function_arguments_12), int32(0))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L11
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_13), int32(_a_F_expand_function_arguments_11))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L11
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	m.G0 = v486 + int32(800)
	v598 = v471
	goto L1
}
func F_simplify_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v553 int32
	_ = v553
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v591 int32
	_ = v591
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v626 int32
	_ = v626
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
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
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v754 int32
	_ = v754
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v800 float64
	_ = v800
	var v801 float64
	_ = v801
	var v804 float64
	_ = v804
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v845 int32
	_ = v845
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v904 int32
	_ = v904
	var v910 int32
	_ = v910
	var v929 int32
	_ = v929
	var v936 int32
	_ = v936
	var v955 int32
	_ = v955
	v7 = l6
	v11 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(112)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v30 = F_SearchSysCache1(m, int32(47), l0)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	if base.B2i32(l8 == int32(0))|v331 != 0 {
		v936 = v331
		goto L58
	} else {
		goto L59
	}
L2:
	;
	if v294|base.B2i32(l8 == int32(0)) != 0 {
		v331 = v294
		goto L1
	} else {
		goto L53
	}
L3:
	;
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+101)))
	if v240 != int32(105) {
		goto L46
	} else {
		goto L47
	}
L4:
	;
	v204 = int32(1)
	v206 = int32(0)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+99)))
	if base.B2i32(v188&v204 == v206)|base.B2i32(v208 != v204) == v206 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v62+v149<<(uint(int32(2))%32))))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	if v171 != int32(7) {
		goto L38
	} else {
		goto L39
	}
L6:
	;
	if v59&int32(1) == int32(0) {
		v188 = v122
		v197 = v123
		goto L4
	} else {
		goto L37
	}
L7:
	;
	return int32(0)
L8:
	;
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	v36 = int32(0)
	if l7 == v36 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L34
	}
L12:
	;
	v51 = v50 + v48
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+100)))
	if v52 != 0 {
		v294 = v36
		goto L2
	} else {
		goto L18
	}
L13:
	;
	v48 = v34
	v49 = v28
	v50 = v35
	goto L12
L14:
	;
	goto L15
L15:
	;
	v40 = F_expand_function_arguments(m, v28, int32(0), l1, v30)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v43 = F_expression_tree_mutator_impl(m, v40, int32(871), l9)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
	v48 = v46
	v49 = v43
	v50 = v47
	goto L12
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+108))
	if v54 == int32(2249) {
		v294 = int32(0)
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v49 == int32(0) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v59 <= int32(0) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v63 = int32(0)
	if v59 == int32(1) {
		v149 = v63
		v151 = v63
		v160 = v11
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v77 = v63
	v79 = v63
	v85 = int32(0)
	v88 = v11
	goto L23
L23:
	;
	v97 = v62 + v77<<(uint(int32(2))%32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v99 != int32(7) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L6
L25:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v112 != int32(7) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v109 = v79
	v110 = int32(1)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+24)))
	v109 = base.B2i32(v103|v79&int32(1) != int32(0))
	v110 = v88
	goto L25
L29:
	;
	v124 = int32(2)
	v125 = v77 + v124
	v127 = v85 + v124
	if v59&int32(2147483646) != v127 {
		v77 = v125
		v79 = v122
		v85 = v127
		v88 = v123
		goto L23
	} else {
		goto L33
	}
L30:
	;
	v122 = v109
	v123 = int32(1)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+24)))
	v122 = base.B2i32(v116|v109&int32(1) != int32(0))
	v123 = v110
	goto L29
L33:
	;
	goto L24
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = l0
	F_errmsg_internal(m, int32(_a_F_simplify_function_0), v26)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_simplify_function_1), int32(4086), int32(_a_F_simplify_function_2))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v149 = v125
	v151 = v122
	v160 = v123
	goto L5
L38:
	;
	v188 = v151
	v197 = int32(1)
	goto L4
L39:
	;
	goto L40
L40:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+24)))
	v188 = base.B2i32(v175|v151&int32(1) != int32(0))
	v197 = v160
	goto L4
L41:
	;
	v214 = F_makeNullConst(m, l1, l2, l3)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L7
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v197 != 0 {
		v294 = int32(0)
		goto L2
	} else {
		goto L45
	}
L44:
	;
	v294 = v214
	goto L2
L45:
	;
	goto L3
L46:
	;
	if v240 != int32(115) {
		v294 = int32(0)
		goto L2
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v253 = F_palloc0(m, int32(36))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	v246 = int32(0)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l9)+16)))
	if v247&int32(1) == v246 {
		v294 = v246
		goto L2
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v253)+28)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v253)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v253)+20)) = l3
	v260 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v253)+16)) = v260
	*(*uint8)(unsafe.Add(mBase, uint32(v253)+13)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v253)+12)) = uint8(v260)
	*(*int32)(unsafe.Add(mBase, uint32(v253)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = int32(15)
	v269 = F_evaluate_expr(m, v253, l1, l2, l3)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v294 = v269
	goto L2
L53:
	;
	v298 = v34 + v35
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v298)+92))
	if v299 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v331 = int32(0)
	goto L1
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = int32(15)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+100)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = l3
	v313 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v313
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+45)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+44)) = uint8(v307)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = int32(457)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v26 + int32(32)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v298)+92))
	v328 = F_OidFunctionCall1Coll(m, v324, v313, v26+int32(88))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v331 = v328
	goto L1
L58:
	;
	F_ReleaseCatCache(m, v30)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L7
	} else {
		goto L225
	}
L59:
	;
	v335 = int32(0)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+22)))
	v338 = v336 + v337
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)+76))
	if v339 != int32(14) {
		v936 = v335
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+96)))
	if v342 != int32(102) {
		v936 = v335
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+97)))
	if v345 != 0 {
		v936 = v335
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+100)))
	if v346 != 0 {
		v936 = v335
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v338)+108))
	if v347 == int32(2249) {
		v936 = v335
		goto L58
	} else {
		goto L64
	}
L64:
	;
	v352 = F_heap_attisnull(m, v30, int32(29), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	if v352 == int32(0) {
		v936 = v335
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338)+104)))
	if v49 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v359 = v357
	goto L69
L68:
	;
	v359 = int32(0)
	goto L69
L69:
	;
	if v359 != v356 {
		v936 = v335
		goto L58
	} else {
		goto L70
	}
L70:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v362 = int32(0)
	if v361 == v362 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v400 != 0 {
		v936 = v335
		goto L58
	} else {
		goto L84
	}
L72:
	;
	v400 = int32(0)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	if v368 <= int32(0) {
		v394 = v362
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v400 = v394
	goto L71
L76:
	;
	v371 = int32(0)
	if v371 < v368 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v374 = v368
	goto L79
L78:
	;
	v374 = v371
	goto L79
L79:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v361)+12))
	v377 = int32(0)
	goto L80
L80:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v375+v377<<(uint(int32(2))%32))))
	v386 = base.B2i32(v385 == l0)
	if v385 == l0 {
		v394 = v386
		goto L75
	} else {
		goto L82
	}
L81:
	;
	v394 = v386
	goto L75
L82:
	;
	v388 = v377 + int32(1)
	if v388 != v374 {
		v377 = v388
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v403 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[0]))
	v405 = F_object_aclcheck(m, int32(1255), l0, v403, int64(128))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	if v405 != 0 {
		v936 = v335
		goto L58
	} else {
		goto L86
	}
L86:
	;
	v408 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[1]))
	if v408 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v409 = m.T0[v408].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L7
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2]))
	v417 = F_AllocSetContextCreateInternal(m, v412, int32(_a_F_simplify_function_3), int32(0), int32(_a_F_simplify_function_4), int32(_a_F_simplify_function_5))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L7
	} else {
		goto L92
	}
L90:
	;
	if v409 != 0 {
		v936 = v335
		goto L58
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v419 = int32(_a_F_simplify_function_6)
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2])) = v417
	v424 = F_palloc0(m, int32(36))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v424)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v424)+28)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v424)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v424)+20)) = l3
	v431 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v424)+16)) = v431
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+13)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v424)+12)) = uint8(v431)
	*(*int32)(unsafe.Add(mBase, uint32(v424)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v424)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v424))) = int32(15)
	v442 = F_SysCacheGetAttrNotNull(m, int32(47), v30, int32(26))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	v444 = F_text_to_cstring(m, v442)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v338 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = int32(873)
	v452 = int32(_a_F_simplify_function_7)
	v453 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[3])) = v26 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v26 + int32(100)
	v466 = F_SysCacheGetAttr(m, int32(47), v30, int32(28), v26+int32(111))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+111)))
	if v468 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[3])) = v929
	v936 = v910
	goto L58
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2])) = v420
	F_MemoryContextDelete(m, v417)
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L7
	} else {
		goto L224
	}
L99:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v626)))
	if v641 != int32(67) {
		goto L98
	} else {
		goto L134
	}
L100:
	;
	v471 = F_text_to_cstring(m, v466)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L7
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v495 = F_prepare_sql_fn_parse_info(m, v30, v424, l4)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L7
	} else {
		goto L112
	}
L103:
	;
	if v487 == int32(0) {
		goto L98
	} else {
		goto L110
	}
L104:
	;
	v473 = F_stringToNode(m, v471)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	if v475 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v473)+12))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)))
	v487 = v479
	goto L103
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v473
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v473
	v485 = F_list_make1_impl(m, int32(1), v26+int32(28))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v487 = v485
	goto L103
L110:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v487)+4))
	if v490 != int32(1) {
		goto L98
	} else {
		goto L111
	}
L111:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v487)+12))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	v626 = v494
	goto L99
L112:
	;
	v497 = F_pg_parse_query(m, v444)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	if v497 == int32(0) {
		goto L98
	} else {
		goto L114
	}
L114:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	if v501 != int32(1) {
		goto L98
	} else {
		goto L115
	}
L115:
	;
	v505 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v505)+4)) = v444
	*(*int32)(unsafe.Add(mBase, uint32(v505)+120)) = v495
	*(*int32)(unsafe.Add(mBase, uint32(v505)+112)) = int32(679)
	*(*int32)(unsafe.Add(mBase, uint32(v505)+108)) = int32(680)
	*(*int32)(unsafe.Add(mBase, uint32(v505)+104)) = int32(0)
	goto L117
L117:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v497)+12))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v515)))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v516)+4))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v517)))
	if v518 != int32(141) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v610 = F_transformStmt(m, v505, v591)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L7
	} else {
		goto L132
	}
L119:
	;
	v591 = v517
	goto L118
L120:
	;
	goto L121
L121:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v517)+68))
	if v521 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v528 = v517
	goto L125
L123:
	;
	v553 = v517
	goto L124
L124:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	if v570 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v528)+76))
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)+68))
	if v546 != 0 {
		v528 = v545
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v553 = v545
	goto L124
L127:
	;
	goto L126
L128:
	;
	v591 = v517
	goto L118
L129:
	;
	goto L130
L130:
	;
	v574 = F_palloc0(m, int32(20))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v574)+4)) = v517
	*(*int32)(unsafe.Add(mBase, uint32(v574))) = int32(242)
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v553)+8))
	v580 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v574)+16)) = uint8(v580)
	*(*int32)(unsafe.Add(mBase, uint32(v574)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v574)+8)) = v579
	*(*int32)(unsafe.Add(mBase, uint32(v553)+8)) = int32(0)
	v591 = v574
	goto L118
L132:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v516)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v610)+156)) = v612
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v516)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v610)+160)) = v614
	F_free_parsestate(m, v505)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v626 = v610
	goto L99
L134:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v626)+4))
	if v644 != int32(1) {
		goto L98
	} else {
		goto L135
	}
L135:
	;
	v647 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+36)))
	if v647 != 0 {
		goto L98
	} else {
		goto L136
	}
L136:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+37)))
	if v648 != 0 {
		goto L98
	} else {
		goto L137
	}
L137:
	;
	v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+38)))
	if v649 != 0 {
		goto L98
	} else {
		goto L138
	}
L138:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v626)+39)))
	if v650 != 0 {
		goto L98
	} else {
		goto L139
	}
L139:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v626)+48))
	if v651 != 0 {
		goto L98
	} else {
		goto L140
	}
L140:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v626)+52))
	if v652 != 0 {
		goto L98
	} else {
		goto L141
	}
L141:
	;
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v626)+60))
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v653)+4))
	if v654 != 0 {
		goto L98
	} else {
		goto L142
	}
L142:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v653)+8))
	if v655 != 0 {
		goto L98
	} else {
		goto L143
	}
L143:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v626)+100))
	if v656 != 0 {
		goto L98
	} else {
		goto L144
	}
L144:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v626)+108))
	if v657 != 0 {
		goto L98
	} else {
		goto L145
	}
L145:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v626)+112))
	if v658 != 0 {
		goto L98
	} else {
		goto L146
	}
L146:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v626)+116))
	if v659 != 0 {
		goto L98
	} else {
		goto L147
	}
L147:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v626)+120))
	if v660 != 0 {
		goto L98
	} else {
		goto L148
	}
L148:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v626)+124))
	if v661 != 0 {
		goto L98
	} else {
		goto L149
	}
L149:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v626)+128))
	if v662 != 0 {
		goto L98
	} else {
		goto L150
	}
L150:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v626)+132))
	if v663 != 0 {
		goto L98
	} else {
		goto L151
	}
L151:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v626)+144))
	if v664 != 0 {
		goto L98
	} else {
		goto L152
	}
L152:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v626)+76))
	if v665 == int32(0) {
		goto L98
	} else {
		goto L153
	}
L153:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v665)+4))
	if v668 != int32(1) {
		goto L98
	} else {
		goto L154
	}
L154:
	;
	v674 = F_get_expr_result_type(m, v424, int32(0), v26+int32(84))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L7
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v626
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v626
	v681 = F_list_make1_impl(m, int32(1), v26+int32(24))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L7
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v681
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v681
	v688 = F_list_make1_impl(m, int32(1), v26+int32(20))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v26)+84))
	v691 = int32(*(*int8)(unsafe.Add(mBase, uint32(v338)+96)))
	v693 = F_check_sql_fn_retval(m, v688, l1, v690, v691, int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L7
	} else {
		goto L158
	}
L158:
	;
	if v693 != 0 {
		goto L98
	} else {
		goto L159
	}
L159:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v681)+12))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	if v626 != v696 {
		goto L98
	} else {
		goto L160
	}
L160:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v626)+76))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+12))
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v699)))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v700)+4))
	v702 = F_exprType(m, v701)
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L7
	} else {
		goto L161
	}
L161:
	;
	if v702 != l1 {
		goto L98
	} else {
		goto L162
	}
L162:
	;
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+101)))
	if v705 == int32(105) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v709 = F_contain_mutable_functions_walker(m, v701, int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L7
	} else {
		goto L166
	}
L164:
	;
	v712 = v705
	goto L165
L165:
	;
	if v712&int32(255) == int32(115) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	if v709 != 0 {
		goto L98
	} else {
		goto L167
	}
L167:
	;
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+101)))
	v712 = v711
	goto L165
L168:
	;
	v718 = F_contain_volatile_functions_walker(m, v701, int32(0))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L7
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+99)))
	if v720 == int32(1) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	if v718 != 0 {
		goto L98
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v724 = F_contain_nonstrict_functions_walker(m, v701, int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L7
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = int32(0)
	v729 = v26 + int32(32)
	v730 = F_contain_context_dependent_node_walker(m, v49, v729)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L7
	} else {
		goto L178
	}
L176:
	;
	if v724 != 0 {
		goto L98
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	if v730 != 0 {
		goto L98
	} else {
		goto L179
	}
L179:
	;
	v732 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338)+104)))
	v735 = F_palloc0(m, v732<<(uint(int32(2))%32))
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L7
	} else {
		goto L180
	}
L180:
	;
	v737 = int32(*(*int16)(unsafe.Add(mBase, uint32(v338)+104)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v735
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v737
	v741 = F_substitute_actual_parameters_mutator(m, v701, v729)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L7
	} else {
		goto L181
	}
L181:
	;
	if v49 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2])) = v420
	v842 = F_copyObjectImpl(m, v741)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L7
	} else {
		goto L203
	}
L183:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v745 <= int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v754 = int32(0)
	goto L185
L185:
	;
	v773 = v754 << (uint(int32(2)) % 32)
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v735+v773)))
	switch v775 {
	case 0:
		goto L188
	case 1:
		goto L187
	default:
		goto L189
	}
L186:
	;
	goto L182
L187:
	;
	v814 = v754 + int32(1)
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v814 < v815 {
		v754 = v814
		goto L185
	} else {
		goto L202
	}
L188:
	;
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+99)))
	if v811 != 0 {
		goto L98
	} else {
		goto L201
	}
L189:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v776+v773)))
	if v778 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v778)))
	if base.Ui32(v779-int32(22)) < base.Ui32(int32(3)) {
		goto L98
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v778
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v778
	v795 = F_list_make1_impl(m, int32(1), v26+int32(16))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L7
	} else {
		goto L196
	}
L193:
	;
	v786 = F_expression_tree_walker_impl(m, v778, int32(857), int32(0))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L7
	} else {
		goto L194
	}
L194:
	;
	if v786 != 0 {
		goto L98
	} else {
		goto L195
	}
L195:
	;
	goto L192
L196:
	;
	F_cost_qual_eval(m, v26+int32(32), v795, int32(0))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L7
	} else {
		goto L197
	}
L197:
	;
	v800 = *(*float64)(unsafe.Add(mBase, uint32(v26)+32))
	v801 = *(*float64)(unsafe.Add(mBase, uint32(v26)+40))
	v804 = *(*float64)(unsafe.Add(mBase, _c_F_simplify_function[4]))
	if base.F64_gt(base.F64_add(v800, v801), base.F64_mul(v804, float64(10))) != 0 {
		goto L98
	} else {
		goto L198
	}
L198:
	;
	v809 = F_contain_volatile_functions_walker(m, v778, int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L7
	} else {
		goto L199
	}
L199:
	;
	if v809 != 0 {
		goto L98
	} else {
		goto L200
	}
L200:
	;
	goto L187
L201:
	;
	goto L187
L202:
	;
	goto L186
L203:
	;
	F_MemoryContextDelete(m, v417)
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L7
	} else {
		goto L204
	}
L204:
	;
	if l3 == int32(0) {
		goto L206
	} else {
		goto L207
	}
L205:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v864 != 0 {
		goto L217
	} else {
		goto L218
	}
L206:
	;
	v863 = v842
	goto L205
L207:
	;
	goto L208
L208:
	;
	v848 = F_exprCollation(m, v842)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L7
	} else {
		goto L209
	}
L209:
	;
	if v848 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v863 = v842
	goto L205
L211:
	;
	goto L212
L212:
	;
	if v848 == l3 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v863 = v842
	goto L205
L214:
	;
	goto L215
L215:
	;
	v854 = F_palloc0(m, int32(16))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L7
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v854)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v854)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v854)+4)) = v842
	*(*int32)(unsafe.Add(mBase, uint32(v854))) = int32(31)
	v863 = v854
	goto L205
L217:
	;
	F_record_plan_function_dependency(m, v864, l0)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L7
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v868 = F_lappend_oid(m, v867, l0)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L7
	} else {
		goto L221
	}
L220:
	;
	goto L219
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9)+8)) = v868
	v871 = F_eval_const_expressions_mutator(m, v863, l9)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L7
	} else {
		goto L222
	}
L222:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v874 = F_list_delete_last(m, v873)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L7
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9)+8)) = v874
	v910 = v871
	goto L97
L224:
	;
	v910 = int32(0)
	goto L97
L225:
	;
	m.G0 = v26 + int32(112)
	return v936
}
