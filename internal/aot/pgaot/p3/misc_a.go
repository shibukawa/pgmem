package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_AccessTempTableNamespace(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v6 = m.G0
	v8 = v6 - int32(128)
	m.G0 = v8
	v10 = int32(_a_F_AccessTempTableNamespace_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[0])) = v12 | int32(1)
	if l0 == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1]))
		if v19 != 0 {
			m.G0 = v8 + int32(128)
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[2]))
			v24 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[3]))
			v26 = F_object_aclcheck(m, int32(1262), v22, v24, int64(1024))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				if v26 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return
						} else {
							v140 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[2]))
							v141 = F_get_database_name(m, v140)
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v141
								F_errmsg(m, int32(_a_F_AccessTempTableNamespace_1), v8+int32(32))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_3), int32(_a_F_AccessTempTableNamespace_4))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
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
					v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[4])))
					if v30 == int32(1) {
						v35 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[5]))
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+308))
						v38 = base.B2i32(v36 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[4])) = uint8(v38)
						v40 = v38
					} else {
						v40 = int32(0)
					}
					if v40 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return
						} else {
							F_errcode(m, int32(100663618))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_AccessTempTableNamespace_5), int32(0))
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_6), int32(_a_F_AccessTempTableNamespace_4))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
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
						v42 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[6]))
						if int32(0) <= v42 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								F_errcode(m, int32(100663618))
								mBase = m.M
								v176 = m.ExcPending
								if v176 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_AccessTempTableNamespace_7), int32(0))
									mBase = m.M
									v180 = m.ExcPending
									if v180 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_8), int32(_a_F_AccessTempTableNamespace_4))
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
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v46
							v49 = v8 + int32(48)
							v54 = F_pg_snprintf(m, v49, int32(64), int32(_a_F_AccessTempTableNamespace_9), v8+int32(16))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return
							} else {
								v57 = base.I64_extend_i32_u(v49)
								v58 = int64(0)
								v61 = F_GetSysCacheOid(m, int32(37), v57, v58, v58, v58)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									if v61 == int32(0) {
										v67 = F_NamespaceCreate(m, v49, int32(10), int32(1))
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
											} else {
												v82 = v67
												v84 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = v84
												v87 = v8 + int32(48)
												v90 = F_pg_snprintf(m, v87, int32(64), int32(_a_F_AccessTempTableNamespace_10), v8)
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return
												} else {
													v93 = int64(0)
													v96 = F_GetSysCacheOid(m, int32(37), v57, v93, v93, v93)
													mBase = m.M
													v97 = m.ExcPending
													if v97 != 0 {
														return
													} else {
														if v96 == int32(0) {
															v102 = F_NamespaceCreate(m, v87, int32(10), int32(1))
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																F_CommandCounterIncrement(m)
																mBase = m.M
																v105 = m.ExcPending
																if v105 != 0 {
																	return
																} else {
																	v106 = v102
																	*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v106
																	*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v82
																	v112 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
																	*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v82
																	v115 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
																	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
																	v118 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v118)
																	*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v116
																	v123 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v123)
																	m.G0 = v8 + int32(128)
																	return
																}
															}
														} else {
															v106 = v96
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v106
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v82
															v112 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
															*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v82
															v115 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
															v118 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v118)
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v116
															v123 = int32(0)
															*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v123)
															m.G0 = v8 + int32(128)
															return
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+124)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = v61
										*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = int32(2615)
										F_performDeletion(m, v8+int32(116), int32(1), int32(29))
										mBase = m.M
										v81 = m.ExcPending
										if v81 != 0 {
											return
										} else {
											v82 = v61
											v84 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v84
											v87 = v8 + int32(48)
											v90 = F_pg_snprintf(m, v87, int32(64), int32(_a_F_AccessTempTableNamespace_10), v8)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												v93 = int64(0)
												v96 = F_GetSysCacheOid(m, int32(37), v57, v93, v93, v93)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													if v96 == int32(0) {
														v102 = F_NamespaceCreate(m, v87, int32(10), int32(1))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																v106 = v102
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v106
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v82
																v112 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
																*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v82
																v115 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
																v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
																v118 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v118)
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v116
																v123 = int32(0)
																*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v123)
																m.G0 = v8 + int32(128)
																return
															}
														}
													} else {
														v106 = v96
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v106
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v82
														v112 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
														*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v82
														v115 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
														v118 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v118)
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v116
														v123 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v123)
														m.G0 = v8 + int32(128)
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
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[2]))
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[3]))
		v26 = F_object_aclcheck(m, int32(1262), v22, v24, int64(1024))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			if v26 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v135 = m.ExcPending
				if v135 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						v140 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[2]))
						v141 = F_get_database_name(m, v140)
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v141
							F_errmsg(m, int32(_a_F_AccessTempTableNamespace_1), v8+int32(32))
							mBase = m.M
							v148 = m.ExcPending
							if v148 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_3), int32(_a_F_AccessTempTableNamespace_4))
								mBase = m.M
								v153 = m.ExcPending
								if v153 != 0 {
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
				v30 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[4])))
				if v30 == int32(1) {
					v35 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[5]))
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+308))
					v38 = base.B2i32(v36 != int32(2))
					*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[4])) = uint8(v38)
					v40 = v38
				} else {
					v40 = int32(0)
				}
				if v40 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return
					} else {
						F_errcode(m, int32(100663618))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_AccessTempTableNamespace_5), int32(0))
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_6), int32(_a_F_AccessTempTableNamespace_4))
								mBase = m.M
								v169 = m.ExcPending
								if v169 != 0 {
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
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[6]))
					if int32(0) <= v42 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v173 = m.ExcPending
						if v173 != 0 {
							return
						} else {
							F_errcode(m, int32(100663618))
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_AccessTempTableNamespace_7), int32(0))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_8), int32(_a_F_AccessTempTableNamespace_4))
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
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v46
						v49 = v8 + int32(48)
						v54 = F_pg_snprintf(m, v49, int32(64), int32(_a_F_AccessTempTableNamespace_9), v8+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							v57 = base.I64_extend_i32_u(v49)
							v58 = int64(0)
							v61 = F_GetSysCacheOid(m, int32(37), v57, v58, v58, v58)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								if v61 == int32(0) {
									v67 = F_NamespaceCreate(m, v49, int32(10), int32(1))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v70 = m.ExcPending
										if v70 != 0 {
											return
										} else {
											v82 = v67
											v84 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v84
											v87 = v8 + int32(48)
											v90 = F_pg_snprintf(m, v87, int32(64), int32(_a_F_AccessTempTableNamespace_10), v8)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return
											} else {
												v93 = int64(0)
												v96 = F_GetSysCacheOid(m, int32(37), v57, v93, v93, v93)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return
												} else {
													if v96 == int32(0) {
														v102 = F_NamespaceCreate(m, v87, int32(10), int32(1))
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v105 = m.ExcPending
															if v105 != 0 {
																return
															} else {
																v106 = v102
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v106
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v82
																v112 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
																*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v82
																v115 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
																v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
																v118 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v118)
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v116
																v123 = int32(0)
																*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v123)
																m.G0 = v8 + int32(128)
																return
															}
														}
													} else {
														v106 = v96
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v106
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v82
														v112 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
														*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v82
														v115 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
														v118 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v118)
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v116
														v123 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v123)
														m.G0 = v8 + int32(128)
														return
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+124)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v8)+120)) = v61
									*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = int32(2615)
									F_performDeletion(m, v8+int32(116), int32(1), int32(29))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										v82 = v61
										v84 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = v84
										v87 = v8 + int32(48)
										v90 = F_pg_snprintf(m, v87, int32(64), int32(_a_F_AccessTempTableNamespace_10), v8)
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return
										} else {
											v93 = int64(0)
											v96 = F_GetSysCacheOid(m, int32(37), v57, v93, v93, v93)
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return
											} else {
												if v96 == int32(0) {
													v102 = F_NamespaceCreate(m, v87, int32(10), int32(1))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return
													} else {
														F_CommandCounterIncrement(m)
														mBase = m.M
														v105 = m.ExcPending
														if v105 != 0 {
															return
														} else {
															v106 = v102
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v106
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v82
															v112 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
															*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v82
															v115 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
															v118 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v118)
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v116
															v123 = int32(0)
															*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v123)
															m.G0 = v8 + int32(128)
															return
														}
													}
												} else {
													v106 = v96
													*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v106
													*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v82
													v112 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
													*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v82
													v115 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
													v118 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v118)
													*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v116
													v123 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v123)
													m.G0 = v8 + int32(128)
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
func F_AioShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[0]))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[1]))
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[2]))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[3]))
	v20 = v15 * (v17 + int32(38))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v11 * v20
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v25
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v28
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v31
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[7]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v34
	if v17 != int32(-38) {
		v42 = int32(0)
		v43 = v2
		v45 = v2
		for {
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[1]))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
			v53 = v50 + v43*int32(164)
			*(*int32)(unsafe.Add(mBase, uint32(v53))) = v45
			v55 = int32(_a_F_AioShmemInit_0)
			v56 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[2]))
			v57 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v57
			v60 = v53 + int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v60
			*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v60
			base.MemoryFill(m, v53+int32(24), v57, int32(128))
			*(*int32)(unsafe.Add(mBase, uint32(v53)+160)) = v57
			v71 = v53 + int32(152)
			*(*int32)(unsafe.Add(mBase, uint32(v53)+156)) = v71
			*(*int32)(unsafe.Add(mBase, uint32(v53)+152)) = v71
			v76 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[2]))
			if v57 < v76 {
				v82 = v42
				v84 = v57
				for {
					v89 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[1]))
					v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
					v91 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
					v92 = int32(7)
					v97 = v90 + v91<<(uint(v92)%32) + v84<<(uint(v92)%32)
					*(*int32)(unsafe.Add(mBase, uint32(v97)+76)) = v82
					*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v43
					*(*int64)(unsafe.Add(mBase, uint32(v97)+48)) = int64(1)
					v102 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v97)+80)) = v102
					*(*uint8)(unsafe.Add(mBase, uint32(v97)+13)) = uint8(v102)
					*(*int32)(unsafe.Add(mBase, uint32(v97)+32)) = v102
					*(*uint16)(unsafe.Add(mBase, uint32(v97)+3)) = uint16(v102)
					v110 = *(*int32)(unsafe.Add(mBase, uint32(v97)+68))
					*(*int32)(unsafe.Add(mBase, uint32(v97)+68)) = v110 & int32(-449)
					v115 = v97 + int32(56)
					atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v115))), uint32(v102))
					*(*int64)(unsafe.Add(mBase, uint32(v115)+4)) = int64(-1)
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
					if v121 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v60
						*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v60
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v97)+28)) = v60
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v97)+24)) = v129
					v132 = v97 + int32(24)
					*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v132
					*(*int32)(unsafe.Add(mBase, uint32(v53)+4)) = v132
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
					v136 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v53)+12)) = v135 + v136
					v140 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[0]))
					v141 = v140 + v82
					v143 = v84 + v136
					v145 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[2]))
					if v143 < v145 {
						v82 = v141
						v84 = v143
						continue
					} else {
						break
					}
					break
				}
				v150 = v141
			} else {
				v150 = v42
			}
			v158 = v43 + int32(1)
			v160 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[3]))
			if base.Ui32(v158) < base.Ui32(v160+int32(38)) {
				v42 = v150
				v43 = v158
				v45 = v45 + v56
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_AioShmemInit[8]))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+12))
	if v175 != 0 {
		v176 = *(*int32)(unsafe.Add(mBase, uint32(v174)+20))
		m.T0[v175].(func(*base.Module, int32))(m, v176)
		mBase = m.M
		v178 = m.ExcPending
		if v178 != 0 {
			return
		} else {
			return
		}
	} else {
		return
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	v13 = v11 + v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+68))
	if v14 == l2 {
		m.G0 = v9 + int32(16)
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
				v61 = m.ExcPending
				if v61 != 0 {
					return
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
						if v67 == int32(0) {
							m.G0 = v9 + int32(16)
							return
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v72 = int32(0)
							F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[1]))
				v22 = F_object_ownercheck(m, int32(_a_F_AlterPublicationOwner_internal_0), v19, v21)
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
							v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[1]))
							F_check_can_set_role(m, v33, l2)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[2]))
								v40 = F_object_aclcheck(m, int32(1262), v38, l2, int64(512))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									if v40 != 0 {
										v44 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[2]))
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
												v49 = F_superuser_arg(m, l2)
												mBase = m.M
												v50 = m.ExcPending
												if v50 != 0 {
													return
												} else {
													if v49 != 0 {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
														F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
														mBase = m.M
														v61 = m.ExcPending
														if v61 != 0 {
															return
														} else {
															v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
															mBase = m.M
															v65 = m.ExcPending
															if v65 != 0 {
																return
															} else {
																v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																if v67 == int32(0) {
																	m.G0 = v9 + int32(16)
																	return
																} else {
																	v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v72 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
																	mBase = m.M
																	v76 = m.ExcPending
																	if v76 != 0 {
																		return
																	} else {
																		m.G0 = v9 + int32(16)
																		return
																	}
																}
															}
														}
													} else {
														v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)))
														if v51 != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return
															} else {
																F_errcode(m, int32(16797828))
																mBase = m.M
																v87 = m.ExcPending
																if v87 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																	F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																		mBase = m.M
																		v97 = m.ExcPending
																		if v97 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																			mBase = m.M
																			v102 = m.ExcPending
																			if v102 != 0 {
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
															v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+73)))
															if v52 != 0 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return
																} else {
																	F_errcode(m, int32(16797828))
																	mBase = m.M
																	v87 = m.ExcPending
																	if v87 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																		F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return
																		} else {
																			F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																			mBase = m.M
																			v97 = m.ExcPending
																			if v97 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																				mBase = m.M
																				v102 = m.ExcPending
																				if v102 != 0 {
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
																v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																v54 = F_is_schema_publication(m, v53)
																mBase = m.M
																v55 = m.ExcPending
																if v55 != 0 {
																	return
																} else {
																	if v54 != 0 {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v84 = m.ExcPending
																		if v84 != 0 {
																			return
																		} else {
																			F_errcode(m, int32(16797828))
																			mBase = m.M
																			v87 = m.ExcPending
																			if v87 != 0 {
																				return
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																				F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																				mBase = m.M
																				v93 = m.ExcPending
																				if v93 != 0 {
																					return
																				} else {
																					F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																					mBase = m.M
																					v97 = m.ExcPending
																					if v97 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																						mBase = m.M
																						v102 = m.ExcPending
																						if v102 != 0 {
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
																		v61 = m.ExcPending
																		if v61 != 0 {
																			return
																		} else {
																			v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
																			mBase = m.M
																			v65 = m.ExcPending
																			if v65 != 0 {
																				return
																			} else {
																				v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																				if v67 == int32(0) {
																					m.G0 = v9 + int32(16)
																					return
																				} else {
																					v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																					v72 = int32(0)
																					F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
																					mBase = m.M
																					v76 = m.ExcPending
																					if v76 != 0 {
																						return
																					} else {
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
										v49 = F_superuser_arg(m, l2)
										mBase = m.M
										v50 = m.ExcPending
										if v50 != 0 {
											return
										} else {
											if v49 != 0 {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
												F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
												mBase = m.M
												v61 = m.ExcPending
												if v61 != 0 {
													return
												} else {
													v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return
													} else {
														v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
														if v67 == int32(0) {
															m.G0 = v9 + int32(16)
															return
														} else {
															v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															v72 = int32(0)
															F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
															mBase = m.M
															v76 = m.ExcPending
															if v76 != 0 {
																return
															} else {
																m.G0 = v9 + int32(16)
																return
															}
														}
													}
												}
											} else {
												v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)))
												if v51 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														F_errcode(m, int32(16797828))
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
															F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
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
													v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+73)))
													if v52 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															F_errcode(m, int32(16797828))
															mBase = m.M
															v87 = m.ExcPending
															if v87 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return
																} else {
																	F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																		mBase = m.M
																		v102 = m.ExcPending
																		if v102 != 0 {
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
														v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v54 = F_is_schema_publication(m, v53)
														mBase = m.M
														v55 = m.ExcPending
														if v55 != 0 {
															return
														} else {
															if v54 != 0 {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v84 = m.ExcPending
																if v84 != 0 {
																	return
																} else {
																	F_errcode(m, int32(16797828))
																	mBase = m.M
																	v87 = m.ExcPending
																	if v87 != 0 {
																		return
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																		F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																		mBase = m.M
																		v93 = m.ExcPending
																		if v93 != 0 {
																			return
																		} else {
																			F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																			mBase = m.M
																			v97 = m.ExcPending
																			if v97 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																				mBase = m.M
																				v102 = m.ExcPending
																				if v102 != 0 {
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
																v61 = m.ExcPending
																if v61 != 0 {
																	return
																} else {
																	v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
																	mBase = m.M
																	v65 = m.ExcPending
																	if v65 != 0 {
																		return
																	} else {
																		v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																		if v67 == int32(0) {
																			m.G0 = v9 + int32(16)
																			return
																		} else {
																			v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v72 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
																			mBase = m.M
																			v76 = m.ExcPending
																			if v76 != 0 {
																				return
																			} else {
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
							}
						}
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[1]))
						F_check_can_set_role(m, v33, l2)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return
						} else {
							v38 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[2]))
							v40 = F_object_aclcheck(m, int32(1262), v38, l2, int64(512))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return
							} else {
								if v40 != 0 {
									v44 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[2]))
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
											v49 = F_superuser_arg(m, l2)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												if v49 != 0 {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
													F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
													mBase = m.M
													v61 = m.ExcPending
													if v61 != 0 {
														return
													} else {
														v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
														mBase = m.M
														v65 = m.ExcPending
														if v65 != 0 {
															return
														} else {
															v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
															if v67 == int32(0) {
																m.G0 = v9 + int32(16)
																return
															} else {
																v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																v72 = int32(0)
																F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
																mBase = m.M
																v76 = m.ExcPending
																if v76 != 0 {
																	return
																} else {
																	m.G0 = v9 + int32(16)
																	return
																}
															}
														}
													}
												} else {
													v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)))
													if v51 != 0 {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v84 = m.ExcPending
														if v84 != 0 {
															return
														} else {
															F_errcode(m, int32(16797828))
															mBase = m.M
															v87 = m.ExcPending
															if v87 != 0 {
																return
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																mBase = m.M
																v93 = m.ExcPending
																if v93 != 0 {
																	return
																} else {
																	F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																	mBase = m.M
																	v97 = m.ExcPending
																	if v97 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																		mBase = m.M
																		v102 = m.ExcPending
																		if v102 != 0 {
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
														v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+73)))
														if v52 != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return
															} else {
																F_errcode(m, int32(16797828))
																mBase = m.M
																v87 = m.ExcPending
																if v87 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																	F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																		mBase = m.M
																		v97 = m.ExcPending
																		if v97 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																			mBase = m.M
																			v102 = m.ExcPending
																			if v102 != 0 {
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
															v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															v54 = F_is_schema_publication(m, v53)
															mBase = m.M
															v55 = m.ExcPending
															if v55 != 0 {
																return
															} else {
																if v54 != 0 {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v84 = m.ExcPending
																	if v84 != 0 {
																		return
																	} else {
																		F_errcode(m, int32(16797828))
																		mBase = m.M
																		v87 = m.ExcPending
																		if v87 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																			F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																			mBase = m.M
																			v93 = m.ExcPending
																			if v93 != 0 {
																				return
																			} else {
																				F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																				mBase = m.M
																				v97 = m.ExcPending
																				if v97 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																					mBase = m.M
																					v102 = m.ExcPending
																					if v102 != 0 {
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
																	v61 = m.ExcPending
																	if v61 != 0 {
																		return
																	} else {
																		v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
																		mBase = m.M
																		v65 = m.ExcPending
																		if v65 != 0 {
																			return
																		} else {
																			v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																			if v67 == int32(0) {
																				m.G0 = v9 + int32(16)
																				return
																			} else {
																				v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				v72 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
																				mBase = m.M
																				v76 = m.ExcPending
																				if v76 != 0 {
																					return
																				} else {
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
									v49 = F_superuser_arg(m, l2)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										if v49 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = l2
											F_CatalogTupleUpdate(m, l0, l1+int32(4), l1)
											mBase = m.M
											v61 = m.ExcPending
											if v61 != 0 {
												return
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
												F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return
												} else {
													v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
													if v67 == int32(0) {
														m.G0 = v9 + int32(16)
														return
													} else {
														v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
														v72 = int32(0)
														F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return
														} else {
															m.G0 = v9 + int32(16)
															return
														}
													}
												}
											}
										} else {
											v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+72)))
											if v51 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v84 = m.ExcPending
												if v84 != 0 {
													return
												} else {
													F_errcode(m, int32(16797828))
													mBase = m.M
													v87 = m.ExcPending
													if v87 != 0 {
														return
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
														F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
														mBase = m.M
														v93 = m.ExcPending
														if v93 != 0 {
															return
														} else {
															F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
															mBase = m.M
															v97 = m.ExcPending
															if v97 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																mBase = m.M
																v102 = m.ExcPending
																if v102 != 0 {
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
												v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+73)))
												if v52 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v84 = m.ExcPending
													if v84 != 0 {
														return
													} else {
														F_errcode(m, int32(16797828))
														mBase = m.M
														v87 = m.ExcPending
														if v87 != 0 {
															return
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
															F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
															mBase = m.M
															v93 = m.ExcPending
															if v93 != 0 {
																return
															} else {
																F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																mBase = m.M
																v97 = m.ExcPending
																if v97 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																	mBase = m.M
																	v102 = m.ExcPending
																	if v102 != 0 {
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
													v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
													v54 = F_is_schema_publication(m, v53)
													mBase = m.M
													v55 = m.ExcPending
													if v55 != 0 {
														return
													} else {
														if v54 != 0 {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v84 = m.ExcPending
															if v84 != 0 {
																return
															} else {
																F_errcode(m, int32(16797828))
																mBase = m.M
																v87 = m.ExcPending
																if v87 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v13 + int32(4)
																	F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																	mBase = m.M
																	v93 = m.ExcPending
																	if v93 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																		mBase = m.M
																		v97 = m.ExcPending
																		if v97 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2215), int32(_a_F_AlterPublicationOwner_internal_4))
																			mBase = m.M
																			v102 = m.ExcPending
																			if v102 != 0 {
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
															v61 = m.ExcPending
															if v61 != 0 {
																return
															} else {
																v63 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v63, l2)
																mBase = m.M
																v65 = m.ExcPending
																if v65 != 0 {
																	return
																} else {
																	v67 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																	if v67 == int32(0) {
																		m.G0 = v9 + int32(16)
																		return
																	} else {
																		v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v72 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v71, v72, v72, v72)
																		mBase = m.M
																		v76 = m.ExcPending
																		if v76 != 0 {
																			return
																		} else {
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
						}
					}
				}
			}
		}
	}
}
func F_AlterSubscription_refresh(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v84 int32
	_ = v84
	var v96 int32
	_ = v96
	var v109 int32
	_ = v109
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
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
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
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v446 int32
	_ = v446
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v471 int32
	_ = v471
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v521 int32
	_ = v521
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
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
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int64
	_ = v637
	var v638 int32
	_ = v638
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v705 int32
	_ = v705
	var v708 int32
	_ = v708
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v883 int32
	_ = v883
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v947 int32
	_ = v947
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v968 int32
	_ = v968
	var v981 int32
	_ = v981
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v999 int32
	_ = v999
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1018 int32
	_ = v1018
	var v1034 int32
	_ = v1034
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1141 int32
	_ = v1141
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1211 int32
	_ = v1211
	var v1227 int32
	_ = v1227
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1237 int64
	_ = v1237
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1277 int32
	_ = v1277
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1317 int32
	_ = v1317
	var v1318 int32
	_ = v1318
	var v1328 int32
	_ = v1328
	var v1345 int32
	_ = v1345
	var v1375 int32
	_ = v1375
	var v1376 int64
	_ = v1376
	var v1380 int32
	_ = v1380
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	v5 = int32(0)
	v30 = m.G0
	v32 = v30 - int32(416)
	m.G0 = v32
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v37 = int32(105)
	goto L3
L2:
	;
	v37 = int32(114)
	goto L3
L3:
	;
	v39 = l0 + int32(24)
	v45 = v5
	v46 = v5
	v47 = v5
	v48 = v5
	v49 = v5
	v50 = v5
	v51 = v5
	v52 = v5
	v53 = v5
	v55 = int32(-1)
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
	m.G0 = v32 + int32(416)
	return
L7:
	;
	goto L6
L8:
	;
	if v55 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v1375 = int32(m.ExcTag)
	v1376 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1375 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L11:
	;
	if l3 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v217 = v45
	v218 = v46
	v219 = v48
	v220 = v49
	v221 = v53
	goto L13
L13:
	;
	if v221 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v45
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v45
	F_load_file(m, int32(_a_F_AlterSubscription_refresh_0), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L10
	} else {
		goto L20
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v45
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_refresh_1), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v45
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_2), int32(1084), int32(_a_F_AlterSubscription_refresh_3))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	goto L4
L20:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	if v123 == int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v129 = v126 ^ int32(1)
	goto L23
L22:
	;
	v129 = int32(0)
	goto L23
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[0]))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v39
	v142 = int32(1)
	v148 = m.T0[v133].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l3, v142, v142, v129&v142, v130, v32+int32(380))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	if v148 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v39
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L10
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[1]))
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2]))
	goto L32
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v39
	F_errcode(m, int32(100663808))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v39
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v32)+380))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+52)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v32)+48)) = v175
	F_errmsg(m, int32(_a_F_AlterSubscription_refresh_4), v32+int32(48))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v39
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_2), int32(1097), int32(_a_F_AlterSubscription_refresh_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	goto L4
L32:
	;
	v211 = v32 + int32(224)
	*(*int32)(unsafe.Add(mBase, uint32(v211)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v32 + int32(60)
	goto L35
L33:
	;
	v217 = v39
	v218 = v148
	v219 = v207
	v220 = v209
	v221 = int32(0)
	goto L13
L35:
	;
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_pg_qsort(m, v277, v407, int32(4), int32(506))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L10
	} else {
		goto L65
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2])) = v32 + int32(224)
	if l2 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[1])) = v219
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2])) = v220
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[0]))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v370)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	m.T0[v371].(func(*base.Module, int32))(m, v218)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L10
	} else {
		goto L63
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_check_publications(m, v218, l2)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v248 = F_fetch_relation_list(m, v218, v239)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L10
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v259 = int32(1)
	v262 = F_GetSubscriptionRelations(m, v250, v259, v259, int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	if v262 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	v268 = v264 << (uint(int32(2)) % 32)
	goto L48
L47:
	;
	v268 = int32(0)
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v277 = F_palloc(m, v268)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v287 = F_palloc(m, v268)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	if v262 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v291 = int32(0)
	v407 = v291
	v410 = v291
	goto L36
L52:
	;
	goto L53
L53:
	;
	v293 = int32(0)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v296 <= v293 {
		v407 = v293
		v410 = v293
		goto L36
	} else {
		goto L54
	}
L54:
	;
	v312 = v293
	v314 = v293
	v317 = v293
	goto L55
L55:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v262)+12))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v328+v312<<(uint(int32(2))%32))))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v342 = F_get_rel_relkind(m, v333)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L10
	} else {
		goto L57
	}
