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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(2281)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+40)) = int64(90194315497)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = int64(111669151977)
	v19 = F_LookupFuncName(m, l0, int32(5), v5+int32(-32), int32(1))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v27 = F_LookupFuncName(m, l0, int32(4), v5+int32(-32), int32(1))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				if v27 == int32(0) {
					v57 = v19
					v58 = F_get_func_rettype(m, v57)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						if v58 == int32(701) {
							if base.Ui32(int32(10000)) <= base.Ui32(v57) {
								v64 = F_superuser(m)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									if v64 != 0 {
										m.G0 = v7 - int32(-64)
										return v57
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(262577), int32(0))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(513781), int32(379), int32(218659))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
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
								v84 = *(*int32)(unsafe.Add(mBase, _consts[168]))
								v86 = F_object_aclcheck(m, int32(1255), v57, v84, int64(128))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									if v86 == int32(0) {
										m.G0 = v7 - int32(-64)
										return v57
									} else {
										v91 = F_NameListToString(m, l0)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											F_aclcheck_error(m, v86, int32(19), v91)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												m.G0 = v7 - int32(-64)
												return v57
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									v107 = F_NameListToString(m, l0)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(576332)
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v107
										F_errmsg(m, int32(199911), v7)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(513781), int32(371), int32(218659))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
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
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(84439172))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = F_NameListToString(m, l0)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v38
								F_errmsg(m, int32(177481), v5+int32(-48))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(513781), int32(356), int32(218659))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
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
				if v27 != 0 {
					v57 = v27
					v58 = F_get_func_rettype(m, v57)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						if v58 == int32(701) {
							if base.Ui32(int32(10000)) <= base.Ui32(v57) {
								v64 = F_superuser(m)
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									if v64 != 0 {
										m.G0 = v7 - int32(-64)
										return v57
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(262577), int32(0))
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(513781), int32(379), int32(218659))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
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
								v84 = *(*int32)(unsafe.Add(mBase, _consts[168]))
								v86 = F_object_aclcheck(m, int32(1255), v57, v84, int64(128))
								mBase = m.M
								v87 = m.ExcPending
								if v87 != 0 {
									return int32(0)
								} else {
									if v86 == int32(0) {
										m.G0 = v7 - int32(-64)
										return v57
									} else {
										v91 = F_NameListToString(m, l0)
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int32(0)
										} else {
											F_aclcheck_error(m, v86, int32(19), v91)
											mBase = m.M
											v94 = m.ExcPending
											if v94 != 0 {
												return int32(0)
											} else {
												m.G0 = v7 - int32(-64)
												return v57
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(117833860))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return int32(0)
								} else {
									v107 = F_NameListToString(m, l0)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(576332)
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v107
										F_errmsg(m, int32(199911), v7)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(513781), int32(371), int32(218659))
											mBase = m.M
											v119 = m.ExcPending
											if v119 != 0 {
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
					v55 = F_LookupFuncName(m, l0, int32(5), v5+int32(-32), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v57 = v55
						v58 = F_get_func_rettype(m, v57)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return int32(0)
						} else {
							if v58 == int32(701) {
								if base.Ui32(int32(10000)) <= base.Ui32(v57) {
									v64 = F_superuser(m)
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return int32(0)
									} else {
										if v64 != 0 {
											m.G0 = v7 - int32(-64)
											return v57
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(16797828))
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(262577), int32(0))
													mBase = m.M
													v76 = m.ExcPending
													if v76 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(513781), int32(379), int32(218659))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
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
									v84 = *(*int32)(unsafe.Add(mBase, _consts[168]))
									v86 = F_object_aclcheck(m, int32(1255), v57, v84, int64(128))
									mBase = m.M
									v87 = m.ExcPending
									if v87 != 0 {
										return int32(0)
									} else {
										if v86 == int32(0) {
											m.G0 = v7 - int32(-64)
											return v57
										} else {
											v91 = F_NameListToString(m, l0)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												F_aclcheck_error(m, v86, int32(19), v91)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 - int32(-64)
													return v57
												}
											}
										}
									}
								}
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(117833860))
									mBase = m.M
									v106 = m.ExcPending
									if v106 != 0 {
										return int32(0)
									} else {
										v107 = F_NameListToString(m, l0)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(576332)
											*(*int32)(unsafe.Add(mBase, uint32(v7))) = v107
											F_errmsg(m, int32(199911), v7)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(513781), int32(371), int32(218659))
												mBase = m.M
												v119 = m.ExcPending
												if v119 != 0 {
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
	var v49 int32
	_ = v49
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
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
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
	return v190
L2:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v109 == int32(0) {
		v190 = v3
		goto L1
	} else {
		goto L20
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v16 == int32(0) {
		v190 = v3
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
	v30 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v30
	v37 = F_hash_create(m, int32(412869), int32(256), v9+int32(-48), int32(1224))
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
	v49 = int32(0)
	goto L13
L13:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55+v49<<(uint(int32(2))%32))))
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
	v69 = v49 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v69 < v70 {
		v49 = v69
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
		v190 = v3
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v190 = v100
	goto L1
L20:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v112 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v190 = v3
	goto L1
L22:
	;
	goto L23
L23:
	;
	v118 = int32(0)
	goto L24
L24:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124+v118<<(uint(int32(2))%32))))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)+8))
	v130 = int32(0)
	v137 = base.B2i32(v129|l1 == v130)
	if v129 == v130 {
		v176 = v137
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v190 = int32(0)
	goto L1
L26:
	;
	if v176 != 0 {
		v190 = v128
		goto L1
	} else {
		goto L38
	}
L27:
	;
	goto L26
L28:
	;
	if l1 == int32(0) {
		v176 = v137
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v143 != v144 {
		v176 = int32(0)
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v146 = int32(1)
	if v143 <= v146 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v149 = v146
	goto L33
L32:
	;
	v149 = v143
	goto L33
L33:
	;
	v150 = int32(8)
	v155 = int32(0)
	goto L34
L34:
	;
	v163 = v155 << (uint(int32(2)) % 32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v129+v150+v163)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+(l1+v150))))
	v168 = base.B2i32(v165 == v167)
	if v167 != v165 {
		v176 = v168
		goto L27
	} else {
		goto L36
	}
