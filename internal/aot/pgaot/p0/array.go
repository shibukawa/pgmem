package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_array_agg_array_serialize(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
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
	var v167 int32
	_ = v167
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v7)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
		F_enlargeStringInfo(m, v7, int32(4))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v21 = int32(24)
			v23 = int32(65280)
			v25 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v18+v19))) = v14<<(uint(v21)%32) | v14&v23<<(uint(v25)%32) | (int32(base.Ui32(v14)>>(uint(v25)%32))&v23 | int32(base.Ui32(v14)>>(uint(v21)%32)))
			v37 = int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18 + v37
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v9)+80))
			F_enlargeStringInfo(m, v7, v37)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v47 = int32(24)
				v49 = int32(65280)
				v51 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v44+v45))) = v40<<(uint(v47)%32) | v40&v49<<(uint(v51)%32) | (int32(base.Ui32(v40)>>(uint(v51)%32))&v49 | int32(base.Ui32(v40)>>(uint(v47)%32)))
				v63 = int32(4)
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v44 + v63
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				F_enlargeStringInfo(m, v7, v63)
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v73 = int32(24)
					v75 = int32(65280)
					v77 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v70+v71))) = v66<<(uint(v73)%32) | v66&v75<<(uint(v77)%32) | (int32(base.Ui32(v66)>>(uint(v77)%32))&v75 | int32(base.Ui32(v66)>>(uint(v73)%32)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v70 + int32(4)
					v92 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v93 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
					F_pq_sendbytes(m, v7, v92, v93)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return int32(0)
					} else {
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
						F_enlargeStringInfo(m, v7, int32(4))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
							v101 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
							v103 = int32(24)
							v105 = int32(65280)
							v107 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v100+v101))) = v96<<(uint(v103)%32) | v96&v105<<(uint(v107)%32) | (int32(base.Ui32(v96)>>(uint(v107)%32))&v105 | int32(base.Ui32(v96)>>(uint(v103)%32)))
							v119 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v100 + v119
							v122 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
							F_enlargeStringInfo(m, v7, v119)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return int32(0)
							} else {
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								v127 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
								v129 = int32(24)
								v131 = int32(65280)
								v133 = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v126+v127))) = v122<<(uint(v129)%32) | v122&v131<<(uint(v133)%32) | (int32(base.Ui32(v122)>>(uint(v133)%32))&v131 | int32(base.Ui32(v122)>>(uint(v129)%32)))
								*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v126 + int32(4)
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
								if v148 != 0 {
									v149 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
									v153 = base.I32_div_s(v149+int32(7), int32(8))
									F_pq_sendbytes(m, v7, v148, v153)
									mBase = m.M
									v155 = m.ExcPending
									if v155 != 0 {
										return int32(0)
									} else {
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										F_enlargeStringInfo(m, v7, int32(4))
										mBase = m.M
										v159 = m.ExcPending
										if v159 != 0 {
											return int32(0)
										} else {
											v160 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
											v161 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
											v163 = int32(24)
											v165 = int32(65280)
											v167 = int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v160+v161))) = v156<<(uint(v163)%32) | v156&v165<<(uint(v167)%32) | (int32(base.Ui32(v156)>>(uint(v167)%32))&v165 | int32(base.Ui32(v156)>>(uint(v163)%32)))
											v179 = int32(4)
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v160 + v179
											v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
											F_enlargeStringInfo(m, v7, v179)
											mBase = m.M
											v185 = m.ExcPending
											if v185 != 0 {
												return int32(0)
											} else {
												v186 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
												v187 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
												v189 = int32(24)
												v191 = int32(65280)
												v193 = int32(8)
												*(*int32)(unsafe.Add(mBase, uint32(v186+v187))) = v182<<(uint(v189)%32) | v182&v191<<(uint(v193)%32) | (int32(base.Ui32(v182)>>(uint(v193)%32))&v191 | int32(base.Ui32(v182)>>(uint(v189)%32)))
												*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v186 + int32(4)
												F_pq_sendbytes(m, v7, v9+int32(32), v189)
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													F_pq_sendbytes(m, v7, v9+int32(56), int32(24))
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int32(0)
													} else {
														v219 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
														v220 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
														*(*int32)(unsafe.Add(mBase, uint32(v219))) = v220 << (uint(int32(2)) % 32)
														m.G0 = v7 + int32(16)
														return v219
													}
												}
											}
										}
									}
								} else {
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
									F_enlargeStringInfo(m, v7, int32(4))
									mBase = m.M
									v159 = m.ExcPending
									if v159 != 0 {
										return int32(0)
									} else {
										v160 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
										v163 = int32(24)
										v165 = int32(65280)
										v167 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v160+v161))) = v156<<(uint(v163)%32) | v156&v165<<(uint(v167)%32) | (int32(base.Ui32(v156)>>(uint(v167)%32))&v165 | int32(base.Ui32(v156)>>(uint(v163)%32)))
										v179 = int32(4)
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v160 + v179
										v182 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
										F_enlargeStringInfo(m, v7, v179)
										mBase = m.M
										v185 = m.ExcPending
										if v185 != 0 {
											return int32(0)
										} else {
											v186 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
											v187 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
											v189 = int32(24)
											v191 = int32(65280)
											v193 = int32(8)
											*(*int32)(unsafe.Add(mBase, uint32(v186+v187))) = v182<<(uint(v189)%32) | v182&v191<<(uint(v193)%32) | (int32(base.Ui32(v182)>>(uint(v193)%32))&v191 | int32(base.Ui32(v182)>>(uint(v189)%32)))
											*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v186 + int32(4)
											F_pq_sendbytes(m, v7, v9+int32(32), v189)
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												F_pq_sendbytes(m, v7, v9+int32(56), int32(24))
												mBase = m.M
												v217 = m.ExcPending
												if v217 != 0 {
													return int32(0)
												} else {
													v219 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
													v220 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
													*(*int32)(unsafe.Add(mBase, uint32(v219))) = v220 << (uint(int32(2)) % 32)
													m.G0 = v7 + int32(16)
													return v219
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
func F_array_append(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = F_fetch_array_arg_replace_nulls(m, l0, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v18 == int32(0) {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v22 = v21
		} else {
			v22 = v2
		}
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
		switch v23 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(1)
			v71 = int32(12)
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
			v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78)+4)))
			v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+6)))
			v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78)+7)))
			v82 = F_array_set_element(m, v14+v71, int32(1), v11+v71, v22, v18, int32(-1), v79, v80, v81)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				m.G0 = v11 + int32(16)
				return v82
			}
		case 1:
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v14)+36))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
			v28 = v25 + v27
			*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v28
			if base.B2i32(v27 < int32(0)) == base.B2i32(v28 < v25) {
				v71 = int32(12)
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+16))
				v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78)+4)))
				v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+6)))
				v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v78)+7)))
				v82 = F_array_set_element(m, v14+v71, int32(1), v11+v71, v22, v18, int32(-1), v79, v80, v81)
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					m.G0 = v11 + int32(16)
					return v82
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(401813), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(494052), int32(167), int32(426154))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
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
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(25655), int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(494052), int32(174), int32(426154))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
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
func F_array_cardinality(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_DatumGetAnyArrayP(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		if v11 == int32(-1) {
			v14 = int32(28)
		} else {
			v14 = int32(4)
		}
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v5+v14)))
		if v11 == int32(-1) {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
			v20 = F_ArrayGetNItems(m, v16, v19)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v20
			}
		} else {
			v25 = F_ArrayGetNItems(m, v16, v5+int32(16))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v25
			}
		}
	}
}
func F_array_contain_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v20 == int32(-1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L16
	} else {
		goto L78
	}
L2:
	;
	v23 = int32(40)
	goto L4
