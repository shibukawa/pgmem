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
						F_errmsg_internal(m, int32(56368), v10+int32(16))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(515123), int32(2170), int32(137751))
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
								F_errmsg(m, int32(52564), v10+int32(32))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(515123), int32(2178), int32(137751))
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
						v34 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
									v45 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
								v45 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
					F_errmsg(m, int32(74548), v10)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515123), int32(2161), int32(137751))
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
				F_errmsg_internal(m, int32(553757), v11)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515123), int32(1278), int32(316127))
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
	*(*int64)(unsafe.Add(mBase, uint32(v6)+13)) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v8
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
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
				F_errmsg_internal(m, int32(553757), v6)
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515123), int32(1143), int32(316247))
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v47 float32
	_ = v47
	var v50 float64
	_ = v50
	var v52 float64
	_ = v52
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
				v47 = *(*float32)(unsafe.Add(mBase, uint32(v18)+80))
				v50 = *(*float64)(unsafe.Add(mBase, _consts[382]))
				v52 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
				*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_promote_f32(v47), v50), v52)
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
				*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v22
				*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(459)
				v34 = F_OidFunctionCall1Coll(m, v19, int32(0), v11+int32(16))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					if v34 != v11+int32(16) {
						v47 = *(*float32)(unsafe.Add(mBase, uint32(v18)+80))
						v50 = *(*float64)(unsafe.Add(mBase, _consts[382]))
						v52 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(base.F64_mul(base.F64_promote_f32(v47), v50), v52)
					} else {
						v39 = *(*float64)(unsafe.Add(mBase, uint32(v11)+32))
						v40 = *(*float64)(unsafe.Add(mBase, uint32(l3)))
						*(*float64)(unsafe.Add(mBase, uint32(l3))) = base.F64_add(v39, v40)
						v43 = *(*float64)(unsafe.Add(mBase, uint32(v11)+40))
						v44 = *(*float64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*float64)(unsafe.Add(mBase, uint32(l3)+8)) = base.F64_add(v43, v44)
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
				F_errmsg_internal(m, int32(47929), v11)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errfinish(m, int32(512965), int32(2133), int32(73611))
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
					F_errmsg(m, int32(263307), int32(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errfinish(m, int32(514176), int32(1811), int32(261825))
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
						F_errmsg(m, int32(113481), int32(0))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
							return
						} else {
							F_errfinish(m, int32(514176), int32(1815), int32(261825))
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
							F_errmsg(m, int32(100173), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return
							} else {
								F_errfinish(m, int32(514176), int32(1819), int32(261825))
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
								*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(325094)
								F_errmsg(m, int32(201938), v5)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return
								} else {
									F_errfinish(m, int32(514176), int32(1824), int32(261825))
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
				F_errmsg(m, int32(401604), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_errfinish(m, int32(514176), int32(1807), int32(261825))
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
	var v15 int64
	_ = v15
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	v7 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v15
	*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v15
	*(*int64)(unsafe.Add(mBase, uint32(l2)+8)) = v15
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
	if l5 != 0 {
	} else {
		if v30^int32(1) != 0 {
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
			*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v43
		}
	}
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v45
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+104)))
	if int32(0) < v47 {
		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v47
		v52 = l2 + int32(28)
		v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+104)))
		v57 = v55 << (uint(int32(2)) % 32)
		if v57 != 0 {
			v58 = F__emscripten_memcpy_bulkmem(m, v52, l1+int32(136), v57)
			mBase = m.M
			v59 = v58
		} else {
			v59 = v52
		}
		v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+104)))
		v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+24))
		F_cfunc_resolve_polymorphic_argtypes(m, v60, v59, int32(0), v63, l5, l1+int32(4))
		mBase = m.M
		v67 = m.ExcPending
		if v67 != 0 {
			return
		} else {
			if l4 == int32(0) {
				m.G0 = v13 + int32(16)
				return
			} else {
				v75 = F_get_call_result_type(m, l0, v13+int32(12), v13+int32(8))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					v77 = int32(1)
					if base.Ui32(v77) < base.Ui32(v75-v77) {
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v81
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
			v75 = F_get_call_result_type(m, l0, v13+int32(12), v13+int32(8))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return
			} else {
				v77 = int32(1)
				if base.Ui32(v77) < base.Ui32(v75-v77) {
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v81
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
	var v20 int32
	_ = v20
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
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v609 int32
	_ = v609
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(400)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
	v26 = v24 + v25
	v28 = v26 + int32(136)
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+104)))
	if l1 == v5 {
		v54 = v29
		v55 = v28
		goto L7
	} else {
		goto L8
	}
L1:
	;
	m.G0 = v22 + int32(400)
	return v609
L2:
	;
	v491 = m.G0
	v493 = v491 - int32(800)
	m.G0 = v493
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495)+22)))
	v497 = int32(0)
	if v476 == v497 {
		v556 = v497
		goto L108
	} else {
		goto L109
	}
