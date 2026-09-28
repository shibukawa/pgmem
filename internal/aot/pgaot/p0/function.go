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
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l1))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
			v21 = v19 + v20
			v22 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v21)+76)))
			v23 = F_SearchSysCache1(m, int32(36), v22)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				if v23 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v83 = m.ExcPending
					if v83 != 0 {
						return int32(0)
					} else {
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v84
						F_errmsg_internal(m, int32(_a_F_CheckFunctionValidatorAccess_0), v10+int32(16))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_CheckFunctionValidatorAccess_1), int32(2136), int32(_a_F_CheckFunctionValidatorAccess_2))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
					v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+22)))
					v29 = v27 + v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
					if v30 != l0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(16797828))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int32(0)
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
								v104 = *(*int32)(unsafe.Add(mBase, uint32(v29)+84))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v104
								*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v103
								*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
								F_errmsg(m, int32(_a_F_CheckFunctionValidatorAccess_3), v10+int32(32))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_CheckFunctionValidatorAccess_1), int32(2144), int32(_a_F_CheckFunctionValidatorAccess_2))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
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
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_CheckFunctionValidatorAccess[0]))
						v37 = F_object_aclcheck(m, int32(2612), v33, v35, int64(256))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							if v37 != 0 {
								F_aclcheck_error(m, v37, int32(21), v29+int32(4))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, _c_F_CheckFunctionValidatorAccess[0]))
									v48 = F_object_aclcheck(m, int32(1255), l1, v46, int64(128))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										if v48 != 0 {
											F_aclcheck_error(m, v48, int32(19), v21+int32(4))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v14)
												mBase = m.M
												v56 = m.ExcPending
												if v56 != 0 {
													return int32(0)
												} else {
													F_ReleaseCatCache(m, v23)
													mBase = m.M
													v58 = m.ExcPending
													if v58 != 0 {
														return int32(0)
													} else {
														m.G0 = v10 + int32(48)
														return int32(1)
													}
												}
											}
										} else {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v23)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
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
								v46 = *(*int32)(unsafe.Add(mBase, _c_F_CheckFunctionValidatorAccess[0]))
								v48 = F_object_aclcheck(m, int32(1255), l1, v46, int64(128))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									if v48 != 0 {
										F_aclcheck_error(m, v48, int32(19), v21+int32(4))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v14)
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
												return int32(0)
											} else {
												F_ReleaseCatCache(m, v23)
												mBase = m.M
												v58 = m.ExcPending
												if v58 != 0 {
													return int32(0)
												} else {
													m.G0 = v10 + int32(48)
													return int32(1)
												}
											}
										}
									} else {
										F_ReleaseCatCache(m, v14)
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
											return int32(0)
										} else {
											F_ReleaseCatCache(m, v23)
											mBase = m.M
											v58 = m.ExcPending
											if v58 != 0 {
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
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					F_errmsg(m, int32(_a_F_CheckFunctionValidatorAccess_4), v10)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_CheckFunctionValidatorAccess_1), int32(2127), int32(_a_F_CheckFunctionValidatorAccess_2))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
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
	v4 = F_ExecScan(m, l0, int32(758), int32(759))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_FunctionCall6Coll(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int64, l4 int64, l5 int64, l6 int64) int64 {
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
	var v43 int64
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
	v11 = v9 - int32(128)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+120)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+112)) = l6
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+104)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = l5
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+88)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = l4
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+72)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = l3
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+56)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = l2
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v8)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = l1
	v31 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+26)) = uint16(v31)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v8)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = l0
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = m.T0[v42].(func(*base.Module, int32) int64)(m, v11+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		return int64(0)
	} else {
		v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)))
		if v47 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int64(0)
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v54
				F_errmsg_internal(m, int32(_a_F_FunctionCall6Coll_0), v11)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_FunctionCall6Coll_1), int32(1280), int32(_a_F_FunctionCall6Coll_2))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			m.G0 = v11 + int32(128)
			return v43
		}
	}
}
func F_SendFunctionCall(m *base.Module, l0 int32, l1 int64) int32 {
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
	var v21 int64
	_ = v21
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
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v3 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v8 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+12)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v6)+17)) = v8
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+40)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = l1
	v16 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v6)+26)) = uint16(v16)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = m.T0[v20].(func(*base.Module, int32) int64)(m, v6+int32(8))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+24)))
		if v25 != int32(1) {
			v28 = base.I32_wrap_i64(v21)
			v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
			if v29&int32(3) != 0 {
				v32 = F_detoast_attr(m, v28)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v32
					m.G0 = v6 + int32(48)
					return v34
				}
			} else {
				v34 = v28
				m.G0 = v6 + int32(48)
				return v34
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v43
				F_errmsg_internal(m, int32(_a_F_SendFunctionCall_0), v6)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_SendFunctionCall_1), int32(1145), int32(_a_F_SendFunctionCall_2))
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
	}
}
func F_add_function_cost(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v23 int64
	_ = v23
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v49 float32
	_ = v49
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v15 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+22)))
			v19 = v17 + v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+92))
			if v20 == int32(0) {
				v49 = *(*float32)(unsafe.Add(mBase, uint32(v19)+80))
				v52 = *(*float64)(unsafe.Add(mBase, _c_F_add_function_cost[0]))
				v54 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
				*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_promote_f32(v49), v52), v54)
				F_ReleaseCatCache(m, v15)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					m.G0 = v11 + int32(48)
					return
				}
			} else {
				v23 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v23
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(467)
				*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v23
				v34 = v11 + int32(16)
				v36 = F_OidFunctionCall1Coll(m, v20, int32(0), base.I64_extend_i32_u(v34))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					if base.I32_wrap_i64(v36) != v34 {
						v49 = *(*float32)(unsafe.Add(mBase, uint32(v19)+80))
						v52 = *(*float64)(unsafe.Add(mBase, _c_F_add_function_cost[0]))
						v54 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_promote_f32(v49), v52), v54)
					} else {
						v40 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
						v41 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v40, v41)
						v44 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
						v45 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(v44, v45)
					}
					F_ReleaseCatCache(m, v15)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
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
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
				F_errmsg_internal(m, int32(_a_F_add_function_cost_0), v11)
				mBase = m.M
				v70 = m.ExcPending
				if v70 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_add_function_cost_1), int32(2343), int32(_a_F_add_function_cost_2))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
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
						F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1826), int32(_a_F_check_transform_function_2))
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
							F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1830), int32(_a_F_check_transform_function_2))
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
								F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1834), int32(_a_F_check_transform_function_2))
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
									F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1839), int32(_a_F_check_transform_function_2))
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
					F_errfinish(m, int32(_a_F_check_transform_function_1), int32(1822), int32(_a_F_check_transform_function_2))
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
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
		v30 = base.B2i32(v27 == int32(448))
	} else {
		v30 = v7
	}
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v30)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v32 != 0 {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
		v36 = base.B2i32(v33 == int32(447))
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
		if base.Ui32(int32(101)) <= base.Ui32(v50) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return
			} else {
				F_errcode(m, int32(50856197))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return
				} else {
					v98 = int32(100)
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = v98
					F_errmsg_plural(m, int32(_a_F_compute_function_hashkey_0), int32(_a_F_compute_function_hashkey_1), v98, v13)
					mBase = m.M
					v104 = m.ExcPending
					if v104 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_compute_function_hashkey_2), int32(310), int32(_a_F_compute_function_hashkey_3))
						mBase = m.M
						v109 = m.ExcPending
						if v109 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v50
			v57 = l2 + int32(28)
			v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+104)))
			v60 = v58 << (uint(int32(2)) % 32)
			if v60 != 0 {
				base.MemoryCopy(m, v57, l1+int32(136), v60)
			} else {
			}
			v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+104)))
			v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
			F_cfunc_resolve_polymorphic_argtypes(m, v64, v57, int32(0), v67, l5, l1+int32(4))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return
			} else {
				if l4 == int32(0) {
					m.G0 = v13 + int32(16)
					return
				} else {
					v80 = F_get_call_result_type(m, l0, v13+int32(12), v13+int32(8))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return
					} else {
						v82 = int32(1)
						if base.Ui32(v82) < base.Ui32(v80-v82) {
						} else {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v86
						}
						m.G0 = v13 + int32(16)
						return
					}
				}
			}
		}
	} else {
		if l4 == int32(0) {
			m.G0 = v13 + int32(16)
			return
		} else {
			v80 = F_get_call_result_type(m, l0, v13+int32(12), v13+int32(8))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				v82 = int32(1)
				if base.Ui32(v82) < base.Ui32(v80-v82) {
				} else {
					v86 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v86
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v206 int64
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v324 int64
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v535 int32
	_ = v535
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v619 int32
	_ = v619
	v5 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(400)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
	v25 = v23 + v24
	v27 = v25 + int32(136)
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+104)))
	if l1 == v5 {
		v54 = v28
		v55 = v27
		goto L9
	} else {
		goto L10
	}
