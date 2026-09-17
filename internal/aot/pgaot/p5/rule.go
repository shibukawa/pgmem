package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_EnableDisableRule(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	v3 = l2
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v15 = F_table_open(m, int32(2618), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		v18 = F_SearchSysCacheCopy(m, int32(60), v12, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			if v18 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_EnableDisableRule[0]))
				v27 = F_object_ownercheck(m, int32(1259), v24, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					if v27 == int32(0) {
						v32 = F_get_rel_relkind(m, v24)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							switch v32 - int32(73) {
							case 0, 32:
								v43 = int32(20)
								v45 = v43
							default:
								v43 = int32(41)
								v45 = v43
							case 10:
								v45 = int32(37)
							case 29:
								v45 = int32(18)
							case 36:
								v45 = int32(23)
							case 45:
								v45 = int32(51)
							}
							v46 = F_get_rel_name(m, v24)
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_aclcheck_error(m, int32(2), v45, v46)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+73)))
									if v50 != v3&int32(255) {
										*(*uint8)(unsafe.Add(mBase, uint32(v23)+73)) = uint8(v3)
										F_CatalogTupleUpdate(m, v15, v18+int32(4), v18)
										mBase = m.M
										v58 = m.ExcPending
										if v58 != 0 {
											return
										} else {
											v60 = *(*int32)(unsafe.Add(mBase, _c_F_EnableDisableRule[1]))
											if v60 != 0 {
												v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
												v63 = int32(0)
												F_RunObjectPostAlterHook(m, int32(2618), v62, v63, v63, v63)
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return
												} else {
													F_pfree(m, v18)
													mBase = m.M
													v69 = m.ExcPending
													if v69 != 0 {
														return
													} else {
														F_relation_close(m, v15, int32(3))
														mBase = m.M
														v72 = m.ExcPending
														if v72 != 0 {
															return
														} else {
															if v3&int32(255) != v50 {
																F_CacheInvalidateRelcache(m, l0)
																mBase = m.M
																v77 = m.ExcPending
																if v77 != 0 {
																	return
																} else {
																	m.G0 = v10 + int32(16)
																	return
																}
															} else {
																m.G0 = v10 + int32(16)
																return
															}
														}
													}
												}
											} else {
												F_pfree(m, v18)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													F_relation_close(m, v15, int32(3))
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return
													} else {
														if v3&int32(255) != v50 {
															F_CacheInvalidateRelcache(m, l0)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return
															} else {
																m.G0 = v10 + int32(16)
																return
															}
														} else {
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											}
										}
									} else {
										v60 = *(*int32)(unsafe.Add(mBase, _c_F_EnableDisableRule[1]))
										if v60 != 0 {
											v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
											v63 = int32(0)
											F_RunObjectPostAlterHook(m, int32(2618), v62, v63, v63, v63)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return
											} else {
												F_pfree(m, v18)
												mBase = m.M
												v69 = m.ExcPending
												if v69 != 0 {
													return
												} else {
													F_relation_close(m, v15, int32(3))
													mBase = m.M
													v72 = m.ExcPending
													if v72 != 0 {
														return
													} else {
														if v3&int32(255) != v50 {
															F_CacheInvalidateRelcache(m, l0)
															mBase = m.M
															v77 = m.ExcPending
															if v77 != 0 {
																return
															} else {
																m.G0 = v10 + int32(16)
																return
															}
														} else {
															m.G0 = v10 + int32(16)
															return
														}
													}
												}
											}
										} else {
											F_pfree(m, v18)
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return
											} else {
												F_relation_close(m, v15, int32(3))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return
												} else {
													if v3&int32(255) != v50 {
														F_CacheInvalidateRelcache(m, l0)
														mBase = m.M
														v77 = m.ExcPending
														if v77 != 0 {
															return
														} else {
															m.G0 = v10 + int32(16)
															return
														}
													} else {
														m.G0 = v10 + int32(16)
														return
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+73)))
						if v50 != v3&int32(255) {
							*(*uint8)(unsafe.Add(mBase, uint32(v23)+73)) = uint8(v3)
							F_CatalogTupleUpdate(m, v15, v18+int32(4), v18)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, _c_F_EnableDisableRule[1]))
								if v60 != 0 {
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
									v63 = int32(0)
									F_RunObjectPostAlterHook(m, int32(2618), v62, v63, v63, v63)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return
									} else {
										F_pfree(m, v18)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											F_relation_close(m, v15, int32(3))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												if v3&int32(255) != v50 {
													F_CacheInvalidateRelcache(m, l0)
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
														return
													} else {
														m.G0 = v10 + int32(16)
														return
													}
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											}
										}
									}
								} else {
									F_pfree(m, v18)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_relation_close(m, v15, int32(3))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											if v3&int32(255) != v50 {
												F_CacheInvalidateRelcache(m, l0)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											} else {
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								}
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, _c_F_EnableDisableRule[1]))
							if v60 != 0 {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
								v63 = int32(0)
								F_RunObjectPostAlterHook(m, int32(2618), v62, v63, v63, v63)
								mBase = m.M
								v67 = m.ExcPending
								if v67 != 0 {
									return
								} else {
									F_pfree(m, v18)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_relation_close(m, v15, int32(3))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											if v3&int32(255) != v50 {
												F_CacheInvalidateRelcache(m, l0)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return
												} else {
													m.G0 = v10 + int32(16)
													return
												}
											} else {
												m.G0 = v10 + int32(16)
												return
											}
										}
									}
								}
							} else {
								F_pfree(m, v18)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									F_relation_close(m, v15, int32(3))
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										if v3&int32(255) != v50 {
											F_CacheInvalidateRelcache(m, l0)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												m.G0 = v10 + int32(16)
												return
											}
										} else {
											m.G0 = v10 + int32(16)
											return
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
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						v88 = F_get_rel_name(m, v12)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v88
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
							F_errmsg(m, int32(_a_F_EnableDisableRule_0), v10)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_EnableDisableRule_1), int32(712), int32(_a_F_EnableDisableRule_2))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
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
func F_checkRuleResultList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
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
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
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
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	if l0 == v5 {
		v119 = v5
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L28
	} else {
		goto L80
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L28
	} else {
		goto L67
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L28
	} else {
		goto L54
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L28
	} else {
		goto L49
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L28
	} else {
		goto L42
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L28
	} else {
		goto L35
	}
