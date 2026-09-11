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
	var v75 int32
	_ = v75
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
											v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
											v77 = base.I32_extend16_s(v72) + v48
											if v75 < v77 {
												v79 = v75
											} else {
												v79 = v77
											}
											v86 = v79
										} else {
											v81 = v48 - base.I32_extend16_s(v72)
											if v81 < v49 {
												v83 = v49
											} else {
												v83 = v81
											}
											v86 = v83
										}
										v87 = int32(0)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+v86&int32(65535)<<(uint(int32(2))%32))+20))
										v101 = F__bt_tuple_before_array_skeys(m, l0, v52, v88+v94&int32(32767), v18, l4, v87, v87, v87)
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
											v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
											v77 = base.I32_extend16_s(v72) + v48
											if v75 < v77 {
												v79 = v75
											} else {
												v79 = v77
											}
											v86 = v79
										} else {
											v81 = v48 - base.I32_extend16_s(v72)
											if v81 < v49 {
												v83 = v49
											} else {
												v83 = v81
											}
											v86 = v83
										}
										v87 = int32(0)
										v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+v86&int32(65535)<<(uint(int32(2))%32))+20))
										v101 = F__bt_tuple_before_array_skeys(m, l0, v52, v88+v94&int32(32767), v18, l4, v87, v87, v87)
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
												v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
												v77 = base.I32_extend16_s(v72) + v48
												if v75 < v77 {
													v79 = v75
												} else {
													v79 = v77
												}
												v86 = v79
											} else {
												v81 = v48 - base.I32_extend16_s(v72)
												if v81 < v49 {
													v83 = v49
												} else {
													v83 = v81
												}
												v86 = v83
											}
											v87 = int32(0)
											v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
											v94 = *(*int32)(unsafe.Add(mBase, uint32(v88+v86&int32(65535)<<(uint(int32(2))%32))+20))
											v101 = F__bt_tuple_before_array_skeys(m, l0, v52, v88+v94&int32(32767), v18, l4, v87, v87, v87)
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v19 = v17 | v18
	if v19&int32(1) != 0 {
		if v19&int32(262144) != 0 {
			v24 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l4)+19)) = uint8(v24)
			v188 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
			v205 = v188
			m.G0 = v15 + int32(16)
			return v205
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
			v40 = v38 & int32(65535)
			switch v40 - int32(1) {
			case 0:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(base.Ui32(v31) < base.Ui32(v29)))
				v205 = v27
				m.G0 = v15 + int32(16)
				return v205
			case 1:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(base.Ui32(v31) <= base.Ui32(v29)))
				v205 = v27
				m.G0 = v15 + int32(16)
				return v205
			case 2:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v29 == v31))
				v205 = v27
				m.G0 = v15 + int32(16)
				return v205
			case 3:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(base.Ui32(v29) <= base.Ui32(v31)))
				v205 = v27
				m.G0 = v15 + int32(16)
				return v205
			case 4:
				*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(base.Ui32(v29) < base.Ui32(v31)))
				v205 = v27
				m.G0 = v15 + int32(16)
				return v205
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v40
					F_errmsg_internal(m, int32(453148), v15)
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(463238), int32(950), int32(143987))
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
			v205 = int32(0)
			m.G0 = v15 + int32(16)
			return v205
		} else {
			v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if l4 == int32(0) {
				v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
				v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
				v121 = v119 << (uint(int32(2)) % 32)
				v122 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
				v126 = *(*int32)(unsafe.Add(mBase, uint32(v121+v122-int32(4))))
				if v118 != 0 {
					v127 = v118
				} else {
					v127 = v126
				}
				v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				if v128 != 0 {
					v129 = v128
				} else {
					v129 = v126
				}
				if v129 != v126 {
					v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
					v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
					v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v154&int32(16777216) != 0 {
						v157 = int32(6) - v152
					} else {
						v157 = v152
					}
					v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
					mBase = m.M
					v160 = m.ExcPending
					if v160 != 0 {
						return int32(0)
					} else {
						if v159 == int32(0) {
							v188 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
							v205 = v188
							m.G0 = v15 + int32(16)
							return v205
						} else {
							v163 = F_get_opcode(m, v159)
							mBase = m.M
							v164 = m.ExcPending
							if v164 != 0 {
								return int32(0)
							} else {
								if v163 == int32(0) {
									v188 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
									v205 = v188
									m.G0 = v15 + int32(16)
									return v205
								} else {
									v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
									v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
									v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
									mBase = m.M
									v171 = m.ExcPending
									if v171 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
										v205 = int32(1)
										m.G0 = v15 + int32(16)
										return v205
									}
								}
							}
						}
					}
				} else {
					v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if v131 != 0 {
						v132 = v131
					} else {
						v132 = v126
					}
					if v127 != v132 {
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
						v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
						v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v154&int32(16777216) != 0 {
							v157 = int32(6) - v152
						} else {
							v157 = v152
						}
						v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
							return int32(0)
						} else {
							if v159 == int32(0) {
								v188 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
								v205 = v188
								m.G0 = v15 + int32(16)
								return v205
							} else {
								v163 = F_get_opcode(m, v159)
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return int32(0)
								} else {
									if v163 == int32(0) {
										v188 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
										v205 = v188
										m.G0 = v15 + int32(16)
										return v205
									} else {
										v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
										v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
										v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
										v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
										mBase = m.M
										v171 = m.ExcPending
										if v171 != 0 {
											return int32(0)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
											v205 = int32(1)
											m.G0 = v15 + int32(16)
											return v205
										}
									}
								}
							}
						}
					} else {
						v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
						v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
						v139 = F_FunctionCall2Coll(m, l1+int32(16), v136, v137, v138)
						mBase = m.M
						v140 = m.ExcPending
						if v140 != 0 {
							return int32(0)
						} else {
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v139 != int32(0)))
							v205 = int32(1)
							m.G0 = v15 + int32(16)
							return v205
						}
					}
				}
			} else {
				if v18&int32(32) != 0 {
					v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
					if v17&int32(32) == int32(0) {
						if v75&int32(65535) != int32(3) {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
							v121 = v119 << (uint(int32(2)) % 32)
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v121+v122-int32(4))))
							if v118 != 0 {
								v127 = v118
							} else {
								v127 = v126
							}
							v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v128 != 0 {
								v129 = v128
							} else {
								v129 = v126
							}
							if v129 != v126 {
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
								v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v154&int32(16777216) != 0 {
									v157 = int32(6) - v152
								} else {
									v157 = v152
								}
								v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return int32(0)
								} else {
									if v159 == int32(0) {
										v188 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
										v205 = v188
										m.G0 = v15 + int32(16)
										return v205
									} else {
										v163 = F_get_opcode(m, v159)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											if v163 == int32(0) {
												v188 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
												v205 = v188
												m.G0 = v15 + int32(16)
												return v205
											} else {
												v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
												v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
												v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
													v205 = int32(1)
													m.G0 = v15 + int32(16)
													return v205
												}
											}
										}
									}
								}
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v131 != 0 {
									v132 = v131
								} else {
									v132 = v126
								}
								if v127 != v132 {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
									v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
									v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v154&int32(16777216) != 0 {
										v157 = int32(6) - v152
									} else {
										v157 = v152
									}
									v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										if v159 == int32(0) {
											v188 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
											v205 = v188
											m.G0 = v15 + int32(16)
											return v205
										} else {
											v163 = F_get_opcode(m, v159)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												if v163 == int32(0) {
													v188 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
													v205 = v188
													m.G0 = v15 + int32(16)
													return v205
												} else {
													v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
													v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
													v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
													v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
														v205 = int32(1)
														m.G0 = v15 + int32(16)
														return v205
													}
												}
											}
										}
									}
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
									v139 = F_FunctionCall2Coll(m, l1+int32(16), v136, v137, v138)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v139 != int32(0)))
										v205 = int32(1)
										m.G0 = v15 + int32(16)
										return v205
									}
								}
							}
						} else {
							v109 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							if v109 != int32(-1) {
								v112 = F__bt_saoparray_shrink(m, l0, l2, l3, l5, l4, l6)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									v205 = v112
									m.G0 = v15 + int32(16)
									return v205
								}
							} else {
								v114 = F__bt_skiparray_shrink(m, l0, l3, l4, l6)
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									v205 = v114
									m.G0 = v15 + int32(16)
									return v205
								}
							}
						}
					} else {
						v82 = int32(3)
						v83 = base.B2i32(v75&int32(65535) != v82)
						v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
						if base.B2i32(v83 == int32(0))&base.B2i32(v86 == v82) != 0 {
							v188 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
							v205 = v188
							m.G0 = v15 + int32(16)
							return v205
						} else {
							if v83 == int32(0) {
								v109 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
								if v109 != int32(-1) {
									v112 = F__bt_saoparray_shrink(m, l0, l2, l3, l5, l4, l6)
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
										return int32(0)
									} else {
										v205 = v112
										m.G0 = v15 + int32(16)
										return v205
									}
								} else {
									v114 = F__bt_skiparray_shrink(m, l0, l3, l4, l6)
									mBase = m.M
									v115 = m.ExcPending
									if v115 != 0 {
										return int32(0)
									} else {
										v205 = v114
										m.G0 = v15 + int32(16)
										return v205
									}
								}
							} else {
								if v86&int32(65535) != int32(3) {
									v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
									v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
									v121 = v119 << (uint(int32(2)) % 32)
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v121+v122-int32(4))))
									if v118 != 0 {
										v127 = v118
									} else {
										v127 = v126
									}
									v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
									if v128 != 0 {
										v129 = v128
									} else {
										v129 = v126
									}
									if v129 != v126 {
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
										v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
										v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
										v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										if v154&int32(16777216) != 0 {
											v157 = int32(6) - v152
										} else {
											v157 = v152
										}
										v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
										mBase = m.M
										v160 = m.ExcPending
										if v160 != 0 {
											return int32(0)
										} else {
											if v159 == int32(0) {
												v188 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
												v205 = v188
												m.G0 = v15 + int32(16)
												return v205
											} else {
												v163 = F_get_opcode(m, v159)
												mBase = m.M
												v164 = m.ExcPending
												if v164 != 0 {
													return int32(0)
												} else {
													if v163 == int32(0) {
														v188 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
														v205 = v188
														m.G0 = v15 + int32(16)
														return v205
													} else {
														v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
														v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
														v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
														v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
														mBase = m.M
														v171 = m.ExcPending
														if v171 != 0 {
															return int32(0)
														} else {
															*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
															v205 = int32(1)
															m.G0 = v15 + int32(16)
															return v205
														}
													}
												}
											}
										}
									} else {
										v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
										if v131 != 0 {
											v132 = v131
										} else {
											v132 = v126
										}
										if v127 != v132 {
											v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
											v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
											v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
											v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											if v154&int32(16777216) != 0 {
												v157 = int32(6) - v152
											} else {
												v157 = v152
											}
											v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
											mBase = m.M
											v160 = m.ExcPending
											if v160 != 0 {
												return int32(0)
											} else {
												if v159 == int32(0) {
													v188 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
													v205 = v188
													m.G0 = v15 + int32(16)
													return v205
												} else {
													v163 = F_get_opcode(m, v159)
													mBase = m.M
													v164 = m.ExcPending
													if v164 != 0 {
														return int32(0)
													} else {
														if v163 == int32(0) {
															v188 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
															v205 = v188
															m.G0 = v15 + int32(16)
															return v205
														} else {
															v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
															v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
															v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
															v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
															mBase = m.M
															v171 = m.ExcPending
															if v171 != 0 {
																return int32(0)
															} else {
																*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
																v205 = int32(1)
																m.G0 = v15 + int32(16)
																return v205
															}
														}
													}
												}
											}
										} else {
											v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
											v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
											v139 = F_FunctionCall2Coll(m, l1+int32(16), v136, v137, v138)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v139 != int32(0)))
												v205 = int32(1)
												m.G0 = v15 + int32(16)
												return v205
											}
										}
									}
								} else {
									v192 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
									if v192 != int32(-1) {
										v195 = F__bt_saoparray_shrink(m, l0, l3, l2, l5, l4, l6)
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											v199 = v195
											v205 = v199
											m.G0 = v15 + int32(16)
											return v205
										}
									} else {
										v197 = F__bt_skiparray_shrink(m, l0, l2, l4, l6)
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int32(0)
										} else {
											v199 = v197
											v205 = v199
											m.G0 = v15 + int32(16)
											return v205
										}
									}
								}
							}
						}
					}
				} else {
					if v17&int32(32) == int32(0) {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
						v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
						v121 = v119 << (uint(int32(2)) % 32)
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
						v126 = *(*int32)(unsafe.Add(mBase, uint32(v121+v122-int32(4))))
						if v118 != 0 {
							v127 = v118
						} else {
							v127 = v126
						}
						v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
						if v128 != 0 {
							v129 = v128
						} else {
							v129 = v126
						}
						if v129 != v126 {
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
							v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
							v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
							v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if v154&int32(16777216) != 0 {
								v157 = int32(6) - v152
							} else {
								v157 = v152
							}
							v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
							mBase = m.M
							v160 = m.ExcPending
							if v160 != 0 {
								return int32(0)
							} else {
								if v159 == int32(0) {
									v188 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
									v205 = v188
									m.G0 = v15 + int32(16)
									return v205
								} else {
									v163 = F_get_opcode(m, v159)
									mBase = m.M
									v164 = m.ExcPending
									if v164 != 0 {
										return int32(0)
									} else {
										if v163 == int32(0) {
											v188 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
											v205 = v188
											m.G0 = v15 + int32(16)
											return v205
										} else {
											v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
											v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
											v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
											v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
											mBase = m.M
											v171 = m.ExcPending
											if v171 != 0 {
												return int32(0)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
												v205 = int32(1)
												m.G0 = v15 + int32(16)
												return v205
											}
										}
									}
								}
							}
						} else {
							v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
							if v131 != 0 {
								v132 = v131
							} else {
								v132 = v126
							}
							if v127 != v132 {
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
								v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v154&int32(16777216) != 0 {
									v157 = int32(6) - v152
								} else {
									v157 = v152
								}
								v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return int32(0)
								} else {
									if v159 == int32(0) {
										v188 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
										v205 = v188
										m.G0 = v15 + int32(16)
										return v205
									} else {
										v163 = F_get_opcode(m, v159)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											if v163 == int32(0) {
												v188 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
												v205 = v188
												m.G0 = v15 + int32(16)
												return v205
											} else {
												v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
												v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
												v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
													v205 = int32(1)
													m.G0 = v15 + int32(16)
													return v205
												}
											}
										}
									}
								}
							} else {
								v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
								v139 = F_FunctionCall2Coll(m, l1+int32(16), v136, v137, v138)
								mBase = m.M
								v140 = m.ExcPending
								if v140 != 0 {
									return int32(0)
								} else {
									*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v139 != int32(0)))
									v205 = int32(1)
									m.G0 = v15 + int32(16)
									return v205
								}
							}
						}
					} else {
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+6)))
						if v100 == int32(3) {
							v192 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							if v192 != int32(-1) {
								v195 = F__bt_saoparray_shrink(m, l0, l3, l2, l5, l4, l6)
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return int32(0)
								} else {
									v199 = v195
									v205 = v199
									m.G0 = v15 + int32(16)
									return v205
								}
							} else {
								v197 = F__bt_skiparray_shrink(m, l0, l2, l4, l6)
								mBase = m.M
								v198 = m.ExcPending
								if v198 != 0 {
									return int32(0)
								} else {
									v199 = v197
									v205 = v199
									m.G0 = v15 + int32(16)
									return v205
								}
							}
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
							v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+4)))
							v121 = v119 << (uint(int32(2)) % 32)
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v70)+212))
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v121+v122-int32(4))))
							if v118 != 0 {
								v127 = v118
							} else {
								v127 = v126
							}
							v128 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							if v128 != 0 {
								v129 = v128
							} else {
								v129 = v126
							}
							if v129 != v126 {
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
								v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
								v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								if v154&int32(16777216) != 0 {
									v157 = int32(6) - v152
								} else {
									v157 = v152
								}
								v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
								mBase = m.M
								v160 = m.ExcPending
								if v160 != 0 {
									return int32(0)
								} else {
									if v159 == int32(0) {
										v188 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
										v205 = v188
										m.G0 = v15 + int32(16)
										return v205
									} else {
										v163 = F_get_opcode(m, v159)
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											if v163 == int32(0) {
												v188 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
												v205 = v188
												m.G0 = v15 + int32(16)
												return v205
											} else {
												v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
												v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
												v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
												v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
												mBase = m.M
												v171 = m.ExcPending
												if v171 != 0 {
													return int32(0)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
													v205 = int32(1)
													m.G0 = v15 + int32(16)
													return v205
												}
											}
										}
									}
								}
							} else {
								v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
								if v131 != 0 {
									v132 = v131
								} else {
									v132 = v126
								}
								if v127 != v132 {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v70)+208))
									v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v121-int32(4))))
									v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
									v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									if v154&int32(16777216) != 0 {
										v157 = int32(6) - v152
									} else {
										v157 = v152
									}
									v159 = F_get_opfamily_member(m, v150, v129, v127, base.I32_extend16_s(v157))
									mBase = m.M
									v160 = m.ExcPending
									if v160 != 0 {
										return int32(0)
									} else {
										if v159 == int32(0) {
											v188 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
											v205 = v188
											m.G0 = v15 + int32(16)
											return v205
										} else {
											v163 = F_get_opcode(m, v159)
											mBase = m.M
											v164 = m.ExcPending
											if v164 != 0 {
												return int32(0)
											} else {
												if v163 == int32(0) {
													v188 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(v188)
													v205 = v188
													m.G0 = v15 + int32(16)
													return v205
												} else {
													v167 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
													v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
													v169 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
													v170 = F_OidFunctionCall2Coll(m, v163, v167, v168, v169)
													mBase = m.M
													v171 = m.ExcPending
													if v171 != 0 {
														return int32(0)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v170 != int32(0)))
														v205 = int32(1)
														m.G0 = v15 + int32(16)
														return v205
													}
												}
											}
										}
									}
								} else {
									v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
									v138 = *(*int32)(unsafe.Add(mBase, uint32(l3)+44))
									v139 = F_FunctionCall2Coll(m, l1+int32(16), v136, v137, v138)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										*(*uint8)(unsafe.Add(mBase, uint32(l6))) = uint8(base.B2i32(v139 != int32(0)))
										v205 = int32(1)
										m.G0 = v15 + int32(16)
										return v205
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
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v10 = int32(1)
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)))
	if base.Ui32(v11) < base.Ui32(int32(25)) {
		v20 = v10
	} else {
		v20 = int32(base.Ui32(v11+int32(262120))>>(uint(int32(2))%32)) + v10
	}
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v22 == int32(1) {
		v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)))
		v31 = F_PageAddItemExtended(m, l0, v21, v25&int32(8191), v20&int32(65535), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			if v31 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = int64(0)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(383888), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errfinish(m, int32(465338), int32(574), int32(315782))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
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
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
		v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+6)))
		if v48&int32(8192) == int32(0) {
			v65 = v48 & int32(8191)
		} else {
			v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+5)))
			if v53&int32(32) == int32(0) {
				v65 = v48 & int32(8191)
			} else {
				v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+2)))
				v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21))))
				v65 = v58 | v59<<(uint(int32(16))%32)
			}
		}
		v67 = v46 * int32(6)
		if int32(1) < v46 {
			v75 = (v65 + v67 + int32(7)) & int32(-8)
		} else {
			v75 = v65
		}
		v76 = F_palloc0(m, v75)
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return
		} else {
			if v65 != 0 {
				v78 = F__emscripten_memcpy_bulkmem(m, v76, v21, v65)
				mBase = m.M
				v79 = v78
			} else {
				v79 = v76
			}
			v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+6)))
			v83 = v80&int32(-8192) | v75
			if int32(2) <= v46 {
				*(*uint16)(unsafe.Add(mBase, uint32(v79)+2)) = uint16(v65)
				v87 = int32(8192)
				v88 = v46 | v87
				*(*uint16)(unsafe.Add(mBase, uint32(v79)+4)) = uint16(v88)
				v91 = v83 | v87
				*(*uint16)(unsafe.Add(mBase, uint32(v79)+6)) = uint16(v91)
				v94 = int32(base.Ui32(v65) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v79))) = uint16(v94)
				if v67 != 0 {
					v97 = F__emscripten_memcpy_bulkmem(m, v79+v65, v47, v67)
					mBase = m.M
				} else {
				}
				v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+6)))
				v107 = v99
			} else {
				v101 = v83 & int32(57343)
				*(*uint16)(unsafe.Add(mBase, uint32(v79)+6)) = uint16(v101)
				v103 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
				*(*int32)(unsafe.Add(mBase, uint32(v79))) = v103
				v105 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v79)+4)) = uint16(v105)
				v107 = v101
			}
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v108<<(uint(int32(2))%32))+46)) = uint16(v112)
			v119 = F_PageAddItemExtended(m, l0, v79, v107&int32(8191), v20&int32(65535), int32(0))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return
			} else {
				if v119 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return
					} else {
						F_errmsg_internal(m, int32(383888), int32(0))
						mBase = m.M
						v147 = m.ExcPending
						if v147 != 0 {
							return
						} else {
							F_errfinish(m, int32(465338), int32(594), int32(315782))
							mBase = m.M
							v152 = m.ExcPending
							if v152 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					F_pfree(m, v79)
					mBase = m.M
					v124 = m.ExcPending
					if v124 != 0 {
						return
					} else {
						v125 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v125 + int32(1)
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
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	v7 = int32(1)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v8&int32(32) == int32(0) {
		v26 = v7
		v28 = l1
	} else {
		v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v13&int32(8192) == int32(0) {
			v26 = v7
			v28 = l1
		} else {
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
			v26 = v13 & int32(4095)
			v28 = l1 + (v20 | v21<<(uint(int32(16))%32))
		}
	}
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v40 = base.B2i32(base.Ui32((v30+(v31+v26)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v29))
	if v40 == int32(0) {
		if v31 <= int32(50) {
		} else {
			v72 = int32(4)
			v73 = int32(1)
			v74 = l0 + v72
			v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
			*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75 + v73
		}
	} else {
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v47 + int32(1)
		v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v52 = int32(6)
		v56 = v26 * v52
		if v56 != 0 {
			v57 = F__emscripten_memcpy_bulkmem(m, v51+v31*v52, v28, v56)
			mBase = m.M
		} else {
		}
		v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v59 + v26
		v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
		v72 = int32(36)
		v73 = (v63&int32(8191)+int32(7))&int32(16376) | int32(4)
		v74 = l0 + v72
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
		*(*int32)(unsafe.Add(mBase, uint32(v74))) = v75 + v73
	}
	return v40
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
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	if l2 < v5 {
		v21 = *(*int32)(unsafe.Add(mBase, _consts[5]))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(l2^int32(-1))<<(uint(int32(2))%32))))
		v35 = v27
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
			v45 = *(*int32)(unsafe.Add(mBase, _consts[5]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v45+(v40^int32(-1))<<(uint(int32(2))%32))))
			v59 = v51
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
					v70 = *(*int32)(unsafe.Add(mBase, _consts[5]))
					v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v65^int32(-1))<<(uint(int32(2))%32))))
					v84 = v76
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, _consts[6]))
					v84 = v78 + v65<<(uint(int32(13))%32) + int32(-8192)
				}
				v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+32))
				if l2 < int32(0) {
					v89 = *(*int32)(unsafe.Add(mBase, _consts[9]))
					v95 = *(*int32)(unsafe.Add(mBase, uint32(v89+(l2^int32(-1))<<(uint(int32(6))%32))+16))
					v104 = v95
				} else {
					v97 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
					v111 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
					if v111 == int32(0) {
						v115 = *(*int32)(unsafe.Add(mBase, uint32(v59+v60)+4))
						v118 = base.B2i32(v115 == int32(0))
					} else {
						v118 = v5
					}
					v121 = F_errstart(m, int32(14), int32(0))
					mBase = m.M
					v122 = m.ExcPending
					if v122 != 0 {
						return
					} else {
						if v121 != 0 {
							if l2 < int32(0) {
								v126 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v132 = *(*int32)(unsafe.Add(mBase, uint32(v126+(l2^int32(-1))<<(uint(int32(6))%32))+16))
								v141 = v132
							} else {
								v134 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v134+l2<<(uint(int32(6))%32)+int32(-64))+16))
								v141 = v140
							}
							if v40 < int32(0) {
								v145 = *(*int32)(unsafe.Add(mBase, _consts[9]))
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v145+(v40^int32(-1))<<(uint(int32(6))%32))+16))
								v160 = v151
							} else {
								v153 = *(*int32)(unsafe.Add(mBase, _consts[10]))
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v153+v40<<(uint(int32(6))%32)+int32(-64))+16))
								v160 = v159
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v160
							*(*int32)(unsafe.Add(mBase, uint32(v16))) = v141
							F_errmsg_internal(m, int32(36053), v16)
							mBase = m.M
							v165 = m.ExcPending
							if v165 != 0 {
								return
							} else {
								F_errfinish(m, int32(462691), int32(2282), int32(94226))
								mBase = m.M
								v170 = m.ExcPending
								if v170 != 0 {
									return
								} else {
									F__bt_insert_parent(m, l0, l1, l2, v40, l3, v108, v118)
									mBase = m.M
									v173 = m.ExcPending
									if v173 != 0 {
										return
									} else {
										m.G0 = v16 + int32(16)
										return
									}
								}
							}
						} else {
							F__bt_insert_parent(m, l0, l1, l2, v40, l3, v108, v118)
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
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
			v111 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
			if v111 == int32(0) {
				v115 = *(*int32)(unsafe.Add(mBase, uint32(v59+v60)+4))
				v118 = base.B2i32(v115 == int32(0))
			} else {
				v118 = v5
			}
			v121 = F_errstart(m, int32(14), int32(0))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return
			} else {
				if v121 != 0 {
					if l2 < int32(0) {
						v126 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v126+(l2^int32(-1))<<(uint(int32(6))%32))+16))
						v141 = v132
					} else {
						v134 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v134+l2<<(uint(int32(6))%32)+int32(-64))+16))
						v141 = v140
					}
					if v40 < int32(0) {
						v145 = *(*int32)(unsafe.Add(mBase, _consts[9]))
						v151 = *(*int32)(unsafe.Add(mBase, uint32(v145+(v40^int32(-1))<<(uint(int32(6))%32))+16))
						v160 = v151
					} else {
						v153 = *(*int32)(unsafe.Add(mBase, _consts[10]))
						v159 = *(*int32)(unsafe.Add(mBase, uint32(v153+v40<<(uint(int32(6))%32)+int32(-64))+16))
						v160 = v159
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v160
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = v141
					F_errmsg_internal(m, int32(36053), v16)
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return
					} else {
						F_errfinish(m, int32(462691), int32(2282), int32(94226))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return
						} else {
							F__bt_insert_parent(m, l0, l1, l2, v40, l3, v108, v118)
							mBase = m.M
							v173 = m.ExcPending
							if v173 != 0 {
								return
							} else {
								m.G0 = v16 + int32(16)
								return
							}
						}
					}
				} else {
					F__bt_insert_parent(m, l0, l1, l2, v40, l3, v108, v118)
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v7 = int32(1)
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+v6<<(uint(v7)%32)-int32(2)))))
	v14 = v12 << (uint(int32(24)) % 32)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v15&v7 != 0 {
		v18 = v15 | v14
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18
		if v15&int32(64) != 0 {
			v87 = int32(3)
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v87)
			return int32(1)
		} else {
			if v15&int32(128) != 0 {
				if v18&int32(33554432) != 0 {
					v86 = int32(5)
				} else {
					v86 = int32(1)
				}
				v87 = v86
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v87)
				return int32(1)
			} else {
				return int32(0)
			}
		}
	} else {
		if v12&int32(1) == int32(0) {
		} else {
			if v15&int32(16777216) != 0 {
			} else {
				v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
				v35 = int32(6) - v34
				*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v35)
			}
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v15 | v14
		if v15&int32(4) == int32(0) {
			return int32(1)
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
			if v44&int32(1) != 0 {
				return int32(0)
			} else {
				v51 = v43
				for {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
					v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+4)))
					v58 = int32(1)
					v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1-int32(2)+v57<<(uint(v58)%32)))))
					if v61&v58 == int32(0) {
					} else {
						if v56&int32(16777216) != 0 {
						} else {
							v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51)+6)))
							v72 = int32(6) - v71
							*(*uint16)(unsafe.Add(mBase, uint32(v51)+6)) = uint16(v72)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v51))) = v56 | v61<<(uint(int32(24))%32)
					if v56&int32(16) == int32(0) {
						v51 = v51 + int32(48)
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
	var v185 int32
	_ = v185
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v248 int32
	_ = v248
	var v251 int64
	_ = v251
	var v252 int32
	_ = v252
	var v253 int64
	_ = v253
	var v258 int64
	_ = v258
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int64
	_ = v288
	var v290 int64
	_ = v290
	var v292 int64
	_ = v292
	var v294 int64
	_ = v294
	var v296 int64
	_ = v296
	var v298 int64
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	goto L3
L1:
	;
	m.G0 = v12 - int32(-64)
	return v387
L2:
	;
	v280 = v87 + int32(16)
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v87)+20))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v285 = F_MemoryContextAlloc(m, v283, int32(48))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
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
	v40 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(v28^int32(-1))<<(uint(int32(2))%32))))
	v54 = v46
	goto L13
