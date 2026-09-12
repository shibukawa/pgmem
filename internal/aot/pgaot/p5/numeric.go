package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_numeric(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int64
	_ = v167
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		v18 = base.I32_extend16_s(v17)
		if base.Ui32(int32(49152)) <= base.Ui32(v17) {
			if v16 < int32(4) {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v61 = F_palloc(m, int32(base.Ui32(v58)>>(uint(int32(2))%32)))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v65 = int32(base.Ui32(v63) >> (uint(int32(2)) % 32))
					if v65 != 0 {
						v66 = F__emscripten_memcpy_bulkmem(m, v61, v12, v65)
						mBase = m.M
					} else {
					}
					v192 = v61
					m.G0 = v9 + int32(32)
					return v192
				}
			} else {
				if v18 == int32(-16384) {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v61 = F_palloc(m, int32(base.Ui32(v58)>>(uint(int32(2))%32)))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v65 = int32(base.Ui32(v63) >> (uint(int32(2)) % 32))
						if v65 != 0 {
							v66 = F__emscripten_memcpy_bulkmem(m, v61, v12, v65)
							mBase = m.M
						} else {
						}
						v192 = v61
						m.G0 = v9 + int32(32)
						return v192
					}
				} else {
					v26 = F_errsave_start(m, int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 == int32(0) {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v61 = F_palloc(m, int32(base.Ui32(v58)>>(uint(int32(2))%32)))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v65 = int32(base.Ui32(v63) >> (uint(int32(2)) % 32))
								if v65 != 0 {
									v66 = F__emscripten_memcpy_bulkmem(m, v61, v12, v65)
									mBase = m.M
								} else {
								}
								v192 = v61
								m.G0 = v9 + int32(32)
								return v192
							}
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(31369), int32(0))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(base.Ui32(v16-int32(4)) >> (uint(int32(16)) % 32))
									v42 = int32(21)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = (v16<<(uint(v42)%32) - int32(8388608)) >> (uint(v42) % 32)
									F_errdetail(m, int32(610461), v9)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										F_errsave_finish(m, int32(0), int32(496807), int32(8138), int32(312180))
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
											v61 = F_palloc(m, int32(base.Ui32(v58)>>(uint(int32(2))%32)))
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												v63 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
												v65 = int32(base.Ui32(v63) >> (uint(int32(2)) % 32))
												if v65 != 0 {
													v66 = F__emscripten_memcpy_bulkmem(m, v61, v12, v65)
													mBase = m.M
												} else {
												}
												v192 = v61
												m.G0 = v9 + int32(32)
												return v192
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
			if v16 <= int32(3) {
				v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v73 = F_palloc(m, int32(base.Ui32(v70)>>(uint(int32(2))%32)))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v77 = int32(base.Ui32(v75) >> (uint(int32(2)) % 32))
					if v77 != 0 {
						v78 = F__emscripten_memcpy_bulkmem(m, v73, v12, v77)
						mBase = m.M
					} else {
					}
					v192 = v73
					m.G0 = v9 + int32(32)
					return v192
				}
			} else {
				if v18 < int32(0) {
					v92 = v17<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v17&int32(63)
				} else {
					v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
					v92 = v91
				}
				v95 = int32(4)
				v101 = int32(21)
				v106 = (v16<<(uint(v101)%32) - int32(8388608)) >> (uint(v101) % 32)
				if int32(base.Ui32(v16-v95)>>(uint(int32(16))%32))-v106 < v92<<(uint(int32(2))%32)+v95 {
					v166 = v9 + int32(24)
					v167 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v166))) = v167
					*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v167
					*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v167
					F_set_var_from_num(m, v12, v9+int32(8))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int32(0)
					} else {
						v180 = F_apply_typmod(m, v9+int32(8), v16, int32(0))
						mBase = m.M
						v181 = m.ExcPending
						if v181 != 0 {
							return int32(0)
						} else {
							v185 = F_make_result_opt_error(m, v9+int32(8), int32(0))
							mBase = m.M
							v186 = m.ExcPending
							if v186 != 0 {
								return int32(0)
							} else {
								v187 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
								if v187 == int32(0) {
									v192 = v185
									m.G0 = v9 + int32(32)
									return v192
								} else {
									F_pfree(m, v187)
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return int32(0)
									} else {
										v192 = v185
										m.G0 = v9 + int32(32)
										return v192
									}
								}
							}
						}
					}
				} else {
					if int32(0) <= v18 {
						v117 = v17 & int32(16383)
					} else {
						v117 = int32(base.Ui32(v17)>>(uint(int32(7))%32)) & int32(63)
					}
					if v106 < v117 {
						v166 = v9 + int32(24)
						v167 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v166))) = v167
						*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v167
						*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v167
						F_set_var_from_num(m, v12, v9+int32(8))
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return int32(0)
						} else {
							v180 = F_apply_typmod(m, v9+int32(8), v16, int32(0))
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return int32(0)
							} else {
								v185 = F_make_result_opt_error(m, v9+int32(8), int32(0))
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									v187 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
									if v187 == int32(0) {
										v192 = v185
										m.G0 = v9 + int32(32)
										return v192
									} else {
										F_pfree(m, v187)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											v192 = v185
											m.G0 = v9 + int32(32)
											return v192
										}
									}
								}
							}
						}
					} else {
						if base.B2i32(v18 < int32(-16384))&base.B2i32(base.Ui32(int32(64)) <= base.Ui32(v106)) != 0 {
							v166 = v9 + int32(24)
							v167 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v166))) = v167
							*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v167
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v167
							F_set_var_from_num(m, v12, v9+int32(8))
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
								return int32(0)
							} else {
								v180 = F_apply_typmod(m, v9+int32(8), v16, int32(0))
								mBase = m.M
								v181 = m.ExcPending
								if v181 != 0 {
									return int32(0)
								} else {
									v185 = F_make_result_opt_error(m, v9+int32(8), int32(0))
									mBase = m.M
									v186 = m.ExcPending
									if v186 != 0 {
										return int32(0)
									} else {
										v187 = *(*int32)(unsafe.Add(mBase, uint32(v166)))
										if v187 == int32(0) {
											v192 = v185
											m.G0 = v9 + int32(32)
											return v192
										} else {
											F_pfree(m, v187)
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return int32(0)
											} else {
												v192 = v185
												m.G0 = v9 + int32(32)
												return v192
											}
										}
									}
								}
							}
						} else {
							v124 = int32(0)
							if v124 < v106 {
								v127 = v106
							} else {
								v127 = v124
							}
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v131 = F_palloc(m, int32(base.Ui32(v128)>>(uint(int32(2))%32)))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int32(0)
							} else {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v135 = int32(base.Ui32(v133) >> (uint(int32(2)) % 32))
								if v135 != 0 {
									v136 = F__emscripten_memcpy_bulkmem(m, v131, v12, v135)
									mBase = m.M
									v137 = v136
								} else {
									v137 = v131
								}
								v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
								if v138&int32(49152) == int32(32768) {
									v147 = v138&int32(41087) | v127<<(uint(int32(7))%32)
									*(*uint16)(unsafe.Add(mBase, uint32(v137)+4)) = uint16(v147)
									v192 = v131
								} else {
									v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v137)+4)))
									v150 = int32(49152)
									v151 = v149 & v150
									if v151 != v150 {
										if v151 != int32(32768) {
											v162 = v151
										} else {
											v162 = v149 << (uint(int32(1)) % 32) & int32(16384)
										}
									} else {
										v162 = v149 & int32(61440)
									}
									v163 = v162 | v127
									*(*uint16)(unsafe.Add(mBase, uint32(v137)+4)) = uint16(v163)
									v192 = v131
								}
								m.G0 = v9 + int32(32)
								return v192
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v35 int64
	_ = v35
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v52 int32
	_ = v52
	var v53 int64
	_ = v53
	var v59 int64
	_ = v59
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v70 int32
	_ = v70
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
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_PGLC_localeconv(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+41)))
		if base.Ui32(int32(10)) < base.Ui32(v13) {
			v16 = int32(2)
		} else {
			v16 = v13
		}
		v17 = base.I32_extend8_s(v16)
		if v17 <= int32(0) {
			v64 = int64(1)
		} else {
			if base.Ui32(v16) < base.Ui32(int32(8)) {
				v40 = int64(1)
			} else {
				v28 = int32(0)
				v29 = int64(1)
				for {
					v35 = v29 * int64(100000000)
					v37 = v28 + int32(8)
					if v37 != v17&int32(120) {
						v28 = v37
						v29 = v35
						continue
					} else {
						break
					}
					break
				}
				v40 = v35
			}
			if v16&int32(7) == int32(0) {
				v64 = v40
			} else {
				v52 = int32(0)
				v53 = v40
				for {
					v59 = v53 * int64(10)
					v61 = v52 + int32(1)
					if v61 != v17&int32(7) {
						v52 = v61
						v53 = v59
						continue
					} else {
						break
					}
					break
				}
				v64 = v59
			}
		}
		v70 = int32(0)
		v73 = F_int64_to_numeric(m, v64)
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			v75 = F_DirectFunctionCall2Coll(m, int32(1278), v70, v7, v73)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int32(0)
			} else {
				v77 = F_DirectFunctionCall1Coll(m, int32(1277), v70, v75)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					v79 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
					v80 = F_Int64GetDatum(m, v79)
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						return v80
					}
				}
			}
		}
	}
}
func F_numeric_float8_no_overflow(m *base.Module, l0 int32) int32 {
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
	var v24 float64
	_ = v24
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v6 = m.G0
	v8 = v6 - int32(32)
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
		if base.Ui32(int32(49152)) <= base.Ui32(v15) {
			if v16 == int32(-4096) {
				v24 = math.Float64frombits(uint64(0xfff0000000000000))
			} else {
				v24 = math.Float64frombits(uint64(0x7ff8000000000000))
			}
			if v16 == int32(-12288) {
				v27 = math.Float64frombits(uint64(0x7ff0000000000000))
			} else {
				v27 = v24
			}
			v86 = v27
			v87 = F_Float8GetDatum(m, v86)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(32)
				return v87
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v34 = base.B2i32(int32(0) <= v16)
			if int32(0) <= v16 {
				v35 = int32(-8)
			} else {
				v35 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(base.Ui32(int32(base.Ui32(v28)>>(uint(int32(2))%32))+v35) >> (uint(int32(1)) % 32))
			if int32(0) <= v16 {
				v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+6)))
				v50 = v40
			} else {
				v50 = v15<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v15&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v50
			v52 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v52
			v61 = base.B2i32(v16 < v52)
			if v16 < v52 {
				v62 = int32(base.Ui32(v15)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v62 = v15 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v62
			v69 = v15 & int32(49152)
			if v69 == int32(32768) {
				v72 = v15 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v72 = v69
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v72
			if v16 < v52 {
				v76 = int32(6)
			} else {
				v76 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v11 + v76
			v81 = F_numericvar_to_double_no_overflow(m, v8+int32(8))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v86 = v81
				v87 = F_Float8GetDatum(m, v86)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(32)
					return v87
				}
			}
		}
	}
}
func F_numeric_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_numeric_int4_opt_error(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_numeric_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_numeric_int8_opt_error(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_Int64GetDatum(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_numeric_mul_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v249 int64
	_ = v249
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v341 int32
	_ = v341
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v15 = base.I32_extend16_s(v14)
	if base.Ui32(v14) <= base.Ui32(int32(49151)) {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v19 = base.I32_extend16_s(v18)
		if base.Ui32(int32(49151)) < base.Ui32(v18) {
			v46 = v19
			if v46&int32(65535) != int32(49152) {
				if v14 != int32(61440) {
					if v14 != int32(53248) {
						v131 = v14 & int32(49152)
						if v46&int32(65535) == int32(53248) {
							if v131 == int32(49152) {
								v164 = F_make_result_opt_error(m, int32(1721372), int32(0))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return int32(0)
								} else {
									v425 = v164
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if int32(0) <= v15 {
									v145 = int32(-8)
								} else {
									v145 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v138)>>(uint(int32(2))%32))+v145) < base.Ui32(int32(2)) {
									v408 = F_make_result_opt_error(m, int32(1721324), int32(0))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return int32(0)
									} else {
										v425 = v408
										m.G0 = v12 + int32(80)
										return v425
									}
								} else {
									if v131 == int32(32768) {
										v155 = v14 << (uint(int32(1)) % 32) & int32(16384)
									} else {
										v155 = v131
									}
									if v155 == int32(16384) {
										v164 = F_make_result_opt_error(m, int32(1721372), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int32(0)
										} else {
											v425 = v164
											m.G0 = v12 + int32(80)
											return v425
										}
									} else {
										v160 = F_make_result_opt_error(m, int32(1721348), int32(0))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											v425 = v160
											m.G0 = v12 + int32(80)
											return v425
										}
									}
								}
							}
						} else {
							if v131 == int32(49152) {
								v194 = F_make_result_opt_error(m, int32(1721348), int32(0))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									v425 = v194
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if int32(0) <= v15 {
									v175 = int32(-8)
								} else {
									v175 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v168)>>(uint(int32(2))%32))+v175) < base.Ui32(int32(2)) {
									v412 = F_make_result_opt_error(m, int32(1721324), int32(0))
									mBase = m.M
									v413 = m.ExcPending
									if v413 != 0 {
										return int32(0)
									} else {
										v425 = v412
										m.G0 = v12 + int32(80)
										return v425
									}
								} else {
									if v131 == int32(32768) {
										v185 = v14 << (uint(int32(1)) % 32) & int32(16384)
									} else {
										v185 = v131
									}
									if v185 == int32(16384) {
										v194 = F_make_result_opt_error(m, int32(1721348), int32(0))
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
											return int32(0)
										} else {
											v425 = v194
											m.G0 = v12 + int32(80)
											return v425
										}
									} else {
										v190 = F_make_result_opt_error(m, int32(1721372), int32(0))
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											v425 = v190
											m.G0 = v12 + int32(80)
											return v425
										}
									}
								}
							}
						}
					} else {
						v65 = v46 & int32(65535)
						v66 = int32(49152)
						v67 = v46 & v66
						if v67 == v66 {
							if v65 == int32(53248) {
								v95 = F_make_result_opt_error(m, int32(1721348), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v425 = v95
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								v420 = F_make_result_opt_error(m, int32(1721372), int32(0))
								mBase = m.M
								v421 = m.ExcPending
								if v421 != 0 {
									return int32(0)
								} else {
									v425 = v420
									m.G0 = v12 + int32(80)
									return v425
								}
							}
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if int32(0) <= base.I32_extend16_s(v46) {
								v80 = int32(-8)
							} else {
								v80 = int32(-6)
							}
							if base.Ui32(int32(base.Ui32(v72)>>(uint(int32(2))%32))+v80) < base.Ui32(int32(2)) {
								v400 = F_make_result_opt_error(m, int32(1721324), int32(0))
								mBase = m.M
								v401 = m.ExcPending
								if v401 != 0 {
									return int32(0)
								} else {
									v425 = v400
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								if v67 == int32(32768) {
									v90 = v65 << (uint(int32(1)) % 32) & int32(16384)
								} else {
									v90 = v67
								}
								if v90 == int32(16384) {
									v420 = F_make_result_opt_error(m, int32(1721372), int32(0))
									mBase = m.M
									v421 = m.ExcPending
									if v421 != 0 {
										return int32(0)
									} else {
										v425 = v420
										m.G0 = v12 + int32(80)
										return v425
									}
								} else {
									v95 = F_make_result_opt_error(m, int32(1721348), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v425 = v95
										m.G0 = v12 + int32(80)
										return v425
									}
								}
							}
						}
					}
				} else {
					v98 = v46 & int32(65535)
					v99 = int32(49152)
					v100 = v46 & v99
					if v100 == v99 {
						if v98 == int32(53248) {
							v128 = F_make_result_opt_error(m, int32(1721372), int32(0))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								v425 = v128
								m.G0 = v12 + int32(80)
								return v425
							}
						} else {
							v416 = F_make_result_opt_error(m, int32(1721348), int32(0))
							mBase = m.M
							v417 = m.ExcPending
							if v417 != 0 {
								return int32(0)
							} else {
								v425 = v416
								m.G0 = v12 + int32(80)
								return v425
							}
						}
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v46) {
							v113 = int32(-8)
						} else {
							v113 = int32(-6)
						}
						if base.Ui32(int32(base.Ui32(v105)>>(uint(int32(2))%32))+v113) < base.Ui32(int32(2)) {
							v404 = F_make_result_opt_error(m, int32(1721324), int32(0))
							mBase = m.M
							v405 = m.ExcPending
							if v405 != 0 {
								return int32(0)
							} else {
								v425 = v404
								m.G0 = v12 + int32(80)
								return v425
							}
						} else {
							if v100 == int32(32768) {
								v123 = v98 << (uint(int32(1)) % 32) & int32(16384)
							} else {
								v123 = v100
							}
							if v123 == int32(16384) {
								v416 = F_make_result_opt_error(m, int32(1721348), int32(0))
								mBase = m.M
								v417 = m.ExcPending
								if v417 != 0 {
									return int32(0)
								} else {
									v425 = v416
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								v128 = F_make_result_opt_error(m, int32(1721372), int32(0))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									v425 = v128
									m.G0 = v12 + int32(80)
									return v425
								}
							}
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(1721324), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v425 = v56
					m.G0 = v12 + int32(80)
					return v425
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
				v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v197 = v196
			} else {
				v197 = v14<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v14&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v197
			v199 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v199
			v208 = base.B2i32(v15 < v199)
			if v15 < v199 {
				v209 = int32(base.Ui32(v14)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v209 = v14 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v209
			v216 = v14 & int32(49152)
			if v216 == int32(32768) {
				v219 = v14 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v219 = v216
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v219
			if v15 < v199 {
				v223 = int32(6)
			} else {
				v223 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l0 + v223
			v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v232 = base.B2i32(int32(0) <= v19)
			if int32(0) <= v19 {
				v233 = int32(-8)
			} else {
				v233 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(base.Ui32(int32(base.Ui32(v226)>>(uint(int32(2))%32))+v233) >> (uint(int32(1)) % 32))
			if int32(0) <= v19 {
				v238 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				v248 = v238
			} else {
				v248 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
			}
			v249 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v249
			*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v249
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v248
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v249
			v261 = v18 & int32(49152)
			if v261 == int32(32768) {
				v264 = v18 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v264 = v261
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v264
			v266 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v266
			v271 = base.B2i32(v19 < v266)
			if v19 < v266 {
				v272 = int32(6)
			} else {
				v272 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l1 + v272
			if v19 < v266 {
				v281 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v281 = v18 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v281
			F_mul_var(m, v12+int32(56), v12+int32(32), v12+int32(8), v281+v209)
			mBase = m.M
			v291 = m.ExcPending
			if v291 != 0 {
				return int32(0)
			} else {
				v292 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				if v292 < int32(16384) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(16383)
					v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					v299 = v297 << (uint(int32(2)) % 32)
					if v299+int32(16387) < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
					} else {
						v311 = base.I32_div_s(v299+int32(16390), int32(4))
						v312 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v312 < v311 {
						} else {
							v314 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v311
							v316 = int32(1)
							v317 = v311 - v316
							v320 = v314 + v317<<(uint(v316)%32)
							v321 = int32(*(*int16)(unsafe.Add(mBase, uint32(v320))))
							v323 = base.I32_rem_s(v321, int32(10))
							v324 = v321 - v323
							*(*uint16)(unsafe.Add(mBase, uint32(v320))) = uint16(v324)
							if v323 < int32(5) {
								v363 = v317
							} else {
								v330 = base.I32_extend16_s(v324)
								if int32(9989) < v330 {
									v333 = int32(-9990)
								} else {
									v333 = int32(10)
								}
								v334 = v333 + v330
								*(*uint16)(unsafe.Add(mBase, uint32(v320))) = uint16(v334)
								if v330 < int32(9990) {
									v363 = v317
								} else {
									v341 = v317
									for {
										v347 = int32(1)
										v348 = v341 - v347
										v351 = v314 + v348<<(uint(v347)%32)
										v354 = int32(*(*int16)(unsafe.Add(mBase, uint32(v351))))
										v356 = base.B2i32(int32(9998) < v354)
										if int32(9998) < v354 {
											v357 = int32(-9999)
										} else {
											v357 = v347
										}
										v358 = v357 + v354
										*(*uint16)(unsafe.Add(mBase, uint32(v351))) = uint16(v358)
										if int32(9998) < v354 {
											v341 = v348
											continue
										} else {
											break
										}
										break
									}
									v363 = v348
								}
							}
							if int32(0) <= v363 {
							} else {
								v371 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v297 + v371
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v311 + v371
								*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v314 - int32(2)
							}
						}
					}
				}
				v391 = F_make_result_opt_error(m, v12+int32(8), l2)
				mBase = m.M
				v392 = m.ExcPending
				if v392 != 0 {
					return int32(0)
				} else {
					v393 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					if v393 == int32(0) {
						v425 = v391
						m.G0 = v12 + int32(80)
						return v425
					} else {
						F_pfree(m, v393)
						mBase = m.M
						v397 = m.ExcPending
						if v397 != 0 {
							return int32(0)
						} else {
							v425 = v391
							m.G0 = v12 + int32(80)
							return v425
						}
					}
				}
			}
		}
	} else {
		if v15 == int32(-16384) {
			v56 = F_make_result_opt_error(m, int32(1721324), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v425 = v56
				m.G0 = v12 + int32(80)
				return v425
			}
		} else {
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v46 = v45
			if v46&int32(65535) != int32(49152) {
				if v14 != int32(61440) {
					if v14 != int32(53248) {
						v131 = v14 & int32(49152)
						if v46&int32(65535) == int32(53248) {
							if v131 == int32(49152) {
								v164 = F_make_result_opt_error(m, int32(1721372), int32(0))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return int32(0)
								} else {
									v425 = v164
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if int32(0) <= v15 {
									v145 = int32(-8)
								} else {
									v145 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v138)>>(uint(int32(2))%32))+v145) < base.Ui32(int32(2)) {
									v408 = F_make_result_opt_error(m, int32(1721324), int32(0))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return int32(0)
									} else {
										v425 = v408
										m.G0 = v12 + int32(80)
										return v425
									}
								} else {
									if v131 == int32(32768) {
										v155 = v14 << (uint(int32(1)) % 32) & int32(16384)
									} else {
										v155 = v131
									}
									if v155 == int32(16384) {
										v164 = F_make_result_opt_error(m, int32(1721372), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int32(0)
										} else {
											v425 = v164
											m.G0 = v12 + int32(80)
											return v425
										}
									} else {
										v160 = F_make_result_opt_error(m, int32(1721348), int32(0))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											v425 = v160
											m.G0 = v12 + int32(80)
											return v425
										}
									}
								}
							}
						} else {
							if v131 == int32(49152) {
								v194 = F_make_result_opt_error(m, int32(1721348), int32(0))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									v425 = v194
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if int32(0) <= v15 {
									v175 = int32(-8)
								} else {
									v175 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v168)>>(uint(int32(2))%32))+v175) < base.Ui32(int32(2)) {
									v412 = F_make_result_opt_error(m, int32(1721324), int32(0))
									mBase = m.M
									v413 = m.ExcPending
									if v413 != 0 {
										return int32(0)
									} else {
										v425 = v412
										m.G0 = v12 + int32(80)
										return v425
									}
								} else {
									if v131 == int32(32768) {
										v185 = v14 << (uint(int32(1)) % 32) & int32(16384)
									} else {
										v185 = v131
									}
									if v185 == int32(16384) {
										v194 = F_make_result_opt_error(m, int32(1721348), int32(0))
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
											return int32(0)
										} else {
											v425 = v194
											m.G0 = v12 + int32(80)
											return v425
										}
									} else {
										v190 = F_make_result_opt_error(m, int32(1721372), int32(0))
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											v425 = v190
											m.G0 = v12 + int32(80)
											return v425
										}
									}
								}
							}
						}
					} else {
						v65 = v46 & int32(65535)
						v66 = int32(49152)
						v67 = v46 & v66
						if v67 == v66 {
							if v65 == int32(53248) {
								v95 = F_make_result_opt_error(m, int32(1721348), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v425 = v95
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								v420 = F_make_result_opt_error(m, int32(1721372), int32(0))
								mBase = m.M
								v421 = m.ExcPending
								if v421 != 0 {
									return int32(0)
								} else {
									v425 = v420
									m.G0 = v12 + int32(80)
									return v425
								}
							}
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if int32(0) <= base.I32_extend16_s(v46) {
								v80 = int32(-8)
							} else {
								v80 = int32(-6)
							}
							if base.Ui32(int32(base.Ui32(v72)>>(uint(int32(2))%32))+v80) < base.Ui32(int32(2)) {
								v400 = F_make_result_opt_error(m, int32(1721324), int32(0))
								mBase = m.M
								v401 = m.ExcPending
								if v401 != 0 {
									return int32(0)
								} else {
									v425 = v400
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								if v67 == int32(32768) {
									v90 = v65 << (uint(int32(1)) % 32) & int32(16384)
								} else {
									v90 = v67
								}
								if v90 == int32(16384) {
									v420 = F_make_result_opt_error(m, int32(1721372), int32(0))
									mBase = m.M
									v421 = m.ExcPending
									if v421 != 0 {
										return int32(0)
									} else {
										v425 = v420
										m.G0 = v12 + int32(80)
										return v425
									}
								} else {
									v95 = F_make_result_opt_error(m, int32(1721348), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v425 = v95
										m.G0 = v12 + int32(80)
										return v425
									}
								}
							}
						}
					}
				} else {
					v98 = v46 & int32(65535)
					v99 = int32(49152)
					v100 = v46 & v99
					if v100 == v99 {
						if v98 == int32(53248) {
							v128 = F_make_result_opt_error(m, int32(1721372), int32(0))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								v425 = v128
								m.G0 = v12 + int32(80)
								return v425
							}
						} else {
							v416 = F_make_result_opt_error(m, int32(1721348), int32(0))
							mBase = m.M
							v417 = m.ExcPending
							if v417 != 0 {
								return int32(0)
							} else {
								v425 = v416
								m.G0 = v12 + int32(80)
								return v425
							}
						}
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v46) {
							v113 = int32(-8)
						} else {
							v113 = int32(-6)
						}
						if base.Ui32(int32(base.Ui32(v105)>>(uint(int32(2))%32))+v113) < base.Ui32(int32(2)) {
							v404 = F_make_result_opt_error(m, int32(1721324), int32(0))
							mBase = m.M
							v405 = m.ExcPending
							if v405 != 0 {
								return int32(0)
							} else {
								v425 = v404
								m.G0 = v12 + int32(80)
								return v425
							}
						} else {
							if v100 == int32(32768) {
								v123 = v98 << (uint(int32(1)) % 32) & int32(16384)
							} else {
								v123 = v100
							}
							if v123 == int32(16384) {
								v416 = F_make_result_opt_error(m, int32(1721348), int32(0))
								mBase = m.M
								v417 = m.ExcPending
								if v417 != 0 {
									return int32(0)
								} else {
									v425 = v416
									m.G0 = v12 + int32(80)
									return v425
								}
							} else {
								v128 = F_make_result_opt_error(m, int32(1721372), int32(0))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									v425 = v128
									m.G0 = v12 + int32(80)
									return v425
								}
							}
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(1721324), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v425 = v56
					m.G0 = v12 + int32(80)
					return v425
				}
			}
		}
	}
}
func F_numeric_pg_lsn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v212 int32
	_ = v212
	var v220 int64
	_ = v220
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v224 int64
	_ = v224
	var v230 int64
	_ = v230
	var v233 int64
	_ = v233
	var v236 int64
	_ = v236
	var v239 int64
	_ = v239
	var v240 int64
	_ = v240
	var v244 int64
	_ = v244
	var v251 int64
	_ = v251
	var v262 int64
	_ = v262
	var v265 int64
	_ = v265
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v273 int64
	_ = v273
	var v275 int32
	_ = v275
	var v286 int64
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
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
	if base.Ui32(int32(49152)) <= base.Ui32(v20) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_pfree(m, v69)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L77
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(243453)
	F_errmsg(m, int32(182290), v13)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L75
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v51 = base.B2i32(int32(0) <= v21)
	if int32(0) <= v21 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v21 == int32(-16384) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(243453)
	F_errmsg(m, int32(181613), v13+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(496807), int32(4884), int32(243388))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
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
	v52 = int32(-8)
	goto L15