L56:
	;
	v407 = v359
	v410 = v360
	goto L36
L57:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	if v342 == int32(83) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v362 = v312 + int32(1)
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v262)+4))
	if v362 < v363 {
		v312 = v362
		v314 = v359
		v317 = v360
		goto L55
	} else {
		goto L62
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v287+v317<<(uint(int32(2))%32)))) = v344
	v359 = v314
	v360 = v317 + int32(1)
	goto L58
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277+v314<<(uint(int32(2))%32)))) = v344
	v359 = v314 + int32(1)
	v360 = v317
	goto L58
L62:
	;
	goto L56
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_pg_re_throw(m)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L10
	} else {
		goto L64
	}
L64:
	;
	goto L4
L65:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+41)))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_check_publications_origin_tables(m, v218, v436, l1, v435, v434, v277, v407, v433)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_pg_qsort(m, v287, v410, int32(4), int32(506))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_check_publications_origin_sequences(m, v218, v461, l1, v460, v287, v410, v459)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L10
	} else {
		goto L68
	}
L68:
	;
	if v248 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_pg_qsort(m, v708, v692, int32(4), int32(506))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L10
	} else {
		goto L95
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v483 = F_palloc(m, int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L10
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v499 = F_palloc(m, v488<<(uint(int32(2))%32))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L10
	} else {
		goto L74
	}
L73:
	;
	v686 = v47
	v689 = v483
	v692 = int32(0)
	v705 = v248 + int32(4)
	v708 = v483
	goto L69
L74:
	;
	v502 = v248 + int32(4)
	v503 = int32(0)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v248)+4))
	if v504 <= v503 {
		v686 = v499
		v689 = v50
		v692 = v504
		v705 = v502
		v708 = v499
		goto L69
	} else {
		goto L75
	}