L7:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v119 != v126 {
		goto L1
	} else {
		goto L34
	}
L8:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 <= int32(0) {
		v119 = v5
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v27 = v5
	v31 = v5
	goto L10
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v31<<(uint(int32(2))%32))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+26)))
	if v39 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v119 = v104
	goto L7
L12:
	;
	v43 = v27 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v44 < v43 {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	v104 = v27
	goto L14
L14:
	;
	v111 = v31 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v111 < v112 {
		v27 = v104
		v31 = v111
		goto L10
	} else {
		goto L33
	}
L15:
	;
	v51 = l1 + v44<<(uint(int32(4))%32) + v27*int32(100)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+111)))
	if v52 == int32(1) {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v56 = v51 + int32(24)
	if l3 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if base.B2i32(v60 == int32(0))|base.B2i32(v60 != v63) != 0 {
		v81 = v60
		v82 = v63
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v85 = F_exprType(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	if v81-v82 != 0 {
		goto L4
	} else {
		goto L27
	}
L21:
	;
	goto L20
L22:
	;
	v66 = v57
	v67 = v56
	goto L23
L23:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)))
	if v71 == int32(0) {
		v81 = v71
		v82 = v70
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v81 = v71
	v82 = v70
	goto L21
L25:
	;
	v74 = int32(1)
	if v71 == v70 {
		v66 = v66 + v74
		v67 = v67 + v74
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	goto L19
L28:
	;
	return
L29:
	;
	v88 = v51 + int32(20)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+68))
	if v85 != v89 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v92 = F_exprTypmod(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+76))
	v96 = int32(-1)
	if base.B2i32(base.B2i32(v92 == v94)|base.B2i32(v92 == v96) == int32(0))&base.B2i32(v94 != v96) != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v104 = v43
	goto L14
