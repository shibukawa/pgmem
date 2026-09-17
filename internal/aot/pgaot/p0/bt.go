package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_blk_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui32(v4) < base.Ui32(v3)) - base.B2i32(base.Ui32(v3) < base.Ui32(v4))
}
func F__bt_build_callback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	if l4 == int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
		if v9 != 0 {
			v12 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l5)+2)) = uint8(v12)
			v14 = v9
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
			v14 = v11
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
		v14 = v11
	}
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	F_tuplesort_putindextuplevalues(m, v15, v16, l1, l2, l3)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return
	} else {
		v19 = *(*float64)(unsafe.Add(mBase, uint32(l5)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l5)+16)) = base.F64_add(v19, float64(1))
		return
	}
}
func F__bt_checkkeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+80))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	v23 = l1 + int32(24)
	v26 = F__bt_check_compare(m, l0, v16, l3, l4, v18, l2, v21, v23, v13+int32(12))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		if l2 == int32(0) {
			v127 = v26
			m.G0 = v13 + int32(16)
			return v127
		} else {
			v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
			if v32 != 0 {
				v127 = v26
				m.G0 = v13 + int32(16)
				return v127
			} else {
				v33 = int32(0)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				v37 = F__bt_tuple_before_array_skeys(m, l0, v16, l3, v18, l4, int32(1), v35, v33)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					if v37 != 0 {
						v39 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v39)
						v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)))
						v43 = v41 + v39
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v43)
						if base.I32_extend16_s(v43) < int32(3) {
							v127 = v33
							m.G0 = v13 + int32(16)
							return v127
						} else {
							v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)))
							v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
							if base.Ui32(v48) < base.Ui32(v49) {
								v127 = v33
								m.G0 = v13 + int32(16)
								return v127
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+80))
								if v52 == int32(1) {
									v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
									if v48 < v55-int32(5) {
										v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+28)))
										if v64 != 0 {
											if int32(203) < v64 {
												v72 = v64
											} else {
												v70 = v64 << (uint(int32(1)) % 32)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v70)
												v72 = v70
											}
										} else {
											v70 = int32(5)
											*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v70)
											v72 = v70
										}
										if v52 == int32(1) {
											v76 = base.I32_extend16_s(v72) + v48
											v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
											if v76 < v77 {
												v79 = v76
											} else {
												v79 = v77
											}
											v86 = v79
										} else {
											v81 = v48 - base.I32_extend16_s(v72)
											if v49 < v81 {
												v83 = v81
											} else {
												v83 = v49
											}
											v86 = v83
										}
										v87 = int32(0)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+v86&int32(_a_F__bt_checkkeys_0)<<(uint(int32(2))%32))+20))
										v101 = F__bt_tuple_before_array_skeys(m, l0, v52, v88+v94&int32(_a_F__bt_checkkeys_1), v18, l4, v87, v87, v87)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											if v101 != 0 {
												if v52 == int32(1) {
													v106 = v86 + int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v106)
													v127 = v87
												} else {
													v109 = v86 - int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v109)
													v127 = v87
												}
											} else {
												v111 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v111)
												v113 = int32(15)
												v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+28)))
												if v114 <= v113 {
													v117 = v113
												} else {
													v117 = v114
												}
												v119 = int32(base.Ui32(v117) >> (uint(int32(3)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v119)
												v127 = v87
											}
											m.G0 = v13 + int32(16)
											return v127
										}
									} else {
										v127 = v33
										m.G0 = v13 + int32(16)
										return v127
									}
								} else {
									if v52 != int32(-1) {
										v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+28)))
										if v64 != 0 {
											if int32(203) < v64 {
												v72 = v64
											} else {
												v70 = v64 << (uint(int32(1)) % 32)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v70)
												v72 = v70
											}
										} else {
											v70 = int32(5)
											*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v70)
											v72 = v70
										}
										if v52 == int32(1) {
											v76 = base.I32_extend16_s(v72) + v48
											v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
											if v76 < v77 {
												v79 = v76
											} else {
												v79 = v77
											}
											v86 = v79
										} else {
											v81 = v48 - base.I32_extend16_s(v72)
											if v49 < v81 {
												v83 = v81
											} else {
												v83 = v49
											}
											v86 = v83
										}
										v87 = int32(0)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+v86&int32(_a_F__bt_checkkeys_0)<<(uint(int32(2))%32))+20))
										v101 = F__bt_tuple_before_array_skeys(m, l0, v52, v88+v94&int32(_a_F__bt_checkkeys_1), v18, l4, v87, v87, v87)
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return int32(0)
										} else {
											if v101 != 0 {
												if v52 == int32(1) {
													v106 = v86 + int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v106)
													v127 = v87
												} else {
													v109 = v86 - int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v109)
													v127 = v87
												}
											} else {
												v111 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v111)
												v113 = int32(15)
												v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+28)))
												if v114 <= v113 {
													v117 = v113
												} else {
													v117 = v114
												}
												v119 = int32(base.Ui32(v117) >> (uint(int32(3)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v119)
												v127 = v87
											}
											m.G0 = v13 + int32(16)
											return v127
										}
									} else {
										if base.Ui32(v48) <= base.Ui32(v49+int32(5)) {
											v127 = v33
											m.G0 = v13 + int32(16)
											return v127
										} else {
											v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+28)))
											if v64 != 0 {
												if int32(203) < v64 {
													v72 = v64
												} else {
													v70 = v64 << (uint(int32(1)) % 32)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v70)
													v72 = v70
												}
											} else {
												v70 = int32(5)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v70)
												v72 = v70
											}
											if v52 == int32(1) {
												v76 = base.I32_extend16_s(v72) + v48
												v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
												if v76 < v77 {
													v79 = v76
												} else {
													v79 = v77
												}
												v86 = v79
											} else {
												v81 = v48 - base.I32_extend16_s(v72)
												if v49 < v81 {
													v83 = v81
												} else {
													v83 = v49
												}
												v86 = v83
											}
											v87 = int32(0)
											v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+v86&int32(_a_F__bt_checkkeys_0)<<(uint(int32(2))%32))+20))
											v101 = F__bt_tuple_before_array_skeys(m, l0, v52, v88+v94&int32(_a_F__bt_checkkeys_1), v18, l4, v87, v87, v87)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												if v101 != 0 {
													if v52 == int32(1) {
														v106 = v86 + int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v106)
														v127 = v87
													} else {
														v109 = v86 - int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(l1)+22)) = uint16(v109)
														v127 = v87
													}
												} else {
													v111 = int32(0)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v111)
													v113 = int32(15)
													v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+28)))
													if v114 <= v113 {
														v117 = v113
													} else {
														v117 = v114
													}
													v119 = int32(base.Ui32(v117) >> (uint(int32(3)) % 32))
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+28)) = uint16(v119)
													v127 = v87
												}
												m.G0 = v13 + int32(16)
												return v127
											}
										}
									}
								}
							}
						}
					} else {
						v122 = F__bt_advance_array_keys(m, l0, l1, l3, l4, v18, v35, int32(1))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							v127 = v122
							m.G0 = v13 + int32(16)
							return v127
						}
					}
				}
			}
		}
	}
}
func F__bt_compare_scankey_args(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
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
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
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
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v19 = v17 | v18
	if v19&int32(1) != 0 {
		if v19&int32(_a_F__bt_compare_scankey_args_0) != 0 {
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l4)+19)) = uint8(v24)
			v178 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
			v196 = v178
			m.G0 = v15 + int32(16)
			return v196
		} else {
			v27 = int32(1)
			v29 = v17 & v27
			v31 = v18 & v27
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
			if v35&int32(2) != 0 {
				v38 = int32(6) - v33
			} else {
				v38 = v33
			}
			v40 = v38 & int32(_a_F__bt_compare_scankey_args_1)
			switch v40 - int32(1) {
			case 0:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(base.Ui32(v31) < base.Ui32(v29)))
				v196 = v27
				m.G0 = v15 + int32(16)
				return v196
			case 1:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(base.Ui32(v31) <= base.Ui32(v29)))
				v196 = v27
				m.G0 = v15 + int32(16)
				return v196
			case 2:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v29 == v31))
				v196 = v27
				m.G0 = v15 + int32(16)
				return v196
			case 3:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(base.Ui32(v29) <= base.Ui32(v31)))
				v196 = v27
				m.G0 = v15 + int32(16)
				return v196
			case 4:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(base.Ui32(v29) < base.Ui32(v31)))
				v196 = v27
				m.G0 = v15 + int32(16)
				return v196
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v40
					F_errmsg_internal(m, int32(_a_F__bt_compare_scankey_args_2), v15)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F__bt_compare_scankey_args_3), int32(950), int32(_a_F__bt_compare_scankey_args_4))
						mBase = m.M
						v67 = m.ExcPending
						if v67 != 0 {
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
		if v19&int32(4) != 0 {
			v196 = int32(0)
			m.G0 = v15 + int32(16)
			return v196
		} else {
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if l4 == int32(0) {
				v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
				v111 = v109 << (uint(int32(2)) % 32)
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
				v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112-int32(4))))
				if v108 != 0 {
					v117 = v108
				} else {
					v117 = v116
				}
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				if v118 != 0 {
					v119 = v118
				} else {
					v119 = v116
				}
				if v119 != v116 {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
					v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v144&int32(16777216) != 0 {
						v147 = int32(6) - v142
					} else {
						v147 = v142
					}
					v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return int32(0)
					} else {
						if v149 == int32(0) {
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
							v196 = v178
							m.G0 = v15 + int32(16)
							return v196
						} else {
							v153 = F_get_opcode(m, v149)
							mBase = m.M
							v154 = m.ExcPending
							if v154 != 0 {
								return int32(0)
							} else {
								if v153 == int32(0) {
									v178 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
									v196 = v178
									m.G0 = v15 + int32(16)
									return v196
								} else {
									v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
									v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
									v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
										v196 = int32(1)
										m.G0 = v15 + int32(16)
										return v196
									}
								}
							}
						}
					}
				} else {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v121 != 0 {
						v122 = v121
					} else {
						v122 = v116
					}
					if v117 != v122 {
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
						v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
						v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v144&int32(16777216) != 0 {
							v147 = int32(6) - v142
						} else {
							v147 = v142
						}
						v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
						mBase = m.M
						v150 = m.ExcPending
						if v150 != 0 {
							return int32(0)
						} else {
							if v149 == int32(0) {
								v178 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
								v196 = v178
								m.G0 = v15 + int32(16)
								return v196
							} else {
								v153 = F_get_opcode(m, v149)
								mBase = m.M
								v154 = m.ExcPending
								if v154 != 0 {
									return int32(0)
								} else {
									if v153 == int32(0) {
										v178 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
										v196 = v178
										m.G0 = v15 + int32(16)
										return v196
									} else {
										v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
										v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
										v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
											v196 = int32(1)
											m.G0 = v15 + int32(16)
											return v196
										}
									}
								}
							}
						}
					} else {
						v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
						v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
						v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int32(0)))
							v196 = int32(1)
							m.G0 = v15 + int32(16)
							return v196
						}
					}
				}
			} else {
				if v18&int32(32) != 0 {
					v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
					if v17&int32(32) == int32(0) {
						if v75 != int32(3) {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
							v111 = v109 << (uint(int32(2)) % 32)
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112-int32(4))))
							if v108 != 0 {
								v117 = v108
							} else {
								v117 = v116
							}
							v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v118 != 0 {
								v119 = v118
							} else {
								v119 = v116
							}
							if v119 != v116 {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
								v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v144&int32(16777216) != 0 {
									v147 = int32(6) - v142
								} else {
									v147 = v142
								}
								v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									if v149 == int32(0) {
										v178 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
										v196 = v178
										m.G0 = v15 + int32(16)
										return v196
									} else {
										v153 = F_get_opcode(m, v149)
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											if v153 == int32(0) {
												v178 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
												v196 = v178
												m.G0 = v15 + int32(16)
												return v196
											} else {
												v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
												v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
												v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
													v196 = int32(1)
													m.G0 = v15 + int32(16)
													return v196
												}
											}
										}
									}
								}
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v121 != 0 {
									v122 = v121
								} else {
									v122 = v116
								}
								if v117 != v122 {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
									v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
									v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
									v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v144&int32(16777216) != 0 {
										v147 = int32(6) - v142
									} else {
										v147 = v142
									}
									v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int32(0)
									} else {
										if v149 == int32(0) {
											v178 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
											v196 = v178
											m.G0 = v15 + int32(16)
											return v196
										} else {
											v153 = F_get_opcode(m, v149)
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return int32(0)
											} else {
												if v153 == int32(0) {
													v178 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
													v196 = v178
													m.G0 = v15 + int32(16)
													return v196
												} else {
													v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
													v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
													v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
													v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
													mBase = m.M
													v161 = m.ExcPending
													if v161 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
														v196 = int32(1)
														m.G0 = v15 + int32(16)
														return v196
													}
												}
											}
										}
									}
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
									v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
									v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int32(0)))
										v196 = int32(1)
										m.G0 = v15 + int32(16)
										return v196
									}
								}
							}
						} else {
							v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							if v99 != int32(-1) {
								v102 = F__bt_saoparray_shrink(m, l0, l2, l3, l5, l4, l6)
								mBase = m.M
								v103 = m.ExcPending
								if v103 != 0 {
									return int32(0)
								} else {
									v196 = v102
									m.G0 = v15 + int32(16)
									return v196
								}
							} else {
								v104 = F__bt_skiparray_shrink(m, l0, l3, l4, l6)
								mBase = m.M
								v105 = m.ExcPending
								if v105 != 0 {
									return int32(0)
								} else {
									v196 = v104
									m.G0 = v15 + int32(16)
									return v196
								}
							}
						}
					} else {
						v80 = int32(3)
						v81 = base.B2i32(v75 == v80)
						v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
						if v81&base.B2i32(v82 == v80) != 0 {
							v178 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
							v196 = v178
							m.G0 = v15 + int32(16)
							return v196
						} else {
							if v75 == v80 {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
								if v99 != int32(-1) {
									v102 = F__bt_saoparray_shrink(m, l0, l2, l3, l5, l4, l6)
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int32(0)
									} else {
										v196 = v102
										m.G0 = v15 + int32(16)
										return v196
									}
								} else {
									v104 = F__bt_skiparray_shrink(m, l0, l3, l4, l6)
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										v196 = v104
										m.G0 = v15 + int32(16)
										return v196
									}
								}
							} else {
								if v82 != int32(3) {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
									v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
									v111 = v109 << (uint(int32(2)) % 32)
									v112 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
									v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112-int32(4))))
									if v108 != 0 {
										v117 = v108
									} else {
										v117 = v116
									}
									v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v118 != 0 {
										v119 = v118
									} else {
										v119 = v116
									}
									if v119 != v116 {
										v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
										v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
										v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
										v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										if v144&int32(16777216) != 0 {
											v147 = int32(6) - v142
										} else {
											v147 = v142
										}
										v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
										mBase = m.M
										v150 = m.ExcPending
										if v150 != 0 {
											return int32(0)
										} else {
											if v149 == int32(0) {
												v178 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
												v196 = v178
												m.G0 = v15 + int32(16)
												return v196
											} else {
												v153 = F_get_opcode(m, v149)
												mBase = m.M
												v154 = m.ExcPending
												if v154 != 0 {
													return int32(0)
												} else {
													if v153 == int32(0) {
														v178 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
														v196 = v178
														m.G0 = v15 + int32(16)
														return v196
													} else {
														v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
														v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
														v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
														v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
															v196 = int32(1)
															m.G0 = v15 + int32(16)
															return v196
														}
													}
												}
											}
										}
									} else {
										v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										if v121 != 0 {
											v122 = v121
										} else {
											v122 = v116
										}
										if v117 != v122 {
											v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
											v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
											v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
											v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											if v144&int32(16777216) != 0 {
												v147 = int32(6) - v142
											} else {
												v147 = v142
											}
											v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
											mBase = m.M
											v150 = m.ExcPending
											if v150 != 0 {
												return int32(0)
											} else {
												if v149 == int32(0) {
													v178 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
													v196 = v178
													m.G0 = v15 + int32(16)
													return v196
												} else {
													v153 = F_get_opcode(m, v149)
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return int32(0)
													} else {
														if v153 == int32(0) {
															v178 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
															v196 = v178
															m.G0 = v15 + int32(16)
															return v196
														} else {
															v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
															v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
															v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
															v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return int32(0)
															} else {
																*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
																v196 = int32(1)
																m.G0 = v15 + int32(16)
																return v196
															}
														}
													}
												}
											}
										} else {
											v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
											v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
											v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int32(0)))
												v196 = int32(1)
												m.G0 = v15 + int32(16)
												return v196
											}
										}
									}
								} else {
									v182 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
									if v182 != int32(-1) {
										v185 = F__bt_saoparray_shrink(m, l0, l3, l2, l5, l4, l6)
										mBase = m.M
										v186 = m.ExcPending
										if v186 != 0 {
											return int32(0)
										} else {
											v189 = v185
											v196 = v189
											m.G0 = v15 + int32(16)
											return v196
										}
									} else {
										v187 = F__bt_skiparray_shrink(m, l0, l2, l4, l6)
										mBase = m.M
										v188 = m.ExcPending
										if v188 != 0 {
											return int32(0)
										} else {
											v189 = v187
											v196 = v189
											m.G0 = v15 + int32(16)
											return v196
										}
									}
								}
							}
						}
					}
				} else {
					if v17&int32(32) == int32(0) {
						v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
						v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
						v111 = v109 << (uint(int32(2)) % 32)
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112-int32(4))))
						if v108 != 0 {
							v117 = v108
						} else {
							v117 = v116
						}
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						if v118 != 0 {
							v119 = v118
						} else {
							v119 = v116
						}
						if v119 != v116 {
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
							v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v144&int32(16777216) != 0 {
								v147 = int32(6) - v142
							} else {
								v147 = v142
							}
							v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								if v149 == int32(0) {
									v178 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
									v196 = v178
									m.G0 = v15 + int32(16)
									return v196
								} else {
									v153 = F_get_opcode(m, v149)
									mBase = m.M
									v154 = m.ExcPending
									if v154 != 0 {
										return int32(0)
									} else {
										if v153 == int32(0) {
											v178 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
											v196 = v178
											m.G0 = v15 + int32(16)
											return v196
										} else {
											v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
											v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
												v196 = int32(1)
												m.G0 = v15 + int32(16)
												return v196
											}
										}
									}
								}
							}
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v121 != 0 {
								v122 = v121
							} else {
								v122 = v116
							}
							if v117 != v122 {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
								v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v144&int32(16777216) != 0 {
									v147 = int32(6) - v142
								} else {
									v147 = v142
								}
								v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									if v149 == int32(0) {
										v178 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
										v196 = v178
										m.G0 = v15 + int32(16)
										return v196
									} else {
										v153 = F_get_opcode(m, v149)
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											if v153 == int32(0) {
												v178 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
												v196 = v178
												m.G0 = v15 + int32(16)
												return v196
											} else {
												v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
												v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
												v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
													v196 = int32(1)
													m.G0 = v15 + int32(16)
													return v196
												}
											}
										}
									}
								}
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
								v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int32(0)))
									v196 = int32(1)
									m.G0 = v15 + int32(16)
									return v196
								}
							}
						}
					} else {
						v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
						if v92 == int32(3) {
							v182 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							if v182 != int32(-1) {
								v185 = F__bt_saoparray_shrink(m, l0, l3, l2, l5, l4, l6)
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									v189 = v185
									v196 = v189
									m.G0 = v15 + int32(16)
									return v196
								}
							} else {
								v187 = F__bt_skiparray_shrink(m, l0, l2, l4, l6)
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return int32(0)
								} else {
									v189 = v187
									v196 = v189
									m.G0 = v15 + int32(16)
									return v196
								}
							}
						} else {
							v108 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							v109 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
							v111 = v109 << (uint(int32(2)) % 32)
							v112 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v111+v112-int32(4))))
							if v108 != 0 {
								v117 = v108
							} else {
								v117 = v116
							}
							v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v118 != 0 {
								v119 = v118
							} else {
								v119 = v116
							}
							if v119 != v116 {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
								v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v144&int32(16777216) != 0 {
									v147 = int32(6) - v142
								} else {
									v147 = v142
								}
								v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
								mBase = m.M
								v150 = m.ExcPending
								if v150 != 0 {
									return int32(0)
								} else {
									if v149 == int32(0) {
										v178 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
										v196 = v178
										m.G0 = v15 + int32(16)
										return v196
									} else {
										v153 = F_get_opcode(m, v149)
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											if v153 == int32(0) {
												v178 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
												v196 = v178
												m.G0 = v15 + int32(16)
												return v196
											} else {
												v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
												v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
												v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
												v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
													v196 = int32(1)
													m.G0 = v15 + int32(16)
													return v196
												}
											}
										}
									}
								}
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v121 != 0 {
									v122 = v121
								} else {
									v122 = v116
								}
								if v117 != v122 {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
									v140 = *(*int32)(unsafe.Add(mBase, uint32(v136+v111-int32(4))))
									v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
									v144 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v144&int32(16777216) != 0 {
										v147 = int32(6) - v142
									} else {
										v147 = v142
									}
									v149 = F_get_opfamily_member(m, v140, v119, v117, base.I32_extend16_s(v147))
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return int32(0)
									} else {
										if v149 == int32(0) {
											v178 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
											v196 = v178
											m.G0 = v15 + int32(16)
											return v196
										} else {
											v153 = F_get_opcode(m, v149)
											mBase = m.M
											v154 = m.ExcPending
											if v154 != 0 {
												return int32(0)
											} else {
												if v153 == int32(0) {
													v178 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v178)
													v196 = v178
													m.G0 = v15 + int32(16)
													return v196
												} else {
													v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
													v158 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
													v159 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
													v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
													mBase = m.M
													v161 = m.ExcPending
													if v161 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int32(0)))
														v196 = int32(1)
														m.G0 = v15 + int32(16)
														return v196
													}
												}
											}
										}
									}
								} else {
									v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
									v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
									v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int32(0)))
										v196 = int32(1)
										m.G0 = v15 + int32(16)
										return v196
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
func F__bt_dedup_finish_pending(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v11 = int32(1)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v12) < base.Ui32(int32(25)) {
		v21 = v11
	} else {
		v21 = int32(base.Ui32(v12+int32(_a_F__bt_dedup_finish_pending_0))>>(uint(int32(2))%32)) + v11
	}
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v23 == int32(1) {
		v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)))
		v32 = F_PageAddItemExtended(m, l0, v22, v26&int32(_a_F__bt_dedup_finish_pending_1), v21&int32(_a_F__bt_dedup_finish_pending_2), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return
		} else {
			if v32 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = int64(0)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(_a_F__bt_dedup_finish_pending_3), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F__bt_dedup_finish_pending_4), int32(574), int32(_a_F__bt_dedup_finish_pending_5))
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
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+6)))
		if v49&int32(_a_F__bt_dedup_finish_pending_6) == int32(0) {
			v66 = v49 & int32(_a_F__bt_dedup_finish_pending_1)
		} else {
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+5)))
			if v54&int32(32) == int32(0) {
				v66 = v49 & int32(_a_F__bt_dedup_finish_pending_1)
			} else {
				v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+2)))
				v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22))))
				v66 = v59 | v60<<(uint(int32(16))%32)
			}
		}
		v68 = v47 * int32(6)
		if int32(1) < v47 {
			v76 = (v66 + v68 + int32(7)) & int32(-8)
		} else {
			v76 = v66
		}
		v77 = F_palloc0(m, v76)
		mBase = m.M
		v78 = m.ExcPending
		if v78 != 0 {
			return
		} else {
			if v66 != 0 {
				base.MemoryCopy(m, v77, v22, v66)
			} else {
			}
			v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)))
			v83 = v80&int32(-8192) | v76
			if int32(2) <= v47 {
				*(*uint16)(unsafe.Add(mBase, uint32(v77)+2)) = uint16(v66)
				v87 = int32(_a_F__bt_dedup_finish_pending_6)
				v88 = v47 | v87
				*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v88)
				v91 = v83 | v87
				*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v91)
				v94 = int32(base.Ui32(v66) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v77))) = uint16(v94)
				if v68 != 0 {
					base.MemoryCopy(m, v77+v66&int32(-65536)+v66&int32(_a_F__bt_dedup_finish_pending_2), v48, v68)
				} else {
				}
				v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)))
				v111 = v103
			} else {
				v105 = v83 & int32(_a_F__bt_dedup_finish_pending_7)
				*(*uint16)(unsafe.Add(mBase, uint32(v77)+6)) = uint16(v105)
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				*(*int32)(unsafe.Add(mBase, uint32(v77))) = v107
				v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v77)+4)) = uint16(v109)
				v111 = v105
			}
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v112<<(uint(int32(2))%32))+46)) = uint16(v116)
			v123 = F_PageAddItemExtended(m, l0, v77, v111&int32(_a_F__bt_dedup_finish_pending_1), v21&int32(_a_F__bt_dedup_finish_pending_2), int32(0))
			mBase = m.M
			v124 = m.ExcPending
			if v124 != 0 {
				return
			} else {
				if v123 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v148 = m.ExcPending
					if v148 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(_a_F__bt_dedup_finish_pending_3), int32(0))
						mBase = m.M
						v152 = m.ExcPending
						if v152 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F__bt_dedup_finish_pending_4), int32(594), int32(_a_F__bt_dedup_finish_pending_5))
							mBase = m.M
							v157 = m.ExcPending
							if v157 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					F_pfree(m, v77)
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return
					} else {
						v129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v129 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = int64(0)
						return
					}
				}
			}
		}
	}
}
func F__bt_dedup_save_htid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	v8 = int32(1)
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v9&int32(32) == int32(0) {
		v27 = v8
		v29 = l1
	} else {
		v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v14&int32(_a_F__bt_dedup_save_htid_0) == int32(0) {
			v27 = v8
			v29 = l1
		} else {
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v27 = v14 & int32(4095)
			v29 = v21 + (l1 + v22<<(uint(int32(16))%32))
		}
	}
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v41 = base.B2i32(base.Ui32((v31+(v32+v27)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v30))
	if v41 == int32(0) {
		if v32 <= int32(50) {
		} else {
			v72 = int32(4)
			v74 = int32(1)
			v75 = l0 + v72
			v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
			*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76 + v74
		}
	} else {
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v48 + int32(1)
		v53 = v27 * int32(6)
		if v53 != 0 {
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			base.MemoryCopy(m, v54+v32*int32(6), v29, v53)
		} else {
		}
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v59 + v27
		v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
		v72 = int32(36)
		v74 = (v63&int32(_a_F__bt_dedup_save_htid_1)+int32(7))&int32(_a_F__bt_dedup_save_htid_2) | int32(4)
		v75 = l0 + v72
		v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
		*(*int32)(unsafe.Add(mBase, uint32(v75))) = v76 + v74
	}
	return v41
}
func F__bt_finish_split(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 < v5 {
		v21 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[0]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l2^int32(-1))<<(uint(int32(2))%32))))
		v35 = v27
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[1]))
		v35 = v29 + l2<<(uint(int32(13))%32) + int32(-8192)
	}
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)))
	v37 = v36 + v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v40 = F__bt_getbuf(m, l0, v38, int32(2))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		return
	} else {
		if v40 < int32(0) {
			v45 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[0]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v45+(v40^int32(-1))<<(uint(int32(2))%32))))
			v59 = v51
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[1]))
			v59 = v53 + v40<<(uint(int32(13))%32) + int32(-8192)
		}
		v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+16)))
		if l3 == int32(0) {
			v65 = F__bt_getbuf(m, l0, int32(0), int32(2))
			mBase = m.M
			v66 = m.ExcPending
			if v66 != 0 {
				return
			} else {
				if v65 < int32(0) {
					v70 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[0]))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v65^int32(-1))<<(uint(int32(2))%32))))
					v84 = v76
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[1]))
					v84 = v78 + v65<<(uint(int32(13))%32) + int32(-8192)
				}
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+32))
				if l2 < int32(0) {
					v89 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[2]))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v89+(l2^int32(-1))<<(uint(int32(6))%32))+16))
					v104 = v95
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v97+l2<<(uint(int32(6))%32)+int32(-64))+16))
					v104 = v103
				}
				F__bt_relbuf(m, v65)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return
				} else {
					v108 = base.B2i32(v85 == v104)
					v109 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					if v109 == int32(0) {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v59+v60)+4))
						v116 = base.B2i32(v113 == int32(0))
					} else {
						v116 = v5
					}
					v119 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v120 = m.ExcPending
					if v120 != 0 {
						return
					} else {
						if v119 != 0 {
							if l2 < int32(0) {
								v124 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[2]))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+(l2^int32(-1))<<(uint(int32(6))%32))+16))
								v139 = v130
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v132+l2<<(uint(int32(6))%32)+int32(-64))+16))
								v139 = v138
							}
							if v40 < int32(0) {
								v143 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[2]))
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+(v40^int32(-1))<<(uint(int32(6))%32))+16))
								v158 = v149
							} else {
								v151 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v151+v40<<(uint(int32(6))%32)+int32(-64))+16))
								v158 = v157
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v158
							*(*int32)(unsafe.Add(mBase, uint32(v16))) = v139
							F_errmsg_internal(m, int32(_a_F__bt_finish_split_0), v16)
							mBase = m.M
							v163 = m.ExcPending
							if v163 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F__bt_finish_split_1), int32(2282), int32(_a_F__bt_finish_split_2))
								mBase = m.M
								v168 = m.ExcPending
								if v168 != 0 {
									return
								} else {
									F__bt_insert_parent(m, l0, l1, l2, v40, l3, v108, v116)
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return
									} else {
										m.G0 = v16 + int32(16)
										return
									}
								}
							}
						} else {
							F__bt_insert_parent(m, l0, l1, l2, v40, l3, v108, v116)
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return
							} else {
								m.G0 = v16 + int32(16)
								return
							}
						}
					}
				}
			}
		} else {
			v108 = v5
			v109 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
			if v109 == int32(0) {
				v113 = *(*int32)(unsafe.Add(mBase, uint32(v59+v60)+4))
				v116 = base.B2i32(v113 == int32(0))
			} else {
				v116 = v5
			}
			v119 = F_errstart(m, int32(14), int32(0))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return
			} else {
				if v119 != 0 {
					if l2 < int32(0) {
						v124 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[2]))
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+(l2^int32(-1))<<(uint(int32(6))%32))+16))
						v139 = v130
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v132+l2<<(uint(int32(6))%32)+int32(-64))+16))
						v139 = v138
					}
					if v40 < int32(0) {
						v143 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[2]))
						v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+(v40^int32(-1))<<(uint(int32(6))%32))+16))
						v158 = v149
					} else {
						v151 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
						v157 = *(*int32)(unsafe.Add(mBase, uint32(v151+v40<<(uint(int32(6))%32)+int32(-64))+16))
						v158 = v157
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v158
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v139
					F_errmsg_internal(m, int32(_a_F__bt_finish_split_0), v16)
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F__bt_finish_split_1), int32(2282), int32(_a_F__bt_finish_split_2))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return
						} else {
							F__bt_insert_parent(m, l0, l1, l2, v40, l3, v108, v116)
							mBase = m.M
							v171 = m.ExcPending
							if v171 != 0 {
								return
							} else {
								m.G0 = v16 + int32(16)
								return
							}
						}
					}
				} else {
					F__bt_insert_parent(m, l0, l1, l2, v40, l3, v108, v116)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return
					} else {
						m.G0 = v16 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F__bt_fix_scankey_strategy(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v7 = int32(1)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v6<<(uint(v7)%32)-int32(2)))))
	v14 = v12 << (uint(int32(24)) % 32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15&v7 != 0 {
		v18 = v15 | v14
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18
		if v15&int32(64) != 0 {
			v93 = int32(3)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v93)
			return int32(1)
		} else {
			if v15&int32(128) != 0 {
				if v18&int32(33554432) != 0 {
					v92 = int32(5)
				} else {
					v92 = int32(1)
				}
				v93 = v92
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v93)
				return int32(1)
			} else {
				return int32(0)
			}
		}
	} else {
		v29 = int32(0)
		if base.B2i32(v12&int32(1) == v29)|v15&int32(16777216) == v29 {
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
			v38 = int32(6) - v37
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v38)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v15 | v14
		if v15&int32(4) == int32(0) {
			return int32(1)
		} else {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
			if v47&int32(1) != 0 {
				return int32(0)
			} else {
				v52 = v46
				for {
					v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+4)))
					v58 = int32(1)
					v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v57<<(uint(v58)%32)-int32(2)))))
					v68 = int32(0)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
					if base.B2i32(v63&v58 == v68)|v70&int32(16777216) == v68 {
						v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)))
						v78 = int32(6) - v77
						*(*uint16)(unsafe.Add(mBase, uint32(v52)+6)) = uint16(v78)
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v70 | v63<<(uint(int32(24))%32)
					if v70&int32(16) == int32(0) {
						v52 = v52 + int32(48)
						continue
					} else {
						break
					}
					break
				}
				return int32(1)
			}
		}
	}
}
func F__bt_getroot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v246 int32
	_ = v246
	var v249 int64
	_ = v249
	var v250 int32
	_ = v250
	var v251 int64
	_ = v251
	var v256 int64
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v286 int64
	_ = v286
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	goto L3
L1:
	;
	m.G0 = v12 - int32(-64)
	return v383