L3:
	;
	if v56&v104 != 0 {
		goto L70
	} else {
		goto L71
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L9
	} else {
		goto L67
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L9
	} else {
		goto L64
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L9
	} else {
		goto L61
	}
L7:
	;
	if l0 != 0 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v34 = F_SysCacheGetAttr(m, int32(47), l3, int32(21), v22)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v38 != 0 {
		v54 = v29
		v55 = v28
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v39 = F_pg_detoast_datum(m, v34)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v41 != int32(1) {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	if v44 < int32(0) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	if v47 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v48 != int32(26) {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v54 = v44
	v55 = v39 + int32(24)
	goto L7
L17:
	;
	v211 = F_SysCacheGetAttrNotNull(m, int32(47), l3, int32(24))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L9
	} else {
		goto L48
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v56 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v187 = int32(0)
	if v54 <= v187 {
		v609 = v187
		goto L1
	} else {
		goto L47
	}
L21:
	;
	if v56 < v54 {
		v194 = v56
		goto L17
	} else {
		goto L46
	}
L22:
	;
	v59 = int32(0)
	if v59 < v56 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v63 = v56
	goto L25
L24:
	;
	v63 = v59
	goto L25
L25:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v66 = v59
	goto L26
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v64+v66<<(uint(int32(2))%32))))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 != int32(16) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if base.Ui32(int32(101)) <= base.Ui32(v54) {
		goto L5
	} else {
		goto L32
	}
L28:
	;
	v92 = v66 + int32(1)
	if v63 != v92 {
		v66 = v92
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L21
L32:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
	v98 = int32(0)
	v103 = F__emscripten_memset_bulkmem(m, v22, base.I32_extend8_s(v98), v54<<(uint(int32(2))%32))
	mBase = m.M
	goto L33
L33:
	;
	v104 = int32(1)
	if v56 == v104 {
		v272 = v98
		v281 = v5
		goto L3
	} else {
		goto L34
	}
L34:
	;
	v110 = v98
	v119 = v5
	v122 = v5
	goto L35
L35:
	;
	v131 = v64 + v110<<(uint(int32(2))%32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	if v133 != int32(16) {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v272 = v163
	v281 = v157
	goto L3
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103+v141<<(uint(int32(2))%32)))) = v140
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	if v148 != int32(16) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v140 = v132
	v141 = v119
	v142 = v119 + int32(1)
	goto L37
L39:
	;
	goto L40
L40:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v140 = v139
	v141 = v138
	v142 = v119
	goto L37
L41:
	;
	v158 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v103+v156<<(uint(v158)%32)))) = v155
	v163 = v110 + v158
	v165 = v122 + v158
	if v165 != v56&int32(2147483646) {
		v110 = v163
		v119 = v157
		v122 = v165
		goto L35
	} else {
		goto L45
	}
L42:
	;
	v155 = v147
	v156 = v142
	v157 = v142 + int32(1)
	goto L41
L43:
	;
	goto L44
L44:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v155 = v154
	v156 = v153
	v157 = v142
	goto L41
L45:
	;
	goto L36
L46:
	;
	v609 = l0
	goto L1
L47:
	;
	v194 = v187
	goto L17
L48:
	;
	v213 = F_text_to_cstring(m, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L9
	} else {
		goto L49
	}