L33:
	;
	goto L11
L34:
	;
	m.G0 = v15 + int32(96)
	return
L35:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L28
	} else {
		goto L36
	}
L36:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v140 = int32(_a_F_checkRuleResultList_0)
	goto L39
L38:
	;
	v140 = int32(_a_F_checkRuleResultList_1)
	goto L39
L39:
	;
	F_errmsg(m, v140, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L28
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_checkRuleResultList_2), int32(533), int32(_a_F_checkRuleResultList_3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L28
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L28
	} else {
		goto L43
	}
L43:
	;
	if l2 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v158 = int32(_a_F_checkRuleResultList_4)
	goto L46
L45:
	;
	v158 = int32(_a_F_checkRuleResultList_5)
	goto L46
L46:
	;
	F_errmsg(m, v158, int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L28
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_checkRuleResultList_2), int32(561), int32(_a_F_checkRuleResultList_3))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L28
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+84)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = v43
	F_errmsg(m, int32(_a_F_checkRuleResultList_6), v15+int32(80))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L28
	} else {
		goto L51
	}
L51:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = v181
	F_errdetail(m, int32(_a_F_checkRuleResultList_7), v15-int32(-64))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L28
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_checkRuleResultList_2), int32(570), int32(_a_F_checkRuleResultList_3))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L28
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
	F_errcode(m, int32(117833860))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L28
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v43
	if l2 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v204 = int32(_a_F_checkRuleResultList_8)
	goto L58
L57:
	;
	v204 = int32(_a_F_checkRuleResultList_9)
	goto L58
L58:
	;
	F_errmsg(m, v204, v15+int32(48))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L28
	} else {
		goto L59
	}
L59:
	;
	v209 = F_format_type_be(m, v85)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L28
	} else {
		goto L60
	}
L60:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v88)+68))
	v212 = F_format_type_be(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L28
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v209
	if l2 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v218 = int32(_a_F_checkRuleResultList_10)
	goto L64
L63:
	;
	v218 = int32(_a_F_checkRuleResultList_11)
	goto L64
L64:
	;
	F_errdetail(m, v218, v15+int32(32))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L28
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_checkRuleResultList_2), int32(588), int32(_a_F_checkRuleResultList_3))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L28
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
	F_errcode(m, int32(117833860))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L28
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v43
	if l2 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v239 = int32(_a_F_checkRuleResultList_12)
	goto L71
L70:
	;
	v239 = int32(_a_F_checkRuleResultList_13)
	goto L71
L71:
	;
	F_errmsg(m, v239, v15+int32(16))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L28
	} else {
		goto L72
	}
L72:
	;
	v244 = F_format_type_with_typemod(m, v85, v92)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L28
	} else {
		goto L73
	}
L73:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v88)+68))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v88)+76))
	v248 = F_format_type_with_typemod(m, v246, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L28
	} else {
		goto L74
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v244
	if l2 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v254 = int32(_a_F_checkRuleResultList_10)
	goto L77
L76:
	;
	v254 = int32(_a_F_checkRuleResultList_11)
	goto L77
L77:
	;
	F_errdetail(m, v254, v15)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L28
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_checkRuleResultList_2), int32(614), int32(_a_F_checkRuleResultList_3))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L28
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L80:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L28
	} else {
		goto L81
	}
L81:
	;
	if l2 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v271 = int32(_a_F_checkRuleResultList_14)
	goto L84
L83:
	;
	v271 = int32(_a_F_checkRuleResultList_15)
	goto L84
L84:
	;
	F_errmsg(m, v271, int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L28
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_checkRuleResultList_2), int32(622), int32(_a_F_checkRuleResultList_3))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L28
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_setRuleCheckAsUser(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l1
	if l0 == int32(0) {
		m.G0 = v6 + int32(16)
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v11 == int32(67) {
			F_setRuleCheckAsUser_Query(m, l0, l1)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		} else {
			v19 = F_expression_tree_walker_impl(m, l0, int32(1039), v6+int32(12))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return
			} else {
				m.G0 = v6 + int32(16)
				return
			}
		}
	}
}