L2:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v87)+16))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v281 = F_MemoryContextAlloc(m, v279, int32(48))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L8
	} else {
		goto L78
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v25 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v107 = F__bt_allocbuf(m, l0, l1)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L8
	} else {
		goto L43
	}
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	v28 = F_ReadBuffer(m, l0, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v80 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L27
	}
L8:
	;
	return int32(0)
L9:
	;
	F_LockBuffer(m, v28, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	F__bt_checkpage(m, l0, v28)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	if v28 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_LockBuffer(m, v28, int32(0))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L21
	}
L13:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)))
	v56 = v55 + v54
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+12)))
	if v57&int32(20) != 0 {
		goto L12
	} else {
		goto L17
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[0]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(v28^int32(-1))<<(uint(int32(2))%32))))
	v54 = v46
	goto L13
L15:
	;
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[1]))
	v54 = v48 + v28<<(uint(int32(13))%32) + int32(-8192)
	goto L13
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	if v60 != v26 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	if v62 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v63 == int32(0) {
		v383 = v28
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L12
L21:
	;
	F_ReleaseBuffer(m, v28)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v71 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	F_pfree(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L8
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = int32(0)
	goto L7
L26:
	;
	goto L25
L27:
	;
	F_LockBuffer(m, v80, int32(1))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	F__bt_checkpage(m, l0, v80)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v87 = F__bt_getmeta(m, l0, v80)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	if v89 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	F_LockBuffer(m, v80, int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	if base.B2i32(l2 != int32(1)) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_ReleaseBuffer(m, v80)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_LockBuffer(m, v80, int32(2))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L8
	} else {
		goto L37
	}
L36:
	;
	v383 = int32(0)
	goto L1
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	if v101 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_LockBuffer(m, v80, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L8
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	goto L4
L41:
	;
	F_ReleaseBuffer(m, v80)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	goto L3
L43:
	;
	if v107 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v107 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[2]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112+(v107^int32(-1))<<(uint(int32(6))%32))+16))
	v127 = v118
	goto L44
