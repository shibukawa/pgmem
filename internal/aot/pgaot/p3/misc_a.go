package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_AccessTempTableNamespace(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	v5 = m.G0
	v7 = v5 - int32(128)
	m.G0 = v7
	v9 = int32(_a_F_AccessTempTableNamespace_0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[0])) = v11 | int32(1)
	if l0 == int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1]))
		if v18 != 0 {
			m.G0 = v7 + int32(128)
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[2]))
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[3]))
			v25 = F_object_aclcheck(m, int32(1262), v21, v23, int64(1024))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v132 = m.ExcPending
					if v132 != 0 {
						return
					} else {
						F_errcode(m, int32(16797828))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
							return
						} else {
							v137 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[2]))
							v138 = F_get_database_name(m, v137)
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v138
								F_errmsg(m, int32(_a_F_AccessTempTableNamespace_1), v7+int32(32))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_3), int32(_a_F_AccessTempTableNamespace_4))
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
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
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[4])))
					if v29 == int32(1) {
						v34 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[5]))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+316))
						v37 = base.B2i32(v35 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[4])) = uint8(v37)
						v39 = v37
					} else {
						v39 = int32(0)
					}
					if v39 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v154 = m.ExcPending
						if v154 != 0 {
							return
						} else {
							F_errcode(m, int32(100663618))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_AccessTempTableNamespace_5), int32(0))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_6), int32(_a_F_AccessTempTableNamespace_4))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
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
						v41 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[6]))
						if int32(0) <= v41 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v170 = m.ExcPending
							if v170 != 0 {
								return
							} else {
								F_errcode(m, int32(100663618))
								mBase = m.M
								v173 = m.ExcPending
								if v173 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_AccessTempTableNamespace_7), int32(0))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_8), int32(_a_F_AccessTempTableNamespace_4))
										mBase = m.M
										v182 = m.ExcPending
										if v182 != 0 {
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
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v45
							v48 = v7 + int32(48)
							v53 = F_pg_snprintf(m, v48, int32(64), int32(_a_F_AccessTempTableNamespace_9), v7+int32(16))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v56 = int32(0)
								v59 = F_GetSysCacheOid(m, int32(37), v48, v56, v56, v56)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return
								} else {
									if v59 == int32(0) {
										v65 = F_NamespaceCreate(m, v48, int32(10), int32(1))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											F_CommandCounterIncrement(m)
											mBase = m.M
											v68 = m.ExcPending
											if v68 != 0 {
												return
											} else {
												v80 = v65
												v82 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
												*(*int32)(unsafe.Add(mBase, uint32(v7))) = v82
												v85 = v7 + int32(48)
												v88 = F_pg_snprintf(m, v85, int32(64), int32(_a_F_AccessTempTableNamespace_10), v7)
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return
												} else {
													v91 = int32(0)
													v94 = F_GetSysCacheOid(m, int32(37), v85, v91, v91, v91)
													mBase = m.M
													v95 = m.ExcPending
													if v95 != 0 {
														return
													} else {
														if v94 == int32(0) {
															v100 = F_NamespaceCreate(m, v85, int32(10), int32(1))
															mBase = m.M
															v101 = m.ExcPending
															if v101 != 0 {
																return
															} else {
																F_CommandCounterIncrement(m)
																mBase = m.M
																v103 = m.ExcPending
																if v103 != 0 {
																	return
																} else {
																	v104 = v100
																	*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v104
																	*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v80
																	v110 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
																	*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = v80
																	v113 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
																	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
																	v116 = int32(1)
																	*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v116)
																	*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v114
																	v121 = int32(0)
																	*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v121)
																	m.G0 = v7 + int32(128)
																	return
																}
															}
														} else {
															v104 = v94
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v104
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v80
															v110 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
															*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = v80
															v113 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
															v116 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v116)
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v114
															v121 = int32(0)
															*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v121)
															m.G0 = v7 + int32(128)
															return
														}
													}
												}
											}
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+124)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+120)) = v59
										*(*int32)(unsafe.Add(mBase, uint32(v7)+116)) = int32(2615)
										F_performDeletion(m, v7+int32(116), int32(1), int32(29))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return
										} else {
											v80 = v59
											v82 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v82
											v85 = v7 + int32(48)
											v88 = F_pg_snprintf(m, v85, int32(64), int32(_a_F_AccessTempTableNamespace_10), v7)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return
											} else {
												v91 = int32(0)
												v94 = F_GetSysCacheOid(m, int32(37), v85, v91, v91, v91)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													if v94 == int32(0) {
														v100 = F_NamespaceCreate(m, v85, int32(10), int32(1))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																v104 = v100
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v104
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v80
																v110 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
																*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = v80
																v113 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
																v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
																v116 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v116)
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v114
																v121 = int32(0)
																*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v121)
																m.G0 = v7 + int32(128)
																return
															}
														}
													} else {
														v104 = v94
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v104
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v80
														v110 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
														*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = v80
														v113 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
														v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
														v116 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v116)
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v114
														v121 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v121)
														m.G0 = v7 + int32(128)
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
		v21 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[2]))
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[3]))
		v25 = F_object_aclcheck(m, int32(1262), v21, v23, int64(1024))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			if v25 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
					return
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return
					} else {
						v137 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[2]))
						v138 = F_get_database_name(m, v137)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v138
							F_errmsg(m, int32(_a_F_AccessTempTableNamespace_1), v7+int32(32))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_3), int32(_a_F_AccessTempTableNamespace_4))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
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
				v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[4])))
				if v29 == int32(1) {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[5]))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+316))
					v37 = base.B2i32(v35 != int32(2))
					*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[4])) = uint8(v37)
					v39 = v37
				} else {
					v39 = int32(0)
				}
				if v39 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v154 = m.ExcPending
					if v154 != 0 {
						return
					} else {
						F_errcode(m, int32(100663618))
						mBase = m.M
						v157 = m.ExcPending
						if v157 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_AccessTempTableNamespace_5), int32(0))
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_6), int32(_a_F_AccessTempTableNamespace_4))
								mBase = m.M
								v166 = m.ExcPending
								if v166 != 0 {
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
					v41 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[6]))
					if int32(0) <= v41 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return
						} else {
							F_errcode(m, int32(100663618))
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_AccessTempTableNamespace_7), int32(0))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_AccessTempTableNamespace_2), int32(_a_F_AccessTempTableNamespace_8), int32(_a_F_AccessTempTableNamespace_4))
									mBase = m.M
									v182 = m.ExcPending
									if v182 != 0 {
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
						v45 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v45
						v48 = v7 + int32(48)
						v53 = F_pg_snprintf(m, v48, int32(64), int32(_a_F_AccessTempTableNamespace_9), v7+int32(16))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return
						} else {
							v56 = int32(0)
							v59 = F_GetSysCacheOid(m, int32(37), v48, v56, v56, v56)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								if v59 == int32(0) {
									v65 = F_NamespaceCreate(m, v48, int32(10), int32(1))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
										return
									} else {
										F_CommandCounterIncrement(m)
										mBase = m.M
										v68 = m.ExcPending
										if v68 != 0 {
											return
										} else {
											v80 = v65
											v82 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v82
											v85 = v7 + int32(48)
											v88 = F_pg_snprintf(m, v85, int32(64), int32(_a_F_AccessTempTableNamespace_10), v7)
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return
											} else {
												v91 = int32(0)
												v94 = F_GetSysCacheOid(m, int32(37), v85, v91, v91, v91)
												mBase = m.M
												v95 = m.ExcPending
												if v95 != 0 {
													return
												} else {
													if v94 == int32(0) {
														v100 = F_NamespaceCreate(m, v85, int32(10), int32(1))
														mBase = m.M
														v101 = m.ExcPending
														if v101 != 0 {
															return
														} else {
															F_CommandCounterIncrement(m)
															mBase = m.M
															v103 = m.ExcPending
															if v103 != 0 {
																return
															} else {
																v104 = v100
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v104
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v80
																v110 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
																*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = v80
																v113 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
																v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
																v116 = int32(1)
																*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v116)
																*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v114
																v121 = int32(0)
																*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v121)
																m.G0 = v7 + int32(128)
																return
															}
														}
													} else {
														v104 = v94
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v104
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v80
														v110 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
														*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = v80
														v113 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
														v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
														v116 = int32(1)
														*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v116)
														*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v114
														v121 = int32(0)
														*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v121)
														m.G0 = v7 + int32(128)
														return
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+124)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v7)+120)) = v59
									*(*int32)(unsafe.Add(mBase, uint32(v7)+116)) = int32(2615)
									F_performDeletion(m, v7+int32(116), int32(1), int32(29))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return
									} else {
										v80 = v59
										v82 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[7]))
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v82
										v85 = v7 + int32(48)
										v88 = F_pg_snprintf(m, v85, int32(64), int32(_a_F_AccessTempTableNamespace_10), v7)
										mBase = m.M
										v89 = m.ExcPending
										if v89 != 0 {
											return
										} else {
											v91 = int32(0)
											v94 = F_GetSysCacheOid(m, int32(37), v85, v91, v91, v91)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return
											} else {
												if v94 == int32(0) {
													v100 = F_NamespaceCreate(m, v85, int32(10), int32(1))
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return
													} else {
														F_CommandCounterIncrement(m)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return
														} else {
															v104 = v100
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v104
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v80
															v110 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
															*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = v80
															v113 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
															v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
															v116 = int32(1)
															*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v116)
															*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v114
															v121 = int32(0)
															*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v121)
															m.G0 = v7 + int32(128)
															return
														}
													}
												} else {
													v104 = v94
													*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[8])) = v104
													*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[1])) = v80
													v110 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[9]))
													*(*int32)(unsafe.Add(mBase, uint32(v110)+68)) = v80
													v113 = *(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[10]))
													v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
													v116 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[11])) = uint8(v116)
													*(*int32)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[12])) = v114
													v121 = int32(0)
													*(*uint8)(unsafe.Add(mBase, _c_F_AccessTempTableNamespace[13])) = uint8(v121)
													m.G0 = v7 + int32(128)
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
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceOldestClogXid[0]))
	v7 = F_LWLockAcquire(m, v3+int32(_a_F_AdvanceOldestClogXid_0), int32(0))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceOldestClogXid[1]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
		if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v11)) == int32(0) {
			v23 = base.B2i32(base.Ui32(v11) < base.Ui32(l0))
		} else {
			v23 = int32(base.Ui32(v11-l0) >> (uint(int32(31)) % 32))
		}
		if v23 != 0 {
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceOldestClogXid[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = l0
		} else {
		}
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_AdvanceOldestClogXid[0]))
		F_LWLockRelease(m, v28+int32(_a_F_AdvanceOldestClogXid_0))
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
					F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
						if v72 == int32(0) {
							m.G0 = v9 + int32(32)
							return
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
							v77 = int32(0)
							F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																	F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9+int32(16))
																	mBase = m.M
																	v100 = m.ExcPending
																	if v100 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																		mBase = m.M
																		v104 = m.ExcPending
																		if v104 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2032), int32(_a_F_AlterPublicationOwner_internal_4))
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
																		F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return
																		} else {
																			v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																			if v72 == int32(0) {
																				m.G0 = v9 + int32(32)
																				return
																			} else {
																				v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				v77 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																					F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																					mBase = m.M
																					v122 = m.ExcPending
																					if v122 != 0 {
																						return
																					} else {
																						F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_5), int32(0))
																						mBase = m.M
																						v126 = m.ExcPending
																						if v126 != 0 {
																							return
																						} else {
																							F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2039), int32(_a_F_AlterPublicationOwner_internal_4))
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
																				F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																				mBase = m.M
																				v70 = m.ExcPending
																				if v70 != 0 {
																					return
																				} else {
																					v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																					if v72 == int32(0) {
																						m.G0 = v9 + int32(32)
																						return
																					} else {
																						v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																						v77 = int32(0)
																						F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																mBase = m.M
																v70 = m.ExcPending
																if v70 != 0 {
																	return
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																	if v72 == int32(0) {
																		m.G0 = v9 + int32(32)
																		return
																	} else {
																		v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v77 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																			F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return
																			} else {
																				F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_5), int32(0))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2039), int32(_a_F_AlterPublicationOwner_internal_4))
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
																		F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return
																		} else {
																			v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																			if v72 == int32(0) {
																				m.G0 = v9 + int32(32)
																				return
																			} else {
																				v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				v77 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
															F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9+int32(16))
															mBase = m.M
															v100 = m.ExcPending
															if v100 != 0 {
																return
															} else {
																F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																mBase = m.M
																v104 = m.ExcPending
																if v104 != 0 {
																	return
																} else {
																	F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2032), int32(_a_F_AlterPublicationOwner_internal_4))
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
																F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																mBase = m.M
																v70 = m.ExcPending
																if v70 != 0 {
																	return
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																	if v72 == int32(0) {
																		m.G0 = v9 + int32(32)
																		return
																	} else {
																		v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v77 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																			F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																			mBase = m.M
																			v122 = m.ExcPending
																			if v122 != 0 {
																				return
																			} else {
																				F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_5), int32(0))
																				mBase = m.M
																				v126 = m.ExcPending
																				if v126 != 0 {
																					return
																				} else {
																					F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2039), int32(_a_F_AlterPublicationOwner_internal_4))
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
																		F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																		mBase = m.M
																		v70 = m.ExcPending
																		if v70 != 0 {
																			return
																		} else {
																			v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																			if v72 == int32(0) {
																				m.G0 = v9 + int32(32)
																				return
																			} else {
																				v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																				v77 = int32(0)
																				F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
														F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
														mBase = m.M
														v70 = m.ExcPending
														if v70 != 0 {
															return
														} else {
															v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
															if v72 == int32(0) {
																m.G0 = v9 + int32(32)
																return
															} else {
																v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																v77 = int32(0)
																F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																	F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																	mBase = m.M
																	v122 = m.ExcPending
																	if v122 != 0 {
																		return
																	} else {
																		F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_5), int32(0))
																		mBase = m.M
																		v126 = m.ExcPending
																		if v126 != 0 {
																			return
																		} else {
																			F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2039), int32(_a_F_AlterPublicationOwner_internal_4))
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
																F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																mBase = m.M
																v70 = m.ExcPending
																if v70 != 0 {
																	return
																} else {
																	v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																	if v72 == int32(0) {
																		m.G0 = v9 + int32(32)
																		return
																	} else {
																		v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																		v77 = int32(0)
																		F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9+int32(16))
																mBase = m.M
																v100 = m.ExcPending
																if v100 != 0 {
																	return
																} else {
																	F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
																	mBase = m.M
																	v104 = m.ExcPending
																	if v104 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2032), int32(_a_F_AlterPublicationOwner_internal_4))
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
																	F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return
																	} else {
																		v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																		if v72 == int32(0) {
																			m.G0 = v9 + int32(32)
																			return
																		} else {
																			v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v77 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																				F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																				mBase = m.M
																				v122 = m.ExcPending
																				if v122 != 0 {
																					return
																				} else {
																					F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_5), int32(0))
																					mBase = m.M
																					v126 = m.ExcPending
																					if v126 != 0 {
																						return
																					} else {
																						F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2039), int32(_a_F_AlterPublicationOwner_internal_4))
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
																			F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																			mBase = m.M
																			v70 = m.ExcPending
																			if v70 != 0 {
																				return
																			} else {
																				v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																				if v72 == int32(0) {
																					m.G0 = v9 + int32(32)
																					return
																				} else {
																					v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																					v77 = int32(0)
																					F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
															F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																if v72 == int32(0) {
																	m.G0 = v9 + int32(32)
																	return
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v77 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																		F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return
																		} else {
																			F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_5), int32(0))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2039), int32(_a_F_AlterPublicationOwner_internal_4))
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
																	F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return
																	} else {
																		v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																		if v72 == int32(0) {
																			m.G0 = v9 + int32(32)
																			return
																		} else {
																			v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v77 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
														F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9+int32(16))
														mBase = m.M
														v100 = m.ExcPending
														if v100 != 0 {
															return
														} else {
															F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_2), int32(0))
															mBase = m.M
															v104 = m.ExcPending
															if v104 != 0 {
																return
															} else {
																F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2032), int32(_a_F_AlterPublicationOwner_internal_4))
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
															F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																if v72 == int32(0) {
																	m.G0 = v9 + int32(32)
																	return
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v77 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																		F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																		mBase = m.M
																		v122 = m.ExcPending
																		if v122 != 0 {
																			return
																		} else {
																			F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_5), int32(0))
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2039), int32(_a_F_AlterPublicationOwner_internal_4))
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
																	F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
																	mBase = m.M
																	v70 = m.ExcPending
																	if v70 != 0 {
																		return
																	} else {
																		v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																		if v72 == int32(0) {
																			m.G0 = v9 + int32(32)
																			return
																		} else {
																			v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																			v77 = int32(0)
																			F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
													F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
													mBase = m.M
													v70 = m.ExcPending
													if v70 != 0 {
														return
													} else {
														v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
														if v72 == int32(0) {
															m.G0 = v9 + int32(32)
															return
														} else {
															v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
															v77 = int32(0)
															F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
																F_errmsg(m, int32(_a_F_AlterPublicationOwner_internal_1), v9)
																mBase = m.M
																v122 = m.ExcPending
																if v122 != 0 {
																	return
																} else {
																	F_errhint(m, int32(_a_F_AlterPublicationOwner_internal_5), int32(0))
																	mBase = m.M
																	v126 = m.ExcPending
																	if v126 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_AlterPublicationOwner_internal_3), int32(2039), int32(_a_F_AlterPublicationOwner_internal_4))
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
															F_changeDependencyOnOwner(m, int32(_a_F_AlterPublicationOwner_internal_0), v68, l2)
															mBase = m.M
															v70 = m.ExcPending
															if v70 != 0 {
																return
															} else {
																v72 = *(*int32)(unsafe.Add(mBase, _c_F_AlterPublicationOwner_internal[0]))
																if v72 == int32(0) {
																	m.G0 = v9 + int32(32)
																	return
																} else {
																	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
																	v77 = int32(0)
																	F_RunObjectPostAlterHook(m, int32(_a_F_AlterPublicationOwner_internal_0), v76, v77, v77, v77)
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v124 int32
	_ = v124
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v276 int32
	_ = v276
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v337 int32
	_ = v337
	var v340 int32
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
	var v371 int32
	_ = v371
	var v382 int32
	_ = v382
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v444 int32
	_ = v444
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int64
	_ = v549
	var v550 int32
	_ = v550
	var v566 int32
	_ = v566
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v604 int32
	_ = v604
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v692 int32
	_ = v692
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v806 int32
	_ = v806
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v871 int32
	_ = v871
	var v885 int32
	_ = v885
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v916 int32
	_ = v916
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v938 int64
	_ = v938
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v996 int32
	_ = v996
	var v1003 int32
	_ = v1003
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1030 int32
	_ = v1030
	var v1048 int32
	_ = v1048
	var v1076 int32
	_ = v1076
	var v1077 int64
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	v4 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(400)
	m.G0 = v30
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = int32(105)
	goto L3
L2:
	;
	v35 = int32(114)
	goto L3
L3:
	;
	v37 = l0 + int32(16)
	v42 = v4
	v43 = v4
	v44 = v4
	v45 = v4
	v46 = v4
	v47 = v4
	v48 = v4
	v49 = v4
	v50 = v4
	v51 = v4
	v52 = int32(-1)
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
	m.G0 = v30 + int32(400)
	return
L7:
	;
	goto L6
L8:
	;
	if v52 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v1076 = int32(m.ExcTag)
	v1077 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v1076 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v46
	F_load_file(m, int32(_a_F_AlterSubscription_refresh_0), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v181 = v43
	v182 = v44
	v183 = v45
	v184 = v46
	v185 = v51
	goto L13
L13:
	;
	if v185 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L14:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v81 == int32(1) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v87 = v84 ^ int32(1)
	goto L17
L16:
	;
	v87 = int32(0)
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[0]))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v37
	v102 = int32(1)
	v108 = m.T0[v92].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, v89, v102, v102, v87&v102, v88, v30+int32(360))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	if v108 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v37
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[1]))
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2]))
	goto L26
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v37
	F_errcode(m, int32(100663808))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L10
	} else {
		goto L23
	}