L49:
	;
	v215 = F_stringToNode(m, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	F_pfree(m, v213)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	if v215 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v221 = v219
	goto L54
L53:
	;
	v221 = int32(0)
	goto L54
L54:
	;
	v222 = v221 + v194
	v223 = v222 - v54
	if v223 < int32(0) {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	if v222 != v54 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v227 = F_list_delete_first_n(m, v215, v223)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L9
	} else {
		goto L59
	}
L57:
	;
	v229 = v215
	goto L58
L58:
	;
	v230 = F_list_concat_copy(m, l0, v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L60
	}
L59:
	;
	v229 = v227
	goto L58
L60:
	;
	v476 = v230
	goto L2
L61:
	;
	F_errmsg_internal(m, int32(160407), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L9
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(513951), int32(4212), int32(127839))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L9
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	F_errmsg_internal(m, int32(128594), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L9
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(513951), int32(4269), int32(127812))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L9
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errmsg_internal(m, int32(128048), int32(0))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L9
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(513951), int32(4341), int32(130595))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L9
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v64+v272<<(uint(int32(2))%32))))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	if v295 == int32(16) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if v54 <= v56 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v294)+4))
	v300 = v299
	v301 = v298
	goto L75
L74:
	;
	v300 = v294
	v301 = v281
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103+v301<<(uint(int32(2))%32)))) = v300
	goto L72
L76:
	;
	if v54 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L77:
	;
	v311 = F_SysCacheGetAttrNotNull(m, int32(47), l3, int32(24))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L9
	} else {
		goto L78
	}
L78:
	;
	v313 = F_text_to_cstring(m, v311)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L9
	} else {
		goto L79
	}
L79:
	;
	v315 = F_stringToNode(m, v313)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L9
	} else {
		goto L80
	}
L80:
	;
	F_pfree(m, v313)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L9
	} else {
		goto L81
	}
L81:
	;
	if v315 == int32(0) {
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	if v321 <= int32(0) {
		goto L76
	} else {
		goto L83
	}
L83:
	;
	v324 = int32(1)
	v327 = int32(*(*int16)(unsafe.Add(mBase, uint32(v96+v97)+106)))
	v328 = v54 - v327
	v329 = int32(0)
	if v321 != v324 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v336 = v328
	v339 = v329
	v342 = int32(0)
	goto L87
L85:
	;
	v386 = v328
	v389 = v329
	goto L86
L86:
	;
	if v321&v324 == int32(0) {
		goto L76
	} else {
		goto L96
	}
L87:
	;
	v355 = v336 << (uint(int32(2)) % 32)
	v356 = v103 + v355
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v356)))
	if v357 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v386 = v381
	v389 = v379
	goto L86
L89:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v360+v339<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v356))) = v364
	goto L91
L90:
	;
	goto L91
L91:
	;
	v368 = v355 + v103 + int32(4)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	if v369 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v372+v339<<(uint(int32(2))%32))+4))
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v376
	goto L94
L93:
	;
	goto L94
L94:
	;
	v378 = int32(2)
	v379 = v339 + v378
	v381 = v336 + v378
	v383 = v342 + v378
	if v383 != v321&int32(2147483646) {
		v336 = v381
		v339 = v379
		v342 = v383
		goto L87
	} else {
		goto L95
	}
L95:
	;
	goto L88
L96:
	;
	v408 = v103 + v386<<(uint(int32(2))%32)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v408)))
	if v409 != 0 {
		goto L76
	} else {
		goto L97
	}
L97:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v410+v389<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v408))) = v414
	goto L76
L98:
	;
	v476 = int32(0)
	goto L2
L99:
	;
	goto L100
L100:
	;
	v438 = int32(1)
	if v54 <= v438 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v441 = v438
	goto L103
L102:
	;
	v441 = v54
	goto L103
L103:
	;
	v442 = int32(0)
	v445 = v442
	v448 = v442
	goto L104
L104:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v103+v445<<(uint(int32(2))%32))))
	v467 = F_lappend(m, v448, v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L9
	} else {
		goto L106
	}
L105:
	;
	v476 = v467
	goto L2
L106:
	;
	v470 = v445 + int32(1)
	if v441 != v470 {
		v445 = v470
		v448 = v467
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	v573 = v54 << (uint(int32(2)) % 32)
	if v573 != 0 {
		goto L122
	} else {
		goto L123
	}
L109:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if int32(101) <= v500 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L9
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v516 <= int32(0) {
		v556 = v497
		goto L108
	} else {
		goto L116
	}
L113:
	;
	F_errmsg_internal(m, int32(128594), int32(0))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L9
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(513951), int32(4395), int32(163629))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L9
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L116:
	;
	v522 = v497
	goto L117
L117:
	;
	v539 = v522 << (uint(int32(2)) % 32)
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v476)+12))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v543+v539)))
	v546 = F_exprType(m, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L9
	} else {
		goto L119
	}