L46:
	;
	goto L47
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[3]))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v120+v107<<(uint(int32(6))%32)+int32(-64))+16))
	v127 = v126
	goto L44
L48:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+16)))
	v147 = v146 + v145
	v148 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v148
	*(*int64)(unsafe.Add(mBase, uint32(v147)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = int32(3)
	if v80 < v148 {
		goto L53
	} else {
		goto L54
	}
L49:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[0]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v107^int32(-1))<<(uint(int32(2))%32))))
	v145 = v137
	goto L48
L50:
	;
	goto L51
L51:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[1]))
	v145 = v139 + v107<<(uint(int32(13))%32) + int32(-8192)
	goto L48
L52:
	;
	v172 = int32(_a_F__bt_getroot_0)
	v174 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[4])) = v174 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if base.Ui32(v178) <= base.Ui32(int32(2)) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[0]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157+(v80^int32(-1))<<(uint(int32(2))%32))))
	v171 = v163
	goto L52
L54:
	;
	goto L55
L55:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[1]))
	v171 = v165 + v80<<(uint(int32(13))%32) + int32(-8192)
	goto L52
L56:
	;
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171)+64)) = uint8(v181)
	*(*int64)(unsafe.Add(mBase, uint32(v171)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v171)+48)) = v181
	*(*int32)(unsafe.Add(mBase, uint32(v171)+28)) = int32(3)
	v189 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+12)) = uint16(v189)
	goto L58
L57:
	;
	goto L58
L58:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v87)+32)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+16)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v87)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = v127
	F_MarkBufferDirty(m, v107)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	F_MarkBufferDirty(m, v80)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+118)))
	if v204 != int32(112) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v260 = int32(_a_F__bt_getroot_0)
	v262 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[4])) = v262 - int32(1)
	F_LockBuffer(m, v107, int32(0))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L8
	} else {
		goto L74
	}
L62:
	;
	v208 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[5]))
	if v208 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v211 != 0 {
		goto L61
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L8
	} else {
		goto L68
	}
L66:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v212 != 0 {
		goto L61
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	F_XLogRegisterBuffer(m, int32(0), v107, int32(6))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	F_XLogRegisterBuffer(m, int32(2), v80, int32(14))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v223
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+52)) = uint8(v231)
	F_XLogRegisterBufData(m, int32(2), v10+int32(-36), int32(28))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v127
	F_XLogRegisterData(m, v10+int32(-8), int32(8))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v249 = F_XLogInsert(m, int32(11), int32(160))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	v251 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = base.I64_rotr(v249, v251)
	*(*uint32)(unsafe.Add(mBase, uint32(v171)+4)) = uint32(v249)
	v256 = int64(base.Ui64(v249) >> (uint(v251) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v171))) = uint32(v256)
	goto L61
L74:
	;
	F_LockBuffer(m, v107, int32(1))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	F_LockBuffer(m, v80, int32(0))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	F_ReleaseBuffer(m, v80)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v383 = v107
	goto L1
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v281
	v284 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+40)) = v284
	v286 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+32)) = v286
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+24)) = v288
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v87)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+16)) = v290
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v281)+8)) = v292
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	*(*int64)(unsafe.Add(mBase, uint32(v281))) = v294
	v298 = v80
	v302 = v278
	goto L80
L79:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	if v359 == v277 {
		v383 = v308
		goto L1
	} else {
		goto L98
	}
L80:
	;
	if v298 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L8
	} else {
		goto L95
	}
L82:
	;
	F_LockBuffer(m, v298, int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L8
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v308 = F_ReleaseAndReadBuffer(m, v298, l0, v302)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L8
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	F_LockBuffer(m, v308, int32(1))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	F__bt_checkpage(m, l0, v308)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	if v308 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v332)+16)))
	v334 = v333 + v332
	v335 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+12)))
	if v335&int32(20) == int32(0) {
		goto L79
	} else {
		goto L93
	}
