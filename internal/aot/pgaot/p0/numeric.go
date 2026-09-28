package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_convert_numeric_to_scalar(m *base.Module, l0 int64, l1 int32, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v32 float64
	_ = v32
	var v56 int32
	_ = v56
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	if l1 <= int32(2201) {
		switch l1 - int32(16) {
		case 0:
			if l0 != int64(0) {
				v32 = float64(1)
			} else {
				v32 = float64(0)
			}
			return v32
		case 1, 2, 3, 6, 9:
			v56 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
			return float64(0)
		case 4:
			return base.F64_convert_i64_s(l0)
		case 5:
			return base.F64_convert_i32_s(base.I32_extend16_s(base.I32_wrap_i64(l0)))
		case 7:
			return base.F64_convert_i32_s(base.I32_wrap_i64(l0))
		case 8, 10:
			return base.F64_convert_i32_u(base.I32_wrap_i64(l0))
		default:
			switch l1 - int32(700) {
			case 0:
				return base.F64_promote_f32(base.F32_reinterpret_i32(base.I32_wrap_i64(l0)))
			case 1:
				return base.F64_reinterpret_i64(l0)
			default:
				if l1 == int32(1700) {
					v62 = F_DirectFunctionCall1Coll(m, int32(1706), int32(0), l0)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return float64(0)
					} else {
						return base.F64_reinterpret_i64(v62)
					}
				} else {
					v56 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
					return float64(0)
				}
			}
		}
	} else {
		if l1 <= int32(3768) {
			if base.B2i32(l1 == int32(3734))|base.B2i32(base.Ui32(l1-int32(2202)) < base.Ui32(int32(5))) != 0 {
				return base.F64_convert_i32_u(base.I32_wrap_i64(l0))
			} else {
				v56 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
				return float64(0)
			}
		} else {
			if l1 <= int32(_a_F_convert_numeric_to_scalar_0) {
				switch l1 - int32(4089) {
				case 0, 7:
					return base.F64_convert_i32_u(base.I32_wrap_i64(l0))
				case 1, 2, 3, 4, 5, 6:
					v56 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
					return float64(0)
				default:
					if l1 != int32(3769) {
						v56 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
						return float64(0)
					} else {
						return base.F64_convert_i32_u(base.I32_wrap_i64(l0))
					}
				}
			} else {
				if base.B2i32(l1 == int32(_a_F_convert_numeric_to_scalar_1))|base.B2i32(l1 == int32(_a_F_convert_numeric_to_scalar_2)) != 0 {
					return base.F64_convert_i32_u(base.I32_wrap_i64(l0))
				} else {
					v56 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v56)
					return float64(0)
				}
			}
		}
	}
}
func F_executeNumericItemMethod(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int64
	_ = v155
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v161 int64
	_ = v161
	var v168 int32
	_ = v168
	v10 = m.G0
	v12 = v10 - int32(112)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if l3 != 0 {
		switch v14 - int32(2) {
		case 0:
			v91 = int32(0)
			v93 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)))
			v94 = F_DirectFunctionCall1Coll(m, l4, v91, v93)
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				v97 = v12 + int32(84)
				v98 = F_jspGetNext(m, l1, v97)
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					if v98|l5 == int32(0) {
						v168 = v91
						m.G0 = v12 + int32(112)
						return v168
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(2)
						v106 = F_pg_detoast_datum(m, base.I32_wrap_i64(v94))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v106
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							if int32(0) < v109 {
								v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
								v115 = F_executeItemOptUnwrapTarget(m, l0, v97, v12+int32(48), l5, v114)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									v168 = v115
									m.G0 = v12 + int32(112)
									return v168
								}
							} else {
								if l5 == int32(0) {
									v168 = v91
									m.G0 = v12 + int32(112)
									return v168
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
									if v120 < v121 {
										v125 = v119 + v120<<(uint(int32(5))%32)
										v126 = *(*int64)(unsafe.Add(mBase, uint32(v12)+72))
										*(*int64)(unsafe.Add(mBase, uint32(v125)+40)) = v126
										v128 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v125)+32)) = v128
										v130 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v125)+24)) = v130
										v132 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v125)+16)) = v132
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
										*(*int32)(unsafe.Add(mBase, uint32(v119))) = v134 + int32(1)
										v168 = v91
										m.G0 = v12 + int32(112)
										return v168
									} else {
										v138 = int32(16)
										v140 = v121 << (uint(int32(1)) % 32)
										if v140 <= v138 {
											v143 = v138
										} else {
											v143 = v140
										}
										v148 = F_palloc(m, v143<<(uint(int32(5))%32)|int32(16))
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v148)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v143
											*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(1)
											v155 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v148)+16)) = v155
											v157 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v148)+24)) = v157
											v159 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
											*(*int64)(unsafe.Add(mBase, uint32(v148)+32)) = v159
											v161 = *(*int64)(unsafe.Add(mBase, uint32(v12)+72))
											*(*int64)(unsafe.Add(mBase, uint32(v148)+40)) = v161
											*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v148
											*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v148
											v168 = v91
											m.G0 = v12 + int32(112)
											return v168
										}
									}
								}
							}
						}
					}
				}
			}
		default:
			v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
			if v69 != int32(1) {
				v168 = int32(2)
				m.G0 = v12 + int32(112)
				return v168
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101449858))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v80 = F_jspOperationName(m, v79)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v80
							F_errmsg(m, int32(_a_F_executeNumericItemMethod_0), v12)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(2407), int32(_a_F_executeNumericItemMethod_2))
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
					}
				}
			}
		case 14:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v45
				F_errmsg_internal(m, int32(_a_F_executeNumericItemMethod_3), v12+int32(32))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(1707), int32(_a_F_executeNumericItemMethod_4))
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
		case 16:
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			if v18&int32(536870912) != 0 {
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
				if v69 != int32(1) {
					v168 = int32(2)
					m.G0 = v12 + int32(112)
					return v168
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(101449858))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v80 = F_jspOperationName(m, v79)
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v80
								F_errmsg(m, int32(_a_F_executeNumericItemMethod_0), v12)
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(2407), int32(_a_F_executeNumericItemMethod_2))
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
						}
					}
				}
			} else {
				if v18&int32(1073741824) != 0 {
					v57 = int32(1)
					v60 = int32(0)
					v62 = F_executeAnyItem(m, l0, l1, v17, l5, v57, v57, v57, v60, v60)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v168 = v62
						m.G0 = v12 + int32(112)
						return v168
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v29
						F_errmsg_internal(m, int32(_a_F_executeNumericItemMethod_5), v12+int32(16))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(3947), int32(_a_F_executeNumericItemMethod_6))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
		if v14 == int32(2) {
			v91 = int32(0)
			v93 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l2)+8)))
			v94 = F_DirectFunctionCall1Coll(m, l4, v91, v93)
			mBase = m.M
			v95 = m.ExcPending
			if v95 != 0 {
				return int32(0)
			} else {
				v97 = v12 + int32(84)
				v98 = F_jspGetNext(m, l1, v97)
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return int32(0)
				} else {
					if v98|l5 == int32(0) {
						v168 = v91
						m.G0 = v12 + int32(112)
						return v168
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(2)
						v106 = F_pg_detoast_datum(m, base.I32_wrap_i64(v94))
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v106
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
							if int32(0) < v109 {
								v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
								v115 = F_executeItemOptUnwrapTarget(m, l0, v97, v12+int32(48), l5, v114)
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int32(0)
								} else {
									v168 = v115
									m.G0 = v12 + int32(112)
									return v168
								}
							} else {
								if l5 == int32(0) {
									v168 = v91
									m.G0 = v12 + int32(112)
									return v168
								} else {
									v119 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
									v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
									v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
									if v120 < v121 {
										v125 = v119 + v120<<(uint(int32(5))%32)
										v126 = *(*int64)(unsafe.Add(mBase, uint32(v12)+72))
										*(*int64)(unsafe.Add(mBase, uint32(v125)+40)) = v126
										v128 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v125)+32)) = v128
										v130 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
										*(*int64)(unsafe.Add(mBase, uint32(v125)+24)) = v130
										v132 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
										*(*int64)(unsafe.Add(mBase, uint32(v125)+16)) = v132
										v134 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
										*(*int32)(unsafe.Add(mBase, uint32(v119))) = v134 + int32(1)
										v168 = v91
										m.G0 = v12 + int32(112)
										return v168
									} else {
										v138 = int32(16)
										v140 = v121 << (uint(int32(1)) % 32)
										if v140 <= v138 {
											v143 = v138
										} else {
											v143 = v140
										}
										v148 = F_palloc(m, v143<<(uint(int32(5))%32)|int32(16))
										mBase = m.M
										v149 = m.ExcPending
										if v149 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v148)+8)) = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v143
											*(*int32)(unsafe.Add(mBase, uint32(v148))) = int32(1)
											v155 = *(*int64)(unsafe.Add(mBase, uint32(v12)+48))
											*(*int64)(unsafe.Add(mBase, uint32(v148)+16)) = v155
											v157 = *(*int64)(unsafe.Add(mBase, uint32(v12)+56))
											*(*int64)(unsafe.Add(mBase, uint32(v148)+24)) = v157
											v159 = *(*int64)(unsafe.Add(mBase, uint32(v12)+64))
											*(*int64)(unsafe.Add(mBase, uint32(v148)+32)) = v159
											v161 = *(*int64)(unsafe.Add(mBase, uint32(v12)+72))
											*(*int64)(unsafe.Add(mBase, uint32(v148)+40)) = v161
											*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = v148
											*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v148
											v168 = v91
											m.G0 = v12 + int32(112)
											return v168
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+34)))
			if v69 != int32(1) {
				v168 = int32(2)
				m.G0 = v12 + int32(112)
				return v168
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(101449858))
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v80 = F_jspOperationName(m, v79)
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v80
							F_errmsg(m, int32(_a_F_executeNumericItemMethod_0), v12)
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_executeNumericItemMethod_1), int32(2407), int32(_a_F_executeNumericItemMethod_2))
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
					}
				}
			}
		}
	}
}
func F_make_numeric_typmod_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if base.Ui32(l0-int32(1001)) <= base.Ui32(int32(-1001)) {
		v14 = int32(-1)
		v15 = F_errsave_start(m, l2)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				v67 = v14
				m.G0 = v8 + int32(32)
				return v67
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(1000)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg(m, int32(_a_F_make_numeric_typmod_safe_0), v8)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errsave_finish(m, l2, int32(_a_F_make_numeric_typmod_safe_1), int32(1321), int32(_a_F_make_numeric_typmod_safe_2))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							v67 = v14
							m.G0 = v8 + int32(32)
							return v67
						}
					}
				}
			}
		}
	} else {
		if base.Ui32(l1-int32(1001)) <= base.Ui32(int32(-2002)) {
			v39 = int32(-1)
			v40 = F_errsave_start(m, l2)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				if v40 == int32(0) {
					v67 = v39
					m.G0 = v8 + int32(32)
					return v67
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = int64(4299262262296)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
						F_errmsg(m, int32(_a_F_make_numeric_typmod_safe_3), v8+int32(16))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l2, int32(_a_F_make_numeric_typmod_safe_1), int32(1326), int32(_a_F_make_numeric_typmod_safe_2))
							mBase = m.M
							v59 = m.ExcPending
							if v59 != 0 {
								return int32(0)
							} else {
								v67 = v39
								m.G0 = v8 + int32(32)
								return v67
							}
						}
					}
				}
			}
		} else {
			v67 = l1&int32(2047) | l0<<(uint(int32(16))%32) + int32(4)
			m.G0 = v8 + int32(32)
			return v67
		}
	}
}
func F_numeric_abbrev_convert(m *base.Module, l0 int64, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v129 int64
	_ = v129
	var v132 int64
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v211 int32
	_ = v211
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v11 = base.I32_wrap_i64(l0)
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v16 + int64(1)
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
		if v20&int32(1) == int32(0) {
			v42 = v12
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			v26 = int32(1)
			v27 = int32(base.Ui32(v20) >> (uint(v26) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v25))) = v27<<(uint(int32(2))%32) + int32(12)
			v34 = v27 - v26
			if v34 == int32(0) {
				v42 = v25
			} else {
				base.MemoryCopy(m, v25+int32(4), v12+int32(1), v34)
				v42 = v25
			}
		}
		v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+4)))
		v45 = base.I32_extend16_s(v44)
		if base.Ui32(int32(_a_F_numeric_abbrev_convert_0)) <= base.Ui32(v44) {
			if v45 == int32(-4096) {
				v53 = int64(9223372036854775807)
			} else {
				v53 = int64(-9223372036854775807 - 1)
			}
			if v45 == int32(-12288) {
				v56 = int64(-9223372036854775807)
			} else {
				v56 = v53
			}
			v204 = v56
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
			v63 = base.B2i32(int32(0) <= v45)
			if int32(0) <= v45 {
				v64 = int32(-8)
			} else {
				v64 = int32(-6)
			}
			v67 = int32(base.Ui32(int32(base.Ui32(v57)>>(uint(int32(2))%32))+v64) >> (uint(int32(1)) % 32))
			if int32(0) <= v45 {
				v68 = int32(*(*int16)(unsafe.Add(mBase, uint32(v42)+6)))
				v78 = v68
			} else {
				v78 = v44<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v44&int32(63)
			}
			v84 = v44 & int32(_a_F_numeric_abbrev_convert_0)
			if v84 == int32(_a_F_numeric_abbrev_convert_1) {
				v87 = v44 << (uint(int32(1)) % 32) & int32(_a_F_numeric_abbrev_convert_2)
			} else {
				v87 = v84
			}
			if v67 == int32(0) {
				v129 = int64(0)
			} else {
				if v78 < int32(-44) {
					v129 = int64(0)
				} else {
					if int32(83) < v78 {
						v129 = int64(9223372036854775807)
					} else {
						if v45 < int32(0) {
							v101 = int32(6)
						} else {
							v101 = int32(8)
						}
						v102 = v42 + v101
						v107 = base.I64_extend_i32_u(v78+int32(44)) << (uint(int64(56)) % 64)
						switch v67 - int32(1) {
						case 0:
							v122 = v107
						case 1:
							v117 = v107
							v118 = int64(*(*int16)(unsafe.Add(mBase, uint32(v102)+2)))
							v122 = v118<<(uint(int64(28))%64) | v117
						case 2:
							v112 = v107
							v113 = int64(*(*int16)(unsafe.Add(mBase, uint32(v102)+4)))
							v117 = v113<<(uint(int64(14))%64) | v112
							v118 = int64(*(*int16)(unsafe.Add(mBase, uint32(v102)+2)))
							v122 = v118<<(uint(int64(28))%64) | v117
						default:
							v110 = int64(*(*int16)(unsafe.Add(mBase, uint32(v102)+6)))
							v112 = v107 | v110
							v113 = int64(*(*int16)(unsafe.Add(mBase, uint32(v102)+4)))
							v117 = v113<<(uint(int64(14))%64) | v112
							v118 = int64(*(*int16)(unsafe.Add(mBase, uint32(v102)+2)))
							v122 = v118<<(uint(int64(28))%64) | v117
						}
						v123 = int64(*(*int16)(unsafe.Add(mBase, uint32(v102))))
						v129 = v123<<(uint(int64(42))%64) | v122
					}
				}
			}
			if v87 != 0 {
				v132 = v129
			} else {
				v132 = int64(0) - v129
			}
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)))
			if v133 != int32(1) {
				v204 = v132
			} else {
				v136 = int32(24)
				v137 = v10 + v136
				v146 = int32(711645284)
				v149 = base.I32_wrap_i64(int64(base.Ui64(v132)>>(uint(int64(32))%64))^v132) - int32(1636608428) ^ v146 - int32(1455628627)
				v154 = v149 ^ int32(-1636608428) - base.I32_rotl(v149, int32(25))
				v159 = v154 ^ v146 - base.I32_rotl(v154, int32(16))
				v163 = v159 ^ v149 - base.I32_rotl(v159, int32(4))
				v167 = v163 ^ v154 - base.I32_rotl(v163, int32(14))
				v171 = v167 ^ v159 - base.I32_rotl(v167, v136)
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
				v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
				v177 = int32(32) - v176
				v179 = v174 + int32(base.Ui32(v171)>>(uint(v177)%32))
				v180 = v171 << (uint(v176) % 32)
				if v180 != 0 {
					v187 = int32(32) - (base.I32_clz(v180) ^ int32(31))
					v188 = int32(255)
					if base.Ui32(v177&v188) < base.Ui32(v187&v188) {
						v193 = v177 + int32(1)
					} else {
						v193 = v187
					}
					v197 = v193
				} else {
					v197 = v177 + int32(1)
				}
				v199 = v197 & int32(255)
				v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
				if base.Ui32(v200) < base.Ui32(v199) {
					v202 = v199
				} else {
					v202 = v200
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v202)
				v204 = v132
			}
		}
		if v12 != v11 {
			F_pfree(m, v12)
			mBase = m.M
			v211 = m.ExcPending
			if v211 != 0 {
				return int64(0)
			} else {
				return v204
			}
		} else {
			return v204
		}
	}
}
func F_numeric_avg_deserialize(m *base.Module, l0 int32) int64 {
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int64
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v120 int64
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
		switch v15 - int32(435) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v42 = F_pg_detoast_datum_packed(m, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
			v48 = int32(1)
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v52 = v50 & v48
			if v52 != 0 {
				v53 = v48
			} else {
				v53 = int32(4)
			}
			if v50 == int32(1) {
				v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
				if v60 == int32(18) {
					v63 = int32(16)
				} else {
					v63 = int32(0)
				}
				if base.Ui32((v60-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v70 = int32(4)
				} else {
					v70 = v63
				}
				v81 = v70
			} else {
				v71 = int32(1)
				if v52 != 0 {
					v81 = int32(base.Ui32(v50)>>(uint(v71)%32)) - v71
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v81
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v42 + v53
			v87 = F_palloc0(m, int32(112))
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int64(0)
			} else {
				v89 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v89)
				v92 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_avg_deserialize[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v92
				v95 = v8 + int32(32)
				v96 = F_pq_getmsgint64(m, v95)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v87)+8)) = v96
					v100 = v8 + int32(8)
					F_numericvar_deserialize(m, v95, v100)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						F_accum_sum_add(m, v87+int32(16), v100)
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int64(0)
						} else {
							v108 = F_pq_getmsgint(m, v95, int32(4))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v87)+72)) = v108
								v111 = F_pq_getmsgint64(m, v95)
								mBase = m.M
								v112 = m.ExcPending
								if v112 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v87)+80)) = v111
									v114 = F_pq_getmsgint64(m, v95)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v87)+88)) = v114
										v117 = F_pq_getmsgint64(m, v95)
										mBase = m.M
										v118 = m.ExcPending
										if v118 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v87)+96)) = v117
											v120 = F_pq_getmsgint64(m, v95)
											mBase = m.M
											v121 = m.ExcPending
											if v121 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v87)+104)) = v120
												F_pq_getmsgend(m, v95)
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int64(0)
												} else {
													v125 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
													if v125 != 0 {
														F_pfree(m, v125)
														mBase = m.M
														v127 = m.ExcPending
														if v127 != 0 {
															return int64(0)
														} else {
															m.G0 = v8 + int32(48)
															return base.I64_extend_i32_u(v87)
														}
													} else {
														m.G0 = v8 + int32(48)
														return base.I64_extend_i32_u(v87)
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
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_avg_deserialize_0), int32(0))
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_avg_deserialize_1), int32(_a_F_numeric_avg_deserialize_2), int32(_a_F_numeric_avg_deserialize_3))
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
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
func F_numeric_div_trunc(m *base.Module, l0 int32) int64 {
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
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = F_pg_detoast_datum(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
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
								v149 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_5), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int64(0)
								} else {
									v257 = v149
									m.G0 = v11 + int32(80)
									return base.I64_extend_i32_u(v257)
								}
							} else {
								if base.Ui32(int32(_a_F_numeric_div_trunc_2)) <= base.Ui32(v53&int32(_a_F_numeric_div_trunc_1)) {
									v75 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_6), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										v257 = v75
										m.G0 = v11 + int32(80)
										return base.I64_extend_i32_u(v257)
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
										v270 = m.ExcPending
										if v270 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(33816706))
											mBase = m.M
											v273 = m.ExcPending
											if v273 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_numeric_div_trunc_7), int32(0))
												mBase = m.M
												v277 = m.ExcPending
												if v277 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_numeric_div_trunc_8), int32(3284), int32(_a_F_numeric_div_trunc_9))
													mBase = m.M
													v282 = m.ExcPending
													if v282 != 0 {
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
										v94 = v53 & int32(_a_F_numeric_div_trunc_2)
										if v94 == int32(_a_F_numeric_div_trunc_10) {
											v97 = v53 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
										} else {
											v97 = v94
										}
										if v97 != int32(_a_F_numeric_div_trunc_11) {
											v102 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_12), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int64(0)
											} else {
												v257 = v102
												m.G0 = v11 + int32(80)
												return base.I64_extend_i32_u(v257)
											}
										} else {
											v106 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_13), int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int64(0)
											} else {
												v257 = v106
												m.G0 = v11 + int32(80)
												return base.I64_extend_i32_u(v257)
											}
										}
									}
								}
							}
						} else {
							if base.Ui32(int32(_a_F_numeric_div_trunc_2)) <= base.Ui32(v53&int32(_a_F_numeric_div_trunc_1)) {
								v114 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_6), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int64(0)
								} else {
									v257 = v114
									m.G0 = v11 + int32(80)
									return base.I64_extend_i32_u(v257)
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
									v286 = m.ExcPending
									if v286 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(33816706))
										mBase = m.M
										v289 = m.ExcPending
										if v289 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_numeric_div_trunc_7), int32(0))
											mBase = m.M
											v293 = m.ExcPending
											if v293 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_numeric_div_trunc_8), int32(3302), int32(_a_F_numeric_div_trunc_9))
												mBase = m.M
												v298 = m.ExcPending
												if v298 != 0 {
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
									v133 = v53 & int32(_a_F_numeric_div_trunc_2)
									if v133 == int32(_a_F_numeric_div_trunc_10) {
										v136 = v53 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
									} else {
										v136 = v133
									}
									if v136 != int32(_a_F_numeric_div_trunc_11) {
										v141 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_13), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int64(0)
										} else {
											v257 = v141
											m.G0 = v11 + int32(80)
											return base.I64_extend_i32_u(v257)
										}
									} else {
										v145 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_12), int32(0))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return int64(0)
										} else {
											v257 = v145
											m.G0 = v11 + int32(80)
											return base.I64_extend_i32_u(v257)
										}
									}
								}
							}
						}
					} else {
						v63 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_6), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int64(0)
						} else {
							v257 = v63
							m.G0 = v11 + int32(80)
							return base.I64_extend_i32_u(v257)
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
						return int64(0)
					} else {
						v250 = F_make_result_safe(m, v243, int32(0))
						mBase = m.M
						v251 = m.ExcPending
						if v251 != 0 {
							return int64(0)
						} else {
							v252 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
							if v252 == int32(0) {
								v257 = v250
								m.G0 = v11 + int32(80)
								return base.I64_extend_i32_u(v257)
							} else {
								F_pfree(m, v252)
								mBase = m.M
								v256 = m.ExcPending
								if v256 != 0 {
									return int64(0)
								} else {
									v257 = v250
									m.G0 = v11 + int32(80)
									return base.I64_extend_i32_u(v257)
								}
							}
						}
					}
				}
			} else {
				if v22 == int32(-16384) {
					v63 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_6), int32(0))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int64(0)
					} else {
						v257 = v63
						m.G0 = v11 + int32(80)
						return base.I64_extend_i32_u(v257)
					}
				} else {
					v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
					v53 = v52
					if v53&int32(_a_F_numeric_div_trunc_1) != int32(_a_F_numeric_div_trunc_2) {
						if v21 != int32(_a_F_numeric_div_trunc_3) {
							if v21 != int32(_a_F_numeric_div_trunc_4) {
								v149 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_5), int32(0))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int64(0)
								} else {
									v257 = v149
									m.G0 = v11 + int32(80)
									return base.I64_extend_i32_u(v257)
								}
							} else {
								if base.Ui32(int32(_a_F_numeric_div_trunc_2)) <= base.Ui32(v53&int32(_a_F_numeric_div_trunc_1)) {
									v75 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_6), int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int64(0)
									} else {
										v257 = v75
										m.G0 = v11 + int32(80)
										return base.I64_extend_i32_u(v257)
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
										v270 = m.ExcPending
										if v270 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(33816706))
											mBase = m.M
											v273 = m.ExcPending
											if v273 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_numeric_div_trunc_7), int32(0))
												mBase = m.M
												v277 = m.ExcPending
												if v277 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_numeric_div_trunc_8), int32(3284), int32(_a_F_numeric_div_trunc_9))
													mBase = m.M
													v282 = m.ExcPending
													if v282 != 0 {
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
										v94 = v53 & int32(_a_F_numeric_div_trunc_2)
										if v94 == int32(_a_F_numeric_div_trunc_10) {
											v97 = v53 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
										} else {
											v97 = v94
										}
										if v97 != int32(_a_F_numeric_div_trunc_11) {
											v102 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_12), int32(0))
											mBase = m.M
											v103 = m.ExcPending
											if v103 != 0 {
												return int64(0)
											} else {
												v257 = v102
												m.G0 = v11 + int32(80)
												return base.I64_extend_i32_u(v257)
											}
										} else {
											v106 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_13), int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int64(0)
											} else {
												v257 = v106
												m.G0 = v11 + int32(80)
												return base.I64_extend_i32_u(v257)
											}
										}
									}
								}
							}
						} else {
							if base.Ui32(int32(_a_F_numeric_div_trunc_2)) <= base.Ui32(v53&int32(_a_F_numeric_div_trunc_1)) {
								v114 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_6), int32(0))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int64(0)
								} else {
									v257 = v114
									m.G0 = v11 + int32(80)
									return base.I64_extend_i32_u(v257)
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
									v286 = m.ExcPending
									if v286 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(33816706))
										mBase = m.M
										v289 = m.ExcPending
										if v289 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_numeric_div_trunc_7), int32(0))
											mBase = m.M
											v293 = m.ExcPending
											if v293 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_numeric_div_trunc_8), int32(3302), int32(_a_F_numeric_div_trunc_9))
												mBase = m.M
												v298 = m.ExcPending
												if v298 != 0 {
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
									v133 = v53 & int32(_a_F_numeric_div_trunc_2)
									if v133 == int32(_a_F_numeric_div_trunc_10) {
										v136 = v53 << (uint(int32(1)) % 32) & int32(_a_F_numeric_div_trunc_11)
									} else {
										v136 = v133
									}
									if v136 != int32(_a_F_numeric_div_trunc_11) {
										v141 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_13), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int64(0)
										} else {
											v257 = v141
											m.G0 = v11 + int32(80)
											return base.I64_extend_i32_u(v257)
										}
									} else {
										v145 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_12), int32(0))
										mBase = m.M
										v146 = m.ExcPending
										if v146 != 0 {
											return int64(0)
										} else {
											v257 = v145
											m.G0 = v11 + int32(80)
											return base.I64_extend_i32_u(v257)
										}
									}
								}
							}
						}
					} else {
						v63 = F_make_result_safe(m, int32(_a_F_numeric_div_trunc_6), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int64(0)
						} else {
							v257 = v63
							m.G0 = v11 + int32(80)
							return base.I64_extend_i32_u(v257)
						}
					}
				}
			}
		}
	}
}
func F_numeric_fac(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v115 int32
	_ = v115
	var v117 int64
	_ = v117
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if int64(0) <= v15 {
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
	if base.Ui64(v15) <= base.Ui64(int64(1)) {
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
	return base.I64_extend_i32_u(v158)
L6:
	;
	v22 = F_make_result_safe(m, int32(_a_F_numeric_fac_0), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if base.Ui64(int64(32178)) <= base.Ui64(v15) {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	return int64(0)
L10:
	;
	v158 = v22
	goto L5
L11:
	;
	v28 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+40)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v28
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v28
	v35 = F_palloc(m, int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v35
	v38 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v38)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = int64(0)
	v45 = v38
	v48 = v35 + int32(12)
	v51 = v15
	goto L13
L13:
	;
	v56 = v48 - int32(2)
	v58 = base.I64_div_u_s(v51, int64(10000))
	v61 = v58*int64(55536) + v51
	*(*uint16)(unsafe.Add(mBase, uint32(v56))) = uint16(v61)
	v64 = v45 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v51) {
		v45 = v64
		v48 = v56
		v51 = v58
		goto L13
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v64
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v56
	if v15 == int64(2) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L14
L16:
	;
	v151 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v151
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v153 == v151 {
		v158 = v141
		goto L5
	} else {
		goto L39
	}
L17:
	;
	v73 = F_make_result_safe(m, v13, int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
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
	v82 = v15
	goto L21
L20:
	;
	v141 = v73
	goto L16
L21:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_fac[0]))
	if v86 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v137 = F_make_result_safe(m, v13, int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L37
	}
L23:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
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
	v90 = m.ExcPending
	if v90 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v92 = F_palloc(m, int32(12))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L9
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v92
	v95 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v92))) = uint16(v95)
	*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = int64(0)
	v103 = v82 - int64(1)
	v104 = v95
	v107 = v92 + int32(12)
	v110 = v103
	goto L32
