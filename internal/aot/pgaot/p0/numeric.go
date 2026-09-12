package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_convert_numeric_to_scalar(m *base.Module, l0 int32, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v28 float64
	_ = v28
	var v35 int64
	_ = v35
	var v41 float64
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	if l1 <= int32(2201) {
		switch l1 - int32(16) {
		case 0:
			if l0 != 0 {
				v28 = float64(1)
			} else {
				v28 = float64(0)
			}
			return v28
		case 1, 2, 3, 6, 9:
			v45 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
			return float64(0)
		case 4:
			v35 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			return base.F64_convert_i64_s(v35)
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
				v41 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
				return v41
			default:
				if l1 == int32(1700) {
					v51 = F_DirectFunctionCall1Coll(m, int32(1508), int32(0), l0)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return float64(0)
					} else {
						v55 = *(*float64)(unsafe.Add(mBase, uint32(v51)))
						return v55
					}
				} else {
					v45 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
					return float64(0)
				}
			}
		}
	} else {
		if l1 <= int32(3768) {
			if base.Ui32(l1-int32(2202)) < base.Ui32(int32(5)) {
				return base.F64_convert_i32_u(l0)
			} else {
				if l1 == int32(3734) {
					return base.F64_convert_i32_u(l0)
				} else {
					v45 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
					return float64(0)
				}
			}
		} else {
			switch l1 - int32(4089) {
			case 0, 7:
				return base.F64_convert_i32_u(l0)
			case 1, 2, 3, 4, 5, 6:
				v45 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
				return float64(0)
			default:
				if l1 == int32(3769) {
					return base.F64_convert_i32_u(l0)
				} else {
					if l1 != int32(4191) {
						v45 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v45)
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
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
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
				v97 = F_jspGetNext(m, l1, v11+int32(44))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					if l5|v97 == int32(0) {
						v142 = v90
						m.G0 = v11 + int32(80)
						return v142
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
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
									v116 = F_executeItemOptUnwrapTarget(m, l0, v11+int32(44), v103, l5, v115)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										v142 = v116
										m.G0 = v11 + int32(80)
										return v142
									}
								} else {
									if l5 == int32(0) {
										v142 = v90
										m.G0 = v11 + int32(80)
										return v142
									} else {
										v120 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
										if v120 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v103
											*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v120
											*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v120
											*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v103
											v129 = F_list_make2_impl(m, v11+int32(40), v11+int32(36))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v129
												v142 = v90
												m.G0 = v11 + int32(80)
												return v142
											}
										} else {
											v134 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
											if v134 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l5))) = v103
												v142 = v90
												m.G0 = v11 + int32(80)
												return v142
											} else {
												v138 = F_lappend(m, v134, v103)
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v138
													v142 = v90
													m.G0 = v11 + int32(80)
													return v142
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
				v142 = int32(2)
				m.G0 = v11 + int32(80)
				return v142
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
							F_errmsg(m, int32(363395), v11)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(523842), int32(2311), int32(442731))
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
				F_errmsg_internal(m, int32(507113), v11+int32(32))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(523842), int32(1680), int32(26735))
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
					v142 = int32(2)
					m.G0 = v11 + int32(80)
					return v142
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
								F_errmsg(m, int32(363395), v11)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(523842), int32(2311), int32(442731))
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
						v142 = v61
						m.G0 = v11 + int32(80)
						return v142
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
						F_errmsg_internal(m, int32(30646), v11+int32(16))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523842), int32(3629), int32(389211))
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
				v97 = F_jspGetNext(m, l1, v11+int32(44))
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					if l5|v97 == int32(0) {
						v142 = v90
						m.G0 = v11 + int32(80)
						return v142
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
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
									v116 = F_executeItemOptUnwrapTarget(m, l0, v11+int32(44), v103, l5, v115)
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return int32(0)
									} else {
										v142 = v116
										m.G0 = v11 + int32(80)
										return v142
									}
								} else {
									if l5 == int32(0) {
										v142 = v90
										m.G0 = v11 + int32(80)
										return v142
									} else {
										v120 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
										if v120 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v103
											*(*int32)(unsafe.Add(mBase, uint32(v11)+76)) = v120
											*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v120
											*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v103
											v129 = F_list_make2_impl(m, v11+int32(40), v11+int32(36))
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l5))) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v129
												v142 = v90
												m.G0 = v11 + int32(80)
												return v142
											}
										} else {
											v134 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
											if v134 == int32(0) {
												*(*int32)(unsafe.Add(mBase, uint32(l5))) = v103
												v142 = v90
												m.G0 = v11 + int32(80)
												return v142
											} else {
												v138 = F_lappend(m, v134, v103)
												mBase = m.M
												v139 = m.ExcPending
												if v139 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l5)+4)) = v138
													v142 = v90
													m.G0 = v11 + int32(80)
													return v142
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
				v142 = int32(2)
				m.G0 = v11 + int32(80)
				return v142
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
							F_errmsg(m, int32(363395), v11)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(523842), int32(2311), int32(442731))
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
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	v3 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v12 = F_pg_detoast_datum_packed(m, l0)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v16 + int64(1)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v20&int32(1) != 0 {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v24 = int32(1)
			v25 = int32(base.Ui32(v20) >> (uint(v24) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v23))) = v25<<(uint(int32(2))%32) + int32(12)
			v36 = v25 - v24
			if v36 != 0 {
				v37 = F__emscripten_memcpy_bulkmem(m, v23+int32(4), v12+v24, v36)
				mBase = m.M
			} else {
			}
			v39 = v23
		} else {
			v39 = v12
		}
		v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
		v42 = base.I32_extend16_s(v41)
		if base.Ui32(int32(49152)) <= base.Ui32(v41) {
			if v42 == int32(-4096) {
				v50 = int32(2147483647)
			} else {
				v50 = int32(-2147483648)
			}
			if v42 == int32(-12288) {
				v53 = int32(-2147483647)
			} else {
				v53 = v50
			}
			v235 = v53
		} else {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
			v60 = base.B2i32(int32(0) <= v42)
			if int32(0) <= v42 {
				v61 = int32(-8)
			} else {
				v61 = int32(-6)
			}
			v62 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) + v61
			if int32(0) <= v42 {
				v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+6)))
				v73 = v63
			} else {
				v73 = v41<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v41&int32(63)
			}
			v79 = v41 & int32(49152)
			if v79 == int32(32768) {
				v82 = v41 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v82 = v79
			}
			if base.Ui32(v62) < base.Ui32(int32(2)) {
				v164 = int32(0)
			} else {
				if v73 < int32(-11) {
					v164 = int32(0)
				} else {
					if int32(20) < v73 {
						v164 = int32(2147483647)
					} else {
						if v42 < int32(0) {
							v96 = int32(6)
						} else {
							v96 = int32(8)
						}
						v97 = v39 + v96
						if base.Ui32(int32(4)) <= base.Ui32(v62) {
							v101 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+2)))
							v102 = v101
						} else {
							v102 = int32(0)
						}
						v104 = v73 << (uint(int32(2)) % 32)
						v105 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97))))
						if int32(1000) <= v105 {
							v109 = base.I32_div_s(v102, int32(10))
							v155 = base.I32_extend16_s(v109) + v105*int32(1000)
							v156 = v104 + int32(47)
						} else {
							if int32(100) <= v105 {
								v155 = v105*int32(10000) + v102
								v156 = v104 + int32(46)
							} else {
								if int32(10) <= v105 {
									if base.Ui32(int32(6)) <= base.Ui32(v62) {
										v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+4)))
										v129 = base.I32_div_s(v127, int32(1000))
										v131 = base.I32_extend16_s(v129)
									} else {
										v131 = v3
									}
									v155 = v131 + (v105*int32(100000) + v102*int32(10))
									v156 = v104 + int32(45)
								} else {
									if base.Ui32(int32(6)) <= base.Ui32(v62) {
										v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(v97)+4)))
										v144 = base.I32_div_s(v142, int32(100))
										v146 = base.I32_extend16_s(v144)
									} else {
										v146 = v3
									}
									v155 = v146 + (v105*int32(1000000) + v102*int32(100))
									v156 = v104 + int32(44)
								}
							}
						}
						v164 = v156<<(uint(int32(24))%32) | v155
					}
				}
			}
			if v82 != 0 {
				v167 = v164
			} else {
				v167 = int32(0) - v164
			}
			v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+16)))
			if v168 != int32(1) {
				v235 = v167
			} else {
				v171 = int32(24)
				v172 = v11 + v171
				v177 = int32(711645284)
				v180 = v167 - int32(1636608428) ^ v177 - int32(1455628627)
				v185 = v180 ^ int32(-1636608428) - base.I32_rotl(v180, int32(25))
				v190 = v185 ^ v177 - base.I32_rotl(v185, int32(16))
				v194 = v190 ^ v180 - base.I32_rotl(v190, int32(4))
				v198 = v194 ^ v185 - base.I32_rotl(v194, int32(14))
				v202 = v198 ^ v190 - base.I32_rotl(v198, v171)
				v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
				v207 = int32(32) - v206
				v208 = v202 << (uint(v206) % 32)
				if v208 != 0 {
					v215 = int32(32) - (base.I32_clz(v208) ^ int32(31))
					v216 = int32(255)
					if base.Ui32(v207&v216) < base.Ui32(v215&v216) {
						v221 = v207 + int32(1)
					} else {
						v221 = v215
					}
					v225 = v221
				} else {
					v225 = v207 + int32(1)
				}
				v226 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
				v228 = v226 + int32(base.Ui32(v202)>>(uint(v207)%32))
				v230 = v225 & int32(255)
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228))))
				if base.Ui32(v231) < base.Ui32(v230) {
					v233 = v230
				} else {
					v233 = v231
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v228))) = uint8(v233)
				v235 = v167
			}
		}
		if l0 != v12 {
			F_pfree(m, v12)
			mBase = m.M
			v243 = m.ExcPending
			if v243 != 0 {
				return int32(0)
			} else {
				return v235
			}
		} else {
			return v235
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
	var v153 int64
	_ = v153
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v15 = base.I32_extend16_s(v14)
	if base.Ui32(v14) <= base.Ui32(int32(49151)) {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v19 = base.I32_extend16_s(v18)
		if base.Ui32(int32(49151)) < base.Ui32(v18) {
			v47 = v19
			if v47&int32(65535) != int32(49152) {
				if v14 != int32(61440) {
					if v14 != int32(53248) {
						if v47&int32(65535) == int32(53248) {
							v94 = F_make_result_opt_error(m, int32(1770500), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v207 = v94
								m.G0 = v12 + int32(80)
								return v207
							}
						} else {
							v98 = F_make_result_opt_error(m, int32(1770524), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v207 = v98
								m.G0 = v12 + int32(80)
								return v207
							}
						}
					} else {
						if v47&int32(65535) == int32(61440) {
							v70 = F_make_result_opt_error(m, int32(1770476), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v207 = v70
								m.G0 = v12 + int32(80)
								return v207
							}
						} else {
							v74 = F_make_result_opt_error(m, int32(1770500), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v207 = v74
								m.G0 = v12 + int32(80)
								return v207
							}
						}
					}
				} else {
					if v47&int32(65535) == int32(53248) {
						v82 = F_make_result_opt_error(m, int32(1770476), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v207 = v82
							m.G0 = v12 + int32(80)
							return v207
						}
					} else {
						v86 = F_make_result_opt_error(m, int32(1770524), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v207 = v86
							m.G0 = v12 + int32(80)
							return v207
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(1770476), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v207 = v56
					m.G0 = v12 + int32(80)
					return v207
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
				v113 = v14 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v113
			v120 = v14 & int32(49152)
			if v120 == int32(32768) {
				v123 = v14 << (uint(int32(1)) % 32) & int32(16384)
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
			v153 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v153
			v156 = v12 + int32(24)
			*(*int64)(unsafe.Add(mBase, uint32(v156))) = v153
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v152
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v153
			v162 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v162
			v171 = base.B2i32(v19 < v162)
			if v19 < v162 {
				v172 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v172 = v18 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v172
			v179 = v18 & int32(49152)
			if v179 == int32(32768) {
				v182 = v18 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v182 = v179
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v182
			if v19 < v162 {
				v186 = int32(6)
			} else {
				v186 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l1 + v186
			F_add_var(m, v12+int32(56), v12+int32(32), v12+int32(8))
			mBase = m.M
			v196 = m.ExcPending
			if v196 != 0 {
				return int32(0)
			} else {
				v199 = F_make_result_opt_error(m, v12+int32(8), l2)
				mBase = m.M
				v200 = m.ExcPending
				if v200 != 0 {
					return int32(0)
				} else {
					v201 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
					if v201 == int32(0) {
						v207 = v199
						m.G0 = v12 + int32(80)
						return v207
					} else {
						F_pfree(m, v201)
						mBase = m.M
						v205 = m.ExcPending
						if v205 != 0 {
							return int32(0)
						} else {
							v207 = v199
							m.G0 = v12 + int32(80)
							return v207
						}
					}
				}
			}
		}
	} else {
		if v15 == int32(-16384) {
			v56 = F_make_result_opt_error(m, int32(1770476), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v207 = v56
				m.G0 = v12 + int32(80)
				return v207
			}
		} else {
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v47 = v45
			if v47&int32(65535) != int32(49152) {
				if v14 != int32(61440) {
					if v14 != int32(53248) {
						if v47&int32(65535) == int32(53248) {
							v94 = F_make_result_opt_error(m, int32(1770500), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v207 = v94
								m.G0 = v12 + int32(80)
								return v207
							}
						} else {
							v98 = F_make_result_opt_error(m, int32(1770524), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v207 = v98
								m.G0 = v12 + int32(80)
								return v207
							}
						}
					} else {
						if v47&int32(65535) == int32(61440) {
							v70 = F_make_result_opt_error(m, int32(1770476), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v207 = v70
								m.G0 = v12 + int32(80)
								return v207
							}
						} else {
							v74 = F_make_result_opt_error(m, int32(1770500), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v207 = v74
								m.G0 = v12 + int32(80)
								return v207
							}
						}
					}
				} else {
					if v47&int32(65535) == int32(53248) {
						v82 = F_make_result_opt_error(m, int32(1770476), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v207 = v82
							m.G0 = v12 + int32(80)
							return v207
						}
					} else {
						v86 = F_make_result_opt_error(m, int32(1770524), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v207 = v86
							m.G0 = v12 + int32(80)
							return v207
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(1770476), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v207 = v56
					m.G0 = v12 + int32(80)
					return v207
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
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	var v126 int32
	_ = v126
	var v130 int64
	_ = v130
	var v131 int32
	_ = v131
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
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
			v46 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v46
			v52 = int32(1)
			v53 = v42 + v52
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v56 = v54 & v52
			if v54 == v52 {
				v59 = int32(4)
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v61&int32(254) == int32(2) {
					v70 = v59
				} else {
					v70 = base.B2i32(v61 == int32(18)) << (uint(v59) % 32)
				}
				if v61 == int32(1) {
					v73 = v59
				} else {
					v73 = v70
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v56 != 0 {
					v84 = int32(base.Ui32(v54)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v84
			if v56 != 0 {
				v90 = v53
			} else {
				v90 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v90
			v93 = F_palloc0(m, int32(112))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v95)
				v98 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v98
				v102 = F_pq_getmsgint64(m, v8+int32(32))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v93)+8)) = v102
					F_numericvar_deserialize(m, v8+int32(32), v8+int32(8))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_accum_sum_add(m, v93+int32(16), v8+int32(8))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							v120 = F_pq_getmsgint(m, v8+int32(32), int32(4))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v93)+72)) = v120
								v125 = F_pq_getmsgint64(m, v8+int32(32))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v93)+80)) = v125
									v130 = F_pq_getmsgint64(m, v8+int32(32))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v93)+88)) = v130
										v135 = F_pq_getmsgint64(m, v8+int32(32))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v93)+96)) = v135
											v140 = F_pq_getmsgint64(m, v8+int32(32))
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
												return int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v93)+104)) = v140
												F_pq_getmsgend(m, v8+int32(32))
												mBase = m.M
												v146 = m.ExcPending
												if v146 != 0 {
													return int32(0)
												} else {
													v147 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
													if v147 != 0 {
														F_pfree(m, v147)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															m.G0 = v8 + int32(48)
															return v93
														}
													} else {
														m.G0 = v8 + int32(48)
														return v93
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
		v157 = m.ExcPending
		if v157 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66863), int32(0))
			mBase = m.M
			v161 = m.ExcPending
			if v161 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523822), int32(5383), int32(358580))
				mBase = m.M
				v166 = m.ExcPending
				if v166 != 0 {
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
	var v204 int64
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
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
			if base.Ui32(v21) <= base.Ui32(int32(49151)) {
				v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
				v26 = base.I32_extend16_s(v25)
				if base.Ui32(int32(49151)) < base.Ui32(v25) {
					v53 = v26
					if v53&int32(65535) != int32(49152) {
						if v21 != int32(61440) {
							if v21 != int32(53248) {
								v149 = F_make_result_opt_error(m, int32(1770548), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									v261 = v149
									m.G0 = v11 + int32(80)
									return v261
								}
							} else {
								if base.Ui32(int32(49152)) <= base.Ui32(v53&int32(65535)) {
									v75 = F_make_result_opt_error(m, int32(1770476), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v261 = v75
										m.G0 = v11 + int32(80)
										return v261
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
										v274 = m.ExcPending
										if v274 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(33816706))
											mBase = m.M
											v277 = m.ExcPending
											if v277 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(251230), int32(0))
												mBase = m.M
												v281 = m.ExcPending
												if v281 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(523822), int32(3403), int32(511306))
													mBase = m.M
													v286 = m.ExcPending
													if v286 != 0 {
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
										v94 = v53 & int32(49152)
										if v94 == int32(32768) {
											v97 = v53 << (uint(int32(1)) % 32) & int32(16384)
										} else {
											v97 = v94
										}
										if v97 != int32(16384) {
											v102 = F_make_result_opt_error(m, int32(1770500), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v261 = v102
												m.G0 = v11 + int32(80)
												return v261
											}
										} else {
											v106 = F_make_result_opt_error(m, int32(1770524), int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												v261 = v106
												m.G0 = v11 + int32(80)
												return v261
											}
										}
									}
								}
							}
						} else {
							if base.Ui32(int32(49152)) <= base.Ui32(v53&int32(65535)) {
								v114 = F_make_result_opt_error(m, int32(1770476), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v261 = v114
									m.G0 = v11 + int32(80)
									return v261
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
									v290 = m.ExcPending
									if v290 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(33816706))
										mBase = m.M
										v293 = m.ExcPending
										if v293 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(251230), int32(0))
											mBase = m.M
											v297 = m.ExcPending
											if v297 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(523822), int32(3421), int32(511306))
												mBase = m.M
												v302 = m.ExcPending
												if v302 != 0 {
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
									v133 = v53 & int32(49152)
									if v133 == int32(32768) {
										v136 = v53 << (uint(int32(1)) % 32) & int32(16384)
									} else {
										v136 = v133
									}
									if v136 != int32(16384) {
										v141 = F_make_result_opt_error(m, int32(1770524), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											v261 = v141
											m.G0 = v11 + int32(80)
											return v261
										}
									} else {
										v145 = F_make_result_opt_error(m, int32(1770500), int32(0))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return int32(0)
										} else {
											v261 = v145
											m.G0 = v11 + int32(80)
											return v261
										}
									}
								}
							}
						}
					} else {
						v63 = F_make_result_opt_error(m, int32(1770476), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v261 = v63
							m.G0 = v11 + int32(80)
							return v261
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
						v164 = v21 & int32(16383)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v164
					v171 = v21 & int32(49152)
					if v171 == int32(32768) {
						v174 = v21 << (uint(int32(1)) % 32) & int32(16384)
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
					v204 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v204
					v207 = v11 + int32(24)
					*(*int64)(unsafe.Add(mBase, uint32(v207))) = v204
					*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v203
					v211 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v211
					*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v204
					v218 = base.B2i32(v26 < v211)
					if v26 < v211 {
						v219 = int32(6)
					} else {
						v219 = int32(8)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+52)) = v19 + v219
					if v26 < v211 {
						v228 = int32(base.Ui32(v25)>>(uint(int32(7))%32)) & int32(63)
					} else {
						v228 = v25 & int32(16383)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+44)) = v228
					v235 = v25 & int32(49152)
					if v235 == int32(32768) {
						v238 = v25 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v238 = v235
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v238
					v246 = int32(0)
					F_div_var(m, v11+int32(56), v11+int32(32), v11+int32(8), v246, v246, int32(1))
					mBase = m.M
					v250 = m.ExcPending
					if v250 != 0 {
						return int32(0)
					} else {
						v254 = F_make_result_opt_error(m, v11+int32(8), int32(0))
						mBase = m.M
						v255 = m.ExcPending
						if v255 != 0 {
							return int32(0)
						} else {
							v256 = *(*int32)(unsafe.Add(mBase, uint32(v207)))
							if v256 == int32(0) {
								v261 = v254
								m.G0 = v11 + int32(80)
								return v261
							} else {
								F_pfree(m, v256)
								mBase = m.M
								v260 = m.ExcPending
								if v260 != 0 {
									return int32(0)
								} else {
									v261 = v254
									m.G0 = v11 + int32(80)
									return v261
								}
							}
						}
					}
				}
			} else {
				if v22 == int32(-16384) {
					v63 = F_make_result_opt_error(m, int32(1770476), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v261 = v63
						m.G0 = v11 + int32(80)
						return v261
					}
				} else {
					v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
					v53 = v52
					if v53&int32(65535) != int32(49152) {
						if v21 != int32(61440) {
							if v21 != int32(53248) {
								v149 = F_make_result_opt_error(m, int32(1770548), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									v261 = v149
									m.G0 = v11 + int32(80)
									return v261
								}
							} else {
								if base.Ui32(int32(49152)) <= base.Ui32(v53&int32(65535)) {
									v75 = F_make_result_opt_error(m, int32(1770476), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v261 = v75
										m.G0 = v11 + int32(80)
										return v261
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
										v274 = m.ExcPending
										if v274 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(33816706))
											mBase = m.M
											v277 = m.ExcPending
											if v277 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(251230), int32(0))
												mBase = m.M
												v281 = m.ExcPending
												if v281 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(523822), int32(3403), int32(511306))
													mBase = m.M
													v286 = m.ExcPending
													if v286 != 0 {
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
										v94 = v53 & int32(49152)
										if v94 == int32(32768) {
											v97 = v53 << (uint(int32(1)) % 32) & int32(16384)
										} else {
											v97 = v94
										}
										if v97 != int32(16384) {
											v102 = F_make_result_opt_error(m, int32(1770500), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int32(0)
											} else {
												v261 = v102
												m.G0 = v11 + int32(80)
												return v261
											}
										} else {
											v106 = F_make_result_opt_error(m, int32(1770524), int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												v261 = v106
												m.G0 = v11 + int32(80)
												return v261
											}
										}
									}
								}
							}
						} else {
							if base.Ui32(int32(49152)) <= base.Ui32(v53&int32(65535)) {
								v114 = F_make_result_opt_error(m, int32(1770476), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v261 = v114
									m.G0 = v11 + int32(80)
									return v261
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
									v290 = m.ExcPending
									if v290 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(33816706))
										mBase = m.M
										v293 = m.ExcPending
										if v293 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(251230), int32(0))
											mBase = m.M
											v297 = m.ExcPending
											if v297 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(523822), int32(3421), int32(511306))
												mBase = m.M
												v302 = m.ExcPending
												if v302 != 0 {
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
									v133 = v53 & int32(49152)
									if v133 == int32(32768) {
										v136 = v53 << (uint(int32(1)) % 32) & int32(16384)
									} else {
										v136 = v133
									}
									if v136 != int32(16384) {
										v141 = F_make_result_opt_error(m, int32(1770524), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											v261 = v141
											m.G0 = v11 + int32(80)
											return v261
										}
									} else {
										v145 = F_make_result_opt_error(m, int32(1770500), int32(0))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return int32(0)
										} else {
											v261 = v145
											m.G0 = v11 + int32(80)
											return v261
										}
									}
								}
							}
						}
					} else {
						v63 = F_make_result_opt_error(m, int32(1770476), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							v261 = v63
							m.G0 = v11 + int32(80)
							return v261
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
	var v48 int32
	_ = v48
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
	var v79 int32
	_ = v79
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
	var v107 int32
	_ = v107
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
	v23 = F_make_result_opt_error(m, int32(1770600), int32(0))
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
	v48 = v36 + int32(12)
	v52 = v16
	goto L13
L13:
	;
	v57 = v48 - int32(2)
	v59 = base.I64_div_u_s(v52, int64(10000))
	v62 = v59*int64(55536) + v52
	*(*uint16)(unsafe.Add(mBase, uint32(v57))) = uint16(v62)
	v65 = v46 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v52) {
		v46 = v65
		v48 = v57
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
	if v16 <= int64(2) {
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
	v79 = int32(0)
	v83 = v16
	goto L21
L20:
	;
	v142 = v74
	goto L16
L21:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	if v79 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	F_pfree(m, v79)
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
	v107 = v93 + int32(12)
	v111 = v104
	goto L32
L32:
	;
	v116 = v107 - int32(2)
	v118 = base.I64_div_u_s(v111, int64(10000))
	v121 = v118*int64(55536) + v111
	*(*uint16)(unsafe.Add(mBase, uint32(v116))) = uint16(v121)
	v124 = v105 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v111) {
		v105 = v124
		v107 = v116
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
	if int64(3) < v83 {
		v79 = v93
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
	F_errmsg(m, int32(473153), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(523822), int32(3753), int32(514298))
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
	F_errmsg(m, int32(119621), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(523822), int32(3763), int32(514298))
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
		if base.Ui32(int32(49152)) <= base.Ui32(v8) {
			if v8 != int32(61440) {
				if v8 != int32(53248) {
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
				v33 = F_DirectFunctionCall1Coll(m, int32(1469), v28, v31)
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
		if base.Ui32(v14) <= base.Ui32(int32(49151)) {
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
				v83 = v47 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v83
			v90 = v47 & int32(49152)
			if v90 == int32(32768) {
				v93 = v47 << (uint(int32(1)) % 32) & int32(16384)
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
									F_errmsg(m, int32(420694), int32(0))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523822), int32(4558), int32(222474))
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
								F_errmsg(m, int32(420694), int32(0))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(523822), int32(4558), int32(222474))
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
					if v14 == int32(49152) {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(236055)
						F_errmsg(m, int32(193475), v10)
						mBase = m.M
						v144 = m.ExcPending
						if v144 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523822), int32(4536), int32(222474))
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
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(236055)
						F_errmsg(m, int32(192798), v8+int32(-48))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523822), int32(4540), int32(222474))
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
		if base.Ui32(v42) < base.Ui32(int32(49152)) {
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
				v83 = v47 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v83
			v90 = v47 & int32(49152)
			if v90 == int32(32768) {
				v93 = v47 << (uint(int32(1)) % 32) & int32(16384)
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
									F_errmsg(m, int32(420694), int32(0))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(523822), int32(4558), int32(222474))
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
								F_errmsg(m, int32(420694), int32(0))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(523822), int32(4558), int32(222474))
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		v17 = base.I32_extend16_s(v16)
		if base.Ui32(int32(49152)) <= base.Ui32(v16) {
			if v17 == int32(-4096) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352583810))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(239569), int32(0))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(523822), int32(3951), int32(287818))
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
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
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v25 = F_palloc(m, int32(base.Ui32(v22)>>(uint(int32(2))%32)))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v29 = int32(base.Ui32(v27) >> (uint(int32(2)) % 32))
					if v29 != 0 {
						v30 = F__emscripten_memcpy_bulkmem(m, v25, v12, v29)
						mBase = m.M
					} else {
					}
					v115 = v25
					m.G0 = v9 + int32(48)
					return v115
				}
			}
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v38 = base.B2i32(int32(0) <= v17)
			if int32(0) <= v17 {
				v39 = int32(-8)
			} else {
				v39 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(int32(base.Ui32(v32)>>(uint(int32(2))%32))+v39) >> (uint(int32(1)) % 32))
			if int32(0) <= v17 {
				v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v54 = v44
			} else {
				v54 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
			}
			v55 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v55
			v58 = v9 + int32(16)
			*(*int64)(unsafe.Add(mBase, uint32(v58))) = v55
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v54
			v62 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v62
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v55
			v69 = base.B2i32(v17 < v62)
			if v17 < v62 {
				v70 = int32(6)
			} else {
				v70 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v12 + v70
			v78 = v16 & int32(49152)
			if v78 == int32(32768) {
				v81 = v16 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v81 = v78
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v81
			if v17 < v62 {
				v89 = int32(base.Ui32(v16)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v89 = v16 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v89
			v92 = v9 + int32(24)
			v97 = F_estimate_ln_dweight(m, v92)
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return int32(0)
			} else {
				v99 = int32(16) - v97
				if v89 < v99 {
					v101 = v99
				} else {
					v101 = v89
				}
				if int32(1000) <= v101 {
					v104 = int32(1000)
				} else {
					v104 = v101
				}
				F_ln_var(m, v92, v9, v104)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					v108 = F_make_result_opt_error(m, v9, int32(0))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
						if v110 == int32(0) {
							v115 = v108
							m.G0 = v9 + int32(48)
							return v115
						} else {
							F_pfree(m, v110)
							mBase = m.M
							v114 = m.ExcPending
							if v114 != 0 {
								return int32(0)
							} else {
								v115 = v108
								m.G0 = v9 + int32(48)
								return v115
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
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
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
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v30 = int32(49152)
			v31 = v29 & v30
			if v31 == v30 {
				if v29 != int32(53248) {
					if v29 != int32(49152) {
						if v28 != int32(61440) {
							v50 = int32(-1)
						} else {
							v50 = int32(0)
						}
						v169 = v50
					} else {
						v169 = base.B2i32(v28 != int32(49152))
					}
				} else {
					if v28 == int32(49152) {
						v45 = int32(-1)
					} else {
						v45 = base.B2i32(v28 != int32(53248))
					}
					v169 = v45
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v28) {
					if v28 == int32(61440) {
						v57 = int32(1)
					} else {
						v57 = int32(-1)
					}
					v169 = v57
				} else {
					v59 = v7 + int32(6)
					v64 = base.B2i32(int32(0) <= base.I32_extend16_s(v29))
					if int32(0) <= base.I32_extend16_s(v29) {
						v65 = int32(-8)
					} else {
						v65 = int32(-6)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if int32(0) <= base.I32_extend16_s(v29) {
						v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59))))
						v79 = v69
					} else {
						v79 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
					}
					v80 = v65 + int32(base.Ui32(v66)>>(uint(int32(2))%32))
					v82 = v14 + int32(6)
					v87 = base.B2i32(int32(0) <= base.I32_extend16_s(v28))
					if int32(0) <= base.I32_extend16_s(v28) {
						v88 = int32(-8)
					} else {
						v88 = int32(-6)
					}
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if int32(0) <= base.I32_extend16_s(v28) {
						v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
						v102 = v92
					} else {
						v102 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
					}
					v103 = v88 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
					v109 = v28 & int32(49152)
					if v109 == int32(32768) {
						v112 = v28 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v112 = v109
					}
					if base.Ui32(v80) <= base.Ui32(int32(1)) {
						if base.Ui32(v103) < base.Ui32(int32(2)) {
							v169 = int32(0)
						} else {
							if v112 == int32(16384) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v169 = v122
						}
					} else {
						if v31 == int32(32768) {
							v129 = v29 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v129 = v31
						}
						if base.Ui32(v103) <= base.Ui32(int32(1)) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v169 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v29) {
								v137 = v7 + int32(8)
							} else {
								v137 = v59
							}
							v139 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v28) {
								v142 = v14 + int32(8)
							} else {
								v142 = v82
							}
							v144 = int32(base.Ui32(v103) >> (uint(int32(1)) % 32))
							if v129 == int32(0) {
								if v112 == int32(16384) {
									v169 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v137, v139, v79, v142, v144, v102)
									mBase = m.M
									v169 = v150
								}
							} else {
								if v112 == int32(0) {
									v169 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v142, v144, v102, v137, v139, v79)
									mBase = m.M
									v169 = v154
								}
							}
						}
					}
				}
			}
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v170 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v174 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							return int32(base.Ui32(v169) >> (uint(int32(31)) % 32))
						}
					} else {
						return int32(base.Ui32(v169) >> (uint(int32(31)) % 32))
					}
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v174 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						return int32(base.Ui32(v169) >> (uint(int32(31)) % 32))
					}
				} else {
					return int32(base.Ui32(v169) >> (uint(int32(31)) % 32))
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
	var v129 int32
	_ = v129
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v142 int32
	_ = v142
	var v145 int64
	_ = v145
	var v150 int32
	_ = v150
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
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
			v150 = v49
			m.G0 = v11 + int32(16)
			return v150
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v51 == int32(0) {
				v150 = v49
				m.G0 = v11 + int32(16)
				return v150
			} else {
				if v49 == int32(0) {
					v56 = int32(4554128)
					v57 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v59
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
						v176 = m.ExcPending
						if v176 != 0 {
							return int32(0)
						} else {
							F_errmsg_internal(m, int32(66863), int32(0))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(523822), int32(5606), int32(370904))
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
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
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v96
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
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v57
							v150 = v99
							m.G0 = v11 + int32(16)
							return v150
						}
					}
				} else {
					v117 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
					if v117 <= int64(0) {
						v150 = v49
					} else {
						v120 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v120 + v117
						v123 = *(*int64)(unsafe.Add(mBase, uint32(v51)+24))
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
						v125 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
						v126 = v124 + v125
						*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v126
						v129 = v49 + int32(24)
						v132 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
						*(*int64)(unsafe.Add(mBase, uint32(v129))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v126) < base.Ui64(v124))) + (v123 + v132)
						v136 = *(*int64)(unsafe.Add(mBase, uint32(v51)+40))
						v137 = *(*int64)(unsafe.Add(mBase, uint32(v49)+32))
						v138 = *(*int64)(unsafe.Add(mBase, uint32(v51)+32))
						v139 = v137 + v138
						*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v139
						v142 = v49 + int32(40)
						v145 = *(*int64)(unsafe.Add(mBase, uint32(v142)))
						*(*int64)(unsafe.Add(mBase, uint32(v142))) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v139) < base.Ui64(v137))) + (v136 + v145)
						v150 = v49
					}
					m.G0 = v11 + int32(16)
					return v150
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v163 = m.ExcPending
		if v163 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66863), int32(0))
			mBase = m.M
			v167 = m.ExcPending
			if v167 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523822), int32(5743), int32(391244))
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
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
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int64
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
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
			v46 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v46
			v52 = int32(1)
			v53 = v42 + v52
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v56 = v54 & v52
			if v54 == v52 {
				v59 = int32(4)
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v61&int32(254) == int32(2) {
					v70 = v59
				} else {
					v70 = base.B2i32(v61 == int32(18)) << (uint(v59) % 32)
				}
				if v61 == int32(1) {
					v73 = v59
				} else {
					v73 = v70
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v56 != 0 {
					v84 = int32(base.Ui32(v54)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v84
			if v56 != 0 {
				v90 = v53
			} else {
				v90 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v90
			v93 = F_palloc0(m, int32(48))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v95)
				v99 = F_pq_getmsgint64(m, v8+int32(32))
				mBase = m.M
				v100 = m.ExcPending
				if v100 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v93)+8)) = v99
					F_numericvar_deserialize(m, v8+int32(32), v8+int32(8))
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						F_numericvar_to_int128(m, v8+int32(8), v93+int32(16))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							F_numericvar_deserialize(m, v8+int32(32), v8+int32(8))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								F_numericvar_to_int128(m, v8+int32(8), v93+int32(32))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int32(0)
								} else {
									F_pq_getmsgend(m, v8+int32(32))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
										if v130 != 0 {
											F_pfree(m, v130)
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												m.G0 = v8 + int32(48)
												return v93
											}
										} else {
											m.G0 = v8 + int32(48)
											return v93
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
		v140 = m.ExcPending
		if v140 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66863), int32(0))
			mBase = m.M
			v144 = m.ExcPending
			if v144 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(523822), int32(5866), int32(358415))
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
		}
	}
}
func F_numeric_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v21 = base.I32_extend16_s(v20)
	v23 = base.B2i32(int32(0) <= v21)
	if int32(0) <= v21 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v24 = int32(-8)
	goto L5
