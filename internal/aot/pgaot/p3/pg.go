package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_PgArchShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchShmemInit[0]))
	if v3&int32(3) == int32(0) {
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v3))) = int32(-1)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_PgArchShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = int32(0)
	return
}
func F_ValidatePgVersion(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int64
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	v8 = m.G0
	v10 = v8 - int32(1232)
	m.G0 = v10
	v17 = F_strtox_2(m, int32(_a_F_ValidatePgVersion_0), v10+int32(204), int32(10), int64(2147483648))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v10)+112)) = l0
	v21 = v10 + int32(208)
	v26 = F_pg_snprintf(m, v21, int32(1024), int32(_a_F_ValidatePgVersion_1), v10+int32(112))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return
	} else {
		v29 = F_AllocateFile(m, v21, int32(_a_F_ValidatePgVersion_2))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return
		} else {
			if v29 == int32(0) {
				v34 = *(*int32)(unsafe.Add(mBase, _c_F_ValidatePgVersion[0]))
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v34 == int32(44) {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v82 = m.ExcPending
						if v82 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
							F_errmsg(m, int32(_a_F_ValidatePgVersion_3), v10+int32(16))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = v10 + int32(208)
								v93 = F_errdetail(m, int32(_a_F_ValidatePgVersion_4), v10)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1742), int32(_a_F_ValidatePgVersion_6))
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
					} else {
						F_errcode_for_file_access(m)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v21
							F_errmsg(m, int32(_a_F_ValidatePgVersion_7), v10+int32(32))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1746), int32(_a_F_ValidatePgVersion_6))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
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
				v54 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+128)) = uint8(v54)
				v57 = v10 + int32(128)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+96)) = v57
				v62 = F_fscanf(m, v29, int32(_a_F_ValidatePgVersion_8), v10+int32(96))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v68 = F_strtox_2(m, v57, v10+int32(204), int32(10), int64(2147483648))
					mBase = m.M
					if v62 != int32(1) {
						F_errstart_cold(m, int32(22), int32(0))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l0
								F_errmsg(m, int32(_a_F_ValidatePgVersion_3), v10-int32(-64))
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(208)
									v119 = F_errdetail(m, int32(_a_F_ValidatePgVersion_9), v10+int32(48))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return
									} else {
										F_errhint(m, int32(_a_F_ValidatePgVersion_10), int32(0))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1760), int32(_a_F_ValidatePgVersion_6))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
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
					} else {
						v72 = *(*int32)(unsafe.Add(mBase, uint32(v10)+204))
						if v57 == v72 {
							F_errstart_cold(m, int32(22), int32(0))
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v106 = m.ExcPending
								if v106 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v10)+64)) = l0
									F_errmsg(m, int32(_a_F_ValidatePgVersion_3), v10-int32(-64))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v10 + int32(208)
										v119 = F_errdetail(m, int32(_a_F_ValidatePgVersion_9), v10+int32(48))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return
										} else {
											F_errhint(m, int32(_a_F_ValidatePgVersion_10), int32(0))
											mBase = m.M
											v124 = m.ExcPending
											if v124 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1760), int32(_a_F_ValidatePgVersion_6))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
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
						} else {
							v74 = F_FreeFile(m, v29)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								if base.I32_wrap_i64(v17) != base.I32_wrap_i64(v68) {
									F_errstart_cold(m, int32(22), int32(0))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return
									} else {
										F_errcode(m, int32(50856066))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return
										} else {
											F_errmsg(m, int32(_a_F_ValidatePgVersion_11), int32(0))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v10)+84)) = int32(_a_F_ValidatePgVersion_0)
												*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v10 + int32(128)
												v149 = F_errdetail(m, int32(_a_F_ValidatePgVersion_12), v10+int32(80))
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1770), int32(_a_F_ValidatePgVersion_6))
													mBase = m.M
													v155 = m.ExcPending
													if v155 != 0 {
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
									m.G0 = v10 + int32(1232)
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
func F__PG_init_plpgsql(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[0])))
	if v4 == int32(0) {
		v9 = int32(0)
		F_DefineCustomEnumVariable(m, int32(_a_F__PG_init_plpgsql_0), int32(_a_F__PG_init_plpgsql_1), v9, int32(_a_F__PG_init_plpgsql_2), v9, int32(_a_F__PG_init_plpgsql_3), int32(5))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v18 = int32(0)
			F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_plpgsql_4), int32(_a_F__PG_init_plpgsql_5), v18, int32(_a_F__PG_init_plpgsql_6), v18, int32(6))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return
			} else {
				F_DefineCustomBoolVariable(m, int32(_a_F__PG_init_plpgsql_7), int32(_a_F__PG_init_plpgsql_8), int32(0), int32(_a_F__PG_init_plpgsql_9), int32(1), int32(6))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_DefineCustomStringVariable(m, int32(_a_F__PG_init_plpgsql_10), int32(_a_F__PG_init_plpgsql_11), int32(_a_F__PG_init_plpgsql_12), int32(_a_F__PG_init_plpgsql_13), int32(6), int32(1), int32(_a_F__PG_init_plpgsql_14), int32(_a_F__PG_init_plpgsql_15))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_DefineCustomStringVariable(m, int32(_a_F__PG_init_plpgsql_16), int32(_a_F__PG_init_plpgsql_17), int32(_a_F__PG_init_plpgsql_18), int32(_a_F__PG_init_plpgsql_13), int32(6), int32(1), int32(_a_F__PG_init_plpgsql_14), int32(_a_F__PG_init_plpgsql_19))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_plpgsql_20))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v56 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[1]))
								v58 = F_MemoryContextAlloc(m, v56, int32(12))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = int32(_a_F__PG_init_plpgsql_21)
									v64 = int32(_a_F__PG_init_plpgsql_22)
									v65 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[2]))
									*(*int32)(unsafe.Add(mBase, uint32(v58))) = v65
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[2])) = v58
									v70 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[1]))
									v72 = F_MemoryContextAlloc(m, v70, int32(12))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return
									} else {
										v74 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = v74
										*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = int32(_a_F__PG_init_plpgsql_23)
										v78 = int32(_a_F__PG_init_plpgsql_24)
										v79 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[3]))
										*(*int32)(unsafe.Add(mBase, uint32(v72))) = v79
										*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[3])) = v72
										v83 = m.G0
										v85 = v83 - int32(48)
										m.G0 = v85
										v88 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[4]))
										if v88 == v74 {
											*(*int64)(unsafe.Add(mBase, uint32(v85)+8)) = int64(292057776192)
											v97 = F_hash_create(m, int32(_a_F__PG_init_plpgsql_25), int64(16), v85, int32(24))
											mBase = m.M
											v98 = m.ExcPending
											if v98 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[4])) = v97
												v100 = v97
												v103 = F_hash_search(m, v100, int32(_a_F__PG_init_plpgsql_26), int32(1), v85)
												mBase = m.M
												v104 = m.ExcPending
												if v104 != 0 {
													return
												} else {
													v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
													if v105 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v103)+64)) = int32(0)
													} else {
													}
													m.G0 = v85 + int32(48)
													v114 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[0])) = uint8(v114)
													*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[5])) = v103 - int32(-64)
													return
												}
											}
										} else {
											v100 = v88
											v103 = F_hash_search(m, v100, int32(_a_F__PG_init_plpgsql_26), int32(1), v85)
											mBase = m.M
											v104 = m.ExcPending
											if v104 != 0 {
												return
											} else {
												v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
												if v105 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v103)+64)) = int32(0)
												} else {
												}
												m.G0 = v85 + int32(48)
												v114 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[0])) = uint8(v114)
												*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[5])) = v103 - int32(-64)
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
		return
	}
}
func F_create_pg_locale_icu(m *base.Module) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	F_errstart_cold(m, int32(21), int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F_errmsg(m, int32(_a_F_create_pg_locale_icu_0), int32(0))
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_create_pg_locale_icu_1), int32(388), int32(_a_F_create_pg_locale_icu_2))
				v18 = m.ExcPending
				if v18 != 0 {
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
func F_do_pg_backup_stop(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v113 int64
	_ = v113
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v132 int64
	_ = v132
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int64
	_ = v195
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v350 int32
	_ = v350
	var v364 int32
	_ = v364
	var v369 int32
	_ = v369
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v378 int64
	_ = v378
	var v379 int64
	_ = v379
	var v381 int64
	_ = v381
	var v382 int64
	_ = v382
	var v385 int64
	_ = v385
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int64
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int64
	_ = v405
	var v406 int64
	_ = v406
	var v408 int64
	_ = v408
	var v409 int64
	_ = v409
	var v412 int64
	_ = v412
	var v420 int32
	_ = v420
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
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v528 int32
	_ = v528
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	v11 = m.G0
	v13 = v11 - int32(2272)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[0])))
	if v16 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[1]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+308))
	v24 = base.B2i32(v22 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[0])) = uint8(v24)
	v26 = v24
	goto L3
L2:
	;
	v26 = int32(0)
	goto L3
L3:
	;
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L13
	} else {
		goto L158
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L13
	} else {
		goto L154
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L13
	} else {
		goto L149
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L13
	} else {
		goto L144
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L13
	} else {
		goto L139
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[2]))
	if v30 <= int32(0) {
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_WALInsertLockAcquireExclusive(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	return
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[1]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+164)) = v37 - int32(1)
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[3])) = uint8(v42)
	F_WALInsertLockRelease(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1064)))
	if (v26|(v46^int32(-1)))&int32(1) == int32(0) {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	if v26 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if l1 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[1]))
	v58 = base.AtomicRmwXchg32(m, v55, int32(440), int32(1))
	if v58 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L13
	} else {
		goto L28
	}
L21:
	;
	F_s_lock(m, v55+int32(440), int32(_a_F_do_pg_backup_stop_0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[1]))
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v65)+432))
	v67 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v65)+440)), uint32(v67))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	if base.Ui64(v70) <= base.Ui64(v66) {
		goto L6
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[4]))
	v77 = F_LWLockAcquire(m, v73+int32(1152), int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[5]))
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v80)+144))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1088)) = v81
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)+152))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1096)) = v83
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[4]))
	F_LWLockRelease(m, v86+int32(1152))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L13
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	F_XLogRegisterData(m, l0+int32(1032), int32(8))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v100 = F_XLogInsert(m, int32(0), int32(80))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1088)) = v100
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[1]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+300))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1096)) = v105
	F_XLogBeginInsert(m)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v111 = F_XLogInsert(m, int32(0), int32(64))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v113 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1104)) = v113
	v115 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v116
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = base.I32_wrap_i64(v115) & (v120 - int32(1))
	v125 = base.I64_extend_i32_s(v120)
	v126 = base.I64_div_u_s(v115, v125)
	v128 = base.I64_div_u_s(int64(4294967296), v125)
	v129 = base.I64_div_u_s(v126, v128)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+132)) = uint32(v129)
	v132 = v126 - v128*v129
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+136)) = uint32(v132)
	v135 = v13 + int32(208)
	v140 = F_pg_snprintf(m, v135, int32(1024), int32(_a_F_do_pg_backup_stop_1), v13+int32(128))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v143 = F_AllocateFile(m, v135, int32(_a_F_do_pg_backup_stop_2))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	if v143 == int32(0) {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v148 = F_build_backup_content(m, l0, int32(1))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v148
	v154 = F_pg_fprintf(m, v143, int32(_a_F_do_pg_backup_stop_3), v13+int32(112))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, v148)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v158 = F_fflush(m, v143)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	if v158 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	goto L41
L41:
	;
	if int32(base.Ui32(v160)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v165 = F_FreeFile(m, v143)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	if v165 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v168 = F_AllocateDir(m, int32(_a_F_do_pg_backup_stop_4))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v171 = F_ReadDir(m, v168, int32(_a_F_do_pg_backup_stop_4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	if v171 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v176 = v171
	goto L50
L48:
	;
	goto L49
L49:
	;
	F_FreeDir(m, v168)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L13
	} else {
		goto L94
	}
L50:
	;
	v184 = v176 + int32(19)
	v185 = F_strlen(m, v184)
	mBase = m.M
	if base.Ui32(v185) < base.Ui32(int32(25)) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v337 = F_ReadDir(m, v168, int32(_a_F_do_pg_backup_stop_4))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L13
	} else {
		goto L92
	}
L53:
	;
	v188 = int32(_a_F_do_pg_backup_stop_5)
	v192 = m.G0
	v194 = v192 - int32(32)
	v195 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v194)+24)) = v195
	*(*int64)(unsafe.Add(mBase, uint32(v194)+16)) = v195
	*(*int64)(unsafe.Add(mBase, uint32(v194)+8)) = v195
	*(*int64)(unsafe.Add(mBase, uint32(v194))) = v195
	v203 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[7])))
	if v203 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v271 != int32(24) {
		goto L52
	} else {
		goto L73
	}
L55:
	;
	v271 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[8])))
	if v207 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v211 = v184
	goto L61
L59:
	;
	goto L60
L60:
	;
	v221 = v188
	v222 = v203
	goto L64
L61:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v217 == v203 {
		v211 = v211 + int32(1)
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v271 = v211 - v184
	goto L54
L63:
	;
	goto L62
L64:
	;
	v229 = v194 + int32(base.Ui32(v222)>>(uint(int32(3))%32))&int32(28)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v231 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v229))) = v230 | v231<<(uint(v222)%32)
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	if v235 != 0 {
		v221 = v221 + v231
		v222 = v235
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v238 == int32(0) {
		v261 = v184
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v271 = v261 - v184
	goto L54
L68:
	;
	v242 = v184
	v243 = v238
	goto L69
L69:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v194+int32(base.Ui32(v243)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v251)>>(uint(v243)%32))&int32(1) == int32(0) {
		v261 = v242
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v261 = v259
	goto L67
L71:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)))
	v259 = v242 + int32(1)
	if v257 != 0 {
		v242 = v259
		v243 = v257
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v276 = v184 + v185 - int32(7)
	v277 = int32(_a_F_do_pg_backup_stop_6)
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276))))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[9])))
	if base.B2i32(v280 == int32(0))|base.B2i32(v280 != v283) != 0 {
		v301 = v280
		v302 = v283
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v301-v302 != 0 {
		goto L52
	} else {
		goto L81
	}
L75:
	;
	goto L74
L76:
	;
	v286 = v276
	v287 = v277
	goto L77
L77:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+1)))
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v286)+1)))
	if v291 == int32(0) {
		v301 = v291
		v302 = v290
		goto L75
	} else {
		goto L79
	}
L78:
	;
	v301 = v291
	v302 = v290
	goto L75
L79:
	;
	v294 = int32(1)
	if v291 == v290 {
		v286 = v286 + v294
		v287 = v287 + v294
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v304 = F_XLogArchiveCheckDone(m, v184)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	if v304 == int32(0) {
		goto L52
	} else {
		goto L83
	}
L83:
	;
	v310 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	if v310 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v184
	F_errmsg_internal(m, int32(_a_F_do_pg_backup_stop_7), v13+int32(80))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L13
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v184
	v325 = v13 + int32(1232)
	v330 = F_pg_snprintf(m, v325, int32(1031), int32(_a_F_do_pg_backup_stop_8), v13-int32(-64))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), int32(_a_F_do_pg_backup_stop_10), int32(_a_F_do_pg_backup_stop_11))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v332 = F_unlink(m, v325)
	mBase = m.M
	F_XLogArchiveCleanup(m, v184)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	goto L52
L92:
	;
	if v337 != 0 {
		v176 = v337
		goto L50
	} else {
		goto L93
	}
L93:
	;
	goto L51
L94:
	;
	goto L17
L95:
	;
	m.G0 = v13 + int32(2272)
	return
L96:
	;
	v364 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[10]))
	if v26 != 0 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	F_errmsg(m, v528, int32(0))
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L13
	} else {
		goto L137
	}
L98:
	;
	v369 = base.B2i32(v364 == int32(2))
	goto L100
L99:
	;
	v369 = base.B2i32(int32(0) < v364)
	goto L100
L100:
	;
	if v369 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v372 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v373
	v378 = int64(*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[6])))
	v379 = base.I64_div_u_s(v372-int64(1), v378)
	v381 = base.I64_div_u_s(int64(4294967296), v378)
	v382 = base.I64_div_u_s(v379, v381)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v382)
	v385 = v379 - v381*v382
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v385)
	v393 = F_pg_snprintf(m, v13+int32(1232), int32(64), int32(_a_F_do_pg_backup_stop_12), v13+int32(32))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L13
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	v522 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L13
	} else {
		goto L135
	}
L104:
	;
	v395 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v396
	v400 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = base.I32_wrap_i64(v395) & (v400 - int32(1))
	v405 = base.I64_extend_i32_s(v400)
	v406 = base.I64_div_u_s(v395, v405)
	v408 = base.I64_div_u_s(int64(4294967296), v405)
	v409 = base.I64_div_u_s(v406, v408)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+20)) = uint32(v409)
	v412 = v406 - v408*v409
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+24)) = uint32(v412)
	v420 = F_pg_snprintf(m, v13+int32(144), int32(64), int32(_a_F_do_pg_backup_stop_13), v13+int32(16))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	v423 = int32(0)
	v425 = v423
	v426 = int32(60)
	v428 = v423
	goto L106
L106:
	;
	v437 = F_XLogArchiveIsBusy(m, v13+int32(1232))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L13
	} else {
		goto L109
	}
L107:
	;
	v514 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L13
	} else {
		goto L133
	}
L108:
	;
	goto L107
L109:
	;
	if v437 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v443 = F_XLogArchiveIsBusy(m, v13+int32(144))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L13
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v448 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[11]))
	if v448 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	if v443 == int32(0) {
		goto L108
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v450 = m.ExcPending
	if v450 != 0 {
		goto L13
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if (base.B2i32(v428 < int32(6))|v425)&int32(1) != 0 {
		v472 = v425
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L117
L119:
	;
	v474 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[12]))
	v478 = F_WaitLatch(m, v474, int32(41), int32(1000), int32(134217732))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L13
	} else {
		goto L125
	}
L120:
	;
	v456 = int32(1)
	v459 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L13
	} else {
		goto L121
	}
L121:
	;
	if v459 == int32(0) {
		v472 = v456
		goto L119
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_14), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L13
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), int32(_a_F_do_pg_backup_stop_15), int32(_a_F_do_pg_backup_stop_16))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	v472 = v456
	goto L119
L125:
	;
	v481 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[12]))
	v482 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v481))) = v482
	v487 = base.AtomicRmwOr32(m, v482, int32(_a_F_do_pg_backup_stop_17), v482)
	goto L126
L126:
	;
	v489 = v428 + int32(1)
	if v489 < v426 {
		v425 = v472
		v428 = v489
		goto L106
	} else {
		goto L127
	}
L127:
	;
	v492 = v426 << (uint(int32(1)) % 32)
	v495 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L13
	} else {
		goto L128
	}
L128:
	;
	if v495 == int32(0) {
		v425 = v472
		v426 = v492
		v428 = v489
		goto L106
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v489
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_18), v13)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L13
	} else {
		goto L130
	}
L130:
	;
	F_errhint(m, int32(_a_F_do_pg_backup_stop_19), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L13
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), int32(_a_F_do_pg_backup_stop_20), int32(_a_F_do_pg_backup_stop_16))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L13
	} else {
		goto L132
	}
L132:
	;
	v425 = v472
	v426 = v492
	v428 = v489
	goto L106
L133:
	;
	if v514 == int32(0) {
		goto L95
	} else {
		goto L134
	}
L134:
	;
	v528 = int32(_a_F_do_pg_backup_stop_21)
	v538 = int32(_a_F_do_pg_backup_stop_22)
	goto L97
L135:
	;
	if v522 == int32(0) {
		goto L95
	} else {
		goto L136
	}
L136:
	;
	v528 = int32(_a_F_do_pg_backup_stop_23)
	v538 = int32(_a_F_do_pg_backup_stop_24)
	goto L97
L137:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), v538, int32(_a_F_do_pg_backup_stop_16))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L13
	} else {
		goto L138
	}
L138:
	;
	goto L95
L139:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_25), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	F_errhint(m, int32(_a_F_do_pg_backup_stop_26), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L13
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), int32(_a_F_do_pg_backup_stop_27), int32(_a_F_do_pg_backup_stop_16))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L13
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L144:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L13
	} else {
		goto L145
	}
L145:
	;
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_28), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L13
	} else {
		goto L146
	}
