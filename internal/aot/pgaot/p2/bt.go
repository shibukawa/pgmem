package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetBTPageStatistics(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int64
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	if l1 < int32(0) {
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_GetBTPageStatistics[0]))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v18+(l1^int32(-1))<<(uint(int32(2))%32))))
		v32 = v24
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_GetBTPageStatistics[1]))
		v32 = v26 + l1<<(uint(int32(13))%32) + int32(-8192)
	}
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = l0
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v34 - int32(24)
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+19)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v41 << (uint(int32(8)) % 32)
	v45 = v34 + v32
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
	if v46&int32(4) != 0 {
		if v46&int32(257) == int32(256) {
			v55 = int32(68)
		} else {
			v55 = int32(100)
		}
		*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v55)
		if v46&int32(256) != 0 {
			v59 = *(*int64)(unsafe.Add(mBase, uint32(v32)+24))
			v60 = int32(0)
			v63 = F_errstart(m, int32(13), v60)
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return
			} else {
				if v63 == int32(0) {
					v122 = v60
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
					v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
					v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
					v134 = int32(0)
					v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
					if v136 != 0 {
						v140 = v134
						v141 = int32(1)
						for {
							v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
							v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
							v160 = int32(_a_F_GetBTPageStatistics_2)
							if v153&v160 != v160 {
								v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
							}
							v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
							if v141 != v136 {
								v140 = v172
								v141 = v141 + int32(1)
								continue
							} else {
								break
							}
							break
						}
						v176 = v172
					} else {
						v176 = v134
					}
					v186 = int32(4)
					v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
					v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
					v189 = v187 - v188
					if v189 <= v186 {
						v192 = v186
					} else {
						v192 = v189
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					v198 = v196 + v197
					if v198 != 0 {
						v199 = base.I32_div_u_s(v176, v198)
						v201 = v199
					} else {
						v201 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
					m.G0 = v13 + int32(32)
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = l0
					*(*uint32)(unsafe.Add(mBase, uint32(v13)+24)) = uint32(v59)
					v70 = int64(base.Ui64(v59) >> (uint(int64(32)) % 64))
					*(*uint32)(unsafe.Add(mBase, uint32(v13)+20)) = uint32(v70)
					F_errmsg_internal(m, int32(_a_F_GetBTPageStatistics_4), v13+int32(16))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_GetBTPageStatistics_5), int32(147), int32(_a_F_GetBTPageStatistics_6))
						mBase = m.M
						v81 = m.ExcPending
						if v81 != 0 {
							return
						} else {
							v122 = v60
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
							v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
							v134 = int32(0)
							v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
							if v136 != 0 {
								v140 = v134
								v141 = int32(1)
								for {
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
									v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
									v160 = int32(_a_F_GetBTPageStatistics_2)
									if v153&v160 != v160 {
										v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
									} else {
										v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
									}
									v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
									if v141 != v136 {
										v140 = v172
										v141 = v141 + int32(1)
										continue
									} else {
										break
									}
									break
								}
								v176 = v172
							} else {
								v176 = v134
							}
							v186 = int32(4)
							v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
							v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
							v189 = v187 - v188
							if v189 <= v186 {
								v192 = v186
							} else {
								v192 = v189
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
							v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v198 = v196 + v197
							if v198 != 0 {
								v199 = base.I32_div_u_s(v176, v198)
								v201 = v199
							} else {
								v201 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
							m.G0 = v13 + int32(32)
							return
						}
					}
				}
			}
		} else {
			v82 = int32(0)
			v85 = F_errstart(m, int32(13), v82)
			mBase = m.M
			v86 = m.ExcPending
			if v86 != 0 {
				return
			} else {
				if v85 == int32(0) {
					v122 = v82
					v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
					v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
					v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
					v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
					v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
					*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
					v134 = int32(0)
					v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
					if v136 != 0 {
						v140 = v134
						v141 = int32(1)
						for {
							v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
							v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
							v160 = int32(_a_F_GetBTPageStatistics_2)
							if v153&v160 != v160 {
								v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
								*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
							}
							v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
							if v141 != v136 {
								v140 = v172
								v141 = v141 + int32(1)
								continue
							} else {
								break
							}
							break
						}
						v176 = v172
					} else {
						v176 = v134
					}
					v186 = int32(4)
					v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
					v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
					v189 = v187 - v188
					if v189 <= v186 {
						v192 = v186
					} else {
						v192 = v189
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
					v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					v198 = v196 + v197
					if v198 != 0 {
						v199 = base.I32_div_u_s(v176, v198)
						v201 = v199
					} else {
						v201 = int32(0)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
					m.G0 = v13 + int32(32)
					return
				} else {
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v89
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = l0
					F_errmsg_internal(m, int32(_a_F_GetBTPageStatistics_7), v13)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_GetBTPageStatistics_5), int32(151), int32(_a_F_GetBTPageStatistics_6))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return
						} else {
							v122 = v82
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
							v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
							v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
							v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
							*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
							v134 = int32(0)
							v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
							if v136 != 0 {
								v140 = v134
								v141 = int32(1)
								for {
									v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
									v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
									v160 = int32(_a_F_GetBTPageStatistics_2)
									if v153&v160 != v160 {
										v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
									} else {
										v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
									}
									v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
									if v141 != v136 {
										v140 = v172
										v141 = v141 + int32(1)
										continue
									} else {
										break
									}
									break
								}
								v176 = v172
							} else {
								v176 = v134
							}
							v186 = int32(4)
							v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
							v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
							v189 = v187 - v188
							if v189 <= v186 {
								v192 = v186
							} else {
								v192 = v189
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
							v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
							v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
							v198 = v196 + v197
							if v198 != 0 {
								v199 = base.I32_div_u_s(v176, v198)
								v201 = v199
							} else {
								v201 = int32(0)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
							m.G0 = v13 + int32(32)
							return
						}
					}
				}
			}
		}
	} else {
		if base.Ui32(int32(25)) <= base.Ui32(v33) {
			v107 = int32(base.Ui32(v33+int32(_a_F_GetBTPageStatistics_8)) >> (uint(int32(2)) % 32))
		} else {
			v107 = int32(0)
		}
		if v46&int32(16) != 0 {
			v110 = int32(101)
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v110)
			v122 = v107
		} else {
			if v46&int32(1) != 0 {
				v114 = int32(108)
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v114)
				v122 = v107
			} else {
				if v46&int32(2) != 0 {
					v118 = int32(114)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v118)
					v122 = v107
				} else {
					v120 = int32(105)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+28)) = uint8(v120)
					v122 = v107
				}
			}
		}
		v124 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+32)) = v124
		v126 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v126
		v128 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+40)) = v128
		v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+12)))
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)) = uint16(v130)
		v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+14)))
		*(*uint16)(unsafe.Add(mBase, uint32(l2)+46)) = uint16(v132)
		v134 = int32(0)
		v136 = v122 & int32(_a_F_GetBTPageStatistics_0)
		if v136 != 0 {
			v140 = v134
			v141 = int32(1)
			for {
				v153 = *(*int32)(unsafe.Add(mBase, uint32(v32+int32(20)+v141<<(uint(int32(2))%32))))
				v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32+v153&int32(_a_F_GetBTPageStatistics_1))+6)))
				v160 = int32(_a_F_GetBTPageStatistics_2)
				if v153&v160 != v160 {
					v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v164 + int32(1)
				} else {
					v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v168 + int32(1)
				}
				v172 = v140 + v157&int32(_a_F_GetBTPageStatistics_3)
				if v141 != v136 {
					v140 = v172
					v141 = v141 + int32(1)
					continue
				} else {
					break
				}
				break
			}
			v176 = v172
		} else {
			v176 = v134
		}
		v186 = int32(4)
		v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+14)))
		v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v32)+12)))
		v189 = v187 - v188
		if v189 <= v186 {
			v192 = v186
		} else {
			v192 = v189
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v192 - int32(4)
		v196 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		v197 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		v198 = v196 + v197
		if v198 != 0 {
			v199 = base.I32_div_u_s(v176, v198)
			v201 = v199
		} else {
			v201 = int32(0)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l2)+24)) = v201
		m.G0 = v13 + int32(32)
		return
	}
}
func F__bt_binsrch_array_skey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v269 int32
	_ = v269
	v9 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v16 = v14 - int32(1)
	if l1 == v9 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(0)
	return v269
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v248
	return v254
L3:
	;
	if v130 < v132 {
		goto L52
	} else {
		goto L53
	}
L4:
	;
	v128 = int32(0)
	v130 = v9
	v131 = int32(-1)
	v132 = v16
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	if l2 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v123 = int32(0)
	v125 = v20 - int32(2)
	if v125 < v123 {
		v248 = v77
		v254 = v123
		goto L2
	} else {
		goto L51
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v116
	return v119
L9:
	;
	v24 = v20 + int32(1)
	if v16 < v24 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v77 = int32(-1)
	v79 = v20 - int32(1)
	if v79 < int32(0) {
		v248 = v77
		v254 = v9
		goto L2
	} else {
		goto L34
	}
L12:
	;
	if v71 <= v16 {
		v128 = v69
		v130 = v71
		v131 = v72
		v132 = v16
		goto L3
	} else {
		goto L33
	}
L13:
	;
	v69 = int32(0)
	v71 = v24
	v72 = int32(-1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v30 = v28 & int32(1)
	if l4 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v69 = v65
	v71 = v20 + int32(2)
	v72 = v24
	goto L12
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v24<<(uint(int32(2))%32))))
	v50 = F_FunctionCall2Coll(m, l0, v44, l3, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(-1)
	return v24
L19:
	;
	if v30 != 0 {
		v269 = v24
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v30 == int32(0) {
		goto L17
	} else {
		goto L24
	}
L22:
	;
	if v28&int32(33554432) != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v65 = int32(1)
	goto L16
L24:
	;
	if v28&int32(33554432) == int32(0) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v65 = int32(1)
	goto L16
L26:
	;
	return int32(0)
L27:
	;
	v54 = int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v55&v54 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v50 < int32(0) {
		v65 = v54
		goto L16
	} else {
		goto L31
	}
L29:
	;
	v62 = v50
	goto L30
L30:
	;
	if v62 <= int32(0) {
		v116 = v62
		v119 = v24
		goto L8
	} else {
		goto L32
	}
L31:
	;
	v62 = int32(0) - v50
	goto L30
L32:
	;
	v65 = v62
	goto L16
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(1)
	return v16
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v84 = v82 & int32(1)
	if l4 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	if v113 < int32(0) {
		v122 = v113
		goto L7
	} else {
		goto L50
	}
L36:
	;
	v113 = int32(0) - v98
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(1)
	return v79
L38:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93+v79<<(uint(int32(2))%32))))
	v98 = F_FunctionCall2Coll(m, l0, v92, l3, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L26
	} else {
		goto L47
	}
L39:
	;
	v122 = int32(-1)
	goto L7
L40:
	;
	if v84 != 0 {
		v269 = v79
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v84 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L43:
	;
	if v82&int32(33554432) != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	if v82&int32(33554432) != 0 {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L39
L47:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v100&int32(1) == int32(0) {
		v113 = v98
		goto L35
	} else {
		goto L48
	}
L48:
	;
	if int32(0) <= v98 {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	goto L37
L50:
	;
	v116 = v113
	v119 = v79
	goto L8
L51:
	;
	v128 = v122
	v130 = v123
	v131 = v79
	v132 = v125
	goto L3
L52:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v138 = v135
	v144 = v130
	v146 = v132
	goto L55
L53:
	;
	v193 = v128
	v200 = v130
	v201 = v131
	goto L54
L54:
	;
	if v200 == v201 {
		goto L85
	} else {
		goto L86
	}
L55:
	;
	v149 = v138 & int32(1)
	v152 = base.I32_div_s(v146-v144, int32(2))
	v153 = v152 + v144
	if l4 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v193 = v183
	v200 = v190
	v201 = v153
	goto L54
L57:
	;
	v186 = base.B2i32(int32(0) < v183)
	if int32(0) < v183 {
		goto L78
	} else {
		goto L79
	}
L58:
	;
	if v149 != 0 {
		v269 = v153
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v149 != 0 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	if v138&int32(33554432) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v158 = int32(-1)
	goto L64
L63:
	;
	v158 = int32(1)
	goto L64
L64:
	;
	v183 = v158
	v184 = v138
	goto L57
L65:
	;
	if v138&int32(33554432) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165+v153<<(uint(int32(2))%32))))
	v170 = F_FunctionCall2Coll(m, l0, v164, l3, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L26
	} else {
		goto L71
	}
L68:
	;
	v163 = int32(1)
	goto L70
L69:
	;
	v163 = int32(-1)
	goto L70
L70:
	;
	v183 = v163
	v184 = v138
	goto L57
L71:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v172&int32(16777216) != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v183 = int32(1)
	v184 = v172
	goto L57
L73:
	;
	if v170 < int32(0) {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	v179 = v170
	goto L75
L75:
	;
	if v179 == int32(0) {
		v269 = v153
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v179 = int32(0) - v170
	goto L75
L77:
	;
	v183 = v179
	v184 = v172
	goto L57
L78:
	;
	v187 = v146
	goto L80
L79:
	;
	v187 = v153
	goto L80
L80:
	;
	if int32(0) < v183 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v190 = v153 + int32(1)
	goto L83
L82:
	;
	v190 = v144
	goto L83
L83:
	;
	if v190 < v187 {
		v138 = v184
		v144 = v190
		v146 = v187
		goto L55
	} else {
		goto L84
	}
L84:
	;
	goto L56
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v193
	return v200
L86:
	;
	goto L87
L87:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v209 = v207 & int32(1)
	if l4 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v209 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v209 != 0 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(0)
	return v200
L92:
	;
	goto L93
L93:
	;
	if v207&int32(33554432) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v217 = int32(-1)
	goto L96
L95:
	;
	v217 = int32(1)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v217
	return v200
L97:
	;
	if v207&int32(33554432) != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228+v200<<(uint(int32(2))%32))))
	v233 = F_FunctionCall2Coll(m, l0, v227, l3, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L26
	} else {
		goto L103
	}
