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
				if base.Ui32(int32(_a_F_ValidateRestrictionEstimator_0)) <= base.Ui32(v17) {
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
									F_errmsg(m, int32(_a_F_ValidateRestrictionEstimator_1), int32(0))
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_ValidateRestrictionEstimator_2), int32(309), int32(_a_F_ValidateRestrictionEstimator_3))
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
					v47 = *(*int32)(unsafe.Add(mBase, _c_F_ValidateRestrictionEstimator[0]))
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
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_ValidateRestrictionEstimator_4)
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v70
							F_errmsg(m, int32(_a_F_ValidateRestrictionEstimator_5), v7)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_ValidateRestrictionEstimator_2), int32(292), int32(_a_F_ValidateRestrictionEstimator_3))
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
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
			v88 = m.ExcPending
			if v88 != 0 {
				return
			} else {
				F_errcode(m, int32(67141764))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_validateConnectbyTupleDesc_0), int32(0))
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						v97 = F_format_type_be(m, int32(23))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v97
							F_errdetail(m, int32(_a_F_validateConnectbyTupleDesc_1), v8+int32(48))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_validateConnectbyTupleDesc_2), int32(1434), int32(_a_F_validateConnectbyTupleDesc_3))
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
			}
		} else {
			if l1 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+388))
				if v22 != int32(25) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_validateConnectbyTupleDesc_0), int32(0))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return
							} else {
								v122 = F_format_type_be(m, int32(25))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v122
									F_errdetail(m, int32(_a_F_validateConnectbyTupleDesc_4), v8+int32(16))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_validateConnectbyTupleDesc_2), int32(1442), int32(_a_F_validateConnectbyTupleDesc_3))
										mBase = m.M
										v134 = m.ExcPending
										if v134 != 0 {
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
									F_errmsg(m, int32(_a_F_validateConnectbyTupleDesc_0), int32(0))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return
									} else {
										v42 = F_format_type_be(m, int32(23))
										mBase = m.M
										v43 = m.ExcPending
										if v43 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8))) = v42
											F_errdetail(m, int32(_a_F_validateConnectbyTupleDesc_5), v8)
											mBase = m.M
											v47 = m.ExcPending
											if v47 != 0 {
												return
											} else {
												F_errfinish(m, int32(_a_F_validateConnectbyTupleDesc_2), int32(1451), int32(_a_F_validateConnectbyTupleDesc_3))
												mBase = m.M
												v52 = m.ExcPending
												if v52 != 0 {
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
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v18)+388))
					if v55 != int32(23) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v138 = m.ExcPending
						if v138 != 0 {
							return
						} else {
							F_errcode(m, int32(67141764))
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_validateConnectbyTupleDesc_0), int32(0))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return
								} else {
									v147 = F_format_type_be(m, int32(23))
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v147
										F_errdetail(m, int32(_a_F_validateConnectbyTupleDesc_6), v8+int32(32))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return
										} else {
											F_errfinish(m, int32(_a_F_validateConnectbyTupleDesc_2), int32(1458), int32(_a_F_validateConnectbyTupleDesc_3))
											mBase = m.M
											v159 = m.ExcPending
											if v159 != 0 {
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
		v64 = m.ExcPending
		if v64 != 0 {
			return
		} else {
			F_errcode(m, int32(67141764))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_validateConnectbyTupleDesc_0), int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return
				} else {
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+68)) = v72
					*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v13
					F_errdetail(m, int32(_a_F_validateConnectbyTupleDesc_7), v8-int32(-64))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_validateConnectbyTupleDesc_2), int32(1424), int32(_a_F_validateConnectbyTupleDesc_3))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	v13 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = F_pg_detoast_datum_packed(m, l1)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = int32(4)
			v21 = int32(1)
			v22 = v17 + v21
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
			v27 = v13 + v21
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			v32 = v30 & v21
			if v32 != 0 {
				v33 = v27
			} else {
				v33 = v13 + v19
			}
			if v30 == int32(1) {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
				if v39 == int32(18) {
					v42 = int32(16)
				} else {
					v42 = int32(0)
				}
				if base.Ui32((v39-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v49 = int32(4)
				} else {
					v49 = v42
				}
				v60 = v49
			} else {
				v50 = int32(1)
				if v32 != 0 {
					v60 = int32(base.Ui32(v30)>>(uint(v50)%32)) - v50
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if v23&v21 != 0 {
				v61 = v22
			} else {
				v61 = v17 + v19
			}
			if v23 == int32(1) {
				v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
				if v67 == int32(18) {
					v70 = int32(16)
				} else {
					v70 = int32(0)
				}
				if base.Ui32((v67-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v77 = int32(4)
				} else {
					v77 = v70
				}
				v90 = v77
			} else {
				v78 = int32(1)
				if v23&v78 != 0 {
					v90 = int32(base.Ui32(v23)>>(uint(v78)%32)) - v78
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
					v90 = int32(base.Ui32(v84)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v91 = F_varstrfastcmp_locale(m, v33, v60, v61, v90, l2)
			mBase = m.M
			v92 = m.ExcPending
			if v92 != 0 {
				return int32(0)
			} else {
				if l0 != v13 {
					F_pfree(m, v13)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						if l1 != v17 {
							F_pfree(m, v17)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								return v91
							}
						} else {
							return v91
						}
					}
				} else {
					if l1 != v17 {
						F_pfree(m, v17)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int32(0)
						} else {
							return v91
						}
					} else {
						return v91
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
	var v18 int32
	_ = v18
	var v25 float64
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 float64
	_ = v43
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
	var v67 int32
	_ = v67
	var v73 float64
	_ = v73
	var v75 float64
	_ = v75
	var v78 int32
	_ = v78
	var v79 float64
	_ = v79
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v94 float64
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v189 float64
	_ = v189
	var v191 float64
	_ = v191
	var v193 float64
	_ = v193
	var v206 float64
	_ = v206
	var v216 float64
	_ = v216
	var v227 float64
	_ = v227
	var v239 float64
	_ = v239
	var v242 float64
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 float64
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v271 float64
	_ = v271
	var v273 float64
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 float64
	_ = v276
	var v279 int32
	_ = v279
	var v280 float64
	_ = v280
	var v285 float64
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v295 int32
	_ = v295
	var v301 float64
	_ = v301
	var v303 float64
	_ = v303
	var v306 int32
	_ = v306
	var v307 float64
	_ = v307
	var v318 float64
	_ = v318
	var v320 float64
	_ = v320
	var v321 float64
	_ = v321
	var v322 float64
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v406 int32
	_ = v406
	var v417 float64
	_ = v417
	var v419 float64
	_ = v419
	var v421 float64
	_ = v421
	var v434 float64
	_ = v434
	var v444 float64
	_ = v444
	var v455 float64
	_ = v455
	var v467 float64
	_ = v467
	var v470 float64
	_ = v470
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v485 float64
	_ = v485
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 float64
	_ = v503
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 float64
	_ = v522
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(80)
	m.G0 = v10
	if l0 < int32(100) {
		v536 = v3
		m.G0 = v10 + int32(80)
		return v536
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v17 = v15 + int32(40)
		v18 = int32(0)
		v25 = float64(0)
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
		if v27 != 0 {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
			if v27 != int32(1) {
				v37 = v18
				v38 = v18
				v43 = v25
				for {
					v45 = float64(1)
					v46 = v37 + v28
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+1)))
					v48 = F_scalbn(m, v45, v47)
					mBase = m.M
					v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
					v52 = F_scalbn(m, v45, v51)
					mBase = m.M
					v57 = base.F64_add(base.F64_add(v43, base.F64_div(v45, v52)), base.F64_div(v45, v48))
					v58 = int32(2)
					v59 = v37 + v58
					v61 = v38 + v58
					if v61 != v27&int32(-2) {
						v37 = v59
						v38 = v61
						v43 = v57
						continue
					} else {
						break
					}
					break
				}
				if v27&int32(1) == int32(0) {
					v90 = v57
				} else {
					v67 = v59
					v73 = v57
					v75 = float64(1)
					v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v28))))
					v79 = F_scalbn(m, v75, v78)
					mBase = m.M
					v90 = base.F64_add(v73, base.F64_div(v75, v79))
				}
			} else {
				v67 = v18
				v73 = v25
				v75 = float64(1)
				v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67+v28))))
				v79 = F_scalbn(m, v75, v78)
				mBase = m.M
				v90 = base.F64_add(v73, base.F64_div(v75, v79))
			}
			v92 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
			v93 = base.F64_div(v92, v90)
			v94 = base.F64_convert_i32_u(v27)
			if base.F64_le(v93, base.F64_mul(v94, float64(2.5))) == int32(0) {
				v206 = v93
				if base.F64_gt(v206, float64(1.4316557653333333e+08)) == int32(0) {
					v227 = v206
				} else {
					v216 = F_log(m, base.F64_add(base.F64_mul(v206, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v227 = base.F64_mul(v216, float64(-4.294967296e+09))
				}
				v239 = v227
			} else {
				v101 = v27 & int32(3)
				v102 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
				v103 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v27) {
					v112 = int32(0)
					v113 = v103
					v114 = v103
					for {
						v121 = v113 + v102
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
						v123 = int32(0)
						v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+1)))
						v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+2)))
						v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+3)))
						v137 = v114 + base.B2i32(v122 == v123) + base.B2i32(v126 == v123) + base.B2i32(v130 == v123) + base.B2i32(v134 == v123)
						v138 = int32(4)
						v139 = v113 + v138
						v141 = v112 + v138
						if v141 != v27&int32(-4) {
							v112 = v141
							v113 = v139
							v114 = v137
							continue
						} else {
							break
						}
						break
					}
					if v101 == int32(0) {
						v178 = v137
					} else {
						v147 = v139
						v148 = v137
						v157 = v147
						v158 = v148
						v161 = v103
						for {
							v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v102))))
							v169 = v158 + base.B2i32(v166 == int32(0))
							v170 = int32(1)
							v173 = v161 + v170
							if v173 != v101 {
								v157 = v157 + v170
								v158 = v169
								v161 = v173
								continue
							} else {
								break
							}
							break
						}
						v178 = v169
					}
				} else {
					v147 = v103
					v148 = v103
					v157 = v147
					v158 = v148
					v161 = v103
					for {
						v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+v102))))
						v169 = v158 + base.B2i32(v166 == int32(0))
						v170 = int32(1)
						v173 = v161 + v170
						if v173 != v101 {
							v157 = v157 + v170
							v158 = v169
							v161 = v173
							continue
						} else {
							break
						}
						break
					}
					v178 = v169
				}
				if v178 == int32(0) {
					v227 = v93
					v239 = v227
				} else {
					v189 = F_log(m, base.F64_div(v94, base.F64_convert_i32_s(v178)))
					mBase = m.M
					v239 = base.F64_mul(v189, v94)
				}
			}
		} else {
			v191 = *(*float64)(unsafe.Add(mBase, uint32(v17)+8))
			v193 = base.F64_div(v191, float64(0))
			if base.F64_le(v193, base.F64_mul(base.F64_convert_i32_u(v27), float64(2.5))) != 0 {
				v227 = v193
			} else {
				v206 = v193
				if base.F64_gt(v206, float64(1.4316557653333333e+08)) == int32(0) {
					v227 = v206
				} else {
					v216 = F_log(m, base.F64_add(base.F64_mul(v206, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v227 = base.F64_mul(v216, float64(-4.294967296e+09))
				}
			}
			v239 = v227
		}
		if base.F64_le(v239, float64(1)) != 0 {
			v242 = float64(1)
		} else {
			v242 = v239
		}
		v245 = v15 - int32(-64)
		v246 = int32(0)
		v253 = float64(0)
		v255 = *(*int32)(unsafe.Add(mBase, uint32(v245)+4))
		if v255 != 0 {
			v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
			if v255 != int32(1) {
				v265 = v246
				v266 = v246
				v271 = v253
				for {
					v273 = float64(1)
					v274 = v265 + v256
					v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
					v276 = F_scalbn(m, v273, v275)
					mBase = m.M
					v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
					v280 = F_scalbn(m, v273, v279)
					mBase = m.M
					v285 = base.F64_add(base.F64_add(v271, base.F64_div(v273, v280)), base.F64_div(v273, v276))
					v286 = int32(2)
					v287 = v265 + v286
					v289 = v266 + v286
					if v289 != v255&int32(-2) {
						v265 = v287
						v266 = v289
						v271 = v285
						continue
					} else {
						break
					}
					break
				}
				if v255&int32(1) == int32(0) {
					v318 = v285
				} else {
					v295 = v287
					v301 = v285
					v303 = float64(1)
					v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v256))))
					v307 = F_scalbn(m, v303, v306)
					mBase = m.M
					v318 = base.F64_add(v301, base.F64_div(v303, v307))
				}
			} else {
				v295 = v246
				v301 = v253
				v303 = float64(1)
				v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295+v256))))
				v307 = F_scalbn(m, v303, v306)
				mBase = m.M
				v318 = base.F64_add(v301, base.F64_div(v303, v307))
			}
			v320 = *(*float64)(unsafe.Add(mBase, uint32(v245)+8))
			v321 = base.F64_div(v320, v318)
			v322 = base.F64_convert_i32_u(v255)
			if base.F64_le(v321, base.F64_mul(v322, float64(2.5))) == int32(0) {
				v434 = v321
				if base.F64_gt(v434, float64(1.4316557653333333e+08)) == int32(0) {
					v455 = v434
				} else {
					v444 = F_log(m, base.F64_add(base.F64_mul(v434, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v455 = base.F64_mul(v444, float64(-4.294967296e+09))
				}
				v467 = v455
			} else {
				v329 = v255 & int32(3)
				v330 = *(*int32)(unsafe.Add(mBase, uint32(v245)+16))
				v331 = int32(0)
				if base.Ui32(int32(4)) <= base.Ui32(v255) {
					v340 = int32(0)
					v341 = v331
					v342 = v331
					for {
						v349 = v341 + v330
						v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
						v351 = int32(0)
						v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
						v358 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+2)))
						v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+3)))
						v365 = v342 + base.B2i32(v350 == v351) + base.B2i32(v354 == v351) + base.B2i32(v358 == v351) + base.B2i32(v362 == v351)
						v366 = int32(4)
						v367 = v341 + v366
						v369 = v340 + v366
						if v369 != v255&int32(-4) {
							v340 = v369
							v341 = v367
							v342 = v365
							continue
						} else {
							break
						}
						break
					}
					if v329 == int32(0) {
						v406 = v365
					} else {
						v375 = v367
						v376 = v365
						v385 = v375
						v386 = v376
						v389 = v331
						for {
							v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+v330))))
							v397 = v386 + base.B2i32(v394 == int32(0))
							v398 = int32(1)
							v401 = v389 + v398
							if v401 != v329 {
								v385 = v385 + v398
								v386 = v397
								v389 = v401
								continue
							} else {
								break
							}
							break
						}
						v406 = v397
					}
				} else {
					v375 = v331
					v376 = v331
					v385 = v375
					v386 = v376
					v389 = v331
					for {
						v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v385+v330))))
						v397 = v386 + base.B2i32(v394 == int32(0))
						v398 = int32(1)
						v401 = v389 + v398
						if v401 != v329 {
							v385 = v385 + v398
							v386 = v397
							v389 = v401
							continue
						} else {
							break
						}
						break
					}
					v406 = v397
				}
				if v406 == int32(0) {
					v455 = v321
					v467 = v455
				} else {
					v417 = F_log(m, base.F64_div(v322, base.F64_convert_i32_s(v406)))
					mBase = m.M
					v467 = base.F64_mul(v417, v322)
				}
			}
		} else {
			v419 = *(*float64)(unsafe.Add(mBase, uint32(v245)+8))
			v421 = base.F64_div(v419, float64(0))
			if base.F64_le(v421, base.F64_mul(base.F64_convert_i32_u(v255), float64(2.5))) != 0 {
				v455 = v421
			} else {
				v434 = v421
				if base.F64_gt(v434, float64(1.4316557653333333e+08)) == int32(0) {
					v455 = v434
				} else {
					v444 = F_log(m, base.F64_add(base.F64_mul(v434, float64(-2.3283064365386963e-10)), float64(1)))
					mBase = m.M
					v455 = base.F64_mul(v444, float64(-4.294967296e+09))
				}
			}
			v467 = v455
		}
		if base.F64_le(v467, float64(1)) != 0 {
			v470 = float64(1)
		} else {
			v470 = v467
		}
		v472 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_varstr_abbrev_abort[0])))
		if v472 != int32(1) {
			v503 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
			if base.F64_lt(base.F64_mul(v470, v503), v242) != 0 {
				if base.Ui32(l0) < base.Ui32(int32(_a_F_varstr_abbrev_abort_0)) {
					v536 = v3
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = base.F64_mul(v503, float64(0.65))
					v536 = v3
				}
				m.G0 = v10 + int32(80)
				return v536
			} else {
				v511 = int32(1)
				v513 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_varstr_abbrev_abort[0])))
				if v513 != v511 {
					v536 = v511
					m.G0 = v10 + int32(80)
					return v536
				} else {
					v518 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v519 = m.ExcPending
					if v519 != 0 {
						return int32(0)
					} else {
						if v518 == int32(0) {
							v536 = v511
							m.G0 = v10 + int32(80)
							return v536
						} else {
							v522 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
							*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v522
							*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v470
							*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v242
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
							F_errmsg_internal(m, int32(_a_F_varstr_abbrev_abort_1), v10)
							mBase = m.M
							v529 = m.ExcPending
							if v529 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_varstr_abbrev_abort_2), int32(2650), int32(_a_F_varstr_abbrev_abort_3))
								mBase = m.M
								v534 = m.ExcPending
								if v534 != 0 {
									return int32(0)
								} else {
									v536 = v511
									m.G0 = v10 + int32(80)
									return v536
								}
							}
						}
					}
				}
			}
		} else {
			v477 = F_errstart(m, int32(15), int32(0))
			mBase = m.M
			v480 = m.ExcPending
			if v480 != 0 {
				return int32(0)
			} else {
				if v477 == int32(0) {
					v503 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
					if base.F64_lt(base.F64_mul(v470, v503), v242) != 0 {
						if base.Ui32(l0) < base.Ui32(int32(_a_F_varstr_abbrev_abort_0)) {
							v536 = v3
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = base.F64_mul(v503, float64(0.65))
							v536 = v3
						}
						m.G0 = v10 + int32(80)
						return v536
					} else {
						v511 = int32(1)
						v513 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_varstr_abbrev_abort[0])))
						if v513 != v511 {
							v536 = v511
							m.G0 = v10 + int32(80)
							return v536
						} else {
							v518 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v519 = m.ExcPending
							if v519 != 0 {
								return int32(0)
							} else {
								if v518 == int32(0) {
									v536 = v511
									m.G0 = v10 + int32(80)
									return v536
								} else {
									v522 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
									*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v522
									*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v470
									*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v242
									*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
									F_errmsg_internal(m, int32(_a_F_varstr_abbrev_abort_1), v10)
									mBase = m.M
									v529 = m.ExcPending
									if v529 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_varstr_abbrev_abort_2), int32(2650), int32(_a_F_varstr_abbrev_abort_3))
										mBase = m.M
										v534 = m.ExcPending
										if v534 != 0 {
											return int32(0)
										} else {
											v536 = v511
											m.G0 = v10 + int32(80)
											return v536
										}
									}
								}
							}
						}
					}
				} else {
					v485 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
					*(*float64)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = v485
					*(*float64)(unsafe.Add(mBase, uint32(v10)+48)) = v470
					*(*float64)(unsafe.Add(mBase, uint32(v10)+40)) = v242
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = l0
					*(*float64)(unsafe.Add(mBase, uint32(v10)+56)) = base.F64_div(v242, base.F64_convert_i32_u(l0))
					F_errmsg_internal(m, int32(_a_F_varstr_abbrev_abort_4), v10+int32(32))
					mBase = m.M
					v497 = m.ExcPending
					if v497 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_varstr_abbrev_abort_2), int32(2584), int32(_a_F_varstr_abbrev_abort_3))
						mBase = m.M
						v502 = m.ExcPending
						if v502 != 0 {
							return int32(0)
						} else {
							v503 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
							if base.F64_lt(base.F64_mul(v470, v503), v242) != 0 {
								if base.Ui32(l0) < base.Ui32(int32(_a_F_varstr_abbrev_abort_0)) {
									v536 = v3
								} else {
									*(*float64)(unsafe.Add(mBase, uint32(v15)+88)) = base.F64_mul(v503, float64(0.65))
									v536 = v3
								}
								m.G0 = v10 + int32(80)
								return v536
							} else {
								v511 = int32(1)
								v513 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_varstr_abbrev_abort[0])))
								if v513 != v511 {
									v536 = v511
									m.G0 = v10 + int32(80)
									return v536
								} else {
									v518 = F_errstart(m, int32(15), int32(0))
									mBase = m.M
									v519 = m.ExcPending
									if v519 != 0 {
										return int32(0)
									} else {
										if v518 == int32(0) {
											v536 = v511
											m.G0 = v10 + int32(80)
											return v536
										} else {
											v522 = *(*float64)(unsafe.Add(mBase, uint32(v15)+88))
											*(*float64)(unsafe.Add(mBase, uint32(v10)+24)) = v522
											*(*float64)(unsafe.Add(mBase, uint32(v10)+16)) = v470
											*(*float64)(unsafe.Add(mBase, uint32(v10)+8)) = v242
											*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
											F_errmsg_internal(m, int32(_a_F_varstr_abbrev_abort_1), v10)
											mBase = m.M
											v529 = m.ExcPending
											if v529 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_varstr_abbrev_abort_2), int32(2650), int32(_a_F_varstr_abbrev_abort_3))
												mBase = m.M
												v534 = m.ExcPending
												if v534 != 0 {
													return int32(0)
												} else {
													v536 = v511
													m.G0 = v10 + int32(80)
													return v536
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
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
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v10 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = F_pg_detoast_datum_packed(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(4)
	v18 = int32(1)
	v19 = v10 + v18
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v24 = v22 & v18
	if v24 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = v19
	goto L6
L5:
	;
	v25 = v10 + v16
	goto L6
L6:
	;
	v26 = int32(1)
	v27 = v14 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v28&v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = v27
	goto L9
L8:
	;
	v31 = v14 + v16
	goto L9
L9:
	;
	if v22 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v28 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v37 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v48 = int32(1)
	if v24 != 0 {
		v58 = int32(base.Ui32(v22)>>(uint(v48)%32)) - v48
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v40 = int32(16)
	goto L16
L15:
	;
	v40 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v47 = int32(4)
	goto L19
L18:
	;
	v47 = v40
	goto L19
L19:
	;
	v58 = v47
	goto L10
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	v88 = base.B2i32(v58 < v87)
	if v58 < v87 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v64 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v75 = int32(1)
	if v28&v75 != 0 {
		v87 = int32(base.Ui32(v28)>>(uint(v75)%32)) - v75
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v67 = int32(16)
	goto L27
L26:
	;
	v67 = int32(0)
	goto L27
L27:
	;
	if base.Ui32((v64-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v74 = int32(4)
	goto L30
L29:
	;
	v74 = v67
	goto L30
L30:
	;
	v87 = v74
	goto L21
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v87 = int32(base.Ui32(v81)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v89 = v58
	goto L34
L33:
	;
	v89 = v87
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if l0 != v10 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v25|v31)&int32(3) != 0 {
		v120 = v25
		v121 = v31
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v25
	v114 = v31
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v25
	v98 = v31
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
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
	F_pfree(m, v10)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if l1 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v151 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v160 = v151
	goto L63
L62:
	;
	v160 = base.B2i32(v87 < v58) - v88
	goto L63
L63:
	;
	return v160
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
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_verify_dictoptions[0])))
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
					F_errmsg_internal(m, int32(_a_F_verify_dictoptions_0), v8)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_verify_dictoptions_1), int32(361), int32(_a_F_verify_dictoptions_2))
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
								F_errmsg(m, int32(_a_F_verify_dictoptions_3), v8+int32(16))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_verify_dictoptions_1), int32(373), int32(_a_F_verify_dictoptions_2))
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v8 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(_a_F_view_query_is_auto_updatable_0)
L2:
	;
	goto L3
L3:
	;
	v11 = int32(_a_F_view_query_is_auto_updatable_1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	if v12 != 0 {
		v140 = v11
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v140
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v13 != 0 {
		v140 = v11
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
	return int32(_a_F_view_query_is_auto_updatable_2)
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
	return int32(_a_F_view_query_is_auto_updatable_3)
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
	return int32(_a_F_view_query_is_auto_updatable_4)
L14:
	;
	goto L15
L15:
	;
	v23 = int32(_a_F_view_query_is_auto_updatable_5)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	if v24 != 0 {
		v140 = v23
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v25 != 0 {
		v140 = v23
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
	return int32(_a_F_view_query_is_auto_updatable_6)
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
	return int32(_a_F_view_query_is_auto_updatable_7)
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
	return int32(_a_F_view_query_is_auto_updatable_8)
L25:
	;
	goto L26
L26:
	;
	v35 = int32(_a_F_view_query_is_auto_updatable_9)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v37 == int32(0) {
		v140 = v35
		goto L4
	} else {
		goto L27
	}
L27:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v40 != int32(1) {
		v140 = v35
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
		v140 = v35
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
		v140 = v35
		goto L4
	} else {
		goto L30
	}
L30:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+21)))
	v60 = v58 - int32(102)
	v65 = int32(1)
	v69 = (v60<<(uint(int32(7))%32) | int32(base.Ui32(v60&int32(254))>>(uint(v65)%32))) & int32(255)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v69))|base.B2i32(v65<<(uint(v69)%32)&int32(353) == int32(0)) != 0 {
		v140 = v35
		goto L4
	} else {
		goto L31
	}
L31:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v56)+32))
	if v81 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v82 = int32(_a_F_view_query_is_auto_updatable_10)
	goto L34
