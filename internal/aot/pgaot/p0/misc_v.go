package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ValidateRestrictionEstimator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(98784250089)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = int64(111669151977)
	v17 = F_LookupFuncName(m, l0, int32(4), v7+int32(16), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = F_get_func_rettype(m, v17)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 == int32(701) {
				if base.Ui32(int32(10000)) <= base.Ui32(v17) {
					v27 = F_superuser(m)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						if v27 != 0 {
							m.G0 = v7 + int32(32)
							return v17
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(264332), int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(517112), int32(309), int32(220117))
										mBase = m.M
										v44 = m.ExcPending
										if v44 != 0 {
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
					v47 = *(*int32)(unsafe.Add(mBase, _consts[4]))
					v49 = F_object_aclcheck(m, int32(1255), v17, v47, int64(128))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 == int32(0) {
							m.G0 = v7 + int32(32)
							return v17
						} else {
							v54 = F_NameListToString(m, l0)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_aclcheck_error(m, v49, int32(19), v54)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									m.G0 = v7 + int32(32)
									return v17
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(117833860))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v70 = F_NameListToString(m, l0)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(580526)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v70
							F_errmsg(m, int32(201250), v7)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(517112), int32(292), int32(220117))
								mBase = m.M
								v82 = m.ExcPending
								if v82 != 0 {
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
func F_validateConnectbyTupleDesc(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v184 int32
	_ = v184
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if l1 != 0 {
		v12 = int32(4)
	} else {
		v12 = int32(3)
	}
	v13 = v12 + l2
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v13 == v14 {
		v18 = l0 + v13<<(uint(int32(4))%32)
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+288))
		if v19 != int32(23) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v101 = m.ExcPending
				if v101 != 0 {
					return
				} else {
					F_errmsg(m, int32(385935), int32(0))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return
					} else {
						v109 = F_format_type_be(m, int32(23))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v109
							F_errdetail(m, int32(632706), v8+int32(48))
							mBase = m.M
							v117 = m.ExcPending
							if v117 != 0 {
								return
							} else {
								F_errfinish(m, int32(523797), int32(1434), int32(510222))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
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
			if l1 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+388))
				if v22 != int32(25) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v131 = m.ExcPending
						if v131 != 0 {
							return
						} else {
							F_errmsg(m, int32(385935), int32(0))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return
							} else {
								v139 = F_format_type_be(m, int32(25))
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v139
									F_errdetail(m, int32(632751), v8+int32(16))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return
									} else {
										F_errfinish(m, int32(523797), int32(1442), int32(510222))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
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
					if l2 == int32(0) {
						m.G0 = v8 + int32(80)
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v18)+488))
						if v27 == int32(23) {
							m.G0 = v8 + int32(80)
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return
								} else {
									F_errmsg(m, int32(385935), int32(0))
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return
									} else {
										v44 = F_format_type_be(m, int32(23))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v44
											F_errdetail(m, int32(632660), v8)
											mBase = m.M
											v50 = m.ExcPending
											if v50 != 0 {
												return
											} else {
												F_errfinish(m, int32(523797), int32(1451), int32(510222))
												mBase = m.M
												v57 = m.ExcPending
												if v57 != 0 {
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
			} else {
				if l2 == int32(0) {
					m.G0 = v8 + int32(80)
					return
				} else {
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v18)+388))
					if v60 != int32(23) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v158 = m.ExcPending
						if v158 != 0 {
							return
						} else {
							F_errcode(m, int32(67141764))
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return
							} else {
								F_errmsg(m, int32(385935), int32(0))
								mBase = m.M
								v167 = m.ExcPending
								if v167 != 0 {
									return
								} else {
									v169 = F_format_type_be(m, int32(23))
									mBase = m.M
									v170 = m.ExcPending
									if v170 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v169
										F_errdetail(m, int32(632613), v8+int32(32))
										mBase = m.M
										v177 = m.ExcPending
										if v177 != 0 {
											return
										} else {
											F_errfinish(m, int32(523797), int32(1458), int32(510222))
											mBase = m.M
											v184 = m.ExcPending
											if v184 != 0 {
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
						m.G0 = v8 + int32(80)
						return
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v69 = m.ExcPending
		if v69 != 0 {
			return
		} else {
			F_errcode(m, int32(67141764))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				F_errmsg(m, int32(385935), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v79
					*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v13
					F_errdetail(m, int32(680420), v8-int32(-64))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						F_errfinish(m, int32(523797), int32(1424), int32(510222))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
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
func F_varlenafastcmp_locale(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
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
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v14 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = F_pg_detoast_datum_packed(m, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(4)
			v22 = int32(1)
			v23 = v18 + v22
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			v30 = v14 + v22
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v33 = v31 & v22
			if v31 == v22 {
				v36 = int32(4)
				v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
				if v38&int32(254) == int32(2) {
					v47 = v36
				} else {
					v47 = base.B2i32(v38 == int32(18)) << (uint(v36) % 32)
				}
				if v38 == int32(1) {
					v50 = v36
				} else {
					v50 = v47
				}
				v61 = v50
			} else {
				v51 = int32(1)
				if v33 != 0 {
					v61 = int32(base.Ui32(v31)>>(uint(v51)%32)) - v51
				} else {
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v61 = int32(base.Ui32(v55)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v33 != 0 {
				v62 = v30
			} else {
				v62 = v14 + v20
			}
			if v24&v22 != 0 {
				v63 = v23
			} else {
				v63 = v18 + v20
			}
			if v24 == int32(1) {
				v66 = int32(4)
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				if v68&int32(254) == int32(2) {
					v77 = v66
				} else {
					v77 = base.B2i32(v68 == int32(18)) << (uint(v66) % 32)
				}
				if v68 == int32(1) {
					v80 = v66
				} else {
					v80 = v77
				}
				v93 = v80
			} else {
				v81 = int32(1)
				if v24&v81 != 0 {
					v93 = int32(base.Ui32(v24)>>(uint(v81)%32)) - v81
				} else {
					v87 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v93 = int32(base.Ui32(v87)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v94 = F_varstrfastcmp_locale(m, v62, v61, v63, v93, l2)
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				if l0 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int32(0)
					} else {
						if l1 != v18 {
							F_pfree(m, v18)
							mBase = m.M
							v101 = m.ExcPending
							if v101 != 0 {
								return int32(0)
							} else {
								return v94
							}
						} else {
							return v94
						}
					}
				} else {
					if l1 != v18 {
						F_pfree(m, v18)
						mBase = m.M
						v101 = m.ExcPending
						if v101 != 0 {
							return int32(0)
						} else {
							return v94
						}
					} else {
						return v94
					}
				}
			}
		}
	}
}
func F_varstr_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v36 float64
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v45 float64
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 float64
	_ = v48
	var v51 int32
	_ = v51
	var v52 float64
	_ = v52
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 float64
	_ = v64
	var v66 int32
	_ = v66
	var v73 float64
	_ = v73
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v176 float64
	_ = v176
	var v178 float64
	_ = v178
	var v180 float64
	_ = v180
	var v186 float64
	_ = v186
	var v203 float64
	_ = v203
	var v207 float64
	_ = v207
	var v226 float64
	_ = v226
	var v229 float64
	_ = v229
	var v232 int32
	_ = v232
	var v233 float64
	_ = v233
	var v235 int32
	_ = v235
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v251 float64
	_ = v251
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v260 float64
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 float64
	_ = v263
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v272 float64
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v279 float64
	_ = v279
	var v281 int32
	_ = v281
	var v288 float64
	_ = v288
	var v291 int32
	_ = v291
	var v292 float64
	_ = v292
	var v295 float64
	_ = v295
	var v296 float64
	_ = v296
	var v297 float64
	_ = v297
	var v298 float64
	_ = v298
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v391 float64
	_ = v391
	var v393 float64
	_ = v393
	var v395 float64
	_ = v395
	var v401 float64
	_ = v401
	var v418 float64
	_ = v418
	var v422 float64
	_ = v422
	var v441 float64
	_ = v441
	var v444 float64
	_ = v444
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v459 float64
	_ = v459
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v477 float64
	_ = v477
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 float64
	_ = v496
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	if l0 < int32(100) {
		v510 = v3
		m.G0 = v10 + int32(80)
		return v510
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v17 = v15 + int32(40)
		v18 = float64(0)
		v20 = int32(0)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		if v27 != 0 {
			v28 = int32(1)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			if v27 == v28 {
				v64 = v18
				v66 = v20
			} else {
				v36 = v18
				v38 = v20
				v40 = v20
				for {
					v45 = float64(1)
					v46 = v38 + v30
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
					v48 = F_ldexp(m, v45, v47)
					mBase = m.M
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
					v52 = F_ldexp(m, v45, v51)
					mBase = m.M
					v57 = base.F64_add(base.F64_add(v36, base.F64_div(v45, v52)), base.F64_div(v45, v48))
					v58 = int32(2)
					v59 = v38 + v58
					v61 = v40 + v58
					if v61 != v27&int32(-2) {
						v36 = v57
						v38 = v59
						v40 = v61
						continue
					} else {
						break
					}
					break
				}
				v64 = v57
				v66 = v59
			}
			if v27&v28 != 0 {
				v73 = float64(1)
				v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66+v30))))
				v77 = F_ldexp(m, v73, v76)
				mBase = m.M
				v80 = base.F64_add(v64, base.F64_div(v73, v77))
			} else {
				v80 = v64
			}
			v81 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
			v82 = base.F64_div(v81, v80)
			v83 = base.F64_convert_i32_u(v27)
			if base.F64_le(v82, base.F64_mul(v83, float64(2.5))) == int32(0) {
				v186 = v82
				if base.F64_gt(v186, float64(1.4316557653333333e+08)) == int32(0) {
					v207 = v186
				} else {
					v203 = F_log(m, base.F64_add(base.F64_mul(v186, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v207 = base.F64_mul(v203, float64(-4.294967296e+09))
				}
				v226 = v207
			} else {
				v90 = v27 & int32(3)
				v91 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v92 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v27) {
					v103 = v92
					v104 = int32(0)
					v105 = v92
					for {
						v110 = v103 + v91
						v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
						v112 = int32(0)
						v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
						v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+2)))
						v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+3)))
						v126 = v105 + base.B2i32(v111 == v112) + base.B2i32(v115 == v112) + base.B2i32(v119 == v112) + base.B2i32(v123 == v112)
						v127 = int32(4)
						v128 = v103 + v127
						v130 = v104 + v127
						if v130 != v27&int32(-4) {
							v103 = v128
							v104 = v130
							v105 = v126
							continue
						} else {
							break
						}
						break
					}
					v135 = v128
					v137 = v126
				} else {
					v135 = v92
					v137 = v92
				}
				if v90 != 0 {
					v145 = v135
					v147 = v137
					v148 = v92
					for {
						v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v91))))
						v156 = v147 + base.B2i32(v153 == int32(0))
						v157 = int32(1)
						v160 = v148 + v157
						if v160 != v90 {
							v145 = v145 + v157
							v147 = v156
							v148 = v160
							continue
						} else {
							break
						}
						break
					}
					v167 = v156
				} else {
					v167 = v137
				}
				if v167 == int32(0) {
					v207 = v82
					v226 = v207
				} else {
					v176 = F_log(m, base.F64_div(v83, base.F64_convert_i32_s(v167)))
					mBase = m.M
					v226 = base.F64_mul(v176, v83)
				}
			}
		} else {
			v178 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
			v180 = base.F64_div(v178, float64(0))
			if base.F64_le(v180, base.F64_mul(base.F64_convert_i32_u(v27), float64(2.5))) != 0 {
				v207 = v180
			} else {
				v186 = v180
				if base.F64_gt(v186, float64(1.4316557653333333e+08)) == int32(0) {
					v207 = v186
				} else {
					v203 = F_log(m, base.F64_add(base.F64_mul(v186, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v207 = base.F64_mul(v203, float64(-4.294967296e+09))
				}
			}
			v226 = v207
		}
		if base.F64_le(v226, float64(1)) != 0 {
			v229 = float64(1)
		} else {
			v229 = v226
		}
		v232 = v15 - int32(-64)
		v233 = float64(0)
		v235 = int32(0)
		v242 = *(*int32)(unsafe.Add(mBase, uint32(v232)+4))
		if v242 != 0 {
			v243 = int32(1)
			v245 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
			if v242 == v243 {
				v279 = v233
				v281 = v235
			} else {
				v251 = v233
				v253 = v235
				v255 = v235
				for {
					v260 = float64(1)
					v261 = v253 + v245
					v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
					v263 = F_ldexp(m, v260, v262)
					mBase = m.M
					v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
					v267 = F_ldexp(m, v260, v266)
					mBase = m.M
					v272 = base.F64_add(base.F64_add(v251, base.F64_div(v260, v267)), base.F64_div(v260, v263))
					v273 = int32(2)
					v274 = v253 + v273
					v276 = v255 + v273
					if v276 != v242&int32(-2) {
						v251 = v272
						v253 = v274
						v255 = v276
						continue
					} else {
						break
					}
					break
				}
				v279 = v272
				v281 = v274
			}
			if v242&v243 != 0 {
				v288 = float64(1)
				v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281+v245))))
				v292 = F_ldexp(m, v288, v291)
				mBase = m.M
				v295 = base.F64_add(v279, base.F64_div(v288, v292))
			} else {
				v295 = v279
			}
			v296 = *(*float64)(unsafe.Add(mBase, uint32(v232)+8))
			v297 = base.F64_div(v296, v295)
			v298 = base.F64_convert_i32_u(v242)
			if base.F64_le(v297, base.F64_mul(v298, float64(2.5))) == int32(0) {
				v401 = v297
				if base.F64_gt(v401, float64(1.4316557653333333e+08)) == int32(0) {
					v422 = v401
				} else {
					v418 = F_log(m, base.F64_add(base.F64_mul(v401, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v422 = base.F64_mul(v418, float64(-4.294967296e+09))
				}
				v441 = v422
			} else {
				v305 = v242 & int32(3)
				v306 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
				v307 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v242) {
					v318 = v307
					v319 = int32(0)
					v320 = v307
					for {
						v325 = v318 + v306
						v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325))))
						v327 = int32(0)
						v330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+1)))
						v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+2)))
						v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v325)+3)))
						v341 = v320 + base.B2i32(v326 == v327) + base.B2i32(v330 == v327) + base.B2i32(v334 == v327) + base.B2i32(v338 == v327)
						v342 = int32(4)
						v343 = v318 + v342
						v345 = v319 + v342
						if v345 != v242&int32(-4) {
							v318 = v343
							v319 = v345
							v320 = v341
							continue
						} else {
							break
						}
						break
					}
					v350 = v343
					v352 = v341
				} else {
					v350 = v307
					v352 = v307
				}
				if v305 != 0 {
					v360 = v350
					v362 = v352
					v363 = v307
					for {
						v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v360+v306))))
						v371 = v362 + base.B2i32(v368 == int32(0))
						v372 = int32(1)
						v375 = v363 + v372
						if v375 != v305 {
							v360 = v360 + v372
							v362 = v371
							v363 = v375
							continue
						} else {
							break
						}
						break
					}
					v382 = v371
				} else {
					v382 = v352
				}
				if v382 == int32(0) {
					v422 = v297
					v441 = v422
				} else {
					v391 = F_log(m, base.F64_div(v298, base.F64_convert_i32_s(v382)))
					mBase = m.M
					v441 = base.F64_mul(v391, v298)
				}
			}
		} else {
			v393 = *(*float64)(unsafe.Add(mBase, uint32(v232)+8))
			v395 = base.F64_div(v393, float64(0))
			if base.F64_le(v395, base.F64_mul(base.F64_convert_i32_u(v242), float64(2.5))) != 0 {
				v422 = v395
			} else {
				v401 = v395
				if base.F64_gt(v401, float64(1.4316557653333333e+08)) == int32(0) {
					v422 = v401
				} else {
					v418 = F_log(m, base.F64_add(base.F64_mul(v401, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v422 = base.F64_mul(v418, float64(-4.294967296e+09))
				}
			}
			v441 = v422
		}
		if base.F64_le(v441, float64(1)) != 0 {
			v444 = float64(1)
		} else {
			v444 = v441
		}
		v446 = int32(*(*uint8)(unsafe.Add(mBase, _consts[33])))
		if v446 != int32(1) {
			v477 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
			if base.F64_lt(base.F64_mul(v444, v477), v229) != 0 {
				if base.Ui32(l0) < base.Ui32(int32(10001)) {
					v510 = v3
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = base.F64_mul(v477, float64(0.65))
					v510 = v3
				}
				m.G0 = v10 + int32(80)
				return v510
			} else {
				v485 = int32(1)
				v487 = int32(*(*uint8)(unsafe.Add(mBase, _consts[33])))
				if v487 != v485 {
					v510 = v485
					m.G0 = v10 + int32(80)
					return v510
				} else {
					v492 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v493 = m.ExcPending
					if v493 != 0 {
						return int32(0)
					} else {
						if v492 == int32(0) {
							v510 = v485
							m.G0 = v10 + int32(80)
							return v510
						} else {
							v496 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
							*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v496
							*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v444
							*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v229
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
							F_errmsg_internal(m, int32(706227), v10)
							mBase = m.M
							v503 = m.ExcPending
							if v503 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(524313), int32(2650), int32(87198))
								mBase = m.M
								v508 = m.ExcPending
								if v508 != 0 {
									return int32(0)
								} else {
									v510 = v485
									m.G0 = v10 + int32(80)
									return v510
								}
							}
						}
					}
				}
			}
		} else {
			v451 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v454 = m.ExcPending
			if v454 != 0 {
				return int32(0)
			} else {
				if v451 == int32(0) {
					v477 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
					if base.F64_lt(base.F64_mul(v444, v477), v229) != 0 {
						if base.Ui32(l0) < base.Ui32(int32(10001)) {
							v510 = v3
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = base.F64_mul(v477, float64(0.65))
							v510 = v3
						}
						m.G0 = v10 + int32(80)
						return v510
					} else {
						v485 = int32(1)
						v487 = int32(*(*uint8)(unsafe.Add(mBase, _consts[33])))
						if v487 != v485 {
							v510 = v485
							m.G0 = v10 + int32(80)
							return v510
						} else {
							v492 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v493 = m.ExcPending
							if v493 != 0 {
								return int32(0)
							} else {
								if v492 == int32(0) {
									v510 = v485
									m.G0 = v10 + int32(80)
									return v510
								} else {
									v496 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
									*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v496
									*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v444
									*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v229
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
									F_errmsg_internal(m, int32(706227), v10)
									mBase = m.M
									v503 = m.ExcPending
									if v503 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(524313), int32(2650), int32(87198))
										mBase = m.M
										v508 = m.ExcPending
										if v508 != 0 {
											return int32(0)
										} else {
											v510 = v485
											m.G0 = v10 + int32(80)
											return v510
										}
									}
								}
							}
						}
					}
				} else {
					v459 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
					*(*float64)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = v459
					*(*float64)(unsafe.Add(mBase, uint32(v10)+48)) = v444
					*(*float64)(unsafe.Add(mBase, uint32(v10)+40)) = v229
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
					*(*float64)(unsafe.Add(mBase, uint32(v10)+56)) = base.F64_div(v229, base.F64_convert_i32_u(l0))
					F_errmsg_internal(m, int32(706324), v10+int32(32))
					mBase = m.M
					v471 = m.ExcPending
					if v471 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(524313), int32(2584), int32(87198))
						mBase = m.M
						v476 = m.ExcPending
						if v476 != 0 {
							return int32(0)
						} else {
							v477 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
							if base.F64_lt(base.F64_mul(v444, v477), v229) != 0 {
								if base.Ui32(l0) < base.Ui32(int32(10001)) {
									v510 = v3
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = base.F64_mul(v477, float64(0.65))
									v510 = v3
								}
								m.G0 = v10 + int32(80)
								return v510
							} else {
								v485 = int32(1)
								v487 = int32(*(*uint8)(unsafe.Add(mBase, _consts[33])))
								if v487 != v485 {
									v510 = v485
									m.G0 = v10 + int32(80)
									return v510
								} else {
									v492 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v493 = m.ExcPending
									if v493 != 0 {
										return int32(0)
									} else {
										if v492 == int32(0) {
											v510 = v485
											m.G0 = v10 + int32(80)
											return v510
										} else {
											v496 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
											*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v496
											*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v444
											*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v229
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
											F_errmsg_internal(m, int32(706227), v10)
											mBase = m.M
											v503 = m.ExcPending
											if v503 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(524313), int32(2650), int32(87198))
												mBase = m.M
												v508 = m.ExcPending
												if v508 != 0 {
													return int32(0)
												} else {
													v510 = v485
													m.G0 = v10 + int32(80)
													return v510
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
func F_varstrfastcmp_c(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	v13 = F_pg_detoast_datum_packed(m, l0)
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
	v17 = F_pg_detoast_datum_packed(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(4)
	v21 = int32(1)
	v22 = v17 + v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v29 = v13 + v21
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	v32 = v30 & v21
	if v30 == v21 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v32 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v35 = int32(4)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v37&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v50 = int32(1)
	if v32 != 0 {
		v60 = int32(base.Ui32(v30)>>(uint(v50)%32)) - v50
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v46 = v35
	goto L10
L9:
	;
	v46 = base.B2i32(v37 == int32(18)) << (uint(v35) % 32)
	goto L10
L10:
	;
	if v37 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v49 = v35
	goto L13
L12:
	;
	v49 = v46
	goto L13
L13:
	;
	v60 = v49
	goto L4
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v61 = v29
	goto L17
L16:
	;
	v61 = v13 + v19
	goto L17
L17:
	;
	if v23&v21 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v62 = v22
	goto L20
L19:
	;
	v62 = v17 + v19
	goto L20
L20:
	;
	if v23 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v93 = base.B2i32(v60 < v92)
	if v60 < v92 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v65 = int32(4)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v67&int32(254) == int32(2) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v80 = int32(1)
	if v23&v80 != 0 {
		v92 = int32(base.Ui32(v23)>>(uint(v80)%32)) - v80
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v76 = v65
	goto L27
L26:
	;
	v76 = base.B2i32(v67 == int32(18)) << (uint(v65) % 32)
	goto L27
L27:
	;
	if v67 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v79 = v65
	goto L30
L29:
	;
	v79 = v76
	goto L30
L30:
	;
	v92 = v79
	goto L21
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v92 = int32(base.Ui32(v86)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v94 = v60
	goto L34
L33:
	;
	v94 = v92
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v94) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if l0 != v13 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v156 = int32(0)
	goto L35
L37:
	;
	v130 = v125
	v131 = v126
	v132 = v127
	goto L47
L38:
	;
	if (v61|v62)&int32(3) != 0 {
		v125 = v61
		v126 = v62
		v127 = v94
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v118 = v61
	v119 = v62
	v120 = v94
	goto L40
L40:
	;
	if v120 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v102 = v61
	v103 = v62
	v104 = v94
	goto L42
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v107 != v108 {
		v125 = v102
		v126 = v103
		v127 = v104
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v118 = v113
	v119 = v111
	v120 = v115
	goto L40
L44:
	;
	v110 = int32(4)
	v111 = v103 + v110
	v113 = v102 + v110
	v115 = v104 - v110
	if base.Ui32(int32(3)) < base.Ui32(v115) {
		v102 = v113
		v103 = v111
		v104 = v115
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v125 = v118
	v126 = v119
	v127 = v120
	goto L37
L47:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
	if v135 == v136 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v156 = v135 - v136
	goto L35
L49:
	;
	v138 = int32(1)
	v143 = v132 - v138
	if v143 != 0 {
		v130 = v130 + v138
		v131 = v131 + v138
		v132 = v143
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v13)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if l1 != v17 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v17)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v156 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v165 = v156
	goto L63
L62:
	;
	v165 = base.B2i32(v92 < v60) - v93
	goto L63
L63:
	;
	return v165
}
func F_verify_dictoptions(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
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
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[120])))
	if v11 == int32(1) {
		v15 = F_SearchSysCache1(m, int32(80), l0)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			if v15 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg_internal(m, int32(53964), v8)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_errfinish(m, int32(517270), int32(361), int32(145342))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
				v21 = v19 + v20
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
				if v22 == int32(0) {
					if l1 == int32(0) {
						F_ReleaseCatCache(m, v15)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							F_errcode(m, int32(16801924))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v21 + int32(4)
								F_errmsg(m, int32(145882), v8+int32(16))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									F_errfinish(m, int32(517270), int32(373), int32(145342))
									mBase = m.M
									v46 = m.ExcPending
									if v46 != 0 {
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
					v48 = F_copyObjectImpl(m, l1)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v50 = F_OidFunctionCall1Coll(m, v22, int32(0), v48)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v15)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	} else {
		m.G0 = v8 + int32(32)
		return
	}
}
func F_view_query_is_auto_updatable(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(667043)
L2:
	;
	goto L3
L3:
	;
	v11 = int32(666840)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v12 != 0 {
		v139 = v11
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v139
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v13 != 0 {
		v139 = v11
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	if v14 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(667157)
L8:
	;
	goto L9
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v17 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(666899)
L11:
	;
	goto L12
L12:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v20 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(667102)
L14:
	;
	goto L15
L15:
	;
	v23 = int32(666977)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v24 != 0 {
		v139 = v23
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v25 != 0 {
		v139 = v23
		goto L4
	} else {
		goto L17
	}
L17:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v26 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(666626)
L19:
	;
	goto L20
L20:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+37)))
	if v29 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return int32(666483)
L22:
	;
	goto L23
L23:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+38)))
	if v32 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(666551)
L25:
	;
	goto L26
L26:
	;
	v35 = int32(666397)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v37 == int32(0) {
		v139 = v35
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 != int32(1) {
		v139 = v35
		goto L4
	} else {
		goto L28
	}
L28:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v45 != int32(63) {
		v139 = v35
		goto L4
	} else {
		goto L29
	}
L29:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50<<(uint(int32(2))%32)-int32(4))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v57 != 0 {
		v139 = v35
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+21)))
	v60 = v58 - int32(102)
	v69 = (v60<<(uint(int32(7))%32) | int32(base.Ui32(v60&int32(254))>>(uint(int32(1))%32))) & int32(255)
	if base.Ui32(int32(8)) < base.Ui32(v69) {
		v139 = v35
		goto L4
	} else {
		goto L31
	}
L31:
	;
	if int32(1)<<(uint(v69)%32)&int32(353) == int32(0) {
		v139 = v35
		goto L4
	} else {
		goto L32
	}
L32:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	if v80 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v81 = int32(667214)
	goto L35
L34:
	;
	v81 = int32(0)
	goto L35
L35:
	;
	if l1 == int32(0) {
		v139 = v81
		goto L4
	} else {
		goto L36
	}
L36:
	;
	if v80 != 0 {
		v139 = v81
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v84 = int32(666697)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v85 == int32(0) {
		v139 = v84
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v88 <= int32(0) {
		v139 = v84
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v91 = int32(0)
	if v91 < v88 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v95 = v88
	goto L42
L41:
	;
	v95 = v91
	goto L42
L42:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v97 = v91
	goto L43
L43:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v96+v97<<(uint(int32(2))%32))))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+26)))
	if v109 != 0 {
		v130 = int32(667276)
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v139 = int32(0)
	goto L4
L45:
	;
	if v130 != 0 {
		goto L55
	} else {
		goto L56
	}
L46:
	;
	v110 = int32(667439)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	if v112 != int32(6) {
		v127 = v110
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v130 = v127
	goto L45
L48:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v115 != v116 {
		v127 = v110
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v111)+28))
	if v118 != 0 {
		v127 = v110
		goto L47
	} else {
		goto L50
	}
L50:
	;
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+8)))
	if v120 < int32(0) {
		v130 = int32(667313)
		goto L45
	} else {
		goto L51
	}