L1:
	;
	m.G0 = v21 + int32(400)
	return v619
L2:
	;
	v499 = m.G0
	v501 = v499 - int32(800)
	m.G0 = v501
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v503)+22)))
	v505 = int32(0)
	v506 = int32(400)
	base.MemoryFill(m, v501+v506, v505, v506)
	if v485 == v505 {
		v568 = v505
		goto L110
	} else {
		goto L111
	}
L3:
	;
	if v54 <= v56 {
		goto L78
	} else {
		goto L79
	}
L4:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v64+v278<<(uint(int32(2))%32))))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)))
	if v292 == int32(16) {
		goto L75
	} else {
		goto L76
	}
L5:
	;
	if v56&int32(1) == int32(0) {
		goto L3
	} else {
		goto L74
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L11
	} else {
		goto L71
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L11
	} else {
		goto L68
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
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
	v33 = F_SysCacheGetAttr(m, int32(47), l3, int32(21), v21)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v37 != 0 {
		v54 = v28
		v55 = v27
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v39 = F_pg_detoast_datum(m, base.I32_wrap_i64(v33))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v41 != int32(1) {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v44 < int32(0) {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v47 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v48 != int32(26) {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v54 = v44
	v55 = v39 + int32(24)
	goto L9
L19:
	;
	v206 = F_SysCacheGetAttrNotNull(m, int32(47), l3, int32(24))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L11
	} else {
		goto L52
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v56 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v183 = int32(0)
	if v54 <= v183 {
		v619 = v183
		goto L1
	} else {
		goto L51
	}
L23:
	;
	if v56 < v54 {
		v190 = v56
		goto L19
	} else {
		goto L50
	}
L24:
	;
	v59 = int32(0)
	if v59 < v56 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v63 = v56
	goto L27
L26:
	;
	v63 = v59
	goto L27
L27:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v66 = v59
	goto L28
L28:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v64+v66<<(uint(int32(2))%32))))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)))
	if v87 != int32(16) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if base.Ui32(int32(101)) <= base.Ui32(v54) {
		goto L7
	} else {
		goto L34
	}