L23:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v37
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v30)+360))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+36)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v30)+32)) = v137
	F_errmsg(m, int32(_a_F_AlterSubscription_refresh_1), v30+int32(32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v37
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_2), int32(853), int32(_a_F_AlterSubscription_refresh_3))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	goto L4
L26:
	;
	v175 = v30 + int32(192)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v30 + int32(44)
	goto L29
L27:
	;
	v181 = v108
	v182 = v171
	v183 = v173
	v184 = v37
	v185 = int32(0)
	goto L13
L29:
	;
	goto L27
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_pg_qsort(m, v629, v630, int32(4), int32(471))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L10
	} else {
		goto L79
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2])) = v30 + int32(192)
	if l2 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[1])) = v182
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2])) = v183
	v592 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[0]))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v592)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	m.T0[v593].(func(*base.Module, int32))(m, v181)
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L10
	} else {
		goto L77
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_check_publications(m, v181, l2)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L10
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v214 = F_fetch_table_list(m, v181, v204)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L10
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v227 = F_GetSubscriptionRelations(m, v216, int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L40
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_pg_qsort(m, v355, v344, int32(4), int32(471))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L10
	} else {
		goto L52
	}
L40:
	;
	if v227 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v240 = int32(0)
	v242 = F_palloc(m, v240)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v256 = F_palloc(m, v244<<(uint(int32(2))%32))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L10
	} else {
		goto L45
	}
