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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
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
				v80 = v13
				m.G0 = v9 - int32(-64)
				return v80
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
							F_errmsg(m, int32(706372), v9)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return int32(0)
							} else {
								F_errhint(m, int32(618197), int32(0))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493653), int32(2279), int32(130760))
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
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
					v52 = F_SearchSysCache3(m, int32(13), l3, v51, v48)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						v64 = v52
						if v64 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v111 = m.ExcPending
								if v111 != 0 {
									return int32(0)
								} else {
									v112 = F_NameListToString(m, l0)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v112
										F_errmsg(m, int32(706314), v7+int32(-32))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(493653), int32(2317), int32(130760))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
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
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
							v69 = v67 + v68
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+84))
							v72 = F_IsBinaryCoercible(m, l1, v71)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								if v72 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67141764))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											v133 = F_NameListToString(m, l0)
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
												return int32(0)
											} else {
												v135 = F_format_type_be(m, l1)
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v135
													*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v133
													F_errmsg(m, int32(193040), v7+int32(-16))
													mBase = m.M
													v143 = m.ExcPending
													if v143 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(493653), int32(2331), int32(130760))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
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
									F_ReleaseCatCache(m, v64)
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v80 = v70
										m.G0 = v9 - int32(-64)
										return v80
									}
								}
							}
						}
					}
				}
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
				v55 = F_OpclassnameGetOpcid(m, l3, v54)
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					if v55 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = l2
								v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+56))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v93
								F_errmsg(m, int32(706314), v7+int32(-48))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493653), int32(2309), int32(130760))
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
						v60 = F_SearchSysCache1(m, int32(14), v55)
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v64 = v60
							if v64 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v111 = m.ExcPending
									if v111 != 0 {
										return int32(0)
									} else {
										v112 = F_NameListToString(m, l0)
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v112
											F_errmsg(m, int32(706314), v7+int32(-32))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(493653), int32(2317), int32(130760))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
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
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+22)))
								v69 = v67 + v68
								v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
								v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+84))
								v72 = F_IsBinaryCoercible(m, l1, v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									if v72 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(67141764))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												v133 = F_NameListToString(m, l0)
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													v135 = F_format_type_be(m, l1)
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v135
														*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v133
														F_errmsg(m, int32(193040), v7+int32(-16))
														mBase = m.M
														v143 = m.ExcPending
														if v143 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(493653), int32(2331), int32(130760))
															mBase = m.M
															v148 = m.ExcPending
															if v148 != 0 {
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
										F_ReleaseCatCache(m, v64)
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return int32(0)
										} else {
											v80 = v70
											m.G0 = v9 - int32(-64)
											return v80
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
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_SearchSysCacheExists(m, int32(3), l0, int32(115), l1, int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
