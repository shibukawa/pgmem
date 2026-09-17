package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValidateJoinEstimator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
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
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(2281)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(90194315497)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(111669151977)
	v17 = v5 + int32(-32)
	v19 = F_LookupFuncName(m, l0, int32(5), v17, int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v25 = F_LookupFuncName(m, l0, int32(4), v17, int32(1))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				if v25 == int32(0) {
					v55 = v19
					v56 = F_get_func_rettype(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						if v56 == int32(701) {
							if base.Ui32(int32(_a_F_ValidateJoinEstimator_0)) <= base.Ui32(v55) {
								v62 = F_superuser(m)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									if v62 != 0 {
										m.G0 = v7 - int32(-64)
										return v55
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_ValidateJoinEstimator_1), int32(0))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_ValidateJoinEstimator_2), int32(379), int32(_a_F_ValidateJoinEstimator_3))
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
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateJoinEstimator[0]))
								v84 = F_object_aclcheck(m, int32(1255), v55, v82, int64(128))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									if v84 == int32(0) {
										m.G0 = v7 - int32(-64)
										return v55
									} else {
										v89 = F_NameListToString(m, l0)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											F_aclcheck_error(m, v84, int32(19), v89)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												m.G0 = v7 - int32(-64)
												return v55
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v105 = F_NameListToString(m, l0)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_ValidateJoinEstimator_4)
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v105
										F_errmsg(m, int32(_a_F_ValidateJoinEstimator_5), v7)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ValidateJoinEstimator_2), int32(371), int32(_a_F_ValidateJoinEstimator_3))
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
							}
						}
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(84439172))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = F_NameListToString(m, l0)
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v36
								F_errmsg(m, int32(_a_F_ValidateJoinEstimator_6), v5+int32(-48))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_ValidateJoinEstimator_2), int32(356), int32(_a_F_ValidateJoinEstimator_3))
									mBase = m.M
									v48 = m.ExcPending
									if v48 != 0 {
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
				if v25 != 0 {
					v55 = v25
					v56 = F_get_func_rettype(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						if v56 == int32(701) {
							if base.Ui32(int32(_a_F_ValidateJoinEstimator_0)) <= base.Ui32(v55) {
								v62 = F_superuser(m)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									if v62 != 0 {
										m.G0 = v7 - int32(-64)
										return v55
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v67 = m.ExcPending
										if v67 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_ValidateJoinEstimator_1), int32(0))
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_ValidateJoinEstimator_2), int32(379), int32(_a_F_ValidateJoinEstimator_3))
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
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateJoinEstimator[0]))
								v84 = F_object_aclcheck(m, int32(1255), v55, v82, int64(128))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									if v84 == int32(0) {
										m.G0 = v7 - int32(-64)
										return v55
									} else {
										v89 = F_NameListToString(m, l0)
										mBase = m.M
										v90 = m.ExcPending
										if v90 != 0 {
											return int32(0)
										} else {
											F_aclcheck_error(m, v84, int32(19), v89)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												m.G0 = v7 - int32(-64)
												return v55
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									v105 = F_NameListToString(m, l0)
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_ValidateJoinEstimator_4)
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v105
										F_errmsg(m, int32(_a_F_ValidateJoinEstimator_5), v7)
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_ValidateJoinEstimator_2), int32(371), int32(_a_F_ValidateJoinEstimator_3))
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
							}
						}
					}
				} else {
					v53 = F_LookupFuncName(m, l0, int32(5), v5+int32(-32), int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v55 = v53
						v56 = F_get_func_rettype(m, v55)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							if v56 == int32(701) {
								if base.Ui32(int32(_a_F_ValidateJoinEstimator_0)) <= base.Ui32(v55) {
									v62 = F_superuser(m)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return int32(0)
									} else {
										if v62 != 0 {
											m.G0 = v7 - int32(-64)
											return v55
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(16797828))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_ValidateJoinEstimator_1), int32(0))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_ValidateJoinEstimator_2), int32(379), int32(_a_F_ValidateJoinEstimator_3))
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
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateJoinEstimator[0]))
									v84 = F_object_aclcheck(m, int32(1255), v55, v82, int64(128))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
										return int32(0)
									} else {
										if v84 == int32(0) {
											m.G0 = v7 - int32(-64)
											return v55
										} else {
											v89 = F_NameListToString(m, l0)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												F_aclcheck_error(m, v84, int32(19), v89)
												mBase = m.M
												v92 = m.ExcPending
												if v92 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 - int32(-64)
													return v55
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										v105 = F_NameListToString(m, l0)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_ValidateJoinEstimator_4)
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v105
											F_errmsg(m, int32(_a_F_ValidateJoinEstimator_5), v7)
											mBase = m.M
											v112 = m.ExcPending
											if v112 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_ValidateJoinEstimator_2), int32(371), int32(_a_F_ValidateJoinEstimator_3))
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
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_find_join_rel(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	v3 = int32(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v13 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 - int32(-64)
	return v192
L2:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v109 == int32(0) {
		v192 = v3
		goto L1
	} else {
		goto L20
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v16 == int32(0) {
		v192 = v3
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v87 = v13
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	v94 = int32(0)
	v96 = F_hash_search(m, v87, v9+int32(-48), v94, v94)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L18
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v19 < int32(33) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = int32(893)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = int32(894)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = int64(34359738372)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_find_join_rel[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v30
	v37 = F_hash_create(m, int32(_a_F_find_join_rel_0), int32(256), v9+int32(-48), int32(1224))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v41 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v37
	if v37 == int32(0) {
		goto L2
	} else {
		goto L17
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v44 <= int32(0) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v50 = int32(0)
	goto L13
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v50<<(uint(int32(2))%32))))
	v65 = F_hash_search(m, v37, v59+int32(8), int32(1), v9+int32(-49))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	goto L10
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+4)) = v59
	v69 = v50 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v69 < v70 {
		v50 = v69
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v87 = v37
	goto L5
L18:
	;
	if v96 == int32(0) {
		v192 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v192 = v100
	goto L1
L20:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v112 <= int32(0) {
		v192 = v3
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v119 = int32(0)
	goto L22
L22:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v119<<(uint(int32(2))%32))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v130 = int32(0)
	if base.B2i32(v129 == v130)|base.B2i32(l1 == v130) != 0 {
		v176 = base.B2i32(v129|l1 == v130)
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v192 = int32(0)
	goto L1
L24:
	;
	if v176 != 0 {
		v192 = v128
		goto L1
	} else {
		goto L35
	}
L25:
	;
	goto L24
L26:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v144 != v145 {
		v176 = int32(0)
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v147 = int32(1)
	if v144 <= v147 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v150 = v147
	goto L30
L29:
	;
	v150 = v144
	goto L30
L30:
	;
	v151 = int32(8)
	v156 = int32(0)
	goto L31
L31:
	;
	v164 = v156 << (uint(int32(2)) % 32)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v129+v151+v164)))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l1+v151+v164)))
	v169 = base.B2i32(v166 == v168)
	if v166 != v168 {
		v176 = v169
		goto L25
	} else {
		goto L33
	}