L15:
	;
	goto L16
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
		v387 = v28
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
	v387 = int32(0)
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
	v112 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112+(v107^int32(-1))<<(uint(int32(6))%32))+16))
	v127 = v118
	goto L44
L46:
	;
	goto L47
L47:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[10]))
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
	v131 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v107^int32(-1))<<(uint(int32(2))%32))))
	v145 = v137
	goto L48
L50:
	;
	goto L51
L51:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v145 = v139 + v107<<(uint(int32(13))%32) + int32(-8192)
	goto L48
L52:
	;
	v172 = int32(4419940)
	v174 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v174 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if base.Ui32(v178) <= base.Ui32(int32(2)) {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157+(v80^int32(-1))<<(uint(int32(2))%32))))
	v171 = v163
	goto L52
L54:
	;
	goto L55
L55:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v171 = v165 + v80<<(uint(int32(13))%32) + int32(-8192)
	goto L52
L56:
	;
	v181 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v171)+12)) = uint16(v181)
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v171-int32(-64)))) = uint8(v185)
	*(*int64)(unsafe.Add(mBase, uint32(v171)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v171)+48)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v171)+28)) = int32(3)
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
	v202 = m.ExcPending
	if v202 != 0 {
		goto L8
	} else {
		goto L59
	}
L59:
	;
	F_MarkBufferDirty(m, v80)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L60
	}