L75:
	;
	v521 = v503
	goto L76
L76:
	;
	v537 = v521 << (uint(int32(2)) % 32)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v537+v538)))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v540)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v551 = int32(0)
	v554 = F_RangeVarGetRelidExtended(m, v541, int32(1), v551, v551, v551)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L10
	} else {
		goto L78
	}
L77:
	;
	v686 = v499
	v689 = v50
	v692 = v677
	v705 = v502
	v708 = v499
	goto L69
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+220)) = v554
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v565 = F_get_rel_relkind(m, v554)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v541)+12))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v541)+8))
	v569 = int32(*(*int8)(unsafe.Add(mBase, uint32(v540)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_CheckSubscriptionRelkind(m, v565, v569, v568, v567)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L10
	} else {
		goto L80
	}
L80:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v32)+220))
	*(*int32)(unsafe.Add(mBase, uint32(v499+v537))) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v592 = v32 + int32(220)
	v595 = F_bsearch(m, v592, v277, v407, int32(4), int32(506))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L10
	} else {
		goto L82
	}
L81:
	;
	v676 = v521 + int32(1)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	if v676 < v677 {
		v521 = v676
		goto L76
	} else {
		goto L94
	}
L82:
	;
	if v595 != 0 {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v607 = F_bsearch(m, v592, v287, v410, int32(4), int32(506))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	if v607 != 0 {
		goto L81
	} else {
		goto L85
	}
L85:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v32)+220))
	F_AddSubscriptionRelState(m, v609, v618, v37, int64(0), int32(1))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L10
	} else {
		goto L86
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v633 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L10
	} else {
		goto L87
	}
L87:
	;
	if v633 == int32(0) {
		goto L81
	} else {
		goto L88
	}
L88:
	;
	v637 = *(*int64)(unsafe.Add(mBase, uint32(v541)+8))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v32)+44)) = v638
	*(*int64)(unsafe.Add(mBase, uint32(v32)+36)) = v637
	if v565 == int32(83) {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v653 = int32(_a_F_AlterSubscription_refresh_5)
	goto L91
L90:
	;
	v653 = int32(_a_F_AlterSubscription_refresh_6)
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+32)) = v653
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_refresh_7), v32+int32(32))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v499
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_2), int32(1184), int32(_a_F_AlterSubscription_refresh_3))
	mBase = m.M
	v672 = m.ExcPending
	if v672 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	goto L81
L94:
	;
	goto L77
L95:
	;
	v721 = int32(0)
	if v721 < v407 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v737 = v51
	v739 = v721
	v742 = v721
	v743 = v721
	goto L99
L97:
	;
	v981 = v51
	v986 = v721
	v987 = v721
	goto L98
L98:
	;
	v999 = int32(0)
	if v999 < v410 {
		goto L129
	} else {
		goto L130
	}
L99:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v277+v739<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+216)) = v758
	if v248 != 0 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	v981 = v961
	v986 = v963
	v987 = v964
	goto L98
L101:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v762 = v761
	goto L103
L102:
	;
	v762 = int32(0)
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v775 = F_bsearch(m, v32+int32(216), v708, v762, int32(4), int32(506))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L10
	} else {
		goto L105
	}
L104:
	;
	v968 = v739 + int32(1)
	if v968 != v407 {
		v737 = v961
		v739 = v968
		v742 = v963
		v743 = v964
		goto L99
	} else {
		goto L128
	}
L105:
	;
	if v775 != 0 {
		v961 = v737
		v963 = v742
		v964 = v743
		goto L104
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v786 = F_palloc(m, int32(8))
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	if v742 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v800 = F_table_open(m, int32(_a_F_AlterSubscription_refresh_8), int32(8))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L10
	} else {
		goto L111
	}
L109:
	;
	v802 = v737
	v803 = v742
	goto L110
L110:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v813 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	v816 = F_GetSubscriptionRelState(m, v804, v813, v32+int32(208))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L10
	} else {
		goto L112
	}
L111:
	;
	v802 = v800
	v803 = v800
	goto L110
L112:
	;
	v818 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v827 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	F_RemoveSubscriptionRel(m, v818, v827)
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	*(*uint8)(unsafe.Add(mBase, uint32(v786)+4)) = uint8(v816)
	*(*int32)(unsafe.Add(mBase, uint32(v786))) = v830
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v841 = F_lappend(m, v743, v786)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L10
	} else {
		goto L114
	}
L114:
	;
	v843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	F_logicalrep_worker_stop(m, int32(1), v843, v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	if v816 != int32(114) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	v869 = v32 + int32(144)
	F_ReplicationOriginNameForLogicalRep(m, v858, v867, v869)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L10
	} else {
		goto L119
	}
L117:
	;
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v895 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L10
	} else {
		goto L121
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_replorigin_drop_by_name(m, v869, int32(1), int32(0))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L10
	} else {
		goto L120
	}
L120:
	;
	goto L118
L121:
	;
	if v895 == int32(0) {
		v961 = v802
		v963 = v803
		v964 = v841
		goto L104
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v907 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	v908 = F_get_rel_namespace(m, v907)
	mBase = m.M
	v909 = m.ExcPending
	if v909 != 0 {
		goto L10
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v918 = F_get_namespace_name(m, v908)
	mBase = m.M
	v919 = m.ExcPending
	if v919 != 0 {
		goto L10
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v32)+216))
	v929 = F_get_rel_name(m, v928)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L10
	} else {
		goto L125
	}
L125:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v32)+24)) = v931
	*(*int32)(unsafe.Add(mBase, uint32(v32)+20)) = v929
	*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = v918
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_refresh_9), v32+int32(16))
	mBase = m.M
	v947 = m.ExcPending
	if v947 != 0 {
		goto L10
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v802
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_2), int32(1261), int32(_a_F_AlterSubscription_refresh_3))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	v961 = v802
	v963 = v803
	v964 = v841
	goto L104
L128:
	;
	goto L100
L129:
	;
	v1014 = v52
	v1015 = v999
	v1018 = v986
	goto L132
L130:
	;
	v1175 = v52
	v1179 = v986
	goto L131
L131:
	;
	if v987 == int32(0) {
		goto L153
	} else {
		goto L154
	}
L132:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v287+v1015<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+140)) = v1034
	if v248 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	v1175 = v1155
	v1179 = v1158
	goto L131
L134:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v705)))
	v1038 = v1037
	goto L136
L135:
	;
	v1038 = int32(0)
	goto L136
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1014
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v1051 = F_bsearch(m, v32+int32(140), v708, v1038, int32(4), int32(506))
	mBase = m.M
	v1052 = m.ExcPending
	if v1052 != 0 {
		goto L10
	} else {
		goto L138
	}
L137:
	;
	v1161 = v1015 + int32(1)
	if v1161 != v410 {
		v1014 = v1155
		v1015 = v1161
		v1018 = v1158
		goto L132
	} else {
		goto L152
	}
L138:
	;
	if v1051 != 0 {
		v1155 = v1014
		v1158 = v1018
		goto L137
	} else {
		goto L139
	}
L139:
	;
	if v1018 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1014
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v1065 = F_table_open(m, int32(_a_F_AlterSubscription_refresh_8), int32(8))
	mBase = m.M
	v1066 = m.ExcPending
	if v1066 != 0 {
		goto L10
	} else {
		goto L143
	}
L141:
	;
	v1067 = v1014
	v1068 = v1018
	goto L142
L142:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v32)+140))
	F_RemoveSubscriptionRel(m, v1069, v1078)
	mBase = m.M
	v1080 = m.ExcPending
	if v1080 != 0 {
		goto L10
	} else {
		goto L144
	}
L143:
	;
	v1067 = v1065
	v1068 = v1065
	goto L142
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v1091 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L10
	} else {
		goto L145
	}
L145:
	;
	if v1091 == int32(0) {
		v1155 = v1067
		v1158 = v1068
		goto L137
	} else {
		goto L146
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v32)+140))
	v1104 = F_get_rel_namespace(m, v1103)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L10
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v1114 = F_get_namespace_name(m, v1104)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L10
	} else {
		goto L148
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v32)+140))
	v1125 = F_get_rel_name(m, v1124)
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L10
	} else {
		goto L149
	}
L149:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v32)+8)) = v1127
	*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v1125
	*(*int32)(unsafe.Add(mBase, uint32(v32))) = v1114
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_refresh_10), v32)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L10
	} else {
		goto L150
	}
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1067
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_2), int32(1289), int32(_a_F_AlterSubscription_refresh_3))
	mBase = m.M
	v1154 = m.ExcPending
	if v1154 != 0 {
		goto L10
	} else {
		goto L151
	}
L151:
	;
	v1155 = v1067
	v1158 = v1068
	goto L137
L152:
	;
	goto L133
L153:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[1])) = v219
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2])) = v220
	v1317 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[0]))
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v1317)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	m.T0[v1318].(func(*base.Module, int32))(m, v218)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L10
	} else {
		goto L164
	}
L154:
	;
	v1194 = int32(0)
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v987)+4))
	if v1195 <= v1194 {
		goto L153
	} else {
		goto L155
	}
L155:
	;
	v1211 = v1194
	goto L156
L156:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(v987)+12))
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v1227+v1211<<(uint(int32(2))%32))))
	v1232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1231)+4)))
	if v1232&int32(254) != int32(114) {
		goto L158
	} else {
		goto L159
	}
L157:
	;
	goto L153
L158:
	;
	v1237 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v32)+120)) = v1237
	*(*int64)(unsafe.Add(mBase, uint32(v32)+112)) = v1237
	*(*int64)(unsafe.Add(mBase, uint32(v32)+104)) = v1237
	*(*int64)(unsafe.Add(mBase, uint32(v32)+96)) = v1237
	*(*int64)(unsafe.Add(mBase, uint32(v32)+88)) = v1237
	*(*int64)(unsafe.Add(mBase, uint32(v32)+80)) = v1237
	*(*int64)(unsafe.Add(mBase, uint32(v32)+72)) = v1237
	*(*int64)(unsafe.Add(mBase, uint32(v32)+64)) = v1237
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1231)))
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	v1264 = v32 - int32(-64)
	F_ReplicationSlotNameForTablesync(m, v1254, v1253, v1264)
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L10
	} else {
		goto L161
	}
L159:
	;
	goto L160
L160:
	;
	v1280 = v1211 + int32(1)
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v987)+4))
	if v1280 < v1281 {
		v1211 = v1280
		goto L156
	} else {
		goto L163
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_ReplicationSlotDropAtPubNode(m, v218, v1264, int32(1))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L10
	} else {
		goto L162
	}
L162:
	;
	goto L160
L163:
	;
	goto L157
L164:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[1])) = v219
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2])) = v220
	if v1179 == int32(0) {
		goto L7
	} else {
		goto L165
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v32)+388)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v32)+384)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(v32)+392)) = v686
	*(*int32)(unsafe.Add(mBase, uint32(v32)+396)) = v689
	*(*int32)(unsafe.Add(mBase, uint32(v32)+400)) = v219
	*(*int32)(unsafe.Add(mBase, uint32(v32)+404)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v32)+408)) = v218
	*(*int32)(unsafe.Add(mBase, uint32(v32)+412)) = v217
	F_relation_close(m, v1179, int32(0))
	mBase = m.M
	v1345 = m.ExcPending
	if v1345 != 0 {
		goto L10
	} else {
		goto L166
	}
L166:
	;
	goto L9
L167:
	;
	v1380 = int32(v1376)
	m.G0 = v32
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+4))
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(v1380)))
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v1383)))
	if v32+int32(60) == v1386 {
		goto L170
	} else {
		goto L171
	}
L168:
	;
	m.ExcPending = 1
	goto L176
L169:
	;
	if v1390 != 0 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	v1388 = *(*int32)(unsafe.Add(mBase, uint32(v1383)+4))
	v1390 = v1388
	goto L172
L171:
	;
	v1390 = int32(0)
	goto L172
L172:
	;
	goto L169
L173:
	;
	v1391 = *(*int32)(unsafe.Add(mBase, uint32(v32)+412))
	v1392 = *(*int32)(unsafe.Add(mBase, uint32(v32)+408))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v32)+404))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v32)+400))
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v32)+396))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v32)+392))
	v1397 = *(*int32)(unsafe.Add(mBase, uint32(v32)+388))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v32)+384))
	v45 = v1391
	v46 = v1392
	v47 = v1396
	v48 = v1394
	v49 = v1393
	v50 = v1395
	v51 = v1397
	v52 = v1398
	v53 = v1382
	v55 = v1390
	goto L5
L174:
	;
	goto L175
L175:
	;
	F___wasm_longjmp(m, v1383, v1382)
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	return
L177:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ApplyLauncherMain(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v423 int32
	_ = v423
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v459 int32
	_ = v459
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v519 int32
	_ = v519
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int64
	_ = v561
	var v562 int64
	_ = v562
	var v571 int64
	_ = v571
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int64
	_ = v584
	var v585 int64
	_ = v585
	var v593 int64
	_ = v593
	var v601 int64
	_ = v601
	var v610 int64
	_ = v610
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v620 int64
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v764 int32
	_ = v764
	var v768 int32
	_ = v768
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v830 int32
	_ = v830
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v846 int32
	_ = v846
	v2 = int32(0)
	v22 = m.G0
	v24 = v22 - int32(16)
	m.G0 = v24
	v28 = F_errstart(m, int32(14), v2)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_errmsg_internal(m, int32(_a_F_ApplyLauncherMain_0), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_before_shmem_exit(m, int32(1054), int64(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	F_errfinish(m, int32(_a_F_ApplyLauncherMain_1), int32(1212), int32(_a_F_ApplyLauncherMain_2))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[0]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = v46
	v52 = m.G0
	v54 = v52 - int32(32)
	m.G0 = v54
	v57 = int32(967)
	switch v57 {
	case 0, 2:
		goto L10
	default:
		goto L11
	}
L9:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L19
	}
L10:
	;
	F_sigemptyset(m, v54+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(268435456)
	switch v57 {
	case 0:
		goto L15
	default:
		goto L13
	case 2:
		goto L14
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[2])) = int32(965)
	goto L10
L12:
	;
	goto L17
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = int32(_a_F_ApplyLauncherMain_3)
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = int32(0)
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v54)+12)) = int32(-2)
	goto L12
