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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+52))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v18
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
	v22 = l1 + int32(28)
	v25 = F__bt_check_compare(m, l0, v15, l3, l4, v17, l2, v20, v22, v13+int32(12))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		if l2 == int32(0) {
			v125 = v25
			m.G0 = v13 + int32(16)
			return v125
		} else {
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
			if v31 != 0 {
				v125 = v25
				m.G0 = v13 + int32(16)
				return v125
			} else {
				v32 = int32(0)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				v36 = F__bt_tuple_before_array_skeys(m, l0, v15, l3, v17, l4, int32(1), v34, v32)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					if v36 != 0 {
						v38 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+28)) = uint8(v38)
						v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)))
						v42 = v40 + v38
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v42)
						if base.I32_extend16_s(v42) < int32(3) {
							v125 = v32
							m.G0 = v13 + int32(16)
							return v125
						} else {
							v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+24)))
							v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							if base.Ui32(v47) < base.Ui32(v48) {
								v125 = v32
								m.G0 = v13 + int32(16)
								return v125
							} else {
								v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v50 == int32(1) {
									v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
									if v47 < v53-int32(5) {
										v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+32)))
										if v62 != 0 {
											if int32(203) < v62 {
												v70 = v62
											} else {
												v68 = v62 << (uint(int32(1)) % 32)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v68)
												v70 = v68
											}
										} else {
											v68 = int32(5)
											*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v68)
											v70 = v68
										}
										if v50 == int32(1) {
											v74 = base.I32_extend16_s(v70) + v47
											v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
											if v74 < v75 {
												v77 = v74
											} else {
												v77 = v75
											}
											v84 = v77
										} else {
											v79 = v47 - base.I32_extend16_s(v70)
											if v48 < v79 {
												v81 = v79
											} else {
												v81 = v48
											}
											v84 = v81
										}
										v85 = int32(0)
										v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v86+v84&int32(_a_F__bt_checkkeys_0)<<(uint(int32(2))%32))+20))
										v99 = F__bt_tuple_before_array_skeys(m, l0, v50, v86+v92&int32(_a_F__bt_checkkeys_1), v17, l4, v85, v85, v85)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											if v99 != 0 {
												if v50 == int32(1) {
													v104 = v84 + int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v104)
													v125 = v85
												} else {
													v107 = v84 - int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v107)
													v125 = v85
												}
											} else {
												v109 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v109)
												v111 = int32(15)
												v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+32)))
												if v112 <= v111 {
													v115 = v111
												} else {
													v115 = v112
												}
												v117 = int32(base.Ui32(v115) >> (uint(int32(3)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v117)
												v125 = v85
											}
											m.G0 = v13 + int32(16)
											return v125
										}
									} else {
										v125 = v32
										m.G0 = v13 + int32(16)
										return v125
									}
								} else {
									if v50 != int32(-1) {
										v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+32)))
										if v62 != 0 {
											if int32(203) < v62 {
												v70 = v62
											} else {
												v68 = v62 << (uint(int32(1)) % 32)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v68)
												v70 = v68
											}
										} else {
											v68 = int32(5)
											*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v68)
											v70 = v68
										}
										if v50 == int32(1) {
											v74 = base.I32_extend16_s(v70) + v47
											v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
											if v74 < v75 {
												v77 = v74
											} else {
												v77 = v75
											}
											v84 = v77
										} else {
											v79 = v47 - base.I32_extend16_s(v70)
											if v48 < v79 {
												v81 = v79
											} else {
												v81 = v48
											}
											v84 = v81
										}
										v85 = int32(0)
										v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v86+v84&int32(_a_F__bt_checkkeys_0)<<(uint(int32(2))%32))+20))
										v99 = F__bt_tuple_before_array_skeys(m, l0, v50, v86+v92&int32(_a_F__bt_checkkeys_1), v17, l4, v85, v85, v85)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											if v99 != 0 {
												if v50 == int32(1) {
													v104 = v84 + int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v104)
													v125 = v85
												} else {
													v107 = v84 - int32(1)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v107)
													v125 = v85
												}
											} else {
												v109 = int32(0)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v109)
												v111 = int32(15)
												v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+32)))
												if v112 <= v111 {
													v115 = v111
												} else {
													v115 = v112
												}
												v117 = int32(base.Ui32(v115) >> (uint(int32(3)) % 32))
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v117)
												v125 = v85
											}
											m.G0 = v13 + int32(16)
											return v125
										}
									} else {
										if base.Ui32(v47) <= base.Ui32(v48+int32(5)) {
											v125 = v32
											m.G0 = v13 + int32(16)
											return v125
										} else {
											v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+32)))
											if v62 != 0 {
												if int32(203) < v62 {
													v70 = v62
												} else {
													v68 = v62 << (uint(int32(1)) % 32)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v68)
													v70 = v68
												}
											} else {
												v68 = int32(5)
												*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v68)
												v70 = v68
											}
											if v50 == int32(1) {
												v74 = base.I32_extend16_s(v70) + v47
												v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
												if v74 < v75 {
													v77 = v74
												} else {
													v77 = v75
												}
												v84 = v77
											} else {
												v79 = v47 - base.I32_extend16_s(v70)
												if v48 < v79 {
													v81 = v79
												} else {
													v81 = v48
												}
												v84 = v81
											}
											v85 = int32(0)
											v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											v92 = *(*int32)(unsafe.Add(mBase, uint32(v86+v84&int32(_a_F__bt_checkkeys_0)<<(uint(int32(2))%32))+20))
											v99 = F__bt_tuple_before_array_skeys(m, l0, v50, v86+v92&int32(_a_F__bt_checkkeys_1), v17, l4, v85, v85, v85)
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
												return int32(0)
											} else {
												if v99 != 0 {
													if v50 == int32(1) {
														v104 = v84 + int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v104)
														v125 = v85
													} else {
														v107 = v84 - int32(1)
														*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v107)
														v125 = v85
													}
												} else {
													v109 = int32(0)
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+30)) = uint16(v109)
													v111 = int32(15)
													v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+32)))
													if v112 <= v111 {
														v115 = v111
													} else {
														v115 = v112
													}
													v117 = int32(base.Ui32(v115) >> (uint(int32(3)) % 32))
													*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v117)
													v125 = v85
												}
												m.G0 = v13 + int32(16)
												return v125
											}
										}
									}
								}
							}
						}
					} else {
						v120 = F__bt_advance_array_keys(m, l0, l1, l3, l4, v17, v34, int32(1))
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							v125 = v120
							m.G0 = v13 + int32(16)
							return v125
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
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v129 int64
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
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v160 int64
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
						F_errfinish(m, int32(_a_F__bt_compare_scankey_args_3), int32(952), int32(_a_F__bt_compare_scankey_args_4))
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
									v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
									v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
									v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
										v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
										v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
										v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
						v127 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
						v128 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
						v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int64(0)))
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
												v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
												v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
												v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
													v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
													v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
													v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
													mBase = m.M
													v161 = m.ExcPending
													if v161 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
									v127 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
									v128 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
									v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int64(0)))
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
														v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
														v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
														v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
														mBase = m.M
														v161 = m.ExcPending
														if v161 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
															v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
															v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
															v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return int32(0)
															} else {
																*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
											v127 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
											v128 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
											v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
											mBase = m.M
											v130 = m.ExcPending
											if v130 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int64(0)))
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
											v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
											v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
											v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
											mBase = m.M
											v161 = m.ExcPending
											if v161 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
												v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
												v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
												v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
								v127 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
								v128 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
								v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int64(0)))
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
												v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
												v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
												v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
												mBase = m.M
												v161 = m.ExcPending
												if v161 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
													v158 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
													v159 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
													v160 = F_OidFunctionCall2Coll(m, v153, v157, v158, v159)
													mBase = m.M
													v161 = m.ExcPending
													if v161 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v160 != int64(0)))
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
									v127 = *(*int64)(unsafe.Add(mBase, uint32(l2)+48))
									v128 = *(*int64)(unsafe.Add(mBase, uint32(l3)+48))
									v129 = F_FunctionCall2Coll(m, l1+int32(16), v126, v127, v128)
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v129 != int64(0)))
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
						F_errfinish(m, int32(_a_F__bt_dedup_finish_pending_4), int32(575), int32(_a_F__bt_dedup_finish_pending_5))
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
	var v98 int32
	_ = v98
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
	var v133 int32
	_ = v133
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
	var v152 int32
	_ = v152
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
	v40 = F__bt_getbuf(m, l0, v38, int32(3))
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
			v65 = F__bt_getbuf(m, l0, int32(0), int32(3))
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
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v89+(l2^int32(-1))*int32(56))+16))
					v104 = v95
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
					v98 = int32(56)
					v103 = *(*int32)(unsafe.Add(mBase, uint32(v97+l2*v98-v98)+16))
					v104 = v103
				}
				F_UnlockReleaseBuffer(m, v65)
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
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+(l2^int32(-1))*int32(56))+16))
								v139 = v130
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
								v133 = int32(56)
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v132+l2*v133-v133)+16))
								v139 = v138
							}
							if v40 < int32(0) {
								v143 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[2]))
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+(v40^int32(-1))*int32(56))+16))
								v158 = v149
							} else {
								v151 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
								v152 = int32(56)
								v157 = *(*int32)(unsafe.Add(mBase, uint32(v151+v40*v152-v152)+16))
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
								F_errfinish(m, int32(_a_F__bt_finish_split_1), int32(2314), int32(_a_F__bt_finish_split_2))
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
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+(l2^int32(-1))*int32(56))+16))
						v139 = v130
					} else {
						v132 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
						v133 = int32(56)
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v132+l2*v133-v133)+16))
						v139 = v138
					}
					if v40 < int32(0) {
						v143 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[2]))
						v149 = *(*int32)(unsafe.Add(mBase, uint32(v143+(v40^int32(-1))*int32(56))+16))
						v158 = v149
					} else {
						v151 = *(*int32)(unsafe.Add(mBase, _c_F__bt_finish_split[3]))
						v152 = int32(56)
						v157 = *(*int32)(unsafe.Add(mBase, uint32(v151+v40*v152-v152)+16))
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
						F_errfinish(m, int32(_a_F__bt_finish_split_1), int32(2314), int32(_a_F__bt_finish_split_2))
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
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
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
						v52 = v52 + int32(56)
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v247 int64
	_ = v247
	var v249 int64
	_ = v249
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int64
	_ = v272
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v280 int64
	_ = v280
	var v282 int64
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	goto L3
L1:
	;
	m.G0 = v12 - int32(-64)
	return v364
L2:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v84)+20))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v84)+16))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v269 = F_MemoryContextAlloc(m, v267, int32(48))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L77
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
	v100 = F__bt_allocbuf(m, l0, l1)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L8
	} else {
		goto L41
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
	v77 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L26
	}