L14:
	;
	v52 = int32(-6)
	goto L15
L15:
	;
	v53 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) + v52
	if int32(0) <= v21 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
	v64 = v54
	goto L18
L17:
	;
	v64 = v20<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v20&int32(63)
	goto L18
L18:
	;
	v66 = v53 & int32(-2)
	v69 = F_palloc(m, v66+int32(2))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v71 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v69))) = uint16(v71)
	if base.Ui32(int32(2)) <= base.Ui32(v53) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v21 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v64 < int32(-1) {
		v286 = v10
		goto L30
	} else {
		goto L31
	}
L23:
	;
	v81 = int32(6)
	goto L25
L24:
	;
	v81 = int32(8)
	goto L25
L25:
	;
	if v66 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L22
L27:
	;
	v83 = F__emscripten_memcpy_bulkmem(m, v69+int32(2), v16+v81, v66)
	mBase = m.M
	goto L29
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	F_pfree(m, v69)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L73
	}
L31:
	;
	v88 = v69 + int32(2)
	v89 = int32(1)
	v92 = (v64 + v89) & int32(1073741823)
	v94 = int32(base.Ui32(v53) >> (uint(v89) % 32))
	if base.Ui32(v94) <= base.Ui32(v92) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	if int32(0) < v131 {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	v131 = v94
	v133 = v88
	v134 = v64
	goto L32
L34:
	;
	goto L35
L35:
	;
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v88+v92<<(uint(int32(1))%32)))))
	if v99 < int32(5000) {
		v131 = v92
		v133 = v88
		v134 = v64
		goto L32
	} else {
		goto L36
	}
L36:
	;
	v103 = v92
	goto L37
L37:
	;
	v112 = int32(1)
	v114 = v69 + v103<<(uint(v112)%32)
	v117 = int32(*(*int16)(unsafe.Add(mBase, uint32(v114))))
	v119 = base.B2i32(int32(9998) < v117)
	if int32(9998) < v117 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if int32(0) <= v124 {
		v131 = v92
		v133 = v88
		v134 = v64
		goto L32
	} else {
		goto L43
	}
L39:
	;
	v120 = int32(-9999)
	goto L41
L40:
	;
	v120 = v112
	goto L41
L41:
	;
	v121 = v120 + v117
	*(*uint16)(unsafe.Add(mBase, uint32(v114))) = uint16(v121)
	v124 = v103 - int32(1)
	if int32(9998) < v117 {
		v103 = v124
		goto L37
	} else {
		goto L42
	}
L42:
	;
	goto L38
L43:
	;
	v127 = int32(1)
	v131 = v92 + v127
	v133 = v69
	v134 = v64 + v127
	goto L32
L44:
	;
	v201 = v20 & int32(49152)
	if v201 == int32(32768) {
		goto L59
	} else {
		goto L60
	}
L45:
	;
	v143 = v131
	v145 = v133
	v146 = v134
	goto L48
L46:
	;
	goto L47
L47:
	;
	if v131 == int32(0) {
		v286 = v10
		goto L30
	} else {
		goto L58
	}
L48:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145))))
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v166 = v143
	goto L54
L50:
	;
	v156 = int32(1)
	if v156 < v143 {
		v143 = v143 - v156
		v145 = v145 + int32(2)
		v146 = v146 - v156
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	goto L49
L53:
	;
	v286 = v10
	goto L30
L54:
	;
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145-int32(2)+v166<<(uint(int32(1))%32)))))
	if v179 != 0 {
		v186 = v166
		v188 = v145
		v189 = v146
		goto L44
	} else {
		goto L56
	}
L55:
	;
	v286 = v10
	goto L30