L30:
	;
	v91 = v66 + int32(1)
	if v63 != v91 {
		v66 = v91
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
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+22)))
	v98 = v54 << (uint(int32(2)) % 32)
	if v98 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	base.MemoryFill(m, v21, int32(0), v98)
	goto L37
L36:
	;
	goto L37
L37:
	;
	v101 = int32(0)
	if v56 == int32(1) {
		v270 = v101
		v278 = v5
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v108 = v101
	v116 = v5
	v124 = v5
	goto L39
L39:
	;
	v128 = v64 + v116<<(uint(int32(2))%32)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	if v130 != int32(16) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L5
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+v138<<(uint(int32(2))%32)))) = v137
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if v145 != int32(16) {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	v137 = v129
	v138 = v108
	v139 = v108 + int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v137 = v136
	v138 = v135
	v139 = v108
	goto L41
L45:
	;
	v155 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v21+v153<<(uint(v155)%32)))) = v152
	v160 = v116 + v155
	v162 = v124 + v155
	if v56&int32(2147483646) != v162 {
		v108 = v154
		v116 = v160
		v124 = v162
		goto L39
	} else {
		goto L49
	}
L46:
	;
	v152 = v144
	v153 = v139
	v154 = v139 + int32(1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v144)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	v152 = v151
	v153 = v150
	v154 = v139
	goto L45
L49:
	;
	goto L40
L50:
	;
	v619 = l0
	goto L1
L51:
	;
	v190 = v183
	goto L19
L52:
	;
	v209 = F_text_to_cstring(m, base.I32_wrap_i64(v206))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v211 = F_stringToNode(m, v209)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	F_pfree(m, v209)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	if v211 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)+4))
	v217 = v215
	goto L58
L57:
	;
	v217 = int32(0)
	goto L58
L58:
	;
	v218 = v217 + v190
	v219 = v218 - v54
	if v219 < int32(0) {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	if v218 != v54 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v223 = F_list_delete_first_n(m, v211, v219)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L63
	}
L61:
	;
	v225 = v211
	goto L62
L62:
	;
	v226 = F_list_concat_copy(m, l0, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L11
	} else {
		goto L64
	}
L63:
	;
	v225 = v223
	goto L62
L64:
	;
	v485 = v226
	goto L2
L65:
	;
	F_errmsg_internal(m, int32(_a_F_expand_function_arguments_0), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_2), int32(_a_F_expand_function_arguments_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
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
	v249 = m.ExcPending
	if v249 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_5), int32(_a_F_expand_function_arguments_6))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
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
	v262 = m.ExcPending
	if v262 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_8), int32(_a_F_expand_function_arguments_9))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	v270 = v154
	v278 = v160
	goto L4
L75:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v291)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v291)+4))
	v297 = v295
	v298 = v296
	goto L77
L76:
	;
	v297 = v270
	v298 = v291
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+v297<<(uint(int32(2))%32)))) = v298
	goto L3
L78:
	;
	if v54 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L79:
	;
	v324 = F_SysCacheGetAttrNotNull(m, int32(47), l3, int32(24))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	v327 = F_text_to_cstring(m, base.I32_wrap_i64(v324))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	v329 = F_stringToNode(m, v327)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	F_pfree(m, v327)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	if v329 == int32(0) {
		goto L78
	} else {
		goto L84
	}
L84:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v335 <= int32(0) {
		goto L78
	} else {
		goto L85
	}
L85:
	;
	v339 = int32(*(*int16)(unsafe.Add(mBase, uint32(v95+v96)+106)))
	v340 = v54 - v339
	v341 = int32(0)
	if v335 != int32(1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v350 = v340
	v353 = v341
	v357 = int32(0)
	goto L89
L87:
	;
	v400 = v340
	v403 = v341
	goto L88
L88:
	;
	v419 = v21 + v400<<(uint(int32(2))%32)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v419)))
	if v420 != 0 {
		goto L78
	} else {
		goto L99
	}