L118:
	;
	v556 = v550
	goto L108
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v539+(v493+int32(400))))) = v546
	v550 = v522 + int32(1)
	v551 = *(*int32)(unsafe.Add(mBase, uint32(v476)+4))
	if v550 < v551 {
		v522 = v550
		goto L117
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v495+v496)+108))
	v581 = F_enforce_generic_type_consistency(m, v575+int32(400), v575, v556, v579, int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L9
	} else {
		goto L125
	}
L122:
	;
	v574 = F__emscripten_memcpy_bulkmem(m, v493, v55, v573)
	mBase = m.M
	v575 = v574
	goto L124
L123:
	;
	v575 = v493
	goto L124
L124:
	;
	goto L121
L125:
	;
	if v581 != l2 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L9
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	F_make_fn_arguments(m, int32(0), v476, v575+int32(400), v575)
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L9
	} else {
		goto L132
	}
L129:
	;
	F_errmsg_internal(m, int32(349158), int32(0))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L9
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(513951), int32(4410), int32(163629))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L9
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	m.G0 = v575 + int32(800)
	v609 = v476
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
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
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
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v567 int32
	_ = v567
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v602 int32
	_ = v602
	var v617 int32
	_ = v617
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
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
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v732 int32
	_ = v732
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 float64
	_ = v778
	var v779 float64
	_ = v779
	var v782 float64
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v906 int32
	_ = v906
	var v913 int32
	_ = v913
	var v932 int32
	_ = v932
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
		goto L4
	} else {
		goto L5
	}
L1:
	;
	if l8 == int32(0) {
		v307 = v272
		goto L53
	} else {
		goto L54
	}
L2:
	;
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+101)))
	if v218 != int32(105) {
		goto L46
	} else {
		goto L47
	}
L3:
	;
	if v59&v62 == int32(0) {
		v183 = v150
		v184 = v162
		goto L36
	} else {
		goto L37
	}
L4:
	;
	return int32(0)
L5:
	;
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+22)))
	v36 = int32(0)
	if l7 == v36 {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L33
	}
L9:
	;
	v51 = v50 + v48
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+100)))
	if v52 != 0 {
		v272 = v36
		goto L1
	} else {
		goto L15
	}
L10:
	;
	v48 = v34
	v49 = v28
	v50 = v35
	goto L9
L11:
	;
	goto L12
L12:
	;
	v40 = F_expand_function_arguments(m, v28, int32(0), l1, v30)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v43 = F_expression_tree_mutator_impl(m, v40, int32(871), l9)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
	v48 = v46
	v49 = v43
	v50 = v47
	goto L9
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+108))
	if v54 == int32(2249) {
		v272 = int32(0)
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v49 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v59 <= int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	v62 = int32(1)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	if v59 == v62 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v67 = int32(0)
	v148 = v67
	v150 = v67
	v162 = v11
	goto L3
L20:
	;
	goto L21
L21:
	;
	v71 = int32(0)
	v78 = v71
	v80 = v71
	v88 = v11
	v92 = v11
	goto L22
L22:
	;
	v98 = v64 + v78<<(uint(int32(2))%32)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v100 != int32(7) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v148 = v126
	v150 = v123
	v162 = v124
	goto L3
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v113 != int32(7) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v110 = v80
	v111 = int32(1)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+24)))
	v110 = base.B2i32(v104|v80&int32(1) != int32(0))
	v111 = v92
	goto L24
L28:
	;
	v125 = int32(2)
	v126 = v78 + v125
	v128 = v88 + v125
	if v59&int32(2147483646) != v128 {
		v78 = v126
		v80 = v123
		v88 = v128
		v92 = v124
		goto L22
	} else {
		goto L32
	}
L29:
	;
	v123 = v110
	v124 = int32(1)
	goto L28
L30:
	;
	goto L31
L31:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+24)))
	v123 = base.B2i32(v117|v110&int32(1) != int32(0))
	v124 = v111
	goto L28
L32:
	;
	goto L23
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = l0
	F_errmsg_internal(m, int32(47929), v26)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(513951), int32(4086), int32(261709))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+99)))
	if v185 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v64+v148<<(uint(int32(2))%32))))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	if v172 != int32(7) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v183 = v150
	v184 = int32(1)
	goto L36
L39:
	;
	goto L40
