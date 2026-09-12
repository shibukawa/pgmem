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
				v26 = *(*int32)(unsafe.Add(mBase, _consts[239]))
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
								v45 = int32(20)
							default:
								v43 = int32(41)
								v45 = v43
							case 10:
								v45 = int32(37)
							case 29:
								v43 = int32(18)
								v45 = v43
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
											v60 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
														F_sequence_close(m, v15, int32(3))
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
													F_sequence_close(m, v15, int32(3))
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
										v60 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
													F_sequence_close(m, v15, int32(3))
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
												F_sequence_close(m, v15, int32(3))
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
								v60 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
											F_sequence_close(m, v15, int32(3))
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
										F_sequence_close(m, v15, int32(3))
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
							v60 = *(*int32)(unsafe.Add(mBase, _consts[230]))
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
										F_sequence_close(m, v15, int32(3))
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
									F_sequence_close(m, v15, int32(3))
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
							F_errmsg(m, int32(70813), v10)
							mBase = m.M
							v94 = m.ExcPending
							if v94 != 0 {
								return
							} else {
								F_errfinish(m, int32(493391), int32(712), int32(378929))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
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
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v253 != v245 {
		goto L85
	} else {
		goto L86
	}
L2:
	;
	v28 = v5
	v29 = v5
	goto L11
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v18 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v245 = v5
	goto L1
L6:
	;
	goto L5
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L31
	} else {
		goto L72
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L31
	} else {
		goto L59
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L31
	} else {
		goto L54
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L31
	} else {
		goto L47
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v29<<(uint(int32(2))%32))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+26)))
	if v41 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L31
	} else {
		goto L40
	}
L13:
	;
	goto L12
L14:
	;
	v45 = v28 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v46 < v45 {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	v99 = v28
	goto L16
L16:
	;
	v106 = v29 + int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v106 < v107 {
		v28 = v99
		v29 = v106
		goto L11
	} else {
		goto L39
	}
L17:
	;
	v53 = l1 + int32(20) + v46<<(uint(int32(4))%32) + v28*int32(100)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+91)))
	if v54 == int32(1) {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	v58 = v53 + int32(4)
	if l3 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v63 == int32(0) {
		v82 = v62
		v83 = v63
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v86 = F_exprType(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	if v83-v82 != 0 {
		goto L9
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	if v62 != v63 {
		v82 = v62
		v83 = v63
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v67 = v59
	v68 = v58
	goto L26
L26:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v72 == int32(0) {
		v82 = v71
		v83 = v72
		goto L23
	} else {
		goto L28
	}
L27:
	;
	v82 = v71
	v83 = v72
	goto L23
L28:
	;
	v75 = int32(1)
	if v71 == v72 {
		v67 = v67 + v75
		v68 = v68 + v75
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L21
L31:
	;
	return
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	if v86 != v88 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	v91 = F_exprTypmod(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v99 = v45
	goto L16
L35:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
	if v91 == v93 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	if v91 == int32(-1) {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	if v93 != int32(-1) {
		goto L7
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v245 = v99
	goto L1
L40:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L31
	} else {
		goto L41
	}
L41:
	;
	if l2 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v118 = int32(165809)
	goto L44
L43:
	;
	v118 = int32(165856)
	goto L44
L44:
	;
	F_errmsg(m, v118, int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(493391), int32(533), int32(75468))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L31
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	if l2 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v136 = int32(32089)
	goto L51
L50:
	;
	v136 = int32(147070)
	goto L51
L51:
	;
	F_errmsg(m, v136, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(493391), int32(561), int32(75468))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L31
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
	v151 = m.ExcPending
	if v151 != 0 {
		goto L31
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v45
	F_errmsg(m, int32(689513), v16+int32(80))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L31
	} else {
		goto L56
	}
L56:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v159
	F_errdetail(m, int32(646270), v16-int32(-64))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L31
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(493391), int32(570), int32(75468))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L31
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L31
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v45
	if l2 != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v182 = int32(689385)
	goto L63
L62:
	;
	v182 = int32(689451)
	goto L63
L63:
	;
	F_errmsg(m, v182, v16+int32(48))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L31
	} else {
		goto L64
	}
L64:
	;
	v187 = F_format_type_be(m, v86)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L31
	} else {
		goto L65
	}
L65:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	v190 = F_format_type_be(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L31
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v190
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v187
	if l2 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v196 = int32(584107)
	goto L69
L68:
	;
	v196 = int32(584049)
	goto L69
L69:
	;
	F_errdetail(m, v196, v16+int32(32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L31
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(493391), int32(588), int32(75468))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L31
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
	F_errcode(m, int32(117833860))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L31
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v45
	if l2 != 0 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v217 = int32(689257)
	goto L76
L75:
	;
	v217 = int32(689323)
	goto L76
L76:
	;
	F_errmsg(m, v217, v16+int32(16))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L31
	} else {
		goto L77
	}
L77:
	;
	v222 = F_format_type_with_typemod(m, v86, v91)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L31
	} else {
		goto L78
	}
L78:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v53)+68))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
	v226 = F_format_type_with_typemod(m, v224, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L31
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v222
	if l2 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v232 = int32(584107)
	goto L82
L81:
	;
	v232 = int32(584049)
	goto L82
L82:
	;
	F_errdetail(m, v232, v16)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L31
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(493391), int32(614), int32(75468))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L31
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L31
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	m.G0 = v16 + int32(96)
	return
L88:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L31
	} else {
		goto L89
	}
L89:
	;
	if l2 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v264 = int32(165892)
	goto L92
L91:
	;
	v264 = int32(165938)
	goto L92
L92:
	;
	F_errmsg(m, v264, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L31
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(493391), int32(622), int32(75468))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L31
	} else {
		goto L94
	}
L94:
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