L51:
	;
	if v120 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v125 = int32(0)
	goto L54
L53:
	;
	v125 = int32(667374)
	goto L54
L54:
	;
	v127 = v125
	goto L47
L55:
	;
	v132 = v97 + int32(1)
	if v95 != v132 {
		v97 = v132
		goto L43
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	goto L44
L58:
	;
	v139 = v84
	goto L4
}
func F_visibilitymap_clear(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v7 = base.I32_div_u_s(l0, int32(32672))
	if l1 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(241692), int32(0))
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(518975), int32(156), int32(241715))
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		if l1 < int32(0) {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(l1^int32(-1))<<(uint(int32(6))%32))+16))
			v28 = v19
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _consts[10]))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+l1<<(uint(int32(6))%32)+int32(-64))+16))
			v28 = v27
		}
		if v28 != v7 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(241692), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(518975), int32(156), int32(241715))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v34 = l2 << (uint(l0<<(uint(int32(1))%32)&int32(6)) % 32)
			v38 = int32(2)
			F_LockBuffer(m, l1, v38)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				if l1 < int32(0) {
					v48 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+(l1^int32(-1))<<(uint(int32(2))%32))))
					v62 = v54
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					v62 = v56 + l1<<(uint(int32(13))%32) + int32(-8192)
				}
				v65 = v62 + int32(base.Ui32(l0-v7*int32(32672))>>(uint(v38)%32)) + int32(24)
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
				v67 = v34 & v66
				if v67 != 0 {
					v70 = v66 & (v34 ^ int32(-1))
					*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v70)
					F_MarkBufferDirty(m, l1)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						F_LockBuffer(m, l1, int32(0))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v67 != int32(0))
						}
					}
				} else {
					F_LockBuffer(m, l1, int32(0))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v67 != int32(0))
					}
				}
			}
		}
	}
}
func F_visibilitymap_get_status(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	v8 = base.I32_div_u_s(l1, int32(32672))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v9 == int32(0) {
		v41 = int32(0)
		v43 = F_vm_readbuf(m, l0, v8, v41)
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v43
			if v43 == int32(0) {
				v84 = v41
			} else {
				v48 = v43
				if v48 < int32(0) {
					v57 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+(v48^int32(-1))<<(uint(int32(2))%32))))
					v71 = v63
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					v71 = v65 + v48<<(uint(int32(13))%32) + int32(-8192)
				}
				v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71+int32(base.Ui32(l1-v8*int32(32672))>>(uint(int32(2))%32)))+24)))
				v84 = v78 >> (uint(l1<<(uint(int32(1))%32)&int32(6)) % 32) & int32(3)
			}
			return v84
		}
	} else {
		if v9 < int32(0) {
			v15 = *(*int32)(unsafe.Add(mBase, _consts[9]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+(v9^int32(-1))<<(uint(int32(6))%32))+16))
			v30 = v21
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _consts[10]))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v23+v9<<(uint(int32(6))%32)+int32(-64))+16))
			v30 = v29
		}
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v30 != v8 {
			F_ReleaseBuffer(m, v31)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
				v41 = int32(0)
				v43 = F_vm_readbuf(m, l0, v8, v41)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v43
					if v43 == int32(0) {
						v84 = v41
					} else {
						v48 = v43
						if v48 < int32(0) {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+(v48^int32(-1))<<(uint(int32(2))%32))))
							v71 = v63
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, _consts[6]))
							v71 = v65 + v48<<(uint(int32(13))%32) + int32(-8192)
						}
						v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71+int32(base.Ui32(l1-v8*int32(32672))>>(uint(int32(2))%32)))+24)))
						v84 = v78 >> (uint(l1<<(uint(int32(1))%32)&int32(6)) % 32) & int32(3)
					}
					return v84
				}
			}
		} else {
			if v31 != 0 {
				v48 = v31
				if v48 < int32(0) {
					v57 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+(v48^int32(-1))<<(uint(int32(2))%32))))
					v71 = v63
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					v71 = v65 + v48<<(uint(int32(13))%32) + int32(-8192)
				}
				v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71+int32(base.Ui32(l1-v8*int32(32672))>>(uint(int32(2))%32)))+24)))
				v84 = v78 >> (uint(l1<<(uint(int32(1))%32)&int32(6)) % 32) & int32(3)
				return v84
			} else {
				v41 = int32(0)
				v43 = F_vm_readbuf(m, l0, v8, v41)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v43
					if v43 == int32(0) {
						v84 = v41
					} else {
						v48 = v43
						if v48 < int32(0) {
							v57 = *(*int32)(unsafe.Add(mBase, _consts[5]))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v57+(v48^int32(-1))<<(uint(int32(2))%32))))
							v71 = v63
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, _consts[6]))
							v71 = v65 + v48<<(uint(int32(13))%32) + int32(-8192)
						}
						v78 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71+int32(base.Ui32(l1-v8*int32(32672))>>(uint(int32(2))%32)))+24)))
						v84 = v78 >> (uint(l1<<(uint(int32(1))%32)&int32(6)) % 32) & int32(3)
					}
					return v84
				}
			}
		}
	}
}
func F_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v7 = m.G0
	v9 = v7 - int32(160)
	m.G0 = v9
	if l1 != 0 {
		v13 = l0
	} else {
		v13 = v9 + int32(158)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v13
	v16 = l1 - int32(1)
	if base.Ui32(v16) <= base.Ui32(l1) {
		v19 = v16
	} else {
		v19 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v19
	v24 = F__emscripten_memset_bulkmem(m, v9, base.I32_extend8_s(int32(0)), int32(144))
	mBase = m.M
	v25 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+76)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v24)+36)) = int32(7693)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+80)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v24)+44)) = v24 + int32(159)
	*(*int32)(unsafe.Add(mBase, uint32(v24)+84)) = v24 + int32(148)
	v37 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13))) = uint8(v37)
	v39 = F_vfprintf(m, v24, l2, l3)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return int32(0)
	} else {
		m.G0 = v24 + int32(160)
		return v39
	}
}