L100:
	;
	v224 = int32(1)
	goto L102
L101:
	;
	v224 = int32(-1)
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v224
	return v200
L103:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v235&int32(1) == int32(0) {
		v248 = v233
		v254 = v200
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v241 = int32(0)
	if v233 < v241 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v245 = int32(1)
	goto L107
L106:
	;
	v245 = v241 - v233
	goto L107
L107:
	;
	v248 = v245
	v254 = v200
	goto L2
}
func F__bt_check_third_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)))
	v18 = (v12&int32(_a_F__bt_check_third_page_0) + int32(7)) & int32(_a_F__bt_check_third_page_1)
	if base.B2i32(base.Ui32(v18) < base.Ui32(int32(2705)))|base.B2i32(l2|base.B2i32(base.Ui32(int32(2712)) < base.Ui32(v18)) == v6) == v6 {
		v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+16)))
		v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v29)+12)))
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			if v31&int32(1) == int32(0) {
				v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v126 + int32(4)
				F_errmsg_internal(m, int32(_a_F__bt_check_third_page_2), v10)
				mBase = m.M
				v133 = m.ExcPending
				if v133 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F__bt_check_third_page_3), int32(_a_F__bt_check_third_page_4), int32(_a_F__bt_check_third_page_5))
					mBase = m.M
					v138 = m.ExcPending
					if v138 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					if l2 != 0 {
						v46 = int32(2704)
					} else {
						v46 = int32(2712)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v46
					if l2 != 0 {
						v50 = int32(4)
					} else {
						v50 = int32(3)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v50
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v43 + int32(4)
					F_errmsg(m, int32(_a_F__bt_check_third_page_6), v10+int32(32))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)))
						if v63&int32(_a_F__bt_check_third_page_7) == int32(0) {
							v91 = l4
						} else {
							v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
							if v68&int32(_a_F__bt_check_third_page_7) == int32(0) {
								v73 = int32(0)
								if v68&int32(_a_F__bt_check_third_page_8) == v73 {
									v89 = v73
									v91 = v89
								} else {
									v91 = l4 + v63&int32(_a_F__bt_check_third_page_0) - int32(6)
								}
							} else {
								v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+2)))
								v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
								v89 = v83 + (l4 + v84<<(uint(int32(16))%32))
								v91 = v89
							}
						}
						v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+2)))
						v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91))))
						v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
						v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v95
						*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v94 + int32(4)
						v100 = int32(16)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v92 | v93<<(uint(v100)%32)
						F_errdetail(m, int32(_a_F__bt_check_third_page_9), v10+v100)
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F__bt_check_third_page_10), int32(0))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return
							} else {
								v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
								F_errtableconstraint(m, l1, v113+int32(4))
								mBase = m.M
								v117 = m.ExcPending
								if v117 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F__bt_check_third_page_3), int32(_a_F__bt_check_third_page_11), int32(_a_F__bt_check_third_page_5))
									mBase = m.M
									v122 = m.ExcPending
									if v122 != 0 {
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
				}
			}
		}
	} else {
		m.G0 = v10 + int32(48)
		return
	}
}
func F__bt_delitems_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	return v3 - v4
}
func F__bt_delitems_delete_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v328 int32
	_ = v328
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v418 int32
	_ = v418
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v525 int32
	_ = v525
	var v541 int32
	_ = v541
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v646 int32
	_ = v646
	var v649 int32
	_ = v649
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v674 int32
	_ = v674
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v769 int32
	_ = v769
	var v772 int32
	_ = v772
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v808 int32
	_ = v808
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v830 int32
	_ = v830
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v845 int64
	_ = v845
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v931 int32
	_ = v931
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(3296)
	m.G0 = v27
	if l1 < v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+188))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+76))
	v49 = m.T0[v48].(func(*base.Module, int32, int32) int32)(m, l2, l3)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[0]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+(l1^int32(-1))<<(uint(int32(2))%32))))
	v46 = v38
	goto L1
L3:
	;
	goto L4
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[1]))
	v46 = v40 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	return
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[2]))
	if v52 < int32(2) {
		v73 = v5
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[2]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	F_pg_qsort(m, v76, v77, int32(8), int32(209))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+118)))
	if v56 != int32(112) {
		v73 = v5
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	goto L10
L10:
	;
	if base.Ui32(v60) < base.Ui32(int32(_a_F__bt_delitems_delete_check_0)) {
		v73 = int32(1)
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
	if v64 == v63 {
		v73 = v63
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+119)))
	switch v68 - int32(109) {
	case 0, 5:
		goto L13
	default:
		v73 = v63
		goto L7
	}
L13:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+104)))
	v73 = v71
	goto L7
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v82 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	m.G0 = v27 + int32(3296)
	return
L16:
	;
	if int32(0) < v82 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v96 = int32(0)
	v100 = v5
	v102 = v5
	v103 = v5
	goto L20
L18:
	;
	v401 = v5
	v403 = v5
	goto L19
L19:
	;
	if l1 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L20:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115+v103<<(uint(int32(3))%32))+6)))
	v122 = v114 + v119*int32(6)
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	if v123 == v96&int32(_a_F__bt_delitems_delete_check_1) {
		v369 = v96
		v373 = v100
		v375 = v102
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v401 = v373
	v403 = v375
	goto L19
L22:
	;
	v388 = v103 + int32(1)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v388 < v389 {
		v96 = v369
		v100 = v373
		v102 = v375
		v103 = v388
		goto L20
	} else {
		goto L62
	}
L23:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v46+int32(20)+v123<<(uint(int32(2))%32))))
	v133 = v46 + v130&int32(_a_F__bt_delitems_delete_check_2)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+7)))
	if v134&int32(32) != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v153 = v137 & int32(4095)
	if v153 == int32(0) {
		v349 = v100
		v351 = v102
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+4)))
	if v137&int32(_a_F__bt_delitems_delete_check_3) != 0 {
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
	if v141 != int32(1) {
		v369 = v96
		v373 = v100
		v375 = v102
		goto L22
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v146 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(1648)+v102<<(uint(v146)%32)))) = uint16(v123)
	v369 = v96
	v373 = v100
	v375 = v102 + v146
	goto L22
L30:
	;
	v369 = v123
	v373 = v349
	v375 = v351
	goto L22
L31:
	;
	v160 = int32(0)
	v164 = v103
	v170 = v160
	v171 = v160
	goto L32
L32:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v186 <= v164 {
		v292 = v164
		v298 = v170
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v298 == int32(0) {
		v349 = v100
		v351 = v102
		goto L30
	} else {
		goto L57
	}
L34:
	;
	v315 = v171 + int32(1)
	if v315 != v153 {
		v164 = v292
		v170 = v298
		v171 = v315
		goto L32
	} else {
		goto L56
	}
L35:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+2)))
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133))))
	v196 = v188 + (v133 + v189<<(uint(int32(16))%32)) + v171*int32(6)
	v200 = v164
	v204 = int32(-1)
	v205 = v186
	goto L36
L36:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v226 = v223 + v200<<(uint(int32(3))%32)
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(v226)+6)))
	v230 = v222 + v227*int32(6)
	v231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230))))
	if v231 != v123 {
		v268 = v200
		v269 = v204
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v269 != 0 {
		v292 = v268
		v298 = v170
		goto L34
	} else {
		goto L50
	}
L38:
	;
	goto L37
L39:
	;
	v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+2)))
	if v233 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+2)))
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226))))
	v241 = int32(16)
	v243 = v239 | v240<<(uint(v241)%32)
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+2)))
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196))))
	v248 = v244 | v245<<(uint(v241)%32)
	if base.Ui32(v243) < base.Ui32(v248) {
		v259 = int32(-1)
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v263 = v204
	v264 = v205
	goto L42
L42:
	;
	v266 = v200 + int32(1)
	if v266 < v264 {
		v200 = v266
		v204 = v263
		v205 = v264
		goto L36
	} else {
		goto L49
	}
L43:
	;
	if int32(0) <= v259 {
		v268 = v200
		v269 = v259
		goto L38
	} else {
		goto L48
	}
L44:
	;
	goto L43
L45:
	;
	if base.Ui32(v248) < base.Ui32(v243) {
		v259 = int32(1)
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v226)+4)))
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+4)))
	if base.Ui32(v253) < base.Ui32(v254) {
		v259 = int32(-1)
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v259 = base.B2i32(base.Ui32(v254) < base.Ui32(v253))
	goto L44
L48:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v263 = v259
	v264 = v262
	goto L42
L49:
	;
	v268 = v266
	v269 = v263
	goto L38
L50:
	;
	if v170 != 0 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v281 = int32(1)
	v282 = v279 + v281
	*(*uint16)(unsafe.Add(mBase, uint32(v280)+6)) = uint16(v282)
	*(*uint16)(unsafe.Add(mBase, uint32(v280+v279&int32(_a_F__bt_delitems_delete_check_1)<<(uint(v281)%32))+8)) = uint16(v171)
	v292 = v268
	v298 = v280
	goto L34
L52:
	;
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+6)))
	v279 = v271
	v280 = v170
	goto L51
L53:
	;
	goto L54
L54:
	;
	v273 = F_palloc(m, v153<<(uint(int32(1))%32)+int32(8))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v275 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v273)+6)) = uint16(v275)
	*(*uint16)(unsafe.Add(mBase, uint32(v273)+4)) = uint16(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v273))) = v133
	v279 = int32(0)
	v280 = v273
	goto L51
L56:
	;
	goto L33
L57:
	;
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298)+6)))
	if v153 == v319 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(1648)+v102<<(uint(int32(1))%32)))) = uint16(v123)
	F_pfree(m, v298)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v100<<(uint(int32(2))%32)))) = v298
	v349 = v100 + int32(1)
	v351 = v102
	goto L30
L61:
	;
	v349 = v100
	v351 = v102 + int32(1)
	goto L30
L62:
	;
	goto L21
L63:
	;
	v433 = int32(0)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435)+118)))
	if v436 != int32(112) {
		v449 = v433
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v418 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[0]))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v418+(l1^int32(-1))<<(uint(int32(2))%32))))
	v432 = v424
	goto L63
L65:
	;
	goto L66
L66:
	;
	v426 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[1]))
	v432 = v426 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L63
L67:
	;
	if int32(0) < v401 {
		goto L72
	} else {
		goto L73
	}
L68:
	;
	v441 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[2]))
	if int32(0) < v441 {
		v449 = int32(1)
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v445 != 0 {
		v449 = int32(0)
		goto L67
	} else {
		goto L70
	}
L70:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v449 = base.B2i32(v446 == int32(0))
	goto L67
L71:
	;
	if int32(0) < v403 {
		goto L112
	} else {
		goto L113
	}
L72:
	;
	v455 = int32(0)
	v459 = v433
	goto L75
L73:
	;
	goto L74
L74:
	;
	v757 = int32(0)
	v758 = int32(_a_F__bt_delitems_delete_check_4)
	v760 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[3]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[3])) = v760 + int32(1)
	v769 = v757
	v772 = v757
	goto L71