L40:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+24)))
	v183 = base.B2i32(v176|v150&int32(1) != int32(0))
	v184 = v162
	goto L36
L41:
	;
	if v184 != 0 {
		v272 = int32(0)
		goto L1
	} else {
		goto L45
	}
L42:
	;
	if v183&int32(1) == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v192 = F_makeNullConst(m, l1, l2, l3)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v272 = v192
	goto L1
L45:
	;
	goto L2
L46:
	;
	if v218 != int32(115) {
		v272 = int32(0)
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v231 = F_palloc0(m, int32(36))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	v224 = int32(0)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l9)+16)))
	if v225&int32(1) == v224 {
		v272 = v224
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+28)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v231)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = l3
	v238 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+16)) = v238
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+13)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v231)+12)) = uint8(v238)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = int32(15)
	v247 = F_evaluate_expr(m, v231, l1, l2, l3)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v272 = v247
	goto L1
L53:
	;
	if l8 == int32(0) {
		v913 = v307
		goto L60
	} else {
		goto L61
	}
L54:
	;
	if v272 != 0 {
		v307 = v272
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v275 = v35 + v34
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+92))
	if v276 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v307 = int32(0)
	goto L53
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = int32(15)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275)+100)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = l3
	v290 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v290
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+45)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+44)) = uint8(v284)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = int32(457)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = v296
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v26 + int32(32)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v275)+92))
	v305 = F_OidFunctionCall1Coll(m, v301, v290, v26+int32(88))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v307 = v305
	goto L53
L60:
	;
	F_ReleaseCatCache(m, v30)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L4
	} else {
		goto L228
	}
L61:
	;
	if v307 != 0 {
		v913 = v307
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v311 = int32(0)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v312)+22)))
	v314 = v312 + v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)+76))
	if v315 != int32(14) {
		v913 = v311
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+96)))
	if v318 != int32(102) {
		v913 = v311
		goto L60
	} else {
		goto L64
	}
L64:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+97)))
	if v321 != 0 {
		v913 = v311
		goto L60
	} else {
		goto L65
	}
L65:
	;
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+100)))
	if v322 != 0 {
		v913 = v311
		goto L60
	} else {
		goto L66
	}
L66:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v314)+108))
	if v323 == int32(2249) {
		v913 = v311
		goto L60
	} else {
		goto L67
	}
L67:
	;
	v328 = F_heap_attisnull(m, v30, int32(29), int32(0))
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	if v328 == int32(0) {
		v913 = v311
		goto L60
	} else {
		goto L69
	}
L69:
	;
	v332 = int32(*(*int16)(unsafe.Add(mBase, uint32(v314)+104)))
	if v49 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v335 = v333
	goto L72
L71:
	;
	v335 = int32(0)
	goto L72
L72:
	;
	if v335 != v332 {
		v913 = v311
		goto L60
	} else {
		goto L73
	}
L73:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v338 = int32(0)
	if v337 == v338 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v376 != 0 {
		v913 = v311
		goto L60
	} else {
		goto L87
	}
L75:
	;
	v376 = int32(0)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v337)+4))
	if v344 <= int32(0) {
		v369 = v338
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v376 = v369
	goto L74
L79:
	;
	v347 = int32(0)
	if v347 < v344 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v350 = v344
	goto L82
L81:
	;
	v350 = v347
	goto L82
L82:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v337)+12))
	v353 = int32(0)
	goto L83
L83:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v351+v353<<(uint(int32(2))%32))))
	v362 = base.B2i32(v361 == l0)
	if v361 == l0 {
		v369 = v362
		goto L78
	} else {
		goto L85
	}
L84:
	;
	v369 = v362
	goto L78
L85:
	;
	v364 = v353 + int32(1)
	if v364 != v350 {
		v353 = v364
		goto L83
	} else {
		goto L86
	}
L86:
	;
	goto L84
L87:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v381 = F_object_aclcheck(m, int32(1255), l0, v379, int64(128))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	if v381 != 0 {
		v913 = v311
		goto L60
	} else {
		goto L89
	}
L89:
	;
	v384 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	if v384 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v385 = m.T0[v384].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v388 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v393 = F_AllocSetContextCreateInternal(m, v388, int32(261982), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L95
	}