L146:
	;
	F_errhint(m, int32(_a_F_do_pg_backup_stop_29), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L13
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), int32(_a_F_do_pg_backup_stop_30), int32(_a_F_do_pg_backup_stop_16))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L13
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_31), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	F_errhint(m, int32(_a_F_do_pg_backup_stop_32), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), int32(_a_F_do_pg_backup_stop_33), int32(_a_F_do_pg_backup_stop_16))
	mBase = m.M
	v618 = m.ExcPending
	if v618 != 0 {
		goto L13
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L13
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v13 + int32(208)
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_34), v13+int32(48))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), int32(_a_F_do_pg_backup_stop_35), int32(_a_F_do_pg_backup_stop_16))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L13
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L13
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = v13 + int32(208)
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_36), v13+int32(96))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_9), int32(_a_F_do_pg_backup_stop_37), int32(_a_F_do_pg_backup_stop_16))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L13
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_b64_decode(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
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
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v161 int32
	_ = v161
	v5 = int32(0)
	if v5 < l1 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 != 0 {
		goto L43
	} else {
		goto L44
	}
L2:
	;
	v14 = l0 + l1
	v15 = l0
	v19 = v5
	v20 = l2
	v22 = v5
	v23 = v5
	goto L5
L3:
	;
	v161 = l2
	goto L4
L4:
	;
	return v161 - l2
L5:
	;
	v27 = v15 + int32(1)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v28 != int32(61) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	if v149 != 0 {
		goto L1
	} else {
		goto L42
	}
L7:
	;
	if base.Ui32(v147) < base.Ui32(v14) {
		v15 = v147
		v19 = v149
		v20 = v150
		v22 = v152
		v23 = v153
		goto L5
	} else {
		goto L41
	}
L8:
	;
	if l3 < v20-l2+int32(1) {
		goto L1
	} else {
		goto L28
	}
L9:
	;
	v105 = v27
	v106 = int32(2)
	v109 = v23 << (uint(int32(6)) % 32)
	goto L8
L10:
	;
	v94 = v91 + v90<<(uint(int32(6))%32)
	v96 = v87 + int32(1)
	if v96 == int32(4) {
		v105 = v88
		v106 = v89
		v109 = v94
		goto L8
	} else {
		goto L27
	}
L11:
	;
	v87 = int32(3)
	v88 = v15 + int32(2)
	v89 = int32(1)
	v90 = v47
	v91 = v42
	goto L10
L12:
	;
	if base.Ui32(int32(125)) < base.Ui32((v66-int32(1))&int32(255)) {
		goto L1
	} else {
		goto L25
	}
L13:
	;
	v32 = v28 - int32(9)
	if base.B2i32(base.Ui32(int32(23)) < base.Ui32(v32))|base.B2i32(int32(1)<<(uint(v32)%32)&int32(_a_F_pg_b64_decode_0) == int32(0)) != 0 {
		v66 = v28
		v67 = v19
		v68 = v27
		v69 = v22
		v70 = v23
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v42 = int32(0)
	if v22 != 0 {
		v87 = v19
		v88 = v27
		v89 = v22
		v90 = v23
		v91 = v42
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L1
L17:
	;
	switch v19 - int32(2) {
	case 0:
		goto L18
	case 1:
		goto L9
	default:
		goto L1
	}
L18:
	;
	if base.Ui32(v14) <= base.Ui32(v27) {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v47 = v23 << (uint(int32(6)) % 32)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v48 == int32(61) {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	v52 = v48 - int32(9)
	if int32(1)<<(uint(v52)%32)&int32(_a_F_pg_b64_decode_0) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v60 = base.B2i32(base.Ui32(v52) <= base.Ui32(int32(23)))
	goto L23
L22:
	;
	v60 = int32(0)
	goto L23
L23:
	;
	if v60 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v66 = v48
	v67 = int32(3)
	v68 = v15 + int32(2)
	v69 = int32(1)
	v70 = v47
	goto L12
L25:
	;
	v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v66)+uint32(_c_F_pg_b64_decode[0]))))
	if v78 < int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v87 = v67
	v88 = v68
	v89 = v69
	v90 = v70
	v91 = v78
	goto L10
L27:
	;
	v147 = v88
	v149 = v96
	v150 = v20
	v152 = v89
	v153 = v94
	goto L7
L28:
	;
	v115 = int32(base.Ui32(v109) >> (uint(int32(16)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v20))) = uint8(v115)
	v118 = v20 + int32(1)
	if base.Ui32(v106) < base.Ui32(int32(2)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v122 = v106
	goto L31
L30:
	;
	v122 = int32(0)
	goto L31
L31:
	;
	if v122 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if l3 < v118-l2+int32(1) {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	v134 = v118
	goto L34
L34:
	;
	if v106 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v130 = int32(base.Ui32(v109) >> (uint(int32(8)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)) = uint8(v130)
	v134 = v20 + int32(2)
	goto L34
L36:
	;
	v147 = v105
	v149 = int32(0)
	v150 = v144
	v152 = v145
	v153 = int32(0)
	goto L7
L37:
	;
	v144 = v134
	v145 = v106
	goto L36
L38:
	;
	goto L39
L39:
	;
	if l3 < v134-l2+int32(1) {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v109)
	v144 = v134 + int32(1)
	v145 = int32(0)
	goto L36
L41:
	;
	goto L6
L42:
	;
	v161 = v150
	goto L4
L43:
	;
	base.MemoryFill(m, l2, int32(0), l3)
	goto L45
L44:
	;
	goto L45
L45:
	;
	return int32(-1)
}
func F_pg_backup_stop(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v2)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_stop[0])))
	v19 = F_get_call_result_type(m, l0, v2, v7+int32(44))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int64(0)
	} else {
		if v19 == int32(1) {
			if v15 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_backup_stop_0), int32(0))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int64(0)
						} else {
							F_errhint(m, int32(_a_F_pg_backup_stop_1), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_backup_stop_2), int32(172), int32(_a_F_pg_backup_stop_3))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
									return int64(0)
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
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[1]))
				F_do_pg_backup_stop(m, v28, base.B2i32(v13 != int64(0)))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[1]))
					v36 = F_build_backup_content(m, v34, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int64(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[1]))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+1088))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v40
						v42 = F_cstring_to_text(m, v36)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int64(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = base.I64_extend_i32_u(v42)
							v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[2]))
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
							v49 = F_cstring_to_text(m, v48)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v7)+32)) = base.I64_extend_i32_u(v49)
								F_pfree(m, v36)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									v56 = int32(0)
									*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[2])) = v56
									*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[1])) = v56
									v62 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[3]))
									F_MemoryContextDelete(m, v62)
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[3])) = int32(0)
										v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
										v73 = F_heap_form_tuple(m, v68, v7+int32(16), v7+int32(12))
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int64(0)
										} else {
											v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
											v76 = F_HeapTupleHeaderGetDatum(m, v75)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return int64(0)
											} else {
												m.G0 = v7 + int32(48)
												return v76
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_backup_stop_4), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_backup_stop_2), int32(166), int32(_a_F_pg_backup_stop_3))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int64(0)
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
func F_pg_base64url_dec_len(m *base.Module, l0 int32, l1 int32) int64 {
	var v3 int32
	_ = v3
	v3 = int32(3)
	return base.I64_extend_i32_u(int32(base.Ui32((l1+v3)&int32(-4)*v3) >> (uint(int32(2)) % 32)))
}
func F_pg_base64url_enc_len(m *base.Module, l0 int32, l1 int32) int64 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = int32(2)
	v6 = base.I32_div_u_s(l1+v3, int32(3))
	return base.I64_extend_i32_u(v6 << (uint(v3) % 32))
}
func F_pg_base64url_encode(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	if l1 == int32(0) {
		v85 = l2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.I64_extend_i32_s(v85 - l2)
L2:
	;
	v11 = l0 + l1
	v12 = l0
	v15 = l2
	goto L3
L3:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v22 = v20 << (uint(int32(16)) % 32)
	v25 = base.B2i32(base.Ui32(v11) <= base.Ui32(v12+int32(1)))
	if base.Ui32(v11) <= base.Ui32(v12+int32(1)) {
		v60 = v22
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v60)>>(uint(int32(18))%32)))+uint32(_c_F_pg_base64url_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v64)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v60)>>(uint(int32(12))%32))&int32(63))+uint32(_c_F_pg_base64url_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v70)
	if base.Ui32(v11) <= base.Ui32(v12+int32(1)) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L4
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v28 = v26 << (uint(int32(8)) % 32)
	v29 = v28 | v22
	if base.Ui32(v11) <= base.Ui32(v12+int32(2)) {
		v60 = v29
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+2)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v20)>>(uint(int32(2))%32)))+uint32(_c_F_pg_base64url_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v36)
	v38 = int32(63)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33&v38)+uint32(_c_F_pg_base64url_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+3)) = uint8(v40)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v29)>>(uint(int32(12))%32))&v38)+uint32(_c_F_pg_base64url_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)) = uint8(v46)
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v33|v28)>>(uint(int32(6))%32))&v38)+uint32(_c_F_pg_base64url_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v53)
	v56 = v15 + int32(4)
	v58 = v12 + int32(3)
	if base.Ui32(v58) < base.Ui32(v11) {
		v12 = v58
		v15 = v56
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v85 = v56
	goto L1
L9:
	;
	v85 = v15 + int32(2)
	goto L1
L10:
	;
	goto L11
L11:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v60)>>(uint(int32(6))%32))&int32(63))+uint32(_c_F_pg_base64url_encode[0]))))
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+2)) = uint8(v78)
	v85 = v15 + int32(3)
	goto L1
}
func F_pg_basetype(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	v7 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	v8 = F_SearchSysCache1(m, int32(82), v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_ReleaseCatCache(m, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L3
	} else {
		goto L15
	}
L2:
	;
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v46)
	return int64(0)
L3:
	;
	return int64(0)
L4:
	;
	if v8 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
	v16 = v14 + v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+79)))
	if v17 != int32(100) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v51 = v8
	v52 = base.I32_wrap_i64(v7)
	goto L1
L7:
	;
	goto L8
L8:
	;
	v22 = v8
	v24 = v16
	goto L9
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+132))
	F_ReleaseCatCache(m, v22)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v51 = v31
	v52 = v26
	goto L1
L11:
	;
	v31 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v26))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L3
	} else {
		goto L12
	}
L12:
	;
	if v31 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+22)))
	v37 = v35 + v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+79)))
	if v38 == int32(100) {
		v22 = v31
		v24 = v37
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	return base.I64_extend_i32_u(v52)
}
func F_pg_big5_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v6 = int32(-1)
	v9 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if v9 < int32(0) {
		v12 = int32(2)
	} else {
		v12 = int32(1)
	}
	if l1 < v12 {
		v26 = v6
	} else {
		if v9 == int32(-115) {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if v16 != int32(32) {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v24 != 0 {
					v25 = v12
				} else {
					v25 = int32(-1)
				}
				v26 = v25
			} else {
				v26 = v6
			}
		} else {
			if int32(0) <= v9 {
				v26 = int32(1)
			} else {
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v24 != 0 {
					v25 = v12
				} else {
					v25 = int32(-1)
				}
				v26 = v25
			}
		}
	}
	return v26
}
func F_pg_cancel_backend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_signal_backend(m, v7, int32(2))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		switch v9 - int32(2) {
		case 0:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_cancel_backend_0), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(_a_F_pg_cancel_backend_1)
						v77 = F_errdetail(m, int32(_a_F_pg_cancel_backend_2), v5+int32(32))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_cancel_backend_3), int32(157), int32(_a_F_pg_cancel_backend_4))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 1:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_cancel_backend_0), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int64(0)
					} else {
						v26 = int32(_a_F_pg_cancel_backend_5)
						*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26
						v31 = F_errdetail(m, int32(_a_F_pg_cancel_backend_6), v5)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_cancel_backend_3), int32(143), int32(_a_F_pg_cancel_backend_4))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 2:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_cancel_backend_0), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(_a_F_pg_cancel_backend_7)
						v54 = F_errdetail(m, int32(_a_F_pg_cancel_backend_8), v5+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_cancel_backend_3), int32(150), int32(_a_F_pg_cancel_backend_4))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		default:
			m.G0 = v5 + int32(48)
			return base.I64_extend_i32_u(base.B2i32(v9 == int32(0)))
		}
	}
}
func F_pg_check_frozen(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14344(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_pg_checksum_update(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(int32(4)) <= base.Ui32(v5-int32(2)) {
		if v5 != int32(1) {
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v13 = m.Env.Pgmem_crc32c(m, v12, l1, l2)
			mBase = m.M
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v13
		}
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = F_pg_cryptohash_update(m, v15, l1, l2)
		mBase = m.M
		if int32(0) <= v16 {
			return int32(0)
		} else {
			return int32(-1)
		}
	}
}
func F_pg_class_aclcheck(m *base.Module, l0 int32, l1 int32, l2 int64) int32 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_pg_class_aclmask_ext(m, l0, l1, l2, int32(1), int32(0))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 == int64(0))
	}
}
func F_pg_class_aclmask(m *base.Module, l0 int32, l1 int32, l2 int64) int64 {
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v4 = int32(0)
	v6 = F_pg_class_aclmask_ext(m, l0, l1, l2, v4, v4)
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_pg_clean_ascii(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
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
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = F_strlen(m, l0)
	mBase = m.M
	v16 = v12<<(uint(int32(2))%32) | int32(1)
	v17 = F_palloc_extended(m, v16, l1)
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
	if v17 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v21 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v10 + int32(16)
	return v17
L6:
	;
	v22 = l0
	v23 = v21
	v24 = v3
	goto L9
L7:
	;
	v53 = v3
	goto L8
L8:
	;
	v59 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v53+v17))) = uint8(v59)
	goto L5
L9:
	;
	v29 = v24 + v17
	if base.Ui32((v23-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v53 = v47
	goto L8
L11:
	;
	v47 = v46 + v24
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v48 != 0 {
		v22 = v22 + int32(1)
		v23 = v48
		v24 = v47
		goto L9
	} else {
		goto L16
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v23 & int32(255)
	v41 = F_pg_snprintf(m, v29, v16-v24, int32(_a_F_pg_clean_ascii_0), v10)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v23)
	v46 = int32(1)
	goto L11
L15:
	;
	v46 = int32(4)
	goto L11
L16:
	;
	goto L10
}
func F_pg_clear_extended_stats(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14)+60)) = v2
	F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_extended_stats_0), v2)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v25 = F_text_to_cstring(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_extended_stats_0), int32(1))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v32 = F_text_to_cstring(m, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_extended_stats_0), int32(2))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int64(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v39 = F_text_to_cstring(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_extended_stats_0), int32(3))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int64(0)
							} else {
								v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
								v46 = F_text_to_cstring(m, v45)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int64(0)
								} else {
									F_stats_check_required_arg(m, l0, int32(_a_F_pg_clear_extended_stats_0), int32(4))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int64(0)
									} else {
										v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
										v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_clear_extended_stats[0])))
										if v55 == int32(1) {
											v60 = *(*int32)(unsafe.Add(mBase, _c_F_pg_clear_extended_stats[1]))
											v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+308))
											v63 = base.B2i32(v61 != int32(2))
											*(*uint8)(unsafe.Add(mBase, _c_F_pg_clear_extended_stats[0])) = uint8(v63)
											v65 = v63
										} else {
											v65 = int32(0)
										}
										if v65 != 0 {
											v68 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v69 = m.ExcPending
											if v69 != 0 {
												return int64(0)
											} else {
												if v68 == int32(0) {
													m.G0 = v14 - int32(-64)
													return int64(0)
												} else {
													F_errcode(m, int32(325))
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_pg_clear_extended_stats_1), int32(0))
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return int64(0)
														} else {
															F_errhint(m, int32(_a_F_pg_clear_extended_stats_2), int32(0))
															mBase = m.M
															v82 = m.ExcPending
															if v82 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_pg_clear_extended_stats_3), int32(1779), int32(_a_F_pg_clear_extended_stats_4))
																mBase = m.M
																v87 = m.ExcPending
																if v87 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v14 - int32(-64)
																	return int64(0)
																}
															}
														}
													}
												}
											}
										} else {
											v89 = F_makeRangeVar(m, v25, v32, int32(-1))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int64(0)
											} else {
												v96 = F_RangeVarGetRelidExtended(m, v89, int32(4), int32(0), int32(1138), v12+int32(-4))
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int64(0)
												} else {
													v99 = F_get_namespace_oid(m, v39, int32(1))
													mBase = m.M
													v100 = m.ExcPending
													if v100 != 0 {
														return int64(0)
													} else {
														if v99 == int32(0) {
															v105 = F_errstart(m, int32(19), int32(0))
															mBase = m.M
															v106 = m.ExcPending
															if v106 != 0 {
																return int64(0)
															} else {
																if v105 == int32(0) {
																	m.G0 = v14 - int32(-64)
																	return int64(0)
																} else {
																	F_errcode(m, int32(67137668))
																	mBase = m.M
																	v111 = m.ExcPending
																	if v111 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v14))) = v39
																		F_errmsg(m, int32(_a_F_pg_clear_extended_stats_5), v14)
																		mBase = m.M
																		v115 = m.ExcPending
																		if v115 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_pg_clear_extended_stats_3), int32(1799), int32(_a_F_pg_clear_extended_stats_4))
																			mBase = m.M
																			v120 = m.ExcPending
																			if v120 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v14 - int32(-64)
																				return int64(0)
																			}
																		}
																	}
																}
															}
														} else {
															v123 = F_table_open(m, int32(3381), int32(3))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int64(0)
															} else {
																v125 = F_get_pg_statistic_ext(m, v123, v99, v46)
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int64(0)
																} else {
																	if v125 == int32(0) {
																		F_relation_close(m, v123, int32(3))
																		mBase = m.M
																		v131 = m.ExcPending
																		if v131 != 0 {
																			return int64(0)
																		} else {
																			v134 = F_errstart(m, int32(19), int32(0))
																			mBase = m.M
																			v135 = m.ExcPending
																			if v135 != 0 {
																				return int64(0)
																			} else {
																				if v134 == int32(0) {
																					m.G0 = v14 - int32(-64)
																					return int64(0)
																				} else {
																					F_errcode(m, int32(67137668))
																					mBase = m.M
																					v140 = m.ExcPending
																					if v140 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v46
																						*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v39
																						F_errmsg(m, int32(_a_F_pg_clear_extended_stats_6), v12+int32(-48))
																						mBase = m.M
																						v147 = m.ExcPending
																						if v147 != 0 {
																							return int64(0)
																						} else {
																							F_errfinish(m, int32(_a_F_pg_clear_extended_stats_3), int32(1812), int32(_a_F_pg_clear_extended_stats_4))
																							mBase = m.M
																							v152 = m.ExcPending
																							if v152 != 0 {
																								return int64(0)
																							} else {
																								m.G0 = v14 - int32(-64)
																								return int64(0)
																							}
																						}
																					}
																				}
																			}
																		}
																	} else {
																		v153 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
																		v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+22)))
																		v155 = v153 + v154
																		v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
																		if v96 != v156 {
																			F_pfree(m, v125)
																			mBase = m.M
																			v159 = m.ExcPending
																			if v159 != 0 {
																				return int64(0)
																			} else {
																				F_relation_close(m, v123, int32(3))
																				mBase = m.M
																				v162 = m.ExcPending
																				if v162 != 0 {
																					return int64(0)
																				} else {
																					v165 = F_errstart(m, int32(19), int32(0))
																					mBase = m.M
																					v166 = m.ExcPending
																					if v166 != 0 {
																						return int64(0)
																					} else {
																						if v165 == int32(0) {
																							m.G0 = v14 - int32(-64)
																							return int64(0)
																						} else {
																							F_errcode(m, int32(50856066))
																							mBase = m.M
																							v171 = m.ExcPending
																							if v171 != 0 {
																								return int64(0)
																							} else {
																								v172 = F_get_namespace_name(m, v99)
																								mBase = m.M
																								v173 = m.ExcPending
																								if v173 != 0 {
																									return int64(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = v32
																									*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v25
																									*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v46
																									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v172
																									F_errmsg(m, int32(_a_F_pg_clear_extended_stats_7), v12+int32(-32))
																									mBase = m.M
																									v182 = m.ExcPending
																									if v182 != 0 {
																										return int64(0)
																									} else {
																										F_errfinish(m, int32(_a_F_pg_clear_extended_stats_3), int32(1830), int32(_a_F_pg_clear_extended_stats_4))
																										mBase = m.M
																										v187 = m.ExcPending
																										if v187 != 0 {
																											return int64(0)
																										} else {
																											m.G0 = v14 - int32(-64)
																											return int64(0)
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v188 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v155))))
																			v191 = F_table_open(m, int32(3429), int32(3))
																			mBase = m.M
																			v192 = m.ExcPending
																			if v192 != 0 {
																				return int64(0)
																			} else {
																				v197 = F_SearchSysCache2(m, int32(62), v188, base.I64_extend_i32_u(base.B2i32(v52 != int64(0))))
																				mBase = m.M
																				v198 = m.ExcPending
																				if v198 != 0 {
																					return int64(0)
																				} else {
																					if v197 != 0 {
																						F_simple_heap_delete(m, v191, v197+int32(4))
																						mBase = m.M
																						v202 = m.ExcPending
																						if v202 != 0 {
																							return int64(0)
																						} else {
																							F_ReleaseCatCache(m, v197)
																							mBase = m.M
																							v204 = m.ExcPending
																							if v204 != 0 {
																								return int64(0)
																							} else {
																								F_relation_close(m, v191, int32(3))
																								mBase = m.M
																								v207 = m.ExcPending
																								if v207 != 0 {
																									return int64(0)
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v209 = m.ExcPending
																									if v209 != 0 {
																										return int64(0)
																									} else {
																										F_pfree(m, v125)
																										mBase = m.M
																										v211 = m.ExcPending
																										if v211 != 0 {
																											return int64(0)
																										} else {
																											F_relation_close(m, v123, int32(3))
																											mBase = m.M
																											v214 = m.ExcPending
																											if v214 != 0 {
																												return int64(0)
																											} else {
																												m.G0 = v14 - int32(-64)
																												return int64(0)
																											}
																										}
																									}
																								}
																							}
																						}
																					} else {
																						F_relation_close(m, v191, int32(3))
																						mBase = m.M
																						v207 = m.ExcPending
																						if v207 != 0 {
																							return int64(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v209 = m.ExcPending
																							if v209 != 0 {
																								return int64(0)
																							} else {
																								F_pfree(m, v125)
																								mBase = m.M
																								v211 = m.ExcPending
																								if v211 != 0 {
																									return int64(0)
																								} else {
																									F_relation_close(m, v123, int32(3))
																									mBase = m.M
																									v214 = m.ExcPending
																									if v214 != 0 {
																										return int64(0)
																									} else {
																										m.G0 = v14 - int32(-64)
																										return int64(0)
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
						}
					}
				}
			}
		}
	}
}
func F_pg_client_encoding(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pg_client_encoding[0]))
	v6 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v5))))
	v7 = F_DirectFunctionCall1Coll(m, int32(534), int32(0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_pg_control_system(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v11 = F_get_call_result_type(m, l0, int32(0), v6+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		if v11 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_system[0]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_system[1]))
				v28 = F_get_controlfile(m, v25, v6+int32(7))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_system[0]))
					F_LWLockRelease(m, v31+int32(1152))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int64(0)
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_pg_control_system_0), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_control_system_1), int32(50), int32(_a_F_pg_control_system_2))
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+8)))
							v40 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v39
							v43 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+12)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v43
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v40)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v47
							v51 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
							v56 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v56)
							*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v51*int64(1000000) - int64(946684800000000)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
							v64 = F_heap_form_tuple(m, v59, v6+int32(16), v6+int32(12))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int64(0)
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
								v67 = F_HeapTupleHeaderGetDatum(m, v66)
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
									return int64(0)
								} else {
									m.G0 = v6 + int32(48)
									return v67
								}
							}
						}
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_control_system_3), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_control_system_1), int32(42), int32(_a_F_pg_control_system_2))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int64(0)
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
func F_pg_convert_to(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = int32(0)
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_convert_to[0]))
	v11 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v10))))
	v12 = F_DirectFunctionCall1Coll(m, int32(534), v5, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = F_DirectFunctionCall3Coll(m, int32(1851), v5, v6, v12, v3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			return v16
		}
	}
}
func F_pg_database_encoding_max_length(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_pg_database_encoding_max_length[0]))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v3*int32(28))+uint32(_c_F_pg_database_encoding_max_length[1])))
	return v8
}
func F_pg_ddl_command_out(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_pg_ddl_command_out_0), int32(358), int32(_a_F_pg_ddl_command_out_1), int32(_a_F_pg_ddl_command_out_2), int32(_a_F_pg_ddl_command_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_pg_dependencies_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
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
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v254 int32
	_ = v254
	var v273 int32
	_ = v273
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v364 int32
	_ = v364
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int64
	_ = v484
	v2 = int32(0)
	v16 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(224)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+200)) = v16
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+192)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+188)) = v2
	*(*int64)(unsafe.Add(mBase, uint32(v19)+172)) = v16
	*(*int32)(unsafe.Add(mBase, uint32(v19)+168)) = v21
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+186)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+184)) = uint16(v2)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+160)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+156)) = int32(1608)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+144)) = int32(1609)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+140)) = int32(1610)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+136)) = int32(1611)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+132)) = int32(1612)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+180)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v19)+164)) = int32(1613)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+152)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+148)) = int32(1614)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+128)) = v19 + int32(168)
	v59 = F_strlen(m, v21)
	mBase = m.M
	v62 = F_makeJsonLexContextCstringLen(m, v2, v21, v59, int32(6), int32(1))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v68 = F_pg_parse_json(m, v62, v19+int32(128))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_freeJsonLexContext(m, v62)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	if v68 != 0 {
		v432 = v2
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v19)+176))
	F_list_free_deep(m, v446)
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L69
	}
L6:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v19)+176))
	if v72 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+4))
	v75 = v73
	goto L9
L8:
	;
	v75 = int32(0)
	goto L9
L9:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	switch v76 {
	case 0:
		goto L12
	default:
		goto L11
	case 7:
		goto L13
	}
L10:
	;
	v145 = F_palloc0(m, v75<<(uint(int32(2))%32)+int32(12))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L30
	}
L11:
	;
	v114 = int32(0)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v19)+180))
	v116 = F_errsave_start(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L24
	}
