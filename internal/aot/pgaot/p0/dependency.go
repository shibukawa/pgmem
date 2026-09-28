package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dependencyLockAndCheckObject(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v94 int32
	_ = v94
	var v111 int32
	_ = v111
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v159 int64
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if l0 != int32(1259) {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = int32(17301504)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = l0
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_dependencyLockAndCheckObject[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v17
		v21 = F_LockHeldByMe(m, v8-int32(-64))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			if v21 != 0 {
				m.G0 = v8 + int32(80)
				return
			} else {
				F_LockDatabaseObject(m, l0, l1, int32(1))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v26 = base.I64_extend_i32_u(l1)
					v27 = F_get_object_catcache_oid(m, l0)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						if v27 != int32(-1) {
							v31 = int64(0)
							v34 = F_SearchSysCacheExists(m, v27, v26, v31, v31, v31)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return
							} else {
								if v34 != 0 {
									m.G0 = v8 + int32(80)
									return
								} else {
									v37 = F_table_open(m, l0, int32(1))
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										v40 = v8 + int32(8)
										v41 = F_get_object_attnum_oid(m, l0)
										mBase = m.M
										v42 = m.ExcPending
										if v42 != 0 {
											return
										} else {
											F_ScanKeyInit(m, v40, v41, int32(3), int32(184), v26)
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return
											} else {
												v47 = F_get_object_oid_index(m, l0)
												mBase = m.M
												v48 = m.ExcPending
												if v48 != 0 {
													return
												} else {
													v49 = int32(1)
													v52 = F_systable_beginscan(m, v37, v47, v49, int32(_a_F_dependencyLockAndCheckObject_0), v49, v40)
													mBase = m.M
													v53 = m.ExcPending
													if v53 != 0 {
														return
													} else {
														v54 = F_systable_getnext(m, v52)
														mBase = m.M
														v55 = m.ExcPending
														if v55 != 0 {
															return
														} else {
															if v54 == int32(0) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return
																} else {
																	F_errcode(m, int32(67137668))
																	mBase = m.M
																	v179 = m.ExcPending
																	if v179 != 0 {
																		return
																	} else {
																		v180 = F_get_object_class_descr(m, l0)
																		mBase = m.M
																		v181 = m.ExcPending
																		if v181 != 0 {
																			return
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v180
																			F_errmsg(m, int32(_a_F_dependencyLockAndCheckObject_1), v8)
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
																				return
																			} else {
																				F_errfinish(m, int32(_a_F_dependencyLockAndCheckObject_2), int32(813), int32(_a_F_dependencyLockAndCheckObject_3))
																				mBase = m.M
																				v190 = m.ExcPending
																				if v190 != 0 {
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
																F_systable_endscan(m, v52)
																mBase = m.M
																v59 = m.ExcPending
																if v59 != 0 {
																	return
																} else {
																	F_relation_close(m, v37, int32(1))
																	mBase = m.M
																	v62 = m.ExcPending
																	if v62 != 0 {
																		return
																	} else {
																		m.G0 = v8 + int32(80)
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
							v37 = F_table_open(m, l0, int32(1))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v40 = v8 + int32(8)
								v41 = F_get_object_attnum_oid(m, l0)
								mBase = m.M
								v42 = m.ExcPending
								if v42 != 0 {
									return
								} else {
									F_ScanKeyInit(m, v40, v41, int32(3), int32(184), v26)
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
										return
									} else {
										v47 = F_get_object_oid_index(m, l0)
										mBase = m.M
										v48 = m.ExcPending
										if v48 != 0 {
											return
										} else {
											v49 = int32(1)
											v52 = F_systable_beginscan(m, v37, v47, v49, int32(_a_F_dependencyLockAndCheckObject_0), v49, v40)
											mBase = m.M
											v53 = m.ExcPending
											if v53 != 0 {
												return
											} else {
												v54 = F_systable_getnext(m, v52)
												mBase = m.M
												v55 = m.ExcPending
												if v55 != 0 {
													return
												} else {
													if v54 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v176 = m.ExcPending
														if v176 != 0 {
															return
														} else {
															F_errcode(m, int32(67137668))
															mBase = m.M
															v179 = m.ExcPending
															if v179 != 0 {
																return
															} else {
																v180 = F_get_object_class_descr(m, l0)
																mBase = m.M
																v181 = m.ExcPending
																if v181 != 0 {
																	return
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v180
																	F_errmsg(m, int32(_a_F_dependencyLockAndCheckObject_1), v8)
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
																		return
																	} else {
																		F_errfinish(m, int32(_a_F_dependencyLockAndCheckObject_2), int32(813), int32(_a_F_dependencyLockAndCheckObject_3))
																		mBase = m.M
																		v190 = m.ExcPending
																		if v190 != 0 {
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
														F_systable_endscan(m, v52)
														mBase = m.M
														v59 = m.ExcPending
														if v59 != 0 {
															return
														} else {
															F_relation_close(m, v37, int32(1))
															mBase = m.M
															v62 = m.ExcPending
															if v62 != 0 {
																return
															} else {
																m.G0 = v8 + int32(80)
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
		v63 = m.G0
		v65 = v63 - int32(16)
		m.G0 = v65
		v69 = int32(1)
		if l1 <= int32(3591) {
			if l1 <= int32(2670) {
				switch l1 - int32(1213) {
				case 0, 1, 19, 20, 47, 48, 49:
					v140 = v69
				case 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46:
					v140 = int32(0)
				default:
					if base.Ui32(int32(2)) <= base.Ui32(l1-int32(2396)) {
						v140 = int32(0)
					} else {
						v140 = v69
					}
				}
			} else {
				v81 = l1 - int32(2671)
				if base.B2i32(base.Ui32(int32(27)) < base.Ui32(v81))|base.B2i32(int32(1)<<(uint(v81)%32)&int32(226492515) == int32(0)) != 0 {
					if base.B2i32(base.Ui32(l1-int32(2964)) < base.Ui32(int32(4)))|base.B2i32(base.Ui32(l1-int32(2846)) < base.Ui32(int32(2))) != 0 {
						v140 = v69
					} else {
						v140 = int32(0)
					}
				} else {
					v140 = v69
				}
			}
		} else {
			if l1 <= int32(_a_F_dependencyLockAndCheckObject_4) {
				v94 = l1 - int32(_a_F_dependencyLockAndCheckObject_5)
				if base.B2i32(base.Ui32(int32(9)) < base.Ui32(v94))|base.B2i32(int32(1)<<(uint(v94)%32)&int32(963) == int32(0)) != 0 {
					if base.Ui32(l1-int32(3592)) < base.Ui32(int32(2)) {
						v140 = v69
					} else {
						if base.Ui32(int32(2)) <= base.Ui32(l1-int32(4060)) {
							v140 = int32(0)
						} else {
							v140 = v69
						}
					}
				} else {
					v140 = v69
				}
			} else {
				switch l1 - int32(_a_F_dependencyLockAndCheckObject_6) {
				case 0, 1, 2, 3, 4, 59, 60:
					v140 = v69
				case 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58:
					v140 = int32(0)
				default:
					if base.Ui32(l1-int32(_a_F_dependencyLockAndCheckObject_7)) < base.Ui32(int32(3)) {
						v140 = v69
					} else {
						v111 = l1 - int32(_a_F_dependencyLockAndCheckObject_8)
						if base.Ui32(int32(15)) < base.Ui32(v111) {
							v140 = int32(0)
						} else {
							if int32(1)<<(uint(v111)%32)&int32(_a_F_dependencyLockAndCheckObject_9) != 0 {
								v140 = v69
							} else {
								v140 = int32(0)
							}
						}
					}
				}
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = int64(72057594037927936)
		*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = l1
		v146 = *(*int32)(unsafe.Add(mBase, _c_F_dependencyLockAndCheckObject[0]))
		if v140 != 0 {
			v147 = int32(0)
		} else {
			v147 = v146
		}
		*(*int32)(unsafe.Add(mBase, uint32(v65))) = v147
		v149 = F_LockHeldByMe(m, v65)
		mBase = m.M
		v150 = m.ExcPending
		if v150 != 0 {
			return
		} else {
			m.G0 = v65 + int32(16)
			if v149 != 0 {
				m.G0 = v8 + int32(80)
				return
			} else {
				F_LockRelationOid(m, l1, int32(1))
				mBase = m.M
				v156 = m.ExcPending
				if v156 != 0 {
					return
				} else {
					v159 = int64(0)
					v162 = F_SearchSysCacheExists(m, int32(57), base.I64_extend_i32_u(l1), v159, v159, v159)
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return
					} else {
						if v162 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return
							} else {
								F_errcode(m, int32(67137668))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return
								} else {
									F_errmsg(m, int32(_a_F_dependencyLockAndCheckObject_10), int32(0))
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_dependencyLockAndCheckObject_2), int32(844), int32(_a_F_dependencyLockAndCheckObject_3))
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
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
							m.G0 = v8 + int32(80)
							return
						}
					}
				}
			}
		}
	}
}
func F_dependency_is_compatible_expression(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
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
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v14 == int32(320) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v264
L2:
	;
	if v188 == int32(27) {
		goto L64
	} else {
		goto L65
	}
L3:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v187 = v184
	v188 = v186
	goto L2
L4:
	;
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159)+20)))
	if v161 != int32(1) {
		v264 = v4
		goto L1
	} else {
		goto L57
	}
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v17 != 0 {
		v264 = v4
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v72 = l0
	v73 = v14
	goto L7
L7:
	;
	switch v73 - int32(17) {
	case 0:
		goto L30
	default:
		v187 = v72
		v188 = v73
		goto L2
	case 3:
		v159 = v72
		goto L4
	case 4:
		goto L29
	}
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = int32(0)
	if v18 == v19 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v64 != int32(1) {
		v264 = v4
		goto L1
	} else {
		goto L25
	}
L10:
	;
	v64 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v27 = int32(1)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v28 <= v27 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v31 = v27
	goto L15
L14:
	;
	v31 = v28
	goto L15
L15:
	;
	v35 = int32(0)
	v37 = v19
	goto L16
L16:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v18+int32(8)+v35<<(uint(int32(2))%32))))
	if v44 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v64 = v56
	goto L9
L18:
	;
	goto L17
L19:
	;
	v45 = int32(2)
	if v37 != 0 {
		v56 = v45
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v51 = v37
	goto L21
L21:
	;
	v53 = v35 + int32(1)
	if v53 != v31 {
		v35 = v53
		v37 = v51
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v46 = int32(1)
	if base.Ui32(v46) < base.Ui32(base.I32_popcnt(v44)) {
		v56 = v45
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v51 = v46
	goto L21
L24:
	;
	v56 = v51
	goto L18
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v159 = int32(0)
	goto L4
L27:
	;
	goto L28
L28:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v72 = v67
	v73 = v71
	goto L7
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	switch v108 - int32(1) {
	case 0:
		goto L43
	case 1:
		goto L42
	default:
		v184 = v72
		goto L3
	}
L30:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	if v76 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
	if v79 != int32(2) {
		v264 = v4
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v84 = F_is_pseudo_constant_clause(m, v83)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	if v84 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v93 = F_is_pseudo_constant_clause(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v101 = v89
	goto L37
L37:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v104 = F_get_oprrest(m, v102)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	if v93 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	v101 = v98 + int32(4)
	goto L37
L40:
	;
	if v104 == int32(101) {
		v184 = v103
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v264 = v4
	goto L1
L42:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	v184 = v158
	goto L3
L43:
	;
	v111 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v111
	v114 = int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	if v115 == v111 {
		v264 = v114
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v118 <= int32(0) {
		v264 = v114
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v124 = v111
	goto L46
L46:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+v124<<(uint(int32(2))%32))))
	v139 = F_dependency_is_compatible_expression(m, v136, l1, v12+int32(12))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L33
	} else {
		goto L49
	}
L47:
	;
	v264 = int32(0)
	goto L1
L48:
	;
	goto L47
L49:
	;
	if v139 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v144 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v146 = v144
	goto L53
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v143
	v146 = v143
	goto L53
L53:
	;
	v147 = F_equal(m, v143, v146)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L33
	} else {
		goto L54
	}
L54:
	;
	if v147 == int32(0) {
		v264 = v147
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v152 = v124 + int32(1)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v152 < v153 {
		v124 = v152
		goto L46
	} else {
		goto L56
	}
L56:
	;
	v264 = v147
	goto L1
L57:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	if v164 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	if v167 != int32(2) {
		v264 = v4
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+4))
	v172 = F_is_pseudo_constant_clause(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L33
	} else {
		goto L60
	}
L60:
	;
	if v172 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v159)+28))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v177)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	v180 = F_get_oprrest(m, v176)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	if v180 != int32(101) {
		v264 = v4
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v184 = v179
	goto L3
L64:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v192 = v191
	goto L66
L65:
	;
	v192 = v187
	goto L66
L66:
	;
	if l1 == int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v195 <= int32(0) {
		v264 = v4
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v205 = v4
	goto L69
L69:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207+v205<<(uint(int32(2))%32))))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211)+16)))
	if v212 != int32(102) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v264 = v4
	goto L1
L71:
	;
	v256 = v205 + int32(1)
	v257 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v256 < v257 {
		v205 = v256
		goto L69
	} else {
		goto L82
	}
L72:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	if v215 == int32(0) {
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v218 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v219 <= v218 {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v225 = v218
	goto L75
L75:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v215)+12))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v231+v225<<(uint(int32(2))%32))))
	v236 = F_equal(m, v192, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L33
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v235
	v264 = int32(1)
	goto L1
L77:
	;
	if v236 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v241 = v225 + int32(1)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	if v241 < v242 {
		v225 = v241
		goto L75
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	goto L76
L81:
	;
	goto L71
L82:
	;
	goto L70
}
func F_recordDependencyOnExpr(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int64
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	v10 = m.G0
	v11 = int32(16)
	v12 = v10 - v11
	m.G0 = v12
	v15 = F_palloc(m, v11)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(137438953472)
		v21 = F_palloc_mul(m, int32(12), int32(32))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v15
			v29 = int32(1)
			v31 = F_list_make1_impl(m, v29, v12)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v31
				v36 = F_find_expr_references_walker(m, l1, v12+int32(8))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					if v39 < int32(2) {
						v105 = v38
						v108 = v39
						F_recordMultipleDependencies(m, l0, v105, v108, int32(110))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return
						} else {
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
							F_pfree(m, v115)
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
								if v118 != 0 {
									F_pfree(m, v118)
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return
									} else {
										F_pfree(m, v15)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
											return
										}
									}
								} else {
									F_pfree(m, v15)
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
										return
									} else {
										m.G0 = v12 + int32(16)
										return
									}
								}
							}
						}
					} else {
						F_pg_qsort(m, v38, v39, int32(12), int32(497))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
							if int32(2) <= v47 {
								v53 = v46
								v56 = v29
								v57 = int32(1)
								for {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
									v64 = v61 + v57*int32(12)
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
									if v60 != v65 {
										v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v76
										v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
										*(*int64)(unsafe.Add(mBase, uint32(v53)+12)) = v78
										v84 = v53 + int32(12)
										v85 = v56 + int32(1)
									} else {
										v67 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
										if v67 != v68 {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v76
											v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
											*(*int64)(unsafe.Add(mBase, uint32(v53)+12)) = v78
											v84 = v53 + int32(12)
											v85 = v56 + int32(1)
										} else {
											v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
											v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
											if v70 == v71 {
												v84 = v53
												v85 = v56
											} else {
												if v70 != 0 {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v53)+20)) = v76
													v78 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
													*(*int64)(unsafe.Add(mBase, uint32(v53)+12)) = v78
													v84 = v53 + int32(12)
													v85 = v56 + int32(1)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v53)+8)) = v71
													v84 = v53
													v85 = v56
												}
											}
										}
									}
									v89 = v57 + int32(1)
									v90 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
									if v89 < v90 {
										v53 = v84
										v56 = v85
										v57 = v89
										continue
									} else {
										break
									}
									break
								}
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
								v95 = v92
								v98 = v85
							} else {
								v95 = v46
								v98 = v29
							}
							*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v98
							v105 = v95
							v108 = v98
							F_recordMultipleDependencies(m, l0, v105, v108, int32(110))
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return
							} else {
								v115 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
								F_pfree(m, v115)
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									v118 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
									if v118 != 0 {
										F_pfree(m, v118)
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return
										} else {
											F_pfree(m, v15)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return
											} else {
												m.G0 = v12 + int32(16)
												return
											}
										}
									} else {
										F_pfree(m, v15)
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											m.G0 = v12 + int32(16)
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
