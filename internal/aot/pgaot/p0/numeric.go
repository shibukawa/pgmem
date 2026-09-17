package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_convert_numeric_to_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v29 float64
	_ = v29
	var v36 int64
	_ = v36
	var v42 float64
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	if l1 <= int32(2201) {
		switch l1 - int32(16) {
		case 0:
			if l0 != 0 {
				v29 = float64(1)
			} else {
				v29 = float64(0)
			}
			return v29
		case 1, 2, 3, 6, 9:
			v46 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v46)
			return float64(0)
		case 4:
			v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			return base.F64_convert_i64_s(v36)
		case 5:
			return base.F64_convert_i32_s(base.I32_extend16_s(l0))
		case 7:
			return base.F64_convert_i32_s(l0)
		case 8, 10:
			return base.F64_convert_i32_u(l0)
		default:
			switch l1 - int32(700) {
			case 0:
				return base.F64_promote_f32(base.F32_reinterpret_i32(l0))
			case 1:
				v42 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
				return v42
			default:
				if l1 == int32(1700) {
					v52 = F_DirectFunctionCall1Coll(m, int32(1492), int32(0), l0)
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return float64(0)
					} else {
						v56 = *(*float64)(unsafe.Add(mBase, uint32(v52)))
						return v56
					}
				} else {
					v46 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v46)
					return float64(0)
				}
			}
		}
	} else {
		if l1 <= int32(3768) {
			if base.B2i32(l1 == int32(3734))|base.B2i32(base.Ui32(l1-int32(2202)) < base.Ui32(int32(5))) != 0 {
				return base.F64_convert_i32_u(l0)
			} else {
				v46 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v46)
				return float64(0)
			}
		} else {
			switch l1 - int32(4089) {
			case 0, 7:
				return base.F64_convert_i32_u(l0)
			case 1, 2, 3, 4, 5, 6:
				v46 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v46)
				return float64(0)
			default:
				if l1 == int32(3769) {
					return base.F64_convert_i32_u(l0)
				} else {
					if l1 != int32(_a_F_convert_numeric_to_scalar_0) {
						v46 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v46)
						return float64(0)
					} else {
						return base.F64_convert_i32_u(l0)
					}
				}
			}
		}
	}
}
func F_executeNumericItemMethod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
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
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l3 != 0 {
		switch v13 - int32(2) {
		case 0:
			v90 = int32(0)
			v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v93 = F_DirectFunctionCall1Coll(m, l4, v90, v92)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v96 = v11 + int32(44)
				v97 = F_jspGetNext(m, l1, v96)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					if v97|l5 == int32(0) {
						v141 = v90
						m.G0 = v11 + int32(80)
						return v141
					} else {
						v103 = F_palloc(m, int32(20))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(2)
							v107 = F_pg_detoast_datum(m, v93)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v107
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if int32(0) < v110 {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
									v114 = F_executeItemOptUnwrapTarget(m, l0, v96, v103, l5, v113)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v141 = v114
										m.G0 = v11 + int32(80)
										return v141
									}
								} else {
									if l5 == int32(0) {
										v141 = v90
										m.G0 = v11 + int32(80)
										return v141
									} else {
										v118 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
										if v118 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v103
											*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v118
											*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v118
											*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v103
											v127 = F_list_make2_impl(m, v11+int32(40), v11+int32(36))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v127
												v141 = v90
												m.G0 = v11 + int32(80)
												return v141
											}
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
											if v132 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l5))) = v103
												v141 = v90
												m.G0 = v11 + int32(80)
												return v141
											} else {
												v136 = F_lappend(m, v132, v103)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v136
													v141 = v90
													m.G0 = v11 + int32(80)
													return v141
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
		default:
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
			if v68 != int32(1) {
				v141 = int32(2)
				m.G0 = v11 + int32(80)
				return v141
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101449858))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v79 = F_jspOperationName(m, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v79
							F_errmsg(m, int32(_a_F_executeNumericItemMethod_0), v11)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(2311), int32(_a_F_executeNumericItemMethod_2))
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
		case 14:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v44
				F_errmsg_internal(m, int32(_a_F_executeNumericItemMethod_3), v11+int32(32))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(1680), int32(_a_F_executeNumericItemMethod_4))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 16:
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17&int32(536870912) != 0 {
				v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
				if v68 != int32(1) {
					v141 = int32(2)
					m.G0 = v11 + int32(80)
					return v141
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(101449858))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v79 = F_jspOperationName(m, v78)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v11))) = v79
								F_errmsg(m, int32(_a_F_executeNumericItemMethod_0), v11)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(2311), int32(_a_F_executeNumericItemMethod_2))
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
			} else {
				if v17&int32(1073741824) != 0 {
					v56 = int32(1)
					v59 = int32(0)
					v61 = F_executeAnyItem(m, l0, l1, v16, l5, v56, v56, v56, v59, v59)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v141 = v61
						m.G0 = v11 + int32(80)
						return v141
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v28
						F_errmsg_internal(m, int32(_a_F_executeNumericItemMethod_5), v11+int32(16))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(3629), int32(_a_F_executeNumericItemMethod_6))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
		if v13 == int32(2) {
			v90 = int32(0)
			v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v93 = F_DirectFunctionCall1Coll(m, l4, v90, v92)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v96 = v11 + int32(44)
				v97 = F_jspGetNext(m, l1, v96)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					if v97|l5 == int32(0) {
						v141 = v90
						m.G0 = v11 + int32(80)
						return v141
					} else {
						v103 = F_palloc(m, int32(20))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v103))) = int32(2)
							v107 = F_pg_detoast_datum(m, v93)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v103)+4)) = v107
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
								if int32(0) < v110 {
									v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
									v114 = F_executeItemOptUnwrapTarget(m, l0, v96, v103, l5, v113)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v141 = v114
										m.G0 = v11 + int32(80)
										return v141
									}
								} else {
									if l5 == int32(0) {
										v141 = v90
										m.G0 = v11 + int32(80)
										return v141
									} else {
										v118 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
										if v118 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v103
											*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v118
											*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v118
											*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v103
											v127 = F_list_make2_impl(m, v11+int32(40), v11+int32(36))
											mBase = m.M
											v128 = m.ExcPending
											if v128 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v127
												v141 = v90
												m.G0 = v11 + int32(80)
												return v141
											}
										} else {
											v132 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
											if v132 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l5))) = v103
												v141 = v90
												m.G0 = v11 + int32(80)
												return v141
											} else {
												v136 = F_lappend(m, v132, v103)
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v136
													v141 = v90
													m.G0 = v11 + int32(80)
													return v141
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
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
			if v68 != int32(1) {
				v141 = int32(2)
				m.G0 = v11 + int32(80)
				return v141
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101449858))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return int32(0)
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v79 = F_jspOperationName(m, v78)
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v79
							F_errmsg(m, int32(_a_F_executeNumericItemMethod_0), v11)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(2311), int32(_a_F_executeNumericItemMethod_2))
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
		}
	}
}
func F_numeric_abbrev_convert(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
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
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v15 + int64(1)
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
		if v19&int32(1) == int32(0) {
			v41 = v11
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v25 = int32(1)
			v26 = int32(base.Ui32(v19) >> (uint(v25) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = v26<<(uint(int32(2))%32) + int32(12)
			v33 = v26 - v25
			if v33 == int32(0) {
				v41 = v24
			} else {
				base.MemoryCopy(m, v24+int32(4), v11+int32(1), v33)
				v41 = v24
			}
		}
		v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
		v44 = base.I32_extend16_s(v43)
		if base.Ui32(int32(_a_F_numeric_abbrev_convert_0)) <= base.Ui32(v43) {
			if v44 == int32(-4096) {
				v52 = int32(2147483647)
			} else {
				v52 = int32(-2147483648)
			}
			if v44 == int32(-12288) {
				v55 = int32(-2147483647)
			} else {
				v55 = v52
			}
			v239 = v55
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
			v62 = base.B2i32(int32(0) <= v44)
			if int32(0) <= v44 {
				v63 = int32(-8)
			} else {
				v63 = int32(-6)
			}
			v64 = int32(base.Ui32(v56)>>(uint(int32(2))%32)) + v63
			if int32(0) <= v44 {
				v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+6)))
				v75 = v65
			} else {
				v75 = v43<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v43&int32(63)
			}
			v81 = v43 & int32(_a_F_numeric_abbrev_convert_0)
			if v81 == int32(_a_F_numeric_abbrev_convert_1) {
				v84 = v43 << (uint(int32(1)) % 32) & int32(_a_F_numeric_abbrev_convert_2)
			} else {
				v84 = v81
			}
			if base.Ui32(v64) < base.Ui32(int32(2)) {
				v168 = int32(0)
			} else {
				if v75 < int32(-11) {
					v168 = int32(0)
				} else {
					if int32(20) < v75 {
						v168 = int32(2147483647)
					} else {
						if v44 < int32(0) {
							v98 = int32(6)
						} else {
							v98 = int32(8)
						}
						v99 = v41 + v98
						if base.Ui32(int32(4)) <= base.Ui32(v64) {
							v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+2)))
							v104 = v103
						} else {
							v104 = int32(0)
						}
						v106 = v75 << (uint(int32(2)) % 32)
						v107 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99))))
						if int32(1000) <= v107 {
							v111 = base.I32_div_s(v104, int32(10))
							v159 = base.I32_extend16_s(v111) + v107*int32(1000)
							v160 = v106 + int32(47)
						} else {
							if int32(100) <= v107 {
								v159 = v107*int32(_a_F_numeric_abbrev_convert_3) + v104
								v160 = v106 + int32(46)
							} else {
								if int32(10) <= v107 {
									if base.Ui32(int32(6)) <= base.Ui32(v64) {
										v129 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+4)))
										v131 = base.I32_div_s(v129, int32(1000))
										v134 = base.I32_extend16_s(v131)
									} else {
										v134 = int32(0)
									}
									v159 = v134 + (v107*int32(_a_F_numeric_abbrev_convert_4) + v104*int32(10))
									v160 = v106 + int32(45)
								} else {
									if base.Ui32(int32(6)) <= base.Ui32(v64) {
										v145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v99)+4)))
										v147 = base.I32_div_s(v145, int32(100))
										v150 = base.I32_extend16_s(v147)
									} else {
										v150 = int32(0)
									}
									v159 = v150 + (v107*int32(_a_F_numeric_abbrev_convert_5) + v104*int32(100))
									v160 = v106 + int32(44)
								}
							}
						}
						v168 = v160<<(uint(int32(24))%32) | v159
					}
				}
			}
			if v84 != 0 {
				v171 = v168
			} else {
				v171 = int32(0) - v168
			}
			v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)))
			if v172 != int32(1) {
				v239 = v171
			} else {
				v175 = int32(24)
				v176 = v10 + v175
				v181 = int32(711645284)
				v184 = v171 - int32(1636608428) ^ v181 - int32(1455628627)
				v189 = v184 ^ int32(-1636608428) - base.I32_rotl(v184, int32(25))
				v194 = v189 ^ v181 - base.I32_rotl(v189, int32(16))
				v198 = v194 ^ v184 - base.I32_rotl(v194, int32(4))
				v202 = v198 ^ v189 - base.I32_rotl(v198, int32(14))
				v206 = v202 ^ v194 - base.I32_rotl(v202, v175)
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
				v212 = int32(32) - v211
				v214 = v209 + int32(base.Ui32(v206)>>(uint(v212)%32))
				v215 = v206 << (uint(v211) % 32)
				if v215 != 0 {
					v222 = int32(32) - (base.I32_clz(v215) ^ int32(31))
					v223 = int32(255)
					if base.Ui32(v212&v223) < base.Ui32(v222&v223) {
						v228 = v212 + int32(1)
					} else {
						v228 = v222
					}
					v232 = v228
				} else {
					v232 = v212 + int32(1)
				}
				v234 = v232 & int32(255)
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214))))
				if base.Ui32(v235) < base.Ui32(v234) {
					v237 = v234
				} else {
					v237 = v235
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v214))) = uint8(v237)
				v239 = v171
			}
		}
		if l0 != v11 {
			F_pfree(m, v11)
			mBase = m.M
			v247 = m.ExcPending
			if v247 != 0 {
				return int32(0)
			} else {
				return v239
			}
		} else {
			return v239
		}
	}
}
func F_numeric_add_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v15 = base.I32_extend16_s(v14)
	if base.Ui32(v14) <= base.Ui32(int32(_a_F_numeric_add_opt_error_0)) {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v19 = base.I32_extend16_s(v18)
		if base.Ui32(int32(_a_F_numeric_add_opt_error_0)) < base.Ui32(v18) {
			v47 = v19
			if v47&int32(_a_F_numeric_add_opt_error_1) != int32(_a_F_numeric_add_opt_error_2) {
				if v14 != int32(_a_F_numeric_add_opt_error_3) {
					if v14 != int32(_a_F_numeric_add_opt_error_4) {
						if v47&int32(_a_F_numeric_add_opt_error_1) == int32(_a_F_numeric_add_opt_error_4) {
							v94 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_5), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v202 = v94
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v98 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_6), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v202 = v98
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					} else {
						if v47&int32(_a_F_numeric_add_opt_error_1) == int32(_a_F_numeric_add_opt_error_3) {
							v70 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_7), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v202 = v70
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v74 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_5), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v202 = v74
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					}
				} else {
					if v47&int32(_a_F_numeric_add_opt_error_1) == int32(_a_F_numeric_add_opt_error_4) {
						v82 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_7), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v202 = v82
							m.G0 = v12 + int32(80)
							return v202
						}
					} else {
						v86 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_6), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v202 = v86
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_7), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v202 = v56
					m.G0 = v12 + int32(80)
					return v202
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = base.B2i32(int32(0) <= v15)
			if int32(0) <= v15 {
				v29 = int32(-8)
			} else {
				v29 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(base.Ui32(int32(base.Ui32(v22)>>(uint(int32(2))%32))+v29) >> (uint(int32(1)) % 32))
			if int32(0) <= v15 {
				v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v101 = v100
			} else {
				v101 = v14<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v14&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v101
			v103 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v103
			v112 = base.B2i32(v15 < v103)
			if v15 < v103 {
				v113 = int32(base.Ui32(v14)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v113 = v14 & int32(_a_F_numeric_add_opt_error_8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v113
			v120 = v14 & int32(_a_F_numeric_add_opt_error_2)
			if v120 == int32(_a_F_numeric_add_opt_error_9) {
				v123 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_add_opt_error_10)
			} else {
				v123 = v120
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v123
			if v15 < v103 {
				v127 = int32(6)
			} else {
				v127 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l0 + v127
			v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v136 = base.B2i32(int32(0) <= v19)
			if int32(0) <= v19 {
				v137 = int32(-8)
			} else {
				v137 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(base.Ui32(int32(base.Ui32(v130)>>(uint(int32(2))%32))+v137) >> (uint(int32(1)) % 32))
			if int32(0) <= v19 {
				v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				v152 = v142
			} else {
				v152 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v152
			v154 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v154
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v154
			*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v154
			v160 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v160
			v169 = base.B2i32(v19 < v160)
			if v19 < v160 {
				v170 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v170 = v18 & int32(_a_F_numeric_add_opt_error_8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v170
			v177 = v18 & int32(_a_F_numeric_add_opt_error_2)
			if v177 == int32(_a_F_numeric_add_opt_error_9) {
				v180 = v18 << (uint(int32(1)) % 32) & int32(_a_F_numeric_add_opt_error_10)
			} else {
				v180 = v177
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v180
			if v19 < v160 {
				v184 = int32(6)
			} else {
				v184 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l1 + v184
			v192 = v12 + int32(8)
			F_add_var(m, v12+int32(56), v12+int32(32), v192)
			mBase = m.M
			v194 = m.ExcPending
			if v194 != 0 {
				return int32(0)
			} else {
				v195 = F_make_result_opt_error(m, v192, l2)
				mBase = m.M
				v196 = m.ExcPending
				if v196 != 0 {
					return int32(0)
				} else {
					v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					if v197 == int32(0) {
						v202 = v195
						m.G0 = v12 + int32(80)
						return v202
					} else {
						F_pfree(m, v197)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return int32(0)
						} else {
							v202 = v195
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			}
		}
	} else {
		if v15 == int32(-16384) {
			v56 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_7), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v202 = v56
				m.G0 = v12 + int32(80)
				return v202
			}
		} else {
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v47 = v45
			if v47&int32(_a_F_numeric_add_opt_error_1) != int32(_a_F_numeric_add_opt_error_2) {
				if v14 != int32(_a_F_numeric_add_opt_error_3) {
					if v14 != int32(_a_F_numeric_add_opt_error_4) {
						if v47&int32(_a_F_numeric_add_opt_error_1) == int32(_a_F_numeric_add_opt_error_4) {
							v94 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_5), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v202 = v94
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v98 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_6), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v202 = v98
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					} else {
						if v47&int32(_a_F_numeric_add_opt_error_1) == int32(_a_F_numeric_add_opt_error_3) {
							v70 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_7), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v202 = v70
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v74 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_5), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v202 = v74
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					}
				} else {
					if v47&int32(_a_F_numeric_add_opt_error_1) == int32(_a_F_numeric_add_opt_error_4) {
						v82 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_7), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v202 = v82
							m.G0 = v12 + int32(80)
							return v202
						}
					} else {
						v86 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_6), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v202 = v86
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(_a_F_numeric_add_opt_error_7), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v202 = v56
					m.G0 = v12 + int32(80)
					return v202
				}
			}
		}
	}
}
func F_numeric_avg_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(429) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v42 = F_pg_detoast_datum_packed(m, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
			v48 = int32(1)
			v49 = v42 + v48
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v52 = v50 & v48
			if v50 == v48 {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
				if v58 == int32(18) {
					v61 = int32(16)
				} else {
					v61 = int32(0)
				}
				if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v68 = int32(4)
				} else {
					v68 = v61
				}
				v79 = v68
			} else {
				v69 = int32(1)
				if v52 != 0 {
					v79 = int32(base.Ui32(v50)>>(uint(v69)%32)) - v69
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v79
			if v52 != 0 {
				v85 = v49
			} else {
				v85 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v85
			v88 = F_palloc0(m, int32(112))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
				v93 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_deserialize[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v93
				v96 = v8 + int32(32)
				v97 = F_pq_getmsgint64(m, v96)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v97
					v101 = v8 + int32(8)
					F_numericvar_deserialize(m, v96, v101)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						F_accum_sum_add(m, v88+int32(16), v101)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							v109 = F_pq_getmsgint(m, v96, int32(4))
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v88)+72)) = v109
								v112 = F_pq_getmsgint64(m, v96)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v88)+80)) = v112
									v115 = F_pq_getmsgint64(m, v96)
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v88)+88)) = v115
										v118 = F_pq_getmsgint64(m, v96)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v88)+96)) = v118
											v121 = F_pq_getmsgint64(m, v96)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v88)+104)) = v121
												F_pq_getmsgend(m, v96)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
													if v126 != 0 {
														F_pfree(m, v126)
														mBase = m.M
														v128 = m.ExcPending
														if v128 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(48)
															return v88
														}
													} else {
														m.G0 = v8 + int32(48)
														return v88
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v136 = m.ExcPending
		if v136 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_avg_deserialize_0), int32(0))
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_avg_deserialize_1), int32(_a_F_numeric_avg_deserialize_2), int32(_a_F_numeric_avg_deserialize_3))
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
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
func F_numeric_div_trunc(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v22 = base.I32_extend16_s(v21)
			if base.Ui32(v21) <= base.Ui32(int32(_a_F_numeric_div_trunc_0)) {
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
				v26 = base.I32_extend16_s(v25)
				if base.Ui32(int32(_a_F_numeric_div_trunc_0)) < base.Ui32(v25) {
					v53 = v26
					if v53&int32(_a_F_numeric_div_trunc_1) != int32(_a_F_numeric_div_trunc_2) {
						if v21 != int32(_a_F_numeric_div_trunc_3) {
							if v21 != int32(_a_F_numeric_div_trunc_4) {
								v149 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_5), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									v257 = v149
									m.G0 = v11 + int32(80)
									return v257
								}
							} else {
								if base.Ui32(int32(_a_F_numeric_div_trunc_2)) <= base.Ui32(v53&int32(_a_F_numeric_div_trunc_1)) {
									v75 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_6), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v257 = v75
										m.G0 = v11 + int32(80)
										return v257
									}
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
									if int32(0) <= base.I32_extend16_s(v53) {
										v85 = int32(-8)
									} else {
										v85 = int32(-6)
									}
									if base.Ui32(int32(base.Ui32(v77)>>(uint(int32(2))%32))+v85) < base.Ui32(int32(2)) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v269 = m.ExcPending
										if v269 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(33816706))
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_numeric_div_trunc_7), int32(0))
												mBase = m.M
												v276 = m.ExcPending
												if v276 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_numeric_div_trunc_8), int32(3403), int32(_a_F_numeric_div_trunc_9))
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
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
										v94 = v53 & int32(_a_F_numeric_div_trunc_2)
										if v94 == int32(_a_F_numeric_div_trunc_10) {
											v97 = v53 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
										} else {
											v97 = v94
										}
										if v97 != int32(_a_F_numeric_div_trunc_11) {
											v102 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_12), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v257 = v102
												m.G0 = v11 + int32(80)
												return v257
											}
										} else {
											v106 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_13), int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												v257 = v106
												m.G0 = v11 + int32(80)
												return v257
											}
										}
									}
								}
							}
						} else {
							if base.Ui32(int32(_a_F_numeric_div_trunc_2)) <= base.Ui32(v53&int32(_a_F_numeric_div_trunc_1)) {
								v114 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_6), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v257 = v114
									m.G0 = v11 + int32(80)
									return v257
								}
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
								if int32(0) <= base.I32_extend16_s(v53) {
									v124 = int32(-8)
								} else {
									v124 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v116)>>(uint(int32(2))%32))+v124) < base.Ui32(int32(2)) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v285 = m.ExcPending
									if v285 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(33816706))
										mBase = m.M
										v288 = m.ExcPending
										if v288 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_numeric_div_trunc_7), int32(0))
											mBase = m.M
											v292 = m.ExcPending
											if v292 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_numeric_div_trunc_8), int32(3421), int32(_a_F_numeric_div_trunc_9))
												mBase = m.M
												v297 = m.ExcPending
												if v297 != 0 {
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
									v133 = v53 & int32(_a_F_numeric_div_trunc_2)
									if v133 == int32(_a_F_numeric_div_trunc_10) {
										v136 = v53 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
									} else {
										v136 = v133
									}
									if v136 != int32(_a_F_numeric_div_trunc_11) {
										v141 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_13), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											v257 = v141
											m.G0 = v11 + int32(80)
											return v257
										}
									} else {
										v145 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_12), int32(0))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return int32(0)
										} else {
											v257 = v145
											m.G0 = v11 + int32(80)
											return v257
										}
									}
								}
							}
						}
					} else {
						v63 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_6), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v257 = v63
							m.G0 = v11 + int32(80)
							return v257
						}
					}
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v35 = base.B2i32(int32(0) <= v22)
					if int32(0) <= v22 {
						v36 = int32(-8)
					} else {
						v36 = int32(-6)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = int32(base.Ui32(int32(base.Ui32(v29)>>(uint(int32(2))%32))+v36) >> (uint(int32(1)) % 32))
					if int32(0) <= v22 {
						v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+6)))
						v152 = v151
					} else {
						v152 = v21<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v21&int32(63)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v152
					v154 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v154
					v163 = base.B2i32(v22 < v154)
					if v22 < v154 {
						v164 = int32(base.Ui32(v21)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v164 = v21 & int32(_a_F_numeric_div_trunc_14)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v164
					v171 = v21 & int32(_a_F_numeric_div_trunc_2)
					if v171 == int32(_a_F_numeric_div_trunc_10) {
						v174 = v21 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
					} else {
						v174 = v171
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v174
					if v22 < v154 {
						v178 = int32(6)
					} else {
						v178 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v14 + v178
					v181 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
					v187 = base.B2i32(int32(0) <= v26)
					if int32(0) <= v26 {
						v188 = int32(-8)
					} else {
						v188 = int32(-6)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = int32(base.Ui32(int32(base.Ui32(v181)>>(uint(int32(2))%32))+v188) >> (uint(int32(1)) % 32))
					if int32(0) <= v26 {
						v193 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+6)))
						v203 = v193
					} else {
						v203 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v203
					v205 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v205
					v207 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v207
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v207
					*(*int64)(unsafe.Add(mBase, uint32(v11)+24)) = v207
					v216 = base.B2i32(v26 < v205)
					if v26 < v205 {
						v217 = int32(6)
					} else {
						v217 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v19 + v217
					if v26 < v205 {
						v226 = int32(base.Ui32(v25)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v226 = v25 & int32(_a_F_numeric_div_trunc_14)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v226
					v233 = v25 & int32(_a_F_numeric_div_trunc_2)
					if v233 == int32(_a_F_numeric_div_trunc_10) {
						v236 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
					} else {
						v236 = v233
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v236
					v243 = v11 + int32(8)
					v244 = int32(0)
					F_div_var(m, v11+int32(56), v11+int32(32), v243, v244, v244, int32(1))
					mBase = m.M
					v248 = m.ExcPending
					if v248 != 0 {
						return int32(0)
					} else {
						v250 = F_make_result_opt_error(m, v243, int32(0))
						mBase = m.M
						v251 = m.ExcPending
						if v251 != 0 {
							return int32(0)
						} else {
							v252 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
							if v252 == int32(0) {
								v257 = v250
								m.G0 = v11 + int32(80)
								return v257
							} else {
								F_pfree(m, v252)
								mBase = m.M
								v256 = m.ExcPending
								if v256 != 0 {
									return int32(0)
								} else {
									v257 = v250
									m.G0 = v11 + int32(80)
									return v257
								}
							}
						}
					}
				}
			} else {
				if v22 == int32(-16384) {
					v63 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_6), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v257 = v63
						m.G0 = v11 + int32(80)
						return v257
					}
				} else {
					v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
					v53 = v52
					if v53&int32(_a_F_numeric_div_trunc_1) != int32(_a_F_numeric_div_trunc_2) {
						if v21 != int32(_a_F_numeric_div_trunc_3) {
							if v21 != int32(_a_F_numeric_div_trunc_4) {
								v149 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_5), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									v257 = v149
									m.G0 = v11 + int32(80)
									return v257
								}
							} else {
								if base.Ui32(int32(_a_F_numeric_div_trunc_2)) <= base.Ui32(v53&int32(_a_F_numeric_div_trunc_1)) {
									v75 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_6), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v257 = v75
										m.G0 = v11 + int32(80)
										return v257
									}
								} else {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
									if int32(0) <= base.I32_extend16_s(v53) {
										v85 = int32(-8)
									} else {
										v85 = int32(-6)
									}
									if base.Ui32(int32(base.Ui32(v77)>>(uint(int32(2))%32))+v85) < base.Ui32(int32(2)) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v269 = m.ExcPending
										if v269 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(33816706))
											mBase = m.M
											v272 = m.ExcPending
											if v272 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(_a_F_numeric_div_trunc_7), int32(0))
												mBase = m.M
												v276 = m.ExcPending
												if v276 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_numeric_div_trunc_8), int32(3403), int32(_a_F_numeric_div_trunc_9))
													mBase = m.M
													v281 = m.ExcPending
													if v281 != 0 {
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
										v94 = v53 & int32(_a_F_numeric_div_trunc_2)
										if v94 == int32(_a_F_numeric_div_trunc_10) {
											v97 = v53 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
										} else {
											v97 = v94
										}
										if v97 != int32(_a_F_numeric_div_trunc_11) {
											v102 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_12), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v257 = v102
												m.G0 = v11 + int32(80)
												return v257
											}
										} else {
											v106 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_13), int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												v257 = v106
												m.G0 = v11 + int32(80)
												return v257
											}
										}
									}
								}
							}
						} else {
							if base.Ui32(int32(_a_F_numeric_div_trunc_2)) <= base.Ui32(v53&int32(_a_F_numeric_div_trunc_1)) {
								v114 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_6), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v257 = v114
									m.G0 = v11 + int32(80)
									return v257
								}
							} else {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
								if int32(0) <= base.I32_extend16_s(v53) {
									v124 = int32(-8)
								} else {
									v124 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v116)>>(uint(int32(2))%32))+v124) < base.Ui32(int32(2)) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v285 = m.ExcPending
									if v285 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(33816706))
										mBase = m.M
										v288 = m.ExcPending
										if v288 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(_a_F_numeric_div_trunc_7), int32(0))
											mBase = m.M
											v292 = m.ExcPending
											if v292 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_numeric_div_trunc_8), int32(3421), int32(_a_F_numeric_div_trunc_9))
												mBase = m.M
												v297 = m.ExcPending
												if v297 != 0 {
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
									v133 = v53 & int32(_a_F_numeric_div_trunc_2)
									if v133 == int32(_a_F_numeric_div_trunc_10) {
										v136 = v53 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
									} else {
										v136 = v133
									}
									if v136 != int32(_a_F_numeric_div_trunc_11) {
										v141 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_13), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											v257 = v141
											m.G0 = v11 + int32(80)
											return v257
										}
									} else {
										v145 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_12), int32(0))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return int32(0)
										} else {
											v257 = v145
											m.G0 = v11 + int32(80)
											return v257
										}
									}
								}
							}
						}
					} else {
						v63 = F_make_result_opt_error(m, int32(_a_F_numeric_div_trunc_6), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v257 = v63
							m.G0 = v11 + int32(80)
							return v257
						}
					}
				}
			}
		}
	}
}
func F_numeric_fac(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v83 int64
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v104 int64
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v124 int32
	_ = v124
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
	if int64(0) <= v16 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L45
	}