L90:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[0]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318+(v308^int32(-1))<<(uint(int32(2))%32))))
	v332 = v324
	goto L89
L91:
	;
	goto L92
L92:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[1]))
	v332 = v326 + v308<<(uint(int32(13))%32) + int32(-8192)
	goto L89
L93:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v334)+4))
	if v340 != 0 {
		v298 = v308
		v302 = v340
		goto L80
	} else {
		goto L94
	}
L94:
	;
	goto L81
L95:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v345 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_getroot_1), v10+int32(-48))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F__bt_getroot_2), int32(548), int32(_a_F__bt_getroot_3))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L8
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v334)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v302
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v365 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_getroot_4), v12)
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F__bt_getroot_2), int32(555), int32(_a_F__bt_getroot_3))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_initmetapage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	v4 = l3
	v5 = int32(_a_F__bt_initmetapage_0)
	v7 = int32(0)
	if v7|(l0&int32(3)|int32(1)) == v7 {
		v23 = l0 + v5
		v25 = l0 + int32(4)
		if base.Ui32(v25) < base.Ui32(v23) {
			v27 = v23
		} else {
			v27 = v25
		}
		v32 = (l0^int32(-1)+v27)&int32(-4) + int32(4)
		if v32 == int32(0) {
		} else {
			base.MemoryFill(m, l0, int32(0), v32)
		}
	} else {
		base.MemoryFill(m, l0, int32(0), v5)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(_a_F__bt_initmetapage_1)
	v46 = int32(_a_F__bt_initmetapage_2)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v46)
	v52 = int32(_a_F__bt_initmetapage_3)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v52)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v52)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(17180209506)
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v68 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v66)+12)) = uint16(v68)
	v70 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v70)
	return
}
func F__bt_leftsib_splitflag(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v4 = int32(0)
	if l1 == v4 {
		return int32(0)
	} else {
		v9 = F_ReadBuffer(m, l0, l1)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_LockBuffer(m, v9, int32(1))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				F__bt_checkpage(m, l0, v9)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					if v9 < int32(0) {
						v21 = *(*int32)(unsafe.Add(mBase, _c_F__bt_leftsib_splitflag[0]))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v9^int32(-1))<<(uint(int32(2))%32))))
						v35 = v27
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, _c_F__bt_leftsib_splitflag[1]))
						v35 = v29 + v9<<(uint(int32(13))%32) + int32(-8192)
					}
					v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)))
					v37 = v36 + v35
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
					if l2 == v38 {
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+12)))
						v43 = int32(base.Ui32(v40) >> (uint(int32(7)) % 32))
					} else {
						v43 = v4
					}
					F_LockBuffer(m, v9, int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_ReleaseBuffer(m, v9)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							return v43
						}
					}
				}
			}
		}
	}
}
func F__bt_mark_scankey_required(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v15 = (v11 - int32(1)) & int32(_a_F__bt_mark_scankey_required_0)
	if base.Ui32(v15) < base.Ui32(int32(5)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(2))%32))+uint32(_c_F__bt_mark_scankey_required[0])))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18 | v21
	if v18&int32(4) == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	m.G0 = v9 + int32(16)
	return
L5:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
	if v28 != v30 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v33 = v29
	v34 = v28
	goto L7
L7:
	;
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+6)))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v38 != v39 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v33))) = v41 | v21
	if v41&int32(16) != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+52)))
	v50 = v34 + int32(1)
	if v46 == v50&int32(_a_F__bt_mark_scankey_required_0) {
		v33 = v33 + int32(48)
		v34 = v50
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	return
L13:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v67
	F_errmsg_internal(m, int32(_a_F__bt_mark_scankey_required_1), v9)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F__bt_mark_scankey_required_2), int32(797), int32(_a_F__bt_mark_scankey_required_3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_parallel_build_main(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int64
	_ = v160
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v175 int64
	_ = v175
	var v177 int64
	_ = v177
	var v181 int64
	_ = v181
	var v183 int64
	_ = v183
	var v187 int64
	_ = v187
	var v189 int64
	_ = v189
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	v3 = int32(0)
	v15 = F_shm_toc_lookup(m, l1, int64(-6917529027641081852), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[0])) = v15
		F_pgstat_report_activity(m, int32(3), v15)
		mBase = m.M
		v22 = F_shm_toc_lookup(m, l1, int64(-6917529027641081855), int32(0))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+10)))
			v25 = *(*int64)(unsafe.Add(mBase, uint32(v22)+16))
			v29 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[1]))
			if v29 == int32(0) {
			} else {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[2])))
				if v33&int32(1) == int32(0) {
				} else {
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v29)+392))
					if int32(1)&base.B2i32(v40 != int64(0)) != 0 {
					} else {
						v44 = int32(_a_F__bt_parallel_build_main_0)
						v46 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[3]))
						v47 = int32(1)
						*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[3])) = v46 + v47
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = v50 + v47
						*(*int64)(unsafe.Add(mBase, uint32(v29)+392)) = v25
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = v50 + int32(2)
						v61 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[3]))
						*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[3])) = v61 - v47
					}
				}
			}
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v24 != 0 {
				v68 = int32(4)
			} else {
				v68 = int32(5)
			}
			v69 = F_table_open(m, v65, v68)
			mBase = m.M
			v70 = m.ExcPending
			if v70 != 0 {
				return
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				if v24 != 0 {
					v74 = int32(3)
				} else {
					v74 = int32(8)
				}
				v75 = F_index_open(m, v71, v74)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return
				} else {
					v78 = F_palloc0(m, int32(16))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v75
						*(*int32)(unsafe.Add(mBase, uint32(v78)+4)) = v69
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
						*(*uint8)(unsafe.Add(mBase, uint32(v78)+12)) = uint8(v82)
						v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
						*(*uint8)(unsafe.Add(mBase, uint32(v78)+13)) = uint8(v84)
						v88 = F_shm_toc_lookup(m, l1, int64(-6917529027641081854), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							F_tuplesort_attach_shared(m, v88, l0)
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return
							} else {
								v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
								if v92 == int32(1) {
									v96 = F_palloc0(m, int32(16))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v98
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
										v101 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v96)+12)) = uint8(v101)
										*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
										v106 = F_shm_toc_lookup(m, l1, int64(-6917529027641081853), v101)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return
										} else {
											F_tuplesort_attach_shared(m, v106, l0)
											mBase = m.M
											v109 = m.ExcPending
											if v109 != 0 {
												return
											} else {
												v110 = v96
												v111 = v106
												base.MemoryCopy(m, int32(_a_F__bt_parallel_build_main_1), int32(_a_F__bt_parallel_build_main_2), int32(128))
												v118 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[4]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[5])) = v118
												v122 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[6]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[7])) = v122
												v126 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[8]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[9])) = v126
												v130 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[10]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[11])) = v130
												v133 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[12]))
												v134 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
												v135 = base.I32_div_s(v133, v134)
												F__bt_parallel_scan_and_sort(m, v78, v110, v22, v88, v111, v135, int32(0))
												mBase = m.M
												v138 = m.ExcPending
												if v138 != 0 {
													return
												} else {
													v141 = F_shm_toc_lookup(m, l1, int64(-6917529027641081850), int32(0))
													mBase = m.M
													v142 = m.ExcPending
													if v142 != 0 {
														return
													} else {
														v145 = F_shm_toc_lookup(m, l1, int64(-6917529027641081851), int32(0))
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return
														} else {
															v148 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[13]))
															v151 = v141 + v148<<(uint(int32(7))%32)
															v154 = v145 + v148<<(uint(int32(5))%32)
															base.MemoryFill(m, v151, int32(0), int32(128))
															F_BufferUsageAccumDiff(m, v151, int32(_a_F__bt_parallel_build_main_1))
															mBase = m.M
															v160 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v154)+24)) = v160
															*(*int64)(unsafe.Add(mBase, uint32(v154)+16)) = v160
															*(*int64)(unsafe.Add(mBase, uint32(v154)+8)) = v160
															*(*int64)(unsafe.Add(mBase, uint32(v154))) = v160
															v169 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[6]))
															v171 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[7]))
															*(*int64)(unsafe.Add(mBase, uint32(v154)+16)) = v169 - v171
															v175 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[10]))
															v177 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[11]))
															*(*int64)(unsafe.Add(mBase, uint32(v154))) = v175 - v177
															v181 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[8]))
															v183 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[9]))
															*(*int64)(unsafe.Add(mBase, uint32(v154)+8)) = v181 - v183
															v187 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[4]))
															v189 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[5]))
															*(*int64)(unsafe.Add(mBase, uint32(v154)+24)) = v187 - v189
															F_relation_close(m, v75, v74)
															mBase = m.M
															v193 = m.ExcPending
															if v193 != 0 {
																return
															} else {
																F_relation_close(m, v69, v68)
																mBase = m.M
																v195 = m.ExcPending
																if v195 != 0 {
																	return
																} else {
																	return
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v110 = v3
									v111 = v3
									base.MemoryCopy(m, int32(_a_F__bt_parallel_build_main_1), int32(_a_F__bt_parallel_build_main_2), int32(128))
									v118 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[4]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[5])) = v118
									v122 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[6]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[7])) = v122
									v126 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[8]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[9])) = v126
									v130 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[10]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[11])) = v130
									v133 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[12]))
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
									v135 = base.I32_div_s(v133, v134)
									F__bt_parallel_scan_and_sort(m, v78, v110, v22, v88, v111, v135, int32(0))
									mBase = m.M
									v138 = m.ExcPending
									if v138 != 0 {
										return
									} else {
										v141 = F_shm_toc_lookup(m, l1, int64(-6917529027641081850), int32(0))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return
										} else {
											v145 = F_shm_toc_lookup(m, l1, int64(-6917529027641081851), int32(0))
											mBase = m.M
											v146 = m.ExcPending
											if v146 != 0 {
												return
											} else {
												v148 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[13]))
												v151 = v141 + v148<<(uint(int32(7))%32)
												v154 = v145 + v148<<(uint(int32(5))%32)
												base.MemoryFill(m, v151, int32(0), int32(128))
												F_BufferUsageAccumDiff(m, v151, int32(_a_F__bt_parallel_build_main_1))
												mBase = m.M
												v160 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v154)+24)) = v160
												*(*int64)(unsafe.Add(mBase, uint32(v154)+16)) = v160
												*(*int64)(unsafe.Add(mBase, uint32(v154)+8)) = v160
												*(*int64)(unsafe.Add(mBase, uint32(v154))) = v160
												v169 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[6]))
												v171 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[7]))
												*(*int64)(unsafe.Add(mBase, uint32(v154)+16)) = v169 - v171
												v175 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[10]))
												v177 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[11]))
												*(*int64)(unsafe.Add(mBase, uint32(v154))) = v175 - v177
												v181 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[8]))
												v183 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[9]))
												*(*int64)(unsafe.Add(mBase, uint32(v154)+8)) = v181 - v183
												v187 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[4]))
												v189 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[5]))
												*(*int64)(unsafe.Add(mBase, uint32(v154)+24)) = v187 - v189
												F_relation_close(m, v75, v74)
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return
												} else {
													F_relation_close(m, v69, v68)
													mBase = m.M
													v195 = m.ExcPending
													if v195 != 0 {
														return
													} else {
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
			}
		}
	}
}
func F__bt_skiparray_shrink(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v5 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+19)) = uint8(v5)
	v13 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v13)
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	switch v15 - v13 {
	case 0, 1:
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
		if v18 != 0 {
			v19 = int32(0)
			v24 = F__bt_compare_scankey_args(m, l0, v18, l1, v18, v19, v19, v9+int32(15))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				if v24 == int32(0) {
					v55 = v19
				} else {
					v30 = int32(1)
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v31 != v30 {
						v55 = v30
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l1
						v55 = int32(1)
					}
				}
				m.G0 = v9 + int32(16)
				return v55
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = l1
			v55 = int32(1)
			m.G0 = v9 + int32(16)
			return v55
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v64 = m.ExcPending
		if v64 != 0 {
			return int32(0)
		} else {
			v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v65
			F_errmsg_internal(m, int32(_a_F__bt_skiparray_shrink_0), v9)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F__bt_skiparray_shrink_1), int32(1338), int32(_a_F__bt_skiparray_shrink_2))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3, 4:
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		if v36 != 0 {
			v37 = int32(0)
			v42 = F__bt_compare_scankey_args(m, l0, v36, l1, v36, v37, v37, v9+int32(15))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				if v42 == int32(0) {
					v55 = v37
				} else {
					v46 = int32(1)
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v47 != v46 {
						v55 = v46
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l1
						v55 = int32(1)
					}
				}
				m.G0 = v9 + int32(16)
				return v55
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = l1
			v55 = int32(1)
			m.G0 = v9 + int32(16)
			return v55
		}
	}
}
func F__bt_splitcmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	return v3 - v4
}
func F__bt_stepright(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v14 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+16)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v33+v32)+4))
	v41 = int32(0)
	v44 = v35
	goto L6
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F__bt_stepright[0]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(v14^int32(-1))<<(uint(int32(2))%32))))
	v32 = v24
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F__bt_stepright[1]))
	v32 = v26 + v14<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F__bt_relbuf(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L31
	}