L12:
	;
	v90 = int32(0)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v19)+180))
	v92 = F_errsave_start(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v75 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errmsg_internal(m, int32(_a_F_pg_dependencies_in_0), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_pg_dependencies_in_1), int32(665), int32(_a_F_pg_dependencies_in_2))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v92 == int32(0) {
		v432 = v90
		goto L5
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+112)) = v21
	F_errmsg(m, int32(_a_F_pg_dependencies_in_3), v19+int32(112))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v107 = F_errdetail(m, int32(_a_F_pg_dependencies_in_4), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F_errsave_finish(m, v91, int32(_a_F_pg_dependencies_in_1), int32(673), int32(_a_F_pg_dependencies_in_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v432 = v90
	goto L5
L24:
	;
	if v116 == int32(0) {
		v432 = v114
		goto L5
	} else {
		goto L25
	}
L25:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v21
	F_errmsg(m, int32(_a_F_pg_dependencies_in_3), v19+int32(32))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v19)+172))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v129
	v134 = F_errdetail(m, int32(_a_F_pg_dependencies_in_5), v19+int32(16))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errsave_finish(m, v115, int32(_a_F_pg_dependencies_in_1), int32(681), int32(_a_F_pg_dependencies_in_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v432 = v114
	goto L5
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v75
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = int64(7320410668)
	if int32(0) < v75 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	F_pfree(m, v145)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L68
	}
L32:
	;
	v310 = v19 + int32(208)
	F_initStringInfo(m, v310)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L53
	}
L33:
	;
	v153 = v145 + int32(12)
	v160 = v2
	goto L36
L34:
	;
	goto L35
L35:
	;
	v291 = F_statext_dependencies_serialize(m, v145)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L52
	}
L36:
	;
	v171 = v160 << (uint(int32(2)) % 32)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v19)+176))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v174+v171)))
	*(*int32)(unsafe.Add(mBase, uint32(v153+v171))) = v176
	if v160 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L35
L38:
	;
	v179 = v176 + int32(10)
	v181 = int32(*(*int16)(unsafe.Add(mBase, uint32(v176)+8)))
	v192 = int32(0)
	goto L41
L39:
	;
	goto L40
L40:
	;
	v273 = v160 + int32(1)
	if v273 != v75 {
		v160 = v273
		goto L36
	} else {
		goto L51
	}
L41:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v153+v192<<(uint(int32(2))%32))))
	v204 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+8)))
	if v181&int32(_a_F_pg_dependencies_in_6) != v204 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L40
L43:
	;
	v254 = v192 + int32(1)
	if v254 != v160 {
		v192 = v254
		goto L41
	} else {
		goto L50
	}
L44:
	;
	if v181 <= int32(0) {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	v213 = int32(0)
	goto L46
L46:
	;
	v228 = v213 << (uint(int32(1)) % 32)
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v179+v228))))
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203+int32(10)+v228))))
	if v230 != v232 {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	goto L32
L48:
	;
	v235 = v213 + int32(1)
	if v181 != v235 {
		v213 = v235
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	goto L42
L51:
	;
	goto L37
L52:
	;
	v414 = v291
	goto L31
L53:
	;
	v313 = int32(*(*int16)(unsafe.Add(mBase, uint32(v176)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+96)) = v313
	F_appendStringInfo(m, v310, int32(_a_F_pg_dependencies_in_7), v19+int32(96))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v321 = v176 + int32(8)
	v322 = int32(*(*int16)(unsafe.Add(mBase, uint32(v176)+8)))
	if int32(2) < v322 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v328 = int32(1)
	goto L58
L56:
	;
	v364 = v322
	goto L57
L57:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v19)+208))
	v380 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321+v364<<(uint(int32(1))%32)))))
	v381 = int32(0)
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v19)+180))
	v383 = F_errsave_start(m, v382)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L62
	}
L58:
	;
	v345 = int32(*(*int16)(unsafe.Add(mBase, uint32(v179+v328<<(uint(int32(1))%32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+80)) = v345
	F_appendStringInfo(m, v19+int32(208), int32(_a_F_pg_dependencies_in_8), v19+int32(80))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L60
	}
L59:
	;
	v364 = v356
	goto L57
L60:
	;
	v354 = int32(1)
	v355 = v328 + v354
	v356 = int32(*(*int16)(unsafe.Add(mBase, uint32(v321))))
	if v355 < v356-v354 {
		v328 = v355
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	if v383 == int32(0) {
		v414 = v381
		goto L31
	} else {
		goto L63
	}
L63:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v21
	F_errmsg(m, int32(_a_F_pg_dependencies_in_3), v19-int32(-64))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = int32(_a_F_pg_dependencies_in_9)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v376
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(_a_F_pg_dependencies_in_10)
	v405 = F_errdetail(m, int32(_a_F_pg_dependencies_in_11), v19+int32(48))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errsave_finish(m, v382, int32(_a_F_pg_dependencies_in_1), int32(718), int32(_a_F_pg_dependencies_in_2))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v414 = v381
	goto L31
L68:
	;
	v432 = v414
	goto L5
L69:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v19)+188))
	F_list_free(m, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v432 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v484 = base.I64_extend_i32_u(v432)
	goto L73
L72:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v19)+180))
	if v453 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	m.G0 = v19 + int32(224)
	return v484
L74:
	;
	v480 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v480)
	v484 = int64(0)
	goto L73
L75:
	;
	v460 = F_errsave_start(m, v453)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L79
	}
L76:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	if v456 != int32(453) {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453)+4)))
	if v459 != 0 {
		goto L74
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	if v460 == int32(0) {
		goto L74
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
	F_errmsg(m, int32(_a_F_pg_dependencies_in_3), v19)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v473 = F_errdetail(m, int32(_a_F_pg_dependencies_in_12), int32(0))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errsave_finish(m, v453, int32(_a_F_pg_dependencies_in_1), int32(803), int32(_a_F_pg_dependencies_in_13))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	goto L74
}
func F_pg_describe_object(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = base.I32_wrap_i64(v11)
	if v10|v12 == int32(0) {
		v16 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v16)
		v39 = int64(0)
		m.G0 = v8 + int32(16)
		return v39
	} else {
		v19 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		*(*uint32)(unsafe.Add(mBase, uint32(v8)+12)) = uint32(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v12
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v10
		v26 = F_getObjectDescription(m, v8+int32(4), int32(1))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int64(0)
		} else {
			if v26 == int32(0) {
				v32 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v32)
				v39 = int64(0)
				m.G0 = v8 + int32(16)
				return v39
			} else {
				v35 = F_cstring_to_text(m, v26)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int64(0)
				} else {
					v39 = base.I64_extend_i32_u(v35)
					m.G0 = v8 + int32(16)
					return v39
				}
			}
		}
	}
}
func F_pg_encrypt_iv(m *base.Module, l0 int32) int64 {
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L2
	} else {
		goto L113
	}
L2:
	;
	return int64(0)
L3:
	;
	v21 = int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v25 = v23 & v21
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v21
	goto L6
L5:
	;
	v26 = int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v56 = F_downcase_truncate_identifier(m, v17+v26, v54, int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v33 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v44 = int32(1)
	if v25 != 0 {
		v54 = int32(base.Ui32(v23)>>(uint(v44)%32)) - v44
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v36 = int32(16)
	goto L13
L12:
	;
	v36 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v43 = int32(4)
	goto L16
L15:
	;
	v43 = v36
	goto L16
L16:
	;
	v54 = v43
	goto L7
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v60 = F_px_find_combo(m, v56, v14+int32(28))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v60 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v56)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L2
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L2
	} else {
		goto L95
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v68 = F_pg_detoast_datum_packed(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v71 = F_pg_detoast_datum_packed(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v74 = F_pg_detoast_datum_packed(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v76 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v106 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	if v82 == int32(18) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v93 = int32(1)
	if v76&v93 != 0 {
		v105 = int32(base.Ui32(v76)>>(uint(v93)%32)) - v93
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v85 = int32(16)
	goto L33
L32:
	;
	v85 = int32(0)
	goto L33
L33:
	;
	if base.Ui32((v82-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v92 = int32(4)
	goto L36
L35:
	;
	v92 = v85
	goto L36
L36:
	;
	v105 = v92
	goto L27
L37:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	v105 = int32(base.Ui32(v99)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v136 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v112 == int32(18) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v123 = int32(1)
	if v106&v123 != 0 {
		v135 = int32(base.Ui32(v106)>>(uint(v123)%32)) - v123
		goto L38
	} else {
		goto L48
	}
L42:
	;
	v115 = int32(16)
	goto L44
L43:
	;
	v115 = int32(0)
	goto L44
L44:
	;
	if base.Ui32((v112-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v122 = int32(4)
	goto L47
L46:
	;
	v122 = v115
	goto L47
L47:
	;
	v135 = v122
	goto L38
L48:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	v135 = int32(base.Ui32(v129)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	v167 = m.T0[v166].(func(*base.Module, int32, int32) int32)(m, v66, v105)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L60
	}
L50:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v142 == int32(18) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v153 = int32(1)
	if v136&v153 != 0 {
		v165 = int32(base.Ui32(v136)>>(uint(v153)%32)) - v153
		goto L49
	} else {
		goto L59
	}
L53:
	;
	v145 = int32(16)
	goto L55
L54:
	;
	v145 = int32(0)
	goto L55
L55:
	;
	if base.Ui32((v142-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v152 = int32(4)
	goto L58
L57:
	;
	v152 = v145
	goto L58
L58:
	;
	v165 = v152
	goto L49
L59:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	v165 = int32(base.Ui32(v159)>>(uint(int32(2))%32)) - int32(4)
	goto L49
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v167
	v172 = F_palloc(m, v167+int32(4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v174 = int32(1)
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v176&v174 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v179 = v174
	goto L64
L63:
	;
	v179 = int32(4)
	goto L64
L64:
	;
	v181 = int32(1)
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v183&v181 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v186 = v181
	goto L67
L66:
	;
	v186 = int32(4)
	goto L67
L67:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v189 = m.T0[v188].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v66, v71+v179, v135, v74+v186, v165)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v189 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	m.T0[v191].(func(*base.Module, int32))(m, v66)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v194 = int32(1)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	if v196&v194 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v287 = v189
	goto L1
L73:
	;
	v199 = v194
	goto L75
L74:
	;
	v199 = int32(4)
	goto L75
L75:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v206 = m.T0[v205].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v66, v68+v199, v105, v172+int32(4), v14+int32(28))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	m.T0[v208].(func(*base.Module, int32))(m, v66)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v206 != 0 {
		v287 = v206
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v211<<(uint(int32(2))%32) + int32(16)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v217 != v68 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_pfree(m, v68)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v221 != v71 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v71)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v225 != v74 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_pfree(m, v74)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L2
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v229 != v17 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	goto L89
L91:
	;
	F_pfree(m, v17)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L2
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	m.G0 = v14 + int32(32)
	return base.I64_extend_i32_u(v172)
L94:
	;
	goto L93
L95:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L2
	} else {
		goto L96
	}
L96:
	;
	if v60 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v56
	F_errmsg(m, int32(_a_F_pg_encrypt_iv_0), v14+int32(16))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L2
	} else {
		goto L111
	}
L98:
	;
	v274 = int32(_a_F_pg_encrypt_iv_1)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v253 = int32(_a_F_pg_encrypt_iv_2)
	goto L102
L101:
	;
	v274 = v268
	goto L97
L102:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v253)+8))
	if v60 != v256 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v253)+12))
	v268 = v266
	goto L101
L104:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v253)+20))
	if v258 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	v274 = int32(_a_F_pg_encrypt_iv_3)
	goto L97
L108:
	;
	goto L109
L109:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v253)+16))
	if v60 != v262 {
		v253 = v253 + int32(16)
		goto L102
	} else {
		goto L110
	}
L110:
	;
	v268 = v258
	goto L101
L111:
	;
	F_errfinish(m, int32(_a_F_pg_encrypt_iv_4), int32(513), int32(_a_F_pg_encrypt_iv_5))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L2
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	if v287 == int32(0) {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v324
	F_errmsg(m, int32(_a_F_pg_encrypt_iv_6), v14)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L2
	} else {
		goto L129
	}
L116:
	;
	v324 = int32(_a_F_pg_encrypt_iv_1)
	goto L115
L117:
	;
	goto L118
L118:
	;
	v303 = int32(_a_F_pg_encrypt_iv_2)
	goto L120
L119:
	;
	v324 = v318
	goto L115
L120:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+8))
	if v287 != v306 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v303)+12))
	v318 = v316
	goto L119
L122:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v303)+20))
	if v308 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	goto L121
L125:
	;
	v324 = int32(_a_F_pg_encrypt_iv_3)
	goto L115
L126:
	;
	goto L127
L127:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v303)+16))
	if v287 != v312 {
		v303 = v303 + int32(16)
		goto L120
	} else {
		goto L128
	}
L128:
	;
	v318 = v308
	goto L119
L129:
	;
	F_errfinish(m, int32(_a_F_pg_encrypt_iv_4), int32(387), int32(_a_F_pg_encrypt_iv_7))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L2
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_euccn_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	v5 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if int32(0) <= v5 {
		v8 = int32(1)
	} else {
		v8 = int32(2)
	}
	if v5&int32(254) == int32(142) {
		v13 = int32(3)
	} else {
		v13 = v8
	}
	return v13
}
func F_pg_eucjp2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v19 - int32(142) {
	case 0:
		goto L10
	case 1:
		goto L9
	default:
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = int32(0)
	return v77
L6:
	;
	goto L5
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v61
	v66 = v18 + int32(1)
	v68 = v14 + int32(4)
	v69 = v15 + v62
	if int32(0) < v69 {
		v13 = v63
		v14 = v68
		v15 = v69
		v18 = v66
		goto L4
	} else {
		goto L18
	}
L8:
	;
	if v19 == int32(0) {
		v73 = v14
		v77 = v18
		goto L6
	} else {
		goto L13
	}
L9:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v73 = v14
		v77 = v18
		goto L6
	} else {
		goto L12
	}
L10:
	;
	if v15 == int32(1) {
		v73 = v14
		v77 = v18
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v61 = v24 | int32(_a_F_pg_eucjp2wchar_with_len_0)
	v62 = int32(-2)
	v63 = v13 + int32(2)
	goto L7
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v36 = v32<<(uint(int32(8))%32) | int32(_a_F_pg_eucjp2wchar_with_len_1)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v61 = v36 | v38
	v62 = int32(-3)
	v63 = v13 + int32(3)
	goto L7
L13:
	;
	if base.I32_extend8_s(v19) < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v15 == int32(1) {
		v73 = v14
		v77 = v18
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v61 = v19
	v62 = int32(-1)
	v63 = v13 + int32(1)
	goto L7
L17:
	;
	v51 = v19 << (uint(int32(8)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v51
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v61 = v51 | v53
	v62 = int32(-2)
	v63 = v13 + int32(2)
	goto L7
L18:
	;
	v73 = v68
	v77 = v66
	goto L6
}
func F_pg_euctw_mblen(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	switch v4 - int32(142) {
	case 0:
		v14 = int32(4)
		v16 = v14
	case 1:
		v16 = int32(3)
	default:
		if int32(0) <= base.I32_extend8_s(v4) {
			v13 = int32(1)
		} else {
			v13 = int32(2)
		}
		v14 = v13
		v16 = v14
	}
	return v16
}
func F_pg_euctw_verifystr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	if l1 <= int32(0) {
		v63 = l0
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v63 - l0
L2:
	;
	v9 = l1
	v10 = l0
	goto L3
L3:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v14 = base.I32_extend8_s(v13)
	if int32(0) <= v14 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v63 = v57
	goto L1
L5:
	;
	v57 = v56 + v10
	v58 = v9 - v56
	if int32(0) < v58 {
		v9 = v58
		v10 = v57
		goto L3
	} else {
		goto L18
	}
L6:
	;
	if v14 == int32(0) {
		v63 = v10
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	switch v13 - int32(142) {
	case 0:
		goto L11
	case 1:
		v63 = v10
		goto L1
	default:
		goto L10
	}
L9:
	;
	v56 = int32(1)
	goto L5
L10:
	;
	if v9 == int32(1) {
		v63 = v10
		goto L1
	} else {
		goto L16
	}
L11:
	;
	if base.Ui32(v9) < base.Ui32(int32(4)) {
		v63 = v10
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32((v24+int32(88))&int32(255)) < base.Ui32(int32(249)) {
		v63 = v10
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+2)))
	if base.Ui32(int32(93)) < base.Ui32((v31+int32(95))&int32(255)) {
		v63 = v10
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+3)))
	if base.Ui32(int32(94)) <= base.Ui32((v38+int32(95))&int32(255)) {
		v63 = v10
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v56 = int32(4)
	goto L5
L16:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if base.Ui32(int32(93)) < base.Ui32((v48+int32(95))&int32(255)) {
		v63 = v10
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v56 = int32(2)
	goto L5
L18:
	;
	goto L4
}
func F_pg_filenode_relation(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v3 == int32(0) {
		v6 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v6)
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v11 = F_RelidByRelfilenumber(m, v10, v3)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			if v11 == int32(0) {
				v17 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v11)
			}
		}
	}
}
func F_pg_fsync(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v4 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_fsync[0])))
	if v4 != int32(1) {
		v18 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v18
L2:
	;
	goto L3
L3:
	;
	v9 = F_fsync(m, l0)
	mBase = m.M
	if v9 != int32(-1) {
		v18 = v9
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v18 = int32(-1)
	goto L1
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_fsync[1]))
	if v13 == int32(27) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
}
func F_pg_gb18030_verifychar(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v4 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0))))
	if int32(0) <= v4 {
		return int32(1)
	} else {
		v10 = v4 & int32(255)
		if int32(4) <= l1 {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
				if base.B2i32(v10 == int32(128))|base.B2i32(v10 == int32(255)) != 0 {
					v69 = int32(-1)
					return v69
				} else {
					v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
					if base.Ui32((v54+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
						return int32(2)
					} else {
						if int32(-2) < v54 {
							v67 = int32(-1)
						} else {
							v67 = int32(2)
						}
						v69 = v67
						return v69
					}
				}
			} else {
				if base.B2i32(v10 == int32(128))|base.B2i32(v10 == int32(255)) != 0 {
					return int32(-1)
				} else {
					v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
					if base.Ui32((v25+int32(1))&int32(255)) < base.Ui32(int32(130)) {
						return int32(-1)
					} else {
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
						if base.Ui32(int32(10)) <= base.Ui32((v32-int32(48))&int32(255)) {
							return int32(-1)
						} else {
							return int32(4)
						}
					}
				}
			}
		} else {
			if int32(2) <= l1 {
				if base.B2i32(v10 == int32(128))|base.B2i32(v10 == int32(255)) != 0 {
					v69 = int32(-1)
					return v69
				} else {
					v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+1)))
					if base.Ui32((v54+int32(-64))&int32(255)) < base.Ui32(int32(63)) {
						return int32(2)
					} else {
						if int32(-2) < v54 {
							v67 = int32(-1)
						} else {
							v67 = int32(2)
						}
						v69 = v67
						return v69
					}
				}
			} else {
				return int32(-1)
			}
		}
	}
}
func F_pg_get_object_address(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v458 int32
	_ = v458
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
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v505 int32
	_ = v505
	var v506 int64
	_ = v506
	var v507 int64
	_ = v507
	var v508 int64
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int64
	_ = v535
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v554 int32
	_ = v554
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	v2 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(288)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_text_to_cstring(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v33 = v2
	goto L15
L5:
	;
	switch v82 - int32(2) {
	case 0, 1:
		goto L99
	default:
		goto L97
	case 3, 9, 11, 29, 31, 42:
		goto L101
	case 22, 24:
		goto L100
	case 23:
		goto L98
	case 30, 49:
		goto L102
	}
L6:
	;
	v340 = F_textarray_to_strvaluelist(m, v26)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L91
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L87
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L83
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L79
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L75
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L71
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L67
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L63
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L59
	}
L15:
	;
	v41 = v33 << (uint(int32(3)) % 32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_pg_get_object_address[0])))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if base.B2i32(v45 == int32(0))|base.B2i32(v45 != v48) != 0 {
		v66 = v45
		v67 = v48
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v33))%64)&int64(4398046543432) != int64(0) {
		goto L13
	} else {
		goto L28
	}
L17:
	;
	if v66-v67 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	v51 = v42
	v52 = v18
	goto L20
L20:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v56 == int32(0) {
		v66 = v56
		v67 = v55
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v66 = v56
	v67 = v55
	goto L18
L22:
	;
	v59 = int32(1)
	if v56 == v55 {
		v51 = v51 + v59
		v52 = v52 + v59
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v70 = v33 + int32(1)
	if v70 != int32(60) {
		v33 = v70
		goto L15
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L16
L27:
	;
	goto L14
L28:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v41)+uint32(_c_F_pg_get_object_address[1])))
	if v33 == int32(58) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	if v33&int32(2147483646) == int32(32) {
		goto L48
	} else {
		goto L49
	}
L30:
	;
	v134 = F_textarray_to_strvaluelist(m, v23)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L46
	}
L31:
	;
	F_deconstruct_array_builtin(m, v23, int32(25), v15+int32(256), v15+int32(284), v15+int32(248))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L41
	}
L32:
	;
	F_deconstruct_array_builtin(m, v23, int32(25), v15+int32(256), v15+int32(284), v15+int32(248))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L36
	}
L33:
	;
	switch v82 - int32(5) {
	case 0, 8:
		goto L32
	case 1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13, 14, 15, 16:
		goto L30
	case 17:
		goto L31
	default:
		goto L34
	}
L34:
	;
	if v82 != int32(50) {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v15)+248))
	if v98 != int32(1) {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v15)+284))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v102 == int32(1) {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v15)+256))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v107 = F_text_to_cstring(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v110 = F_typeStringToTypeName(m, v107, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v138 = v2
	v139 = v110
	goto L29
L41:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v15)+248))
	if v121 != int32(1) {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v15)+284))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v125 == int32(1) {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v15)+256))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v130 = F_text_to_cstring(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v132 = F_makeFloat(m, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v337 = v132
	v338 = v2
	v339 = v2
	goto L6
L46:
	;
	if v134 == int32(0) {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v138 = v134
	v139 = v2
	goto L29
L48:
	;
	F_deconstruct_array_builtin(m, v26, int32(25), v15+int32(256), v15+int32(284), v15+int32(248))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	switch v82 - int32(1) {
	case 0, 4, 18, 24, 28, 34:
		goto L48
	default:
		v337 = int32(0)
		v338 = v138
		v339 = v139
		goto L6
	}
L50:
	;
	v157 = int32(0)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v15)+248))
	if v159 <= v157 {
		v344 = v157
		v345 = v138
		v346 = v157
		v349 = v139
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v167 = v157
	v168 = int32(0)
	goto L52
L52:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v15)+284))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v168))))
	if v177 == int32(1) {
		goto L7
	} else {
		goto L54
	}
L53:
	;
	v344 = v157
	v345 = v138
	v346 = v190
	v349 = v139
	goto L5
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v15)+256))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v180+v168<<(uint(int32(3))%32))))
	v185 = F_text_to_cstring(m, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v188 = F_typeStringToTypeName(m, v185, int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v190 = F_lappend(m, v167, v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v193 = v168 + int32(1)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v15)+248))
	if v193 < v194 {
		v167 = v190
		v168 = v193
		goto L52
	} else {
		goto L58
	}