L3:
	;
	v23 = int32(12)
	goto L4
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0+v23)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v28 == int32(-1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v31 = int32(40)
	goto L7
L6:
	;
	v31 = int32(12)
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1+v31)))
	if v25 == v33 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v35 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L16
	} else {
		goto L74
	}
L11:
	;
	v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v48)+11)))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+10)))
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v48)+8)))
	if v49 == int32(-1) {
		goto L20
	} else {
		goto L21
	}
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v36 == v25 {
		v48 = v35
		v49 = v28
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v39 = F_lookup_type_cache(m, v25, int32(32))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	return int32(0)
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	if v43 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v39
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v48 = v39
	v49 = v47
	goto L11
L19:
	;
	v73 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+62)) = uint16(v73)
	v75 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+60)) = uint8(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v16)+48)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v48 + int32(76)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v85 == int32(-1) {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	F_deconstruct_expanded_array(m, l1)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L16
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	F_deconstruct_array(m, l1, v52, v51&int32(1), v50, v16+int32(40), v16+int32(36), v16+int32(32))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L16
	} else {
		goto L24
	}
L23:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v59
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v61
	goto L19
L24:
	;
	goto L19
L25:
	;
	v88 = int32(28)
	goto L27
L26:
	;
	v88 = int32(4)
	goto L27
L27:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0+v88)))
	if v85 == int32(-1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v97 = F_ArrayGetNItems(m, v90, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L16
	} else {
		goto L32
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v96 = v93
	goto L28
L30:
	;
	goto L31
L31:
	;
	v96 = l0 + int32(16)
	goto L28
L32:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v99 == int32(-1) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = int32(1)
	if v97 <= int32(0) {
		v272 = l3
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v102 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = int64(0)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v135 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v105 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v104
	v158 = v105
	goto L33
L38:
	;
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+12)) = int64(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	if v112 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v111 + (v115<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v158 = int32(0)
	goto L33
L41:
	;
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v111 + v112
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v158 = v111 + v127<<(uint(int32(3))%32) + int32(16)
	goto L33
L43:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l0 + (v138<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v158 = int32(0)
	goto L33
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l0 + v135
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v158 = l0 + v150<<(uint(int32(3))%32) + int32(16)
	goto L33
L46:
	;
	m.G0 = v16 + int32(80)
	return v272
L47:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v16)+32))
	v171 = int32(0)
	goto L48
L48:
	;
	v186 = F_array_iter_next(m, v16+int32(12), v16+int32(11), v171, v52, v51&int32(1), v50)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L16
	} else {
		goto L50
	}
L49:
	;
	v272 = l3
	goto L46
L50:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+11)))
	if v188 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v267 = v171 + int32(1)
	if v267 != v97 {
		v171 = v267
		goto L48
	} else {
		goto L73
	}
L52:
	;
	v191 = int32(0)
	if v191 < v165 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L54
L54:
	;
	if l3 == int32(0) {
		goto L51
	} else {
		goto L72
	}
L55:
	;
	if l3 != 0 {
		goto L51
	} else {
		goto L71
	}
L56:
	;
	v194 = v191
	goto L59
L57:
	;
	goto L58
L58:
	;
	if l3 == int32(0) {
		goto L51
	} else {
		goto L70
	}
L59:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v164+v194<<(uint(int32(2))%32))))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v16)+36))
	if v211 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L58
L61:
	;
	v231 = v194 + int32(1)
	if v231 != v165 {
		v194 = v231
		goto L59
	} else {
		goto L69
	}
L62:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v211))))
	if v213 != 0 {
		goto L61
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v214 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v214)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v210
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+68)) = uint8(v214)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v186
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+60)) = uint8(v214)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	v226 = m.T0[v225].(func(*base.Module, int32) int32)(m, v16+int32(44))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L16
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+60)))
	if v228 != 0 {
		goto L61
	} else {
		goto L67
	}
L67:
	;
	if v226 != 0 {
		goto L55
	} else {
		goto L68
	}
L68:
	;
	goto L61
L69:
	;
	goto L60
L70:
	;
	v272 = int32(0)
	goto L46
L71:
	;
	v272 = int32(1)
	goto L46
L72:
	;
	v272 = int32(0)
	goto L46
L73:
	;
	goto L49
L74:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L16
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(162107), int32(0))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L16
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(493945), int32(4408), int32(364755))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L16
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	v309 = F_format_type_be(m, v25)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L16
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v309
	F_errmsg(m, int32(189157), v16)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L16
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(493945), int32(4426), int32(364755))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L16
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_contains_nulls(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = l0 + int32(16)
	v14 = F_ArrayGetNItems(m, v11, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v24 = v13 + v19<<(uint(int32(3))%32)
	goto L8
L7:
	;
	v24 = int32(0)
	goto L8
L8:
	;
	if v14 <= int32(7) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return v73
L10:
	;
	if v42 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	v42 = v14
	v43 = v24
	goto L10
L12:
	;
	goto L13
L13:
	;
	v29 = v24
	v30 = v14
	goto L14
L14:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v33 != int32(255) {
		v73 = int32(1)
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v42 = v41
	v43 = v37
	goto L10
L16:
	;
	v37 = v29 + int32(1)
	v41 = v30 - int32(8)
	if base.Ui32(int32(15)) < base.Ui32(v30) {
		v29 = v37
		v30 = v41
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return int32(0)
L19:
	;
	goto L20
L20:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	v53 = v42
	v55 = int32(1)
	goto L21
L21:
	;
	v60 = base.B2i32(v55&v51 == int32(0))
	if v55&v51 == int32(0) {
		v73 = v60
		goto L9
	} else {
		goto L23
	}
L22:
	;
	v73 = v60
	goto L9
L23:
	;
	v63 = int32(1)
	if v63 < v53 {
		v53 = v53 - v63
		v55 = v55 << (uint(v63) % 32)
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
}
func F_array_fill_with_lower_bounds(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(67108994))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(301975), int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(493945), int32(6010), int32(171994))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
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
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
		if v8 == int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(67108994))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(301975), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(493945), int32(6010), int32(171994))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v12 = F_pg_detoast_datum(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v17 = F_pg_detoast_datum(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
					if v19 == int32(0) {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v23 = v22
					} else {
						v23 = int32(0)
					}
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v26 = F_get_fn_expr_argtype(m, v24, int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						if v26 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								F_errmsg_internal(m, int32(64940), int32(0))
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493945), int32(6028), int32(171994))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = F_array_fill_internal(m, v12, v17, v23, v19, v26, l0)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								return v30
							}
						}
					}
				}
			}
		}
	}
}
func F_array_iterator(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = F_ArrayGetNItems(m, v8, l0+int32(16))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 < int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L31
	}
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v19 = F_array_contains_nulls(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L27
	}
L7:
	;
	if v19 != 0 {
		goto L3
	} else {
		goto L8
	}
L8:
	;
	if l3 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	goto L11
L10:
	;
	goto L11
L11:
	;
	if int32(0) < v11 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return v76
L13:
	;
	if v18 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v76 = int32(0)
	goto L12
L16:
	;
	v31 = v18
	goto L18