L6:
	;
	v47 = F__bt_relandgetbuf(m, l0, v41, v44, int32(2))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L28
	}
L8:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+16)))
	v68 = v67 + v66
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+12)))
	if v69&int32(128) != 0 {
		goto L14
	} else {
		goto L15
	}
L9:
	;
	return
L10:
	;
	if int32(0) <= v47 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F__bt_stepright[1]))
	v66 = v52 + v47<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L12:
	;
	goto L13
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _c_F__bt_stepright[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v59+(v47^int32(-1))<<(uint(int32(2))%32))))
	v66 = v65
	goto L8
L14:
	;
	v76 = v47
	goto L17
L15:
	;
	v114 = v47
	v115 = v68
	v118 = v69
	goto L16
L16:
	;
	if v118&int32(20) == int32(0) {
		goto L5
	} else {
		goto L26
	}
L17:
	;
	F__bt_finish_split(m, l0, l1, v76, l3)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v114 = v85
	v115 = v106
	v118 = v107
	goto L16
L19:
	;
	v85 = F__bt_relandgetbuf(m, l0, int32(0), v44, int32(2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L21
	}
L20:
	;
	v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+16)))
	v106 = v105 + v104
	v107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v106)+12)))
	if v107&int32(128) != 0 {
		v76 = v85
		goto L17
	} else {
		goto L25
	}
L21:
	;
	if v85 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F__bt_stepright[0]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90+(v85^int32(-1))<<(uint(int32(2))%32))))
	v104 = v96
	goto L20
L23:
	;
	goto L24
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F__bt_stepright[1]))
	v104 = v98 + v85<<(uint(int32(13))%32) + int32(-8192)
	goto L20
L25:
	;
	goto L18
L26:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v123 != 0 {
		v41 = v114
		v44 = v123
		goto L6
	} else {
		goto L27
	}
L27:
	;
	goto L7
L28:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v128 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_stepright_0), v12)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F__bt_stepright_1), int32(1064), int32(_a_F__bt_stepright_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v143 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v143)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v114
	m.G0 = v12 + int32(16)
	return
}
func F__bt_swap_posting(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v16 = v14 & int32(4095)
	if base.B2i32(l2 <= v4)|base.B2i32(base.Ui32(v16) <= base.Ui32(l2)) == v4 {
		v21 = F_CopyIndexTuple(m, l1)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+2)))
			v31 = int32(6)
			v33 = v21 + v25<<(uint(int32(16))%32) + v29 + l2*v31
			v38 = (v16 + (l2 ^ int32(-1))) * v31
			if v38 != 0 {
				base.MemoryCopy(m, v33+int32(6), v33, v38)
			} else {
			}
			v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)) = uint16(v42)
			v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			*(*int32)(unsafe.Add(mBase, uint32(v33))) = v44
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
			if v46&int32(32) == int32(0) {
				v69 = l1
			} else {
				v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				if v51&int32(_a_F__bt_swap_posting_0) == int32(0) {
					v69 = l1
				} else {
					v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
					v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
					v64 = int32(6)
					v69 = v56 + (l1 + v57<<(uint(int32(16))%32)) + v51&int32(4095)*v64 - v64
				}
			}
			v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v71)
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v73
			m.G0 = v10 + int32(16)
			return v21
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v10))) = v16
			F_errmsg_internal(m, int32(_a_F__bt_swap_posting_1), v10)
			mBase = m.M
			v87 = m.ExcPending
			if v87 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F__bt_swap_posting_2), int32(1044), int32(_a_F__bt_swap_posting_3))
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
	}
}
func F__bt_truncate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+10)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v19 != int32(1) {
		v65 = v18
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v76 = m.G0
	v78 = v76 - int32(160)
	m.G0 = v78
	if v65 < v18 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v22 = int32(1)
	if v18 <= int32(0) {
		v65 = v22
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = v22
	v32 = l3 + int32(16)
	goto L4
L4:
	;
	v42 = F_index_getattr_2(m, l1, v29, v16, v14+int32(15))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v65 = v18 + int32(1)
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	v48 = F_index_getattr_2(m, l2, v29, v16, v14+int32(14))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)))
	if v50 != v51 {
		v65 = v29
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v50 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	v58 = F_FunctionCall2Coll(m, v32+int32(16), v57, v42, v48)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v29 != v18 {
		v29 = v29 + int32(1)
		v32 = v32 + int32(48)
		goto L4
	} else {
		goto L15
	}
L13:
	;
	if v58 != 0 {
		v65 = v29
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	goto L5
L16:
	;
	m.G0 = v78 + int32(160)
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)))
	if v126&int32(_a_F__bt_truncate_0) == int32(0) {
		v146 = v126
		goto L32
	} else {
		goto L33
	}
L17:
	;
	v81 = v65
	goto L19
L18:
	;
	v81 = v18
	goto L19
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v81 == v82 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v86 = v84 & int32(_a_F__bt_truncate_1)
	v87 = F_palloc(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v92 = F_CreateTupleDescTruncatedCopy(m, v16, v81)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	if v86 == int32(0) {
		v120 = v87
		goto L16
	} else {
		goto L24
	}
L24:
	;
	base.MemoryCopy(m, v87, l2, v86)
	v120 = v87
	goto L16
L25:
	;
	v95 = v78 + int32(32)
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	if int32(0) <= base.I32_extend16_s(v98) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v102 = int32(8)
	goto L28
L27:
	;
	v102 = int32(16)
	goto L28
L28:
	;
	F_index_deform_tuple_internal(m, v92, v95, v78, l2+v102, l2+int32(8), int32(base.Ui32(v98)>>(uint(int32(15))%32)))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F__bt_truncate[0]))
	v112 = F_index_form_tuple_context(m, v92, v95, v78, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)) = uint16(v114)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v116
	F_pfree(m, v92)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	v120 = v112
	goto L16
L32:
	;
	if v65 <= v18 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
	if v131&int32(32) == int32(0) {
		v146 = v126
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v137 = v126 & int32(_a_F__bt_truncate_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)) = uint16(v137)
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v144 = (v139+int32(7))&int32(-8200) | v137
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)) = uint16(v144)
	v146 = v144
	goto L32
L35:
	;
	m.G0 = v14 + int32(16)
	return v226
L36:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)) = uint16(v65)
	v150 = v146 | int32(_a_F__bt_truncate_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)) = uint16(v150)
	v226 = v120
	goto L35
L37:
	;
	goto L38
L38:
	;
	v159 = (v146&int32(_a_F__bt_truncate_1)+int32(7))&int32(_a_F__bt_truncate_3) + int32(8)
	v160 = F_palloc0(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)))
	v168 = (v162&int32(_a_F__bt_truncate_1) + int32(7)) & int32(_a_F__bt_truncate_3)
	if v168 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	base.MemoryCopy(m, v160, v120, v168)
	goto L42
L41:
	;
	goto L42
L42:
	;
	F_pfree(m, v120)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v173 = v18 | int32(_a_F__bt_truncate_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v160)+4)) = uint16(v173)
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+6)))
	v179 = int32(_a_F__bt_truncate_0)
	v180 = v159 | v175&int32(_a_F__bt_truncate_5) | v179
	*(*uint16)(unsafe.Add(mBase, uint32(v160)+6)) = uint16(v180)
	if v18&v179 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v196 = v194 + (v160 + v193)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v197&int32(32) == int32(0) {
		v220 = l1
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v193 = v159 & int32(_a_F__bt_truncate_6)
	v194 = int32(-6)
	goto L44
L46:
	;
	goto L47
L47:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160))))
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+2)))
	v193 = v189 << (uint(int32(16)) % 32)
	v194 = v192
	goto L44
L48:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+4)) = uint16(v222)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v224
	v226 = v160
	goto L35
L49:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v202&int32(_a_F__bt_truncate_0) == int32(0) {
		v220 = l1
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v215 = int32(6)
	v220 = v207 + (l1 + v208<<(uint(int32(16))%32)) + v202&int32(4095)*v215 - v215
	goto L48
}
func F__bt_tuple_before_array_skeys(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v215 int32
	_ = v215
	var v242 int32
	_ = v242
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if l7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v22)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v24 <= l6 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v19 + int32(16)
	return v242
L5:
	;
	v242 = int32(0)
	goto L4
L6:
	;
	v26 = v24
	v32 = l6
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v45 = v42 + v32*int32(48)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
	if v46&int32(3) == int32(0) {
		goto L5
	} else {
		goto L9
	}
L8:
	;
	goto L5
L9:
	;
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+4)))
	if l4 < v51 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if l7 == int32(0) {
		goto L5
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)))
	if v57 != int32(3) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v55)
	goto L5
L14:
	;
	v215 = v32 + int32(1)
	if v215 < v198 {
		v26 = v198
		v32 = v215
		goto L7
	} else {
		goto L61
	}
L15:
	;
	if l5 == int32(0) {
		v198 = v26
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v64 = F_index_getattr_2(m, l2, v51, l3, v19+int32(15))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L5
L19:
	;
	return int32(0)
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v68&int32(_a_F__bt_tuple_before_array_skeys_0) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v184 = int32(1)
	v187 = int32(0)
	if base.B2i32(l1 == v184)&base.B2i32(v168 < v187)|base.B2i32(l1 == int32(-1))&base.B2i32(v187 < v168) != 0 {
		v242 = v184
		goto L4
	} else {
		goto L59
	}
L22:
	;
	v71 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v73 <= v71 {
		v115 = v71
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v125 = int32(1)
	v126 = v68 & v125
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	if v127 == v125 {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	F__bt_binsrch_skiparray_skey(m, int32(0), l1, v64, v119, v115, v45, v19+int32(8))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L19
	} else {
		goto L31
	}
L26:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v78 = int32(0)
	goto L27
L27:
	;
	v96 = v76 + v78<<(uint(int32(5))%32)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	if v97 == v32 {
		v115 = v96
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v115 = v96
	goto L25
L29:
	;
	v100 = v78 + int32(1)
	if v100 < v73 {
		v78 = v100
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v124 != 0 {
		v168 = v124
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v242 = v71
	goto L4
L33:
	;
	v168 = int32(1)
	goto L21
L34:
	;
	if v159&int32(_a_F__bt_tuple_before_array_skeys_1) != 0 {
		goto L56
	} else {
		goto L57
	}
L35:
	;
	if v126 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	goto L37
L37:
	;
	if v126 != 0 {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
	v159 = v68
	goto L34
L39:
	;
	goto L40
L40:
	;
	if v68&int32(33554432) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v136 = int32(-1)
	goto L43
L42:
	;
	v136 = int32(1)
	goto L43
L43:
	;
	v168 = v136
	goto L21
L44:
	;
	if v68&int32(33554432) != 0 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v45)+44))
	v148 = F_FunctionCall2Coll(m, v142+v32*int32(28), v146, v64, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L19
	} else {
		goto L50
	}
L47:
	;
	v141 = int32(1)
	goto L49
L48:
	;
	v141 = int32(-1)
	goto L49
L49:
	;
	v168 = v141
	goto L21
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v150&int32(16777216) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v148 < int32(0) {
		goto L33
	} else {
		goto L54
	}
L52:
	;
	v157 = v148
	goto L53
L53:
	;
	if v157 != 0 {
		v168 = v157
		goto L21
	} else {
		goto L55
	}
L54:
	;
	v157 = int32(0) - v148
	goto L53
L55:
	;
	v159 = v150
	goto L34
L56:
	;
	v168 = int32(-1)
	goto L21
L57:
	;
	goto L58
L58:
	;
	v168 = int32(base.Ui32(v159)>>(uint(int32(22))%32)) & int32(1)
	goto L21
L59:
	;
	if v168|l5 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v198 = v197
	goto L14
L61:
	;
	goto L8
}
func F__bt_update_posting(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
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
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v16 = v12&int32(4095) - v15
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	v21 = v19 << (uint(int32(16)) % 32)
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)))
	v23 = v21 | v22
	if int32(1) < v16 {
		v31 = (v16*int32(6) + v23 + int32(7)) & int32(-8)
	} else {
		v31 = v23
	}
	v32 = F_palloc0(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		return
	} else {
		if v23 != 0 {
			base.MemoryCopy(m, v32, v11, v23)
		} else {
		}
		v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+6)))
		v38 = v35&int32(-8192) | v31
		if int32(2) <= v16 {
			*(*uint16)(unsafe.Add(mBase, uint32(v32)+2)) = uint16(v22)
			*(*uint16)(unsafe.Add(mBase, uint32(v32))) = uint16(v19)
			v43 = int32(_a_F__bt_update_posting_0)
			v44 = v16 | v43
			*(*uint16)(unsafe.Add(mBase, uint32(v32)+4)) = uint16(v44)
			v52 = v38 | v43
			v53 = v32 + v21 + v22
		} else {
			v52 = v38 & int32(_a_F__bt_update_posting_1)
			v53 = v32
		}
		*(*uint16)(unsafe.Add(mBase, uint32(v32)+6)) = uint16(v52)
		v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		if v55&int32(4095) != 0 {
			v60 = int32(0)
			v64 = v60
			v65 = v60
			v68 = v55
			v69 = v60
			for {
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
				if v73 <= v65 {
					v82 = int32(6)
					v84 = v53 + v69*v82
					v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)))
					v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
					v93 = v85 + (v11 + v86<<(uint(int32(16))%32)) + v64*v82
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = v94
					v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)) = uint16(v96)
					v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
					v101 = v65
					v102 = v100
					v103 = v69 + int32(1)
				} else {
					v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(8)+v65<<(uint(int32(1))%32)))))
					if v64 != v78 {
						v82 = int32(6)
						v84 = v53 + v69*v82
						v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)))
						v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
						v93 = v85 + (v11 + v86<<(uint(int32(16))%32)) + v64*v82
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
						*(*int32)(unsafe.Add(mBase, uint32(v84))) = v94
						v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)) = uint16(v96)
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
						v101 = v65
						v102 = v100
						v103 = v69 + int32(1)
					} else {
						v101 = v65 + int32(1)
						v102 = v68
						v103 = v69
					}
				}
				v106 = v64 + int32(1)
				if base.Ui32(v106) < base.Ui32(v102&int32(4095)) {
					v64 = v106
					v65 = v101
					v68 = v102
					v69 = v103
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32
		return
	}
}
func F_bt_child_highkey_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int64
	_ = v590
	var v595 int64
	_ = v595
	var v601 int32
	_ = v601
	var v606 int32
	_ = v606
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v655 int32
	_ = v655
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v677 int32
	_ = v677
	var v685 int32
	_ = v685
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v717 int32
	_ = v717
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v739 int32
	_ = v739
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v781 int32
	_ = v781
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int64
	_ = v805
	var v810 int64
	_ = v810
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int64
	_ = v840
	var v845 int64
	_ = v845
	var v851 int32
	_ = v851
	var v856 int32
	_ = v856
	var v860 int32
	_ = v860
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int64
	_ = v875
	var v880 int64
	_ = v880
	var v886 int32
	_ = v886
	var v891 int32
	_ = v891
	v5 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(384)
	m.G0 = v22
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.Ui32((l1-int32(1))&int32(_a_F_bt_child_highkey_check_0)) <= base.Ui32(int32(2047)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v35 = F_PageGetItemIdCareful_2(m, l0, v33, v34, l1)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v48 = v5
	goto L3
L3:
	;
	v50 = base.B2i32(v26 != int32(-1))
	if v26 != int32(-1) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	return
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	v41 = v37 + v38&int32(_a_F_bt_child_highkey_check_1)
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41))))
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+2)))
	v48 = v42<<(uint(int32(16))%32) | v45
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L4
	} else {
		goto L209
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L4
	} else {
		goto L204
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L4
	} else {
		goto L199
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L4
	} else {
		goto L194
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L4
	} else {
		goto L189
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L4
	} else {
		goto L184
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L4
	} else {
		goto L179
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L4
	} else {
		goto L175
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L4
	} else {
		goto L170
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L4
	} else {
		goto L165
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L161
	}
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)) = uint8(v532)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v530
	m.G0 = v22 + int32(384)
	return