L56:
	;
	v180 = int32(1)
	if v180 < v166 {
		v166 = v166 - v180
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v186 = v131
	v188 = v133
	v189 = v134
	goto L44
L59:
	;
	v204 = v20 << (uint(int32(1)) % 32) & int32(16384)
	goto L61
L60:
	;
	v204 = v201
	goto L61
L61:
	;
	if v204 == int32(16384) {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	v208 = int64(*(*int16)(unsafe.Add(mBase, uint32(v188))))
	if v189 <= int32(0) {
		v286 = v208
		goto L30
	} else {
		goto L63
	}
L63:
	;
	v212 = int32(1)
	v220 = v208
	goto L64
L64:
	;
	v222 = v13 + int32(32)
	v223 = int64(0)
	v224 = int64(10000)
	v230 = int64(32)
	v233 = int64(base.Ui64(v220) >> (uint(v230) % 64))
	v236 = int64(4294967295)
	v239 = v220 & v236
	v240 = v224 * v239
	v244 = int64(base.Ui64(v240)>>(uint(v230)%64)) + v224*v233
	v251 = v239*v223 + v244&v236
	*(*int64)(unsafe.Add(mBase, uint32(v222)+8)) = v220*v223 + v223 + v223*v233 + int64(base.Ui64(v244)>>(uint(v230)%64)) + int64(base.Ui64(v251)>>(uint(v230)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v222))) = v240&v236 | v251<<(uint(v230)%64)
	goto L66
L65:
	;
	v286 = v273
	goto L30
L66:
	;
	v262 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	if v262 != int64(0) {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v265 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	if v212 < v186 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v270 = int64(*(*int16)(unsafe.Add(mBase, uint32(v188+v212<<(uint(int32(1))%32)))))
	v271 = v265 + v270
	if base.Ui64(v271) < base.Ui64(v265) {
		goto L3
	} else {
		goto L71
	}
L69:
	;
	v273 = v265
	goto L70
L70:
	;
	v275 = v212 + int32(1)
	if v275 <= v189 {
		v212 = v275
		v220 = v273
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v273 = v271
	goto L70
L72:
	;
	goto L65
L73:
	;
	v289 = F_Int64GetDatum(m, v286)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	m.G0 = v13 + int32(48)
	return v289
L75:
	;
	F_errfinish(m, int32(496807), int32(4880), int32(243388))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(400056), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(496807), int32(4893), int32(243388))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_numeric_poly_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v122 int64
	_ = v122
	var v130 int64
	_ = v130
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v166 int64
	_ = v166
	var v167 int64
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int64
	_ = v192
	var v197 int64
	_ = v197
	var v200 int64
	_ = v200
	var v203 int64
	_ = v203
	var v206 int64
	_ = v206
	var v207 int64
	_ = v207
	var v211 int64
	_ = v211
	var v218 int64
	_ = v218
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v236 int64
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v265 int32
	_ = v265
	var v266 int64
	_ = v266
	var v267 int64
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v284 int64
	_ = v284
	var v292 int64
	_ = v292
	var v299 int64
	_ = v299
	var v300 int64
	_ = v300
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v315 int64
	_ = v315
	var v316 int64
	_ = v316
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v340 int64
	_ = v340
	var v341 int64
	_ = v341
	var v346 int64
	_ = v346
	var v349 int64
	_ = v349
	var v352 int64
	_ = v352
	var v355 int64
	_ = v355
	var v356 int64
	_ = v356
	var v360 int64
	_ = v360
	var v367 int64
	_ = v367
	var v379 int32
	_ = v379
	var v380 int64
	_ = v380
	var v381 int64
	_ = v381
	var v385 int64
	_ = v385
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(112)
	m.G0 = v17
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 == v2 {
		v49 = int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
		switch v24 - int32(429) {
		case 0:
			v49 = int32(1)
		case 1:
			v49 = int32(2)
		default:
			v49 = int32(0)
		}
	}
	if v49 != 0 {
		v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = int64(0)
		F_pq_begintypsend(m, v17+int32(96))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
			F_enlargeStringInfo(m, v17+int32(96), int32(8))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int32(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v17)+96))
				v68 = int64(56)
				v70 = int64(65280)
				v72 = int64(40)
				v75 = int64(16711680)
				v77 = int64(24)
				v79 = int64(4278190080)
				v81 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v65+v66))) = v59<<(uint(v68)%64) | v59&v70<<(uint(v72)%64) | (v59&v75<<(uint(v77)%64) | v59&v79<<(uint(v81)%64)) | (int64(base.Ui64(v59)>>(uint(v81)%64))&v79 | int64(base.Ui64(v59)>>(uint(v77)%64))&v75 | (int64(base.Ui64(v59)>>(uint(v72)%64))&v70 | int64(base.Ui64(v59)>>(uint(v68)%64))))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+100)) = v65 + int32(8)
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v50)+24))
				v108 = *(*int64)(unsafe.Add(mBase, uint32(v50)+16))
				v110 = F_palloc(m, int32(22))
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v110
					v113 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v110))) = uint16(v113)
					*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v110 + int32(2)
					if v107 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = int64(16384)
						v122 = int64(0)
						v149 = v122 - (v107 + base.I64_extend_i32_u(base.B2i32(v108 != v122)))
						v150 = v122 - v108
						v156 = int32(0)
						v159 = v110 + int32(22)
						v166 = v149
						v167 = v150
						for {
							v171 = v17 + int32(56)
							v174 = m.G0
							v175 = int32(16)
							v176 = v174 - v175
							m.G0 = v176
							F___udivmodti4(m, v176, v167, v166, int64(10000), int64(0))
							mBase = m.M
							v180 = *(*int64)(unsafe.Add(mBase, uint32(v176)))
							v181 = *(*int64)(unsafe.Add(mBase, uint32(v176)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v181
							*(*int64)(unsafe.Add(mBase, uint32(v171))) = v180
							m.G0 = v176 + v175
							v188 = v17 + int32(40)
							v189 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
							v190 = *(*int64)(unsafe.Add(mBase, uint32(v17-int32(-64))))
							v191 = int64(55536)
							v192 = int64(0)
							v197 = int64(32)
							v200 = int64(base.Ui64(v189) >> (uint(v197) % 64))
							v203 = int64(4294967295)
							v206 = v189 & v203
							v207 = v191 * v206
							v211 = int64(base.Ui64(v207)>>(uint(v197)%64)) + v191*v200
							v218 = v206*v192 + v211&v203
							*(*int64)(unsafe.Add(mBase, uint32(v188)+8)) = v189*v192 + v190*v191 + v192*v200 + int64(base.Ui64(v211)>>(uint(v197)%64)) + int64(base.Ui64(v218)>>(uint(v197)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v188))) = v207&v203 | v218<<(uint(v197)%64)
							v230 = v159 - int32(2)
							v231 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
							v232 = v231 + v167
							*(*uint16)(unsafe.Add(mBase, uint32(v230))) = uint16(v232)
							v236 = int64(0)
							v241 = v156 + int32(1)
							if v166 == v236 {
								v242 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v167))
							} else {
								v242 = base.B2i32(v166 != v236)
							}
							if v242 != 0 {
								v156 = v241
								v159 = v230
								v166 = v190
								v167 = v189
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v230
						v244 = v241
						v248 = v156
					} else {
						v130 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v130
						if v107|v108 != v130 {
							v149 = v107
							v150 = v108
							v156 = int32(0)
							v159 = v110 + int32(22)
							v166 = v149
							v167 = v150
							for {
								v171 = v17 + int32(56)
								v174 = m.G0
								v175 = int32(16)
								v176 = v174 - v175
								m.G0 = v176
								F___udivmodti4(m, v176, v167, v166, int64(10000), int64(0))
								mBase = m.M
								v180 = *(*int64)(unsafe.Add(mBase, uint32(v176)))
								v181 = *(*int64)(unsafe.Add(mBase, uint32(v176)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v171)+8)) = v181
								*(*int64)(unsafe.Add(mBase, uint32(v171))) = v180
								m.G0 = v176 + v175
								v188 = v17 + int32(40)
								v189 = *(*int64)(unsafe.Add(mBase, uint32(v17)+56))
								v190 = *(*int64)(unsafe.Add(mBase, uint32(v17-int32(-64))))
								v191 = int64(55536)
								v192 = int64(0)
								v197 = int64(32)
								v200 = int64(base.Ui64(v189) >> (uint(v197) % 64))
								v203 = int64(4294967295)
								v206 = v189 & v203
								v207 = v191 * v206
								v211 = int64(base.Ui64(v207)>>(uint(v197)%64)) + v191*v200
								v218 = v206*v192 + v211&v203
								*(*int64)(unsafe.Add(mBase, uint32(v188)+8)) = v189*v192 + v190*v191 + v192*v200 + int64(base.Ui64(v211)>>(uint(v197)%64)) + int64(base.Ui64(v218)>>(uint(v197)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v188))) = v207&v203 | v218<<(uint(v197)%64)
								v230 = v159 - int32(2)
								v231 = *(*int64)(unsafe.Add(mBase, uint32(v17)+40))
								v232 = v231 + v167
								*(*uint16)(unsafe.Add(mBase, uint32(v230))) = uint16(v232)
								v236 = int64(0)
								v241 = v156 + int32(1)
								if v166 == v236 {
									v242 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v167))
								} else {
									v242 = base.B2i32(v166 != v236)
								}
								if v242 != 0 {
									v156 = v241
									v159 = v230
									v166 = v190
									v167 = v189
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v230
							v244 = v241
							v248 = v156
						} else {
							v244 = int32(0)
							v248 = v2
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v248
					*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v244
					F_numericvar_serialize(m, v17+int32(96), v17+int32(72))
					mBase = m.M
					v265 = m.ExcPending
					if v265 != 0 {
						return int32(0)
					} else {
						v266 = *(*int64)(unsafe.Add(mBase, uint32(v50)+40))
						v267 = *(*int64)(unsafe.Add(mBase, uint32(v50)+32))
						F_pfree(m, v110)
						mBase = m.M
						v269 = m.ExcPending
						if v269 != 0 {
							return int32(0)
						} else {
							v271 = F_palloc(m, int32(22))
							mBase = m.M
							v272 = m.ExcPending
							if v272 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v17)+88)) = v271
								v274 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v271))) = uint16(v274)
								*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v271 + int32(2)
								if v266 < int64(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = int64(16384)
									v284 = int64(0)
									v299 = v284 - (v266 + base.I64_extend_i32_u(base.B2i32(v267 != v284)))
									v300 = v284 - v267
									v305 = v274
									v308 = v271 + int32(22)
									v315 = v299
									v316 = v300
									for {
										v320 = v17 + int32(24)
										v323 = m.G0
										v324 = int32(16)
										v325 = v323 - v324
										m.G0 = v325
										F___udivmodti4(m, v325, v316, v315, int64(10000), int64(0))
										mBase = m.M
										v329 = *(*int64)(unsafe.Add(mBase, uint32(v325)))
										v330 = *(*int64)(unsafe.Add(mBase, uint32(v325)+8))
										*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v330
										*(*int64)(unsafe.Add(mBase, uint32(v320))) = v329
										m.G0 = v325 + v324
										v337 = v17 + int32(8)
										v338 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
										v339 = *(*int64)(unsafe.Add(mBase, uint32(v17+int32(32))))
										v340 = int64(55536)
										v341 = int64(0)
										v346 = int64(32)
										v349 = int64(base.Ui64(v338) >> (uint(v346) % 64))
										v352 = int64(4294967295)
										v355 = v338 & v352
										v356 = v340 * v355
										v360 = int64(base.Ui64(v356)>>(uint(v346)%64)) + v340*v349
										v367 = v355*v341 + v360&v352
										*(*int64)(unsafe.Add(mBase, uint32(v337)+8)) = v338*v341 + v339*v340 + v341*v349 + int64(base.Ui64(v360)>>(uint(v346)%64)) + int64(base.Ui64(v367)>>(uint(v346)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v337))) = v356&v352 | v367<<(uint(v346)%64)
										v379 = v308 - int32(2)
										v380 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
										v381 = v380 + v316
										*(*uint16)(unsafe.Add(mBase, uint32(v379))) = uint16(v381)
										v385 = int64(0)
										v390 = v305 + int32(1)
										if v315 == v385 {
											v391 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v316))
										} else {
											v391 = base.B2i32(v315 != v385)
										}
										if v391 != 0 {
											v305 = v390
											v308 = v379
											v315 = v339
											v316 = v338
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v379
									v393 = v390
									v397 = v305
								} else {
									v292 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = v292
									if v266|v267 == v292 {
										v393 = v274
										v397 = int32(0)
									} else {
										v299 = v266
										v300 = v267
										v305 = v274
										v308 = v271 + int32(22)
										v315 = v299
										v316 = v300
										for {
											v320 = v17 + int32(24)
											v323 = m.G0
											v324 = int32(16)
											v325 = v323 - v324
											m.G0 = v325
											F___udivmodti4(m, v325, v316, v315, int64(10000), int64(0))
											mBase = m.M
											v329 = *(*int64)(unsafe.Add(mBase, uint32(v325)))
											v330 = *(*int64)(unsafe.Add(mBase, uint32(v325)+8))
											*(*int64)(unsafe.Add(mBase, uint32(v320)+8)) = v330
											*(*int64)(unsafe.Add(mBase, uint32(v320))) = v329
											m.G0 = v325 + v324
											v337 = v17 + int32(8)
											v338 = *(*int64)(unsafe.Add(mBase, uint32(v17)+24))
											v339 = *(*int64)(unsafe.Add(mBase, uint32(v17+int32(32))))
											v340 = int64(55536)
											v341 = int64(0)
											v346 = int64(32)
											v349 = int64(base.Ui64(v338) >> (uint(v346) % 64))
											v352 = int64(4294967295)
											v355 = v338 & v352
											v356 = v340 * v355
											v360 = int64(base.Ui64(v356)>>(uint(v346)%64)) + v340*v349
											v367 = v355*v341 + v360&v352
											*(*int64)(unsafe.Add(mBase, uint32(v337)+8)) = v338*v341 + v339*v340 + v341*v349 + int64(base.Ui64(v360)>>(uint(v346)%64)) + int64(base.Ui64(v367)>>(uint(v346)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v337))) = v356&v352 | v367<<(uint(v346)%64)
											v379 = v308 - int32(2)
											v380 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
											v381 = v380 + v316
											*(*uint16)(unsafe.Add(mBase, uint32(v379))) = uint16(v381)
											v385 = int64(0)
											v390 = v305 + int32(1)
											if v315 == v385 {
												v391 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v316))
											} else {
												v391 = base.B2i32(v315 != v385)
											}
											if v391 != 0 {
												v305 = v390
												v308 = v379
												v315 = v339
												v316 = v338
												continue
											} else {
												break
											}
											break
										}
										*(*int32)(unsafe.Add(mBase, uint32(v17)+92)) = v379
										v393 = v390
										v397 = v305
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v397
								*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v393
								F_numericvar_serialize(m, v17+int32(96), v17+int32(72))
								mBase = m.M
								v414 = m.ExcPending
								if v414 != 0 {
									return int32(0)
								} else {
									v416 = v17 + int32(96)
									v418 = *(*int32)(unsafe.Add(mBase, uint32(v416)))
									v419 = *(*int32)(unsafe.Add(mBase, uint32(v416)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v418))) = v419 << (uint(int32(2)) % 32)
									F_pfree(m, v271)
									mBase = m.M
									v424 = m.ExcPending
									if v424 != 0 {
										return int32(0)
									} else {
										m.G0 = v17 + int32(112)
										return v418
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
		v139 = m.ExcPending
		if v139 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(60895), int32(0))
			mBase = m.M
			v143 = m.ExcPending
			if v143 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(496807), int32(5809), int32(339961))
				mBase = m.M
				v148 = m.ExcPending
				if v148 != 0 {
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
func F_numeric_power(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v268 int32
	_ = v268
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v431 int32
	_ = v431
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v493 int32
	_ = v493
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v567 int32
	_ = v567
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v628 int32
	_ = v628
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v669 int32
	_ = v669
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v786 int32
	_ = v786
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v932 int32
	_ = v932
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v941 int32
	_ = v941
	var v944 int32
	_ = v944
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v957 int32
	_ = v957
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v989 int32
	_ = v989
	var v998 int32
	_ = v998
	var v1001 int32
	_ = v1001
	var v1005 int64
	_ = v1005
	var v1011 int32
	_ = v1011
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int64
	_ = v1139
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1148 float64
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1163 float64
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1170 float64
	_ = v1170
	var v1175 int32
	_ = v1175
	var v1179 float64
	_ = v1179
	var v1187 float64
	_ = v1187
	var v1188 int32
	_ = v1188
	var v1199 int64
	_ = v1199
	var v1214 int32
	_ = v1214
	var v1216 int64
	_ = v1216
	var v1226 int64
	_ = v1226
	var v1230 int64
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1239 float64
	_ = v1239
	var v1241 float64
	_ = v1241
	var v1254 float64
	_ = v1254
	var v1257 float64
	_ = v1257
	var v1262 float64
	_ = v1262
	var v1263 float64
	_ = v1263
	var v1264 float64
	_ = v1264
	var v1265 float64
	_ = v1265
	var v1270 float64
	_ = v1270
	var v1271 float64
	_ = v1271
	var v1272 float64
	_ = v1272
	var v1297 float64
	_ = v1297
	var v1309 float64
	_ = v1309
	var v1331 float64
	_ = v1331
	var v1335 float64
	_ = v1335
	var v1340 float64
	_ = v1340
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1350 int64
	_ = v1350
	var v1359 int32
	_ = v1359
	var v1361 int32
	_ = v1361
	var v1362 int32
	_ = v1362
	var v1364 int32
	_ = v1364
	var v1366 int32
	_ = v1366
	var v1369 int32
	_ = v1369
	var v1373 int32
	_ = v1373
	var v1374 int32
	_ = v1374
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1389 int64
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1397 int32
	_ = v1397
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1407 int32
	_ = v1407
	var v1408 int64
	_ = v1408
	var v1410 int64
	_ = v1410
	var v1413 int32
	_ = v1413
	var v1422 int32
	_ = v1422
	var v1425 int32
	_ = v1425
	var v1434 int32
	_ = v1434
	var v1436 int32
	_ = v1436
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1451 int32
	_ = v1451
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1477 int32
	_ = v1477
	var v1481 int32
	_ = v1481
	var v1487 int32
	_ = v1487
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1515 int32
	_ = v1515
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1523 int32
	_ = v1523
	var v1544 int32
	_ = v1544
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1553 int32
	_ = v1553
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1564 int64
	_ = v1564
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int64
	_ = v1580
	var v1584 int64
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1589 int32
	_ = v1589
	var v1594 float64
	_ = v1594
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1614 int64
	_ = v1614
	var v1616 int64
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1629 int64
	_ = v1629
	var v1632 int64
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1643 int32
	_ = v1643
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1668 int32
	_ = v1668
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1685 int32
	_ = v1685
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1707 int32
	_ = v1707
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1718 int64
	_ = v1718
	var v1722 int32
	_ = v1722
	var v1726 int32
	_ = v1726
	var v1727 int32
	_ = v1727
	var v1729 int32
	_ = v1729
	var v1736 int32
	_ = v1736
	var v1741 int32
	_ = v1741
	var v1742 int32
	_ = v1742
	var v1743 int32
	_ = v1743
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1754 int64
	_ = v1754
	var v1760 int64
	_ = v1760
	var v1780 int32
	_ = v1780
	var v1788 int32
	_ = v1788
	var v1794 int32
	_ = v1794
	var v1798 int32
	_ = v1798
	var v1800 int32
	_ = v1800
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1813 int64
	_ = v1813
	var v1819 int64
	_ = v1819
	var v1823 int32
	_ = v1823
	var v1826 int32
	_ = v1826
	var v1827 int32
	_ = v1827
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1839 int32
	_ = v1839
	var v1847 int32
	_ = v1847
	var v1850 float64
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1857 int32
	_ = v1857
	var v1859 int32
	_ = v1859
	var v1862 int64
	_ = v1862
	var v1869 float64
	_ = v1869
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1913 int32
	_ = v1913
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1928 int32
	_ = v1928
	var v1932 int32
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1944 int32
	_ = v1944
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1955 int32
	_ = v1955
	var v1960 int32
	_ = v1960
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1976 int32
	_ = v1976
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2008 int32
	_ = v2008
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2019 int32
	_ = v2019
	var v2024 int32
	_ = v2024
	var v2028 int32
	_ = v2028
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2040 int32
	_ = v2040
	var v2044 int32
	_ = v2044
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2056 int32
	_ = v2056
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2085 int32
	_ = v2085
	var v2086 int32
	_ = v2086
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2103 int32
	_ = v2103
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2117 int32
	_ = v2117
	var v2122 int32
	_ = v2122
	var v2126 int32
	_ = v2126
	var v2132 int32
	_ = v2132
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2142 int32
	_ = v2142
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2152 int32
	_ = v2152
	var v2160 int32
	_ = v2160
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2168 int32
	_ = v2168
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2209 int32
	_ = v2209
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2219 int32
	_ = v2219
	v17 = m.G0
	v19 = v17 - int32(144)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
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
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v27 = F_pg_detoast_datum(m, v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
	v30 = base.I32_extend16_s(v29)
	if base.Ui32(v29) <= base.Ui32(int32(49151)) {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	m.G0 = v19 + int32(144)
	return v2219
L5:
	;
	v2215 = F_make_result_opt_error(m, int32(1721324), int32(0))
	mBase = m.M
	v2216 = m.ExcPending
	if v2216 != 0 {
		goto L1
	} else {
		goto L703
	}
L6:
	;
	v2203 = F_make_result_opt_error(m, v19, int32(0))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L1
	} else {
		goto L700
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v1369
	v2067 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v2070 = v1369 + v2067<<(uint(int32(2))%32)
	if v2070+int32(4) < int32(0) {
		goto L674
	} else {
		goto L675
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2044 = m.ExcPending
	if v2044 != 0 {
		goto L1
	} else {
		goto L669
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2028 = m.ExcPending
	if v2028 != 0 {
		goto L1
	} else {
		goto L665
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2012 = m.ExcPending
	if v2012 != 0 {
		goto L1
	} else {
		goto L661
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L1
	} else {
		goto L657
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1980 = m.ExcPending
	if v1980 != 0 {
		goto L1
	} else {
		goto L653
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1964 = m.ExcPending
	if v1964 != 0 {
		goto L1
	} else {
		goto L649
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L1
	} else {
		goto L645
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L1
	} else {
		goto L641
	}
L16:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if int32(0) <= base.I32_extend16_s(v33) {
		goto L344
	} else {
		goto L345
	}
L17:
	;
	if v30 == int32(-12288) {
		goto L98
	} else {
		goto L99
	}
L18:
	;
	if v56&int32(65535) == int32(49152) {
		goto L5
	} else {
		goto L96
	}
L19:
	;
	if base.Ui32(v46) < base.Ui32(int32(2)) {
		goto L5
	} else {
		goto L39
	}
L20:
	;
	v83 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+6)))
	v84 = v83
	goto L19
L21:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+4)))
	if base.Ui32(v33) <= base.Ui32(int32(49151)) {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+4)))
	if v30 != int32(-16384) {
		goto L18
	} else {
		goto L30
	}
L24:
	;
	if v33 != int32(49152) {
		v287 = v33
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v44 = base.B2i32(int32(0) <= v30)
	if int32(0) <= v30 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v45 = int32(-8)
	goto L28
L27:
	;
	v45 = int32(-6)
	goto L28
L28:
	;
	v46 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) + v45
	if int32(0) <= v30 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v84 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
	goto L19
L30:
	;
	if base.Ui32(int32(49151)) < base.Ui32(v56&int32(65535)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v81 = F_make_result_opt_error(m, int32(1721324), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L38
	}
L32:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if int32(0) <= base.I32_extend16_s(v56) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v71 = int32(-8)
	goto L35
L34:
	;
	v71 = int32(-6)
	goto L35
L35:
	;
	if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v63)>>(uint(int32(2))%32))+v71) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v77 = F_make_result_opt_error(m, int32(1721448), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v2219 = v77
	goto L4
L38:
	;
	v2219 = v81
	goto L4
L39:
	;
	v92 = v29 & int32(49152)
	if v92 == int32(32768) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v95 = v29 << (uint(int32(1)) % 32) & int32(16384)
	goto L42
L41:
	;
	v95 = v92
	goto L42
L42:
	;
	if v95 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	if v30 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v100 = int32(6)
	goto L46
L45:
	;
	v100 = int32(8)
	goto L46
L46:
	;
	v101 = v22 + v100
	v102 = int32(1)
	v103 = int32(base.Ui32(v46) >> (uint(v102) % 32))
	v106 = int32(0)
	if base.B2i32(v106 < v84)&base.B2i32(v106 < v103) == v106 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	if v278 != 0 {
		goto L5
	} else {
		goto L94
	}
L48:
	;
	v278 = v268
	goto L47
L49:
	;
	if v106 <= v138 {
		v173 = v106
		v175 = v106
		goto L58
	} else {
		goto L59
	}
L50:
	;
	v138 = v84
	v142 = v106
	goto L49
L51:
	;
	goto L52
L52:
	;
	v119 = v84
	v123 = v106
	goto L53
L53:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+v123<<(uint(int32(1))%32)))))
	if v129 != 0 {
		v268 = int32(1)
		goto L48
	} else {
		goto L55
	}
L54:
	;
	v138 = v133
	v142 = v131
	goto L49
L55:
	;
	v130 = int32(1)
	v131 = v123 + v130
	v133 = v119 - v130
	if v133 <= v106 {
		v138 = v133
		v142 = v131
		goto L49
	} else {
		goto L56
	}
L56:
	;
	if v131 < v103 {
		v119 = v133
		v123 = v131
		goto L53
	} else {
		goto L57
	}
L57:
	;
	goto L54
L58:
	;
	if v138 != v173 {
		v214 = v142
		v215 = v175
		goto L66
	} else {
		goto L67
	}
L59:
	;
	goto L60
L60:
	;
	v154 = v106
	v156 = v106
	goto L61
L61:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156<<(uint(int32(1))%32))+uint32(_consts[1089]))))
	if v161 != 0 {
		v268 = int32(-1)
		goto L48
	} else {
		goto L63
	}
L62:
	;
	v173 = v165
	v175 = v163
	goto L58
L63:
	;
	v162 = int32(1)
	v163 = v156 + v162
	v165 = v154 - v162
	if v165 <= v138 {
		v173 = v165
		v175 = v163
		goto L58
	} else {
		goto L64
	}
L64:
	;
	if v163 < v102 {
		v154 = v165
		v156 = v163
		goto L61
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	if v103 < v214 {
		goto L76
	} else {
		goto L77
	}
L67:
	;
	v184 = v142
	v185 = v175
	goto L68
L68:
	;
	if v103 <= v184 {
		v214 = v184
		v215 = v185
		goto L66
	} else {
		goto L70
	}
L69:
	;
	if base.I32_extend16_s(v200) < base.I32_extend16_s(v198) {
		goto L73
	} else {
		goto L74
	}
L70:
	;
	if v102 <= v185 {
		v214 = v184
		v215 = v185
		goto L66
	} else {
		goto L71
	}
L71:
	;
	v189 = int32(1)
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+v184<<(uint(v189)%32)))))
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185<<(uint(v189)%32))+uint32(_consts[1089]))))
	if v198 == v200 {
		v184 = v184 + v189
		v185 = v185 + v189
		goto L68
	} else {
		goto L72
	}
L72:
	;
	goto L69
L73:
	;
	v207 = int32(1)
	goto L75