L60:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+118)))
	if v206 != int32(112) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v262 = int32(4419940)
	v264 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v264 - int32(1)
	F_LockBuffer(m, v107, int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L8
	} else {
		goto L74
	}
L62:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v210 <= int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v213 != 0 {
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
	v216 = m.ExcPending
	if v216 != 0 {
		goto L8
	} else {
		goto L68
	}
L66:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v214 != 0 {
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
	v220 = m.ExcPending
	if v220 != 0 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	F_XLogRegisterBuffer(m, int32(2), v80, int32(14))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L8
	} else {
		goto L70
	}
L70:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+44)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v127
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v225
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+52)) = uint8(v233)
	F_XLogRegisterBufData(m, int32(2), v10+int32(-36), int32(28))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
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
	v248 = m.ExcPending
	if v248 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	v251 = F_XLogInsert(m, int32(11), int32(160))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	v253 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v145))) = base.I64_rotr(v251, v253)
	*(*uint32)(unsafe.Add(mBase, uint32(v171)+4)) = uint32(v251)
	v258 = int64(base.Ui64(v251) >> (uint(v253) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v171))) = uint32(v258)
	goto L61
L74:
	;
	F_LockBuffer(m, v107, int32(1))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	F_LockBuffer(m, v80, int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	F_ReleaseBuffer(m, v80)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	v387 = v107
	goto L1
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v285
	v288 = *(*int64)(unsafe.Add(mBase, uint32(v87)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v285)+40)) = v288
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v87)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v285)+32)) = v290
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v285)+24)) = v292
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v280)))
	*(*int64)(unsafe.Add(mBase, uint32(v285)+16)) = v294
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v87)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v285)+8)) = v296
	v298 = *(*int64)(unsafe.Add(mBase, uint32(v87)))
	*(*int64)(unsafe.Add(mBase, uint32(v285))) = v298
	v302 = v80
	v306 = v281
	goto L80