L44:
	;
	v337 = v47
	v340 = v242
	v344 = v240
	v355 = v242
	goto L39
L45:
	;
	v258 = int32(0)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v258 < v259 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v276 = v258
	goto L49
L47:
	;
	goto L48
L48:
	;
	v337 = v256
	v340 = v50
	v344 = v244
	v355 = v256
	goto L39
L49:
	;
	v290 = v276 << (uint(int32(2)) % 32)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v292+v290)))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	*(*int32)(unsafe.Add(mBase, uint32(v256+v290))) = v295
	v298 = v276 + int32(1)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	if v298 < v299 {
		v276 = v298
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L48
L51:
	;
	goto L50
L52:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_check_publications_origin(m, v181, v371, l1, v370, v355, v344, v369)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L10
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v394 = F_palloc(m, v344<<(uint(int32(3))%32))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	if v214 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v407 = int32(0)
	v409 = F_palloc(m, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L10
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v423 = F_palloc(m, v411<<(uint(int32(2))%32))
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L10
	} else {
		goto L59
	}
L58:
	;
	v620 = v42
	v627 = v409
	v629 = v409
	v630 = v407
	goto L30
L59:
	;
	v425 = int32(0)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v426 <= v425 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v620 = v423
	v627 = v49
	v629 = v423
	v630 = v426
	goto L30
L61:
	;
	goto L62
L62:
	;
	v444 = v425
	goto L63
L63:
	;
	v457 = v444 << (uint(int32(2)) % 32)
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v457+v458)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v471 = int32(0)
	v474 = F_RangeVarGetRelidExtended(m, v460, int32(1), v471, v471, v471)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L10
	} else {
		goto L65
	}
L64:
	;
	v620 = v423
	v627 = v49
	v629 = v423
	v630 = v585
	goto L30
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+188)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v486 = F_get_rel_relkind(m, v474)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L10
	} else {
		goto L66
	}
L66:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v460)+12))
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v460)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_CheckSubscriptionRelkind(m, v486, v489, v488)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L10
	} else {
		goto L67
	}
L67:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v30)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v423+v457))) = v502
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v517 = F_bsearch(m, v30+int32(188), v355, v344, int32(4), int32(471))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L10
	} else {
		goto L69
	}
L68:
	;
	v584 = v444 + int32(1)
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v584 < v585 {
		v444 = v584
		goto L63
	} else {
		goto L76
	}
L69:
	;
	if v517 != 0 {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v30)+188))
	F_AddSubscriptionRelState(m, v519, v529, v35, int64(0), int32(1))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v545 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	if v545 == int32(0) {
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v549 = *(*int64)(unsafe.Add(mBase, uint32(v460)+8))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v550
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v549
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_refresh_4), v30+int32(16))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L10
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v423
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_2), int32(924), int32(_a_F_AlterSubscription_refresh_3))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L10
	} else {
		goto L75
	}
L75:
	;
	goto L68
L76:
	;
	goto L64
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v47
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_pg_re_throw(m)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	goto L4
L79:
	;
	if v344 <= int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[1])) = v182
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2])) = v183
	v1018 = *(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[0]))
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1018)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	m.T0[v1019].(func(*base.Module, int32))(m, v181)
	mBase = m.M
	v1030 = m.ExcPending
	if v1030 != 0 {
		goto L10
	} else {
		goto L124
	}
L81:
	;
	v996 = v48
	v1003 = int32(0)
	goto L80
L82:
	;
	goto L83
L83:
	;
	v659 = int32(0)
	v672 = v48
	v676 = v659
	v679 = v659
	v680 = v659
	goto L84
L84:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v355+v676<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+184)) = v692
	if v214 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v900 = int32(0)
	if v894 <= v900 {
		v996 = v891
		v1003 = v893
		goto L80
	} else {
		goto L115
	}
L86:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v696 = v695
	goto L88
L87:
	;
	v696 = int32(0)
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v710 = F_bsearch(m, v30+int32(184), v629, v696, int32(4), int32(471))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L10
	} else {
		goto L89
	}
L89:
	;
	if v710 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	if v679 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	v891 = v672
	v893 = v679
	v894 = v680
	goto L92
L92:
	;
	v898 = v676 + int32(1)
	if v898 != v344 {
		v672 = v891
		v676 = v898
		v679 = v893
		v680 = v894
		goto L84
	} else {
		goto L114
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v672
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v727 = F_table_open(m, int32(_a_F_AlterSubscription_refresh_5), int32(8))
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L10
	} else {
		goto L96
	}
L94:
	;
	v729 = v672
	v730 = v679
	goto L95
L95:
	;
	v731 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v30)+184))
	v744 = F_GetSubscriptionRelState(m, v731, v741, v30+int32(176))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L10
	} else {
		goto L97
	}
L96:
	;
	v729 = v727
	v730 = v727
	goto L95
L97:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v30)+184))
	v749 = v394 + v680<<(uint(int32(3))%32)
	*(*uint8)(unsafe.Add(mBase, uint32(v749)+4)) = uint8(v744)
	*(*int32)(unsafe.Add(mBase, uint32(v749))) = v746
	v752 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_RemoveSubscriptionRel(m, v752, v746)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L10
	} else {
		goto L98
	}
L98:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v30)+184))
	F_logicalrep_worker_stop(m, v764, v774)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L10
	} else {
		goto L99
	}
L99:
	;
	if v744 != int32(114) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v30)+184))
	v791 = v30 + int32(112)
	F_ReplicationOriginNameForLogicalRep(m, v779, v789, v791)
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L10
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v819 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L10
	} else {
		goto L105
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_replorigin_drop_by_name(m, v791, int32(1), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	goto L102
L105:
	;
	if v819 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v30)+184))
	v831 = F_get_rel_namespace(m, v830)
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L10
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v891 = v729
	v893 = v730
	v894 = v680 + int32(1)
	goto L92
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v842 = F_get_namespace_name(m, v831)
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v30)+184))
	v854 = F_get_rel_name(m, v853)
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L10
	} else {
		goto L111
	}
