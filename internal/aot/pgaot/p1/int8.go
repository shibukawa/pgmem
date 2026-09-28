package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_int8_accum(m *base.Module, l0 int32) int64 {
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v93 int64
	_ = v93
	var v97 int64
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	v2 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v14 == v2 {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v17 != 0 {
			v70 = v17
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v71 == int32(0) {
				v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
				v78 = F_palloc(m, int32(12))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v78
					v81 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v78))) = uint16(v81)
					*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v78 + int32(2)
					if v74 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
						v97 = int64(0) - v74
						v100 = v81
						v104 = v78 + int32(12)
						v107 = v97
						for {
							v110 = v104 - int32(2)
							v112 = base.I64_div_u_s(v107, int64(10000))
							v115 = v112*int64(55536) + v107
							*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v115)
							v118 = v100 + int32(1)
							if base.Ui64(int64(9999)) < base.Ui64(v107) {
								v100 = v118
								v104 = v110
								v107 = v112
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
						v122 = v118
						v127 = v100
					} else {
						v93 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v93
						if v74 == v93 {
							v122 = v81
							v127 = v2
						} else {
							v97 = v74
							v100 = v81
							v104 = v78 + int32(12)
							v107 = v97
							for {
								v110 = v104 - int32(2)
								v112 = base.I64_div_u_s(v107, int64(10000))
								v115 = v112*int64(55536) + v107
								*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v115)
								v118 = v100 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v107) {
									v100 = v118
									v104 = v110
									v107 = v112
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
							v122 = v118
							v127 = v100
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v127
					*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v122
					v136 = F_make_result_safe(m, v12+int32(8), int32(0))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return int64(0)
					} else {
						F_pfree(m, v78)
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int64(0)
						} else {
							F_do_numeric_accum(m, v70, v136)
							mBase = m.M
							v141 = m.ExcPending
							if v141 != 0 {
								return int64(0)
							} else {
								m.G0 = v12 + int32(32)
								return base.I64_extend_i32_u(v70)
							}
						}
					}
				}
			} else {
				m.G0 = v12 + int32(32)
				return base.I64_extend_i32_u(v70)
			}
		} else {
			v20 = v12 + int32(8)
			v21 = int32(0)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v22 == v21 {
				v39 = int32(0)
				if v20 == v39 {
					v47 = v39
				} else {
					v42 = v39
					v43 = v21
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
					v47 = v43
				}
				v50 = v47
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				switch v25 - int32(435) {
				case 0:
					if v20 == int32(0) {
						v50 = int32(1)
					} else {
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+168))
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
						v42 = v32
						v43 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
						v47 = v43
						v50 = v47
					}
				case 1:
					if v20 == int32(0) {
						v50 = int32(2)
					} else {
						v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+376))
						v42 = v37
						v43 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
						v47 = v43
						v50 = v47
					}
				default:
					v39 = int32(0)
					if v20 == v39 {
						v47 = v39
					} else {
						v42 = v39
						v43 = v21
						*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
						v47 = v43
					}
					v50 = v47
				}
			}
			if v50 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v159 = m.ExcPending
				if v159 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_int8_accum_0), int32(0))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_int8_accum_1), int32(_a_F_int8_accum_2), int32(_a_F_int8_accum_3))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v53 = int32(_a_F_int8_accum_4)
				v54 = *(*int32)(unsafe.Add(mBase, _c_F_int8_accum[0]))
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int32)(unsafe.Add(mBase, _c_F_int8_accum[0])) = v56
				v59 = F_palloc0(m, int32(112))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int64(0)
				} else {
					v63 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v65
					*(*int32)(unsafe.Add(mBase, _c_F_int8_accum[0])) = v54
					v70 = v59
					v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
					if v71 == int32(0) {
						v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
						*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
						v78 = F_palloc(m, int32(12))
						mBase = m.M
						v79 = m.ExcPending
						if v79 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v78
							v81 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(v78))) = uint16(v81)
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v78 + int32(2)
							if v74 < int64(0) {
								*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
								v97 = int64(0) - v74
								v100 = v81
								v104 = v78 + int32(12)
								v107 = v97
								for {
									v110 = v104 - int32(2)
									v112 = base.I64_div_u_s(v107, int64(10000))
									v115 = v112*int64(55536) + v107
									*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v115)
									v118 = v100 + int32(1)
									if base.Ui64(int64(9999)) < base.Ui64(v107) {
										v100 = v118
										v104 = v110
										v107 = v112
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
								v122 = v118
								v127 = v100
							} else {
								v93 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v93
								if v74 == v93 {
									v122 = v81
									v127 = v2
								} else {
									v97 = v74
									v100 = v81
									v104 = v78 + int32(12)
									v107 = v97
									for {
										v110 = v104 - int32(2)
										v112 = base.I64_div_u_s(v107, int64(10000))
										v115 = v112*int64(55536) + v107
										*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v115)
										v118 = v100 + int32(1)
										if base.Ui64(int64(9999)) < base.Ui64(v107) {
											v100 = v118
											v104 = v110
											v107 = v112
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
									v122 = v118
									v127 = v100
								}
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v127
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v122
							v136 = F_make_result_safe(m, v12+int32(8), int32(0))
							mBase = m.M
							v137 = m.ExcPending
							if v137 != 0 {
								return int64(0)
							} else {
								F_pfree(m, v78)
								mBase = m.M
								v139 = m.ExcPending
								if v139 != 0 {
									return int64(0)
								} else {
									F_do_numeric_accum(m, v70, v136)
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
										return int64(0)
									} else {
										m.G0 = v12 + int32(32)
										return base.I64_extend_i32_u(v70)
									}
								}
							}
						}
					} else {
						m.G0 = v12 + int32(32)
						return base.I64_extend_i32_u(v70)
					}
				}
			}
		}
	} else {
		v20 = v12 + int32(8)
		v21 = int32(0)
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v22 == v21 {
			v39 = int32(0)
			if v20 == v39 {
				v47 = v39
			} else {
				v42 = v39
				v43 = v21
				*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
				v47 = v43
			}
			v50 = v47
		} else {
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			switch v25 - int32(435) {
			case 0:
				if v20 == int32(0) {
					v50 = int32(1)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+168))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
					v42 = v32
					v43 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
					v47 = v43
					v50 = v47
				}
			case 1:
				if v20 == int32(0) {
					v50 = int32(2)
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v22)+376))
					v42 = v37
					v43 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
					v47 = v43
					v50 = v47
				}
			default:
				v39 = int32(0)
				if v20 == v39 {
					v47 = v39
				} else {
					v42 = v39
					v43 = v21
					*(*int32)(unsafe.Add(mBase, uint32(v20))) = v42
					v47 = v43
				}
				v50 = v47
			}
		}
		if v50 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v159 = m.ExcPending
			if v159 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_int8_accum_0), int32(0))
				mBase = m.M
				v163 = m.ExcPending
				if v163 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_int8_accum_1), int32(_a_F_int8_accum_2), int32(_a_F_int8_accum_3))
					mBase = m.M
					v168 = m.ExcPending
					if v168 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v53 = int32(_a_F_int8_accum_4)
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_int8_accum[0]))
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			*(*int32)(unsafe.Add(mBase, _c_F_int8_accum[0])) = v56
			v59 = F_palloc0(m, int32(112))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int64(0)
			} else {
				v63 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v59)+4)) = v65
				*(*int32)(unsafe.Add(mBase, _c_F_int8_accum[0])) = v54
				v70 = v59
				v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
				if v71 == int32(0) {
					v74 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(0)
					v78 = F_palloc(m, int32(12))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v78
						v81 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v78))) = uint16(v81)
						*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v78 + int32(2)
						if v74 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = int64(16384)
							v97 = int64(0) - v74
							v100 = v81
							v104 = v78 + int32(12)
							v107 = v97
							for {
								v110 = v104 - int32(2)
								v112 = base.I64_div_u_s(v107, int64(10000))
								v115 = v112*int64(55536) + v107
								*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v115)
								v118 = v100 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v107) {
									v100 = v118
									v104 = v110
									v107 = v112
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
							v122 = v118
							v127 = v100
						} else {
							v93 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v93
							if v74 == v93 {
								v122 = v81
								v127 = v2
							} else {
								v97 = v74
								v100 = v81
								v104 = v78 + int32(12)
								v107 = v97
								for {
									v110 = v104 - int32(2)
									v112 = base.I64_div_u_s(v107, int64(10000))
									v115 = v112*int64(55536) + v107
									*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v115)
									v118 = v100 + int32(1)
									if base.Ui64(int64(9999)) < base.Ui64(v107) {
										v100 = v118
										v104 = v110
										v107 = v112
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v110
								v122 = v118
								v127 = v100
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v127
						*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v122
						v136 = F_make_result_safe(m, v12+int32(8), int32(0))
						mBase = m.M
						v137 = m.ExcPending
						if v137 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v78)
							mBase = m.M
							v139 = m.ExcPending
							if v139 != 0 {
								return int64(0)
							} else {
								F_do_numeric_accum(m, v70, v136)
								mBase = m.M
								v141 = m.ExcPending
								if v141 != 0 {
									return int64(0)
								} else {
									m.G0 = v12 + int32(32)
									return base.I64_extend_i32_u(v70)
								}
							}
						}
					}
				} else {
					m.G0 = v12 + int32(32)
					return base.I64_extend_i32_u(v70)
				}
			}
		}
	}
}
func F_int8_bytea(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_cash_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_int8_cash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v42 int64
	_ = v42
	var v44 int32
	_ = v44
	var v51 int64
	_ = v51
	var v62 int64
	_ = v62
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v82 int64
	_ = v82
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v104 int64
	_ = v104
	var v111 int64
	_ = v111
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_PGLC_localeconv(m)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+41)))
		if base.Ui32(int32(10)) < base.Ui32(v19) {
			v22 = int32(2)
		} else {
			v22 = v19
		}
		v23 = base.I32_extend8_s(v22)
		if v23 <= int32(0) {
			v75 = int64(1)
		} else {
			v27 = int64(1)
			if base.Ui32(int32(8)) <= base.Ui32(v22) {
				v34 = v27
				v36 = int32(0)
				for {
					v42 = v34 * int64(100000000)
					v44 = v36 + int32(8)
					if v44 != v23&int32(120) {
						v34 = v42
						v36 = v44
						continue
					} else {
						break
					}
					break
				}
				if v22&int32(7) == int32(0) {
					v75 = v42
				} else {
					v51 = v42
					v62 = v51
					v64 = int32(0)
					for {
						v70 = v62 * int64(10)
						v72 = v64 + int32(1)
						if v72 != v23&int32(7) {
							v62 = v70
							v64 = v72
							continue
						} else {
							break
						}
						break
					}
					v75 = v70
				}
			} else {
				v51 = v27
				v62 = v51
				v64 = int32(0)
				for {
					v70 = v62 * int64(10)
					v72 = v64 + int32(1)
					if v72 != v23&int32(7) {
						v62 = v70
						v64 = v72
						continue
					} else {
						break
					}
					break
				}
				v75 = v70
			}
		}
		v82 = int64(63)
		v90 = int64(32)
		v91 = int64(base.Ui64(v75) >> (uint(v90) % 64))
		v93 = int64(base.Ui64(v13) >> (uint(v90) % 64))
		v96 = int64(4294967295)
		v97 = v75 & v96
		v99 = v13 & v96
		v100 = v97 * v99
		v104 = int64(base.Ui64(v100)>>(uint(v90)%64)) + v97*v93
		v111 = v99*v91 + v104&v96
		*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v13*(v75>>(uint(v82)%64)) + v13>>(uint(v82)%64)*v75 + v91*v93 + int64(base.Ui64(v104)>>(uint(v90)%64)) + int64(base.Ui64(v111)>>(uint(v90)%64))
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = v100&v96 | v111<<(uint(v90)%64)
		v122 = *(*int64)(unsafe.Add(mBase, uint32(v11)+8))
		v123 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		if v122 == v123>>(uint(int64(63))%64) {
			v146 = v123
			m.G0 = v11 + int32(16)
			return v146
		} else {
			v127 = int64(0)
			v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v129 = F_errsave_start(m, v128)
			mBase = m.M
			v130 = m.ExcPending
			if v130 != 0 {
				return int64(0)
			} else {
				if v129 == int32(0) {
					v146 = v127
					m.G0 = v11 + int32(16)
					return v146
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_int8_cash_0), int32(0))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, v128, int32(_a_F_int8_cash_1), int32(1245), int32(_a_F_int8_cash_2))
							mBase = m.M
							v144 = m.ExcPending
							if v144 != 0 {
								return int64(0)
							} else {
								v146 = v127
								m.G0 = v11 + int32(16)
								return v146
							}
						}
					}
				}
			}
		}
	}
}
func F_int8_increment(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	v5 = base.B2i32(l1 == int64(9223372036854775807))
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v5)
	if l1 == int64(9223372036854775807) {
		v10 = int64(0)
	} else {
		v10 = l1 + int64(1)
	}
	return v10
}