L17:
	;
	goto L18
L18:
	;
	v86 = F___sigaction(m, int32(1), v54+int32(12), int32(0))
	mBase = m.M
	m.G0 = v54 + int32(32)
	goto L9
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[3]))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+192)))
	if v94&int32(2) != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	v143 = F_SearchNamedReplicationSlot(m, int32(_a_F_ApplyLauncherMain_4), int32(1))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L34
	}
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L31
	}
L22:
	;
	v97 = int32(0)
	F_InitPostgres(m, v97, v97, v97, v97, v97, v97)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[4]))
	if v106 != int32(1) {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[4])) = int32(2)
	goto L20
L27:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(_a_F_ApplyLauncherMain_5), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_ApplyLauncherMain_6), int32(882), int32(_a_F_ApplyLauncherMain_7))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errmsg(m, int32(_a_F_ApplyLauncherMain_8), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errfinish(m, int32(_a_F_ApplyLauncherMain_6), int32(892), int32(_a_F_ApplyLauncherMain_7))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
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
	if v143 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_ReplicationSlotAcquire(m, int32(_a_F_ApplyLauncherMain_4), int32(1), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v164 = v2
	goto L39
L38:
	;
	goto L37
L39:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[5]))
	if v172 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v175 = int32(0)
	v177 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[6]))
	v182 = F_AllocSetContextCreateInternal(m, v177, int32(_a_F_ApplyLauncherMain_9), v175, int32(_a_F_ApplyLauncherMain_10), int32(_a_F_ApplyLauncherMain_11))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v184 = int32(_a_F_ApplyLauncherMain_12)
	v185 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[7])) = v182
	F_StartTransactionCommand(m)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v192 = F_table_open(m, int32(_a_F_ApplyLauncherMain_13), int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v194 = int32(0)
	v196 = F_table_beginscan_catalog(m, v192, v194, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v205 = v175
	goto L49
L49:
	;
	v219 = F_heap_getnext(m, v196)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v196)))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+188))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	m.T0[v254].(func(*base.Module, int32))(m, v196)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L58
	}
L51:
	;
	if v219 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v221 = int32(_a_F_ApplyLauncherMain_12)
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[7]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+22)))
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[7])) = v182
	v228 = F_palloc0(m, int32(72))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	goto L50
L55:
	;
	v230 = v223 + v224
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+4)) = v231
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v230)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+8)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v230)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v228)+28)) = v235
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+84)))
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+33)) = uint8(v237)
	v241 = F_pstrdup(m, v230+int32(16))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v228)+24)) = v241
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+92)))
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+41)) = uint8(v244)
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(v228)+48)) = uint8(v246)
	v248 = F_lappend(m, v205, v228)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[7])) = v222
	v205 = v248
	goto L49
L58:
	;
	F_relation_close(m, v192, int32(1))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	if v205 != 0 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[7])) = v185
	F_list_free(m, v164)
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
		goto L1
	} else {
		goto L194
	}
L62:
	;
	F_ReplicationSlotDropAcquired(m, int32(0))
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L193
	}
L63:
	;
	v715 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8]))
	if v715 == int32(0) {
		v788 = v697
		v792 = v701
		goto L61
	} else {
		goto L177
	}
L64:
	;
	v262 = int32(1)
	v263 = int32(_a_F_ApplyLauncherMain_14)
	v264 = int32(0)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v269 <= v264 {
		v696 = v264
		v697 = v263
		v698 = v262
		v701 = v264
		v706 = v264
		goto L63
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v687 = int32(0)
	v688 = int32(_a_F_ApplyLauncherMain_14)
	v690 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8]))
	if v690 == v687 {
		v788 = v688
		v792 = v687
		goto L61
	} else {
		goto L176
	}
L67:
	;
	v275 = v264
	v276 = v263
	v277 = v262
	v280 = v264
	v283 = v264
	v284 = v264
	v285 = v264
	goto L68
L68:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v205)+12))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v293+v283<<(uint(int32(2))%32))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+41)))
	if v298 == int32(0) {
		v389 = v277
		v390 = v280
		v391 = v284
		v392 = v285
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v696 = v665
	v697 = v666
	v698 = v667
	v701 = v390
	v706 = v392
	goto L63
L70:
	;
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+33)))
	if v393 != int32(1) {
		v665 = v275
		v666 = v276
		v667 = v389
		goto L106
	} else {
		goto L107
	}
L71:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8]))
	if v302 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v307 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	v329 = int32(1)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+48)))
	if v330 != v329 {
		v389 = v277
		v390 = v280
		v391 = v284
		v392 = v329
		goto L70
	} else {
		goto L83
	}
L75:
	;
	if v307 != 0 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	F_errmsg(m, int32(_a_F_ApplyLauncherMain_15), int32(0))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v319 = int32(0)
	F_ReplicationSlotCreate(m, int32(_a_F_ApplyLauncherMain_4), v319, v319, v319, v319, v319, v319)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	F_errfinish(m, int32(_a_F_ApplyLauncherMain_1), int32(1643), int32(_a_F_ApplyLauncherMain_16))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	F_reset_conflict_slot_xmin_to_safe_horizon(m)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	goto L74
L83:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v297)+8))
	v334 = F_list_append_unique_oid(m, v280, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+33)))
	v337 = v277 & v336
	if v284 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8]))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)+96))
	if v342 != 0 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v389 = v337
	v390 = v334
	v391 = int32(1)
	v392 = v329
	goto L70
L88:
	;
	v343 = int32(0)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v297)+8))
	if v164 == v343 {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	goto L90
L90:
	;
	F_reset_conflict_slot_xmin_to_safe_horizon(m)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L105
	}
L91:
	;
	if v383 != 0 {
		v389 = v337
		v390 = v334
		v391 = v343
		v392 = v329
		goto L70
	} else {
		goto L104
	}
L92:
	;
	v383 = int32(0)
	goto L91
L93:
	;
	goto L94
L94:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v351 <= int32(0) {
		v377 = v343
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v383 = v377
	goto L91
L96:
	;
	v354 = int32(0)
	if v354 < v351 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v357 = v351
	goto L99
L98:
	;
	v357 = v354
	goto L99
L99:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v360 = int32(0)
	goto L100
L100:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v358+v360<<(uint(int32(2))%32))))
	v369 = base.B2i32(v368 == v344)
	if v368 == v344 {
		v377 = v369
		goto L95
	} else {
		goto L102
	}
L101:
	;
	v377 = v369
	goto L95
L102:
	;
	v371 = v360 + int32(1)
	if v371 != v357 {
		v360 = v371
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	goto L90
L105:
	;
	goto L87
L106:
	;
	v684 = v283 + int32(1)
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	if v684 < v685 {
		v275 = v665
		v276 = v666
		v277 = v667
		v280 = v390
		v283 = v684
		v284 = v391
		v285 = v392
		goto L68
	} else {
		goto L175
	}
L107:
	;
	v398 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[9]))
	v402 = F_LWLockAcquire(m, v398+int32(_a_F_ApplyLauncherMain_17), int32(1))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v405 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[10]))
	if v405 <= int32(0) {
		v459 = int32(0)
		goto L109
	} else {
		goto L110
	}
L109:
	;
	if v389&int32(1) == int32(0) {
		goto L121
	} else {
		goto L122
	}
L110:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v411 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[0]))
	v423 = int32(0)
	goto L111
L111:
	;
	v437 = v411 + int32(16) + v423<<(uint(int32(7))%32)
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+16)))
	if v438 != int32(1) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v459 = int32(0)
	goto L109
L113:
	;
	v454 = v423 + int32(1)
	if v454 != v405 {
		v423 = v454
		goto L111
	} else {
		goto L118
	}
L114:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v437)))
	if v441 == int32(4) {
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v437)+32))
	if base.B2i32(v444 != v408)|base.B2i32(v441 != int32(3)) != 0 {
		goto L113
	} else {
		goto L116
	}
L116:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v437)+36))
	if v449 == int32(0) {
		v459 = v437
		goto L109
	} else {
		goto L117
	}
L117:
	;
	goto L113
L118:
	;
	goto L112
L119:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = v540
	F_logicalrep_launcher_attach_dshmem(m)
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L145
	}
L120:
	;
	if v459 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	v489 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[9]))
	F_LWLockRelease(m, v489+int32(_a_F_ApplyLauncherMain_17))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L125
	}
L122:
	;
	v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+41)))
	if v482 != int32(1) {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+48)))
	if v485 == int32(1) {
		goto L120
	} else {
		goto L124
	}
L124:
	;
	goto L121
L125:
	;
	if v459 != 0 {
		v665 = v275
		v666 = v276
		v667 = v389
		goto L106
	} else {
		goto L126
	}
L126:
	;
	v539 = v389
	goto L119
L127:
	;
	v498 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[9]))
	F_LWLockRelease(m, v498+int32(_a_F_ApplyLauncherMain_17))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v505 = base.AtomicRmwXchg32(m, v459, int32(56), int32(1))
	if v505 != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v539 = int32(0)
	goto L119
L131:
	;
	F_s_lock(m, v459+int32(56), int32(_a_F_ApplyLauncherMain_18))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v459)+72))
	v512 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v459)+56)), uint32(v512))
	if v511 == v512 {
		v530 = v275
		goto L135
	} else {
		goto L136
	}
L134:
	;
	goto L133
L135:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[9]))
	F_LWLockRelease(m, v534+int32(_a_F_ApplyLauncherMain_17))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L144
	}
L136:
	;
	if v275 == int32(0) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v530 = v511
	goto L135
L138:
	;
	v519 = int32(3)
	if base.B2i32(base.Ui32(v275) < base.Ui32(v519))|base.B2i32(base.Ui32(v511) < base.Ui32(v519)) == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	if v511-v275 < int32(0) {
		goto L137
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	if base.Ui32(v275) <= base.Ui32(v511) {
		v530 = v275
		goto L135
	} else {
		goto L143
	}
L142:
	;
	v530 = v275
	goto L135
L143:
	;
	goto L137
L144:
	;
	v665 = v530
	v666 = v276
	v667 = base.B2i32(v511 != int32(0))
	goto L106
L145:
	;
	v545 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[11]))
	v549 = F_dshash_find(m, v545, v24+int32(4), int32(0))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	v659 = v615 - v613
	if v276 < v659 {
		goto L172
	} else {
		goto L173
	}
L147:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v621
	F_logicalrep_launcher_attach_dshmem(m)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L161
	}
L148:
	;
	if v549 == int32(0) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v556 = m.G0
	v557 = int32(16)
	v558 = v556 - v557
	m.G0 = v558
	F_gettimeofday(m, v558)
	mBase = m.M
	v561 = *(*int64)(unsafe.Add(mBase, uint32(v558)))
	v562 = int64(*(*int32)(unsafe.Add(mBase, uint32(v558)+8)))
	m.G0 = v558 + v557
	goto L152
L150:
	;
	goto L151
L151:
	;
	v571 = *(*int64)(unsafe.Add(mBase, uint32(v549)+8))
	v573 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[11]))
	F_dshash_release_lock(m, v573, v549)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L153
	}
L152:
	;
	v620 = v562 + v561*int64(1000000) - int64(946684800000000)
	goto L147
L153:
	;
	v579 = m.G0
	v580 = int32(16)
	v581 = v579 - v580
	m.G0 = v581
	F_gettimeofday(m, v581)
	mBase = m.M
	v584 = *(*int64)(unsafe.Add(mBase, uint32(v581)))
	v585 = int64(*(*int32)(unsafe.Add(mBase, uint32(v581)+8)))
	m.G0 = v581 + v580
	v593 = v585 + v584*int64(1000000) - int64(946684800000000)
	goto L154
L154:
	;
	if v571 == int64(0) {
		v620 = v593
		goto L147
	} else {
		goto L155
	}
L155:
	;
	if v593 <= v571 {
		v613 = int32(0)
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v615 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[12]))
	if v613 < v615 {
		goto L146
	} else {
		goto L160
	}
L157:
	;
	goto L156
L158:
	;
	v601 = v593 - v571
	if base.B2i32(int64(0) < v571)^base.B2i32(v601 < v593)|base.B2i32(int64(2147483646000) < v601) != 0 {
		v613 = int32(2147483647)
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v610 = base.I64_div_s(v601+int64(999), int64(1000))
	v613 = base.I32_wrap_i64(v610)
	goto L157
L160:
	;
	v620 = v593
	goto L147
L161:
	;
	v626 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[11]))
	v631 = F_dshash_find_or_insert_extended(m, v626, v24+int32(12), v24+int32(11))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v631)+8)) = v620
	v635 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[11]))
	F_dshash_release_lock(m, v635, v631)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v297)+8))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v297)+4))
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v297)+24))
	v642 = *(*int32)(unsafe.Add(mBase, uint32(v297)+28))
	v643 = int32(0)
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+41)))
	if v645 == int32(1) {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297)+48)))
	v650 = v648
	goto L166
L165:
	;
	v650 = int32(0)
	goto L166
L166:
	;
	v653 = F_logicalrep_worker_launch(m, int32(3), v639, v640, v641, v642, v643, v643, v650&int32(1))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	if v653 != 0 {
		v665 = v275
		v666 = v276
		v667 = v539
		goto L106
	} else {
		goto L168
	}
L168:
	;
	v656 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[12]))
	if v276 < v656 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v658 = v276
	goto L171
L170:
	;
	v658 = v656
	goto L171
L171:
	;
	v665 = v275
	v666 = v658
	v667 = v539
	goto L106
L172:
	;
	v661 = v276
	goto L174
L173:
	;
	v661 = v659
	goto L174
L174:
	;
	v665 = v275
	v666 = v661
	v667 = v539
	goto L106