L17:
	;
	v31 = (v15<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L18
L18:
	;
	v33 = l0 + v31
	v37 = v11
	goto L19
L19:
	;
	v41 = F_DirectFunctionCall2Coll(m, l1, int32(0), v33, l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	if v41 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if l3 == int32(0) {
		v76 = int32(1)
		goto L12
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v57 = int32(1)
	if v57 < v37 {
		v33 = v33 + (int32(base.Ui32(v49)>>(uint(int32(2))%32))+int32(3))&int32(2147483644)
		v37 = v37 - v57
		goto L19
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v33
	return int32(1)
L26:
	;
	goto L20
L27:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errmsg(m, int32(312607), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(495234), int32(46), int32(208762))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
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
	F_errcode(m, int32(67108994))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(152400), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(495234), int32(50), int32(208762))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_prepend_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	if v5 != int32(464) {
		v23 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if v10 == int32(0) {
			v23 = v2
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			if v13 != int32(8) {
				v23 = v2
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
				if v16 != 0 {
					v23 = v2
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
					if v18 == v19 {
						v21 = v10
					} else {
						v21 = int32(0)
					}
					v23 = v21
				}
			}
		}
	}
	return v23
}
func F_array_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v310 int32
	_ = v310
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v391 int32
	_ = v391
	var v408 int32
	_ = v408
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v589 int32
	_ = v589
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v646 int32
	_ = v646
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v665 int32
	_ = v665
	var v670 int32
	_ = v670
	v19 = m.G0
	v21 = v19 - int32(176)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L9
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L8
	} else {
		goto L140
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L8
	} else {
		goto L136
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L8
	} else {
		goto L132
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L8
	} else {
		goto L127
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L8
	} else {
		goto L121
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L8
	} else {
		goto L117
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L8
	} else {
		goto L113
	}
L8:
	;
	return int32(0)
L9:
	;
	if int32(0) <= v27 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	if base.Ui32(int32(7)) <= base.Ui32(v27) {
		goto L7
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L8
	} else {
		goto L109
	}
L13:
	;
	v36 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v36) {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v41 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L8
	} else {
		goto L17
	}
L16:
	;
	if v27 != 0 {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	if v41 == v24 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(int32(9999)) < base.Ui32(v41) {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	if base.Ui32(v24) <= base.Ui32(int32(9999)) {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	v50 = int32(0)
	goto L24
L22:
	;
	goto L23
L23:
	;
	v106 = F_ArrayGetNItems(m, v27, v21+int32(128))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L29
	}
L24:
	;
	v68 = v50 << (uint(int32(2)) % 32)
	v73 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L26
	}
L25:
	;
	goto L23
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68+(v21+int32(128))))) = v73
	v80 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+int32(96)+v68))) = v80
	v84 = v50 + int32(1)
	if v84 != v27 {
		v50 = v84
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	F_ArrayCheckBounds(m, v27, v21+int32(128), v21+int32(96))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)+16))
	if v115 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	if v106 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L32:
	;
	F_get_type_io_data(m, v24, int32(2), v131+int32(4), v131+int32(6), v131+int32(7), v131+int32(8), v131+int32(12), v131+int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L8
	} else {
		goto L38
	}
L33:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v114)+20))
	v120 = F_MemoryContextAlloc(m, v118, int32(48))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L8
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	if v129 == v24 {
		v157 = v115
		goto L31
	} else {
		goto L37
	}
L36:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v122)+16)) = v120
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v24 ^ int32(-1)
	v131 = v125
	goto L32
L37:
	;
	v131 = v115
	goto L32
L38:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v131)+16))
	if v147 == int32(0) {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	F_fmgr_info_cxt(m, v147, v131+int32(20), v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L8
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v24
	v157 = v131
	goto L31
L41:
	;
	m.G0 = v21 + int32(176)
	return v494
L42:
	;
	v162 = F_palloc0(m, int32(16))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+6)))
	v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v157)+4)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+7)))
	v175 = F_palloc(m, v106<<(uint(int32(2))%32))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L46
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v162))) = int64(64)
	v494 = v162
	goto L41
L46:
	;
	v177 = F_palloc(m, v106)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v179 = int32(0)
	if v106 <= v179 {
		v421 = v179
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v461 = F_palloc0(m, v460)
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L8
	} else {
		goto L97
	}
L49:
	;
	v457 = int32(0)
	v460 = v421 + (v27<<(uint(int32(3))%32)+int32(23))&int32(120)
	goto L48
L50:
	;
	v183 = v157 + int32(20)
	v188 = v179
	goto L51
L51:
	;
	v203 = F_pq_getmsgint(m, v25, int32(4))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L53
	}
L52:
	;
	v252 = int32(0)
	v257 = v252
	v259 = v252
	v264 = v252
	goto L65
L53:
	;
	if v203 < int32(-1) {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v207-v208 < v203 {
		goto L3
	} else {
		goto L55
	}
L55:
	;
	if v203 == int32(-1) {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v246 = v188 + int32(1)
	if v246 != v106 {
		v188 = v246
		goto L51
	} else {
		goto L63
	}
L57:
	;
	v217 = F_ReceiveFunctionCall(m, v183, int32(0), v169, v23)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+168)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+160)) = v208 + v223
	*(*int32)(unsafe.Add(mBase, uint32(v21)+164)) = v203
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v208 + v203
	v236 = F_ReceiveFunctionCall(m, v183, v21+int32(160), v169, v23)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L8
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175+v188<<(uint(int32(2))%32)))) = v217
	v221 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v188+v177))) = uint8(v221)
	goto L56
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175+v188<<(uint(int32(2))%32)))) = v236
	v240 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v188+v177))) = uint8(v240)
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v21)+172))
	if v242 != v203 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	goto L56
L63:
	;
	goto L52
L64:
	;
	v408 = base.I32_div_s(v106+int32(7), int32(8))
	v415 = (v408 + v27<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v457 = v415
	v460 = v391 + v415
	goto L48
L65:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257+v177))))
	if v274 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	if v310&int32(1) == int32(0) {
		v421 = v377
		goto L49
	} else {
		goto L96
	}
L67:
	;
	v279 = v257
	goto L70
L68:
	;
	v303 = v257
	v310 = v264
	goto L69
L69:
	;
	if base.B2i32(v171 == int32(-1)) == int32(0) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v296 = v279 + int32(1)
	if v296 == v106 {
		v391 = v259
		goto L64
	} else {
		goto L72
	}
L71:
	;
	v303 = v296
	v310 = int32(1)
	goto L69
L72:
	;
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v177))))
	if v299 != 0 {
		v279 = v296
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v364 = v362 + v259
	switch v172 - int32(99) {
	case 0:
		v377 = v364
		goto L90
	case 1:
		goto L92
	default:
		goto L91
	case 6:
		goto L93
	}
L75:
	;
	if int32(0) < v171 {
		v362 = v171
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v332 = v175 + v303<<(uint(int32(2))%32)
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v334 = F_pg_detoast_datum(m, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L8
	} else {
		goto L79
	}
L78:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v175+v303<<(uint(int32(2))%32))))
	v327 = F_strlen(m, v326)
	mBase = m.M
	v362 = v327 + int32(1)
	goto L74
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v332))) = v334
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334))))
	if v337 == int32(1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v334)+1)))
	if base.Ui32((v341-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v362 = int32(6)
		goto L74
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	if v337&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L83:
	;
	v348 = int32(18)
	if v341&int32(255) == v348 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v354 = v348
	goto L86
L85:
	;
	v354 = int32(2)
	goto L86
L86:
	;
	v362 = v354
	goto L74
L87:
	;
	v362 = int32(base.Ui32(v337) >> (uint(int32(1)) % 32))
	goto L74
L88:
	;
	goto L89
L89:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v334)))
	v362 = int32(base.Ui32(v359) >> (uint(int32(2)) % 32))
	goto L74
L90:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v377) {
		goto L1
	} else {
		goto L94
	}
L91:
	;
	v377 = (v364 + int32(1)) & int32(-2)
	goto L90
L92:
	;
	v377 = (v364 + int32(7)) & int32(-8)
	goto L90
L93:
	;
	v377 = (v364 + int32(3)) & int32(-4)
	goto L90
L94:
	;
	v381 = v303 + int32(1)
	if v381 != v106 {
		v257 = v381
		v259 = v377
		v264 = v310
		goto L65
	} else {
		goto L95
	}
L95:
	;
	goto L66