L58:
	;
	goto L53
L59:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+208)) = v18
	F_errmsg(m, int32(_a_F_pg_get_object_address_0), v15+int32(208))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2670), int32(_a_F_pg_get_object_address_2))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+192)) = v18
	F_errmsg(m, int32(_a_F_pg_get_object_address_3), v15+int32(192))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2179), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+144)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_5), v15+int32(144))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2198), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errmsg(m, int32(_a_F_pg_get_object_address_6), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2202), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+176)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_5), v15+int32(176))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2215), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errmsg(m, int32(_a_F_pg_get_object_address_7), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2219), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+160)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_8), v15+int32(160))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2228), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	F_errmsg(m, int32(_a_F_pg_get_object_address_6), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2257), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	v344 = v337
	v345 = v338
	v346 = v340
	v349 = v339
	goto L5
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L1
	} else {
		goto L173
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L1
	} else {
		goto L169
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L165
	}
L95:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L161
	}
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L1
	} else {
		goto L157
	}
L97:
	;
	switch v82 {
	case 0, 9, 14, 15, 16, 17, 21, 27, 30, 34, 37, 39, 43:
		goto L135
	case 1, 19, 25, 29, 35:
		goto L126
	case 2, 3:
		goto L129
	case 4, 6, 7, 8, 10, 18, 20, 23, 24, 26, 28, 36, 38, 40, 41, 42, 45, 46, 47, 48, 49, 52:
		goto L128
	case 5, 13, 44:
		goto L133
	case 11:
		goto L130
	case 12, 50:
		goto L134
	default:
		v474 = v344
		goto L127
	case 31, 33:
		goto L132
	case 32, 51:
		goto L131
	}
L98:
	;
	if v346 == int32(0) {
		goto L94
	} else {
		goto L123
	}
L99:
	;
	if v345 == int32(0) {
		goto L95
	} else {
		goto L121
	}
L100:
	;
	if v345 != 0 {
		goto L113
	} else {
		goto L114
	}
L101:
	;
	if v346 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	if v345 == int32(0) {
		goto L96
	} else {
		goto L103
	}
L103:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v358 != int32(1) {
		goto L96
	} else {
		goto L104
	}
L104:
	;
	goto L101
L105:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v361 == int32(1) {
		goto L97
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	goto L107
L109:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+80)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_9), v15+int32(80))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2292), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if int32(1) < v383 {
		goto L97
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L117
	}
L116:
	;
	goto L115
L117:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+96)) = int32(2)
	F_errmsg(m, int32(_a_F_pg_get_object_address_8), v15+int32(96))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2299), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v407 <= int32(2) {
		goto L95
	} else {
		goto L122
	}
L122:
	;
	goto L98
L123:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v346)+4))
	if v412 != int32(2) {
		goto L94
	} else {
		goto L124
	}
L124:
	;
	goto L97
L125:
	;
	F_get_object_address(m, v15+int32(256), v82, v496, v15+int32(248), int32(1), int32(0))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L1
	} else {
		goto L148
	}
L126:
	;
	v490 = F_palloc0(m, int32(20))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L147
	}
L127:
	;
	if v474 != 0 {
		v496 = v474
		goto L125
	} else {
		goto L143
	}
L128:
	;
	v474 = v345
	goto L127
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+216)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v15)+220)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v15)+60)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v15)+56)) = v346
	v472 = F_list_make2_impl(m, v15+int32(60), v15+int32(56))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L142
	}
L130:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v460)))
	v462 = F_lcons(m, v461, v345)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L141
	}
L131:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v446)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+228)) = v447
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+224)) = v450
	*(*int32)(unsafe.Add(mBase, uint32(v15)+52)) = v447
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v450
	v458 = F_list_make2_impl(m, v15+int32(52), v15+int32(48))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L140
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+236)) = v345
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+232)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v15)+44)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v15)+40)) = v436
	v444 = F_list_make2_impl(m, v15+int32(44), v15+int32(40))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L139
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+244)) = v349
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v346)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v423)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+240)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v424
	v432 = F_list_make2_impl(m, v15+int32(36), v15+int32(32))
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L138
	}
L134:
	;
	v474 = v349
	goto L127
L135:
	;
	if v345 == int32(0) {
		goto L93
	} else {
		goto L136
	}
L136:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v417 != int32(1) {
		goto L93
	} else {
		goto L137
	}
L137:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	v474 = v421
	goto L127
L138:
	;
	v474 = v432
	goto L127
L139:
	;
	v474 = v444
	goto L127
L140:
	;
	v474 = v458
	goto L127
L141:
	;
	v474 = v462
	goto L127
L142:
	;
	v474 = v472
	goto L127
L143:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v82
	F_errmsg_internal(m, int32(_a_F_pg_get_object_address_10), v15)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2412), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v490)+8)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v490)+4)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v490))) = int32(153)
	v496 = v490
	goto L125
L148:
	;
	v506 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+264)))
	v507 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+260)))
	v508 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v15)+256)))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v15)+248))
	if v509 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	F_relation_close(m, v509, int32(1))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v516 = F_get_call_result_type(m, l0, int32(0), v15+int32(284))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L153
	}
L152:
	;
	goto L151
L153:
	;
	if v516 != int32(1) {
		goto L92
	} else {
		goto L154
	}
L154:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v15)+272)) = v506
	*(*int64)(unsafe.Add(mBase, uint32(v15)+264)) = v507
	*(*int64)(unsafe.Add(mBase, uint32(v15)+256)) = v508
	v523 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+255)) = uint8(v523)
	*(*uint16)(unsafe.Add(mBase, uint32(v15)+253)) = uint16(v523)
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v15)+284))
	v532 = F_heap_form_tuple(m, v527, v15+int32(256), v15+int32(253))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v532)+16))
	v535 = F_HeapTupleHeaderGetDatum(m, v534)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	m.G0 = v15 + int32(288)
	return v535
L157:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+64)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_5), v15-int32(-64))
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2280), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+112)) = int32(3)
	F_errmsg(m, int32(_a_F_pg_get_object_address_8), v15+int32(112))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2306), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+128)) = int32(2)
	F_errmsg(m, int32(_a_F_pg_get_object_address_9), v15+int32(128))
	mBase = m.M
	v592 = m.ExcPending
	if v592 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2313), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_5), v15+int32(16))
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2365), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_errmsg_internal(m, int32(_a_F_pg_get_object_address_11), int32(0))
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2422), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_replica_identity_index(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_table_open(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = F_RelationGetReplicaIndex(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			F_relation_close(m, v6, int32(1))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				if v10 != 0 {
					return base.I64_extend_i32_u(v10)
				} else {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
					return int64(0)
				}
			}
		}
	}
}
func F_pg_get_triggerdef_worker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v215 int32
	_ = v215
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
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
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int64
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int64
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v520 int32
	_ = v520
	var v523 int64
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v561 int32
	_ = v561
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v582 int32
	_ = v582
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v652 int32
	_ = v652
	v16 = m.G0
	v18 = v16 - int32(336)
	m.G0 = v18
	v22 = F_table_open(m, int32(2620), int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v27 = v18 + int32(256)
	F_ScanKeyInit(m, v27, int32(1), int32(3), int32(184), base.I64_extend_i32_u(l0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v35 = int32(1)
	v38 = F_systable_beginscan(m, v22, int32(2702), v35, int32(0), v35, v27)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v18 + int32(336)
	return v652
L5:
	;
	v40 = F_systable_getnext(m, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v40 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_systable_endscan(m, v38)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+16))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
	v52 = v18 + int32(320)
	F_initStringInfo(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	F_relation_close(m, v22, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v652 = int32(0)
	goto L4
L12:
	;
	v55 = v49 + v50
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+92))
	v59 = F_quote_identifier(m, v55+int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+116)) = v59
	if v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v64 = int32(_a_F_pg_get_triggerdef_worker_0)
	goto L16
L15:
	;
	v64 = int32(_a_F_pg_get_triggerdef_worker_1)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+112)) = v64
	F_appendStringInfo(m, v52, int32(_a_F_pg_get_triggerdef_worker_2), v18+int32(112))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+80)))
	v74 = v72 & int32(66)
	switch v74 {
	case 0:
		goto L19
	case 1:
		goto L21
	case 2:
		v95 = int32(_a_F_pg_get_triggerdef_worker_3)
		goto L18
	default:
		goto L22
	}
L18:
	;
	v97 = v18 + int32(320)
	F_appendStringInfoString(m, v97, v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L27
	}
L19:
	;
	v95 = int32(_a_F_pg_get_triggerdef_worker_4)
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L24
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L22:
	;
	switch v74 - int32(65) {
	case 0:
		goto L21
	case 1:
		goto L20
	default:
		goto L23
	}
L23:
	;
	v95 = int32(_a_F_pg_get_triggerdef_worker_5)
	goto L18
L24:
	;
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+80)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+96)) = v82
	F_errmsg_internal(m, int32(_a_F_pg_get_triggerdef_worker_6), v18+int32(96))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_pg_get_triggerdef_worker_7), int32(960), int32(_a_F_pg_get_triggerdef_worker_8))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+80)))
	if v100&int32(4) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L28:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if l1 != 0 {
		goto L64
	} else {
		goto L65
	}
L29:
	;
	F_appendStringInfoString(m, v97, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L62
	}
L30:
	;
	if v100&int32(32) == int32(0) {
		goto L28
	} else {
		goto L61
	}
L31:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+80)))
	if v215&int32(32) == int32(0) {
		goto L28
	} else {
		goto L60
	}
L32:
	;
	F_appendStringInfoString(m, v97, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L45
	}
L33:
	;
	if v100&int32(16) == int32(0) {
		goto L30
	} else {
		goto L44
	}
L34:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+80)))
	if v124&int32(16) == int32(0) {
		goto L31
	} else {
		goto L43
	}
L35:
	;
	F_appendStringInfoString(m, v97, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L42
	}
L36:
	;
	if v100&int32(8) == int32(0) {
		goto L33
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_appendStringInfoString(m, v18+int32(320), int32(_a_F_pg_get_triggerdef_worker_9))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	v121 = int32(_a_F_pg_get_triggerdef_worker_10)
	goto L35
L40:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+80)))
	if v115&int32(8) == int32(0) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v121 = int32(_a_F_pg_get_triggerdef_worker_11)
	goto L35
L42:
	;
	goto L34
L43:
	;
	v135 = int32(_a_F_pg_get_triggerdef_worker_12)
	goto L32
L44:
	;
	v135 = int32(_a_F_pg_get_triggerdef_worker_13)
	goto L32
L45:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v55)+116))
	if v138 <= int32(0) {
		goto L31
	} else {
		goto L46
	}
L46:
	;
	v142 = v18 + int32(320)
	F_appendStringInfoString(m, v142, int32(_a_F_pg_get_triggerdef_worker_14))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v55)+116))
	if v146 <= int32(0) {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v150 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+124)))
	v152 = F_get_attname(m, v149, v150, int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v154 = F_quote_identifier(m, v152)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_appendStringInfoString(m, v142, v154)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v55)+116))
	if v158 < int32(2) {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	v167 = int32(1)
	goto L53
L53:
	;
	v180 = v18 + int32(320)
	F_appendStringInfoString(m, v180, int32(_a_F_pg_get_triggerdef_worker_15))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L55
	}
L54:
	;
	goto L31
L55:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v188 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55+int32(124)+v167<<(uint(int32(1))%32)))))
	v190 = F_get_attname(m, v184, v188, int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v192 = F_quote_identifier(m, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_appendStringInfoString(m, v180, v192)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v197 = v167 + int32(1)
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v55)+116))
	if v197 < v198 {
		v167 = v197
		goto L53
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	v241 = int32(_a_F_pg_get_triggerdef_worker_16)
	goto L29
L61:
	;
	v241 = int32(_a_F_pg_get_triggerdef_worker_17)
	goto L29
L62:
	;
	goto L28
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v265
	v268 = v18 + int32(320)
	F_appendStringInfo(m, v268, int32(_a_F_pg_get_triggerdef_worker_18), v18+int32(80))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L69
	}
L64:
	;
	v261 = F_generate_relation_name(m, v259, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v263 = F_generate_qualified_relation_name(m, v259)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L68
	}
L67:
	;
	v265 = v261
	goto L63
L68:
	;
	v265 = v263
	goto L63
L69:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v55)+92))
	if v274 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v55)+84))
	if v275 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v308 = v18 + int32(255)
	v309 = F_fastgetattr_3(m, v40, int32(18), v306, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L87
	}
L73:
	;
	v277 = F_generate_relation_name(m, v275, int32(0))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+96)))
	if v285 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v277
	F_appendStringInfo(m, v268, int32(_a_F_pg_get_triggerdef_worker_19), v18-int32(-64))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	F_appendStringInfoString(m, v18+int32(320), int32(_a_F_pg_get_triggerdef_worker_20))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v294 = v18 + int32(320)
	F_appendStringInfoString(m, v294, int32(_a_F_pg_get_triggerdef_worker_21))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L82
	}
L81:
	;
	goto L80
L82:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+97)))
	if v300 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v301 = int32(_a_F_pg_get_triggerdef_worker_22)
	goto L85
L84:
	;
	v301 = int32(_a_F_pg_get_triggerdef_worker_23)
	goto L85
L85:
	;
	F_appendStringInfoString(m, v294, v301)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L72
L87:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+255)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v314 = F_fastgetattr_3(m, v40, int32(19), v313, v308)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	v316 = int32(-1)
	v318 = base.I32_wrap_i64(v309)
	v319 = int32(0)
	v321 = (v311 ^ v316) & base.B2i32(v318 != v319)
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+255)))
	v325 = base.I32_wrap_i64(v314)
	v328 = (v322 ^ v316) & base.B2i32(v325 != v319)
	if v321|v328 == v319 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v359 = v18 + int32(320)
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+80)))
	if v362&int32(1) != 0 {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	v333 = v18 + int32(320)
	F_appendStringInfoString(m, v333, int32(_a_F_pg_get_triggerdef_worker_24))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	if v321 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v337 = F_quote_identifier(m, v318)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	if v328 == int32(0) {
		goto L89
	} else {
		goto L97
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v337
	F_appendStringInfo(m, v333, int32(_a_F_pg_get_triggerdef_worker_25), v18+int32(48))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	v347 = F_quote_identifier(m, v325)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v347
	F_appendStringInfo(m, v18+int32(320), int32(_a_F_pg_get_triggerdef_worker_26), v18+int32(32))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L89
L100:
	;
	v365 = int32(_a_F_pg_get_triggerdef_worker_27)
	goto L102
L101:
	;
	v365 = int32(_a_F_pg_get_triggerdef_worker_28)
	goto L102
L102:
	;
	F_appendStringInfoString(m, v359, v365)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v372 = F_fastgetattr_3(m, v40, int32(17), v369, v18+int32(255))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+255)))
	if v374 == int32(0) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	F_appendStringInfoString(m, v359, int32(_a_F_pg_get_triggerdef_worker_29))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v55)+76))
	v500 = int32(0)
	v506 = F_generate_function_name(m, v499, v500, v500, v500, v500, v500, v500)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L125
	}
L108:
	;
	v381 = F_text_to_cstring(m, base.I32_wrap_i64(v372))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	v383 = F_stringToNode(m, v381)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v386 = F_get_rel_relkind(m, v385)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v389 = F_palloc0(m, int32(136))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v391 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v389)+12)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v389))) = int32(101)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v389)+24)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v389)+21)) = uint8(v386)
	*(*int32)(unsafe.Add(mBase, uint32(v389)+16)) = v395
	v402 = F_makeAlias(m, int32(_a_F_pg_get_triggerdef_worker_30), v391)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v389)+8)) = v402
	*(*int32)(unsafe.Add(mBase, uint32(v389)+4)) = v402
	v406 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v389)+124)) = uint16(v406)
	v408 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v389)+20)) = uint8(v408)
	v411 = F_palloc0(m, int32(136))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	v413 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v411)+12)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = int32(101)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v411)+24)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v411)+21)) = uint8(v386)
	*(*int32)(unsafe.Add(mBase, uint32(v411)+16)) = v417
	v424 = F_makeAlias(m, int32(_a_F_pg_get_triggerdef_worker_31), v413)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v411)+8)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v411)+4)) = v424
	v428 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v411)+124)) = uint16(v428)
	v430 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v411)+20)) = uint8(v430)
	base.MemoryFill(m, v18+int32(136), v430, int32(76))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+124)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v411
	v445 = F_list_make2_impl(m, v18+int32(28), v18+int32(24))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v447 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+152)) = v447
	*(*int64)(unsafe.Add(mBase, uint32(v18)+144)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+132)) = v445
	v453 = v18 + int32(132)
	F_set_rtable_names(m, v453, v447, v447)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_set_simple_column_names(m, v453)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+120)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v18)+212)) = v359
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v453
	v466 = F_list_make1_impl(m, int32(1), v18+int32(20))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	v468 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+244)) = uint8(v468)
	v470 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+228)) = v470
	*(*int64)(unsafe.Add(mBase, uint32(v18)+220)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+216)) = v466
	*(*int32)(unsafe.Add(mBase, uint32(v18)+248)) = v470
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+247)) = uint8(v470)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+245)) = uint16(v468)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+236)) = int64(34359738368)
	if l1 != 0 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v485 = int32(7)
	goto L122
L121:
	;
	v485 = int32(2)
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+232)) = v485
	F_get_rule_expr(m, v383, v18+int32(212), int32(0))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_appendStringInfoString(m, v359, int32(_a_F_pg_get_triggerdef_worker_32))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	goto L107
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v506
	F_appendStringInfo(m, v18+int32(320), int32(_a_F_pg_get_triggerdef_worker_33), v18+int32(16))
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v516 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+98)))
	if v516 <= int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	F_appendStringInfoChar(m, v18+int32(320), int32(41))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L159
	}
L128:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(v22)+52))
	v523 = F_fastgetattr_3(m, v40, int32(16), v520, v18+int32(255))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v525 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+255)))
	if v525 != int32(1) {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	v529 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v523))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L156
	}
L133:
	;
	v531 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+98)))
	if v531 <= int32(0) {
		goto L127
	} else {
		goto L134
	}
L134:
	;
	v534 = int32(1)
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529))))
	if v536&v534 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v539 = v534
	goto L137
L136:
	;
	v539 = int32(4)
	goto L137
L137:
	;
	v542 = int32(0)
	v543 = v529 + v539
	goto L138
L138:
	;
	if v542 != 0 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	F_appendStringInfoString(m, v18+int32(320), int32(_a_F_pg_get_triggerdef_worker_15))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	F_appendStringInfoChar(m, v18+int32(320), int32(39))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	v570 = v543
	goto L145
L145:
	;
	v582 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v570))))
	if v582 != int32(39) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	F_appendStringInfoChar(m, v18+int32(320), base.I32_extend8_s(v582))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L155
	}
L148:
	;
	if v582 != 0 {
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	F_appendStringInfoChar(m, v18+int32(320), int32(39))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L154
	}
L151:
	;
	F_appendStringInfoChar(m, v18+int32(320), int32(39))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v590 = F_strlen(m, v543)
	mBase = m.M
	v592 = int32(1)
	v595 = v542 + v592
	v596 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55)+98)))
	if v595 < v596 {
		v542 = v595
		v543 = v590 + v543 + v592
		goto L138
	} else {
		goto L153
	}
L153:
	;
	goto L127
L154:
	;
	goto L147
L155:
	;
	v570 = v570 + int32(1)
	goto L145
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_triggerdef_worker_34), v18)
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	F_errfinish(m, int32(_a_F_pg_get_triggerdef_worker_7), int32(1143), int32(_a_F_pg_get_triggerdef_worker_8))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L159:
	;
	F_systable_endscan(m, v38)
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_relation_close(m, v22, int32(1))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v18)+320))
	v652 = v648
	goto L4
}
func F_pg_get_viewdef_name_ext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = F_textToQualifiedNameList(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = F_makeRangeVarFromNameList(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				v14 = int32(0)
				v18 = F_RangeVarGetRelidExtended(m, v12, v14, v14, v14, v14)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					if v9 == int64(0) {
						v24 = int32(2)
					} else {
						v24 = int32(7)
					}
					v26 = F_pg_get_viewdef_worker(m, v18, v24, int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int64(0)
					} else {
						if v26 == int32(0) {
							v30 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
							return int64(0)
						} else {
							v34 = F_cstring_to_text(m, v26)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								F_pfree(m, v26)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v34)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_gmtime(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v85 int64
	_ = v85
	var v90 int32
	_ = v90
	var __phi90 int32
	_ = __phi90
	var v95 int64
	_ = v95
	var __phi95 int64
	_ = __phi95
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int64
	_ = v116
	var v124 int32
	_ = v124
	var v126 int64
	_ = v126
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v209 int64
	_ = v209
	var v212 int64
	_ = v212
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v237 int64
	_ = v237
	var v240 int64
	_ = v240
	var v244 int64
	_ = v244
	var v245 int64
	_ = v245
	var v255 int32
	_ = v255
	var v256 int64
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int64
	_ = v367
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v513 int32
	_ = v513
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[0]))
	if v4 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[0]))
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v7 = F_emscripten_builtin_malloc(m, int32(_a_F_pg_gmtime_0))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[0])) = v7
	if v7 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v15 = F_tzload(m, int32(_a_F_pg_gmtime_1), int32(0), v7)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	if v15 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v23 = F_tzparse(m, int32(_a_F_pg_gmtime_1), v7, int32(1))
	mBase = m.M
	goto L1
L9:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[1])) = v27 + int32(_a_F_pg_gmtime_2)
	return v513
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v38 = v37
	goto L12
L11:
	;
	v38 = v25
	goto L12
L12:
	;
	v44 = v38
	goto L15
L13:
	;
	v81 = int64(86400)
	v82 = base.I64_div_s(v77, v81)
	v85 = v77 - v82*v81
	__phi90 = int32(1970)
	__phi95 = v82
	v90 = __phi90
	v95 = __phi95
	goto L22
L14:
	;
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v77 = v74
	v78 = int64(0)
	v80 = int32(0)
	goto L13
L15:
	;
	v54 = v44 - int32(1)
	if v54 < int32(0) {
		goto L14
	} else {
		goto L17
	}
L16:
	;
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	if v57 != v61 {
		v77 = v57
		v78 = v63
		v80 = int32(0)
		goto L13
	} else {
		goto L19
	}
L17:
	;
	v57 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v60 = v27 + int32(_a_F_pg_gmtime_3) + v54<<(uint(int32(4))%32)
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	if v57 < v61 {
		v44 = v54
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if v54 == int32(0) {
		v77 = v57
		v78 = v63
		v80 = base.B2i32(int64(0) < v63)
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v60-int32(8))))
	v77 = v57
	v78 = v63
	v80 = base.B2i32(v72 < v63)
	goto L13
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[2])) = int32(61)
	v513 = int32(0)
	goto L9
L22:
	;
	v100 = base.B2i32(v95 < int64(0))
	if v100 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v203 = base.I32_wrap_i64(v95)
	v204 = base.I64_extend_i32_s(v25)
	v206 = v204 - v78 + v85
	if v206 < int64(0) {
		goto L53
	} else {
		goto L54
	}
L24:
	;
	goto L23