L74:
	;
	v207 = int32(-1)
	goto L75
L75:
	;
	v278 = v207
	goto L47
L76:
	;
	v218 = v214
	goto L78
L77:
	;
	v218 = v103
	goto L78
L78:
	;
	v225 = v214
	goto L79
L79:
	;
	if v218 == v225 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v268 = v251
	goto L48
L81:
	;
	if v102 < v215 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	v251 = int32(1)
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101+v225<<(uint(v251)%32)))))
	if v257 == int32(0) {
		v225 = v225 + v251
		goto L79
	} else {
		goto L93
	}
L84:
	;
	v230 = v215
	goto L86
L85:
	;
	v230 = v102
	goto L86
L86:
	;
	v238 = v215
	goto L87
L87:
	;
	if v230 == v238 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v268 = int32(-1)
	goto L48
L89:
	;
	v278 = int32(0)
	goto L47
L90:
	;
	goto L91
L91:
	;
	v242 = int32(1)
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238<<(uint(v242)%32))+uint32(_consts[1089]))))
	if v247 == int32(0) {
		v238 = v238 + v242
		goto L87
	} else {
		goto L92
	}
L92:
	;
	goto L88
L93:
	;
	goto L80
L94:
	;
	v281 = F_make_result_opt_error(m, int32(1721448), int32(0))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	v2219 = v281
	goto L4
L96:
	;
	v287 = v56
	goto L17
L97:
	;
	v323 = v287 & int32(65535)
	v324 = int32(49152)
	v325 = v287 & v324
	if v325 == v324 {
		goto L113
	} else {
		goto L114
	}
L98:
	;
	v292 = int32(1)
	goto L100
L99:
	;
	v292 = int32(-1)
	goto L100
L100:
	;
	v293 = int32(49152)
	v294 = v29 & v293
	if v294 == v293 {
		v321 = v292
		goto L97
	} else {
		goto L101
	}
L101:
	;
	v297 = int32(0)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v297 <= v30 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v305 = int32(-8)
	goto L104
L103:
	;
	v305 = int32(-6)
	goto L104
L104:
	;
	if base.Ui32(int32(base.Ui32(v298)>>(uint(int32(2))%32))+v305) < base.Ui32(int32(2)) {
		v321 = v297
		goto L97
	} else {
		goto L105
	}
L105:
	;
	v310 = int32(1)
	if v294 == int32(32768) {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v317 = v29 << (uint(v310) % 32) & int32(16384)
	goto L108
L107:
	;
	v317 = v294
	goto L108
L108:
	;
	if v317 == int32(16384) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v320 = int32(-1)
	goto L111
L110:
	;
	v320 = v310
	goto L111
L111:
	;
	v321 = v320
	goto L97
L112:
	;
	if base.Ui32(int32(49151)) < base.Ui32(v287&int32(65535)) {
		goto L133
	} else {
		goto L134
	}
L113:
	;
	if v323 == int32(53248) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if int32(0) <= base.I32_extend16_s(v287) {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	v362 = int32(base.Ui32(v360) >> (uint(int32(31)) % 32))
	v364 = base.B2i32(v321 == int32(0))
	if v321 != 0 {
		v367 = v360
		v368 = v362
		v369 = v364
		goto L112
	} else {
		goto L131
	}
L116:
	;
	v332 = int32(1)
	goto L118
L117:
	;
	v332 = int32(-1)
	goto L118
L118:
	;
	v360 = v332
	goto L115
L119:
	;
	v341 = int32(-8)
	goto L121
L120:
	;
	v341 = int32(-6)
	goto L121
L121:
	;
	if base.Ui32(int32(base.Ui32(v333)>>(uint(int32(2))%32))+v341) <= base.Ui32(int32(1)) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v345 = int32(0)
	v367 = v345
	v368 = int32(0)
	v369 = base.B2i32(v321 == v345)
	goto L112
L123:
	;
	goto L124
L124:
	;
	v349 = int32(1)
	if v325 == int32(32768) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v356 = v323 << (uint(v349) % 32) & int32(16384)
	goto L127
L126:
	;
	v356 = v325
	goto L127
L127:
	;
	if v356 == int32(16384) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v359 = int32(-1)
	goto L130
L129:
	;
	v359 = v349
	goto L130
L130:
	;
	v360 = v359
	goto L115
L131:
	;
	if v360 < int32(0) {
		goto L15
	} else {
		goto L132
	}
L132:
	;
	v367 = v360
	v368 = v362
	v369 = v364
	goto L112
L133:
	;
	if base.Ui32(int32(-16385)) < base.Ui32(v30) {
		goto L144
	} else {
		goto L145
	}
L134:
	;
	if int32(0) <= v321 {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v383 = base.B2i32(int32(0) <= base.I32_extend16_s(v287))
	if int32(0) <= base.I32_extend16_s(v287) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v384 = int32(-8)
	goto L138
L137:
	;
	v384 = int32(-6)
	goto L138
L138:
	;
	v385 = int32(base.Ui32(v376)>>(uint(int32(2))%32)) + v384
	if int32(0) <= base.I32_extend16_s(v287) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v386 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+6)))
	v396 = v386
	goto L141
L140:
	;
	v396 = v323<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v323&int32(63)
	goto L141
L141:
	;
	if base.Ui32(v385) < base.Ui32(int32(2)) {
		goto L133
	} else {
		goto L142
	}
L142:
	;
	v399 = int32(1)
	if v396+v399 < int32(base.Ui32(v385)>>(uint(v399)%32)) {
		goto L14
	} else {
		goto L143
	}
L143:
	;
	goto L133
L144:
	;
	if v367 == int32(0) {
		goto L214
	} else {
		goto L215
	}
L145:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v415 = base.B2i32(int32(0) <= v30)
	if int32(0) <= v30 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v416 = int32(-8)
	goto L148
L147:
	;
	v416 = int32(-6)
	goto L148
L148:
	;
	v417 = int32(base.Ui32(v409)>>(uint(int32(2))%32)) + v416
	v419 = int32(base.Ui32(v417) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v419
	if int32(0) <= v30 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v421 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+6)))
	v431 = v421
	goto L151
L150:
	;
	v431 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
	goto L151
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v431
	if v294 != int32(49152) {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v443
	v445 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v445
	v454 = base.B2i32(v30 < v445)
	if v30 < v445 {
		goto L157
	} else {
		goto L158
	}
L153:
	;
	if v294 != int32(32768) {
		v443 = v294
		goto L152
	} else {
		goto L156
	}
L154:
	;
	goto L155
L155:
	;
	v443 = v29 & int32(61440)
	goto L152
L156:
	;
	v443 = v29 << (uint(int32(1)) % 32) & int32(16384)
	goto L152
L157:
	;
	v455 = int32(base.Ui32(v29)>>(uint(int32(7))%32)) & int32(63)
	goto L159
L158:
	;
	v455 = v29 & int32(16383)
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v455
	if v30 < v445 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v459 = int32(6)
	goto L162
L161:
	;
	v459 = int32(8)
	goto L162
L162:
	;
	v460 = v22 + v459
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v460
	if base.Ui32(v417) < base.Ui32(int32(2)) {
		goto L144
	} else {
		goto L163
	}
L163:
	;
	if v443 != 0 {
		goto L144
	} else {
		goto L164
	}
L164:
	;
	v465 = int32(1)
	v466 = int32(0)
	if base.B2i32(v466 < v431)&base.B2i32(v466 < v419) == v466 {
		goto L168
	} else {
		goto L169
	}
L165:
	;
	if v638 != 0 {
		goto L144
	} else {
		goto L212
	}
L166:
	;
	v638 = v628
	goto L165
L167:
	;
	if v466 <= v498 {
		v533 = v466
		v535 = v466
		goto L176
	} else {
		goto L177
	}
L168:
	;
	v498 = v431
	v502 = v466
	goto L167
L169:
	;
	goto L170
L170:
	;
	v479 = v431
	v483 = v466
	goto L171
L171:
	;
	v489 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v460+v483<<(uint(int32(1))%32)))))
	if v489 != 0 {
		v628 = int32(1)
		goto L166
	} else {
		goto L173
	}
L172:
	;
	v498 = v493
	v502 = v491
	goto L167
L173:
	;
	v490 = int32(1)
	v491 = v483 + v490
	v493 = v479 - v490
	if v493 <= v466 {
		v498 = v493
		v502 = v491
		goto L167
	} else {
		goto L174
	}
L174:
	;
	if v491 < v419 {
		v479 = v493
		v483 = v491
		goto L171
	} else {
		goto L175
	}
L175:
	;
	goto L172
L176:
	;
	if v498 != v533 {
		v574 = v502
		v575 = v535
		goto L184
	} else {
		goto L185
	}
L177:
	;
	goto L178
L178:
	;
	v514 = v466
	v516 = v466
	goto L179
L179:
	;
	v521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v516<<(uint(int32(1))%32))+uint32(_consts[1089]))))
	if v521 != 0 {
		v628 = int32(-1)
		goto L166
	} else {
		goto L181
	}
L180:
	;
	v533 = v525
	v535 = v523
	goto L176
L181:
	;
	v522 = int32(1)
	v523 = v516 + v522
	v525 = v514 - v522
	if v525 <= v498 {
		v533 = v525
		v535 = v523
		goto L176
	} else {
		goto L182
	}
L182:
	;
	if v523 < v465 {
		v514 = v525
		v516 = v523
		goto L179
	} else {
		goto L183
	}
L183:
	;
	goto L180
L184:
	;
	if v419 < v574 {
		goto L194
	} else {
		goto L195
	}
L185:
	;
	v544 = v502
	v545 = v535
	goto L186
L186:
	;
	if v419 <= v544 {
		v574 = v544
		v575 = v545
		goto L184
	} else {
		goto L188
	}
L187:
	;
	if base.I32_extend16_s(v560) < base.I32_extend16_s(v558) {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	if v465 <= v545 {
		v574 = v544
		v575 = v545
		goto L184
	} else {
		goto L189
	}
L189:
	;
	v549 = int32(1)
	v558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v460+v544<<(uint(v549)%32)))))
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v545<<(uint(v549)%32))+uint32(_consts[1089]))))
	if v558 == v560 {
		v544 = v544 + v549
		v545 = v545 + v549
		goto L186
	} else {
		goto L190
	}
L190:
	;
	goto L187
L191:
	;
	v567 = int32(1)
	goto L193
L192:
	;
	v567 = int32(-1)
	goto L193
L193:
	;
	v638 = v567
	goto L165
L194:
	;
	v578 = v574
	goto L196
L195:
	;
	v578 = v419
	goto L196
L196:
	;
	v585 = v574
	goto L197
L197:
	;
	if v578 == v585 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v628 = v611
	goto L166
L199:
	;
	if v465 < v575 {
		goto L202
	} else {
		goto L203
	}
L200:
	;
	goto L201
L201:
	;
	v611 = int32(1)
	v617 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v460+v585<<(uint(v611)%32)))))
	if v617 == int32(0) {
		v585 = v585 + v611
		goto L197
	} else {
		goto L211
	}
L202:
	;
	v590 = v575
	goto L204
L203:
	;
	v590 = v465
	goto L204
L204:
	;
	v598 = v575
	goto L205
L205:
	;
	if v590 == v598 {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	v628 = int32(-1)
	goto L166
L207:
	;
	v638 = int32(0)
	goto L165
L208:
	;
	goto L209
L209:
	;
	v602 = int32(1)
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v598<<(uint(v602)%32))+uint32(_consts[1089]))))
	if v607 == int32(0) {
		v598 = v598 + v602
		goto L205
	} else {
		goto L210
	}
L210:
	;
	goto L206
L211:
	;
	goto L198
L212:
	;
	v641 = F_make_result_opt_error(m, int32(1721448), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v2219 = v641
	goto L4
L214:
	;
	v652 = F_make_result_opt_error(m, int32(1721448), int32(0))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	if base.B2i32(int32(0) < v367)&v369 != 0 {
		goto L218
	} else {
		goto L219
	}
L217:
	;
	v2219 = v652
	goto L4
L218:
	;
	v659 = F_make_result_opt_error(m, int32(1721396), int32(0))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L1
	} else {
		goto L221
	}
L219:
	;
	goto L220
L220:
	;
	if v287&int32(57343) == int32(53248) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v2219 = v659
	goto L4
L222:
	;
	if base.Ui32(v30) <= base.Ui32(int32(-16385)) {
		goto L225
	} else {
		goto L226
	}
L223:
	;
	goto L224
L224:
	;
	if v30 == int32(-12288) {
		goto L306
	} else {
		goto L307
	}
L225:
	;
	v669 = v19 + int32(48)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v677 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
	if int32(0) <= v677 {
		goto L229
	} else {
		goto L230
	}
L226:
	;
	v845 = int32(1)
	goto L227
L227:
	;
	if v845 == base.B2i32(int32(0) < v367) {
		goto L301
	} else {
		goto L302
	}
L228:
	;
	v740 = v19 + int32(48)
	v741 = int32(1721420)
	v748 = *(*int32)(unsafe.Add(mBase, _consts[1090]))
	v749 = *(*int32)(unsafe.Add(mBase, _consts[1091]))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v740)))
	if v750 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L229:
	;
	v680 = int32(-8)
	goto L231
L230:
	;
	v680 = int32(-6)
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669))) = int32(base.Ui32(int32(base.Ui32(v672)>>(uint(int32(2))%32))+v680) >> (uint(int32(1)) % 32))
	v685 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
	if v685 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+4)) = v701
	v703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
	v704 = int32(49152)
	v705 = v703 & v704
	if v705 != v704 {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	v689 = v685 & int32(65535)
	v701 = v689<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v689&int32(63)
	goto L232
L234:
	;
	goto L235
L235:
	;
	v699 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+6)))
	v701 = v699
	goto L232
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+8)) = v716
	v718 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
	if v718 < int32(0) {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	if v705 != int32(32768) {
		v716 = v705
		goto L236
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v716 = v703 & int32(61440)
	goto L236
L240:
	;
	v716 = v703 << (uint(int32(1)) % 32) & int32(16384)
	goto L236
L241:
	;
	v727 = int32(base.Ui32(v718)>>(uint(int32(7))%32)) & int32(63)
	goto L243
L242:
	;
	v727 = v718 & int32(16383)
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+12)) = v727
	v729 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+4)))
	v730 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v669)+16)) = v730
	if v729 < v730 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v736 = int32(6)
	goto L246
L245:
	;
	v736 = int32(8)
	goto L246
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v669)+20)) = v22 + v736
	goto L228
L247:
	;
	if v786 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L248:
	;
	if v749 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L249:
	;
	goto L250
L250:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v740)+8))
	if v749 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L251:
	;
	v786 = int32(0)
	goto L247
L252:
	;
	goto L253
L253:
	;
	if v748 == int32(16384) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v760 = int32(1)
	goto L256
L255:
	;
	v760 = int32(-1)
	goto L256
L256:
	;
	v786 = v760
	goto L247
L257:
	;
	if v761 != 0 {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L259
L259:
	;
	v767 = *(*int32)(unsafe.Add(mBase, _consts[1092]))
	v768 = *(*int32)(unsafe.Add(mBase, _consts[1093]))
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v740)+4))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v740)+20))
	if v761 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v766 = int32(-1)
	goto L262
L261:
	;
	v766 = int32(1)
	goto L262
L262:
	;
	v786 = v766
	goto L247
L263:
	;
	if v748 == int32(16384) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	goto L265
L265:
	;
	if v748 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L266:
	;
	v786 = int32(1)
	goto L247
L267:
	;
	goto L268
L268:
	;
	v776 = F_cmp_abs_common(m, v770, v750, v769, v768, v749, v767)
	mBase = m.M
	v786 = v776
	goto L247
L269:
	;
	v786 = int32(-1)
	goto L247
L270:
	;
	goto L271
L271:
	;
	v780 = F_cmp_abs_common(m, v768, v749, v767, v770, v750, v769)
	mBase = m.M
	v786 = v780
	goto L247
L272:
	;
	v791 = F_make_result_opt_error(m, int32(1721448), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v793 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v793
	v796 = v19 + int32(48)
	v797 = int32(1721448)
	v804 = *(*int32)(unsafe.Add(mBase, _consts[1083]))
	v805 = *(*int32)(unsafe.Add(mBase, _consts[1084]))
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v796)))
	if v806 == v793 {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	v2219 = v791
	goto L4
L276:
	;
	v845 = base.B2i32(int32(0) < v842)
	goto L227
L277:
	;
	if v805 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L279
L279:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v796)+8))
	if v805 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L280:
	;
	v842 = int32(0)
	goto L276
L281:
	;
	goto L282
L282:
	;
	if v804 == int32(16384) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v816 = int32(1)
	goto L285
L284:
	;
	v816 = int32(-1)
	goto L285
L285:
	;
	v842 = v816
	goto L276
L286:
	;
	if v817 != 0 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	goto L288
L288:
	;
	v823 = *(*int32)(unsafe.Add(mBase, _consts[1094]))
	v824 = *(*int32)(unsafe.Add(mBase, _consts[1095]))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v796)+4))
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v796)+20))
	if v817 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	v822 = int32(-1)
	goto L291
L290:
	;
	v822 = int32(1)
	goto L291
L291:
	;
	v842 = v822
	goto L276
L292:
	;
	if v804 == int32(16384) {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	goto L294
L294:
	;
	if v804 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	v842 = int32(1)
	goto L276
L296:
	;
	goto L297
L297:
	;
	v832 = F_cmp_abs_common(m, v826, v806, v825, v824, v805, v823)
	mBase = m.M
	v842 = v832
	goto L276
L298:
	;
	v842 = int32(-1)
	goto L276
L299:
	;
	goto L300
L300:
	;
	v836 = F_cmp_abs_common(m, v824, v805, v823, v826, v806, v825)
	mBase = m.M
	v842 = v836
	goto L276
L301:
	;
	v851 = F_make_result_opt_error(m, int32(1721348), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	v855 = F_make_result_opt_error(m, int32(1721396), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L305
	}
L304:
	;
	v2219 = v851
	goto L4
L305:
	;
	v2219 = v855
	goto L4
L306:
	;
	if int32(0) < v367 {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L308
L308:
	;
	if v368 != 0 {
		goto L314
	} else {
		goto L315
	}
L309:
	;
	v863 = F_make_result_opt_error(m, int32(1721348), int32(0))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v867 = F_make_result_opt_error(m, int32(1721396), int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L1
	} else {
		goto L313
	}
L312:
	;
	v2219 = v863
	goto L4
L313:
	;
	v2219 = v867
	goto L4
L314:
	;
	v871 = F_make_result_opt_error(m, int32(1721396), int32(0))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L1
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v874 = v19 + int32(24)
	v877 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v882 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+4)))
	if int32(0) <= v882 {
		goto L319
	} else {
		goto L320
	}
L317:
	;
	v2219 = v871
	goto L4
L318:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v944 <= int32(0) {
		goto L337
	} else {
		goto L338
	}
L319:
	;
	v885 = int32(-8)
	goto L321
L320:
	;
	v885 = int32(-6)
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v874))) = int32(base.Ui32(int32(base.Ui32(v877)>>(uint(int32(2))%32))+v885) >> (uint(int32(1)) % 32))
	v890 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+4)))
	if v890 < int32(0) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v874)+4)) = v906
	v908 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+4)))
	v909 = int32(49152)
	v910 = v908 & v909
	if v910 != v909 {
		goto L327
	} else {
		goto L328
	}
L323:
	;
	v894 = v890 & int32(65535)
	v906 = v894<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v894&int32(63)
	goto L322
L324:
	;
	goto L325
L325:
	;
	v904 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+6)))
	v906 = v904
	goto L322
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v874)+8)) = v921
	v923 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+4)))
	if v923 < int32(0) {
		goto L331
	} else {
		goto L332
	}