L8:
	;
	return int32(0)
L9:
	;
	F_LockBufferInternal(m, v28, int32(1))
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
	F_UnlockReleaseBuffer(m, v28)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
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
		v364 = v28
		goto L1
	} else {
		goto L20
	}
L20:
	;
	goto L12
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v68 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	F_pfree(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = int32(0)
	goto L7
L25:
	;
	goto L24
L26:
	;
	F_LockBufferInternal(m, v77, int32(1))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	F__bt_checkpage(m, l0, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v84 = F__bt_getmeta(m, l0, v77)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v86 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	if base.B2i32(l2 != int32(1)) == int32(0) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	F_UnlockReleaseBuffer(m, v77)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	F_UnlockBuffer(m, v77)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L35
	}
L34:
	;
	v364 = int32(0)
	goto L1
L35:
	;
	F_LockBufferInternal(m, v77, int32(3))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L8
	} else {
		goto L36
	}
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v84)+8))
	if v97 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_UnlockReleaseBuffer(m, v77)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L8
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	goto L4
L40:
	;
	goto L3
L41:
	;
	if v100 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	if v100 < int32(0) {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[2]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105+(v100^int32(-1))*int32(56))+16))
	v120 = v111
	goto L42
L44:
	;
	goto L45
L45:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[3]))
	v114 = int32(56)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v113+v100*v114-v114)+16))
	v120 = v119
	goto L42
L46:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+16)))
	v140 = v139 + v138
	v141 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v140))) = v141
	*(*int64)(unsafe.Add(mBase, uint32(v140)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v140)+12)) = int32(3)
	if v77 < v141 {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[0]))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124+(v100^int32(-1))<<(uint(int32(2))%32))))
	v138 = v130
	goto L46
L48:
	;
	goto L49
L49:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[1]))
	v138 = v132 + v100<<(uint(int32(13))%32) + int32(-8192)
	goto L46
L50:
	;
	v165 = int32(_a_F__bt_getroot_0)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[4])) = v167 + int32(1)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	if base.Ui32(v171) <= base.Ui32(int32(2)) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[0]))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150+(v77^int32(-1))<<(uint(int32(2))%32))))
	v164 = v156
	goto L50
L52:
	;
	goto L53
L53:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[1]))
	v164 = v158 + v77<<(uint(int32(13))%32) + int32(-8192)
	goto L50
L54:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v164)+64)) = uint8(v174)
	*(*int64)(unsafe.Add(mBase, uint32(v164)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v164)+48)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v164)+28)) = int32(3)
	v182 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v164)+12)) = uint16(v182)
	goto L56
L55:
	;
	goto L56
L56:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v84)+32)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(v84)+20)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v84)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = v120
	F_MarkBufferDirty(m, v100)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L8
	} else {
		goto L57
	}
L57:
	;
	F_MarkBufferDirty(m, v77)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L8
	} else {
		goto L58
	}
L58:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+118)))
	if v197 != int32(112) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v249 = base.I64_rotl(v247, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v138))) = v249
	*(*int64)(unsafe.Add(mBase, uint32(v164))) = v249
	v252 = int32(_a_F__bt_getroot_0)
	v254 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[4])) = v254 - int32(1)
	F_UnlockBuffer(m, v100)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L74
	}
L60:
	;
	v244 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L8
	} else {
		goto L73
	}
L61:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[5]))
	if v201 <= int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v204 != 0 {
		goto L60
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L8
	} else {
		goto L67
	}
L65:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v205 != 0 {
		goto L60
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	F_XLogRegisterBuffer(m, int32(0), v100, int32(6))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	F_XLogRegisterBuffer(m, int32(2), v77, int32(14))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v216
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+52)) = uint8(v224)
	F_XLogRegisterBufData(m, int32(2), v10+int32(-36), int32(28))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v120
	F_XLogRegisterData(m, v10+int32(-8), int32(8))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	v242 = F_XLogInsert(m, int32(11), int32(160))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v247 = v242
	goto L59
L73:
	;
	v247 = v244
	goto L59
L74:
	;
	F_LockBufferInternal(m, v100, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	F_UnlockReleaseBuffer(m, v77)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	v364 = v100
	goto L1
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v269
	v272 = *(*int64)(unsafe.Add(mBase, uint32(v84)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v269)+40)) = v272
	v274 = *(*int64)(unsafe.Add(mBase, uint32(v84)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v269)+32)) = v274
	v276 = *(*int64)(unsafe.Add(mBase, uint32(v84)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v269)+24)) = v276
	v278 = *(*int64)(unsafe.Add(mBase, uint32(v84)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v269)+16)) = v278
	v280 = *(*int64)(unsafe.Add(mBase, uint32(v84)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v269)+8)) = v280
	v282 = *(*int64)(unsafe.Add(mBase, uint32(v84)))
	*(*int64)(unsafe.Add(mBase, uint32(v269))) = v282
	v286 = v77
	v290 = v266
	goto L79
L78:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	if v340 == v265 {
		v364 = v294
		goto L1
	} else {
		goto L91
	}
L79:
	;
	v294 = F__bt_relandgetbuf(m, l0, v286, v290, int32(1))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L8
	} else {
		goto L82
	}
L80:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L8
	} else {
		goto L88
	}
L81:
	;
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313)+16)))
	v315 = v314 + v313
	v316 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+12)))
	if v316&int32(20) == int32(0) {
		goto L78
	} else {
		goto L86
	}
L82:
	;
	if v294 < int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v299 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[0]))
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v299+(v294^int32(-1))<<(uint(int32(2))%32))))
	v313 = v305
	goto L81
L84:
	;
	goto L85
L85:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F__bt_getroot[1]))
	v313 = v307 + v294<<(uint(int32(13))%32) + int32(-8192)
	goto L81
L86:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v315)+4))
	if v321 != 0 {
		v286 = v294
		v290 = v321
		goto L79
	} else {
		goto L87
	}
L87:
	;
	goto L80
L88:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v326 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_getroot_1), v10+int32(-48))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F__bt_getroot_2), int32(553), int32(_a_F__bt_getroot_3))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L8
	} else {
		goto L90
	}