L75:
	;
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v455<<(uint(int32(2))%32))))
	F__bt_update_posting(m, v482)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L5
	} else {
		goto L77
	}
L76:
	;
	v501 = int32(0)
	if v449 != 0 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+6)))
	v488 = int32(1)
	v491 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v482)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(2480)+v455<<(uint(v488)%32)))) = uint16(v491)
	v497 = v459 + v485<<(uint(v488)%32) + int32(2)
	v499 = v455 + v488
	if v499 != v401 {
		v455 = v499
		v459 = v497
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v504 = int32(0)
	v505 = F_palloc(m, v497)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L5
	} else {
		goto L82
	}
L80:
	;
	v646 = v501
	v649 = v501
	goto L81
L81:
	;
	v666 = int32(_a_F__bt_delitems_delete_check_4)
	v668 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[3]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[3])) = v668 + int32(1)
	v674 = v501
	goto L98
L82:
	;
	if v401 != int32(1) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v646 = v497
	v649 = v505
	goto L81
L84:
	;
	v517 = v504
	v518 = v501
	v525 = int32(0)
	goto L87
L85:
	;
	v579 = v504
	v580 = v501
	goto L86
L86:
	;
	v600 = v580 + v505
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v579<<(uint(int32(2))%32))))
	v607 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v606)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v600))) = uint16(v607)
	v610 = v607 << (uint(int32(1)) % 32)
	if v610 == int32(0) {
		goto L83
	} else {
		goto L97
	}
L87:
	;
	v541 = int32(2)
	v543 = v27 + int32(16) + v517<<(uint(v541)%32)
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)))
	v545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v544)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v518+v505))) = uint16(v545)
	v548 = v518 + v541
	v550 = v545 << (uint(int32(1)) % 32)
	if v550 != 0 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	if v401&int32(1) == int32(0) {
		goto L83
	} else {
		goto L96
	}
L89:
	;
	base.MemoryCopy(m, v548+v505, v544+int32(8), v550)
	goto L91
L90:
	;
	goto L91
L91:
	;
	v555 = v548 + v550
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v543)+4))
	v558 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v557)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v505+v555))) = uint16(v558)
	v561 = v555 + int32(2)
	v563 = v558 << (uint(int32(1)) % 32)
	if v563 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	base.MemoryCopy(m, v561+v505, v557+int32(8), v563)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v568 = v561 + v563
	v569 = int32(2)
	v570 = v517 + v569
	v572 = v525 + v569
	if v572 != v401&int32(2147483646) {
		v517 = v570
		v518 = v568
		v525 = v572
		goto L87
	} else {
		goto L95
	}
L95:
	;
	goto L88
L96:
	;
	v579 = v570
	v580 = v568
	goto L86
L97:
	;
	base.MemoryCopy(m, v600+int32(2), v606+int32(8), v610)
	goto L83
L98:
	;
	v701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(2480)+v674<<(uint(int32(1))%32)))))
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v674<<(uint(int32(2))%32))))
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v707)))
	v709 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v708)+6)))
	v716 = F_PageIndexTupleOverwrite(m, v432, v701, v708, (v709&int32(_a_F__bt_delitems_delete_check_5)+int32(7))&int32(_a_F__bt_delitems_delete_check_6))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L5
	} else {
		goto L100
	}
L99:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L5
	} else {
		goto L105
	}
L100:
	;
	if v716 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v719 = v674 + int32(1)
	if v719 != v401 {
		v674 = v719
		goto L98
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	goto L99
L104:
	;
	v769 = v646
	v772 = v649
	goto L71
L105:
	;
	if l1 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v743
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v744 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_delitems_delete_check_7), v27)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L5
	} else {
		goto L110
	}
L107:
	;
	v728 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[4]))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v728+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v743 = v734
	goto L106
L108:
	;
	goto L109
L109:
	;
	v736 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[5]))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v736+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v743 = v742
	goto L106
L110:
	;
	F_errfinish(m, int32(_a_F__bt_delitems_delete_check_8), int32(1320), int32(_a_F__bt_delitems_delete_check_9))
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L5
	} else {
		goto L111
	}
L111:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L112:
	;
	F_PageIndexMultiDelete(m, v432, v27+int32(1648), v403)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L5
	} else {
		goto L115
	}
L113:
	;
	goto L114
L114:
	;
	v795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v432)+16)))
	v796 = v432 + v795
	v797 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v796)+12)))
	v799 = v797 & int32(_a_F__bt_delitems_delete_check_10)
	*(*uint16)(unsafe.Add(mBase, uint32(v796)+12)) = uint16(v799)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L5
	} else {
		goto L116
	}
L115:
	;
	goto L114
L116:
	;
	if v449 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2476)) = uint8(v73)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+2474)) = uint16(v401)
	v805 = int32(0)
	if v805 < v75 {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	goto L119
L119:
	;
	v851 = int32(_a_F__bt_delitems_delete_check_4)
	v853 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[3]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delitems_delete_check[3])) = v853 - int32(1)
	if v772 != 0 {
		goto L136
	} else {
		goto L137
	}
L120:
	;
	v808 = v49
	goto L122
L121:
	;
	v808 = v805
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+2468)) = v808
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+2472)) = uint16(v403)
	F_XLogBeginInsert(m)
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L123
	}
L123:
	;
	F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
	mBase = m.M
	v816 = m.ExcPending
	if v816 != 0 {
		goto L5
	} else {
		goto L124
	}
L124:
	;
	F_XLogRegisterData(m, v27+int32(2468), int32(9))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L5
	} else {
		goto L125
	}
L125:
	;
	if int32(0) < v403 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_XLogRegisterBufData(m, int32(0), v27+int32(1648), v403<<(uint(int32(1))%32))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L5
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	if int32(0) < v401 {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	goto L128
L130:
	;
	F_XLogRegisterBufData(m, int32(0), v27+int32(2480), v401<<(uint(int32(1))%32))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L5
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v845 = F_XLogInsert(m, int32(11), int32(112))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	F_XLogRegisterBufData(m, int32(0), v772, v769)
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L5
	} else {
		goto L134
	}
L134:
	;
	goto L132
L135:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v432))) = base.I64_rotr(v845, int64(32))
	goto L119
L136:
	;
	F_pfree(m, v772)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L5
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	if v401 <= int32(0) {
		goto L15
	} else {
		goto L140
	}
L139:
	;
	goto L138
L140:
	;
	v863 = int32(0)
	goto L141
L141:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v863<<(uint(int32(2))%32))))
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v890)))
	F_pfree(m, v891)
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L5
	} else {
		goto L143
	}
L142:
	;
	v900 = int32(0)
	goto L145
L143:
	;
	v895 = v863 + int32(1)
	if v895 != v401 {
		v863 = v895
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v900<<(uint(int32(2))%32))))
	F_pfree(m, v927)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L5
	} else {
		goto L147
	}
L146:
	;
	goto L15
L147:
	;
	v931 = v900 + int32(1)
	if v931 != v401 {
		v900 = v931
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
}
func F__bt_killitems(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v267 int32
	_ = v267
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v300 int32
	_ = v300
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v393 int32
	_ = v393
	v2 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v2
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+40)))
	if v24 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__bt_relbuf(m, v375)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L6
	} else {
		goto L68
	}
L2:
	;
	if v39 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+56))
	F_LockBuffer(m, v27, int32(1))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+60))
	v33 = F__bt_getbuf(m, v19, v31, int32(1))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	v39 = v27
	goto L2
L8:
	;
	v35 = F_BufferGetLSNAtomic(m, v33)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v20)+72))
	if v35 != v37 {
		v375 = v33
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v39 = v33
	goto L2
L11:
	;
	if v21 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F__bt_killitems[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v39^int32(-1))<<(uint(int32(2))%32))))
	v57 = v49
	goto L11
L13:
	;
	goto L14
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F__bt_killitems[1]))
	v57 = v51 + v39<<(uint(int32(13))%32) + int32(-8192)
	goto L11
L15:
	;
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+40)))
	if v371 != 0 {
		v375 = v39
		goto L1
	} else {
		goto L66
	}
L16:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
	v63 = v57 + v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	if v64 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = int32(2)
	goto L19
L18:
	;
	v65 = int32(1)
	goto L19
L19:
	;
	v69 = v20 + int32(104)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v70) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v78 = int32(base.Ui32(v70+int32(_a_F__bt_killitems_0)) >> (uint(int32(2)) % 32))
	goto L22
L21:
	;
	v78 = int32(0)
	goto L22
L22:
	;
	v80 = v78 & int32(_a_F__bt_killitems_1)
	v87 = v2
	v98 = v2
	goto L23
L23:
	;
	v100 = v87 + int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101+v87<<(uint(int32(2))%32))))
	v108 = v69 + v105*int32(10)
	v109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+6)))
	if base.B2i32(base.Ui32(v109) < base.Ui32(v65))|base.B2i32(base.Ui32(v80) < base.Ui32(v109)) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)))
	v348 = v346 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)) = uint16(v348)
	F_MarkBufferDirtyHint(m, v39, int32(1))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L6
	} else {
		goto L65
	}
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v276 | int32(_a_F__bt_killitems_2)
	if v21 != v100 {
		v87 = v100
		v98 = int32(1)
		goto L23
	} else {
		goto L64
	}
L27:
	;
	v122 = v109
	v124 = v108
	goto L30
L28:
	;
	goto L29
L29:
	;
	if v21 != v100 {
		v87 = v100
		goto L23
	} else {
		goto L62
	}
L30:
	;
	v137 = v57 + int32(20) + v122&int32(_a_F__bt_killitems_1)<<(uint(int32(2))%32)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v141 = v57 + v138&int32(_a_F__bt_killitems_3)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+7)))
	if v142&int32(32) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	goto L29
L32:
	;
	v300 = v122 + int32(1)
	if base.Ui32(v300&int32(_a_F__bt_killitems_1)) <= base.Ui32(v80) {
		v122 = v300
		v124 = v290
		goto L30
	} else {
		goto L61
	}
L33:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v277 = int32(_a_F__bt_killitems_2)
	if v276&v277 != v277 {
		goto L26
	} else {
		goto L60
	}
L34:
	;
	if v247 != v154 {
		v290 = v248
		goto L32
	} else {
		goto L59
	}
L35:
	;
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+2)))
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141))))
	v222 = int32(16)
	v225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+2)))
	v226 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v220|v221<<(uint(v222)%32) == v225|v226<<(uint(v222)%32) {
		goto L54
	} else {
		goto L55
	}
L36:
	;
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+4)))
	if v147&int32(_a_F__bt_killitems_4) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v152 = int32(0)
	v154 = v147 & int32(4095)
	if v154 == v152 {
		v247 = v152
		v248 = v124
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v157 = v100
	v165 = v152
	v166 = v124
	goto L39
L39:
	;
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+2)))
	v176 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141))))
	v177 = int32(16)
	v183 = v175 + (v141 + v176<<(uint(v177)%32)) + v165*int32(6)
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+2)))
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183))))
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+2)))
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166))))
	if v184|v185<<(uint(v177)%32) == v189|v190<<(uint(v177)%32) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v267 = v215
	goto L33
L41:
	;
	if v200 == int32(0) {
		v247 = v165
		v248 = v166
		goto L34
	} else {
		goto L47
	}
L42:
	;
	goto L41
L43:
	;
	v196 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+4)))
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+4)))
	if v196 == v197 {
		v200 = int32(1)
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v200 = int32(0)
	goto L42
L46:
	;
	goto L45
L47:
	;
	if v157 < v21 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v204+v157<<(uint(int32(2))%32))))
	v214 = v157 + int32(1)
	v215 = v69 + v208*int32(10)
	goto L50
L49:
	;
	v214 = v157
	v215 = v166
	goto L50
L50:
	;
	v217 = v165 + int32(1)
	if v217 != v154 {
		v157 = v214
		v165 = v217
		v166 = v215
		goto L39
	} else {
		goto L51
	}
L51:
	;
	goto L40
L52:
	;
	if v236 == int32(0) {
		v290 = v124
		goto L32
	} else {
		goto L58
	}
L53:
	;
	goto L52