L96:
	;
	v391 = v377
	goto L64
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v461)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v461)+8)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v461)+4)) = v27
	v466 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v461))) = v460 << (uint(v466) % 32)
	v470 = v461 + int32(16)
	v474 = v27 << (uint(v466) % 32)
	if v474 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	if v474 != 0 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v475 = F__emscripten_memcpy_bulkmem(m, v470, v21+int32(128), v474)
	mBase = m.M
	v476 = v475
	goto L101
L100:
	;
	v476 = v470
	goto L101
L101:
	;
	goto L98
L102:
	;
	v482 = int32(1)
	F_CopyArrayEls(m, v461, v175, v177, v106, v171, v170&v482, base.I32_extend8_s(v172), v482)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L8
	} else {
		goto L106
	}
L103:
	;
	v480 = F__emscripten_memcpy_bulkmem(m, v476+v474, v21+int32(96), v474)
	mBase = m.M
	goto L105
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	F_pfree(m, v175)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	F_pfree(m, v177)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	v494 = v461
	goto L41
L109:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L8
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v27
	F_errmsg(m, int32(481223), v21)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	F_errfinish(m, int32(493945), int32(1301), int32(36231))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L113:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v27
	F_errmsg(m, int32(663843), v21+int32(16))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L8
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(493945), int32(1306), int32(36231))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L8
	} else {
		goto L118
	}
L118:
	;
	F_errmsg(m, int32(156637), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L8
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(493945), int32(1312), int32(36231))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L8
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L8
	} else {
		goto L122
	}
L122:
	;
	v575 = F_format_type_extended(m, v41, int32(-1), int32(2))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L8
	} else {
		goto L123
	}
L123:
	;
	v579 = F_format_type_extended(m, v24, int32(-1), int32(2))
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L8
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+92)) = v579
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v575
	*(*int32)(unsafe.Add(mBase, uint32(v21)+80)) = v41
	F_errmsg(m, int32(658974), v21+int32(80))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L8
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(493945), int32(1340), int32(36231))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L8
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v601 = m.ExcPending
	if v601 != 0 {
		goto L8
	} else {
		goto L128
	}
L128:
	;
	v602 = F_format_type_be(m, v24)
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L8
	} else {
		goto L129
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v602
	F_errmsg(m, int32(189614), v21+int32(32))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L8
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(493945), int32(1379), int32(36231))
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L8
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L8
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(403977), int32(0))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(493945), int32(1481), int32(17839))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L8
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(50462850))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L8
	} else {
		goto L137
	}
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v188 + int32(1)
	F_errmsg(m, int32(467588), v21-int32(-64))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L8
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(493945), int32(1510), int32(17839))
	mBase = m.M
	v651 = m.ExcPending
	if v651 != 0 {
		goto L8
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L140:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L8
	} else {
		goto L141
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = int32(1073741823)
	F_errmsg(m, int32(663745), v21+int32(48))
	mBase = m.M
	v665 = m.ExcPending
	if v665 != 0 {
		goto L8
	} else {
		goto L142
	}
L142:
	;
	F_errfinish(m, int32(493945), int32(1534), int32(17839))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L8
	} else {
		goto L143
	}
L143:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_sort(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = int32(0)
		v9 = F_array_sort_internal(m, v3, v7, v7, l0)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_array_sort_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v16 <= int32(0) {
		v190 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L9
	} else {
		goto L57
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L9
	} else {
		goto L52
	}
L3:
	;
	m.G0 = v14 + int32(32)
	return v190
L4:
	;
	v20 = l0 + int32(16)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v21 < int32(2) {
		v190 = l0
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
	if v26 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v31 = F_MemoryContextAllocZero(m, v29, int32(60))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v37 = v26
	goto L8
L8:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v38 != v39 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	return int32(0)
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+16)) = v31
	v37 = v31
	goto L8
L11:
	;
	v42 = F_lookup_type_cache(m, v38, int32(6))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if v16 != int32(1) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v38
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v37)+4)) = uint16(v45)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+6)) = uint8(v47)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+11)))
	*(*uint8)(unsafe.Add(mBase, uint32(v37)+7)) = uint8(v49)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v42)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+48)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v42)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+52)) = v53
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v42)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v37)+56)) = v55
	goto L13
L15:
	;
	v77 = int32(0)
	v79 = *(*int32)(unsafe.Add(mBase, _consts[326]))
	v81 = F_tuplesort_begin_datum(m, v74, v73, v24, l2, v79, v77)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L9
	} else {
		goto L27
	}
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v37)+56))
	if v60 == int32(0) {
		goto L2
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	if l1 != 0 {
		goto L23
	} else {
		goto L24
	}
L19:
	;
	if l1 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v65 = int32(1073)
	goto L22
L21:
	;
	v65 = int32(1072)
	goto L22
L22:
	;
	v73 = v65
	v74 = v60
	goto L15
L23:
	;
	v68 = int32(52)
	goto L25
L24:
	;
	v68 = int32(48)
	goto L25
L25:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v37+v68)))
	if v70 == int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v73 = v70
	v74 = v38
	goto L15
L27:
	;
	v85 = F_array_create_iterator(m, l0, v16-int32(1), v37)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	v91 = F_array_iterate(m, v85, v14+int32(28), v14+int32(27))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	if v91 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	goto L33
L31:
	;
	goto L32
L32:
	;
	F_array_free_iterator(m, v85)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L9
	} else {
		goto L38
	}
L33:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	F_tuplesort_putdatum(m, v81, v104, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L9
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	v112 = F_array_iterate(m, v85, v14+int32(28), v14+int32(27))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L9
	} else {
		goto L36
	}
L36:
	;
	if v112 != 0 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	F_tuplesort_performsort(m, v81)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	v131 = int32(0)
	v137 = F_tuplesort_getdatum(m, v81, int32(1), v131, v14+int32(28), v14+int32(27), v131)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L9
	} else {
		goto L40
	}
L40:
	;
	if v137 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v140 = v77
	goto L44
L42:
	;
	v166 = v77
	goto L43
L43:
	;
	F_tuplesort_end(m, v81)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L9
	} else {
		goto L49
	}
L44:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+27)))
	v153 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v154 = F_accumArrayResultAny(m, v140, v150, v151, v74, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L9
	} else {
		goto L46
	}
L45:
	;
	v166 = v154
	goto L43
L46:
	;
	v157 = int32(0)
	v163 = F_tuplesort_getdatum(m, v81, int32(1), v157, v14+int32(28), v14+int32(27), v157)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L9
	} else {
		goto L47
	}
L47:
	;
	if v163 != 0 {
		v140 = v154
		goto L44
	} else {
		goto L48
	}
L48:
	;
	goto L45
L49:
	;
	v179 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v180 = F_makeArrayResultAny(m, v166, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	v182 = F_pg_detoast_datum(m, v180)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v16<<(uint(int32(2))%32)+v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v182+v184<<(uint(int32(2))%32))+16)) = v188
	v190 = v182
	goto L3
L52:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	v212 = F_format_type_be(m, v38)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v212
	F_errmsg(m, int32(193150), v14+int32(16))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(494052), int32(1951), int32(310341))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L9
	} else {
		goto L58
	}
L58:
	;
	v232 = F_format_type_be(m, v38)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L9
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v232
	F_errmsg(m, int32(189368), v14)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L9
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(494052), int32(1964), int32(310341))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L9
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_to_tsvector(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
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
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
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
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v359 int32
	_ = v359
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v381 int32
	_ = v381
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
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
	var v421 int32
	_ = v421
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_deconstruct_array_builtin(m, v17, int32(25), v14+int32(12), v14+int32(8), v14+int32(4))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v31 <= int32(0) {
		v359 = v31
		v367 = int32(8)
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L107
	}
L5:
	;
	v370 = v367 + v359<<(uint(int32(2))%32)
	v371 = F_palloc0(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L92
	}
L6:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
	v37 = v2
	goto L8
L7:
	;
	v79 = int32(1)
	if v31 != v79 {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v35))))
	if v48 == int32(1) {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v34+v37<<(uint(int32(2))%32))))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	if v55&int32(-4) != int32(16) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v61 = v37 + int32(1)
	if v61 == v31 {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	goto L9
L14:
	;
	v37 = v61
	goto L8
L15:
	;
	F_errcode(m, int32(369098882))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(155859), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(495198), int32(776), int32(207895))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	F_pg_qsort(m, v82, v31, int32(4), int32(1538))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v244 = v31
	goto L21