L90:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v265
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v347
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v346 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_getroot_4), v12)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L8
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F__bt_getroot_2), int32(560), int32(_a_F__bt_getroot_3))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
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
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
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
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+60)))
	v50 = v34 + int32(1)
	if v46 == v50&int32(_a_F__bt_mark_scankey_required_0) {
		v33 = v33 + int32(56)
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
	F_errfinish(m, int32(_a_F__bt_mark_scankey_required_2), int32(799), int32(_a_F__bt_mark_scankey_required_3))
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
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v126 int64
	_ = v126
	var v130 int64
	_ = v130
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
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
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v172 int64
	_ = v172
	var v183 int64
	_ = v183
	var v185 int64
	_ = v185
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v195 int64
	_ = v195
	var v197 int64
	_ = v197
	var v201 int64
	_ = v201
	var v203 int64
	_ = v203
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
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
						v54 = int32(0)
						v56 = int32(_a_F__bt_parallel_build_main_1)
						v57 = base.AtomicRmwOr32(m, v54, v56, v54)
						*(*int64)(unsafe.Add(mBase, uint32(v29)+392)) = v25
						v62 = base.AtomicRmwOr32(m, v54, v56, v54)
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = v63 + v47
						v69 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[3]))
						*(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[3])) = v69 - v47
					}
				}
			}
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v24 != 0 {
				v76 = int32(4)
			} else {
				v76 = int32(5)
			}
			v77 = F_table_open(m, v73, v76)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				if v24 != 0 {
					v82 = int32(3)
				} else {
					v82 = int32(8)
				}
				v83 = F_index_open(m, v79, v82)
				mBase = m.M
				v84 = m.ExcPending
				if v84 != 0 {
					return
				} else {
					v86 = F_palloc0(m, int32(16))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v83
						*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v77
						v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
						*(*uint8)(unsafe.Add(mBase, uint32(v86)+12)) = uint8(v90)
						v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
						*(*uint8)(unsafe.Add(mBase, uint32(v86)+13)) = uint8(v92)
						v96 = F_shm_toc_lookup(m, l1, int64(-6917529027641081854), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							F_tuplesort_attach_shared(m, v96, l0)
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
								if v100 == int32(1) {
									v104 = F_palloc0(m, int32(16))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return
									} else {
										v106 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v104)+4)) = v106
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v86)+8))
										v109 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v104)+12)) = uint8(v109)
										*(*int32)(unsafe.Add(mBase, uint32(v104)+8)) = v108
										v114 = F_shm_toc_lookup(m, l1, int64(-6917529027641081853), v109)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											F_tuplesort_attach_shared(m, v114, l0)
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return
											} else {
												v118 = v104
												v119 = v114
												base.MemoryCopy(m, int32(_a_F__bt_parallel_build_main_2), int32(_a_F__bt_parallel_build_main_3), int32(128))
												v126 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[4]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[5])) = v126
												v130 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[6]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[7])) = v130
												v134 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[8]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[9])) = v134
												v138 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[10]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[11])) = v138
												v142 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[12]))
												*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[13])) = v142
												v145 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[14]))
												v146 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
												v147 = base.I32_div_s(v145, v146)
												F__bt_parallel_scan_and_sort(m, v86, v118, v22, v96, v119, v147, int32(0))
												mBase = m.M
												v150 = m.ExcPending
												if v150 != 0 {
													return
												} else {
													v153 = F_shm_toc_lookup(m, l1, int64(-6917529027641081850), int32(0))
													mBase = m.M
													v154 = m.ExcPending
													if v154 != 0 {
														return
													} else {
														v157 = F_shm_toc_lookup(m, l1, int64(-6917529027641081851), int32(0))
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return
														} else {
															v160 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[15]))
															v163 = v153 + v160<<(uint(int32(7))%32)
															v166 = v157 + v160*int32(40)
															base.MemoryFill(m, v163, int32(0), int32(128))
															F_BufferUsageAccumDiff(m, v163, int32(_a_F__bt_parallel_build_main_2))
															mBase = m.M
															v172 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v166)+32)) = v172
															*(*int64)(unsafe.Add(mBase, uint32(v166)+24)) = v172
															*(*int64)(unsafe.Add(mBase, uint32(v166)+16)) = v172
															*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v172
															*(*int64)(unsafe.Add(mBase, uint32(v166))) = v172
															v183 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[8]))
															v185 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[9]))
															*(*int64)(unsafe.Add(mBase, uint32(v166)+16)) = v183 - v185
															v189 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[12]))
															v191 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[13]))
															*(*int64)(unsafe.Add(mBase, uint32(v166))) = v189 - v191
															v195 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[10]))
															v197 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[11]))
															*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v195 - v197
															v201 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[6]))
															v203 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[7]))
															*(*int64)(unsafe.Add(mBase, uint32(v166)+24)) = v201 - v203
															v207 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[4]))
															v209 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[5]))
															*(*int64)(unsafe.Add(mBase, uint32(v166)+32)) = v207 - v209
															F_relation_close(m, v83, v82)
															mBase = m.M
															v213 = m.ExcPending
															if v213 != 0 {
																return
															} else {
																F_relation_close(m, v77, v76)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
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
									v118 = v3
									v119 = v3
									base.MemoryCopy(m, int32(_a_F__bt_parallel_build_main_2), int32(_a_F__bt_parallel_build_main_3), int32(128))
									v126 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[4]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[5])) = v126
									v130 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[6]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[7])) = v130
									v134 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[8]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[9])) = v134
									v138 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[10]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[11])) = v138
									v142 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[12]))
									*(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[13])) = v142
									v145 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[14]))
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
									v147 = base.I32_div_s(v145, v146)
									F__bt_parallel_scan_and_sort(m, v86, v118, v22, v96, v119, v147, int32(0))
									mBase = m.M
									v150 = m.ExcPending
									if v150 != 0 {
										return
									} else {
										v153 = F_shm_toc_lookup(m, l1, int64(-6917529027641081850), int32(0))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return
										} else {
											v157 = F_shm_toc_lookup(m, l1, int64(-6917529027641081851), int32(0))
											mBase = m.M
											v158 = m.ExcPending
											if v158 != 0 {
												return
											} else {
												v160 = *(*int32)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[15]))
												v163 = v153 + v160<<(uint(int32(7))%32)
												v166 = v157 + v160*int32(40)
												base.MemoryFill(m, v163, int32(0), int32(128))
												F_BufferUsageAccumDiff(m, v163, int32(_a_F__bt_parallel_build_main_2))
												mBase = m.M
												v172 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v166)+32)) = v172
												*(*int64)(unsafe.Add(mBase, uint32(v166)+24)) = v172
												*(*int64)(unsafe.Add(mBase, uint32(v166)+16)) = v172
												*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v172
												*(*int64)(unsafe.Add(mBase, uint32(v166))) = v172
												v183 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[8]))
												v185 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[9]))
												*(*int64)(unsafe.Add(mBase, uint32(v166)+16)) = v183 - v185
												v189 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[12]))
												v191 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[13]))
												*(*int64)(unsafe.Add(mBase, uint32(v166))) = v189 - v191
												v195 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[10]))
												v197 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[11]))
												*(*int64)(unsafe.Add(mBase, uint32(v166)+8)) = v195 - v197
												v201 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[6]))
												v203 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[7]))
												*(*int64)(unsafe.Add(mBase, uint32(v166)+24)) = v201 - v203
												v207 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[4]))
												v209 = *(*int64)(unsafe.Add(mBase, _c_F__bt_parallel_build_main[5]))
												*(*int64)(unsafe.Add(mBase, uint32(v166)+32)) = v207 - v209
												F_relation_close(m, v83, v82)
												mBase = m.M
												v213 = m.ExcPending
												if v213 != 0 {
													return
												} else {
													F_relation_close(m, v77, v76)
													mBase = m.M
													v215 = m.ExcPending
													if v215 != 0 {
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
				F_errfinish(m, int32(_a_F__bt_skiparray_shrink_1), int32(1340), int32(_a_F__bt_skiparray_shrink_2))
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
	F_UnlockReleaseBuffer(m, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L9
	} else {
		goto L31
	}
L6:
	;
	v47 = F__bt_relandgetbuf(m, l0, v41, v44, int32(3))
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
	v85 = F__bt_relandgetbuf(m, l0, int32(0), v44, int32(3))
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
	F_errfinish(m, int32(_a_F__bt_stepright_1), int32(1078), int32(_a_F__bt_stepright_2))
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
	var v42 int64
	_ = v42
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
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
		v66 = v18
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v77 = m.G0
	v79 = v77 - int32(288)
	m.G0 = v79
	if v66 < v18 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v22 = int32(1)
	if v18 <= int32(0) {
		v66 = v22
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
	v66 = v18 + int32(1)
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
		v66 = v29
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
		v32 = v32 + int32(56)
		goto L4
	} else {
		goto L15
	}
L13:
	;
	if base.I32_wrap_i64(v58) != 0 {
		v66 = v29
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
	m.G0 = v79 + int32(288)
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)))
	if v126&int32(_a_F__bt_truncate_0) == int32(0) {
		v146 = v126
		goto L31
	} else {
		goto L32
	}
L17:
	;
	v82 = v66
	goto L19
L18:
	;
	v82 = v18
	goto L19
L19:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v82 == v83 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	v87 = v85 & int32(_a_F__bt_truncate_1)
	v88 = F_palloc(m, v87)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L6
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v93 = F_CreateTupleDescTruncatedCopy(m, v16, v82)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	if v87 == int32(0) {
		v120 = v88
		goto L16
	} else {
		goto L24
	}
L24:
	;
	base.MemoryCopy(m, v88, l2, v87)
	v120 = v88
	goto L16
L25:
	;
	v96 = v79 + int32(32)
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	if int32(0) <= base.I32_extend16_s(v99) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v103 = int32(8)
	goto L28
L27:
	;
	v103 = int32(16)
	goto L28
L28:
	;
	F_index_deform_tuple_internal(m, v93, v96, v79, l2+v103, l2+int32(8), int32(base.Ui32(v99)>>(uint(int32(15))%32)))
	mBase = m.M
	v111 = *(*int32)(unsafe.Add(mBase, _c_F__bt_truncate[0]))
	v112 = F_index_form_tuple_context(m, v93, v96, v79, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v112)+4)) = uint16(v114)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v116
	F_pfree(m, v93)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v120 = v112
	goto L16
L31:
	;
	if v66 <= v18 {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)))
	if v131&int32(32) == int32(0) {
		v146 = v126
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v137 = v126 & int32(_a_F__bt_truncate_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)) = uint16(v137)
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v144 = (v139+int32(7))&int32(-8200) | v137
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)) = uint16(v144)
	v146 = v144
	goto L31
