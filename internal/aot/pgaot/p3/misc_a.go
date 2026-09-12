package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AccessTempTableNamespace(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v8 = int32(4384564)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[430]))
	*(*int32)(unsafe.Add(mBase, _consts[430])) = v10 | int32(1)
	if l0 == int32(0) {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[135]))
		if v17 != 0 {
			m.G0 = v6 + int32(128)
			return
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, _consts[128]))
			v22 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v24 = F_object_aclcheck(m, int32(1262), v20, v22, int64(1024))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				if v24 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v141 = m.ExcPending
						if v141 != 0 {
							return
						} else {
							v143 = *(*int32)(unsafe.Add(mBase, _consts[128]))
							v144 = F_get_database_name(m, v143)
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v144
								F_errmsg(m, int32(697771), v6+int32(32))
								mBase = m.M
								v151 = m.ExcPending
								if v151 != 0 {
									return
								} else {
									F_errfinish(m, int32(498918), int32(4416), int32(418237))
									mBase = m.M
									v156 = m.ExcPending
									if v156 != 0 {
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
					v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
					if v28 == int32(1) {
						v33 = *(*int32)(unsafe.Add(mBase, _consts[2]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+316))
						v36 = base.B2i32(v34 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v36)
						v38 = v36
					} else {
						v38 = int32(0)
					}
					if v38 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return
						} else {
							F_errcode(m, int32(100663618))
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return
							} else {
								F_errmsg(m, int32(14475), int32(0))
								mBase = m.M
								v167 = m.ExcPending
								if v167 != 0 {
									return
								} else {
									F_errfinish(m, int32(498918), int32(4431), int32(418237))
									mBase = m.M
									v172 = m.ExcPending
									if v172 != 0 {
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
						v40 = *(*int32)(unsafe.Add(mBase, _consts[431]))
						if int32(0) <= v40 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return
							} else {
								F_errcode(m, int32(100663618))
								mBase = m.M
								v179 = m.ExcPending
								if v179 != 0 {
									return
								} else {
									F_errmsg(m, int32(260131), int32(0))
									mBase = m.M
									v183 = m.ExcPending
									if v183 != 0 {
										return
									} else {
										F_errfinish(m, int32(498918), int32(4437), int32(418237))
										mBase = m.M
										v188 = m.ExcPending
										if v188 != 0 {
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
							v44 = *(*int32)(unsafe.Add(mBase, _consts[109]))
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v44
							v52 = F_pg_snprintf(m, v6+int32(48), int32(64), int32(465348), v6+int32(16))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v57 = int32(0)
								v60 = F_GetSysCacheOid(m, int32(37), v6+int32(48), v57, v57, v57)
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									if v60 == int32(0) {
										v68 = F_NamespaceCreate(m, v6+int32(48), int32(10), int32(1))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v71 = m.ExcPending
											if v71 != 0 {
												return
											} else {
												v83 = v68
												v85 = *(*int32)(unsafe.Add(mBase, _consts[109]))
												*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85
												v91 = F_pg_snprintf(m, v6+int32(48), int32(64), int32(465331), v6)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return
												} else {
													v96 = int32(0)
													v99 = F_GetSysCacheOid(m, int32(37), v6+int32(48), v96, v96, v96)
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return
													} else {
														if v99 == int32(0) {
															v107 = F_NamespaceCreate(m, v6+int32(48), int32(10), int32(1))
															mBase = m.M
															v108 = m.ExcPending
															if v108 != 0 {
																return
															} else {
																F_CommandCounterIncrement(m)
																mBase = m.M
																v110 = m.ExcPending
																if v110 != 0 {
																	return
																} else {
																	v111 = v107
																	*(*int32)(unsafe.Add(mBase, _consts[136])) = v111
																	*(*int32)(unsafe.Add(mBase, _consts[135])) = v83
																	v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
																	*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v83
																	v120 = *(*int32)(unsafe.Add(mBase, _consts[72]))
																	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
																	v123 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v123)
																	*(*int32)(unsafe.Add(mBase, _consts[180])) = v121
																	v128 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v128)
																	m.G0 = v6 + int32(128)
																	return
																}
															}
														} else {
															v111 = v99
															*(*int32)(unsafe.Add(mBase, _consts[136])) = v111
															*(*int32)(unsafe.Add(mBase, _consts[135])) = v83
															v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
															*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v83
															v120 = *(*int32)(unsafe.Add(mBase, _consts[72]))
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
															v123 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v123)
															*(*int32)(unsafe.Add(mBase, _consts[180])) = v121
															v128 = int32(0)
															*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v128)
															m.G0 = v6 + int32(128)
															return
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+124)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v6)+120)) = v60
										*(*int32)(unsafe.Add(mBase, uint32(v6)+116)) = int32(2615)
										F_performDeletion(m, v6+int32(116), int32(1), int32(29))
										mBase = m.M
										v82 = m.ExcPending
										if v82 != 0 {
											return
										} else {
											v83 = v60
											v85 = *(*int32)(unsafe.Add(mBase, _consts[109]))
											*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85
											v91 = F_pg_snprintf(m, v6+int32(48), int32(64), int32(465331), v6)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v96 = int32(0)
												v99 = F_GetSysCacheOid(m, int32(37), v6+int32(48), v96, v96, v96)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													if v99 == int32(0) {
														v107 = F_NamespaceCreate(m, v6+int32(48), int32(10), int32(1))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return
															} else {
																v111 = v107
																*(*int32)(unsafe.Add(mBase, _consts[136])) = v111
																*(*int32)(unsafe.Add(mBase, _consts[135])) = v83
																v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
																*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v83
																v120 = *(*int32)(unsafe.Add(mBase, _consts[72]))
																v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
																v123 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v123)
																*(*int32)(unsafe.Add(mBase, _consts[180])) = v121
																v128 = int32(0)
																*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v128)
																m.G0 = v6 + int32(128)
																return
															}
														}
													} else {
														v111 = v99
														*(*int32)(unsafe.Add(mBase, _consts[136])) = v111
														*(*int32)(unsafe.Add(mBase, _consts[135])) = v83
														v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
														*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v83
														v120 = *(*int32)(unsafe.Add(mBase, _consts[72]))
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
														v123 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v123)
														*(*int32)(unsafe.Add(mBase, _consts[180])) = v121
														v128 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v128)
														m.G0 = v6 + int32(128)
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
		v20 = *(*int32)(unsafe.Add(mBase, _consts[128]))
		v22 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v24 = F_object_aclcheck(m, int32(1262), v20, v22, int64(1024))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			if v24 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v138 = m.ExcPending
				if v138 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return
					} else {
						v143 = *(*int32)(unsafe.Add(mBase, _consts[128]))
						v144 = F_get_database_name(m, v143)
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v144
							F_errmsg(m, int32(697771), v6+int32(32))
							mBase = m.M
							v151 = m.ExcPending
							if v151 != 0 {
								return
							} else {
								F_errfinish(m, int32(498918), int32(4416), int32(418237))
								mBase = m.M
								v156 = m.ExcPending
								if v156 != 0 {
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
				v28 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1])))
				if v28 == int32(1) {
					v33 = *(*int32)(unsafe.Add(mBase, _consts[2]))
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+316))
					v36 = base.B2i32(v34 != int32(2))
					*(*uint8)(unsafe.Add(mBase, _consts[1])) = uint8(v36)
					v38 = v36
				} else {
					v38 = int32(0)
				}
				if v38 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return
					} else {
						F_errcode(m, int32(100663618))
						mBase = m.M
						v163 = m.ExcPending
						if v163 != 0 {
							return
						} else {
							F_errmsg(m, int32(14475), int32(0))
							mBase = m.M
							v167 = m.ExcPending
							if v167 != 0 {
								return
							} else {
								F_errfinish(m, int32(498918), int32(4431), int32(418237))
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
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
					v40 = *(*int32)(unsafe.Add(mBase, _consts[431]))
					if int32(0) <= v40 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return
						} else {
							F_errcode(m, int32(100663618))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return
							} else {
								F_errmsg(m, int32(260131), int32(0))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return
								} else {
									F_errfinish(m, int32(498918), int32(4437), int32(418237))
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
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
						v44 = *(*int32)(unsafe.Add(mBase, _consts[109]))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v44
						v52 = F_pg_snprintf(m, v6+int32(48), int32(64), int32(465348), v6+int32(16))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v57 = int32(0)
							v60 = F_GetSysCacheOid(m, int32(37), v6+int32(48), v57, v57, v57)
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								if v60 == int32(0) {
									v68 = F_NamespaceCreate(m, v6+int32(48), int32(10), int32(1))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v71 = m.ExcPending
										if v71 != 0 {
											return
										} else {
											v83 = v68
											v85 = *(*int32)(unsafe.Add(mBase, _consts[109]))
											*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85
											v91 = F_pg_snprintf(m, v6+int32(48), int32(64), int32(465331), v6)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return
											} else {
												v96 = int32(0)
												v99 = F_GetSysCacheOid(m, int32(37), v6+int32(48), v96, v96, v96)
												mBase = m.M
												v100 = m.ExcPending
												if v100 != 0 {
													return
												} else {
													if v99 == int32(0) {
														v107 = F_NamespaceCreate(m, v6+int32(48), int32(10), int32(1))
														mBase = m.M
														v108 = m.ExcPending
														if v108 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return
															} else {
																v111 = v107
																*(*int32)(unsafe.Add(mBase, _consts[136])) = v111
																*(*int32)(unsafe.Add(mBase, _consts[135])) = v83
																v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
																*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v83
																v120 = *(*int32)(unsafe.Add(mBase, _consts[72]))
																v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
																v123 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v123)
																*(*int32)(unsafe.Add(mBase, _consts[180])) = v121
																v128 = int32(0)
																*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v128)
																m.G0 = v6 + int32(128)
																return
															}
														}
													} else {
														v111 = v99
														*(*int32)(unsafe.Add(mBase, _consts[136])) = v111
														*(*int32)(unsafe.Add(mBase, _consts[135])) = v83
														v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
														*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v83
														v120 = *(*int32)(unsafe.Add(mBase, _consts[72]))
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
														v123 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v123)
														*(*int32)(unsafe.Add(mBase, _consts[180])) = v121
														v128 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v128)
														m.G0 = v6 + int32(128)
														return
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+124)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+120)) = v60
									*(*int32)(unsafe.Add(mBase, uint32(v6)+116)) = int32(2615)
									F_performDeletion(m, v6+int32(116), int32(1), int32(29))
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return
									} else {
										v83 = v60
										v85 = *(*int32)(unsafe.Add(mBase, _consts[109]))
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = v85
										v91 = F_pg_snprintf(m, v6+int32(48), int32(64), int32(465331), v6)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return
										} else {
											v96 = int32(0)
											v99 = F_GetSysCacheOid(m, int32(37), v6+int32(48), v96, v96, v96)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return
											} else {
												if v99 == int32(0) {
													v107 = F_NamespaceCreate(m, v6+int32(48), int32(10), int32(1))
													mBase = m.M
													v108 = m.ExcPending
													if v108 != 0 {
														return
													} else {
														F_CommandCounterIncrement(m)
														mBase = m.M
														v110 = m.ExcPending
														if v110 != 0 {
															return
														} else {
															v111 = v107
															*(*int32)(unsafe.Add(mBase, _consts[136])) = v111
															*(*int32)(unsafe.Add(mBase, _consts[135])) = v83
															v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
															*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v83
															v120 = *(*int32)(unsafe.Add(mBase, _consts[72]))
															v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
															v123 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v123)
															*(*int32)(unsafe.Add(mBase, _consts[180])) = v121
															v128 = int32(0)
															*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v128)
															m.G0 = v6 + int32(128)
															return
														}
													}
												} else {
													v111 = v99
													*(*int32)(unsafe.Add(mBase, _consts[136])) = v111
													*(*int32)(unsafe.Add(mBase, _consts[135])) = v83
													v117 = *(*int32)(unsafe.Add(mBase, _consts[137]))
													*(*int32)(unsafe.Add(mBase, uint32(v117)+68)) = v83
													v120 = *(*int32)(unsafe.Add(mBase, _consts[72]))
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
													v123 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _consts[432])) = uint8(v123)
													*(*int32)(unsafe.Add(mBase, _consts[180])) = v121
													v128 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _consts[433])) = uint8(v128)
													m.G0 = v6 + int32(128)
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
}
func F_AdvanceOldestClogXid(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v3 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v7 = F_LWLockAcquire(m, v3+int32(5632), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[87]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v11)) == int32(0) {
			v23 = base.B2i32(base.Ui32(v11) < base.Ui32(l0))
		} else {
			v23 = int32(base.Ui32(v11-l0) >> (uint(int32(31)) % 32))
		}
		if v23 != 0 {
			v25 = *(*int32)(unsafe.Add(mBase, _consts[87]))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = l0
		} else {
		}
		v28 = *(*int32)(unsafe.Add(mBase, _consts[44]))
		F_LWLockRelease(m, v28+int32(5632))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			return
		}
	}
}
func F_AlignedAllocFree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	v3 = l0 - int32(8)
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	F_pfree(m, v3-base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(34))%64)))&int32(1073741822))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		return
	}
}
func F_AlignedAllocGetChunkContext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v3 = l0 - int32(8)
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v11 = F_GetMemoryChunkContext(m, v3-base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(34))%64)))&int32(1073741822))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
func F_AlterPublicationOwner_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	v13 = v11 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	if v14 == l2 {
		m.G0 = v9 + int32(32)
		return
	} else {
		v16 = F_superuser(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			if v16 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
				F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					F_changeDependencyOnOwner(m, int32(6104), v68, l2)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
						if v72 == int32(0) {
							m.G0 = v9 + int32(32)
							return
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v77 = int32(0)
							F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v21 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v22 = F_object_ownercheck(m, int32(6104), v19, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					if v22 == int32(0) {
						F_aclcheck_error(m, int32(2), int32(30), v13+int32(4))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							F_check_can_set_role(m, v33, l2)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, _consts[128]))
								v40 = F_object_aclcheck(m, int32(1262), v38, l2, int64(512))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									if v40 != 0 {
										v44 = *(*int32)(unsafe.Add(mBase, _consts[128]))
										v45 = F_get_database_name(m, v44)
										mBase = m.M
										v46 = m.ExcPending
										if v46 != 0 {
											return
										} else {
											F_aclcheck_error(m, v40, int32(9), v45)
											mBase = m.M
											v48 = m.ExcPending
											if v48 != 0 {
												return
											} else {
												v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)))
												if v49 == int32(1) {
													v52 = F_superuser_arg(m, l2)
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return
													} else {
														if v52 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return
															} else {
																F_errcode(m, int32(16797828))
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13 + int32(4)
																	F_errmsg(m, int32(693405), v9+int32(16))
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(593919), int32(0))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(493767), int32(2032), int32(310845))
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
															}
														} else {
															v56 = F_superuser_arg(m, l2)
															mBase = m.M
															v57 = m.ExcPending
															if v57 != 0 {
																return
															} else {
																if v56 != 0 {
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
																	F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
																	mBase = m.M
																	v66 = m.ExcPending
																	if v66 != 0 {
																		return
																	} else {
																		v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return
																		} else {
																			v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																			if v72 == int32(0) {
																				m.G0 = v9 + int32(32)
																				return
																			} else {
																				v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				v77 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																				mBase = m.M
																				v81 = m.ExcPending
																				if v81 != 0 {
																					return
																				} else {
																					m.G0 = v9 + int32(32)
																					return
																				}
																			}
																		}
																	}
																} else {
																	v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v59 = F_is_schema_publication(m, v58)
																	mBase = m.M
																	v60 = m.ExcPending
																	if v60 != 0 {
																		return
																	} else {
																		if v59 != 0 {
																			F_errstart_cold(m, int32(21), int32(0))
																			mBase = m.M
																			v113 = m.ExcPending
																			if v113 != 0 {
																				return
																			} else {
																				F_errcode(m, int32(16797828))
																				mBase = m.M
																				v116 = m.ExcPending
																				if v116 != 0 {
																					return
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																					F_errmsg(m, int32(693405), v9)
																					mBase = m.M
																					v122 = m.ExcPending
																					if v122 != 0 {
																						return
																					} else {
																						F_errhint(m, int32(593982), int32(0))
																						mBase = m.M
																						v126 = m.ExcPending
																						if v126 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(493767), int32(2039), int32(310845))
																							mBase = m.M
																							v131 = m.ExcPending
																							if v131 != 0 {
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
																			*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
																			F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
																			mBase = m.M
																			v66 = m.ExcPending
																			if v66 != 0 {
																				return
																			} else {
																				v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																				mBase = m.M
																				v70 = m.ExcPending
																				if v70 != 0 {
																					return
																				} else {
																					v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																					if v72 == int32(0) {
																						m.G0 = v9 + int32(32)
																						return
																					} else {
																						v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																						v77 = int32(0)
																						F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																						mBase = m.M
																						v81 = m.ExcPending
																						if v81 != 0 {
																							return
																						} else {
																							m.G0 = v9 + int32(32)
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
												} else {
													v56 = F_superuser_arg(m, l2)
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
														return
													} else {
														if v56 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
															F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
															mBase = m.M
															v66 = m.ExcPending
															if v66 != 0 {
																return
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																mBase = m.M
																v70 = m.ExcPending
																if v70 != 0 {
																	return
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																	if v72 == int32(0) {
																		m.G0 = v9 + int32(32)
																		return
																	} else {
																		v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v77 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(32)
																			return
																		}
																	}
																}
															}
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															v59 = F_is_schema_publication(m, v58)
															mBase = m.M
															v60 = m.ExcPending
															if v60 != 0 {
																return
															} else {
																if v59 != 0 {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v113 = m.ExcPending
																	if v113 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(16797828))
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																			F_errmsg(m, int32(693405), v9)
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return
																			} else {
																				F_errhint(m, int32(593982), int32(0))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(493767), int32(2039), int32(310845))
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
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
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
																	F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
																	mBase = m.M
																	v66 = m.ExcPending
																	if v66 != 0 {
																		return
																	} else {
																		v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return
																		} else {
																			v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																			if v72 == int32(0) {
																				m.G0 = v9 + int32(32)
																				return
																			} else {
																				v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				v77 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																				mBase = m.M
																				v81 = m.ExcPending
																				if v81 != 0 {
																					return
																				} else {
																					m.G0 = v9 + int32(32)
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
									} else {
										v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)))
										if v49 == int32(1) {
											v52 = F_superuser_arg(m, l2)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												if v52 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return
													} else {
														F_errcode(m, int32(16797828))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13 + int32(4)
															F_errmsg(m, int32(693405), v9+int32(16))
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																F_errhint(m, int32(593919), int32(0))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(493767), int32(2032), int32(310845))
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
													}
												} else {
													v56 = F_superuser_arg(m, l2)
													mBase = m.M
													v57 = m.ExcPending
													if v57 != 0 {
														return
													} else {
														if v56 != 0 {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
															F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
															mBase = m.M
															v66 = m.ExcPending
															if v66 != 0 {
																return
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																mBase = m.M
																v70 = m.ExcPending
																if v70 != 0 {
																	return
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																	if v72 == int32(0) {
																		m.G0 = v9 + int32(32)
																		return
																	} else {
																		v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v77 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(32)
																			return
																		}
																	}
																}
															}
														} else {
															v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															v59 = F_is_schema_publication(m, v58)
															mBase = m.M
															v60 = m.ExcPending
															if v60 != 0 {
																return
															} else {
																if v59 != 0 {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v113 = m.ExcPending
																	if v113 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(16797828))
																		mBase = m.M
																		v116 = m.ExcPending
																		if v116 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																			F_errmsg(m, int32(693405), v9)
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return
																			} else {
																				F_errhint(m, int32(593982), int32(0))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(493767), int32(2039), int32(310845))
																					mBase = m.M
																					v131 = m.ExcPending
																					if v131 != 0 {
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
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
																	F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
																	mBase = m.M
																	v66 = m.ExcPending
																	if v66 != 0 {
																		return
																	} else {
																		v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return
																		} else {
																			v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																			if v72 == int32(0) {
																				m.G0 = v9 + int32(32)
																				return
																			} else {
																				v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				v77 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																				mBase = m.M
																				v81 = m.ExcPending
																				if v81 != 0 {
																					return
																				} else {
																					m.G0 = v9 + int32(32)
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
										} else {
											v56 = F_superuser_arg(m, l2)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return
											} else {
												if v56 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
													F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														F_changeDependencyOnOwner(m, int32(6104), v68, l2)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
															if v72 == int32(0) {
																m.G0 = v9 + int32(32)
																return
															} else {
																v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																v77 = int32(0)
																F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																mBase = m.M
																v81 = m.ExcPending
																if v81 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(32)
																	return
																}
															}
														}
													}
												} else {
													v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													v59 = F_is_schema_publication(m, v58)
													mBase = m.M
													v60 = m.ExcPending
													if v60 != 0 {
														return
													} else {
														if v59 != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v113 = m.ExcPending
															if v113 != 0 {
																return
															} else {
																F_errcode(m, int32(16797828))
																mBase = m.M
																v116 = m.ExcPending
																if v116 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																	F_errmsg(m, int32(693405), v9)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(593982), int32(0))
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(493767), int32(2039), int32(310845))
																			mBase = m.M
																			v131 = m.ExcPending
																			if v131 != 0 {
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
															*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
															F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
															mBase = m.M
															v66 = m.ExcPending
															if v66 != 0 {
																return
															} else {
																v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																mBase = m.M
																v70 = m.ExcPending
																if v70 != 0 {
																	return
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																	if v72 == int32(0) {
																		m.G0 = v9 + int32(32)
																		return
																	} else {
																		v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v77 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																		mBase = m.M
																		v81 = m.ExcPending
																		if v81 != 0 {
																			return
																		} else {
																			m.G0 = v9 + int32(32)
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
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						F_check_can_set_role(m, v33, l2)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, _consts[128]))
							v40 = F_object_aclcheck(m, int32(1262), v38, l2, int64(512))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								if v40 != 0 {
									v44 = *(*int32)(unsafe.Add(mBase, _consts[128]))
									v45 = F_get_database_name(m, v44)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										F_aclcheck_error(m, v40, int32(9), v45)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)))
											if v49 == int32(1) {
												v52 = F_superuser_arg(m, l2)
												mBase = m.M
												v53 = m.ExcPending
												if v53 != 0 {
													return
												} else {
													if v52 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return
														} else {
															F_errcode(m, int32(16797828))
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13 + int32(4)
																F_errmsg(m, int32(693405), v9+int32(16))
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	F_errhint(m, int32(593919), int32(0))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(493767), int32(2032), int32(310845))
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
														}
													} else {
														v56 = F_superuser_arg(m, l2)
														mBase = m.M
														v57 = m.ExcPending
														if v57 != 0 {
															return
														} else {
															if v56 != 0 {
																*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
																F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
																mBase = m.M
																v66 = m.ExcPending
																if v66 != 0 {
																	return
																} else {
																	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return
																	} else {
																		v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																		if v72 == int32(0) {
																			m.G0 = v9 + int32(32)
																			return
																		} else {
																			v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v77 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return
																			} else {
																				m.G0 = v9 + int32(32)
																				return
																			}
																		}
																	}
																}
															} else {
																v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																v59 = F_is_schema_publication(m, v58)
																mBase = m.M
																v60 = m.ExcPending
																if v60 != 0 {
																	return
																} else {
																	if v59 != 0 {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v113 = m.ExcPending
																		if v113 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(16797828))
																			mBase = m.M
																			v116 = m.ExcPending
																			if v116 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																				F_errmsg(m, int32(693405), v9)
																				mBase = m.M
																				v122 = m.ExcPending
																				if v122 != 0 {
																					return
																				} else {
																					F_errhint(m, int32(593982), int32(0))
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(493767), int32(2039), int32(310845))
																						mBase = m.M
																						v131 = m.ExcPending
																						if v131 != 0 {
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
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
																		F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
																		mBase = m.M
																		v66 = m.ExcPending
																		if v66 != 0 {
																			return
																		} else {
																			v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																			mBase = m.M
																			v70 = m.ExcPending
																			if v70 != 0 {
																				return
																			} else {
																				v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																				if v72 == int32(0) {
																					m.G0 = v9 + int32(32)
																					return
																				} else {
																					v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																					v77 = int32(0)
																					F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																					mBase = m.M
																					v81 = m.ExcPending
																					if v81 != 0 {
																						return
																					} else {
																						m.G0 = v9 + int32(32)
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
											} else {
												v56 = F_superuser_arg(m, l2)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return
												} else {
													if v56 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
														F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_changeDependencyOnOwner(m, int32(6104), v68, l2)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																if v72 == int32(0) {
																	m.G0 = v9 + int32(32)
																	return
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v77 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																}
															}
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v59 = F_is_schema_publication(m, v58)
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
															return
														} else {
															if v59 != 0 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v113 = m.ExcPending
																if v113 != 0 {
																	return
																} else {
																	F_errcode(m, int32(16797828))
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																		F_errmsg(m, int32(693405), v9)
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return
																		} else {
																			F_errhint(m, int32(593982), int32(0))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(493767), int32(2039), int32(310845))
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
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
																*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
																F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
																mBase = m.M
																v66 = m.ExcPending
																if v66 != 0 {
																	return
																} else {
																	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return
																	} else {
																		v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																		if v72 == int32(0) {
																			m.G0 = v9 + int32(32)
																			return
																		} else {
																			v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v77 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return
																			} else {
																				m.G0 = v9 + int32(32)
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
								} else {
									v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)))
									if v49 == int32(1) {
										v52 = F_superuser_arg(m, l2)
										mBase = m.M
										v53 = m.ExcPending
										if v53 != 0 {
											return
										} else {
											if v52 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													F_errcode(m, int32(16797828))
													mBase = m.M
													v92 = m.ExcPending
													if v92 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v13 + int32(4)
														F_errmsg(m, int32(693405), v9+int32(16))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															F_errhint(m, int32(593919), int32(0))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																F_errfinish(m, int32(493767), int32(2032), int32(310845))
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
												}
											} else {
												v56 = F_superuser_arg(m, l2)
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
													return
												} else {
													if v56 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
														F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_changeDependencyOnOwner(m, int32(6104), v68, l2)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																if v72 == int32(0) {
																	m.G0 = v9 + int32(32)
																	return
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v77 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
																		return
																	}
																}
															}
														}
													} else {
														v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v59 = F_is_schema_publication(m, v58)
														mBase = m.M
														v60 = m.ExcPending
														if v60 != 0 {
															return
														} else {
															if v59 != 0 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v113 = m.ExcPending
																if v113 != 0 {
																	return
																} else {
																	F_errcode(m, int32(16797828))
																	mBase = m.M
																	v116 = m.ExcPending
																	if v116 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																		F_errmsg(m, int32(693405), v9)
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return
																		} else {
																			F_errhint(m, int32(593982), int32(0))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(493767), int32(2039), int32(310845))
																				mBase = m.M
																				v131 = m.ExcPending
																				if v131 != 0 {
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
																*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
																F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
																mBase = m.M
																v66 = m.ExcPending
																if v66 != 0 {
																	return
																} else {
																	v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_changeDependencyOnOwner(m, int32(6104), v68, l2)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return
																	} else {
																		v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																		if v72 == int32(0) {
																			m.G0 = v9 + int32(32)
																			return
																		} else {
																			v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v77 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																			mBase = m.M
																			v81 = m.ExcPending
																			if v81 != 0 {
																				return
																			} else {
																				m.G0 = v9 + int32(32)
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
									} else {
										v56 = F_superuser_arg(m, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return
										} else {
											if v56 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
												F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return
												} else {
													v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													F_changeDependencyOnOwner(m, int32(6104), v68, l2)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
														if v72 == int32(0) {
															m.G0 = v9 + int32(32)
															return
														} else {
															v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															v77 = int32(0)
															F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
															mBase = m.M
															v81 = m.ExcPending
															if v81 != 0 {
																return
															} else {
																m.G0 = v9 + int32(32)
																return
															}
														}
													}
												}
											} else {
												v58 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												v59 = F_is_schema_publication(m, v58)
												mBase = m.M
												v60 = m.ExcPending
												if v60 != 0 {
													return
												} else {
													if v59 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v113 = m.ExcPending
														if v113 != 0 {
															return
														} else {
															F_errcode(m, int32(16797828))
															mBase = m.M
															v116 = m.ExcPending
															if v116 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																F_errmsg(m, int32(693405), v9)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return
																} else {
																	F_errhint(m, int32(593982), int32(0))
																	mBase = m.M
																	v126 = m.ExcPending
																	if v126 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(493767), int32(2039), int32(310845))
																		mBase = m.M
																		v131 = m.ExcPending
																		if v131 != 0 {
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
														*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
														F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
														mBase = m.M
														v66 = m.ExcPending
														if v66 != 0 {
															return
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_changeDependencyOnOwner(m, int32(6104), v68, l2)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, _consts[442]))
																if v72 == int32(0) {
																	m.G0 = v9 + int32(32)
																	return
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v77 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(6104), v76, v77, v77, v77)
																	mBase = m.M
																	v81 = m.ExcPending
																	if v81 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(32)
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
					}
				}
			}
		}
	}
}
func F_AlterSubscription_refresh(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v180 int32
	_ = v180
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v224 int32
	_ = v224
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v401 int32
	_ = v401
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v487 int32
	_ = v487
	var v497 int32
	_ = v497
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v538 int32
	_ = v538
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v626 int32
	_ = v626
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v663 int32
	_ = v663
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v751 int32
	_ = v751
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int64
	_ = v774
	var v775 int32
	_ = v775
	var v798 int32
	_ = v798
	var v819 int32
	_ = v819
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v836 int32
	_ = v836
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v863 int32
	_ = v863
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v905 int32
	_ = v905
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v915 int32
	_ = v915
	var v926 int32
	_ = v926
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1063 int32
	_ = v1063
	var v1083 int32
	_ = v1083
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1123 int32
	_ = v1123
	var v1124 int32
	_ = v1124
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1184 int32
	_ = v1184
	var v1205 int32
	_ = v1205
	var v1211 int32
	_ = v1211
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1245 int32
	_ = v1245
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1266 int64
	_ = v1266
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1301 int32
	_ = v1301
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1353 int32
	_ = v1353
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1385 int32
	_ = v1385
	var v1403 int32
	_ = v1403
	var v1428 int32
	_ = v1428
	var v1457 int32
	_ = v1457
	var v1464 int32
	_ = v1464
	var v1465 int64
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1471 int32
	_ = v1471
	var v1472 int32
	_ = v1472
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1480 int32
	_ = v1480
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1488 int32
	_ = v1488
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1497 int32
	_ = v1497
	v4 = int32(0)
	v36 = m.G0
	v38 = v36 - int32(112)
	m.G0 = v38
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v43 = int32(105)
	goto L3
L2:
	;
	v43 = int32(114)
	goto L3
L3:
	;
	v48 = v4
	v49 = v4
	v50 = v4
	v51 = v4
	v52 = v4
	v53 = v4
	v54 = v4
	v55 = v4
	v56 = v4
	v57 = v4
	v58 = v4
	v59 = v4
	v60 = v4
	v61 = v4
	v62 = v4
	v63 = v4
	v66 = int32(-1)
	v72 = v38
	v73 = v4
	goto L5
L4:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L5:
	;
	goto L8
L6:
	;
	m.G0 = v38 + int32(112)
	return
L7:
	;
	goto L6
L8:
	;
	if v66 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v1464 = int32(m.ExcTag)
	v1465 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1464 == int32(0) {
		goto L123
	} else {
		goto L124
	}
L11:
	;
	v81 = int32(16)
	v82 = v72 - v81
	m.G0 = v82
	v85 = v82 - int32(160)
	m.G0 = v85
	v88 = v85 - v81
	m.G0 = v88
	v91 = v88 - v81
	m.G0 = v91
	v94 = v91 - v81
	m.G0 = v94
	v97 = v94 + int32(-64)
	m.G0 = v97
	v100 = v94 + int32(-128)
	m.G0 = v100
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v82
	F_load_file(m, int32(214713), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		v1457 = v100
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v256 = v48
	v257 = v49
	v258 = v50
	v259 = v51
	v260 = v53
	v261 = v54
	v262 = v55
	v263 = v56
	v264 = v58
	v265 = v59
	v266 = v60
	v271 = v72
	v272 = v73
	goto L13
L13:
	;
	v273 = int32(0)
	if v272 != 0 {
		v1334 = v52
		v1339 = v57
		v1343 = v61
		v1344 = v62
		v1345 = v63
		v1353 = v273
		goto L30
	} else {
		goto L31
	}
L14:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v123 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v129 = v126 ^ int32(1)
	goto L17
L16:
	;
	v129 = int32(0)
	goto L17
L17:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v133 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v50
	v144 = l0 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v82
	v153 = int32(1)
	v157 = m.T0[v134].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v131, v153, v153, v129&v153, v130, v82)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		v1457 = v100
		goto L10
	} else {
		goto L18
	}
L18:
	;
	if v157 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v82
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		v1457 = v100
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v248 = *(*int32)(unsafe.Add(mBase, _consts[49]))
	v250 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	goto L26
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v82
	F_errcode(m, int32(100663808))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		v1457 = v100
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v38)+36)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = v200
	F_errmsg(m, int32(201419), v38+int32(32))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		v1457 = v100
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v97
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v91
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v82
	F_errfinish(m, int32(493717), int32(853), int32(322277))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		v1457 = v100
		goto L10
	} else {
		goto L25
	}
L25:
	;
	goto L4
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v38 + int32(44)
	goto L29
L27:
	;
	v256 = v91
	v257 = v100
	v258 = v157
	v259 = v88
	v260 = v144
	v261 = v82
	v262 = v85
	v263 = v97
	v264 = v248
	v265 = v250
	v266 = v94
	v271 = v100
	v272 = int32(0)
	goto L13
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v264
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v265
	v1366 = *(*int32)(unsafe.Add(mBase, _consts[493]))
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(v1366)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v1343
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v1334
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v1344
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v1339
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v1345
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	m.T0[v1367].(func(*base.Module, int32))(m, v258)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L116
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v262
	if l2 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_check_publications(m, v258, l2)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v311 = F_fetch_table_list(m, v258, v294)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v331 = F_GetSubscriptionRelations(m, v313, int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_pg_qsort(m, v497, v487, int32(4), int32(471))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L50
	}
L38:
	;
	if v331 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v351 = int32(0)
	v353 = F_palloc(m, v351)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v374 = F_palloc(m, v355<<(uint(int32(2))%32))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L43
	}
L42:
	;
	v475 = v57
	v481 = v353
	v487 = v351
	v497 = v353
	goto L37
L43:
	;
	v376 = int32(0)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v376 < v377 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v401 = v376
	goto L47
L45:
	;
	goto L46
L46:
	;
	v475 = v374
	v481 = v63
	v487 = v355
	v497 = v374
	goto L37
L47:
	;
	v416 = v401 << (uint(int32(2)) % 32)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v418+v416)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	*(*int32)(unsafe.Add(mBase, uint32(v374+v416))) = v421
	v424 = v401 + int32(1)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v424 < v425 {
		v401 = v424
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	goto L48
L50:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_check_publications_origin(m, v258, v520, l1, v519, v497, v487, v518)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v557 = F_palloc(m, v487<<(uint(int32(3))%32))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L52
	}
L52:
	;
	if v311 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_pg_qsort(m, v863, v849, int32(4), int32(471))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L74
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v577 = int32(0)
	v579 = F_palloc(m, v577)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v600 = F_palloc(m, v581<<(uint(int32(2))%32))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L58
	}
L57:
	;
	v836 = v52
	v846 = v579
	v849 = v577
	v863 = v579
	goto L53
L58:
	;
	v602 = int32(0)
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	if v603 <= v602 {
		v836 = v600
		v846 = v62
		v849 = v603
		v863 = v600
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v626 = v602
	goto L60
L60:
	;
	v642 = v626 << (uint(int32(2)) % 32)
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v311)+12))
	v645 = *(*int32)(unsafe.Add(mBase, uint32(v642+v643)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v663 = int32(0)
	v666 = F_RangeVarGetRelidExtended(m, v645, int32(1), v663, v663, v663)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L62
	}
L61:
	;
	v836 = v600
	v846 = v62
	v849 = v826
	v863 = v600
	goto L53
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v259))) = v666
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v685 = F_get_rel_relkind(m, v666)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L63
	}
L63:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v645)+12))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v645)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_CheckSubscriptionRelkind(m, v685, v688, v687)
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L64
	}