L25:
	;
	if v90&int32(3) != 0 {
		v113 = int32(0)
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	if base.Ui64(int64(1571958030700)) < base.Ui64(v95+int64(785979015533)) {
		goto L21
	} else {
		goto L32
	}
L28:
	;
	v116 = int64(*(*int32)(unsafe.Add(mBase, uint32(v113<<(uint(int32(2))%32))+uint32(_c_F_pg_gmtime[3]))))
	if v95 < v116 {
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v108 = base.I32_rem_s(v90, int32(100))
	if v108 != 0 {
		v113 = int32(1)
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v110 = base.I32_rem_s(v90, int32(400))
	v113 = base.B2i32(v110 == int32(0))
	goto L28
L31:
	;
	goto L27
L32:
	;
	if v95 < int64(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v124 = int32(-1)
	goto L35
L34:
	;
	v124 = int32(1)
	goto L35
L35:
	;
	v126 = base.I64_div_s(v95, int64(366))
	if base.Ui64(v95+int64(365)) < base.Ui64(int64(731)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v132 = v124
	goto L38
L37:
	;
	v132 = base.I32_wrap_i64(v126)
	goto L38
L38:
	;
	if int32(0) <= v90 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v141 = v132 + v90
	v143 = v141 - int32(1)
	if v143 < int32(0) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	if v132 <= v90^int32(2147483647) {
		goto L39
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v132 < int32(-2147483648)-v90 {
		goto L21
	} else {
		goto L44
	}
L43:
	;
	goto L21
L44:
	;
	goto L39
L45:
	;
	v175 = v90 - int32(1)
	if v175 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v147 = int32(0) - v141
	v151 = base.I32_div_u_s(v147, int32(100))
	v154 = base.I32_div_u_s(v147, int32(400))
	v167 = int32(base.Ui32(v147)>>(uint(int32(2))%32)) - v151 + v154 ^ int32(-1)
	goto L45
L47:
	;
	goto L48
L48:
	;
	v161 = base.I32_div_u_s(v143, int32(100))
	v164 = base.I32_div_u_s(v143, int32(400))
	v167 = int32(base.Ui32(v143)>>(uint(int32(2))%32)) - v161 + v164
	goto L45
L49:
	;
	__phi90 = v141
	__phi95 = (base.I64_extend_i32_s(v141)-base.I64_extend_i32_s(v90))*int64(-365) + v95 - base.I64_extend_i32_s(v167-v199)
	v90 = __phi90
	v95 = __phi95
	goto L22
L50:
	;
	v179 = int32(0) - v90
	v183 = base.I32_div_u_s(v179, int32(100))
	v186 = base.I32_div_u_s(v179, int32(400))
	v199 = int32(base.Ui32(v179)>>(uint(int32(2))%32)) - v183 + v186 ^ int32(-1)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v193 = base.I32_div_u_s(v175, int32(100))
	v196 = base.I32_div_u_s(v175, int32(400))
	v199 = int32(base.Ui32(v175)>>(uint(int32(2))%32)) - v193 + v196
	goto L49
L53:
	;
	v209 = int64(-86400)
	if base.Ui64(v206) <= base.Ui64(v209) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	v233 = v203
	v234 = v206
	goto L55
L55:
	;
	if base.Ui64(int64(86400)) <= base.Ui64(v234) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v212 = v209
	goto L58
L57:
	;
	v212 = v206
	goto L58
L58:
	;
	v214 = v78 + v212 - v85
	v216 = base.I64_extend_i32_u(base.B2i32(v214 != v204))
	v219 = int64(86400)
	v220 = base.I64_div_u_s(v214-(v216+v204), v219)
	v221 = v220 + v216
	v233 = base.I32_wrap_i64(v221) ^ int32(-1) + v203
	v234 = v85 + v221*v219 + v204 - v78 + v219
	goto L55
L59:
	;
	v237 = int64(172799)
	if v237 <= v234 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v255 = v233
	v256 = v234
	goto L61
L61:
	;
	if v255 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	v240 = v237
	goto L64
L63:
	;
	v240 = v234
	goto L64
L64:
	;
	v244 = int64(86400)
	v245 = base.I64_div_u_s(v234-v240+int64(86399), v244)
	v255 = v233 + base.I32_wrap_i64(v245) + int32(1)
	v256 = v234 + v245*int64(-86400) - v244
	goto L61
L65:
	;
	v260 = v255
	v263 = v90
	goto L68
L66:
	;
	v293 = v255
	v296 = v90
	goto L67
L67:
	;
	v305 = v293
	v308 = v296
	goto L75
L68:
	;
	if v263 == int32(-2147483648) {
		goto L21
	} else {
		goto L70
	}
L69:
	;
	v293 = v290
	v296 = v276
	goto L67
L70:
	;
	v276 = v263 - int32(1)
	if v276&int32(3) != 0 {
		v286 = int32(0)
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v286<<(uint(int32(2))%32))+uint32(_c_F_pg_gmtime[3])))
	v290 = v289 + v260
	if v290 < int32(0) {
		v260 = v290
		v263 = v276
		goto L68
	} else {
		goto L74
	}
L72:
	;
	v281 = base.I32_rem_s(v276, int32(100))
	if v281 != 0 {
		v286 = int32(1)
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v283 = base.I32_rem_s(v276, int32(400))
	v286 = base.B2i32(v283 == int32(0))
	goto L71
L74:
	;
	goto L69
L75:
	;
	v318 = v308 & int32(3)
	if v318 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[4])) = v308
	if v308 < int32(-2147481748) {
		goto L21
	} else {
		goto L89
	}
L77:
	;
	goto L76
L78:
	;
	if v308 == int32(2147483647) {
		goto L21
	} else {
		goto L88
	}
L79:
	;
	v322 = base.I32_rem_s(v308, int32(100))
	if v322 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L80:
	;
	goto L81
L81:
	;
	if v305 < int32(365) {
		goto L77
	} else {
		goto L87
	}
L82:
	;
	v326 = base.I32_rem_s(v308, int32(400))
	v328 = base.B2i32(v326 == int32(0))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v328<<(uint(int32(2))%32))+uint32(_c_F_pg_gmtime[3])))
	if v305 < v331 {
		goto L77
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	if v305 < int32(366) {
		goto L77
	} else {
		goto L86
	}
L85:
	;
	v340 = v328
	goto L78
L86:
	;
	v340 = int32(1)
	goto L78
L87:
	;
	v340 = int32(0)
	goto L78
L88:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v340<<(uint(int32(2))%32))+uint32(_c_F_pg_gmtime[3])))
	v305 = v305 - v347
	v308 = v308 + int32(1)
	goto L75
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[5])) = v305
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[4])) = v308 - int32(1900)
	v363 = base.I32_rem_s(v308-int32(1970), int32(7))
	v364 = int32(0)
	v367 = base.I64_div_u_s(v256, int64(3600))
	*(*uint32)(unsafe.Add(mBase, _c_F_pg_gmtime[6])) = uint32(v367)
	if v308 <= v364 {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v400 = int32(7)
	v401 = base.I32_rem_s(v395+(v305+v363)-int32(473), v400)
	if v401 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L91:
	;
	v373 = int32(0) - v308
	v377 = base.I32_div_u_s(v373, int32(100))
	v380 = base.I32_div_u_s(v373, int32(400))
	v395 = int32(base.Ui32(v373)>>(uint(int32(2))%32)) - v377 + v380 ^ int32(-1)
	goto L90
L92:
	;
	goto L93
L93:
	;
	v385 = v308 - int32(1)
	v389 = base.I32_div_u_s(v385, int32(100))
	v392 = base.I32_div_u_s(v385, int32(400))
	v395 = int32(base.Ui32(v385)>>(uint(int32(2))%32)) - v389 + v392
	goto L90
L94:
	;
	v406 = v401 + v400
	goto L96
L95:
	;
	v406 = v401
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[7])) = v406
	v412 = base.I32_wrap_i64(v256 - v367*int64(3600))
	v413 = int32(_a_F_pg_gmtime_4)
	v415 = int32(60)
	v416 = base.I32_div_u_s(v412&v413, v415)
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[8])) = v416
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[9])) = v80 + (v412-v416*v415)&v413
	if v318 != 0 {
		v434 = int32(0)
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v436 = v434 * int32(48)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v436)+uint32(_c_F_pg_gmtime[10])))
	if v437 <= v305 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	v429 = base.I32_rem_s(v308, int32(100))
	if v429 != 0 {
		v434 = int32(1)
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v431 = base.I32_rem_s(v308, int32(400))
	v434 = base.B2i32(v431 == int32(0))
	goto L97
L100:
	;
	v441 = v305
	v443 = v364
	v444 = v437
	goto L103
L101:
	;
	v461 = v305
	v463 = v364
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[11])) = v25
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[12])) = v463
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[13])) = v461 + int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_pg_gmtime[14])) = int32(0)
	v513 = int32(_a_F_pg_gmtime_5)
	goto L9
L103:
	;
	v453 = v441 - v444
	v455 = v443 + int32(1)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v436+int32(_a_F_pg_gmtime_6)+v455<<(uint(int32(2))%32))))
	if v459 <= v453 {
		v441 = v453
		v443 = v455
		v444 = v459
		goto L103
	} else {
		goto L105
	}
L104:
	;
	v461 = v453
	v463 = v455
	goto L102
L105:
	;
	goto L104
}
func F_pg_indexam_has_property(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = F_text_to_cstring(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v11 = int32(0)
			v13 = F_indexam_property(m, l0, v9, v3, v11, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int64(0)
			} else {
				return v13
			}
		}
	}
}
func F_pg_indexam_progress_phasename(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_GetIndexAmRoutineByAmId(m, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+80))
			if v11 != 0 {
				v18 = m.T0[v11].(func(*base.Module, int64) int32)(m, base.I64_extend32_s(v4))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					if v18 == int32(0) {
						v22 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
						return int64(0)
					} else {
						v26 = F_cstring_to_text(m, v18)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v26)
						}
					}
				}
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int64(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int64(0)
		}
	}
}
func F_pg_initialize_timing(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int64
	_ = v6
	var v12 int32
	_ = v12
	v2 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_initialize_timing[0])))
	if v2 == int32(0) {
		v6 = int64(0)
		*(*int64)(unsafe.Add(mBase, _c_F_pg_initialize_timing[1])) = v6
		*(*int64)(unsafe.Add(mBase, _c_F_pg_initialize_timing[2])) = v6
		v12 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_initialize_timing[0])) = uint8(v12)
	} else {
	}
	return
}
func F_pg_iswalpha(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v4 == int32(0) {
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v21 = int32(0)
			return v21
		} else {
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_pg_iswalpha[0]))))
			return int32(base.Ui32(v10&int32(2)) >> (uint(int32(1)) % 32))
		}
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
		v17 = m.T0[v16].(func(*base.Module, int32, int32) int32)(m, l0, l1)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			return v21
		}
	}
}
func F_pg_itoa(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	if int32(0) <= l0 {
		v12 = l0
		v13 = int32(0)
	} else {
		v7 = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v7)
		v12 = int32(0) - l0
		v13 = int32(1)
	}
	v14 = l1 + v13
	v15 = int32(0)
	if v12 == v15 {
		v24 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v24)
		v134 = int32(1)
	} else {
		v30 = int32(1233)
		v35 = int32(base.Ui32((base.I32_clz(v12)^int32(31))*v30+v30) >> (uint(int32(12)) % 32))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v35<<(uint(int32(2))%32))+uint32(_c_F_pg_itoa[0])))
		v40 = v35 + base.B2i32(base.Ui32(v38) <= base.Ui32(v12))
		if base.Ui32(int32(_a_F_pg_itoa_0)) <= base.Ui32(v12) {
			v44 = v12
			v46 = v15
			for {
				v53 = v14 + v40 - v46
				v54 = int32(4)
				v57 = base.I32_div_u_s(v44, int32(_a_F_pg_itoa_0))
				v60 = v44 + v57*int32(-10000)
				v61 = int32(100)
				v62 = base.I32_div_u_s(v60, v61)
				v63 = int32(1)
				v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62<<(uint(v63)%32))+uint32(_c_F_pg_itoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v53-v54))) = uint16(v65)
				v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v60-v62*v61)<<(uint(v63)%32))+uint32(_c_F_pg_itoa[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v53-int32(2)))) = uint16(v74)
				v77 = v46 + v54
				if base.Ui32(int32(99999999)) < base.Ui32(v44) {
					v44 = v57
					v46 = v77
					continue
				} else {
					break
				}
				break
			}
			v80 = v57
			v82 = v77
		} else {
			v80 = v12
			v82 = v15
		}
		if base.Ui32(int32(100)) <= base.Ui32(v80) {
			v93 = int32(2)
			v95 = int32(_a_F_pg_itoa_1)
			v97 = int32(100)
			v98 = base.I32_div_u_s(v80&v95, v97)
			v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v80-v98*v97)&v95<<(uint(int32(1))%32))+uint32(_c_F_pg_itoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v40-v82-v93))) = uint16(v106)
			v110 = v98
			v111 = v82 | v93
		} else {
			v110 = v80
			v111 = v82
		}
		if base.Ui32(int32(10)) <= base.Ui32(v110) {
			v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110<<(uint(int32(1))%32))+uint32(_c_F_pg_itoa[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(v14+v40-v111-int32(2)))) = uint16(v120)
			v134 = v40
		} else {
			v123 = v110 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v123)
			v134 = v40
		}
	}
	v135 = v134 + v13
	v137 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1+v135))) = uint8(v137)
	return v135
}
func F_pg_last_wal_receive_lsn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v3 = int32(0)
	v5 = F_GetWalRcvFlushRecPtr(m, v3, v3)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int64(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
		} else {
		}
		return v5
	}
}
func F_pg_log_standby_snapshot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_log_standby_snapshot[0])))
	if v9 == int32(1) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_pg_log_standby_snapshot[1]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+308))
		v17 = base.B2i32(v15 != int32(2))
		*(*uint8)(unsafe.Add(mBase, _c_F_pg_log_standby_snapshot[0])) = uint8(v17)
		v19 = v17
	} else {
		v19 = int32(0)
	}
	if v19 == int32(0) {
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_pg_log_standby_snapshot[2]))
		if v23 <= int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_log_standby_snapshot_0), int32(0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_log_standby_snapshot_1), int32(246), int32(_a_F_pg_log_standby_snapshot_2))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v26 = F_LogStandbySnapshot(m)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				m.G0 = v5 + int32(16)
				return v26
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_log_standby_snapshot_3), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_log_standby_snapshot_4)
					F_errhint(m, int32(_a_F_pg_log_standby_snapshot_5), v5)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_log_standby_snapshot_1), int32(241), int32(_a_F_pg_log_standby_snapshot_2))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
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
func F_pg_lsn_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v4 = F_pg_lsn_in_safe(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_pg_lsn_mii(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14348(m, l0, int32(19), int32(_a_F_pg_lsn_mii_0), int32(291), int32(_a_F_pg_lsn_mii_1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_pg_lsn_pli(m *base.Module, l0 int32) int64 {
	var v6 int64
	_ = v6
	var v9 int32
	_ = v9
	v6 = Fn14348(m, l0, int32(1406), int32(_a_F_pg_lsn_pli_0), int32(257), int32(_a_F_pg_lsn_pli_1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return v6
	}
}
func F_pg_mblen_cstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mblen_cstr[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_c_F_pg_mblen_cstr[1])))
	v13 = m.T0[v12].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_report_invalid_encoding_db(m, l0, v13, v20)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L11
	}
L2:
	;
	return int32(0)
L3:
	;
	if int32(1) < v13 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v20 = int32(1)
	goto L7
L5:
	;
	goto L6
L6:
	;
	return v13
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v20))))
	if v23 == int32(0) {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	v27 = v20 + int32(1)
	if v27 != v13 {
		v20 = v27
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_mblen_with_len(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pg_mblen_with_len[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v6*int32(28))+uint32(_c_F_pg_mblen_with_len[1])))
	v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if l1 < v12 {
			F_report_invalid_encoding_db(m, l0, v12, l1)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return v12
		}
	}
}
func F_pg_ndistinct_recv(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14235(m, l0, int32(_a_F_pg_ndistinct_recv_0), int32(836), int32(_a_F_pg_ndistinct_recv_1), int32(_a_F_pg_ndistinct_recv_2), int32(_a_F_pg_ndistinct_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_pg_newlocale_from_collation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v148 int64
	_ = v148
	var v153 int32
	_ = v153
	var v154 int64
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v173 int64
	_ = v173
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int64
	_ = v266
	var v268 int64
	_ = v268
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int64
	_ = v374
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int64
	_ = v422
	var v424 int64
	_ = v424
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int64
	_ = v447
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v502 int32
	_ = v502
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v539 int32
	_ = v539
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
	var v554 int64
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v567 int64
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
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
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v651 int32
	_ = v651
	var v656 int32
	_ = v656
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v711 int32
	_ = v711
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	if l0 != int32(100) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L9
	} else {
		goto L173
	}
L2:
	;
	m.G0 = v17 + int32(112)
	return v730
L3:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[0]))
	if l0 == v58 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	if l0 == int32(950) {
		v730 = int32(_a_F_pg_newlocale_from_collation_0)
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[1]))
	if v43 != 0 {
		v730 = v43
		goto L2
	} else {
		goto L13
	}
L7:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = int32(0)
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_1), v17+int32(96))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_2), int32(1211), int32(_a_F_pg_newlocale_from_collation_3))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L9
	} else {
		goto L14
	}
L14:
	;
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_4), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L9
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_2), int32(1198), int32(_a_F_pg_newlocale_from_collation_3))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[2]))
	v730 = v61
	goto L2
L18:
	;
	goto L19
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[3]))
	if v63 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v96 = int32(16)
	v100 = (int32(base.Ui32(l0)>>(uint(v96)%32)) ^ l0) * int32(-2048144789)
	v105 = (int32(base.Ui32(v100)>>(uint(int32(13))%32)) ^ v100) * int32(-1028477387)
	v108 = int32(base.Ui32(v105)>>(uint(v96)%32)) ^ v105
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v112 = base.B2i32(base.Ui32(v109) < base.Ui32(v95))
	goto L27
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v94 = v63
	v95 = v64
	goto L20
L22:
	;
	goto L23
L23:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[4]))
	v72 = F_AllocSetContextCreateInternal(m, v67, int32(_a_F_pg_newlocale_from_collation_5), int32(0), int32(_a_F_pg_newlocale_from_collation_6), int32(_a_F_pg_newlocale_from_collation_7))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[5])) = v72
	v76 = F_MemoryContextAllocZero(m, v72, int32(32))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = v72
	v83 = F_MemoryContextAllocExtended(m, v72, int32(512), int32(5))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v76)+12)) = int64(120259084319)
	*(*int64)(unsafe.Add(mBase, uint32(v76))) = int64(32)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = v83
	*(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[3])) = v76
	v94 = v76
	v95 = int32(28)
	goto L20
L27:
	;
	if v112 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v726 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v726
	v112 = v726
	goto L27
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L9
	} else {
		goto L170
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L9
	} else {
		goto L167
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[2])) = v664
	*(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[0])) = l0
	v730 = v664
	goto L2
L33:
	;
	v510 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[5]))
	v513 = F_SearchSysCache1(m, int32(16), base.I64_extend_i32_u(l0))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L9
	} else {
		goto L119
	}
L34:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v486 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v485 + v486
	*(*uint8)(unsafe.Add(mBase, uint32(v478)+12)) = uint8(v486)
	*(*int32)(unsafe.Add(mBase, uint32(v478)+8)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v478))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v478)+4)) = int32(0)
	v502 = v478
	goto L33
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L9
	} else {
		goto L116
	}
L36:
	;
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
	if v127 == int64(4294967296) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v313 = v312 & v108
	v316 = v311 + v313<<(uint(int32(4))%32)
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+12)))
	if v317 == int32(0) {
		v478 = v316
		goto L34
	} else {
		goto L80
	}
L39:
	;
	v130 = int32(0)
	v132 = int64(2)
	v134 = v127 << (uint(int64(1)) % 64)
	if base.Ui64(v134) <= base.Ui64(v132) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v112 = int32(1)
	goto L27
L41:
	;
	v137 = v132
	goto L43
L42:
	;
	v137 = v134
	goto L43
L43:
	;
	v138 = int64(1)
	if v137&(v137-v138) == int64(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v148 = v137
	goto L46
L45:
	;
	v148 = v138 << (uint(int64(64)-base.I64_clz(v137)) % 64)
	goto L46
L46:
	;
	if base.Ui64(v148<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v94)+20))
	v154 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v94)+24))
	v160 = F_MemoryContextAllocExtended(m, v155, base.I32_wrap_i64(v148)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	goto L1
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+20)) = v160
	v163 = int64(1)
	if v148&(v148-v163) == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v173 = v148
	goto L53
L52:
	;
	v173 = v163 << (uint(int64(64)-base.I64_clz(v148)) % 64)
	goto L53
L53:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v173<<(uint(int64(4))%64)) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v94))) = v173
	v181 = base.I32_wrap_i64(v173) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v181
	if v173 == int64(4294967296) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v190 = int32(-85899346)
	goto L57
L56:
	;
	v190 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v173), float64(0.9)))
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v94)+16)) = v190
	if v154 != int64(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v195 = v130
	goto L62
L59:
	;
	goto L60
L60:
	;
	F_pfree(m, v153)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L9
	} else {
		goto L79
	}
L61:
	;
	v224 = v222
	v230 = v130
	goto L67
L62:
	;
	v210 = v153 + v195<<(uint(int32(4))%32)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210)+12)))
	if v211 != int32(1) {
		v222 = v195
		goto L61
	} else {
		goto L64
	}
L63:
	;
	v222 = int32(0)
	goto L61
L64:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v210)+8))
	if v214&v181 == v195 {
		v222 = v195
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v218 = v195 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v218)) < base.Ui64(v154) {
		v195 = v218
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v239 = v153 + v224<<(uint(int32(4))%32)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+12)))
	if v240 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L60
L69:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v239)+8))
	v249 = v244
	goto L72
L70:
	;
	goto L71
L71:
	;
	v285 = v224 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v285)) < base.Ui64(v154) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v259 = v249 & v243
	v264 = v160 + v259<<(uint(int32(4))%32)
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+12)))
	if v265 != 0 {
		v249 = v259 + int32(1)
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v239)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v264)+8)) = v266
	v268 = *(*int64)(unsafe.Add(mBase, uint32(v239)))
	*(*int64)(unsafe.Add(mBase, uint32(v264))) = v268
	goto L71
L74:
	;
	goto L73
L75:
	;
	v289 = v285
	goto L77
L76:
	;
	v289 = int32(0)
	goto L77
L77:
	;
	v291 = v230 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v291)) < base.Ui64(v154) {
		v224 = v289
		v230 = v291
		goto L67
	} else {
		goto L78
	}
L78:
	;
	goto L68
L79:
	;
	goto L40
L80:
	;
	v322 = v313
	v325 = int32(0)
	v328 = v316
	goto L81
L81:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v328)+8))
	if v335 == v108 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v457 != 0 {
		v664 = v457
		goto L32
	} else {
		goto L115
	}
L83:
	;
	goto L82
L84:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	if v337 == l0 {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v340 = v322 + int32(1)
	v341 = v312 & v335
	if base.Ui32(v322) < base.Ui32(v341) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L86
L88:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v345 = v322 + v343
	goto L90
L89:
	;
	v345 = v322
	goto L90
L90:
	;
	if base.Ui32(v345-v341) < base.Ui32(v325) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v348 = v340 & v312
	v351 = v311 + v348<<(uint(int32(4))%32)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v351)+12)))
	if v352 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v442 = v325 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v442) {
		goto L110
	} else {
		goto L111
	}