L33:
	;
	v82 = int32(0)
	goto L34
L34:
	;
	if base.B2i32(l1 == int32(0))|v81 != 0 {
		v140 = v82
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v86 = int32(_a_F_view_query_is_auto_updatable_11)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v87 == int32(0) {
		v140 = v86
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v90 <= int32(0) {
		v140 = v86
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v93 = int32(0)
	if v93 < v90 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v97 = v90
	goto L40
L39:
	;
	v97 = v93
	goto L40
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v99 = v93
	goto L41
L41:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v98+v99<<(uint(int32(2))%32))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+26)))
	if v111 != 0 {
		v132 = int32(_a_F_view_query_is_auto_updatable_12)
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v140 = int32(0)
	goto L4
L43:
	;
	if v132 != 0 {
		goto L53
	} else {
		goto L54
	}
L44:
	;
	v112 = int32(_a_F_view_query_is_auto_updatable_13)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	if v114 != int32(6) {
		v129 = v112
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v132 = v129
	goto L43
L46:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v117 != v118 {
		v129 = v112
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v113)+28))
	if v120 != 0 {
		v129 = v112
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v122 = int32(*(*int16)(unsafe.Add(mBase, uint32(v113)+8)))
	if v122 < int32(0) {
		v132 = int32(_a_F_view_query_is_auto_updatable_14)
		goto L43
	} else {
		goto L49
	}
L49:
	;
	if v122 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v127 = int32(0)
	goto L52
L51:
	;
	v127 = int32(_a_F_view_query_is_auto_updatable_15)
	goto L52
L52:
	;
	v129 = v127
	goto L45
L53:
	;
	v134 = v99 + int32(1)
	if v97 != v134 {
		v99 = v134
		goto L41
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L42
L56:
	;
	v140 = v86
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v7 = base.I32_div_u_s(l0, int32(_a_F_visibilitymap_clear_0))
	if l1 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v81 = m.ExcPending
		if v81 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_visibilitymap_clear_1), int32(0))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_visibilitymap_clear_2), int32(156), int32(_a_F_visibilitymap_clear_3))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
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
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_clear[0]))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(l1^int32(-1))<<(uint(int32(6))%32))+16))
			v28 = v19
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_clear[1]))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+l1<<(uint(int32(6))%32)+int32(-64))+16))
			v28 = v27
		}
		if v28 != v7 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_visibilitymap_clear_1), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_visibilitymap_clear_2), int32(156), int32(_a_F_visibilitymap_clear_3))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
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
					v48 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_clear[2]))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v48+(l1^int32(-1))<<(uint(int32(2))%32))))
					v62 = v54
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_clear[3]))
					v62 = v56 + l1<<(uint(int32(13))%32) + int32(-8192)
				}
				v63 = v62 + int32(base.Ui32(l0-v7*int32(_a_F_visibilitymap_clear_0))>>(uint(v38)%32))
				v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+24)))
				v65 = v34 & v64
				if v65 != 0 {
					v68 = v64 & (v34 ^ int32(-1))
					*(*uint8)(unsafe.Add(mBase, uint32(v63)+24)) = uint8(v68)
					F_MarkBufferDirty(m, l1)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_LockBuffer(m, l1, int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v65 != int32(0))
						}
					}
				} else {
					F_LockBuffer(m, l1, int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v65 != int32(0))
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	v8 = base.I32_div_u_s(l1, int32(_a_F_visibilitymap_get_status_0))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v9 == int32(0) {
		v40 = int32(0)
		v42 = F_vm_readbuf(m, l0, v8, v40)
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v42
			if v42 == int32(0) {
				v83 = v40
			} else {
				v47 = v42
				if v47 < int32(0) {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[0]))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+(v47^int32(-1))<<(uint(int32(2))%32))))
					v70 = v62
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[1]))
					v70 = v64 + v47<<(uint(int32(13))%32) + int32(-8192)
				}
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(base.Ui32(l1-v8*int32(_a_F_visibilitymap_get_status_0))>>(uint(int32(2))%32)))+24)))
				v83 = int32(base.Ui32(v77)>>(uint(l1<<(uint(int32(1))%32)&int32(6))%32)) & int32(3)
			}
			return v83
		}
	} else {
		if v9 < int32(0) {
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[2]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v15+(v9^int32(-1))<<(uint(int32(6))%32))+16))
			v30 = v21
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[3]))
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
				v40 = int32(0)
				v42 = F_vm_readbuf(m, l0, v8, v40)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v42
					if v42 == int32(0) {
						v83 = v40
					} else {
						v47 = v42
						if v47 < int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[0]))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+(v47^int32(-1))<<(uint(int32(2))%32))))
							v70 = v62
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[1]))
							v70 = v64 + v47<<(uint(int32(13))%32) + int32(-8192)
						}
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(base.Ui32(l1-v8*int32(_a_F_visibilitymap_get_status_0))>>(uint(int32(2))%32)))+24)))
						v83 = int32(base.Ui32(v77)>>(uint(l1<<(uint(int32(1))%32)&int32(6))%32)) & int32(3)
					}
					return v83
				}
			}
		} else {
			if v31 != 0 {
				v47 = v31
				if v47 < int32(0) {
					v56 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[0]))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+(v47^int32(-1))<<(uint(int32(2))%32))))
					v70 = v62
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[1]))
					v70 = v64 + v47<<(uint(int32(13))%32) + int32(-8192)
				}
				v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(base.Ui32(l1-v8*int32(_a_F_visibilitymap_get_status_0))>>(uint(int32(2))%32)))+24)))
				v83 = int32(base.Ui32(v77)>>(uint(l1<<(uint(int32(1))%32)&int32(6))%32)) & int32(3)
				return v83
			} else {
				v40 = int32(0)
				v42 = F_vm_readbuf(m, l0, v8, v40)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v42
					if v42 == int32(0) {
						v83 = v40
					} else {
						v47 = v42
						if v47 < int32(0) {
							v56 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[0]))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v56+(v47^int32(-1))<<(uint(int32(2))%32))))
							v70 = v62
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_get_status[1]))
							v70 = v64 + v47<<(uint(int32(13))%32) + int32(-8192)
						}
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+int32(base.Ui32(l1-v8*int32(_a_F_visibilitymap_get_status_0))>>(uint(int32(2))%32)))+24)))
						v83 = int32(base.Ui32(v77)>>(uint(l1<<(uint(int32(1))%32)&int32(6))%32)) & int32(3)
					}
					return v83
				}
			}
		}
	}
}
func F_vsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v6 = m.G0
	v8 = v6 - int32(160)
	m.G0 = v8
	if l1 != 0 {
		v12 = l0
	} else {
		v12 = v8 + int32(158)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = v12
	v14 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+152)) = l1 - base.B2i32(l1 != v14)
	base.MemoryFill(m, v8, v14, int32(144))
	v21 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_vsnprintf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v8 + int32(159)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v8 + int32(148)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v14)
	v35 = F_vfprintf(m, v8, l2, l3)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		m.G0 = v8 + int32(160)
		return v35
	}
}
