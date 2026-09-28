package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ResolveOpClass(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
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
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	if l0 == int32(0) {
		v13 = F_GetDefaultOpClass(m, l1, l3)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v13 != 0 {
				v82 = v13
				m.G0 = v9 - int32(-64)
				return v82
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v24 = F_format_type_be(m, l1)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v24
							F_errmsg(m, int32(_a_F_ResolveOpClass_0), v9)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								F_errhint(m, int32(_a_F_ResolveOpClass_1), int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ResolveOpClass_2), int32(2321), int32(_a_F_ResolveOpClass_3))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
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
		}
	} else {
		F_DeconstructQualifiedName(m, l0, v7+int32(-4), v7+int32(-8))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+60))
			if v46 != 0 {
				v48 = F_LookupExplicitNamespace(m, v46, int32(0))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v52 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+56)))
					v54 = F_SearchSysCache3(m, int32(13), base.I64_extend_i32_u(l3), v52, base.I64_extend_i32_u(v48))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						v67 = v54
						if v67 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v111 = m.ExcPending
							if v111 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int32(0)
								} else {
									v115 = F_NameListToString(m, l0)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v115
										F_errmsg(m, int32(_a_F_ResolveOpClass_4), v7+int32(-32))
										mBase = m.M
										v123 = m.ExcPending
										if v123 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ResolveOpClass_2), int32(2359), int32(_a_F_ResolveOpClass_3))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
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
						} else {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+22)))
							v72 = v70 + v71
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+84))
							v75 = F_IsBinaryCoercible(m, l1, v74)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return int32(0)
							} else {
								if v75 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67141764))
										mBase = m.M
										v135 = m.ExcPending
										if v135 != 0 {
											return int32(0)
										} else {
											v136 = F_NameListToString(m, l0)
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int32(0)
											} else {
												v138 = F_format_type_be(m, l1)
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v138
													*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v136
													F_errmsg(m, int32(_a_F_ResolveOpClass_5), v7+int32(-16))
													mBase = m.M
													v146 = m.ExcPending
													if v146 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_ResolveOpClass_2), int32(2373), int32(_a_F_ResolveOpClass_3))
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
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
									F_ReleaseCatCache(m, v67)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										v82 = v73
										m.G0 = v9 - int32(-64)
										return v82
									}
								}
							}
						}
					}
				}
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
				v57 = F_OpclassnameGetOpcid(m, l3, v56)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					if v57 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l2
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v96
								F_errmsg(m, int32(_a_F_ResolveOpClass_4), v7+int32(-48))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ResolveOpClass_2), int32(2351), int32(_a_F_ResolveOpClass_3))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
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
						v63 = F_SearchSysCache1(m, int32(14), base.I64_extend_i32_u(v57))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v67 = v63
							if v67 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										v115 = F_NameListToString(m, l0)
										mBase = m.M
										v116 = m.ExcPending
										if v116 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v115
											F_errmsg(m, int32(_a_F_ResolveOpClass_4), v7+int32(-32))
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_ResolveOpClass_2), int32(2359), int32(_a_F_ResolveOpClass_3))
												mBase = m.M
												v128 = m.ExcPending
												if v128 != 0 {
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
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+22)))
								v72 = v70 + v71
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v72)+84))
								v75 = F_IsBinaryCoercible(m, l1, v74)
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									if v75 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(67141764))
											mBase = m.M
											v135 = m.ExcPending
											if v135 != 0 {
												return int32(0)
											} else {
												v136 = F_NameListToString(m, l0)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													v138 = F_format_type_be(m, l1)
													mBase = m.M
													v139 = m.ExcPending
													if v139 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v138
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v136
														F_errmsg(m, int32(_a_F_ResolveOpClass_5), v7+int32(-16))
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_ResolveOpClass_2), int32(2373), int32(_a_F_ResolveOpClass_3))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
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
										F_ReleaseCatCache(m, v67)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											v82 = v73
											m.G0 = v9 - int32(-64)
											return v82
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
func F_op_in_opfamily(m *base.Module, l0 int32, l1 int32) int32 {
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v8 = F_SearchSysCacheExists(m, int32(3), base.I64_extend_i32_u(l0), int64(115), base.I64_extend_i32_u(l1), int64(0))
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		return v8
	}
}