L79:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	if v363 == v282 {
		v387 = v312
		goto L1
	} else {
		goto L98
	}
L80:
	;
	if v302 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L8
	} else {
		goto L95
	}
L82:
	;
	F_LockBuffer(m, v302, int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L8
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	v312 = F_ReleaseAndReadBuffer(m, v302, l0, v306)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L8
	} else {
		goto L86
	}
L85:
	;
	goto L84
L86:
	;
	F_LockBuffer(m, v312, int32(1))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L8
	} else {
		goto L87
	}
L87:
	;
	F__bt_checkpage(m, l0, v312)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	if v312 < int32(0) {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	v337 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v336)+16)))
	v338 = v337 + v336
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+12)))
	if v339&int32(20) == int32(0) {
		goto L79
	} else {
		goto L93
	}
L90:
	;
	v322 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v322+(v312^int32(-1))<<(uint(int32(2))%32))))
	v336 = v328
	goto L89
L91:
	;
	goto L92
L92:
	;
	v330 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v336 = v330 + v312<<(uint(int32(13))%32) + int32(-8192)
	goto L89
L93:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v344 != 0 {
		v302 = v312
		v306 = v344
		goto L80
	} else {
		goto L94
	}
L94:
	;
	goto L81
L95:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v349 + int32(4)
	F_errmsg_internal(m, int32(644135), v10+int32(-48))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L8
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(468658), int32(548), int32(79000))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
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
	v368 = m.ExcPending
	if v368 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v338)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v306
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v369 + int32(4)
	F_errmsg_internal(m, int32(52192), v12)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(468658), int32(555), int32(79000))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
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
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v4 = l3
	if l0&int32(3) != 0 {
	} else {
	}
	v31 = F___memset(m, l0, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+10)) = int32(1572864)
	v37 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)) = uint16(v37)
	v43 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)) = uint16(v43)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+14)) = uint16(v43)
	*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint8(v4)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(17180209506)
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+16)))
	v61 = int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0+v59)+12)) = uint16(v61)
	v63 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v63)
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
						v21 = *(*int32)(unsafe.Add(mBase, _consts[5]))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v9^int32(-1))<<(uint(int32(2))%32))))
						v35 = v27
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	v15 = (v11 - int32(1)) & int32(65535)
	if base.Ui32(v15) < base.Ui32(int32(5)) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15<<(uint(int32(2))%32))+uint32(_consts[45])))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18 | v23
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
	v70 = m.ExcPending
	if v70 != 0 {
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
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)))
	if v30 != v32 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v35 = v31
	v36 = v30
	goto L7
