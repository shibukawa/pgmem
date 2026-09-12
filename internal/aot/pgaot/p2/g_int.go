package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_consistent(m *base.Module, l0 int32) int32 {
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
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
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
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum_copy(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v19)
		v22 = v18 & int32(65535)
		*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(base.B2i32(v22 == int32(6)))
		if v22 == int32(20) {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+16)))
			v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29+v30)+12)))
			v35 = F_execconsistent(m, v13, v28, v32&int32(1))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_pfree(m, v13)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v111 = v35
					m.G0 = v9 + int32(16)
					return v111 & int32(255)
				}
			}
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			if v39 != 0 {
				v40 = F_array_contains_nulls(m, v13)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					if v40 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(152400), int32(0))
								mBase = m.M
								v134 = m.ExcPending
								if v134 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(492055), int32(71), int32(91732))
									mBase = m.M
									v141 = m.ExcPending
									if v141 != 0 {
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
						v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v45 = F_ArrayGetNItems(m, v42, v13+int32(16))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							v47 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v47)
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
							if v49 != 0 {
								v57 = v49
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v57 = (v50<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							F_isort(m, v57+v13, v45, v9+int32(14))
							mBase = m.M
							v62 = F__int_unique(m, v13)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								switch v18&int32(65535) - int32(3) {
								case 0:
									v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
									v69 = F_inner_int_overlap(m, v68, v62)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v69)
										F_pfree(m, v62)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
											v111 = v110
											m.G0 = v9 + int32(16)
											return v111 & int32(255)
										}
									}
								default:
									v104 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v104)
									F_pfree(m, v62)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
										v111 = v110
										m.G0 = v9 + int32(16)
										return v111 & int32(255)
									}
								case 3:
									v72 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+16)))
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v74)+12)))
									if v76&int32(1) != 0 {
										v83 = F_DirectFunctionCall3Coll(m, int32(6382), int32(0), v72, v62, v9+int32(15))
										mBase = m.M
										v84 = m.ExcPending
										if v84 != 0 {
											return int32(0)
										} else {
											F_pfree(m, v62)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
												v111 = v110
												m.G0 = v9 + int32(16)
												return v111 & int32(255)
											}
										}
									} else {
										v85 = F_inner_int_contains(m, v72, v62)
										mBase = m.M
										v86 = m.ExcPending
										if v86 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v85)
											F_pfree(m, v62)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
												v111 = v110
												m.G0 = v9 + int32(16)
												return v111 & int32(255)
											}
										}
									}
								case 4, 10:
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
									v89 = F_inner_int_contains(m, v88, v62)
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v89)
										F_pfree(m, v62)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
											v111 = v110
											m.G0 = v9 + int32(16)
											return v111 & int32(255)
										}
									}
								case 5, 11:
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
									v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+16)))
									v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v93)+12)))
									if v95&int32(1) != 0 {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
										v99 = F_inner_int_contains(m, v62, v98)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v99)
											F_pfree(m, v62)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return int32(0)
											} else {
												v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
												v111 = v110
												m.G0 = v9 + int32(16)
												return v111 & int32(255)
											}
										}
									} else {
										v102 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v102)
										F_pfree(m, v62)
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return int32(0)
										} else {
											v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
											v111 = v110
											m.G0 = v9 + int32(16)
											return v111 & int32(255)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
				v45 = F_ArrayGetNItems(m, v42, v13+int32(16))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					v47 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+14)) = uint8(v47)
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
					if v49 != 0 {
						v57 = v49
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
						v57 = (v50<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					}
					F_isort(m, v57+v13, v45, v9+int32(14))
					mBase = m.M
					v62 = F__int_unique(m, v13)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						switch v18&int32(65535) - int32(3) {
						case 0:
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v69 = F_inner_int_overlap(m, v68, v62)
							mBase = m.M
							v70 = m.ExcPending
							if v70 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v69)
								F_pfree(m, v62)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
									v111 = v110
									m.G0 = v9 + int32(16)
									return v111 & int32(255)
								}
							}
						default:
							v104 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v104)
							F_pfree(m, v62)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
								v111 = v110
								m.G0 = v9 + int32(16)
								return v111 & int32(255)
							}
						case 3:
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v73)+16)))
							v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73+v74)+12)))
							if v76&int32(1) != 0 {
								v83 = F_DirectFunctionCall3Coll(m, int32(6382), int32(0), v72, v62, v9+int32(15))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v62)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
										v111 = v110
										m.G0 = v9 + int32(16)
										return v111 & int32(255)
									}
								}
							} else {
								v85 = F_inner_int_contains(m, v72, v62)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v85)
									F_pfree(m, v62)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
										v111 = v110
										m.G0 = v9 + int32(16)
										return v111 & int32(255)
									}
								}
							}
						case 4, 10:
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
							v89 = F_inner_int_contains(m, v88, v62)
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v89)
								F_pfree(m, v62)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
									v111 = v110
									m.G0 = v9 + int32(16)
									return v111 & int32(255)
								}
							}
						case 5, 11:
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v92)+16)))
							v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92+v93)+12)))
							if v95&int32(1) != 0 {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
								v99 = F_inner_int_contains(m, v62, v98)
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v99)
									F_pfree(m, v62)
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return int32(0)
									} else {
										v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
										v111 = v110
										m.G0 = v9 + int32(16)
										return v111 & int32(255)
									}
								}
							} else {
								v102 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)) = uint8(v102)
								F_pfree(m, v62)
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int32(0)
								} else {
									v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
									v111 = v110
									m.G0 = v9 + int32(16)
									return v111 & int32(255)
								}
							}
						}
					}
				}
			}
		}
	}
}
