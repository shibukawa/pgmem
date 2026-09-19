package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

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
								F_errdetail(m, int32(_a_F_ValidatePgVersion_4), v10)
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1793), int32(_a_F_ValidatePgVersion_6))
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
								F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1797), int32(_a_F_ValidatePgVersion_6))
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
									F_errdetail(m, int32(_a_F_ValidatePgVersion_9), v10+int32(48))
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
											F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1811), int32(_a_F_ValidatePgVersion_6))
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
										F_errdetail(m, int32(_a_F_ValidatePgVersion_9), v10+int32(48))
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
												F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1811), int32(_a_F_ValidatePgVersion_6))
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
												F_errdetail(m, int32(_a_F_ValidatePgVersion_12), v10+int32(80))
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return
												} else {
													F_errfinish(m, int32(_a_F_ValidatePgVersion_5), int32(1821), int32(_a_F_ValidatePgVersion_6))
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
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
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
					F_DefineCustomStringVariable(m, int32(_a_F__PG_init_plpgsql_10), int32(_a_F__PG_init_plpgsql_11), int32(_a_F__PG_init_plpgsql_12), int32(_a_F__PG_init_plpgsql_13))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_DefineCustomStringVariable(m, int32(_a_F__PG_init_plpgsql_14), int32(_a_F__PG_init_plpgsql_15), int32(_a_F__PG_init_plpgsql_16), int32(_a_F__PG_init_plpgsql_17))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_MarkGUCPrefixReserved(m, int32(_a_F__PG_init_plpgsql_18))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								v48 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[1]))
								v50 = F_MemoryContextAlloc(m, v48, int32(12))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v50)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v50)+4)) = int32(_a_F__PG_init_plpgsql_19)
									v56 = int32(_a_F__PG_init_plpgsql_20)
									v57 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[2]))
									*(*int32)(unsafe.Add(mBase, uint32(v50))) = v57
									*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[2])) = v50
									v62 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[1]))
									v64 = F_MemoryContextAlloc(m, v62, int32(12))
									mBase = m.M
									v65 = m.ExcPending
									if v65 != 0 {
										return
									} else {
										v66 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v64)+8)) = v66
										*(*int32)(unsafe.Add(mBase, uint32(v64)+4)) = int32(_a_F__PG_init_plpgsql_21)
										v70 = int32(_a_F__PG_init_plpgsql_22)
										v71 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[3]))
										*(*int32)(unsafe.Add(mBase, uint32(v64))) = v71
										*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[3])) = v64
										v75 = m.G0
										v77 = v75 - int32(48)
										m.G0 = v77
										v80 = *(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[4]))
										if v80 == v66 {
											*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = int64(292057776192)
											v89 = F_hash_create(m, int32(_a_F__PG_init_plpgsql_23), int32(16), v77, int32(24))
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[4])) = v89
												v92 = v89
												v95 = F_hash_search(m, v92, int32(_a_F__PG_init_plpgsql_24), int32(1), v77)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return
												} else {
													v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
													if v97 == int32(0) {
														*(*int32)(unsafe.Add(mBase, uint32(v95)+64)) = int32(0)
													} else {
													}
													m.G0 = v77 + int32(48)
													v106 = int32(1)
													*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[0])) = uint8(v106)
													*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[5])) = v95 - int32(-64)
													return
												}
											}
										} else {
											v92 = v80
											v95 = F_hash_search(m, v92, int32(_a_F__PG_init_plpgsql_24), int32(1), v77)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return
											} else {
												v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
												if v97 == int32(0) {
													*(*int32)(unsafe.Add(mBase, uint32(v95)+64)) = int32(0)
												} else {
												}
												m.G0 = v77 + int32(48)
												v106 = int32(1)
												*(*uint8)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[0])) = uint8(v106)
												*(*int32)(unsafe.Add(mBase, _c_F__PG_init_plpgsql[5])) = v95 - int32(-64)
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
				F_errfinish(m, int32(_a_F_create_pg_locale_icu_1), int32(215), int32(_a_F_create_pg_locale_icu_2))
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
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v136 int64
	_ = v136
	var v139 int32
	_ = v139
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
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int64
	_ = v199
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v354 int32
	_ = v354
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v376 int64
	_ = v376
	var v377 int32
	_ = v377
	var v382 int64
	_ = v382
	var v383 int64
	_ = v383
	var v385 int64
	_ = v385
	var v386 int64
	_ = v386
	var v389 int64
	_ = v389
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int64
	_ = v399
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int64
	_ = v409
	var v410 int64
	_ = v410
	var v412 int64
	_ = v412
	var v413 int64
	_ = v413
	var v416 int64
	_ = v416
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+316))
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
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L28
	}
L21:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[1]))
	F_s_lock(m, v60+int32(440), int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_1), int32(_a_F_do_pg_backup_stop_2))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[1]))
	v70 = *(*int64)(unsafe.Add(mBase, uint32(v69)+432))
	v71 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v69)+440)), uint32(v71))
	v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	if base.Ui64(v74) <= base.Ui64(v70) {
		goto L6
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[4]))
	v81 = F_LWLockAcquire(m, v77+int32(1152), int32(1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L26
	}
L26:
	;
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[5]))
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v84)+136))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1088)) = v85
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+144))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1096)) = v87
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[4]))
	F_LWLockRelease(m, v90+int32(1152))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
	v101 = m.ExcPending
	if v101 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	v104 = F_XLogInsert(m, int32(0), int32(80))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L13
	} else {
		goto L30
	}
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1088)) = v104
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[1]))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)+308))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+1096)) = v109
	F_XLogBeginInsert(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v115 = F_XLogInsert(m, int32(0), int32(64))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v117 = F_time(m)
	mBase = m.M
	*(*int64)(unsafe.Add(mBase, uint32(l0)+1104)) = v117
	v119 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = v120
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+140)) = base.I32_wrap_i64(v119) & (v124 - int32(1))
	v129 = base.I64_extend_i32_s(v124)
	v130 = base.I64_div_u_s(v119, v129)
	v132 = base.I64_div_u_s(int64(4294967296), v129)
	v133 = base.I64_div_u_s(v130, v132)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+132)) = uint32(v133)
	v136 = v130 - v132*v133
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+136)) = uint32(v136)
	v139 = v13 + int32(208)
	v144 = F_pg_snprintf(m, v139, int32(1024), int32(_a_F_do_pg_backup_stop_3), v13+int32(128))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v147 = F_AllocateFile(m, v139, int32(_a_F_do_pg_backup_stop_4))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	if v147 == int32(0) {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v152 = F_build_backup_content(m, l0, int32(1))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = v152
	v158 = F_pg_fprintf(m, v147, int32(_a_F_do_pg_backup_stop_5), v13+int32(112))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_pfree(m, v152)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	v162 = F_fflush(m, v147)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L13
	} else {
		goto L39
	}
L39:
	;
	if v162 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	goto L41
L41:
	;
	if int32(base.Ui32(v164)>>(uint(int32(5))%32))&int32(1) != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v169 = F_FreeFile(m, v147)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L13
	} else {
		goto L43
	}
L43:
	;
	if v169 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	v172 = F_AllocateDir(m, int32(_a_F_do_pg_backup_stop_6))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v175 = F_ReadDir(m, v172, int32(_a_F_do_pg_backup_stop_6))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	if v175 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v180 = v175
	goto L50
L48:
	;
	goto L49
L49:
	;
	F_FreeDir(m, v172)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L13
	} else {
		goto L94
	}
L50:
	;
	v188 = v180 + int32(19)
	v189 = F_strlen(m, v188)
	mBase = m.M
	if base.Ui32(v189) < base.Ui32(int32(25)) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L49
L52:
	;
	v341 = F_ReadDir(m, v172, int32(_a_F_do_pg_backup_stop_6))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L13
	} else {
		goto L92
	}
L53:
	;
	v192 = int32(_a_F_do_pg_backup_stop_7)
	v196 = m.G0
	v198 = v196 - int32(32)
	v199 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v198)+24)) = v199
	*(*int64)(unsafe.Add(mBase, uint32(v198)+16)) = v199
	*(*int64)(unsafe.Add(mBase, uint32(v198)+8)) = v199
	*(*int64)(unsafe.Add(mBase, uint32(v198))) = v199
	v207 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[7])))
	if v207 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v275 != int32(24) {
		goto L52
	} else {
		goto L73
	}
L55:
	;
	v275 = int32(0)
	goto L54
L56:
	;
	goto L57
L57:
	;
	v211 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[8])))
	if v211 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v215 = v188
	goto L61
L59:
	;
	goto L60
L60:
	;
	v225 = v192
	v226 = v207
	goto L64
L61:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v221 == v207 {
		v215 = v215 + int32(1)
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v275 = v215 - v188
	goto L54
L63:
	;
	goto L62
L64:
	;
	v233 = v198 + int32(base.Ui32(v226)>>(uint(int32(3))%32))&int32(28)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v235 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v234 | v235<<(uint(v226)%32)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v225)+1)))
	if v239 != 0 {
		v225 = v225 + v235
		v226 = v239
		goto L64
	} else {
		goto L66
	}
L65:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v242 == int32(0) {
		v265 = v188
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L65
L67:
	;
	v275 = v265 - v188
	goto L54
L68:
	;
	v246 = v188
	v247 = v242
	goto L69
L69:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v198+int32(base.Ui32(v247)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v255)>>(uint(v247)%32))&int32(1) == int32(0) {
		v265 = v246
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v265 = v263
	goto L67
L71:
	;
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	v263 = v246 + int32(1)
	if v261 != 0 {
		v246 = v263
		v247 = v261
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v280 = v188 + v189 - int32(7)
	v281 = int32(_a_F_do_pg_backup_stop_8)
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v280))))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[9])))
	if base.B2i32(v284 == int32(0))|base.B2i32(v284 != v287) != 0 {
		v305 = v284
		v306 = v287
		goto L75
	} else {
		goto L76
	}
L74:
	;
	if v305-v306 != 0 {
		goto L52
	} else {
		goto L81
	}
L75:
	;
	goto L74
L76:
	;
	v290 = v280
	v291 = v281
	goto L77
L77:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+1)))
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)))
	if v295 == int32(0) {
		v305 = v295
		v306 = v294
		goto L75
	} else {
		goto L79
	}
L78:
	;
	v305 = v295
	v306 = v294
	goto L75
L79:
	;
	v298 = int32(1)
	if v295 == v294 {
		v290 = v290 + v298
		v291 = v291 + v298
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v308 = F_XLogArchiveCheckDone(m, v188)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	if v308 == int32(0) {
		goto L52
	} else {
		goto L83
	}
L83:
	;
	v314 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	if v314 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v188
	F_errmsg_internal(m, int32(_a_F_do_pg_backup_stop_9), v13+int32(80))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L13
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = v188
	v329 = v13 + int32(1232)
	v334 = F_pg_snprintf(m, v329, int32(1031), int32(_a_F_do_pg_backup_stop_10), v13-int32(-64))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L13
	} else {
		goto L90
	}
L88:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_11), int32(_a_F_do_pg_backup_stop_12))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L13
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v336 = F_unlink(m, v329)
	mBase = m.M
	F_XLogArchiveCleanup(m, v188)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	goto L52
L92:
	;
	if v341 != 0 {
		v180 = v341
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
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[10]))
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
	v373 = base.B2i32(v368 == int32(2))
	goto L100
L99:
	;
	v373 = base.B2i32(int32(0) < v368)
	goto L100
L100:
	;
	if v373 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v376 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1088))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v377
	v382 = int64(*(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[6])))
	v383 = base.I64_div_u_s(v376-int64(1), v382)
	v385 = base.I64_div_u_s(int64(4294967296), v382)
	v386 = base.I64_div_u_s(v383, v385)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+36)) = uint32(v386)
	v389 = v383 - v385*v386
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+40)) = uint32(v389)
	v397 = F_pg_snprintf(m, v13+int32(1232), int32(64), int32(_a_F_do_pg_backup_stop_13), v13+int32(32))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
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
	v399 = *(*int64)(unsafe.Add(mBase, uint32(l0)+1032))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(l0)+1096))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v400
	v404 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = base.I32_wrap_i64(v399) & (v404 - int32(1))
	v409 = base.I64_extend_i32_s(v404)
	v410 = base.I64_div_u_s(v399, v409)
	v412 = base.I64_div_u_s(int64(4294967296), v409)
	v413 = base.I64_div_u_s(v410, v412)
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+20)) = uint32(v413)
	v416 = v410 - v412*v413
	*(*uint32)(unsafe.Add(mBase, uint32(v13)+24)) = uint32(v416)
	v424 = F_pg_snprintf(m, v13+int32(144), int32(64), int32(_a_F_do_pg_backup_stop_14), v13+int32(16))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L13
	} else {
		goto L105
	}
L105:
	;
	v427 = int32(0)
	v429 = v427
	v430 = int32(60)
	v432 = v427
	goto L106
L106:
	;
	v441 = F_XLogArchiveIsBusy(m, v13+int32(1232))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
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
	if v441 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v447 = F_XLogArchiveIsBusy(m, v13+int32(144))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L13
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v452 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[11]))
	if v452 != 0 {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	if v447 == int32(0) {
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
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if (base.B2i32(v432 < int32(6))|v429)&int32(1) != 0 {
		v476 = v429
		goto L119
	} else {
		goto L120
	}
L118:
	;
	goto L117
L119:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[12]))
	v482 = F_WaitLatch(m, v478, int32(41), int32(1000), int32(134217732))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L13
	} else {
		goto L125
	}
L120:
	;
	v460 = int32(1)
	v463 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L13
	} else {
		goto L121
	}
L121:
	;
	if v463 == int32(0) {
		v476 = v460
		goto L119
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_15), int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L13
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_16), int32(_a_F_do_pg_backup_stop_2))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L13
	} else {
		goto L124
	}
L124:
	;
	v476 = v460
	goto L119
L125:
	;
	v485 = *(*int32)(unsafe.Add(mBase, _c_F_do_pg_backup_stop[12]))
	*(*int32)(unsafe.Add(mBase, uint32(v485))) = int32(0)
	goto L126
L126:
	;
	v489 = v432 + int32(1)
	if v489 < v430 {
		v429 = v476
		v432 = v489
		goto L106
	} else {
		goto L127
	}
L127:
	;
	v492 = v430 << (uint(int32(1)) % 32)
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
		v429 = v476
		v430 = v492
		v432 = v489
		goto L106
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v489
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_17), v13)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L13
	} else {
		goto L130
	}
L130:
	;
	F_errhint(m, int32(_a_F_do_pg_backup_stop_18), int32(0))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L13
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_19), int32(_a_F_do_pg_backup_stop_2))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L13
	} else {
		goto L132
	}
L132:
	;
	v429 = v476
	v430 = v492
	v432 = v489
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
	v528 = int32(_a_F_do_pg_backup_stop_20)
	v538 = int32(_a_F_do_pg_backup_stop_21)
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
	v528 = int32(_a_F_do_pg_backup_stop_22)
	v538 = int32(_a_F_do_pg_backup_stop_23)
	goto L97
L137:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), v538, int32(_a_F_do_pg_backup_stop_2))
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
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_24), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	F_errhint(m, int32(_a_F_do_pg_backup_stop_25), int32(0))
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L13
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_26), int32(_a_F_do_pg_backup_stop_2))
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
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_27), int32(0))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L13
	} else {
		goto L146
	}
L146:
	;
	F_errhint(m, int32(_a_F_do_pg_backup_stop_28), int32(0))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L13
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_29), int32(_a_F_do_pg_backup_stop_2))
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
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_30), int32(0))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	F_errhint(m, int32(_a_F_do_pg_backup_stop_31), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_32), int32(_a_F_do_pg_backup_stop_2))
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
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_33), v13+int32(48))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L13
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_34), int32(_a_F_do_pg_backup_stop_2))
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
	F_errmsg(m, int32(_a_F_do_pg_backup_stop_35), v13+int32(96))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L13
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_do_pg_backup_stop_0), int32(_a_F_do_pg_backup_stop_36), int32(_a_F_do_pg_backup_stop_2))
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
func F_pg_backup_stop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
	var v76 int32
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
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v2)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_backup_stop[0])))
	v19 = F_get_call_result_type(m, l0, v2, v7+int32(28))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		if v19 == int32(1) {
			if v15 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v101 = m.ExcPending
					if v101 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_backup_stop_0), int32(0))
						mBase = m.M
						v105 = m.ExcPending
						if v105 != 0 {
							return int32(0)
						} else {
							F_errhint(m, int32(_a_F_pg_backup_stop_1), int32(0))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pg_backup_stop_2), int32(142), int32(_a_F_pg_backup_stop_3))
								mBase = m.M
								v114 = m.ExcPending
								if v114 != 0 {
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
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[1]))
				F_do_pg_backup_stop(m, v28, base.B2i32(v13 != int32(0)))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[1]))
					v36 = F_build_backup_content(m, v34, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v39 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[1]))
						v40 = *(*int64)(unsafe.Add(mBase, uint32(v39)+1088))
						v41 = F_Int64GetDatum(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v41
							v44 = F_cstring_to_text(m, v36)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v44
								v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[2]))
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
								v50 = F_cstring_to_text(m, v49)
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v50
									F_pfree(m, v36)
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										v56 = int32(0)
										*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[2])) = v56
										*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[1])) = v56
										v62 = *(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[3]))
										F_MemoryContextDelete(m, v62)
										mBase = m.M
										v64 = m.ExcPending
										if v64 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, _c_F_pg_backup_stop[3])) = int32(0)
											v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
											v73 = F_heap_form_tuple(m, v68, v7+int32(16), v7+int32(12))
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v73)+16))
												v76 = F_HeapTupleHeaderGetDatum(m, v75)
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													m.G0 = v7 + int32(32)
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
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_backup_stop_4), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_backup_stop_2), int32(136), int32(_a_F_pg_backup_stop_3))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
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
func F_pg_basetype(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_SearchSysCache1(m, int32(82), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_ReleaseCatCache(m, v47)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L13
	}
L2:
	;
	v41 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v41)
	return int32(0)