L4:
	;
	v24 = int32(-6)
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v28 = v24 + int32(base.Ui32(v25)>>(uint(int32(2))%32))
	if int32(0) <= v21 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+6)))
	v40 = v29
	goto L8
L7:
	;
	v40 = base.I32_extend16_s(v21<<(uint(int32(9))%32))>>(uint(int32(15))%32)&int32(-64) | v21&int32(63)
	goto L8
L8:
	;
	v42 = int32(base.Ui32(v28) >> (uint(int32(1)) % 32))
	v43 = int32(49152)
	v44 = v20 & v43
	if v44 != v43 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	F_pq_begintypsend(m, v11)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L14
	}
L10:
	;
	if v44 != int32(32768) {
		v55 = v44
		goto L9
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v55 = v20 & int32(61440)
	goto L9
L13:
	;
	v55 = v20 << (uint(int32(1)) % 32) & int32(16384)
	goto L9
L14:
	;
	F_enlargeStringInfo(m, v11, int32(2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v64 = int32(8)
	v70 = v42<<(uint(v64)%32) | int32(base.Ui32(v42&int32(65280))>>(uint(v64)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v61+v62))) = uint16(v70)
	v72 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v61 + v72
	F_enlargeStringInfo(m, v11, v72)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v81 = int32(8)
	v87 = v40<<(uint(v81)%32) | int32(base.Ui32(v40&int32(65280))>>(uint(v81)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v78+v79))) = uint16(v87)
	v89 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v78 + v89
	F_enlargeStringInfo(m, v11, v89)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v99 = int32(base.Ui32(v55) >> (uint(int32(8)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v95+v96))) = uint16(v99)
	v101 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v95 + v101
	F_enlargeStringInfo(m, v11, v101)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v107 = int32(0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v118 = base.B2i32(v21 < v107)
	if v21 < v107 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v119 = int32(base.Ui32(v21)>>(uint(int32(7))%32)) & int32(63)
	goto L21
L20:
	;
	v119 = v21 & int32(16383)
	goto L21
L21:
	;
	v120 = int32(8)
	v124 = v119<<(uint(v120)%32) | int32(base.Ui32(v119)>>(uint(v120)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v108+v109))) = uint16(v124)
	v126 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v108 + v126
	if base.Ui32(v126) <= base.Ui32(v28) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v21 < v107 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v178))) = v179 << (uint(int32(2)) % 32)
	goto L35
L25:
	;
	v133 = int32(6)
	goto L27
L26:
	;
	v133 = int32(8)
	goto L27
L27:
	;
	v135 = int32(1)
	if base.Ui32(v42) <= base.Ui32(v135) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v138 = v135
	goto L30
L29:
	;
	v138 = v42
	goto L30
L30:
	;
	v139 = v107
	goto L31
L31:
	;
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16+v133+v139<<(uint(int32(1))%32)))))
	F_enlargeStringInfo(m, v11, int32(2))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L33
	}
L32:
	;
	goto L24
L33:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v157 = int32(8)
	v161 = v150<<(uint(v157)%32) | int32(base.Ui32(v150)>>(uint(v157)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v154+v155))) = uint16(v161)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v154 + int32(2)
	v167 = v139 + int32(1)
	if v167 != v138 {
		v139 = v167
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	m.G0 = v11 + int32(16)
	return v178
}
func F_numeric_stddev_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		v10 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	}
	v11 = int32(0)
	v15 = F_numeric_stddev_internal(m, v10, v11, v11, v6+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
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