L54:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+4)))
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	if v232 == v233 {
		v236 = int32(1)
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v236 = int32(0)
	goto L53
L57:
	;
	goto L56
L58:
	;
	v267 = v124
	goto L33
L59:
	;
	v267 = v248
	goto L33
L60:
	;
	v290 = v267
	goto L32
L61:
	;
	goto L31
L62:
	;
	if v98 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	goto L15
L64:
	;
	goto L25
L65:
	;
	goto L15
L66:
	;
	F__bt_unlockbuf(m, v39)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	return
L68:
	;
	return
}
func F__bt_mkscankey(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v118 int32
	_ = v118
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+10)))
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v64 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v64)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)) = uint16(v64)
	if v63 < v23 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v24&int32(32) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v56 = F_palloc(m, v23*int32(48)+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L12
	}
L5:
	;
	v42 = F_palloc(m, v23*int32(48)+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+8)))
	v37 = v35
	goto L5
L7:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v29&int32(_a_F__bt_mkscankey_0) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v37 = v29 & int32(4095)
	goto L5
L9:
	;
	return int32(0)
L10:
	;
	F__bt_metaversion(m, l0, v42, v42+int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v60 = v42
	v61 = v50
	v63 = v37
	goto L1
L12:
	;
	v58 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v56))) = uint16(v58)
	v60 = v56
	v61 = int32(1)
	v63 = int32(0)
	goto L1
L13:
	;
	v69 = v63
	goto L15
L14:
	;
	v69 = v23
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v69
	v71 = int32(0)
	if base.B2i32(l1 == v71)|base.B2i32(v61&int32(1) == v71) != 0 {
		v107 = v71
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v107
	if int32(0) < v23 {
		goto L23
	} else {
		goto L24
	}
L17:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v79&int32(_a_F__bt_mkscankey_0) == int32(0) {
		v107 = l1
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v84&int32(_a_F__bt_mkscankey_0) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v89 = int32(0)
	if v84&int32(_a_F__bt_mkscankey_1) == v89 {
		v107 = v89
		goto L16
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v107 = v99 + (l1 + v100<<(uint(int32(16))%32))
	goto L16
L22:
	;
	v107 = l1 + v79&int32(_a_F__bt_mkscankey_2) - int32(6)
	goto L16
L23:
	;
	v118 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+13)))
	if v194 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	v129 = int32(1)
	v131 = v118 + v129
	v132 = base.I32_extend16_s(v131)
	v134 = F_index_getprocinfo(m, l0, v132, v129)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L9
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	if v118 < v63 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v140 = F_index_getattr_2(m, l1, v131, v21, v18+int32(15))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L9
	} else {
		goto L32
	}
L30:
	;
	v143 = v129
	v144 = int32(0)
	goto L31
L31:
	;
	v147 = v60 + int32(16) + v118*int32(48)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v118<<(uint(int32(1))%32)))))
	v155 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v118<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+44)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v147)+8)) = v155
	*(*uint16)(unsafe.Add(mBase, uint32(v147)+6)) = uint16(v155)
	*(*uint16)(unsafe.Add(mBase, uint32(v147)+4)) = uint16(v132)
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = v143 | v151<<(uint(int32(24))%32)
	v171 = *(*int32)(unsafe.Add(mBase, _c_F__bt_mkscankey[0]))
	F_fmgr_info_copy(m, v147+int32(16), v134, v171)
	mBase = m.M
	goto L33
L32:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	v143 = v142
	v144 = v140
	goto L31
L33:
	;
	if v143&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v175)
	goto L36
L35:
	;
	goto L36
L36:
	;
	if v131 != v23 {
		v118 = v131
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v197)
	goto L40
L39:
	;
	goto L40
L40:
	;
	m.G0 = v18 + int32(16)
	return v60
}
func F__bt_parallel_seize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v218 int32
	_ = v218
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = int64(-4294967296)
	v25 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+88)) = uint16(v25)
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v15 + int32(16)
	return v218
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v33 = v17 + v32
	v37 = v33 + int32(40)
	v39 = v33 + int32(12)
	goto L7
L3:
	;
	v27 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+19)) = uint8(v27)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+17)) = uint16(v27)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+17)))
	if v31 != 0 {
		v218 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	v53 = F_LWLockAcquire(m, v39, int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return int32(0)
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v57 != int32(2) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v60 = int32(0)
	switch v57 - int32(1) {
	case 0:
		goto L17
	default:
		goto L16
	case 2:
		goto L18
	case 3:
		v173 = v60
		v174 = v5
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_LWLockRelease(m, v39)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L9
	} else {
		goto L50
	}
L14:
	;
	F_LWLockRelease(m, v39)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L9
	} else {
		goto L38
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v165
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v169
	v173 = int32(1)
	v174 = v5
	goto L14
L16:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v165 = v164
	goto L15
L17:
	;
	if l3 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v63 != 0 {
		v165 = v63
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v173 = v60
	v174 = int32(1)
	goto L14
L20:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+19)) = uint8(v156)
	v158 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+17)) = uint16(v158)
	F_LWLockRelease(m, v39)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L9
	} else {
		goto L36
	}
L21:
	;
	v67 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v37 + v69<<(uint(v67)%32)
	if v69 <= int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v78 = int32(0)
	goto L23
L23:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v93 = v90 + v78<<(uint(int32(5))%32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	v97 = v89 + v94*int32(48)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v98 != int32(-1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L20
L25:
	;
	v141 = v78 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	if v141 < v142 {
		v78 = v141
		goto L23
	} else {
		goto L35
	}
L26:
	;
	v101 = int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v37+v78<<(uint(v101)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v104
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v106+v104<<(uint(v101)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v97)+44)) = v110
	goto L25
L27:
	;
	goto L28
L28:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+18)))
	if v112 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+44)) = int32(0)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	*(*int32)(unsafe.Add(mBase, uint32(v97))) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v124 + int32(4)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
	if v128&int32(24) != 0 {
		goto L25
	} else {
		goto L33
	}
L30:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v97)+44))
	if v113 == int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	F_pfree(m, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L9
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v135 = F_datumRestore(m, v15+int32(12), v15+int32(11))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L9
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+44)) = v135
	goto L25
L35:
	;
	goto L24
L36:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	v218 = l3
	goto L1
L38:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L9
	} else {
		goto L39
	}
L39:
	;
	if v174 == int32(0) {
		v218 = v173
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v181 == int32(0) {
		v218 = v173
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+17)))
	if v185 != 0 {
		v218 = v173
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)+28))
	v187 = v181 + v186
	v189 = v187 + int32(12)
	v191 = F_LWLockAcquire(m, v189, int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L9
	} else {
		goto L43
	}
L43:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)+8))
	if v193 != int32(4) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v187)+8)) = int32(4)
	F_LWLockRelease(m, v189)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L9
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_LWLockRelease(m, v189)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L9
	} else {
		goto L49
	}
L47:
	;
	F_ConditionVariableBroadcast(m, v187+int32(28))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L9
	} else {
		goto L48
	}
L48:
	;
	v218 = v173
	goto L1
L49:
	;
	v218 = v173
	goto L1
L50:
	;
	F_ConditionVariableSleep(m, v33+int32(28), int32(134217735))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L51
	}
L51:
	;
	goto L7
}
func F__bt_readnextpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v82 int32
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
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v426 int64
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = l1
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l3 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+88)) = uint8(v25)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+89)) = uint8(v27)
	goto L1
L5:
	;
	m.G0 = v17 + int32(32)
	return v468
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(-4294967296)
	F__bt_parallel_done(m, l0)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L22
	} else {
		goto L124
	}
L7:
	;
	if l3 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v35 = int32(4)
	goto L10
L9:
	;
	v35 = int32(0)
	goto L10
L10:
	;
	v40 = l4
	goto L11
L11:
	;
	v51 = base.B2i32(l3 != int32(1))
	if v51 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v421 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L13:
	;
	v61 = v40 & int32(1)
	if v61 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+89)))
	if v54 == int32(0) {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+88)))
	if v57 == int32(0) {
		goto L6
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	goto L13
L19:
	;
	if v51 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v62 == int32(0) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v70 = F__bt_parallel_seize(m, l0, v17+int32(28), v17+int32(24), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return int32(0)
L23:
	;
	if v70 != 0 {
		goto L19
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(-4294967296)
	v468 = int32(0)
	goto L5
L25:
	;
	if v342 < int32(0) {
		goto L92
	} else {
		goto L93
	}
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[0]))
	if v80 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v96 = v89
	v98 = v88
	goto L36
L29:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L22
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v85 = F__bt_getbuf(m, v22, v83, int32(1))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L22
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v85
	v342 = v85
	goto L25
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L22
	} else {
		goto L88
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L22
	} else {
		goto L85
	}
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[0]))
	if v105 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	F__bt_relbuf(m, v287)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L22
	} else {
		goto L84
	}
L38:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L22
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v110 = F__bt_getbuf(m, v22, v108, int32(1))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L22
	} else {
		goto L43
	}
L41:
	;
	goto L40
L42:
	;
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v129)+16)))
	v131 = v130 + v129
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v133 = int32(0)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+12)))
	if base.B2i32(v134&int32(4) == v133)&base.B2i32(v132 == v96) == v133 {
		goto L48
	} else {
		goto L49
	}
L43:
	;
	if v110 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115+(v110^int32(-1))<<(uint(int32(2))%32))))
	v129 = v121
	goto L42
L45:
	;
	goto L46
L46:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v129 = v123 + v110<<(uint(int32(13))%32) + int32(-8192)
	goto L42
L47:
	;
	v212 = F__bt_relandgetbuf(m, v22, v147, v96, int32(1))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L22
	} else {
		goto L63
	}
L48:
	;
	v144 = v132
	v147 = v110
	v153 = v133
	goto L51
L49:
	;
	v198 = v110
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v198
	if v198 == int32(0) {
		goto L6
	} else {
		goto L60
	}
L51:
	;
	if base.B2i32(v144 == int32(0))|base.B2i32(v153 == int32(4)) != 0 {
		goto L47
	} else {
		goto L53
	}
L52:
	;
	v198 = v164
	goto L50
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v144
	v164 = F__bt_relandgetbuf(m, v22, v147, v144, int32(1))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L22
	} else {
		goto L55
	}
L54:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v187 = v183 + v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+12)))
	if v189&int32(4)|base.B2i32(v188 != v96) != 0 {
		v144 = v188
		v147 = v164
		v153 = v153 + int32(1)
		goto L51
	} else {
		goto L59
	}
L55:
	;
	if v164 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v169+(v164^int32(-1))<<(uint(int32(2))%32))))
	v183 = v175
	goto L54
L57:
	;
	goto L58
L58:
	;
	v177 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v183 = v177 + v164<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L59:
	;
	goto L52
L60:
	;
	v342 = v198
	goto L25
L61:
	;
	if v297 != 0 {
		goto L80
	} else {
		goto L81
	}
L62:
	;
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v231)+16)))
	v233 = v232 + v231
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+12)))
	if v234&int32(4) != 0 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	if v212 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217+(v212^int32(-1))<<(uint(int32(2))%32))))
	v231 = v223
	goto L62
L65:
	;
	goto L66
L66:
	;
	v225 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v231 = v225 + v212<<(uint(int32(13))%32) + int32(-8192)
	goto L62
L67:
	;
	v238 = v233
	v241 = v212
	goto L70
L68:
	;
	goto L69
L69:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v233)))
	if v281 == v98 {
		goto L34
	} else {
		goto L79
	}
L70:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	if v251 == int32(0) {
		goto L35
	} else {
		goto L72
	}
L71:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	v287 = v255
	v289 = v251
	v297 = v280
	goto L61
L72:
	;
	v255 = F__bt_relandgetbuf(m, v22, v241, v251, int32(1))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L22
	} else {
		goto L74
	}
L73:
	;
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274)+16)))
	v276 = v275 + v274
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v276)+12)))
	if v277&int32(4) != 0 {
		v238 = v276
		v241 = v255
		goto L70
	} else {
		goto L78
	}
L74:
	;
	if v255 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v260+(v255^int32(-1))<<(uint(int32(2))%32))))
	v274 = v266
	goto L73
L76:
	;
	goto L77
L77:
	;
	v268 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v274 = v268 + v255<<(uint(int32(13))%32) + int32(-8192)
	goto L73
L78:
	;
	goto L71
L79:
	;
	v287 = v212
	v289 = v96
	v297 = v281
	goto L61
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v297
	F__bt_relbuf(m, v287)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L22
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	goto L37
L83:
	;
	v96 = v289
	v98 = v297
	goto L36
L84:
	;
	goto L6
L85:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v307 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_readnextpage_0), v17+int32(16))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L22
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(_a_F__bt_readnextpage_1), int32(2563), int32(_a_F__bt_readnextpage_2))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L22
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v325 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_readnextpage_3), v17)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L22
	} else {
		goto L89
	}