L3:
	;
	return int32(0)
L4:
	;
	if v7 == int32(0) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+22)))
	v15 = v13 + v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+79)))
	if v16 != int32(100) {
		v47 = v7
		v48 = v6
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v20 = v15
	v21 = v7
	goto L7
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)+132))
	F_ReleaseCatCache(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L9
	}
L8:
	;
	v47 = v27
	v48 = v23
	goto L1
L9:
	;
	v27 = F_SearchSysCache1(m, int32(82), v23)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	if v27 == int32(0) {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+22)))
	v33 = v31 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+79)))
	if v34 == int32(100) {
		v20 = v33
		v21 = v27
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
L13:
	;
	return v48
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
func F_pg_buffercache_numa_pages(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
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
	var v30 int64
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v59 int64
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v14 != 0 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
		if base.Ui64(v17) < base.Ui64(v18) {
			v22 = base.I32_wrap_i64(v17) * int32(24)
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
			v25 = v22 + v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			v27 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v27)
			*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v26
			v30 = *(*int64)(unsafe.Add(mBase, uint32(v25)+8))
			v31 = F_Int64GetDatum(m, v30)
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v35)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v31
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)+16))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v38+v22)+16))
				v42 = int32(base.Ui32(v40) >> (uint(int32(31)) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+31)) = uint8(v42)
				if v35 < v40 {
					v47 = v40
				} else {
					v47 = v35
				}
				*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v54 = F_heap_form_tuple(m, v49, v11+int32(32), v11+int32(29))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
					v57 = F_HeapTupleHeaderGetDatum(m, v56)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						v59 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
						*(*int64)(unsafe.Add(mBase, uint32(v16))) = v59 + int64(1)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v63)+20)) = int32(1)
						v74 = v57
						m.G0 = v11 + int32(48)
						return v74
					}
				}
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v68)+20)) = int32(2)
				v71 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v71)
				v74 = int32(0)
				m.G0 = v11 + int32(48)
				return v74
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v84 = m.ExcPending
		if v84 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_pg_buffercache_numa_pages_0), int32(0))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_pg_buffercache_numa_pages_1), int32(329), int32(_a_F_pg_buffercache_numa_pages_2))
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
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
func F_pg_buffercache_pages(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int64
	_ = v205
	var v206 int64
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int64
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L53
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L5
	} else {
		goto L50
	}
L3:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+16))
	goto L36
L4:
	;
	v16 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int32(0)
L6:
	;
	v20 = int32(_a_F_pg_buffercache_pages_0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[0])) = v23
	v26 = F_palloc(m, int32(8))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v31 = F_get_call_result_type(m, l0, int32(0), v10+int32(-4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	if v31 != int32(1) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if base.Ui32(v36-int32(10)) <= base.Ui32(int32(-3)) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v41 = F_CreateTemplateTupleDesc(m, v36)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	F_TupleDescInitEntry(m, v41, int32(1), int32(_a_F_pg_buffercache_pages_1), int32(23), int32(-1), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	F_TupleDescInitEntry(m, v41, int32(2), int32(_a_F_pg_buffercache_pages_2), int32(26), int32(-1), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	F_TupleDescInitEntry(m, v41, int32(3), int32(_a_F_pg_buffercache_pages_3), int32(26), int32(-1), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	F_TupleDescInitEntry(m, v41, int32(4), int32(_a_F_pg_buffercache_pages_4), int32(26), int32(-1), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	F_TupleDescInitEntry(m, v41, int32(5), int32(_a_F_pg_buffercache_pages_5), int32(21), int32(-1), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	F_TupleDescInitEntry(m, v41, int32(6), int32(_a_F_pg_buffercache_pages_6), int32(20), int32(-1), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_TupleDescInitEntry(m, v41, int32(7), int32(_a_F_pg_buffercache_pages_7), int32(16), int32(-1), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	F_TupleDescInitEntry(m, v41, int32(8), int32(_a_F_pg_buffercache_pages_8), int32(21), int32(-1), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v12)+60))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v100 == int32(9) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_TupleDescInitEntry(m, v41, int32(9), int32(_a_F_pg_buffercache_pages_9), int32(23), int32(-1), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v110 = F_BlessTupleDesc(m, v41)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v110
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[0]))
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[1]))
	v119 = F_MemoryContextAllocHuge(m, v114, v116<<(uint(int32(5))%32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v119
	v123 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v123
	*(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[0])) = v21
	if v123 <= int64(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v136 = int32(0)
	goto L27
L27:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[2]))
	if v140 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L3
L29:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[3]))
	v147 = v144 + v136<<(uint(int32(6))%32)
	v148 = F_LockBufHdr(m, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v153 = v150 + v136<<(uint(int32(5))%32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v147)+20))
	v155 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v154 + v155
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v147)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v147)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+28)) = v148 & int32(_a_F_pg_buffercache_pages_10)
	v173 = int32(base.Ui32(v148)>>(uint(int32(18))%32)) & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v153)+26)) = uint16(v173)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = v166
	v179 = int32(base.Ui32(v148)>>(uint(int32(23))%32)) & v155
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+25)) = uint8(v179)
	v181 = int32(50331648)
	*(*uint8)(unsafe.Add(mBase, uint32(v153)+24)) = uint8(base.B2i32(v148&v181 == v181))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = v148 & int32(-4194305)
	v190 = v136 + v155
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_pages[1]))
	if v190 < v192 {
		v136 = v190
		goto L27
	} else {
		goto L34
	}
L34:
	;
	goto L28
L35:
	;
	m.G0 = v12 - int32(-64)
	return v290
L36:
	;
	v205 = *(*int64)(unsafe.Add(mBase, uint32(v204)))
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v204)+8))
	if base.Ui64(v205) < base.Ui64(v206) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204)+16))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v210 = base.I32_wrap_i64(v205)
	v213 = v209 + v210<<(uint(int32(5))%32)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v213)))
	v215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)) = uint8(v215)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v214
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v213)+20))
	if v218 != int32(-1) {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	goto L39
L39:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L5
	} else {
		goto L49
	}
L40:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	v270 = F_heap_form_tuple(m, v265, v10+int32(-48), v10+int32(-57))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L5
	} else {
		goto L47
	}
L41:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v213)+4))
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)) = uint8(v225)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v224
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v213)+8))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)) = uint8(v225)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v228
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v213)+12))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+10)) = uint8(v225)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v232
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v213)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)) = uint8(v225)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v236
	v241 = F_Int64GetDatum(m, base.I64_extend_i32_u(v218))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L46
	}
L42:
	;
	v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+24)))
	if v221 != 0 {
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(72340172838076673)
	goto L40
L45:
	;
	goto L44
L46:
	;
	v243 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v241
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v208)+4))
	v249 = v246 + v210<<(uint(int32(5))%32)
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+25)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v250
	v254 = int32(*(*int16)(unsafe.Add(mBase, uint32(v249)+26)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v254
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v249)+28))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)) = uint8(v243)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v258
	goto L40
L47:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v270)+16))
	v273 = F_HeapTupleHeaderGetDatum(m, v272)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v275 = *(*int64)(unsafe.Add(mBase, uint32(v204)))
	*(*int64)(unsafe.Add(mBase, uint32(v204))) = v275 + int64(1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v279)+20)) = int32(1)
	v290 = v273
	goto L35
L49:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+20)) = int32(2)
	v287 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v287)
	v290 = int32(0)
	goto L35
L50:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_pages_11), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_pages_12), int32(141), int32(_a_F_pg_buffercache_pages_13))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_pg_buffercache_pages_14), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_pg_buffercache_pages_12), int32(145), int32(_a_F_pg_buffercache_pages_13))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_buffercache_usage_counts(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
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
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
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
	var v157 int32
	_ = v157
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(128)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+112)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v23)+104)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v23)+80)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v23)+72)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v23)+64)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v23)+48)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v23)+40)) = v26
	*(*int64)(unsafe.Add(mBase, uint32(v23)+32)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v2
	F_InitMaterializedSRF(m, l0, v2)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[0]))
	if v51 < v53 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[1]))
	v63 = v57
	v65 = v2
	goto L6
L4:
	;
	v145 = v51
	v148 = v2
	v149 = v2
	v150 = v2
	v151 = v2
	v152 = v2
	v153 = v2
	v154 = v2
	v155 = v2
	v156 = v2
	v157 = v2
	v158 = v2
	v159 = v2
	v160 = v2
	v161 = v2
	v162 = v2
	v163 = v2
	v164 = v2
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(0)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	v173 = v23 + int32(16)
	v175 = v23 + int32(12)
	F_tuplestore_putvalues(m, v170, v171, v173, v175)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L19
	}
L6:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v63+v65<<(uint(int32(6))%32))+24))
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[2]))
	if v83 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v23)+84))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v23)+116))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v23)+80))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v23)+112))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v23)+44))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v23)+108))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v23)+72))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v23)+104))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v23)+36))
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v23)+100))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v23)+64))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v23)+96))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
	v145 = v144
	v148 = v141
	v149 = v142
	v150 = v130
	v151 = v129
	v152 = v127
	v153 = v128
	v154 = v131
	v155 = v132
	v156 = v133
	v157 = v134
	v158 = v135
	v159 = v136
	v160 = v137
	v161 = v138
	v162 = v139
	v163 = v140
	v164 = v143
	goto L5
L8:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	v88 = v63
	goto L10
L10:
	;
	v94 = int32(base.Ui32(v81)>>(uint(int32(18))%32)) & int32(15) << (uint(int32(2)) % 32)
	v97 = v94 + (v23 + int32(96))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v98 + int32(1)
	if v81&int32(_a_F_pg_buffercache_usage_counts_0) != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[1]))
	v88 = v87
	goto L10
L12:
	;
	v106 = v23 - int32(-64) + v94
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v107 + int32(1)
	goto L14
L13:
	;
	goto L14
L14:
	;
	if v81&int32(_a_F_pg_buffercache_usage_counts_1) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v116 = v23 + int32(32) + v94
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v117 + int32(1)
	goto L17
L16:
	;
	goto L17
L17:
	;
	v123 = v65 + int32(1)
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_pg_buffercache_usage_counts[0]))
	if v123 < v125 {
		v63 = v88
		v65 = v123
		goto L6
	} else {
		goto L18
	}
L18:
	;
	goto L7
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v163
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(1)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v183, v184, v173, v175)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v159
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(2)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v192, v193, v173, v175)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v156
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(3)
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v201, v202, v173, v175)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v151
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(4)
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v210, v211, v173, v175)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v152
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v153
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = int32(5)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	F_tuplestore_putvalues(m, v219, v220, v173, v175)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	m.G0 = v23 + int32(128)
	return int32(0)
}
func F_pg_cancel_backend(m *base.Module, l0 int32) int32 {
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
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_signal_backend(m, v7, int32(2))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		switch v9 - int32(2) {
		case 0:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_cancel_backend_0), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = int32(_a_F_pg_cancel_backend_1)
						F_errdetail(m, int32(_a_F_pg_cancel_backend_2), v5+int32(32))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_cancel_backend_3), int32(159), int32(_a_F_pg_cancel_backend_4))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
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
		case 1:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_cancel_backend_0), int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = int32(_a_F_pg_cancel_backend_5)
						*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v26
						*(*int32)(unsafe.Add(mBase, uint32(v5))) = v26
						F_errdetail(m, int32(_a_F_pg_cancel_backend_6), v5)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_cancel_backend_3), int32(145), int32(_a_F_pg_cancel_backend_4))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
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
		case 2:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_cancel_backend_0), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(_a_F_pg_cancel_backend_7)
						F_errdetail(m, int32(_a_F_pg_cancel_backend_8), v5+int32(16))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_cancel_backend_3), int32(152), int32(_a_F_pg_cancel_backend_4))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
		default:
			m.G0 = v5 + int32(48)
			return base.B2i32(v9 == int32(0))
		}
	}
}
func F_pg_check_frozen(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13961(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
func F_pg_client_encoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_pg_client_encoding[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v7 = F_DirectFunctionCall1Coll(m, int32(500), int32(0), v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_control_system(m *base.Module, l0 int32) int32 {
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
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int64
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v11 = F_get_call_result_type(m, l0, int32(0), v6+int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_system[0]))
			v22 = F_LWLockAcquire(m, v18+int32(1152), int32(1))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_system[1]))
				v28 = F_get_controlfile(m, v25, v6+int32(7))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_control_system[0]))
					F_LWLockRelease(m, v31+int32(1152))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+7)))
						if v36 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_pg_control_system_0), int32(0))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_pg_control_system_1), int32(50), int32(_a_F_pg_control_system_2))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
							v40 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v39
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v40)
							*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v43
							v47 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
							v48 = F_Int64GetDatum(m, v47)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								v50 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v50)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v48
								v53 = *(*int64)(unsafe.Add(mBase, uint32(v28)+24))
								v58 = F_Int64GetDatum(m, v53*int64(1000000)-int64(946684800000000))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v60 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)) = uint8(v60)
									*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v58
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
									v68 = F_heap_form_tuple(m, v63, v6+int32(16), v6+int32(12))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return int32(0)
									} else {
										v70 = *(*int32)(unsafe.Add(mBase, uint32(v68)+16))
										v71 = F_HeapTupleHeaderGetDatum(m, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											m.G0 = v6 + int32(32)
											return v71
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
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_pg_control_system_3), int32(0))
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_control_system_1), int32(42), int32(_a_F_pg_control_system_2))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
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
func F_pg_convert_to(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
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
	var v17 int32
	_ = v17
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_convert_to[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v12 = F_DirectFunctionCall1Coll(m, int32(500), v2, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = F_DirectFunctionCall3Coll(m, int32(1636), v2, v6, v12, v3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
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
func F_pg_ddl_command_out(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_pg_ddl_command_out_0), int32(358), int32(_a_F_pg_ddl_command_out_1), int32(_a_F_pg_ddl_command_out_2), int32(_a_F_pg_ddl_command_out_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_dependencies_in(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_pg_dependencies_in_0), int32(661), int32(_a_F_pg_dependencies_in_1), int32(_a_F_pg_dependencies_in_2), int32(_a_F_pg_dependencies_in_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_pg_describe_object(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v9|v10 == int32(0) {
		v14 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
		v36 = int32(0)
		m.G0 = v7 + int32(16)
		return v36
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v17
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v9
		v24 = F_getObjectDescription(m, v7+int32(4), int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			if v24 == int32(0) {
				v30 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v30)
				v36 = int32(0)
				m.G0 = v7 + int32(16)
				return v36
			} else {
				v33 = F_cstring_to_text(m, v24)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v36 = v33
					m.G0 = v7 + int32(16)
					return v36
				}
			}
		}
	}
}
func F_pg_encrypt_iv(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
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
	return int32(0)
L3:
	;
	v21 = int32(1)
	v22 = v17 + v21
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v27 = v25 & v21
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v22
	goto L6
L5:
	;
	v28 = v17 + int32(4)
	goto L6
L6:
	;
	if v25 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = F_downcase_truncate_identifier(m, v28, v55, int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L2
	} else {
		goto L18
	}
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v34 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v45 = int32(1)
	if v27 != 0 {
		v55 = int32(base.Ui32(v25)>>(uint(v45)%32)) - v45
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v37 = int32(16)
	goto L13
L12:
	;
	v37 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v44 = int32(4)
	goto L16
L15:
	;
	v44 = v37
	goto L16
L16:
	;
	v55 = v44
	goto L7
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v61 = F_px_find_combo(m, v57, v14+int32(28))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	if v61 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	F_pfree(m, v57)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
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
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v69 = F_pg_detoast_datum_packed(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v72 = F_pg_detoast_datum_packed(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v75 = F_pg_detoast_datum_packed(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v77 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v107 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+1)))
	if v83 == int32(18) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L30
L30:
	;
	v94 = int32(1)
	if v77&v94 != 0 {
		v106 = int32(base.Ui32(v77)>>(uint(v94)%32)) - v94
		goto L27
	} else {
		goto L37
	}
L31:
	;
	v86 = int32(16)
	goto L33
L32:
	;
	v86 = int32(0)
	goto L33
L33:
	;
	if base.Ui32((v83-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v93 = int32(4)
	goto L36
L35:
	;
	v93 = v86
	goto L36
L36:
	;
	v106 = v93
	goto L27
L37:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v106 = int32(base.Ui32(v100)>>(uint(int32(2))%32)) - int32(4)
	goto L27
L38:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v137 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L39:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v113 == int32(18) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	v124 = int32(1)
	if v107&v124 != 0 {
		v136 = int32(base.Ui32(v107)>>(uint(v124)%32)) - v124
		goto L38
	} else {
		goto L48
	}
L42:
	;
	v116 = int32(16)
	goto L44
L43:
	;
	v116 = int32(0)
	goto L44
L44:
	;
	if base.Ui32((v113-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v123 = int32(4)
	goto L47
L46:
	;
	v123 = v116
	goto L47
L47:
	;
	v136 = v123
	goto L38
L48:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v136 = int32(base.Ui32(v130)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L49:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v67)+12))
	v168 = m.T0[v167].(func(*base.Module, int32, int32) int32)(m, v67, v106)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L2
	} else {
		goto L60
	}
L50:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+1)))
	if v143 == int32(18) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	v154 = int32(1)
	if v137&v154 != 0 {
		v166 = int32(base.Ui32(v137)>>(uint(v154)%32)) - v154
		goto L49
	} else {
		goto L59
	}
L53:
	;
	v146 = int32(16)
	goto L55
L54:
	;
	v146 = int32(0)
	goto L55
L55:
	;
	if base.Ui32((v143-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v153 = int32(4)
	goto L58
L57:
	;
	v153 = v146
	goto L58
L58:
	;
	v166 = v153
	goto L49
L59:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v166 = int32(base.Ui32(v160)>>(uint(int32(2))%32)) - int32(4)
	goto L49
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v168
	v173 = F_palloc(m, v168+int32(4))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	v175 = int32(1)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v177&v175 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v180 = v175
	goto L64
L63:
	;
	v180 = int32(4)
	goto L64
L64:
	;
	v182 = int32(1)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75))))
	if v184&v182 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v187 = v182
	goto L67
L66:
	;
	v187 = int32(4)
	goto L67
L67:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
	v190 = m.T0[v189].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v67, v72+v180, v136, v75+v187, v166)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	if v190 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	m.T0[v192].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v195 = int32(1)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v197&v195 != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v287 = v190
	goto L1
L73:
	;
	v200 = v195
	goto L75
L74:
	;
	v200 = int32(4)
	goto L75
L75:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	v207 = m.T0[v206].(func(*base.Module, int32, int32, int32, int32, int32) int32)(m, v67, v69+v200, v106, v173+int32(4), v14+int32(28))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L76
	}