L175:
	;
	goto L69
L176:
	;
	v764 = v688
	v768 = v687
	goto L62
L177:
	;
	if v706 == int32(0) {
		v764 = v697
		v768 = v701
		goto L62
	} else {
		goto L178
	}
L178:
	;
	if v698&int32(1) == int32(0) {
		v788 = v697
		v792 = v701
		goto L61
	} else {
		goto L179
	}
L179:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v715)+96))
	if v724 == v696 {
		v788 = v697
		v792 = v701
		goto L61
	} else {
		goto L180
	}
L180:
	;
	v728 = base.AtomicRmwXchg32(m, v715, int32(0), int32(1))
	if v728 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	F_s_lock(m, v715, int32(_a_F_ApplyLauncherMain_18))
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L184
	}
L182:
	;
	goto L183
L183:
	;
	v733 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8]))
	*(*int32)(unsafe.Add(mBase, uint32(v733)+96)) = v696
	*(*int32)(unsafe.Add(mBase, uint32(v733)+16)) = v696
	v736 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v733))), uint32(v736))
	v741 = F_errstart(m, int32(14), v736)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L185
	}
L184:
	;
	goto L183
L185:
	;
	if v741 != 0 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8]))
	v745 = *(*int32)(unsafe.Add(mBase, uint32(v744)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v745
	F_errmsg_internal(m, int32(_a_F_ApplyLauncherMain_19), v24)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L1
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L1
	} else {
		goto L191
	}
L189:
	;
	F_errfinish(m, int32(_a_F_ApplyLauncherMain_1), int32(1555), int32(_a_F_ApplyLauncherMain_20))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	v788 = v697
	v792 = v701
	goto L61
L193:
	;
	v788 = v764
	v792 = v768
	goto L61
L194:
	;
	v809 = F_list_copy(m, v792)
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	F_MemoryContextDelete(m, v182)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v814 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[13]))
	v817 = F_WaitLatch(m, v814, int32(41), v788, int32(83886088))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L1
	} else {
		goto L198
	}
L197:
	;
	v838 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[14]))
	if v838 == int32(0) {
		v164 = v809
		goto L39
	} else {
		goto L203
	}
L198:
	;
	if v817&int32(1) == int32(0) {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	v824 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[13]))
	v825 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v824))) = v825
	v830 = base.AtomicRmwOr32(m, v825, int32(_a_F_ApplyLauncherMain_21), v825)
	goto L200
L200:
	;
	v832 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[5]))
	if v832 == int32(0) {
		goto L197
	} else {
		goto L201
	}
L201:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v836 = m.ExcPending
	if v836 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	goto L197
L203:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[14])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v164 = v809
	goto L39
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
	var v38 int32
	_ = v38
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[0]))
	if v13 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L19
	}
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[1]))
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
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[0]))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+48)))
	if v23 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[1]))
	F_LWLockRelease(m, v73+int32(2304))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L18
	}
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[2]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(-1)
	goto L7
L11:
	;
	v32 = v27 + int32(8)
	v38 = int32(0)
	goto L12
L12:
	;
	v42 = v32 + v38<<(uint(int32(2))%32)
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
	v56 = int32(_a_F_AtAbort_Twophase_0)
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[2]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v58
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v22
	goto L7
L14:
	;
	v46 = v38 + int32(1)
	if v28 != v46 {
		v38 = v46
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
	*(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[0])) = int32(0)
	goto L4
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22
	F_errmsg_internal(m, int32(_a_F_AtAbort_Twophase_1), v10)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_AtAbort_Twophase_2), int32(658), int32(_a_F_AtAbort_Twophase_3))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
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
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
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
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v146 int64
	_ = v146
	var v157 int32
	_ = v157
	var v171 int64
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[0]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[0])) = int32(0)
	goto L3
L6:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[1])) = uint8(v19)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[2]))
	if v22 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[3]))
	v17 = F_GetBackendTypeDesc(m, v16)
	mBase = m.M
	goto L9
L9:
	;
	goto L6
L10:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+12)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[2]))
	v129 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v128)+20)), uint32(v129))
	v132 = int32(_a_F_AuxiliaryProcessMainCommon_0)
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5])) = v122
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v138 = base.I32_div_s(v122-v135, int32(768))
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[6])) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v122)+576)) = v129
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+572)) = uint8(v129)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+416)) = v129
	v146 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v122)+4)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v122)+48)) = v146
	*(*int64)(unsafe.Add(mBase, uint32(v122)+40)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v122)+28)) = v129
	*(*int64)(unsafe.Add(mBase, uint32(v122)+20)) = v146
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+336)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v122)+16)) = v157
	*(*int64)(unsafe.Add(mBase, uint32(v122)+384)) = v146
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+344)) = uint16(v129)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+36)) = uint8(v129)
	*(*int64)(unsafe.Add(mBase, uint32(v122)+392)) = v146
	v171 = base.AtomicRmwXchg64(m, v122, int32(408), v146)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+340)) = v129
	F_OwnLatch(m, v173+int32(316))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L41
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L38
	}
L12:
	;
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L35
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[7]))
	if v26 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5]))
	if v30 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[8])))
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
	v39 = v22
	goto L18
L18:
	;
	v42 = base.AtomicRmwXchg32(m, v39, int32(20), int32(1))
	if v42 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[2]))
	v39 = v38
	goto L18
L20:
	;
	F_s_lock(m, v39+int32(20), int32(_a_F_AuxiliaryProcessMainCommon_1))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[2]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+72))
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[9])) = v51
	goto L24
L23:
	;
	goto L22
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[7]))
	v57 = int32(0)
	goto L25
L25:
	;
	v62 = v55 + v57*int32(768)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if v63 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[2]))
	v80 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v79)+20)), uint32(v80))
	F_errstart_cold(m, int32(22), v80)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L32
	}
L27:
	;
	v122 = v62
	v123 = v57
	goto L10
L28:
	;
	goto L29
L29:
	;
	v67 = v57 | int32(1)
	v70 = v55 + v67*int32(768)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+12))
	if v71 == int32(0) {
		v122 = v70
		v123 = v67
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v75 = v57 + int32(2)
	if v75 != int32(38) {
		v57 = v75
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	F_errmsg_internal(m, int32(_a_F_AuxiliaryProcessMainCommon_2), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_AuxiliaryProcessMainCommon_3), int32(658), int32(_a_F_AuxiliaryProcessMainCommon_4))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_AuxiliaryProcessMainCommon_5), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_AuxiliaryProcessMainCommon_3), int32(627), int32(_a_F_AuxiliaryProcessMainCommon_4))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_AuxiliaryProcessMainCommon_6), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_AuxiliaryProcessMainCommon_3), int32(630), int32(_a_F_AuxiliaryProcessMainCommon_4))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
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
	v181 = m.ExcPending
	if v181 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[10])) = v183 + int32(648)
	goto L43
L43:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5]))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+332))
	F_PGSemaphoreReset(m, v190)
	mBase = m.M
	F_on_shmem_exit(m, int32(1229), base.I64_extend_i32_u(v123))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[8])))
	if v197 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_AttachSharedMemoryStructs(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	F_BaseInit(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v204 = int32(0)
	F_ProcSignalInit(m, v204, v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_InitializeProcessXLogLogicalInfo(m)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_CreateAuxProcessResourceOwner(m)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	F_pgstat_beinit(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_pgstat_bestart_initial(m)
	mBase = m.M
	F_pgstat_bestart_final(m)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_before_shmem_exit(m, int32(977), int64(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[11])) = int32(2)
	return
}
func F_accumulate_append_subpath(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v33 int32
	_ = v33
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v6 - int32(293) {
	case 0:
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v9 == int32(1) {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			if v12 != 0 {
				if l2 == int32(0) {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v33 = F_lappend(m, v32, l0)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
						return
					}
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
					v18 = F_list_copy_tail(m, v17, v12)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return
					} else {
						v20 = F_list_concat(m, v16, v18)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v20
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
							v25 = F_list_copy_head(m, v23, v24)
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								v28 = F_list_concat(m, v27, v25)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v28
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
									v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
									v47 = F_lappend(m, v44, v46)
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
										v51 = F_list_concat(m, v47, v50)
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v51
											return
										}
									}
								}
							}
						}
					}
				}
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
				v39 = F_list_concat(m, v37, v38)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					v47 = F_lappend(m, v44, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
						v51 = F_list_concat(m, v47, v50)
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v51
							return
						}
					}
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
			v39 = F_list_concat(m, v37, v38)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
				v47 = F_lappend(m, v44, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
					v51 = F_list_concat(m, v47, v50)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v51
						return
					}
				}
			}
		}
	case 1:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		v39 = F_list_concat(m, v37, v38)
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v39
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
			v47 = F_lappend(m, v44, v46)
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
				v51 = F_list_concat(m, v47, v50)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v51
					return
				}
			}
		}
	default:
		v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v33 = F_lappend(m, v32, l0)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v33
			return
		}
	}
}
func F_acldefault(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v25 int64
	_ = v25
	var v29 int64
	_ = v29
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v84 int32
	_ = v84
	v3 = int64(0)
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v15 = int32(1)
	switch l0 - int32(6) {
	case 0:
		v55 = int64(0)
		v56 = v3
		v57 = int32(0)
		v58 = v15
		v59 = int32(1)
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
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
			F_errmsg_internal(m, int32(_a_F_acldefault_0), v12)
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_acldefault_1), int32(895), int32(_a_F_acldefault_2))
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
		v55 = int64(3584)
		v56 = int64(3072)
		v57 = int32(2)
		v58 = int32(0)
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 6, 15, 44:
		v29 = int64(256)
		v55 = v29
		v56 = v29
		v57 = int32(2)
		v58 = int32(0)
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 10, 11:
		v55 = int64(256)
		v56 = v3
		v57 = v15
		v58 = v15
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 13:
		v25 = int64(128)
		v55 = v25
		v56 = v25
		v57 = int32(2)
		v58 = int32(0)
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 16:
		v55 = int64(6)
		v56 = v3
		v57 = v15
		v58 = v15
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 21:
		v55 = int64(12288)
		v56 = v3
		v57 = v15
		v58 = v15
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 31:
		v55 = int64(768)
		v56 = v3
		v57 = v15
		v58 = v15
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 32:
		v55 = int64(262)
		v56 = v3
		v57 = v15
		v58 = v15
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 36:
		v55 = int64(16511)
		v56 = v3
		v57 = v15
		v58 = v15
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	case 37:
		v55 = int64(512)
		v56 = v3
		v57 = v15
		v58 = v15
		v59 = v5
		v63 = v57<<(uint(int32(4))%32) + int32(24)
		v64 = F_palloc0(m, v63)
		mBase = m.M
		v65 = m.ExcPending
		if v65 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v64)+20)) = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+12)) = int32(1033)
			*(*int64)(unsafe.Add(mBase, uint32(v64)+4)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(v64))) = v63 << (uint(int32(2)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v64)+16)) = v57
			if v58 != 0 {
				v84 = v64 + int32(24)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v64)+32)) = v56
				*(*int32)(unsafe.Add(mBase, uint32(v64)+28)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v64)+24)) = int32(0)
				v84 = v64 + int32(40)
			}
			if v59 == int32(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v84))) = l1
			} else {
			}
			m.G0 = v12 + int32(16)
			return v64
		}
	}
}
func F_aclexplode(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v223 int64
	_ = v223
	var v225 int64
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int64
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v271 int64
	_ = v271
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum(m, v16)
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	if v22 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_check_acl(m, v17)
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
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	goto L34
L6:
	;
	v27 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v29 = int32(_a_F_aclexplode_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_aclexplode[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_aclexplode[0])) = v32
	v35 = F_CreateTemplateTupleDesc(m, int32(4))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v35, int32(1), int32(_a_F_aclexplode_1), int32(26), int32(-1), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_TupleDescInitEntry(m, v35, int32(2), int32(_a_F_aclexplode_2), int32(26), int32(-1), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	F_TupleDescInitEntry(m, v35, int32(3), int32(_a_F_aclexplode_3), int32(25), int32(-1), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v35, int32(4), int32(_a_F_aclexplode_4), int32(16), int32(-1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v65 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v65 < v74 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v152 = F_BlessTupleDesc(m, v35)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L32
	}
L14:
	;
	v78 = v35 + int32(28)
	v85 = v65
	v86 = v74
	v88 = v65
	goto L18
L15:
	;
	v142 = v65
	v149 = v74
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v142
	goto L13
L17:
	;
	v142 = v136
	v149 = v115
	goto L16
L18:
	;
	v94 = v78 + v74<<(uint(int32(3))%32) + v85*int32(100)
	v97 = v78 + v85<<(uint(int32(3))%32)
	if v74 != v86 {
		v115 = v86
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v136 = v74
	goto L17
L20:
	;
	v116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+2)))
	if v116 <= int32(0) {
		v136 = v85
		goto L17
	} else {
		goto L28
	}
L21:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+7)))
	if v99 != int32(118) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v115 = v85
	goto L20
L23:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+4)))
	if v102 != int32(1) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+6)))
	if v105&int32(6) != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+2)))
	if v108 <= int32(0) {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+90)))
	if v111 != int32(118) {
		v115 = v74
		goto L20
	} else {
		goto L27
	}
L27:
	;
	goto L22
L28:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+90)))
	if v119 == int32(118) {
		v136 = v85
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+5)))
	v128 = (v88 + v122 - int32(1)) & (int32(0) - v122)
	if int32(_a_F_aclexplode_5) < v128 {
		v136 = v85
		goto L17
	} else {
		goto L30
	}
L30:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v97))) = uint16(v128)
	v134 = v85 + int32(1)
	if v134 != v74 {
		v85 = v134
		v86 = v115
		v88 = v128 + v116
		goto L18
	} else {
		goto L31
	}
L31:
	;
	goto L19
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v152
	v157 = F_palloc_mul(m, int32(4), int32(2))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v157))) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v157
	*(*int32)(unsafe.Add(mBase, _c_F_aclexplode[0])) = v30
	goto L5
L34:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v170 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v178 = v170
	goto L37
L36:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v178 = (v171<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L37
L37:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v183 = v180
	goto L41
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L54
	}
L39:
	;
	m.G0 = v14 + int32(48)
	return v271
L40:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L53
	}