L64:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	*(*int32)(unsafe.Add(mBase, uint32(v600+v642))) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v728 = F_bsearch(m, v259, v497, v487, int32(4), int32(471))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L66
	}
L65:
	;
	v825 = v626 + int32(1)
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	if v825 < v826 {
		v626 = v825
		goto L60
	} else {
		goto L73
	}
L66:
	;
	if v728 != 0 {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_AddSubscriptionRelState(m, v731, v730, v43, int64(0), int32(1))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v770 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L69
	}
L69:
	;
	if v770 == int32(0) {
		goto L65
	} else {
		goto L70
	}
L70:
	;
	v774 = *(*int64)(unsafe.Add(mBase, uint32(v645)+8))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = v775
	*(*int64)(unsafe.Add(mBase, uint32(v38)+16)) = v774
	F_errmsg_internal(m, int32(687099), v38+int32(16))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v600
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_errfinish(m, int32(493717), int32(924), int32(322277))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L72
	}
L72:
	;
	goto L65
L73:
	;
	goto L61
L74:
	;
	if v487 <= int32(0) {
		v1334 = v836
		v1339 = v475
		v1343 = v61
		v1344 = v846
		v1345 = v481
		v1353 = v273
		goto L30
	} else {
		goto L75
	}
L75:
	;
	v886 = int32(0)
	v905 = v61
	v909 = v886
	v911 = v886
	v915 = v273
	goto L76