L94:
	;
	v358 = v348
	v363 = int32(0)
	goto L97
L95:
	;
	v387 = v351
	v390 = v348
	goto L96
L96:
	;
	if v390 != v322 {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	v369 = v363 + int32(1)
	if int32(151) <= v369 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v387 = v384
	v390 = v381
	goto L96
L99:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v374 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v372), base.F64_convert_i64_u(v374)), float64(0.1)) != 0 {
		goto L29
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v381 = (v358 + int32(1)) & v312
	v384 = v311 + v381<<(uint(int32(4))%32)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+12)))
	if v385 != 0 {
		v358 = v381
		v363 = v369
		goto L97
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	goto L98
L104:
	;
	v402 = v387
	v405 = v390
	goto L107
L105:
	;
	goto L106
L106:
	;
	v478 = v328
	goto L34
L107:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v418 = v415 & (v405 - int32(1))
	v421 = v311 + v418<<(uint(int32(4))%32)
	v422 = *(*int64)(unsafe.Add(mBase, uint32(v421)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v402)+8)) = v422
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v421)))
	*(*int64)(unsafe.Add(mBase, uint32(v402))) = v424
	if v418 != v322 {
		v402 = v421
		v405 = v418
		goto L107
	} else {
		goto L109
	}
L108:
	;
	goto L106
L109:
	;
	goto L108
L110:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v94)+8))
	v447 = *(*int64)(unsafe.Add(mBase, uint32(v94)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v445), base.F64_convert_i64_u(v447)), float64(0.1)) != 0 {
		goto L29
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v452 = v340 & v312
	v455 = v311 + v452<<(uint(int32(4))%32)
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+12)))
	if v456 != 0 {
		v322 = v452
		v325 = v442
		v328 = v455
		goto L81
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	v478 = v455
	goto L34
L115:
	;
	v502 = v328
	goto L33
L116:
	;
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_8), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L9
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_9), int32(635), int32(_a_F_pg_newlocale_from_collation_10))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L9
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	if v513 == int32(0) {
		goto L31
	} else {
		goto L120
	}
L120:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513)+16))
	v518 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v517)+22)))
	v519 = v517 + v518
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+76)))
	switch v520 - int32(98) {
	case 0:
		goto L122
	case 1:
		goto L124
	default:
		goto L123
	case 7:
		goto L125
	}
L121:
	;
	v548 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v547)+3)) = uint8(v548)
	v554 = F_SysCacheGetAttr(m, int32(16), v513, int32(12), v17+int32(111))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L9
	} else {
		goto L132
	}
L122:
	;
	v545 = F_create_pg_locale_builtin(m, l0, v510)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L9
	} else {
		goto L131
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L9
	} else {
		goto L128
	}
L124:
	;
	v525 = F_create_pg_locale_libc(m, l0, v510)
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L9
	} else {
		goto L127
	}
L125:
	;
	v523 = F_create_pg_locale_icu(m)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L9
	} else {
		goto L126
	}
L126:
	;
	v547 = v523
	goto L121
L127:
	;
	v547 = v525
	goto L121
L128:
	;
	v531 = int32(*(*int8)(unsafe.Add(mBase, uint32(v519)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v531
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(_a_F_pg_newlocale_from_collation_11)
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_12), v17+int32(16))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L9
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_2), int32(1070), int32(_a_F_pg_newlocale_from_collation_11))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L9
	} else {
		goto L130
	}
L130:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L131:
	;
	v547 = v545
	goto L121
L132:
	;
	v556 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+111)))
	if v556 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	F_ReleaseCatCache(m, v513)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L9
	} else {
		goto L166
	}
L134:
	;
	v558 = F_text_to_cstring(m, base.I32_wrap_i64(v554))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L9
	} else {
		goto L135
	}
L135:
	;
	v563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+76)))
	if v563 == int32(99) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v566 = int32(8)
	goto L138
L137:
	;
	v566 = int32(10)
	goto L138
L138:
	;
	v567 = F_SysCacheGetAttrNotNull(m, int32(16), v513, v566)
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L9
	} else {
		goto L139
	}
L139:
	;
	v569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519)+76)))
	v571 = F_text_to_cstring(m, base.I32_wrap_i64(v567))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L9
	} else {
		goto L140
	}
L140:
	;
	switch v569 - int32(98) {
	case 0:
		goto L143
	case 1:
		goto L142
	default:
		goto L30
	}
L141:
	;
	if v589 == int32(0) {
		goto L30
	} else {
		goto L149
	}
L142:
	;
	v578 = F_pg_strcasecmp(m, int32(_a_F_pg_newlocale_from_collation_13), v571)
	mBase = m.M
	if v578 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v575 = F_get_collation_actual_version_builtin(m, v571)
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L9
	} else {
		goto L144
	}
L144:
	;
	v589 = v575
	goto L141
L145:
	;
	v589 = int32(0)
	goto L141
L146:
	;
	goto L145
L147:
	;
	v583 = F_pg_strncasecmp(m, int32(_a_F_pg_newlocale_from_collation_14), v571, int32(2))
	mBase = m.M
	if v583 == int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v587 = F_pg_strcasecmp(m, int32(_a_F_pg_newlocale_from_collation_15), v571)
	mBase = m.M
	goto L146
L149:
	;
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589))))
	v597 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558))))
	if base.B2i32(v594 == int32(0))|base.B2i32(v594 != v597) != 0 {
		v615 = v594
		v616 = v597
		goto L151
	} else {
		goto L152
	}
L150:
	;
	if v615-v616 == int32(0) {
		goto L133
	} else {
		goto L157
	}
L151:
	;
	goto L150
L152:
	;
	v600 = v589
	v601 = v558
	goto L153
L153:
	;
	v604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601)+1)))
	v605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v600)+1)))
	if v605 == int32(0) {
		v615 = v605
		v616 = v604
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v615 = v605
	v616 = v604
	goto L151
L155:
	;
	v608 = int32(1)
	if v605 == v604 {
		v600 = v600 + v608
		v601 = v601 + v608
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v622 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L9
	} else {
		goto L158
	}
L158:
	;
	if v622 == int32(0) {
		goto L133
	} else {
		goto L159
	}
L159:
	;
	v627 = v519 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v627
	F_errmsg(m, int32(_a_F_pg_newlocale_from_collation_16), v17+int32(80))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L9
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v589
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v558
	v639 = F_errdetail(m, int32(_a_F_pg_newlocale_from_collation_17), v17-int32(-64))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L9
	} else {
		goto L161
	}
L161:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v519)+68))
	v642 = F_get_namespace_name(m, v641)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L9
	} else {
		goto L162
	}
L162:
	;
	v644 = F_quote_qualified_identifier(m, v642, v627)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L9
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v644
	F_errhint(m, int32(_a_F_pg_newlocale_from_collation_18), v17+int32(48))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L9
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_2), int32(1119), int32(_a_F_pg_newlocale_from_collation_11))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L9
	} else {
		goto L165
	}
L165:
	;
	goto L133
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v502)+4)) = v547
	v664 = v547
	goto L32
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_1), v17)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L9
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_2), int32(1059), int32(_a_F_pg_newlocale_from_collation_11))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L9
	} else {
		goto L169
	}
L169:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v519 + int32(4)
	F_errmsg(m, int32(_a_F_pg_newlocale_from_collation_19), v17+int32(32))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L9
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_2), int32(1105), int32(_a_F_pg_newlocale_from_collation_11))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L9
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L173:
	;
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_20), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L9
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_9), int32(332), int32(_a_F_pg_newlocale_from_collation_21))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L9
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_notification_queue_usage(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	F_asyncQueueAdvanceTail(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_notification_queue_usage[0]))
		v14 = F_LWLockAcquire(m, v10+int32(3456), int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_notification_queue_usage[1]))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
			if v18 != v19 {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_notification_queue_usage[2]))
				v28 = base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v18-v19), base.F64_convert_i32_s(v24)))
			} else {
				v28 = int64(0)
			}
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_notification_queue_usage[0]))
			F_LWLockRelease(m, v30+int32(3456))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				return v28
			}
		}
	}
}
func F_pg_numa_available(m *base.Module, l0 int32) int64 {
	return int64(0)
}
func F_pg_opfamily_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_OpfamilyIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_options_to_table(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_untransformRelOptions(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v13 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	m.G0 = v10 + int32(32)
	return int64(0)
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v23 <= int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v27 = int32(0)
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v27<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v40 = F_cstring_to_text(m, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	v42 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v42)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = base.I64_extend_i32_u(v40)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v46 == v42 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v57)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v56
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	F_tuplestore_putvalues(m, v60, v61, v10+int32(16), v10+int32(14))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L15
	}
L11:
	;
	v56 = int64(0)
	v57 = int32(1)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v52 = F_cstring_to_text(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v56 = base.I64_extend_i32_u(v52)
	v57 = int32(0)
	goto L10
L15:
	;
	v69 = v27 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v69 < v70 {
		v27 = v69
		goto L7
	} else {
		goto L16
	}
L16:
	;
	goto L8
}
func F_pg_partition_ancestors(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v8 == int32(0) {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v19 = int64(0)
			v22 = F_SearchSysCacheExists(m, int32(57), v11&int64(4294967295), v19, v19, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				if v22 == int32(0) {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int64(0)
					} else {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(2)
						v95 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v95)
						return int64(0)
					}
				} else {
					v26 = base.I32_wrap_i64(v11)
					v27 = F_get_rel_relkind(m, v26)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int64(0)
					} else {
						v29 = F_get_rel_relispartition(m, v26)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int64(0)
						} else {
							if base.B2i32(v29|base.B2i32(v27 == int32(112)) == int32(0))&base.B2i32(v27&int32(255) != int32(73)) != 0 {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int64(0)
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(2)
									v95 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v95)
									return int64(0)
								}
							} else {
								v41 = int32(_a_F_pg_partition_ancestors_0)
								v42 = *(*int32)(unsafe.Add(mBase, _c_F_pg_partition_ancestors[0]))
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
								*(*int32)(unsafe.Add(mBase, _c_F_pg_partition_ancestors[0])) = v44
								v46 = F_get_partition_ancestors(m, v26)
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int64(0)
								} else {
									v48 = F_lcons_oid(m, v26, v46)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v48
										*(*int32)(unsafe.Add(mBase, _c_F_pg_partition_ancestors[0])) = v42
										v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
										if v59 == int32(0) {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return int64(0)
											} else {
												v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = int32(2)
												v84 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
												return int64(0)
											}
										} else {
											v62 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
											v63 = int64(*(*int32)(unsafe.Add(mBase, uint32(v59)+4)))
											if base.Ui64(v63) <= base.Ui64(v62) {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return int64(0)
												} else {
													v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = int32(2)
													v84 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
													return int64(0)
												}
											} else {
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
												v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65+base.I32_wrap_i64(v62)<<(uint(int32(2))%32)))))
												*(*int64)(unsafe.Add(mBase, uint32(v58))) = v62 + int64(1)
												v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(1)
												return v70
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
		v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
		v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
		if v59 == int32(0) {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int64(0)
			} else {
				v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = int32(2)
				v84 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
				return int64(0)
			}
		} else {
			v62 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
			v63 = int64(*(*int32)(unsafe.Add(mBase, uint32(v59)+4)))
			if base.Ui64(v63) <= base.Ui64(v62) {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return int64(0)
				} else {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v81)+20)) = int32(2)
					v84 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v84)
					return int64(0)
				}
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
				v70 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v65+base.I32_wrap_i64(v62)<<(uint(int32(2))%32)))))
				*(*int64)(unsafe.Add(mBase, uint32(v58))) = v62 + int64(1)
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(1)
				return v70
			}
		}
	}
}
func F_pg_partition_tree(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v163 int64
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v176 int64
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v220 int64
	_ = v220
	v9 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = base.I32_wrap_i64(v15)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v18 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v220
L2:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L49
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L46
	}
L4:
	;
	v21 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	goto L18
L7:
	;
	return int64(0)
L8:
	;
	v28 = int64(0)
	v31 = F_SearchSysCacheExists(m, int32(57), v15&int64(4294967295), v28, v28, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v31 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v35 = F_get_rel_relkind(m, v16)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v37 = F_get_rel_relispartition(m, v16)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if base.B2i32(v37|base.B2i32(v35 == int32(112)) == int32(0))&base.B2i32(v35&int32(255) != int32(73)) != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v49 = int32(_a_F_pg_partition_tree_0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_pg_partition_tree[0]))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_partition_tree[0])) = v52
	v56 = F_find_all_inheritors(m, v16, int32(1), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v61 = F_get_call_result_type(m, l0, int32(0), v13+int32(16))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	if v61 != int32(1) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v65
	*(*int32)(unsafe.Add(mBase, _c_F_pg_partition_tree[0])) = v50
	goto L6
L17:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L7
	} else {
		goto L45
	}
L18:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
	if v76 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	v80 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+4)))
	if base.Ui64(v80) <= base.Ui64(v79) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v82 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v82
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v82
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = int32(0)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v92+base.I32_wrap_i64(v79)<<(uint(int32(2))%32))))
	v98 = F_get_rel_relkind(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v100 = F_get_partition_ancestors(m, v97)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = base.I64_extend_i32_u(v97)
	if v100 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v163
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v75)+28))
	v171 = F_heap_form_tuple(m, v166, v13+int32(16), v13+int32(12))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L43
	}
L24:
	;
	v128 = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v128 < v129 {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = base.I64_extend_i32_u(v105)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = base.I64_extend_i32_u(base.B2i32(v98 != int32(112)) & base.B2i32(v98 != int32(73)))
	if v97 == v16 {
		v163 = v9
		goto L23
	} else {
		goto L32
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v105 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v107)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = base.I64_extend_i32_u(base.B2i32(v98 != int32(112)) & base.B2i32(v98 != int32(73)))
	if v97 == v16 {
		v163 = v9
		goto L23
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v100 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v163 = v9
	goto L23
L32:
	;
	goto L24
L33:
	;
	v133 = v129
	goto L35
L34:
	;
	v133 = v128
	goto L35
L35:
	;
	v135 = v128
	goto L36
L36:
	;
	if v135 == v133 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v163 = base.I64_extend_i32_s(v153)
	goto L23
L38:
	;
	goto L37
L39:
	;
	v153 = v133
	goto L38
L40:
	;
	goto L41
L41:
	;
	v148 = v135 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v100)+12))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v135<<(uint(int32(2))%32)+v149)))
	if v151 != v16 {
		v135 = v148
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v153 = v148
	goto L38
L43:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v171)+16))
	v174 = F_HeapTupleHeaderGetDatum(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v176 = *(*int64)(unsafe.Add(mBase, uint32(v75)))
	*(*int64)(unsafe.Add(mBase, uint32(v75))) = v176 + int64(1)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(1)
	v220 = v174
	goto L1
L45:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
	v189 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
	v220 = v9
	goto L1
L46:
	;
	F_errmsg_internal(m, int32(_a_F_pg_partition_tree_1), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L7
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_pg_partition_tree_2), int32(91), int32(_a_F_pg_partition_tree_3))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
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
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+20)) = int32(2)
	v210 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v210)
	v220 = v9
	goto L1
}
func F_pg_postmaster_start_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_pg_postmaster_start_time[0]))
	return v3
}
func F_pg_prepared_statement(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v104 int32
	_ = v104
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v131 int64
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v172 int64
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
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
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int64
	_ = v269
	var v272 int32
	_ = v272
	var v279 int64
	_ = v279
	var v282 int32
	_ = v282
	var v289 int64
	_ = v289
	var v292 int32
	_ = v292
	var v299 int64
	_ = v299
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v340 int64
	_ = v340
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v380 int64
	_ = v380
	var v382 int32
	_ = v382
	var v383 int64
	_ = v383
	var v385 int64
	_ = v385
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_statement[0]))
	if v25 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v16 + int32(112)
	return int64(0)
L4:
	;
	v29 = v16 + int32(92)
	F_hash_seq_init(m, v29, v25)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v32 = F_hash_seq_search(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v32 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v41 = v32
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+52))
	v53 = F_cstring_to_text(m, v41)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+16)) = base.I64_extend_i32_u(v53)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v59 = F_cstring_to_text(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+24)) = base.I64_extend_i32_u(v59)
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v41)+72))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v69 = F_palloc_mul(m, int32(8), v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v68 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v193 = F_construct_array_builtin(m, v69, v68, int32(2206))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L25
	}
L14:
	;
	v74 = v68 & int32(3)
	v75 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v68) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v82 = v75
	v91 = int32(0)
	goto L18
L16:
	;
	v140 = v75
	goto L17
L17:
	;
	v153 = v140
	v157 = v75
	goto L22
L18:
	;
	v95 = int32(3)
	v98 = int32(2)
	v101 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v66+v82<<(uint(v98)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v82<<(uint(v95)%32)))) = v101
	v104 = v82 | int32(1)
	v111 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v66+v104<<(uint(v98)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v104<<(uint(v95)%32)))) = v111
	v114 = v82 | v98
	v121 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v66+v114<<(uint(v98)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v114<<(uint(v95)%32)))) = v121
	v124 = v82 | v95
	v131 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v66+v124<<(uint(v98)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v124<<(uint(v95)%32)))) = v131
	v133 = int32(4)
	v134 = v82 + v133
	v136 = v91 + v133
	if v136 != v68&int32(2147483644) {
		v82 = v134
		v91 = v136
		goto L18
	} else {
		goto L20
	}
L19:
	;
	if v74 == int32(0) {
		goto L13
	} else {
		goto L21
	}
L20:
	;
	goto L19
L21:
	;
	v140 = v134
	goto L17
L22:
	;
	v172 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v66+v153<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v153<<(uint(int32(3))%32)))) = v172
	v174 = int32(1)
	v177 = v157 + v174
	if v177 != v74 {
		v153 = v153 + v174
		v157 = v177
		goto L22
	} else {
		goto L24
	}
L23:
	;
	goto L13
L24:
	;
	goto L23
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = base.I64_extend_i32_u(v193)
	if v52 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v380 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v41)+68)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+56)) = v380
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v41)+64))
	v383 = *(*int64)(unsafe.Add(mBase, uint32(v382)+136))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v383
	v385 = *(*int64)(unsafe.Add(mBase, uint32(v382)+128))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+72)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	F_tuplestore_putvalues(m, v387, v388, v16+int32(16), v16+int32(8))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L52
	}
L27:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v199 = F_palloc_mul(m, int32(4), v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v365 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)) = uint8(v365)
	goto L26
L30:
	;
	v201 = int32(0)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v202 <= v201 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v361 = F_construct_array_builtin(m, v351, v348, int32(2206))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L51
	}
L32:
	;
	v206 = F_palloc_mul(m, int32(8), v202)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v208 = v201
	v209 = v202
	goto L36
L35:
	;
	v348 = v202
	v351 = v206
	goto L31
L36:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v52+v209<<(uint(int32(3))%32)+v208*int32(100))+96))
	*(*int32)(unsafe.Add(mBase, uint32(v199+v208<<(uint(int32(2))%32)))) = v230
	v233 = v208 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v233 < v234 {
		v208 = v233
		v209 = v234
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v237 = F_palloc_mul(m, int32(8), v234)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	if v234 <= int32(0) {
		v348 = v234
		v351 = v237
		goto L31
	} else {
		goto L40
	}
L40:
	;
	v242 = v234 & int32(3)
	v243 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v234) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v250 = v243
	v259 = int32(0)
	goto L44
L42:
	;
	v308 = v243
	goto L43
L43:
	;
	v321 = v308
	v328 = v243
	goto L48
L44:
	;
	v263 = int32(3)
	v266 = int32(2)
	v269 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v199+v250<<(uint(v266)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v237+v250<<(uint(v263)%32)))) = v269
	v272 = v250 | int32(1)
	v279 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v199+v272<<(uint(v266)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v237+v272<<(uint(v263)%32)))) = v279
	v282 = v250 | v266
	v289 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v199+v282<<(uint(v266)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v237+v282<<(uint(v263)%32)))) = v289
	v292 = v250 | v263
	v299 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v199+v292<<(uint(v266)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v237+v292<<(uint(v263)%32)))) = v299
	v301 = int32(4)
	v302 = v250 + v301
	v304 = v259 + v301
	if v304 != v234&int32(2147483644) {
		v250 = v302
		v259 = v304
		goto L44
	} else {
		goto L46
	}