L32:
	;
	v176 = v169
	goto L25
L33:
	;
	v172 = v156 + int32(1)
	if v172 != v150 {
		v156 = v172
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v182 = v119 + int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v182 < v183 {
		v119 = v182
		goto L22
	} else {
		goto L36
	}
L36:
	;
	goto L23
}
func F_join_clause_is_movable_to(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_bms_is_member(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v115 & int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v6 == int32(0) {
		v115 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_bms_is_member(m, v12, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	if v14 != 0 {
		v115 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v18 = int32(0)
	if base.B2i32(v16 == v18)|base.B2i32(v17 == v18) != 0 {
		v63 = v18
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v63 != 0 {
		v115 = v3
		goto L1
	} else {
		goto L20
	}
L8:
	;
	goto L7
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v28 < v29 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v31 = v28
	goto L12
L11:
	;
	v31 = v29
	goto L12
L12:
	;
	if v31 <= int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v34 = int32(1)
	goto L15
L14:
	;
	v34 = v31
	goto L15
L15:
	;
	v35 = int32(8)
	v40 = int32(0)
	goto L16
L16:
	;
	v47 = v40 << (uint(int32(2)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17+v35+v47)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v16+v35+v47)))
	v52 = v49 & v51
	v54 = base.B2i32(v52 != int32(0))
	if v52 != 0 {
		v63 = v54
		goto L8
	} else {
		goto L18
	}
L17:
	;
	v63 = v54
	goto L8
L18:
	;
	v56 = v40 + int32(1)
	if v56 != v34 {
		v40 = v56
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v66 = int32(0)
	if base.B2i32(v64 == v66)|base.B2i32(v65 == v66) != 0 {
		v111 = v66
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v111 != 0 {
		v115 = v3
		goto L1
	} else {
		goto L34
	}
L22:
	;
	goto L21
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v76 < v77 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v79 = v76
	goto L26
L25:
	;
	v79 = v77
	goto L26
L26:
	;
	if v79 <= int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v82 = int32(1)
	goto L29
L28:
	;
	v82 = v79
	goto L29
L29:
	;
	v83 = int32(8)
	v88 = int32(0)
	goto L30
L30:
	;
	v95 = v88 << (uint(int32(2)) % 32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v65+v83+v95)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v64+v83+v95)))
	v100 = v97 & v99
	v102 = base.B2i32(v100 != int32(0))
	if v100 != 0 {
		v111 = v102
		goto L22
	} else {
		goto L32
	}
L31:
	;
	v111 = v102
	goto L22
L32:
	;
	v104 = v88 + int32(1)
	if v104 != v82 {
		v88 = v104
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v115 = v112 ^ int32(1)
	goto L1
}
func F_join_selectivity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 float64
	_ = v82
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v104 float64
	_ = v104
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = F_get_oprjoin(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return float64(0)
	} else {
		if v15 == int32(0) {
			v104 = float64(0.5)
			m.G0 = v13 + int32(16)
			return v104
		} else {
			v23 = m.G0
			v25 = v23 - int32(96)
			m.G0 = v25
			v28 = v25 + int32(8)
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_join_selectivity[0]))
			F_fmgr_info_cxt_security(m, v15, v28, v30, int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return float64(0)
			} else {
				v34 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+92)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+88)) = l5
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+84)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+80)) = base.I32_extend16_s(l4)
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+76)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = l2
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+68)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = l1
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+60)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+56)) = l0
				v49 = int32(5)
				*(*uint16)(unsafe.Add(mBase, uint32(v25)+54)) = uint16(v49)
				*(*uint8)(unsafe.Add(mBase, uint32(v25)+52)) = uint8(v34)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = l3
				*(*int64)(unsafe.Add(mBase, uint32(v25)+40)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v28
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				v60 = m.T0[v59].(func(*base.Module, int32) int32)(m, v25+int32(36))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return float64(0)
				} else {
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+52)))
					if v62 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return float64(0)
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v69
							F_errmsg_internal(m, int32(_a_F_join_selectivity_0), v25)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(_a_F_join_selectivity_1), int32(1246), int32(_a_F_join_selectivity_2))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return float64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						m.G0 = v25 + int32(96)
						v82 = *(*float64)(unsafe.Add(mBase, uint32(v60)))
						if base.F64_lt(v82, float64(0))|base.F64_gt(v82, float64(1)) == int32(0) {
							v104 = v82
							m.G0 = v13 + int32(16)
							return v104
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return float64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v13))) = v82
								F_errmsg_internal(m, int32(_a_F_join_selectivity_3), v13)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(_a_F_join_selectivity_4), int32(2048), int32(_a_F_join_selectivity_5))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return float64(0)
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
}