L93:
	;
	if v385 != 0 {
		v913 = v311
		goto L60
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v395 = int32(4549024)
	v396 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v393
	v400 = F_palloc0(m, int32(36))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L4
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v400)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v400)+28)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v400)+24)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v400)+20)) = l3
	v407 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v400)+16)) = v407
	*(*uint8)(unsafe.Add(mBase, uint32(v400)+13)) = uint8(v7)
	*(*uint8)(unsafe.Add(mBase, uint32(v400)+12)) = uint8(v407)
	*(*int32)(unsafe.Add(mBase, uint32(v400)+8)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v400)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v400))) = int32(15)
	v418 = F_SysCacheGetAttrNotNull(m, int32(47), v30, int32(26))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L97
	}
L97:
	;
	v420 = F_text_to_cstring(m, v418)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v314 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+92)) = int32(873)
	v428 = int32(4541928)
	v429 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v26 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+88)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v26)+96)) = v26 + int32(100)
	v442 = F_SysCacheGetAttr(m, int32(47), v30, int32(28), v26+int32(111))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+111)))
	if v444 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v906
	v913 = v887
	goto L60
L101:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v396
	F_MemoryContextDelete(m, v393)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L4
	} else {
		goto L227
	}
L102:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v602)))
	if v617 != int32(67) {
		goto L101
	} else {
		goto L137
	}
L103:
	;
	v447 = F_text_to_cstring(m, v442)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L4
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	v471 = F_prepare_sql_fn_parse_info(m, v30, v400, l4)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L4
	} else {
		goto L115
	}
L106:
	;
	if v463 == int32(0) {
		goto L101
	} else {
		goto L113
	}
L107:
	;
	v449 = F_stringToNode(m, v447)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	if v451 == int32(1) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v449)+12))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)))
	v463 = v455
	goto L106
L110:
	;
	goto L111
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v449
	*(*int32)(unsafe.Add(mBase, uint32(v26)+80)) = v449
	v461 = F_list_make1_impl(m, int32(1), v26+int32(28))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L4
	} else {
		goto L112
	}
L112:
	;
	v463 = v461
	goto L106
L113:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v463)+4))
	if v466 != int32(1) {
		goto L101
	} else {
		goto L114
	}
L114:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v463)+12))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v469)))
	v602 = v470
	goto L102
L115:
	;
	v473 = F_pg_parse_query(m, v420)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L116
	}
L116:
	;
	if v473 == int32(0) {
		goto L101
	} else {
		goto L117
	}
L117:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473)+4))
	if v477 != int32(1) {
		goto L101
	} else {
		goto L118
	}
L118:
	;
	v481 = F_make_parsestate(m, int32(0))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v481)+4)) = v420
	*(*int32)(unsafe.Add(mBase, uint32(v481)+120)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v481)+112)) = int32(679)
	*(*int32)(unsafe.Add(mBase, uint32(v481)+108)) = int32(680)
	*(*int32)(unsafe.Add(mBase, uint32(v481)+104)) = int32(0)
	goto L120
L120:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v473)+12))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)))
	if v494 != int32(141) {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v586 = F_transformStmt(m, v481, v567)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L4
	} else {
		goto L135
	}
L122:
	;
	v567 = v493
	goto L121
L123:
	;
	goto L124
L124:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v493)+68))
	if v497 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v504 = v493
	goto L128
L126:
	;
	v529 = v493
	goto L127
L127:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v529)+8))
	if v546 == int32(0) {
		goto L131
	} else {
		goto L132
	}
L128:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v504)+76))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v521)+68))
	if v522 != 0 {
		v504 = v521
		goto L128
	} else {
		goto L130
	}
L129:
	;
	v529 = v521
	goto L127
L130:
	;
	goto L129
L131:
	;
	v567 = v493
	goto L121
L132:
	;
	goto L133
L133:
	;
	v550 = F_palloc0(m, int32(20))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L4
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v550)+4)) = v493
	*(*int32)(unsafe.Add(mBase, uint32(v550))) = int32(242)
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v529)+8))
	v556 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v550)+16)) = uint8(v556)
	*(*int32)(unsafe.Add(mBase, uint32(v550)+12)) = int32(41)
	*(*int32)(unsafe.Add(mBase, uint32(v550)+8)) = v555
	*(*int32)(unsafe.Add(mBase, uint32(v529)+8)) = int32(0)
	v567 = v550
	goto L121
L135:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v492)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v586)+156)) = v588
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v586)+160)) = v590
	F_free_parsestate(m, v481)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L4
	} else {
		goto L136
	}