L76:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	m.T0[v209].(func(*base.Module, int32))(m, v67)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	if v207 != 0 {
		v287 = v207
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v212<<(uint(int32(2))%32) + int32(16)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v218 != v69 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	F_pfree(m, v69)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v222 != v72 {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	goto L81
L83:
	;
	F_pfree(m, v72)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v226 != v75 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	goto L85
L87:
	;
	F_pfree(m, v75)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	if v230 != v17 {
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
	v233 = m.ExcPending
	if v233 != 0 {
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
	return v173
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
	if v61 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v274
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57
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
	if v61 != v256 {
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
	if v61 != v262 {
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
func F_pg_filenode_relation(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v3 != 0 {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v5 = F_RelidByRelfilenumber(m, v4, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			if v5 != 0 {
				v13 = v5
			} else {
				v10 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
				v13 = int32(0)
			}
			return v13
		}
	} else {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
		v13 = int32(0)
		return v13
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
func F_pg_get_object_address(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v440 int32
	_ = v440
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
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
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
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(272)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_text_to_cstring(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = F_pg_detoast_datum(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v30 = v2
	goto L15
L5:
	;
	switch v76 - int32(2) {
	case 0, 1:
		goto L96
	default:
		goto L94
	case 3, 9, 11, 30, 41:
		goto L98
	case 22, 24:
		goto L97
	case 23:
		goto L95
	case 29, 48:
		goto L99
	}
L6:
	;
	v325 = F_textarray_to_strvaluelist(m, v23)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L88
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L1
	} else {
		goto L84
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L80
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L76
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L72
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L68
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L64
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L60
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L56
	}
L15:
	;
	v35 = v30 << (uint(int32(3)) % 32)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_pg_get_object_address[0])))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if base.B2i32(v39 == int32(0))|base.B2i32(v39 != v42) != 0 {
		v60 = v39
		v61 = v42
		goto L18
	} else {
		goto L19
	}
L16:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v30))%64)&int64(4398046543432) != int64(0) {
		goto L13
	} else {
		goto L28
	}
L17:
	;
	if v60-v61 != 0 {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	v45 = v36
	v46 = v15
	goto L20
L20:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+1)))
	if v50 == int32(0) {
		v60 = v50
		v61 = v49
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v60 = v50
	v61 = v49
	goto L18
L22:
	;
	v53 = int32(1)
	if v50 == v49 {
		v45 = v45 + v53
		v46 = v46 + v53
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v64 = v30 + int32(1)
	if v64 != int32(59) {
		v30 = v64
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
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_pg_get_object_address[1])))
	switch v76 - int32(5) {
	case 0, 7, 8, 38, 44:
		goto L32
	default:
		goto L30
	case 17:
		goto L31
	}
L29:
	;
	if v30&int32(2147483646) == int32(32) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v124 = F_textarray_to_strvaluelist(m, v20)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L43
	}
L31:
	;
	F_deconstruct_array_builtin(m, v20, int32(25), v12+int32(256), v12+int32(268), v12+int32(248))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L38
	}
L32:
	;
	F_deconstruct_array_builtin(m, v20, int32(25), v12+int32(256), v12+int32(268), v12+int32(248))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+248))
	if v88 != int32(1) {
		goto L12
	} else {
		goto L34
	}
L34:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v12)+268))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v92 == int32(1) {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v12)+256))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v97 = F_text_to_cstring(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v100 = F_typeStringToTypeName(m, v97, int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v128 = v2
	v129 = v100
	goto L29
L38:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+248))
	if v111 != int32(1) {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v12)+268))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v115 == int32(1) {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v12)+256))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)))
	v120 = F_text_to_cstring(m, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v122 = F_makeFloat(m, v120)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v322 = v122
	v323 = v2
	v324 = v2
	goto L6
L43:
	;
	if v124 == int32(0) {
		goto L8
	} else {
		goto L44
	}
L44:
	;
	v128 = v124
	v129 = v2
	goto L29
L45:
	;
	F_deconstruct_array_builtin(m, v23, int32(25), v12+int32(256), v12+int32(268), v12+int32(248))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L47
	}
L46:
	;
	switch v76 - int32(1) {
	case 0, 4, 18, 24, 28, 33:
		goto L45
	default:
		v322 = int32(0)
		v323 = v128
		v324 = v129
		goto L6
	}
L47:
	;
	v147 = int32(0)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+248))
	if v149 <= v147 {
		v329 = v147
		v330 = v128
		v331 = v147
		v334 = v129
		goto L5
	} else {
		goto L48
	}
L48:
	;
	v157 = v147
	v158 = int32(0)
	goto L49
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+268))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162+v158))))
	if v164 == int32(1) {
		goto L7
	} else {
		goto L51
	}
L50:
	;
	v329 = v147
	v330 = v128
	v331 = v177
	v334 = v129
	goto L5
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v12)+256))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v158<<(uint(int32(2))%32))))
	v172 = F_text_to_cstring(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v175 = F_typeStringToTypeName(m, v172, int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v177 = F_lappend(m, v157, v175)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v180 = v158 + int32(1)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+248))
	if v180 < v181 {
		v157 = v177
		v158 = v180
		goto L49
	} else {
		goto L55
	}
L55:
	;
	goto L50
L56:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+208)) = v15
	F_errmsg(m, int32(_a_F_pg_get_object_address_0), v12+int32(208))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2620), int32(_a_F_pg_get_object_address_2))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+192)) = v15
	F_errmsg(m, int32(_a_F_pg_get_object_address_3), v12+int32(192))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2132), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_5), v12+int32(160))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2151), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errmsg(m, int32(_a_F_pg_get_object_address_6), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2155), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+176)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_5), v12+int32(176))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2168), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	F_errmsg(m, int32(_a_F_pg_get_object_address_7), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2172), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_8), v12)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2181), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_pg_get_object_address_6), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2210), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v329 = v322
	v330 = v323
	v331 = v325
	v334 = v324
	goto L5
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v604 = m.ExcPending
	if v604 != 0 {
		goto L1
	} else {
		goto L170
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L166
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L162
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L1
	} else {
		goto L158
	}
L93:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L154
	}
L94:
	;
	switch v76 {
	case 0, 9, 14, 15, 16, 17, 21, 27, 30, 33, 36, 38, 42:
		goto L132
	case 1, 19, 25, 29, 34:
		goto L123
	case 2, 3:
		goto L126
	case 4, 6, 7, 8, 10, 18, 20, 23, 24, 26, 28, 35, 37, 39, 40, 41, 44, 45, 46, 47, 48, 51:
		goto L125
	case 5, 13, 43:
		goto L130
	case 11:
		goto L127
	case 12, 49:
		goto L131
	default:
		v456 = v329
		goto L124
	case 31, 50:
		goto L128
	case 32:
		goto L129
	}
L95:
	;
	if v331 == int32(0) {
		goto L91
	} else {
		goto L120
	}
L96:
	;
	if v330 == int32(0) {
		goto L92
	} else {
		goto L118
	}
L97:
	;
	if v330 != 0 {
		goto L110
	} else {
		goto L111
	}
L98:
	;
	if v331 != 0 {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	if v330 == int32(0) {
		goto L93
	} else {
		goto L100
	}
L100:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if v340 != int32(1) {
		goto L93
	} else {
		goto L101
	}
L101:
	;
	goto L98
L102:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v343 == int32(1) {
		goto L94
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_9), v12+int32(96))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2244), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if int32(1) < v365 {
		goto L94
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = int32(2)
	F_errmsg(m, int32(_a_F_pg_get_object_address_8), v12+int32(112))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2251), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if v389 <= int32(2) {
		goto L92
	} else {
		goto L119
	}
L119:
	;
	goto L95
L120:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v331)+4))
	if v394 != int32(2) {
		goto L91
	} else {
		goto L121
	}
L121:
	;
	goto L94
L122:
	;
	F_get_object_address(m, v12+int32(256), v76, v480, v12+int32(248), int32(1), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L145
	}
L123:
	;
	v474 = F_palloc0(m, int32(20))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L144
	}
L124:
	;
	if v456 != 0 {
		v480 = v456
		goto L122
	} else {
		goto L140
	}
L125:
	;
	v456 = v330
	goto L124
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+216)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(v12)+220)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v331
	v454 = F_list_make2_impl(m, v12+int32(76), v12+int32(72))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L1
	} else {
		goto L139
	}
L127:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)))
	v444 = F_lcons(m, v443, v330)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L138
	}
L128:
	;
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v428)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v429
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v431)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v432
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v432
	v440 = F_list_make2_impl(m, v12+int32(68), v12-int32(-64))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L137
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v330
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v417)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v418
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v418
	v426 = F_list_make2_impl(m, v12+int32(60), v12+int32(56))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L136
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v334
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v331)+12))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v406
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v406
	v414 = F_list_make2_impl(m, v12+int32(52), v12+int32(48))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L135
	}
L131:
	;
	v456 = v334
	goto L124
L132:
	;
	if v330 == int32(0) {
		goto L90
	} else {
		goto L133
	}
L133:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	if v399 != int32(1) {
		goto L90
	} else {
		goto L134
	}
L134:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v330)+12))
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v402)))
	v456 = v403
	goto L124
L135:
	;
	v456 = v414
	goto L124
L136:
	;
	v456 = v426
	goto L124
L137:
	;
	v456 = v440
	goto L124
L138:
	;
	v456 = v444
	goto L124
L139:
	;
	v456 = v454
	goto L124
L140:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v76
	F_errmsg_internal(m, int32(_a_F_pg_get_object_address_10), v12+int32(16))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2363), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
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
	*(*int32)(unsafe.Add(mBase, uint32(v474)+8)) = v331
	*(*int32)(unsafe.Add(mBase, uint32(v474)+4)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = int32(153)
	v480 = v474
	goto L122
L145:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v12)+264))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v12)+260))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v12)+256))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v12)+248))
	if v493 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	F_relation_close(m, v493, int32(1))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v500 = F_get_call_result_type(m, l0, int32(0), v12+int32(268))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L150
	}
L149:
	;
	goto L148
L150:
	;
	if v500 != int32(1) {
		goto L89
	} else {
		goto L151
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v491
	*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v492
	v507 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+255)) = uint8(v507)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+253)) = uint16(v507)
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v12)+268))
	v516 = F_heap_form_tuple(m, v511, v12+int32(256), v12+int32(253))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v516)+16))
	v519 = F_HeapTupleHeaderGetDatum(m, v518)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	m.G0 = v12 + int32(272)
	return v519
L154:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_5), v12+int32(80))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2233), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = int32(3)
	F_errmsg(m, int32(_a_F_pg_get_object_address_8), v12+int32(128))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2258), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L162:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = int32(2)
	F_errmsg(m, int32(_a_F_pg_get_object_address_9), v12+int32(144))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2265), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(1)
	F_errmsg(m, int32(_a_F_pg_get_object_address_5), v12+int32(32))
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2317), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_pg_get_object_address_11), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L1
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_pg_get_object_address_1), int32(2373), int32(_a_F_pg_get_object_address_4))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_replica_identity_index(m *base.Module, l0 int32) int32 {
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
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_table_open(m, v4, int32(1))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_RelationGetReplicaIndex(m, v6)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_relation_close(m, v6, int32(1))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v10 == int32(0) {
					v17 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v17)
				} else {
				}
				return v10
			}
		}
	}
}
func F_pg_get_triggerdef_worker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v208 int32
	_ = v208
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
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
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v471 int32
	_ = v471
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
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
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v584 int32
	_ = v584
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	v14 = m.G0
	v16 = v14 - int32(320)
	m.G0 = v16
	v20 = F_table_open(m, int32(2620), int32(1))
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
	v25 = v16 + int32(256)
	F_ScanKeyInit(m, v25, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = int32(1)
	v35 = F_systable_beginscan(m, v20, int32(2702), v32, int32(0), v32, v25)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L162
	}
L5:
	;
	m.G0 = v16 + int32(320)
	return v627
L6:
	;
	v37 = F_systable_getnext(m, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	if v37 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v37)+16))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+22)))
	v49 = v16 + int32(304)
	F_initStringInfo(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	F_relation_close(m, v20, int32(1))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v627 = int32(0)
	goto L5
L13:
	;
	v52 = v46 + v47
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+92))
	v56 = F_quote_identifier(m, v52+int32(12))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = v56
	if v53 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v61 = int32(_a_F_pg_get_triggerdef_worker_0)
	goto L17
L16:
	;
	v61 = int32(_a_F_pg_get_triggerdef_worker_1)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = v61
	F_appendStringInfo(m, v49, int32(_a_F_pg_get_triggerdef_worker_2), v16+int32(112))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+80)))
	v71 = v69 & int32(66)
	switch v71 {
	case 0:
		goto L20
	case 1:
		goto L22
	case 2:
		v92 = int32(_a_F_pg_get_triggerdef_worker_3)
		goto L19
	default:
		goto L23
	}
L19:
	;
	v94 = v16 + int32(304)
	F_appendStringInfoString(m, v94, v92)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L28
	}
L20:
	;
	v92 = int32(_a_F_pg_get_triggerdef_worker_4)
	goto L19
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L25
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	switch v71 - int32(65) {
	case 0:
		goto L22
	case 1:
		goto L21
	default:
		goto L24
	}
L24:
	;
	v92 = int32(_a_F_pg_get_triggerdef_worker_5)
	goto L19
L25:
	;
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+80)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v79
	F_errmsg_internal(m, int32(_a_F_pg_get_triggerdef_worker_6), v16+int32(96))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_pg_get_triggerdef_worker_7), int32(957), int32(_a_F_pg_get_triggerdef_worker_8))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
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
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+80)))
	if v97&int32(4) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L29:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if l1 != 0 {
		goto L65
	} else {
		goto L66
	}
L30:
	;
	F_appendStringInfoString(m, v94, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L63
	}
L31:
	;
	if v97&int32(32) == int32(0) {
		goto L29
	} else {
		goto L62
	}
L32:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+80)))
	if v208&int32(32) == int32(0) {
		goto L29
	} else {
		goto L61
	}
L33:
	;
	F_appendStringInfoString(m, v94, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L46
	}
L34:
	;
	if v97&int32(16) == int32(0) {
		goto L31
	} else {
		goto L45
	}
L35:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+80)))
	if v121&int32(16) == int32(0) {
		goto L32
	} else {
		goto L44
	}
L36:
	;
	F_appendStringInfoString(m, v94, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L43
	}
L37:
	;
	if v97&int32(8) == int32(0) {
		goto L34
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_appendStringInfoString(m, v16+int32(304), int32(_a_F_pg_get_triggerdef_worker_9))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	v118 = int32(_a_F_pg_get_triggerdef_worker_10)
	goto L36
L41:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+80)))
	if v112&int32(8) == int32(0) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	v118 = int32(_a_F_pg_get_triggerdef_worker_11)
	goto L36
L43:
	;
	goto L35
L44:
	;
	v132 = int32(_a_F_pg_get_triggerdef_worker_12)
	goto L33
L45:
	;
	v132 = int32(_a_F_pg_get_triggerdef_worker_13)
	goto L33
L46:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v52)+116))
	if v135 <= int32(0) {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	v139 = v16 + int32(304)
	F_appendStringInfoString(m, v139, int32(_a_F_pg_get_triggerdef_worker_14))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v52)+116))
	if v143 <= int32(0) {
		goto L32
	} else {
		goto L49
	}