L18:
	;
	v51 = v26
	goto L20
L19:
	;
	v51 = v48
	goto L20
L20:
	;
	if v51|v48 == int32(0) {
		v530 = int32(-1)
		v532 = v5
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v57 = int32(1)
	v58 = l3 - v57
	v66 = base.B2i32(v51 == int32(0))
	v67 = v51
	v68 = v57
	v71 = v50 & v25
	goto L22
L22:
	;
	if v66&int32(1) != 0 {
		goto L16
	} else {
		goto L24
	}
L23:
	;
	v530 = int32(-1)
	v532 = v520
	goto L17
L24:
	;
	if l2 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+16)))
	v87 = v86 + v85
	v89 = v68 & int32(1)
	if v89 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	if v67 == v48 {
		v85 = l2
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v83 = F_palloc_btree_page(m, l0, v67)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v85 = v83
	goto L25
L31:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+12)))
	if v99&int32(260) != int32(4) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v92 != int32(-1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v95 = F_bt_leftmost_ignoring_half_dead(m, l0, v67, v87)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v95 == int32(0) {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	if v104 != v58 {
		goto L14
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v89 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v67 == v108 {
		goto L13
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v67 == v110 {
		goto L13
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v112 = base.B2i32(v67 == v48)
	if v112|v99&int32(20) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+12)))
	v292 = v290 & int32(128)
	v294 = int32(base.Ui32(v292) >> (uint(int32(7)) % 32))
	if v294|v290&int32(16) != 0 {
		goto L89
	} else {
		goto L90
	}
L46:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+16)))
	v117 = v85 + v116
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+12)))
	if v118&int32(2) != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v71 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v125 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v118&int32(1) != 0 {
		goto L12
	} else {
		goto L57
	}
L51:
	;
	if v125 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+160)) = v133 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_child_highkey_check_2), v22+int32(160))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+144)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v22)+140)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v22)+136)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v22)+132)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v22)+128)) = v67
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_3), v22+int32(128))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2610), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	goto L45
L57:
	;
	v163 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	if v163 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+272)) = v166 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_child_highkey_check_6), v22+int32(272))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v117)+4))
	if v183 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2635), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v184 = int32(2)
	goto L66
L65:
	;
	v184 = int32(1)
	goto L66
L66:
	;
	v185 = F_PageGetItemIdCareful_2(m, l0, v67, v85, v184)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	v190 = v85 + v187&int32(_a_F_bt_child_highkey_check_1)
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190))))
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+2)))
	v199 = v180
	v201 = v191<<(uint(int32(16))%32) | v194
	goto L68
L68:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_bt_child_highkey_check[0]))
	if v216 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v223&int32(4) != 0 {
		goto L10
	} else {
		goto L84
	}
L70:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v219 = F_palloc_btree_page(m, l0, v201)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L4
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+16)))
	v222 = v219 + v221
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v222)+12)))
	if v223&int32(1) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v229 = v199 - int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	if v229 != v230 {
		goto L11
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L69
L78:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v234 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v235 = int32(2)
	goto L81
L80:
	;
	v235 = int32(1)
	goto L81
L81:
	;
	v236 = F_PageGetItemIdCareful_2(m, l0, v201, v219, v235)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	v241 = v219 + v238&int32(_a_F_bt_child_highkey_check_1)
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241)+2)))
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241))))
	F_pfree(m, v219)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v199 = v230
	v201 = v242 | v243<<(uint(int32(16))%32)
	goto L68
L84:
	;
	if v223&int32(16) == int32(0) {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	if v255 == int32(0) {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	v259 = F_PageGetItemIdCareful_2(m, l0, v201, v219, int32(1))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v264 = v219 + v261&int32(_a_F_bt_child_highkey_check_1)
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264))))
	v268 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v264)+2)))
	if v265<<(uint(int32(16))%32)|v268 != v67 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	goto L45
L89:
	;
	if v67 == v48 {
		goto L153
	} else {
		goto L154
	}
L90:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v298 == int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v302 = F_PageGetItemIdCareful_2(m, l0, v67, v85, int32(1))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v305 = l1 + v112
	v306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306)+16)))
	v308 = v306 + v307
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308)+12)))
	if v309&int32(1) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v358 = v85 + v304&int32(_a_F_bt_child_highkey_check_1)
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v358)+6)))
	v360 = int32(_a_F_bt_child_highkey_check_7)
	v361 = v359 & v360
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v354)+6)))
	if v361 != v362&v360 {
		goto L6
	} else {
		goto L111
	}
L94:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v350 == int32(0) {
		goto L7
	} else {
		goto L110
	}
L95:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	if v318 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v306)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v323) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v319 = int32(2)
	goto L100
L99:
	;
	v319 = int32(1)
	goto L100
L100:
	;
	if v305&int32(_a_F_bt_child_highkey_check_0) == v319 {
		goto L94
	} else {
		goto L101
	}
L101:
	;
	goto L97
L102:
	;
	v331 = int32(base.Ui32(v323+int32(_a_F_bt_child_highkey_check_8)) >> (uint(int32(2)) % 32))
	goto L104
L103:
	;
	v331 = int32(0)
	goto L104
L104:
	;
	if base.Ui32(v331&int32(_a_F_bt_child_highkey_check_0)) < base.Ui32(v305&int32(_a_F_bt_child_highkey_check_0)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v308)+4))
	if v336 == int32(0) {
		goto L8
	} else {
		goto L108
	}
L106:
	;
	v339 = v305
	goto L107
L107:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v343 = F_PageGetItemIdCareful_2(m, l0, v340, v306, v339&int32(_a_F_bt_child_highkey_check_0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L109
	}
L108:
	;
	v339 = int32(1)
	goto L107
L109:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v343)))
	v354 = v345 + v346&int32(_a_F_bt_child_highkey_check_1)
	goto L93
L110:
	;
	v354 = v350
	goto L93
L111:
	;
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v366 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v369 = int32(4)
	v370 = v358 + v369
	v372 = v354 + v369
	v374 = v361 - v369
	if base.Ui32(v369) <= base.Ui32(v374) {
		goto L118
	} else {
		goto L119
	}
L113:
	;
	goto L114
L114:
	;
	v439 = int32(6)
	v440 = v358 + v439
	v442 = v354 + v439
	v444 = v361 - v439
	if base.Ui32(int32(4)) <= base.Ui32(v444) {
		goto L137
	} else {
		goto L138
	}
L115:
	;
	if v436 == int32(0) {
		goto L89
	} else {
		goto L133
	}
L116:
	;
	v436 = int32(0)
	goto L115
L117:
	;
	v410 = v405
	v411 = v406
	v412 = v407
	goto L127
L118:
	;
	if (v370|v372)&int32(3) != 0 {
		v405 = v370
		v406 = v372
		v407 = v374
		goto L117
	} else {
		goto L121
	}
L119:
	;
	v398 = v370
	v399 = v372
	v400 = v374
	goto L120
L120:
	;
	if v400 == int32(0) {
		goto L116
	} else {
		goto L126
	}
L121:
	;
	v382 = v370
	v383 = v372
	v384 = v374
	goto L122
L122:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v382)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	if v387 != v388 {
		v405 = v382
		v406 = v383
		v407 = v384
		goto L117
	} else {
		goto L124
	}
L123:
	;
	v398 = v393
	v399 = v391
	v400 = v395
	goto L120
L124:
	;
	v390 = int32(4)
	v391 = v383 + v390
	v393 = v382 + v390
	v395 = v384 - v390
	if base.Ui32(int32(3)) < base.Ui32(v395) {
		v382 = v393
		v383 = v391
		v384 = v395
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v405 = v398
	v406 = v399
	v407 = v400
	goto L117
L127:
	;
	v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410))))
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v411))))
	if v415 == v416 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v436 = v415 - v416
	goto L115
L129:
	;
	v418 = int32(1)
	v423 = v412 - v418
	if v423 != 0 {
		v410 = v410 + v418
		v411 = v411 + v418
		v412 = v423
		goto L127
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	goto L128
L132:
	;
	goto L116
L133:
	;
	goto L6
L134:
	;
	if v506 != 0 {
		goto L6
	} else {
		goto L152
	}
L135:
	;
	v506 = int32(0)
	goto L134
L136:
	;
	v480 = v475
	v481 = v476
	v482 = v477
	goto L146
L137:
	;
	if (v440|v442)&int32(3) != 0 {
		v475 = v440
		v476 = v442
		v477 = v444
		goto L136
	} else {
		goto L140
	}
L138:
	;
	v468 = v440
	v469 = v442
	v470 = v444
	goto L139
L139:
	;
	if v470 == int32(0) {
		goto L135
	} else {
		goto L145
	}
L140:
	;
	v452 = v440
	v453 = v442
	v454 = v444
	goto L141
L141:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v452)))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v453)))
	if v457 != v458 {
		v475 = v452
		v476 = v453
		v477 = v454
		goto L136
	} else {
		goto L143
	}
L142:
	;
	v468 = v463
	v469 = v461
	v470 = v465
	goto L139
L143:
	;
	v460 = int32(4)
	v461 = v453 + v460
	v463 = v452 + v460
	v465 = v454 - v460
	if base.Ui32(int32(3)) < base.Ui32(v465) {
		v452 = v463
		v453 = v461
		v454 = v465
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v475 = v468
	v476 = v469
	v477 = v470
	goto L136
L146:
	;
	v485 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v480))))
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v481))))
	if v485 == v486 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v506 = v485 - v486
	goto L134