L136:
	;
	v602 = v586
	goto L102
L137:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v602)+4))
	if v620 != int32(1) {
		goto L101
	} else {
		goto L138
	}
L138:
	;
	v623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+36)))
	if v623 != 0 {
		goto L101
	} else {
		goto L139
	}
L139:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+37)))
	if v624 != 0 {
		goto L101
	} else {
		goto L140
	}
L140:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+38)))
	if v625 != 0 {
		goto L101
	} else {
		goto L141
	}
L141:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v602)+39)))
	if v626 != 0 {
		goto L101
	} else {
		goto L142
	}
L142:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v602)+48))
	if v627 != 0 {
		goto L101
	} else {
		goto L143
	}
L143:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v602)+52))
	if v628 != 0 {
		goto L101
	} else {
		goto L144
	}
L144:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v602)+60))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v629)+4))
	if v630 != 0 {
		goto L101
	} else {
		goto L145
	}
L145:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v629)+8))
	if v631 != 0 {
		goto L101
	} else {
		goto L146
	}
L146:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v602)+100))
	if v632 != 0 {
		goto L101
	} else {
		goto L147
	}
L147:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v602)+108))
	if v633 != 0 {
		goto L101
	} else {
		goto L148
	}
L148:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v602)+112))
	if v634 != 0 {
		goto L101
	} else {
		goto L149
	}
L149:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v602)+116))
	if v635 != 0 {
		goto L101
	} else {
		goto L150
	}
L150:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v602)+120))
	if v636 != 0 {
		goto L101
	} else {
		goto L151
	}
L151:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v602)+124))
	if v637 != 0 {
		goto L101
	} else {
		goto L152
	}
L152:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v602)+128))
	if v638 != 0 {
		goto L101
	} else {
		goto L153
	}
L153:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v602)+132))
	if v639 != 0 {
		goto L101
	} else {
		goto L154
	}
L154:
	;
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v602)+144))
	if v640 != 0 {
		goto L101
	} else {
		goto L155
	}
L155:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v602)+76))
	if v641 == int32(0) {
		goto L101
	} else {
		goto L156
	}
L156:
	;
	v644 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
	if v644 != int32(1) {
		goto L101
	} else {
		goto L157
	}
L157:
	;
	v650 = F_get_expr_result_type(m, v400, int32(0), v26+int32(84))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = v602
	*(*int32)(unsafe.Add(mBase, uint32(v26)+76)) = v602
	v657 = F_list_make1_impl(m, int32(1), v26+int32(24))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v657
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v657
	v664 = F_list_make1_impl(m, int32(1), v26+int32(20))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L4
	} else {
		goto L160
	}
L160:
	;
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v26)+84))
	v667 = int32(*(*int8)(unsafe.Add(mBase, uint32(v314)+96)))
	v669 = F_check_sql_fn_retval(m, v664, l1, v666, v667, int32(0))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	if v669 != 0 {
		goto L101
	} else {
		goto L162
	}
L162:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v657)+12))
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	if v602 != v672 {
		goto L101
	} else {
		goto L163
	}
L163:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v602)+76))
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v674)+12))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v675)))
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v676)+4))
	v678 = F_exprType(m, v677)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	if v678 != l1 {
		goto L101
	} else {
		goto L165
	}
L165:
	;
	v681 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+101)))
	if v681 == int32(105) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v685 = F_contain_mutable_functions_walker(m, v677, int32(0))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L169
	}
L167:
	;
	v688 = v681
	goto L168
L168:
	;
	if v688&int32(255) == int32(115) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	if v685 != 0 {
		goto L101
	} else {
		goto L170
	}
L170:
	;
	v687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+101)))
	v688 = v687
	goto L168
L171:
	;
	v694 = F_contain_volatile_functions_walker(m, v677, int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L4
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+99)))
	if v696 == int32(1) {
		goto L176
	} else {
		goto L177
	}
L174:
	;
	if v694 != 0 {
		goto L101
	} else {
		goto L175
	}
L175:
	;
	goto L173
L176:
	;
	v700 = F_contain_nonstrict_functions_walker(m, v677, int32(0))
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = int32(0)
	v706 = F_contain_context_dependent_node_walker(m, v49, v26+int32(32))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L4
	} else {
		goto L181
	}
L179:
	;
	if v700 != 0 {
		goto L101
	} else {
		goto L180
	}