L49:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v147 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+124)))
	v149 = F_get_attname(m, v146, v147, int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v151 = F_quote_identifier(m, v149)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_appendStringInfoString(m, v139, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v52)+116))
	if v155 < int32(2) {
		goto L32
	} else {
		goto L53
	}
L53:
	;
	v164 = int32(1)
	goto L54
L54:
	;
	v175 = v16 + int32(304)
	F_appendStringInfoString(m, v175, int32(_a_F_pg_get_triggerdef_worker_15))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	goto L32
L56:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v183 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52+int32(124)+v164<<(uint(int32(1))%32)))))
	v185 = F_get_attname(m, v179, v183, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v187 = F_quote_identifier(m, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_appendStringInfoString(m, v175, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v192 = v164 + int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v52)+116))
	if v192 < v193 {
		v164 = v192
		goto L54
	} else {
		goto L60
	}
L60:
	;
	goto L55
L61:
	;
	v232 = int32(_a_F_pg_get_triggerdef_worker_16)
	goto L30
L62:
	;
	v232 = int32(_a_F_pg_get_triggerdef_worker_17)
	goto L30
L63:
	;
	goto L29
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v254
	v257 = v16 + int32(304)
	F_appendStringInfo(m, v257, int32(_a_F_pg_get_triggerdef_worker_18), v16+int32(80))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L70
	}
L65:
	;
	v250 = F_generate_relation_name(m, v248, int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v252 = F_generate_qualified_relation_name(m, v248)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L69
	}
L68:
	;
	v254 = v250
	goto L64
L69:
	;
	v254 = v252
	goto L64
L70:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v52)+92))
	if v263 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v52)+84))
	if v264 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v297 = v16 + int32(255)
	v298 = F_fastgetattr_3(m, v37, int32(18), v295, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L88
	}
L74:
	;
	v266 = F_generate_relation_name(m, v264, int32(0))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+96)))
	if v274 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v266
	F_appendStringInfo(m, v257, int32(_a_F_pg_get_triggerdef_worker_19), v16-int32(-64))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	F_appendStringInfoString(m, v16+int32(304), int32(_a_F_pg_get_triggerdef_worker_20))
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v283 = v16 + int32(304)
	F_appendStringInfoString(m, v283, int32(_a_F_pg_get_triggerdef_worker_21))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+97)))
	if v289 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v290 = int32(_a_F_pg_get_triggerdef_worker_22)
	goto L86
L85:
	;
	v290 = int32(_a_F_pg_get_triggerdef_worker_23)
	goto L86
L86:
	;
	F_appendStringInfoString(m, v283, v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	goto L73
L88:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)))
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v303 = F_fastgetattr_3(m, v37, int32(19), v302, v297)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v305 = int32(-1)
	v307 = int32(0)
	v309 = (v300 ^ v305) & base.B2i32(v298 != v307)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)))
	v315 = (v310 ^ v305) & base.B2i32(v303 != v307)
	if v309|v315 == v307 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v346 = v16 + int32(304)
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+80)))
	if v349&int32(1) != 0 {
		goto L101
	} else {
		goto L102
	}
L91:
	;
	v320 = v16 + int32(304)
	F_appendStringInfoString(m, v320, int32(_a_F_pg_get_triggerdef_worker_24))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	if v309 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v324 = F_quote_identifier(m, v298)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v315 == int32(0) {
		goto L90
	} else {
		goto L98
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v324
	F_appendStringInfo(m, v320, int32(_a_F_pg_get_triggerdef_worker_25), v16+int32(48))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v334 = F_quote_identifier(m, v303)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v334
	F_appendStringInfo(m, v16+int32(304), int32(_a_F_pg_get_triggerdef_worker_26), v16+int32(32))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	goto L90
L101:
	;
	v352 = int32(_a_F_pg_get_triggerdef_worker_27)
	goto L103
L102:
	;
	v352 = int32(_a_F_pg_get_triggerdef_worker_28)
	goto L103
L103:
	;
	F_appendStringInfoString(m, v346, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v359 = F_fastgetattr_3(m, v37, int32(17), v356, v16+int32(255))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)))
	if v361 == int32(0) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	F_appendStringInfoString(m, v346, int32(_a_F_pg_get_triggerdef_worker_29))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v52)+76))
	v486 = int32(0)
	v492 = F_generate_function_name(m, v485, v486, v486, v486, v486, v486, v486)
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L126
	}
L109:
	;
	v367 = F_text_to_cstring(m, v359)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	v369 = F_stringToNode(m, v367)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v372 = F_get_rel_relkind(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	v375 = F_palloc0(m, int32(136))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v377 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v375)+12)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(v375))) = int32(101)
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v375)+24)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v375)+21)) = uint8(v372)
	*(*int32)(unsafe.Add(mBase, uint32(v375)+16)) = v381
	v388 = F_makeAlias(m, int32(_a_F_pg_get_triggerdef_worker_30), v377)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v375)+8)) = v388
	*(*int32)(unsafe.Add(mBase, uint32(v375)+4)) = v388
	v392 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v375)+124)) = uint16(v392)
	v394 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v375)+20)) = uint8(v394)
	v397 = F_palloc0(m, int32(136))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	v399 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+12)) = v399
	*(*int32)(unsafe.Add(mBase, uint32(v397))) = int32(101)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v397)+24)) = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+21)) = uint8(v372)
	*(*int32)(unsafe.Add(mBase, uint32(v397)+16)) = v403
	v410 = F_makeAlias(m, int32(_a_F_pg_get_triggerdef_worker_31), v399)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v397)+8)) = v410
	*(*int32)(unsafe.Add(mBase, uint32(v397)+4)) = v410
	v414 = int32(256)
	*(*uint16)(unsafe.Add(mBase, uint32(v397)+124)) = uint16(v414)
	v416 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v397)+20)) = uint8(v416)
	base.MemoryFill(m, v16+int32(136), v416, int32(76))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+124)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v16)+128)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v375
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v397
	v431 = F_list_make2_impl(m, v16+int32(28), v16+int32(24))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v433 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+152)) = v433
	*(*int64)(unsafe.Add(mBase, uint32(v16)+144)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+132)) = v431
	v439 = v16 + int32(132)
	F_set_rtable_names(m, v439, v433, v433)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	F_set_simple_column_names(m, v439)
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v16)+212)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v439
	v452 = F_list_make1_impl(m, int32(1), v16+int32(20))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	v454 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+244)) = uint8(v454)
	v456 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+228)) = v456
	*(*int64)(unsafe.Add(mBase, uint32(v16)+220)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+216)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v16)+248)) = v456
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+247)) = uint8(v456)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+245)) = uint16(v454)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+236)) = int64(34359738368)
	if l1 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v471 = int32(7)
	goto L123
L122:
	;
	v471 = int32(2)
	goto L123
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+232)) = v471
	F_get_rule_expr(m, v369, v16+int32(212), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_appendStringInfoString(m, v346, int32(_a_F_pg_get_triggerdef_worker_32))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L108
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v492
	F_appendStringInfo(m, v16+int32(304), int32(_a_F_pg_get_triggerdef_worker_33), v16+int32(16))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	v502 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+98)))
	if v502 <= int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_appendStringInfoChar(m, v16+int32(304), int32(41))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L1
	} else {
		goto L159
	}
L129:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v20)+52))
	v509 = F_fastgetattr_3(m, v37, int32(16), v506, v16+int32(255))
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+255)))
	if v511 == int32(1) {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	v514 = F_pg_detoast_datum_packed(m, v509)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	v517 = F_pg_detoast_datum_packed(m, v509)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v519 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+98)))
	if v519 <= int32(0) {
		goto L128
	} else {
		goto L134
	}
L134:
	;
	v522 = int32(1)
	if v516&v522 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v526 = v522
	goto L137
L136:
	;
	v526 = int32(4)
	goto L137
L137:
	;
	v529 = v517 + v526
	v530 = int32(0)
	goto L138
L138:
	;
	if v530 != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	goto L128
L140:
	;
	F_appendStringInfoString(m, v16+int32(304), int32(_a_F_pg_get_triggerdef_worker_15))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L143
	}
L141:
	;
	goto L142
L142:
	;
	F_appendStringInfoChar(m, v16+int32(304), int32(39))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	goto L142
L144:
	;
	v555 = v529
	goto L145
L145:
	;
	v565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v555))))
	v566 = base.I32_extend8_s(v565)
	if v565 != int32(39) {
		goto L149
	} else {
		goto L150
	}
L146:
	;
	F_appendStringInfoChar(m, v16+int32(304), int32(39))
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L157
	}
L147:
	;
	goto L146
L148:
	;
	F_appendStringInfoChar(m, v16+int32(304), v566)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L1
	} else {
		goto L156
	}
L149:
	;
	if v565 == int32(0) {
		goto L147
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	F_appendStringInfoChar(m, v16+int32(304), v566)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L1
	} else {
		goto L155
	}
L152:
	;
	if v566 != int32(92) {
		goto L148
	} else {
		goto L153
	}
L153:
	;
	v574 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_get_triggerdef_worker[0])))
	if v574&int32(1) != 0 {
		goto L148
	} else {
		goto L154
	}
L154:
	;
	goto L151
L155:
	;
	goto L148
L156:
	;
	v555 = v555 + int32(1)
	goto L145
L157:
	;
	v592 = F_strlen(m, v529)
	mBase = m.M
	v594 = int32(1)
	v597 = v530 + v594
	v598 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+98)))
	if v597 < v598 {
		v529 = v592 + v529 + v594
		v530 = v597
		goto L138
	} else {
		goto L158
	}
L158:
	;
	goto L139
L159:
	;
	F_systable_endscan(m, v35)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	F_relation_close(m, v20, int32(1))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v16)+304))
	v627 = v623
	goto L5
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_get_triggerdef_worker_34), v16)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_pg_get_triggerdef_worker_7), int32(1140), int32(_a_F_pg_get_triggerdef_worker_8))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_get_viewdef_name_ext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v10 = F_textToQualifiedNameList(m, v5)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = F_makeRangeVarFromNameList(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				v14 = int32(0)
				v18 = F_RangeVarGetRelidExtended(m, v12, v14, v14, v14, v14)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					if v9 != 0 {
						v22 = int32(7)
					} else {
						v22 = int32(2)
					}
					v24 = F_pg_get_viewdef_worker(m, v18, v22, int32(0))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						if v24 == int32(0) {
							v28 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
							return int32(0)
						} else {
							v32 = F_cstring_to_text(m, v24)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_pfree(m, v24)
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									return v32
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
func F_pg_indexam_has_property(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = F_pg_detoast_datum_packed(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = F_text_to_cstring(m, v5)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = int32(0)
			v13 = F_indexam_property(m, l0, v9, v3, v11, v11)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				return v13
			}
		}
	}
}
func F_pg_indexam_progress_phasename(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+28)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_GetIndexAmRoutineByAmId(m, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)+80))
			if v11 != 0 {
				v17 = m.T0[v11].(func(*base.Module, int64) int32)(m, v4)
				mBase = m.M
				if v17 == int32(0) {
					v20 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v20)
					return int32(0)
				} else {
					v24 = F_cstring_to_text(m, v17)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return v24
					}
				}
			} else {
				v13 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
				return int32(0)
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
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
func F_pg_last_wal_receive_lsn(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(0)
	v5 = F_GetWalRcvFlushRecPtr(m, v3, v3)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		if v5 == int64(0) {
			v11 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v11)
			return int32(0)
		} else {
			v15 = F_Int64GetDatum(m, v5)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_pg_log_standby_snapshot(m *base.Module, l0 int32) int32 {
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_log_standby_snapshot[0])))
	if v9 == int32(1) {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_pg_log_standby_snapshot[1]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+316))
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
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_log_standby_snapshot_0), int32(0))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_log_standby_snapshot_1), int32(216), int32(_a_F_pg_log_standby_snapshot_2))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
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
			v26 = F_LogStandbySnapshot(m)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = F_Int64GetDatum(m, v26)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v30
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_log_standby_snapshot_3), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_pg_log_standby_snapshot_4)
					F_errhint(m, int32(_a_F_pg_log_standby_snapshot_5), v5)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_log_standby_snapshot_1), int32(211), int32(_a_F_pg_log_standby_snapshot_2))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
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
func F_pg_ndistinct_recv(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn13854(m, l0, int32(_a_F_pg_ndistinct_recv_0), int32(396), int32(_a_F_pg_ndistinct_recv_1), int32(_a_F_pg_ndistinct_recv_2), int32(_a_F_pg_ndistinct_recv_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v140 int64
	_ = v140
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v165 int64
	_ = v165
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int64
	_ = v258
	var v260 int64
	_ = v260
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v366 int64
	_ = v366
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	var v414 int64
	_ = v414
	var v416 int64
	_ = v416
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v439 int64
	_ = v439
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v470 int32
	_ = v470
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v494 int32
	_ = v494
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v530 int32
	_ = v530
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
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
	var v553 int32
	_ = v553
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
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v622 int32
	_ = v622
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
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v687 int32
	_ = v687
	var v695 int32
	_ = v695
	var v700 int32
	_ = v700
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v743 int32
	_ = v743
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	if l0 != int32(100) {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L19
	} else {
		goto L173
	}
L2:
	;
	m.G0 = v17 + int32(112)
	return v719
L3:
	;
	v104 = v102
	goto L27
L4:
	;
	v102 = int32(0)
	goto L3
L5:
	;
	v102 = int32(1)
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L19
	} else {
		goto L24
	}
L7:
	;
	if l0 == int32(950) {
		v719 = int32(_a_F_pg_newlocale_from_collation_0)
		goto L2
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[0]))
	v719 = v83
	goto L2
L10:
	;
	if l0 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[1]))
	if l0 == v27 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[2]))
	v719 = v30
	goto L2
L13:
	;
	goto L14
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[3]))
	if v32 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v67 = int32(16)
	v71 = (int32(base.Ui32(l0)>>(uint(v67)%32)) ^ l0) * int32(-2048144789)
	v76 = (int32(base.Ui32(v71)>>(uint(int32(13))%32)) ^ v71) * int32(-1028477387)
	v79 = int32(base.Ui32(v76)>>(uint(v67)%32)) ^ v76
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	if base.Ui32(v66) <= base.Ui32(v80) {
		goto L4
	} else {
		goto L23
	}
L16:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v65 = v32
	v66 = v33
	goto L15
L17:
	;
	goto L18
L18:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[4]))
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(_a_F_pg_newlocale_from_collation_1), int32(0), int32(_a_F_pg_newlocale_from_collation_2), int32(_a_F_pg_newlocale_from_collation_3))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[5])) = v41
	v47 = F_MemoryContextAllocZero(m, v41, int32(32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47)+28)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+24)) = v41
	v54 = F_MemoryContextAllocExtended(m, v41, int32(512), int32(5))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v47)+12)) = int64(120259084319)
	*(*int64)(unsafe.Add(mBase, uint32(v47))) = int64(32)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+20)) = v54
	*(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[3])) = v47
	v65 = v47
	v66 = int32(28)
	goto L15
L23:
	;
	goto L5
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = int32(0)
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_4), v17+int32(96))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L19
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_5), int32(1212), int32(_a_F_pg_newlocale_from_collation_6))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L19
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
	if v104 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L29:
	;
	v715 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v715
	v104 = v715
	goto L27
L30:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L19
	} else {
		goto L170
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L19
	} else {
		goto L167
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[2])) = v653
	*(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[1])) = l0
	v719 = v653
	goto L2
L33:
	;
	v502 = *(*int32)(unsafe.Add(mBase, _c_F_pg_newlocale_from_collation[5]))
	v504 = F_SearchSysCache1(m, int32(16), l0)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L19
	} else {
		goto L119
	}
L34:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v478 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+8)) = v477 + v478
	*(*uint8)(unsafe.Add(mBase, uint32(v470)+12)) = uint8(v478)
	*(*int32)(unsafe.Add(mBase, uint32(v470)+8)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v470))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v470)+4)) = int32(0)
	v494 = v470
	goto L33
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L19
	} else {
		goto L116
	}
L36:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	if v119 == int64(4294967296) {
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v305 = v304 & v79
	v308 = v303 + v305<<(uint(int32(4))%32)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+12)))
	if v309 == int32(0) {
		v470 = v308
		goto L34
	} else {
		goto L80
	}
L39:
	;
	v122 = int32(0)
	v124 = int64(2)
	v126 = v119 << (uint(int64(1)) % 64)
	if base.Ui64(v126) <= base.Ui64(v124) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v104 = int32(1)
	goto L27
L41:
	;
	v129 = v124
	goto L43
L42:
	;
	v129 = v126
	goto L43
L43:
	;
	v130 = int64(1)
	if v129&(v129-v130) == int64(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v140 = v129
	goto L46
L45:
	;
	v140 = v130 << (uint(int64(64)-base.I64_clz(v129)) % 64)
	goto L46
L46:
	;
	if base.Ui64(v140<<(uint(int64(4))%64)) < base.Ui64(int64(2147483647)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v152 = F_MemoryContextAllocExtended(m, v147, base.I32_wrap_i64(v140)<<(uint(int32(4))%32), int32(5))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L19
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
	*(*int32)(unsafe.Add(mBase, uint32(v65)+20)) = v152
	v155 = int64(1)
	if v140&(v140-v155) == int64(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v165 = v140
	goto L53
L52:
	;
	v165 = v155 << (uint(int64(64)-base.I64_clz(v140)) % 64)
	goto L53
L53:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v165<<(uint(int64(4))%64)) {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v165
	v173 = base.I32_wrap_i64(v165) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v65)+12)) = v173
	if v165 == int64(4294967296) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v182 = int32(-85899346)
	goto L57