L111:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v856
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v854
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v842
	F_errmsg_internal(m, int32(_a_F_AlterSubscription_refresh_6), v30)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L10
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v729
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_errfinish(m, int32(_a_F_AlterSubscription_refresh_2), int32(1000), int32(_a_F_AlterSubscription_refresh_3))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L10
	} else {
		goto L113
	}
L113:
	;
	goto L108
L114:
	;
	goto L85
L115:
	;
	v916 = v900
	goto L116
L116:
	;
	v932 = v394 + v916<<(uint(int32(3))%32)
	v933 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932)+4)))
	if v933&int32(254) != int32(114) {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v996 = v891
	v1003 = v893
	goto L80
L118:
	;
	v938 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+104)) = v938
	*(*int64)(unsafe.Add(mBase, uint32(v30)+96)) = v938
	*(*int64)(unsafe.Add(mBase, uint32(v30)+88)) = v938
	*(*int64)(unsafe.Add(mBase, uint32(v30)+80)) = v938
	*(*int64)(unsafe.Add(mBase, uint32(v30)+72)) = v938
	*(*int64)(unsafe.Add(mBase, uint32(v30)+64)) = v938
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v938
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v938
	v954 = *(*int32)(unsafe.Add(mBase, uint32(v932)))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	v966 = v30 + int32(48)
	F_ReplicationSlotNameForTablesync(m, v955, v954, v966)
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L10
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v984 = v916 + int32(1)
	if v984 != v894 {
		v916 = v984
		goto L116
	} else {
		goto L123
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v891
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_ReplicationSlotDropAtPubNode(m, v181, v966, int32(1))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L10
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	goto L117
L124:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[1])) = v182
	*(*int32)(unsafe.Add(mBase, _c_F_AlterSubscription_refresh[2])) = v183
	if v1003 == int32(0) {
		goto L7
	} else {
		goto L125
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+368)) = v620
	*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = v996
	*(*int32)(unsafe.Add(mBase, uint32(v30)+372)) = v627
	*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v30)+384)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v30)+388)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(v30)+392)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v30)+396)) = v184
	F_relation_close(m, v1003, int32(0))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L10
	} else {
		goto L126
	}
L126:
	;
	goto L9
L127:
	;
	v1081 = int32(v1077)
	m.G0 = v30
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1081)))
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v1084)))
	if v30+int32(44) == v1087 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	m.ExcPending = 1
	goto L136
L129:
	;
	if v1091 != 0 {
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1084)+4))
	v1091 = v1089
	goto L132
L131:
	;
	v1091 = int32(0)
	goto L132
L132:
	;
	goto L129
L133:
	;
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v30)+396))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v30)+392))
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v30)+388))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v30)+384))
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v30)+380))
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v30)+376))
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v30)+372))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v30)+368))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v30)+364))
	v42 = v1099
	v43 = v1093
	v44 = v1095
	v45 = v1094
	v46 = v1092
	v47 = v1097
	v48 = v1100
	v49 = v1098
	v50 = v1096
	v51 = v1083
	v52 = v1091
	goto L5
L134:
	;
	goto L135
L135:
	;
	F___wasm_longjmp(m, v1084, v1083)
	mBase = m.M
	v1102 = m.ExcPending
	if v1102 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	return
L137:
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
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v213 int32
	_ = v213
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int64
	_ = v371
	var v372 int64
	_ = v372
	var v381 int64
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int64
	_ = v394
	var v395 int64
	_ = v395
	var v403 int64
	_ = v403
	var v411 int64
	_ = v411
	var v420 int64
	_ = v420
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int64
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
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
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v535 int32
	_ = v535
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
	F_errmsg_internal(m, int32(_a_F_ApplyLauncherMain_0), int32(0))
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
	F_errfinish(m, int32(_a_F_ApplyLauncherMain_1), int32(1135), int32(_a_F_ApplyLauncherMain_2))
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
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[0]))
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v39
	v42 = int32(914)
	v44 = m.G0
	v46 = v44 - int32(32)
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
	v75 = int32(295)
	v77 = m.G0
	v79 = v77 - int32(32)
	m.G0 = v79
	switch int32(297) {
	case 0, 2:
		v89 = v75
		goto L16
	default:
		goto L17
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v56
	F_sigemptyset(m, v46+int32(16))
	mBase = m.M
	goto L13
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[2])) = v42
	v56 = int32(_a_F_ApplyLauncherMain_3)
	goto L10
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+24)) = int32(268435456)
	v70 = F___sigaction(m, int32(1), v46+int32(12), int32(0))
	mBase = m.M
	m.G0 = v46 + int32(32)
	goto L9
L15:
	;
	F_BackgroundWorkerUnblockSignals(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L21
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+12)) = v89
	F_sigemptyset(m, v79+int32(16))
	mBase = m.M
	goto L19
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[3])) = v75
	v89 = int32(_a_F_ApplyLauncherMain_3)
	goto L16
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+24)) = int32(268435456)
	v103 = F___sigaction(m, int32(15), v79+int32(12), int32(0))
	mBase = m.M
	m.G0 = v79 + int32(32)
	goto L15
L21:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[4]))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+192)))
	if v111&int32(2) != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L36
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L33
	}
L24:
	;
	v114 = int32(0)
	F_InitPostgres(m, v114, v114, v114, v114, v114, v114)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[5]))
	if v123 != int32(1) {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[5])) = int32(2)
	goto L22
L29:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_ApplyLauncherMain_4), int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(_a_F_ApplyLauncherMain_5), int32(869), int32(_a_F_ApplyLauncherMain_6))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	F_errmsg(m, int32(_a_F_ApplyLauncherMain_7), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_ApplyLauncherMain_5), int32(879), int32(_a_F_ApplyLauncherMain_6))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
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
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[6]))
	if v173 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v176 = int32(0)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[7]))
	v183 = F_AllocSetContextCreateInternal(m, v178, int32(_a_F_ApplyLauncherMain_8), v176, int32(_a_F_ApplyLauncherMain_9), int32(_a_F_ApplyLauncherMain_10))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v185 = int32(_a_F_ApplyLauncherMain_11)
	v186 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8]))
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8])) = v183
	F_StartTransactionCommand(m)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v193 = F_table_open(m, int32(_a_F_ApplyLauncherMain_12), int32(1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v195 = int32(0)
	v197 = F_table_beginscan_catalog(m, v193, v195, v195)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v203 = v176
	goto L46
L46:
	;
	v213 = F_heap_getnext(m, v197)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v197)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+188))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	m.T0[v244].(func(*base.Module, int32))(m, v197)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L55
	}
L48:
	;
	if v213 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v215 = int32(_a_F_ApplyLauncherMain_11)
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8]))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v217)+22)))
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8])) = v183
	v222 = F_palloc0(m, int32(56))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L47
L52:
	;
	v224 = v217 + v218
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	*(*int32)(unsafe.Add(mBase, uint32(v222))) = v225
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v222)+4)) = v227
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v224)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v222)+20)) = v229
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+84)))
	*(*uint8)(unsafe.Add(mBase, uint32(v222)+25)) = uint8(v231)
	v235 = F_pstrdup(m, v224+int32(16))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v222)+16)) = v235
	v238 = F_lappend(m, v203, v222)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8])) = v216
	v203 = v238
	goto L46
L55:
	;
	F_relation_close(m, v193, int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	if v203 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[8])) = v186
	F_MemoryContextDelete(m, v183)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L110
	}
L59:
	;
	v489 = int32(_a_F_ApplyLauncherMain_13)
	goto L58
L60:
	;
	goto L61
L61:
	;
	v255 = int32(_a_F_ApplyLauncherMain_13)
	v256 = int32(0)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	if v257 <= v256 {
		v489 = v255
		goto L58
	} else {
		goto L62
	}
L62:
	;
	v261 = v255
	v266 = v256
	goto L63
L63:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274+v266<<(uint(int32(2))%32))))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+25)))
	if v279 == int32(0) {
		v471 = v261
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v489 = v471
	goto L58
L65:
	;
	v485 = v266 + int32(1)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v203)+4))
	if v485 < v486 {
		v261 = v471
		v266 = v485
		goto L63
	} else {
		goto L109
	}
L66:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[9]))
	v287 = F_LWLockAcquire(m, v283+int32(_a_F_ApplyLauncherMain_14), int32(1))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[10]))
	if int32(0) < v290 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v465 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[9]))
	F_LWLockRelease(m, v465+int32(_a_F_ApplyLauncherMain_14))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L108
	}
L69:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v296 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[0]))
	v299 = int32(0)
	goto L72
L70:
	;
	goto L71
L71:
	;
	v345 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[9]))
	F_LWLockRelease(m, v345+int32(_a_F_ApplyLauncherMain_14))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L80
	}
L72:
	;
	v315 = v296 + int32(16) + v299*int32(112)
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+16)))
	if v316 != int32(1) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	goto L71
L74:
	;
	v328 = v299 + int32(1)
	if v328 != v290 {
		v299 = v328
		goto L72
	} else {
		goto L79
	}