L34:
	;
	m.G0 = v14 + int32(16)
	return v226
L35:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)) = uint16(v66)
	v150 = v146 | int32(_a_F__bt_truncate_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)) = uint16(v150)
	v226 = v120
	goto L34
L36:
	;
	goto L37
L37:
	;
	v159 = (v146&int32(_a_F__bt_truncate_1)+int32(7))&int32(_a_F__bt_truncate_3) + int32(8)
	v160 = F_palloc0(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+6)))
	v168 = (v162&int32(_a_F__bt_truncate_1) + int32(7)) & int32(_a_F__bt_truncate_3)
	if v168 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	base.MemoryCopy(m, v160, v120, v168)
	goto L41
L40:
	;
	goto L41
L41:
	;
	F_pfree(m, v120)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v173 = v18 | int32(_a_F__bt_truncate_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v160)+4)) = uint16(v173)
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+6)))
	v179 = int32(_a_F__bt_truncate_0)
	v180 = v159 | v175&int32(_a_F__bt_truncate_5) | v179
	*(*uint16)(unsafe.Add(mBase, uint32(v160)+6)) = uint16(v180)
	if v18&v179 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v196 = v194 + (v160 + v193)
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v197&int32(32) == int32(0) {
		v220 = l1
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v193 = v159 & int32(_a_F__bt_truncate_6)
	v194 = int32(-6)
	goto L43
L45:
	;
	goto L46
L46:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160))))
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160)+2)))
	v193 = v189 << (uint(int32(16)) % 32)
	v194 = v192
	goto L43
L47:
	;
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+4)) = uint16(v222)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v224
	v226 = v160
	goto L34
L48:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v202&int32(_a_F__bt_truncate_0) == int32(0) {
		v220 = l1
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v215 = int32(6)
	v220 = v207 + (l1 + v208<<(uint(int32(16))%32)) + v202&int32(4095)*v215 - v215
	goto L47
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
	var v64 int64
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
	var v114 int32
	_ = v114
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
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v216 int32
	_ = v216
	var v243 int32
	_ = v243
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
	return v243
L5:
	;
	v243 = int32(0)
	goto L4
L6:
	;
	v26 = v24
	v32 = l6
	goto L7
L7:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v45 = v42 + v32*int32(56)
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
	v216 = v32 + int32(1)
	if v216 < v199 {
		v26 = v199
		v32 = v216
		goto L7
	} else {
		goto L61
	}
L15:
	;
	if l5 == int32(0) {
		v199 = v26
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
	v185 = int32(1)
	v188 = int32(0)
	if base.B2i32(l1 == v185)&base.B2i32(v169 < v188)|base.B2i32(l1 == int32(-1))&base.B2i32(v188 < v169) != 0 {
		v243 = v185
		goto L4
	} else {
		goto L59
	}
L22:
	;
	v71 = int32(0)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v73 <= v71 {
		v114 = v71
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
	F__bt_binsrch_skiparray_skey(m, int32(0), l1, v64, v119, v114, v45, v19+int32(8))
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
		v114 = v96
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v114 = v96
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
		v169 = v124
		goto L21
	} else {
		goto L32
	}
L32:
	;
	v243 = v71
	goto L4
L33:
	;
	v169 = int32(1)
	goto L21
L34:
	;
	if v160&int32(_a_F__bt_tuple_before_array_skeys_1) != 0 {
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
	v160 = v68
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
	v169 = v136
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
	v147 = *(*int64)(unsafe.Add(mBase, uint32(v45)+48))
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
	v169 = v141
	goto L21
L50:
	;
	v150 = base.I32_wrap_i64(v148)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v151&int32(16777216) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v150 < int32(0) {
		goto L33
	} else {
		goto L54
	}
L52:
	;
	v158 = v150
	goto L53
L53:
	;
	if v158 != 0 {
		v169 = v158
		goto L21
	} else {
		goto L55
	}
L54:
	;
	v158 = int32(0) - v150
	goto L53
L55:
	;
	v160 = v151
	goto L34
L56:
	;
	v169 = int32(-1)
	goto L21
L57:
	;
	goto L58
L58:
	;
	v169 = int32(base.Ui32(v160)>>(uint(int32(22))%32)) & int32(1)
	goto L21
L59:
	;
	if v169|l5 != 0 {
		goto L5
	} else {
		goto L60
	}
L60:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v199 = v198
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
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
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
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
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v501 int32
	_ = v501
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int64
	_ = v583
	var v588 int64
	_ = v588
	var v594 int32
	_ = v594
	var v599 int32
	_ = v599
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v670 int32
	_ = v670
	var v674 int64
	_ = v674
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v734 int32
	_ = v734
	var v739 int64
	_ = v739
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v774 int64
	_ = v774
	var v780 int32
	_ = v780
	var v785 int32
	_ = v785
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int64
	_ = v804
	var v809 int64
	_ = v809
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v827 int32
	_ = v827
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int64
	_ = v839
	var v844 int64
	_ = v844
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v859 int32
	_ = v859
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int64
	_ = v874
	var v879 int64
	_ = v879
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(384)
	m.G0 = v20
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if base.Ui32((l1-int32(1))&int32(_a_F_bt_child_highkey_check_0)) <= base.Ui32(int32(2047)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = F_PageGetItemIdCareful_2(m, l0, v31, v32, l1)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v46 = v5
	goto L3
L3:
	;
	v48 = base.B2i32(v24 != int32(-1))
	if v24 != int32(-1) {
		goto L18
	} else {
		goto L19
	}
L4:
	;
	return
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v39 = v35 + v36&int32(_a_F_bt_child_highkey_check_1)
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	v43 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+2)))
	v46 = v40<<(uint(int32(16))%32) | v43
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L4
	} else {
		goto L209
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L4
	} else {
		goto L204
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L4
	} else {
		goto L199
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L4
	} else {
		goto L194
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L4
	} else {
		goto L189
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L4
	} else {
		goto L184
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L4
	} else {
		goto L179
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L4
	} else {
		goto L175
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L4
	} else {
		goto L170
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L4
	} else {
		goto L165
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v545 = m.ExcPending
	if v545 != 0 {
		goto L4
	} else {
		goto L161
	}
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)) = uint8(v527)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v525
	m.G0 = v20 + int32(384)
	return
L18:
	;
	v49 = v24
	goto L20
L19:
	;
	v49 = v46
	goto L20
L20:
	;
	if v49|v46 == int32(0) {
		v525 = int32(-1)
		v527 = v5
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v55 = int32(1)
	v56 = l3 - v55
	v64 = base.B2i32(v49 == int32(0))
	v65 = v49
	v66 = v55
	v69 = v48 & v23
	goto L22
L22:
	;
	if v64&int32(1) != 0 {
		goto L16
	} else {
		goto L24
	}
L23:
	;
	v525 = int32(-1)
	v527 = v515
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
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+16)))
	v83 = v82 + v81
	v85 = v66 & int32(1)
	if v85 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L26:
	;
	if v65 == v46 {
		v81 = l2
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v79 = F_palloc_btree_page(m, l0, v65)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v81 = v79
	goto L25
L31:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+12)))
	if v95&int32(260) != int32(4) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v88 != int32(-1) {
		goto L31
	} else {
		goto L33
	}
L33:
	;
	v91 = F_bt_leftmost_ignoring_half_dead(m, l0, v65, v83)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v91 == int32(0) {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	goto L31
L36:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	if v100 != v56 {
		goto L14
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	if v85 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v65 == v104 {
		goto L13
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	if v65 == v106 {
		goto L13
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	v108 = base.B2i32(v65 == v46)
	if v108|v95&int32(20) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+12)))
	v287 = v285 & int32(128)
	v289 = int32(base.Ui32(v287) >> (uint(int32(7)) % 32))
	if v289|v285&int32(16) != 0 {
		goto L89
	} else {
		goto L90
	}
L46:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v81)+16)))
	v113 = v81 + v112
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+12)))
	if v114&int32(2) != 0 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v81)))
	v119 = base.I64_rotl(v117, int64(32))
	if v69 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v122 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	if v114&int32(1) != 0 {
		goto L12
	} else {
		goto L57
	}
L51:
	;
	if v122 == int32(0) {
		goto L45
	} else {
		goto L52
	}
L52:
	;
	F_errcode(m, int32(128))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v130 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_child_highkey_check_2), v20+int32(160))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+144)) = uint32(v119)
	v143 = int64(base.Ui64(v119) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+140)) = uint32(v143)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v65
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_3), v20+int32(128))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2610), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	goto L45
L57:
	;
	v162 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	if v162 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = v165 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_child_highkey_check_6), v20+int32(272))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v113)+4))
	if v182 != 0 {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2635), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v183 = int32(2)
	goto L66
L65:
	;
	v183 = int32(1)
	goto L66
L66:
	;
	v184 = F_PageGetItemIdCareful_2(m, l0, v65, v81, v183)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v189 = v81 + v186&int32(_a_F_bt_child_highkey_check_1)
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189))))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189)+2)))
	v198 = v179
	v200 = v190<<(uint(int32(16))%32) | v193
	goto L68