L56:
	;
	v182 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v165), float64(0.9)))
	goto L57
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v65)+16)) = v182
	if v146 != int64(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v187 = v122
	goto L62
L59:
	;
	goto L60
L60:
	;
	F_pfree(m, v145)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L19
	} else {
		goto L79
	}
L61:
	;
	v216 = v214
	v222 = v122
	goto L67
L62:
	;
	v202 = v145 + v187<<(uint(int32(4))%32)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+12)))
	if v203 != int32(1) {
		v214 = v187
		goto L61
	} else {
		goto L64
	}
L63:
	;
	v214 = int32(0)
	goto L61
L64:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202)+8))
	if v206&v173 == v187 {
		v214 = v187
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v210 = v187 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v210)) < base.Ui64(v146) {
		v187 = v210
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v231 = v145 + v216<<(uint(int32(4))%32)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+12)))
	if v232 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	goto L60
L69:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v231)+8))
	v241 = v236
	goto L72
L70:
	;
	goto L71
L71:
	;
	v277 = v216 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v277)) < base.Ui64(v146) {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	v251 = v241 & v235
	v256 = v152 + v251<<(uint(int32(4))%32)
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v256)+12)))
	if v257 != 0 {
		v241 = v251 + int32(1)
		goto L72
	} else {
		goto L74
	}
L73:
	;
	v258 = *(*int64)(unsafe.Add(mBase, uint32(v231)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v256)+8)) = v258
	v260 = *(*int64)(unsafe.Add(mBase, uint32(v231)))
	*(*int64)(unsafe.Add(mBase, uint32(v256))) = v260
	goto L71
L74:
	;
	goto L73
L75:
	;
	v281 = v277
	goto L77
L76:
	;
	v281 = int32(0)
	goto L77
L77:
	;
	v283 = v222 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v283)) < base.Ui64(v146) {
		v216 = v281
		v222 = v283
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
	v314 = v305
	v317 = int32(0)
	v320 = v308
	goto L81
L81:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v320)+8))
	if v327 == v79 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v449 != 0 {
		v653 = v449
		goto L32
	} else {
		goto L115
	}
L83:
	;
	goto L82
L84:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v320)))
	if v329 == l0 {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v332 = v314 + int32(1)
	v333 = v304 & v327
	if base.Ui32(v314) < base.Ui32(v333) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	goto L86
L88:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v337 = v314 + v335
	goto L90
L89:
	;
	v337 = v314
	goto L90
L90:
	;
	if base.Ui32(v337-v333) < base.Ui32(v317) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v340 = v332 & v304
	v343 = v303 + v340<<(uint(int32(4))%32)
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+12)))
	if v344 != 0 {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v434 = v317 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v434) {
		goto L110
	} else {
		goto L111
	}
L94:
	;
	v350 = v340
	v355 = int32(0)
	goto L97
L95:
	;
	v379 = v343
	v382 = v340
	goto L96
L96:
	;
	if v382 != v314 {
		goto L104
	} else {
		goto L105
	}
L97:
	;
	v361 = v355 + int32(1)
	if int32(151) <= v361 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v379 = v376
	v382 = v373
	goto L96
L99:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v366 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v364), base.F64_convert_i64_u(v366)), float64(0.1)) != 0 {
		goto L29
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v373 = (v350 + int32(1)) & v304
	v376 = v303 + v373<<(uint(int32(4))%32)
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v376)+12)))
	if v377 != 0 {
		v350 = v373
		v355 = v361
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
	v394 = v379
	v397 = v382
	goto L107
L105:
	;
	goto L106
L106:
	;
	v470 = v320
	goto L34
L107:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v410 = v407 & (v397 - int32(1))
	v413 = v303 + v410<<(uint(int32(4))%32)
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v413)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v394)+8)) = v414
	v416 = *(*int64)(unsafe.Add(mBase, uint32(v413)))
	*(*int64)(unsafe.Add(mBase, uint32(v394))) = v416
	if v410 != v314 {
		v394 = v413
		v397 = v410
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
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v65)+8))
	v439 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v437), base.F64_convert_i64_u(v439)), float64(0.1)) != 0 {
		goto L29
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v444 = v332 & v304
	v447 = v303 + v444<<(uint(int32(4))%32)
	v448 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+12)))
	if v448 != 0 {
		v314 = v444
		v317 = v434
		v320 = v447
		goto L81
	} else {
		goto L114
	}
L113:
	;
	goto L112
L114:
	;
	v470 = v447
	goto L34
L115:
	;
	v494 = v320
	goto L33
L116:
	;
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_7), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L19
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_8), int32(630), int32(_a_F_pg_newlocale_from_collation_9))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L19
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
	if v504 == int32(0) {
		goto L31
	} else {
		goto L120
	}
L120:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v504)+16))
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+22)))
	v510 = v508 + v509
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+76)))
	switch v511 - int32(98) {
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
	v539 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+4)) = uint8(v539)
	v545 = F_SysCacheGetAttr(m, int32(16), v504, int32(12), v17+int32(111))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L19
	} else {
		goto L132
	}
L122:
	;
	v536 = F_create_pg_locale_builtin(m, l0, v502)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L19
	} else {
		goto L131
	}
L123:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L19
	} else {
		goto L128
	}
L124:
	;
	v516 = F_create_pg_locale_libc(m, l0, v502)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L19
	} else {
		goto L127
	}
L125:
	;
	v514 = F_create_pg_locale_icu(m)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L19
	} else {
		goto L126
	}
L126:
	;
	v538 = v514
	goto L121
L127:
	;
	v538 = v516
	goto L121
L128:
	;
	v522 = int32(*(*int8)(unsafe.Add(mBase, uint32(v510)+76)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v522
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(_a_F_pg_newlocale_from_collation_10)
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_11), v17+int32(16))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L19
	} else {
		goto L129
	}
L129:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_5), int32(1096), int32(_a_F_pg_newlocale_from_collation_10))
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L19
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
	v538 = v536
	goto L121
L132:
	;
	v547 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+111)))
	if v547 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	F_ReleaseCatCache(m, v504)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L19
	} else {
		goto L166
	}
L134:
	;
	v548 = F_text_to_cstring(m, v545)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L19
	} else {
		goto L135
	}
L135:
	;
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+76)))
	if v553 == int32(99) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v556 = int32(8)
	goto L138
L137:
	;
	v556 = int32(10)
	goto L138
L138:
	;
	v557 = F_SysCacheGetAttrNotNull(m, int32(16), v504, v556)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L19
	} else {
		goto L139
	}
L139:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v510)+76)))
	v560 = F_text_to_cstring(m, v557)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L19
	} else {
		goto L140
	}
L140:
	;
	switch v559 - int32(98) {
	case 0:
		goto L143
	case 1:
		goto L142
	default:
		goto L30
	}
L141:
	;
	if v578 == int32(0) {
		goto L30
	} else {
		goto L149
	}
L142:
	;
	v567 = F_pg_strcasecmp(m, int32(_a_F_pg_newlocale_from_collation_12), v560)
	mBase = m.M
	if v567 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	v564 = F_get_collation_actual_version_builtin(m, v560)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L19
	} else {
		goto L144
	}
L144:
	;
	v578 = v564
	goto L141
L145:
	;
	v578 = int32(0)
	goto L141
L146:
	;
	goto L145
L147:
	;
	v572 = F_pg_strncasecmp(m, int32(_a_F_pg_newlocale_from_collation_13), v560, int32(2))
	mBase = m.M
	if v572 == int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v576 = F_pg_strcasecmp(m, int32(_a_F_pg_newlocale_from_collation_14), v560)
	mBase = m.M
	goto L146
L149:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v578))))
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
	if base.B2i32(v583 == int32(0))|base.B2i32(v583 != v586) != 0 {
		v604 = v583
		v605 = v586
		goto L151
	} else {
		goto L152
	}
L150:
	;
	if v604-v605 == int32(0) {
		goto L133
	} else {
		goto L157
	}
L151:
	;
	goto L150
L152:
	;
	v589 = v578
	v590 = v548
	goto L153
L153:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590)+1)))
	v594 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v589)+1)))
	if v594 == int32(0) {
		v604 = v594
		v605 = v593
		goto L151
	} else {
		goto L155
	}
L154:
	;
	v604 = v594
	v605 = v593
	goto L151
L155:
	;
	v597 = int32(1)
	if v594 == v593 {
		v589 = v589 + v597
		v590 = v590 + v597
		goto L153
	} else {
		goto L156
	}
L156:
	;
	goto L154
L157:
	;
	v611 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L19
	} else {
		goto L158
	}
L158:
	;
	if v611 == int32(0) {
		goto L133
	} else {
		goto L159
	}
L159:
	;
	v616 = v510 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v616
	F_errmsg(m, int32(_a_F_pg_newlocale_from_collation_15), v17+int32(80))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L19
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v578
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v548
	F_errdetail(m, int32(_a_F_pg_newlocale_from_collation_16), v17-int32(-64))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L19
	} else {
		goto L161
	}
L161:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v510)+68))
	v631 = F_get_namespace_name(m, v630)
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L19
	} else {
		goto L162
	}
L162:
	;
	v633 = F_quote_qualified_identifier(m, v631, v616)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L19
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v633
	F_errhint(m, int32(_a_F_pg_newlocale_from_collation_17), v17+int32(48))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L19
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_5), int32(1142), int32(_a_F_pg_newlocale_from_collation_10))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L19
	} else {
		goto L165
	}
L165:
	;
	goto L133
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v494)+4)) = v538
	v653 = v538
	goto L32
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_4), v17)
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L19
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_5), int32(1085), int32(_a_F_pg_newlocale_from_collation_10))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L19
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v510 + int32(4)
	F_errmsg(m, int32(_a_F_pg_newlocale_from_collation_18), v17+int32(32))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L19
	} else {
		goto L171
	}
L171:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_5), int32(1128), int32(_a_F_pg_newlocale_from_collation_10))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L19
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
	F_errmsg_internal(m, int32(_a_F_pg_newlocale_from_collation_19), int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L19
	} else {
		goto L174
	}
L174:
	;
	F_errfinish(m, int32(_a_F_pg_newlocale_from_collation_8), int32(327), int32(_a_F_pg_newlocale_from_collation_20))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L19
	} else {
		goto L175
	}
L175:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_notification_queue_usage(m *base.Module, l0 int32) int32 {
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
	var v27 float64
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	F_asyncQueueAdvanceTail(m)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_pg_notification_queue_usage[0]))
		v14 = F_LWLockAcquire(m, v10+int32(3456), int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_notification_queue_usage[1]))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
			v19 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
			if v18 != v19 {
				v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_notification_queue_usage[2]))
				v27 = base.F64_div(base.F64_convert_i64_s(v18-v19), base.F64_convert_i32_s(v24))
			} else {
				v27 = float64(0)
			}
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_notification_queue_usage[0]))
			F_LWLockRelease(m, v29+int32(3456))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = F_Float8GetDatum(m, v27)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return v34
				}
			}
		}
	}
}
func F_pg_numa_available(m *base.Module, l0 int32) int32 {
	return int32(0)
}
func F_pg_opfamily_is_visible(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_OpfamilyIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_options_to_table(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
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
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
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
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	m.G0 = v10 + int32(16)
	return int32(0)
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v27<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v41 = F_cstring_to_text(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+6)) = uint8(v43)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v41
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v50 = F_cstring_to_text(m, v49)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v52 = int32(1)
	v53 = int32(0)
	goto L12
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+7)) = uint8(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	F_tuplestore_putvalues(m, v56, v57, v10+int32(8), v10+int32(6))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v52 = int32(0)
	v53 = v50
	goto L12
L14:
	;
	v65 = v27 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v65 < v66 {
		v27 = v65
		goto L7
	} else {
		goto L15
	}
L15:
	;
	goto L8
}
func F_pg_partition_ancestors(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+16))
	if v7 == int32(0) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v11 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v16 = int32(0)
			v19 = F_SearchSysCacheExists(m, int32(57), v10, v16, v16, v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if v19 == int32(0) {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = int32(2)
						v89 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
						return int32(0)
					}
				} else {
					v23 = F_get_rel_relkind(m, v10)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = F_get_rel_relispartition(m, v10)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							if base.B2i32(v25|base.B2i32(v23 == int32(112)) == int32(0))&base.B2i32(v23&int32(255) != int32(73)) != 0 {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = int32(2)
									v89 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
									return int32(0)
								}
							} else {
								v37 = int32(_a_F_pg_partition_ancestors_0)
								v38 = *(*int32)(unsafe.Add(mBase, _c_F_pg_partition_ancestors[0]))
								v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
								*(*int32)(unsafe.Add(mBase, _c_F_pg_partition_ancestors[0])) = v40
								v42 = F_get_partition_ancestors(m, v10)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = F_lcons_oid(m, v10, v42)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v44
										*(*int32)(unsafe.Add(mBase, _c_F_pg_partition_ancestors[0])) = v38
										v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
										v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
										if v54 == int32(0) {
											F_end_MultiFuncCall(m, l0)
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
												v79 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
												return int32(0)
											}
										} else {
											v57 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
											v58 = int64(*(*int32)(unsafe.Add(mBase, uint32(v54)+4)))
											if base.Ui64(v58) <= base.Ui64(v57) {
												F_end_MultiFuncCall(m, l0)
												mBase = m.M
												v75 = m.ExcPending
												if v75 != 0 {
													return int32(0)
												} else {
													v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
													v79 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
													return int32(0)
												}
											} else {
												v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
												v65 = *(*int32)(unsafe.Add(mBase, uint32(v60+base.I32_wrap_i64(v57)<<(uint(int32(2))%32))))
												*(*int64)(unsafe.Add(mBase, uint32(v53))) = v57 + int64(1)
												v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v69)+20)) = int32(1)
												return v65
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
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+16))
		if v54 == int32(0) {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
				v79 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
				return int32(0)
			}
		} else {
			v57 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
			v58 = int64(*(*int32)(unsafe.Add(mBase, uint32(v54)+4)))
			if base.Ui64(v58) <= base.Ui64(v57) {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v76)+20)) = int32(2)
					v79 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v79)
					return int32(0)
				}
			} else {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v60+base.I32_wrap_i64(v57)<<(uint(int32(2))%32))))
				*(*int64)(unsafe.Add(mBase, uint32(v53))) = v57 + int64(1)
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v69)+20)) = int32(1)
				return v65
			}
		}
	}
}
func F_pg_partition_tree(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
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
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == v2 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(32)
	return v203
L2:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L48
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L45
	}
L4:
	;
	v20 = F_init_MultiFuncCall(m, l0)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	goto L18
L7:
	;
	return int32(0)
L8:
	;
	v25 = int32(0)
	v28 = F_SearchSysCacheExists(m, int32(57), v15, v25, v25, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v28 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v32 = F_get_rel_relkind(m, v15)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v34 = F_get_rel_relispartition(m, v15)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	if base.B2i32(v34|base.B2i32(v32 == int32(112)) == int32(0))&base.B2i32(v32&int32(255) != int32(73)) != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v46 = int32(_a_F_pg_partition_tree_0)
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pg_partition_tree[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_partition_tree[0])) = v49
	v53 = F_find_all_inheritors(m, v15, int32(1), int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v58 = F_get_call_result_type(m, l0, int32(0), v13+int32(16))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	if v58 != int32(1) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v62
	*(*int32)(unsafe.Add(mBase, _c_F_pg_partition_tree[0])) = v47
	goto L6
L17:
	;
	F_end_MultiFuncCall(m, l0)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L44
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
	if v73 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	v77 = int64(*(*int32)(unsafe.Add(mBase, uint32(v73)+4)))
	if base.Ui64(v77) <= base.Ui64(v76) {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v79 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v79
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v79
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v83
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v86+base.I32_wrap_i64(v76)<<(uint(int32(2))%32))))
	v92 = F_get_rel_relkind(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v94 = F_get_partition_ancestors(m, v91)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v91
	if v94 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v145
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v72)+28))
	v158 = F_heap_form_tuple(m, v153, v13+int32(16), v13+int32(12))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L42
	}
L24:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	v119 = int32(0)
	if v119 < v118 {
		goto L33
	} else {
		goto L34
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = base.B2i32(v92 != int32(112)) & base.B2i32(v92 != int32(73))
	if v91 == v15 {
		v145 = v83
		goto L23
	} else {
		goto L32
	}
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v98 != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v100 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+13)) = uint8(v100)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = base.B2i32(v92 != int32(112)) & base.B2i32(v92 != int32(73))
	if v91 == v15 {
		v145 = v83
		goto L23
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v94 != 0 {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v145 = v83
	goto L23
L32:
	;
	goto L24
L33:
	;
	v122 = v118
	goto L35
L34:
	;
	v122 = v119
	goto L35
L35:
	;
	v126 = v83
	goto L36
L36:
	;
	if v122 == v126 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v145 = v137
	goto L23
L38:
	;
	v145 = v122
	goto L23
L39:
	;
	goto L40
L40:
	;
	v137 = v126 + int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v126<<(uint(int32(2))%32)+v138)))
	if v140 != v15 {
		v126 = v137
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v161 = F_HeapTupleHeaderGetDatum(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v163 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
	*(*int64)(unsafe.Add(mBase, uint32(v72))) = v163 + int64(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+20)) = int32(1)
	v203 = v161
	goto L1
L44:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+20)) = int32(2)
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v176)
	v203 = int32(0)
	goto L1