L75:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	if v319 == int32(3) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v315)+32))
	if v322 != v293 {
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v315)+36))
	if v324 == int32(0) {
		goto L68
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	goto L73
L80:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v350
	F_logicalrep_launcher_attach_dshmem(m)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[11]))
	v359 = F_dshash_find(m, v355, v17+int32(4), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	v461 = v425 - v423
	if v261 < v461 {
		goto L105
	} else {
		goto L106
	}
L83:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v431
	F_logicalrep_launcher_attach_dshmem(m)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L97
	}
L84:
	;
	if v359 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v366 = m.G0
	v367 = int32(16)
	v368 = v366 - v367
	m.G0 = v368
	F_gettimeofday(m, v368)
	mBase = m.M
	v371 = *(*int64)(unsafe.Add(mBase, uint32(v368)))
	v372 = int64(*(*int32)(unsafe.Add(mBase, uint32(v368)+8)))
	m.G0 = v368 + v367
	goto L88
L86:
	;
	goto L87
L87:
	;
	v381 = *(*int64)(unsafe.Add(mBase, uint32(v359)+8))
	v383 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[11]))
	F_dshash_release_lock(m, v383, v359)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L89
	}
L88:
	;
	v429 = v372 + v371*int64(1000000) - int64(946684800000000)
	goto L83
L89:
	;
	v389 = m.G0
	v390 = int32(16)
	v391 = v389 - v390
	m.G0 = v391
	F_gettimeofday(m, v391)
	mBase = m.M
	v394 = *(*int64)(unsafe.Add(mBase, uint32(v391)))
	v395 = int64(*(*int32)(unsafe.Add(mBase, uint32(v391)+8)))
	m.G0 = v391 + v390
	v403 = v395 + v394*int64(1000000) - int64(946684800000000)
	goto L90
L90:
	;
	if v381 == int64(0) {
		v429 = v403
		goto L83
	} else {
		goto L91
	}
L91:
	;
	if v403 <= v381 {
		v423 = int32(0)
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v425 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[12]))
	if v423 < v425 {
		goto L82
	} else {
		goto L96
	}
L93:
	;
	goto L92
L94:
	;
	v411 = v403 - v381
	if base.B2i32(int64(0) < v381)^base.B2i32(v411 < v403)|base.B2i32(int64(2147483646000) < v411) != 0 {
		v423 = int32(2147483647)
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v420 = base.I64_div_s(v411+int64(999), int64(1000))
	v423 = base.I32_wrap_i64(v420)
	goto L93
L96:
	;
	v429 = v403
	goto L83
L97:
	;
	v436 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[11]))
	v441 = F_dshash_find_or_insert(m, v436, v17+int32(12), v17+int32(11))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v441)+8)) = v429
	v445 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[11]))
	F_dshash_release_lock(m, v445, v441)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v278)+4))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v278)))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v278)+16))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v278)+20))
	v453 = int32(0)
	v455 = F_logicalrep_worker_launch(m, int32(2), v449, v450, v451, v452, v453, v453)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	if v455 != 0 {
		v471 = v261
		goto L65
	} else {
		goto L101
	}
L101:
	;
	v458 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[12]))
	if v261 < v458 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v460 = v261
	goto L104
L103:
	;
	v460 = v458
	goto L104
L104:
	;
	v471 = v460
	goto L65
L105:
	;
	v463 = v261
	goto L107
L106:
	;
	v463 = v461
	goto L107
L107:
	;
	v471 = v463
	goto L65
L108:
	;
	v471 = v261
	goto L65
L109:
	;
	goto L64
L110:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[13]))
	v510 = F_WaitLatch(m, v507, int32(41), v489, int32(83886088))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	v527 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[14]))
	if v527 == int32(0) {
		goto L36
	} else {
		goto L117
	}
L112:
	;
	if v510&int32(1) == int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v517 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[13]))
	*(*int32)(unsafe.Add(mBase, uint32(v517))) = int32(0)
	goto L114
L114:
	;
	v521 = *(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[6]))
	if v521 == int32(0) {
		goto L111
	} else {
		goto L115
	}
L115:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	goto L111
L117:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ApplyLauncherMain[14])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L36
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
	v97 = m.ExcPending
	if v97 != 0 {
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
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+44)))
	if v23 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_AtAbort_Twophase[1]))
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
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = int32(-1)
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
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v22
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
	F_errmsg_internal(m, int32(_a_F_AtAbort_Twophase_0), v10)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(_a_F_AtAbort_Twophase_1), int32(650), int32(_a_F_AtAbort_Twophase_2))
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
	var v41 int32
	_ = v41
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
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
	*(*int32)(unsafe.Add(mBase, uint32(v122)+44)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5]))
	v129 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v128))), uint32(v129))
	v132 = int32(_a_F_AuxiliaryProcessMainCommon_0)
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[6])) = v122
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[2]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v140 = base.I32_div_s(v122-v137, int32(640))
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[7])) = v140
	v142 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v122))) = v142
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+612)) = v129
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+608)) = uint8(v129)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+16)) = v129
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+124)) = uint8(v129)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+120)) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v145)+52)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v145)+36)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v145)+92)) = v142
	*(*uint16)(unsafe.Add(mBase, uint32(v145)+74)) = uint16(v129)
	*(*int64)(unsafe.Add(mBase, uint32(v145)+56)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v145)+64)) = v142
	*(*uint8)(unsafe.Add(mBase, uint32(v145)+72)) = uint8(v129)
	v172 = base.AtomicRmwXchg64(m, v145, int32(112), v142)
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[6]))
	F_OwnLatch(m, v174+int32(20))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
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
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L35
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[8]))
	if v26 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[6]))
	if v30 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[9])))
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
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5]))
	v41 = base.AtomicRmwXchg32(m, v38, int32(0), int32(1))
	if v41 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5]))
	F_s_lock(m, v43, int32(_a_F_AuxiliaryProcessMainCommon_1), int32(640), int32(_a_F_AuxiliaryProcessMainCommon_2))
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
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[2]))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+68))
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[10])) = v51
	goto L24
L23:
	;
	goto L22
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[8]))
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
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[5]))
	v80 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v79))), uint32(v80))
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
	v123 = v56
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
		v122 = v70
		v123 = v67
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
	F_errmsg_internal(m, int32(_a_F_AuxiliaryProcessMainCommon_3), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_AuxiliaryProcessMainCommon_1), int32(656), int32(_a_F_AuxiliaryProcessMainCommon_2))
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
	F_errmsg_internal(m, int32(_a_F_AuxiliaryProcessMainCommon_4), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_AuxiliaryProcessMainCommon_1), int32(625), int32(_a_F_AuxiliaryProcessMainCommon_2))
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
	F_errmsg_internal(m, int32(_a_F_AuxiliaryProcessMainCommon_5), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_AuxiliaryProcessMainCommon_1), int32(628), int32(_a_F_AuxiliaryProcessMainCommon_2))
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
	v180 = m.ExcPending
	if v180 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[6]))
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[11])) = v182 + int32(548)
	goto L43
L43:
	;
	v188 = *(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[6]))
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	F_PGSemaphoreReset(m, v189)
	mBase = m.M
	F_on_shmem_exit(m, int32(1118), v123)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[9])))
	if v195 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_AttachSharedMemoryStructs(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
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
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v202 = int32(0)
	F_ProcSignalInit(m, v202, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_CreateAuxProcessResourceOwner(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_pgstat_beinit(m)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	F_pgstat_bestart_initial(m)
	mBase = m.M
	F_pgstat_bestart_final(m)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_before_shmem_exit(m, int32(922), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AuxiliaryProcessMainCommon[12])) = int32(2)
	return
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
				F_errfinish(m, int32(_a_F_acldefault_1), int32(872), int32(_a_F_acldefault_2))
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
	case 6, 15, 43:
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
	case 30:
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
	case 31:
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
	case 35:
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
	case 36:
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
func F_aclexplode(m *base.Module, l0 int32) int32 {
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
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	goto L15
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
	v65 = F_BlessTupleDesc(m, v35)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v65
	v69 = F_palloc(m, int32(8))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v69
	*(*int32)(unsafe.Add(mBase, _c_F_aclexplode[0])) = v30
	goto L5
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v82 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v90 = v82
	goto L18
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v90 = (v83<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L18:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v94 = v92
	goto L22
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L35
	}
L20:
	;
	m.G0 = v14 + int32(32)
	return v180
L21:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L34
	}
L22:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v104 <= v94 {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = int32(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v137
	if base.Ui32(int32(15)) <= base.Ui32(v121) {
		goto L19
	} else {
		goto L30
	}
L24:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v108 = v106 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v108
	if v108 == int32(15) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v112 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v112
	v116 = v94 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v118 <= v116 {
		goto L21
	} else {
		goto L28
	}
L26:
	;
	v120 = v94
	v121 = v108
	goto L27
L27:
	;
	v123 = base.I64_extend_i32_u(v121)
	v124 = int64(1) << (uint(v123) % 64)
	v127 = v90 + v17 + v120<<(uint(int32(4))%32)
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v127)+8))
	if base.I32_wrap_i64(v124&v128) == int32(0) {
		v94 = v120
		goto L22
	} else {
		goto L29
	}
L28:
	;
	v120 = v116
	v121 = v112
	goto L27
L29:
	;
	goto L23
L30:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v121<<(uint(int32(2))%32))+uint32(_c_F_aclexplode[1])))
	v144 = F_cstring_to_text(m, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v144
	v147 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v127)+12)))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = base.I32_wrap_i64(int64(base.Ui64(v147)>>(uint(v123)%64))) & int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v158 = F_heap_form_tuple(m, v153, v14+int32(16), v14+int32(12))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v161 = F_HeapTupleHeaderGetDatum(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v80)))
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v163 + int64(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+20)) = int32(1)
	v180 = v161
	goto L20