L21:
	;
	if v244 <= int32(0) {
		v359 = v244
		v367 = int32(8)
		goto L5
	} else {
		goto L79
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v87) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v93 = v79
	v94 = int32(0)
	goto L26
L24:
	;
	v232 = v87
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v232
	v244 = v232
	goto L21
L26:
	;
	v103 = int32(2)
	v105 = v91 + v93<<(uint(v103)%32)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	v116 = int32(1)
	v117 = v115 + v116
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v120 = v118 & v116
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v91+v94<<(uint(v103)%32))))
	if v118 == v116 {
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v232 = v222 + int32(1)
	goto L25
L28:
	;
	v225 = v93 + int32(1)
	if v225 != v87 {
		v93 = v225
		v94 = v222
		goto L26
	} else {
		goto L78
	}
L29:
	;
	if v211 == int32(0) {
		v222 = v94
		goto L28
	} else {
		goto L76
	}
L30:
	;
	v150 = int32(1)
	v151 = v121 + v150
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v154 = v152 & v150
	if v152 == v150 {
		goto L42
	} else {
		goto L43
	}
L31:
	;
	v124 = int32(4)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if v126&int32(254) == int32(2) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v139 = int32(1)
	if v120 != 0 {
		v149 = int32(base.Ui32(v118)>>(uint(v139)%32)) - v139
		goto L30
	} else {
		goto L40
	}
L34:
	;
	v135 = v124
	goto L36
L35:
	;
	v135 = base.B2i32(v126 == int32(18)) << (uint(v124) % 32)
	goto L36
L36:
	;
	if v126 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v138 = v124
	goto L39
L38:
	;
	v138 = v135
	goto L39
L39:
	;
	v149 = v138
	goto L30
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	v149 = int32(base.Ui32(v143)>>(uint(int32(2))%32)) - int32(4)
	goto L30
L41:
	;
	if v149 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L42:
	;
	v157 = int32(4)
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v159&int32(254) == int32(2) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v172 = int32(1)
	if v154 != 0 {
		v182 = int32(base.Ui32(v152)>>(uint(v172)%32)) - v172
		goto L41
	} else {
		goto L51
	}
L45:
	;
	v168 = v157
	goto L47
L46:
	;
	v168 = base.B2i32(v159 == int32(18)) << (uint(v157) % 32)
	goto L47
L47:
	;
	if v159 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v171 = v157
	goto L50
L49:
	;
	v171 = v168
	goto L50
L50:
	;
	v182 = v171
	goto L41
L51:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v182 = int32(base.Ui32(v176)>>(uint(int32(2))%32)) - int32(4)
	goto L41
L52:
	;
	v186 = int32(0)
	if v186 < v182 {
		goto L55
	} else {
		goto L56
	}
L53:
	;
	goto L54
L54:
	;
	if v182 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v189 = int32(-1)
	goto L57
L56:
	;
	v189 = v186
	goto L57
L57:
	;
	v211 = v189
	goto L29
L58:
	;
	v211 = base.B2i32(int32(0) < v149)
	goto L29
L59:
	;
	goto L60
L60:
	;
	if v120 != 0 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v211 = v209
	goto L29
L62:
	;
	v196 = v117
	goto L64
L63:
	;
	v196 = v115 + int32(4)
	goto L64
L64:
	;
	if v154 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v199 = v151
	goto L67
L66:
	;
	v199 = v121 + int32(4)
	goto L67
L67:
	;
	if base.Ui32(v149) < base.Ui32(v182) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v201 = v149
	goto L70
L69:
	;
	v201 = v182
	goto L70
L70:
	;
	v202 = F_memcmp(m, v196, v199, v201)
	mBase = m.M
	if v202 != 0 {
		v209 = v202
		goto L61
	} else {
		goto L71
	}
L71:
	;
	if v182 == v149 {
		v209 = int32(0)
		goto L61
	} else {
		goto L72
	}
L72:
	;
	if v149 < v182 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v208 = int32(-1)
	goto L75
L74:
	;
	v208 = int32(1)
	goto L75
L75:
	;
	v209 = v208
	goto L61
L76:
	;
	v215 = v94 + int32(1)
	if v93 == v215 {
		v222 = v93
		goto L28
	} else {
		goto L77
	}
L77:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	*(*int32)(unsafe.Add(mBase, uint32(v91+v215<<(uint(int32(2))%32)))) = v220
	v222 = v215
	goto L28
L78:
	;
	goto L27
L79:
	;
	v255 = int32(0)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	if base.Ui32(int32(4)) <= base.Ui32(v244) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v264 = v255
	v265 = v255
	v270 = v2
	goto L83
L81:
	;
	v305 = v255
	v306 = v255
	goto L82
L82:
	;
	v316 = v244 & int32(3)
	if v316 != 0 {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	v274 = int32(2)
	v276 = v256 + v264<<(uint(v274)%32)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v276)+8))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v276)+12))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v298 = v265 + int32(base.Ui32(v278)>>(uint(v274)%32)) + int32(base.Ui32(v283)>>(uint(v274)%32)) + int32(base.Ui32(v288)>>(uint(v274)%32)) + int32(base.Ui32(v293)>>(uint(v274)%32)) - int32(16)
	v299 = int32(4)
	v300 = v264 + v299
	v302 = v270 + v299
	if v302 != v244&int32(2147483644) {
		v264 = v300
		v265 = v298
		v270 = v302
		goto L83
	} else {
		goto L85
	}
L84:
	;
	v305 = v300
	v306 = v298
	goto L82
L85:
	;
	goto L84
L86:
	;
	v318 = v305
	v319 = v306
	v323 = v255
	goto L89
L87:
	;
	v345 = v306
	goto L88
L88:
	;
	v359 = v244
	v367 = v345 + int32(8)
	goto L5
L89:
	;
	v328 = int32(2)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v256+v318<<(uint(v328)%32))))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)))
	v337 = v319 + int32(base.Ui32(v332)>>(uint(v328)%32)) - int32(4)
	v338 = int32(1)
	v341 = v323 + v338
	if v341 != v316 {
		v318 = v318 + v338
		v319 = v337
		v323 = v341
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v345 = v337
	goto L88
L91:
	;
	goto L90
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371))) = v370 << (uint(int32(2)) % 32)
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = v376
	if int32(0) < v376 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v381 = v371 + int32(8)
	v387 = int32(0)
	v388 = v381 + v376<<(uint(int32(2))%32)
	goto L96
L94:
	;
	goto L95
L95:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v446 != v17 {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	v397 = int32(2)
	v398 = v387 << (uint(v397) % 32)
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v398+v399)))
	v402 = int32(4)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v408 = int32(base.Ui32(v404)>>(uint(v397)%32)) - v402
	if v408 != 0 {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	goto L95
L98:
	;
	v411 = v398 + v381
	v412 = int32(1)
	v415 = v408 << (uint(v412) % 32) & int32(4094)
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = v415 | v416&int32(-4096)
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v371)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v411))) = (v410-(v381+v421<<(uint(int32(2))%32)))<<(uint(int32(12))%32) | v415
	v432 = v387 + v412
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v432 < v433 {
		v387 = v432
		v388 = v410 + v408
		goto L96
	} else {
		goto L102
	}