L7:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+6)))
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v40 != v41 {
		goto L4
	} else {
		goto L9
	}
L8:
	;
	goto L4
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v43 | v23
	if v43&int32(16) != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35+int32(52)))))
	v54 = v36 + int32(1)
	if v52 == v54&int32(65535) {
		v35 = v35 + int32(48)
		v36 = v54
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
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v71
	F_errmsg_internal(m, int32(453148), v9)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(463238), int32(797), int32(423476))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
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
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int64
	_ = v116
	var v120 int64
	_ = v120
	var v124 int64
	_ = v124
	var v128 int64
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
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
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	v3 = int32(0)
	v15 = F_shm_toc_lookup(m, l1, int64(-6917529027641081852), int32(1))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[46])) = v15
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
			v29 = *(*int32)(unsafe.Add(mBase, _consts[31]))
			if v29 == int32(0) {
			} else {
				v33 = int32(*(*uint8)(unsafe.Add(mBase, _consts[34])))
				if v33 != int32(1) {
				} else {
					v38 = *(*int64)(unsafe.Add(mBase, uint32(v29)+392))
					if int32(1)&base.B2i32(v38 != int64(0)) != 0 {
					} else {
						v42 = int32(4419940)
						v44 = *(*int32)(unsafe.Add(mBase, _consts[7]))
						v45 = int32(1)
						*(*int32)(unsafe.Add(mBase, _consts[7])) = v44 + v45
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = v48 + v45
						*(*int64)(unsafe.Add(mBase, uint32(v29)+392)) = v25
						*(*int32)(unsafe.Add(mBase, uint32(v29))) = v48 + int32(2)
						v59 = *(*int32)(unsafe.Add(mBase, _consts[7]))
						*(*int32)(unsafe.Add(mBase, _consts[7])) = v59 - v45
					}
				}
			}
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
			if v24 != 0 {
				v66 = int32(4)
			} else {
				v66 = int32(5)
			}
			v67 = F_table_open(m, v63, v66)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return
			} else {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				if v24 != 0 {
					v72 = int32(3)
				} else {
					v72 = int32(8)
				}
				v73 = F_index_open(m, v69, v72)
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					v76 = F_palloc0(m, int32(16))
					mBase = m.M
					v77 = m.ExcPending
					if v77 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = v67
						v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
						*(*uint8)(unsafe.Add(mBase, uint32(v76)+12)) = uint8(v80)
						v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+9)))
						*(*uint8)(unsafe.Add(mBase, uint32(v76)+13)) = uint8(v82)
						v86 = F_shm_toc_lookup(m, l1, int64(-6917529027641081854), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return
						} else {
							F_tuplesort_attach_shared(m, v86, l0)
							mBase = m.M
							v89 = m.ExcPending
							if v89 != 0 {
								return
							} else {
								v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+8)))
								if v90 == int32(1) {
									v94 = F_palloc0(m, int32(16))
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return
									} else {
										v96 = *(*int32)(unsafe.Add(mBase, uint32(v76)+4))
										*(*int32)(unsafe.Add(mBase, uint32(v94)+4)) = v96
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
										v99 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v94)+12)) = uint8(v99)
										*(*int32)(unsafe.Add(mBase, uint32(v94)+8)) = v98
										v104 = F_shm_toc_lookup(m, l1, int64(-6917529027641081853), v99)
										mBase = m.M
										v105 = m.ExcPending
										if v105 != 0 {
											return
										} else {
											F_tuplesort_attach_shared(m, v104, l0)
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return
											} else {
												v108 = v104
												v109 = v94
												v113 = F___memcpy(m, int32(4323704), int32(4323544), int32(128))
												mBase = m.M
												v116 = *(*int64)(unsafe.Add(mBase, _consts[47]))
												*(*int64)(unsafe.Add(mBase, _consts[48])) = v116
												v120 = *(*int64)(unsafe.Add(mBase, _consts[49]))
												*(*int64)(unsafe.Add(mBase, _consts[50])) = v120
												v124 = *(*int64)(unsafe.Add(mBase, _consts[51]))
												*(*int64)(unsafe.Add(mBase, _consts[52])) = v124
												v128 = *(*int64)(unsafe.Add(mBase, _consts[53]))
												*(*int64)(unsafe.Add(mBase, _consts[54])) = v128
												v131 = *(*int32)(unsafe.Add(mBase, _consts[32]))
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
												v133 = base.I32_div_s(v131, v132)
												F__bt_parallel_scan_and_sort(m, v76, v109, v22, v86, v108, v133, int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return
												} else {
													v139 = F_shm_toc_lookup(m, l1, int64(-6917529027641081850), int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return
													} else {
														v143 = F_shm_toc_lookup(m, l1, int64(-6917529027641081851), int32(0))
														mBase = m.M
														v144 = m.ExcPending
														if v144 != 0 {
															return
														} else {
															v146 = *(*int32)(unsafe.Add(mBase, _consts[55]))
															v152 = v143 + v146<<(uint(int32(5))%32)
															v157 = F___memset(m, v139+v146<<(uint(int32(7))%32), int32(0), int32(128))
															mBase = m.M
															F_BufferUsageAccumDiff(m, v157, int32(4323704))
															mBase = m.M
															v161 = v152 + int32(24)
															v162 = int64(0)
															*(*int64)(unsafe.Add(mBase, uint32(v161))) = v162
															v165 = v152 + int32(16)
															*(*int64)(unsafe.Add(mBase, uint32(v165))) = v162
															v169 = v152 + int32(8)
															*(*int64)(unsafe.Add(mBase, uint32(v169))) = v162
															*(*int64)(unsafe.Add(mBase, uint32(v152))) = v162
															v175 = *(*int64)(unsafe.Add(mBase, _consts[49]))
															v177 = *(*int64)(unsafe.Add(mBase, _consts[50]))
															*(*int64)(unsafe.Add(mBase, uint32(v165))) = v175 - v177
															v181 = *(*int64)(unsafe.Add(mBase, _consts[53]))
															v183 = *(*int64)(unsafe.Add(mBase, _consts[54]))
															*(*int64)(unsafe.Add(mBase, uint32(v152))) = v181 - v183
															v187 = *(*int64)(unsafe.Add(mBase, _consts[51]))
															v189 = *(*int64)(unsafe.Add(mBase, _consts[52]))
															*(*int64)(unsafe.Add(mBase, uint32(v169))) = v187 - v189
															v193 = *(*int64)(unsafe.Add(mBase, _consts[47]))
															v195 = *(*int64)(unsafe.Add(mBase, _consts[48]))
															*(*int64)(unsafe.Add(mBase, uint32(v161))) = v193 - v195
															F_relation_close(m, v73, v72)
															mBase = m.M
															v199 = m.ExcPending
															if v199 != 0 {
																return
															} else {
																F_sequence_close(m, v67, v66)
																mBase = m.M
																v201 = m.ExcPending
																if v201 != 0 {
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
									v108 = v3
									v109 = v3
									v113 = F___memcpy(m, int32(4323704), int32(4323544), int32(128))
									mBase = m.M
									v116 = *(*int64)(unsafe.Add(mBase, _consts[47]))
									*(*int64)(unsafe.Add(mBase, _consts[48])) = v116
									v120 = *(*int64)(unsafe.Add(mBase, _consts[49]))
									*(*int64)(unsafe.Add(mBase, _consts[50])) = v120
									v124 = *(*int64)(unsafe.Add(mBase, _consts[51]))
									*(*int64)(unsafe.Add(mBase, _consts[52])) = v124
									v128 = *(*int64)(unsafe.Add(mBase, _consts[53]))
									*(*int64)(unsafe.Add(mBase, _consts[54])) = v128
									v131 = *(*int32)(unsafe.Add(mBase, _consts[32]))
									v132 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
									v133 = base.I32_div_s(v131, v132)
									F__bt_parallel_scan_and_sort(m, v76, v109, v22, v86, v108, v133, int32(0))
									mBase = m.M
									v136 = m.ExcPending
									if v136 != 0 {
										return
									} else {
										v139 = F_shm_toc_lookup(m, l1, int64(-6917529027641081850), int32(0))
										mBase = m.M
										v140 = m.ExcPending
										if v140 != 0 {
											return
										} else {
											v143 = F_shm_toc_lookup(m, l1, int64(-6917529027641081851), int32(0))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return
											} else {
												v146 = *(*int32)(unsafe.Add(mBase, _consts[55]))
												v152 = v143 + v146<<(uint(int32(5))%32)
												v157 = F___memset(m, v139+v146<<(uint(int32(7))%32), int32(0), int32(128))
												mBase = m.M
												F_BufferUsageAccumDiff(m, v157, int32(4323704))
												mBase = m.M
												v161 = v152 + int32(24)
												v162 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v161))) = v162
												v165 = v152 + int32(16)
												*(*int64)(unsafe.Add(mBase, uint32(v165))) = v162
												v169 = v152 + int32(8)
												*(*int64)(unsafe.Add(mBase, uint32(v169))) = v162
												*(*int64)(unsafe.Add(mBase, uint32(v152))) = v162
												v175 = *(*int64)(unsafe.Add(mBase, _consts[49]))
												v177 = *(*int64)(unsafe.Add(mBase, _consts[50]))
												*(*int64)(unsafe.Add(mBase, uint32(v165))) = v175 - v177
												v181 = *(*int64)(unsafe.Add(mBase, _consts[53]))
												v183 = *(*int64)(unsafe.Add(mBase, _consts[54]))
												*(*int64)(unsafe.Add(mBase, uint32(v152))) = v181 - v183
												v187 = *(*int64)(unsafe.Add(mBase, _consts[51]))
												v189 = *(*int64)(unsafe.Add(mBase, _consts[52]))
												*(*int64)(unsafe.Add(mBase, uint32(v169))) = v187 - v189
												v193 = *(*int64)(unsafe.Add(mBase, _consts[47]))
												v195 = *(*int64)(unsafe.Add(mBase, _consts[48]))
												*(*int64)(unsafe.Add(mBase, uint32(v161))) = v193 - v195
												F_relation_close(m, v73, v72)
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return
												} else {
													F_sequence_close(m, v67, v66)
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
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
			F_errmsg_internal(m, int32(453148), v9)
			mBase = m.M
			v69 = m.ExcPending
			if v69 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(463238), int32(1338), int32(296284))
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
	v18 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(v14^int32(-1))<<(uint(int32(2))%32))))
	v32 = v24
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	v52 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v66 = v52 + v47<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L12:
	;
	goto L13
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, _consts[5]))
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
	v115 = v69
	v118 = v68
	goto L16