L89:
	;
	v369 = v21 + v350<<(uint(int32(2))%32)
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	if v370 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	if v335&int32(1) == int32(0) {
		goto L78
	} else {
		goto L98
	}
L91:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v373+v353<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v369))) = v377
	goto L93
L92:
	;
	goto L93
L93:
	;
	v380 = v369 + int32(4)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	if v381 == int32(0) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384+v353<<(uint(int32(2))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v380))) = v388
	goto L96
L95:
	;
	goto L96
L96:
	;
	v390 = int32(2)
	v391 = v353 + v390
	v393 = v350 + v390
	v395 = v357 + v390
	if v395 != v335&int32(2147483646) {
		v350 = v393
		v353 = v391
		v357 = v395
		goto L89
	} else {
		goto L97
	}
L97:
	;
	goto L90
L98:
	;
	v400 = v393
	v403 = v391
	goto L88
L99:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v421+v403<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v419))) = v425
	goto L78
L100:
	;
	v485 = int32(0)
	goto L2
L101:
	;
	goto L102
L102:
	;
	v448 = int32(1)
	if v54 <= v448 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v451 = v448
	goto L105
L104:
	;
	v451 = v54
	goto L105
L105:
	;
	v452 = int32(0)
	v455 = v452
	v458 = v452
	goto L106
L106:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v21+v455<<(uint(int32(2))%32))))
	v476 = F_lappend(m, v458, v475)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L11
	} else {
		goto L108
	}
L107:
	;
	v485 = v476
	goto L2
L108:
	;
	v479 = v455 + int32(1)
	if v479 != v451 {
		v455 = v479
		v458 = v476
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v584 = v54 << (uint(int32(2)) % 32)
	if v584 != 0 {
		goto L123
	} else {
		goto L124
	}
L111:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	if int32(101) <= v513 {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L11
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	if v529 <= int32(0) {
		v568 = v505
		goto L110
	} else {
		goto L118
	}
L115:
	;
	F_errmsg_internal(m, int32(_a_F_expand_function_arguments_4), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_10), int32(_a_F_expand_function_arguments_11))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v535 = v505
	goto L119
L119:
	;
	v551 = v535 << (uint(int32(2)) % 32)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v485)+12))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v555+v551)))
	v558 = F_exprType(m, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L11
	} else {
		goto L121
	}
L120:
	;
	v568 = v562
	goto L110
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v551+(v501+int32(400))))) = v558
	v562 = v535 + int32(1)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v485)+4))
	if v562 < v563 {
		v535 = v562
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	base.MemoryCopy(m, v501, v55, v584)
	goto L125
L124:
	;
	goto L125
L125:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v503+v504)+108))
	v591 = F_enforce_generic_type_consistency(m, v501+int32(400), v501, v568, v589, int32(0))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L11
	} else {
		goto L126
	}
L126:
	;
	if v591 != l2 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L11
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	F_make_fn_arguments(m, int32(0), v485, v501+int32(400), v501)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L11
	} else {
		goto L133
	}
L130:
	;
	F_errmsg_internal(m, int32(_a_F_expand_function_arguments_12), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_expand_function_arguments_1), int32(_a_F_expand_function_arguments_13), int32(_a_F_expand_function_arguments_11))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	m.G0 = v501 + int32(800)
	v619 = v485
	goto L1
}
func F_simplify_function(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v336 int64
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v451 int64
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v476 int64
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v604 int32
	_ = v604
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v640 int32
	_ = v640
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v760 int32
	_ = v760
	var v769 int32
	_ = v769
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 float64
	_ = v816
	var v817 float64
	_ = v817
	var v820 float64
	_ = v820
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v922 int32
	_ = v922
	var v928 int32
	_ = v928
	var v948 int32
	_ = v948
	var v955 int32
	_ = v955
	var v975 int32
	_ = v975
	v7 = l6
	v11 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(112)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v32 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	if base.B2i32(l8 == int32(0))|v340 != 0 {
		v955 = v340
		goto L58
	} else {
		goto L59
	}
L2:
	;
	if v301|base.B2i32(l8 == int32(0)) != 0 {
		v340 = v301
		goto L1
	} else {
		goto L53
	}
L3:
	;
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+101)))
	if v246 != int32(105) {
		goto L46
	} else {
		goto L47
	}
L4:
	;
	v209 = int32(1)
	v211 = int32(0)
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+99)))
	if base.B2i32(v192&v209 == v211)|base.B2i32(v213 != v209) == v211 {
		goto L41
	} else {
		goto L42
	}
L5:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v64+v152<<(uint(int32(2))%32))))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	if v175 != int32(7) {
		goto L38
	} else {
		goto L39
	}
L6:
	;
	if v61&int32(1) == int32(0) {
		v192 = v125
		v201 = v126
		goto L4
	} else {
		goto L37
	}
L7:
	;
	return int32(0)
L8:
	;
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+22)))
	v38 = int32(0)
	if l7 == v38 {
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
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L34
	}
L12:
	;
	v53 = v52 + v50
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+100)))
	if v54 != 0 {
		v301 = v38
		goto L2
	} else {
		goto L18
	}