L99:
	;
	v409 = F__emscripten_memcpy_bulkmem(m, v388, v401+v402, v408)
	mBase = m.M
	v410 = v409
	goto L101
L100:
	;
	v410 = v388
	goto L101
L101:
	;
	goto L98
L102:
	;
	goto L97
L103:
	;
	F_pfree(m, v17)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	m.G0 = v14 + int32(16)
	return v371
L106:
	;
	goto L105
L107:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errmsg(m, int32(152313), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	F_errfinish(m, int32(495198), int32(771), int32(207895))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_deconstruct_array(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	v14 = m.G0
	v15 = int32(16)
	v16 = v14 - v15
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = l0 + v15
	v21 = F_ArrayGetNItems(m, v18, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = F_palloc(m, v21<<(uint(int32(2))%32))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v25
	if l5 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = F_palloc0(m, v21)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v31 = int32(0)
	goto L6
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v21
	if int32(0) < v21 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v28
	v31 = v28
	goto L6
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v35 << (uint(int32(3)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	m.G0 = v16 + int32(16)
	return
L11:
	;
	v41 = v20 + v37
	goto L13
L12:
	;
	v41 = int32(0)
	goto L13
L13:
	;
	if v40 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v46 = v40
	goto L16
L15:
	;
	v46 = (v37 + int32(23)) & int32(-8)
	goto L16
L16:
	;
	v48 = int32(1)
	v54 = v48
	v58 = l0 + v46
	v59 = v41
	v60 = int32(0)
	goto L17
L17:
	;
	if v59 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L10
L19:
	;
	v175 = int32(1)
	v177 = v54 << (uint(v175) % 32)
	v179 = base.B2i32(v177 == int32(256))
	if v177 == int32(256) {
		goto L62
	} else {
		goto L63
	}
L20:
	;
	if l2 != 0 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v54&v69 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v60<<(uint(int32(2))%32)))) = int32(0)
	if v31 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v77 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60+v31))) = uint8(v77)
	v174 = v58
	goto L19
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(60347), int32(0))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(493945), int32(3669), int32(24191))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	switch l3 - int32(99) {
	case 0:
		v174 = v160
		goto L19
	case 1:
		goto L60
	default:
		goto L59
	case 6:
		goto L61
	}
L31:
	;
	switch l1 - v48 {
	case 0:
		goto L35
	case 1:
		goto L38
	default:
		goto L36
	case 3:
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v60<<(uint(int32(2))%32)))) = v58
	if int32(0) < l1 {
		v160 = l1 + v58
		goto L30
	} else {
		goto L42
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25+v60<<(uint(int32(2))%32)))) = v114
	v160 = l1 + v58
	goto L30
L35:
	;
	v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(v58))))
	v114 = v113
	goto L34
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v114 = v99
	goto L34
L38:
	;
	v98 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58))))
	v114 = v98
	goto L34
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = l1
	F_errmsg_internal(m, int32(482718), v16)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(326157), int32(70), int32(67716))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	if l1 == int32(-1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
	if v126 == int32(1) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	v155 = F_strlen(m, v58)
	mBase = m.M
	v160 = v155 + v58 + int32(1)
	goto L30
L46:
	;
	v129 = int32(6)
	v131 = int32(18)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+1)))
	if v133 == v131 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	goto L48
L48:
	;
	v146 = int32(1)
	if v126&v146 != 0 {
		v160 = v58 + int32(base.Ui32(v126)>>(uint(v146)%32))
		goto L30
	} else {
		goto L58
	}
L49:
	;
	v136 = v131
	goto L51
L50:
	;
	v136 = int32(2)
	goto L51
L51:
	;
	if v133&int32(254) == int32(2) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v141 = v129
	goto L54
L53:
	;
	v141 = v136
	goto L54
L54:
	;
	if v133 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v144 = v129
	goto L57
L56:
	;
	v144 = v141
	goto L57
L57:
	;
	v160 = v58 + v144
	goto L30
L58:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v160 = v58 + int32(base.Ui32(v151)>>(uint(int32(2))%32))
	goto L30
L59:
	;
	v174 = (v160 + int32(1)) & int32(-2)
	goto L19
L60:
	;
	v174 = (v160 + int32(7)) & int32(-8)
	goto L19
L61:
	;
	v174 = (v160 + int32(3)) & int32(-4)
	goto L19
L62:
	;
	v180 = v175
	goto L64
L63:
	;
	v180 = v177
	goto L64
L64:
	;
	if v59 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v181 = v180
	goto L67
L66:
	;
	v181 = v54
	goto L67
L67:
	;
	if v59 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v184 = v179 + v59
	goto L70
L69:
	;
	v184 = int32(0)
	goto L70
L70:
	;
	v186 = v60 + int32(1)
	if v186 != v21 {
		v54 = v181
		v58 = v174
		v59 = v184
		v60 = v186
		goto L17
	} else {
		goto L71
	}