L2:
	;
	if base.Ui64(v16) <= base.Ui64(int64(1)) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L41
	}
L5:
	;
	m.G0 = v13 + int32(48)
	return v159
L6:
	;
	v23 = F_make_result_opt_error(m, int32(_a_F_numeric_fac_0), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui64(int64(32178)) <= base.Ui64(v16) {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	v159 = v23
	goto L5
L11:
	;
	v29 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v29
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v29
	v36 = F_palloc(m, int32(12))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v36
	v39 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v36))) = uint16(v39)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
	v46 = v39
	v49 = v36 + int32(12)
	v52 = v16
	goto L13
L13:
	;
	v57 = v49 - int32(2)
	v59 = base.I64_div_u_s(v52, int64(10000))
	v62 = v59*int64(55536) + v52
	*(*uint16)(unsafe.Add(mBase, uint32(v57))) = uint16(v62)
	v65 = v46 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v52) {
		v46 = v65
		v49 = v57
		v52 = v59
		goto L13
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v57
	if v16 == int64(2) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L14
L16:
	;
	v152 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v152
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v154 == v152 {
		v159 = v142
		goto L5
	} else {
		goto L39
	}
L17:
	;
	v74 = F_make_result_opt_error(m, v13, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L9
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v80 = int32(0)
	v83 = v16
	goto L21
L20:
	;
	v142 = v74
	goto L16
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_fac[0]))
	if v87 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v138 = F_make_result_opt_error(m, v13, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L37
	}
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	F_pfree(m, v80)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v93 = F_palloc(m, int32(12))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L9
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v93
	v96 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v93))) = uint16(v96)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = int64(0)
	v104 = v83 - int64(1)
	v105 = v96
	v108 = v93 + int32(12)
	v111 = v104
	goto L32