L68:
	;
	v213 = *(*int32)(unsafe.Add(mBase, _c_F_bt_child_highkey_check[0]))
	if v213 != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v220&int32(4) != 0 {
		goto L10
	} else {
		goto L84
	}
L70:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v216 = F_palloc_btree_page(m, l0, v200)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L4
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216)+16)))
	v219 = v216 + v218
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+12)))
	if v220&int32(1) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v226 = v198 - int32(1)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	if v226 != v227 {
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
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v231 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v232 = int32(2)
	goto L81
L80:
	;
	v232 = int32(1)
	goto L81
L81:
	;
	v233 = F_PageGetItemIdCareful_2(m, l0, v200, v216, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	v238 = v216 + v235&int32(_a_F_bt_child_highkey_check_1)
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238)+2)))
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v238))))
	F_pfree(m, v216)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v198 = v227
	v200 = v239 | v240<<(uint(int32(16))%32)
	goto L68
L84:
	;
	if v220&int32(16) == int32(0) {
		goto L9
	} else {
		goto L85
	}
L85:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	if v252 == int32(0) {
		goto L9
	} else {
		goto L86
	}
L86:
	;
	v256 = F_PageGetItemIdCareful_2(m, l0, v200, v216, int32(1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v256)))
	v261 = v216 + v258&int32(_a_F_bt_child_highkey_check_1)
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261))))
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+2)))
	if v262<<(uint(int32(16))%32)|v265 != v65 {
		goto L9
	} else {
		goto L88
	}
L88:
	;
	goto L45
L89:
	;
	if v65 == v46 {
		goto L153
	} else {
		goto L154
	}
L90:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v293 == int32(0) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v297 = F_PageGetItemIdCareful_2(m, l0, v65, v81, int32(1))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L4
	} else {
		goto L92
	}
L92:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v297)))
	v300 = l1 + v108
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301)+16)))
	v303 = v301 + v302
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303)+12)))
	if v304&int32(1) == int32(0) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	v353 = v81 + v299&int32(_a_F_bt_child_highkey_check_1)
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v353)+6)))
	v355 = int32(_a_F_bt_child_highkey_check_7)
	v356 = v354 & v355
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v349)+6)))
	if v356 != v357&v355 {
		goto L6
	} else {
		goto L111
	}
L94:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v345 == int32(0) {
		goto L7
	} else {
		goto L110
	}
L95:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v313 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v318) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	v314 = int32(2)
	goto L100
L99:
	;
	v314 = int32(1)
	goto L100
L100:
	;
	if v300&int32(_a_F_bt_child_highkey_check_0) == v314 {
		goto L94
	} else {
		goto L101
	}
L101:
	;
	goto L97
L102:
	;
	v326 = int32(base.Ui32(v318+int32(_a_F_bt_child_highkey_check_8)) >> (uint(int32(2)) % 32))
	goto L104
L103:
	;
	v326 = int32(0)
	goto L104
L104:
	;
	if base.Ui32(v326&int32(_a_F_bt_child_highkey_check_0)) < base.Ui32(v300&int32(_a_F_bt_child_highkey_check_0)) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v331 == int32(0) {
		goto L8
	} else {
		goto L108
	}
L106:
	;
	v334 = v300
	goto L107
L107:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v338 = F_PageGetItemIdCareful_2(m, l0, v335, v301, v334&int32(_a_F_bt_child_highkey_check_0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L4
	} else {
		goto L109
	}
L108:
	;
	v334 = int32(1)
	goto L107
L109:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	v349 = v340 + v341&int32(_a_F_bt_child_highkey_check_1)
	goto L93
L110:
	;
	v349 = v345
	goto L93
L111:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v361 == int32(1) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v364 = int32(4)
	v365 = v353 + v364
	v367 = v349 + v364
	v369 = v356 - v364
	if base.Ui32(v364) <= base.Ui32(v369) {
		goto L118
	} else {
		goto L119
	}
L113:
	;
	goto L114
L114:
	;
	v434 = int32(6)
	v435 = v353 + v434
	v437 = v349 + v434
	v439 = v356 - v434
	if base.Ui32(int32(4)) <= base.Ui32(v439) {
		goto L137
	} else {
		goto L138
	}
L115:
	;
	if v431 == int32(0) {
		goto L89
	} else {
		goto L133
	}
L116:
	;
	v431 = int32(0)
	goto L115
L117:
	;
	v405 = v400
	v406 = v401
	v407 = v402
	goto L127
L118:
	;
	if (v365|v367)&int32(3) != 0 {
		v400 = v365
		v401 = v367
		v402 = v369
		goto L117
	} else {
		goto L121
	}
L119:
	;
	v393 = v365
	v394 = v367
	v395 = v369
	goto L120
L120:
	;
	if v395 == int32(0) {
		goto L116
	} else {
		goto L126
	}
L121:
	;
	v377 = v365
	v378 = v367
	v379 = v369
	goto L122
L122:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v377)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	if v382 != v383 {
		v400 = v377
		v401 = v378
		v402 = v379
		goto L117
	} else {
		goto L124
	}
L123:
	;
	v393 = v388
	v394 = v386
	v395 = v390
	goto L120
L124:
	;
	v385 = int32(4)
	v386 = v378 + v385
	v388 = v377 + v385
	v390 = v379 - v385
	if base.Ui32(int32(3)) < base.Ui32(v390) {
		v377 = v388
		v378 = v386
		v379 = v390
		goto L122
	} else {
		goto L125
	}
L125:
	;
	goto L123
L126:
	;
	v400 = v393
	v401 = v394
	v402 = v395
	goto L117
L127:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v405))))
	v411 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if v410 == v411 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v431 = v410 - v411
	goto L115
L129:
	;
	v413 = int32(1)
	v418 = v407 - v413
	if v418 != 0 {
		v405 = v405 + v413
		v406 = v406 + v413
		v407 = v418
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
	if v501 != 0 {
		goto L6
	} else {
		goto L152
	}
L135:
	;
	v501 = int32(0)
	goto L134
L136:
	;
	v475 = v470
	v476 = v471
	v477 = v472
	goto L146
L137:
	;
	if (v435|v437)&int32(3) != 0 {
		v470 = v435
		v471 = v437
		v472 = v439
		goto L136
	} else {
		goto L140
	}
L138:
	;
	v463 = v435
	v464 = v437
	v465 = v439
	goto L139
L139:
	;
	if v465 == int32(0) {
		goto L135
	} else {
		goto L145
	}
L140:
	;
	v447 = v435
	v448 = v437
	v449 = v439
	goto L141
L141:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v448)))
	if v452 != v453 {
		v470 = v447
		v471 = v448
		v472 = v449
		goto L136
	} else {
		goto L143
	}
L142:
	;
	v463 = v458
	v464 = v456
	v465 = v460
	goto L139
L143:
	;
	v455 = int32(4)
	v456 = v448 + v455
	v458 = v447 + v455
	v460 = v449 - v455
	if base.Ui32(int32(3)) < base.Ui32(v460) {
		v447 = v458
		v448 = v456
		v449 = v460
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v470 = v463
	v471 = v464
	v472 = v465
	goto L136
L146:
	;
	v480 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
	v481 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v476))))
	if v480 == v481 {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	v501 = v480 - v481
	goto L134
L148:
	;
	v483 = int32(1)
	v488 = v477 - v483
	if v488 != 0 {
		v475 = v475 + v483
		v476 = v476 + v483
		v477 = v488
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
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	v525 = v510
	v527 = base.B2i32(v287 != int32(0))
	goto L17
L154:
	;
	goto L155
L155:
	;
	v511 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if l2 != v81 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	F_pfree(m, v81)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L4
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v515 = int32(0)
	if v511|v46 != 0 {
		v64 = base.B2i32(v511 == v515)
		v65 = v511
		v66 = v515
		v69 = v289
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
	v548 = m.ExcPending
	if v548 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v549)+48))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v551
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v550 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_9), v20)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2211), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
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
	v571 = m.ExcPending
	if v571 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v572)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+368)) = v573 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_11), v20+int32(368))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v583 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+364)) = uint32(v583)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+356)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v20)+352)) = v582
	v588 = int64(base.Ui64(v583) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+360)) = uint32(v588)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_12), v20+int32(352))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L4
	} else {
		goto L168
	}
L168:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2230), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
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
	v606 = m.ExcPending
	if v606 != 0 {
		goto L4
	} else {
		goto L171
	}
L171:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v607)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+336)) = v608 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_13), v20+int32(336))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L4
	} else {
		goto L172
	}
L172:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+328)) = v617
	*(*int32)(unsafe.Add(mBase, uint32(v20)+324)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v20)+320)) = v65
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_14), v20+int32(320))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L4
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2240), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
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
	v637 = m.ExcPending
	if v637 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v639 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_15), v20+int32(16))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L4
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2247), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
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
	v660 = m.ExcPending
	if v660 != 0 {
		goto L4
	} else {
		goto L180
	}
L180:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v661)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+304)) = v662 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_16), v20+int32(304))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+288)) = v65
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+296)) = uint32(v119)
	v674 = int64(base.Ui64(v119) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+292)) = uint32(v674)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_17), v20+int32(288))
	mBase = m.M
	v680 = m.ExcPending
	if v680 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2631), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
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
	v692 = m.ExcPending
	if v692 != 0 {
		goto L4
	} else {
		goto L185
	}