L76:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v497+v909<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v256))) = v926
	if v311 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v1220 = int32(0)
	if v1213 <= v1220 {
		v1334 = v836
		v1339 = v475
		v1343 = v1211
		v1344 = v846
		v1345 = v481
		v1353 = v1216
		goto L30
	} else {
		goto L107
	}
L78:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	v930 = v929
	goto L80
L79:
	;
	v930 = int32(0)
	goto L80
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v905
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v949 = F_bsearch(m, v256, v863, v930, int32(4), int32(471))
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L81
	}
L81:
	;
	if v949 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	if v915 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	v1211 = v905
	v1213 = v911
	v1216 = v915
	goto L84
L84:
	;
	v1218 = v909 + int32(1)
	if v1218 != v487 {
		v905 = v1211
		v909 = v1218
		v911 = v1213
		v915 = v1216
		goto L76
	} else {
		goto L106
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v905
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v973 = F_table_open(m, int32(6102), int32(8))
	mBase = m.M
	v974 = m.ExcPending
	if v974 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L88
	}
L86:
	;
	v975 = v905
	v976 = v915
	goto L87
L87:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v995 = F_GetSubscriptionRelState(m, v978, v977, v266)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L89
	}
L88:
	;
	v975 = v973
	v976 = v973
	goto L87
L89:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v1000 = v557 + v911<<(uint(int32(3))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v1000)+4)) = uint8(v995)
	*(*int32)(unsafe.Add(mBase, uint32(v1000))) = v997
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_RemoveSubscriptionRel(m, v1003, v997)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L90
	}
L90:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_logicalrep_worker_stop(m, v1023, v1022)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L91
	}
L91:
	;
	if v995 != int32(114) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_ReplicationOriginNameForLogicalRep(m, v1045, v1044, v263)
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v1104 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_replorigin_drop_by_name(m, v263, int32(1), int32(0))
	mBase = m.M
	v1083 = m.ExcPending
	if v1083 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	if v1104 != 0 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v1106 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v1123 = F_get_rel_namespace(m, v1106)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v1211 = v975
	v1213 = v911 + int32(1)
	v1216 = v976
	goto L84
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v1141 = F_get_namespace_name(m, v1123)
	mBase = m.M
	v1142 = m.ExcPending
	if v1142 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L102
	}
L102:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	v1160 = F_get_rel_name(m, v1143)
	mBase = m.M
	v1161 = m.ExcPending
	if v1161 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L103
	}
L103:
	;
	v1162 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v1162
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v1141
	F_errmsg_internal(m, int32(687234), v38)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v975
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_errfinish(m, int32(493717), int32(1000), int32(322277))
	mBase = m.M
	v1205 = m.ExcPending
	if v1205 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L105
	}
L105:
	;
	goto L100
L106:
	;
	goto L77
L107:
	;
	v1245 = v1220
	goto L108
L108:
	;
	v1260 = v557 + v1245<<(uint(int32(3))%32)
	v1261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1260)+4)))
	if v1261&int32(254) != int32(114) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v1334 = v836
	v1339 = v475
	v1343 = v1211
	v1344 = v846
	v1345 = v481
	v1353 = v1216
	goto L30