L32:
	;
	v116 = v108 - int32(2)
	v118 = base.I64_div_u_s(v111, int64(10000))
	v121 = v118*int64(55536) + v111
	*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v121)
	v124 = v105 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v111) {
		v105 = v124
		v108 = v116
		v111 = v118
		goto L32
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v124
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v116
	F_mul_var(m, v13, v13+int32(24), v13, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	if base.Ui64(int64(3)) < base.Ui64(v83) {
		v80 = v93
		v83 = v104
		goto L21
	} else {
		goto L36
	}
L36:
	;
	goto L22
L37:
	;
	F_pfree(m, v93)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	v142 = v138
	goto L16
L39:
	;
	F_pfree(m, v154)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	v159 = v142
	goto L5
L41:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L9
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_numeric_fac_1), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_numeric_fac_2), int32(3753), int32(_a_F_numeric_fac_3))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L9
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
	F_errcode(m, int32(50331778))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L9
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_numeric_fac_4), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_numeric_fac_2), int32(3763), int32(_a_F_numeric_fac_3))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_numeric_float8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v36 int32
	_ = v36
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
		if base.Ui32(int32(_a_F_numeric_float8_0)) <= base.Ui32(v8) {
			if v8 != int32(_a_F_numeric_float8_1) {
				if v8 != int32(_a_F_numeric_float8_2) {
					v24 = F_Float8GetDatum(m, math.Float64frombits(uint64(0x7ff8000000000000)))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						return v24
					}
				} else {
					v16 = F_Float8GetDatum(m, math.Float64frombits(uint64(0x7ff0000000000000)))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						return v16
					}
				}
			} else {
				v20 = F_Float8GetDatum(m, math.Float64frombits(uint64(0xfff0000000000000)))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			}
		} else {
			v28 = int32(0)
			v31 = F_DirectFunctionCall1Coll(m, int32(618), v28, v4)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int32(0)
			} else {
				v33 = F_DirectFunctionCall1Coll(m, int32(1453), v28, v31)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v31)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						return v33
					}
				}
			}
		}
	}
}
func F_numeric_int4_opt_error(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int64
	_ = v106
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	if l1 == v3 {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		if base.Ui32(v14) <= base.Ui32(int32(_a_F_numeric_int4_opt_error_0)) {
			v47 = v14
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v53 = base.I32_extend16_s(v47)
			v55 = base.B2i32(int32(0) <= v53)
			if int32(0) <= v53 {
				v56 = int32(-8)
			} else {
				v56 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(base.Ui32(int32(base.Ui32(v48)>>(uint(int32(2))%32))+v56) >> (uint(int32(1)) % 32))
			if int32(0) <= v53 {
				v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v71 = v61
			} else {
				v71 = v47<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v47&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v71
			v73 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v73
			v82 = base.B2i32(v53 < v73)
			if v53 < v73 {
				v83 = int32(base.Ui32(v47)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v83 = v47 & int32(_a_F_numeric_int4_opt_error_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v83
			v90 = v47 & int32(_a_F_numeric_int4_opt_error_2)
			if v90 == int32(_a_F_numeric_int4_opt_error_3) {
				v93 = v47 << (uint(int32(1)) % 32) & int32(_a_F_numeric_int4_opt_error_4)
			} else {
				v93 = v90
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v93
			if v53 < v73 {
				v97 = int32(6)
			} else {
				v97 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l0 + v97
			v104 = F_numericvar_to_int64(m, v8+int32(-32), v8+int32(-8))
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int32(0)
			} else {
				if v104 != 0 {
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v10)+56))
					if base.Ui64(int64(-4294967297)) < base.Ui64(v106-int64(2147483648)) {
						v133 = base.I32_wrap_i64(v106)
						m.G0 = v10 - int32(-64)
						return v133
					} else {
						if l1 != 0 {
							v112 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v112)
							v133 = int32(0)
							m.G0 = v10 - int32(-64)
							return v133
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50331778))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_numeric_int4_opt_error_5), int32(0))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_numeric_int4_opt_error_6), int32(_a_F_numeric_int4_opt_error_7), int32(_a_F_numeric_int4_opt_error_8))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
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
					if l1 != 0 {
						v112 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v112)
						v133 = int32(0)
						m.G0 = v10 - int32(-64)
						return v133
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_numeric_int4_opt_error_5), int32(0))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_numeric_int4_opt_error_6), int32(_a_F_numeric_int4_opt_error_7), int32(_a_F_numeric_int4_opt_error_8))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
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
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v14 == int32(_a_F_numeric_int4_opt_error_2) {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_numeric_int4_opt_error_9)
						F_errmsg(m, int32(_a_F_numeric_int4_opt_error_10), v10)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_int4_opt_error_6), int32(_a_F_numeric_int4_opt_error_11), int32(_a_F_numeric_int4_opt_error_8))
							mBase = m.M
							v149 = m.ExcPending
							if v149 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_numeric_int4_opt_error_9)
						F_errmsg(m, int32(_a_F_numeric_int4_opt_error_12), v8+int32(-48))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_int4_opt_error_6), int32(_a_F_numeric_int4_opt_error_13), int32(_a_F_numeric_int4_opt_error_8))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
		v40 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v40)
		v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		if base.Ui32(v42) < base.Ui32(int32(_a_F_numeric_int4_opt_error_2)) {
			v47 = v42
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v53 = base.I32_extend16_s(v47)
			v55 = base.B2i32(int32(0) <= v53)
			if int32(0) <= v53 {
				v56 = int32(-8)
			} else {
				v56 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(base.Ui32(int32(base.Ui32(v48)>>(uint(int32(2))%32))+v56) >> (uint(int32(1)) % 32))
			if int32(0) <= v53 {
				v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v71 = v61
			} else {
				v71 = v47<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v47&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v71
			v73 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v73
			v82 = base.B2i32(v53 < v73)
			if v53 < v73 {
				v83 = int32(base.Ui32(v47)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v83 = v47 & int32(_a_F_numeric_int4_opt_error_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v83
			v90 = v47 & int32(_a_F_numeric_int4_opt_error_2)
			if v90 == int32(_a_F_numeric_int4_opt_error_3) {
				v93 = v47 << (uint(int32(1)) % 32) & int32(_a_F_numeric_int4_opt_error_4)
			} else {
				v93 = v90
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v93
			if v53 < v73 {
				v97 = int32(6)
			} else {
				v97 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l0 + v97
			v104 = F_numericvar_to_int64(m, v8+int32(-32), v8+int32(-8))
			mBase = m.M
			v105 = m.ExcPending
			if v105 != 0 {
				return int32(0)
			} else {
				if v104 != 0 {
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v10)+56))
					if base.Ui64(int64(-4294967297)) < base.Ui64(v106-int64(2147483648)) {
						v133 = base.I32_wrap_i64(v106)
						m.G0 = v10 - int32(-64)
						return v133
					} else {
						if l1 != 0 {
							v112 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v112)
							v133 = int32(0)
							m.G0 = v10 - int32(-64)
							return v133
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50331778))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_numeric_int4_opt_error_5), int32(0))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_numeric_int4_opt_error_6), int32(_a_F_numeric_int4_opt_error_7), int32(_a_F_numeric_int4_opt_error_8))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
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
					if l1 != 0 {
						v112 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v112)
						v133 = int32(0)
						m.G0 = v10 - int32(-64)
						return v133
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_numeric_int4_opt_error_5), int32(0))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_numeric_int4_opt_error_6), int32(_a_F_numeric_int4_opt_error_7), int32(_a_F_numeric_int4_opt_error_8))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
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
			v45 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v45)
			v133 = v3
			m.G0 = v10 - int32(-64)
			return v133
		}
	}
}
func F_numeric_ln(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		v16 = base.I32_extend16_s(v15)
		if base.Ui32(int32(_a_F_numeric_ln_0)) <= base.Ui32(v15) {
			if v16 == int32(-4096) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352583810))
					mBase = m.M
					v125 = m.ExcPending
					if v125 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_numeric_ln_1), int32(0))
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_ln_2), int32(3951), int32(_a_F_numeric_ln_3))
							mBase = m.M
							v134 = m.ExcPending
							if v134 != 0 {
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v24 = F_palloc(m, int32(base.Ui32(v21)>>(uint(int32(2))%32)))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v28 = int32(base.Ui32(v26) >> (uint(int32(2)) % 32))
					if v28 == int32(0) {
						v111 = v24
					} else {
						base.MemoryCopy(m, v24, v11, v28)
						v111 = v24
					}
					m.G0 = v8 + int32(48)
					return v111
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v38 = base.B2i32(int32(0) <= v16)
			if int32(0) <= v16 {
				v39 = int32(-8)
			} else {
				v39 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(base.Ui32(int32(base.Ui32(v32)>>(uint(int32(2))%32))+v39) >> (uint(int32(1)) % 32))
			if int32(0) <= v16 {
				v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+6)))
				v54 = v44
			} else {
				v54 = v15<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v15&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v54
			v56 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v56
			v58 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v8))) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v58
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v58
			v67 = base.B2i32(v16 < v56)
			if v16 < v56 {
				v68 = int32(6)
			} else {
				v68 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v11 + v68
			v76 = v15 & int32(_a_F_numeric_ln_0)
			if v76 == int32(_a_F_numeric_ln_4) {
				v79 = v15 << (uint(int32(1)) % 32) & int32(_a_F_numeric_ln_5)
			} else {
				v79 = v76
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v79
			if v16 < v56 {
				v87 = int32(base.Ui32(v15)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v87 = v15 & int32(_a_F_numeric_ln_6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v87
			v90 = v8 + int32(24)
			v93 = F_estimate_ln_dweight(m, v90)
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = int32(16) - v93
				if v87 < v95 {
					v97 = v95
				} else {
					v97 = v87
				}
				if int32(1000) <= v97 {
					v100 = int32(1000)
				} else {
					v100 = v97
				}
				F_ln_var(m, v90, v8, v100)
				mBase = m.M
				v102 = m.ExcPending
				if v102 != 0 {
					return int32(0)
				} else {
					v104 = F_make_result_opt_error(m, v8, int32(0))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						if v106 == int32(0) {
							v111 = v104
							m.G0 = v8 + int32(48)
							return v111
						} else {
							F_pfree(m, v106)
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int32(0)
							} else {
								v111 = v104
								m.G0 = v8 + int32(48)
								return v111
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_lt(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(_a_F_numeric_lt_0)
			v27 = v25 & v26
			if v27 == v26 {
				if v25 != int32(_a_F_numeric_lt_1) {
					if v25 != int32(_a_F_numeric_lt_0) {
						if v24 != int32(_a_F_numeric_lt_2) {
							v46 = int32(-1)
						} else {
							v46 = int32(0)
						}
						v163 = v46
					} else {
						v163 = base.B2i32(v24 != int32(_a_F_numeric_lt_0))
					}
				} else {
					if v24 == int32(_a_F_numeric_lt_0) {
						v41 = int32(-1)
					} else {
						v41 = base.B2i32(v24 != int32(_a_F_numeric_lt_1))
					}
					v163 = v41
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_lt_0)) <= base.Ui32(v24) {
					if v24 == int32(_a_F_numeric_lt_2) {
						v53 = int32(1)
					} else {
						v53 = int32(-1)
					}
					v163 = v53
				} else {
					v55 = v6 + int32(6)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v63 = base.B2i32(int32(0) <= base.I32_extend16_s(v25))
					if int32(0) <= base.I32_extend16_s(v25) {
						v64 = int32(-8)
					} else {
						v64 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v25) {
						v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55))))
						v76 = v66
					} else {
						v76 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v78 = int32(base.Ui32(int32(base.Ui32(v56)>>(uint(int32(2))%32))+v64) >> (uint(int32(1)) % 32))
					v80 = v11 + int32(6)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v88 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v89 = int32(-8)
					} else {
						v89 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v24) {
						v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v80))))
						v101 = v91
					} else {
						v101 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v102 = int32(1)
					v103 = int32(base.Ui32(int32(base.Ui32(v81)>>(uint(int32(2))%32))+v89) >> (uint(v102) % 32))
					v109 = v24 & int32(_a_F_numeric_lt_0)
					if v109 == int32(_a_F_numeric_lt_3) {
						v112 = v24 << (uint(v102) % 32) & int32(_a_F_numeric_lt_4)
					} else {
						v112 = v109
					}
					if v78 == int32(0) {
						if v103 == int32(0) {
							v163 = int32(0)
						} else {
							if v112 == int32(_a_F_numeric_lt_4) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v163 = v122
						}
					} else {
						if v27 == int32(_a_F_numeric_lt_3) {
							v129 = v25 << (uint(int32(1)) % 32) & int32(_a_F_numeric_lt_4)
						} else {
							v129 = v27
						}
						if v103 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v163 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v25) {
								v137 = v6 + int32(8)
							} else {
								v137 = v55
							}
							if int32(0) <= base.I32_extend16_s(v24) {
								v140 = v11 + int32(8)
							} else {
								v140 = v80
							}
							if v129 == int32(0) {
								if v112 == int32(_a_F_numeric_lt_4) {
									v163 = int32(1)
								} else {
									v146 = F_cmp_abs_common(m, v137, v78, v76, v140, v103, v101)
									mBase = m.M
									v163 = v146
								}
							} else {
								if v112 == int32(0) {
									v163 = int32(-1)
								} else {
									v150 = F_cmp_abs_common(m, v140, v103, v101, v137, v78, v76)
									mBase = m.M
									v163 = v150
								}
							}
						}
					}
				}
			}
			v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v164 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v168 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v163) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v163) >> (uint(int32(31)) % 32))
					}
				}
			} else {
				v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v168 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						return int32(base.Ui32(v163) >> (uint(int32(31)) % 32))
					}
				} else {
					return int32(base.Ui32(v163) >> (uint(int32(31)) % 32))
				}
			}
		}
	}
}
func F_numeric_poly_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v141 int64
	_ = v141
	var v145 int32
	_ = v145
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = v11 + int32(8)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 == v2 {
		v33 = int32(0)
		if v14 == v33 {
			v41 = v33
		} else {
			v36 = v33
			v37 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
			v41 = v37
		}
		v44 = v41
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		switch v19 - int32(429) {
		case 0:
			if v14 == int32(0) {
				v44 = int32(1)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+168))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
				v36 = v26
				v37 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
				v44 = v41
			}
		case 1:
			if v14 == int32(0) {
				v44 = int32(2)
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+368))
				v36 = v31
				v37 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
				v44 = v41
			}
		default:
			v33 = int32(0)
			if v14 == v33 {
				v41 = v33
			} else {
				v36 = v33
				v37 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = v36
				v41 = v37
			}
			v44 = v41
		}
	}
	if v44 != 0 {
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v45 == int32(0) {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v49 = v48
		} else {
			v49 = v2
		}
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v50 != 0 {
			v145 = v49
			m.G0 = v11 + int32(16)
			return v145
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v51 == int32(0) {
				v145 = v49
				m.G0 = v11 + int32(16)
				return v145
			} else {
				if v49 == int32(0) {
					v56 = int32(_a_F_numeric_poly_combine_0)
					v57 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_poly_combine[0]))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int32)(unsafe.Add(mBase, _c_F_numeric_poly_combine[0])) = v59
					v62 = v11 + int32(12)
					v63 = int32(0)
					v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v64 == v63 {
						v81 = int32(0)
						if v62 == v81 {
							v89 = v81
						} else {
							v84 = v81
							v85 = v63
							*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
							v89 = v85
						}
						v92 = v89
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
						switch v67 - int32(429) {
						case 0:
							if v62 == int32(0) {
								v92 = int32(1)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v64)+168))
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+20))
								v84 = v74
								v85 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
								v92 = v89
							}
						case 1:
							if v62 == int32(0) {
								v92 = int32(2)
							} else {
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v64)+368))
								v84 = v79
								v85 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
								v92 = v89
							}
						default:
							v81 = int32(0)
							if v62 == v81 {
								v89 = v81
							} else {
								v84 = v81
								v85 = v63
								*(*int32)(unsafe.Add(mBase, uint32(v62))) = v84
								v89 = v85
							}
							v92 = v89
						}
					}
					if v92 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_numeric_poly_combine_1), int32(0))
							mBase = m.M
							v175 = m.ExcPending
							if v175 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_numeric_poly_combine_2), int32(_a_F_numeric_poly_combine_3), int32(_a_F_numeric_poly_combine_4))
								mBase = m.M
								v180 = m.ExcPending
								if v180 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, _c_F_numeric_poly_combine[0])) = v96
						v99 = F_palloc0(m, int32(48))
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							v103 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v99))) = uint8(v103)
							v105 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v99)+8)) = v105
							v107 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
							v108 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v99)+24)) = v108
							*(*int64)(unsafe.Add(mBase, uint32(v99)+16)) = v107
							v111 = *(*int64)(unsafe.Add(mBase, uint32(v51)+32))
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v51)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v99)+40)) = v112
							*(*int64)(unsafe.Add(mBase, uint32(v99)+32)) = v111
							*(*int32)(unsafe.Add(mBase, _c_F_numeric_poly_combine[0])) = v57
							v145 = v99
							m.G0 = v11 + int32(16)
							return v145
						}
					}
				} else {
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
					if v117 <= int64(0) {
						v145 = v49
					} else {
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v120 + v117
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
						v125 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
						v126 = v124 + v125
						*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v126
						v130 = *(*int64)(unsafe.Add(mBase, uint32(v49)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v126) < base.Ui64(v124))) + (v123 + v130)
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v51)+40))
						v135 = *(*int64)(unsafe.Add(mBase, uint32(v49)+32))
						v136 = *(*int64)(unsafe.Add(mBase, uint32(v51)+32))
						v137 = v135 + v136
						*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v137
						v141 = *(*int64)(unsafe.Add(mBase, uint32(v49)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v137) < base.Ui64(v135))) + (v134 + v141)
						v145 = v49
					}
					m.G0 = v11 + int32(16)
					return v145
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v158 = m.ExcPending
		if v158 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_poly_combine_1), int32(0))
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_poly_combine_2), int32(_a_F_numeric_poly_combine_5), int32(_a_F_numeric_poly_combine_6))
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
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
func F_numeric_poly_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(429) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v42 = F_pg_detoast_datum_packed(m, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
			v48 = int32(1)
			v49 = v42 + v48
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v52 = v50 & v48
			if v50 == v48 {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
				if v58 == int32(18) {
					v61 = int32(16)
				} else {
					v61 = int32(0)
				}
				if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v68 = int32(4)
				} else {
					v68 = v61
				}
				v79 = v68
			} else {
				v69 = int32(1)
				if v52 != 0 {
					v79 = int32(base.Ui32(v50)>>(uint(v69)%32)) - v69
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v79
			if v52 != 0 {
				v85 = v49
			} else {
				v85 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v85
			v88 = F_palloc0(m, int32(48))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
				v93 = v8 + int32(32)
				v94 = F_pq_getmsgint64(m, v93)
				mBase = m.M
				v95 = m.ExcPending
				if v95 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v94
					v98 = v8 + int32(8)
					F_numericvar_deserialize(m, v93, v98)
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						F_numericvar_to_int128(m, v98, v88+int32(16))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							F_numericvar_deserialize(m, v93, v98)
							mBase = m.M
							v106 = m.ExcPending
							if v106 != 0 {
								return int32(0)
							} else {
								F_numericvar_to_int128(m, v98, v88+int32(32))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									F_pq_getmsgend(m, v93)
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
										if v113 != 0 {
											F_pfree(m, v113)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return int32(0)
											} else {
												m.G0 = v8 + int32(48)
												return v88
											}
										} else {
											m.G0 = v8 + int32(48)
											return v88
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
		v123 = m.ExcPending
		if v123 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_poly_deserialize_0), int32(0))
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_poly_deserialize_1), int32(_a_F_numeric_poly_deserialize_2), int32(_a_F_numeric_poly_deserialize_3))
				mBase = m.M
				v132 = m.ExcPending
				if v132 != 0 {
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
func F_numeric_send(m *base.Module, l0 int32) int32 {
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
	var v28 int32
	_ = v28
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum(m, v14)
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
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
	v20 = base.I32_extend16_s(v19)
	v22 = base.B2i32(int32(0) <= v20)
	if int32(0) <= v20 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = int32(-8)
	goto L5
L4:
	;
	v23 = int32(-6)
	goto L5
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if int32(0) <= v20 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
	v39 = v28
	goto L8
L7:
	;
	v39 = base.I32_extend16_s(v20<<(uint(int32(9))%32))>>(uint(int32(15))%32)&int32(-64) | v20&int32(63)
	goto L8
L8:
	;
	v41 = int32(base.Ui32(v23+int32(base.Ui32(v24)>>(uint(int32(2))%32))) >> (uint(int32(1)) % 32))
	v42 = int32(_a_F_numeric_send_0)
	v43 = v19 & v42
	if v43 != v42 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	F_pq_begintypsend(m, v10)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	if v43 != int32(_a_F_numeric_send_1) {
		v54 = v43
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v54 = v19 & int32(_a_F_numeric_send_2)
	goto L9
L13:
	;
	v54 = v19 << (uint(int32(1)) % 32) & int32(_a_F_numeric_send_3)
	goto L9
L14:
	;
	F_enlargeStringInfo(m, v10, int32(2))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v63 = int32(8)
	v69 = v41<<(uint(v63)%32) | int32(base.Ui32(v41&int32(_a_F_numeric_send_4))>>(uint(v63)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v60+v61))) = uint16(v69)
	v71 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v60 + v71
	F_enlargeStringInfo(m, v10, v71)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v80 = int32(8)
	v86 = v39<<(uint(v80)%32) | int32(base.Ui32(v39&int32(_a_F_numeric_send_4))>>(uint(v80)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v77+v78))) = uint16(v86)
	v88 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v77 + v88
	F_enlargeStringInfo(m, v10, v88)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v98 = int32(base.Ui32(v54) >> (uint(int32(8)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v94+v95))) = uint16(v98)
	v100 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v94 + v100
	F_enlargeStringInfo(m, v10, v100)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v106 = int32(0)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v117 = base.B2i32(v20 < v106)
	if v20 < v106 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v118 = int32(base.Ui32(v20)>>(uint(int32(7))%32)) & int32(63)
	goto L21
L20:
	;
	v118 = v20 & int32(_a_F_numeric_send_5)
	goto L21
L21:
	;
	v119 = int32(8)
	v123 = v118<<(uint(v119)%32) | int32(base.Ui32(v118)>>(uint(v119)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v107+v108))) = uint16(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v107 + int32(2)
	if v41 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v20 < v106 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v169))) = v170 << (uint(int32(2)) % 32)
	goto L32
L25:
	;
	v130 = int32(6)
	goto L27
L26:
	;
	v130 = int32(8)
	goto L27
L27:
	;
	v132 = v106
	goto L28
L28:
	;
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15+v130+v132<<(uint(int32(1))%32)))))
	F_enlargeStringInfo(m, v10, int32(2))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L24
L30:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v149 = int32(8)
	v153 = v142<<(uint(v149)%32) | int32(base.Ui32(v142)>>(uint(v149)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v146+v147))) = uint16(v153)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v146 + int32(2)
	v159 = v132 + int32(1)
	if v159 != v41 {
		v132 = v159
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	m.G0 = v10 + int32(16)
	return v169
}
func F_numeric_stddev_pop(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = Fn13951(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_numeric_sub(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = F_numeric_sub_opt_error(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