L327:
	;
	if v910 != int32(32768) {
		v921 = v910
		goto L326
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v921 = v908 & int32(61440)
	goto L326
L330:
	;
	v921 = v908 << (uint(int32(1)) % 32) & int32(16384)
	goto L326
L331:
	;
	v932 = int32(base.Ui32(v923)>>(uint(int32(7))%32)) & int32(63)
	goto L333
L332:
	;
	v932 = v923 & int32(16383)
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v874)+12)) = v932
	v934 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+4)))
	v935 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v874)+16)) = v935
	if v934 < v935 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v941 = int32(6)
	goto L336
L335:
	;
	v941 = int32(8)
	goto L336
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v874)+20)) = v27 + v941
	goto L318
L337:
	;
	v968 = F_make_result_opt_error(m, int32(1721348), int32(0))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L342
	}
L338:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	if v944 != v947+int32(1) {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	v952 = int32(1)
	v957 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951+v944<<(uint(v952)%32)-int32(2)))))
	if v957&v952 == int32(0) {
		goto L337
	} else {
		goto L340
	}
L340:
	;
	v964 = F_make_result_opt_error(m, int32(1721372), int32(0))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v2219 = v964
	goto L4
L342:
	;
	v2219 = v968
	goto L4
L343:
	;
	v1005 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1005
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v1005
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1005
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v1016 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+4)))
	v1017 = base.I32_extend16_s(v1016)
	v1019 = base.B2i32(int32(0) <= v1017)
	if int32(0) <= v1017 {
		goto L356
	} else {
		goto L357
	}
L344:
	;
	v978 = int32(-8)
	goto L346
L345:
	;
	v978 = int32(-6)
	goto L346
L346:
	;
	if base.Ui32(int32(base.Ui32(v970)>>(uint(int32(2))%32))+v978) < base.Ui32(int32(2)) {
		goto L343
	} else {
		goto L347
	}
L347:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if int32(0) <= v30 {
		goto L348
	} else {
		goto L349
	}
L348:
	;
	v989 = int32(-8)
	goto L350
L349:
	;
	v989 = int32(-6)
	goto L350
L350:
	;
	if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v982)>>(uint(int32(2))%32))+v989) {
		goto L343
	} else {
		goto L351
	}
L351:
	;
	v998 = v33 & int32(49152)
	if v998 == int32(32768) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1001 = v33 << (uint(int32(1)) % 32) & int32(16384)
	goto L354
L353:
	;
	v1001 = v998
	goto L354
L354:
	;
	if v1001 == int32(16384) {
		goto L13
	} else {
		goto L355
	}
L355:
	;
	goto L343
L356:
	;
	v1020 = int32(-8)
	goto L358
L357:
	;
	v1020 = int32(-6)
	goto L358
L358:
	;
	v1021 = int32(base.Ui32(v1011)>>(uint(int32(2))%32)) + v1020
	v1023 = int32(base.Ui32(v1021) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = v1023
	if int32(0) <= v1017 {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1025 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+6)))
	v1035 = v1025
	goto L361
L360:
	;
	v1035 = v1016<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v1016&int32(63)
	goto L361
L361:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+52)) = v1035
	v1037 = int32(49152)
	v1038 = v1016 & v1037
	if v1038 != v1037 {
		goto L363
	} else {
		goto L364
	}
L362:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v1049
	v1051 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v1051
	v1060 = base.B2i32(v1017 < v1051)
	if v1017 < v1051 {
		goto L367
	} else {
		goto L368
	}
L363:
	;
	if v1038 != int32(32768) {
		v1049 = v1038
		goto L362
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1049 = v1016 & int32(61440)
	goto L362
L366:
	;
	v1049 = v1016 << (uint(int32(1)) % 32) & int32(16384)
	goto L362
L367:
	;
	v1061 = int32(base.Ui32(v1016)>>(uint(int32(7))%32)) & int32(63)
	goto L369
L368:
	;
	v1061 = v1016 & int32(16383)
	goto L369
L369:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+60)) = v1061
	if v1017 < v1051 {
		goto L370
	} else {
		goto L371
	}
L370:
	;
	v1065 = int32(6)
	goto L372
L371:
	;
	v1065 = int32(8)
	goto L372
L372:
	;
	v1066 = v22 + v1065
	*(*int32)(unsafe.Add(mBase, uint32(v19)+68)) = v1066
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v1073 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+4)))
	v1074 = base.I32_extend16_s(v1073)
	v1076 = base.B2i32(int32(0) <= v1074)
	if int32(0) <= v1074 {
		goto L373
	} else {
		goto L374
	}
L373:
	;
	v1077 = int32(-8)
	goto L375
L374:
	;
	v1077 = int32(-6)
	goto L375
L375:
	;
	v1078 = int32(base.Ui32(v1068)>>(uint(int32(2))%32)) + v1077
	v1080 = int32(base.Ui32(v1078) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v1080
	if int32(0) <= v1074 {
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v1082 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+6)))
	v1092 = v1082
	goto L378
L377:
	;
	v1092 = v1073<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v1073&int32(63)
	goto L378
L378:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v1092
	v1094 = int32(49152)
	v1095 = v1073 & v1094
	if v1095 != v1094 {
		goto L380
	} else {
		goto L381
	}
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v1106
	v1108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v1108
	v1117 = base.B2i32(v1074 < v1108)
	if v1074 < v1108 {
		goto L384
	} else {
		goto L385
	}
L380:
	;
	if v1095 != int32(32768) {
		v1106 = v1095
		goto L379
	} else {
		goto L383
	}
L381:
	;
	goto L382
L382:
	;
	v1106 = v1073 & int32(61440)
	goto L379
L383:
	;
	v1106 = v1073 << (uint(int32(1)) % 32) & int32(16384)
	goto L379
L384:
	;
	v1118 = int32(base.Ui32(v1073)>>(uint(int32(7))%32)) & int32(63)
	goto L386
L385:
	;
	v1118 = v1073 & int32(16383)
	goto L386
L386:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = v1118
	if v1074 < v1108 {
		goto L387
	} else {
		goto L388
	}
L387:
	;
	v1122 = int32(6)
	goto L389
L388:
	;
	v1122 = int32(8)
	goto L389
L389:
	;
	v1123 = v27 + v1122
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v1123
	v1126 = v1092 + int32(1)
	if base.B2i32(v1126 < v1080)&base.B2i32(base.Ui32(int32(2)) <= base.Ui32(v1078)) != 0 {
		goto L390
	} else {
		goto L391
	}
L390:
	;
	if base.Ui32(v1021) <= base.Ui32(int32(1)) {
		goto L568
	} else {
		goto L569
	}
L391:
	;
	v1135 = F_numericvar_to_int64(m, v19+int32(24), v19+int32(96))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L392
	}
L392:
	;
	if v1135 == int32(0) {
		goto L390
	} else {
		goto L393
	}
L393:
	;
	v1139 = *(*int64)(unsafe.Add(mBase, uint32(v19)+96))
	if base.Ui64(int64(4294967295)) < base.Ui64(v1139+int64(2147483648)) {
		goto L390
	} else {
		goto L394
	}
L394:
	;
	v1144 = base.I32_wrap_i64(v1139)
	if base.Ui32(int32(2)) <= base.Ui32(v1021) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1147 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1066))))
	v1148 = base.F64_convert_i32_s(v1147)
	if base.Ui32(v1021) < base.Ui32(int32(4)) {
		goto L398
	} else {
		goto L399
	}
L396:
	;
	v1340 = float64(0)
	goto L397
L397:
	;
	if base.F64_lt(base.F64_add(v1340, float64(1)), float64(-1000)) != 0 {
		goto L424
	} else {
		goto L425
	}
L398:
	;
	v1187 = v1148
	v1188 = v1035 << (uint(int32(2)) % 32)
	goto L400
L399:
	;
	v1153 = int32(2)
	v1155 = v1023 - v1153
	if base.Ui32(v1153) <= base.Ui32(v1155) {
		goto L401
	} else {
		goto L402
	}
L400:
	;
	v1199 = base.I64_reinterpret_f64(v1187)
	if v1199 <= int64(4503599627370495) {
		goto L411
	} else {
		goto L412
	}
L401:
	;
	v1158 = v1153
	goto L403
L402:
	;
	v1158 = v1155
	goto L403
L403:
	;
	v1161 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1066)+2)))
	v1163 = base.F64_add(base.F64_mul(v1148, float64(10000)), base.F64_convert_i32_s(v1161))
	if v1155 == int32(0) {
		v1179 = v1163
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1187 = v1179
	v1188 = (v1035-v1158)<<(uint(int32(2))%32) - int32(4)
	goto L400
L405:
	;
	v1168 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1066)+4)))
	v1170 = base.F64_add(base.F64_mul(v1163, float64(10000)), base.F64_convert_i32_s(v1168))
	if v1155 == int32(1) {
		v1179 = v1170
		goto L404
	} else {
		goto L406
	}
L406:
	;
	v1175 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1066)+6)))
	v1179 = base.F64_add(base.F64_mul(v1170, float64(10000)), base.F64_convert_i32_s(v1175))
	goto L404
L407:
	;
	v1335 = base.F64_mul(base.F64_add(v1331, base.F64_convert_i32_s(v1188)), base.F64_convert_i32_s(v1144))
	if base.F64_gt(v1335, float64(131072)) != 0 {
		goto L12
	} else {
		goto L423
	}
L408:
	;
	v1331 = v1309
	goto L407
L409:
	;
	v1235 = v1233 + int32(614242)
	v1239 = base.F64_convert_i32_s(int32(base.Ui32(v1235)>>(uint(int32(20))%32)) + v1232)
	v1241 = base.F64_mul(v1239, float64(0.30102999566361177))
	v1254 = base.F64_add(base.F64_reinterpret_i64(v1230&int64(4294967295)|base.I64_extend_i32_u(v1235&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v1257 = base.F64_mul(v1254, base.F64_mul(v1254, float64(0.5)))
	v1262 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v1254, v1257)) & int64(-4294967296))
	v1263 = float64(0.4342944818781689)
	v1264 = base.F64_mul(v1262, v1263)
	v1265 = base.F64_add(v1241, v1264)
	v1270 = base.F64_div(v1254, base.F64_add(v1254, float64(2)))
	v1271 = base.F64_mul(v1270, v1270)
	v1272 = base.F64_mul(v1271, v1271)
	v1297 = base.F64_add(base.F64_mul(v1270, base.F64_add(v1257, base.F64_add(base.F64_mul(v1272, base.F64_add(base.F64_mul(v1272, base.F64_add(base.F64_mul(v1272, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v1271, base.F64_add(base.F64_mul(v1272, base.F64_add(base.F64_mul(v1272, base.F64_add(base.F64_mul(v1272, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v1254, v1262), v1257))
	v1309 = base.F64_add(v1265, base.F64_add(base.F64_add(v1264, base.F64_sub(v1241, v1265)), base.F64_add(base.F64_mul(v1297, v1263), base.F64_add(base.F64_mul(v1239, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v1297, v1262), float64(2.5082946711645275e-11))))))
	goto L408
L410:
	;
	v1226 = base.I64_reinterpret_f64(base.F64_mul(v1187, float64(1.8014398509481984e+16)))
	v1230 = v1226
	v1232 = int32(-1077)
	v1233 = base.I32_wrap_i64(int64(base.Ui64(v1226) >> (uint(int64(32)) % 64)))
	goto L409
L411:
	;
	if base.F64_eq(v1187, float64(0)) != 0 {
		goto L414
	} else {
		goto L415
	}
L412:
	;
	goto L413
L413:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v1199) {
		v1309 = v1187
		goto L408
	} else {
		goto L418
	}
L414:
	;
	v1331 = base.F64_div(float64(-1), base.F64_mul(v1187, v1187))
	goto L407
L415:
	;
	goto L416
L416:
	;
	if int64(0) <= v1199 {
		goto L410
	} else {
		goto L417
	}
L417:
	;
	v1331 = base.F64_div(base.F64_sub(v1187, v1187), float64(0))
	goto L407
L418:
	;
	v1214 = int32(-1023)
	v1216 = int64(base.Ui64(v1199) >> (uint(int64(32)) % 64))
	if v1216 != int64(1072693248) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1230 = v1199
	v1232 = v1214
	v1233 = base.I32_wrap_i64(v1216)
	goto L409
L420:
	;
	goto L421
L421:
	;
	if base.I32_wrap_i64(v1199) != 0 {
		v1230 = v1199
		v1232 = v1214
		v1233 = int32(1072693248)
		goto L409
	} else {
		goto L422
	}
L422:
	;
	v1331 = float64(0)
	goto L407
L423:
	;
	v1340 = v1335
	goto L397
L424:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1345 != 0 {
		goto L427
	} else {
		goto L428
	}
L425:
	;
	goto L426
L426:
	;
	if base.F64_lt(base.F64_abs(v1340), float64(2.147483648e+09)) != 0 {
		goto L432
	} else {
		goto L433
	}
L427:
	;
	F_pfree(m, v1345)
	mBase = m.M
	v1347 = m.ExcPending
	if v1347 != 0 {
		goto L1
	} else {
		goto L430
	}
L428:
	;
	goto L429
L429:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(4294967296000)
	v1350 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1350
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1350
	goto L6
L430:
	;
	goto L429
L431:
	;
	v1362 = int32(16) - v1361
	if v1061 < v1362 {
		goto L435
	} else {
		goto L436
	}
L432:
	;
	v1359 = base.I32_trunc_f64_s(v1340)
	v1361 = v1359
	goto L431
L433:
	;
	goto L434
L434:
	;
	v1361 = int32(-2147483648)
	goto L431
L435:
	;
	v1364 = v1362
	goto L437
L436:
	;
	v1364 = v1061
	goto L437
L437:
	;
	if v1118 < v1364 {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v1366 = v1364
	goto L440
L439:
	;
	v1366 = v1118
	goto L440
L440:
	;
	if int32(1000) <= v1366 {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v1369 = int32(1000)
	goto L443
L442:
	;
	v1369 = v1366
	goto L443
L443:
	;
	switch v1144 + int32(1) {
	case 0:
		goto L446
	case 1:
		goto L448
	case 2:
		goto L447
	case 3:
		goto L445
	default:
		goto L444
	}
L444:
	;
	if base.Ui32(v1021) <= base.Ui32(int32(1)) {
		goto L495
	} else {
		goto L496
	}
L445:
	;
	v1549 = v19 + int32(48)
	F_mul_var(m, v1549, v1549, v19, v1369)
	mBase = m.M
	v1553 = m.ExcPending
	if v1553 != 0 {
		goto L1
	} else {
		goto L494
	}
L446:
	;
	v1544 = int32(1)
	F_div_var(m, int32(1721448), v19+int32(48), v19, v1369, v1544, v1544)
	mBase = m.M
	v1547 = m.ExcPending
	if v1547 != 0 {
		goto L1
	} else {
		goto L493
	}
L447:
	;
	v1392 = v1021 & int32(-2)
	v1395 = F_palloc(m, v1392+int32(2))
	mBase = m.M
	v1396 = m.ExcPending
	if v1396 != 0 {
		goto L1
	} else {
		goto L454
	}
L448:
	;
	v1373 = F_palloc(m, int32(4))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L1
	} else {
		goto L449
	}
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1373))) = int32(65536)
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1379 != 0 {
		goto L450
	} else {
		goto L451
	}
L450:
	;
	F_pfree(m, v1379)
	mBase = m.M
	v1381 = m.ExcPending
	if v1381 != 0 {
		goto L1
	} else {
		goto L453
	}
L451:
	;
	goto L452
L452:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, _consts[1083]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v1383
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v1373 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1373
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v1369
	v1389 = *(*int64)(unsafe.Add(mBase, _consts[1084]))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1389
	goto L6
L453:
	;
	goto L452
L454:
	;
	v1397 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1395))) = uint16(v1397)
	if base.Ui32(int32(2)) <= base.Ui32(v1021) {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	if v1392 != 0 {
		goto L459
	} else {
		goto L460
	}
L456:
	;
	goto L457
L457:
	;
	v1405 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1405 != 0 {
		goto L462
	} else {
		goto L463
	}
L458:
	;
	goto L457
L459:
	;
	v1403 = F__emscripten_memcpy_bulkmem(m, v1395+int32(2), v1066, v1392)
	mBase = m.M
	goto L461
L460:
	;
	goto L461
L461:
	;
	goto L458
L462:
	;
	F_pfree(m, v1405)
	mBase = m.M
	v1407 = m.ExcPending
	if v1407 != 0 {
		goto L1
	} else {
		goto L465
	}
L463:
	;
	goto L464
L464:
	;
	v1408 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v1408
	v1410 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1410
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1395
	v1413 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v1395 + v1413
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v1369
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v1425 = v1369 + v1422<<(uint(v1413)%32)
	if v1425+int32(4) < int32(0) {
		goto L467
	} else {
		goto L468
	}
L465:
	;
	goto L464
L466:
	;
	goto L6
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(0)
	goto L466
L468:
	;
	goto L469
L469:
	;
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v1436 = v1369 & int32(3)
	v1440 = base.I32_div_s(v1425+int32(7), int32(4))
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v1441 <= v1440 {
		goto L474
	} else {
		goto L475
	}
L470:
	;
	goto L466
L471:
	;
	if int32(0) <= v1507 {
		goto L470
	} else {
		goto L492
	}
L472:
	;
	v1487 = v1481
	goto L486
L473:
	;
	v1454 = int32(1)
	v1455 = v1440 - v1454
	v1458 = v1434 + v1455<<(uint(v1454)%32)
	v1459 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1458))))
	v1460 = int32(2)
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1436<<(uint(v1460)%32))+uint32(_consts[1081])))
	v1465 = base.I32_rem_s(v1459, v1464)
	v1466 = v1459 - v1465
	*(*uint16)(unsafe.Add(mBase, uint32(v1458))) = uint16(v1466)
	v1469 = base.I32_div_s(v1464, v1460)
	if v1465 < v1469 {
		v1507 = v1455
		goto L471
	} else {
		goto L481
	}
L474:
	;
	if v1436 == int32(0) {
		goto L470
	} else {
		goto L477
	}
L475:
	;
	goto L476
L476:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1440
	if v1436 != 0 {
		goto L473
	} else {
		goto L479
	}
L477:
	;
	if v1440 != v1441 {
		goto L470
	} else {
		goto L478
	}
L478:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1440
	goto L473
L479:
	;
	v1451 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1434+v1440<<(uint(int32(1))%32)))))
	if v1451 <= int32(4999) {
		v1507 = v1440
		goto L471
	} else {
		goto L480
	}
L480:
	;
	v1481 = v1440
	goto L472
L481:
	;
	v1472 = v1464 + base.I32_extend16_s(v1466)
	if int32(9999) < v1472 {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	v1477 = v1472 + int32(55536)
	goto L484
L483:
	;
	v1477 = v1472
	goto L484
L484:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1458))) = uint16(v1477)
	if v1472 < int32(10000) {
		v1507 = v1455
		goto L471
	} else {
		goto L485
	}
L485:
	;
	v1481 = v1455
	goto L472
L486:
	;
	v1493 = int32(1)
	v1494 = v1487 - v1493
	v1497 = v1434 + v1494<<(uint(v1493)%32)
	v1500 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1497))))
	v1502 = base.B2i32(int32(9998) < v1500)
	if int32(9998) < v1500 {
		goto L488
	} else {
		goto L489
	}
L487:
	;
	v1507 = v1494
	goto L471
L488:
	;
	v1503 = int32(-9999)
	goto L490
L489:
	;
	v1503 = v1493
	goto L490
L490:
	;
	v1504 = v1503 + v1500
	*(*uint16)(unsafe.Add(mBase, uint32(v1497))) = uint16(v1504)
	if int32(9998) < v1500 {
		v1487 = v1494
		goto L486
	} else {
		goto L491
	}
L491:
	;
	goto L487
L492:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v1515 - int32(2)
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v1520 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v1519 + v1520
	v1523 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v1523 + v1520
	goto L470
L493:
	;
	goto L6
L494:
	;
	goto L6