L110:
	;
	v1266 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v257))) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v257)+56)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v257)+48)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v257)+40)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v257)+32)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v257)+24)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v257)+16)) = v1266
	*(*int64)(unsafe.Add(mBase, uint32(v257)+8)) = v1266
	v1282 = *(*int32)(unsafe.Add(mBase, uint32(v1260)))
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_ReplicationSlotNameForTablesync(m, v1283, v1282, v257)
	mBase = m.M
	v1301 = m.ExcPending
	if v1301 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v1324 = v1245 + int32(1)
	if v1324 != v1213 {
		v1245 = v1324
		goto L108
	} else {
		goto L115
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v836
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v1211
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v846
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v475
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v481
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_ReplicationSlotDropAtPubNode(m, v258, v257, int32(1))
	mBase = m.M
	v1320 = m.ExcPending
	if v1320 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	goto L109
L116:
	;
	if v272 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v1334
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v1343
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v1344
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v1339
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v1345
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_pg_re_throw(m)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	*(*int32)(unsafe.Add(mBase, _consts[49])) = v264
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v265
	if v1353 == int32(0) {
		goto L7
	} else {
		goto L121
	}
L120:
	;
	goto L4
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+52)) = v1334
	*(*int32)(unsafe.Add(mBase, uint32(v38)+48)) = v1343
	*(*int32)(unsafe.Add(mBase, uint32(v38)+56)) = v1344
	*(*int32)(unsafe.Add(mBase, uint32(v38)+60)) = v1339
	*(*int32)(unsafe.Add(mBase, uint32(v38)+64)) = v1345
	*(*int32)(unsafe.Add(mBase, uint32(v38)+68)) = v264
	*(*int32)(unsafe.Add(mBase, uint32(v38)+72)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v38)+80)) = v260
	*(*int32)(unsafe.Add(mBase, uint32(v38)+84)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v38)+88)) = v263
	*(*int32)(unsafe.Add(mBase, uint32(v38)+92)) = v266
	*(*int32)(unsafe.Add(mBase, uint32(v38)+96)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v38)+100)) = v259
	*(*int32)(unsafe.Add(mBase, uint32(v38)+104)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v38)+108)) = v261
	F_sequence_close(m, v1353, int32(0))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		v1457 = v271
		goto L10
	} else {
		goto L122
	}
L122:
	;
	goto L9
L123:
	;
	v1469 = int32(v1465)
	m.G0 = v1457
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1469)+4))
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1469)))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1472)))
	if v38+int32(44) == v1476 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	m.ExcPending = 1
	goto L132
L125:
	;
	if v1479 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1472)+4))
	v1479 = v1478
	goto L128
L127:
	;
	v1479 = int32(0)
	goto L128
L128:
	;
	goto L125
L129:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v38)+108))
	v1481 = *(*int32)(unsafe.Add(mBase, uint32(v38)+104))
	v1482 = *(*int32)(unsafe.Add(mBase, uint32(v38)+100))
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v38)+96))
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v38)+92))
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(v38)+88))
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(v38)+84))
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v38)+80))
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v38)+76))
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v38)+72))
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v38)+68))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v38)+56))
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
	v48 = v1483
	v49 = v1486
	v50 = v1488
	v51 = v1482
	v52 = v1494
	v53 = v1487
	v54 = v1480
	v55 = v1481
	v56 = v1485
	v57 = v1492
	v58 = v1490
	v59 = v1489
	v60 = v1484
	v61 = v1495
	v62 = v1493
	v63 = v1491
	v66 = v1479
	v72 = v1457
	v73 = v1471
	goto L5
L130:
	;
	goto L131
L131:
	;
	F___wasm_longjmp(m, v1472, v1471)
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	return
L133:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ApplyLauncherMain(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int64
	_ = v389
	var v390 int64
	_ = v390
	var v399 int64
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v421 int64
	_ = v421
	var v426 int32
	_ = v426
	var v429 int64
	_ = v429
	var v437 int64
	_ = v437
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int64
	_ = v446
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v552 int32
	_ = v552
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v21 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v21 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errmsg_internal(m, int32(444043), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_before_shmem_exit(m, int32(989), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errfinish(m, int32(494873), int32(1135), int32(278432))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	v39 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v39
	v42 = int32(914)
	v44 = m.G0
	v46 = v44 - int32(144)
	m.G0 = v46
	switch int32(916) {
	case 0, 2:
		v56 = v42
		goto L10
	default:
		goto L11
	}
L9:
	;
	v84 = int32(295)
	v86 = m.G0
	v88 = v86 - int32(144)
	m.G0 = v88
	switch int32(297) {
	case 0, 2:
		v98 = v84
		goto L23
	default:
		goto L24
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v56
	F_sigemptyset(m, v46+int32(8))
	mBase = m.M
	goto L13
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[625])) = v42
	v56 = int32(4730)
	goto L10
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+136)) = int32(268435456)
	v68 = v46 + int32(4)
	goto L17
L15:
	;
	m.G0 = v46 + int32(144)
	goto L9
L17:
	;
	goto L18
L18:
	;
	if v68 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v79 = F___memcpy(m, int32(4655052), v68, int32(140))
	mBase = m.M
	goto L21
L20:
	;
	goto L21
L21:
	;
	goto L15
L22:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L35
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v98
	F_sigemptyset(m, v88+int32(8))
	mBase = m.M
	goto L26
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[622])) = v84
	v98 = int32(4730)
	goto L23
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+136)) = int32(268435456)
	v110 = v88 + int32(4)
	goto L30
L28:
	;
	m.G0 = v88 + int32(144)
	goto L22
L30:
	;
	goto L31
L31:
	;
	if v110 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v121 = F___memcpy(m, int32(4657012), v110, int32(140))
	mBase = m.M
	goto L34
L33:
	;
	goto L34
L34:
	;
	goto L28
L35:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128)+192)))
	if v129&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L50
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L47
	}
L38:
	;
	v132 = int32(0)
	F_InitPostgres(m, v132, v132, v132, v132, v132, v132)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _consts[84]))
	if v141 != int32(1) {
		goto L37
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, _consts[84])) = int32(2)
	goto L36
L43:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(258954), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(494819), int32(869), int32(255059))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	F_errmsg(m, int32(220738), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(494819), int32(879), int32(255059))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v191 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v191 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v194 = int32(0)
	v196 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v201 = F_AllocSetContextCreateInternal(m, v196, int32(73675), v194, int32(8192), int32(8388608))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v203 = int32(4489440)
	v204 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v201
	F_StartTransactionCommand(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v211 = F_table_open(m, int32(6100), int32(1))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v213 = int32(0)
	v215 = F_table_beginscan_catalog(m, v211, v213, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v221 = v194
	goto L60
L60:
	;
	v231 = F_heap_getnext(m, v215)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+188))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+12))
	m.T0[v262].(func(*base.Module, int32))(m, v215)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L69
	}
L62:
	;
	if v231 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v233 = int32(4489440)
	v234 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+16))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+22)))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v201
	v240 = F_palloc0(m, int32(56))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	goto L61
L66:
	;
	v242 = v235 + v236
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)))
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v245
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v242)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v240)+20)) = v247
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+84)))
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+25)) = uint8(v249)
	v253 = F_pstrdup(m, v242+int32(16))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+16)) = v253
	v256 = F_lappend(m, v221, v240)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v234
	v221 = v256
	goto L60
L69:
	;
	F_sequence_close(m, v211, int32(1))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	if v221 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v204
	F_MemoryContextDelete(m, v201)
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L1
	} else {
		goto L125
	}
L73:
	;
	v508 = int32(180000)
	goto L72
L74:
	;
	goto L75
L75:
	;
	v273 = int32(0)
	v274 = int32(180000)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v275 <= v273 {
		v508 = v274
		goto L72
	} else {
		goto L76
	}
L76:
	;
	v281 = v274
	v284 = v273
	goto L77
L77:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v221)+12))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v292+v284<<(uint(int32(2))%32))))
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296)+25)))
	if v297 != int32(1) {
		v490 = v281
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v508 = v490
	goto L72
L79:
	;
	v502 = v284 + int32(1)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v221)+4))
	if v502 < v503 {
		v281 = v490
		v284 = v502
		goto L77
	} else {
		goto L124
	}
L80:
	;
	v301 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v305 = F_LWLockAcquire(m, v301+int32(5504), int32(1))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v308 = *(*int32)(unsafe.Add(mBase, _consts[671]))
	if int32(0) < v308 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v482 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v482+int32(5504))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L123
	}
L83:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v314 = *(*int32)(unsafe.Add(mBase, _consts[670]))
	v317 = int32(0)
	goto L86
L84:
	;
	goto L85
L85:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v363+int32(5504))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L94
	}
L86:
	;
	v333 = v314 + int32(16) + v317*int32(112)
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+16)))
	if v334 != int32(1) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L85
L88:
	;
	v346 = v317 + int32(1)
	if v346 != v308 {
		v317 = v346
		goto L86
	} else {
		goto L93
	}
L89:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v333)))
	if v337 == int32(3) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v333)+32))
	if v340 != v311 {
		goto L88
	} else {
		goto L91
	}
L91:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v333)+36))
	if v342 == int32(0) {
		goto L82
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	goto L87
L94:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v368
	F_logicalrep_launcher_attach_dshmem(m)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v373 = *(*int32)(unsafe.Add(mBase, _consts[672]))
	v377 = F_dshash_find(m, v373, v17+int32(4), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	v478 = v442 - v440
	if v281 < v478 {
		goto L120
	} else {
		goto L121
	}
L97:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v448
	F_logicalrep_launcher_attach_dshmem(m)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L112
	}
L98:
	;
	if v377 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v384 = m.G0
	v385 = int32(16)
	v386 = v384 - v385
	m.G0 = v386
	F___gettimeofday(m, v386)
	mBase = m.M
	v389 = *(*int64)(unsafe.Add(mBase, uint32(v386)))
	v390 = int64(*(*int32)(unsafe.Add(mBase, uint32(v386)+8)))
	m.G0 = v386 + v385
	goto L102
L100:
	;
	goto L101
L101:
	;
	v399 = *(*int64)(unsafe.Add(mBase, uint32(v377)+8))
	v401 = *(*int32)(unsafe.Add(mBase, _consts[672]))
	F_dshash_release_lock(m, v401, v377)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L103
	}
L102:
	;
	v446 = v390 + v389*int64(1000000) - int64(946684800000000)
	goto L97
L103:
	;
	v407 = m.G0
	v408 = int32(16)
	v409 = v407 - v408
	m.G0 = v409
	F___gettimeofday(m, v409)
	mBase = m.M
	v412 = *(*int64)(unsafe.Add(mBase, uint32(v409)))
	v413 = int64(*(*int32)(unsafe.Add(mBase, uint32(v409)+8)))
	m.G0 = v409 + v408
	v421 = v413 + v412*int64(1000000) - int64(946684800000000)
	goto L104
L104:
	;
	if v399 == int64(0) {
		v446 = v421
		goto L97
	} else {
		goto L105
	}
L105:
	;
	if v421 <= v399 {
		v440 = int32(0)
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _consts[392]))
	if v440 < v442 {
		goto L96
	} else {
		goto L111
	}
L107:
	;
	goto L106
L108:
	;
	v426 = int32(2147483647)
	v429 = v421 - v399
	if base.B2i32(int64(0) < v399)^base.B2i32(v429 < v421) != 0 {
		v440 = v426
		goto L107
	} else {
		goto L109
	}
L109:
	;
	if int64(2147483646000) < v429 {
		v440 = v426
		goto L107
	} else {
		goto L110
	}
L110:
	;
	v437 = base.I64_div_s(v429+int64(999), int64(1000))
	v440 = base.I32_wrap_i64(v437)
	goto L107
L111:
	;
	v446 = v421
	goto L97
L112:
	;
	v453 = *(*int32)(unsafe.Add(mBase, _consts[672]))
	v458 = F_dshash_find_or_insert(m, v453, v17+int32(12), v17+int32(11))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v458)+8)) = v446
	v462 = *(*int32)(unsafe.Add(mBase, _consts[672]))
	F_dshash_release_lock(m, v462, v458)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v296)+4))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v296)))
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v296)+16))
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v296)+20))
	v470 = int32(0)
	v472 = F_logicalrep_worker_launch(m, int32(2), v466, v467, v468, v469, v470, v470)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	if v472 != 0 {
		v490 = v281
		goto L79
	} else {
		goto L116
	}
L116:
	;
	v475 = *(*int32)(unsafe.Add(mBase, _consts[392]))
	if v281 < v475 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v477 = v281
	goto L119
L118:
	;
	v477 = v475
	goto L119
L119:
	;
	v490 = v477
	goto L79
L120:
	;
	v480 = v281
	goto L122
L121:
	;
	v480 = v478
	goto L122
L122:
	;
	v490 = v480
	goto L79
L123:
	;
	v490 = v281
	goto L79
L124:
	;
	goto L78
L125:
	;
	v524 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v527 = F_WaitLatch(m, v524, int32(41), v508, int32(83886088))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L127
	}
L126:
	;
	v544 = *(*int32)(unsafe.Add(mBase, _consts[346]))
	if v544 == int32(0) {
		goto L50
	} else {
		goto L132
	}
L127:
	;
	if v527&int32(1) == int32(0) {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v534))) = int32(0)
	goto L129
L129:
	;
	v538 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v538 == int32(0) {
		goto L126
	} else {
		goto L130
	}