L35:
	;
	v176 = v168
	goto L27
L36:
	;
	v171 = v155 + int32(1)
	if v171 != v149 {
		v155 = v171
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v181 = v118 + int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	if v181 < v182 {
		v118 = v181
		goto L24
	} else {
		goto L39
	}
L39:
	;
	goto L25
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
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
	return v113 & int32(1)
L2:
	;
	return int32(0)
L3:
	;
	if v6 == int32(0) {
		v113 = v3
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
		v113 = v3
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v18 = int32(0)
	if v16 == v18 {
		v59 = v18
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v59 != 0 {
		v113 = v3
		goto L1
	} else {
		goto L21
	}
L8:
	;
	goto L7
L9:
	;
	if v17 == int32(0) {
		v59 = v18
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v27 < v28 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v30 = v27
	goto L13
L12:
	;
	v30 = v28
	goto L13
L13:
	;
	if v30 <= int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v33 = int32(1)
	goto L16
L15:
	;
	v33 = v30
	goto L16
L16:
	;
	v34 = int32(8)
	v39 = int32(0)
	goto L17
L17:
	;
	v46 = v39 << (uint(int32(2)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v17+v34+v46)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46+(v16+v34))))
	v51 = v48 & v50
	v53 = base.B2i32(v51 != int32(0))
	if v51 != 0 {
		v59 = v53
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v59 = v53
	goto L8
L19:
	;
	v55 = v39 + int32(1)
	if v55 != v33 {
		v39 = v55
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v65 = int32(0)
	if v63 == v65 {
		v106 = v65
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v106 != 0 {
		v113 = v3
		goto L1
	} else {
		goto L36
	}
L23:
	;
	goto L22
L24:
	;
	if v64 == int32(0) {
		v106 = v65
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v74 < v75 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v77 = v74
	goto L28
L27:
	;
	v77 = v75
	goto L28
L28:
	;
	if v77 <= int32(1) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v80 = int32(1)
	goto L31
L30:
	;
	v80 = v77
	goto L31
L31:
	;
	v81 = int32(8)
	v86 = int32(0)
	goto L32
L32:
	;
	v93 = v86 << (uint(int32(2)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v64+v81+v93)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+(v63+v81))))
	v98 = v95 & v97
	v100 = base.B2i32(v98 != int32(0))
	if v98 != 0 {
		v106 = v100
		goto L23
	} else {
		goto L34
	}
L33:
	;
	v106 = v100
	goto L23
L34:
	;
	v102 = v86 + int32(1)
	if v102 != v80 {
		v86 = v102
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v113 = v110 ^ int32(1)
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v49 int32
	_ = v49
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 float64
	_ = v84
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 float64
	_ = v108
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
			v108 = float64(0.5)
			m.G0 = v13 + int32(16)
			return v108
		} else {
			v23 = m.G0
			v25 = v23 - int32(96)
			m.G0 = v25
			v30 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			F_fmgr_info_cxt_security(m, v15, v25+int32(8), v30, int32(0))
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
				*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v25 + int32(8)
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
				v62 = m.T0[v61].(func(*base.Module, int32) int32)(m, v25+int32(36))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return float64(0)
				} else {
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+52)))
					if v64 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return float64(0)
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v25))) = v71
							F_errmsg_internal(m, int32(553407), v25)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return float64(0)
							} else {
								F_errfinish(m, int32(514773), int32(1246), int32(315939))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
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
						v84 = *(*float64)(unsafe.Add(mBase, uint32(v62)))
						if base.F64_lt(v84, float64(0))|base.F64_gt(v84, float64(1)) == int32(0) {
							v108 = v84
							m.G0 = v13 + int32(16)
							return v108
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return float64(0)
							} else {
								*(*float64)(unsafe.Add(mBase, uint32(v13))) = v84
								F_errmsg_internal(m, int32(354033), v13)
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return float64(0)
								} else {
									F_errfinish(m, int32(512615), int32(2048), int32(9971))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
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