L16:
	;
	if v115&int32(20) == int32(0) {
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
	v115 = v107
	v118 = v106
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
	v90 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90+(v85^int32(-1))<<(uint(int32(2))%32))))
	v104 = v96
	goto L20
L23:
	;
	goto L24
L24:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v104 = v98 + v85<<(uint(int32(13))%32) + int32(-8192)
	goto L20
L25:
	;
	goto L18
L26:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
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
	F_errmsg_internal(m, int32(646006), v12)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(462691), int32(1064), int32(96084))
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v14 = v12 & int32(4095)
	if l2 <= int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L55
	}
L2:
	;
	if base.Ui32(v14) <= base.Ui32(l2) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = F_CopyIndexTuple(m, l1)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+2)))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18))))
	v28 = int32(6)
	v30 = v18 + (v22 | v23<<(uint(int32(16))%32)) + l2*v28
	v32 = v30 + v28
	v37 = (v14 + (l2 ^ int32(-1))) * v28
	if v32 == v30 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v183 = l0 + int32(4)
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183))))
	*(*uint16)(unsafe.Add(mBase, uint32(v30)+4)) = uint16(v184)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v186
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v188&int32(32) == int32(0) {
		v211 = l1
		goto L52
	} else {
		goto L53
	}
L7:
	;
	goto L6