L185:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v693)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v694 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_child_highkey_check_18), v20+int32(192))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L4
	} else {
		goto L186
	}
L186:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v219)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+188)) = v703
	*(*int32)(unsafe.Add(mBase, uint32(v20)+184)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v20)+180)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v65
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_19), v20+int32(176))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L4
	} else {
		goto L187
	}
L187:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2659), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
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
	v724 = m.ExcPending
	if v724 != 0 {
		goto L4
	} else {
		goto L190
	}
L190:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v726 = *(*int32)(unsafe.Add(mBase, uint32(v725)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v726 + int32(4)
	F_errmsg_internal(m, int32(_a_F_bt_child_highkey_check_20), v20+int32(256))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L4
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v65
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+252)) = uint32(v119)
	v739 = int64(base.Ui64(v119) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+248)) = uint32(v739)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_21), v20+int32(240))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L4
	} else {
		goto L192
	}
L192:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2697), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
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
	v758 = m.ExcPending
	if v758 != 0 {
		goto L4
	} else {
		goto L195
	}
L195:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v759)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v760 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_22), v20+int32(224))
	mBase = m.M
	v768 = m.ExcPending
	if v768 != 0 {
		goto L4
	} else {
		goto L196
	}
L196:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v113)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v65
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+220)) = uint32(v119)
	v774 = int64(base.Ui64(v119) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+216)) = uint32(v774)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_23), v20+int32(208))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L197
	}
L197:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2723), int32(_a_F_bt_child_highkey_check_5))
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
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
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L200
	}
L200:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v793)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v794 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_24), v20+int32(112))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L4
	} else {
		goto L201
	}
L201:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v804 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+108)) = uint32(v804)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+100)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v803
	v809 = int64(base.Ui64(v804) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+104)) = uint32(v809)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_12), v20+int32(96))
	mBase = m.M
	v815 = m.ExcPending
	if v815 != 0 {
		goto L4
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2316), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
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
	v827 = m.ExcPending
	if v827 != 0 {
		goto L4
	} else {
		goto L205
	}
L205:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v829 = *(*int32)(unsafe.Add(mBase, uint32(v828)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v829 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_25), v20+int32(48))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L4
	} else {
		goto L206
	}
L206:
	;
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v839 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+44)) = uint32(v839)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v838
	v844 = int64(base.Ui64(v839) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+40)) = uint32(v844)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_12), v20+int32(32))
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L4
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2346), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
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
	v862 = m.ExcPending
	if v862 != 0 {
		goto L4
	} else {
		goto L210
	}
L210:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v863)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v864 + int32(4)
	F_errmsg(m, int32(_a_F_bt_child_highkey_check_26), v20+int32(80))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L4
	} else {
		goto L211
	}