L13:
	;
	v50 = v36
	v51 = v29
	v52 = v37
	goto L12
L14:
	;
	goto L15
L15:
	;
	v42 = F_expand_function_arguments(m, v29, int32(0), l1, v32)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	v45 = F_expression_tree_mutator_impl(m, v42, int32(919), l9)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v45
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+22)))
	v50 = v48
	v51 = v45
	v52 = v49
	goto L12
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+108))
	if v56 == int32(2249) {
		v301 = int32(0)
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v51 == int32(0) {
		goto L3
	} else {
		goto L20
	}
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v61 <= int32(0) {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v65 = int32(0)
	if v61 == int32(1) {
		v152 = v65
		v154 = v65
		v163 = v11
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v79 = v65
	v81 = v65
	v87 = int32(0)
	v90 = v11
	goto L23
L23:
	;
	v100 = v64 + v79<<(uint(int32(2))%32)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	if v102 != int32(7) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L6
L25:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v115 != int32(7) {
		goto L30
	} else {
		goto L31
	}
L26:
	;
	v112 = v81
	v113 = int32(1)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+32)))
	v112 = base.B2i32(v106|v81&int32(1) != int32(0))
	v113 = v90
	goto L25
L29:
	;
	v127 = int32(2)
	v128 = v79 + v127
	v130 = v87 + v127
	if v61&int32(2147483646) != v130 {
		v79 = v128
		v81 = v125
		v87 = v130
		v90 = v126
		goto L23
	} else {
		goto L33
	}
L30:
	;
	v125 = v112
	v126 = int32(1)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+32)))
	v125 = base.B2i32(v119|v112&int32(1) != int32(0))
	v126 = v113
	goto L29
L33:
	;
	goto L24
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = l0
	F_errmsg_internal(m, int32(_a_F_simplify_function_0), v27)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_simplify_function_1), int32(_a_F_simplify_function_2), int32(_a_F_simplify_function_3))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
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
	v152 = v128
	v154 = v125
	v163 = v126
	goto L5
L38:
	;
	v192 = v154
	v201 = int32(1)
	goto L4
L39:
	;
	goto L40
L40:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+32)))
	v192 = base.B2i32(v179|v154&int32(1) != int32(0))
	v201 = v163
	goto L4
L41:
	;
	v219 = F_makeNullConst(m, l1, l2, l3)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v201 != 0 {
		v301 = int32(0)
		goto L2
	} else {
		goto L45
	}
L44:
	;
	v301 = v219
	goto L2
L45:
	;
	goto L3
L46:
	;
	if v246 != int32(115) {
		v301 = int32(0)
		goto L2
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v259 = F_palloc0(m, int32(36))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L7
	} else {
		goto L51
	}
L49:
	;
	v252 = int32(0)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l9)+16)))
	if v253&int32(1) == v252 {
		v301 = v252
		goto L2
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v259)+28)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v259)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v259)+20)) = l3
	v266 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v259)+16)) = v266
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+13)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v259)+12)) = uint8(v266)
	*(*int32)(unsafe.Add(mBase, uint32(v259)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v259)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = int32(15)
	v275 = F_evaluate_expr(m, v259, l1, l2, l3)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	v301 = v275
	goto L2
L53:
	;
	v305 = v36 + v37
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)+92))
	if v306 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v340 = int32(0)
	goto L1
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = int32(15)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305)+100)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+64)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+56)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v27)+52)) = l3
	v320 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v320
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+45)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+44)) = uint8(v314)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = int32(463)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+92)) = v326
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v27 + int32(32)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v305)+92))
	v336 = F_OidFunctionCall1Coll(m, v331, v320, base.I64_extend_i32_u(v27+int32(88)))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v340 = base.I32_wrap_i64(v336)
	goto L1
L58:
	;
	F_ReleaseCatCache(m, v32)
	mBase = m.M
	v975 = m.ExcPending
	if v975 != 0 {
		goto L7
	} else {
		goto L225
	}
L59:
	;
	v344 = int32(0)
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v345)+22)))
	v347 = v345 + v346
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v347)+76))
	if v348 != int32(14) {
		v955 = v344
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+96)))
	if v351 != int32(102) {
		v955 = v344
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+97)))
	if v354 != 0 {
		v955 = v344
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+100)))
	if v355 != 0 {
		v955 = v344
		goto L58
	} else {
		goto L63
	}
L63:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v347)+108))
	if v356 == int32(2249) {
		v955 = v344
		goto L58
	} else {
		goto L64
	}