L495:
	;
	if v1139 < int64(0) {
		goto L11
	} else {
		goto L498
	}
L496:
	;
	goto L497
L497:
	;
	v1569 = v1021 & int32(-2)
	v1571 = v1569 + int32(2)
	v1572 = F_palloc(m, v1571)
	mBase = m.M
	v1573 = m.ExcPending
	if v1573 != 0 {
		goto L1
	} else {
		goto L503
	}
L498:
	;
	v1558 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1558 != 0 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	F_pfree(m, v1558)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L1
	} else {
		goto L502
	}
L500:
	;
	goto L501
L501:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v1369
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
	v1564 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1564
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1564
	goto L6
L502:
	;
	goto L501
L503:
	;
	v1574 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1572))) = uint16(v1574)
	v1577 = v1572 + int32(2)
	if v1569 != 0 {
		goto L505
	} else {
		goto L506
	}
L504:
	;
	v1580 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v1580
	*(*int32)(unsafe.Add(mBase, uint32(v19)+140)) = v1579
	*(*int32)(unsafe.Add(mBase, uint32(v19)+136)) = v1572
	v1584 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = v1584
	v1587 = v1144 >> (uint(int32(31)) % 32)
	v1589 = v1144 ^ v1587 - v1587
	v1594 = F_log(m, base.F64_abs(base.F64_convert_i32_s(v1144)))
	mBase = m.M
	if base.F64_lt(base.F64_abs(v1594), float64(2.147483648e+09)) != 0 {
		goto L509
	} else {
		goto L510
	}
L505:
	;
	v1578 = F__emscripten_memcpy_bulkmem(m, v1577, v1066, v1569)
	mBase = m.M
	v1579 = v1578
	goto L507
L506:
	;
	v1579 = v1577
	goto L507
L507:
	;
	goto L504
L508:
	;
	if v1589&int32(1) != 0 {
		goto L513
	} else {
		goto L514
	}
L509:
	;
	v1598 = base.I32_trunc_f64_s(v1594)
	v1600 = v1598
	goto L508
L510:
	;
	goto L511
L511:
	;
	v1600 = int32(-2147483648)
	goto L508
L512:
	;
	v1638 = v1600 + v1361 + v1369 + int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v1636
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1634
	v1643 = v1589
	goto L531
L513:
	;
	v1603 = F_palloc(m, v1571)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L1
	} else {
		goto L516
	}
L514:
	;
	goto L515
L515:
	;
	v1619 = F_palloc(m, int32(4))
	mBase = m.M
	v1620 = m.ExcPending
	if v1620 != 0 {
		goto L1
	} else {
		goto L525
	}
L516:
	;
	v1605 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1603))) = uint16(v1605)
	v1608 = v1603 + int32(2)
	if v1569 != 0 {
		goto L518
	} else {
		goto L519
	}
L517:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1611 != 0 {
		goto L521
	} else {
		goto L522
	}
L518:
	;
	v1609 = F__emscripten_memcpy_bulkmem(m, v1608, v1066, v1569)
	mBase = m.M
	goto L520
L519:
	;
	goto L520
L520:
	;
	goto L517
L521:
	;
	F_pfree(m, v1611)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L1
	} else {
		goto L524
	}
L522:
	;
	goto L523
L523:
	;
	v1614 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v1614
	v1616 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1616
	v1634 = v1603
	v1636 = v1608
	goto L512
L524:
	;
	goto L523
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1619))) = int32(65536)
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1623 != 0 {
		goto L526
	} else {
		goto L527
	}
L526:
	;
	F_pfree(m, v1623)
	mBase = m.M
	v1625 = m.ExcPending
	if v1625 != 0 {
		goto L1
	} else {
		goto L529
	}
L527:
	;
	goto L528
L528:
	;
	v1629 = *(*int64)(unsafe.Add(mBase, _consts[1083]))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v1629
	v1632 = *(*int64)(unsafe.Add(mBase, _consts[1084]))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1632
	v1634 = v1619
	v1636 = v1619 + int32(2)
	goto L512
L529:
	;
	goto L528
L530:
	;
	v1727 = *(*int32)(unsafe.Add(mBase, uint32(v19)+136))
	if v1727 != 0 {
		goto L562
	} else {
		goto L563
	}
L531:
	;
	if base.Ui32(v1643) < base.Ui32(int32(2)) {
		goto L530
	} else {
		goto L533
	}
L532:
	;
	if int64(0) <= v1139 {
		goto L10
	} else {
		goto L555
	}
L533:
	;
	v1660 = v19 + int32(120)
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v19)+124))
	v1668 = v1638 - v1665<<(uint(int32(3))%32)
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v19)+132))
	v1671 = v1669 << (uint(int32(1)) % 32)
	if v1668 < v1671 {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	v1673 = v1668
	goto L536
L535:
	;
	v1673 = v1671
	goto L536
L536:
	;
	v1674 = int32(0)
	if v1674 < v1673 {
		goto L537
	} else {
		goto L538
	}
L537:
	;
	v1677 = v1673
	goto L539
L538:
	;
	v1677 = v1674
	goto L539
L539:
	;
	F_mul_var(m, v1660, v1660, v1660, v1677)
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L1
	} else {
		goto L540
	}
L540:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v19)+124))
	if v1643&int32(2) != 0 {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v1685 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v1689 = v1638 - (v1685+v1680)<<(uint(int32(2))%32)
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v19)+132))
	v1692 = v1690 + v1691
	if v1689 < v1692 {
		goto L544
	} else {
		goto L545
	}
L542:
	;
	goto L543
L543:
	;
	if v1680 <= int32(32767) {
		goto L551
	} else {
		goto L552
	}
L544:
	;
	v1694 = v1689
	goto L546
L545:
	;
	v1694 = v1692
	goto L546
L546:
	;
	v1695 = int32(0)
	if v1695 < v1694 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v1698 = v1694
	goto L549
L548:
	;
	v1698 = v1695
	goto L549
L549:
	;
	F_mul_var(m, v19+int32(120), v19, v19, v1698)
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L1
	} else {
		goto L550
	}
L550:
	;
	goto L543
L551:
	;
	v1707 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v1707 < int32(32768) {
		v1643 = int32(base.Ui32(v1643) >> (uint(int32(1)) % 32))
		goto L531
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	goto L532
L554:
	;
	goto L553
L555:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1713 != 0 {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	F_pfree(m, v1713)
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L1
	} else {
		goto L559
	}
L557:
	;
	goto L558
L558:
	;
	v1716 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v1716
	v1718 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1718
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1718
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v19)+136))
	if v1722 == v1716 {
		goto L7
	} else {
		goto L560
	}
L559:
	;
	goto L558
L560:
	;
	F_pfree(m, v1722)
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L1
	} else {
		goto L561
	}
L561:
	;
	goto L7
L562:
	;
	F_pfree(m, v1727)
	mBase = m.M
	v1729 = m.ExcPending
	if v1729 != 0 {
		goto L1
	} else {
		goto L565
	}
L563:
	;
	goto L564
L564:
	;
	if int64(0) <= v1139 {
		goto L7
	} else {
		goto L566
	}
L565:
	;
	goto L564
L566:
	;
	F_div_var(m, int32(1721448), v19, v19, v1369, int32(1), int32(0))
	mBase = m.M
	v1736 = m.ExcPending
	if v1736 != 0 {
		goto L1
	} else {
		goto L567
	}
L567:
	;
	goto L6
L568:
	;
	v1741 = F_palloc(m, int32(2))
	mBase = m.M
	v1742 = m.ExcPending
	if v1742 != 0 {
		goto L1
	} else {
		goto L571
	}
L569:
	;
	goto L570
L570:
	;
	v1760 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+136)) = v1760
	*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v1760
	*(*int64)(unsafe.Add(mBase, uint32(v19)+104)) = v1760
	*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v1760
	*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v1760
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v1760
	*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = v1760
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v1760
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v1760
	if v1049 != int32(16384) {
		goto L577
	} else {
		goto L578
	}
L571:
	;
	v1743 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1741))) = uint16(v1743)
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1745 != 0 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	F_pfree(m, v1745)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L1
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v1749 = *(*int32)(unsafe.Add(mBase, _consts[1085]))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v1749
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(16)
	v1754 = *(*int64)(unsafe.Add(mBase, _consts[1086]))
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1754
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v1741
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v1741 + int32(2)
	goto L6
L575:
	;
	goto L574
L576:
	;
	v1831 = F_estimate_ln_dweight(m, v1827)
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L1
	} else {
		goto L591
	}
L577:
	;
	v1780 = int32(0)
	v1823 = v1780
	v1826 = v1780
	v1827 = v19 + int32(48)
	goto L576
L578:
	;
	goto L579
L579:
	;
	if base.Ui32(v1078) < base.Ui32(int32(2)) {
		goto L581
	} else {
		goto L582
	}
L580:
	;
	v1800 = v1021 & int32(-2)
	v1803 = F_palloc(m, v1800+int32(2))
	mBase = m.M
	v1804 = m.ExcPending
	if v1804 != 0 {
		goto L1
	} else {
		goto L586
	}
L581:
	;
	v1798 = int32(0)
	goto L580
L582:
	;
	if v1126 < v1080 {
		goto L9
	} else {
		goto L583
	}
L583:
	;
	if v1126 != v1080 {
		goto L581
	} else {
		goto L584
	}
L584:
	;
	v1788 = int32(1)
	v1794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1123+v1080<<(uint(v1788)%32)-int32(2)))))
	if v1794&v1788 != 0 {
		v1798 = v1788
		goto L580
	} else {
		goto L585
	}
L585:
	;
	goto L581
L586:
	;
	v1805 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1803))) = uint16(v1805)
	v1808 = v1803 + int32(2)
	if v1800 != 0 {
		goto L588
	} else {
		goto L589
	}
L587:
	;
	v1812 = v19 + int32(128)
	v1813 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v1812))) = v1813
	*(*int32)(unsafe.Add(mBase, uint32(v1812))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+140)) = v1810
	*(*int32)(unsafe.Add(mBase, uint32(v19)+136)) = v1803
	v1819 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v19)+120)) = v1819
	v1823 = v1803
	v1826 = v1798
	v1827 = v19 + int32(120)
	goto L576
L588:
	;
	v1809 = F__emscripten_memcpy_bulkmem(m, v1808, v1066, v1800)
	mBase = m.M
	v1810 = v1809
	goto L590
L589:
	;
	v1810 = v1808
	goto L590
L590:
	;
	goto L587
L591:
	;
	v1833 = int32(8) - v1831
	v1834 = int32(0)
	if v1834 < v1833 {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v1837 = v1833
	goto L594
L593:
	;
	v1837 = v1834
	goto L594
L594:
	;
	F_ln_var(m, v1827, v19+int32(96), v1837)
	mBase = m.M
	v1839 = m.ExcPending
	if v1839 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	F_mul_var(m, v19+int32(96), v19+int32(24), v19+int32(72), v1837)
	mBase = m.M
	v1847 = m.ExcPending
	if v1847 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	v1850 = F_numericvar_to_double_no_overflow(m, v19+int32(72))
	mBase = m.M
	v1851 = m.ExcPending
	if v1851 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	if base.F64_gt(base.F64_abs(v1850), float64(6020)) != 0 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	if base.F64_gt(v1850, float64(0)) != 0 {
		goto L8
	} else {
		goto L601
	}
L599:
	;
	goto L600
L600:
	;
	v1869 = base.F64_mul(v1850, float64(0.434294481903252))
	if base.F64_lt(base.F64_abs(v1869), float64(2.147483648e+09)) != 0 {
		goto L607
	} else {
		goto L608
	}
L601:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v1857 != 0 {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	F_pfree(m, v1857)
	mBase = m.M
	v1859 = m.ExcPending
	if v1859 != 0 {
		goto L1
	} else {
		goto L605
	}
L603:
	;
	goto L604
L604:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(4294967296000)
	v1862 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = v1862
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v1862
	goto L6
L605:
	;
	goto L604
L606:
	;
	v1878 = int32(16) - v1875
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v1827)+12))
	if v1879 < v1878 {
		goto L610
	} else {
		goto L611
	}
L607:
	;
	v1873 = base.I32_trunc_f64_s(v1869)
	v1875 = v1873
	goto L606
L608:
	;
	goto L609
L609:
	;
	v1875 = int32(-2147483648)
	goto L606
L610:
	;
	v1881 = v1878
	goto L612
L611:
	;
	v1881 = v1879
	goto L612
L612:
	;
	if v1118 < v1881 {
		goto L613
	} else {
		goto L614
	}
L613:
	;
	v1883 = v1881
	goto L615
L614:
	;
	v1883 = v1118
	goto L615
L615:
	;
	if int32(1000) <= v1883 {
		goto L616
	} else {
		goto L617
	}
L616:
	;
	v1886 = int32(1000)
	goto L618
L617:
	;
	v1886 = v1883
	goto L618
L618:
	;
	v1887 = v1875 + v1886
	v1888 = int32(0)
	if v1888 < v1887 {
		goto L619
	} else {
		goto L620
	}
L619:
	;
	v1891 = v1887
	goto L621
L620:
	;
	v1891 = v1888
	goto L621
L621:
	;
	v1894 = v1891 - v1831 + int32(8)
	v1895 = int32(0)
	if v1895 < v1894 {
		goto L622
	} else {
		goto L623
	}
L622:
	;
	v1898 = v1894
	goto L624
L623:
	;
	v1898 = v1895
	goto L624
L624:
	;
	F_ln_var(m, v1827, v19+int32(96), v1898)
	mBase = m.M
	v1900 = m.ExcPending
	if v1900 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	F_mul_var(m, v19+int32(96), v19+int32(24), v19+int32(72), v1898)
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	F_exp_var(m, v19+int32(72), v19, v1886)
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L1
	} else {
		goto L627
	}
L627:
	;
	v1913 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if base.B2i32(int32(0) < v1913)&v1826 != 0 {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(16384)
	goto L630
L629:
	;
	goto L630
L630:
	;
	v1919 = *(*int32)(unsafe.Add(mBase, uint32(v19)+88))
	if v1919 != 0 {
		goto L631
	} else {
		goto L632
	}
L631:
	;
	F_pfree(m, v1919)
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L1
	} else {
		goto L634
	}
L632:
	;
	goto L633
L633:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v19)+112))
	if v1922 != 0 {
		goto L635
	} else {
		goto L636
	}
L634:
	;
	goto L633
L635:
	;
	F_pfree(m, v1922)
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L1
	} else {
		goto L638
	}
L636:
	;
	goto L637
L637:
	;
	if v1823 == int32(0) {
		goto L6
	} else {
		goto L639
	}
L638:
	;
	goto L637
L639:
	;
	F_pfree(m, v1823)
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L1
	} else {
		goto L640
	}
L640:
	;
	goto L6
L641:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L1
	} else {
		goto L642
	}
L642:
	;
	F_errmsg(m, int32(450000), int32(0))
	mBase = m.M
	v1939 = m.ExcPending
	if v1939 != 0 {
		goto L1
	} else {
		goto L643
	}
L643:
	;
	F_errfinish(m, int32(496807), int32(4101), int32(212849))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L1
	} else {
		goto L644
	}
L644:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L645:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L1
	} else {
		goto L646
	}
L646:
	;
	F_errmsg(m, int32(97382), int32(0))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	F_errfinish(m, int32(496807), int32(4105), int32(212849))
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L1
	} else {
		goto L648
	}
L648:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L649:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L1
	} else {
		goto L650
	}
L650:
	;
	F_errmsg(m, int32(450000), int32(0))
	mBase = m.M
	v1971 = m.ExcPending
	if v1971 != 0 {
		goto L1
	} else {
		goto L651
	}
L651:
	;
	F_errfinish(m, int32(496807), int32(4213), int32(212849))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L1
	} else {
		goto L652
	}
L652:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L653:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L1
	} else {
		goto L654
	}
L654:
	;
	F_errmsg(m, int32(111609), int32(0))
	mBase = m.M
	v1987 = m.ExcPending
	if v1987 != 0 {
		goto L1
	} else {
		goto L655
	}
L655:
	;
	F_errfinish(m, int32(496807), int32(11498), int32(90637))
	mBase = m.M
	v1992 = m.ExcPending
	if v1992 != 0 {
		goto L1
	} else {
		goto L656
	}
L656:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L657:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v1999 = m.ExcPending
	if v1999 != 0 {
		goto L1
	} else {
		goto L658
	}
L658:
	;
	F_errmsg(m, int32(238346), int32(0))
	mBase = m.M
	v2003 = m.ExcPending
	if v2003 != 0 {
		goto L1
	} else {
		goto L659
	}
L659:
	;
	F_errfinish(m, int32(496807), int32(11551), int32(90637))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L1
	} else {
		goto L660
	}
L660:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L661:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L662
	}
L662:
	;
	F_errmsg(m, int32(111609), int32(0))
	mBase = m.M
	v2019 = m.ExcPending
	if v2019 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	F_errfinish(m, int32(496807), int32(11633), int32(90637))
	mBase = m.M
	v2024 = m.ExcPending
	if v2024 != 0 {
		goto L1
	} else {
		goto L664
	}
L664:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L665:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L1
	} else {
		goto L666
	}
L666:
	;
	F_errmsg(m, int32(97382), int32(0))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L1
	} else {
		goto L667
	}
L667:
	;
	F_errfinish(m, int32(496807), int32(11346), int32(228222))
	mBase = m.M
	v2040 = m.ExcPending
	if v2040 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L669:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	F_errmsg(m, int32(111609), int32(0))
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	F_errfinish(m, int32(496807), int32(11404), int32(228222))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L1
	} else {
		goto L672
	}
L672:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L673:
	;
	goto L6
L674:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19))) = int64(0)
	goto L673
L675:
	;
	goto L676
L676:
	;
	v2079 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v2081 = v1369 & int32(3)
	v2085 = base.I32_div_s(v2070+int32(7), int32(4))
	v2086 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v2086 <= v2085 {
		goto L681
	} else {
		goto L682
	}
L677:
	;
	goto L673
L678:
	;
	if int32(0) <= v2152 {
		goto L677
	} else {
		goto L699
	}
L679:
	;
	v2132 = v2126
	goto L693
L680:
	;
	v2099 = int32(1)
	v2100 = v2085 - v2099
	v2103 = v2079 + v2100<<(uint(v2099)%32)
	v2104 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2103))))
	v2105 = int32(2)
	v2109 = *(*int32)(unsafe.Add(mBase, uint32(v2081<<(uint(v2105)%32))+uint32(_consts[1081])))
	v2110 = base.I32_rem_s(v2104, v2109)
	v2111 = v2104 - v2110
	*(*uint16)(unsafe.Add(mBase, uint32(v2103))) = uint16(v2111)
	v2114 = base.I32_div_s(v2109, v2105)
	if v2110 < v2114 {
		v2152 = v2100
		goto L678
	} else {
		goto L688
	}
L681:
	;
	if v2081 == int32(0) {
		goto L677
	} else {
		goto L684
	}
L682:
	;
	goto L683
L683:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v2085
	if v2081 != 0 {
		goto L680
	} else {
		goto L686
	}
L684:
	;
	if v2085 != v2086 {
		goto L677
	} else {
		goto L685
	}
L685:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v2085
	goto L680
L686:
	;
	v2096 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2079+v2085<<(uint(int32(1))%32)))))
	if v2096 <= int32(4999) {
		v2152 = v2085
		goto L678
	} else {
		goto L687
	}
L687:
	;
	v2126 = v2085
	goto L679
L688:
	;
	v2117 = v2109 + base.I32_extend16_s(v2111)
	if int32(9999) < v2117 {
		goto L689
	} else {
		goto L690
	}
L689:
	;
	v2122 = v2117 + int32(55536)
	goto L691
L690:
	;
	v2122 = v2117
	goto L691
L691:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2103))) = uint16(v2122)
	if v2117 < int32(10000) {
		v2152 = v2100
		goto L678
	} else {
		goto L692
	}