L130:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	goto L126
L132:
	;
	*(*int32)(unsafe.Add(mBase, _consts[346])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	goto L50
}
func F_AtAbort_Twophase(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L19
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	v19 = F_LWLockAcquire(m, v15+int32(2304), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	m.G0 = v10 + int32(16)
	return
L5:
	;
	return
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[163]))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+44)))
	if v23 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	F_LWLockRelease(m, v69+int32(2304))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L5
	} else {
		goto L18
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[164]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v28 <= int32(0) {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = int32(-1)
	goto L7
L11:
	;
	v32 = v27 + int32(8)
	v34 = int32(0)
	goto L12
L12:
	;
	v42 = v32 + v34<<(uint(int32(2))%32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v43 != v22 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v49 = v28 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v49
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v32+v49<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v22
	goto L7
L14:
	;
	v46 = v34 + int32(1)
	if v28 != v46 {
		v34 = v46
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	goto L13
L17:
	;
	goto L1
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[163])) = int32(0)
	goto L4
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22
	F_errmsg_internal(m, int32(25149), v10)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(498113), int32(650), int32(111782))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_AuxiliaryProcessMainCommon(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v143 int32
	_ = v143
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	v6 = *(*int32)(unsafe.Add(mBase, _consts[613]))
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_MemoryContextDelete(m, v6)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	goto L7
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, _consts[613])) = int32(0)
	goto L3
L6:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[614])) = uint8(v19)
	v22 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	if v22 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[198]))
	v17 = F_GetBackendTypeDesc(m, v16)
	mBase = m.M
	goto L9
L9:
	;
	goto L6
L10:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[138]))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+44)) = v124
	v127 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	v128 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v128
	v130 = int32(4412876)
	*(*int32)(unsafe.Add(mBase, _consts[137])) = v121
	v134 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	v138 = base.I32_div_s(v121-v135, int32(640))
	*(*int32)(unsafe.Add(mBase, _consts[109])) = v138
	v140 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v121))) = v140
	v143 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+612)) = v128
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+608)) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+16)) = v128
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+124)) = uint8(v128)
	*(*int32)(unsafe.Add(mBase, uint32(v143)+120)) = v128
	*(*int32)(unsafe.Add(mBase, uint32(v143)+52)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v143)+36)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v143)+92)) = v140
	*(*uint16)(unsafe.Add(mBase, uint32(v143)+74)) = uint16(v128)
	*(*int64)(unsafe.Add(mBase, uint32(v143)+56)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v143)+112)) = v140
	*(*int64)(unsafe.Add(mBase, uint32(v143-int32(-64)))) = v140
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+72)) = uint8(v128)
	v173 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	F_OwnLatch(m, v173+int32(20))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L4
	} else {
		goto L41
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L38
	}
L12:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L35
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[616]))
	if v26 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v30 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _consts[101])))
	if v32 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_RegisterPostmasterChildActive(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = int32(1)
	if v39 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	F_s_lock(m, v43, int32(499295), int32(640), int32(129936))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[150]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+68))
	*(*int32)(unsafe.Add(mBase, _consts[333])) = v51
	goto L24
L23:
	;
	goto L22
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[616]))
	v56 = int32(0)
	goto L25
L25:
	;
	v62 = v55 + v56*int32(640)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+44))
	if v63 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v80
	F_errstart_cold(m, int32(22), v80)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L32
	}
L27:
	;
	v121 = v62
	v122 = v56
	goto L10
L28:
	;
	goto L29
L29:
	;
	v67 = v56 | int32(1)
	v70 = v55 + v67*int32(640)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+44))
	if v71 == int32(0) {
		v121 = v70
		v122 = v67
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v75 = v56 + int32(2)
	if v75 != int32(38) {
		v56 = v75
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	F_errmsg_internal(m, int32(359767), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(499295), int32(656), int32(129936))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errmsg_internal(m, int32(437677), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(499295), int32(625), int32(129936))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_errmsg_internal(m, int32(68553), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(499295), int32(628), int32(129936))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_SwitchToSharedLatch(m)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	*(*int32)(unsafe.Add(mBase, _consts[39])) = v181 + int32(548)
	goto L43
L43:
	;
	F_on_shmem_exit(m, int32(1115), v122)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_BaseInit(m)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v194 = int32(0)
	F_ProcSignalInit(m, v194, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_CreateAuxProcessResourceOwner(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	F_pgstat_beinit(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L48
	}
L48:
	;
	F_pgstat_bestart_initial(m)
	mBase = m.M
	F_pgstat_bestart_final(m)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	F_before_shmem_exit(m, int32(922), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, _consts[84])) = int32(2)
	return
}
func F_acldefault(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int64
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v25 int64
	_ = v25
	var v33 int64
	_ = v33
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v84 int32
	_ = v84
	v3 = int32(0)
	v8 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = int32(1)
	switch l0 - int32(6) {
	case 0:
		v55 = int32(0)
		v56 = v15
		v57 = int32(1)
		v58 = int64(0)
		v59 = v8
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
			F_errmsg_internal(m, int32(484222), v12)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(496905), int32(872), int32(98126))
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
	case 3:
		v55 = int32(2)
		v56 = int32(0)
		v57 = v3
		v58 = int64(3584)
		v59 = int64(3072)
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 6, 15, 43:
		v33 = int64(256)
		v55 = int32(2)
		v56 = int32(0)
		v57 = v3
		v58 = v33
		v59 = v33
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 10, 11:
		v55 = v15
		v56 = v15
		v57 = v3
		v58 = int64(256)
		v59 = v8
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 13:
		v25 = int64(128)
		v55 = int32(2)
		v56 = int32(0)
		v57 = v3
		v58 = v25
		v59 = v25
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 16:
		v55 = v15
		v56 = v15
		v57 = v3
		v58 = int64(6)
		v59 = v8
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 21:
		v55 = v15
		v56 = v15
		v57 = v3
		v58 = int64(12288)
		v59 = v8
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 30:
		v55 = v15
		v56 = v15
		v57 = v3
		v58 = int64(768)
		v59 = v8
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 31:
		v55 = v15
		v56 = v15
		v57 = v3
		v58 = int64(262)
		v59 = v8
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 35:
		v55 = v15
		v56 = v15
		v57 = v3
		v58 = int64(16511)
		v59 = v8
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 36:
		v55 = v15
		v56 = v15
		v57 = v3
		v58 = int64(512)
		v59 = v8
		v63 = v55<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v55
			if v56 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v59
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v57 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v58
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	}
}
func F_aclexplode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int64
	_ = v134
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v211 int32
	_ = v211
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
	var v221 int64
	_ = v221
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	if v23 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_check_acl(m, v18)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	goto L15
L6:
	;
	v28 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v30 = int32(4489440)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v33
	v36 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v36, int32(1), int32(207716), int32(26), int32(-1), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v36, int32(2), int32(409252), int32(26), int32(-1), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v36, int32(3), int32(366143), int32(25), int32(-1), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v36, int32(4), int32(391082), int32(16), int32(-1), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v66 = F_BlessTupleDesc(m, v36)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v66
	v70 = F_palloc(m, int32(8))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v70))) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v70
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v31
	goto L5
L15:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	if v83 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v91 = v83
	goto L18
L17:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v91 = (v84<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L18:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v95 = v18 + int32(16)
	v97 = v82 + int32(4)
	v99 = v93
	goto L21
L19:
	;
	m.G0 = v15 + int32(32)
	return v238
L20:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L67
	}
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v110 <= v99 {
		goto L20
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v141
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v143
	v145 = base.I32_wrap_i64(v130)
	if v145 <= int32(127) {
		goto L41
	} else {
		goto L42
	}
L23:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v114 = v112 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v114
	if v114 == int32(15) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v119 = v99 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v119
	v121 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v121
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v124 <= v119 {
		goto L20
	} else {
		goto L27
	}
L25:
	;
	v126 = v99
	v127 = v114
	goto L26
L26:
	;
	v129 = base.I64_extend_i32_u(v127)
	v130 = int64(1) << (uint(v129) % 64)
	v133 = v91 + v18 + v126<<(uint(int32(4))%32)
	v134 = *(*int64)(unsafe.Add(mBase, uint32(v133)+8))
	if base.I32_wrap_i64(v130&v134) == int32(0) {
		v99 = v126
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v126 = v119
	v127 = v121
	goto L26
L28:
	;
	goto L22
L29:
	;
	v202 = F_cstring_to_text(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L64
	}
L30:
	;
	v201 = int32(521510)
	goto L29
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L61
	}
L32:
	;
	v201 = int32(531016)
	goto L29
L33:
	;
	v201 = int32(520871)
	goto L29
L34:
	;
	v201 = int32(537839)
	goto L29
L35:
	;
	v201 = int32(537163)
	goto L29
L36:
	;
	v201 = int32(525284)
	goto L29
L37:
	;
	v201 = int32(522904)
	goto L29
L38:
	;
	v201 = int32(538208)
	goto L29
L39:
	;
	v201 = int32(537520)
	goto L29
L40:
	;
	v201 = int32(538186)
	goto L29
L41:
	;
	switch v145 - int32(1) {
	case 0:
		v201 = int32(516883)
		goto L29
	case 1:
		goto L30
	default:
		goto L31
	case 3:
		goto L40
	case 7:
		goto L39
	case 15:
		goto L38
	case 31:
		goto L37
	case 63:
		goto L36
	}
L42:
	;
	goto L43
L43:
	;
	if v145 <= int32(2047) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if v145 <= int32(511) {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	if v145 <= int32(8191) {
		goto L54
	} else {
		goto L55
	}
L47:
	;
	if v145 == int32(128) {
		goto L35
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if v145 == int32(512) {
		goto L34
	} else {
		goto L52
	}
L50:
	;
	if v145 != int32(256) {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	v201 = int32(540858)
	goto L29
L52:
	;
	if v145 != int32(1024) {
		goto L31
	} else {
		goto L53
	}
L53:
	;
	v201 = int32(508022)
	goto L29
L54:
	;
	if v145 == int32(2048) {
		goto L33
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	if v145 == int32(8192) {
		goto L32
	} else {
		goto L59
	}
L57:
	;
	if v145 != int32(4096) {
		goto L31
	} else {
		goto L58
	}
L58:
	;
	v201 = int32(520453)
	goto L29
L59:
	;
	if v145 != int32(16384) {
		goto L31
	} else {
		goto L60
	}
L60:
	;
	v201 = int32(529711)
	goto L29
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v145
	F_errmsg_internal(m, int32(480945), v15)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(496905), int32(1770), int32(329398))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v202
	v205 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v133)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = base.I32_wrap_i64(int64(base.Ui64(v205)>>(uint(v129)%64))) & int32(1)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v81)+28))
	v216 = F_heap_form_tuple(m, v211, v15+int32(16), v15+int32(12))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
	v219 = F_HeapTupleHeaderGetDatum(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v221 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	*(*int64)(unsafe.Add(mBase, uint32(v81))) = v221 + int64(1)
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+20)) = int32(1)
	v238 = v219
	goto L19
L67:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+20)) = int32(2)
	v235 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v235)
	v238 = int32(0)
	goto L19
}
func F_aclitemout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v56 int32
	_ = v56
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v149 int32
	_ = v149
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v175 int64
	_ = v175
	var v178 int64
	_ = v178
	var v180 int64
	_ = v180
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v192 int32
	_ = v192
	var v193 int64
	_ = v193
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v241 int32
	_ = v241
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_palloc(m, int32(293))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v23)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v25 == v23 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v149 = v19
	goto L33
L4:
	;
	v29 = F_SearchSysCache1(m, int32(11), v25)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	v35 = v33 + int32(4)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+4)))
	if v37 == int32(0) {
		v85 = int32(1)
		v93 = v19
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v130
	v135 = F_pg_sprintf(m, v19, int32(59399), v15+int32(16))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L32
	}
L9:
	;
	v94 = v93
	v95 = v35
	goto L21
L10:
	;
	v40 = v35
	v42 = v37
	goto L11
L11:
	;
	if int32(0) <= base.I32_extend8_s(v42) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v85 = v77
	v93 = v19
	goto L9
L13:
	;
	v77 = int32(1)
	v79 = v40 + v77
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v80 != 0 {
		v40 = v79
		v42 = v80
		goto L11
	} else {
		goto L20
	}
L14:
	;
	v56 = v42 & int32(255)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v72 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v72)
	v85 = int32(0)
	v93 = v19 + int32(1)
	goto L9
L17:
	;
	if v56 == int32(95) {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if base.B2i32(base.Ui32(v56-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v56|int32(32)-int32(97)) < base.Ui32(int32(26))) != 0 {
		goto L13
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	goto L12
L21:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	if v106 != int32(34) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v123))) = uint8(v124)
	v126 = int32(1)
	v94 = v123 + v126
	v95 = v95 + v126
	goto L21
L24:
	;
	if v106 != 0 {
		v123 = v94
		v124 = v106
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v118 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v118)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
	v123 = v94 + int32(1)
	v124 = v122
	goto L23
L27:
	;
	if v85 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v113 = v94
	goto L30
L29:
	;
	v109 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v94))) = uint8(v109)
	v113 = v94 + int32(1)
	goto L30
L30:
	;
	v114 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v113))) = uint8(v114)
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L3
L32:
	;
	goto L3
L33:
	;
	v162 = v149 + int32(1)
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149))))
	if v163 != 0 {
		v149 = v162
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v164 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v164)
	v166 = v162
	v175 = int64(0)
	goto L36
L35:
	;
	goto L34
L36:
	;
	v178 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v180 = int64(1) << (uint(v175) % 64)
	if v178&v180 != int64(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v208 = int32(47)
	*(*uint16)(unsafe.Add(mBase, uint32(v203))) = uint16(v208)
	v210 = int32(1)
	v212 = v203 + v210
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v215 = F_SearchSysCache1(m, int32(11), v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L46
	}
L38:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v175))+uint32(_consts[947]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v187)
	v189 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v192 = v166 + int32(1)
	v193 = v189
	goto L40