L8:
	;
	v41 = v32 + v37
	if base.Ui32(v30-v41) <= base.Ui32(int32(0)-v37<<(uint(int32(1))%32)) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v48 = F___memcpy(m, v32, v30, v37)
	mBase = m.M
	goto L6
L10:
	;
	goto L11
L11:
	;
	v51 = (v32 ^ v30) & int32(3)
	if base.Ui32(v32) < base.Ui32(v30) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	if v153 == int32(0) {
		goto L7
	} else {
		goto L48
	}
L13:
	;
	if base.Ui32(v131) <= base.Ui32(int32(3)) {
		v152 = v130
		v153 = v131
		v154 = v132
		goto L12
	} else {
		goto L44
	}
L14:
	;
	if v51 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v51 != 0 {
		v113 = v37
		goto L27
	} else {
		goto L28
	}
L17:
	;
	v152 = v30
	v153 = v37
	v154 = v32
	goto L12
L18:
	;
	goto L19
L19:
	;
	if v32&int32(3) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v130 = v30
	v131 = v37
	v132 = v32
	goto L13
L21:
	;
	goto L22
L22:
	;
	v58 = v30
	v59 = v37
	v60 = v32
	goto L23
L23:
	;
	if v59 == int32(0) {
		goto L7
	} else {
		goto L25
	}
L24:
	;
	v130 = v67
	v131 = v69
	v132 = v71
	goto L13
L25:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v64)
	v66 = int32(1)
	v67 = v58 + v66
	v69 = v59 - v66
	v71 = v60 + v66
	if v71&int32(3) != 0 {
		v58 = v67
		v59 = v69
		v60 = v71
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	if v113 == int32(0) {
		goto L7
	} else {
		goto L40
	}
L28:
	;
	if v41&int32(3) != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v78 = v37
	goto L32
L30:
	;
	v93 = v37
	goto L31
L31:
	;
	if base.Ui32(v93) <= base.Ui32(int32(3)) {
		v113 = v93
		goto L27
	} else {
		goto L36
	}
L32:
	;
	if v78 == int32(0) {
		goto L7
	} else {
		goto L34
	}
L33:
	;
	v93 = v84
	goto L31
L34:
	;
	v84 = v78 - int32(1)
	v85 = v32 + v84
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v84))))
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v87)
	if v85&int32(3) != 0 {
		v78 = v84
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v100 = v93
	goto L37
L37:
	;
	v104 = v100 - int32(4)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v30+v104)))
	*(*int32)(unsafe.Add(mBase, uint32(v32+v104))) = v107
	if base.Ui32(int32(3)) < base.Ui32(v104) {
		v100 = v104
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v113 = v104
	goto L27
L39:
	;
	goto L38
L40:
	;
	v120 = v113
	goto L41
L41:
	;
	v124 = v120 - int32(1)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v124))))
	*(*uint8)(unsafe.Add(mBase, uint32(v32+v124))) = uint8(v127)
	if v124 != 0 {
		v120 = v124
		goto L41
	} else {
		goto L43
	}
L42:
	;
	goto L7
L43:
	;
	goto L42
L44:
	;
	v137 = v130
	v138 = v131
	v139 = v132
	goto L45
L45:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v141
	v143 = int32(4)
	v144 = v137 + v143
	v146 = v139 + v143
	v148 = v138 - v143
	if base.Ui32(int32(3)) < base.Ui32(v148) {
		v137 = v144
		v138 = v148
		v139 = v146
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v152 = v144
	v153 = v148
	v154 = v146
	goto L12
L47:
	;
	goto L46
L48:
	;
	v159 = v152
	v160 = v153
	v161 = v154
	goto L49
L49:
	;
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	*(*uint8)(unsafe.Add(mBase, uint32(v161))) = uint8(v163)
	v165 = int32(1)
	v170 = v160 - v165
	if v170 != 0 {
		v159 = v159 + v165
		v160 = v170
		v161 = v161 + v165
		goto L49
	} else {
		goto L51
	}
L50:
	;
	goto L7
L51:
	;
	goto L50
L52:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v213
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v183))) = uint16(v215)
	m.G0 = v10 + int32(16)
	return v18
L53:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v193&int32(8192) == int32(0) {
		v211 = l1
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v206 = int32(6)
	v211 = l1 + (v198 | v199<<(uint(int32(16))%32)) + v193&int32(4095)*v206 - v206
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v14
	F_errmsg_internal(m, int32(440022), v10)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(465338), int32(1044), int32(308285))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
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
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
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
	var v229 int32
	_ = v229
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
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)))
	if v126&int32(8192) == int32(0) {
		v146 = v126
		goto L35
	} else {
		goto L36
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
	v86 = v84 & int32(8191)
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
	v91 = F_CreateTupleDescTruncatedCopy(m, v16, v81)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L6
	} else {
		goto L28
	}
L23:
	;
	if v86 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v122 = v87
	goto L16
L25:
	;
	v89 = F__emscripten_memcpy_bulkmem(m, v87, l2, v86)
	mBase = m.M
	goto L27
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+6)))
	if int32(0) <= base.I32_extend16_s(v97) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v101 = int32(8)
	goto L31
L30:
	;
	v101 = int32(16)
	goto L31
L31:
	;
	F_index_deform_tuple_internal(m, v91, v78+int32(32), v78, l2+v101, l2+int32(8), int32(base.Ui32(v97)>>(uint(int32(15))%32)))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v113 = F_index_form_tuple_context(m, v91, v78+int32(32), v78, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v113)+4)) = uint16(v115)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v117
	F_pfree(m, v91)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v122 = v113
	goto L16
L35:
	;
	if v65 <= v18 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
	if v131&int32(32) == int32(0) {
		v146 = v126
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v137 = v126 & int32(57344)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v137)
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+2)))
	v144 = (v139+int32(7))&int32(-8200) | v137
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v144)
	v146 = v144
	goto L35
L38:
	;
	m.G0 = v14 + int32(16)
	return v229
L39:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)) = uint16(v65)
	v150 = v146 | int32(8192)
	*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)) = uint16(v150)
	v229 = v122
	goto L38
L40:
	;
	goto L41
L41:
	;
	v159 = (v146&int32(8191)+int32(7))&int32(16376) + int32(8)
	v160 = F_palloc0(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+6)))
	v168 = (v162&int32(8191) + int32(7)) & int32(16376)
	if v168 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	F_pfree(m, v122)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L47
	}