L45:
	;
	F_errmsg_internal(m, int32(_a_F_pg_partition_tree_1), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pg_partition_tree_2), int32(91), int32(_a_F_pg_partition_tree_3))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L7
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
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+20)) = int32(2)
	v198 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v198)
	v203 = v2
	goto L1
}
func F_pg_postmaster_start_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v3 = *(*int64)(unsafe.Add(mBase, _c_F_pg_postmaster_start_time[0]))
	v4 = F_Int64GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_pg_prepared_statement(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v144 int32
	_ = v144
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	F_InitMaterializedSRF(m, l0, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_pg_prepared_statement[0]))
	if v26 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v17 + int32(80)
	return int32(0)
L4:
	;
	v30 = v17 + int32(60)
	F_hash_seq_init(m, v30, v26)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v33 = F_hash_seq_search(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v33 == int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v44 = v33
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v44)+64))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+52))
	v55 = F_cstring_to_text(m, v44)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L3
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v44)+64))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v60 = F_cstring_to_text(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v60
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v44)+72))
	v64 = F_Int64GetDatum(m, v63)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v64
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v44)+64))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v67)+24))
	v72 = F_palloc(m, v69<<(uint(int32(2))%32))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v69 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v184 = F_construct_array_builtin(m, v72, v69, int32(2206))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L26
	}
L15:
	;
	v77 = v69 & int32(3)
	v78 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v69) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v85 = v78
	v93 = int32(0)
	goto L19
L17:
	;
	v130 = v78
	goto L18
L18:
	;
	v144 = v130
	v155 = v78
	goto L23
L19:
	;
	v100 = v85 << (uint(int32(2)) % 32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v68+v100)))
	*(*int32)(unsafe.Add(mBase, uint32(v72+v100))) = v103
	v105 = int32(4)
	v106 = v100 | v105
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v68+v106)))
	*(*int32)(unsafe.Add(mBase, uint32(v72+v106))) = v109
	v112 = v100 | int32(8)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v68+v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v72+v112))) = v115
	v118 = v100 | int32(12)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v68+v118)))
	*(*int32)(unsafe.Add(mBase, uint32(v72+v118))) = v121
	v124 = v85 + v105
	v126 = v93 + v105
	if v126 != v69&int32(2147483644) {
		v85 = v124
		v93 = v126
		goto L19
	} else {
		goto L21
	}
L20:
	;
	if v77 == int32(0) {
		goto L14
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v130 = v124
	goto L18
L23:
	;
	v159 = v144 << (uint(int32(2)) % 32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v68+v159)))
	*(*int32)(unsafe.Add(mBase, uint32(v72+v159))) = v162
	v164 = int32(1)
	v167 = v155 + v164
	if v167 != v77 {
		v144 = v144 + v164
		v155 = v167
		goto L23
	} else {
		goto L25
	}
L24:
	;
	goto L14
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v184
	if v54 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+68)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v362
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v44)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v364 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v364 + int32(136)
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	F_tuplestore_putvalues(m, v371, v372, v17+int32(16), v17+int32(8))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L53
	}
L28:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v190 = F_palloc(m, v187<<(uint(int32(2))%32))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v346 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+12)) = uint8(v346)
	goto L27
L31:
	;
	v192 = int32(0)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v193 <= v192 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v343 = F_construct_array_builtin(m, v332, v329, int32(2206))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L52
	}
L33:
	;
	v198 = F_palloc(m, v193<<(uint(int32(2))%32))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v200 = v192
	v201 = v193
	goto L37
L36:
	;
	v329 = v193
	v332 = v198
	goto L32
L37:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v54+v201<<(uint(int32(4))%32)+v200*int32(100))+88))
	*(*int32)(unsafe.Add(mBase, uint32(v190+v200<<(uint(int32(2))%32)))) = v223
	v226 = v200 + int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v226 < v227 {
		v200 = v226
		v201 = v227
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v231 = F_palloc(m, v227<<(uint(int32(2))%32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	if v227 <= int32(0) {
		v329 = v227
		v332 = v231
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v236 = v227 & int32(3)
	v237 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v227) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v244 = v237
	v252 = int32(0)
	goto L45
L43:
	;
	v289 = v237
	goto L44
L44:
	;
	v303 = v289
	v312 = v237
	goto L49
L45:
	;
	v259 = v244 << (uint(int32(2)) % 32)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v259+v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v259))) = v262
	v264 = int32(4)
	v265 = v259 | v264
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v190+v265)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v265))) = v268
	v271 = v259 | int32(8)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v190+v271)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v271))) = v274
	v277 = v259 | int32(12)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v277+v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v277))) = v280
	v283 = v244 + v264
	v285 = v252 + v264
	if v285 != v227&int32(2147483644) {
		v244 = v283
		v252 = v285
		goto L45
	} else {
		goto L47
	}
L46:
	;
	if v236 == int32(0) {
		v329 = v227
		v332 = v231
		goto L32
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v289 = v283
	goto L44
L49:
	;
	v318 = v303 << (uint(int32(2)) % 32)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v318+v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v231+v318))) = v321
	v323 = int32(1)
	v326 = v312 + v323
	if v326 != v236 {
		v303 = v303 + v323
		v312 = v326
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v329 = v227
	v332 = v231
	goto L32
L51:
	;
	goto L50
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v343
	goto L27
L53:
	;
	v381 = F_hash_seq_search(m, v17+int32(60))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v381 != 0 {
		v44 = v381
		goto L8
	} else {
		goto L55
	}
L55:
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
func F_pg_random_bytes(m *base.Module, l0 int32) int32 {
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(-1025)) < base.Ui32(v4-int32(1025)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_px_THROW_ERROR(m, int32(-17))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
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
	v76 = m.ExcPending
	if v76 != 0 {
		goto L5
	} else {
		goto L21
	}
L5:
	;
	return int32(0)
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
	return v11
L21:
	;
	F_errcode(m, int32(579))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(_a_F_pg_random_bytes_1), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_pg_random_bytes_2), int32(465), int32(_a_F_pg_random_bytes_3))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
func F_pg_random_uuid(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_gen_random_uuid(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_pg_read_binary_file_off_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
		if int64(0) <= v11 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
			v16 = F_convert_and_check_filename(m, v6)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v19 = F_read_binary_file(m, v16, v15, v11, int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					if v19 == int32(0) {
						v23 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
						return int32(0)
					} else {
						return v19
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_read_binary_file_off_len_0), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_read_binary_file_off_len_1), int32(269), int32(_a_F_pg_read_binary_file_off_len_2))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
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
func F_pg_relpages(m *base.Module, l0 int32) int32 {
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_superuser(m)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			if v7 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_pg_relpages_0), int32(0))
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_relpages_1), int32(401), int32(_a_F_pg_relpages_2))
							mBase = m.M
							v26 = m.ExcPending
							if v26 != 0 {
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
				v27 = F_textToQualifiedNameList(m, v3)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = F_makeRangeVarFromNameList(m, v27)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v32 = F_relation_openrv(m, v29, int32(1))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = F_pg_relpages_impl(m, v32)
							mBase = m.M
							v35 = m.ExcPending
							if v35 != 0 {
								return int32(0)
							} else {
								v36 = F_Int64GetDatum(m, v34)
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									return v36
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_relpages_v1_5(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_textToQualifiedNameList(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = F_makeRangeVarFromNameList(m, v7)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				v12 = F_relation_openrv(m, v9, int32(1))
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					v14 = F_pg_relpages_impl(m, v12)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						v16 = F_Int64GetDatum(m, v14)
						mBase = m.M
						v17 = m.ExcPending
						if v17 != 0 {
							return int32(0)
						} else {
							return v16
						}
					}
				}
			}
		}
	}
}
func F_pg_relpagesbyid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_superuser(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_relpagesbyid_0), int32(0))
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_relpagesbyid_1), int32(433), int32(_a_F_pg_relpagesbyid_2))
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
		} else {
			v26 = F_relation_open(m, v2, int32(1))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				v28 = F_pg_relpages_impl(m, v26)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = F_Int64GetDatum(m, v28)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						return v30
					}
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
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
					v94 = int32(1)
					m.G0 = v8 + int32(16)
					return v94
				} else {
					v80 = int32(_a_F_pg_signal_backend_0)
					v82 = int32(74)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg(m, v80, v8)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_signal_backend_1), v82, int32(_a_F_pg_signal_backend_2))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							v94 = int32(1)
							m.G0 = v8 + int32(16)
							return v94
						}
					}
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
			if v24 != 0 {
				v25 = F_superuser_arg(m, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[0]))
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v10)+64))
						v57 = F_has_privs_of_role(m, v55, v56)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v66 = int32(0)
								v69 = F_pgmem_kill(m, v66-l0, l1)
								mBase = m.M
								if v69 == v66 {
									v94 = v66
									m.G0 = v8 + int32(16)
									return v94
								} else {
									v74 = F_errstart(m, int32(19), int32(0))
									mBase = m.M
									v75 = m.ExcPending
									if v75 != 0 {
										return int32(0)
									} else {
										if v74 == int32(0) {
											v94 = int32(1)
											m.G0 = v8 + int32(16)
											return v94
										} else {
											v80 = int32(_a_F_pg_signal_backend_3)
											v82 = int32(123)
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
											F_errmsg(m, v80, v8)
											mBase = m.M
											v85 = m.ExcPending
											if v85 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_pg_signal_backend_1), v82, int32(_a_F_pg_signal_backend_2))
												mBase = m.M
												v89 = m.ExcPending
												if v89 != 0 {
													return int32(0)
												} else {
													v94 = int32(1)
													m.G0 = v8 + int32(16)
													return v94
												}
											}
										}
									}
								}
							} else {
								v60 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[0]))
								v62 = F_has_privs_of_role(m, v60, int32(_a_F_pg_signal_backend_4))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									if v62 != 0 {
										v66 = int32(0)
										v69 = F_pgmem_kill(m, v66-l0, l1)
										mBase = m.M
										if v69 == v66 {
											v94 = v66
											m.G0 = v8 + int32(16)
											return v94
										} else {
											v74 = F_errstart(m, int32(19), int32(0))
											mBase = m.M
											v75 = m.ExcPending
											if v75 != 0 {
												return int32(0)
											} else {
												if v74 == int32(0) {
													v94 = int32(1)
													m.G0 = v8 + int32(16)
													return v94
												} else {
													v80 = int32(_a_F_pg_signal_backend_3)
													v82 = int32(123)
													*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
													F_errmsg(m, v80, v8)
													mBase = m.M
													v85 = m.ExcPending
													if v85 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_pg_signal_backend_1), v82, int32(_a_F_pg_signal_backend_2))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															v94 = int32(1)
															m.G0 = v8 + int32(16)
															return v94
														}
													}
												}
											}
										}
									} else {
										v94 = int32(2)
										m.G0 = v8 + int32(16)
										return v94
									}
								}
							}
						}
					} else {
						v29 = int32(4)
						v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[1]))
						v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[2]))
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v37 = base.I32_div_s(v10-v34, int32(640))
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v31+v37*int32(408))+8))
						if v41 == v29 {
							v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[0]))
							v47 = F_has_privs_of_role(m, v45, int32(_a_F_pg_signal_backend_5))
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								if v47 == int32(0) {
									v94 = v29
									m.G0 = v8 + int32(16)
									return v94
								} else {
									v66 = int32(0)
									v69 = F_pgmem_kill(m, v66-l0, l1)
									mBase = m.M
									if v69 == v66 {
										v94 = v66
										m.G0 = v8 + int32(16)
										return v94
									} else {
										v74 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											if v74 == int32(0) {
												v94 = int32(1)
												m.G0 = v8 + int32(16)
												return v94
											} else {
												v80 = int32(_a_F_pg_signal_backend_3)
												v82 = int32(123)
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
												F_errmsg(m, v80, v8)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_signal_backend_1), v82, int32(_a_F_pg_signal_backend_2))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v94 = int32(1)
														m.G0 = v8 + int32(16)
														return v94
													}
												}
											}
										}
									}
								}
							}
						} else {
							v51 = F_superuser(m)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								if v51 != 0 {
									v66 = int32(0)
									v69 = F_pgmem_kill(m, v66-l0, l1)
									mBase = m.M
									if v69 == v66 {
										v94 = v66
										m.G0 = v8 + int32(16)
										return v94
									} else {
										v74 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int32(0)
										} else {
											if v74 == int32(0) {
												v94 = int32(1)
												m.G0 = v8 + int32(16)
												return v94
											} else {
												v80 = int32(_a_F_pg_signal_backend_3)
												v82 = int32(123)
												*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
												F_errmsg(m, v80, v8)
												mBase = m.M
												v85 = m.ExcPending
												if v85 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_pg_signal_backend_1), v82, int32(_a_F_pg_signal_backend_2))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v94 = int32(1)
														m.G0 = v8 + int32(16)
														return v94
													}
												}
											}
										}
									}
								} else {
									v94 = int32(3)
									m.G0 = v8 + int32(16)
									return v94
								}
							}
						}
					}
				}
			} else {
				v29 = int32(4)
				v31 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[1]))
				v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[2]))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				v37 = base.I32_div_s(v10-v34, int32(640))
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v31+v37*int32(408))+8))
				if v41 == v29 {
					v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_signal_backend[0]))
					v47 = F_has_privs_of_role(m, v45, int32(_a_F_pg_signal_backend_5))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						if v47 == int32(0) {
							v94 = v29
							m.G0 = v8 + int32(16)
							return v94
						} else {
							v66 = int32(0)
							v69 = F_pgmem_kill(m, v66-l0, l1)
							mBase = m.M
							if v69 == v66 {
								v94 = v66
								m.G0 = v8 + int32(16)
								return v94
							} else {
								v74 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									if v74 == int32(0) {
										v94 = int32(1)
										m.G0 = v8 + int32(16)
										return v94
									} else {
										v80 = int32(_a_F_pg_signal_backend_3)
										v82 = int32(123)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg(m, v80, v8)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_signal_backend_1), v82, int32(_a_F_pg_signal_backend_2))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v94 = int32(1)
												m.G0 = v8 + int32(16)
												return v94
											}
										}
									}
								}
							}
						}
					}
				} else {
					v51 = F_superuser(m)
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						if v51 != 0 {
							v66 = int32(0)
							v69 = F_pgmem_kill(m, v66-l0, l1)
							mBase = m.M
							if v69 == v66 {
								v94 = v66
								m.G0 = v8 + int32(16)
								return v94
							} else {
								v74 = F_errstart(m, int32(19), int32(0))
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int32(0)
								} else {
									if v74 == int32(0) {
										v94 = int32(1)
										m.G0 = v8 + int32(16)
										return v94
									} else {
										v80 = int32(_a_F_pg_signal_backend_3)
										v82 = int32(123)
										*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
										F_errmsg(m, v80, v8)
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_pg_signal_backend_1), v82, int32(_a_F_pg_signal_backend_2))
											mBase = m.M
											v89 = m.ExcPending
											if v89 != 0 {
												return int32(0)
											} else {
												v94 = int32(1)
												m.G0 = v8 + int32(16)
												return v94
											}
										}
									}
								}
							}
						} else {
							v94 = int32(3)
							m.G0 = v8 + int32(16)
							return v94
						}
					}
				}
			}
		}
	}
}
func F_pg_snapshot_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v16 = v8 - int32(-64)
	F_initStringInfo(m, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v19
	F_appendStringInfo(m, v16, int32(_a_F_pg_snapshot_out_0), v8+int32(48))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v26
	F_appendStringInfo(m, v16, int32(_a_F_pg_snapshot_out_0), v8+int32(32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v33 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
	m.G0 = v8 + int32(80)
	return v76
L7:
	;
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v36
	F_appendStringInfo(m, v16, int32(_a_F_pg_snapshot_out_1), v8+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v43) < base.Ui32(int32(2)) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v49 = int32(1)
	goto L10
L10:
	;
	v55 = v8 - int32(-64)
	F_appendStringInfoChar(m, v55, int32(44))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	goto L6
L12:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v11+int32(24)+v49<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v62
	F_appendStringInfo(m, v55, int32(_a_F_pg_snapshot_out_1), v8)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v68 = v49 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if base.Ui32(v68) < base.Ui32(v69) {
		v49 = v68
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
}
func F_pg_snapshot_recv(m *base.Module, l0 int32) int32 {
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
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L22
	}
L2:
	;
	return int32(0)
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
	return v36
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
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	F_errmsg(m, int32(_a_F_pg_snapshot_recv_0), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	F_errfinish(m, int32(_a_F_pg_snapshot_recv_1), int32(522), int32(_a_F_pg_snapshot_recv_2))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
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
func F_pg_split_walfile_name(m *base.Module, l0 int32) int32 {
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
	var v47 int32
	_ = v47
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v145 int32
	_ = v145
	var v149 int64
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int64
	_ = v164
	var v165 int64
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int64
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	v10 = m.G0
	v12 = v10 - int32(336)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+322)) = uint16(v21)
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
	v28 = v23
	v29 = v25
	goto L8
L6:
	;
	goto L7
L7:
	;
	v59 = F_strlen(m, v23)
	mBase = m.M
	if v59 != int32(24) {
		goto L16
	} else {
		goto L17
	}
L8:
	;
	if base.Ui32((v29-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v45)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v47 != 0 {
		v28 = v28 + int32(1)
		v29 = v47
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v43 = v29 - int32(32)
	goto L13
L12:
	;
	v43 = v29
	goto L13
L13:
	;
	v45 = v43 & int32(255)
	goto L10
L14:
	;
	goto L9
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L50
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L46
	}
L17:
	;
	v62 = int32(_a_F_pg_split_walfile_name_0)
	v66 = m.G0
	v68 = v66 - int32(32)
	v69 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v68)+24)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v68)+16)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v68)+8)) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v68))) = v69
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_split_walfile_name[0])))
	if v77 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v145 != int32(24) {
		goto L16
	} else {
		goto L37
	}