L34:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v174)+20)) = int32(2)
	v177 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v177)
	v180 = int32(0)
	goto L20
L35:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v14))) = uint32(v124)
	F_errmsg_internal(m, int32(_a_F_aclexplode_5), v14)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_aclexplode_6), int32(1770), int32(_a_F_aclexplode_7))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_aclitemout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v58 int32
	_ = v58
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
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
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v169 int64
	_ = v169
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v179 int32
	_ = v179
	var v181 int64
	_ = v181
	var v184 int32
	_ = v184
	var v185 int64
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v197 int64
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
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
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_palloc(m, int32(293))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v24)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v26 == v24 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v153 = F_strlen(m, v20)
	mBase = m.M
	v154 = v153 + v20
	v155 = int32(61)
	*(*uint8)(unsafe.Add(mBase, uint32(v154))) = uint8(v155)
	v160 = v154 + int32(1)
	v169 = int64(0)
	goto L32
L4:
	;
	v30 = F_SearchSysCache1(m, int32(11), v26)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v30)+16))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+22)))
	v34 = v32 + v33
	v36 = v34 + int32(4)
	v37 = int32(1)
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+4)))
	if v38 == int32(0) {
		v87 = v37
		v95 = v20
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v133
	v138 = F_pg_sprintf(m, v20, int32(_a_F_aclitemout_0), v16+int32(16))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L31
	}
L9:
	;
	v96 = v95
	v98 = v36
	goto L20
L10:
	;
	v41 = v36
	v42 = v38
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
	v87 = v37
	v95 = v20
	goto L9
L13:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v81 != 0 {
		v41 = v41 + int32(1)
		v42 = v81
		goto L11
	} else {
		goto L19
	}
L14:
	;
	v58 = v42 & int32(255)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v74 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v74)
	v87 = int32(0)
	v95 = v20 + int32(1)
	goto L9
L17:
	;
	if base.B2i32(base.Ui32(v58-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v58|int32(32)-int32(97)) < base.Ui32(int32(26)))|base.B2i32(v58 == int32(95)) != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L12
L20:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	if v109 != int32(34) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v126))) = uint8(v127)
	v129 = int32(1)
	v96 = v126 + v129
	v98 = v98 + v129
	goto L20
L23:
	;
	if v109 != 0 {
		v126 = v96
		v127 = v109
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v121 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v121)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v126 = v96 + int32(1)
	v127 = v125
	goto L22
L26:
	;
	if v87 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v116 = v96
	goto L29
L28:
	;
	v112 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v112)
	v116 = v96 + int32(1)
	goto L29
L29:
	;
	v117 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v117)
	F_ReleaseCatCache(m, v30)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L3
L31:
	;
	goto L3
L32:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	v174 = int64(1) << (uint(v169) % 64)
	if v172&v174 != int64(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v200 = int32(47)
	*(*uint16)(unsafe.Add(mBase, uint32(v195))) = uint16(v200)
	v202 = int32(1)
	v204 = v195 + v202
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v207 = F_SearchSysCache1(m, int32(11), v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L42
	}
L34:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v169))+uint32(_c_F_aclitemout[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v179)
	v181 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	v184 = v160 + int32(1)
	v185 = v181
	goto L36
L35:
	;
	v184 = v160
	v185 = v172
	goto L36
L36:
	;
	if int64(base.Ui64(v185)>>(uint(int64(32))%64))&v174 != int64(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v191 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v184))) = uint8(v191)
	v195 = v184 + int32(1)
	goto L39
L38:
	;
	v195 = v184
	goto L39
L39:
	;
	v197 = v169 + int64(1)
	if v197 != int64(15) {
		v160 = v195
		v169 = v197
		goto L32
	} else {
		goto L40
	}
L40:
	;
	goto L33
L41:
	;
	m.G0 = v16 + int32(32)
	return v20
L42:
	;
	if v207 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v207)+16))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+22)))
	v211 = v209 + v210
	v213 = v211 + int32(4)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+4)))
	if v214 == int32(0) {
		v257 = v204
		v265 = v202
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v307
	v310 = F_pg_sprintf(m, v204, int32(_a_F_aclitemout_0), v16)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L66
	}
L46:
	;
	v270 = v257
	v272 = v213
	goto L55
L47:
	;
	v220 = v214
	v222 = v213
	goto L48
L48:
	;
	if base.I32_extend8_s(v220) < int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v252 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+1)) = uint8(v252)
	v257 = v195 + int32(2)
	v265 = int32(0)
	goto L46
L50:
	;
	goto L49
L51:
	;
	goto L52
L52:
	;
	if base.B2i32(base.B2i32(base.Ui32(v220-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v220|int32(32)-int32(97)) < base.Ui32(int32(26))) == int32(0))&base.B2i32(v220 != int32(95)) != 0 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v222)+1)))
	if v251 != 0 {
		v220 = v251
		v222 = v222 + int32(1)
		goto L48
	} else {
		goto L54
	}
L54:
	;
	v257 = v204
	v265 = v202
	goto L46
L55:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	if v283 != int32(34) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v300))) = uint8(v301)
	v303 = int32(1)
	v270 = v300 + v303
	v272 = v272 + v303
	goto L55
L58:
	;
	if v283 != 0 {
		v300 = v270
		v301 = v283
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v295 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(v295)
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	v300 = v270 + int32(1)
	v301 = v299
	goto L57
L61:
	;
	if v265 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v290 = v270
	goto L64
L63:
	;
	v286 = int32(34)
	*(*uint8)(unsafe.Add(mBase, uint32(v270))) = uint8(v286)
	v290 = v270 + int32(1)
	goto L64
L64:
	;
	v291 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v291)
	F_ReleaseCatCache(m, v207)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	goto L41
L66:
	;
	goto L41
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
					F_errfinish(m, int32(_a_F_aclnewowner_1), int32(433), int32(_a_F_aclnewowner_2))
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
	F_errfinish(m, int32(_a_F_add_parameter_name_2), int32(930), int32(_a_F_add_parameter_name_3))
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
				F_errmsg(m, int32(_a_F_add_size_0), int32(0))
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_add_size_1), int32(502), int32(_a_F_add_size_2))
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
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
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
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
	var v217 int32
	_ = v217
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
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
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
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
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
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
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
	var v394 int32
	_ = v394
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
	var v404 int32
	_ = v404
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
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
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
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
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
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
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v570 int64
	_ = v570
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
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
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l0 == v3 {
		v595 = v3
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L16
	} else {
		goto L227
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v662 = m.ExcPending
	if v662 != 0 {
		goto L16
	} else {
		goto L224
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L16
	} else {
		goto L221
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L16
	} else {
		goto L217
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L16
	} else {
		goto L213
	}
L6:
	;
	m.G0 = v14 + int32(32)
	return v595
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v20 - int32(318) {
	case 0:
		goto L9
	case 1:
		goto L10
	default:
		goto L11
	}
L8:
	;
	v591 = F_expression_tree_mutator_impl(m, l0, int32(854), l1)
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L16
	} else {
		goto L212
	}
L9:
	;
	v306 = F_palloc0(m, int32(168))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L16
	} else {
		goto L114
	}
L10:
	;
	v255 = F_expression_tree_mutator_impl(m, l0, int32(854), l1)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L16
	} else {
		goto L93
	}
L11:
	;
	if v20 != int32(58) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v20 != int32(6) {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v227 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L16
	} else {
		goto L85
	}
L15:
	;
	v27 = F_copyObjectImpl(m, l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v31 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v595 = v27
	goto L6
L19:
	;
	goto L20
L20:
	;
	if v18 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v595 = v27
	goto L6
L22:
	;
	goto L23
L23:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v39 = v3
	goto L25
L24:
	;
	if v34 != int32(-4) {
		goto L65
	} else {
		goto L66
	}
L25:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v19+v39<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v50 != v34 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
	v56 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+40)) = uint16(v56)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v55
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
	if v56 < v61 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v53 = v39 + int32(1)
	if v18 != v53 {
		v39 = v53
		goto L25
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	goto L24
L31:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	if v64 == int32(0) {
		goto L5
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	if v61 != 0 {
		goto L47
	} else {
		goto L48
	}
L34:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v67 < v61 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69+v61<<(uint(int32(2))%32)-int32(4))))
	v76 = F_copyObjectImpl(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L16
	} else {
		goto L36
	}
L36:
	;
	if v76 == int32(0) {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v81 == int32(6) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+32)) = v80
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v76)+24))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	v87 = F_bms_add_members(m, v85, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L16
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v80 != 0 {
		goto L3
	} else {
		goto L42
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = v87
	v595 = v76
	goto L6
L42:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v90 == int32(0) {
		v595 = v76
		goto L6
	} else {
		goto L43
	}
L43:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L16
	} else {
		goto L44
	}
L44:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_0), int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L16
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(306), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L16
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
	v595 = v27
	goto L6