L64:
	;
	v361 = F_heap_attisnull(m, v32, int32(29), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	if v361 == int32(0) {
		v955 = v344
		goto L58
	} else {
		goto L66
	}
L66:
	;
	v365 = int32(*(*int16)(unsafe.Add(mBase, uint32(v347)+104)))
	if v51 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	v368 = v366
	goto L69
L68:
	;
	v368 = int32(0)
	goto L69
L69:
	;
	if v368 != v365 {
		v955 = v344
		goto L58
	} else {
		goto L70
	}
L70:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v371 = int32(0)
	if v370 == v371 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v409 != 0 {
		v955 = v344
		goto L58
	} else {
		goto L84
	}
L72:
	;
	v409 = int32(0)
	goto L71
L73:
	;
	goto L74
L74:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v370)+4))
	if v377 <= int32(0) {
		v403 = v371
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v409 = v403
	goto L71
L76:
	;
	v380 = int32(0)
	if v380 < v377 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v383 = v377
	goto L79
L78:
	;
	v383 = v380
	goto L79
L79:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v370)+12))
	v386 = int32(0)
	goto L80
L80:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v384+v386<<(uint(int32(2))%32))))
	v395 = base.B2i32(v394 == l0)
	if v394 == l0 {
		v403 = v395
		goto L75
	} else {
		goto L82
	}
L81:
	;
	v403 = v395
	goto L75
L82:
	;
	v397 = v386 + int32(1)
	if v397 != v383 {
		v386 = v397
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	v412 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[0]))
	v414 = F_object_aclcheck(m, int32(1255), l0, v412, int64(128))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L7
	} else {
		goto L85
	}
L85:
	;
	if v414 != 0 {
		v955 = v344
		goto L58
	} else {
		goto L86
	}
L86:
	;
	v417 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[1]))
	if v417 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v418 = m.T0[v417].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L7
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2]))
	v426 = F_AllocSetContextCreateInternal(m, v421, int32(_a_F_simplify_function_4), int32(0), int32(_a_F_simplify_function_5), int32(_a_F_simplify_function_6))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L7
	} else {
		goto L92
	}
L90:
	;
	if v418 != 0 {
		v955 = v344
		goto L58
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v428 = int32(_a_F_simplify_function_7)
	v429 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2])) = v426
	v433 = F_palloc0(m, int32(36))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L7
	} else {
		goto L93
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v433)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+28)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v433)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v433)+20)) = l3
	v440 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+16)) = v440
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+13)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v433)+12)) = uint8(v440)
	*(*int32)(unsafe.Add(mBase, uint32(v433)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v433)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v433))) = int32(15)
	v451 = F_SysCacheGetAttrNotNull(m, int32(47), v32, int32(26))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L7
	} else {
		goto L94
	}
L94:
	;
	v454 = F_text_to_cstring(m, base.I32_wrap_i64(v451))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L7
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+104)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v347 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+92)) = int32(921)
	v462 = int32(_a_F_simplify_function_8)
	v463 = *(*int32)(unsafe.Add(mBase, _c_F_simplify_function[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[3])) = v27 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = v463
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v27 + int32(100)
	v476 = F_SysCacheGetAttr(m, int32(47), v32, int32(28), v27+int32(111))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+111)))
	if v478 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L97:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[3])) = v948
	v955 = v928
	goto L58
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2])) = v429
	F_MemoryContextDelete(m, v426)
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L7
	} else {
		goto L224
	}
L99:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v640)))
	if v656 != int32(67) {
		goto L98
	} else {
		goto L134
	}
L100:
	;
	v482 = F_text_to_cstring(m, base.I32_wrap_i64(v476))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L7
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v506 = F_prepare_sql_fn_parse_info(m, v32, v433, l4)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L7
	} else {
		goto L112
	}
L103:
	;
	if v498 == int32(0) {
		goto L98
	} else {
		goto L110
	}
L104:
	;
	v484 = F_stringToNode(m, v482)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L7
	} else {
		goto L105
	}
L105:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v484)))
	if v486 == int32(1) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v484)+12))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	v498 = v490
	goto L103
L107:
	;
	goto L108
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v484
	*(*int32)(unsafe.Add(mBase, uint32(v27)+80)) = v484
	v496 = F_list_make1_impl(m, int32(1), v27+int32(28))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	v498 = v496
	goto L103
L110:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v501 != int32(1) {
		goto L98
	} else {
		goto L111
	}
L111:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v504)))
	v640 = v505
	goto L99
L112:
	;
	v508 = F_pg_parse_query(m, v454)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	if v508 == int32(0) {
		goto L98
	} else {
		goto L114
	}
L114:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508)+4))
	if v512 != int32(1) {
		goto L98
	} else {
		goto L115
	}
L115:
	;
	v516 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L7
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v516)+4)) = v454
	*(*int32)(unsafe.Add(mBase, uint32(v516)+116)) = v506
	*(*int32)(unsafe.Add(mBase, uint32(v516)+108)) = int32(725)
	*(*int32)(unsafe.Add(mBase, uint32(v516)+104)) = int32(726)
	*(*int32)(unsafe.Add(mBase, uint32(v516)+100)) = int32(0)
	goto L117
L117:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v508)+12))
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v526)))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v527)+4))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)))
	if v529 != int32(141) {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v624 = F_transformStmt(m, v516, v604)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L7
	} else {
		goto L132
	}
L119:
	;
	v604 = v528
	goto L118
L120:
	;
	goto L121