L71:
	;
	goto L18
}
func F_deconstruct_array_builtin(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(99)
	v15 = int32(1)
	switch l1 - int32(18) {
	case 0:
		v51 = v15
		v52 = v14
		v53 = v15
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 1, 2, 4, 6:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
			F_errmsg_internal(m, int32(668140), v12)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				F_errfinish(m, int32(493945), int32(3756), int32(274908))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 3:
		v51 = int32(2)
		v52 = int32(115)
		v53 = v15
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 5, 8:
		v51 = int32(4)
		v52 = int32(105)
		v53 = v15
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 7:
		v51 = int32(-1)
		v52 = int32(105)
		v53 = int32(0)
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	case 9:
		v51 = int32(6)
		v52 = int32(115)
		v53 = int32(0)
		F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return
		} else {
			m.G0 = v12 + int32(16)
			return
		}
	default:
		if l1 == int32(701) {
			v51 = int32(8)
			v52 = int32(100)
			v53 = int32(0)
			F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return
			} else {
				m.G0 = v12 + int32(16)
				return
			}
		} else {
			if l1 != int32(2275) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12))) = l1
					F_errmsg_internal(m, int32(668140), v12)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, int32(493945), int32(3756), int32(274908))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v51 = int32(-2)
				v52 = v14
				v53 = int32(0)
				F_deconstruct_array(m, l0, v51, v53, v52, l2, l3, l4)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					m.G0 = v12 + int32(16)
					return
				}
			}
		}
	}
}
func F_expand_array(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
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
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
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
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int64
	_ = v170
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v16 = F_AllocSetContextCreateInternal(m, l1, int32(26042), int32(0), int32(1024), int32(8388608))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v21 = F_MemoryContextAlloc(m, v16, int32(80))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v24 = int32(513)
			*(*uint16)(unsafe.Add(mBase, uint32(v21)+18)) = uint16(v24)
			v26 = int32(769)
			*(*uint16)(unsafe.Add(mBase, uint32(v21)+12)) = uint16(v26)
			*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = int32(1641504)
			*(*int32)(unsafe.Add(mBase, uint32(v21))) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v21)+14)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(689375833)
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v36 != int32(1) {
				v114 = l2
				v116 = int32(4489440)
				v117 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
				v120 = F_pg_detoast_datum_copy(m, l0)
				mBase = m.M
				v121 = m.ExcPending
				if v121 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v117
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
					v126 = v120 + int32(16)
					*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v126
					*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v124
					v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v126 + v129<<(uint(int32(2))%32)
					v134 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v134
					if v114 == int32(0) {
						F_get_typlenbyvalalign(m, v134, v21+int32(44), v21+int32(46), v21+int32(47))
						mBase = m.M
						v145 = m.ExcPending
						if v145 != 0 {
							return int32(0)
						} else {
							v170 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
							*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
							v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
							if v179 != 0 {
								v187 = v179
							} else {
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
								v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
							v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
							v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
							*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
							m.G0 = v10 + int32(48)
							return v21 + int32(12)
						}
					} else {
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
						if v134 == v146 {
							v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v148)
							v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
							*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)) = uint8(v150)
							v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
							*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)) = uint8(v152)
							v170 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
							*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
							*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
							v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
							if v179 != 0 {
								v187 = v179
							} else {
								v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
								v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
							v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
							v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
							*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
							m.G0 = v10 + int32(48)
							return v21 + int32(12)
						} else {
							F_get_typlenbyvalalign(m, v134, v21+int32(44), v21+int32(46), v21+int32(47))
							mBase = m.M
							v161 = m.ExcPending
							if v161 != 0 {
								return int32(0)
							} else {
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v114))) = v162
								v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)))
								*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)) = uint16(v164)
								v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)))
								*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)) = uint8(v166)
								v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
								*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)) = uint8(v168)
								v170 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
								*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
								if v179 != 0 {
									v187 = v179
								} else {
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
									v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
								v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
								v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
								m.G0 = v10 + int32(48)
								return v21 + int32(12)
							}
						}
					}
				}
			} else {
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v39&int32(254) != int32(2) {
					v114 = l2
					v116 = int32(4489440)
					v117 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
					v120 = F_pg_detoast_datum_copy(m, l0)
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v117
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
						v126 = v120 + int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v126
						*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v124
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v126 + v129<<(uint(int32(2))%32)
						v134 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v134
						if v114 == int32(0) {
							F_get_typlenbyvalalign(m, v134, v21+int32(44), v21+int32(46), v21+int32(47))
							mBase = m.M
							v145 = m.ExcPending
							if v145 != 0 {
								return int32(0)
							} else {
								v170 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
								*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
								if v179 != 0 {
									v187 = v179
								} else {
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
									v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
								v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
								v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
								m.G0 = v10 + int32(48)
								return v21 + int32(12)
							}
						} else {
							v146 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
							if v134 == v146 {
								v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v148)
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
								*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)) = uint8(v150)
								v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
								*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)) = uint8(v152)
								v170 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
								*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
								*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
								v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
								if v179 != 0 {
									v187 = v179
								} else {
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
									v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
								v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
								v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
								*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
								m.G0 = v10 + int32(48)
								return v21 + int32(12)
							} else {
								F_get_typlenbyvalalign(m, v134, v21+int32(44), v21+int32(46), v21+int32(47))
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int32(0)
								} else {
									v162 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
									*(*int32)(unsafe.Add(mBase, uint32(v114))) = v162
									v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)))
									*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)) = uint16(v164)
									v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)))
									*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)) = uint8(v166)
									v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
									*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)) = uint8(v168)
									v170 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
									*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
									*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
									if v179 != 0 {
										v187 = v179
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
										v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
									v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
									v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
									m.G0 = v10 + int32(48)
									return v21 + int32(12)
								}
							}
						}
					}
				} else {
					if l2 != 0 {
						v44 = l2
					} else {
						v44 = v10
					}
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v44))) = v46
					v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+44)))
					*(*uint16)(unsafe.Add(mBase, uint32(v44)+4)) = uint16(v48)
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+46)))
					*(*uint8)(unsafe.Add(mBase, uint32(v44)+6)) = uint8(v50)
					v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+47)))
					*(*uint8)(unsafe.Add(mBase, uint32(v44)+7)) = uint8(v52)
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+46)))
					if v54 != int32(1) {
						v114 = v44
						v116 = int32(4489440)
						v117 = *(*int32)(unsafe.Add(mBase, _consts[0]))
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
						v120 = F_pg_detoast_datum_copy(m, l0)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v117
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
							v126 = v120 + int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v126
							*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v124
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v126 + v129<<(uint(int32(2))%32)
							v134 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v134
							if v114 == int32(0) {
								F_get_typlenbyvalalign(m, v134, v21+int32(44), v21+int32(46), v21+int32(47))
								mBase = m.M
								v145 = m.ExcPending
								if v145 != 0 {
									return int32(0)
								} else {
									v170 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
									*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
									*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
									if v179 != 0 {
										v187 = v179
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
										v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
									v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
									v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
									m.G0 = v10 + int32(48)
									return v21 + int32(12)
								}
							} else {
								v146 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
								if v134 == v146 {
									v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
									*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v148)
									v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
									*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)) = uint8(v150)
									v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
									*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)) = uint8(v152)
									v170 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
									*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
									*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
									*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
									if v179 != 0 {
										v187 = v179
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
										v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
									}
									*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
									v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
									v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
									*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
									m.G0 = v10 + int32(48)
									return v21 + int32(12)
								} else {
									F_get_typlenbyvalalign(m, v134, v21+int32(44), v21+int32(46), v21+int32(47))
									mBase = m.M
									v161 = m.ExcPending
									if v161 != 0 {
										return int32(0)
									} else {
										v162 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
										*(*int32)(unsafe.Add(mBase, uint32(v114))) = v162
										v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)))
										*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)) = uint16(v164)
										v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)))
										*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)) = uint8(v166)
										v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
										*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)) = uint8(v168)
										v170 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
										*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
										*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
										v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
										if v179 != 0 {
											v187 = v179
										} else {
											v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
											v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
										v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
										v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
										m.G0 = v10 + int32(48)
										return v21 + int32(12)
									}
								}
							}
						}
					} else {
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
						if v57 == int32(0) {
							v114 = v44
							v116 = int32(4489440)
							v117 = *(*int32)(unsafe.Add(mBase, _consts[0]))
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v16
							v120 = F_pg_detoast_datum_copy(m, l0)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, _consts[0])) = v117
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
								v126 = v120 + int32(16)
								*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v126
								*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v124
								v129 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v126 + v129<<(uint(int32(2))%32)
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v134
								if v114 == int32(0) {
									F_get_typlenbyvalalign(m, v134, v21+int32(44), v21+int32(46), v21+int32(47))
									mBase = m.M
									v145 = m.ExcPending
									if v145 != 0 {
										return int32(0)
									} else {
										v170 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
										*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
										*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
										v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
										if v179 != 0 {
											v187 = v179
										} else {
											v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
											v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
										v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
										v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
										m.G0 = v10 + int32(48)
										return v21 + int32(12)
									}
								} else {
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
									if v134 == v146 {
										v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)))
										*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v148)
										v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)))
										*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)) = uint8(v150)
										v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)))
										*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)) = uint8(v152)
										v170 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
										*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
										*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
										*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
										v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
										if v179 != 0 {
											v187 = v179
										} else {
											v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
											v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
										}
										*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
										v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
										v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
										*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
										m.G0 = v10 + int32(48)
										return v21 + int32(12)
									} else {
										F_get_typlenbyvalalign(m, v134, v21+int32(44), v21+int32(46), v21+int32(47))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											v162 = *(*int32)(unsafe.Add(mBase, uint32(v21)+40))
											*(*int32)(unsafe.Add(mBase, uint32(v114))) = v162
											v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)))
											*(*uint16)(unsafe.Add(mBase, uint32(v114)+4)) = uint16(v164)
											v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)))
											*(*uint8)(unsafe.Add(mBase, uint32(v114)+6)) = uint8(v166)
											v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
											*(*uint8)(unsafe.Add(mBase, uint32(v114)+7)) = uint8(v168)
											v170 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v170
											*(*int32)(unsafe.Add(mBase, uint32(v21)+68)) = v120
											*(*int32)(unsafe.Add(mBase, uint32(v21-int32(-64)))) = int32(0)
											*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = v170
											v179 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
											if v179 != 0 {
												v187 = v179
											} else {
												v180 = *(*int32)(unsafe.Add(mBase, uint32(v120)+4))
												v187 = (v180<<(uint(int32(3))%32) + int32(23)) & int32(-8)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v187 + v120
											v190 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
											v199 = v120 + int32(base.Ui32(v190)>>(uint(int32(2))%32))
											*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
											m.G0 = v10 + int32(48)
											return v21 + int32(12)
										}
									}
								}
							}
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v45)+56))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v45)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
							v66 = F_MemoryContextAlloc(m, v63, v61<<(uint(int32(3))%32))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v66
								v70 = v61 << (uint(int32(2)) % 32)
								*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v66 + v70
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v45)+32))
								if v70 != 0 {
									v74 = F__emscripten_memcpy_bulkmem(m, v66, v73, v70)
									mBase = m.M
								} else {
								}
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v45)+36))
								if v70 != 0 {
									v78 = F__emscripten_memcpy_bulkmem(m, v76, v77, v70)
									mBase = m.M
								} else {
								}
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v45)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v80
								v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+44)))
								*(*uint16)(unsafe.Add(mBase, uint32(v21)+44)) = uint16(v82)
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+46)))
								*(*uint8)(unsafe.Add(mBase, uint32(v21)+46)) = uint8(v84)
								v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+47)))
								*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)) = uint8(v86)
								v89 = v60 << (uint(int32(2)) % 32)
								v90 = F_MemoryContextAlloc(m, v63, v89)
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v90
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v45)+48))
									if v89 != 0 {
										v94 = F__emscripten_memcpy_bulkmem(m, v90, v93, v89)
										mBase = m.M
									} else {
									}
									v96 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
									if v96 != 0 {
										v97 = F_MemoryContextAlloc(m, v63, v60)
										mBase = m.M
										v98 = m.ExcPending
										if v98 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v97
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
											if v60 != 0 {
												v101 = F__emscripten_memcpy_bulkmem(m, v97, v100, v60)
												mBase = m.M
											} else {
											}
											*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v60
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v45)+60))
											*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v107
											v109 = *(*int32)(unsafe.Add(mBase, uint32(v45)+64))
											*(*int64)(unsafe.Add(mBase, uint32(v21)+68)) = int64(0)
											*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v109
											v199 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
											m.G0 = v10 + int32(48)
											return v21 + int32(12)
										}
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v60
										v107 = *(*int32)(unsafe.Add(mBase, uint32(v45)+60))
										*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v107
										v109 = *(*int32)(unsafe.Add(mBase, uint32(v45)+64))
										*(*int64)(unsafe.Add(mBase, uint32(v21)+68)) = int64(0)
										*(*int32)(unsafe.Add(mBase, uint32(v21)+64)) = v109
										v199 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v21)+76)) = v199
										m.G0 = v10 + int32(48)
										return v21 + int32(12)
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
func F_get_array_element_start(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v8 <= v9 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		v16 = v11 + v8<<(uint(int32(2))%32) - int32(4)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17 + int32(1)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v22 = v21
	} else {
		v22 = v9
	}
	if v22 < v8 {
		return int32(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v25 = int32(1)
		v26 = v8 - v25
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v26))))
		if v28 != v25 {
			return int32(0)
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v31 == int32(0) {
				return int32(0)
			} else {
				v35 = v26 << (uint(int32(2)) % 32)
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v35+v36)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v35+v31)))
				if v38 != v40 {
					return int32(0)
				} else {
					if v8 < v22 {
						v44 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v8+v24))) = uint8(v44)
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
						v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
						if v51 != int32(1) {
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v61
							return int32(0)
						} else {
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
							if v54 != int32(1) {
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v61
								return int32(0)
							} else {
								v57 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)) = uint8(v57)
								return int32(0)
							}
						}
					}
				}
			}
		}
	}
}
func F_makeArrayResult(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(1)
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
	v16 = int32(4489440)
	v17 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = l1
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+26)))
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+27)))
	v32 = F_construct_md_array(m, v20, v21, base.B2i32(int32(0) < v11), v9+int32(12), v9+int32(8), v28, v29, v30, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[0])) = v17
		if v15 == int32(1) {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			F_MemoryContextDelete(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				m.G0 = v9 + int32(16)
				return v32
			}
		} else {
			m.G0 = v9 + int32(16)
			return v32
		}
	}
}
func F_makeArrayTypeName(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v13 = F_makeObjectName(m, int32(741336), l0, v3)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = int32(0)
	v19 = F_SearchSysCacheExists(m, int32(81), v13, l1, v17, v17)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = v13
	v25 = v3
	goto L7
L5:
	;
	v50 = v13
	goto L6
L6:
	;
	m.G0 = v8 + int32(80)
	return v50
L7:
	;
	F_pfree(m, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v50 = v41
	goto L6
L9:
	;
	v29 = v25 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v29
	v35 = F_pg_snprintf(m, v8+int32(16), int32(64), int32(487662), v8)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v41 = F_makeObjectName(m, int32(741336), l0, v8+int32(16))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v43 = int32(0)
	v45 = F_SearchSysCacheExists(m, int32(81), v41, l1, v43, v43)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	if v45 != 0 {
		v24 = v41
		v25 = v29
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L8
}
func F_parse_array_element(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v12 = m.T0[v8].(func(*base.Module, int32, int32) int32)(m, v9, base.B2i32(v6 == int32(11)))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			if v12 != 0 {
				v33 = v12
				return v33
			} else {
				switch v6 - int32(3) {
				case 0:
					v19 = F_parse_object(m, l0, l1)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int32(0)
					} else {
						v25 = v19
						if v25 != 0 {
							v33 = v25
							return v33
						} else {
							if v7 != 0 {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									if v29 != 0 {
										v33 = v29
									} else {
										v33 = int32(0)
									}
									return v33
								}
							} else {
								v33 = int32(0)
								return v33
							}
						}
					}
				default:
					v23 = F_parse_scalar(m, l0, l1)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = v23
						if v25 != 0 {
							v33 = v25
							return v33
						} else {
							if v7 != 0 {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									if v29 != 0 {
										v33 = v29
									} else {
										v33 = int32(0)
									}
									return v33
								}
							} else {
								v33 = int32(0)
								return v33
							}
						}
					}
				case 2:
					v21 = F_parse_array(m, l0, l1)
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						v25 = v21
						if v25 != 0 {
							v33 = v25
							return v33
						} else {
							if v7 != 0 {
								v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return int32(0)
								} else {
									if v29 != 0 {
										v33 = v29
									} else {
										v33 = int32(0)
									}
									return v33
								}
							} else {
								v33 = int32(0)
								return v33
							}
						}
					}
				}
			}
		}
	} else {
		switch v6 - int32(3) {
		case 0:
			v19 = F_parse_object(m, l0, l1)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v25 = v19
				if v25 != 0 {
					v33 = v25
					return v33
				} else {
					if v7 != 0 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = v29
							} else {
								v33 = int32(0)
							}
							return v33
						}
					} else {
						v33 = int32(0)
						return v33
					}
				}
			}
		default:
			v23 = F_parse_scalar(m, l0, l1)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = v23
				if v25 != 0 {
					v33 = v25
					return v33
				} else {
					if v7 != 0 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = v29
							} else {
								v33 = int32(0)
							}
							return v33
						}
					} else {
						v33 = int32(0)
						return v33
					}
				}
			}
		case 2:
			v21 = F_parse_array(m, l0, l1)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v25 = v21
				if v25 != 0 {
					v33 = v25
					return v33
				} else {
					if v7 != 0 {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v29 = m.T0[v7].(func(*base.Module, int32, int32) int32)(m, v26, base.B2i32(v6 == int32(11)))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							if v29 != 0 {
								v33 = v29
							} else {
								v33 = int32(0)
							}
							return v33
						}
					} else {
						v33 = int32(0)
						return v33
					}
				}
			}
		}
	}
}