L41:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v192 <= v183 {
		goto L40
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(0)
	v223 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v215)+4)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v223
	v225 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v215))))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v225
	if base.Ui32(int32(15)) <= base.Ui32(v209) {
		goto L38
	} else {
		goto L49
	}
L43:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v196 = v194 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v196
	if v196 == int32(15) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v200 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v169)+4)) = v200
	v204 = v183 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v204
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v206 <= v204 {
		goto L40
	} else {
		goto L47
	}
L45:
	;
	v208 = v183
	v209 = v196
	goto L46
L46:
	;
	v211 = base.I64_extend_i32_u(v209)
	v212 = int64(1) << (uint(v211) % 64)
	v215 = v178 + v17 + v208<<(uint(int32(4))%32)
	v216 = *(*int64)(unsafe.Add(mBase, uint32(v215)+8))
	if base.I32_wrap_i64(v212&v216) == int32(0) {
		v183 = v208
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v208 = v204
	v209 = v200
	goto L46
L48:
	;
	goto L42
L49:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v209<<(uint(int32(2))%32))+uint32(_c_F_aclexplode[1])))
	v232 = F_cstring_to_text(m, v231)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = base.I64_extend_i32_u(v232)
	v236 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v215)+12)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = int64(base.Ui64(v236)>>(uint(v211)%64)) & int64(1)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v168)+28))
	v246 = F_heap_form_tuple(m, v241, v14+int32(16), v14+int32(12))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v246)+16))
	v249 = F_HeapTupleHeaderGetDatum(m, v248)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v168)))
	*(*int64)(unsafe.Add(mBase, uint32(v168))) = v251 + int64(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+20)) = int32(1)
	v271 = v249
	goto L39
L53:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v262)+20)) = int32(2)
	v265 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v265)
	v271 = int64(0)
	goto L39
L54:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v212)
	F_errmsg_internal(m, int32(_a_F_aclexplode_6), v14)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_aclexplode_7), int32(1793), int32(_a_F_aclexplode_8))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_aclitemout(m *base.Module, l0 int32) int64 {
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v165 int64
	_ = v165
	var v168 int64
	_ = v168
	var v170 int64
	_ = v170
	var v175 int32
	_ = v175
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v181 int64
	_ = v181
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int64
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
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
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v150 = F_strlen(m, v19)
	mBase = m.M
	v151 = v150 + v19
	v152 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v151))) = uint8(v152)
	v156 = v151 + int32(1)
	v165 = int64(0)
	goto L34
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_aclitemout[0]))
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v47 = v45 + int32(4)
	v48 = int32(1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+4)))
	if v49 == int32(0) {
		v93 = v48
		v101 = v19
		goto L12
	} else {
		goto L13
	}
L6:
	;
	v32 = F_SearchSysCache1(m, int32(11), base.I64_extend_i32_u(v25))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v36 = v25
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v36
	v41 = F_pg_sprintf(m, v19, int32(_a_F_aclitemout_0), v15+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	if v32 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v36 = v34
	goto L8
L11:
	;
	goto L3
L12:
	;
	v102 = v101
	v103 = v47
	goto L23
L13:
	;
	v52 = v47
	v54 = v49
	goto L14
L14:
	;
	if int32(0) <= base.I32_extend8_s(v54) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v93 = v48
	v101 = v19
	goto L12
L16:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v88 != 0 {
		v52 = v52 + int32(1)
		v54 = v88
		goto L14
	} else {
		goto L22
	}
L17:
	;
	goto L20
L18:
	;
	goto L19
L19:
	;
	v81 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v19))) = uint8(v81)
	v93 = int32(0)
	v101 = v19 + int32(1)
	goto L12
L20:
	;
	if base.B2i32(base.Ui32(v54-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v54|int32(32)-int32(97)) < base.Ui32(int32(26)))|base.B2i32(v54 == int32(95)) != 0 {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	goto L15
L23:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v114 != int32(34) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v132)
	v134 = int32(1)
	v102 = v131 + v134
	v103 = v103 + v134
	goto L23
L26:
	;
	if v114 != 0 {
		v131 = v102
		v132 = v114
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v126 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v126)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	v131 = v102 + int32(1)
	v132 = v130
	goto L25
L29:
	;
	if v93 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v121 = v102
	goto L32
L31:
	;
	v117 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v117)
	v121 = v102 + int32(1)
	goto L32
L32:
	;
	v122 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v122)
	F_ReleaseCatCache(m, v32)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L3
L34:
	;
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v170 = int64(1) << (uint(v165) % 64)
	if v168&v170 != int64(0) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v196 = int32(47)
	*(*uint16)(unsafe.Add(mBase, uint32(v191))) = uint16(v196)
	v199 = v191 + int32(1)
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_aclitemout[0]))
	if v201 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L36:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v165))+uint32(_c_F_aclitemout[1]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v175)
	v177 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
	v180 = v156 + int32(1)
	v181 = v177
	goto L38
L37:
	;
	v180 = v156
	v181 = v168
	goto L38
L38:
	;
	if int64(base.Ui64(v181)>>(uint(int64(32))%64))&v170 != int64(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v187 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v180))) = uint8(v187)
	v191 = v180 + int32(1)
	goto L41
L40:
	;
	v191 = v180
	goto L41
L41:
	;
	v193 = v165 + int64(1)
	if v193 != int64(15) {
		v156 = v191
		v165 = v193
		goto L34
	} else {
		goto L42
	}
L42:
	;
	goto L35
L43:
	;
	m.G0 = v15 + int32(32)
	return base.I64_extend_i32_u(v19)
L44:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v308
	v311 = F_pg_sprintf(m, v199, int32(_a_F_aclitemout_0), v15)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L70
	}
L45:
	;
	v205 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v17)+4)))
	v206 = F_SearchSysCache1(m, int32(11), v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v206 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+22)))
	v212 = v210 + v211
	v214 = v212 + int32(4)
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212)+4)))
	if v215 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v272 = v260
	v277 = v214
	goto L59
L49:
	;
	v260 = v199
	v262 = int32(1)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v221 = v215
	v223 = v214
	goto L52
L52:
	;
	if base.I32_extend8_s(v221) < int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v254 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)) = uint8(v254)
	v260 = v191 + int32(2)
	v262 = int32(0)
	goto L48
L54:
	;
	goto L53
L55:
	;
	goto L56
L56:
	;
	if base.B2i32(base.B2i32(base.Ui32(v221-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v221|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v221 != int32(95)) != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v250 = int32(1)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	if v253 != 0 {
		v221 = v253
		v223 = v223 + v250
		goto L52
	} else {
		goto L58
	}
L58:
	;
	v260 = v199
	v262 = v250
	goto L48
L59:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	if v283 != int32(34) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v301))) = uint8(v300)
	v303 = int32(1)
	v272 = v301 + v303
	v277 = v277 + v303
	goto L59
L62:
	;
	if v283 != 0 {
		v300 = v283
		v301 = v272
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v295 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v295)
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277))))
	v300 = v299
	v301 = v272 + int32(1)
	goto L61
L65:
	;
	if v262 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v290 = v272
	goto L68
L67:
	;
	v286 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v286)
	v290 = v272 + int32(1)
	goto L68
L68:
	;
	v291 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v291)
	F_ReleaseCatCache(m, v206)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	goto L43
L70:
	;
	goto L43
}
func F_aclnewowner(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int64
	_ = v130
	var v136 int32
	_ = v136
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v161 int32
	_ = v161
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	F_check_acl(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v21 == int32(0) {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v31 = (v24<<(uint(int32(3))%32) + int32(23)) & int32(-8)
		} else {
			v31 = v21
		}
		if int32(0) <= v20 {
			v35 = v20 << (uint(int32(4)) % 32)
			v37 = v35 + int32(24)
			v38 = F_palloc0(m, v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v38)+20)) = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = int32(1033)
				*(*int64)(unsafe.Add(mBase, uint32(v38)+4)) = int64(1)
				*(*int32)(unsafe.Add(mBase, uint32(v38))) = v37 << (uint(int32(2)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v20
				v51 = v38 + int32(24)
				if v35 != 0 {
					base.MemoryCopy(m, v51, l0+v31, v35)
				} else {
				}
				if v20 == int32(0) {
				} else {
					v57 = v51
					v61 = int32(0)
					v63 = int32(0)
					for {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
						if l1 == v68 {
							*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = l2
							v73 = v63
						} else {
							v73 = base.B2i32(l2 == v68) | v63
						}
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
						if l1 == v74 {
							*(*int32)(unsafe.Add(mBase, uint32(v57))) = l2
							v79 = v73
						} else {
							v79 = base.B2i32(l2 == v74) | v73
						}
						v83 = v61 + int32(1)
						if v83 != v20 {
							v57 = v57 + int32(16)
							v61 = v83
							v63 = v79
							continue
						} else {
							break
						}
						break
					}
					if v79&int32(1) == int32(0) {
					} else {
						v89 = int32(0)
						v92 = v51
						v94 = v89
						v95 = v89
						for {
							v103 = v95 + int32(1)
							v104 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
							if v104 != int64(0) {
								if v103 < v20 {
									v108 = v92
									v110 = v103
									for {
										v120 = v108 + int32(16)
										v121 = *(*int64)(unsafe.Add(mBase, uint32(v108)+24))
										if v121 == int64(0) {
										} else {
											v124 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
											v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
											if v124 != v125 {
											} else {
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
												v128 = *(*int32)(unsafe.Add(mBase, uint32(v108)+20))
												if v127 != v128 {
												} else {
													v130 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
													*(*int64)(unsafe.Add(mBase, uint32(v92)+8)) = v130 | v121
													*(*int64)(unsafe.Add(mBase, uint32(v108)+24)) = int64(0)
												}
											}
										}
										v136 = v110 + int32(1)
										if v136 != v20 {
											v108 = v120
											v110 = v136
											continue
										} else {
											break
										}
										break
									}
								} else {
								}
								v151 = v51 + v94<<(uint(int32(4))%32)
								v152 = *(*int64)(unsafe.Add(mBase, uint32(v92)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v151)+8)) = v152
								v154 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
								*(*int64)(unsafe.Add(mBase, uint32(v151))) = v154
								v161 = v94 + int32(1)
							} else {
								v161 = v94
							}
							if v103 != v20 {
								v92 = v92 + int32(16)
								v94 = v161
								v95 = v103
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v38)+16)) = v161
						*(*int32)(unsafe.Add(mBase, uint32(v38))) = v161<<(uint(int32(6))%32) + int32(96)
					}
				}
				m.G0 = v14 + int32(16)
				return v38
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v196 = m.ExcPending
			if v196 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v20
				F_errmsg_internal(m, int32(_a_F_aclnewowner_0), v14)
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_aclnewowner_1), int32(446), int32(_a_F_aclnewowner_2))
					mBase = m.M
					v205 = m.ExcPending
					if v205 != 0 {
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
	var v48 int32
	_ = v48
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
	v48 = int32(0)
	goto L14
L14:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v39+v48<<(uint(int32(2))%32))))
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
	v73 = v48 + int32(1)
	if v73 != v37 {
		v48 = v73
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
	var v15 int32
	_ = v15
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
	v15 = v3
	goto L4
L4:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v15<<(uint(int32(2))%32))))
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
	v52 = v15 + int32(1)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v52 < v53 {
		v15 = v52
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v149 int32
	_ = v149
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_add_parameter_name[0]))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	if v149 != 0 {
		goto L37
	} else {
		goto L38
	}
L2:
	;
	goto L1
L3:
	;
	v149 = v36
	goto L2
L5:
	;
	v149 = int32(0)
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
	if v43|base.B2i32(v38 == int32(1))&int32(0) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	goto L3
L15:
	;
	goto L16
L16:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 != 0 {
		v36 = v55
		v38 = v56
		goto L12
	} else {
		goto L18
	}
L18:
	;
	goto L13
L19:
	;
	goto L5
L37:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_add_parameter_name_0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	F_plpgsql_ns_additem(m, l0, l1, l2)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L40
	} else {
		goto L45
	}
L40:
	;
	return
L41:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = l2
	F_errmsg(m, int32(_a_F_add_parameter_name_1), v7)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_add_parameter_name_2), int32(934), int32(_a_F_add_parameter_name_3))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L40
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	m.G0 = v7 + int32(16)
	return
}
func F_add_size(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = l0 + l1
	if base.Ui32(v4) < base.Ui32(l0) {
		v6 = m.G0
		v8 = v6 - int32(16)
		m.G0 = v8
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
				F_errmsg(m, int32(_a_F_add_size_0), v8)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_add_size_1), int32(1755), int32(_a_F_add_size_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
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
		return v4
	}
}
func F_adjust_appendrel_attrs_mutator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
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
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
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
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
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
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
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
	var v330 int32
	_ = v330
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
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
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
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
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v481 int32
	_ = v481
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
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
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
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
	var v563 int32
	_ = v563
	var v572 int32
	_ = v572
	var v576 int64
	_ = v576
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v620 int64
	_ = v620
	var v622 int64
	_ = v622
	var v624 int64
	_ = v624
	var v626 int64
	_ = v626
	var v628 int64
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
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
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
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
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v698 int32
	_ = v698
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l0 == v3 {
		v732 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(32)
	return v732
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 <= int32(319) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L21
	} else {
		goto L244
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L21
	} else {
		goto L241
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L21
	} else {
		goto L238
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L21
	} else {
		goto L234
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L21
	} else {
		goto L230
	}
L8:
	;
	v651 = F_expression_tree_mutator_impl(m, l0, int32(902), l1)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L21
	} else {
		goto L229
	}
L9:
	;
	v616 = F_palloc0(m, int32(40))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L21
	} else {
		goto L221
	}
L10:
	;
	v597 = F_palloc0(m, int32(40))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L21
	} else {
		goto L217
	}
L11:
	;
	v312 = F_palloc0(m, int32(168))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L21
	} else {
		goto L119
	}
L12:
	;
	v260 = F_expression_tree_mutator_impl(m, l0, int32(902), l1)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L21
	} else {
		goto L98
	}
L13:
	;
	v237 = int32(0)
	goto L92
L14:
	;
	v39 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L21
	} else {
		goto L24
	}
