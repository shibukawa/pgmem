package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_consistent(m *base.Module, l0 int32) int64 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int64
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_copy(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v20)
		*(*uint8)(unsafe.Add(mBase, uint32(v18))) = uint8(base.B2i32(v19 == int32(6)))
		if v19 == int32(20) {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+16)))
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28+v29)+12)))
			v34 = F_execconsistent(m, v14, v27, v31&int32(1))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				F_pfree(m, v14)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int64(0)
				} else {
					v111 = v34
					m.G0 = v10 + int32(16)
					return base.I64_extend_i32_u(v111) & int64(255)
				}
			}
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
			if v38 != 0 {
				v39 = F_array_contains_nulls(m, v14)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					if v39 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_g_int_consistent_0), int32(0))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_g_int_consistent_1), int32(72), int32(_a_F_g_int_consistent_2))
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
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
						v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						v44 = F_ArrayGetNItemsSafe(m, v41, v14+int32(16))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							v46 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v46)
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							if v48 != 0 {
								v56 = v48
							} else {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
								v56 = (v49<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							F_isort(m, v56+v14, v44, v10+int32(14))
							mBase = m.M
							v61 = F__int_unique(m, v14)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int64(0)
							} else {
								switch v19 - int32(3) {
								case 0:
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v66 = F_inner_int_overlap(m, v65, v61)
									mBase = m.M
									v67 = m.ExcPending
									if v67 != 0 {
										return int64(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v66)
										F_pfree(m, v61)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int64(0)
										} else {
											v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											v111 = v110
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v111) & int64(255)
										}
									}
								default:
									v104 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v104)
									F_pfree(m, v61)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										v111 = v110
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v111) & int64(255)
									}
								case 3:
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+16)))
									v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v71)+12)))
									if v73&int32(1) != 0 {
										v82 = F_DirectFunctionCall3Coll(m, int32(_a_F_g_int_consistent_3), int32(0), v69, base.I64_extend_i32_u(v61), base.I64_extend_i32_u(v10+int32(15)))
										mBase = m.M
										v83 = m.ExcPending
										if v83 != 0 {
											return int64(0)
										} else {
											F_pfree(m, v61)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												v111 = v110
												m.G0 = v10 + int32(16)
												return base.I64_extend_i32_u(v111) & int64(255)
											}
										}
									} else {
										v85 = F_inner_int_contains(m, base.I32_wrap_i64(v69), v61)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int64(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v85)
											F_pfree(m, v61)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												v111 = v110
												m.G0 = v10 + int32(16)
												return base.I64_extend_i32_u(v111) & int64(255)
											}
										}
									}
								case 4, 10:
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
									v89 = F_inner_int_contains(m, v88, v61)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int64(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v89)
										F_pfree(m, v61)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int64(0)
										} else {
											v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											v111 = v110
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v111) & int64(255)
										}
									}
								case 5, 11:
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
									v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+16)))
									v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v93)+12)))
									if v95&int32(1) != 0 {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										v99 = F_inner_int_contains(m, v61, v98)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int64(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v99)
											F_pfree(m, v61)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int64(0)
											} else {
												v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
												v111 = v110
												m.G0 = v10 + int32(16)
												return base.I64_extend_i32_u(v111) & int64(255)
											}
										}
									} else {
										v102 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v102)
										F_pfree(m, v61)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int64(0)
										} else {
											v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
											v111 = v110
											m.G0 = v10 + int32(16)
											return base.I64_extend_i32_u(v111) & int64(255)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v41 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				v44 = F_ArrayGetNItemsSafe(m, v41, v14+int32(16))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int64(0)
				} else {
					v46 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)) = uint8(v46)
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
					if v48 != 0 {
						v56 = v48
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
						v56 = (v49<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					}
					F_isort(m, v56+v14, v44, v10+int32(14))
					mBase = m.M
					v61 = F__int_unique(m, v14)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int64(0)
					} else {
						switch v19 - int32(3) {
						case 0:
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v66 = F_inner_int_overlap(m, v65, v61)
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int64(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v66)
								F_pfree(m, v61)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
								} else {
									v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									v111 = v110
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v111) & int64(255)
								}
							}
						default:
							v104 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v104)
							F_pfree(m, v61)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
							} else {
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
								v111 = v110
								m.G0 = v10 + int32(16)
								return base.I64_extend_i32_u(v111) & int64(255)
							}
						case 3:
							v69 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+16)))
							v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70+v71)+12)))
							if v73&int32(1) != 0 {
								v82 = F_DirectFunctionCall3Coll(m, int32(_a_F_g_int_consistent_3), int32(0), v69, base.I64_extend_i32_u(v61), base.I64_extend_i32_u(v10+int32(15)))
								mBase = m.M
								v83 = m.ExcPending
								if v83 != 0 {
									return int64(0)
								} else {
									F_pfree(m, v61)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										v111 = v110
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v111) & int64(255)
									}
								}
							} else {
								v85 = F_inner_int_contains(m, base.I32_wrap_i64(v69), v61)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int64(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v85)
									F_pfree(m, v61)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										v111 = v110
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v111) & int64(255)
									}
								}
							}
						case 4, 10:
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v89 = F_inner_int_contains(m, v88, v61)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int64(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v89)
								F_pfree(m, v61)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
								} else {
									v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									v111 = v110
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v111) & int64(255)
								}
							}
						case 5, 11:
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+16)))
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v93)+12)))
							if v95&int32(1) != 0 {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
								v99 = F_inner_int_contains(m, v61, v98)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int64(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v99)
									F_pfree(m, v61)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int64(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
										v111 = v110
										m.G0 = v10 + int32(16)
										return base.I64_extend_i32_u(v111) & int64(255)
									}
								}
							} else {
								v102 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)) = uint8(v102)
								F_pfree(m, v61)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
								} else {
									v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
									v111 = v110
									m.G0 = v10 + int32(16)
									return base.I64_extend_i32_u(v111) & int64(255)
								}
							}
						}
					}
				}
			}
		}
	}
}