L148:
	;
	v488 = int32(1)
	v493 = v482 - v488
	if v493 != 0 {
		v480 = v480 + v488
		v481 = v481 + v488
		v482 = v493
		goto L146
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	goto L147
L151:
	;
	goto L135
L152:
	;
	goto L89
L153:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v530 = v515
	v532 = base.B2i32(v292 != int32(0))
	goto L17
L154:
	;
	goto L155
L155:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if l2 != v85 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	F_pfree(m, v85)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L4
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v520 = int32(0)
	if v516|v48 != 0 {
		v66 = base.B2i32(v516 == v520)
		v67 = v516
		v68 = v520
		v71 = v294
		goto L22
	} else {
		goto L160
	}
L159:
	;
	goto L158
L160:
	;
	goto L23
L161:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v556)+48))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v558
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v557 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_9), v22)
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2211), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L4
	} else {
		goto L164
	}
L164:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L165:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v579)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+368)) = v580 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_11), v22+int32(368))
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v590 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+364)) = uint32(v590)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+356)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v22)+352)) = v589
	v595 = int64(base.Ui64(v590) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+360)) = uint32(v595)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_12), v22+int32(352))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2230), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
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
	F_errcode(m, int32(33557032))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v614)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+336)) = v615 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_13), v22+int32(336))
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v87)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+328)) = v624
	*(*int32)(unsafe.Add(mBase, uint32(v22)+324)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v22)+320)) = v67
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_14), v22+int32(320))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2240), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L4
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v646 = *(*int32)(unsafe.Add(mBase, uint32(v645)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v646 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_15), v22+int32(16))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2247), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L4
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v668 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v668)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+304)) = v669 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_16), v22+int32(304))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+296)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v22)+292)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v22)+288)) = v67
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_17), v22+int32(288))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2631), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L4
	} else {
		goto L183
	}
L183:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L184:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v698)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+192)) = v699 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_child_highkey_check_18), v22+int32(192))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v222)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+188)) = v708
	*(*int32)(unsafe.Add(mBase, uint32(v22)+184)) = v229
	*(*int32)(unsafe.Add(mBase, uint32(v22)+180)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v22)+176)) = v67
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_19), v22+int32(176))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2659), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L4
	} else {
		goto L188
	}
L188:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L189:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v731 = *(*int32)(unsafe.Add(mBase, uint32(v730)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+256)) = v731 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_child_highkey_check_20), v22+int32(256))
	mBase = m.M
	v739 = m.ExcPending
	if v739 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+252)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v22)+248)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v22)+244)) = v201
	*(*int32)(unsafe.Add(mBase, uint32(v22)+240)) = v67
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_21), v22+int32(240))
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2697), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L4
	} else {
		goto L193
	}
L193:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L194:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v762)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+224)) = v763 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_22), v22+int32(224))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+220)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v22)+216)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v22)+212)) = v772
	*(*int32)(unsafe.Add(mBase, uint32(v22)+208)) = v67
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_23), v22+int32(208))
	mBase = m.M
	v781 = m.ExcPending
	if v781 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2723), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L199:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(v794)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+112)) = v795 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_24), v22+int32(112))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
	;
	v804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v805 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+108)) = uint32(v805)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+100)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v22)+96)) = v804
	v810 = int64(base.Ui64(v805) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+104)) = uint32(v810)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_12), v22+int32(96))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2316), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L4
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(v829)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v830 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_25), v22+int32(48))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	v839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v840 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+44)) = uint32(v840)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v839
	v845 = int64(base.Ui64(v840) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+40)) = uint32(v845)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_12), v22+int32(32))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2346), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v864)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = v865 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_26), v22+int32(80))
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L4
	} else {
		goto L211
	}
L211:
	;
	v874 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v875 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+76)) = uint32(v875)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+68)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v874
	v880 = int64(base.Ui64(v875) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v22)+72)) = uint32(v880)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_12), v22-int32(-64))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2358), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L4
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bt_metap(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int64
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int64
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 float64
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
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
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	v6 = m.G0
	v8 = v6 - int32(192)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = F_superuser(m)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15 != 0 {
				v17 = F_textToQualifiedNameList(m, v11)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = F_makeRangeVarFromNameList(m, v17)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v22 = F_relation_openrv(m, v19, int32(1))
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
							v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+119)))
							if v25 != int32(105) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int32(0)
									} else {
										v190 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+132)) = int32(_a_F_bt_metap_0)
										*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = v190 + int32(4)
										F_errmsg(m, int32(_a_F_bt_metap_1), v8+int32(128))
										mBase = m.M
										v200 = m.ExcPending
										if v200 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_bt_metap_2), int32(864), int32(_a_F_bt_metap_3))
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
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
								v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+84))
								if v28 != int32(403) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v186 = m.ExcPending
									if v186 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return int32(0)
										} else {
											v190 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+132)) = int32(_a_F_bt_metap_0)
											*(*int32)(unsafe.Add(mBase, uint32(v8)+128)) = v190 + int32(4)
											F_errmsg(m, int32(_a_F_bt_metap_1), v8+int32(128))
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_bt_metap_2), int32(864), int32(_a_F_bt_metap_3))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
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
									v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+118)))
									if v31 == int32(116) {
										v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
										if v34 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(_a_F_bt_metap_4), int32(0))
													mBase = m.M
													v216 = m.ExcPending
													if v216 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_bt_metap_2), int32(874), int32(_a_F_bt_metap_3))
														mBase = m.M
														v221 = m.ExcPending
														if v221 != 0 {
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
											v38 = F_ReadBuffer(m, v22, int32(0))
											mBase = m.M
											v39 = m.ExcPending
											if v39 != 0 {
												return int32(0)
											} else {
												F_LockBuffer(m, v38, int32(1))
												mBase = m.M
												v42 = m.ExcPending
												if v42 != 0 {
													return int32(0)
												} else {
													if v38 < int32(0) {
														v46 = *(*int32)(unsafe.Add(mBase, _c_F_bt_metap[0]))
														v52 = *(*int32)(unsafe.Add(mBase, uint32(v46+(v38^int32(-1))<<(uint(int32(2))%32))))
														v60 = v52
													} else {
														v54 = *(*int32)(unsafe.Add(mBase, _c_F_bt_metap[1]))
														v60 = v54 + v38<<(uint(int32(13))%32) + int32(-8192)
													}
													v64 = F_get_call_result_type(m, l0, int32(0), v8+int32(188))
													mBase = m.M
													v65 = m.ExcPending
													if v65 != 0 {
														return int32(0)
													} else {
														if v64 != int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v225 = m.ExcPending
															if v225 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_bt_metap_5), int32(0))
																mBase = m.M
																v229 = m.ExcPending
																if v229 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_metap_2), int32(884), int32(_a_F_bt_metap_3))
																	mBase = m.M
																	v234 = m.ExcPending
																	if v234 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
															if v69 <= int32(8) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v238 = m.ExcPending
																if v238 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(50724996))
																	mBase = m.M
																	v241 = m.ExcPending
																	if v241 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(_a_F_bt_metap_6), int32(0))
																		mBase = m.M
																		v245 = m.ExcPending
																		if v245 != 0 {
																			return int32(0)
																		} else {
																			F_errhint(m, int32(_a_F_bt_metap_7), int32(0))
																			mBase = m.M
																			v249 = m.ExcPending
																			if v249 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_metap_2), int32(898), int32(_a_F_bt_metap_3))
																				mBase = m.M
																				v254 = m.ExcPending
																				if v254 != 0 {
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
																v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
																*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v72
																v77 = F_psprintf(m, int32(_a_F_bt_metap_8), v8+int32(112))
																mBase = m.M
																v78 = m.ExcPending
																if v78 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v77
																	v80 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v80
																	v85 = F_psprintf(m, int32(_a_F_bt_metap_8), v8+int32(96))
																	mBase = m.M
																	v86 = m.ExcPending
																	if v86 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = v85
																		v88 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+32)))
																		*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = v88
																		v93 = F_psprintf(m, int32(_a_F_bt_metap_9), v8+int32(80))
																		mBase = m.M
																		v94 = m.ExcPending
																		if v94 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v8)+152)) = v93
																			v96 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+36)))
																			*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v96
																			v101 = F_psprintf(m, int32(_a_F_bt_metap_9), v8-int32(-64))
																			mBase = m.M
																			v102 = m.ExcPending
																			if v102 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v8)+156)) = v101
																				v104 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+40)))
																				*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v104
																				v109 = F_psprintf(m, int32(_a_F_bt_metap_9), v8+int32(48))
																				mBase = m.M
																				v110 = m.ExcPending
																				if v110 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = v109
																					v112 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+44)))
																					*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v112
																					v117 = F_psprintf(m, int32(_a_F_bt_metap_9), v8+int32(32))
																					mBase = m.M
																					v118 = m.ExcPending
																					if v118 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = v117
																						v120 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
																						if base.Ui32(int32(3)) <= base.Ui32(v120) {
																							v123 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+48)))
																							*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v123
																							v128 = F_psprintf(m, int32(_a_F_bt_metap_9), v8+int32(16))
																							mBase = m.M
																							v129 = m.ExcPending
																							if v129 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = v128
																								v131 = *(*float64)(unsafe.Add(mBase, uint32(v60)+56))
																								*(*float64)(unsafe.Add(mBase, uint32(v8))) = v131
																								v134 = F_psprintf(m, int32(_a_F_bt_metap_10), v8)
																								mBase = m.M
																								v135 = m.ExcPending
																								if v135 != 0 {
																									return int32(0)
																								} else {
																									v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+64)))
																									if v138 != 0 {
																										v139 = int32(_a_F_bt_metap_11)
																									} else {
																										v139 = int32(_a_F_bt_metap_12)
																									}
																									v144 = v134
																									v145 = v139
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+176)) = v145
																									*(*int32)(unsafe.Add(mBase, uint32(v8)+172)) = v144
																									v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
																									v149 = F_TupleDescGetAttInMetadata(m, v148)
																									mBase = m.M
																									v150 = m.ExcPending
																									if v150 != 0 {
																										return int32(0)
																									} else {
																										v153 = F_BuildTupleFromCStrings(m, v149, v8+int32(144))
																										mBase = m.M
																										v154 = m.ExcPending
																										if v154 != 0 {
																											return int32(0)
																										} else {
																											v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
																											v156 = F_HeapTupleHeaderGetDatum(m, v155)
																											mBase = m.M
																											v157 = m.ExcPending
																											if v157 != 0 {
																												return int32(0)
																											} else {
																												F_UnlockReleaseBuffer(m, v38)
																												mBase = m.M
																												v159 = m.ExcPending
																												if v159 != 0 {
																													return int32(0)
																												} else {
																													F_relation_close(m, v22, int32(1))
																													mBase = m.M
																													v162 = m.ExcPending
																													if v162 != 0 {
																														return int32(0)
																													} else {
																														m.G0 = v8 + int32(192)
																														return v156
																													}
																												}
																											}
																										}
																									}
																								}
																							}
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = int32(_a_F_bt_metap_13)
																							v144 = int32(_a_F_bt_metap_14)
																							v145 = int32(_a_F_bt_metap_12)
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+176)) = v145
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+172)) = v144
																							v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
																							v149 = F_TupleDescGetAttInMetadata(m, v148)
																							mBase = m.M
																							v150 = m.ExcPending
																							if v150 != 0 {
																								return int32(0)
																							} else {
																								v153 = F_BuildTupleFromCStrings(m, v149, v8+int32(144))
																								mBase = m.M
																								v154 = m.ExcPending
																								if v154 != 0 {
																									return int32(0)
																								} else {
																									v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
																									v156 = F_HeapTupleHeaderGetDatum(m, v155)
																									mBase = m.M
																									v157 = m.ExcPending
																									if v157 != 0 {
																										return int32(0)
																									} else {
																										F_UnlockReleaseBuffer(m, v38)
																										mBase = m.M
																										v159 = m.ExcPending
																										if v159 != 0 {
																											return int32(0)
																										} else {
																											F_relation_close(m, v22, int32(1))
																											mBase = m.M
																											v162 = m.ExcPending
																											if v162 != 0 {
																												return int32(0)
																											} else {
																												m.G0 = v8 + int32(192)
																												return v156
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
														}
													}
												}
											}
										}
									} else {
										v38 = F_ReadBuffer(m, v22, int32(0))
										mBase = m.M
										v39 = m.ExcPending
										if v39 != 0 {
											return int32(0)
										} else {
											F_LockBuffer(m, v38, int32(1))
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return int32(0)
											} else {
												if v38 < int32(0) {
													v46 = *(*int32)(unsafe.Add(mBase, _c_F_bt_metap[0]))
													v52 = *(*int32)(unsafe.Add(mBase, uint32(v46+(v38^int32(-1))<<(uint(int32(2))%32))))
													v60 = v52
												} else {
													v54 = *(*int32)(unsafe.Add(mBase, _c_F_bt_metap[1]))
													v60 = v54 + v38<<(uint(int32(13))%32) + int32(-8192)
												}
												v64 = F_get_call_result_type(m, l0, int32(0), v8+int32(188))
												mBase = m.M
												v65 = m.ExcPending
												if v65 != 0 {
													return int32(0)
												} else {
													if v64 != int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_bt_metap_5), int32(0))
															mBase = m.M
															v229 = m.ExcPending
															if v229 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_bt_metap_2), int32(884), int32(_a_F_bt_metap_3))
																mBase = m.M
																v234 = m.ExcPending
																if v234 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v68 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
														if v69 <= int32(8) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v238 = m.ExcPending
															if v238 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(50724996))
																mBase = m.M
																v241 = m.ExcPending
																if v241 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(_a_F_bt_metap_6), int32(0))
																	mBase = m.M
																	v245 = m.ExcPending
																	if v245 != 0 {
																		return int32(0)
																	} else {
																		F_errhint(m, int32(_a_F_bt_metap_7), int32(0))
																		mBase = m.M
																		v249 = m.ExcPending
																		if v249 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(_a_F_bt_metap_2), int32(898), int32(_a_F_bt_metap_3))
																			mBase = m.M
																			v254 = m.ExcPending
																			if v254 != 0 {
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
															v72 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
															*(*int32)(unsafe.Add(mBase, uint32(v8)+112)) = v72
															v77 = F_psprintf(m, int32(_a_F_bt_metap_8), v8+int32(112))
															mBase = m.M
															v78 = m.ExcPending
															if v78 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v8)+144)) = v77
																v80 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
																*(*int32)(unsafe.Add(mBase, uint32(v8)+96)) = v80
																v85 = F_psprintf(m, int32(_a_F_bt_metap_8), v8+int32(96))
																mBase = m.M
																v86 = m.ExcPending
																if v86 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v8)+148)) = v85
																	v88 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+32)))
																	*(*int64)(unsafe.Add(mBase, uint32(v8)+80)) = v88
																	v93 = F_psprintf(m, int32(_a_F_bt_metap_9), v8+int32(80))
																	mBase = m.M
																	v94 = m.ExcPending
																	if v94 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v8)+152)) = v93
																		v96 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+36)))
																		*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = v96
																		v101 = F_psprintf(m, int32(_a_F_bt_metap_9), v8-int32(-64))
																		mBase = m.M
																		v102 = m.ExcPending
																		if v102 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v8)+156)) = v101
																			v104 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+40)))
																			*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v104
																			v109 = F_psprintf(m, int32(_a_F_bt_metap_9), v8+int32(48))
																			mBase = m.M
																			v110 = m.ExcPending
																			if v110 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v8)+160)) = v109
																				v112 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+44)))
																				*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v112
																				v117 = F_psprintf(m, int32(_a_F_bt_metap_9), v8+int32(32))
																				mBase = m.M
																				v118 = m.ExcPending
																				if v118 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v8)+164)) = v117
																					v120 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
																					if base.Ui32(int32(3)) <= base.Ui32(v120) {
																						v123 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v60)+48)))
																						*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v123
																						v128 = F_psprintf(m, int32(_a_F_bt_metap_9), v8+int32(16))
																						mBase = m.M
																						v129 = m.ExcPending
																						if v129 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = v128
																							v131 = *(*float64)(unsafe.Add(mBase, uint32(v60)+56))
																							*(*float64)(unsafe.Add(mBase, uint32(v8))) = v131
																							v134 = F_psprintf(m, int32(_a_F_bt_metap_10), v8)
																							mBase = m.M
																							v135 = m.ExcPending
																							if v135 != 0 {
																								return int32(0)
																							} else {
																								v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+64)))
																								if v138 != 0 {
																									v139 = int32(_a_F_bt_metap_11)
																								} else {
																									v139 = int32(_a_F_bt_metap_12)
																								}
																								v144 = v134
																								v145 = v139
																								*(*int32)(unsafe.Add(mBase, uint32(v8)+176)) = v145
																								*(*int32)(unsafe.Add(mBase, uint32(v8)+172)) = v144
																								v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
																								v149 = F_TupleDescGetAttInMetadata(m, v148)
																								mBase = m.M
																								v150 = m.ExcPending
																								if v150 != 0 {
																									return int32(0)
																								} else {
																									v153 = F_BuildTupleFromCStrings(m, v149, v8+int32(144))
																									mBase = m.M
																									v154 = m.ExcPending
																									if v154 != 0 {
																										return int32(0)
																									} else {
																										v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
																										v156 = F_HeapTupleHeaderGetDatum(m, v155)
																										mBase = m.M
																										v157 = m.ExcPending
																										if v157 != 0 {
																											return int32(0)
																										} else {
																											F_UnlockReleaseBuffer(m, v38)
																											mBase = m.M
																											v159 = m.ExcPending
																											if v159 != 0 {
																												return int32(0)
																											} else {
																												F_relation_close(m, v22, int32(1))
																												mBase = m.M
																												v162 = m.ExcPending
																												if v162 != 0 {
																													return int32(0)
																												} else {
																													m.G0 = v8 + int32(192)
																													return v156
																												}
																											}
																										}
																									}
																								}
																							}
																						}
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+168)) = int32(_a_F_bt_metap_13)
																						v144 = int32(_a_F_bt_metap_14)
																						v145 = int32(_a_F_bt_metap_12)
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+176)) = v145
																						*(*int32)(unsafe.Add(mBase, uint32(v8)+172)) = v144
																						v148 = *(*int32)(unsafe.Add(mBase, uint32(v8)+188))
																						v149 = F_TupleDescGetAttInMetadata(m, v148)
																						mBase = m.M
																						v150 = m.ExcPending
																						if v150 != 0 {
																							return int32(0)
																						} else {
																							v153 = F_BuildTupleFromCStrings(m, v149, v8+int32(144))
																							mBase = m.M
																							v154 = m.ExcPending
																							if v154 != 0 {
																								return int32(0)
																							} else {
																								v155 = *(*int32)(unsafe.Add(mBase, uint32(v153)+16))
																								v156 = F_HeapTupleHeaderGetDatum(m, v155)
																								mBase = m.M
																								v157 = m.ExcPending
																								if v157 != 0 {
																									return int32(0)
																								} else {
																									F_UnlockReleaseBuffer(m, v38)
																									mBase = m.M
																									v159 = m.ExcPending
																									if v159 != 0 {
																										return int32(0)
																									} else {
																										F_relation_close(m, v22, int32(1))
																										mBase = m.M
																										v162 = m.ExcPending
																										if v162 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v8 + int32(192)
																											return v156
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
				v170 = m.ExcPending
				if v170 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_bt_metap_15), int32(0))
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_bt_metap_2), int32(855), int32(_a_F_bt_metap_3))
							mBase = m.M
							v182 = m.ExcPending
							if v182 != 0 {
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
func F_bt_page_print_tuples(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v291 int32
	_ = v291
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v349 int32
	_ = v349
	var v364 int32
	_ = v364
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
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v24 = v18 + v19<<(uint(int32(2))%32) + int32(20)
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L17
	} else {
		goto L72
	}
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v28 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v28)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = base.I32_extend16_s(v19)
	v36 = v18 + v27&int32(_a_F_bt_page_print_tuples_0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v36
	v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = int32(base.Ui32(v38) >> (uint(int32(15)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = int32(base.Ui32(v38)>>(uint(int32(14))%32)) & int32(1)
	v48 = v38 & int32(_a_F_bt_page_print_tuples_1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v48
	if v28 <= base.I32_extend16_s(v38) {
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
	v381 = m.ExcPending
	if v381 != 0 {
		goto L17
	} else {
		goto L69
	}
L5:
	;
	if base.Ui32(int32(_a_F_bt_page_print_tuples_2)) <= base.Ui32(v76) {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v55 = int32(8)
	goto L8
L7:
	;
	v55 = int32(16)
	goto L8
L8:
	;
	v56 = v48 - v55
	if v38&int32(_a_F_bt_page_print_tuples_2) == int32(0) {
		v76 = v56
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	if v61&int32(_a_F_bt_page_print_tuples_2) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	v76 = v64 | v65<<(uint(int32(16))%32) - v55
	goto L5
L11:
	;
	goto L12
L12:
	;
	if v61&int32(_a_F_bt_page_print_tuples_3) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v74 = v56 - int32(8)
	goto L15
L14:
	;
	v74 = v56
	goto L15
L15:
	;
	v76 = v74
	goto L5
L16:
	;
	v79 = int32(1)
	v84 = F_palloc0(m, v76*int32(3)+v79)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	if v76 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v141 = F_cstring_to_text(m, v84)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L17
	} else {
		goto L27
	}
L20:
	;
	v90 = v36 + v55
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v91
	v96 = F_pg_sprintf(m, v84, int32(_a_F_bt_page_print_tuples_4), v16+int32(32))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v76 == int32(1) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v101 = v84
	v104 = v79
	goto L23
L23:
	;
	v113 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)) = uint8(v113)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104+v90))))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v116
	v119 = v101 + int32(3)
	v123 = F_pg_sprintf(m, v119, int32(_a_F_bt_page_print_tuples_4), v16+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L17
	} else {
		goto L25
	}
L24:
	;
	goto L19
L25:
	;
	v126 = v104 + int32(1)
	if v126 != v76 {
		v101 = v119
		v104 = v126
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+84)) = v141
	F_pfree(m, v84)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v146 = int32(1)
	v149 = v26 & (base.B2i32(v19 != v146) | v25)
	if v149&v146 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v369 = F_heap_form_tuple(m, v364, v16-int32(-64), v16+int32(48))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L17
	} else {
		goto L67
	}