L692:
	;
	v2126 = v2100
	goto L679
L693:
	;
	v2138 = int32(1)
	v2139 = v2132 - v2138
	v2142 = v2079 + v2139<<(uint(v2138)%32)
	v2145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2142))))
	v2147 = base.B2i32(int32(9998) < v2145)
	if int32(9998) < v2145 {
		goto L695
	} else {
		goto L696
	}
L694:
	;
	v2152 = v2139
	goto L678
L695:
	;
	v2148 = int32(-9999)
	goto L697
L696:
	;
	v2148 = v2138
	goto L697
L697:
	;
	v2149 = v2148 + v2145
	*(*uint16)(unsafe.Add(mBase, uint32(v2142))) = uint16(v2149)
	if int32(9998) < v2145 {
		v2132 = v2139
		goto L693
	} else {
		goto L698
	}
L698:
	;
	goto L694
L699:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v2160 - int32(2)
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v2165 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v2164 + v2165
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v2168 + v2165
	goto L677
L700:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
	if v2205 == int32(0) {
		v2219 = v2203
		goto L4
	} else {
		goto L701
	}
L701:
	;
	F_pfree(m, v2205)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L1
	} else {
		goto L702
	}
L702:
	;
	v2219 = v2203
	goto L4
L703:
	;
	v2219 = v2215
	goto L4
}
func F_numeric_random(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v136 int64
	_ = v136
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v146 int64
	_ = v146
	var v151 int64
	_ = v151
	var v156 int64
	_ = v156
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v293 int64
	_ = v293
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v325 int32
	_ = v325
	var v327 int64
	_ = v327
	var v330 int32
	_ = v330
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v367 int64
	_ = v367
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v492 int32
	_ = v492
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v519 int64
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v530 int64
	_ = v530
	var v532 int64
	_ = v532
	var v536 int64
	_ = v536
	var v539 int64
	_ = v539
	var v541 int64
	_ = v541
	var v545 int64
	_ = v545
	var v548 int64
	_ = v548
	var v551 int32
	_ = v551
	var v553 int64
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int64
	_ = v566
	var v578 int32
	_ = v578
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int64
	_ = v604
	var v605 int64
	_ = v605
	var v612 int64
	_ = v612
	var v614 int64
	_ = v614
	var v615 int64
	_ = v615
	var v618 int64
	_ = v618
	var v620 int64
	_ = v620
	var v624 int64
	_ = v624
	var v626 int64
	_ = v626
	var v634 int64
	_ = v634
	var v639 int64
	_ = v639
	var v652 int64
	_ = v652
	var v654 int32
	_ = v654
	var v655 int64
	_ = v655
	var v662 int64
	_ = v662
	var v664 int64
	_ = v664
	var v665 int64
	_ = v665
	var v668 int64
	_ = v668
	var v670 int64
	_ = v670
	var v674 int64
	_ = v674
	var v676 int64
	_ = v676
	var v684 int64
	_ = v684
	var v689 int64
	_ = v689
	var v702 int64
	_ = v702
	var v703 int64
	_ = v703
	var v706 int32
	_ = v706
	var v707 int64
	_ = v707
	var v708 int64
	_ = v708
	var v711 int64
	_ = v711
	var v715 int32
	_ = v715
	var v718 int64
	_ = v718
	var v725 int64
	_ = v725
	var v727 int64
	_ = v727
	var v732 int64
	_ = v732
	var v734 int64
	_ = v734
	var v737 int32
	_ = v737
	var v739 int32
	_ = v739
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v777 int64
	_ = v777
	var v779 int64
	_ = v779
	var v780 int64
	_ = v780
	var v783 int64
	_ = v783
	var v785 int64
	_ = v785
	var v789 int64
	_ = v789
	var v791 int64
	_ = v791
	var v799 int64
	_ = v799
	var v804 int64
	_ = v804
	var v808 int64
	_ = v808
	var v819 int64
	_ = v819
	var v822 int64
	_ = v822
	var v823 int64
	_ = v823
	var v824 int64
	_ = v824
	var v827 int64
	_ = v827
	var v829 int64
	_ = v829
	var v833 int64
	_ = v833
	var v836 int32
	_ = v836
	var v838 int32
	_ = v838
	var v865 int32
	_ = v865
	var v894 int32
	_ = v894
	var v903 int64
	_ = v903
	var v905 int64
	_ = v905
	var v906 int64
	_ = v906
	var v909 int64
	_ = v909
	var v911 int64
	_ = v911
	var v915 int64
	_ = v915
	var v917 int64
	_ = v917
	var v925 int64
	_ = v925
	var v930 int64
	_ = v930
	var v934 int64
	_ = v934
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int64
	_ = v979
	var v981 int32
	_ = v981
	var v984 int64
	_ = v984
	var v991 int64
	_ = v991
	var v993 int64
	_ = v993
	var v994 int64
	_ = v994
	var v997 int64
	_ = v997
	var v999 int64
	_ = v999
	var v1003 int64
	_ = v1003
	var v1005 int64
	_ = v1005
	var v1013 int64
	_ = v1013
	var v1018 int64
	_ = v1018
	var v1031 int64
	_ = v1031
	var v1032 int64
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1076 int32
	_ = v1076
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1140 int32
	_ = v1140
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1222 int32
	_ = v1222
	var v1231 int32
	_ = v1231
	var v1233 int32
	_ = v1233
	var v1240 int32
	_ = v1240
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1251 int32
	_ = v1251
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1271 int32
	_ = v1271
	var v1275 int32
	_ = v1275
	var v1280 int32
	_ = v1280
	var v1284 int32
	_ = v1284
	var v1290 int32
	_ = v1290
	var v1301 int32
	_ = v1301
	var v1311 int32
	_ = v1311
	var v1325 int32
	_ = v1325
	var v1383 int32
	_ = v1383
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1432 int32
	_ = v1432
	var v1437 int32
	_ = v1437
	var v1441 int32
	_ = v1441
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1455 int32
	_ = v1455
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v33 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1106])))
	if v36 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = int32(16)
	v41 = int32(0)
	v45 = m.G0
	v47 = v45 - v40
	m.G0 = v47
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v41
	v53 = F_open(m, int32(286717), v41, v47)
	mBase = m.M
	if v53 != int32(-1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	v171 = m.G0
	v173 = v171 - int32(96)
	m.G0 = v173
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	v176 = base.I32_extend16_s(v175)
	if base.Ui32(int32(49152)) <= base.Ui32(v175) {
		goto L36
	} else {
		goto L37
	}
L7:
	;
	v169 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1106])) = uint8(v169)
	goto L6
L8:
	;
	if v86 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	goto L13
L10:
	;
	v86 = v41
	goto L11
L11:
	;
	m.G0 = v47 + int32(16)
	goto L8
L12:
	;
	v81 = F_close(m, v53)
	mBase = m.M
	v86 = v79
	goto L11
L13:
	;
	v59 = int32(4471776)
	v60 = v40
	goto L14
L14:
	;
	v65 = F_read(m, v53, v59, v60)
	mBase = m.M
	if v65 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v79 = int32(1)
	goto L12
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	if v69 == int32(27) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v74 = v60 - v65
	if v74 != 0 {
		v59 = v59 + v65
		v60 = v74
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v79 = int32(0)
	goto L12
L20:
	;
	goto L15
L21:
	;
	v91 = int32(4471776)
	v92 = *(*int64)(unsafe.Add(mBase, _consts[1107]))
	if v92 != int64(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v103 = int32(4471776)
	v107 = m.G0
	v108 = int32(16)
	v109 = v107 - v108
	m.G0 = v109
	F___gettimeofday(m, v109)
	mBase = m.M
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
	v113 = int64(*(*int32)(unsafe.Add(mBase, uint32(v109)+8)))
	m.G0 = v109 + v108
	goto L29
L24:
	;
	goto L7
L25:
	;
	goto L24
L26:
	;
	v95 = *(*int64)(unsafe.Add(mBase, _consts[1108]))
	if v95 != int64(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1108])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _consts[1107])) = int64(6364136223846793005)
	goto L25
L29:
	;
	v123 = int64(*(*uint32)(unsafe.Add(mBase, _consts[451])))
	v126 = v113 + v112*int64(1000000) - int64(946684800000000) ^ v123<<(uint(int64(32))%64)
	v130 = v126 + int64(4354685564936845354)
	v131 = int64(30)
	v134 = int64(-4658895280553007687)
	v135 = (int64(base.Ui64(v130)>>(uint(v131)%64)) ^ v130) * v134
	v136 = int64(27)
	v139 = int64(-7723592293110705685)
	v140 = (int64(base.Ui64(v135)>>(uint(v136)%64)) ^ v135) * v139
	v141 = int64(31)
	*(*int64)(unsafe.Add(mBase, _consts[1108])) = int64(base.Ui64(v140)>>(uint(v141)%64)) ^ v140
	v146 = v126 - int64(7046029254386353131)
	v151 = (int64(base.Ui64(v146)>>(uint(v131)%64)) ^ v146) * v134
	v156 = (int64(base.Ui64(v151)>>(uint(v136)%64)) ^ v151) * v139
	*(*int64)(unsafe.Add(mBase, _consts[1107])) = int64(base.Ui64(v156)>>(uint(v141)%64)) ^ v156
	if v146|v130 == int64(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L7
L31:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1108])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _consts[1107])) = int64(6364136223846793005)
	goto L33
L32:
	;
	goto L33
L33:
	;
	goto L30
L34:
	;
	return v1414
L35:
	;
	F_errmsg(m, int32(524374), int32(0))
	mBase = m.M
	v1450 = m.ExcPending
	if v1450 != 0 {
		goto L1
	} else {
		goto L277
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)))
	v198 = base.I32_extend16_s(v197)
	if base.Ui32(int32(49152)) <= base.Ui32(v197) {
		goto L45
	} else {
		goto L46
	}
L39:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v176 == int32(-16384) {
		goto L35
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(11349), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(496807), int32(4364), int32(487533))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errmsg(m, int32(524426), int32(0))
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L1
	} else {
		goto L275
	}
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v225 = base.B2i32(int32(0) <= v176)
	if int32(0) <= v176 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	if v198 == int32(-16384) {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(11380), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(496807), int32(4375), int32(487533))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
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
	v226 = int32(-8)
	goto L55
L54:
	;
	v226 = int32(-6)
	goto L55
L55:
	;
	v227 = int32(base.Ui32(v219)>>(uint(int32(2))%32)) + v226
	*(*int32)(unsafe.Add(mBase, uint32(v173)+48)) = int32(base.Ui32(v227) >> (uint(int32(1)) % 32))
	if int32(0) <= v176 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v231 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+6)))
	v241 = v231
	goto L58
L57:
	;
	v241 = v175<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v175&int32(63)
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+52)) = v241
	v243 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+64)) = v243
	v252 = base.B2i32(v176 < v243)
	if v176 < v243 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v253 = int32(base.Ui32(v175)>>(uint(int32(7))%32)) & int32(63)
	goto L61
L60:
	;
	v253 = v175 & int32(16383)
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+60)) = v253
	v260 = v175 & int32(49152)
	if v260 == int32(32768) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v263 = v175 << (uint(int32(1)) % 32) & int32(16384)
	goto L64
L63:
	;
	v263 = v260
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+56)) = v263
	if v176 < v243 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v267 = int32(6)
	goto L67
L66:
	;
	v267 = int32(8)
	goto L67
L67:
	;
	v268 = v28 + v267
	*(*int32)(unsafe.Add(mBase, uint32(v173)+68)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v276 = base.B2i32(int32(0) <= v198)
	if int32(0) <= v198 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v277 = int32(-8)
	goto L70
L69:
	;
	v277 = int32(-6)
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+24)) = int32(base.Ui32(int32(base.Ui32(v270)>>(uint(int32(2))%32))+v277) >> (uint(int32(1)) % 32))
	if int32(0) <= v198 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v282 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+6)))
	v292 = v282
	goto L73
L72:
	;
	v292 = v197<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v197&int32(63)
	goto L73
L73:
	;
	v293 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v173)+8)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v173)+16)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v173)+28)) = v292
	*(*int64)(unsafe.Add(mBase, uint32(v173))) = v293
	v305 = v197 & int32(49152)
	if v305 == int32(32768) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v308 = v197 << (uint(int32(1)) % 32) & int32(16384)
	goto L76
L75:
	;
	v308 = v305
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+32)) = v308
	v310 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v173)+40)) = v310
	v315 = base.B2i32(v198 < v310)
	if v198 < v310 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v316 = int32(6)
	goto L79
L78:
	;
	v316 = int32(8)
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+44)) = v33 + v316
	if v198 < v310 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v325 = int32(base.Ui32(v197)>>(uint(int32(7))%32)) & int32(63)
	goto L82
L81:
	;
	v325 = v197 & int32(16383)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+36)) = v325
	v327 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v173)+88)) = v327
	v330 = v173 + int32(80)
	*(*int64)(unsafe.Add(mBase, uint32(v330))) = v327
	*(*int64)(unsafe.Add(mBase, uint32(v173)+72)) = v327
	F_sub_var(m, v173+int32(24), v173+int32(48), v173+int32(72))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	if v343 != int32(16384) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if base.Ui32(v325) < base.Ui32(v253) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L1
	} else {
		goto L271
	}
L87:
	;
	v347 = v253
	goto L89
L88:
	;
	v347 = v325
	goto L89
L89:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v173)+72))
	if v348 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(v173)+88))
	if v1410 != 0 {
		goto L262
	} else {
		goto L263
	}
L91:
	;
	v352 = v227 & int32(-2)
	v355 = F_palloc(m, v352+int32(2))
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v374 = int32(1)
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v173)+76))
	v377 = v347 + int32(3)
	v380 = v375 + int32(base.Ui32(v377)>>(uint(int32(2))%32))
	v382 = v380 + v374
	v385 = v377 & int32(32764)
	v386 = v385 - v347
	if v386 <= int32(0) {
		v492 = v374
		goto L102
	} else {
		goto L103
	}
L94:
	;
	v357 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v355))) = uint16(v357)
	if base.Ui32(int32(2)) <= base.Ui32(v227) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	if v352 != 0 {
		goto L99
	} else {
		goto L100
	}
L96:
	;
	goto L97
L97:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v173)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = v365
	v367 = *(*int64)(unsafe.Add(mBase, uint32(v173)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v173))) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v173)+16)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v173)+20)) = v355 + int32(2)
	goto L90
L98:
	;
	goto L97
L99:
	;
	v363 = F__emscripten_memcpy_bulkmem(m, v355+int32(2), v268, v352)
	mBase = m.M
	goto L101
L100:
	;
	goto L101
L101:
	;
	goto L98
L102:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v173)+92))
	v515 = int64(*(*int16)(unsafe.Add(mBase, uint32(v514))))
	if v382 < int32(2) {
		v551 = v374
		v553 = v515
		goto L114
	} else {
		goto L115
	}
L103:
	;
	v390 = v386 & int32(7)
	if base.Ui32(v347-v385) <= base.Ui32(int32(-8)) {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v397 = int32(0)
	v401 = v374
	goto L107
L105:
	;
	v432 = v374
	goto L106
L106:
	;
	if v390 == int32(0) {
		v492 = v432
		goto L102
	} else {
		goto L110
	}
L107:
	;
	v424 = v401 * int32(100000000)
	v426 = v397 + int32(8)
	if v426 != v386&int32(2147483640) {
		v397 = v426
		v401 = v424
		goto L107
	} else {
		goto L109
	}
L108:
	;
	v432 = v424
	goto L106
L109:
	;
	goto L108
L110:
	;
	v457 = int32(0)
	v461 = v432
	goto L111
L111:
	;
	v484 = v461 * int32(10)
	v486 = v457 + int32(1)
	if v486 != v390 {
		v457 = v486
		v461 = v484
		goto L111
	} else {
		goto L113
	}
L112:
	;
	v492 = v484
	goto L102
L113:
	;
	goto L112
L114:
	;
	v554 = int32(1)
	v555 = base.B2i32(v492 != v554)
	v559 = v382 << (uint(v554) % 32)
	if v492 != v554 {
		goto L128
	} else {
		goto L129
	}
L115:
	;
	v519 = v515 * int64(10000)
	v520 = int32(2)
	v522 = v380 - int32(1)
	if base.Ui32(v520) <= base.Ui32(v522) {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v525 = v520
	goto L118
L117:
	;
	v525 = v522
	goto L118
L118:
	;
	v527 = v525 + int32(2)
	if int32(1) < v348 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v530 = int64(*(*int16)(unsafe.Add(mBase, uint32(v514)+2)))
	v532 = v519 + v530
	goto L121
L120:
	;
	v532 = v519
	goto L121
L121:
	;
	if v522 == int32(0) {
		v551 = v527
		v553 = v532
		goto L114
	} else {
		goto L122
	}
L122:
	;
	v536 = v532 * int64(10000)
	if int32(3) <= v348 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v539 = int64(*(*int16)(unsafe.Add(mBase, uint32(v514)+4)))
	v541 = v536 + v539
	goto L125
L124:
	;
	v541 = v536
	goto L125
L125:
	;
	if v527 == int32(3) {
		v551 = v527
		v553 = v541
		goto L114
	} else {
		goto L126
	}
L126:
	;
	v545 = v541 * int64(10000)
	if v348 < int32(4) {
		v551 = v527
		v553 = v545
		goto L114
	} else {
		goto L127
	}
L127:
	;
	v548 = int64(*(*int16)(unsafe.Add(mBase, uint32(v514)+6)))
	v551 = v527
	v553 = v545 + v548
	goto L114
L128:
	;
	v562 = v380
	goto L130
L129:
	;
	v562 = v382
	goto L130
L130:
	;
	v564 = v562 - int32(3)
	v566 = base.I64_extend_i32_s(v492)
	v578 = int32(0)
	goto L134
L131:
	;
	F_add_var(m, v173, v173+int32(48), v173)
	mBase = m.M
	v1383 = m.ExcPending
	if v1383 != 0 {
		goto L1
	} else {
		goto L261
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+20)) = v1325
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v173))) = int64(0)
	goto L131
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v173)+16)) = v599
	v1325 = v737
	goto L132
L134:
	;
	if v578 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+20)) = v1116
	*(*int32)(unsafe.Add(mBase, uint32(v173)+16)) = v599
	*(*int32)(unsafe.Add(mBase, uint32(v173)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v1114
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v173)+4)) = v1119
	goto L131
L136:
	;
	F_pfree(m, v578)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v599 = F_palloc(m, v559+int32(2))
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L1
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	v601 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v599))) = uint16(v601)
	if v555&base.B2i32(v551 == v382) != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v706 = v599 + v551<<(uint(int32(1))%32)
	v707 = int64(10000)
	v708 = base.I64_div_u_s(v703, v707)
	v711 = v703 - v708*v707
	*(*uint16)(unsafe.Add(mBase, uint32(v706))) = uint16(v711)
	if v551 < int32(2) {
		goto L159
	} else {
		goto L160
	}
L142:
	;
	v603 = int32(4471776)
	v604 = int64(0)
	v605 = base.I64_div_u_s(v553, v566)
	if base.Ui64(v605) <= base.Ui64(v604) {
		goto L146
	} else {
		goto L147
	}
L143:
	;
	goto L144
L144:
	;
	v654 = int32(4471776)
	v655 = int64(0)
	if base.Ui64(v553) <= base.Ui64(v655) {
		goto L153
	} else {
		goto L154
	}
L145:
	;
	v703 = v652 * v566
	goto L141
L146:
	;
	v652 = v604
	goto L145
L147:
	;
	goto L148
L148:
	;
	v612 = v605 - v604
	v614 = *(*int64)(unsafe.Add(mBase, _consts[1108]))
	v615 = *(*int64)(unsafe.Add(mBase, _consts[1107]))
	v618 = v615
	v620 = v614
	goto L149