L180:
	;
	goto L178
L181:
	;
	if v706 != 0 {
		goto L101
	} else {
		goto L182
	}
L182:
	;
	v708 = int32(*(*int16)(unsafe.Add(mBase, uint32(v314)+104)))
	v711 = F_palloc0(m, v708<<(uint(int32(2))%32))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	v713 = int32(*(*int16)(unsafe.Add(mBase, uint32(v314)+104)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v711
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v713
	v719 = F_substitute_actual_parameters_mutator(m, v677, v26+int32(32))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L184
	}
L184:
	;
	if v49 == int32(0) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v396
	v820 = F_copyObjectImpl(m, v719)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L206
	}
L186:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v723 <= int32(0) {
		goto L185
	} else {
		goto L187
	}
L187:
	;
	v732 = int32(0)
	goto L188
L188:
	;
	v751 = v732 << (uint(int32(2)) % 32)
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v711+v751)))
	switch v753 {
	case 0:
		goto L191
	case 1:
		goto L190
	default:
		goto L192
	}
L189:
	;
	goto L185
L190:
	;
	v792 = v732 + int32(1)
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v792 < v793 {
		v732 = v792
		goto L188
	} else {
		goto L205
	}
L191:
	;
	v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+99)))
	if v789 != 0 {
		goto L101
	} else {
		goto L204
	}
L192:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v754+v751)))
	if v756 != 0 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v756)))
	if base.Ui32(v757-int32(22)) < base.Ui32(int32(3)) {
		goto L101
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v756
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v756
	v773 = F_list_make1_impl(m, int32(1), v26+int32(16))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L4
	} else {
		goto L199
	}
L196:
	;
	v764 = F_expression_tree_walker_impl(m, v756, int32(857), int32(0))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	if v764 != 0 {
		goto L101
	} else {
		goto L198
	}
L198:
	;
	goto L195
L199:
	;
	F_cost_qual_eval(m, v26+int32(32), v773, int32(0))
	mBase = m.M
	v777 = m.ExcPending
	if v777 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v778 = *(*float64)(unsafe.Add(mBase, uint32(v26)+32))
	v779 = *(*float64)(unsafe.Add(mBase, uint32(v26)+40))
	v782 = *(*float64)(unsafe.Add(mBase, _consts[382]))
	if base.F64_gt(base.F64_add(v778, v779), base.F64_mul(v782, float64(10))) != 0 {
		goto L101
	} else {
		goto L201
	}
L201:
	;
	v787 = F_contain_volatile_functions_walker(m, v756, int32(0))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	if v787 != 0 {
		goto L101
	} else {
		goto L203
	}
L203:
	;
	goto L190
L204:
	;
	goto L190
L205:
	;
	goto L189
L206:
	;
	F_MemoryContextDelete(m, v393)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	if l3 == int32(0) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(l9)+4))
	if v841 != 0 {
		goto L220
	} else {
		goto L221
	}
L209:
	;
	v840 = v820
	goto L208
L210:
	;
	goto L211
L211:
	;
	v826 = F_exprCollation(m, v820)
	mBase = m.M
	v827 = m.ExcPending
	if v827 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	if v826 == int32(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v840 = v820
	goto L208
L214:
	;
	goto L215
L215:
	;
	if l3 == v826 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v840 = v820
	goto L208
L217:
	;
	goto L218
L218:
	;
	v832 = F_palloc0(m, int32(16))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v832)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v832)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v832)+4)) = v820
	*(*int32)(unsafe.Add(mBase, uint32(v832))) = int32(31)
	v840 = v832
	goto L208
L220:
	;
	F_record_plan_function_dependency(m, v841, l0)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L4
	} else {
		goto L223
	}
L221:
	;
	goto L222
L222:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v845 = F_lappend_oid(m, v844, l0)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L4
	} else {
		goto L224
	}
L223:
	;
	goto L222
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9)+8)) = v845
	v848 = F_eval_const_expressions_mutator(m, v840, l9)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L4
	} else {
		goto L225
	}
L225:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(l9)+8))
	v851 = F_list_delete_last(m, v850)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L4
	} else {
		goto L226
	}
L226:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9)+8)) = v851
	v887 = v848
	goto L100
L227:
	;
	v887 = int32(0)
	goto L100
L228:
	;
	m.G0 = v26 + int32(112)
	return v913
}