L30:
	;
	v349 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+56)) = uint8(v349)
	goto L29
L31:
	;
	v344 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+55)) = uint8(v344)
	goto L30
L32:
	;
	v170 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	if v170&int32(_a_F_bt_page_print_tuples_2) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v153 = int32(_a_F_bt_page_print_tuples_5)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = base.B2i32(v152&v153 == v153)
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	if v158&int32(_a_F_bt_page_print_tuples_2) != 0 {
		v169 = v158
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v162 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+54)) = uint8(v162)
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
	if v164&int32(_a_F_bt_page_print_tuples_2) == int32(0) {
		goto L31
	} else {
		goto L37
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v36
	goto L30
L37:
	;
	v169 = v164
	goto L32
L38:
	;
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+4)))
	if v205&int32(_a_F_bt_page_print_tuples_2) == int32(0) {
		goto L30
	} else {
		goto L51
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v187 + v36 + v186
	goto L38
L40:
	;
	v198 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+55)) = uint8(v198)
	goto L38
L41:
	;
	v188 = int32(1)
	if v170&int32(_a_F_bt_page_print_tuples_2) != 0 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	if v170&int32(_a_F_bt_page_print_tuples_3) == int32(0) {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	v186 = v182
	v187 = v183 << (uint(int32(16)) % 32)
	goto L41
L45:
	;
	v186 = int32(-6)
	v187 = v169 & int32(_a_F_bt_page_print_tuples_1)
	goto L41
L46:
	;
	v193 = v149 & v188
	goto L48
L47:
	;
	v193 = v188
	goto L48
L48:
	;
	if v193 == int32(0) {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	if v18 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	goto L40
L51:
	;
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
	v213 = v205 & int32(4095)
	v216 = F_palloc(m, v213<<(uint(int32(2))%32))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	if v213 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v339 = F_construct_array_builtin(m, v216, v213, int32(27))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L17
	} else {
		goto L65
	}
L54:
	;
	v223 = v36 + v210<<(uint(int32(16))%32) + v211
	v225 = v213 & int32(3)
	v226 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v213) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v237 = v226
	v242 = int32(0)
	goto L58
L56:
	;
	v291 = v226
	goto L57
L57:
	;
	v304 = v291
	v305 = v226
	goto L62
L58:
	;
	v246 = int32(2)
	v249 = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v216+v237<<(uint(v246)%32)))) = v223 + v237*v249
	v254 = v237 | int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v216+v254<<(uint(v246)%32)))) = v223 + v254*v249
	v263 = v237 | v246
	*(*int32)(unsafe.Add(mBase, uint32(v216+v263<<(uint(v246)%32)))) = v223 + v263*v249
	v272 = v237 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v216+v272<<(uint(v246)%32)))) = v223 + v272*v249
	v280 = int32(4)
	v281 = v237 + v280
	v283 = v242 + v280
	if v283 != v213&int32(4092) {
		v237 = v281
		v242 = v283
		goto L58
	} else {
		goto L60
	}
L59:
	;
	if v225 == int32(0) {
		goto L53
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v291 = v281
	goto L57
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216+v304<<(uint(int32(2))%32)))) = v223 + v304*int32(6)
	v320 = int32(1)
	v323 = v305 + v320
	if v323 != v225 {
		v304 = v304 + v320
		v305 = v323
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L53
L64:
	;
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+96)) = v339
	F_pfree(m, v216)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L17
	} else {
		goto L66
	}
L66:
	;
	goto L29
L67:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v369)+16))
	v372 = F_HeapTupleHeaderGetDatum(m, v371)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L17
	} else {
		goto L68
	}
L68:
	;
	m.G0 = v16 + int32(112)
	return v372
L69:
	;
	F_errmsg_internal(m, int32(_a_F_bt_page_print_tuples_6), int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L17
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_bt_page_print_tuples_7), int32(503), int32(_a_F_bt_page_print_tuples_8))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L17
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
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v76
	F_errmsg_internal(m, int32(_a_F_bt_page_print_tuples_9), v16)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L17
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_bt_page_print_tuples_7), int32(548), int32(_a_F_bt_page_print_tuples_8))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L17
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bt_page_stats_1_9(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_stats_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