L39:
	;
	v192 = v166
	v193 = v178
	goto L40
L40:
	;
	if int64(base.Ui64(v193)>>(uint(int64(32))%64))&v180 != int64(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v199 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v192))) = uint8(v199)
	v203 = v192 + int32(1)
	goto L43
L42:
	;
	v203 = v192
	goto L43
L43:
	;
	v205 = v175 + int64(1)
	if v205 != int64(15) {
		v166 = v203
		v175 = v205
		goto L36
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	m.G0 = v15 + int32(32)
	return v19
L46:
	;
	if v215 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+22)))
	v219 = v217 + v218
	v221 = v219 + int32(4)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+4)))
	if v222 == int32(0) {
		v268 = v212
		v273 = v210
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v315
	v318 = F_pg_sprintf(m, v212, int32(59399), v15)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L70
	}
L50:
	;
	v280 = v268
	v281 = v221
	goto L59
L51:
	;
	v228 = v222
	v229 = v221
	goto L52
L52:
	;
	if base.I32_extend8_s(v228) < int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v262 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)) = uint8(v262)
	v268 = v203 + int32(2)
	v273 = int32(0)
	goto L50
L54:
	;
	goto L53
L55:
	;
	v241 = v228 & int32(255)
	goto L56
L56:
	;
	if base.B2i32(base.B2i32(base.Ui32(v241-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v241|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v241 != int32(95)) != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v259 = v229 + int32(1)
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259))))
	if v260 != 0 {
		v228 = v260
		v229 = v259
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v268 = v212
	v273 = v210
	goto L50
L59:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v291 != int32(34) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v308)
	v311 = int32(1)
	v280 = v309 + v311
	v281 = v281 + v311
	goto L59
L62:
	;
	if v291 != 0 {
		v308 = v291
		v309 = v280
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v303 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v303)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	v308 = v307
	v309 = v280 + int32(1)
	goto L61
L65:
	;
	if v273 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v298 = v280
	goto L68
L67:
	;
	v294 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v280))) = uint8(v294)
	v298 = v280 + int32(1)
	goto L68
L68:
	;
	v299 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v298))) = uint8(v299)
	F_ReleaseCatCache(m, v215)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L45
L70:
	;
	goto L45
}
func F_aclnewowner(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v159 int32
	_ = v159
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v174 int32
	_ = v174
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	F_check_acl(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v22 == int32(0) {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v32 = (v25<<(uint(int32(3))%32) + int32(23)) & int32(-8)
		} else {
			v32 = v22
		}
		if int32(0) <= v21 {
			v36 = v21 << (uint(int32(4)) % 32)
			v38 = v36 + int32(24)
			v39 = F_palloc0(m, v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = int32(1033)
				*(*int64)(unsafe.Add(mBase, uint32(v39)+4)) = int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = v38 << (uint(int32(2)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v21
				v52 = v39 + int32(24)
				if v36 != 0 {
					v54 = F__emscripten_memcpy_bulkmem(m, v52, l0+v32, v36)
					mBase = m.M
					v55 = v54
				} else {
					v55 = v52
				}
				if v21 == int32(0) {
				} else {
					v58 = int32(0)
					v60 = v55
					v63 = v58
					v67 = v58
					for {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
						if l1 == v72 {
							*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = l2
							v77 = v63
						} else {
							v77 = base.B2i32(l2 == v72) | v63
						}
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
						if l1 == v78 {
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = l2
							v83 = v77
						} else {
							v83 = base.B2i32(l2 == v78) | v77
						}
						v87 = v67 + int32(1)
						if v87 != v21 {
							v60 = v60 + int32(16)
							v63 = v83
							v67 = v87
							continue
						} else {
							break
						}
						break
					}
					if v83&int32(1) == int32(0) {
					} else {
						v95 = v55
						v98 = int32(0)
						v102 = int32(0)
						for {
							v107 = v98 + int32(1)
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v95)+8))
							if v108 != int64(0) {
								if v107 < v21 {
									v112 = v95
									v114 = v107
									for {
										v125 = v112 + int32(16)
										v127 = v112 + int32(24)
										v128 = *(*int64)(unsafe.Add(mBase, uint32(v127)))
										if v128 == int64(0) {
										} else {
											v131 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
											v132 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
											if v131 != v132 {
											} else {
												v134 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
												v135 = *(*int32)(unsafe.Add(mBase, uint32(v112)+20))
												if v134 != v135 {
												} else {
													v137 = *(*int64)(unsafe.Add(mBase, uint32(v95)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v95)+8)) = v137 | v128
													*(*int64)(unsafe.Add(mBase, uint32(v127))) = int64(0)
												}
											}
										}
										v143 = v114 + int32(1)
										if v143 != v21 {
											v112 = v125
											v114 = v143
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v159 = v55 + v102<<(uint(int32(4))%32)
								v160 = *(*int64)(unsafe.Add(mBase, uint32(v95)))
								*(*int64)(unsafe.Add(mBase, uint32(v159))) = v160
								v162 = *(*int64)(unsafe.Add(mBase, uint32(v95)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v159)+8)) = v162
								v174 = v102 + int32(1)
							} else {
								v174 = v102
							}
							if v107 != v21 {
								v95 = v95 + int32(16)
								v98 = v107
								v102 = v174
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v39)+16)) = v174
						*(*int32)(unsafe.Add(mBase, uint32(v39))) = v174<<(uint(int32(6))%32) + int32(96)
					}
				}
				m.G0 = v15 + int32(16)
				return v39
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v21
				F_errmsg_internal(m, int32(483043), v15)
				mBase = m.M
				v210 = m.ExcPending
				if v210 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496905), int32(433), int32(308051))
					mBase = m.M
					v215 = m.ExcPending
					if v215 != 0 {
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
func F_addArc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	switch v17 + int32(4) {
	case 0:
		goto L4
	case 1:
		goto L1
	default:
		goto L3
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v28 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v25 == int32(-4) {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	if l3 != int32(-4) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v22 != int32(-4) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	goto L1
L7:
	;
	goto L2
L8:
	;
	v88 = F_palloc(m, int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L26
	} else {
		goto L27
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v31 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v34 = int32(0)
	if v34 < v31 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = v31
	goto L13
L12:
	;
	v37 = v34
	goto L13
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v46 = int32(0)
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v39+v46<<(uint(int32(2))%32))))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v56 != v38 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L8
L16:
	;
	v73 = v46 + int32(1)
	if v73 != v37 {
		v46 = v73
		goto L14
	} else {
		goto L25
	}
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v58 == int32(-3) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v61 != int32(-3) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v61 != v64 {
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v58 == v68 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v58 != v66 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	goto L1
L24:
	;
	goto L16
L25:
	;
	goto L15
L26:
	;
	return
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v94 = F_hash_search(m, v90, l4, int32(1), v15+int32(15))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v96 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v99 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v99
	v101 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v94)+12)) = v101
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v103 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+36)) = v99
	*(*int64)(unsafe.Add(mBase, uint32(v94)+28)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = v103 ^ int32(-1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v115 = F_lappend(m, v114, v94)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L26
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+12)) = v94
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = v120
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v122
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v126 = F_lappend(m, v125, v88)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L26
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v115
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v126
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v129 + int32(1)
	goto L1
}
func F_add_new_columns_to_pathtarget(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = v3
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v14<<(uint(int32(2))%32))))
	v22 = F_list_member(m, v16, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v52 = v14 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v52 < v53 {
		v14 = v52
		goto L4
	} else {
		goto L19
	}
L7:
	;
	return
L8:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v25 = F_lappend(m, v24, v21)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v25 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v44 != int32(2) {
		goto L6
	} else {
		goto L18
	}
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v31 = v29
	goto L16
L15:
	;
	v31 = int32(0)
	goto L16
L16:
	;
	v33 = v31 << (uint(int32(2)) % 32)
	v34 = F_repalloc(m, v28, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v33+v34-int32(4)))) = int32(0)
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = int32(0)
	goto L6
L19:
	;
	goto L5
}
func F_add_nullingrels_if_needed(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	if v10 == int32(0) {
		v120 = l1
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L36
	} else {
		goto L51
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L36
	} else {
		goto L48
	}
L3:
	;
	m.G0 = v8 + int32(16)
	return v120
L4:
	;
	v13 = int32(0)
	if l1 == v13 {
		v80 = v13
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L6:
	;
	v86 = v80
	goto L5
L7:
	;
	v18 = l1
	goto L8
L8:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	switch v23 - int32(6) {
	case 0:
		goto L14
	case 1, 2, 3, 4, 5, 6, 7, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 24, 25, 26, 27, 28, 29, 30, 31:
		v80 = v13
		goto L6
	case 9:
		goto L13
	case 21, 22, 23:
		goto L12
	case 32:
		goto L11
	default:
		goto L15
	}
L9:
	;
	v80 = v13
	goto L6
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v75 != 0 {
		v18 = v75
		goto L8
	} else {
		goto L32
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v45 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	v74 = v18 + int32(4)
	goto L10
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v36 != int32(2) {
		v80 = v13
		goto L6
	} else {
		goto L19
	}
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v32 != v33 {
		v80 = v13
		goto L6
	} else {
		goto L18
	}
L15:
	;
	if v23 != int32(319) {
		v80 = v13
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v28 != v29 {
		v80 = v13
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v86 = int32(1)
	goto L5
L18:
	;
	v86 = int32(1)
	goto L5
L19:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v39 == int32(0) {
		v80 = v13
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v74 = v42
	goto L10
L21:
	;
	v86 = int32(1)
	goto L5
L22:
	;
	goto L23
L23:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v50 <= int32(0) {
		v80 = int32(1)
		goto L6
	} else {
		goto L24
	}
L24:
	;
	v53 = int32(0)
	if v53 < v50 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v57 = v50
	goto L27
L26:
	;
	v57 = v53
	goto L27
L27:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v59 = v53
	goto L28
L28:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v58+v59<<(uint(int32(2))%32))))
	v68 = F_is_standard_join_alias_expression(m, v67, l2)
	mBase = m.M
	if v68 == int32(0) {
		v80 = v68
		goto L6
	} else {
		goto L30
	}
L29:
	;
	v80 = v68
	goto L6
L30:
	;
	v72 = v59 + int32(1)
	if v72 != v57 {
		v59 = v72
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	goto L9
L33:
	;
	F_adjust_standard_join_alias_expression(m, l1, l2)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	if l0 == int32(0) {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	return int32(0)
L37:
	;
	v120 = l1
	goto L3
L38:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l0
	v96 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v96
	v102 = F_query_or_expression_tree_walker_impl(m, l1, int32(896), v8+int32(4), v96)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v104 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v112 = v104
	goto L42
L41:
	;
	if v93 != 0 {
		goto L2
	} else {
		goto L43
	}
L42:
	;
	v113 = F_make_placeholder_expr(m, l0, l1, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L36
	} else {
		goto L46
	}
L43:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v107 = F_get_relids_for_join(m, v105, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v110 = F_bms_del_member(m, v107, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v112 = v110
	goto L42
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+20)) = v93
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v117 = F_bms_copy(m, v116)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L36
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+12)) = v117
	v120 = v113
	goto L3
L48:
	;
	F_errmsg_internal(m, int32(268983), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(494980), int32(1200), int32(461298))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L36
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	F_errmsg_internal(m, int32(268983), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L36
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(494980), int32(1215), int32(461298))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L36
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_add_parameter_name(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1254]))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	if v142 != 0 {
		goto L38
	} else {
		goto L39
	}
L2:
	;
	goto L1
L3:
	;
	v142 = v36
	goto L2
L5:
	;
	v142 = int32(0)
	goto L2
L6:
	;
	goto L7
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v31 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v36 = v10
	v38 = v31
	goto L12
L10:
	;
	goto L11
L11:
	;
	goto L19
L12:
	;
	v43 = F_strcmp(m, v36+int32(12), l2)
	mBase = m.M
	if v43 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 != 0 {
		v36 = v52
		v38 = v53
		goto L12
	} else {
		goto L18
	}
L15:
	;
	if base.B2i32(v38 == int32(1))&int32(0) != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	goto L3
L18:
	;
	goto L13
L19:
	;
	goto L5
L38:
	;
	F_errstart_cold(m, int32(21), int32(549150))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	F_plpgsql_ns_additem(m, l0, l1, l2)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L41
	} else {
		goto L46
	}
L41:
	;
	return
L42:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
	F_errmsg(m, int32(414572), v7)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(495264), int32(930), int32(378212))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	m.G0 = v7 + int32(16)
	return
}
func F_add_size(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v3 = l0 + l1
	if base.Ui32(v3) < base.Ui32(l1) {
		F_errstart_cold(m, int32(21), int32(0))
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(112414), int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(496575), int32(502), int32(340935))
					v22 = m.ExcPending
					if v22 != 0 {
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
		return v3
	}
}
func F_adjust_appendrel_attrs_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
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
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
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
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v594 int64
	_ = v594
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
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
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
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
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	if l0 == v3 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L18
	} else {
		goto L238
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L18
	} else {
		goto L235
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L18
	} else {
		goto L232
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L18
	} else {
		goto L228
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L18
	} else {
		goto L224
	}
L6:
	;
	m.G0 = v15 + int32(32)
	return v623
L7:
	;
	v623 = v3
	goto L6
L8:
	;
	goto L9
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v21 - int32(318) {
	case 0:
		goto L11
	case 1:
		goto L12
	default:
		goto L13
	}
L10:
	;
	v615 = F_expression_tree_mutator_impl(m, l0, int32(854), l1)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L18
	} else {
		goto L223
	}
L11:
	;
	v315 = F_palloc0(m, int32(168))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L18
	} else {
		goto L119
	}