L121:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528)+68))
	if v532 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v539 = v528
	goto L125
L123:
	;
	v565 = v528
	goto L124
L124:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v565)+8))
	if v583 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v539)+76))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v557)+68))
	if v558 != 0 {
		v539 = v557
		goto L125
	} else {
		goto L127
	}
L126:
	;
	v565 = v557
	goto L124
L127:
	;
	goto L126
L128:
	;
	v604 = v528
	goto L118
L129:
	;
	goto L130
L130:
	;
	v587 = F_palloc0(m, int32(20))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587)+4)) = v528
	*(*int32)(unsafe.Add(mBase, uint32(v587))) = int32(242)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v565)+8))
	v593 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v587)+16)) = uint8(v593)
	*(*int32)(unsafe.Add(mBase, uint32(v587)+12)) = int32(42)
	*(*int32)(unsafe.Add(mBase, uint32(v587)+8)) = v592
	*(*int32)(unsafe.Add(mBase, uint32(v565)+8)) = int32(0)
	v604 = v587
	goto L118
L132:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v527)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v624)+156)) = v626
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v527)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v624)+160)) = v628
	F_free_parsestate(m, v516)
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L7
	} else {
		goto L133
	}
L133:
	;
	v640 = v624
	goto L99
L134:
	;
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	if v659 != int32(1) {
		goto L98
	} else {
		goto L135
	}
L135:
	;
	v662 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+36)))
	if v662 != 0 {
		goto L98
	} else {
		goto L136
	}
L136:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+37)))
	if v663 != 0 {
		goto L98
	} else {
		goto L137
	}
L137:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+38)))
	if v664 != 0 {
		goto L98
	} else {
		goto L138
	}
L138:
	;
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+39)))
	if v665 != 0 {
		goto L98
	} else {
		goto L139
	}
L139:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v640)+48))
	if v666 != 0 {
		goto L98
	} else {
		goto L140
	}
L140:
	;
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v640)+52))
	if v667 != 0 {
		goto L98
	} else {
		goto L141
	}
L141:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v640)+60))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)+4))
	if v669 != 0 {
		goto L98
	} else {
		goto L142
	}
L142:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v668)+8))
	if v670 != 0 {
		goto L98
	} else {
		goto L143
	}
L143:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v640)+100))
	if v671 != 0 {
		goto L98
	} else {
		goto L144
	}
L144:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v640)+108))
	if v672 != 0 {
		goto L98
	} else {
		goto L145
	}
L145:
	;
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v640)+112))
	if v673 != 0 {
		goto L98
	} else {
		goto L146
	}
L146:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v640)+116))
	if v674 != 0 {
		goto L98
	} else {
		goto L147
	}
L147:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v640)+120))
	if v675 != 0 {
		goto L98
	} else {
		goto L148
	}
L148:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v640)+124))
	if v676 != 0 {
		goto L98
	} else {
		goto L149
	}
L149:
	;
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v640)+128))
	if v677 != 0 {
		goto L98
	} else {
		goto L150
	}
L150:
	;
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v640)+132))
	if v678 != 0 {
		goto L98
	} else {
		goto L151
	}
L151:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v640)+144))
	if v679 != 0 {
		goto L98
	} else {
		goto L152
	}
L152:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v640)+76))
	if v680 == int32(0) {
		goto L98
	} else {
		goto L153
	}
L153:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	if v683 != int32(1) {
		goto L98
	} else {
		goto L154
	}
L154:
	;
	v689 = F_get_expr_result_type(m, v433, int32(0), v27+int32(84))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L7
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+24)) = v640
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = v640
	v696 = F_list_make1_impl(m, int32(1), v27+int32(24))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L7
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v27)+72)) = v696
	v703 = F_list_make1_impl(m, int32(1), v27+int32(20))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L7
	} else {
		goto L157
	}
L157:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v27)+84))
	v706 = int32(*(*int8)(unsafe.Add(mBase, uint32(v347)+96)))
	v708 = F_check_sql_fn_retval(m, v703, l1, v705, v706, int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L7
	} else {
		goto L158
	}
L158:
	;
	if v708 != 0 {
		goto L98
	} else {
		goto L159
	}
L159:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v696)+12))
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	if v640 != v711 {
		goto L98
	} else {
		goto L160
	}
L160:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v640)+76))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v713)+12))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v715)+4))
	v717 = F_exprType(m, v716)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L7
	} else {
		goto L161
	}
L161:
	;
	if v717 != l1 {
		goto L98
	} else {
		goto L162
	}
L162:
	;
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+101)))
	if v720 == int32(105) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v724 = F_contain_mutable_functions_walker(m, v716, int32(0))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L7
	} else {
		goto L166
	}
L164:
	;
	v727 = v720
	goto L165
L165:
	;
	if v727&int32(255) == int32(115) {
		goto L168
	} else {
		goto L169
	}
L166:
	;
	if v724 != 0 {
		goto L98
	} else {
		goto L167
	}
L167:
	;
	v726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+101)))
	v727 = v726
	goto L165