L45:
	;
	if v242 == int32(0) {
		v348 = v234
		v351 = v237
		goto L31
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	v308 = v302
	goto L43
L48:
	;
	v340 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v199+v321<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v237+v321<<(uint(int32(3))%32)))) = v340
	v342 = int32(1)
	v345 = v328 + v342
	if v345 != v242 {
		v321 = v321 + v342
		v328 = v345
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v348 = v234
	v351 = v237
	goto L31
L50:
	;
	goto L49
L51:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = base.I64_extend_i32_u(v361)
	goto L26
L52:
	;
	v397 = F_hash_seq_search(m, v16+int32(92))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v397 != 0 {
		v41 = v397
		goto L8
	} else {
		goto L54
	}
L54:
	;
	goto L9
}
func F_pg_printf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v5 = m.G0
	v7 = v5 - int32(1072)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_pg_printf[0]))
	if v11 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_pg_printf[1])) = int32(28)
		m.G0 = v7 + int32(1072)
		return
	} else {
		v18 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v7)+1068)) = uint8(v18)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1064)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1060)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1056)) = v7 + int32(1040)
		v27 = v7 + int32(16)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1052)) = v27
		*(*int32)(unsafe.Add(mBase, uint32(v7)+1048)) = v27
		F_dopr(m, v7+int32(1048), l0, l1)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1068)))
			if v34 == int32(0) {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1048))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1052))
				if v37 != v38 {
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v7)+1060))
					v48 = F_fwrite(m, v38, int32(1), v37-v38, v47)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						m.G0 = v7 + int32(1072)
						return
					}
				} else {
					m.G0 = v7 + int32(1072)
					return
				}
			} else {
				m.G0 = v7 + int32(1072)
				return
			}
		}
	}
}
func F_pg_qsort_med3(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v7 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l1, l2)
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			if v7 < int32(0) {
				if v11 < int32(0) {
					v30 = l1
					return v30
				} else {
					v17 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l0, l2)
					v18 = m.ExcPending
					if v18 != 0 {
						return int32(0)
					} else {
						if v17 < int32(0) {
							v21 = l2
						} else {
							v21 = l0
						}
						return v21
					}
				}
			} else {
				if int32(0) < v11 {
					v30 = l1
					return v30
				} else {
					v25 = m.T0[l3].(func(*base.Module, int32, int32) int32)(m, l0, l2)
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if v25 < int32(0) {
							v29 = l0
						} else {
							v29 = l2
						}
						v30 = v29
						return v30
					}
				}
			}
		}
	}
}
func F_pg_random_bytes(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui32(int32(-1025)) < base.Ui32(v4-int32(1025)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_px_THROW_ERROR(m, int32(-17))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L25
	}
L2:
	;
	v10 = v4 + int32(4)
	v11 = F_palloc(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L21
	}
L5:
	;
	return int64(0)
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v10 << (uint(int32(2)) % 32)
	v20 = int32(0)
	v24 = m.G0
	v26 = v24 - int32(16)
	m.G0 = v26
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v20
	v32 = F_open(m, int32(_a_F_pg_random_bytes_0), v20, v26)
	mBase = m.M
	if v32 != int32(-1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v65 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v35 = int32(1)
	if v4 == int32(0) {
		v58 = v35
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v65 = v20
	goto L10
L10:
	;
	m.G0 = v26 + int32(16)
	goto L7
L11:
	;
	v60 = F_close(m, v32)
	mBase = m.M
	v65 = v58
	goto L10
L12:
	;
	v38 = v11 + int32(4)
	v39 = v4
	goto L13
L13:
	;
	v44 = F_read(m, v32, v38, v39)
	mBase = m.M
	if v44 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v58 = v35
	goto L11
L15:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_random_bytes[0]))
	if v48 == int32(27) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v53 = v39 - v44
	if v53 != 0 {
		v38 = v38 + v44
		v39 = v53
		goto L13
	} else {
		goto L19
	}
L18:
	;
	v58 = int32(0)
	goto L11
L19:
	;
	goto L14
L20:
	;
	return base.I64_extend_i32_u(v11)
L21:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_pg_random_bytes_1), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_pg_random_bytes_2), int32(465), int32(_a_F_pg_random_bytes_3))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_random_uuid(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_gen_random_uuid(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_pg_read_binary_file_off_len(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		if int64(0) <= v10 {
			v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			v14 = F_convert_and_check_filename(m, v6)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				v17 = F_read_binary_file(m, v14, v13, v10, int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int64(0)
				} else {
					if v17 == int32(0) {
						v21 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v17)
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_read_binary_file_off_len_0), int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_read_binary_file_off_len_1), int32(269), int32(_a_F_pg_read_binary_file_off_len_2))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
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
func F_pg_relpages(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
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
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = F_superuser(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int64(0)
		} else {
			if v7 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_relpages_0), int32(0))
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_relpages_1), int32(422), int32(_a_F_pg_relpages_2))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v27 = F_textToQualifiedNameList(m, v3)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v29 = F_makeRangeVarFromNameList(m, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int64(0)
					} else {
						v32 = F_relation_openrv(m, v29, int32(1))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int64(0)
						} else {
							v34 = F_pg_relpages_impl(m, v32)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int64(0)
							} else {
								return v34
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_relpages_v1_5(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = F_textToQualifiedNameList(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int64(0)
		} else {
			v9 = F_makeRangeVarFromNameList(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int64(0)
			} else {
				v12 = F_relation_openrv(m, v9, int32(1))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int64(0)
				} else {
					v14 = F_pg_relpages_impl(m, v12)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int64(0)
					} else {
						return v14
					}
				}
			}
		}
	}
}
func F_pg_relpagesbyid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_relpagesbyid_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_relpagesbyid_1), int32(454), int32(_a_F_pg_relpagesbyid_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v28 = F_relation_open(m, base.I32_wrap_i64(v3), int32(1))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = F_pg_relpages_impl(m, v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					return v30
				}
			}
		}
	}
}
func F_pg_server_to_client(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	if l1 <= int32(0) {
		v35 = l0
		return v35
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_pg_server_to_client[0]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		if v9 == int32(0) {
			v35 = l0
			return v35
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_server_to_client[1]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
			if v9 == v14 {
				v35 = l0
				return v35
			} else {
				if v14 == int32(0) {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v9*int32(28))+uint32(_c_F_pg_server_to_client[2])))
					v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, l0, l1)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						if l1 == v23 {
							v35 = l0
							return v35
						} else {
							F_report_invalid_encoding(m, v9, l0+v23, l1-v23)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v33 = F_perform_default_encoding_conversion(m, l0, l1, int32(0))
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						v35 = v33
						return v35
					}
				}
			}
		}
	}
}
func F_pg_signal_backend(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_BackendPidGetProc(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 == int32(0) {
			v18 = F_errstart(m, int32(19), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 == int32(0) {
					v83 = int32(1)
					m.G0 = v8 + int32(16)
					return v83
				} else {
					v69 = int32(_a_F_pg_signal_backend_0)
					v71 = int32(75)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg(m, v69, v8)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_signal_backend_1), v71, int32(_a_F_pg_signal_backend_2))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							v83 = int32(1)
							m.G0 = v8 + int32(16)
							return v83
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
			if v24 != 0 {
				v25 = F_superuser_arg(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v44 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[0]))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						v46 = F_has_privs_of_role(m, v44, v45)
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v46 != 0 {
								v55 = int32(0)
								v58 = F_pgmem_kill(m, v55-l0, l1)
								mBase = m.M
								if v58 == v55 {
									v83 = v55
									m.G0 = v8 + int32(16)
									return v83
								} else {
									v63 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v64 = m.ExcPending
									if v64 != 0 {
										return int32(0)
									} else {
										if v63 == int32(0) {
											v83 = int32(1)
											m.G0 = v8 + int32(16)
											return v83
										} else {
											v69 = int32(_a_F_pg_signal_backend_3)
											v71 = int32(121)
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
											F_errmsg(m, v69, v8)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_pg_signal_backend_1), v71, int32(_a_F_pg_signal_backend_2))
												mBase = m.M
												v78 = m.ExcPending
												if v78 != 0 {
													return int32(0)
												} else {
													v83 = int32(1)
													m.G0 = v8 + int32(16)
													return v83
												}
											}
										}
									}
								}
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[0]))
								v51 = F_has_privs_of_role(m, v49, int32(_a_F_pg_signal_backend_4))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									if v51 != 0 {
										v55 = int32(0)
										v58 = F_pgmem_kill(m, v55-l0, l1)
										mBase = m.M
										if v58 == v55 {
											v83 = v55
											m.G0 = v8 + int32(16)
											return v83
										} else {
											v63 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												if v63 == int32(0) {
													v83 = int32(1)
													m.G0 = v8 + int32(16)
													return v83
												} else {
													v69 = int32(_a_F_pg_signal_backend_3)
													v71 = int32(121)
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
													F_errmsg(m, v69, v8)
													mBase = m.M
													v74 = m.ExcPending
													if v74 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_signal_backend_1), v71, int32(_a_F_pg_signal_backend_2))
														mBase = m.M
														v78 = m.ExcPending
														if v78 != 0 {
															return int32(0)
														} else {
															v83 = int32(1)
															m.G0 = v8 + int32(16)
															return v83
														}
													}
												}
											}
										}
									} else {
										v83 = int32(2)
										m.G0 = v8 + int32(16)
										return v83
									}
								}
							}
						}
					} else {
						v29 = int32(4)
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
						if v30 == v29 {
							v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[0]))
							v36 = F_has_privs_of_role(m, v34, int32(_a_F_pg_signal_backend_5))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								if v36 == int32(0) {
									v83 = v29
									m.G0 = v8 + int32(16)
									return v83
								} else {
									v55 = int32(0)
									v58 = F_pgmem_kill(m, v55-l0, l1)
									mBase = m.M
									if v58 == v55 {
										v83 = v55
										m.G0 = v8 + int32(16)
										return v83
									} else {
										v63 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											if v63 == int32(0) {
												v83 = int32(1)
												m.G0 = v8 + int32(16)
												return v83
											} else {
												v69 = int32(_a_F_pg_signal_backend_3)
												v71 = int32(121)
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
												F_errmsg(m, v69, v8)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_signal_backend_1), v71, int32(_a_F_pg_signal_backend_2))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														v83 = int32(1)
														m.G0 = v8 + int32(16)
														return v83
													}
												}
											}
										}
									}
								}
							}
						} else {
							v40 = F_superuser(m)
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								if v40 != 0 {
									v55 = int32(0)
									v58 = F_pgmem_kill(m, v55-l0, l1)
									mBase = m.M
									if v58 == v55 {
										v83 = v55
										m.G0 = v8 + int32(16)
										return v83
									} else {
										v63 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											if v63 == int32(0) {
												v83 = int32(1)
												m.G0 = v8 + int32(16)
												return v83
											} else {
												v69 = int32(_a_F_pg_signal_backend_3)
												v71 = int32(121)
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
												F_errmsg(m, v69, v8)
												mBase = m.M
												v74 = m.ExcPending
												if v74 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_signal_backend_1), v71, int32(_a_F_pg_signal_backend_2))
													mBase = m.M
													v78 = m.ExcPending
													if v78 != 0 {
														return int32(0)
													} else {
														v83 = int32(1)
														m.G0 = v8 + int32(16)
														return v83
													}
												}
											}
										}
									}
								} else {
									v83 = int32(3)
									m.G0 = v8 + int32(16)
									return v83
								}
							}
						}
					}
				}
			} else {
				v29 = int32(4)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
				if v30 == v29 {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[0]))
					v36 = F_has_privs_of_role(m, v34, int32(_a_F_pg_signal_backend_5))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						if v36 == int32(0) {
							v83 = v29
							m.G0 = v8 + int32(16)
							return v83
						} else {
							v55 = int32(0)
							v58 = F_pgmem_kill(m, v55-l0, l1)
							mBase = m.M
							if v58 == v55 {
								v83 = v55
								m.G0 = v8 + int32(16)
								return v83
							} else {
								v63 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									if v63 == int32(0) {
										v83 = int32(1)
										m.G0 = v8 + int32(16)
										return v83
									} else {
										v69 = int32(_a_F_pg_signal_backend_3)
										v71 = int32(121)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg(m, v69, v8)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_signal_backend_1), v71, int32(_a_F_pg_signal_backend_2))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v83 = int32(1)
												m.G0 = v8 + int32(16)
												return v83
											}
										}
									}
								}
							}
						}
					}
				} else {
					v40 = F_superuser(m)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						if v40 != 0 {
							v55 = int32(0)
							v58 = F_pgmem_kill(m, v55-l0, l1)
							mBase = m.M
							if v58 == v55 {
								v83 = v55
								m.G0 = v8 + int32(16)
								return v83
							} else {
								v63 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return int32(0)
								} else {
									if v63 == int32(0) {
										v83 = int32(1)
										m.G0 = v8 + int32(16)
										return v83
									} else {
										v69 = int32(_a_F_pg_signal_backend_3)
										v71 = int32(121)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg(m, v69, v8)
										mBase = m.M
										v74 = m.ExcPending
										if v74 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_signal_backend_1), v71, int32(_a_F_pg_signal_backend_2))
											mBase = m.M
											v78 = m.ExcPending
											if v78 != 0 {
												return int32(0)
											} else {
												v83 = int32(1)
												m.G0 = v8 + int32(16)
												return v83
											}
										}
									}
								}
							}
						} else {
							v83 = int32(3)
							m.G0 = v8 + int32(16)
							return v83
						}
					}
				}
			}
		}
	}
}
func F_pg_snapshot_out(m *base.Module, l0 int32) int64 {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v17 = v9 - int32(-64)
	F_initStringInfo(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+48)) = v20
	F_appendStringInfo(m, v17, int32(_a_F_pg_snapshot_out_0), v9+int32(48))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+32)) = v27
	F_appendStringInfo(m, v17, int32(_a_F_pg_snapshot_out_0), v9+int32(32))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v34 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v79 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v9)+64)))
	m.G0 = v9 + int32(80)
	return v79
L7:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v37
	F_appendStringInfo(m, v17, int32(_a_F_pg_snapshot_out_1), v9+int32(16))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(v44) < base.Ui32(int32(2)) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v50 = int32(1)
	goto L10
L10:
	;
	v57 = v9 - int32(-64)
	F_appendStringInfoChar(m, v57, int32(44))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L6
L12:
	;
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v12+int32(24)+v50<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v9))) = v64
	F_appendStringInfo(m, v57, int32(_a_F_pg_snapshot_out_1), v9)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v70 = v50 + int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(v70) < base.Ui32(v71) {
		v50 = v70
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
}
func F_pg_snapshot_recv(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int32
	_ = v56
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pq_getmsgint(m, v10, int32(4))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L22
	}
L2:
	;
	return int64(0)
L3:
	;
	if base.Ui32(int32(134217724)) < base.Ui32(v12) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = F_pq_getmsgint64(m, v10)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v23 = F_pq_getmsgint64(m, v10)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if base.B2i32(base.I32_wrap_i64(v18) == int32(0))|base.B2i32(v23&int64(4294967295) == int64(0))|base.B2i32(base.Ui64(v23) < base.Ui64(v18)) != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v36 = F_palloc(m, v12<<(uint(int32(3))%32)+int32(24))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v18
	if v12 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v76<<(uint(int32(5))%32) + int32(96)
	return base.I64_extend_i32_u(v36)
L10:
	;
	v76 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v46 = int32(0)
	v47 = v12
	v54 = int64(0)
	goto L13
L13:
	;
	v55 = F_pq_getmsgint64(m, v10)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L2
	} else {
		goto L15
	}
L14:
	;
	v76 = v72
	goto L9
L15:
	;
	if base.B2i32(base.Ui64(v55) < base.Ui64(v54))|base.B2i32(base.Ui64(v55) < base.Ui64(v18))|base.B2i32(base.Ui64(v23) < base.Ui64(v55)) != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	if v55 == v54 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v71 < v72 {
		v46 = v71
		v47 = v72
		v54 = v73
		goto L13
	} else {
		goto L21
	}
L18:
	;
	v71 = v46
	v72 = v47 - int32(1)
	v73 = v54
	goto L17
L19:
	;
	goto L20
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v36+int32(24)+v46<<(uint(int32(3))%32)))) = v55
	v71 = v46 + int32(1)
	v72 = v47
	v73 = v55
	goto L17
L21:
	;
	goto L14
L22:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_pg_snapshot_recv_0), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_pg_snapshot_recv_1), int32(523), int32(_a_F_pg_snapshot_recv_2))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_sockaddr_cidr_mask(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int64
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v116 int64
	_ = v116
	var v118 int64
	_ = v118
	var v134 int32
	_ = v134
	v3 = l2
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l1 == int32(0) {
		if v3 == int32(2) {
			v18 = int32(32)
		} else {
			v18 = int32(128)
		}
		v31 = v18
		v33 = int32(-1)
		switch v3 - int32(2) {
		case 0:
			if base.Ui32(int32(32)) < base.Ui32(v31) {
				v134 = v33
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
				if v31 != 0 {
					v43 = int32(-1) << (uint(int32(32)-v31) % 32)
					v44 = int32(16711935)
					v55 = base.I32_rotr(v43&v44, int32(8)) | base.I32_rotr(v43, int32(24))&v44
				} else {
					v55 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v55
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
				v134 = int32(0)
			}
		default:
			v134 = v33
		case 8:
			if base.Ui32(int32(128)) < base.Ui32(v31) {
				v134 = v33
			} else {
				v61 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v61
				v64 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v64
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v64
				*(*int64)(unsafe.Add(mBase, uint32(v10))) = v64
				v73 = v61
				v76 = v31
				for {
					v79 = v73 + (v10 + int32(8))
					v80 = int32(0)
					if v76 <= v80 {
						v90 = v80
					} else {
						if base.Ui32(int32(7)) < base.Ui32(v76) {
							v90 = int32(255)
						} else {
							v90 = int32(255) << (uint(int32(8)-v76) % 32)
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v90)
					v92 = int32(0)
					v94 = v76 - int32(8)
					if v94 <= v92 {
						v104 = v92
					} else {
						if base.Ui32(int32(7)) < base.Ui32(v94) {
							v104 = int32(255)
						} else {
							v104 = int32(255) << (uint(int32(16)-v76) % 32)
						}
					}
					*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)) = uint8(v104)
					v106 = int32(16)
					v109 = v73 + int32(2)
					if v109 != v106 {
						v73 = v109
						v76 = v76 - v106
						continue
					} else {
						break
					}
					break
				}
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v112
				v114 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v114
				v116 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v116
				v118 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v118
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
				v134 = int32(0)
			}
		}
	} else {
		v23 = F_strtox_2(m, l1, v10+int32(28), int32(10), int64(2147483648))
		mBase = m.M
		v25 = int32(-1)
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
		if v26 == int32(0) {
			v134 = v25
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
			if v30 != 0 {
				v134 = v25
			} else {
				v31 = base.I32_wrap_i64(v23)
				v33 = int32(-1)
				switch v3 - int32(2) {
				case 0:
					if base.Ui32(int32(32)) < base.Ui32(v31) {
						v134 = v33
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
						if v31 != 0 {
							v43 = int32(-1) << (uint(int32(32)-v31) % 32)
							v44 = int32(16711935)
							v55 = base.I32_rotr(v43&v44, int32(8)) | base.I32_rotr(v43, int32(24))&v44
						} else {
							v55 = int32(0)
						}
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
						v134 = int32(0)
					}
				default:
					v134 = v33
				case 8:
					if base.Ui32(int32(128)) < base.Ui32(v31) {
						v134 = v33
					} else {
						v61 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v61
						v64 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v64
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v64
						*(*int64)(unsafe.Add(mBase, uint32(v10))) = v64
						v73 = v61
						v76 = v31
						for {
							v79 = v73 + (v10 + int32(8))
							v80 = int32(0)
							if v76 <= v80 {
								v90 = v80
							} else {
								if base.Ui32(int32(7)) < base.Ui32(v76) {
									v90 = int32(255)
								} else {
									v90 = int32(255) << (uint(int32(8)-v76) % 32)
								}
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v79))) = uint8(v90)
							v92 = int32(0)
							v94 = v76 - int32(8)
							if v94 <= v92 {
								v104 = v92
							} else {
								if base.Ui32(int32(7)) < base.Ui32(v94) {
									v104 = int32(255)
								} else {
									v104 = int32(255) << (uint(int32(16)-v76) % 32)
								}
							}
							*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)) = uint8(v104)
							v106 = int32(16)
							v109 = v73 + int32(2)
							if v109 != v106 {
								v73 = v109
								v76 = v76 - v106
								continue
							} else {
								break
							}
							break
						}
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v112
						v114 = *(*int64)(unsafe.Add(mBase, uint32(v10)+16))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v114
						v116 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
						*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v116
						v118 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
						*(*int64)(unsafe.Add(mBase, uint32(l0))) = v118
						*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v3)
						v134 = int32(0)
					}
				}
			}
		}
	}
	m.G0 = v10 + int32(32)
	return v134
}
func F_pg_split_walfile_name(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int64
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v10 = m.G0
	v12 = v10 - int32(352)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = F_text_to_cstring(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+318)) = uint16(v21)
	v23 = F_pstrdup(m, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v28 = v25
	v29 = v23
	goto L8
L6:
	;
	goto L7
L7:
	;
	v57 = F_strlen(m, v23)
	mBase = m.M
	if v57 != int32(24) {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	if base.Ui32((v28-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v43 = v28 - int32(32)
	goto L12
L11:
	;
	v43 = v28
	goto L12
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v43)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v45 != 0 {
		v28 = v45
		v29 = v29 + int32(1)
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L48
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L44
	}
L16:
	;
	v60 = int32(_a_F_pg_split_walfile_name_0)
	v64 = m.G0
	v66 = v64 - int32(32)
	v67 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v66)+24)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v66)+16)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v66))) = v67
	v75 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_split_walfile_name[0])))
	if v75 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v143 != int32(24) {
		goto L15
	} else {
		goto L36
	}
L18:
	;
	v143 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_split_walfile_name[1])))
	if v79 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v83 = v23
	goto L24
L22:
	;
	goto L23
L23:
	;
	v93 = v60
	v94 = v75
	goto L27
L24:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v89 == v75 {
		v83 = v83 + int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v143 = v83 - v23
	goto L17
L26:
	;
	goto L25
L27:
	;
	v101 = v66 + int32(base.Ui32(v94)>>(uint(int32(3))%32))&int32(28)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v103 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v102 | v103<<(uint(v94)%32)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+1)))
	if v107 != 0 {
		v93 = v93 + v103
		v94 = v107
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v110 == int32(0) {
		v133 = v23
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v143 = v133 - v23
	goto L17
L31:
	;
	v114 = v23
	v115 = v110
	goto L32
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v66+int32(base.Ui32(v115)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v123)>>(uint(v115)%32))&int32(1) == int32(0) {
		v133 = v114
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v133 = v131
	goto L30
L34:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	v131 = v114 + int32(1)
	if v129 != 0 {
		v114 = v131
		v115 = v129
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v147 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_split_walfile_name[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(348)
	v152 = v12 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v152
	v155 = v12 + int32(320)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v155
	v160 = F_sscanf(m, v23, int32(_a_F_pg_split_walfile_name_1), v12+int32(16))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v162 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+320)))
	v163 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+48)))
	v167 = F_get_call_result_type(m, l0, int32(0), v12+int32(312))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v167 != int32(1) {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v172 = base.I64_div_u_s(int64(4294967296), v147)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v172*v163 + v162
	v178 = F_pg_snprintf(m, v152, int32(256), int32(_a_F_pg_split_walfile_name_2), v12)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v185 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), base.I64_extend_i32_u(v152), int64(0), int64(-1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+320)) = v185
	v188 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+348)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+328)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v12)+312))
	v193 = F_heap_form_tuple(m, v190, v155, v12+int32(318))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v193)+16))
	v196 = F_HeapTupleHeaderGetDatum(m, v195)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	m.G0 = v12 + int32(352)
	return v196
L44:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v19
	F_errmsg(m, int32(_a_F_pg_split_walfile_name_3), v12+int32(32))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_split_walfile_name_4), int32(518), int32(_a_F_pg_split_walfile_name_5))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errmsg_internal(m, int32(_a_F_pg_split_walfile_name_6), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_pg_split_walfile_name_4), int32(523), int32(_a_F_pg_split_walfile_name_5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_statistics_obj_is_visible(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
	_ = v25
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_StatisticsObjIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v25 = int64(0)
		} else {
			v25 = base.I64_extend_i32_u(v14)
		}
		m.G0 = v7 + int32(16)
		return v25
	}
}
func F_pg_stats_ext_mcvlist_items(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int64
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int64
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int64
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int64
	_ = v169
	var v170 int32
	_ = v170
	var v172 int64
	_ = v172
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v216 int64
	_ = v216
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 + int32(-64)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L43
	}
L2:
	;
	v20 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	goto L17
L5:
	;
	return int64(0)
L6:
	;
	v24 = int32(_a_F_pg_stats_ext_mcvlist_items_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0])) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v30 = F_pg_detoast_datum(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v32 = F_statext_mcv_deserialize(m, v30)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v32
	if v32 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v32)+8)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v37
	goto L11
L10:
	;
	goto L11
L11:
	;
	v42 = F_get_call_result_type(m, l0, int32(0), v12+int32(-48))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if v42 != int32(1) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v47 = F_BlessTupleDesc(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v47
	v50 = F_TupleDescGetAttInMetadata(m, v47)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v50
	*(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0])) = v25
	goto L4
L16:
	;
	m.G0 = v14 - int32(-64)
	return v216
L17:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v59)+8))
	if base.Ui64(v60) < base.Ui64(v61) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	v69 = v63 + base.I32_wrap_i64(v60)*int32(24) + int32(48)
	v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+12)))
	if int32(0) < v70 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L42
	}
L21:
	;
	v79 = int32(0)
	v82 = v2
	v83 = v2
	goto L24
L22:
	;
	v154 = v2
	v155 = v2
	v159 = v60
	goto L23
L23:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = base.I64_extend32_s(v159)
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v164 = F_makeArrayResult(m, v155, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L38
	}
L24:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v89 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v87+v79))))
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v94 = F_accumArrayResult(m, v82, v89, int32(0), int32(16), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	v154 = v94
	v155 = v142
	v159 = v147
	goto L23
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96+v79))))
	if v98 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v144 = v79 + int32(1)
	v145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+12)))
	if v144 < v145 {
		v79 = v144
		v82 = v94
		v83 = v142
		goto L24
	} else {
		goto L37
	}