L48:
	;
	goto L49
L49:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	if v106 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	if v106 == v107 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+52))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v124+v125<<(uint(int32(2))%32)-int32(4))))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	v133 = F_copyObjectImpl(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L16
	} else {
		goto L57
	}
L53:
	;
	v595 = v27
	goto L6
L54:
	;
	goto L55
L55:
	;
	v110 = F_palloc0(m, int32(20))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L16
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v110))) = int32(30)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v110)+12)) = int64(-4294967294)
	*(*int32)(unsafe.Add(mBase, uint32(v110)+8)) = v115
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v119
	v595 = v110
	goto L6
L57:
	;
	v136 = F_palloc0(m, int32(24))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L16
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = int32(36)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v136)+12)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v141
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v131)+8))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+8))
	v147 = F_copyObjectImpl(m, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L16
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+20)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v136)+16)) = v147
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	if v152 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v27)+24))
	if v153 == int32(0) {
		v595 = v136
		goto L6
	} else {
		goto L61
	}
L61:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L16
	} else {
		goto L62
	}
L62:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_0), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L16
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(364), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L16
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L65:
	;
	v595 = v27
	goto L6
L66:
	;
	goto L67
L67:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v171)+124))
	v173 = int32(0)
	v178 = v173
	v179 = v173
	goto L68
L68:
	;
	v188 = v19 + v179<<(uint(int32(2))%32)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)+8))
	v191 = F_bms_is_member(m, v190, v172)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L16
	} else {
		goto L70
	}
L69:
	;
	if v195 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	if v191 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	if v178 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	v195 = v178
	goto L73
L73:
	;
	v197 = v179 + int32(1)
	if v197 != v18 {
		v178 = v195
		v179 = v197
		goto L68
	} else {
		goto L75
	}
L74:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v193)+8))
	v195 = v194
	goto L73
L75:
	;
	goto L69
L76:
	;
	v595 = v27
	goto L6
L77:
	;
	goto L78
L78:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+132))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v204 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v203+v204<<(uint(int32(2))%32)-int32(4))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+16))
	v212 = F_bms_is_member(m, v195, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	if v212 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	v215 = F_copyObjectImpl(m, v214)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L16
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v225 = F_makeNullConst(m, v222, v223, v224)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L16
	} else {
		goto L84
	}
L83:
	;
	v217 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+40)) = uint16(v217)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+36)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v215)+4)) = v195
	v595 = v215
	goto L6
L84:
	;
	v595 = v225
	goto L6
L85:
	;
	if v18 <= int32(0) {
		v595 = v227
		goto L6
	} else {
		goto L86
	}
L86:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	v236 = v3
	goto L87
L87:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v19+v236<<(uint(int32(2))%32))))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	if v247 != v231 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v246)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v227)+4)) = v252
	v595 = v227
	goto L6
L89:
	;
	v250 = v236 + int32(1)
	if v18 != v250 {
		v236 = v250
		goto L87
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	goto L88
L92:
	;
	v595 = v227
	goto L6
L93:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v255)+20))
	if v257 != 0 {
		v595 = v255
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	if int32(0) < v18 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v264 = v3
	v265 = v3
	goto L98
L96:
	;
	v295 = v3
	goto L97
L97:
	;
	if v295 != 0 {
		goto L111
	} else {
		goto L112
	}
L98:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v19+v265<<(uint(int32(2))%32))))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v277 = F_bms_is_member(m, v276, v258)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L16
	} else {
		goto L100
	}
L99:
	;
	v295 = v288
	goto L97
L100:
	;
	if v277 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if v264 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	v288 = v264
	goto L103
L103:
	;
	v290 = v265 + int32(1)
	if v290 != v18 {
		v264 = v288
		v265 = v290
		goto L98
	} else {
		goto L110
	}
L104:
	;
	v281 = v264
	goto L106
L105:
	;
	v279 = F_bms_copy(m, v258)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L16
	} else {
		goto L107
	}
L106:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	v283 = F_bms_del_member(m, v281, v282)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L16
	} else {
		goto L108
	}
L107:
	;
	v281 = v279
	goto L106
L108:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v275)+8))
	v286 = F_bms_add_member(m, v283, v285)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L16
	} else {
		goto L109
	}
L109:
	;
	v288 = v286
	goto L103
L110:
	;
	goto L99
L111:
	;
	v303 = v295
	goto L113
L112:
	;
	v303 = v258
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v255)+8)) = v303
	v595 = v255
	goto L6
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306))) = int32(318)
	base.MemoryCopy(m, v306, l0, int32(168))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v313 = F_adjust_appendrel_attrs_mutator(m, v312, l1)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L16
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+4)) = v313
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v317 = F_adjust_appendrel_attrs_mutator(m, v316, l1)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L16
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+52)) = v317
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v321 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v328 = v3
	v329 = v3
	goto L120
L118:
	;
	v359 = v3
	goto L119
L119:
	;
	if v359 != 0 {
		goto L133
	} else {
		goto L134
	}
L120:
	;
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v324+v329<<(uint(int32(2))%32))))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v341 = F_bms_is_member(m, v340, v320)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L16
	} else {
		goto L122
	}
L121:
	;
	v359 = v352
	goto L119
L122:
	;
	if v341 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	if v328 != 0 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	v352 = v328
	goto L125
L125:
	;
	v354 = v329 + int32(1)
	if v354 != v321 {
		v328 = v352
		v329 = v354
		goto L120
	} else {
		goto L132
	}
L126:
	;
	v345 = v328
	goto L128
L127:
	;
	v343 = F_bms_copy(m, v320)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L16
	} else {
		goto L129
	}
L128:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	v347 = F_bms_del_member(m, v345, v346)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L16
	} else {
		goto L130
	}
L129:
	;
	v345 = v343
	goto L128
L130:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v339)+8))
	v350 = F_bms_add_member(m, v347, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L16
	} else {
		goto L131
	}
L131:
	;
	v352 = v350
	goto L125
L132:
	;
	goto L121
L133:
	;
	v367 = v359
	goto L135
L134:
	;
	v367 = v320
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+28)) = v367
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v370 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v379 = int32(0)
	v381 = v3
	goto L139
L137:
	;
	v412 = v3
	goto L138
L138:
	;
	if v412 != 0 {
		goto L152
	} else {
		goto L153
	}
L139:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v373+v379<<(uint(int32(2))%32))))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v391 = F_bms_is_member(m, v390, v369)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L16
	} else {
		goto L141
	}
L140:
	;
	v412 = v402
	goto L138
L141:
	;
	if v391 != 0 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	if v381 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v402 = v381
	goto L144
L144:
	;
	v404 = v379 + int32(1)
	if v404 != v370 {
		v379 = v404
		v381 = v402
		goto L139
	} else {
		goto L151
	}
L145:
	;
	v395 = v381
	goto L147
L146:
	;
	v393 = F_bms_copy(m, v369)
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L16
	} else {
		goto L148
	}
L147:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v397 = F_bms_del_member(m, v395, v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L16
	} else {
		goto L149
	}
L148:
	;
	v395 = v393
	goto L147
L149:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v389)+8))
	v400 = F_bms_add_member(m, v397, v399)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L16
	} else {
		goto L150
	}
L150:
	;
	v402 = v400
	goto L144
L151:
	;
	goto L140
L152:
	;
	v417 = v412
	goto L154
L153:
	;
	v417 = v369
	goto L154
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+32)) = v417
	v419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v420 = int32(0)
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v420 < v422 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v430 = v420
	v431 = int32(0)
	goto L158
L156:
	;
	v461 = v420
	goto L157
L157:
	;
	if v461 != 0 {
		goto L171
	} else {
		goto L172
	}
L158:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v425+v431<<(uint(int32(2))%32))))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	v443 = F_bms_is_member(m, v442, v419)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L16
	} else {
		goto L160
	}
L159:
	;
	v461 = v454
	goto L157
L160:
	;
	if v443 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	if v430 != 0 {
		goto L164
	} else {
		goto L165
	}
L162:
	;
	v454 = v430
	goto L163
L163:
	;
	v456 = v431 + int32(1)
	if v456 != v422 {
		v430 = v454
		v431 = v456
		goto L158
	} else {
		goto L170
	}
L164:
	;
	v447 = v430
	goto L166
L165:
	;
	v445 = F_bms_copy(m, v419)
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L16
	} else {
		goto L167
	}
L166:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	v449 = F_bms_del_member(m, v447, v448)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L16
	} else {
		goto L168
	}
L167:
	;
	v447 = v445
	goto L166
L168:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v441)+8))
	v452 = F_bms_add_member(m, v449, v451)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L16
	} else {
		goto L169
	}
L169:
	;
	v454 = v452
	goto L163
L170:
	;
	goto L159
L171:
	;
	v469 = v461
	goto L173
L172:
	;
	v469 = v419
	goto L173
L173:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+40)) = v469
	v471 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v472 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if int32(0) < v472 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v481 = int32(0)
	v483 = v420
	goto L177
L175:
	;
	v514 = v420
	goto L176