L19:
	;
	v145 = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_split_walfile_name[1])))
	if v81 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v85 = v23
	goto L25
L23:
	;
	goto L24
L24:
	;
	v95 = v62
	v96 = v77
	goto L28
L25:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85))))
	if v91 == v77 {
		v85 = v85 + int32(1)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v145 = v85 - v23
	goto L18
L27:
	;
	goto L26
L28:
	;
	v103 = v68 + int32(base.Ui32(v96)>>(uint(int32(3))%32))&int32(28)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	v105 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v104 | v105<<(uint(v96)%32)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
	if v109 != 0 {
		v95 = v95 + v105
		v96 = v109
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	if v112 == int32(0) {
		v135 = v23
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v145 = v135 - v23
	goto L18
L32:
	;
	v116 = v23
	v117 = v112
	goto L33
L33:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v68+int32(base.Ui32(v117)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v125)>>(uint(v117)%32))&int32(1) == int32(0) {
		v135 = v116
		goto L31
	} else {
		goto L35
	}
L34:
	;
	v135 = v133
	goto L31
L35:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+1)))
	v133 = v116 + int32(1)
	if v131 != 0 {
		v116 = v133
		v117 = v131
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v149 = int64(*(*int32)(unsafe.Add(mBase, _c_F_pg_split_walfile_name[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v12 + int32(332)
	v154 = v12 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v154
	v157 = v12 + int32(324)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v157
	v162 = F_sscanf(m, v23, int32(_a_F_pg_split_walfile_name_1), v12+int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v164 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+324)))
	v165 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+48)))
	v169 = F_get_call_result_type(m, l0, int32(0), v12+int32(316))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v169 != int32(1) {
		goto L15
	} else {
		goto L40
	}
L40:
	;
	v174 = base.I64_div_u_s(int64(4294967296), v149)
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v174*v165 + v164
	v180 = F_pg_snprintf(m, v154, int32(256), int32(_a_F_pg_split_walfile_name_2), v12)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v183 = int32(0)
	v186 = F_DirectFunctionCall3Coll(m, int32(408), v183, v154, v183, int32(-1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+324)) = v186
	v189 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v12)+332)))
	v190 = F_Int64GetDatum(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+328)) = v190
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v12)+316))
	v196 = F_heap_form_tuple(m, v193, v157, v12+int32(322))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v196)+16))
	v199 = F_HeapTupleHeaderGetDatum(m, v198)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	m.G0 = v12 + int32(336)
	return v199
L46:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v19
	F_errmsg(m, int32(_a_F_pg_split_walfile_name_3), v12+int32(32))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_pg_split_walfile_name_4), int32(487), int32(_a_F_pg_split_walfile_name_5))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_pg_split_walfile_name_6), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_pg_split_walfile_name_4), int32(492), int32(_a_F_pg_split_walfile_name_5))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_statistics_obj_is_visible(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	v2 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)) = uint8(v2)
	v14 = F_StatisticsObjIsVisibleExt(m, v9, v7+int32(15))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
		if v18 == int32(1) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v23 = v2
		} else {
			v23 = v14
		}
		m.G0 = v7 + int32(16)
		return v23
	}
}
func F_pg_stats_ext_mcvlist_items(m *base.Module, l0 int32) int32 {
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
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
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 float64
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 float64
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(48)
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
	v217 = m.ExcPending
	if v217 != 0 {
		goto L5
	} else {
		goto L45
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
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
	goto L17
L5:
	;
	return int32(0)
L6:
	;
	v24 = int32(_a_F_pg_stats_ext_mcvlist_items_0)
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0])) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	v40 = F_get_call_result_type(m, l0, int32(0), v14)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if v40 != int32(1) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v45 = F_BlessTupleDesc(m, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v45
	v48 = F_TupleDescGetAttInMetadata(m, v45)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v48
	*(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0])) = v25
	goto L4
L16:
	;
	m.G0 = v14 + int32(48)
	return v201
L17:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
	if base.Ui64(v58) < base.Ui64(v59) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57)+16))
	v62 = base.I32_wrap_i64(v58)
	v67 = v61 + v62*int32(24) + int32(48)
	v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+12)))
	if int32(0) < v68 {
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
	v192 = m.ExcPending
	if v192 != 0 {
		goto L5
	} else {
		goto L44
	}
L21:
	;
	v76 = int32(0)
	v80 = v2
	v81 = v2
	goto L24
L22:
	;
	v146 = v2
	v147 = v2
	v151 = v62
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v151
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v155 = F_makeArrayResult(m, v147, v154)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L38
	}
L24:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85+v76))))
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v92 = F_accumArrayResult(m, v80, v87, int32(0), int32(16), v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L26
	}
L25:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v146 = v92
	v147 = v134
	v151 = v139
	goto L23
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v76))))
	if v96 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v136 = v76 + int32(1)
	v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+12)))
	if v136 < v137 {
		v76 = v136
		v80 = v92
		v81 = v134
		goto L24
	} else {
		goto L37
	}
L28:
	;
	v100 = v76 << (uint(int32(2)) % 32)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v61+int32(16)+v100)))
	F_getTypeOutputInfo(m, v102, v14+int32(40), v14+int32(39))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L5
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v131 = F_accumArrayResult(m, v81, int32(0), int32(1), int32(25), v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L36
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	F_fmgr_info(m, v109, v14)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v67)+20))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v113+v100)))
	v116 = F_FunctionCall1Coll(m, v14, int32(0), v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L33
	}
L33:
	;
	v118 = F_cstring_to_text(m, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v124 = F_accumArrayResult(m, v81, v118, int32(0), int32(25), v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v134 = v124
	goto L27
L36:
	;
	v134 = v131
	goto L27
L37:
	;
	goto L25
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v155
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stats_ext_mcvlist_items[0]))
	v160 = F_makeArrayResult(m, v146, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v160
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v67)))
	v164 = F_Float8GetDatum(m, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v164
	v167 = *(*float64)(unsafe.Add(mBase, uint32(v67)+8))
	v168 = F_Float8GetDatum(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v170 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v168
	*(*uint8)(unsafe.Add(mBase, uint32(v14)+44)) = uint8(v170)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v179 = F_heap_form_tuple(m, v176, v14, v14+int32(40))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v179)+16))
	v182 = F_HeapTupleHeaderGetDatum(m, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v184 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = v184 + int64(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v188)+20)) = int32(1)
	v201 = v182
	goto L16
L44:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v193)+20)) = int32(2)
	v196 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v196)
	v201 = int32(0)
	goto L16
L45:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_pg_stats_ext_mcvlist_items_1), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_pg_stats_ext_mcvlist_items_2), int32(1369), int32(_a_F_pg_stats_ext_mcvlist_items_3))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L5
	} else {
		goto L48
	}
L48:
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
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
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
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
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v396 int32
	_ = v396
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v480 int32
	_ = v480
	v6 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	switch v14 - int32(98) {
	case 0:
		goto L2
	case 1:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v480
L2:
	;
	v448 = m.G0
	v450 = v448 - int32(16)
	m.G0 = v450
	v452 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v450)+8)) = v452
	*(*int32)(unsafe.Add(mBase, uint32(v450)+4)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v450))) = l2
	v456 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v450)+15)) = uint8(v452)
	*(*uint16)(unsafe.Add(mBase, uint32(v450)+13)) = uint16(v452)
	v461 = int32(1)
	v462 = v456 ^ v461
	*(*uint8)(unsafe.Add(mBase, uint32(v450)+12)) = uint8(v462)
	v466 = F_convert_case(m, l0, l1, l2, l3, v461, v456, int32(1462), v450)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L15
	} else {
		goto L156
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L15
	} else {
		goto L153
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strtitle[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v19*int32(28))+uint32(_c_F_pg_strtitle[1])))
	goto L7
L5:
	;
	v480 = v431
	goto L1
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L15
	} else {
		goto L149
	}
L7:
	;
	if int32(2) <= v24 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if l3 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L11:
	;
	v30 = F_strlen(m, l2)
	mBase = m.M
	v31 = v30
	goto L13
L12:
	;
	v31 = l3
	goto L13
L13:
	;
	v33 = v31 + int32(1)
	if base.Ui32(int32(536870912)) <= base.Ui32(v33) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v38 = F_palloc(m, v33<<(uint(int32(2))%32))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	F_char2wchar(m, v38, v33, l2, v31, l4)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v44 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = v44
	v50 = int32(0)
	v52 = v45
	goto L21
L19:
	;
	v82 = v44
	goto L20
L20:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strtitle[0]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v91*int32(28))+uint32(_c_F_pg_strtitle[1])))
	goto L34
L21:
	;
	if v50 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v82 = v75
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38+v49<<(uint(int32(2))%32)))) = v63
	if base.Ui32(int32(10)) <= base.Ui32(v63-int32(48)) {
		goto L30
	} else {
		goto L31
	}
L24:
	;
	v60 = F_casemap(m, v52, int32(0))
	mBase = m.M
	goto L27
L25:
	;
	goto L26
L26:
	;
	v62 = F_casemap(m, v52, int32(1))
	mBase = m.M
	goto L28
L27:
	;
	v63 = v60
	goto L23
L28:
	;
	v63 = v62
	goto L23
L29:
	;
	v75 = v49 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v38+v75<<(uint(int32(2))%32))))
	if v79 != 0 {
		v49 = v75
		v50 = v73
		v52 = v79
		goto L21
	} else {
		goto L33
	}
L30:
	;
	v69 = F_iswalpha(m, v63)
	mBase = m.M
	v73 = base.B2i32(v69 != int32(0))
	goto L32
L31:
	;
	v73 = int32(1)
	goto L32
L32:
	;
	goto L29
L33:
	;
	goto L22
L34:
	;
	v99 = v96*v82 + int32(1)
	v100 = F_palloc(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	if v99 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strtitle[2]))
	if v102 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v306 = v6
	goto L38
L38:
	;
	if base.Ui32(v306+int32(1)) <= base.Ui32(l1) {
		goto L102
	} else {
		goto L103
	}
L39:
	;
	v116 = int32(0)
	v122 = m.G0
	v123 = int32(16)
	v124 = v122 - v123
	m.G0 = v124
	*(*int32)(unsafe.Add(mBase, uint32(v124)+12)) = v38
	v128 = v124 + int32(12)
	v129 = m.G0
	v131 = v129 - v123
	m.G0 = v131
	if v100 != 0 {
		goto L54
	} else {
		goto L55
	}
L40:
	;
	if v102 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	if v105 == int32(_a_F_pg_strtitle_0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v110 = int32(_a_F_pg_strtitle_0)
	goto L45
L44:
	;
	v110 = v102
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_strtitle[2])) = v110
	goto L42
L46:
	;
	v115 = int32(-1)
	goto L48
L47:
	;
	v115 = v105
	goto L48
L48:
	;
	goto L39
L49:
	;
	if v115 != 0 {
		goto L93
	} else {
		goto L94
	}
L50:
	;
	v287 = int32(16)
	m.G0 = v131 + v287
	m.G0 = v124 + v287
	goto L49
L51:
	;
	v283 = v99 - v268
	goto L50
L52:
	;
	if v205 != 0 {
		goto L77
	} else {
		goto L78
	}
L53:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v164 = v99
	v165 = v100
	v169 = v163
	goto L66
L54:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v99) {
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v136 == int32(0) {
		v283 = v116
		goto L50
	} else {
		goto L58
	}
L57:
	;
	v205 = v99
	v206 = v100
	goto L52
L58:
	;
	v139 = v136
	v140 = v135
	v142 = v116
	goto L59
L59:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v139) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v283 = v162
	goto L50
L61:
	;
	v151 = int32(-1)
	v154 = F_wcrtomb(m, v131+int32(12), v139)
	mBase = m.M
	if v154 == v151 {
		v283 = v151
		goto L50
	} else {
		goto L64
	}
L62:
	;
	v157 = int32(1)
	goto L63
L63:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	v162 = v157 + v142
	if v159 != 0 {
		v139 = v159
		v140 = v140 + int32(4)
		v142 = v162
		goto L59
	} else {
		goto L65
	}
L64:
	;
	v157 = v154
	goto L63
L65:
	;
	goto L60
L66:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if base.Ui32(v173-int32(128)) <= base.Ui32(int32(-128)) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v205 = v195
	v206 = v198
	goto L52
L68:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v201 = v199 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v201
	if base.Ui32(int32(3)) < base.Ui32(v195) {
		v164 = v195
		v165 = v198
		v169 = v201
		goto L66
	} else {
		goto L76
	}
L69:
	;
	if v173 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v173)
	v191 = int32(1)
	v195 = v164 - v191
	v198 = v165 + v191
	goto L68
L72:
	;
	v180 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v165))) = uint8(v180)
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v180
	v268 = v164
	goto L51
L73:
	;
	goto L74
L74:
	;
	v184 = int32(-1)
	v185 = F_wcrtomb(m, v165, v173)
	mBase = m.M
	if v185 == v184 {
		v283 = v184
		goto L50
	} else {
		goto L75
	}
L75:
	;
	v195 = v164 - v185
	v198 = v165 + v185
	goto L68
L76:
	;
	goto L67
L77:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v215 = v205
	v216 = v206
	v218 = v214
	goto L80
L78:
	;
	goto L79
L79:
	;
	v283 = v99
	goto L50
L80:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	if base.Ui32(v224-int32(128)) <= base.Ui32(int32(-128)) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L79
L82:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
	v257 = v255 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v257
	if v251 != 0 {
		v215 = v251
		v216 = v254
		v218 = v257
		goto L80
	} else {
		goto L91
	}
L83:
	;
	if v224 == int32(0) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v224)
	v247 = int32(1)
	v251 = v215 - v247
	v254 = v216 + v247
	goto L82
L86:
	;
	v231 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v216))) = uint8(v231)
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = v231
	v268 = v215
	goto L51
L87:
	;
	goto L88
L88:
	;
	v235 = int32(-1)
	v238 = F_wcrtomb(m, v131+int32(12), v224)
	mBase = m.M
	if v238 == v235 {
		v283 = v235
		goto L50
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(v215) < base.Ui32(v238) {
		v268 = v215
		goto L51
	} else {
		goto L90
	}
L90:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v218)))
	v243 = F_wcrtomb(m, v216, v242)
	mBase = m.M
	v251 = v215 - v238
	v254 = v216 + v238
	goto L82
L91:
	;
	goto L81
L92:
	;
	v306 = v283
	goto L38
L93:
	;
	if v115 == int32(-1) {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	goto L99
L96:
	;
	v300 = int32(_a_F_pg_strtitle_0)
	goto L98
L97:
	;
	v300 = v115
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_strtitle[2])) = v300
	goto L95
L99:
	;
	goto L101
L101:
	;
	goto L92
L102:
	;
	if v306 != 0 {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L104
L104:
	;
	F_pfree(m, v38)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L15
	} else {
		goto L108
	}
L105:
	;
	base.MemoryCopy(m, l0, v100, v306)
	goto L107
L106:
	;
	goto L107
L107:
	;
	v312 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v306))) = uint8(v312)
	goto L104
L108:
	;
	F_pfree(m, v100)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L15
	} else {
		goto L109
	}
L109:
	;
	v431 = v306
	goto L5
L110:
	;
	v320 = F_strlen(m, l2)
	mBase = m.M
	v321 = v320
	goto L112
L111:
	;
	v321 = l3
	goto L112
L112:
	;
	if base.Ui32(l1) < base.Ui32(v321+int32(1)) {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v431 = v321
	goto L5
L114:
	;
	if v321 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	base.MemoryCopy(m, l0, l2, v321)
	goto L117
L116:
	;
	goto L117
L117:
	;
	v328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v321))) = uint8(v328)
	v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v330 == v328 {
		goto L113
	} else {
		goto L118
	}
L118:
	;
	v333 = l0
	v334 = v330
	v338 = v6
	goto L119
L119:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v342 == int32(1) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L113
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v333))) = uint8(v379)
	v384 = v379 & int32(255)
	goto L147