L28:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(16)+v79<<(uint(int32(2))%32))))
	F_getTypeOutputInfo(m, v104, v12+int32(-56), v12+int32(-57))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v139 = F_accumArrayResult(m, v83, int64(0), int32(1), int32(25), v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L36
	}
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v113 = v12 + int32(-48)
	F_fmgr_info(m, v111, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v69)+20))
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v117+v79<<(uint(int32(3))%32))))
	v122 = F_FunctionCall1Coll(m, v113, int32(0), v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v125 = F_cstring_to_text(m, base.I32_wrap_i64(v122))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v132 = F_accumArrayResult(m, v83, base.I64_extend_i32_u(v125), int32(0), int32(25), v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v142 = v132
	goto L27
L36:
	;
	v142 = v139
	goto L27
L37:
	;
	goto L25
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+24)) = v164
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v169 = F_makeArrayResult(m, v154, v168)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v169
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+40)) = v172
	v174 = *(*int64)(unsafe.Add(mBase, uint32(v69)+8))
	v175 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v175
	*(*int64)(unsafe.Add(mBase, uint32(v14)+48)) = v174
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+12)) = uint8(v175)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v186 = F_heap_form_tuple(m, v181, v12+int32(-48), v12+int32(-56))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v186)+16))
	v189 = F_HeapTupleHeaderGetDatum(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v191 = *(*int64)(unsafe.Add(mBase, uint32(v59)))
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = v191 + int64(1)
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+20)) = int32(1)
	v216 = v189
	goto L16
L42:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v200)+20)) = int32(2)
	v203 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v203)
	v216 = int64(0)
	goto L16
L43:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L44
	}
L44:
	;
	F_errmsg(m, int32(_a_F_pg_stats_ext_mcvlist_items_1), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(_a_F_pg_stats_ext_mcvlist_items_2), int32(1367), int32(_a_F_pg_stats_ext_mcvlist_items_3))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_strtitle(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	v6 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v10 == v6 {
		v13 = int32(0)
		if base.B2i32(l1 == v13)|base.B2i32(l3 == v13) == v13 {
			v20 = int32(1)
			v21 = l3 - v20
			v23 = l1 - v20
			if base.Ui32(v21) < base.Ui32(v23) {
				v25 = v21
			} else {
				v25 = v23
			}
			v31 = int32(0)
			v33 = v6
			for {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v31))))
				if v33 != 0 {
					v42 = v37 - int32(65)
					if base.Ui32(v42&int32(255)) < base.Ui32(int32(26)) {
						v47 = v37 | int32(32)
					} else {
						v47 = v37
					}
					v59 = v42
					v60 = v47
				} else {
					if base.Ui32((v37-int32(97))&int32(255)) < base.Ui32(int32(26)) {
						v58 = v37 - int32(32)
					} else {
						v58 = v37
					}
					v59 = v37 - int32(65)
					v60 = v58
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v31))) = uint8(v60)
				v65 = int32(255)
				v71 = int32(26)
				v82 = v31 + int32(1)
				if v31 != v25 {
					v31 = v82
					v33 = base.B2i32(base.Ui32((v37-int32(48))&v65) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v59&v65) < base.Ui32(v71)) | base.B2i32(base.Ui32((v37-int32(97))&v65) < base.Ui32(v71))
					continue
				} else {
					break
				}
				break
			}
			if base.Ui32(l1) <= base.Ui32(v82) {
				v111 = l3
				return v111
			} else {
				v94 = v25 + int32(1)
				v100 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v94))) = uint8(v100)
				return l3
			}
		} else {
			v87 = int32(0)
			if l1 == v87 {
				v111 = l3
				return v111
			} else {
				v94 = v87
				v100 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v94))) = uint8(v100)
				return l3
			}
		}
	} else {
		v103 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v104 = m.T0[v103].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v107 = m.ExcPending
		if v107 != 0 {
			return int32(0)
		} else {
			v111 = v104
			return v111
		}
	}
}
func F_pg_strupper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v111 int32
	_ = v111
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	v6 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v11 == v6 {
		v14 = int32(0)
		if base.B2i32(l1 == v14)|base.B2i32(l3 == v14) == v14 {
			v21 = int32(1)
			v22 = l3 - v21
			v24 = l1 - v21
			if base.Ui32(v22) < base.Ui32(v24) {
				v26 = v22
			} else {
				v26 = v24
			}
			if v26 == int32(0) {
				v86 = int32(0)
				v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v86))))
				if base.Ui32((v94-int32(97))&int32(255)) < base.Ui32(int32(26)) {
					v103 = v94 - int32(32)
				} else {
					v103 = v94
				}
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v86))) = uint8(v103)
				v111 = v86 + int32(1)
			} else {
				v30 = int32(1)
				v31 = v26 + v30
				v41 = int32(0)
				v46 = v6
				for {
					v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v41))))
					if base.Ui32((v49-int32(97))&int32(255)) < base.Ui32(int32(26)) {
						v58 = v49 - int32(32)
					} else {
						v58 = v49
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v41))) = uint8(v58)
					v61 = v41 | int32(1)
					v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v61))))
					if base.Ui32((v64-int32(97))&int32(255)) < base.Ui32(int32(26)) {
						v73 = v64 - int32(32)
					} else {
						v73 = v64
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v61))) = uint8(v73)
					v75 = int32(2)
					v76 = v41 + v75
					v78 = v46 + v75
					if v78 != v31&int32(-2) {
						v41 = v76
						v46 = v78
						continue
					} else {
						break
					}
					break
				}
				if v31&v30 == int32(0) {
					v111 = v76
				} else {
					v86 = v76
					v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2+v86))))
					if base.Ui32((v94-int32(97))&int32(255)) < base.Ui32(int32(26)) {
						v103 = v94 - int32(32)
					} else {
						v103 = v94
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l0+v86))) = uint8(v103)
					v111 = v86 + int32(1)
				}
			}
			if base.Ui32(l1) <= base.Ui32(v111) {
				v145 = l3
				return v145
			} else {
				v127 = v26 + int32(1)
				v134 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v127))) = uint8(v134)
				return l3
			}
		} else {
			v120 = int32(0)
			if l1 == v120 {
				v145 = l3
				return v145
			} else {
				v127 = v120
				v134 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l0+v127))) = uint8(v134)
				return l3
			}
		}
	} else {
		v137 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
		v138 = m.T0[v137].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, l0, l1, l2, l3, l4)
		mBase = m.M
		v141 = m.ExcPending
		if v141 != 0 {
			return int32(0)
		} else {
			v145 = v138
			return v145
		}
	}
}
func F_pg_timezone_abbrevs_abbrevs(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int64
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v169 int32
	_ = v169
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
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
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
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int64
	_ = v318
	var v319 int64
	_ = v319
	var v322 int64
	_ = v322
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int64
	_ = v350
	var v351 int32
	_ = v351
	var v352 int64
	_ = v352
	var v356 int32
	_ = v356
	var v366 int64
	_ = v366
	var v374 int32
	_ = v374
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+78)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+76)) = uint16(v2)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	if v18 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L5
	} else {
		goto L86
	}
L2:
	;
	v21 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	goto L10
L5:
	;
	return int64(0)
L6:
	;
	v25 = int32(_a_F_pg_timezone_abbrevs_abbrevs_0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_abbrevs[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_abbrevs[0])) = v28
	v31 = F_palloc(m, int32(4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v31
	v39 = F_get_call_result_type(m, l0, v33, v11+int32(80))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v39 != int32(1) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v43
	*(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_abbrevs[0])) = v26
	goto L4
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_abbrevs[1]))
	if v53 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	m.G0 = v11 + int32(112)
	return v366
L12:
	;
	v69 = v53 + v55<<(uint(int32(4))%32)
	v71 = v69 + int32(8)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+19)))
	switch v72 - int32(5) {
	case 0:
		goto L19
	case 1:
		goto L22
	case 2:
		goto L21
	default:
		goto L20
	}
L13:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v55 < v56 {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v62)+20)) = int32(2)
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v65)
	v366 = int64(0)
	goto L11
L18:
	;
	v146 = v11 + int32(65)
	goto L40
L19:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v143 = v140
	v144 = int64(0)
	goto L18
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L5
	} else {
		goto L34
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v78 = v53 + v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v79 != 0 {
		v112 = v79
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v143 = v75
	v144 = int64(1)
	goto L18
L23:
	;
	v116 = *(*int64)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_abbrevs[2]))
	v119 = F_DetermineTimeZoneAbbrevOffsetTS(m, v116, v71, v112, v11+int32(80))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L33
	}
L24:
	;
	v81 = v78 + int32(4)
	v82 = F_pg_tzset(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v78))) = v82
	if v82 != 0 {
		v112 = v82
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v85 = int32(0)
	v87 = F_errsave_start(m, v85)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	if v87 == int32(0) {
		v112 = v85
		goto L23
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v81
	F_errmsg(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_1), v11+int32(32))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v71
	v104 = F_errdetail(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_2), v11+int32(16))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errsave_finish(m, int32(0), int32(_a_F_pg_timezone_abbrevs_abbrevs_3), int32(_a_F_pg_timezone_abbrevs_abbrevs_4), int32(_a_F_pg_timezone_abbrevs_abbrevs_5))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v112 = v85
	goto L23
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	v143 = int32(0) - v119
	v144 = base.I64_extend_i32_u(base.B2i32(v122 != int32(0)))
	goto L18
L34:
	;
	v130 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v130
	F_errmsg_internal(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_6), v11)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	F_errfinish(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_3), int32(_a_F_pg_timezone_abbrevs_abbrevs_7), int32(_a_F_pg_timezone_abbrevs_abbrevs_8))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+65)))
	if v266 != 0 {
		goto L68
	} else {
		goto L69
	}
L38:
	;
	v263 = F_strlen(m, v252)
	mBase = m.M
	goto L37
L40:
	;
	goto L41
L41:
	;
	v153 = int32(10)
	if (v146^v71)&int32(3) != 0 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v256 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v253))) = uint8(v256)
	goto L38
L43:
	;
	v237 = v232
	v238 = v233
	v239 = v234
	goto L64
L44:
	;
	if v227 == int32(0) {
		v252 = v225
		v253 = v226
		goto L42
	} else {
		goto L63
	}
L45:
	;
	v225 = v71
	v226 = v146
	v227 = v153
	goto L44
L46:
	;
	goto L47
L47:
	;
	v157 = int32(0)
	if base.B2i32(v71&int32(3) == v157)|int32(0) == v157 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	if v193 == int32(0) {
		v252 = v190
		v253 = v191
		goto L42
	} else {
		goto L57
	}
L49:
	;
	v169 = v71
	v170 = v146
	v171 = v153
	goto L52
L50:
	;
	goto L51
L51:
	;
	v190 = v71
	v191 = v146
	v192 = v153
	v193 = int32(1)
	goto L48
L52:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169))))
	*(*uint8)(unsafe.Add(mBase, uint32(v170))) = uint8(v173)
	if v173 == int32(0) {
		v232 = v169
		v233 = v170
		v234 = v171
		goto L43
	} else {
		goto L54
	}
L53:
	;
	v190 = v184
	v191 = v178
	v192 = v180
	v193 = v182
	goto L48
L54:
	;
	v177 = int32(1)
	v178 = v170 + v177
	v180 = v171 - v177
	v181 = int32(0)
	v182 = base.B2i32(v180 != v181)
	v184 = v169 + v177
	if v184&int32(3) == v181 {
		v190 = v184
		v191 = v178
		v192 = v180
		v193 = v182
		goto L48
	} else {
		goto L55
	}
L55:
	;
	if v180 != 0 {
		v169 = v184
		v170 = v178
		v171 = v180
		goto L52
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if base.B2i32(v196 == int32(0))|base.B2i32(base.Ui32(v192) < base.Ui32(int32(4))) != 0 {
		v225 = v190
		v226 = v191
		v227 = v192
		goto L44
	} else {
		goto L58
	}
L58:
	;
	v203 = v190
	v204 = v191
	v205 = v192
	goto L59
L59:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v211 = int32(-2139062144)
	if (int32(16843008)-v208|v208)&v211 != v211 {
		v232 = v203
		v233 = v204
		v234 = v205
		goto L43
	} else {
		goto L61
	}
L60:
	;
	v225 = v219
	v226 = v217
	v227 = v221
	goto L44
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v208
	v216 = int32(4)
	v217 = v204 + v216
	v219 = v203 + v216
	v221 = v205 - v216
	if base.Ui32(int32(3)) < base.Ui32(v221) {
		v203 = v219
		v204 = v217
		v205 = v221
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v232 = v225
	v233 = v226
	v234 = v227
	goto L43
L64:
	;
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	*(*uint8)(unsafe.Add(mBase, uint32(v238))) = uint8(v241)
	if v241 == int32(0) {
		v252 = v237
		v253 = v238
		goto L42
	} else {
		goto L66
	}
L65:
	;
	v252 = v248
	v253 = v246
	goto L42
L66:
	;
	v245 = int32(1)
	v246 = v238 + v245
	v248 = v237 + v245
	v250 = v239 - v245
	if v250 != 0 {
		v237 = v248
		v238 = v246
		v239 = v250
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v269 = v146
	v270 = v266
	goto L71
L69:
	;
	goto L70
L70:
	;
	v300 = F_cstring_to_text(m, v11+int32(65))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L5
	} else {
		goto L78
	}
L71:
	;
	if base.Ui32((v270-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L70
L73:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v269))) = uint8(v285)
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	if v287 != 0 {
		v269 = v269 + int32(1)
		v270 = v287
		goto L71
	} else {
		goto L77
	}
L74:
	;
	v283 = v270 - int32(32)
	goto L76
L75:
	;
	v283 = v270
	goto L76
L76:
	;
	v285 = v283 & int32(255)
	goto L73
L77:
	;
	goto L72
L78:
	;
	v302 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v302
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v302
	*(*int64)(unsafe.Add(mBase, uint32(v11)+80)) = base.I64_extend_i32_u(v300)
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = base.I64_extend_i32_s(v143) * int64(1000000)
	v313 = v11 + int32(40)
	v315 = F_palloc(m, int32(16))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v318 = int64(*(*int32)(unsafe.Add(mBase, uint32(v313)+12)))
	v319 = int64(*(*int32)(unsafe.Add(mBase, uint32(v313)+16)))
	v322 = v318 + v319*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v322-int64(2147483648)) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v11)+96)) = v144
	*(*int64)(unsafe.Add(mBase, uint32(v11)+88)) = base.I64_extend_i32_u(v315)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v338 + int32(1)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v347 = F_heap_form_tuple(m, v342, v11+int32(80), v11+int32(76))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L5
	} else {
		goto L84
	}
L81:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v315)+12)) = uint32(v322)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v313)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+8)) = v328
	v330 = *(*int64)(unsafe.Add(mBase, uint32(v313)))
	*(*int64)(unsafe.Add(mBase, uint32(v315))) = v330
	goto L83
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v347)+16))
	v350 = F_HeapTupleHeaderGetDatum(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v352 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v352 + int64(1)
	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v356)+20)) = int32(1)
	v366 = v350
	goto L11
L86:
	;
	F_errmsg_internal(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_9), int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_3), int32(_a_F_pg_timezone_abbrevs_abbrevs_10), int32(_a_F_pg_timezone_abbrevs_abbrevs_8))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_total_relation_size(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_try_relation_open(m, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int64(0)
		} else {
			v17 = F_calculate_table_size(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = F_calculate_indexes_size(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					F_relation_close(m, v7, int32(1))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return v17 + v19
					}
				}
			}
		}
	}
}
func F_pg_typeof(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = F_get_fn_expr_argtype(m, v2, int32(0))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v4)
	}
}
func F_pg_ultostr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	v3 = int32(0)
	if l1 == v3 {
		v12 = int32(48)
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v12)
		v122 = int32(1)
	} else {
		v18 = int32(1233)
		v23 = int32(base.Ui32((base.I32_clz(l1)^int32(31))*v18+v18) >> (uint(int32(12)) % 32))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v23<<(uint(int32(2))%32))+uint32(_c_F_pg_ultostr[0])))
		v28 = v23 + base.B2i32(base.Ui32(v26) <= base.Ui32(l1))
		if base.Ui32(int32(_a_F_pg_ultostr_0)) <= base.Ui32(l1) {
			v32 = l1
			v34 = v3
			for {
				v41 = l0 + v28 - v34
				v42 = int32(4)
				v45 = base.I32_div_u_s(v32, int32(_a_F_pg_ultostr_0))
				v48 = v32 + v45*int32(-10000)
				v49 = int32(100)
				v50 = base.I32_div_u_s(v48, v49)
				v51 = int32(1)
				v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50<<(uint(v51)%32))+uint32(_c_F_pg_ultostr[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v41-v42))) = uint16(v53)
				v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v48-v50*v49)<<(uint(v51)%32))+uint32(_c_F_pg_ultostr[1]))))
				*(*uint16)(unsafe.Add(mBase, uint32(v41-int32(2)))) = uint16(v62)
				v65 = v34 + v42
				if base.Ui32(int32(99999999)) < base.Ui32(v32) {
					v32 = v45
					v34 = v65
					continue
				} else {
					break
				}
				break
			}
			v68 = v45
			v70 = v65
		} else {
			v68 = l1
			v70 = v3
		}
		if base.Ui32(int32(100)) <= base.Ui32(v68) {
			v81 = int32(2)
			v83 = int32(_a_F_pg_ultostr_1)
			v85 = int32(100)
			v86 = base.I32_div_u_s(v68&v83, v85)
			v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32((v68-v86*v85)&v83<<(uint(int32(1))%32))+uint32(_c_F_pg_ultostr[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l0+v28-v70-v81))) = uint16(v94)
			v98 = v86
			v99 = v70 | v81
		} else {
			v98 = v68
			v99 = v70
		}
		if base.Ui32(int32(10)) <= base.Ui32(v98) {
			v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98<<(uint(int32(1))%32))+uint32(_c_F_pg_ultostr[1]))))
			*(*uint16)(unsafe.Add(mBase, uint32(l0+v28-v99-int32(2)))) = uint16(v108)
			v122 = v28
		} else {
			v111 = v98 | int32(48)
			*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v111)
			v122 = v28
		}
	}
	return v122 + l0
}
func F_pg_utf2wchar_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v35 int32
	_ = v35
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v19 == int32(0) {
		v109 = v14
		v113 = v18
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(0)
	return v113
L6:
	;
	goto L5
L7:
	;
	if int32(0) <= base.I32_extend8_s(v19) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v97
	v102 = v18 + int32(1)
	v104 = v14 + int32(4)
	v105 = v15 + v98
	if int32(0) < v105 {
		v13 = v99
		v14 = v104
		v15 = v105
		v18 = v102
		goto L4
	} else {
		goto L21
	}
L9:
	;
	v97 = v19
	v98 = int32(-1)
	v99 = v13 + int32(1)
	goto L8
L10:
	;
	if v19&int32(224) == int32(192) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v15 == int32(1) {
		v109 = v14
		v113 = v18
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v19&int32(240) == int32(224) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v97 = v19<<(uint(int32(6))%32)&int32(1984) | v35&int32(63)
	v98 = int32(-2)
	v99 = v13 + int32(2)
	goto L8
L15:
	;
	if base.Ui32(v15) < base.Ui32(int32(3)) {
		v109 = v14
		v113 = v18
		goto L6
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if v19&int32(248) != int32(240) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v49 = int32(63)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v97 = v48&v49 | (v19<<(uint(int32(12))%32)&int32(_a_F_pg_utf2wchar_with_len_0) | v55&v49<<(uint(int32(6))%32))
	v98 = int32(-3)
	v99 = v13 + int32(3)
	goto L8
L19:
	;
	if base.Ui32(v15) < base.Ui32(int32(4)) {
		v109 = v14
		v113 = v18
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+3)))
	v72 = int32(63)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+2)))
	v97 = v71&v72 | (v19<<(uint(int32(18))%32)&int32(_a_F_pg_utf2wchar_with_len_1) | v78&v72<<(uint(int32(12))%32) | v84&v72<<(uint(int32(6))%32))
	v98 = int32(-4)
	v99 = v13 + int32(4)
	goto L8
L21:
	;
	v109 = v104
	v113 = v102
	goto L6
}
func F_pg_verifymbstr(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_verifymbstr[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7*int32(28))+uint32(_c_F_pg_verifymbstr[1])))
	v13 = m.T0[v12].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if l1 != v13 {
			F_report_invalid_encoding(m, v7, l0+v13, l1-v13)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return
		}
	}
}
func F_pg_wchar2utf_with_len(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v4 = int32(0)
	if l2 <= v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v9 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v9)
	return v9
L2:
	;
	goto L3
L3:
	;
	v13 = l0
	v14 = l1
	v15 = l2
	v18 = v4
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v105 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v105)
	return v104
L6:
	;
	if base.Ui32(v19) <= base.Ui32(int32(127)) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v100 = v14
	v104 = v18
	goto L8
L8:
	;
	goto L5
L9:
	;
	v91 = int32(1)
	v95 = v14 + v90
	v96 = v90 + v18
	if v91 < v15 {
		v13 = v13 + int32(4)
		v14 = v95
		v15 = v15 - v91
		v18 = v96
		goto L4
	} else {
		goto L24
	}
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v19)
	v90 = int32(1)
	goto L9
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v19) <= base.Ui32(int32(2047)) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v75 = v19&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v70+v14))) = uint8(v75)
	if v69&int32(224) == int32(192) {
		v90 = int32(2)
		goto L9
	} else {
		goto L20
	}
L14:
	;
	v29 = int32(base.Ui32(v19)>>(uint(int32(6))%32)) | int32(-64)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v29)
	v69 = v29
	v70 = int32(1)
	goto L13
L15:
	;
	goto L16
L16:
	;
	if base.Ui32(v19) <= base.Ui32(int32(_a_F_pg_wchar2utf_with_len_0)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v37 = int32(base.Ui32(v19)>>(uint(int32(12))%32)) | int32(-32)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v37)
	v44 = int32(base.Ui32(v19)>>(uint(int32(6))%32))&int32(63) | int32(128)
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v44)
	v69 = v37
	v70 = int32(2)
	goto L13
L18:
	;
	goto L19
L19:
	;
	v49 = int32(63)
	v51 = int32(128)
	v52 = int32(base.Ui32(v19)>>(uint(int32(6))%32))&v49 | v51
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+2)) = uint8(v52)
	v59 = int32(base.Ui32(v19)>>(uint(int32(12))%32))&v49 | v51
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)) = uint8(v59)
	v66 = int32(base.Ui32(v19)>>(uint(int32(18))%32))&int32(7) | int32(-16)
	*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v66)
	v69 = v66
	v70 = int32(3)
	goto L13
L20:
	;
	if v69&int32(240) == int32(224) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v88 = int32(3)
	goto L23
L22:
	;
	v88 = int32(4)
	goto L23
L23:
	;
	v90 = v88
	goto L9
L24:
	;
	v100 = v95
	v104 = v96
	goto L8
}
func F_pg_xact_commit_timestamp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_TransactionIdGetCommitTsData(m, v8, v6+int32(8), int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v22 = int64(0)
		} else {
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			v22 = v21
		}
		m.G0 = v6 + int32(16)
		return v22
	}
}