L168:
	;
	v733 = F_contain_volatile_functions_walker(m, v716, int32(0))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L7
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+99)))
	if v735 == int32(1) {
		goto L173
	} else {
		goto L174
	}
L171:
	;
	if v733 != 0 {
		goto L98
	} else {
		goto L172
	}
L172:
	;
	goto L170
L173:
	;
	v739 = F_contain_nonstrict_functions_walker(m, v716, int32(0))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L7
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = int32(0)
	v744 = v27 + int32(32)
	v745 = F_contain_context_dependent_node_walker(m, v51, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L7
	} else {
		goto L178
	}
L176:
	;
	if v739 != 0 {
		goto L98
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	if v745 != 0 {
		goto L98
	} else {
		goto L179
	}
L179:
	;
	v747 = int32(*(*int16)(unsafe.Add(mBase, uint32(v347)+104)))
	v750 = F_palloc0(m, v747<<(uint(int32(2))%32))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L7
	} else {
		goto L180
	}
L180:
	;
	v752 = int32(*(*int16)(unsafe.Add(mBase, uint32(v347)+104)))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v752
	v756 = F_substitute_actual_parameters_mutator(m, v716, v744)
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L7
	} else {
		goto L181
	}
L181:
	;
	if v51 == int32(0) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_simplify_function[2])) = v429
	v859 = F_copyObjectImpl(m, v756)
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L7
	} else {
		goto L203
	}
L183:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v760 <= int32(0) {
		goto L182
	} else {
		goto L184
	}
L184:
	;
	v769 = int32(0)
	goto L185
L185:
	;
	v789 = v769 << (uint(int32(2)) % 32)
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v750+v789)))
	switch v791 {
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
	v830 = v769 + int32(1)
	v831 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v830 < v831 {
		v769 = v830
		goto L185
	} else {
		goto L202
	}
L188:
	;
	v827 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347)+99)))
	if v827 != 0 {
		goto L98
	} else {
		goto L201
	}
L189:
	;
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v792+v789)))
	if v794 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)))
	if base.Ui32(v795-int32(22)) < base.Ui32(int32(3)) {
		goto L98
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v794
	*(*int32)(unsafe.Add(mBase, uint32(v27)+68)) = v794
	v811 = F_list_make1_impl(m, int32(1), v27+int32(16))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L7
	} else {
		goto L196
	}
L193:
	;
	v802 = F_expression_tree_walker_impl(m, v794, int32(905), int32(0))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L7
	} else {
		goto L194
	}
L194:
	;
	if v802 != 0 {
		goto L98
	} else {
		goto L195
	}
L195:
	;
	goto L192
L196:
	;
	F_cost_qual_eval(m, v27+int32(32), v811, int32(0))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L7
	} else {
		goto L197
	}
L197:
	;
	v816 = *(*float64)(unsafe.Add(mBase, uint32(v27)+32))
	v817 = *(*float64)(unsafe.Add(mBase, uint32(v27)+40))
	v820 = *(*float64)(unsafe.Add(mBase, _c_F_simplify_function[4]))
	if base.F64_gt(base.F64_add(v816, v817), base.F64_mul(v820, float64(10))) != 0 {
		goto L98
	} else {
		goto L198
	}
L198:
	;
	v825 = F_contain_volatile_functions_walker(m, v794, int32(0))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L7
	} else {
		goto L199
	}
L199:
	;
	if v825 != 0 {
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
	F_MemoryContextDelete(m, v426)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
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
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v881 != 0 {
		goto L217
	} else {
		goto L218
	}
L206:
	;
	v880 = v859
	goto L205
L207:
	;
	goto L208
L208:
	;
	v865 = F_exprCollation(m, v859)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L7
	} else {
		goto L209
	}
L209:
	;
	if v865 == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v880 = v859
	goto L205
L211:
	;
	goto L212
L212:
	;
	if v865 == l3 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v880 = v859
	goto L205
L214:
	;
	goto L215
L215:
	;
	v871 = F_palloc0(m, int32(16))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L7
	} else {
		goto L216
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v871)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v871)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v871)+4)) = v859
	*(*int32)(unsafe.Add(mBase, uint32(v871))) = int32(31)
	v880 = v871
	goto L205
L217:
	;
	F_record_plan_function_dependency(m, v881, l0)
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L7
	} else {
		goto L220
	}
L218:
	;
	goto L219
L219:
	;
	v884 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v885 = F_lappend_oid(m, v884, l0)
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L7
	} else {
		goto L221
	}
L220:
	;
	goto L219
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9)+8)) = v885
	v888 = F_eval_const_expressions_mutator(m, v880, l9)
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L7
	} else {
		goto L222
	}
L222:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v891 = F_list_delete_last(m, v890)
	mBase = m.M
	v892 = m.ExcPending
	if v892 != 0 {
		goto L7
	} else {
		goto L223
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9)+8)) = v891
	v928 = v888
	goto L97
L224:
	;
	v928 = int32(0)
	goto L97
L225:
	;
	m.G0 = v27 + int32(112)
	return v955
}