L15:
	;
	switch v20 - int32(271) {
	case 0:
		goto L10
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L8
	case 9:
		goto L9
	default:
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	switch v20 - int32(320) {
	case 0:
		goto L11
	case 1:
		goto L12
	default:
		goto L8
	}
L18:
	;
	if v20 == int32(6) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	if v20 != int32(58) {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v29 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(0)
L22:
	;
	if v18 <= int32(0) {
		v732 = v29
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	goto L13
L24:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
	if v41 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v732 = v39
	goto L1
L26:
	;
	goto L27
L27:
	;
	if v18 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v732 = v39
	goto L1
L29:
	;
	goto L30
L30:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v46 = int32(0)
	goto L32
L31:
	;
	if v44 != int32(-4) {
		goto L72
	} else {
		goto L73
	}
L32:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v19+v46<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v61 != v44 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v67 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+40)) = uint16(v67)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v66
	v72 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+8)))
	if v67 < v72 {
		goto L38
	} else {
		goto L39
	}
L34:
	;
	v64 = v46 + int32(1)
	if v18 != v64 {
		v46 = v64
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	goto L31
L38:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	if v75 == int32(0) {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v72 != 0 {
		goto L54
	} else {
		goto L55
	}
L41:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 < v72 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v80+v72<<(uint(int32(2))%32)-int32(4))))
	v87 = F_copyObjectImpl(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L21
	} else {
		goto L43
	}
L43:
	;
	if v87 == int32(0) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v92 == int32(6) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+32)) = v91
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v87)+24))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v98 = F_bms_add_members(m, v96, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L21
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v91 != 0 {
		goto L5
	} else {
		goto L49
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87)+24)) = v98
	v732 = v87
	goto L1
L49:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v101 == int32(0) {
		v732 = v87
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_0), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L21
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(307), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	v732 = v39
	goto L1
L55:
	;
	goto L56
L56:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v117 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	if v117 == v118 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v135+v136<<(uint(int32(2))%32)-int32(4))))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	v144 = F_copyObjectImpl(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L21
	} else {
		goto L64
	}
L60:
	;
	v732 = v39
	goto L1
L61:
	;
	goto L62
L62:
	;
	v121 = F_palloc0(m, int32(20))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L21
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v121))) = int32(30)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v121)+12)) = int64(-4294967294)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v126
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v130
	v732 = v121
	goto L1
L64:
	;
	v147 = F_palloc0(m, int32(24))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L21
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+4)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = int32(36)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+8)) = v152
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v142)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+8))
	v158 = F_copyObjectImpl(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L21
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+20)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+16)) = v158
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	if v163 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v164 == int32(0) {
		v732 = v147
		goto L1
	} else {
		goto L68
	}
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L21
	} else {
		goto L69
	}
L69:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_0), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L21
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(365), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L21
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L72:
	;
	v732 = v39
	goto L1
L73:
	;
	goto L74
L74:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+132))
	v185 = int32(0)
	v189 = v3
	goto L75
L75:
	;
	v198 = v19 + v185<<(uint(int32(2))%32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+8))
	v201 = F_bms_is_member(m, v200, v183)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L21
	} else {
		goto L77
	}
L76:
	;
	if v205 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	if v201 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	if v189 != 0 {
		goto L3
	} else {
		goto L81
	}
L79:
	;
	v205 = v189
	goto L80
L80:
	;
	v207 = v185 + int32(1)
	if v207 != v18 {
		v185 = v207
		v189 = v205
		goto L75
	} else {
		goto L82
	}
L81:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	v205 = v204
	goto L80
L82:
	;
	goto L76
L83:
	;
	v732 = v39
	goto L1
L84:
	;
	goto L85
L85:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+140))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+12))
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+8)))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v213+v214<<(uint(int32(2))%32)-int32(4))))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	v222 = F_bms_is_member(m, v205, v221)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L21
	} else {
		goto L86
	}
L86:
	;
	if v222 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v225 = F_copyObjectImpl(m, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L21
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v235 = F_makeNullConst(m, v232, v233, v234)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L21
	} else {
		goto L91
	}
L90:
	;
	v227 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v225)+40)) = uint16(v227)
	*(*int32)(unsafe.Add(mBase, uint32(v225)+36)) = v227
	*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v205
	v732 = v225
	goto L1
L91:
	;
	v732 = v235
	goto L1
L92:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v19+v237<<(uint(int32(2))%32))))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v251)+4))
	if v252 != v35 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v257
	v732 = v29
	goto L1
L94:
	;
	v255 = v237 + int32(1)
	if v18 != v255 {
		v237 = v255
		goto L92
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	goto L93
L97:
	;
	v732 = v29
	goto L1
L98:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
	if v262 != 0 {
		v732 = v260
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+8))
	if int32(0) < v18 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v267 = int32(0)
	v271 = v3
	goto L103
L101:
	;
	v302 = v3
	goto L102
L102:
	;
	if v302 != 0 {
		goto L116
	} else {
		goto L117
	}
L103:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v19+v267<<(uint(int32(2))%32))))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v283 = F_bms_is_member(m, v282, v263)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L21
	} else {
		goto L105
	}
L104:
	;
	v302 = v294
	goto L102
L105:
	;
	if v283 != 0 {
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
	v294 = v271
	goto L108
L108:
	;
	v296 = v267 + int32(1)
	if v296 != v18 {
		v267 = v296
		v271 = v294
		goto L103
	} else {
		goto L115
	}
L109:
	;
	v287 = v271
	goto L111
L110:
	;
	v285 = F_bms_copy(m, v263)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L21
	} else {
		goto L112
	}
L111:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v281)+4))
	v289 = F_bms_del_member(m, v287, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L21
	} else {
		goto L113
	}
L112:
	;
	v287 = v285
	goto L111
L113:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v281)+8))
	v292 = F_bms_add_member(m, v289, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L21
	} else {
		goto L114
	}
L114:
	;
	v294 = v292
	goto L108
L115:
	;
	goto L104
L116:
	;
	v309 = v302
	goto L118
L117:
	;
	v309 = v263
	goto L118
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260)+8)) = v309
	v732 = v260
	goto L1
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312))) = int32(320)
	base.MemoryCopy(m, v312, l0, int32(168))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v319 = F_adjust_appendrel_attrs_mutator(m, v318, l1)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L21
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+4)) = v319
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v323 = F_adjust_appendrel_attrs_mutator(m, v322, l1)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L21
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+52)) = v323
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v327 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v335 = v3
	v339 = v3
	goto L125
L123:
	;
	v366 = v3
	goto L124
L124:
	;
	if v366 != 0 {
		goto L138
	} else {
		goto L139
	}
L125:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v330+v339<<(uint(int32(2))%32))))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v347 = F_bms_is_member(m, v346, v326)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L21
	} else {
		goto L127
	}
L126:
	;
	v366 = v358
	goto L124
L127:
	;
	if v347 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if v335 != 0 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	v358 = v335
	goto L130
L130:
	;
	v360 = v339 + int32(1)
	if v360 != v327 {
		v335 = v358
		v339 = v360
		goto L125
	} else {
		goto L137
	}
L131:
	;
	v351 = v335
	goto L133
L132:
	;
	v349 = F_bms_copy(m, v326)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L21
	} else {
		goto L134
	}
L133:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	v353 = F_bms_del_member(m, v351, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L21
	} else {
		goto L135
	}
L134:
	;
	v351 = v349
	goto L133
L135:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v345)+8))
	v356 = F_bms_add_member(m, v353, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L21
	} else {
		goto L136
	}
L136:
	;
	v358 = v356
	goto L130
L137:
	;
	goto L126
L138:
	;
	v373 = v366
	goto L140
L139:
	;
	v373 = v326
	goto L140
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+28)) = v373
	v375 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v376 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v389 = int32(0)
	v390 = v3
	goto L144
L142:
	;
	v421 = v3
	goto L143
L143:
	;
	if v421 != 0 {
		goto L157
	} else {
		goto L158
	}
L144:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v379+v389<<(uint(int32(2))%32))))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v397 = F_bms_is_member(m, v396, v375)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L21
	} else {
		goto L146
	}
L145:
	;
	v421 = v408
	goto L143
L146:
	;
	if v397 != 0 {
		goto L147
	} else {
		goto L148
	}
L147:
	;
	if v390 != 0 {
		goto L150
	} else {
		goto L151
	}
L148:
	;
	v408 = v390
	goto L149
L149:
	;
	v410 = v389 + int32(1)
	if v410 != v376 {
		v389 = v410
		v390 = v408
		goto L144
	} else {
		goto L156
	}
L150:
	;
	v401 = v390
	goto L152
L151:
	;
	v399 = F_bms_copy(m, v375)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L21
	} else {
		goto L153
	}
L152:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	v403 = F_bms_del_member(m, v401, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L21
	} else {
		goto L154
	}
L153:
	;
	v401 = v399
	goto L152
L154:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v395)+8))
	v406 = F_bms_add_member(m, v403, v405)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L21
	} else {
		goto L155
	}
L155:
	;
	v408 = v406
	goto L149
L156:
	;
	goto L145
L157:
	;
	v423 = v421
	goto L159
L158:
	;
	v423 = v375
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+32)) = v423
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v426 = int32(0)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v426 < v428 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v437 = v426
	v441 = int32(0)
	goto L163
L161:
	;
	v468 = v426
	goto L162
L162:
	;
	if v468 != 0 {
		goto L176
	} else {
		goto L177
	}
L163:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v431+v441<<(uint(int32(2))%32))))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	v449 = F_bms_is_member(m, v448, v425)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L21
	} else {
		goto L165
	}
L164:
	;
	v468 = v460
	goto L162
L165:
	;
	if v449 != 0 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	if v437 != 0 {
		goto L169
	} else {
		goto L170
	}
L167:
	;
	v460 = v437
	goto L168
L168:
	;
	v462 = v441 + int32(1)
	if v462 != v428 {
		v437 = v460
		v441 = v462
		goto L163
	} else {
		goto L175
	}
L169:
	;
	v453 = v437
	goto L171
L170:
	;
	v451 = F_bms_copy(m, v425)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L21
	} else {
		goto L172
	}
L171:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v447)+4))
	v455 = F_bms_del_member(m, v453, v454)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L21
	} else {
		goto L173
	}
L172:
	;
	v453 = v451
	goto L171
L173:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v447)+8))
	v458 = F_bms_add_member(m, v455, v457)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L21
	} else {
		goto L174
	}
L174:
	;
	v460 = v458
	goto L168
L175:
	;
	goto L164
L176:
	;
	v475 = v468
	goto L178
L177:
	;
	v475 = v425
	goto L178
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+40)) = v475
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v478 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v491 = int32(0)
	v492 = v426
	goto L182
L180:
	;
	v523 = v426
	goto L181
L181:
	;
	if v523 != 0 {
		goto L195
	} else {
		goto L196
	}
L182:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v481+v491<<(uint(int32(2))%32))))
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v499 = F_bms_is_member(m, v498, v477)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L21
	} else {
		goto L184
	}
L183:
	;
	v523 = v510
	goto L181
L184:
	;
	if v499 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	if v492 != 0 {
		goto L188
	} else {
		goto L189
	}
L186:
	;
	v510 = v492
	goto L187
L187:
	;
	v512 = v491 + int32(1)
	if v512 != v478 {
		v491 = v512
		v492 = v510
		goto L182
	} else {
		goto L194
	}
L188:
	;
	v503 = v492
	goto L190
L189:
	;
	v501 = F_bms_copy(m, v477)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L21
	} else {
		goto L191
	}
L190:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v505 = F_bms_del_member(m, v503, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L21
	} else {
		goto L192
	}
L191:
	;
	v503 = v501
	goto L190
L192:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v497)+8))
	v508 = F_bms_add_member(m, v505, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L21
	} else {
		goto L193
	}
L193:
	;
	v510 = v508
	goto L187
L194:
	;
	goto L183
L195:
	;
	v525 = v523
	goto L197
L196:
	;
	v525 = v477
	goto L197
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+44)) = v525
	v527 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v528 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v528 < v529 {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v541 = v528
	v542 = int32(0)
	goto L201
L199:
	;
	v572 = v528
	goto L200
L200:
	;
	v576 = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v312)+152)) = v576
	*(*int64)(unsafe.Add(mBase, uint32(v312)+144)) = v576
	*(*int64)(unsafe.Add(mBase, uint32(v312)+136)) = v576
	*(*int64)(unsafe.Add(mBase, uint32(v312)+128)) = v576
	*(*int32)(unsafe.Add(mBase, uint32(v312)+116)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v312)+108)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v312)+88)) = v576
	*(*int64)(unsafe.Add(mBase, uint32(v312)+80)) = v576
	*(*int64)(unsafe.Add(mBase, uint32(v312)+64)) = v576
	if v572 != 0 {
		goto L214
	} else {
		goto L215
	}
L201:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v532+v542<<(uint(int32(2))%32))))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	v550 = F_bms_is_member(m, v549, v527)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L21
	} else {
		goto L203
	}
L202:
	;
	v572 = v561
	goto L200
L203:
	;
	if v550 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	if v541 != 0 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	v561 = v541
	goto L206
L206:
	;
	v563 = v542 + int32(1)
	if v563 != v529 {
		v541 = v561
		v542 = v563
		goto L201
	} else {
		goto L213
	}
L207:
	;
	v554 = v541
	goto L209
L208:
	;
	v552 = F_bms_copy(m, v527)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L21
	} else {
		goto L210
	}
L209:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v548)+4))
	v556 = F_bms_del_member(m, v554, v555)
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L21
	} else {
		goto L211
	}
L210:
	;
	v554 = v552
	goto L209
L211:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v548)+8))
	v559 = F_bms_add_member(m, v556, v558)
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L21
	} else {
		goto L212
	}
L212:
	;
	v561 = v559
	goto L206
L213:
	;
	goto L202
L214:
	;
	v594 = v572
	goto L216
L215:
	;
	v594 = v527
	goto L216
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+48)) = v594
	v732 = v312
	goto L1
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597))) = int32(271)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v602 = F_adjust_appendrel_attrs_mutator(m, v601, l1)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L21
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597)+4)) = v602
	v605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v606 = F_adjust_appendrel_attrs_mutator(m, v605, l1)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L21
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597)+8)) = v606
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v597)+12)) = v609
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v612 = F_adjust_appendrel_attrs_mutator(m, v611, l1)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L21
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597)+16)) = v612
	v732 = v597
	goto L1
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616))) = int32(280)
	v620 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v616)+8)) = v620
	v622 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v616)+16)) = v622
	v624 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v616)+24)) = v624
	v626 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v616)+32)) = v626
	v628 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v616))) = v628
	v630 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v631 = F_adjust_appendrel_attrs_mutator(m, v630, l1)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L21
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616)+4)) = v631
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v634 == int32(0) {
		v732 = v616
		goto L1
	} else {
		goto L223
	}