L32:
	;
	v115 = v107 - int32(2)
	v117 = base.I64_div_u_s(v110, int64(10000))
	v120 = v117*int64(55536) + v110
	*(*uint16)(unsafe.Add(mBase, uint32(v115))) = uint16(v120)
	v123 = v104 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v110) {
		v104 = v123
		v107 = v115
		v110 = v117
		goto L32
	} else {
		goto L34
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v115
	F_mul_var(m, v13, v13+int32(24), v13, int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	if base.Ui64(int64(3)) < base.Ui64(v82) {
		v79 = v92
		v82 = v103
		goto L21
	} else {
		goto L36
	}
L36:
	;
	goto L22
L37:
	;
	F_pfree(m, v92)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L9
	} else {
		goto L38
	}
L38:
	;
	v141 = v137
	goto L16
L39:
	;
	F_pfree(m, v153)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	v158 = v141
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
	F_errfinish(m, int32(_a_F_numeric_fac_2), int32(3621), int32(_a_F_numeric_fac_3))
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
	F_errfinish(m, int32(_a_F_numeric_fac_2), int32(3631), int32(_a_F_numeric_fac_3))
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
func F_numeric_float8(m *base.Module, l0 int32) int64 {
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
	var v23 int64
	_ = v23
	var v26 int64
	_ = v26
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int64
	_ = v47
	var v49 int64
	_ = v49
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		if base.Ui32(int32(_a_F_numeric_float8_0)) <= base.Ui32(v15) {
			if v15 == int32(_a_F_numeric_float8_1) {
				v23 = int64(-4503599627370496)
			} else {
				v23 = int64(9221120237041090560)
			}
			if v15 == int32(_a_F_numeric_float8_2) {
				v26 = int64(9218868437227405312)
			} else {
				v26 = v23
			}
			v49 = v26
			m.G0 = v8 + int32(16)
			return v49
		} else {
			v31 = F_DirectFunctionCall1Coll(m, int32(664), int32(0), base.I64_extend_i32_u(v11))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				v33 = base.I32_wrap_i64(v31)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v38 = F_DirectInputFunctionCallSafe(m, int32(1600), v33, int32(-1), v35, v8+int32(8))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int64(0)
				} else {
					if v38 == int32(0) {
						F_pfree(m, v33)
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int64(0)
						} else {
							v44 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v44)
							v49 = int64(0)
							m.G0 = v8 + int32(16)
							return v49
						}
					} else {
						v47 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
						v49 = v47
						m.G0 = v8 + int32(16)
						return v49
					}
				}
			}
		}
	}
}
func F_numeric_ln(m *base.Module, l0 int32) int64 {
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
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		v16 = base.I32_extend16_s(v15)
		if base.Ui32(int32(_a_F_numeric_ln_0)) <= base.Ui32(v15) {
			if v16 == int32(-4096) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(352583810))
					mBase = m.M
					v126 = m.ExcPending
					if v126 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_numeric_ln_1), int32(0))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_ln_2), int32(3819), int32(_a_F_numeric_ln_3))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
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
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v24 = F_palloc(m, int32(base.Ui32(v21)>>(uint(int32(2))%32)))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
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
					return base.I64_extend_i32_u(v111)
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
				return int64(0)
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
					return int64(0)
				} else {
					v104 = F_make_result_safe(m, v8, int32(0))
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int64(0)
					} else {
						v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
						if v106 == int32(0) {
							v111 = v104
							m.G0 = v8 + int32(48)
							return base.I64_extend_i32_u(v111)
						} else {
							F_pfree(m, v106)
							mBase = m.M
							v110 = m.ExcPending
							if v110 != 0 {
								return int64(0)
							} else {
								v111 = v104
								m.G0 = v8 + int32(48)
								return base.I64_extend_i32_u(v111)
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_lt(m *base.Module, l0 int32) int64 {
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
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v25 = base.I32_extend16_s(v24)
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			if base.Ui32(int32(_a_F_numeric_lt_0)) <= base.Ui32(v26) {
				if v26 != int32(_a_F_numeric_lt_1) {
					if v26 != int32(_a_F_numeric_lt_0) {
						if v25 != int32(-4096) {
							v45 = int32(-1)
						} else {
							v45 = int32(0)
						}
						v166 = v45
					} else {
						v166 = base.B2i32(v25 != int32(-16384))
					}
				} else {
					if v25 == int32(-16384) {
						v40 = int32(-1)
					} else {
						v40 = base.B2i32(v25 != int32(-12288))
					}
					v166 = v40
				}
			} else {
				if base.Ui32(int32(-16384)) <= base.Ui32(v25) {
					if v25 == int32(-4096) {
						v52 = int32(1)
					} else {
						v52 = int32(-1)
					}
					v166 = v52
				} else {
					v54 = v6 + int32(6)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v60 = base.I32_extend16_s(v26)
					v62 = base.B2i32(int32(0) <= v60)
					if int32(0) <= v60 {
						v63 = int32(-8)
					} else {
						v63 = int32(-6)
					}
					if int32(0) <= v60 {
						v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54))))
						v75 = v65
					} else {
						v75 = v26<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v26&int32(63)
					}
					v77 = int32(base.Ui32(int32(base.Ui32(v55)>>(uint(int32(2))%32))+v63) >> (uint(int32(1)) % 32))
					v79 = v11 + int32(6)
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v86 = base.B2i32(int32(0) <= v25)
					if int32(0) <= v25 {
						v87 = int32(-8)
					} else {
						v87 = int32(-6)
					}
					if int32(0) <= v25 {
						v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79))))
						v99 = v89
					} else {
						v99 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v100 = int32(1)
					v101 = int32(base.Ui32(int32(base.Ui32(v80)>>(uint(int32(2))%32))+v87) >> (uint(v100) % 32))
					v107 = v24 & int32(_a_F_numeric_lt_0)
					if v107 == int32(_a_F_numeric_lt_2) {
						v110 = v24 << (uint(v100) % 32) & int32(_a_F_numeric_lt_3)
					} else {
						v110 = v107
					}
					if v77 == int32(0) {
						if v101 == int32(0) {
							v166 = int32(0)
						} else {
							if v110 == int32(_a_F_numeric_lt_3) {
								v120 = int32(1)
							} else {
								v120 = int32(-1)
							}
							v166 = v120
						}
					} else {
						v126 = v26 & int32(_a_F_numeric_lt_0)
						if v126 == int32(_a_F_numeric_lt_2) {
							v129 = v26 << (uint(int32(1)) % 32) & int32(_a_F_numeric_lt_3)
						} else {
							v129 = v126
						}
						if v101 == int32(0) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v166 = v134
						} else {
							if v60 < int32(0) {
								v139 = v54
							} else {
								v139 = v6 + int32(8)
							}
							if v25 < int32(0) {
								v144 = v79
							} else {
								v144 = v11 + int32(8)
							}
							if v129 == int32(0) {
								if v110 == int32(_a_F_numeric_lt_3) {
									v166 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v139, v77, v75, v144, v101, v99)
									mBase = m.M
									v166 = v150
								}
							} else {
								if v110 == int32(0) {
									v166 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v144, v101, v99, v139, v77, v75)
									mBase = m.M
									v166 = v154
								}
							}
						}
					}
				}
			}
			v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v167 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v170 = m.ExcPending
				if v170 != 0 {
					return int64(0)
				} else {
					v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v171 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(int32(base.Ui32(v166) >> (uint(int32(31)) % 32)))
						}
					} else {
						return base.I64_extend_i32_u(int32(base.Ui32(v166) >> (uint(int32(31)) % 32)))
					}
				}
			} else {
				v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v171 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(int32(base.Ui32(v166) >> (uint(int32(31)) % 32)))
					}
				} else {
					return base.I64_extend_i32_u(int32(base.Ui32(v166) >> (uint(int32(31)) % 32)))
				}
			}
		}
	}
}
func F_numeric_mod_safe(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v149 int64
	_ = v149
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int64
	_ = v187
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v15 = base.I32_extend16_s(v14)
	if base.Ui32(v14) <= base.Ui32(int32(_a_F_numeric_mod_safe_0)) {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v19 = base.I32_extend16_s(v18)
		if base.Ui32(int32(_a_F_numeric_mod_safe_0)) < base.Ui32(v18) {
			v47 = v19
			if v47&int32(_a_F_numeric_mod_safe_1) != int32(_a_F_numeric_mod_safe_2) {
				if v15&int32(-8193) == int32(-12288) {
					if base.Ui32(v47&int32(_a_F_numeric_mod_safe_1)) <= base.Ui32(int32(_a_F_numeric_mod_safe_0)) {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v47) {
							v76 = int32(-8)
						} else {
							v76 = int32(-6)
						}
						if base.Ui32(int32(base.Ui32(v68)>>(uint(int32(2))%32))+v76) < base.Ui32(int32(2)) {
							v224 = int32(0)
							v225 = F_errsave_start(m, l2)
							mBase = m.M
							v226 = m.ExcPending
							if v226 != 0 {
								return int32(0)
							} else {
								if v225 == int32(0) {
									v242 = v224
									m.G0 = v12 + int32(96)
									return v242
								} else {
									F_errcode(m, int32(33816706))
									mBase = m.M
									v231 = m.ExcPending
									if v231 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_numeric_mod_safe_3), int32(0))
										mBase = m.M
										v235 = m.ExcPending
										if v235 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, l2, int32(_a_F_numeric_mod_safe_4), int32(3414), int32(_a_F_numeric_mod_safe_5))
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												v242 = v224
												m.G0 = v12 + int32(96)
												return v242
											}
										}
									}
								}
							}
						} else {
							v82 = F_make_result_safe(m, int32(_a_F_numeric_mod_safe_6), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								v242 = v82
								m.G0 = v12 + int32(96)
								return v242
							}
						}
					} else {
						v82 = F_make_result_safe(m, int32(_a_F_numeric_mod_safe_6), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v242 = v82
							m.G0 = v12 + int32(96)
							return v242
						}
					}
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v87 = F_palloc(m, int32(base.Ui32(v84)>>(uint(int32(2))%32)))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v91 = int32(base.Ui32(v89) >> (uint(int32(2)) % 32))
						if v91 == int32(0) {
							v242 = v87
						} else {
							base.MemoryCopy(m, v87, l0, v91)
							v242 = v87
						}
						m.G0 = v12 + int32(96)
						return v242
					}
				}
			} else {
				v56 = F_make_result_safe(m, int32(_a_F_numeric_mod_safe_6), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v242 = v56
					m.G0 = v12 + int32(96)
					return v242
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
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(base.Ui32(int32(base.Ui32(v22)>>(uint(int32(2))%32))+v29) >> (uint(int32(1)) % 32))
			if int32(0) <= v15 {
				v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v96 = v95
			} else {
				v96 = v14<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v14&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v96
			v98 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v98
			v107 = base.B2i32(v15 < v98)
			if v15 < v98 {
				v108 = int32(base.Ui32(v14)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v108 = v14 & int32(_a_F_numeric_mod_safe_7)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v108
			v115 = v14 & int32(_a_F_numeric_mod_safe_2)
			if v115 == int32(_a_F_numeric_mod_safe_8) {
				v118 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mod_safe_9)
			} else {
				v118 = v115
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v118
			if v15 < v98 {
				v122 = int32(6)
			} else {
				v122 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = l0 + v122
			v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v131 = base.B2i32(int32(0) <= v19)
			if int32(0) <= v19 {
				v132 = int32(-8)
			} else {
				v132 = int32(-6)
			}
			v135 = int32(base.Ui32(int32(base.Ui32(v125)>>(uint(int32(2))%32))+v132) >> (uint(int32(1)) % 32))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v135
			if int32(0) <= v19 {
				v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				v147 = v137
			} else {
				v147 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v147
			v149 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v149
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v149
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v149
			v155 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v155
			v164 = base.B2i32(v19 < v155)
			if v19 < v155 {
				v165 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v165 = v18 & int32(_a_F_numeric_mod_safe_7)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v165
			v172 = v18 & int32(_a_F_numeric_mod_safe_2)
			if v172 == int32(_a_F_numeric_mod_safe_8) {
				v175 = v18 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mod_safe_9)
			} else {
				v175 = v172
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v175
			if v19 < v155 {
				v179 = int32(6)
			} else {
				v179 = int32(8)
			}
			v180 = l1 + v179
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v180
			if v135 == int32(0) {
				v224 = int32(0)
				v225 = F_errsave_start(m, l2)
				mBase = m.M
				v226 = m.ExcPending
				if v226 != 0 {
					return int32(0)
				} else {
					if v225 == int32(0) {
						v242 = v224
						m.G0 = v12 + int32(96)
						return v242
					} else {
						F_errcode(m, int32(33816706))
						mBase = m.M
						v231 = m.ExcPending
						if v231 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_numeric_mod_safe_3), int32(0))
							mBase = m.M
							v235 = m.ExcPending
							if v235 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, l2, int32(_a_F_numeric_mod_safe_4), int32(3414), int32(_a_F_numeric_mod_safe_5))
								mBase = m.M
								v240 = m.ExcPending
								if v240 != 0 {
									return int32(0)
								} else {
									v242 = v224
									m.G0 = v12 + int32(96)
									return v242
								}
							}
						}
					}
				}
			} else {
				v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180))))
				if v184 == int32(0) {
					v224 = int32(0)
					v225 = F_errsave_start(m, l2)
					mBase = m.M
					v226 = m.ExcPending
					if v226 != 0 {
						return int32(0)
					} else {
						if v225 == int32(0) {
							v242 = v224
							m.G0 = v12 + int32(96)
							return v242
						} else {
							F_errcode(m, int32(33816706))
							mBase = m.M
							v231 = m.ExcPending
							if v231 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_numeric_mod_safe_3), int32(0))
								mBase = m.M
								v235 = m.ExcPending
								if v235 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, l2, int32(_a_F_numeric_mod_safe_4), int32(3414), int32(_a_F_numeric_mod_safe_5))
									mBase = m.M
									v240 = m.ExcPending
									if v240 != 0 {
										return int32(0)
									} else {
										v242 = v224
										m.G0 = v12 + int32(96)
										return v242
									}
								}
							}
						}
					}
				} else {
					v187 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v187
					*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v187
					*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v187
					v194 = v12 + int32(48)
					v196 = v12 + int32(24)
					v198 = v12 + int32(72)
					v199 = int32(0)
					F_div_var(m, v194, v196, v198, v199, v199, int32(1))
					mBase = m.M
					v203 = m.ExcPending
					if v203 != 0 {
						return int32(0)
					} else {
						F_mul_var(m, v196, v198, v198, v165)
						mBase = m.M
						v205 = m.ExcPending
						if v205 != 0 {
							return int32(0)
						} else {
							F_sub_var(m, v194, v198, v12)
							mBase = m.M
							v207 = m.ExcPending
							if v207 != 0 {
								return int32(0)
							} else {
								v208 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
								if v208 != 0 {
									F_pfree(m, v208)
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
										return int32(0)
									} else {
										v211 = F_make_result_safe(m, v12, l2)
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
											return int32(0)
										} else {
											v213 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
											if v213 == int32(0) {
												v242 = v211
												m.G0 = v12 + int32(96)
												return v242
											} else {
												F_pfree(m, v213)
												mBase = m.M
												v217 = m.ExcPending
												if v217 != 0 {
													return int32(0)
												} else {
													v242 = v211
													m.G0 = v12 + int32(96)
													return v242
												}
											}
										}
									}
								} else {
									v211 = F_make_result_safe(m, v12, l2)
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return int32(0)
									} else {
										v213 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
										if v213 == int32(0) {
											v242 = v211
											m.G0 = v12 + int32(96)
											return v242
										} else {
											F_pfree(m, v213)
											mBase = m.M
											v217 = m.ExcPending
											if v217 != 0 {
												return int32(0)
											} else {
												v242 = v211
												m.G0 = v12 + int32(96)
												return v242
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
		if v15 == int32(-16384) {
			v56 = F_make_result_safe(m, int32(_a_F_numeric_mod_safe_6), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v242 = v56
				m.G0 = v12 + int32(96)
				return v242
			}
		} else {
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v47 = v45
			if v47&int32(_a_F_numeric_mod_safe_1) != int32(_a_F_numeric_mod_safe_2) {
				if v15&int32(-8193) == int32(-12288) {
					if base.Ui32(v47&int32(_a_F_numeric_mod_safe_1)) <= base.Ui32(int32(_a_F_numeric_mod_safe_0)) {
						v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v47) {
							v76 = int32(-8)
						} else {
							v76 = int32(-6)
						}
						if base.Ui32(int32(base.Ui32(v68)>>(uint(int32(2))%32))+v76) < base.Ui32(int32(2)) {
							v224 = int32(0)
							v225 = F_errsave_start(m, l2)
							mBase = m.M
							v226 = m.ExcPending
							if v226 != 0 {
								return int32(0)
							} else {
								if v225 == int32(0) {
									v242 = v224
									m.G0 = v12 + int32(96)
									return v242
								} else {
									F_errcode(m, int32(33816706))
									mBase = m.M
									v231 = m.ExcPending
									if v231 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_numeric_mod_safe_3), int32(0))
										mBase = m.M
										v235 = m.ExcPending
										if v235 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, l2, int32(_a_F_numeric_mod_safe_4), int32(3414), int32(_a_F_numeric_mod_safe_5))
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												v242 = v224
												m.G0 = v12 + int32(96)
												return v242
											}
										}
									}
								}
							}
						} else {
							v82 = F_make_result_safe(m, int32(_a_F_numeric_mod_safe_6), int32(0))
							mBase = m.M
							v83 = m.ExcPending
							if v83 != 0 {
								return int32(0)
							} else {
								v242 = v82
								m.G0 = v12 + int32(96)
								return v242
							}
						}
					} else {
						v82 = F_make_result_safe(m, int32(_a_F_numeric_mod_safe_6), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v242 = v82
							m.G0 = v12 + int32(96)
							return v242
						}
					}
				} else {
					v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v87 = F_palloc(m, int32(base.Ui32(v84)>>(uint(int32(2))%32)))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v91 = int32(base.Ui32(v89) >> (uint(int32(2)) % 32))
						if v91 == int32(0) {
							v242 = v87
						} else {
							base.MemoryCopy(m, v87, l0, v91)
							v242 = v87
						}
						m.G0 = v12 + int32(96)
						return v242
					}
				}
			} else {
				v56 = F_make_result_safe(m, int32(_a_F_numeric_mod_safe_6), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v242 = v56
					m.G0 = v12 + int32(96)
					return v242
				}
			}
		}
	}
}
func F_numeric_poly_combine(m *base.Module, l0 int32) int64 {
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v111 int64
	_ = v111
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v116 int64
	_ = v116
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v145 int64
	_ = v145
	var v149 int32
	_ = v149
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
		switch v19 - int32(435) {
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
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v16)+376))
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
		v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v45 == int32(0) {
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v49 = v48
		} else {
			v49 = v2
		}
		v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v50 == int32(0) {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v53 != 0 {
				if v49 == int32(0) {
					v60 = int32(_a_F_numeric_poly_combine_0)
					v61 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_poly_combine[0]))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
					*(*int32)(unsafe.Add(mBase, _c_F_numeric_poly_combine[0])) = v63
					v66 = v11 + int32(12)
					v67 = int32(0)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v68 == v67 {
						v85 = int32(0)
						if v66 == v85 {
							v93 = v85
						} else {
							v88 = v85
							v89 = v67
							*(*int32)(unsafe.Add(mBase, uint32(v66))) = v88
							v93 = v89
						}
						v96 = v93
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
						switch v71 - int32(435) {
						case 0:
							if v66 == int32(0) {
								v96 = int32(1)
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v68)+168))
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
								v88 = v78
								v89 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v66))) = v88
								v93 = v89
								v96 = v93
							}
						case 1:
							if v66 == int32(0) {
								v96 = int32(2)
							} else {
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v68)+376))
								v88 = v83
								v89 = int32(2)
								*(*int32)(unsafe.Add(mBase, uint32(v66))) = v88
								v93 = v89
								v96 = v93
							}
						default:
							v85 = int32(0)
							if v66 == v85 {
								v93 = v85
							} else {
								v88 = v85
								v89 = v67
								*(*int32)(unsafe.Add(mBase, uint32(v66))) = v88
								v93 = v89
							}
							v96 = v93
						}
					}
					if v96 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return int64(0)
						} else {
							F_errmsg_internal(m, int32(_a_F_numeric_poly_combine_1), int32(0))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_numeric_poly_combine_2), int32(_a_F_numeric_poly_combine_3), int32(_a_F_numeric_poly_combine_4))
								mBase = m.M
								v185 = m.ExcPending
								if v185 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
						*(*int32)(unsafe.Add(mBase, _c_F_numeric_poly_combine[0])) = v100
						v103 = F_palloc0(m, int32(48))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int64(0)
						} else {
							v107 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v107)
							v109 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v103)+8)) = v109
							v111 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
							v112 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v103)+24)) = v112
							*(*int64)(unsafe.Add(mBase, uint32(v103)+16)) = v111
							v115 = *(*int64)(unsafe.Add(mBase, uint32(v53)+32))
							v116 = *(*int64)(unsafe.Add(mBase, uint32(v53)+40))
							*(*int64)(unsafe.Add(mBase, uint32(v103)+40)) = v116
							*(*int64)(unsafe.Add(mBase, uint32(v103)+32)) = v115
							*(*int32)(unsafe.Add(mBase, _c_F_numeric_poly_combine[0])) = v61
							v149 = v103
							m.G0 = v11 + int32(16)
							return base.I64_extend_i32_u(v149)
						}
					}
				} else {
					v121 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
					if v121 <= int64(0) {
						v149 = v49
					} else {
						v124 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v124 + v121
						v127 = *(*int64)(unsafe.Add(mBase, uint32(v53)+24))
						v128 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
						v129 = *(*int64)(unsafe.Add(mBase, uint32(v53)+16))
						v130 = v128 + v129
						*(*int64)(unsafe.Add(mBase, uint32(v49)+16)) = v130
						v134 = *(*int64)(unsafe.Add(mBase, uint32(v49)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+24)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v130) < base.Ui64(v128))) + (v127 + v134)
						v138 = *(*int64)(unsafe.Add(mBase, uint32(v53)+40))
						v139 = *(*int64)(unsafe.Add(mBase, uint32(v49)+32))
						v140 = *(*int64)(unsafe.Add(mBase, uint32(v53)+32))
						v141 = v139 + v140
						*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v141
						v145 = *(*int64)(unsafe.Add(mBase, uint32(v49)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v49)+40)) = base.I64_extend_i32_u(base.B2i32(base.Ui64(v141) < base.Ui64(v139))) + (v138 + v145)
						v149 = v49
					}
					m.G0 = v11 + int32(16)
					return base.I64_extend_i32_u(v149)
				}
			} else {
				if v49 != 0 {
					v149 = v49
				} else {
					v55 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
					v149 = int32(0)
				}
				m.G0 = v11 + int32(16)
				return base.I64_extend_i32_u(v149)
			}
		} else {
			if v49 != 0 {
				v149 = v49
			} else {
				v55 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
				v149 = int32(0)
			}
			m.G0 = v11 + int32(16)
			return base.I64_extend_i32_u(v149)
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v163 = m.ExcPending
		if v163 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_poly_combine_1), int32(0))
			mBase = m.M
			v167 = m.ExcPending
			if v167 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_poly_combine_2), int32(_a_F_numeric_poly_combine_5), int32(_a_F_numeric_poly_combine_6))
				mBase = m.M
				v172 = m.ExcPending
				if v172 != 0 {
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
func F_numeric_poly_deserialize(m *base.Module, l0 int32) int64 {
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == int32(0) {
		v41 = int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		switch v16 - int32(435) {
		case 0:
			v41 = int32(1)
		case 1:
			v41 = int32(2)
		default:
			v41 = int32(0)
		}
	}
	if v41 != 0 {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v43 = F_pg_detoast_datum_packed(m, v42)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int64(0)
		} else {
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v76
			v80 = int32(1)
			if v47&v80 != 0 {
				v84 = v80
			} else {
				v84 = int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v43 + v84
			v88 = F_palloc0(m, int32(48))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int64(0)
			} else {
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
				v92 = F_pq_getmsgint64(m, v9)
				mBase = m.M
				v93 = m.ExcPending
				if v93 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v92
					v95 = F_pq_getmsgint64(m, v9)
					mBase = m.M
					v96 = m.ExcPending
					if v96 != 0 {
						return int64(0)
					} else {
						v97 = F_pq_getmsgint64(m, v9)
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v88)+24)) = v95
							*(*int64)(unsafe.Add(mBase, uint32(v88)+16)) = v97
							v101 = F_pq_getmsgint64(m, v9)
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return int64(0)
							} else {
								v103 = F_pq_getmsgint64(m, v9)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v88)+40)) = v101
									*(*int64)(unsafe.Add(mBase, uint32(v88)+32)) = v103
									F_pq_getmsgend(m, v9)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int64(0)
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_u(v88)
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
		v117 = m.ExcPending
		if v117 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_poly_deserialize_0), int32(0))
			mBase = m.M
			v121 = m.ExcPending
			if v121 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_poly_deserialize_1), int32(_a_F_numeric_poly_deserialize_2), int32(_a_F_numeric_poly_deserialize_3))
				mBase = m.M
				v126 = m.ExcPending
				if v126 != 0 {
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
func F_numeric_send(m *base.Module, l0 int32) int64 {
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
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	return base.I64_extend_i32_u(v169)
}
func F_numeric_stddev_pop(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = Fn14336(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_numeric_sub(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v11 = F_numeric_sub_safe(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v11)
			}
		}
	}
}