L44:
	;
	v169 = F__emscripten_memcpy_bulkmem(m, v160, v122, v168)
	mBase = m.M
	v170 = v169
	goto L46
L45:
	;
	v170 = v160
	goto L46
L46:
	;
	goto L43
L47:
	;
	v174 = v18 | int32(4096)
	*(*uint16)(unsafe.Add(mBase, uint32(v170)+4)) = uint16(v174)
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+6)))
	v180 = int32(8192)
	v181 = v159 | v176&int32(49152) | v180
	*(*uint16)(unsafe.Add(mBase, uint32(v170)+6)) = uint16(v181)
	if v18&v180 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+2)))
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170))))
	v196 = v170 + (v185 | v186<<(uint(int32(16))%32))
	goto L50
L49:
	;
	v196 = v170 + v159&int32(8184) - int32(6)
	goto L50
L50:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v197&int32(32) == int32(0) {
		v220 = l1
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v196))) = v222
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+4)) = uint16(v224)
	v229 = v160
	goto L38
L52:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v202&int32(8192) == int32(0) {
		v220 = l1
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v215 = int32(6)
	v220 = l1 + (v207 | v208<<(uint(int32(16))%32)) + v202&int32(4095)*v215 - v215
	goto L51
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
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
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v247 int32
	_ = v247
	var v258 int32
	_ = v258
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
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
	if l6 < v24 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v19 + int32(16)
	return v282 & v275
L5:
	;
	v32 = l6
	v34 = v24
	goto L8
L6:
	;
	v258 = int32(0)
	goto L7
L7:
	;
	v275 = v258
	v282 = int32(0)
	goto L4
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v45 = v42 + v32*int32(48)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+2)))
	if v46&int32(3) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v258 = v239
	goto L7
L10:
	;
	v247 = v32 + int32(1)
	if v247 < v238 {
		v32 = v247
		v34 = v238
		goto L8
	} else {
		goto L67
	}
L11:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	v238 = v229
	v239 = v182
	goto L10
L12:
	;
	v275 = v221
	v282 = int32(1)
	goto L4
L13:
	;
	v221 = int32(0)
	goto L12
L14:
	;
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+4)))
	if l4 < v51 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if l7 == int32(0) {
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+6)))
	if v57 != int32(3) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v55 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v55)
	goto L13
L19:
	;
	if l5 != 0 {
		goto L13
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v62 = F_index_getattr_2(m, l2, v51, l3, v19+int32(15))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v238 = v34
	v239 = v51
	goto L10
L23:
	;
	return int32(0)
L24:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v66&int32(1572864) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v182 = int32(1)
	if base.B2i32(l1 == v182)&base.B2i32(v174 < int32(0)) != 0 {
		v221 = v182
		goto L12
	} else {
		goto L63
	}
L26:
	;
	v69 = int32(0)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v71 <= v69 {
		v113 = v69
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v123 = int32(1)
	v124 = v66 & v123
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	if v125 == v123 {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+15)))
	F__bt_binsrch_skiparray_skey(m, int32(0), l1, v62, v117, v113, v45, v19+int32(8))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L23
	} else {
		goto L35
	}
L30:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v84 = int32(0)
	goto L31
L31:
	;
	v94 = v74 + v84<<(uint(int32(5))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v95 == v32 {
		v113 = v94
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v113 = v94
	goto L29
L33:
	;
	v98 = v84 + int32(1)
	if v98 < v71 {
		v84 = v98
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v122 != 0 {
		v174 = v122
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v221 = v69
	goto L12
L37:
	;
	v174 = int32(1)
	goto L25
L38:
	;
	if v157&int32(2097152) != 0 {
		goto L60
	} else {
		goto L61
	}
L39:
	;
	if v124 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	if v124 != 0 {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(0)
	v157 = v66
	goto L38
L43:
	;
	goto L44
L44:
	;
	if v66&int32(33554432) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v134 = int32(-1)
	goto L47
L46:
	;
	v134 = int32(1)
	goto L47
L47:
	;
	v174 = v134
	goto L25
L48:
	;
	if v66&int32(33554432) != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v45)+44))
	v146 = F_FunctionCall2Coll(m, v140+v32*int32(28), v144, v62, v145)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L23
	} else {
		goto L54
	}
L51:
	;
	v139 = int32(1)
	goto L53
L52:
	;
	v139 = int32(-1)
	goto L53
L53:
	;
	v174 = v139
	goto L25
L54:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v148&int32(16777216) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v146 < int32(0) {
		goto L37
	} else {
		goto L58
	}
L56:
	;
	v155 = v146
	goto L57
L57:
	;
	if v155 != 0 {
		v174 = v155
		goto L25
	} else {
		goto L59
	}
L58:
	;
	v155 = int32(0) - v146
	goto L57
L59:
	;
	v157 = v148
	goto L38
L60:
	;
	v174 = int32(-1)
	goto L25
L61:
	;
	goto L62
L62:
	;
	v174 = int32(base.Ui32(v157)>>(uint(int32(22))%32)) & int32(1)
	goto L25
L63:
	;
	if base.B2i32(l1 == int32(-1))&base.B2i32(int32(0) < v174) != 0 {
		v221 = v182
		goto L12
	} else {
		goto L64
	}
L64:
	;
	if l5 != 0 {
		goto L13
	} else {
		goto L65
	}
L65:
	;
	if v174 == int32(0) {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	goto L13
L67:
	;
	goto L9
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v67 int32
	_ = v67
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
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
	v23 = v19 | v20<<(uint(int32(16))%32)
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
			v34 = F__emscripten_memcpy_bulkmem(m, v32, v11, v23)
			mBase = m.M
			v35 = v34
		} else {
			v35 = v32
		}
		v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+6)))
		v39 = v36&int32(-8192) | v31
		if int32(2) <= v16 {
			*(*uint16)(unsafe.Add(mBase, uint32(v35)+2)) = uint16(v19)
			*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v20)
			v44 = int32(8192)
			v45 = v16 | v44
			*(*uint16)(unsafe.Add(mBase, uint32(v35)+4)) = uint16(v45)
			v52 = v39 | v44
			v53 = v23 + v35
		} else {
			v52 = v39 & int32(57343)
			v53 = v35
		}
		*(*uint16)(unsafe.Add(mBase, uint32(v35)+6)) = uint16(v52)
		v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		if v55&int32(4095) != 0 {
			v60 = int32(0)
			v64 = v60
			v65 = v60
			v67 = v60
			v69 = v55
			for {
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
				if v73 <= v65 {
					v82 = int32(6)
					v84 = v53 + v67*v82
					v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)))
					v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
					v93 = v11 + (v85 | v86<<(uint(int32(16))%32)) + v64*v82
					v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
					*(*int32)(unsafe.Add(mBase, uint32(v84))) = v94
					v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)) = uint16(v96)
					v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
					v101 = v65
					v102 = v67 + int32(1)
					v103 = v100
				} else {
					v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0+int32(8)+v65<<(uint(int32(1))%32)))))
					if v64 != v78 {
						v82 = int32(6)
						v84 = v53 + v67*v82
						v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+2)))
						v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11))))
						v93 = v11 + (v85 | v86<<(uint(int32(16))%32)) + v64*v82
						v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
						*(*int32)(unsafe.Add(mBase, uint32(v84))) = v94
						v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v93)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v84)+4)) = uint16(v96)
						v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
						v101 = v65
						v102 = v67 + int32(1)
						v103 = v100
					} else {
						v101 = v65 + int32(1)
						v102 = v67
						v103 = v69
					}
				}
				v106 = v64 + int32(1)
				if base.Ui32(v106) < base.Ui32(v103&int32(4095)) {
					v64 = v106
					v65 = v101
					v67 = v102
					v69 = v103
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v35
		return
	}
}