L211:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v874 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+76)) = uint32(v874)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+68)) = v65
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v873
	v879 = int64(base.Ui64(v874) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v20)+72)) = uint32(v879)
	F_errdetail_internal(m, int32(_a_F_bt_child_highkey_check_12), v20-int32(-64))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L4
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_bt_child_highkey_check_4), int32(2358), int32(_a_F_bt_child_highkey_check_10))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
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
func F_bt_metap(m *base.Module, l0 int32) int64 {
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
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 float64
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
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
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	v7 = m.G0
	v9 = v7 - int32(192)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = F_superuser(m)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			if v16 != 0 {
				v18 = F_textToQualifiedNameList(m, v12)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = F_makeRangeVarFromNameList(m, v18)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int64(0)
					} else {
						v23 = F_relation_openrv(m, v20, int32(1))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int64(0)
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
							v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+119)))
							if v26 != int32(105) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v187 = m.ExcPending
								if v187 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(151027844))
									mBase = m.M
									v190 = m.ExcPending
									if v190 != 0 {
										return int64(0)
									} else {
										v191 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = int32(_a_F_bt_metap_0)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v191 + int32(4)
										F_errmsg(m, int32(_a_F_bt_metap_1), v9+int32(128))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_bt_metap_2), int32(865), int32(_a_F_bt_metap_3))
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
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
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+84))
								if v29 != int32(403) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v187 = m.ExcPending
									if v187 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(151027844))
										mBase = m.M
										v190 = m.ExcPending
										if v190 != 0 {
											return int64(0)
										} else {
											v191 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+132)) = int32(_a_F_bt_metap_0)
											*(*int32)(unsafe.Add(mBase, uint32(v9)+128)) = v191 + int32(4)
											F_errmsg(m, int32(_a_F_bt_metap_1), v9+int32(128))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_bt_metap_2), int32(865), int32(_a_F_bt_metap_3))
												mBase = m.M
												v206 = m.ExcPending
												if v206 != 0 {
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
									v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+118)))
									if v32 == int32(116) {
										v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)))
										if v35 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v210 = m.ExcPending
											if v210 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v213 = m.ExcPending
												if v213 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_bt_metap_4), int32(0))
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_bt_metap_2), int32(875), int32(_a_F_bt_metap_3))
														mBase = m.M
														v222 = m.ExcPending
														if v222 != 0 {
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
											v39 = F_ReadBuffer(m, v23, int32(0))
											mBase = m.M
											v40 = m.ExcPending
											if v40 != 0 {
												return int64(0)
											} else {
												F_LockBufferInternal(m, v39, int32(1))
												mBase = m.M
												v43 = m.ExcPending
												if v43 != 0 {
													return int64(0)
												} else {
													if v39 < int32(0) {
														v47 = *(*int32)(unsafe.Add(mBase, _c_F_bt_metap[0]))
														v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v39^int32(-1))<<(uint(int32(2))%32))))
														v61 = v53
													} else {
														v55 = *(*int32)(unsafe.Add(mBase, _c_F_bt_metap[1]))
														v61 = v55 + v39<<(uint(int32(13))%32) + int32(-8192)
													}
													v65 = F_get_call_result_type(m, l0, int32(0), v9+int32(188))
													mBase = m.M
													v66 = m.ExcPending
													if v66 != 0 {
														return int64(0)
													} else {
														if v65 != int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v226 = m.ExcPending
															if v226 != 0 {
																return int64(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_bt_metap_5), int32(0))
																mBase = m.M
																v230 = m.ExcPending
																if v230 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_metap_2), int32(885), int32(_a_F_bt_metap_3))
																	mBase = m.M
																	v235 = m.ExcPending
																	if v235 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
															v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
															if v70 <= int32(8) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v239 = m.ExcPending
																if v239 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(50724996))
																	mBase = m.M
																	v242 = m.ExcPending
																	if v242 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_bt_metap_6), int32(0))
																		mBase = m.M
																		v246 = m.ExcPending
																		if v246 != 0 {
																			return int64(0)
																		} else {
																			F_errhint(m, int32(_a_F_bt_metap_7), int32(0))
																			mBase = m.M
																			v250 = m.ExcPending
																			if v250 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_bt_metap_2), int32(899), int32(_a_F_bt_metap_3))
																				mBase = m.M
																				v255 = m.ExcPending
																				if v255 != 0 {
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
															} else {
																v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v73
																v78 = F_psprintf(m, int32(_a_F_bt_metap_8), v9+int32(112))
																mBase = m.M
																v79 = m.ExcPending
																if v79 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v78
																	v81 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v81
																	v86 = F_psprintf(m, int32(_a_F_bt_metap_8), v9+int32(96))
																	mBase = m.M
																	v87 = m.ExcPending
																	if v87 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v86
																		v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v89
																		v94 = F_psprintf(m, int32(_a_F_bt_metap_9), v9+int32(80))
																		mBase = m.M
																		v95 = m.ExcPending
																		if v95 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v94
																			v97 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v97
																			v102 = F_psprintf(m, int32(_a_F_bt_metap_9), v9-int32(-64))
																			mBase = m.M
																			v103 = m.ExcPending
																			if v103 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+156)) = v102
																				v105 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v105
																				v110 = F_psprintf(m, int32(_a_F_bt_metap_9), v9+int32(48))
																				mBase = m.M
																				v111 = m.ExcPending
																				if v111 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v110
																					v113 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
																					*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v113
																					v118 = F_psprintf(m, int32(_a_F_bt_metap_9), v9+int32(32))
																					mBase = m.M
																					v119 = m.ExcPending
																					if v119 != 0 {
																						return int64(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v118
																						v121 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
																						if base.Ui32(int32(3)) <= base.Ui32(v121) {
																							v124 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
																							v129 = F_psprintf(m, int32(_a_F_bt_metap_9), v9+int32(16))
																							mBase = m.M
																							v130 = m.ExcPending
																							if v130 != 0 {
																								return int64(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v9)+168)) = v129
																								v132 = *(*float64)(unsafe.Add(mBase, uint32(v61)+56))
																								*(*float64)(unsafe.Add(mBase, uint32(v9))) = v132
																								v135 = F_psprintf(m, int32(_a_F_bt_metap_10), v9)
																								mBase = m.M
																								v136 = m.ExcPending
																								if v136 != 0 {
																									return int64(0)
																								} else {
																									v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+64)))
																									if v139 != 0 {
																										v140 = int32(_a_F_bt_metap_11)
																									} else {
																										v140 = int32(_a_F_bt_metap_12)
																									}
																									v145 = v135
																									v146 = v140
																									*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v146
																									*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v145
																									v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
																									v150 = F_TupleDescGetAttInMetadata(m, v149)
																									mBase = m.M
																									v151 = m.ExcPending
																									if v151 != 0 {
																										return int64(0)
																									} else {
																										v154 = F_BuildTupleFromCStrings(m, v150, v9+int32(144))
																										mBase = m.M
																										v155 = m.ExcPending
																										if v155 != 0 {
																											return int64(0)
																										} else {
																											v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
																											v157 = F_HeapTupleHeaderGetDatum(m, v156)
																											mBase = m.M
																											v158 = m.ExcPending
																											if v158 != 0 {
																												return int64(0)
																											} else {
																												F_UnlockReleaseBuffer(m, v39)
																												mBase = m.M
																												v160 = m.ExcPending
																												if v160 != 0 {
																													return int64(0)
																												} else {
																													F_relation_close(m, v23, int32(1))
																													mBase = m.M
																													v163 = m.ExcPending
																													if v163 != 0 {
																														return int64(0)
																													} else {
																														m.G0 = v9 + int32(192)
																														return v157
																													}
																												}
																											}
																										}
																									}
																								}
																							}
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+168)) = int32(_a_F_bt_metap_13)
																							v145 = int32(_a_F_bt_metap_14)
																							v146 = int32(_a_F_bt_metap_12)
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v146
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v145
																							v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
																							v150 = F_TupleDescGetAttInMetadata(m, v149)
																							mBase = m.M
																							v151 = m.ExcPending
																							if v151 != 0 {
																								return int64(0)
																							} else {
																								v154 = F_BuildTupleFromCStrings(m, v150, v9+int32(144))
																								mBase = m.M
																								v155 = m.ExcPending
																								if v155 != 0 {
																									return int64(0)
																								} else {
																									v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
																									v157 = F_HeapTupleHeaderGetDatum(m, v156)
																									mBase = m.M
																									v158 = m.ExcPending
																									if v158 != 0 {
																										return int64(0)
																									} else {
																										F_UnlockReleaseBuffer(m, v39)
																										mBase = m.M
																										v160 = m.ExcPending
																										if v160 != 0 {
																											return int64(0)
																										} else {
																											F_relation_close(m, v23, int32(1))
																											mBase = m.M
																											v163 = m.ExcPending
																											if v163 != 0 {
																												return int64(0)
																											} else {
																												m.G0 = v9 + int32(192)
																												return v157
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
										v39 = F_ReadBuffer(m, v23, int32(0))
										mBase = m.M
										v40 = m.ExcPending
										if v40 != 0 {
											return int64(0)
										} else {
											F_LockBufferInternal(m, v39, int32(1))
											mBase = m.M
											v43 = m.ExcPending
											if v43 != 0 {
												return int64(0)
											} else {
												if v39 < int32(0) {
													v47 = *(*int32)(unsafe.Add(mBase, _c_F_bt_metap[0]))
													v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v39^int32(-1))<<(uint(int32(2))%32))))
													v61 = v53
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, _c_F_bt_metap[1]))
													v61 = v55 + v39<<(uint(int32(13))%32) + int32(-8192)
												}
												v65 = F_get_call_result_type(m, l0, int32(0), v9+int32(188))
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int64(0)
												} else {
													if v65 != int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v226 = m.ExcPending
														if v226 != 0 {
															return int64(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_bt_metap_5), int32(0))
															mBase = m.M
															v230 = m.ExcPending
															if v230 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_bt_metap_2), int32(885), int32(_a_F_bt_metap_3))
																mBase = m.M
																v235 = m.ExcPending
																if v235 != 0 {
																	return int64(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v69 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
														v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
														if v70 <= int32(8) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v239 = m.ExcPending
															if v239 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(50724996))
																mBase = m.M
																v242 = m.ExcPending
																if v242 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_bt_metap_6), int32(0))
																	mBase = m.M
																	v246 = m.ExcPending
																	if v246 != 0 {
																		return int64(0)
																	} else {
																		F_errhint(m, int32(_a_F_bt_metap_7), int32(0))
																		mBase = m.M
																		v250 = m.ExcPending
																		if v250 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_bt_metap_2), int32(899), int32(_a_F_bt_metap_3))
																			mBase = m.M
																			v255 = m.ExcPending
																			if v255 != 0 {
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
														} else {
															v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)+24))
															*(*int32)(unsafe.Add(mBase, uint32(v9)+112)) = v73
															v78 = F_psprintf(m, int32(_a_F_bt_metap_8), v9+int32(112))
															mBase = m.M
															v79 = m.ExcPending
															if v79 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v9)+144)) = v78
																v81 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
																*(*int32)(unsafe.Add(mBase, uint32(v9)+96)) = v81
																v86 = F_psprintf(m, int32(_a_F_bt_metap_8), v9+int32(96))
																mBase = m.M
																v87 = m.ExcPending
																if v87 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+148)) = v86
																	v89 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
																	*(*int32)(unsafe.Add(mBase, uint32(v9)+80)) = v89
																	v94 = F_psprintf(m, int32(_a_F_bt_metap_9), v9+int32(80))
																	mBase = m.M
																	v95 = m.ExcPending
																	if v95 != 0 {
																		return int64(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+152)) = v94
																		v97 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
																		*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v97
																		v102 = F_psprintf(m, int32(_a_F_bt_metap_9), v9-int32(-64))
																		mBase = m.M
																		v103 = m.ExcPending
																		if v103 != 0 {
																			return int64(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+156)) = v102
																			v105 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
																			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v105
																			v110 = F_psprintf(m, int32(_a_F_bt_metap_9), v9+int32(48))
																			mBase = m.M
																			v111 = m.ExcPending
																			if v111 != 0 {
																				return int64(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+160)) = v110
																				v113 = *(*int32)(unsafe.Add(mBase, uint32(v61)+44))
																				*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v113
																				v118 = F_psprintf(m, int32(_a_F_bt_metap_9), v9+int32(32))
																				mBase = m.M
																				v119 = m.ExcPending
																				if v119 != 0 {
																					return int64(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v9)+164)) = v118
																					v121 = *(*int32)(unsafe.Add(mBase, uint32(v61)+28))
																					if base.Ui32(int32(3)) <= base.Ui32(v121) {
																						v124 = *(*int32)(unsafe.Add(mBase, uint32(v61)+48))
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v124
																						v129 = F_psprintf(m, int32(_a_F_bt_metap_9), v9+int32(16))
																						mBase = m.M
																						v130 = m.ExcPending
																						if v130 != 0 {
																							return int64(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v9)+168)) = v129
																							v132 = *(*float64)(unsafe.Add(mBase, uint32(v61)+56))
																							*(*float64)(unsafe.Add(mBase, uint32(v9))) = v132
																							v135 = F_psprintf(m, int32(_a_F_bt_metap_10), v9)
																							mBase = m.M
																							v136 = m.ExcPending
																							if v136 != 0 {
																								return int64(0)
																							} else {
																								v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+64)))
																								if v139 != 0 {
																									v140 = int32(_a_F_bt_metap_11)
																								} else {
																									v140 = int32(_a_F_bt_metap_12)
																								}
																								v145 = v135
																								v146 = v140
																								*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v146
																								*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v145
																								v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
																								v150 = F_TupleDescGetAttInMetadata(m, v149)
																								mBase = m.M
																								v151 = m.ExcPending
																								if v151 != 0 {
																									return int64(0)
																								} else {
																									v154 = F_BuildTupleFromCStrings(m, v150, v9+int32(144))
																									mBase = m.M
																									v155 = m.ExcPending
																									if v155 != 0 {
																										return int64(0)
																									} else {
																										v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
																										v157 = F_HeapTupleHeaderGetDatum(m, v156)
																										mBase = m.M
																										v158 = m.ExcPending
																										if v158 != 0 {
																											return int64(0)
																										} else {
																											F_UnlockReleaseBuffer(m, v39)
																											mBase = m.M
																											v160 = m.ExcPending
																											if v160 != 0 {
																												return int64(0)
																											} else {
																												F_relation_close(m, v23, int32(1))
																												mBase = m.M
																												v163 = m.ExcPending
																												if v163 != 0 {
																													return int64(0)
																												} else {
																													m.G0 = v9 + int32(192)
																													return v157
																												}
																											}
																										}
																									}
																								}
																							}
																						}
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+168)) = int32(_a_F_bt_metap_13)
																						v145 = int32(_a_F_bt_metap_14)
																						v146 = int32(_a_F_bt_metap_12)
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+176)) = v146
																						*(*int32)(unsafe.Add(mBase, uint32(v9)+172)) = v145
																						v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+188))
																						v150 = F_TupleDescGetAttInMetadata(m, v149)
																						mBase = m.M
																						v151 = m.ExcPending
																						if v151 != 0 {
																							return int64(0)
																						} else {
																							v154 = F_BuildTupleFromCStrings(m, v150, v9+int32(144))
																							mBase = m.M
																							v155 = m.ExcPending
																							if v155 != 0 {
																								return int64(0)
																							} else {
																								v156 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
																								v157 = F_HeapTupleHeaderGetDatum(m, v156)
																								mBase = m.M
																								v158 = m.ExcPending
																								if v158 != 0 {
																									return int64(0)
																								} else {
																									F_UnlockReleaseBuffer(m, v39)
																									mBase = m.M
																									v160 = m.ExcPending
																									if v160 != 0 {
																										return int64(0)
																									} else {
																										F_relation_close(m, v23, int32(1))
																										mBase = m.M
																										v163 = m.ExcPending
																										if v163 != 0 {
																											return int64(0)
																										} else {
																											m.G0 = v9 + int32(192)
																											return v157
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
				v171 = m.ExcPending
				if v171 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(16797828))
					mBase = m.M
					v174 = m.ExcPending
					if v174 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_bt_metap_15), int32(0))
						mBase = m.M
						v178 = m.ExcPending
						if v178 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_bt_metap_2), int32(856), int32(_a_F_bt_metap_3))
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
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
		}
	}
}
func F_bt_page_print_tuples(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v307 int32
	_ = v307
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int64
	_ = v394
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	v15 = m.G0
	v17 = v15 - int32(144)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v25 = v19 + v20<<(uint(int32(2))%32) + int32(20)
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L17
	} else {
		goto L72
	}