L89:
	;
	F_errfinish(m, int32(_a_F__bt_readnextpage_1), int32(2581), int32(_a_F__bt_readnextpage_2))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L22
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
	v370 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369)+16)))
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v371
	v373 = v369 + v370
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+12)))
	if v374&int32(20) == int32(0) {
		goto L97
	} else {
		goto L98
	}
L92:
	;
	v355 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[1]))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v355+(v342^int32(-1))<<(uint(int32(2))%32))))
	v369 = v361
	goto L91
L93:
	;
	goto L94
L94:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F__bt_readnextpage[2]))
	v369 = v363 + v342<<(uint(int32(13))%32) + int32(-8192)
	goto L91
L95:
	;
	goto L12
L96:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	F__bt_relbuf(m, v414)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L22
	} else {
		goto L115
	}
L97:
	;
	if v51 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v373+v35)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v408 == int32(0) {
		goto L96
	} else {
		goto L113
	}
L100:
	;
	v381 = int32(1)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
	if v384 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	goto L102
L102:
	;
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v369)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v390) {
		goto L108
	} else {
		goto L109
	}
L103:
	;
	v385 = int32(2)
	goto L105
L104:
	;
	v385 = v381
	goto L105
L105:
	;
	v386 = F__bt_readpage(m, l0, v381, v385, v61)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L22
	} else {
		goto L106
	}
L106:
	;
	if v386 != 0 {
		goto L95
	} else {
		goto L107
	}
L107:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v388
	goto L96
L108:
	;
	v398 = int32(base.Ui32(v390+int32(_a_F__bt_readnextpage_4)) >> (uint(int32(2)) % 32))
	goto L110
L109:
	;
	v398 = int32(0)
	goto L110
L110:
	;
	v401 = F__bt_readpage(m, l0, l3, v398&int32(_a_F__bt_readnextpage_5), v61)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L22
	} else {
		goto L111
	}
L111:
	;
	if v401 != 0 {
		goto L95
	} else {
		goto L112
	}
L112:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v403
	goto L96
L113:
	;
	F__bt_parallel_release(m, l0, v406, v371)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L22
	} else {
		goto L114
	}
L114:
	;
	goto L96
L115:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v418 != 0 {
		v40 = int32(0)
		goto L11
	} else {
		goto L116
	}
L116:
	;
	goto L6
L117:
	;
	v468 = int32(1)
	goto L5
L118:
	;
	F__bt_unlockbuf(m, v420)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L22
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v426 = F_BufferGetLSNAtomic(m, v420)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L22
	} else {
		goto L122
	}
L121:
	;
	goto L117
L122:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = v426
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	F__bt_relbuf(m, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L22
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = int32(0)
	goto L117
L124:
	;
	v468 = int32(0)
	goto L5
}
func F__bt_recsplitloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
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
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	v2 = l1
	v5 = int32(0)
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v9 == v2 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = v11 + (l2 - v12)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v17&int32(1) != 0 {
			v65 = v16
			v66 = v14
			v68 = v15
			v70 = v65
			v71 = v66
			v72 = int32(-8)
			v73 = v68
			v75 = v70
			v76 = v71
			v77 = v72
			v78 = v73
			v79 = int32(1)
		} else {
			v75 = v16
			v76 = v14
			v77 = v5
			v78 = v15
			v79 = v5
		}
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if base.B2i32(v20 != int32(1))|base.B2i32(base.Ui32(l3) < base.Ui32(int32(65))) == int32(0) {
			v28 = int32(-8)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v2<<(uint(int32(2))%32))+20))
			v36 = v29 + v33&int32(_a_F__bt_recsplitloc_0)
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+6)))
			if v37&int32(_a_F__bt_recsplitloc_1) == int32(0) {
				v53 = v28
			} else {
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+5)))
				if v42&int32(32) == int32(0) {
					v53 = v28
				} else {
					v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+2)))
					v53 = v37&int32(_a_F__bt_recsplitloc_2) - v49 - int32(8)
				}
			}
			v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v70 = l3
			v71 = v54 + (l2 - v55)
			v72 = v53
			v73 = v58
			v75 = v70
			v76 = v71
			v77 = v72
			v78 = v73
			v79 = int32(1)
		} else {
			v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v62 = v59 + (l2 - v60)
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v20 != 0 {
				v65 = l3
				v66 = v62
				v68 = v63
				v70 = v65
				v71 = v66
				v72 = int32(-8)
				v73 = v68
				v75 = v70
				v76 = v71
				v77 = v72
				v78 = v73
				v79 = int32(1)
			} else {
				v75 = l3
				v76 = v62
				v77 = int32(0)
				v78 = v63
				v79 = v5
			}
		}
	}
	v82 = v77 + v78 - (l2 + v75)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v79 != 0 {
		v88 = int32(0)
	} else {
		v88 = v75 + int32(_a_F__bt_recsplitloc_3)
	}
	v89 = v76 - v83 + v88
	if (v82|v89)&int32(_a_F__bt_recsplitloc_4) == int32(0) {
		v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if base.Ui32(v95) < base.Ui32(v75) {
			v97 = v95
		} else {
			v97 = v75
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v97
		v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v101 = int32(10)
		v104 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v99+v100*v101))) = uint16(v104)
		v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v106+v107*v101)+2)) = uint16(v82)
		v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v112+v113*v101)+4)) = uint16(v89)
		v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v118+v119*v101)+6)) = uint16(v2)
		v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint8)(unsafe.Add(mBase, uint32(v124+v125*v101)+8)) = uint8(v104)
		v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v131 + int32(1)
	} else {
	}
	return
}
func F__bt_restore_meta(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_XLogInitBufferForRedo(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = v9 + int32(12)
		v16 = int32(0)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
		if v18 < l1 {
			v40 = v16
			v43 = v40
		} else {
			v22 = v17 + l1*int32(52)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+76)))
			if v23 != int32(1) {
				v40 = v16
				v43 = v40
			} else {
				v27 = v22 + int32(76)
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+43)))
				if v28 == int32(0) {
					if v15 == int32(0) {
						v40 = v16
						v43 = v40
					} else {
						v33 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v33
						v43 = v33
					}
				} else {
					if v15 != 0 {
						v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+48)))
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v36
					} else {
					}
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
					v40 = v38
					v43 = v40
				}
			}
		}
		if v12 < int32(0) {
			v47 = *(*int32)(unsafe.Add(mBase, _c_F__bt_restore_meta[0]))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v12^int32(-1))<<(uint(int32(2))%32))))
			v61 = v53
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, _c_F__bt_restore_meta[1]))
			v61 = v55 + v12<<(uint(int32(13))%32) + int32(-8192)
		}
		F_PageInit(m, v61, int32(_a_F__bt_restore_meta_0), int32(16))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = int32(_a_F__bt_restore_meta_1)
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+28)) = v67
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v69
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v71
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = v73
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = v75
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v61)+56)) = int64(-4616189618054758400)
		*(*int32)(unsafe.Add(mBase, uint32(v61)+48)) = v77
		v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
		*(*uint8)(unsafe.Add(mBase, uint32(v61)+64)) = uint8(v81)
		v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+16)))
		v85 = int32(8)
		*(*uint16)(unsafe.Add(mBase, uint32(v61+v83)+12)) = uint16(v85)
		*(*uint32)(unsafe.Add(mBase, uint32(v61)+4)) = uint32(v11)
		v89 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v61))) = uint32(v89)
		v91 = int32(72)
		*(*uint16)(unsafe.Add(mBase, uint32(v61)+12)) = uint16(v91)
		F_MarkBufferDirty(m, v12)
		mBase = m.M
		v94 = m.ExcPending
		if v94 != 0 {
			return
		} else {
			F_UnlockReleaseBuffer(m, v12)
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F__bt_restore_page(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(2448)
	m.G0 = v9
	if v4 < l2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = l1
	v18 = v4
	goto L4
L2:
	;
	v45 = v4
	goto L3
L3:
	;
	v48 = v45
	goto L8
L4:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(816)+v18<<(uint(int32(2))%32)))) = v15
	v27 = int32(1)
	v35 = (v20&int32(_a_F__bt_restore_page_0) + int32(7)) & int32(_a_F__bt_restore_page_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v9+v18<<(uint(v27)%32)))) = uint16(v35)
	v38 = v18 + v27
	v39 = v15 + v35
	if base.Ui32(v39) < base.Ui32(l1+l2) {
		v15 = v39
		v18 = v38
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v45 = v38
	goto L3
L6:
	;
	goto L5
L7:
	;
	m.G0 = v9 + int32(2448)
	return
L8:
	;
	v54 = v48 - int32(1)
	if v54 < int32(0) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L11
	} else {
		goto L14
	}
L10:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(816)+v54<<(uint(int32(2))%32))))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+v54<<(uint(int32(1))%32)))))
	v71 = F_PageAddItemExtended(m, l0, v62, v66, (v45-v54)&int32(_a_F__bt_restore_page_2), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	if v71 != 0 {
		v48 = v54
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	F_errmsg_internal(m, int32(_a_F__bt_restore_page_3), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(_a_F__bt_restore_page_4), int32(77), int32(_a_F__bt_restore_page_5))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_scanbehind_checkkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v16&int32(32) == v4 {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+192))
		v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
		v30 = v28
	} else {
		v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
		if v21&int32(_a_F__bt_scanbehind_checkkeys_0) != 0 {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)+192))
			v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
			v30 = v28
		} else {
			v30 = v21 & int32(4095)
		}
	}
	v31 = int32(0)
	v35 = F__bt_tuple_before_array_skeys(m, l0, l1, l2, v15, v30, v31, v31, v11+int32(7))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		if v35 != 0 {
			v89 = v4
			m.G0 = v11 + int32(16)
			return v89
		} else {
			v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
			if v39&int32(1) != 0 {
				v89 = v4
				m.G0 = v11 + int32(16)
				return v89
			} else {
				v42 = int32(1)
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+19)))
				if v43 != v42 {
					v89 = v42
					m.G0 = v11 + int32(16)
					return v89
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
					if v48&int32(32) == int32(0) {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+192))
						v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+8)))
						v62 = v60
					} else {
						v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
						if v53&int32(_a_F__bt_scanbehind_checkkeys_0) != 0 {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+192))
							v60 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59)+8)))
							v62 = v60
						} else {
							v62 = v53 & int32(4095)
						}
					}
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v64 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v64
					v75 = F__bt_check_compare(m, l0, v64-l1, l2, v62, v47, v64, v64, v11+int32(15), v11+int32(8))
					mBase = m.M
					v76 = m.ExcPending
					if v76 != 0 {
						return int32(0)
					} else {
						v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
						if v77 == int32(0) {
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+v81*int32(48))+6)))
							if v85 != int32(3) {
								v89 = v64
							} else {
								v89 = int32(1)
							}
						} else {
							v89 = int32(1)
						}
						m.G0 = v11 + int32(16)
						return v89
					}
				}
			}
		}
	}
}
func F__bt_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v15 = F__bt_getroot(m, l0, l1, l4)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v15
	if v15 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v25 = base.B2i32(l4 == int32(2))
	v33 = v15
	v34 = int32(1)
	v35 = int32(0)
	goto L6
L6:
	;
	v41 = F__bt_moveright(m, l0, l1, l2, v33, v25, v35, v34)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	if base.B2i32(l4 != int32(2))|base.B2i32(v34 != int32(1)) == int32(0) {
		goto L29
	} else {
		goto L30
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v41
	if v41 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+16)))
	v63 = v62 + v61
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+12)))
	if v64&int32(1) == int32(0) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F__bt_search[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v41^int32(-1))<<(uint(int32(2))%32))))
	v61 = v53
	goto L9
L11:
	;
	goto L12
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F__bt_search[1]))
	v61 = v55 + v41<<(uint(int32(13))%32) + int32(-8192)
	goto L9
L13:
	;
	v69 = F__bt_binsrch(m, l0, l2, v41)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	goto L7
L16:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v61+v69<<(uint(int32(2))%32))+20))
	v77 = v61 + v74&int32(_a_F__bt_search_0)
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77)+2)))
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v77))))
	v81 = F_palloc(m, int32(12))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v83 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v35
	*(*uint16)(unsafe.Add(mBase, uint32(v81)+4)) = uint16(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v81))) = v102
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	if v111 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F__bt_search[2]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v83^int32(-1))<<(uint(int32(6))%32))+16))
	v102 = v93
	goto L18
L20:
	;
	goto L21
L21:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F__bt_search[3]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+v83<<(uint(int32(6))%32)+int32(-64))+16))
	v102 = v101
	goto L18