L149:
	;
	v624 = v618 ^ v620
	v626 = base.I64_rotl(v624, int64(37))
	v634 = v624 ^ (v624<<(uint(int64(16))%64) ^ base.I64_rotl(v618, int64(24)))
	v639 = int64(base.Ui64(base.I64_rotl(v618*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v612)) % 64))
	if base.Ui64(v612) < base.Ui64(v639) {
		v618 = v634
		v620 = v626
		goto L149
	} else {
		goto L151
	}
L150:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1108])) = v626
	*(*int64)(unsafe.Add(mBase, _consts[1107])) = v634
	v652 = v604 + v639
	goto L145
L151:
	;
	goto L150
L152:
	;
	v703 = v702
	goto L141
L153:
	;
	v702 = v655
	goto L152
L154:
	;
	goto L155
L155:
	;
	v662 = v553 - v655
	v664 = *(*int64)(unsafe.Add(mBase, _consts[1108]))
	v665 = *(*int64)(unsafe.Add(mBase, _consts[1107]))
	v668 = v665
	v670 = v664
	goto L156
L156:
	;
	v674 = v668 ^ v670
	v676 = base.I64_rotl(v674, int64(37))
	v684 = v674 ^ (v674<<(uint(int64(16))%64) ^ base.I64_rotl(v668, int64(24)))
	v689 = int64(base.Ui64(base.I64_rotl(v668*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v662)) % 64))
	if base.Ui64(v662) < base.Ui64(v689) {
		v668 = v684
		v670 = v676
		goto L156
	} else {
		goto L158
	}
L157:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1108])) = v676
	*(*int64)(unsafe.Add(mBase, _consts[1107])) = v684
	v702 = v655 + v689
	goto L152
L158:
	;
	goto L157
L159:
	;
	v737 = v599 + int32(2)
	if v551 < v564 {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	v715 = int32(2)
	v718 = base.I64_rem_u_s(v708, int64(10000))
	*(*uint16)(unsafe.Add(mBase, uint32(v706-v715))) = uint16(v718)
	if v551 == v715 {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v725 = base.I64_div_u_s(v703, int64(100000000))
	v727 = base.I64_rem_u_s(v725, int64(10000))
	*(*uint16)(unsafe.Add(mBase, uint32(v706-int32(4)))) = uint16(v727)
	if v551 < int32(4) {
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v732 = base.I64_div_u_s(v703, int64(1000000000000))
	v734 = base.I64_rem_u_s(v732, int64(10000))
	*(*uint16)(unsafe.Add(mBase, uint32(v706-int32(6)))) = uint16(v734)
	goto L159
L163:
	;
	v739 = v551
	goto L166
L164:
	;
	v838 = v551
	goto L165
L165:
	;
	if v838 < v562 {
		goto L176
	} else {
		goto L177
	}
L166:
	;
	v767 = v737 + v739<<(uint(int32(1))%32)
	v768 = int32(4471776)
	goto L170
L167:
	;
	v838 = v836
	goto L165
L168:
	;
	v819 = base.I64_div_u_s(v808, int64(1000000000000))
	*(*uint16)(unsafe.Add(mBase, uint32(v767)+6)) = uint16(v819)
	v822 = base.I64_div_u_s(v808, int64(100000000))
	v823 = int64(10000)
	v824 = base.I64_rem_u_s(v822, v823)
	*(*uint16)(unsafe.Add(mBase, uint32(v767)+4)) = uint16(v824)
	v827 = base.I64_div_u_s(v808, v823)
	v829 = base.I64_rem_u_s(v827, v823)
	*(*uint16)(unsafe.Add(mBase, uint32(v767)+2)) = uint16(v829)
	v833 = v808 - v827*v823
	*(*uint16)(unsafe.Add(mBase, uint32(v767))) = uint16(v833)
	v836 = v739 + int32(4)
	if v836 < v564 {
		v739 = v836
		goto L166
	} else {
		goto L175
	}
L170:
	;
	goto L171
L171:
	;
	v777 = int64(9999999999999999)
	v779 = *(*int64)(unsafe.Add(mBase, _consts[1108]))
	v780 = *(*int64)(unsafe.Add(mBase, _consts[1107]))
	v783 = v780
	v785 = v779
	goto L172
L172:
	;
	v789 = v783 ^ v785
	v791 = base.I64_rotl(v789, int64(37))
	v799 = v789 ^ (v789<<(uint(int64(16))%64) ^ base.I64_rotl(v783, int64(24)))
	v804 = int64(base.Ui64(base.I64_rotl(v783*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v777)) % 64))
	if base.Ui64(v777) < base.Ui64(v804) {
		v783 = v799
		v785 = v791
		goto L172
	} else {
		goto L174
	}
L173:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1108])) = v791
	*(*int64)(unsafe.Add(mBase, _consts[1107])) = v799
	v808 = int64(0) + v804
	goto L168
L174:
	;
	goto L173
L175:
	;
	goto L167
L176:
	;
	v865 = v838
	goto L179
L177:
	;
	v948 = v838
	goto L178
L178:
	;
	if v948 < v382 {
		goto L189
	} else {
		goto L190
	}
L179:
	;
	v894 = int32(4471776)
	goto L183
L180:
	;
	v948 = v562
	goto L178
L181:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v737+v865<<(uint(int32(1))%32)))) = uint16(v934)
	v946 = v865 + int32(1)
	if v946 != v562 {
		v865 = v946
		goto L179
	} else {
		goto L188
	}
L183:
	;
	goto L184
L184:
	;
	v903 = int64(9999)
	v905 = *(*int64)(unsafe.Add(mBase, _consts[1108]))
	v906 = *(*int64)(unsafe.Add(mBase, _consts[1107]))
	v909 = v906
	v911 = v905
	goto L185
L185:
	;
	v915 = v909 ^ v911
	v917 = base.I64_rotl(v915, int64(37))
	v925 = v915 ^ (v915<<(uint(int64(16))%64) ^ base.I64_rotl(v909, int64(24)))
	v930 = int64(base.Ui64(base.I64_rotl(v909*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v903)) % 64))
	if base.Ui64(v903) < base.Ui64(v930) {
		v909 = v925
		v911 = v917
		goto L185
	} else {
		goto L187
	}
L186:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1108])) = v917
	*(*int64)(unsafe.Add(mBase, _consts[1107])) = v925
	v934 = int64(0) + v930
	goto L181
L187:
	;
	goto L186
L188:
	;
	goto L180
L189:
	;
	v975 = int32(1)
	v978 = int32(4471776)
	v979 = int64(0)
	v981 = base.I32_div_s(int32(10000), v492)
	v984 = base.I64_extend_i32_s(v981 - v975)
	if base.Ui64(v984) <= base.Ui64(v979) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	goto L191
L191:
	;
	if base.B2i32(base.Ui32(int32(2147483646)) < base.Ui32(v380)) == int32(0) {
		goto L200
	} else {
		goto L201
	}
L192:
	;
	v1032 = v1031 * base.I64_extend_i32_u(v492)
	*(*uint16)(unsafe.Add(mBase, uint32(v737+v948<<(uint(v975)%32)))) = uint16(v1032)
	goto L191
L193:
	;
	v1031 = v979
	goto L192
L194:
	;
	goto L195
L195:
	;
	v991 = v984 - v979
	v993 = *(*int64)(unsafe.Add(mBase, _consts[1108]))
	v994 = *(*int64)(unsafe.Add(mBase, _consts[1107]))
	v997 = v994
	v999 = v993
	goto L196
L196:
	;
	v1003 = v997 ^ v999
	v1005 = base.I64_rotl(v1003, int64(37))
	v1013 = v1003 ^ (v1003<<(uint(int64(16))%64) ^ base.I64_rotl(v997, int64(24)))
	v1018 = int64(base.Ui64(base.I64_rotl(v997*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v991)) % 64))
	if base.Ui64(v991) < base.Ui64(v1018) {
		v997 = v1013
		v999 = v1005
		goto L196
	} else {
		goto L198
	}
L197:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1108])) = v1005
	*(*int64)(unsafe.Add(mBase, _consts[1107])) = v1013
	v1031 = v979 + v1018
	goto L192
L198:
	;
	goto L197
L199:
	;
	v1140 = int32(0)
	if base.B2i32(v375 < v1119)&base.B2i32(v1140 < v1114) == v1140 {
		goto L216
	} else {
		goto L217
	}
L200:
	;
	v1036 = v382
	v1038 = v737
	v1041 = v375
	goto L204
L201:
	;
	goto L202
L202:
	;
	if v382 == int32(0) {
		goto L133
	} else {
		goto L212
	}
L203:
	;
	v1076 = v1036
	goto L208
L204:
	;
	v1062 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1038))))
	if v1062 != 0 {
		goto L203
	} else {
		goto L206
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v173)+16)) = v599
	v1325 = v737 + v559
	goto L132
L206:
	;
	v1063 = int32(1)
	if v1063 < v1036 {
		v1036 = v1036 - v1063
		v1038 = v1038 + int32(2)
		v1041 = v1041 - v1063
		goto L204
	} else {
		goto L207
	}
L207:
	;
	goto L205
L208:
	;
	v1105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1038-int32(2)+v1076<<(uint(int32(1))%32)))))
	if v1105 != 0 {
		v1114 = v1076
		v1116 = v1038
		v1119 = v1041
		goto L199
	} else {
		goto L210
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v173)+12)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v173)+16)) = v599
	v1325 = v1038
	goto L132
L210:
	;
	v1106 = int32(1)
	if v1106 < v1076 {
		v1076 = v1076 - v1106
		goto L208
	} else {
		goto L211
	}
L211:
	;
	goto L209
L212:
	;
	v1114 = v382
	v1116 = v737
	v1119 = v375
	goto L199
L213:
	;
	if int32(0) < v1311 {
		v578 = v599
		goto L134
	} else {
		goto L260
	}
L214:
	;
	v1311 = v1301
	goto L213
L215:
	;
	if v375 <= v1171 {
		v1206 = v375
		v1208 = v1140
		goto L224
	} else {
		goto L225
	}
L216:
	;
	v1171 = v1119
	v1175 = v1140
	goto L215
L217:
	;
	goto L218
L218:
	;
	v1152 = v1119
	v1156 = v1140
	goto L219
L219:
	;
	v1162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116+v1156<<(uint(int32(1))%32)))))
	if v1162 != 0 {
		v1301 = int32(1)
		goto L214
	} else {
		goto L221
	}
L220:
	;
	v1171 = v1166
	v1175 = v1164
	goto L215
L221:
	;
	v1163 = int32(1)
	v1164 = v1156 + v1163
	v1166 = v1152 - v1163
	if v1166 <= v375 {
		v1171 = v1166
		v1175 = v1164
		goto L215
	} else {
		goto L222
	}
L222:
	;
	if v1164 < v1114 {
		v1152 = v1166
		v1156 = v1164
		goto L219
	} else {
		goto L223
	}
L223:
	;
	goto L220
L224:
	;
	if v1171 != v1206 {
		v1247 = v1175
		v1248 = v1208
		goto L232
	} else {
		goto L233
	}
L225:
	;
	if v348 <= int32(0) {
		v1206 = v375
		v1208 = v1140
		goto L224
	} else {
		goto L226
	}
L226:
	;
	v1187 = v375
	v1189 = v1140
	goto L227
L227:
	;
	v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514+v1189<<(uint(int32(1))%32)))))
	if v1194 != 0 {
		v1301 = int32(-1)
		goto L214
	} else {
		goto L229
	}
L228:
	;
	v1206 = v1198
	v1208 = v1196
	goto L224
L229:
	;
	v1195 = int32(1)
	v1196 = v1189 + v1195
	v1198 = v1187 - v1195
	if v1198 <= v1171 {
		v1206 = v1198
		v1208 = v1196
		goto L224
	} else {
		goto L230
	}
L230:
	;
	if v1196 < v348 {
		v1187 = v1198
		v1189 = v1196
		goto L227
	} else {
		goto L231
	}
L231:
	;
	goto L228
L232:
	;
	if v1114 < v1247 {
		goto L242
	} else {
		goto L243
	}
L233:
	;
	v1217 = v1175
	v1218 = v1208
	goto L234
L234:
	;
	if v1114 <= v1217 {
		v1247 = v1217
		v1248 = v1218
		goto L232
	} else {
		goto L236
	}
L235:
	;
	if base.I32_extend16_s(v1233) < base.I32_extend16_s(v1231) {
		goto L239
	} else {
		goto L240
	}
L236:
	;
	if v348 <= v1218 {
		v1247 = v1217
		v1248 = v1218
		goto L232
	} else {
		goto L237
	}
L237:
	;
	v1222 = int32(1)
	v1231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116+v1217<<(uint(v1222)%32)))))
	v1233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1218<<(uint(v1222)%32)+v514))))
	if v1231 == v1233 {
		v1217 = v1217 + v1222
		v1218 = v1218 + v1222
		goto L234
	} else {
		goto L238
	}
L238:
	;
	goto L235
L239:
	;
	v1240 = int32(1)
	goto L241
L240:
	;
	v1240 = int32(-1)
	goto L241
L241:
	;
	v1311 = v1240
	goto L213
L242:
	;
	v1251 = v1247
	goto L244
L243:
	;
	v1251 = v1114
	goto L244
L244:
	;
	v1258 = v1247
	goto L245
L245:
	;
	if v1251 == v1258 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v1301 = v1284
	goto L214
L247:
	;
	if v348 < v1248 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	goto L249
L249:
	;
	v1284 = int32(1)
	v1290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116+v1258<<(uint(v1284)%32)))))
	if v1290 == int32(0) {
		v1258 = v1258 + v1284
		goto L245
	} else {
		goto L259
	}
L250:
	;
	v1263 = v1248
	goto L252
L251:
	;
	v1263 = v348
	goto L252
L252:
	;
	v1271 = v1248
	goto L253
L253:
	;
	if v1263 == v1271 {
		goto L255
	} else {
		goto L256
	}
L254:
	;
	v1301 = int32(-1)
	goto L214
L255:
	;
	v1311 = int32(0)
	goto L213
L256:
	;
	goto L257
L257:
	;
	v1275 = int32(1)
	v1280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v514+v1271<<(uint(v1275)%32)))))
	if v1280 == int32(0) {
		v1271 = v1271 + v1275
		goto L253
	} else {
		goto L258
	}
L258:
	;
	goto L254
L259:
	;
	goto L246
L260:
	;
	goto L135
L261:
	;
	goto L90
L262:
	;
	F_pfree(m, v1410)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L1
	} else {
		goto L265
	}
L263:
	;
	goto L264
L264:
	;
	v1414 = F_make_result_opt_error(m, v173, int32(0))
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L1
	} else {
		goto L266
	}
L265:
	;
	goto L264
L266:
	;
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v173)+16))
	if v1416 != 0 {
		goto L267
	} else {
		goto L268
	}
L267:
	;
	F_pfree(m, v1416)
	mBase = m.M
	v1418 = m.ExcPending
	if v1418 != 0 {
		goto L1
	} else {
		goto L270
	}
L268:
	;
	goto L269
L269:
	;
	m.G0 = v173 + int32(96)
	goto L34
L270:
	;
	goto L269
L271:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1428 = m.ExcPending
	if v1428 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	F_errmsg(m, int32(421948), int32(0))
	mBase = m.M
	v1432 = m.ExcPending
	if v1432 != 0 {
		goto L1
	} else {
		goto L273
	}
L273:
	;
	F_errfinish(m, int32(496807), int32(11702), int32(228247))
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L274
	}
L274:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L275:
	;
	F_errfinish(m, int32(496807), int32(4371), int32(487533))
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L277:
	;
	F_errfinish(m, int32(496807), int32(4360), int32(487533))
	mBase = m.M
	v1455 = m.ExcPending
	if v1455 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_numeric_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(1465)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+20)))
	if v8 == int32(1) {
		v11 = int32(4486928)
		v12 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*int32)(unsafe.Add(mBase, _consts[9])) = v14
		v17 = F_palloc(m, int32(48))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v22 = F_palloc(m, int32(132))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)) = uint8(v24)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v22
				F_initHyperLogLog(m, v17+int32(24), int32(10))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1466)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = int32(1467)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(1468)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v39
					*(*int32)(unsafe.Add(mBase, _consts[9])) = v12
					return int32(0)
				}
			}
		}
	} else {
		return int32(0)
	}
}
func F_numeric_sub_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
							v94 = F_make_result_opt_error(m, int32(1721372), int32(0))
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
							v98 = F_make_result_opt_error(m, int32(1721348), int32(0))
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
						if v47&int32(65535) == int32(53248) {
							v70 = F_make_result_opt_error(m, int32(1721324), int32(0))
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
							v74 = F_make_result_opt_error(m, int32(1721348), int32(0))
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
					if v47&int32(65535) == int32(61440) {
						v82 = F_make_result_opt_error(m, int32(1721324), int32(0))
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
						v86 = F_make_result_opt_error(m, int32(1721372), int32(0))
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
				v56 = F_make_result_opt_error(m, int32(1721324), int32(0))
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
			F_sub_var(m, v12+int32(56), v12+int32(32), v12+int32(8))
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
			v56 = F_make_result_opt_error(m, int32(1721324), int32(0))
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
							v94 = F_make_result_opt_error(m, int32(1721372), int32(0))
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
							v98 = F_make_result_opt_error(m, int32(1721348), int32(0))
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
						if v47&int32(65535) == int32(53248) {
							v70 = F_make_result_opt_error(m, int32(1721324), int32(0))
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
							v74 = F_make_result_opt_error(m, int32(1721348), int32(0))
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
					if v47&int32(65535) == int32(61440) {
						v82 = F_make_result_opt_error(m, int32(1721324), int32(0))
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
						v86 = F_make_result_opt_error(m, int32(1721372), int32(0))
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
				v56 = F_make_result_opt_error(m, int32(1721324), int32(0))
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
func F_numeric_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(457) {
		v43 = v2
		return v43
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v43 = v2
			return v43
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v43 = v2
				return v43
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					if int32(4) <= v22 {
						if v18 < int32(4) {
							v43 = v2
							return v43
						} else {
							if (v22^v18)&int32(2047) != 0 {
								v43 = v2
								return v43
							} else {
								v30 = int32(4)
								v32 = int32(16)
								if base.Ui32(int32(base.Ui32(v22-v30)>>(uint(v32)%32))) < base.Ui32(int32(base.Ui32(v18-v30)>>(uint(v32)%32))) {
									v43 = v2
									return v43
								} else {
									v39 = F_relabel_to_typmod(m, v17, v22)
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
										return int32(0)
									} else {
										v43 = v39
										return v43
									}
								}
							}
						}
					} else {
						v39 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v43 = v39
							return v43
						}
					}
				}
			}
		}
	}
}
func F_numeric_uminus(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v12 = F_palloc(m, int32(base.Ui32(v9)>>(uint(int32(2))%32)))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v16 != 0 {
				v17 = F__emscripten_memcpy_bulkmem(m, v12, v5, v16)
				mBase = m.M
				v18 = v17
			} else {
				v18 = v12
			}
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+4)))
			if base.Ui32(int32(49152)) <= base.Ui32(v19) {
				if v19 == int32(49152) {
					return v18
				} else {
					v25 = v19 ^ int32(8192)
					*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v25)
					return v18
				}
			} else {
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v33 = base.I32_extend16_s(v19)
				if int32(0) <= v33 {
					v36 = int32(-8)
				} else {
					v36 = int32(-6)
				}
				if base.Ui32(int32(base.Ui32(v28)>>(uint(int32(2))%32))+v36) < base.Ui32(int32(2)) {
					return v18
				} else {
					if v33 <= int32(-16385) {
						v43 = v19 ^ int32(8192)
						*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v43)
						return v18
					} else {
						if base.Ui32(v19) <= base.Ui32(int32(16383)) {
							v49 = v19 | int32(16384)
							*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v49)
							return v18
						} else {
							v53 = v19 & int32(16383)
							*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v53)
							return v18
						}
					}
				}
			}
		}
	}
}