L2:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v29 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+56)) = uint8(v29)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+48)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = base.I64_extend16_s(base.I64_extend_i32_u(v20))
	v38 = v19 + v28&int32(_a_F_bt_page_print_tuples_0)
	v39 = base.I64_extend_i32_u(v38)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+72)) = v39
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+88)) = base.I64_extend_i32_u(int32(base.Ui32(v41) >> (uint(int32(15)) % 32)))
	v47 = v41 & int32(_a_F_bt_page_print_tuples_1)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+80)) = base.I64_extend_i32_u(v47)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+96)) = base.I64_extend_i32_u(int32(base.Ui32(v41)>>(uint(int32(14))%32)) & int32(1))
	if v29 <= base.I32_extend16_s(v41) {
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
	v403 = m.ExcPending
	if v403 != 0 {
		goto L17
	} else {
		goto L69
	}
L5:
	;
	if base.Ui32(int32(_a_F_bt_page_print_tuples_2)) <= base.Ui32(v82) {
		goto L1
	} else {
		goto L16
	}
L6:
	;
	v61 = int32(8)
	goto L8
L7:
	;
	v61 = int32(16)
	goto L8
L8:
	;
	v62 = v47 - v61
	if v41&int32(_a_F_bt_page_print_tuples_2) == int32(0) {
		v82 = v62
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	if v67&int32(_a_F_bt_page_print_tuples_2) != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+2)))
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v82 = v70 | v71<<(uint(int32(16))%32) - v61
	goto L5
L11:
	;
	goto L12
L12:
	;
	if v67&int32(_a_F_bt_page_print_tuples_3) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v80 = v62 - int32(8)
	goto L15
L14:
	;
	v80 = v62
	goto L15
L15:
	;
	v82 = v80
	goto L5
L16:
	;
	v85 = int32(1)
	v90 = F_palloc0(m, v82*int32(3)+v85)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int64(0)
L18:
	;
	if v82 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v149 = F_cstring_to_text(m, v90)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L17
	} else {
		goto L27
	}
L20:
	;
	v96 = v38 + v61
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v97
	v102 = F_pg_sprintf(m, v90, int32(_a_F_bt_page_print_tuples_4), v17+int32(32))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	if v82 == int32(1) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v107 = v90
	v110 = v85
	goto L23
L23:
	;
	v120 = int32(32)
	*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)) = uint8(v120)
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110+v96))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v123
	v126 = v107 + int32(3)
	v130 = F_pg_sprintf(m, v126, int32(_a_F_bt_page_print_tuples_4), v17+int32(16))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L17
	} else {
		goto L25
	}
L24:
	;
	goto L19
L25:
	;
	v133 = v110 + int32(1)
	if v133 != v82 {
		v107 = v126
		v110 = v133
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+104)) = base.I64_extend_i32_u(v149)
	F_pfree(m, v90)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L17
	} else {
		goto L28
	}
L28:
	;
	v155 = int32(1)
	v158 = v27 & (base.B2i32(v20 != v155) | v26)
	if v158&v155 != 0 {
		goto L33
	} else {
		goto L34
	}
L29:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v391 = F_heap_form_tuple(m, v386, v17-int32(-64), v17+int32(48))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L17
	} else {
		goto L67
	}
L30:
	;
	v370 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+56)) = uint8(v370)
	goto L29
L31:
	;
	v365 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+55)) = uint8(v365)
	goto L30
L32:
	;
	v180 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	if v180&int32(_a_F_bt_page_print_tuples_2) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L33:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v162 = int32(_a_F_bt_page_print_tuples_5)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+112)) = base.I64_extend_i32_u(base.B2i32(v161&v162 == v162))
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)))
	if v168&int32(_a_F_bt_page_print_tuples_2) != 0 {
		v179 = v168
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v172 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+54)) = uint8(v172)
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)))
	if v174&int32(_a_F_bt_page_print_tuples_2) == int32(0) {
		goto L31
	} else {
		goto L37
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = v39
	goto L30
L37:
	;
	v179 = v174
	goto L32
L38:
	;
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	if v216&int32(_a_F_bt_page_print_tuples_2) == int32(0) {
		goto L30
	} else {
		goto L51
	}
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v17)+120)) = base.I64_extend_i32_u(v197 + v38 + v196)
	goto L38
L40:
	;
	v208 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+55)) = uint8(v208)
	goto L38
L41:
	;
	v198 = int32(1)
	if v180&int32(_a_F_bt_page_print_tuples_2) != 0 {
		goto L46
	} else {
		goto L47
	}
L42:
	;
	if v180&int32(_a_F_bt_page_print_tuples_3) == int32(0) {
		goto L40
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+2)))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v196 = v192
	v197 = v193 << (uint(int32(16)) % 32)
	goto L41
L45:
	;
	v196 = int32(-6)
	v197 = v179 & int32(_a_F_bt_page_print_tuples_1)
	goto L41
L46:
	;
	v203 = v158 & v198
	goto L48
L47:
	;
	v203 = v198
	goto L48
L48:
	;
	if v203 == int32(0) {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	if v19 != 0 {
		goto L39
	} else {
		goto L50
	}
L50:
	;
	goto L40
L51:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+2)))
	v224 = v216 & int32(4095)
	v227 = F_palloc(m, v224<<(uint(int32(3))%32))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L17
	} else {
		goto L52
	}
L52:
	;
	if v224 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v359 = F_construct_array_builtin(m, v227, v224, int32(27))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L17
	} else {
		goto L65
	}
L54:
	;
	v234 = v38 + v221<<(uint(int32(16))%32) + v222
	v236 = v224 & int32(3)
	v237 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v224) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v248 = v237
	v253 = int32(0)
	goto L58
L56:
	;
	v307 = v237
	goto L57
L57:
	;
	v321 = v307
	v322 = v237
	goto L62
L58:
	;
	v258 = int32(3)
	v261 = int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v227+v248<<(uint(v258)%32)))) = base.I64_extend_i32_u(v234 + v248*v261)
	v267 = v248 | int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v227+v267<<(uint(v258)%32)))) = base.I64_extend_i32_u(v234 + v267*v261)
	v277 = v248 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v227+v277<<(uint(v258)%32)))) = base.I64_extend_i32_u(v234 + v277*v261)
	v287 = v248 | v258
	*(*int64)(unsafe.Add(mBase, uint32(v227+v287<<(uint(v258)%32)))) = base.I64_extend_i32_u(v234 + v287*v261)
	v296 = int32(4)
	v297 = v248 + v296
	v299 = v253 + v296
	if v299 != v224&int32(4092) {
		v248 = v297
		v253 = v299
		goto L58
	} else {
		goto L60
	}
L59:
	;
	if v236 == int32(0) {
		goto L53
	} else {
		goto L61
	}
L60:
	;
	goto L59
L61:
	;
	v307 = v297
	goto L57
L62:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v227+v321<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v234 + v321*int32(6))
	v339 = int32(1)
	v342 = v322 + v339
	if v342 != v236 {
		v321 = v321 + v339
		v322 = v342
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
	*(*int64)(unsafe.Add(mBase, uint32(v17)+128)) = base.I64_extend_i32_u(v359)
	F_pfree(m, v227)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L17
	} else {
		goto L66
	}
L66:
	;
	goto L29
L67:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	v394 = F_HeapTupleHeaderGetDatum(m, v393)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L17
	} else {
		goto L68
	}
L68:
	;
	m.G0 = v17 + int32(144)
	return v394
L69:
	;
	F_errmsg_internal(m, int32(_a_F_bt_page_print_tuples_6), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L17
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_bt_page_print_tuples_7), int32(504), int32(_a_F_bt_page_print_tuples_8))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v82
	F_errmsg_internal(m, int32(_a_F_bt_page_print_tuples_9), v17)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L17
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_bt_page_print_tuples_7), int32(549), int32(_a_F_bt_page_print_tuples_8))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
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
func F_bt_page_stats_1_9(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_bt_page_stats_internal(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