L223:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v637 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+4))
	v642 = v638 << (uint(int32(2)) % 32)
	goto L226
L225:
	;
	v642 = int32(0)
	goto L226
L226:
	;
	v643 = F_palloc(m, v642)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L21
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v616)+8)) = v643
	if v642 == int32(0) {
		v732 = v616
		goto L1
	} else {
		goto L228
	}
L228:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	base.MemoryCopy(m, v643, v648, v642)
	v732 = v616
	goto L1
L229:
	;
	v732 = v651
	goto L1
L230:
	;
	v657 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+8)))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	v659 = F_get_rel_name(m, v658)
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L21
	} else {
		goto L231
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v659
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v657
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_3), v14)
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L21
	} else {
		goto L232
	}
L232:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(288), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L21
	} else {
		goto L233
	}
L233:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L234:
	;
	v675 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+8)))
	v676 = *(*int32)(unsafe.Add(mBase, uint32(v60)+32))
	v677 = F_get_rel_name(m, v676)
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L21
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v677
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v675
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_3), v14+int32(16))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L21
	} else {
		goto L236
	}
L236:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(293), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L21
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
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_4), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L21
	} else {
		goto L239
	}
L239:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(305), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L21
	} else {
		goto L240
	}
L240:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L241:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_4), int32(0))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L21
	} else {
		goto L242
	}
L242:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(363), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L21
	} else {
		goto L243
	}
L243:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L244:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_5), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L21
	} else {
		goto L245
	}
L245:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(391), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L21
	} else {
		goto L246
	}
L246:
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
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
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
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
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L32
	}
L4:
	;
	m.G0 = v11 + int32(16)
	return v77
L5:
	;
	v20 = F_adjust_inherited_attnums_multilevel(m, l0, l1, v18, l3)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v24 = l1
	goto L7
L7:
	;
	v25 = int32(0)
	v27 = m.G0
	v29 = v27 - int32(32)
	m.G0 = v29
	if v24 == v25 {
		v77 = v25
		goto L12
	} else {
		goto L13
	}
L8:
	;
	return int32(0)
L9:
	;
	v24 = v20
	goto L7
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L28
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L24
	}
L12:
	;
	m.G0 = v29 + int32(32)
	goto L4
L13:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v33 <= int32(0) {
		v77 = v25
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v36 = v25
	v38 = v25
	goto L15
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v44+v36<<(uint(int32(2))%32)))))
	if v48 <= int32(0) {
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v77 = v69
	goto L12
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
	if v51 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v54 < v48 {
		goto L11
	} else {
		goto L19
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+v48<<(uint(int32(2))%32)-int32(4))))
	if v62 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v65 != int32(6) {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62)+8)))
	v69 = F_lappend_int(m, v38, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v72 = v36 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	if v72 < v73 {
		v36 = v72
		v38 = v69
		goto L15
	} else {
		goto L23
	}
L23:
	;
	goto L16
L24:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v92 = F_get_rel_name(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+4)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v48
	F_errmsg_internal(m, int32(_a_F_adjust_inherited_attnums_multilevel_0), v29)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_adjust_inherited_attnums_multilevel_1), int32(722), int32(_a_F_adjust_inherited_attnums_multilevel_2))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v109 = F_get_rel_name(m, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+20)) = v109
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v48
	F_errmsg_internal(m, int32(_a_F_adjust_inherited_attnums_multilevel_0), v29+int32(16))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_adjust_inherited_attnums_multilevel_1), int32(726), int32(_a_F_adjust_inherited_attnums_multilevel_2))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l2
	F_errmsg_internal(m, int32(_a_F_adjust_inherited_attnums_multilevel_3), v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_adjust_inherited_attnums_multilevel_1), int32(744), int32(_a_F_adjust_inherited_attnums_multilevel_4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
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
	var v28 int32
	_ = v28
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
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
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v109 == int32(0) {
		goto L1
	} else {
		goto L36
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
	v28 = int32(0)
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v38 = v35 + v28<<(uint(int32(2))%32)
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
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v70+v44<<(uint(int32(2))%32))))
	if v88 != 0 {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v74 == v39 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v71 = v39
	goto L17
L19:
	;
	goto L20
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.B2i32(v62 == int32(0))|base.B2i32(v67 <= v44) != 0 {
		v71 = v45
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v70 != 0 {
		goto L16
	} else {
		goto L22
	}
L22:
	;
	v71 = v45
	goto L17
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+56)) = v71
	goto L25
L24:
	;
	goto L25
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	if v77 == v39 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+60)) = v71
	goto L28
L27:
	;
	goto L28
L28:
	;
	v81 = v28 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v81 < v82 {
		v28 = v81
		goto L9
	} else {
		goto L29
	}
L29:
	;
	goto L6
L30:
	;
	v89 = F_create_set_projection_path(m, l0, l1, v45, v84)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v93 = F_apply_projection_to_path(m, l0, l1, v45, v84)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L33
	} else {
		goto L35
	}
L33:
	;
	return
L34:
	;
	v44 = v44 + int32(1)
	v45 = v89
	goto L11
L35:
	;
	v44 = v44 + int32(1)
	v45 = v93
	goto L11
L36:
	;
	v112 = int32(0)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v113 <= v112 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v123 = v112
	goto L38
L38:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v132 = v129 + v123<<(uint(int32(2))%32)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v138 = int32(0)
	v139 = v133
	goto L40
L40:
	;
	v146 = int32(0)
	if l2 == v146 {
		v156 = v146
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if l3 != 0 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v150 <= v138 {
		v156 = int32(0)
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v156 = v152 + v138<<(uint(int32(2))%32)
	goto L42
L45:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v164+v138<<(uint(int32(2))%32))))
	if v177 != 0 {
		goto L54
	} else {
		goto L55
	}
L46:
	;
	v157 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if base.B2i32(v156 == v157)|base.B2i32(v159 <= v138) == v157 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v167 = v133
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v167
	v170 = v123 + int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v170 < v171 {
		v123 = v170
		goto L38
	} else {
		goto L53
	}
L49:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v164 != 0 {
		goto L45
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v167 = v139
	goto L48
L52:
	;
	goto L51
L53:
	;
	goto L1
L54:
	;
	v178 = F_create_set_projection_path(m, l0, l1, v139, v173)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L33
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v182 = F_create_projection_path(m, l0, l1, v139, v173)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L33
	} else {
		goto L58
	}
L57:
	;
	v138 = v138 + int32(1)
	v139 = v178
	goto L40
L58:
	;
	v138 = v138 + int32(1)
	v139 = v182
	goto L40
}
func F_advance_transition_function(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int64
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	v9 = int32(1)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+224))
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)))
	if v11 == v9 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0])) = v103
	goto L1
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v66 = int32(_a_F_advance_transition_function_0)
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0])) = v70
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = l1
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v73
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v76 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)) = uint8(v76)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+32)) = uint8(v75)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = m.T0[v80].(func(*base.Module, int32) int64)(m, v10)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L16
	} else {
		goto L19
	}
L6:
	;
	v21 = v9
	goto L9
L7:
	;
	goto L8
L8:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)))
	if v40 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v21<<(uint(int32(4))%32))+32)))
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
	if v30 <= v14 {
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
	v43 = int32(_a_F_advance_transition_function_0)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0])) = v47
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+191)))
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+188)))
	v52 = F_datumCopy(m, v49, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	if v57 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	v54 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)) = uint16(v54)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v52
	v103 = v44
	goto L2
L18:
	;
	goto L5
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+191)))
	if v85 != 0 {
		v94 = v81
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v94
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v97)
	v103 = v67
	goto L2
L21:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	if base.I32_wrap_i64(v81) == base.I32_wrap_i64(v87) {
		v94 = v81
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
	v92 = F_ExecAggCopyTransValue(m, l0, l1, v81, v90, v87, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v94 = v92
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
				F_errmsg(m, int32(_a_F_alen_scalar_0), int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_alen_scalar_1), int32(1922), int32(_a_F_alen_scalar_2))
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
func F_anycompatible_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anycompatible_out_0), int32(376), int32(_a_F_anycompatible_out_1), int32(_a_F_anycompatible_out_2), int32(_a_F_anycompatible_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anyelement_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anyelement_out_0), int32(374), int32(_a_F_anyelement_out_1), int32(_a_F_anyelement_out_2), int32(_a_F_anyelement_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_anyenum_in(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_anyenum_in_0), int32(194), int32(_a_F_anyenum_in_1), int32(_a_F_anyenum_in_2), int32(_a_F_anyenum_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_arrayexpr_startup_fn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v6 = F_palloc(m, int32(40))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v6
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(17)
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v11
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v14 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v14
		*(*uint8)(unsafe.Add(mBase, uint32(v6)+16)) = uint8(v14)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v13
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v21
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v24 = F_list_copy(m, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v24
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v30
			if v30 != 0 {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
				v34 = v32
			} else {
				v34 = int32(0)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = v34
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
	*(*int32)(unsafe.Add(mBase, _c_F_assign_datestyle[0])) = v4
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_datestyle[1])) = v7
	return
}
func F_assign_restrict_nonsystem_relation_kind(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, _c_F_assign_restrict_nonsystem_relation_kind[0])) = v4
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
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_facility[0]))
	if l0 != v4 {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_assign_syslog_facility[1])))
		if v7 != 0 {
			v9 = m.G0
			v10 = int32(16)
			v11 = v9 - v10
			m.G0 = v11
			v13 = int32(_a_F_assign_syslog_facility_0)
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_facility[2]))
			v15 = F_close(m, v14)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_facility[2])) = int32(-1)
			m.G0 = v11 + v10
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, _c_F_assign_syslog_facility[1])) = uint8(v24)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, _c_F_assign_syslog_facility[0])) = l0
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
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
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
	return v100
L2:
	;
	v100 = v13 + int32(1)
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
	v22 = v15 + v16<<(uint(int32(3))%32) + v13*int32(100)
	v24 = v22 + int32(32)
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
	if v39 == int32(0) {
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
	v32 = int32(0)
	if base.B2i32(v24 == v32)|base.B2i32(l1 == v32) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v39 = v31
	goto L17
L16:
	;
	v38 = F_strncmp(m, v24, l1, int32(64))
	mBase = m.M
	v39 = v38
	goto L17
L17:
	;
	goto L8
L18:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+119)))
	if v42 != int32(1) {
		goto L2
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v46 = v13 + int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(v47)+120)))
	if v46 < v48 {
		v13 = v46
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
	v58 = F_strcmp(m, int32(_a_F_attnameAttNum_0), l1)
	mBase = m.M
	if v58 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v87 == int32(0) {
		goto L23
	} else {
		goto L44
	}
L26:
	;
	v87 = int32(_a_F_attnameAttNum_1)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v63 = F_strcmp(m, int32(_a_F_attnameAttNum_2), l1)
	mBase = m.M
	if v63 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v87 = int32(_a_F_attnameAttNum_3)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v68 = F_strcmp(m, int32(_a_F_attnameAttNum_4), l1)
	mBase = m.M
	if v68 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v87 = int32(_a_F_attnameAttNum_5)
	goto L25
L33:
	;
	goto L34
L34:
	;
	v73 = F_strcmp(m, int32(_a_F_attnameAttNum_6), l1)
	mBase = m.M
	if v73 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v87 = int32(_a_F_attnameAttNum_7)
	goto L25
L36:
	;
	goto L37
L37:
	;
	v78 = F_strcmp(m, int32(_a_F_attnameAttNum_8), l1)
	mBase = m.M
	if v78 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v87 = int32(_a_F_attnameAttNum_9)
	goto L25
L39:
	;
	goto L40
L40:
	;
	v85 = F_strcmp(m, int32(_a_F_attnameAttNum_10), l1)
	mBase = m.M
	if v85 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v86 = int32(0)
	goto L43
L42:
	;
	v86 = int32(_a_F_attnameAttNum_11)
	goto L43
L43:
	;
	v87 = v86
	goto L25
L44:
	;
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+74)))
	if v90 != 0 {
		v100 = v90
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L23
}
func F_autoprewarm_dump_now(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(192)
	m.G0 = v8
	v12 = v2
	v13 = v2
	v14 = int32(-1)
	v15 = v2
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v14 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v95 = int32(m.ExcTag)
	v96 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v95 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v13
	v23 = F_GetNamedDSMSegment(m, v8+int32(183))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v44 = v12
	v45 = v13
	v46 = v15
	goto L8
L8:
	;
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[0])) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v13
	F_before_shmem_exit(m, int32(_a_F_autoprewarm_dump_now_0), int64(0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[1]))
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[2]))
	goto L11
L11:
	;
	v37 = v8 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v8 + int32(12)
	goto L14
L12:
	;
	v44 = v33
	v45 = v35
	v46 = int32(0)
	goto L8
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v44
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[2])) = v8 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v45
	v57 = F_apw_dump_now(m, int32(0), int32(1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[1])) = v44
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[2])) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v45
	F_cancel_before_shmem_exit(m, int32(_a_F_autoprewarm_dump_now_0), int64(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v44
	F_cancel_before_shmem_exit(m, int32(_a_F_autoprewarm_dump_now_0), int64(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[1])) = v44
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[2])) = v45
	m.G0 = v8 + int32(192)
	return base.I64_extend_i32_s(v57)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v44
	F_apw_detach_shmem(m, v8, int64(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v44
	F_pg_re_throw(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	goto L4
L23:
	;
	v100 = int32(v96)
	m.G0 = v8
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v8+int32(12) == v106 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	m.ExcPending = 1
	goto L32
L25:
	;
	if v110 != 0 {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v110 = v108
	goto L28
L27:
	;
	v110 = int32(0)
	goto L28
L28:
	;
	goto L25
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+184))
	v12 = v112
	v13 = v111
	v14 = v110
	v15 = v102
	goto L1
L30:
	;
	goto L31
L31:
	;
	F___wasm_longjmp(m, v103, v102)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int64(0)
L33:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