L12:
	;
	v261 = F_expression_tree_mutator_impl(m, l0, int32(854), l1)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L18
	} else {
		goto L97
	}
L13:
	;
	if v21 != int32(58) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v21 != int32(6) {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v231 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L18
	} else {
		goto L89
	}
L17:
	;
	v28 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	if v32 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v623 = v28
	goto L6
L21:
	;
	goto L22
L22:
	;
	if v19 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v623 = v28
	goto L6
L24:
	;
	goto L25
L25:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v39 = int32(0)
	goto L27
L26:
	;
	if v35 != int32(-4) {
		goto L69
	} else {
		goto L70
	}
L27:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v20+v39<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v53 != v35 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v59 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+40)) = uint16(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v28)+4)) = v58
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+8)))
	if v59 < v65 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v56 = v39 + int32(1)
	if v19 != v56 {
		v39 = v56
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	goto L26
L33:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	if v68 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	if v65 != 0 {
		goto L51
	} else {
		goto L52
	}
L36:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = v69
	goto L38
L37:
	;
	v70 = v59
	goto L38
L38:
	;
	if v70 < v65 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72+v65<<(uint(int32(2))%32)-int32(4))))
	v79 = F_copyObjectImpl(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	if v79 == int32(0) {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v84 == int32(6) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+32)) = v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v79)+24))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	v90 = F_bms_add_members(m, v88, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L18
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	if v83 != 0 {
		goto L3
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = v90
	v623 = v79
	goto L6
L46:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	if v93 == int32(0) {
		v623 = v79
		goto L6
	} else {
		goto L47
	}
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	F_errmsg_internal(m, int32(230323), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(495574), int32(306), int32(208422))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L51:
	;
	v623 = v28
	goto L6
L52:
	;
	goto L53
L53:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	if v109 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if v109 == v110 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v127+v128<<(uint(int32(2))%32)-int32(4))))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	v136 = F_copyObjectImpl(m, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L61
	}
L57:
	;
	v623 = v28
	goto L6
L58:
	;
	goto L59
L59:
	;
	v113 = F_palloc0(m, int32(20))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v113)+4)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = int32(30)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v113)+12)) = int64(-4294967294)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+8)) = v118
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+12)) = v122
	v623 = v113
	goto L6
L61:
	;
	v139 = F_palloc0(m, int32(24))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L18
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+4)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(36)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v139)+12)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+8)) = v144
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	v150 = F_copyObjectImpl(m, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L18
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139)+20)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+16)) = v150
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	if v155 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v28)+24))
	if v156 == int32(0) {
		v623 = v139
		goto L6
	} else {
		goto L65
	}
L65:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	F_errmsg_internal(m, int32(230323), int32(0))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(495574), int32(364), int32(208422))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	v623 = v28
	goto L6
L70:
	;
	goto L71
L71:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+124))
	v176 = int32(0)
	v180 = v176
	v181 = v176
	goto L72
L72:
	;
	v192 = v20 + v180<<(uint(int32(2))%32)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	v195 = F_bms_is_member(m, v194, v175)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L18
	} else {
		goto L74
	}
L73:
	;
	if v199 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L74:
	;
	if v195 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	if v181 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	v199 = v181
	goto L77
L77:
	;
	v201 = v180 + int32(1)
	if v201 != v19 {
		v180 = v201
		v181 = v199
		goto L72
	} else {
		goto L79
	}
L78:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+8))
	v199 = v198
	goto L77
L79:
	;
	goto L73
L80:
	;
	v623 = v28
	goto L6
L81:
	;
	goto L82
L82:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+132))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+12))
	v208 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+8)))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v207+v208<<(uint(int32(2))%32)-int32(4))))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)+16))
	v216 = F_bms_is_member(m, v199, v215)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L18
	} else {
		goto L83
	}
L83:
	;
	if v216 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v219 = F_copyObjectImpl(m, v218)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L18
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
	v229 = F_makeNullConst(m, v226, v227, v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L18
	} else {
		goto L88
	}
L87:
	;
	v221 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v219)+40)) = uint16(v221)
	*(*int32)(unsafe.Add(mBase, uint32(v219)+36)) = v221
	*(*int32)(unsafe.Add(mBase, uint32(v219)+4)) = v199
	v623 = v219
	goto L6
L88:
	;
	v623 = v229
	goto L6
L89:
	;
	if v19 <= int32(0) {
		v623 = v231
		goto L6
	} else {
		goto L90
	}
L90:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231)+4))
	v239 = int32(0)
	goto L91
L91:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v20+v239<<(uint(int32(2))%32))))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+4))
	if v253 != v235 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v252)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v258
	v623 = v231
	goto L6
L93:
	;
	v256 = v239 + int32(1)
	if v19 != v256 {
		v239 = v256
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	goto L92
L96:
	;
	v623 = v231
	goto L6
L97:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v261)+20))
	if v263 != 0 {
		v623 = v261
		goto L6
	} else {
		goto L98
	}
L98:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)+8))
	v265 = int32(0)
	if v19 <= v265 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v303 != 0 {
		goto L116
	} else {
		goto L117
	}
L100:
	;
	v303 = v3
	goto L99
L101:
	;
	goto L102
L102:
	;
	v270 = v265
	v271 = v3
	goto L103
L103:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v20+v270<<(uint(int32(2))%32))))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v285 = F_bms_is_member(m, v284, v264)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L18
	} else {
		goto L105
	}
L104:
	;
	v303 = v296
	goto L99
L105:
	;
	if v285 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	if v271 != 0 {
		goto L109
	} else {
		goto L110
	}
L107:
	;
	v296 = v271
	goto L108
L108:
	;
	v298 = v270 + int32(1)
	if v298 != v19 {
		v270 = v298
		v271 = v296
		goto L103
	} else {
		goto L115
	}
L109:
	;
	v289 = v271
	goto L111
L110:
	;
	v287 = F_bms_copy(m, v264)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L18
	} else {
		goto L112
	}
L111:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v283)+4))
	v291 = F_bms_del_member(m, v289, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L18
	} else {
		goto L113
	}
L112:
	;
	v289 = v287
	goto L111
L113:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v283)+8))
	v294 = F_bms_add_member(m, v291, v293)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L18
	} else {
		goto L114
	}
L114:
	;
	v296 = v294
	goto L108
L115:
	;
	goto L104
L116:
	;
	v312 = v303
	goto L118
L117:
	;
	v312 = v264
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v261)+8)) = v312
	v623 = v261
	goto L6
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = int32(318)
	goto L121
L120:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v323 = F_adjust_appendrel_attrs_mutator(m, v322, l1)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L18
	} else {
		goto L124
	}
L121:
	;
	v320 = F__emscripten_memcpy_bulkmem(m, v315, l0, int32(168))
	mBase = m.M
	goto L123
L123:
	;
	goto L120
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+4)) = v323
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v327 = F_adjust_appendrel_attrs_mutator(m, v326, l1)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L18
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+52)) = v327
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v331 = int32(0)
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v331 < v332 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v342 = v3
	v345 = int32(0)
	goto L129
L127:
	;
	v374 = v3
	goto L128
L128:
	;
	if v374 != 0 {
		goto L142
	} else {
		goto L143
	}
L129:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v335+v345<<(uint(int32(2))%32))))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v354 = F_bms_is_member(m, v353, v330)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L18
	} else {
		goto L131
	}
L130:
	;
	v374 = v365
	goto L128
L131:
	;
	if v354 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	if v342 != 0 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	v365 = v342
	goto L134
L134:
	;
	v367 = v345 + int32(1)
	if v367 != v332 {
		v342 = v365
		v345 = v367
		goto L129
	} else {
		goto L141
	}
L135:
	;
	v358 = v342
	goto L137
L136:
	;
	v356 = F_bms_copy(m, v330)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L18
	} else {
		goto L138
	}
L137:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v352)+4))
	v360 = F_bms_del_member(m, v358, v359)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L18
	} else {
		goto L139
	}
L138:
	;
	v358 = v356
	goto L137
L139:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v352)+8))
	v363 = F_bms_add_member(m, v360, v362)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L18
	} else {
		goto L140
	}
L140:
	;
	v365 = v363
	goto L134
L141:
	;
	goto L130
L142:
	;
	v381 = v374
	goto L144
L143:
	;
	v381 = v330
	goto L144
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+28)) = v381
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v384 <= int32(0) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if v426 != 0 {
		goto L162
	} else {
		goto L163
	}
L146:
	;
	v426 = int32(0)
	goto L145
L147:
	;
	goto L148
L148:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v392 = v331
	v394 = int32(0)
	goto L149
L149:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v388+v392<<(uint(int32(2))%32))))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v407 = F_bms_is_member(m, v406, v383)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L18
	} else {
		goto L151
	}
L150:
	;
	v426 = v418
	goto L145
L151:
	;
	if v407 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v394 != 0 {
		goto L155
	} else {
		goto L156
	}
L153:
	;
	v418 = v394
	goto L154
L154:
	;
	v420 = v392 + int32(1)
	if v420 != v384 {
		v392 = v420
		v394 = v418
		goto L149
	} else {
		goto L161
	}
L155:
	;
	v411 = v394
	goto L157
L156:
	;
	v409 = F_bms_copy(m, v383)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L18
	} else {
		goto L158
	}
L157:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v405)+4))
	v413 = F_bms_del_member(m, v411, v412)
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L18
	} else {
		goto L159
	}
L158:
	;
	v411 = v409
	goto L157
L159:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v405)+8))
	v416 = F_bms_add_member(m, v413, v415)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L18
	} else {
		goto L160
	}
L160:
	;
	v418 = v416
	goto L154
L161:
	;
	goto L150
L162:
	;
	v434 = v426
	goto L164
L163:
	;
	v434 = v383
	goto L164
L164:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+32)) = v434
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v437 = int32(0)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v437 < v439 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v449 = v437
	v452 = int32(0)
	goto L168
L166:
	;
	v481 = v437
	goto L167
L167:
	;
	if v481 != 0 {
		goto L181
	} else {
		goto L182
	}
L168:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v442+v452<<(uint(int32(2))%32))))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	v461 = F_bms_is_member(m, v460, v436)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L18
	} else {
		goto L170
	}
L169:
	;
	v481 = v472
	goto L167
L170:
	;
	if v461 != 0 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	if v449 != 0 {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	v472 = v449
	goto L173
L173:
	;
	v474 = v452 + int32(1)
	if v474 != v439 {
		v449 = v472
		v452 = v474
		goto L168
	} else {
		goto L180
	}
L174:
	;
	v465 = v449
	goto L176
L175:
	;
	v463 = F_bms_copy(m, v436)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L18
	} else {
		goto L177
	}
L176:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	v467 = F_bms_del_member(m, v465, v466)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L18
	} else {
		goto L178
	}
L177:
	;
	v465 = v463
	goto L176
L178:
	;
	v469 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	v470 = F_bms_add_member(m, v467, v469)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L18
	} else {
		goto L179
	}
L179:
	;
	v472 = v470
	goto L173
L180:
	;
	goto L169
L181:
	;
	v488 = v481
	goto L183
L182:
	;
	v488 = v436
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+40)) = v488
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v491 <= int32(0) {
		goto L185
	} else {
		goto L186
	}
L184:
	;
	if v533 != 0 {
		goto L201
	} else {
		goto L202
	}
L185:
	;
	v533 = int32(0)
	goto L184
L186:
	;
	goto L187
L187:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v499 = v437
	v501 = int32(0)
	goto L188
L188:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v495+v499<<(uint(int32(2))%32))))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	v514 = F_bms_is_member(m, v513, v490)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L18
	} else {
		goto L190
	}
L189:
	;
	v533 = v525
	goto L184
L190:
	;
	if v514 != 0 {
		goto L191
	} else {
		goto L192
	}
L191:
	;
	if v501 != 0 {
		goto L194
	} else {
		goto L195
	}
L192:
	;
	v525 = v501
	goto L193
L193:
	;
	v527 = v499 + int32(1)
	if v527 != v491 {
		v499 = v527
		v501 = v525
		goto L188
	} else {
		goto L200
	}
L194:
	;
	v518 = v501
	goto L196
L195:
	;
	v516 = F_bms_copy(m, v490)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L18
	} else {
		goto L197
	}
L196:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v512)+4))
	v520 = F_bms_del_member(m, v518, v519)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L18
	} else {
		goto L198
	}
L197:
	;
	v518 = v516
	goto L196
L198:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v512)+8))
	v523 = F_bms_add_member(m, v520, v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L18
	} else {
		goto L199
	}
L199:
	;
	v525 = v523
	goto L193
L200:
	;
	goto L189
L201:
	;
	v541 = v533
	goto L203
L202:
	;
	v541 = v490
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+44)) = v541
	v543 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v544 = int32(0)
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v544 < v545 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v552 = int32(0)
	v554 = v544
	goto L207
L205:
	;
	v586 = v544
	goto L206
L206:
	;
	v594 = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v320)+152)) = v594
	*(*int64)(unsafe.Add(mBase, uint32(v320)+144)) = v594
	*(*int64)(unsafe.Add(mBase, uint32(v320)+136)) = v594
	*(*int64)(unsafe.Add(mBase, uint32(v320)+128)) = v594
	*(*int32)(unsafe.Add(mBase, uint32(v320)+116)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v320)+108)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v320)+88)) = v594
	*(*int64)(unsafe.Add(mBase, uint32(v320)+80)) = v594
	*(*int64)(unsafe.Add(mBase, uint32(v320)+64)) = v594
	if v586 != 0 {
		goto L220
	} else {
		goto L221
	}
L207:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v548+v552<<(uint(int32(2))%32))))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	v567 = F_bms_is_member(m, v566, v543)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L18
	} else {
		goto L209
	}
L208:
	;
	v586 = v578
	goto L206