L176:
	;
	if v514 != 0 {
		goto L190
	} else {
		goto L191
	}
L177:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v475+v481<<(uint(int32(2))%32))))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	v493 = F_bms_is_member(m, v492, v471)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L16
	} else {
		goto L179
	}
L178:
	;
	v514 = v504
	goto L176
L179:
	;
	if v493 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	if v483 != 0 {
		goto L183
	} else {
		goto L184
	}
L181:
	;
	v504 = v483
	goto L182
L182:
	;
	v506 = v481 + int32(1)
	if v506 != v472 {
		v481 = v506
		v483 = v504
		goto L177
	} else {
		goto L189
	}
L183:
	;
	v497 = v483
	goto L185
L184:
	;
	v495 = F_bms_copy(m, v471)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L16
	} else {
		goto L186
	}
L185:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	v499 = F_bms_del_member(m, v497, v498)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L16
	} else {
		goto L187
	}
L186:
	;
	v497 = v495
	goto L185
L187:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v491)+8))
	v502 = F_bms_add_member(m, v499, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L16
	} else {
		goto L188
	}
L188:
	;
	v504 = v502
	goto L182
L189:
	;
	goto L178
L190:
	;
	v519 = v514
	goto L192
L191:
	;
	v519 = v471
	goto L192
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+44)) = v519
	v521 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v522 = int32(0)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v522 < v523 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v531 = v522
	v532 = int32(0)
	goto L196
L194:
	;
	v562 = v522
	goto L195
L195:
	;
	v570 = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v306)+152)) = v570
	*(*int64)(unsafe.Add(mBase, uint32(v306)+144)) = v570
	*(*int64)(unsafe.Add(mBase, uint32(v306)+136)) = v570
	*(*int64)(unsafe.Add(mBase, uint32(v306)+128)) = v570
	*(*int32)(unsafe.Add(mBase, uint32(v306)+116)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v306)+108)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v306)+88)) = v570
	*(*int64)(unsafe.Add(mBase, uint32(v306)+80)) = v570
	*(*int64)(unsafe.Add(mBase, uint32(v306)+64)) = v570
	if v562 != 0 {
		goto L209
	} else {
		goto L210
	}
L196:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v526+v532<<(uint(int32(2))%32))))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	v544 = F_bms_is_member(m, v543, v521)
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L16
	} else {
		goto L198
	}
L197:
	;
	v562 = v555
	goto L195
L198:
	;
	if v544 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	if v531 != 0 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	v555 = v531
	goto L201
L201:
	;
	v557 = v532 + int32(1)
	if v557 != v523 {
		v531 = v555
		v532 = v557
		goto L196
	} else {
		goto L208
	}
L202:
	;
	v548 = v531
	goto L204
L203:
	;
	v546 = F_bms_copy(m, v521)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L16
	} else {
		goto L205
	}
L204:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	v550 = F_bms_del_member(m, v548, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L16
	} else {
		goto L206
	}
L205:
	;
	v548 = v546
	goto L204
L206:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v542)+8))
	v553 = F_bms_add_member(m, v550, v552)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L16
	} else {
		goto L207
	}
L207:
	;
	v555 = v553
	goto L201
L208:
	;
	goto L197
L209:
	;
	v588 = v562
	goto L211
L210:
	;
	v588 = v521
	goto L211
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v306)+48)) = v588
	v595 = v306
	goto L6
L212:
	;
	v595 = v591
	goto L6
L213:
	;
	v612 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v49)+32))
	v614 = F_get_rel_name(m, v613)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L16
	} else {
		goto L214
	}
L214:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v612
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_3), v14)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L16
	} else {
		goto L215
	}
L215:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(287), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L16
	} else {
		goto L216
	}
L216:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L217:
	;
	v630 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v49)+32))
	v632 = F_get_rel_name(m, v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L16
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v632
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v630
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_3), v14+int32(16))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L16
	} else {
		goto L219
	}
L219:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(292), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L16
	} else {
		goto L220
	}
L220:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L221:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_4), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L16
	} else {
		goto L222
	}
L222:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(304), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L16
	} else {
		goto L223
	}
L223:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L224:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_4), int32(0))
	mBase = m.M
	v666 = m.ExcPending
	if v666 != 0 {
		goto L16
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(362), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L16
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	F_errmsg_internal(m, int32(_a_F_adjust_appendrel_attrs_mutator_5), int32(0))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L16
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(_a_F_adjust_appendrel_attrs_mutator_1), int32(390), int32(_a_F_adjust_appendrel_attrs_mutator_2))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L16
	} else {
		goto L229
	}
L229:
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
	F_errfinish(m, int32(_a_F_adjust_inherited_attnums_multilevel_1), int32(670), int32(_a_F_adjust_inherited_attnums_multilevel_2))
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
	F_errfinish(m, int32(_a_F_adjust_inherited_attnums_multilevel_1), int32(674), int32(_a_F_adjust_inherited_attnums_multilevel_2))
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
	F_errfinish(m, int32(_a_F_adjust_inherited_attnums_multilevel_1), int32(692), int32(_a_F_adjust_inherited_attnums_multilevel_4))
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
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
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
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
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
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v71
	goto L25
L24:
	;
	goto L25
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v77 == v39 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v71
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
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
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)+212))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+42)))
	if v10 == v8 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0])) = v98
	goto L1
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if int32(0) < v13 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v62 = int32(_a_F_advance_transition_function_0)
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0])) = v66
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = l1
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v69
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)) = uint8(v72)
	*(*uint8)(unsafe.Add(mBase, uint32(v9)+24)) = uint8(v71)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = m.T0[v76].(func(*base.Module, int32) int32)(m, v9)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L16
	} else {
		goto L19
	}
L6:
	;
	v19 = v8
	goto L9
L7:
	;
	goto L8
L8:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)))
	if v37 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v19<<(uint(int32(3))%32))+24)))
	if v26 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v28 = v19 + int32(1)
	if v28 <= v13 {
		v19 = v28
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v40 = int32(_a_F_advance_transition_function_0)
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_advance_transition_function[0])) = v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
	v48 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+184)))
	v49 = F_datumCopy(m, v46, v47, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	if v54 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	return
L17:
	;
	v51 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v51)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v49
	v98 = v41
	goto L2
L18:
	;
	goto L5
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = int32(0)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+187)))
	if v81 != 0 {
		v88 = v77
		goto L20
	} else {
		goto L21
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v88
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v91)
	v98 = v63
	goto L2
L21:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v77 == v82 {
		v88 = v77
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v86 = F_ExecAggCopyTransValue(m, l0, l1, v77, v84, v82, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v88 = v86
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
func F_anycompatible_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_anycompatible_out_0), int32(376), int32(_a_F_anycompatible_out_1), int32(_a_F_anycompatible_out_2), int32(_a_F_anycompatible_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anyelement_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_anyelement_out_0), int32(374), int32(_a_F_anyelement_out_1), int32(_a_F_anyelement_out_2), int32(_a_F_anyelement_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_anyenum_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_anyenum_in_0), int32(194), int32(_a_F_anyenum_in_1), int32(_a_F_anyenum_in_2), int32(_a_F_anyenum_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
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
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+111)))
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
func F_autoprewarm_dump_now(m *base.Module, l0 int32) int32 {
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
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int64
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
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
	v106 = int32(m.ExcTag)
	v107 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v106 == int32(0) {
		goto L25
	} else {
		goto L26
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
	v50 = v12
	v51 = v13
	v53 = v15
	goto L8
L8:
	;
	if v53 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[0])) = v23
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v12
	F_LWLockRegisterTranche(m, v26, int32(_a_F_autoprewarm_dump_now_0))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v13
	F_before_shmem_exit(m, int32(_a_F_autoprewarm_dump_now_1), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[1]))
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[2]))
	goto L12
L12:
	;
	v43 = v8 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v8 + int32(12)
	goto L15
L13:
	;
	v50 = v39
	v51 = v41
	v53 = int32(0)
	goto L8
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v50
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[2])) = v8 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v51
	v64 = F_apw_dump_now(m, int32(0), int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L5
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[1])) = v50
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[2])) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v51
	F_cancel_before_shmem_exit(m, int32(_a_F_autoprewarm_dump_now_1), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v50
	F_cancel_before_shmem_exit(m, int32(_a_F_autoprewarm_dump_now_1), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[1])) = v50
	*(*int32)(unsafe.Add(mBase, _c_F_autoprewarm_dump_now[2])) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v51
	v79 = F_Int64GetDatum(m, base.I64_extend_i32_s(v64))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	m.G0 = v8 + int32(192)
	return v79
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v50
	F_apw_detach_shmem(m, v8, v8)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+188)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v8)+184)) = v50
	F_pg_re_throw(m)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	goto L4
L25:
	;
	v111 = int32(v107)
	m.G0 = v8
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	if v8+int32(12) == v117 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	m.ExcPending = 1
	goto L34
L27:
	;
	if v121 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	v121 = v119
	goto L30
L29:
	;
	v121 = int32(0)
	goto L30
L30:
	;
	goto L27
L31:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v8)+184))
	v12 = v123
	v13 = v122
	v14 = v121
	v15 = v113
	goto L1
L32:
	;
	goto L33
L33:
	;
	F___wasm_longjmp(m, v114, v113)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(0)
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