L122:
	;
	if v338 != 0 {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	if v338 != 0 {
		goto L136
	} else {
		goto L137
	}
L125:
	;
	if base.Ui32((v334-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	goto L127
L127:
	;
	if base.Ui32((v334-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L133
	} else {
		goto L134
	}
L128:
	;
	v379 = v353
	goto L121
L129:
	;
	v353 = v334 | int32(32)
	goto L131
L130:
	;
	v353 = v334
	goto L131
L131:
	;
	goto L128
L132:
	;
	v379 = v362 & int32(255)
	goto L121
L133:
	;
	v362 = v334 - int32(32)
	goto L135
L134:
	;
	v362 = v334
	goto L135
L135:
	;
	goto L132
L136:
	;
	if base.Ui32(v334-int32(65)) < base.Ui32(int32(26)) {
		goto L140
	} else {
		goto L141
	}
L137:
	;
	goto L138
L138:
	;
	if base.Ui32(v334-int32(97)) < base.Ui32(int32(26)) {
		goto L144
	} else {
		goto L145
	}
L139:
	;
	v379 = v371
	goto L121
L140:
	;
	v371 = v334 | int32(32)
	goto L142
L141:
	;
	v371 = v334
	goto L142
L142:
	;
	goto L139
L143:
	;
	v379 = v378
	goto L121
L144:
	;
	v378 = v334 & int32(95)
	goto L146
L145:
	;
	v378 = v334
	goto L146
L146:
	;
	goto L143
L147:
	;
	v396 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v333)+1)))
	if v396 != 0 {
		v333 = v333 + int32(1)
		v334 = v396
		v338 = base.B2i32(base.Ui32(v384-int32(48)) < base.Ui32(int32(10))) | base.B2i32(base.Ui32(v384|int32(32)-int32(97)) < base.Ui32(int32(26)))
		goto L119
	} else {
		goto L148
	}
L148:
	;
	goto L120
L149:
	;
	F_errcode(m, int32(_a_F_pg_strtitle_1))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L15
	} else {
		goto L150
	}
L150:
	;
	F_errmsg(m, int32(_a_F_pg_strtitle_2), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L15
	} else {
		goto L151
	}
L151:
	;
	F_errfinish(m, int32(_a_F_pg_strtitle_3), int32(302), int32(_a_F_pg_strtitle_4))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L15
	} else {
		goto L152
	}
L152:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L153:
	;
	v436 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4))))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v436
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_pg_strtitle_5)
	F_errmsg_internal(m, int32(_a_F_pg_strtitle_6), v12)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L15
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_pg_strtitle_7), int32(1303), int32(_a_F_pg_strtitle_5))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L15
	} else {
		goto L155
	}
L155:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L156:
	;
	m.G0 = v450 + int32(16)
	v480 = v466
	goto L1
}
func F_pg_strupper(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
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
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v253 int32
	_ = v253
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v419 int32
	_ = v419
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	switch v13 - int32(98) {
	case 0:
		goto L2
	case 1:
		goto L4
	default:
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v419
L2:
	;
	v406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+16)))
	v407 = int32(0)
	v409 = F_convert_case(m, l0, l1, l2, l3, int32(2), v406, v407, v407)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L15
	} else {
		goto L132
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L15
	} else {
		goto L129
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strupper[0]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18*int32(28))+uint32(_c_F_pg_strupper[1])))
	goto L7
L5:
	;
	v419 = v388
	goto L1
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L15
	} else {
		goto L125
	}
L7:
	;
	if int32(2) <= v23 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l3 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	if l3 < int32(0) {
		goto L101
	} else {
		goto L102
	}
L11:
	;
	v29 = F_strlen(m, l2)
	mBase = m.M
	v30 = v29
	goto L13
L12:
	;
	v30 = l3
	goto L13
L13:
	;
	v32 = v30 + int32(1)
	if base.Ui32(int32(536870912)) <= base.Ui32(v32) {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v37 = F_palloc(m, v32<<(uint(int32(2))%32))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	F_char2wchar(m, v37, v32, l2, v30, l4)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v45 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v48 = v43
	v49 = v45
	goto L21
L19:
	;
	v68 = v43
	goto L20
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strupper[0]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76*int32(28))+uint32(_c_F_pg_strupper[1])))
	goto L25
L21:
	;
	v58 = F_casemap(m, v49, int32(1))
	mBase = m.M
	goto L23
L22:
	;
	v68 = v61
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37+v48<<(uint(int32(2))%32)))) = v58
	v61 = v48 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v37+v61<<(uint(int32(2))%32))))
	if v65 != 0 {
		v48 = v61
		v49 = v65
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v84 = v81*v68 + int32(1)
	v85 = F_palloc(m, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L15
	} else {
		goto L26
	}
L26:
	;
	if v84 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_pg_strupper[2]))
	if v87 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v291 = v43
	goto L29
L29:
	;
	if base.Ui32(v291+int32(1)) <= base.Ui32(l1) {
		goto L93
	} else {
		goto L94
	}
L30:
	;
	v101 = int32(0)
	v107 = m.G0
	v108 = int32(16)
	v109 = v107 - v108
	m.G0 = v109
	*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v37
	v113 = v109 + int32(12)
	v114 = m.G0
	v116 = v114 - v108
	m.G0 = v116
	if v85 != 0 {
		goto L45
	} else {
		goto L46
	}
L31:
	;
	if v87 == int32(-1) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v90 == int32(_a_F_pg_strupper_0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v95 = int32(_a_F_pg_strupper_0)
	goto L36
L35:
	;
	v95 = v87
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_strupper[2])) = v95
	goto L33
L37:
	;
	v100 = int32(-1)
	goto L39
L38:
	;
	v100 = v90
	goto L39
L39:
	;
	goto L30
L40:
	;
	if v100 != 0 {
		goto L84
	} else {
		goto L85
	}
L41:
	;
	v272 = int32(16)
	m.G0 = v116 + v272
	m.G0 = v109 + v272
	goto L40
L42:
	;
	v268 = v84 - v253
	goto L41
L43:
	;
	if v190 != 0 {
		goto L68
	} else {
		goto L69
	}
L44:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v149 = v84
	v150 = v85
	v154 = v148
	goto L57
L45:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v84) {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v121 == int32(0) {
		v268 = v101
		goto L41
	} else {
		goto L49
	}
L48:
	;
	v190 = v84
	v191 = v85
	goto L43
L49:
	;
	v124 = v121
	v125 = v120
	v127 = v101
	goto L50
L50:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v124) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v268 = v147
	goto L41
L52:
	;
	v136 = int32(-1)
	v139 = F_wcrtomb(m, v116+int32(12), v124)
	mBase = m.M
	if v139 == v136 {
		v268 = v136
		goto L41
	} else {
		goto L55
	}
L53:
	;
	v142 = int32(1)
	goto L54
L54:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	v147 = v142 + v127
	if v144 != 0 {
		v124 = v144
		v125 = v125 + int32(4)
		v127 = v147
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v142 = v139
	goto L54
L56:
	;
	goto L51
L57:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	if base.Ui32(v158-int32(128)) <= base.Ui32(int32(-128)) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v190 = v180
	v191 = v183
	goto L43
L59:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v186 = v184 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v186
	if base.Ui32(int32(3)) < base.Ui32(v180) {
		v149 = v180
		v150 = v183
		v154 = v186
		goto L57
	} else {
		goto L67
	}
L60:
	;
	if v158 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v158)
	v176 = int32(1)
	v180 = v149 - v176
	v183 = v150 + v176
	goto L59
L63:
	;
	v165 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v165)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v165
	v253 = v149
	goto L42
L64:
	;
	goto L65
L65:
	;
	v169 = int32(-1)
	v170 = F_wcrtomb(m, v150, v158)
	mBase = m.M
	if v170 == v169 {
		v268 = v169
		goto L41
	} else {
		goto L66
	}
L66:
	;
	v180 = v149 - v170
	v183 = v150 + v170
	goto L59
L67:
	;
	goto L58
L68:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v200 = v190
	v201 = v191
	v203 = v199
	goto L71
L69:
	;
	goto L70
L70:
	;
	v268 = v84
	goto L41
L71:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if base.Ui32(v209-int32(128)) <= base.Ui32(int32(-128)) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L70
L73:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v242 = v240 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v242
	if v236 != 0 {
		v200 = v236
		v201 = v239
		v203 = v242
		goto L71
	} else {
		goto L82
	}
L74:
	;
	if v209 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v209)
	v232 = int32(1)
	v236 = v200 - v232
	v239 = v201 + v232
	goto L73
L77:
	;
	v216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v201))) = uint8(v216)
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v216
	v253 = v200
	goto L42
L78:
	;
	goto L79
L79:
	;
	v220 = int32(-1)
	v223 = F_wcrtomb(m, v116+int32(12), v209)
	mBase = m.M
	if v223 == v220 {
		v268 = v220
		goto L41
	} else {
		goto L80
	}
L80:
	;
	if base.Ui32(v200) < base.Ui32(v223) {
		v253 = v200
		goto L42
	} else {
		goto L81
	}
L81:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v228 = F_wcrtomb(m, v201, v227)
	mBase = m.M
	v236 = v200 - v223
	v239 = v201 + v223
	goto L73
L82:
	;
	goto L72
L83:
	;
	v291 = v268
	goto L29
L84:
	;
	if v100 == int32(-1) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	goto L90
L87:
	;
	v285 = int32(_a_F_pg_strupper_0)
	goto L89
L88:
	;
	v285 = v100
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_pg_strupper[2])) = v285
	goto L86
L90:
	;
	goto L92
L92:
	;
	goto L83
L93:
	;
	if v291 != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	F_pfree(m, v37)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L15
	} else {
		goto L99
	}
L96:
	;
	base.MemoryCopy(m, l0, v85, v291)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v297 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v291))) = uint8(v297)
	goto L95
L99:
	;
	F_pfree(m, v85)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L15
	} else {
		goto L100
	}
L100:
	;
	v388 = v291
	goto L5
L101:
	;
	v305 = F_strlen(m, l2)
	mBase = m.M
	v306 = v305
	goto L103
L102:
	;
	v306 = l3
	goto L103
L103:
	;
	if base.Ui32(l1) < base.Ui32(v306+int32(1)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v388 = v306
	goto L5
L105:
	;
	if v306 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	base.MemoryCopy(m, l0, l2, v306)
	goto L108
L107:
	;
	goto L108
L108:
	;
	v313 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v306))) = uint8(v313)
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v315 == v313 {
		goto L104
	} else {
		goto L109
	}
L109:
	;
	v318 = l0
	v320 = v315
	goto L110
L110:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+4)))
	if v326 == int32(1) {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L104
L112:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v318))) = uint8(v351)
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v318)+1)))
	if v353 != 0 {
		v318 = v318 + int32(1)
		v320 = v353
		goto L110
	} else {
		goto L124
	}
L113:
	;
	v329 = int32(255)
	v330 = v320 & v329
	if base.Ui32((v330-int32(97))&v329) < base.Ui32(int32(26)) {
		goto L117
	} else {
		goto L118
	}
L114:
	;
	goto L115
L115:
	;
	v343 = v320 & int32(255)
	if base.Ui32(v343-int32(97)) < base.Ui32(int32(26)) {
		goto L121
	} else {
		goto L122
	}
L116:
	;
	v351 = v339 & int32(255)
	goto L112
L117:
	;
	v339 = v330 - int32(32)
	goto L119
L118:
	;
	v339 = v330
	goto L119
L119:
	;
	goto L116
L120:
	;
	v351 = v350
	goto L112
L121:
	;
	v350 = v343 & int32(95)
	goto L123
L122:
	;
	v350 = v343
	goto L123
L123:
	;
	goto L120
L124:
	;
	goto L111
L125:
	;
	F_errcode(m, int32(_a_F_pg_strupper_1))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L15
	} else {
		goto L126
	}
L126:
	;
	F_errmsg(m, int32(_a_F_pg_strupper_2), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L15
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_pg_strupper_3), int32(390), int32(_a_F_pg_strupper_4))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L15
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	v393 = int32(*(*int8)(unsafe.Add(mBase, uint32(l4))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_pg_strupper_5)
	F_errmsg_internal(m, int32(_a_F_pg_strupper_6), v11)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L15
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_pg_strupper_7), int32(1322), int32(_a_F_pg_strupper_5))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L15
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	v419 = v409
	goto L1
}
func F_pg_timezone_abbrevs_abbrevs(m *base.Module, l0 int32) int32 {
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
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
	var v144 int32
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
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int64
	_ = v317
	var v318 int64
	_ = v318
	var v321 int64
	_ = v321
	var v327 int32
	_ = v327
	var v329 int64
	_ = v329
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
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
	var v350 int64
	_ = v350
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+82)) = uint8(v2)
	*(*uint16)(unsafe.Add(mBase, uint32(v11)+80)) = uint16(v2)
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
	v372 = m.ExcPending
	if v372 != 0 {
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
	return int32(0)
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
	v39 = F_get_call_result_type(m, l0, v33, v11+int32(40))
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
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
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
	m.G0 = v11 + int32(96)
	return v359
L12:
	;
	v70 = v53 + v55<<(uint(int32(4))%32)
	v72 = v70 + int32(8)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+19)))
	switch v73 - int32(5) {
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
	v359 = int32(0)
	goto L11
L18:
	;
	v146 = v11 + int32(69)
	goto L40
L19:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v143 = v140
	v144 = int32(0)
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
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v79 = v53 + v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v80 != 0 {
		v113 = v80
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)+12))
	v143 = v76
	v144 = int32(1)
	goto L18
L23:
	;
	v117 = *(*int64)(unsafe.Add(mBase, _c_F_pg_timezone_abbrevs_abbrevs[2]))
	v120 = F_DetermineTimeZoneAbbrevOffsetTS(m, v117, v72, v113, v11+int32(40))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L33
	}
L24:
	;
	v82 = v79 + int32(4)
	v83 = F_pg_tzset(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v83
	if v83 != 0 {
		v113 = v83
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v86 = int32(0)
	v88 = F_errsave_start(m, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	if v88 == int32(0) {
		v113 = v86
		goto L23
	} else {
		goto L28
	}
L28:
	;
	F_errcode(m, int32(22))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v82
	F_errmsg(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_1), v11+int32(32))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v72
	F_errdetail(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_2), v11+int32(16))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	F_errsave_finish(m, int32(0), int32(_a_F_pg_timezone_abbrevs_abbrevs_3), int32(_a_F_pg_timezone_abbrevs_abbrevs_4), int32(_a_F_pg_timezone_abbrevs_abbrevs_5))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v113 = v86
	goto L23
L33:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v143 = int32(0) - v120
	v144 = base.B2i32(v123 != int32(0))
	goto L18
L34:
	;
	v130 = int32(*(*int8)(unsafe.Add(mBase, uint32(v72)+11)))
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
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+69)))
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
	if (v146^v72)&int32(3) != 0 {
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
	v225 = v72
	v226 = v146
	v227 = v153
	goto L44
L46:
	;
	goto L47
L47:
	;
	v157 = int32(0)
	if base.B2i32(v72&int32(3) == v157)|int32(0) == v157 {
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
	v169 = v72
	v170 = v146
	v171 = v153
	goto L52
L50:
	;
	goto L51
L51:
	;
	v190 = v72
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
	v300 = F_cstring_to_text(m, v11+int32(69))
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v300
	*(*int64)(unsafe.Add(mBase, uint32(v11)+56)) = v302
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = base.I64_extend_i32_s(v143) * int64(1000000)
	v312 = v11 + int32(40)
	v314 = F_palloc(m, int32(16))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L5
	} else {
		goto L79
	}
L79:
	;
	v317 = int64(*(*int32)(unsafe.Add(mBase, uint32(v312)+12)))
	v318 = int64(*(*int32)(unsafe.Add(mBase, uint32(v312)+16)))
	v321 = v317 + v318*int64(12)
	if base.Ui64(int64(-4294967296)) <= base.Ui64(v321-int64(2147483648)) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+92)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = v314
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v336 + int32(1)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v51)+28))
	v345 = F_heap_form_tuple(m, v340, v11+int32(84), v11+int32(80))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L5
	} else {
		goto L84
	}
L81:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v314)+12)) = uint32(v321)
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v314)+8)) = v327
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
	*(*int64)(unsafe.Add(mBase, uint32(v314))) = v329
	goto L83
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
	v348 = F_HeapTupleHeaderGetDatum(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L5
	} else {
		goto L85
	}
L85:
	;
	v350 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
	*(*int64)(unsafe.Add(mBase, uint32(v51))) = v350 + int64(1)
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v354)+20)) = int32(1)
	v359 = v348
	goto L11
L86:
	;
	F_errmsg_internal(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_9), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L5
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(_a_F_pg_timezone_abbrevs_abbrevs_3), int32(_a_F_pg_timezone_abbrevs_abbrevs_10), int32(_a_F_pg_timezone_abbrevs_abbrevs_8))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
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
func F_pg_total_relation_size(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_try_relation_open(m, v5, int32(1))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		if v7 == int32(0) {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			return int32(0)
		} else {
			v17 = F_calculate_table_size(m, v7)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v19 = F_calculate_indexes_size(m, v7)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					F_relation_close(m, v7, int32(1))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						v25 = F_Int64GetDatum(m, v17+v19)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return v25
						}
					}
				}
			}
		}
	}
}
func F_pg_typeof(m *base.Module, l0 int32) int32 {
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
		return int32(0)
	} else {
		return v4
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
func F_pg_xact_commit_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v20 int64
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_TransactionIdGetCommitTsData(m, v8, v6+int32(8), v2)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 == int32(0) {
			v18 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v18)
			v23 = v2
			m.G0 = v6 + int32(16)
			return v23
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, uint32(v6)+8))
			v21 = F_Int64GetDatum(m, v20)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = v21
				m.G0 = v6 + int32(16)
				return v23
			}
		}
	}
}