L209:
	;
	if v567 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	if v554 != 0 {
		goto L213
	} else {
		goto L214
	}
L211:
	;
	v578 = v554
	goto L212
L212:
	;
	v580 = v552 + int32(1)
	if v580 != v545 {
		v552 = v580
		v554 = v578
		goto L207
	} else {
		goto L219
	}
L213:
	;
	v571 = v554
	goto L215
L214:
	;
	v569 = F_bms_copy(m, v543)
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L18
	} else {
		goto L216
	}
L215:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v565)+4))
	v573 = F_bms_del_member(m, v571, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L18
	} else {
		goto L217
	}
L216:
	;
	v571 = v569
	goto L215
L217:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v565)+8))
	v576 = F_bms_add_member(m, v573, v575)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L18
	} else {
		goto L218
	}
L218:
	;
	v578 = v576
	goto L212
L219:
	;
	goto L208
L220:
	;
	v612 = v586
	goto L222
L221:
	;
	v612 = v543
	goto L222
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v320)+48)) = v612
	v623 = v315
	goto L6
L223:
	;
	v623 = v615
	goto L6
L224:
	;
	v637 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+8)))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v639 = F_get_rel_name(m, v638)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L18
	} else {
		goto L225
	}
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v639
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v637
	F_errmsg_internal(m, int32(71412), v15)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L18
	} else {
		goto L226
	}
L226:
	;
	F_errfinish(m, int32(495574), int32(287), int32(208422))
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L18
	} else {
		goto L227
	}
L227:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L228:
	;
	v655 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+8)))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v657 = F_get_rel_name(m, v656)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L18
	} else {
		goto L229
	}
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v657
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v655
	F_errmsg_internal(m, int32(71412), v15+int32(16))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L18
	} else {
		goto L230
	}
L230:
	;
	F_errfinish(m, int32(495574), int32(292), int32(208422))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L18
	} else {
		goto L231
	}
L231:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L232:
	;
	F_errmsg_internal(m, int32(230364), int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L18
	} else {
		goto L233
	}
L233:
	;
	F_errfinish(m, int32(495574), int32(304), int32(208422))
	mBase = m.M
	v683 = m.ExcPending
	if v683 != 0 {
		goto L18
	} else {
		goto L234
	}
L234:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L235:
	;
	F_errmsg_internal(m, int32(230364), int32(0))
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L18
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(495574), int32(362), int32(208422))
	mBase = m.M
	v696 = m.ExcPending
	if v696 != 0 {
		goto L18
	} else {
		goto L237
	}
L237:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L238:
	;
	F_errmsg_internal(m, int32(173200), int32(0))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L18
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(495574), int32(390), int32(208422))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L18
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_adjust_inherited_attnums_multilevel(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13+l2<<(uint(int32(2))%32))))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if l3 != v18 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L7
	} else {
		goto L34
	}
L4:
	;
	v20 = F_adjust_inherited_attnums_multilevel(m, l0, l1, v18, l3)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v24 = l1
	goto L6
L6:
	;
	v25 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(32)
	m.G0 = v29
	if v24 != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	return int32(0)
L8:
	;
	v24 = v20
	goto L6
L9:
	;
	m.G0 = v29 + int32(32)
	m.G0 = v11 + int32(16)
	return v112
L10:
	;
	v34 = v25
	v36 = v25
	goto L16
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if int32(0) < v31 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v112 = v25
	goto L9
L14:
	;
	goto L13
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L7
	} else {
		goto L30
	}
L16:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42+v34<<(uint(int32(2))%32)))))
	if v46 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L26
	}
L18:
	;
	goto L17
L19:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v49 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v52 < v46 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+v46<<(uint(int32(2))%32)-int32(4))))
	if v60 == int32(0) {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	if v63 != int32(6) {
		goto L15
	} else {
		goto L23
	}
L23:
	;
	v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v60)+8)))
	v67 = F_lappend_int(m, v36, v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L24
	}
L24:
	;
	v70 = v34 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v70 < v71 {
		v34 = v70
		v36 = v67
		goto L16
	} else {
		goto L25
	}
L25:
	;
	v112 = v67
	goto L9
L26:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v79 = F_get_rel_name(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v46
	F_errmsg_internal(m, int32(71412), v29)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(495574), int32(670), int32(150459))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v96 = F_get_rel_name(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v46
	F_errmsg_internal(m, int32(71412), v29+int32(16))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(495574), int32(674), int32(150459))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
	F_errmsg_internal(m, int32(24429), v11)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(495574), int32(692), int32(305451))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_adjust_paths_for_srfs(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v13 == int32(1) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v108 == int32(0) {
		goto L1
	} else {
		goto L37
	}
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v31 = int32(0)
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v38 = v35 + v31<<(uint(int32(2))%32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v44 = int32(0)
	v45 = v39
	goto L11
L11:
	;
	v52 = int32(0)
	if l2 == v52 {
		v62 = v52
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l3 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v56 <= v44 {
		v62 = int32(0)
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v62 = v58 + v44<<(uint(int32(2))%32)
	goto L13
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	if v87 != 0 {
		goto L31
	} else {
		goto L32
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v73
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if v76 == v39 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v73 = v39
	goto L17
L19:
	;
	goto L20
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v65 <= v44 {
		v73 = v45
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v62 == int32(0) {
		v73 = v45
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v72 = v69 + v44<<(uint(int32(2))%32)
	if v72 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v73 = v45
	goto L17
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v73
	goto L26
L25:
	;
	goto L26
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v79 == v39 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v73
	goto L29
L28:
	;
	goto L29
L29:
	;
	v83 = v31 + int32(1)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v83 < v84 {
		v31 = v83
		goto L9
	} else {
		goto L30
	}
L30:
	;
	goto L6
L31:
	;
	v88 = F_create_set_projection_path(m, l0, l1, v45, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v92 = F_apply_projection_to_path(m, l0, l1, v45, v86)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L34
	} else {
		goto L36
	}
L34:
	;
	return
L35:
	;
	v44 = v44 + int32(1)
	v45 = v88
	goto L11
L36:
	;
	v44 = v44 + int32(1)
	v45 = v92
	goto L11
L37:
	;
	v111 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v112 <= v111 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v123 = v111
	goto L39
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v131 = v128 + v123<<(uint(int32(2))%32)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v137 = int32(0)
	v138 = v132
	goto L41
L41:
	;
	v145 = int32(0)
	if l2 == v145 {
		v155 = v145
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if l3 != 0 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v149 <= v137 {
		v155 = int32(0)
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v155 = v151 + v137<<(uint(int32(2))%32)
	goto L43
L46:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	if v173 != 0 {
		goto L55
	} else {
		goto L56
	}
L47:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v156 <= v137 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v166 = v132
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v166
	v169 = v123 + int32(1)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v169 < v170 {
		v123 = v169
		goto L39
	} else {
		goto L54
	}
L50:
	;
	v166 = v138
	goto L49
L51:
	;
	if v155 == int32(0) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v163 = v160 + v137<<(uint(int32(2))%32)
	if v163 != 0 {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	goto L50
L54:
	;
	goto L1
L55:
	;
	v174 = F_create_set_projection_path(m, l0, l1, v138, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L34
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v178 = F_create_projection_path(m, l0, l1, v138, v172)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L34
	} else {
		goto L59
	}
L58:
	;
	v137 = v137 + int32(1)
	v138 = v174
	goto L41
L59:
	;
	v137 = v137 + int32(1)
	v138 = v178
	goto L41
}
func F_advance_transition_function(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
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
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)))
	if v9 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v100
	goto L1
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v12 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v64 = int32(4489440)
	v65 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v68
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = l1
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)) = uint8(v74)
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)) = uint8(v73)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	v79 = m.T0[v78].(func(*base.Module, int32) int32)(m, v8)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L16
	} else {
		goto L19
	}
L6:
	;
	v21 = int32(1)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v39 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+int32(24)+v21<<(uint(int32(3))%32)))))
	if v28 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v30 = v21 + int32(1)
	if v30 <= v12 {
		v21 = v30
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v42 = int32(4489440)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
	v50 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+184)))
	v51 = F_datumCopy(m, v48, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v56 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	v53 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v51
	v100 = v43
	goto L2
L18:
	;
	goto L5
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
	if v83 != 0 {
		v90 = v79
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v90
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v93)
	v100 = v65
	goto L2
L21:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v79 == v84 {
		v90 = v79
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v88 = F_ExecAggCopyTransValue(m, l0, l1, v79, v86, v84, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v90 = v88
	goto L20
}
func F_alen_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
	if v5 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(229561), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(494070), int32(1922), int32(229317))
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
		return int32(0)
	}
}
func F_anycompatible_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(390136)
			F_errmsg(m, int32(192226), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493543), int32(376), int32(67280))
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
}
func F_anyelement_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(95902)
			F_errmsg(m, int32(192226), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493543), int32(374), int32(66697))
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
}
func F_anyenum_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(286842)
			F_errmsg(m, int32(192260), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(493543), int32(194), int32(279188))
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
}
func F_arrayexpr_startup_fn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	v7 = F_palloc(m, int32(40))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v7
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(17)
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v15 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v15
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+16)) = uint8(v15)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v14
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v22
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v25 = F_list_copy(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v25
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v31
			if v31 != 0 {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
				v34 = v33
			} else {
				v34 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v34
			return
		}
	}
}
func F_ascii_safe_strlcpy(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = l2 - int32(1)
	if v7 == int32(0) {
		v44 = l0
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v49 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v44))) = uint8(v49)
	goto L3
L5:
	;
	v10 = l0
	v11 = l1
	v13 = v7
	goto L6
L6:
	;
	v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11))))
	if v15 == int32(0) {
		v44 = v10
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v44 = v41
	goto L4
L8:
	;
	if int32(31) < v15 {
		v35 = v15
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10))) = uint8(v35)
	v41 = v10 + v37
	v43 = v13 - v37
	if v43 != 0 {
		v10 = v41
		v11 = v11 + v37
		v13 = v43
		goto L6
	} else {
		goto L12
	}
L10:
	;
	v21 = v15 - int32(9)
	if base.Ui32(int32(4)) < base.Ui32(v21&int32(255)) {
		v35 = int32(63)
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v35 = base.I32_wrap_i64(int64(base.Ui64(int64(56895670793)) >> (uint(base.I64_extend_i32_u(v21<<(uint(int32(3))%32))&int64(248)) % 64)))
	goto L9
L12:
	;
	goto L7
}
func F_assign_datestyle(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[522])) = v4
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, _consts[523])) = v7
	return
}
func F_assign_restrict_nonsystem_relation_kind(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _consts[894])) = v4
	return
}
func F_assign_syslog_facility(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, _consts[1109]))
	if l0 != v4 {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1110])))
		if v7 != 0 {
			v9 = m.G0
			v10 = int32(16)
			v11 = v9 - v10
			m.G0 = v11
			v13 = int32(4383808)
			v14 = *(*int32)(unsafe.Add(mBase, _consts[1111]))
			v15 = F_close(m, v14)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _consts[1111])) = int32(-1)
			m.G0 = v11 + v10
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _consts[1110])) = uint8(v24)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, _consts[1109])) = l0
	} else {
	}
	return
}
func F_atexit_callback(m *base.Module) {
	var v3 int32
	_ = v3
	F_proc_exit_prepare(m, int32(-1))
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_attnameAttNum(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v7 = int32(*(*int16)(unsafe.Add(mBase, uint32(v6)+120)))
	if v4 < v7 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v98
L2:
	;
	v98 = v13 + int32(1)
	goto L1
L3:
	;
	v13 = v4
	goto L6
L4:
	;
	goto L5
L5:
	;
	if l2 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v22 = v15 + v16<<(uint(int32(4))%32) + v13*int32(100)
	v24 = v22 + int32(24)
	if v24|l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	if v38 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	v30 = int32(-1)
	goto L11
L10:
	;
	v30 = int32(0)
	goto L11
L11:
	;
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v31 = int32(1)
	goto L14
L13:
	;
	v31 = v30
	goto L14
L14:
	;
	if v24 == int32(0) {
		v38 = v31
		goto L15
	} else {
		goto L16
	}
L15:
	;
	goto L8
L16:
	;
	if l1 == int32(0) {
		v38 = v31
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v37 = F_strncmp(m, v24, l1, int32(64))
	mBase = m.M
	v38 = v37
	goto L15
L18:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+111)))
	if v41 != int32(1) {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v45 = v13 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+120)))
	if v45 < v47 {
		v13 = v45
		goto L6
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	goto L7
L23:
	;
	return int32(0)
L24:
	;
	v57 = F_strcmp(m, int32(749340), l1)
	mBase = m.M
	if v57 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v86 == int32(0) {
		goto L23
	} else {
		goto L44
	}
L26:
	;
	v86 = int32(749336)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v62 = F_strcmp(m, int32(749440), l1)
	mBase = m.M
	if v62 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = int32(749436)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v67 = F_strcmp(m, int32(749540), l1)
	mBase = m.M
	if v67 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v86 = int32(749536)
	goto L25
L33:
	;
	goto L34
L34:
	;
	v72 = F_strcmp(m, int32(749640), l1)
	mBase = m.M
	if v72 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v86 = int32(749636)
	goto L25
L36:
	;
	goto L37
L37:
	;
	v77 = F_strcmp(m, int32(749740), l1)
	mBase = m.M
	if v77 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v86 = int32(749736)
	goto L25
L39:
	;
	goto L40
L40:
	;
	v84 = F_strcmp(m, int32(749840), l1)
	mBase = m.M
	if v84 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v85 = int32(0)
	goto L43
L42:
	;
	v85 = int32(749836)
	goto L43
L43:
	;
	v86 = v85
	goto L25
L44:
	;
	v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v86)+74)))
	if v89 != 0 {
		v98 = v89
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L23
}