L22:
	;
	v114 = int32(2)
	goto L24
L23:
	;
	v114 = v34
	goto L24
L24:
	;
	if l4 == int32(2) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v115 = v114
	goto L27
L26:
	;
	v115 = v34
	goto L27
L27:
	;
	v116 = F__bt_relandgetbuf(m, l0, v106, v78|v79<<(uint(int32(16))%32), v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v116
	v33 = v116
	v34 = v115
	v35 = v81
	goto L6
L29:
	;
	F__bt_unlockbuf(m, v41)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	return v35
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_LockBuffer(m, v128, int32(2))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v135 = F__bt_moveright(m, l0, l1, l2, v132, int32(1), v35, int32(2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v135
	goto L31
}
func F__bt_setup_array_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v19 = (v15 - int32(1)) << (uint(int32(2)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+212))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v21)))
	if v23 == l2 {
		v26 = F_index_getprocinfo(m, v20, v15, int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v28
			v30 = *(*int64)(unsafe.Add(mBase, uint32(v26)+16))
			*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v30
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v32
			v34 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = v34
			if l4 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = l3
			}
			m.G0 = v13 - int32(-64)
			return
		}
	} else {
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v40+v19)))
		v44 = F_get_opfamily_proc(m, v42, v23, l2, int32(1))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			if v44 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return
				} else {
					v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v81 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v80
					*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
					F_errmsg_internal(m, int32(_a_F__bt_setup_array_cmp_0), v13)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F__bt_setup_array_cmp_1), int32(2691), int32(_a_F__bt_setup_array_cmp_2))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
				F_fmgr_info_cxt(m, v44, l3, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					if l4 == int32(0) {
						m.G0 = v13 - int32(-64)
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
						v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54<<(uint(int32(2))%32)-int32(4))))
						v62 = F_get_opfamily_proc(m, v60, l2, l2, int32(1))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							if v62 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return
								} else {
									v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v103 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v102
									*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(1)
									F_errmsg_internal(m, int32(_a_F__bt_setup_array_cmp_0), v11+int32(-32))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F__bt_setup_array_cmp_1), int32(2713), int32(_a_F__bt_setup_array_cmp_2))
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
								F_fmgr_info_cxt(m, v62, v66, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									m.G0 = v13 - int32(-64)
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
func F__bt_upgrademetapage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	v2 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	v10 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v10)
	return
}
func F_bt_index_block_validate(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
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
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+119)))
	if v10 != int32(105) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_bt_index_block_validate_0)
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v36 + int32(4)
				F_errmsg(m, int32(_a_F_bt_index_block_validate_1), v7)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(231), int32(_a_F_bt_index_block_validate_3))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+84))
		if v13 != int32(403) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_bt_index_block_validate_0)
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = v36 + int32(4)
					F_errmsg(m, int32(_a_F_bt_index_block_validate_1), v7)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(231), int32(_a_F_bt_index_block_validate_3))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+118)))
			if v16 == int32(116) {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v19 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_bt_index_block_validate_4), int32(0))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(241), int32(_a_F_bt_index_block_validate_3))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if l1 == int64(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
							return
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_bt_index_block_validate_5), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(246), int32(_a_F_bt_index_block_validate_3))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						F_check_relation_block_range(m, l0, l1)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			} else {
				if l1 == int64(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_bt_index_block_validate_5), int32(0))
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_bt_index_block_validate_2), int32(246), int32(_a_F_bt_index_block_validate_3))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					F_check_relation_block_range(m, l0, l1)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_bt_multi_page_stats(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int64
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v221 int64
	_ = v221
	var v222 int64
	_ = v222
	var v225 int64
	_ = v225
	var v229 int64
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	v10 = m.G0
	v12 = v10 - int32(272)
	m.G0 = v12
	v14 = F_superuser(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
			if v19 == int32(0) {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v23 = F_pg_detoast_datum_packed(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v26 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v28 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
					v29 = F_init_MultiFuncCall(m, l0)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v31 = F_textToQualifiedNameList(m, v23)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = F_makeRangeVarFromNameList(m, v31)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								v36 = F_relation_openrv(m, v33, int32(1))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_bt_index_block_validate(m, v36, v28)
									mBase = m.M
									v39 = m.ExcPending
									if v39 != 0 {
										return int32(0)
									} else {
										if int64(2) <= v26 {
											F_check_relation_block_range(m, v36, v26+v28-int64(1))
											mBase = m.M
											v46 = m.ExcPending
											if v46 != 0 {
												return int32(0)
											} else {
												v47 = int32(_a_F_bt_multi_page_stats_0)
												v48 = *(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0]))
												v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
												*(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0])) = v50
												v53 = F_palloc(m, int32(32))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
												} else {
													v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
													v57 = int64(base.Ui64(v26) >> (uint(int64(63)) % 64))
													*(*uint8)(unsafe.Add(mBase, uint32(v53)+24)) = uint8(v57)
													*(*int64)(unsafe.Add(mBase, uint32(v53)+16)) = v26
													*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v28
													*(*int32)(unsafe.Add(mBase, uint32(v53))) = v55
													*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v53
													*(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0])) = v48
													F_relation_close(m, v36, int32(0))
													mBase = m.M
													v67 = m.ExcPending
													if v67 != 0 {
														return int32(0)
													} else {
														v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
														v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
														v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
														v80 = F_relation_open(m, v78, int32(0))
														mBase = m.M
														v81 = m.ExcPending
														if v81 != 0 {
															return int32(0)
														} else {
															v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+24)))
															if v82 == int32(0) {
																v85 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
																v93 = v85
																if int64(0) < v93 {
																	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
																	v97 = F_ReadBuffer(m, v80, v96)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return int32(0)
																	} else {
																		F_LockBuffer(m, v97, int32(1))
																		mBase = m.M
																		v101 = m.ExcPending
																		if v101 != 0 {
																			return int32(0)
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
																			v104 = int32(0)
																			*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v104)
																			*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
																			v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
																			F_GetBTPageStatistics(m, v108, v97, v12+int32(176))
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return int32(0)
																			} else {
																				F_UnlockReleaseBuffer(m, v97)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v80, int32(0))
																					mBase = m.M
																					v117 = m.ExcPending
																					if v117 != 0 {
																						return int32(0)
																					} else {
																						v121 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
																						mBase = m.M
																						v122 = m.ExcPending
																						if v122 != 0 {
																							return int32(0)
																						} else {
																							if v121 != int32(1) {
																								F_errstart_cold(m, int32(21), int32(0))
																								mBase = m.M
																								v272 = m.ExcPending
																								if v272 != 0 {
																									return int32(0)
																								} else {
																									F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																									mBase = m.M
																									v276 = m.ExcPending
																									if v276 != 0 {
																										return int32(0)
																									} else {
																										F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(437), int32(_a_F_bt_multi_page_stats_3))
																										mBase = m.M
																										v281 = m.ExcPending
																										if v281 != 0 {
																											return int32(0)
																										} else {
																											base.Wasm_trap_unreachable()
																											for {
																											}
																										}
																									}
																								}
																							} else {
																								v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v125
																								v130 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
																								mBase = m.M
																								v131 = m.ExcPending
																								if v131 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v130
																									v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v133
																									v138 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																									mBase = m.M
																									v139 = m.ExcPending
																									if v139 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v138
																										v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v141
																										v146 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																										mBase = m.M
																										v147 = m.ExcPending
																										if v147 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v146
																											v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v149
																											v154 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																											mBase = m.M
																											v155 = m.ExcPending
																											if v155 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v154
																												v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v157
																												v162 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																												mBase = m.M
																												v163 = m.ExcPending
																												if v163 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v162
																													v165 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v165
																													v170 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																													mBase = m.M
																													v171 = m.ExcPending
																													if v171 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v170
																														v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v173
																														v178 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																														mBase = m.M
																														v179 = m.ExcPending
																														if v179 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v178
																															v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v181
																															v186 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																															mBase = m.M
																															v187 = m.ExcPending
																															if v187 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v186
																																v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v189
																																v194 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																																mBase = m.M
																																v195 = m.ExcPending
																																if v195 != 0 {
																																	return int32(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v194
																																	v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v197
																																	v202 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																																	mBase = m.M
																																	v203 = m.ExcPending
																																	if v203 != 0 {
																																		return int32(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v202
																																		v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v205
																																		v208 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																																		mBase = m.M
																																		v209 = m.ExcPending
																																		if v209 != 0 {
																																			return int32(0)
																																		} else {
																																			*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v208
																																			v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																																			v212 = F_TupleDescGetAttInMetadata(m, v211)
																																			mBase = m.M
																																			v213 = m.ExcPending
																																			if v213 != 0 {
																																				return int32(0)
																																			} else {
																																				v216 = F_BuildTupleFromCStrings(m, v212, v12+int32(224))
																																				mBase = m.M
																																				v217 = m.ExcPending
																																				if v217 != 0 {
																																					return int32(0)
																																				} else {
																																					v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
																																					v219 = F_HeapTupleHeaderGetDatum(m, v218)
																																					mBase = m.M
																																					v220 = m.ExcPending
																																					if v220 != 0 {
																																						return int32(0)
																																					} else {
																																						v221 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
																																						v222 = int64(1)
																																						*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v221 + v222
																																						v225 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
																																						*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v225 - v222
																																						v229 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
																																						*(*int64)(unsafe.Add(mBase, uint32(v76))) = v229 + v222
																																						v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																																						*(*int32)(unsafe.Add(mBase, uint32(v233)+20)) = int32(1)
																																						v247 = v219
																																						m.G0 = v12 + int32(272)
																																						return v247
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
																	F_relation_close(m, v80, int32(1))
																	mBase = m.M
																	v238 = m.ExcPending
																	if v238 != 0 {
																		return int32(0)
																	} else {
																		F_end_MultiFuncCall(m, l0)
																		mBase = m.M
																		v240 = m.ExcPending
																		if v240 != 0 {
																			return int32(0)
																		} else {
																			v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(2)
																			v244 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v244)
																			v247 = int32(0)
																			m.G0 = v12 + int32(272)
																			return v247
																		}
																	}
																}
															} else {
																v87 = F_RelationGetNumberOfBlocksInFork(m, v80, int32(0))
																mBase = m.M
																v88 = m.ExcPending
																if v88 != 0 {
																	return int32(0)
																} else {
																	v90 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
																	v91 = base.I64_extend_i32_u(v87) - v90
																	*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v91
																	v93 = v91
																	if int64(0) < v93 {
																		v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
																		v97 = F_ReadBuffer(m, v80, v96)
																		mBase = m.M
																		v98 = m.ExcPending
																		if v98 != 0 {
																			return int32(0)
																		} else {
																			F_LockBuffer(m, v97, int32(1))
																			mBase = m.M
																			v101 = m.ExcPending
																			if v101 != 0 {
																				return int32(0)
																			} else {
																				*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
																				v104 = int32(0)
																				*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v104)
																				*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
																				v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
																				F_GetBTPageStatistics(m, v108, v97, v12+int32(176))
																				mBase = m.M
																				v112 = m.ExcPending
																				if v112 != 0 {
																					return int32(0)
																				} else {
																					F_UnlockReleaseBuffer(m, v97)
																					mBase = m.M
																					v114 = m.ExcPending
																					if v114 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v80, int32(0))
																						mBase = m.M
																						v117 = m.ExcPending
																						if v117 != 0 {
																							return int32(0)
																						} else {
																							v121 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
																							mBase = m.M
																							v122 = m.ExcPending
																							if v122 != 0 {
																								return int32(0)
																							} else {
																								if v121 != int32(1) {
																									F_errstart_cold(m, int32(21), int32(0))
																									mBase = m.M
																									v272 = m.ExcPending
																									if v272 != 0 {
																										return int32(0)
																									} else {
																										F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																										mBase = m.M
																										v276 = m.ExcPending
																										if v276 != 0 {
																											return int32(0)
																										} else {
																											F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(437), int32(_a_F_bt_multi_page_stats_3))
																											mBase = m.M
																											v281 = m.ExcPending
																											if v281 != 0 {
																												return int32(0)
																											} else {
																												base.Wasm_trap_unreachable()
																												for {
																												}
																											}
																										}
																									}
																								} else {
																									v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v125
																									v130 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
																									mBase = m.M
																									v131 = m.ExcPending
																									if v131 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v130
																										v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v133
																										v138 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																										mBase = m.M
																										v139 = m.ExcPending
																										if v139 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v138
																											v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v141
																											v146 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																											mBase = m.M
																											v147 = m.ExcPending
																											if v147 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v146
																												v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v149
																												v154 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																												mBase = m.M
																												v155 = m.ExcPending
																												if v155 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v154
																													v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v157
																													v162 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																													mBase = m.M
																													v163 = m.ExcPending
																													if v163 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v162
																														v165 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v165
																														v170 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																														mBase = m.M
																														v171 = m.ExcPending
																														if v171 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v170
																															v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v173
																															v178 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																															mBase = m.M
																															v179 = m.ExcPending
																															if v179 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v178
																																v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v181
																																v186 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																																mBase = m.M
																																v187 = m.ExcPending
																																if v187 != 0 {
																																	return int32(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v186
																																	v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v189
																																	v194 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																																	mBase = m.M
																																	v195 = m.ExcPending
																																	if v195 != 0 {
																																		return int32(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v194
																																		v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v197
																																		v202 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																																		mBase = m.M
																																		v203 = m.ExcPending
																																		if v203 != 0 {
																																			return int32(0)
																																		} else {
																																			*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v202
																																			v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																																			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v205
																																			v208 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																																			mBase = m.M
																																			v209 = m.ExcPending
																																			if v209 != 0 {
																																				return int32(0)
																																			} else {
																																				*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v208
																																				v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																																				v212 = F_TupleDescGetAttInMetadata(m, v211)
																																				mBase = m.M
																																				v213 = m.ExcPending
																																				if v213 != 0 {
																																					return int32(0)
																																				} else {
																																					v216 = F_BuildTupleFromCStrings(m, v212, v12+int32(224))
																																					mBase = m.M
																																					v217 = m.ExcPending
																																					if v217 != 0 {
																																						return int32(0)
																																					} else {
																																						v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
																																						v219 = F_HeapTupleHeaderGetDatum(m, v218)
																																						mBase = m.M
																																						v220 = m.ExcPending
																																						if v220 != 0 {
																																							return int32(0)
																																						} else {
																																							v221 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
																																							v222 = int64(1)
																																							*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v221 + v222
																																							v225 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
																																							*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v225 - v222
																																							v229 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
																																							*(*int64)(unsafe.Add(mBase, uint32(v76))) = v229 + v222
																																							v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																																							*(*int32)(unsafe.Add(mBase, uint32(v233)+20)) = int32(1)
																																							v247 = v219
																																							m.G0 = v12 + int32(272)
																																							return v247
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
																		F_relation_close(m, v80, int32(1))
																		mBase = m.M
																		v238 = m.ExcPending
																		if v238 != 0 {
																			return int32(0)
																		} else {
																			F_end_MultiFuncCall(m, l0)
																			mBase = m.M
																			v240 = m.ExcPending
																			if v240 != 0 {
																				return int32(0)
																			} else {
																				v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																				*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(2)
																				v244 = int32(1)
																				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v244)
																				v247 = int32(0)
																				m.G0 = v12 + int32(272)
																				return v247
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
											v47 = int32(_a_F_bt_multi_page_stats_0)
											v48 = *(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0]))
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+24))
											*(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0])) = v50
											v53 = F_palloc(m, int32(32))
											mBase = m.M
											v54 = m.ExcPending
											if v54 != 0 {
												return int32(0)
											} else {
												v55 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
												v57 = int64(base.Ui64(v26) >> (uint(int64(63)) % 64))
												*(*uint8)(unsafe.Add(mBase, uint32(v53)+24)) = uint8(v57)
												*(*int64)(unsafe.Add(mBase, uint32(v53)+16)) = v26
												*(*int64)(unsafe.Add(mBase, uint32(v53)+8)) = v28
												*(*int32)(unsafe.Add(mBase, uint32(v53))) = v55
												*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = v53
												*(*int32)(unsafe.Add(mBase, _c_F_bt_multi_page_stats[0])) = v48
												F_relation_close(m, v36, int32(0))
												mBase = m.M
												v67 = m.ExcPending
												if v67 != 0 {
													return int32(0)
												} else {
													v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
													v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
													v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
													v80 = F_relation_open(m, v78, int32(0))
													mBase = m.M
													v81 = m.ExcPending
													if v81 != 0 {
														return int32(0)
													} else {
														v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+24)))
														if v82 == int32(0) {
															v85 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
															v93 = v85
															if int64(0) < v93 {
																v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
																v97 = F_ReadBuffer(m, v80, v96)
																mBase = m.M
																v98 = m.ExcPending
																if v98 != 0 {
																	return int32(0)
																} else {
																	F_LockBuffer(m, v97, int32(1))
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
																		return int32(0)
																	} else {
																		*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
																		v104 = int32(0)
																		*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v104)
																		*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
																		v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
																		F_GetBTPageStatistics(m, v108, v97, v12+int32(176))
																		mBase = m.M
																		v112 = m.ExcPending
																		if v112 != 0 {
																			return int32(0)
																		} else {
																			F_UnlockReleaseBuffer(m, v97)
																			mBase = m.M
																			v114 = m.ExcPending
																			if v114 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v80, int32(0))
																				mBase = m.M
																				v117 = m.ExcPending
																				if v117 != 0 {
																					return int32(0)
																				} else {
																					v121 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
																					mBase = m.M
																					v122 = m.ExcPending
																					if v122 != 0 {
																						return int32(0)
																					} else {
																						if v121 != int32(1) {
																							F_errstart_cold(m, int32(21), int32(0))
																							mBase = m.M
																							v272 = m.ExcPending
																							if v272 != 0 {
																								return int32(0)
																							} else {
																								F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																								mBase = m.M
																								v276 = m.ExcPending
																								if v276 != 0 {
																									return int32(0)
																								} else {
																									F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(437), int32(_a_F_bt_multi_page_stats_3))
																									mBase = m.M
																									v281 = m.ExcPending
																									if v281 != 0 {
																										return int32(0)
																									} else {
																										base.Wasm_trap_unreachable()
																										for {
																										}
																									}
																								}
																							}
																						} else {
																							v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v125
																							v130 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
																							mBase = m.M
																							v131 = m.ExcPending
																							if v131 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v130
																								v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v133
																								v138 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																								mBase = m.M
																								v139 = m.ExcPending
																								if v139 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v138
																									v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v141
																									v146 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																									mBase = m.M
																									v147 = m.ExcPending
																									if v147 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v146
																										v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v149
																										v154 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																										mBase = m.M
																										v155 = m.ExcPending
																										if v155 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v154
																											v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v157
																											v162 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																											mBase = m.M
																											v163 = m.ExcPending
																											if v163 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v162
																												v165 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v165
																												v170 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																												mBase = m.M
																												v171 = m.ExcPending
																												if v171 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v170
																													v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v173
																													v178 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																													mBase = m.M
																													v179 = m.ExcPending
																													if v179 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v178
																														v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v181
																														v186 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																														mBase = m.M
																														v187 = m.ExcPending
																														if v187 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v186
																															v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v189
																															v194 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																															mBase = m.M
																															v195 = m.ExcPending
																															if v195 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v194
																																v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v197
																																v202 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																																mBase = m.M
																																v203 = m.ExcPending
																																if v203 != 0 {
																																	return int32(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v202
																																	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																																	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v205
																																	v208 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																																	mBase = m.M
																																	v209 = m.ExcPending
																																	if v209 != 0 {
																																		return int32(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v208
																																		v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																																		v212 = F_TupleDescGetAttInMetadata(m, v211)
																																		mBase = m.M
																																		v213 = m.ExcPending
																																		if v213 != 0 {
																																			return int32(0)
																																		} else {
																																			v216 = F_BuildTupleFromCStrings(m, v212, v12+int32(224))
																																			mBase = m.M
																																			v217 = m.ExcPending
																																			if v217 != 0 {
																																				return int32(0)
																																			} else {
																																				v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
																																				v219 = F_HeapTupleHeaderGetDatum(m, v218)
																																				mBase = m.M
																																				v220 = m.ExcPending
																																				if v220 != 0 {
																																					return int32(0)
																																				} else {
																																					v221 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
																																					v222 = int64(1)
																																					*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v221 + v222
																																					v225 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
																																					*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v225 - v222
																																					v229 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
																																					*(*int64)(unsafe.Add(mBase, uint32(v76))) = v229 + v222
																																					v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																																					*(*int32)(unsafe.Add(mBase, uint32(v233)+20)) = int32(1)
																																					v247 = v219
																																					m.G0 = v12 + int32(272)
																																					return v247
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
																F_relation_close(m, v80, int32(1))
																mBase = m.M
																v238 = m.ExcPending
																if v238 != 0 {
																	return int32(0)
																} else {
																	F_end_MultiFuncCall(m, l0)
																	mBase = m.M
																	v240 = m.ExcPending
																	if v240 != 0 {
																		return int32(0)
																	} else {
																		v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																		*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(2)
																		v244 = int32(1)
																		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v244)
																		v247 = int32(0)
																		m.G0 = v12 + int32(272)
																		return v247
																	}
																}
															}
														} else {
															v87 = F_RelationGetNumberOfBlocksInFork(m, v80, int32(0))
															mBase = m.M
															v88 = m.ExcPending
															if v88 != 0 {
																return int32(0)
															} else {
																v90 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
																v91 = base.I64_extend_i32_u(v87) - v90
																*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v91
																v93 = v91
																if int64(0) < v93 {
																	v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
																	v97 = F_ReadBuffer(m, v80, v96)
																	mBase = m.M
																	v98 = m.ExcPending
																	if v98 != 0 {
																		return int32(0)
																	} else {
																		F_LockBuffer(m, v97, int32(1))
																		mBase = m.M
																		v101 = m.ExcPending
																		if v101 != 0 {
																			return int32(0)
																		} else {
																			*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
																			v104 = int32(0)
																			*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v104)
																			*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
																			v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
																			F_GetBTPageStatistics(m, v108, v97, v12+int32(176))
																			mBase = m.M
																			v112 = m.ExcPending
																			if v112 != 0 {
																				return int32(0)
																			} else {
																				F_UnlockReleaseBuffer(m, v97)
																				mBase = m.M
																				v114 = m.ExcPending
																				if v114 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v80, int32(0))
																					mBase = m.M
																					v117 = m.ExcPending
																					if v117 != 0 {
																						return int32(0)
																					} else {
																						v121 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
																						mBase = m.M
																						v122 = m.ExcPending
																						if v122 != 0 {
																							return int32(0)
																						} else {
																							if v121 != int32(1) {
																								F_errstart_cold(m, int32(21), int32(0))
																								mBase = m.M
																								v272 = m.ExcPending
																								if v272 != 0 {
																									return int32(0)
																								} else {
																									F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																									mBase = m.M
																									v276 = m.ExcPending
																									if v276 != 0 {
																										return int32(0)
																									} else {
																										F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(437), int32(_a_F_bt_multi_page_stats_3))
																										mBase = m.M
																										v281 = m.ExcPending
																										if v281 != 0 {
																											return int32(0)
																										} else {
																											base.Wasm_trap_unreachable()
																											for {
																											}
																										}
																									}
																								}
																							} else {
																								v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v125
																								v130 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
																								mBase = m.M
																								v131 = m.ExcPending
																								if v131 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v130
																									v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v133
																									v138 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																									mBase = m.M
																									v139 = m.ExcPending
																									if v139 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v138
																										v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v141
																										v146 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																										mBase = m.M
																										v147 = m.ExcPending
																										if v147 != 0 {
																											return int32(0)
																										} else {
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v146
																											v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																											*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v149
																											v154 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																											mBase = m.M
																											v155 = m.ExcPending
																											if v155 != 0 {
																												return int32(0)
																											} else {
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v154
																												v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																												*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v157
																												v162 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																												mBase = m.M
																												v163 = m.ExcPending
																												if v163 != 0 {
																													return int32(0)
																												} else {
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v162
																													v165 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																													*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v165
																													v170 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																													mBase = m.M
																													v171 = m.ExcPending
																													if v171 != 0 {
																														return int32(0)
																													} else {
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v170
																														v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																														*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v173
																														v178 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																														mBase = m.M
																														v179 = m.ExcPending
																														if v179 != 0 {
																															return int32(0)
																														} else {
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v178
																															v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																															*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v181
																															v186 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																															mBase = m.M
																															v187 = m.ExcPending
																															if v187 != 0 {
																																return int32(0)
																															} else {
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v186
																																v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																																*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v189
																																v194 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																																mBase = m.M
																																v195 = m.ExcPending
																																if v195 != 0 {
																																	return int32(0)
																																} else {
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v194
																																	v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																																	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v197
																																	v202 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																																	mBase = m.M
																																	v203 = m.ExcPending
																																	if v203 != 0 {
																																		return int32(0)
																																	} else {
																																		*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v202
																																		v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																																		*(*int32)(unsafe.Add(mBase, uint32(v12))) = v205
																																		v208 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																																		mBase = m.M
																																		v209 = m.ExcPending
																																		if v209 != 0 {
																																			return int32(0)
																																		} else {
																																			*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v208
																																			v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																																			v212 = F_TupleDescGetAttInMetadata(m, v211)
																																			mBase = m.M
																																			v213 = m.ExcPending
																																			if v213 != 0 {
																																				return int32(0)
																																			} else {
																																				v216 = F_BuildTupleFromCStrings(m, v212, v12+int32(224))
																																				mBase = m.M
																																				v217 = m.ExcPending
																																				if v217 != 0 {
																																					return int32(0)
																																				} else {
																																					v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
																																					v219 = F_HeapTupleHeaderGetDatum(m, v218)
																																					mBase = m.M
																																					v220 = m.ExcPending
																																					if v220 != 0 {
																																						return int32(0)
																																					} else {
																																						v221 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
																																						v222 = int64(1)
																																						*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v221 + v222
																																						v225 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
																																						*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v225 - v222
																																						v229 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
																																						*(*int64)(unsafe.Add(mBase, uint32(v76))) = v229 + v222
																																						v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																																						*(*int32)(unsafe.Add(mBase, uint32(v233)+20)) = int32(1)
																																						v247 = v219
																																						m.G0 = v12 + int32(272)
																																						return v247
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
																	F_relation_close(m, v80, int32(1))
																	mBase = m.M
																	v238 = m.ExcPending
																	if v238 != 0 {
																		return int32(0)
																	} else {
																		F_end_MultiFuncCall(m, l0)
																		mBase = m.M
																		v240 = m.ExcPending
																		if v240 != 0 {
																			return int32(0)
																		} else {
																			v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																			*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(2)
																			v244 = int32(1)
																			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v244)
																			v247 = int32(0)
																			m.G0 = v12 + int32(272)
																			return v247
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
				v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+16))
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+16))
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
				v80 = F_relation_open(m, v78, int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+24)))
					if v82 == int32(0) {
						v85 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
						v93 = v85
						if int64(0) < v93 {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
							v97 = F_ReadBuffer(m, v80, v96)
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								F_LockBuffer(m, v97, int32(1))
								mBase = m.M
								v101 = m.ExcPending
								if v101 != 0 {
									return int32(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
									v104 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v104)
									*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
									F_GetBTPageStatistics(m, v108, v97, v12+int32(176))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return int32(0)
									} else {
										F_UnlockReleaseBuffer(m, v97)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											F_relation_close(m, v80, int32(0))
											mBase = m.M
											v117 = m.ExcPending
											if v117 != 0 {
												return int32(0)
											} else {
												v121 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
												mBase = m.M
												v122 = m.ExcPending
												if v122 != 0 {
													return int32(0)
												} else {
													if v121 != int32(1) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v272 = m.ExcPending
														if v272 != 0 {
															return int32(0)
														} else {
															F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
															mBase = m.M
															v276 = m.ExcPending
															if v276 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(437), int32(_a_F_bt_multi_page_stats_3))
																mBase = m.M
																v281 = m.ExcPending
																if v281 != 0 {
																	return int32(0)
																} else {
																	base.Wasm_trap_unreachable()
																	for {
																	}
																}
															}
														}
													} else {
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
														*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v125
														v130 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v130
															v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v133
															v138 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v138
																v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v141
																v146 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																mBase = m.M
																v147 = m.ExcPending
																if v147 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v146
																	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v149
																	v154 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																	mBase = m.M
																	v155 = m.ExcPending
																	if v155 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v154
																		v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v157
																		v162 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																		mBase = m.M
																		v163 = m.ExcPending
																		if v163 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v162
																			v165 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v165
																			v170 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																			mBase = m.M
																			v171 = m.ExcPending
																			if v171 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v170
																				v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v173
																				v178 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																				mBase = m.M
																				v179 = m.ExcPending
																				if v179 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v178
																					v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v181
																					v186 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																					mBase = m.M
																					v187 = m.ExcPending
																					if v187 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v186
																						v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v189
																						v194 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																						mBase = m.M
																						v195 = m.ExcPending
																						if v195 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v194
																							v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v197
																							v202 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																							mBase = m.M
																							v203 = m.ExcPending
																							if v203 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v202
																								v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																								*(*int32)(unsafe.Add(mBase, uint32(v12))) = v205
																								v208 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																								mBase = m.M
																								v209 = m.ExcPending
																								if v209 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v208
																									v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																									v212 = F_TupleDescGetAttInMetadata(m, v211)
																									mBase = m.M
																									v213 = m.ExcPending
																									if v213 != 0 {
																										return int32(0)
																									} else {
																										v216 = F_BuildTupleFromCStrings(m, v212, v12+int32(224))
																										mBase = m.M
																										v217 = m.ExcPending
																										if v217 != 0 {
																											return int32(0)
																										} else {
																											v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
																											v219 = F_HeapTupleHeaderGetDatum(m, v218)
																											mBase = m.M
																											v220 = m.ExcPending
																											if v220 != 0 {
																												return int32(0)
																											} else {
																												v221 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
																												v222 = int64(1)
																												*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v221 + v222
																												v225 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
																												*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v225 - v222
																												v229 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
																												*(*int64)(unsafe.Add(mBase, uint32(v76))) = v229 + v222
																												v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																												*(*int32)(unsafe.Add(mBase, uint32(v233)+20)) = int32(1)
																												v247 = v219
																												m.G0 = v12 + int32(272)
																												return v247
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
							F_relation_close(m, v80, int32(1))
							mBase = m.M
							v238 = m.ExcPending
							if v238 != 0 {
								return int32(0)
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v240 = m.ExcPending
								if v240 != 0 {
									return int32(0)
								} else {
									v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(2)
									v244 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v244)
									v247 = int32(0)
									m.G0 = v12 + int32(272)
									return v247
								}
							}
						}
					} else {
						v87 = F_RelationGetNumberOfBlocksInFork(m, v80, int32(0))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int32(0)
						} else {
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
							v91 = base.I64_extend_i32_u(v87) - v90
							*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v91
							v93 = v91
							if int64(0) < v93 {
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
								v97 = F_ReadBuffer(m, v80, v96)
								mBase = m.M
								v98 = m.ExcPending
								if v98 != 0 {
									return int32(0)
								} else {
									F_LockBuffer(m, v97, int32(1))
									mBase = m.M
									v101 = m.ExcPending
									if v101 != 0 {
										return int32(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v12)+208)) = int64(-1)
										v104 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)) = uint16(v104)
										*(*int64)(unsafe.Add(mBase, uint32(v12)+196)) = int64(0)
										v108 = *(*int32)(unsafe.Add(mBase, uint32(v77)+8))
										F_GetBTPageStatistics(m, v108, v97, v12+int32(176))
										mBase = m.M
										v112 = m.ExcPending
										if v112 != 0 {
											return int32(0)
										} else {
											F_UnlockReleaseBuffer(m, v97)
											mBase = m.M
											v114 = m.ExcPending
											if v114 != 0 {
												return int32(0)
											} else {
												F_relation_close(m, v80, int32(0))
												mBase = m.M
												v117 = m.ExcPending
												if v117 != 0 {
													return int32(0)
												} else {
													v121 = F_get_call_result_type(m, l0, int32(0), v12+int32(172))
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int32(0)
													} else {
														if v121 != int32(1) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v272 = m.ExcPending
															if v272 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_bt_multi_page_stats_1), int32(0))
																mBase = m.M
																v276 = m.ExcPending
																if v276 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(437), int32(_a_F_bt_multi_page_stats_3))
																	mBase = m.M
																	v281 = m.ExcPending
																	if v281 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v125 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
															*(*int32)(unsafe.Add(mBase, uint32(v12)+160)) = v125
															v130 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(160))
															mBase = m.M
															v131 = m.ExcPending
															if v131 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v12)+224)) = v130
																v133 = int32(*(*int8)(unsafe.Add(mBase, uint32(v12)+204)))
																*(*int32)(unsafe.Add(mBase, uint32(v12)+144)) = v133
																v138 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_5), v12+int32(144))
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+228)) = v138
																	v141 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
																	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = v141
																	v146 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(128))
																	mBase = m.M
																	v147 = m.ExcPending
																	if v147 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+232)) = v146
																		v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = v149
																		v154 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(112))
																		mBase = m.M
																		v155 = m.ExcPending
																		if v155 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+236)) = v154
																			v157 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
																			*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = v157
																			v162 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(96))
																			mBase = m.M
																			v163 = m.ExcPending
																			if v163 != 0 {
																				return int32(0)
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+240)) = v162
																				v165 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
																				*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v165
																				v170 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(80))
																				mBase = m.M
																				v171 = m.ExcPending
																				if v171 != 0 {
																					return int32(0)
																				} else {
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+244)) = v170
																					v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
																					*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v173
																					v178 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12-int32(-64))
																					mBase = m.M
																					v179 = m.ExcPending
																					if v179 != 0 {
																						return int32(0)
																					} else {
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+248)) = v178
																						v181 = *(*int32)(unsafe.Add(mBase, uint32(v12)+208))
																						*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v181
																						v186 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(48))
																						mBase = m.M
																						v187 = m.ExcPending
																						if v187 != 0 {
																							return int32(0)
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+252)) = v186
																							v189 = *(*int32)(unsafe.Add(mBase, uint32(v12)+212))
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v189
																							v194 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(32))
																							mBase = m.M
																							v195 = m.ExcPending
																							if v195 != 0 {
																								return int32(0)
																							} else {
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+256)) = v194
																								v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+216))
																								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v197
																								v202 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_4), v12+int32(16))
																								mBase = m.M
																								v203 = m.ExcPending
																								if v203 != 0 {
																									return int32(0)
																								} else {
																									*(*int32)(unsafe.Add(mBase, uint32(v12)+260)) = v202
																									v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+220)))
																									*(*int32)(unsafe.Add(mBase, uint32(v12))) = v205
																									v208 = F_psprintf(m, int32(_a_F_bt_multi_page_stats_6), v12)
																									mBase = m.M
																									v209 = m.ExcPending
																									if v209 != 0 {
																										return int32(0)
																									} else {
																										*(*int32)(unsafe.Add(mBase, uint32(v12)+264)) = v208
																										v211 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
																										v212 = F_TupleDescGetAttInMetadata(m, v211)
																										mBase = m.M
																										v213 = m.ExcPending
																										if v213 != 0 {
																											return int32(0)
																										} else {
																											v216 = F_BuildTupleFromCStrings(m, v212, v12+int32(224))
																											mBase = m.M
																											v217 = m.ExcPending
																											if v217 != 0 {
																												return int32(0)
																											} else {
																												v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)+16))
																												v219 = F_HeapTupleHeaderGetDatum(m, v218)
																												mBase = m.M
																												v220 = m.ExcPending
																												if v220 != 0 {
																													return int32(0)
																												} else {
																													v221 = *(*int64)(unsafe.Add(mBase, uint32(v77)+8))
																													v222 = int64(1)
																													*(*int64)(unsafe.Add(mBase, uint32(v77)+8)) = v221 + v222
																													v225 = *(*int64)(unsafe.Add(mBase, uint32(v77)+16))
																													*(*int64)(unsafe.Add(mBase, uint32(v77)+16)) = v225 - v222
																													v229 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
																													*(*int64)(unsafe.Add(mBase, uint32(v76))) = v229 + v222
																													v233 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																													*(*int32)(unsafe.Add(mBase, uint32(v233)+20)) = int32(1)
																													v247 = v219
																													m.G0 = v12 + int32(272)
																													return v247
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
								F_relation_close(m, v80, int32(1))
								mBase = m.M
								v238 = m.ExcPending
								if v238 != 0 {
									return int32(0)
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v240 = m.ExcPending
									if v240 != 0 {
										return int32(0)
									} else {
										v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v241)+20)) = int32(2)
										v244 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v244)
										v247 = int32(0)
										m.G0 = v12 + int32(272)
										return v247
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
			v256 = m.ExcPending
			if v256 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v259 = m.ExcPending
				if v259 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_bt_multi_page_stats_7), int32(0))
					mBase = m.M
					v263 = m.ExcPending
					if v263 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bt_multi_page_stats_2), int32(353), int32(_a_F_bt_multi_page_stats_3))
						mBase = m.M
						v268 = m.ExcPending
						if v268 != 0 {
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
